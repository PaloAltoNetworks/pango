package samlidp

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
	suffix = []string{"server-profile", "saml-idp", "$name"}
)

type Entry struct {
	Name                            string
	AdminUseOnly                    *bool
	AttributeNameAccessDomainImport *string
	AttributeNameAdminRoleImport    *string
	AttributeNameUsergroupImport    *string
	AttributeNameUsernameImport     *string
	Certificate                     *string
	EntityId                        *string
	MaxClockSkew                    *int64
	SloBindings                     *string
	SloUrl                          *string
	SsoBindings                     *string
	SsoUrl                          *string
	ValidateIdpCertificate          *bool
	WantAuthRequestsSigned          *bool
	Misc                            []generic.Xml
	MiscAttributes                  []xml.Attr
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
	XMLName                         xml.Name      `xml:"entry"`
	Name                            string        `xml:"name,attr"`
	AdminUseOnly                    *string       `xml:"admin-use-only,omitempty"`
	AttributeNameAccessDomainImport *string       `xml:"attribute-name-access-domain-import,omitempty"`
	AttributeNameAdminRoleImport    *string       `xml:"attribute-name-admin-role-import,omitempty"`
	AttributeNameUsergroupImport    *string       `xml:"attribute-name-usergroup-import,omitempty"`
	AttributeNameUsernameImport     *string       `xml:"attribute-name-username-import,omitempty"`
	Certificate                     *string       `xml:"certificate,omitempty"`
	EntityId                        *string       `xml:"entity-id,omitempty"`
	MaxClockSkew                    *int64        `xml:"max-clock-skew,omitempty"`
	SloBindings                     *string       `xml:"slo-bindings,omitempty"`
	SloUrl                          *string       `xml:"slo-url,omitempty"`
	SsoBindings                     *string       `xml:"sso-bindings,omitempty"`
	SsoUrl                          *string       `xml:"sso-url,omitempty"`
	ValidateIdpCertificate          *string       `xml:"validate-idp-certificate,omitempty"`
	WantAuthRequestsSigned          *string       `xml:"want-auth-requests-signed,omitempty"`
	Misc                            []generic.Xml `xml:",any"`
	MiscAttributes                  []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AdminUseOnly = util.YesNo(s.AdminUseOnly, nil)
	o.AttributeNameAccessDomainImport = s.AttributeNameAccessDomainImport
	o.AttributeNameAdminRoleImport = s.AttributeNameAdminRoleImport
	o.AttributeNameUsergroupImport = s.AttributeNameUsergroupImport
	o.AttributeNameUsernameImport = s.AttributeNameUsernameImport
	o.Certificate = s.Certificate
	o.EntityId = s.EntityId
	o.MaxClockSkew = s.MaxClockSkew
	o.SloBindings = s.SloBindings
	o.SloUrl = s.SloUrl
	o.SsoBindings = s.SsoBindings
	o.SsoUrl = s.SsoUrl
	o.ValidateIdpCertificate = util.YesNo(s.ValidateIdpCertificate, nil)
	o.WantAuthRequestsSigned = util.YesNo(s.WantAuthRequestsSigned, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:                            o.Name,
		AdminUseOnly:                    util.AsBool(o.AdminUseOnly, nil),
		AttributeNameAccessDomainImport: o.AttributeNameAccessDomainImport,
		AttributeNameAdminRoleImport:    o.AttributeNameAdminRoleImport,
		AttributeNameUsergroupImport:    o.AttributeNameUsergroupImport,
		AttributeNameUsernameImport:     o.AttributeNameUsernameImport,
		Certificate:                     o.Certificate,
		EntityId:                        o.EntityId,
		MaxClockSkew:                    o.MaxClockSkew,
		SloBindings:                     o.SloBindings,
		SloUrl:                          o.SloUrl,
		SsoBindings:                     o.SsoBindings,
		SsoUrl:                          o.SsoUrl,
		ValidateIdpCertificate:          util.AsBool(o.ValidateIdpCertificate, nil),
		WantAuthRequestsSigned:          util.AsBool(o.WantAuthRequestsSigned, nil),
		Misc:                            o.Misc,
		MiscAttributes:                  o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "admin_use_only" || v == "AdminUseOnly" {
		return e.AdminUseOnly, nil
	}
	if v == "attribute_name_access_domain_import" || v == "AttributeNameAccessDomainImport" {
		return e.AttributeNameAccessDomainImport, nil
	}
	if v == "attribute_name_admin_role_import" || v == "AttributeNameAdminRoleImport" {
		return e.AttributeNameAdminRoleImport, nil
	}
	if v == "attribute_name_usergroup_import" || v == "AttributeNameUsergroupImport" {
		return e.AttributeNameUsergroupImport, nil
	}
	if v == "attribute_name_username_import" || v == "AttributeNameUsernameImport" {
		return e.AttributeNameUsernameImport, nil
	}
	if v == "certificate" || v == "Certificate" {
		return e.Certificate, nil
	}
	if v == "entity_id" || v == "EntityId" {
		return e.EntityId, nil
	}
	if v == "max_clock_skew" || v == "MaxClockSkew" {
		return e.MaxClockSkew, nil
	}
	if v == "slo_bindings" || v == "SloBindings" {
		return e.SloBindings, nil
	}
	if v == "slo_url" || v == "SloUrl" {
		return e.SloUrl, nil
	}
	if v == "sso_bindings" || v == "SsoBindings" {
		return e.SsoBindings, nil
	}
	if v == "sso_url" || v == "SsoUrl" {
		return e.SsoUrl, nil
	}
	if v == "validate_idp_certificate" || v == "ValidateIdpCertificate" {
		return e.ValidateIdpCertificate, nil
	}
	if v == "want_auth_requests_signed" || v == "WantAuthRequestsSigned" {
		return e.WantAuthRequestsSigned, nil
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
	if !util.BoolsMatch(o.AdminUseOnly, other.AdminUseOnly) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameAccessDomainImport, other.AttributeNameAccessDomainImport) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameAdminRoleImport, other.AttributeNameAdminRoleImport) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameUsergroupImport, other.AttributeNameUsergroupImport) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameUsernameImport, other.AttributeNameUsernameImport) {
		return false
	}
	if !util.StringsMatch(o.Certificate, other.Certificate) {
		return false
	}
	if !util.StringsMatch(o.EntityId, other.EntityId) {
		return false
	}
	if !util.Ints64Match(o.MaxClockSkew, other.MaxClockSkew) {
		return false
	}
	if !util.StringsMatch(o.SloBindings, other.SloBindings) {
		return false
	}
	if !util.StringsMatch(o.SloUrl, other.SloUrl) {
		return false
	}
	if !util.StringsMatch(o.SsoBindings, other.SsoBindings) {
		return false
	}
	if !util.StringsMatch(o.SsoUrl, other.SsoUrl) {
		return false
	}
	if !util.BoolsMatch(o.ValidateIdpCertificate, other.ValidateIdpCertificate) {
		return false
	}
	if !util.BoolsMatch(o.WantAuthRequestsSigned, other.WantAuthRequestsSigned) {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
