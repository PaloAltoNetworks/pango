package virtualwire

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
	suffix = []string{"network", "virtual-wire", "$name"}
)

type Entry struct {
	Name                 string
	Interface1           *string
	Interface2           *string
	LinkStatePassThrough *LinkStatePassThrough
	MulticastFirewalling *MulticastFirewalling
	TagAllowed           *string
	Misc                 []generic.Xml
	MiscAttributes       []xml.Attr
}
type LinkStatePassThrough struct {
	Enable         *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MulticastFirewalling struct {
	Enable         *bool
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
	XMLName              xml.Name                 `xml:"entry"`
	Name                 string                   `xml:"name,attr"`
	Interface1           *string                  `xml:"interface1,omitempty"`
	Interface2           *string                  `xml:"interface2,omitempty"`
	LinkStatePassThrough *linkStatePassThroughXml `xml:"link-state-pass-through,omitempty"`
	MulticastFirewalling *multicastFirewallingXml `xml:"multicast-firewalling,omitempty"`
	TagAllowed           *string                  `xml:"tag-allowed,omitempty"`
	Misc                 []generic.Xml            `xml:",any"`
	MiscAttributes       []xml.Attr               `xml:",any,attr"`
}
type linkStatePassThroughXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type multicastFirewallingXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Interface1 = s.Interface1
	o.Interface2 = s.Interface2
	if s.LinkStatePassThrough != nil {
		var obj linkStatePassThroughXml
		obj.MarshalFromObject(*s.LinkStatePassThrough)
		o.LinkStatePassThrough = &obj
	}
	if s.MulticastFirewalling != nil {
		var obj multicastFirewallingXml
		obj.MarshalFromObject(*s.MulticastFirewalling)
		o.MulticastFirewalling = &obj
	}
	o.TagAllowed = s.TagAllowed
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var linkStatePassThroughVal *LinkStatePassThrough
	if o.LinkStatePassThrough != nil {
		obj, err := o.LinkStatePassThrough.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkStatePassThroughVal = obj
	}
	var multicastFirewallingVal *MulticastFirewalling
	if o.MulticastFirewalling != nil {
		obj, err := o.MulticastFirewalling.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastFirewallingVal = obj
	}

	result := &Entry{
		Name:                 o.Name,
		Interface1:           o.Interface1,
		Interface2:           o.Interface2,
		LinkStatePassThrough: linkStatePassThroughVal,
		MulticastFirewalling: multicastFirewallingVal,
		TagAllowed:           o.TagAllowed,
		Misc:                 o.Misc,
		MiscAttributes:       o.MiscAttributes,
	}
	return result, nil
}
func (o *linkStatePassThroughXml) MarshalFromObject(s LinkStatePassThrough) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o linkStatePassThroughXml) UnmarshalToObject() (*LinkStatePassThrough, error) {

	result := &LinkStatePassThrough{
		Enable:         util.AsBool(o.Enable, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *multicastFirewallingXml) MarshalFromObject(s MulticastFirewalling) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o multicastFirewallingXml) UnmarshalToObject() (*MulticastFirewalling, error) {

	result := &MulticastFirewalling{
		Enable:         util.AsBool(o.Enable, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "interface1" || v == "Interface1" {
		return e.Interface1, nil
	}
	if v == "interface2" || v == "Interface2" {
		return e.Interface2, nil
	}
	if v == "link_state_pass_through" || v == "LinkStatePassThrough" {
		return e.LinkStatePassThrough, nil
	}
	if v == "multicast_firewalling" || v == "MulticastFirewalling" {
		return e.MulticastFirewalling, nil
	}
	if v == "tag_allowed" || v == "TagAllowed" {
		return e.TagAllowed, nil
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
	if !util.StringsMatch(o.Interface1, other.Interface1) {
		return false
	}
	if !util.StringsMatch(o.Interface2, other.Interface2) {
		return false
	}
	if !o.LinkStatePassThrough.matches(other.LinkStatePassThrough) {
		return false
	}
	if !o.MulticastFirewalling.matches(other.MulticastFirewalling) {
		return false
	}
	if !util.StringsMatch(o.TagAllowed, other.TagAllowed) {
		return false
	}

	return true
}

func (o *LinkStatePassThrough) matches(other *LinkStatePassThrough) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}

	return true
}

func (o *MulticastFirewalling) matches(other *MulticastFirewalling) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
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
