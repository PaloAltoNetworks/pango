package certificate

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
	suffix = []string{"certificate", "$name"}
)

type Entry struct {
	Name                            string
	BlockExpiredCertificate         *bool
	BlockTimeoutCertificate         *bool
	BlockUnauthenticatedCertificate *bool
	BlockUnknownCertificate         *bool
	Certificate                     []Certificate
	CertificateStatusTimeout        *int64
	CrlReceiveTimeout               *int64
	Domain                          *string
	OcspExcludeNonce                *bool
	OcspReceiveTimeout              *int64
	UseCrl                          *bool
	UseOcsp                         *bool
	UsernameField                   *UsernameField
	Misc                            []generic.Xml
}
type Certificate struct {
	Name                  string
	DefaultOcspUrl        *string
	OcspVerifyCertificate *string
	TemplateName          *string
	Misc                  []generic.Xml
}
type UsernameField struct {
	Subject    *string
	SubjectAlt *string
	Misc       []generic.Xml
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
	XMLName                         xml.Name                 `xml:"entry"`
	Name                            string                   `xml:"name,attr"`
	BlockExpiredCertificate         *string                  `xml:"block-expired-cert,omitempty"`
	BlockTimeoutCertificate         *string                  `xml:"block-timeout-cert,omitempty"`
	BlockUnauthenticatedCertificate *string                  `xml:"block-unauthenticated-cert,omitempty"`
	BlockUnknownCertificate         *string                  `xml:"block-unknown-cert,omitempty"`
	Certificate                     *certificateContainerXml `xml:"CA,omitempty"`
	CertificateStatusTimeout        *int64                   `xml:"cert-status-timeout,omitempty"`
	CrlReceiveTimeout               *int64                   `xml:"crl-receive-timeout,omitempty"`
	Domain                          *string                  `xml:"domain,omitempty"`
	OcspExcludeNonce                *string                  `xml:"ocsp-exclude-nonce,omitempty"`
	OcspReceiveTimeout              *int64                   `xml:"ocsp-receive-timeout,omitempty"`
	UseCrl                          *string                  `xml:"use-crl,omitempty"`
	UseOcsp                         *string                  `xml:"use-ocsp,omitempty"`
	UsernameField                   *usernameFieldXml        `xml:"username-field,omitempty"`
	Misc                            []generic.Xml            `xml:",any"`
}
type certificateContainerXml struct {
	Entries []certificateXml `xml:"entry"`
}
type certificateXml struct {
	XMLName               xml.Name      `xml:"entry"`
	Name                  string        `xml:"name,attr"`
	DefaultOcspUrl        *string       `xml:"default-ocsp-url,omitempty"`
	OcspVerifyCertificate *string       `xml:"ocsp-verify-cert,omitempty"`
	TemplateName          *string       `xml:"template-name,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type usernameFieldXml struct {
	Subject    *string       `xml:"subject,omitempty"`
	SubjectAlt *string       `xml:"subject-alt,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockTimeoutCertificate = util.YesNo(s.BlockTimeoutCertificate, nil)
	o.BlockUnauthenticatedCertificate = util.YesNo(s.BlockUnauthenticatedCertificate, nil)
	o.BlockUnknownCertificate = util.YesNo(s.BlockUnknownCertificate, nil)
	if s.Certificate != nil {
		var objs []certificateXml
		for _, elt := range s.Certificate {
			var obj certificateXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Certificate = &certificateContainerXml{Entries: objs}
	}
	o.CertificateStatusTimeout = s.CertificateStatusTimeout
	o.CrlReceiveTimeout = s.CrlReceiveTimeout
	o.Domain = s.Domain
	o.OcspExcludeNonce = util.YesNo(s.OcspExcludeNonce, nil)
	o.OcspReceiveTimeout = s.OcspReceiveTimeout
	o.UseCrl = util.YesNo(s.UseCrl, nil)
	o.UseOcsp = util.YesNo(s.UseOcsp, nil)
	if s.UsernameField != nil {
		var obj usernameFieldXml
		obj.MarshalFromObject(*s.UsernameField)
		o.UsernameField = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var certificateVal []Certificate
	if o.Certificate != nil {
		for _, elt := range o.Certificate.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			certificateVal = append(certificateVal, *obj)
		}
	}
	var usernameFieldVal *UsernameField
	if o.UsernameField != nil {
		obj, err := o.UsernameField.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		usernameFieldVal = obj
	}

	result := &Entry{
		Name:                            o.Name,
		BlockExpiredCertificate:         util.AsBool(o.BlockExpiredCertificate, nil),
		BlockTimeoutCertificate:         util.AsBool(o.BlockTimeoutCertificate, nil),
		BlockUnauthenticatedCertificate: util.AsBool(o.BlockUnauthenticatedCertificate, nil),
		BlockUnknownCertificate:         util.AsBool(o.BlockUnknownCertificate, nil),
		Certificate:                     certificateVal,
		CertificateStatusTimeout:        o.CertificateStatusTimeout,
		CrlReceiveTimeout:               o.CrlReceiveTimeout,
		Domain:                          o.Domain,
		OcspExcludeNonce:                util.AsBool(o.OcspExcludeNonce, nil),
		OcspReceiveTimeout:              o.OcspReceiveTimeout,
		UseCrl:                          util.AsBool(o.UseCrl, nil),
		UseOcsp:                         util.AsBool(o.UseOcsp, nil),
		UsernameField:                   usernameFieldVal,
		Misc:                            o.Misc,
	}
	return result, nil
}
func (o *certificateXml) MarshalFromObject(s Certificate) {
	o.Name = s.Name
	o.DefaultOcspUrl = s.DefaultOcspUrl
	o.OcspVerifyCertificate = s.OcspVerifyCertificate
	o.TemplateName = s.TemplateName
	o.Misc = s.Misc
}

func (o certificateXml) UnmarshalToObject() (*Certificate, error) {

	result := &Certificate{
		Name:                  o.Name,
		DefaultOcspUrl:        o.DefaultOcspUrl,
		OcspVerifyCertificate: o.OcspVerifyCertificate,
		TemplateName:          o.TemplateName,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *usernameFieldXml) MarshalFromObject(s UsernameField) {
	o.Subject = s.Subject
	o.SubjectAlt = s.SubjectAlt
	o.Misc = s.Misc
}

func (o usernameFieldXml) UnmarshalToObject() (*UsernameField, error) {

	result := &UsernameField{
		Subject:    o.Subject,
		SubjectAlt: o.SubjectAlt,
		Misc:       o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "block_expired_certificate" || v == "BlockExpiredCertificate" {
		return e.BlockExpiredCertificate, nil
	}
	if v == "block_timeout_certificate" || v == "BlockTimeoutCertificate" {
		return e.BlockTimeoutCertificate, nil
	}
	if v == "block_unauthenticated_certificate" || v == "BlockUnauthenticatedCertificate" {
		return e.BlockUnauthenticatedCertificate, nil
	}
	if v == "block_unknown_certificate" || v == "BlockUnknownCertificate" {
		return e.BlockUnknownCertificate, nil
	}
	if v == "certificate" || v == "Certificate" {
		return e.Certificate, nil
	}
	if v == "certificate|LENGTH" || v == "Certificate|LENGTH" {
		return int64(len(e.Certificate)), nil
	}
	if v == "certificate_status_timeout" || v == "CertificateStatusTimeout" {
		return e.CertificateStatusTimeout, nil
	}
	if v == "crl_receive_timeout" || v == "CrlReceiveTimeout" {
		return e.CrlReceiveTimeout, nil
	}
	if v == "domain" || v == "Domain" {
		return e.Domain, nil
	}
	if v == "ocsp_exclude_nonce" || v == "OcspExcludeNonce" {
		return e.OcspExcludeNonce, nil
	}
	if v == "ocsp_receive_timeout" || v == "OcspReceiveTimeout" {
		return e.OcspReceiveTimeout, nil
	}
	if v == "use_crl" || v == "UseCrl" {
		return e.UseCrl, nil
	}
	if v == "use_ocsp" || v == "UseOcsp" {
		return e.UseOcsp, nil
	}
	if v == "username_field" || v == "UsernameField" {
		return e.UsernameField, nil
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
	if !util.BoolsMatch(o.BlockExpiredCertificate, other.BlockExpiredCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockTimeoutCertificate, other.BlockTimeoutCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnauthenticatedCertificate, other.BlockUnauthenticatedCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnknownCertificate, other.BlockUnknownCertificate) {
		return false
	}
	if len(o.Certificate) != len(other.Certificate) {
		return false
	}
	for idx := range o.Certificate {
		if !o.Certificate[idx].matches(&other.Certificate[idx]) {
			return false
		}
	}
	if !util.Ints64Match(o.CertificateStatusTimeout, other.CertificateStatusTimeout) {
		return false
	}
	if !util.Ints64Match(o.CrlReceiveTimeout, other.CrlReceiveTimeout) {
		return false
	}
	if !util.StringsMatch(o.Domain, other.Domain) {
		return false
	}
	if !util.BoolsMatch(o.OcspExcludeNonce, other.OcspExcludeNonce) {
		return false
	}
	if !util.Ints64Match(o.OcspReceiveTimeout, other.OcspReceiveTimeout) {
		return false
	}
	if !util.BoolsMatch(o.UseCrl, other.UseCrl) {
		return false
	}
	if !util.BoolsMatch(o.UseOcsp, other.UseOcsp) {
		return false
	}
	if !o.UsernameField.matches(other.UsernameField) {
		return false
	}

	return true
}

func (o *Certificate) matches(other *Certificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.DefaultOcspUrl, other.DefaultOcspUrl) {
		return false
	}
	if !util.StringsMatch(o.OcspVerifyCertificate, other.OcspVerifyCertificate) {
		return false
	}
	if !util.StringsMatch(o.TemplateName, other.TemplateName) {
		return false
	}

	return true
}

func (o *UsernameField) matches(other *UsernameField) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Subject, other.Subject) {
		return false
	}
	if !util.StringsMatch(o.SubjectAlt, other.SubjectAlt) {
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
