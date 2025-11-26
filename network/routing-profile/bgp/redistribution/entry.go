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
	suffix = []string{"network", "routing-profile", "bgp", "redistribution-profile", "$name"}
)

type Entry struct {
	Name           string
	Ipv4           *Ipv4
	Ipv6           *Ipv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4 struct {
	Unicast        *Ipv4Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4Unicast struct {
	Connected      *Ipv4UnicastConnected
	Ospf           *Ipv4UnicastOspf
	Rip            *Ipv4UnicastRip
	Static         *Ipv4UnicastStatic
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastConnected struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastOspf struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastRip struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastStatic struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6 struct {
	Unicast        *Ipv6Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6Unicast struct {
	Connected      *Ipv6UnicastConnected
	Ospfv3         *Ipv6UnicastOspfv3
	Static         *Ipv6UnicastStatic
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastConnected struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastOspfv3 struct {
	Enable         *bool
	Metric         *int64
	RouteMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastStatic struct {
	Enable         *bool
	Metric         *int64
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Ipv4           *ipv4Xml      `xml:"ipv4,omitempty"`
	Ipv6           *ipv6Xml      `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4Xml struct {
	Unicast        *ipv4UnicastXml `xml:"unicast,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type ipv4UnicastXml struct {
	Connected      *ipv4UnicastConnectedXml `xml:"connected,omitempty"`
	Ospf           *ipv4UnicastOspfXml      `xml:"ospf,omitempty"`
	Rip            *ipv4UnicastRipXml       `xml:"rip,omitempty"`
	Static         *ipv4UnicastStaticXml    `xml:"static,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ipv4UnicastConnectedXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastOspfXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastRipXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastStaticXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml struct {
	Unicast        *ipv6UnicastXml `xml:"unicast,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type ipv6UnicastXml struct {
	Connected      *ipv6UnicastConnectedXml `xml:"connected,omitempty"`
	Ospfv3         *ipv6UnicastOspfv3Xml    `xml:"ospfv3,omitempty"`
	Static         *ipv6UnicastStaticXml    `xml:"static,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ipv6UnicastConnectedXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastOspfv3Xml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastStaticXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Metric         *int64        `xml:"metric,omitempty"`
	RouteMap       *string       `xml:"route-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Ipv4 != nil {
		var obj ipv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj ipv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var ipv4Val *Ipv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &Entry{
		Name:           o.Name,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4Xml) MarshalFromObject(s Ipv4) {
	if s.Unicast != nil {
		var obj ipv4UnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4Xml) UnmarshalToObject() (*Ipv4, error) {
	var unicastVal *Ipv4Unicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &Ipv4{
		Unicast:        unicastVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastXml) MarshalFromObject(s Ipv4Unicast) {
	if s.Connected != nil {
		var obj ipv4UnicastConnectedXml
		obj.MarshalFromObject(*s.Connected)
		o.Connected = &obj
	}
	if s.Ospf != nil {
		var obj ipv4UnicastOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Rip != nil {
		var obj ipv4UnicastRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	if s.Static != nil {
		var obj ipv4UnicastStaticXml
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastXml) UnmarshalToObject() (*Ipv4Unicast, error) {
	var connectedVal *Ipv4UnicastConnected
	if o.Connected != nil {
		obj, err := o.Connected.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectedVal = obj
	}
	var ospfVal *Ipv4UnicastOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ripVal *Ipv4UnicastRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}
	var staticVal *Ipv4UnicastStatic
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}

	result := &Ipv4Unicast{
		Connected:      connectedVal,
		Ospf:           ospfVal,
		Rip:            ripVal,
		Static:         staticVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastConnectedXml) MarshalFromObject(s Ipv4UnicastConnected) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastConnectedXml) UnmarshalToObject() (*Ipv4UnicastConnected, error) {

	result := &Ipv4UnicastConnected{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastOspfXml) MarshalFromObject(s Ipv4UnicastOspf) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastOspfXml) UnmarshalToObject() (*Ipv4UnicastOspf, error) {

	result := &Ipv4UnicastOspf{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastRipXml) MarshalFromObject(s Ipv4UnicastRip) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastRipXml) UnmarshalToObject() (*Ipv4UnicastRip, error) {

	result := &Ipv4UnicastRip{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastStaticXml) MarshalFromObject(s Ipv4UnicastStatic) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastStaticXml) UnmarshalToObject() (*Ipv4UnicastStatic, error) {

	result := &Ipv4UnicastStatic{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6Xml) MarshalFromObject(s Ipv6) {
	if s.Unicast != nil {
		var obj ipv6UnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6Xml) UnmarshalToObject() (*Ipv6, error) {
	var unicastVal *Ipv6Unicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &Ipv6{
		Unicast:        unicastVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastXml) MarshalFromObject(s Ipv6Unicast) {
	if s.Connected != nil {
		var obj ipv6UnicastConnectedXml
		obj.MarshalFromObject(*s.Connected)
		o.Connected = &obj
	}
	if s.Ospfv3 != nil {
		var obj ipv6UnicastOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Static != nil {
		var obj ipv6UnicastStaticXml
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastXml) UnmarshalToObject() (*Ipv6Unicast, error) {
	var connectedVal *Ipv6UnicastConnected
	if o.Connected != nil {
		obj, err := o.Connected.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectedVal = obj
	}
	var ospfv3Val *Ipv6UnicastOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var staticVal *Ipv6UnicastStatic
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}

	result := &Ipv6Unicast{
		Connected:      connectedVal,
		Ospfv3:         ospfv3Val,
		Static:         staticVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastConnectedXml) MarshalFromObject(s Ipv6UnicastConnected) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastConnectedXml) UnmarshalToObject() (*Ipv6UnicastConnected, error) {

	result := &Ipv6UnicastConnected{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastOspfv3Xml) MarshalFromObject(s Ipv6UnicastOspfv3) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastOspfv3Xml) UnmarshalToObject() (*Ipv6UnicastOspfv3, error) {

	result := &Ipv6UnicastOspfv3{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
		RouteMap:       o.RouteMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastStaticXml) MarshalFromObject(s Ipv6UnicastStatic) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Metric = s.Metric
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastStaticXml) UnmarshalToObject() (*Ipv6UnicastStatic, error) {

	result := &Ipv6UnicastStatic{
		Enable:         util.AsBool(o.Enable, nil),
		Metric:         o.Metric,
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
	if v == "ipv4" || v == "Ipv4" {
		return e.Ipv4, nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
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
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *Ipv4) matches(other *Ipv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *Ipv4Unicast) matches(other *Ipv4Unicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Connected.matches(other.Connected) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}
	if !o.Static.matches(other.Static) {
		return false
	}

	return true
}

func (o *Ipv4UnicastConnected) matches(other *Ipv4UnicastConnected) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastOspf) matches(other *Ipv4UnicastOspf) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastRip) matches(other *Ipv4UnicastRip) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastStatic) matches(other *Ipv4UnicastStatic) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv6) matches(other *Ipv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *Ipv6Unicast) matches(other *Ipv6Unicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Connected.matches(other.Connected) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Static.matches(other.Static) {
		return false
	}

	return true
}

func (o *Ipv6UnicastConnected) matches(other *Ipv6UnicastConnected) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv6UnicastOspfv3) matches(other *Ipv6UnicastOspfv3) bool {
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
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *Ipv6UnicastStatic) matches(other *Ipv6UnicastStatic) bool {
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
