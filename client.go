package pango

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/plugin"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

type LoggingInfo struct {
	SLogHandler   slog.Handler `json:"-"`
	LogCategories LogCategory  `json:"-"`
	LogSymbols    []string     `json:"log_symbols"`
	LogLevel      slog.Level   `json:"log_level"`
}

// Client is an XML API client connection.  If provides wrapper functions
// for invoking the various PAN-OS API methods.  After creating the client,
// invoke Setup() followed by Initialize() to prepare it for use.
type Client struct {
	Hostname        string            `json:"hostname"`
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	ApiKey          string            `json:"api_key"`
	Protocol        string            `json:"protocol"`
	Port            int               `json:"port"`
	Target          string            `json:"target"`
	ApiKeyInRequest bool              `json:"api_key_in_request"`
	Headers         map[string]string `json:"headers"`
	Logging         LoggingInfo       `json:"logging"`

	// Set to true if you want to check environment variables
	// for auth and connection properties.
	CheckEnvironment bool   `json:"-"`
	AuthFile         string `json:"-"`

	// HTTP transport options.  Note that the SkipVerifyCertificate setting
	// is only used if you do not specify a HTTP transport yourself.
	SkipVerifyCertificate bool            `json:"skip_verify_certificate"`
	Transport             *http.Transport `json:"-"`

	// Variables determined at runtime.
	Version    version.Number    `json:"-"`
	systemInfo map[string]string `json:"-"`
	plugins    []plugin.Info     `json:"-"`

	// Internal variables.
	con        *http.Client
	api_url    string
	configTree *generic.Xml
	logger     *categoryLogger

	// Variables for testing, response bytes, headers, and response index.
	testInput       []*http.Request
	testOutput      []*http.Response
	testIndex       int
	authFileContent []byte
}

// Versioning returns the version number of PAN-OS.
func (c *Client) Versioning() version.Number {
	return c.Version
}

func (c *Client) SystemInfo(ctx context.Context) (map[string]string, error) {
	if c.systemInfo == nil {
		if err := c.RetrieveSystemInfo(ctx); err != nil {
			return nil, err
		}
	}

	return c.systemInfo, nil
}

func (c *Client) Plugins(ctx context.Context) ([]plugin.Info, error) {
	if c.plugins == nil {
		if err := c.RetrievePlugins(ctx); err != nil {
			return nil, err
		}
	}

	return c.plugins, nil
}

// GetTarget returns the Target param, used in certain API calls.
func (c *Client) GetTarget() string {
	return c.Target
}

// IsPanorama returns true if this is Panorama.
func (c *Client) IsPanorama() (bool, error) {
	if len(c.systemInfo) == 0 {
		return false, fmt.Errorf("SystemInfo is nil")
	}

	model, ok := c.systemInfo["model"]
	if !ok {
		return false, fmt.Errorf("model not present in SystemInfo")
	}

	return model == "Panorama" || strings.HasPrefix(model, "M-"), nil
}

// IsFirewall returns true if this PAN-OS seems to be a NGFW instance.
func (c *Client) IsFirewall() (bool, error) {
	ans, err := c.IsPanorama()
	if err != nil {
		return ans, err
	}

	return !ans, nil
}

