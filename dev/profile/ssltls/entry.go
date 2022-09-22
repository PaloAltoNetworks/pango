package ssltls

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// SSL/TLS service profile.
//
// The schema for the various booleans is default=yes/optional=yes/type=bool,
// and there are some prune-on's thrown in for good measure.  As such, the XML
// created for SET/EDIT will only contain the booleans which have been set to false,
//
// PAN-OS 7.0+.
type Entry struct {
	Name                      string
	Certificate               string
	MinVersion                string
	MaxVersion                string
	AllowAlgorithmRsa         bool // PAN-OS 8.1+
	AllowAlgorithmDhe         bool // PAN-OS 8.1+
	AllowAlgorithmEcdhe       bool // PAN-OS 8.1+
	AllowAlgorithm3des        bool // PAN-OS 8.1+
	AllowAlgorithmRc4         bool // PAN-OS 8.1+
	AllowAlgorithmAes128Cbc   bool // PAN-OS 8.1+
	AllowAlgorithmAes256Cbc   bool // PAN-OS 8.1+
	AllowAlgorithmAes128Gcm   bool // PAN-OS 8.1+
	AllowAlgorithmAes256Gcm   bool // PAN-OS 8.1+
	AllowAuthenticationSha1   bool // PAN-OS 8.1+
	AllowAuthenticationSha256 bool // PAN-OS 8.1+
	AllowAuthenticationSha384 bool // PAN-OS 8.1+
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Certificate = s.Certificate
	o.MinVersion = s.MinVersion
	o.MaxVersion = s.MaxVersion
	o.AllowAlgorithmRsa = s.AllowAlgorithmRsa
	o.AllowAlgorithmDhe = s.AllowAlgorithmDhe
	o.AllowAlgorithmEcdhe = s.AllowAlgorithmEcdhe
	o.AllowAlgorithm3des = s.AllowAlgorithm3des
	o.AllowAlgorithmRc4 = s.AllowAlgorithmRc4
	o.AllowAlgorithmAes128Cbc = s.AllowAlgorithmAes128Cbc
	o.AllowAlgorithmAes256Cbc = s.AllowAlgorithmAes256Cbc
	o.AllowAlgorithmAes128Gcm = s.AllowAlgorithmAes128Gcm
	o.AllowAlgorithmAes256Gcm = s.AllowAlgorithmAes256Gcm
	o.AllowAuthenticationSha1 = s.AllowAuthenticationSha1
	o.AllowAuthenticationSha256 = s.AllowAuthenticationSha256
	o.AllowAuthenticationSha384 = s.AllowAuthenticationSha384
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:        o.Name,
		Certificate: o.Certificate,
		MinVersion:  o.MinVersion,
		MaxVersion:  o.MaxVersion,
	}

	return ans
}

