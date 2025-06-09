package ssltls

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"ssl-tls-service-profile", "$name"}
)

type Entry struct {
	Name             string
	Certificate      *string
	ProtocolSettings *ProtocolSettings
	Misc             []generic.Xml
}
type ProtocolSettings struct {
	AllowAlgorithm3des        *bool
	AllowAlgorithmAes128Cbc   *bool
	AllowAlgorithmAes128Gcm   *bool
	AllowAlgorithmAes256Cbc   *bool
	AllowAlgorithmAes256Gcm   *bool
	AllowAlgorithmDhe         *bool
	AllowAlgorithmEcdhe       *bool
	AllowAlgorithmRc4         *bool
	AllowAlgorithmRsa         *bool
	AllowAuthenticationSha1   *bool
	AllowAuthenticationSha256 *bool
	AllowAuthenticationSha384 *bool
	MaxVersion                *string
	MinVersion                *string
	Misc                      []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName          xml.Name             `xml:"entry"`
	Name             string               `xml:"name,attr"`
	Certificate      *string              `xml:"certificate,omitempty"`
	ProtocolSettings *protocolSettingsXml `xml:"protocol-settings,omitempty"`
	Misc             []generic.Xml        `xml:",any"`
}
type protocolSettingsXml struct {
	AllowAlgorithm3des        *string       `xml:"enc-algo-3des,omitempty"`
	AllowAlgorithmAes128Cbc   *string       `xml:"enc-algo-aes-128-cbc,omitempty"`
	AllowAlgorithmAes128Gcm   *string       `xml:"enc-algo-aes-128-gcm,omitempty"`
	AllowAlgorithmAes256Cbc   *string       `xml:"enc-algo-aes-256-cbc,omitempty"`
	AllowAlgorithmAes256Gcm   *string       `xml:"enc-algo-aes-256-gcm,omitempty"`
	AllowAlgorithmDhe         *string       `xml:"keyxchg-algo-dhe,omitempty"`
	AllowAlgorithmEcdhe       *string       `xml:"keyxchg-algo-ecdhe,omitempty"`
	AllowAlgorithmRc4         *string       `xml:"enc-algo-rc4,omitempty"`
	AllowAlgorithmRsa         *string       `xml:"keyxchg-algo-rsa,omitempty"`
	AllowAuthenticationSha1   *string       `xml:"auth-algo-sha1,omitempty"`
	AllowAuthenticationSha256 *string       `xml:"auth-algo-sha256,omitempty"`
	AllowAuthenticationSha384 *string       `xml:"auth-algo-sha384,omitempty"`
	MaxVersion                *string       `xml:"max-version,omitempty"`
	MinVersion                *string       `xml:"min-version,omitempty"`
	Misc                      []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Certificate = s.Certificate
	if s.ProtocolSettings != nil {
		var obj protocolSettingsXml
		obj.MarshalFromObject(*s.ProtocolSettings)
		o.ProtocolSettings = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var protocolSettingsVal *ProtocolSettings
	if o.ProtocolSettings != nil {
		obj, err := o.ProtocolSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolSettingsVal = obj
	}

	result := &Entry{
		Name:             o.Name,
		Certificate:      o.Certificate,
		ProtocolSettings: protocolSettingsVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolSettingsXml) MarshalFromObject(s ProtocolSettings) {
	o.AllowAlgorithm3des = util.YesNo(s.AllowAlgorithm3des, nil)
	o.AllowAlgorithmAes128Cbc = util.YesNo(s.AllowAlgorithmAes128Cbc, nil)
	o.AllowAlgorithmAes128Gcm = util.YesNo(s.AllowAlgorithmAes128Gcm, nil)
	o.AllowAlgorithmAes256Cbc = util.YesNo(s.AllowAlgorithmAes256Cbc, nil)
	o.AllowAlgorithmAes256Gcm = util.YesNo(s.AllowAlgorithmAes256Gcm, nil)
	o.AllowAlgorithmDhe = util.YesNo(s.AllowAlgorithmDhe, nil)
	o.AllowAlgorithmEcdhe = util.YesNo(s.AllowAlgorithmEcdhe, nil)
	o.AllowAlgorithmRc4 = util.YesNo(s.AllowAlgorithmRc4, nil)
	o.AllowAlgorithmRsa = util.YesNo(s.AllowAlgorithmRsa, nil)
	o.AllowAuthenticationSha1 = util.YesNo(s.AllowAuthenticationSha1, nil)
	o.AllowAuthenticationSha256 = util.YesNo(s.AllowAuthenticationSha256, nil)
	o.AllowAuthenticationSha384 = util.YesNo(s.AllowAuthenticationSha384, nil)
	o.MaxVersion = s.MaxVersion
	o.MinVersion = s.MinVersion
	o.Misc = s.Misc
}

func (o protocolSettingsXml) UnmarshalToObject() (*ProtocolSettings, error) {

	result := &ProtocolSettings{
		AllowAlgorithm3des:        util.AsBool(o.AllowAlgorithm3des, nil),
		AllowAlgorithmAes128Cbc:   util.AsBool(o.AllowAlgorithmAes128Cbc, nil),
		AllowAlgorithmAes128Gcm:   util.AsBool(o.AllowAlgorithmAes128Gcm, nil),
		AllowAlgorithmAes256Cbc:   util.AsBool(o.AllowAlgorithmAes256Cbc, nil),
		AllowAlgorithmAes256Gcm:   util.AsBool(o.AllowAlgorithmAes256Gcm, nil),
		AllowAlgorithmDhe:         util.AsBool(o.AllowAlgorithmDhe, nil),
		AllowAlgorithmEcdhe:       util.AsBool(o.AllowAlgorithmEcdhe, nil),
		AllowAlgorithmRc4:         util.AsBool(o.AllowAlgorithmRc4, nil),
		AllowAlgorithmRsa:         util.AsBool(o.AllowAlgorithmRsa, nil),
		AllowAuthenticationSha1:   util.AsBool(o.AllowAuthenticationSha1, nil),
		AllowAuthenticationSha256: util.AsBool(o.AllowAuthenticationSha256, nil),
		AllowAuthenticationSha384: util.AsBool(o.AllowAuthenticationSha384, nil),
		MaxVersion:                o.MaxVersion,
		MinVersion:                o.MinVersion,
		Misc:                      o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "certificate" || v == "Certificate" {
		return e.Certificate, nil
	}
	if v == "protocol_settings" || v == "ProtocolSettings" {
		return e.ProtocolSettings, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Certificate, other.Certificate) {
		return false
	}
	if !o.ProtocolSettings.matches(other.ProtocolSettings) {
		return false
	}

	return true
}

func (o *ProtocolSettings) matches(other *ProtocolSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithm3des, other.AllowAlgorithm3des) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmAes128Cbc, other.AllowAlgorithmAes128Cbc) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmAes128Gcm, other.AllowAlgorithmAes128Gcm) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmAes256Cbc, other.AllowAlgorithmAes256Cbc) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmAes256Gcm, other.AllowAlgorithmAes256Gcm) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmDhe, other.AllowAlgorithmDhe) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmEcdhe, other.AllowAlgorithmEcdhe) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmRc4, other.AllowAlgorithmRc4) {
		return false
	}
	if !util.BoolsMatch(o.AllowAlgorithmRsa, other.AllowAlgorithmRsa) {
		return false
	}
	if !util.BoolsMatch(o.AllowAuthenticationSha1, other.AllowAuthenticationSha1) {
		return false
	}
	if !util.BoolsMatch(o.AllowAuthenticationSha256, other.AllowAuthenticationSha256) {
		return false
	}
	if !util.BoolsMatch(o.AllowAuthenticationSha384, other.AllowAuthenticationSha384) {
		return false
	}
	if !util.StringsMatch(o.MaxVersion, other.MaxVersion) {
		return false
	}
	if !util.StringsMatch(o.MinVersion, other.MinVersion) {
		return false
	}

	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
