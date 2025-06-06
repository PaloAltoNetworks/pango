package group

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
	suffix = []string{"address-group", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Tag             []string
	Dynamic         *Dynamic
	Static          []string
	Misc            []generic.Xml
}
type Dynamic struct {
	Filter *string
	Misc   []generic.Xml
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
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	Description     *string          `xml:"description,omitempty"`
	DisableOverride *string          `xml:"disable-override,omitempty"`
	Tag             *util.MemberType `xml:"tag,omitempty"`
	Dynamic         *dynamicXml      `xml:"dynamic,omitempty"`
	Static          *util.MemberType `xml:"static,omitempty"`
	Misc            []generic.Xml    `xml:",any"`
}
type dynamicXml struct {
	Filter *string       `xml:"filter,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	if s.Dynamic != nil {
		var obj dynamicXml
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.Static != nil {
		o.Static = util.StrToMem(s.Static)
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}
	var dynamicVal *Dynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var staticVal []string
	if o.Static != nil {
		staticVal = util.MemToStr(o.Static)
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		DisableOverride: o.DisableOverride,
		Tag:             tagVal,
		Dynamic:         dynamicVal,
		Static:          staticVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *dynamicXml) MarshalFromObject(s Dynamic) {
	o.Filter = s.Filter
	o.Misc = s.Misc
}

func (o dynamicXml) UnmarshalToObject() (*Dynamic, error) {

	result := &Dynamic{
		Filter: o.Filter,
		Misc:   o.Misc,
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
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
	}
	if v == "dynamic" || v == "Dynamic" {
		return e.Dynamic, nil
	}
	if v == "static" || v == "Static" {
		return e.Static, nil
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
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}
	if !o.Dynamic.matches(other.Dynamic) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Static, other.Static) {
		return false
	}

	return true
}

func (o *Dynamic) matches(other *Dynamic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Filter, other.Filter) {
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