// Setup does validation and initialization in preparation to start executing API
// commands against PAN-OS.
func (c *Client) Setup() error {
	var err error

	// Load up the JSON config file.
	var json_client Client
	if c.AuthFile != "" {
		var b []byte
		if len(c.testOutput) == 0 {
			b, err = os.ReadFile(c.AuthFile)
		} else {
			b = c.authFileContent
		}

		if err != nil {
			return err
		}

		if err = json.Unmarshal(b, &json_client); err != nil {
			return err
		}
	}

	// Hostname.
	if c.Hostname == "" {
		if val := os.Getenv("PANOS_HOSTNAME"); c.CheckEnvironment && val != "" {
			c.Hostname = val
		} else if json_client.Hostname != "" {
			c.Hostname = json_client.Hostname
		}
	}

	// Username.
	if c.Username == "" {
		if val := os.Getenv("PANOS_USERNAME"); c.CheckEnvironment && val != "" {
			c.Username = val
		} else {
			c.Username = json_client.Username
		}
	}

	// Password.
	if c.Password == "" {
		if val := os.Getenv("PANOS_PASSWORD"); c.CheckEnvironment && val != "" {
			c.Password = val
		} else {
			c.Password = json_client.Password
		}
	}

	// API key.
	if c.ApiKey == "" {
		if val := os.Getenv("PANOS_API_KEY"); c.CheckEnvironment && val != "" {
			c.ApiKey = val
		} else {
			c.ApiKey = json_client.ApiKey
		}
	}

	// Protocol.
	if c.Protocol == "" {
		if val := os.Getenv("PANOS_PROTOCOL"); c.CheckEnvironment && val != "" {
			c.Protocol = val
		} else if json_client.Protocol != "" {
			c.Protocol = json_client.Protocol
		} else {
			c.Protocol = "https"
		}
	}
	if c.Protocol != "http" && c.Protocol != "https" {
		return fmt.Errorf("Invalid protocol %q.  Must be \"http\" or \"https\"", c.Protocol)
	}

	// Port.
	if c.Port == 0 {
		if val := os.Getenv("PANOS_PORT"); c.CheckEnvironment && val != "" {
			if cp, err := strconv.Atoi(val); err != nil {
				return fmt.Errorf("Failed to parse the env port number: %s", err)
			} else {
				c.Port = cp
			}
		} else if json_client.Port != 0 {
			c.Port = json_client.Port
		}
	}
	if c.Port > 65535 {
		return fmt.Errorf("Port %d is out of bounds", c.Port)
	}

	// Target.
	if c.Target == "" {
		if val := os.Getenv("PANOS_TARGET"); c.CheckEnvironment && val != "" {
			c.Target = val
		} else {
			c.Target = json_client.Target
		}
	}

	// Headers.
	if len(c.Headers) == 0 {
		if val := os.Getenv("PANOS_HEADERS"); c.CheckEnvironment && val != "" {
			if err := json.Unmarshal([]byte(val), &c.Headers); err != nil {
				return err
			}
		}
		if len(c.Headers) == 0 && len(json_client.Headers) > 0 {
			c.Headers = make(map[string]string)
			for k, v := range json_client.Headers {
				c.Headers[k] = v
			}
		}
	}

	// Skip verify cert.
	if !c.SkipVerifyCertificate {
		if val := os.Getenv("PANOS_SKIP_VERIFY_CERTIFICATE"); c.CheckEnvironment && val != "" {
			if vcb, err := strconv.ParseBool(val); err != nil {
				return err
			} else if vcb {
				c.SkipVerifyCertificate = vcb
			}
		}
		if !c.SkipVerifyCertificate && json_client.SkipVerifyCertificate {
			c.SkipVerifyCertificate = true
		}
	}

	err = c.setupLogging(c.Logging)
	if err != nil {
		return err
	}

	// Setup the client.
	if c.Transport == nil {
		c.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.SkipVerifyCertificate,
			},
		}
	}
	c.con = &http.Client{
		Transport: c.Transport,
	}

	// Sanity check.
	if c.Hostname == "" {
		return fmt.Errorf("hostname must be specified")
	}
	if c.ApiKey == "" && (c.Username == "" && c.Password == "") {
		return fmt.Errorf("either API key or both username and password must be specified")
	}

	// Configure the api url.
	if c.Port == 0 {
		c.api_url = fmt.Sprintf("%s://%s/api", c.Protocol, c.Hostname)
	} else {
		c.api_url = fmt.Sprintf("%s://%s:%d/api", c.Protocol, c.Hostname, c.Port)
	}

	return nil
}

// SetupLocalInspection configures the client for local inspection.
//
// The version number is taken from the schema if it's present.  If it is not
// present, then the version number falls back to what's specified in panosVersion.
func (c *Client) SetupLocalInspection(schema, panosVersion string) error {
	var err error

	if err = c.LoadPanosConfig([]byte(schema)); err != nil {
		return err
	}

	if c.Version.Major == 0 && c.Version.Minor == 0 && c.Version.Patch == 0 {
		if panosVersion == "" {
			return fmt.Errorf("no version found in the schema; the version must be specified")
		}

		c.Version, err = version.New(panosVersion)
		if err != nil {
			return err
		}
	}

	return nil
}

// Initialize retrieves the API key if needed then retrieves the system info.
func (c *Client) Initialize(ctx context.Context) error {
	var err error

	if c.ApiKey == "" {
		if err = c.RetrieveApiKey(ctx); err != nil {
			return err
		}
	}

	return nil
}

