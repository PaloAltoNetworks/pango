package communitylist

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
	suffix = []string{"network", "routing-profile", "filters", "community-list", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Type           *Type
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Type struct {
	Extended       *TypeExtended
	Large          *TypeLarge
	Regular        *TypeRegular
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeExtended struct {
	ExtendedEntry  []TypeExtendedExtendedEntry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeExtendedExtendedEntry struct {
	Name           string
	Action         *string
	EcRegex        []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeLarge struct {
	LargeEntry     []TypeLargeLargeEntry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeLargeLargeEntry struct {
	Name           string
	Action         *string
	LcRegex        []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeRegular struct {
	RegularEntry   []TypeRegularRegularEntry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeRegularRegularEntry struct {
	Name           string
	Action         *string
	Community      []string
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	Type           *typeXml      `xml:"type,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type typeXml struct {
	Extended       *typeExtendedXml `xml:"extended,omitempty"`
	Large          *typeLargeXml    `xml:"large,omitempty"`
	Regular        *typeRegularXml  `xml:"regular,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type typeExtendedXml struct {
	ExtendedEntry  *typeExtendedExtendedEntryContainerXml `xml:"extended-entry,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type typeExtendedExtendedEntryContainerXml struct {
	Entries []typeExtendedExtendedEntryXml `xml:"entry"`
}
type typeExtendedExtendedEntryXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Action         *string          `xml:"action,omitempty"`
	EcRegex        *util.MemberType `xml:"ec-regex,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type typeLargeXml struct {
	LargeEntry     *typeLargeLargeEntryContainerXml `xml:"large-entry,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type typeLargeLargeEntryContainerXml struct {
	Entries []typeLargeLargeEntryXml `xml:"entry"`
}
type typeLargeLargeEntryXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Action         *string          `xml:"action,omitempty"`
	LcRegex        *util.MemberType `xml:"lc-regex,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type typeRegularXml struct {
	RegularEntry   *typeRegularRegularEntryContainerXml `xml:"regular-entry,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type typeRegularRegularEntryContainerXml struct {
	Entries []typeRegularRegularEntryXml `xml:"entry"`
}
type typeRegularRegularEntryXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Action         *string          `xml:"action,omitempty"`
	Community      *util.MemberType `xml:"community,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Type != nil {
		var obj typeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var typeVal *Type
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		Type:           typeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeXml) MarshalFromObject(s Type) {
	if s.Extended != nil {
		var obj typeExtendedXml
		obj.MarshalFromObject(*s.Extended)
		o.Extended = &obj
	}
	if s.Large != nil {
		var obj typeLargeXml
		obj.MarshalFromObject(*s.Large)
		o.Large = &obj
	}
	if s.Regular != nil {
		var obj typeRegularXml
		obj.MarshalFromObject(*s.Regular)
		o.Regular = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeXml) UnmarshalToObject() (*Type, error) {
	var extendedVal *TypeExtended
	if o.Extended != nil {
		obj, err := o.Extended.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedVal = obj
	}
	var largeVal *TypeLarge
	if o.Large != nil {
		obj, err := o.Large.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		largeVal = obj
	}
	var regularVal *TypeRegular
	if o.Regular != nil {
		obj, err := o.Regular.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		regularVal = obj
	}

	result := &Type{
		Extended:       extendedVal,
		Large:          largeVal,
		Regular:        regularVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeExtendedXml) MarshalFromObject(s TypeExtended) {
	if s.ExtendedEntry != nil {
		var objs []typeExtendedExtendedEntryXml
		for _, elt := range s.ExtendedEntry {
			var obj typeExtendedExtendedEntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExtendedEntry = &typeExtendedExtendedEntryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeExtendedXml) UnmarshalToObject() (*TypeExtended, error) {
	var extendedEntryVal []TypeExtendedExtendedEntry
	if o.ExtendedEntry != nil {
		for _, elt := range o.ExtendedEntry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			extendedEntryVal = append(extendedEntryVal, *obj)
		}
	}

	result := &TypeExtended{
		ExtendedEntry:  extendedEntryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeExtendedExtendedEntryXml) MarshalFromObject(s TypeExtendedExtendedEntry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.EcRegex != nil {
		o.EcRegex = util.StrToMem(s.EcRegex)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeExtendedExtendedEntryXml) UnmarshalToObject() (*TypeExtendedExtendedEntry, error) {
	var ecRegexVal []string
	if o.EcRegex != nil {
		ecRegexVal = util.MemToStr(o.EcRegex)
	}

	result := &TypeExtendedExtendedEntry{
		Name:           o.Name,
		Action:         o.Action,
		EcRegex:        ecRegexVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeLargeXml) MarshalFromObject(s TypeLarge) {
	if s.LargeEntry != nil {
		var objs []typeLargeLargeEntryXml
		for _, elt := range s.LargeEntry {
			var obj typeLargeLargeEntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.LargeEntry = &typeLargeLargeEntryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeLargeXml) UnmarshalToObject() (*TypeLarge, error) {
	var largeEntryVal []TypeLargeLargeEntry
	if o.LargeEntry != nil {
		for _, elt := range o.LargeEntry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			largeEntryVal = append(largeEntryVal, *obj)
		}
	}

	result := &TypeLarge{
		LargeEntry:     largeEntryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeLargeLargeEntryXml) MarshalFromObject(s TypeLargeLargeEntry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.LcRegex != nil {
		o.LcRegex = util.StrToMem(s.LcRegex)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeLargeLargeEntryXml) UnmarshalToObject() (*TypeLargeLargeEntry, error) {
	var lcRegexVal []string
	if o.LcRegex != nil {
		lcRegexVal = util.MemToStr(o.LcRegex)
	}

	result := &TypeLargeLargeEntry{
		Name:           o.Name,
		Action:         o.Action,
		LcRegex:        lcRegexVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeRegularXml) MarshalFromObject(s TypeRegular) {
	if s.RegularEntry != nil {
		var objs []typeRegularRegularEntryXml
		for _, elt := range s.RegularEntry {
			var obj typeRegularRegularEntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RegularEntry = &typeRegularRegularEntryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeRegularXml) UnmarshalToObject() (*TypeRegular, error) {
	var regularEntryVal []TypeRegularRegularEntry
	if o.RegularEntry != nil {
		for _, elt := range o.RegularEntry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			regularEntryVal = append(regularEntryVal, *obj)
		}
	}

	result := &TypeRegular{
		RegularEntry:   regularEntryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeRegularRegularEntryXml) MarshalFromObject(s TypeRegularRegularEntry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.Community != nil {
		o.Community = util.StrToMem(s.Community)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeRegularRegularEntryXml) UnmarshalToObject() (*TypeRegularRegularEntry, error) {
	var communityVal []string
	if o.Community != nil {
		communityVal = util.MemToStr(o.Community)
	}

	result := &TypeRegularRegularEntry{
		Name:           o.Name,
		Action:         o.Action,
		Community:      communityVal,
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
	if v == "type" || v == "Type" {
		return e.Type, nil
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
	if !o.Type.matches(other.Type) {
		return false
	}

	return true
}

func (o *Type) matches(other *Type) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Extended.matches(other.Extended) {
		return false
	}
	if !o.Large.matches(other.Large) {
		return false
	}
	if !o.Regular.matches(other.Regular) {
		return false
	}

	return true
}

func (o *TypeExtended) matches(other *TypeExtended) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.ExtendedEntry) != len(other.ExtendedEntry) {
		return false
	}
	for idx := range o.ExtendedEntry {
		if !o.ExtendedEntry[idx].matches(&other.ExtendedEntry[idx]) {
			return false
		}
	}

	return true
}

func (o *TypeExtendedExtendedEntry) matches(other *TypeExtendedExtendedEntry) bool {
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
	if !util.OrderedListsMatch[string](o.EcRegex, other.EcRegex) {
		return false
	}

	return true
}

func (o *TypeLarge) matches(other *TypeLarge) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.LargeEntry) != len(other.LargeEntry) {
		return false
	}
	for idx := range o.LargeEntry {
		if !o.LargeEntry[idx].matches(&other.LargeEntry[idx]) {
			return false
		}
	}

	return true
}

func (o *TypeLargeLargeEntry) matches(other *TypeLargeLargeEntry) bool {
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
	if !util.OrderedListsMatch[string](o.LcRegex, other.LcRegex) {
		return false
	}

	return true
}

func (o *TypeRegular) matches(other *TypeRegular) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.RegularEntry) != len(other.RegularEntry) {
		return false
	}
	for idx := range o.RegularEntry {
		if !o.RegularEntry[idx].matches(&other.RegularEntry[idx]) {
			return false
		}
	}

	return true
}

func (o *TypeRegularRegularEntry) matches(other *TypeRegularRegularEntry) bool {
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
	if !util.OrderedListsMatch[string](o.Community, other.Community) {
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
