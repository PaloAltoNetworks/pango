package ssldecrypt

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	DisabledSslExcludeCertFromPredefined []string
	RootCaExcludeList                    []string
	SslExcludeCert                       []SslExcludeCert
	TrustedRootCa                        []string
	ForwardTrustCertificateEcdsa         *string
	ForwardTrustCertificateRsa           *string
	ForwardUntrustCertificateEcdsa       *string
	ForwardUntrustCertificateRsa         *string

	Misc map[string][]generic.Xml
}
type SslExcludeCert struct {
	Name        string
	Description *string
	Exclude     *bool
}
type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}
type configXml struct {
	XMLName                              xml.Name            `xml:"system"`
	DisabledSslExcludeCertFromPredefined *util.MemberType    `xml:"disabled-ssl-exclude-cert-from-predefined,omitempty"`
	ForwardTrustCertificateEcdsa         *string             `xml:"forward-trust-certificate>ecdsa,omitempty"`
	ForwardTrustCertificateRsa           *string             `xml:"forward-trust-certificate>rsa,omitempty"`
	ForwardUntrustCertificateEcdsa       *string             `xml:"forward-untrust-certificate>ecdsa,omitempty"`
	ForwardUntrustCertificateRsa         *string             `xml:"forward-untrust-certificate>rsa,omitempty"`
	RootCaExcludeList                    *util.MemberType    `xml:"root-ca-exclude-list,omitempty"`
	SslExcludeCert                       []SslExcludeCertXml `xml:"ssl-exclude-cert>entry,omitempty"`
	TrustedRootCa                        *util.MemberType    `xml:"trusted-root-CA,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SslExcludeCertXml struct {
	Description *string `xml:"description,omitempty"`
	Exclude     *string `xml:"exclude,omitempty"`
	Name        string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func specifyConfig(o *Config) (any, error) {
	config := configXml{}
	config.DisabledSslExcludeCertFromPredefined = util.StrToMem(o.DisabledSslExcludeCertFromPredefined)
	config.ForwardTrustCertificateEcdsa = o.ForwardTrustCertificateEcdsa
	config.ForwardTrustCertificateRsa = o.ForwardTrustCertificateRsa
	config.ForwardUntrustCertificateEcdsa = o.ForwardUntrustCertificateEcdsa
	config.ForwardUntrustCertificateRsa = o.ForwardUntrustCertificateRsa
	config.RootCaExcludeList = util.StrToMem(o.RootCaExcludeList)
	var nestedSslExcludeCertCol []SslExcludeCertXml
	if o.SslExcludeCert != nil {
		nestedSslExcludeCertCol = []SslExcludeCertXml{}
		for _, oSslExcludeCert := range o.SslExcludeCert {
			nestedSslExcludeCert := SslExcludeCertXml{}
			if _, ok := o.Misc["SslExcludeCert"]; ok {
				nestedSslExcludeCert.Misc = o.Misc["SslExcludeCert"]
			}
			if oSslExcludeCert.Name != "" {
				nestedSslExcludeCert.Name = oSslExcludeCert.Name
			}
			if oSslExcludeCert.Description != nil {
				nestedSslExcludeCert.Description = oSslExcludeCert.Description
			}
			if oSslExcludeCert.Exclude != nil {
				nestedSslExcludeCert.Exclude = util.YesNo(oSslExcludeCert.Exclude, nil)
			}
			nestedSslExcludeCertCol = append(nestedSslExcludeCertCol, nestedSslExcludeCert)
		}
		config.SslExcludeCert = nestedSslExcludeCertCol
	}

	config.TrustedRootCa = util.StrToMem(o.TrustedRootCa)

	config.Misc = o.Misc["Config"]

	return config, nil
}
func (c *configXmlContainer) Normalize() ([]*Config, error) {
	configList := make([]*Config, 0, len(c.Answer))
	for _, o := range c.Answer {
		config := &Config{
			Misc: make(map[string][]generic.Xml),
		}
		config.DisabledSslExcludeCertFromPredefined = util.MemToStr(o.DisabledSslExcludeCertFromPredefined)
		config.ForwardTrustCertificateEcdsa = o.ForwardTrustCertificateEcdsa
		config.ForwardTrustCertificateRsa = o.ForwardTrustCertificateRsa
		config.ForwardUntrustCertificateEcdsa = o.ForwardUntrustCertificateEcdsa
		config.ForwardUntrustCertificateRsa = o.ForwardUntrustCertificateRsa
		config.RootCaExcludeList = util.MemToStr(o.RootCaExcludeList)
		var nestedSslExcludeCertCol []SslExcludeCert
		if o.SslExcludeCert != nil {
			nestedSslExcludeCertCol = []SslExcludeCert{}
			for _, oSslExcludeCert := range o.SslExcludeCert {
				nestedSslExcludeCert := SslExcludeCert{}
				if oSslExcludeCert.Misc != nil {
					config.Misc["SslExcludeCert"] = oSslExcludeCert.Misc
				}
				if oSslExcludeCert.Name != "" {
					nestedSslExcludeCert.Name = oSslExcludeCert.Name
				}
				if oSslExcludeCert.Description != nil {
					nestedSslExcludeCert.Description = oSslExcludeCert.Description
				}
				if oSslExcludeCert.Exclude != nil {
					nestedSslExcludeCert.Exclude = util.AsBool(oSslExcludeCert.Exclude, nil)
				}
				nestedSslExcludeCertCol = append(nestedSslExcludeCertCol, nestedSslExcludeCert)
			}
			config.SslExcludeCert = nestedSslExcludeCertCol
		}

		config.TrustedRootCa = util.MemToStr(o.TrustedRootCa)

		config.Misc["Config"] = o.Misc

		configList = append(configList, config)
	}

	return configList, nil
}

func SpecMatches(a, b *Config) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !util.OrderedListsMatch(a.DisabledSslExcludeCertFromPredefined, b.DisabledSslExcludeCertFromPredefined) {
		return false
	}
	if !util.OrderedListsMatch(a.RootCaExcludeList, b.RootCaExcludeList) {
		return false
	}
	if !matchSslExcludeCert(a.SslExcludeCert, b.SslExcludeCert) {
		return false
	}
	if !util.OrderedListsMatch(a.TrustedRootCa, b.TrustedRootCa) {
		return false
	}
	if !util.StringsMatch(a.ForwardTrustCertificateEcdsa, b.ForwardTrustCertificateEcdsa) {
		return false
	}
	if !util.StringsMatch(a.ForwardTrustCertificateRsa, b.ForwardTrustCertificateRsa) {
		return false
	}
	if !util.StringsMatch(a.ForwardUntrustCertificateEcdsa, b.ForwardUntrustCertificateEcdsa) {
		return false
	}
	if !util.StringsMatch(a.ForwardUntrustCertificateRsa, b.ForwardUntrustCertificateRsa) {
		return false
	}

	return true
}

func matchSslExcludeCert(a []SslExcludeCert, b []SslExcludeCert) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
			if !util.BoolsMatch(a.Exclude, b.Exclude) {
				return false
			}
		}
	}
	return true
}