// RetrieveSystemInfo performs "show system info" and saves it SystemInfo.
func (c *Client) RetrieveSystemInfo(ctx context.Context) error {
	var err error

	type req struct {
		XMLName xml.Name `xml:"show"`
		Cmd     string   `xml:"system>info"`
	}

	type system struct {
		Tags []generic.Xml `xml:",any"`
	}

	type resp struct {
		System system `xml:"result>system"`
	}

	cmd := &xmlapi.Op{
		Command: req{},
		Target:  c.Target,
	}

	var ans resp
	if _, _, err = c.Communicate(ctx, cmd, false, &ans); err != nil {
		return err
	}

	c.systemInfo = make(map[string]string, len(ans.System.Tags))
	for _, t := range ans.System.Tags {
		if t.TrimmedText == nil {
			continue
		}
		c.systemInfo[t.XMLName.Local] = *t.TrimmedText
		if t.XMLName.Local == "sw-version" {
			c.Version, err = version.New(*t.TrimmedText)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// RetrievePanosConfig retrieves the running config, candidate config,
// or the specified saved config file.
//
// If the name is candidate, then the candidate config is retrieved.  If the
// name is running, then the running config is retrieved.  Otherwise, then
// the name is assumed to be the name of the saved config.
func (c *Client) RetrievePanosConfig(ctx context.Context, name string) ([]byte, error) {
	type req struct {
		XMLName   xml.Name `xml:"show"`
		Running   *string  `xml:"config>running"`
		Candidate *string  `xml:"config>candidate"`
		Saved     *string  `xml:"config>saved"`
	}

	type rdata struct {
		Data []byte `xml:",innerxml"`
	}

	type resp struct {
		XMLName xml.Name `xml:"response"`
		Result  rdata    `xml:"result"`
	}

	var s string
	cs := req{}
	switch name {
	case "candidate":
		cs.Candidate = &s
	case "running":
		cs.Running = &s
	default:
		cs.Saved = &name
	}

	cmd := &xmlapi.Op{
		Command: cs,
		Target:  c.Target,
	}

	var ans resp
	if _, _, err := c.Communicate(ctx, cmd, false, &ans); err != nil {
		return nil, err
	}

	return ans.Result.Data, nil
}

// RetrieveApiKey refreshes the API key.
//
// This function unsets the ApiKey value and thus requires that the Username and Password
// be specified.
func (c *Client) RetrieveApiKey(ctx context.Context) error {
	type key_gen_ans struct {
		Key string `xml:"result>key"`
	}

	cmd := &xmlapi.KeyGen{
		Username: c.Username,
		Password: c.Password,
	}

	var ans key_gen_ans
	_, _, err := c.Communicate(ctx, cmd, false, &ans)
	if err != nil {
		return err
	}

	c.ApiKey = ans.Key

	return nil
}

// RetrievePlugins refreshes the Plugins info from PAN-OS.
func (c *Client) RetrievePlugins(ctx context.Context) error {
	cmd := &xmlapi.Op{
		Command: plugin.GetPlugins{},
	}

	var ans plugin.PackageListing
	_, _, err := c.Communicate(ctx, cmd, false, &ans)
	if err != nil {
		return err
	}

	c.plugins = ans.Listing()

	return nil
}

// LoadPanosConfig stores the given XML document into this client, allowing
// the user to use various namespace functions to query the config.  This
// is referred to as local inspection mode.
//
// The config given must be in the form of `<config>...</config>`.
//
// If the 'detail-version' attribute is present in the XML, then it's saved to this
// instance's Version attribute.
func (c *Client) LoadPanosConfig(config []byte) error {
	var ans generic.Xml
	if err := xml.Unmarshal(config, &ans); err != nil {
		return err
	}

	if ans.XMLName.Local != "config" {
		return fmt.Errorf("Expected \"config\" at the root, found %q", ans.XMLName.Local)
	}

	if ans.DetailedVersion != nil {
		vn, err := version.New(*ans.DetailedVersion)
		if err != nil {
			return err
		}
		c.Version = vn
	}

	c.configTree = &generic.Xml{
		XMLName: xml.Name{
			Local: "a",
		},
		Nodes: []generic.Xml{ans},
	}

	return nil
}

// ReadFromConfig returns the XML at the given XPATH location.  This is
// referred to as local inspection mode.
//
// If the XPATH is a listing, then set withPackaging to false.
func (c *Client) ReadFromConfig(ctx context.Context, path []string, withPackaging bool, ans any) ([]byte, error) {
	if c.configTree == nil {
		return nil, fmt.Errorf("no config loaded")
	}

	if len(path) == 0 {
		return nil, fmt.Errorf("path is empty")
	}

	entryPrefix := "entry[@name='"
	entrySuffix := "']"

	config := c.configTree
	for _, pp := range path {
		var tag, name string
		if strings.HasPrefix(pp, entryPrefix) && strings.HasSuffix(pp, entrySuffix) {
			tag = "entry"
			name = strings.TrimSuffix(strings.TrimPrefix(pp, entryPrefix), entrySuffix)
		} else {
			tag = pp
		}

		found := false
		for _, node := range config.Nodes {
			if node.XMLName.Local == tag && (node.Name == nil || *node.Name == name) {
				found = true
				config = &node
				break
			}
		}

		if !found {
			config = nil
			break
		}
	}

	if config == nil {
		return nil, errors.ObjectNotFound()
	}

	b, err := xml.Marshal(config)
	if err != nil {
		return b, err
	}

	var newb []byte
	if !withPackaging {
		newb = append([]byte(nil), b...)
	} else {
		newb = make([]byte, 0, len(b)+7)
		newb = append(newb, []byte("<q>")...)
		newb = append(newb, b...)
		newb = append(newb, []byte("</q>")...)
	}

	if ans == nil {
		return newb, nil
	}

	err = xml.Unmarshal(newb, ans)
	return newb, err
}

func (c *Client) Clock(ctx context.Context) (time.Time, error) {
	type creq struct {
		XMLName xml.Name `xml:"show"`
		Cmd     string   `xml:"clock"`
	}

	type cresp struct {
		Result string `xml:"result"`
	}

	cmd := &xmlapi.Op{
		Command: creq{},
		Target:  c.Target,
	}
	var ans cresp

	if _, _, err := c.Communicate(ctx, cmd, false, &ans); err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.UnixDate+"\n", ans.Result)
}

// ChunkedMultiConfig wraps
func (c *Client) ChunkedMultiConfig(ctx context.Context, mc *xmlapi.MultiConfig, strict bool, extras url.Values) ([]xmlapi.ChunkedMultiConfigResponse, error) {
	if mc.BatchSize == 0 {
		return nil, fmt.Errorf("cannot use ChunkedMultiConfig() with batchSize of 0")
	}

	if len(mc.Operations) == 0 {
		return nil, nil
	}

	var chunked [][]xmlapi.MultiConfigOperation
	for i := 0; i < len(mc.Operations); i += mc.BatchSize {
		end := i + mc.BatchSize
		if end > len(mc.Operations) {
			end = len(mc.Operations)
		}

		chunked = append(chunked, mc.Operations[i:end])
	}

	var response []xmlapi.ChunkedMultiConfigResponse
	for _, chunk := range chunked {
		updates := &xmlapi.MultiConfig{
			Operations: chunk,
		}

		data, resp, mcResp, err := c.MultiConfig(ctx, updates, strict, extras)
		if err != nil {
			return response, err
		}

		response = append(response, xmlapi.ChunkedMultiConfigResponse{
			Data:                data,
			HttpResponse:        resp,
			MultiConfigResponse: mcResp,
		})
	}

	return response, nil
}

// MultiConfig does a "multi-config" type command.
//
// Param strict should be true if you want strict transactional support.
//
// Note that the error returned from this function is only if there was an error
// unmarshaling the response into the the multi config response struct.  If the
// multi config itself failed, then the reason can be found in its results.
func (c *Client) MultiConfig(ctx context.Context, mc *xmlapi.MultiConfig, strict bool, extras url.Values) ([]byte, *http.Response, *xmlapi.MultiConfigResponse, error) {
	cmd := &xmlapi.Config{
		Action:              "multi-config",
		Element:             mc,
		StrictTransactional: strict,
		Target:              c.Target,
	}

	text, httpResp, err := c.Communicate(ctx, cmd, false, nil)

	// If the text is empty, then the err will have a real error we should report to the
	// invoker.  However, if the text is not empty, then ignore the error and parse the
	// multi config results using the specialized struct custom designed to handle it.
	if len(text) == 0 {
		return text, httpResp, nil, err
	}

	mcResp, err := xmlapi.NewMultiConfigResponse(text)
	if err != nil {
		return text, httpResp, nil, err
	}

	if !mcResp.Ok() {
		return text, httpResp, mcResp, mcResp
	}

	return text, httpResp, mcResp, nil
}

// RequestPasswordHash requests a password hash of the given string.
func (c *Client) RequestPasswordHash(ctx context.Context, v string) (string, error) {
	type phash_req struct {
		XMLName xml.Name `xml:"request"`
		Val     string   `xml:"password-hash>password"`
	}

	cmd := &xmlapi.Op{
		Command: phash_req{
			Val: v,
		},
		Target: c.Target,
	}

	type phash_ans struct {
		Hash string `xml:"result>phash"`
	}

	var ans phash_ans
	if _, _, err := c.Communicate(ctx, cmd, false, &ans); err != nil {
		return "", err
	}

	return ans.Hash, nil
}

// Generate API Key for the provided username/password pair.
//
// Generation of the API Key will also invalidate the existing key for the given user.
func (c *Client) GenerateApiKey(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf("username and password cannot be empty")
	}

	data := url.Values{}
	data.Set("user", username)
	data.Set("password", password)

	url := fmt.Sprintf("%s/?type=keygen", c.api_url)
	request, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.con.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("server returned an error when generating API Key: %s", string(bodyBytes))
		return "", err
	}

	type xmlResponse struct {
		XMLName xml.Name `xml:"response"`
		Key     string   `xml:"result>key"`
	}

	response := &xmlResponse{}
	err = xml.Unmarshal(bodyBytes, &response)
	if err != nil {
		return "", err
	}

	return response.Key, nil

}

