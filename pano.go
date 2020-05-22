package pango

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	// Various namespace imports.
	"github.com/PaloAltoNetworks/pango/dev"
	"github.com/PaloAltoNetworks/pango/licen"
	"github.com/PaloAltoNetworks/pango/netw"
	"github.com/PaloAltoNetworks/pango/objs"
	"github.com/PaloAltoNetworks/pango/pnrm"
	"github.com/PaloAltoNetworks/pango/poli"
	"github.com/PaloAltoNetworks/pango/userid"
)

// Panorama is a panorama specific client, providing version safe functions
// for the PAN-OS Xpath API methods.  After creating the object, invoke
// Initialize() to prepare it for use.
//
// It has the following namespaces:
//      * Licensing
//      * UserId
type Panorama struct {
	Client

	// Namespaces
	Device    *dev.PanoDev
	Licensing *licen.Licen
	UserId    *userid.UserId
	Panorama  *pnrm.Pnrm
	Objects   *objs.PanoObjs
	Policies  *poli.PanoPoli
	Network   *netw.PanoNetw
}

// Initialize does some initial setup of the Panorama connection, retrieves
// the API key if it was not already present, then performs "show system
// info" to get the PAN-OS version.  The full results are saved into the
// client's SystemInfo map.
//
// If not specified, the following is assumed:
//  * Protocol: https
//  * Port: (unspecified)
//  * Timeout: 10
//  * Logging: LogAction | LogUid
func (c *Panorama) Initialize() error {
	if len(c.rb) == 0 {
		var e error

		if e = c.initCon(); e != nil {
			return e
		} else if e = c.initApiKey(); e != nil {
			return e
		} else if e = c.initSystemInfo(); e != nil {
			return e
		}
		c.initPlugins()
	} else {
		c.Hostname = "localhost"
		c.ApiKey = "password"
	}
	c.initNamespaces()

	return nil
}

// InitializeUsing does Initialize(), but takes in a filename that contains
// fallback authentication credentials if they aren't specified.
//
// The order of preference for auth / connection settings is:
//
// * explicitly set
// * environment variable (set chkenv to true to enable this)
// * json file
func (c *Panorama) InitializeUsing(filename string, chkenv bool) error {
	c.CheckEnvironment = chkenv
	c.credsFile = filename

	return c.Initialize()
}

// CreateVmAuthKey creates a VM auth key to bootstrap a VM-Series firewall.
//
// VM auth keys are only valid for the number of hours specified.
func (c *Panorama) CreateVmAuthKey(hours int) (VmAuthKey, error) {
	clock, err := c.Clock()
	if err != nil {
		c.LogOp("(op) Failed to get/parse system time: %s", err)
	}

	type ak_req struct {
		XMLName  xml.Name `xml:"request"`
		Duration int      `xml:"bootstrap>vm-auth-key>generate>lifetime"`
	}

	type ak_resp struct {
		Msg string `xml:"result"`
	}

	req := ak_req{Duration: hours}
	ans := ak_resp{}

	c.LogOp("(op) generating a vm auth code")
	if b, err := c.Op(req, "", nil, &ans); err != nil {
		return VmAuthKey{}, err
	} else if ans.Msg == "" {
		return VmAuthKey{}, fmt.Errorf("No msg: %s", b)
	} else if !strings.HasPrefix(ans.Msg, "VM auth key ") {
		return VmAuthKey{}, fmt.Errorf("Wrong resp prefix: %s", b)
	}

	tokens := strings.Fields(ans.Msg)
	if len(tokens) != 9 {
		return VmAuthKey{}, fmt.Errorf("Got %d of 9 fields from: %s", len(tokens), ans.Msg)
	}

	key := VmAuthKey{
		AuthKey: tokens[3],
		Expiry:  strings.Join(tokens[7:], " "),
	}
	key.ParseExpires(clock)

	return key, nil
}

