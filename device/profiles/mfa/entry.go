package mfa

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
	suffix = []string{"server-profile", "mfa-server-profile", "$name"}
)

type Entry struct {
	Name           string
	MfaCertProfile *string
	MfaConfig      []MfaConfig
	MfaVendorType  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MfaConfig struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	XMLName        xml.Name               `xml:"entry"`
	Name           string                 `xml:"name,attr"`
	MfaCertProfile *string                `xml:"mfa-cert-profile,omitempty"`
	MfaConfig      *mfaConfigContainerXml `xml:"mfa-config,omitempty"`
	MfaVendorType  *string                `xml:"mfa-vendor-type,omitempty"`
	Misc           []generic.Xml          `xml:",any"`
	MiscAttributes []xml.Attr             `xml:",any,attr"`
}
type mfaConfigContainerXml struct {
	Entries []mfaConfigXml `xml:"entry"`
}
type mfaConfigXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.MfaCertProfile = s.MfaCertProfile
	if s.MfaConfig != nil {
		var objs []mfaConfigXml
		for _, elt := range s.MfaConfig {
			var obj mfaConfigXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MfaConfig = &mfaConfigContainerXml{Entries: objs}
	}
	o.MfaVendorType = s.MfaVendorType
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var mfaConfigVal []MfaConfig
	if o.MfaConfig != nil {
		for _, elt := range o.MfaConfig.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			mfaConfigVal = append(mfaConfigVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		MfaCertProfile: o.MfaCertProfile,
		MfaConfig:      mfaConfigVal,
		MfaVendorType:  o.MfaVendorType,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *mfaConfigXml) MarshalFromObject(s MfaConfig) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o mfaConfigXml) UnmarshalToObject() (*MfaConfig, error) {

	result := &MfaConfig{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "mfa_cert_profile" || v == "MfaCertProfile" {
		return e.MfaCertProfile, nil
	}
	if v == "mfa_config" || v == "MfaConfig" {
		return e.MfaConfig, nil
	}
	if v == "mfa_config|LENGTH" || v == "MfaConfig|LENGTH" {
		return int64(len(e.MfaConfig)), nil
	}
	if v == "mfa_vendor_type" || v == "MfaVendorType" {
		return e.MfaVendorType, nil
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
	if !util.StringsMatch(o.MfaCertProfile, other.MfaCertProfile) {
		return false
	}
	if len(o.MfaConfig) != len(other.MfaConfig) {
		return false
	}
	for idx := range o.MfaConfig {
		if !o.MfaConfig[idx].matches(&other.MfaConfig[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.MfaVendorType, other.MfaVendorType) {
		return false
	}

	return true
}

func (o *MfaConfig) matches(other *MfaConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
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
