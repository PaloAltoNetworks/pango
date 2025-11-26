package aspathacl

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
	suffix = []string{"network", "routing-profile", "filters", "as-path-access-list", "$name"}
)

type Entry struct {
	Name           string
	AspathEntry    []AspathEntry
	Description    *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AspathEntry struct {
	Name           string
	Action         *string
	AspathRegex    *string
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
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	AspathEntry    *aspathEntryContainerXml `xml:"aspath-entry,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type aspathEntryContainerXml struct {
	Entries []aspathEntryXml `xml:"entry"`
}
type aspathEntryXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Action         *string       `xml:"action,omitempty"`
	AspathRegex    *string       `xml:"aspath-regex,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AspathEntry != nil {
		var objs []aspathEntryXml
		for _, elt := range s.AspathEntry {
			var obj aspathEntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AspathEntry = &aspathEntryContainerXml{Entries: objs}
	}
	o.Description = s.Description
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var aspathEntryVal []AspathEntry
	if o.AspathEntry != nil {
		for _, elt := range o.AspathEntry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			aspathEntryVal = append(aspathEntryVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		AspathEntry:    aspathEntryVal,
		Description:    o.Description,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *aspathEntryXml) MarshalFromObject(s AspathEntry) {
	o.Name = s.Name
	o.Action = s.Action
	o.AspathRegex = s.AspathRegex
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o aspathEntryXml) UnmarshalToObject() (*AspathEntry, error) {

	result := &AspathEntry{
		Name:           o.Name,
		Action:         o.Action,
		AspathRegex:    o.AspathRegex,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "aspath_entry" || v == "AspathEntry" {
		return e.AspathEntry, nil
	}
	if v == "aspath_entry|LENGTH" || v == "AspathEntry|LENGTH" {
		return int64(len(e.AspathEntry)), nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
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
	if len(o.AspathEntry) != len(other.AspathEntry) {
		return false
	}
	for idx := range o.AspathEntry {
		if !o.AspathEntry[idx].matches(&other.AspathEntry[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}

	return true
}

func (o *AspathEntry) matches(other *AspathEntry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.AspathRegex, other.AspathRegex) {
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
