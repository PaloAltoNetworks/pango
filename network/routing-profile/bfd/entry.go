package bfd

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
	suffix = []string{"network", "routing-profile", "bfd", "$name"}
)

type Entry struct {
	Name                string
	DetectionMultiplier *int64
	HoldTime            *int64
	MinRxInterval       *int64
	MinTxInterval       *int64
	Mode                *string
	Multihop            *Multihop
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type Multihop struct {
	MinReceivedTtl *int64
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
	XMLName             xml.Name      `xml:"entry"`
	Name                string        `xml:"name,attr"`
	DetectionMultiplier *int64        `xml:"detection-multiplier,omitempty"`
	HoldTime            *int64        `xml:"hold-time,omitempty"`
	MinRxInterval       *int64        `xml:"min-rx-interval,omitempty"`
	MinTxInterval       *int64        `xml:"min-tx-interval,omitempty"`
	Mode                *string       `xml:"mode,omitempty"`
	Multihop            *multihopXml  `xml:"multihop,omitempty"`
	Misc                []generic.Xml `xml:",any"`
	MiscAttributes      []xml.Attr    `xml:",any,attr"`
}
type multihopXml struct {
	MinReceivedTtl *int64        `xml:"min-received-ttl,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DetectionMultiplier = s.DetectionMultiplier
	o.HoldTime = s.HoldTime
	o.MinRxInterval = s.MinRxInterval
	o.MinTxInterval = s.MinTxInterval
	o.Mode = s.Mode
	if s.Multihop != nil {
		var obj multihopXml
		obj.MarshalFromObject(*s.Multihop)
		o.Multihop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var multihopVal *Multihop
	if o.Multihop != nil {
		obj, err := o.Multihop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multihopVal = obj
	}

	result := &Entry{
		Name:                o.Name,
		DetectionMultiplier: o.DetectionMultiplier,
		HoldTime:            o.HoldTime,
		MinRxInterval:       o.MinRxInterval,
		MinTxInterval:       o.MinTxInterval,
		Mode:                o.Mode,
		Multihop:            multihopVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *multihopXml) MarshalFromObject(s Multihop) {
	o.MinReceivedTtl = s.MinReceivedTtl
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o multihopXml) UnmarshalToObject() (*Multihop, error) {

	result := &Multihop{
		MinReceivedTtl: o.MinReceivedTtl,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "detection_multiplier" || v == "DetectionMultiplier" {
		return e.DetectionMultiplier, nil
	}
	if v == "hold_time" || v == "HoldTime" {
		return e.HoldTime, nil
	}
	if v == "min_rx_interval" || v == "MinRxInterval" {
		return e.MinRxInterval, nil
	}
	if v == "min_tx_interval" || v == "MinTxInterval" {
		return e.MinTxInterval, nil
	}
	if v == "mode" || v == "Mode" {
		return e.Mode, nil
	}
	if v == "multihop" || v == "Multihop" {
		return e.Multihop, nil
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
	if !util.Ints64Match(o.DetectionMultiplier, other.DetectionMultiplier) {
		return false
	}
	if !util.Ints64Match(o.HoldTime, other.HoldTime) {
		return false
	}
	if !util.Ints64Match(o.MinRxInterval, other.MinRxInterval) {
		return false
	}
	if !util.Ints64Match(o.MinTxInterval, other.MinTxInterval) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !o.Multihop.matches(other.Multihop) {
		return false
	}

	return true
}

func (o *Multihop) matches(other *Multihop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.MinReceivedTtl, other.MinReceivedTtl) {
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
