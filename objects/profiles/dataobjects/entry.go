package dataobjects

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
	suffix = []string{"profiles", "data-objects", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	PatternType     *PatternType
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type PatternType struct {
	FileProperties *PatternTypeFileProperties
	Predefined     *PatternTypePredefined
	Regex          *PatternTypeRegex
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypeFileProperties struct {
	Pattern        []PatternTypeFilePropertiesPattern
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypeFilePropertiesPattern struct {
	Name           string
	FileType       *string
	FileProperty   *string
	PropertyValue  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypePredefined struct {
	Pattern        []PatternTypePredefinedPattern
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypePredefinedPattern struct {
	Name           string
	FileType       []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypeRegex struct {
	Pattern        []PatternTypeRegexPattern
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PatternTypeRegexPattern struct {
	Name           string
	FileType       []string
	Regex          *string
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
	XMLName         xml.Name        `xml:"entry"`
	Name            string          `xml:"name,attr"`
	Description     *string         `xml:"description,omitempty"`
	DisableOverride *string         `xml:"disable-override,omitempty"`
	PatternType     *patternTypeXml `xml:"pattern-type,omitempty"`
	Misc            []generic.Xml   `xml:",any"`
	MiscAttributes  []xml.Attr      `xml:",any,attr"`
}
type patternTypeXml struct {
	FileProperties *patternTypeFilePropertiesXml `xml:"file-properties,omitempty"`
	Predefined     *patternTypePredefinedXml     `xml:"predefined,omitempty"`
	Regex          *patternTypeRegexXml          `xml:"regex,omitempty"`
	Misc           []generic.Xml                 `xml:",any"`
	MiscAttributes []xml.Attr                    `xml:",any,attr"`
}
type patternTypeFilePropertiesXml struct {
	Pattern        *patternTypeFilePropertiesPatternContainerXml `xml:"pattern,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type patternTypeFilePropertiesPatternContainerXml struct {
	Entries []patternTypeFilePropertiesPatternXml `xml:"entry"`
}
type patternTypeFilePropertiesPatternXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	FileType       *string       `xml:"file-type,omitempty"`
	FileProperty   *string       `xml:"file-property,omitempty"`
	PropertyValue  *string       `xml:"property-value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type patternTypePredefinedXml struct {
	Pattern        *patternTypePredefinedPatternContainerXml `xml:"pattern,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
	MiscAttributes []xml.Attr                                `xml:",any,attr"`
}
type patternTypePredefinedPatternContainerXml struct {
	Entries []patternTypePredefinedPatternXml `xml:"entry"`
}
type patternTypePredefinedPatternXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	FileType       *util.MemberType `xml:"file-type,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type patternTypeRegexXml struct {
	Pattern        *patternTypeRegexPatternContainerXml `xml:"pattern,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type patternTypeRegexPatternContainerXml struct {
	Entries []patternTypeRegexPatternXml `xml:"entry"`
}
type patternTypeRegexPatternXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	FileType       *util.MemberType `xml:"file-type,omitempty"`
	Regex          *string          `xml:"regex,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.PatternType != nil {
		var obj patternTypeXml
		obj.MarshalFromObject(*s.PatternType)
		o.PatternType = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var patternTypeVal *PatternType
	if o.PatternType != nil {
		obj, err := o.PatternType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		patternTypeVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		DisableOverride: o.DisableOverride,
		PatternType:     patternTypeVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypeXml) MarshalFromObject(s PatternType) {
	if s.FileProperties != nil {
		var obj patternTypeFilePropertiesXml
		obj.MarshalFromObject(*s.FileProperties)
		o.FileProperties = &obj
	}
	if s.Predefined != nil {
		var obj patternTypePredefinedXml
		obj.MarshalFromObject(*s.Predefined)
		o.Predefined = &obj
	}
	if s.Regex != nil {
		var obj patternTypeRegexXml
		obj.MarshalFromObject(*s.Regex)
		o.Regex = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypeXml) UnmarshalToObject() (*PatternType, error) {
	var filePropertiesVal *PatternTypeFileProperties
	if o.FileProperties != nil {
		obj, err := o.FileProperties.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filePropertiesVal = obj
	}
	var predefinedVal *PatternTypePredefined
	if o.Predefined != nil {
		obj, err := o.Predefined.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		predefinedVal = obj
	}
	var regexVal *PatternTypeRegex
	if o.Regex != nil {
		obj, err := o.Regex.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		regexVal = obj
	}

	result := &PatternType{
		FileProperties: filePropertiesVal,
		Predefined:     predefinedVal,
		Regex:          regexVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypeFilePropertiesXml) MarshalFromObject(s PatternTypeFileProperties) {
	if s.Pattern != nil {
		var objs []patternTypeFilePropertiesPatternXml
		for _, elt := range s.Pattern {
			var obj patternTypeFilePropertiesPatternXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Pattern = &patternTypeFilePropertiesPatternContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypeFilePropertiesXml) UnmarshalToObject() (*PatternTypeFileProperties, error) {
	var patternVal []PatternTypeFilePropertiesPattern
	if o.Pattern != nil {
		for _, elt := range o.Pattern.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			patternVal = append(patternVal, *obj)
		}
	}

	result := &PatternTypeFileProperties{
		Pattern:        patternVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypeFilePropertiesPatternXml) MarshalFromObject(s PatternTypeFilePropertiesPattern) {
	o.Name = s.Name
	o.FileType = s.FileType
	o.FileProperty = s.FileProperty
	o.PropertyValue = s.PropertyValue
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypeFilePropertiesPatternXml) UnmarshalToObject() (*PatternTypeFilePropertiesPattern, error) {

	result := &PatternTypeFilePropertiesPattern{
		Name:           o.Name,
		FileType:       o.FileType,
		FileProperty:   o.FileProperty,
		PropertyValue:  o.PropertyValue,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypePredefinedXml) MarshalFromObject(s PatternTypePredefined) {
	if s.Pattern != nil {
		var objs []patternTypePredefinedPatternXml
		for _, elt := range s.Pattern {
			var obj patternTypePredefinedPatternXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Pattern = &patternTypePredefinedPatternContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypePredefinedXml) UnmarshalToObject() (*PatternTypePredefined, error) {
	var patternVal []PatternTypePredefinedPattern
	if o.Pattern != nil {
		for _, elt := range o.Pattern.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			patternVal = append(patternVal, *obj)
		}
	}

	result := &PatternTypePredefined{
		Pattern:        patternVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypePredefinedPatternXml) MarshalFromObject(s PatternTypePredefinedPattern) {
	o.Name = s.Name
	if s.FileType != nil {
		o.FileType = util.StrToMem(s.FileType)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypePredefinedPatternXml) UnmarshalToObject() (*PatternTypePredefinedPattern, error) {
	var fileTypeVal []string
	if o.FileType != nil {
		fileTypeVal = util.MemToStr(o.FileType)
	}

	result := &PatternTypePredefinedPattern{
		Name:           o.Name,
		FileType:       fileTypeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypeRegexXml) MarshalFromObject(s PatternTypeRegex) {
	if s.Pattern != nil {
		var objs []patternTypeRegexPatternXml
		for _, elt := range s.Pattern {
			var obj patternTypeRegexPatternXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Pattern = &patternTypeRegexPatternContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypeRegexXml) UnmarshalToObject() (*PatternTypeRegex, error) {
	var patternVal []PatternTypeRegexPattern
	if o.Pattern != nil {
		for _, elt := range o.Pattern.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			patternVal = append(patternVal, *obj)
		}
	}

	result := &PatternTypeRegex{
		Pattern:        patternVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *patternTypeRegexPatternXml) MarshalFromObject(s PatternTypeRegexPattern) {
	o.Name = s.Name
	if s.FileType != nil {
		o.FileType = util.StrToMem(s.FileType)
	}
	o.Regex = s.Regex
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o patternTypeRegexPatternXml) UnmarshalToObject() (*PatternTypeRegexPattern, error) {
	var fileTypeVal []string
	if o.FileType != nil {
		fileTypeVal = util.MemToStr(o.FileType)
	}

	result := &PatternTypeRegexPattern{
		Name:           o.Name,
		FileType:       fileTypeVal,
		Regex:          o.Regex,
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
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "pattern_type" || v == "PatternType" {
		return e.PatternType, nil
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
	if !o.PatternType.matches(other.PatternType) {
		return false
	}

	return true
}

func (o *PatternType) matches(other *PatternType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.FileProperties.matches(other.FileProperties) {
		return false
	}
	if !o.Predefined.matches(other.Predefined) {
		return false
	}
	if !o.Regex.matches(other.Regex) {
		return false
	}

	return true
}

func (o *PatternTypeFileProperties) matches(other *PatternTypeFileProperties) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Pattern) != len(other.Pattern) {
		return false
	}
	for idx := range o.Pattern {
		if !o.Pattern[idx].matches(&other.Pattern[idx]) {
			return false
		}
	}

	return true
}

func (o *PatternTypeFilePropertiesPattern) matches(other *PatternTypeFilePropertiesPattern) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.FileType, other.FileType) {
		return false
	}
	if !util.StringsMatch(o.FileProperty, other.FileProperty) {
		return false
	}
	if !util.StringsMatch(o.PropertyValue, other.PropertyValue) {
		return false
	}

	return true
}

func (o *PatternTypePredefined) matches(other *PatternTypePredefined) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Pattern) != len(other.Pattern) {
		return false
	}
	for idx := range o.Pattern {
		if !o.Pattern[idx].matches(&other.Pattern[idx]) {
			return false
		}
	}

	return true
}

func (o *PatternTypePredefinedPattern) matches(other *PatternTypePredefinedPattern) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.FileType, other.FileType) {
		return false
	}

	return true
}

func (o *PatternTypeRegex) matches(other *PatternTypeRegex) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Pattern) != len(other.Pattern) {
		return false
	}
	for idx := range o.Pattern {
		if !o.Pattern[idx].matches(&other.Pattern[idx]) {
			return false
		}
	}

	return true
}

func (o *PatternTypeRegexPattern) matches(other *PatternTypeRegexPattern) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.FileType, other.FileType) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
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
