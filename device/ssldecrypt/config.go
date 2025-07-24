package ssldecrypt

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	DisabledSslExcludeCertFromPredefined []string
	ForwardTrustCertificateEcdsa         *string
	ForwardTrustCertificateRsa           *string
	ForwardUntrustCertificateEcdsa       *string
	ForwardUntrustCertificateRsa         *string
	RootCaExcludeList                    []string
	SslExcludeCert                       []SslExcludeCert
	TrustedRootCa                        []string
	Misc                                 []generic.Xml
	MiscAttributes                       []xml.Attr
}
type SslExcludeCert struct {
	Name           string
	Description    *string
	Exclude        *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}

func (o *configXmlContainer) Normalize() ([]*Config, error) {
	entries := make([]*Config, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyConfig(source *Config) (any, error) {
	var obj configXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type configXml struct {
	XMLName                              xml.Name                    `xml:"system"`
	DisabledSslExcludeCertFromPredefined *util.MemberType            `xml:"disabled-ssl-exclude-cert-from-predefined,omitempty"`
	ForwardTrustCertificateEcdsa         *string                     `xml:"forward-trust-certificate>ecdsa,omitempty"`
	ForwardTrustCertificateRsa           *string                     `xml:"forward-trust-certificate>rsa,omitempty"`
	ForwardUntrustCertificateEcdsa       *string                     `xml:"forward-untrust-certificate>ecdsa,omitempty"`
	ForwardUntrustCertificateRsa         *string                     `xml:"forward-untrust-certificate>rsa,omitempty"`
	RootCaExcludeList                    *util.MemberType            `xml:"root-ca-exclude-list,omitempty"`
	SslExcludeCert                       *sslExcludeCertContainerXml `xml:"ssl-exclude-cert,omitempty"`
	TrustedRootCa                        *util.MemberType            `xml:"trusted-root-CA,omitempty"`
	Misc                                 []generic.Xml               `xml:",any"`
	MiscAttributes                       []xml.Attr                  `xml:",any,attr"`
}
type sslExcludeCertContainerXml struct {
	Entries []sslExcludeCertXml `xml:"entry"`
}
type sslExcludeCertXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	Exclude        *string       `xml:"exclude,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *configXml) MarshalFromObject(s Config) {
	if s.DisabledSslExcludeCertFromPredefined != nil {
		o.DisabledSslExcludeCertFromPredefined = util.StrToMem(s.DisabledSslExcludeCertFromPredefined)
	}
	o.ForwardTrustCertificateEcdsa = s.ForwardTrustCertificateEcdsa
	o.ForwardTrustCertificateRsa = s.ForwardTrustCertificateRsa
	o.ForwardUntrustCertificateEcdsa = s.ForwardUntrustCertificateEcdsa
	o.ForwardUntrustCertificateRsa = s.ForwardUntrustCertificateRsa
	if s.RootCaExcludeList != nil {
		o.RootCaExcludeList = util.StrToMem(s.RootCaExcludeList)
	}
	if s.SslExcludeCert != nil {
		var objs []sslExcludeCertXml
		for _, elt := range s.SslExcludeCert {
			var obj sslExcludeCertXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SslExcludeCert = &sslExcludeCertContainerXml{Entries: objs}
	}
	if s.TrustedRootCa != nil {
		o.TrustedRootCa = util.StrToMem(s.TrustedRootCa)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var disabledSslExcludeCertFromPredefinedVal []string
	if o.DisabledSslExcludeCertFromPredefined != nil {
		disabledSslExcludeCertFromPredefinedVal = util.MemToStr(o.DisabledSslExcludeCertFromPredefined)
	}
	var rootCaExcludeListVal []string
	if o.RootCaExcludeList != nil {
		rootCaExcludeListVal = util.MemToStr(o.RootCaExcludeList)
	}
	var sslExcludeCertVal []SslExcludeCert
	if o.SslExcludeCert != nil {
		for _, elt := range o.SslExcludeCert.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			sslExcludeCertVal = append(sslExcludeCertVal, *obj)
		}
	}
	var trustedRootCaVal []string
	if o.TrustedRootCa != nil {
		trustedRootCaVal = util.MemToStr(o.TrustedRootCa)
	}

	result := &Config{
		DisabledSslExcludeCertFromPredefined: disabledSslExcludeCertFromPredefinedVal,
		ForwardTrustCertificateEcdsa:         o.ForwardTrustCertificateEcdsa,
		ForwardTrustCertificateRsa:           o.ForwardTrustCertificateRsa,
		ForwardUntrustCertificateEcdsa:       o.ForwardUntrustCertificateEcdsa,
		ForwardUntrustCertificateRsa:         o.ForwardUntrustCertificateRsa,
		RootCaExcludeList:                    rootCaExcludeListVal,
		SslExcludeCert:                       sslExcludeCertVal,
		TrustedRootCa:                        trustedRootCaVal,
		Misc:                                 o.Misc,
		MiscAttributes:                       o.MiscAttributes,
	}
	return result, nil
}
func (o *sslExcludeCertXml) MarshalFromObject(s SslExcludeCert) {
	o.Name = s.Name
	o.Description = s.Description
	o.Exclude = util.YesNo(s.Exclude, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslExcludeCertXml) UnmarshalToObject() (*SslExcludeCert, error) {

	result := &SslExcludeCert{
		Name:           o.Name,
		Description:    o.Description,
		Exclude:        util.AsBool(o.Exclude, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func SpecMatches(a, b *Config) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Config) matches(other *Config) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DisabledSslExcludeCertFromPredefined, other.DisabledSslExcludeCertFromPredefined) {
		return false
	}
	if !util.StringsMatch(o.ForwardTrustCertificateEcdsa, other.ForwardTrustCertificateEcdsa) {
		return false
	}
	if !util.StringsMatch(o.ForwardTrustCertificateRsa, other.ForwardTrustCertificateRsa) {
		return false
	}
	if !util.StringsMatch(o.ForwardUntrustCertificateEcdsa, other.ForwardUntrustCertificateEcdsa) {
		return false
	}
	if !util.StringsMatch(o.ForwardUntrustCertificateRsa, other.ForwardUntrustCertificateRsa) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RootCaExcludeList, other.RootCaExcludeList) {
		return false
	}
	if len(o.SslExcludeCert) != len(other.SslExcludeCert) {
		return false
	}
	for idx := range o.SslExcludeCert {
		if !o.SslExcludeCert[idx].matches(&other.SslExcludeCert[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.TrustedRootCa, other.TrustedRootCa) {
		return false
	}

	return true
}

func (o *SslExcludeCert) matches(other *SslExcludeCert) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.BoolsMatch(o.Exclude, other.Exclude) {
		return false
	}

	return true
}
