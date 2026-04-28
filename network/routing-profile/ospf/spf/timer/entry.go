package timer

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
	suffix = []string{"network", "routing-profile", "ospf", "spf-timer-profile", "$name"}
)

type Entry struct {
	Name                string
	InitialHoldTime     *int64
	LsaInterval         *int64
	MaxHoldTime         *int64
	SpfCalculationDelay *int64
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
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
	XMLName             xml.Name      `xml:"entry"`
	Name                string        `xml:"name,attr"`
	InitialHoldTime     *int64        `xml:"initial-hold-time,omitempty"`
	LsaInterval         *int64        `xml:"lsa-interval,omitempty"`
	MaxHoldTime         *int64        `xml:"max-hold-time,omitempty"`
	SpfCalculationDelay *int64        `xml:"spf-calculation-delay,omitempty"`
	Misc                []generic.Xml `xml:",any"`
	MiscAttributes      []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.InitialHoldTime = s.InitialHoldTime
	o.LsaInterval = s.LsaInterval
	o.MaxHoldTime = s.MaxHoldTime
	o.SpfCalculationDelay = s.SpfCalculationDelay
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:                o.Name,
		InitialHoldTime:     o.InitialHoldTime,
		LsaInterval:         o.LsaInterval,
		MaxHoldTime:         o.MaxHoldTime,
		SpfCalculationDelay: o.SpfCalculationDelay,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "initial_hold_time" || v == "InitialHoldTime" {
		return e.InitialHoldTime, nil
	}
	if v == "lsa_interval" || v == "LsaInterval" {
		return e.LsaInterval, nil
	}
	if v == "max_hold_time" || v == "MaxHoldTime" {
		return e.MaxHoldTime, nil
	}
	if v == "spf_calculation_delay" || v == "SpfCalculationDelay" {
		return e.SpfCalculationDelay, nil
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
	if !util.Ints64Match(o.InitialHoldTime, other.InitialHoldTime) {
		return false
	}
	if !util.Ints64Match(o.LsaInterval, other.LsaInterval) {
		return false
	}
	if !util.Ints64Match(o.MaxHoldTime, other.MaxHoldTime) {
		return false
	}
	if !util.Ints64Match(o.SpfCalculationDelay, other.SpfCalculationDelay) {
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
