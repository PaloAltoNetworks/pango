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
	Suffix = []string{"service-group"}
)

type Entry struct {
	Name        string
	Description *string
	Members     []string
	Tags        []string

	Misc map[string][]generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	Description *string          `xml:"description,omitempty"`
	Members     *util.MemberType `xml:"members,omitempty"`
	Tags        *util.MemberType `xml:"tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "members" || v == "Members" {
		return e.Members, nil
	}
	if v == "members|LENGTH" || v == "Members|LENGTH" {
		return int64(len(e.Members)), nil
	}
	if v == "tags" || v == "Tags" {
		return e.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return int64(len(e.Tags)), nil
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
	entry.Members = util.StrToMem(o.Members)
	entry.Tags = util.StrToMem(o.Tags)

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
		entry.Members = util.MemToStr(o.Members)
		entry.Tags = util.MemToStr(o.Tags)

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
	if !util.OrderedListsMatch(a.Members, b.Members) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
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
