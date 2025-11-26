package routemapsredistribution

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
	suffix = []string{"network", "routing-profile", "filters", "route-maps", "redistribution", "redist-entry", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	Bgp             *Bgp
	ConnectedStatic *ConnectedStatic
	Ospf            *Ospf
	Ospfv3          *Ospfv3
	Rip             *Rip
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type Bgp struct {
	Ospf           *BgpOspf
	Ospfv3         *BgpOspfv3
	Rib            *BgpRib
	Rip            *BgpRip
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspf struct {
	RouteMap       []BgpOspfRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *BgpOspfRouteMapMatch
	Set            *BgpOspfRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapMatch struct {
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
	Ipv4              *BgpOspfRouteMapMatchIpv4
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type BgpOspfRouteMapMatchIpv4 struct {
	Address        *BgpOspfRouteMapMatchIpv4Address
	NextHop        *BgpOspfRouteMapMatchIpv4NextHop
	RouteSource    *BgpOspfRouteMapMatchIpv4RouteSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapMatchIpv4RouteSource struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapSet struct {
	Metric         *BgpOspfRouteMapSetMetric
	MetricType     *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3 struct {
	RouteMap       []BgpOspfv3RouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *BgpOspfv3RouteMapMatch
	Set            *BgpOspfv3RouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMapMatch struct {
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
	Ipv6              *BgpOspfv3RouteMapMatchIpv6
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type BgpOspfv3RouteMapMatchIpv6 struct {
	Address        *BgpOspfv3RouteMapMatchIpv6Address
	NextHop        *BgpOspfv3RouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMapSet struct {
	Metric         *BgpOspfv3RouteMapSetMetric
	MetricType     *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpOspfv3RouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRib struct {
	RouteMap       []BgpRibRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *BgpRibRouteMapMatch
	Set            *BgpRibRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatch struct {
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
	Ipv4              *BgpRibRouteMapMatchIpv4
	Ipv6              *BgpRibRouteMapMatchIpv6
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type BgpRibRouteMapMatchIpv4 struct {
	Address        *BgpRibRouteMapMatchIpv4Address
	NextHop        *BgpRibRouteMapMatchIpv4NextHop
	RouteSource    *BgpRibRouteMapMatchIpv4RouteSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv4RouteSource struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv6 struct {
	Address        *BgpRibRouteMapMatchIpv6Address
	NextHop        *BgpRibRouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRibRouteMapSet struct {
	SourceAddress  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRip struct {
	RouteMap       []BgpRipRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *BgpRipRouteMapMatch
	Set            *BgpRipRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapMatch struct {
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
	Ipv4              *BgpRipRouteMapMatchIpv4
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type BgpRipRouteMapMatchIpv4 struct {
	Address        *BgpRipRouteMapMatchIpv4Address
	NextHop        *BgpRipRouteMapMatchIpv4NextHop
	RouteSource    *BgpRipRouteMapMatchIpv4RouteSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapMatchIpv4RouteSource struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapSet struct {
	Metric         *BgpRipRouteMapSetMetric
	NextHop        *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type BgpRipRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStatic struct {
	Bgp            *ConnectedStaticBgp
	Ospf           *ConnectedStaticOspf
	Ospfv3         *ConnectedStaticOspfv3
	Rib            *ConnectedStaticRib
	Rip            *ConnectedStaticRip
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgp struct {
	RouteMap       []ConnectedStaticBgpRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *ConnectedStaticBgpRouteMapMatch
	Set            *ConnectedStaticBgpRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Ipv4           *ConnectedStaticBgpRouteMapMatchIpv4
	Ipv6           *ConnectedStaticBgpRouteMapMatchIpv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv4 struct {
	Address        *ConnectedStaticBgpRouteMapMatchIpv4Address
	NextHop        *ConnectedStaticBgpRouteMapMatchIpv4NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv6 struct {
	Address        *ConnectedStaticBgpRouteMapMatchIpv6Address
	NextHop        *ConnectedStaticBgpRouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapSet struct {
	Aggregator        *ConnectedStaticBgpRouteMapSetAggregator
	Metric            *ConnectedStaticBgpRouteMapSetMetric
	Ipv4              *ConnectedStaticBgpRouteMapSetIpv4
	Ipv6              *ConnectedStaticBgpRouteMapSetIpv6
	Tag               *int64
	LocalPreference   *int64
	Weight            *int64
	Origin            *string
	AtomicAggregate   *bool
	OriginatorId      *string
	AspathPrepend     []int64
	RegularCommunity  []string
	LargeCommunity    []string
	ExtendedCommunity []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type ConnectedStaticBgpRouteMapSetAggregator struct {
	As             *int64
	RouterId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapSetIpv4 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticBgpRouteMapSetIpv6 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspf struct {
	RouteMap       []ConnectedStaticOspfRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *ConnectedStaticOspfRouteMapMatch
	Set            *ConnectedStaticOspfRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Ipv4           *ConnectedStaticOspfRouteMapMatchIpv4
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapMatchIpv4 struct {
	Address        *ConnectedStaticOspfRouteMapMatchIpv4Address
	NextHop        *ConnectedStaticOspfRouteMapMatchIpv4NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapSet struct {
	Metric         *ConnectedStaticOspfRouteMapSetMetric
	MetricType     *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3 struct {
	RouteMap       []ConnectedStaticOspfv3RouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *ConnectedStaticOspfv3RouteMapMatch
	Set            *ConnectedStaticOspfv3RouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Ipv6           *ConnectedStaticOspfv3RouteMapMatchIpv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapMatchIpv6 struct {
	Address        *ConnectedStaticOspfv3RouteMapMatchIpv6Address
	NextHop        *ConnectedStaticOspfv3RouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapSet struct {
	Metric         *ConnectedStaticOspfv3RouteMapSetMetric
	MetricType     *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticOspfv3RouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRib struct {
	RouteMap       []ConnectedStaticRibRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *ConnectedStaticRibRouteMapMatch
	Set            *ConnectedStaticRibRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Ipv4           *ConnectedStaticRibRouteMapMatchIpv4
	Ipv6           *ConnectedStaticRibRouteMapMatchIpv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv4 struct {
	Address        *ConnectedStaticRibRouteMapMatchIpv4Address
	NextHop        *ConnectedStaticRibRouteMapMatchIpv4NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv6 struct {
	Address        *ConnectedStaticRibRouteMapMatchIpv6Address
	NextHop        *ConnectedStaticRibRouteMapMatchIpv6NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv6Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapMatchIpv6NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRibRouteMapSet struct {
	SourceAddress  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRip struct {
	RouteMap       []ConnectedStaticRipRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *ConnectedStaticRipRouteMapMatch
	Set            *ConnectedStaticRipRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Ipv4           *ConnectedStaticRipRouteMapMatchIpv4
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapMatchIpv4 struct {
	Address        *ConnectedStaticRipRouteMapMatchIpv4Address
	NextHop        *ConnectedStaticRipRouteMapMatchIpv4NextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapMatchIpv4Address struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapMatchIpv4NextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapSet struct {
	Metric         *ConnectedStaticRipRouteMapSetMetric
	NextHop        *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConnectedStaticRipRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospf struct {
	Bgp            *OspfBgp
	Rib            *OspfRib
	Rip            *OspfRip
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgp struct {
	RouteMap       []OspfBgpRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *OspfBgpRouteMapMatch
	Set            *OspfBgpRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *OspfBgpRouteMapMatchAddress
	NextHop        *OspfBgpRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapSet struct {
	Aggregator        *OspfBgpRouteMapSetAggregator
	Metric            *OspfBgpRouteMapSetMetric
	Ipv4              *OspfBgpRouteMapSetIpv4
	Tag               *int64
	LocalPreference   *int64
	Weight            *int64
	Origin            *string
	AtomicAggregate   *bool
	OriginatorId      *string
	AspathPrepend     []int64
	RegularCommunity  []string
	LargeCommunity    []string
	ExtendedCommunity []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type OspfBgpRouteMapSetAggregator struct {
	As             *int64
	RouterId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfBgpRouteMapSetIpv4 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRib struct {
	RouteMap       []OspfRibRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRibRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *OspfRibRouteMapMatch
	Set            *OspfRibRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRibRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *OspfRibRouteMapMatchAddress
	NextHop        *OspfRibRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRibRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRibRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRibRouteMapSet struct {
	SourceAddress  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRip struct {
	RouteMap       []OspfRipRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *OspfRipRouteMapMatch
	Set            *OspfRipRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *OspfRipRouteMapMatchAddress
	NextHop        *OspfRipRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMapSet struct {
	Metric         *OspfRipRouteMapSetMetric
	NextHop        *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OspfRipRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3 struct {
	Bgp            *Ospfv3Bgp
	Rib            *Ospfv3Rib
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3Bgp struct {
	RouteMap       []Ospfv3BgpRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *Ospfv3BgpRouteMapMatch
	Set            *Ospfv3BgpRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *Ospfv3BgpRouteMapMatchAddress
	NextHop        *Ospfv3BgpRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapMatchAddress struct {
	PrefixList     *string
	AccessList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapSet struct {
	Aggregator        *Ospfv3BgpRouteMapSetAggregator
	Metric            *Ospfv3BgpRouteMapSetMetric
	Ipv6              *Ospfv3BgpRouteMapSetIpv6
	Tag               *int64
	LocalPreference   *int64
	Weight            *int64
	Origin            *string
	AtomicAggregate   *bool
	OriginatorId      *string
	AspathPrepend     []int64
	RegularCommunity  []string
	LargeCommunity    []string
	ExtendedCommunity []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ospfv3BgpRouteMapSetAggregator struct {
	As             *int64
	RouterId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3BgpRouteMapSetIpv6 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3Rib struct {
	RouteMap       []Ospfv3RibRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3RibRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *Ospfv3RibRouteMapMatch
	Set            *Ospfv3RibRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3RibRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *Ospfv3RibRouteMapMatchAddress
	NextHop        *Ospfv3RibRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3RibRouteMapMatchAddress struct {
	PrefixList     *string
	AccessList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3RibRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ospfv3RibRouteMapSet struct {
	SourceAddress  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Rip struct {
	Bgp            *RipBgp
	Ospf           *RipOspf
	Rib            *RipRib
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgp struct {
	RouteMap       []RipBgpRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *RipBgpRouteMapMatch
	Set            *RipBgpRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *RipBgpRouteMapMatchAddress
	NextHop        *RipBgpRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapSet struct {
	Aggregator        *RipBgpRouteMapSetAggregator
	Metric            *RipBgpRouteMapSetMetric
	Ipv4              *RipBgpRouteMapSetIpv4
	Tag               *int64
	LocalPreference   *int64
	Weight            *int64
	Origin            *string
	AtomicAggregate   *bool
	OriginatorId      *string
	AspathPrepend     []int64
	RegularCommunity  []string
	LargeCommunity    []string
	ExtendedCommunity []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type RipBgpRouteMapSetAggregator struct {
	As             *int64
	RouterId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipBgpRouteMapSetIpv4 struct {
	SourceAddress  *string
	NextHop        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspf struct {
	RouteMap       []RipOspfRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *RipOspfRouteMapMatch
	Set            *RipOspfRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *RipOspfRouteMapMatchAddress
	NextHop        *RipOspfRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMapSet struct {
	Metric         *RipOspfRouteMapSetMetric
	MetricType     *string
	Tag            *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipOspfRouteMapSetMetric struct {
	Value          *int64
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRib struct {
	RouteMap       []RipRibRouteMap
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRibRouteMap struct {
	Name           string
	Action         *string
	Description    *string
	Match          *RipRibRouteMapMatch
	Set            *RipRibRouteMapSet
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRibRouteMapMatch struct {
	Interface      *string
	Metric         *int64
	Tag            *int64
	Address        *RipRibRouteMapMatchAddress
	NextHop        *RipRibRouteMapMatchNextHop
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRibRouteMapMatchAddress struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRibRouteMapMatchNextHop struct {
	AccessList     *string
	PrefixList     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RipRibRouteMapSet struct {
	SourceAddress  *string
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
	XMLName         xml.Name            `xml:"entry"`
	Name            string              `xml:"name,attr"`
	Description     *string             `xml:"description,omitempty"`
	Bgp             *bgpXml             `xml:"bgp,omitempty"`
	ConnectedStatic *connectedStaticXml `xml:"connected-static,omitempty"`
	Ospf            *ospfXml            `xml:"ospf,omitempty"`
	Ospfv3          *ospfv3Xml          `xml:"ospfv3,omitempty"`
	Rip             *ripXml             `xml:"rip,omitempty"`
	Misc            []generic.Xml       `xml:",any"`
	MiscAttributes  []xml.Attr          `xml:",any,attr"`
}
type bgpXml struct {
	Ospf           *bgpOspfXml   `xml:"ospf,omitempty"`
	Ospfv3         *bgpOspfv3Xml `xml:"ospfv3,omitempty"`
	Rib            *bgpRibXml    `xml:"rib,omitempty"`
	Rip            *bgpRipXml    `xml:"rip,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfXml struct {
	RouteMap       *bgpOspfRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type bgpOspfRouteMapContainerXml struct {
	Entries []bgpOspfRouteMapXml `xml:"entry"`
}
type bgpOspfRouteMapXml struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *bgpOspfRouteMapMatchXml `xml:"match,omitempty"`
	Set            *bgpOspfRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type bgpOspfRouteMapMatchXml struct {
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
	Ipv4              *bgpOspfRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Misc              []generic.Xml                `xml:",any"`
	MiscAttributes    []xml.Attr                   `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4Xml struct {
	Address        *bgpOspfRouteMapMatchIpv4AddressXml     `xml:"address,omitempty"`
	NextHop        *bgpOspfRouteMapMatchIpv4NextHopXml     `xml:"next-hop,omitempty"`
	RouteSource    *bgpOspfRouteMapMatchIpv4RouteSourceXml `xml:"route-source,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4RouteSourceXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapSetXml struct {
	Metric         *bgpOspfRouteMapSetMetricXml `xml:"metric,omitempty"`
	MetricType     *string                      `xml:"metric-type,omitempty"`
	Tag            *int64                       `xml:"tag,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type bgpOspfRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3Xml struct {
	RouteMap       *bgpOspfv3RouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type bgpOspfv3RouteMapContainerXml struct {
	Entries []bgpOspfv3RouteMapXml `xml:"entry"`
}
type bgpOspfv3RouteMapXml struct {
	XMLName        xml.Name                   `xml:"entry"`
	Name           string                     `xml:"name,attr"`
	Action         *string                    `xml:"action,omitempty"`
	Description    *string                    `xml:"description,omitempty"`
	Match          *bgpOspfv3RouteMapMatchXml `xml:"match,omitempty"`
	Set            *bgpOspfv3RouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchXml struct {
	AsPathAccessList  *string                        `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                        `xml:"regular-community,omitempty"`
	LargeCommunity    *string                        `xml:"large-community,omitempty"`
	ExtendedCommunity *string                        `xml:"extended-community,omitempty"`
	Interface         *string                        `xml:"interface,omitempty"`
	Origin            *string                        `xml:"origin,omitempty"`
	Metric            *int64                         `xml:"metric,omitempty"`
	Tag               *int64                         `xml:"tag,omitempty"`
	LocalPreference   *int64                         `xml:"local-preference,omitempty"`
	Peer              *string                        `xml:"peer,omitempty"`
	Ipv6              *bgpOspfv3RouteMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc              []generic.Xml                  `xml:",any"`
	MiscAttributes    []xml.Attr                     `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6Xml struct {
	Address        *bgpOspfv3RouteMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *bgpOspfv3RouteMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3RouteMapSetXml struct {
	Metric         *bgpOspfv3RouteMapSetMetricXml `xml:"metric,omitempty"`
	MetricType     *string                        `xml:"metric-type,omitempty"`
	Tag            *int64                         `xml:"tag,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type bgpOspfv3RouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibXml struct {
	RouteMap       *bgpRibRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type bgpRibRouteMapContainerXml struct {
	Entries []bgpRibRouteMapXml `xml:"entry"`
}
type bgpRibRouteMapXml struct {
	XMLName        xml.Name                `xml:"entry"`
	Name           string                  `xml:"name,attr"`
	Action         *string                 `xml:"action,omitempty"`
	Description    *string                 `xml:"description,omitempty"`
	Match          *bgpRibRouteMapMatchXml `xml:"match,omitempty"`
	Set            *bgpRibRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type bgpRibRouteMapMatchXml struct {
	AsPathAccessList  *string                     `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                     `xml:"regular-community,omitempty"`
	LargeCommunity    *string                     `xml:"large-community,omitempty"`
	ExtendedCommunity *string                     `xml:"extended-community,omitempty"`
	Interface         *string                     `xml:"interface,omitempty"`
	Origin            *string                     `xml:"origin,omitempty"`
	Metric            *int64                      `xml:"metric,omitempty"`
	Tag               *int64                      `xml:"tag,omitempty"`
	LocalPreference   *int64                      `xml:"local-preference,omitempty"`
	Peer              *string                     `xml:"peer,omitempty"`
	Ipv4              *bgpRibRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6              *bgpRibRouteMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc              []generic.Xml               `xml:",any"`
	MiscAttributes    []xml.Attr                  `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4Xml struct {
	Address        *bgpRibRouteMapMatchIpv4AddressXml     `xml:"address,omitempty"`
	NextHop        *bgpRibRouteMapMatchIpv4NextHopXml     `xml:"next-hop,omitempty"`
	RouteSource    *bgpRibRouteMapMatchIpv4RouteSourceXml `xml:"route-source,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4RouteSourceXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6Xml struct {
	Address        *bgpRibRouteMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *bgpRibRouteMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapSetXml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipXml struct {
	RouteMap       *bgpRipRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type bgpRipRouteMapContainerXml struct {
	Entries []bgpRipRouteMapXml `xml:"entry"`
}
type bgpRipRouteMapXml struct {
	XMLName        xml.Name                `xml:"entry"`
	Name           string                  `xml:"name,attr"`
	Action         *string                 `xml:"action,omitempty"`
	Description    *string                 `xml:"description,omitempty"`
	Match          *bgpRipRouteMapMatchXml `xml:"match,omitempty"`
	Set            *bgpRipRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type bgpRipRouteMapMatchXml struct {
	AsPathAccessList  *string                     `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                     `xml:"regular-community,omitempty"`
	LargeCommunity    *string                     `xml:"large-community,omitempty"`
	ExtendedCommunity *string                     `xml:"extended-community,omitempty"`
	Interface         *string                     `xml:"interface,omitempty"`
	Origin            *string                     `xml:"origin,omitempty"`
	Metric            *int64                      `xml:"metric,omitempty"`
	Tag               *int64                      `xml:"tag,omitempty"`
	LocalPreference   *int64                      `xml:"local-preference,omitempty"`
	Peer              *string                     `xml:"peer,omitempty"`
	Ipv4              *bgpRipRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Misc              []generic.Xml               `xml:",any"`
	MiscAttributes    []xml.Attr                  `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4Xml struct {
	Address        *bgpRipRouteMapMatchIpv4AddressXml     `xml:"address,omitempty"`
	NextHop        *bgpRipRouteMapMatchIpv4NextHopXml     `xml:"next-hop,omitempty"`
	RouteSource    *bgpRipRouteMapMatchIpv4RouteSourceXml `xml:"route-source,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4RouteSourceXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapSetXml struct {
	Metric         *bgpRipRouteMapSetMetricXml `xml:"metric,omitempty"`
	NextHop        *string                     `xml:"next-hop,omitempty"`
	Tag            *int64                      `xml:"tag,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type bgpRipRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticXml struct {
	Bgp            *connectedStaticBgpXml    `xml:"bgp,omitempty"`
	Ospf           *connectedStaticOspfXml   `xml:"ospf,omitempty"`
	Ospfv3         *connectedStaticOspfv3Xml `xml:"ospfv3,omitempty"`
	Rib            *connectedStaticRibXml    `xml:"rib,omitempty"`
	Rip            *connectedStaticRipXml    `xml:"rip,omitempty"`
	Misc           []generic.Xml             `xml:",any"`
	MiscAttributes []xml.Attr                `xml:",any,attr"`
}
type connectedStaticBgpXml struct {
	RouteMap       *connectedStaticBgpRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticBgpRouteMapContainerXml struct {
	Entries []connectedStaticBgpRouteMapXml `xml:"entry"`
}
type connectedStaticBgpRouteMapXml struct {
	XMLName        xml.Name                            `xml:"entry"`
	Name           string                              `xml:"name,attr"`
	Action         *string                             `xml:"action,omitempty"`
	Description    *string                             `xml:"description,omitempty"`
	Match          *connectedStaticBgpRouteMapMatchXml `xml:"match,omitempty"`
	Set            *connectedStaticBgpRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchXml struct {
	Interface      *string                                 `xml:"interface,omitempty"`
	Metric         *int64                                  `xml:"metric,omitempty"`
	Ipv4           *connectedStaticBgpRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6           *connectedStaticBgpRouteMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4Xml struct {
	Address        *connectedStaticBgpRouteMapMatchIpv4AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticBgpRouteMapMatchIpv4NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6Xml struct {
	Address        *connectedStaticBgpRouteMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticBgpRouteMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetXml struct {
	Aggregator        *connectedStaticBgpRouteMapSetAggregatorXml `xml:"aggregator,omitempty"`
	Metric            *connectedStaticBgpRouteMapSetMetricXml     `xml:"metric,omitempty"`
	Ipv4              *connectedStaticBgpRouteMapSetIpv4Xml       `xml:"ipv4,omitempty"`
	Ipv6              *connectedStaticBgpRouteMapSetIpv6Xml       `xml:"ipv6,omitempty"`
	Tag               *int64                                      `xml:"tag,omitempty"`
	LocalPreference   *int64                                      `xml:"local-preference,omitempty"`
	Weight            *int64                                      `xml:"weight,omitempty"`
	Origin            *string                                     `xml:"origin,omitempty"`
	AtomicAggregate   *string                                     `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                                     `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                            `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                            `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                            `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                            `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                               `xml:",any"`
	MiscAttributes    []xml.Attr                                  `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetAggregatorXml struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetIpv4Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetIpv6Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfXml struct {
	RouteMap       *connectedStaticOspfRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type connectedStaticOspfRouteMapContainerXml struct {
	Entries []connectedStaticOspfRouteMapXml `xml:"entry"`
}
type connectedStaticOspfRouteMapXml struct {
	XMLName        xml.Name                             `xml:"entry"`
	Name           string                               `xml:"name,attr"`
	Action         *string                              `xml:"action,omitempty"`
	Description    *string                              `xml:"description,omitempty"`
	Match          *connectedStaticOspfRouteMapMatchXml `xml:"match,omitempty"`
	Set            *connectedStaticOspfRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchXml struct {
	Interface      *string                                  `xml:"interface,omitempty"`
	Metric         *int64                                   `xml:"metric,omitempty"`
	Ipv4           *connectedStaticOspfRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4Xml struct {
	Address        *connectedStaticOspfRouteMapMatchIpv4AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticOspfRouteMapMatchIpv4NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfRouteMapSetXml struct {
	Metric         *connectedStaticOspfRouteMapSetMetricXml `xml:"metric,omitempty"`
	MetricType     *string                                  `xml:"metric-type,omitempty"`
	Tag            *int64                                   `xml:"tag,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type connectedStaticOspfRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3Xml struct {
	RouteMap       *connectedStaticOspfv3RouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapContainerXml struct {
	Entries []connectedStaticOspfv3RouteMapXml `xml:"entry"`
}
type connectedStaticOspfv3RouteMapXml struct {
	XMLName        xml.Name                               `xml:"entry"`
	Name           string                                 `xml:"name,attr"`
	Action         *string                                `xml:"action,omitempty"`
	Description    *string                                `xml:"description,omitempty"`
	Match          *connectedStaticOspfv3RouteMapMatchXml `xml:"match,omitempty"`
	Set            *connectedStaticOspfv3RouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchXml struct {
	Interface      *string                                    `xml:"interface,omitempty"`
	Metric         *int64                                     `xml:"metric,omitempty"`
	Ipv6           *connectedStaticOspfv3RouteMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6Xml struct {
	Address        *connectedStaticOspfv3RouteMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticOspfv3RouteMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapSetXml struct {
	Metric         *connectedStaticOspfv3RouteMapSetMetricXml `xml:"metric,omitempty"`
	MetricType     *string                                    `xml:"metric-type,omitempty"`
	Tag            *int64                                     `xml:"tag,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibXml struct {
	RouteMap       *connectedStaticRibRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticRibRouteMapContainerXml struct {
	Entries []connectedStaticRibRouteMapXml `xml:"entry"`
}
type connectedStaticRibRouteMapXml struct {
	XMLName        xml.Name                            `xml:"entry"`
	Name           string                              `xml:"name,attr"`
	Action         *string                             `xml:"action,omitempty"`
	Description    *string                             `xml:"description,omitempty"`
	Match          *connectedStaticRibRouteMapMatchXml `xml:"match,omitempty"`
	Set            *connectedStaticRibRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchXml struct {
	Interface      *string                                 `xml:"interface,omitempty"`
	Metric         *int64                                  `xml:"metric,omitempty"`
	Ipv4           *connectedStaticRibRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6           *connectedStaticRibRouteMapMatchIpv6Xml `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4Xml struct {
	Address        *connectedStaticRibRouteMapMatchIpv4AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticRibRouteMapMatchIpv4NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6Xml struct {
	Address        *connectedStaticRibRouteMapMatchIpv6AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticRibRouteMapMatchIpv6NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapSetXml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipXml struct {
	RouteMap       *connectedStaticRipRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticRipRouteMapContainerXml struct {
	Entries []connectedStaticRipRouteMapXml `xml:"entry"`
}
type connectedStaticRipRouteMapXml struct {
	XMLName        xml.Name                            `xml:"entry"`
	Name           string                              `xml:"name,attr"`
	Action         *string                             `xml:"action,omitempty"`
	Description    *string                             `xml:"description,omitempty"`
	Match          *connectedStaticRipRouteMapMatchXml `xml:"match,omitempty"`
	Set            *connectedStaticRipRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchXml struct {
	Interface      *string                                 `xml:"interface,omitempty"`
	Metric         *int64                                  `xml:"metric,omitempty"`
	Ipv4           *connectedStaticRipRouteMapMatchIpv4Xml `xml:"ipv4,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4Xml struct {
	Address        *connectedStaticRipRouteMapMatchIpv4AddressXml `xml:"address,omitempty"`
	NextHop        *connectedStaticRipRouteMapMatchIpv4NextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4AddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4NextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipRouteMapSetXml struct {
	Metric         *connectedStaticRipRouteMapSetMetricXml `xml:"metric,omitempty"`
	NextHop        *string                                 `xml:"next-hop,omitempty"`
	Tag            *int64                                  `xml:"tag,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type connectedStaticRipRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfXml struct {
	Bgp            *ospfBgpXml   `xml:"bgp,omitempty"`
	Rib            *ospfRibXml   `xml:"rib,omitempty"`
	Rip            *ospfRipXml   `xml:"rip,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpXml struct {
	RouteMap       *ospfBgpRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ospfBgpRouteMapContainerXml struct {
	Entries []ospfBgpRouteMapXml `xml:"entry"`
}
type ospfBgpRouteMapXml struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *ospfBgpRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ospfBgpRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ospfBgpRouteMapMatchXml struct {
	Interface      *string                         `xml:"interface,omitempty"`
	Metric         *int64                          `xml:"metric,omitempty"`
	Tag            *int64                          `xml:"tag,omitempty"`
	Address        *ospfBgpRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ospfBgpRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfBgpRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetXml struct {
	Aggregator        *ospfBgpRouteMapSetAggregatorXml `xml:"aggregator,omitempty"`
	Metric            *ospfBgpRouteMapSetMetricXml     `xml:"metric,omitempty"`
	Ipv4              *ospfBgpRouteMapSetIpv4Xml       `xml:"ipv4,omitempty"`
	Tag               *int64                           `xml:"tag,omitempty"`
	LocalPreference   *int64                           `xml:"local-preference,omitempty"`
	Weight            *int64                           `xml:"weight,omitempty"`
	Origin            *string                          `xml:"origin,omitempty"`
	AtomicAggregate   *string                          `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                          `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                 `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                 `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                 `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                 `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                    `xml:",any"`
	MiscAttributes    []xml.Attr                       `xml:",any,attr"`
}
type ospfBgpRouteMapSetAggregatorXml struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetIpv4Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibXml struct {
	RouteMap       *ospfRibRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ospfRibRouteMapContainerXml struct {
	Entries []ospfRibRouteMapXml `xml:"entry"`
}
type ospfRibRouteMapXml struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *ospfRibRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ospfRibRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ospfRibRouteMapMatchXml struct {
	Interface      *string                         `xml:"interface,omitempty"`
	Metric         *int64                          `xml:"metric,omitempty"`
	Tag            *int64                          `xml:"tag,omitempty"`
	Address        *ospfRibRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ospfRibRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfRibRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibRouteMapSetXml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipXml struct {
	RouteMap       *ospfRipRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ospfRipRouteMapContainerXml struct {
	Entries []ospfRipRouteMapXml `xml:"entry"`
}
type ospfRipRouteMapXml struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *ospfRipRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ospfRipRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ospfRipRouteMapMatchXml struct {
	Interface      *string                         `xml:"interface,omitempty"`
	Metric         *int64                          `xml:"metric,omitempty"`
	Tag            *int64                          `xml:"tag,omitempty"`
	Address        *ospfRipRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ospfRipRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfRipRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipRouteMapSetXml struct {
	Metric         *ospfRipRouteMapSetMetricXml `xml:"metric,omitempty"`
	NextHop        *string                      `xml:"next-hop,omitempty"`
	Tag            *int64                       `xml:"tag,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ospfRipRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3Xml struct {
	Bgp            *ospfv3BgpXml `xml:"bgp,omitempty"`
	Rib            *ospfv3RibXml `xml:"rib,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpXml struct {
	RouteMap       *ospfv3BgpRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ospfv3BgpRouteMapContainerXml struct {
	Entries []ospfv3BgpRouteMapXml `xml:"entry"`
}
type ospfv3BgpRouteMapXml struct {
	XMLName        xml.Name                   `xml:"entry"`
	Name           string                     `xml:"name,attr"`
	Action         *string                    `xml:"action,omitempty"`
	Description    *string                    `xml:"description,omitempty"`
	Match          *ospfv3BgpRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ospfv3BgpRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchXml struct {
	Interface      *string                           `xml:"interface,omitempty"`
	Metric         *int64                            `xml:"metric,omitempty"`
	Tag            *int64                            `xml:"tag,omitempty"`
	Address        *ospfv3BgpRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ospfv3BgpRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchAddressXml struct {
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	AccessList     *string       `xml:"access-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetXml struct {
	Aggregator        *ospfv3BgpRouteMapSetAggregatorXml `xml:"aggregator,omitempty"`
	Metric            *ospfv3BgpRouteMapSetMetricXml     `xml:"metric,omitempty"`
	Ipv6              *ospfv3BgpRouteMapSetIpv6Xml       `xml:"ipv6,omitempty"`
	Tag               *int64                             `xml:"tag,omitempty"`
	LocalPreference   *int64                             `xml:"local-preference,omitempty"`
	Weight            *int64                             `xml:"weight,omitempty"`
	Origin            *string                            `xml:"origin,omitempty"`
	AtomicAggregate   *string                            `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                            `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                   `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                   `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                   `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                   `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                      `xml:",any"`
	MiscAttributes    []xml.Attr                         `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetAggregatorXml struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetIpv6Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibXml struct {
	RouteMap       *ospfv3RibRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ospfv3RibRouteMapContainerXml struct {
	Entries []ospfv3RibRouteMapXml `xml:"entry"`
}
type ospfv3RibRouteMapXml struct {
	XMLName        xml.Name                   `xml:"entry"`
	Name           string                     `xml:"name,attr"`
	Action         *string                    `xml:"action,omitempty"`
	Description    *string                    `xml:"description,omitempty"`
	Match          *ospfv3RibRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ospfv3RibRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchXml struct {
	Interface      *string                           `xml:"interface,omitempty"`
	Metric         *int64                            `xml:"metric,omitempty"`
	Tag            *int64                            `xml:"tag,omitempty"`
	Address        *ospfv3RibRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ospfv3RibRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchAddressXml struct {
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	AccessList     *string       `xml:"access-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibRouteMapSetXml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripXml struct {
	Bgp            *ripBgpXml    `xml:"bgp,omitempty"`
	Ospf           *ripOspfXml   `xml:"ospf,omitempty"`
	Rib            *ripRibXml    `xml:"rib,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpXml struct {
	RouteMap       *ripBgpRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type ripBgpRouteMapContainerXml struct {
	Entries []ripBgpRouteMapXml `xml:"entry"`
}
type ripBgpRouteMapXml struct {
	XMLName        xml.Name                `xml:"entry"`
	Name           string                  `xml:"name,attr"`
	Action         *string                 `xml:"action,omitempty"`
	Description    *string                 `xml:"description,omitempty"`
	Match          *ripBgpRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ripBgpRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type ripBgpRouteMapMatchXml struct {
	Interface      *string                        `xml:"interface,omitempty"`
	Metric         *int64                         `xml:"metric,omitempty"`
	Tag            *int64                         `xml:"tag,omitempty"`
	Address        *ripBgpRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ripBgpRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ripBgpRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetXml struct {
	Aggregator        *ripBgpRouteMapSetAggregatorXml `xml:"aggregator,omitempty"`
	Metric            *ripBgpRouteMapSetMetricXml     `xml:"metric,omitempty"`
	Ipv4              *ripBgpRouteMapSetIpv4Xml       `xml:"ipv4,omitempty"`
	Tag               *int64                          `xml:"tag,omitempty"`
	LocalPreference   *int64                          `xml:"local-preference,omitempty"`
	Weight            *int64                          `xml:"weight,omitempty"`
	Origin            *string                         `xml:"origin,omitempty"`
	AtomicAggregate   *string                         `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                         `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                   `xml:",any"`
	MiscAttributes    []xml.Attr                      `xml:",any,attr"`
}
type ripBgpRouteMapSetAggregatorXml struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetIpv4Xml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfXml struct {
	RouteMap       *ripOspfRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ripOspfRouteMapContainerXml struct {
	Entries []ripOspfRouteMapXml `xml:"entry"`
}
type ripOspfRouteMapXml struct {
	XMLName        xml.Name                 `xml:"entry"`
	Name           string                   `xml:"name,attr"`
	Action         *string                  `xml:"action,omitempty"`
	Description    *string                  `xml:"description,omitempty"`
	Match          *ripOspfRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ripOspfRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ripOspfRouteMapMatchXml struct {
	Interface      *string                         `xml:"interface,omitempty"`
	Metric         *int64                          `xml:"metric,omitempty"`
	Tag            *int64                          `xml:"tag,omitempty"`
	Address        *ripOspfRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ripOspfRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ripOspfRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfRouteMapSetXml struct {
	Metric         *ripOspfRouteMapSetMetricXml `xml:"metric,omitempty"`
	MetricType     *string                      `xml:"metric-type,omitempty"`
	Tag            *int64                       `xml:"tag,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ripOspfRouteMapSetMetricXml struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibXml struct {
	RouteMap       *ripRibRouteMapContainerXml `xml:"route-map,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type ripRibRouteMapContainerXml struct {
	Entries []ripRibRouteMapXml `xml:"entry"`
}
type ripRibRouteMapXml struct {
	XMLName        xml.Name                `xml:"entry"`
	Name           string                  `xml:"name,attr"`
	Action         *string                 `xml:"action,omitempty"`
	Description    *string                 `xml:"description,omitempty"`
	Match          *ripRibRouteMapMatchXml `xml:"match,omitempty"`
	Set            *ripRibRouteMapSetXml   `xml:"set,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type ripRibRouteMapMatchXml struct {
	Interface      *string                        `xml:"interface,omitempty"`
	Metric         *int64                         `xml:"metric,omitempty"`
	Tag            *int64                         `xml:"tag,omitempty"`
	Address        *ripRibRouteMapMatchAddressXml `xml:"address,omitempty"`
	NextHop        *ripRibRouteMapMatchNextHopXml `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ripRibRouteMapMatchAddressXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibRouteMapMatchNextHopXml struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibRouteMapSetXml struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type entryXml_11_0_2 struct {
	XMLName         xml.Name                   `xml:"entry"`
	Name            string                     `xml:"name,attr"`
	Description     *string                    `xml:"description,omitempty"`
	Bgp             *bgpXml_11_0_2             `xml:"bgp,omitempty"`
	ConnectedStatic *connectedStaticXml_11_0_2 `xml:"connected-static,omitempty"`
	Ospf            *ospfXml_11_0_2            `xml:"ospf,omitempty"`
	Ospfv3          *ospfv3Xml_11_0_2          `xml:"ospfv3,omitempty"`
	Rip             *ripXml_11_0_2             `xml:"rip,omitempty"`
	Misc            []generic.Xml              `xml:",any"`
	MiscAttributes  []xml.Attr                 `xml:",any,attr"`
}
type bgpXml_11_0_2 struct {
	Ospf           *bgpOspfXml_11_0_2   `xml:"ospf,omitempty"`
	Ospfv3         *bgpOspfv3Xml_11_0_2 `xml:"ospfv3,omitempty"`
	Rib            *bgpRibXml_11_0_2    `xml:"rib,omitempty"`
	Rip            *bgpRipXml_11_0_2    `xml:"rip,omitempty"`
	Misc           []generic.Xml        `xml:",any"`
	MiscAttributes []xml.Attr           `xml:",any,attr"`
}
type bgpOspfXml_11_0_2 struct {
	RouteMap       *bgpOspfRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type bgpOspfRouteMapContainerXml_11_0_2 struct {
	Entries []bgpOspfRouteMapXml_11_0_2 `xml:"entry"`
}
type bgpOspfRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Action         *string                         `xml:"action,omitempty"`
	Description    *string                         `xml:"description,omitempty"`
	Match          *bgpOspfRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *bgpOspfRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type bgpOspfRouteMapMatchXml_11_0_2 struct {
	AsPathAccessList  *string                             `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                             `xml:"regular-community,omitempty"`
	LargeCommunity    *string                             `xml:"large-community,omitempty"`
	ExtendedCommunity *string                             `xml:"extended-community,omitempty"`
	Interface         *string                             `xml:"interface,omitempty"`
	Origin            *string                             `xml:"origin,omitempty"`
	Metric            *int64                              `xml:"metric,omitempty"`
	Tag               *int64                              `xml:"tag,omitempty"`
	LocalPreference   *int64                              `xml:"local-preference,omitempty"`
	Peer              *string                             `xml:"peer,omitempty"`
	Ipv4              *bgpOspfRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Misc              []generic.Xml                       `xml:",any"`
	MiscAttributes    []xml.Attr                          `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *bgpOspfRouteMapMatchIpv4AddressXml_11_0_2     `xml:"address,omitempty"`
	NextHop        *bgpOspfRouteMapMatchIpv4NextHopXml_11_0_2     `xml:"next-hop,omitempty"`
	RouteSource    *bgpOspfRouteMapMatchIpv4RouteSourceXml_11_0_2 `xml:"route-source,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapMatchIpv4RouteSourceXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfRouteMapSetXml_11_0_2 struct {
	Metric         *bgpOspfRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	MetricType     *string                             `xml:"metric-type,omitempty"`
	Tag            *int64                              `xml:"tag,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type bgpOspfRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3Xml_11_0_2 struct {
	RouteMap       *bgpOspfv3RouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type bgpOspfv3RouteMapContainerXml_11_0_2 struct {
	Entries []bgpOspfv3RouteMapXml_11_0_2 `xml:"entry"`
}
type bgpOspfv3RouteMapXml_11_0_2 struct {
	XMLName        xml.Name                          `xml:"entry"`
	Name           string                            `xml:"name,attr"`
	Action         *string                           `xml:"action,omitempty"`
	Description    *string                           `xml:"description,omitempty"`
	Match          *bgpOspfv3RouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *bgpOspfv3RouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchXml_11_0_2 struct {
	AsPathAccessList  *string                               `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                               `xml:"regular-community,omitempty"`
	LargeCommunity    *string                               `xml:"large-community,omitempty"`
	ExtendedCommunity *string                               `xml:"extended-community,omitempty"`
	Interface         *string                               `xml:"interface,omitempty"`
	Origin            *string                               `xml:"origin,omitempty"`
	Metric            *int64                                `xml:"metric,omitempty"`
	Tag               *int64                                `xml:"tag,omitempty"`
	LocalPreference   *int64                                `xml:"local-preference,omitempty"`
	Peer              *string                               `xml:"peer,omitempty"`
	Ipv6              *bgpOspfv3RouteMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc              []generic.Xml                         `xml:",any"`
	MiscAttributes    []xml.Attr                            `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6Xml_11_0_2 struct {
	Address        *bgpOspfv3RouteMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *bgpOspfv3RouteMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                `xml:",any"`
	MiscAttributes []xml.Attr                                   `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3RouteMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpOspfv3RouteMapSetXml_11_0_2 struct {
	Metric         *bgpOspfv3RouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	MetricType     *string                               `xml:"metric-type,omitempty"`
	Tag            *int64                                `xml:"tag,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type bgpOspfv3RouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibXml_11_0_2 struct {
	RouteMap       *bgpRibRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type bgpRibRouteMapContainerXml_11_0_2 struct {
	Entries []bgpRibRouteMapXml_11_0_2 `xml:"entry"`
}
type bgpRibRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Action         *string                        `xml:"action,omitempty"`
	Description    *string                        `xml:"description,omitempty"`
	Match          *bgpRibRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *bgpRibRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type bgpRibRouteMapMatchXml_11_0_2 struct {
	AsPathAccessList  *string                            `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                            `xml:"regular-community,omitempty"`
	LargeCommunity    *string                            `xml:"large-community,omitempty"`
	ExtendedCommunity *string                            `xml:"extended-community,omitempty"`
	Interface         *string                            `xml:"interface,omitempty"`
	Origin            *string                            `xml:"origin,omitempty"`
	Metric            *int64                             `xml:"metric,omitempty"`
	Tag               *int64                             `xml:"tag,omitempty"`
	LocalPreference   *int64                             `xml:"local-preference,omitempty"`
	Peer              *string                            `xml:"peer,omitempty"`
	Ipv4              *bgpRibRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6              *bgpRibRouteMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc              []generic.Xml                      `xml:",any"`
	MiscAttributes    []xml.Attr                         `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *bgpRibRouteMapMatchIpv4AddressXml_11_0_2     `xml:"address,omitempty"`
	NextHop        *bgpRibRouteMapMatchIpv4NextHopXml_11_0_2     `xml:"next-hop,omitempty"`
	RouteSource    *bgpRibRouteMapMatchIpv4RouteSourceXml_11_0_2 `xml:"route-source,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv4RouteSourceXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6Xml_11_0_2 struct {
	Address        *bgpRibRouteMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *bgpRibRouteMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
	MiscAttributes []xml.Attr                                `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRibRouteMapSetXml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipXml_11_0_2 struct {
	RouteMap       *bgpRipRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type bgpRipRouteMapContainerXml_11_0_2 struct {
	Entries []bgpRipRouteMapXml_11_0_2 `xml:"entry"`
}
type bgpRipRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Action         *string                        `xml:"action,omitempty"`
	Description    *string                        `xml:"description,omitempty"`
	Match          *bgpRipRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *bgpRipRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type bgpRipRouteMapMatchXml_11_0_2 struct {
	AsPathAccessList  *string                            `xml:"as-path-access-list,omitempty"`
	RegularCommunity  *string                            `xml:"regular-community,omitempty"`
	LargeCommunity    *string                            `xml:"large-community,omitempty"`
	ExtendedCommunity *string                            `xml:"extended-community,omitempty"`
	Interface         *string                            `xml:"interface,omitempty"`
	Origin            *string                            `xml:"origin,omitempty"`
	Metric            *int64                             `xml:"metric,omitempty"`
	Tag               *int64                             `xml:"tag,omitempty"`
	LocalPreference   *int64                             `xml:"local-preference,omitempty"`
	Peer              *string                            `xml:"peer,omitempty"`
	Ipv4              *bgpRipRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Misc              []generic.Xml                      `xml:",any"`
	MiscAttributes    []xml.Attr                         `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *bgpRipRouteMapMatchIpv4AddressXml_11_0_2     `xml:"address,omitempty"`
	NextHop        *bgpRipRouteMapMatchIpv4NextHopXml_11_0_2     `xml:"next-hop,omitempty"`
	RouteSource    *bgpRipRouteMapMatchIpv4RouteSourceXml_11_0_2 `xml:"route-source,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapMatchIpv4RouteSourceXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bgpRipRouteMapSetXml_11_0_2 struct {
	Metric         *bgpRipRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	NextHop        *string                            `xml:"next-hop,omitempty"`
	Tag            *int64                             `xml:"tag,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type bgpRipRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticXml_11_0_2 struct {
	Bgp            *connectedStaticBgpXml_11_0_2    `xml:"bgp,omitempty"`
	Ospf           *connectedStaticOspfXml_11_0_2   `xml:"ospf,omitempty"`
	Ospfv3         *connectedStaticOspfv3Xml_11_0_2 `xml:"ospfv3,omitempty"`
	Rib            *connectedStaticRibXml_11_0_2    `xml:"rib,omitempty"`
	Rip            *connectedStaticRipXml_11_0_2    `xml:"rip,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type connectedStaticBgpXml_11_0_2 struct {
	RouteMap       *connectedStaticBgpRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticBgpRouteMapContainerXml_11_0_2 struct {
	Entries []connectedStaticBgpRouteMapXml_11_0_2 `xml:"entry"`
}
type connectedStaticBgpRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Action         *string                                    `xml:"action,omitempty"`
	Description    *string                                    `xml:"description,omitempty"`
	Match          *connectedStaticBgpRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *connectedStaticBgpRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                        `xml:"interface,omitempty"`
	Metric         *int64                                         `xml:"metric,omitempty"`
	Ipv4           *connectedStaticBgpRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6           *connectedStaticBgpRouteMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *connectedStaticBgpRouteMapMatchIpv4AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticBgpRouteMapMatchIpv4NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6Xml_11_0_2 struct {
	Address        *connectedStaticBgpRouteMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticBgpRouteMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetXml_11_0_2 struct {
	Aggregator        *connectedStaticBgpRouteMapSetAggregatorXml_11_0_2 `xml:"aggregator,omitempty"`
	Metric            *connectedStaticBgpRouteMapSetMetricXml_11_0_2     `xml:"metric,omitempty"`
	Ipv4              *connectedStaticBgpRouteMapSetIpv4Xml_11_0_2       `xml:"ipv4,omitempty"`
	Ipv6              *connectedStaticBgpRouteMapSetIpv6Xml_11_0_2       `xml:"ipv6,omitempty"`
	Tag               *int64                                             `xml:"tag,omitempty"`
	LocalPreference   *int64                                             `xml:"local-preference,omitempty"`
	Weight            *int64                                             `xml:"weight,omitempty"`
	Origin            *string                                            `xml:"origin,omitempty"`
	AtomicAggregate   *string                                            `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                                            `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                                   `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                                   `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                                   `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                                   `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                      `xml:",any"`
	MiscAttributes    []xml.Attr                                         `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetAggregatorXml_11_0_2 struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetIpv4Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticBgpRouteMapSetIpv6Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfXml_11_0_2 struct {
	RouteMap       *connectedStaticOspfRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type connectedStaticOspfRouteMapContainerXml_11_0_2 struct {
	Entries []connectedStaticOspfRouteMapXml_11_0_2 `xml:"entry"`
}
type connectedStaticOspfRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                                    `xml:"entry"`
	Name           string                                      `xml:"name,attr"`
	Action         *string                                     `xml:"action,omitempty"`
	Description    *string                                     `xml:"description,omitempty"`
	Match          *connectedStaticOspfRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *connectedStaticOspfRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                         `xml:"interface,omitempty"`
	Metric         *int64                                          `xml:"metric,omitempty"`
	Ipv4           *connectedStaticOspfRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *connectedStaticOspfRouteMapMatchIpv4AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticOspfRouteMapMatchIpv4NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                          `xml:",any"`
	MiscAttributes []xml.Attr                                             `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfRouteMapSetXml_11_0_2 struct {
	Metric         *connectedStaticOspfRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	MetricType     *string                                         `xml:"metric-type,omitempty"`
	Tag            *int64                                          `xml:"tag,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type connectedStaticOspfRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3Xml_11_0_2 struct {
	RouteMap       *connectedStaticOspfv3RouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapContainerXml_11_0_2 struct {
	Entries []connectedStaticOspfv3RouteMapXml_11_0_2 `xml:"entry"`
}
type connectedStaticOspfv3RouteMapXml_11_0_2 struct {
	XMLName        xml.Name                                      `xml:"entry"`
	Name           string                                        `xml:"name,attr"`
	Action         *string                                       `xml:"action,omitempty"`
	Description    *string                                       `xml:"description,omitempty"`
	Match          *connectedStaticOspfv3RouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *connectedStaticOspfv3RouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchXml_11_0_2 struct {
	Interface      *string                                           `xml:"interface,omitempty"`
	Metric         *int64                                            `xml:"metric,omitempty"`
	Ipv6           *connectedStaticOspfv3RouteMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6Xml_11_0_2 struct {
	Address        *connectedStaticOspfv3RouteMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticOspfv3RouteMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapSetXml_11_0_2 struct {
	Metric         *connectedStaticOspfv3RouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	MetricType     *string                                           `xml:"metric-type,omitempty"`
	Tag            *int64                                            `xml:"tag,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type connectedStaticOspfv3RouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibXml_11_0_2 struct {
	RouteMap       *connectedStaticRibRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRibRouteMapContainerXml_11_0_2 struct {
	Entries []connectedStaticRibRouteMapXml_11_0_2 `xml:"entry"`
}
type connectedStaticRibRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Action         *string                                    `xml:"action,omitempty"`
	Description    *string                                    `xml:"description,omitempty"`
	Match          *connectedStaticRibRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *connectedStaticRibRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                        `xml:"interface,omitempty"`
	Metric         *int64                                         `xml:"metric,omitempty"`
	Ipv4           *connectedStaticRibRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6           *connectedStaticRibRouteMapMatchIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *connectedStaticRibRouteMapMatchIpv4AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticRibRouteMapMatchIpv4NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6Xml_11_0_2 struct {
	Address        *connectedStaticRibRouteMapMatchIpv6AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticRibRouteMapMatchIpv6NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapMatchIpv6NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRibRouteMapSetXml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipXml_11_0_2 struct {
	RouteMap       *connectedStaticRipRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRipRouteMapContainerXml_11_0_2 struct {
	Entries []connectedStaticRipRouteMapXml_11_0_2 `xml:"entry"`
}
type connectedStaticRipRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Action         *string                                    `xml:"action,omitempty"`
	Description    *string                                    `xml:"description,omitempty"`
	Match          *connectedStaticRipRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *connectedStaticRipRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                        `xml:"interface,omitempty"`
	Metric         *int64                                         `xml:"metric,omitempty"`
	Ipv4           *connectedStaticRipRouteMapMatchIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4Xml_11_0_2 struct {
	Address        *connectedStaticRipRouteMapMatchIpv4AddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *connectedStaticRipRouteMapMatchIpv4NextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4AddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipRouteMapMatchIpv4NextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type connectedStaticRipRouteMapSetXml_11_0_2 struct {
	Metric         *connectedStaticRipRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	NextHop        *string                                        `xml:"next-hop,omitempty"`
	Tag            *int64                                         `xml:"tag,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type connectedStaticRipRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfXml_11_0_2 struct {
	Bgp            *ospfBgpXml_11_0_2 `xml:"bgp,omitempty"`
	Rib            *ospfRibXml_11_0_2 `xml:"rib,omitempty"`
	Rip            *ospfRipXml_11_0_2 `xml:"rip,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
	MiscAttributes []xml.Attr         `xml:",any,attr"`
}
type ospfBgpXml_11_0_2 struct {
	RouteMap       *ospfBgpRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ospfBgpRouteMapContainerXml_11_0_2 struct {
	Entries []ospfBgpRouteMapXml_11_0_2 `xml:"entry"`
}
type ospfBgpRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Action         *string                         `xml:"action,omitempty"`
	Description    *string                         `xml:"description,omitempty"`
	Match          *ospfBgpRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ospfBgpRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfBgpRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                `xml:"interface,omitempty"`
	Metric         *int64                                 `xml:"metric,omitempty"`
	Tag            *int64                                 `xml:"tag,omitempty"`
	Address        *ospfBgpRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ospfBgpRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ospfBgpRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetXml_11_0_2 struct {
	Aggregator        *ospfBgpRouteMapSetAggregatorXml_11_0_2 `xml:"aggregator,omitempty"`
	Metric            *ospfBgpRouteMapSetMetricXml_11_0_2     `xml:"metric,omitempty"`
	Ipv4              *ospfBgpRouteMapSetIpv4Xml_11_0_2       `xml:"ipv4,omitempty"`
	Tag               *int64                                  `xml:"tag,omitempty"`
	LocalPreference   *int64                                  `xml:"local-preference,omitempty"`
	Weight            *int64                                  `xml:"weight,omitempty"`
	Origin            *string                                 `xml:"origin,omitempty"`
	AtomicAggregate   *string                                 `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                                 `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                        `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                        `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                        `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                        `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                           `xml:",any"`
	MiscAttributes    []xml.Attr                              `xml:",any,attr"`
}
type ospfBgpRouteMapSetAggregatorXml_11_0_2 struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfBgpRouteMapSetIpv4Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibXml_11_0_2 struct {
	RouteMap       *ospfRibRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ospfRibRouteMapContainerXml_11_0_2 struct {
	Entries []ospfRibRouteMapXml_11_0_2 `xml:"entry"`
}
type ospfRibRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Action         *string                         `xml:"action,omitempty"`
	Description    *string                         `xml:"description,omitempty"`
	Match          *ospfRibRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ospfRibRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfRibRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                `xml:"interface,omitempty"`
	Metric         *int64                                 `xml:"metric,omitempty"`
	Tag            *int64                                 `xml:"tag,omitempty"`
	Address        *ospfRibRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ospfRibRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ospfRibRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRibRouteMapSetXml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipXml_11_0_2 struct {
	RouteMap       *ospfRipRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ospfRipRouteMapContainerXml_11_0_2 struct {
	Entries []ospfRipRouteMapXml_11_0_2 `xml:"entry"`
}
type ospfRipRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Action         *string                         `xml:"action,omitempty"`
	Description    *string                         `xml:"description,omitempty"`
	Match          *ospfRipRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ospfRipRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ospfRipRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                `xml:"interface,omitempty"`
	Metric         *int64                                 `xml:"metric,omitempty"`
	Tag            *int64                                 `xml:"tag,omitempty"`
	Address        *ospfRipRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ospfRipRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ospfRipRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfRipRouteMapSetXml_11_0_2 struct {
	Metric         *ospfRipRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	NextHop        *string                             `xml:"next-hop,omitempty"`
	Tag            *int64                              `xml:"tag,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ospfRipRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3Xml_11_0_2 struct {
	Bgp            *ospfv3BgpXml_11_0_2 `xml:"bgp,omitempty"`
	Rib            *ospfv3RibXml_11_0_2 `xml:"rib,omitempty"`
	Misc           []generic.Xml        `xml:",any"`
	MiscAttributes []xml.Attr           `xml:",any,attr"`
}
type ospfv3BgpXml_11_0_2 struct {
	RouteMap       *ospfv3BgpRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type ospfv3BgpRouteMapContainerXml_11_0_2 struct {
	Entries []ospfv3BgpRouteMapXml_11_0_2 `xml:"entry"`
}
type ospfv3BgpRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                          `xml:"entry"`
	Name           string                            `xml:"name,attr"`
	Action         *string                           `xml:"action,omitempty"`
	Description    *string                           `xml:"description,omitempty"`
	Match          *ospfv3BgpRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ospfv3BgpRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                  `xml:"interface,omitempty"`
	Metric         *int64                                   `xml:"metric,omitempty"`
	Tag            *int64                                   `xml:"tag,omitempty"`
	Address        *ospfv3BgpRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ospfv3BgpRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchAddressXml_11_0_2 struct {
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	AccessList     *string       `xml:"access-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetXml_11_0_2 struct {
	Aggregator        *ospfv3BgpRouteMapSetAggregatorXml_11_0_2 `xml:"aggregator,omitempty"`
	Metric            *ospfv3BgpRouteMapSetMetricXml_11_0_2     `xml:"metric,omitempty"`
	Ipv6              *ospfv3BgpRouteMapSetIpv6Xml_11_0_2       `xml:"ipv6,omitempty"`
	Tag               *int64                                    `xml:"tag,omitempty"`
	LocalPreference   *int64                                    `xml:"local-preference,omitempty"`
	Weight            *int64                                    `xml:"weight,omitempty"`
	Origin            *string                                   `xml:"origin,omitempty"`
	AtomicAggregate   *string                                   `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                                   `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                          `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                          `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                          `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                          `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                             `xml:",any"`
	MiscAttributes    []xml.Attr                                `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetAggregatorXml_11_0_2 struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3BgpRouteMapSetIpv6Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibXml_11_0_2 struct {
	RouteMap       *ospfv3RibRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type ospfv3RibRouteMapContainerXml_11_0_2 struct {
	Entries []ospfv3RibRouteMapXml_11_0_2 `xml:"entry"`
}
type ospfv3RibRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                          `xml:"entry"`
	Name           string                            `xml:"name,attr"`
	Action         *string                           `xml:"action,omitempty"`
	Description    *string                           `xml:"description,omitempty"`
	Match          *ospfv3RibRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ospfv3RibRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                  `xml:"interface,omitempty"`
	Metric         *int64                                   `xml:"metric,omitempty"`
	Tag            *int64                                   `xml:"tag,omitempty"`
	Address        *ospfv3RibRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ospfv3RibRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchAddressXml_11_0_2 struct {
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	AccessList     *string       `xml:"access-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ospfv3RibRouteMapSetXml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripXml_11_0_2 struct {
	Bgp            *ripBgpXml_11_0_2  `xml:"bgp,omitempty"`
	Ospf           *ripOspfXml_11_0_2 `xml:"ospf,omitempty"`
	Rib            *ripRibXml_11_0_2  `xml:"rib,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
	MiscAttributes []xml.Attr         `xml:",any,attr"`
}
type ripBgpXml_11_0_2 struct {
	RouteMap       *ripBgpRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type ripBgpRouteMapContainerXml_11_0_2 struct {
	Entries []ripBgpRouteMapXml_11_0_2 `xml:"entry"`
}
type ripBgpRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Action         *string                        `xml:"action,omitempty"`
	Description    *string                        `xml:"description,omitempty"`
	Match          *ripBgpRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ripBgpRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ripBgpRouteMapMatchXml_11_0_2 struct {
	Interface      *string                               `xml:"interface,omitempty"`
	Metric         *int64                                `xml:"metric,omitempty"`
	Tag            *int64                                `xml:"tag,omitempty"`
	Address        *ripBgpRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ripBgpRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type ripBgpRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetXml_11_0_2 struct {
	Aggregator        *ripBgpRouteMapSetAggregatorXml_11_0_2 `xml:"aggregator,omitempty"`
	Metric            *ripBgpRouteMapSetMetricXml_11_0_2     `xml:"metric,omitempty"`
	Ipv4              *ripBgpRouteMapSetIpv4Xml_11_0_2       `xml:"ipv4,omitempty"`
	Tag               *int64                                 `xml:"tag,omitempty"`
	LocalPreference   *int64                                 `xml:"local-preference,omitempty"`
	Weight            *int64                                 `xml:"weight,omitempty"`
	Origin            *string                                `xml:"origin,omitempty"`
	AtomicAggregate   *string                                `xml:"atomic-aggregate,omitempty"`
	OriginatorId      *string                                `xml:"originator-id,omitempty"`
	AspathPrepend     *util.MemberType                       `xml:"aspath-prepend,omitempty"`
	RegularCommunity  *util.MemberType                       `xml:"regular-community,omitempty"`
	LargeCommunity    *util.MemberType                       `xml:"large-community,omitempty"`
	ExtendedCommunity *util.MemberType                       `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                          `xml:",any"`
	MiscAttributes    []xml.Attr                             `xml:",any,attr"`
}
type ripBgpRouteMapSetAggregatorXml_11_0_2 struct {
	As             *int64        `xml:"as,omitempty"`
	RouterId       *string       `xml:"router-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripBgpRouteMapSetIpv4Xml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	NextHop        *string       `xml:"next-hop,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfXml_11_0_2 struct {
	RouteMap       *ripOspfRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ripOspfRouteMapContainerXml_11_0_2 struct {
	Entries []ripOspfRouteMapXml_11_0_2 `xml:"entry"`
}
type ripOspfRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Action         *string                         `xml:"action,omitempty"`
	Description    *string                         `xml:"description,omitempty"`
	Match          *ripOspfRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ripOspfRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ripOspfRouteMapMatchXml_11_0_2 struct {
	Interface      *string                                `xml:"interface,omitempty"`
	Metric         *int64                                 `xml:"metric,omitempty"`
	Tag            *int64                                 `xml:"tag,omitempty"`
	Address        *ripOspfRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ripOspfRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ripOspfRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripOspfRouteMapSetXml_11_0_2 struct {
	Metric         *ripOspfRouteMapSetMetricXml_11_0_2 `xml:"metric,omitempty"`
	MetricType     *string                             `xml:"metric-type,omitempty"`
	Tag            *int64                              `xml:"tag,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ripOspfRouteMapSetMetricXml_11_0_2 struct {
	Value          *int64        `xml:"value,omitempty"`
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibXml_11_0_2 struct {
	RouteMap       *ripRibRouteMapContainerXml_11_0_2 `xml:"route-map,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type ripRibRouteMapContainerXml_11_0_2 struct {
	Entries []ripRibRouteMapXml_11_0_2 `xml:"entry"`
}
type ripRibRouteMapXml_11_0_2 struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Action         *string                        `xml:"action,omitempty"`
	Description    *string                        `xml:"description,omitempty"`
	Match          *ripRibRouteMapMatchXml_11_0_2 `xml:"match,omitempty"`
	Set            *ripRibRouteMapSetXml_11_0_2   `xml:"set,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ripRibRouteMapMatchXml_11_0_2 struct {
	Interface      *string                               `xml:"interface,omitempty"`
	Metric         *int64                                `xml:"metric,omitempty"`
	Tag            *int64                                `xml:"tag,omitempty"`
	Address        *ripRibRouteMapMatchAddressXml_11_0_2 `xml:"address,omitempty"`
	NextHop        *ripRibRouteMapMatchNextHopXml_11_0_2 `xml:"next-hop,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type ripRibRouteMapMatchAddressXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibRouteMapMatchNextHopXml_11_0_2 struct {
	AccessList     *string       `xml:"access-list,omitempty"`
	PrefixList     *string       `xml:"prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ripRibRouteMapSetXml_11_0_2 struct {
	SourceAddress  *string       `xml:"source-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Bgp != nil {
		var obj bgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.ConnectedStatic != nil {
		var obj connectedStaticXml
		obj.MarshalFromObject(*s.ConnectedStatic)
		o.ConnectedStatic = &obj
	}
	if s.Ospf != nil {
		var obj ospfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj ospfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rip != nil {
		var obj ripXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
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
	var connectedStaticVal *ConnectedStatic
	if o.ConnectedStatic != nil {
		obj, err := o.ConnectedStatic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectedStaticVal = obj
	}
	var ospfVal *Ospf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *Ospfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ripVal *Rip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		Bgp:             bgpVal,
		ConnectedStatic: connectedStaticVal,
		Ospf:            ospfVal,
		Ospfv3:          ospfv3Val,
		Rip:             ripVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpXml) MarshalFromObject(s Bgp) {
	if s.Ospf != nil {
		var obj bgpOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj bgpOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rib != nil {
		var obj bgpRibXml
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj bgpRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpXml) UnmarshalToObject() (*Bgp, error) {
	var ospfVal *BgpOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *BgpOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ribVal *BgpRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *BgpRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Bgp{
		Ospf:           ospfVal,
		Ospfv3:         ospfv3Val,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfXml) MarshalFromObject(s BgpOspf) {
	if s.RouteMap != nil {
		var objs []bgpOspfRouteMapXml
		for _, elt := range s.RouteMap {
			var obj bgpOspfRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpOspfRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfXml) UnmarshalToObject() (*BgpOspf, error) {
	var routeMapVal []BgpOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapXml) MarshalFromObject(s BgpOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpOspfRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpOspfRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapXml) UnmarshalToObject() (*BgpOspfRouteMap, error) {
	var matchVal *BgpOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpOspfRouteMap{
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
func (o *bgpOspfRouteMapMatchXml) MarshalFromObject(s BgpOspfRouteMapMatch) {
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
		var obj bgpOspfRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchXml) UnmarshalToObject() (*BgpOspfRouteMapMatch, error) {
	var ipv4Val *BgpOspfRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &BgpOspfRouteMapMatch{
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
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4Xml) MarshalFromObject(s BgpOspfRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpOspfRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpOspfRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpOspfRouteMapMatchIpv4RouteSourceXml
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4Xml) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4, error) {
	var addressVal *BgpOspfRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpOspfRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpOspfRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpOspfRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4AddressXml) MarshalFromObject(s BgpOspfRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4Address, error) {

	result := &BgpOspfRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4NextHopXml) MarshalFromObject(s BgpOspfRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4NextHop, error) {

	result := &BgpOspfRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4RouteSourceXml) MarshalFromObject(s BgpOspfRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4RouteSourceXml) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4RouteSource, error) {

	result := &BgpOspfRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapSetXml) MarshalFromObject(s BgpOspfRouteMapSet) {
	if s.Metric != nil {
		var obj bgpOspfRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapSetXml) UnmarshalToObject() (*BgpOspfRouteMapSet, error) {
	var metricVal *BgpOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapSetMetricXml) MarshalFromObject(s BgpOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapSetMetricXml) UnmarshalToObject() (*BgpOspfRouteMapSetMetric, error) {

	result := &BgpOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3Xml) MarshalFromObject(s BgpOspfv3) {
	if s.RouteMap != nil {
		var objs []bgpOspfv3RouteMapXml
		for _, elt := range s.RouteMap {
			var obj bgpOspfv3RouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpOspfv3RouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3Xml) UnmarshalToObject() (*BgpOspfv3, error) {
	var routeMapVal []BgpOspfv3RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpOspfv3{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapXml) MarshalFromObject(s BgpOspfv3RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpOspfv3RouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpOspfv3RouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapXml) UnmarshalToObject() (*BgpOspfv3RouteMap, error) {
	var matchVal *BgpOspfv3RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpOspfv3RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpOspfv3RouteMap{
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
func (o *bgpOspfv3RouteMapMatchXml) MarshalFromObject(s BgpOspfv3RouteMapMatch) {
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
	if s.Ipv6 != nil {
		var obj bgpOspfv3RouteMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchXml) UnmarshalToObject() (*BgpOspfv3RouteMapMatch, error) {
	var ipv6Val *BgpOspfv3RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &BgpOspfv3RouteMapMatch{
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
		Ipv6:              ipv6Val,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6Xml) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj bgpOspfv3RouteMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpOspfv3RouteMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6Xml) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6, error) {
	var addressVal *BgpOspfv3RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpOspfv3RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &BgpOspfv3RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6AddressXml) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6AddressXml) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6Address, error) {

	result := &BgpOspfv3RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6NextHopXml) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6NextHopXml) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6NextHop, error) {

	result := &BgpOspfv3RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapSetXml) MarshalFromObject(s BgpOspfv3RouteMapSet) {
	if s.Metric != nil {
		var obj bgpOspfv3RouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapSetXml) UnmarshalToObject() (*BgpOspfv3RouteMapSet, error) {
	var metricVal *BgpOspfv3RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpOspfv3RouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapSetMetricXml) MarshalFromObject(s BgpOspfv3RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapSetMetricXml) UnmarshalToObject() (*BgpOspfv3RouteMapSetMetric, error) {

	result := &BgpOspfv3RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibXml) MarshalFromObject(s BgpRib) {
	if s.RouteMap != nil {
		var objs []bgpRibRouteMapXml
		for _, elt := range s.RouteMap {
			var obj bgpRibRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpRibRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibXml) UnmarshalToObject() (*BgpRib, error) {
	var routeMapVal []BgpRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapXml) MarshalFromObject(s BgpRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpRibRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpRibRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapXml) UnmarshalToObject() (*BgpRibRouteMap, error) {
	var matchVal *BgpRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpRibRouteMap{
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
func (o *bgpRibRouteMapMatchXml) MarshalFromObject(s BgpRibRouteMapMatch) {
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
		var obj bgpRibRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj bgpRibRouteMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchXml) UnmarshalToObject() (*BgpRibRouteMapMatch, error) {
	var ipv4Val *BgpRibRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *BgpRibRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &BgpRibRouteMapMatch{
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
func (o *bgpRibRouteMapMatchIpv4Xml) MarshalFromObject(s BgpRibRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpRibRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRibRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpRibRouteMapMatchIpv4RouteSourceXml
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4Xml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4, error) {
	var addressVal *BgpRibRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRibRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpRibRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpRibRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4AddressXml) MarshalFromObject(s BgpRibRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4Address, error) {

	result := &BgpRibRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4NextHopXml) MarshalFromObject(s BgpRibRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4NextHop, error) {

	result := &BgpRibRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4RouteSourceXml) MarshalFromObject(s BgpRibRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4RouteSourceXml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4RouteSource, error) {

	result := &BgpRibRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6Xml) MarshalFromObject(s BgpRibRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj bgpRibRouteMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRibRouteMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6Xml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6, error) {
	var addressVal *BgpRibRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRibRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &BgpRibRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6AddressXml) MarshalFromObject(s BgpRibRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6AddressXml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6Address, error) {

	result := &BgpRibRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6NextHopXml) MarshalFromObject(s BgpRibRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6NextHopXml) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6NextHop, error) {

	result := &BgpRibRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapSetXml) MarshalFromObject(s BgpRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapSetXml) UnmarshalToObject() (*BgpRibRouteMapSet, error) {

	result := &BgpRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipXml) MarshalFromObject(s BgpRip) {
	if s.RouteMap != nil {
		var objs []bgpRipRouteMapXml
		for _, elt := range s.RouteMap {
			var obj bgpRipRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpRipRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipXml) UnmarshalToObject() (*BgpRip, error) {
	var routeMapVal []BgpRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapXml) MarshalFromObject(s BgpRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpRipRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpRipRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapXml) UnmarshalToObject() (*BgpRipRouteMap, error) {
	var matchVal *BgpRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpRipRouteMap{
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
func (o *bgpRipRouteMapMatchXml) MarshalFromObject(s BgpRipRouteMapMatch) {
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
		var obj bgpRipRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchXml) UnmarshalToObject() (*BgpRipRouteMapMatch, error) {
	var ipv4Val *BgpRipRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &BgpRipRouteMapMatch{
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
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4Xml) MarshalFromObject(s BgpRipRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpRipRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRipRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpRipRouteMapMatchIpv4RouteSourceXml
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4Xml) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4, error) {
	var addressVal *BgpRipRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRipRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpRipRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpRipRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4AddressXml) MarshalFromObject(s BgpRipRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4Address, error) {

	result := &BgpRipRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4NextHopXml) MarshalFromObject(s BgpRipRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4NextHop, error) {

	result := &BgpRipRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4RouteSourceXml) MarshalFromObject(s BgpRipRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4RouteSourceXml) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4RouteSource, error) {

	result := &BgpRipRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapSetXml) MarshalFromObject(s BgpRipRouteMapSet) {
	if s.Metric != nil {
		var obj bgpRipRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapSetXml) UnmarshalToObject() (*BgpRipRouteMapSet, error) {
	var metricVal *BgpRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapSetMetricXml) MarshalFromObject(s BgpRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapSetMetricXml) UnmarshalToObject() (*BgpRipRouteMapSetMetric, error) {

	result := &BgpRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticXml) MarshalFromObject(s ConnectedStatic) {
	if s.Bgp != nil {
		var obj connectedStaticBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj connectedStaticOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj connectedStaticOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rib != nil {
		var obj connectedStaticRibXml
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj connectedStaticRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticXml) UnmarshalToObject() (*ConnectedStatic, error) {
	var bgpVal *ConnectedStaticBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *ConnectedStaticOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *ConnectedStaticOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ribVal *ConnectedStaticRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *ConnectedStaticRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &ConnectedStatic{
		Bgp:            bgpVal,
		Ospf:           ospfVal,
		Ospfv3:         ospfv3Val,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpXml) MarshalFromObject(s ConnectedStaticBgp) {
	if s.RouteMap != nil {
		var objs []connectedStaticBgpRouteMapXml
		for _, elt := range s.RouteMap {
			var obj connectedStaticBgpRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticBgpRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpXml) UnmarshalToObject() (*ConnectedStaticBgp, error) {
	var routeMapVal []ConnectedStaticBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapXml) MarshalFromObject(s ConnectedStaticBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticBgpRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticBgpRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMap, error) {
	var matchVal *ConnectedStaticBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticBgpRouteMap{
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
func (o *connectedStaticBgpRouteMapMatchXml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticBgpRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticBgpRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4Xml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4Xml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticBgpRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticBgpRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticBgpRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4AddressXml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4NextHopXml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6Xml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6Xml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticBgpRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticBgpRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticBgpRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6AddressXml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6AddressXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6NextHopXml) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6NextHopXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetXml) MarshalFromObject(s ConnectedStaticBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj connectedStaticBgpRouteMapSetAggregatorXml
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj connectedStaticBgpRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj connectedStaticBgpRouteMapSetIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticBgpRouteMapSetIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSet, error) {
	var aggregatorVal *ConnectedStaticBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *ConnectedStaticBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *ConnectedStaticBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticBgpRouteMapSetIpv6
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &ConnectedStaticBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Ipv6:              ipv6Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetAggregatorXml) MarshalFromObject(s ConnectedStaticBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetAggregatorXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetAggregator, error) {

	result := &ConnectedStaticBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetMetricXml) MarshalFromObject(s ConnectedStaticBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetMetricXml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetMetric, error) {

	result := &ConnectedStaticBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetIpv4Xml) MarshalFromObject(s ConnectedStaticBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetIpv4Xml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetIpv4, error) {

	result := &ConnectedStaticBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetIpv6Xml) MarshalFromObject(s ConnectedStaticBgpRouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetIpv6Xml) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetIpv6, error) {

	result := &ConnectedStaticBgpRouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfXml) MarshalFromObject(s ConnectedStaticOspf) {
	if s.RouteMap != nil {
		var objs []connectedStaticOspfRouteMapXml
		for _, elt := range s.RouteMap {
			var obj connectedStaticOspfRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticOspfRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfXml) UnmarshalToObject() (*ConnectedStaticOspf, error) {
	var routeMapVal []ConnectedStaticOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapXml) MarshalFromObject(s ConnectedStaticOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticOspfRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticOspfRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMap, error) {
	var matchVal *ConnectedStaticOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticOspfRouteMap{
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
func (o *connectedStaticOspfRouteMapMatchXml) MarshalFromObject(s ConnectedStaticOspfRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticOspfRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &ConnectedStaticOspfRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4Xml) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4Xml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticOspfRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticOspfRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticOspfRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4AddressXml) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticOspfRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4NextHopXml) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticOspfRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapSetXml) MarshalFromObject(s ConnectedStaticOspfRouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticOspfRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapSetXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapSet, error) {
	var metricVal *ConnectedStaticOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapSetMetricXml) MarshalFromObject(s ConnectedStaticOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapSetMetricXml) UnmarshalToObject() (*ConnectedStaticOspfRouteMapSetMetric, error) {

	result := &ConnectedStaticOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3Xml) MarshalFromObject(s ConnectedStaticOspfv3) {
	if s.RouteMap != nil {
		var objs []connectedStaticOspfv3RouteMapXml
		for _, elt := range s.RouteMap {
			var obj connectedStaticOspfv3RouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticOspfv3RouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3Xml) UnmarshalToObject() (*ConnectedStaticOspfv3, error) {
	var routeMapVal []ConnectedStaticOspfv3RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticOspfv3{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticOspfv3RouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticOspfv3RouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMap, error) {
	var matchVal *ConnectedStaticOspfv3RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticOspfv3RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMap{
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
func (o *connectedStaticOspfv3RouteMapMatchXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv6 != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatch, error) {
	var ipv6Val *ConnectedStaticOspfv3RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticOspfv3RouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6Xml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6Xml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticOspfv3RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticOspfv3RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6AddressXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6AddressXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6NextHopXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6NextHopXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapSetXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticOspfv3RouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapSetXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapSet, error) {
	var metricVal *ConnectedStaticOspfv3RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapSetMetricXml) MarshalFromObject(s ConnectedStaticOspfv3RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapSetMetricXml) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapSetMetric, error) {

	result := &ConnectedStaticOspfv3RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibXml) MarshalFromObject(s ConnectedStaticRib) {
	if s.RouteMap != nil {
		var objs []connectedStaticRibRouteMapXml
		for _, elt := range s.RouteMap {
			var obj connectedStaticRibRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticRibRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibXml) UnmarshalToObject() (*ConnectedStaticRib, error) {
	var routeMapVal []ConnectedStaticRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapXml) MarshalFromObject(s ConnectedStaticRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticRibRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticRibRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapXml) UnmarshalToObject() (*ConnectedStaticRibRouteMap, error) {
	var matchVal *ConnectedStaticRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticRibRouteMap{
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
func (o *connectedStaticRibRouteMapMatchXml) MarshalFromObject(s ConnectedStaticRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticRibRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticRibRouteMapMatchIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticRibRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticRibRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4Xml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticRibRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRibRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4Xml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticRibRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRibRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRibRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4AddressXml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4NextHopXml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6Xml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticRibRouteMapMatchIpv6AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRibRouteMapMatchIpv6NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6Xml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticRibRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRibRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRibRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6AddressXml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6AddressXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6NextHopXml) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6NextHopXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapSetXml) MarshalFromObject(s ConnectedStaticRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapSetXml) UnmarshalToObject() (*ConnectedStaticRibRouteMapSet, error) {

	result := &ConnectedStaticRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipXml) MarshalFromObject(s ConnectedStaticRip) {
	if s.RouteMap != nil {
		var objs []connectedStaticRipRouteMapXml
		for _, elt := range s.RouteMap {
			var obj connectedStaticRipRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticRipRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipXml) UnmarshalToObject() (*ConnectedStaticRip, error) {
	var routeMapVal []ConnectedStaticRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapXml) MarshalFromObject(s ConnectedStaticRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticRipRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticRipRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapXml) UnmarshalToObject() (*ConnectedStaticRipRouteMap, error) {
	var matchVal *ConnectedStaticRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticRipRouteMap{
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
func (o *connectedStaticRipRouteMapMatchXml) MarshalFromObject(s ConnectedStaticRipRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticRipRouteMapMatchIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchXml) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticRipRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &ConnectedStaticRipRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4Xml) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticRipRouteMapMatchIpv4AddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRipRouteMapMatchIpv4NextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4Xml) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticRipRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRipRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRipRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4AddressXml) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4AddressXml) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticRipRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4NextHopXml) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4NextHopXml) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticRipRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapSetXml) MarshalFromObject(s ConnectedStaticRipRouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticRipRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapSetXml) UnmarshalToObject() (*ConnectedStaticRipRouteMapSet, error) {
	var metricVal *ConnectedStaticRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapSetMetricXml) MarshalFromObject(s ConnectedStaticRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapSetMetricXml) UnmarshalToObject() (*ConnectedStaticRipRouteMapSetMetric, error) {

	result := &ConnectedStaticRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfXml) MarshalFromObject(s Ospf) {
	if s.Bgp != nil {
		var obj ospfBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Rib != nil {
		var obj ospfRibXml
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj ospfRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfXml) UnmarshalToObject() (*Ospf, error) {
	var bgpVal *OspfBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ribVal *OspfRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *OspfRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Ospf{
		Bgp:            bgpVal,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpXml) MarshalFromObject(s OspfBgp) {
	if s.RouteMap != nil {
		var objs []ospfBgpRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ospfBgpRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfBgpRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpXml) UnmarshalToObject() (*OspfBgp, error) {
	var routeMapVal []OspfBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapXml) MarshalFromObject(s OspfBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfBgpRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfBgpRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapXml) UnmarshalToObject() (*OspfBgpRouteMap, error) {
	var matchVal *OspfBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfBgpRouteMap{
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
func (o *ospfBgpRouteMapMatchXml) MarshalFromObject(s OspfBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfBgpRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfBgpRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchXml) UnmarshalToObject() (*OspfBgpRouteMapMatch, error) {
	var addressVal *OspfBgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfBgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapMatchAddressXml) MarshalFromObject(s OspfBgpRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchAddressXml) UnmarshalToObject() (*OspfBgpRouteMapMatchAddress, error) {

	result := &OspfBgpRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapMatchNextHopXml) MarshalFromObject(s OspfBgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchNextHopXml) UnmarshalToObject() (*OspfBgpRouteMapMatchNextHop, error) {

	result := &OspfBgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetXml) MarshalFromObject(s OspfBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ospfBgpRouteMapSetAggregatorXml
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ospfBgpRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj ospfBgpRouteMapSetIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetXml) UnmarshalToObject() (*OspfBgpRouteMapSet, error) {
	var aggregatorVal *OspfBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *OspfBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *OspfBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &OspfBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetAggregatorXml) MarshalFromObject(s OspfBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetAggregatorXml) UnmarshalToObject() (*OspfBgpRouteMapSetAggregator, error) {

	result := &OspfBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetMetricXml) MarshalFromObject(s OspfBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetMetricXml) UnmarshalToObject() (*OspfBgpRouteMapSetMetric, error) {

	result := &OspfBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetIpv4Xml) MarshalFromObject(s OspfBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetIpv4Xml) UnmarshalToObject() (*OspfBgpRouteMapSetIpv4, error) {

	result := &OspfBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibXml) MarshalFromObject(s OspfRib) {
	if s.RouteMap != nil {
		var objs []ospfRibRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ospfRibRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfRibRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibXml) UnmarshalToObject() (*OspfRib, error) {
	var routeMapVal []OspfRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapXml) MarshalFromObject(s OspfRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfRibRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfRibRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapXml) UnmarshalToObject() (*OspfRibRouteMap, error) {
	var matchVal *OspfRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfRibRouteMap{
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
func (o *ospfRibRouteMapMatchXml) MarshalFromObject(s OspfRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfRibRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfRibRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchXml) UnmarshalToObject() (*OspfRibRouteMapMatch, error) {
	var addressVal *OspfRibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfRibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapMatchAddressXml) MarshalFromObject(s OspfRibRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchAddressXml) UnmarshalToObject() (*OspfRibRouteMapMatchAddress, error) {

	result := &OspfRibRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapMatchNextHopXml) MarshalFromObject(s OspfRibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchNextHopXml) UnmarshalToObject() (*OspfRibRouteMapMatchNextHop, error) {

	result := &OspfRibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapSetXml) MarshalFromObject(s OspfRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapSetXml) UnmarshalToObject() (*OspfRibRouteMapSet, error) {

	result := &OspfRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipXml) MarshalFromObject(s OspfRip) {
	if s.RouteMap != nil {
		var objs []ospfRipRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ospfRipRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfRipRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipXml) UnmarshalToObject() (*OspfRip, error) {
	var routeMapVal []OspfRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapXml) MarshalFromObject(s OspfRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfRipRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfRipRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapXml) UnmarshalToObject() (*OspfRipRouteMap, error) {
	var matchVal *OspfRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfRipRouteMap{
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
func (o *ospfRipRouteMapMatchXml) MarshalFromObject(s OspfRipRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfRipRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfRipRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchXml) UnmarshalToObject() (*OspfRipRouteMapMatch, error) {
	var addressVal *OspfRipRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfRipRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfRipRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapMatchAddressXml) MarshalFromObject(s OspfRipRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchAddressXml) UnmarshalToObject() (*OspfRipRouteMapMatchAddress, error) {

	result := &OspfRipRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapMatchNextHopXml) MarshalFromObject(s OspfRipRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchNextHopXml) UnmarshalToObject() (*OspfRipRouteMapMatchNextHop, error) {

	result := &OspfRipRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapSetXml) MarshalFromObject(s OspfRipRouteMapSet) {
	if s.Metric != nil {
		var obj ospfRipRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapSetXml) UnmarshalToObject() (*OspfRipRouteMapSet, error) {
	var metricVal *OspfRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &OspfRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapSetMetricXml) MarshalFromObject(s OspfRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapSetMetricXml) UnmarshalToObject() (*OspfRipRouteMapSetMetric, error) {

	result := &OspfRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3Xml) MarshalFromObject(s Ospfv3) {
	if s.Bgp != nil {
		var obj ospfv3BgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Rib != nil {
		var obj ospfv3RibXml
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3Xml) UnmarshalToObject() (*Ospfv3, error) {
	var bgpVal *Ospfv3Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ribVal *Ospfv3Rib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}

	result := &Ospfv3{
		Bgp:            bgpVal,
		Rib:            ribVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpXml) MarshalFromObject(s Ospfv3Bgp) {
	if s.RouteMap != nil {
		var objs []ospfv3BgpRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ospfv3BgpRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfv3BgpRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpXml) UnmarshalToObject() (*Ospfv3Bgp, error) {
	var routeMapVal []Ospfv3BgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Ospfv3Bgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapXml) MarshalFromObject(s Ospfv3BgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfv3BgpRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfv3BgpRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapXml) UnmarshalToObject() (*Ospfv3BgpRouteMap, error) {
	var matchVal *Ospfv3BgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *Ospfv3BgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &Ospfv3BgpRouteMap{
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
func (o *ospfv3BgpRouteMapMatchXml) MarshalFromObject(s Ospfv3BgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfv3BgpRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfv3BgpRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchXml) UnmarshalToObject() (*Ospfv3BgpRouteMapMatch, error) {
	var addressVal *Ospfv3BgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *Ospfv3BgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &Ospfv3BgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapMatchAddressXml) MarshalFromObject(s Ospfv3BgpRouteMapMatchAddress) {
	o.PrefixList = s.PrefixList
	o.AccessList = s.AccessList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchAddressXml) UnmarshalToObject() (*Ospfv3BgpRouteMapMatchAddress, error) {

	result := &Ospfv3BgpRouteMapMatchAddress{
		PrefixList:     o.PrefixList,
		AccessList:     o.AccessList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapMatchNextHopXml) MarshalFromObject(s Ospfv3BgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchNextHopXml) UnmarshalToObject() (*Ospfv3BgpRouteMapMatchNextHop, error) {

	result := &Ospfv3BgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetXml) MarshalFromObject(s Ospfv3BgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ospfv3BgpRouteMapSetAggregatorXml
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ospfv3BgpRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv6 != nil {
		var obj ospfv3BgpRouteMapSetIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetXml) UnmarshalToObject() (*Ospfv3BgpRouteMapSet, error) {
	var aggregatorVal *Ospfv3BgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *Ospfv3BgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv6Val *Ospfv3BgpRouteMapSetIpv6
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &Ospfv3BgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv6:              ipv6Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetAggregatorXml) MarshalFromObject(s Ospfv3BgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetAggregatorXml) UnmarshalToObject() (*Ospfv3BgpRouteMapSetAggregator, error) {

	result := &Ospfv3BgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetMetricXml) MarshalFromObject(s Ospfv3BgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetMetricXml) UnmarshalToObject() (*Ospfv3BgpRouteMapSetMetric, error) {

	result := &Ospfv3BgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetIpv6Xml) MarshalFromObject(s Ospfv3BgpRouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetIpv6Xml) UnmarshalToObject() (*Ospfv3BgpRouteMapSetIpv6, error) {

	result := &Ospfv3BgpRouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibXml) MarshalFromObject(s Ospfv3Rib) {
	if s.RouteMap != nil {
		var objs []ospfv3RibRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ospfv3RibRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfv3RibRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibXml) UnmarshalToObject() (*Ospfv3Rib, error) {
	var routeMapVal []Ospfv3RibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Ospfv3Rib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapXml) MarshalFromObject(s Ospfv3RibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfv3RibRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfv3RibRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapXml) UnmarshalToObject() (*Ospfv3RibRouteMap, error) {
	var matchVal *Ospfv3RibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *Ospfv3RibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &Ospfv3RibRouteMap{
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
func (o *ospfv3RibRouteMapMatchXml) MarshalFromObject(s Ospfv3RibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfv3RibRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfv3RibRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchXml) UnmarshalToObject() (*Ospfv3RibRouteMapMatch, error) {
	var addressVal *Ospfv3RibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *Ospfv3RibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &Ospfv3RibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapMatchAddressXml) MarshalFromObject(s Ospfv3RibRouteMapMatchAddress) {
	o.PrefixList = s.PrefixList
	o.AccessList = s.AccessList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchAddressXml) UnmarshalToObject() (*Ospfv3RibRouteMapMatchAddress, error) {

	result := &Ospfv3RibRouteMapMatchAddress{
		PrefixList:     o.PrefixList,
		AccessList:     o.AccessList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapMatchNextHopXml) MarshalFromObject(s Ospfv3RibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchNextHopXml) UnmarshalToObject() (*Ospfv3RibRouteMapMatchNextHop, error) {

	result := &Ospfv3RibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapSetXml) MarshalFromObject(s Ospfv3RibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapSetXml) UnmarshalToObject() (*Ospfv3RibRouteMapSet, error) {

	result := &Ospfv3RibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripXml) MarshalFromObject(s Rip) {
	if s.Bgp != nil {
		var obj ripBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj ripOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Rib != nil {
		var obj ripRibXml
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripXml) UnmarshalToObject() (*Rip, error) {
	var bgpVal *RipBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *RipOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ribVal *RipRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}

	result := &Rip{
		Bgp:            bgpVal,
		Ospf:           ospfVal,
		Rib:            ribVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpXml) MarshalFromObject(s RipBgp) {
	if s.RouteMap != nil {
		var objs []ripBgpRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ripBgpRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripBgpRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpXml) UnmarshalToObject() (*RipBgp, error) {
	var routeMapVal []RipBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapXml) MarshalFromObject(s RipBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripBgpRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripBgpRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapXml) UnmarshalToObject() (*RipBgpRouteMap, error) {
	var matchVal *RipBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipBgpRouteMap{
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
func (o *ripBgpRouteMapMatchXml) MarshalFromObject(s RipBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripBgpRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripBgpRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchXml) UnmarshalToObject() (*RipBgpRouteMapMatch, error) {
	var addressVal *RipBgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipBgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapMatchAddressXml) MarshalFromObject(s RipBgpRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchAddressXml) UnmarshalToObject() (*RipBgpRouteMapMatchAddress, error) {

	result := &RipBgpRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapMatchNextHopXml) MarshalFromObject(s RipBgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchNextHopXml) UnmarshalToObject() (*RipBgpRouteMapMatchNextHop, error) {

	result := &RipBgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetXml) MarshalFromObject(s RipBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ripBgpRouteMapSetAggregatorXml
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ripBgpRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj ripBgpRouteMapSetIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetXml) UnmarshalToObject() (*RipBgpRouteMapSet, error) {
	var aggregatorVal *RipBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *RipBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *RipBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &RipBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetAggregatorXml) MarshalFromObject(s RipBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetAggregatorXml) UnmarshalToObject() (*RipBgpRouteMapSetAggregator, error) {

	result := &RipBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetMetricXml) MarshalFromObject(s RipBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetMetricXml) UnmarshalToObject() (*RipBgpRouteMapSetMetric, error) {

	result := &RipBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetIpv4Xml) MarshalFromObject(s RipBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetIpv4Xml) UnmarshalToObject() (*RipBgpRouteMapSetIpv4, error) {

	result := &RipBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfXml) MarshalFromObject(s RipOspf) {
	if s.RouteMap != nil {
		var objs []ripOspfRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ripOspfRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripOspfRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfXml) UnmarshalToObject() (*RipOspf, error) {
	var routeMapVal []RipOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapXml) MarshalFromObject(s RipOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripOspfRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripOspfRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapXml) UnmarshalToObject() (*RipOspfRouteMap, error) {
	var matchVal *RipOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipOspfRouteMap{
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
func (o *ripOspfRouteMapMatchXml) MarshalFromObject(s RipOspfRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripOspfRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripOspfRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchXml) UnmarshalToObject() (*RipOspfRouteMapMatch, error) {
	var addressVal *RipOspfRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipOspfRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipOspfRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapMatchAddressXml) MarshalFromObject(s RipOspfRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchAddressXml) UnmarshalToObject() (*RipOspfRouteMapMatchAddress, error) {

	result := &RipOspfRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapMatchNextHopXml) MarshalFromObject(s RipOspfRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchNextHopXml) UnmarshalToObject() (*RipOspfRouteMapMatchNextHop, error) {

	result := &RipOspfRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapSetXml) MarshalFromObject(s RipOspfRouteMapSet) {
	if s.Metric != nil {
		var obj ripOspfRouteMapSetMetricXml
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapSetXml) UnmarshalToObject() (*RipOspfRouteMapSet, error) {
	var metricVal *RipOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &RipOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapSetMetricXml) MarshalFromObject(s RipOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapSetMetricXml) UnmarshalToObject() (*RipOspfRouteMapSetMetric, error) {

	result := &RipOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibXml) MarshalFromObject(s RipRib) {
	if s.RouteMap != nil {
		var objs []ripRibRouteMapXml
		for _, elt := range s.RouteMap {
			var obj ripRibRouteMapXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripRibRouteMapContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibXml) UnmarshalToObject() (*RipRib, error) {
	var routeMapVal []RipRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapXml) MarshalFromObject(s RipRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripRibRouteMapMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripRibRouteMapSetXml
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapXml) UnmarshalToObject() (*RipRibRouteMap, error) {
	var matchVal *RipRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipRibRouteMap{
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
func (o *ripRibRouteMapMatchXml) MarshalFromObject(s RipRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripRibRouteMapMatchAddressXml
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripRibRouteMapMatchNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchXml) UnmarshalToObject() (*RipRibRouteMapMatch, error) {
	var addressVal *RipRibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipRibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapMatchAddressXml) MarshalFromObject(s RipRibRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchAddressXml) UnmarshalToObject() (*RipRibRouteMapMatchAddress, error) {

	result := &RipRibRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapMatchNextHopXml) MarshalFromObject(s RipRibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchNextHopXml) UnmarshalToObject() (*RipRibRouteMapMatchNextHop, error) {

	result := &RipRibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapSetXml) MarshalFromObject(s RipRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapSetXml) UnmarshalToObject() (*RipRibRouteMapSet, error) {

	result := &RipRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Bgp != nil {
		var obj bgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.ConnectedStatic != nil {
		var obj connectedStaticXml_11_0_2
		obj.MarshalFromObject(*s.ConnectedStatic)
		o.ConnectedStatic = &obj
	}
	if s.Ospf != nil {
		var obj ospfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj ospfv3Xml_11_0_2
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rip != nil {
		var obj ripXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var bgpVal *Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var connectedStaticVal *ConnectedStatic
	if o.ConnectedStatic != nil {
		obj, err := o.ConnectedStatic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectedStaticVal = obj
	}
	var ospfVal *Ospf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *Ospfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ripVal *Rip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		Bgp:             bgpVal,
		ConnectedStatic: connectedStaticVal,
		Ospf:            ospfVal,
		Ospfv3:          ospfv3Val,
		Rip:             ripVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpXml_11_0_2) MarshalFromObject(s Bgp) {
	if s.Ospf != nil {
		var obj bgpOspfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj bgpOspfv3Xml_11_0_2
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rib != nil {
		var obj bgpRibXml_11_0_2
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj bgpRipXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpXml_11_0_2) UnmarshalToObject() (*Bgp, error) {
	var ospfVal *BgpOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *BgpOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ribVal *BgpRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *BgpRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Bgp{
		Ospf:           ospfVal,
		Ospfv3:         ospfv3Val,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfXml_11_0_2) MarshalFromObject(s BgpOspf) {
	if s.RouteMap != nil {
		var objs []bgpOspfRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj bgpOspfRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpOspfRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfXml_11_0_2) UnmarshalToObject() (*BgpOspf, error) {
	var routeMapVal []BgpOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapXml_11_0_2) MarshalFromObject(s BgpOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpOspfRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpOspfRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMap, error) {
	var matchVal *BgpOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpOspfRouteMap{
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
func (o *bgpOspfRouteMapMatchXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapMatch) {
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
		var obj bgpOspfRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapMatch, error) {
	var ipv4Val *BgpOspfRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &BgpOspfRouteMapMatch{
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
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s BgpOspfRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpOspfRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpOspfRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpOspfRouteMapMatchIpv4RouteSourceXml_11_0_2
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4, error) {
	var addressVal *BgpOspfRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpOspfRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpOspfRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpOspfRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4Address, error) {

	result := &BgpOspfRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4NextHop, error) {

	result := &BgpOspfRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapMatchIpv4RouteSourceXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapMatchIpv4RouteSourceXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapMatchIpv4RouteSource, error) {

	result := &BgpOspfRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapSetXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapSet) {
	if s.Metric != nil {
		var obj bgpOspfRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapSetXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapSet, error) {
	var metricVal *BgpOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfRouteMapSetMetricXml_11_0_2) MarshalFromObject(s BgpOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*BgpOspfRouteMapSetMetric, error) {

	result := &BgpOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3Xml_11_0_2) MarshalFromObject(s BgpOspfv3) {
	if s.RouteMap != nil {
		var objs []bgpOspfv3RouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj bgpOspfv3RouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpOspfv3RouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3Xml_11_0_2) UnmarshalToObject() (*BgpOspfv3, error) {
	var routeMapVal []BgpOspfv3RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpOspfv3{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpOspfv3RouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpOspfv3RouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMap, error) {
	var matchVal *BgpOspfv3RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpOspfv3RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpOspfv3RouteMap{
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
func (o *bgpOspfv3RouteMapMatchXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapMatch) {
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
	if s.Ipv6 != nil {
		var obj bgpOspfv3RouteMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapMatch, error) {
	var ipv6Val *BgpOspfv3RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &BgpOspfv3RouteMapMatch{
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
		Ipv6:              ipv6Val,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6Xml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj bgpOspfv3RouteMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpOspfv3RouteMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6, error) {
	var addressVal *BgpOspfv3RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpOspfv3RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &BgpOspfv3RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6Address, error) {

	result := &BgpOspfv3RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapMatchIpv6NextHop, error) {

	result := &BgpOspfv3RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapSetXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapSet) {
	if s.Metric != nil {
		var obj bgpOspfv3RouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapSetXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapSet, error) {
	var metricVal *BgpOspfv3RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpOspfv3RouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpOspfv3RouteMapSetMetricXml_11_0_2) MarshalFromObject(s BgpOspfv3RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpOspfv3RouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*BgpOspfv3RouteMapSetMetric, error) {

	result := &BgpOspfv3RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibXml_11_0_2) MarshalFromObject(s BgpRib) {
	if s.RouteMap != nil {
		var objs []bgpRibRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj bgpRibRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpRibRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibXml_11_0_2) UnmarshalToObject() (*BgpRib, error) {
	var routeMapVal []BgpRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapXml_11_0_2) MarshalFromObject(s BgpRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpRibRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpRibRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMap, error) {
	var matchVal *BgpRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpRibRouteMap{
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
func (o *bgpRibRouteMapMatchXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatch) {
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
		var obj bgpRibRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj bgpRibRouteMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatch, error) {
	var ipv4Val *BgpRibRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *BgpRibRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &BgpRibRouteMapMatch{
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
func (o *bgpRibRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpRibRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRibRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpRibRouteMapMatchIpv4RouteSourceXml_11_0_2
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4, error) {
	var addressVal *BgpRibRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRibRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpRibRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpRibRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4Address, error) {

	result := &BgpRibRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4NextHop, error) {

	result := &BgpRibRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv4RouteSourceXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv4RouteSourceXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv4RouteSource, error) {

	result := &BgpRibRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6Xml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj bgpRibRouteMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRibRouteMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6, error) {
	var addressVal *BgpRibRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRibRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &BgpRibRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6Address, error) {

	result := &BgpRibRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s BgpRibRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapMatchIpv6NextHop, error) {

	result := &BgpRibRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRibRouteMapSetXml_11_0_2) MarshalFromObject(s BgpRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRibRouteMapSetXml_11_0_2) UnmarshalToObject() (*BgpRibRouteMapSet, error) {

	result := &BgpRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipXml_11_0_2) MarshalFromObject(s BgpRip) {
	if s.RouteMap != nil {
		var objs []bgpRipRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj bgpRipRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &bgpRipRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipXml_11_0_2) UnmarshalToObject() (*BgpRip, error) {
	var routeMapVal []BgpRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &BgpRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapXml_11_0_2) MarshalFromObject(s BgpRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj bgpRipRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj bgpRipRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMap, error) {
	var matchVal *BgpRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *BgpRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &BgpRipRouteMap{
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
func (o *bgpRipRouteMapMatchXml_11_0_2) MarshalFromObject(s BgpRipRouteMapMatch) {
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
		var obj bgpRipRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapMatch, error) {
	var ipv4Val *BgpRipRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &BgpRipRouteMapMatch{
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
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s BgpRipRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj bgpRipRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj bgpRipRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.RouteSource != nil {
		var obj bgpRipRouteMapMatchIpv4RouteSourceXml_11_0_2
		obj.MarshalFromObject(*s.RouteSource)
		o.RouteSource = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4, error) {
	var addressVal *BgpRipRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *BgpRipRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var routeSourceVal *BgpRipRouteMapMatchIpv4RouteSource
	if o.RouteSource != nil {
		obj, err := o.RouteSource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeSourceVal = obj
	}

	result := &BgpRipRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		RouteSource:    routeSourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s BgpRipRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4Address, error) {

	result := &BgpRipRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s BgpRipRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4NextHop, error) {

	result := &BgpRipRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapMatchIpv4RouteSourceXml_11_0_2) MarshalFromObject(s BgpRipRouteMapMatchIpv4RouteSource) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapMatchIpv4RouteSourceXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapMatchIpv4RouteSource, error) {

	result := &BgpRipRouteMapMatchIpv4RouteSource{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapSetXml_11_0_2) MarshalFromObject(s BgpRipRouteMapSet) {
	if s.Metric != nil {
		var obj bgpRipRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapSetXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapSet, error) {
	var metricVal *BgpRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &BgpRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bgpRipRouteMapSetMetricXml_11_0_2) MarshalFromObject(s BgpRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bgpRipRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*BgpRipRouteMapSetMetric, error) {

	result := &BgpRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticXml_11_0_2) MarshalFromObject(s ConnectedStatic) {
	if s.Bgp != nil {
		var obj connectedStaticBgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj connectedStaticOspfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj connectedStaticOspfv3Xml_11_0_2
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Rib != nil {
		var obj connectedStaticRibXml_11_0_2
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj connectedStaticRipXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticXml_11_0_2) UnmarshalToObject() (*ConnectedStatic, error) {
	var bgpVal *ConnectedStaticBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *ConnectedStaticOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *ConnectedStaticOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ribVal *ConnectedStaticRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *ConnectedStaticRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &ConnectedStatic{
		Bgp:            bgpVal,
		Ospf:           ospfVal,
		Ospfv3:         ospfv3Val,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpXml_11_0_2) MarshalFromObject(s ConnectedStaticBgp) {
	if s.RouteMap != nil {
		var objs []connectedStaticBgpRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj connectedStaticBgpRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticBgpRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgp, error) {
	var routeMapVal []ConnectedStaticBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticBgpRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticBgpRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMap, error) {
	var matchVal *ConnectedStaticBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticBgpRouteMap{
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
func (o *connectedStaticBgpRouteMapMatchXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticBgpRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticBgpRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticBgpRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticBgpRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticBgpRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticBgpRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6Xml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticBgpRouteMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticBgpRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticBgpRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticBgpRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticBgpRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj connectedStaticBgpRouteMapSetAggregatorXml_11_0_2
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj connectedStaticBgpRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj connectedStaticBgpRouteMapSetIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticBgpRouteMapSetIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSet, error) {
	var aggregatorVal *ConnectedStaticBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *ConnectedStaticBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *ConnectedStaticBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticBgpRouteMapSetIpv6
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &ConnectedStaticBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Ipv6:              ipv6Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetAggregatorXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetAggregatorXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetAggregator, error) {

	result := &ConnectedStaticBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetMetricXml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetMetric, error) {

	result := &ConnectedStaticBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetIpv4Xml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetIpv4Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetIpv4, error) {

	result := &ConnectedStaticBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticBgpRouteMapSetIpv6Xml_11_0_2) MarshalFromObject(s ConnectedStaticBgpRouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticBgpRouteMapSetIpv6Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticBgpRouteMapSetIpv6, error) {

	result := &ConnectedStaticBgpRouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfXml_11_0_2) MarshalFromObject(s ConnectedStaticOspf) {
	if s.RouteMap != nil {
		var objs []connectedStaticOspfRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj connectedStaticOspfRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticOspfRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspf, error) {
	var routeMapVal []ConnectedStaticOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticOspfRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticOspfRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMap, error) {
	var matchVal *ConnectedStaticOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticOspfRouteMap{
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
func (o *connectedStaticOspfRouteMapMatchXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticOspfRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &ConnectedStaticOspfRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticOspfRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticOspfRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticOspfRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticOspfRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticOspfRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticOspfRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapSetXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticOspfRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapSetXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapSet, error) {
	var metricVal *ConnectedStaticOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfRouteMapSetMetricXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfRouteMapSetMetric, error) {

	result := &ConnectedStaticOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3Xml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3) {
	if s.RouteMap != nil {
		var objs []connectedStaticOspfv3RouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj connectedStaticOspfv3RouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticOspfv3RouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3, error) {
	var routeMapVal []ConnectedStaticOspfv3RouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticOspfv3{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticOspfv3RouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticOspfv3RouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMap, error) {
	var matchVal *ConnectedStaticOspfv3RouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticOspfv3RouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMap{
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
func (o *connectedStaticOspfv3RouteMapMatchXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv6 != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatch, error) {
	var ipv6Val *ConnectedStaticOspfv3RouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticOspfv3RouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6Xml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticOspfv3RouteMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticOspfv3RouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticOspfv3RouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticOspfv3RouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapSetXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticOspfv3RouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapSetXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapSet, error) {
	var metricVal *ConnectedStaticOspfv3RouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticOspfv3RouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticOspfv3RouteMapSetMetricXml_11_0_2) MarshalFromObject(s ConnectedStaticOspfv3RouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticOspfv3RouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*ConnectedStaticOspfv3RouteMapSetMetric, error) {

	result := &ConnectedStaticOspfv3RouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibXml_11_0_2) MarshalFromObject(s ConnectedStaticRib) {
	if s.RouteMap != nil {
		var objs []connectedStaticRibRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj connectedStaticRibRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticRibRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRib, error) {
	var routeMapVal []ConnectedStaticRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticRibRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticRibRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMap, error) {
	var matchVal *ConnectedStaticRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticRibRouteMap{
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
func (o *connectedStaticRibRouteMapMatchXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticRibRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj connectedStaticRibRouteMapMatchIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticRibRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *ConnectedStaticRibRouteMapMatchIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &ConnectedStaticRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticRibRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRibRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticRibRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRibRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRibRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6Xml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6) {
	if s.Address != nil {
		var obj connectedStaticRibRouteMapMatchIpv6AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRibRouteMapMatchIpv6NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6, error) {
	var addressVal *ConnectedStaticRibRouteMapMatchIpv6Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRibRouteMapMatchIpv6NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRibRouteMapMatchIpv6{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6Address, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv6Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapMatchIpv6NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapMatchIpv6NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapMatchIpv6NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapMatchIpv6NextHop, error) {

	result := &ConnectedStaticRibRouteMapMatchIpv6NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRibRouteMapSetXml_11_0_2) MarshalFromObject(s ConnectedStaticRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRibRouteMapSetXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRibRouteMapSet, error) {

	result := &ConnectedStaticRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipXml_11_0_2) MarshalFromObject(s ConnectedStaticRip) {
	if s.RouteMap != nil {
		var objs []connectedStaticRipRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj connectedStaticRipRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &connectedStaticRipRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRip, error) {
	var routeMapVal []ConnectedStaticRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &ConnectedStaticRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj connectedStaticRipRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj connectedStaticRipRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMap, error) {
	var matchVal *ConnectedStaticRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *ConnectedStaticRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &ConnectedStaticRipRouteMap{
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
func (o *connectedStaticRipRouteMapMatchXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Ipv4 != nil {
		var obj connectedStaticRipRouteMapMatchIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatch, error) {
	var ipv4Val *ConnectedStaticRipRouteMapMatchIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}

	result := &ConnectedStaticRipRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Ipv4:           ipv4Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4Xml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4) {
	if s.Address != nil {
		var obj connectedStaticRipRouteMapMatchIpv4AddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj connectedStaticRipRouteMapMatchIpv4NextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4Xml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4, error) {
	var addressVal *ConnectedStaticRipRouteMapMatchIpv4Address
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *ConnectedStaticRipRouteMapMatchIpv4NextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &ConnectedStaticRipRouteMapMatchIpv4{
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4AddressXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4Address) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4AddressXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4Address, error) {

	result := &ConnectedStaticRipRouteMapMatchIpv4Address{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapMatchIpv4NextHopXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapMatchIpv4NextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapMatchIpv4NextHopXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapMatchIpv4NextHop, error) {

	result := &ConnectedStaticRipRouteMapMatchIpv4NextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapSetXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapSet) {
	if s.Metric != nil {
		var obj connectedStaticRipRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapSetXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapSet, error) {
	var metricVal *ConnectedStaticRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &ConnectedStaticRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *connectedStaticRipRouteMapSetMetricXml_11_0_2) MarshalFromObject(s ConnectedStaticRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o connectedStaticRipRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*ConnectedStaticRipRouteMapSetMetric, error) {

	result := &ConnectedStaticRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfXml_11_0_2) MarshalFromObject(s Ospf) {
	if s.Bgp != nil {
		var obj ospfBgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Rib != nil {
		var obj ospfRibXml_11_0_2
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	if s.Rip != nil {
		var obj ospfRipXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfXml_11_0_2) UnmarshalToObject() (*Ospf, error) {
	var bgpVal *OspfBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ribVal *OspfRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}
	var ripVal *OspfRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Ospf{
		Bgp:            bgpVal,
		Rib:            ribVal,
		Rip:            ripVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpXml_11_0_2) MarshalFromObject(s OspfBgp) {
	if s.RouteMap != nil {
		var objs []ospfBgpRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ospfBgpRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfBgpRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpXml_11_0_2) UnmarshalToObject() (*OspfBgp, error) {
	var routeMapVal []OspfBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapXml_11_0_2) MarshalFromObject(s OspfBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfBgpRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfBgpRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMap, error) {
	var matchVal *OspfBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfBgpRouteMap{
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
func (o *ospfBgpRouteMapMatchXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfBgpRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfBgpRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapMatch, error) {
	var addressVal *OspfBgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfBgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapMatchAddress, error) {

	result := &OspfBgpRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapMatchNextHop, error) {

	result := &OspfBgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ospfBgpRouteMapSetAggregatorXml_11_0_2
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ospfBgpRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj ospfBgpRouteMapSetIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapSet, error) {
	var aggregatorVal *OspfBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *OspfBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *OspfBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &OspfBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetAggregatorXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetAggregatorXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapSetAggregator, error) {

	result := &OspfBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetMetricXml_11_0_2) MarshalFromObject(s OspfBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapSetMetric, error) {

	result := &OspfBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfBgpRouteMapSetIpv4Xml_11_0_2) MarshalFromObject(s OspfBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfBgpRouteMapSetIpv4Xml_11_0_2) UnmarshalToObject() (*OspfBgpRouteMapSetIpv4, error) {

	result := &OspfBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibXml_11_0_2) MarshalFromObject(s OspfRib) {
	if s.RouteMap != nil {
		var objs []ospfRibRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ospfRibRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfRibRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibXml_11_0_2) UnmarshalToObject() (*OspfRib, error) {
	var routeMapVal []OspfRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapXml_11_0_2) MarshalFromObject(s OspfRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfRibRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfRibRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapXml_11_0_2) UnmarshalToObject() (*OspfRibRouteMap, error) {
	var matchVal *OspfRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfRibRouteMap{
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
func (o *ospfRibRouteMapMatchXml_11_0_2) MarshalFromObject(s OspfRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfRibRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfRibRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchXml_11_0_2) UnmarshalToObject() (*OspfRibRouteMapMatch, error) {
	var addressVal *OspfRibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfRibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s OspfRibRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*OspfRibRouteMapMatchAddress, error) {

	result := &OspfRibRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s OspfRibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*OspfRibRouteMapMatchNextHop, error) {

	result := &OspfRibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRibRouteMapSetXml_11_0_2) MarshalFromObject(s OspfRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRibRouteMapSetXml_11_0_2) UnmarshalToObject() (*OspfRibRouteMapSet, error) {

	result := &OspfRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipXml_11_0_2) MarshalFromObject(s OspfRip) {
	if s.RouteMap != nil {
		var objs []ospfRipRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ospfRipRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfRipRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipXml_11_0_2) UnmarshalToObject() (*OspfRip, error) {
	var routeMapVal []OspfRipRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &OspfRip{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapXml_11_0_2) MarshalFromObject(s OspfRipRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfRipRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfRipRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMap, error) {
	var matchVal *OspfRipRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *OspfRipRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &OspfRipRouteMap{
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
func (o *ospfRipRouteMapMatchXml_11_0_2) MarshalFromObject(s OspfRipRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfRipRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfRipRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMapMatch, error) {
	var addressVal *OspfRipRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *OspfRipRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &OspfRipRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s OspfRipRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMapMatchAddress, error) {

	result := &OspfRipRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s OspfRipRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMapMatchNextHop, error) {

	result := &OspfRipRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapSetXml_11_0_2) MarshalFromObject(s OspfRipRouteMapSet) {
	if s.Metric != nil {
		var obj ospfRipRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.NextHop = s.NextHop
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapSetXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMapSet, error) {
	var metricVal *OspfRipRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &OspfRipRouteMapSet{
		Metric:         metricVal,
		NextHop:        o.NextHop,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfRipRouteMapSetMetricXml_11_0_2) MarshalFromObject(s OspfRipRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfRipRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*OspfRipRouteMapSetMetric, error) {

	result := &OspfRipRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3Xml_11_0_2) MarshalFromObject(s Ospfv3) {
	if s.Bgp != nil {
		var obj ospfv3BgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Rib != nil {
		var obj ospfv3RibXml_11_0_2
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3Xml_11_0_2) UnmarshalToObject() (*Ospfv3, error) {
	var bgpVal *Ospfv3Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ribVal *Ospfv3Rib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}

	result := &Ospfv3{
		Bgp:            bgpVal,
		Rib:            ribVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpXml_11_0_2) MarshalFromObject(s Ospfv3Bgp) {
	if s.RouteMap != nil {
		var objs []ospfv3BgpRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ospfv3BgpRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfv3BgpRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpXml_11_0_2) UnmarshalToObject() (*Ospfv3Bgp, error) {
	var routeMapVal []Ospfv3BgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Ospfv3Bgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfv3BgpRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfv3BgpRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMap, error) {
	var matchVal *Ospfv3BgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *Ospfv3BgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &Ospfv3BgpRouteMap{
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
func (o *ospfv3BgpRouteMapMatchXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfv3BgpRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfv3BgpRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapMatch, error) {
	var addressVal *Ospfv3BgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *Ospfv3BgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &Ospfv3BgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapMatchAddress) {
	o.PrefixList = s.PrefixList
	o.AccessList = s.AccessList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapMatchAddress, error) {

	result := &Ospfv3BgpRouteMapMatchAddress{
		PrefixList:     o.PrefixList,
		AccessList:     o.AccessList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapMatchNextHop, error) {

	result := &Ospfv3BgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ospfv3BgpRouteMapSetAggregatorXml_11_0_2
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ospfv3BgpRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv6 != nil {
		var obj ospfv3BgpRouteMapSetIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapSet, error) {
	var aggregatorVal *Ospfv3BgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *Ospfv3BgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv6Val *Ospfv3BgpRouteMapSetIpv6
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &Ospfv3BgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv6:              ipv6Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetAggregatorXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetAggregatorXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapSetAggregator, error) {

	result := &Ospfv3BgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetMetricXml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapSetMetric, error) {

	result := &Ospfv3BgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3BgpRouteMapSetIpv6Xml_11_0_2) MarshalFromObject(s Ospfv3BgpRouteMapSetIpv6) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3BgpRouteMapSetIpv6Xml_11_0_2) UnmarshalToObject() (*Ospfv3BgpRouteMapSetIpv6, error) {

	result := &Ospfv3BgpRouteMapSetIpv6{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibXml_11_0_2) MarshalFromObject(s Ospfv3Rib) {
	if s.RouteMap != nil {
		var objs []ospfv3RibRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ospfv3RibRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ospfv3RibRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibXml_11_0_2) UnmarshalToObject() (*Ospfv3Rib, error) {
	var routeMapVal []Ospfv3RibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &Ospfv3Rib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapXml_11_0_2) MarshalFromObject(s Ospfv3RibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ospfv3RibRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ospfv3RibRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapXml_11_0_2) UnmarshalToObject() (*Ospfv3RibRouteMap, error) {
	var matchVal *Ospfv3RibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *Ospfv3RibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &Ospfv3RibRouteMap{
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
func (o *ospfv3RibRouteMapMatchXml_11_0_2) MarshalFromObject(s Ospfv3RibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ospfv3RibRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ospfv3RibRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchXml_11_0_2) UnmarshalToObject() (*Ospfv3RibRouteMapMatch, error) {
	var addressVal *Ospfv3RibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *Ospfv3RibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &Ospfv3RibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s Ospfv3RibRouteMapMatchAddress) {
	o.PrefixList = s.PrefixList
	o.AccessList = s.AccessList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*Ospfv3RibRouteMapMatchAddress, error) {

	result := &Ospfv3RibRouteMapMatchAddress{
		PrefixList:     o.PrefixList,
		AccessList:     o.AccessList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s Ospfv3RibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*Ospfv3RibRouteMapMatchNextHop, error) {

	result := &Ospfv3RibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ospfv3RibRouteMapSetXml_11_0_2) MarshalFromObject(s Ospfv3RibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ospfv3RibRouteMapSetXml_11_0_2) UnmarshalToObject() (*Ospfv3RibRouteMapSet, error) {

	result := &Ospfv3RibRouteMapSet{
		SourceAddress:  o.SourceAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripXml_11_0_2) MarshalFromObject(s Rip) {
	if s.Bgp != nil {
		var obj ripBgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj ripOspfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Rib != nil {
		var obj ripRibXml_11_0_2
		obj.MarshalFromObject(*s.Rib)
		o.Rib = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripXml_11_0_2) UnmarshalToObject() (*Rip, error) {
	var bgpVal *RipBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *RipOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ribVal *RipRib
	if o.Rib != nil {
		obj, err := o.Rib.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribVal = obj
	}

	result := &Rip{
		Bgp:            bgpVal,
		Ospf:           ospfVal,
		Rib:            ribVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpXml_11_0_2) MarshalFromObject(s RipBgp) {
	if s.RouteMap != nil {
		var objs []ripBgpRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ripBgpRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripBgpRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpXml_11_0_2) UnmarshalToObject() (*RipBgp, error) {
	var routeMapVal []RipBgpRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipBgp{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapXml_11_0_2) MarshalFromObject(s RipBgpRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripBgpRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripBgpRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMap, error) {
	var matchVal *RipBgpRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipBgpRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipBgpRouteMap{
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
func (o *ripBgpRouteMapMatchXml_11_0_2) MarshalFromObject(s RipBgpRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripBgpRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripBgpRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapMatch, error) {
	var addressVal *RipBgpRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipBgpRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipBgpRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s RipBgpRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapMatchAddress, error) {

	result := &RipBgpRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s RipBgpRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapMatchNextHop, error) {

	result := &RipBgpRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetXml_11_0_2) MarshalFromObject(s RipBgpRouteMapSet) {
	if s.Aggregator != nil {
		var obj ripBgpRouteMapSetAggregatorXml_11_0_2
		obj.MarshalFromObject(*s.Aggregator)
		o.Aggregator = &obj
	}
	if s.Metric != nil {
		var obj ripBgpRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	if s.Ipv4 != nil {
		var obj ripBgpRouteMapSetIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	o.Tag = s.Tag
	o.LocalPreference = s.LocalPreference
	o.Weight = s.Weight
	o.Origin = s.Origin
	o.AtomicAggregate = util.YesNo(s.AtomicAggregate, nil)
	o.OriginatorId = s.OriginatorId
	if s.AspathPrepend != nil {
		o.AspathPrepend = util.Int64ToMem(s.AspathPrepend)
	}
	if s.RegularCommunity != nil {
		o.RegularCommunity = util.StrToMem(s.RegularCommunity)
	}
	if s.LargeCommunity != nil {
		o.LargeCommunity = util.StrToMem(s.LargeCommunity)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapSet, error) {
	var aggregatorVal *RipBgpRouteMapSetAggregator
	if o.Aggregator != nil {
		obj, err := o.Aggregator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregatorVal = obj
	}
	var metricVal *RipBgpRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}
	var ipv4Val *RipBgpRouteMapSetIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
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
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &RipBgpRouteMapSet{
		Aggregator:        aggregatorVal,
		Metric:            metricVal,
		Ipv4:              ipv4Val,
		Tag:               o.Tag,
		LocalPreference:   o.LocalPreference,
		Weight:            o.Weight,
		Origin:            o.Origin,
		AtomicAggregate:   util.AsBool(o.AtomicAggregate, nil),
		OriginatorId:      o.OriginatorId,
		AspathPrepend:     aspathPrependVal,
		RegularCommunity:  regularCommunityVal,
		LargeCommunity:    largeCommunityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetAggregatorXml_11_0_2) MarshalFromObject(s RipBgpRouteMapSetAggregator) {
	o.As = s.As
	o.RouterId = s.RouterId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetAggregatorXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapSetAggregator, error) {

	result := &RipBgpRouteMapSetAggregator{
		As:             o.As,
		RouterId:       o.RouterId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetMetricXml_11_0_2) MarshalFromObject(s RipBgpRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapSetMetric, error) {

	result := &RipBgpRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripBgpRouteMapSetIpv4Xml_11_0_2) MarshalFromObject(s RipBgpRouteMapSetIpv4) {
	o.SourceAddress = s.SourceAddress
	o.NextHop = s.NextHop
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripBgpRouteMapSetIpv4Xml_11_0_2) UnmarshalToObject() (*RipBgpRouteMapSetIpv4, error) {

	result := &RipBgpRouteMapSetIpv4{
		SourceAddress:  o.SourceAddress,
		NextHop:        o.NextHop,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfXml_11_0_2) MarshalFromObject(s RipOspf) {
	if s.RouteMap != nil {
		var objs []ripOspfRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ripOspfRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripOspfRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfXml_11_0_2) UnmarshalToObject() (*RipOspf, error) {
	var routeMapVal []RipOspfRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipOspf{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapXml_11_0_2) MarshalFromObject(s RipOspfRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripOspfRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripOspfRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMap, error) {
	var matchVal *RipOspfRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipOspfRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipOspfRouteMap{
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
func (o *ripOspfRouteMapMatchXml_11_0_2) MarshalFromObject(s RipOspfRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripOspfRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripOspfRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMapMatch, error) {
	var addressVal *RipOspfRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipOspfRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipOspfRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s RipOspfRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMapMatchAddress, error) {

	result := &RipOspfRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s RipOspfRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMapMatchNextHop, error) {

	result := &RipOspfRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapSetXml_11_0_2) MarshalFromObject(s RipOspfRouteMapSet) {
	if s.Metric != nil {
		var obj ripOspfRouteMapSetMetricXml_11_0_2
		obj.MarshalFromObject(*s.Metric)
		o.Metric = &obj
	}
	o.MetricType = s.MetricType
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapSetXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMapSet, error) {
	var metricVal *RipOspfRouteMapSetMetric
	if o.Metric != nil {
		obj, err := o.Metric.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		metricVal = obj
	}

	result := &RipOspfRouteMapSet{
		Metric:         metricVal,
		MetricType:     o.MetricType,
		Tag:            o.Tag,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripOspfRouteMapSetMetricXml_11_0_2) MarshalFromObject(s RipOspfRouteMapSetMetric) {
	o.Value = s.Value
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripOspfRouteMapSetMetricXml_11_0_2) UnmarshalToObject() (*RipOspfRouteMapSetMetric, error) {

	result := &RipOspfRouteMapSetMetric{
		Value:          o.Value,
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibXml_11_0_2) MarshalFromObject(s RipRib) {
	if s.RouteMap != nil {
		var objs []ripRibRouteMapXml_11_0_2
		for _, elt := range s.RouteMap {
			var obj ripRibRouteMapXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RouteMap = &ripRibRouteMapContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibXml_11_0_2) UnmarshalToObject() (*RipRib, error) {
	var routeMapVal []RipRibRouteMap
	if o.RouteMap != nil {
		for _, elt := range o.RouteMap.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			routeMapVal = append(routeMapVal, *obj)
		}
	}

	result := &RipRib{
		RouteMap:       routeMapVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapXml_11_0_2) MarshalFromObject(s RipRibRouteMap) {
	o.Name = s.Name
	o.Action = s.Action
	o.Description = s.Description
	if s.Match != nil {
		var obj ripRibRouteMapMatchXml_11_0_2
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Set != nil {
		var obj ripRibRouteMapSetXml_11_0_2
		obj.MarshalFromObject(*s.Set)
		o.Set = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapXml_11_0_2) UnmarshalToObject() (*RipRibRouteMap, error) {
	var matchVal *RipRibRouteMapMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var setVal *RipRibRouteMapSet
	if o.Set != nil {
		obj, err := o.Set.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setVal = obj
	}

	result := &RipRibRouteMap{
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
func (o *ripRibRouteMapMatchXml_11_0_2) MarshalFromObject(s RipRibRouteMapMatch) {
	o.Interface = s.Interface
	o.Metric = s.Metric
	o.Tag = s.Tag
	if s.Address != nil {
		var obj ripRibRouteMapMatchAddressXml_11_0_2
		obj.MarshalFromObject(*s.Address)
		o.Address = &obj
	}
	if s.NextHop != nil {
		var obj ripRibRouteMapMatchNextHopXml_11_0_2
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchXml_11_0_2) UnmarshalToObject() (*RipRibRouteMapMatch, error) {
	var addressVal *RipRibRouteMapMatchAddress
	if o.Address != nil {
		obj, err := o.Address.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressVal = obj
	}
	var nextHopVal *RipRibRouteMapMatchNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}

	result := &RipRibRouteMapMatch{
		Interface:      o.Interface,
		Metric:         o.Metric,
		Tag:            o.Tag,
		Address:        addressVal,
		NextHop:        nextHopVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapMatchAddressXml_11_0_2) MarshalFromObject(s RipRibRouteMapMatchAddress) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchAddressXml_11_0_2) UnmarshalToObject() (*RipRibRouteMapMatchAddress, error) {

	result := &RipRibRouteMapMatchAddress{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapMatchNextHopXml_11_0_2) MarshalFromObject(s RipRibRouteMapMatchNextHop) {
	o.AccessList = s.AccessList
	o.PrefixList = s.PrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapMatchNextHopXml_11_0_2) UnmarshalToObject() (*RipRibRouteMapMatchNextHop, error) {

	result := &RipRibRouteMapMatchNextHop{
		AccessList:     o.AccessList,
		PrefixList:     o.PrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ripRibRouteMapSetXml_11_0_2) MarshalFromObject(s RipRibRouteMapSet) {
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ripRibRouteMapSetXml_11_0_2) UnmarshalToObject() (*RipRibRouteMapSet, error) {

	result := &RipRibRouteMapSet{
		SourceAddress:  o.SourceAddress,
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
	if v == "bgp" || v == "Bgp" {
		return e.Bgp, nil
	}
	if v == "connected_static" || v == "ConnectedStatic" {
		return e.ConnectedStatic, nil
	}
	if v == "ospf" || v == "Ospf" {
		return e.Ospf, nil
	}
	if v == "ospfv3" || v == "Ospfv3" {
		return e.Ospfv3, nil
	}
	if v == "rip" || v == "Rip" {
		return e.Rip, nil
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
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.ConnectedStatic.matches(other.ConnectedStatic) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
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
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Rib.matches(other.Rib) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *BgpOspf) matches(other *BgpOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *BgpOspfRouteMap) matches(other *BgpOspfRouteMap) bool {
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

func (o *BgpOspfRouteMapMatch) matches(other *BgpOspfRouteMapMatch) bool {
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

	return true
}

func (o *BgpOspfRouteMapMatchIpv4) matches(other *BgpOspfRouteMapMatchIpv4) bool {
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

func (o *BgpOspfRouteMapMatchIpv4Address) matches(other *BgpOspfRouteMapMatchIpv4Address) bool {
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

func (o *BgpOspfRouteMapMatchIpv4NextHop) matches(other *BgpOspfRouteMapMatchIpv4NextHop) bool {
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

func (o *BgpOspfRouteMapMatchIpv4RouteSource) matches(other *BgpOspfRouteMapMatchIpv4RouteSource) bool {
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

func (o *BgpOspfRouteMapSet) matches(other *BgpOspfRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *BgpOspfRouteMapSetMetric) matches(other *BgpOspfRouteMapSetMetric) bool {
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

func (o *BgpOspfv3) matches(other *BgpOspfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *BgpOspfv3RouteMap) matches(other *BgpOspfv3RouteMap) bool {
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

func (o *BgpOspfv3RouteMapMatch) matches(other *BgpOspfv3RouteMapMatch) bool {
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
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *BgpOspfv3RouteMapMatchIpv6) matches(other *BgpOspfv3RouteMapMatchIpv6) bool {
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

func (o *BgpOspfv3RouteMapMatchIpv6Address) matches(other *BgpOspfv3RouteMapMatchIpv6Address) bool {
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

func (o *BgpOspfv3RouteMapMatchIpv6NextHop) matches(other *BgpOspfv3RouteMapMatchIpv6NextHop) bool {
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

func (o *BgpOspfv3RouteMapSet) matches(other *BgpOspfv3RouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *BgpOspfv3RouteMapSetMetric) matches(other *BgpOspfv3RouteMapSetMetric) bool {
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

func (o *BgpRib) matches(other *BgpRib) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *BgpRibRouteMap) matches(other *BgpRibRouteMap) bool {
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

func (o *BgpRibRouteMapMatch) matches(other *BgpRibRouteMapMatch) bool {
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

func (o *BgpRibRouteMapMatchIpv4) matches(other *BgpRibRouteMapMatchIpv4) bool {
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

func (o *BgpRibRouteMapMatchIpv4Address) matches(other *BgpRibRouteMapMatchIpv4Address) bool {
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

func (o *BgpRibRouteMapMatchIpv4NextHop) matches(other *BgpRibRouteMapMatchIpv4NextHop) bool {
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

func (o *BgpRibRouteMapMatchIpv4RouteSource) matches(other *BgpRibRouteMapMatchIpv4RouteSource) bool {
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

func (o *BgpRibRouteMapMatchIpv6) matches(other *BgpRibRouteMapMatchIpv6) bool {
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

func (o *BgpRibRouteMapMatchIpv6Address) matches(other *BgpRibRouteMapMatchIpv6Address) bool {
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

func (o *BgpRibRouteMapMatchIpv6NextHop) matches(other *BgpRibRouteMapMatchIpv6NextHop) bool {
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

func (o *BgpRibRouteMapSet) matches(other *BgpRibRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}

	return true
}

func (o *BgpRip) matches(other *BgpRip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *BgpRipRouteMap) matches(other *BgpRipRouteMap) bool {
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

func (o *BgpRipRouteMapMatch) matches(other *BgpRipRouteMapMatch) bool {
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

	return true
}

func (o *BgpRipRouteMapMatchIpv4) matches(other *BgpRipRouteMapMatchIpv4) bool {
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

func (o *BgpRipRouteMapMatchIpv4Address) matches(other *BgpRipRouteMapMatchIpv4Address) bool {
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

func (o *BgpRipRouteMapMatchIpv4NextHop) matches(other *BgpRipRouteMapMatchIpv4NextHop) bool {
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

func (o *BgpRipRouteMapMatchIpv4RouteSource) matches(other *BgpRipRouteMapMatchIpv4RouteSource) bool {
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

func (o *BgpRipRouteMapSet) matches(other *BgpRipRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.NextHop, other.NextHop) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *BgpRipRouteMapSetMetric) matches(other *BgpRipRouteMapSetMetric) bool {
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

func (o *ConnectedStatic) matches(other *ConnectedStatic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Rib.matches(other.Rib) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *ConnectedStaticBgp) matches(other *ConnectedStaticBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *ConnectedStaticBgpRouteMap) matches(other *ConnectedStaticBgpRouteMap) bool {
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

func (o *ConnectedStaticBgpRouteMapMatch) matches(other *ConnectedStaticBgpRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv4) matches(other *ConnectedStaticBgpRouteMapMatchIpv4) bool {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv4Address) matches(other *ConnectedStaticBgpRouteMapMatchIpv4Address) bool {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv4NextHop) matches(other *ConnectedStaticBgpRouteMapMatchIpv4NextHop) bool {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv6) matches(other *ConnectedStaticBgpRouteMapMatchIpv6) bool {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv6Address) matches(other *ConnectedStaticBgpRouteMapMatchIpv6Address) bool {
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

func (o *ConnectedStaticBgpRouteMapMatchIpv6NextHop) matches(other *ConnectedStaticBgpRouteMapMatchIpv6NextHop) bool {
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

func (o *ConnectedStaticBgpRouteMapSet) matches(other *ConnectedStaticBgpRouteMapSet) bool {
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
	if !util.OrderedListsMatch[int64](o.AspathPrepend, other.AspathPrepend) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ConnectedStaticBgpRouteMapSetAggregator) matches(other *ConnectedStaticBgpRouteMapSetAggregator) bool {
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

func (o *ConnectedStaticBgpRouteMapSetMetric) matches(other *ConnectedStaticBgpRouteMapSetMetric) bool {
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

func (o *ConnectedStaticBgpRouteMapSetIpv4) matches(other *ConnectedStaticBgpRouteMapSetIpv4) bool {
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

func (o *ConnectedStaticBgpRouteMapSetIpv6) matches(other *ConnectedStaticBgpRouteMapSetIpv6) bool {
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

func (o *ConnectedStaticOspf) matches(other *ConnectedStaticOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *ConnectedStaticOspfRouteMap) matches(other *ConnectedStaticOspfRouteMap) bool {
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

func (o *ConnectedStaticOspfRouteMapMatch) matches(other *ConnectedStaticOspfRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}

	return true
}

func (o *ConnectedStaticOspfRouteMapMatchIpv4) matches(other *ConnectedStaticOspfRouteMapMatchIpv4) bool {
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

func (o *ConnectedStaticOspfRouteMapMatchIpv4Address) matches(other *ConnectedStaticOspfRouteMapMatchIpv4Address) bool {
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

func (o *ConnectedStaticOspfRouteMapMatchIpv4NextHop) matches(other *ConnectedStaticOspfRouteMapMatchIpv4NextHop) bool {
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

func (o *ConnectedStaticOspfRouteMapSet) matches(other *ConnectedStaticOspfRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *ConnectedStaticOspfRouteMapSetMetric) matches(other *ConnectedStaticOspfRouteMapSetMetric) bool {
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

func (o *ConnectedStaticOspfv3) matches(other *ConnectedStaticOspfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *ConnectedStaticOspfv3RouteMap) matches(other *ConnectedStaticOspfv3RouteMap) bool {
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

func (o *ConnectedStaticOspfv3RouteMapMatch) matches(other *ConnectedStaticOspfv3RouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *ConnectedStaticOspfv3RouteMapMatchIpv6) matches(other *ConnectedStaticOspfv3RouteMapMatchIpv6) bool {
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

func (o *ConnectedStaticOspfv3RouteMapMatchIpv6Address) matches(other *ConnectedStaticOspfv3RouteMapMatchIpv6Address) bool {
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

func (o *ConnectedStaticOspfv3RouteMapMatchIpv6NextHop) matches(other *ConnectedStaticOspfv3RouteMapMatchIpv6NextHop) bool {
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

func (o *ConnectedStaticOspfv3RouteMapSet) matches(other *ConnectedStaticOspfv3RouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *ConnectedStaticOspfv3RouteMapSetMetric) matches(other *ConnectedStaticOspfv3RouteMapSetMetric) bool {
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

func (o *ConnectedStaticRib) matches(other *ConnectedStaticRib) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *ConnectedStaticRibRouteMap) matches(other *ConnectedStaticRibRouteMap) bool {
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

func (o *ConnectedStaticRibRouteMapMatch) matches(other *ConnectedStaticRibRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
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

func (o *ConnectedStaticRibRouteMapMatchIpv4) matches(other *ConnectedStaticRibRouteMapMatchIpv4) bool {
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

func (o *ConnectedStaticRibRouteMapMatchIpv4Address) matches(other *ConnectedStaticRibRouteMapMatchIpv4Address) bool {
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

func (o *ConnectedStaticRibRouteMapMatchIpv4NextHop) matches(other *ConnectedStaticRibRouteMapMatchIpv4NextHop) bool {
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

func (o *ConnectedStaticRibRouteMapMatchIpv6) matches(other *ConnectedStaticRibRouteMapMatchIpv6) bool {
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

func (o *ConnectedStaticRibRouteMapMatchIpv6Address) matches(other *ConnectedStaticRibRouteMapMatchIpv6Address) bool {
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

func (o *ConnectedStaticRibRouteMapMatchIpv6NextHop) matches(other *ConnectedStaticRibRouteMapMatchIpv6NextHop) bool {
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

func (o *ConnectedStaticRibRouteMapSet) matches(other *ConnectedStaticRibRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}

	return true
}

func (o *ConnectedStaticRip) matches(other *ConnectedStaticRip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *ConnectedStaticRipRouteMap) matches(other *ConnectedStaticRipRouteMap) bool {
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

func (o *ConnectedStaticRipRouteMapMatch) matches(other *ConnectedStaticRipRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}

	return true
}

func (o *ConnectedStaticRipRouteMapMatchIpv4) matches(other *ConnectedStaticRipRouteMapMatchIpv4) bool {
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

func (o *ConnectedStaticRipRouteMapMatchIpv4Address) matches(other *ConnectedStaticRipRouteMapMatchIpv4Address) bool {
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

func (o *ConnectedStaticRipRouteMapMatchIpv4NextHop) matches(other *ConnectedStaticRipRouteMapMatchIpv4NextHop) bool {
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

func (o *ConnectedStaticRipRouteMapSet) matches(other *ConnectedStaticRipRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.NextHop, other.NextHop) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *ConnectedStaticRipRouteMapSetMetric) matches(other *ConnectedStaticRipRouteMapSetMetric) bool {
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

func (o *Ospf) matches(other *Ospf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Rib.matches(other.Rib) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *OspfBgp) matches(other *OspfBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *OspfBgpRouteMap) matches(other *OspfBgpRouteMap) bool {
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

func (o *OspfBgpRouteMapMatch) matches(other *OspfBgpRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *OspfBgpRouteMapMatchAddress) matches(other *OspfBgpRouteMapMatchAddress) bool {
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

func (o *OspfBgpRouteMapMatchNextHop) matches(other *OspfBgpRouteMapMatchNextHop) bool {
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

func (o *OspfBgpRouteMapSet) matches(other *OspfBgpRouteMapSet) bool {
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
	if !util.OrderedListsMatch[int64](o.AspathPrepend, other.AspathPrepend) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *OspfBgpRouteMapSetAggregator) matches(other *OspfBgpRouteMapSetAggregator) bool {
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

func (o *OspfBgpRouteMapSetMetric) matches(other *OspfBgpRouteMapSetMetric) bool {
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

func (o *OspfBgpRouteMapSetIpv4) matches(other *OspfBgpRouteMapSetIpv4) bool {
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

func (o *OspfRib) matches(other *OspfRib) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *OspfRibRouteMap) matches(other *OspfRibRouteMap) bool {
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

func (o *OspfRibRouteMapMatch) matches(other *OspfRibRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *OspfRibRouteMapMatchAddress) matches(other *OspfRibRouteMapMatchAddress) bool {
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

func (o *OspfRibRouteMapMatchNextHop) matches(other *OspfRibRouteMapMatchNextHop) bool {
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

func (o *OspfRibRouteMapSet) matches(other *OspfRibRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}

	return true
}

func (o *OspfRip) matches(other *OspfRip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *OspfRipRouteMap) matches(other *OspfRipRouteMap) bool {
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

func (o *OspfRipRouteMapMatch) matches(other *OspfRipRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *OspfRipRouteMapMatchAddress) matches(other *OspfRipRouteMapMatchAddress) bool {
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

func (o *OspfRipRouteMapMatchNextHop) matches(other *OspfRipRouteMapMatchNextHop) bool {
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

func (o *OspfRipRouteMapSet) matches(other *OspfRipRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.NextHop, other.NextHop) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *OspfRipRouteMapSetMetric) matches(other *OspfRipRouteMapSetMetric) bool {
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

func (o *Ospfv3) matches(other *Ospfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Rib.matches(other.Rib) {
		return false
	}

	return true
}

func (o *Ospfv3Bgp) matches(other *Ospfv3Bgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *Ospfv3BgpRouteMap) matches(other *Ospfv3BgpRouteMap) bool {
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

func (o *Ospfv3BgpRouteMapMatch) matches(other *Ospfv3BgpRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *Ospfv3BgpRouteMapMatchAddress) matches(other *Ospfv3BgpRouteMapMatchAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}

	return true
}

func (o *Ospfv3BgpRouteMapMatchNextHop) matches(other *Ospfv3BgpRouteMapMatchNextHop) bool {
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

func (o *Ospfv3BgpRouteMapSet) matches(other *Ospfv3BgpRouteMapSet) bool {
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
	if !util.OrderedListsMatch[int64](o.AspathPrepend, other.AspathPrepend) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *Ospfv3BgpRouteMapSetAggregator) matches(other *Ospfv3BgpRouteMapSetAggregator) bool {
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

func (o *Ospfv3BgpRouteMapSetMetric) matches(other *Ospfv3BgpRouteMapSetMetric) bool {
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

func (o *Ospfv3BgpRouteMapSetIpv6) matches(other *Ospfv3BgpRouteMapSetIpv6) bool {
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

func (o *Ospfv3Rib) matches(other *Ospfv3Rib) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *Ospfv3RibRouteMap) matches(other *Ospfv3RibRouteMap) bool {
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

func (o *Ospfv3RibRouteMapMatch) matches(other *Ospfv3RibRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *Ospfv3RibRouteMapMatchAddress) matches(other *Ospfv3RibRouteMapMatchAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.PrefixList, other.PrefixList) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}

	return true
}

func (o *Ospfv3RibRouteMapMatchNextHop) matches(other *Ospfv3RibRouteMapMatchNextHop) bool {
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

func (o *Ospfv3RibRouteMapSet) matches(other *Ospfv3RibRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}

	return true
}

func (o *Rip) matches(other *Rip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Rib.matches(other.Rib) {
		return false
	}

	return true
}

func (o *RipBgp) matches(other *RipBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *RipBgpRouteMap) matches(other *RipBgpRouteMap) bool {
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

func (o *RipBgpRouteMapMatch) matches(other *RipBgpRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *RipBgpRouteMapMatchAddress) matches(other *RipBgpRouteMapMatchAddress) bool {
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

func (o *RipBgpRouteMapMatchNextHop) matches(other *RipBgpRouteMapMatchNextHop) bool {
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

func (o *RipBgpRouteMapSet) matches(other *RipBgpRouteMapSet) bool {
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
	if !util.OrderedListsMatch[int64](o.AspathPrepend, other.AspathPrepend) {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegularCommunity, other.RegularCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.LargeCommunity, other.LargeCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *RipBgpRouteMapSetAggregator) matches(other *RipBgpRouteMapSetAggregator) bool {
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

func (o *RipBgpRouteMapSetMetric) matches(other *RipBgpRouteMapSetMetric) bool {
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

func (o *RipBgpRouteMapSetIpv4) matches(other *RipBgpRouteMapSetIpv4) bool {
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

func (o *RipOspf) matches(other *RipOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *RipOspfRouteMap) matches(other *RipOspfRouteMap) bool {
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

func (o *RipOspfRouteMapMatch) matches(other *RipOspfRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *RipOspfRouteMapMatchAddress) matches(other *RipOspfRouteMapMatchAddress) bool {
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

func (o *RipOspfRouteMapMatchNextHop) matches(other *RipOspfRouteMapMatchNextHop) bool {
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

func (o *RipOspfRouteMapSet) matches(other *RipOspfRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Metric.matches(other.Metric) {
		return false
	}
	if !util.StringsMatch(o.MetricType, other.MetricType) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *RipOspfRouteMapSetMetric) matches(other *RipOspfRouteMapSetMetric) bool {
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

func (o *RipRib) matches(other *RipRib) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *RipRibRouteMap) matches(other *RipRibRouteMap) bool {
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

func (o *RipRibRouteMapMatch) matches(other *RipRibRouteMapMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Tag, other.Tag) {
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

func (o *RipRibRouteMapMatchAddress) matches(other *RipRibRouteMapMatchAddress) bool {
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

func (o *RipRibRouteMapMatchNextHop) matches(other *RipRibRouteMapMatchNextHop) bool {
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

func (o *RipRibRouteMapSet) matches(other *RipRibRouteMapSet) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
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