// ValidateConfig performs a commit config validation check.
//
// Use WaitForJob and the uint returned from this function to get the
// results of the job.
func (c *Client) ValidateConfig(ctx context.Context, sleep time.Duration) (uint, error) {
	type req struct {
		XMLName xml.Name `xml:"validate"`
		Cmd     string   `xml:"full"`
	}

	cmd := &xmlapi.Op{
		Command: req{},
		Target:  c.Target,
	}

	id, _, _, err := c.StartJob(ctx, cmd)
	return id, err
}

// WaitForLogs polls PAN-OS until the given log retrieval job is complete.
//
// The sleep param is the time to wait between polling. Note that a sleep
// time less than 2 seconds may cause PAN-OS to take longer to finish the
// job.
func (c *Client) WaitForLogs(ctx context.Context, id uint, sleep time.Duration, resp any) ([]byte, error) {
	if id == 0 {
		return nil, fmt.Errorf("job ID must be specified")
	}

	var err error
	var data []byte
	var prev string

	cmd := &xmlapi.Log{
		Action: "get",
		JobId:  id,
	}

	var ans util.BasicJob
	for {
		ans = util.BasicJob{}

		data, _, err = c.Communicate(ctx, cmd, false, &ans)
		if err != nil {
			return data, err
		}

		if ans.Status != prev {
			prev = ans.Status
			// log %d id and %s prev
		}

		if ans.Status == "FIN" {
			break
		}

		if sleep > 0 {
			time.Sleep(sleep)
		}
	}

	if ans.Result == "FAIL" {
		if len(ans.Details.Lines) > 0 {
			return data, fmt.Errorf(ans.Details.String())
		} else {
			return data, fmt.Errorf("Job %d has failed", id)
		}
	}

	if resp == nil {
		return data, nil
	}

	err = xml.Unmarshal(data, resp)
	return data, err
}

