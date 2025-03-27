package ssltls

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	Suffix = []string{"ssl-tls-service-profile"}
)

type Entry struct {
	Name             string
	Certificate      *string
	ProtocolSettings *ProtocolSettings

	Misc map[string][]generic.Xml
}

type ProtocolSettings struct {
	AllowAuthenticationSha1   *bool
	AllowAuthenticationSha256 *bool
	AllowAuthenticationSha384 *bool
	AllowAlgorithm3des        *bool
	AllowAlgorithmAes128Cbc   *bool
	AllowAlgorithmAes128Gcm   *bool
	AllowAlgorithmAes256Cbc   *bool
	AllowAlgorithmAes256Gcm   *bool
	AllowAlgorithmRc4         *bool
	AllowAlgorithmDhe         *bool
	AllowAlgorithmEcdhe       *bool
	AllowAlgorithmRsa         *bool
	MaxVersion                *string
	MinVersion                *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName          xml.Name             `xml:"entry"`
	Name             string               `xml:"name,attr"`
	Certificate      *string              `xml:"certificate,omitempty"`
	ProtocolSettings *ProtocolSettingsXml `xml:"protocol-settings,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolSettingsXml struct {
	AllowAlgorithm3des        *string `xml:"enc-algo-3des,omitempty"`
	AllowAlgorithmAes128Cbc   *string `xml:"enc-algo-aes-128-cbc,omitempty"`
	AllowAlgorithmAes128Gcm   *string `xml:"enc-algo-aes-128-gcm,omitempty"`
	AllowAlgorithmAes256Cbc   *string `xml:"enc-algo-aes-256-cbc,omitempty"`
	AllowAlgorithmAes256Gcm   *string `xml:"enc-algo-aes-256-gcm,omitempty"`
	AllowAlgorithmDhe         *string `xml:"keyxchg-algo-dhe,omitempty"`
	AllowAlgorithmEcdhe       *string `xml:"keyxchg-algo-ecdhe,omitempty"`
	AllowAlgorithmRc4         *string `xml:"enc-algo-rc4,omitempty"`
	AllowAlgorithmRsa         *string `xml:"keyxchg-algo-rsa,omitempty"`
	AllowAuthenticationSha1   *string `xml:"auth-algo-sha1,omitempty"`
	AllowAuthenticationSha256 *string `xml:"auth-algo-sha256,omitempty"`
	AllowAuthenticationSha384 *string `xml:"auth-algo-sha384,omitempty"`
	MaxVersion                *string `xml:"max-version,omitempty"`
	MinVersion                *string `xml:"min-version,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Certificate = o.Certificate
	var nestedProtocolSettings *ProtocolSettingsXml
	if o.ProtocolSettings != nil {
		nestedProtocolSettings = &ProtocolSettingsXml{}
		if _, ok := o.Misc["ProtocolSettings"]; ok {
			nestedProtocolSettings.Misc = o.Misc["ProtocolSettings"]
		}
		if o.ProtocolSettings.AllowAuthenticationSha1 != nil {
			nestedProtocolSettings.AllowAuthenticationSha1 = util.YesNo(o.ProtocolSettings.AllowAuthenticationSha1, nil)
		}
		if o.ProtocolSettings.AllowAuthenticationSha256 != nil {
			nestedProtocolSettings.AllowAuthenticationSha256 = util.YesNo(o.ProtocolSettings.AllowAuthenticationSha256, nil)
		}
		if o.ProtocolSettings.AllowAuthenticationSha384 != nil {
			nestedProtocolSettings.AllowAuthenticationSha384 = util.YesNo(o.ProtocolSettings.AllowAuthenticationSha384, nil)
		}
		if o.ProtocolSettings.AllowAlgorithm3des != nil {
			nestedProtocolSettings.AllowAlgorithm3des = util.YesNo(o.ProtocolSettings.AllowAlgorithm3des, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmAes128Cbc != nil {
			nestedProtocolSettings.AllowAlgorithmAes128Cbc = util.YesNo(o.ProtocolSettings.AllowAlgorithmAes128Cbc, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmAes128Gcm != nil {
			nestedProtocolSettings.AllowAlgorithmAes128Gcm = util.YesNo(o.ProtocolSettings.AllowAlgorithmAes128Gcm, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmAes256Cbc != nil {
			nestedProtocolSettings.AllowAlgorithmAes256Cbc = util.YesNo(o.ProtocolSettings.AllowAlgorithmAes256Cbc, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmAes256Gcm != nil {
			nestedProtocolSettings.AllowAlgorithmAes256Gcm = util.YesNo(o.ProtocolSettings.AllowAlgorithmAes256Gcm, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmRc4 != nil {
			nestedProtocolSettings.AllowAlgorithmRc4 = util.YesNo(o.ProtocolSettings.AllowAlgorithmRc4, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmDhe != nil {
			nestedProtocolSettings.AllowAlgorithmDhe = util.YesNo(o.ProtocolSettings.AllowAlgorithmDhe, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmEcdhe != nil {
			nestedProtocolSettings.AllowAlgorithmEcdhe = util.YesNo(o.ProtocolSettings.AllowAlgorithmEcdhe, nil)
		}
		if o.ProtocolSettings.AllowAlgorithmRsa != nil {
			nestedProtocolSettings.AllowAlgorithmRsa = util.YesNo(o.ProtocolSettings.AllowAlgorithmRsa, nil)
		}
		if o.ProtocolSettings.MaxVersion != nil {
			nestedProtocolSettings.MaxVersion = o.ProtocolSettings.MaxVersion
		}
		if o.ProtocolSettings.MinVersion != nil {
			nestedProtocolSettings.MinVersion = o.ProtocolSettings.MinVersion
		}
	}
	entry.ProtocolSettings = nestedProtocolSettings

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		entry.Certificate = o.Certificate
		var nestedProtocolSettings *ProtocolSettings
		if o.ProtocolSettings != nil {
			nestedProtocolSettings = &ProtocolSettings{}
			if o.ProtocolSettings.Misc != nil {
				entry.Misc["ProtocolSettings"] = o.ProtocolSettings.Misc
			}
			if o.ProtocolSettings.AllowAuthenticationSha1 != nil {
				nestedProtocolSettings.AllowAuthenticationSha1 = util.AsBool(o.ProtocolSettings.AllowAuthenticationSha1, nil)
			}
			if o.ProtocolSettings.AllowAuthenticationSha256 != nil {
				nestedProtocolSettings.AllowAuthenticationSha256 = util.AsBool(o.ProtocolSettings.AllowAuthenticationSha256, nil)
			}
			if o.ProtocolSettings.AllowAuthenticationSha384 != nil {
				nestedProtocolSettings.AllowAuthenticationSha384 = util.AsBool(o.ProtocolSettings.AllowAuthenticationSha384, nil)
			}
			if o.ProtocolSettings.AllowAlgorithm3des != nil {
				nestedProtocolSettings.AllowAlgorithm3des = util.AsBool(o.ProtocolSettings.AllowAlgorithm3des, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmAes128Cbc != nil {
				nestedProtocolSettings.AllowAlgorithmAes128Cbc = util.AsBool(o.ProtocolSettings.AllowAlgorithmAes128Cbc, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmAes128Gcm != nil {
				nestedProtocolSettings.AllowAlgorithmAes128Gcm = util.AsBool(o.ProtocolSettings.AllowAlgorithmAes128Gcm, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmAes256Cbc != nil {
				nestedProtocolSettings.AllowAlgorithmAes256Cbc = util.AsBool(o.ProtocolSettings.AllowAlgorithmAes256Cbc, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmAes256Gcm != nil {
				nestedProtocolSettings.AllowAlgorithmAes256Gcm = util.AsBool(o.ProtocolSettings.AllowAlgorithmAes256Gcm, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmRc4 != nil {
				nestedProtocolSettings.AllowAlgorithmRc4 = util.AsBool(o.ProtocolSettings.AllowAlgorithmRc4, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmDhe != nil {
				nestedProtocolSettings.AllowAlgorithmDhe = util.AsBool(o.ProtocolSettings.AllowAlgorithmDhe, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmEcdhe != nil {
				nestedProtocolSettings.AllowAlgorithmEcdhe = util.AsBool(o.ProtocolSettings.AllowAlgorithmEcdhe, nil)
			}
			if o.ProtocolSettings.AllowAlgorithmRsa != nil {
				nestedProtocolSettings.AllowAlgorithmRsa = util.AsBool(o.ProtocolSettings.AllowAlgorithmRsa, nil)
			}
			if o.ProtocolSettings.MaxVersion != nil {
				nestedProtocolSettings.MaxVersion = o.ProtocolSettings.MaxVersion
			}
			if o.ProtocolSettings.MinVersion != nil {
				nestedProtocolSettings.MinVersion = o.ProtocolSettings.MinVersion
			}
		}
		entry.ProtocolSettings = nestedProtocolSettings

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !util.StringsMatch(a.Certificate, b.Certificate) {
		return false
	}
	if !matchProtocolSettings(a.ProtocolSettings, b.ProtocolSettings) {
		return false
	}

	return true
}

func matchProtocolSettings(a *ProtocolSettings, b *ProtocolSettings) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AllowAuthenticationSha1, b.AllowAuthenticationSha1) {
		return false
	}
	if !util.BoolsMatch(a.AllowAuthenticationSha256, b.AllowAuthenticationSha256) {
		return false
	}
	if !util.BoolsMatch(a.AllowAuthenticationSha384, b.AllowAuthenticationSha384) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithm3des, b.AllowAlgorithm3des) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmAes128Cbc, b.AllowAlgorithmAes128Cbc) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmAes128Gcm, b.AllowAlgorithmAes128Gcm) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmAes256Cbc, b.AllowAlgorithmAes256Cbc) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmAes256Gcm, b.AllowAlgorithmAes256Gcm) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmRc4, b.AllowAlgorithmRc4) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmDhe, b.AllowAlgorithmDhe) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmEcdhe, b.AllowAlgorithmEcdhe) {
		return false
	}
	if !util.BoolsMatch(a.AllowAlgorithmRsa, b.AllowAlgorithmRsa) {
		return false
	}
	if !util.StringsMatch(a.MaxVersion, b.MaxVersion) {
		return false
	}
	if !util.StringsMatch(a.MinVersion, b.MinVersion) {
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
