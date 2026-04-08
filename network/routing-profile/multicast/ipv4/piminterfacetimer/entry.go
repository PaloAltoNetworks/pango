package piminterfacetimer

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
	suffix = []string{"network", "routing-profile", "multicast", "pim-interface-timer-profile", "$name"}
)

type Entry struct {
	Name              string
	AssertInterval    *int64
	HelloInterval     *int64
	JoinPruneInterval *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
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
	XMLName           xml.Name      `xml:"entry"`
	Name              string        `xml:"name,attr"`
	AssertInterval    *int64        `xml:"assert-interval,omitempty"`
	HelloInterval     *int64        `xml:"hello-interval,omitempty"`
	JoinPruneInterval *int64        `xml:"join-prune-interval,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AssertInterval = s.AssertInterval
	o.HelloInterval = s.HelloInterval
	o.JoinPruneInterval = s.JoinPruneInterval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:              o.Name,
		AssertInterval:    o.AssertInterval,
		HelloInterval:     o.HelloInterval,
		JoinPruneInterval: o.JoinPruneInterval,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "assert_interval" || v == "AssertInterval" {
		return e.AssertInterval, nil
	}
	if v == "hello_interval" || v == "HelloInterval" {
		return e.HelloInterval, nil
	}
	if v == "join_prune_interval" || v == "JoinPruneInterval" {
		return e.JoinPruneInterval, nil
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
	if !util.Ints64Match(o.AssertInterval, other.AssertInterval) {
		return false
	}
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.JoinPruneInterval, other.JoinPruneInterval) {
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
