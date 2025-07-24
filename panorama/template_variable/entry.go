package template_variable

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
	suffix = []string{"variable", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Type           *Type
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Type struct {
	IpNetmask      *string
	IpRange        *string
	Fqdn           *string
	GroupId        *string
	DevicePriority *string
	DeviceId       *string
	Interface      *string
	AsNumber       *string
	QosProfile     *string
	EgressMax      *string
	LinkTag        *string
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	Type           *typeXml      `xml:"type,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type typeXml struct {
	IpNetmask      *string       `xml:"ip-netmask,omitempty"`
	IpRange        *string       `xml:"ip-range,omitempty"`
	Fqdn           *string       `xml:"fqdn,omitempty"`
	GroupId        *string       `xml:"group-id,omitempty"`
	DevicePriority *string       `xml:"device-priority,omitempty"`
	DeviceId       *string       `xml:"device-id,omitempty"`
	Interface      *string       `xml:"interface,omitempty"`
	AsNumber       *string       `xml:"as-number,omitempty"`
	QosProfile     *string       `xml:"qos-profile,omitempty"`
	EgressMax      *string       `xml:"egress-max,omitempty"`
	LinkTag        *string       `xml:"link-tag,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Type != nil {
		var obj typeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var typeVal *Type
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		Type:           typeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeXml) MarshalFromObject(s Type) {
	o.IpNetmask = s.IpNetmask
	o.IpRange = s.IpRange
	o.Fqdn = s.Fqdn
	o.GroupId = s.GroupId
	o.DevicePriority = s.DevicePriority
	o.DeviceId = s.DeviceId
	o.Interface = s.Interface
	o.AsNumber = s.AsNumber
	o.QosProfile = s.QosProfile
	o.EgressMax = s.EgressMax
	o.LinkTag = s.LinkTag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeXml) UnmarshalToObject() (*Type, error) {

	result := &Type{
		IpNetmask:      o.IpNetmask,
		IpRange:        o.IpRange,
		Fqdn:           o.Fqdn,
		GroupId:        o.GroupId,
		DevicePriority: o.DevicePriority,
		DeviceId:       o.DeviceId,
		Interface:      o.Interface,
		AsNumber:       o.AsNumber,
		QosProfile:     o.QosProfile,
		EgressMax:      o.EgressMax,
		LinkTag:        o.LinkTag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "type" || v == "Type" {
		return e.Type, nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}

	return true
}

func (o *Type) matches(other *Type) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.IpNetmask, other.IpNetmask) {
		return false
	}
	if !util.StringsMatch(o.IpRange, other.IpRange) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !util.StringsMatch(o.GroupId, other.GroupId) {
		return false
	}
	if !util.StringsMatch(o.DevicePriority, other.DevicePriority) {
		return false
	}
	if !util.StringsMatch(o.DeviceId, other.DeviceId) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.AsNumber, other.AsNumber) {
		return false
	}
	if !util.StringsMatch(o.QosProfile, other.QosProfile) {
		return false
	}
	if !util.StringsMatch(o.EgressMax, other.EgressMax) {
		return false
	}
	if !util.StringsMatch(o.LinkTag, other.LinkTag) {
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
