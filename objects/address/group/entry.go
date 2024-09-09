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
	Suffix = []string{"address-group"}
)

type Entry struct {
	Name        string
	Description *string
	Tags        []string
	Dynamic     *string
	Static      []string

	Misc map[string][]generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	Description *string          `xml:"description,omitempty"`
	Tags        *util.MemberType `xml:"tag,omitempty"`
	Dynamic     *string          `xml:"dynamic>filter,omitempty"`
	Static      *util.MemberType `xml:"static,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "tags" || v == "Tags" {
		return e.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return int64(len(e.Tags)), nil
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

func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}

	entry.Name = o.Name
	entry.Description = o.Description
	entry.Tags = util.StrToMem(o.Tags)
	entry.Dynamic = o.Dynamic
	entry.Static = util.StrToMem(o.Static)

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
		entry.Description = o.Description
		entry.Tags = util.MemToStr(o.Tags)
		entry.Dynamic = o.Dynamic
		entry.Static = util.MemToStr(o.Static)

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
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.Dynamic, b.Dynamic) {
		return false
	}
	if !util.OrderedListsMatch(a.Static, b.Static) {
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
