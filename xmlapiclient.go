package pango

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	_ "mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/plugin"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

// XmlApiClient is an XML API client connection.  If provides wrapper functions
// for invoking the various PAN-OS API methods.  After creating the client,
// invoke Setup() followed by Initialize() to prepare it for use.
type XmlApiClient struct {
	Hostname        string            `json:"hostname"`
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	ApiKey          string            `json:"api_key"`
	Protocol        string            `json:"protocol"`
	Port            int               `json:"port"`
	Target          string            `json:"target"`
	ApiKeyInRequest bool              `json:"api_key_in_request"`
	Headers         map[string]string `json:"headers"`
	Logging         map[string]bool   `json:"logging"`

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
	SystemInfo map[string]string `json:"-"`
	Plugin     []plugin.Info     `json:"-"`

	// Internal variables.
	credsFile  string
	con        *http.Client
	api_url    string
	configTree *generic.Xml

	// Variables for testing, response bytes, headers, and response index.
	testInput       []*http.Request
	testOutput      []*http.Response
	testIndex       int
	authFileContent []byte
}

func (c *XmlApiClient) Versioning() version.Number { return c.Version }

func (c *XmlApiClient) Plugins() []plugin.Info {
	return c.Plugin
}

// Setup does validation and initialization in preparation to start executing API
// commands against PAN-OS.
func (c *XmlApiClient) Setup() error {
	var err error

	// Load up the JSON config file.
	var json_client XmlApiClient
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

	// Logging.
	if len(c.Logging) == 0 {
		if val := os.Getenv("PANOS_LOGGING"); c.CheckEnvironment && val != "" {
			ll := strings.Split(val, ",")
			for _, k := range ll {
				c.Logging[k] = true
			}
		} else if len(json_client.Logging) != 0 {
			c.Logging = make(map[string]bool)
			for k, v := range json_client.Logging {
				c.Logging[k] = v
			}
		}
	}
	if len(c.Logging) == 0 {
		c.Logging = map[string]bool{
			"quiet": true,
		}
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

// Initialize retrieves the API key if needed then retrieves the system info.
func (c *XmlApiClient) Initialize(ctx context.Context) error {
	if c.ApiKey == "" {
		if err := c.RetrieveApiKey(ctx); err != nil {
			return err
		}
	}

	if err := c.RetrieveSystemInfo(ctx); err != nil {
		return err
	}

	return nil
}

// RetrieveSystemInfo performs "show system info" and saves it SystemInfo.
func (c *XmlApiClient) RetrieveSystemInfo(ctx context.Context) error {
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

	c.SystemInfo = make(map[string]string, len(ans.System.Tags))
	for _, t := range ans.System.Tags {
		if t.TrimmedText == nil {
			continue
		}
		c.SystemInfo[t.XMLName.Local] = *t.TrimmedText
		if t.XMLName.Local == "sw-version" {
			c.Version, err = version.New(*t.TrimmedText)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// RetrieveApiKey refreshes the API key.
//
// This function unsets the ApiKey value and thus requires that the Username and Password
// be specified.
func (c *XmlApiClient) RetrieveApiKey(ctx context.Context) error {
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
func (c *XmlApiClient) RetrievePlugins(ctx context.Context) error {
	cmd := &xmlapi.Op{
		Command: plugin.GetPlugins{},
	}

	var ans plugin.PackageListing
	_, _, err := c.Communicate(ctx, cmd, false, &ans)
	if err != nil {
		return err
	}

	c.Plugin = ans.Listing()

	return nil
}

// GetTarget returns the Target param, used in certain API calls.
func (c *XmlApiClient) GetTarget() string {
	return c.Target
}

// MultiConfig does a "multi-config" type command.
//
// Param strict should be true if you want strict transactional support.
//
// Note that the error returned from this function is only if there was an error
// unmarshaling the response into the the multi config response struct.  If the
// multi config itself failed, then the reason can be found in its results.
func (c *XmlApiClient) MultiConfig(ctx context.Context, mc *xmlapi.MultiConfig, strict bool, extras url.Values) ([]byte, *http.Response, *xmlapi.MultiConfigResponse, error) {
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
func (c *XmlApiClient) RequestPasswordHash(ctx context.Context, v string) (string, error) {
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

// Communicate sends the given content to PAN-OS.
//
// The API key is sent either in the request body or as a header.
//
// The timeout for the operation is taken from the context.
func (c *XmlApiClient) Communicate(ctx context.Context, cmd util.PangoCommand, strip bool, ans interface{}) ([]byte, *http.Response, error) {
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

	req, err := http.NewRequestWithContext(ctx, "POST", c.api_url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.sendRequest(ctx, req, strip, ans)
}

// CommunicateFile sends the given file to PAN-OS.
//
// The API key is sent either in the request body or as a header.
//
// The timeout for the operation is taken from the context.
func (c *XmlApiClient) CommunicateFile(ctx context.Context, content, filename, fp string, data url.Values, strip bool, ans interface{}) ([]byte, *http.Response, error) {
	var err error

	if c.ApiKeyInRequest && c.ApiKey != "" && data.Get("key") == "" {
		data.Set("key", c.ApiKey)
	}

	//c.logSend(data)

	buf := bytes.Buffer{}
	w := multipart.NewWriter(&buf)

	for k := range data {
		w.WriteField(k, data.Get(k))
	}

	w2, err := w.CreateFormFile(fp, filename)
	if err != nil {
		return nil, nil, err
	}

	if _, err = io.Copy(w2, strings.NewReader(content)); err != nil {
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

//
// Internal functions.
//

func (c *XmlApiClient) sendRequest(ctx context.Context, req *http.Request, strip bool, ans interface{}) ([]byte, *http.Response, error) {
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

	if true {
		log.Printf("Response = %s", body)
	}

	// Check the response for errors.
	if err = errors.Parse(body); err != nil {
		return body, resp, err
	}

	if strip {
		var index int
		gt := []byte(">")
		lt := []byte("<")

		// Remove 'response'.
		index = bytes.Index(body, gt)
		body = body[index+1:]
		index = bytes.LastIndex(body, lt)
		body = body[:index]
		log.Printf("new body: %q", body)
	}

	if ans == nil {
		return body, resp, nil
	}

	// Optional: unmarshal using the struct passed in.
	err = xml.Unmarshal(body, ans)
	if err != nil {
		return body, resp, fmt.Errorf("err unmarshaling into provided interface: %s", err)
	}

	return body, resp, nil
}
