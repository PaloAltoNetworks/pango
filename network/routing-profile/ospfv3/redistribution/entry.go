package redistribution

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
	suffix = []string{"network", "routing-profile", "ospfv3", "redistribution-profile", "$name"}
)

type Entry struct {
	Name           string
	Bgp            *Bgp
	Connected      *Connected
	DefaultRoute   *DefaultRoute
	Static         *Static
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Bgp struct {
	Enable         *bool
	Metric         *int64
	MetricType     *string
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Connected struct {
	Enable         *bool
	Metric         *int64
	MetricType     *string
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultRoute struct {
	Always         *bool
	Enable         *bool
	Metric         *int64
	MetricType     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Static struct {
	Enable         *bool
	Metric         *int64
	MetricType     *string
	RouteMap       *string
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
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Bgp            *bgpXml          `xml:"bgp,omitempty"`
	Connected      *connectedXml    `xml:"connected,omitempty"`
	DefaultRoute   *defaultRouteXml `xml:"default-route,omitempty"`
	Static         *staticXml       `xml:"static,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type bgpXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	MetricType     *string       `xml:"metric-type,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	MetricType     *string       `xml:"metric-type,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultRouteXml struct {
	Always         *string       `xml:"always,omitempty"`
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	MetricType     *string       `xml:"metric-type,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type staticXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	MetricType     *string       `xml:"metric-type,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Bgp != nil {
		var obj bgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Connected != nil {
		var obj connectedXml
		obj.MarshalFromObject(*s.Connected)
		o.Connected = &obj
	}
	if s.DefaultRoute != nil {
		var obj defaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	if s.Static != nil {
		var obj staticXml
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var bgpVal *Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var connectedVal *Connected
	if o.Connected != nil {
		obj, err := o.Connected.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectedVal = obj
	}
	var defaultRouteVal *DefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}
	var staticVal *Static
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Bgp:            bgpVal,
		Connected:      connectedVal,
		DefaultRoute:   defaultRouteVal,
		Static:         staticVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpXml) MarshalFromObject(s Bgp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpXml) UnmarshalToObject() (*Bgp, error) {

	result := &Bgp{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		MetricType:     o.MetricType,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedXml) MarshalFromObject(s Connected) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedXml) UnmarshalToObject() (*Connected, error) {

	result := &Connected{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		MetricType:     o.MetricType,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultRouteXml) MarshalFromObject(s DefaultRoute) {
	o.Always = util.YesNo(s.Always, nil)
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultRouteXml) UnmarshalToObject() (*DefaultRoute, error) {

	result := &DefaultRoute{
		Always:         util.AsBool(o.Always, nil),
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		MetricType:     o.MetricType,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *staticXml) MarshalFromObject(s Static) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o staticXml) UnmarshalToObject() (*Static, error) {

	result := &Static{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		MetricType:     o.MetricType,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "bgp" || v == "Bgp" {
		return e.Bgp, nil
	}
	if v == "connected" || v == "Connected" {
		return e.Connected, nil
	}
	if v == "default_route" || v == "DefaultRoute" {
		return e.DefaultRoute, nil
	}
	if v == "static" || v == "Static" {
		return e.Static, nil
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
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Connected.matches(other.Connected) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
		return false
	}
	if !o.Static.matches(other.Static) {
		return false
	}

	return true
}

func (o *Bgp) matches(other *Bgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Connected) matches(other *Connected) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *DefaultRoute) matches(other *DefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Always, other.Always) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}

	return true
}

func (o *Static) matches(other *Static) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
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