type entry_v1 struct {
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Certificate string   `xml:"certificate"`
	MinVersion  string   `xml:"protocol-settings>min-version,omitempty"`
	MaxVersion  string   `xml:"protocol-settings>max-version,omitempty"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{
		MinVersion: Tls1_0,
		MaxVersion: TlsMax,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v1(ans)
	return nil
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:        e.Name,
		Certificate: e.Certificate,
		MinVersion:  e.MinVersion,
		MaxVersion:  e.MaxVersion,
	}

	return ans
}

// PAN-OS 8.1
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:                      o.Name,
		Certificate:               o.Certificate,
		MinVersion:                o.MinVersion,
		MaxVersion:                o.MaxVersion,
		AllowAlgorithmRsa:         util.AsBool(o.AllowAlgorithmRsa),
		AllowAlgorithmDhe:         util.AsBool(o.AllowAlgorithmDhe),
		AllowAlgorithmEcdhe:       util.AsBool(o.AllowAlgorithmEcdhe),
		AllowAlgorithm3des:        util.AsBool(o.AllowAlgorithm3des),
		AllowAlgorithmRc4:         util.AsBool(o.AllowAlgorithmRc4),
		AllowAlgorithmAes128Cbc:   util.AsBool(o.AllowAlgorithmAes128Cbc),
		AllowAlgorithmAes256Cbc:   util.AsBool(o.AllowAlgorithmAes256Cbc),
		AllowAlgorithmAes128Gcm:   util.AsBool(o.AllowAlgorithmAes128Gcm),
		AllowAlgorithmAes256Gcm:   util.AsBool(o.AllowAlgorithmAes256Gcm),
		AllowAuthenticationSha1:   util.AsBool(o.AllowAuthenticationSha1),
		AllowAuthenticationSha256: util.AsBool(o.AllowAuthenticationSha256),
		AllowAuthenticationSha384: util.AsBool(o.AllowAuthenticationSha384),
	}

	return ans
}

type entry_v2 struct {
	XMLName                   xml.Name `xml:"entry"`
	Name                      string   `xml:"name,attr"`
	Certificate               string   `xml:"certificate"`
	MinVersion                string   `xml:"protocol-settings>min-version,omitempty"`
	MaxVersion                string   `xml:"protocol-settings>max-version,omitempty"`
	AllowAlgorithmRsa         string   `xml:"protocol-settings>keyxchg-algo-rsa,omitempty"`
	AllowAlgorithmDhe         string   `xml:"protocol-settings>keyxchg-algo-dhe,omitempty"`
	AllowAlgorithmEcdhe       string   `xml:"protocol-settings>keyxchg-algo-ecdhe,omitempty"`
	AllowAlgorithm3des        string   `xml:"protocol-settings>enc-algo-3des,omitempty"`
	AllowAlgorithmRc4         string   `xml:"protocol-settings>enc-algo-rc4,omitempty"`
	AllowAlgorithmAes128Cbc   string   `xml:"protocol-settings>enc-algo-aes-128-cbc,omitempty"`
	AllowAlgorithmAes256Cbc   string   `xml:"protocol-settings>enc-algo-aes-256-cbc,omitempty"`
	AllowAlgorithmAes128Gcm   string   `xml:"protocol-settings>enc-algo-aes-128-gcm,omitempty"`
	AllowAlgorithmAes256Gcm   string   `xml:"protocol-settings>enc-algo-aes-256-gcm,omitempty"`
	AllowAuthenticationSha1   string   `xml:"protocol-settings>auth-algo-sha1,omitempty"`
	AllowAuthenticationSha256 string   `xml:"protocol-settings>auth-algo-sha256,omitempty"`
	AllowAuthenticationSha384 string   `xml:"protocol-settings>auth-algo-sha384,omitempty"`
}

func (e *entry_v2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v2
	ans := local{
		MinVersion:                Tls1_0,
		MaxVersion:                TlsMax,
		AllowAlgorithmRsa:         util.YesNo(true),
		AllowAlgorithmDhe:         util.YesNo(true),
		AllowAlgorithmEcdhe:       util.YesNo(true),
		AllowAlgorithm3des:        util.YesNo(true),
		AllowAlgorithmRc4:         util.YesNo(true),
		AllowAlgorithmAes128Cbc:   util.YesNo(true),
		AllowAlgorithmAes256Cbc:   util.YesNo(true),
		AllowAlgorithmAes128Gcm:   util.YesNo(true),
		AllowAlgorithmAes256Gcm:   util.YesNo(true),
		AllowAuthenticationSha1:   util.YesNo(true),
		AllowAuthenticationSha256: util.YesNo(true),
		AllowAuthenticationSha384: util.YesNo(true),
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v2(ans)
	return nil
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:        e.Name,
		Certificate: e.Certificate,
		MinVersion:  e.MinVersion,
		MaxVersion:  e.MaxVersion,
	}

	if !e.AllowAlgorithmRsa {
		ans.AllowAlgorithmRsa = util.YesNo(false)
	}

	if !e.AllowAlgorithmDhe {
		ans.AllowAlgorithmDhe = util.YesNo(false)
	}

	if !e.AllowAlgorithmEcdhe {
		ans.AllowAlgorithmEcdhe = util.YesNo(false)
	}

	if !e.AllowAlgorithm3des {
		ans.AllowAlgorithm3des = util.YesNo(false)
	}

	if !e.AllowAlgorithmRc4 {
		ans.AllowAlgorithmRc4 = util.YesNo(false)
	}

	if !e.AllowAlgorithmAes128Cbc {
		ans.AllowAlgorithmAes128Cbc = util.YesNo(false)
	}

	if !e.AllowAlgorithmAes256Cbc {
		ans.AllowAlgorithmAes256Cbc = util.YesNo(false)
	}

	if !e.AllowAlgorithmAes128Gcm {
		ans.AllowAlgorithmAes128Gcm = util.YesNo(false)
	}

	if !e.AllowAlgorithmAes256Gcm {
		ans.AllowAlgorithmAes256Gcm = util.YesNo(false)
	}

	if !e.AllowAuthenticationSha1 {
		ans.AllowAuthenticationSha1 = util.YesNo(false)
	}

	if !e.AllowAuthenticationSha256 {
		ans.AllowAuthenticationSha256 = util.YesNo(false)
	}

	if !e.AllowAuthenticationSha384 {
		ans.AllowAuthenticationSha384 = util.YesNo(false)
	}

	return ans
}
