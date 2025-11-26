package filtering

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
	suffix = []string{"network", "routing-profile", "bgp", "filtering-profile", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Ipv4           *Ipv4
	Ipv6           *Ipv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4 struct {
	Multicast      *Ipv4Multicast
	Unicast        *Ipv4Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4Multicast struct {
	ConditionalAdvertisement *Ipv4MulticastConditionalAdvertisement
	FilterList               *Ipv4MulticastFilterList
	InboundNetworkFilters    *Ipv4MulticastInboundNetworkFilters
	Inherit                  *bool
	OutboundNetworkFilters   *Ipv4MulticastOutboundNetworkFilters
	RouteMaps                *Ipv4MulticastRouteMaps
	UnsuppressMap            *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type Ipv4MulticastConditionalAdvertisement struct {
	Exist          *Ipv4MulticastConditionalAdvertisementExist
	NonExist       *Ipv4MulticastConditionalAdvertisementNonExist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastConditionalAdvertisementExist struct {
	AdvertiseMap   *string
	ExistMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastConditionalAdvertisementNonExist struct {
	AdvertiseMap   *string
	NonExistMap    *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastFilterList struct {
	Inbound        *string
	Outbound       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastInboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastOutboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastRouteMaps struct {
	Inbound        *string
	Outbound       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4Unicast struct {
	ConditionalAdvertisement *Ipv4UnicastConditionalAdvertisement
	FilterList               *Ipv4UnicastFilterList
	InboundNetworkFilters    *Ipv4UnicastInboundNetworkFilters
	OutboundNetworkFilters   *Ipv4UnicastOutboundNetworkFilters
	RouteMaps                *Ipv4UnicastRouteMaps
	UnsuppressMap            *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type Ipv4UnicastConditionalAdvertisement struct {
	Exist          *Ipv4UnicastConditionalAdvertisementExist
	NonExist       *Ipv4UnicastConditionalAdvertisementNonExist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastConditionalAdvertisementExist struct {
	AdvertiseMap   *string
	ExistMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastConditionalAdvertisementNonExist struct {
	AdvertiseMap   *string
	NonExistMap    *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastFilterList struct {
	Inbound        *string
	Outbound       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastInboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastOutboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastRouteMaps struct {
	Inbound        *string
	Outbound       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6 struct {
	Unicast        *Ipv6Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6Unicast struct {
	ConditionalAdvertisement *Ipv6UnicastConditionalAdvertisement
	FilterList               *Ipv6UnicastFilterList
	InboundNetworkFilters    *Ipv6UnicastInboundNetworkFilters
	OutboundNetworkFilters   *Ipv6UnicastOutboundNetworkFilters
	RouteMaps                *Ipv6UnicastRouteMaps
	UnsuppressMap            *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type Ipv6UnicastConditionalAdvertisement struct {
	Exist          *Ipv6UnicastConditionalAdvertisementExist
	NonExist       *Ipv6UnicastConditionalAdvertisementNonExist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastConditionalAdvertisementExist struct {
	AdvertiseMap   *string
	ExistMap       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastConditionalAdvertisementNonExist struct {
	AdvertiseMap   *string
	NonExistMap    *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastFilterList struct {
	Inbound        *string
	Outbound       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastInboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastOutboundNetworkFilters struct {
	DistributeList *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastRouteMaps struct {
	Inbound        *string
	Outbound       *string
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
	Description    *string       `xml:"description,omitempty"`
	Ipv4           *ipv4Xml      `xml:"ipv4,omitempty"`
	Ipv6           *ipv6Xml      `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4Xml struct {
	Multicast      *ipv4MulticastXml `xml:"multicast,omitempty"`
	Unicast        *ipv4UnicastXml   `xml:"unicast,omitempty"`
	Misc           []generic.Xml     `xml:",any"`
	MiscAttributes []xml.Attr        `xml:",any,attr"`
}
type ipv4MulticastXml struct {
	ConditionalAdvertisement *ipv4MulticastConditionalAdvertisementXml `xml:"conditional-advertisement,omitempty"`
	FilterList               *ipv4MulticastFilterListXml               `xml:"filter-list,omitempty"`
	InboundNetworkFilters    *ipv4MulticastInboundNetworkFiltersXml    `xml:"inbound-network-filters,omitempty"`
	Inherit                  *string                                   `xml:"inherit,omitempty"`
	OutboundNetworkFilters   *ipv4MulticastOutboundNetworkFiltersXml   `xml:"outbound-network-filters,omitempty"`
	RouteMaps                *ipv4MulticastRouteMapsXml                `xml:"route-maps,omitempty"`
	UnsuppressMap            *string                                   `xml:"unsuppress-map,omitempty"`
	Misc                     []generic.Xml                             `xml:",any"`
	MiscAttributes           []xml.Attr                                `xml:",any,attr"`
}
type ipv4MulticastConditionalAdvertisementXml struct {
	Exist          *ipv4MulticastConditionalAdvertisementExistXml    `xml:"exist,omitempty"`
	NonExist       *ipv4MulticastConditionalAdvertisementNonExistXml `xml:"non-exist,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type ipv4MulticastConditionalAdvertisementExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	ExistMap       *string       `xml:"exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastConditionalAdvertisementNonExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	NonExistMap    *string       `xml:"non-exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastFilterListXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastInboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastOutboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastRouteMapsXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastXml struct {
	ConditionalAdvertisement *ipv4UnicastConditionalAdvertisementXml `xml:"conditional-advertisement,omitempty"`
	FilterList               *ipv4UnicastFilterListXml               `xml:"filter-list,omitempty"`
	InboundNetworkFilters    *ipv4UnicastInboundNetworkFiltersXml    `xml:"inbound-network-filters,omitempty"`
	OutboundNetworkFilters   *ipv4UnicastOutboundNetworkFiltersXml   `xml:"outbound-network-filters,omitempty"`
	RouteMaps                *ipv4UnicastRouteMapsXml                `xml:"route-maps,omitempty"`
	UnsuppressMap            *string                                 `xml:"unsuppress-map,omitempty"`
	Misc                     []generic.Xml                           `xml:",any"`
	MiscAttributes           []xml.Attr                              `xml:",any,attr"`
}
type ipv4UnicastConditionalAdvertisementXml struct {
	Exist          *ipv4UnicastConditionalAdvertisementExistXml    `xml:"exist,omitempty"`
	NonExist       *ipv4UnicastConditionalAdvertisementNonExistXml `xml:"non-exist,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type ipv4UnicastConditionalAdvertisementExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	ExistMap       *string       `xml:"exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastConditionalAdvertisementNonExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	NonExistMap    *string       `xml:"non-exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastFilterListXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastInboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastOutboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastRouteMapsXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml struct {
	Unicast        *ipv6UnicastXml `xml:"unicast,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type ipv6UnicastXml struct {
	ConditionalAdvertisement *ipv6UnicastConditionalAdvertisementXml `xml:"conditional-advertisement,omitempty"`
	FilterList               *ipv6UnicastFilterListXml               `xml:"filter-list,omitempty"`
	InboundNetworkFilters    *ipv6UnicastInboundNetworkFiltersXml    `xml:"inbound-network-filters,omitempty"`
	OutboundNetworkFilters   *ipv6UnicastOutboundNetworkFiltersXml   `xml:"outbound-network-filters,omitempty"`
	RouteMaps                *ipv6UnicastRouteMapsXml                `xml:"route-maps,omitempty"`
	UnsuppressMap            *string                                 `xml:"unsuppress-map,omitempty"`
	Misc                     []generic.Xml                           `xml:",any"`
	MiscAttributes           []xml.Attr                              `xml:",any,attr"`
}
type ipv6UnicastConditionalAdvertisementXml struct {
	Exist          *ipv6UnicastConditionalAdvertisementExistXml    `xml:"exist,omitempty"`
	NonExist       *ipv6UnicastConditionalAdvertisementNonExistXml `xml:"non-exist,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type ipv6UnicastConditionalAdvertisementExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	ExistMap       *string       `xml:"exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastConditionalAdvertisementNonExistXml struct {
	AdvertiseMap   *string       `xml:"advertise-map,omitempty"`
	NonExistMap    *string       `xml:"non-exist-map,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastFilterListXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastInboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastOutboundNetworkFiltersXml struct {
	DistributeList *string       `xml:"distribute-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastRouteMapsXml struct {
	Inbound        *string       `xml:"inbound,omitempty"`
	Outbound       *string       `xml:"outbound,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
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
		Description:    o.Description,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4Xml) MarshalFromObject(s Ipv4) {
	if s.Multicast != nil {
		var obj ipv4MulticastXml
		obj.MarshalFromObject(*s.Multicast)
		o.Multicast = &obj
	}
	if s.Unicast != nil {
		var obj ipv4UnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4Xml) UnmarshalToObject() (*Ipv4, error) {
	var multicastVal *Ipv4Multicast
	if o.Multicast != nil {
		obj, err := o.Multicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastVal = obj
	}
	var unicastVal *Ipv4Unicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &Ipv4{
		Multicast:      multicastVal,
		Unicast:        unicastVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastXml) MarshalFromObject(s Ipv4Multicast) {
	if s.ConditionalAdvertisement != nil {
		var obj ipv4MulticastConditionalAdvertisementXml
		obj.MarshalFromObject(*s.ConditionalAdvertisement)
		o.ConditionalAdvertisement = &obj
	}
	if s.FilterList != nil {
		var obj ipv4MulticastFilterListXml
		obj.MarshalFromObject(*s.FilterList)
		o.FilterList = &obj
	}
	if s.InboundNetworkFilters != nil {
		var obj ipv4MulticastInboundNetworkFiltersXml
		obj.MarshalFromObject(*s.InboundNetworkFilters)
		o.InboundNetworkFilters = &obj
	}
	o.Inherit = util.YesNo(s.Inherit, nil)
	if s.OutboundNetworkFilters != nil {
		var obj ipv4MulticastOutboundNetworkFiltersXml
		obj.MarshalFromObject(*s.OutboundNetworkFilters)
		o.OutboundNetworkFilters = &obj
	}
	if s.RouteMaps != nil {
		var obj ipv4MulticastRouteMapsXml
		obj.MarshalFromObject(*s.RouteMaps)
		o.RouteMaps = &obj
	}
	o.UnsuppressMap = s.UnsuppressMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastXml) UnmarshalToObject() (*Ipv4Multicast, error) {
	var conditionalAdvertisementVal *Ipv4MulticastConditionalAdvertisement
	if o.ConditionalAdvertisement != nil {
		obj, err := o.ConditionalAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		conditionalAdvertisementVal = obj
	}
	var filterListVal *Ipv4MulticastFilterList
	if o.FilterList != nil {
		obj, err := o.FilterList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filterListVal = obj
	}
	var inboundNetworkFiltersVal *Ipv4MulticastInboundNetworkFilters
	if o.InboundNetworkFilters != nil {
		obj, err := o.InboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inboundNetworkFiltersVal = obj
	}
	var outboundNetworkFiltersVal *Ipv4MulticastOutboundNetworkFilters
	if o.OutboundNetworkFilters != nil {
		obj, err := o.OutboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		outboundNetworkFiltersVal = obj
	}
	var routeMapsVal *Ipv4MulticastRouteMaps
	if o.RouteMaps != nil {
		obj, err := o.RouteMaps.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeMapsVal = obj
	}

	result := &Ipv4Multicast{
		ConditionalAdvertisement: conditionalAdvertisementVal,
		FilterList:               filterListVal,
		InboundNetworkFilters:    inboundNetworkFiltersVal,
		Inherit:                  util.AsBool(o.Inherit, nil),
		OutboundNetworkFilters:   outboundNetworkFiltersVal,
		RouteMaps:                routeMapsVal,
		UnsuppressMap:            o.UnsuppressMap,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastConditionalAdvertisementXml) MarshalFromObject(s Ipv4MulticastConditionalAdvertisement) {
	if s.Exist != nil {
		var obj ipv4MulticastConditionalAdvertisementExistXml
		obj.MarshalFromObject(*s.Exist)
		o.Exist = &obj
	}
	if s.NonExist != nil {
		var obj ipv4MulticastConditionalAdvertisementNonExistXml
		obj.MarshalFromObject(*s.NonExist)
		o.NonExist = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastConditionalAdvertisementXml) UnmarshalToObject() (*Ipv4MulticastConditionalAdvertisement, error) {
	var existVal *Ipv4MulticastConditionalAdvertisementExist
	if o.Exist != nil {
		obj, err := o.Exist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		existVal = obj
	}
	var nonExistVal *Ipv4MulticastConditionalAdvertisementNonExist
	if o.NonExist != nil {
		obj, err := o.NonExist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nonExistVal = obj
	}

	result := &Ipv4MulticastConditionalAdvertisement{
		Exist:          existVal,
		NonExist:       nonExistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastConditionalAdvertisementExistXml) MarshalFromObject(s Ipv4MulticastConditionalAdvertisementExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.ExistMap = s.ExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastConditionalAdvertisementExistXml) UnmarshalToObject() (*Ipv4MulticastConditionalAdvertisementExist, error) {

	result := &Ipv4MulticastConditionalAdvertisementExist{
		AdvertiseMap:   o.AdvertiseMap,
		ExistMap:       o.ExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastConditionalAdvertisementNonExistXml) MarshalFromObject(s Ipv4MulticastConditionalAdvertisementNonExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.NonExistMap = s.NonExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastConditionalAdvertisementNonExistXml) UnmarshalToObject() (*Ipv4MulticastConditionalAdvertisementNonExist, error) {

	result := &Ipv4MulticastConditionalAdvertisementNonExist{
		AdvertiseMap:   o.AdvertiseMap,
		NonExistMap:    o.NonExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastFilterListXml) MarshalFromObject(s Ipv4MulticastFilterList) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastFilterListXml) UnmarshalToObject() (*Ipv4MulticastFilterList, error) {

	result := &Ipv4MulticastFilterList{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastInboundNetworkFiltersXml) MarshalFromObject(s Ipv4MulticastInboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastInboundNetworkFiltersXml) UnmarshalToObject() (*Ipv4MulticastInboundNetworkFilters, error) {

	result := &Ipv4MulticastInboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastOutboundNetworkFiltersXml) MarshalFromObject(s Ipv4MulticastOutboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastOutboundNetworkFiltersXml) UnmarshalToObject() (*Ipv4MulticastOutboundNetworkFilters, error) {

	result := &Ipv4MulticastOutboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastRouteMapsXml) MarshalFromObject(s Ipv4MulticastRouteMaps) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastRouteMapsXml) UnmarshalToObject() (*Ipv4MulticastRouteMaps, error) {

	result := &Ipv4MulticastRouteMaps{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastXml) MarshalFromObject(s Ipv4Unicast) {
	if s.ConditionalAdvertisement != nil {
		var obj ipv4UnicastConditionalAdvertisementXml
		obj.MarshalFromObject(*s.ConditionalAdvertisement)
		o.ConditionalAdvertisement = &obj
	}
	if s.FilterList != nil {
		var obj ipv4UnicastFilterListXml
		obj.MarshalFromObject(*s.FilterList)
		o.FilterList = &obj
	}
	if s.InboundNetworkFilters != nil {
		var obj ipv4UnicastInboundNetworkFiltersXml
		obj.MarshalFromObject(*s.InboundNetworkFilters)
		o.InboundNetworkFilters = &obj
	}
	if s.OutboundNetworkFilters != nil {
		var obj ipv4UnicastOutboundNetworkFiltersXml
		obj.MarshalFromObject(*s.OutboundNetworkFilters)
		o.OutboundNetworkFilters = &obj
	}
	if s.RouteMaps != nil {
		var obj ipv4UnicastRouteMapsXml
		obj.MarshalFromObject(*s.RouteMaps)
		o.RouteMaps = &obj
	}
	o.UnsuppressMap = s.UnsuppressMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastXml) UnmarshalToObject() (*Ipv4Unicast, error) {
	var conditionalAdvertisementVal *Ipv4UnicastConditionalAdvertisement
	if o.ConditionalAdvertisement != nil {
		obj, err := o.ConditionalAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		conditionalAdvertisementVal = obj
	}
	var filterListVal *Ipv4UnicastFilterList
	if o.FilterList != nil {
		obj, err := o.FilterList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filterListVal = obj
	}
	var inboundNetworkFiltersVal *Ipv4UnicastInboundNetworkFilters
	if o.InboundNetworkFilters != nil {
		obj, err := o.InboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inboundNetworkFiltersVal = obj
	}
	var outboundNetworkFiltersVal *Ipv4UnicastOutboundNetworkFilters
	if o.OutboundNetworkFilters != nil {
		obj, err := o.OutboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		outboundNetworkFiltersVal = obj
	}
	var routeMapsVal *Ipv4UnicastRouteMaps
	if o.RouteMaps != nil {
		obj, err := o.RouteMaps.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeMapsVal = obj
	}

	result := &Ipv4Unicast{
		ConditionalAdvertisement: conditionalAdvertisementVal,
		FilterList:               filterListVal,
		InboundNetworkFilters:    inboundNetworkFiltersVal,
		OutboundNetworkFilters:   outboundNetworkFiltersVal,
		RouteMaps:                routeMapsVal,
		UnsuppressMap:            o.UnsuppressMap,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastConditionalAdvertisementXml) MarshalFromObject(s Ipv4UnicastConditionalAdvertisement) {
	if s.Exist != nil {
		var obj ipv4UnicastConditionalAdvertisementExistXml
		obj.MarshalFromObject(*s.Exist)
		o.Exist = &obj
	}
	if s.NonExist != nil {
		var obj ipv4UnicastConditionalAdvertisementNonExistXml
		obj.MarshalFromObject(*s.NonExist)
		o.NonExist = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastConditionalAdvertisementXml) UnmarshalToObject() (*Ipv4UnicastConditionalAdvertisement, error) {
	var existVal *Ipv4UnicastConditionalAdvertisementExist
	if o.Exist != nil {
		obj, err := o.Exist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		existVal = obj
	}
	var nonExistVal *Ipv4UnicastConditionalAdvertisementNonExist
	if o.NonExist != nil {
		obj, err := o.NonExist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nonExistVal = obj
	}

	result := &Ipv4UnicastConditionalAdvertisement{
		Exist:          existVal,
		NonExist:       nonExistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastConditionalAdvertisementExistXml) MarshalFromObject(s Ipv4UnicastConditionalAdvertisementExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.ExistMap = s.ExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastConditionalAdvertisementExistXml) UnmarshalToObject() (*Ipv4UnicastConditionalAdvertisementExist, error) {

	result := &Ipv4UnicastConditionalAdvertisementExist{
		AdvertiseMap:   o.AdvertiseMap,
		ExistMap:       o.ExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastConditionalAdvertisementNonExistXml) MarshalFromObject(s Ipv4UnicastConditionalAdvertisementNonExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.NonExistMap = s.NonExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastConditionalAdvertisementNonExistXml) UnmarshalToObject() (*Ipv4UnicastConditionalAdvertisementNonExist, error) {

	result := &Ipv4UnicastConditionalAdvertisementNonExist{
		AdvertiseMap:   o.AdvertiseMap,
		NonExistMap:    o.NonExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastFilterListXml) MarshalFromObject(s Ipv4UnicastFilterList) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastFilterListXml) UnmarshalToObject() (*Ipv4UnicastFilterList, error) {

	result := &Ipv4UnicastFilterList{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastInboundNetworkFiltersXml) MarshalFromObject(s Ipv4UnicastInboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastInboundNetworkFiltersXml) UnmarshalToObject() (*Ipv4UnicastInboundNetworkFilters, error) {

	result := &Ipv4UnicastInboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastOutboundNetworkFiltersXml) MarshalFromObject(s Ipv4UnicastOutboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastOutboundNetworkFiltersXml) UnmarshalToObject() (*Ipv4UnicastOutboundNetworkFilters, error) {

	result := &Ipv4UnicastOutboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastRouteMapsXml) MarshalFromObject(s Ipv4UnicastRouteMaps) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastRouteMapsXml) UnmarshalToObject() (*Ipv4UnicastRouteMaps, error) {

	result := &Ipv4UnicastRouteMaps{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
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
	if s.ConditionalAdvertisement != nil {
		var obj ipv6UnicastConditionalAdvertisementXml
		obj.MarshalFromObject(*s.ConditionalAdvertisement)
		o.ConditionalAdvertisement = &obj
	}
	if s.FilterList != nil {
		var obj ipv6UnicastFilterListXml
		obj.MarshalFromObject(*s.FilterList)
		o.FilterList = &obj
	}
	if s.InboundNetworkFilters != nil {
		var obj ipv6UnicastInboundNetworkFiltersXml
		obj.MarshalFromObject(*s.InboundNetworkFilters)
		o.InboundNetworkFilters = &obj
	}
	if s.OutboundNetworkFilters != nil {
		var obj ipv6UnicastOutboundNetworkFiltersXml
		obj.MarshalFromObject(*s.OutboundNetworkFilters)
		o.OutboundNetworkFilters = &obj
	}
	if s.RouteMaps != nil {
		var obj ipv6UnicastRouteMapsXml
		obj.MarshalFromObject(*s.RouteMaps)
		o.RouteMaps = &obj
	}
	o.UnsuppressMap = s.UnsuppressMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastXml) UnmarshalToObject() (*Ipv6Unicast, error) {
	var conditionalAdvertisementVal *Ipv6UnicastConditionalAdvertisement
	if o.ConditionalAdvertisement != nil {
		obj, err := o.ConditionalAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		conditionalAdvertisementVal = obj
	}
	var filterListVal *Ipv6UnicastFilterList
	if o.FilterList != nil {
		obj, err := o.FilterList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filterListVal = obj
	}
	var inboundNetworkFiltersVal *Ipv6UnicastInboundNetworkFilters
	if o.InboundNetworkFilters != nil {
		obj, err := o.InboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inboundNetworkFiltersVal = obj
	}
	var outboundNetworkFiltersVal *Ipv6UnicastOutboundNetworkFilters
	if o.OutboundNetworkFilters != nil {
		obj, err := o.OutboundNetworkFilters.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		outboundNetworkFiltersVal = obj
	}
	var routeMapsVal *Ipv6UnicastRouteMaps
	if o.RouteMaps != nil {
		obj, err := o.RouteMaps.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeMapsVal = obj
	}

	result := &Ipv6Unicast{
		ConditionalAdvertisement: conditionalAdvertisementVal,
		FilterList:               filterListVal,
		InboundNetworkFilters:    inboundNetworkFiltersVal,
		OutboundNetworkFilters:   outboundNetworkFiltersVal,
		RouteMaps:                routeMapsVal,
		UnsuppressMap:            o.UnsuppressMap,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastConditionalAdvertisementXml) MarshalFromObject(s Ipv6UnicastConditionalAdvertisement) {
	if s.Exist != nil {
		var obj ipv6UnicastConditionalAdvertisementExistXml
		obj.MarshalFromObject(*s.Exist)
		o.Exist = &obj
	}
	if s.NonExist != nil {
		var obj ipv6UnicastConditionalAdvertisementNonExistXml
		obj.MarshalFromObject(*s.NonExist)
		o.NonExist = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastConditionalAdvertisementXml) UnmarshalToObject() (*Ipv6UnicastConditionalAdvertisement, error) {
	var existVal *Ipv6UnicastConditionalAdvertisementExist
	if o.Exist != nil {
		obj, err := o.Exist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		existVal = obj
	}
	var nonExistVal *Ipv6UnicastConditionalAdvertisementNonExist
	if o.NonExist != nil {
		obj, err := o.NonExist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nonExistVal = obj
	}

	result := &Ipv6UnicastConditionalAdvertisement{
		Exist:          existVal,
		NonExist:       nonExistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastConditionalAdvertisementExistXml) MarshalFromObject(s Ipv6UnicastConditionalAdvertisementExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.ExistMap = s.ExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastConditionalAdvertisementExistXml) UnmarshalToObject() (*Ipv6UnicastConditionalAdvertisementExist, error) {

	result := &Ipv6UnicastConditionalAdvertisementExist{
		AdvertiseMap:   o.AdvertiseMap,
		ExistMap:       o.ExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastConditionalAdvertisementNonExistXml) MarshalFromObject(s Ipv6UnicastConditionalAdvertisementNonExist) {
	o.AdvertiseMap = s.AdvertiseMap
	o.NonExistMap = s.NonExistMap
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastConditionalAdvertisementNonExistXml) UnmarshalToObject() (*Ipv6UnicastConditionalAdvertisementNonExist, error) {

	result := &Ipv6UnicastConditionalAdvertisementNonExist{
		AdvertiseMap:   o.AdvertiseMap,
		NonExistMap:    o.NonExistMap,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastFilterListXml) MarshalFromObject(s Ipv6UnicastFilterList) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastFilterListXml) UnmarshalToObject() (*Ipv6UnicastFilterList, error) {

	result := &Ipv6UnicastFilterList{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastInboundNetworkFiltersXml) MarshalFromObject(s Ipv6UnicastInboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastInboundNetworkFiltersXml) UnmarshalToObject() (*Ipv6UnicastInboundNetworkFilters, error) {

	result := &Ipv6UnicastInboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastOutboundNetworkFiltersXml) MarshalFromObject(s Ipv6UnicastOutboundNetworkFilters) {
	o.DistributeList = s.DistributeList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastOutboundNetworkFiltersXml) UnmarshalToObject() (*Ipv6UnicastOutboundNetworkFilters, error) {

	result := &Ipv6UnicastOutboundNetworkFilters{
		DistributeList: o.DistributeList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastRouteMapsXml) MarshalFromObject(s Ipv6UnicastRouteMaps) {
	o.Inbound = s.Inbound
	o.Outbound = s.Outbound
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastRouteMapsXml) UnmarshalToObject() (*Ipv6UnicastRouteMaps, error) {

	result := &Ipv6UnicastRouteMaps{
		Inbound:        o.Inbound,
		Outbound:       o.Outbound,
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
	if !util.StringsMatch(o.Description, other.Description) {
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
	if !o.Multicast.matches(other.Multicast) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *Ipv4Multicast) matches(other *Ipv4Multicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.ConditionalAdvertisement.matches(other.ConditionalAdvertisement) {
		return false
	}
	if !o.FilterList.matches(other.FilterList) {
		return false
	}
	if !o.InboundNetworkFilters.matches(other.InboundNetworkFilters) {
		return false
	}
	if !util.BoolsMatch(o.Inherit, other.Inherit) {
		return false
	}
	if !o.OutboundNetworkFilters.matches(other.OutboundNetworkFilters) {
		return false
	}
	if !o.RouteMaps.matches(other.RouteMaps) {
		return false
	}
	if !util.StringsMatch(o.UnsuppressMap, other.UnsuppressMap) {
		return false
	}

	return true
}

func (o *Ipv4MulticastConditionalAdvertisement) matches(other *Ipv4MulticastConditionalAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Exist.matches(other.Exist) {
		return false
	}
	if !o.NonExist.matches(other.NonExist) {
		return false
	}

	return true
}

func (o *Ipv4MulticastConditionalAdvertisementExist) matches(other *Ipv4MulticastConditionalAdvertisementExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.ExistMap, other.ExistMap) {
		return false
	}

	return true
}

func (o *Ipv4MulticastConditionalAdvertisementNonExist) matches(other *Ipv4MulticastConditionalAdvertisementNonExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.NonExistMap, other.NonExistMap) {
		return false
	}

	return true
}

func (o *Ipv4MulticastFilterList) matches(other *Ipv4MulticastFilterList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
		return false
	}

	return true
}

func (o *Ipv4MulticastInboundNetworkFilters) matches(other *Ipv4MulticastInboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv4MulticastOutboundNetworkFilters) matches(other *Ipv4MulticastOutboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv4MulticastRouteMaps) matches(other *Ipv4MulticastRouteMaps) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
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
	if !o.ConditionalAdvertisement.matches(other.ConditionalAdvertisement) {
		return false
	}
	if !o.FilterList.matches(other.FilterList) {
		return false
	}
	if !o.InboundNetworkFilters.matches(other.InboundNetworkFilters) {
		return false
	}
	if !o.OutboundNetworkFilters.matches(other.OutboundNetworkFilters) {
		return false
	}
	if !o.RouteMaps.matches(other.RouteMaps) {
		return false
	}
	if !util.StringsMatch(o.UnsuppressMap, other.UnsuppressMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastConditionalAdvertisement) matches(other *Ipv4UnicastConditionalAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Exist.matches(other.Exist) {
		return false
	}
	if !o.NonExist.matches(other.NonExist) {
		return false
	}

	return true
}

func (o *Ipv4UnicastConditionalAdvertisementExist) matches(other *Ipv4UnicastConditionalAdvertisementExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.ExistMap, other.ExistMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastConditionalAdvertisementNonExist) matches(other *Ipv4UnicastConditionalAdvertisementNonExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.NonExistMap, other.NonExistMap) {
		return false
	}

	return true
}

func (o *Ipv4UnicastFilterList) matches(other *Ipv4UnicastFilterList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
		return false
	}

	return true
}

func (o *Ipv4UnicastInboundNetworkFilters) matches(other *Ipv4UnicastInboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv4UnicastOutboundNetworkFilters) matches(other *Ipv4UnicastOutboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv4UnicastRouteMaps) matches(other *Ipv4UnicastRouteMaps) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
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
	if !o.ConditionalAdvertisement.matches(other.ConditionalAdvertisement) {
		return false
	}
	if !o.FilterList.matches(other.FilterList) {
		return false
	}
	if !o.InboundNetworkFilters.matches(other.InboundNetworkFilters) {
		return false
	}
	if !o.OutboundNetworkFilters.matches(other.OutboundNetworkFilters) {
		return false
	}
	if !o.RouteMaps.matches(other.RouteMaps) {
		return false
	}
	if !util.StringsMatch(o.UnsuppressMap, other.UnsuppressMap) {
		return false
	}

	return true
}

func (o *Ipv6UnicastConditionalAdvertisement) matches(other *Ipv6UnicastConditionalAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Exist.matches(other.Exist) {
		return false
	}
	if !o.NonExist.matches(other.NonExist) {
		return false
	}

	return true
}

func (o *Ipv6UnicastConditionalAdvertisementExist) matches(other *Ipv6UnicastConditionalAdvertisementExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.ExistMap, other.ExistMap) {
		return false
	}

	return true
}

func (o *Ipv6UnicastConditionalAdvertisementNonExist) matches(other *Ipv6UnicastConditionalAdvertisementNonExist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AdvertiseMap, other.AdvertiseMap) {
		return false
	}
	if !util.StringsMatch(o.NonExistMap, other.NonExistMap) {
		return false
	}

	return true
}

func (o *Ipv6UnicastFilterList) matches(other *Ipv6UnicastFilterList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
		return false
	}

	return true
}

func (o *Ipv6UnicastInboundNetworkFilters) matches(other *Ipv6UnicastInboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv6UnicastOutboundNetworkFilters) matches(other *Ipv6UnicastOutboundNetworkFilters) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DistributeList, other.DistributeList) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}

	return true
}

func (o *Ipv6UnicastRouteMaps) matches(other *Ipv6UnicastRouteMaps) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Inbound, other.Inbound) {
		return false
	}
	if !util.StringsMatch(o.Outbound, other.Outbound) {
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