// GetVmAuthKeys gets the list of VM auth keys.
func (c *Panorama) GetVmAuthKeys() ([]VmAuthKey, error) {
	clock, err := c.Clock()
	if err != nil {
		c.LogOp("(op) Failed to get/parse system time: %s", err)
	}

	type l_req struct {
		XMLName xml.Name `xml:"request"`
		Msg     string   `xml:"bootstrap>vm-auth-key>show"`
	}

	type l_resp struct {
		List []VmAuthKey `xml:"result>bootstrap-vm-auth-keys>entry"`
	}

	req := l_req{}
	ans := l_resp{}

	c.LogOp("(op) listing vm auth codes")
	if _, err := c.Op(req, "", nil, &ans); err != nil {
		return nil, err
	}

	for i := range ans.List {
		ans.List[i].ParseExpires(clock)
	}

	return ans.List, nil
}

/** Public structs **/

// VmAuthKey is a VM auth key paired with when it expires.
//
// The Expiry field is the string returned from PAN-OS, while the Expires
// field is an attempt at parsing the Expiry field.
type VmAuthKey struct {
	AuthKey string `xml:"vm-auth-key"`
	Expiry  string `xml:"expiry-time"`
	Expires time.Time
}

// ParseExpires sets Expires from the Expiry field.
//
// Since PAN-OS does not output timezone information with the expirations,
// the current PAN-OS time is retrieved, which does contain timezone
// information.  Then in the string parsing for Expires, the location
// information of the system clock is applied.
func (o *VmAuthKey) ParseExpires(clock time.Time) {
	format := "2006/01/02 15:04:05"

	if t, err := time.ParseInLocation(format, o.Expiry, clock.Location()); err == nil {
		o.Expires = t
	}
}

/** Private functions **/

func (c *Panorama) initNamespaces() {
	c.Device = &dev.PanoDev{}
	c.Device.Initialize(c)

	c.Licensing = &licen.Licen{}
	c.Licensing.Initialize(c)

	c.UserId = &userid.UserId{}
	c.UserId.Initialize(c)

	c.Panorama = &pnrm.Pnrm{}
	c.Panorama.Initialize(c)

	c.Objects = &objs.PanoObjs{}
	c.Objects.Initialize(c)

	c.Policies = &poli.PanoPoli{}
	c.Policies.Initialize(c)

	c.Network = &netw.PanoNetw{}
	c.Network.Initialize(c)
}

func (c *Panorama) initPlugins() {
	c.LogOp("(op) getting plugin info")

	type plugin_req struct {
		XMLName xml.Name `xml:"show"`
		Cmd     string   `xml:"plugins>packages"`
	}

	type relNote struct {
		ReleaseNoteUrl string `xml:",cdata"`
	}

	type pkgInfo struct {
		Name        string  `xml:"name"`
		Version     string  `xml:"version"`
		ReleaseDate string  `xml:"release-date"`
		RelNote     relNote `xml:"release-note-url"`
		PackageFile string  `xml:"pkg-file"`
		Size        string  `xml:"size"`
		Platform    string  `xml:"platform"`
		Installed   string  `xml:"installed"`
		Downloaded  string  `xml:"downloaded"`
	}

	type pluginResp struct {
		Answer []pkgInfo `xml:"result>plugins>entry"`
	}

	req := plugin_req{}
	ans := pluginResp{}

	_, err := c.Op(req, "", nil, &ans)
	if err != nil {
		c.LogAction("WARNING: Failed to get plugin info: %s", err)
		return
	}

	c.Plugin = make([]map[string]string, 0, len(ans.Answer))
	for _, data := range ans.Answer {
		c.Plugin = append(c.Plugin, map[string]string{
			"name":             data.Name,
			"version":          data.Version,
			"release-date":     data.ReleaseDate,
			"release-note-url": data.RelNote.ReleaseNoteUrl,
			"package-file":     data.PackageFile,
			"size":             data.Size,
			"platform":         data.Platform,
			"installed":        data.Installed,
			"downloaded":       data.Downloaded,
		})
	}
}
