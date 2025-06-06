package devicegroup

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
	suffix = []string{"device-group", "$name"}
)

type Entry struct {
	Name              string
	Description       *string
	Templates         []string
	Devices           []Devices
	AuthorizationCode *string
	Misc              []generic.Xml
}
type Devices struct {
	Name string
	Vsys []string
	Misc []generic.Xml
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
	XMLName           xml.Name             `xml:"entry"`
	Name              string               `xml:"name,attr"`
	Description       *string              `xml:"description,omitempty"`
	Templates         *util.MemberType     `xml:"reference-templates,omitempty"`
	Devices           *devicesContainerXml `xml:"devices,omitempty"`
	AuthorizationCode *string              `xml:"authorization-code,omitempty"`
	Misc              []generic.Xml        `xml:",any"`
}
type devicesContainerXml struct {
	Entries []devicesXml `xml:"entry"`
}
type devicesXml struct {
	XMLName xml.Name         `xml:"entry"`
	Name    string           `xml:"name,attr"`
	Vsys    *util.MemberType `xml:"vsys,omitempty"`
	Misc    []generic.Xml    `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Templates != nil {
		o.Templates = util.StrToMem(s.Templates)
	}
	if s.Devices != nil {
		var objs []devicesXml
		for _, elt := range s.Devices {
			var obj devicesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Devices = &devicesContainerXml{Entries: objs}
	}
	o.AuthorizationCode = s.AuthorizationCode
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var templatesVal []string
	if o.Templates != nil {
		templatesVal = util.MemToStr(o.Templates)
	}
	var devicesVal []Devices
	if o.Devices != nil {
		for _, elt := range o.Devices.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			devicesVal = append(devicesVal, *obj)
		}
	}

	result := &Entry{
		Name:              o.Name,
		Description:       o.Description,
		Templates:         templatesVal,
		Devices:           devicesVal,
		AuthorizationCode: o.AuthorizationCode,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *devicesXml) MarshalFromObject(s Devices) {
	o.Name = s.Name
	if s.Vsys != nil {
		o.Vsys = util.StrToMem(s.Vsys)
	}
	o.Misc = s.Misc
}

func (o devicesXml) UnmarshalToObject() (*Devices, error) {
	var vsysVal []string
	if o.Vsys != nil {
		vsysVal = util.MemToStr(o.Vsys)
	}

	result := &Devices{
		Name: o.Name,
		Vsys: vsysVal,
		Misc: o.Misc,
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
	if v == "authorization_code" || v == "AuthorizationCode" {
		return e.AuthorizationCode, nil
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
	if len(o.Devices) != len(other.Devices) {
		return false
	}
	for idx := range o.Devices {
		if !o.Devices[idx].matches(&other.Devices[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.AuthorizationCode, other.AuthorizationCode) {
		return false
	}

	return true
}

func (o *Devices) matches(other *Devices) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vsys, other.Vsys) {
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
