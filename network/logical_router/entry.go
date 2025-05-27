package logical_router

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
	Suffix = []string{"network", "logical-router"}
)

type Entry struct {
	Name string
	Vrf  []Vrf

	Misc map[string][]generic.Xml
}

type Vrf struct {
	Name         string
	Interface    []string
	AdminDists   *VrfAdminDists
	RibFilter    *VrfRibFilter
	Bgp          *VrfBgp
	RoutingTable *VrfRoutingTable
	Ospf         *VrfOspf
	Ospfv3       *VrfOspfv3
	Ecmp         *VrfEcmp
	Multicast    *VrfMulticast
	Rip          *VrfRip
}
type VrfAdminDists struct {
	Static      *int64
	StaticIpv6  *int64
	OspfInter   *int64
	OspfIntra   *int64
	OspfExt     *int64
	Ospfv3Inter *int64
	Ospfv3Intra *int64
	Ospfv3Ext   *int64
	BgpInternal *int64
	BgpExternal *int64
	BgpLocal    *int64
	Rip         *int64
}
type VrfBgp struct {
	Enable                      *bool
	RouterId                    *string
	LocalAs                     *string
	InstallRoute                *bool
	EnforceFirstAs              *bool
	FastExternalFailover        *bool
	EcmpMultiAs                 *bool
	DefaultLocalPreference      *int64
	GracefulShutdown            *bool
	AlwaysAdvertiseNetworkRoute *bool
	Med                         *VrfBgpMed
	GracefulRestart             *VrfBgpGracefulRestart
	GlobalBfd                   *VrfBgpGlobalBfd
	RedistributionProfile       *VrfBgpRedistributionProfile
	AdvertiseNetwork            *VrfBgpAdvertiseNetwork
	PeerGroup                   []VrfBgpPeerGroup
	AggregateRoutes             []VrfBgpAggregateRoutes
}
type VrfBgpAdvertiseNetwork struct {
	Ipv4 *VrfBgpAdvertiseNetworkIpv4
	Ipv6 *VrfBgpAdvertiseNetworkIpv6
}
type VrfBgpAdvertiseNetworkIpv4 struct {
	Network []VrfBgpAdvertiseNetworkIpv4Network
}
type VrfBgpAdvertiseNetworkIpv4Network struct {
	Name      string
	Unicast   *bool
	Multicast *bool
	Backdoor  *bool
}
type VrfBgpAdvertiseNetworkIpv6 struct {
	Network []VrfBgpAdvertiseNetworkIpv6Network
}
type VrfBgpAdvertiseNetworkIpv6Network struct {
	Name    string
	Unicast *bool
}
type VrfBgpAggregateRoutes struct {
	Name        string
	Description *string
	Enable      *bool
	SummaryOnly *bool
	AsSet       *bool
	SameMed     *bool
	Type        *VrfBgpAggregateRoutesType
}
type VrfBgpAggregateRoutesType struct {
	Ipv4 *VrfBgpAggregateRoutesTypeIpv4
	Ipv6 *VrfBgpAggregateRoutesTypeIpv6
}
type VrfBgpAggregateRoutesTypeIpv4 struct {
	SummaryPrefix *string
	SuppressMap   *string
	AttributeMap  *string
}
type VrfBgpAggregateRoutesTypeIpv6 struct {
	SummaryPrefix *string
	SuppressMap   *string
	AttributeMap  *string
}
type VrfBgpGlobalBfd struct {
	Profile *string
}
type VrfBgpGracefulRestart struct {
	Enable             *bool
	StaleRouteTime     *int64
	MaxPeerRestartTime *int64
	LocalRestartTime   *int64
}
type VrfBgpMed struct {
	AlwaysCompareMed           *bool
	DeterministicMedComparison *bool
}
type VrfBgpPeerGroup struct {
	Name              string
	Enable            *bool
	Type              *VrfBgpPeerGroupType
	AddressFamily     *VrfBgpPeerGroupAddressFamily
	FilteringProfile  *VrfBgpPeerGroupFilteringProfile
	ConnectionOptions *VrfBgpPeerGroupConnectionOptions
	Peer              []VrfBgpPeerGroupPeer
}
type VrfBgpPeerGroupAddressFamily struct {
	Ipv4 *string
	Ipv6 *string
}
type VrfBgpPeerGroupConnectionOptions struct {
	Timers         *string
	Multihop       *int64
	Authentication *string
	Dampening      *string
}
type VrfBgpPeerGroupFilteringProfile struct {
	Ipv4 *string
	Ipv6 *string
}
type VrfBgpPeerGroupPeer struct {
	Name                          string
	Enable                        *bool
	Passive                       *bool
	PeerAs                        *string
	EnableSenderSideLoopDetection *bool
	Inherit                       *VrfBgpPeerGroupPeerInherit
	LocalAddress                  *VrfBgpPeerGroupPeerLocalAddress
	PeerAddress                   *VrfBgpPeerGroupPeerPeerAddress
	ConnectionOptions             *VrfBgpPeerGroupPeerConnectionOptions
	Bfd                           *VrfBgpPeerGroupPeerBfd
}
type VrfBgpPeerGroupPeerBfd struct {
	Profile *string
}
type VrfBgpPeerGroupPeerConnectionOptions struct {
	Timers         *string
	Multihop       *string
	Authentication *string
	Dampening      *string
}
type VrfBgpPeerGroupPeerInherit struct {
	Yes *VrfBgpPeerGroupPeerInheritYes
	No  *VrfBgpPeerGroupPeerInheritNo
}
type VrfBgpPeerGroupPeerInheritNo struct {
	AddressFamily    *VrfBgpPeerGroupPeerInheritNoAddressFamily
	FilteringProfile *VrfBgpPeerGroupPeerInheritNoFilteringProfile
}
type VrfBgpPeerGroupPeerInheritNoAddressFamily struct {
	Ipv4 *string
	Ipv6 *string
}
type VrfBgpPeerGroupPeerInheritNoFilteringProfile struct {
	Ipv4 *string
	Ipv6 *string
}
type VrfBgpPeerGroupPeerInheritYes struct {
}
type VrfBgpPeerGroupPeerLocalAddress struct {
	Interface *string
	Ip        *string
}
type VrfBgpPeerGroupPeerPeerAddress struct {
	Ip   *string
	Fqdn *string
}
type VrfBgpPeerGroupType struct {
	Ibgp *VrfBgpPeerGroupTypeIbgp
	Ebgp *VrfBgpPeerGroupTypeEbgp
}
type VrfBgpPeerGroupTypeEbgp struct {
}
type VrfBgpPeerGroupTypeIbgp struct {
}
type VrfBgpRedistributionProfile struct {
	Ipv4 *VrfBgpRedistributionProfileIpv4
	Ipv6 *VrfBgpRedistributionProfileIpv6
}
type VrfBgpRedistributionProfileIpv4 struct {
	Unicast *string
}
type VrfBgpRedistributionProfileIpv6 struct {
	Unicast *string
}
type VrfEcmp struct {
	Enable           *bool
	MaxPath          *int64
	SymmetricReturn  *bool
	StrictSourcePath *bool
	Algorithm        *VrfEcmpAlgorithm
}
type VrfEcmpAlgorithm struct {
	IpModulo           *VrfEcmpAlgorithmIpModulo
	IpHash             *VrfEcmpAlgorithmIpHash
	WeightedRoundRobin *VrfEcmpAlgorithmWeightedRoundRobin
	BalancedRoundRobin *VrfEcmpAlgorithmBalancedRoundRobin
}
type VrfEcmpAlgorithmBalancedRoundRobin struct {
}
type VrfEcmpAlgorithmIpHash struct {
	SrcOnly  *bool
	UsePort  *bool
	HashSeed *int64
}
type VrfEcmpAlgorithmIpModulo struct {
}
type VrfEcmpAlgorithmWeightedRoundRobin struct {
	Interface []VrfEcmpAlgorithmWeightedRoundRobinInterface
}
type VrfEcmpAlgorithmWeightedRoundRobinInterface struct {
	Name   string
	Weight *int64
}
type VrfMulticast struct {
	Enable      *bool
	StaticRoute []VrfMulticastStaticRoute
	Pim         *VrfMulticastPim
	Igmp        *VrfMulticastIgmp
	Msdp        *VrfMulticastMsdp
}
type VrfMulticastIgmp struct {
	Enable  *bool
	Dynamic *VrfMulticastIgmpDynamic
	Static  []VrfMulticastIgmpStatic
}
type VrfMulticastIgmpDynamic struct {
	Interface []VrfMulticastIgmpDynamicInterface
}
type VrfMulticastIgmpDynamicInterface struct {
	Name                string
	Version             *string
	Robustness          *string
	GroupFilter         *string
	MaxGroups           *string
	MaxSources          *string
	QueryProfile        *string
	RouterAlertPolicing *bool
}
type VrfMulticastIgmpStatic struct {
	Name          string
	Interface     *string
	GroupAddress  *string
	SourceAddress *string
}
type VrfMulticastMsdp struct {
	Enable               *bool
	GlobalTimer          *string
	GlobalAuthentication *string
	OriginatorId         *VrfMulticastMsdpOriginatorId
	Peer                 []VrfMulticastMsdpPeer
}
type VrfMulticastMsdpOriginatorId struct {
	Interface *string
	Ip        *string
}
type VrfMulticastMsdpPeer struct {
	Name             string
	Enable           *bool
	PeerAs           *string
	Authentication   *string
	MaxSa            *int64
	InboundSaFilter  *string
	OutboundSaFilter *string
	LocalAddress     *VrfMulticastMsdpPeerLocalAddress
	PeerAddress      *VrfMulticastMsdpPeerPeerAddress
}
type VrfMulticastMsdpPeerLocalAddress struct {
	Interface *string
	Ip        *string
}
type VrfMulticastMsdpPeerPeerAddress struct {
	Ip   *string
	Fqdn *string
}
type VrfMulticastPim struct {
	Enable          *bool
	RpfLookupMode   *string
	RouteAgeoutTime *int64
	IfTimerGlobal   *string
	GroupPermission *string
	SsmAddressSpace *VrfMulticastPimSsmAddressSpace
	Rp              *VrfMulticastPimRp
	SptThreshold    []VrfMulticastPimSptThreshold
	Interface       []VrfMulticastPimInterface
}
type VrfMulticastPimInterface struct {
	Name           string
	Description    *string
	DrPriority     *int64
	SendBsm        *bool
	IfTimer        *string
	NeighborFilter *string
}
type VrfMulticastPimRp struct {
	LocalRp    *VrfMulticastPimRpLocalRp
	ExternalRp []VrfMulticastPimRpExternalRp
}
type VrfMulticastPimRpExternalRp struct {
	Name      string
	GroupList *string
	Override  *bool
}
type VrfMulticastPimRpLocalRp struct {
	StaticRp    *VrfMulticastPimRpLocalRpStaticRp
	CandidateRp *VrfMulticastPimRpLocalRpCandidateRp
}
type VrfMulticastPimRpLocalRpCandidateRp struct {
	Interface             *string
	Address               *string
	Priority              *int64
	AdvertisementInterval *int64
	GroupList             *string
}
type VrfMulticastPimRpLocalRpStaticRp struct {
	Interface *string
	Address   *string
	Override  *bool
	GroupList *string
}
type VrfMulticastPimSptThreshold struct {
	Name      string
	Threshold *string
}
type VrfMulticastPimSsmAddressSpace struct {
	GroupList *string
}
type VrfMulticastStaticRoute struct {
	Name        string
	Destination *string
	Interface   *string
	Preference  *int64
	Nexthop     *VrfMulticastStaticRouteNexthop
}
type VrfMulticastStaticRouteNexthop struct {
	IpAddress *string
}
type VrfOspf struct {
	RouterId              *string
	Enable                *bool
	Rfc1583               *bool
	SpfTimer              *string
	GlobalIfTimer         *string
	RedistributionProfile *string
	GlobalBfd             *VrfOspfGlobalBfd
	GracefulRestart       *VrfOspfGracefulRestart
	Area                  []VrfOspfArea
}
type VrfOspfArea struct {
	Name           string
	Authentication *string
	Type           *VrfOspfAreaType
	Range          []VrfOspfAreaRange
	Interface      []VrfOspfAreaInterface
	VirtualLink    []VrfOspfAreaVirtualLink
}
type VrfOspfAreaInterface struct {
	Name           string
	Enable         *bool
	MtuIgnore      *bool
	Passive        *bool
	Priority       *int64
	Metric         *int64
	Authentication *string
	Timing         *string
	LinkType       *VrfOspfAreaInterfaceLinkType
	Bfd            *VrfOspfAreaInterfaceBfd
}
type VrfOspfAreaInterfaceBfd struct {
	Profile *string
}
type VrfOspfAreaInterfaceLinkType struct {
	Broadcast *VrfOspfAreaInterfaceLinkTypeBroadcast
	P2p       *VrfOspfAreaInterfaceLinkTypeP2p
	P2mp      *VrfOspfAreaInterfaceLinkTypeP2mp
}
type VrfOspfAreaInterfaceLinkTypeBroadcast struct {
}
type VrfOspfAreaInterfaceLinkTypeP2mp struct {
	Neighbor []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor
}
type VrfOspfAreaInterfaceLinkTypeP2mpNeighbor struct {
	Name     string
	Priority *int64
}
type VrfOspfAreaInterfaceLinkTypeP2p struct {
}
type VrfOspfAreaRange struct {
	Name      string
	Advertise *bool
}
type VrfOspfAreaType struct {
	Normal *VrfOspfAreaTypeNormal
	Stub   *VrfOspfAreaTypeStub
	Nssa   *VrfOspfAreaTypeNssa
}
type VrfOspfAreaTypeNormal struct {
	Abr *VrfOspfAreaTypeNormalAbr
}
type VrfOspfAreaTypeNormalAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
}
type VrfOspfAreaTypeNssa struct {
	NoSummary                   *bool
	DefaultInformationOriginate *VrfOspfAreaTypeNssaDefaultInformationOriginate
	Abr                         *VrfOspfAreaTypeNssaAbr
}
type VrfOspfAreaTypeNssaAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	NssaExtRange       []VrfOspfAreaTypeNssaAbrNssaExtRange
}
type VrfOspfAreaTypeNssaAbrNssaExtRange struct {
	Name      string
	Advertise *bool
}
type VrfOspfAreaTypeNssaDefaultInformationOriginate struct {
	Metric     *int64
	MetricType *string
}
type VrfOspfAreaTypeStub struct {
	NoSummary          *bool
	Abr                *VrfOspfAreaTypeStubAbr
	DefaultRouteMetric *int64
}
type VrfOspfAreaTypeStubAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
}
type VrfOspfAreaVirtualLink struct {
	Name           string
	NeighborId     *string
	TransitAreaId  *string
	Enable         *bool
	InstanceId     *int64
	Timing         *string
	Authentication *string
	Bfd            *VrfOspfAreaVirtualLinkBfd
}
type VrfOspfAreaVirtualLinkBfd struct {
	Profile *string
}
type VrfOspfGlobalBfd struct {
	Profile *string
}
type VrfOspfGracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	StrictLSAChecking      *bool
	MaxNeighborRestartTime *int64
}
type VrfOspfv3 struct {
	Enable                *bool
	RouterId              *string
	DisableTransitTraffic *bool
	SpfTimer              *string
	GlobalIfTimer         *string
	RedistributionProfile *string
	GlobalBfd             *VrfOspfv3GlobalBfd
	GracefulRestart       *VrfOspfv3GracefulRestart
	Area                  []VrfOspfv3Area
}
type VrfOspfv3Area struct {
	Name           string
	Authentication *string
	Type           *VrfOspfv3AreaType
	Range          []VrfOspfv3AreaRange
	Interface      []VrfOspfv3AreaInterface
	VirtualLink    []VrfOspfv3AreaVirtualLink
}
type VrfOspfv3AreaInterface struct {
	Name           string
	Enable         *bool
	MtuIgnore      *bool
	Passive        *bool
	Priority       *int64
	Metric         *int64
	InstanceId     *int64
	Authentication *string
	Timing         *string
	LinkType       *VrfOspfv3AreaInterfaceLinkType
	Bfd            *VrfOspfv3AreaInterfaceBfd
}
type VrfOspfv3AreaInterfaceBfd struct {
	Profile *string
}
type VrfOspfv3AreaInterfaceLinkType struct {
	Broadcast *VrfOspfv3AreaInterfaceLinkTypeBroadcast
	P2p       *VrfOspfv3AreaInterfaceLinkTypeP2p
	P2mp      *VrfOspfv3AreaInterfaceLinkTypeP2mp
}
type VrfOspfv3AreaInterfaceLinkTypeBroadcast struct {
}
type VrfOspfv3AreaInterfaceLinkTypeP2mp struct {
	Neighbor []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor struct {
	Name     string
	Priority *int64
}
type VrfOspfv3AreaInterfaceLinkTypeP2p struct {
}
type VrfOspfv3AreaRange struct {
	Name      string
	Advertise *bool
}
type VrfOspfv3AreaType struct {
	Normal *VrfOspfv3AreaTypeNormal
	Stub   *VrfOspfv3AreaTypeStub
	Nssa   *VrfOspfv3AreaTypeNssa
}
type VrfOspfv3AreaTypeNormal struct {
	Abr *VrfOspfv3AreaTypeNormalAbr
}
type VrfOspfv3AreaTypeNormalAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
}
type VrfOspfv3AreaTypeNssa struct {
	NoSummary                   *bool
	DefaultInformationOriginate *VrfOspfv3AreaTypeNssaDefaultInformationOriginate
	Abr                         *VrfOspfv3AreaTypeNssaAbr
}
type VrfOspfv3AreaTypeNssaAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	NssaExtRange       []VrfOspfv3AreaTypeNssaAbrNssaExtRange
}
type VrfOspfv3AreaTypeNssaAbrNssaExtRange struct {
	Name      string
	Advertise *bool
}
type VrfOspfv3AreaTypeNssaDefaultInformationOriginate struct {
	Metric     *int64
	MetricType *string
}
type VrfOspfv3AreaTypeStub struct {
	NoSummary          *bool
	Abr                *VrfOspfv3AreaTypeStubAbr
	DefaultRouteMetric *int64
}
type VrfOspfv3AreaTypeStubAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
}
type VrfOspfv3AreaVirtualLink struct {
	Name           string
	NeighborId     *string
	TransitAreaId  *string
	Enable         *bool
	InstanceId     *int64
	Timing         *string
	Authentication *string
}
type VrfOspfv3GlobalBfd struct {
	Profile *string
}
type VrfOspfv3GracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	StrictLSAChecking      *bool
	MaxNeighborRestartTime *int64
}
type VrfRibFilter struct {
	Ipv4 *VrfRibFilterIpv4
	Ipv6 *VrfRibFilterIpv6
}
type VrfRibFilterIpv4 struct {
	Static *VrfRibFilterIpv4Static
	Bgp    *VrfRibFilterIpv4Bgp
	Ospf   *VrfRibFilterIpv4Ospf
	Rip    *VrfRibFilterIpv4Rip
}
type VrfRibFilterIpv4Bgp struct {
	RouteMap *string
}
type VrfRibFilterIpv4Ospf struct {
	RouteMap *string
}
type VrfRibFilterIpv4Rip struct {
	RouteMap *string
}
type VrfRibFilterIpv4Static struct {
	RouteMap *string
}
type VrfRibFilterIpv6 struct {
	Static *VrfRibFilterIpv6Static
	Bgp    *VrfRibFilterIpv6Bgp
	Ospfv3 *VrfRibFilterIpv6Ospfv3
}
type VrfRibFilterIpv6Bgp struct {
	RouteMap *string
}
type VrfRibFilterIpv6Ospfv3 struct {
	RouteMap *string
}
type VrfRibFilterIpv6Static struct {
	RouteMap *string
}
type VrfRip struct {
	Enable                       *bool
	DefaultInformationOriginate  *bool
	GlobalTimer                  *string
	AuthProfile                  *string
	RedistributionProfile        *string
	GlobalBfd                    *VrfRipGlobalBfd
	GlobalInboundDistributeList  *VrfRipGlobalInboundDistributeList
	GlobalOutboundDistributeList *VrfRipGlobalOutboundDistributeList
	Interface                    []VrfRipInterface
}
type VrfRipGlobalBfd struct {
	Profile *string
}
type VrfRipGlobalInboundDistributeList struct {
	AccessList *string
}
type VrfRipGlobalOutboundDistributeList struct {
	AccessList *string
}
type VrfRipInterface struct {
	Name                            string
	Enable                          *bool
	Mode                            *string
	SplitHorizon                    *string
	Authentication                  *string
	Bfd                             *VrfRipInterfaceBfd
	InterfaceInboundDistributeList  *VrfRipInterfaceInterfaceInboundDistributeList
	InterfaceOutboundDistributeList *VrfRipInterfaceInterfaceOutboundDistributeList
}
type VrfRipInterfaceBfd struct {
	Profile *string
}
type VrfRipInterfaceInterfaceInboundDistributeList struct {
	AccessList *string
	Metric     *int64
}
type VrfRipInterfaceInterfaceOutboundDistributeList struct {
	AccessList *string
	Metric     *int64
}
type VrfRoutingTable struct {
	Ip   *VrfRoutingTableIp
	Ipv6 *VrfRoutingTableIpv6
}
type VrfRoutingTableIp struct {
	StaticRoute []VrfRoutingTableIpStaticRoute
}
type VrfRoutingTableIpStaticRoute struct {
	Name        string
	Destination *string
	Interface   *string
	AdminDist   *int64
	Metric      *int64
	Nexthop     *VrfRoutingTableIpStaticRouteNexthop
	Bfd         *VrfRoutingTableIpStaticRouteBfd
	PathMonitor *VrfRoutingTableIpStaticRoutePathMonitor
}
type VrfRoutingTableIpStaticRouteBfd struct {
	Profile *string
}
type VrfRoutingTableIpStaticRouteNexthop struct {
	Discard   *VrfRoutingTableIpStaticRouteNexthopDiscard
	IpAddress *string
	NextLr    *string
	Fqdn      *string
}
type VrfRoutingTableIpStaticRouteNexthopDiscard struct {
}
type VrfRoutingTableIpStaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations
}
type VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations struct {
	Name        string
	Enable      *bool
	Source      *string
	Destination *string
	Interval    *int64
	Count       *int64
}
type VrfRoutingTableIpv6 struct {
	StaticRoute []VrfRoutingTableIpv6StaticRoute
}
type VrfRoutingTableIpv6StaticRoute struct {
	Name        string
	Destination *string
	Interface   *string
	AdminDist   *int64
	Metric      *int64
	Nexthop     *VrfRoutingTableIpv6StaticRouteNexthop
	Bfd         *VrfRoutingTableIpv6StaticRouteBfd
	PathMonitor *VrfRoutingTableIpv6StaticRoutePathMonitor
}
type VrfRoutingTableIpv6StaticRouteBfd struct {
	Profile *string
}
type VrfRoutingTableIpv6StaticRouteNexthop struct {
	Discard     *VrfRoutingTableIpv6StaticRouteNexthopDiscard
	Ipv6Address *string
	Fqdn        *string
	NextLr      *string
}
type VrfRoutingTableIpv6StaticRouteNexthopDiscard struct {
}
type VrfRoutingTableIpv6StaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations
}
type VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations struct {
	Name        string
	Enable      *bool
	Source      *string
	Destination *string
	Interval    *int64
	Count       *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}
type entryXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Vrf     []VrfXml `xml:"vrf>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName xml.Name        `xml:"entry"`
	Name    string          `xml:"name,attr"`
	Vrf     []VrfXml_11_0_2 `xml:"vrf>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfXml struct {
	AdminDists   *VrfAdminDistsXml   `xml:"admin-dists,omitempty"`
	Bgp          *VrfBgpXml          `xml:"bgp,omitempty"`
	Ecmp         *VrfEcmpXml         `xml:"ecmp,omitempty"`
	Interface    *util.MemberType    `xml:"interface,omitempty"`
	Multicast    *VrfMulticastXml    `xml:"multicast,omitempty"`
	Name         string              `xml:"name,attr"`
	Ospf         *VrfOspfXml         `xml:"ospf,omitempty"`
	Ospfv3       *VrfOspfv3Xml       `xml:"ospfv3,omitempty"`
	RibFilter    *VrfRibFilterXml    `xml:"rib-filter,omitempty"`
	Rip          *VrfRipXml          `xml:"rip,omitempty"`
	RoutingTable *VrfRoutingTableXml `xml:"routing-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfAdminDistsXml struct {
	BgpExternal *int64 `xml:"bgp-external,omitempty"`
	BgpInternal *int64 `xml:"bgp-internal,omitempty"`
	BgpLocal    *int64 `xml:"bgp-local,omitempty"`
	OspfExt     *int64 `xml:"ospf-ext,omitempty"`
	OspfInter   *int64 `xml:"ospf-inter,omitempty"`
	OspfIntra   *int64 `xml:"ospf-intra,omitempty"`
	Ospfv3Ext   *int64 `xml:"ospfv3-ext,omitempty"`
	Ospfv3Inter *int64 `xml:"ospfv3-inter,omitempty"`
	Ospfv3Intra *int64 `xml:"ospfv3-intra,omitempty"`
	Rip         *int64 `xml:"rip,omitempty"`
	Static      *int64 `xml:"static,omitempty"`
	StaticIpv6  *int64 `xml:"static-ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpXml struct {
	AdvertiseNetwork            *VrfBgpAdvertiseNetworkXml      `xml:"advertise-network,omitempty"`
	AggregateRoutes             []VrfBgpAggregateRoutesXml      `xml:"aggregate-routes>entry,omitempty"`
	AlwaysAdvertiseNetworkRoute *string                         `xml:"always-advertise-network-route,omitempty"`
	DefaultLocalPreference      *int64                          `xml:"default-local-preference,omitempty"`
	EcmpMultiAs                 *string                         `xml:"ecmp-multi-as,omitempty"`
	Enable                      *string                         `xml:"enable,omitempty"`
	EnforceFirstAs              *string                         `xml:"enforce-first-as,omitempty"`
	FastExternalFailover        *string                         `xml:"fast-external-failover,omitempty"`
	GlobalBfd                   *VrfBgpGlobalBfdXml             `xml:"global-bfd,omitempty"`
	GracefulRestart             *VrfBgpGracefulRestartXml       `xml:"graceful-restart,omitempty"`
	GracefulShutdown            *string                         `xml:"graceful-shutdown,omitempty"`
	InstallRoute                *string                         `xml:"install-route,omitempty"`
	LocalAs                     *string                         `xml:"local-as,omitempty"`
	Med                         *VrfBgpMedXml                   `xml:"med,omitempty"`
	PeerGroup                   []VrfBgpPeerGroupXml            `xml:"peer-group>entry,omitempty"`
	RedistributionProfile       *VrfBgpRedistributionProfileXml `xml:"redistribution-profile,omitempty"`
	RouterId                    *string                         `xml:"router-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkXml struct {
	Ipv4 *VrfBgpAdvertiseNetworkIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpAdvertiseNetworkIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv4Xml struct {
	Network []VrfBgpAdvertiseNetworkIpv4NetworkXml `xml:"network>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv4NetworkXml struct {
	Backdoor  *string `xml:"backdoor,omitempty"`
	Multicast *string `xml:"multicast,omitempty"`
	Name      string  `xml:"name,attr"`
	Unicast   *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv6Xml struct {
	Network []VrfBgpAdvertiseNetworkIpv6NetworkXml `xml:"network>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv6NetworkXml struct {
	Name    string  `xml:"name,attr"`
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesXml struct {
	AsSet       *string                       `xml:"as-set,omitempty"`
	Description *string                       `xml:"description,omitempty"`
	Enable      *string                       `xml:"enable,omitempty"`
	Name        string                        `xml:"name,attr"`
	SameMed     *string                       `xml:"same-med,omitempty"`
	SummaryOnly *string                       `xml:"summary-only,omitempty"`
	Type        *VrfBgpAggregateRoutesTypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeXml struct {
	Ipv4 *VrfBgpAggregateRoutesTypeIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpAggregateRoutesTypeIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeIpv4Xml struct {
	AttributeMap  *string `xml:"attribute-map,omitempty"`
	SummaryPrefix *string `xml:"summary-prefix,omitempty"`
	SuppressMap   *string `xml:"suppress-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeIpv6Xml struct {
	AttributeMap  *string `xml:"attribute-map,omitempty"`
	SummaryPrefix *string `xml:"summary-prefix,omitempty"`
	SuppressMap   *string `xml:"suppress-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpGracefulRestartXml struct {
	Enable             *string `xml:"enable,omitempty"`
	LocalRestartTime   *int64  `xml:"local-restart-time,omitempty"`
	MaxPeerRestartTime *int64  `xml:"max-peer-restart-time,omitempty"`
	StaleRouteTime     *int64  `xml:"stale-route-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpMedXml struct {
	AlwaysCompareMed           *string `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string `xml:"deterministic-med-comparison,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupXml struct {
	AddressFamily     *VrfBgpPeerGroupAddressFamilyXml     `xml:"address-family,omitempty"`
	ConnectionOptions *VrfBgpPeerGroupConnectionOptionsXml `xml:"connection-options,omitempty"`
	Enable            *string                              `xml:"enable,omitempty"`
	FilteringProfile  *VrfBgpPeerGroupFilteringProfileXml  `xml:"filtering-profile,omitempty"`
	Name              string                               `xml:"name,attr"`
	Peer              []VrfBgpPeerGroupPeerXml             `xml:"peer>entry,omitempty"`
	Type              *VrfBgpPeerGroupTypeXml              `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupAddressFamilyXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupConnectionOptionsXml struct {
	Authentication *string `xml:"authentication,omitempty"`
	Dampening      *string `xml:"dampening,omitempty"`
	Multihop       *int64  `xml:"multihop,omitempty"`
	Timers         *string `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupFilteringProfileXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerXml struct {
	Bfd                           *VrfBgpPeerGroupPeerBfdXml               `xml:"bfd,omitempty"`
	ConnectionOptions             *VrfBgpPeerGroupPeerConnectionOptionsXml `xml:"connection-options,omitempty"`
	Enable                        *string                                  `xml:"enable,omitempty"`
	EnableSenderSideLoopDetection *string                                  `xml:"enable-sender-side-loop-detection,omitempty"`
	Inherit                       *VrfBgpPeerGroupPeerInheritXml           `xml:"inherit,omitempty"`
	LocalAddress                  *VrfBgpPeerGroupPeerLocalAddressXml      `xml:"local-address,omitempty"`
	Name                          string                                   `xml:"name,attr"`
	Passive                       *string                                  `xml:"passive,omitempty"`
	PeerAddress                   *VrfBgpPeerGroupPeerPeerAddressXml       `xml:"peer-address,omitempty"`
	PeerAs                        *string                                  `xml:"peer-as,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerConnectionOptionsXml struct {
	Authentication *string `xml:"authentication,omitempty"`
	Dampening      *string `xml:"dampening,omitempty"`
	Multihop       *string `xml:"multihop,omitempty"`
	Timers         *string `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritXml struct {
	No  *VrfBgpPeerGroupPeerInheritNoXml  `xml:"no,omitempty"`
	Yes *VrfBgpPeerGroupPeerInheritYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoXml struct {
	AddressFamily    *VrfBgpPeerGroupPeerInheritNoAddressFamilyXml    `xml:"address-family,omitempty"`
	FilteringProfile *VrfBgpPeerGroupPeerInheritNoFilteringProfileXml `xml:"filtering-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoAddressFamilyXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoFilteringProfileXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritYesXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerLocalAddressXml struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerPeerAddressXml struct {
	Fqdn *string `xml:"fqdn,omitempty"`
	Ip   *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeXml struct {
	Ebgp *VrfBgpPeerGroupTypeEbgpXml `xml:"ebgp,omitempty"`
	Ibgp *VrfBgpPeerGroupTypeIbgpXml `xml:"ibgp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeEbgpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeIbgpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileXml struct {
	Ipv4 *VrfBgpRedistributionProfileIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpRedistributionProfileIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileIpv4Xml struct {
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileIpv6Xml struct {
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpXml struct {
	Algorithm        *VrfEcmpAlgorithmXml `xml:"algorithm,omitempty"`
	Enable           *string              `xml:"enable,omitempty"`
	MaxPath          *int64               `xml:"max-path,omitempty"`
	StrictSourcePath *string              `xml:"strict-source-path,omitempty"`
	SymmetricReturn  *string              `xml:"symmetric-return,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmXml struct {
	BalancedRoundRobin *VrfEcmpAlgorithmBalancedRoundRobinXml `xml:"balanced-round-robin,omitempty"`
	IpHash             *VrfEcmpAlgorithmIpHashXml             `xml:"ip-hash,omitempty"`
	IpModulo           *VrfEcmpAlgorithmIpModuloXml           `xml:"ip-modulo,omitempty"`
	WeightedRoundRobin *VrfEcmpAlgorithmWeightedRoundRobinXml `xml:"weighted-round-robin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmBalancedRoundRobinXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmIpHashXml struct {
	HashSeed *int64  `xml:"hash-seed,omitempty"`
	SrcOnly  *string `xml:"src-only,omitempty"`
	UsePort  *string `xml:"use-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmIpModuloXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmWeightedRoundRobinXml struct {
	Interface []VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml struct {
	Name   string `xml:"name,attr"`
	Weight *int64 `xml:"weight,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastXml struct {
	Enable      *string                      `xml:"enable,omitempty"`
	Igmp        *VrfMulticastIgmpXml         `xml:"igmp,omitempty"`
	Msdp        *VrfMulticastMsdpXml         `xml:"msdp,omitempty"`
	Pim         *VrfMulticastPimXml          `xml:"pim,omitempty"`
	StaticRoute []VrfMulticastStaticRouteXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpXml struct {
	Dynamic *VrfMulticastIgmpDynamicXml `xml:"dynamic,omitempty"`
	Enable  *string                     `xml:"enable,omitempty"`
	Static  []VrfMulticastIgmpStaticXml `xml:"static>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpDynamicXml struct {
	Interface []VrfMulticastIgmpDynamicInterfaceXml `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpDynamicInterfaceXml struct {
	GroupFilter         *string `xml:"group-filter,omitempty"`
	MaxGroups           *string `xml:"max-groups,omitempty"`
	MaxSources          *string `xml:"max-sources,omitempty"`
	Name                string  `xml:"name,attr"`
	QueryProfile        *string `xml:"query-profile,omitempty"`
	Robustness          *string `xml:"robustness,omitempty"`
	RouterAlertPolicing *string `xml:"router-alert-policing,omitempty"`
	Version             *string `xml:"version,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpStaticXml struct {
	GroupAddress  *string `xml:"group-address,omitempty"`
	Interface     *string `xml:"interface,omitempty"`
	Name          string  `xml:"name,attr"`
	SourceAddress *string `xml:"source-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpXml struct {
	Enable               *string                          `xml:"enable,omitempty"`
	GlobalAuthentication *string                          `xml:"global-authentication,omitempty"`
	GlobalTimer          *string                          `xml:"global-timer,omitempty"`
	OriginatorId         *VrfMulticastMsdpOriginatorIdXml `xml:"originator-id,omitempty"`
	Peer                 []VrfMulticastMsdpPeerXml        `xml:"peer>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpOriginatorIdXml struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerXml struct {
	Authentication   *string                              `xml:"authentication,omitempty"`
	Enable           *string                              `xml:"enable,omitempty"`
	InboundSaFilter  *string                              `xml:"inbound-sa-filter,omitempty"`
	LocalAddress     *VrfMulticastMsdpPeerLocalAddressXml `xml:"local-address,omitempty"`
	MaxSa            *int64                               `xml:"max-sa,omitempty"`
	Name             string                               `xml:"name,attr"`
	OutboundSaFilter *string                              `xml:"outbound-sa-filter,omitempty"`
	PeerAddress      *VrfMulticastMsdpPeerPeerAddressXml  `xml:"peer-address,omitempty"`
	PeerAs           *string                              `xml:"peer-as,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerLocalAddressXml struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerPeerAddressXml struct {
	Fqdn *string `xml:"fqdn,omitempty"`
	Ip   *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimXml struct {
	Enable          *string                            `xml:"enable,omitempty"`
	GroupPermission *string                            `xml:"group-permission,omitempty"`
	IfTimerGlobal   *string                            `xml:"if-timer-global,omitempty"`
	Interface       []VrfMulticastPimInterfaceXml      `xml:"interface>entry,omitempty"`
	RouteAgeoutTime *int64                             `xml:"route-ageout-time,omitempty"`
	Rp              *VrfMulticastPimRpXml              `xml:"rp,omitempty"`
	RpfLookupMode   *string                            `xml:"rpf-lookup-mode,omitempty"`
	SptThreshold    []VrfMulticastPimSptThresholdXml   `xml:"spt-threshold>entry,omitempty"`
	SsmAddressSpace *VrfMulticastPimSsmAddressSpaceXml `xml:"ssm-address-space,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimInterfaceXml struct {
	Description    *string `xml:"description,omitempty"`
	DrPriority     *int64  `xml:"dr-priority,omitempty"`
	IfTimer        *string `xml:"if-timer,omitempty"`
	Name           string  `xml:"name,attr"`
	NeighborFilter *string `xml:"neighbor-filter,omitempty"`
	SendBsm        *string `xml:"send-bsm,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpXml struct {
	ExternalRp []VrfMulticastPimRpExternalRpXml `xml:"external-rp>entry,omitempty"`
	LocalRp    *VrfMulticastPimRpLocalRpXml     `xml:"local-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpExternalRpXml struct {
	GroupList *string `xml:"group-list,omitempty"`
	Name      string  `xml:"name,attr"`
	Override  *string `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpXml struct {
	CandidateRp *VrfMulticastPimRpLocalRpCandidateRpXml `xml:"candidate-rp,omitempty"`
	StaticRp    *VrfMulticastPimRpLocalRpStaticRpXml    `xml:"static-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpCandidateRpXml struct {
	Address               *string `xml:"address,omitempty"`
	AdvertisementInterval *int64  `xml:"advertisement-interval,omitempty"`
	GroupList             *string `xml:"group-list,omitempty"`
	Interface             *string `xml:"interface,omitempty"`
	Priority              *int64  `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpStaticRpXml struct {
	Address   *string `xml:"address,omitempty"`
	GroupList *string `xml:"group-list,omitempty"`
	Interface *string `xml:"interface,omitempty"`
	Override  *string `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimSptThresholdXml struct {
	Name      string  `xml:"name,attr"`
	Threshold *string `xml:"threshold,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimSsmAddressSpaceXml struct {
	GroupList *string `xml:"group-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastStaticRouteXml struct {
	Destination *string                            `xml:"destination,omitempty"`
	Interface   *string                            `xml:"interface,omitempty"`
	Name        string                             `xml:"name,attr"`
	Nexthop     *VrfMulticastStaticRouteNexthopXml `xml:"nexthop,omitempty"`
	Preference  *int64                             `xml:"preference,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastStaticRouteNexthopXml struct {
	IpAddress *string `xml:"ip-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfXml struct {
	Area                  []VrfOspfAreaXml           `xml:"area>entry,omitempty"`
	Enable                *string                    `xml:"enable,omitempty"`
	GlobalBfd             *VrfOspfGlobalBfdXml       `xml:"global-bfd,omitempty"`
	GlobalIfTimer         *string                    `xml:"global-if-timer,omitempty"`
	GracefulRestart       *VrfOspfGracefulRestartXml `xml:"graceful-restart,omitempty"`
	RedistributionProfile *string                    `xml:"redistribution-profile,omitempty"`
	Rfc1583               *string                    `xml:"rfc1583,omitempty"`
	RouterId              *string                    `xml:"router-id,omitempty"`
	SpfTimer              *string                    `xml:"spf-timer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaXml struct {
	Authentication *string                     `xml:"authentication,omitempty"`
	Interface      []VrfOspfAreaInterfaceXml   `xml:"interface>entry,omitempty"`
	Name           string                      `xml:"name,attr"`
	Range          []VrfOspfAreaRangeXml       `xml:"range>entry,omitempty"`
	Type           *VrfOspfAreaTypeXml         `xml:"type,omitempty"`
	VirtualLink    []VrfOspfAreaVirtualLinkXml `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceXml struct {
	Authentication *string                          `xml:"authentication,omitempty"`
	Bfd            *VrfOspfAreaInterfaceBfdXml      `xml:"bfd,omitempty"`
	Enable         *string                          `xml:"enable,omitempty"`
	LinkType       *VrfOspfAreaInterfaceLinkTypeXml `xml:"link-type,omitempty"`
	Metric         *int64                           `xml:"metric,omitempty"`
	MtuIgnore      *string                          `xml:"mtu-ignore,omitempty"`
	Name           string                           `xml:"name,attr"`
	Passive        *string                          `xml:"passive,omitempty"`
	Priority       *int64                           `xml:"priority,omitempty"`
	Timing         *string                          `xml:"timing,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeXml struct {
	Broadcast *VrfOspfAreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2mp      *VrfOspfAreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	P2p       *VrfOspfAreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2mpXml struct {
	Neighbor []VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml `xml:"neighbor>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml struct {
	Name     string `xml:"name,attr"`
	Priority *int64 `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaRangeXml struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeXml struct {
	Normal *VrfOspfAreaTypeNormalXml `xml:"normal,omitempty"`
	Nssa   *VrfOspfAreaTypeNssaXml   `xml:"nssa,omitempty"`
	Stub   *VrfOspfAreaTypeStubXml   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNormalXml struct {
	Abr *VrfOspfAreaTypeNormalAbrXml `xml:"abr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNormalAbrXml struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaXml struct {
	Abr                         *VrfOspfAreaTypeNssaAbrXml                         `xml:"abr,omitempty"`
	DefaultInformationOriginate *VrfOspfAreaTypeNssaDefaultInformationOriginateXml `xml:"default-information-originate,omitempty"`
	NoSummary                   *string                                            `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaAbrXml struct {
	ExportList         *string                                 `xml:"export-list,omitempty"`
	ImportList         *string                                 `xml:"import-list,omitempty"`
	InboundFilterList  *string                                 `xml:"inbound-filter-list,omitempty"`
	NssaExtRange       []VrfOspfAreaTypeNssaAbrNssaExtRangeXml `xml:"nssa-ext-range>entry,omitempty"`
	OutboundFilterList *string                                 `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaAbrNssaExtRangeXml struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaDefaultInformationOriginateXml struct {
	Metric     *int64  `xml:"metric,omitempty"`
	MetricType *string `xml:"metric-type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeStubXml struct {
	Abr                *VrfOspfAreaTypeStubAbrXml `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                     `xml:"default-route-metric,omitempty"`
	NoSummary          *string                    `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeStubAbrXml struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaVirtualLinkXml struct {
	Authentication *string                       `xml:"authentication,omitempty"`
	Bfd            *VrfOspfAreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	Enable         *string                       `xml:"enable,omitempty"`
	InstanceId     *int64                        `xml:"instance-id,omitempty"`
	Name           string                        `xml:"name,attr"`
	NeighborId     *string                       `xml:"neighbor-id,omitempty"`
	Timing         *string                       `xml:"timing,omitempty"`
	TransitAreaId  *string                       `xml:"transit-area-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaVirtualLinkBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfGracefulRestartXml struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3Xml struct {
	Area                  []VrfOspfv3AreaXml           `xml:"area>entry,omitempty"`
	DisableTransitTraffic *string                      `xml:"disable-transit-traffic,omitempty"`
	Enable                *string                      `xml:"enable,omitempty"`
	GlobalBfd             *VrfOspfv3GlobalBfdXml       `xml:"global-bfd,omitempty"`
	GlobalIfTimer         *string                      `xml:"global-if-timer,omitempty"`
	GracefulRestart       *VrfOspfv3GracefulRestartXml `xml:"graceful-restart,omitempty"`
	RedistributionProfile *string                      `xml:"redistribution-profile,omitempty"`
	RouterId              *string                      `xml:"router-id,omitempty"`
	SpfTimer              *string                      `xml:"spf-timer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaXml struct {
	Authentication *string                       `xml:"authentication,omitempty"`
	Interface      []VrfOspfv3AreaInterfaceXml   `xml:"interface>entry,omitempty"`
	Name           string                        `xml:"name,attr"`
	Range          []VrfOspfv3AreaRangeXml       `xml:"range>entry,omitempty"`
	Type           *VrfOspfv3AreaTypeXml         `xml:"type,omitempty"`
	VirtualLink    []VrfOspfv3AreaVirtualLinkXml `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceXml struct {
	Authentication *string                            `xml:"authentication,omitempty"`
	Bfd            *VrfOspfv3AreaInterfaceBfdXml      `xml:"bfd,omitempty"`
	Enable         *string                            `xml:"enable,omitempty"`
	InstanceId     *int64                             `xml:"instance-id,omitempty"`
	LinkType       *VrfOspfv3AreaInterfaceLinkTypeXml `xml:"link-type,omitempty"`
	Metric         *int64                             `xml:"metric,omitempty"`
	MtuIgnore      *string                            `xml:"mtu-ignore,omitempty"`
	Name           string                             `xml:"name,attr"`
	Passive        *string                            `xml:"passive,omitempty"`
	Priority       *int64                             `xml:"priority,omitempty"`
	Timing         *string                            `xml:"timing,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeXml struct {
	Broadcast *VrfOspfv3AreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2mp      *VrfOspfv3AreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	P2p       *VrfOspfv3AreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpXml struct {
	Neighbor []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml `xml:"neighbor>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml struct {
	Name     string `xml:"name,attr"`
	Priority *int64 `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaRangeXml struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeXml struct {
	Normal *VrfOspfv3AreaTypeNormalXml `xml:"normal,omitempty"`
	Nssa   *VrfOspfv3AreaTypeNssaXml   `xml:"nssa,omitempty"`
	Stub   *VrfOspfv3AreaTypeStubXml   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNormalXml struct {
	Abr *VrfOspfv3AreaTypeNormalAbrXml `xml:"abr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNormalAbrXml struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaXml struct {
	Abr                         *VrfOspfv3AreaTypeNssaAbrXml                         `xml:"abr,omitempty"`
	DefaultInformationOriginate *VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml `xml:"default-information-originate,omitempty"`
	NoSummary                   *string                                              `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaAbrXml struct {
	ExportList         *string                                   `xml:"export-list,omitempty"`
	ImportList         *string                                   `xml:"import-list,omitempty"`
	InboundFilterList  *string                                   `xml:"inbound-filter-list,omitempty"`
	NssaExtRange       []VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml `xml:"nssa-ext-range>entry,omitempty"`
	OutboundFilterList *string                                   `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml struct {
	Metric     *int64  `xml:"metric,omitempty"`
	MetricType *string `xml:"metric-type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeStubXml struct {
	Abr                *VrfOspfv3AreaTypeStubAbrXml `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                       `xml:"default-route-metric,omitempty"`
	NoSummary          *string                      `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeStubAbrXml struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaVirtualLinkXml struct {
	Authentication *string `xml:"authentication,omitempty"`
	Enable         *string `xml:"enable,omitempty"`
	InstanceId     *int64  `xml:"instance-id,omitempty"`
	Name           string  `xml:"name,attr"`
	NeighborId     *string `xml:"neighbor-id,omitempty"`
	Timing         *string `xml:"timing,omitempty"`
	TransitAreaId  *string `xml:"transit-area-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3GlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3GracefulRestartXml struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterXml struct {
	Ipv4 *VrfRibFilterIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *VrfRibFilterIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4Xml struct {
	Bgp    *VrfRibFilterIpv4BgpXml    `xml:"bgp,omitempty"`
	Ospf   *VrfRibFilterIpv4OspfXml   `xml:"ospf,omitempty"`
	Rip    *VrfRibFilterIpv4RipXml    `xml:"rip,omitempty"`
	Static *VrfRibFilterIpv4StaticXml `xml:"static,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4BgpXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4OspfXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4RipXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4StaticXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6Xml struct {
	Bgp    *VrfRibFilterIpv6BgpXml    `xml:"bgp,omitempty"`
	Ospfv3 *VrfRibFilterIpv6Ospfv3Xml `xml:"ospfv3,omitempty"`
	Static *VrfRibFilterIpv6StaticXml `xml:"static,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6BgpXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6Ospfv3Xml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6StaticXml struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipXml struct {
	AuthProfile                  *string                                `xml:"auth-profile,omitempty"`
	DefaultInformationOriginate  *string                                `xml:"default-information-originate,omitempty"`
	Enable                       *string                                `xml:"enable,omitempty"`
	GlobalBfd                    *VrfRipGlobalBfdXml                    `xml:"global-bfd,omitempty"`
	GlobalInboundDistributeList  *VrfRipGlobalInboundDistributeListXml  `xml:"global-inbound-distribute-list,omitempty"`
	GlobalOutboundDistributeList *VrfRipGlobalOutboundDistributeListXml `xml:"global-outbound-distribute-list,omitempty"`
	GlobalTimer                  *string                                `xml:"global-timer,omitempty"`
	Interface                    []VrfRipInterfaceXml                   `xml:"interface>entry,omitempty"`
	RedistributionProfile        *string                                `xml:"redistribution-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalInboundDistributeListXml struct {
	AccessList *string `xml:"access-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalOutboundDistributeListXml struct {
	AccessList *string `xml:"access-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceXml struct {
	Authentication                  *string                                            `xml:"authentication,omitempty"`
	Bfd                             *VrfRipInterfaceBfdXml                             `xml:"bfd,omitempty"`
	Enable                          *string                                            `xml:"enable,omitempty"`
	InterfaceInboundDistributeList  *VrfRipInterfaceInterfaceInboundDistributeListXml  `xml:"interface-inbound-distribute-list,omitempty"`
	InterfaceOutboundDistributeList *VrfRipInterfaceInterfaceOutboundDistributeListXml `xml:"interface-outbound-distribute-list,omitempty"`
	Mode                            *string                                            `xml:"mode,omitempty"`
	Name                            string                                             `xml:"name,attr"`
	SplitHorizon                    *string                                            `xml:"split-horizon,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceInterfaceInboundDistributeListXml struct {
	AccessList *string `xml:"access-list,omitempty"`
	Metric     *int64  `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceInterfaceOutboundDistributeListXml struct {
	AccessList *string `xml:"access-list,omitempty"`
	Metric     *int64  `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableXml struct {
	Ip   *VrfRoutingTableIpXml   `xml:"ip,omitempty"`
	Ipv6 *VrfRoutingTableIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpXml struct {
	StaticRoute []VrfRoutingTableIpStaticRouteXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteXml struct {
	AdminDist   *int64                                      `xml:"admin-dist,omitempty"`
	Bfd         *VrfRoutingTableIpStaticRouteBfdXml         `xml:"bfd,omitempty"`
	Destination *string                                     `xml:"destination,omitempty"`
	Interface   *string                                     `xml:"interface,omitempty"`
	Metric      *int64                                      `xml:"metric,omitempty"`
	Name        string                                      `xml:"name,attr"`
	Nexthop     *VrfRoutingTableIpStaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	PathMonitor *VrfRoutingTableIpStaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteNexthopXml struct {
	Discard   *VrfRoutingTableIpStaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	Fqdn      *string                                        `xml:"fqdn,omitempty"`
	IpAddress *string                                        `xml:"ip-address,omitempty"`
	NextLr    *string                                        `xml:"next-lr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRoutePathMonitorXml struct {
	Enable              *string                                                         `xml:"enable,omitempty"`
	FailureCondition    *string                                                         `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                          `xml:"hold-time,omitempty"`
	MonitorDestinations []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml struct {
	Count       *int64  `xml:"count,omitempty"`
	Destination *string `xml:"destination,omitempty"`
	Enable      *string `xml:"enable,omitempty"`
	Interval    *int64  `xml:"interval,omitempty"`
	Name        string  `xml:"name,attr"`
	Source      *string `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6Xml struct {
	StaticRoute []VrfRoutingTableIpv6StaticRouteXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteXml struct {
	AdminDist   *int64                                        `xml:"admin-dist,omitempty"`
	Bfd         *VrfRoutingTableIpv6StaticRouteBfdXml         `xml:"bfd,omitempty"`
	Destination *string                                       `xml:"destination,omitempty"`
	Interface   *string                                       `xml:"interface,omitempty"`
	Metric      *int64                                        `xml:"metric,omitempty"`
	Name        string                                        `xml:"name,attr"`
	Nexthop     *VrfRoutingTableIpv6StaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	PathMonitor *VrfRoutingTableIpv6StaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteNexthopXml struct {
	Discard     *VrfRoutingTableIpv6StaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	Fqdn        *string                                          `xml:"fqdn,omitempty"`
	Ipv6Address *string                                          `xml:"ipv6-address,omitempty"`
	NextLr      *string                                          `xml:"next-lr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRoutePathMonitorXml struct {
	Enable              *string                                                           `xml:"enable,omitempty"`
	FailureCondition    *string                                                           `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                            `xml:"hold-time,omitempty"`
	MonitorDestinations []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml struct {
	Count       *int64  `xml:"count,omitempty"`
	Destination *string `xml:"destination,omitempty"`
	Enable      *string `xml:"enable,omitempty"`
	Interval    *int64  `xml:"interval,omitempty"`
	Name        string  `xml:"name,attr"`
	Source      *string `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfXml_11_0_2 struct {
	AdminDists   *VrfAdminDistsXml_11_0_2   `xml:"admin-dists,omitempty"`
	Bgp          *VrfBgpXml_11_0_2          `xml:"bgp,omitempty"`
	Ecmp         *VrfEcmpXml_11_0_2         `xml:"ecmp,omitempty"`
	Interface    *util.MemberType           `xml:"interface,omitempty"`
	Multicast    *VrfMulticastXml_11_0_2    `xml:"multicast,omitempty"`
	Name         string                     `xml:"name,attr"`
	Ospf         *VrfOspfXml_11_0_2         `xml:"ospf,omitempty"`
	Ospfv3       *VrfOspfv3Xml_11_0_2       `xml:"ospfv3,omitempty"`
	RibFilter    *VrfRibFilterXml_11_0_2    `xml:"rib-filter,omitempty"`
	Rip          *VrfRipXml_11_0_2          `xml:"rip,omitempty"`
	RoutingTable *VrfRoutingTableXml_11_0_2 `xml:"routing-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfAdminDistsXml_11_0_2 struct {
	BgpExternal *int64 `xml:"bgp-external,omitempty"`
	BgpInternal *int64 `xml:"bgp-internal,omitempty"`
	BgpLocal    *int64 `xml:"bgp-local,omitempty"`
	OspfExt     *int64 `xml:"ospf-ext,omitempty"`
	OspfInter   *int64 `xml:"ospf-inter,omitempty"`
	OspfIntra   *int64 `xml:"ospf-intra,omitempty"`
	Ospfv3Ext   *int64 `xml:"ospfv3-ext,omitempty"`
	Ospfv3Inter *int64 `xml:"ospfv3-inter,omitempty"`
	Ospfv3Intra *int64 `xml:"ospfv3-intra,omitempty"`
	Rip         *int64 `xml:"rip,omitempty"`
	Static      *int64 `xml:"static,omitempty"`
	StaticIpv6  *int64 `xml:"static-ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpXml_11_0_2 struct {
	AdvertiseNetwork            *VrfBgpAdvertiseNetworkXml_11_0_2      `xml:"advertise-network,omitempty"`
	AggregateRoutes             []VrfBgpAggregateRoutesXml_11_0_2      `xml:"aggregate-routes>entry,omitempty"`
	AlwaysAdvertiseNetworkRoute *string                                `xml:"always-advertise-network-route,omitempty"`
	DefaultLocalPreference      *int64                                 `xml:"default-local-preference,omitempty"`
	EcmpMultiAs                 *string                                `xml:"ecmp-multi-as,omitempty"`
	Enable                      *string                                `xml:"enable,omitempty"`
	EnforceFirstAs              *string                                `xml:"enforce-first-as,omitempty"`
	FastExternalFailover        *string                                `xml:"fast-external-failover,omitempty"`
	GlobalBfd                   *VrfBgpGlobalBfdXml_11_0_2             `xml:"global-bfd,omitempty"`
	GracefulRestart             *VrfBgpGracefulRestartXml_11_0_2       `xml:"graceful-restart,omitempty"`
	GracefulShutdown            *string                                `xml:"graceful-shutdown,omitempty"`
	InstallRoute                *string                                `xml:"install-route,omitempty"`
	LocalAs                     *string                                `xml:"local-as,omitempty"`
	Med                         *VrfBgpMedXml_11_0_2                   `xml:"med,omitempty"`
	PeerGroup                   []VrfBgpPeerGroupXml_11_0_2            `xml:"peer-group>entry,omitempty"`
	RedistributionProfile       *VrfBgpRedistributionProfileXml_11_0_2 `xml:"redistribution-profile,omitempty"`
	RouterId                    *string                                `xml:"router-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkXml_11_0_2 struct {
	Ipv4 *VrfBgpAdvertiseNetworkIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpAdvertiseNetworkIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv4Xml_11_0_2 struct {
	Network []VrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2 `xml:"network>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2 struct {
	Backdoor  *string `xml:"backdoor,omitempty"`
	Multicast *string `xml:"multicast,omitempty"`
	Name      string  `xml:"name,attr"`
	Unicast   *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv6Xml_11_0_2 struct {
	Network []VrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2 `xml:"network>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2 struct {
	Name    string  `xml:"name,attr"`
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesXml_11_0_2 struct {
	AsSet       *string                              `xml:"as-set,omitempty"`
	Description *string                              `xml:"description,omitempty"`
	Enable      *string                              `xml:"enable,omitempty"`
	Name        string                               `xml:"name,attr"`
	SameMed     *string                              `xml:"same-med,omitempty"`
	SummaryOnly *string                              `xml:"summary-only,omitempty"`
	Type        *VrfBgpAggregateRoutesTypeXml_11_0_2 `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeXml_11_0_2 struct {
	Ipv4 *VrfBgpAggregateRoutesTypeIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpAggregateRoutesTypeIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeIpv4Xml_11_0_2 struct {
	AttributeMap  *string `xml:"attribute-map,omitempty"`
	SummaryPrefix *string `xml:"summary-prefix,omitempty"`
	SuppressMap   *string `xml:"suppress-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpAggregateRoutesTypeIpv6Xml_11_0_2 struct {
	AttributeMap  *string `xml:"attribute-map,omitempty"`
	SummaryPrefix *string `xml:"summary-prefix,omitempty"`
	SuppressMap   *string `xml:"suppress-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpGlobalBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpGracefulRestartXml_11_0_2 struct {
	Enable             *string `xml:"enable,omitempty"`
	LocalRestartTime   *int64  `xml:"local-restart-time,omitempty"`
	MaxPeerRestartTime *int64  `xml:"max-peer-restart-time,omitempty"`
	StaleRouteTime     *int64  `xml:"stale-route-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpMedXml_11_0_2 struct {
	AlwaysCompareMed           *string `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string `xml:"deterministic-med-comparison,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupXml_11_0_2 struct {
	AddressFamily     *VrfBgpPeerGroupAddressFamilyXml_11_0_2     `xml:"address-family,omitempty"`
	ConnectionOptions *VrfBgpPeerGroupConnectionOptionsXml_11_0_2 `xml:"connection-options,omitempty"`
	Enable            *string                                     `xml:"enable,omitempty"`
	FilteringProfile  *VrfBgpPeerGroupFilteringProfileXml_11_0_2  `xml:"filtering-profile,omitempty"`
	Name              string                                      `xml:"name,attr"`
	Peer              []VrfBgpPeerGroupPeerXml_11_0_2             `xml:"peer>entry,omitempty"`
	Type              *VrfBgpPeerGroupTypeXml_11_0_2              `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupAddressFamilyXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupConnectionOptionsXml_11_0_2 struct {
	Authentication *string `xml:"authentication,omitempty"`
	Dampening      *string `xml:"dampening,omitempty"`
	Multihop       *int64  `xml:"multihop,omitempty"`
	Timers         *string `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupFilteringProfileXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerXml_11_0_2 struct {
	Bfd                           *VrfBgpPeerGroupPeerBfdXml_11_0_2               `xml:"bfd,omitempty"`
	ConnectionOptions             *VrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2 `xml:"connection-options,omitempty"`
	Enable                        *string                                         `xml:"enable,omitempty"`
	EnableSenderSideLoopDetection *string                                         `xml:"enable-sender-side-loop-detection,omitempty"`
	Inherit                       *VrfBgpPeerGroupPeerInheritXml_11_0_2           `xml:"inherit,omitempty"`
	LocalAddress                  *VrfBgpPeerGroupPeerLocalAddressXml_11_0_2      `xml:"local-address,omitempty"`
	Name                          string                                          `xml:"name,attr"`
	Passive                       *string                                         `xml:"passive,omitempty"`
	PeerAddress                   *VrfBgpPeerGroupPeerPeerAddressXml_11_0_2       `xml:"peer-address,omitempty"`
	PeerAs                        *string                                         `xml:"peer-as,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2 struct {
	Authentication *string `xml:"authentication,omitempty"`
	Dampening      *string `xml:"dampening,omitempty"`
	Multihop       *string `xml:"multihop,omitempty"`
	Timers         *string `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritXml_11_0_2 struct {
	No  *VrfBgpPeerGroupPeerInheritNoXml_11_0_2  `xml:"no,omitempty"`
	Yes *VrfBgpPeerGroupPeerInheritYesXml_11_0_2 `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoXml_11_0_2 struct {
	AddressFamily    *VrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2    `xml:"address-family,omitempty"`
	FilteringProfile *VrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2 `xml:"filtering-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerInheritYesXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerLocalAddressXml_11_0_2 struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupPeerPeerAddressXml_11_0_2 struct {
	Fqdn *string `xml:"fqdn,omitempty"`
	Ip   *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeXml_11_0_2 struct {
	Ebgp *VrfBgpPeerGroupTypeEbgpXml_11_0_2 `xml:"ebgp,omitempty"`
	Ibgp *VrfBgpPeerGroupTypeIbgpXml_11_0_2 `xml:"ibgp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeEbgpXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpPeerGroupTypeIbgpXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileXml_11_0_2 struct {
	Ipv4 *VrfBgpRedistributionProfileIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *VrfBgpRedistributionProfileIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileIpv4Xml_11_0_2 struct {
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfBgpRedistributionProfileIpv6Xml_11_0_2 struct {
	Unicast *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpXml_11_0_2 struct {
	Algorithm        *VrfEcmpAlgorithmXml_11_0_2 `xml:"algorithm,omitempty"`
	Enable           *string                     `xml:"enable,omitempty"`
	MaxPath          *int64                      `xml:"max-path,omitempty"`
	StrictSourcePath *string                     `xml:"strict-source-path,omitempty"`
	SymmetricReturn  *string                     `xml:"symmetric-return,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmXml_11_0_2 struct {
	BalancedRoundRobin *VrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2 `xml:"balanced-round-robin,omitempty"`
	IpHash             *VrfEcmpAlgorithmIpHashXml_11_0_2             `xml:"ip-hash,omitempty"`
	IpModulo           *VrfEcmpAlgorithmIpModuloXml_11_0_2           `xml:"ip-modulo,omitempty"`
	WeightedRoundRobin *VrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2 `xml:"weighted-round-robin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmIpHashXml_11_0_2 struct {
	HashSeed *int64  `xml:"hash-seed,omitempty"`
	SrcOnly  *string `xml:"src-only,omitempty"`
	UsePort  *string `xml:"use-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmIpModuloXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2 struct {
	Interface []VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2 `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2 struct {
	Name   string `xml:"name,attr"`
	Weight *int64 `xml:"weight,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastXml_11_0_2 struct {
	Enable      *string                             `xml:"enable,omitempty"`
	Igmp        *VrfMulticastIgmpXml_11_0_2         `xml:"igmp,omitempty"`
	Msdp        *VrfMulticastMsdpXml_11_0_2         `xml:"msdp,omitempty"`
	Pim         *VrfMulticastPimXml_11_0_2          `xml:"pim,omitempty"`
	StaticRoute []VrfMulticastStaticRouteXml_11_0_2 `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpXml_11_0_2 struct {
	Dynamic *VrfMulticastIgmpDynamicXml_11_0_2 `xml:"dynamic,omitempty"`
	Enable  *string                            `xml:"enable,omitempty"`
	Static  []VrfMulticastIgmpStaticXml_11_0_2 `xml:"static>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpDynamicXml_11_0_2 struct {
	Interface []VrfMulticastIgmpDynamicInterfaceXml_11_0_2 `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpDynamicInterfaceXml_11_0_2 struct {
	GroupFilter         *string `xml:"group-filter,omitempty"`
	MaxGroups           *string `xml:"max-groups,omitempty"`
	MaxSources          *string `xml:"max-sources,omitempty"`
	Name                string  `xml:"name,attr"`
	QueryProfile        *string `xml:"query-profile,omitempty"`
	Robustness          *string `xml:"robustness,omitempty"`
	RouterAlertPolicing *string `xml:"router-alert-policing,omitempty"`
	Version             *string `xml:"version,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastIgmpStaticXml_11_0_2 struct {
	GroupAddress  *string `xml:"group-address,omitempty"`
	Interface     *string `xml:"interface,omitempty"`
	Name          string  `xml:"name,attr"`
	SourceAddress *string `xml:"source-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpXml_11_0_2 struct {
	Enable               *string                                 `xml:"enable,omitempty"`
	GlobalAuthentication *string                                 `xml:"global-authentication,omitempty"`
	GlobalTimer          *string                                 `xml:"global-timer,omitempty"`
	OriginatorId         *VrfMulticastMsdpOriginatorIdXml_11_0_2 `xml:"originator-id,omitempty"`
	Peer                 []VrfMulticastMsdpPeerXml_11_0_2        `xml:"peer>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpOriginatorIdXml_11_0_2 struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerXml_11_0_2 struct {
	Authentication   *string                                     `xml:"authentication,omitempty"`
	Enable           *string                                     `xml:"enable,omitempty"`
	InboundSaFilter  *string                                     `xml:"inbound-sa-filter,omitempty"`
	LocalAddress     *VrfMulticastMsdpPeerLocalAddressXml_11_0_2 `xml:"local-address,omitempty"`
	MaxSa            *int64                                      `xml:"max-sa,omitempty"`
	Name             string                                      `xml:"name,attr"`
	OutboundSaFilter *string                                     `xml:"outbound-sa-filter,omitempty"`
	PeerAddress      *VrfMulticastMsdpPeerPeerAddressXml_11_0_2  `xml:"peer-address,omitempty"`
	PeerAs           *string                                     `xml:"peer-as,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerLocalAddressXml_11_0_2 struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastMsdpPeerPeerAddressXml_11_0_2 struct {
	Fqdn *string `xml:"fqdn,omitempty"`
	Ip   *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimXml_11_0_2 struct {
	Enable          *string                                   `xml:"enable,omitempty"`
	GroupPermission *string                                   `xml:"group-permission,omitempty"`
	IfTimerGlobal   *string                                   `xml:"if-timer-global,omitempty"`
	Interface       []VrfMulticastPimInterfaceXml_11_0_2      `xml:"interface>entry,omitempty"`
	RouteAgeoutTime *int64                                    `xml:"route-ageout-time,omitempty"`
	Rp              *VrfMulticastPimRpXml_11_0_2              `xml:"rp,omitempty"`
	RpfLookupMode   *string                                   `xml:"rpf-lookup-mode,omitempty"`
	SptThreshold    []VrfMulticastPimSptThresholdXml_11_0_2   `xml:"spt-threshold>entry,omitempty"`
	SsmAddressSpace *VrfMulticastPimSsmAddressSpaceXml_11_0_2 `xml:"ssm-address-space,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimInterfaceXml_11_0_2 struct {
	Description    *string `xml:"description,omitempty"`
	DrPriority     *int64  `xml:"dr-priority,omitempty"`
	IfTimer        *string `xml:"if-timer,omitempty"`
	Name           string  `xml:"name,attr"`
	NeighborFilter *string `xml:"neighbor-filter,omitempty"`
	SendBsm        *string `xml:"send-bsm,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpXml_11_0_2 struct {
	ExternalRp []VrfMulticastPimRpExternalRpXml_11_0_2 `xml:"external-rp>entry,omitempty"`
	LocalRp    *VrfMulticastPimRpLocalRpXml_11_0_2     `xml:"local-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpExternalRpXml_11_0_2 struct {
	GroupList *string `xml:"group-list,omitempty"`
	Name      string  `xml:"name,attr"`
	Override  *string `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpXml_11_0_2 struct {
	CandidateRp *VrfMulticastPimRpLocalRpCandidateRpXml_11_0_2 `xml:"candidate-rp,omitempty"`
	StaticRp    *VrfMulticastPimRpLocalRpStaticRpXml_11_0_2    `xml:"static-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpCandidateRpXml_11_0_2 struct {
	Address               *string `xml:"address,omitempty"`
	AdvertisementInterval *int64  `xml:"advertisement-interval,omitempty"`
	GroupList             *string `xml:"group-list,omitempty"`
	Interface             *string `xml:"interface,omitempty"`
	Priority              *int64  `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimRpLocalRpStaticRpXml_11_0_2 struct {
	Address   *string `xml:"address,omitempty"`
	GroupList *string `xml:"group-list,omitempty"`
	Interface *string `xml:"interface,omitempty"`
	Override  *string `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimSptThresholdXml_11_0_2 struct {
	Name      string  `xml:"name,attr"`
	Threshold *string `xml:"threshold,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastPimSsmAddressSpaceXml_11_0_2 struct {
	GroupList *string `xml:"group-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastStaticRouteXml_11_0_2 struct {
	Destination *string                                   `xml:"destination,omitempty"`
	Interface   *string                                   `xml:"interface,omitempty"`
	Name        string                                    `xml:"name,attr"`
	Nexthop     *VrfMulticastStaticRouteNexthopXml_11_0_2 `xml:"nexthop,omitempty"`
	Preference  *int64                                    `xml:"preference,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfMulticastStaticRouteNexthopXml_11_0_2 struct {
	IpAddress *string `xml:"ip-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfXml_11_0_2 struct {
	Area                  []VrfOspfAreaXml_11_0_2           `xml:"area>entry,omitempty"`
	Enable                *string                           `xml:"enable,omitempty"`
	GlobalBfd             *VrfOspfGlobalBfdXml_11_0_2       `xml:"global-bfd,omitempty"`
	GlobalIfTimer         *string                           `xml:"global-if-timer,omitempty"`
	GracefulRestart       *VrfOspfGracefulRestartXml_11_0_2 `xml:"graceful-restart,omitempty"`
	RedistributionProfile *string                           `xml:"redistribution-profile,omitempty"`
	Rfc1583               *string                           `xml:"rfc1583,omitempty"`
	RouterId              *string                           `xml:"router-id,omitempty"`
	SpfTimer              *string                           `xml:"spf-timer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaXml_11_0_2 struct {
	Authentication *string                            `xml:"authentication,omitempty"`
	Interface      []VrfOspfAreaInterfaceXml_11_0_2   `xml:"interface>entry,omitempty"`
	Name           string                             `xml:"name,attr"`
	Range          []VrfOspfAreaRangeXml_11_0_2       `xml:"range>entry,omitempty"`
	Type           *VrfOspfAreaTypeXml_11_0_2         `xml:"type,omitempty"`
	VirtualLink    []VrfOspfAreaVirtualLinkXml_11_0_2 `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceXml_11_0_2 struct {
	Authentication *string                                 `xml:"authentication,omitempty"`
	Bfd            *VrfOspfAreaInterfaceBfdXml_11_0_2      `xml:"bfd,omitempty"`
	Enable         *string                                 `xml:"enable,omitempty"`
	LinkType       *VrfOspfAreaInterfaceLinkTypeXml_11_0_2 `xml:"link-type,omitempty"`
	Metric         *int64                                  `xml:"metric,omitempty"`
	MtuIgnore      *string                                 `xml:"mtu-ignore,omitempty"`
	Name           string                                  `xml:"name,attr"`
	Passive        *string                                 `xml:"passive,omitempty"`
	Priority       *int64                                  `xml:"priority,omitempty"`
	Timing         *string                                 `xml:"timing,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeXml_11_0_2 struct {
	Broadcast *VrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2 `xml:"broadcast,omitempty"`
	P2mp      *VrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2      `xml:"p2mp,omitempty"`
	P2p       *VrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2 struct {
	Neighbor []VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 `xml:"neighbor>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 struct {
	Name     string `xml:"name,attr"`
	Priority *int64 `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaRangeXml_11_0_2 struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeXml_11_0_2 struct {
	Normal *VrfOspfAreaTypeNormalXml_11_0_2 `xml:"normal,omitempty"`
	Nssa   *VrfOspfAreaTypeNssaXml_11_0_2   `xml:"nssa,omitempty"`
	Stub   *VrfOspfAreaTypeStubXml_11_0_2   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNormalXml_11_0_2 struct {
	Abr *VrfOspfAreaTypeNormalAbrXml_11_0_2 `xml:"abr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNormalAbrXml_11_0_2 struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaXml_11_0_2 struct {
	Abr                         *VrfOspfAreaTypeNssaAbrXml_11_0_2                         `xml:"abr,omitempty"`
	DefaultInformationOriginate *VrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2 `xml:"default-information-originate,omitempty"`
	NoSummary                   *string                                                   `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaAbrXml_11_0_2 struct {
	ExportList         *string                                        `xml:"export-list,omitempty"`
	ImportList         *string                                        `xml:"import-list,omitempty"`
	InboundFilterList  *string                                        `xml:"inbound-filter-list,omitempty"`
	NssaExtRange       []VrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2 `xml:"nssa-ext-range>entry,omitempty"`
	OutboundFilterList *string                                        `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2 struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2 struct {
	Metric     *int64  `xml:"metric,omitempty"`
	MetricType *string `xml:"metric-type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeStubXml_11_0_2 struct {
	Abr                *VrfOspfAreaTypeStubAbrXml_11_0_2 `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                            `xml:"default-route-metric,omitempty"`
	NoSummary          *string                           `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaTypeStubAbrXml_11_0_2 struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaVirtualLinkXml_11_0_2 struct {
	Authentication *string                              `xml:"authentication,omitempty"`
	Bfd            *VrfOspfAreaVirtualLinkBfdXml_11_0_2 `xml:"bfd,omitempty"`
	Enable         *string                              `xml:"enable,omitempty"`
	InstanceId     *int64                               `xml:"instance-id,omitempty"`
	Name           string                               `xml:"name,attr"`
	NeighborId     *string                              `xml:"neighbor-id,omitempty"`
	Timing         *string                              `xml:"timing,omitempty"`
	TransitAreaId  *string                              `xml:"transit-area-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfAreaVirtualLinkBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfGlobalBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfGracefulRestartXml_11_0_2 struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3Xml_11_0_2 struct {
	Area                  []VrfOspfv3AreaXml_11_0_2           `xml:"area>entry,omitempty"`
	DisableTransitTraffic *string                             `xml:"disable-transit-traffic,omitempty"`
	Enable                *string                             `xml:"enable,omitempty"`
	GlobalBfd             *VrfOspfv3GlobalBfdXml_11_0_2       `xml:"global-bfd,omitempty"`
	GlobalIfTimer         *string                             `xml:"global-if-timer,omitempty"`
	GracefulRestart       *VrfOspfv3GracefulRestartXml_11_0_2 `xml:"graceful-restart,omitempty"`
	RedistributionProfile *string                             `xml:"redistribution-profile,omitempty"`
	RouterId              *string                             `xml:"router-id,omitempty"`
	SpfTimer              *string                             `xml:"spf-timer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaXml_11_0_2 struct {
	Authentication *string                              `xml:"authentication,omitempty"`
	Interface      []VrfOspfv3AreaInterfaceXml_11_0_2   `xml:"interface>entry,omitempty"`
	Name           string                               `xml:"name,attr"`
	Range          []VrfOspfv3AreaRangeXml_11_0_2       `xml:"range>entry,omitempty"`
	Type           *VrfOspfv3AreaTypeXml_11_0_2         `xml:"type,omitempty"`
	VirtualLink    []VrfOspfv3AreaVirtualLinkXml_11_0_2 `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceXml_11_0_2 struct {
	Authentication *string                                   `xml:"authentication,omitempty"`
	Bfd            *VrfOspfv3AreaInterfaceBfdXml_11_0_2      `xml:"bfd,omitempty"`
	Enable         *string                                   `xml:"enable,omitempty"`
	InstanceId     *int64                                    `xml:"instance-id,omitempty"`
	LinkType       *VrfOspfv3AreaInterfaceLinkTypeXml_11_0_2 `xml:"link-type,omitempty"`
	Metric         *int64                                    `xml:"metric,omitempty"`
	MtuIgnore      *string                                   `xml:"mtu-ignore,omitempty"`
	Name           string                                    `xml:"name,attr"`
	Passive        *string                                   `xml:"passive,omitempty"`
	Priority       *int64                                    `xml:"priority,omitempty"`
	Timing         *string                                   `xml:"timing,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeXml_11_0_2 struct {
	Broadcast *VrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2 `xml:"broadcast,omitempty"`
	P2mp      *VrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2      `xml:"p2mp,omitempty"`
	P2p       *VrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2 struct {
	Neighbor []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 `xml:"neighbor>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 struct {
	Name     string `xml:"name,attr"`
	Priority *int64 `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaRangeXml_11_0_2 struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeXml_11_0_2 struct {
	Normal *VrfOspfv3AreaTypeNormalXml_11_0_2 `xml:"normal,omitempty"`
	Nssa   *VrfOspfv3AreaTypeNssaXml_11_0_2   `xml:"nssa,omitempty"`
	Stub   *VrfOspfv3AreaTypeStubXml_11_0_2   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNormalXml_11_0_2 struct {
	Abr *VrfOspfv3AreaTypeNormalAbrXml_11_0_2 `xml:"abr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNormalAbrXml_11_0_2 struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaXml_11_0_2 struct {
	Abr                         *VrfOspfv3AreaTypeNssaAbrXml_11_0_2                         `xml:"abr,omitempty"`
	DefaultInformationOriginate *VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2 `xml:"default-information-originate,omitempty"`
	NoSummary                   *string                                                     `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaAbrXml_11_0_2 struct {
	ExportList         *string                                          `xml:"export-list,omitempty"`
	ImportList         *string                                          `xml:"import-list,omitempty"`
	InboundFilterList  *string                                          `xml:"inbound-filter-list,omitempty"`
	NssaExtRange       []VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2 `xml:"nssa-ext-range>entry,omitempty"`
	OutboundFilterList *string                                          `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2 struct {
	Advertise *string `xml:"advertise,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2 struct {
	Metric     *int64  `xml:"metric,omitempty"`
	MetricType *string `xml:"metric-type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeStubXml_11_0_2 struct {
	Abr                *VrfOspfv3AreaTypeStubAbrXml_11_0_2 `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                              `xml:"default-route-metric,omitempty"`
	NoSummary          *string                             `xml:"no-summary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaTypeStubAbrXml_11_0_2 struct {
	ExportList         *string `xml:"export-list,omitempty"`
	ImportList         *string `xml:"import-list,omitempty"`
	InboundFilterList  *string `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string `xml:"outbound-filter-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3AreaVirtualLinkXml_11_0_2 struct {
	Authentication *string `xml:"authentication,omitempty"`
	Enable         *string `xml:"enable,omitempty"`
	InstanceId     *int64  `xml:"instance-id,omitempty"`
	Name           string  `xml:"name,attr"`
	NeighborId     *string `xml:"neighbor-id,omitempty"`
	Timing         *string `xml:"timing,omitempty"`
	TransitAreaId  *string `xml:"transit-area-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3GlobalBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfOspfv3GracefulRestartXml_11_0_2 struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterXml_11_0_2 struct {
	Ipv4 *VrfRibFilterIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *VrfRibFilterIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4Xml_11_0_2 struct {
	Bgp    *VrfRibFilterIpv4BgpXml_11_0_2    `xml:"bgp,omitempty"`
	Ospf   *VrfRibFilterIpv4OspfXml_11_0_2   `xml:"ospf,omitempty"`
	Rip    *VrfRibFilterIpv4RipXml_11_0_2    `xml:"rip,omitempty"`
	Static *VrfRibFilterIpv4StaticXml_11_0_2 `xml:"static,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4BgpXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4OspfXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4RipXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv4StaticXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6Xml_11_0_2 struct {
	Bgp    *VrfRibFilterIpv6BgpXml_11_0_2    `xml:"bgp,omitempty"`
	Ospfv3 *VrfRibFilterIpv6Ospfv3Xml_11_0_2 `xml:"ospfv3,omitempty"`
	Static *VrfRibFilterIpv6StaticXml_11_0_2 `xml:"static,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6BgpXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6Ospfv3Xml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRibFilterIpv6StaticXml_11_0_2 struct {
	RouteMap *string `xml:"route-map,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipXml_11_0_2 struct {
	AuthProfile                  *string                                       `xml:"auth-profile,omitempty"`
	DefaultInformationOriginate  *string                                       `xml:"default-information-originate,omitempty"`
	Enable                       *string                                       `xml:"enable,omitempty"`
	GlobalBfd                    *VrfRipGlobalBfdXml_11_0_2                    `xml:"global-bfd,omitempty"`
	GlobalInboundDistributeList  *VrfRipGlobalInboundDistributeListXml_11_0_2  `xml:"global-inbound-distribute-list,omitempty"`
	GlobalOutboundDistributeList *VrfRipGlobalOutboundDistributeListXml_11_0_2 `xml:"global-outbound-distribute-list,omitempty"`
	GlobalTimer                  *string                                       `xml:"global-timer,omitempty"`
	Interface                    []VrfRipInterfaceXml_11_0_2                   `xml:"interface>entry,omitempty"`
	RedistributionProfile        *string                                       `xml:"redistribution-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalInboundDistributeListXml_11_0_2 struct {
	AccessList *string `xml:"access-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipGlobalOutboundDistributeListXml_11_0_2 struct {
	AccessList *string `xml:"access-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceXml_11_0_2 struct {
	Authentication                  *string                                                   `xml:"authentication,omitempty"`
	Bfd                             *VrfRipInterfaceBfdXml_11_0_2                             `xml:"bfd,omitempty"`
	Enable                          *string                                                   `xml:"enable,omitempty"`
	InterfaceInboundDistributeList  *VrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2  `xml:"interface-inbound-distribute-list,omitempty"`
	InterfaceOutboundDistributeList *VrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2 `xml:"interface-outbound-distribute-list,omitempty"`
	Mode                            *string                                                   `xml:"mode,omitempty"`
	Name                            string                                                    `xml:"name,attr"`
	SplitHorizon                    *string                                                   `xml:"split-horizon,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2 struct {
	AccessList *string `xml:"access-list,omitempty"`
	Metric     *int64  `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2 struct {
	AccessList *string `xml:"access-list,omitempty"`
	Metric     *int64  `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableXml_11_0_2 struct {
	Ip   *VrfRoutingTableIpXml_11_0_2   `xml:"ip,omitempty"`
	Ipv6 *VrfRoutingTableIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpXml_11_0_2 struct {
	StaticRoute []VrfRoutingTableIpStaticRouteXml_11_0_2 `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteXml_11_0_2 struct {
	AdminDist   *int64                                             `xml:"admin-dist,omitempty"`
	Bfd         *VrfRoutingTableIpStaticRouteBfdXml_11_0_2         `xml:"bfd,omitempty"`
	Destination *string                                            `xml:"destination,omitempty"`
	Interface   *string                                            `xml:"interface,omitempty"`
	Metric      *int64                                             `xml:"metric,omitempty"`
	Name        string                                             `xml:"name,attr"`
	Nexthop     *VrfRoutingTableIpStaticRouteNexthopXml_11_0_2     `xml:"nexthop,omitempty"`
	PathMonitor *VrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2 `xml:"path-monitor,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteNexthopXml_11_0_2 struct {
	Discard   *VrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2 `xml:"discard,omitempty"`
	Fqdn      *string                                               `xml:"fqdn,omitempty"`
	IpAddress *string                                               `xml:"ip-address,omitempty"`
	NextLr    *string                                               `xml:"next-lr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2 struct {
	Enable              *string                                                                `xml:"enable,omitempty"`
	FailureCondition    *string                                                                `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                 `xml:"hold-time,omitempty"`
	MonitorDestinations []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2 `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2 struct {
	Count       *int64  `xml:"count,omitempty"`
	Destination *string `xml:"destination,omitempty"`
	Enable      *string `xml:"enable,omitempty"`
	Interval    *int64  `xml:"interval,omitempty"`
	Name        string  `xml:"name,attr"`
	Source      *string `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6Xml_11_0_2 struct {
	StaticRoute []VrfRoutingTableIpv6StaticRouteXml_11_0_2 `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteXml_11_0_2 struct {
	AdminDist   *int64                                               `xml:"admin-dist,omitempty"`
	Bfd         *VrfRoutingTableIpv6StaticRouteBfdXml_11_0_2         `xml:"bfd,omitempty"`
	Destination *string                                              `xml:"destination,omitempty"`
	Interface   *string                                              `xml:"interface,omitempty"`
	Metric      *int64                                               `xml:"metric,omitempty"`
	Name        string                                               `xml:"name,attr"`
	Nexthop     *VrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2     `xml:"nexthop,omitempty"`
	PathMonitor *VrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2 `xml:"path-monitor,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteBfdXml_11_0_2 struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2 struct {
	Discard     *VrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2 `xml:"discard,omitempty"`
	Fqdn        *string                                                 `xml:"fqdn,omitempty"`
	Ipv6Address *string                                                 `xml:"ipv6-address,omitempty"`
	NextLr      *string                                                 `xml:"next-lr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2 struct {
	Enable              *string                                                                  `xml:"enable,omitempty"`
	FailureCondition    *string                                                                  `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                   `xml:"hold-time,omitempty"`
	MonitorDestinations []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2 `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2 struct {
	Count       *int64  `xml:"count,omitempty"`
	Destination *string `xml:"destination,omitempty"`
	Enable      *string `xml:"enable,omitempty"`
	Interval    *int64  `xml:"interval,omitempty"`
	Name        string  `xml:"name,attr"`
	Source      *string `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "vrf" || v == "Vrf" {
		return e.Vrf, nil
	}
	if v == "vrf|LENGTH" || v == "Vrf|LENGTH" {
		return int64(len(e.Vrf)), nil
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedVrfCol []VrfXml
	if o.Vrf != nil {
		nestedVrfCol = []VrfXml{}
		for _, oVrf := range o.Vrf {
			nestedVrf := VrfXml{}
			if _, ok := o.Misc["Vrf"]; ok {
				nestedVrf.Misc = o.Misc["Vrf"]
			}
			if oVrf.Name != "" {
				nestedVrf.Name = oVrf.Name
			}
			if oVrf.Interface != nil {
				nestedVrf.Interface = util.StrToMem(oVrf.Interface)
			}
			if oVrf.AdminDists != nil {
				nestedVrf.AdminDists = &VrfAdminDistsXml{}
				if _, ok := o.Misc["VrfAdminDists"]; ok {
					nestedVrf.AdminDists.Misc = o.Misc["VrfAdminDists"]
				}
				if oVrf.AdminDists.Static != nil {
					nestedVrf.AdminDists.Static = oVrf.AdminDists.Static
				}
				if oVrf.AdminDists.StaticIpv6 != nil {
					nestedVrf.AdminDists.StaticIpv6 = oVrf.AdminDists.StaticIpv6
				}
				if oVrf.AdminDists.OspfInter != nil {
					nestedVrf.AdminDists.OspfInter = oVrf.AdminDists.OspfInter
				}
				if oVrf.AdminDists.OspfIntra != nil {
					nestedVrf.AdminDists.OspfIntra = oVrf.AdminDists.OspfIntra
				}
				if oVrf.AdminDists.OspfExt != nil {
					nestedVrf.AdminDists.OspfExt = oVrf.AdminDists.OspfExt
				}
				if oVrf.AdminDists.Ospfv3Inter != nil {
					nestedVrf.AdminDists.Ospfv3Inter = oVrf.AdminDists.Ospfv3Inter
				}
				if oVrf.AdminDists.Ospfv3Intra != nil {
					nestedVrf.AdminDists.Ospfv3Intra = oVrf.AdminDists.Ospfv3Intra
				}
				if oVrf.AdminDists.Ospfv3Ext != nil {
					nestedVrf.AdminDists.Ospfv3Ext = oVrf.AdminDists.Ospfv3Ext
				}
				if oVrf.AdminDists.BgpInternal != nil {
					nestedVrf.AdminDists.BgpInternal = oVrf.AdminDists.BgpInternal
				}
				if oVrf.AdminDists.BgpExternal != nil {
					nestedVrf.AdminDists.BgpExternal = oVrf.AdminDists.BgpExternal
				}
				if oVrf.AdminDists.BgpLocal != nil {
					nestedVrf.AdminDists.BgpLocal = oVrf.AdminDists.BgpLocal
				}
				if oVrf.AdminDists.Rip != nil {
					nestedVrf.AdminDists.Rip = oVrf.AdminDists.Rip
				}
			}
			if oVrf.RibFilter != nil {
				nestedVrf.RibFilter = &VrfRibFilterXml{}
				if _, ok := o.Misc["VrfRibFilter"]; ok {
					nestedVrf.RibFilter.Misc = o.Misc["VrfRibFilter"]
				}
				if oVrf.RibFilter.Ipv4 != nil {
					nestedVrf.RibFilter.Ipv4 = &VrfRibFilterIpv4Xml{}
					if _, ok := o.Misc["VrfRibFilterIpv4"]; ok {
						nestedVrf.RibFilter.Ipv4.Misc = o.Misc["VrfRibFilterIpv4"]
					}
					if oVrf.RibFilter.Ipv4.Static != nil {
						nestedVrf.RibFilter.Ipv4.Static = &VrfRibFilterIpv4StaticXml{}
						if _, ok := o.Misc["VrfRibFilterIpv4Static"]; ok {
							nestedVrf.RibFilter.Ipv4.Static.Misc = o.Misc["VrfRibFilterIpv4Static"]
						}
						if oVrf.RibFilter.Ipv4.Static.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Static.RouteMap = oVrf.RibFilter.Ipv4.Static.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Bgp != nil {
						nestedVrf.RibFilter.Ipv4.Bgp = &VrfRibFilterIpv4BgpXml{}
						if _, ok := o.Misc["VrfRibFilterIpv4Bgp"]; ok {
							nestedVrf.RibFilter.Ipv4.Bgp.Misc = o.Misc["VrfRibFilterIpv4Bgp"]
						}
						if oVrf.RibFilter.Ipv4.Bgp.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Bgp.RouteMap = oVrf.RibFilter.Ipv4.Bgp.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Ospf != nil {
						nestedVrf.RibFilter.Ipv4.Ospf = &VrfRibFilterIpv4OspfXml{}
						if _, ok := o.Misc["VrfRibFilterIpv4Ospf"]; ok {
							nestedVrf.RibFilter.Ipv4.Ospf.Misc = o.Misc["VrfRibFilterIpv4Ospf"]
						}
						if oVrf.RibFilter.Ipv4.Ospf.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Ospf.RouteMap = oVrf.RibFilter.Ipv4.Ospf.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Rip != nil {
						nestedVrf.RibFilter.Ipv4.Rip = &VrfRibFilterIpv4RipXml{}
						if _, ok := o.Misc["VrfRibFilterIpv4Rip"]; ok {
							nestedVrf.RibFilter.Ipv4.Rip.Misc = o.Misc["VrfRibFilterIpv4Rip"]
						}
						if oVrf.RibFilter.Ipv4.Rip.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Rip.RouteMap = oVrf.RibFilter.Ipv4.Rip.RouteMap
						}
					}
				}
				if oVrf.RibFilter.Ipv6 != nil {
					nestedVrf.RibFilter.Ipv6 = &VrfRibFilterIpv6Xml{}
					if _, ok := o.Misc["VrfRibFilterIpv6"]; ok {
						nestedVrf.RibFilter.Ipv6.Misc = o.Misc["VrfRibFilterIpv6"]
					}
					if oVrf.RibFilter.Ipv6.Static != nil {
						nestedVrf.RibFilter.Ipv6.Static = &VrfRibFilterIpv6StaticXml{}
						if _, ok := o.Misc["VrfRibFilterIpv6Static"]; ok {
							nestedVrf.RibFilter.Ipv6.Static.Misc = o.Misc["VrfRibFilterIpv6Static"]
						}
						if oVrf.RibFilter.Ipv6.Static.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Static.RouteMap = oVrf.RibFilter.Ipv6.Static.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv6.Bgp != nil {
						nestedVrf.RibFilter.Ipv6.Bgp = &VrfRibFilterIpv6BgpXml{}
						if _, ok := o.Misc["VrfRibFilterIpv6Bgp"]; ok {
							nestedVrf.RibFilter.Ipv6.Bgp.Misc = o.Misc["VrfRibFilterIpv6Bgp"]
						}
						if oVrf.RibFilter.Ipv6.Bgp.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Bgp.RouteMap = oVrf.RibFilter.Ipv6.Bgp.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv6.Ospfv3 != nil {
						nestedVrf.RibFilter.Ipv6.Ospfv3 = &VrfRibFilterIpv6Ospfv3Xml{}
						if _, ok := o.Misc["VrfRibFilterIpv6Ospfv3"]; ok {
							nestedVrf.RibFilter.Ipv6.Ospfv3.Misc = o.Misc["VrfRibFilterIpv6Ospfv3"]
						}
						if oVrf.RibFilter.Ipv6.Ospfv3.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Ospfv3.RouteMap = oVrf.RibFilter.Ipv6.Ospfv3.RouteMap
						}
					}
				}
			}
			if oVrf.Bgp != nil {
				nestedVrf.Bgp = &VrfBgpXml{}
				if _, ok := o.Misc["VrfBgp"]; ok {
					nestedVrf.Bgp.Misc = o.Misc["VrfBgp"]
				}
				if oVrf.Bgp.Enable != nil {
					nestedVrf.Bgp.Enable = util.YesNo(oVrf.Bgp.Enable, nil)
				}
				if oVrf.Bgp.RouterId != nil {
					nestedVrf.Bgp.RouterId = oVrf.Bgp.RouterId
				}
				if oVrf.Bgp.LocalAs != nil {
					nestedVrf.Bgp.LocalAs = oVrf.Bgp.LocalAs
				}
				if oVrf.Bgp.InstallRoute != nil {
					nestedVrf.Bgp.InstallRoute = util.YesNo(oVrf.Bgp.InstallRoute, nil)
				}
				if oVrf.Bgp.EnforceFirstAs != nil {
					nestedVrf.Bgp.EnforceFirstAs = util.YesNo(oVrf.Bgp.EnforceFirstAs, nil)
				}
				if oVrf.Bgp.FastExternalFailover != nil {
					nestedVrf.Bgp.FastExternalFailover = util.YesNo(oVrf.Bgp.FastExternalFailover, nil)
				}
				if oVrf.Bgp.EcmpMultiAs != nil {
					nestedVrf.Bgp.EcmpMultiAs = util.YesNo(oVrf.Bgp.EcmpMultiAs, nil)
				}
				if oVrf.Bgp.DefaultLocalPreference != nil {
					nestedVrf.Bgp.DefaultLocalPreference = oVrf.Bgp.DefaultLocalPreference
				}
				if oVrf.Bgp.GracefulShutdown != nil {
					nestedVrf.Bgp.GracefulShutdown = util.YesNo(oVrf.Bgp.GracefulShutdown, nil)
				}
				if oVrf.Bgp.AlwaysAdvertiseNetworkRoute != nil {
					nestedVrf.Bgp.AlwaysAdvertiseNetworkRoute = util.YesNo(oVrf.Bgp.AlwaysAdvertiseNetworkRoute, nil)
				}
				if oVrf.Bgp.Med != nil {
					nestedVrf.Bgp.Med = &VrfBgpMedXml{}
					if _, ok := o.Misc["VrfBgpMed"]; ok {
						nestedVrf.Bgp.Med.Misc = o.Misc["VrfBgpMed"]
					}
					if oVrf.Bgp.Med.AlwaysCompareMed != nil {
						nestedVrf.Bgp.Med.AlwaysCompareMed = util.YesNo(oVrf.Bgp.Med.AlwaysCompareMed, nil)
					}
					if oVrf.Bgp.Med.DeterministicMedComparison != nil {
						nestedVrf.Bgp.Med.DeterministicMedComparison = util.YesNo(oVrf.Bgp.Med.DeterministicMedComparison, nil)
					}
				}
				if oVrf.Bgp.GracefulRestart != nil {
					nestedVrf.Bgp.GracefulRestart = &VrfBgpGracefulRestartXml{}
					if _, ok := o.Misc["VrfBgpGracefulRestart"]; ok {
						nestedVrf.Bgp.GracefulRestart.Misc = o.Misc["VrfBgpGracefulRestart"]
					}
					if oVrf.Bgp.GracefulRestart.Enable != nil {
						nestedVrf.Bgp.GracefulRestart.Enable = util.YesNo(oVrf.Bgp.GracefulRestart.Enable, nil)
					}
					if oVrf.Bgp.GracefulRestart.StaleRouteTime != nil {
						nestedVrf.Bgp.GracefulRestart.StaleRouteTime = oVrf.Bgp.GracefulRestart.StaleRouteTime
					}
					if oVrf.Bgp.GracefulRestart.MaxPeerRestartTime != nil {
						nestedVrf.Bgp.GracefulRestart.MaxPeerRestartTime = oVrf.Bgp.GracefulRestart.MaxPeerRestartTime
					}
					if oVrf.Bgp.GracefulRestart.LocalRestartTime != nil {
						nestedVrf.Bgp.GracefulRestart.LocalRestartTime = oVrf.Bgp.GracefulRestart.LocalRestartTime
					}
				}
				if oVrf.Bgp.GlobalBfd != nil {
					nestedVrf.Bgp.GlobalBfd = &VrfBgpGlobalBfdXml{}
					if _, ok := o.Misc["VrfBgpGlobalBfd"]; ok {
						nestedVrf.Bgp.GlobalBfd.Misc = o.Misc["VrfBgpGlobalBfd"]
					}
					if oVrf.Bgp.GlobalBfd.Profile != nil {
						nestedVrf.Bgp.GlobalBfd.Profile = oVrf.Bgp.GlobalBfd.Profile
					}
				}
				if oVrf.Bgp.RedistributionProfile != nil {
					nestedVrf.Bgp.RedistributionProfile = &VrfBgpRedistributionProfileXml{}
					if _, ok := o.Misc["VrfBgpRedistributionProfile"]; ok {
						nestedVrf.Bgp.RedistributionProfile.Misc = o.Misc["VrfBgpRedistributionProfile"]
					}
					if oVrf.Bgp.RedistributionProfile.Ipv4 != nil {
						nestedVrf.Bgp.RedistributionProfile.Ipv4 = &VrfBgpRedistributionProfileIpv4Xml{}
						if _, ok := o.Misc["VrfBgpRedistributionProfileIpv4"]; ok {
							nestedVrf.Bgp.RedistributionProfile.Ipv4.Misc = o.Misc["VrfBgpRedistributionProfileIpv4"]
						}
						if oVrf.Bgp.RedistributionProfile.Ipv4.Unicast != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv4.Unicast = oVrf.Bgp.RedistributionProfile.Ipv4.Unicast
						}
					}
					if oVrf.Bgp.RedistributionProfile.Ipv6 != nil {
						nestedVrf.Bgp.RedistributionProfile.Ipv6 = &VrfBgpRedistributionProfileIpv6Xml{}
						if _, ok := o.Misc["VrfBgpRedistributionProfileIpv6"]; ok {
							nestedVrf.Bgp.RedistributionProfile.Ipv6.Misc = o.Misc["VrfBgpRedistributionProfileIpv6"]
						}
						if oVrf.Bgp.RedistributionProfile.Ipv6.Unicast != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv6.Unicast = oVrf.Bgp.RedistributionProfile.Ipv6.Unicast
						}
					}
				}
				if oVrf.Bgp.AdvertiseNetwork != nil {
					nestedVrf.Bgp.AdvertiseNetwork = &VrfBgpAdvertiseNetworkXml{}
					if _, ok := o.Misc["VrfBgpAdvertiseNetwork"]; ok {
						nestedVrf.Bgp.AdvertiseNetwork.Misc = o.Misc["VrfBgpAdvertiseNetwork"]
					}
					if oVrf.Bgp.AdvertiseNetwork.Ipv4 != nil {
						nestedVrf.Bgp.AdvertiseNetwork.Ipv4 = &VrfBgpAdvertiseNetworkIpv4Xml{}
						if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv4"]; ok {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv4"]
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv4.Network != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = []VrfBgpAdvertiseNetworkIpv4NetworkXml{}
							for _, oVrfBgpAdvertiseNetworkIpv4Network := range oVrf.Bgp.AdvertiseNetwork.Ipv4.Network {
								nestedVrfBgpAdvertiseNetworkIpv4Network := VrfBgpAdvertiseNetworkIpv4NetworkXml{}
								if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv4Network"]; ok {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv4Network"]
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Name != "" {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Name = oVrfBgpAdvertiseNetworkIpv4Network.Name
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Unicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Unicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Unicast, nil)
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Multicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Multicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Multicast, nil)
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Backdoor != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Backdoor = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Backdoor, nil)
								}
								nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network, nestedVrfBgpAdvertiseNetworkIpv4Network)
							}
						}
					}
					if oVrf.Bgp.AdvertiseNetwork.Ipv6 != nil {
						nestedVrf.Bgp.AdvertiseNetwork.Ipv6 = &VrfBgpAdvertiseNetworkIpv6Xml{}
						if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv6"]; ok {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv6"]
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv6.Network != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = []VrfBgpAdvertiseNetworkIpv6NetworkXml{}
							for _, oVrfBgpAdvertiseNetworkIpv6Network := range oVrf.Bgp.AdvertiseNetwork.Ipv6.Network {
								nestedVrfBgpAdvertiseNetworkIpv6Network := VrfBgpAdvertiseNetworkIpv6NetworkXml{}
								if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv6Network"]; ok {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv6Network"]
								}
								if oVrfBgpAdvertiseNetworkIpv6Network.Name != "" {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Name = oVrfBgpAdvertiseNetworkIpv6Network.Name
								}
								if oVrfBgpAdvertiseNetworkIpv6Network.Unicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Unicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv6Network.Unicast, nil)
								}
								nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network, nestedVrfBgpAdvertiseNetworkIpv6Network)
							}
						}
					}
				}
				if oVrf.Bgp.PeerGroup != nil {
					nestedVrf.Bgp.PeerGroup = []VrfBgpPeerGroupXml{}
					for _, oVrfBgpPeerGroup := range oVrf.Bgp.PeerGroup {
						nestedVrfBgpPeerGroup := VrfBgpPeerGroupXml{}
						if _, ok := o.Misc["VrfBgpPeerGroup"]; ok {
							nestedVrfBgpPeerGroup.Misc = o.Misc["VrfBgpPeerGroup"]
						}
						if oVrfBgpPeerGroup.Name != "" {
							nestedVrfBgpPeerGroup.Name = oVrfBgpPeerGroup.Name
						}
						if oVrfBgpPeerGroup.Enable != nil {
							nestedVrfBgpPeerGroup.Enable = util.YesNo(oVrfBgpPeerGroup.Enable, nil)
						}
						if oVrfBgpPeerGroup.Type != nil {
							nestedVrfBgpPeerGroup.Type = &VrfBgpPeerGroupTypeXml{}
							if _, ok := o.Misc["VrfBgpPeerGroupType"]; ok {
								nestedVrfBgpPeerGroup.Type.Misc = o.Misc["VrfBgpPeerGroupType"]
							}
							if oVrfBgpPeerGroup.Type.Ibgp != nil {
								nestedVrfBgpPeerGroup.Type.Ibgp = &VrfBgpPeerGroupTypeIbgpXml{}
								if _, ok := o.Misc["VrfBgpPeerGroupTypeIbgp"]; ok {
									nestedVrfBgpPeerGroup.Type.Ibgp.Misc = o.Misc["VrfBgpPeerGroupTypeIbgp"]
								}
							}
							if oVrfBgpPeerGroup.Type.Ebgp != nil {
								nestedVrfBgpPeerGroup.Type.Ebgp = &VrfBgpPeerGroupTypeEbgpXml{}
								if _, ok := o.Misc["VrfBgpPeerGroupTypeEbgp"]; ok {
									nestedVrfBgpPeerGroup.Type.Ebgp.Misc = o.Misc["VrfBgpPeerGroupTypeEbgp"]
								}
							}
						}
						if oVrfBgpPeerGroup.AddressFamily != nil {
							nestedVrfBgpPeerGroup.AddressFamily = &VrfBgpPeerGroupAddressFamilyXml{}
							if _, ok := o.Misc["VrfBgpPeerGroupAddressFamily"]; ok {
								nestedVrfBgpPeerGroup.AddressFamily.Misc = o.Misc["VrfBgpPeerGroupAddressFamily"]
							}
							if oVrfBgpPeerGroup.AddressFamily.Ipv4 != nil {
								nestedVrfBgpPeerGroup.AddressFamily.Ipv4 = oVrfBgpPeerGroup.AddressFamily.Ipv4
							}
							if oVrfBgpPeerGroup.AddressFamily.Ipv6 != nil {
								nestedVrfBgpPeerGroup.AddressFamily.Ipv6 = oVrfBgpPeerGroup.AddressFamily.Ipv6
							}
						}
						if oVrfBgpPeerGroup.FilteringProfile != nil {
							nestedVrfBgpPeerGroup.FilteringProfile = &VrfBgpPeerGroupFilteringProfileXml{}
							if _, ok := o.Misc["VrfBgpPeerGroupFilteringProfile"]; ok {
								nestedVrfBgpPeerGroup.FilteringProfile.Misc = o.Misc["VrfBgpPeerGroupFilteringProfile"]
							}
							if oVrfBgpPeerGroup.FilteringProfile.Ipv4 != nil {
								nestedVrfBgpPeerGroup.FilteringProfile.Ipv4 = oVrfBgpPeerGroup.FilteringProfile.Ipv4
							}
							if oVrfBgpPeerGroup.FilteringProfile.Ipv6 != nil {
								nestedVrfBgpPeerGroup.FilteringProfile.Ipv6 = oVrfBgpPeerGroup.FilteringProfile.Ipv6
							}
						}
						if oVrfBgpPeerGroup.ConnectionOptions != nil {
							nestedVrfBgpPeerGroup.ConnectionOptions = &VrfBgpPeerGroupConnectionOptionsXml{}
							if _, ok := o.Misc["VrfBgpPeerGroupConnectionOptions"]; ok {
								nestedVrfBgpPeerGroup.ConnectionOptions.Misc = o.Misc["VrfBgpPeerGroupConnectionOptions"]
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Timers != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Timers = oVrfBgpPeerGroup.ConnectionOptions.Timers
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Multihop != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Multihop = oVrfBgpPeerGroup.ConnectionOptions.Multihop
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Authentication != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Authentication = oVrfBgpPeerGroup.ConnectionOptions.Authentication
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Dampening != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Dampening = oVrfBgpPeerGroup.ConnectionOptions.Dampening
							}
						}
						if oVrfBgpPeerGroup.Peer != nil {
							nestedVrfBgpPeerGroup.Peer = []VrfBgpPeerGroupPeerXml{}
							for _, oVrfBgpPeerGroupPeer := range oVrfBgpPeerGroup.Peer {
								nestedVrfBgpPeerGroupPeer := VrfBgpPeerGroupPeerXml{}
								if _, ok := o.Misc["VrfBgpPeerGroupPeer"]; ok {
									nestedVrfBgpPeerGroupPeer.Misc = o.Misc["VrfBgpPeerGroupPeer"]
								}
								if oVrfBgpPeerGroupPeer.Name != "" {
									nestedVrfBgpPeerGroupPeer.Name = oVrfBgpPeerGroupPeer.Name
								}
								if oVrfBgpPeerGroupPeer.Enable != nil {
									nestedVrfBgpPeerGroupPeer.Enable = util.YesNo(oVrfBgpPeerGroupPeer.Enable, nil)
								}
								if oVrfBgpPeerGroupPeer.Passive != nil {
									nestedVrfBgpPeerGroupPeer.Passive = util.YesNo(oVrfBgpPeerGroupPeer.Passive, nil)
								}
								if oVrfBgpPeerGroupPeer.PeerAs != nil {
									nestedVrfBgpPeerGroupPeer.PeerAs = oVrfBgpPeerGroupPeer.PeerAs
								}
								if oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
									nestedVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.YesNo(oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
								}
								if oVrfBgpPeerGroupPeer.Inherit != nil {
									nestedVrfBgpPeerGroupPeer.Inherit = &VrfBgpPeerGroupPeerInheritXml{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerInherit"]; ok {
										nestedVrfBgpPeerGroupPeer.Inherit.Misc = o.Misc["VrfBgpPeerGroupPeerInherit"]
									}
									if oVrfBgpPeerGroupPeer.Inherit.Yes != nil {
										nestedVrfBgpPeerGroupPeer.Inherit.Yes = &VrfBgpPeerGroupPeerInheritYesXml{}
										if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritYes"]; ok {
											nestedVrfBgpPeerGroupPeer.Inherit.Yes.Misc = o.Misc["VrfBgpPeerGroupPeerInheritYes"]
										}
									}
									if oVrfBgpPeerGroupPeer.Inherit.No != nil {
										nestedVrfBgpPeerGroupPeer.Inherit.No = &VrfBgpPeerGroupPeerInheritNoXml{}
										if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNo"]; ok {
											nestedVrfBgpPeerGroupPeer.Inherit.No.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNo"]
										}
										if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily = &VrfBgpPeerGroupPeerInheritNoAddressFamilyXml{}
											if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"]; ok {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"]
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6
											}
										}
										if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile = &VrfBgpPeerGroupPeerInheritNoFilteringProfileXml{}
											if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"]; ok {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"]
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6
											}
										}
									}
								}
								if oVrfBgpPeerGroupPeer.LocalAddress != nil {
									nestedVrfBgpPeerGroupPeer.LocalAddress = &VrfBgpPeerGroupPeerLocalAddressXml{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerLocalAddress"]; ok {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Misc = o.Misc["VrfBgpPeerGroupPeerLocalAddress"]
									}
									if oVrfBgpPeerGroupPeer.LocalAddress.Interface != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Interface = oVrfBgpPeerGroupPeer.LocalAddress.Interface
									}
									if oVrfBgpPeerGroupPeer.LocalAddress.Ip != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Ip = oVrfBgpPeerGroupPeer.LocalAddress.Ip
									}
								}
								if oVrfBgpPeerGroupPeer.PeerAddress != nil {
									nestedVrfBgpPeerGroupPeer.PeerAddress = &VrfBgpPeerGroupPeerPeerAddressXml{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerPeerAddress"]; ok {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Misc = o.Misc["VrfBgpPeerGroupPeerPeerAddress"]
									}
									if oVrfBgpPeerGroupPeer.PeerAddress.Ip != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Ip = oVrfBgpPeerGroupPeer.PeerAddress.Ip
									}
									if oVrfBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Fqdn = oVrfBgpPeerGroupPeer.PeerAddress.Fqdn
									}
								}
								if oVrfBgpPeerGroupPeer.ConnectionOptions != nil {
									nestedVrfBgpPeerGroupPeer.ConnectionOptions = &VrfBgpPeerGroupPeerConnectionOptionsXml{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerConnectionOptions"]; ok {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Misc = o.Misc["VrfBgpPeerGroupPeerConnectionOptions"]
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Timers != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Timers = oVrfBgpPeerGroupPeer.ConnectionOptions.Timers
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Multihop = oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Authentication = oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Dampening = oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening
									}
								}
								if oVrfBgpPeerGroupPeer.Bfd != nil {
									nestedVrfBgpPeerGroupPeer.Bfd = &VrfBgpPeerGroupPeerBfdXml{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerBfd"]; ok {
										nestedVrfBgpPeerGroupPeer.Bfd.Misc = o.Misc["VrfBgpPeerGroupPeerBfd"]
									}
									if oVrfBgpPeerGroupPeer.Bfd.Profile != nil {
										nestedVrfBgpPeerGroupPeer.Bfd.Profile = oVrfBgpPeerGroupPeer.Bfd.Profile
									}
								}
								nestedVrfBgpPeerGroup.Peer = append(nestedVrfBgpPeerGroup.Peer, nestedVrfBgpPeerGroupPeer)
							}
						}
						nestedVrf.Bgp.PeerGroup = append(nestedVrf.Bgp.PeerGroup, nestedVrfBgpPeerGroup)
					}
				}
				if oVrf.Bgp.AggregateRoutes != nil {
					nestedVrf.Bgp.AggregateRoutes = []VrfBgpAggregateRoutesXml{}
					for _, oVrfBgpAggregateRoutes := range oVrf.Bgp.AggregateRoutes {
						nestedVrfBgpAggregateRoutes := VrfBgpAggregateRoutesXml{}
						if _, ok := o.Misc["VrfBgpAggregateRoutes"]; ok {
							nestedVrfBgpAggregateRoutes.Misc = o.Misc["VrfBgpAggregateRoutes"]
						}
						if oVrfBgpAggregateRoutes.Name != "" {
							nestedVrfBgpAggregateRoutes.Name = oVrfBgpAggregateRoutes.Name
						}
						if oVrfBgpAggregateRoutes.Description != nil {
							nestedVrfBgpAggregateRoutes.Description = oVrfBgpAggregateRoutes.Description
						}
						if oVrfBgpAggregateRoutes.Enable != nil {
							nestedVrfBgpAggregateRoutes.Enable = util.YesNo(oVrfBgpAggregateRoutes.Enable, nil)
						}
						if oVrfBgpAggregateRoutes.SummaryOnly != nil {
							nestedVrfBgpAggregateRoutes.SummaryOnly = util.YesNo(oVrfBgpAggregateRoutes.SummaryOnly, nil)
						}
						if oVrfBgpAggregateRoutes.AsSet != nil {
							nestedVrfBgpAggregateRoutes.AsSet = util.YesNo(oVrfBgpAggregateRoutes.AsSet, nil)
						}
						if oVrfBgpAggregateRoutes.SameMed != nil {
							nestedVrfBgpAggregateRoutes.SameMed = util.YesNo(oVrfBgpAggregateRoutes.SameMed, nil)
						}
						if oVrfBgpAggregateRoutes.Type != nil {
							nestedVrfBgpAggregateRoutes.Type = &VrfBgpAggregateRoutesTypeXml{}
							if _, ok := o.Misc["VrfBgpAggregateRoutesType"]; ok {
								nestedVrfBgpAggregateRoutes.Type.Misc = o.Misc["VrfBgpAggregateRoutesType"]
							}
							if oVrfBgpAggregateRoutes.Type.Ipv4 != nil {
								nestedVrfBgpAggregateRoutes.Type.Ipv4 = &VrfBgpAggregateRoutesTypeIpv4Xml{}
								if _, ok := o.Misc["VrfBgpAggregateRoutesTypeIpv4"]; ok {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.Misc = o.Misc["VrfBgpAggregateRoutesTypeIpv4"]
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap
								}
							}
							if oVrfBgpAggregateRoutes.Type.Ipv6 != nil {
								nestedVrfBgpAggregateRoutes.Type.Ipv6 = &VrfBgpAggregateRoutesTypeIpv6Xml{}
								if _, ok := o.Misc["VrfBgpAggregateRoutesTypeIpv6"]; ok {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.Misc = o.Misc["VrfBgpAggregateRoutesTypeIpv6"]
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap
								}
							}
						}
						nestedVrf.Bgp.AggregateRoutes = append(nestedVrf.Bgp.AggregateRoutes, nestedVrfBgpAggregateRoutes)
					}
				}
			}
			if oVrf.RoutingTable != nil {
				nestedVrf.RoutingTable = &VrfRoutingTableXml{}
				if _, ok := o.Misc["VrfRoutingTable"]; ok {
					nestedVrf.RoutingTable.Misc = o.Misc["VrfRoutingTable"]
				}
				if oVrf.RoutingTable.Ip != nil {
					nestedVrf.RoutingTable.Ip = &VrfRoutingTableIpXml{}
					if _, ok := o.Misc["VrfRoutingTableIp"]; ok {
						nestedVrf.RoutingTable.Ip.Misc = o.Misc["VrfRoutingTableIp"]
					}
					if oVrf.RoutingTable.Ip.StaticRoute != nil {
						nestedVrf.RoutingTable.Ip.StaticRoute = []VrfRoutingTableIpStaticRouteXml{}
						for _, oVrfRoutingTableIpStaticRoute := range oVrf.RoutingTable.Ip.StaticRoute {
							nestedVrfRoutingTableIpStaticRoute := VrfRoutingTableIpStaticRouteXml{}
							if _, ok := o.Misc["VrfRoutingTableIpStaticRoute"]; ok {
								nestedVrfRoutingTableIpStaticRoute.Misc = o.Misc["VrfRoutingTableIpStaticRoute"]
							}
							if oVrfRoutingTableIpStaticRoute.Name != "" {
								nestedVrfRoutingTableIpStaticRoute.Name = oVrfRoutingTableIpStaticRoute.Name
							}
							if oVrfRoutingTableIpStaticRoute.Destination != nil {
								nestedVrfRoutingTableIpStaticRoute.Destination = oVrfRoutingTableIpStaticRoute.Destination
							}
							if oVrfRoutingTableIpStaticRoute.Interface != nil {
								nestedVrfRoutingTableIpStaticRoute.Interface = oVrfRoutingTableIpStaticRoute.Interface
							}
							if oVrfRoutingTableIpStaticRoute.AdminDist != nil {
								nestedVrfRoutingTableIpStaticRoute.AdminDist = oVrfRoutingTableIpStaticRoute.AdminDist
							}
							if oVrfRoutingTableIpStaticRoute.Metric != nil {
								nestedVrfRoutingTableIpStaticRoute.Metric = oVrfRoutingTableIpStaticRoute.Metric
							}
							if oVrfRoutingTableIpStaticRoute.Nexthop != nil {
								nestedVrfRoutingTableIpStaticRoute.Nexthop = &VrfRoutingTableIpStaticRouteNexthopXml{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRouteNexthop"]; ok {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Misc = o.Misc["VrfRoutingTableIpStaticRouteNexthop"]
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.Discard != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard = &VrfRoutingTableIpStaticRouteNexthopDiscardXml{}
									if _, ok := o.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"]; ok {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc = o.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"]
									}
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.IpAddress = oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.NextLr != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.NextLr = oVrfRoutingTableIpStaticRoute.Nexthop.NextLr
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn
								}
							}
							if oVrfRoutingTableIpStaticRoute.Bfd != nil {
								nestedVrfRoutingTableIpStaticRoute.Bfd = &VrfRoutingTableIpStaticRouteBfdXml{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRouteBfd"]; ok {
									nestedVrfRoutingTableIpStaticRoute.Bfd.Misc = o.Misc["VrfRoutingTableIpStaticRouteBfd"]
								}
								if oVrfRoutingTableIpStaticRoute.Bfd.Profile != nil {
									nestedVrfRoutingTableIpStaticRoute.Bfd.Profile = oVrfRoutingTableIpStaticRoute.Bfd.Profile
								}
							}
							if oVrfRoutingTableIpStaticRoute.PathMonitor != nil {
								nestedVrfRoutingTableIpStaticRoute.PathMonitor = &VrfRoutingTableIpStaticRoutePathMonitorXml{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRoutePathMonitor"]; ok {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.Misc = o.Misc["VrfRoutingTableIpStaticRoutePathMonitor"]
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.Enable = util.YesNo(oVrfRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml{}
									for _, oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
										nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml{}
										if _, ok := o.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"]; ok {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"]
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
										}
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
									}
								}
							}
							nestedVrf.RoutingTable.Ip.StaticRoute = append(nestedVrf.RoutingTable.Ip.StaticRoute, nestedVrfRoutingTableIpStaticRoute)
						}
					}
				}
				if oVrf.RoutingTable.Ipv6 != nil {
					nestedVrf.RoutingTable.Ipv6 = &VrfRoutingTableIpv6Xml{}
					if _, ok := o.Misc["VrfRoutingTableIpv6"]; ok {
						nestedVrf.RoutingTable.Ipv6.Misc = o.Misc["VrfRoutingTableIpv6"]
					}
					if oVrf.RoutingTable.Ipv6.StaticRoute != nil {
						nestedVrf.RoutingTable.Ipv6.StaticRoute = []VrfRoutingTableIpv6StaticRouteXml{}
						for _, oVrfRoutingTableIpv6StaticRoute := range oVrf.RoutingTable.Ipv6.StaticRoute {
							nestedVrfRoutingTableIpv6StaticRoute := VrfRoutingTableIpv6StaticRouteXml{}
							if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoute"]; ok {
								nestedVrfRoutingTableIpv6StaticRoute.Misc = o.Misc["VrfRoutingTableIpv6StaticRoute"]
							}
							if oVrfRoutingTableIpv6StaticRoute.Name != "" {
								nestedVrfRoutingTableIpv6StaticRoute.Name = oVrfRoutingTableIpv6StaticRoute.Name
							}
							if oVrfRoutingTableIpv6StaticRoute.Destination != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Destination = oVrfRoutingTableIpv6StaticRoute.Destination
							}
							if oVrfRoutingTableIpv6StaticRoute.Interface != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Interface = oVrfRoutingTableIpv6StaticRoute.Interface
							}
							if oVrfRoutingTableIpv6StaticRoute.AdminDist != nil {
								nestedVrfRoutingTableIpv6StaticRoute.AdminDist = oVrfRoutingTableIpv6StaticRoute.AdminDist
							}
							if oVrfRoutingTableIpv6StaticRoute.Metric != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Metric = oVrfRoutingTableIpv6StaticRoute.Metric
							}
							if oVrfRoutingTableIpv6StaticRoute.Nexthop != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Nexthop = &VrfRoutingTableIpv6StaticRouteNexthopXml{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteNexthop"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteNexthop"]
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard = &VrfRoutingTableIpv6StaticRouteNexthopDiscardXml{}
									if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"]; ok {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"]
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr = oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr
								}
							}
							if oVrfRoutingTableIpv6StaticRoute.Bfd != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Bfd = &VrfRoutingTableIpv6StaticRouteBfdXml{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteBfd"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteBfd"]
								}
								if oVrfRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd.Profile = oVrfRoutingTableIpv6StaticRoute.Bfd.Profile
								}
							}
							if oVrfRoutingTableIpv6StaticRoute.PathMonitor != nil {
								nestedVrfRoutingTableIpv6StaticRoute.PathMonitor = &VrfRoutingTableIpv6StaticRoutePathMonitorXml{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc = o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"]
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.YesNo(oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml{}
									for _, oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
										nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml{}
										if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]; ok {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
										}
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
									}
								}
							}
							nestedVrf.RoutingTable.Ipv6.StaticRoute = append(nestedVrf.RoutingTable.Ipv6.StaticRoute, nestedVrfRoutingTableIpv6StaticRoute)
						}
					}
				}
			}
			if oVrf.Ospf != nil {
				nestedVrf.Ospf = &VrfOspfXml{}
				if _, ok := o.Misc["VrfOspf"]; ok {
					nestedVrf.Ospf.Misc = o.Misc["VrfOspf"]
				}
				if oVrf.Ospf.RouterId != nil {
					nestedVrf.Ospf.RouterId = oVrf.Ospf.RouterId
				}
				if oVrf.Ospf.Enable != nil {
					nestedVrf.Ospf.Enable = util.YesNo(oVrf.Ospf.Enable, nil)
				}
				if oVrf.Ospf.Rfc1583 != nil {
					nestedVrf.Ospf.Rfc1583 = util.YesNo(oVrf.Ospf.Rfc1583, nil)
				}
				if oVrf.Ospf.SpfTimer != nil {
					nestedVrf.Ospf.SpfTimer = oVrf.Ospf.SpfTimer
				}
				if oVrf.Ospf.GlobalIfTimer != nil {
					nestedVrf.Ospf.GlobalIfTimer = oVrf.Ospf.GlobalIfTimer
				}
				if oVrf.Ospf.RedistributionProfile != nil {
					nestedVrf.Ospf.RedistributionProfile = oVrf.Ospf.RedistributionProfile
				}
				if oVrf.Ospf.GlobalBfd != nil {
					nestedVrf.Ospf.GlobalBfd = &VrfOspfGlobalBfdXml{}
					if _, ok := o.Misc["VrfOspfGlobalBfd"]; ok {
						nestedVrf.Ospf.GlobalBfd.Misc = o.Misc["VrfOspfGlobalBfd"]
					}
					if oVrf.Ospf.GlobalBfd.Profile != nil {
						nestedVrf.Ospf.GlobalBfd.Profile = oVrf.Ospf.GlobalBfd.Profile
					}
				}
				if oVrf.Ospf.GracefulRestart != nil {
					nestedVrf.Ospf.GracefulRestart = &VrfOspfGracefulRestartXml{}
					if _, ok := o.Misc["VrfOspfGracefulRestart"]; ok {
						nestedVrf.Ospf.GracefulRestart.Misc = o.Misc["VrfOspfGracefulRestart"]
					}
					if oVrf.Ospf.GracefulRestart.Enable != nil {
						nestedVrf.Ospf.GracefulRestart.Enable = util.YesNo(oVrf.Ospf.GracefulRestart.Enable, nil)
					}
					if oVrf.Ospf.GracefulRestart.GracePeriod != nil {
						nestedVrf.Ospf.GracefulRestart.GracePeriod = oVrf.Ospf.GracefulRestart.GracePeriod
					}
					if oVrf.Ospf.GracefulRestart.HelperEnable != nil {
						nestedVrf.Ospf.GracefulRestart.HelperEnable = util.YesNo(oVrf.Ospf.GracefulRestart.HelperEnable, nil)
					}
					if oVrf.Ospf.GracefulRestart.StrictLSAChecking != nil {
						nestedVrf.Ospf.GracefulRestart.StrictLSAChecking = util.YesNo(oVrf.Ospf.GracefulRestart.StrictLSAChecking, nil)
					}
					if oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedVrf.Ospf.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime
					}
				}
				if oVrf.Ospf.Area != nil {
					nestedVrf.Ospf.Area = []VrfOspfAreaXml{}
					for _, oVrfOspfArea := range oVrf.Ospf.Area {
						nestedVrfOspfArea := VrfOspfAreaXml{}
						if _, ok := o.Misc["VrfOspfArea"]; ok {
							nestedVrfOspfArea.Misc = o.Misc["VrfOspfArea"]
						}
						if oVrfOspfArea.Name != "" {
							nestedVrfOspfArea.Name = oVrfOspfArea.Name
						}
						if oVrfOspfArea.Authentication != nil {
							nestedVrfOspfArea.Authentication = oVrfOspfArea.Authentication
						}
						if oVrfOspfArea.Type != nil {
							nestedVrfOspfArea.Type = &VrfOspfAreaTypeXml{}
							if _, ok := o.Misc["VrfOspfAreaType"]; ok {
								nestedVrfOspfArea.Type.Misc = o.Misc["VrfOspfAreaType"]
							}
							if oVrfOspfArea.Type.Normal != nil {
								nestedVrfOspfArea.Type.Normal = &VrfOspfAreaTypeNormalXml{}
								if _, ok := o.Misc["VrfOspfAreaTypeNormal"]; ok {
									nestedVrfOspfArea.Type.Normal.Misc = o.Misc["VrfOspfAreaTypeNormal"]
								}
								if oVrfOspfArea.Type.Normal.Abr != nil {
									nestedVrfOspfArea.Type.Normal.Abr = &VrfOspfAreaTypeNormalAbrXml{}
									if _, ok := o.Misc["VrfOspfAreaTypeNormalAbr"]; ok {
										nestedVrfOspfArea.Type.Normal.Abr.Misc = o.Misc["VrfOspfAreaTypeNormalAbr"]
									}
									if oVrfOspfArea.Type.Normal.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.ImportList = oVrfOspfArea.Type.Normal.Abr.ImportList
									}
									if oVrfOspfArea.Type.Normal.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.ExportList = oVrfOspfArea.Type.Normal.Abr.ExportList
									}
									if oVrfOspfArea.Type.Normal.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.InboundFilterList = oVrfOspfArea.Type.Normal.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Normal.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.OutboundFilterList = oVrfOspfArea.Type.Normal.Abr.OutboundFilterList
									}
								}
							}
							if oVrfOspfArea.Type.Stub != nil {
								nestedVrfOspfArea.Type.Stub = &VrfOspfAreaTypeStubXml{}
								if _, ok := o.Misc["VrfOspfAreaTypeStub"]; ok {
									nestedVrfOspfArea.Type.Stub.Misc = o.Misc["VrfOspfAreaTypeStub"]
								}
								if oVrfOspfArea.Type.Stub.NoSummary != nil {
									nestedVrfOspfArea.Type.Stub.NoSummary = util.YesNo(oVrfOspfArea.Type.Stub.NoSummary, nil)
								}
								if oVrfOspfArea.Type.Stub.Abr != nil {
									nestedVrfOspfArea.Type.Stub.Abr = &VrfOspfAreaTypeStubAbrXml{}
									if _, ok := o.Misc["VrfOspfAreaTypeStubAbr"]; ok {
										nestedVrfOspfArea.Type.Stub.Abr.Misc = o.Misc["VrfOspfAreaTypeStubAbr"]
									}
									if oVrfOspfArea.Type.Stub.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.ImportList = oVrfOspfArea.Type.Stub.Abr.ImportList
									}
									if oVrfOspfArea.Type.Stub.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.ExportList = oVrfOspfArea.Type.Stub.Abr.ExportList
									}
									if oVrfOspfArea.Type.Stub.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.InboundFilterList = oVrfOspfArea.Type.Stub.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Stub.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.OutboundFilterList = oVrfOspfArea.Type.Stub.Abr.OutboundFilterList
									}
								}
								if oVrfOspfArea.Type.Stub.DefaultRouteMetric != nil {
									nestedVrfOspfArea.Type.Stub.DefaultRouteMetric = oVrfOspfArea.Type.Stub.DefaultRouteMetric
								}
							}
							if oVrfOspfArea.Type.Nssa != nil {
								nestedVrfOspfArea.Type.Nssa = &VrfOspfAreaTypeNssaXml{}
								if _, ok := o.Misc["VrfOspfAreaTypeNssa"]; ok {
									nestedVrfOspfArea.Type.Nssa.Misc = o.Misc["VrfOspfAreaTypeNssa"]
								}
								if oVrfOspfArea.Type.Nssa.NoSummary != nil {
									nestedVrfOspfArea.Type.Nssa.NoSummary = util.YesNo(oVrfOspfArea.Type.Nssa.NoSummary, nil)
								}
								if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate != nil {
									nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate = &VrfOspfAreaTypeNssaDefaultInformationOriginateXml{}
									if _, ok := o.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"]; ok {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc = o.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"]
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType
									}
								}
								if oVrfOspfArea.Type.Nssa.Abr != nil {
									nestedVrfOspfArea.Type.Nssa.Abr = &VrfOspfAreaTypeNssaAbrXml{}
									if _, ok := o.Misc["VrfOspfAreaTypeNssaAbr"]; ok {
										nestedVrfOspfArea.Type.Nssa.Abr.Misc = o.Misc["VrfOspfAreaTypeNssaAbr"]
									}
									if oVrfOspfArea.Type.Nssa.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.ImportList = oVrfOspfArea.Type.Nssa.Abr.ImportList
									}
									if oVrfOspfArea.Type.Nssa.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.ExportList = oVrfOspfArea.Type.Nssa.Abr.ExportList
									}
									if oVrfOspfArea.Type.Nssa.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.InboundFilterList = oVrfOspfArea.Type.Nssa.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.OutboundFilterList = oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList
									}
									if oVrfOspfArea.Type.Nssa.Abr.NssaExtRange != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = []VrfOspfAreaTypeNssaAbrNssaExtRangeXml{}
										for _, oVrfOspfAreaTypeNssaAbrNssaExtRange := range oVrfOspfArea.Type.Nssa.Abr.NssaExtRange {
											nestedVrfOspfAreaTypeNssaAbrNssaExtRange := VrfOspfAreaTypeNssaAbrNssaExtRangeXml{}
											if _, ok := o.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"]; ok {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Misc = o.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"]
											}
											if oVrfOspfAreaTypeNssaAbrNssaExtRange.Name != "" {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Name = oVrfOspfAreaTypeNssaAbrNssaExtRange.Name
											}
											if oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise != nil {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise = util.YesNo(oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise, nil)
											}
											nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfAreaTypeNssaAbrNssaExtRange)
										}
									}
								}
							}
						}
						if oVrfOspfArea.Range != nil {
							nestedVrfOspfArea.Range = []VrfOspfAreaRangeXml{}
							for _, oVrfOspfAreaRange := range oVrfOspfArea.Range {
								nestedVrfOspfAreaRange := VrfOspfAreaRangeXml{}
								if _, ok := o.Misc["VrfOspfAreaRange"]; ok {
									nestedVrfOspfAreaRange.Misc = o.Misc["VrfOspfAreaRange"]
								}
								if oVrfOspfAreaRange.Name != "" {
									nestedVrfOspfAreaRange.Name = oVrfOspfAreaRange.Name
								}
								if oVrfOspfAreaRange.Advertise != nil {
									nestedVrfOspfAreaRange.Advertise = util.YesNo(oVrfOspfAreaRange.Advertise, nil)
								}
								nestedVrfOspfArea.Range = append(nestedVrfOspfArea.Range, nestedVrfOspfAreaRange)
							}
						}
						if oVrfOspfArea.Interface != nil {
							nestedVrfOspfArea.Interface = []VrfOspfAreaInterfaceXml{}
							for _, oVrfOspfAreaInterface := range oVrfOspfArea.Interface {
								nestedVrfOspfAreaInterface := VrfOspfAreaInterfaceXml{}
								if _, ok := o.Misc["VrfOspfAreaInterface"]; ok {
									nestedVrfOspfAreaInterface.Misc = o.Misc["VrfOspfAreaInterface"]
								}
								if oVrfOspfAreaInterface.Name != "" {
									nestedVrfOspfAreaInterface.Name = oVrfOspfAreaInterface.Name
								}
								if oVrfOspfAreaInterface.Enable != nil {
									nestedVrfOspfAreaInterface.Enable = util.YesNo(oVrfOspfAreaInterface.Enable, nil)
								}
								if oVrfOspfAreaInterface.MtuIgnore != nil {
									nestedVrfOspfAreaInterface.MtuIgnore = util.YesNo(oVrfOspfAreaInterface.MtuIgnore, nil)
								}
								if oVrfOspfAreaInterface.Passive != nil {
									nestedVrfOspfAreaInterface.Passive = util.YesNo(oVrfOspfAreaInterface.Passive, nil)
								}
								if oVrfOspfAreaInterface.Priority != nil {
									nestedVrfOspfAreaInterface.Priority = oVrfOspfAreaInterface.Priority
								}
								if oVrfOspfAreaInterface.Metric != nil {
									nestedVrfOspfAreaInterface.Metric = oVrfOspfAreaInterface.Metric
								}
								if oVrfOspfAreaInterface.Authentication != nil {
									nestedVrfOspfAreaInterface.Authentication = oVrfOspfAreaInterface.Authentication
								}
								if oVrfOspfAreaInterface.Timing != nil {
									nestedVrfOspfAreaInterface.Timing = oVrfOspfAreaInterface.Timing
								}
								if oVrfOspfAreaInterface.LinkType != nil {
									nestedVrfOspfAreaInterface.LinkType = &VrfOspfAreaInterfaceLinkTypeXml{}
									if _, ok := o.Misc["VrfOspfAreaInterfaceLinkType"]; ok {
										nestedVrfOspfAreaInterface.LinkType.Misc = o.Misc["VrfOspfAreaInterfaceLinkType"]
									}
									if oVrfOspfAreaInterface.LinkType.Broadcast != nil {
										nestedVrfOspfAreaInterface.LinkType.Broadcast = &VrfOspfAreaInterfaceLinkTypeBroadcastXml{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"]; ok {
											nestedVrfOspfAreaInterface.LinkType.Broadcast.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"]
										}
									}
									if oVrfOspfAreaInterface.LinkType.P2p != nil {
										nestedVrfOspfAreaInterface.LinkType.P2p = &VrfOspfAreaInterfaceLinkTypeP2pXml{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2p"]; ok {
											nestedVrfOspfAreaInterface.LinkType.P2p.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2p"]
										}
									}
									if oVrfOspfAreaInterface.LinkType.P2mp != nil {
										nestedVrfOspfAreaInterface.LinkType.P2mp = &VrfOspfAreaInterfaceLinkTypeP2mpXml{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"]; ok {
											nestedVrfOspfAreaInterface.LinkType.P2mp.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"]
										}
										if oVrfOspfAreaInterface.LinkType.P2mp.Neighbor != nil {
											nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = []VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml{}
											for _, oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfAreaInterface.LinkType.P2mp.Neighbor {
												nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml{}
												if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"]; ok {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"]
												}
												if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name
												}
												if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority
												}
												nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor)
											}
										}
									}
								}
								if oVrfOspfAreaInterface.Bfd != nil {
									nestedVrfOspfAreaInterface.Bfd = &VrfOspfAreaInterfaceBfdXml{}
									if _, ok := o.Misc["VrfOspfAreaInterfaceBfd"]; ok {
										nestedVrfOspfAreaInterface.Bfd.Misc = o.Misc["VrfOspfAreaInterfaceBfd"]
									}
									if oVrfOspfAreaInterface.Bfd.Profile != nil {
										nestedVrfOspfAreaInterface.Bfd.Profile = oVrfOspfAreaInterface.Bfd.Profile
									}
								}
								nestedVrfOspfArea.Interface = append(nestedVrfOspfArea.Interface, nestedVrfOspfAreaInterface)
							}
						}
						if oVrfOspfArea.VirtualLink != nil {
							nestedVrfOspfArea.VirtualLink = []VrfOspfAreaVirtualLinkXml{}
							for _, oVrfOspfAreaVirtualLink := range oVrfOspfArea.VirtualLink {
								nestedVrfOspfAreaVirtualLink := VrfOspfAreaVirtualLinkXml{}
								if _, ok := o.Misc["VrfOspfAreaVirtualLink"]; ok {
									nestedVrfOspfAreaVirtualLink.Misc = o.Misc["VrfOspfAreaVirtualLink"]
								}
								if oVrfOspfAreaVirtualLink.Name != "" {
									nestedVrfOspfAreaVirtualLink.Name = oVrfOspfAreaVirtualLink.Name
								}
								if oVrfOspfAreaVirtualLink.NeighborId != nil {
									nestedVrfOspfAreaVirtualLink.NeighborId = oVrfOspfAreaVirtualLink.NeighborId
								}
								if oVrfOspfAreaVirtualLink.TransitAreaId != nil {
									nestedVrfOspfAreaVirtualLink.TransitAreaId = oVrfOspfAreaVirtualLink.TransitAreaId
								}
								if oVrfOspfAreaVirtualLink.Enable != nil {
									nestedVrfOspfAreaVirtualLink.Enable = util.YesNo(oVrfOspfAreaVirtualLink.Enable, nil)
								}
								if oVrfOspfAreaVirtualLink.InstanceId != nil {
									nestedVrfOspfAreaVirtualLink.InstanceId = oVrfOspfAreaVirtualLink.InstanceId
								}
								if oVrfOspfAreaVirtualLink.Timing != nil {
									nestedVrfOspfAreaVirtualLink.Timing = oVrfOspfAreaVirtualLink.Timing
								}
								if oVrfOspfAreaVirtualLink.Authentication != nil {
									nestedVrfOspfAreaVirtualLink.Authentication = oVrfOspfAreaVirtualLink.Authentication
								}
								if oVrfOspfAreaVirtualLink.Bfd != nil {
									nestedVrfOspfAreaVirtualLink.Bfd = &VrfOspfAreaVirtualLinkBfdXml{}
									if _, ok := o.Misc["VrfOspfAreaVirtualLinkBfd"]; ok {
										nestedVrfOspfAreaVirtualLink.Bfd.Misc = o.Misc["VrfOspfAreaVirtualLinkBfd"]
									}
									if oVrfOspfAreaVirtualLink.Bfd.Profile != nil {
										nestedVrfOspfAreaVirtualLink.Bfd.Profile = oVrfOspfAreaVirtualLink.Bfd.Profile
									}
								}
								nestedVrfOspfArea.VirtualLink = append(nestedVrfOspfArea.VirtualLink, nestedVrfOspfAreaVirtualLink)
							}
						}
						nestedVrf.Ospf.Area = append(nestedVrf.Ospf.Area, nestedVrfOspfArea)
					}
				}
			}
			if oVrf.Ospfv3 != nil {
				nestedVrf.Ospfv3 = &VrfOspfv3Xml{}
				if _, ok := o.Misc["VrfOspfv3"]; ok {
					nestedVrf.Ospfv3.Misc = o.Misc["VrfOspfv3"]
				}
				if oVrf.Ospfv3.Enable != nil {
					nestedVrf.Ospfv3.Enable = util.YesNo(oVrf.Ospfv3.Enable, nil)
				}
				if oVrf.Ospfv3.RouterId != nil {
					nestedVrf.Ospfv3.RouterId = oVrf.Ospfv3.RouterId
				}
				if oVrf.Ospfv3.DisableTransitTraffic != nil {
					nestedVrf.Ospfv3.DisableTransitTraffic = util.YesNo(oVrf.Ospfv3.DisableTransitTraffic, nil)
				}
				if oVrf.Ospfv3.SpfTimer != nil {
					nestedVrf.Ospfv3.SpfTimer = oVrf.Ospfv3.SpfTimer
				}
				if oVrf.Ospfv3.GlobalIfTimer != nil {
					nestedVrf.Ospfv3.GlobalIfTimer = oVrf.Ospfv3.GlobalIfTimer
				}
				if oVrf.Ospfv3.RedistributionProfile != nil {
					nestedVrf.Ospfv3.RedistributionProfile = oVrf.Ospfv3.RedistributionProfile
				}
				if oVrf.Ospfv3.GlobalBfd != nil {
					nestedVrf.Ospfv3.GlobalBfd = &VrfOspfv3GlobalBfdXml{}
					if _, ok := o.Misc["VrfOspfv3GlobalBfd"]; ok {
						nestedVrf.Ospfv3.GlobalBfd.Misc = o.Misc["VrfOspfv3GlobalBfd"]
					}
					if oVrf.Ospfv3.GlobalBfd.Profile != nil {
						nestedVrf.Ospfv3.GlobalBfd.Profile = oVrf.Ospfv3.GlobalBfd.Profile
					}
				}
				if oVrf.Ospfv3.GracefulRestart != nil {
					nestedVrf.Ospfv3.GracefulRestart = &VrfOspfv3GracefulRestartXml{}
					if _, ok := o.Misc["VrfOspfv3GracefulRestart"]; ok {
						nestedVrf.Ospfv3.GracefulRestart.Misc = o.Misc["VrfOspfv3GracefulRestart"]
					}
					if oVrf.Ospfv3.GracefulRestart.Enable != nil {
						nestedVrf.Ospfv3.GracefulRestart.Enable = util.YesNo(oVrf.Ospfv3.GracefulRestart.Enable, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.GracePeriod != nil {
						nestedVrf.Ospfv3.GracefulRestart.GracePeriod = oVrf.Ospfv3.GracefulRestart.GracePeriod
					}
					if oVrf.Ospfv3.GracefulRestart.HelperEnable != nil {
						nestedVrf.Ospfv3.GracefulRestart.HelperEnable = util.YesNo(oVrf.Ospfv3.GracefulRestart.HelperEnable, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
						nestedVrf.Ospfv3.GracefulRestart.StrictLSAChecking = util.YesNo(oVrf.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime
					}
				}
				if oVrf.Ospfv3.Area != nil {
					nestedVrf.Ospfv3.Area = []VrfOspfv3AreaXml{}
					for _, oVrfOspfv3Area := range oVrf.Ospfv3.Area {
						nestedVrfOspfv3Area := VrfOspfv3AreaXml{}
						if _, ok := o.Misc["VrfOspfv3Area"]; ok {
							nestedVrfOspfv3Area.Misc = o.Misc["VrfOspfv3Area"]
						}
						if oVrfOspfv3Area.Name != "" {
							nestedVrfOspfv3Area.Name = oVrfOspfv3Area.Name
						}
						if oVrfOspfv3Area.Authentication != nil {
							nestedVrfOspfv3Area.Authentication = oVrfOspfv3Area.Authentication
						}
						if oVrfOspfv3Area.Type != nil {
							nestedVrfOspfv3Area.Type = &VrfOspfv3AreaTypeXml{}
							if _, ok := o.Misc["VrfOspfv3AreaType"]; ok {
								nestedVrfOspfv3Area.Type.Misc = o.Misc["VrfOspfv3AreaType"]
							}
							if oVrfOspfv3Area.Type.Normal != nil {
								nestedVrfOspfv3Area.Type.Normal = &VrfOspfv3AreaTypeNormalXml{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeNormal"]; ok {
									nestedVrfOspfv3Area.Type.Normal.Misc = o.Misc["VrfOspfv3AreaTypeNormal"]
								}
								if oVrfOspfv3Area.Type.Normal.Abr != nil {
									nestedVrfOspfv3Area.Type.Normal.Abr = &VrfOspfv3AreaTypeNormalAbrXml{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNormalAbr"]; ok {
										nestedVrfOspfv3Area.Type.Normal.Abr.Misc = o.Misc["VrfOspfv3AreaTypeNormalAbr"]
									}
									if oVrfOspfv3Area.Type.Normal.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.ImportList = oVrfOspfv3Area.Type.Normal.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.ExportList = oVrfOspfv3Area.Type.Normal.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.InboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList
									}
								}
							}
							if oVrfOspfv3Area.Type.Stub != nil {
								nestedVrfOspfv3Area.Type.Stub = &VrfOspfv3AreaTypeStubXml{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeStub"]; ok {
									nestedVrfOspfv3Area.Type.Stub.Misc = o.Misc["VrfOspfv3AreaTypeStub"]
								}
								if oVrfOspfv3Area.Type.Stub.NoSummary != nil {
									nestedVrfOspfv3Area.Type.Stub.NoSummary = util.YesNo(oVrfOspfv3Area.Type.Stub.NoSummary, nil)
								}
								if oVrfOspfv3Area.Type.Stub.Abr != nil {
									nestedVrfOspfv3Area.Type.Stub.Abr = &VrfOspfv3AreaTypeStubAbrXml{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeStubAbr"]; ok {
										nestedVrfOspfv3Area.Type.Stub.Abr.Misc = o.Misc["VrfOspfv3AreaTypeStubAbr"]
									}
									if oVrfOspfv3Area.Type.Stub.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.ImportList = oVrfOspfv3Area.Type.Stub.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.ExportList = oVrfOspfv3Area.Type.Stub.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.InboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList
									}
								}
								if oVrfOspfv3Area.Type.Stub.DefaultRouteMetric != nil {
									nestedVrfOspfv3Area.Type.Stub.DefaultRouteMetric = oVrfOspfv3Area.Type.Stub.DefaultRouteMetric
								}
							}
							if oVrfOspfv3Area.Type.Nssa != nil {
								nestedVrfOspfv3Area.Type.Nssa = &VrfOspfv3AreaTypeNssaXml{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeNssa"]; ok {
									nestedVrfOspfv3Area.Type.Nssa.Misc = o.Misc["VrfOspfv3AreaTypeNssa"]
								}
								if oVrfOspfv3Area.Type.Nssa.NoSummary != nil {
									nestedVrfOspfv3Area.Type.Nssa.NoSummary = util.YesNo(oVrfOspfv3Area.Type.Nssa.NoSummary, nil)
								}
								if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate != nil {
									nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate = &VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"]; ok {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc = o.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"]
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType
									}
								}
								if oVrfOspfv3Area.Type.Nssa.Abr != nil {
									nestedVrfOspfv3Area.Type.Nssa.Abr = &VrfOspfv3AreaTypeNssaAbrXml{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNssaAbr"]; ok {
										nestedVrfOspfv3Area.Type.Nssa.Abr.Misc = o.Misc["VrfOspfv3AreaTypeNssaAbr"]
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.ImportList = oVrfOspfv3Area.Type.Nssa.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.ExportList = oVrfOspfv3Area.Type.Nssa.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = []VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml{}
										for _, oVrfOspfv3AreaTypeNssaAbrNssaExtRange := range oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange {
											nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange := VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml{}
											if _, ok := o.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"]; ok {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc = o.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"]
											}
											if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name != "" {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name
											}
											if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise != nil {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise = util.YesNo(oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise, nil)
											}
											nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange)
										}
									}
								}
							}
						}
						if oVrfOspfv3Area.Range != nil {
							nestedVrfOspfv3Area.Range = []VrfOspfv3AreaRangeXml{}
							for _, oVrfOspfv3AreaRange := range oVrfOspfv3Area.Range {
								nestedVrfOspfv3AreaRange := VrfOspfv3AreaRangeXml{}
								if _, ok := o.Misc["VrfOspfv3AreaRange"]; ok {
									nestedVrfOspfv3AreaRange.Misc = o.Misc["VrfOspfv3AreaRange"]
								}
								if oVrfOspfv3AreaRange.Name != "" {
									nestedVrfOspfv3AreaRange.Name = oVrfOspfv3AreaRange.Name
								}
								if oVrfOspfv3AreaRange.Advertise != nil {
									nestedVrfOspfv3AreaRange.Advertise = util.YesNo(oVrfOspfv3AreaRange.Advertise, nil)
								}
								nestedVrfOspfv3Area.Range = append(nestedVrfOspfv3Area.Range, nestedVrfOspfv3AreaRange)
							}
						}
						if oVrfOspfv3Area.Interface != nil {
							nestedVrfOspfv3Area.Interface = []VrfOspfv3AreaInterfaceXml{}
							for _, oVrfOspfv3AreaInterface := range oVrfOspfv3Area.Interface {
								nestedVrfOspfv3AreaInterface := VrfOspfv3AreaInterfaceXml{}
								if _, ok := o.Misc["VrfOspfv3AreaInterface"]; ok {
									nestedVrfOspfv3AreaInterface.Misc = o.Misc["VrfOspfv3AreaInterface"]
								}
								if oVrfOspfv3AreaInterface.Name != "" {
									nestedVrfOspfv3AreaInterface.Name = oVrfOspfv3AreaInterface.Name
								}
								if oVrfOspfv3AreaInterface.Enable != nil {
									nestedVrfOspfv3AreaInterface.Enable = util.YesNo(oVrfOspfv3AreaInterface.Enable, nil)
								}
								if oVrfOspfv3AreaInterface.MtuIgnore != nil {
									nestedVrfOspfv3AreaInterface.MtuIgnore = util.YesNo(oVrfOspfv3AreaInterface.MtuIgnore, nil)
								}
								if oVrfOspfv3AreaInterface.Passive != nil {
									nestedVrfOspfv3AreaInterface.Passive = util.YesNo(oVrfOspfv3AreaInterface.Passive, nil)
								}
								if oVrfOspfv3AreaInterface.Priority != nil {
									nestedVrfOspfv3AreaInterface.Priority = oVrfOspfv3AreaInterface.Priority
								}
								if oVrfOspfv3AreaInterface.Metric != nil {
									nestedVrfOspfv3AreaInterface.Metric = oVrfOspfv3AreaInterface.Metric
								}
								if oVrfOspfv3AreaInterface.InstanceId != nil {
									nestedVrfOspfv3AreaInterface.InstanceId = oVrfOspfv3AreaInterface.InstanceId
								}
								if oVrfOspfv3AreaInterface.Authentication != nil {
									nestedVrfOspfv3AreaInterface.Authentication = oVrfOspfv3AreaInterface.Authentication
								}
								if oVrfOspfv3AreaInterface.Timing != nil {
									nestedVrfOspfv3AreaInterface.Timing = oVrfOspfv3AreaInterface.Timing
								}
								if oVrfOspfv3AreaInterface.LinkType != nil {
									nestedVrfOspfv3AreaInterface.LinkType = &VrfOspfv3AreaInterfaceLinkTypeXml{}
									if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkType"]; ok {
										nestedVrfOspfv3AreaInterface.LinkType.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkType"]
									}
									if oVrfOspfv3AreaInterface.LinkType.Broadcast != nil {
										nestedVrfOspfv3AreaInterface.LinkType.Broadcast = &VrfOspfv3AreaInterfaceLinkTypeBroadcastXml{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.Broadcast.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"]
										}
									}
									if oVrfOspfv3AreaInterface.LinkType.P2p != nil {
										nestedVrfOspfv3AreaInterface.LinkType.P2p = &VrfOspfv3AreaInterfaceLinkTypeP2pXml{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.P2p.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"]
										}
									}
									if oVrfOspfv3AreaInterface.LinkType.P2mp != nil {
										nestedVrfOspfv3AreaInterface.LinkType.P2mp = &VrfOspfv3AreaInterfaceLinkTypeP2mpXml{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"]
										}
										if oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml{}
											for _, oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor {
												nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml{}
												if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"]; ok {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"]
												}
												if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name
												}
												if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority
												}
												nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor)
											}
										}
									}
								}
								if oVrfOspfv3AreaInterface.Bfd != nil {
									nestedVrfOspfv3AreaInterface.Bfd = &VrfOspfv3AreaInterfaceBfdXml{}
									if _, ok := o.Misc["VrfOspfv3AreaInterfaceBfd"]; ok {
										nestedVrfOspfv3AreaInterface.Bfd.Misc = o.Misc["VrfOspfv3AreaInterfaceBfd"]
									}
									if oVrfOspfv3AreaInterface.Bfd.Profile != nil {
										nestedVrfOspfv3AreaInterface.Bfd.Profile = oVrfOspfv3AreaInterface.Bfd.Profile
									}
								}
								nestedVrfOspfv3Area.Interface = append(nestedVrfOspfv3Area.Interface, nestedVrfOspfv3AreaInterface)
							}
						}
						if oVrfOspfv3Area.VirtualLink != nil {
							nestedVrfOspfv3Area.VirtualLink = []VrfOspfv3AreaVirtualLinkXml{}
							for _, oVrfOspfv3AreaVirtualLink := range oVrfOspfv3Area.VirtualLink {
								nestedVrfOspfv3AreaVirtualLink := VrfOspfv3AreaVirtualLinkXml{}
								if _, ok := o.Misc["VrfOspfv3AreaVirtualLink"]; ok {
									nestedVrfOspfv3AreaVirtualLink.Misc = o.Misc["VrfOspfv3AreaVirtualLink"]
								}
								if oVrfOspfv3AreaVirtualLink.Name != "" {
									nestedVrfOspfv3AreaVirtualLink.Name = oVrfOspfv3AreaVirtualLink.Name
								}
								if oVrfOspfv3AreaVirtualLink.NeighborId != nil {
									nestedVrfOspfv3AreaVirtualLink.NeighborId = oVrfOspfv3AreaVirtualLink.NeighborId
								}
								if oVrfOspfv3AreaVirtualLink.TransitAreaId != nil {
									nestedVrfOspfv3AreaVirtualLink.TransitAreaId = oVrfOspfv3AreaVirtualLink.TransitAreaId
								}
								if oVrfOspfv3AreaVirtualLink.Enable != nil {
									nestedVrfOspfv3AreaVirtualLink.Enable = util.YesNo(oVrfOspfv3AreaVirtualLink.Enable, nil)
								}
								if oVrfOspfv3AreaVirtualLink.InstanceId != nil {
									nestedVrfOspfv3AreaVirtualLink.InstanceId = oVrfOspfv3AreaVirtualLink.InstanceId
								}
								if oVrfOspfv3AreaVirtualLink.Timing != nil {
									nestedVrfOspfv3AreaVirtualLink.Timing = oVrfOspfv3AreaVirtualLink.Timing
								}
								if oVrfOspfv3AreaVirtualLink.Authentication != nil {
									nestedVrfOspfv3AreaVirtualLink.Authentication = oVrfOspfv3AreaVirtualLink.Authentication
								}
								nestedVrfOspfv3Area.VirtualLink = append(nestedVrfOspfv3Area.VirtualLink, nestedVrfOspfv3AreaVirtualLink)
							}
						}
						nestedVrf.Ospfv3.Area = append(nestedVrf.Ospfv3.Area, nestedVrfOspfv3Area)
					}
				}
			}
			if oVrf.Ecmp != nil {
				nestedVrf.Ecmp = &VrfEcmpXml{}
				if _, ok := o.Misc["VrfEcmp"]; ok {
					nestedVrf.Ecmp.Misc = o.Misc["VrfEcmp"]
				}
				if oVrf.Ecmp.Enable != nil {
					nestedVrf.Ecmp.Enable = util.YesNo(oVrf.Ecmp.Enable, nil)
				}
				if oVrf.Ecmp.MaxPath != nil {
					nestedVrf.Ecmp.MaxPath = oVrf.Ecmp.MaxPath
				}
				if oVrf.Ecmp.SymmetricReturn != nil {
					nestedVrf.Ecmp.SymmetricReturn = util.YesNo(oVrf.Ecmp.SymmetricReturn, nil)
				}
				if oVrf.Ecmp.StrictSourcePath != nil {
					nestedVrf.Ecmp.StrictSourcePath = util.YesNo(oVrf.Ecmp.StrictSourcePath, nil)
				}
				if oVrf.Ecmp.Algorithm != nil {
					nestedVrf.Ecmp.Algorithm = &VrfEcmpAlgorithmXml{}
					if _, ok := o.Misc["VrfEcmpAlgorithm"]; ok {
						nestedVrf.Ecmp.Algorithm.Misc = o.Misc["VrfEcmpAlgorithm"]
					}
					if oVrf.Ecmp.Algorithm.IpModulo != nil {
						nestedVrf.Ecmp.Algorithm.IpModulo = &VrfEcmpAlgorithmIpModuloXml{}
						if _, ok := o.Misc["VrfEcmpAlgorithmIpModulo"]; ok {
							nestedVrf.Ecmp.Algorithm.IpModulo.Misc = o.Misc["VrfEcmpAlgorithmIpModulo"]
						}
					}
					if oVrf.Ecmp.Algorithm.IpHash != nil {
						nestedVrf.Ecmp.Algorithm.IpHash = &VrfEcmpAlgorithmIpHashXml{}
						if _, ok := o.Misc["VrfEcmpAlgorithmIpHash"]; ok {
							nestedVrf.Ecmp.Algorithm.IpHash.Misc = o.Misc["VrfEcmpAlgorithmIpHash"]
						}
						if oVrf.Ecmp.Algorithm.IpHash.SrcOnly != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.SrcOnly = util.YesNo(oVrf.Ecmp.Algorithm.IpHash.SrcOnly, nil)
						}
						if oVrf.Ecmp.Algorithm.IpHash.UsePort != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.UsePort = util.YesNo(oVrf.Ecmp.Algorithm.IpHash.UsePort, nil)
						}
						if oVrf.Ecmp.Algorithm.IpHash.HashSeed != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.HashSeed = oVrf.Ecmp.Algorithm.IpHash.HashSeed
						}
					}
					if oVrf.Ecmp.Algorithm.WeightedRoundRobin != nil {
						nestedVrf.Ecmp.Algorithm.WeightedRoundRobin = &VrfEcmpAlgorithmWeightedRoundRobinXml{}
						if _, ok := o.Misc["VrfEcmpAlgorithmWeightedRoundRobin"]; ok {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc = o.Misc["VrfEcmpAlgorithmWeightedRoundRobin"]
						}
						if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = []VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml{}
							for _, oVrfEcmpAlgorithmWeightedRoundRobinInterface := range oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface {
								nestedVrfEcmpAlgorithmWeightedRoundRobinInterface := VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml{}
								if _, ok := o.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"]; ok {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc = o.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"]
								}
								if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Name = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name
								}
								if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight
								}
								nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = append(nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface, nestedVrfEcmpAlgorithmWeightedRoundRobinInterface)
							}
						}
					}
					if oVrf.Ecmp.Algorithm.BalancedRoundRobin != nil {
						nestedVrf.Ecmp.Algorithm.BalancedRoundRobin = &VrfEcmpAlgorithmBalancedRoundRobinXml{}
						if _, ok := o.Misc["VrfEcmpAlgorithmBalancedRoundRobin"]; ok {
							nestedVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc = o.Misc["VrfEcmpAlgorithmBalancedRoundRobin"]
						}
					}
				}
			}
			if oVrf.Multicast != nil {
				nestedVrf.Multicast = &VrfMulticastXml{}
				if _, ok := o.Misc["VrfMulticast"]; ok {
					nestedVrf.Multicast.Misc = o.Misc["VrfMulticast"]
				}
				if oVrf.Multicast.Enable != nil {
					nestedVrf.Multicast.Enable = util.YesNo(oVrf.Multicast.Enable, nil)
				}
				if oVrf.Multicast.StaticRoute != nil {
					nestedVrf.Multicast.StaticRoute = []VrfMulticastStaticRouteXml{}
					for _, oVrfMulticastStaticRoute := range oVrf.Multicast.StaticRoute {
						nestedVrfMulticastStaticRoute := VrfMulticastStaticRouteXml{}
						if _, ok := o.Misc["VrfMulticastStaticRoute"]; ok {
							nestedVrfMulticastStaticRoute.Misc = o.Misc["VrfMulticastStaticRoute"]
						}
						if oVrfMulticastStaticRoute.Name != "" {
							nestedVrfMulticastStaticRoute.Name = oVrfMulticastStaticRoute.Name
						}
						if oVrfMulticastStaticRoute.Destination != nil {
							nestedVrfMulticastStaticRoute.Destination = oVrfMulticastStaticRoute.Destination
						}
						if oVrfMulticastStaticRoute.Interface != nil {
							nestedVrfMulticastStaticRoute.Interface = oVrfMulticastStaticRoute.Interface
						}
						if oVrfMulticastStaticRoute.Preference != nil {
							nestedVrfMulticastStaticRoute.Preference = oVrfMulticastStaticRoute.Preference
						}
						if oVrfMulticastStaticRoute.Nexthop != nil {
							nestedVrfMulticastStaticRoute.Nexthop = &VrfMulticastStaticRouteNexthopXml{}
							if _, ok := o.Misc["VrfMulticastStaticRouteNexthop"]; ok {
								nestedVrfMulticastStaticRoute.Nexthop.Misc = o.Misc["VrfMulticastStaticRouteNexthop"]
							}
							if oVrfMulticastStaticRoute.Nexthop.IpAddress != nil {
								nestedVrfMulticastStaticRoute.Nexthop.IpAddress = oVrfMulticastStaticRoute.Nexthop.IpAddress
							}
						}
						nestedVrf.Multicast.StaticRoute = append(nestedVrf.Multicast.StaticRoute, nestedVrfMulticastStaticRoute)
					}
				}
				if oVrf.Multicast.Pim != nil {
					nestedVrf.Multicast.Pim = &VrfMulticastPimXml{}
					if _, ok := o.Misc["VrfMulticastPim"]; ok {
						nestedVrf.Multicast.Pim.Misc = o.Misc["VrfMulticastPim"]
					}
					if oVrf.Multicast.Pim.Enable != nil {
						nestedVrf.Multicast.Pim.Enable = util.YesNo(oVrf.Multicast.Pim.Enable, nil)
					}
					if oVrf.Multicast.Pim.RpfLookupMode != nil {
						nestedVrf.Multicast.Pim.RpfLookupMode = oVrf.Multicast.Pim.RpfLookupMode
					}
					if oVrf.Multicast.Pim.RouteAgeoutTime != nil {
						nestedVrf.Multicast.Pim.RouteAgeoutTime = oVrf.Multicast.Pim.RouteAgeoutTime
					}
					if oVrf.Multicast.Pim.IfTimerGlobal != nil {
						nestedVrf.Multicast.Pim.IfTimerGlobal = oVrf.Multicast.Pim.IfTimerGlobal
					}
					if oVrf.Multicast.Pim.GroupPermission != nil {
						nestedVrf.Multicast.Pim.GroupPermission = oVrf.Multicast.Pim.GroupPermission
					}
					if oVrf.Multicast.Pim.SsmAddressSpace != nil {
						nestedVrf.Multicast.Pim.SsmAddressSpace = &VrfMulticastPimSsmAddressSpaceXml{}
						if _, ok := o.Misc["VrfMulticastPimSsmAddressSpace"]; ok {
							nestedVrf.Multicast.Pim.SsmAddressSpace.Misc = o.Misc["VrfMulticastPimSsmAddressSpace"]
						}
						if oVrf.Multicast.Pim.SsmAddressSpace.GroupList != nil {
							nestedVrf.Multicast.Pim.SsmAddressSpace.GroupList = oVrf.Multicast.Pim.SsmAddressSpace.GroupList
						}
					}
					if oVrf.Multicast.Pim.Rp != nil {
						nestedVrf.Multicast.Pim.Rp = &VrfMulticastPimRpXml{}
						if _, ok := o.Misc["VrfMulticastPimRp"]; ok {
							nestedVrf.Multicast.Pim.Rp.Misc = o.Misc["VrfMulticastPimRp"]
						}
						if oVrf.Multicast.Pim.Rp.LocalRp != nil {
							nestedVrf.Multicast.Pim.Rp.LocalRp = &VrfMulticastPimRpLocalRpXml{}
							if _, ok := o.Misc["VrfMulticastPimRpLocalRp"]; ok {
								nestedVrf.Multicast.Pim.Rp.LocalRp.Misc = o.Misc["VrfMulticastPimRpLocalRp"]
							}
							if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp = &VrfMulticastPimRpLocalRpStaticRpXml{}
								if _, ok := o.Misc["VrfMulticastPimRpLocalRpStaticRp"]; ok {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc = o.Misc["VrfMulticastPimRpLocalRpStaticRp"]
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override = util.YesNo(oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override, nil)
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList
								}
							}
							if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp = &VrfMulticastPimRpLocalRpCandidateRpXml{}
								if _, ok := o.Misc["VrfMulticastPimRpLocalRpCandidateRp"]; ok {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc = o.Misc["VrfMulticastPimRpLocalRpCandidateRp"]
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList
								}
							}
						}
						if oVrf.Multicast.Pim.Rp.ExternalRp != nil {
							nestedVrf.Multicast.Pim.Rp.ExternalRp = []VrfMulticastPimRpExternalRpXml{}
							for _, oVrfMulticastPimRpExternalRp := range oVrf.Multicast.Pim.Rp.ExternalRp {
								nestedVrfMulticastPimRpExternalRp := VrfMulticastPimRpExternalRpXml{}
								if _, ok := o.Misc["VrfMulticastPimRpExternalRp"]; ok {
									nestedVrfMulticastPimRpExternalRp.Misc = o.Misc["VrfMulticastPimRpExternalRp"]
								}
								if oVrfMulticastPimRpExternalRp.Name != "" {
									nestedVrfMulticastPimRpExternalRp.Name = oVrfMulticastPimRpExternalRp.Name
								}
								if oVrfMulticastPimRpExternalRp.GroupList != nil {
									nestedVrfMulticastPimRpExternalRp.GroupList = oVrfMulticastPimRpExternalRp.GroupList
								}
								if oVrfMulticastPimRpExternalRp.Override != nil {
									nestedVrfMulticastPimRpExternalRp.Override = util.YesNo(oVrfMulticastPimRpExternalRp.Override, nil)
								}
								nestedVrf.Multicast.Pim.Rp.ExternalRp = append(nestedVrf.Multicast.Pim.Rp.ExternalRp, nestedVrfMulticastPimRpExternalRp)
							}
						}
					}
					if oVrf.Multicast.Pim.SptThreshold != nil {
						nestedVrf.Multicast.Pim.SptThreshold = []VrfMulticastPimSptThresholdXml{}
						for _, oVrfMulticastPimSptThreshold := range oVrf.Multicast.Pim.SptThreshold {
							nestedVrfMulticastPimSptThreshold := VrfMulticastPimSptThresholdXml{}
							if _, ok := o.Misc["VrfMulticastPimSptThreshold"]; ok {
								nestedVrfMulticastPimSptThreshold.Misc = o.Misc["VrfMulticastPimSptThreshold"]
							}
							if oVrfMulticastPimSptThreshold.Name != "" {
								nestedVrfMulticastPimSptThreshold.Name = oVrfMulticastPimSptThreshold.Name
							}
							if oVrfMulticastPimSptThreshold.Threshold != nil {
								nestedVrfMulticastPimSptThreshold.Threshold = oVrfMulticastPimSptThreshold.Threshold
							}
							nestedVrf.Multicast.Pim.SptThreshold = append(nestedVrf.Multicast.Pim.SptThreshold, nestedVrfMulticastPimSptThreshold)
						}
					}
					if oVrf.Multicast.Pim.Interface != nil {
						nestedVrf.Multicast.Pim.Interface = []VrfMulticastPimInterfaceXml{}
						for _, oVrfMulticastPimInterface := range oVrf.Multicast.Pim.Interface {
							nestedVrfMulticastPimInterface := VrfMulticastPimInterfaceXml{}
							if _, ok := o.Misc["VrfMulticastPimInterface"]; ok {
								nestedVrfMulticastPimInterface.Misc = o.Misc["VrfMulticastPimInterface"]
							}
							if oVrfMulticastPimInterface.Name != "" {
								nestedVrfMulticastPimInterface.Name = oVrfMulticastPimInterface.Name
							}
							if oVrfMulticastPimInterface.Description != nil {
								nestedVrfMulticastPimInterface.Description = oVrfMulticastPimInterface.Description
							}
							if oVrfMulticastPimInterface.DrPriority != nil {
								nestedVrfMulticastPimInterface.DrPriority = oVrfMulticastPimInterface.DrPriority
							}
							if oVrfMulticastPimInterface.SendBsm != nil {
								nestedVrfMulticastPimInterface.SendBsm = util.YesNo(oVrfMulticastPimInterface.SendBsm, nil)
							}
							if oVrfMulticastPimInterface.IfTimer != nil {
								nestedVrfMulticastPimInterface.IfTimer = oVrfMulticastPimInterface.IfTimer
							}
							if oVrfMulticastPimInterface.NeighborFilter != nil {
								nestedVrfMulticastPimInterface.NeighborFilter = oVrfMulticastPimInterface.NeighborFilter
							}
							nestedVrf.Multicast.Pim.Interface = append(nestedVrf.Multicast.Pim.Interface, nestedVrfMulticastPimInterface)
						}
					}
				}
				if oVrf.Multicast.Igmp != nil {
					nestedVrf.Multicast.Igmp = &VrfMulticastIgmpXml{}
					if _, ok := o.Misc["VrfMulticastIgmp"]; ok {
						nestedVrf.Multicast.Igmp.Misc = o.Misc["VrfMulticastIgmp"]
					}
					if oVrf.Multicast.Igmp.Enable != nil {
						nestedVrf.Multicast.Igmp.Enable = util.YesNo(oVrf.Multicast.Igmp.Enable, nil)
					}
					if oVrf.Multicast.Igmp.Dynamic != nil {
						nestedVrf.Multicast.Igmp.Dynamic = &VrfMulticastIgmpDynamicXml{}
						if _, ok := o.Misc["VrfMulticastIgmpDynamic"]; ok {
							nestedVrf.Multicast.Igmp.Dynamic.Misc = o.Misc["VrfMulticastIgmpDynamic"]
						}
						if oVrf.Multicast.Igmp.Dynamic.Interface != nil {
							nestedVrf.Multicast.Igmp.Dynamic.Interface = []VrfMulticastIgmpDynamicInterfaceXml{}
							for _, oVrfMulticastIgmpDynamicInterface := range oVrf.Multicast.Igmp.Dynamic.Interface {
								nestedVrfMulticastIgmpDynamicInterface := VrfMulticastIgmpDynamicInterfaceXml{}
								if _, ok := o.Misc["VrfMulticastIgmpDynamicInterface"]; ok {
									nestedVrfMulticastIgmpDynamicInterface.Misc = o.Misc["VrfMulticastIgmpDynamicInterface"]
								}
								if oVrfMulticastIgmpDynamicInterface.Name != "" {
									nestedVrfMulticastIgmpDynamicInterface.Name = oVrfMulticastIgmpDynamicInterface.Name
								}
								if oVrfMulticastIgmpDynamicInterface.Version != nil {
									nestedVrfMulticastIgmpDynamicInterface.Version = oVrfMulticastIgmpDynamicInterface.Version
								}
								if oVrfMulticastIgmpDynamicInterface.Robustness != nil {
									nestedVrfMulticastIgmpDynamicInterface.Robustness = oVrfMulticastIgmpDynamicInterface.Robustness
								}
								if oVrfMulticastIgmpDynamicInterface.GroupFilter != nil {
									nestedVrfMulticastIgmpDynamicInterface.GroupFilter = oVrfMulticastIgmpDynamicInterface.GroupFilter
								}
								if oVrfMulticastIgmpDynamicInterface.MaxGroups != nil {
									nestedVrfMulticastIgmpDynamicInterface.MaxGroups = oVrfMulticastIgmpDynamicInterface.MaxGroups
								}
								if oVrfMulticastIgmpDynamicInterface.MaxSources != nil {
									nestedVrfMulticastIgmpDynamicInterface.MaxSources = oVrfMulticastIgmpDynamicInterface.MaxSources
								}
								if oVrfMulticastIgmpDynamicInterface.QueryProfile != nil {
									nestedVrfMulticastIgmpDynamicInterface.QueryProfile = oVrfMulticastIgmpDynamicInterface.QueryProfile
								}
								if oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing != nil {
									nestedVrfMulticastIgmpDynamicInterface.RouterAlertPolicing = util.YesNo(oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing, nil)
								}
								nestedVrf.Multicast.Igmp.Dynamic.Interface = append(nestedVrf.Multicast.Igmp.Dynamic.Interface, nestedVrfMulticastIgmpDynamicInterface)
							}
						}
					}
					if oVrf.Multicast.Igmp.Static != nil {
						nestedVrf.Multicast.Igmp.Static = []VrfMulticastIgmpStaticXml{}
						for _, oVrfMulticastIgmpStatic := range oVrf.Multicast.Igmp.Static {
							nestedVrfMulticastIgmpStatic := VrfMulticastIgmpStaticXml{}
							if _, ok := o.Misc["VrfMulticastIgmpStatic"]; ok {
								nestedVrfMulticastIgmpStatic.Misc = o.Misc["VrfMulticastIgmpStatic"]
							}
							if oVrfMulticastIgmpStatic.Name != "" {
								nestedVrfMulticastIgmpStatic.Name = oVrfMulticastIgmpStatic.Name
							}
							if oVrfMulticastIgmpStatic.Interface != nil {
								nestedVrfMulticastIgmpStatic.Interface = oVrfMulticastIgmpStatic.Interface
							}
							if oVrfMulticastIgmpStatic.GroupAddress != nil {
								nestedVrfMulticastIgmpStatic.GroupAddress = oVrfMulticastIgmpStatic.GroupAddress
							}
							if oVrfMulticastIgmpStatic.SourceAddress != nil {
								nestedVrfMulticastIgmpStatic.SourceAddress = oVrfMulticastIgmpStatic.SourceAddress
							}
							nestedVrf.Multicast.Igmp.Static = append(nestedVrf.Multicast.Igmp.Static, nestedVrfMulticastIgmpStatic)
						}
					}
				}
				if oVrf.Multicast.Msdp != nil {
					nestedVrf.Multicast.Msdp = &VrfMulticastMsdpXml{}
					if _, ok := o.Misc["VrfMulticastMsdp"]; ok {
						nestedVrf.Multicast.Msdp.Misc = o.Misc["VrfMulticastMsdp"]
					}
					if oVrf.Multicast.Msdp.Enable != nil {
						nestedVrf.Multicast.Msdp.Enable = util.YesNo(oVrf.Multicast.Msdp.Enable, nil)
					}
					if oVrf.Multicast.Msdp.GlobalTimer != nil {
						nestedVrf.Multicast.Msdp.GlobalTimer = oVrf.Multicast.Msdp.GlobalTimer
					}
					if oVrf.Multicast.Msdp.GlobalAuthentication != nil {
						nestedVrf.Multicast.Msdp.GlobalAuthentication = oVrf.Multicast.Msdp.GlobalAuthentication
					}
					if oVrf.Multicast.Msdp.OriginatorId != nil {
						nestedVrf.Multicast.Msdp.OriginatorId = &VrfMulticastMsdpOriginatorIdXml{}
						if _, ok := o.Misc["VrfMulticastMsdpOriginatorId"]; ok {
							nestedVrf.Multicast.Msdp.OriginatorId.Misc = o.Misc["VrfMulticastMsdpOriginatorId"]
						}
						if oVrf.Multicast.Msdp.OriginatorId.Interface != nil {
							nestedVrf.Multicast.Msdp.OriginatorId.Interface = oVrf.Multicast.Msdp.OriginatorId.Interface
						}
						if oVrf.Multicast.Msdp.OriginatorId.Ip != nil {
							nestedVrf.Multicast.Msdp.OriginatorId.Ip = oVrf.Multicast.Msdp.OriginatorId.Ip
						}
					}
					if oVrf.Multicast.Msdp.Peer != nil {
						nestedVrf.Multicast.Msdp.Peer = []VrfMulticastMsdpPeerXml{}
						for _, oVrfMulticastMsdpPeer := range oVrf.Multicast.Msdp.Peer {
							nestedVrfMulticastMsdpPeer := VrfMulticastMsdpPeerXml{}
							if _, ok := o.Misc["VrfMulticastMsdpPeer"]; ok {
								nestedVrfMulticastMsdpPeer.Misc = o.Misc["VrfMulticastMsdpPeer"]
							}
							if oVrfMulticastMsdpPeer.Name != "" {
								nestedVrfMulticastMsdpPeer.Name = oVrfMulticastMsdpPeer.Name
							}
							if oVrfMulticastMsdpPeer.Enable != nil {
								nestedVrfMulticastMsdpPeer.Enable = util.YesNo(oVrfMulticastMsdpPeer.Enable, nil)
							}
							if oVrfMulticastMsdpPeer.PeerAs != nil {
								nestedVrfMulticastMsdpPeer.PeerAs = oVrfMulticastMsdpPeer.PeerAs
							}
							if oVrfMulticastMsdpPeer.Authentication != nil {
								nestedVrfMulticastMsdpPeer.Authentication = oVrfMulticastMsdpPeer.Authentication
							}
							if oVrfMulticastMsdpPeer.MaxSa != nil {
								nestedVrfMulticastMsdpPeer.MaxSa = oVrfMulticastMsdpPeer.MaxSa
							}
							if oVrfMulticastMsdpPeer.InboundSaFilter != nil {
								nestedVrfMulticastMsdpPeer.InboundSaFilter = oVrfMulticastMsdpPeer.InboundSaFilter
							}
							if oVrfMulticastMsdpPeer.OutboundSaFilter != nil {
								nestedVrfMulticastMsdpPeer.OutboundSaFilter = oVrfMulticastMsdpPeer.OutboundSaFilter
							}
							if oVrfMulticastMsdpPeer.LocalAddress != nil {
								nestedVrfMulticastMsdpPeer.LocalAddress = &VrfMulticastMsdpPeerLocalAddressXml{}
								if _, ok := o.Misc["VrfMulticastMsdpPeerLocalAddress"]; ok {
									nestedVrfMulticastMsdpPeer.LocalAddress.Misc = o.Misc["VrfMulticastMsdpPeerLocalAddress"]
								}
								if oVrfMulticastMsdpPeer.LocalAddress.Interface != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress.Interface = oVrfMulticastMsdpPeer.LocalAddress.Interface
								}
								if oVrfMulticastMsdpPeer.LocalAddress.Ip != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress.Ip = oVrfMulticastMsdpPeer.LocalAddress.Ip
								}
							}
							if oVrfMulticastMsdpPeer.PeerAddress != nil {
								nestedVrfMulticastMsdpPeer.PeerAddress = &VrfMulticastMsdpPeerPeerAddressXml{}
								if _, ok := o.Misc["VrfMulticastMsdpPeerPeerAddress"]; ok {
									nestedVrfMulticastMsdpPeer.PeerAddress.Misc = o.Misc["VrfMulticastMsdpPeerPeerAddress"]
								}
								if oVrfMulticastMsdpPeer.PeerAddress.Ip != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress.Ip = oVrfMulticastMsdpPeer.PeerAddress.Ip
								}
								if oVrfMulticastMsdpPeer.PeerAddress.Fqdn != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress.Fqdn = oVrfMulticastMsdpPeer.PeerAddress.Fqdn
								}
							}
							nestedVrf.Multicast.Msdp.Peer = append(nestedVrf.Multicast.Msdp.Peer, nestedVrfMulticastMsdpPeer)
						}
					}
				}
			}
			if oVrf.Rip != nil {
				nestedVrf.Rip = &VrfRipXml{}
				if _, ok := o.Misc["VrfRip"]; ok {
					nestedVrf.Rip.Misc = o.Misc["VrfRip"]
				}
				if oVrf.Rip.Enable != nil {
					nestedVrf.Rip.Enable = util.YesNo(oVrf.Rip.Enable, nil)
				}
				if oVrf.Rip.DefaultInformationOriginate != nil {
					nestedVrf.Rip.DefaultInformationOriginate = util.YesNo(oVrf.Rip.DefaultInformationOriginate, nil)
				}
				if oVrf.Rip.GlobalTimer != nil {
					nestedVrf.Rip.GlobalTimer = oVrf.Rip.GlobalTimer
				}
				if oVrf.Rip.AuthProfile != nil {
					nestedVrf.Rip.AuthProfile = oVrf.Rip.AuthProfile
				}
				if oVrf.Rip.RedistributionProfile != nil {
					nestedVrf.Rip.RedistributionProfile = oVrf.Rip.RedistributionProfile
				}
				if oVrf.Rip.GlobalBfd != nil {
					nestedVrf.Rip.GlobalBfd = &VrfRipGlobalBfdXml{}
					if _, ok := o.Misc["VrfRipGlobalBfd"]; ok {
						nestedVrf.Rip.GlobalBfd.Misc = o.Misc["VrfRipGlobalBfd"]
					}
					if oVrf.Rip.GlobalBfd.Profile != nil {
						nestedVrf.Rip.GlobalBfd.Profile = oVrf.Rip.GlobalBfd.Profile
					}
				}
				if oVrf.Rip.GlobalInboundDistributeList != nil {
					nestedVrf.Rip.GlobalInboundDistributeList = &VrfRipGlobalInboundDistributeListXml{}
					if _, ok := o.Misc["VrfRipGlobalInboundDistributeList"]; ok {
						nestedVrf.Rip.GlobalInboundDistributeList.Misc = o.Misc["VrfRipGlobalInboundDistributeList"]
					}
					if oVrf.Rip.GlobalInboundDistributeList.AccessList != nil {
						nestedVrf.Rip.GlobalInboundDistributeList.AccessList = oVrf.Rip.GlobalInboundDistributeList.AccessList
					}
				}
				if oVrf.Rip.GlobalOutboundDistributeList != nil {
					nestedVrf.Rip.GlobalOutboundDistributeList = &VrfRipGlobalOutboundDistributeListXml{}
					if _, ok := o.Misc["VrfRipGlobalOutboundDistributeList"]; ok {
						nestedVrf.Rip.GlobalOutboundDistributeList.Misc = o.Misc["VrfRipGlobalOutboundDistributeList"]
					}
					if oVrf.Rip.GlobalOutboundDistributeList.AccessList != nil {
						nestedVrf.Rip.GlobalOutboundDistributeList.AccessList = oVrf.Rip.GlobalOutboundDistributeList.AccessList
					}
				}
				if oVrf.Rip.Interface != nil {
					nestedVrf.Rip.Interface = []VrfRipInterfaceXml{}
					for _, oVrfRipInterface := range oVrf.Rip.Interface {
						nestedVrfRipInterface := VrfRipInterfaceXml{}
						if _, ok := o.Misc["VrfRipInterface"]; ok {
							nestedVrfRipInterface.Misc = o.Misc["VrfRipInterface"]
						}
						if oVrfRipInterface.Name != "" {
							nestedVrfRipInterface.Name = oVrfRipInterface.Name
						}
						if oVrfRipInterface.Enable != nil {
							nestedVrfRipInterface.Enable = util.YesNo(oVrfRipInterface.Enable, nil)
						}
						if oVrfRipInterface.Mode != nil {
							nestedVrfRipInterface.Mode = oVrfRipInterface.Mode
						}
						if oVrfRipInterface.SplitHorizon != nil {
							nestedVrfRipInterface.SplitHorizon = oVrfRipInterface.SplitHorizon
						}
						if oVrfRipInterface.Authentication != nil {
							nestedVrfRipInterface.Authentication = oVrfRipInterface.Authentication
						}
						if oVrfRipInterface.Bfd != nil {
							nestedVrfRipInterface.Bfd = &VrfRipInterfaceBfdXml{}
							if _, ok := o.Misc["VrfRipInterfaceBfd"]; ok {
								nestedVrfRipInterface.Bfd.Misc = o.Misc["VrfRipInterfaceBfd"]
							}
							if oVrfRipInterface.Bfd.Profile != nil {
								nestedVrfRipInterface.Bfd.Profile = oVrfRipInterface.Bfd.Profile
							}
						}
						if oVrfRipInterface.InterfaceInboundDistributeList != nil {
							nestedVrfRipInterface.InterfaceInboundDistributeList = &VrfRipInterfaceInterfaceInboundDistributeListXml{}
							if _, ok := o.Misc["VrfRipInterfaceInterfaceInboundDistributeList"]; ok {
								nestedVrfRipInterface.InterfaceInboundDistributeList.Misc = o.Misc["VrfRipInterfaceInterfaceInboundDistributeList"]
							}
							if oVrfRipInterface.InterfaceInboundDistributeList.AccessList != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList.AccessList = oVrfRipInterface.InterfaceInboundDistributeList.AccessList
							}
							if oVrfRipInterface.InterfaceInboundDistributeList.Metric != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList.Metric = oVrfRipInterface.InterfaceInboundDistributeList.Metric
							}
						}
						if oVrfRipInterface.InterfaceOutboundDistributeList != nil {
							nestedVrfRipInterface.InterfaceOutboundDistributeList = &VrfRipInterfaceInterfaceOutboundDistributeListXml{}
							if _, ok := o.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"]; ok {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.Misc = o.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"]
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList.AccessList != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.AccessList = oVrfRipInterface.InterfaceOutboundDistributeList.AccessList
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList.Metric != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.Metric = oVrfRipInterface.InterfaceOutboundDistributeList.Metric
							}
						}
						nestedVrf.Rip.Interface = append(nestedVrf.Rip.Interface, nestedVrfRipInterface)
					}
				}
			}
			nestedVrfCol = append(nestedVrfCol, nestedVrf)
		}
		entry.Vrf = nestedVrfCol
	}

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	var nestedVrfCol []VrfXml_11_0_2
	if o.Vrf != nil {
		nestedVrfCol = []VrfXml_11_0_2{}
		for _, oVrf := range o.Vrf {
			nestedVrf := VrfXml_11_0_2{}
			if _, ok := o.Misc["Vrf"]; ok {
				nestedVrf.Misc = o.Misc["Vrf"]
			}
			if oVrf.Name != "" {
				nestedVrf.Name = oVrf.Name
			}
			if oVrf.Interface != nil {
				nestedVrf.Interface = util.StrToMem(oVrf.Interface)
			}
			if oVrf.AdminDists != nil {
				nestedVrf.AdminDists = &VrfAdminDistsXml_11_0_2{}
				if _, ok := o.Misc["VrfAdminDists"]; ok {
					nestedVrf.AdminDists.Misc = o.Misc["VrfAdminDists"]
				}
				if oVrf.AdminDists.Static != nil {
					nestedVrf.AdminDists.Static = oVrf.AdminDists.Static
				}
				if oVrf.AdminDists.StaticIpv6 != nil {
					nestedVrf.AdminDists.StaticIpv6 = oVrf.AdminDists.StaticIpv6
				}
				if oVrf.AdminDists.OspfInter != nil {
					nestedVrf.AdminDists.OspfInter = oVrf.AdminDists.OspfInter
				}
				if oVrf.AdminDists.OspfIntra != nil {
					nestedVrf.AdminDists.OspfIntra = oVrf.AdminDists.OspfIntra
				}
				if oVrf.AdminDists.OspfExt != nil {
					nestedVrf.AdminDists.OspfExt = oVrf.AdminDists.OspfExt
				}
				if oVrf.AdminDists.Ospfv3Inter != nil {
					nestedVrf.AdminDists.Ospfv3Inter = oVrf.AdminDists.Ospfv3Inter
				}
				if oVrf.AdminDists.Ospfv3Intra != nil {
					nestedVrf.AdminDists.Ospfv3Intra = oVrf.AdminDists.Ospfv3Intra
				}
				if oVrf.AdminDists.Ospfv3Ext != nil {
					nestedVrf.AdminDists.Ospfv3Ext = oVrf.AdminDists.Ospfv3Ext
				}
				if oVrf.AdminDists.BgpInternal != nil {
					nestedVrf.AdminDists.BgpInternal = oVrf.AdminDists.BgpInternal
				}
				if oVrf.AdminDists.BgpExternal != nil {
					nestedVrf.AdminDists.BgpExternal = oVrf.AdminDists.BgpExternal
				}
				if oVrf.AdminDists.BgpLocal != nil {
					nestedVrf.AdminDists.BgpLocal = oVrf.AdminDists.BgpLocal
				}
				if oVrf.AdminDists.Rip != nil {
					nestedVrf.AdminDists.Rip = oVrf.AdminDists.Rip
				}
			}
			if oVrf.RibFilter != nil {
				nestedVrf.RibFilter = &VrfRibFilterXml_11_0_2{}
				if _, ok := o.Misc["VrfRibFilter"]; ok {
					nestedVrf.RibFilter.Misc = o.Misc["VrfRibFilter"]
				}
				if oVrf.RibFilter.Ipv4 != nil {
					nestedVrf.RibFilter.Ipv4 = &VrfRibFilterIpv4Xml_11_0_2{}
					if _, ok := o.Misc["VrfRibFilterIpv4"]; ok {
						nestedVrf.RibFilter.Ipv4.Misc = o.Misc["VrfRibFilterIpv4"]
					}
					if oVrf.RibFilter.Ipv4.Static != nil {
						nestedVrf.RibFilter.Ipv4.Static = &VrfRibFilterIpv4StaticXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv4Static"]; ok {
							nestedVrf.RibFilter.Ipv4.Static.Misc = o.Misc["VrfRibFilterIpv4Static"]
						}
						if oVrf.RibFilter.Ipv4.Static.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Static.RouteMap = oVrf.RibFilter.Ipv4.Static.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Bgp != nil {
						nestedVrf.RibFilter.Ipv4.Bgp = &VrfRibFilterIpv4BgpXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv4Bgp"]; ok {
							nestedVrf.RibFilter.Ipv4.Bgp.Misc = o.Misc["VrfRibFilterIpv4Bgp"]
						}
						if oVrf.RibFilter.Ipv4.Bgp.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Bgp.RouteMap = oVrf.RibFilter.Ipv4.Bgp.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Ospf != nil {
						nestedVrf.RibFilter.Ipv4.Ospf = &VrfRibFilterIpv4OspfXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv4Ospf"]; ok {
							nestedVrf.RibFilter.Ipv4.Ospf.Misc = o.Misc["VrfRibFilterIpv4Ospf"]
						}
						if oVrf.RibFilter.Ipv4.Ospf.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Ospf.RouteMap = oVrf.RibFilter.Ipv4.Ospf.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv4.Rip != nil {
						nestedVrf.RibFilter.Ipv4.Rip = &VrfRibFilterIpv4RipXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv4Rip"]; ok {
							nestedVrf.RibFilter.Ipv4.Rip.Misc = o.Misc["VrfRibFilterIpv4Rip"]
						}
						if oVrf.RibFilter.Ipv4.Rip.RouteMap != nil {
							nestedVrf.RibFilter.Ipv4.Rip.RouteMap = oVrf.RibFilter.Ipv4.Rip.RouteMap
						}
					}
				}
				if oVrf.RibFilter.Ipv6 != nil {
					nestedVrf.RibFilter.Ipv6 = &VrfRibFilterIpv6Xml_11_0_2{}
					if _, ok := o.Misc["VrfRibFilterIpv6"]; ok {
						nestedVrf.RibFilter.Ipv6.Misc = o.Misc["VrfRibFilterIpv6"]
					}
					if oVrf.RibFilter.Ipv6.Static != nil {
						nestedVrf.RibFilter.Ipv6.Static = &VrfRibFilterIpv6StaticXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv6Static"]; ok {
							nestedVrf.RibFilter.Ipv6.Static.Misc = o.Misc["VrfRibFilterIpv6Static"]
						}
						if oVrf.RibFilter.Ipv6.Static.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Static.RouteMap = oVrf.RibFilter.Ipv6.Static.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv6.Bgp != nil {
						nestedVrf.RibFilter.Ipv6.Bgp = &VrfRibFilterIpv6BgpXml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv6Bgp"]; ok {
							nestedVrf.RibFilter.Ipv6.Bgp.Misc = o.Misc["VrfRibFilterIpv6Bgp"]
						}
						if oVrf.RibFilter.Ipv6.Bgp.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Bgp.RouteMap = oVrf.RibFilter.Ipv6.Bgp.RouteMap
						}
					}
					if oVrf.RibFilter.Ipv6.Ospfv3 != nil {
						nestedVrf.RibFilter.Ipv6.Ospfv3 = &VrfRibFilterIpv6Ospfv3Xml_11_0_2{}
						if _, ok := o.Misc["VrfRibFilterIpv6Ospfv3"]; ok {
							nestedVrf.RibFilter.Ipv6.Ospfv3.Misc = o.Misc["VrfRibFilterIpv6Ospfv3"]
						}
						if oVrf.RibFilter.Ipv6.Ospfv3.RouteMap != nil {
							nestedVrf.RibFilter.Ipv6.Ospfv3.RouteMap = oVrf.RibFilter.Ipv6.Ospfv3.RouteMap
						}
					}
				}
			}
			if oVrf.Bgp != nil {
				nestedVrf.Bgp = &VrfBgpXml_11_0_2{}
				if _, ok := o.Misc["VrfBgp"]; ok {
					nestedVrf.Bgp.Misc = o.Misc["VrfBgp"]
				}
				if oVrf.Bgp.Enable != nil {
					nestedVrf.Bgp.Enable = util.YesNo(oVrf.Bgp.Enable, nil)
				}
				if oVrf.Bgp.RouterId != nil {
					nestedVrf.Bgp.RouterId = oVrf.Bgp.RouterId
				}
				if oVrf.Bgp.LocalAs != nil {
					nestedVrf.Bgp.LocalAs = oVrf.Bgp.LocalAs
				}
				if oVrf.Bgp.InstallRoute != nil {
					nestedVrf.Bgp.InstallRoute = util.YesNo(oVrf.Bgp.InstallRoute, nil)
				}
				if oVrf.Bgp.EnforceFirstAs != nil {
					nestedVrf.Bgp.EnforceFirstAs = util.YesNo(oVrf.Bgp.EnforceFirstAs, nil)
				}
				if oVrf.Bgp.FastExternalFailover != nil {
					nestedVrf.Bgp.FastExternalFailover = util.YesNo(oVrf.Bgp.FastExternalFailover, nil)
				}
				if oVrf.Bgp.EcmpMultiAs != nil {
					nestedVrf.Bgp.EcmpMultiAs = util.YesNo(oVrf.Bgp.EcmpMultiAs, nil)
				}
				if oVrf.Bgp.DefaultLocalPreference != nil {
					nestedVrf.Bgp.DefaultLocalPreference = oVrf.Bgp.DefaultLocalPreference
				}
				if oVrf.Bgp.GracefulShutdown != nil {
					nestedVrf.Bgp.GracefulShutdown = util.YesNo(oVrf.Bgp.GracefulShutdown, nil)
				}
				if oVrf.Bgp.AlwaysAdvertiseNetworkRoute != nil {
					nestedVrf.Bgp.AlwaysAdvertiseNetworkRoute = util.YesNo(oVrf.Bgp.AlwaysAdvertiseNetworkRoute, nil)
				}
				if oVrf.Bgp.Med != nil {
					nestedVrf.Bgp.Med = &VrfBgpMedXml_11_0_2{}
					if _, ok := o.Misc["VrfBgpMed"]; ok {
						nestedVrf.Bgp.Med.Misc = o.Misc["VrfBgpMed"]
					}
					if oVrf.Bgp.Med.AlwaysCompareMed != nil {
						nestedVrf.Bgp.Med.AlwaysCompareMed = util.YesNo(oVrf.Bgp.Med.AlwaysCompareMed, nil)
					}
					if oVrf.Bgp.Med.DeterministicMedComparison != nil {
						nestedVrf.Bgp.Med.DeterministicMedComparison = util.YesNo(oVrf.Bgp.Med.DeterministicMedComparison, nil)
					}
				}
				if oVrf.Bgp.GracefulRestart != nil {
					nestedVrf.Bgp.GracefulRestart = &VrfBgpGracefulRestartXml_11_0_2{}
					if _, ok := o.Misc["VrfBgpGracefulRestart"]; ok {
						nestedVrf.Bgp.GracefulRestart.Misc = o.Misc["VrfBgpGracefulRestart"]
					}
					if oVrf.Bgp.GracefulRestart.Enable != nil {
						nestedVrf.Bgp.GracefulRestart.Enable = util.YesNo(oVrf.Bgp.GracefulRestart.Enable, nil)
					}
					if oVrf.Bgp.GracefulRestart.StaleRouteTime != nil {
						nestedVrf.Bgp.GracefulRestart.StaleRouteTime = oVrf.Bgp.GracefulRestart.StaleRouteTime
					}
					if oVrf.Bgp.GracefulRestart.MaxPeerRestartTime != nil {
						nestedVrf.Bgp.GracefulRestart.MaxPeerRestartTime = oVrf.Bgp.GracefulRestart.MaxPeerRestartTime
					}
					if oVrf.Bgp.GracefulRestart.LocalRestartTime != nil {
						nestedVrf.Bgp.GracefulRestart.LocalRestartTime = oVrf.Bgp.GracefulRestart.LocalRestartTime
					}
				}
				if oVrf.Bgp.GlobalBfd != nil {
					nestedVrf.Bgp.GlobalBfd = &VrfBgpGlobalBfdXml_11_0_2{}
					if _, ok := o.Misc["VrfBgpGlobalBfd"]; ok {
						nestedVrf.Bgp.GlobalBfd.Misc = o.Misc["VrfBgpGlobalBfd"]
					}
					if oVrf.Bgp.GlobalBfd.Profile != nil {
						nestedVrf.Bgp.GlobalBfd.Profile = oVrf.Bgp.GlobalBfd.Profile
					}
				}
				if oVrf.Bgp.RedistributionProfile != nil {
					nestedVrf.Bgp.RedistributionProfile = &VrfBgpRedistributionProfileXml_11_0_2{}
					if _, ok := o.Misc["VrfBgpRedistributionProfile"]; ok {
						nestedVrf.Bgp.RedistributionProfile.Misc = o.Misc["VrfBgpRedistributionProfile"]
					}
					if oVrf.Bgp.RedistributionProfile.Ipv4 != nil {
						nestedVrf.Bgp.RedistributionProfile.Ipv4 = &VrfBgpRedistributionProfileIpv4Xml_11_0_2{}
						if _, ok := o.Misc["VrfBgpRedistributionProfileIpv4"]; ok {
							nestedVrf.Bgp.RedistributionProfile.Ipv4.Misc = o.Misc["VrfBgpRedistributionProfileIpv4"]
						}
						if oVrf.Bgp.RedistributionProfile.Ipv4.Unicast != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv4.Unicast = oVrf.Bgp.RedistributionProfile.Ipv4.Unicast
						}
					}
					if oVrf.Bgp.RedistributionProfile.Ipv6 != nil {
						nestedVrf.Bgp.RedistributionProfile.Ipv6 = &VrfBgpRedistributionProfileIpv6Xml_11_0_2{}
						if _, ok := o.Misc["VrfBgpRedistributionProfileIpv6"]; ok {
							nestedVrf.Bgp.RedistributionProfile.Ipv6.Misc = o.Misc["VrfBgpRedistributionProfileIpv6"]
						}
						if oVrf.Bgp.RedistributionProfile.Ipv6.Unicast != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv6.Unicast = oVrf.Bgp.RedistributionProfile.Ipv6.Unicast
						}
					}
				}
				if oVrf.Bgp.AdvertiseNetwork != nil {
					nestedVrf.Bgp.AdvertiseNetwork = &VrfBgpAdvertiseNetworkXml_11_0_2{}
					if _, ok := o.Misc["VrfBgpAdvertiseNetwork"]; ok {
						nestedVrf.Bgp.AdvertiseNetwork.Misc = o.Misc["VrfBgpAdvertiseNetwork"]
					}
					if oVrf.Bgp.AdvertiseNetwork.Ipv4 != nil {
						nestedVrf.Bgp.AdvertiseNetwork.Ipv4 = &VrfBgpAdvertiseNetworkIpv4Xml_11_0_2{}
						if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv4"]; ok {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv4"]
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv4.Network != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = []VrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2{}
							for _, oVrfBgpAdvertiseNetworkIpv4Network := range oVrf.Bgp.AdvertiseNetwork.Ipv4.Network {
								nestedVrfBgpAdvertiseNetworkIpv4Network := VrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2{}
								if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv4Network"]; ok {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv4Network"]
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Name != "" {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Name = oVrfBgpAdvertiseNetworkIpv4Network.Name
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Unicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Unicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Unicast, nil)
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Multicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Multicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Multicast, nil)
								}
								if oVrfBgpAdvertiseNetworkIpv4Network.Backdoor != nil {
									nestedVrfBgpAdvertiseNetworkIpv4Network.Backdoor = util.YesNo(oVrfBgpAdvertiseNetworkIpv4Network.Backdoor, nil)
								}
								nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network, nestedVrfBgpAdvertiseNetworkIpv4Network)
							}
						}
					}
					if oVrf.Bgp.AdvertiseNetwork.Ipv6 != nil {
						nestedVrf.Bgp.AdvertiseNetwork.Ipv6 = &VrfBgpAdvertiseNetworkIpv6Xml_11_0_2{}
						if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv6"]; ok {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv6"]
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv6.Network != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = []VrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2{}
							for _, oVrfBgpAdvertiseNetworkIpv6Network := range oVrf.Bgp.AdvertiseNetwork.Ipv6.Network {
								nestedVrfBgpAdvertiseNetworkIpv6Network := VrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2{}
								if _, ok := o.Misc["VrfBgpAdvertiseNetworkIpv6Network"]; ok {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Misc = o.Misc["VrfBgpAdvertiseNetworkIpv6Network"]
								}
								if oVrfBgpAdvertiseNetworkIpv6Network.Name != "" {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Name = oVrfBgpAdvertiseNetworkIpv6Network.Name
								}
								if oVrfBgpAdvertiseNetworkIpv6Network.Unicast != nil {
									nestedVrfBgpAdvertiseNetworkIpv6Network.Unicast = util.YesNo(oVrfBgpAdvertiseNetworkIpv6Network.Unicast, nil)
								}
								nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network, nestedVrfBgpAdvertiseNetworkIpv6Network)
							}
						}
					}
				}
				if oVrf.Bgp.PeerGroup != nil {
					nestedVrf.Bgp.PeerGroup = []VrfBgpPeerGroupXml_11_0_2{}
					for _, oVrfBgpPeerGroup := range oVrf.Bgp.PeerGroup {
						nestedVrfBgpPeerGroup := VrfBgpPeerGroupXml_11_0_2{}
						if _, ok := o.Misc["VrfBgpPeerGroup"]; ok {
							nestedVrfBgpPeerGroup.Misc = o.Misc["VrfBgpPeerGroup"]
						}
						if oVrfBgpPeerGroup.Name != "" {
							nestedVrfBgpPeerGroup.Name = oVrfBgpPeerGroup.Name
						}
						if oVrfBgpPeerGroup.Enable != nil {
							nestedVrfBgpPeerGroup.Enable = util.YesNo(oVrfBgpPeerGroup.Enable, nil)
						}
						if oVrfBgpPeerGroup.Type != nil {
							nestedVrfBgpPeerGroup.Type = &VrfBgpPeerGroupTypeXml_11_0_2{}
							if _, ok := o.Misc["VrfBgpPeerGroupType"]; ok {
								nestedVrfBgpPeerGroup.Type.Misc = o.Misc["VrfBgpPeerGroupType"]
							}
							if oVrfBgpPeerGroup.Type.Ibgp != nil {
								nestedVrfBgpPeerGroup.Type.Ibgp = &VrfBgpPeerGroupTypeIbgpXml_11_0_2{}
								if _, ok := o.Misc["VrfBgpPeerGroupTypeIbgp"]; ok {
									nestedVrfBgpPeerGroup.Type.Ibgp.Misc = o.Misc["VrfBgpPeerGroupTypeIbgp"]
								}
							}
							if oVrfBgpPeerGroup.Type.Ebgp != nil {
								nestedVrfBgpPeerGroup.Type.Ebgp = &VrfBgpPeerGroupTypeEbgpXml_11_0_2{}
								if _, ok := o.Misc["VrfBgpPeerGroupTypeEbgp"]; ok {
									nestedVrfBgpPeerGroup.Type.Ebgp.Misc = o.Misc["VrfBgpPeerGroupTypeEbgp"]
								}
							}
						}
						if oVrfBgpPeerGroup.AddressFamily != nil {
							nestedVrfBgpPeerGroup.AddressFamily = &VrfBgpPeerGroupAddressFamilyXml_11_0_2{}
							if _, ok := o.Misc["VrfBgpPeerGroupAddressFamily"]; ok {
								nestedVrfBgpPeerGroup.AddressFamily.Misc = o.Misc["VrfBgpPeerGroupAddressFamily"]
							}
							if oVrfBgpPeerGroup.AddressFamily.Ipv4 != nil {
								nestedVrfBgpPeerGroup.AddressFamily.Ipv4 = oVrfBgpPeerGroup.AddressFamily.Ipv4
							}
							if oVrfBgpPeerGroup.AddressFamily.Ipv6 != nil {
								nestedVrfBgpPeerGroup.AddressFamily.Ipv6 = oVrfBgpPeerGroup.AddressFamily.Ipv6
							}
						}
						if oVrfBgpPeerGroup.FilteringProfile != nil {
							nestedVrfBgpPeerGroup.FilteringProfile = &VrfBgpPeerGroupFilteringProfileXml_11_0_2{}
							if _, ok := o.Misc["VrfBgpPeerGroupFilteringProfile"]; ok {
								nestedVrfBgpPeerGroup.FilteringProfile.Misc = o.Misc["VrfBgpPeerGroupFilteringProfile"]
							}
							if oVrfBgpPeerGroup.FilteringProfile.Ipv4 != nil {
								nestedVrfBgpPeerGroup.FilteringProfile.Ipv4 = oVrfBgpPeerGroup.FilteringProfile.Ipv4
							}
							if oVrfBgpPeerGroup.FilteringProfile.Ipv6 != nil {
								nestedVrfBgpPeerGroup.FilteringProfile.Ipv6 = oVrfBgpPeerGroup.FilteringProfile.Ipv6
							}
						}
						if oVrfBgpPeerGroup.ConnectionOptions != nil {
							nestedVrfBgpPeerGroup.ConnectionOptions = &VrfBgpPeerGroupConnectionOptionsXml_11_0_2{}
							if _, ok := o.Misc["VrfBgpPeerGroupConnectionOptions"]; ok {
								nestedVrfBgpPeerGroup.ConnectionOptions.Misc = o.Misc["VrfBgpPeerGroupConnectionOptions"]
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Timers != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Timers = oVrfBgpPeerGroup.ConnectionOptions.Timers
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Multihop != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Multihop = oVrfBgpPeerGroup.ConnectionOptions.Multihop
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Authentication != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Authentication = oVrfBgpPeerGroup.ConnectionOptions.Authentication
							}
							if oVrfBgpPeerGroup.ConnectionOptions.Dampening != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions.Dampening = oVrfBgpPeerGroup.ConnectionOptions.Dampening
							}
						}
						if oVrfBgpPeerGroup.Peer != nil {
							nestedVrfBgpPeerGroup.Peer = []VrfBgpPeerGroupPeerXml_11_0_2{}
							for _, oVrfBgpPeerGroupPeer := range oVrfBgpPeerGroup.Peer {
								nestedVrfBgpPeerGroupPeer := VrfBgpPeerGroupPeerXml_11_0_2{}
								if _, ok := o.Misc["VrfBgpPeerGroupPeer"]; ok {
									nestedVrfBgpPeerGroupPeer.Misc = o.Misc["VrfBgpPeerGroupPeer"]
								}
								if oVrfBgpPeerGroupPeer.Name != "" {
									nestedVrfBgpPeerGroupPeer.Name = oVrfBgpPeerGroupPeer.Name
								}
								if oVrfBgpPeerGroupPeer.Enable != nil {
									nestedVrfBgpPeerGroupPeer.Enable = util.YesNo(oVrfBgpPeerGroupPeer.Enable, nil)
								}
								if oVrfBgpPeerGroupPeer.Passive != nil {
									nestedVrfBgpPeerGroupPeer.Passive = util.YesNo(oVrfBgpPeerGroupPeer.Passive, nil)
								}
								if oVrfBgpPeerGroupPeer.PeerAs != nil {
									nestedVrfBgpPeerGroupPeer.PeerAs = oVrfBgpPeerGroupPeer.PeerAs
								}
								if oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
									nestedVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.YesNo(oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
								}
								if oVrfBgpPeerGroupPeer.Inherit != nil {
									nestedVrfBgpPeerGroupPeer.Inherit = &VrfBgpPeerGroupPeerInheritXml_11_0_2{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerInherit"]; ok {
										nestedVrfBgpPeerGroupPeer.Inherit.Misc = o.Misc["VrfBgpPeerGroupPeerInherit"]
									}
									if oVrfBgpPeerGroupPeer.Inherit.Yes != nil {
										nestedVrfBgpPeerGroupPeer.Inherit.Yes = &VrfBgpPeerGroupPeerInheritYesXml_11_0_2{}
										if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritYes"]; ok {
											nestedVrfBgpPeerGroupPeer.Inherit.Yes.Misc = o.Misc["VrfBgpPeerGroupPeerInheritYes"]
										}
									}
									if oVrfBgpPeerGroupPeer.Inherit.No != nil {
										nestedVrfBgpPeerGroupPeer.Inherit.No = &VrfBgpPeerGroupPeerInheritNoXml_11_0_2{}
										if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNo"]; ok {
											nestedVrfBgpPeerGroupPeer.Inherit.No.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNo"]
										}
										if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily = &VrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2{}
											if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"]; ok {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"]
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6
											}
										}
										if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile = &VrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2{}
											if _, ok := o.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"]; ok {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc = o.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"]
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6
											}
										}
									}
								}
								if oVrfBgpPeerGroupPeer.LocalAddress != nil {
									nestedVrfBgpPeerGroupPeer.LocalAddress = &VrfBgpPeerGroupPeerLocalAddressXml_11_0_2{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerLocalAddress"]; ok {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Misc = o.Misc["VrfBgpPeerGroupPeerLocalAddress"]
									}
									if oVrfBgpPeerGroupPeer.LocalAddress.Interface != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Interface = oVrfBgpPeerGroupPeer.LocalAddress.Interface
									}
									if oVrfBgpPeerGroupPeer.LocalAddress.Ip != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress.Ip = oVrfBgpPeerGroupPeer.LocalAddress.Ip
									}
								}
								if oVrfBgpPeerGroupPeer.PeerAddress != nil {
									nestedVrfBgpPeerGroupPeer.PeerAddress = &VrfBgpPeerGroupPeerPeerAddressXml_11_0_2{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerPeerAddress"]; ok {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Misc = o.Misc["VrfBgpPeerGroupPeerPeerAddress"]
									}
									if oVrfBgpPeerGroupPeer.PeerAddress.Ip != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Ip = oVrfBgpPeerGroupPeer.PeerAddress.Ip
									}
									if oVrfBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress.Fqdn = oVrfBgpPeerGroupPeer.PeerAddress.Fqdn
									}
								}
								if oVrfBgpPeerGroupPeer.ConnectionOptions != nil {
									nestedVrfBgpPeerGroupPeer.ConnectionOptions = &VrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerConnectionOptions"]; ok {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Misc = o.Misc["VrfBgpPeerGroupPeerConnectionOptions"]
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Timers != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Timers = oVrfBgpPeerGroupPeer.ConnectionOptions.Timers
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Multihop = oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Authentication = oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions.Dampening = oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening
									}
								}
								if oVrfBgpPeerGroupPeer.Bfd != nil {
									nestedVrfBgpPeerGroupPeer.Bfd = &VrfBgpPeerGroupPeerBfdXml_11_0_2{}
									if _, ok := o.Misc["VrfBgpPeerGroupPeerBfd"]; ok {
										nestedVrfBgpPeerGroupPeer.Bfd.Misc = o.Misc["VrfBgpPeerGroupPeerBfd"]
									}
									if oVrfBgpPeerGroupPeer.Bfd.Profile != nil {
										nestedVrfBgpPeerGroupPeer.Bfd.Profile = oVrfBgpPeerGroupPeer.Bfd.Profile
									}
								}
								nestedVrfBgpPeerGroup.Peer = append(nestedVrfBgpPeerGroup.Peer, nestedVrfBgpPeerGroupPeer)
							}
						}
						nestedVrf.Bgp.PeerGroup = append(nestedVrf.Bgp.PeerGroup, nestedVrfBgpPeerGroup)
					}
				}
				if oVrf.Bgp.AggregateRoutes != nil {
					nestedVrf.Bgp.AggregateRoutes = []VrfBgpAggregateRoutesXml_11_0_2{}
					for _, oVrfBgpAggregateRoutes := range oVrf.Bgp.AggregateRoutes {
						nestedVrfBgpAggregateRoutes := VrfBgpAggregateRoutesXml_11_0_2{}
						if _, ok := o.Misc["VrfBgpAggregateRoutes"]; ok {
							nestedVrfBgpAggregateRoutes.Misc = o.Misc["VrfBgpAggregateRoutes"]
						}
						if oVrfBgpAggregateRoutes.Name != "" {
							nestedVrfBgpAggregateRoutes.Name = oVrfBgpAggregateRoutes.Name
						}
						if oVrfBgpAggregateRoutes.Description != nil {
							nestedVrfBgpAggregateRoutes.Description = oVrfBgpAggregateRoutes.Description
						}
						if oVrfBgpAggregateRoutes.Enable != nil {
							nestedVrfBgpAggregateRoutes.Enable = util.YesNo(oVrfBgpAggregateRoutes.Enable, nil)
						}
						if oVrfBgpAggregateRoutes.SummaryOnly != nil {
							nestedVrfBgpAggregateRoutes.SummaryOnly = util.YesNo(oVrfBgpAggregateRoutes.SummaryOnly, nil)
						}
						if oVrfBgpAggregateRoutes.AsSet != nil {
							nestedVrfBgpAggregateRoutes.AsSet = util.YesNo(oVrfBgpAggregateRoutes.AsSet, nil)
						}
						if oVrfBgpAggregateRoutes.SameMed != nil {
							nestedVrfBgpAggregateRoutes.SameMed = util.YesNo(oVrfBgpAggregateRoutes.SameMed, nil)
						}
						if oVrfBgpAggregateRoutes.Type != nil {
							nestedVrfBgpAggregateRoutes.Type = &VrfBgpAggregateRoutesTypeXml_11_0_2{}
							if _, ok := o.Misc["VrfBgpAggregateRoutesType"]; ok {
								nestedVrfBgpAggregateRoutes.Type.Misc = o.Misc["VrfBgpAggregateRoutesType"]
							}
							if oVrfBgpAggregateRoutes.Type.Ipv4 != nil {
								nestedVrfBgpAggregateRoutes.Type.Ipv4 = &VrfBgpAggregateRoutesTypeIpv4Xml_11_0_2{}
								if _, ok := o.Misc["VrfBgpAggregateRoutesTypeIpv4"]; ok {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.Misc = o.Misc["VrfBgpAggregateRoutesTypeIpv4"]
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap
								}
							}
							if oVrfBgpAggregateRoutes.Type.Ipv6 != nil {
								nestedVrfBgpAggregateRoutes.Type.Ipv6 = &VrfBgpAggregateRoutesTypeIpv6Xml_11_0_2{}
								if _, ok := o.Misc["VrfBgpAggregateRoutesTypeIpv6"]; ok {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.Misc = o.Misc["VrfBgpAggregateRoutesTypeIpv6"]
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap
								}
							}
						}
						nestedVrf.Bgp.AggregateRoutes = append(nestedVrf.Bgp.AggregateRoutes, nestedVrfBgpAggregateRoutes)
					}
				}
			}
			if oVrf.RoutingTable != nil {
				nestedVrf.RoutingTable = &VrfRoutingTableXml_11_0_2{}
				if _, ok := o.Misc["VrfRoutingTable"]; ok {
					nestedVrf.RoutingTable.Misc = o.Misc["VrfRoutingTable"]
				}
				if oVrf.RoutingTable.Ip != nil {
					nestedVrf.RoutingTable.Ip = &VrfRoutingTableIpXml_11_0_2{}
					if _, ok := o.Misc["VrfRoutingTableIp"]; ok {
						nestedVrf.RoutingTable.Ip.Misc = o.Misc["VrfRoutingTableIp"]
					}
					if oVrf.RoutingTable.Ip.StaticRoute != nil {
						nestedVrf.RoutingTable.Ip.StaticRoute = []VrfRoutingTableIpStaticRouteXml_11_0_2{}
						for _, oVrfRoutingTableIpStaticRoute := range oVrf.RoutingTable.Ip.StaticRoute {
							nestedVrfRoutingTableIpStaticRoute := VrfRoutingTableIpStaticRouteXml_11_0_2{}
							if _, ok := o.Misc["VrfRoutingTableIpStaticRoute"]; ok {
								nestedVrfRoutingTableIpStaticRoute.Misc = o.Misc["VrfRoutingTableIpStaticRoute"]
							}
							if oVrfRoutingTableIpStaticRoute.Name != "" {
								nestedVrfRoutingTableIpStaticRoute.Name = oVrfRoutingTableIpStaticRoute.Name
							}
							if oVrfRoutingTableIpStaticRoute.Destination != nil {
								nestedVrfRoutingTableIpStaticRoute.Destination = oVrfRoutingTableIpStaticRoute.Destination
							}
							if oVrfRoutingTableIpStaticRoute.Interface != nil {
								nestedVrfRoutingTableIpStaticRoute.Interface = oVrfRoutingTableIpStaticRoute.Interface
							}
							if oVrfRoutingTableIpStaticRoute.AdminDist != nil {
								nestedVrfRoutingTableIpStaticRoute.AdminDist = oVrfRoutingTableIpStaticRoute.AdminDist
							}
							if oVrfRoutingTableIpStaticRoute.Metric != nil {
								nestedVrfRoutingTableIpStaticRoute.Metric = oVrfRoutingTableIpStaticRoute.Metric
							}
							if oVrfRoutingTableIpStaticRoute.Nexthop != nil {
								nestedVrfRoutingTableIpStaticRoute.Nexthop = &VrfRoutingTableIpStaticRouteNexthopXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRouteNexthop"]; ok {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Misc = o.Misc["VrfRoutingTableIpStaticRouteNexthop"]
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.Discard != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard = &VrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2{}
									if _, ok := o.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"]; ok {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc = o.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"]
									}
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.IpAddress = oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.NextLr != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.NextLr = oVrfRoutingTableIpStaticRoute.Nexthop.NextLr
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn
								}
							}
							if oVrfRoutingTableIpStaticRoute.Bfd != nil {
								nestedVrfRoutingTableIpStaticRoute.Bfd = &VrfRoutingTableIpStaticRouteBfdXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRouteBfd"]; ok {
									nestedVrfRoutingTableIpStaticRoute.Bfd.Misc = o.Misc["VrfRoutingTableIpStaticRouteBfd"]
								}
								if oVrfRoutingTableIpStaticRoute.Bfd.Profile != nil {
									nestedVrfRoutingTableIpStaticRoute.Bfd.Profile = oVrfRoutingTableIpStaticRoute.Bfd.Profile
								}
							}
							if oVrfRoutingTableIpStaticRoute.PathMonitor != nil {
								nestedVrfRoutingTableIpStaticRoute.PathMonitor = &VrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpStaticRoutePathMonitor"]; ok {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.Misc = o.Misc["VrfRoutingTableIpStaticRoutePathMonitor"]
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.Enable = util.YesNo(oVrfRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2{}
									for _, oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
										nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2{}
										if _, ok := o.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"]; ok {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"]
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
										}
										if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
										}
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
									}
								}
							}
							nestedVrf.RoutingTable.Ip.StaticRoute = append(nestedVrf.RoutingTable.Ip.StaticRoute, nestedVrfRoutingTableIpStaticRoute)
						}
					}
				}
				if oVrf.RoutingTable.Ipv6 != nil {
					nestedVrf.RoutingTable.Ipv6 = &VrfRoutingTableIpv6Xml_11_0_2{}
					if _, ok := o.Misc["VrfRoutingTableIpv6"]; ok {
						nestedVrf.RoutingTable.Ipv6.Misc = o.Misc["VrfRoutingTableIpv6"]
					}
					if oVrf.RoutingTable.Ipv6.StaticRoute != nil {
						nestedVrf.RoutingTable.Ipv6.StaticRoute = []VrfRoutingTableIpv6StaticRouteXml_11_0_2{}
						for _, oVrfRoutingTableIpv6StaticRoute := range oVrf.RoutingTable.Ipv6.StaticRoute {
							nestedVrfRoutingTableIpv6StaticRoute := VrfRoutingTableIpv6StaticRouteXml_11_0_2{}
							if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoute"]; ok {
								nestedVrfRoutingTableIpv6StaticRoute.Misc = o.Misc["VrfRoutingTableIpv6StaticRoute"]
							}
							if oVrfRoutingTableIpv6StaticRoute.Name != "" {
								nestedVrfRoutingTableIpv6StaticRoute.Name = oVrfRoutingTableIpv6StaticRoute.Name
							}
							if oVrfRoutingTableIpv6StaticRoute.Destination != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Destination = oVrfRoutingTableIpv6StaticRoute.Destination
							}
							if oVrfRoutingTableIpv6StaticRoute.Interface != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Interface = oVrfRoutingTableIpv6StaticRoute.Interface
							}
							if oVrfRoutingTableIpv6StaticRoute.AdminDist != nil {
								nestedVrfRoutingTableIpv6StaticRoute.AdminDist = oVrfRoutingTableIpv6StaticRoute.AdminDist
							}
							if oVrfRoutingTableIpv6StaticRoute.Metric != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Metric = oVrfRoutingTableIpv6StaticRoute.Metric
							}
							if oVrfRoutingTableIpv6StaticRoute.Nexthop != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Nexthop = &VrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteNexthop"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteNexthop"]
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard = &VrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2{}
									if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"]; ok {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"]
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr = oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr
								}
							}
							if oVrfRoutingTableIpv6StaticRoute.Bfd != nil {
								nestedVrfRoutingTableIpv6StaticRoute.Bfd = &VrfRoutingTableIpv6StaticRouteBfdXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRouteBfd"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd.Misc = o.Misc["VrfRoutingTableIpv6StaticRouteBfd"]
								}
								if oVrfRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd.Profile = oVrfRoutingTableIpv6StaticRoute.Bfd.Profile
								}
							}
							if oVrfRoutingTableIpv6StaticRoute.PathMonitor != nil {
								nestedVrfRoutingTableIpv6StaticRoute.PathMonitor = &VrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2{}
								if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"]; ok {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc = o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"]
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.YesNo(oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2{}
									for _, oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
										nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2{}
										if _, ok := o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]; ok {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
										}
										if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
										}
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
									}
								}
							}
							nestedVrf.RoutingTable.Ipv6.StaticRoute = append(nestedVrf.RoutingTable.Ipv6.StaticRoute, nestedVrfRoutingTableIpv6StaticRoute)
						}
					}
				}
			}
			if oVrf.Ospf != nil {
				nestedVrf.Ospf = &VrfOspfXml_11_0_2{}
				if _, ok := o.Misc["VrfOspf"]; ok {
					nestedVrf.Ospf.Misc = o.Misc["VrfOspf"]
				}
				if oVrf.Ospf.RouterId != nil {
					nestedVrf.Ospf.RouterId = oVrf.Ospf.RouterId
				}
				if oVrf.Ospf.Enable != nil {
					nestedVrf.Ospf.Enable = util.YesNo(oVrf.Ospf.Enable, nil)
				}
				if oVrf.Ospf.Rfc1583 != nil {
					nestedVrf.Ospf.Rfc1583 = util.YesNo(oVrf.Ospf.Rfc1583, nil)
				}
				if oVrf.Ospf.SpfTimer != nil {
					nestedVrf.Ospf.SpfTimer = oVrf.Ospf.SpfTimer
				}
				if oVrf.Ospf.GlobalIfTimer != nil {
					nestedVrf.Ospf.GlobalIfTimer = oVrf.Ospf.GlobalIfTimer
				}
				if oVrf.Ospf.RedistributionProfile != nil {
					nestedVrf.Ospf.RedistributionProfile = oVrf.Ospf.RedistributionProfile
				}
				if oVrf.Ospf.GlobalBfd != nil {
					nestedVrf.Ospf.GlobalBfd = &VrfOspfGlobalBfdXml_11_0_2{}
					if _, ok := o.Misc["VrfOspfGlobalBfd"]; ok {
						nestedVrf.Ospf.GlobalBfd.Misc = o.Misc["VrfOspfGlobalBfd"]
					}
					if oVrf.Ospf.GlobalBfd.Profile != nil {
						nestedVrf.Ospf.GlobalBfd.Profile = oVrf.Ospf.GlobalBfd.Profile
					}
				}
				if oVrf.Ospf.GracefulRestart != nil {
					nestedVrf.Ospf.GracefulRestart = &VrfOspfGracefulRestartXml_11_0_2{}
					if _, ok := o.Misc["VrfOspfGracefulRestart"]; ok {
						nestedVrf.Ospf.GracefulRestart.Misc = o.Misc["VrfOspfGracefulRestart"]
					}
					if oVrf.Ospf.GracefulRestart.Enable != nil {
						nestedVrf.Ospf.GracefulRestart.Enable = util.YesNo(oVrf.Ospf.GracefulRestart.Enable, nil)
					}
					if oVrf.Ospf.GracefulRestart.GracePeriod != nil {
						nestedVrf.Ospf.GracefulRestart.GracePeriod = oVrf.Ospf.GracefulRestart.GracePeriod
					}
					if oVrf.Ospf.GracefulRestart.HelperEnable != nil {
						nestedVrf.Ospf.GracefulRestart.HelperEnable = util.YesNo(oVrf.Ospf.GracefulRestart.HelperEnable, nil)
					}
					if oVrf.Ospf.GracefulRestart.StrictLSAChecking != nil {
						nestedVrf.Ospf.GracefulRestart.StrictLSAChecking = util.YesNo(oVrf.Ospf.GracefulRestart.StrictLSAChecking, nil)
					}
					if oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedVrf.Ospf.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime
					}
				}
				if oVrf.Ospf.Area != nil {
					nestedVrf.Ospf.Area = []VrfOspfAreaXml_11_0_2{}
					for _, oVrfOspfArea := range oVrf.Ospf.Area {
						nestedVrfOspfArea := VrfOspfAreaXml_11_0_2{}
						if _, ok := o.Misc["VrfOspfArea"]; ok {
							nestedVrfOspfArea.Misc = o.Misc["VrfOspfArea"]
						}
						if oVrfOspfArea.Name != "" {
							nestedVrfOspfArea.Name = oVrfOspfArea.Name
						}
						if oVrfOspfArea.Authentication != nil {
							nestedVrfOspfArea.Authentication = oVrfOspfArea.Authentication
						}
						if oVrfOspfArea.Type != nil {
							nestedVrfOspfArea.Type = &VrfOspfAreaTypeXml_11_0_2{}
							if _, ok := o.Misc["VrfOspfAreaType"]; ok {
								nestedVrfOspfArea.Type.Misc = o.Misc["VrfOspfAreaType"]
							}
							if oVrfOspfArea.Type.Normal != nil {
								nestedVrfOspfArea.Type.Normal = &VrfOspfAreaTypeNormalXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaTypeNormal"]; ok {
									nestedVrfOspfArea.Type.Normal.Misc = o.Misc["VrfOspfAreaTypeNormal"]
								}
								if oVrfOspfArea.Type.Normal.Abr != nil {
									nestedVrfOspfArea.Type.Normal.Abr = &VrfOspfAreaTypeNormalAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaTypeNormalAbr"]; ok {
										nestedVrfOspfArea.Type.Normal.Abr.Misc = o.Misc["VrfOspfAreaTypeNormalAbr"]
									}
									if oVrfOspfArea.Type.Normal.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.ImportList = oVrfOspfArea.Type.Normal.Abr.ImportList
									}
									if oVrfOspfArea.Type.Normal.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.ExportList = oVrfOspfArea.Type.Normal.Abr.ExportList
									}
									if oVrfOspfArea.Type.Normal.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.InboundFilterList = oVrfOspfArea.Type.Normal.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Normal.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Normal.Abr.OutboundFilterList = oVrfOspfArea.Type.Normal.Abr.OutboundFilterList
									}
								}
							}
							if oVrfOspfArea.Type.Stub != nil {
								nestedVrfOspfArea.Type.Stub = &VrfOspfAreaTypeStubXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaTypeStub"]; ok {
									nestedVrfOspfArea.Type.Stub.Misc = o.Misc["VrfOspfAreaTypeStub"]
								}
								if oVrfOspfArea.Type.Stub.NoSummary != nil {
									nestedVrfOspfArea.Type.Stub.NoSummary = util.YesNo(oVrfOspfArea.Type.Stub.NoSummary, nil)
								}
								if oVrfOspfArea.Type.Stub.Abr != nil {
									nestedVrfOspfArea.Type.Stub.Abr = &VrfOspfAreaTypeStubAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaTypeStubAbr"]; ok {
										nestedVrfOspfArea.Type.Stub.Abr.Misc = o.Misc["VrfOspfAreaTypeStubAbr"]
									}
									if oVrfOspfArea.Type.Stub.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.ImportList = oVrfOspfArea.Type.Stub.Abr.ImportList
									}
									if oVrfOspfArea.Type.Stub.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.ExportList = oVrfOspfArea.Type.Stub.Abr.ExportList
									}
									if oVrfOspfArea.Type.Stub.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.InboundFilterList = oVrfOspfArea.Type.Stub.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Stub.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Stub.Abr.OutboundFilterList = oVrfOspfArea.Type.Stub.Abr.OutboundFilterList
									}
								}
								if oVrfOspfArea.Type.Stub.DefaultRouteMetric != nil {
									nestedVrfOspfArea.Type.Stub.DefaultRouteMetric = oVrfOspfArea.Type.Stub.DefaultRouteMetric
								}
							}
							if oVrfOspfArea.Type.Nssa != nil {
								nestedVrfOspfArea.Type.Nssa = &VrfOspfAreaTypeNssaXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaTypeNssa"]; ok {
									nestedVrfOspfArea.Type.Nssa.Misc = o.Misc["VrfOspfAreaTypeNssa"]
								}
								if oVrfOspfArea.Type.Nssa.NoSummary != nil {
									nestedVrfOspfArea.Type.Nssa.NoSummary = util.YesNo(oVrfOspfArea.Type.Nssa.NoSummary, nil)
								}
								if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate != nil {
									nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate = &VrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"]; ok {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc = o.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"]
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType
									}
								}
								if oVrfOspfArea.Type.Nssa.Abr != nil {
									nestedVrfOspfArea.Type.Nssa.Abr = &VrfOspfAreaTypeNssaAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaTypeNssaAbr"]; ok {
										nestedVrfOspfArea.Type.Nssa.Abr.Misc = o.Misc["VrfOspfAreaTypeNssaAbr"]
									}
									if oVrfOspfArea.Type.Nssa.Abr.ImportList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.ImportList = oVrfOspfArea.Type.Nssa.Abr.ImportList
									}
									if oVrfOspfArea.Type.Nssa.Abr.ExportList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.ExportList = oVrfOspfArea.Type.Nssa.Abr.ExportList
									}
									if oVrfOspfArea.Type.Nssa.Abr.InboundFilterList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.InboundFilterList = oVrfOspfArea.Type.Nssa.Abr.InboundFilterList
									}
									if oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.OutboundFilterList = oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList
									}
									if oVrfOspfArea.Type.Nssa.Abr.NssaExtRange != nil {
										nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = []VrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2{}
										for _, oVrfOspfAreaTypeNssaAbrNssaExtRange := range oVrfOspfArea.Type.Nssa.Abr.NssaExtRange {
											nestedVrfOspfAreaTypeNssaAbrNssaExtRange := VrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2{}
											if _, ok := o.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"]; ok {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Misc = o.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"]
											}
											if oVrfOspfAreaTypeNssaAbrNssaExtRange.Name != "" {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Name = oVrfOspfAreaTypeNssaAbrNssaExtRange.Name
											}
											if oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise != nil {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise = util.YesNo(oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise, nil)
											}
											nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfAreaTypeNssaAbrNssaExtRange)
										}
									}
								}
							}
						}
						if oVrfOspfArea.Range != nil {
							nestedVrfOspfArea.Range = []VrfOspfAreaRangeXml_11_0_2{}
							for _, oVrfOspfAreaRange := range oVrfOspfArea.Range {
								nestedVrfOspfAreaRange := VrfOspfAreaRangeXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaRange"]; ok {
									nestedVrfOspfAreaRange.Misc = o.Misc["VrfOspfAreaRange"]
								}
								if oVrfOspfAreaRange.Name != "" {
									nestedVrfOspfAreaRange.Name = oVrfOspfAreaRange.Name
								}
								if oVrfOspfAreaRange.Advertise != nil {
									nestedVrfOspfAreaRange.Advertise = util.YesNo(oVrfOspfAreaRange.Advertise, nil)
								}
								nestedVrfOspfArea.Range = append(nestedVrfOspfArea.Range, nestedVrfOspfAreaRange)
							}
						}
						if oVrfOspfArea.Interface != nil {
							nestedVrfOspfArea.Interface = []VrfOspfAreaInterfaceXml_11_0_2{}
							for _, oVrfOspfAreaInterface := range oVrfOspfArea.Interface {
								nestedVrfOspfAreaInterface := VrfOspfAreaInterfaceXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaInterface"]; ok {
									nestedVrfOspfAreaInterface.Misc = o.Misc["VrfOspfAreaInterface"]
								}
								if oVrfOspfAreaInterface.Name != "" {
									nestedVrfOspfAreaInterface.Name = oVrfOspfAreaInterface.Name
								}
								if oVrfOspfAreaInterface.Enable != nil {
									nestedVrfOspfAreaInterface.Enable = util.YesNo(oVrfOspfAreaInterface.Enable, nil)
								}
								if oVrfOspfAreaInterface.MtuIgnore != nil {
									nestedVrfOspfAreaInterface.MtuIgnore = util.YesNo(oVrfOspfAreaInterface.MtuIgnore, nil)
								}
								if oVrfOspfAreaInterface.Passive != nil {
									nestedVrfOspfAreaInterface.Passive = util.YesNo(oVrfOspfAreaInterface.Passive, nil)
								}
								if oVrfOspfAreaInterface.Priority != nil {
									nestedVrfOspfAreaInterface.Priority = oVrfOspfAreaInterface.Priority
								}
								if oVrfOspfAreaInterface.Metric != nil {
									nestedVrfOspfAreaInterface.Metric = oVrfOspfAreaInterface.Metric
								}
								if oVrfOspfAreaInterface.Authentication != nil {
									nestedVrfOspfAreaInterface.Authentication = oVrfOspfAreaInterface.Authentication
								}
								if oVrfOspfAreaInterface.Timing != nil {
									nestedVrfOspfAreaInterface.Timing = oVrfOspfAreaInterface.Timing
								}
								if oVrfOspfAreaInterface.LinkType != nil {
									nestedVrfOspfAreaInterface.LinkType = &VrfOspfAreaInterfaceLinkTypeXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaInterfaceLinkType"]; ok {
										nestedVrfOspfAreaInterface.LinkType.Misc = o.Misc["VrfOspfAreaInterfaceLinkType"]
									}
									if oVrfOspfAreaInterface.LinkType.Broadcast != nil {
										nestedVrfOspfAreaInterface.LinkType.Broadcast = &VrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"]; ok {
											nestedVrfOspfAreaInterface.LinkType.Broadcast.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"]
										}
									}
									if oVrfOspfAreaInterface.LinkType.P2p != nil {
										nestedVrfOspfAreaInterface.LinkType.P2p = &VrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2p"]; ok {
											nestedVrfOspfAreaInterface.LinkType.P2p.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2p"]
										}
									}
									if oVrfOspfAreaInterface.LinkType.P2mp != nil {
										nestedVrfOspfAreaInterface.LinkType.P2mp = &VrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"]; ok {
											nestedVrfOspfAreaInterface.LinkType.P2mp.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"]
										}
										if oVrfOspfAreaInterface.LinkType.P2mp.Neighbor != nil {
											nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = []VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2{}
											for _, oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfAreaInterface.LinkType.P2mp.Neighbor {
												nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := VrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2{}
												if _, ok := o.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"]; ok {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc = o.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"]
												}
												if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name
												}
												if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority
												}
												nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor)
											}
										}
									}
								}
								if oVrfOspfAreaInterface.Bfd != nil {
									nestedVrfOspfAreaInterface.Bfd = &VrfOspfAreaInterfaceBfdXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaInterfaceBfd"]; ok {
										nestedVrfOspfAreaInterface.Bfd.Misc = o.Misc["VrfOspfAreaInterfaceBfd"]
									}
									if oVrfOspfAreaInterface.Bfd.Profile != nil {
										nestedVrfOspfAreaInterface.Bfd.Profile = oVrfOspfAreaInterface.Bfd.Profile
									}
								}
								nestedVrfOspfArea.Interface = append(nestedVrfOspfArea.Interface, nestedVrfOspfAreaInterface)
							}
						}
						if oVrfOspfArea.VirtualLink != nil {
							nestedVrfOspfArea.VirtualLink = []VrfOspfAreaVirtualLinkXml_11_0_2{}
							for _, oVrfOspfAreaVirtualLink := range oVrfOspfArea.VirtualLink {
								nestedVrfOspfAreaVirtualLink := VrfOspfAreaVirtualLinkXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfAreaVirtualLink"]; ok {
									nestedVrfOspfAreaVirtualLink.Misc = o.Misc["VrfOspfAreaVirtualLink"]
								}
								if oVrfOspfAreaVirtualLink.Name != "" {
									nestedVrfOspfAreaVirtualLink.Name = oVrfOspfAreaVirtualLink.Name
								}
								if oVrfOspfAreaVirtualLink.NeighborId != nil {
									nestedVrfOspfAreaVirtualLink.NeighborId = oVrfOspfAreaVirtualLink.NeighborId
								}
								if oVrfOspfAreaVirtualLink.TransitAreaId != nil {
									nestedVrfOspfAreaVirtualLink.TransitAreaId = oVrfOspfAreaVirtualLink.TransitAreaId
								}
								if oVrfOspfAreaVirtualLink.Enable != nil {
									nestedVrfOspfAreaVirtualLink.Enable = util.YesNo(oVrfOspfAreaVirtualLink.Enable, nil)
								}
								if oVrfOspfAreaVirtualLink.InstanceId != nil {
									nestedVrfOspfAreaVirtualLink.InstanceId = oVrfOspfAreaVirtualLink.InstanceId
								}
								if oVrfOspfAreaVirtualLink.Timing != nil {
									nestedVrfOspfAreaVirtualLink.Timing = oVrfOspfAreaVirtualLink.Timing
								}
								if oVrfOspfAreaVirtualLink.Authentication != nil {
									nestedVrfOspfAreaVirtualLink.Authentication = oVrfOspfAreaVirtualLink.Authentication
								}
								if oVrfOspfAreaVirtualLink.Bfd != nil {
									nestedVrfOspfAreaVirtualLink.Bfd = &VrfOspfAreaVirtualLinkBfdXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfAreaVirtualLinkBfd"]; ok {
										nestedVrfOspfAreaVirtualLink.Bfd.Misc = o.Misc["VrfOspfAreaVirtualLinkBfd"]
									}
									if oVrfOspfAreaVirtualLink.Bfd.Profile != nil {
										nestedVrfOspfAreaVirtualLink.Bfd.Profile = oVrfOspfAreaVirtualLink.Bfd.Profile
									}
								}
								nestedVrfOspfArea.VirtualLink = append(nestedVrfOspfArea.VirtualLink, nestedVrfOspfAreaVirtualLink)
							}
						}
						nestedVrf.Ospf.Area = append(nestedVrf.Ospf.Area, nestedVrfOspfArea)
					}
				}
			}
			if oVrf.Ospfv3 != nil {
				nestedVrf.Ospfv3 = &VrfOspfv3Xml_11_0_2{}
				if _, ok := o.Misc["VrfOspfv3"]; ok {
					nestedVrf.Ospfv3.Misc = o.Misc["VrfOspfv3"]
				}
				if oVrf.Ospfv3.Enable != nil {
					nestedVrf.Ospfv3.Enable = util.YesNo(oVrf.Ospfv3.Enable, nil)
				}
				if oVrf.Ospfv3.RouterId != nil {
					nestedVrf.Ospfv3.RouterId = oVrf.Ospfv3.RouterId
				}
				if oVrf.Ospfv3.DisableTransitTraffic != nil {
					nestedVrf.Ospfv3.DisableTransitTraffic = util.YesNo(oVrf.Ospfv3.DisableTransitTraffic, nil)
				}
				if oVrf.Ospfv3.SpfTimer != nil {
					nestedVrf.Ospfv3.SpfTimer = oVrf.Ospfv3.SpfTimer
				}
				if oVrf.Ospfv3.GlobalIfTimer != nil {
					nestedVrf.Ospfv3.GlobalIfTimer = oVrf.Ospfv3.GlobalIfTimer
				}
				if oVrf.Ospfv3.RedistributionProfile != nil {
					nestedVrf.Ospfv3.RedistributionProfile = oVrf.Ospfv3.RedistributionProfile
				}
				if oVrf.Ospfv3.GlobalBfd != nil {
					nestedVrf.Ospfv3.GlobalBfd = &VrfOspfv3GlobalBfdXml_11_0_2{}
					if _, ok := o.Misc["VrfOspfv3GlobalBfd"]; ok {
						nestedVrf.Ospfv3.GlobalBfd.Misc = o.Misc["VrfOspfv3GlobalBfd"]
					}
					if oVrf.Ospfv3.GlobalBfd.Profile != nil {
						nestedVrf.Ospfv3.GlobalBfd.Profile = oVrf.Ospfv3.GlobalBfd.Profile
					}
				}
				if oVrf.Ospfv3.GracefulRestart != nil {
					nestedVrf.Ospfv3.GracefulRestart = &VrfOspfv3GracefulRestartXml_11_0_2{}
					if _, ok := o.Misc["VrfOspfv3GracefulRestart"]; ok {
						nestedVrf.Ospfv3.GracefulRestart.Misc = o.Misc["VrfOspfv3GracefulRestart"]
					}
					if oVrf.Ospfv3.GracefulRestart.Enable != nil {
						nestedVrf.Ospfv3.GracefulRestart.Enable = util.YesNo(oVrf.Ospfv3.GracefulRestart.Enable, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.GracePeriod != nil {
						nestedVrf.Ospfv3.GracefulRestart.GracePeriod = oVrf.Ospfv3.GracefulRestart.GracePeriod
					}
					if oVrf.Ospfv3.GracefulRestart.HelperEnable != nil {
						nestedVrf.Ospfv3.GracefulRestart.HelperEnable = util.YesNo(oVrf.Ospfv3.GracefulRestart.HelperEnable, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
						nestedVrf.Ospfv3.GracefulRestart.StrictLSAChecking = util.YesNo(oVrf.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
					}
					if oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime
					}
				}
				if oVrf.Ospfv3.Area != nil {
					nestedVrf.Ospfv3.Area = []VrfOspfv3AreaXml_11_0_2{}
					for _, oVrfOspfv3Area := range oVrf.Ospfv3.Area {
						nestedVrfOspfv3Area := VrfOspfv3AreaXml_11_0_2{}
						if _, ok := o.Misc["VrfOspfv3Area"]; ok {
							nestedVrfOspfv3Area.Misc = o.Misc["VrfOspfv3Area"]
						}
						if oVrfOspfv3Area.Name != "" {
							nestedVrfOspfv3Area.Name = oVrfOspfv3Area.Name
						}
						if oVrfOspfv3Area.Authentication != nil {
							nestedVrfOspfv3Area.Authentication = oVrfOspfv3Area.Authentication
						}
						if oVrfOspfv3Area.Type != nil {
							nestedVrfOspfv3Area.Type = &VrfOspfv3AreaTypeXml_11_0_2{}
							if _, ok := o.Misc["VrfOspfv3AreaType"]; ok {
								nestedVrfOspfv3Area.Type.Misc = o.Misc["VrfOspfv3AreaType"]
							}
							if oVrfOspfv3Area.Type.Normal != nil {
								nestedVrfOspfv3Area.Type.Normal = &VrfOspfv3AreaTypeNormalXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeNormal"]; ok {
									nestedVrfOspfv3Area.Type.Normal.Misc = o.Misc["VrfOspfv3AreaTypeNormal"]
								}
								if oVrfOspfv3Area.Type.Normal.Abr != nil {
									nestedVrfOspfv3Area.Type.Normal.Abr = &VrfOspfv3AreaTypeNormalAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNormalAbr"]; ok {
										nestedVrfOspfv3Area.Type.Normal.Abr.Misc = o.Misc["VrfOspfv3AreaTypeNormalAbr"]
									}
									if oVrfOspfv3Area.Type.Normal.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.ImportList = oVrfOspfv3Area.Type.Normal.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.ExportList = oVrfOspfv3Area.Type.Normal.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.InboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList
									}
								}
							}
							if oVrfOspfv3Area.Type.Stub != nil {
								nestedVrfOspfv3Area.Type.Stub = &VrfOspfv3AreaTypeStubXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeStub"]; ok {
									nestedVrfOspfv3Area.Type.Stub.Misc = o.Misc["VrfOspfv3AreaTypeStub"]
								}
								if oVrfOspfv3Area.Type.Stub.NoSummary != nil {
									nestedVrfOspfv3Area.Type.Stub.NoSummary = util.YesNo(oVrfOspfv3Area.Type.Stub.NoSummary, nil)
								}
								if oVrfOspfv3Area.Type.Stub.Abr != nil {
									nestedVrfOspfv3Area.Type.Stub.Abr = &VrfOspfv3AreaTypeStubAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeStubAbr"]; ok {
										nestedVrfOspfv3Area.Type.Stub.Abr.Misc = o.Misc["VrfOspfv3AreaTypeStubAbr"]
									}
									if oVrfOspfv3Area.Type.Stub.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.ImportList = oVrfOspfv3Area.Type.Stub.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.ExportList = oVrfOspfv3Area.Type.Stub.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.InboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList
									}
								}
								if oVrfOspfv3Area.Type.Stub.DefaultRouteMetric != nil {
									nestedVrfOspfv3Area.Type.Stub.DefaultRouteMetric = oVrfOspfv3Area.Type.Stub.DefaultRouteMetric
								}
							}
							if oVrfOspfv3Area.Type.Nssa != nil {
								nestedVrfOspfv3Area.Type.Nssa = &VrfOspfv3AreaTypeNssaXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaTypeNssa"]; ok {
									nestedVrfOspfv3Area.Type.Nssa.Misc = o.Misc["VrfOspfv3AreaTypeNssa"]
								}
								if oVrfOspfv3Area.Type.Nssa.NoSummary != nil {
									nestedVrfOspfv3Area.Type.Nssa.NoSummary = util.YesNo(oVrfOspfv3Area.Type.Nssa.NoSummary, nil)
								}
								if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate != nil {
									nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate = &VrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"]; ok {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc = o.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"]
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType
									}
								}
								if oVrfOspfv3Area.Type.Nssa.Abr != nil {
									nestedVrfOspfv3Area.Type.Nssa.Abr = &VrfOspfv3AreaTypeNssaAbrXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaTypeNssaAbr"]; ok {
										nestedVrfOspfv3Area.Type.Nssa.Abr.Misc = o.Misc["VrfOspfv3AreaTypeNssaAbr"]
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.ImportList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.ImportList = oVrfOspfv3Area.Type.Nssa.Abr.ImportList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.ExportList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.ExportList = oVrfOspfv3Area.Type.Nssa.Abr.ExportList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList
									}
									if oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = []VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2{}
										for _, oVrfOspfv3AreaTypeNssaAbrNssaExtRange := range oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange {
											nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange := VrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2{}
											if _, ok := o.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"]; ok {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc = o.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"]
											}
											if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name != "" {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name
											}
											if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise != nil {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise = util.YesNo(oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise, nil)
											}
											nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange)
										}
									}
								}
							}
						}
						if oVrfOspfv3Area.Range != nil {
							nestedVrfOspfv3Area.Range = []VrfOspfv3AreaRangeXml_11_0_2{}
							for _, oVrfOspfv3AreaRange := range oVrfOspfv3Area.Range {
								nestedVrfOspfv3AreaRange := VrfOspfv3AreaRangeXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaRange"]; ok {
									nestedVrfOspfv3AreaRange.Misc = o.Misc["VrfOspfv3AreaRange"]
								}
								if oVrfOspfv3AreaRange.Name != "" {
									nestedVrfOspfv3AreaRange.Name = oVrfOspfv3AreaRange.Name
								}
								if oVrfOspfv3AreaRange.Advertise != nil {
									nestedVrfOspfv3AreaRange.Advertise = util.YesNo(oVrfOspfv3AreaRange.Advertise, nil)
								}
								nestedVrfOspfv3Area.Range = append(nestedVrfOspfv3Area.Range, nestedVrfOspfv3AreaRange)
							}
						}
						if oVrfOspfv3Area.Interface != nil {
							nestedVrfOspfv3Area.Interface = []VrfOspfv3AreaInterfaceXml_11_0_2{}
							for _, oVrfOspfv3AreaInterface := range oVrfOspfv3Area.Interface {
								nestedVrfOspfv3AreaInterface := VrfOspfv3AreaInterfaceXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaInterface"]; ok {
									nestedVrfOspfv3AreaInterface.Misc = o.Misc["VrfOspfv3AreaInterface"]
								}
								if oVrfOspfv3AreaInterface.Name != "" {
									nestedVrfOspfv3AreaInterface.Name = oVrfOspfv3AreaInterface.Name
								}
								if oVrfOspfv3AreaInterface.Enable != nil {
									nestedVrfOspfv3AreaInterface.Enable = util.YesNo(oVrfOspfv3AreaInterface.Enable, nil)
								}
								if oVrfOspfv3AreaInterface.MtuIgnore != nil {
									nestedVrfOspfv3AreaInterface.MtuIgnore = util.YesNo(oVrfOspfv3AreaInterface.MtuIgnore, nil)
								}
								if oVrfOspfv3AreaInterface.Passive != nil {
									nestedVrfOspfv3AreaInterface.Passive = util.YesNo(oVrfOspfv3AreaInterface.Passive, nil)
								}
								if oVrfOspfv3AreaInterface.Priority != nil {
									nestedVrfOspfv3AreaInterface.Priority = oVrfOspfv3AreaInterface.Priority
								}
								if oVrfOspfv3AreaInterface.Metric != nil {
									nestedVrfOspfv3AreaInterface.Metric = oVrfOspfv3AreaInterface.Metric
								}
								if oVrfOspfv3AreaInterface.InstanceId != nil {
									nestedVrfOspfv3AreaInterface.InstanceId = oVrfOspfv3AreaInterface.InstanceId
								}
								if oVrfOspfv3AreaInterface.Authentication != nil {
									nestedVrfOspfv3AreaInterface.Authentication = oVrfOspfv3AreaInterface.Authentication
								}
								if oVrfOspfv3AreaInterface.Timing != nil {
									nestedVrfOspfv3AreaInterface.Timing = oVrfOspfv3AreaInterface.Timing
								}
								if oVrfOspfv3AreaInterface.LinkType != nil {
									nestedVrfOspfv3AreaInterface.LinkType = &VrfOspfv3AreaInterfaceLinkTypeXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkType"]; ok {
										nestedVrfOspfv3AreaInterface.LinkType.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkType"]
									}
									if oVrfOspfv3AreaInterface.LinkType.Broadcast != nil {
										nestedVrfOspfv3AreaInterface.LinkType.Broadcast = &VrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.Broadcast.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"]
										}
									}
									if oVrfOspfv3AreaInterface.LinkType.P2p != nil {
										nestedVrfOspfv3AreaInterface.LinkType.P2p = &VrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.P2p.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"]
										}
									}
									if oVrfOspfv3AreaInterface.LinkType.P2mp != nil {
										nestedVrfOspfv3AreaInterface.LinkType.P2mp = &VrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2{}
										if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"]; ok {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"]
										}
										if oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2{}
											for _, oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor {
												nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := VrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2{}
												if _, ok := o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"]; ok {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc = o.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"]
												}
												if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name
												}
												if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority
												}
												nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor)
											}
										}
									}
								}
								if oVrfOspfv3AreaInterface.Bfd != nil {
									nestedVrfOspfv3AreaInterface.Bfd = &VrfOspfv3AreaInterfaceBfdXml_11_0_2{}
									if _, ok := o.Misc["VrfOspfv3AreaInterfaceBfd"]; ok {
										nestedVrfOspfv3AreaInterface.Bfd.Misc = o.Misc["VrfOspfv3AreaInterfaceBfd"]
									}
									if oVrfOspfv3AreaInterface.Bfd.Profile != nil {
										nestedVrfOspfv3AreaInterface.Bfd.Profile = oVrfOspfv3AreaInterface.Bfd.Profile
									}
								}
								nestedVrfOspfv3Area.Interface = append(nestedVrfOspfv3Area.Interface, nestedVrfOspfv3AreaInterface)
							}
						}
						if oVrfOspfv3Area.VirtualLink != nil {
							nestedVrfOspfv3Area.VirtualLink = []VrfOspfv3AreaVirtualLinkXml_11_0_2{}
							for _, oVrfOspfv3AreaVirtualLink := range oVrfOspfv3Area.VirtualLink {
								nestedVrfOspfv3AreaVirtualLink := VrfOspfv3AreaVirtualLinkXml_11_0_2{}
								if _, ok := o.Misc["VrfOspfv3AreaVirtualLink"]; ok {
									nestedVrfOspfv3AreaVirtualLink.Misc = o.Misc["VrfOspfv3AreaVirtualLink"]
								}
								if oVrfOspfv3AreaVirtualLink.Name != "" {
									nestedVrfOspfv3AreaVirtualLink.Name = oVrfOspfv3AreaVirtualLink.Name
								}
								if oVrfOspfv3AreaVirtualLink.NeighborId != nil {
									nestedVrfOspfv3AreaVirtualLink.NeighborId = oVrfOspfv3AreaVirtualLink.NeighborId
								}
								if oVrfOspfv3AreaVirtualLink.TransitAreaId != nil {
									nestedVrfOspfv3AreaVirtualLink.TransitAreaId = oVrfOspfv3AreaVirtualLink.TransitAreaId
								}
								if oVrfOspfv3AreaVirtualLink.Enable != nil {
									nestedVrfOspfv3AreaVirtualLink.Enable = util.YesNo(oVrfOspfv3AreaVirtualLink.Enable, nil)
								}
								if oVrfOspfv3AreaVirtualLink.InstanceId != nil {
									nestedVrfOspfv3AreaVirtualLink.InstanceId = oVrfOspfv3AreaVirtualLink.InstanceId
								}
								if oVrfOspfv3AreaVirtualLink.Timing != nil {
									nestedVrfOspfv3AreaVirtualLink.Timing = oVrfOspfv3AreaVirtualLink.Timing
								}
								if oVrfOspfv3AreaVirtualLink.Authentication != nil {
									nestedVrfOspfv3AreaVirtualLink.Authentication = oVrfOspfv3AreaVirtualLink.Authentication
								}
								nestedVrfOspfv3Area.VirtualLink = append(nestedVrfOspfv3Area.VirtualLink, nestedVrfOspfv3AreaVirtualLink)
							}
						}
						nestedVrf.Ospfv3.Area = append(nestedVrf.Ospfv3.Area, nestedVrfOspfv3Area)
					}
				}
			}
			if oVrf.Ecmp != nil {
				nestedVrf.Ecmp = &VrfEcmpXml_11_0_2{}
				if _, ok := o.Misc["VrfEcmp"]; ok {
					nestedVrf.Ecmp.Misc = o.Misc["VrfEcmp"]
				}
				if oVrf.Ecmp.Enable != nil {
					nestedVrf.Ecmp.Enable = util.YesNo(oVrf.Ecmp.Enable, nil)
				}
				if oVrf.Ecmp.MaxPath != nil {
					nestedVrf.Ecmp.MaxPath = oVrf.Ecmp.MaxPath
				}
				if oVrf.Ecmp.SymmetricReturn != nil {
					nestedVrf.Ecmp.SymmetricReturn = util.YesNo(oVrf.Ecmp.SymmetricReturn, nil)
				}
				if oVrf.Ecmp.StrictSourcePath != nil {
					nestedVrf.Ecmp.StrictSourcePath = util.YesNo(oVrf.Ecmp.StrictSourcePath, nil)
				}
				if oVrf.Ecmp.Algorithm != nil {
					nestedVrf.Ecmp.Algorithm = &VrfEcmpAlgorithmXml_11_0_2{}
					if _, ok := o.Misc["VrfEcmpAlgorithm"]; ok {
						nestedVrf.Ecmp.Algorithm.Misc = o.Misc["VrfEcmpAlgorithm"]
					}
					if oVrf.Ecmp.Algorithm.IpModulo != nil {
						nestedVrf.Ecmp.Algorithm.IpModulo = &VrfEcmpAlgorithmIpModuloXml_11_0_2{}
						if _, ok := o.Misc["VrfEcmpAlgorithmIpModulo"]; ok {
							nestedVrf.Ecmp.Algorithm.IpModulo.Misc = o.Misc["VrfEcmpAlgorithmIpModulo"]
						}
					}
					if oVrf.Ecmp.Algorithm.IpHash != nil {
						nestedVrf.Ecmp.Algorithm.IpHash = &VrfEcmpAlgorithmIpHashXml_11_0_2{}
						if _, ok := o.Misc["VrfEcmpAlgorithmIpHash"]; ok {
							nestedVrf.Ecmp.Algorithm.IpHash.Misc = o.Misc["VrfEcmpAlgorithmIpHash"]
						}
						if oVrf.Ecmp.Algorithm.IpHash.SrcOnly != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.SrcOnly = util.YesNo(oVrf.Ecmp.Algorithm.IpHash.SrcOnly, nil)
						}
						if oVrf.Ecmp.Algorithm.IpHash.UsePort != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.UsePort = util.YesNo(oVrf.Ecmp.Algorithm.IpHash.UsePort, nil)
						}
						if oVrf.Ecmp.Algorithm.IpHash.HashSeed != nil {
							nestedVrf.Ecmp.Algorithm.IpHash.HashSeed = oVrf.Ecmp.Algorithm.IpHash.HashSeed
						}
					}
					if oVrf.Ecmp.Algorithm.WeightedRoundRobin != nil {
						nestedVrf.Ecmp.Algorithm.WeightedRoundRobin = &VrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2{}
						if _, ok := o.Misc["VrfEcmpAlgorithmWeightedRoundRobin"]; ok {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc = o.Misc["VrfEcmpAlgorithmWeightedRoundRobin"]
						}
						if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = []VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2{}
							for _, oVrfEcmpAlgorithmWeightedRoundRobinInterface := range oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface {
								nestedVrfEcmpAlgorithmWeightedRoundRobinInterface := VrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2{}
								if _, ok := o.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"]; ok {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc = o.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"]
								}
								if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Name = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name
								}
								if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight
								}
								nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = append(nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface, nestedVrfEcmpAlgorithmWeightedRoundRobinInterface)
							}
						}
					}
					if oVrf.Ecmp.Algorithm.BalancedRoundRobin != nil {
						nestedVrf.Ecmp.Algorithm.BalancedRoundRobin = &VrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2{}
						if _, ok := o.Misc["VrfEcmpAlgorithmBalancedRoundRobin"]; ok {
							nestedVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc = o.Misc["VrfEcmpAlgorithmBalancedRoundRobin"]
						}
					}
				}
			}
			if oVrf.Multicast != nil {
				nestedVrf.Multicast = &VrfMulticastXml_11_0_2{}
				if _, ok := o.Misc["VrfMulticast"]; ok {
					nestedVrf.Multicast.Misc = o.Misc["VrfMulticast"]
				}
				if oVrf.Multicast.Enable != nil {
					nestedVrf.Multicast.Enable = util.YesNo(oVrf.Multicast.Enable, nil)
				}
				if oVrf.Multicast.StaticRoute != nil {
					nestedVrf.Multicast.StaticRoute = []VrfMulticastStaticRouteXml_11_0_2{}
					for _, oVrfMulticastStaticRoute := range oVrf.Multicast.StaticRoute {
						nestedVrfMulticastStaticRoute := VrfMulticastStaticRouteXml_11_0_2{}
						if _, ok := o.Misc["VrfMulticastStaticRoute"]; ok {
							nestedVrfMulticastStaticRoute.Misc = o.Misc["VrfMulticastStaticRoute"]
						}
						if oVrfMulticastStaticRoute.Name != "" {
							nestedVrfMulticastStaticRoute.Name = oVrfMulticastStaticRoute.Name
						}
						if oVrfMulticastStaticRoute.Destination != nil {
							nestedVrfMulticastStaticRoute.Destination = oVrfMulticastStaticRoute.Destination
						}
						if oVrfMulticastStaticRoute.Interface != nil {
							nestedVrfMulticastStaticRoute.Interface = oVrfMulticastStaticRoute.Interface
						}
						if oVrfMulticastStaticRoute.Preference != nil {
							nestedVrfMulticastStaticRoute.Preference = oVrfMulticastStaticRoute.Preference
						}
						if oVrfMulticastStaticRoute.Nexthop != nil {
							nestedVrfMulticastStaticRoute.Nexthop = &VrfMulticastStaticRouteNexthopXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastStaticRouteNexthop"]; ok {
								nestedVrfMulticastStaticRoute.Nexthop.Misc = o.Misc["VrfMulticastStaticRouteNexthop"]
							}
							if oVrfMulticastStaticRoute.Nexthop.IpAddress != nil {
								nestedVrfMulticastStaticRoute.Nexthop.IpAddress = oVrfMulticastStaticRoute.Nexthop.IpAddress
							}
						}
						nestedVrf.Multicast.StaticRoute = append(nestedVrf.Multicast.StaticRoute, nestedVrfMulticastStaticRoute)
					}
				}
				if oVrf.Multicast.Pim != nil {
					nestedVrf.Multicast.Pim = &VrfMulticastPimXml_11_0_2{}
					if _, ok := o.Misc["VrfMulticastPim"]; ok {
						nestedVrf.Multicast.Pim.Misc = o.Misc["VrfMulticastPim"]
					}
					if oVrf.Multicast.Pim.Enable != nil {
						nestedVrf.Multicast.Pim.Enable = util.YesNo(oVrf.Multicast.Pim.Enable, nil)
					}
					if oVrf.Multicast.Pim.RpfLookupMode != nil {
						nestedVrf.Multicast.Pim.RpfLookupMode = oVrf.Multicast.Pim.RpfLookupMode
					}
					if oVrf.Multicast.Pim.RouteAgeoutTime != nil {
						nestedVrf.Multicast.Pim.RouteAgeoutTime = oVrf.Multicast.Pim.RouteAgeoutTime
					}
					if oVrf.Multicast.Pim.IfTimerGlobal != nil {
						nestedVrf.Multicast.Pim.IfTimerGlobal = oVrf.Multicast.Pim.IfTimerGlobal
					}
					if oVrf.Multicast.Pim.GroupPermission != nil {
						nestedVrf.Multicast.Pim.GroupPermission = oVrf.Multicast.Pim.GroupPermission
					}
					if oVrf.Multicast.Pim.SsmAddressSpace != nil {
						nestedVrf.Multicast.Pim.SsmAddressSpace = &VrfMulticastPimSsmAddressSpaceXml_11_0_2{}
						if _, ok := o.Misc["VrfMulticastPimSsmAddressSpace"]; ok {
							nestedVrf.Multicast.Pim.SsmAddressSpace.Misc = o.Misc["VrfMulticastPimSsmAddressSpace"]
						}
						if oVrf.Multicast.Pim.SsmAddressSpace.GroupList != nil {
							nestedVrf.Multicast.Pim.SsmAddressSpace.GroupList = oVrf.Multicast.Pim.SsmAddressSpace.GroupList
						}
					}
					if oVrf.Multicast.Pim.Rp != nil {
						nestedVrf.Multicast.Pim.Rp = &VrfMulticastPimRpXml_11_0_2{}
						if _, ok := o.Misc["VrfMulticastPimRp"]; ok {
							nestedVrf.Multicast.Pim.Rp.Misc = o.Misc["VrfMulticastPimRp"]
						}
						if oVrf.Multicast.Pim.Rp.LocalRp != nil {
							nestedVrf.Multicast.Pim.Rp.LocalRp = &VrfMulticastPimRpLocalRpXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastPimRpLocalRp"]; ok {
								nestedVrf.Multicast.Pim.Rp.LocalRp.Misc = o.Misc["VrfMulticastPimRpLocalRp"]
							}
							if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp = &VrfMulticastPimRpLocalRpStaticRpXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastPimRpLocalRpStaticRp"]; ok {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc = o.Misc["VrfMulticastPimRpLocalRpStaticRp"]
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override = util.YesNo(oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override, nil)
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList
								}
							}
							if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp = &VrfMulticastPimRpLocalRpCandidateRpXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastPimRpLocalRpCandidateRp"]; ok {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc = o.Misc["VrfMulticastPimRpLocalRpCandidateRp"]
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList
								}
							}
						}
						if oVrf.Multicast.Pim.Rp.ExternalRp != nil {
							nestedVrf.Multicast.Pim.Rp.ExternalRp = []VrfMulticastPimRpExternalRpXml_11_0_2{}
							for _, oVrfMulticastPimRpExternalRp := range oVrf.Multicast.Pim.Rp.ExternalRp {
								nestedVrfMulticastPimRpExternalRp := VrfMulticastPimRpExternalRpXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastPimRpExternalRp"]; ok {
									nestedVrfMulticastPimRpExternalRp.Misc = o.Misc["VrfMulticastPimRpExternalRp"]
								}
								if oVrfMulticastPimRpExternalRp.Name != "" {
									nestedVrfMulticastPimRpExternalRp.Name = oVrfMulticastPimRpExternalRp.Name
								}
								if oVrfMulticastPimRpExternalRp.GroupList != nil {
									nestedVrfMulticastPimRpExternalRp.GroupList = oVrfMulticastPimRpExternalRp.GroupList
								}
								if oVrfMulticastPimRpExternalRp.Override != nil {
									nestedVrfMulticastPimRpExternalRp.Override = util.YesNo(oVrfMulticastPimRpExternalRp.Override, nil)
								}
								nestedVrf.Multicast.Pim.Rp.ExternalRp = append(nestedVrf.Multicast.Pim.Rp.ExternalRp, nestedVrfMulticastPimRpExternalRp)
							}
						}
					}
					if oVrf.Multicast.Pim.SptThreshold != nil {
						nestedVrf.Multicast.Pim.SptThreshold = []VrfMulticastPimSptThresholdXml_11_0_2{}
						for _, oVrfMulticastPimSptThreshold := range oVrf.Multicast.Pim.SptThreshold {
							nestedVrfMulticastPimSptThreshold := VrfMulticastPimSptThresholdXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastPimSptThreshold"]; ok {
								nestedVrfMulticastPimSptThreshold.Misc = o.Misc["VrfMulticastPimSptThreshold"]
							}
							if oVrfMulticastPimSptThreshold.Name != "" {
								nestedVrfMulticastPimSptThreshold.Name = oVrfMulticastPimSptThreshold.Name
							}
							if oVrfMulticastPimSptThreshold.Threshold != nil {
								nestedVrfMulticastPimSptThreshold.Threshold = oVrfMulticastPimSptThreshold.Threshold
							}
							nestedVrf.Multicast.Pim.SptThreshold = append(nestedVrf.Multicast.Pim.SptThreshold, nestedVrfMulticastPimSptThreshold)
						}
					}
					if oVrf.Multicast.Pim.Interface != nil {
						nestedVrf.Multicast.Pim.Interface = []VrfMulticastPimInterfaceXml_11_0_2{}
						for _, oVrfMulticastPimInterface := range oVrf.Multicast.Pim.Interface {
							nestedVrfMulticastPimInterface := VrfMulticastPimInterfaceXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastPimInterface"]; ok {
								nestedVrfMulticastPimInterface.Misc = o.Misc["VrfMulticastPimInterface"]
							}
							if oVrfMulticastPimInterface.Name != "" {
								nestedVrfMulticastPimInterface.Name = oVrfMulticastPimInterface.Name
							}
							if oVrfMulticastPimInterface.Description != nil {
								nestedVrfMulticastPimInterface.Description = oVrfMulticastPimInterface.Description
							}
							if oVrfMulticastPimInterface.DrPriority != nil {
								nestedVrfMulticastPimInterface.DrPriority = oVrfMulticastPimInterface.DrPriority
							}
							if oVrfMulticastPimInterface.SendBsm != nil {
								nestedVrfMulticastPimInterface.SendBsm = util.YesNo(oVrfMulticastPimInterface.SendBsm, nil)
							}
							if oVrfMulticastPimInterface.IfTimer != nil {
								nestedVrfMulticastPimInterface.IfTimer = oVrfMulticastPimInterface.IfTimer
							}
							if oVrfMulticastPimInterface.NeighborFilter != nil {
								nestedVrfMulticastPimInterface.NeighborFilter = oVrfMulticastPimInterface.NeighborFilter
							}
							nestedVrf.Multicast.Pim.Interface = append(nestedVrf.Multicast.Pim.Interface, nestedVrfMulticastPimInterface)
						}
					}
				}
				if oVrf.Multicast.Igmp != nil {
					nestedVrf.Multicast.Igmp = &VrfMulticastIgmpXml_11_0_2{}
					if _, ok := o.Misc["VrfMulticastIgmp"]; ok {
						nestedVrf.Multicast.Igmp.Misc = o.Misc["VrfMulticastIgmp"]
					}
					if oVrf.Multicast.Igmp.Enable != nil {
						nestedVrf.Multicast.Igmp.Enable = util.YesNo(oVrf.Multicast.Igmp.Enable, nil)
					}
					if oVrf.Multicast.Igmp.Dynamic != nil {
						nestedVrf.Multicast.Igmp.Dynamic = &VrfMulticastIgmpDynamicXml_11_0_2{}
						if _, ok := o.Misc["VrfMulticastIgmpDynamic"]; ok {
							nestedVrf.Multicast.Igmp.Dynamic.Misc = o.Misc["VrfMulticastIgmpDynamic"]
						}
						if oVrf.Multicast.Igmp.Dynamic.Interface != nil {
							nestedVrf.Multicast.Igmp.Dynamic.Interface = []VrfMulticastIgmpDynamicInterfaceXml_11_0_2{}
							for _, oVrfMulticastIgmpDynamicInterface := range oVrf.Multicast.Igmp.Dynamic.Interface {
								nestedVrfMulticastIgmpDynamicInterface := VrfMulticastIgmpDynamicInterfaceXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastIgmpDynamicInterface"]; ok {
									nestedVrfMulticastIgmpDynamicInterface.Misc = o.Misc["VrfMulticastIgmpDynamicInterface"]
								}
								if oVrfMulticastIgmpDynamicInterface.Name != "" {
									nestedVrfMulticastIgmpDynamicInterface.Name = oVrfMulticastIgmpDynamicInterface.Name
								}
								if oVrfMulticastIgmpDynamicInterface.Version != nil {
									nestedVrfMulticastIgmpDynamicInterface.Version = oVrfMulticastIgmpDynamicInterface.Version
								}
								if oVrfMulticastIgmpDynamicInterface.Robustness != nil {
									nestedVrfMulticastIgmpDynamicInterface.Robustness = oVrfMulticastIgmpDynamicInterface.Robustness
								}
								if oVrfMulticastIgmpDynamicInterface.GroupFilter != nil {
									nestedVrfMulticastIgmpDynamicInterface.GroupFilter = oVrfMulticastIgmpDynamicInterface.GroupFilter
								}
								if oVrfMulticastIgmpDynamicInterface.MaxGroups != nil {
									nestedVrfMulticastIgmpDynamicInterface.MaxGroups = oVrfMulticastIgmpDynamicInterface.MaxGroups
								}
								if oVrfMulticastIgmpDynamicInterface.MaxSources != nil {
									nestedVrfMulticastIgmpDynamicInterface.MaxSources = oVrfMulticastIgmpDynamicInterface.MaxSources
								}
								if oVrfMulticastIgmpDynamicInterface.QueryProfile != nil {
									nestedVrfMulticastIgmpDynamicInterface.QueryProfile = oVrfMulticastIgmpDynamicInterface.QueryProfile
								}
								if oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing != nil {
									nestedVrfMulticastIgmpDynamicInterface.RouterAlertPolicing = util.YesNo(oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing, nil)
								}
								nestedVrf.Multicast.Igmp.Dynamic.Interface = append(nestedVrf.Multicast.Igmp.Dynamic.Interface, nestedVrfMulticastIgmpDynamicInterface)
							}
						}
					}
					if oVrf.Multicast.Igmp.Static != nil {
						nestedVrf.Multicast.Igmp.Static = []VrfMulticastIgmpStaticXml_11_0_2{}
						for _, oVrfMulticastIgmpStatic := range oVrf.Multicast.Igmp.Static {
							nestedVrfMulticastIgmpStatic := VrfMulticastIgmpStaticXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastIgmpStatic"]; ok {
								nestedVrfMulticastIgmpStatic.Misc = o.Misc["VrfMulticastIgmpStatic"]
							}
							if oVrfMulticastIgmpStatic.Name != "" {
								nestedVrfMulticastIgmpStatic.Name = oVrfMulticastIgmpStatic.Name
							}
							if oVrfMulticastIgmpStatic.Interface != nil {
								nestedVrfMulticastIgmpStatic.Interface = oVrfMulticastIgmpStatic.Interface
							}
							if oVrfMulticastIgmpStatic.GroupAddress != nil {
								nestedVrfMulticastIgmpStatic.GroupAddress = oVrfMulticastIgmpStatic.GroupAddress
							}
							if oVrfMulticastIgmpStatic.SourceAddress != nil {
								nestedVrfMulticastIgmpStatic.SourceAddress = oVrfMulticastIgmpStatic.SourceAddress
							}
							nestedVrf.Multicast.Igmp.Static = append(nestedVrf.Multicast.Igmp.Static, nestedVrfMulticastIgmpStatic)
						}
					}
				}
				if oVrf.Multicast.Msdp != nil {
					nestedVrf.Multicast.Msdp = &VrfMulticastMsdpXml_11_0_2{}
					if _, ok := o.Misc["VrfMulticastMsdp"]; ok {
						nestedVrf.Multicast.Msdp.Misc = o.Misc["VrfMulticastMsdp"]
					}
					if oVrf.Multicast.Msdp.Enable != nil {
						nestedVrf.Multicast.Msdp.Enable = util.YesNo(oVrf.Multicast.Msdp.Enable, nil)
					}
					if oVrf.Multicast.Msdp.GlobalTimer != nil {
						nestedVrf.Multicast.Msdp.GlobalTimer = oVrf.Multicast.Msdp.GlobalTimer
					}
					if oVrf.Multicast.Msdp.GlobalAuthentication != nil {
						nestedVrf.Multicast.Msdp.GlobalAuthentication = oVrf.Multicast.Msdp.GlobalAuthentication
					}
					if oVrf.Multicast.Msdp.OriginatorId != nil {
						nestedVrf.Multicast.Msdp.OriginatorId = &VrfMulticastMsdpOriginatorIdXml_11_0_2{}
						if _, ok := o.Misc["VrfMulticastMsdpOriginatorId"]; ok {
							nestedVrf.Multicast.Msdp.OriginatorId.Misc = o.Misc["VrfMulticastMsdpOriginatorId"]
						}
						if oVrf.Multicast.Msdp.OriginatorId.Interface != nil {
							nestedVrf.Multicast.Msdp.OriginatorId.Interface = oVrf.Multicast.Msdp.OriginatorId.Interface
						}
						if oVrf.Multicast.Msdp.OriginatorId.Ip != nil {
							nestedVrf.Multicast.Msdp.OriginatorId.Ip = oVrf.Multicast.Msdp.OriginatorId.Ip
						}
					}
					if oVrf.Multicast.Msdp.Peer != nil {
						nestedVrf.Multicast.Msdp.Peer = []VrfMulticastMsdpPeerXml_11_0_2{}
						for _, oVrfMulticastMsdpPeer := range oVrf.Multicast.Msdp.Peer {
							nestedVrfMulticastMsdpPeer := VrfMulticastMsdpPeerXml_11_0_2{}
							if _, ok := o.Misc["VrfMulticastMsdpPeer"]; ok {
								nestedVrfMulticastMsdpPeer.Misc = o.Misc["VrfMulticastMsdpPeer"]
							}
							if oVrfMulticastMsdpPeer.Name != "" {
								nestedVrfMulticastMsdpPeer.Name = oVrfMulticastMsdpPeer.Name
							}
							if oVrfMulticastMsdpPeer.Enable != nil {
								nestedVrfMulticastMsdpPeer.Enable = util.YesNo(oVrfMulticastMsdpPeer.Enable, nil)
							}
							if oVrfMulticastMsdpPeer.PeerAs != nil {
								nestedVrfMulticastMsdpPeer.PeerAs = oVrfMulticastMsdpPeer.PeerAs
							}
							if oVrfMulticastMsdpPeer.Authentication != nil {
								nestedVrfMulticastMsdpPeer.Authentication = oVrfMulticastMsdpPeer.Authentication
							}
							if oVrfMulticastMsdpPeer.MaxSa != nil {
								nestedVrfMulticastMsdpPeer.MaxSa = oVrfMulticastMsdpPeer.MaxSa
							}
							if oVrfMulticastMsdpPeer.InboundSaFilter != nil {
								nestedVrfMulticastMsdpPeer.InboundSaFilter = oVrfMulticastMsdpPeer.InboundSaFilter
							}
							if oVrfMulticastMsdpPeer.OutboundSaFilter != nil {
								nestedVrfMulticastMsdpPeer.OutboundSaFilter = oVrfMulticastMsdpPeer.OutboundSaFilter
							}
							if oVrfMulticastMsdpPeer.LocalAddress != nil {
								nestedVrfMulticastMsdpPeer.LocalAddress = &VrfMulticastMsdpPeerLocalAddressXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastMsdpPeerLocalAddress"]; ok {
									nestedVrfMulticastMsdpPeer.LocalAddress.Misc = o.Misc["VrfMulticastMsdpPeerLocalAddress"]
								}
								if oVrfMulticastMsdpPeer.LocalAddress.Interface != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress.Interface = oVrfMulticastMsdpPeer.LocalAddress.Interface
								}
								if oVrfMulticastMsdpPeer.LocalAddress.Ip != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress.Ip = oVrfMulticastMsdpPeer.LocalAddress.Ip
								}
							}
							if oVrfMulticastMsdpPeer.PeerAddress != nil {
								nestedVrfMulticastMsdpPeer.PeerAddress = &VrfMulticastMsdpPeerPeerAddressXml_11_0_2{}
								if _, ok := o.Misc["VrfMulticastMsdpPeerPeerAddress"]; ok {
									nestedVrfMulticastMsdpPeer.PeerAddress.Misc = o.Misc["VrfMulticastMsdpPeerPeerAddress"]
								}
								if oVrfMulticastMsdpPeer.PeerAddress.Ip != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress.Ip = oVrfMulticastMsdpPeer.PeerAddress.Ip
								}
								if oVrfMulticastMsdpPeer.PeerAddress.Fqdn != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress.Fqdn = oVrfMulticastMsdpPeer.PeerAddress.Fqdn
								}
							}
							nestedVrf.Multicast.Msdp.Peer = append(nestedVrf.Multicast.Msdp.Peer, nestedVrfMulticastMsdpPeer)
						}
					}
				}
			}
			if oVrf.Rip != nil {
				nestedVrf.Rip = &VrfRipXml_11_0_2{}
				if _, ok := o.Misc["VrfRip"]; ok {
					nestedVrf.Rip.Misc = o.Misc["VrfRip"]
				}
				if oVrf.Rip.Enable != nil {
					nestedVrf.Rip.Enable = util.YesNo(oVrf.Rip.Enable, nil)
				}
				if oVrf.Rip.DefaultInformationOriginate != nil {
					nestedVrf.Rip.DefaultInformationOriginate = util.YesNo(oVrf.Rip.DefaultInformationOriginate, nil)
				}
				if oVrf.Rip.GlobalTimer != nil {
					nestedVrf.Rip.GlobalTimer = oVrf.Rip.GlobalTimer
				}
				if oVrf.Rip.AuthProfile != nil {
					nestedVrf.Rip.AuthProfile = oVrf.Rip.AuthProfile
				}
				if oVrf.Rip.RedistributionProfile != nil {
					nestedVrf.Rip.RedistributionProfile = oVrf.Rip.RedistributionProfile
				}
				if oVrf.Rip.GlobalBfd != nil {
					nestedVrf.Rip.GlobalBfd = &VrfRipGlobalBfdXml_11_0_2{}
					if _, ok := o.Misc["VrfRipGlobalBfd"]; ok {
						nestedVrf.Rip.GlobalBfd.Misc = o.Misc["VrfRipGlobalBfd"]
					}
					if oVrf.Rip.GlobalBfd.Profile != nil {
						nestedVrf.Rip.GlobalBfd.Profile = oVrf.Rip.GlobalBfd.Profile
					}
				}
				if oVrf.Rip.GlobalInboundDistributeList != nil {
					nestedVrf.Rip.GlobalInboundDistributeList = &VrfRipGlobalInboundDistributeListXml_11_0_2{}
					if _, ok := o.Misc["VrfRipGlobalInboundDistributeList"]; ok {
						nestedVrf.Rip.GlobalInboundDistributeList.Misc = o.Misc["VrfRipGlobalInboundDistributeList"]
					}
					if oVrf.Rip.GlobalInboundDistributeList.AccessList != nil {
						nestedVrf.Rip.GlobalInboundDistributeList.AccessList = oVrf.Rip.GlobalInboundDistributeList.AccessList
					}
				}
				if oVrf.Rip.GlobalOutboundDistributeList != nil {
					nestedVrf.Rip.GlobalOutboundDistributeList = &VrfRipGlobalOutboundDistributeListXml_11_0_2{}
					if _, ok := o.Misc["VrfRipGlobalOutboundDistributeList"]; ok {
						nestedVrf.Rip.GlobalOutboundDistributeList.Misc = o.Misc["VrfRipGlobalOutboundDistributeList"]
					}
					if oVrf.Rip.GlobalOutboundDistributeList.AccessList != nil {
						nestedVrf.Rip.GlobalOutboundDistributeList.AccessList = oVrf.Rip.GlobalOutboundDistributeList.AccessList
					}
				}
				if oVrf.Rip.Interface != nil {
					nestedVrf.Rip.Interface = []VrfRipInterfaceXml_11_0_2{}
					for _, oVrfRipInterface := range oVrf.Rip.Interface {
						nestedVrfRipInterface := VrfRipInterfaceXml_11_0_2{}
						if _, ok := o.Misc["VrfRipInterface"]; ok {
							nestedVrfRipInterface.Misc = o.Misc["VrfRipInterface"]
						}
						if oVrfRipInterface.Name != "" {
							nestedVrfRipInterface.Name = oVrfRipInterface.Name
						}
						if oVrfRipInterface.Enable != nil {
							nestedVrfRipInterface.Enable = util.YesNo(oVrfRipInterface.Enable, nil)
						}
						if oVrfRipInterface.Mode != nil {
							nestedVrfRipInterface.Mode = oVrfRipInterface.Mode
						}
						if oVrfRipInterface.SplitHorizon != nil {
							nestedVrfRipInterface.SplitHorizon = oVrfRipInterface.SplitHorizon
						}
						if oVrfRipInterface.Authentication != nil {
							nestedVrfRipInterface.Authentication = oVrfRipInterface.Authentication
						}
						if oVrfRipInterface.Bfd != nil {
							nestedVrfRipInterface.Bfd = &VrfRipInterfaceBfdXml_11_0_2{}
							if _, ok := o.Misc["VrfRipInterfaceBfd"]; ok {
								nestedVrfRipInterface.Bfd.Misc = o.Misc["VrfRipInterfaceBfd"]
							}
							if oVrfRipInterface.Bfd.Profile != nil {
								nestedVrfRipInterface.Bfd.Profile = oVrfRipInterface.Bfd.Profile
							}
						}
						if oVrfRipInterface.InterfaceInboundDistributeList != nil {
							nestedVrfRipInterface.InterfaceInboundDistributeList = &VrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2{}
							if _, ok := o.Misc["VrfRipInterfaceInterfaceInboundDistributeList"]; ok {
								nestedVrfRipInterface.InterfaceInboundDistributeList.Misc = o.Misc["VrfRipInterfaceInterfaceInboundDistributeList"]
							}
							if oVrfRipInterface.InterfaceInboundDistributeList.AccessList != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList.AccessList = oVrfRipInterface.InterfaceInboundDistributeList.AccessList
							}
							if oVrfRipInterface.InterfaceInboundDistributeList.Metric != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList.Metric = oVrfRipInterface.InterfaceInboundDistributeList.Metric
							}
						}
						if oVrfRipInterface.InterfaceOutboundDistributeList != nil {
							nestedVrfRipInterface.InterfaceOutboundDistributeList = &VrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2{}
							if _, ok := o.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"]; ok {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.Misc = o.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"]
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList.AccessList != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.AccessList = oVrfRipInterface.InterfaceOutboundDistributeList.AccessList
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList.Metric != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList.Metric = oVrfRipInterface.InterfaceOutboundDistributeList.Metric
							}
						}
						nestedVrf.Rip.Interface = append(nestedVrf.Rip.Interface, nestedVrfRipInterface)
					}
				}
			}
			nestedVrfCol = append(nestedVrfCol, nestedVrf)
		}
		entry.Vrf = nestedVrfCol
	}

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}
func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedVrfCol []Vrf
		if o.Vrf != nil {
			nestedVrfCol = []Vrf{}
			for _, oVrf := range o.Vrf {
				nestedVrf := Vrf{}
				if oVrf.Misc != nil {
					entry.Misc["Vrf"] = oVrf.Misc
				}
				if oVrf.Name != "" {
					nestedVrf.Name = oVrf.Name
				}
				if oVrf.Interface != nil {
					nestedVrf.Interface = util.MemToStr(oVrf.Interface)
				}
				if oVrf.AdminDists != nil {
					nestedVrf.AdminDists = &VrfAdminDists{}
					if oVrf.AdminDists.Misc != nil {
						entry.Misc["VrfAdminDists"] = oVrf.AdminDists.Misc
					}
					if oVrf.AdminDists.Static != nil {
						nestedVrf.AdminDists.Static = oVrf.AdminDists.Static
					}
					if oVrf.AdminDists.StaticIpv6 != nil {
						nestedVrf.AdminDists.StaticIpv6 = oVrf.AdminDists.StaticIpv6
					}
					if oVrf.AdminDists.OspfInter != nil {
						nestedVrf.AdminDists.OspfInter = oVrf.AdminDists.OspfInter
					}
					if oVrf.AdminDists.OspfIntra != nil {
						nestedVrf.AdminDists.OspfIntra = oVrf.AdminDists.OspfIntra
					}
					if oVrf.AdminDists.OspfExt != nil {
						nestedVrf.AdminDists.OspfExt = oVrf.AdminDists.OspfExt
					}
					if oVrf.AdminDists.Ospfv3Inter != nil {
						nestedVrf.AdminDists.Ospfv3Inter = oVrf.AdminDists.Ospfv3Inter
					}
					if oVrf.AdminDists.Ospfv3Intra != nil {
						nestedVrf.AdminDists.Ospfv3Intra = oVrf.AdminDists.Ospfv3Intra
					}
					if oVrf.AdminDists.Ospfv3Ext != nil {
						nestedVrf.AdminDists.Ospfv3Ext = oVrf.AdminDists.Ospfv3Ext
					}
					if oVrf.AdminDists.BgpInternal != nil {
						nestedVrf.AdminDists.BgpInternal = oVrf.AdminDists.BgpInternal
					}
					if oVrf.AdminDists.BgpExternal != nil {
						nestedVrf.AdminDists.BgpExternal = oVrf.AdminDists.BgpExternal
					}
					if oVrf.AdminDists.BgpLocal != nil {
						nestedVrf.AdminDists.BgpLocal = oVrf.AdminDists.BgpLocal
					}
					if oVrf.AdminDists.Rip != nil {
						nestedVrf.AdminDists.Rip = oVrf.AdminDists.Rip
					}
				}
				if oVrf.RibFilter != nil {
					nestedVrf.RibFilter = &VrfRibFilter{}
					if oVrf.RibFilter.Misc != nil {
						entry.Misc["VrfRibFilter"] = oVrf.RibFilter.Misc
					}
					if oVrf.RibFilter.Ipv4 != nil {
						nestedVrf.RibFilter.Ipv4 = &VrfRibFilterIpv4{}
						if oVrf.RibFilter.Ipv4.Misc != nil {
							entry.Misc["VrfRibFilterIpv4"] = oVrf.RibFilter.Ipv4.Misc
						}
						if oVrf.RibFilter.Ipv4.Static != nil {
							nestedVrf.RibFilter.Ipv4.Static = &VrfRibFilterIpv4Static{}
							if oVrf.RibFilter.Ipv4.Static.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Static"] = oVrf.RibFilter.Ipv4.Static.Misc
							}
							if oVrf.RibFilter.Ipv4.Static.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Static.RouteMap = oVrf.RibFilter.Ipv4.Static.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Bgp != nil {
							nestedVrf.RibFilter.Ipv4.Bgp = &VrfRibFilterIpv4Bgp{}
							if oVrf.RibFilter.Ipv4.Bgp.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Bgp"] = oVrf.RibFilter.Ipv4.Bgp.Misc
							}
							if oVrf.RibFilter.Ipv4.Bgp.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Bgp.RouteMap = oVrf.RibFilter.Ipv4.Bgp.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Ospf != nil {
							nestedVrf.RibFilter.Ipv4.Ospf = &VrfRibFilterIpv4Ospf{}
							if oVrf.RibFilter.Ipv4.Ospf.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Ospf"] = oVrf.RibFilter.Ipv4.Ospf.Misc
							}
							if oVrf.RibFilter.Ipv4.Ospf.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Ospf.RouteMap = oVrf.RibFilter.Ipv4.Ospf.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Rip != nil {
							nestedVrf.RibFilter.Ipv4.Rip = &VrfRibFilterIpv4Rip{}
							if oVrf.RibFilter.Ipv4.Rip.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Rip"] = oVrf.RibFilter.Ipv4.Rip.Misc
							}
							if oVrf.RibFilter.Ipv4.Rip.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Rip.RouteMap = oVrf.RibFilter.Ipv4.Rip.RouteMap
							}
						}
					}
					if oVrf.RibFilter.Ipv6 != nil {
						nestedVrf.RibFilter.Ipv6 = &VrfRibFilterIpv6{}
						if oVrf.RibFilter.Ipv6.Misc != nil {
							entry.Misc["VrfRibFilterIpv6"] = oVrf.RibFilter.Ipv6.Misc
						}
						if oVrf.RibFilter.Ipv6.Static != nil {
							nestedVrf.RibFilter.Ipv6.Static = &VrfRibFilterIpv6Static{}
							if oVrf.RibFilter.Ipv6.Static.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Static"] = oVrf.RibFilter.Ipv6.Static.Misc
							}
							if oVrf.RibFilter.Ipv6.Static.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Static.RouteMap = oVrf.RibFilter.Ipv6.Static.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv6.Bgp != nil {
							nestedVrf.RibFilter.Ipv6.Bgp = &VrfRibFilterIpv6Bgp{}
							if oVrf.RibFilter.Ipv6.Bgp.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Bgp"] = oVrf.RibFilter.Ipv6.Bgp.Misc
							}
							if oVrf.RibFilter.Ipv6.Bgp.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Bgp.RouteMap = oVrf.RibFilter.Ipv6.Bgp.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv6.Ospfv3 != nil {
							nestedVrf.RibFilter.Ipv6.Ospfv3 = &VrfRibFilterIpv6Ospfv3{}
							if oVrf.RibFilter.Ipv6.Ospfv3.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Ospfv3"] = oVrf.RibFilter.Ipv6.Ospfv3.Misc
							}
							if oVrf.RibFilter.Ipv6.Ospfv3.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Ospfv3.RouteMap = oVrf.RibFilter.Ipv6.Ospfv3.RouteMap
							}
						}
					}
				}
				if oVrf.Bgp != nil {
					nestedVrf.Bgp = &VrfBgp{}
					if oVrf.Bgp.Misc != nil {
						entry.Misc["VrfBgp"] = oVrf.Bgp.Misc
					}
					if oVrf.Bgp.Enable != nil {
						nestedVrf.Bgp.Enable = util.AsBool(oVrf.Bgp.Enable, nil)
					}
					if oVrf.Bgp.RouterId != nil {
						nestedVrf.Bgp.RouterId = oVrf.Bgp.RouterId
					}
					if oVrf.Bgp.LocalAs != nil {
						nestedVrf.Bgp.LocalAs = oVrf.Bgp.LocalAs
					}
					if oVrf.Bgp.InstallRoute != nil {
						nestedVrf.Bgp.InstallRoute = util.AsBool(oVrf.Bgp.InstallRoute, nil)
					}
					if oVrf.Bgp.EnforceFirstAs != nil {
						nestedVrf.Bgp.EnforceFirstAs = util.AsBool(oVrf.Bgp.EnforceFirstAs, nil)
					}
					if oVrf.Bgp.FastExternalFailover != nil {
						nestedVrf.Bgp.FastExternalFailover = util.AsBool(oVrf.Bgp.FastExternalFailover, nil)
					}
					if oVrf.Bgp.EcmpMultiAs != nil {
						nestedVrf.Bgp.EcmpMultiAs = util.AsBool(oVrf.Bgp.EcmpMultiAs, nil)
					}
					if oVrf.Bgp.DefaultLocalPreference != nil {
						nestedVrf.Bgp.DefaultLocalPreference = oVrf.Bgp.DefaultLocalPreference
					}
					if oVrf.Bgp.GracefulShutdown != nil {
						nestedVrf.Bgp.GracefulShutdown = util.AsBool(oVrf.Bgp.GracefulShutdown, nil)
					}
					if oVrf.Bgp.AlwaysAdvertiseNetworkRoute != nil {
						nestedVrf.Bgp.AlwaysAdvertiseNetworkRoute = util.AsBool(oVrf.Bgp.AlwaysAdvertiseNetworkRoute, nil)
					}
					if oVrf.Bgp.Med != nil {
						nestedVrf.Bgp.Med = &VrfBgpMed{}
						if oVrf.Bgp.Med.Misc != nil {
							entry.Misc["VrfBgpMed"] = oVrf.Bgp.Med.Misc
						}
						if oVrf.Bgp.Med.AlwaysCompareMed != nil {
							nestedVrf.Bgp.Med.AlwaysCompareMed = util.AsBool(oVrf.Bgp.Med.AlwaysCompareMed, nil)
						}
						if oVrf.Bgp.Med.DeterministicMedComparison != nil {
							nestedVrf.Bgp.Med.DeterministicMedComparison = util.AsBool(oVrf.Bgp.Med.DeterministicMedComparison, nil)
						}
					}
					if oVrf.Bgp.GracefulRestart != nil {
						nestedVrf.Bgp.GracefulRestart = &VrfBgpGracefulRestart{}
						if oVrf.Bgp.GracefulRestart.Misc != nil {
							entry.Misc["VrfBgpGracefulRestart"] = oVrf.Bgp.GracefulRestart.Misc
						}
						if oVrf.Bgp.GracefulRestart.Enable != nil {
							nestedVrf.Bgp.GracefulRestart.Enable = util.AsBool(oVrf.Bgp.GracefulRestart.Enable, nil)
						}
						if oVrf.Bgp.GracefulRestart.StaleRouteTime != nil {
							nestedVrf.Bgp.GracefulRestart.StaleRouteTime = oVrf.Bgp.GracefulRestart.StaleRouteTime
						}
						if oVrf.Bgp.GracefulRestart.MaxPeerRestartTime != nil {
							nestedVrf.Bgp.GracefulRestart.MaxPeerRestartTime = oVrf.Bgp.GracefulRestart.MaxPeerRestartTime
						}
						if oVrf.Bgp.GracefulRestart.LocalRestartTime != nil {
							nestedVrf.Bgp.GracefulRestart.LocalRestartTime = oVrf.Bgp.GracefulRestart.LocalRestartTime
						}
					}
					if oVrf.Bgp.GlobalBfd != nil {
						nestedVrf.Bgp.GlobalBfd = &VrfBgpGlobalBfd{}
						if oVrf.Bgp.GlobalBfd.Misc != nil {
							entry.Misc["VrfBgpGlobalBfd"] = oVrf.Bgp.GlobalBfd.Misc
						}
						if oVrf.Bgp.GlobalBfd.Profile != nil {
							nestedVrf.Bgp.GlobalBfd.Profile = oVrf.Bgp.GlobalBfd.Profile
						}
					}
					if oVrf.Bgp.RedistributionProfile != nil {
						nestedVrf.Bgp.RedistributionProfile = &VrfBgpRedistributionProfile{}
						if oVrf.Bgp.RedistributionProfile.Misc != nil {
							entry.Misc["VrfBgpRedistributionProfile"] = oVrf.Bgp.RedistributionProfile.Misc
						}
						if oVrf.Bgp.RedistributionProfile.Ipv4 != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv4 = &VrfBgpRedistributionProfileIpv4{}
							if oVrf.Bgp.RedistributionProfile.Ipv4.Misc != nil {
								entry.Misc["VrfBgpRedistributionProfileIpv4"] = oVrf.Bgp.RedistributionProfile.Ipv4.Misc
							}
							if oVrf.Bgp.RedistributionProfile.Ipv4.Unicast != nil {
								nestedVrf.Bgp.RedistributionProfile.Ipv4.Unicast = oVrf.Bgp.RedistributionProfile.Ipv4.Unicast
							}
						}
						if oVrf.Bgp.RedistributionProfile.Ipv6 != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv6 = &VrfBgpRedistributionProfileIpv6{}
							if oVrf.Bgp.RedistributionProfile.Ipv6.Misc != nil {
								entry.Misc["VrfBgpRedistributionProfileIpv6"] = oVrf.Bgp.RedistributionProfile.Ipv6.Misc
							}
							if oVrf.Bgp.RedistributionProfile.Ipv6.Unicast != nil {
								nestedVrf.Bgp.RedistributionProfile.Ipv6.Unicast = oVrf.Bgp.RedistributionProfile.Ipv6.Unicast
							}
						}
					}
					if oVrf.Bgp.AdvertiseNetwork != nil {
						nestedVrf.Bgp.AdvertiseNetwork = &VrfBgpAdvertiseNetwork{}
						if oVrf.Bgp.AdvertiseNetwork.Misc != nil {
							entry.Misc["VrfBgpAdvertiseNetwork"] = oVrf.Bgp.AdvertiseNetwork.Misc
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv4 != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4 = &VrfBgpAdvertiseNetworkIpv4{}
							if oVrf.Bgp.AdvertiseNetwork.Ipv4.Misc != nil {
								entry.Misc["VrfBgpAdvertiseNetworkIpv4"] = oVrf.Bgp.AdvertiseNetwork.Ipv4.Misc
							}
							if oVrf.Bgp.AdvertiseNetwork.Ipv4.Network != nil {
								nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = []VrfBgpAdvertiseNetworkIpv4Network{}
								for _, oVrfBgpAdvertiseNetworkIpv4Network := range oVrf.Bgp.AdvertiseNetwork.Ipv4.Network {
									nestedVrfBgpAdvertiseNetworkIpv4Network := VrfBgpAdvertiseNetworkIpv4Network{}
									if oVrfBgpAdvertiseNetworkIpv4Network.Misc != nil {
										entry.Misc["VrfBgpAdvertiseNetworkIpv4Network"] = oVrfBgpAdvertiseNetworkIpv4Network.Misc
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Name != "" {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Name = oVrfBgpAdvertiseNetworkIpv4Network.Name
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Unicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Unicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Unicast, nil)
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Multicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Multicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Multicast, nil)
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Backdoor != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Backdoor = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Backdoor, nil)
									}
									nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network, nestedVrfBgpAdvertiseNetworkIpv4Network)
								}
							}
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv6 != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6 = &VrfBgpAdvertiseNetworkIpv6{}
							if oVrf.Bgp.AdvertiseNetwork.Ipv6.Misc != nil {
								entry.Misc["VrfBgpAdvertiseNetworkIpv6"] = oVrf.Bgp.AdvertiseNetwork.Ipv6.Misc
							}
							if oVrf.Bgp.AdvertiseNetwork.Ipv6.Network != nil {
								nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = []VrfBgpAdvertiseNetworkIpv6Network{}
								for _, oVrfBgpAdvertiseNetworkIpv6Network := range oVrf.Bgp.AdvertiseNetwork.Ipv6.Network {
									nestedVrfBgpAdvertiseNetworkIpv6Network := VrfBgpAdvertiseNetworkIpv6Network{}
									if oVrfBgpAdvertiseNetworkIpv6Network.Misc != nil {
										entry.Misc["VrfBgpAdvertiseNetworkIpv6Network"] = oVrfBgpAdvertiseNetworkIpv6Network.Misc
									}
									if oVrfBgpAdvertiseNetworkIpv6Network.Name != "" {
										nestedVrfBgpAdvertiseNetworkIpv6Network.Name = oVrfBgpAdvertiseNetworkIpv6Network.Name
									}
									if oVrfBgpAdvertiseNetworkIpv6Network.Unicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv6Network.Unicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv6Network.Unicast, nil)
									}
									nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network, nestedVrfBgpAdvertiseNetworkIpv6Network)
								}
							}
						}
					}
					if oVrf.Bgp.PeerGroup != nil {
						nestedVrf.Bgp.PeerGroup = []VrfBgpPeerGroup{}
						for _, oVrfBgpPeerGroup := range oVrf.Bgp.PeerGroup {
							nestedVrfBgpPeerGroup := VrfBgpPeerGroup{}
							if oVrfBgpPeerGroup.Misc != nil {
								entry.Misc["VrfBgpPeerGroup"] = oVrfBgpPeerGroup.Misc
							}
							if oVrfBgpPeerGroup.Name != "" {
								nestedVrfBgpPeerGroup.Name = oVrfBgpPeerGroup.Name
							}
							if oVrfBgpPeerGroup.Enable != nil {
								nestedVrfBgpPeerGroup.Enable = util.AsBool(oVrfBgpPeerGroup.Enable, nil)
							}
							if oVrfBgpPeerGroup.Type != nil {
								nestedVrfBgpPeerGroup.Type = &VrfBgpPeerGroupType{}
								if oVrfBgpPeerGroup.Type.Misc != nil {
									entry.Misc["VrfBgpPeerGroupType"] = oVrfBgpPeerGroup.Type.Misc
								}
								if oVrfBgpPeerGroup.Type.Ibgp != nil {
									nestedVrfBgpPeerGroup.Type.Ibgp = &VrfBgpPeerGroupTypeIbgp{}
									if oVrfBgpPeerGroup.Type.Ibgp.Misc != nil {
										entry.Misc["VrfBgpPeerGroupTypeIbgp"] = oVrfBgpPeerGroup.Type.Ibgp.Misc
									}
								}
								if oVrfBgpPeerGroup.Type.Ebgp != nil {
									nestedVrfBgpPeerGroup.Type.Ebgp = &VrfBgpPeerGroupTypeEbgp{}
									if oVrfBgpPeerGroup.Type.Ebgp.Misc != nil {
										entry.Misc["VrfBgpPeerGroupTypeEbgp"] = oVrfBgpPeerGroup.Type.Ebgp.Misc
									}
								}
							}
							if oVrfBgpPeerGroup.AddressFamily != nil {
								nestedVrfBgpPeerGroup.AddressFamily = &VrfBgpPeerGroupAddressFamily{}
								if oVrfBgpPeerGroup.AddressFamily.Misc != nil {
									entry.Misc["VrfBgpPeerGroupAddressFamily"] = oVrfBgpPeerGroup.AddressFamily.Misc
								}
								if oVrfBgpPeerGroup.AddressFamily.Ipv4 != nil {
									nestedVrfBgpPeerGroup.AddressFamily.Ipv4 = oVrfBgpPeerGroup.AddressFamily.Ipv4
								}
								if oVrfBgpPeerGroup.AddressFamily.Ipv6 != nil {
									nestedVrfBgpPeerGroup.AddressFamily.Ipv6 = oVrfBgpPeerGroup.AddressFamily.Ipv6
								}
							}
							if oVrfBgpPeerGroup.FilteringProfile != nil {
								nestedVrfBgpPeerGroup.FilteringProfile = &VrfBgpPeerGroupFilteringProfile{}
								if oVrfBgpPeerGroup.FilteringProfile.Misc != nil {
									entry.Misc["VrfBgpPeerGroupFilteringProfile"] = oVrfBgpPeerGroup.FilteringProfile.Misc
								}
								if oVrfBgpPeerGroup.FilteringProfile.Ipv4 != nil {
									nestedVrfBgpPeerGroup.FilteringProfile.Ipv4 = oVrfBgpPeerGroup.FilteringProfile.Ipv4
								}
								if oVrfBgpPeerGroup.FilteringProfile.Ipv6 != nil {
									nestedVrfBgpPeerGroup.FilteringProfile.Ipv6 = oVrfBgpPeerGroup.FilteringProfile.Ipv6
								}
							}
							if oVrfBgpPeerGroup.ConnectionOptions != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions = &VrfBgpPeerGroupConnectionOptions{}
								if oVrfBgpPeerGroup.ConnectionOptions.Misc != nil {
									entry.Misc["VrfBgpPeerGroupConnectionOptions"] = oVrfBgpPeerGroup.ConnectionOptions.Misc
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Timers != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Timers = oVrfBgpPeerGroup.ConnectionOptions.Timers
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Multihop != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Multihop = oVrfBgpPeerGroup.ConnectionOptions.Multihop
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Authentication != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Authentication = oVrfBgpPeerGroup.ConnectionOptions.Authentication
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Dampening != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Dampening = oVrfBgpPeerGroup.ConnectionOptions.Dampening
								}
							}
							if oVrfBgpPeerGroup.Peer != nil {
								nestedVrfBgpPeerGroup.Peer = []VrfBgpPeerGroupPeer{}
								for _, oVrfBgpPeerGroupPeer := range oVrfBgpPeerGroup.Peer {
									nestedVrfBgpPeerGroupPeer := VrfBgpPeerGroupPeer{}
									if oVrfBgpPeerGroupPeer.Misc != nil {
										entry.Misc["VrfBgpPeerGroupPeer"] = oVrfBgpPeerGroupPeer.Misc
									}
									if oVrfBgpPeerGroupPeer.Name != "" {
										nestedVrfBgpPeerGroupPeer.Name = oVrfBgpPeerGroupPeer.Name
									}
									if oVrfBgpPeerGroupPeer.Enable != nil {
										nestedVrfBgpPeerGroupPeer.Enable = util.AsBool(oVrfBgpPeerGroupPeer.Enable, nil)
									}
									if oVrfBgpPeerGroupPeer.Passive != nil {
										nestedVrfBgpPeerGroupPeer.Passive = util.AsBool(oVrfBgpPeerGroupPeer.Passive, nil)
									}
									if oVrfBgpPeerGroupPeer.PeerAs != nil {
										nestedVrfBgpPeerGroupPeer.PeerAs = oVrfBgpPeerGroupPeer.PeerAs
									}
									if oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
										nestedVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.AsBool(oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
									}
									if oVrfBgpPeerGroupPeer.Inherit != nil {
										nestedVrfBgpPeerGroupPeer.Inherit = &VrfBgpPeerGroupPeerInherit{}
										if oVrfBgpPeerGroupPeer.Inherit.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerInherit"] = oVrfBgpPeerGroupPeer.Inherit.Misc
										}
										if oVrfBgpPeerGroupPeer.Inherit.Yes != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.Yes = &VrfBgpPeerGroupPeerInheritYes{}
											if oVrfBgpPeerGroupPeer.Inherit.Yes.Misc != nil {
												entry.Misc["VrfBgpPeerGroupPeerInheritYes"] = oVrfBgpPeerGroupPeer.Inherit.Yes.Misc
											}
										}
										if oVrfBgpPeerGroupPeer.Inherit.No != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No = &VrfBgpPeerGroupPeerInheritNo{}
											if oVrfBgpPeerGroupPeer.Inherit.No.Misc != nil {
												entry.Misc["VrfBgpPeerGroupPeerInheritNo"] = oVrfBgpPeerGroupPeer.Inherit.No.Misc
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily = &VrfBgpPeerGroupPeerInheritNoAddressFamily{}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc != nil {
													entry.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"] = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6
												}
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile = &VrfBgpPeerGroupPeerInheritNoFilteringProfile{}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc != nil {
													entry.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"] = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6
												}
											}
										}
									}
									if oVrfBgpPeerGroupPeer.LocalAddress != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress = &VrfBgpPeerGroupPeerLocalAddress{}
										if oVrfBgpPeerGroupPeer.LocalAddress.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerLocalAddress"] = oVrfBgpPeerGroupPeer.LocalAddress.Misc
										}
										if oVrfBgpPeerGroupPeer.LocalAddress.Interface != nil {
											nestedVrfBgpPeerGroupPeer.LocalAddress.Interface = oVrfBgpPeerGroupPeer.LocalAddress.Interface
										}
										if oVrfBgpPeerGroupPeer.LocalAddress.Ip != nil {
											nestedVrfBgpPeerGroupPeer.LocalAddress.Ip = oVrfBgpPeerGroupPeer.LocalAddress.Ip
										}
									}
									if oVrfBgpPeerGroupPeer.PeerAddress != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress = &VrfBgpPeerGroupPeerPeerAddress{}
										if oVrfBgpPeerGroupPeer.PeerAddress.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerPeerAddress"] = oVrfBgpPeerGroupPeer.PeerAddress.Misc
										}
										if oVrfBgpPeerGroupPeer.PeerAddress.Ip != nil {
											nestedVrfBgpPeerGroupPeer.PeerAddress.Ip = oVrfBgpPeerGroupPeer.PeerAddress.Ip
										}
										if oVrfBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
											nestedVrfBgpPeerGroupPeer.PeerAddress.Fqdn = oVrfBgpPeerGroupPeer.PeerAddress.Fqdn
										}
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions = &VrfBgpPeerGroupPeerConnectionOptions{}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerConnectionOptions"] = oVrfBgpPeerGroupPeer.ConnectionOptions.Misc
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Timers != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Timers = oVrfBgpPeerGroupPeer.ConnectionOptions.Timers
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Multihop = oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Authentication = oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Dampening = oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening
										}
									}
									if oVrfBgpPeerGroupPeer.Bfd != nil {
										nestedVrfBgpPeerGroupPeer.Bfd = &VrfBgpPeerGroupPeerBfd{}
										if oVrfBgpPeerGroupPeer.Bfd.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerBfd"] = oVrfBgpPeerGroupPeer.Bfd.Misc
										}
										if oVrfBgpPeerGroupPeer.Bfd.Profile != nil {
											nestedVrfBgpPeerGroupPeer.Bfd.Profile = oVrfBgpPeerGroupPeer.Bfd.Profile
										}
									}
									nestedVrfBgpPeerGroup.Peer = append(nestedVrfBgpPeerGroup.Peer, nestedVrfBgpPeerGroupPeer)
								}
							}
							nestedVrf.Bgp.PeerGroup = append(nestedVrf.Bgp.PeerGroup, nestedVrfBgpPeerGroup)
						}
					}
					if oVrf.Bgp.AggregateRoutes != nil {
						nestedVrf.Bgp.AggregateRoutes = []VrfBgpAggregateRoutes{}
						for _, oVrfBgpAggregateRoutes := range oVrf.Bgp.AggregateRoutes {
							nestedVrfBgpAggregateRoutes := VrfBgpAggregateRoutes{}
							if oVrfBgpAggregateRoutes.Misc != nil {
								entry.Misc["VrfBgpAggregateRoutes"] = oVrfBgpAggregateRoutes.Misc
							}
							if oVrfBgpAggregateRoutes.Name != "" {
								nestedVrfBgpAggregateRoutes.Name = oVrfBgpAggregateRoutes.Name
							}
							if oVrfBgpAggregateRoutes.Description != nil {
								nestedVrfBgpAggregateRoutes.Description = oVrfBgpAggregateRoutes.Description
							}
							if oVrfBgpAggregateRoutes.Enable != nil {
								nestedVrfBgpAggregateRoutes.Enable = util.AsBool(oVrfBgpAggregateRoutes.Enable, nil)
							}
							if oVrfBgpAggregateRoutes.SummaryOnly != nil {
								nestedVrfBgpAggregateRoutes.SummaryOnly = util.AsBool(oVrfBgpAggregateRoutes.SummaryOnly, nil)
							}
							if oVrfBgpAggregateRoutes.AsSet != nil {
								nestedVrfBgpAggregateRoutes.AsSet = util.AsBool(oVrfBgpAggregateRoutes.AsSet, nil)
							}
							if oVrfBgpAggregateRoutes.SameMed != nil {
								nestedVrfBgpAggregateRoutes.SameMed = util.AsBool(oVrfBgpAggregateRoutes.SameMed, nil)
							}
							if oVrfBgpAggregateRoutes.Type != nil {
								nestedVrfBgpAggregateRoutes.Type = &VrfBgpAggregateRoutesType{}
								if oVrfBgpAggregateRoutes.Type.Misc != nil {
									entry.Misc["VrfBgpAggregateRoutesType"] = oVrfBgpAggregateRoutes.Type.Misc
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4 != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4 = &VrfBgpAggregateRoutesTypeIpv4{}
									if oVrfBgpAggregateRoutes.Type.Ipv4.Misc != nil {
										entry.Misc["VrfBgpAggregateRoutesTypeIpv4"] = oVrfBgpAggregateRoutes.Type.Ipv4.Misc
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap
									}
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6 != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6 = &VrfBgpAggregateRoutesTypeIpv6{}
									if oVrfBgpAggregateRoutes.Type.Ipv6.Misc != nil {
										entry.Misc["VrfBgpAggregateRoutesTypeIpv6"] = oVrfBgpAggregateRoutes.Type.Ipv6.Misc
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap
									}
								}
							}
							nestedVrf.Bgp.AggregateRoutes = append(nestedVrf.Bgp.AggregateRoutes, nestedVrfBgpAggregateRoutes)
						}
					}
				}
				if oVrf.RoutingTable != nil {
					nestedVrf.RoutingTable = &VrfRoutingTable{}
					if oVrf.RoutingTable.Misc != nil {
						entry.Misc["VrfRoutingTable"] = oVrf.RoutingTable.Misc
					}
					if oVrf.RoutingTable.Ip != nil {
						nestedVrf.RoutingTable.Ip = &VrfRoutingTableIp{}
						if oVrf.RoutingTable.Ip.Misc != nil {
							entry.Misc["VrfRoutingTableIp"] = oVrf.RoutingTable.Ip.Misc
						}
						if oVrf.RoutingTable.Ip.StaticRoute != nil {
							nestedVrf.RoutingTable.Ip.StaticRoute = []VrfRoutingTableIpStaticRoute{}
							for _, oVrfRoutingTableIpStaticRoute := range oVrf.RoutingTable.Ip.StaticRoute {
								nestedVrfRoutingTableIpStaticRoute := VrfRoutingTableIpStaticRoute{}
								if oVrfRoutingTableIpStaticRoute.Misc != nil {
									entry.Misc["VrfRoutingTableIpStaticRoute"] = oVrfRoutingTableIpStaticRoute.Misc
								}
								if oVrfRoutingTableIpStaticRoute.Name != "" {
									nestedVrfRoutingTableIpStaticRoute.Name = oVrfRoutingTableIpStaticRoute.Name
								}
								if oVrfRoutingTableIpStaticRoute.Destination != nil {
									nestedVrfRoutingTableIpStaticRoute.Destination = oVrfRoutingTableIpStaticRoute.Destination
								}
								if oVrfRoutingTableIpStaticRoute.Interface != nil {
									nestedVrfRoutingTableIpStaticRoute.Interface = oVrfRoutingTableIpStaticRoute.Interface
								}
								if oVrfRoutingTableIpStaticRoute.AdminDist != nil {
									nestedVrfRoutingTableIpStaticRoute.AdminDist = oVrfRoutingTableIpStaticRoute.AdminDist
								}
								if oVrfRoutingTableIpStaticRoute.Metric != nil {
									nestedVrfRoutingTableIpStaticRoute.Metric = oVrfRoutingTableIpStaticRoute.Metric
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop = &VrfRoutingTableIpStaticRouteNexthop{}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRouteNexthop"] = oVrfRoutingTableIpStaticRoute.Nexthop.Misc
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Discard != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard = &VrfRoutingTableIpStaticRouteNexthopDiscard{}
										if oVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc != nil {
											entry.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"] = oVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc
										}
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.IpAddress = oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.NextLr != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.NextLr = oVrfRoutingTableIpStaticRoute.Nexthop.NextLr
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn
									}
								}
								if oVrfRoutingTableIpStaticRoute.Bfd != nil {
									nestedVrfRoutingTableIpStaticRoute.Bfd = &VrfRoutingTableIpStaticRouteBfd{}
									if oVrfRoutingTableIpStaticRoute.Bfd.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRouteBfd"] = oVrfRoutingTableIpStaticRoute.Bfd.Misc
									}
									if oVrfRoutingTableIpStaticRoute.Bfd.Profile != nil {
										nestedVrfRoutingTableIpStaticRoute.Bfd.Profile = oVrfRoutingTableIpStaticRoute.Bfd.Profile
									}
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor = &VrfRoutingTableIpStaticRoutePathMonitor{}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRoutePathMonitor"] = oVrfRoutingTableIpStaticRoute.PathMonitor.Misc
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.Enable = util.AsBool(oVrfRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
										for _, oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc != nil {
												entry.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"] = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
											}
											nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
										}
									}
								}
								nestedVrf.RoutingTable.Ip.StaticRoute = append(nestedVrf.RoutingTable.Ip.StaticRoute, nestedVrfRoutingTableIpStaticRoute)
							}
						}
					}
					if oVrf.RoutingTable.Ipv6 != nil {
						nestedVrf.RoutingTable.Ipv6 = &VrfRoutingTableIpv6{}
						if oVrf.RoutingTable.Ipv6.Misc != nil {
							entry.Misc["VrfRoutingTableIpv6"] = oVrf.RoutingTable.Ipv6.Misc
						}
						if oVrf.RoutingTable.Ipv6.StaticRoute != nil {
							nestedVrf.RoutingTable.Ipv6.StaticRoute = []VrfRoutingTableIpv6StaticRoute{}
							for _, oVrfRoutingTableIpv6StaticRoute := range oVrf.RoutingTable.Ipv6.StaticRoute {
								nestedVrfRoutingTableIpv6StaticRoute := VrfRoutingTableIpv6StaticRoute{}
								if oVrfRoutingTableIpv6StaticRoute.Misc != nil {
									entry.Misc["VrfRoutingTableIpv6StaticRoute"] = oVrfRoutingTableIpv6StaticRoute.Misc
								}
								if oVrfRoutingTableIpv6StaticRoute.Name != "" {
									nestedVrfRoutingTableIpv6StaticRoute.Name = oVrfRoutingTableIpv6StaticRoute.Name
								}
								if oVrfRoutingTableIpv6StaticRoute.Destination != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Destination = oVrfRoutingTableIpv6StaticRoute.Destination
								}
								if oVrfRoutingTableIpv6StaticRoute.Interface != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Interface = oVrfRoutingTableIpv6StaticRoute.Interface
								}
								if oVrfRoutingTableIpv6StaticRoute.AdminDist != nil {
									nestedVrfRoutingTableIpv6StaticRoute.AdminDist = oVrfRoutingTableIpv6StaticRoute.AdminDist
								}
								if oVrfRoutingTableIpv6StaticRoute.Metric != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Metric = oVrfRoutingTableIpv6StaticRoute.Metric
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop = &VrfRoutingTableIpv6StaticRouteNexthop{}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRouteNexthop"] = oVrfRoutingTableIpv6StaticRoute.Nexthop.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard = &VrfRoutingTableIpv6StaticRouteNexthopDiscard{}
										if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc != nil {
											entry.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"] = oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc
										}
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr = oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.Bfd != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd = &VrfRoutingTableIpv6StaticRouteBfd{}
									if oVrfRoutingTableIpv6StaticRoute.Bfd.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRouteBfd"] = oVrfRoutingTableIpv6StaticRoute.Bfd.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Bfd.Profile = oVrfRoutingTableIpv6StaticRoute.Bfd.Profile
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor = &VrfRoutingTableIpv6StaticRoutePathMonitor{}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"] = oVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.AsBool(oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
										for _, oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc != nil {
												entry.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"] = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
											}
											nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
										}
									}
								}
								nestedVrf.RoutingTable.Ipv6.StaticRoute = append(nestedVrf.RoutingTable.Ipv6.StaticRoute, nestedVrfRoutingTableIpv6StaticRoute)
							}
						}
					}
				}
				if oVrf.Ospf != nil {
					nestedVrf.Ospf = &VrfOspf{}
					if oVrf.Ospf.Misc != nil {
						entry.Misc["VrfOspf"] = oVrf.Ospf.Misc
					}
					if oVrf.Ospf.RouterId != nil {
						nestedVrf.Ospf.RouterId = oVrf.Ospf.RouterId
					}
					if oVrf.Ospf.Enable != nil {
						nestedVrf.Ospf.Enable = util.AsBool(oVrf.Ospf.Enable, nil)
					}
					if oVrf.Ospf.Rfc1583 != nil {
						nestedVrf.Ospf.Rfc1583 = util.AsBool(oVrf.Ospf.Rfc1583, nil)
					}
					if oVrf.Ospf.SpfTimer != nil {
						nestedVrf.Ospf.SpfTimer = oVrf.Ospf.SpfTimer
					}
					if oVrf.Ospf.GlobalIfTimer != nil {
						nestedVrf.Ospf.GlobalIfTimer = oVrf.Ospf.GlobalIfTimer
					}
					if oVrf.Ospf.RedistributionProfile != nil {
						nestedVrf.Ospf.RedistributionProfile = oVrf.Ospf.RedistributionProfile
					}
					if oVrf.Ospf.GlobalBfd != nil {
						nestedVrf.Ospf.GlobalBfd = &VrfOspfGlobalBfd{}
						if oVrf.Ospf.GlobalBfd.Misc != nil {
							entry.Misc["VrfOspfGlobalBfd"] = oVrf.Ospf.GlobalBfd.Misc
						}
						if oVrf.Ospf.GlobalBfd.Profile != nil {
							nestedVrf.Ospf.GlobalBfd.Profile = oVrf.Ospf.GlobalBfd.Profile
						}
					}
					if oVrf.Ospf.GracefulRestart != nil {
						nestedVrf.Ospf.GracefulRestart = &VrfOspfGracefulRestart{}
						if oVrf.Ospf.GracefulRestart.Misc != nil {
							entry.Misc["VrfOspfGracefulRestart"] = oVrf.Ospf.GracefulRestart.Misc
						}
						if oVrf.Ospf.GracefulRestart.Enable != nil {
							nestedVrf.Ospf.GracefulRestart.Enable = util.AsBool(oVrf.Ospf.GracefulRestart.Enable, nil)
						}
						if oVrf.Ospf.GracefulRestart.GracePeriod != nil {
							nestedVrf.Ospf.GracefulRestart.GracePeriod = oVrf.Ospf.GracefulRestart.GracePeriod
						}
						if oVrf.Ospf.GracefulRestart.HelperEnable != nil {
							nestedVrf.Ospf.GracefulRestart.HelperEnable = util.AsBool(oVrf.Ospf.GracefulRestart.HelperEnable, nil)
						}
						if oVrf.Ospf.GracefulRestart.StrictLSAChecking != nil {
							nestedVrf.Ospf.GracefulRestart.StrictLSAChecking = util.AsBool(oVrf.Ospf.GracefulRestart.StrictLSAChecking, nil)
						}
						if oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
							nestedVrf.Ospf.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime
						}
					}
					if oVrf.Ospf.Area != nil {
						nestedVrf.Ospf.Area = []VrfOspfArea{}
						for _, oVrfOspfArea := range oVrf.Ospf.Area {
							nestedVrfOspfArea := VrfOspfArea{}
							if oVrfOspfArea.Misc != nil {
								entry.Misc["VrfOspfArea"] = oVrfOspfArea.Misc
							}
							if oVrfOspfArea.Name != "" {
								nestedVrfOspfArea.Name = oVrfOspfArea.Name
							}
							if oVrfOspfArea.Authentication != nil {
								nestedVrfOspfArea.Authentication = oVrfOspfArea.Authentication
							}
							if oVrfOspfArea.Type != nil {
								nestedVrfOspfArea.Type = &VrfOspfAreaType{}
								if oVrfOspfArea.Type.Misc != nil {
									entry.Misc["VrfOspfAreaType"] = oVrfOspfArea.Type.Misc
								}
								if oVrfOspfArea.Type.Normal != nil {
									nestedVrfOspfArea.Type.Normal = &VrfOspfAreaTypeNormal{}
									if oVrfOspfArea.Type.Normal.Misc != nil {
										entry.Misc["VrfOspfAreaTypeNormal"] = oVrfOspfArea.Type.Normal.Misc
									}
									if oVrfOspfArea.Type.Normal.Abr != nil {
										nestedVrfOspfArea.Type.Normal.Abr = &VrfOspfAreaTypeNormalAbr{}
										if oVrfOspfArea.Type.Normal.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNormalAbr"] = oVrfOspfArea.Type.Normal.Abr.Misc
										}
										if oVrfOspfArea.Type.Normal.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.ImportList = oVrfOspfArea.Type.Normal.Abr.ImportList
										}
										if oVrfOspfArea.Type.Normal.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.ExportList = oVrfOspfArea.Type.Normal.Abr.ExportList
										}
										if oVrfOspfArea.Type.Normal.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.InboundFilterList = oVrfOspfArea.Type.Normal.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Normal.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.OutboundFilterList = oVrfOspfArea.Type.Normal.Abr.OutboundFilterList
										}
									}
								}
								if oVrfOspfArea.Type.Stub != nil {
									nestedVrfOspfArea.Type.Stub = &VrfOspfAreaTypeStub{}
									if oVrfOspfArea.Type.Stub.Misc != nil {
										entry.Misc["VrfOspfAreaTypeStub"] = oVrfOspfArea.Type.Stub.Misc
									}
									if oVrfOspfArea.Type.Stub.NoSummary != nil {
										nestedVrfOspfArea.Type.Stub.NoSummary = util.AsBool(oVrfOspfArea.Type.Stub.NoSummary, nil)
									}
									if oVrfOspfArea.Type.Stub.Abr != nil {
										nestedVrfOspfArea.Type.Stub.Abr = &VrfOspfAreaTypeStubAbr{}
										if oVrfOspfArea.Type.Stub.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeStubAbr"] = oVrfOspfArea.Type.Stub.Abr.Misc
										}
										if oVrfOspfArea.Type.Stub.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.ImportList = oVrfOspfArea.Type.Stub.Abr.ImportList
										}
										if oVrfOspfArea.Type.Stub.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.ExportList = oVrfOspfArea.Type.Stub.Abr.ExportList
										}
										if oVrfOspfArea.Type.Stub.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.InboundFilterList = oVrfOspfArea.Type.Stub.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Stub.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.OutboundFilterList = oVrfOspfArea.Type.Stub.Abr.OutboundFilterList
										}
									}
									if oVrfOspfArea.Type.Stub.DefaultRouteMetric != nil {
										nestedVrfOspfArea.Type.Stub.DefaultRouteMetric = oVrfOspfArea.Type.Stub.DefaultRouteMetric
									}
								}
								if oVrfOspfArea.Type.Nssa != nil {
									nestedVrfOspfArea.Type.Nssa = &VrfOspfAreaTypeNssa{}
									if oVrfOspfArea.Type.Nssa.Misc != nil {
										entry.Misc["VrfOspfAreaTypeNssa"] = oVrfOspfArea.Type.Nssa.Misc
									}
									if oVrfOspfArea.Type.Nssa.NoSummary != nil {
										nestedVrfOspfArea.Type.Nssa.NoSummary = util.AsBool(oVrfOspfArea.Type.Nssa.NoSummary, nil)
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate = &VrfOspfAreaTypeNssaDefaultInformationOriginate{}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"] = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc
										}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric != nil {
											nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric
										}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
											nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType
										}
									}
									if oVrfOspfArea.Type.Nssa.Abr != nil {
										nestedVrfOspfArea.Type.Nssa.Abr = &VrfOspfAreaTypeNssaAbr{}
										if oVrfOspfArea.Type.Nssa.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNssaAbr"] = oVrfOspfArea.Type.Nssa.Abr.Misc
										}
										if oVrfOspfArea.Type.Nssa.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.ImportList = oVrfOspfArea.Type.Nssa.Abr.ImportList
										}
										if oVrfOspfArea.Type.Nssa.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.ExportList = oVrfOspfArea.Type.Nssa.Abr.ExportList
										}
										if oVrfOspfArea.Type.Nssa.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.InboundFilterList = oVrfOspfArea.Type.Nssa.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.OutboundFilterList = oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList
										}
										if oVrfOspfArea.Type.Nssa.Abr.NssaExtRange != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = []VrfOspfAreaTypeNssaAbrNssaExtRange{}
											for _, oVrfOspfAreaTypeNssaAbrNssaExtRange := range oVrfOspfArea.Type.Nssa.Abr.NssaExtRange {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange := VrfOspfAreaTypeNssaAbrNssaExtRange{}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Misc != nil {
													entry.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"] = oVrfOspfAreaTypeNssaAbrNssaExtRange.Misc
												}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Name != "" {
													nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Name = oVrfOspfAreaTypeNssaAbrNssaExtRange.Name
												}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise != nil {
													nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise = util.AsBool(oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise, nil)
												}
												nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfAreaTypeNssaAbrNssaExtRange)
											}
										}
									}
								}
							}
							if oVrfOspfArea.Range != nil {
								nestedVrfOspfArea.Range = []VrfOspfAreaRange{}
								for _, oVrfOspfAreaRange := range oVrfOspfArea.Range {
									nestedVrfOspfAreaRange := VrfOspfAreaRange{}
									if oVrfOspfAreaRange.Misc != nil {
										entry.Misc["VrfOspfAreaRange"] = oVrfOspfAreaRange.Misc
									}
									if oVrfOspfAreaRange.Name != "" {
										nestedVrfOspfAreaRange.Name = oVrfOspfAreaRange.Name
									}
									if oVrfOspfAreaRange.Advertise != nil {
										nestedVrfOspfAreaRange.Advertise = util.AsBool(oVrfOspfAreaRange.Advertise, nil)
									}
									nestedVrfOspfArea.Range = append(nestedVrfOspfArea.Range, nestedVrfOspfAreaRange)
								}
							}
							if oVrfOspfArea.Interface != nil {
								nestedVrfOspfArea.Interface = []VrfOspfAreaInterface{}
								for _, oVrfOspfAreaInterface := range oVrfOspfArea.Interface {
									nestedVrfOspfAreaInterface := VrfOspfAreaInterface{}
									if oVrfOspfAreaInterface.Misc != nil {
										entry.Misc["VrfOspfAreaInterface"] = oVrfOspfAreaInterface.Misc
									}
									if oVrfOspfAreaInterface.Name != "" {
										nestedVrfOspfAreaInterface.Name = oVrfOspfAreaInterface.Name
									}
									if oVrfOspfAreaInterface.Enable != nil {
										nestedVrfOspfAreaInterface.Enable = util.AsBool(oVrfOspfAreaInterface.Enable, nil)
									}
									if oVrfOspfAreaInterface.MtuIgnore != nil {
										nestedVrfOspfAreaInterface.MtuIgnore = util.AsBool(oVrfOspfAreaInterface.MtuIgnore, nil)
									}
									if oVrfOspfAreaInterface.Passive != nil {
										nestedVrfOspfAreaInterface.Passive = util.AsBool(oVrfOspfAreaInterface.Passive, nil)
									}
									if oVrfOspfAreaInterface.Priority != nil {
										nestedVrfOspfAreaInterface.Priority = oVrfOspfAreaInterface.Priority
									}
									if oVrfOspfAreaInterface.Metric != nil {
										nestedVrfOspfAreaInterface.Metric = oVrfOspfAreaInterface.Metric
									}
									if oVrfOspfAreaInterface.Authentication != nil {
										nestedVrfOspfAreaInterface.Authentication = oVrfOspfAreaInterface.Authentication
									}
									if oVrfOspfAreaInterface.Timing != nil {
										nestedVrfOspfAreaInterface.Timing = oVrfOspfAreaInterface.Timing
									}
									if oVrfOspfAreaInterface.LinkType != nil {
										nestedVrfOspfAreaInterface.LinkType = &VrfOspfAreaInterfaceLinkType{}
										if oVrfOspfAreaInterface.LinkType.Misc != nil {
											entry.Misc["VrfOspfAreaInterfaceLinkType"] = oVrfOspfAreaInterface.LinkType.Misc
										}
										if oVrfOspfAreaInterface.LinkType.Broadcast != nil {
											nestedVrfOspfAreaInterface.LinkType.Broadcast = &VrfOspfAreaInterfaceLinkTypeBroadcast{}
											if oVrfOspfAreaInterface.LinkType.Broadcast.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"] = oVrfOspfAreaInterface.LinkType.Broadcast.Misc
											}
										}
										if oVrfOspfAreaInterface.LinkType.P2p != nil {
											nestedVrfOspfAreaInterface.LinkType.P2p = &VrfOspfAreaInterfaceLinkTypeP2p{}
											if oVrfOspfAreaInterface.LinkType.P2p.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeP2p"] = oVrfOspfAreaInterface.LinkType.P2p.Misc
											}
										}
										if oVrfOspfAreaInterface.LinkType.P2mp != nil {
											nestedVrfOspfAreaInterface.LinkType.P2mp = &VrfOspfAreaInterfaceLinkTypeP2mp{}
											if oVrfOspfAreaInterface.LinkType.P2mp.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"] = oVrfOspfAreaInterface.LinkType.P2mp.Misc
											}
											if oVrfOspfAreaInterface.LinkType.P2mp.Neighbor != nil {
												nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{}
												for _, oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfAreaInterface.LinkType.P2mp.Neighbor {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc != nil {
														entry.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"] = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc
													}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
														nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name
													}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
														nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority
													}
													nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor)
												}
											}
										}
									}
									if oVrfOspfAreaInterface.Bfd != nil {
										nestedVrfOspfAreaInterface.Bfd = &VrfOspfAreaInterfaceBfd{}
										if oVrfOspfAreaInterface.Bfd.Misc != nil {
											entry.Misc["VrfOspfAreaInterfaceBfd"] = oVrfOspfAreaInterface.Bfd.Misc
										}
										if oVrfOspfAreaInterface.Bfd.Profile != nil {
											nestedVrfOspfAreaInterface.Bfd.Profile = oVrfOspfAreaInterface.Bfd.Profile
										}
									}
									nestedVrfOspfArea.Interface = append(nestedVrfOspfArea.Interface, nestedVrfOspfAreaInterface)
								}
							}
							if oVrfOspfArea.VirtualLink != nil {
								nestedVrfOspfArea.VirtualLink = []VrfOspfAreaVirtualLink{}
								for _, oVrfOspfAreaVirtualLink := range oVrfOspfArea.VirtualLink {
									nestedVrfOspfAreaVirtualLink := VrfOspfAreaVirtualLink{}
									if oVrfOspfAreaVirtualLink.Misc != nil {
										entry.Misc["VrfOspfAreaVirtualLink"] = oVrfOspfAreaVirtualLink.Misc
									}
									if oVrfOspfAreaVirtualLink.Name != "" {
										nestedVrfOspfAreaVirtualLink.Name = oVrfOspfAreaVirtualLink.Name
									}
									if oVrfOspfAreaVirtualLink.NeighborId != nil {
										nestedVrfOspfAreaVirtualLink.NeighborId = oVrfOspfAreaVirtualLink.NeighborId
									}
									if oVrfOspfAreaVirtualLink.TransitAreaId != nil {
										nestedVrfOspfAreaVirtualLink.TransitAreaId = oVrfOspfAreaVirtualLink.TransitAreaId
									}
									if oVrfOspfAreaVirtualLink.Enable != nil {
										nestedVrfOspfAreaVirtualLink.Enable = util.AsBool(oVrfOspfAreaVirtualLink.Enable, nil)
									}
									if oVrfOspfAreaVirtualLink.InstanceId != nil {
										nestedVrfOspfAreaVirtualLink.InstanceId = oVrfOspfAreaVirtualLink.InstanceId
									}
									if oVrfOspfAreaVirtualLink.Timing != nil {
										nestedVrfOspfAreaVirtualLink.Timing = oVrfOspfAreaVirtualLink.Timing
									}
									if oVrfOspfAreaVirtualLink.Authentication != nil {
										nestedVrfOspfAreaVirtualLink.Authentication = oVrfOspfAreaVirtualLink.Authentication
									}
									if oVrfOspfAreaVirtualLink.Bfd != nil {
										nestedVrfOspfAreaVirtualLink.Bfd = &VrfOspfAreaVirtualLinkBfd{}
										if oVrfOspfAreaVirtualLink.Bfd.Misc != nil {
											entry.Misc["VrfOspfAreaVirtualLinkBfd"] = oVrfOspfAreaVirtualLink.Bfd.Misc
										}
										if oVrfOspfAreaVirtualLink.Bfd.Profile != nil {
											nestedVrfOspfAreaVirtualLink.Bfd.Profile = oVrfOspfAreaVirtualLink.Bfd.Profile
										}
									}
									nestedVrfOspfArea.VirtualLink = append(nestedVrfOspfArea.VirtualLink, nestedVrfOspfAreaVirtualLink)
								}
							}
							nestedVrf.Ospf.Area = append(nestedVrf.Ospf.Area, nestedVrfOspfArea)
						}
					}
				}
				if oVrf.Ospfv3 != nil {
					nestedVrf.Ospfv3 = &VrfOspfv3{}
					if oVrf.Ospfv3.Misc != nil {
						entry.Misc["VrfOspfv3"] = oVrf.Ospfv3.Misc
					}
					if oVrf.Ospfv3.Enable != nil {
						nestedVrf.Ospfv3.Enable = util.AsBool(oVrf.Ospfv3.Enable, nil)
					}
					if oVrf.Ospfv3.RouterId != nil {
						nestedVrf.Ospfv3.RouterId = oVrf.Ospfv3.RouterId
					}
					if oVrf.Ospfv3.DisableTransitTraffic != nil {
						nestedVrf.Ospfv3.DisableTransitTraffic = util.AsBool(oVrf.Ospfv3.DisableTransitTraffic, nil)
					}
					if oVrf.Ospfv3.SpfTimer != nil {
						nestedVrf.Ospfv3.SpfTimer = oVrf.Ospfv3.SpfTimer
					}
					if oVrf.Ospfv3.GlobalIfTimer != nil {
						nestedVrf.Ospfv3.GlobalIfTimer = oVrf.Ospfv3.GlobalIfTimer
					}
					if oVrf.Ospfv3.RedistributionProfile != nil {
						nestedVrf.Ospfv3.RedistributionProfile = oVrf.Ospfv3.RedistributionProfile
					}
					if oVrf.Ospfv3.GlobalBfd != nil {
						nestedVrf.Ospfv3.GlobalBfd = &VrfOspfv3GlobalBfd{}
						if oVrf.Ospfv3.GlobalBfd.Misc != nil {
							entry.Misc["VrfOspfv3GlobalBfd"] = oVrf.Ospfv3.GlobalBfd.Misc
						}
						if oVrf.Ospfv3.GlobalBfd.Profile != nil {
							nestedVrf.Ospfv3.GlobalBfd.Profile = oVrf.Ospfv3.GlobalBfd.Profile
						}
					}
					if oVrf.Ospfv3.GracefulRestart != nil {
						nestedVrf.Ospfv3.GracefulRestart = &VrfOspfv3GracefulRestart{}
						if oVrf.Ospfv3.GracefulRestart.Misc != nil {
							entry.Misc["VrfOspfv3GracefulRestart"] = oVrf.Ospfv3.GracefulRestart.Misc
						}
						if oVrf.Ospfv3.GracefulRestart.Enable != nil {
							nestedVrf.Ospfv3.GracefulRestart.Enable = util.AsBool(oVrf.Ospfv3.GracefulRestart.Enable, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.GracePeriod != nil {
							nestedVrf.Ospfv3.GracefulRestart.GracePeriod = oVrf.Ospfv3.GracefulRestart.GracePeriod
						}
						if oVrf.Ospfv3.GracefulRestart.HelperEnable != nil {
							nestedVrf.Ospfv3.GracefulRestart.HelperEnable = util.AsBool(oVrf.Ospfv3.GracefulRestart.HelperEnable, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
							nestedVrf.Ospfv3.GracefulRestart.StrictLSAChecking = util.AsBool(oVrf.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
							nestedVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime
						}
					}
					if oVrf.Ospfv3.Area != nil {
						nestedVrf.Ospfv3.Area = []VrfOspfv3Area{}
						for _, oVrfOspfv3Area := range oVrf.Ospfv3.Area {
							nestedVrfOspfv3Area := VrfOspfv3Area{}
							if oVrfOspfv3Area.Misc != nil {
								entry.Misc["VrfOspfv3Area"] = oVrfOspfv3Area.Misc
							}
							if oVrfOspfv3Area.Name != "" {
								nestedVrfOspfv3Area.Name = oVrfOspfv3Area.Name
							}
							if oVrfOspfv3Area.Authentication != nil {
								nestedVrfOspfv3Area.Authentication = oVrfOspfv3Area.Authentication
							}
							if oVrfOspfv3Area.Type != nil {
								nestedVrfOspfv3Area.Type = &VrfOspfv3AreaType{}
								if oVrfOspfv3Area.Type.Misc != nil {
									entry.Misc["VrfOspfv3AreaType"] = oVrfOspfv3Area.Type.Misc
								}
								if oVrfOspfv3Area.Type.Normal != nil {
									nestedVrfOspfv3Area.Type.Normal = &VrfOspfv3AreaTypeNormal{}
									if oVrfOspfv3Area.Type.Normal.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeNormal"] = oVrfOspfv3Area.Type.Normal.Misc
									}
									if oVrfOspfv3Area.Type.Normal.Abr != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr = &VrfOspfv3AreaTypeNormalAbr{}
										if oVrfOspfv3Area.Type.Normal.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNormalAbr"] = oVrfOspfv3Area.Type.Normal.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Normal.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.ImportList = oVrfOspfv3Area.Type.Normal.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.ExportList = oVrfOspfv3Area.Type.Normal.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.InboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList
										}
									}
								}
								if oVrfOspfv3Area.Type.Stub != nil {
									nestedVrfOspfv3Area.Type.Stub = &VrfOspfv3AreaTypeStub{}
									if oVrfOspfv3Area.Type.Stub.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeStub"] = oVrfOspfv3Area.Type.Stub.Misc
									}
									if oVrfOspfv3Area.Type.Stub.NoSummary != nil {
										nestedVrfOspfv3Area.Type.Stub.NoSummary = util.AsBool(oVrfOspfv3Area.Type.Stub.NoSummary, nil)
									}
									if oVrfOspfv3Area.Type.Stub.Abr != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr = &VrfOspfv3AreaTypeStubAbr{}
										if oVrfOspfv3Area.Type.Stub.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeStubAbr"] = oVrfOspfv3Area.Type.Stub.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Stub.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.ImportList = oVrfOspfv3Area.Type.Stub.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.ExportList = oVrfOspfv3Area.Type.Stub.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.InboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList
										}
									}
									if oVrfOspfv3Area.Type.Stub.DefaultRouteMetric != nil {
										nestedVrfOspfv3Area.Type.Stub.DefaultRouteMetric = oVrfOspfv3Area.Type.Stub.DefaultRouteMetric
									}
								}
								if oVrfOspfv3Area.Type.Nssa != nil {
									nestedVrfOspfv3Area.Type.Nssa = &VrfOspfv3AreaTypeNssa{}
									if oVrfOspfv3Area.Type.Nssa.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeNssa"] = oVrfOspfv3Area.Type.Nssa.Misc
									}
									if oVrfOspfv3Area.Type.Nssa.NoSummary != nil {
										nestedVrfOspfv3Area.Type.Nssa.NoSummary = util.AsBool(oVrfOspfv3Area.Type.Nssa.NoSummary, nil)
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate = &VrfOspfv3AreaTypeNssaDefaultInformationOriginate{}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"] = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc
										}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric != nil {
											nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric
										}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
											nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType
										}
									}
									if oVrfOspfv3Area.Type.Nssa.Abr != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr = &VrfOspfv3AreaTypeNssaAbr{}
										if oVrfOspfv3Area.Type.Nssa.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNssaAbr"] = oVrfOspfv3Area.Type.Nssa.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.ImportList = oVrfOspfv3Area.Type.Nssa.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.ExportList = oVrfOspfv3Area.Type.Nssa.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = []VrfOspfv3AreaTypeNssaAbrNssaExtRange{}
											for _, oVrfOspfv3AreaTypeNssaAbrNssaExtRange := range oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange := VrfOspfv3AreaTypeNssaAbrNssaExtRange{}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc != nil {
													entry.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"] = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc
												}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name != "" {
													nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name
												}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise != nil {
													nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise = util.AsBool(oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise, nil)
												}
												nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange)
											}
										}
									}
								}
							}
							if oVrfOspfv3Area.Range != nil {
								nestedVrfOspfv3Area.Range = []VrfOspfv3AreaRange{}
								for _, oVrfOspfv3AreaRange := range oVrfOspfv3Area.Range {
									nestedVrfOspfv3AreaRange := VrfOspfv3AreaRange{}
									if oVrfOspfv3AreaRange.Misc != nil {
										entry.Misc["VrfOspfv3AreaRange"] = oVrfOspfv3AreaRange.Misc
									}
									if oVrfOspfv3AreaRange.Name != "" {
										nestedVrfOspfv3AreaRange.Name = oVrfOspfv3AreaRange.Name
									}
									if oVrfOspfv3AreaRange.Advertise != nil {
										nestedVrfOspfv3AreaRange.Advertise = util.AsBool(oVrfOspfv3AreaRange.Advertise, nil)
									}
									nestedVrfOspfv3Area.Range = append(nestedVrfOspfv3Area.Range, nestedVrfOspfv3AreaRange)
								}
							}
							if oVrfOspfv3Area.Interface != nil {
								nestedVrfOspfv3Area.Interface = []VrfOspfv3AreaInterface{}
								for _, oVrfOspfv3AreaInterface := range oVrfOspfv3Area.Interface {
									nestedVrfOspfv3AreaInterface := VrfOspfv3AreaInterface{}
									if oVrfOspfv3AreaInterface.Misc != nil {
										entry.Misc["VrfOspfv3AreaInterface"] = oVrfOspfv3AreaInterface.Misc
									}
									if oVrfOspfv3AreaInterface.Name != "" {
										nestedVrfOspfv3AreaInterface.Name = oVrfOspfv3AreaInterface.Name
									}
									if oVrfOspfv3AreaInterface.Enable != nil {
										nestedVrfOspfv3AreaInterface.Enable = util.AsBool(oVrfOspfv3AreaInterface.Enable, nil)
									}
									if oVrfOspfv3AreaInterface.MtuIgnore != nil {
										nestedVrfOspfv3AreaInterface.MtuIgnore = util.AsBool(oVrfOspfv3AreaInterface.MtuIgnore, nil)
									}
									if oVrfOspfv3AreaInterface.Passive != nil {
										nestedVrfOspfv3AreaInterface.Passive = util.AsBool(oVrfOspfv3AreaInterface.Passive, nil)
									}
									if oVrfOspfv3AreaInterface.Priority != nil {
										nestedVrfOspfv3AreaInterface.Priority = oVrfOspfv3AreaInterface.Priority
									}
									if oVrfOspfv3AreaInterface.Metric != nil {
										nestedVrfOspfv3AreaInterface.Metric = oVrfOspfv3AreaInterface.Metric
									}
									if oVrfOspfv3AreaInterface.InstanceId != nil {
										nestedVrfOspfv3AreaInterface.InstanceId = oVrfOspfv3AreaInterface.InstanceId
									}
									if oVrfOspfv3AreaInterface.Authentication != nil {
										nestedVrfOspfv3AreaInterface.Authentication = oVrfOspfv3AreaInterface.Authentication
									}
									if oVrfOspfv3AreaInterface.Timing != nil {
										nestedVrfOspfv3AreaInterface.Timing = oVrfOspfv3AreaInterface.Timing
									}
									if oVrfOspfv3AreaInterface.LinkType != nil {
										nestedVrfOspfv3AreaInterface.LinkType = &VrfOspfv3AreaInterfaceLinkType{}
										if oVrfOspfv3AreaInterface.LinkType.Misc != nil {
											entry.Misc["VrfOspfv3AreaInterfaceLinkType"] = oVrfOspfv3AreaInterface.LinkType.Misc
										}
										if oVrfOspfv3AreaInterface.LinkType.Broadcast != nil {
											nestedVrfOspfv3AreaInterface.LinkType.Broadcast = &VrfOspfv3AreaInterfaceLinkTypeBroadcast{}
											if oVrfOspfv3AreaInterface.LinkType.Broadcast.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"] = oVrfOspfv3AreaInterface.LinkType.Broadcast.Misc
											}
										}
										if oVrfOspfv3AreaInterface.LinkType.P2p != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2p = &VrfOspfv3AreaInterfaceLinkTypeP2p{}
											if oVrfOspfv3AreaInterface.LinkType.P2p.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"] = oVrfOspfv3AreaInterface.LinkType.P2p.Misc
											}
										}
										if oVrfOspfv3AreaInterface.LinkType.P2mp != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp = &VrfOspfv3AreaInterfaceLinkTypeP2mp{}
											if oVrfOspfv3AreaInterface.LinkType.P2mp.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"] = oVrfOspfv3AreaInterface.LinkType.P2mp.Misc
											}
											if oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor != nil {
												nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{}
												for _, oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc != nil {
														entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"] = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc
													}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
														nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name
													}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
														nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority
													}
													nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor)
												}
											}
										}
									}
									if oVrfOspfv3AreaInterface.Bfd != nil {
										nestedVrfOspfv3AreaInterface.Bfd = &VrfOspfv3AreaInterfaceBfd{}
										if oVrfOspfv3AreaInterface.Bfd.Misc != nil {
											entry.Misc["VrfOspfv3AreaInterfaceBfd"] = oVrfOspfv3AreaInterface.Bfd.Misc
										}
										if oVrfOspfv3AreaInterface.Bfd.Profile != nil {
											nestedVrfOspfv3AreaInterface.Bfd.Profile = oVrfOspfv3AreaInterface.Bfd.Profile
										}
									}
									nestedVrfOspfv3Area.Interface = append(nestedVrfOspfv3Area.Interface, nestedVrfOspfv3AreaInterface)
								}
							}
							if oVrfOspfv3Area.VirtualLink != nil {
								nestedVrfOspfv3Area.VirtualLink = []VrfOspfv3AreaVirtualLink{}
								for _, oVrfOspfv3AreaVirtualLink := range oVrfOspfv3Area.VirtualLink {
									nestedVrfOspfv3AreaVirtualLink := VrfOspfv3AreaVirtualLink{}
									if oVrfOspfv3AreaVirtualLink.Misc != nil {
										entry.Misc["VrfOspfv3AreaVirtualLink"] = oVrfOspfv3AreaVirtualLink.Misc
									}
									if oVrfOspfv3AreaVirtualLink.Name != "" {
										nestedVrfOspfv3AreaVirtualLink.Name = oVrfOspfv3AreaVirtualLink.Name
									}
									if oVrfOspfv3AreaVirtualLink.NeighborId != nil {
										nestedVrfOspfv3AreaVirtualLink.NeighborId = oVrfOspfv3AreaVirtualLink.NeighborId
									}
									if oVrfOspfv3AreaVirtualLink.TransitAreaId != nil {
										nestedVrfOspfv3AreaVirtualLink.TransitAreaId = oVrfOspfv3AreaVirtualLink.TransitAreaId
									}
									if oVrfOspfv3AreaVirtualLink.Enable != nil {
										nestedVrfOspfv3AreaVirtualLink.Enable = util.AsBool(oVrfOspfv3AreaVirtualLink.Enable, nil)
									}
									if oVrfOspfv3AreaVirtualLink.InstanceId != nil {
										nestedVrfOspfv3AreaVirtualLink.InstanceId = oVrfOspfv3AreaVirtualLink.InstanceId
									}
									if oVrfOspfv3AreaVirtualLink.Timing != nil {
										nestedVrfOspfv3AreaVirtualLink.Timing = oVrfOspfv3AreaVirtualLink.Timing
									}
									if oVrfOspfv3AreaVirtualLink.Authentication != nil {
										nestedVrfOspfv3AreaVirtualLink.Authentication = oVrfOspfv3AreaVirtualLink.Authentication
									}
									nestedVrfOspfv3Area.VirtualLink = append(nestedVrfOspfv3Area.VirtualLink, nestedVrfOspfv3AreaVirtualLink)
								}
							}
							nestedVrf.Ospfv3.Area = append(nestedVrf.Ospfv3.Area, nestedVrfOspfv3Area)
						}
					}
				}
				if oVrf.Ecmp != nil {
					nestedVrf.Ecmp = &VrfEcmp{}
					if oVrf.Ecmp.Misc != nil {
						entry.Misc["VrfEcmp"] = oVrf.Ecmp.Misc
					}
					if oVrf.Ecmp.Enable != nil {
						nestedVrf.Ecmp.Enable = util.AsBool(oVrf.Ecmp.Enable, nil)
					}
					if oVrf.Ecmp.MaxPath != nil {
						nestedVrf.Ecmp.MaxPath = oVrf.Ecmp.MaxPath
					}
					if oVrf.Ecmp.SymmetricReturn != nil {
						nestedVrf.Ecmp.SymmetricReturn = util.AsBool(oVrf.Ecmp.SymmetricReturn, nil)
					}
					if oVrf.Ecmp.StrictSourcePath != nil {
						nestedVrf.Ecmp.StrictSourcePath = util.AsBool(oVrf.Ecmp.StrictSourcePath, nil)
					}
					if oVrf.Ecmp.Algorithm != nil {
						nestedVrf.Ecmp.Algorithm = &VrfEcmpAlgorithm{}
						if oVrf.Ecmp.Algorithm.Misc != nil {
							entry.Misc["VrfEcmpAlgorithm"] = oVrf.Ecmp.Algorithm.Misc
						}
						if oVrf.Ecmp.Algorithm.IpModulo != nil {
							nestedVrf.Ecmp.Algorithm.IpModulo = &VrfEcmpAlgorithmIpModulo{}
							if oVrf.Ecmp.Algorithm.IpModulo.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmIpModulo"] = oVrf.Ecmp.Algorithm.IpModulo.Misc
							}
						}
						if oVrf.Ecmp.Algorithm.IpHash != nil {
							nestedVrf.Ecmp.Algorithm.IpHash = &VrfEcmpAlgorithmIpHash{}
							if oVrf.Ecmp.Algorithm.IpHash.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmIpHash"] = oVrf.Ecmp.Algorithm.IpHash.Misc
							}
							if oVrf.Ecmp.Algorithm.IpHash.SrcOnly != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.SrcOnly = util.AsBool(oVrf.Ecmp.Algorithm.IpHash.SrcOnly, nil)
							}
							if oVrf.Ecmp.Algorithm.IpHash.UsePort != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.UsePort = util.AsBool(oVrf.Ecmp.Algorithm.IpHash.UsePort, nil)
							}
							if oVrf.Ecmp.Algorithm.IpHash.HashSeed != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.HashSeed = oVrf.Ecmp.Algorithm.IpHash.HashSeed
							}
						}
						if oVrf.Ecmp.Algorithm.WeightedRoundRobin != nil {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin = &VrfEcmpAlgorithmWeightedRoundRobin{}
							if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmWeightedRoundRobin"] = oVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc
							}
							if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
								nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = []VrfEcmpAlgorithmWeightedRoundRobinInterface{}
								for _, oVrfEcmpAlgorithmWeightedRoundRobinInterface := range oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface := VrfEcmpAlgorithmWeightedRoundRobinInterface{}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc != nil {
										entry.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"] = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc
									}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
										nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Name = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name
									}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
										nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight
									}
									nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = append(nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface, nestedVrfEcmpAlgorithmWeightedRoundRobinInterface)
								}
							}
						}
						if oVrf.Ecmp.Algorithm.BalancedRoundRobin != nil {
							nestedVrf.Ecmp.Algorithm.BalancedRoundRobin = &VrfEcmpAlgorithmBalancedRoundRobin{}
							if oVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmBalancedRoundRobin"] = oVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc
							}
						}
					}
				}
				if oVrf.Multicast != nil {
					nestedVrf.Multicast = &VrfMulticast{}
					if oVrf.Multicast.Misc != nil {
						entry.Misc["VrfMulticast"] = oVrf.Multicast.Misc
					}
					if oVrf.Multicast.Enable != nil {
						nestedVrf.Multicast.Enable = util.AsBool(oVrf.Multicast.Enable, nil)
					}
					if oVrf.Multicast.StaticRoute != nil {
						nestedVrf.Multicast.StaticRoute = []VrfMulticastStaticRoute{}
						for _, oVrfMulticastStaticRoute := range oVrf.Multicast.StaticRoute {
							nestedVrfMulticastStaticRoute := VrfMulticastStaticRoute{}
							if oVrfMulticastStaticRoute.Misc != nil {
								entry.Misc["VrfMulticastStaticRoute"] = oVrfMulticastStaticRoute.Misc
							}
							if oVrfMulticastStaticRoute.Name != "" {
								nestedVrfMulticastStaticRoute.Name = oVrfMulticastStaticRoute.Name
							}
							if oVrfMulticastStaticRoute.Destination != nil {
								nestedVrfMulticastStaticRoute.Destination = oVrfMulticastStaticRoute.Destination
							}
							if oVrfMulticastStaticRoute.Interface != nil {
								nestedVrfMulticastStaticRoute.Interface = oVrfMulticastStaticRoute.Interface
							}
							if oVrfMulticastStaticRoute.Preference != nil {
								nestedVrfMulticastStaticRoute.Preference = oVrfMulticastStaticRoute.Preference
							}
							if oVrfMulticastStaticRoute.Nexthop != nil {
								nestedVrfMulticastStaticRoute.Nexthop = &VrfMulticastStaticRouteNexthop{}
								if oVrfMulticastStaticRoute.Nexthop.Misc != nil {
									entry.Misc["VrfMulticastStaticRouteNexthop"] = oVrfMulticastStaticRoute.Nexthop.Misc
								}
								if oVrfMulticastStaticRoute.Nexthop.IpAddress != nil {
									nestedVrfMulticastStaticRoute.Nexthop.IpAddress = oVrfMulticastStaticRoute.Nexthop.IpAddress
								}
							}
							nestedVrf.Multicast.StaticRoute = append(nestedVrf.Multicast.StaticRoute, nestedVrfMulticastStaticRoute)
						}
					}
					if oVrf.Multicast.Pim != nil {
						nestedVrf.Multicast.Pim = &VrfMulticastPim{}
						if oVrf.Multicast.Pim.Misc != nil {
							entry.Misc["VrfMulticastPim"] = oVrf.Multicast.Pim.Misc
						}
						if oVrf.Multicast.Pim.Enable != nil {
							nestedVrf.Multicast.Pim.Enable = util.AsBool(oVrf.Multicast.Pim.Enable, nil)
						}
						if oVrf.Multicast.Pim.RpfLookupMode != nil {
							nestedVrf.Multicast.Pim.RpfLookupMode = oVrf.Multicast.Pim.RpfLookupMode
						}
						if oVrf.Multicast.Pim.RouteAgeoutTime != nil {
							nestedVrf.Multicast.Pim.RouteAgeoutTime = oVrf.Multicast.Pim.RouteAgeoutTime
						}
						if oVrf.Multicast.Pim.IfTimerGlobal != nil {
							nestedVrf.Multicast.Pim.IfTimerGlobal = oVrf.Multicast.Pim.IfTimerGlobal
						}
						if oVrf.Multicast.Pim.GroupPermission != nil {
							nestedVrf.Multicast.Pim.GroupPermission = oVrf.Multicast.Pim.GroupPermission
						}
						if oVrf.Multicast.Pim.SsmAddressSpace != nil {
							nestedVrf.Multicast.Pim.SsmAddressSpace = &VrfMulticastPimSsmAddressSpace{}
							if oVrf.Multicast.Pim.SsmAddressSpace.Misc != nil {
								entry.Misc["VrfMulticastPimSsmAddressSpace"] = oVrf.Multicast.Pim.SsmAddressSpace.Misc
							}
							if oVrf.Multicast.Pim.SsmAddressSpace.GroupList != nil {
								nestedVrf.Multicast.Pim.SsmAddressSpace.GroupList = oVrf.Multicast.Pim.SsmAddressSpace.GroupList
							}
						}
						if oVrf.Multicast.Pim.Rp != nil {
							nestedVrf.Multicast.Pim.Rp = &VrfMulticastPimRp{}
							if oVrf.Multicast.Pim.Rp.Misc != nil {
								entry.Misc["VrfMulticastPimRp"] = oVrf.Multicast.Pim.Rp.Misc
							}
							if oVrf.Multicast.Pim.Rp.LocalRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp = &VrfMulticastPimRpLocalRp{}
								if oVrf.Multicast.Pim.Rp.LocalRp.Misc != nil {
									entry.Misc["VrfMulticastPimRpLocalRp"] = oVrf.Multicast.Pim.Rp.LocalRp.Misc
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp = &VrfMulticastPimRpLocalRpStaticRp{}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpLocalRpStaticRp"] = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override = util.AsBool(oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override, nil)
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList
									}
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp = &VrfMulticastPimRpLocalRpCandidateRp{}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpLocalRpCandidateRp"] = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList
									}
								}
							}
							if oVrf.Multicast.Pim.Rp.ExternalRp != nil {
								nestedVrf.Multicast.Pim.Rp.ExternalRp = []VrfMulticastPimRpExternalRp{}
								for _, oVrfMulticastPimRpExternalRp := range oVrf.Multicast.Pim.Rp.ExternalRp {
									nestedVrfMulticastPimRpExternalRp := VrfMulticastPimRpExternalRp{}
									if oVrfMulticastPimRpExternalRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpExternalRp"] = oVrfMulticastPimRpExternalRp.Misc
									}
									if oVrfMulticastPimRpExternalRp.Name != "" {
										nestedVrfMulticastPimRpExternalRp.Name = oVrfMulticastPimRpExternalRp.Name
									}
									if oVrfMulticastPimRpExternalRp.GroupList != nil {
										nestedVrfMulticastPimRpExternalRp.GroupList = oVrfMulticastPimRpExternalRp.GroupList
									}
									if oVrfMulticastPimRpExternalRp.Override != nil {
										nestedVrfMulticastPimRpExternalRp.Override = util.AsBool(oVrfMulticastPimRpExternalRp.Override, nil)
									}
									nestedVrf.Multicast.Pim.Rp.ExternalRp = append(nestedVrf.Multicast.Pim.Rp.ExternalRp, nestedVrfMulticastPimRpExternalRp)
								}
							}
						}
						if oVrf.Multicast.Pim.SptThreshold != nil {
							nestedVrf.Multicast.Pim.SptThreshold = []VrfMulticastPimSptThreshold{}
							for _, oVrfMulticastPimSptThreshold := range oVrf.Multicast.Pim.SptThreshold {
								nestedVrfMulticastPimSptThreshold := VrfMulticastPimSptThreshold{}
								if oVrfMulticastPimSptThreshold.Misc != nil {
									entry.Misc["VrfMulticastPimSptThreshold"] = oVrfMulticastPimSptThreshold.Misc
								}
								if oVrfMulticastPimSptThreshold.Name != "" {
									nestedVrfMulticastPimSptThreshold.Name = oVrfMulticastPimSptThreshold.Name
								}
								if oVrfMulticastPimSptThreshold.Threshold != nil {
									nestedVrfMulticastPimSptThreshold.Threshold = oVrfMulticastPimSptThreshold.Threshold
								}
								nestedVrf.Multicast.Pim.SptThreshold = append(nestedVrf.Multicast.Pim.SptThreshold, nestedVrfMulticastPimSptThreshold)
							}
						}
						if oVrf.Multicast.Pim.Interface != nil {
							nestedVrf.Multicast.Pim.Interface = []VrfMulticastPimInterface{}
							for _, oVrfMulticastPimInterface := range oVrf.Multicast.Pim.Interface {
								nestedVrfMulticastPimInterface := VrfMulticastPimInterface{}
								if oVrfMulticastPimInterface.Misc != nil {
									entry.Misc["VrfMulticastPimInterface"] = oVrfMulticastPimInterface.Misc
								}
								if oVrfMulticastPimInterface.Name != "" {
									nestedVrfMulticastPimInterface.Name = oVrfMulticastPimInterface.Name
								}
								if oVrfMulticastPimInterface.Description != nil {
									nestedVrfMulticastPimInterface.Description = oVrfMulticastPimInterface.Description
								}
								if oVrfMulticastPimInterface.DrPriority != nil {
									nestedVrfMulticastPimInterface.DrPriority = oVrfMulticastPimInterface.DrPriority
								}
								if oVrfMulticastPimInterface.SendBsm != nil {
									nestedVrfMulticastPimInterface.SendBsm = util.AsBool(oVrfMulticastPimInterface.SendBsm, nil)
								}
								if oVrfMulticastPimInterface.IfTimer != nil {
									nestedVrfMulticastPimInterface.IfTimer = oVrfMulticastPimInterface.IfTimer
								}
								if oVrfMulticastPimInterface.NeighborFilter != nil {
									nestedVrfMulticastPimInterface.NeighborFilter = oVrfMulticastPimInterface.NeighborFilter
								}
								nestedVrf.Multicast.Pim.Interface = append(nestedVrf.Multicast.Pim.Interface, nestedVrfMulticastPimInterface)
							}
						}
					}
					if oVrf.Multicast.Igmp != nil {
						nestedVrf.Multicast.Igmp = &VrfMulticastIgmp{}
						if oVrf.Multicast.Igmp.Misc != nil {
							entry.Misc["VrfMulticastIgmp"] = oVrf.Multicast.Igmp.Misc
						}
						if oVrf.Multicast.Igmp.Enable != nil {
							nestedVrf.Multicast.Igmp.Enable = util.AsBool(oVrf.Multicast.Igmp.Enable, nil)
						}
						if oVrf.Multicast.Igmp.Dynamic != nil {
							nestedVrf.Multicast.Igmp.Dynamic = &VrfMulticastIgmpDynamic{}
							if oVrf.Multicast.Igmp.Dynamic.Misc != nil {
								entry.Misc["VrfMulticastIgmpDynamic"] = oVrf.Multicast.Igmp.Dynamic.Misc
							}
							if oVrf.Multicast.Igmp.Dynamic.Interface != nil {
								nestedVrf.Multicast.Igmp.Dynamic.Interface = []VrfMulticastIgmpDynamicInterface{}
								for _, oVrfMulticastIgmpDynamicInterface := range oVrf.Multicast.Igmp.Dynamic.Interface {
									nestedVrfMulticastIgmpDynamicInterface := VrfMulticastIgmpDynamicInterface{}
									if oVrfMulticastIgmpDynamicInterface.Misc != nil {
										entry.Misc["VrfMulticastIgmpDynamicInterface"] = oVrfMulticastIgmpDynamicInterface.Misc
									}
									if oVrfMulticastIgmpDynamicInterface.Name != "" {
										nestedVrfMulticastIgmpDynamicInterface.Name = oVrfMulticastIgmpDynamicInterface.Name
									}
									if oVrfMulticastIgmpDynamicInterface.Version != nil {
										nestedVrfMulticastIgmpDynamicInterface.Version = oVrfMulticastIgmpDynamicInterface.Version
									}
									if oVrfMulticastIgmpDynamicInterface.Robustness != nil {
										nestedVrfMulticastIgmpDynamicInterface.Robustness = oVrfMulticastIgmpDynamicInterface.Robustness
									}
									if oVrfMulticastIgmpDynamicInterface.GroupFilter != nil {
										nestedVrfMulticastIgmpDynamicInterface.GroupFilter = oVrfMulticastIgmpDynamicInterface.GroupFilter
									}
									if oVrfMulticastIgmpDynamicInterface.MaxGroups != nil {
										nestedVrfMulticastIgmpDynamicInterface.MaxGroups = oVrfMulticastIgmpDynamicInterface.MaxGroups
									}
									if oVrfMulticastIgmpDynamicInterface.MaxSources != nil {
										nestedVrfMulticastIgmpDynamicInterface.MaxSources = oVrfMulticastIgmpDynamicInterface.MaxSources
									}
									if oVrfMulticastIgmpDynamicInterface.QueryProfile != nil {
										nestedVrfMulticastIgmpDynamicInterface.QueryProfile = oVrfMulticastIgmpDynamicInterface.QueryProfile
									}
									if oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing != nil {
										nestedVrfMulticastIgmpDynamicInterface.RouterAlertPolicing = util.AsBool(oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing, nil)
									}
									nestedVrf.Multicast.Igmp.Dynamic.Interface = append(nestedVrf.Multicast.Igmp.Dynamic.Interface, nestedVrfMulticastIgmpDynamicInterface)
								}
							}
						}
						if oVrf.Multicast.Igmp.Static != nil {
							nestedVrf.Multicast.Igmp.Static = []VrfMulticastIgmpStatic{}
							for _, oVrfMulticastIgmpStatic := range oVrf.Multicast.Igmp.Static {
								nestedVrfMulticastIgmpStatic := VrfMulticastIgmpStatic{}
								if oVrfMulticastIgmpStatic.Misc != nil {
									entry.Misc["VrfMulticastIgmpStatic"] = oVrfMulticastIgmpStatic.Misc
								}
								if oVrfMulticastIgmpStatic.Name != "" {
									nestedVrfMulticastIgmpStatic.Name = oVrfMulticastIgmpStatic.Name
								}
								if oVrfMulticastIgmpStatic.Interface != nil {
									nestedVrfMulticastIgmpStatic.Interface = oVrfMulticastIgmpStatic.Interface
								}
								if oVrfMulticastIgmpStatic.GroupAddress != nil {
									nestedVrfMulticastIgmpStatic.GroupAddress = oVrfMulticastIgmpStatic.GroupAddress
								}
								if oVrfMulticastIgmpStatic.SourceAddress != nil {
									nestedVrfMulticastIgmpStatic.SourceAddress = oVrfMulticastIgmpStatic.SourceAddress
								}
								nestedVrf.Multicast.Igmp.Static = append(nestedVrf.Multicast.Igmp.Static, nestedVrfMulticastIgmpStatic)
							}
						}
					}
					if oVrf.Multicast.Msdp != nil {
						nestedVrf.Multicast.Msdp = &VrfMulticastMsdp{}
						if oVrf.Multicast.Msdp.Misc != nil {
							entry.Misc["VrfMulticastMsdp"] = oVrf.Multicast.Msdp.Misc
						}
						if oVrf.Multicast.Msdp.Enable != nil {
							nestedVrf.Multicast.Msdp.Enable = util.AsBool(oVrf.Multicast.Msdp.Enable, nil)
						}
						if oVrf.Multicast.Msdp.GlobalTimer != nil {
							nestedVrf.Multicast.Msdp.GlobalTimer = oVrf.Multicast.Msdp.GlobalTimer
						}
						if oVrf.Multicast.Msdp.GlobalAuthentication != nil {
							nestedVrf.Multicast.Msdp.GlobalAuthentication = oVrf.Multicast.Msdp.GlobalAuthentication
						}
						if oVrf.Multicast.Msdp.OriginatorId != nil {
							nestedVrf.Multicast.Msdp.OriginatorId = &VrfMulticastMsdpOriginatorId{}
							if oVrf.Multicast.Msdp.OriginatorId.Misc != nil {
								entry.Misc["VrfMulticastMsdpOriginatorId"] = oVrf.Multicast.Msdp.OriginatorId.Misc
							}
							if oVrf.Multicast.Msdp.OriginatorId.Interface != nil {
								nestedVrf.Multicast.Msdp.OriginatorId.Interface = oVrf.Multicast.Msdp.OriginatorId.Interface
							}
							if oVrf.Multicast.Msdp.OriginatorId.Ip != nil {
								nestedVrf.Multicast.Msdp.OriginatorId.Ip = oVrf.Multicast.Msdp.OriginatorId.Ip
							}
						}
						if oVrf.Multicast.Msdp.Peer != nil {
							nestedVrf.Multicast.Msdp.Peer = []VrfMulticastMsdpPeer{}
							for _, oVrfMulticastMsdpPeer := range oVrf.Multicast.Msdp.Peer {
								nestedVrfMulticastMsdpPeer := VrfMulticastMsdpPeer{}
								if oVrfMulticastMsdpPeer.Misc != nil {
									entry.Misc["VrfMulticastMsdpPeer"] = oVrfMulticastMsdpPeer.Misc
								}
								if oVrfMulticastMsdpPeer.Name != "" {
									nestedVrfMulticastMsdpPeer.Name = oVrfMulticastMsdpPeer.Name
								}
								if oVrfMulticastMsdpPeer.Enable != nil {
									nestedVrfMulticastMsdpPeer.Enable = util.AsBool(oVrfMulticastMsdpPeer.Enable, nil)
								}
								if oVrfMulticastMsdpPeer.PeerAs != nil {
									nestedVrfMulticastMsdpPeer.PeerAs = oVrfMulticastMsdpPeer.PeerAs
								}
								if oVrfMulticastMsdpPeer.Authentication != nil {
									nestedVrfMulticastMsdpPeer.Authentication = oVrfMulticastMsdpPeer.Authentication
								}
								if oVrfMulticastMsdpPeer.MaxSa != nil {
									nestedVrfMulticastMsdpPeer.MaxSa = oVrfMulticastMsdpPeer.MaxSa
								}
								if oVrfMulticastMsdpPeer.InboundSaFilter != nil {
									nestedVrfMulticastMsdpPeer.InboundSaFilter = oVrfMulticastMsdpPeer.InboundSaFilter
								}
								if oVrfMulticastMsdpPeer.OutboundSaFilter != nil {
									nestedVrfMulticastMsdpPeer.OutboundSaFilter = oVrfMulticastMsdpPeer.OutboundSaFilter
								}
								if oVrfMulticastMsdpPeer.LocalAddress != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress = &VrfMulticastMsdpPeerLocalAddress{}
									if oVrfMulticastMsdpPeer.LocalAddress.Misc != nil {
										entry.Misc["VrfMulticastMsdpPeerLocalAddress"] = oVrfMulticastMsdpPeer.LocalAddress.Misc
									}
									if oVrfMulticastMsdpPeer.LocalAddress.Interface != nil {
										nestedVrfMulticastMsdpPeer.LocalAddress.Interface = oVrfMulticastMsdpPeer.LocalAddress.Interface
									}
									if oVrfMulticastMsdpPeer.LocalAddress.Ip != nil {
										nestedVrfMulticastMsdpPeer.LocalAddress.Ip = oVrfMulticastMsdpPeer.LocalAddress.Ip
									}
								}
								if oVrfMulticastMsdpPeer.PeerAddress != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress = &VrfMulticastMsdpPeerPeerAddress{}
									if oVrfMulticastMsdpPeer.PeerAddress.Misc != nil {
										entry.Misc["VrfMulticastMsdpPeerPeerAddress"] = oVrfMulticastMsdpPeer.PeerAddress.Misc
									}
									if oVrfMulticastMsdpPeer.PeerAddress.Ip != nil {
										nestedVrfMulticastMsdpPeer.PeerAddress.Ip = oVrfMulticastMsdpPeer.PeerAddress.Ip
									}
									if oVrfMulticastMsdpPeer.PeerAddress.Fqdn != nil {
										nestedVrfMulticastMsdpPeer.PeerAddress.Fqdn = oVrfMulticastMsdpPeer.PeerAddress.Fqdn
									}
								}
								nestedVrf.Multicast.Msdp.Peer = append(nestedVrf.Multicast.Msdp.Peer, nestedVrfMulticastMsdpPeer)
							}
						}
					}
				}
				if oVrf.Rip != nil {
					nestedVrf.Rip = &VrfRip{}
					if oVrf.Rip.Misc != nil {
						entry.Misc["VrfRip"] = oVrf.Rip.Misc
					}
					if oVrf.Rip.Enable != nil {
						nestedVrf.Rip.Enable = util.AsBool(oVrf.Rip.Enable, nil)
					}
					if oVrf.Rip.DefaultInformationOriginate != nil {
						nestedVrf.Rip.DefaultInformationOriginate = util.AsBool(oVrf.Rip.DefaultInformationOriginate, nil)
					}
					if oVrf.Rip.GlobalTimer != nil {
						nestedVrf.Rip.GlobalTimer = oVrf.Rip.GlobalTimer
					}
					if oVrf.Rip.AuthProfile != nil {
						nestedVrf.Rip.AuthProfile = oVrf.Rip.AuthProfile
					}
					if oVrf.Rip.RedistributionProfile != nil {
						nestedVrf.Rip.RedistributionProfile = oVrf.Rip.RedistributionProfile
					}
					if oVrf.Rip.GlobalBfd != nil {
						nestedVrf.Rip.GlobalBfd = &VrfRipGlobalBfd{}
						if oVrf.Rip.GlobalBfd.Misc != nil {
							entry.Misc["VrfRipGlobalBfd"] = oVrf.Rip.GlobalBfd.Misc
						}
						if oVrf.Rip.GlobalBfd.Profile != nil {
							nestedVrf.Rip.GlobalBfd.Profile = oVrf.Rip.GlobalBfd.Profile
						}
					}
					if oVrf.Rip.GlobalInboundDistributeList != nil {
						nestedVrf.Rip.GlobalInboundDistributeList = &VrfRipGlobalInboundDistributeList{}
						if oVrf.Rip.GlobalInboundDistributeList.Misc != nil {
							entry.Misc["VrfRipGlobalInboundDistributeList"] = oVrf.Rip.GlobalInboundDistributeList.Misc
						}
						if oVrf.Rip.GlobalInboundDistributeList.AccessList != nil {
							nestedVrf.Rip.GlobalInboundDistributeList.AccessList = oVrf.Rip.GlobalInboundDistributeList.AccessList
						}
					}
					if oVrf.Rip.GlobalOutboundDistributeList != nil {
						nestedVrf.Rip.GlobalOutboundDistributeList = &VrfRipGlobalOutboundDistributeList{}
						if oVrf.Rip.GlobalOutboundDistributeList.Misc != nil {
							entry.Misc["VrfRipGlobalOutboundDistributeList"] = oVrf.Rip.GlobalOutboundDistributeList.Misc
						}
						if oVrf.Rip.GlobalOutboundDistributeList.AccessList != nil {
							nestedVrf.Rip.GlobalOutboundDistributeList.AccessList = oVrf.Rip.GlobalOutboundDistributeList.AccessList
						}
					}
					if oVrf.Rip.Interface != nil {
						nestedVrf.Rip.Interface = []VrfRipInterface{}
						for _, oVrfRipInterface := range oVrf.Rip.Interface {
							nestedVrfRipInterface := VrfRipInterface{}
							if oVrfRipInterface.Misc != nil {
								entry.Misc["VrfRipInterface"] = oVrfRipInterface.Misc
							}
							if oVrfRipInterface.Name != "" {
								nestedVrfRipInterface.Name = oVrfRipInterface.Name
							}
							if oVrfRipInterface.Enable != nil {
								nestedVrfRipInterface.Enable = util.AsBool(oVrfRipInterface.Enable, nil)
							}
							if oVrfRipInterface.Mode != nil {
								nestedVrfRipInterface.Mode = oVrfRipInterface.Mode
							}
							if oVrfRipInterface.SplitHorizon != nil {
								nestedVrfRipInterface.SplitHorizon = oVrfRipInterface.SplitHorizon
							}
							if oVrfRipInterface.Authentication != nil {
								nestedVrfRipInterface.Authentication = oVrfRipInterface.Authentication
							}
							if oVrfRipInterface.Bfd != nil {
								nestedVrfRipInterface.Bfd = &VrfRipInterfaceBfd{}
								if oVrfRipInterface.Bfd.Misc != nil {
									entry.Misc["VrfRipInterfaceBfd"] = oVrfRipInterface.Bfd.Misc
								}
								if oVrfRipInterface.Bfd.Profile != nil {
									nestedVrfRipInterface.Bfd.Profile = oVrfRipInterface.Bfd.Profile
								}
							}
							if oVrfRipInterface.InterfaceInboundDistributeList != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList = &VrfRipInterfaceInterfaceInboundDistributeList{}
								if oVrfRipInterface.InterfaceInboundDistributeList.Misc != nil {
									entry.Misc["VrfRipInterfaceInterfaceInboundDistributeList"] = oVrfRipInterface.InterfaceInboundDistributeList.Misc
								}
								if oVrfRipInterface.InterfaceInboundDistributeList.AccessList != nil {
									nestedVrfRipInterface.InterfaceInboundDistributeList.AccessList = oVrfRipInterface.InterfaceInboundDistributeList.AccessList
								}
								if oVrfRipInterface.InterfaceInboundDistributeList.Metric != nil {
									nestedVrfRipInterface.InterfaceInboundDistributeList.Metric = oVrfRipInterface.InterfaceInboundDistributeList.Metric
								}
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList = &VrfRipInterfaceInterfaceOutboundDistributeList{}
								if oVrfRipInterface.InterfaceOutboundDistributeList.Misc != nil {
									entry.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"] = oVrfRipInterface.InterfaceOutboundDistributeList.Misc
								}
								if oVrfRipInterface.InterfaceOutboundDistributeList.AccessList != nil {
									nestedVrfRipInterface.InterfaceOutboundDistributeList.AccessList = oVrfRipInterface.InterfaceOutboundDistributeList.AccessList
								}
								if oVrfRipInterface.InterfaceOutboundDistributeList.Metric != nil {
									nestedVrfRipInterface.InterfaceOutboundDistributeList.Metric = oVrfRipInterface.InterfaceOutboundDistributeList.Metric
								}
							}
							nestedVrf.Rip.Interface = append(nestedVrf.Rip.Interface, nestedVrfRipInterface)
						}
					}
				}
				nestedVrfCol = append(nestedVrfCol, nestedVrf)
			}
			entry.Vrf = nestedVrfCol
		}

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}
func (c *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedVrfCol []Vrf
		if o.Vrf != nil {
			nestedVrfCol = []Vrf{}
			for _, oVrf := range o.Vrf {
				nestedVrf := Vrf{}
				if oVrf.Misc != nil {
					entry.Misc["Vrf"] = oVrf.Misc
				}
				if oVrf.Name != "" {
					nestedVrf.Name = oVrf.Name
				}
				if oVrf.Interface != nil {
					nestedVrf.Interface = util.MemToStr(oVrf.Interface)
				}
				if oVrf.AdminDists != nil {
					nestedVrf.AdminDists = &VrfAdminDists{}
					if oVrf.AdminDists.Misc != nil {
						entry.Misc["VrfAdminDists"] = oVrf.AdminDists.Misc
					}
					if oVrf.AdminDists.Static != nil {
						nestedVrf.AdminDists.Static = oVrf.AdminDists.Static
					}
					if oVrf.AdminDists.StaticIpv6 != nil {
						nestedVrf.AdminDists.StaticIpv6 = oVrf.AdminDists.StaticIpv6
					}
					if oVrf.AdminDists.OspfInter != nil {
						nestedVrf.AdminDists.OspfInter = oVrf.AdminDists.OspfInter
					}
					if oVrf.AdminDists.OspfIntra != nil {
						nestedVrf.AdminDists.OspfIntra = oVrf.AdminDists.OspfIntra
					}
					if oVrf.AdminDists.OspfExt != nil {
						nestedVrf.AdminDists.OspfExt = oVrf.AdminDists.OspfExt
					}
					if oVrf.AdminDists.Ospfv3Inter != nil {
						nestedVrf.AdminDists.Ospfv3Inter = oVrf.AdminDists.Ospfv3Inter
					}
					if oVrf.AdminDists.Ospfv3Intra != nil {
						nestedVrf.AdminDists.Ospfv3Intra = oVrf.AdminDists.Ospfv3Intra
					}
					if oVrf.AdminDists.Ospfv3Ext != nil {
						nestedVrf.AdminDists.Ospfv3Ext = oVrf.AdminDists.Ospfv3Ext
					}
					if oVrf.AdminDists.BgpInternal != nil {
						nestedVrf.AdminDists.BgpInternal = oVrf.AdminDists.BgpInternal
					}
					if oVrf.AdminDists.BgpExternal != nil {
						nestedVrf.AdminDists.BgpExternal = oVrf.AdminDists.BgpExternal
					}
					if oVrf.AdminDists.BgpLocal != nil {
						nestedVrf.AdminDists.BgpLocal = oVrf.AdminDists.BgpLocal
					}
					if oVrf.AdminDists.Rip != nil {
						nestedVrf.AdminDists.Rip = oVrf.AdminDists.Rip
					}
				}
				if oVrf.RibFilter != nil {
					nestedVrf.RibFilter = &VrfRibFilter{}
					if oVrf.RibFilter.Misc != nil {
						entry.Misc["VrfRibFilter"] = oVrf.RibFilter.Misc
					}
					if oVrf.RibFilter.Ipv4 != nil {
						nestedVrf.RibFilter.Ipv4 = &VrfRibFilterIpv4{}
						if oVrf.RibFilter.Ipv4.Misc != nil {
							entry.Misc["VrfRibFilterIpv4"] = oVrf.RibFilter.Ipv4.Misc
						}
						if oVrf.RibFilter.Ipv4.Static != nil {
							nestedVrf.RibFilter.Ipv4.Static = &VrfRibFilterIpv4Static{}
							if oVrf.RibFilter.Ipv4.Static.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Static"] = oVrf.RibFilter.Ipv4.Static.Misc
							}
							if oVrf.RibFilter.Ipv4.Static.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Static.RouteMap = oVrf.RibFilter.Ipv4.Static.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Bgp != nil {
							nestedVrf.RibFilter.Ipv4.Bgp = &VrfRibFilterIpv4Bgp{}
							if oVrf.RibFilter.Ipv4.Bgp.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Bgp"] = oVrf.RibFilter.Ipv4.Bgp.Misc
							}
							if oVrf.RibFilter.Ipv4.Bgp.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Bgp.RouteMap = oVrf.RibFilter.Ipv4.Bgp.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Ospf != nil {
							nestedVrf.RibFilter.Ipv4.Ospf = &VrfRibFilterIpv4Ospf{}
							if oVrf.RibFilter.Ipv4.Ospf.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Ospf"] = oVrf.RibFilter.Ipv4.Ospf.Misc
							}
							if oVrf.RibFilter.Ipv4.Ospf.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Ospf.RouteMap = oVrf.RibFilter.Ipv4.Ospf.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv4.Rip != nil {
							nestedVrf.RibFilter.Ipv4.Rip = &VrfRibFilterIpv4Rip{}
							if oVrf.RibFilter.Ipv4.Rip.Misc != nil {
								entry.Misc["VrfRibFilterIpv4Rip"] = oVrf.RibFilter.Ipv4.Rip.Misc
							}
							if oVrf.RibFilter.Ipv4.Rip.RouteMap != nil {
								nestedVrf.RibFilter.Ipv4.Rip.RouteMap = oVrf.RibFilter.Ipv4.Rip.RouteMap
							}
						}
					}
					if oVrf.RibFilter.Ipv6 != nil {
						nestedVrf.RibFilter.Ipv6 = &VrfRibFilterIpv6{}
						if oVrf.RibFilter.Ipv6.Misc != nil {
							entry.Misc["VrfRibFilterIpv6"] = oVrf.RibFilter.Ipv6.Misc
						}
						if oVrf.RibFilter.Ipv6.Static != nil {
							nestedVrf.RibFilter.Ipv6.Static = &VrfRibFilterIpv6Static{}
							if oVrf.RibFilter.Ipv6.Static.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Static"] = oVrf.RibFilter.Ipv6.Static.Misc
							}
							if oVrf.RibFilter.Ipv6.Static.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Static.RouteMap = oVrf.RibFilter.Ipv6.Static.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv6.Bgp != nil {
							nestedVrf.RibFilter.Ipv6.Bgp = &VrfRibFilterIpv6Bgp{}
							if oVrf.RibFilter.Ipv6.Bgp.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Bgp"] = oVrf.RibFilter.Ipv6.Bgp.Misc
							}
							if oVrf.RibFilter.Ipv6.Bgp.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Bgp.RouteMap = oVrf.RibFilter.Ipv6.Bgp.RouteMap
							}
						}
						if oVrf.RibFilter.Ipv6.Ospfv3 != nil {
							nestedVrf.RibFilter.Ipv6.Ospfv3 = &VrfRibFilterIpv6Ospfv3{}
							if oVrf.RibFilter.Ipv6.Ospfv3.Misc != nil {
								entry.Misc["VrfRibFilterIpv6Ospfv3"] = oVrf.RibFilter.Ipv6.Ospfv3.Misc
							}
							if oVrf.RibFilter.Ipv6.Ospfv3.RouteMap != nil {
								nestedVrf.RibFilter.Ipv6.Ospfv3.RouteMap = oVrf.RibFilter.Ipv6.Ospfv3.RouteMap
							}
						}
					}
				}
				if oVrf.Bgp != nil {
					nestedVrf.Bgp = &VrfBgp{}
					if oVrf.Bgp.Misc != nil {
						entry.Misc["VrfBgp"] = oVrf.Bgp.Misc
					}
					if oVrf.Bgp.Enable != nil {
						nestedVrf.Bgp.Enable = util.AsBool(oVrf.Bgp.Enable, nil)
					}
					if oVrf.Bgp.RouterId != nil {
						nestedVrf.Bgp.RouterId = oVrf.Bgp.RouterId
					}
					if oVrf.Bgp.LocalAs != nil {
						nestedVrf.Bgp.LocalAs = oVrf.Bgp.LocalAs
					}
					if oVrf.Bgp.InstallRoute != nil {
						nestedVrf.Bgp.InstallRoute = util.AsBool(oVrf.Bgp.InstallRoute, nil)
					}
					if oVrf.Bgp.EnforceFirstAs != nil {
						nestedVrf.Bgp.EnforceFirstAs = util.AsBool(oVrf.Bgp.EnforceFirstAs, nil)
					}
					if oVrf.Bgp.FastExternalFailover != nil {
						nestedVrf.Bgp.FastExternalFailover = util.AsBool(oVrf.Bgp.FastExternalFailover, nil)
					}
					if oVrf.Bgp.EcmpMultiAs != nil {
						nestedVrf.Bgp.EcmpMultiAs = util.AsBool(oVrf.Bgp.EcmpMultiAs, nil)
					}
					if oVrf.Bgp.DefaultLocalPreference != nil {
						nestedVrf.Bgp.DefaultLocalPreference = oVrf.Bgp.DefaultLocalPreference
					}
					if oVrf.Bgp.GracefulShutdown != nil {
						nestedVrf.Bgp.GracefulShutdown = util.AsBool(oVrf.Bgp.GracefulShutdown, nil)
					}
					if oVrf.Bgp.AlwaysAdvertiseNetworkRoute != nil {
						nestedVrf.Bgp.AlwaysAdvertiseNetworkRoute = util.AsBool(oVrf.Bgp.AlwaysAdvertiseNetworkRoute, nil)
					}
					if oVrf.Bgp.Med != nil {
						nestedVrf.Bgp.Med = &VrfBgpMed{}
						if oVrf.Bgp.Med.Misc != nil {
							entry.Misc["VrfBgpMed"] = oVrf.Bgp.Med.Misc
						}
						if oVrf.Bgp.Med.AlwaysCompareMed != nil {
							nestedVrf.Bgp.Med.AlwaysCompareMed = util.AsBool(oVrf.Bgp.Med.AlwaysCompareMed, nil)
						}
						if oVrf.Bgp.Med.DeterministicMedComparison != nil {
							nestedVrf.Bgp.Med.DeterministicMedComparison = util.AsBool(oVrf.Bgp.Med.DeterministicMedComparison, nil)
						}
					}
					if oVrf.Bgp.GracefulRestart != nil {
						nestedVrf.Bgp.GracefulRestart = &VrfBgpGracefulRestart{}
						if oVrf.Bgp.GracefulRestart.Misc != nil {
							entry.Misc["VrfBgpGracefulRestart"] = oVrf.Bgp.GracefulRestart.Misc
						}
						if oVrf.Bgp.GracefulRestart.Enable != nil {
							nestedVrf.Bgp.GracefulRestart.Enable = util.AsBool(oVrf.Bgp.GracefulRestart.Enable, nil)
						}
						if oVrf.Bgp.GracefulRestart.StaleRouteTime != nil {
							nestedVrf.Bgp.GracefulRestart.StaleRouteTime = oVrf.Bgp.GracefulRestart.StaleRouteTime
						}
						if oVrf.Bgp.GracefulRestart.MaxPeerRestartTime != nil {
							nestedVrf.Bgp.GracefulRestart.MaxPeerRestartTime = oVrf.Bgp.GracefulRestart.MaxPeerRestartTime
						}
						if oVrf.Bgp.GracefulRestart.LocalRestartTime != nil {
							nestedVrf.Bgp.GracefulRestart.LocalRestartTime = oVrf.Bgp.GracefulRestart.LocalRestartTime
						}
					}
					if oVrf.Bgp.GlobalBfd != nil {
						nestedVrf.Bgp.GlobalBfd = &VrfBgpGlobalBfd{}
						if oVrf.Bgp.GlobalBfd.Misc != nil {
							entry.Misc["VrfBgpGlobalBfd"] = oVrf.Bgp.GlobalBfd.Misc
						}
						if oVrf.Bgp.GlobalBfd.Profile != nil {
							nestedVrf.Bgp.GlobalBfd.Profile = oVrf.Bgp.GlobalBfd.Profile
						}
					}
					if oVrf.Bgp.RedistributionProfile != nil {
						nestedVrf.Bgp.RedistributionProfile = &VrfBgpRedistributionProfile{}
						if oVrf.Bgp.RedistributionProfile.Misc != nil {
							entry.Misc["VrfBgpRedistributionProfile"] = oVrf.Bgp.RedistributionProfile.Misc
						}
						if oVrf.Bgp.RedistributionProfile.Ipv4 != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv4 = &VrfBgpRedistributionProfileIpv4{}
							if oVrf.Bgp.RedistributionProfile.Ipv4.Misc != nil {
								entry.Misc["VrfBgpRedistributionProfileIpv4"] = oVrf.Bgp.RedistributionProfile.Ipv4.Misc
							}
							if oVrf.Bgp.RedistributionProfile.Ipv4.Unicast != nil {
								nestedVrf.Bgp.RedistributionProfile.Ipv4.Unicast = oVrf.Bgp.RedistributionProfile.Ipv4.Unicast
							}
						}
						if oVrf.Bgp.RedistributionProfile.Ipv6 != nil {
							nestedVrf.Bgp.RedistributionProfile.Ipv6 = &VrfBgpRedistributionProfileIpv6{}
							if oVrf.Bgp.RedistributionProfile.Ipv6.Misc != nil {
								entry.Misc["VrfBgpRedistributionProfileIpv6"] = oVrf.Bgp.RedistributionProfile.Ipv6.Misc
							}
							if oVrf.Bgp.RedistributionProfile.Ipv6.Unicast != nil {
								nestedVrf.Bgp.RedistributionProfile.Ipv6.Unicast = oVrf.Bgp.RedistributionProfile.Ipv6.Unicast
							}
						}
					}
					if oVrf.Bgp.AdvertiseNetwork != nil {
						nestedVrf.Bgp.AdvertiseNetwork = &VrfBgpAdvertiseNetwork{}
						if oVrf.Bgp.AdvertiseNetwork.Misc != nil {
							entry.Misc["VrfBgpAdvertiseNetwork"] = oVrf.Bgp.AdvertiseNetwork.Misc
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv4 != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv4 = &VrfBgpAdvertiseNetworkIpv4{}
							if oVrf.Bgp.AdvertiseNetwork.Ipv4.Misc != nil {
								entry.Misc["VrfBgpAdvertiseNetworkIpv4"] = oVrf.Bgp.AdvertiseNetwork.Ipv4.Misc
							}
							if oVrf.Bgp.AdvertiseNetwork.Ipv4.Network != nil {
								nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = []VrfBgpAdvertiseNetworkIpv4Network{}
								for _, oVrfBgpAdvertiseNetworkIpv4Network := range oVrf.Bgp.AdvertiseNetwork.Ipv4.Network {
									nestedVrfBgpAdvertiseNetworkIpv4Network := VrfBgpAdvertiseNetworkIpv4Network{}
									if oVrfBgpAdvertiseNetworkIpv4Network.Misc != nil {
										entry.Misc["VrfBgpAdvertiseNetworkIpv4Network"] = oVrfBgpAdvertiseNetworkIpv4Network.Misc
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Name != "" {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Name = oVrfBgpAdvertiseNetworkIpv4Network.Name
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Unicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Unicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Unicast, nil)
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Multicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Multicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Multicast, nil)
									}
									if oVrfBgpAdvertiseNetworkIpv4Network.Backdoor != nil {
										nestedVrfBgpAdvertiseNetworkIpv4Network.Backdoor = util.AsBool(oVrfBgpAdvertiseNetworkIpv4Network.Backdoor, nil)
									}
									nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv4.Network, nestedVrfBgpAdvertiseNetworkIpv4Network)
								}
							}
						}
						if oVrf.Bgp.AdvertiseNetwork.Ipv6 != nil {
							nestedVrf.Bgp.AdvertiseNetwork.Ipv6 = &VrfBgpAdvertiseNetworkIpv6{}
							if oVrf.Bgp.AdvertiseNetwork.Ipv6.Misc != nil {
								entry.Misc["VrfBgpAdvertiseNetworkIpv6"] = oVrf.Bgp.AdvertiseNetwork.Ipv6.Misc
							}
							if oVrf.Bgp.AdvertiseNetwork.Ipv6.Network != nil {
								nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = []VrfBgpAdvertiseNetworkIpv6Network{}
								for _, oVrfBgpAdvertiseNetworkIpv6Network := range oVrf.Bgp.AdvertiseNetwork.Ipv6.Network {
									nestedVrfBgpAdvertiseNetworkIpv6Network := VrfBgpAdvertiseNetworkIpv6Network{}
									if oVrfBgpAdvertiseNetworkIpv6Network.Misc != nil {
										entry.Misc["VrfBgpAdvertiseNetworkIpv6Network"] = oVrfBgpAdvertiseNetworkIpv6Network.Misc
									}
									if oVrfBgpAdvertiseNetworkIpv6Network.Name != "" {
										nestedVrfBgpAdvertiseNetworkIpv6Network.Name = oVrfBgpAdvertiseNetworkIpv6Network.Name
									}
									if oVrfBgpAdvertiseNetworkIpv6Network.Unicast != nil {
										nestedVrfBgpAdvertiseNetworkIpv6Network.Unicast = util.AsBool(oVrfBgpAdvertiseNetworkIpv6Network.Unicast, nil)
									}
									nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network = append(nestedVrf.Bgp.AdvertiseNetwork.Ipv6.Network, nestedVrfBgpAdvertiseNetworkIpv6Network)
								}
							}
						}
					}
					if oVrf.Bgp.PeerGroup != nil {
						nestedVrf.Bgp.PeerGroup = []VrfBgpPeerGroup{}
						for _, oVrfBgpPeerGroup := range oVrf.Bgp.PeerGroup {
							nestedVrfBgpPeerGroup := VrfBgpPeerGroup{}
							if oVrfBgpPeerGroup.Misc != nil {
								entry.Misc["VrfBgpPeerGroup"] = oVrfBgpPeerGroup.Misc
							}
							if oVrfBgpPeerGroup.Name != "" {
								nestedVrfBgpPeerGroup.Name = oVrfBgpPeerGroup.Name
							}
							if oVrfBgpPeerGroup.Enable != nil {
								nestedVrfBgpPeerGroup.Enable = util.AsBool(oVrfBgpPeerGroup.Enable, nil)
							}
							if oVrfBgpPeerGroup.Type != nil {
								nestedVrfBgpPeerGroup.Type = &VrfBgpPeerGroupType{}
								if oVrfBgpPeerGroup.Type.Misc != nil {
									entry.Misc["VrfBgpPeerGroupType"] = oVrfBgpPeerGroup.Type.Misc
								}
								if oVrfBgpPeerGroup.Type.Ibgp != nil {
									nestedVrfBgpPeerGroup.Type.Ibgp = &VrfBgpPeerGroupTypeIbgp{}
									if oVrfBgpPeerGroup.Type.Ibgp.Misc != nil {
										entry.Misc["VrfBgpPeerGroupTypeIbgp"] = oVrfBgpPeerGroup.Type.Ibgp.Misc
									}
								}
								if oVrfBgpPeerGroup.Type.Ebgp != nil {
									nestedVrfBgpPeerGroup.Type.Ebgp = &VrfBgpPeerGroupTypeEbgp{}
									if oVrfBgpPeerGroup.Type.Ebgp.Misc != nil {
										entry.Misc["VrfBgpPeerGroupTypeEbgp"] = oVrfBgpPeerGroup.Type.Ebgp.Misc
									}
								}
							}
							if oVrfBgpPeerGroup.AddressFamily != nil {
								nestedVrfBgpPeerGroup.AddressFamily = &VrfBgpPeerGroupAddressFamily{}
								if oVrfBgpPeerGroup.AddressFamily.Misc != nil {
									entry.Misc["VrfBgpPeerGroupAddressFamily"] = oVrfBgpPeerGroup.AddressFamily.Misc
								}
								if oVrfBgpPeerGroup.AddressFamily.Ipv4 != nil {
									nestedVrfBgpPeerGroup.AddressFamily.Ipv4 = oVrfBgpPeerGroup.AddressFamily.Ipv4
								}
								if oVrfBgpPeerGroup.AddressFamily.Ipv6 != nil {
									nestedVrfBgpPeerGroup.AddressFamily.Ipv6 = oVrfBgpPeerGroup.AddressFamily.Ipv6
								}
							}
							if oVrfBgpPeerGroup.FilteringProfile != nil {
								nestedVrfBgpPeerGroup.FilteringProfile = &VrfBgpPeerGroupFilteringProfile{}
								if oVrfBgpPeerGroup.FilteringProfile.Misc != nil {
									entry.Misc["VrfBgpPeerGroupFilteringProfile"] = oVrfBgpPeerGroup.FilteringProfile.Misc
								}
								if oVrfBgpPeerGroup.FilteringProfile.Ipv4 != nil {
									nestedVrfBgpPeerGroup.FilteringProfile.Ipv4 = oVrfBgpPeerGroup.FilteringProfile.Ipv4
								}
								if oVrfBgpPeerGroup.FilteringProfile.Ipv6 != nil {
									nestedVrfBgpPeerGroup.FilteringProfile.Ipv6 = oVrfBgpPeerGroup.FilteringProfile.Ipv6
								}
							}
							if oVrfBgpPeerGroup.ConnectionOptions != nil {
								nestedVrfBgpPeerGroup.ConnectionOptions = &VrfBgpPeerGroupConnectionOptions{}
								if oVrfBgpPeerGroup.ConnectionOptions.Misc != nil {
									entry.Misc["VrfBgpPeerGroupConnectionOptions"] = oVrfBgpPeerGroup.ConnectionOptions.Misc
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Timers != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Timers = oVrfBgpPeerGroup.ConnectionOptions.Timers
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Multihop != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Multihop = oVrfBgpPeerGroup.ConnectionOptions.Multihop
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Authentication != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Authentication = oVrfBgpPeerGroup.ConnectionOptions.Authentication
								}
								if oVrfBgpPeerGroup.ConnectionOptions.Dampening != nil {
									nestedVrfBgpPeerGroup.ConnectionOptions.Dampening = oVrfBgpPeerGroup.ConnectionOptions.Dampening
								}
							}
							if oVrfBgpPeerGroup.Peer != nil {
								nestedVrfBgpPeerGroup.Peer = []VrfBgpPeerGroupPeer{}
								for _, oVrfBgpPeerGroupPeer := range oVrfBgpPeerGroup.Peer {
									nestedVrfBgpPeerGroupPeer := VrfBgpPeerGroupPeer{}
									if oVrfBgpPeerGroupPeer.Misc != nil {
										entry.Misc["VrfBgpPeerGroupPeer"] = oVrfBgpPeerGroupPeer.Misc
									}
									if oVrfBgpPeerGroupPeer.Name != "" {
										nestedVrfBgpPeerGroupPeer.Name = oVrfBgpPeerGroupPeer.Name
									}
									if oVrfBgpPeerGroupPeer.Enable != nil {
										nestedVrfBgpPeerGroupPeer.Enable = util.AsBool(oVrfBgpPeerGroupPeer.Enable, nil)
									}
									if oVrfBgpPeerGroupPeer.Passive != nil {
										nestedVrfBgpPeerGroupPeer.Passive = util.AsBool(oVrfBgpPeerGroupPeer.Passive, nil)
									}
									if oVrfBgpPeerGroupPeer.PeerAs != nil {
										nestedVrfBgpPeerGroupPeer.PeerAs = oVrfBgpPeerGroupPeer.PeerAs
									}
									if oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
										nestedVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.AsBool(oVrfBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
									}
									if oVrfBgpPeerGroupPeer.Inherit != nil {
										nestedVrfBgpPeerGroupPeer.Inherit = &VrfBgpPeerGroupPeerInherit{}
										if oVrfBgpPeerGroupPeer.Inherit.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerInherit"] = oVrfBgpPeerGroupPeer.Inherit.Misc
										}
										if oVrfBgpPeerGroupPeer.Inherit.Yes != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.Yes = &VrfBgpPeerGroupPeerInheritYes{}
											if oVrfBgpPeerGroupPeer.Inherit.Yes.Misc != nil {
												entry.Misc["VrfBgpPeerGroupPeerInheritYes"] = oVrfBgpPeerGroupPeer.Inherit.Yes.Misc
											}
										}
										if oVrfBgpPeerGroupPeer.Inherit.No != nil {
											nestedVrfBgpPeerGroupPeer.Inherit.No = &VrfBgpPeerGroupPeerInheritNo{}
											if oVrfBgpPeerGroupPeer.Inherit.No.Misc != nil {
												entry.Misc["VrfBgpPeerGroupPeerInheritNo"] = oVrfBgpPeerGroupPeer.Inherit.No.Misc
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily = &VrfBgpPeerGroupPeerInheritNoAddressFamily{}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc != nil {
													entry.Misc["VrfBgpPeerGroupPeerInheritNoAddressFamily"] = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Misc
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv4
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.AddressFamily.Ipv6
												}
											}
											if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile != nil {
												nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile = &VrfBgpPeerGroupPeerInheritNoFilteringProfile{}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc != nil {
													entry.Misc["VrfBgpPeerGroupPeerInheritNoFilteringProfile"] = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Misc
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv4
												}
												if oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 != nil {
													nestedVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6 = oVrfBgpPeerGroupPeer.Inherit.No.FilteringProfile.Ipv6
												}
											}
										}
									}
									if oVrfBgpPeerGroupPeer.LocalAddress != nil {
										nestedVrfBgpPeerGroupPeer.LocalAddress = &VrfBgpPeerGroupPeerLocalAddress{}
										if oVrfBgpPeerGroupPeer.LocalAddress.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerLocalAddress"] = oVrfBgpPeerGroupPeer.LocalAddress.Misc
										}
										if oVrfBgpPeerGroupPeer.LocalAddress.Interface != nil {
											nestedVrfBgpPeerGroupPeer.LocalAddress.Interface = oVrfBgpPeerGroupPeer.LocalAddress.Interface
										}
										if oVrfBgpPeerGroupPeer.LocalAddress.Ip != nil {
											nestedVrfBgpPeerGroupPeer.LocalAddress.Ip = oVrfBgpPeerGroupPeer.LocalAddress.Ip
										}
									}
									if oVrfBgpPeerGroupPeer.PeerAddress != nil {
										nestedVrfBgpPeerGroupPeer.PeerAddress = &VrfBgpPeerGroupPeerPeerAddress{}
										if oVrfBgpPeerGroupPeer.PeerAddress.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerPeerAddress"] = oVrfBgpPeerGroupPeer.PeerAddress.Misc
										}
										if oVrfBgpPeerGroupPeer.PeerAddress.Ip != nil {
											nestedVrfBgpPeerGroupPeer.PeerAddress.Ip = oVrfBgpPeerGroupPeer.PeerAddress.Ip
										}
										if oVrfBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
											nestedVrfBgpPeerGroupPeer.PeerAddress.Fqdn = oVrfBgpPeerGroupPeer.PeerAddress.Fqdn
										}
									}
									if oVrfBgpPeerGroupPeer.ConnectionOptions != nil {
										nestedVrfBgpPeerGroupPeer.ConnectionOptions = &VrfBgpPeerGroupPeerConnectionOptions{}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerConnectionOptions"] = oVrfBgpPeerGroupPeer.ConnectionOptions.Misc
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Timers != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Timers = oVrfBgpPeerGroupPeer.ConnectionOptions.Timers
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Multihop = oVrfBgpPeerGroupPeer.ConnectionOptions.Multihop
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Authentication = oVrfBgpPeerGroupPeer.ConnectionOptions.Authentication
										}
										if oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening != nil {
											nestedVrfBgpPeerGroupPeer.ConnectionOptions.Dampening = oVrfBgpPeerGroupPeer.ConnectionOptions.Dampening
										}
									}
									if oVrfBgpPeerGroupPeer.Bfd != nil {
										nestedVrfBgpPeerGroupPeer.Bfd = &VrfBgpPeerGroupPeerBfd{}
										if oVrfBgpPeerGroupPeer.Bfd.Misc != nil {
											entry.Misc["VrfBgpPeerGroupPeerBfd"] = oVrfBgpPeerGroupPeer.Bfd.Misc
										}
										if oVrfBgpPeerGroupPeer.Bfd.Profile != nil {
											nestedVrfBgpPeerGroupPeer.Bfd.Profile = oVrfBgpPeerGroupPeer.Bfd.Profile
										}
									}
									nestedVrfBgpPeerGroup.Peer = append(nestedVrfBgpPeerGroup.Peer, nestedVrfBgpPeerGroupPeer)
								}
							}
							nestedVrf.Bgp.PeerGroup = append(nestedVrf.Bgp.PeerGroup, nestedVrfBgpPeerGroup)
						}
					}
					if oVrf.Bgp.AggregateRoutes != nil {
						nestedVrf.Bgp.AggregateRoutes = []VrfBgpAggregateRoutes{}
						for _, oVrfBgpAggregateRoutes := range oVrf.Bgp.AggregateRoutes {
							nestedVrfBgpAggregateRoutes := VrfBgpAggregateRoutes{}
							if oVrfBgpAggregateRoutes.Misc != nil {
								entry.Misc["VrfBgpAggregateRoutes"] = oVrfBgpAggregateRoutes.Misc
							}
							if oVrfBgpAggregateRoutes.Name != "" {
								nestedVrfBgpAggregateRoutes.Name = oVrfBgpAggregateRoutes.Name
							}
							if oVrfBgpAggregateRoutes.Description != nil {
								nestedVrfBgpAggregateRoutes.Description = oVrfBgpAggregateRoutes.Description
							}
							if oVrfBgpAggregateRoutes.Enable != nil {
								nestedVrfBgpAggregateRoutes.Enable = util.AsBool(oVrfBgpAggregateRoutes.Enable, nil)
							}
							if oVrfBgpAggregateRoutes.SummaryOnly != nil {
								nestedVrfBgpAggregateRoutes.SummaryOnly = util.AsBool(oVrfBgpAggregateRoutes.SummaryOnly, nil)
							}
							if oVrfBgpAggregateRoutes.AsSet != nil {
								nestedVrfBgpAggregateRoutes.AsSet = util.AsBool(oVrfBgpAggregateRoutes.AsSet, nil)
							}
							if oVrfBgpAggregateRoutes.SameMed != nil {
								nestedVrfBgpAggregateRoutes.SameMed = util.AsBool(oVrfBgpAggregateRoutes.SameMed, nil)
							}
							if oVrfBgpAggregateRoutes.Type != nil {
								nestedVrfBgpAggregateRoutes.Type = &VrfBgpAggregateRoutesType{}
								if oVrfBgpAggregateRoutes.Type.Misc != nil {
									entry.Misc["VrfBgpAggregateRoutesType"] = oVrfBgpAggregateRoutes.Type.Misc
								}
								if oVrfBgpAggregateRoutes.Type.Ipv4 != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv4 = &VrfBgpAggregateRoutesTypeIpv4{}
									if oVrfBgpAggregateRoutes.Type.Ipv4.Misc != nil {
										entry.Misc["VrfBgpAggregateRoutesTypeIpv4"] = oVrfBgpAggregateRoutes.Type.Ipv4.Misc
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv4.SummaryPrefix
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv4.SuppressMap
									}
									if oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv4.AttributeMap
									}
								}
								if oVrfBgpAggregateRoutes.Type.Ipv6 != nil {
									nestedVrfBgpAggregateRoutes.Type.Ipv6 = &VrfBgpAggregateRoutesTypeIpv6{}
									if oVrfBgpAggregateRoutes.Type.Ipv6.Misc != nil {
										entry.Misc["VrfBgpAggregateRoutesTypeIpv6"] = oVrfBgpAggregateRoutes.Type.Ipv6.Misc
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix = oVrfBgpAggregateRoutes.Type.Ipv6.SummaryPrefix
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap = oVrfBgpAggregateRoutes.Type.Ipv6.SuppressMap
									}
									if oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap != nil {
										nestedVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap = oVrfBgpAggregateRoutes.Type.Ipv6.AttributeMap
									}
								}
							}
							nestedVrf.Bgp.AggregateRoutes = append(nestedVrf.Bgp.AggregateRoutes, nestedVrfBgpAggregateRoutes)
						}
					}
				}
				if oVrf.RoutingTable != nil {
					nestedVrf.RoutingTable = &VrfRoutingTable{}
					if oVrf.RoutingTable.Misc != nil {
						entry.Misc["VrfRoutingTable"] = oVrf.RoutingTable.Misc
					}
					if oVrf.RoutingTable.Ip != nil {
						nestedVrf.RoutingTable.Ip = &VrfRoutingTableIp{}
						if oVrf.RoutingTable.Ip.Misc != nil {
							entry.Misc["VrfRoutingTableIp"] = oVrf.RoutingTable.Ip.Misc
						}
						if oVrf.RoutingTable.Ip.StaticRoute != nil {
							nestedVrf.RoutingTable.Ip.StaticRoute = []VrfRoutingTableIpStaticRoute{}
							for _, oVrfRoutingTableIpStaticRoute := range oVrf.RoutingTable.Ip.StaticRoute {
								nestedVrfRoutingTableIpStaticRoute := VrfRoutingTableIpStaticRoute{}
								if oVrfRoutingTableIpStaticRoute.Misc != nil {
									entry.Misc["VrfRoutingTableIpStaticRoute"] = oVrfRoutingTableIpStaticRoute.Misc
								}
								if oVrfRoutingTableIpStaticRoute.Name != "" {
									nestedVrfRoutingTableIpStaticRoute.Name = oVrfRoutingTableIpStaticRoute.Name
								}
								if oVrfRoutingTableIpStaticRoute.Destination != nil {
									nestedVrfRoutingTableIpStaticRoute.Destination = oVrfRoutingTableIpStaticRoute.Destination
								}
								if oVrfRoutingTableIpStaticRoute.Interface != nil {
									nestedVrfRoutingTableIpStaticRoute.Interface = oVrfRoutingTableIpStaticRoute.Interface
								}
								if oVrfRoutingTableIpStaticRoute.AdminDist != nil {
									nestedVrfRoutingTableIpStaticRoute.AdminDist = oVrfRoutingTableIpStaticRoute.AdminDist
								}
								if oVrfRoutingTableIpStaticRoute.Metric != nil {
									nestedVrfRoutingTableIpStaticRoute.Metric = oVrfRoutingTableIpStaticRoute.Metric
								}
								if oVrfRoutingTableIpStaticRoute.Nexthop != nil {
									nestedVrfRoutingTableIpStaticRoute.Nexthop = &VrfRoutingTableIpStaticRouteNexthop{}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRouteNexthop"] = oVrfRoutingTableIpStaticRoute.Nexthop.Misc
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Discard != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Discard = &VrfRoutingTableIpStaticRouteNexthopDiscard{}
										if oVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc != nil {
											entry.Misc["VrfRoutingTableIpStaticRouteNexthopDiscard"] = oVrfRoutingTableIpStaticRoute.Nexthop.Discard.Misc
										}
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.IpAddress = oVrfRoutingTableIpStaticRoute.Nexthop.IpAddress
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.NextLr != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.NextLr = oVrfRoutingTableIpStaticRoute.Nexthop.NextLr
									}
									if oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
										nestedVrfRoutingTableIpStaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpStaticRoute.Nexthop.Fqdn
									}
								}
								if oVrfRoutingTableIpStaticRoute.Bfd != nil {
									nestedVrfRoutingTableIpStaticRoute.Bfd = &VrfRoutingTableIpStaticRouteBfd{}
									if oVrfRoutingTableIpStaticRoute.Bfd.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRouteBfd"] = oVrfRoutingTableIpStaticRoute.Bfd.Misc
									}
									if oVrfRoutingTableIpStaticRoute.Bfd.Profile != nil {
										nestedVrfRoutingTableIpStaticRoute.Bfd.Profile = oVrfRoutingTableIpStaticRoute.Bfd.Profile
									}
								}
								if oVrfRoutingTableIpStaticRoute.PathMonitor != nil {
									nestedVrfRoutingTableIpStaticRoute.PathMonitor = &VrfRoutingTableIpStaticRoutePathMonitor{}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.Misc != nil {
										entry.Misc["VrfRoutingTableIpStaticRoutePathMonitor"] = oVrfRoutingTableIpStaticRoute.PathMonitor.Misc
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.Enable = util.AsBool(oVrfRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpStaticRoute.PathMonitor.FailureCondition
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpStaticRoute.PathMonitor.HoldTime
									}
									if oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
										nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
										for _, oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
											nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc != nil {
												entry.Misc["VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations"] = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
											}
											if oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
												nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
											}
											nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
										}
									}
								}
								nestedVrf.RoutingTable.Ip.StaticRoute = append(nestedVrf.RoutingTable.Ip.StaticRoute, nestedVrfRoutingTableIpStaticRoute)
							}
						}
					}
					if oVrf.RoutingTable.Ipv6 != nil {
						nestedVrf.RoutingTable.Ipv6 = &VrfRoutingTableIpv6{}
						if oVrf.RoutingTable.Ipv6.Misc != nil {
							entry.Misc["VrfRoutingTableIpv6"] = oVrf.RoutingTable.Ipv6.Misc
						}
						if oVrf.RoutingTable.Ipv6.StaticRoute != nil {
							nestedVrf.RoutingTable.Ipv6.StaticRoute = []VrfRoutingTableIpv6StaticRoute{}
							for _, oVrfRoutingTableIpv6StaticRoute := range oVrf.RoutingTable.Ipv6.StaticRoute {
								nestedVrfRoutingTableIpv6StaticRoute := VrfRoutingTableIpv6StaticRoute{}
								if oVrfRoutingTableIpv6StaticRoute.Misc != nil {
									entry.Misc["VrfRoutingTableIpv6StaticRoute"] = oVrfRoutingTableIpv6StaticRoute.Misc
								}
								if oVrfRoutingTableIpv6StaticRoute.Name != "" {
									nestedVrfRoutingTableIpv6StaticRoute.Name = oVrfRoutingTableIpv6StaticRoute.Name
								}
								if oVrfRoutingTableIpv6StaticRoute.Destination != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Destination = oVrfRoutingTableIpv6StaticRoute.Destination
								}
								if oVrfRoutingTableIpv6StaticRoute.Interface != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Interface = oVrfRoutingTableIpv6StaticRoute.Interface
								}
								if oVrfRoutingTableIpv6StaticRoute.AdminDist != nil {
									nestedVrfRoutingTableIpv6StaticRoute.AdminDist = oVrfRoutingTableIpv6StaticRoute.AdminDist
								}
								if oVrfRoutingTableIpv6StaticRoute.Metric != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Metric = oVrfRoutingTableIpv6StaticRoute.Metric
								}
								if oVrfRoutingTableIpv6StaticRoute.Nexthop != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Nexthop = &VrfRoutingTableIpv6StaticRouteNexthop{}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRouteNexthop"] = oVrfRoutingTableIpv6StaticRoute.Nexthop.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Discard = &VrfRoutingTableIpv6StaticRouteNexthopDiscard{}
										if oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc != nil {
											entry.Misc["VrfRoutingTableIpv6StaticRouteNexthopDiscard"] = oVrfRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc
										}
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oVrfRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn = oVrfRoutingTableIpv6StaticRoute.Nexthop.Fqdn
									}
									if oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr = oVrfRoutingTableIpv6StaticRoute.Nexthop.NextLr
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.Bfd != nil {
									nestedVrfRoutingTableIpv6StaticRoute.Bfd = &VrfRoutingTableIpv6StaticRouteBfd{}
									if oVrfRoutingTableIpv6StaticRoute.Bfd.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRouteBfd"] = oVrfRoutingTableIpv6StaticRoute.Bfd.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
										nestedVrfRoutingTableIpv6StaticRoute.Bfd.Profile = oVrfRoutingTableIpv6StaticRoute.Bfd.Profile
									}
								}
								if oVrfRoutingTableIpv6StaticRoute.PathMonitor != nil {
									nestedVrfRoutingTableIpv6StaticRoute.PathMonitor = &VrfRoutingTableIpv6StaticRoutePathMonitor{}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc != nil {
										entry.Misc["VrfRoutingTableIpv6StaticRoutePathMonitor"] = oVrfRoutingTableIpv6StaticRoute.PathMonitor.Misc
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.AsBool(oVrfRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oVrfRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oVrfRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
									}
									if oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
										nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
										for _, oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
											nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc != nil {
												entry.Misc["VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"] = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
											}
											if oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
												nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
											}
											nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedVrfRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
										}
									}
								}
								nestedVrf.RoutingTable.Ipv6.StaticRoute = append(nestedVrf.RoutingTable.Ipv6.StaticRoute, nestedVrfRoutingTableIpv6StaticRoute)
							}
						}
					}
				}
				if oVrf.Ospf != nil {
					nestedVrf.Ospf = &VrfOspf{}
					if oVrf.Ospf.Misc != nil {
						entry.Misc["VrfOspf"] = oVrf.Ospf.Misc
					}
					if oVrf.Ospf.RouterId != nil {
						nestedVrf.Ospf.RouterId = oVrf.Ospf.RouterId
					}
					if oVrf.Ospf.Enable != nil {
						nestedVrf.Ospf.Enable = util.AsBool(oVrf.Ospf.Enable, nil)
					}
					if oVrf.Ospf.Rfc1583 != nil {
						nestedVrf.Ospf.Rfc1583 = util.AsBool(oVrf.Ospf.Rfc1583, nil)
					}
					if oVrf.Ospf.SpfTimer != nil {
						nestedVrf.Ospf.SpfTimer = oVrf.Ospf.SpfTimer
					}
					if oVrf.Ospf.GlobalIfTimer != nil {
						nestedVrf.Ospf.GlobalIfTimer = oVrf.Ospf.GlobalIfTimer
					}
					if oVrf.Ospf.RedistributionProfile != nil {
						nestedVrf.Ospf.RedistributionProfile = oVrf.Ospf.RedistributionProfile
					}
					if oVrf.Ospf.GlobalBfd != nil {
						nestedVrf.Ospf.GlobalBfd = &VrfOspfGlobalBfd{}
						if oVrf.Ospf.GlobalBfd.Misc != nil {
							entry.Misc["VrfOspfGlobalBfd"] = oVrf.Ospf.GlobalBfd.Misc
						}
						if oVrf.Ospf.GlobalBfd.Profile != nil {
							nestedVrf.Ospf.GlobalBfd.Profile = oVrf.Ospf.GlobalBfd.Profile
						}
					}
					if oVrf.Ospf.GracefulRestart != nil {
						nestedVrf.Ospf.GracefulRestart = &VrfOspfGracefulRestart{}
						if oVrf.Ospf.GracefulRestart.Misc != nil {
							entry.Misc["VrfOspfGracefulRestart"] = oVrf.Ospf.GracefulRestart.Misc
						}
						if oVrf.Ospf.GracefulRestart.Enable != nil {
							nestedVrf.Ospf.GracefulRestart.Enable = util.AsBool(oVrf.Ospf.GracefulRestart.Enable, nil)
						}
						if oVrf.Ospf.GracefulRestart.GracePeriod != nil {
							nestedVrf.Ospf.GracefulRestart.GracePeriod = oVrf.Ospf.GracefulRestart.GracePeriod
						}
						if oVrf.Ospf.GracefulRestart.HelperEnable != nil {
							nestedVrf.Ospf.GracefulRestart.HelperEnable = util.AsBool(oVrf.Ospf.GracefulRestart.HelperEnable, nil)
						}
						if oVrf.Ospf.GracefulRestart.StrictLSAChecking != nil {
							nestedVrf.Ospf.GracefulRestart.StrictLSAChecking = util.AsBool(oVrf.Ospf.GracefulRestart.StrictLSAChecking, nil)
						}
						if oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
							nestedVrf.Ospf.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospf.GracefulRestart.MaxNeighborRestartTime
						}
					}
					if oVrf.Ospf.Area != nil {
						nestedVrf.Ospf.Area = []VrfOspfArea{}
						for _, oVrfOspfArea := range oVrf.Ospf.Area {
							nestedVrfOspfArea := VrfOspfArea{}
							if oVrfOspfArea.Misc != nil {
								entry.Misc["VrfOspfArea"] = oVrfOspfArea.Misc
							}
							if oVrfOspfArea.Name != "" {
								nestedVrfOspfArea.Name = oVrfOspfArea.Name
							}
							if oVrfOspfArea.Authentication != nil {
								nestedVrfOspfArea.Authentication = oVrfOspfArea.Authentication
							}
							if oVrfOspfArea.Type != nil {
								nestedVrfOspfArea.Type = &VrfOspfAreaType{}
								if oVrfOspfArea.Type.Misc != nil {
									entry.Misc["VrfOspfAreaType"] = oVrfOspfArea.Type.Misc
								}
								if oVrfOspfArea.Type.Normal != nil {
									nestedVrfOspfArea.Type.Normal = &VrfOspfAreaTypeNormal{}
									if oVrfOspfArea.Type.Normal.Misc != nil {
										entry.Misc["VrfOspfAreaTypeNormal"] = oVrfOspfArea.Type.Normal.Misc
									}
									if oVrfOspfArea.Type.Normal.Abr != nil {
										nestedVrfOspfArea.Type.Normal.Abr = &VrfOspfAreaTypeNormalAbr{}
										if oVrfOspfArea.Type.Normal.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNormalAbr"] = oVrfOspfArea.Type.Normal.Abr.Misc
										}
										if oVrfOspfArea.Type.Normal.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.ImportList = oVrfOspfArea.Type.Normal.Abr.ImportList
										}
										if oVrfOspfArea.Type.Normal.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.ExportList = oVrfOspfArea.Type.Normal.Abr.ExportList
										}
										if oVrfOspfArea.Type.Normal.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.InboundFilterList = oVrfOspfArea.Type.Normal.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Normal.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Normal.Abr.OutboundFilterList = oVrfOspfArea.Type.Normal.Abr.OutboundFilterList
										}
									}
								}
								if oVrfOspfArea.Type.Stub != nil {
									nestedVrfOspfArea.Type.Stub = &VrfOspfAreaTypeStub{}
									if oVrfOspfArea.Type.Stub.Misc != nil {
										entry.Misc["VrfOspfAreaTypeStub"] = oVrfOspfArea.Type.Stub.Misc
									}
									if oVrfOspfArea.Type.Stub.NoSummary != nil {
										nestedVrfOspfArea.Type.Stub.NoSummary = util.AsBool(oVrfOspfArea.Type.Stub.NoSummary, nil)
									}
									if oVrfOspfArea.Type.Stub.Abr != nil {
										nestedVrfOspfArea.Type.Stub.Abr = &VrfOspfAreaTypeStubAbr{}
										if oVrfOspfArea.Type.Stub.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeStubAbr"] = oVrfOspfArea.Type.Stub.Abr.Misc
										}
										if oVrfOspfArea.Type.Stub.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.ImportList = oVrfOspfArea.Type.Stub.Abr.ImportList
										}
										if oVrfOspfArea.Type.Stub.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.ExportList = oVrfOspfArea.Type.Stub.Abr.ExportList
										}
										if oVrfOspfArea.Type.Stub.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.InboundFilterList = oVrfOspfArea.Type.Stub.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Stub.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Stub.Abr.OutboundFilterList = oVrfOspfArea.Type.Stub.Abr.OutboundFilterList
										}
									}
									if oVrfOspfArea.Type.Stub.DefaultRouteMetric != nil {
										nestedVrfOspfArea.Type.Stub.DefaultRouteMetric = oVrfOspfArea.Type.Stub.DefaultRouteMetric
									}
								}
								if oVrfOspfArea.Type.Nssa != nil {
									nestedVrfOspfArea.Type.Nssa = &VrfOspfAreaTypeNssa{}
									if oVrfOspfArea.Type.Nssa.Misc != nil {
										entry.Misc["VrfOspfAreaTypeNssa"] = oVrfOspfArea.Type.Nssa.Misc
									}
									if oVrfOspfArea.Type.Nssa.NoSummary != nil {
										nestedVrfOspfArea.Type.Nssa.NoSummary = util.AsBool(oVrfOspfArea.Type.Nssa.NoSummary, nil)
									}
									if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate != nil {
										nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate = &VrfOspfAreaTypeNssaDefaultInformationOriginate{}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNssaDefaultInformationOriginate"] = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Misc
										}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric != nil {
											nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.Metric
										}
										if oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
											nestedVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfArea.Type.Nssa.DefaultInformationOriginate.MetricType
										}
									}
									if oVrfOspfArea.Type.Nssa.Abr != nil {
										nestedVrfOspfArea.Type.Nssa.Abr = &VrfOspfAreaTypeNssaAbr{}
										if oVrfOspfArea.Type.Nssa.Abr.Misc != nil {
											entry.Misc["VrfOspfAreaTypeNssaAbr"] = oVrfOspfArea.Type.Nssa.Abr.Misc
										}
										if oVrfOspfArea.Type.Nssa.Abr.ImportList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.ImportList = oVrfOspfArea.Type.Nssa.Abr.ImportList
										}
										if oVrfOspfArea.Type.Nssa.Abr.ExportList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.ExportList = oVrfOspfArea.Type.Nssa.Abr.ExportList
										}
										if oVrfOspfArea.Type.Nssa.Abr.InboundFilterList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.InboundFilterList = oVrfOspfArea.Type.Nssa.Abr.InboundFilterList
										}
										if oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.OutboundFilterList = oVrfOspfArea.Type.Nssa.Abr.OutboundFilterList
										}
										if oVrfOspfArea.Type.Nssa.Abr.NssaExtRange != nil {
											nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = []VrfOspfAreaTypeNssaAbrNssaExtRange{}
											for _, oVrfOspfAreaTypeNssaAbrNssaExtRange := range oVrfOspfArea.Type.Nssa.Abr.NssaExtRange {
												nestedVrfOspfAreaTypeNssaAbrNssaExtRange := VrfOspfAreaTypeNssaAbrNssaExtRange{}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Misc != nil {
													entry.Misc["VrfOspfAreaTypeNssaAbrNssaExtRange"] = oVrfOspfAreaTypeNssaAbrNssaExtRange.Misc
												}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Name != "" {
													nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Name = oVrfOspfAreaTypeNssaAbrNssaExtRange.Name
												}
												if oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise != nil {
													nestedVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise = util.AsBool(oVrfOspfAreaTypeNssaAbrNssaExtRange.Advertise, nil)
												}
												nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfArea.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfAreaTypeNssaAbrNssaExtRange)
											}
										}
									}
								}
							}
							if oVrfOspfArea.Range != nil {
								nestedVrfOspfArea.Range = []VrfOspfAreaRange{}
								for _, oVrfOspfAreaRange := range oVrfOspfArea.Range {
									nestedVrfOspfAreaRange := VrfOspfAreaRange{}
									if oVrfOspfAreaRange.Misc != nil {
										entry.Misc["VrfOspfAreaRange"] = oVrfOspfAreaRange.Misc
									}
									if oVrfOspfAreaRange.Name != "" {
										nestedVrfOspfAreaRange.Name = oVrfOspfAreaRange.Name
									}
									if oVrfOspfAreaRange.Advertise != nil {
										nestedVrfOspfAreaRange.Advertise = util.AsBool(oVrfOspfAreaRange.Advertise, nil)
									}
									nestedVrfOspfArea.Range = append(nestedVrfOspfArea.Range, nestedVrfOspfAreaRange)
								}
							}
							if oVrfOspfArea.Interface != nil {
								nestedVrfOspfArea.Interface = []VrfOspfAreaInterface{}
								for _, oVrfOspfAreaInterface := range oVrfOspfArea.Interface {
									nestedVrfOspfAreaInterface := VrfOspfAreaInterface{}
									if oVrfOspfAreaInterface.Misc != nil {
										entry.Misc["VrfOspfAreaInterface"] = oVrfOspfAreaInterface.Misc
									}
									if oVrfOspfAreaInterface.Name != "" {
										nestedVrfOspfAreaInterface.Name = oVrfOspfAreaInterface.Name
									}
									if oVrfOspfAreaInterface.Enable != nil {
										nestedVrfOspfAreaInterface.Enable = util.AsBool(oVrfOspfAreaInterface.Enable, nil)
									}
									if oVrfOspfAreaInterface.MtuIgnore != nil {
										nestedVrfOspfAreaInterface.MtuIgnore = util.AsBool(oVrfOspfAreaInterface.MtuIgnore, nil)
									}
									if oVrfOspfAreaInterface.Passive != nil {
										nestedVrfOspfAreaInterface.Passive = util.AsBool(oVrfOspfAreaInterface.Passive, nil)
									}
									if oVrfOspfAreaInterface.Priority != nil {
										nestedVrfOspfAreaInterface.Priority = oVrfOspfAreaInterface.Priority
									}
									if oVrfOspfAreaInterface.Metric != nil {
										nestedVrfOspfAreaInterface.Metric = oVrfOspfAreaInterface.Metric
									}
									if oVrfOspfAreaInterface.Authentication != nil {
										nestedVrfOspfAreaInterface.Authentication = oVrfOspfAreaInterface.Authentication
									}
									if oVrfOspfAreaInterface.Timing != nil {
										nestedVrfOspfAreaInterface.Timing = oVrfOspfAreaInterface.Timing
									}
									if oVrfOspfAreaInterface.LinkType != nil {
										nestedVrfOspfAreaInterface.LinkType = &VrfOspfAreaInterfaceLinkType{}
										if oVrfOspfAreaInterface.LinkType.Misc != nil {
											entry.Misc["VrfOspfAreaInterfaceLinkType"] = oVrfOspfAreaInterface.LinkType.Misc
										}
										if oVrfOspfAreaInterface.LinkType.Broadcast != nil {
											nestedVrfOspfAreaInterface.LinkType.Broadcast = &VrfOspfAreaInterfaceLinkTypeBroadcast{}
											if oVrfOspfAreaInterface.LinkType.Broadcast.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeBroadcast"] = oVrfOspfAreaInterface.LinkType.Broadcast.Misc
											}
										}
										if oVrfOspfAreaInterface.LinkType.P2p != nil {
											nestedVrfOspfAreaInterface.LinkType.P2p = &VrfOspfAreaInterfaceLinkTypeP2p{}
											if oVrfOspfAreaInterface.LinkType.P2p.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeP2p"] = oVrfOspfAreaInterface.LinkType.P2p.Misc
											}
										}
										if oVrfOspfAreaInterface.LinkType.P2mp != nil {
											nestedVrfOspfAreaInterface.LinkType.P2mp = &VrfOspfAreaInterfaceLinkTypeP2mp{}
											if oVrfOspfAreaInterface.LinkType.P2mp.Misc != nil {
												entry.Misc["VrfOspfAreaInterfaceLinkTypeP2mp"] = oVrfOspfAreaInterface.LinkType.P2mp.Misc
											}
											if oVrfOspfAreaInterface.LinkType.P2mp.Neighbor != nil {
												nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{}
												for _, oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfAreaInterface.LinkType.P2mp.Neighbor {
													nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor := VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc != nil {
														entry.Misc["VrfOspfAreaInterfaceLinkTypeP2mpNeighbor"] = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Misc
													}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
														nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Name
													}
													if oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
														nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfAreaInterfaceLinkTypeP2mpNeighbor.Priority
													}
													nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfAreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfAreaInterfaceLinkTypeP2mpNeighbor)
												}
											}
										}
									}
									if oVrfOspfAreaInterface.Bfd != nil {
										nestedVrfOspfAreaInterface.Bfd = &VrfOspfAreaInterfaceBfd{}
										if oVrfOspfAreaInterface.Bfd.Misc != nil {
											entry.Misc["VrfOspfAreaInterfaceBfd"] = oVrfOspfAreaInterface.Bfd.Misc
										}
										if oVrfOspfAreaInterface.Bfd.Profile != nil {
											nestedVrfOspfAreaInterface.Bfd.Profile = oVrfOspfAreaInterface.Bfd.Profile
										}
									}
									nestedVrfOspfArea.Interface = append(nestedVrfOspfArea.Interface, nestedVrfOspfAreaInterface)
								}
							}
							if oVrfOspfArea.VirtualLink != nil {
								nestedVrfOspfArea.VirtualLink = []VrfOspfAreaVirtualLink{}
								for _, oVrfOspfAreaVirtualLink := range oVrfOspfArea.VirtualLink {
									nestedVrfOspfAreaVirtualLink := VrfOspfAreaVirtualLink{}
									if oVrfOspfAreaVirtualLink.Misc != nil {
										entry.Misc["VrfOspfAreaVirtualLink"] = oVrfOspfAreaVirtualLink.Misc
									}
									if oVrfOspfAreaVirtualLink.Name != "" {
										nestedVrfOspfAreaVirtualLink.Name = oVrfOspfAreaVirtualLink.Name
									}
									if oVrfOspfAreaVirtualLink.NeighborId != nil {
										nestedVrfOspfAreaVirtualLink.NeighborId = oVrfOspfAreaVirtualLink.NeighborId
									}
									if oVrfOspfAreaVirtualLink.TransitAreaId != nil {
										nestedVrfOspfAreaVirtualLink.TransitAreaId = oVrfOspfAreaVirtualLink.TransitAreaId
									}
									if oVrfOspfAreaVirtualLink.Enable != nil {
										nestedVrfOspfAreaVirtualLink.Enable = util.AsBool(oVrfOspfAreaVirtualLink.Enable, nil)
									}
									if oVrfOspfAreaVirtualLink.InstanceId != nil {
										nestedVrfOspfAreaVirtualLink.InstanceId = oVrfOspfAreaVirtualLink.InstanceId
									}
									if oVrfOspfAreaVirtualLink.Timing != nil {
										nestedVrfOspfAreaVirtualLink.Timing = oVrfOspfAreaVirtualLink.Timing
									}
									if oVrfOspfAreaVirtualLink.Authentication != nil {
										nestedVrfOspfAreaVirtualLink.Authentication = oVrfOspfAreaVirtualLink.Authentication
									}
									if oVrfOspfAreaVirtualLink.Bfd != nil {
										nestedVrfOspfAreaVirtualLink.Bfd = &VrfOspfAreaVirtualLinkBfd{}
										if oVrfOspfAreaVirtualLink.Bfd.Misc != nil {
											entry.Misc["VrfOspfAreaVirtualLinkBfd"] = oVrfOspfAreaVirtualLink.Bfd.Misc
										}
										if oVrfOspfAreaVirtualLink.Bfd.Profile != nil {
											nestedVrfOspfAreaVirtualLink.Bfd.Profile = oVrfOspfAreaVirtualLink.Bfd.Profile
										}
									}
									nestedVrfOspfArea.VirtualLink = append(nestedVrfOspfArea.VirtualLink, nestedVrfOspfAreaVirtualLink)
								}
							}
							nestedVrf.Ospf.Area = append(nestedVrf.Ospf.Area, nestedVrfOspfArea)
						}
					}
				}
				if oVrf.Ospfv3 != nil {
					nestedVrf.Ospfv3 = &VrfOspfv3{}
					if oVrf.Ospfv3.Misc != nil {
						entry.Misc["VrfOspfv3"] = oVrf.Ospfv3.Misc
					}
					if oVrf.Ospfv3.Enable != nil {
						nestedVrf.Ospfv3.Enable = util.AsBool(oVrf.Ospfv3.Enable, nil)
					}
					if oVrf.Ospfv3.RouterId != nil {
						nestedVrf.Ospfv3.RouterId = oVrf.Ospfv3.RouterId
					}
					if oVrf.Ospfv3.DisableTransitTraffic != nil {
						nestedVrf.Ospfv3.DisableTransitTraffic = util.AsBool(oVrf.Ospfv3.DisableTransitTraffic, nil)
					}
					if oVrf.Ospfv3.SpfTimer != nil {
						nestedVrf.Ospfv3.SpfTimer = oVrf.Ospfv3.SpfTimer
					}
					if oVrf.Ospfv3.GlobalIfTimer != nil {
						nestedVrf.Ospfv3.GlobalIfTimer = oVrf.Ospfv3.GlobalIfTimer
					}
					if oVrf.Ospfv3.RedistributionProfile != nil {
						nestedVrf.Ospfv3.RedistributionProfile = oVrf.Ospfv3.RedistributionProfile
					}
					if oVrf.Ospfv3.GlobalBfd != nil {
						nestedVrf.Ospfv3.GlobalBfd = &VrfOspfv3GlobalBfd{}
						if oVrf.Ospfv3.GlobalBfd.Misc != nil {
							entry.Misc["VrfOspfv3GlobalBfd"] = oVrf.Ospfv3.GlobalBfd.Misc
						}
						if oVrf.Ospfv3.GlobalBfd.Profile != nil {
							nestedVrf.Ospfv3.GlobalBfd.Profile = oVrf.Ospfv3.GlobalBfd.Profile
						}
					}
					if oVrf.Ospfv3.GracefulRestart != nil {
						nestedVrf.Ospfv3.GracefulRestart = &VrfOspfv3GracefulRestart{}
						if oVrf.Ospfv3.GracefulRestart.Misc != nil {
							entry.Misc["VrfOspfv3GracefulRestart"] = oVrf.Ospfv3.GracefulRestart.Misc
						}
						if oVrf.Ospfv3.GracefulRestart.Enable != nil {
							nestedVrf.Ospfv3.GracefulRestart.Enable = util.AsBool(oVrf.Ospfv3.GracefulRestart.Enable, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.GracePeriod != nil {
							nestedVrf.Ospfv3.GracefulRestart.GracePeriod = oVrf.Ospfv3.GracefulRestart.GracePeriod
						}
						if oVrf.Ospfv3.GracefulRestart.HelperEnable != nil {
							nestedVrf.Ospfv3.GracefulRestart.HelperEnable = util.AsBool(oVrf.Ospfv3.GracefulRestart.HelperEnable, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
							nestedVrf.Ospfv3.GracefulRestart.StrictLSAChecking = util.AsBool(oVrf.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
						}
						if oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
							nestedVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime = oVrf.Ospfv3.GracefulRestart.MaxNeighborRestartTime
						}
					}
					if oVrf.Ospfv3.Area != nil {
						nestedVrf.Ospfv3.Area = []VrfOspfv3Area{}
						for _, oVrfOspfv3Area := range oVrf.Ospfv3.Area {
							nestedVrfOspfv3Area := VrfOspfv3Area{}
							if oVrfOspfv3Area.Misc != nil {
								entry.Misc["VrfOspfv3Area"] = oVrfOspfv3Area.Misc
							}
							if oVrfOspfv3Area.Name != "" {
								nestedVrfOspfv3Area.Name = oVrfOspfv3Area.Name
							}
							if oVrfOspfv3Area.Authentication != nil {
								nestedVrfOspfv3Area.Authentication = oVrfOspfv3Area.Authentication
							}
							if oVrfOspfv3Area.Type != nil {
								nestedVrfOspfv3Area.Type = &VrfOspfv3AreaType{}
								if oVrfOspfv3Area.Type.Misc != nil {
									entry.Misc["VrfOspfv3AreaType"] = oVrfOspfv3Area.Type.Misc
								}
								if oVrfOspfv3Area.Type.Normal != nil {
									nestedVrfOspfv3Area.Type.Normal = &VrfOspfv3AreaTypeNormal{}
									if oVrfOspfv3Area.Type.Normal.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeNormal"] = oVrfOspfv3Area.Type.Normal.Misc
									}
									if oVrfOspfv3Area.Type.Normal.Abr != nil {
										nestedVrfOspfv3Area.Type.Normal.Abr = &VrfOspfv3AreaTypeNormalAbr{}
										if oVrfOspfv3Area.Type.Normal.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNormalAbr"] = oVrfOspfv3Area.Type.Normal.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Normal.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.ImportList = oVrfOspfv3Area.Type.Normal.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.ExportList = oVrfOspfv3Area.Type.Normal.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.InboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Normal.Abr.OutboundFilterList
										}
									}
								}
								if oVrfOspfv3Area.Type.Stub != nil {
									nestedVrfOspfv3Area.Type.Stub = &VrfOspfv3AreaTypeStub{}
									if oVrfOspfv3Area.Type.Stub.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeStub"] = oVrfOspfv3Area.Type.Stub.Misc
									}
									if oVrfOspfv3Area.Type.Stub.NoSummary != nil {
										nestedVrfOspfv3Area.Type.Stub.NoSummary = util.AsBool(oVrfOspfv3Area.Type.Stub.NoSummary, nil)
									}
									if oVrfOspfv3Area.Type.Stub.Abr != nil {
										nestedVrfOspfv3Area.Type.Stub.Abr = &VrfOspfv3AreaTypeStubAbr{}
										if oVrfOspfv3Area.Type.Stub.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeStubAbr"] = oVrfOspfv3Area.Type.Stub.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Stub.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.ImportList = oVrfOspfv3Area.Type.Stub.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.ExportList = oVrfOspfv3Area.Type.Stub.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.InboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Stub.Abr.OutboundFilterList
										}
									}
									if oVrfOspfv3Area.Type.Stub.DefaultRouteMetric != nil {
										nestedVrfOspfv3Area.Type.Stub.DefaultRouteMetric = oVrfOspfv3Area.Type.Stub.DefaultRouteMetric
									}
								}
								if oVrfOspfv3Area.Type.Nssa != nil {
									nestedVrfOspfv3Area.Type.Nssa = &VrfOspfv3AreaTypeNssa{}
									if oVrfOspfv3Area.Type.Nssa.Misc != nil {
										entry.Misc["VrfOspfv3AreaTypeNssa"] = oVrfOspfv3Area.Type.Nssa.Misc
									}
									if oVrfOspfv3Area.Type.Nssa.NoSummary != nil {
										nestedVrfOspfv3Area.Type.Nssa.NoSummary = util.AsBool(oVrfOspfv3Area.Type.Nssa.NoSummary, nil)
									}
									if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate != nil {
										nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate = &VrfOspfv3AreaTypeNssaDefaultInformationOriginate{}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNssaDefaultInformationOriginate"] = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Misc
										}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric != nil {
											nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.Metric
										}
										if oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType != nil {
											nestedVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType = oVrfOspfv3Area.Type.Nssa.DefaultInformationOriginate.MetricType
										}
									}
									if oVrfOspfv3Area.Type.Nssa.Abr != nil {
										nestedVrfOspfv3Area.Type.Nssa.Abr = &VrfOspfv3AreaTypeNssaAbr{}
										if oVrfOspfv3Area.Type.Nssa.Abr.Misc != nil {
											entry.Misc["VrfOspfv3AreaTypeNssaAbr"] = oVrfOspfv3Area.Type.Nssa.Abr.Misc
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.ImportList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.ImportList = oVrfOspfv3Area.Type.Nssa.Abr.ImportList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.ExportList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.ExportList = oVrfOspfv3Area.Type.Nssa.Abr.ExportList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.InboundFilterList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList = oVrfOspfv3Area.Type.Nssa.Abr.OutboundFilterList
										}
										if oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange != nil {
											nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = []VrfOspfv3AreaTypeNssaAbrNssaExtRange{}
											for _, oVrfOspfv3AreaTypeNssaAbrNssaExtRange := range oVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange {
												nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange := VrfOspfv3AreaTypeNssaAbrNssaExtRange{}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc != nil {
													entry.Misc["VrfOspfv3AreaTypeNssaAbrNssaExtRange"] = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Misc
												}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name != "" {
													nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name = oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Name
												}
												if oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise != nil {
													nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise = util.AsBool(oVrfOspfv3AreaTypeNssaAbrNssaExtRange.Advertise, nil)
												}
												nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange = append(nestedVrfOspfv3Area.Type.Nssa.Abr.NssaExtRange, nestedVrfOspfv3AreaTypeNssaAbrNssaExtRange)
											}
										}
									}
								}
							}
							if oVrfOspfv3Area.Range != nil {
								nestedVrfOspfv3Area.Range = []VrfOspfv3AreaRange{}
								for _, oVrfOspfv3AreaRange := range oVrfOspfv3Area.Range {
									nestedVrfOspfv3AreaRange := VrfOspfv3AreaRange{}
									if oVrfOspfv3AreaRange.Misc != nil {
										entry.Misc["VrfOspfv3AreaRange"] = oVrfOspfv3AreaRange.Misc
									}
									if oVrfOspfv3AreaRange.Name != "" {
										nestedVrfOspfv3AreaRange.Name = oVrfOspfv3AreaRange.Name
									}
									if oVrfOspfv3AreaRange.Advertise != nil {
										nestedVrfOspfv3AreaRange.Advertise = util.AsBool(oVrfOspfv3AreaRange.Advertise, nil)
									}
									nestedVrfOspfv3Area.Range = append(nestedVrfOspfv3Area.Range, nestedVrfOspfv3AreaRange)
								}
							}
							if oVrfOspfv3Area.Interface != nil {
								nestedVrfOspfv3Area.Interface = []VrfOspfv3AreaInterface{}
								for _, oVrfOspfv3AreaInterface := range oVrfOspfv3Area.Interface {
									nestedVrfOspfv3AreaInterface := VrfOspfv3AreaInterface{}
									if oVrfOspfv3AreaInterface.Misc != nil {
										entry.Misc["VrfOspfv3AreaInterface"] = oVrfOspfv3AreaInterface.Misc
									}
									if oVrfOspfv3AreaInterface.Name != "" {
										nestedVrfOspfv3AreaInterface.Name = oVrfOspfv3AreaInterface.Name
									}
									if oVrfOspfv3AreaInterface.Enable != nil {
										nestedVrfOspfv3AreaInterface.Enable = util.AsBool(oVrfOspfv3AreaInterface.Enable, nil)
									}
									if oVrfOspfv3AreaInterface.MtuIgnore != nil {
										nestedVrfOspfv3AreaInterface.MtuIgnore = util.AsBool(oVrfOspfv3AreaInterface.MtuIgnore, nil)
									}
									if oVrfOspfv3AreaInterface.Passive != nil {
										nestedVrfOspfv3AreaInterface.Passive = util.AsBool(oVrfOspfv3AreaInterface.Passive, nil)
									}
									if oVrfOspfv3AreaInterface.Priority != nil {
										nestedVrfOspfv3AreaInterface.Priority = oVrfOspfv3AreaInterface.Priority
									}
									if oVrfOspfv3AreaInterface.Metric != nil {
										nestedVrfOspfv3AreaInterface.Metric = oVrfOspfv3AreaInterface.Metric
									}
									if oVrfOspfv3AreaInterface.InstanceId != nil {
										nestedVrfOspfv3AreaInterface.InstanceId = oVrfOspfv3AreaInterface.InstanceId
									}
									if oVrfOspfv3AreaInterface.Authentication != nil {
										nestedVrfOspfv3AreaInterface.Authentication = oVrfOspfv3AreaInterface.Authentication
									}
									if oVrfOspfv3AreaInterface.Timing != nil {
										nestedVrfOspfv3AreaInterface.Timing = oVrfOspfv3AreaInterface.Timing
									}
									if oVrfOspfv3AreaInterface.LinkType != nil {
										nestedVrfOspfv3AreaInterface.LinkType = &VrfOspfv3AreaInterfaceLinkType{}
										if oVrfOspfv3AreaInterface.LinkType.Misc != nil {
											entry.Misc["VrfOspfv3AreaInterfaceLinkType"] = oVrfOspfv3AreaInterface.LinkType.Misc
										}
										if oVrfOspfv3AreaInterface.LinkType.Broadcast != nil {
											nestedVrfOspfv3AreaInterface.LinkType.Broadcast = &VrfOspfv3AreaInterfaceLinkTypeBroadcast{}
											if oVrfOspfv3AreaInterface.LinkType.Broadcast.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeBroadcast"] = oVrfOspfv3AreaInterface.LinkType.Broadcast.Misc
											}
										}
										if oVrfOspfv3AreaInterface.LinkType.P2p != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2p = &VrfOspfv3AreaInterfaceLinkTypeP2p{}
											if oVrfOspfv3AreaInterface.LinkType.P2p.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2p"] = oVrfOspfv3AreaInterface.LinkType.P2p.Misc
											}
										}
										if oVrfOspfv3AreaInterface.LinkType.P2mp != nil {
											nestedVrfOspfv3AreaInterface.LinkType.P2mp = &VrfOspfv3AreaInterfaceLinkTypeP2mp{}
											if oVrfOspfv3AreaInterface.LinkType.P2mp.Misc != nil {
												entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mp"] = oVrfOspfv3AreaInterface.LinkType.P2mp.Misc
											}
											if oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor != nil {
												nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{}
												for _, oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := range oVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor {
													nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor := VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc != nil {
														entry.Misc["VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor"] = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Misc
													}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name != "" {
														nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Name
													}
													if oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority != nil {
														nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority = oVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor.Priority
													}
													nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor = append(nestedVrfOspfv3AreaInterface.LinkType.P2mp.Neighbor, nestedVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor)
												}
											}
										}
									}
									if oVrfOspfv3AreaInterface.Bfd != nil {
										nestedVrfOspfv3AreaInterface.Bfd = &VrfOspfv3AreaInterfaceBfd{}
										if oVrfOspfv3AreaInterface.Bfd.Misc != nil {
											entry.Misc["VrfOspfv3AreaInterfaceBfd"] = oVrfOspfv3AreaInterface.Bfd.Misc
										}
										if oVrfOspfv3AreaInterface.Bfd.Profile != nil {
											nestedVrfOspfv3AreaInterface.Bfd.Profile = oVrfOspfv3AreaInterface.Bfd.Profile
										}
									}
									nestedVrfOspfv3Area.Interface = append(nestedVrfOspfv3Area.Interface, nestedVrfOspfv3AreaInterface)
								}
							}
							if oVrfOspfv3Area.VirtualLink != nil {
								nestedVrfOspfv3Area.VirtualLink = []VrfOspfv3AreaVirtualLink{}
								for _, oVrfOspfv3AreaVirtualLink := range oVrfOspfv3Area.VirtualLink {
									nestedVrfOspfv3AreaVirtualLink := VrfOspfv3AreaVirtualLink{}
									if oVrfOspfv3AreaVirtualLink.Misc != nil {
										entry.Misc["VrfOspfv3AreaVirtualLink"] = oVrfOspfv3AreaVirtualLink.Misc
									}
									if oVrfOspfv3AreaVirtualLink.Name != "" {
										nestedVrfOspfv3AreaVirtualLink.Name = oVrfOspfv3AreaVirtualLink.Name
									}
									if oVrfOspfv3AreaVirtualLink.NeighborId != nil {
										nestedVrfOspfv3AreaVirtualLink.NeighborId = oVrfOspfv3AreaVirtualLink.NeighborId
									}
									if oVrfOspfv3AreaVirtualLink.TransitAreaId != nil {
										nestedVrfOspfv3AreaVirtualLink.TransitAreaId = oVrfOspfv3AreaVirtualLink.TransitAreaId
									}
									if oVrfOspfv3AreaVirtualLink.Enable != nil {
										nestedVrfOspfv3AreaVirtualLink.Enable = util.AsBool(oVrfOspfv3AreaVirtualLink.Enable, nil)
									}
									if oVrfOspfv3AreaVirtualLink.InstanceId != nil {
										nestedVrfOspfv3AreaVirtualLink.InstanceId = oVrfOspfv3AreaVirtualLink.InstanceId
									}
									if oVrfOspfv3AreaVirtualLink.Timing != nil {
										nestedVrfOspfv3AreaVirtualLink.Timing = oVrfOspfv3AreaVirtualLink.Timing
									}
									if oVrfOspfv3AreaVirtualLink.Authentication != nil {
										nestedVrfOspfv3AreaVirtualLink.Authentication = oVrfOspfv3AreaVirtualLink.Authentication
									}
									nestedVrfOspfv3Area.VirtualLink = append(nestedVrfOspfv3Area.VirtualLink, nestedVrfOspfv3AreaVirtualLink)
								}
							}
							nestedVrf.Ospfv3.Area = append(nestedVrf.Ospfv3.Area, nestedVrfOspfv3Area)
						}
					}
				}
				if oVrf.Ecmp != nil {
					nestedVrf.Ecmp = &VrfEcmp{}
					if oVrf.Ecmp.Misc != nil {
						entry.Misc["VrfEcmp"] = oVrf.Ecmp.Misc
					}
					if oVrf.Ecmp.Enable != nil {
						nestedVrf.Ecmp.Enable = util.AsBool(oVrf.Ecmp.Enable, nil)
					}
					if oVrf.Ecmp.MaxPath != nil {
						nestedVrf.Ecmp.MaxPath = oVrf.Ecmp.MaxPath
					}
					if oVrf.Ecmp.SymmetricReturn != nil {
						nestedVrf.Ecmp.SymmetricReturn = util.AsBool(oVrf.Ecmp.SymmetricReturn, nil)
					}
					if oVrf.Ecmp.StrictSourcePath != nil {
						nestedVrf.Ecmp.StrictSourcePath = util.AsBool(oVrf.Ecmp.StrictSourcePath, nil)
					}
					if oVrf.Ecmp.Algorithm != nil {
						nestedVrf.Ecmp.Algorithm = &VrfEcmpAlgorithm{}
						if oVrf.Ecmp.Algorithm.Misc != nil {
							entry.Misc["VrfEcmpAlgorithm"] = oVrf.Ecmp.Algorithm.Misc
						}
						if oVrf.Ecmp.Algorithm.IpModulo != nil {
							nestedVrf.Ecmp.Algorithm.IpModulo = &VrfEcmpAlgorithmIpModulo{}
							if oVrf.Ecmp.Algorithm.IpModulo.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmIpModulo"] = oVrf.Ecmp.Algorithm.IpModulo.Misc
							}
						}
						if oVrf.Ecmp.Algorithm.IpHash != nil {
							nestedVrf.Ecmp.Algorithm.IpHash = &VrfEcmpAlgorithmIpHash{}
							if oVrf.Ecmp.Algorithm.IpHash.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmIpHash"] = oVrf.Ecmp.Algorithm.IpHash.Misc
							}
							if oVrf.Ecmp.Algorithm.IpHash.SrcOnly != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.SrcOnly = util.AsBool(oVrf.Ecmp.Algorithm.IpHash.SrcOnly, nil)
							}
							if oVrf.Ecmp.Algorithm.IpHash.UsePort != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.UsePort = util.AsBool(oVrf.Ecmp.Algorithm.IpHash.UsePort, nil)
							}
							if oVrf.Ecmp.Algorithm.IpHash.HashSeed != nil {
								nestedVrf.Ecmp.Algorithm.IpHash.HashSeed = oVrf.Ecmp.Algorithm.IpHash.HashSeed
							}
						}
						if oVrf.Ecmp.Algorithm.WeightedRoundRobin != nil {
							nestedVrf.Ecmp.Algorithm.WeightedRoundRobin = &VrfEcmpAlgorithmWeightedRoundRobin{}
							if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmWeightedRoundRobin"] = oVrf.Ecmp.Algorithm.WeightedRoundRobin.Misc
							}
							if oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
								nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = []VrfEcmpAlgorithmWeightedRoundRobinInterface{}
								for _, oVrfEcmpAlgorithmWeightedRoundRobinInterface := range oVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface {
									nestedVrfEcmpAlgorithmWeightedRoundRobinInterface := VrfEcmpAlgorithmWeightedRoundRobinInterface{}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc != nil {
										entry.Misc["VrfEcmpAlgorithmWeightedRoundRobinInterface"] = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Misc
									}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
										nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Name = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Name
									}
									if oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
										nestedVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight = oVrfEcmpAlgorithmWeightedRoundRobinInterface.Weight
									}
									nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface = append(nestedVrf.Ecmp.Algorithm.WeightedRoundRobin.Interface, nestedVrfEcmpAlgorithmWeightedRoundRobinInterface)
								}
							}
						}
						if oVrf.Ecmp.Algorithm.BalancedRoundRobin != nil {
							nestedVrf.Ecmp.Algorithm.BalancedRoundRobin = &VrfEcmpAlgorithmBalancedRoundRobin{}
							if oVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc != nil {
								entry.Misc["VrfEcmpAlgorithmBalancedRoundRobin"] = oVrf.Ecmp.Algorithm.BalancedRoundRobin.Misc
							}
						}
					}
				}
				if oVrf.Multicast != nil {
					nestedVrf.Multicast = &VrfMulticast{}
					if oVrf.Multicast.Misc != nil {
						entry.Misc["VrfMulticast"] = oVrf.Multicast.Misc
					}
					if oVrf.Multicast.Enable != nil {
						nestedVrf.Multicast.Enable = util.AsBool(oVrf.Multicast.Enable, nil)
					}
					if oVrf.Multicast.StaticRoute != nil {
						nestedVrf.Multicast.StaticRoute = []VrfMulticastStaticRoute{}
						for _, oVrfMulticastStaticRoute := range oVrf.Multicast.StaticRoute {
							nestedVrfMulticastStaticRoute := VrfMulticastStaticRoute{}
							if oVrfMulticastStaticRoute.Misc != nil {
								entry.Misc["VrfMulticastStaticRoute"] = oVrfMulticastStaticRoute.Misc
							}
							if oVrfMulticastStaticRoute.Name != "" {
								nestedVrfMulticastStaticRoute.Name = oVrfMulticastStaticRoute.Name
							}
							if oVrfMulticastStaticRoute.Destination != nil {
								nestedVrfMulticastStaticRoute.Destination = oVrfMulticastStaticRoute.Destination
							}
							if oVrfMulticastStaticRoute.Interface != nil {
								nestedVrfMulticastStaticRoute.Interface = oVrfMulticastStaticRoute.Interface
							}
							if oVrfMulticastStaticRoute.Preference != nil {
								nestedVrfMulticastStaticRoute.Preference = oVrfMulticastStaticRoute.Preference
							}
							if oVrfMulticastStaticRoute.Nexthop != nil {
								nestedVrfMulticastStaticRoute.Nexthop = &VrfMulticastStaticRouteNexthop{}
								if oVrfMulticastStaticRoute.Nexthop.Misc != nil {
									entry.Misc["VrfMulticastStaticRouteNexthop"] = oVrfMulticastStaticRoute.Nexthop.Misc
								}
								if oVrfMulticastStaticRoute.Nexthop.IpAddress != nil {
									nestedVrfMulticastStaticRoute.Nexthop.IpAddress = oVrfMulticastStaticRoute.Nexthop.IpAddress
								}
							}
							nestedVrf.Multicast.StaticRoute = append(nestedVrf.Multicast.StaticRoute, nestedVrfMulticastStaticRoute)
						}
					}
					if oVrf.Multicast.Pim != nil {
						nestedVrf.Multicast.Pim = &VrfMulticastPim{}
						if oVrf.Multicast.Pim.Misc != nil {
							entry.Misc["VrfMulticastPim"] = oVrf.Multicast.Pim.Misc
						}
						if oVrf.Multicast.Pim.Enable != nil {
							nestedVrf.Multicast.Pim.Enable = util.AsBool(oVrf.Multicast.Pim.Enable, nil)
						}
						if oVrf.Multicast.Pim.RpfLookupMode != nil {
							nestedVrf.Multicast.Pim.RpfLookupMode = oVrf.Multicast.Pim.RpfLookupMode
						}
						if oVrf.Multicast.Pim.RouteAgeoutTime != nil {
							nestedVrf.Multicast.Pim.RouteAgeoutTime = oVrf.Multicast.Pim.RouteAgeoutTime
						}
						if oVrf.Multicast.Pim.IfTimerGlobal != nil {
							nestedVrf.Multicast.Pim.IfTimerGlobal = oVrf.Multicast.Pim.IfTimerGlobal
						}
						if oVrf.Multicast.Pim.GroupPermission != nil {
							nestedVrf.Multicast.Pim.GroupPermission = oVrf.Multicast.Pim.GroupPermission
						}
						if oVrf.Multicast.Pim.SsmAddressSpace != nil {
							nestedVrf.Multicast.Pim.SsmAddressSpace = &VrfMulticastPimSsmAddressSpace{}
							if oVrf.Multicast.Pim.SsmAddressSpace.Misc != nil {
								entry.Misc["VrfMulticastPimSsmAddressSpace"] = oVrf.Multicast.Pim.SsmAddressSpace.Misc
							}
							if oVrf.Multicast.Pim.SsmAddressSpace.GroupList != nil {
								nestedVrf.Multicast.Pim.SsmAddressSpace.GroupList = oVrf.Multicast.Pim.SsmAddressSpace.GroupList
							}
						}
						if oVrf.Multicast.Pim.Rp != nil {
							nestedVrf.Multicast.Pim.Rp = &VrfMulticastPimRp{}
							if oVrf.Multicast.Pim.Rp.Misc != nil {
								entry.Misc["VrfMulticastPimRp"] = oVrf.Multicast.Pim.Rp.Misc
							}
							if oVrf.Multicast.Pim.Rp.LocalRp != nil {
								nestedVrf.Multicast.Pim.Rp.LocalRp = &VrfMulticastPimRpLocalRp{}
								if oVrf.Multicast.Pim.Rp.LocalRp.Misc != nil {
									entry.Misc["VrfMulticastPimRpLocalRp"] = oVrf.Multicast.Pim.Rp.LocalRp.Misc
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp = &VrfMulticastPimRpLocalRpStaticRp{}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpLocalRpStaticRp"] = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Misc
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Interface
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Address
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override = util.AsBool(oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.Override, nil)
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.StaticRp.GroupList
									}
								}
								if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp != nil {
									nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp = &VrfMulticastPimRpLocalRpCandidateRp{}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpLocalRpCandidateRp"] = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Misc
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Interface
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Address
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.Priority
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.AdvertisementInterval
									}
									if oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList != nil {
										nestedVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList = oVrf.Multicast.Pim.Rp.LocalRp.CandidateRp.GroupList
									}
								}
							}
							if oVrf.Multicast.Pim.Rp.ExternalRp != nil {
								nestedVrf.Multicast.Pim.Rp.ExternalRp = []VrfMulticastPimRpExternalRp{}
								for _, oVrfMulticastPimRpExternalRp := range oVrf.Multicast.Pim.Rp.ExternalRp {
									nestedVrfMulticastPimRpExternalRp := VrfMulticastPimRpExternalRp{}
									if oVrfMulticastPimRpExternalRp.Misc != nil {
										entry.Misc["VrfMulticastPimRpExternalRp"] = oVrfMulticastPimRpExternalRp.Misc
									}
									if oVrfMulticastPimRpExternalRp.Name != "" {
										nestedVrfMulticastPimRpExternalRp.Name = oVrfMulticastPimRpExternalRp.Name
									}
									if oVrfMulticastPimRpExternalRp.GroupList != nil {
										nestedVrfMulticastPimRpExternalRp.GroupList = oVrfMulticastPimRpExternalRp.GroupList
									}
									if oVrfMulticastPimRpExternalRp.Override != nil {
										nestedVrfMulticastPimRpExternalRp.Override = util.AsBool(oVrfMulticastPimRpExternalRp.Override, nil)
									}
									nestedVrf.Multicast.Pim.Rp.ExternalRp = append(nestedVrf.Multicast.Pim.Rp.ExternalRp, nestedVrfMulticastPimRpExternalRp)
								}
							}
						}
						if oVrf.Multicast.Pim.SptThreshold != nil {
							nestedVrf.Multicast.Pim.SptThreshold = []VrfMulticastPimSptThreshold{}
							for _, oVrfMulticastPimSptThreshold := range oVrf.Multicast.Pim.SptThreshold {
								nestedVrfMulticastPimSptThreshold := VrfMulticastPimSptThreshold{}
								if oVrfMulticastPimSptThreshold.Misc != nil {
									entry.Misc["VrfMulticastPimSptThreshold"] = oVrfMulticastPimSptThreshold.Misc
								}
								if oVrfMulticastPimSptThreshold.Name != "" {
									nestedVrfMulticastPimSptThreshold.Name = oVrfMulticastPimSptThreshold.Name
								}
								if oVrfMulticastPimSptThreshold.Threshold != nil {
									nestedVrfMulticastPimSptThreshold.Threshold = oVrfMulticastPimSptThreshold.Threshold
								}
								nestedVrf.Multicast.Pim.SptThreshold = append(nestedVrf.Multicast.Pim.SptThreshold, nestedVrfMulticastPimSptThreshold)
							}
						}
						if oVrf.Multicast.Pim.Interface != nil {
							nestedVrf.Multicast.Pim.Interface = []VrfMulticastPimInterface{}
							for _, oVrfMulticastPimInterface := range oVrf.Multicast.Pim.Interface {
								nestedVrfMulticastPimInterface := VrfMulticastPimInterface{}
								if oVrfMulticastPimInterface.Misc != nil {
									entry.Misc["VrfMulticastPimInterface"] = oVrfMulticastPimInterface.Misc
								}
								if oVrfMulticastPimInterface.Name != "" {
									nestedVrfMulticastPimInterface.Name = oVrfMulticastPimInterface.Name
								}
								if oVrfMulticastPimInterface.Description != nil {
									nestedVrfMulticastPimInterface.Description = oVrfMulticastPimInterface.Description
								}
								if oVrfMulticastPimInterface.DrPriority != nil {
									nestedVrfMulticastPimInterface.DrPriority = oVrfMulticastPimInterface.DrPriority
								}
								if oVrfMulticastPimInterface.SendBsm != nil {
									nestedVrfMulticastPimInterface.SendBsm = util.AsBool(oVrfMulticastPimInterface.SendBsm, nil)
								}
								if oVrfMulticastPimInterface.IfTimer != nil {
									nestedVrfMulticastPimInterface.IfTimer = oVrfMulticastPimInterface.IfTimer
								}
								if oVrfMulticastPimInterface.NeighborFilter != nil {
									nestedVrfMulticastPimInterface.NeighborFilter = oVrfMulticastPimInterface.NeighborFilter
								}
								nestedVrf.Multicast.Pim.Interface = append(nestedVrf.Multicast.Pim.Interface, nestedVrfMulticastPimInterface)
							}
						}
					}
					if oVrf.Multicast.Igmp != nil {
						nestedVrf.Multicast.Igmp = &VrfMulticastIgmp{}
						if oVrf.Multicast.Igmp.Misc != nil {
							entry.Misc["VrfMulticastIgmp"] = oVrf.Multicast.Igmp.Misc
						}
						if oVrf.Multicast.Igmp.Enable != nil {
							nestedVrf.Multicast.Igmp.Enable = util.AsBool(oVrf.Multicast.Igmp.Enable, nil)
						}
						if oVrf.Multicast.Igmp.Dynamic != nil {
							nestedVrf.Multicast.Igmp.Dynamic = &VrfMulticastIgmpDynamic{}
							if oVrf.Multicast.Igmp.Dynamic.Misc != nil {
								entry.Misc["VrfMulticastIgmpDynamic"] = oVrf.Multicast.Igmp.Dynamic.Misc
							}
							if oVrf.Multicast.Igmp.Dynamic.Interface != nil {
								nestedVrf.Multicast.Igmp.Dynamic.Interface = []VrfMulticastIgmpDynamicInterface{}
								for _, oVrfMulticastIgmpDynamicInterface := range oVrf.Multicast.Igmp.Dynamic.Interface {
									nestedVrfMulticastIgmpDynamicInterface := VrfMulticastIgmpDynamicInterface{}
									if oVrfMulticastIgmpDynamicInterface.Misc != nil {
										entry.Misc["VrfMulticastIgmpDynamicInterface"] = oVrfMulticastIgmpDynamicInterface.Misc
									}
									if oVrfMulticastIgmpDynamicInterface.Name != "" {
										nestedVrfMulticastIgmpDynamicInterface.Name = oVrfMulticastIgmpDynamicInterface.Name
									}
									if oVrfMulticastIgmpDynamicInterface.Version != nil {
										nestedVrfMulticastIgmpDynamicInterface.Version = oVrfMulticastIgmpDynamicInterface.Version
									}
									if oVrfMulticastIgmpDynamicInterface.Robustness != nil {
										nestedVrfMulticastIgmpDynamicInterface.Robustness = oVrfMulticastIgmpDynamicInterface.Robustness
									}
									if oVrfMulticastIgmpDynamicInterface.GroupFilter != nil {
										nestedVrfMulticastIgmpDynamicInterface.GroupFilter = oVrfMulticastIgmpDynamicInterface.GroupFilter
									}
									if oVrfMulticastIgmpDynamicInterface.MaxGroups != nil {
										nestedVrfMulticastIgmpDynamicInterface.MaxGroups = oVrfMulticastIgmpDynamicInterface.MaxGroups
									}
									if oVrfMulticastIgmpDynamicInterface.MaxSources != nil {
										nestedVrfMulticastIgmpDynamicInterface.MaxSources = oVrfMulticastIgmpDynamicInterface.MaxSources
									}
									if oVrfMulticastIgmpDynamicInterface.QueryProfile != nil {
										nestedVrfMulticastIgmpDynamicInterface.QueryProfile = oVrfMulticastIgmpDynamicInterface.QueryProfile
									}
									if oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing != nil {
										nestedVrfMulticastIgmpDynamicInterface.RouterAlertPolicing = util.AsBool(oVrfMulticastIgmpDynamicInterface.RouterAlertPolicing, nil)
									}
									nestedVrf.Multicast.Igmp.Dynamic.Interface = append(nestedVrf.Multicast.Igmp.Dynamic.Interface, nestedVrfMulticastIgmpDynamicInterface)
								}
							}
						}
						if oVrf.Multicast.Igmp.Static != nil {
							nestedVrf.Multicast.Igmp.Static = []VrfMulticastIgmpStatic{}
							for _, oVrfMulticastIgmpStatic := range oVrf.Multicast.Igmp.Static {
								nestedVrfMulticastIgmpStatic := VrfMulticastIgmpStatic{}
								if oVrfMulticastIgmpStatic.Misc != nil {
									entry.Misc["VrfMulticastIgmpStatic"] = oVrfMulticastIgmpStatic.Misc
								}
								if oVrfMulticastIgmpStatic.Name != "" {
									nestedVrfMulticastIgmpStatic.Name = oVrfMulticastIgmpStatic.Name
								}
								if oVrfMulticastIgmpStatic.Interface != nil {
									nestedVrfMulticastIgmpStatic.Interface = oVrfMulticastIgmpStatic.Interface
								}
								if oVrfMulticastIgmpStatic.GroupAddress != nil {
									nestedVrfMulticastIgmpStatic.GroupAddress = oVrfMulticastIgmpStatic.GroupAddress
								}
								if oVrfMulticastIgmpStatic.SourceAddress != nil {
									nestedVrfMulticastIgmpStatic.SourceAddress = oVrfMulticastIgmpStatic.SourceAddress
								}
								nestedVrf.Multicast.Igmp.Static = append(nestedVrf.Multicast.Igmp.Static, nestedVrfMulticastIgmpStatic)
							}
						}
					}
					if oVrf.Multicast.Msdp != nil {
						nestedVrf.Multicast.Msdp = &VrfMulticastMsdp{}
						if oVrf.Multicast.Msdp.Misc != nil {
							entry.Misc["VrfMulticastMsdp"] = oVrf.Multicast.Msdp.Misc
						}
						if oVrf.Multicast.Msdp.Enable != nil {
							nestedVrf.Multicast.Msdp.Enable = util.AsBool(oVrf.Multicast.Msdp.Enable, nil)
						}
						if oVrf.Multicast.Msdp.GlobalTimer != nil {
							nestedVrf.Multicast.Msdp.GlobalTimer = oVrf.Multicast.Msdp.GlobalTimer
						}
						if oVrf.Multicast.Msdp.GlobalAuthentication != nil {
							nestedVrf.Multicast.Msdp.GlobalAuthentication = oVrf.Multicast.Msdp.GlobalAuthentication
						}
						if oVrf.Multicast.Msdp.OriginatorId != nil {
							nestedVrf.Multicast.Msdp.OriginatorId = &VrfMulticastMsdpOriginatorId{}
							if oVrf.Multicast.Msdp.OriginatorId.Misc != nil {
								entry.Misc["VrfMulticastMsdpOriginatorId"] = oVrf.Multicast.Msdp.OriginatorId.Misc
							}
							if oVrf.Multicast.Msdp.OriginatorId.Interface != nil {
								nestedVrf.Multicast.Msdp.OriginatorId.Interface = oVrf.Multicast.Msdp.OriginatorId.Interface
							}
							if oVrf.Multicast.Msdp.OriginatorId.Ip != nil {
								nestedVrf.Multicast.Msdp.OriginatorId.Ip = oVrf.Multicast.Msdp.OriginatorId.Ip
							}
						}
						if oVrf.Multicast.Msdp.Peer != nil {
							nestedVrf.Multicast.Msdp.Peer = []VrfMulticastMsdpPeer{}
							for _, oVrfMulticastMsdpPeer := range oVrf.Multicast.Msdp.Peer {
								nestedVrfMulticastMsdpPeer := VrfMulticastMsdpPeer{}
								if oVrfMulticastMsdpPeer.Misc != nil {
									entry.Misc["VrfMulticastMsdpPeer"] = oVrfMulticastMsdpPeer.Misc
								}
								if oVrfMulticastMsdpPeer.Name != "" {
									nestedVrfMulticastMsdpPeer.Name = oVrfMulticastMsdpPeer.Name
								}
								if oVrfMulticastMsdpPeer.Enable != nil {
									nestedVrfMulticastMsdpPeer.Enable = util.AsBool(oVrfMulticastMsdpPeer.Enable, nil)
								}
								if oVrfMulticastMsdpPeer.PeerAs != nil {
									nestedVrfMulticastMsdpPeer.PeerAs = oVrfMulticastMsdpPeer.PeerAs
								}
								if oVrfMulticastMsdpPeer.Authentication != nil {
									nestedVrfMulticastMsdpPeer.Authentication = oVrfMulticastMsdpPeer.Authentication
								}
								if oVrfMulticastMsdpPeer.MaxSa != nil {
									nestedVrfMulticastMsdpPeer.MaxSa = oVrfMulticastMsdpPeer.MaxSa
								}
								if oVrfMulticastMsdpPeer.InboundSaFilter != nil {
									nestedVrfMulticastMsdpPeer.InboundSaFilter = oVrfMulticastMsdpPeer.InboundSaFilter
								}
								if oVrfMulticastMsdpPeer.OutboundSaFilter != nil {
									nestedVrfMulticastMsdpPeer.OutboundSaFilter = oVrfMulticastMsdpPeer.OutboundSaFilter
								}
								if oVrfMulticastMsdpPeer.LocalAddress != nil {
									nestedVrfMulticastMsdpPeer.LocalAddress = &VrfMulticastMsdpPeerLocalAddress{}
									if oVrfMulticastMsdpPeer.LocalAddress.Misc != nil {
										entry.Misc["VrfMulticastMsdpPeerLocalAddress"] = oVrfMulticastMsdpPeer.LocalAddress.Misc
									}
									if oVrfMulticastMsdpPeer.LocalAddress.Interface != nil {
										nestedVrfMulticastMsdpPeer.LocalAddress.Interface = oVrfMulticastMsdpPeer.LocalAddress.Interface
									}
									if oVrfMulticastMsdpPeer.LocalAddress.Ip != nil {
										nestedVrfMulticastMsdpPeer.LocalAddress.Ip = oVrfMulticastMsdpPeer.LocalAddress.Ip
									}
								}
								if oVrfMulticastMsdpPeer.PeerAddress != nil {
									nestedVrfMulticastMsdpPeer.PeerAddress = &VrfMulticastMsdpPeerPeerAddress{}
									if oVrfMulticastMsdpPeer.PeerAddress.Misc != nil {
										entry.Misc["VrfMulticastMsdpPeerPeerAddress"] = oVrfMulticastMsdpPeer.PeerAddress.Misc
									}
									if oVrfMulticastMsdpPeer.PeerAddress.Ip != nil {
										nestedVrfMulticastMsdpPeer.PeerAddress.Ip = oVrfMulticastMsdpPeer.PeerAddress.Ip
									}
									if oVrfMulticastMsdpPeer.PeerAddress.Fqdn != nil {
										nestedVrfMulticastMsdpPeer.PeerAddress.Fqdn = oVrfMulticastMsdpPeer.PeerAddress.Fqdn
									}
								}
								nestedVrf.Multicast.Msdp.Peer = append(nestedVrf.Multicast.Msdp.Peer, nestedVrfMulticastMsdpPeer)
							}
						}
					}
				}
				if oVrf.Rip != nil {
					nestedVrf.Rip = &VrfRip{}
					if oVrf.Rip.Misc != nil {
						entry.Misc["VrfRip"] = oVrf.Rip.Misc
					}
					if oVrf.Rip.Enable != nil {
						nestedVrf.Rip.Enable = util.AsBool(oVrf.Rip.Enable, nil)
					}
					if oVrf.Rip.DefaultInformationOriginate != nil {
						nestedVrf.Rip.DefaultInformationOriginate = util.AsBool(oVrf.Rip.DefaultInformationOriginate, nil)
					}
					if oVrf.Rip.GlobalTimer != nil {
						nestedVrf.Rip.GlobalTimer = oVrf.Rip.GlobalTimer
					}
					if oVrf.Rip.AuthProfile != nil {
						nestedVrf.Rip.AuthProfile = oVrf.Rip.AuthProfile
					}
					if oVrf.Rip.RedistributionProfile != nil {
						nestedVrf.Rip.RedistributionProfile = oVrf.Rip.RedistributionProfile
					}
					if oVrf.Rip.GlobalBfd != nil {
						nestedVrf.Rip.GlobalBfd = &VrfRipGlobalBfd{}
						if oVrf.Rip.GlobalBfd.Misc != nil {
							entry.Misc["VrfRipGlobalBfd"] = oVrf.Rip.GlobalBfd.Misc
						}
						if oVrf.Rip.GlobalBfd.Profile != nil {
							nestedVrf.Rip.GlobalBfd.Profile = oVrf.Rip.GlobalBfd.Profile
						}
					}
					if oVrf.Rip.GlobalInboundDistributeList != nil {
						nestedVrf.Rip.GlobalInboundDistributeList = &VrfRipGlobalInboundDistributeList{}
						if oVrf.Rip.GlobalInboundDistributeList.Misc != nil {
							entry.Misc["VrfRipGlobalInboundDistributeList"] = oVrf.Rip.GlobalInboundDistributeList.Misc
						}
						if oVrf.Rip.GlobalInboundDistributeList.AccessList != nil {
							nestedVrf.Rip.GlobalInboundDistributeList.AccessList = oVrf.Rip.GlobalInboundDistributeList.AccessList
						}
					}
					if oVrf.Rip.GlobalOutboundDistributeList != nil {
						nestedVrf.Rip.GlobalOutboundDistributeList = &VrfRipGlobalOutboundDistributeList{}
						if oVrf.Rip.GlobalOutboundDistributeList.Misc != nil {
							entry.Misc["VrfRipGlobalOutboundDistributeList"] = oVrf.Rip.GlobalOutboundDistributeList.Misc
						}
						if oVrf.Rip.GlobalOutboundDistributeList.AccessList != nil {
							nestedVrf.Rip.GlobalOutboundDistributeList.AccessList = oVrf.Rip.GlobalOutboundDistributeList.AccessList
						}
					}
					if oVrf.Rip.Interface != nil {
						nestedVrf.Rip.Interface = []VrfRipInterface{}
						for _, oVrfRipInterface := range oVrf.Rip.Interface {
							nestedVrfRipInterface := VrfRipInterface{}
							if oVrfRipInterface.Misc != nil {
								entry.Misc["VrfRipInterface"] = oVrfRipInterface.Misc
							}
							if oVrfRipInterface.Name != "" {
								nestedVrfRipInterface.Name = oVrfRipInterface.Name
							}
							if oVrfRipInterface.Enable != nil {
								nestedVrfRipInterface.Enable = util.AsBool(oVrfRipInterface.Enable, nil)
							}
							if oVrfRipInterface.Mode != nil {
								nestedVrfRipInterface.Mode = oVrfRipInterface.Mode
							}
							if oVrfRipInterface.SplitHorizon != nil {
								nestedVrfRipInterface.SplitHorizon = oVrfRipInterface.SplitHorizon
							}
							if oVrfRipInterface.Authentication != nil {
								nestedVrfRipInterface.Authentication = oVrfRipInterface.Authentication
							}
							if oVrfRipInterface.Bfd != nil {
								nestedVrfRipInterface.Bfd = &VrfRipInterfaceBfd{}
								if oVrfRipInterface.Bfd.Misc != nil {
									entry.Misc["VrfRipInterfaceBfd"] = oVrfRipInterface.Bfd.Misc
								}
								if oVrfRipInterface.Bfd.Profile != nil {
									nestedVrfRipInterface.Bfd.Profile = oVrfRipInterface.Bfd.Profile
								}
							}
							if oVrfRipInterface.InterfaceInboundDistributeList != nil {
								nestedVrfRipInterface.InterfaceInboundDistributeList = &VrfRipInterfaceInterfaceInboundDistributeList{}
								if oVrfRipInterface.InterfaceInboundDistributeList.Misc != nil {
									entry.Misc["VrfRipInterfaceInterfaceInboundDistributeList"] = oVrfRipInterface.InterfaceInboundDistributeList.Misc
								}
								if oVrfRipInterface.InterfaceInboundDistributeList.AccessList != nil {
									nestedVrfRipInterface.InterfaceInboundDistributeList.AccessList = oVrfRipInterface.InterfaceInboundDistributeList.AccessList
								}
								if oVrfRipInterface.InterfaceInboundDistributeList.Metric != nil {
									nestedVrfRipInterface.InterfaceInboundDistributeList.Metric = oVrfRipInterface.InterfaceInboundDistributeList.Metric
								}
							}
							if oVrfRipInterface.InterfaceOutboundDistributeList != nil {
								nestedVrfRipInterface.InterfaceOutboundDistributeList = &VrfRipInterfaceInterfaceOutboundDistributeList{}
								if oVrfRipInterface.InterfaceOutboundDistributeList.Misc != nil {
									entry.Misc["VrfRipInterfaceInterfaceOutboundDistributeList"] = oVrfRipInterface.InterfaceOutboundDistributeList.Misc
								}
								if oVrfRipInterface.InterfaceOutboundDistributeList.AccessList != nil {
									nestedVrfRipInterface.InterfaceOutboundDistributeList.AccessList = oVrfRipInterface.InterfaceOutboundDistributeList.AccessList
								}
								if oVrfRipInterface.InterfaceOutboundDistributeList.Metric != nil {
									nestedVrfRipInterface.InterfaceOutboundDistributeList.Metric = oVrfRipInterface.InterfaceOutboundDistributeList.Metric
								}
							}
							nestedVrf.Rip.Interface = append(nestedVrf.Rip.Interface, nestedVrfRipInterface)
						}
					}
				}
				nestedVrfCol = append(nestedVrfCol, nestedVrf)
			}
			entry.Vrf = nestedVrfCol
		}

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !matchVrf(a.Vrf, b.Vrf) {
		return false
	}

	return true
}

func matchVrfAdminDists(a *VrfAdminDists, b *VrfAdminDists) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Static, b.Static) {
		return false
	}
	if !util.Ints64Match(a.StaticIpv6, b.StaticIpv6) {
		return false
	}
	if !util.Ints64Match(a.OspfInter, b.OspfInter) {
		return false
	}
	if !util.Ints64Match(a.OspfIntra, b.OspfIntra) {
		return false
	}
	if !util.Ints64Match(a.OspfExt, b.OspfExt) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Inter, b.Ospfv3Inter) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Intra, b.Ospfv3Intra) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Ext, b.Ospfv3Ext) {
		return false
	}
	if !util.Ints64Match(a.BgpInternal, b.BgpInternal) {
		return false
	}
	if !util.Ints64Match(a.BgpExternal, b.BgpExternal) {
		return false
	}
	if !util.Ints64Match(a.BgpLocal, b.BgpLocal) {
		return false
	}
	if !util.Ints64Match(a.Rip, b.Rip) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv4Static(a *VrfRibFilterIpv4Static, b *VrfRibFilterIpv4Static) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv4Bgp(a *VrfRibFilterIpv4Bgp, b *VrfRibFilterIpv4Bgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv4Ospf(a *VrfRibFilterIpv4Ospf, b *VrfRibFilterIpv4Ospf) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv4Rip(a *VrfRibFilterIpv4Rip, b *VrfRibFilterIpv4Rip) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv4(a *VrfRibFilterIpv4, b *VrfRibFilterIpv4) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRibFilterIpv4Static(a.Static, b.Static) {
		return false
	}
	if !matchVrfRibFilterIpv4Bgp(a.Bgp, b.Bgp) {
		return false
	}
	if !matchVrfRibFilterIpv4Ospf(a.Ospf, b.Ospf) {
		return false
	}
	if !matchVrfRibFilterIpv4Rip(a.Rip, b.Rip) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv6Static(a *VrfRibFilterIpv6Static, b *VrfRibFilterIpv6Static) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv6Bgp(a *VrfRibFilterIpv6Bgp, b *VrfRibFilterIpv6Bgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv6Ospfv3(a *VrfRibFilterIpv6Ospfv3, b *VrfRibFilterIpv6Ospfv3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouteMap, b.RouteMap) {
		return false
	}
	return true
}
func matchVrfRibFilterIpv6(a *VrfRibFilterIpv6, b *VrfRibFilterIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRibFilterIpv6Static(a.Static, b.Static) {
		return false
	}
	if !matchVrfRibFilterIpv6Bgp(a.Bgp, b.Bgp) {
		return false
	}
	if !matchVrfRibFilterIpv6Ospfv3(a.Ospfv3, b.Ospfv3) {
		return false
	}
	return true
}
func matchVrfRibFilter(a *VrfRibFilter, b *VrfRibFilter) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRibFilterIpv4(a.Ipv4, b.Ipv4) {
		return false
	}
	if !matchVrfRibFilterIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpMed(a *VrfBgpMed, b *VrfBgpMed) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AlwaysCompareMed, b.AlwaysCompareMed) {
		return false
	}
	if !util.BoolsMatch(a.DeterministicMedComparison, b.DeterministicMedComparison) {
		return false
	}
	return true
}
func matchVrfBgpGracefulRestart(a *VrfBgpGracefulRestart, b *VrfBgpGracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.StaleRouteTime, b.StaleRouteTime) {
		return false
	}
	if !util.Ints64Match(a.MaxPeerRestartTime, b.MaxPeerRestartTime) {
		return false
	}
	if !util.Ints64Match(a.LocalRestartTime, b.LocalRestartTime) {
		return false
	}
	return true
}
func matchVrfBgpGlobalBfd(a *VrfBgpGlobalBfd, b *VrfBgpGlobalBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfBgpRedistributionProfileIpv4(a *VrfBgpRedistributionProfileIpv4, b *VrfBgpRedistributionProfileIpv4) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Unicast, b.Unicast) {
		return false
	}
	return true
}
func matchVrfBgpRedistributionProfileIpv6(a *VrfBgpRedistributionProfileIpv6, b *VrfBgpRedistributionProfileIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Unicast, b.Unicast) {
		return false
	}
	return true
}
func matchVrfBgpRedistributionProfile(a *VrfBgpRedistributionProfile, b *VrfBgpRedistributionProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpRedistributionProfileIpv4(a.Ipv4, b.Ipv4) {
		return false
	}
	if !matchVrfBgpRedistributionProfileIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpAdvertiseNetworkIpv4Network(a []VrfBgpAdvertiseNetworkIpv4Network, b []VrfBgpAdvertiseNetworkIpv4Network) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Unicast, b.Unicast) {
				return false
			}
			if !util.BoolsMatch(a.Multicast, b.Multicast) {
				return false
			}
			if !util.BoolsMatch(a.Backdoor, b.Backdoor) {
				return false
			}
		}
	}
	return true
}
func matchVrfBgpAdvertiseNetworkIpv4(a *VrfBgpAdvertiseNetworkIpv4, b *VrfBgpAdvertiseNetworkIpv4) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpAdvertiseNetworkIpv4Network(a.Network, b.Network) {
		return false
	}
	return true
}
func matchVrfBgpAdvertiseNetworkIpv6Network(a []VrfBgpAdvertiseNetworkIpv6Network, b []VrfBgpAdvertiseNetworkIpv6Network) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Unicast, b.Unicast) {
				return false
			}
		}
	}
	return true
}
func matchVrfBgpAdvertiseNetworkIpv6(a *VrfBgpAdvertiseNetworkIpv6, b *VrfBgpAdvertiseNetworkIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpAdvertiseNetworkIpv6Network(a.Network, b.Network) {
		return false
	}
	return true
}
func matchVrfBgpAdvertiseNetwork(a *VrfBgpAdvertiseNetwork, b *VrfBgpAdvertiseNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpAdvertiseNetworkIpv4(a.Ipv4, b.Ipv4) {
		return false
	}
	if !matchVrfBgpAdvertiseNetworkIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupTypeIbgp(a *VrfBgpPeerGroupTypeIbgp, b *VrfBgpPeerGroupTypeIbgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfBgpPeerGroupTypeEbgp(a *VrfBgpPeerGroupTypeEbgp, b *VrfBgpPeerGroupTypeEbgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfBgpPeerGroupType(a *VrfBgpPeerGroupType, b *VrfBgpPeerGroupType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpPeerGroupTypeIbgp(a.Ibgp, b.Ibgp) {
		return false
	}
	if !matchVrfBgpPeerGroupTypeEbgp(a.Ebgp, b.Ebgp) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupAddressFamily(a *VrfBgpPeerGroupAddressFamily, b *VrfBgpPeerGroupAddressFamily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupFilteringProfile(a *VrfBgpPeerGroupFilteringProfile, b *VrfBgpPeerGroupFilteringProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupConnectionOptions(a *VrfBgpPeerGroupConnectionOptions, b *VrfBgpPeerGroupConnectionOptions) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Timers, b.Timers) {
		return false
	}
	if !util.Ints64Match(a.Multihop, b.Multihop) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.Dampening, b.Dampening) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerInheritYes(a *VrfBgpPeerGroupPeerInheritYes, b *VrfBgpPeerGroupPeerInheritYes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfBgpPeerGroupPeerInheritNoAddressFamily(a *VrfBgpPeerGroupPeerInheritNoAddressFamily, b *VrfBgpPeerGroupPeerInheritNoAddressFamily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerInheritNoFilteringProfile(a *VrfBgpPeerGroupPeerInheritNoFilteringProfile, b *VrfBgpPeerGroupPeerInheritNoFilteringProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerInheritNo(a *VrfBgpPeerGroupPeerInheritNo, b *VrfBgpPeerGroupPeerInheritNo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpPeerGroupPeerInheritNoAddressFamily(a.AddressFamily, b.AddressFamily) {
		return false
	}
	if !matchVrfBgpPeerGroupPeerInheritNoFilteringProfile(a.FilteringProfile, b.FilteringProfile) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerInherit(a *VrfBgpPeerGroupPeerInherit, b *VrfBgpPeerGroupPeerInherit) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpPeerGroupPeerInheritYes(a.Yes, b.Yes) {
		return false
	}
	if !matchVrfBgpPeerGroupPeerInheritNo(a.No, b.No) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerLocalAddress(a *VrfBgpPeerGroupPeerLocalAddress, b *VrfBgpPeerGroupPeerLocalAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerPeerAddress(a *VrfBgpPeerGroupPeerPeerAddress, b *VrfBgpPeerGroupPeerPeerAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerConnectionOptions(a *VrfBgpPeerGroupPeerConnectionOptions, b *VrfBgpPeerGroupPeerConnectionOptions) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Timers, b.Timers) {
		return false
	}
	if !util.StringsMatch(a.Multihop, b.Multihop) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.Dampening, b.Dampening) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeerBfd(a *VrfBgpPeerGroupPeerBfd, b *VrfBgpPeerGroupPeerBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfBgpPeerGroupPeer(a []VrfBgpPeerGroupPeer, b []VrfBgpPeerGroupPeer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.Passive, b.Passive) {
				return false
			}
			if !util.StringsMatch(a.PeerAs, b.PeerAs) {
				return false
			}
			if !util.BoolsMatch(a.EnableSenderSideLoopDetection, b.EnableSenderSideLoopDetection) {
				return false
			}
			if !matchVrfBgpPeerGroupPeerInherit(a.Inherit, b.Inherit) {
				return false
			}
			if !matchVrfBgpPeerGroupPeerLocalAddress(a.LocalAddress, b.LocalAddress) {
				return false
			}
			if !matchVrfBgpPeerGroupPeerPeerAddress(a.PeerAddress, b.PeerAddress) {
				return false
			}
			if !matchVrfBgpPeerGroupPeerConnectionOptions(a.ConnectionOptions, b.ConnectionOptions) {
				return false
			}
			if !matchVrfBgpPeerGroupPeerBfd(a.Bfd, b.Bfd) {
				return false
			}
		}
	}
	return true
}
func matchVrfBgpPeerGroup(a []VrfBgpPeerGroup, b []VrfBgpPeerGroup) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !matchVrfBgpPeerGroupType(a.Type, b.Type) {
				return false
			}
			if !matchVrfBgpPeerGroupAddressFamily(a.AddressFamily, b.AddressFamily) {
				return false
			}
			if !matchVrfBgpPeerGroupFilteringProfile(a.FilteringProfile, b.FilteringProfile) {
				return false
			}
			if !matchVrfBgpPeerGroupConnectionOptions(a.ConnectionOptions, b.ConnectionOptions) {
				return false
			}
			if !matchVrfBgpPeerGroupPeer(a.Peer, b.Peer) {
				return false
			}
		}
	}
	return true
}
func matchVrfBgpAggregateRoutesTypeIpv4(a *VrfBgpAggregateRoutesTypeIpv4, b *VrfBgpAggregateRoutesTypeIpv4) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SummaryPrefix, b.SummaryPrefix) {
		return false
	}
	if !util.StringsMatch(a.SuppressMap, b.SuppressMap) {
		return false
	}
	if !util.StringsMatch(a.AttributeMap, b.AttributeMap) {
		return false
	}
	return true
}
func matchVrfBgpAggregateRoutesTypeIpv6(a *VrfBgpAggregateRoutesTypeIpv6, b *VrfBgpAggregateRoutesTypeIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SummaryPrefix, b.SummaryPrefix) {
		return false
	}
	if !util.StringsMatch(a.SuppressMap, b.SuppressMap) {
		return false
	}
	if !util.StringsMatch(a.AttributeMap, b.AttributeMap) {
		return false
	}
	return true
}
func matchVrfBgpAggregateRoutesType(a *VrfBgpAggregateRoutesType, b *VrfBgpAggregateRoutesType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfBgpAggregateRoutesTypeIpv4(a.Ipv4, b.Ipv4) {
		return false
	}
	if !matchVrfBgpAggregateRoutesTypeIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfBgpAggregateRoutes(a []VrfBgpAggregateRoutes, b []VrfBgpAggregateRoutes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.SummaryOnly, b.SummaryOnly) {
				return false
			}
			if !util.BoolsMatch(a.AsSet, b.AsSet) {
				return false
			}
			if !util.BoolsMatch(a.SameMed, b.SameMed) {
				return false
			}
			if !matchVrfBgpAggregateRoutesType(a.Type, b.Type) {
				return false
			}
		}
	}
	return true
}
func matchVrfBgp(a *VrfBgp, b *VrfBgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !util.StringsMatch(a.LocalAs, b.LocalAs) {
		return false
	}
	if !util.BoolsMatch(a.InstallRoute, b.InstallRoute) {
		return false
	}
	if !util.BoolsMatch(a.EnforceFirstAs, b.EnforceFirstAs) {
		return false
	}
	if !util.BoolsMatch(a.FastExternalFailover, b.FastExternalFailover) {
		return false
	}
	if !util.BoolsMatch(a.EcmpMultiAs, b.EcmpMultiAs) {
		return false
	}
	if !util.Ints64Match(a.DefaultLocalPreference, b.DefaultLocalPreference) {
		return false
	}
	if !util.BoolsMatch(a.GracefulShutdown, b.GracefulShutdown) {
		return false
	}
	if !util.BoolsMatch(a.AlwaysAdvertiseNetworkRoute, b.AlwaysAdvertiseNetworkRoute) {
		return false
	}
	if !matchVrfBgpMed(a.Med, b.Med) {
		return false
	}
	if !matchVrfBgpGracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	if !matchVrfBgpGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !matchVrfBgpRedistributionProfile(a.RedistributionProfile, b.RedistributionProfile) {
		return false
	}
	if !matchVrfBgpAdvertiseNetwork(a.AdvertiseNetwork, b.AdvertiseNetwork) {
		return false
	}
	if !matchVrfBgpPeerGroup(a.PeerGroup, b.PeerGroup) {
		return false
	}
	if !matchVrfBgpAggregateRoutes(a.AggregateRoutes, b.AggregateRoutes) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpStaticRouteNexthopDiscard(a *VrfRoutingTableIpStaticRouteNexthopDiscard, b *VrfRoutingTableIpStaticRouteNexthopDiscard) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfRoutingTableIpStaticRouteNexthop(a *VrfRoutingTableIpStaticRouteNexthop, b *VrfRoutingTableIpStaticRouteNexthop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRoutingTableIpStaticRouteNexthopDiscard(a.Discard, b.Discard) {
		return false
	}
	if !util.StringsMatch(a.IpAddress, b.IpAddress) {
		return false
	}
	if !util.StringsMatch(a.NextLr, b.NextLr) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpStaticRouteBfd(a *VrfRoutingTableIpStaticRouteBfd, b *VrfRoutingTableIpStaticRouteBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations(a []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations, b []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.Source, b.Source) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.Ints64Match(a.Interval, b.Interval) {
				return false
			}
			if !util.Ints64Match(a.Count, b.Count) {
				return false
			}
		}
	}
	return true
}
func matchVrfRoutingTableIpStaticRoutePathMonitor(a *VrfRoutingTableIpStaticRoutePathMonitor, b *VrfRoutingTableIpStaticRoutePathMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.FailureCondition, b.FailureCondition) {
		return false
	}
	if !util.Ints64Match(a.HoldTime, b.HoldTime) {
		return false
	}
	if !matchVrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations(a.MonitorDestinations, b.MonitorDestinations) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpStaticRoute(a []VrfRoutingTableIpStaticRoute, b []VrfRoutingTableIpStaticRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !matchVrfRoutingTableIpStaticRouteNexthop(a.Nexthop, b.Nexthop) {
				return false
			}
			if !matchVrfRoutingTableIpStaticRouteBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !matchVrfRoutingTableIpStaticRoutePathMonitor(a.PathMonitor, b.PathMonitor) {
				return false
			}
		}
	}
	return true
}
func matchVrfRoutingTableIp(a *VrfRoutingTableIp, b *VrfRoutingTableIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRoutingTableIpStaticRoute(a.StaticRoute, b.StaticRoute) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRouteNexthopDiscard(a *VrfRoutingTableIpv6StaticRouteNexthopDiscard, b *VrfRoutingTableIpv6StaticRouteNexthopDiscard) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRouteNexthop(a *VrfRoutingTableIpv6StaticRouteNexthop, b *VrfRoutingTableIpv6StaticRouteNexthop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRoutingTableIpv6StaticRouteNexthopDiscard(a.Discard, b.Discard) {
		return false
	}
	if !util.StringsMatch(a.Ipv6Address, b.Ipv6Address) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.NextLr, b.NextLr) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRouteBfd(a *VrfRoutingTableIpv6StaticRouteBfd, b *VrfRoutingTableIpv6StaticRouteBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations(a []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations, b []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.Source, b.Source) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.Ints64Match(a.Interval, b.Interval) {
				return false
			}
			if !util.Ints64Match(a.Count, b.Count) {
				return false
			}
		}
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRoutePathMonitor(a *VrfRoutingTableIpv6StaticRoutePathMonitor, b *VrfRoutingTableIpv6StaticRoutePathMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.FailureCondition, b.FailureCondition) {
		return false
	}
	if !util.Ints64Match(a.HoldTime, b.HoldTime) {
		return false
	}
	if !matchVrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations(a.MonitorDestinations, b.MonitorDestinations) {
		return false
	}
	return true
}
func matchVrfRoutingTableIpv6StaticRoute(a []VrfRoutingTableIpv6StaticRoute, b []VrfRoutingTableIpv6StaticRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !matchVrfRoutingTableIpv6StaticRouteNexthop(a.Nexthop, b.Nexthop) {
				return false
			}
			if !matchVrfRoutingTableIpv6StaticRouteBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !matchVrfRoutingTableIpv6StaticRoutePathMonitor(a.PathMonitor, b.PathMonitor) {
				return false
			}
		}
	}
	return true
}
func matchVrfRoutingTableIpv6(a *VrfRoutingTableIpv6, b *VrfRoutingTableIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRoutingTableIpv6StaticRoute(a.StaticRoute, b.StaticRoute) {
		return false
	}
	return true
}
func matchVrfRoutingTable(a *VrfRoutingTable, b *VrfRoutingTable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfRoutingTableIp(a.Ip, b.Ip) {
		return false
	}
	if !matchVrfRoutingTableIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchVrfOspfGlobalBfd(a *VrfOspfGlobalBfd, b *VrfOspfGlobalBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfOspfGracefulRestart(a *VrfOspfGracefulRestart, b *VrfOspfGracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.GracePeriod, b.GracePeriod) {
		return false
	}
	if !util.BoolsMatch(a.HelperEnable, b.HelperEnable) {
		return false
	}
	if !util.BoolsMatch(a.StrictLSAChecking, b.StrictLSAChecking) {
		return false
	}
	if !util.Ints64Match(a.MaxNeighborRestartTime, b.MaxNeighborRestartTime) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeNormalAbr(a *VrfOspfAreaTypeNormalAbr, b *VrfOspfAreaTypeNormalAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeNormal(a *VrfOspfAreaTypeNormal, b *VrfOspfAreaTypeNormal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfAreaTypeNormalAbr(a.Abr, b.Abr) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeStubAbr(a *VrfOspfAreaTypeStubAbr, b *VrfOspfAreaTypeStubAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeStub(a *VrfOspfAreaTypeStub, b *VrfOspfAreaTypeStub) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.NoSummary, b.NoSummary) {
		return false
	}
	if !matchVrfOspfAreaTypeStubAbr(a.Abr, b.Abr) {
		return false
	}
	if !util.Ints64Match(a.DefaultRouteMetric, b.DefaultRouteMetric) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeNssaDefaultInformationOriginate(a *VrfOspfAreaTypeNssaDefaultInformationOriginate, b *VrfOspfAreaTypeNssaDefaultInformationOriginate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	if !util.StringsMatch(a.MetricType, b.MetricType) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeNssaAbrNssaExtRange(a []VrfOspfAreaTypeNssaAbrNssaExtRange, b []VrfOspfAreaTypeNssaAbrNssaExtRange) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfAreaTypeNssaAbr(a *VrfOspfAreaTypeNssaAbr, b *VrfOspfAreaTypeNssaAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	if !matchVrfOspfAreaTypeNssaAbrNssaExtRange(a.NssaExtRange, b.NssaExtRange) {
		return false
	}
	return true
}
func matchVrfOspfAreaTypeNssa(a *VrfOspfAreaTypeNssa, b *VrfOspfAreaTypeNssa) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.NoSummary, b.NoSummary) {
		return false
	}
	if !matchVrfOspfAreaTypeNssaDefaultInformationOriginate(a.DefaultInformationOriginate, b.DefaultInformationOriginate) {
		return false
	}
	if !matchVrfOspfAreaTypeNssaAbr(a.Abr, b.Abr) {
		return false
	}
	return true
}
func matchVrfOspfAreaType(a *VrfOspfAreaType, b *VrfOspfAreaType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfAreaTypeNormal(a.Normal, b.Normal) {
		return false
	}
	if !matchVrfOspfAreaTypeStub(a.Stub, b.Stub) {
		return false
	}
	if !matchVrfOspfAreaTypeNssa(a.Nssa, b.Nssa) {
		return false
	}
	return true
}
func matchVrfOspfAreaRange(a []VrfOspfAreaRange, b []VrfOspfAreaRange) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfAreaInterfaceLinkTypeBroadcast(a *VrfOspfAreaInterfaceLinkTypeBroadcast, b *VrfOspfAreaInterfaceLinkTypeBroadcast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfOspfAreaInterfaceLinkTypeP2p(a *VrfOspfAreaInterfaceLinkTypeP2p, b *VrfOspfAreaInterfaceLinkTypeP2p) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfOspfAreaInterfaceLinkTypeP2mpNeighbor(a []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor, b []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfAreaInterfaceLinkTypeP2mp(a *VrfOspfAreaInterfaceLinkTypeP2mp, b *VrfOspfAreaInterfaceLinkTypeP2mp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfAreaInterfaceLinkTypeP2mpNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	return true
}
func matchVrfOspfAreaInterfaceLinkType(a *VrfOspfAreaInterfaceLinkType, b *VrfOspfAreaInterfaceLinkType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfAreaInterfaceLinkTypeBroadcast(a.Broadcast, b.Broadcast) {
		return false
	}
	if !matchVrfOspfAreaInterfaceLinkTypeP2p(a.P2p, b.P2p) {
		return false
	}
	if !matchVrfOspfAreaInterfaceLinkTypeP2mp(a.P2mp, b.P2mp) {
		return false
	}
	return true
}
func matchVrfOspfAreaInterfaceBfd(a *VrfOspfAreaInterfaceBfd, b *VrfOspfAreaInterfaceBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfOspfAreaInterface(a []VrfOspfAreaInterface, b []VrfOspfAreaInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.MtuIgnore, b.MtuIgnore) {
				return false
			}
			if !util.BoolsMatch(a.Passive, b.Passive) {
				return false
			}
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.StringsMatch(a.Timing, b.Timing) {
				return false
			}
			if !matchVrfOspfAreaInterfaceLinkType(a.LinkType, b.LinkType) {
				return false
			}
			if !matchVrfOspfAreaInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfAreaVirtualLinkBfd(a *VrfOspfAreaVirtualLinkBfd, b *VrfOspfAreaVirtualLinkBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfOspfAreaVirtualLink(a []VrfOspfAreaVirtualLink, b []VrfOspfAreaVirtualLink) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.NeighborId, b.NeighborId) {
				return false
			}
			if !util.StringsMatch(a.TransitAreaId, b.TransitAreaId) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.Ints64Match(a.InstanceId, b.InstanceId) {
				return false
			}
			if !util.StringsMatch(a.Timing, b.Timing) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchVrfOspfAreaVirtualLinkBfd(a.Bfd, b.Bfd) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfArea(a []VrfOspfArea, b []VrfOspfArea) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchVrfOspfAreaType(a.Type, b.Type) {
				return false
			}
			if !matchVrfOspfAreaRange(a.Range, b.Range) {
				return false
			}
			if !matchVrfOspfAreaInterface(a.Interface, b.Interface) {
				return false
			}
			if !matchVrfOspfAreaVirtualLink(a.VirtualLink, b.VirtualLink) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspf(a *VrfOspf, b *VrfOspf) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.Rfc1583, b.Rfc1583) {
		return false
	}
	if !util.StringsMatch(a.SpfTimer, b.SpfTimer) {
		return false
	}
	if !util.StringsMatch(a.GlobalIfTimer, b.GlobalIfTimer) {
		return false
	}
	if !util.StringsMatch(a.RedistributionProfile, b.RedistributionProfile) {
		return false
	}
	if !matchVrfOspfGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !matchVrfOspfGracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	if !matchVrfOspfArea(a.Area, b.Area) {
		return false
	}
	return true
}
func matchVrfOspfv3GlobalBfd(a *VrfOspfv3GlobalBfd, b *VrfOspfv3GlobalBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfOspfv3GracefulRestart(a *VrfOspfv3GracefulRestart, b *VrfOspfv3GracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.GracePeriod, b.GracePeriod) {
		return false
	}
	if !util.BoolsMatch(a.HelperEnable, b.HelperEnable) {
		return false
	}
	if !util.BoolsMatch(a.StrictLSAChecking, b.StrictLSAChecking) {
		return false
	}
	if !util.Ints64Match(a.MaxNeighborRestartTime, b.MaxNeighborRestartTime) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeNormalAbr(a *VrfOspfv3AreaTypeNormalAbr, b *VrfOspfv3AreaTypeNormalAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeNormal(a *VrfOspfv3AreaTypeNormal, b *VrfOspfv3AreaTypeNormal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfv3AreaTypeNormalAbr(a.Abr, b.Abr) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeStubAbr(a *VrfOspfv3AreaTypeStubAbr, b *VrfOspfv3AreaTypeStubAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeStub(a *VrfOspfv3AreaTypeStub, b *VrfOspfv3AreaTypeStub) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.NoSummary, b.NoSummary) {
		return false
	}
	if !matchVrfOspfv3AreaTypeStubAbr(a.Abr, b.Abr) {
		return false
	}
	if !util.Ints64Match(a.DefaultRouteMetric, b.DefaultRouteMetric) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeNssaDefaultInformationOriginate(a *VrfOspfv3AreaTypeNssaDefaultInformationOriginate, b *VrfOspfv3AreaTypeNssaDefaultInformationOriginate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	if !util.StringsMatch(a.MetricType, b.MetricType) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeNssaAbrNssaExtRange(a []VrfOspfv3AreaTypeNssaAbrNssaExtRange, b []VrfOspfv3AreaTypeNssaAbrNssaExtRange) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3AreaTypeNssaAbr(a *VrfOspfv3AreaTypeNssaAbr, b *VrfOspfv3AreaTypeNssaAbr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ImportList, b.ImportList) {
		return false
	}
	if !util.StringsMatch(a.ExportList, b.ExportList) {
		return false
	}
	if !util.StringsMatch(a.InboundFilterList, b.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(a.OutboundFilterList, b.OutboundFilterList) {
		return false
	}
	if !matchVrfOspfv3AreaTypeNssaAbrNssaExtRange(a.NssaExtRange, b.NssaExtRange) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaTypeNssa(a *VrfOspfv3AreaTypeNssa, b *VrfOspfv3AreaTypeNssa) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.NoSummary, b.NoSummary) {
		return false
	}
	if !matchVrfOspfv3AreaTypeNssaDefaultInformationOriginate(a.DefaultInformationOriginate, b.DefaultInformationOriginate) {
		return false
	}
	if !matchVrfOspfv3AreaTypeNssaAbr(a.Abr, b.Abr) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaType(a *VrfOspfv3AreaType, b *VrfOspfv3AreaType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfv3AreaTypeNormal(a.Normal, b.Normal) {
		return false
	}
	if !matchVrfOspfv3AreaTypeStub(a.Stub, b.Stub) {
		return false
	}
	if !matchVrfOspfv3AreaTypeNssa(a.Nssa, b.Nssa) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaRange(a []VrfOspfv3AreaRange, b []VrfOspfv3AreaRange) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3AreaInterfaceLinkTypeBroadcast(a *VrfOspfv3AreaInterfaceLinkTypeBroadcast, b *VrfOspfv3AreaInterfaceLinkTypeBroadcast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfOspfv3AreaInterfaceLinkTypeP2p(a *VrfOspfv3AreaInterfaceLinkTypeP2p, b *VrfOspfv3AreaInterfaceLinkTypeP2p) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor(a []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor, b []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3AreaInterfaceLinkTypeP2mp(a *VrfOspfv3AreaInterfaceLinkTypeP2mp, b *VrfOspfv3AreaInterfaceLinkTypeP2mp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaInterfaceLinkType(a *VrfOspfv3AreaInterfaceLinkType, b *VrfOspfv3AreaInterfaceLinkType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfOspfv3AreaInterfaceLinkTypeBroadcast(a.Broadcast, b.Broadcast) {
		return false
	}
	if !matchVrfOspfv3AreaInterfaceLinkTypeP2p(a.P2p, b.P2p) {
		return false
	}
	if !matchVrfOspfv3AreaInterfaceLinkTypeP2mp(a.P2mp, b.P2mp) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaInterfaceBfd(a *VrfOspfv3AreaInterfaceBfd, b *VrfOspfv3AreaInterfaceBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfOspfv3AreaInterface(a []VrfOspfv3AreaInterface, b []VrfOspfv3AreaInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.MtuIgnore, b.MtuIgnore) {
				return false
			}
			if !util.BoolsMatch(a.Passive, b.Passive) {
				return false
			}
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.Ints64Match(a.InstanceId, b.InstanceId) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.StringsMatch(a.Timing, b.Timing) {
				return false
			}
			if !matchVrfOspfv3AreaInterfaceLinkType(a.LinkType, b.LinkType) {
				return false
			}
			if !matchVrfOspfv3AreaInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3AreaVirtualLink(a []VrfOspfv3AreaVirtualLink, b []VrfOspfv3AreaVirtualLink) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.NeighborId, b.NeighborId) {
				return false
			}
			if !util.StringsMatch(a.TransitAreaId, b.TransitAreaId) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.Ints64Match(a.InstanceId, b.InstanceId) {
				return false
			}
			if !util.StringsMatch(a.Timing, b.Timing) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3Area(a []VrfOspfv3Area, b []VrfOspfv3Area) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchVrfOspfv3AreaType(a.Type, b.Type) {
				return false
			}
			if !matchVrfOspfv3AreaRange(a.Range, b.Range) {
				return false
			}
			if !matchVrfOspfv3AreaInterface(a.Interface, b.Interface) {
				return false
			}
			if !matchVrfOspfv3AreaVirtualLink(a.VirtualLink, b.VirtualLink) {
				return false
			}
		}
	}
	return true
}
func matchVrfOspfv3(a *VrfOspfv3, b *VrfOspfv3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !util.BoolsMatch(a.DisableTransitTraffic, b.DisableTransitTraffic) {
		return false
	}
	if !util.StringsMatch(a.SpfTimer, b.SpfTimer) {
		return false
	}
	if !util.StringsMatch(a.GlobalIfTimer, b.GlobalIfTimer) {
		return false
	}
	if !util.StringsMatch(a.RedistributionProfile, b.RedistributionProfile) {
		return false
	}
	if !matchVrfOspfv3GlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !matchVrfOspfv3GracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	if !matchVrfOspfv3Area(a.Area, b.Area) {
		return false
	}
	return true
}
func matchVrfEcmpAlgorithmIpModulo(a *VrfEcmpAlgorithmIpModulo, b *VrfEcmpAlgorithmIpModulo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfEcmpAlgorithmIpHash(a *VrfEcmpAlgorithmIpHash, b *VrfEcmpAlgorithmIpHash) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.SrcOnly, b.SrcOnly) {
		return false
	}
	if !util.BoolsMatch(a.UsePort, b.UsePort) {
		return false
	}
	if !util.Ints64Match(a.HashSeed, b.HashSeed) {
		return false
	}
	return true
}
func matchVrfEcmpAlgorithmWeightedRoundRobinInterface(a []VrfEcmpAlgorithmWeightedRoundRobinInterface, b []VrfEcmpAlgorithmWeightedRoundRobinInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.Ints64Match(a.Weight, b.Weight) {
				return false
			}
		}
	}
	return true
}
func matchVrfEcmpAlgorithmWeightedRoundRobin(a *VrfEcmpAlgorithmWeightedRoundRobin, b *VrfEcmpAlgorithmWeightedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfEcmpAlgorithmWeightedRoundRobinInterface(a.Interface, b.Interface) {
		return false
	}
	return true
}
func matchVrfEcmpAlgorithmBalancedRoundRobin(a *VrfEcmpAlgorithmBalancedRoundRobin, b *VrfEcmpAlgorithmBalancedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchVrfEcmpAlgorithm(a *VrfEcmpAlgorithm, b *VrfEcmpAlgorithm) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfEcmpAlgorithmIpModulo(a.IpModulo, b.IpModulo) {
		return false
	}
	if !matchVrfEcmpAlgorithmIpHash(a.IpHash, b.IpHash) {
		return false
	}
	if !matchVrfEcmpAlgorithmWeightedRoundRobin(a.WeightedRoundRobin, b.WeightedRoundRobin) {
		return false
	}
	if !matchVrfEcmpAlgorithmBalancedRoundRobin(a.BalancedRoundRobin, b.BalancedRoundRobin) {
		return false
	}
	return true
}
func matchVrfEcmp(a *VrfEcmp, b *VrfEcmp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.MaxPath, b.MaxPath) {
		return false
	}
	if !util.BoolsMatch(a.SymmetricReturn, b.SymmetricReturn) {
		return false
	}
	if !util.BoolsMatch(a.StrictSourcePath, b.StrictSourcePath) {
		return false
	}
	if !matchVrfEcmpAlgorithm(a.Algorithm, b.Algorithm) {
		return false
	}
	return true
}
func matchVrfMulticastStaticRouteNexthop(a *VrfMulticastStaticRouteNexthop, b *VrfMulticastStaticRouteNexthop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.IpAddress, b.IpAddress) {
		return false
	}
	return true
}
func matchVrfMulticastStaticRoute(a []VrfMulticastStaticRoute, b []VrfMulticastStaticRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !util.Ints64Match(a.Preference, b.Preference) {
				return false
			}
			if !matchVrfMulticastStaticRouteNexthop(a.Nexthop, b.Nexthop) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastPimSsmAddressSpace(a *VrfMulticastPimSsmAddressSpace, b *VrfMulticastPimSsmAddressSpace) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.GroupList, b.GroupList) {
		return false
	}
	return true
}
func matchVrfMulticastPimRpLocalRpStaticRp(a *VrfMulticastPimRpLocalRpStaticRp, b *VrfMulticastPimRpLocalRpStaticRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Override, b.Override) {
		return false
	}
	if !util.StringsMatch(a.GroupList, b.GroupList) {
		return false
	}
	return true
}
func matchVrfMulticastPimRpLocalRpCandidateRp(a *VrfMulticastPimRpLocalRpCandidateRp, b *VrfMulticastPimRpLocalRpCandidateRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.Address, b.Address) {
		return false
	}
	if !util.Ints64Match(a.Priority, b.Priority) {
		return false
	}
	if !util.Ints64Match(a.AdvertisementInterval, b.AdvertisementInterval) {
		return false
	}
	if !util.StringsMatch(a.GroupList, b.GroupList) {
		return false
	}
	return true
}
func matchVrfMulticastPimRpLocalRp(a *VrfMulticastPimRpLocalRp, b *VrfMulticastPimRpLocalRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfMulticastPimRpLocalRpStaticRp(a.StaticRp, b.StaticRp) {
		return false
	}
	if !matchVrfMulticastPimRpLocalRpCandidateRp(a.CandidateRp, b.CandidateRp) {
		return false
	}
	return true
}
func matchVrfMulticastPimRpExternalRp(a []VrfMulticastPimRpExternalRp, b []VrfMulticastPimRpExternalRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.GroupList, b.GroupList) {
				return false
			}
			if !util.BoolsMatch(a.Override, b.Override) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastPimRp(a *VrfMulticastPimRp, b *VrfMulticastPimRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfMulticastPimRpLocalRp(a.LocalRp, b.LocalRp) {
		return false
	}
	if !matchVrfMulticastPimRpExternalRp(a.ExternalRp, b.ExternalRp) {
		return false
	}
	return true
}
func matchVrfMulticastPimSptThreshold(a []VrfMulticastPimSptThreshold, b []VrfMulticastPimSptThreshold) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Threshold, b.Threshold) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastPimInterface(a []VrfMulticastPimInterface, b []VrfMulticastPimInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
			if !util.Ints64Match(a.DrPriority, b.DrPriority) {
				return false
			}
			if !util.BoolsMatch(a.SendBsm, b.SendBsm) {
				return false
			}
			if !util.StringsMatch(a.IfTimer, b.IfTimer) {
				return false
			}
			if !util.StringsMatch(a.NeighborFilter, b.NeighborFilter) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastPim(a *VrfMulticastPim, b *VrfMulticastPim) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.RpfLookupMode, b.RpfLookupMode) {
		return false
	}
	if !util.Ints64Match(a.RouteAgeoutTime, b.RouteAgeoutTime) {
		return false
	}
	if !util.StringsMatch(a.IfTimerGlobal, b.IfTimerGlobal) {
		return false
	}
	if !util.StringsMatch(a.GroupPermission, b.GroupPermission) {
		return false
	}
	if !matchVrfMulticastPimSsmAddressSpace(a.SsmAddressSpace, b.SsmAddressSpace) {
		return false
	}
	if !matchVrfMulticastPimRp(a.Rp, b.Rp) {
		return false
	}
	if !matchVrfMulticastPimSptThreshold(a.SptThreshold, b.SptThreshold) {
		return false
	}
	if !matchVrfMulticastPimInterface(a.Interface, b.Interface) {
		return false
	}
	return true
}
func matchVrfMulticastIgmpDynamicInterface(a []VrfMulticastIgmpDynamicInterface, b []VrfMulticastIgmpDynamicInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Version, b.Version) {
				return false
			}
			if !util.StringsMatch(a.Robustness, b.Robustness) {
				return false
			}
			if !util.StringsMatch(a.GroupFilter, b.GroupFilter) {
				return false
			}
			if !util.StringsMatch(a.MaxGroups, b.MaxGroups) {
				return false
			}
			if !util.StringsMatch(a.MaxSources, b.MaxSources) {
				return false
			}
			if !util.StringsMatch(a.QueryProfile, b.QueryProfile) {
				return false
			}
			if !util.BoolsMatch(a.RouterAlertPolicing, b.RouterAlertPolicing) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastIgmpDynamic(a *VrfMulticastIgmpDynamic, b *VrfMulticastIgmpDynamic) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVrfMulticastIgmpDynamicInterface(a.Interface, b.Interface) {
		return false
	}
	return true
}
func matchVrfMulticastIgmpStatic(a []VrfMulticastIgmpStatic, b []VrfMulticastIgmpStatic) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !util.StringsMatch(a.GroupAddress, b.GroupAddress) {
				return false
			}
			if !util.StringsMatch(a.SourceAddress, b.SourceAddress) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastIgmp(a *VrfMulticastIgmp, b *VrfMulticastIgmp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchVrfMulticastIgmpDynamic(a.Dynamic, b.Dynamic) {
		return false
	}
	if !matchVrfMulticastIgmpStatic(a.Static, b.Static) {
		return false
	}
	return true
}
func matchVrfMulticastMsdpOriginatorId(a *VrfMulticastMsdpOriginatorId, b *VrfMulticastMsdpOriginatorId) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchVrfMulticastMsdpPeerLocalAddress(a *VrfMulticastMsdpPeerLocalAddress, b *VrfMulticastMsdpPeerLocalAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchVrfMulticastMsdpPeerPeerAddress(a *VrfMulticastMsdpPeerPeerAddress, b *VrfMulticastMsdpPeerPeerAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	return true
}
func matchVrfMulticastMsdpPeer(a []VrfMulticastMsdpPeer, b []VrfMulticastMsdpPeer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.PeerAs, b.PeerAs) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.Ints64Match(a.MaxSa, b.MaxSa) {
				return false
			}
			if !util.StringsMatch(a.InboundSaFilter, b.InboundSaFilter) {
				return false
			}
			if !util.StringsMatch(a.OutboundSaFilter, b.OutboundSaFilter) {
				return false
			}
			if !matchVrfMulticastMsdpPeerLocalAddress(a.LocalAddress, b.LocalAddress) {
				return false
			}
			if !matchVrfMulticastMsdpPeerPeerAddress(a.PeerAddress, b.PeerAddress) {
				return false
			}
		}
	}
	return true
}
func matchVrfMulticastMsdp(a *VrfMulticastMsdp, b *VrfMulticastMsdp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.GlobalTimer, b.GlobalTimer) {
		return false
	}
	if !util.StringsMatch(a.GlobalAuthentication, b.GlobalAuthentication) {
		return false
	}
	if !matchVrfMulticastMsdpOriginatorId(a.OriginatorId, b.OriginatorId) {
		return false
	}
	if !matchVrfMulticastMsdpPeer(a.Peer, b.Peer) {
		return false
	}
	return true
}
func matchVrfMulticast(a *VrfMulticast, b *VrfMulticast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchVrfMulticastStaticRoute(a.StaticRoute, b.StaticRoute) {
		return false
	}
	if !matchVrfMulticastPim(a.Pim, b.Pim) {
		return false
	}
	if !matchVrfMulticastIgmp(a.Igmp, b.Igmp) {
		return false
	}
	if !matchVrfMulticastMsdp(a.Msdp, b.Msdp) {
		return false
	}
	return true
}
func matchVrfRipGlobalBfd(a *VrfRipGlobalBfd, b *VrfRipGlobalBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfRipGlobalInboundDistributeList(a *VrfRipGlobalInboundDistributeList, b *VrfRipGlobalInboundDistributeList) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AccessList, b.AccessList) {
		return false
	}
	return true
}
func matchVrfRipGlobalOutboundDistributeList(a *VrfRipGlobalOutboundDistributeList, b *VrfRipGlobalOutboundDistributeList) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AccessList, b.AccessList) {
		return false
	}
	return true
}
func matchVrfRipInterfaceBfd(a *VrfRipInterfaceBfd, b *VrfRipInterfaceBfd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVrfRipInterfaceInterfaceInboundDistributeList(a *VrfRipInterfaceInterfaceInboundDistributeList, b *VrfRipInterfaceInterfaceInboundDistributeList) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AccessList, b.AccessList) {
		return false
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	return true
}
func matchVrfRipInterfaceInterfaceOutboundDistributeList(a *VrfRipInterfaceInterfaceOutboundDistributeList, b *VrfRipInterfaceInterfaceOutboundDistributeList) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AccessList, b.AccessList) {
		return false
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	return true
}
func matchVrfRipInterface(a []VrfRipInterface, b []VrfRipInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.Mode, b.Mode) {
				return false
			}
			if !util.StringsMatch(a.SplitHorizon, b.SplitHorizon) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchVrfRipInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !matchVrfRipInterfaceInterfaceInboundDistributeList(a.InterfaceInboundDistributeList, b.InterfaceInboundDistributeList) {
				return false
			}
			if !matchVrfRipInterfaceInterfaceOutboundDistributeList(a.InterfaceOutboundDistributeList, b.InterfaceOutboundDistributeList) {
				return false
			}
		}
	}
	return true
}
func matchVrfRip(a *VrfRip, b *VrfRip) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.DefaultInformationOriginate, b.DefaultInformationOriginate) {
		return false
	}
	if !util.StringsMatch(a.GlobalTimer, b.GlobalTimer) {
		return false
	}
	if !util.StringsMatch(a.AuthProfile, b.AuthProfile) {
		return false
	}
	if !util.StringsMatch(a.RedistributionProfile, b.RedistributionProfile) {
		return false
	}
	if !matchVrfRipGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !matchVrfRipGlobalInboundDistributeList(a.GlobalInboundDistributeList, b.GlobalInboundDistributeList) {
		return false
	}
	if !matchVrfRipGlobalOutboundDistributeList(a.GlobalOutboundDistributeList, b.GlobalOutboundDistributeList) {
		return false
	}
	if !matchVrfRipInterface(a.Interface, b.Interface) {
		return false
	}
	return true
}
func matchVrf(a []Vrf, b []Vrf) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.OrderedListsMatch(a.Interface, b.Interface) {
				return false
			}
			if !matchVrfAdminDists(a.AdminDists, b.AdminDists) {
				return false
			}
			if !matchVrfRibFilter(a.RibFilter, b.RibFilter) {
				return false
			}
			if !matchVrfBgp(a.Bgp, b.Bgp) {
				return false
			}
			if !matchVrfRoutingTable(a.RoutingTable, b.RoutingTable) {
				return false
			}
			if !matchVrfOspf(a.Ospf, b.Ospf) {
				return false
			}
			if !matchVrfOspfv3(a.Ospfv3, b.Ospfv3) {
				return false
			}
			if !matchVrfEcmp(a.Ecmp, b.Ecmp) {
				return false
			}
			if !matchVrfMulticast(a.Multicast, b.Multicast) {
				return false
			}
			if !matchVrfRip(a.Rip, b.Rip) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