// WaitForJob polls PAN-OS until the given job finishes.
//
// The sleep param is the time to wait between polling. Note that a sleep
// time less than 2 seconds may cause PAN-OS to take longer to finish the
// job.
func (c *Client) WaitForJob(ctx context.Context, id uint, sleep time.Duration, resp any) error {
	var err error
	var prev uint
	var data []byte
	dp := false
	all_ok := true

	type req struct {
		XMLName xml.Name `xml:"show"`
		Id      uint     `xml:"jobs>id"`
	}

	cmd := &xmlapi.Op{
		Command: req{
			Id: id,
		},
		Target: c.Target,
	}

	var ans util.BasicJob
	for {
		ans = util.BasicJob{}

		data, _, err = c.Communicate(ctx, cmd, false, &ans)
		if err != nil {
			return err
		}

		if ans.Progress != prev {
			prev = ans.Progress
			// log the change.
		}

		// Check for device commits.
		all_done := true
		for _, d := range ans.Devices {
			// log %q d.Serial %s d.Result
			if d.Result == "PEND" {
				all_done = false
				break
			} else if d.Result != "OK" && all_ok {
				all_ok = false
			}
		}

		// Check if the job's done.
		if ans.Progress == 100 {
			if all_done {
				break
			} else if !dp {
				// log waiting for %d devices len(ans.Devices)
				dp = true
			}
		}

		if sleep > 0 {
			time.Sleep(sleep)
		}
	}

	if ans.Result == "FAIL" {
		if len(ans.Details.Lines) > 0 {
			return fmt.Errorf(ans.Details.String())
		} else {
			return fmt.Errorf("Job %d has failed", id)
		}
	} else if !all_ok {
		return fmt.Errorf("Commit failed on one or more devices")
	}

	if resp == nil {
		return nil
	}

	return xml.Unmarshal(data, resp)
}

