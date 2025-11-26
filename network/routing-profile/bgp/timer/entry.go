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
	suffix = []string{"network", "routing-profile", "bgp", "timer-profile", "$name"}
)

type Entry struct {
	Name                   string
	HoldTime               *string
	KeepAliveInterval      *string
	MinRouteAdvInterval    *int64
	OpenDelayTime          *int64
	ReconnectRetryInterval *int64
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
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
	XMLName                xml.Name      `xml:"entry"`
	Name                   string        `xml:"name,attr"`
	HoldTime               *string       `xml:"hold-time,omitempty"`
	KeepAliveInterval      *string       `xml:"keep-alive-interval,omitempty"`
	MinRouteAdvInterval    *int64        `xml:"min-route-adv-interval,omitempty"`
	OpenDelayTime          *int64        `xml:"open-delay-time,omitempty"`
	ReconnectRetryInterval *int64        `xml:"reconnect-retry-interval,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
	MiscAttributes         []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.HoldTime = s.HoldTime
	o.KeepAliveInterval = s.KeepAliveInterval
	o.MinRouteAdvInterval = s.MinRouteAdvInterval
	o.OpenDelayTime = s.OpenDelayTime
	o.ReconnectRetryInterval = s.ReconnectRetryInterval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:                   o.Name,
		HoldTime:               o.HoldTime,
		KeepAliveInterval:      o.KeepAliveInterval,
		MinRouteAdvInterval:    o.MinRouteAdvInterval,
		OpenDelayTime:          o.OpenDelayTime,
		ReconnectRetryInterval: o.ReconnectRetryInterval,
		Misc:                   o.Misc,
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "hold_time" || v == "HoldTime" {
		return e.HoldTime, nil
	}
	if v == "keep_alive_interval" || v == "KeepAliveInterval" {
		return e.KeepAliveInterval, nil
	}
	if v == "min_route_adv_interval" || v == "MinRouteAdvInterval" {
		return e.MinRouteAdvInterval, nil
	}
	if v == "open_delay_time" || v == "OpenDelayTime" {
		return e.OpenDelayTime, nil
	}
	if v == "reconnect_retry_interval" || v == "ReconnectRetryInterval" {
		return e.ReconnectRetryInterval, nil
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
	if !util.StringsMatch(o.HoldTime, other.HoldTime) {
		return false
	}
	if !util.StringsMatch(o.KeepAliveInterval, other.KeepAliveInterval) {
		return false
	}
	if !util.Ints64Match(o.MinRouteAdvInterval, other.MinRouteAdvInterval) {
		return false
	}
	if !util.Ints64Match(o.OpenDelayTime, other.OpenDelayTime) {
		return false
	}
	if !util.Ints64Match(o.ReconnectRetryInterval, other.ReconnectRetryInterval) {
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
