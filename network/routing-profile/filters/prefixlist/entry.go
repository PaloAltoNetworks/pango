package prefixlist

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
	suffix = []string{"network", "routing-profile", "filters", "prefix-list", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Type           *Type
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Type struct {
	Ipv4           *TypeIpv4
	Ipv6           *TypeIpv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv4 struct {
	Ipv4Entry      []TypeIpv4Ipv4Entry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv4Ipv4Entry struct {
	Name           string
	Action         *string
	Prefix         *TypeIpv4Ipv4EntryPrefix
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv4Ipv4EntryPrefix struct {
	Network        *string
	Entry          *TypeIpv4Ipv4EntryPrefixEntry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv4Ipv4EntryPrefixEntry struct {
	Network            *string
	GreaterThanOrEqual *int64
	LessThanOrEqual    *int64
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type TypeIpv6 struct {
	Ipv6Entry      []TypeIpv6Ipv6Entry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv6Ipv6Entry struct {
	Name           string
	Action         *string
	Prefix         *TypeIpv6Ipv6EntryPrefix
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv6Ipv6EntryPrefix struct {
	Network        *string
	Entry          *TypeIpv6Ipv6EntryPrefixEntry
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TypeIpv6Ipv6EntryPrefixEntry struct {
	Network            *string
	GreaterThanOrEqual *int64
	LessThanOrEqual    *int64
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
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
	Ipv4           *typeIpv4Xml  `xml:"ipv4,omitempty"`
	Ipv6           *typeIpv6Xml  `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type typeIpv4Xml struct {
	Ipv4Entry      *typeIpv4Ipv4EntryContainerXml `xml:"ipv4-entry,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type typeIpv4Ipv4EntryContainerXml struct {
	Entries []typeIpv4Ipv4EntryXml `xml:"entry"`
}
type typeIpv4Ipv4EntryXml struct {
	XMLName        xml.Name                    `xml:"entry"`
	Name           string                      `xml:"name,attr"`
	Action         *string                     `xml:"action,omitempty"`
	Prefix         *typeIpv4Ipv4EntryPrefixXml `xml:"prefix,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type typeIpv4Ipv4EntryPrefixXml struct {
	Network        *string                          `xml:"network,omitempty"`
	Entry          *typeIpv4Ipv4EntryPrefixEntryXml `xml:"entry,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type typeIpv4Ipv4EntryPrefixEntryXml struct {
	Network            *string       `xml:"network,omitempty"`
	GreaterThanOrEqual *int64        `xml:"greater-than-or-equal,omitempty"`
	LessThanOrEqual    *int64        `xml:"less-than-or-equal,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type typeIpv6Xml struct {
	Ipv6Entry      *typeIpv6Ipv6EntryContainerXml `xml:"ipv6-entry,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type typeIpv6Ipv6EntryContainerXml struct {
	Entries []typeIpv6Ipv6EntryXml `xml:"entry"`
}
type typeIpv6Ipv6EntryXml struct {
	XMLName        xml.Name                    `xml:"entry"`
	Name           string                      `xml:"name,attr"`
	Action         *string                     `xml:"action,omitempty"`
	Prefix         *typeIpv6Ipv6EntryPrefixXml `xml:"prefix,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type typeIpv6Ipv6EntryPrefixXml struct {
	Network        *string                          `xml:"network,omitempty"`
	Entry          *typeIpv6Ipv6EntryPrefixEntryXml `xml:"entry,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type typeIpv6Ipv6EntryPrefixEntryXml struct {
	Network            *string       `xml:"network,omitempty"`
	GreaterThanOrEqual *int64        `xml:"greater-than-or-equal,omitempty"`
	LessThanOrEqual    *int64        `xml:"less-than-or-equal,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
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
	if s.Ipv4 != nil {
		var obj typeIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj typeIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeXml) UnmarshalToObject() (*Type, error) {
	var ipv4Val *TypeIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *TypeIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &Type{
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv4Xml) MarshalFromObject(s TypeIpv4) {
	if s.Ipv4Entry != nil {
		var objs []typeIpv4Ipv4EntryXml
		for _, elt := range s.Ipv4Entry {
			var obj typeIpv4Ipv4EntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ipv4Entry = &typeIpv4Ipv4EntryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv4Xml) UnmarshalToObject() (*TypeIpv4, error) {
	var ipv4EntryVal []TypeIpv4Ipv4Entry
	if o.Ipv4Entry != nil {
		for _, elt := range o.Ipv4Entry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipv4EntryVal = append(ipv4EntryVal, *obj)
		}
	}

	result := &TypeIpv4{
		Ipv4Entry:      ipv4EntryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv4Ipv4EntryXml) MarshalFromObject(s TypeIpv4Ipv4Entry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.Prefix != nil {
		var obj typeIpv4Ipv4EntryPrefixXml
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv4Ipv4EntryXml) UnmarshalToObject() (*TypeIpv4Ipv4Entry, error) {
	var prefixVal *TypeIpv4Ipv4EntryPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}

	result := &TypeIpv4Ipv4Entry{
		Name:           o.Name,
		Action:         o.Action,
		Prefix:         prefixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv4Ipv4EntryPrefixXml) MarshalFromObject(s TypeIpv4Ipv4EntryPrefix) {
	o.Network = s.Network
	if s.Entry != nil {
		var obj typeIpv4Ipv4EntryPrefixEntryXml
		obj.MarshalFromObject(*s.Entry)
		o.Entry = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv4Ipv4EntryPrefixXml) UnmarshalToObject() (*TypeIpv4Ipv4EntryPrefix, error) {
	var entryVal *TypeIpv4Ipv4EntryPrefixEntry
	if o.Entry != nil {
		obj, err := o.Entry.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entryVal = obj
	}

	result := &TypeIpv4Ipv4EntryPrefix{
		Network:        o.Network,
		Entry:          entryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv4Ipv4EntryPrefixEntryXml) MarshalFromObject(s TypeIpv4Ipv4EntryPrefixEntry) {
	o.Network = s.Network
	o.GreaterThanOrEqual = s.GreaterThanOrEqual
	o.LessThanOrEqual = s.LessThanOrEqual
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv4Ipv4EntryPrefixEntryXml) UnmarshalToObject() (*TypeIpv4Ipv4EntryPrefixEntry, error) {

	result := &TypeIpv4Ipv4EntryPrefixEntry{
		Network:            o.Network,
		GreaterThanOrEqual: o.GreaterThanOrEqual,
		LessThanOrEqual:    o.LessThanOrEqual,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv6Xml) MarshalFromObject(s TypeIpv6) {
	if s.Ipv6Entry != nil {
		var objs []typeIpv6Ipv6EntryXml
		for _, elt := range s.Ipv6Entry {
			var obj typeIpv6Ipv6EntryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ipv6Entry = &typeIpv6Ipv6EntryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv6Xml) UnmarshalToObject() (*TypeIpv6, error) {
	var ipv6EntryVal []TypeIpv6Ipv6Entry
	if o.Ipv6Entry != nil {
		for _, elt := range o.Ipv6Entry.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipv6EntryVal = append(ipv6EntryVal, *obj)
		}
	}

	result := &TypeIpv6{
		Ipv6Entry:      ipv6EntryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv6Ipv6EntryXml) MarshalFromObject(s TypeIpv6Ipv6Entry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.Prefix != nil {
		var obj typeIpv6Ipv6EntryPrefixXml
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv6Ipv6EntryXml) UnmarshalToObject() (*TypeIpv6Ipv6Entry, error) {
	var prefixVal *TypeIpv6Ipv6EntryPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}

	result := &TypeIpv6Ipv6Entry{
		Name:           o.Name,
		Action:         o.Action,
		Prefix:         prefixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv6Ipv6EntryPrefixXml) MarshalFromObject(s TypeIpv6Ipv6EntryPrefix) {
	o.Network = s.Network
	if s.Entry != nil {
		var obj typeIpv6Ipv6EntryPrefixEntryXml
		obj.MarshalFromObject(*s.Entry)
		o.Entry = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv6Ipv6EntryPrefixXml) UnmarshalToObject() (*TypeIpv6Ipv6EntryPrefix, error) {
	var entryVal *TypeIpv6Ipv6EntryPrefixEntry
	if o.Entry != nil {
		obj, err := o.Entry.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entryVal = obj
	}

	result := &TypeIpv6Ipv6EntryPrefix{
		Network:        o.Network,
		Entry:          entryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *typeIpv6Ipv6EntryPrefixEntryXml) MarshalFromObject(s TypeIpv6Ipv6EntryPrefixEntry) {
	o.Network = s.Network
	o.GreaterThanOrEqual = s.GreaterThanOrEqual
	o.LessThanOrEqual = s.LessThanOrEqual
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o typeIpv6Ipv6EntryPrefixEntryXml) UnmarshalToObject() (*TypeIpv6Ipv6EntryPrefixEntry, error) {

	result := &TypeIpv6Ipv6EntryPrefixEntry{
		Network:            o.Network,
		GreaterThanOrEqual: o.GreaterThanOrEqual,
		LessThanOrEqual:    o.LessThanOrEqual,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
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
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *TypeIpv4) matches(other *TypeIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Ipv4Entry) != len(other.Ipv4Entry) {
		return false
	}
	for idx := range o.Ipv4Entry {
		if !o.Ipv4Entry[idx].matches(&other.Ipv4Entry[idx]) {
			return false
		}
	}

	return true
}

func (o *TypeIpv4Ipv4Entry) matches(other *TypeIpv4Ipv4Entry) bool {
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
	if !o.Prefix.matches(other.Prefix) {
		return false
	}

	return true
}

func (o *TypeIpv4Ipv4EntryPrefix) matches(other *TypeIpv4Ipv4EntryPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Network, other.Network) {
		return false
	}
	if !o.Entry.matches(other.Entry) {
		return false
	}

	return true
}

func (o *TypeIpv4Ipv4EntryPrefixEntry) matches(other *TypeIpv4Ipv4EntryPrefixEntry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Network, other.Network) {
		return false
	}
	if !util.Ints64Match(o.GreaterThanOrEqual, other.GreaterThanOrEqual) {
		return false
	}
	if !util.Ints64Match(o.LessThanOrEqual, other.LessThanOrEqual) {
		return false
	}

	return true
}

func (o *TypeIpv6) matches(other *TypeIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Ipv6Entry) != len(other.Ipv6Entry) {
		return false
	}
	for idx := range o.Ipv6Entry {
		if !o.Ipv6Entry[idx].matches(&other.Ipv6Entry[idx]) {
			return false
		}
	}

	return true
}

func (o *TypeIpv6Ipv6Entry) matches(other *TypeIpv6Ipv6Entry) bool {
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
	if !o.Prefix.matches(other.Prefix) {
		return false
	}

	return true
}

func (o *TypeIpv6Ipv6EntryPrefix) matches(other *TypeIpv6Ipv6EntryPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Network, other.Network) {
		return false
	}
	if !o.Entry.matches(other.Entry) {
		return false
	}

	return true
}

func (o *TypeIpv6Ipv6EntryPrefixEntry) matches(other *TypeIpv6Ipv6EntryPrefixEntry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Network, other.Network) {
		return false
	}
	if !util.Ints64Match(o.GreaterThanOrEqual, other.GreaterThanOrEqual) {
		return false
	}
	if !util.Ints64Match(o.LessThanOrEqual, other.LessThanOrEqual) {
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
