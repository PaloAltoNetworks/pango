package bgproutemap

import (
	"encoding/xml"
	"errors"
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
	suffix = []string{"network", "routing-profile", "filters", "route-maps", "bgp", "bgp-entry", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	RouteMap       []RouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *RouteMapMatch
	Set            *RouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatch struct {
	AsPathAccessList  *string
	RegularCommunity  *string
	LargeCommunity    *string
	ExtendedCommunity *string
	Interface         *string
	Origin            *string
	Metric            *int64
	Tag               *int64
	LocalPreference   *int64
	Peer              *string
	Ipv4              *RouteMapMatchIpv4
	Ipv6              *RouteMapMatchIpv6
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type RouteMapMatchIpv4 struct {
	Address        *RouteMapMatchIpv4Address
	NextHop        *RouteMapMatchIpv4NextHop
	RouteSource    *RouteMapMatchIpv4RouteSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv4RouteSource struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv6 struct {
	Address        *RouteMapMatchIpv6Address
	NextHop        *RouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapSet struct {
	Aggregator                *RouteMapSetAggregator
	Metric                    *RouteMapSetMetric
	Ipv4                      *RouteMapSetIpv4
	Ipv6                      *RouteMapSetIpv6
	Tag                       *int64
	LocalPreference           *int64
	Weight                    *int64
	Origin                    *string
	AtomicAggregate           *bool
	OriginatorId              *string
	Ipv6NexthopPreferGlobal   *bool
	OverwriteRegularCommunity *bool
	OverwriteLargeCommunity   *bool
	RemoveRegularCommunity    *string
	RemoveLargeCommunity      *string
	AspathPrepend             []int64
	RegularCommunity          []string
	LargeCommunity            []string
	AspathExclude             []int64
	ExtendedCommunity         []string
	Misc                      []generic.Xml
	MiscAttributes            []xml.Attr
}
type RouteMapSetAggregator struct {
	As             *int64
	RouterId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapSetIpv4 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RouteMapSetIpv6 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
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
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName        xml.Name              `xml:"entry"`
	Name           string                `xml:"name,attr"`
	Description    *string               `xml:"description,omitempty"`
	RouteMap       *routeMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
	MiscAttributes []xml.Attr            `xml:",any,attr"`
}
type routeMapContainerXml struct {
	Entries []routeMapXml `xml:"entry"`
}
type routeMapXml struct {
	XMLName        xml.Name          `xml:"entry"`
	Name           string            `xml:"name,attr"`
	Action         *string           `xml:"action,omitempty"`
	Description    *string           `xml:"description,omitempty"`
	Match          *routeMapMatchXml `xml:"match,omitempty"`
	Set            *routeMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml     `xml:",any"`
	MiscAttributes []xml.Attr        `xml:",any,attr"`
}
type routeMapMatchXml struct {
	AsPathAccessList  *string               `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string               `xml:"regular-community,omitempty"`
	LargeCommunity    *string               `xml:"large-community,omitempty"`
	ExtendedCommunity *string               `xml:"extended-community,omitempty"`
	Interface         *string               `xml:"interface,omitempty"`
	Origin            *string               `xml:"origin,omitempty"`
	Metric            *int64                `xml:"metric,omitempty"`
	Tag               *int64                `xml:"tag,omitempty"`
	LocalPreference   *int64                `xml:"local-preference,omitempty"`
	Peer              *string               `xml:"peer,omitempty"`
	Ipv4              *routeMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6              *routeMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc              []generic.Xml         `xml:",any"`
	MiscAttributes    []xml.Attr            `xml:",any,attr"`
}
type routeMapMatchIpv4Xml struct {
	Address        *routeMapMatchIpv4AddressXml     `xml:"address,omitempty"`
	NextHop        *routeMapMatchIpv4NextHopXml     `xml:"next-hop,omitempty"`
	RouteSource    *routeMapMatchIpv4RouteSourceXml `xml:"route-source,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type routeMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv4RouteSourceXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv6Xml struct {
	Address        *routeMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *routeMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type routeMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetXml struct {
	Aggregator                *routeMapSetAggregatorXml `xml:"aggregator,omitempty"`
	Metric                    *routeMapSetMetricXml     `xml:"metric,omitempty"`
	Ipv4                      *routeMapSetIpv4Xml       `xml:"ipv4,omitempty"`
	Ipv6                      *routeMapSetIpv6Xml       `xml:"ipv6,omitempty"`
	Tag                       *int64                    `xml:"tag,omitempty"`
	LocalPreference           *int64                    `xml:"local-preference,omitempty"`
	Weight                    *int64                    `xml:"weight,omitempty"`
	Origin                    *string                   `xml:"origin,omitempty"`
	AtomicAggregate           *string                   `xml:"atomic-aggregate,omitempty"`
	OriginatorId              *string                   `xml:"originator-id,omitempty"`
	Ipv6NexthopPreferGlobal   *string                   `xml:"ipv6-nexthop-prefer-global,omitempty"`
	OverwriteRegularCommunity *string                   `xml:"overwrite-regular-community,omitempty"`
	OverwriteLargeCommunity   *string                   `xml:"overwrite-large-community,omitempty"`
	RemoveRegularCommunity    *string                   `xml:"remove-regular-community,omitempty"`
	RemoveLargeCommunity      *string                   `xml:"remove-large-community,omitempty"`
	AspathPrepend             *util.MemberType          `xml:"aspath-prepend,omitempty"`
	RegularCommunity          *util.MemberType          `xml:"regular-community,omitempty"`
	LargeCommunity            *util.MemberType          `xml:"large-community,omitempty"`
	AspathExclude             *util.MemberType          `xml:"aspath-exclude,omitempty"`
	ExtendedCommunity         *util.MemberType          `xml:"extended-community,omitempty"`
	Misc                      []generic.Xml             `xml:",any"`
	MiscAttributes            []xml.Attr                `xml:",any,attr"`
}
type routeMapSetAggregatorXml struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetIpv4Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetIpv6Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type entryXml_11_0_2 struct {
	XMLName        xml.Name                     `xml:"entry"`
	Name           string                       `xml:"name,attr"`
	Description    *string                      `xml:"description,omitempty"`
	RouteMap       *routeMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type routeMapContainerXml_11_0_2 struct {
	Entries []routeMapXml_11_0_2 `xml:"entry"`
}
type routeMapXml_11_0_2 struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *routeMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *routeMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type routeMapMatchXml_11_0_2 struct {
	AsPathAccessList  *string                      `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                      `xml:"regular-community,omitempty"`
	LargeCommunity    *string                      `xml:"large-community,omitempty"`
	ExtendedCommunity *string                      `xml:"extended-community,omitempty"`
	Interface         *string                      `xml:"interface,omitempty"`
	Origin            *string                      `xml:"origin,omitempty"`
	Metric            *int64                       `xml:"metric,omitempty"`
	Tag               *int64                       `xml:"tag,omitempty"`
	LocalPreference   *int64                       `xml:"local-preference,omitempty"`
	Peer              *string                      `xml:"peer,omitempty"`
	Ipv4              *routeMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6              *routeMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc              []generic.Xml                `xml:",any"`
	MiscAttributes    []xml.Attr                   `xml:",any,attr"`
}
type routeMapMatchIpv4Xml_11_0_2 struct {
	Address        *routeMapMatchIpv4AddressXml_11_0_2     `xml:"address,omitempty"`
	NextHop        *routeMapMatchIpv4NextHopXml_11_0_2     `xml:"next-hop,omitempty"`
	RouteSource    *routeMapMatchIpv4RouteSourceXml_11_0_2 `xml:"route-source,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type routeMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv4RouteSourceXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv6Xml_11_0_2 struct {
	Address        *routeMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *routeMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type routeMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetXml_11_0_2 struct {
	Aggregator                *routeMapSetAggregatorXml_11_0_2 `xml:"aggregator,omitempty"`
	Metric                    *routeMapSetMetricXml_11_0_2     `xml:"metric,omitempty"`
	Ipv4                      *routeMapSetIpv4Xml_11_0_2       `xml:"ipv4,omitempty"`
	Ipv6                      *routeMapSetIpv6Xml_11_0_2       `xml:"ipv6,omitempty"`
	Tag                       *int64                           `xml:"tag,omitempty"`
	LocalPreference           *int64                           `xml:"local-preference,omitempty"`
	Weight                    *int64                           `xml:"weight,omitempty"`
	Origin                    *string                          `xml:"origin,omitempty"`
	AtomicAggregate           *string                          `xml:"atomic-aggregate,omitempty"`
	OriginatorId              *string                          `xml:"originator-id,omitempty"`
	Ipv6NexthopPreferGlobal   *string                          `xml:"ipv6-nexthop-prefer-global,omitempty"`
	OverwriteRegularCommunity *string                          `xml:"overwrite-regular-community,omitempty"`
	OverwriteLargeCommunity   *string                          `xml:"overwrite-large-community,omitempty"`
	RemoveRegularCommunity    *string                          `xml:"remove-regular-community,omitempty"`
	RemoveLargeCommunity      *string                          `xml:"remove-large-community,omitempty"`
	AspathPrepend             *util.MemberType                 `xml:"aspath-prepend,omitempty"`
	RegularCommunity          *util.MemberType                 `xml:"regular-community,omitempty"`
	LargeCommunity            *util.MemberType                 `xml:"large-community,omitempty"`
	AspathExclude             *util.MemberType                 `xml:"aspath-exclude,omitempty"`
	ExtendedCommunity         *util.MemberType                 `xml:"extended-community,omitempty"`
	Misc                      []generic.Xml                    `xml:",any"`
	MiscAttributes            []xml.Attr                       `xml:",any,attr"`
}
type routeMapSetAggregatorXml_11_0_2 struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetIpv4Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type routeMapSetIpv6Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.RouteMap != nil {
		var objs []routeMapXml
		for _, elt := range s.RouteMap {
			var obj routeMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &routeMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var routeMapVal []RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapXml) MarshalFromObject(s RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj routeMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj routeMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapXml) UnmarshalToObject() (*RouteMap, error) {
	var matchVal *RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RouteMap{
		Name:           o.Name,
		Action:         o.Action,
		Description:    o.Description,
		Match:          matchVal,
		Set:            setVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchXml) MarshalFromObject(s RouteMapMatch) {
	o.AsPathAccessList = s.AsPathAccessList
	o.RegularCommunity = s.RegularCommunity
	o.LargeCommunity = s.LargeCommunity
	o.ExtendedCommunity = s.ExtendedCommunity
	o.Interface = s.Interface
	o.Origin = s.Origin
	o.Metric = s.Metric
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Peer = s.Peer
	if s.Ipv4 != nil {
		var obj routeMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj routeMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchXml) UnmarshalToObject() (*RouteMapMatch, error) {
	var ipv4Val *RouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &RouteMapMatch{
		AsPathAccessList:  o.AsPathAccessList,
		RegularCommunity:  o.RegularCommunity,
		LargeCommunity:    o.LargeCommunity,
		ExtendedCommunity: o.ExtendedCommunity,
		Interface:         o.Interface,
		Origin:            o.Origin,
		Metric:            o.Metric,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Peer:              o.Peer,
		Ipv4:              ipv4Val,
		Ipv6:              ipv6Val,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4Xml) MarshalFromObject(s RouteMapMatchIpv4) {
	if s.Address != nil {
		var obj routeMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj routeMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj routeMapMatchIpv4RouteSourceXml
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4Xml) UnmarshalToObject() (*RouteMapMatchIpv4, error) {
	var addressVal *RouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *RouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &RouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4AddressXml) MarshalFromObject(s RouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4AddressXml) UnmarshalToObject() (*RouteMapMatchIpv4Address, error) {

	result := &RouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4NextHopXml) MarshalFromObject(s RouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4NextHopXml) UnmarshalToObject() (*RouteMapMatchIpv4NextHop, error) {

	result := &RouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4RouteSourceXml) MarshalFromObject(s RouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4RouteSourceXml) UnmarshalToObject() (*RouteMapMatchIpv4RouteSource, error) {

	result := &RouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6Xml) MarshalFromObject(s RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj routeMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj routeMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6Xml) UnmarshalToObject() (*RouteMapMatchIpv6, error) {
	var addressVal *RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6AddressXml) MarshalFromObject(s RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6AddressXml) UnmarshalToObject() (*RouteMapMatchIpv6Address, error) {

	result := &RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6NextHopXml) MarshalFromObject(s RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6NextHopXml) UnmarshalToObject() (*RouteMapMatchIpv6NextHop, error) {

	result := &RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetXml) MarshalFromObject(s RouteMapSet) {
	if s.Aggregator != nil {
		var obj routeMapSetAggregatorXml
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj routeMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj routeMapSetIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj routeMapSetIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	o.Ipv6NexthopPreferGlobal = util.YesNo(s.Ipv6NexthopPreferGlobal, nil)
	o.OverwriteRegularCommunity = util.YesNo(s.OverwriteRegularCommunity, nil)
	o.OverwriteLargeCommunity = util.YesNo(s.OverwriteLargeCommunity, nil)
	o.RemoveRegularCommunity = s.RemoveRegularCommunity
	o.RemoveLargeCommunity = s.RemoveLargeCommunity
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.AspathExclude != nil {
		o.AspathExclude = util.Int64ToMem(s.AspathExclude)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetXml) UnmarshalToObject() (*RouteMapSet, error) {
	var aggregatorVal *RouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *RouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *RouteMapSetIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var aspathPrependVal []int64
	if o.AspathPrepend != nil {
		var err error
		aspathPrependVal, err = util.MemToInt64(o.AspathPrepend)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}
	var regularCommunityVal []string
	if o.RegularCommunity != nil {
		regularCommunityVal = util.MemToStr(o.RegularCommunity)
	}
	var largeCommunityVal []string
	if o.LargeCommunity != nil {
		largeCommunityVal = util.MemToStr(o.LargeCommunity)
	}
	var aspathExcludeVal []int64
	if o.AspathExclude != nil {
		var err error
		aspathExcludeVal, err = util.MemToInt64(o.AspathExclude)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &RouteMapSet{
		Aggregator:                aggregatorVal,
		Metric:                    metricVal,
		Ipv4:                      ipv4Val,
		Ipv6:                      ipv6Val,
		Tag:                       o.Tag,
		LocalPreference:           o.LocalPreference,
		Weight:                    o.Weight,
		Origin:                    o.Origin,
		AtomicAggregate:           util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:              o.OriginatorId,
		Ipv6NexthopPreferGlobal:   util.AsBool(o.Ipv6NexthopPreferGlobal, nil),
		OverwriteRegularCommunity: util.AsBool(o.OverwriteRegularCommunity, nil),
		OverwriteLargeCommunity:   util.AsBool(o.OverwriteLargeCommunity, nil),
		RemoveRegularCommunity:    o.RemoveRegularCommunity,
		RemoveLargeCommunity:      o.RemoveLargeCommunity,
		AspathPrepend:             aspathPrependVal,
		RegularCommunity:          regularCommunityVal,
		LargeCommunity:            largeCommunityVal,
		AspathExclude:             aspathExcludeVal,
		ExtendedCommunity:         extendedCommunityVal,
		Misc:                      o.Misc,
		MiscAttributes:            o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetAggregatorXml) MarshalFromObject(s RouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetAggregatorXml) UnmarshalToObject() (*RouteMapSetAggregator, error) {

	result := &RouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetMetricXml) MarshalFromObject(s RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetMetricXml) UnmarshalToObject() (*RouteMapSetMetric, error) {

	result := &RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetIpv4Xml) MarshalFromObject(s RouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetIpv4Xml) UnmarshalToObject() (*RouteMapSetIpv4, error) {

	result := &RouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetIpv6Xml) MarshalFromObject(s RouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetIpv6Xml) UnmarshalToObject() (*RouteMapSetIpv6, error) {

	result := &RouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.RouteMap != nil {
		var objs []routeMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj routeMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &routeMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var routeMapVal []RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapXml_11_0_2) MarshalFromObject(s RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj routeMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj routeMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapXml_11_0_2) UnmarshalToObject() (*RouteMap, error) {
	var matchVal *RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RouteMap{
		Name:           o.Name,
		Action:         o.Action,
		Description:    o.Description,
		Match:          matchVal,
		Set:            setVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchXml_11_0_2) MarshalFromObject(s RouteMapMatch) {
	o.AsPathAccessList = s.AsPathAccessList
	o.RegularCommunity = s.RegularCommunity
	o.LargeCommunity = s.LargeCommunity
	o.ExtendedCommunity = s.ExtendedCommunity
	o.Interface = s.Interface
	o.Origin = s.Origin
	o.Metric = s.Metric
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Peer = s.Peer
	if s.Ipv4 != nil {
		var obj routeMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj routeMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchXml_11_0_2) UnmarshalToObject() (*RouteMapMatch, error) {
	var ipv4Val *RouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &RouteMapMatch{
		AsPathAccessList:  o.AsPathAccessList,
		RegularCommunity:  o.RegularCommunity,
		LargeCommunity:    o.LargeCommunity,
		ExtendedCommunity: o.ExtendedCommunity,
		Interface:         o.Interface,
		Origin:            o.Origin,
		Metric:            o.Metric,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Peer:              o.Peer,
		Ipv4:              ipv4Val,
		Ipv6:              ipv6Val,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4Xml_11_0_2) MarshalFromObject(s RouteMapMatchIpv4) {
	if s.Address != nil {
		var obj routeMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj routeMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj routeMapMatchIpv4RouteSourceXml_11_0_2
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv4, error) {
	var addressVal *RouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *RouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &RouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s RouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv4Address, error) {

	result := &RouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s RouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv4NextHop, error) {

	result := &RouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv4RouteSourceXml_11_0_2) MarshalFromObject(s RouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv4RouteSourceXml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv4RouteSource, error) {

	result := &RouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6Xml_11_0_2) MarshalFromObject(s RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj routeMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj routeMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv6, error) {
	var addressVal *RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv6Address, error) {

	result := &RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*RouteMapMatchIpv6NextHop, error) {

	result := &RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetXml_11_0_2) MarshalFromObject(s RouteMapSet) {
	if s.Aggregator != nil {
		var obj routeMapSetAggregatorXml_11_0_2
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj routeMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj routeMapSetIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj routeMapSetIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	o.Ipv6NexthopPreferGlobal = util.YesNo(s.Ipv6NexthopPreferGlobal, nil)
	o.OverwriteRegularCommunity = util.YesNo(s.OverwriteRegularCommunity, nil)
	o.OverwriteLargeCommunity = util.YesNo(s.OverwriteLargeCommunity, nil)
	o.RemoveRegularCommunity = s.RemoveRegularCommunity
	o.RemoveLargeCommunity = s.RemoveLargeCommunity
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.AspathExclude != nil {
		o.AspathExclude = util.Int64ToMem(s.AspathExclude)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetXml_11_0_2) UnmarshalToObject() (*RouteMapSet, error) {
	var aggregatorVal *RouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *RouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *RouteMapSetIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var aspathPrependVal []int64
	if o.AspathPrepend != nil {
		var err error
		aspathPrependVal, err = util.MemToInt64(o.AspathPrepend)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}
	var regularCommunityVal []string
	if o.RegularCommunity != nil {
		regularCommunityVal = util.MemToStr(o.RegularCommunity)
	}
	var largeCommunityVal []string
	if o.LargeCommunity != nil {
		largeCommunityVal = util.MemToStr(o.LargeCommunity)
	}
	var aspathExcludeVal []int64
	if o.AspathExclude != nil {
		var err error
		aspathExcludeVal, err = util.MemToInt64(o.AspathExclude)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &RouteMapSet{
		Aggregator:                aggregatorVal,
		Metric:                    metricVal,
		Ipv4:                      ipv4Val,
		Ipv6:                      ipv6Val,
		Tag:                       o.Tag,
		LocalPreference:           o.LocalPreference,
		Weight:                    o.Weight,
		Origin:                    o.Origin,
		AtomicAggregate:           util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:              o.OriginatorId,
		Ipv6NexthopPreferGlobal:   util.AsBool(o.Ipv6NexthopPreferGlobal, nil),
		OverwriteRegularCommunity: util.AsBool(o.OverwriteRegularCommunity, nil),
		OverwriteLargeCommunity:   util.AsBool(o.OverwriteLargeCommunity, nil),
		RemoveRegularCommunity:    o.RemoveRegularCommunity,
		RemoveLargeCommunity:      o.RemoveLargeCommunity,
		AspathPrepend:             aspathPrependVal,
		RegularCommunity:          regularCommunityVal,
		LargeCommunity:            largeCommunityVal,
		AspathExclude:             aspathExcludeVal,
		ExtendedCommunity:         extendedCommunityVal,
		Misc:                      o.Misc,
		MiscAttributes:            o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetAggregatorXml_11_0_2) MarshalFromObject(s RouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetAggregatorXml_11_0_2) UnmarshalToObject() (*RouteMapSetAggregator, error) {

	result := &RouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetMetricXml_11_0_2) MarshalFromObject(s RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetMetricXml_11_0_2) UnmarshalToObject() (*RouteMapSetMetric, error) {

	result := &RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetIpv4Xml_11_0_2) MarshalFromObject(s RouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetIpv4Xml_11_0_2) UnmarshalToObject() (*RouteMapSetIpv4, error) {

	result := &RouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *routeMapSetIpv6Xml_11_0_2) MarshalFromObject(s RouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o routeMapSetIpv6Xml_11_0_2) UnmarshalToObject() (*RouteMapSetIpv6, error) {

	result := &RouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
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
	if v == "route_map" || v == "RouteMap" {
		return e.RouteMap, nil
	}
	if v == "route_map|LENGTH" || v == "RouteMap|LENGTH" {
		return int64(len(e.RouteMap)), nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

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
	if len(o.RouteMap) != len(other.RouteMap) {
		return false
	}
	for idx := range o.RouteMap {
		if !o.RouteMap[idx].matches(&other.RouteMap[idx]) {
			return false
		}
	}

	return true
}

func (o *RouteMap) matches(other *RouteMap) bool {
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !o.Match.matches(other.Match) {
		return false
	}
	if !o.Set.matches(other.Set) {
		return false
	}

	return true
}

func (o *RouteMapMatch) matches(other *RouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AsPathAccessList, other.AsPathAccessList) {
		return false
	}
	if !util.StringsMatch(o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.StringsMatch(o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.StringsMatch(o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Origin, other.Origin) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}
	if !util.Ints64Match(o.LocalPreference, other.LocalPreference) {
		return false
	}
	if !util.StringsMatch(o.Peer, other.Peer) {
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

func (o *RouteMapMatchIpv4) matches(other *RouteMapMatchIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Address.matches(other.Address) {
		return false
	}
	if !o.NextHop.matches(other.NextHop) {
		return false
	}
	if !o.RouteSource.matches(other.RouteSource) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv4Address) matches(other *RouteMapMatchIpv4Address) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv4NextHop) matches(other *RouteMapMatchIpv4NextHop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv4RouteSource) matches(other *RouteMapMatchIpv4RouteSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv6) matches(other *RouteMapMatchIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Address.matches(other.Address) {
		return false
	}
	if !o.NextHop.matches(other.NextHop) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv6Address) matches(other *RouteMapMatchIpv6Address) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *RouteMapMatchIpv6NextHop) matches(other *RouteMapMatchIpv6NextHop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *RouteMapSet) matches(other *RouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Aggregator.matches(other.Aggregator) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}
	if !util.Ints64Match(o.LocalPreference, other.LocalPreference) {
		return false
	}
	if !util.Ints64Match(o.Weight, other.Weight) {
		return false
	}
	if !util.StringsMatch(o.Origin, other.Origin) {
		return false
	}
	if !util.BoolsMatch(o.AtomicAggregate, other.AtomicAggregate) {
		return false
	}
	if !util.StringsMatch(o.OriginatorId, other.OriginatorId) {
		return false
	}
	if !util.BoolsMatch(o.Ipv6NexthopPreferGlobal, other.Ipv6NexthopPreferGlobal) {
		return false
	}
	if !util.BoolsMatch(o.OverwriteRegularCommunity, other.OverwriteRegularCommunity) {
		return false
	}
	if !util.BoolsMatch(o.OverwriteLargeCommunity, other.OverwriteLargeCommunity) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegularCommunity, other.RemoveRegularCommunity) {
		return false
	}
	if !util.StringsMatch(o.RemoveLargeCommunity, other.RemoveLargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[int64](o.AspathPrepend, other.AspathPrepend) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[int64](o.AspathExclude, other.AspathExclude) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *RouteMapSetAggregator) matches(other *RouteMapSetAggregator) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.As, other.As) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}

	return true
}

func (o *RouteMapSetMetric) matches(other *RouteMapSetMetric) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}

	return true
}

func (o *RouteMapSetIpv4) matches(other *RouteMapSetIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}
	if !util.StringsMatch(o.NextHop, other.NextHop) {
		return false
	}

	return true
}

func (o *RouteMapSetIpv6) matches(other *RouteMapSetIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}
	if !util.StringsMatch(o.NextHop, other.NextHop) {
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
