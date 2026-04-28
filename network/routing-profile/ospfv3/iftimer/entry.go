package iftimer

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
	suffix = []string{"network", "routing-profile", "ospfv3", "if-timer-profile", "$name"}
)

type Entry struct {
	Name               string
	DeadCounts         *int64
	GrDelay            *int64
	HelloInterval      *int64
	RetransmitInterval *int64
	TransitDelay       *int64
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
	XMLName            xml.Name      `xml:"entry"`
	Name               string        `xml:"name,attr"`
	DeadCounts         *int64        `xml:"dead-counts,omitempty"`
	GrDelay            *int64        `xml:"gr-delay,omitempty"`
	HelloInterval      *int64        `xml:"hello-interval,omitempty"`
	RetransmitInterval *int64        `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64        `xml:"transit-delay,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DeadCounts = s.DeadCounts
	o.GrDelay = s.GrDelay
	o.HelloInterval = s.HelloInterval
	o.RetransmitInterval = s.RetransmitInterval
	o.TransitDelay = s.TransitDelay
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:               o.Name,
		DeadCounts:         o.DeadCounts,
		GrDelay:            o.GrDelay,
		HelloInterval:      o.HelloInterval,
		RetransmitInterval: o.RetransmitInterval,
		TransitDelay:       o.TransitDelay,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "dead_counts" || v == "DeadCounts" {
		return e.DeadCounts, nil
	}
	if v == "gr_delay" || v == "GrDelay" {
		return e.GrDelay, nil
	}
	if v == "hello_interval" || v == "HelloInterval" {
		return e.HelloInterval, nil
	}
	if v == "retransmit_interval" || v == "RetransmitInterval" {
		return e.RetransmitInterval, nil
	}
	if v == "transit_delay" || v == "TransitDelay" {
		return e.TransitDelay, nil
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
	if !util.Ints64Match(o.DeadCounts, other.DeadCounts) {
		return false
	}
	if !util.Ints64Match(o.GrDelay, other.GrDelay) {
		return false
	}
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.RetransmitInterval, other.RetransmitInterval) {
		return false
	}
	if !util.Ints64Match(o.TransitDelay, other.TransitDelay) {
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
