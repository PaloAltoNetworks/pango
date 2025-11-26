package dampening

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
	suffix = []string{"network", "routing-profile", "bgp", "dampening-profile", "$name"}
)

type Entry struct {
	Name             string
	Description      *string
	HalfLife         *int64
	MaxSuppressLimit *int64
	ReuseLimit       *int64
	SuppressLimit    *int64
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
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
	XMLName          xml.Name      `xml:"entry"`
	Name             string        `xml:"name,attr"`
	Description      *string       `xml:"description,omitempty"`
	HalfLife         *int64        `xml:"half-life,omitempty"`
	MaxSuppressLimit *int64        `xml:"max-suppress-limit,omitempty"`
	ReuseLimit       *int64        `xml:"reuse-limit,omitempty"`
	SuppressLimit    *int64        `xml:"suppress-limit,omitempty"`
	Misc             []generic.Xml `xml:",any"`
	MiscAttributes   []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.HalfLife = s.HalfLife
	o.MaxSuppressLimit = s.MaxSuppressLimit
	o.ReuseLimit = s.ReuseLimit
	o.SuppressLimit = s.SuppressLimit
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:             o.Name,
		Description:      o.Description,
		HalfLife:         o.HalfLife,
		MaxSuppressLimit: o.MaxSuppressLimit,
		ReuseLimit:       o.ReuseLimit,
		SuppressLimit:    o.SuppressLimit,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
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
	if v == "half_life" || v == "HalfLife" {
		return e.HalfLife, nil
	}
	if v == "max_suppress_limit" || v == "MaxSuppressLimit" {
		return e.MaxSuppressLimit, nil
	}
	if v == "reuse_limit" || v == "ReuseLimit" {
		return e.ReuseLimit, nil
	}
	if v == "suppress_limit" || v == "SuppressLimit" {
		return e.SuppressLimit, nil
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
	if !util.Ints64Match(o.HalfLife, other.HalfLife) {
		return false
	}
	if !util.Ints64Match(o.MaxSuppressLimit, other.MaxSuppressLimit) {
		return false
	}
	if !util.Ints64Match(o.ReuseLimit, other.ReuseLimit) {
		return false
	}
	if !util.Ints64Match(o.SuppressLimit, other.SuppressLimit) {
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
