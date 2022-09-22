package saml

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// SAML identity provider profile.
//
// PAN-OS 8.0+
type Entry struct {
	Name                                string
	AdminUseOnly                        bool
	IdentityProviderId                  string
	IdentityProviderCertificate         string // Required in PAN-OS 10.0+
	SsoUrl                              string
	SsoBinding                          string
	SloUrl                              string
	SloBinding                          string
	ValidateIdentityProviderCertificate bool
	SignSamlMessage                     bool
	MaxClockSkew                        int
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.AdminUseOnly = s.AdminUseOnly
	o.IdentityProviderId = s.IdentityProviderId
	o.IdentityProviderCertificate = s.IdentityProviderCertificate
	o.SsoUrl = s.SsoUrl
	o.SsoBinding = s.SsoBinding
	o.SloUrl = s.SloUrl
	o.SloBinding = s.SloBinding
	o.ValidateIdentityProviderCertificate = s.ValidateIdentityProviderCertificate
	o.SignSamlMessage = s.SignSamlMessage
	o.MaxClockSkew = s.MaxClockSkew
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

type entry_v1 struct {
	XMLName                             xml.Name `xml:"entry"`
	Name                                string   `xml:"name,attr"`
	AdminUseOnly                        string   `xml:"admin-use-only"`
	IdentityProviderId                  string   `xml:"entity-id"`
	IdentityProviderCertificate         string   `xml:"certificate"`
	SsoUrl                              string   `xml:"sso-url"`
	SsoBinding                          string   `xml:"sso-bindings"`
	SloUrl                              string   `xml:"slo-url,omitempty"`
	SloBinding                          string   `xml:"slo-bindings,omitempty"`
	ValidateIdentityProviderCertificate string   `xml:"validate-idp-certificate"`
	SignSamlMessage                     string   `xml:"want-auth-requests-signed"`
	MaxClockSkew                        int      `xml:"max-clock-skew,omitempty"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{
		SloBinding:                          BindingPost,
		ValidateIdentityProviderCertificate: util.YesNo(true),
		SignSamlMessage:                     util.YesNo(true),
		MaxClockSkew:                        60,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v1(ans)
	return nil
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:                                o.Name,
		AdminUseOnly:                        util.AsBool(o.AdminUseOnly),
		IdentityProviderId:                  o.IdentityProviderId,
		IdentityProviderCertificate:         o.IdentityProviderCertificate,
		SsoUrl:                              o.SsoUrl,
		SsoBinding:                          o.SsoBinding,
		SloUrl:                              o.SloUrl,
		SloBinding:                          o.SloBinding,
		ValidateIdentityProviderCertificate: util.AsBool(o.ValidateIdentityProviderCertificate),
		SignSamlMessage:                     util.AsBool(o.SignSamlMessage),
		MaxClockSkew:                        o.MaxClockSkew,
	}

	return ans
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:                                e.Name,
		AdminUseOnly:                        util.YesNo(e.AdminUseOnly),
		IdentityProviderId:                  e.IdentityProviderId,
		IdentityProviderCertificate:         e.IdentityProviderCertificate,
		SsoUrl:                              e.SsoUrl,
		SsoBinding:                          e.SsoBinding,
		SloUrl:                              e.SloUrl,
		SloBinding:                          e.SloBinding,
		ValidateIdentityProviderCertificate: util.YesNo(e.ValidateIdentityProviderCertificate),
		SignSamlMessage:                     util.YesNo(e.SignSamlMessage),
		MaxClockSkew:                        e.MaxClockSkew,
	}

	return ans
}
