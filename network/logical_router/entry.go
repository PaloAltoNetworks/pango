package logical_router

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"network", "logical-router", "$name"}
)

type Entry struct {
	Name string
	Vrf  []Vrf
	Misc []generic.Xml
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
	Misc         []generic.Xml
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
	Misc        []generic.Xml
}
type VrfRibFilter struct {
	Ipv4 *VrfRibFilterIpv4
	Ipv6 *VrfRibFilterIpv6
	Misc []generic.Xml
}
type VrfRibFilterIpv4 struct {
	Static *VrfRibFilterIpv4Static
	Bgp    *VrfRibFilterIpv4Bgp
	Ospf   *VrfRibFilterIpv4Ospf
	Rip    *VrfRibFilterIpv4Rip
	Misc   []generic.Xml
}
type VrfRibFilterIpv4Static struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv4Bgp struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv4Ospf struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv4Rip struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv6 struct {
	Static *VrfRibFilterIpv6Static
	Bgp    *VrfRibFilterIpv6Bgp
	Ospfv3 *VrfRibFilterIpv6Ospfv3
	Misc   []generic.Xml
}
type VrfRibFilterIpv6Static struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv6Bgp struct {
	RouteMap *string
	Misc     []generic.Xml
}
type VrfRibFilterIpv6Ospfv3 struct {
	RouteMap *string
	Misc     []generic.Xml
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
	Misc                        []generic.Xml
}
type VrfBgpMed struct {
	AlwaysCompareMed           *bool
	DeterministicMedComparison *bool
	Misc                       []generic.Xml
}
type VrfBgpGracefulRestart struct {
	Enable             *bool
	StaleRouteTime     *int64
	MaxPeerRestartTime *int64
	LocalRestartTime   *int64
	Misc               []generic.Xml
}
type VrfBgpGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfBgpRedistributionProfile struct {
	Ipv4 *VrfBgpRedistributionProfileIpv4
	Ipv6 *VrfBgpRedistributionProfileIpv6
	Misc []generic.Xml
}
type VrfBgpRedistributionProfileIpv4 struct {
	Unicast *string
	Misc    []generic.Xml
}
type VrfBgpRedistributionProfileIpv6 struct {
	Unicast *string
	Misc    []generic.Xml
}
type VrfBgpAdvertiseNetwork struct {
	Ipv4 *VrfBgpAdvertiseNetworkIpv4
	Ipv6 *VrfBgpAdvertiseNetworkIpv6
	Misc []generic.Xml
}
type VrfBgpAdvertiseNetworkIpv4 struct {
	Network []VrfBgpAdvertiseNetworkIpv4Network
	Misc    []generic.Xml
}
type VrfBgpAdvertiseNetworkIpv4Network struct {
	Name      string
	Unicast   *bool
	Multicast *bool
	Backdoor  *bool
	Misc      []generic.Xml
}
type VrfBgpAdvertiseNetworkIpv6 struct {
	Network []VrfBgpAdvertiseNetworkIpv6Network
	Misc    []generic.Xml
}
type VrfBgpAdvertiseNetworkIpv6Network struct {
	Name    string
	Unicast *bool
	Misc    []generic.Xml
}
type VrfBgpPeerGroup struct {
	Name              string
	Enable            *bool
	Type              *VrfBgpPeerGroupType
	AddressFamily     *VrfBgpPeerGroupAddressFamily
	FilteringProfile  *VrfBgpPeerGroupFilteringProfile
	ConnectionOptions *VrfBgpPeerGroupConnectionOptions
	Peer              []VrfBgpPeerGroupPeer
	Misc              []generic.Xml
}
type VrfBgpPeerGroupType struct {
	Ibgp *VrfBgpPeerGroupTypeIbgp
	Ebgp *VrfBgpPeerGroupTypeEbgp
	Misc []generic.Xml
}
type VrfBgpPeerGroupTypeIbgp struct {
	Misc []generic.Xml
}
type VrfBgpPeerGroupTypeEbgp struct {
	Misc []generic.Xml
}
type VrfBgpPeerGroupAddressFamily struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type VrfBgpPeerGroupFilteringProfile struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type VrfBgpPeerGroupConnectionOptions struct {
	Timers         *string
	Multihop       *int64
	Authentication *string
	Dampening      *string
	Misc           []generic.Xml
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
	Misc                          []generic.Xml
}
type VrfBgpPeerGroupPeerInherit struct {
	Yes  *VrfBgpPeerGroupPeerInheritYes
	No   *VrfBgpPeerGroupPeerInheritNo
	Misc []generic.Xml
}
type VrfBgpPeerGroupPeerInheritYes struct {
	Misc []generic.Xml
}
type VrfBgpPeerGroupPeerInheritNo struct {
	AddressFamily    *VrfBgpPeerGroupPeerInheritNoAddressFamily
	FilteringProfile *VrfBgpPeerGroupPeerInheritNoFilteringProfile
	Misc             []generic.Xml
}
type VrfBgpPeerGroupPeerInheritNoAddressFamily struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type VrfBgpPeerGroupPeerInheritNoFilteringProfile struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type VrfBgpPeerGroupPeerLocalAddress struct {
	Interface *string
	Ip        *string
	Misc      []generic.Xml
}
type VrfBgpPeerGroupPeerPeerAddress struct {
	Ip   *string
	Fqdn *string
	Misc []generic.Xml
}
type VrfBgpPeerGroupPeerConnectionOptions struct {
	Timers         *string
	Multihop       *string
	Authentication *string
	Dampening      *string
	Misc           []generic.Xml
}
type VrfBgpPeerGroupPeerBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfBgpAggregateRoutes struct {
	Name        string
	Description *string
	Enable      *bool
	SummaryOnly *bool
	AsSet       *bool
	SameMed     *bool
	Type        *VrfBgpAggregateRoutesType
	Misc        []generic.Xml
}
type VrfBgpAggregateRoutesType struct {
	Ipv4 *VrfBgpAggregateRoutesTypeIpv4
	Ipv6 *VrfBgpAggregateRoutesTypeIpv6
	Misc []generic.Xml
}
type VrfBgpAggregateRoutesTypeIpv4 struct {
	SummaryPrefix *string
	SuppressMap   *string
	AttributeMap  *string
	Misc          []generic.Xml
}
type VrfBgpAggregateRoutesTypeIpv6 struct {
	SummaryPrefix *string
	SuppressMap   *string
	AttributeMap  *string
	Misc          []generic.Xml
}
type VrfRoutingTable struct {
	Ip   *VrfRoutingTableIp
	Ipv6 *VrfRoutingTableIpv6
	Misc []generic.Xml
}
type VrfRoutingTableIp struct {
	StaticRoute []VrfRoutingTableIpStaticRoute
	Misc        []generic.Xml
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
	Misc        []generic.Xml
}
type VrfRoutingTableIpStaticRouteNexthop struct {
	Discard   *VrfRoutingTableIpStaticRouteNexthopDiscard
	IpAddress *string
	NextLr    *string
	Fqdn      *string
	Misc      []generic.Xml
}
type VrfRoutingTableIpStaticRouteNexthopDiscard struct {
	Misc []generic.Xml
}
type VrfRoutingTableIpStaticRouteBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfRoutingTableIpStaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations
	Misc                []generic.Xml
}
type VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations struct {
	Name        string
	Enable      *bool
	Source      *string
	Destination *string
	Interval    *int64
	Count       *int64
	Misc        []generic.Xml
}
type VrfRoutingTableIpv6 struct {
	StaticRoute []VrfRoutingTableIpv6StaticRoute
	Misc        []generic.Xml
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
	Misc        []generic.Xml
}
type VrfRoutingTableIpv6StaticRouteNexthop struct {
	Discard     *VrfRoutingTableIpv6StaticRouteNexthopDiscard
	Ipv6Address *string
	Fqdn        *string
	NextLr      *string
	Misc        []generic.Xml
}
type VrfRoutingTableIpv6StaticRouteNexthopDiscard struct {
	Misc []generic.Xml
}
type VrfRoutingTableIpv6StaticRouteBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfRoutingTableIpv6StaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations
	Misc                []generic.Xml
}
type VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations struct {
	Name        string
	Enable      *bool
	Source      *string
	Destination *string
	Interval    *int64
	Count       *int64
	Misc        []generic.Xml
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
	Misc                  []generic.Xml
}
type VrfOspfGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfOspfGracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	StrictLSAChecking      *bool
	MaxNeighborRestartTime *int64
	Misc                   []generic.Xml
}
type VrfOspfArea struct {
	Name           string
	Authentication *string
	Type           *VrfOspfAreaType
	Range          []VrfOspfAreaRange
	Interface      []VrfOspfAreaInterface
	VirtualLink    []VrfOspfAreaVirtualLink
	Misc           []generic.Xml
}
type VrfOspfAreaType struct {
	Normal *VrfOspfAreaTypeNormal
	Stub   *VrfOspfAreaTypeStub
	Nssa   *VrfOspfAreaTypeNssa
	Misc   []generic.Xml
}
type VrfOspfAreaTypeNormal struct {
	Abr  *VrfOspfAreaTypeNormalAbr
	Misc []generic.Xml
}
type VrfOspfAreaTypeNormalAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	Misc               []generic.Xml
}
type VrfOspfAreaTypeStub struct {
	NoSummary          *bool
	Abr                *VrfOspfAreaTypeStubAbr
	DefaultRouteMetric *int64
	Misc               []generic.Xml
}
type VrfOspfAreaTypeStubAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	Misc               []generic.Xml
}
type VrfOspfAreaTypeNssa struct {
	NoSummary                   *bool
	DefaultInformationOriginate *VrfOspfAreaTypeNssaDefaultInformationOriginate
	Abr                         *VrfOspfAreaTypeNssaAbr
	Misc                        []generic.Xml
}
type VrfOspfAreaTypeNssaDefaultInformationOriginate struct {
	Metric     *int64
	MetricType *string
	Misc       []generic.Xml
}
type VrfOspfAreaTypeNssaAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	NssaExtRange       []VrfOspfAreaTypeNssaAbrNssaExtRange
	Misc               []generic.Xml
}
type VrfOspfAreaTypeNssaAbrNssaExtRange struct {
	Name      string
	Advertise *bool
	Misc      []generic.Xml
}
type VrfOspfAreaRange struct {
	Name      string
	Advertise *bool
	Misc      []generic.Xml
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
	Misc           []generic.Xml
}
type VrfOspfAreaInterfaceLinkType struct {
	Broadcast *VrfOspfAreaInterfaceLinkTypeBroadcast
	P2p       *VrfOspfAreaInterfaceLinkTypeP2p
	P2mp      *VrfOspfAreaInterfaceLinkTypeP2mp
	Misc      []generic.Xml
}
type VrfOspfAreaInterfaceLinkTypeBroadcast struct {
	Misc []generic.Xml
}
type VrfOspfAreaInterfaceLinkTypeP2p struct {
	Misc []generic.Xml
}
type VrfOspfAreaInterfaceLinkTypeP2mp struct {
	Neighbor []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor
	Misc     []generic.Xml
}
type VrfOspfAreaInterfaceLinkTypeP2mpNeighbor struct {
	Name     string
	Priority *int64
	Misc     []generic.Xml
}
type VrfOspfAreaInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
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
	Misc           []generic.Xml
}
type VrfOspfAreaVirtualLinkBfd struct {
	Profile *string
	Misc    []generic.Xml
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
	Misc                  []generic.Xml
}
type VrfOspfv3GlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfOspfv3GracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	StrictLSAChecking      *bool
	MaxNeighborRestartTime *int64
	Misc                   []generic.Xml
}
type VrfOspfv3Area struct {
	Name           string
	Authentication *string
	Type           *VrfOspfv3AreaType
	Range          []VrfOspfv3AreaRange
	Interface      []VrfOspfv3AreaInterface
	VirtualLink    []VrfOspfv3AreaVirtualLink
	Misc           []generic.Xml
}
type VrfOspfv3AreaType struct {
	Normal *VrfOspfv3AreaTypeNormal
	Stub   *VrfOspfv3AreaTypeStub
	Nssa   *VrfOspfv3AreaTypeNssa
	Misc   []generic.Xml
}
type VrfOspfv3AreaTypeNormal struct {
	Abr  *VrfOspfv3AreaTypeNormalAbr
	Misc []generic.Xml
}
type VrfOspfv3AreaTypeNormalAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	Misc               []generic.Xml
}
type VrfOspfv3AreaTypeStub struct {
	NoSummary          *bool
	Abr                *VrfOspfv3AreaTypeStubAbr
	DefaultRouteMetric *int64
	Misc               []generic.Xml
}
type VrfOspfv3AreaTypeStubAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	Misc               []generic.Xml
}
type VrfOspfv3AreaTypeNssa struct {
	NoSummary                   *bool
	DefaultInformationOriginate *VrfOspfv3AreaTypeNssaDefaultInformationOriginate
	Abr                         *VrfOspfv3AreaTypeNssaAbr
	Misc                        []generic.Xml
}
type VrfOspfv3AreaTypeNssaDefaultInformationOriginate struct {
	Metric     *int64
	MetricType *string
	Misc       []generic.Xml
}
type VrfOspfv3AreaTypeNssaAbr struct {
	ImportList         *string
	ExportList         *string
	InboundFilterList  *string
	OutboundFilterList *string
	NssaExtRange       []VrfOspfv3AreaTypeNssaAbrNssaExtRange
	Misc               []generic.Xml
}
type VrfOspfv3AreaTypeNssaAbrNssaExtRange struct {
	Name      string
	Advertise *bool
	Misc      []generic.Xml
}
type VrfOspfv3AreaRange struct {
	Name      string
	Advertise *bool
	Misc      []generic.Xml
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
	Misc           []generic.Xml
}
type VrfOspfv3AreaInterfaceLinkType struct {
	Broadcast *VrfOspfv3AreaInterfaceLinkTypeBroadcast
	P2p       *VrfOspfv3AreaInterfaceLinkTypeP2p
	P2mp      *VrfOspfv3AreaInterfaceLinkTypeP2mp
	Misc      []generic.Xml
}
type VrfOspfv3AreaInterfaceLinkTypeBroadcast struct {
	Misc []generic.Xml
}
type VrfOspfv3AreaInterfaceLinkTypeP2p struct {
	Misc []generic.Xml
}
type VrfOspfv3AreaInterfaceLinkTypeP2mp struct {
	Neighbor []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor
	Misc     []generic.Xml
}
type VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor struct {
	Name     string
	Priority *int64
	Misc     []generic.Xml
}
type VrfOspfv3AreaInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfOspfv3AreaVirtualLink struct {
	Name           string
	NeighborId     *string
	TransitAreaId  *string
	Enable         *bool
	InstanceId     *int64
	Timing         *string
	Authentication *string
	Misc           []generic.Xml
}
type VrfEcmp struct {
	Enable           *bool
	MaxPath          *int64
	SymmetricReturn  *bool
	StrictSourcePath *bool
	Algorithm        *VrfEcmpAlgorithm
	Misc             []generic.Xml
}
type VrfEcmpAlgorithm struct {
	IpModulo           *VrfEcmpAlgorithmIpModulo
	IpHash             *VrfEcmpAlgorithmIpHash
	WeightedRoundRobin *VrfEcmpAlgorithmWeightedRoundRobin
	BalancedRoundRobin *VrfEcmpAlgorithmBalancedRoundRobin
	Misc               []generic.Xml
}
type VrfEcmpAlgorithmIpModulo struct {
	Misc []generic.Xml
}
type VrfEcmpAlgorithmIpHash struct {
	SrcOnly  *bool
	UsePort  *bool
	HashSeed *int64
	Misc     []generic.Xml
}
type VrfEcmpAlgorithmWeightedRoundRobin struct {
	Interface []VrfEcmpAlgorithmWeightedRoundRobinInterface
	Misc      []generic.Xml
}
type VrfEcmpAlgorithmWeightedRoundRobinInterface struct {
	Name   string
	Weight *int64
	Misc   []generic.Xml
}
type VrfEcmpAlgorithmBalancedRoundRobin struct {
	Misc []generic.Xml
}
type VrfMulticast struct {
	Enable      *bool
	StaticRoute []VrfMulticastStaticRoute
	Pim         *VrfMulticastPim
	Igmp        *VrfMulticastIgmp
	Msdp        *VrfMulticastMsdp
	Misc        []generic.Xml
}
type VrfMulticastStaticRoute struct {
	Name        string
	Destination *string
	Interface   *string
	Preference  *int64
	Nexthop     *VrfMulticastStaticRouteNexthop
	Misc        []generic.Xml
}
type VrfMulticastStaticRouteNexthop struct {
	IpAddress *string
	Misc      []generic.Xml
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
	Misc            []generic.Xml
}
type VrfMulticastPimSsmAddressSpace struct {
	GroupList *string
	Misc      []generic.Xml
}
type VrfMulticastPimRp struct {
	LocalRp    *VrfMulticastPimRpLocalRp
	ExternalRp []VrfMulticastPimRpExternalRp
	Misc       []generic.Xml
}
type VrfMulticastPimRpLocalRp struct {
	StaticRp    *VrfMulticastPimRpLocalRpStaticRp
	CandidateRp *VrfMulticastPimRpLocalRpCandidateRp
	Misc        []generic.Xml
}
type VrfMulticastPimRpLocalRpStaticRp struct {
	Interface *string
	Address   *string
	Override  *bool
	GroupList *string
	Misc      []generic.Xml
}
type VrfMulticastPimRpLocalRpCandidateRp struct {
	Interface             *string
	Address               *string
	Priority              *int64
	AdvertisementInterval *int64
	GroupList             *string
	Misc                  []generic.Xml
}
type VrfMulticastPimRpExternalRp struct {
	Name      string
	GroupList *string
	Override  *bool
	Misc      []generic.Xml
}
type VrfMulticastPimSptThreshold struct {
	Name      string
	Threshold *string
	Misc      []generic.Xml
}
type VrfMulticastPimInterface struct {
	Name           string
	Description    *string
	DrPriority     *int64
	SendBsm        *bool
	IfTimer        *string
	NeighborFilter *string
	Misc           []generic.Xml
}
type VrfMulticastIgmp struct {
	Enable  *bool
	Dynamic *VrfMulticastIgmpDynamic
	Static  []VrfMulticastIgmpStatic
	Misc    []generic.Xml
}
type VrfMulticastIgmpDynamic struct {
	Interface []VrfMulticastIgmpDynamicInterface
	Misc      []generic.Xml
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
	Misc                []generic.Xml
}
type VrfMulticastIgmpStatic struct {
	Name          string
	Interface     *string
	GroupAddress  *string
	SourceAddress *string
	Misc          []generic.Xml
}
type VrfMulticastMsdp struct {
	Enable               *bool
	GlobalTimer          *string
	GlobalAuthentication *string
	OriginatorId         *VrfMulticastMsdpOriginatorId
	Peer                 []VrfMulticastMsdpPeer
	Misc                 []generic.Xml
}
type VrfMulticastMsdpOriginatorId struct {
	Interface *string
	Ip        *string
	Misc      []generic.Xml
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
	Misc             []generic.Xml
}
type VrfMulticastMsdpPeerLocalAddress struct {
	Interface *string
	Ip        *string
	Misc      []generic.Xml
}
type VrfMulticastMsdpPeerPeerAddress struct {
	Ip   *string
	Fqdn *string
	Misc []generic.Xml
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
	Misc                         []generic.Xml
}
type VrfRipGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfRipGlobalInboundDistributeList struct {
	AccessList *string
	Misc       []generic.Xml
}
type VrfRipGlobalOutboundDistributeList struct {
	AccessList *string
	Misc       []generic.Xml
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
	Misc                            []generic.Xml
}
type VrfRipInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type VrfRipInterfaceInterfaceInboundDistributeList struct {
	AccessList *string
	Metric     *int64
	Misc       []generic.Xml
}
type VrfRipInterfaceInterfaceOutboundDistributeList struct {
	AccessList *string
	Metric     *int64
	Misc       []generic.Xml
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
	XMLName xml.Name         `xml:"entry"`
	Name    string           `xml:"name,attr"`
	Vrf     *vrfContainerXml `xml:"vrf,omitempty"`
	Misc    []generic.Xml    `xml:",any"`
}
type vrfContainerXml struct {
	Entries []vrfXml `xml:"entry"`
}
type vrfXml struct {
	XMLName      xml.Name            `xml:"entry"`
	Name         string              `xml:"name,attr"`
	Interface    *util.MemberType    `xml:"interface,omitempty"`
	AdminDists   *vrfAdminDistsXml   `xml:"admin-dists,omitempty"`
	RibFilter    *vrfRibFilterXml    `xml:"rib-filter,omitempty"`
	Bgp          *vrfBgpXml          `xml:"bgp,omitempty"`
	RoutingTable *vrfRoutingTableXml `xml:"routing-table,omitempty"`
	Ospf         *vrfOspfXml         `xml:"ospf,omitempty"`
	Ospfv3       *vrfOspfv3Xml       `xml:"ospfv3,omitempty"`
	Ecmp         *vrfEcmpXml         `xml:"ecmp,omitempty"`
	Multicast    *vrfMulticastXml    `xml:"multicast,omitempty"`
	Rip          *vrfRipXml          `xml:"rip,omitempty"`
	Misc         []generic.Xml       `xml:",any"`
}
type vrfAdminDistsXml struct {
	Static      *int64        `xml:"static,omitempty"`
	StaticIpv6  *int64        `xml:"static-ipv6,omitempty"`
	OspfInter   *int64        `xml:"ospf-inter,omitempty"`
	OspfIntra   *int64        `xml:"ospf-intra,omitempty"`
	OspfExt     *int64        `xml:"ospf-ext,omitempty"`
	Ospfv3Inter *int64        `xml:"ospfv3-inter,omitempty"`
	Ospfv3Intra *int64        `xml:"ospfv3-intra,omitempty"`
	Ospfv3Ext   *int64        `xml:"ospfv3-ext,omitempty"`
	BgpInternal *int64        `xml:"bgp-internal,omitempty"`
	BgpExternal *int64        `xml:"bgp-external,omitempty"`
	BgpLocal    *int64        `xml:"bgp-local,omitempty"`
	Rip         *int64        `xml:"rip,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfRibFilterXml struct {
	Ipv4 *vrfRibFilterIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *vrfRibFilterIpv6Xml `xml:"ipv6,omitempty"`
	Misc []generic.Xml        `xml:",any"`
}
type vrfRibFilterIpv4Xml struct {
	Static *vrfRibFilterIpv4StaticXml `xml:"static,omitempty"`
	Bgp    *vrfRibFilterIpv4BgpXml    `xml:"bgp,omitempty"`
	Ospf   *vrfRibFilterIpv4OspfXml   `xml:"ospf,omitempty"`
	Rip    *vrfRibFilterIpv4RipXml    `xml:"rip,omitempty"`
	Misc   []generic.Xml              `xml:",any"`
}
type vrfRibFilterIpv4StaticXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4BgpXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4OspfXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4RipXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6Xml struct {
	Static *vrfRibFilterIpv6StaticXml `xml:"static,omitempty"`
	Bgp    *vrfRibFilterIpv6BgpXml    `xml:"bgp,omitempty"`
	Ospfv3 *vrfRibFilterIpv6Ospfv3Xml `xml:"ospfv3,omitempty"`
	Misc   []generic.Xml              `xml:",any"`
}
type vrfRibFilterIpv6StaticXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6BgpXml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6Ospfv3Xml struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfBgpXml struct {
	Enable                      *string                            `xml:"enable,omitempty"`
	RouterId                    *string                            `xml:"router-id,omitempty"`
	LocalAs                     *string                            `xml:"local-as,omitempty"`
	InstallRoute                *string                            `xml:"install-route,omitempty"`
	EnforceFirstAs              *string                            `xml:"enforce-first-as,omitempty"`
	FastExternalFailover        *string                            `xml:"fast-external-failover,omitempty"`
	EcmpMultiAs                 *string                            `xml:"ecmp-multi-as,omitempty"`
	DefaultLocalPreference      *int64                             `xml:"default-local-preference,omitempty"`
	GracefulShutdown            *string                            `xml:"graceful-shutdown,omitempty"`
	AlwaysAdvertiseNetworkRoute *string                            `xml:"always-advertise-network-route,omitempty"`
	Med                         *vrfBgpMedXml                      `xml:"med,omitempty"`
	GracefulRestart             *vrfBgpGracefulRestartXml          `xml:"graceful-restart,omitempty"`
	GlobalBfd                   *vrfBgpGlobalBfdXml                `xml:"global-bfd,omitempty"`
	RedistributionProfile       *vrfBgpRedistributionProfileXml    `xml:"redistribution-profile,omitempty"`
	AdvertiseNetwork            *vrfBgpAdvertiseNetworkXml         `xml:"advertise-network,omitempty"`
	PeerGroup                   *vrfBgpPeerGroupContainerXml       `xml:"peer-group,omitempty"`
	AggregateRoutes             *vrfBgpAggregateRoutesContainerXml `xml:"aggregate-routes,omitempty"`
	Misc                        []generic.Xml                      `xml:",any"`
}
type vrfBgpMedXml struct {
	AlwaysCompareMed           *string       `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string       `xml:"deterministic-med-comparison,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type vrfBgpGracefulRestartXml struct {
	Enable             *string       `xml:"enable,omitempty"`
	StaleRouteTime     *int64        `xml:"stale-route-time,omitempty"`
	MaxPeerRestartTime *int64        `xml:"max-peer-restart-time,omitempty"`
	LocalRestartTime   *int64        `xml:"local-restart-time,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfBgpGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpRedistributionProfileXml struct {
	Ipv4 *vrfBgpRedistributionProfileIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpRedistributionProfileIpv6Xml `xml:"ipv6,omitempty"`
	Misc []generic.Xml                       `xml:",any"`
}
type vrfBgpRedistributionProfileIpv4Xml struct {
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpRedistributionProfileIpv6Xml struct {
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpAdvertiseNetworkXml struct {
	Ipv4 *vrfBgpAdvertiseNetworkIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpAdvertiseNetworkIpv6Xml `xml:"ipv6,omitempty"`
	Misc []generic.Xml                  `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv4Xml struct {
	Network *vrfBgpAdvertiseNetworkIpv4NetworkContainerXml `xml:"network,omitempty"`
	Misc    []generic.Xml                                  `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv4NetworkContainerXml struct {
	Entries []vrfBgpAdvertiseNetworkIpv4NetworkXml `xml:"entry"`
}
type vrfBgpAdvertiseNetworkIpv4NetworkXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Unicast   *string       `xml:"unicast,omitempty"`
	Multicast *string       `xml:"multicast,omitempty"`
	Backdoor  *string       `xml:"backdoor,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv6Xml struct {
	Network *vrfBgpAdvertiseNetworkIpv6NetworkContainerXml `xml:"network,omitempty"`
	Misc    []generic.Xml                                  `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv6NetworkContainerXml struct {
	Entries []vrfBgpAdvertiseNetworkIpv6NetworkXml `xml:"entry"`
}
type vrfBgpAdvertiseNetworkIpv6NetworkXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupContainerXml struct {
	Entries []vrfBgpPeerGroupXml `xml:"entry"`
}
type vrfBgpPeerGroupXml struct {
	XMLName           xml.Name                             `xml:"entry"`
	Name              string                               `xml:"name,attr"`
	Enable            *string                              `xml:"enable,omitempty"`
	Type              *vrfBgpPeerGroupTypeXml              `xml:"type,omitempty"`
	AddressFamily     *vrfBgpPeerGroupAddressFamilyXml     `xml:"address-family,omitempty"`
	FilteringProfile  *vrfBgpPeerGroupFilteringProfileXml  `xml:"filtering-profile,omitempty"`
	ConnectionOptions *vrfBgpPeerGroupConnectionOptionsXml `xml:"connection-options,omitempty"`
	Peer              *vrfBgpPeerGroupPeerContainerXml     `xml:"peer,omitempty"`
	Misc              []generic.Xml                        `xml:",any"`
}
type vrfBgpPeerGroupTypeXml struct {
	Ibgp *vrfBgpPeerGroupTypeIbgpXml `xml:"ibgp,omitempty"`
	Ebgp *vrfBgpPeerGroupTypeEbgpXml `xml:"ebgp,omitempty"`
	Misc []generic.Xml               `xml:",any"`
}
type vrfBgpPeerGroupTypeIbgpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupTypeEbgpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupAddressFamilyXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupFilteringProfileXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupConnectionOptionsXml struct {
	Timers         *string       `xml:"timers,omitempty"`
	Multihop       *int64        `xml:"multihop,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Dampening      *string       `xml:"dampening,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerContainerXml struct {
	Entries []vrfBgpPeerGroupPeerXml `xml:"entry"`
}
type vrfBgpPeerGroupPeerXml struct {
	XMLName                       xml.Name                                 `xml:"entry"`
	Name                          string                                   `xml:"name,attr"`
	Enable                        *string                                  `xml:"enable,omitempty"`
	Passive                       *string                                  `xml:"passive,omitempty"`
	PeerAs                        *string                                  `xml:"peer-as,omitempty"`
	EnableSenderSideLoopDetection *string                                  `xml:"enable-sender-side-loop-detection,omitempty"`
	Inherit                       *vrfBgpPeerGroupPeerInheritXml           `xml:"inherit,omitempty"`
	LocalAddress                  *vrfBgpPeerGroupPeerLocalAddressXml      `xml:"local-address,omitempty"`
	PeerAddress                   *vrfBgpPeerGroupPeerPeerAddressXml       `xml:"peer-address,omitempty"`
	ConnectionOptions             *vrfBgpPeerGroupPeerConnectionOptionsXml `xml:"connection-options,omitempty"`
	Bfd                           *vrfBgpPeerGroupPeerBfdXml               `xml:"bfd,omitempty"`
	Misc                          []generic.Xml                            `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritXml struct {
	Yes  *vrfBgpPeerGroupPeerInheritYesXml `xml:"yes,omitempty"`
	No   *vrfBgpPeerGroupPeerInheritNoXml  `xml:"no,omitempty"`
	Misc []generic.Xml                     `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritYesXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoXml struct {
	AddressFamily    *vrfBgpPeerGroupPeerInheritNoAddressFamilyXml    `xml:"address-family,omitempty"`
	FilteringProfile *vrfBgpPeerGroupPeerInheritNoFilteringProfileXml `xml:"filtering-profile,omitempty"`
	Misc             []generic.Xml                                    `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoAddressFamilyXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoFilteringProfileXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerLocalAddressXml struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerPeerAddressXml struct {
	Ip   *string       `xml:"ip,omitempty"`
	Fqdn *string       `xml:"fqdn,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerConnectionOptionsXml struct {
	Timers         *string       `xml:"timers,omitempty"`
	Multihop       *string       `xml:"multihop,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Dampening      *string       `xml:"dampening,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpAggregateRoutesContainerXml struct {
	Entries []vrfBgpAggregateRoutesXml `xml:"entry"`
}
type vrfBgpAggregateRoutesXml struct {
	XMLName     xml.Name                      `xml:"entry"`
	Name        string                        `xml:"name,attr"`
	Description *string                       `xml:"description,omitempty"`
	Enable      *string                       `xml:"enable,omitempty"`
	SummaryOnly *string                       `xml:"summary-only,omitempty"`
	AsSet       *string                       `xml:"as-set,omitempty"`
	SameMed     *string                       `xml:"same-med,omitempty"`
	Type        *vrfBgpAggregateRoutesTypeXml `xml:"type,omitempty"`
	Misc        []generic.Xml                 `xml:",any"`
}
type vrfBgpAggregateRoutesTypeXml struct {
	Ipv4 *vrfBgpAggregateRoutesTypeIpv4Xml `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpAggregateRoutesTypeIpv6Xml `xml:"ipv6,omitempty"`
	Misc []generic.Xml                     `xml:",any"`
}
type vrfBgpAggregateRoutesTypeIpv4Xml struct {
	SummaryPrefix *string       `xml:"summary-prefix,omitempty"`
	SuppressMap   *string       `xml:"suppress-map,omitempty"`
	AttributeMap  *string       `xml:"attribute-map,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfBgpAggregateRoutesTypeIpv6Xml struct {
	SummaryPrefix *string       `xml:"summary-prefix,omitempty"`
	SuppressMap   *string       `xml:"suppress-map,omitempty"`
	AttributeMap  *string       `xml:"attribute-map,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfRoutingTableXml struct {
	Ip   *vrfRoutingTableIpXml   `xml:"ip,omitempty"`
	Ipv6 *vrfRoutingTableIpv6Xml `xml:"ipv6,omitempty"`
	Misc []generic.Xml           `xml:",any"`
}
type vrfRoutingTableIpXml struct {
	StaticRoute *vrfRoutingTableIpStaticRouteContainerXml `xml:"static-route,omitempty"`
	Misc        []generic.Xml                             `xml:",any"`
}
type vrfRoutingTableIpStaticRouteContainerXml struct {
	Entries []vrfRoutingTableIpStaticRouteXml `xml:"entry"`
}
type vrfRoutingTableIpStaticRouteXml struct {
	XMLName     xml.Name                                    `xml:"entry"`
	Name        string                                      `xml:"name,attr"`
	Destination *string                                     `xml:"destination,omitempty"`
	Interface   *string                                     `xml:"interface,omitempty"`
	AdminDist   *int64                                      `xml:"admin-dist,omitempty"`
	Metric      *int64                                      `xml:"metric,omitempty"`
	Nexthop     *vrfRoutingTableIpStaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	Bfd         *vrfRoutingTableIpStaticRouteBfdXml         `xml:"bfd,omitempty"`
	PathMonitor *vrfRoutingTableIpStaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`
	Misc        []generic.Xml                               `xml:",any"`
}
type vrfRoutingTableIpStaticRouteNexthopXml struct {
	Discard   *vrfRoutingTableIpStaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	IpAddress *string                                        `xml:"ip-address,omitempty"`
	NextLr    *string                                        `xml:"next-lr,omitempty"`
	Fqdn      *string                                        `xml:"fqdn,omitempty"`
	Misc      []generic.Xml                                  `xml:",any"`
}
type vrfRoutingTableIpStaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpStaticRouteBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpStaticRoutePathMonitorXml struct {
	Enable              *string                                                                 `xml:"enable,omitempty"`
	FailureCondition    *string                                                                 `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                  `xml:"hold-time,omitempty"`
	MonitorDestinations *vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml `xml:"monitor-destinations,omitempty"`
	Misc                []generic.Xml                                                           `xml:",any"`
}
type vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml struct {
	Entries []vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml `xml:"entry"`
}
type vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Enable      *string       `xml:"enable,omitempty"`
	Source      *string       `xml:"source,omitempty"`
	Destination *string       `xml:"destination,omitempty"`
	Interval    *int64        `xml:"interval,omitempty"`
	Count       *int64        `xml:"count,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6Xml struct {
	StaticRoute *vrfRoutingTableIpv6StaticRouteContainerXml `xml:"static-route,omitempty"`
	Misc        []generic.Xml                               `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteContainerXml struct {
	Entries []vrfRoutingTableIpv6StaticRouteXml `xml:"entry"`
}
type vrfRoutingTableIpv6StaticRouteXml struct {
	XMLName     xml.Name                                      `xml:"entry"`
	Name        string                                        `xml:"name,attr"`
	Destination *string                                       `xml:"destination,omitempty"`
	Interface   *string                                       `xml:"interface,omitempty"`
	AdminDist   *int64                                        `xml:"admin-dist,omitempty"`
	Metric      *int64                                        `xml:"metric,omitempty"`
	Nexthop     *vrfRoutingTableIpv6StaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	Bfd         *vrfRoutingTableIpv6StaticRouteBfdXml         `xml:"bfd,omitempty"`
	PathMonitor *vrfRoutingTableIpv6StaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`
	Misc        []generic.Xml                                 `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteNexthopXml struct {
	Discard     *vrfRoutingTableIpv6StaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	Ipv6Address *string                                          `xml:"ipv6-address,omitempty"`
	Fqdn        *string                                          `xml:"fqdn,omitempty"`
	NextLr      *string                                          `xml:"next-lr,omitempty"`
	Misc        []generic.Xml                                    `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorXml struct {
	Enable              *string                                                                   `xml:"enable,omitempty"`
	FailureCondition    *string                                                                   `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                    `xml:"hold-time,omitempty"`
	MonitorDestinations *vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml `xml:"monitor-destinations,omitempty"`
	Misc                []generic.Xml                                                             `xml:",any"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml struct {
	Entries []vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml `xml:"entry"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Enable      *string       `xml:"enable,omitempty"`
	Source      *string       `xml:"source,omitempty"`
	Destination *string       `xml:"destination,omitempty"`
	Interval    *int64        `xml:"interval,omitempty"`
	Count       *int64        `xml:"count,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfOspfXml struct {
	RouterId              *string                    `xml:"router-id,omitempty"`
	Enable                *string                    `xml:"enable,omitempty"`
	Rfc1583               *string                    `xml:"rfc1583,omitempty"`
	SpfTimer              *string                    `xml:"spf-timer,omitempty"`
	GlobalIfTimer         *string                    `xml:"global-if-timer,omitempty"`
	RedistributionProfile *string                    `xml:"redistribution-profile,omitempty"`
	GlobalBfd             *vrfOspfGlobalBfdXml       `xml:"global-bfd,omitempty"`
	GracefulRestart       *vrfOspfGracefulRestartXml `xml:"graceful-restart,omitempty"`
	Area                  *vrfOspfAreaContainerXml   `xml:"area,omitempty"`
	Misc                  []generic.Xml              `xml:",any"`
}
type vrfOspfGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfGracefulRestartXml struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type vrfOspfAreaContainerXml struct {
	Entries []vrfOspfAreaXml `xml:"entry"`
}
type vrfOspfAreaXml struct {
	XMLName        xml.Name                            `xml:"entry"`
	Name           string                              `xml:"name,attr"`
	Authentication *string                             `xml:"authentication,omitempty"`
	Type           *vrfOspfAreaTypeXml                 `xml:"type,omitempty"`
	Range          *vrfOspfAreaRangeContainerXml       `xml:"range,omitempty"`
	Interface      *vrfOspfAreaInterfaceContainerXml   `xml:"interface,omitempty"`
	VirtualLink    *vrfOspfAreaVirtualLinkContainerXml `xml:"virtual-link,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
}
type vrfOspfAreaTypeXml struct {
	Normal *vrfOspfAreaTypeNormalXml `xml:"normal,omitempty"`
	Stub   *vrfOspfAreaTypeStubXml   `xml:"stub,omitempty"`
	Nssa   *vrfOspfAreaTypeNssaXml   `xml:"nssa,omitempty"`
	Misc   []generic.Xml             `xml:",any"`
}
type vrfOspfAreaTypeNormalXml struct {
	Abr  *vrfOspfAreaTypeNormalAbrXml `xml:"abr,omitempty"`
	Misc []generic.Xml                `xml:",any"`
}
type vrfOspfAreaTypeNormalAbrXml struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeStubXml struct {
	NoSummary          *string                    `xml:"no-summary,omitempty"`
	Abr                *vrfOspfAreaTypeStubAbrXml `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                     `xml:"default-route-metric,omitempty"`
	Misc               []generic.Xml              `xml:",any"`
}
type vrfOspfAreaTypeStubAbrXml struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeNssaXml struct {
	NoSummary                   *string                                            `xml:"no-summary,omitempty"`
	DefaultInformationOriginate *vrfOspfAreaTypeNssaDefaultInformationOriginateXml `xml:"default-information-originate,omitempty"`
	Abr                         *vrfOspfAreaTypeNssaAbrXml                         `xml:"abr,omitempty"`
	Misc                        []generic.Xml                                      `xml:",any"`
}
type vrfOspfAreaTypeNssaDefaultInformationOriginateXml struct {
	Metric     *int64        `xml:"metric,omitempty"`
	MetricType *string       `xml:"metric-type,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeNssaAbrXml struct {
	ImportList         *string                                         `xml:"import-list,omitempty"`
	ExportList         *string                                         `xml:"export-list,omitempty"`
	InboundFilterList  *string                                         `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string                                         `xml:"outbound-filter-list,omitempty"`
	NssaExtRange       *vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml `xml:"nssa-ext-range,omitempty"`
	Misc               []generic.Xml                                   `xml:",any"`
}
type vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml struct {
	Entries []vrfOspfAreaTypeNssaAbrNssaExtRangeXml `xml:"entry"`
}
type vrfOspfAreaTypeNssaAbrNssaExtRangeXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfAreaRangeContainerXml struct {
	Entries []vrfOspfAreaRangeXml `xml:"entry"`
}
type vrfOspfAreaRangeXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceContainerXml struct {
	Entries []vrfOspfAreaInterfaceXml `xml:"entry"`
}
type vrfOspfAreaInterfaceXml struct {
	XMLName        xml.Name                         `xml:"entry"`
	Name           string                           `xml:"name,attr"`
	Enable         *string                          `xml:"enable,omitempty"`
	MtuIgnore      *string                          `xml:"mtu-ignore,omitempty"`
	Passive        *string                          `xml:"passive,omitempty"`
	Priority       *int64                           `xml:"priority,omitempty"`
	Metric         *int64                           `xml:"metric,omitempty"`
	Authentication *string                          `xml:"authentication,omitempty"`
	Timing         *string                          `xml:"timing,omitempty"`
	LinkType       *vrfOspfAreaInterfaceLinkTypeXml `xml:"link-type,omitempty"`
	Bfd            *vrfOspfAreaInterfaceBfdXml      `xml:"bfd,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeXml struct {
	Broadcast *vrfOspfAreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2p       *vrfOspfAreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`
	P2mp      *vrfOspfAreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                             `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpXml struct {
	Neighbor *vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml `xml:"neighbor,omitempty"`
	Misc     []generic.Xml                                         `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml struct {
	Entries []vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml `xml:"entry"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Priority *int64        `xml:"priority,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfAreaVirtualLinkContainerXml struct {
	Entries []vrfOspfAreaVirtualLinkXml `xml:"entry"`
}
type vrfOspfAreaVirtualLinkXml struct {
	XMLName        xml.Name                      `xml:"entry"`
	Name           string                        `xml:"name,attr"`
	NeighborId     *string                       `xml:"neighbor-id,omitempty"`
	TransitAreaId  *string                       `xml:"transit-area-id,omitempty"`
	Enable         *string                       `xml:"enable,omitempty"`
	InstanceId     *int64                        `xml:"instance-id,omitempty"`
	Timing         *string                       `xml:"timing,omitempty"`
	Authentication *string                       `xml:"authentication,omitempty"`
	Bfd            *vrfOspfAreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	Misc           []generic.Xml                 `xml:",any"`
}
type vrfOspfAreaVirtualLinkBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3Xml struct {
	Enable                *string                      `xml:"enable,omitempty"`
	RouterId              *string                      `xml:"router-id,omitempty"`
	DisableTransitTraffic *string                      `xml:"disable-transit-traffic,omitempty"`
	SpfTimer              *string                      `xml:"spf-timer,omitempty"`
	GlobalIfTimer         *string                      `xml:"global-if-timer,omitempty"`
	RedistributionProfile *string                      `xml:"redistribution-profile,omitempty"`
	GlobalBfd             *vrfOspfv3GlobalBfdXml       `xml:"global-bfd,omitempty"`
	GracefulRestart       *vrfOspfv3GracefulRestartXml `xml:"graceful-restart,omitempty"`
	Area                  *vrfOspfv3AreaContainerXml   `xml:"area,omitempty"`
	Misc                  []generic.Xml                `xml:",any"`
}
type vrfOspfv3GlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3GracefulRestartXml struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaContainerXml struct {
	Entries []vrfOspfv3AreaXml `xml:"entry"`
}
type vrfOspfv3AreaXml struct {
	XMLName        xml.Name                              `xml:"entry"`
	Name           string                                `xml:"name,attr"`
	Authentication *string                               `xml:"authentication,omitempty"`
	Type           *vrfOspfv3AreaTypeXml                 `xml:"type,omitempty"`
	Range          *vrfOspfv3AreaRangeContainerXml       `xml:"range,omitempty"`
	Interface      *vrfOspfv3AreaInterfaceContainerXml   `xml:"interface,omitempty"`
	VirtualLink    *vrfOspfv3AreaVirtualLinkContainerXml `xml:"virtual-link,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
}
type vrfOspfv3AreaTypeXml struct {
	Normal *vrfOspfv3AreaTypeNormalXml `xml:"normal,omitempty"`
	Stub   *vrfOspfv3AreaTypeStubXml   `xml:"stub,omitempty"`
	Nssa   *vrfOspfv3AreaTypeNssaXml   `xml:"nssa,omitempty"`
	Misc   []generic.Xml               `xml:",any"`
}
type vrfOspfv3AreaTypeNormalXml struct {
	Abr  *vrfOspfv3AreaTypeNormalAbrXml `xml:"abr,omitempty"`
	Misc []generic.Xml                  `xml:",any"`
}
type vrfOspfv3AreaTypeNormalAbrXml struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeStubXml struct {
	NoSummary          *string                      `xml:"no-summary,omitempty"`
	Abr                *vrfOspfv3AreaTypeStubAbrXml `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                       `xml:"default-route-metric,omitempty"`
	Misc               []generic.Xml                `xml:",any"`
}
type vrfOspfv3AreaTypeStubAbrXml struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeNssaXml struct {
	NoSummary                   *string                                              `xml:"no-summary,omitempty"`
	DefaultInformationOriginate *vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml `xml:"default-information-originate,omitempty"`
	Abr                         *vrfOspfv3AreaTypeNssaAbrXml                         `xml:"abr,omitempty"`
	Misc                        []generic.Xml                                        `xml:",any"`
}
type vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml struct {
	Metric     *int64        `xml:"metric,omitempty"`
	MetricType *string       `xml:"metric-type,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeNssaAbrXml struct {
	ImportList         *string                                           `xml:"import-list,omitempty"`
	ExportList         *string                                           `xml:"export-list,omitempty"`
	InboundFilterList  *string                                           `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string                                           `xml:"outbound-filter-list,omitempty"`
	NssaExtRange       *vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml `xml:"nssa-ext-range,omitempty"`
	Misc               []generic.Xml                                     `xml:",any"`
}
type vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml struct {
	Entries []vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml `xml:"entry"`
}
type vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaRangeContainerXml struct {
	Entries []vrfOspfv3AreaRangeXml `xml:"entry"`
}
type vrfOspfv3AreaRangeXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceContainerXml struct {
	Entries []vrfOspfv3AreaInterfaceXml `xml:"entry"`
}
type vrfOspfv3AreaInterfaceXml struct {
	XMLName        xml.Name                           `xml:"entry"`
	Name           string                             `xml:"name,attr"`
	Enable         *string                            `xml:"enable,omitempty"`
	MtuIgnore      *string                            `xml:"mtu-ignore,omitempty"`
	Passive        *string                            `xml:"passive,omitempty"`
	Priority       *int64                             `xml:"priority,omitempty"`
	Metric         *int64                             `xml:"metric,omitempty"`
	InstanceId     *int64                             `xml:"instance-id,omitempty"`
	Authentication *string                            `xml:"authentication,omitempty"`
	Timing         *string                            `xml:"timing,omitempty"`
	LinkType       *vrfOspfv3AreaInterfaceLinkTypeXml `xml:"link-type,omitempty"`
	Bfd            *vrfOspfv3AreaInterfaceBfdXml      `xml:"bfd,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeXml struct {
	Broadcast *vrfOspfv3AreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2p       *vrfOspfv3AreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`
	P2mp      *vrfOspfv3AreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                               `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpXml struct {
	Neighbor *vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml `xml:"neighbor,omitempty"`
	Misc     []generic.Xml                                           `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml struct {
	Entries []vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml `xml:"entry"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Priority *int64        `xml:"priority,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaVirtualLinkContainerXml struct {
	Entries []vrfOspfv3AreaVirtualLinkXml `xml:"entry"`
}
type vrfOspfv3AreaVirtualLinkXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	NeighborId     *string       `xml:"neighbor-id,omitempty"`
	TransitAreaId  *string       `xml:"transit-area-id,omitempty"`
	Enable         *string       `xml:"enable,omitempty"`
	InstanceId     *int64        `xml:"instance-id,omitempty"`
	Timing         *string       `xml:"timing,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfEcmpXml struct {
	Enable           *string              `xml:"enable,omitempty"`
	MaxPath          *int64               `xml:"max-path,omitempty"`
	SymmetricReturn  *string              `xml:"symmetric-return,omitempty"`
	StrictSourcePath *string              `xml:"strict-source-path,omitempty"`
	Algorithm        *vrfEcmpAlgorithmXml `xml:"algorithm,omitempty"`
	Misc             []generic.Xml        `xml:",any"`
}
type vrfEcmpAlgorithmXml struct {
	IpModulo           *vrfEcmpAlgorithmIpModuloXml           `xml:"ip-modulo,omitempty"`
	IpHash             *vrfEcmpAlgorithmIpHashXml             `xml:"ip-hash,omitempty"`
	WeightedRoundRobin *vrfEcmpAlgorithmWeightedRoundRobinXml `xml:"weighted-round-robin,omitempty"`
	BalancedRoundRobin *vrfEcmpAlgorithmBalancedRoundRobinXml `xml:"balanced-round-robin,omitempty"`
	Misc               []generic.Xml                          `xml:",any"`
}
type vrfEcmpAlgorithmIpModuloXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmIpHashXml struct {
	SrcOnly  *string       `xml:"src-only,omitempty"`
	UsePort  *string       `xml:"use-port,omitempty"`
	HashSeed *int64        `xml:"hash-seed,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmWeightedRoundRobinXml struct {
	Interface *vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml `xml:"interface,omitempty"`
	Misc      []generic.Xml                                            `xml:",any"`
}
type vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml struct {
	Entries []vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml `xml:"entry"`
}
type vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Weight  *int64        `xml:"weight,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmBalancedRoundRobinXml struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfMulticastXml struct {
	Enable      *string                              `xml:"enable,omitempty"`
	StaticRoute *vrfMulticastStaticRouteContainerXml `xml:"static-route,omitempty"`
	Pim         *vrfMulticastPimXml                  `xml:"pim,omitempty"`
	Igmp        *vrfMulticastIgmpXml                 `xml:"igmp,omitempty"`
	Msdp        *vrfMulticastMsdpXml                 `xml:"msdp,omitempty"`
	Misc        []generic.Xml                        `xml:",any"`
}
type vrfMulticastStaticRouteContainerXml struct {
	Entries []vrfMulticastStaticRouteXml `xml:"entry"`
}
type vrfMulticastStaticRouteXml struct {
	XMLName     xml.Name                           `xml:"entry"`
	Name        string                             `xml:"name,attr"`
	Destination *string                            `xml:"destination,omitempty"`
	Interface   *string                            `xml:"interface,omitempty"`
	Preference  *int64                             `xml:"preference,omitempty"`
	Nexthop     *vrfMulticastStaticRouteNexthopXml `xml:"nexthop,omitempty"`
	Misc        []generic.Xml                      `xml:",any"`
}
type vrfMulticastStaticRouteNexthopXml struct {
	IpAddress *string       `xml:"ip-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimXml struct {
	Enable          *string                                  `xml:"enable,omitempty"`
	RpfLookupMode   *string                                  `xml:"rpf-lookup-mode,omitempty"`
	RouteAgeoutTime *int64                                   `xml:"route-ageout-time,omitempty"`
	IfTimerGlobal   *string                                  `xml:"if-timer-global,omitempty"`
	GroupPermission *string                                  `xml:"group-permission,omitempty"`
	SsmAddressSpace *vrfMulticastPimSsmAddressSpaceXml       `xml:"ssm-address-space,omitempty"`
	Rp              *vrfMulticastPimRpXml                    `xml:"rp,omitempty"`
	SptThreshold    *vrfMulticastPimSptThresholdContainerXml `xml:"spt-threshold,omitempty"`
	Interface       *vrfMulticastPimInterfaceContainerXml    `xml:"interface,omitempty"`
	Misc            []generic.Xml                            `xml:",any"`
}
type vrfMulticastPimSsmAddressSpaceXml struct {
	GroupList *string       `xml:"group-list,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpXml struct {
	LocalRp    *vrfMulticastPimRpLocalRpXml             `xml:"local-rp,omitempty"`
	ExternalRp *vrfMulticastPimRpExternalRpContainerXml `xml:"external-rp,omitempty"`
	Misc       []generic.Xml                            `xml:",any"`
}
type vrfMulticastPimRpLocalRpXml struct {
	StaticRp    *vrfMulticastPimRpLocalRpStaticRpXml    `xml:"static-rp,omitempty"`
	CandidateRp *vrfMulticastPimRpLocalRpCandidateRpXml `xml:"candidate-rp,omitempty"`
	Misc        []generic.Xml                           `xml:",any"`
}
type vrfMulticastPimRpLocalRpStaticRpXml struct {
	Interface *string       `xml:"interface,omitempty"`
	Address   *string       `xml:"address,omitempty"`
	Override  *string       `xml:"override,omitempty"`
	GroupList *string       `xml:"group-list,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpLocalRpCandidateRpXml struct {
	Interface             *string       `xml:"interface,omitempty"`
	Address               *string       `xml:"address,omitempty"`
	Priority              *int64        `xml:"priority,omitempty"`
	AdvertisementInterval *int64        `xml:"advertisement-interval,omitempty"`
	GroupList             *string       `xml:"group-list,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpExternalRpContainerXml struct {
	Entries []vrfMulticastPimRpExternalRpXml `xml:"entry"`
}
type vrfMulticastPimRpExternalRpXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	GroupList *string       `xml:"group-list,omitempty"`
	Override  *string       `xml:"override,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimSptThresholdContainerXml struct {
	Entries []vrfMulticastPimSptThresholdXml `xml:"entry"`
}
type vrfMulticastPimSptThresholdXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Threshold *string       `xml:"threshold,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimInterfaceContainerXml struct {
	Entries []vrfMulticastPimInterfaceXml `xml:"entry"`
}
type vrfMulticastPimInterfaceXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	DrPriority     *int64        `xml:"dr-priority,omitempty"`
	SendBsm        *string       `xml:"send-bsm,omitempty"`
	IfTimer        *string       `xml:"if-timer,omitempty"`
	NeighborFilter *string       `xml:"neighbor-filter,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfMulticastIgmpXml struct {
	Enable  *string                             `xml:"enable,omitempty"`
	Dynamic *vrfMulticastIgmpDynamicXml         `xml:"dynamic,omitempty"`
	Static  *vrfMulticastIgmpStaticContainerXml `xml:"static,omitempty"`
	Misc    []generic.Xml                       `xml:",any"`
}
type vrfMulticastIgmpDynamicXml struct {
	Interface *vrfMulticastIgmpDynamicInterfaceContainerXml `xml:"interface,omitempty"`
	Misc      []generic.Xml                                 `xml:",any"`
}
type vrfMulticastIgmpDynamicInterfaceContainerXml struct {
	Entries []vrfMulticastIgmpDynamicInterfaceXml `xml:"entry"`
}
type vrfMulticastIgmpDynamicInterfaceXml struct {
	XMLName             xml.Name      `xml:"entry"`
	Name                string        `xml:"name,attr"`
	Version             *string       `xml:"version,omitempty"`
	Robustness          *string       `xml:"robustness,omitempty"`
	GroupFilter         *string       `xml:"group-filter,omitempty"`
	MaxGroups           *string       `xml:"max-groups,omitempty"`
	MaxSources          *string       `xml:"max-sources,omitempty"`
	QueryProfile        *string       `xml:"query-profile,omitempty"`
	RouterAlertPolicing *string       `xml:"router-alert-policing,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type vrfMulticastIgmpStaticContainerXml struct {
	Entries []vrfMulticastIgmpStaticXml `xml:"entry"`
}
type vrfMulticastIgmpStaticXml struct {
	XMLName       xml.Name      `xml:"entry"`
	Name          string        `xml:"name,attr"`
	Interface     *string       `xml:"interface,omitempty"`
	GroupAddress  *string       `xml:"group-address,omitempty"`
	SourceAddress *string       `xml:"source-address,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpXml struct {
	Enable               *string                           `xml:"enable,omitempty"`
	GlobalTimer          *string                           `xml:"global-timer,omitempty"`
	GlobalAuthentication *string                           `xml:"global-authentication,omitempty"`
	OriginatorId         *vrfMulticastMsdpOriginatorIdXml  `xml:"originator-id,omitempty"`
	Peer                 *vrfMulticastMsdpPeerContainerXml `xml:"peer,omitempty"`
	Misc                 []generic.Xml                     `xml:",any"`
}
type vrfMulticastMsdpOriginatorIdXml struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpPeerContainerXml struct {
	Entries []vrfMulticastMsdpPeerXml `xml:"entry"`
}
type vrfMulticastMsdpPeerXml struct {
	XMLName          xml.Name                             `xml:"entry"`
	Name             string                               `xml:"name,attr"`
	Enable           *string                              `xml:"enable,omitempty"`
	PeerAs           *string                              `xml:"peer-as,omitempty"`
	Authentication   *string                              `xml:"authentication,omitempty"`
	MaxSa            *int64                               `xml:"max-sa,omitempty"`
	InboundSaFilter  *string                              `xml:"inbound-sa-filter,omitempty"`
	OutboundSaFilter *string                              `xml:"outbound-sa-filter,omitempty"`
	LocalAddress     *vrfMulticastMsdpPeerLocalAddressXml `xml:"local-address,omitempty"`
	PeerAddress      *vrfMulticastMsdpPeerPeerAddressXml  `xml:"peer-address,omitempty"`
	Misc             []generic.Xml                        `xml:",any"`
}
type vrfMulticastMsdpPeerLocalAddressXml struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpPeerPeerAddressXml struct {
	Ip   *string       `xml:"ip,omitempty"`
	Fqdn *string       `xml:"fqdn,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfRipXml struct {
	Enable                       *string                                `xml:"enable,omitempty"`
	DefaultInformationOriginate  *string                                `xml:"default-information-originate,omitempty"`
	GlobalTimer                  *string                                `xml:"global-timer,omitempty"`
	AuthProfile                  *string                                `xml:"auth-profile,omitempty"`
	RedistributionProfile        *string                                `xml:"redistribution-profile,omitempty"`
	GlobalBfd                    *vrfRipGlobalBfdXml                    `xml:"global-bfd,omitempty"`
	GlobalInboundDistributeList  *vrfRipGlobalInboundDistributeListXml  `xml:"global-inbound-distribute-list,omitempty"`
	GlobalOutboundDistributeList *vrfRipGlobalOutboundDistributeListXml `xml:"global-outbound-distribute-list,omitempty"`
	Interface                    *vrfRipInterfaceContainerXml           `xml:"interface,omitempty"`
	Misc                         []generic.Xml                          `xml:",any"`
}
type vrfRipGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRipGlobalInboundDistributeListXml struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipGlobalOutboundDistributeListXml struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipInterfaceContainerXml struct {
	Entries []vrfRipInterfaceXml `xml:"entry"`
}
type vrfRipInterfaceXml struct {
	XMLName                         xml.Name                                           `xml:"entry"`
	Name                            string                                             `xml:"name,attr"`
	Enable                          *string                                            `xml:"enable,omitempty"`
	Mode                            *string                                            `xml:"mode,omitempty"`
	SplitHorizon                    *string                                            `xml:"split-horizon,omitempty"`
	Authentication                  *string                                            `xml:"authentication,omitempty"`
	Bfd                             *vrfRipInterfaceBfdXml                             `xml:"bfd,omitempty"`
	InterfaceInboundDistributeList  *vrfRipInterfaceInterfaceInboundDistributeListXml  `xml:"interface-inbound-distribute-list,omitempty"`
	InterfaceOutboundDistributeList *vrfRipInterfaceInterfaceOutboundDistributeListXml `xml:"interface-outbound-distribute-list,omitempty"`
	Misc                            []generic.Xml                                      `xml:",any"`
}
type vrfRipInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRipInterfaceInterfaceInboundDistributeListXml struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Metric     *int64        `xml:"metric,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipInterfaceInterfaceOutboundDistributeListXml struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Metric     *int64        `xml:"metric,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName xml.Name                `xml:"entry"`
	Name    string                  `xml:"name,attr"`
	Vrf     *vrfContainerXml_11_0_2 `xml:"vrf,omitempty"`
	Misc    []generic.Xml           `xml:",any"`
}
type vrfContainerXml_11_0_2 struct {
	Entries []vrfXml_11_0_2 `xml:"entry"`
}
type vrfXml_11_0_2 struct {
	XMLName      xml.Name                   `xml:"entry"`
	Name         string                     `xml:"name,attr"`
	Interface    *util.MemberType           `xml:"interface,omitempty"`
	AdminDists   *vrfAdminDistsXml_11_0_2   `xml:"admin-dists,omitempty"`
	RibFilter    *vrfRibFilterXml_11_0_2    `xml:"rib-filter,omitempty"`
	Bgp          *vrfBgpXml_11_0_2          `xml:"bgp,omitempty"`
	RoutingTable *vrfRoutingTableXml_11_0_2 `xml:"routing-table,omitempty"`
	Ospf         *vrfOspfXml_11_0_2         `xml:"ospf,omitempty"`
	Ospfv3       *vrfOspfv3Xml_11_0_2       `xml:"ospfv3,omitempty"`
	Ecmp         *vrfEcmpXml_11_0_2         `xml:"ecmp,omitempty"`
	Multicast    *vrfMulticastXml_11_0_2    `xml:"multicast,omitempty"`
	Rip          *vrfRipXml_11_0_2          `xml:"rip,omitempty"`
	Misc         []generic.Xml              `xml:",any"`
}
type vrfAdminDistsXml_11_0_2 struct {
	Static      *int64        `xml:"static,omitempty"`
	StaticIpv6  *int64        `xml:"static-ipv6,omitempty"`
	OspfInter   *int64        `xml:"ospf-inter,omitempty"`
	OspfIntra   *int64        `xml:"ospf-intra,omitempty"`
	OspfExt     *int64        `xml:"ospf-ext,omitempty"`
	Ospfv3Inter *int64        `xml:"ospfv3-inter,omitempty"`
	Ospfv3Intra *int64        `xml:"ospfv3-intra,omitempty"`
	Ospfv3Ext   *int64        `xml:"ospfv3-ext,omitempty"`
	BgpInternal *int64        `xml:"bgp-internal,omitempty"`
	BgpExternal *int64        `xml:"bgp-external,omitempty"`
	BgpLocal    *int64        `xml:"bgp-local,omitempty"`
	Rip         *int64        `xml:"rip,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfRibFilterXml_11_0_2 struct {
	Ipv4 *vrfRibFilterIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *vrfRibFilterIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc []generic.Xml               `xml:",any"`
}
type vrfRibFilterIpv4Xml_11_0_2 struct {
	Static *vrfRibFilterIpv4StaticXml_11_0_2 `xml:"static,omitempty"`
	Bgp    *vrfRibFilterIpv4BgpXml_11_0_2    `xml:"bgp,omitempty"`
	Ospf   *vrfRibFilterIpv4OspfXml_11_0_2   `xml:"ospf,omitempty"`
	Rip    *vrfRibFilterIpv4RipXml_11_0_2    `xml:"rip,omitempty"`
	Misc   []generic.Xml                     `xml:",any"`
}
type vrfRibFilterIpv4StaticXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4BgpXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4OspfXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv4RipXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6Xml_11_0_2 struct {
	Static *vrfRibFilterIpv6StaticXml_11_0_2 `xml:"static,omitempty"`
	Bgp    *vrfRibFilterIpv6BgpXml_11_0_2    `xml:"bgp,omitempty"`
	Ospfv3 *vrfRibFilterIpv6Ospfv3Xml_11_0_2 `xml:"ospfv3,omitempty"`
	Misc   []generic.Xml                     `xml:",any"`
}
type vrfRibFilterIpv6StaticXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6BgpXml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfRibFilterIpv6Ospfv3Xml_11_0_2 struct {
	RouteMap *string       `xml:"route-map,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfBgpXml_11_0_2 struct {
	Enable                      *string                                   `xml:"enable,omitempty"`
	RouterId                    *string                                   `xml:"router-id,omitempty"`
	LocalAs                     *string                                   `xml:"local-as,omitempty"`
	InstallRoute                *string                                   `xml:"install-route,omitempty"`
	EnforceFirstAs              *string                                   `xml:"enforce-first-as,omitempty"`
	FastExternalFailover        *string                                   `xml:"fast-external-failover,omitempty"`
	EcmpMultiAs                 *string                                   `xml:"ecmp-multi-as,omitempty"`
	DefaultLocalPreference      *int64                                    `xml:"default-local-preference,omitempty"`
	GracefulShutdown            *string                                   `xml:"graceful-shutdown,omitempty"`
	AlwaysAdvertiseNetworkRoute *string                                   `xml:"always-advertise-network-route,omitempty"`
	Med                         *vrfBgpMedXml_11_0_2                      `xml:"med,omitempty"`
	GracefulRestart             *vrfBgpGracefulRestartXml_11_0_2          `xml:"graceful-restart,omitempty"`
	GlobalBfd                   *vrfBgpGlobalBfdXml_11_0_2                `xml:"global-bfd,omitempty"`
	RedistributionProfile       *vrfBgpRedistributionProfileXml_11_0_2    `xml:"redistribution-profile,omitempty"`
	AdvertiseNetwork            *vrfBgpAdvertiseNetworkXml_11_0_2         `xml:"advertise-network,omitempty"`
	PeerGroup                   *vrfBgpPeerGroupContainerXml_11_0_2       `xml:"peer-group,omitempty"`
	AggregateRoutes             *vrfBgpAggregateRoutesContainerXml_11_0_2 `xml:"aggregate-routes,omitempty"`
	Misc                        []generic.Xml                             `xml:",any"`
}
type vrfBgpMedXml_11_0_2 struct {
	AlwaysCompareMed           *string       `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string       `xml:"deterministic-med-comparison,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type vrfBgpGracefulRestartXml_11_0_2 struct {
	Enable             *string       `xml:"enable,omitempty"`
	StaleRouteTime     *int64        `xml:"stale-route-time,omitempty"`
	MaxPeerRestartTime *int64        `xml:"max-peer-restart-time,omitempty"`
	LocalRestartTime   *int64        `xml:"local-restart-time,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfBgpGlobalBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpRedistributionProfileXml_11_0_2 struct {
	Ipv4 *vrfBgpRedistributionProfileIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpRedistributionProfileIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc []generic.Xml                              `xml:",any"`
}
type vrfBgpRedistributionProfileIpv4Xml_11_0_2 struct {
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpRedistributionProfileIpv6Xml_11_0_2 struct {
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpAdvertiseNetworkXml_11_0_2 struct {
	Ipv4 *vrfBgpAdvertiseNetworkIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpAdvertiseNetworkIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc []generic.Xml                         `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv4Xml_11_0_2 struct {
	Network *vrfBgpAdvertiseNetworkIpv4NetworkContainerXml_11_0_2 `xml:"network,omitempty"`
	Misc    []generic.Xml                                         `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv4NetworkContainerXml_11_0_2 struct {
	Entries []vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2 `xml:"entry"`
}
type vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Unicast   *string       `xml:"unicast,omitempty"`
	Multicast *string       `xml:"multicast,omitempty"`
	Backdoor  *string       `xml:"backdoor,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv6Xml_11_0_2 struct {
	Network *vrfBgpAdvertiseNetworkIpv6NetworkContainerXml_11_0_2 `xml:"network,omitempty"`
	Misc    []generic.Xml                                         `xml:",any"`
}
type vrfBgpAdvertiseNetworkIpv6NetworkContainerXml_11_0_2 struct {
	Entries []vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2 `xml:"entry"`
}
type vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Unicast *string       `xml:"unicast,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupContainerXml_11_0_2 struct {
	Entries []vrfBgpPeerGroupXml_11_0_2 `xml:"entry"`
}
type vrfBgpPeerGroupXml_11_0_2 struct {
	XMLName           xml.Name                                    `xml:"entry"`
	Name              string                                      `xml:"name,attr"`
	Enable            *string                                     `xml:"enable,omitempty"`
	Type              *vrfBgpPeerGroupTypeXml_11_0_2              `xml:"type,omitempty"`
	AddressFamily     *vrfBgpPeerGroupAddressFamilyXml_11_0_2     `xml:"address-family,omitempty"`
	FilteringProfile  *vrfBgpPeerGroupFilteringProfileXml_11_0_2  `xml:"filtering-profile,omitempty"`
	ConnectionOptions *vrfBgpPeerGroupConnectionOptionsXml_11_0_2 `xml:"connection-options,omitempty"`
	Peer              *vrfBgpPeerGroupPeerContainerXml_11_0_2     `xml:"peer,omitempty"`
	Misc              []generic.Xml                               `xml:",any"`
}
type vrfBgpPeerGroupTypeXml_11_0_2 struct {
	Ibgp *vrfBgpPeerGroupTypeIbgpXml_11_0_2 `xml:"ibgp,omitempty"`
	Ebgp *vrfBgpPeerGroupTypeEbgpXml_11_0_2 `xml:"ebgp,omitempty"`
	Misc []generic.Xml                      `xml:",any"`
}
type vrfBgpPeerGroupTypeIbgpXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupTypeEbgpXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupAddressFamilyXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupFilteringProfileXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupConnectionOptionsXml_11_0_2 struct {
	Timers         *string       `xml:"timers,omitempty"`
	Multihop       *int64        `xml:"multihop,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Dampening      *string       `xml:"dampening,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerContainerXml_11_0_2 struct {
	Entries []vrfBgpPeerGroupPeerXml_11_0_2 `xml:"entry"`
}
type vrfBgpPeerGroupPeerXml_11_0_2 struct {
	XMLName                       xml.Name                                        `xml:"entry"`
	Name                          string                                          `xml:"name,attr"`
	Enable                        *string                                         `xml:"enable,omitempty"`
	Passive                       *string                                         `xml:"passive,omitempty"`
	PeerAs                        *string                                         `xml:"peer-as,omitempty"`
	EnableSenderSideLoopDetection *string                                         `xml:"enable-sender-side-loop-detection,omitempty"`
	Inherit                       *vrfBgpPeerGroupPeerInheritXml_11_0_2           `xml:"inherit,omitempty"`
	LocalAddress                  *vrfBgpPeerGroupPeerLocalAddressXml_11_0_2      `xml:"local-address,omitempty"`
	PeerAddress                   *vrfBgpPeerGroupPeerPeerAddressXml_11_0_2       `xml:"peer-address,omitempty"`
	ConnectionOptions             *vrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2 `xml:"connection-options,omitempty"`
	Bfd                           *vrfBgpPeerGroupPeerBfdXml_11_0_2               `xml:"bfd,omitempty"`
	Misc                          []generic.Xml                                   `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritXml_11_0_2 struct {
	Yes  *vrfBgpPeerGroupPeerInheritYesXml_11_0_2 `xml:"yes,omitempty"`
	No   *vrfBgpPeerGroupPeerInheritNoXml_11_0_2  `xml:"no,omitempty"`
	Misc []generic.Xml                            `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritYesXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoXml_11_0_2 struct {
	AddressFamily    *vrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2    `xml:"address-family,omitempty"`
	FilteringProfile *vrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2 `xml:"filtering-profile,omitempty"`
	Misc             []generic.Xml                                           `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerLocalAddressXml_11_0_2 struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerPeerAddressXml_11_0_2 struct {
	Ip   *string       `xml:"ip,omitempty"`
	Fqdn *string       `xml:"fqdn,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2 struct {
	Timers         *string       `xml:"timers,omitempty"`
	Multihop       *string       `xml:"multihop,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Dampening      *string       `xml:"dampening,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfBgpPeerGroupPeerBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfBgpAggregateRoutesContainerXml_11_0_2 struct {
	Entries []vrfBgpAggregateRoutesXml_11_0_2 `xml:"entry"`
}
type vrfBgpAggregateRoutesXml_11_0_2 struct {
	XMLName     xml.Name                             `xml:"entry"`
	Name        string                               `xml:"name,attr"`
	Description *string                              `xml:"description,omitempty"`
	Enable      *string                              `xml:"enable,omitempty"`
	SummaryOnly *string                              `xml:"summary-only,omitempty"`
	AsSet       *string                              `xml:"as-set,omitempty"`
	SameMed     *string                              `xml:"same-med,omitempty"`
	Type        *vrfBgpAggregateRoutesTypeXml_11_0_2 `xml:"type,omitempty"`
	Misc        []generic.Xml                        `xml:",any"`
}
type vrfBgpAggregateRoutesTypeXml_11_0_2 struct {
	Ipv4 *vrfBgpAggregateRoutesTypeIpv4Xml_11_0_2 `xml:"ipv4,omitempty"`
	Ipv6 *vrfBgpAggregateRoutesTypeIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc []generic.Xml                            `xml:",any"`
}
type vrfBgpAggregateRoutesTypeIpv4Xml_11_0_2 struct {
	SummaryPrefix *string       `xml:"summary-prefix,omitempty"`
	SuppressMap   *string       `xml:"suppress-map,omitempty"`
	AttributeMap  *string       `xml:"attribute-map,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfBgpAggregateRoutesTypeIpv6Xml_11_0_2 struct {
	SummaryPrefix *string       `xml:"summary-prefix,omitempty"`
	SuppressMap   *string       `xml:"suppress-map,omitempty"`
	AttributeMap  *string       `xml:"attribute-map,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfRoutingTableXml_11_0_2 struct {
	Ip   *vrfRoutingTableIpXml_11_0_2   `xml:"ip,omitempty"`
	Ipv6 *vrfRoutingTableIpv6Xml_11_0_2 `xml:"ipv6,omitempty"`
	Misc []generic.Xml                  `xml:",any"`
}
type vrfRoutingTableIpXml_11_0_2 struct {
	StaticRoute *vrfRoutingTableIpStaticRouteContainerXml_11_0_2 `xml:"static-route,omitempty"`
	Misc        []generic.Xml                                    `xml:",any"`
}
type vrfRoutingTableIpStaticRouteContainerXml_11_0_2 struct {
	Entries []vrfRoutingTableIpStaticRouteXml_11_0_2 `xml:"entry"`
}
type vrfRoutingTableIpStaticRouteXml_11_0_2 struct {
	XMLName     xml.Name                                           `xml:"entry"`
	Name        string                                             `xml:"name,attr"`
	Destination *string                                            `xml:"destination,omitempty"`
	Interface   *string                                            `xml:"interface,omitempty"`
	AdminDist   *int64                                             `xml:"admin-dist,omitempty"`
	Metric      *int64                                             `xml:"metric,omitempty"`
	Nexthop     *vrfRoutingTableIpStaticRouteNexthopXml_11_0_2     `xml:"nexthop,omitempty"`
	Bfd         *vrfRoutingTableIpStaticRouteBfdXml_11_0_2         `xml:"bfd,omitempty"`
	PathMonitor *vrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2 `xml:"path-monitor,omitempty"`
	Misc        []generic.Xml                                      `xml:",any"`
}
type vrfRoutingTableIpStaticRouteNexthopXml_11_0_2 struct {
	Discard   *vrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2 `xml:"discard,omitempty"`
	IpAddress *string                                               `xml:"ip-address,omitempty"`
	NextLr    *string                                               `xml:"next-lr,omitempty"`
	Fqdn      *string                                               `xml:"fqdn,omitempty"`
	Misc      []generic.Xml                                         `xml:",any"`
}
type vrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpStaticRouteBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2 struct {
	Enable              *string                                                                        `xml:"enable,omitempty"`
	FailureCondition    *string                                                                        `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                         `xml:"hold-time,omitempty"`
	MonitorDestinations *vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2 `xml:"monitor-destinations,omitempty"`
	Misc                []generic.Xml                                                                  `xml:",any"`
}
type vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2 struct {
	Entries []vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2 `xml:"entry"`
}
type vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2 struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Enable      *string       `xml:"enable,omitempty"`
	Source      *string       `xml:"source,omitempty"`
	Destination *string       `xml:"destination,omitempty"`
	Interval    *int64        `xml:"interval,omitempty"`
	Count       *int64        `xml:"count,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6Xml_11_0_2 struct {
	StaticRoute *vrfRoutingTableIpv6StaticRouteContainerXml_11_0_2 `xml:"static-route,omitempty"`
	Misc        []generic.Xml                                      `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteContainerXml_11_0_2 struct {
	Entries []vrfRoutingTableIpv6StaticRouteXml_11_0_2 `xml:"entry"`
}
type vrfRoutingTableIpv6StaticRouteXml_11_0_2 struct {
	XMLName     xml.Name                                             `xml:"entry"`
	Name        string                                               `xml:"name,attr"`
	Destination *string                                              `xml:"destination,omitempty"`
	Interface   *string                                              `xml:"interface,omitempty"`
	AdminDist   *int64                                               `xml:"admin-dist,omitempty"`
	Metric      *int64                                               `xml:"metric,omitempty"`
	Nexthop     *vrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2     `xml:"nexthop,omitempty"`
	Bfd         *vrfRoutingTableIpv6StaticRouteBfdXml_11_0_2         `xml:"bfd,omitempty"`
	PathMonitor *vrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2 `xml:"path-monitor,omitempty"`
	Misc        []generic.Xml                                        `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2 struct {
	Discard     *vrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2 `xml:"discard,omitempty"`
	Ipv6Address *string                                                 `xml:"ipv6-address,omitempty"`
	Fqdn        *string                                                 `xml:"fqdn,omitempty"`
	NextLr      *string                                                 `xml:"next-lr,omitempty"`
	Misc        []generic.Xml                                           `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6StaticRouteBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2 struct {
	Enable              *string                                                                          `xml:"enable,omitempty"`
	FailureCondition    *string                                                                          `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                                           `xml:"hold-time,omitempty"`
	MonitorDestinations *vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2 `xml:"monitor-destinations,omitempty"`
	Misc                []generic.Xml                                                                    `xml:",any"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2 struct {
	Entries []vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2 `xml:"entry"`
}
type vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2 struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Enable      *string       `xml:"enable,omitempty"`
	Source      *string       `xml:"source,omitempty"`
	Destination *string       `xml:"destination,omitempty"`
	Interval    *int64        `xml:"interval,omitempty"`
	Count       *int64        `xml:"count,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type vrfOspfXml_11_0_2 struct {
	RouterId              *string                           `xml:"router-id,omitempty"`
	Enable                *string                           `xml:"enable,omitempty"`
	Rfc1583               *string                           `xml:"rfc1583,omitempty"`
	SpfTimer              *string                           `xml:"spf-timer,omitempty"`
	GlobalIfTimer         *string                           `xml:"global-if-timer,omitempty"`
	RedistributionProfile *string                           `xml:"redistribution-profile,omitempty"`
	GlobalBfd             *vrfOspfGlobalBfdXml_11_0_2       `xml:"global-bfd,omitempty"`
	GracefulRestart       *vrfOspfGracefulRestartXml_11_0_2 `xml:"graceful-restart,omitempty"`
	Area                  *vrfOspfAreaContainerXml_11_0_2   `xml:"area,omitempty"`
	Misc                  []generic.Xml                     `xml:",any"`
}
type vrfOspfGlobalBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfGracefulRestartXml_11_0_2 struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type vrfOspfAreaContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaXml_11_0_2 struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Authentication *string                                    `xml:"authentication,omitempty"`
	Type           *vrfOspfAreaTypeXml_11_0_2                 `xml:"type,omitempty"`
	Range          *vrfOspfAreaRangeContainerXml_11_0_2       `xml:"range,omitempty"`
	Interface      *vrfOspfAreaInterfaceContainerXml_11_0_2   `xml:"interface,omitempty"`
	VirtualLink    *vrfOspfAreaVirtualLinkContainerXml_11_0_2 `xml:"virtual-link,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
}
type vrfOspfAreaTypeXml_11_0_2 struct {
	Normal *vrfOspfAreaTypeNormalXml_11_0_2 `xml:"normal,omitempty"`
	Stub   *vrfOspfAreaTypeStubXml_11_0_2   `xml:"stub,omitempty"`
	Nssa   *vrfOspfAreaTypeNssaXml_11_0_2   `xml:"nssa,omitempty"`
	Misc   []generic.Xml                    `xml:",any"`
}
type vrfOspfAreaTypeNormalXml_11_0_2 struct {
	Abr  *vrfOspfAreaTypeNormalAbrXml_11_0_2 `xml:"abr,omitempty"`
	Misc []generic.Xml                       `xml:",any"`
}
type vrfOspfAreaTypeNormalAbrXml_11_0_2 struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeStubXml_11_0_2 struct {
	NoSummary          *string                           `xml:"no-summary,omitempty"`
	Abr                *vrfOspfAreaTypeStubAbrXml_11_0_2 `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                            `xml:"default-route-metric,omitempty"`
	Misc               []generic.Xml                     `xml:",any"`
}
type vrfOspfAreaTypeStubAbrXml_11_0_2 struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeNssaXml_11_0_2 struct {
	NoSummary                   *string                                                   `xml:"no-summary,omitempty"`
	DefaultInformationOriginate *vrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2 `xml:"default-information-originate,omitempty"`
	Abr                         *vrfOspfAreaTypeNssaAbrXml_11_0_2                         `xml:"abr,omitempty"`
	Misc                        []generic.Xml                                             `xml:",any"`
}
type vrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2 struct {
	Metric     *int64        `xml:"metric,omitempty"`
	MetricType *string       `xml:"metric-type,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfOspfAreaTypeNssaAbrXml_11_0_2 struct {
	ImportList         *string                                                `xml:"import-list,omitempty"`
	ExportList         *string                                                `xml:"export-list,omitempty"`
	InboundFilterList  *string                                                `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string                                                `xml:"outbound-filter-list,omitempty"`
	NssaExtRange       *vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2 `xml:"nssa-ext-range,omitempty"`
	Misc               []generic.Xml                                          `xml:",any"`
}
type vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfAreaRangeContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaRangeXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaRangeXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaInterfaceXml_11_0_2 struct {
	XMLName        xml.Name                                `xml:"entry"`
	Name           string                                  `xml:"name,attr"`
	Enable         *string                                 `xml:"enable,omitempty"`
	MtuIgnore      *string                                 `xml:"mtu-ignore,omitempty"`
	Passive        *string                                 `xml:"passive,omitempty"`
	Priority       *int64                                  `xml:"priority,omitempty"`
	Metric         *int64                                  `xml:"metric,omitempty"`
	Authentication *string                                 `xml:"authentication,omitempty"`
	Timing         *string                                 `xml:"timing,omitempty"`
	LinkType       *vrfOspfAreaInterfaceLinkTypeXml_11_0_2 `xml:"link-type,omitempty"`
	Bfd            *vrfOspfAreaInterfaceBfdXml_11_0_2      `xml:"bfd,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeXml_11_0_2 struct {
	Broadcast *vrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2 `xml:"broadcast,omitempty"`
	P2p       *vrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2       `xml:"p2p,omitempty"`
	P2mp      *vrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                                    `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2 struct {
	Neighbor *vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2 `xml:"neighbor,omitempty"`
	Misc     []generic.Xml                                                `xml:",any"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Priority *int64        `xml:"priority,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfOspfAreaInterfaceBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfAreaVirtualLinkContainerXml_11_0_2 struct {
	Entries []vrfOspfAreaVirtualLinkXml_11_0_2 `xml:"entry"`
}
type vrfOspfAreaVirtualLinkXml_11_0_2 struct {
	XMLName        xml.Name                             `xml:"entry"`
	Name           string                               `xml:"name,attr"`
	NeighborId     *string                              `xml:"neighbor-id,omitempty"`
	TransitAreaId  *string                              `xml:"transit-area-id,omitempty"`
	Enable         *string                              `xml:"enable,omitempty"`
	InstanceId     *int64                               `xml:"instance-id,omitempty"`
	Timing         *string                              `xml:"timing,omitempty"`
	Authentication *string                              `xml:"authentication,omitempty"`
	Bfd            *vrfOspfAreaVirtualLinkBfdXml_11_0_2 `xml:"bfd,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
}
type vrfOspfAreaVirtualLinkBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3Xml_11_0_2 struct {
	Enable                *string                             `xml:"enable,omitempty"`
	RouterId              *string                             `xml:"router-id,omitempty"`
	DisableTransitTraffic *string                             `xml:"disable-transit-traffic,omitempty"`
	SpfTimer              *string                             `xml:"spf-timer,omitempty"`
	GlobalIfTimer         *string                             `xml:"global-if-timer,omitempty"`
	RedistributionProfile *string                             `xml:"redistribution-profile,omitempty"`
	GlobalBfd             *vrfOspfv3GlobalBfdXml_11_0_2       `xml:"global-bfd,omitempty"`
	GracefulRestart       *vrfOspfv3GracefulRestartXml_11_0_2 `xml:"graceful-restart,omitempty"`
	Area                  *vrfOspfv3AreaContainerXml_11_0_2   `xml:"area,omitempty"`
	Misc                  []generic.Xml                       `xml:",any"`
}
type vrfOspfv3GlobalBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3GracefulRestartXml_11_0_2 struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaXml_11_0_2 struct {
	XMLName        xml.Name                                     `xml:"entry"`
	Name           string                                       `xml:"name,attr"`
	Authentication *string                                      `xml:"authentication,omitempty"`
	Type           *vrfOspfv3AreaTypeXml_11_0_2                 `xml:"type,omitempty"`
	Range          *vrfOspfv3AreaRangeContainerXml_11_0_2       `xml:"range,omitempty"`
	Interface      *vrfOspfv3AreaInterfaceContainerXml_11_0_2   `xml:"interface,omitempty"`
	VirtualLink    *vrfOspfv3AreaVirtualLinkContainerXml_11_0_2 `xml:"virtual-link,omitempty"`
	Misc           []generic.Xml                                `xml:",any"`
}
type vrfOspfv3AreaTypeXml_11_0_2 struct {
	Normal *vrfOspfv3AreaTypeNormalXml_11_0_2 `xml:"normal,omitempty"`
	Stub   *vrfOspfv3AreaTypeStubXml_11_0_2   `xml:"stub,omitempty"`
	Nssa   *vrfOspfv3AreaTypeNssaXml_11_0_2   `xml:"nssa,omitempty"`
	Misc   []generic.Xml                      `xml:",any"`
}
type vrfOspfv3AreaTypeNormalXml_11_0_2 struct {
	Abr  *vrfOspfv3AreaTypeNormalAbrXml_11_0_2 `xml:"abr,omitempty"`
	Misc []generic.Xml                         `xml:",any"`
}
type vrfOspfv3AreaTypeNormalAbrXml_11_0_2 struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeStubXml_11_0_2 struct {
	NoSummary          *string                             `xml:"no-summary,omitempty"`
	Abr                *vrfOspfv3AreaTypeStubAbrXml_11_0_2 `xml:"abr,omitempty"`
	DefaultRouteMetric *int64                              `xml:"default-route-metric,omitempty"`
	Misc               []generic.Xml                       `xml:",any"`
}
type vrfOspfv3AreaTypeStubAbrXml_11_0_2 struct {
	ImportList         *string       `xml:"import-list,omitempty"`
	ExportList         *string       `xml:"export-list,omitempty"`
	InboundFilterList  *string       `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string       `xml:"outbound-filter-list,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeNssaXml_11_0_2 struct {
	NoSummary                   *string                                                     `xml:"no-summary,omitempty"`
	DefaultInformationOriginate *vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2 `xml:"default-information-originate,omitempty"`
	Abr                         *vrfOspfv3AreaTypeNssaAbrXml_11_0_2                         `xml:"abr,omitempty"`
	Misc                        []generic.Xml                                               `xml:",any"`
}
type vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2 struct {
	Metric     *int64        `xml:"metric,omitempty"`
	MetricType *string       `xml:"metric-type,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaTypeNssaAbrXml_11_0_2 struct {
	ImportList         *string                                                  `xml:"import-list,omitempty"`
	ExportList         *string                                                  `xml:"export-list,omitempty"`
	InboundFilterList  *string                                                  `xml:"inbound-filter-list,omitempty"`
	OutboundFilterList *string                                                  `xml:"outbound-filter-list,omitempty"`
	NssaExtRange       *vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2 `xml:"nssa-ext-range,omitempty"`
	Misc               []generic.Xml                                            `xml:",any"`
}
type vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaRangeContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaRangeXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaRangeXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Advertise *string       `xml:"advertise,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaInterfaceXml_11_0_2 struct {
	XMLName        xml.Name                                  `xml:"entry"`
	Name           string                                    `xml:"name,attr"`
	Enable         *string                                   `xml:"enable,omitempty"`
	MtuIgnore      *string                                   `xml:"mtu-ignore,omitempty"`
	Passive        *string                                   `xml:"passive,omitempty"`
	Priority       *int64                                    `xml:"priority,omitempty"`
	Metric         *int64                                    `xml:"metric,omitempty"`
	InstanceId     *int64                                    `xml:"instance-id,omitempty"`
	Authentication *string                                   `xml:"authentication,omitempty"`
	Timing         *string                                   `xml:"timing,omitempty"`
	LinkType       *vrfOspfv3AreaInterfaceLinkTypeXml_11_0_2 `xml:"link-type,omitempty"`
	Bfd            *vrfOspfv3AreaInterfaceBfdXml_11_0_2      `xml:"bfd,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeXml_11_0_2 struct {
	Broadcast *vrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2 `xml:"broadcast,omitempty"`
	P2p       *vrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2       `xml:"p2p,omitempty"`
	P2mp      *vrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                                      `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2 struct {
	Neighbor *vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2 `xml:"neighbor,omitempty"`
	Misc     []generic.Xml                                                  `xml:",any"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Priority *int64        `xml:"priority,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaInterfaceBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfOspfv3AreaVirtualLinkContainerXml_11_0_2 struct {
	Entries []vrfOspfv3AreaVirtualLinkXml_11_0_2 `xml:"entry"`
}
type vrfOspfv3AreaVirtualLinkXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	NeighborId     *string       `xml:"neighbor-id,omitempty"`
	TransitAreaId  *string       `xml:"transit-area-id,omitempty"`
	Enable         *string       `xml:"enable,omitempty"`
	InstanceId     *int64        `xml:"instance-id,omitempty"`
	Timing         *string       `xml:"timing,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfEcmpXml_11_0_2 struct {
	Enable           *string                     `xml:"enable,omitempty"`
	MaxPath          *int64                      `xml:"max-path,omitempty"`
	SymmetricReturn  *string                     `xml:"symmetric-return,omitempty"`
	StrictSourcePath *string                     `xml:"strict-source-path,omitempty"`
	Algorithm        *vrfEcmpAlgorithmXml_11_0_2 `xml:"algorithm,omitempty"`
	Misc             []generic.Xml               `xml:",any"`
}
type vrfEcmpAlgorithmXml_11_0_2 struct {
	IpModulo           *vrfEcmpAlgorithmIpModuloXml_11_0_2           `xml:"ip-modulo,omitempty"`
	IpHash             *vrfEcmpAlgorithmIpHashXml_11_0_2             `xml:"ip-hash,omitempty"`
	WeightedRoundRobin *vrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2 `xml:"weighted-round-robin,omitempty"`
	BalancedRoundRobin *vrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2 `xml:"balanced-round-robin,omitempty"`
	Misc               []generic.Xml                                 `xml:",any"`
}
type vrfEcmpAlgorithmIpModuloXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmIpHashXml_11_0_2 struct {
	SrcOnly  *string       `xml:"src-only,omitempty"`
	UsePort  *string       `xml:"use-port,omitempty"`
	HashSeed *int64        `xml:"hash-seed,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2 struct {
	Interface *vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml_11_0_2 `xml:"interface,omitempty"`
	Misc      []generic.Xml                                                   `xml:",any"`
}
type vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml_11_0_2 struct {
	Entries []vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Weight  *int64        `xml:"weight,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type vrfMulticastXml_11_0_2 struct {
	Enable      *string                                     `xml:"enable,omitempty"`
	StaticRoute *vrfMulticastStaticRouteContainerXml_11_0_2 `xml:"static-route,omitempty"`
	Pim         *vrfMulticastPimXml_11_0_2                  `xml:"pim,omitempty"`
	Igmp        *vrfMulticastIgmpXml_11_0_2                 `xml:"igmp,omitempty"`
	Msdp        *vrfMulticastMsdpXml_11_0_2                 `xml:"msdp,omitempty"`
	Misc        []generic.Xml                               `xml:",any"`
}
type vrfMulticastStaticRouteContainerXml_11_0_2 struct {
	Entries []vrfMulticastStaticRouteXml_11_0_2 `xml:"entry"`
}
type vrfMulticastStaticRouteXml_11_0_2 struct {
	XMLName     xml.Name                                  `xml:"entry"`
	Name        string                                    `xml:"name,attr"`
	Destination *string                                   `xml:"destination,omitempty"`
	Interface   *string                                   `xml:"interface,omitempty"`
	Preference  *int64                                    `xml:"preference,omitempty"`
	Nexthop     *vrfMulticastStaticRouteNexthopXml_11_0_2 `xml:"nexthop,omitempty"`
	Misc        []generic.Xml                             `xml:",any"`
}
type vrfMulticastStaticRouteNexthopXml_11_0_2 struct {
	IpAddress *string       `xml:"ip-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimXml_11_0_2 struct {
	Enable          *string                                         `xml:"enable,omitempty"`
	RpfLookupMode   *string                                         `xml:"rpf-lookup-mode,omitempty"`
	RouteAgeoutTime *int64                                          `xml:"route-ageout-time,omitempty"`
	IfTimerGlobal   *string                                         `xml:"if-timer-global,omitempty"`
	GroupPermission *string                                         `xml:"group-permission,omitempty"`
	SsmAddressSpace *vrfMulticastPimSsmAddressSpaceXml_11_0_2       `xml:"ssm-address-space,omitempty"`
	Rp              *vrfMulticastPimRpXml_11_0_2                    `xml:"rp,omitempty"`
	SptThreshold    *vrfMulticastPimSptThresholdContainerXml_11_0_2 `xml:"spt-threshold,omitempty"`
	Interface       *vrfMulticastPimInterfaceContainerXml_11_0_2    `xml:"interface,omitempty"`
	Misc            []generic.Xml                                   `xml:",any"`
}
type vrfMulticastPimSsmAddressSpaceXml_11_0_2 struct {
	GroupList *string       `xml:"group-list,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpXml_11_0_2 struct {
	LocalRp    *vrfMulticastPimRpLocalRpXml_11_0_2             `xml:"local-rp,omitempty"`
	ExternalRp *vrfMulticastPimRpExternalRpContainerXml_11_0_2 `xml:"external-rp,omitempty"`
	Misc       []generic.Xml                                   `xml:",any"`
}
type vrfMulticastPimRpLocalRpXml_11_0_2 struct {
	StaticRp    *vrfMulticastPimRpLocalRpStaticRpXml_11_0_2    `xml:"static-rp,omitempty"`
	CandidateRp *vrfMulticastPimRpLocalRpCandidateRpXml_11_0_2 `xml:"candidate-rp,omitempty"`
	Misc        []generic.Xml                                  `xml:",any"`
}
type vrfMulticastPimRpLocalRpStaticRpXml_11_0_2 struct {
	Interface *string       `xml:"interface,omitempty"`
	Address   *string       `xml:"address,omitempty"`
	Override  *string       `xml:"override,omitempty"`
	GroupList *string       `xml:"group-list,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpLocalRpCandidateRpXml_11_0_2 struct {
	Interface             *string       `xml:"interface,omitempty"`
	Address               *string       `xml:"address,omitempty"`
	Priority              *int64        `xml:"priority,omitempty"`
	AdvertisementInterval *int64        `xml:"advertisement-interval,omitempty"`
	GroupList             *string       `xml:"group-list,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type vrfMulticastPimRpExternalRpContainerXml_11_0_2 struct {
	Entries []vrfMulticastPimRpExternalRpXml_11_0_2 `xml:"entry"`
}
type vrfMulticastPimRpExternalRpXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	GroupList *string       `xml:"group-list,omitempty"`
	Override  *string       `xml:"override,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimSptThresholdContainerXml_11_0_2 struct {
	Entries []vrfMulticastPimSptThresholdXml_11_0_2 `xml:"entry"`
}
type vrfMulticastPimSptThresholdXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Threshold *string       `xml:"threshold,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastPimInterfaceContainerXml_11_0_2 struct {
	Entries []vrfMulticastPimInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfMulticastPimInterfaceXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	DrPriority     *int64        `xml:"dr-priority,omitempty"`
	SendBsm        *string       `xml:"send-bsm,omitempty"`
	IfTimer        *string       `xml:"if-timer,omitempty"`
	NeighborFilter *string       `xml:"neighbor-filter,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type vrfMulticastIgmpXml_11_0_2 struct {
	Enable  *string                                    `xml:"enable,omitempty"`
	Dynamic *vrfMulticastIgmpDynamicXml_11_0_2         `xml:"dynamic,omitempty"`
	Static  *vrfMulticastIgmpStaticContainerXml_11_0_2 `xml:"static,omitempty"`
	Misc    []generic.Xml                              `xml:",any"`
}
type vrfMulticastIgmpDynamicXml_11_0_2 struct {
	Interface *vrfMulticastIgmpDynamicInterfaceContainerXml_11_0_2 `xml:"interface,omitempty"`
	Misc      []generic.Xml                                        `xml:",any"`
}
type vrfMulticastIgmpDynamicInterfaceContainerXml_11_0_2 struct {
	Entries []vrfMulticastIgmpDynamicInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfMulticastIgmpDynamicInterfaceXml_11_0_2 struct {
	XMLName             xml.Name      `xml:"entry"`
	Name                string        `xml:"name,attr"`
	Version             *string       `xml:"version,omitempty"`
	Robustness          *string       `xml:"robustness,omitempty"`
	GroupFilter         *string       `xml:"group-filter,omitempty"`
	MaxGroups           *string       `xml:"max-groups,omitempty"`
	MaxSources          *string       `xml:"max-sources,omitempty"`
	QueryProfile        *string       `xml:"query-profile,omitempty"`
	RouterAlertPolicing *string       `xml:"router-alert-policing,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type vrfMulticastIgmpStaticContainerXml_11_0_2 struct {
	Entries []vrfMulticastIgmpStaticXml_11_0_2 `xml:"entry"`
}
type vrfMulticastIgmpStaticXml_11_0_2 struct {
	XMLName       xml.Name      `xml:"entry"`
	Name          string        `xml:"name,attr"`
	Interface     *string       `xml:"interface,omitempty"`
	GroupAddress  *string       `xml:"group-address,omitempty"`
	SourceAddress *string       `xml:"source-address,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpXml_11_0_2 struct {
	Enable               *string                                  `xml:"enable,omitempty"`
	GlobalTimer          *string                                  `xml:"global-timer,omitempty"`
	GlobalAuthentication *string                                  `xml:"global-authentication,omitempty"`
	OriginatorId         *vrfMulticastMsdpOriginatorIdXml_11_0_2  `xml:"originator-id,omitempty"`
	Peer                 *vrfMulticastMsdpPeerContainerXml_11_0_2 `xml:"peer,omitempty"`
	Misc                 []generic.Xml                            `xml:",any"`
}
type vrfMulticastMsdpOriginatorIdXml_11_0_2 struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpPeerContainerXml_11_0_2 struct {
	Entries []vrfMulticastMsdpPeerXml_11_0_2 `xml:"entry"`
}
type vrfMulticastMsdpPeerXml_11_0_2 struct {
	XMLName          xml.Name                                    `xml:"entry"`
	Name             string                                      `xml:"name,attr"`
	Enable           *string                                     `xml:"enable,omitempty"`
	PeerAs           *string                                     `xml:"peer-as,omitempty"`
	Authentication   *string                                     `xml:"authentication,omitempty"`
	MaxSa            *int64                                      `xml:"max-sa,omitempty"`
	InboundSaFilter  *string                                     `xml:"inbound-sa-filter,omitempty"`
	OutboundSaFilter *string                                     `xml:"outbound-sa-filter,omitempty"`
	LocalAddress     *vrfMulticastMsdpPeerLocalAddressXml_11_0_2 `xml:"local-address,omitempty"`
	PeerAddress      *vrfMulticastMsdpPeerPeerAddressXml_11_0_2  `xml:"peer-address,omitempty"`
	Misc             []generic.Xml                               `xml:",any"`
}
type vrfMulticastMsdpPeerLocalAddressXml_11_0_2 struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type vrfMulticastMsdpPeerPeerAddressXml_11_0_2 struct {
	Ip   *string       `xml:"ip,omitempty"`
	Fqdn *string       `xml:"fqdn,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type vrfRipXml_11_0_2 struct {
	Enable                       *string                                       `xml:"enable,omitempty"`
	DefaultInformationOriginate  *string                                       `xml:"default-information-originate,omitempty"`
	GlobalTimer                  *string                                       `xml:"global-timer,omitempty"`
	AuthProfile                  *string                                       `xml:"auth-profile,omitempty"`
	RedistributionProfile        *string                                       `xml:"redistribution-profile,omitempty"`
	GlobalBfd                    *vrfRipGlobalBfdXml_11_0_2                    `xml:"global-bfd,omitempty"`
	GlobalInboundDistributeList  *vrfRipGlobalInboundDistributeListXml_11_0_2  `xml:"global-inbound-distribute-list,omitempty"`
	GlobalOutboundDistributeList *vrfRipGlobalOutboundDistributeListXml_11_0_2 `xml:"global-outbound-distribute-list,omitempty"`
	Interface                    *vrfRipInterfaceContainerXml_11_0_2           `xml:"interface,omitempty"`
	Misc                         []generic.Xml                                 `xml:",any"`
}
type vrfRipGlobalBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRipGlobalInboundDistributeListXml_11_0_2 struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipGlobalOutboundDistributeListXml_11_0_2 struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipInterfaceContainerXml_11_0_2 struct {
	Entries []vrfRipInterfaceXml_11_0_2 `xml:"entry"`
}
type vrfRipInterfaceXml_11_0_2 struct {
	XMLName                         xml.Name                                                  `xml:"entry"`
	Name                            string                                                    `xml:"name,attr"`
	Enable                          *string                                                   `xml:"enable,omitempty"`
	Mode                            *string                                                   `xml:"mode,omitempty"`
	SplitHorizon                    *string                                                   `xml:"split-horizon,omitempty"`
	Authentication                  *string                                                   `xml:"authentication,omitempty"`
	Bfd                             *vrfRipInterfaceBfdXml_11_0_2                             `xml:"bfd,omitempty"`
	InterfaceInboundDistributeList  *vrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2  `xml:"interface-inbound-distribute-list,omitempty"`
	InterfaceOutboundDistributeList *vrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2 `xml:"interface-outbound-distribute-list,omitempty"`
	Misc                            []generic.Xml                                             `xml:",any"`
}
type vrfRipInterfaceBfdXml_11_0_2 struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type vrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2 struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Metric     *int64        `xml:"metric,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type vrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2 struct {
	AccessList *string       `xml:"access-list,omitempty"`
	Metric     *int64        `xml:"metric,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Vrf != nil {
		var objs []vrfXml
		for _, elt := range s.Vrf {
			var obj vrfXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vrf = &vrfContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var vrfVal []Vrf
	if o.Vrf != nil {
		for _, elt := range o.Vrf.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vrfVal = append(vrfVal, *obj)
		}
	}

	result := &Entry{
		Name: o.Name,
		Vrf:  vrfVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfXml) MarshalFromObject(s Vrf) {
	o.Name = s.Name
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.AdminDists != nil {
		var obj vrfAdminDistsXml
		obj.MarshalFromObject(*s.AdminDists)
		o.AdminDists = &obj
	}
	if s.RibFilter != nil {
		var obj vrfRibFilterXml
		obj.MarshalFromObject(*s.RibFilter)
		o.RibFilter = &obj
	}
	if s.Bgp != nil {
		var obj vrfBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.RoutingTable != nil {
		var obj vrfRoutingTableXml
		obj.MarshalFromObject(*s.RoutingTable)
		o.RoutingTable = &obj
	}
	if s.Ospf != nil {
		var obj vrfOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj vrfOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Ecmp != nil {
		var obj vrfEcmpXml
		obj.MarshalFromObject(*s.Ecmp)
		o.Ecmp = &obj
	}
	if s.Multicast != nil {
		var obj vrfMulticastXml
		obj.MarshalFromObject(*s.Multicast)
		o.Multicast = &obj
	}
	if s.Rip != nil {
		var obj vrfRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
}

func (o vrfXml) UnmarshalToObject() (*Vrf, error) {
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var adminDistsVal *VrfAdminDists
	if o.AdminDists != nil {
		obj, err := o.AdminDists.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adminDistsVal = obj
	}
	var ribFilterVal *VrfRibFilter
	if o.RibFilter != nil {
		obj, err := o.RibFilter.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribFilterVal = obj
	}
	var bgpVal *VrfBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var routingTableVal *VrfRoutingTable
	if o.RoutingTable != nil {
		obj, err := o.RoutingTable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingTableVal = obj
	}
	var ospfVal *VrfOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *VrfOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ecmpVal *VrfEcmp
	if o.Ecmp != nil {
		obj, err := o.Ecmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ecmpVal = obj
	}
	var multicastVal *VrfMulticast
	if o.Multicast != nil {
		obj, err := o.Multicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastVal = obj
	}
	var ripVal *VrfRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Vrf{
		Name:         o.Name,
		Interface:    interfaceVal,
		AdminDists:   adminDistsVal,
		RibFilter:    ribFilterVal,
		Bgp:          bgpVal,
		RoutingTable: routingTableVal,
		Ospf:         ospfVal,
		Ospfv3:       ospfv3Val,
		Ecmp:         ecmpVal,
		Multicast:    multicastVal,
		Rip:          ripVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *vrfAdminDistsXml) MarshalFromObject(s VrfAdminDists) {
	o.Static = s.Static
	o.StaticIpv6 = s.StaticIpv6
	o.OspfInter = s.OspfInter
	o.OspfIntra = s.OspfIntra
	o.OspfExt = s.OspfExt
	o.Ospfv3Inter = s.Ospfv3Inter
	o.Ospfv3Intra = s.Ospfv3Intra
	o.Ospfv3Ext = s.Ospfv3Ext
	o.BgpInternal = s.BgpInternal
	o.BgpExternal = s.BgpExternal
	o.BgpLocal = s.BgpLocal
	o.Rip = s.Rip
	o.Misc = s.Misc
}

func (o vrfAdminDistsXml) UnmarshalToObject() (*VrfAdminDists, error) {

	result := &VrfAdminDists{
		Static:      o.Static,
		StaticIpv6:  o.StaticIpv6,
		OspfInter:   o.OspfInter,
		OspfIntra:   o.OspfIntra,
		OspfExt:     o.OspfExt,
		Ospfv3Inter: o.Ospfv3Inter,
		Ospfv3Intra: o.Ospfv3Intra,
		Ospfv3Ext:   o.Ospfv3Ext,
		BgpInternal: o.BgpInternal,
		BgpExternal: o.BgpExternal,
		BgpLocal:    o.BgpLocal,
		Rip:         o.Rip,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterXml) MarshalFromObject(s VrfRibFilter) {
	if s.Ipv4 != nil {
		var obj vrfRibFilterIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfRibFilterIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterXml) UnmarshalToObject() (*VrfRibFilter, error) {
	var ipv4Val *VrfRibFilterIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfRibFilterIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfRibFilter{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4Xml) MarshalFromObject(s VrfRibFilterIpv4) {
	if s.Static != nil {
		var obj vrfRibFilterIpv4StaticXml
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	if s.Bgp != nil {
		var obj vrfRibFilterIpv4BgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj vrfRibFilterIpv4OspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Rip != nil {
		var obj vrfRibFilterIpv4RipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4Xml) UnmarshalToObject() (*VrfRibFilterIpv4, error) {
	var staticVal *VrfRibFilterIpv4Static
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}
	var bgpVal *VrfRibFilterIpv4Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *VrfRibFilterIpv4Ospf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ripVal *VrfRibFilterIpv4Rip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &VrfRibFilterIpv4{
		Static: staticVal,
		Bgp:    bgpVal,
		Ospf:   ospfVal,
		Rip:    ripVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4StaticXml) MarshalFromObject(s VrfRibFilterIpv4Static) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4StaticXml) UnmarshalToObject() (*VrfRibFilterIpv4Static, error) {

	result := &VrfRibFilterIpv4Static{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4BgpXml) MarshalFromObject(s VrfRibFilterIpv4Bgp) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4BgpXml) UnmarshalToObject() (*VrfRibFilterIpv4Bgp, error) {

	result := &VrfRibFilterIpv4Bgp{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4OspfXml) MarshalFromObject(s VrfRibFilterIpv4Ospf) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4OspfXml) UnmarshalToObject() (*VrfRibFilterIpv4Ospf, error) {

	result := &VrfRibFilterIpv4Ospf{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4RipXml) MarshalFromObject(s VrfRibFilterIpv4Rip) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4RipXml) UnmarshalToObject() (*VrfRibFilterIpv4Rip, error) {

	result := &VrfRibFilterIpv4Rip{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6Xml) MarshalFromObject(s VrfRibFilterIpv6) {
	if s.Static != nil {
		var obj vrfRibFilterIpv6StaticXml
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	if s.Bgp != nil {
		var obj vrfRibFilterIpv6BgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospfv3 != nil {
		var obj vrfRibFilterIpv6Ospfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6Xml) UnmarshalToObject() (*VrfRibFilterIpv6, error) {
	var staticVal *VrfRibFilterIpv6Static
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}
	var bgpVal *VrfRibFilterIpv6Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfv3Val *VrfRibFilterIpv6Ospfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}

	result := &VrfRibFilterIpv6{
		Static: staticVal,
		Bgp:    bgpVal,
		Ospfv3: ospfv3Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6StaticXml) MarshalFromObject(s VrfRibFilterIpv6Static) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6StaticXml) UnmarshalToObject() (*VrfRibFilterIpv6Static, error) {

	result := &VrfRibFilterIpv6Static{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6BgpXml) MarshalFromObject(s VrfRibFilterIpv6Bgp) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6BgpXml) UnmarshalToObject() (*VrfRibFilterIpv6Bgp, error) {

	result := &VrfRibFilterIpv6Bgp{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6Ospfv3Xml) MarshalFromObject(s VrfRibFilterIpv6Ospfv3) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6Ospfv3Xml) UnmarshalToObject() (*VrfRibFilterIpv6Ospfv3, error) {

	result := &VrfRibFilterIpv6Ospfv3{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfBgpXml) MarshalFromObject(s VrfBgp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RouterId = s.RouterId
	o.LocalAs = s.LocalAs
	o.InstallRoute = util.YesNo(s.InstallRoute, nil)
	o.EnforceFirstAs = util.YesNo(s.EnforceFirstAs, nil)
	o.FastExternalFailover = util.YesNo(s.FastExternalFailover, nil)
	o.EcmpMultiAs = util.YesNo(s.EcmpMultiAs, nil)
	o.DefaultLocalPreference = s.DefaultLocalPreference
	o.GracefulShutdown = util.YesNo(s.GracefulShutdown, nil)
	o.AlwaysAdvertiseNetworkRoute = util.YesNo(s.AlwaysAdvertiseNetworkRoute, nil)
	if s.Med != nil {
		var obj vrfBgpMedXml
		obj.MarshalFromObject(*s.Med)
		o.Med = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfBgpGracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.GlobalBfd != nil {
		var obj vrfBgpGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.RedistributionProfile != nil {
		var obj vrfBgpRedistributionProfileXml
		obj.MarshalFromObject(*s.RedistributionProfile)
		o.RedistributionProfile = &obj
	}
	if s.AdvertiseNetwork != nil {
		var obj vrfBgpAdvertiseNetworkXml
		obj.MarshalFromObject(*s.AdvertiseNetwork)
		o.AdvertiseNetwork = &obj
	}
	if s.PeerGroup != nil {
		var objs []vrfBgpPeerGroupXml
		for _, elt := range s.PeerGroup {
			var obj vrfBgpPeerGroupXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.PeerGroup = &vrfBgpPeerGroupContainerXml{Entries: objs}
	}
	if s.AggregateRoutes != nil {
		var objs []vrfBgpAggregateRoutesXml
		for _, elt := range s.AggregateRoutes {
			var obj vrfBgpAggregateRoutesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AggregateRoutes = &vrfBgpAggregateRoutesContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpXml) UnmarshalToObject() (*VrfBgp, error) {
	var medVal *VrfBgpMed
	if o.Med != nil {
		obj, err := o.Med.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		medVal = obj
	}
	var gracefulRestartVal *VrfBgpGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var globalBfdVal *VrfBgpGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var redistributionProfileVal *VrfBgpRedistributionProfile
	if o.RedistributionProfile != nil {
		obj, err := o.RedistributionProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redistributionProfileVal = obj
	}
	var advertiseNetworkVal *VrfBgpAdvertiseNetwork
	if o.AdvertiseNetwork != nil {
		obj, err := o.AdvertiseNetwork.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseNetworkVal = obj
	}
	var peerGroupVal []VrfBgpPeerGroup
	if o.PeerGroup != nil {
		for _, elt := range o.PeerGroup.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerGroupVal = append(peerGroupVal, *obj)
		}
	}
	var aggregateRoutesVal []VrfBgpAggregateRoutes
	if o.AggregateRoutes != nil {
		for _, elt := range o.AggregateRoutes.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			aggregateRoutesVal = append(aggregateRoutesVal, *obj)
		}
	}

	result := &VrfBgp{
		Enable:                      util.AsBool(o.Enable, nil),
		RouterId:                    o.RouterId,
		LocalAs:                     o.LocalAs,
		InstallRoute:                util.AsBool(o.InstallRoute, nil),
		EnforceFirstAs:              util.AsBool(o.EnforceFirstAs, nil),
		FastExternalFailover:        util.AsBool(o.FastExternalFailover, nil),
		EcmpMultiAs:                 util.AsBool(o.EcmpMultiAs, nil),
		DefaultLocalPreference:      o.DefaultLocalPreference,
		GracefulShutdown:            util.AsBool(o.GracefulShutdown, nil),
		AlwaysAdvertiseNetworkRoute: util.AsBool(o.AlwaysAdvertiseNetworkRoute, nil),
		Med:                         medVal,
		GracefulRestart:             gracefulRestartVal,
		GlobalBfd:                   globalBfdVal,
		RedistributionProfile:       redistributionProfileVal,
		AdvertiseNetwork:            advertiseNetworkVal,
		PeerGroup:                   peerGroupVal,
		AggregateRoutes:             aggregateRoutesVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfBgpMedXml) MarshalFromObject(s VrfBgpMed) {
	o.AlwaysCompareMed = util.YesNo(s.AlwaysCompareMed, nil)
	o.DeterministicMedComparison = util.YesNo(s.DeterministicMedComparison, nil)
	o.Misc = s.Misc
}

func (o vrfBgpMedXml) UnmarshalToObject() (*VrfBgpMed, error) {

	result := &VrfBgpMed{
		AlwaysCompareMed:           util.AsBool(o.AlwaysCompareMed, nil),
		DeterministicMedComparison: util.AsBool(o.DeterministicMedComparison, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *vrfBgpGracefulRestartXml) MarshalFromObject(s VrfBgpGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.StaleRouteTime = s.StaleRouteTime
	o.MaxPeerRestartTime = s.MaxPeerRestartTime
	o.LocalRestartTime = s.LocalRestartTime
	o.Misc = s.Misc
}

func (o vrfBgpGracefulRestartXml) UnmarshalToObject() (*VrfBgpGracefulRestart, error) {

	result := &VrfBgpGracefulRestart{
		Enable:             util.AsBool(o.Enable, nil),
		StaleRouteTime:     o.StaleRouteTime,
		MaxPeerRestartTime: o.MaxPeerRestartTime,
		LocalRestartTime:   o.LocalRestartTime,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfBgpGlobalBfdXml) MarshalFromObject(s VrfBgpGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfBgpGlobalBfdXml) UnmarshalToObject() (*VrfBgpGlobalBfd, error) {

	result := &VrfBgpGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileXml) MarshalFromObject(s VrfBgpRedistributionProfile) {
	if s.Ipv4 != nil {
		var obj vrfBgpRedistributionProfileIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpRedistributionProfileIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileXml) UnmarshalToObject() (*VrfBgpRedistributionProfile, error) {
	var ipv4Val *VrfBgpRedistributionProfileIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpRedistributionProfileIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpRedistributionProfile{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileIpv4Xml) MarshalFromObject(s VrfBgpRedistributionProfileIpv4) {
	o.Unicast = s.Unicast
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileIpv4Xml) UnmarshalToObject() (*VrfBgpRedistributionProfileIpv4, error) {

	result := &VrfBgpRedistributionProfileIpv4{
		Unicast: o.Unicast,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileIpv6Xml) MarshalFromObject(s VrfBgpRedistributionProfileIpv6) {
	o.Unicast = s.Unicast
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileIpv6Xml) UnmarshalToObject() (*VrfBgpRedistributionProfileIpv6, error) {

	result := &VrfBgpRedistributionProfileIpv6{
		Unicast: o.Unicast,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkXml) MarshalFromObject(s VrfBgpAdvertiseNetwork) {
	if s.Ipv4 != nil {
		var obj vrfBgpAdvertiseNetworkIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpAdvertiseNetworkIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkXml) UnmarshalToObject() (*VrfBgpAdvertiseNetwork, error) {
	var ipv4Val *VrfBgpAdvertiseNetworkIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpAdvertiseNetworkIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpAdvertiseNetwork{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv4Xml) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv4) {
	if s.Network != nil {
		var objs []vrfBgpAdvertiseNetworkIpv4NetworkXml
		for _, elt := range s.Network {
			var obj vrfBgpAdvertiseNetworkIpv4NetworkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Network = &vrfBgpAdvertiseNetworkIpv4NetworkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv4Xml) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv4, error) {
	var networkVal []VrfBgpAdvertiseNetworkIpv4Network
	if o.Network != nil {
		for _, elt := range o.Network.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			networkVal = append(networkVal, *obj)
		}
	}

	result := &VrfBgpAdvertiseNetworkIpv4{
		Network: networkVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv4NetworkXml) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv4Network) {
	o.Name = s.Name
	o.Unicast = util.YesNo(s.Unicast, nil)
	o.Multicast = util.YesNo(s.Multicast, nil)
	o.Backdoor = util.YesNo(s.Backdoor, nil)
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv4NetworkXml) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv4Network, error) {

	result := &VrfBgpAdvertiseNetworkIpv4Network{
		Name:      o.Name,
		Unicast:   util.AsBool(o.Unicast, nil),
		Multicast: util.AsBool(o.Multicast, nil),
		Backdoor:  util.AsBool(o.Backdoor, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv6Xml) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv6) {
	if s.Network != nil {
		var objs []vrfBgpAdvertiseNetworkIpv6NetworkXml
		for _, elt := range s.Network {
			var obj vrfBgpAdvertiseNetworkIpv6NetworkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Network = &vrfBgpAdvertiseNetworkIpv6NetworkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv6Xml) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv6, error) {
	var networkVal []VrfBgpAdvertiseNetworkIpv6Network
	if o.Network != nil {
		for _, elt := range o.Network.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			networkVal = append(networkVal, *obj)
		}
	}

	result := &VrfBgpAdvertiseNetworkIpv6{
		Network: networkVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv6NetworkXml) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv6Network) {
	o.Name = s.Name
	o.Unicast = util.YesNo(s.Unicast, nil)
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv6NetworkXml) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv6Network, error) {

	result := &VrfBgpAdvertiseNetworkIpv6Network{
		Name:    o.Name,
		Unicast: util.AsBool(o.Unicast, nil),
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupXml) MarshalFromObject(s VrfBgpPeerGroup) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Type != nil {
		var obj vrfBgpPeerGroupTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.AddressFamily != nil {
		var obj vrfBgpPeerGroupAddressFamilyXml
		obj.MarshalFromObject(*s.AddressFamily)
		o.AddressFamily = &obj
	}
	if s.FilteringProfile != nil {
		var obj vrfBgpPeerGroupFilteringProfileXml
		obj.MarshalFromObject(*s.FilteringProfile)
		o.FilteringProfile = &obj
	}
	if s.ConnectionOptions != nil {
		var obj vrfBgpPeerGroupConnectionOptionsXml
		obj.MarshalFromObject(*s.ConnectionOptions)
		o.ConnectionOptions = &obj
	}
	if s.Peer != nil {
		var objs []vrfBgpPeerGroupPeerXml
		for _, elt := range s.Peer {
			var obj vrfBgpPeerGroupPeerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Peer = &vrfBgpPeerGroupPeerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupXml) UnmarshalToObject() (*VrfBgpPeerGroup, error) {
	var typeVal *VrfBgpPeerGroupType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var addressFamilyVal *VrfBgpPeerGroupAddressFamily
	if o.AddressFamily != nil {
		obj, err := o.AddressFamily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressFamilyVal = obj
	}
	var filteringProfileVal *VrfBgpPeerGroupFilteringProfile
	if o.FilteringProfile != nil {
		obj, err := o.FilteringProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filteringProfileVal = obj
	}
	var connectionOptionsVal *VrfBgpPeerGroupConnectionOptions
	if o.ConnectionOptions != nil {
		obj, err := o.ConnectionOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectionOptionsVal = obj
	}
	var peerVal []VrfBgpPeerGroupPeer
	if o.Peer != nil {
		for _, elt := range o.Peer.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerVal = append(peerVal, *obj)
		}
	}

	result := &VrfBgpPeerGroup{
		Name:              o.Name,
		Enable:            util.AsBool(o.Enable, nil),
		Type:              typeVal,
		AddressFamily:     addressFamilyVal,
		FilteringProfile:  filteringProfileVal,
		ConnectionOptions: connectionOptionsVal,
		Peer:              peerVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeXml) MarshalFromObject(s VrfBgpPeerGroupType) {
	if s.Ibgp != nil {
		var obj vrfBgpPeerGroupTypeIbgpXml
		obj.MarshalFromObject(*s.Ibgp)
		o.Ibgp = &obj
	}
	if s.Ebgp != nil {
		var obj vrfBgpPeerGroupTypeEbgpXml
		obj.MarshalFromObject(*s.Ebgp)
		o.Ebgp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeXml) UnmarshalToObject() (*VrfBgpPeerGroupType, error) {
	var ibgpVal *VrfBgpPeerGroupTypeIbgp
	if o.Ibgp != nil {
		obj, err := o.Ibgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ibgpVal = obj
	}
	var ebgpVal *VrfBgpPeerGroupTypeEbgp
	if o.Ebgp != nil {
		obj, err := o.Ebgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ebgpVal = obj
	}

	result := &VrfBgpPeerGroupType{
		Ibgp: ibgpVal,
		Ebgp: ebgpVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeIbgpXml) MarshalFromObject(s VrfBgpPeerGroupTypeIbgp) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeIbgpXml) UnmarshalToObject() (*VrfBgpPeerGroupTypeIbgp, error) {

	result := &VrfBgpPeerGroupTypeIbgp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeEbgpXml) MarshalFromObject(s VrfBgpPeerGroupTypeEbgp) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeEbgpXml) UnmarshalToObject() (*VrfBgpPeerGroupTypeEbgp, error) {

	result := &VrfBgpPeerGroupTypeEbgp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupAddressFamilyXml) MarshalFromObject(s VrfBgpPeerGroupAddressFamily) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupAddressFamilyXml) UnmarshalToObject() (*VrfBgpPeerGroupAddressFamily, error) {

	result := &VrfBgpPeerGroupAddressFamily{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupFilteringProfileXml) MarshalFromObject(s VrfBgpPeerGroupFilteringProfile) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupFilteringProfileXml) UnmarshalToObject() (*VrfBgpPeerGroupFilteringProfile, error) {

	result := &VrfBgpPeerGroupFilteringProfile{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupConnectionOptionsXml) MarshalFromObject(s VrfBgpPeerGroupConnectionOptions) {
	o.Timers = s.Timers
	o.Multihop = s.Multihop
	o.Authentication = s.Authentication
	o.Dampening = s.Dampening
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupConnectionOptionsXml) UnmarshalToObject() (*VrfBgpPeerGroupConnectionOptions, error) {

	result := &VrfBgpPeerGroupConnectionOptions{
		Timers:         o.Timers,
		Multihop:       o.Multihop,
		Authentication: o.Authentication,
		Dampening:      o.Dampening,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerXml) MarshalFromObject(s VrfBgpPeerGroupPeer) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.PeerAs = s.PeerAs
	o.EnableSenderSideLoopDetection = util.YesNo(s.EnableSenderSideLoopDetection, nil)
	if s.Inherit != nil {
		var obj vrfBgpPeerGroupPeerInheritXml
		obj.MarshalFromObject(*s.Inherit)
		o.Inherit = &obj
	}
	if s.LocalAddress != nil {
		var obj vrfBgpPeerGroupPeerLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj vrfBgpPeerGroupPeerPeerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	if s.ConnectionOptions != nil {
		var obj vrfBgpPeerGroupPeerConnectionOptionsXml
		obj.MarshalFromObject(*s.ConnectionOptions)
		o.ConnectionOptions = &obj
	}
	if s.Bfd != nil {
		var obj vrfBgpPeerGroupPeerBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerXml) UnmarshalToObject() (*VrfBgpPeerGroupPeer, error) {
	var inheritVal *VrfBgpPeerGroupPeerInherit
	if o.Inherit != nil {
		obj, err := o.Inherit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritVal = obj
	}
	var localAddressVal *VrfBgpPeerGroupPeerLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *VrfBgpPeerGroupPeerPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var connectionOptionsVal *VrfBgpPeerGroupPeerConnectionOptions
	if o.ConnectionOptions != nil {
		obj, err := o.ConnectionOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectionOptionsVal = obj
	}
	var bfdVal *VrfBgpPeerGroupPeerBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfBgpPeerGroupPeer{
		Name:                          o.Name,
		Enable:                        util.AsBool(o.Enable, nil),
		Passive:                       util.AsBool(o.Passive, nil),
		PeerAs:                        o.PeerAs,
		EnableSenderSideLoopDetection: util.AsBool(o.EnableSenderSideLoopDetection, nil),
		Inherit:                       inheritVal,
		LocalAddress:                  localAddressVal,
		PeerAddress:                   peerAddressVal,
		ConnectionOptions:             connectionOptionsVal,
		Bfd:                           bfdVal,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritXml) MarshalFromObject(s VrfBgpPeerGroupPeerInherit) {
	if s.Yes != nil {
		var obj vrfBgpPeerGroupPeerInheritYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	if s.No != nil {
		var obj vrfBgpPeerGroupPeerInheritNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerInherit, error) {
	var yesVal *VrfBgpPeerGroupPeerInheritYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}
	var noVal *VrfBgpPeerGroupPeerInheritNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}

	result := &VrfBgpPeerGroupPeerInherit{
		Yes:  yesVal,
		No:   noVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritYesXml) MarshalFromObject(s VrfBgpPeerGroupPeerInheritYes) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritYesXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritYes, error) {

	result := &VrfBgpPeerGroupPeerInheritYes{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoXml) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNo) {
	if s.AddressFamily != nil {
		var obj vrfBgpPeerGroupPeerInheritNoAddressFamilyXml
		obj.MarshalFromObject(*s.AddressFamily)
		o.AddressFamily = &obj
	}
	if s.FilteringProfile != nil {
		var obj vrfBgpPeerGroupPeerInheritNoFilteringProfileXml
		obj.MarshalFromObject(*s.FilteringProfile)
		o.FilteringProfile = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNo, error) {
	var addressFamilyVal *VrfBgpPeerGroupPeerInheritNoAddressFamily
	if o.AddressFamily != nil {
		obj, err := o.AddressFamily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressFamilyVal = obj
	}
	var filteringProfileVal *VrfBgpPeerGroupPeerInheritNoFilteringProfile
	if o.FilteringProfile != nil {
		obj, err := o.FilteringProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filteringProfileVal = obj
	}

	result := &VrfBgpPeerGroupPeerInheritNo{
		AddressFamily:    addressFamilyVal,
		FilteringProfile: filteringProfileVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoAddressFamilyXml) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNoAddressFamily) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoAddressFamilyXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNoAddressFamily, error) {

	result := &VrfBgpPeerGroupPeerInheritNoAddressFamily{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoFilteringProfileXml) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNoFilteringProfile) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoFilteringProfileXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNoFilteringProfile, error) {

	result := &VrfBgpPeerGroupPeerInheritNoFilteringProfile{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerLocalAddressXml) MarshalFromObject(s VrfBgpPeerGroupPeerLocalAddress) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerLocalAddressXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerLocalAddress, error) {

	result := &VrfBgpPeerGroupPeerLocalAddress{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerPeerAddressXml) MarshalFromObject(s VrfBgpPeerGroupPeerPeerAddress) {
	o.Ip = s.Ip
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerPeerAddressXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerPeerAddress, error) {

	result := &VrfBgpPeerGroupPeerPeerAddress{
		Ip:   o.Ip,
		Fqdn: o.Fqdn,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerConnectionOptionsXml) MarshalFromObject(s VrfBgpPeerGroupPeerConnectionOptions) {
	o.Timers = s.Timers
	o.Multihop = s.Multihop
	o.Authentication = s.Authentication
	o.Dampening = s.Dampening
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerConnectionOptionsXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerConnectionOptions, error) {

	result := &VrfBgpPeerGroupPeerConnectionOptions{
		Timers:         o.Timers,
		Multihop:       o.Multihop,
		Authentication: o.Authentication,
		Dampening:      o.Dampening,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerBfdXml) MarshalFromObject(s VrfBgpPeerGroupPeerBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerBfdXml) UnmarshalToObject() (*VrfBgpPeerGroupPeerBfd, error) {

	result := &VrfBgpPeerGroupPeerBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesXml) MarshalFromObject(s VrfBgpAggregateRoutes) {
	o.Name = s.Name
	o.Description = s.Description
	o.Enable = util.YesNo(s.Enable, nil)
	o.SummaryOnly = util.YesNo(s.SummaryOnly, nil)
	o.AsSet = util.YesNo(s.AsSet, nil)
	o.SameMed = util.YesNo(s.SameMed, nil)
	if s.Type != nil {
		var obj vrfBgpAggregateRoutesTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesXml) UnmarshalToObject() (*VrfBgpAggregateRoutes, error) {
	var typeVal *VrfBgpAggregateRoutesType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &VrfBgpAggregateRoutes{
		Name:        o.Name,
		Description: o.Description,
		Enable:      util.AsBool(o.Enable, nil),
		SummaryOnly: util.AsBool(o.SummaryOnly, nil),
		AsSet:       util.AsBool(o.AsSet, nil),
		SameMed:     util.AsBool(o.SameMed, nil),
		Type:        typeVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeXml) MarshalFromObject(s VrfBgpAggregateRoutesType) {
	if s.Ipv4 != nil {
		var obj vrfBgpAggregateRoutesTypeIpv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpAggregateRoutesTypeIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeXml) UnmarshalToObject() (*VrfBgpAggregateRoutesType, error) {
	var ipv4Val *VrfBgpAggregateRoutesTypeIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpAggregateRoutesTypeIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpAggregateRoutesType{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeIpv4Xml) MarshalFromObject(s VrfBgpAggregateRoutesTypeIpv4) {
	o.SummaryPrefix = s.SummaryPrefix
	o.SuppressMap = s.SuppressMap
	o.AttributeMap = s.AttributeMap
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeIpv4Xml) UnmarshalToObject() (*VrfBgpAggregateRoutesTypeIpv4, error) {

	result := &VrfBgpAggregateRoutesTypeIpv4{
		SummaryPrefix: o.SummaryPrefix,
		SuppressMap:   o.SuppressMap,
		AttributeMap:  o.AttributeMap,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeIpv6Xml) MarshalFromObject(s VrfBgpAggregateRoutesTypeIpv6) {
	o.SummaryPrefix = s.SummaryPrefix
	o.SuppressMap = s.SuppressMap
	o.AttributeMap = s.AttributeMap
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeIpv6Xml) UnmarshalToObject() (*VrfBgpAggregateRoutesTypeIpv6, error) {

	result := &VrfBgpAggregateRoutesTypeIpv6{
		SummaryPrefix: o.SummaryPrefix,
		SuppressMap:   o.SuppressMap,
		AttributeMap:  o.AttributeMap,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableXml) MarshalFromObject(s VrfRoutingTable) {
	if s.Ip != nil {
		var obj vrfRoutingTableIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfRoutingTableIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableXml) UnmarshalToObject() (*VrfRoutingTable, error) {
	var ipVal *VrfRoutingTableIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}
	var ipv6Val *VrfRoutingTableIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfRoutingTable{
		Ip:   ipVal,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpXml) MarshalFromObject(s VrfRoutingTableIp) {
	if s.StaticRoute != nil {
		var objs []vrfRoutingTableIpStaticRouteXml
		for _, elt := range s.StaticRoute {
			var obj vrfRoutingTableIpStaticRouteXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfRoutingTableIpStaticRouteContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpXml) UnmarshalToObject() (*VrfRoutingTableIp, error) {
	var staticRouteVal []VrfRoutingTableIpStaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}

	result := &VrfRoutingTableIp{
		StaticRoute: staticRouteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteXml) MarshalFromObject(s VrfRoutingTableIpStaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.AdminDist = s.AdminDist
	o.Metric = s.Metric
	if s.Nexthop != nil {
		var obj vrfRoutingTableIpStaticRouteNexthopXml
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	if s.Bfd != nil {
		var obj vrfRoutingTableIpStaticRouteBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.PathMonitor != nil {
		var obj vrfRoutingTableIpStaticRoutePathMonitorXml
		obj.MarshalFromObject(*s.PathMonitor)
		o.PathMonitor = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRoute, error) {
	var nexthopVal *VrfRoutingTableIpStaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}
	var bfdVal *VrfRoutingTableIpStaticRouteBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var pathMonitorVal *VrfRoutingTableIpStaticRoutePathMonitor
	if o.PathMonitor != nil {
		obj, err := o.PathMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pathMonitorVal = obj
	}

	result := &VrfRoutingTableIpStaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		AdminDist:   o.AdminDist,
		Metric:      o.Metric,
		Nexthop:     nexthopVal,
		Bfd:         bfdVal,
		PathMonitor: pathMonitorVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteNexthopXml) MarshalFromObject(s VrfRoutingTableIpStaticRouteNexthop) {
	if s.Discard != nil {
		var obj vrfRoutingTableIpStaticRouteNexthopDiscardXml
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	o.IpAddress = s.IpAddress
	o.NextLr = s.NextLr
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteNexthopXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteNexthop, error) {
	var discardVal *VrfRoutingTableIpStaticRouteNexthopDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}

	result := &VrfRoutingTableIpStaticRouteNexthop{
		Discard:   discardVal,
		IpAddress: o.IpAddress,
		NextLr:    o.NextLr,
		Fqdn:      o.Fqdn,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteNexthopDiscardXml) MarshalFromObject(s VrfRoutingTableIpStaticRouteNexthopDiscard) {
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteNexthopDiscardXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteNexthopDiscard, error) {

	result := &VrfRoutingTableIpStaticRouteNexthopDiscard{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteBfdXml) MarshalFromObject(s VrfRoutingTableIpStaticRouteBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteBfdXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteBfd, error) {

	result := &VrfRoutingTableIpStaticRouteBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRoutePathMonitorXml) MarshalFromObject(s VrfRoutingTableIpStaticRoutePathMonitor) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FailureCondition = s.FailureCondition
	o.HoldTime = s.HoldTime
	if s.MonitorDestinations != nil {
		var objs []vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml
		for _, elt := range s.MonitorDestinations {
			var obj vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MonitorDestinations = &vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRoutePathMonitorXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRoutePathMonitor, error) {
	var monitorDestinationsVal []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations
	if o.MonitorDestinations != nil {
		for _, elt := range o.MonitorDestinations.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			monitorDestinationsVal = append(monitorDestinationsVal, *obj)
		}
	}

	result := &VrfRoutingTableIpStaticRoutePathMonitor{
		Enable:              util.AsBool(o.Enable, nil),
		FailureCondition:    o.FailureCondition,
		HoldTime:            o.HoldTime,
		MonitorDestinations: monitorDestinationsVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml) MarshalFromObject(s VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Source = s.Source
	o.Destination = s.Destination
	o.Interval = s.Interval
	o.Count = s.Count
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml) UnmarshalToObject() (*VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations, error) {

	result := &VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{
		Name:        o.Name,
		Enable:      util.AsBool(o.Enable, nil),
		Source:      o.Source,
		Destination: o.Destination,
		Interval:    o.Interval,
		Count:       o.Count,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6Xml) MarshalFromObject(s VrfRoutingTableIpv6) {
	if s.StaticRoute != nil {
		var objs []vrfRoutingTableIpv6StaticRouteXml
		for _, elt := range s.StaticRoute {
			var obj vrfRoutingTableIpv6StaticRouteXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfRoutingTableIpv6StaticRouteContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6Xml) UnmarshalToObject() (*VrfRoutingTableIpv6, error) {
	var staticRouteVal []VrfRoutingTableIpv6StaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}

	result := &VrfRoutingTableIpv6{
		StaticRoute: staticRouteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.AdminDist = s.AdminDist
	o.Metric = s.Metric
	if s.Nexthop != nil {
		var obj vrfRoutingTableIpv6StaticRouteNexthopXml
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	if s.Bfd != nil {
		var obj vrfRoutingTableIpv6StaticRouteBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.PathMonitor != nil {
		var obj vrfRoutingTableIpv6StaticRoutePathMonitorXml
		obj.MarshalFromObject(*s.PathMonitor)
		o.PathMonitor = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoute, error) {
	var nexthopVal *VrfRoutingTableIpv6StaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}
	var bfdVal *VrfRoutingTableIpv6StaticRouteBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var pathMonitorVal *VrfRoutingTableIpv6StaticRoutePathMonitor
	if o.PathMonitor != nil {
		obj, err := o.PathMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pathMonitorVal = obj
	}

	result := &VrfRoutingTableIpv6StaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		AdminDist:   o.AdminDist,
		Metric:      o.Metric,
		Nexthop:     nexthopVal,
		Bfd:         bfdVal,
		PathMonitor: pathMonitorVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteNexthopXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteNexthop) {
	if s.Discard != nil {
		var obj vrfRoutingTableIpv6StaticRouteNexthopDiscardXml
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	o.Ipv6Address = s.Ipv6Address
	o.Fqdn = s.Fqdn
	o.NextLr = s.NextLr
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteNexthopXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteNexthop, error) {
	var discardVal *VrfRoutingTableIpv6StaticRouteNexthopDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}

	result := &VrfRoutingTableIpv6StaticRouteNexthop{
		Discard:     discardVal,
		Ipv6Address: o.Ipv6Address,
		Fqdn:        o.Fqdn,
		NextLr:      o.NextLr,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteNexthopDiscardXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteNexthopDiscard) {
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteNexthopDiscardXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteNexthopDiscard, error) {

	result := &VrfRoutingTableIpv6StaticRouteNexthopDiscard{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteBfdXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteBfdXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteBfd, error) {

	result := &VrfRoutingTableIpv6StaticRouteBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRoutePathMonitorXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRoutePathMonitor) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FailureCondition = s.FailureCondition
	o.HoldTime = s.HoldTime
	if s.MonitorDestinations != nil {
		var objs []vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml
		for _, elt := range s.MonitorDestinations {
			var obj vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MonitorDestinations = &vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRoutePathMonitorXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoutePathMonitor, error) {
	var monitorDestinationsVal []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations
	if o.MonitorDestinations != nil {
		for _, elt := range o.MonitorDestinations.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			monitorDestinationsVal = append(monitorDestinationsVal, *obj)
		}
	}

	result := &VrfRoutingTableIpv6StaticRoutePathMonitor{
		Enable:              util.AsBool(o.Enable, nil),
		FailureCondition:    o.FailureCondition,
		HoldTime:            o.HoldTime,
		MonitorDestinations: monitorDestinationsVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml) MarshalFromObject(s VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Source = s.Source
	o.Destination = s.Destination
	o.Interval = s.Interval
	o.Count = s.Count
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations, error) {

	result := &VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{
		Name:        o.Name,
		Enable:      util.AsBool(o.Enable, nil),
		Source:      o.Source,
		Destination: o.Destination,
		Interval:    o.Interval,
		Count:       o.Count,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfXml) MarshalFromObject(s VrfOspf) {
	o.RouterId = s.RouterId
	o.Enable = util.YesNo(s.Enable, nil)
	o.Rfc1583 = util.YesNo(s.Rfc1583, nil)
	o.SpfTimer = s.SpfTimer
	o.GlobalIfTimer = s.GlobalIfTimer
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfOspfGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfOspfGracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.Area != nil {
		var objs []vrfOspfAreaXml
		for _, elt := range s.Area {
			var obj vrfOspfAreaXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &vrfOspfAreaContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfXml) UnmarshalToObject() (*VrfOspf, error) {
	var globalBfdVal *VrfOspfGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *VrfOspfGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var areaVal []VrfOspfArea
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}

	result := &VrfOspf{
		RouterId:              o.RouterId,
		Enable:                util.AsBool(o.Enable, nil),
		Rfc1583:               util.AsBool(o.Rfc1583, nil),
		SpfTimer:              o.SpfTimer,
		GlobalIfTimer:         o.GlobalIfTimer,
		RedistributionProfile: o.RedistributionProfile,
		GlobalBfd:             globalBfdVal,
		GracefulRestart:       gracefulRestartVal,
		Area:                  areaVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfOspfGlobalBfdXml) MarshalFromObject(s VrfOspfGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfGlobalBfdXml) UnmarshalToObject() (*VrfOspfGlobalBfd, error) {

	result := &VrfOspfGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfGracefulRestartXml) MarshalFromObject(s VrfOspfGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.Misc = s.Misc
}

func (o vrfOspfGracefulRestartXml) UnmarshalToObject() (*VrfOspfGracefulRestart, error) {

	result := &VrfOspfGracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaXml) MarshalFromObject(s VrfOspfArea) {
	o.Name = s.Name
	o.Authentication = s.Authentication
	if s.Type != nil {
		var obj vrfOspfAreaTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []vrfOspfAreaRangeXml
		for _, elt := range s.Range {
			var obj vrfOspfAreaRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &vrfOspfAreaRangeContainerXml{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfOspfAreaInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfOspfAreaInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfOspfAreaInterfaceContainerXml{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []vrfOspfAreaVirtualLinkXml
		for _, elt := range s.VirtualLink {
			var obj vrfOspfAreaVirtualLinkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &vrfOspfAreaVirtualLinkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaXml) UnmarshalToObject() (*VrfOspfArea, error) {
	var typeVal *VrfOspfAreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []VrfOspfAreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []VrfOspfAreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []VrfOspfAreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &VrfOspfArea{
		Name:           o.Name,
		Authentication: o.Authentication,
		Type:           typeVal,
		Range:          rangeVal,
		Interface:      interfaceVal,
		VirtualLink:    virtualLinkVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeXml) MarshalFromObject(s VrfOspfAreaType) {
	if s.Normal != nil {
		var obj vrfOspfAreaTypeNormalXml
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj vrfOspfAreaTypeStubXml
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj vrfOspfAreaTypeNssaXml
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeXml) UnmarshalToObject() (*VrfOspfAreaType, error) {
	var normalVal *VrfOspfAreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *VrfOspfAreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *VrfOspfAreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &VrfOspfAreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNormalXml) MarshalFromObject(s VrfOspfAreaTypeNormal) {
	if s.Abr != nil {
		var obj vrfOspfAreaTypeNormalAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNormalXml) UnmarshalToObject() (*VrfOspfAreaTypeNormal, error) {
	var abrVal *VrfOspfAreaTypeNormalAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeNormal{
		Abr:  abrVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNormalAbrXml) MarshalFromObject(s VrfOspfAreaTypeNormalAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNormalAbrXml) UnmarshalToObject() (*VrfOspfAreaTypeNormalAbr, error) {

	result := &VrfOspfAreaTypeNormalAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeStubXml) MarshalFromObject(s VrfOspfAreaTypeStub) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.Abr != nil {
		var obj vrfOspfAreaTypeStubAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeStubXml) UnmarshalToObject() (*VrfOspfAreaTypeStub, error) {
	var abrVal *VrfOspfAreaTypeStubAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeStub{
		NoSummary:          util.AsBool(o.NoSummary, nil),
		Abr:                abrVal,
		DefaultRouteMetric: o.DefaultRouteMetric,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeStubAbrXml) MarshalFromObject(s VrfOspfAreaTypeStubAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeStubAbrXml) UnmarshalToObject() (*VrfOspfAreaTypeStubAbr, error) {

	result := &VrfOspfAreaTypeStubAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaXml) MarshalFromObject(s VrfOspfAreaTypeNssa) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.DefaultInformationOriginate != nil {
		var obj vrfOspfAreaTypeNssaDefaultInformationOriginateXml
		obj.MarshalFromObject(*s.DefaultInformationOriginate)
		o.DefaultInformationOriginate = &obj
	}
	if s.Abr != nil {
		var obj vrfOspfAreaTypeNssaAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaXml) UnmarshalToObject() (*VrfOspfAreaTypeNssa, error) {
	var defaultInformationOriginateVal *VrfOspfAreaTypeNssaDefaultInformationOriginate
	if o.DefaultInformationOriginate != nil {
		obj, err := o.DefaultInformationOriginate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultInformationOriginateVal = obj
	}
	var abrVal *VrfOspfAreaTypeNssaAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeNssa{
		NoSummary:                   util.AsBool(o.NoSummary, nil),
		DefaultInformationOriginate: defaultInformationOriginateVal,
		Abr:                         abrVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaDefaultInformationOriginateXml) MarshalFromObject(s VrfOspfAreaTypeNssaDefaultInformationOriginate) {
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaDefaultInformationOriginateXml) UnmarshalToObject() (*VrfOspfAreaTypeNssaDefaultInformationOriginate, error) {

	result := &VrfOspfAreaTypeNssaDefaultInformationOriginate{
		Metric:     o.Metric,
		MetricType: o.MetricType,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaAbrXml) MarshalFromObject(s VrfOspfAreaTypeNssaAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	if s.NssaExtRange != nil {
		var objs []vrfOspfAreaTypeNssaAbrNssaExtRangeXml
		for _, elt := range s.NssaExtRange {
			var obj vrfOspfAreaTypeNssaAbrNssaExtRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaAbrXml) UnmarshalToObject() (*VrfOspfAreaTypeNssaAbr, error) {
	var nssaExtRangeVal []VrfOspfAreaTypeNssaAbrNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &VrfOspfAreaTypeNssaAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		NssaExtRange:       nssaExtRangeVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaAbrNssaExtRangeXml) MarshalFromObject(s VrfOspfAreaTypeNssaAbrNssaExtRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaAbrNssaExtRangeXml) UnmarshalToObject() (*VrfOspfAreaTypeNssaAbrNssaExtRange, error) {

	result := &VrfOspfAreaTypeNssaAbrNssaExtRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaRangeXml) MarshalFromObject(s VrfOspfAreaRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfAreaRangeXml) UnmarshalToObject() (*VrfOspfAreaRange, error) {

	result := &VrfOspfAreaRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceXml) MarshalFromObject(s VrfOspfAreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.MtuIgnore = util.YesNo(s.MtuIgnore, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.Priority = s.Priority
	o.Metric = s.Metric
	o.Authentication = s.Authentication
	o.Timing = s.Timing
	if s.LinkType != nil {
		var obj vrfOspfAreaInterfaceLinkTypeXml
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Bfd != nil {
		var obj vrfOspfAreaInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceXml) UnmarshalToObject() (*VrfOspfAreaInterface, error) {
	var linkTypeVal *VrfOspfAreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var bfdVal *VrfOspfAreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfAreaInterface{
		Name:           o.Name,
		Enable:         util.AsBool(o.Enable, nil),
		MtuIgnore:      util.AsBool(o.MtuIgnore, nil),
		Passive:        util.AsBool(o.Passive, nil),
		Priority:       o.Priority,
		Metric:         o.Metric,
		Authentication: o.Authentication,
		Timing:         o.Timing,
		LinkType:       linkTypeVal,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeXml) MarshalFromObject(s VrfOspfAreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj vrfOspfAreaInterfaceLinkTypeBroadcastXml
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj vrfOspfAreaInterfaceLinkTypeP2pXml
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj vrfOspfAreaInterfaceLinkTypeP2mpXml
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeXml) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkType, error) {
	var broadcastVal *VrfOspfAreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *VrfOspfAreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *VrfOspfAreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &VrfOspfAreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeBroadcastXml) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeBroadcastXml) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeBroadcast, error) {

	result := &VrfOspfAreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2pXml) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2pXml) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2p, error) {

	result := &VrfOspfAreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2mpXml) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2mp) {
	if s.Neighbor != nil {
		var objs []vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml
		for _, elt := range s.Neighbor {
			var obj vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2mpXml) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2mp, error) {
	var neighborVal []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &VrfOspfAreaInterfaceLinkTypeP2mp{
		Neighbor: neighborVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2mpNeighbor) {
	o.Name = s.Name
	o.Priority = s.Priority
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2mpNeighbor, error) {

	result := &VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{
		Name:     o.Name,
		Priority: o.Priority,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceBfdXml) MarshalFromObject(s VrfOspfAreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceBfdXml) UnmarshalToObject() (*VrfOspfAreaInterfaceBfd, error) {

	result := &VrfOspfAreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaVirtualLinkXml) MarshalFromObject(s VrfOspfAreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.Timing = s.Timing
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj vrfOspfAreaVirtualLinkBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaVirtualLinkXml) UnmarshalToObject() (*VrfOspfAreaVirtualLink, error) {
	var bfdVal *VrfOspfAreaVirtualLinkBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfAreaVirtualLink{
		Name:           o.Name,
		NeighborId:     o.NeighborId,
		TransitAreaId:  o.TransitAreaId,
		Enable:         util.AsBool(o.Enable, nil),
		InstanceId:     o.InstanceId,
		Timing:         o.Timing,
		Authentication: o.Authentication,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaVirtualLinkBfdXml) MarshalFromObject(s VrfOspfAreaVirtualLinkBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfAreaVirtualLinkBfdXml) UnmarshalToObject() (*VrfOspfAreaVirtualLinkBfd, error) {

	result := &VrfOspfAreaVirtualLinkBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3Xml) MarshalFromObject(s VrfOspfv3) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RouterId = s.RouterId
	o.DisableTransitTraffic = util.YesNo(s.DisableTransitTraffic, nil)
	o.SpfTimer = s.SpfTimer
	o.GlobalIfTimer = s.GlobalIfTimer
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfOspfv3GlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfOspfv3GracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.Area != nil {
		var objs []vrfOspfv3AreaXml
		for _, elt := range s.Area {
			var obj vrfOspfv3AreaXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &vrfOspfv3AreaContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3Xml) UnmarshalToObject() (*VrfOspfv3, error) {
	var globalBfdVal *VrfOspfv3GlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *VrfOspfv3GracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var areaVal []VrfOspfv3Area
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}

	result := &VrfOspfv3{
		Enable:                util.AsBool(o.Enable, nil),
		RouterId:              o.RouterId,
		DisableTransitTraffic: util.AsBool(o.DisableTransitTraffic, nil),
		SpfTimer:              o.SpfTimer,
		GlobalIfTimer:         o.GlobalIfTimer,
		RedistributionProfile: o.RedistributionProfile,
		GlobalBfd:             globalBfdVal,
		GracefulRestart:       gracefulRestartVal,
		Area:                  areaVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3GlobalBfdXml) MarshalFromObject(s VrfOspfv3GlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfv3GlobalBfdXml) UnmarshalToObject() (*VrfOspfv3GlobalBfd, error) {

	result := &VrfOspfv3GlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3GracefulRestartXml) MarshalFromObject(s VrfOspfv3GracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.Misc = s.Misc
}

func (o vrfOspfv3GracefulRestartXml) UnmarshalToObject() (*VrfOspfv3GracefulRestart, error) {

	result := &VrfOspfv3GracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaXml) MarshalFromObject(s VrfOspfv3Area) {
	o.Name = s.Name
	o.Authentication = s.Authentication
	if s.Type != nil {
		var obj vrfOspfv3AreaTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []vrfOspfv3AreaRangeXml
		for _, elt := range s.Range {
			var obj vrfOspfv3AreaRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &vrfOspfv3AreaRangeContainerXml{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfOspfv3AreaInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfOspfv3AreaInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfOspfv3AreaInterfaceContainerXml{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []vrfOspfv3AreaVirtualLinkXml
		for _, elt := range s.VirtualLink {
			var obj vrfOspfv3AreaVirtualLinkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &vrfOspfv3AreaVirtualLinkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaXml) UnmarshalToObject() (*VrfOspfv3Area, error) {
	var typeVal *VrfOspfv3AreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []VrfOspfv3AreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []VrfOspfv3AreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []VrfOspfv3AreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &VrfOspfv3Area{
		Name:           o.Name,
		Authentication: o.Authentication,
		Type:           typeVal,
		Range:          rangeVal,
		Interface:      interfaceVal,
		VirtualLink:    virtualLinkVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeXml) MarshalFromObject(s VrfOspfv3AreaType) {
	if s.Normal != nil {
		var obj vrfOspfv3AreaTypeNormalXml
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj vrfOspfv3AreaTypeStubXml
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj vrfOspfv3AreaTypeNssaXml
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeXml) UnmarshalToObject() (*VrfOspfv3AreaType, error) {
	var normalVal *VrfOspfv3AreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *VrfOspfv3AreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *VrfOspfv3AreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &VrfOspfv3AreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNormalXml) MarshalFromObject(s VrfOspfv3AreaTypeNormal) {
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeNormalAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNormalXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNormal, error) {
	var abrVal *VrfOspfv3AreaTypeNormalAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeNormal{
		Abr:  abrVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNormalAbrXml) MarshalFromObject(s VrfOspfv3AreaTypeNormalAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNormalAbrXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNormalAbr, error) {

	result := &VrfOspfv3AreaTypeNormalAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeStubXml) MarshalFromObject(s VrfOspfv3AreaTypeStub) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeStubAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeStubXml) UnmarshalToObject() (*VrfOspfv3AreaTypeStub, error) {
	var abrVal *VrfOspfv3AreaTypeStubAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeStub{
		NoSummary:          util.AsBool(o.NoSummary, nil),
		Abr:                abrVal,
		DefaultRouteMetric: o.DefaultRouteMetric,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeStubAbrXml) MarshalFromObject(s VrfOspfv3AreaTypeStubAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeStubAbrXml) UnmarshalToObject() (*VrfOspfv3AreaTypeStubAbr, error) {

	result := &VrfOspfv3AreaTypeStubAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaXml) MarshalFromObject(s VrfOspfv3AreaTypeNssa) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.DefaultInformationOriginate != nil {
		var obj vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml
		obj.MarshalFromObject(*s.DefaultInformationOriginate)
		o.DefaultInformationOriginate = &obj
	}
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeNssaAbrXml
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNssa, error) {
	var defaultInformationOriginateVal *VrfOspfv3AreaTypeNssaDefaultInformationOriginate
	if o.DefaultInformationOriginate != nil {
		obj, err := o.DefaultInformationOriginate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultInformationOriginateVal = obj
	}
	var abrVal *VrfOspfv3AreaTypeNssaAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeNssa{
		NoSummary:                   util.AsBool(o.NoSummary, nil),
		DefaultInformationOriginate: defaultInformationOriginateVal,
		Abr:                         abrVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml) MarshalFromObject(s VrfOspfv3AreaTypeNssaDefaultInformationOriginate) {
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaDefaultInformationOriginate, error) {

	result := &VrfOspfv3AreaTypeNssaDefaultInformationOriginate{
		Metric:     o.Metric,
		MetricType: o.MetricType,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaAbrXml) MarshalFromObject(s VrfOspfv3AreaTypeNssaAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	if s.NssaExtRange != nil {
		var objs []vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml
		for _, elt := range s.NssaExtRange {
			var obj vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaAbrXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaAbr, error) {
	var nssaExtRangeVal []VrfOspfv3AreaTypeNssaAbrNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &VrfOspfv3AreaTypeNssaAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		NssaExtRange:       nssaExtRangeVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml) MarshalFromObject(s VrfOspfv3AreaTypeNssaAbrNssaExtRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaAbrNssaExtRange, error) {

	result := &VrfOspfv3AreaTypeNssaAbrNssaExtRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaRangeXml) MarshalFromObject(s VrfOspfv3AreaRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaRangeXml) UnmarshalToObject() (*VrfOspfv3AreaRange, error) {

	result := &VrfOspfv3AreaRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceXml) MarshalFromObject(s VrfOspfv3AreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.MtuIgnore = util.YesNo(s.MtuIgnore, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.Priority = s.Priority
	o.Metric = s.Metric
	o.InstanceId = s.InstanceId
	o.Authentication = s.Authentication
	o.Timing = s.Timing
	if s.LinkType != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeXml
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Bfd != nil {
		var obj vrfOspfv3AreaInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceXml) UnmarshalToObject() (*VrfOspfv3AreaInterface, error) {
	var linkTypeVal *VrfOspfv3AreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var bfdVal *VrfOspfv3AreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfv3AreaInterface{
		Name:           o.Name,
		Enable:         util.AsBool(o.Enable, nil),
		MtuIgnore:      util.AsBool(o.MtuIgnore, nil),
		Passive:        util.AsBool(o.Passive, nil),
		Priority:       o.Priority,
		Metric:         o.Metric,
		InstanceId:     o.InstanceId,
		Authentication: o.Authentication,
		Timing:         o.Timing,
		LinkType:       linkTypeVal,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeXml) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeBroadcastXml
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeP2pXml
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeP2mpXml
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkType, error) {
	var broadcastVal *VrfOspfv3AreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *VrfOspfv3AreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *VrfOspfv3AreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &VrfOspfv3AreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeBroadcastXml) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeBroadcastXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeBroadcast, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2pXml) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2pXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2p, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2mpXml) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2mp) {
	if s.Neighbor != nil {
		var objs []vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml
		for _, elt := range s.Neighbor {
			var obj vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2mpXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2mp, error) {
	var neighborVal []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &VrfOspfv3AreaInterfaceLinkTypeP2mp{
		Neighbor: neighborVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor) {
	o.Name = s.Name
	o.Priority = s.Priority
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{
		Name:     o.Name,
		Priority: o.Priority,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceBfdXml) MarshalFromObject(s VrfOspfv3AreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceBfdXml) UnmarshalToObject() (*VrfOspfv3AreaInterfaceBfd, error) {

	result := &VrfOspfv3AreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaVirtualLinkXml) MarshalFromObject(s VrfOspfv3AreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.Timing = s.Timing
	o.Authentication = s.Authentication
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaVirtualLinkXml) UnmarshalToObject() (*VrfOspfv3AreaVirtualLink, error) {

	result := &VrfOspfv3AreaVirtualLink{
		Name:           o.Name,
		NeighborId:     o.NeighborId,
		TransitAreaId:  o.TransitAreaId,
		Enable:         util.AsBool(o.Enable, nil),
		InstanceId:     o.InstanceId,
		Timing:         o.Timing,
		Authentication: o.Authentication,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpXml) MarshalFromObject(s VrfEcmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.MaxPath = s.MaxPath
	o.SymmetricReturn = util.YesNo(s.SymmetricReturn, nil)
	o.StrictSourcePath = util.YesNo(s.StrictSourcePath, nil)
	if s.Algorithm != nil {
		var obj vrfEcmpAlgorithmXml
		obj.MarshalFromObject(*s.Algorithm)
		o.Algorithm = &obj
	}
	o.Misc = s.Misc
}

func (o vrfEcmpXml) UnmarshalToObject() (*VrfEcmp, error) {
	var algorithmVal *VrfEcmpAlgorithm
	if o.Algorithm != nil {
		obj, err := o.Algorithm.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		algorithmVal = obj
	}

	result := &VrfEcmp{
		Enable:           util.AsBool(o.Enable, nil),
		MaxPath:          o.MaxPath,
		SymmetricReturn:  util.AsBool(o.SymmetricReturn, nil),
		StrictSourcePath: util.AsBool(o.StrictSourcePath, nil),
		Algorithm:        algorithmVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmXml) MarshalFromObject(s VrfEcmpAlgorithm) {
	if s.IpModulo != nil {
		var obj vrfEcmpAlgorithmIpModuloXml
		obj.MarshalFromObject(*s.IpModulo)
		o.IpModulo = &obj
	}
	if s.IpHash != nil {
		var obj vrfEcmpAlgorithmIpHashXml
		obj.MarshalFromObject(*s.IpHash)
		o.IpHash = &obj
	}
	if s.WeightedRoundRobin != nil {
		var obj vrfEcmpAlgorithmWeightedRoundRobinXml
		obj.MarshalFromObject(*s.WeightedRoundRobin)
		o.WeightedRoundRobin = &obj
	}
	if s.BalancedRoundRobin != nil {
		var obj vrfEcmpAlgorithmBalancedRoundRobinXml
		obj.MarshalFromObject(*s.BalancedRoundRobin)
		o.BalancedRoundRobin = &obj
	}
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmXml) UnmarshalToObject() (*VrfEcmpAlgorithm, error) {
	var ipModuloVal *VrfEcmpAlgorithmIpModulo
	if o.IpModulo != nil {
		obj, err := o.IpModulo.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipModuloVal = obj
	}
	var ipHashVal *VrfEcmpAlgorithmIpHash
	if o.IpHash != nil {
		obj, err := o.IpHash.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipHashVal = obj
	}
	var weightedRoundRobinVal *VrfEcmpAlgorithmWeightedRoundRobin
	if o.WeightedRoundRobin != nil {
		obj, err := o.WeightedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weightedRoundRobinVal = obj
	}
	var balancedRoundRobinVal *VrfEcmpAlgorithmBalancedRoundRobin
	if o.BalancedRoundRobin != nil {
		obj, err := o.BalancedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		balancedRoundRobinVal = obj
	}

	result := &VrfEcmpAlgorithm{
		IpModulo:           ipModuloVal,
		IpHash:             ipHashVal,
		WeightedRoundRobin: weightedRoundRobinVal,
		BalancedRoundRobin: balancedRoundRobinVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmIpModuloXml) MarshalFromObject(s VrfEcmpAlgorithmIpModulo) {
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmIpModuloXml) UnmarshalToObject() (*VrfEcmpAlgorithmIpModulo, error) {

	result := &VrfEcmpAlgorithmIpModulo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmIpHashXml) MarshalFromObject(s VrfEcmpAlgorithmIpHash) {
	o.SrcOnly = util.YesNo(s.SrcOnly, nil)
	o.UsePort = util.YesNo(s.UsePort, nil)
	o.HashSeed = s.HashSeed
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmIpHashXml) UnmarshalToObject() (*VrfEcmpAlgorithmIpHash, error) {

	result := &VrfEcmpAlgorithmIpHash{
		SrcOnly:  util.AsBool(o.SrcOnly, nil),
		UsePort:  util.AsBool(o.UsePort, nil),
		HashSeed: o.HashSeed,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmWeightedRoundRobinXml) MarshalFromObject(s VrfEcmpAlgorithmWeightedRoundRobin) {
	if s.Interface != nil {
		var objs []vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmWeightedRoundRobinXml) UnmarshalToObject() (*VrfEcmpAlgorithmWeightedRoundRobin, error) {
	var interfaceVal []VrfEcmpAlgorithmWeightedRoundRobinInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfEcmpAlgorithmWeightedRoundRobin{
		Interface: interfaceVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml) MarshalFromObject(s VrfEcmpAlgorithmWeightedRoundRobinInterface) {
	o.Name = s.Name
	o.Weight = s.Weight
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml) UnmarshalToObject() (*VrfEcmpAlgorithmWeightedRoundRobinInterface, error) {

	result := &VrfEcmpAlgorithmWeightedRoundRobinInterface{
		Name:   o.Name,
		Weight: o.Weight,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmBalancedRoundRobinXml) MarshalFromObject(s VrfEcmpAlgorithmBalancedRoundRobin) {
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmBalancedRoundRobinXml) UnmarshalToObject() (*VrfEcmpAlgorithmBalancedRoundRobin, error) {

	result := &VrfEcmpAlgorithmBalancedRoundRobin{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastXml) MarshalFromObject(s VrfMulticast) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.StaticRoute != nil {
		var objs []vrfMulticastStaticRouteXml
		for _, elt := range s.StaticRoute {
			var obj vrfMulticastStaticRouteXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfMulticastStaticRouteContainerXml{Entries: objs}
	}
	if s.Pim != nil {
		var obj vrfMulticastPimXml
		obj.MarshalFromObject(*s.Pim)
		o.Pim = &obj
	}
	if s.Igmp != nil {
		var obj vrfMulticastIgmpXml
		obj.MarshalFromObject(*s.Igmp)
		o.Igmp = &obj
	}
	if s.Msdp != nil {
		var obj vrfMulticastMsdpXml
		obj.MarshalFromObject(*s.Msdp)
		o.Msdp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastXml) UnmarshalToObject() (*VrfMulticast, error) {
	var staticRouteVal []VrfMulticastStaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}
	var pimVal *VrfMulticastPim
	if o.Pim != nil {
		obj, err := o.Pim.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pimVal = obj
	}
	var igmpVal *VrfMulticastIgmp
	if o.Igmp != nil {
		obj, err := o.Igmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		igmpVal = obj
	}
	var msdpVal *VrfMulticastMsdp
	if o.Msdp != nil {
		obj, err := o.Msdp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		msdpVal = obj
	}

	result := &VrfMulticast{
		Enable:      util.AsBool(o.Enable, nil),
		StaticRoute: staticRouteVal,
		Pim:         pimVal,
		Igmp:        igmpVal,
		Msdp:        msdpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastStaticRouteXml) MarshalFromObject(s VrfMulticastStaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.Preference = s.Preference
	if s.Nexthop != nil {
		var obj vrfMulticastStaticRouteNexthopXml
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastStaticRouteXml) UnmarshalToObject() (*VrfMulticastStaticRoute, error) {
	var nexthopVal *VrfMulticastStaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}

	result := &VrfMulticastStaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		Preference:  o.Preference,
		Nexthop:     nexthopVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastStaticRouteNexthopXml) MarshalFromObject(s VrfMulticastStaticRouteNexthop) {
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
}

func (o vrfMulticastStaticRouteNexthopXml) UnmarshalToObject() (*VrfMulticastStaticRouteNexthop, error) {

	result := &VrfMulticastStaticRouteNexthop{
		IpAddress: o.IpAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimXml) MarshalFromObject(s VrfMulticastPim) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RpfLookupMode = s.RpfLookupMode
	o.RouteAgeoutTime = s.RouteAgeoutTime
	o.IfTimerGlobal = s.IfTimerGlobal
	o.GroupPermission = s.GroupPermission
	if s.SsmAddressSpace != nil {
		var obj vrfMulticastPimSsmAddressSpaceXml
		obj.MarshalFromObject(*s.SsmAddressSpace)
		o.SsmAddressSpace = &obj
	}
	if s.Rp != nil {
		var obj vrfMulticastPimRpXml
		obj.MarshalFromObject(*s.Rp)
		o.Rp = &obj
	}
	if s.SptThreshold != nil {
		var objs []vrfMulticastPimSptThresholdXml
		for _, elt := range s.SptThreshold {
			var obj vrfMulticastPimSptThresholdXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SptThreshold = &vrfMulticastPimSptThresholdContainerXml{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfMulticastPimInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfMulticastPimInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfMulticastPimInterfaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimXml) UnmarshalToObject() (*VrfMulticastPim, error) {
	var ssmAddressSpaceVal *VrfMulticastPimSsmAddressSpace
	if o.SsmAddressSpace != nil {
		obj, err := o.SsmAddressSpace.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ssmAddressSpaceVal = obj
	}
	var rpVal *VrfMulticastPimRp
	if o.Rp != nil {
		obj, err := o.Rp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		rpVal = obj
	}
	var sptThresholdVal []VrfMulticastPimSptThreshold
	if o.SptThreshold != nil {
		for _, elt := range o.SptThreshold.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			sptThresholdVal = append(sptThresholdVal, *obj)
		}
	}
	var interfaceVal []VrfMulticastPimInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfMulticastPim{
		Enable:          util.AsBool(o.Enable, nil),
		RpfLookupMode:   o.RpfLookupMode,
		RouteAgeoutTime: o.RouteAgeoutTime,
		IfTimerGlobal:   o.IfTimerGlobal,
		GroupPermission: o.GroupPermission,
		SsmAddressSpace: ssmAddressSpaceVal,
		Rp:              rpVal,
		SptThreshold:    sptThresholdVal,
		Interface:       interfaceVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimSsmAddressSpaceXml) MarshalFromObject(s VrfMulticastPimSsmAddressSpace) {
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimSsmAddressSpaceXml) UnmarshalToObject() (*VrfMulticastPimSsmAddressSpace, error) {

	result := &VrfMulticastPimSsmAddressSpace{
		GroupList: o.GroupList,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpXml) MarshalFromObject(s VrfMulticastPimRp) {
	if s.LocalRp != nil {
		var obj vrfMulticastPimRpLocalRpXml
		obj.MarshalFromObject(*s.LocalRp)
		o.LocalRp = &obj
	}
	if s.ExternalRp != nil {
		var objs []vrfMulticastPimRpExternalRpXml
		for _, elt := range s.ExternalRp {
			var obj vrfMulticastPimRpExternalRpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExternalRp = &vrfMulticastPimRpExternalRpContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpXml) UnmarshalToObject() (*VrfMulticastPimRp, error) {
	var localRpVal *VrfMulticastPimRpLocalRp
	if o.LocalRp != nil {
		obj, err := o.LocalRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localRpVal = obj
	}
	var externalRpVal []VrfMulticastPimRpExternalRp
	if o.ExternalRp != nil {
		for _, elt := range o.ExternalRp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			externalRpVal = append(externalRpVal, *obj)
		}
	}

	result := &VrfMulticastPimRp{
		LocalRp:    localRpVal,
		ExternalRp: externalRpVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpXml) MarshalFromObject(s VrfMulticastPimRpLocalRp) {
	if s.StaticRp != nil {
		var obj vrfMulticastPimRpLocalRpStaticRpXml
		obj.MarshalFromObject(*s.StaticRp)
		o.StaticRp = &obj
	}
	if s.CandidateRp != nil {
		var obj vrfMulticastPimRpLocalRpCandidateRpXml
		obj.MarshalFromObject(*s.CandidateRp)
		o.CandidateRp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpXml) UnmarshalToObject() (*VrfMulticastPimRpLocalRp, error) {
	var staticRpVal *VrfMulticastPimRpLocalRpStaticRp
	if o.StaticRp != nil {
		obj, err := o.StaticRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticRpVal = obj
	}
	var candidateRpVal *VrfMulticastPimRpLocalRpCandidateRp
	if o.CandidateRp != nil {
		obj, err := o.CandidateRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		candidateRpVal = obj
	}

	result := &VrfMulticastPimRpLocalRp{
		StaticRp:    staticRpVal,
		CandidateRp: candidateRpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpStaticRpXml) MarshalFromObject(s VrfMulticastPimRpLocalRpStaticRp) {
	o.Interface = s.Interface
	o.Address = s.Address
	o.Override = util.YesNo(s.Override, nil)
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpStaticRpXml) UnmarshalToObject() (*VrfMulticastPimRpLocalRpStaticRp, error) {

	result := &VrfMulticastPimRpLocalRpStaticRp{
		Interface: o.Interface,
		Address:   o.Address,
		Override:  util.AsBool(o.Override, nil),
		GroupList: o.GroupList,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpCandidateRpXml) MarshalFromObject(s VrfMulticastPimRpLocalRpCandidateRp) {
	o.Interface = s.Interface
	o.Address = s.Address
	o.Priority = s.Priority
	o.AdvertisementInterval = s.AdvertisementInterval
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpCandidateRpXml) UnmarshalToObject() (*VrfMulticastPimRpLocalRpCandidateRp, error) {

	result := &VrfMulticastPimRpLocalRpCandidateRp{
		Interface:             o.Interface,
		Address:               o.Address,
		Priority:              o.Priority,
		AdvertisementInterval: o.AdvertisementInterval,
		GroupList:             o.GroupList,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpExternalRpXml) MarshalFromObject(s VrfMulticastPimRpExternalRp) {
	o.Name = s.Name
	o.GroupList = s.GroupList
	o.Override = util.YesNo(s.Override, nil)
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpExternalRpXml) UnmarshalToObject() (*VrfMulticastPimRpExternalRp, error) {

	result := &VrfMulticastPimRpExternalRp{
		Name:      o.Name,
		GroupList: o.GroupList,
		Override:  util.AsBool(o.Override, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimSptThresholdXml) MarshalFromObject(s VrfMulticastPimSptThreshold) {
	o.Name = s.Name
	o.Threshold = s.Threshold
	o.Misc = s.Misc
}

func (o vrfMulticastPimSptThresholdXml) UnmarshalToObject() (*VrfMulticastPimSptThreshold, error) {

	result := &VrfMulticastPimSptThreshold{
		Name:      o.Name,
		Threshold: o.Threshold,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimInterfaceXml) MarshalFromObject(s VrfMulticastPimInterface) {
	o.Name = s.Name
	o.Description = s.Description
	o.DrPriority = s.DrPriority
	o.SendBsm = util.YesNo(s.SendBsm, nil)
	o.IfTimer = s.IfTimer
	o.NeighborFilter = s.NeighborFilter
	o.Misc = s.Misc
}

func (o vrfMulticastPimInterfaceXml) UnmarshalToObject() (*VrfMulticastPimInterface, error) {

	result := &VrfMulticastPimInterface{
		Name:           o.Name,
		Description:    o.Description,
		DrPriority:     o.DrPriority,
		SendBsm:        util.AsBool(o.SendBsm, nil),
		IfTimer:        o.IfTimer,
		NeighborFilter: o.NeighborFilter,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpXml) MarshalFromObject(s VrfMulticastIgmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Dynamic != nil {
		var obj vrfMulticastIgmpDynamicXml
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.Static != nil {
		var objs []vrfMulticastIgmpStaticXml
		for _, elt := range s.Static {
			var obj vrfMulticastIgmpStaticXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Static = &vrfMulticastIgmpStaticContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpXml) UnmarshalToObject() (*VrfMulticastIgmp, error) {
	var dynamicVal *VrfMulticastIgmpDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var staticVal []VrfMulticastIgmpStatic
	if o.Static != nil {
		for _, elt := range o.Static.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticVal = append(staticVal, *obj)
		}
	}

	result := &VrfMulticastIgmp{
		Enable:  util.AsBool(o.Enable, nil),
		Dynamic: dynamicVal,
		Static:  staticVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpDynamicXml) MarshalFromObject(s VrfMulticastIgmpDynamic) {
	if s.Interface != nil {
		var objs []vrfMulticastIgmpDynamicInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfMulticastIgmpDynamicInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfMulticastIgmpDynamicInterfaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpDynamicXml) UnmarshalToObject() (*VrfMulticastIgmpDynamic, error) {
	var interfaceVal []VrfMulticastIgmpDynamicInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfMulticastIgmpDynamic{
		Interface: interfaceVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpDynamicInterfaceXml) MarshalFromObject(s VrfMulticastIgmpDynamicInterface) {
	o.Name = s.Name
	o.Version = s.Version
	o.Robustness = s.Robustness
	o.GroupFilter = s.GroupFilter
	o.MaxGroups = s.MaxGroups
	o.MaxSources = s.MaxSources
	o.QueryProfile = s.QueryProfile
	o.RouterAlertPolicing = util.YesNo(s.RouterAlertPolicing, nil)
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpDynamicInterfaceXml) UnmarshalToObject() (*VrfMulticastIgmpDynamicInterface, error) {

	result := &VrfMulticastIgmpDynamicInterface{
		Name:                o.Name,
		Version:             o.Version,
		Robustness:          o.Robustness,
		GroupFilter:         o.GroupFilter,
		MaxGroups:           o.MaxGroups,
		MaxSources:          o.MaxSources,
		QueryProfile:        o.QueryProfile,
		RouterAlertPolicing: util.AsBool(o.RouterAlertPolicing, nil),
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpStaticXml) MarshalFromObject(s VrfMulticastIgmpStatic) {
	o.Name = s.Name
	o.Interface = s.Interface
	o.GroupAddress = s.GroupAddress
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpStaticXml) UnmarshalToObject() (*VrfMulticastIgmpStatic, error) {

	result := &VrfMulticastIgmpStatic{
		Name:          o.Name,
		Interface:     o.Interface,
		GroupAddress:  o.GroupAddress,
		SourceAddress: o.SourceAddress,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpXml) MarshalFromObject(s VrfMulticastMsdp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GlobalTimer = s.GlobalTimer
	o.GlobalAuthentication = s.GlobalAuthentication
	if s.OriginatorId != nil {
		var obj vrfMulticastMsdpOriginatorIdXml
		obj.MarshalFromObject(*s.OriginatorId)
		o.OriginatorId = &obj
	}
	if s.Peer != nil {
		var objs []vrfMulticastMsdpPeerXml
		for _, elt := range s.Peer {
			var obj vrfMulticastMsdpPeerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Peer = &vrfMulticastMsdpPeerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpXml) UnmarshalToObject() (*VrfMulticastMsdp, error) {
	var originatorIdVal *VrfMulticastMsdpOriginatorId
	if o.OriginatorId != nil {
		obj, err := o.OriginatorId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		originatorIdVal = obj
	}
	var peerVal []VrfMulticastMsdpPeer
	if o.Peer != nil {
		for _, elt := range o.Peer.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerVal = append(peerVal, *obj)
		}
	}

	result := &VrfMulticastMsdp{
		Enable:               util.AsBool(o.Enable, nil),
		GlobalTimer:          o.GlobalTimer,
		GlobalAuthentication: o.GlobalAuthentication,
		OriginatorId:         originatorIdVal,
		Peer:                 peerVal,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpOriginatorIdXml) MarshalFromObject(s VrfMulticastMsdpOriginatorId) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpOriginatorIdXml) UnmarshalToObject() (*VrfMulticastMsdpOriginatorId, error) {

	result := &VrfMulticastMsdpOriginatorId{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerXml) MarshalFromObject(s VrfMulticastMsdpPeer) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.PeerAs = s.PeerAs
	o.Authentication = s.Authentication
	o.MaxSa = s.MaxSa
	o.InboundSaFilter = s.InboundSaFilter
	o.OutboundSaFilter = s.OutboundSaFilter
	if s.LocalAddress != nil {
		var obj vrfMulticastMsdpPeerLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj vrfMulticastMsdpPeerPeerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerXml) UnmarshalToObject() (*VrfMulticastMsdpPeer, error) {
	var localAddressVal *VrfMulticastMsdpPeerLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *VrfMulticastMsdpPeerPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}

	result := &VrfMulticastMsdpPeer{
		Name:             o.Name,
		Enable:           util.AsBool(o.Enable, nil),
		PeerAs:           o.PeerAs,
		Authentication:   o.Authentication,
		MaxSa:            o.MaxSa,
		InboundSaFilter:  o.InboundSaFilter,
		OutboundSaFilter: o.OutboundSaFilter,
		LocalAddress:     localAddressVal,
		PeerAddress:      peerAddressVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerLocalAddressXml) MarshalFromObject(s VrfMulticastMsdpPeerLocalAddress) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerLocalAddressXml) UnmarshalToObject() (*VrfMulticastMsdpPeerLocalAddress, error) {

	result := &VrfMulticastMsdpPeerLocalAddress{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerPeerAddressXml) MarshalFromObject(s VrfMulticastMsdpPeerPeerAddress) {
	o.Ip = s.Ip
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerPeerAddressXml) UnmarshalToObject() (*VrfMulticastMsdpPeerPeerAddress, error) {

	result := &VrfMulticastMsdpPeerPeerAddress{
		Ip:   o.Ip,
		Fqdn: o.Fqdn,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRipXml) MarshalFromObject(s VrfRip) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.DefaultInformationOriginate = util.YesNo(s.DefaultInformationOriginate, nil)
	o.GlobalTimer = s.GlobalTimer
	o.AuthProfile = s.AuthProfile
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfRipGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GlobalInboundDistributeList != nil {
		var obj vrfRipGlobalInboundDistributeListXml
		obj.MarshalFromObject(*s.GlobalInboundDistributeList)
		o.GlobalInboundDistributeList = &obj
	}
	if s.GlobalOutboundDistributeList != nil {
		var obj vrfRipGlobalOutboundDistributeListXml
		obj.MarshalFromObject(*s.GlobalOutboundDistributeList)
		o.GlobalOutboundDistributeList = &obj
	}
	if s.Interface != nil {
		var objs []vrfRipInterfaceXml
		for _, elt := range s.Interface {
			var obj vrfRipInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfRipInterfaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRipXml) UnmarshalToObject() (*VrfRip, error) {
	var globalBfdVal *VrfRipGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var globalInboundDistributeListVal *VrfRipGlobalInboundDistributeList
	if o.GlobalInboundDistributeList != nil {
		obj, err := o.GlobalInboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalInboundDistributeListVal = obj
	}
	var globalOutboundDistributeListVal *VrfRipGlobalOutboundDistributeList
	if o.GlobalOutboundDistributeList != nil {
		obj, err := o.GlobalOutboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalOutboundDistributeListVal = obj
	}
	var interfaceVal []VrfRipInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfRip{
		Enable:                       util.AsBool(o.Enable, nil),
		DefaultInformationOriginate:  util.AsBool(o.DefaultInformationOriginate, nil),
		GlobalTimer:                  o.GlobalTimer,
		AuthProfile:                  o.AuthProfile,
		RedistributionProfile:        o.RedistributionProfile,
		GlobalBfd:                    globalBfdVal,
		GlobalInboundDistributeList:  globalInboundDistributeListVal,
		GlobalOutboundDistributeList: globalOutboundDistributeListVal,
		Interface:                    interfaceVal,
		Misc:                         o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalBfdXml) MarshalFromObject(s VrfRipGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRipGlobalBfdXml) UnmarshalToObject() (*VrfRipGlobalBfd, error) {

	result := &VrfRipGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalInboundDistributeListXml) MarshalFromObject(s VrfRipGlobalInboundDistributeList) {
	o.AccessList = s.AccessList
	o.Misc = s.Misc
}

func (o vrfRipGlobalInboundDistributeListXml) UnmarshalToObject() (*VrfRipGlobalInboundDistributeList, error) {

	result := &VrfRipGlobalInboundDistributeList{
		AccessList: o.AccessList,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalOutboundDistributeListXml) MarshalFromObject(s VrfRipGlobalOutboundDistributeList) {
	o.AccessList = s.AccessList
	o.Misc = s.Misc
}

func (o vrfRipGlobalOutboundDistributeListXml) UnmarshalToObject() (*VrfRipGlobalOutboundDistributeList, error) {

	result := &VrfRipGlobalOutboundDistributeList{
		AccessList: o.AccessList,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceXml) MarshalFromObject(s VrfRipInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Mode = s.Mode
	o.SplitHorizon = s.SplitHorizon
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj vrfRipInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.InterfaceInboundDistributeList != nil {
		var obj vrfRipInterfaceInterfaceInboundDistributeListXml
		obj.MarshalFromObject(*s.InterfaceInboundDistributeList)
		o.InterfaceInboundDistributeList = &obj
	}
	if s.InterfaceOutboundDistributeList != nil {
		var obj vrfRipInterfaceInterfaceOutboundDistributeListXml
		obj.MarshalFromObject(*s.InterfaceOutboundDistributeList)
		o.InterfaceOutboundDistributeList = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRipInterfaceXml) UnmarshalToObject() (*VrfRipInterface, error) {
	var bfdVal *VrfRipInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var interfaceInboundDistributeListVal *VrfRipInterfaceInterfaceInboundDistributeList
	if o.InterfaceInboundDistributeList != nil {
		obj, err := o.InterfaceInboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceInboundDistributeListVal = obj
	}
	var interfaceOutboundDistributeListVal *VrfRipInterfaceInterfaceOutboundDistributeList
	if o.InterfaceOutboundDistributeList != nil {
		obj, err := o.InterfaceOutboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceOutboundDistributeListVal = obj
	}

	result := &VrfRipInterface{
		Name:                            o.Name,
		Enable:                          util.AsBool(o.Enable, nil),
		Mode:                            o.Mode,
		SplitHorizon:                    o.SplitHorizon,
		Authentication:                  o.Authentication,
		Bfd:                             bfdVal,
		InterfaceInboundDistributeList:  interfaceInboundDistributeListVal,
		InterfaceOutboundDistributeList: interfaceOutboundDistributeListVal,
		Misc:                            o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceBfdXml) MarshalFromObject(s VrfRipInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRipInterfaceBfdXml) UnmarshalToObject() (*VrfRipInterfaceBfd, error) {

	result := &VrfRipInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceInterfaceInboundDistributeListXml) MarshalFromObject(s VrfRipInterfaceInterfaceInboundDistributeList) {
	o.AccessList = s.AccessList
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o vrfRipInterfaceInterfaceInboundDistributeListXml) UnmarshalToObject() (*VrfRipInterfaceInterfaceInboundDistributeList, error) {

	result := &VrfRipInterfaceInterfaceInboundDistributeList{
		AccessList: o.AccessList,
		Metric:     o.Metric,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceInterfaceOutboundDistributeListXml) MarshalFromObject(s VrfRipInterfaceInterfaceOutboundDistributeList) {
	o.AccessList = s.AccessList
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o vrfRipInterfaceInterfaceOutboundDistributeListXml) UnmarshalToObject() (*VrfRipInterfaceInterfaceOutboundDistributeList, error) {

	result := &VrfRipInterfaceInterfaceOutboundDistributeList{
		AccessList: o.AccessList,
		Metric:     o.Metric,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Vrf != nil {
		var objs []vrfXml_11_0_2
		for _, elt := range s.Vrf {
			var obj vrfXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vrf = &vrfContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var vrfVal []Vrf
	if o.Vrf != nil {
		for _, elt := range o.Vrf.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vrfVal = append(vrfVal, *obj)
		}
	}

	result := &Entry{
		Name: o.Name,
		Vrf:  vrfVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfXml_11_0_2) MarshalFromObject(s Vrf) {
	o.Name = s.Name
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.AdminDists != nil {
		var obj vrfAdminDistsXml_11_0_2
		obj.MarshalFromObject(*s.AdminDists)
		o.AdminDists = &obj
	}
	if s.RibFilter != nil {
		var obj vrfRibFilterXml_11_0_2
		obj.MarshalFromObject(*s.RibFilter)
		o.RibFilter = &obj
	}
	if s.Bgp != nil {
		var obj vrfBgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.RoutingTable != nil {
		var obj vrfRoutingTableXml_11_0_2
		obj.MarshalFromObject(*s.RoutingTable)
		o.RoutingTable = &obj
	}
	if s.Ospf != nil {
		var obj vrfOspfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj vrfOspfv3Xml_11_0_2
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Ecmp != nil {
		var obj vrfEcmpXml_11_0_2
		obj.MarshalFromObject(*s.Ecmp)
		o.Ecmp = &obj
	}
	if s.Multicast != nil {
		var obj vrfMulticastXml_11_0_2
		obj.MarshalFromObject(*s.Multicast)
		o.Multicast = &obj
	}
	if s.Rip != nil {
		var obj vrfRipXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
}

func (o vrfXml_11_0_2) UnmarshalToObject() (*Vrf, error) {
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var adminDistsVal *VrfAdminDists
	if o.AdminDists != nil {
		obj, err := o.AdminDists.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adminDistsVal = obj
	}
	var ribFilterVal *VrfRibFilter
	if o.RibFilter != nil {
		obj, err := o.RibFilter.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ribFilterVal = obj
	}
	var bgpVal *VrfBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var routingTableVal *VrfRoutingTable
	if o.RoutingTable != nil {
		obj, err := o.RoutingTable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingTableVal = obj
	}
	var ospfVal *VrfOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *VrfOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var ecmpVal *VrfEcmp
	if o.Ecmp != nil {
		obj, err := o.Ecmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ecmpVal = obj
	}
	var multicastVal *VrfMulticast
	if o.Multicast != nil {
		obj, err := o.Multicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastVal = obj
	}
	var ripVal *VrfRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Vrf{
		Name:         o.Name,
		Interface:    interfaceVal,
		AdminDists:   adminDistsVal,
		RibFilter:    ribFilterVal,
		Bgp:          bgpVal,
		RoutingTable: routingTableVal,
		Ospf:         ospfVal,
		Ospfv3:       ospfv3Val,
		Ecmp:         ecmpVal,
		Multicast:    multicastVal,
		Rip:          ripVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *vrfAdminDistsXml_11_0_2) MarshalFromObject(s VrfAdminDists) {
	o.Static = s.Static
	o.StaticIpv6 = s.StaticIpv6
	o.OspfInter = s.OspfInter
	o.OspfIntra = s.OspfIntra
	o.OspfExt = s.OspfExt
	o.Ospfv3Inter = s.Ospfv3Inter
	o.Ospfv3Intra = s.Ospfv3Intra
	o.Ospfv3Ext = s.Ospfv3Ext
	o.BgpInternal = s.BgpInternal
	o.BgpExternal = s.BgpExternal
	o.BgpLocal = s.BgpLocal
	o.Rip = s.Rip
	o.Misc = s.Misc
}

func (o vrfAdminDistsXml_11_0_2) UnmarshalToObject() (*VrfAdminDists, error) {

	result := &VrfAdminDists{
		Static:      o.Static,
		StaticIpv6:  o.StaticIpv6,
		OspfInter:   o.OspfInter,
		OspfIntra:   o.OspfIntra,
		OspfExt:     o.OspfExt,
		Ospfv3Inter: o.Ospfv3Inter,
		Ospfv3Intra: o.Ospfv3Intra,
		Ospfv3Ext:   o.Ospfv3Ext,
		BgpInternal: o.BgpInternal,
		BgpExternal: o.BgpExternal,
		BgpLocal:    o.BgpLocal,
		Rip:         o.Rip,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterXml_11_0_2) MarshalFromObject(s VrfRibFilter) {
	if s.Ipv4 != nil {
		var obj vrfRibFilterIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfRibFilterIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterXml_11_0_2) UnmarshalToObject() (*VrfRibFilter, error) {
	var ipv4Val *VrfRibFilterIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfRibFilterIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfRibFilter{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4Xml_11_0_2) MarshalFromObject(s VrfRibFilterIpv4) {
	if s.Static != nil {
		var obj vrfRibFilterIpv4StaticXml_11_0_2
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	if s.Bgp != nil {
		var obj vrfRibFilterIpv4BgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj vrfRibFilterIpv4OspfXml_11_0_2
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Rip != nil {
		var obj vrfRibFilterIpv4RipXml_11_0_2
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4Xml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv4, error) {
	var staticVal *VrfRibFilterIpv4Static
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}
	var bgpVal *VrfRibFilterIpv4Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *VrfRibFilterIpv4Ospf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ripVal *VrfRibFilterIpv4Rip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &VrfRibFilterIpv4{
		Static: staticVal,
		Bgp:    bgpVal,
		Ospf:   ospfVal,
		Rip:    ripVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4StaticXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv4Static) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4StaticXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv4Static, error) {

	result := &VrfRibFilterIpv4Static{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4BgpXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv4Bgp) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4BgpXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv4Bgp, error) {

	result := &VrfRibFilterIpv4Bgp{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4OspfXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv4Ospf) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4OspfXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv4Ospf, error) {

	result := &VrfRibFilterIpv4Ospf{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv4RipXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv4Rip) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv4RipXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv4Rip, error) {

	result := &VrfRibFilterIpv4Rip{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6Xml_11_0_2) MarshalFromObject(s VrfRibFilterIpv6) {
	if s.Static != nil {
		var obj vrfRibFilterIpv6StaticXml_11_0_2
		obj.MarshalFromObject(*s.Static)
		o.Static = &obj
	}
	if s.Bgp != nil {
		var obj vrfRibFilterIpv6BgpXml_11_0_2
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospfv3 != nil {
		var obj vrfRibFilterIpv6Ospfv3Xml_11_0_2
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6Xml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv6, error) {
	var staticVal *VrfRibFilterIpv6Static
	if o.Static != nil {
		obj, err := o.Static.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticVal = obj
	}
	var bgpVal *VrfRibFilterIpv6Bgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfv3Val *VrfRibFilterIpv6Ospfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}

	result := &VrfRibFilterIpv6{
		Static: staticVal,
		Bgp:    bgpVal,
		Ospfv3: ospfv3Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6StaticXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv6Static) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6StaticXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv6Static, error) {

	result := &VrfRibFilterIpv6Static{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6BgpXml_11_0_2) MarshalFromObject(s VrfRibFilterIpv6Bgp) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6BgpXml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv6Bgp, error) {

	result := &VrfRibFilterIpv6Bgp{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfRibFilterIpv6Ospfv3Xml_11_0_2) MarshalFromObject(s VrfRibFilterIpv6Ospfv3) {
	o.RouteMap = s.RouteMap
	o.Misc = s.Misc
}

func (o vrfRibFilterIpv6Ospfv3Xml_11_0_2) UnmarshalToObject() (*VrfRibFilterIpv6Ospfv3, error) {

	result := &VrfRibFilterIpv6Ospfv3{
		RouteMap: o.RouteMap,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfBgpXml_11_0_2) MarshalFromObject(s VrfBgp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RouterId = s.RouterId
	o.LocalAs = s.LocalAs
	o.InstallRoute = util.YesNo(s.InstallRoute, nil)
	o.EnforceFirstAs = util.YesNo(s.EnforceFirstAs, nil)
	o.FastExternalFailover = util.YesNo(s.FastExternalFailover, nil)
	o.EcmpMultiAs = util.YesNo(s.EcmpMultiAs, nil)
	o.DefaultLocalPreference = s.DefaultLocalPreference
	o.GracefulShutdown = util.YesNo(s.GracefulShutdown, nil)
	o.AlwaysAdvertiseNetworkRoute = util.YesNo(s.AlwaysAdvertiseNetworkRoute, nil)
	if s.Med != nil {
		var obj vrfBgpMedXml_11_0_2
		obj.MarshalFromObject(*s.Med)
		o.Med = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfBgpGracefulRestartXml_11_0_2
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.GlobalBfd != nil {
		var obj vrfBgpGlobalBfdXml_11_0_2
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.RedistributionProfile != nil {
		var obj vrfBgpRedistributionProfileXml_11_0_2
		obj.MarshalFromObject(*s.RedistributionProfile)
		o.RedistributionProfile = &obj
	}
	if s.AdvertiseNetwork != nil {
		var obj vrfBgpAdvertiseNetworkXml_11_0_2
		obj.MarshalFromObject(*s.AdvertiseNetwork)
		o.AdvertiseNetwork = &obj
	}
	if s.PeerGroup != nil {
		var objs []vrfBgpPeerGroupXml_11_0_2
		for _, elt := range s.PeerGroup {
			var obj vrfBgpPeerGroupXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.PeerGroup = &vrfBgpPeerGroupContainerXml_11_0_2{Entries: objs}
	}
	if s.AggregateRoutes != nil {
		var objs []vrfBgpAggregateRoutesXml_11_0_2
		for _, elt := range s.AggregateRoutes {
			var obj vrfBgpAggregateRoutesXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AggregateRoutes = &vrfBgpAggregateRoutesContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpXml_11_0_2) UnmarshalToObject() (*VrfBgp, error) {
	var medVal *VrfBgpMed
	if o.Med != nil {
		obj, err := o.Med.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		medVal = obj
	}
	var gracefulRestartVal *VrfBgpGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var globalBfdVal *VrfBgpGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var redistributionProfileVal *VrfBgpRedistributionProfile
	if o.RedistributionProfile != nil {
		obj, err := o.RedistributionProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redistributionProfileVal = obj
	}
	var advertiseNetworkVal *VrfBgpAdvertiseNetwork
	if o.AdvertiseNetwork != nil {
		obj, err := o.AdvertiseNetwork.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseNetworkVal = obj
	}
	var peerGroupVal []VrfBgpPeerGroup
	if o.PeerGroup != nil {
		for _, elt := range o.PeerGroup.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerGroupVal = append(peerGroupVal, *obj)
		}
	}
	var aggregateRoutesVal []VrfBgpAggregateRoutes
	if o.AggregateRoutes != nil {
		for _, elt := range o.AggregateRoutes.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			aggregateRoutesVal = append(aggregateRoutesVal, *obj)
		}
	}

	result := &VrfBgp{
		Enable:                      util.AsBool(o.Enable, nil),
		RouterId:                    o.RouterId,
		LocalAs:                     o.LocalAs,
		InstallRoute:                util.AsBool(o.InstallRoute, nil),
		EnforceFirstAs:              util.AsBool(o.EnforceFirstAs, nil),
		FastExternalFailover:        util.AsBool(o.FastExternalFailover, nil),
		EcmpMultiAs:                 util.AsBool(o.EcmpMultiAs, nil),
		DefaultLocalPreference:      o.DefaultLocalPreference,
		GracefulShutdown:            util.AsBool(o.GracefulShutdown, nil),
		AlwaysAdvertiseNetworkRoute: util.AsBool(o.AlwaysAdvertiseNetworkRoute, nil),
		Med:                         medVal,
		GracefulRestart:             gracefulRestartVal,
		GlobalBfd:                   globalBfdVal,
		RedistributionProfile:       redistributionProfileVal,
		AdvertiseNetwork:            advertiseNetworkVal,
		PeerGroup:                   peerGroupVal,
		AggregateRoutes:             aggregateRoutesVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfBgpMedXml_11_0_2) MarshalFromObject(s VrfBgpMed) {
	o.AlwaysCompareMed = util.YesNo(s.AlwaysCompareMed, nil)
	o.DeterministicMedComparison = util.YesNo(s.DeterministicMedComparison, nil)
	o.Misc = s.Misc
}

func (o vrfBgpMedXml_11_0_2) UnmarshalToObject() (*VrfBgpMed, error) {

	result := &VrfBgpMed{
		AlwaysCompareMed:           util.AsBool(o.AlwaysCompareMed, nil),
		DeterministicMedComparison: util.AsBool(o.DeterministicMedComparison, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *vrfBgpGracefulRestartXml_11_0_2) MarshalFromObject(s VrfBgpGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.StaleRouteTime = s.StaleRouteTime
	o.MaxPeerRestartTime = s.MaxPeerRestartTime
	o.LocalRestartTime = s.LocalRestartTime
	o.Misc = s.Misc
}

func (o vrfBgpGracefulRestartXml_11_0_2) UnmarshalToObject() (*VrfBgpGracefulRestart, error) {

	result := &VrfBgpGracefulRestart{
		Enable:             util.AsBool(o.Enable, nil),
		StaleRouteTime:     o.StaleRouteTime,
		MaxPeerRestartTime: o.MaxPeerRestartTime,
		LocalRestartTime:   o.LocalRestartTime,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfBgpGlobalBfdXml_11_0_2) MarshalFromObject(s VrfBgpGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfBgpGlobalBfdXml_11_0_2) UnmarshalToObject() (*VrfBgpGlobalBfd, error) {

	result := &VrfBgpGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileXml_11_0_2) MarshalFromObject(s VrfBgpRedistributionProfile) {
	if s.Ipv4 != nil {
		var obj vrfBgpRedistributionProfileIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpRedistributionProfileIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileXml_11_0_2) UnmarshalToObject() (*VrfBgpRedistributionProfile, error) {
	var ipv4Val *VrfBgpRedistributionProfileIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpRedistributionProfileIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpRedistributionProfile{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileIpv4Xml_11_0_2) MarshalFromObject(s VrfBgpRedistributionProfileIpv4) {
	o.Unicast = s.Unicast
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileIpv4Xml_11_0_2) UnmarshalToObject() (*VrfBgpRedistributionProfileIpv4, error) {

	result := &VrfBgpRedistributionProfileIpv4{
		Unicast: o.Unicast,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpRedistributionProfileIpv6Xml_11_0_2) MarshalFromObject(s VrfBgpRedistributionProfileIpv6) {
	o.Unicast = s.Unicast
	o.Misc = s.Misc
}

func (o vrfBgpRedistributionProfileIpv6Xml_11_0_2) UnmarshalToObject() (*VrfBgpRedistributionProfileIpv6, error) {

	result := &VrfBgpRedistributionProfileIpv6{
		Unicast: o.Unicast,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkXml_11_0_2) MarshalFromObject(s VrfBgpAdvertiseNetwork) {
	if s.Ipv4 != nil {
		var obj vrfBgpAdvertiseNetworkIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpAdvertiseNetworkIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkXml_11_0_2) UnmarshalToObject() (*VrfBgpAdvertiseNetwork, error) {
	var ipv4Val *VrfBgpAdvertiseNetworkIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpAdvertiseNetworkIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpAdvertiseNetwork{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv4Xml_11_0_2) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv4) {
	if s.Network != nil {
		var objs []vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2
		for _, elt := range s.Network {
			var obj vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Network = &vrfBgpAdvertiseNetworkIpv4NetworkContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv4Xml_11_0_2) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv4, error) {
	var networkVal []VrfBgpAdvertiseNetworkIpv4Network
	if o.Network != nil {
		for _, elt := range o.Network.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			networkVal = append(networkVal, *obj)
		}
	}

	result := &VrfBgpAdvertiseNetworkIpv4{
		Network: networkVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv4Network) {
	o.Name = s.Name
	o.Unicast = util.YesNo(s.Unicast, nil)
	o.Multicast = util.YesNo(s.Multicast, nil)
	o.Backdoor = util.YesNo(s.Backdoor, nil)
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv4NetworkXml_11_0_2) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv4Network, error) {

	result := &VrfBgpAdvertiseNetworkIpv4Network{
		Name:      o.Name,
		Unicast:   util.AsBool(o.Unicast, nil),
		Multicast: util.AsBool(o.Multicast, nil),
		Backdoor:  util.AsBool(o.Backdoor, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv6Xml_11_0_2) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv6) {
	if s.Network != nil {
		var objs []vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2
		for _, elt := range s.Network {
			var obj vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Network = &vrfBgpAdvertiseNetworkIpv6NetworkContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv6Xml_11_0_2) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv6, error) {
	var networkVal []VrfBgpAdvertiseNetworkIpv6Network
	if o.Network != nil {
		for _, elt := range o.Network.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			networkVal = append(networkVal, *obj)
		}
	}

	result := &VrfBgpAdvertiseNetworkIpv6{
		Network: networkVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2) MarshalFromObject(s VrfBgpAdvertiseNetworkIpv6Network) {
	o.Name = s.Name
	o.Unicast = util.YesNo(s.Unicast, nil)
	o.Misc = s.Misc
}

func (o vrfBgpAdvertiseNetworkIpv6NetworkXml_11_0_2) UnmarshalToObject() (*VrfBgpAdvertiseNetworkIpv6Network, error) {

	result := &VrfBgpAdvertiseNetworkIpv6Network{
		Name:    o.Name,
		Unicast: util.AsBool(o.Unicast, nil),
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroup) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Type != nil {
		var obj vrfBgpPeerGroupTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.AddressFamily != nil {
		var obj vrfBgpPeerGroupAddressFamilyXml_11_0_2
		obj.MarshalFromObject(*s.AddressFamily)
		o.AddressFamily = &obj
	}
	if s.FilteringProfile != nil {
		var obj vrfBgpPeerGroupFilteringProfileXml_11_0_2
		obj.MarshalFromObject(*s.FilteringProfile)
		o.FilteringProfile = &obj
	}
	if s.ConnectionOptions != nil {
		var obj vrfBgpPeerGroupConnectionOptionsXml_11_0_2
		obj.MarshalFromObject(*s.ConnectionOptions)
		o.ConnectionOptions = &obj
	}
	if s.Peer != nil {
		var objs []vrfBgpPeerGroupPeerXml_11_0_2
		for _, elt := range s.Peer {
			var obj vrfBgpPeerGroupPeerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Peer = &vrfBgpPeerGroupPeerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroup, error) {
	var typeVal *VrfBgpPeerGroupType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var addressFamilyVal *VrfBgpPeerGroupAddressFamily
	if o.AddressFamily != nil {
		obj, err := o.AddressFamily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressFamilyVal = obj
	}
	var filteringProfileVal *VrfBgpPeerGroupFilteringProfile
	if o.FilteringProfile != nil {
		obj, err := o.FilteringProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filteringProfileVal = obj
	}
	var connectionOptionsVal *VrfBgpPeerGroupConnectionOptions
	if o.ConnectionOptions != nil {
		obj, err := o.ConnectionOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectionOptionsVal = obj
	}
	var peerVal []VrfBgpPeerGroupPeer
	if o.Peer != nil {
		for _, elt := range o.Peer.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerVal = append(peerVal, *obj)
		}
	}

	result := &VrfBgpPeerGroup{
		Name:              o.Name,
		Enable:            util.AsBool(o.Enable, nil),
		Type:              typeVal,
		AddressFamily:     addressFamilyVal,
		FilteringProfile:  filteringProfileVal,
		ConnectionOptions: connectionOptionsVal,
		Peer:              peerVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupType) {
	if s.Ibgp != nil {
		var obj vrfBgpPeerGroupTypeIbgpXml_11_0_2
		obj.MarshalFromObject(*s.Ibgp)
		o.Ibgp = &obj
	}
	if s.Ebgp != nil {
		var obj vrfBgpPeerGroupTypeEbgpXml_11_0_2
		obj.MarshalFromObject(*s.Ebgp)
		o.Ebgp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupType, error) {
	var ibgpVal *VrfBgpPeerGroupTypeIbgp
	if o.Ibgp != nil {
		obj, err := o.Ibgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ibgpVal = obj
	}
	var ebgpVal *VrfBgpPeerGroupTypeEbgp
	if o.Ebgp != nil {
		obj, err := o.Ebgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ebgpVal = obj
	}

	result := &VrfBgpPeerGroupType{
		Ibgp: ibgpVal,
		Ebgp: ebgpVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeIbgpXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupTypeIbgp) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeIbgpXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupTypeIbgp, error) {

	result := &VrfBgpPeerGroupTypeIbgp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupTypeEbgpXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupTypeEbgp) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupTypeEbgpXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupTypeEbgp, error) {

	result := &VrfBgpPeerGroupTypeEbgp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupAddressFamilyXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupAddressFamily) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupAddressFamilyXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupAddressFamily, error) {

	result := &VrfBgpPeerGroupAddressFamily{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupFilteringProfileXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupFilteringProfile) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupFilteringProfileXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupFilteringProfile, error) {

	result := &VrfBgpPeerGroupFilteringProfile{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupConnectionOptionsXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupConnectionOptions) {
	o.Timers = s.Timers
	o.Multihop = s.Multihop
	o.Authentication = s.Authentication
	o.Dampening = s.Dampening
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupConnectionOptionsXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupConnectionOptions, error) {

	result := &VrfBgpPeerGroupConnectionOptions{
		Timers:         o.Timers,
		Multihop:       o.Multihop,
		Authentication: o.Authentication,
		Dampening:      o.Dampening,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeer) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.PeerAs = s.PeerAs
	o.EnableSenderSideLoopDetection = util.YesNo(s.EnableSenderSideLoopDetection, nil)
	if s.Inherit != nil {
		var obj vrfBgpPeerGroupPeerInheritXml_11_0_2
		obj.MarshalFromObject(*s.Inherit)
		o.Inherit = &obj
	}
	if s.LocalAddress != nil {
		var obj vrfBgpPeerGroupPeerLocalAddressXml_11_0_2
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj vrfBgpPeerGroupPeerPeerAddressXml_11_0_2
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	if s.ConnectionOptions != nil {
		var obj vrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2
		obj.MarshalFromObject(*s.ConnectionOptions)
		o.ConnectionOptions = &obj
	}
	if s.Bfd != nil {
		var obj vrfBgpPeerGroupPeerBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeer, error) {
	var inheritVal *VrfBgpPeerGroupPeerInherit
	if o.Inherit != nil {
		obj, err := o.Inherit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritVal = obj
	}
	var localAddressVal *VrfBgpPeerGroupPeerLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *VrfBgpPeerGroupPeerPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var connectionOptionsVal *VrfBgpPeerGroupPeerConnectionOptions
	if o.ConnectionOptions != nil {
		obj, err := o.ConnectionOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectionOptionsVal = obj
	}
	var bfdVal *VrfBgpPeerGroupPeerBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfBgpPeerGroupPeer{
		Name:                          o.Name,
		Enable:                        util.AsBool(o.Enable, nil),
		Passive:                       util.AsBool(o.Passive, nil),
		PeerAs:                        o.PeerAs,
		EnableSenderSideLoopDetection: util.AsBool(o.EnableSenderSideLoopDetection, nil),
		Inherit:                       inheritVal,
		LocalAddress:                  localAddressVal,
		PeerAddress:                   peerAddressVal,
		ConnectionOptions:             connectionOptionsVal,
		Bfd:                           bfdVal,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerInherit) {
	if s.Yes != nil {
		var obj vrfBgpPeerGroupPeerInheritYesXml_11_0_2
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	if s.No != nil {
		var obj vrfBgpPeerGroupPeerInheritNoXml_11_0_2
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerInherit, error) {
	var yesVal *VrfBgpPeerGroupPeerInheritYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}
	var noVal *VrfBgpPeerGroupPeerInheritNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}

	result := &VrfBgpPeerGroupPeerInherit{
		Yes:  yesVal,
		No:   noVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritYesXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerInheritYes) {
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritYesXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritYes, error) {

	result := &VrfBgpPeerGroupPeerInheritYes{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNo) {
	if s.AddressFamily != nil {
		var obj vrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2
		obj.MarshalFromObject(*s.AddressFamily)
		o.AddressFamily = &obj
	}
	if s.FilteringProfile != nil {
		var obj vrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2
		obj.MarshalFromObject(*s.FilteringProfile)
		o.FilteringProfile = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNo, error) {
	var addressFamilyVal *VrfBgpPeerGroupPeerInheritNoAddressFamily
	if o.AddressFamily != nil {
		obj, err := o.AddressFamily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addressFamilyVal = obj
	}
	var filteringProfileVal *VrfBgpPeerGroupPeerInheritNoFilteringProfile
	if o.FilteringProfile != nil {
		obj, err := o.FilteringProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filteringProfileVal = obj
	}

	result := &VrfBgpPeerGroupPeerInheritNo{
		AddressFamily:    addressFamilyVal,
		FilteringProfile: filteringProfileVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNoAddressFamily) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoAddressFamilyXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNoAddressFamily, error) {

	result := &VrfBgpPeerGroupPeerInheritNoAddressFamily{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerInheritNoFilteringProfile) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerInheritNoFilteringProfileXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerInheritNoFilteringProfile, error) {

	result := &VrfBgpPeerGroupPeerInheritNoFilteringProfile{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerLocalAddressXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerLocalAddress) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerLocalAddressXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerLocalAddress, error) {

	result := &VrfBgpPeerGroupPeerLocalAddress{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerPeerAddressXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerPeerAddress) {
	o.Ip = s.Ip
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerPeerAddressXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerPeerAddress, error) {

	result := &VrfBgpPeerGroupPeerPeerAddress{
		Ip:   o.Ip,
		Fqdn: o.Fqdn,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerConnectionOptions) {
	o.Timers = s.Timers
	o.Multihop = s.Multihop
	o.Authentication = s.Authentication
	o.Dampening = s.Dampening
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerConnectionOptionsXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerConnectionOptions, error) {

	result := &VrfBgpPeerGroupPeerConnectionOptions{
		Timers:         o.Timers,
		Multihop:       o.Multihop,
		Authentication: o.Authentication,
		Dampening:      o.Dampening,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfBgpPeerGroupPeerBfdXml_11_0_2) MarshalFromObject(s VrfBgpPeerGroupPeerBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfBgpPeerGroupPeerBfdXml_11_0_2) UnmarshalToObject() (*VrfBgpPeerGroupPeerBfd, error) {

	result := &VrfBgpPeerGroupPeerBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesXml_11_0_2) MarshalFromObject(s VrfBgpAggregateRoutes) {
	o.Name = s.Name
	o.Description = s.Description
	o.Enable = util.YesNo(s.Enable, nil)
	o.SummaryOnly = util.YesNo(s.SummaryOnly, nil)
	o.AsSet = util.YesNo(s.AsSet, nil)
	o.SameMed = util.YesNo(s.SameMed, nil)
	if s.Type != nil {
		var obj vrfBgpAggregateRoutesTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesXml_11_0_2) UnmarshalToObject() (*VrfBgpAggregateRoutes, error) {
	var typeVal *VrfBgpAggregateRoutesType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &VrfBgpAggregateRoutes{
		Name:        o.Name,
		Description: o.Description,
		Enable:      util.AsBool(o.Enable, nil),
		SummaryOnly: util.AsBool(o.SummaryOnly, nil),
		AsSet:       util.AsBool(o.AsSet, nil),
		SameMed:     util.AsBool(o.SameMed, nil),
		Type:        typeVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeXml_11_0_2) MarshalFromObject(s VrfBgpAggregateRoutesType) {
	if s.Ipv4 != nil {
		var obj vrfBgpAggregateRoutesTypeIpv4Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfBgpAggregateRoutesTypeIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeXml_11_0_2) UnmarshalToObject() (*VrfBgpAggregateRoutesType, error) {
	var ipv4Val *VrfBgpAggregateRoutesTypeIpv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *VrfBgpAggregateRoutesTypeIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfBgpAggregateRoutesType{
		Ipv4: ipv4Val,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeIpv4Xml_11_0_2) MarshalFromObject(s VrfBgpAggregateRoutesTypeIpv4) {
	o.SummaryPrefix = s.SummaryPrefix
	o.SuppressMap = s.SuppressMap
	o.AttributeMap = s.AttributeMap
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeIpv4Xml_11_0_2) UnmarshalToObject() (*VrfBgpAggregateRoutesTypeIpv4, error) {

	result := &VrfBgpAggregateRoutesTypeIpv4{
		SummaryPrefix: o.SummaryPrefix,
		SuppressMap:   o.SuppressMap,
		AttributeMap:  o.AttributeMap,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfBgpAggregateRoutesTypeIpv6Xml_11_0_2) MarshalFromObject(s VrfBgpAggregateRoutesTypeIpv6) {
	o.SummaryPrefix = s.SummaryPrefix
	o.SuppressMap = s.SuppressMap
	o.AttributeMap = s.AttributeMap
	o.Misc = s.Misc
}

func (o vrfBgpAggregateRoutesTypeIpv6Xml_11_0_2) UnmarshalToObject() (*VrfBgpAggregateRoutesTypeIpv6, error) {

	result := &VrfBgpAggregateRoutesTypeIpv6{
		SummaryPrefix: o.SummaryPrefix,
		SuppressMap:   o.SuppressMap,
		AttributeMap:  o.AttributeMap,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableXml_11_0_2) MarshalFromObject(s VrfRoutingTable) {
	if s.Ip != nil {
		var obj vrfRoutingTableIpXml_11_0_2
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	if s.Ipv6 != nil {
		var obj vrfRoutingTableIpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableXml_11_0_2) UnmarshalToObject() (*VrfRoutingTable, error) {
	var ipVal *VrfRoutingTableIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}
	var ipv6Val *VrfRoutingTableIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &VrfRoutingTable{
		Ip:   ipVal,
		Ipv6: ipv6Val,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpXml_11_0_2) MarshalFromObject(s VrfRoutingTableIp) {
	if s.StaticRoute != nil {
		var objs []vrfRoutingTableIpStaticRouteXml_11_0_2
		for _, elt := range s.StaticRoute {
			var obj vrfRoutingTableIpStaticRouteXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfRoutingTableIpStaticRouteContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIp, error) {
	var staticRouteVal []VrfRoutingTableIpStaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}

	result := &VrfRoutingTableIp{
		StaticRoute: staticRouteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.AdminDist = s.AdminDist
	o.Metric = s.Metric
	if s.Nexthop != nil {
		var obj vrfRoutingTableIpStaticRouteNexthopXml_11_0_2
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	if s.Bfd != nil {
		var obj vrfRoutingTableIpStaticRouteBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.PathMonitor != nil {
		var obj vrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2
		obj.MarshalFromObject(*s.PathMonitor)
		o.PathMonitor = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRoute, error) {
	var nexthopVal *VrfRoutingTableIpStaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}
	var bfdVal *VrfRoutingTableIpStaticRouteBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var pathMonitorVal *VrfRoutingTableIpStaticRoutePathMonitor
	if o.PathMonitor != nil {
		obj, err := o.PathMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pathMonitorVal = obj
	}

	result := &VrfRoutingTableIpStaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		AdminDist:   o.AdminDist,
		Metric:      o.Metric,
		Nexthop:     nexthopVal,
		Bfd:         bfdVal,
		PathMonitor: pathMonitorVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteNexthopXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRouteNexthop) {
	if s.Discard != nil {
		var obj vrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	o.IpAddress = s.IpAddress
	o.NextLr = s.NextLr
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteNexthopXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteNexthop, error) {
	var discardVal *VrfRoutingTableIpStaticRouteNexthopDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}

	result := &VrfRoutingTableIpStaticRouteNexthop{
		Discard:   discardVal,
		IpAddress: o.IpAddress,
		NextLr:    o.NextLr,
		Fqdn:      o.Fqdn,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRouteNexthopDiscard) {
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteNexthopDiscardXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteNexthopDiscard, error) {

	result := &VrfRoutingTableIpStaticRouteNexthopDiscard{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRouteBfdXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRouteBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRouteBfdXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRouteBfd, error) {

	result := &VrfRoutingTableIpStaticRouteBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRoutePathMonitor) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FailureCondition = s.FailureCondition
	o.HoldTime = s.HoldTime
	if s.MonitorDestinations != nil {
		var objs []vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2
		for _, elt := range s.MonitorDestinations {
			var obj vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MonitorDestinations = &vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRoutePathMonitorXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRoutePathMonitor, error) {
	var monitorDestinationsVal []VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations
	if o.MonitorDestinations != nil {
		for _, elt := range o.MonitorDestinations.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			monitorDestinationsVal = append(monitorDestinationsVal, *obj)
		}
	}

	result := &VrfRoutingTableIpStaticRoutePathMonitor{
		Enable:              util.AsBool(o.Enable, nil),
		FailureCondition:    o.FailureCondition,
		HoldTime:            o.HoldTime,
		MonitorDestinations: monitorDestinationsVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Source = s.Source
	o.Destination = s.Destination
	o.Interval = s.Interval
	o.Count = s.Count
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations, error) {

	result := &VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations{
		Name:        o.Name,
		Enable:      util.AsBool(o.Enable, nil),
		Source:      o.Source,
		Destination: o.Destination,
		Interval:    o.Interval,
		Count:       o.Count,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6Xml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6) {
	if s.StaticRoute != nil {
		var objs []vrfRoutingTableIpv6StaticRouteXml_11_0_2
		for _, elt := range s.StaticRoute {
			var obj vrfRoutingTableIpv6StaticRouteXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfRoutingTableIpv6StaticRouteContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6Xml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6, error) {
	var staticRouteVal []VrfRoutingTableIpv6StaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}

	result := &VrfRoutingTableIpv6{
		StaticRoute: staticRouteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.AdminDist = s.AdminDist
	o.Metric = s.Metric
	if s.Nexthop != nil {
		var obj vrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	if s.Bfd != nil {
		var obj vrfRoutingTableIpv6StaticRouteBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.PathMonitor != nil {
		var obj vrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2
		obj.MarshalFromObject(*s.PathMonitor)
		o.PathMonitor = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoute, error) {
	var nexthopVal *VrfRoutingTableIpv6StaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}
	var bfdVal *VrfRoutingTableIpv6StaticRouteBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var pathMonitorVal *VrfRoutingTableIpv6StaticRoutePathMonitor
	if o.PathMonitor != nil {
		obj, err := o.PathMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pathMonitorVal = obj
	}

	result := &VrfRoutingTableIpv6StaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		AdminDist:   o.AdminDist,
		Metric:      o.Metric,
		Nexthop:     nexthopVal,
		Bfd:         bfdVal,
		PathMonitor: pathMonitorVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteNexthop) {
	if s.Discard != nil {
		var obj vrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	o.Ipv6Address = s.Ipv6Address
	o.Fqdn = s.Fqdn
	o.NextLr = s.NextLr
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteNexthopXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteNexthop, error) {
	var discardVal *VrfRoutingTableIpv6StaticRouteNexthopDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}

	result := &VrfRoutingTableIpv6StaticRouteNexthop{
		Discard:     discardVal,
		Ipv6Address: o.Ipv6Address,
		Fqdn:        o.Fqdn,
		NextLr:      o.NextLr,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteNexthopDiscard) {
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteNexthopDiscardXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteNexthopDiscard, error) {

	result := &VrfRoutingTableIpv6StaticRouteNexthopDiscard{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRouteBfdXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRouteBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRouteBfdXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRouteBfd, error) {

	result := &VrfRoutingTableIpv6StaticRouteBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRoutePathMonitor) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FailureCondition = s.FailureCondition
	o.HoldTime = s.HoldTime
	if s.MonitorDestinations != nil {
		var objs []vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2
		for _, elt := range s.MonitorDestinations {
			var obj vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MonitorDestinations = &vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRoutePathMonitorXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoutePathMonitor, error) {
	var monitorDestinationsVal []VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations
	if o.MonitorDestinations != nil {
		for _, elt := range o.MonitorDestinations.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			monitorDestinationsVal = append(monitorDestinationsVal, *obj)
		}
	}

	result := &VrfRoutingTableIpv6StaticRoutePathMonitor{
		Enable:              util.AsBool(o.Enable, nil),
		FailureCondition:    o.FailureCondition,
		HoldTime:            o.HoldTime,
		MonitorDestinations: monitorDestinationsVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2) MarshalFromObject(s VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Source = s.Source
	o.Destination = s.Destination
	o.Interval = s.Interval
	o.Count = s.Count
	o.Misc = s.Misc
}

func (o vrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml_11_0_2) UnmarshalToObject() (*VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations, error) {

	result := &VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{
		Name:        o.Name,
		Enable:      util.AsBool(o.Enable, nil),
		Source:      o.Source,
		Destination: o.Destination,
		Interval:    o.Interval,
		Count:       o.Count,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfXml_11_0_2) MarshalFromObject(s VrfOspf) {
	o.RouterId = s.RouterId
	o.Enable = util.YesNo(s.Enable, nil)
	o.Rfc1583 = util.YesNo(s.Rfc1583, nil)
	o.SpfTimer = s.SpfTimer
	o.GlobalIfTimer = s.GlobalIfTimer
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfOspfGlobalBfdXml_11_0_2
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfOspfGracefulRestartXml_11_0_2
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.Area != nil {
		var objs []vrfOspfAreaXml_11_0_2
		for _, elt := range s.Area {
			var obj vrfOspfAreaXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &vrfOspfAreaContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfXml_11_0_2) UnmarshalToObject() (*VrfOspf, error) {
	var globalBfdVal *VrfOspfGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *VrfOspfGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var areaVal []VrfOspfArea
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}

	result := &VrfOspf{
		RouterId:              o.RouterId,
		Enable:                util.AsBool(o.Enable, nil),
		Rfc1583:               util.AsBool(o.Rfc1583, nil),
		SpfTimer:              o.SpfTimer,
		GlobalIfTimer:         o.GlobalIfTimer,
		RedistributionProfile: o.RedistributionProfile,
		GlobalBfd:             globalBfdVal,
		GracefulRestart:       gracefulRestartVal,
		Area:                  areaVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfOspfGlobalBfdXml_11_0_2) MarshalFromObject(s VrfOspfGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfGlobalBfdXml_11_0_2) UnmarshalToObject() (*VrfOspfGlobalBfd, error) {

	result := &VrfOspfGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfGracefulRestartXml_11_0_2) MarshalFromObject(s VrfOspfGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.Misc = s.Misc
}

func (o vrfOspfGracefulRestartXml_11_0_2) UnmarshalToObject() (*VrfOspfGracefulRestart, error) {

	result := &VrfOspfGracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaXml_11_0_2) MarshalFromObject(s VrfOspfArea) {
	o.Name = s.Name
	o.Authentication = s.Authentication
	if s.Type != nil {
		var obj vrfOspfAreaTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []vrfOspfAreaRangeXml_11_0_2
		for _, elt := range s.Range {
			var obj vrfOspfAreaRangeXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &vrfOspfAreaRangeContainerXml_11_0_2{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfOspfAreaInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfOspfAreaInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfOspfAreaInterfaceContainerXml_11_0_2{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []vrfOspfAreaVirtualLinkXml_11_0_2
		for _, elt := range s.VirtualLink {
			var obj vrfOspfAreaVirtualLinkXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &vrfOspfAreaVirtualLinkContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaXml_11_0_2) UnmarshalToObject() (*VrfOspfArea, error) {
	var typeVal *VrfOspfAreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []VrfOspfAreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []VrfOspfAreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []VrfOspfAreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &VrfOspfArea{
		Name:           o.Name,
		Authentication: o.Authentication,
		Type:           typeVal,
		Range:          rangeVal,
		Interface:      interfaceVal,
		VirtualLink:    virtualLinkVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeXml_11_0_2) MarshalFromObject(s VrfOspfAreaType) {
	if s.Normal != nil {
		var obj vrfOspfAreaTypeNormalXml_11_0_2
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj vrfOspfAreaTypeStubXml_11_0_2
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj vrfOspfAreaTypeNssaXml_11_0_2
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaType, error) {
	var normalVal *VrfOspfAreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *VrfOspfAreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *VrfOspfAreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &VrfOspfAreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNormalXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNormal) {
	if s.Abr != nil {
		var obj vrfOspfAreaTypeNormalAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNormalXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNormal, error) {
	var abrVal *VrfOspfAreaTypeNormalAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeNormal{
		Abr:  abrVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNormalAbrXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNormalAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNormalAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNormalAbr, error) {

	result := &VrfOspfAreaTypeNormalAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeStubXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeStub) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.Abr != nil {
		var obj vrfOspfAreaTypeStubAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeStubXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeStub, error) {
	var abrVal *VrfOspfAreaTypeStubAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeStub{
		NoSummary:          util.AsBool(o.NoSummary, nil),
		Abr:                abrVal,
		DefaultRouteMetric: o.DefaultRouteMetric,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeStubAbrXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeStubAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeStubAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeStubAbr, error) {

	result := &VrfOspfAreaTypeStubAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNssa) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.DefaultInformationOriginate != nil {
		var obj vrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2
		obj.MarshalFromObject(*s.DefaultInformationOriginate)
		o.DefaultInformationOriginate = &obj
	}
	if s.Abr != nil {
		var obj vrfOspfAreaTypeNssaAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNssa, error) {
	var defaultInformationOriginateVal *VrfOspfAreaTypeNssaDefaultInformationOriginate
	if o.DefaultInformationOriginate != nil {
		obj, err := o.DefaultInformationOriginate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultInformationOriginateVal = obj
	}
	var abrVal *VrfOspfAreaTypeNssaAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfAreaTypeNssa{
		NoSummary:                   util.AsBool(o.NoSummary, nil),
		DefaultInformationOriginate: defaultInformationOriginateVal,
		Abr:                         abrVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNssaDefaultInformationOriginate) {
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaDefaultInformationOriginateXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNssaDefaultInformationOriginate, error) {

	result := &VrfOspfAreaTypeNssaDefaultInformationOriginate{
		Metric:     o.Metric,
		MetricType: o.MetricType,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaAbrXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNssaAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	if s.NssaExtRange != nil {
		var objs []vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2
		for _, elt := range s.NssaExtRange {
			var obj vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &vrfOspfAreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNssaAbr, error) {
	var nssaExtRangeVal []VrfOspfAreaTypeNssaAbrNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &VrfOspfAreaTypeNssaAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		NssaExtRange:       nssaExtRangeVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2) MarshalFromObject(s VrfOspfAreaTypeNssaAbrNssaExtRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfAreaTypeNssaAbrNssaExtRangeXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaTypeNssaAbrNssaExtRange, error) {

	result := &VrfOspfAreaTypeNssaAbrNssaExtRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaRangeXml_11_0_2) MarshalFromObject(s VrfOspfAreaRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfAreaRangeXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaRange, error) {

	result := &VrfOspfAreaRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.MtuIgnore = util.YesNo(s.MtuIgnore, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.Priority = s.Priority
	o.Metric = s.Metric
	o.Authentication = s.Authentication
	o.Timing = s.Timing
	if s.LinkType != nil {
		var obj vrfOspfAreaInterfaceLinkTypeXml_11_0_2
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Bfd != nil {
		var obj vrfOspfAreaInterfaceBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterface, error) {
	var linkTypeVal *VrfOspfAreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var bfdVal *VrfOspfAreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfAreaInterface{
		Name:           o.Name,
		Enable:         util.AsBool(o.Enable, nil),
		MtuIgnore:      util.AsBool(o.MtuIgnore, nil),
		Passive:        util.AsBool(o.Passive, nil),
		Priority:       o.Priority,
		Metric:         o.Metric,
		Authentication: o.Authentication,
		Timing:         o.Timing,
		LinkType:       linkTypeVal,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj vrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj vrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj vrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkType, error) {
	var broadcastVal *VrfOspfAreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *VrfOspfAreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *VrfOspfAreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &VrfOspfAreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeBroadcastXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeBroadcast, error) {

	result := &VrfOspfAreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2pXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2p, error) {

	result := &VrfOspfAreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2mp) {
	if s.Neighbor != nil {
		var objs []vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &vrfOspfAreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2mpXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2mp, error) {
	var neighborVal []VrfOspfAreaInterfaceLinkTypeP2mpNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &VrfOspfAreaInterfaceLinkTypeP2mp{
		Neighbor: neighborVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceLinkTypeP2mpNeighbor) {
	o.Name = s.Name
	o.Priority = s.Priority
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceLinkTypeP2mpNeighborXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceLinkTypeP2mpNeighbor, error) {

	result := &VrfOspfAreaInterfaceLinkTypeP2mpNeighbor{
		Name:     o.Name,
		Priority: o.Priority,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaInterfaceBfdXml_11_0_2) MarshalFromObject(s VrfOspfAreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfAreaInterfaceBfdXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaInterfaceBfd, error) {

	result := &VrfOspfAreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaVirtualLinkXml_11_0_2) MarshalFromObject(s VrfOspfAreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.Timing = s.Timing
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj vrfOspfAreaVirtualLinkBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfAreaVirtualLinkXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaVirtualLink, error) {
	var bfdVal *VrfOspfAreaVirtualLinkBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfAreaVirtualLink{
		Name:           o.Name,
		NeighborId:     o.NeighborId,
		TransitAreaId:  o.TransitAreaId,
		Enable:         util.AsBool(o.Enable, nil),
		InstanceId:     o.InstanceId,
		Timing:         o.Timing,
		Authentication: o.Authentication,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfAreaVirtualLinkBfdXml_11_0_2) MarshalFromObject(s VrfOspfAreaVirtualLinkBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfAreaVirtualLinkBfdXml_11_0_2) UnmarshalToObject() (*VrfOspfAreaVirtualLinkBfd, error) {

	result := &VrfOspfAreaVirtualLinkBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3Xml_11_0_2) MarshalFromObject(s VrfOspfv3) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RouterId = s.RouterId
	o.DisableTransitTraffic = util.YesNo(s.DisableTransitTraffic, nil)
	o.SpfTimer = s.SpfTimer
	o.GlobalIfTimer = s.GlobalIfTimer
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfOspfv3GlobalBfdXml_11_0_2
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj vrfOspfv3GracefulRestartXml_11_0_2
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.Area != nil {
		var objs []vrfOspfv3AreaXml_11_0_2
		for _, elt := range s.Area {
			var obj vrfOspfv3AreaXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &vrfOspfv3AreaContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3Xml_11_0_2) UnmarshalToObject() (*VrfOspfv3, error) {
	var globalBfdVal *VrfOspfv3GlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *VrfOspfv3GracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var areaVal []VrfOspfv3Area
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}

	result := &VrfOspfv3{
		Enable:                util.AsBool(o.Enable, nil),
		RouterId:              o.RouterId,
		DisableTransitTraffic: util.AsBool(o.DisableTransitTraffic, nil),
		SpfTimer:              o.SpfTimer,
		GlobalIfTimer:         o.GlobalIfTimer,
		RedistributionProfile: o.RedistributionProfile,
		GlobalBfd:             globalBfdVal,
		GracefulRestart:       gracefulRestartVal,
		Area:                  areaVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3GlobalBfdXml_11_0_2) MarshalFromObject(s VrfOspfv3GlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfv3GlobalBfdXml_11_0_2) UnmarshalToObject() (*VrfOspfv3GlobalBfd, error) {

	result := &VrfOspfv3GlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3GracefulRestartXml_11_0_2) MarshalFromObject(s VrfOspfv3GracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.Misc = s.Misc
}

func (o vrfOspfv3GracefulRestartXml_11_0_2) UnmarshalToObject() (*VrfOspfv3GracefulRestart, error) {

	result := &VrfOspfv3GracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaXml_11_0_2) MarshalFromObject(s VrfOspfv3Area) {
	o.Name = s.Name
	o.Authentication = s.Authentication
	if s.Type != nil {
		var obj vrfOspfv3AreaTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []vrfOspfv3AreaRangeXml_11_0_2
		for _, elt := range s.Range {
			var obj vrfOspfv3AreaRangeXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &vrfOspfv3AreaRangeContainerXml_11_0_2{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfOspfv3AreaInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfOspfv3AreaInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfOspfv3AreaInterfaceContainerXml_11_0_2{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []vrfOspfv3AreaVirtualLinkXml_11_0_2
		for _, elt := range s.VirtualLink {
			var obj vrfOspfv3AreaVirtualLinkXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &vrfOspfv3AreaVirtualLinkContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaXml_11_0_2) UnmarshalToObject() (*VrfOspfv3Area, error) {
	var typeVal *VrfOspfv3AreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []VrfOspfv3AreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []VrfOspfv3AreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []VrfOspfv3AreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &VrfOspfv3Area{
		Name:           o.Name,
		Authentication: o.Authentication,
		Type:           typeVal,
		Range:          rangeVal,
		Interface:      interfaceVal,
		VirtualLink:    virtualLinkVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaType) {
	if s.Normal != nil {
		var obj vrfOspfv3AreaTypeNormalXml_11_0_2
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj vrfOspfv3AreaTypeStubXml_11_0_2
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj vrfOspfv3AreaTypeNssaXml_11_0_2
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaType, error) {
	var normalVal *VrfOspfv3AreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *VrfOspfv3AreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *VrfOspfv3AreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &VrfOspfv3AreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNormalXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNormal) {
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeNormalAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNormalXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNormal, error) {
	var abrVal *VrfOspfv3AreaTypeNormalAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeNormal{
		Abr:  abrVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNormalAbrXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNormalAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNormalAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNormalAbr, error) {

	result := &VrfOspfv3AreaTypeNormalAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeStubXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeStub) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeStubAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeStubXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeStub, error) {
	var abrVal *VrfOspfv3AreaTypeStubAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeStub{
		NoSummary:          util.AsBool(o.NoSummary, nil),
		Abr:                abrVal,
		DefaultRouteMetric: o.DefaultRouteMetric,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeStubAbrXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeStubAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeStubAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeStubAbr, error) {

	result := &VrfOspfv3AreaTypeStubAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNssa) {
	o.NoSummary = util.YesNo(s.NoSummary, nil)
	if s.DefaultInformationOriginate != nil {
		var obj vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2
		obj.MarshalFromObject(*s.DefaultInformationOriginate)
		o.DefaultInformationOriginate = &obj
	}
	if s.Abr != nil {
		var obj vrfOspfv3AreaTypeNssaAbrXml_11_0_2
		obj.MarshalFromObject(*s.Abr)
		o.Abr = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNssa, error) {
	var defaultInformationOriginateVal *VrfOspfv3AreaTypeNssaDefaultInformationOriginate
	if o.DefaultInformationOriginate != nil {
		obj, err := o.DefaultInformationOriginate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultInformationOriginateVal = obj
	}
	var abrVal *VrfOspfv3AreaTypeNssaAbr
	if o.Abr != nil {
		obj, err := o.Abr.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		abrVal = obj
	}

	result := &VrfOspfv3AreaTypeNssa{
		NoSummary:                   util.AsBool(o.NoSummary, nil),
		DefaultInformationOriginate: defaultInformationOriginateVal,
		Abr:                         abrVal,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNssaDefaultInformationOriginate) {
	o.Metric = s.Metric
	o.MetricType = s.MetricType
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaDefaultInformationOriginateXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaDefaultInformationOriginate, error) {

	result := &VrfOspfv3AreaTypeNssaDefaultInformationOriginate{
		Metric:     o.Metric,
		MetricType: o.MetricType,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaAbrXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNssaAbr) {
	o.ImportList = s.ImportList
	o.ExportList = s.ExportList
	o.InboundFilterList = s.InboundFilterList
	o.OutboundFilterList = s.OutboundFilterList
	if s.NssaExtRange != nil {
		var objs []vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2
		for _, elt := range s.NssaExtRange {
			var obj vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &vrfOspfv3AreaTypeNssaAbrNssaExtRangeContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaAbrXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaAbr, error) {
	var nssaExtRangeVal []VrfOspfv3AreaTypeNssaAbrNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &VrfOspfv3AreaTypeNssaAbr{
		ImportList:         o.ImportList,
		ExportList:         o.ExportList,
		InboundFilterList:  o.InboundFilterList,
		OutboundFilterList: o.OutboundFilterList,
		NssaExtRange:       nssaExtRangeVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaTypeNssaAbrNssaExtRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaTypeNssaAbrNssaExtRangeXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaTypeNssaAbrNssaExtRange, error) {

	result := &VrfOspfv3AreaTypeNssaAbrNssaExtRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaRangeXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaRange) {
	o.Name = s.Name
	o.Advertise = util.YesNo(s.Advertise, nil)
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaRangeXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaRange, error) {

	result := &VrfOspfv3AreaRange{
		Name:      o.Name,
		Advertise: util.AsBool(o.Advertise, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.MtuIgnore = util.YesNo(s.MtuIgnore, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.Priority = s.Priority
	o.Metric = s.Metric
	o.InstanceId = s.InstanceId
	o.Authentication = s.Authentication
	o.Timing = s.Timing
	if s.LinkType != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeXml_11_0_2
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Bfd != nil {
		var obj vrfOspfv3AreaInterfaceBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterface, error) {
	var linkTypeVal *VrfOspfv3AreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var bfdVal *VrfOspfv3AreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &VrfOspfv3AreaInterface{
		Name:           o.Name,
		Enable:         util.AsBool(o.Enable, nil),
		MtuIgnore:      util.AsBool(o.MtuIgnore, nil),
		Passive:        util.AsBool(o.Passive, nil),
		Priority:       o.Priority,
		Metric:         o.Metric,
		InstanceId:     o.InstanceId,
		Authentication: o.Authentication,
		Timing:         o.Timing,
		LinkType:       linkTypeVal,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj vrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkType, error) {
	var broadcastVal *VrfOspfv3AreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *VrfOspfv3AreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *VrfOspfv3AreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &VrfOspfv3AreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeBroadcastXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeBroadcast, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2pXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2p, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2mp) {
	if s.Neighbor != nil {
		var objs []vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2mpXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2mp, error) {
	var neighborVal []VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &VrfOspfv3AreaInterfaceLinkTypeP2mp{
		Neighbor: neighborVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor) {
	o.Name = s.Name
	o.Priority = s.Priority
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceLinkTypeP2mpNeighborXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor, error) {

	result := &VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor{
		Name:     o.Name,
		Priority: o.Priority,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaInterfaceBfdXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaInterfaceBfdXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaInterfaceBfd, error) {

	result := &VrfOspfv3AreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfOspfv3AreaVirtualLinkXml_11_0_2) MarshalFromObject(s VrfOspfv3AreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.Timing = s.Timing
	o.Authentication = s.Authentication
	o.Misc = s.Misc
}

func (o vrfOspfv3AreaVirtualLinkXml_11_0_2) UnmarshalToObject() (*VrfOspfv3AreaVirtualLink, error) {

	result := &VrfOspfv3AreaVirtualLink{
		Name:           o.Name,
		NeighborId:     o.NeighborId,
		TransitAreaId:  o.TransitAreaId,
		Enable:         util.AsBool(o.Enable, nil),
		InstanceId:     o.InstanceId,
		Timing:         o.Timing,
		Authentication: o.Authentication,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpXml_11_0_2) MarshalFromObject(s VrfEcmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.MaxPath = s.MaxPath
	o.SymmetricReturn = util.YesNo(s.SymmetricReturn, nil)
	o.StrictSourcePath = util.YesNo(s.StrictSourcePath, nil)
	if s.Algorithm != nil {
		var obj vrfEcmpAlgorithmXml_11_0_2
		obj.MarshalFromObject(*s.Algorithm)
		o.Algorithm = &obj
	}
	o.Misc = s.Misc
}

func (o vrfEcmpXml_11_0_2) UnmarshalToObject() (*VrfEcmp, error) {
	var algorithmVal *VrfEcmpAlgorithm
	if o.Algorithm != nil {
		obj, err := o.Algorithm.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		algorithmVal = obj
	}

	result := &VrfEcmp{
		Enable:           util.AsBool(o.Enable, nil),
		MaxPath:          o.MaxPath,
		SymmetricReturn:  util.AsBool(o.SymmetricReturn, nil),
		StrictSourcePath: util.AsBool(o.StrictSourcePath, nil),
		Algorithm:        algorithmVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithm) {
	if s.IpModulo != nil {
		var obj vrfEcmpAlgorithmIpModuloXml_11_0_2
		obj.MarshalFromObject(*s.IpModulo)
		o.IpModulo = &obj
	}
	if s.IpHash != nil {
		var obj vrfEcmpAlgorithmIpHashXml_11_0_2
		obj.MarshalFromObject(*s.IpHash)
		o.IpHash = &obj
	}
	if s.WeightedRoundRobin != nil {
		var obj vrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2
		obj.MarshalFromObject(*s.WeightedRoundRobin)
		o.WeightedRoundRobin = &obj
	}
	if s.BalancedRoundRobin != nil {
		var obj vrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2
		obj.MarshalFromObject(*s.BalancedRoundRobin)
		o.BalancedRoundRobin = &obj
	}
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithm, error) {
	var ipModuloVal *VrfEcmpAlgorithmIpModulo
	if o.IpModulo != nil {
		obj, err := o.IpModulo.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipModuloVal = obj
	}
	var ipHashVal *VrfEcmpAlgorithmIpHash
	if o.IpHash != nil {
		obj, err := o.IpHash.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipHashVal = obj
	}
	var weightedRoundRobinVal *VrfEcmpAlgorithmWeightedRoundRobin
	if o.WeightedRoundRobin != nil {
		obj, err := o.WeightedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weightedRoundRobinVal = obj
	}
	var balancedRoundRobinVal *VrfEcmpAlgorithmBalancedRoundRobin
	if o.BalancedRoundRobin != nil {
		obj, err := o.BalancedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		balancedRoundRobinVal = obj
	}

	result := &VrfEcmpAlgorithm{
		IpModulo:           ipModuloVal,
		IpHash:             ipHashVal,
		WeightedRoundRobin: weightedRoundRobinVal,
		BalancedRoundRobin: balancedRoundRobinVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmIpModuloXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithmIpModulo) {
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmIpModuloXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithmIpModulo, error) {

	result := &VrfEcmpAlgorithmIpModulo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmIpHashXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithmIpHash) {
	o.SrcOnly = util.YesNo(s.SrcOnly, nil)
	o.UsePort = util.YesNo(s.UsePort, nil)
	o.HashSeed = s.HashSeed
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmIpHashXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithmIpHash, error) {

	result := &VrfEcmpAlgorithmIpHash{
		SrcOnly:  util.AsBool(o.SrcOnly, nil),
		UsePort:  util.AsBool(o.UsePort, nil),
		HashSeed: o.HashSeed,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithmWeightedRoundRobin) {
	if s.Interface != nil {
		var objs []vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfEcmpAlgorithmWeightedRoundRobinInterfaceContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmWeightedRoundRobinXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithmWeightedRoundRobin, error) {
	var interfaceVal []VrfEcmpAlgorithmWeightedRoundRobinInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfEcmpAlgorithmWeightedRoundRobin{
		Interface: interfaceVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithmWeightedRoundRobinInterface) {
	o.Name = s.Name
	o.Weight = s.Weight
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmWeightedRoundRobinInterfaceXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithmWeightedRoundRobinInterface, error) {

	result := &VrfEcmpAlgorithmWeightedRoundRobinInterface{
		Name:   o.Name,
		Weight: o.Weight,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *vrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2) MarshalFromObject(s VrfEcmpAlgorithmBalancedRoundRobin) {
	o.Misc = s.Misc
}

func (o vrfEcmpAlgorithmBalancedRoundRobinXml_11_0_2) UnmarshalToObject() (*VrfEcmpAlgorithmBalancedRoundRobin, error) {

	result := &VrfEcmpAlgorithmBalancedRoundRobin{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastXml_11_0_2) MarshalFromObject(s VrfMulticast) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.StaticRoute != nil {
		var objs []vrfMulticastStaticRouteXml_11_0_2
		for _, elt := range s.StaticRoute {
			var obj vrfMulticastStaticRouteXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticRoute = &vrfMulticastStaticRouteContainerXml_11_0_2{Entries: objs}
	}
	if s.Pim != nil {
		var obj vrfMulticastPimXml_11_0_2
		obj.MarshalFromObject(*s.Pim)
		o.Pim = &obj
	}
	if s.Igmp != nil {
		var obj vrfMulticastIgmpXml_11_0_2
		obj.MarshalFromObject(*s.Igmp)
		o.Igmp = &obj
	}
	if s.Msdp != nil {
		var obj vrfMulticastMsdpXml_11_0_2
		obj.MarshalFromObject(*s.Msdp)
		o.Msdp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastXml_11_0_2) UnmarshalToObject() (*VrfMulticast, error) {
	var staticRouteVal []VrfMulticastStaticRoute
	if o.StaticRoute != nil {
		for _, elt := range o.StaticRoute.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticRouteVal = append(staticRouteVal, *obj)
		}
	}
	var pimVal *VrfMulticastPim
	if o.Pim != nil {
		obj, err := o.Pim.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pimVal = obj
	}
	var igmpVal *VrfMulticastIgmp
	if o.Igmp != nil {
		obj, err := o.Igmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		igmpVal = obj
	}
	var msdpVal *VrfMulticastMsdp
	if o.Msdp != nil {
		obj, err := o.Msdp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		msdpVal = obj
	}

	result := &VrfMulticast{
		Enable:      util.AsBool(o.Enable, nil),
		StaticRoute: staticRouteVal,
		Pim:         pimVal,
		Igmp:        igmpVal,
		Msdp:        msdpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastStaticRouteXml_11_0_2) MarshalFromObject(s VrfMulticastStaticRoute) {
	o.Name = s.Name
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.Preference = s.Preference
	if s.Nexthop != nil {
		var obj vrfMulticastStaticRouteNexthopXml_11_0_2
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastStaticRouteXml_11_0_2) UnmarshalToObject() (*VrfMulticastStaticRoute, error) {
	var nexthopVal *VrfMulticastStaticRouteNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}

	result := &VrfMulticastStaticRoute{
		Name:        o.Name,
		Destination: o.Destination,
		Interface:   o.Interface,
		Preference:  o.Preference,
		Nexthop:     nexthopVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastStaticRouteNexthopXml_11_0_2) MarshalFromObject(s VrfMulticastStaticRouteNexthop) {
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
}

func (o vrfMulticastStaticRouteNexthopXml_11_0_2) UnmarshalToObject() (*VrfMulticastStaticRouteNexthop, error) {

	result := &VrfMulticastStaticRouteNexthop{
		IpAddress: o.IpAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimXml_11_0_2) MarshalFromObject(s VrfMulticastPim) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.RpfLookupMode = s.RpfLookupMode
	o.RouteAgeoutTime = s.RouteAgeoutTime
	o.IfTimerGlobal = s.IfTimerGlobal
	o.GroupPermission = s.GroupPermission
	if s.SsmAddressSpace != nil {
		var obj vrfMulticastPimSsmAddressSpaceXml_11_0_2
		obj.MarshalFromObject(*s.SsmAddressSpace)
		o.SsmAddressSpace = &obj
	}
	if s.Rp != nil {
		var obj vrfMulticastPimRpXml_11_0_2
		obj.MarshalFromObject(*s.Rp)
		o.Rp = &obj
	}
	if s.SptThreshold != nil {
		var objs []vrfMulticastPimSptThresholdXml_11_0_2
		for _, elt := range s.SptThreshold {
			var obj vrfMulticastPimSptThresholdXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SptThreshold = &vrfMulticastPimSptThresholdContainerXml_11_0_2{Entries: objs}
	}
	if s.Interface != nil {
		var objs []vrfMulticastPimInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfMulticastPimInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfMulticastPimInterfaceContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimXml_11_0_2) UnmarshalToObject() (*VrfMulticastPim, error) {
	var ssmAddressSpaceVal *VrfMulticastPimSsmAddressSpace
	if o.SsmAddressSpace != nil {
		obj, err := o.SsmAddressSpace.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ssmAddressSpaceVal = obj
	}
	var rpVal *VrfMulticastPimRp
	if o.Rp != nil {
		obj, err := o.Rp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		rpVal = obj
	}
	var sptThresholdVal []VrfMulticastPimSptThreshold
	if o.SptThreshold != nil {
		for _, elt := range o.SptThreshold.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			sptThresholdVal = append(sptThresholdVal, *obj)
		}
	}
	var interfaceVal []VrfMulticastPimInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfMulticastPim{
		Enable:          util.AsBool(o.Enable, nil),
		RpfLookupMode:   o.RpfLookupMode,
		RouteAgeoutTime: o.RouteAgeoutTime,
		IfTimerGlobal:   o.IfTimerGlobal,
		GroupPermission: o.GroupPermission,
		SsmAddressSpace: ssmAddressSpaceVal,
		Rp:              rpVal,
		SptThreshold:    sptThresholdVal,
		Interface:       interfaceVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimSsmAddressSpaceXml_11_0_2) MarshalFromObject(s VrfMulticastPimSsmAddressSpace) {
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimSsmAddressSpaceXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimSsmAddressSpace, error) {

	result := &VrfMulticastPimSsmAddressSpace{
		GroupList: o.GroupList,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpXml_11_0_2) MarshalFromObject(s VrfMulticastPimRp) {
	if s.LocalRp != nil {
		var obj vrfMulticastPimRpLocalRpXml_11_0_2
		obj.MarshalFromObject(*s.LocalRp)
		o.LocalRp = &obj
	}
	if s.ExternalRp != nil {
		var objs []vrfMulticastPimRpExternalRpXml_11_0_2
		for _, elt := range s.ExternalRp {
			var obj vrfMulticastPimRpExternalRpXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExternalRp = &vrfMulticastPimRpExternalRpContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimRp, error) {
	var localRpVal *VrfMulticastPimRpLocalRp
	if o.LocalRp != nil {
		obj, err := o.LocalRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localRpVal = obj
	}
	var externalRpVal []VrfMulticastPimRpExternalRp
	if o.ExternalRp != nil {
		for _, elt := range o.ExternalRp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			externalRpVal = append(externalRpVal, *obj)
		}
	}

	result := &VrfMulticastPimRp{
		LocalRp:    localRpVal,
		ExternalRp: externalRpVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpXml_11_0_2) MarshalFromObject(s VrfMulticastPimRpLocalRp) {
	if s.StaticRp != nil {
		var obj vrfMulticastPimRpLocalRpStaticRpXml_11_0_2
		obj.MarshalFromObject(*s.StaticRp)
		o.StaticRp = &obj
	}
	if s.CandidateRp != nil {
		var obj vrfMulticastPimRpLocalRpCandidateRpXml_11_0_2
		obj.MarshalFromObject(*s.CandidateRp)
		o.CandidateRp = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimRpLocalRp, error) {
	var staticRpVal *VrfMulticastPimRpLocalRpStaticRp
	if o.StaticRp != nil {
		obj, err := o.StaticRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticRpVal = obj
	}
	var candidateRpVal *VrfMulticastPimRpLocalRpCandidateRp
	if o.CandidateRp != nil {
		obj, err := o.CandidateRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		candidateRpVal = obj
	}

	result := &VrfMulticastPimRpLocalRp{
		StaticRp:    staticRpVal,
		CandidateRp: candidateRpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpStaticRpXml_11_0_2) MarshalFromObject(s VrfMulticastPimRpLocalRpStaticRp) {
	o.Interface = s.Interface
	o.Address = s.Address
	o.Override = util.YesNo(s.Override, nil)
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpStaticRpXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimRpLocalRpStaticRp, error) {

	result := &VrfMulticastPimRpLocalRpStaticRp{
		Interface: o.Interface,
		Address:   o.Address,
		Override:  util.AsBool(o.Override, nil),
		GroupList: o.GroupList,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpLocalRpCandidateRpXml_11_0_2) MarshalFromObject(s VrfMulticastPimRpLocalRpCandidateRp) {
	o.Interface = s.Interface
	o.Address = s.Address
	o.Priority = s.Priority
	o.AdvertisementInterval = s.AdvertisementInterval
	o.GroupList = s.GroupList
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpLocalRpCandidateRpXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimRpLocalRpCandidateRp, error) {

	result := &VrfMulticastPimRpLocalRpCandidateRp{
		Interface:             o.Interface,
		Address:               o.Address,
		Priority:              o.Priority,
		AdvertisementInterval: o.AdvertisementInterval,
		GroupList:             o.GroupList,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimRpExternalRpXml_11_0_2) MarshalFromObject(s VrfMulticastPimRpExternalRp) {
	o.Name = s.Name
	o.GroupList = s.GroupList
	o.Override = util.YesNo(s.Override, nil)
	o.Misc = s.Misc
}

func (o vrfMulticastPimRpExternalRpXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimRpExternalRp, error) {

	result := &VrfMulticastPimRpExternalRp{
		Name:      o.Name,
		GroupList: o.GroupList,
		Override:  util.AsBool(o.Override, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimSptThresholdXml_11_0_2) MarshalFromObject(s VrfMulticastPimSptThreshold) {
	o.Name = s.Name
	o.Threshold = s.Threshold
	o.Misc = s.Misc
}

func (o vrfMulticastPimSptThresholdXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimSptThreshold, error) {

	result := &VrfMulticastPimSptThreshold{
		Name:      o.Name,
		Threshold: o.Threshold,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastPimInterfaceXml_11_0_2) MarshalFromObject(s VrfMulticastPimInterface) {
	o.Name = s.Name
	o.Description = s.Description
	o.DrPriority = s.DrPriority
	o.SendBsm = util.YesNo(s.SendBsm, nil)
	o.IfTimer = s.IfTimer
	o.NeighborFilter = s.NeighborFilter
	o.Misc = s.Misc
}

func (o vrfMulticastPimInterfaceXml_11_0_2) UnmarshalToObject() (*VrfMulticastPimInterface, error) {

	result := &VrfMulticastPimInterface{
		Name:           o.Name,
		Description:    o.Description,
		DrPriority:     o.DrPriority,
		SendBsm:        util.AsBool(o.SendBsm, nil),
		IfTimer:        o.IfTimer,
		NeighborFilter: o.NeighborFilter,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpXml_11_0_2) MarshalFromObject(s VrfMulticastIgmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Dynamic != nil {
		var obj vrfMulticastIgmpDynamicXml_11_0_2
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.Static != nil {
		var objs []vrfMulticastIgmpStaticXml_11_0_2
		for _, elt := range s.Static {
			var obj vrfMulticastIgmpStaticXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Static = &vrfMulticastIgmpStaticContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpXml_11_0_2) UnmarshalToObject() (*VrfMulticastIgmp, error) {
	var dynamicVal *VrfMulticastIgmpDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var staticVal []VrfMulticastIgmpStatic
	if o.Static != nil {
		for _, elt := range o.Static.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticVal = append(staticVal, *obj)
		}
	}

	result := &VrfMulticastIgmp{
		Enable:  util.AsBool(o.Enable, nil),
		Dynamic: dynamicVal,
		Static:  staticVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpDynamicXml_11_0_2) MarshalFromObject(s VrfMulticastIgmpDynamic) {
	if s.Interface != nil {
		var objs []vrfMulticastIgmpDynamicInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfMulticastIgmpDynamicInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfMulticastIgmpDynamicInterfaceContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpDynamicXml_11_0_2) UnmarshalToObject() (*VrfMulticastIgmpDynamic, error) {
	var interfaceVal []VrfMulticastIgmpDynamicInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfMulticastIgmpDynamic{
		Interface: interfaceVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpDynamicInterfaceXml_11_0_2) MarshalFromObject(s VrfMulticastIgmpDynamicInterface) {
	o.Name = s.Name
	o.Version = s.Version
	o.Robustness = s.Robustness
	o.GroupFilter = s.GroupFilter
	o.MaxGroups = s.MaxGroups
	o.MaxSources = s.MaxSources
	o.QueryProfile = s.QueryProfile
	o.RouterAlertPolicing = util.YesNo(s.RouterAlertPolicing, nil)
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpDynamicInterfaceXml_11_0_2) UnmarshalToObject() (*VrfMulticastIgmpDynamicInterface, error) {

	result := &VrfMulticastIgmpDynamicInterface{
		Name:                o.Name,
		Version:             o.Version,
		Robustness:          o.Robustness,
		GroupFilter:         o.GroupFilter,
		MaxGroups:           o.MaxGroups,
		MaxSources:          o.MaxSources,
		QueryProfile:        o.QueryProfile,
		RouterAlertPolicing: util.AsBool(o.RouterAlertPolicing, nil),
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastIgmpStaticXml_11_0_2) MarshalFromObject(s VrfMulticastIgmpStatic) {
	o.Name = s.Name
	o.Interface = s.Interface
	o.GroupAddress = s.GroupAddress
	o.SourceAddress = s.SourceAddress
	o.Misc = s.Misc
}

func (o vrfMulticastIgmpStaticXml_11_0_2) UnmarshalToObject() (*VrfMulticastIgmpStatic, error) {

	result := &VrfMulticastIgmpStatic{
		Name:          o.Name,
		Interface:     o.Interface,
		GroupAddress:  o.GroupAddress,
		SourceAddress: o.SourceAddress,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpXml_11_0_2) MarshalFromObject(s VrfMulticastMsdp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GlobalTimer = s.GlobalTimer
	o.GlobalAuthentication = s.GlobalAuthentication
	if s.OriginatorId != nil {
		var obj vrfMulticastMsdpOriginatorIdXml_11_0_2
		obj.MarshalFromObject(*s.OriginatorId)
		o.OriginatorId = &obj
	}
	if s.Peer != nil {
		var objs []vrfMulticastMsdpPeerXml_11_0_2
		for _, elt := range s.Peer {
			var obj vrfMulticastMsdpPeerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Peer = &vrfMulticastMsdpPeerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpXml_11_0_2) UnmarshalToObject() (*VrfMulticastMsdp, error) {
	var originatorIdVal *VrfMulticastMsdpOriginatorId
	if o.OriginatorId != nil {
		obj, err := o.OriginatorId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		originatorIdVal = obj
	}
	var peerVal []VrfMulticastMsdpPeer
	if o.Peer != nil {
		for _, elt := range o.Peer.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerVal = append(peerVal, *obj)
		}
	}

	result := &VrfMulticastMsdp{
		Enable:               util.AsBool(o.Enable, nil),
		GlobalTimer:          o.GlobalTimer,
		GlobalAuthentication: o.GlobalAuthentication,
		OriginatorId:         originatorIdVal,
		Peer:                 peerVal,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpOriginatorIdXml_11_0_2) MarshalFromObject(s VrfMulticastMsdpOriginatorId) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpOriginatorIdXml_11_0_2) UnmarshalToObject() (*VrfMulticastMsdpOriginatorId, error) {

	result := &VrfMulticastMsdpOriginatorId{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerXml_11_0_2) MarshalFromObject(s VrfMulticastMsdpPeer) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.PeerAs = s.PeerAs
	o.Authentication = s.Authentication
	o.MaxSa = s.MaxSa
	o.InboundSaFilter = s.InboundSaFilter
	o.OutboundSaFilter = s.OutboundSaFilter
	if s.LocalAddress != nil {
		var obj vrfMulticastMsdpPeerLocalAddressXml_11_0_2
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj vrfMulticastMsdpPeerPeerAddressXml_11_0_2
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerXml_11_0_2) UnmarshalToObject() (*VrfMulticastMsdpPeer, error) {
	var localAddressVal *VrfMulticastMsdpPeerLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *VrfMulticastMsdpPeerPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}

	result := &VrfMulticastMsdpPeer{
		Name:             o.Name,
		Enable:           util.AsBool(o.Enable, nil),
		PeerAs:           o.PeerAs,
		Authentication:   o.Authentication,
		MaxSa:            o.MaxSa,
		InboundSaFilter:  o.InboundSaFilter,
		OutboundSaFilter: o.OutboundSaFilter,
		LocalAddress:     localAddressVal,
		PeerAddress:      peerAddressVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerLocalAddressXml_11_0_2) MarshalFromObject(s VrfMulticastMsdpPeerLocalAddress) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerLocalAddressXml_11_0_2) UnmarshalToObject() (*VrfMulticastMsdpPeerLocalAddress, error) {

	result := &VrfMulticastMsdpPeerLocalAddress{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *vrfMulticastMsdpPeerPeerAddressXml_11_0_2) MarshalFromObject(s VrfMulticastMsdpPeerPeerAddress) {
	o.Ip = s.Ip
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o vrfMulticastMsdpPeerPeerAddressXml_11_0_2) UnmarshalToObject() (*VrfMulticastMsdpPeerPeerAddress, error) {

	result := &VrfMulticastMsdpPeerPeerAddress{
		Ip:   o.Ip,
		Fqdn: o.Fqdn,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *vrfRipXml_11_0_2) MarshalFromObject(s VrfRip) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.DefaultInformationOriginate = util.YesNo(s.DefaultInformationOriginate, nil)
	o.GlobalTimer = s.GlobalTimer
	o.AuthProfile = s.AuthProfile
	o.RedistributionProfile = s.RedistributionProfile
	if s.GlobalBfd != nil {
		var obj vrfRipGlobalBfdXml_11_0_2
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GlobalInboundDistributeList != nil {
		var obj vrfRipGlobalInboundDistributeListXml_11_0_2
		obj.MarshalFromObject(*s.GlobalInboundDistributeList)
		o.GlobalInboundDistributeList = &obj
	}
	if s.GlobalOutboundDistributeList != nil {
		var obj vrfRipGlobalOutboundDistributeListXml_11_0_2
		obj.MarshalFromObject(*s.GlobalOutboundDistributeList)
		o.GlobalOutboundDistributeList = &obj
	}
	if s.Interface != nil {
		var objs []vrfRipInterfaceXml_11_0_2
		for _, elt := range s.Interface {
			var obj vrfRipInterfaceXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &vrfRipInterfaceContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o vrfRipXml_11_0_2) UnmarshalToObject() (*VrfRip, error) {
	var globalBfdVal *VrfRipGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var globalInboundDistributeListVal *VrfRipGlobalInboundDistributeList
	if o.GlobalInboundDistributeList != nil {
		obj, err := o.GlobalInboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalInboundDistributeListVal = obj
	}
	var globalOutboundDistributeListVal *VrfRipGlobalOutboundDistributeList
	if o.GlobalOutboundDistributeList != nil {
		obj, err := o.GlobalOutboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalOutboundDistributeListVal = obj
	}
	var interfaceVal []VrfRipInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &VrfRip{
		Enable:                       util.AsBool(o.Enable, nil),
		DefaultInformationOriginate:  util.AsBool(o.DefaultInformationOriginate, nil),
		GlobalTimer:                  o.GlobalTimer,
		AuthProfile:                  o.AuthProfile,
		RedistributionProfile:        o.RedistributionProfile,
		GlobalBfd:                    globalBfdVal,
		GlobalInboundDistributeList:  globalInboundDistributeListVal,
		GlobalOutboundDistributeList: globalOutboundDistributeListVal,
		Interface:                    interfaceVal,
		Misc:                         o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalBfdXml_11_0_2) MarshalFromObject(s VrfRipGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRipGlobalBfdXml_11_0_2) UnmarshalToObject() (*VrfRipGlobalBfd, error) {

	result := &VrfRipGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalInboundDistributeListXml_11_0_2) MarshalFromObject(s VrfRipGlobalInboundDistributeList) {
	o.AccessList = s.AccessList
	o.Misc = s.Misc
}

func (o vrfRipGlobalInboundDistributeListXml_11_0_2) UnmarshalToObject() (*VrfRipGlobalInboundDistributeList, error) {

	result := &VrfRipGlobalInboundDistributeList{
		AccessList: o.AccessList,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipGlobalOutboundDistributeListXml_11_0_2) MarshalFromObject(s VrfRipGlobalOutboundDistributeList) {
	o.AccessList = s.AccessList
	o.Misc = s.Misc
}

func (o vrfRipGlobalOutboundDistributeListXml_11_0_2) UnmarshalToObject() (*VrfRipGlobalOutboundDistributeList, error) {

	result := &VrfRipGlobalOutboundDistributeList{
		AccessList: o.AccessList,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceXml_11_0_2) MarshalFromObject(s VrfRipInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Mode = s.Mode
	o.SplitHorizon = s.SplitHorizon
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj vrfRipInterfaceBfdXml_11_0_2
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	if s.InterfaceInboundDistributeList != nil {
		var obj vrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2
		obj.MarshalFromObject(*s.InterfaceInboundDistributeList)
		o.InterfaceInboundDistributeList = &obj
	}
	if s.InterfaceOutboundDistributeList != nil {
		var obj vrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2
		obj.MarshalFromObject(*s.InterfaceOutboundDistributeList)
		o.InterfaceOutboundDistributeList = &obj
	}
	o.Misc = s.Misc
}

func (o vrfRipInterfaceXml_11_0_2) UnmarshalToObject() (*VrfRipInterface, error) {
	var bfdVal *VrfRipInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var interfaceInboundDistributeListVal *VrfRipInterfaceInterfaceInboundDistributeList
	if o.InterfaceInboundDistributeList != nil {
		obj, err := o.InterfaceInboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceInboundDistributeListVal = obj
	}
	var interfaceOutboundDistributeListVal *VrfRipInterfaceInterfaceOutboundDistributeList
	if o.InterfaceOutboundDistributeList != nil {
		obj, err := o.InterfaceOutboundDistributeList.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceOutboundDistributeListVal = obj
	}

	result := &VrfRipInterface{
		Name:                            o.Name,
		Enable:                          util.AsBool(o.Enable, nil),
		Mode:                            o.Mode,
		SplitHorizon:                    o.SplitHorizon,
		Authentication:                  o.Authentication,
		Bfd:                             bfdVal,
		InterfaceInboundDistributeList:  interfaceInboundDistributeListVal,
		InterfaceOutboundDistributeList: interfaceOutboundDistributeListVal,
		Misc:                            o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceBfdXml_11_0_2) MarshalFromObject(s VrfRipInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o vrfRipInterfaceBfdXml_11_0_2) UnmarshalToObject() (*VrfRipInterfaceBfd, error) {

	result := &VrfRipInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2) MarshalFromObject(s VrfRipInterfaceInterfaceInboundDistributeList) {
	o.AccessList = s.AccessList
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o vrfRipInterfaceInterfaceInboundDistributeListXml_11_0_2) UnmarshalToObject() (*VrfRipInterfaceInterfaceInboundDistributeList, error) {

	result := &VrfRipInterfaceInterfaceInboundDistributeList{
		AccessList: o.AccessList,
		Metric:     o.Metric,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *vrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2) MarshalFromObject(s VrfRipInterfaceInterfaceOutboundDistributeList) {
	o.AccessList = s.AccessList
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o vrfRipInterfaceInterfaceOutboundDistributeListXml_11_0_2) UnmarshalToObject() (*VrfRipInterfaceInterfaceOutboundDistributeList, error) {

	result := &VrfRipInterfaceInterfaceOutboundDistributeList{
		AccessList: o.AccessList,
		Metric:     o.Metric,
		Misc:       o.Misc,
	}
	return result, nil
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
	if len(o.Vrf) != len(other.Vrf) {
		return false
	}
	for idx := range o.Vrf {
		if !o.Vrf[idx].matches(&other.Vrf[idx]) {
			return false
		}
	}

	return true
}

func (o *Vrf) matches(other *Vrf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !o.AdminDists.matches(other.AdminDists) {
		return false
	}
	if !o.RibFilter.matches(other.RibFilter) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.RoutingTable.matches(other.RoutingTable) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Ecmp.matches(other.Ecmp) {
		return false
	}
	if !o.Multicast.matches(other.Multicast) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *VrfAdminDists) matches(other *VrfAdminDists) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Static, other.Static) {
		return false
	}
	if !util.Ints64Match(o.StaticIpv6, other.StaticIpv6) {
		return false
	}
	if !util.Ints64Match(o.OspfInter, other.OspfInter) {
		return false
	}
	if !util.Ints64Match(o.OspfIntra, other.OspfIntra) {
		return false
	}
	if !util.Ints64Match(o.OspfExt, other.OspfExt) {
		return false
	}
	if !util.Ints64Match(o.Ospfv3Inter, other.Ospfv3Inter) {
		return false
	}
	if !util.Ints64Match(o.Ospfv3Intra, other.Ospfv3Intra) {
		return false
	}
	if !util.Ints64Match(o.Ospfv3Ext, other.Ospfv3Ext) {
		return false
	}
	if !util.Ints64Match(o.BgpInternal, other.BgpInternal) {
		return false
	}
	if !util.Ints64Match(o.BgpExternal, other.BgpExternal) {
		return false
	}
	if !util.Ints64Match(o.BgpLocal, other.BgpLocal) {
		return false
	}
	if !util.Ints64Match(o.Rip, other.Rip) {
		return false
	}

	return true
}

func (o *VrfRibFilter) matches(other *VrfRibFilter) bool {
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

func (o *VrfRibFilterIpv4) matches(other *VrfRibFilterIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Static.matches(other.Static) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv4Static) matches(other *VrfRibFilterIpv4Static) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv4Bgp) matches(other *VrfRibFilterIpv4Bgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv4Ospf) matches(other *VrfRibFilterIpv4Ospf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv4Rip) matches(other *VrfRibFilterIpv4Rip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv6) matches(other *VrfRibFilterIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Static.matches(other.Static) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv6Static) matches(other *VrfRibFilterIpv6Static) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv6Bgp) matches(other *VrfRibFilterIpv6Bgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfRibFilterIpv6Ospfv3) matches(other *VrfRibFilterIpv6Ospfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteMap, other.RouteMap) {
		return false
	}

	return true
}

func (o *VrfBgp) matches(other *VrfBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !util.StringsMatch(o.LocalAs, other.LocalAs) {
		return false
	}
	if !util.BoolsMatch(o.InstallRoute, other.InstallRoute) {
		return false
	}
	if !util.BoolsMatch(o.EnforceFirstAs, other.EnforceFirstAs) {
		return false
	}
	if !util.BoolsMatch(o.FastExternalFailover, other.FastExternalFailover) {
		return false
	}
	if !util.BoolsMatch(o.EcmpMultiAs, other.EcmpMultiAs) {
		return false
	}
	if !util.Ints64Match(o.DefaultLocalPreference, other.DefaultLocalPreference) {
		return false
	}
	if !util.BoolsMatch(o.GracefulShutdown, other.GracefulShutdown) {
		return false
	}
	if !util.BoolsMatch(o.AlwaysAdvertiseNetworkRoute, other.AlwaysAdvertiseNetworkRoute) {
		return false
	}
	if !o.Med.matches(other.Med) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.RedistributionProfile.matches(other.RedistributionProfile) {
		return false
	}
	if !o.AdvertiseNetwork.matches(other.AdvertiseNetwork) {
		return false
	}
	if len(o.PeerGroup) != len(other.PeerGroup) {
		return false
	}
	for idx := range o.PeerGroup {
		if !o.PeerGroup[idx].matches(&other.PeerGroup[idx]) {
			return false
		}
	}
	if len(o.AggregateRoutes) != len(other.AggregateRoutes) {
		return false
	}
	for idx := range o.AggregateRoutes {
		if !o.AggregateRoutes[idx].matches(&other.AggregateRoutes[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfBgpMed) matches(other *VrfBgpMed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AlwaysCompareMed, other.AlwaysCompareMed) {
		return false
	}
	if !util.BoolsMatch(o.DeterministicMedComparison, other.DeterministicMedComparison) {
		return false
	}

	return true
}

func (o *VrfBgpGracefulRestart) matches(other *VrfBgpGracefulRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.StaleRouteTime, other.StaleRouteTime) {
		return false
	}
	if !util.Ints64Match(o.MaxPeerRestartTime, other.MaxPeerRestartTime) {
		return false
	}
	if !util.Ints64Match(o.LocalRestartTime, other.LocalRestartTime) {
		return false
	}

	return true
}

func (o *VrfBgpGlobalBfd) matches(other *VrfBgpGlobalBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfBgpRedistributionProfile) matches(other *VrfBgpRedistributionProfile) bool {
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

func (o *VrfBgpRedistributionProfileIpv4) matches(other *VrfBgpRedistributionProfileIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Unicast, other.Unicast) {
		return false
	}

	return true
}

func (o *VrfBgpRedistributionProfileIpv6) matches(other *VrfBgpRedistributionProfileIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Unicast, other.Unicast) {
		return false
	}

	return true
}

func (o *VrfBgpAdvertiseNetwork) matches(other *VrfBgpAdvertiseNetwork) bool {
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

func (o *VrfBgpAdvertiseNetworkIpv4) matches(other *VrfBgpAdvertiseNetworkIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Network) != len(other.Network) {
		return false
	}
	for idx := range o.Network {
		if !o.Network[idx].matches(&other.Network[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfBgpAdvertiseNetworkIpv4Network) matches(other *VrfBgpAdvertiseNetworkIpv4Network) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Unicast, other.Unicast) {
		return false
	}
	if !util.BoolsMatch(o.Multicast, other.Multicast) {
		return false
	}
	if !util.BoolsMatch(o.Backdoor, other.Backdoor) {
		return false
	}

	return true
}

func (o *VrfBgpAdvertiseNetworkIpv6) matches(other *VrfBgpAdvertiseNetworkIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Network) != len(other.Network) {
		return false
	}
	for idx := range o.Network {
		if !o.Network[idx].matches(&other.Network[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfBgpAdvertiseNetworkIpv6Network) matches(other *VrfBgpAdvertiseNetworkIpv6Network) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Unicast, other.Unicast) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroup) matches(other *VrfBgpPeerGroup) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}
	if !o.AddressFamily.matches(other.AddressFamily) {
		return false
	}
	if !o.FilteringProfile.matches(other.FilteringProfile) {
		return false
	}
	if !o.ConnectionOptions.matches(other.ConnectionOptions) {
		return false
	}
	if len(o.Peer) != len(other.Peer) {
		return false
	}
	for idx := range o.Peer {
		if !o.Peer[idx].matches(&other.Peer[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfBgpPeerGroupType) matches(other *VrfBgpPeerGroupType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ibgp.matches(other.Ibgp) {
		return false
	}
	if !o.Ebgp.matches(other.Ebgp) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupTypeIbgp) matches(other *VrfBgpPeerGroupTypeIbgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupTypeEbgp) matches(other *VrfBgpPeerGroupTypeEbgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupAddressFamily) matches(other *VrfBgpPeerGroupAddressFamily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupFilteringProfile) matches(other *VrfBgpPeerGroupFilteringProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupConnectionOptions) matches(other *VrfBgpPeerGroupConnectionOptions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Timers, other.Timers) {
		return false
	}
	if !util.Ints64Match(o.Multihop, other.Multihop) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Dampening, other.Dampening) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeer) matches(other *VrfBgpPeerGroupPeer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.Passive, other.Passive) {
		return false
	}
	if !util.StringsMatch(o.PeerAs, other.PeerAs) {
		return false
	}
	if !util.BoolsMatch(o.EnableSenderSideLoopDetection, other.EnableSenderSideLoopDetection) {
		return false
	}
	if !o.Inherit.matches(other.Inherit) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !o.PeerAddress.matches(other.PeerAddress) {
		return false
	}
	if !o.ConnectionOptions.matches(other.ConnectionOptions) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerInherit) matches(other *VrfBgpPeerGroupPeerInherit) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Yes.matches(other.Yes) {
		return false
	}
	if !o.No.matches(other.No) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerInheritYes) matches(other *VrfBgpPeerGroupPeerInheritYes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerInheritNo) matches(other *VrfBgpPeerGroupPeerInheritNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AddressFamily.matches(other.AddressFamily) {
		return false
	}
	if !o.FilteringProfile.matches(other.FilteringProfile) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerInheritNoAddressFamily) matches(other *VrfBgpPeerGroupPeerInheritNoAddressFamily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerInheritNoFilteringProfile) matches(other *VrfBgpPeerGroupPeerInheritNoFilteringProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerLocalAddress) matches(other *VrfBgpPeerGroupPeerLocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerPeerAddress) matches(other *VrfBgpPeerGroupPeerPeerAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerConnectionOptions) matches(other *VrfBgpPeerGroupPeerConnectionOptions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Timers, other.Timers) {
		return false
	}
	if !util.StringsMatch(o.Multihop, other.Multihop) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Dampening, other.Dampening) {
		return false
	}

	return true
}

func (o *VrfBgpPeerGroupPeerBfd) matches(other *VrfBgpPeerGroupPeerBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfBgpAggregateRoutes) matches(other *VrfBgpAggregateRoutes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.SummaryOnly, other.SummaryOnly) {
		return false
	}
	if !util.BoolsMatch(o.AsSet, other.AsSet) {
		return false
	}
	if !util.BoolsMatch(o.SameMed, other.SameMed) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}

	return true
}

func (o *VrfBgpAggregateRoutesType) matches(other *VrfBgpAggregateRoutesType) bool {
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

func (o *VrfBgpAggregateRoutesTypeIpv4) matches(other *VrfBgpAggregateRoutesTypeIpv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SummaryPrefix, other.SummaryPrefix) {
		return false
	}
	if !util.StringsMatch(o.SuppressMap, other.SuppressMap) {
		return false
	}
	if !util.StringsMatch(o.AttributeMap, other.AttributeMap) {
		return false
	}

	return true
}

func (o *VrfBgpAggregateRoutesTypeIpv6) matches(other *VrfBgpAggregateRoutesTypeIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SummaryPrefix, other.SummaryPrefix) {
		return false
	}
	if !util.StringsMatch(o.SuppressMap, other.SuppressMap) {
		return false
	}
	if !util.StringsMatch(o.AttributeMap, other.AttributeMap) {
		return false
	}

	return true
}

func (o *VrfRoutingTable) matches(other *VrfRoutingTable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIp) matches(other *VrfRoutingTableIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.StaticRoute) != len(other.StaticRoute) {
		return false
	}
	for idx := range o.StaticRoute {
		if !o.StaticRoute[idx].matches(&other.StaticRoute[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfRoutingTableIpStaticRoute) matches(other *VrfRoutingTableIpStaticRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.AdminDist, other.AdminDist) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Nexthop.matches(other.Nexthop) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}
	if !o.PathMonitor.matches(other.PathMonitor) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpStaticRouteNexthop) matches(other *VrfRoutingTableIpStaticRouteNexthop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Discard.matches(other.Discard) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}
	if !util.StringsMatch(o.NextLr, other.NextLr) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpStaticRouteNexthopDiscard) matches(other *VrfRoutingTableIpStaticRouteNexthopDiscard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpStaticRouteBfd) matches(other *VrfRoutingTableIpStaticRouteBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpStaticRoutePathMonitor) matches(other *VrfRoutingTableIpStaticRoutePathMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.FailureCondition, other.FailureCondition) {
		return false
	}
	if !util.Ints64Match(o.HoldTime, other.HoldTime) {
		return false
	}
	if len(o.MonitorDestinations) != len(other.MonitorDestinations) {
		return false
	}
	for idx := range o.MonitorDestinations {
		if !o.MonitorDestinations[idx].matches(&other.MonitorDestinations[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations) matches(other *VrfRoutingTableIpStaticRoutePathMonitorMonitorDestinations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Source, other.Source) {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Count, other.Count) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpv6) matches(other *VrfRoutingTableIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.StaticRoute) != len(other.StaticRoute) {
		return false
	}
	for idx := range o.StaticRoute {
		if !o.StaticRoute[idx].matches(&other.StaticRoute[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRoute) matches(other *VrfRoutingTableIpv6StaticRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.AdminDist, other.AdminDist) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Nexthop.matches(other.Nexthop) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}
	if !o.PathMonitor.matches(other.PathMonitor) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRouteNexthop) matches(other *VrfRoutingTableIpv6StaticRouteNexthop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Discard.matches(other.Discard) {
		return false
	}
	if !util.StringsMatch(o.Ipv6Address, other.Ipv6Address) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !util.StringsMatch(o.NextLr, other.NextLr) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRouteNexthopDiscard) matches(other *VrfRoutingTableIpv6StaticRouteNexthopDiscard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRouteBfd) matches(other *VrfRoutingTableIpv6StaticRouteBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRoutePathMonitor) matches(other *VrfRoutingTableIpv6StaticRoutePathMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.FailureCondition, other.FailureCondition) {
		return false
	}
	if !util.Ints64Match(o.HoldTime, other.HoldTime) {
		return false
	}
	if len(o.MonitorDestinations) != len(other.MonitorDestinations) {
		return false
	}
	for idx := range o.MonitorDestinations {
		if !o.MonitorDestinations[idx].matches(&other.MonitorDestinations[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) matches(other *VrfRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Source, other.Source) {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Count, other.Count) {
		return false
	}

	return true
}

func (o *VrfOspf) matches(other *VrfOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.Rfc1583, other.Rfc1583) {
		return false
	}
	if !util.StringsMatch(o.SpfTimer, other.SpfTimer) {
		return false
	}
	if !util.StringsMatch(o.GlobalIfTimer, other.GlobalIfTimer) {
		return false
	}
	if !util.StringsMatch(o.RedistributionProfile, other.RedistributionProfile) {
		return false
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if len(o.Area) != len(other.Area) {
		return false
	}
	for idx := range o.Area {
		if !o.Area[idx].matches(&other.Area[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfGlobalBfd) matches(other *VrfOspfGlobalBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfOspfGracefulRestart) matches(other *VrfOspfGracefulRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.GracePeriod, other.GracePeriod) {
		return false
	}
	if !util.BoolsMatch(o.HelperEnable, other.HelperEnable) {
		return false
	}
	if !util.BoolsMatch(o.StrictLSAChecking, other.StrictLSAChecking) {
		return false
	}
	if !util.Ints64Match(o.MaxNeighborRestartTime, other.MaxNeighborRestartTime) {
		return false
	}

	return true
}

func (o *VrfOspfArea) matches(other *VrfOspfArea) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}
	if len(o.Range) != len(other.Range) {
		return false
	}
	for idx := range o.Range {
		if !o.Range[idx].matches(&other.Range[idx]) {
			return false
		}
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}
	if len(o.VirtualLink) != len(other.VirtualLink) {
		return false
	}
	for idx := range o.VirtualLink {
		if !o.VirtualLink[idx].matches(&other.VirtualLink[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfAreaType) matches(other *VrfOspfAreaType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Normal.matches(other.Normal) {
		return false
	}
	if !o.Stub.matches(other.Stub) {
		return false
	}
	if !o.Nssa.matches(other.Nssa) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeNormal) matches(other *VrfOspfAreaTypeNormal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeNormalAbr) matches(other *VrfOspfAreaTypeNormalAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeStub) matches(other *VrfOspfAreaTypeStub) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.NoSummary, other.NoSummary) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}
	if !util.Ints64Match(o.DefaultRouteMetric, other.DefaultRouteMetric) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeStubAbr) matches(other *VrfOspfAreaTypeStubAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeNssa) matches(other *VrfOspfAreaTypeNssa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.NoSummary, other.NoSummary) {
		return false
	}
	if !o.DefaultInformationOriginate.matches(other.DefaultInformationOriginate) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}

	return true
}

func (o *VrfOspfAreaTypeNssaDefaultInformationOriginate) matches(other *VrfOspfAreaTypeNssaDefaultInformationOriginate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *VrfOspfAreaTypeNssaAbr) matches(other *VrfOspfAreaTypeNssaAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}
	if len(o.NssaExtRange) != len(other.NssaExtRange) {
		return false
	}
	for idx := range o.NssaExtRange {
		if !o.NssaExtRange[idx].matches(&other.NssaExtRange[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfAreaTypeNssaAbrNssaExtRange) matches(other *VrfOspfAreaTypeNssaAbrNssaExtRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Advertise, other.Advertise) {
		return false
	}

	return true
}

func (o *VrfOspfAreaRange) matches(other *VrfOspfAreaRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Advertise, other.Advertise) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterface) matches(other *VrfOspfAreaInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.MtuIgnore, other.MtuIgnore) {
		return false
	}
	if !util.BoolsMatch(o.Passive, other.Passive) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Timing, other.Timing) {
		return false
	}
	if !o.LinkType.matches(other.LinkType) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterfaceLinkType) matches(other *VrfOspfAreaInterfaceLinkType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Broadcast.matches(other.Broadcast) {
		return false
	}
	if !o.P2p.matches(other.P2p) {
		return false
	}
	if !o.P2mp.matches(other.P2mp) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterfaceLinkTypeBroadcast) matches(other *VrfOspfAreaInterfaceLinkTypeBroadcast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterfaceLinkTypeP2p) matches(other *VrfOspfAreaInterfaceLinkTypeP2p) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterfaceLinkTypeP2mp) matches(other *VrfOspfAreaInterfaceLinkTypeP2mp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Neighbor) != len(other.Neighbor) {
		return false
	}
	for idx := range o.Neighbor {
		if !o.Neighbor[idx].matches(&other.Neighbor[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfAreaInterfaceLinkTypeP2mpNeighbor) matches(other *VrfOspfAreaInterfaceLinkTypeP2mpNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}

	return true
}

func (o *VrfOspfAreaInterfaceBfd) matches(other *VrfOspfAreaInterfaceBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfOspfAreaVirtualLink) matches(other *VrfOspfAreaVirtualLink) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.NeighborId, other.NeighborId) {
		return false
	}
	if !util.StringsMatch(o.TransitAreaId, other.TransitAreaId) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.InstanceId, other.InstanceId) {
		return false
	}
	if !util.StringsMatch(o.Timing, other.Timing) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *VrfOspfAreaVirtualLinkBfd) matches(other *VrfOspfAreaVirtualLinkBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfOspfv3) matches(other *VrfOspfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !util.BoolsMatch(o.DisableTransitTraffic, other.DisableTransitTraffic) {
		return false
	}
	if !util.StringsMatch(o.SpfTimer, other.SpfTimer) {
		return false
	}
	if !util.StringsMatch(o.GlobalIfTimer, other.GlobalIfTimer) {
		return false
	}
	if !util.StringsMatch(o.RedistributionProfile, other.RedistributionProfile) {
		return false
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if len(o.Area) != len(other.Area) {
		return false
	}
	for idx := range o.Area {
		if !o.Area[idx].matches(&other.Area[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfv3GlobalBfd) matches(other *VrfOspfv3GlobalBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfOspfv3GracefulRestart) matches(other *VrfOspfv3GracefulRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.GracePeriod, other.GracePeriod) {
		return false
	}
	if !util.BoolsMatch(o.HelperEnable, other.HelperEnable) {
		return false
	}
	if !util.BoolsMatch(o.StrictLSAChecking, other.StrictLSAChecking) {
		return false
	}
	if !util.Ints64Match(o.MaxNeighborRestartTime, other.MaxNeighborRestartTime) {
		return false
	}

	return true
}

func (o *VrfOspfv3Area) matches(other *VrfOspfv3Area) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}
	if len(o.Range) != len(other.Range) {
		return false
	}
	for idx := range o.Range {
		if !o.Range[idx].matches(&other.Range[idx]) {
			return false
		}
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}
	if len(o.VirtualLink) != len(other.VirtualLink) {
		return false
	}
	for idx := range o.VirtualLink {
		if !o.VirtualLink[idx].matches(&other.VirtualLink[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfv3AreaType) matches(other *VrfOspfv3AreaType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Normal.matches(other.Normal) {
		return false
	}
	if !o.Stub.matches(other.Stub) {
		return false
	}
	if !o.Nssa.matches(other.Nssa) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeNormal) matches(other *VrfOspfv3AreaTypeNormal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeNormalAbr) matches(other *VrfOspfv3AreaTypeNormalAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeStub) matches(other *VrfOspfv3AreaTypeStub) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.NoSummary, other.NoSummary) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}
	if !util.Ints64Match(o.DefaultRouteMetric, other.DefaultRouteMetric) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeStubAbr) matches(other *VrfOspfv3AreaTypeStubAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeNssa) matches(other *VrfOspfv3AreaTypeNssa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.NoSummary, other.NoSummary) {
		return false
	}
	if !o.DefaultInformationOriginate.matches(other.DefaultInformationOriginate) {
		return false
	}
	if !o.Abr.matches(other.Abr) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaTypeNssaDefaultInformationOriginate) matches(other *VrfOspfv3AreaTypeNssaDefaultInformationOriginate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *VrfOspfv3AreaTypeNssaAbr) matches(other *VrfOspfv3AreaTypeNssaAbr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportList, other.ImportList) {
		return false
	}
	if !util.StringsMatch(o.ExportList, other.ExportList) {
		return false
	}
	if !util.StringsMatch(o.InboundFilterList, other.InboundFilterList) {
		return false
	}
	if !util.StringsMatch(o.OutboundFilterList, other.OutboundFilterList) {
		return false
	}
	if len(o.NssaExtRange) != len(other.NssaExtRange) {
		return false
	}
	for idx := range o.NssaExtRange {
		if !o.NssaExtRange[idx].matches(&other.NssaExtRange[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfv3AreaTypeNssaAbrNssaExtRange) matches(other *VrfOspfv3AreaTypeNssaAbrNssaExtRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Advertise, other.Advertise) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaRange) matches(other *VrfOspfv3AreaRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Advertise, other.Advertise) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterface) matches(other *VrfOspfv3AreaInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.MtuIgnore, other.MtuIgnore) {
		return false
	}
	if !util.BoolsMatch(o.Passive, other.Passive) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.InstanceId, other.InstanceId) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Timing, other.Timing) {
		return false
	}
	if !o.LinkType.matches(other.LinkType) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceLinkType) matches(other *VrfOspfv3AreaInterfaceLinkType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Broadcast.matches(other.Broadcast) {
		return false
	}
	if !o.P2p.matches(other.P2p) {
		return false
	}
	if !o.P2mp.matches(other.P2mp) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceLinkTypeBroadcast) matches(other *VrfOspfv3AreaInterfaceLinkTypeBroadcast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceLinkTypeP2p) matches(other *VrfOspfv3AreaInterfaceLinkTypeP2p) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceLinkTypeP2mp) matches(other *VrfOspfv3AreaInterfaceLinkTypeP2mp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Neighbor) != len(other.Neighbor) {
		return false
	}
	for idx := range o.Neighbor {
		if !o.Neighbor[idx].matches(&other.Neighbor[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor) matches(other *VrfOspfv3AreaInterfaceLinkTypeP2mpNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaInterfaceBfd) matches(other *VrfOspfv3AreaInterfaceBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfOspfv3AreaVirtualLink) matches(other *VrfOspfv3AreaVirtualLink) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.NeighborId, other.NeighborId) {
		return false
	}
	if !util.StringsMatch(o.TransitAreaId, other.TransitAreaId) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.InstanceId, other.InstanceId) {
		return false
	}
	if !util.StringsMatch(o.Timing, other.Timing) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}

	return true
}

func (o *VrfEcmp) matches(other *VrfEcmp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.MaxPath, other.MaxPath) {
		return false
	}
	if !util.BoolsMatch(o.SymmetricReturn, other.SymmetricReturn) {
		return false
	}
	if !util.BoolsMatch(o.StrictSourcePath, other.StrictSourcePath) {
		return false
	}
	if !o.Algorithm.matches(other.Algorithm) {
		return false
	}

	return true
}

func (o *VrfEcmpAlgorithm) matches(other *VrfEcmpAlgorithm) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.IpModulo.matches(other.IpModulo) {
		return false
	}
	if !o.IpHash.matches(other.IpHash) {
		return false
	}
	if !o.WeightedRoundRobin.matches(other.WeightedRoundRobin) {
		return false
	}
	if !o.BalancedRoundRobin.matches(other.BalancedRoundRobin) {
		return false
	}

	return true
}

func (o *VrfEcmpAlgorithmIpModulo) matches(other *VrfEcmpAlgorithmIpModulo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfEcmpAlgorithmIpHash) matches(other *VrfEcmpAlgorithmIpHash) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.SrcOnly, other.SrcOnly) {
		return false
	}
	if !util.BoolsMatch(o.UsePort, other.UsePort) {
		return false
	}
	if !util.Ints64Match(o.HashSeed, other.HashSeed) {
		return false
	}

	return true
}

func (o *VrfEcmpAlgorithmWeightedRoundRobin) matches(other *VrfEcmpAlgorithmWeightedRoundRobin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfEcmpAlgorithmWeightedRoundRobinInterface) matches(other *VrfEcmpAlgorithmWeightedRoundRobinInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Weight, other.Weight) {
		return false
	}

	return true
}

func (o *VrfEcmpAlgorithmBalancedRoundRobin) matches(other *VrfEcmpAlgorithmBalancedRoundRobin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *VrfMulticast) matches(other *VrfMulticast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.StaticRoute) != len(other.StaticRoute) {
		return false
	}
	for idx := range o.StaticRoute {
		if !o.StaticRoute[idx].matches(&other.StaticRoute[idx]) {
			return false
		}
	}
	if !o.Pim.matches(other.Pim) {
		return false
	}
	if !o.Igmp.matches(other.Igmp) {
		return false
	}
	if !o.Msdp.matches(other.Msdp) {
		return false
	}

	return true
}

func (o *VrfMulticastStaticRoute) matches(other *VrfMulticastStaticRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Preference, other.Preference) {
		return false
	}
	if !o.Nexthop.matches(other.Nexthop) {
		return false
	}

	return true
}

func (o *VrfMulticastStaticRouteNexthop) matches(other *VrfMulticastStaticRouteNexthop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}

	return true
}

func (o *VrfMulticastPim) matches(other *VrfMulticastPim) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.RpfLookupMode, other.RpfLookupMode) {
		return false
	}
	if !util.Ints64Match(o.RouteAgeoutTime, other.RouteAgeoutTime) {
		return false
	}
	if !util.StringsMatch(o.IfTimerGlobal, other.IfTimerGlobal) {
		return false
	}
	if !util.StringsMatch(o.GroupPermission, other.GroupPermission) {
		return false
	}
	if !o.SsmAddressSpace.matches(other.SsmAddressSpace) {
		return false
	}
	if !o.Rp.matches(other.Rp) {
		return false
	}
	if len(o.SptThreshold) != len(other.SptThreshold) {
		return false
	}
	for idx := range o.SptThreshold {
		if !o.SptThreshold[idx].matches(&other.SptThreshold[idx]) {
			return false
		}
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfMulticastPimSsmAddressSpace) matches(other *VrfMulticastPimSsmAddressSpace) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.GroupList, other.GroupList) {
		return false
	}

	return true
}

func (o *VrfMulticastPimRp) matches(other *VrfMulticastPimRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.LocalRp.matches(other.LocalRp) {
		return false
	}
	if len(o.ExternalRp) != len(other.ExternalRp) {
		return false
	}
	for idx := range o.ExternalRp {
		if !o.ExternalRp[idx].matches(&other.ExternalRp[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfMulticastPimRpLocalRp) matches(other *VrfMulticastPimRpLocalRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.StaticRp.matches(other.StaticRp) {
		return false
	}
	if !o.CandidateRp.matches(other.CandidateRp) {
		return false
	}

	return true
}

func (o *VrfMulticastPimRpLocalRpStaticRp) matches(other *VrfMulticastPimRpLocalRpStaticRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.BoolsMatch(o.Override, other.Override) {
		return false
	}
	if !util.StringsMatch(o.GroupList, other.GroupList) {
		return false
	}

	return true
}

func (o *VrfMulticastPimRpLocalRpCandidateRp) matches(other *VrfMulticastPimRpLocalRpCandidateRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.Ints64Match(o.AdvertisementInterval, other.AdvertisementInterval) {
		return false
	}
	if !util.StringsMatch(o.GroupList, other.GroupList) {
		return false
	}

	return true
}

func (o *VrfMulticastPimRpExternalRp) matches(other *VrfMulticastPimRpExternalRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.GroupList, other.GroupList) {
		return false
	}
	if !util.BoolsMatch(o.Override, other.Override) {
		return false
	}

	return true
}

func (o *VrfMulticastPimSptThreshold) matches(other *VrfMulticastPimSptThreshold) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Threshold, other.Threshold) {
		return false
	}

	return true
}

func (o *VrfMulticastPimInterface) matches(other *VrfMulticastPimInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.Ints64Match(o.DrPriority, other.DrPriority) {
		return false
	}
	if !util.BoolsMatch(o.SendBsm, other.SendBsm) {
		return false
	}
	if !util.StringsMatch(o.IfTimer, other.IfTimer) {
		return false
	}
	if !util.StringsMatch(o.NeighborFilter, other.NeighborFilter) {
		return false
	}

	return true
}

func (o *VrfMulticastIgmp) matches(other *VrfMulticastIgmp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Dynamic.matches(other.Dynamic) {
		return false
	}
	if len(o.Static) != len(other.Static) {
		return false
	}
	for idx := range o.Static {
		if !o.Static[idx].matches(&other.Static[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfMulticastIgmpDynamic) matches(other *VrfMulticastIgmpDynamic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfMulticastIgmpDynamicInterface) matches(other *VrfMulticastIgmpDynamicInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Version, other.Version) {
		return false
	}
	if !util.StringsMatch(o.Robustness, other.Robustness) {
		return false
	}
	if !util.StringsMatch(o.GroupFilter, other.GroupFilter) {
		return false
	}
	if !util.StringsMatch(o.MaxGroups, other.MaxGroups) {
		return false
	}
	if !util.StringsMatch(o.MaxSources, other.MaxSources) {
		return false
	}
	if !util.StringsMatch(o.QueryProfile, other.QueryProfile) {
		return false
	}
	if !util.BoolsMatch(o.RouterAlertPolicing, other.RouterAlertPolicing) {
		return false
	}

	return true
}

func (o *VrfMulticastIgmpStatic) matches(other *VrfMulticastIgmpStatic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.GroupAddress, other.GroupAddress) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}

	return true
}

func (o *VrfMulticastMsdp) matches(other *VrfMulticastMsdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.GlobalTimer, other.GlobalTimer) {
		return false
	}
	if !util.StringsMatch(o.GlobalAuthentication, other.GlobalAuthentication) {
		return false
	}
	if !o.OriginatorId.matches(other.OriginatorId) {
		return false
	}
	if len(o.Peer) != len(other.Peer) {
		return false
	}
	for idx := range o.Peer {
		if !o.Peer[idx].matches(&other.Peer[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfMulticastMsdpOriginatorId) matches(other *VrfMulticastMsdpOriginatorId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *VrfMulticastMsdpPeer) matches(other *VrfMulticastMsdpPeer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.PeerAs, other.PeerAs) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.Ints64Match(o.MaxSa, other.MaxSa) {
		return false
	}
	if !util.StringsMatch(o.InboundSaFilter, other.InboundSaFilter) {
		return false
	}
	if !util.StringsMatch(o.OutboundSaFilter, other.OutboundSaFilter) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !o.PeerAddress.matches(other.PeerAddress) {
		return false
	}

	return true
}

func (o *VrfMulticastMsdpPeerLocalAddress) matches(other *VrfMulticastMsdpPeerLocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *VrfMulticastMsdpPeerPeerAddress) matches(other *VrfMulticastMsdpPeerPeerAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}

	return true
}

func (o *VrfRip) matches(other *VrfRip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.DefaultInformationOriginate, other.DefaultInformationOriginate) {
		return false
	}
	if !util.StringsMatch(o.GlobalTimer, other.GlobalTimer) {
		return false
	}
	if !util.StringsMatch(o.AuthProfile, other.AuthProfile) {
		return false
	}
	if !util.StringsMatch(o.RedistributionProfile, other.RedistributionProfile) {
		return false
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.GlobalInboundDistributeList.matches(other.GlobalInboundDistributeList) {
		return false
	}
	if !o.GlobalOutboundDistributeList.matches(other.GlobalOutboundDistributeList) {
		return false
	}
	if len(o.Interface) != len(other.Interface) {
		return false
	}
	for idx := range o.Interface {
		if !o.Interface[idx].matches(&other.Interface[idx]) {
			return false
		}
	}

	return true
}

func (o *VrfRipGlobalBfd) matches(other *VrfRipGlobalBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfRipGlobalInboundDistributeList) matches(other *VrfRipGlobalInboundDistributeList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}

	return true
}

func (o *VrfRipGlobalOutboundDistributeList) matches(other *VrfRipGlobalOutboundDistributeList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}

	return true
}

func (o *VrfRipInterface) matches(other *VrfRipInterface) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !util.StringsMatch(o.SplitHorizon, other.SplitHorizon) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}
	if !o.InterfaceInboundDistributeList.matches(other.InterfaceInboundDistributeList) {
		return false
	}
	if !o.InterfaceOutboundDistributeList.matches(other.InterfaceOutboundDistributeList) {
		return false
	}

	return true
}

func (o *VrfRipInterfaceBfd) matches(other *VrfRipInterfaceBfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VrfRipInterfaceInterfaceInboundDistributeList) matches(other *VrfRipInterfaceInterfaceInboundDistributeList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *VrfRipInterfaceInterfaceOutboundDistributeList) matches(other *VrfRipInterfaceInterfaceOutboundDistributeList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessList, other.AccessList) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
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
