package wildfireanalysis

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
	suffix = []string{"profiles", "wildfire-analysis", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Rules           []Rules
	Misc            []generic.Xml
}
type Rules struct {
	Name        string
	Application []string
	FileType    []string
	Direction   *string
	Analysis    *string
	Misc        []generic.Xml
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
	XMLName         xml.Name           `xml:"entry"`
	Name            string             `xml:"name,attr"`
	Description     *string            `xml:"description,omitempty"`
	DisableOverride *string            `xml:"disable-override,omitempty"`
	Rules           *rulesContainerXml `xml:"rules,omitempty"`
	Misc            []generic.Xml      `xml:",any"`
}
type rulesContainerXml struct {
	Entries []rulesXml `xml:"entry"`
}
type rulesXml struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	Application *util.MemberType `xml:"application,omitempty"`
	FileType    *util.MemberType `xml:"file-type,omitempty"`
	Direction   *string          `xml:"direction,omitempty"`
	Analysis    *string          `xml:"analysis,omitempty"`
	Misc        []generic.Xml    `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.Rules != nil {
		var objs []rulesXml
		for _, elt := range s.Rules {
			var obj rulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Rules = &rulesContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var rulesVal []Rules
	if o.Rules != nil {
		for _, elt := range o.Rules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rulesVal = append(rulesVal, *obj)
		}
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		DisableOverride: o.DisableOverride,
		Rules:           rulesVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *rulesXml) MarshalFromObject(s Rules) {
	o.Name = s.Name
	if s.Application != nil {
		o.Application = util.StrToMem(s.Application)
	}
	if s.FileType != nil {
		o.FileType = util.StrToMem(s.FileType)
	}
	o.Direction = s.Direction
	o.Analysis = s.Analysis
	o.Misc = s.Misc
}

func (o rulesXml) UnmarshalToObject() (*Rules, error) {
	var applicationVal []string
	if o.Application != nil {
		applicationVal = util.MemToStr(o.Application)
	}
	var fileTypeVal []string
	if o.FileType != nil {
		fileTypeVal = util.MemToStr(o.FileType)
	}

	result := &Rules{
		Name:        o.Name,
		Application: applicationVal,
		FileType:    fileTypeVal,
		Direction:   o.Direction,
		Analysis:    o.Analysis,
		Misc:        o.Misc,
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
	if len(o.Rules) != len(other.Rules) {
		return false
	}
	for idx := range o.Rules {
		if !o.Rules[idx].matches(&other.Rules[idx]) {
			return false
		}
	}

	return true
}

func (o *Rules) matches(other *Rules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Application, other.Application) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FileType, other.FileType) {
		return false
	}
	if !util.StringsMatch(o.Direction, other.Direction) {
		return false
	}
	if !util.StringsMatch(o.Analysis, other.Analysis) {
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
