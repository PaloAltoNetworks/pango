package template_stack

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
	suffix = []string{"template-stack", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	Templates       []string
	Devices         []string
	DefaultVsys     *string
	UserGroupSource *UserGroupSource
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type UserGroupSource struct {
	MasterDevice   *string
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
	XMLName         xml.Name            `xml:"entry"`
	Name            string              `xml:"name,attr"`
	Description     *string             `xml:"description,omitempty"`
	Templates       *util.MemberType    `xml:"templates,omitempty"`
	Devices         *util.MemberType    `xml:"devices,omitempty"`
	DefaultVsys     *string             `xml:"settings>default-vsys,omitempty"`
	UserGroupSource *userGroupSourceXml `xml:"user-group-source,omitempty"`
	Misc            []generic.Xml       `xml:",any"`
	MiscAttributes  []xml.Attr          `xml:",any,attr"`
}
type userGroupSourceXml struct {
	MasterDevice   *string       `xml:"master-device,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Templates != nil {
		o.Templates = util.StrToMem(s.Templates)
	}
	if s.Devices != nil {
		o.Devices = util.StrToMem(s.Devices)
	}
	o.DefaultVsys = s.DefaultVsys
	if s.UserGroupSource != nil {
		var obj userGroupSourceXml
		obj.MarshalFromObject(*s.UserGroupSource)
		o.UserGroupSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var templatesVal []string
	if o.Templates != nil {
		templatesVal = util.MemToStr(o.Templates)
	}
	var devicesVal []string
	if o.Devices != nil {
		devicesVal = util.MemToStr(o.Devices)
	}
	var userGroupSourceVal *UserGroupSource
	if o.UserGroupSource != nil {
		obj, err := o.UserGroupSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		userGroupSourceVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		Templates:       templatesVal,
		Devices:         devicesVal,
		DefaultVsys:     o.DefaultVsys,
		UserGroupSource: userGroupSourceVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *userGroupSourceXml) MarshalFromObject(s UserGroupSource) {
	o.MasterDevice = s.MasterDevice
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o userGroupSourceXml) UnmarshalToObject() (*UserGroupSource, error) {

	result := &UserGroupSource{
		MasterDevice:   o.MasterDevice,
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
	if v == "templates" || v == "Templates" {
		return e.Templates, nil
	}
	if v == "templates|LENGTH" || v == "Templates|LENGTH" {
		return int64(len(e.Templates)), nil
	}
	if v == "devices" || v == "Devices" {
		return e.Devices, nil
	}
	if v == "devices|LENGTH" || v == "Devices|LENGTH" {
		return int64(len(e.Devices)), nil
	}
	if v == "default_vsys" || v == "DefaultVsys" {
		return e.DefaultVsys, nil
	}
	if v == "user_group_source" || v == "UserGroupSource" {
		return e.UserGroupSource, nil
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
	if !util.OrderedListsMatch[string](o.Templates, other.Templates) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Devices, other.Devices) {
		return false
	}
	if !util.StringsMatch(o.DefaultVsys, other.DefaultVsys) {
		return false
	}
	if !o.UserGroupSource.matches(other.UserGroupSource) {
		return false
	}

	return true
}

func (o *UserGroupSource) matches(other *UserGroupSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.MasterDevice, other.MasterDevice) {
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
