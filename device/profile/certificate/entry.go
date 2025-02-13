package certificate

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
	Suffix = []string{}
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

	Misc map[string][]generic.Xml
}

type Certificate struct {
	DefaultOcspUrl        *string
	Name                  string
	OcspVerifyCertificate *string
	TemplateName          *string
}
type UsernameField struct {
	Subject    *string
	SubjectAlt *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                         xml.Name          `xml:"entry"`
	Name                            string            `xml:"name,attr"`
	BlockExpiredCertificate         *string           `xml:"block-expired-cert,omitempty"`
	BlockTimeoutCertificate         *string           `xml:"block-timeout-cert,omitempty"`
	BlockUnauthenticatedCertificate *string           `xml:"block-unauthenticated-cert,omitempty"`
	BlockUnknownCertificate         *string           `xml:"block-unknown-cert,omitempty"`
	Certificate                     []CertificateXml  `xml:"CA>entry,omitempty"`
	CertificateStatusTimeout        *int64            `xml:"cert-status-timeout,omitempty"`
	CrlReceiveTimeout               *int64            `xml:"crl-receive-timeout,omitempty"`
	Domain                          *string           `xml:"domain,omitempty"`
	OcspExcludeNonce                *string           `xml:"ocsp-exclude-nonce,omitempty"`
	OcspReceiveTimeout              *int64            `xml:"ocsp-receive-timeout,omitempty"`
	UseCrl                          *string           `xml:"use-crl,omitempty"`
	UseOcsp                         *string           `xml:"use-ocsp,omitempty"`
	UsernameField                   *UsernameFieldXml `xml:"username-field,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type CertificateXml struct {
	DefaultOcspUrl        *string  `xml:"default-ocsp-url,omitempty"`
	XMLName               xml.Name `xml:"entry"`
	Name                  string   `xml:"name,attr"`
	OcspVerifyCertificate *string  `xml:"ocsp-verify-cert,omitempty"`
	TemplateName          *string  `xml:"template-name,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UsernameFieldXml struct {
	Subject    *string `xml:"subject,omitempty"`
	SubjectAlt *string `xml:"subject-alt,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.BlockExpiredCertificate = util.YesNo(o.BlockExpiredCertificate, nil)
	entry.BlockTimeoutCertificate = util.YesNo(o.BlockTimeoutCertificate, nil)
	entry.BlockUnauthenticatedCertificate = util.YesNo(o.BlockUnauthenticatedCertificate, nil)
	entry.BlockUnknownCertificate = util.YesNo(o.BlockUnknownCertificate, nil)
	var nestedCertificateCol []CertificateXml
	if o.Certificate != nil {
		nestedCertificateCol = []CertificateXml{}
		for _, oCertificate := range o.Certificate {
			nestedCertificate := CertificateXml{}
			if _, ok := o.Misc["Certificate"]; ok {
				nestedCertificate.Misc = o.Misc["Certificate"]
			}
			if oCertificate.Name != "" {
				nestedCertificate.Name = oCertificate.Name
			}
			if oCertificate.DefaultOcspUrl != nil {
				nestedCertificate.DefaultOcspUrl = oCertificate.DefaultOcspUrl
			}
			if oCertificate.OcspVerifyCertificate != nil {
				nestedCertificate.OcspVerifyCertificate = oCertificate.OcspVerifyCertificate
			}
			if oCertificate.TemplateName != nil {
				nestedCertificate.TemplateName = oCertificate.TemplateName
			}
			nestedCertificateCol = append(nestedCertificateCol, nestedCertificate)
		}
		entry.Certificate = nestedCertificateCol
	}

	entry.CertificateStatusTimeout = o.CertificateStatusTimeout
	entry.CrlReceiveTimeout = o.CrlReceiveTimeout
	entry.Domain = o.Domain
	entry.OcspExcludeNonce = util.YesNo(o.OcspExcludeNonce, nil)
	entry.OcspReceiveTimeout = o.OcspReceiveTimeout
	entry.UseCrl = util.YesNo(o.UseCrl, nil)
	entry.UseOcsp = util.YesNo(o.UseOcsp, nil)
	var nestedUsernameField *UsernameFieldXml
	if o.UsernameField != nil {
		nestedUsernameField = &UsernameFieldXml{}
		if _, ok := o.Misc["UsernameField"]; ok {
			nestedUsernameField.Misc = o.Misc["UsernameField"]
		}
		if o.UsernameField.SubjectAlt != nil {
			nestedUsernameField.SubjectAlt = o.UsernameField.SubjectAlt
		}
		if o.UsernameField.Subject != nil {
			nestedUsernameField.Subject = o.UsernameField.Subject
		}
	}
	entry.UsernameField = nestedUsernameField

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
		entry.BlockExpiredCertificate = util.AsBool(o.BlockExpiredCertificate, nil)
		entry.BlockTimeoutCertificate = util.AsBool(o.BlockTimeoutCertificate, nil)
		entry.BlockUnauthenticatedCertificate = util.AsBool(o.BlockUnauthenticatedCertificate, nil)
		entry.BlockUnknownCertificate = util.AsBool(o.BlockUnknownCertificate, nil)
		var nestedCertificateCol []Certificate
		if o.Certificate != nil {
			nestedCertificateCol = []Certificate{}
			for _, oCertificate := range o.Certificate {
				nestedCertificate := Certificate{}
				if oCertificate.Misc != nil {
					entry.Misc["Certificate"] = oCertificate.Misc
				}
				if oCertificate.DefaultOcspUrl != nil {
					nestedCertificate.DefaultOcspUrl = oCertificate.DefaultOcspUrl
				}
				if oCertificate.OcspVerifyCertificate != nil {
					nestedCertificate.OcspVerifyCertificate = oCertificate.OcspVerifyCertificate
				}
				if oCertificate.TemplateName != nil {
					nestedCertificate.TemplateName = oCertificate.TemplateName
				}
				if oCertificate.Name != "" {
					nestedCertificate.Name = oCertificate.Name
				}
				nestedCertificateCol = append(nestedCertificateCol, nestedCertificate)
			}
			entry.Certificate = nestedCertificateCol
		}

		entry.CertificateStatusTimeout = o.CertificateStatusTimeout
		entry.CrlReceiveTimeout = o.CrlReceiveTimeout
		entry.Domain = o.Domain
		entry.OcspExcludeNonce = util.AsBool(o.OcspExcludeNonce, nil)
		entry.OcspReceiveTimeout = o.OcspReceiveTimeout
		entry.UseCrl = util.AsBool(o.UseCrl, nil)
		entry.UseOcsp = util.AsBool(o.UseOcsp, nil)
		var nestedUsernameField *UsernameField
		if o.UsernameField != nil {
			nestedUsernameField = &UsernameField{}
			if o.UsernameField.Misc != nil {
				entry.Misc["UsernameField"] = o.UsernameField.Misc
			}
			if o.UsernameField.Subject != nil {
				nestedUsernameField.Subject = o.UsernameField.Subject
			}
			if o.UsernameField.SubjectAlt != nil {
				nestedUsernameField.SubjectAlt = o.UsernameField.SubjectAlt
			}
		}
		entry.UsernameField = nestedUsernameField

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
	if !util.BoolsMatch(a.BlockExpiredCertificate, b.BlockExpiredCertificate) {
		return false
	}
	if !util.BoolsMatch(a.BlockTimeoutCertificate, b.BlockTimeoutCertificate) {
		return false
	}
	if !util.BoolsMatch(a.BlockUnauthenticatedCertificate, b.BlockUnauthenticatedCertificate) {
		return false
	}
	if !util.BoolsMatch(a.BlockUnknownCertificate, b.BlockUnknownCertificate) {
		return false
	}
	if !matchCertificate(a.Certificate, b.Certificate) {
		return false
	}
	if !util.Ints64Match(a.CertificateStatusTimeout, b.CertificateStatusTimeout) {
		return false
	}
	if !util.Ints64Match(a.CrlReceiveTimeout, b.CrlReceiveTimeout) {
		return false
	}
	if !util.StringsMatch(a.Domain, b.Domain) {
		return false
	}
	if !util.BoolsMatch(a.OcspExcludeNonce, b.OcspExcludeNonce) {
		return false
	}
	if !util.Ints64Match(a.OcspReceiveTimeout, b.OcspReceiveTimeout) {
		return false
	}
	if !util.BoolsMatch(a.UseCrl, b.UseCrl) {
		return false
	}
	if !util.BoolsMatch(a.UseOcsp, b.UseOcsp) {
		return false
	}
	if !matchUsernameField(a.UsernameField, b.UsernameField) {
		return false
	}

	return true
}

func matchCertificate(a []Certificate, b []Certificate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.TemplateName, b.TemplateName) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.DefaultOcspUrl, b.DefaultOcspUrl) {
				return false
			}
			if !util.StringsMatch(a.OcspVerifyCertificate, b.OcspVerifyCertificate) {
				return false
			}
		}
	}
	return true
}
func matchUsernameField(a *UsernameField, b *UsernameField) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Subject, b.Subject) {
		return false
	}
	if !util.StringsMatch(a.SubjectAlt, b.SubjectAlt) {
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