// RevertToRunningConfig discards any changes made and reverts to the last
// committed config.
func (c *Client) RevertToRunningConfig(ctx context.Context, vsys string) error {
	type req struct {
		XMLName xml.Name `xml:"load"`
		Cmd     string   `xml:"config>from"`
	}

	cmd := &xmlapi.Op{
		Command: req{
			Cmd: "running-config.xml",
		},
		Vsys:   vsys,
		Target: c.Target,
	}

	_, _, err := c.Communicate(ctx, cmd, false, nil)
	return err
}

// StartJob sends the given command, which starts a job on PAN-OS.
//
// The uint returned is the job ID.
func (c *Client) StartJob(ctx context.Context, cmd util.PangoCommand) (uint, []byte, *http.Response, error) {
	var ans util.JobResponse
	b, resp, err := c.Communicate(ctx, cmd, false, &ans)
	if err != nil {
		return 0, b, resp, err
	}

	return ans.Id, b, resp, nil
}

// Communicate sends the given content to PAN-OS.
//
// The API key is sent either in the request body or as a header.
//
// The timeout for the operation is taken from the context.
func (c *Client) Communicate(ctx context.Context, cmd util.PangoCommand, strip bool, ans any) ([]byte, *http.Response, error) {
	if cmd == nil {
		return nil, nil, fmt.Errorf("cmd is nil")
	}

	data, err := cmd.AsUrlValues()
	if err != nil {
		return nil, nil, err
	}

	if c.ApiKeyInRequest && c.ApiKey != "" && data.Get("key") == "" {
		data.Set("key", c.ApiKey)
	}

	c.logSend(data)

	req, err := http.NewRequestWithContext(ctx, "POST", c.api_url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.sendRequest(ctx, req, strip, ans)
}

// ImportFile imports the given file into PAN-OS.
func (c *Client) ImportFile(ctx context.Context, cmd *xmlapi.Import, content []byte, filename, fp string, strip bool, ans any) ([]byte, *http.Response, error) {
	if cmd == nil {
		return nil, nil, fmt.Errorf("cmd is nil")
	}

	data, err := cmd.AsUrlValues()
	if err != nil {
		return nil, nil, err
	}

	if c.ApiKeyInRequest && c.ApiKey != "" && data.Get("key") == "" {
		data.Set("key", c.ApiKey)
	}

	c.logSend(data)

	buf := bytes.Buffer{}
	w := multipart.NewWriter(&buf)

	for k := range data {
		err = w.WriteField(k, data.Get(k))
		if err != nil {
			return nil, nil, err
		}
	}

	w2, err := w.CreateFormFile(fp, filename)
	if err != nil {
		return nil, nil, err
	}

	if _, err = io.Copy(w2, bytes.NewReader(content)); err != nil {
		return nil, nil, err
	}

	w.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", c.api_url, &buf)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	return c.sendRequest(ctx, req, strip, ans)
}

// ExportFile retrieves a file from PAN-OS.
func (c *Client) ExportFile(ctx context.Context, cmd *xmlapi.Export, ans any) (string, []byte, *http.Response, error) {
	if cmd == nil {
		return "", nil, nil, fmt.Errorf("cmd is nil")
	}

	data, err := cmd.AsUrlValues()
	if err != nil {
		return "", nil, nil, err
	}

	if c.ApiKeyInRequest && c.ApiKey != "" && data.Get("key") == "" {
		data.Set("key", c.ApiKey)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.api_url, strings.NewReader(data.Encode()))
	if err != nil {
		return "", nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	b, resp, err := c.sendRequest(ctx, req, false, ans)
	if err != nil {
		return "", b, resp, err
	}

	var filename string
	mediatype, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err == nil && mediatype == "attachment" {
		filename = params["filename"]
	}

	return filename, b, resp, nil
}

// GetTechSupportFile returns the tech support file .tgz file.
func (c *Client) GetTechSupportFile(ctx context.Context) (string, []byte, error) {
	cmd := &xmlapi.Export{
		Category: "tech-support",
	}

	id, _, _, err := c.StartJob(ctx, cmd)
	if err != nil {
		return "", nil, err
	}

	cmd.Action = "status"
	cmd.JobId = id

	var resp util.BasicJob
	var prev uint
	for {
		resp = util.BasicJob{}

		if _, _, _, err = c.ExportFile(ctx, cmd, &resp); err != nil {
			return "", nil, err
		}

		// The progress is not an uint when the job completes, so don't print
		// the progress as 0 when the job is actually complete.
		if resp.Progress != prev && resp.Progress != 0 {
			prev = resp.Progress
			// log %d resp.Id at %d prev percentage complete.
		}

		if resp.Status == "FIN" {
			break
		}

		time.Sleep(2 * time.Second)
	}

	if resp.Result == "FAIL" {
		return "", nil, fmt.Errorf(resp.Details.String())
	}

	cmd.Action = "get"

	filename, b, _, err := c.ExportFile(ctx, cmd, nil)
	return filename, b, err
}

//
// Internal functions.
//

func (c *Client) setupLogging(logging LoggingInfo) error {
	var logger *slog.Logger

	var logLevel slog.Level
	var levelStr string
	if levelStr = os.Getenv("PANOS_LOG_LEVEL"); c.CheckEnvironment && levelStr != "" {
		err := logLevel.UnmarshalText([]byte(levelStr))
		if err != nil {
			return err
		}
	} else {
		logLevel = slog.LevelInfo
	}

	if logging.SLogHandler == nil {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))
		logger.Info("No slog handler provided, creating default os.Stderr handler.", "LogLevel", logging.LogLevel.Level())
	} else {
		logger = slog.New(logging.SLogHandler)
		if logging.LogLevel != 0 {
			logger.Warn("LogLevel is ignored when using custom SLog handler.")
		}
	}

	// 1. logging.LogCategories has the highest priority
	// 2. If logging.LogCategories is not set, we check logging.LogSymbols
	// 3. If logging.LogSymbols is empty and c.CheckEnvironment is true we consult
	//    PANOS_LOG_CATEGORIES environment variable.
	// 4. If no logging categories have been selected, default to basic library logging
	//    (i.e. "pango" category)
	logMask := logging.LogCategories
	var err error
	if logMask == 0 {
		logMask, err = LogCategoryFromStrings(logging.LogSymbols)
		if err != nil {
			return err
		}

		if logMask == 0 {
			if val := os.Getenv("PANOS_LOG_CATEGORIES"); c.CheckEnvironment && val != "" {
				symbols := strings.Split(val, ",")
				logMask, err = LogCategoryFromStrings(symbols)
				if err != nil {
					return err
				}
			}
		}
	}

	// To disable logging completely, use custom SLog handler that discards all logs,
	// e.g. https://github.com/golang/go/issues/62005 for the example of such handler.
	if logMask == 0 {
		logMask = LogCategoryPango
	}

	enabledLogging, _ := LogCategoryAsStrings(logMask)
	logger.Info("Pango logging configured", "symbols", enabledLogging)

	c.logger = newCategoryLogger(logger, logMask)

	return nil
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, strip bool, ans any) ([]byte, *http.Response, error) {
	// Optional: set the API key in the header.
	if !c.ApiKeyInRequest && c.ApiKey != "" {
		req.Header.Set("X-PAN-KEY", c.ApiKey)
	}

	// Configure all user given headers.
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	// Perform the operation.
	var err error
	var resp *http.Response
	if len(c.testOutput) != 0 {
		c.testInput = append(c.testInput, req)
		resp = c.testOutput[c.testIndex%len(c.testOutput)]
		c.testIndex++
	} else {
		resp, err = c.con.Do(req)
	}

	if err != nil {
		return nil, nil, err
	}

	// Read in the response.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, resp, err
	}

	// Check the response for errors.
	if err = errors.Parse(body); err != nil {
		return body, resp, err
	}

	// Optional: strip the "response" tag from the XML returned.
	if strip {
		var index int
		gt := []byte(">")
		lt := []byte("<")

		index = bytes.Index(body, gt)
		if index > 0 {
			body = body[index+1:]
			index = bytes.LastIndex(body, lt)
			if index > 0 {
				body = body[:index]
			}
		}
	}

	if ans == nil {
		return body, resp, nil
	}

	// Optional: unmarshal using the struct passed in.
	err = xml.Unmarshal(body, ans)
	if err != nil {
		return body, resp, fmt.Errorf("err unmarshaling into provided interface: %s", err)
	}

	c.logger.WithLogCategory(LogCategoryReceive).Debug("Received data from server", "body", c.prepareReceiveDataForLogging(body))

	return body, resp, nil
}

