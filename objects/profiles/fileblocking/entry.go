package fileblocking

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
	Suffix = []string{}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Rules           []Rules

	Misc map[string][]generic.Xml
}

type Rules struct {
	Action      *string
	Application []string
	Direction   *string
	FileType    []string
	Name        string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName         xml.Name   `xml:"entry"`
	Name            string     `xml:"name,attr"`
	Description     *string    `xml:"description,omitempty"`
	DisableOverride *string    `xml:"disable-override,omitempty"`
	Rules           []RulesXml `xml:"rules>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesXml struct {
	Action      *string          `xml:"action,omitempty"`
	Application *util.MemberType `xml:"application,omitempty"`
	Direction   *string          `xml:"direction,omitempty"`
	FileType    *util.MemberType `xml:"file-type,omitempty"`
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
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
	if v == "rules" || v == "Rules" {
		return e.Rules, nil
	}
	if v == "rules|LENGTH" || v == "Rules|LENGTH" {
		return int64(len(e.Rules)), nil
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
	entry.DisableOverride = o.DisableOverride
	var nestedRulesCol []RulesXml
	if o.Rules != nil {
		nestedRulesCol = []RulesXml{}
		for _, oRules := range o.Rules {
			nestedRules := RulesXml{}
			if _, ok := o.Misc["Rules"]; ok {
				nestedRules.Misc = o.Misc["Rules"]
			}
			if oRules.Application != nil {
				nestedRules.Application = util.StrToMem(oRules.Application)
			}
			if oRules.FileType != nil {
				nestedRules.FileType = util.StrToMem(oRules.FileType)
			}
			if oRules.Direction != nil {
				nestedRules.Direction = oRules.Direction
			}
			if oRules.Action != nil {
				nestedRules.Action = oRules.Action
			}
			if oRules.Name != "" {
				nestedRules.Name = oRules.Name
			}
			nestedRulesCol = append(nestedRulesCol, nestedRules)
		}
		entry.Rules = nestedRulesCol
	}

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
		entry.DisableOverride = o.DisableOverride
		var nestedRulesCol []Rules
		if o.Rules != nil {
			nestedRulesCol = []Rules{}
			for _, oRules := range o.Rules {
				nestedRules := Rules{}
				if oRules.Misc != nil {
					entry.Misc["Rules"] = oRules.Misc
				}
				if oRules.Action != nil {
					nestedRules.Action = oRules.Action
				}
				if oRules.Name != "" {
					nestedRules.Name = oRules.Name
				}
				if oRules.Application != nil {
					nestedRules.Application = util.MemToStr(oRules.Application)
				}
				if oRules.FileType != nil {
					nestedRules.FileType = util.MemToStr(oRules.FileType)
				}
				if oRules.Direction != nil {
					nestedRules.Direction = oRules.Direction
				}
				nestedRulesCol = append(nestedRulesCol, nestedRules)
			}
			entry.Rules = nestedRulesCol
		}

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
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !matchRules(a.Rules, b.Rules) {
		return false
	}

	return true
}

func matchRules(a []Rules, b []Rules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.OrderedListsMatch(a.Application, b.Application) {
				return false
			}
			if !util.OrderedListsMatch(a.FileType, b.FileType) {
				return false
			}
			if !util.StringsMatch(a.Direction, b.Direction) {
				return false
			}
			if !util.StringsMatch(a.Action, b.Action) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
