package vlan

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
	suffix = []string{"network", "vlan", "$name"}
)

type Entry struct {
	Name             string
	Interface        []string
	VirtualInterface *VirtualInterface
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
}
type VirtualInterface struct {
	Interface      *string
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
	XMLName          xml.Name             `xml:"entry"`
	Name             string               `xml:"name,attr"`
	Interface        *util.MemberType     `xml:"interface,omitempty"`
	VirtualInterface *virtualInterfaceXml `xml:"virtual-interface,omitempty"`
	Misc             []generic.Xml        `xml:",any"`
	MiscAttributes   []xml.Attr           `xml:",any,attr"`
}
type virtualInterfaceXml struct {
	Interface      *string       `xml:"interface,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.VirtualInterface != nil {
		var obj virtualInterfaceXml
		obj.MarshalFromObject(*s.VirtualInterface)
		o.VirtualInterface = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var virtualInterfaceVal *VirtualInterface
	if o.VirtualInterface != nil {
		obj, err := o.VirtualInterface.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		virtualInterfaceVal = obj
	}

	result := &Entry{
		Name:             o.Name,
		Interface:        interfaceVal,
		VirtualInterface: virtualInterfaceVal,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *virtualInterfaceXml) MarshalFromObject(s VirtualInterface) {
	o.Interface = s.Interface
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o virtualInterfaceXml) UnmarshalToObject() (*VirtualInterface, error) {

	result := &VirtualInterface{
		Interface:      o.Interface,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "interface" || v == "Interface" {
		return e.Interface, nil
	}
	if v == "interface|LENGTH" || v == "Interface|LENGTH" {
		return int64(len(e.Interface)), nil
	}
	if v == "virtual_interface" || v == "VirtualInterface" {
		return e.VirtualInterface, nil
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
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !o.VirtualInterface.matches(other.VirtualInterface) {
		return false
	}

	return true
}

func (o *VirtualInterface) matches(other *VirtualInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
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