func (c *Client) logSend(data url.Values) {
	c.logger.WithLogCategory(LogCategoryPango).Debug("Hello World!")
	sendData := slog.Group("data", c.prepareSendDataForLogging(data)...)
	if c.logger.enabledFor(LogCategoryCurl) {
		curlEquivalent := slog.Group("curl", c.prepareSendDataAsCurl(data)...)
		c.logger.WithLogCategory(LogCategorySend).Debug("data sent to the server", sendData, curlEquivalent)
	} else {
		c.logger.WithLogCategory(LogCategorySend).Debug("data sent to the server", sendData)
	}
}

func (c *Client) prepareSendDataForLogging(data url.Values) []any {
	filterSensitive := !c.logger.enabledFor(LogCategorySensitive)

	preparedValues := make([]any, 0)

	for key, values := range data {
		for _, value := range values {
			preparedValues = append(preparedValues, key)
			if filterSensitive {
				switch key {
				case "key", "password":
					preparedValues = append(preparedValues, "***")
				default:
					preparedValues = append(preparedValues, value)
				}
			} else {
				preparedValues = append(preparedValues, value)
			}
		}
	}

	return preparedValues
}

func (c *Client) prepareSendDataAsCurl(data url.Values) []any {
	filterSensitive := !c.logger.enabledFor(LogCategorySensitive)

	// A map of items and their values for items that require additional
	// processing, and not be sent as is.
	special := map[string]string{
		"element": "",
	}

	sensitive := []string{"user", "password", "key", "host"}

	values := url.Values{}
	for key, value := range data {
		isSpecial := false
		for needle := range special {
			if key == needle {
				isSpecial = true
				special[key] = value[0]
				break
			}
		}

		isSensitive := false
		for _, needle := range sensitive {
			if key == needle {
				isSensitive = true
				break
			}
		}

		// Only non-special values should be part of the
		// encoded url, and sensitive values should only be
		// visible when LogCategorySensitiveis set.
		if !isSpecial {
			if isSensitive && filterSensitive {
				values[key] = []string{strings.ToUpper(key)}
			} else {
				values[key] = value
			}
		}
	}

	curlCmd := []string{"curl"}

	if c.SkipVerifyCertificate {
		curlCmd = append(curlCmd, "-k")
	}

	if len(c.Headers) > 0 && !filterSensitive {
		for k, v := range c.Headers {
			curlCmd = append(curlCmd, "--header")
			var header string
			if v == "" {
				header = fmt.Sprintf("'%s;'", k)
			} else {
				header = fmt.Sprintf("'%s: %s'", k, v)
			}
			curlCmd = append(curlCmd, header)
		}
	}

	hostname := c.Hostname
	if filterSensitive {
		hostname = "HOST"
	}
	port := ""
	if c.Port != 0 {
		port = fmt.Sprintf(":%d", c.Port)
	}

	encodedValues := ""
	if len(values) > 0 {
		encodedValues = fmt.Sprintf("?%s", values.Encode())
	}

	urlTemplate := "'%s://%s%s/api%s'"
	apiUrl := fmt.Sprintf(urlTemplate, c.Protocol, hostname, port, encodedValues)

	curlCmd = append(curlCmd, apiUrl)

	curlData := []any{"command", strings.Join(curlCmd, " ")}

	if special["element"] != "" {
		curlData = append(curlData, "element.xml")
		curlData = append(curlData, special["element"])
	}

	return curlData
}

func (c *Client) prepareReceiveDataForLogging(body []byte) string {
	replacedBody := string(body)
	keyStartIdx := strings.Index(replacedBody, "<key>")
	keyEndIdx := strings.Index(replacedBody, "</key>")
	if keyStartIdx != -1 && keyEndIdx != -1 {
		replacedBody = replacedBody[:keyStartIdx] + "***" + replacedBody[keyEndIdx+len("</key>"):]
	}
	return replacedBody
}
