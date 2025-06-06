package virtual_router

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
	suffix = []string{"network", "virtual-router", "$name"}
)

type Entry struct {
	Name       string
	AdminDists *AdminDists
	Ecmp       *Ecmp
	Interface  []string
	Multicast  *Multicast
	Protocol   *Protocol
	Misc       []generic.Xml
}
type AdminDists struct {
	Ebgp       *int64
	Ibgp       *int64
	OspfExt    *int64
	OspfInt    *int64
	Ospfv3Ext  *int64
	Ospfv3Int  *int64
	Rip        *int64
	Static     *int64
	StaticIpv6 *int64
	Misc       []generic.Xml
}
type Ecmp struct {
	Algorithm        *EcmpAlgorithm
	Enable           *bool
	MaxPath          *int64
	StrictSourcePath *bool
	SymmetricReturn  *bool
	Misc             []generic.Xml
}
type EcmpAlgorithm struct {
	BalancedRoundRobin *EcmpAlgorithmBalancedRoundRobin
	IpHash             *EcmpAlgorithmIpHash
	IpModulo           *EcmpAlgorithmIpModulo
	WeightedRoundRobin *EcmpAlgorithmWeightedRoundRobin
	Misc               []generic.Xml
}
type EcmpAlgorithmBalancedRoundRobin struct {
	Misc []generic.Xml
}
type EcmpAlgorithmIpHash struct {
	HashSeed *int64
	SrcOnly  *bool
	UsePort  *bool
	Misc     []generic.Xml
}
type EcmpAlgorithmIpModulo struct {
	Misc []generic.Xml
}
type EcmpAlgorithmWeightedRoundRobin struct {
	Interface []EcmpAlgorithmWeightedRoundRobinInterface
	Misc      []generic.Xml
}
type EcmpAlgorithmWeightedRoundRobinInterface struct {
	Name   string
	Weight *int64
	Misc   []generic.Xml
}
type Multicast struct {
	Enable          *bool
	InterfaceGroup  []MulticastInterfaceGroup
	RouteAgeoutTime *int64
	Rp              *MulticastRp
	SptThreshold    []MulticastSptThreshold
	SsmAddressSpace []MulticastSsmAddressSpace
	Misc            []generic.Xml
}
type MulticastInterfaceGroup struct {
	Name            string
	Description     *string
	Interface       []string
	GroupPermission *MulticastInterfaceGroupGroupPermission
	Igmp            *MulticastInterfaceGroupIgmp
	Pim             *MulticastInterfaceGroupPim
	Misc            []generic.Xml
}
type MulticastInterfaceGroupGroupPermission struct {
	AnySourceMulticast      []MulticastInterfaceGroupGroupPermissionAnySourceMulticast
	SourceSpecificMulticast []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast
	Misc                    []generic.Xml
}
type MulticastInterfaceGroupGroupPermissionAnySourceMulticast struct {
	Name         string
	GroupAddress *string
	Included     *bool
	Misc         []generic.Xml
}
type MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast struct {
	Name          string
	GroupAddress  *string
	SourceAddress *string
	Included      *bool
	Misc          []generic.Xml
}
type MulticastInterfaceGroupIgmp struct {
	Enable                  *bool
	Version                 *string
	MaxQueryResponseTime    *float64
	QueryInterval           *int64
	LastMemberQueryInterval *float64
	ImmediateLeave          *bool
	Robustness              *string
	MaxGroups               *string
	MaxSources              *string
	RouterAlertPolicing     *bool
	Misc                    []generic.Xml
}
type MulticastInterfaceGroupPim struct {
	Enable            *bool
	AssertInterval    *int64
	HelloInterval     *int64
	JoinPruneInterval *int64
	DrPriority        *int64
	BsrBorder         *bool
	AllowedNeighbors  []MulticastInterfaceGroupPimAllowedNeighbors
	Misc              []generic.Xml
}
type MulticastInterfaceGroupPimAllowedNeighbors struct {
	Name string
	Misc []generic.Xml
}
type MulticastRp struct {
	ExternalRp []MulticastRpExternalRp
	LocalRp    *MulticastRpLocalRp
	Misc       []generic.Xml
}
type MulticastRpExternalRp struct {
	Name           string
	GroupAddresses []string
	Override       *bool
	Misc           []generic.Xml
}
type MulticastRpLocalRp struct {
	CandidateRp *MulticastRpLocalRpCandidateRp
	StaticRp    *MulticastRpLocalRpStaticRp
	Misc        []generic.Xml
}
type MulticastRpLocalRpCandidateRp struct {
	Address               *string
	AdvertisementInterval *int64
	GroupAddresses        []string
	Interface             *string
	Priority              *int64
	Misc                  []generic.Xml
}
type MulticastRpLocalRpStaticRp struct {
	Address        *string
	GroupAddresses []string
	Interface      *string
	Override       *bool
	Misc           []generic.Xml
}
type MulticastSptThreshold struct {
	Name      string
	Threshold *string
	Misc      []generic.Xml
}
type MulticastSsmAddressSpace struct {
	Name         string
	GroupAddress *string
	Included     *bool
	Misc         []generic.Xml
}
type Protocol struct {
	Bgp               *ProtocolBgp
	Ospf              *ProtocolOspf
	Ospfv3            *ProtocolOspfv3
	RedistProfile     []ProtocolRedistProfile
	RedistProfileIpv6 []ProtocolRedistProfileIpv6
	Rip               *ProtocolRip
	Misc              []generic.Xml
}
type ProtocolBgp struct {
	AllowRedistDefaultRoute *bool
	AuthProfile             []ProtocolBgpAuthProfile
	DampeningProfile        []ProtocolBgpDampeningProfile
	EcmpMultiAs             *bool
	Enable                  *bool
	EnforceFirstAs          *bool
	GlobalBfd               *ProtocolBgpGlobalBfd
	InstallRoute            *bool
	LocalAs                 *string
	PeerGroup               []ProtocolBgpPeerGroup
	Policy                  *ProtocolBgpPolicy
	RedistRules             []ProtocolBgpRedistRules
	RejectDefaultRoute      *bool
	RouterId                *string
	RoutingOptions          *ProtocolBgpRoutingOptions
	Misc                    []generic.Xml
}
type ProtocolBgpAuthProfile struct {
	Name   string
	Secret *string
	Misc   []generic.Xml
}
type ProtocolBgpDampeningProfile struct {
	Name                     string
	Enable                   *bool
	Cutoff                   *float64
	Reuse                    *float64
	MaxHoldTime              *int64
	DecayHalfLifeReachable   *int64
	DecayHalfLifeUnreachable *int64
	Misc                     []generic.Xml
}
type ProtocolBgpGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolBgpPeerGroup struct {
	Name                    string
	Enable                  *bool
	AggregatedConfedAsPath  *bool
	SoftResetWithStoredInfo *bool
	Type                    *ProtocolBgpPeerGroupType
	Peer                    []ProtocolBgpPeerGroupPeer
	Misc                    []generic.Xml
}
type ProtocolBgpPeerGroupType struct {
	Ibgp       *ProtocolBgpPeerGroupTypeIbgp
	EbgpConfed *ProtocolBgpPeerGroupTypeEbgpConfed
	IbgpConfed *ProtocolBgpPeerGroupTypeIbgpConfed
	Ebgp       *ProtocolBgpPeerGroupTypeEbgp
	Misc       []generic.Xml
}
type ProtocolBgpPeerGroupTypeIbgp struct {
	ExportNexthop *string
	Misc          []generic.Xml
}
type ProtocolBgpPeerGroupTypeEbgpConfed struct {
	ExportNexthop *string
	Misc          []generic.Xml
}
type ProtocolBgpPeerGroupTypeIbgpConfed struct {
	ExportNexthop *string
	Misc          []generic.Xml
}
type ProtocolBgpPeerGroupTypeEbgp struct {
	ImportNexthop   *string
	ExportNexthop   *string
	RemovePrivateAs *bool
	Misc            []generic.Xml
}
type ProtocolBgpPeerGroupPeer struct {
	Name                              string
	Enable                            *bool
	PeerAs                            *string
	EnableMpBgp                       *bool
	AddressFamilyIdentifier           *string
	EnableSenderSideLoopDetection     *bool
	ReflectorClient                   *string
	PeeringType                       *string
	MaxPrefixes                       *string
	SubsequentAddressFamilyIdentifier *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier
	LocalAddress                      *ProtocolBgpPeerGroupPeerLocalAddress
	PeerAddress                       *ProtocolBgpPeerGroupPeerPeerAddress
	ConnectionOptions                 *ProtocolBgpPeerGroupPeerConnectionOptions
	Bfd                               *ProtocolBgpPeerGroupPeerBfd
	Misc                              []generic.Xml
}
type ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier struct {
	Unicast   *bool
	Multicast *bool
	Misc      []generic.Xml
}
type ProtocolBgpPeerGroupPeerLocalAddress struct {
	Interface *string
	Ip        *string
	Misc      []generic.Xml
}
type ProtocolBgpPeerGroupPeerPeerAddress struct {
	Ip   *string
	Fqdn *string
	Misc []generic.Xml
}
type ProtocolBgpPeerGroupPeerConnectionOptions struct {
	Authentication        *string
	KeepAliveInterval     *string
	MinRouteAdvInterval   *int64
	Multihop              *int64
	OpenDelayTime         *int64
	HoldTime              *string
	IdleHoldTime          *int64
	IncomingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection
	OutgoingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection
	Misc                  []generic.Xml
}
type ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection struct {
	RemotePort *int64
	Allow      *bool
	Misc       []generic.Xml
}
type ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection struct {
	LocalPort *int64
	Allow     *bool
	Misc      []generic.Xml
}
type ProtocolBgpPeerGroupPeerBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolBgpPolicy struct {
	Aggregation              *ProtocolBgpPolicyAggregation
	ConditionalAdvertisement *ProtocolBgpPolicyConditionalAdvertisement
	Export                   *ProtocolBgpPolicyExport
	Import                   *ProtocolBgpPolicyImport
	Misc                     []generic.Xml
}
type ProtocolBgpPolicyAggregation struct {
	Address []ProtocolBgpPolicyAggregationAddress
	Misc    []generic.Xml
}
type ProtocolBgpPolicyAggregationAddress struct {
	Name                     string
	Prefix                   *string
	Enable                   *bool
	Summary                  *bool
	AsSet                    *bool
	AggregateRouteAttributes *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes
	SuppressFilters          []ProtocolBgpPolicyAggregationAddressSuppressFilters
	AdvertiseFilters         []ProtocolBgpPolicyAggregationAddressAdvertiseFilters
	Misc                     []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes struct {
	LocalPreference   *int64
	Med               *int64
	Weight            *int64
	Nexthop           *string
	Origin            *string
	AsPathLimit       *int64
	AsPath            *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath
	Community         *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath struct {
	None    *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone
	Prepend *int64
	Misc    []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity struct {
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity struct {
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFilters struct {
	Name   string
	Enable *bool
	Match  *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch
	Misc   []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath
	Community         *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix struct {
	Name  string
	Exact *bool
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFilters struct {
	Name   string
	Enable *bool
	Match  *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch
	Misc   []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath
	Community         *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix struct {
	Name  string
	Exact *bool
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisement struct {
	Policy []ProtocolBgpPolicyConditionalAdvertisementPolicy
	Misc   []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicy struct {
	Name             string
	Enable           *bool
	UsedBy           []string
	NonExistFilters  []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters
	AdvertiseFilters []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters
	Misc             []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters struct {
	Name   string
	Enable *bool
	Match  *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch
	Misc   []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix struct {
	Name string
	Misc []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters struct {
	Name   string
	Enable *bool
	Match  *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch
	Misc   []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix struct {
	Name string
	Misc []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExport struct {
	Rules []ProtocolBgpPolicyExportRules
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRules struct {
	Name   string
	Enable *bool
	UsedBy []string
	Match  *ProtocolBgpPolicyExportRulesMatch
	Action *ProtocolBgpPolicyExportRulesAction
	Misc   []generic.Xml
}
type ProtocolBgpPolicyExportRulesMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyExportRulesMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyExportRulesMatchAsPath
	Community         *ProtocolBgpPolicyExportRulesMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyExportRulesMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyExportRulesMatchAddressPrefix struct {
	Name  string
	Exact *bool
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRulesMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRulesMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRulesMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRulesAction struct {
	Deny  *ProtocolBgpPolicyExportRulesActionDeny
	Allow *ProtocolBgpPolicyExportRulesActionAllow
	Misc  []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionDeny struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllow struct {
	Update *ProtocolBgpPolicyExportRulesActionAllowUpdate
	Misc   []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdate struct {
	LocalPreference   *int64
	Med               *int64
	Nexthop           *string
	Origin            *string
	AsPathLimit       *int64
	AsPath            *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath
	Community         *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity
	ExtendedCommunity *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath struct {
	None             *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone
	Remove           *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove
	Prepend          *int64
	RemoveAndPrepend *int64
	Misc             []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity struct {
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity struct {
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImport struct {
	Rules []ProtocolBgpPolicyImportRules
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRules struct {
	Name   string
	Enable *bool
	UsedBy []string
	Match  *ProtocolBgpPolicyImportRulesMatch
	Action *ProtocolBgpPolicyImportRulesAction
	Misc   []generic.Xml
}
type ProtocolBgpPolicyImportRulesMatch struct {
	RouteTable        *string
	Med               *int64
	AddressPrefix     []ProtocolBgpPolicyImportRulesMatchAddressPrefix
	Nexthop           []string
	FromPeer          []string
	AsPath            *ProtocolBgpPolicyImportRulesMatchAsPath
	Community         *ProtocolBgpPolicyImportRulesMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyImportRulesMatchExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyImportRulesMatchAddressPrefix struct {
	Name  string
	Exact *bool
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRulesMatchAsPath struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRulesMatchCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRulesMatchExtendedCommunity struct {
	Regex *string
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRulesAction struct {
	Deny  *ProtocolBgpPolicyImportRulesActionDeny
	Allow *ProtocolBgpPolicyImportRulesActionAllow
	Misc  []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionDeny struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllow struct {
	Dampening *string
	Update    *ProtocolBgpPolicyImportRulesActionAllowUpdate
	Misc      []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdate struct {
	LocalPreference   *int64
	Med               *int64
	Weight            *int64
	Nexthop           *string
	Origin            *string
	AsPathLimit       *int64
	AsPath            *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath
	Community         *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity
	ExtendedCommunity *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity
	Misc              []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath struct {
	None   *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone
	Remove *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove
	Misc   []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity struct {
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity struct {
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll
	RemoveRegex *string
	Append      []string
	Overwrite   []string
	Misc        []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone struct {
	Misc []generic.Xml
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll struct {
	Misc []generic.Xml
}
type ProtocolBgpRedistRules struct {
	Name                    string
	AddressFamilyIdentifier *string
	RouteTable              *string
	Enable                  *bool
	SetOrigin               *string
	SetMed                  *int64
	SetLocalPreference      *int64
	SetAsPathLimit          *int64
	Metric                  *int64
	SetCommunity            []string
	SetExtendedCommunity    []string
	Misc                    []generic.Xml
}
type ProtocolBgpRoutingOptions struct {
	Aggregate              *ProtocolBgpRoutingOptionsAggregate
	AsFormat               *string
	ConfederationMemberAs  *string
	DefaultLocalPreference *int64
	GracefulRestart        *ProtocolBgpRoutingOptionsGracefulRestart
	Med                    *ProtocolBgpRoutingOptionsMed
	ReflectorClusterId     *string
	Misc                   []generic.Xml
}
type ProtocolBgpRoutingOptionsAggregate struct {
	AggregateMed *bool
	Misc         []generic.Xml
}
type ProtocolBgpRoutingOptionsGracefulRestart struct {
	Enable             *bool
	LocalRestartTime   *int64
	MaxPeerRestartTime *int64
	StaleRouteTime     *int64
	Misc               []generic.Xml
}
type ProtocolBgpRoutingOptionsMed struct {
	AlwaysCompareMed           *bool
	DeterministicMedComparison *bool
	Misc                       []generic.Xml
}
type ProtocolOspf struct {
	AllowRedistDefaultRoute *bool
	Area                    []ProtocolOspfArea
	AuthProfile             []ProtocolOspfAuthProfile
	Enable                  *bool
	ExportRules             []ProtocolOspfExportRules
	GlobalBfd               *ProtocolOspfGlobalBfd
	GracefulRestart         *ProtocolOspfGracefulRestart
	RejectDefaultRoute      *bool
	Rfc1583                 *bool
	RouterId                *string
	Timers                  *ProtocolOspfTimers
	Misc                    []generic.Xml
}
type ProtocolOspfArea struct {
	Name        string
	Type        *ProtocolOspfAreaType
	Range       []ProtocolOspfAreaRange
	Interface   []ProtocolOspfAreaInterface
	VirtualLink []ProtocolOspfAreaVirtualLink
	Misc        []generic.Xml
}
type ProtocolOspfAreaType struct {
	Normal *ProtocolOspfAreaTypeNormal
	Stub   *ProtocolOspfAreaTypeStub
	Nssa   *ProtocolOspfAreaTypeNssa
	Misc   []generic.Xml
}
type ProtocolOspfAreaTypeNormal struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaTypeStub struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfAreaTypeStubDefaultRoute
	Misc          []generic.Xml
}
type ProtocolOspfAreaTypeStubDefaultRoute struct {
	Disable   *ProtocolOspfAreaTypeStubDefaultRouteDisable
	Advertise *ProtocolOspfAreaTypeStubDefaultRouteAdvertise
	Misc      []generic.Xml
}
type ProtocolOspfAreaTypeStubDefaultRouteDisable struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaTypeStubDefaultRouteAdvertise struct {
	Metric *int64
	Misc   []generic.Xml
}
type ProtocolOspfAreaTypeNssa struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfAreaTypeNssaDefaultRoute
	NssaExtRange  []ProtocolOspfAreaTypeNssaNssaExtRange
	Misc          []generic.Xml
}
type ProtocolOspfAreaTypeNssaDefaultRoute struct {
	Disable   *ProtocolOspfAreaTypeNssaDefaultRouteDisable
	Advertise *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise
	Misc      []generic.Xml
}
type ProtocolOspfAreaTypeNssaDefaultRouteDisable struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaTypeNssaDefaultRouteAdvertise struct {
	Metric *int64
	Type   *string
	Misc   []generic.Xml
}
type ProtocolOspfAreaTypeNssaNssaExtRange struct {
	Name      string
	Advertise *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise
	Suppress  *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress
	Misc      []generic.Xml
}
type ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaTypeNssaNssaExtRangeSuppress struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaRange struct {
	Name      string
	Advertise *ProtocolOspfAreaRangeAdvertise
	Suppress  *ProtocolOspfAreaRangeSuppress
	Misc      []generic.Xml
}
type ProtocolOspfAreaRangeAdvertise struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaRangeSuppress struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaInterface struct {
	Name               string
	Enable             *bool
	Passive            *bool
	Metric             *int64
	Priority           *int64
	HelloInterval      *int64
	DeadCounts         *int64
	RetransmitInterval *int64
	TransitDelay       *int64
	Authentication     *string
	GrDelay            *int64
	LinkType           *ProtocolOspfAreaInterfaceLinkType
	Neighbor           []ProtocolOspfAreaInterfaceNeighbor
	Bfd                *ProtocolOspfAreaInterfaceBfd
	Misc               []generic.Xml
}
type ProtocolOspfAreaInterfaceLinkType struct {
	Broadcast *ProtocolOspfAreaInterfaceLinkTypeBroadcast
	P2p       *ProtocolOspfAreaInterfaceLinkTypeP2p
	P2mp      *ProtocolOspfAreaInterfaceLinkTypeP2mp
	Misc      []generic.Xml
}
type ProtocolOspfAreaInterfaceLinkTypeBroadcast struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaInterfaceLinkTypeP2p struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaInterfaceLinkTypeP2mp struct {
	Misc []generic.Xml
}
type ProtocolOspfAreaInterfaceNeighbor struct {
	Name string
	Misc []generic.Xml
}
type ProtocolOspfAreaInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfAreaVirtualLink struct {
	Name               string
	NeighborId         *string
	TransitAreaId      *string
	Enable             *bool
	HelloInterval      *int64
	DeadCounts         *int64
	RetransmitInterval *int64
	TransitDelay       *int64
	Authentication     *string
	Bfd                *ProtocolOspfAreaVirtualLinkBfd
	Misc               []generic.Xml
}
type ProtocolOspfAreaVirtualLinkBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfAuthProfile struct {
	Name     string
	Password *string
	Md5      []ProtocolOspfAuthProfileMd5
	Misc     []generic.Xml
}
type ProtocolOspfAuthProfileMd5 struct {
	Name      string
	Key       *string
	Preferred *bool
	Misc      []generic.Xml
}
type ProtocolOspfExportRules struct {
	Name        string
	NewPathType *string
	NewTag      *string
	Metric      *int64
	Misc        []generic.Xml
}
type ProtocolOspfGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfGracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	MaxNeighborRestartTime *int64
	StrictLSAChecking      *bool
	Misc                   []generic.Xml
}
type ProtocolOspfTimers struct {
	LsaInterval         *float64
	SpfCalculationDelay *float64
	Misc                []generic.Xml
}
type ProtocolOspfv3 struct {
	AllowRedistDefaultRoute *bool
	Area                    []ProtocolOspfv3Area
	AuthProfile             []ProtocolOspfv3AuthProfile
	DisableTransitTraffic   *bool
	Enable                  *bool
	ExportRules             []ProtocolOspfv3ExportRules
	GlobalBfd               *ProtocolOspfv3GlobalBfd
	GracefulRestart         *ProtocolOspfv3GracefulRestart
	RejectDefaultRoute      *bool
	RouterId                *string
	Timers                  *ProtocolOspfv3Timers
	Misc                    []generic.Xml
}
type ProtocolOspfv3Area struct {
	Name           string
	Authentication *string
	Type           *ProtocolOspfv3AreaType
	Range          []ProtocolOspfv3AreaRange
	Interface      []ProtocolOspfv3AreaInterface
	VirtualLink    []ProtocolOspfv3AreaVirtualLink
	Misc           []generic.Xml
}
type ProtocolOspfv3AreaType struct {
	Normal *ProtocolOspfv3AreaTypeNormal
	Stub   *ProtocolOspfv3AreaTypeStub
	Nssa   *ProtocolOspfv3AreaTypeNssa
	Misc   []generic.Xml
}
type ProtocolOspfv3AreaTypeNormal struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaTypeStub struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfv3AreaTypeStubDefaultRoute
	Misc          []generic.Xml
}
type ProtocolOspfv3AreaTypeStubDefaultRoute struct {
	Disable   *ProtocolOspfv3AreaTypeStubDefaultRouteDisable
	Advertise *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise
	Misc      []generic.Xml
}
type ProtocolOspfv3AreaTypeStubDefaultRouteDisable struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise struct {
	Metric *int64
	Misc   []generic.Xml
}
type ProtocolOspfv3AreaTypeNssa struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfv3AreaTypeNssaDefaultRoute
	NssaExtRange  []ProtocolOspfv3AreaTypeNssaNssaExtRange
	Misc          []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaDefaultRoute struct {
	Disable   *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable
	Advertise *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise
	Misc      []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteDisable struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise struct {
	Metric *int64
	Type   *string
	Misc   []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaNssaExtRange struct {
	Name      string
	Advertise *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise
	Suppress  *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress
	Misc      []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaRange struct {
	Name      string
	Advertise *ProtocolOspfv3AreaRangeAdvertise
	Suppress  *ProtocolOspfv3AreaRangeSuppress
	Misc      []generic.Xml
}
type ProtocolOspfv3AreaRangeAdvertise struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaRangeSuppress struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaInterface struct {
	Name               string
	Enable             *bool
	InstanceId         *int64
	Passive            *bool
	Metric             *int64
	Priority           *int64
	HelloInterval      *int64
	DeadCounts         *int64
	RetransmitInterval *int64
	TransitDelay       *int64
	Authentication     *string
	GrDelay            *int64
	LinkType           *ProtocolOspfv3AreaInterfaceLinkType
	Neighbor           []ProtocolOspfv3AreaInterfaceNeighbor
	Bfd                *ProtocolOspfv3AreaInterfaceBfd
	Misc               []generic.Xml
}
type ProtocolOspfv3AreaInterfaceLinkType struct {
	Broadcast *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast
	P2p       *ProtocolOspfv3AreaInterfaceLinkTypeP2p
	P2mp      *ProtocolOspfv3AreaInterfaceLinkTypeP2mp
	Misc      []generic.Xml
}
type ProtocolOspfv3AreaInterfaceLinkTypeBroadcast struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2p struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2mp struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AreaInterfaceNeighbor struct {
	Name string
	Misc []generic.Xml
}
type ProtocolOspfv3AreaInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfv3AreaVirtualLink struct {
	Name               string
	NeighborId         *string
	TransitAreaId      *string
	Enable             *bool
	InstanceId         *int64
	HelloInterval      *int64
	DeadCounts         *int64
	RetransmitInterval *int64
	TransitDelay       *int64
	Authentication     *string
	Bfd                *ProtocolOspfv3AreaVirtualLinkBfd
	Misc               []generic.Xml
}
type ProtocolOspfv3AreaVirtualLinkBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfv3AuthProfile struct {
	Name string
	Spi  *string
	Esp  *ProtocolOspfv3AuthProfileEsp
	Ah   *ProtocolOspfv3AuthProfileAh
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEsp struct {
	Authentication *ProtocolOspfv3AuthProfileEspAuthentication
	Encryption     *ProtocolOspfv3AuthProfileEspEncryption
	Misc           []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthentication struct {
	Md5    *ProtocolOspfv3AuthProfileEspAuthenticationMd5
	Sha1   *ProtocolOspfv3AuthProfileEspAuthenticationSha1
	Sha256 *ProtocolOspfv3AuthProfileEspAuthenticationSha256
	Sha384 *ProtocolOspfv3AuthProfileEspAuthenticationSha384
	Sha512 *ProtocolOspfv3AuthProfileEspAuthenticationSha512
	None   *ProtocolOspfv3AuthProfileEspAuthenticationNone
	Misc   []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationMd5 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha1 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha256 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha384 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha512 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspAuthenticationNone struct {
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileEspEncryption struct {
	Algorithm *string
	Key       *string
	Misc      []generic.Xml
}
type ProtocolOspfv3AuthProfileAh struct {
	Md5    *ProtocolOspfv3AuthProfileAhMd5
	Sha1   *ProtocolOspfv3AuthProfileAhSha1
	Sha256 *ProtocolOspfv3AuthProfileAhSha256
	Sha384 *ProtocolOspfv3AuthProfileAhSha384
	Sha512 *ProtocolOspfv3AuthProfileAhSha512
	Misc   []generic.Xml
}
type ProtocolOspfv3AuthProfileAhMd5 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileAhSha1 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileAhSha256 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileAhSha384 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3AuthProfileAhSha512 struct {
	Key  *string
	Misc []generic.Xml
}
type ProtocolOspfv3ExportRules struct {
	Name        string
	NewPathType *string
	NewTag      *string
	Metric      *int64
	Misc        []generic.Xml
}
type ProtocolOspfv3GlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolOspfv3GracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	MaxNeighborRestartTime *int64
	StrictLSAChecking      *bool
	Misc                   []generic.Xml
}
type ProtocolOspfv3Timers struct {
	LsaInterval         *float64
	SpfCalculationDelay *float64
	Misc                []generic.Xml
}
type ProtocolRedistProfile struct {
	Name     string
	Priority *int64
	Filter   *ProtocolRedistProfileFilter
	Action   *ProtocolRedistProfileAction
	Misc     []generic.Xml
}
type ProtocolRedistProfileFilter struct {
	Type        []string
	Interface   []string
	Destination []string
	Nexthop     []string
	Ospf        *ProtocolRedistProfileFilterOspf
	Bgp         *ProtocolRedistProfileFilterBgp
	Misc        []generic.Xml
}
type ProtocolRedistProfileFilterOspf struct {
	PathType []string
	Area     []string
	Tag      []string
	Misc     []generic.Xml
}
type ProtocolRedistProfileFilterBgp struct {
	Community         []string
	ExtendedCommunity []string
	Misc              []generic.Xml
}
type ProtocolRedistProfileAction struct {
	NoRedist *ProtocolRedistProfileActionNoRedist
	Redist   *ProtocolRedistProfileActionRedist
	Misc     []generic.Xml
}
type ProtocolRedistProfileActionNoRedist struct {
	Misc []generic.Xml
}
type ProtocolRedistProfileActionRedist struct {
	Misc []generic.Xml
}
type ProtocolRedistProfileIpv6 struct {
	Name     string
	Priority *int64
	Filter   *ProtocolRedistProfileIpv6Filter
	Action   *ProtocolRedistProfileIpv6Action
	Misc     []generic.Xml
}
type ProtocolRedistProfileIpv6Filter struct {
	Type        []string
	Interface   []string
	Destination []string
	Nexthop     []string
	Ospfv3      *ProtocolRedistProfileIpv6FilterOspfv3
	Bgp         *ProtocolRedistProfileIpv6FilterBgp
	Misc        []generic.Xml
}
type ProtocolRedistProfileIpv6FilterOspfv3 struct {
	PathType []string
	Area     []string
	Tag      []string
	Misc     []generic.Xml
}
type ProtocolRedistProfileIpv6FilterBgp struct {
	Community         []string
	ExtendedCommunity []string
	Misc              []generic.Xml
}
type ProtocolRedistProfileIpv6Action struct {
	NoRedist *ProtocolRedistProfileIpv6ActionNoRedist
	Redist   *ProtocolRedistProfileIpv6ActionRedist
	Misc     []generic.Xml
}
type ProtocolRedistProfileIpv6ActionNoRedist struct {
	Misc []generic.Xml
}
type ProtocolRedistProfileIpv6ActionRedist struct {
	Misc []generic.Xml
}
type ProtocolRip struct {
	AllowRedistDefaultRoute *bool
	AuthProfile             []ProtocolRipAuthProfile
	Enable                  *bool
	ExportRules             []ProtocolRipExportRules
	GlobalBfd               *ProtocolRipGlobalBfd
	Interface               []ProtocolRipInterface
	RejectDefaultRoute      *bool
	Timers                  *ProtocolRipTimers
	Misc                    []generic.Xml
}
type ProtocolRipAuthProfile struct {
	Name     string
	Password *string
	Md5      []ProtocolRipAuthProfileMd5
	Misc     []generic.Xml
}
type ProtocolRipAuthProfileMd5 struct {
	Name      string
	Key       *string
	Preferred *bool
	Misc      []generic.Xml
}
type ProtocolRipExportRules struct {
	Name   string
	Metric *int64
	Misc   []generic.Xml
}
type ProtocolRipGlobalBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolRipInterface struct {
	Name           string
	Enable         *bool
	Authentication *string
	Mode           *string
	DefaultRoute   *ProtocolRipInterfaceDefaultRoute
	Bfd            *ProtocolRipInterfaceBfd
	Misc           []generic.Xml
}
type ProtocolRipInterfaceDefaultRoute struct {
	Disable   *ProtocolRipInterfaceDefaultRouteDisable
	Advertise *ProtocolRipInterfaceDefaultRouteAdvertise
	Misc      []generic.Xml
}
type ProtocolRipInterfaceDefaultRouteDisable struct {
	Misc []generic.Xml
}
type ProtocolRipInterfaceDefaultRouteAdvertise struct {
	Metric *int64
	Misc   []generic.Xml
}
type ProtocolRipInterfaceBfd struct {
	Profile *string
	Misc    []generic.Xml
}
type ProtocolRipTimers struct {
	DeleteIntervals *int64
	ExpireIntervals *int64
	IntervalSeconds *int64
	UpdateIntervals *int64
	Misc            []generic.Xml
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
	XMLName    xml.Name         `xml:"entry"`
	Name       string           `xml:"name,attr"`
	AdminDists *adminDistsXml   `xml:"admin-dists,omitempty"`
	Ecmp       *ecmpXml         `xml:"ecmp,omitempty"`
	Interface  *util.MemberType `xml:"interface,omitempty"`
	Multicast  *multicastXml    `xml:"multicast,omitempty"`
	Protocol   *protocolXml     `xml:"protocol,omitempty"`
	Misc       []generic.Xml    `xml:",any"`
}
type adminDistsXml struct {
	Ebgp       *int64        `xml:"ebgp,omitempty"`
	Ibgp       *int64        `xml:"ibgp,omitempty"`
	OspfExt    *int64        `xml:"ospf-ext,omitempty"`
	OspfInt    *int64        `xml:"ospf-int,omitempty"`
	Ospfv3Ext  *int64        `xml:"ospfv3-ext,omitempty"`
	Ospfv3Int  *int64        `xml:"ospfv3-int,omitempty"`
	Rip        *int64        `xml:"rip,omitempty"`
	Static     *int64        `xml:"static,omitempty"`
	StaticIpv6 *int64        `xml:"static-ipv6,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type ecmpXml struct {
	Algorithm        *ecmpAlgorithmXml `xml:"algorithm,omitempty"`
	Enable           *string           `xml:"enable,omitempty"`
	MaxPath          *int64            `xml:"max-path,omitempty"`
	StrictSourcePath *string           `xml:"strict-source-path,omitempty"`
	SymmetricReturn  *string           `xml:"symmetric-return,omitempty"`
	Misc             []generic.Xml     `xml:",any"`
}
type ecmpAlgorithmXml struct {
	BalancedRoundRobin *ecmpAlgorithmBalancedRoundRobinXml `xml:"balanced-round-robin,omitempty"`
	IpHash             *ecmpAlgorithmIpHashXml             `xml:"ip-hash,omitempty"`
	IpModulo           *ecmpAlgorithmIpModuloXml           `xml:"ip-modulo,omitempty"`
	WeightedRoundRobin *ecmpAlgorithmWeightedRoundRobinXml `xml:"weighted-round-robin,omitempty"`
	Misc               []generic.Xml                       `xml:",any"`
}
type ecmpAlgorithmBalancedRoundRobinXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ecmpAlgorithmIpHashXml struct {
	HashSeed *int64        `xml:"hash-seed,omitempty"`
	SrcOnly  *string       `xml:"src-only,omitempty"`
	UsePort  *string       `xml:"use-port,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type ecmpAlgorithmIpModuloXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ecmpAlgorithmWeightedRoundRobinXml struct {
	Interface *ecmpAlgorithmWeightedRoundRobinInterfaceContainerXml `xml:"interface,omitempty"`
	Misc      []generic.Xml                                         `xml:",any"`
}
type ecmpAlgorithmWeightedRoundRobinInterfaceContainerXml struct {
	Entries []ecmpAlgorithmWeightedRoundRobinInterfaceXml `xml:"entry"`
}
type ecmpAlgorithmWeightedRoundRobinInterfaceXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Weight  *int64        `xml:"weight,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type multicastXml struct {
	Enable          *string                               `xml:"enable,omitempty"`
	InterfaceGroup  *multicastInterfaceGroupContainerXml  `xml:"interface-group,omitempty"`
	RouteAgeoutTime *int64                                `xml:"route-ageout-time,omitempty"`
	Rp              *multicastRpXml                       `xml:"rp,omitempty"`
	SptThreshold    *multicastSptThresholdContainerXml    `xml:"spt-threshold,omitempty"`
	SsmAddressSpace *multicastSsmAddressSpaceContainerXml `xml:"ssm-address-space,omitempty"`
	Misc            []generic.Xml                         `xml:",any"`
}
type multicastInterfaceGroupContainerXml struct {
	Entries []multicastInterfaceGroupXml `xml:"entry"`
}
type multicastInterfaceGroupXml struct {
	XMLName         xml.Name                                   `xml:"entry"`
	Name            string                                     `xml:"name,attr"`
	Description     *string                                    `xml:"description,omitempty"`
	Interface       *util.MemberType                           `xml:"interface,omitempty"`
	GroupPermission *multicastInterfaceGroupGroupPermissionXml `xml:"group-permission,omitempty"`
	Igmp            *multicastInterfaceGroupIgmpXml            `xml:"igmp,omitempty"`
	Pim             *multicastInterfaceGroupPimXml             `xml:"pim,omitempty"`
	Misc            []generic.Xml                              `xml:",any"`
}
type multicastInterfaceGroupGroupPermissionXml struct {
	AnySourceMulticast      *multicastInterfaceGroupGroupPermissionAnySourceMulticastContainerXml      `xml:"any-source-multicast,omitempty"`
	SourceSpecificMulticast *multicastInterfaceGroupGroupPermissionSourceSpecificMulticastContainerXml `xml:"source-specific-multicast,omitempty"`
	Misc                    []generic.Xml                                                              `xml:",any"`
}
type multicastInterfaceGroupGroupPermissionAnySourceMulticastContainerXml struct {
	Entries []multicastInterfaceGroupGroupPermissionAnySourceMulticastXml `xml:"entry"`
}
type multicastInterfaceGroupGroupPermissionAnySourceMulticastXml struct {
	XMLName      xml.Name      `xml:"entry"`
	Name         string        `xml:"name,attr"`
	GroupAddress *string       `xml:"group-address,omitempty"`
	Included     *string       `xml:"included,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type multicastInterfaceGroupGroupPermissionSourceSpecificMulticastContainerXml struct {
	Entries []multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml `xml:"entry"`
}
type multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml struct {
	XMLName       xml.Name      `xml:"entry"`
	Name          string        `xml:"name,attr"`
	GroupAddress  *string       `xml:"group-address,omitempty"`
	SourceAddress *string       `xml:"source-address,omitempty"`
	Included      *string       `xml:"included,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type multicastInterfaceGroupIgmpXml struct {
	Enable                  *string       `xml:"enable,omitempty"`
	Version                 *string       `xml:"version,omitempty"`
	MaxQueryResponseTime    *float64      `xml:"max-query-response-time,omitempty"`
	QueryInterval           *int64        `xml:"query-interval,omitempty"`
	LastMemberQueryInterval *float64      `xml:"last-member-query-interval,omitempty"`
	ImmediateLeave          *string       `xml:"immediate-leave,omitempty"`
	Robustness              *string       `xml:"robustness,omitempty"`
	MaxGroups               *string       `xml:"max-groups,omitempty"`
	MaxSources              *string       `xml:"max-sources,omitempty"`
	RouterAlertPolicing     *string       `xml:"router-alert-policing,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type multicastInterfaceGroupPimXml struct {
	Enable            *string                                                 `xml:"enable,omitempty"`
	AssertInterval    *int64                                                  `xml:"assert-interval,omitempty"`
	HelloInterval     *int64                                                  `xml:"hello-interval,omitempty"`
	JoinPruneInterval *int64                                                  `xml:"join-prune-interval,omitempty"`
	DrPriority        *int64                                                  `xml:"dr-priority,omitempty"`
	BsrBorder         *string                                                 `xml:"bsr-border,omitempty"`
	AllowedNeighbors  *multicastInterfaceGroupPimAllowedNeighborsContainerXml `xml:"allowed-neighbors,omitempty"`
	Misc              []generic.Xml                                           `xml:",any"`
}
type multicastInterfaceGroupPimAllowedNeighborsContainerXml struct {
	Entries []multicastInterfaceGroupPimAllowedNeighborsXml `xml:"entry"`
}
type multicastInterfaceGroupPimAllowedNeighborsXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type multicastRpXml struct {
	ExternalRp *multicastRpExternalRpContainerXml `xml:"external-rp,omitempty"`
	LocalRp    *multicastRpLocalRpXml             `xml:"local-rp,omitempty"`
	Misc       []generic.Xml                      `xml:",any"`
}
type multicastRpExternalRpContainerXml struct {
	Entries []multicastRpExternalRpXml `xml:"entry"`
}
type multicastRpExternalRpXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	GroupAddresses *util.MemberType `xml:"group-addresses,omitempty"`
	Override       *string          `xml:"override,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
}
type multicastRpLocalRpXml struct {
	CandidateRp *multicastRpLocalRpCandidateRpXml `xml:"candidate-rp,omitempty"`
	StaticRp    *multicastRpLocalRpStaticRpXml    `xml:"static-rp,omitempty"`
	Misc        []generic.Xml                     `xml:",any"`
}
type multicastRpLocalRpCandidateRpXml struct {
	Address               *string          `xml:"address,omitempty"`
	AdvertisementInterval *int64           `xml:"advertisement-interval,omitempty"`
	GroupAddresses        *util.MemberType `xml:"group-addresses,omitempty"`
	Interface             *string          `xml:"interface,omitempty"`
	Priority              *int64           `xml:"priority,omitempty"`
	Misc                  []generic.Xml    `xml:",any"`
}
type multicastRpLocalRpStaticRpXml struct {
	Address        *string          `xml:"address,omitempty"`
	GroupAddresses *util.MemberType `xml:"group-addresses,omitempty"`
	Interface      *string          `xml:"interface,omitempty"`
	Override       *string          `xml:"override,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
}
type multicastSptThresholdContainerXml struct {
	Entries []multicastSptThresholdXml `xml:"entry"`
}
type multicastSptThresholdXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Threshold *string       `xml:"threshold,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type multicastSsmAddressSpaceContainerXml struct {
	Entries []multicastSsmAddressSpaceXml `xml:"entry"`
}
type multicastSsmAddressSpaceXml struct {
	XMLName      xml.Name      `xml:"entry"`
	Name         string        `xml:"name,attr"`
	GroupAddress *string       `xml:"group-address,omitempty"`
	Included     *string       `xml:"included,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type protocolXml struct {
	Bgp               *protocolBgpXml                        `xml:"bgp,omitempty"`
	Ospf              *protocolOspfXml                       `xml:"ospf,omitempty"`
	Ospfv3            *protocolOspfv3Xml                     `xml:"ospfv3,omitempty"`
	RedistProfile     *protocolRedistProfileContainerXml     `xml:"redist-profile,omitempty"`
	RedistProfileIpv6 *protocolRedistProfileIpv6ContainerXml `xml:"redist-profile-ipv6,omitempty"`
	Rip               *protocolRipXml                        `xml:"rip,omitempty"`
	Misc              []generic.Xml                          `xml:",any"`
}
type protocolBgpXml struct {
	AllowRedistDefaultRoute *string                                  `xml:"allow-redist-default-route,omitempty"`
	AuthProfile             *protocolBgpAuthProfileContainerXml      `xml:"auth-profile,omitempty"`
	DampeningProfile        *protocolBgpDampeningProfileContainerXml `xml:"dampening-profile,omitempty"`
	EcmpMultiAs             *string                                  `xml:"ecmp-multi-as,omitempty"`
	Enable                  *string                                  `xml:"enable,omitempty"`
	EnforceFirstAs          *string                                  `xml:"enforce-first-as,omitempty"`
	GlobalBfd               *protocolBgpGlobalBfdXml                 `xml:"global-bfd,omitempty"`
	InstallRoute            *string                                  `xml:"install-route,omitempty"`
	LocalAs                 *string                                  `xml:"local-as,omitempty"`
	PeerGroup               *protocolBgpPeerGroupContainerXml        `xml:"peer-group,omitempty"`
	Policy                  *protocolBgpPolicyXml                    `xml:"policy,omitempty"`
	RedistRules             *protocolBgpRedistRulesContainerXml      `xml:"redist-rules,omitempty"`
	RejectDefaultRoute      *string                                  `xml:"reject-default-route,omitempty"`
	RouterId                *string                                  `xml:"router-id,omitempty"`
	RoutingOptions          *protocolBgpRoutingOptionsXml            `xml:"routing-options,omitempty"`
	Misc                    []generic.Xml                            `xml:",any"`
}
type protocolBgpAuthProfileContainerXml struct {
	Entries []protocolBgpAuthProfileXml `xml:"entry"`
}
type protocolBgpAuthProfileXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Secret  *string       `xml:"secret,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpDampeningProfileContainerXml struct {
	Entries []protocolBgpDampeningProfileXml `xml:"entry"`
}
type protocolBgpDampeningProfileXml struct {
	XMLName                  xml.Name      `xml:"entry"`
	Name                     string        `xml:"name,attr"`
	Enable                   *string       `xml:"enable,omitempty"`
	Cutoff                   *float64      `xml:"cutoff,omitempty"`
	Reuse                    *float64      `xml:"reuse,omitempty"`
	MaxHoldTime              *int64        `xml:"max-hold-time,omitempty"`
	DecayHalfLifeReachable   *int64        `xml:"decay-half-life-reachable,omitempty"`
	DecayHalfLifeUnreachable *int64        `xml:"decay-half-life-unreachable,omitempty"`
	Misc                     []generic.Xml `xml:",any"`
}
type protocolBgpGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupContainerXml struct {
	Entries []protocolBgpPeerGroupXml `xml:"entry"`
}
type protocolBgpPeerGroupXml struct {
	XMLName                 xml.Name                              `xml:"entry"`
	Name                    string                                `xml:"name,attr"`
	Enable                  *string                               `xml:"enable,omitempty"`
	AggregatedConfedAsPath  *string                               `xml:"aggregated-confed-as-path,omitempty"`
	SoftResetWithStoredInfo *string                               `xml:"soft-reset-with-stored-info,omitempty"`
	Type                    *protocolBgpPeerGroupTypeXml          `xml:"type,omitempty"`
	Peer                    *protocolBgpPeerGroupPeerContainerXml `xml:"peer,omitempty"`
	Misc                    []generic.Xml                         `xml:",any"`
}
type protocolBgpPeerGroupTypeXml struct {
	Ibgp       *protocolBgpPeerGroupTypeIbgpXml       `xml:"ibgp,omitempty"`
	EbgpConfed *protocolBgpPeerGroupTypeEbgpConfedXml `xml:"ebgp-confed,omitempty"`
	IbgpConfed *protocolBgpPeerGroupTypeIbgpConfedXml `xml:"ibgp-confed,omitempty"`
	Ebgp       *protocolBgpPeerGroupTypeEbgpXml       `xml:"ebgp,omitempty"`
	Misc       []generic.Xml                          `xml:",any"`
}
type protocolBgpPeerGroupTypeIbgpXml struct {
	ExportNexthop *string       `xml:"export-nexthop,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupTypeEbgpConfedXml struct {
	ExportNexthop *string       `xml:"export-nexthop,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupTypeIbgpConfedXml struct {
	ExportNexthop *string       `xml:"export-nexthop,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupTypeEbgpXml struct {
	ImportNexthop   *string       `xml:"import-nexthop,omitempty"`
	ExportNexthop   *string       `xml:"export-nexthop,omitempty"`
	RemovePrivateAs *string       `xml:"remove-private-as,omitempty"`
	Misc            []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerContainerXml struct {
	Entries []protocolBgpPeerGroupPeerXml `xml:"entry"`
}
type protocolBgpPeerGroupPeerXml struct {
	XMLName                           xml.Name                                                      `xml:"entry"`
	Name                              string                                                        `xml:"name,attr"`
	Enable                            *string                                                       `xml:"enable,omitempty"`
	PeerAs                            *string                                                       `xml:"peer-as,omitempty"`
	EnableMpBgp                       *string                                                       `xml:"enable-mp-bgp,omitempty"`
	AddressFamilyIdentifier           *string                                                       `xml:"address-family-identifier,omitempty"`
	EnableSenderSideLoopDetection     *string                                                       `xml:"enable-sender-side-loop-detection,omitempty"`
	ReflectorClient                   *string                                                       `xml:"reflector-client,omitempty"`
	PeeringType                       *string                                                       `xml:"peering-type,omitempty"`
	MaxPrefixes                       *string                                                       `xml:"max-prefixes,omitempty"`
	SubsequentAddressFamilyIdentifier *protocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml `xml:"subsequent-address-family-identifier,omitempty"`
	LocalAddress                      *protocolBgpPeerGroupPeerLocalAddressXml                      `xml:"local-address,omitempty"`
	PeerAddress                       *protocolBgpPeerGroupPeerPeerAddressXml                       `xml:"peer-address,omitempty"`
	ConnectionOptions                 *protocolBgpPeerGroupPeerConnectionOptionsXml                 `xml:"connection-options,omitempty"`
	Bfd                               *protocolBgpPeerGroupPeerBfdXml                               `xml:"bfd,omitempty"`
	Misc                              []generic.Xml                                                 `xml:",any"`
}
type protocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml struct {
	Unicast   *string       `xml:"unicast,omitempty"`
	Multicast *string       `xml:"multicast,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerLocalAddressXml struct {
	Interface *string       `xml:"interface,omitempty"`
	Ip        *string       `xml:"ip,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerPeerAddressXml struct {
	Ip   *string       `xml:"ip,omitempty"`
	Fqdn *string       `xml:"fqdn,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerConnectionOptionsXml struct {
	Authentication        *string                                                            `xml:"authentication,omitempty"`
	KeepAliveInterval     *string                                                            `xml:"keep-alive-interval,omitempty"`
	MinRouteAdvInterval   *int64                                                             `xml:"min-route-adv-interval,omitempty"`
	Multihop              *int64                                                             `xml:"multihop,omitempty"`
	OpenDelayTime         *int64                                                             `xml:"open-delay-time,omitempty"`
	HoldTime              *string                                                            `xml:"hold-time,omitempty"`
	IdleHoldTime          *int64                                                             `xml:"idle-hold-time,omitempty"`
	IncomingBgpConnection *protocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml `xml:"incoming-bgp-connection,omitempty"`
	OutgoingBgpConnection *protocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml `xml:"outgoing-bgp-connection,omitempty"`
	Misc                  []generic.Xml                                                      `xml:",any"`
}
type protocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml struct {
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Allow      *string       `xml:"allow,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml struct {
	LocalPort *int64        `xml:"local-port,omitempty"`
	Allow     *string       `xml:"allow,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolBgpPeerGroupPeerBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyXml struct {
	Aggregation              *protocolBgpPolicyAggregationXml              `xml:"aggregation,omitempty"`
	ConditionalAdvertisement *protocolBgpPolicyConditionalAdvertisementXml `xml:"conditional-advertisement,omitempty"`
	Export                   *protocolBgpPolicyExportXml                   `xml:"export,omitempty"`
	Import                   *protocolBgpPolicyImportXml                   `xml:"import,omitempty"`
	Misc                     []generic.Xml                                 `xml:",any"`
}
type protocolBgpPolicyAggregationXml struct {
	Address *protocolBgpPolicyAggregationAddressContainerXml `xml:"address,omitempty"`
	Misc    []generic.Xml                                    `xml:",any"`
}
type protocolBgpPolicyAggregationAddressContainerXml struct {
	Entries []protocolBgpPolicyAggregationAddressXml `xml:"entry"`
}
type protocolBgpPolicyAggregationAddressXml struct {
	XMLName                  xml.Name                                                         `xml:"entry"`
	Name                     string                                                           `xml:"name,attr"`
	Prefix                   *string                                                          `xml:"prefix,omitempty"`
	Enable                   *string                                                          `xml:"enable,omitempty"`
	Summary                  *string                                                          `xml:"summary,omitempty"`
	AsSet                    *string                                                          `xml:"as-set,omitempty"`
	AggregateRouteAttributes *protocolBgpPolicyAggregationAddressAggregateRouteAttributesXml  `xml:"aggregate-route-attributes,omitempty"`
	SuppressFilters          *protocolBgpPolicyAggregationAddressSuppressFiltersContainerXml  `xml:"suppress-filters,omitempty"`
	AdvertiseFilters         *protocolBgpPolicyAggregationAddressAdvertiseFiltersContainerXml `xml:"advertise-filters,omitempty"`
	Misc                     []generic.Xml                                                    `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesXml struct {
	LocalPreference   *int64                                                                           `xml:"local-preference,omitempty"`
	Med               *int64                                                                           `xml:"med,omitempty"`
	Weight            *int64                                                                           `xml:"weight,omitempty"`
	Nexthop           *string                                                                          `xml:"nexthop,omitempty"`
	Origin            *string                                                                          `xml:"origin,omitempty"`
	AsPathLimit       *int64                                                                           `xml:"as-path-limit,omitempty"`
	AsPath            *protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml            `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                                    `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml struct {
	None    *protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml `xml:"none,omitempty"`
	Prepend *int64                                                                    `xml:"prepend,omitempty"`
	Misc    []generic.Xml                                                             `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml struct {
	None        *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                           `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                                  `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                                  `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                                     `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml struct {
	None        *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                                   `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                                          `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                                          `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                                             `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersContainerXml struct {
	Entries []protocolBgpPolicyAggregationAddressSuppressFiltersXml `xml:"entry"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersXml struct {
	XMLName xml.Name                                                    `xml:"entry"`
	Name    string                                                      `xml:"name,attr"`
	Enable  *string                                                     `xml:"enable,omitempty"`
	Match   *protocolBgpPolicyAggregationAddressSuppressFiltersMatchXml `xml:"match,omitempty"`
	Misc    []generic.Xml                                               `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchXml struct {
	RouteTable        *string                                                                           `xml:"route-table,omitempty"`
	Med               *int64                                                                            `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                                                  `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                                                  `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                                     `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Exact   *string       `xml:"exact,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersContainerXml struct {
	Entries []protocolBgpPolicyAggregationAddressAdvertiseFiltersXml `xml:"entry"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersXml struct {
	XMLName xml.Name                                                     `xml:"entry"`
	Name    string                                                       `xml:"name,attr"`
	Enable  *string                                                      `xml:"enable,omitempty"`
	Match   *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml `xml:"match,omitempty"`
	Misc    []generic.Xml                                                `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml struct {
	RouteTable        *string                                                                            `xml:"route-table,omitempty"`
	Med               *int64                                                                             `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                                                   `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                                                   `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                                      `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Exact   *string       `xml:"exact,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementXml struct {
	Policy *protocolBgpPolicyConditionalAdvertisementPolicyContainerXml `xml:"policy,omitempty"`
	Misc   []generic.Xml                                                `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyContainerXml struct {
	Entries []protocolBgpPolicyConditionalAdvertisementPolicyXml `xml:"entry"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyXml struct {
	XMLName          xml.Name                                                                     `xml:"entry"`
	Name             string                                                                       `xml:"name,attr"`
	Enable           *string                                                                      `xml:"enable,omitempty"`
	UsedBy           *util.MemberType                                                             `xml:"used-by,omitempty"`
	NonExistFilters  *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersContainerXml  `xml:"non-exist-filters,omitempty"`
	AdvertiseFilters *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersContainerXml `xml:"advertise-filters,omitempty"`
	Misc             []generic.Xml                                                                `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersContainerXml struct {
	Entries []protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml `xml:"entry"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml struct {
	XMLName xml.Name                                                                `xml:"entry"`
	Name    string                                                                  `xml:"name,attr"`
	Enable  *string                                                                 `xml:"enable,omitempty"`
	Match   *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml `xml:"match,omitempty"`
	Misc    []generic.Xml                                                           `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml struct {
	RouteTable        *string                                                                                       `xml:"route-table,omitempty"`
	Med               *int64                                                                                        `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                                                              `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                                                              `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                                                 `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersContainerXml struct {
	Entries []protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml `xml:"entry"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml struct {
	XMLName xml.Name                                                                 `xml:"entry"`
	Name    string                                                                   `xml:"name,attr"`
	Enable  *string                                                                  `xml:"enable,omitempty"`
	Match   *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml `xml:"match,omitempty"`
	Misc    []generic.Xml                                                            `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml struct {
	RouteTable        *string                                                                                        `xml:"route-table,omitempty"`
	Med               *int64                                                                                         `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                                                               `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                                                               `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                                                  `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportXml struct {
	Rules *protocolBgpPolicyExportRulesContainerXml `xml:"rules,omitempty"`
	Misc  []generic.Xml                             `xml:",any"`
}
type protocolBgpPolicyExportRulesContainerXml struct {
	Entries []protocolBgpPolicyExportRulesXml `xml:"entry"`
}
type protocolBgpPolicyExportRulesXml struct {
	XMLName xml.Name                               `xml:"entry"`
	Name    string                                 `xml:"name,attr"`
	Enable  *string                                `xml:"enable,omitempty"`
	UsedBy  *util.MemberType                       `xml:"used-by,omitempty"`
	Match   *protocolBgpPolicyExportRulesMatchXml  `xml:"match,omitempty"`
	Action  *protocolBgpPolicyExportRulesActionXml `xml:"action,omitempty"`
	Misc    []generic.Xml                          `xml:",any"`
}
type protocolBgpPolicyExportRulesMatchXml struct {
	RouteTable        *string                                                     `xml:"route-table,omitempty"`
	Med               *int64                                                      `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyExportRulesMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                            `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                            `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyExportRulesMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyExportRulesMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyExportRulesMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                               `xml:",any"`
}
type protocolBgpPolicyExportRulesMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyExportRulesMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyExportRulesMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Exact   *string       `xml:"exact,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionXml struct {
	Deny  *protocolBgpPolicyExportRulesActionDenyXml  `xml:"deny,omitempty"`
	Allow *protocolBgpPolicyExportRulesActionAllowXml `xml:"allow,omitempty"`
	Misc  []generic.Xml                               `xml:",any"`
}
type protocolBgpPolicyExportRulesActionDenyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowXml struct {
	Update *protocolBgpPolicyExportRulesActionAllowUpdateXml `xml:"update,omitempty"`
	Misc   []generic.Xml                                     `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateXml struct {
	LocalPreference   *int64                                                             `xml:"local-preference,omitempty"`
	Med               *int64                                                             `xml:"med,omitempty"`
	Nexthop           *string                                                            `xml:"nexthop,omitempty"`
	Origin            *string                                                            `xml:"origin,omitempty"`
	AsPathLimit       *int64                                                             `xml:"as-path-limit,omitempty"`
	AsPath            *protocolBgpPolicyExportRulesActionAllowUpdateAsPathXml            `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyExportRulesActionAllowUpdateCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                      `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateAsPathXml struct {
	None             *protocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml   `xml:"none,omitempty"`
	Remove           *protocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml `xml:"remove,omitempty"`
	Prepend          *int64                                                        `xml:"prepend,omitempty"`
	RemoveAndPrepend *int64                                                        `xml:"remove-and-prepend,omitempty"`
	Misc             []generic.Xml                                                 `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateCommunityXml struct {
	None        *protocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                             `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                    `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                    `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                       `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml struct {
	None        *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                     `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                            `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                            `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                               `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportXml struct {
	Rules *protocolBgpPolicyImportRulesContainerXml `xml:"rules,omitempty"`
	Misc  []generic.Xml                             `xml:",any"`
}
type protocolBgpPolicyImportRulesContainerXml struct {
	Entries []protocolBgpPolicyImportRulesXml `xml:"entry"`
}
type protocolBgpPolicyImportRulesXml struct {
	XMLName xml.Name                               `xml:"entry"`
	Name    string                                 `xml:"name,attr"`
	Enable  *string                                `xml:"enable,omitempty"`
	UsedBy  *util.MemberType                       `xml:"used-by,omitempty"`
	Match   *protocolBgpPolicyImportRulesMatchXml  `xml:"match,omitempty"`
	Action  *protocolBgpPolicyImportRulesActionXml `xml:"action,omitempty"`
	Misc    []generic.Xml                          `xml:",any"`
}
type protocolBgpPolicyImportRulesMatchXml struct {
	RouteTable        *string                                                     `xml:"route-table,omitempty"`
	Med               *int64                                                      `xml:"med,omitempty"`
	AddressPrefix     *protocolBgpPolicyImportRulesMatchAddressPrefixContainerXml `xml:"address-prefix,omitempty"`
	Nexthop           *util.MemberType                                            `xml:"nexthop,omitempty"`
	FromPeer          *util.MemberType                                            `xml:"from-peer,omitempty"`
	AsPath            *protocolBgpPolicyImportRulesMatchAsPathXml                 `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyImportRulesMatchCommunityXml              `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyImportRulesMatchExtendedCommunityXml      `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                               `xml:",any"`
}
type protocolBgpPolicyImportRulesMatchAddressPrefixContainerXml struct {
	Entries []protocolBgpPolicyImportRulesMatchAddressPrefixXml `xml:"entry"`
}
type protocolBgpPolicyImportRulesMatchAddressPrefixXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Exact   *string       `xml:"exact,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesMatchAsPathXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesMatchCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesMatchExtendedCommunityXml struct {
	Regex *string       `xml:"regex,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionXml struct {
	Deny  *protocolBgpPolicyImportRulesActionDenyXml  `xml:"deny,omitempty"`
	Allow *protocolBgpPolicyImportRulesActionAllowXml `xml:"allow,omitempty"`
	Misc  []generic.Xml                               `xml:",any"`
}
type protocolBgpPolicyImportRulesActionDenyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowXml struct {
	Dampening *string                                           `xml:"dampening,omitempty"`
	Update    *protocolBgpPolicyImportRulesActionAllowUpdateXml `xml:"update,omitempty"`
	Misc      []generic.Xml                                     `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateXml struct {
	LocalPreference   *int64                                                             `xml:"local-preference,omitempty"`
	Med               *int64                                                             `xml:"med,omitempty"`
	Weight            *int64                                                             `xml:"weight,omitempty"`
	Nexthop           *string                                                            `xml:"nexthop,omitempty"`
	Origin            *string                                                            `xml:"origin,omitempty"`
	AsPathLimit       *int64                                                             `xml:"as-path-limit,omitempty"`
	AsPath            *protocolBgpPolicyImportRulesActionAllowUpdateAsPathXml            `xml:"as-path,omitempty"`
	Community         *protocolBgpPolicyImportRulesActionAllowUpdateCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml `xml:"extended-community,omitempty"`
	Misc              []generic.Xml                                                      `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateAsPathXml struct {
	None   *protocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml   `xml:"none,omitempty"`
	Remove *protocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml `xml:"remove,omitempty"`
	Misc   []generic.Xml                                                 `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateCommunityXml struct {
	None        *protocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                             `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                    `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                    `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                       `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml struct {
	None        *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml      `xml:"none,omitempty"`
	RemoveAll   *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                     `xml:"remove-regex,omitempty"`
	Append      *util.MemberType                                                            `xml:"append,omitempty"`
	Overwrite   *util.MemberType                                                            `xml:"overwrite,omitempty"`
	Misc        []generic.Xml                                                               `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolBgpRedistRulesContainerXml struct {
	Entries []protocolBgpRedistRulesXml `xml:"entry"`
}
type protocolBgpRedistRulesXml struct {
	XMLName                 xml.Name         `xml:"entry"`
	Name                    string           `xml:"name,attr"`
	AddressFamilyIdentifier *string          `xml:"address-family-identifier,omitempty"`
	RouteTable              *string          `xml:"route-table,omitempty"`
	Enable                  *string          `xml:"enable,omitempty"`
	SetOrigin               *string          `xml:"set-origin,omitempty"`
	SetMed                  *int64           `xml:"set-med,omitempty"`
	SetLocalPreference      *int64           `xml:"set-local-preference,omitempty"`
	SetAsPathLimit          *int64           `xml:"set-as-path-limit,omitempty"`
	Metric                  *int64           `xml:"metric,omitempty"`
	SetCommunity            *util.MemberType `xml:"set-community,omitempty"`
	SetExtendedCommunity    *util.MemberType `xml:"set-extended-community,omitempty"`
	Misc                    []generic.Xml    `xml:",any"`
}
type protocolBgpRoutingOptionsXml struct {
	Aggregate              *protocolBgpRoutingOptionsAggregateXml       `xml:"aggregate,omitempty"`
	AsFormat               *string                                      `xml:"as-format,omitempty"`
	ConfederationMemberAs  *string                                      `xml:"confederation-member-as,omitempty"`
	DefaultLocalPreference *int64                                       `xml:"default-local-preference,omitempty"`
	GracefulRestart        *protocolBgpRoutingOptionsGracefulRestartXml `xml:"graceful-restart,omitempty"`
	Med                    *protocolBgpRoutingOptionsMedXml             `xml:"med,omitempty"`
	ReflectorClusterId     *string                                      `xml:"reflector-cluster-id,omitempty"`
	Misc                   []generic.Xml                                `xml:",any"`
}
type protocolBgpRoutingOptionsAggregateXml struct {
	AggregateMed *string       `xml:"aggregate-med,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type protocolBgpRoutingOptionsGracefulRestartXml struct {
	Enable             *string       `xml:"enable,omitempty"`
	LocalRestartTime   *int64        `xml:"local-restart-time,omitempty"`
	MaxPeerRestartTime *int64        `xml:"max-peer-restart-time,omitempty"`
	StaleRouteTime     *int64        `xml:"stale-route-time,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type protocolBgpRoutingOptionsMedXml struct {
	AlwaysCompareMed           *string       `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string       `xml:"deterministic-med-comparison,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type protocolOspfXml struct {
	AllowRedistDefaultRoute *string                              `xml:"allow-redist-default-route,omitempty"`
	Area                    *protocolOspfAreaContainerXml        `xml:"area,omitempty"`
	AuthProfile             *protocolOspfAuthProfileContainerXml `xml:"auth-profile,omitempty"`
	Enable                  *string                              `xml:"enable,omitempty"`
	ExportRules             *protocolOspfExportRulesContainerXml `xml:"export-rules,omitempty"`
	GlobalBfd               *protocolOspfGlobalBfdXml            `xml:"global-bfd,omitempty"`
	GracefulRestart         *protocolOspfGracefulRestartXml      `xml:"graceful-restart,omitempty"`
	RejectDefaultRoute      *string                              `xml:"reject-default-route,omitempty"`
	Rfc1583                 *string                              `xml:"rfc1583,omitempty"`
	RouterId                *string                              `xml:"router-id,omitempty"`
	Timers                  *protocolOspfTimersXml               `xml:"timers,omitempty"`
	Misc                    []generic.Xml                        `xml:",any"`
}
type protocolOspfAreaContainerXml struct {
	Entries []protocolOspfAreaXml `xml:"entry"`
}
type protocolOspfAreaXml struct {
	XMLName     xml.Name                                 `xml:"entry"`
	Name        string                                   `xml:"name,attr"`
	Type        *protocolOspfAreaTypeXml                 `xml:"type,omitempty"`
	Range       *protocolOspfAreaRangeContainerXml       `xml:"range,omitempty"`
	Interface   *protocolOspfAreaInterfaceContainerXml   `xml:"interface,omitempty"`
	VirtualLink *protocolOspfAreaVirtualLinkContainerXml `xml:"virtual-link,omitempty"`
	Misc        []generic.Xml                            `xml:",any"`
}
type protocolOspfAreaTypeXml struct {
	Normal *protocolOspfAreaTypeNormalXml `xml:"normal,omitempty"`
	Stub   *protocolOspfAreaTypeStubXml   `xml:"stub,omitempty"`
	Nssa   *protocolOspfAreaTypeNssaXml   `xml:"nssa,omitempty"`
	Misc   []generic.Xml                  `xml:",any"`
}
type protocolOspfAreaTypeNormalXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeStubXml struct {
	AcceptSummary *string                                  `xml:"accept-summary,omitempty"`
	DefaultRoute  *protocolOspfAreaTypeStubDefaultRouteXml `xml:"default-route,omitempty"`
	Misc          []generic.Xml                            `xml:",any"`
}
type protocolOspfAreaTypeStubDefaultRouteXml struct {
	Disable   *protocolOspfAreaTypeStubDefaultRouteDisableXml   `xml:"disable,omitempty"`
	Advertise *protocolOspfAreaTypeStubDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Misc      []generic.Xml                                     `xml:",any"`
}
type protocolOspfAreaTypeStubDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeStubDefaultRouteAdvertiseXml struct {
	Metric *int64        `xml:"metric,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeNssaXml struct {
	AcceptSummary *string                                           `xml:"accept-summary,omitempty"`
	DefaultRoute  *protocolOspfAreaTypeNssaDefaultRouteXml          `xml:"default-route,omitempty"`
	NssaExtRange  *protocolOspfAreaTypeNssaNssaExtRangeContainerXml `xml:"nssa-ext-range,omitempty"`
	Misc          []generic.Xml                                     `xml:",any"`
}
type protocolOspfAreaTypeNssaDefaultRouteXml struct {
	Disable   *protocolOspfAreaTypeNssaDefaultRouteDisableXml   `xml:"disable,omitempty"`
	Advertise *protocolOspfAreaTypeNssaDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Misc      []generic.Xml                                     `xml:",any"`
}
type protocolOspfAreaTypeNssaDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeNssaDefaultRouteAdvertiseXml struct {
	Metric *int64        `xml:"metric,omitempty"`
	Type   *string       `xml:"type,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeNssaNssaExtRangeContainerXml struct {
	Entries []protocolOspfAreaTypeNssaNssaExtRangeXml `xml:"entry"`
}
type protocolOspfAreaTypeNssaNssaExtRangeXml struct {
	XMLName   xml.Name                                          `xml:"entry"`
	Name      string                                            `xml:"name,attr"`
	Advertise *protocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *protocolOspfAreaTypeNssaNssaExtRangeSuppressXml  `xml:"suppress,omitempty"`
	Misc      []generic.Xml                                     `xml:",any"`
}
type protocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaTypeNssaNssaExtRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaRangeContainerXml struct {
	Entries []protocolOspfAreaRangeXml `xml:"entry"`
}
type protocolOspfAreaRangeXml struct {
	XMLName   xml.Name                           `xml:"entry"`
	Name      string                             `xml:"name,attr"`
	Advertise *protocolOspfAreaRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *protocolOspfAreaRangeSuppressXml  `xml:"suppress,omitempty"`
	Misc      []generic.Xml                      `xml:",any"`
}
type protocolOspfAreaRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaInterfaceContainerXml struct {
	Entries []protocolOspfAreaInterfaceXml `xml:"entry"`
}
type protocolOspfAreaInterfaceXml struct {
	XMLName            xml.Name                                       `xml:"entry"`
	Name               string                                         `xml:"name,attr"`
	Enable             *string                                        `xml:"enable,omitempty"`
	Passive            *string                                        `xml:"passive,omitempty"`
	Metric             *int64                                         `xml:"metric,omitempty"`
	Priority           *int64                                         `xml:"priority,omitempty"`
	HelloInterval      *int64                                         `xml:"hello-interval,omitempty"`
	DeadCounts         *int64                                         `xml:"dead-counts,omitempty"`
	RetransmitInterval *int64                                         `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                                         `xml:"transit-delay,omitempty"`
	Authentication     *string                                        `xml:"authentication,omitempty"`
	GrDelay            *int64                                         `xml:"gr-delay,omitempty"`
	LinkType           *protocolOspfAreaInterfaceLinkTypeXml          `xml:"link-type,omitempty"`
	Neighbor           *protocolOspfAreaInterfaceNeighborContainerXml `xml:"neighbor,omitempty"`
	Bfd                *protocolOspfAreaInterfaceBfdXml               `xml:"bfd,omitempty"`
	Misc               []generic.Xml                                  `xml:",any"`
}
type protocolOspfAreaInterfaceLinkTypeXml struct {
	Broadcast *protocolOspfAreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2p       *protocolOspfAreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`
	P2mp      *protocolOspfAreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                                  `xml:",any"`
}
type protocolOspfAreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaInterfaceLinkTypeP2mpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfAreaInterfaceNeighborContainerXml struct {
	Entries []protocolOspfAreaInterfaceNeighborXml `xml:"entry"`
}
type protocolOspfAreaInterfaceNeighborXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfAreaInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfAreaVirtualLinkContainerXml struct {
	Entries []protocolOspfAreaVirtualLinkXml `xml:"entry"`
}
type protocolOspfAreaVirtualLinkXml struct {
	XMLName            xml.Name                           `xml:"entry"`
	Name               string                             `xml:"name,attr"`
	NeighborId         *string                            `xml:"neighbor-id,omitempty"`
	TransitAreaId      *string                            `xml:"transit-area-id,omitempty"`
	Enable             *string                            `xml:"enable,omitempty"`
	HelloInterval      *int64                             `xml:"hello-interval,omitempty"`
	DeadCounts         *int64                             `xml:"dead-counts,omitempty"`
	RetransmitInterval *int64                             `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                             `xml:"transit-delay,omitempty"`
	Authentication     *string                            `xml:"authentication,omitempty"`
	Bfd                *protocolOspfAreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	Misc               []generic.Xml                      `xml:",any"`
}
type protocolOspfAreaVirtualLinkBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfAuthProfileContainerXml struct {
	Entries []protocolOspfAuthProfileXml `xml:"entry"`
}
type protocolOspfAuthProfileXml struct {
	XMLName  xml.Name                                `xml:"entry"`
	Name     string                                  `xml:"name,attr"`
	Password *string                                 `xml:"password,omitempty"`
	Md5      *protocolOspfAuthProfileMd5ContainerXml `xml:"md5,omitempty"`
	Misc     []generic.Xml                           `xml:",any"`
}
type protocolOspfAuthProfileMd5ContainerXml struct {
	Entries []protocolOspfAuthProfileMd5Xml `xml:"entry"`
}
type protocolOspfAuthProfileMd5Xml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Key       *string       `xml:"key,omitempty"`
	Preferred *string       `xml:"preferred,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolOspfExportRulesContainerXml struct {
	Entries []protocolOspfExportRulesXml `xml:"entry"`
}
type protocolOspfExportRulesXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	NewPathType *string       `xml:"new-path-type,omitempty"`
	NewTag      *string       `xml:"new-tag,omitempty"`
	Metric      *int64        `xml:"metric,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type protocolOspfGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfGracefulRestartXml struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type protocolOspfTimersXml struct {
	LsaInterval         *float64      `xml:"lsa-interval,omitempty"`
	SpfCalculationDelay *float64      `xml:"spf-calculation-delay,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type protocolOspfv3Xml struct {
	AllowRedistDefaultRoute *string                                `xml:"allow-redist-default-route,omitempty"`
	Area                    *protocolOspfv3AreaContainerXml        `xml:"area,omitempty"`
	AuthProfile             *protocolOspfv3AuthProfileContainerXml `xml:"auth-profile,omitempty"`
	DisableTransitTraffic   *string                                `xml:"disable-transit-traffic,omitempty"`
	Enable                  *string                                `xml:"enable,omitempty"`
	ExportRules             *protocolOspfv3ExportRulesContainerXml `xml:"export-rules,omitempty"`
	GlobalBfd               *protocolOspfv3GlobalBfdXml            `xml:"global-bfd,omitempty"`
	GracefulRestart         *protocolOspfv3GracefulRestartXml      `xml:"graceful-restart,omitempty"`
	RejectDefaultRoute      *string                                `xml:"reject-default-route,omitempty"`
	RouterId                *string                                `xml:"router-id,omitempty"`
	Timers                  *protocolOspfv3TimersXml               `xml:"timers,omitempty"`
	Misc                    []generic.Xml                          `xml:",any"`
}
type protocolOspfv3AreaContainerXml struct {
	Entries []protocolOspfv3AreaXml `xml:"entry"`
}
type protocolOspfv3AreaXml struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Authentication *string                                    `xml:"authentication,omitempty"`
	Type           *protocolOspfv3AreaTypeXml                 `xml:"type,omitempty"`
	Range          *protocolOspfv3AreaRangeContainerXml       `xml:"range,omitempty"`
	Interface      *protocolOspfv3AreaInterfaceContainerXml   `xml:"interface,omitempty"`
	VirtualLink    *protocolOspfv3AreaVirtualLinkContainerXml `xml:"virtual-link,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
}
type protocolOspfv3AreaTypeXml struct {
	Normal *protocolOspfv3AreaTypeNormalXml `xml:"normal,omitempty"`
	Stub   *protocolOspfv3AreaTypeStubXml   `xml:"stub,omitempty"`
	Nssa   *protocolOspfv3AreaTypeNssaXml   `xml:"nssa,omitempty"`
	Misc   []generic.Xml                    `xml:",any"`
}
type protocolOspfv3AreaTypeNormalXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeStubXml struct {
	AcceptSummary *string                                    `xml:"accept-summary,omitempty"`
	DefaultRoute  *protocolOspfv3AreaTypeStubDefaultRouteXml `xml:"default-route,omitempty"`
	Misc          []generic.Xml                              `xml:",any"`
}
type protocolOspfv3AreaTypeStubDefaultRouteXml struct {
	Disable   *protocolOspfv3AreaTypeStubDefaultRouteDisableXml   `xml:"disable,omitempty"`
	Advertise *protocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Misc      []generic.Xml                                       `xml:",any"`
}
type protocolOspfv3AreaTypeStubDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml struct {
	Metric *int64        `xml:"metric,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeNssaXml struct {
	AcceptSummary *string                                             `xml:"accept-summary,omitempty"`
	DefaultRoute  *protocolOspfv3AreaTypeNssaDefaultRouteXml          `xml:"default-route,omitempty"`
	NssaExtRange  *protocolOspfv3AreaTypeNssaNssaExtRangeContainerXml `xml:"nssa-ext-range,omitempty"`
	Misc          []generic.Xml                                       `xml:",any"`
}
type protocolOspfv3AreaTypeNssaDefaultRouteXml struct {
	Disable   *protocolOspfv3AreaTypeNssaDefaultRouteDisableXml   `xml:"disable,omitempty"`
	Advertise *protocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Misc      []generic.Xml                                       `xml:",any"`
}
type protocolOspfv3AreaTypeNssaDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml struct {
	Metric *int64        `xml:"metric,omitempty"`
	Type   *string       `xml:"type,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeNssaNssaExtRangeContainerXml struct {
	Entries []protocolOspfv3AreaTypeNssaNssaExtRangeXml `xml:"entry"`
}
type protocolOspfv3AreaTypeNssaNssaExtRangeXml struct {
	XMLName   xml.Name                                            `xml:"entry"`
	Name      string                                              `xml:"name,attr"`
	Advertise *protocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *protocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml  `xml:"suppress,omitempty"`
	Misc      []generic.Xml                                       `xml:",any"`
}
type protocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaRangeContainerXml struct {
	Entries []protocolOspfv3AreaRangeXml `xml:"entry"`
}
type protocolOspfv3AreaRangeXml struct {
	XMLName   xml.Name                             `xml:"entry"`
	Name      string                               `xml:"name,attr"`
	Advertise *protocolOspfv3AreaRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *protocolOspfv3AreaRangeSuppressXml  `xml:"suppress,omitempty"`
	Misc      []generic.Xml                        `xml:",any"`
}
type protocolOspfv3AreaRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaInterfaceContainerXml struct {
	Entries []protocolOspfv3AreaInterfaceXml `xml:"entry"`
}
type protocolOspfv3AreaInterfaceXml struct {
	XMLName            xml.Name                                         `xml:"entry"`
	Name               string                                           `xml:"name,attr"`
	Enable             *string                                          `xml:"enable,omitempty"`
	InstanceId         *int64                                           `xml:"instance-id,omitempty"`
	Passive            *string                                          `xml:"passive,omitempty"`
	Metric             *int64                                           `xml:"metric,omitempty"`
	Priority           *int64                                           `xml:"priority,omitempty"`
	HelloInterval      *int64                                           `xml:"hello-interval,omitempty"`
	DeadCounts         *int64                                           `xml:"dead-counts,omitempty"`
	RetransmitInterval *int64                                           `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                                           `xml:"transit-delay,omitempty"`
	Authentication     *string                                          `xml:"authentication,omitempty"`
	GrDelay            *int64                                           `xml:"gr-delay,omitempty"`
	LinkType           *protocolOspfv3AreaInterfaceLinkTypeXml          `xml:"link-type,omitempty"`
	Neighbor           *protocolOspfv3AreaInterfaceNeighborContainerXml `xml:"neighbor,omitempty"`
	Bfd                *protocolOspfv3AreaInterfaceBfdXml               `xml:"bfd,omitempty"`
	Misc               []generic.Xml                                    `xml:",any"`
}
type protocolOspfv3AreaInterfaceLinkTypeXml struct {
	Broadcast *protocolOspfv3AreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2p       *protocolOspfv3AreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`
	P2mp      *protocolOspfv3AreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	Misc      []generic.Xml                                    `xml:",any"`
}
type protocolOspfv3AreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaInterfaceLinkTypeP2mpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaInterfaceNeighborContainerXml struct {
	Entries []protocolOspfv3AreaInterfaceNeighborXml `xml:"entry"`
}
type protocolOspfv3AreaInterfaceNeighborXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfv3AreaVirtualLinkContainerXml struct {
	Entries []protocolOspfv3AreaVirtualLinkXml `xml:"entry"`
}
type protocolOspfv3AreaVirtualLinkXml struct {
	XMLName            xml.Name                             `xml:"entry"`
	Name               string                               `xml:"name,attr"`
	NeighborId         *string                              `xml:"neighbor-id,omitempty"`
	TransitAreaId      *string                              `xml:"transit-area-id,omitempty"`
	Enable             *string                              `xml:"enable,omitempty"`
	InstanceId         *int64                               `xml:"instance-id,omitempty"`
	HelloInterval      *int64                               `xml:"hello-interval,omitempty"`
	DeadCounts         *int64                               `xml:"dead-counts,omitempty"`
	RetransmitInterval *int64                               `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                               `xml:"transit-delay,omitempty"`
	Authentication     *string                              `xml:"authentication,omitempty"`
	Bfd                *protocolOspfv3AreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	Misc               []generic.Xml                        `xml:",any"`
}
type protocolOspfv3AreaVirtualLinkBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileContainerXml struct {
	Entries []protocolOspfv3AuthProfileXml `xml:"entry"`
}
type protocolOspfv3AuthProfileXml struct {
	XMLName xml.Name                         `xml:"entry"`
	Name    string                           `xml:"name,attr"`
	Spi     *string                          `xml:"spi,omitempty"`
	Esp     *protocolOspfv3AuthProfileEspXml `xml:"esp,omitempty"`
	Ah      *protocolOspfv3AuthProfileAhXml  `xml:"ah,omitempty"`
	Misc    []generic.Xml                    `xml:",any"`
}
type protocolOspfv3AuthProfileEspXml struct {
	Authentication *protocolOspfv3AuthProfileEspAuthenticationXml `xml:"authentication,omitempty"`
	Encryption     *protocolOspfv3AuthProfileEspEncryptionXml     `xml:"encryption,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationXml struct {
	Md5    *protocolOspfv3AuthProfileEspAuthenticationMd5Xml    `xml:"md5,omitempty"`
	Sha1   *protocolOspfv3AuthProfileEspAuthenticationSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *protocolOspfv3AuthProfileEspAuthenticationSha256Xml `xml:"sha256,omitempty"`
	Sha384 *protocolOspfv3AuthProfileEspAuthenticationSha384Xml `xml:"sha384,omitempty"`
	Sha512 *protocolOspfv3AuthProfileEspAuthenticationSha512Xml `xml:"sha512,omitempty"`
	None   *protocolOspfv3AuthProfileEspAuthenticationNoneXml   `xml:"none,omitempty"`
	Misc   []generic.Xml                                        `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationMd5Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationSha1Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationSha256Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationSha384Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationSha512Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspAuthenticationNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileEspEncryptionXml struct {
	Algorithm *string       `xml:"algorithm,omitempty"`
	Key       *string       `xml:"key,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileAhXml struct {
	Md5    *protocolOspfv3AuthProfileAhMd5Xml    `xml:"md5,omitempty"`
	Sha1   *protocolOspfv3AuthProfileAhSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *protocolOspfv3AuthProfileAhSha256Xml `xml:"sha256,omitempty"`
	Sha384 *protocolOspfv3AuthProfileAhSha384Xml `xml:"sha384,omitempty"`
	Sha512 *protocolOspfv3AuthProfileAhSha512Xml `xml:"sha512,omitempty"`
	Misc   []generic.Xml                         `xml:",any"`
}
type protocolOspfv3AuthProfileAhMd5Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileAhSha1Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileAhSha256Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileAhSha384Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3AuthProfileAhSha512Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type protocolOspfv3ExportRulesContainerXml struct {
	Entries []protocolOspfv3ExportRulesXml `xml:"entry"`
}
type protocolOspfv3ExportRulesXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	NewPathType *string       `xml:"new-path-type,omitempty"`
	NewTag      *string       `xml:"new-tag,omitempty"`
	Metric      *int64        `xml:"metric,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type protocolOspfv3GlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolOspfv3GracefulRestartXml struct {
	Enable                 *string       `xml:"enable,omitempty"`
	GracePeriod            *int64        `xml:"grace-period,omitempty"`
	HelperEnable           *string       `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64        `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string       `xml:"strict-LSA-checking,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type protocolOspfv3TimersXml struct {
	LsaInterval         *float64      `xml:"lsa-interval,omitempty"`
	SpfCalculationDelay *float64      `xml:"spf-calculation-delay,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type protocolRedistProfileContainerXml struct {
	Entries []protocolRedistProfileXml `xml:"entry"`
}
type protocolRedistProfileXml struct {
	XMLName  xml.Name                        `xml:"entry"`
	Name     string                          `xml:"name,attr"`
	Priority *int64                          `xml:"priority,omitempty"`
	Filter   *protocolRedistProfileFilterXml `xml:"filter,omitempty"`
	Action   *protocolRedistProfileActionXml `xml:"action,omitempty"`
	Misc     []generic.Xml                   `xml:",any"`
}
type protocolRedistProfileFilterXml struct {
	Type        *util.MemberType                    `xml:"type,omitempty"`
	Interface   *util.MemberType                    `xml:"interface,omitempty"`
	Destination *util.MemberType                    `xml:"destination,omitempty"`
	Nexthop     *util.MemberType                    `xml:"nexthop,omitempty"`
	Ospf        *protocolRedistProfileFilterOspfXml `xml:"ospf,omitempty"`
	Bgp         *protocolRedistProfileFilterBgpXml  `xml:"bgp,omitempty"`
	Misc        []generic.Xml                       `xml:",any"`
}
type protocolRedistProfileFilterOspfXml struct {
	PathType *util.MemberType `xml:"path-type,omitempty"`
	Area     *util.MemberType `xml:"area,omitempty"`
	Tag      *util.MemberType `xml:"tag,omitempty"`
	Misc     []generic.Xml    `xml:",any"`
}
type protocolRedistProfileFilterBgpXml struct {
	Community         *util.MemberType `xml:"community,omitempty"`
	ExtendedCommunity *util.MemberType `xml:"extended-community,omitempty"`
	Misc              []generic.Xml    `xml:",any"`
}
type protocolRedistProfileActionXml struct {
	NoRedist *protocolRedistProfileActionNoRedistXml `xml:"no-redist,omitempty"`
	Redist   *protocolRedistProfileActionRedistXml   `xml:"redist,omitempty"`
	Misc     []generic.Xml                           `xml:",any"`
}
type protocolRedistProfileActionNoRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolRedistProfileActionRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolRedistProfileIpv6ContainerXml struct {
	Entries []protocolRedistProfileIpv6Xml `xml:"entry"`
}
type protocolRedistProfileIpv6Xml struct {
	XMLName  xml.Name                            `xml:"entry"`
	Name     string                              `xml:"name,attr"`
	Priority *int64                              `xml:"priority,omitempty"`
	Filter   *protocolRedistProfileIpv6FilterXml `xml:"filter,omitempty"`
	Action   *protocolRedistProfileIpv6ActionXml `xml:"action,omitempty"`
	Misc     []generic.Xml                       `xml:",any"`
}
type protocolRedistProfileIpv6FilterXml struct {
	Type        *util.MemberType                          `xml:"type,omitempty"`
	Interface   *util.MemberType                          `xml:"interface,omitempty"`
	Destination *util.MemberType                          `xml:"destination,omitempty"`
	Nexthop     *util.MemberType                          `xml:"nexthop,omitempty"`
	Ospfv3      *protocolRedistProfileIpv6FilterOspfv3Xml `xml:"ospfv3,omitempty"`
	Bgp         *protocolRedistProfileIpv6FilterBgpXml    `xml:"bgp,omitempty"`
	Misc        []generic.Xml                             `xml:",any"`
}
type protocolRedistProfileIpv6FilterOspfv3Xml struct {
	PathType *util.MemberType `xml:"path-type,omitempty"`
	Area     *util.MemberType `xml:"area,omitempty"`
	Tag      *util.MemberType `xml:"tag,omitempty"`
	Misc     []generic.Xml    `xml:",any"`
}
type protocolRedistProfileIpv6FilterBgpXml struct {
	Community         *util.MemberType `xml:"community,omitempty"`
	ExtendedCommunity *util.MemberType `xml:"extended-community,omitempty"`
	Misc              []generic.Xml    `xml:",any"`
}
type protocolRedistProfileIpv6ActionXml struct {
	NoRedist *protocolRedistProfileIpv6ActionNoRedistXml `xml:"no-redist,omitempty"`
	Redist   *protocolRedistProfileIpv6ActionRedistXml   `xml:"redist,omitempty"`
	Misc     []generic.Xml                               `xml:",any"`
}
type protocolRedistProfileIpv6ActionNoRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolRedistProfileIpv6ActionRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolRipXml struct {
	AllowRedistDefaultRoute *string                             `xml:"allow-redist-default-route,omitempty"`
	AuthProfile             *protocolRipAuthProfileContainerXml `xml:"auth-profile,omitempty"`
	Enable                  *string                             `xml:"enable,omitempty"`
	ExportRules             *protocolRipExportRulesContainerXml `xml:"export-rules,omitempty"`
	GlobalBfd               *protocolRipGlobalBfdXml            `xml:"global-bfd,omitempty"`
	Interface               *protocolRipInterfaceContainerXml   `xml:"interface,omitempty"`
	RejectDefaultRoute      *string                             `xml:"reject-default-route,omitempty"`
	Timers                  *protocolRipTimersXml               `xml:"timers,omitempty"`
	Misc                    []generic.Xml                       `xml:",any"`
}
type protocolRipAuthProfileContainerXml struct {
	Entries []protocolRipAuthProfileXml `xml:"entry"`
}
type protocolRipAuthProfileXml struct {
	XMLName  xml.Name                               `xml:"entry"`
	Name     string                                 `xml:"name,attr"`
	Password *string                                `xml:"password,omitempty"`
	Md5      *protocolRipAuthProfileMd5ContainerXml `xml:"md5,omitempty"`
	Misc     []generic.Xml                          `xml:",any"`
}
type protocolRipAuthProfileMd5ContainerXml struct {
	Entries []protocolRipAuthProfileMd5Xml `xml:"entry"`
}
type protocolRipAuthProfileMd5Xml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	Key       *string       `xml:"key,omitempty"`
	Preferred *string       `xml:"preferred,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type protocolRipExportRulesContainerXml struct {
	Entries []protocolRipExportRulesXml `xml:"entry"`
}
type protocolRipExportRulesXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Metric  *int64        `xml:"metric,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolRipGlobalBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolRipInterfaceContainerXml struct {
	Entries []protocolRipInterfaceXml `xml:"entry"`
}
type protocolRipInterfaceXml struct {
	XMLName        xml.Name                             `xml:"entry"`
	Name           string                               `xml:"name,attr"`
	Enable         *string                              `xml:"enable,omitempty"`
	Authentication *string                              `xml:"authentication,omitempty"`
	Mode           *string                              `xml:"mode,omitempty"`
	DefaultRoute   *protocolRipInterfaceDefaultRouteXml `xml:"default-route,omitempty"`
	Bfd            *protocolRipInterfaceBfdXml          `xml:"bfd,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
}
type protocolRipInterfaceDefaultRouteXml struct {
	Disable   *protocolRipInterfaceDefaultRouteDisableXml   `xml:"disable,omitempty"`
	Advertise *protocolRipInterfaceDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Misc      []generic.Xml                                 `xml:",any"`
}
type protocolRipInterfaceDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type protocolRipInterfaceDefaultRouteAdvertiseXml struct {
	Metric *int64        `xml:"metric,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolRipInterfaceBfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type protocolRipTimersXml struct {
	DeleteIntervals *int64        `xml:"delete-intervals,omitempty"`
	ExpireIntervals *int64        `xml:"expire-intervals,omitempty"`
	IntervalSeconds *int64        `xml:"interval-seconds,omitempty"`
	UpdateIntervals *int64        `xml:"update-intervals,omitempty"`
	Misc            []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AdminDists != nil {
		var obj adminDistsXml
		obj.MarshalFromObject(*s.AdminDists)
		o.AdminDists = &obj
	}
	if s.Ecmp != nil {
		var obj ecmpXml
		obj.MarshalFromObject(*s.Ecmp)
		o.Ecmp = &obj
	}
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.Multicast != nil {
		var obj multicastXml
		obj.MarshalFromObject(*s.Multicast)
		o.Multicast = &obj
	}
	if s.Protocol != nil {
		var obj protocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var adminDistsVal *AdminDists
	if o.AdminDists != nil {
		obj, err := o.AdminDists.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adminDistsVal = obj
	}
	var ecmpVal *Ecmp
	if o.Ecmp != nil {
		obj, err := o.Ecmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ecmpVal = obj
	}
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var multicastVal *Multicast
	if o.Multicast != nil {
		obj, err := o.Multicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastVal = obj
	}
	var protocolVal *Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &Entry{
		Name:       o.Name,
		AdminDists: adminDistsVal,
		Ecmp:       ecmpVal,
		Interface:  interfaceVal,
		Multicast:  multicastVal,
		Protocol:   protocolVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *adminDistsXml) MarshalFromObject(s AdminDists) {
	o.Ebgp = s.Ebgp
	o.Ibgp = s.Ibgp
	o.OspfExt = s.OspfExt
	o.OspfInt = s.OspfInt
	o.Ospfv3Ext = s.Ospfv3Ext
	o.Ospfv3Int = s.Ospfv3Int
	o.Rip = s.Rip
	o.Static = s.Static
	o.StaticIpv6 = s.StaticIpv6
	o.Misc = s.Misc
}

func (o adminDistsXml) UnmarshalToObject() (*AdminDists, error) {

	result := &AdminDists{
		Ebgp:       o.Ebgp,
		Ibgp:       o.Ibgp,
		OspfExt:    o.OspfExt,
		OspfInt:    o.OspfInt,
		Ospfv3Ext:  o.Ospfv3Ext,
		Ospfv3Int:  o.Ospfv3Int,
		Rip:        o.Rip,
		Static:     o.Static,
		StaticIpv6: o.StaticIpv6,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *ecmpXml) MarshalFromObject(s Ecmp) {
	if s.Algorithm != nil {
		var obj ecmpAlgorithmXml
		obj.MarshalFromObject(*s.Algorithm)
		o.Algorithm = &obj
	}
	o.Enable = util.YesNo(s.Enable, nil)
	o.MaxPath = s.MaxPath
	o.StrictSourcePath = util.YesNo(s.StrictSourcePath, nil)
	o.SymmetricReturn = util.YesNo(s.SymmetricReturn, nil)
	o.Misc = s.Misc
}

func (o ecmpXml) UnmarshalToObject() (*Ecmp, error) {
	var algorithmVal *EcmpAlgorithm
	if o.Algorithm != nil {
		obj, err := o.Algorithm.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		algorithmVal = obj
	}

	result := &Ecmp{
		Algorithm:        algorithmVal,
		Enable:           util.AsBool(o.Enable, nil),
		MaxPath:          o.MaxPath,
		StrictSourcePath: util.AsBool(o.StrictSourcePath, nil),
		SymmetricReturn:  util.AsBool(o.SymmetricReturn, nil),
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmXml) MarshalFromObject(s EcmpAlgorithm) {
	if s.BalancedRoundRobin != nil {
		var obj ecmpAlgorithmBalancedRoundRobinXml
		obj.MarshalFromObject(*s.BalancedRoundRobin)
		o.BalancedRoundRobin = &obj
	}
	if s.IpHash != nil {
		var obj ecmpAlgorithmIpHashXml
		obj.MarshalFromObject(*s.IpHash)
		o.IpHash = &obj
	}
	if s.IpModulo != nil {
		var obj ecmpAlgorithmIpModuloXml
		obj.MarshalFromObject(*s.IpModulo)
		o.IpModulo = &obj
	}
	if s.WeightedRoundRobin != nil {
		var obj ecmpAlgorithmWeightedRoundRobinXml
		obj.MarshalFromObject(*s.WeightedRoundRobin)
		o.WeightedRoundRobin = &obj
	}
	o.Misc = s.Misc
}

func (o ecmpAlgorithmXml) UnmarshalToObject() (*EcmpAlgorithm, error) {
	var balancedRoundRobinVal *EcmpAlgorithmBalancedRoundRobin
	if o.BalancedRoundRobin != nil {
		obj, err := o.BalancedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		balancedRoundRobinVal = obj
	}
	var ipHashVal *EcmpAlgorithmIpHash
	if o.IpHash != nil {
		obj, err := o.IpHash.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipHashVal = obj
	}
	var ipModuloVal *EcmpAlgorithmIpModulo
	if o.IpModulo != nil {
		obj, err := o.IpModulo.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipModuloVal = obj
	}
	var weightedRoundRobinVal *EcmpAlgorithmWeightedRoundRobin
	if o.WeightedRoundRobin != nil {
		obj, err := o.WeightedRoundRobin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weightedRoundRobinVal = obj
	}

	result := &EcmpAlgorithm{
		BalancedRoundRobin: balancedRoundRobinVal,
		IpHash:             ipHashVal,
		IpModulo:           ipModuloVal,
		WeightedRoundRobin: weightedRoundRobinVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmBalancedRoundRobinXml) MarshalFromObject(s EcmpAlgorithmBalancedRoundRobin) {
	o.Misc = s.Misc
}

func (o ecmpAlgorithmBalancedRoundRobinXml) UnmarshalToObject() (*EcmpAlgorithmBalancedRoundRobin, error) {

	result := &EcmpAlgorithmBalancedRoundRobin{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmIpHashXml) MarshalFromObject(s EcmpAlgorithmIpHash) {
	o.HashSeed = s.HashSeed
	o.SrcOnly = util.YesNo(s.SrcOnly, nil)
	o.UsePort = util.YesNo(s.UsePort, nil)
	o.Misc = s.Misc
}

func (o ecmpAlgorithmIpHashXml) UnmarshalToObject() (*EcmpAlgorithmIpHash, error) {

	result := &EcmpAlgorithmIpHash{
		HashSeed: o.HashSeed,
		SrcOnly:  util.AsBool(o.SrcOnly, nil),
		UsePort:  util.AsBool(o.UsePort, nil),
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmIpModuloXml) MarshalFromObject(s EcmpAlgorithmIpModulo) {
	o.Misc = s.Misc
}

func (o ecmpAlgorithmIpModuloXml) UnmarshalToObject() (*EcmpAlgorithmIpModulo, error) {

	result := &EcmpAlgorithmIpModulo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmWeightedRoundRobinXml) MarshalFromObject(s EcmpAlgorithmWeightedRoundRobin) {
	if s.Interface != nil {
		var objs []ecmpAlgorithmWeightedRoundRobinInterfaceXml
		for _, elt := range s.Interface {
			var obj ecmpAlgorithmWeightedRoundRobinInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &ecmpAlgorithmWeightedRoundRobinInterfaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o ecmpAlgorithmWeightedRoundRobinXml) UnmarshalToObject() (*EcmpAlgorithmWeightedRoundRobin, error) {
	var interfaceVal []EcmpAlgorithmWeightedRoundRobinInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}

	result := &EcmpAlgorithmWeightedRoundRobin{
		Interface: interfaceVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *ecmpAlgorithmWeightedRoundRobinInterfaceXml) MarshalFromObject(s EcmpAlgorithmWeightedRoundRobinInterface) {
	o.Name = s.Name
	o.Weight = s.Weight
	o.Misc = s.Misc
}

func (o ecmpAlgorithmWeightedRoundRobinInterfaceXml) UnmarshalToObject() (*EcmpAlgorithmWeightedRoundRobinInterface, error) {

	result := &EcmpAlgorithmWeightedRoundRobinInterface{
		Name:   o.Name,
		Weight: o.Weight,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *multicastXml) MarshalFromObject(s Multicast) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.InterfaceGroup != nil {
		var objs []multicastInterfaceGroupXml
		for _, elt := range s.InterfaceGroup {
			var obj multicastInterfaceGroupXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.InterfaceGroup = &multicastInterfaceGroupContainerXml{Entries: objs}
	}
	o.RouteAgeoutTime = s.RouteAgeoutTime
	if s.Rp != nil {
		var obj multicastRpXml
		obj.MarshalFromObject(*s.Rp)
		o.Rp = &obj
	}
	if s.SptThreshold != nil {
		var objs []multicastSptThresholdXml
		for _, elt := range s.SptThreshold {
			var obj multicastSptThresholdXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SptThreshold = &multicastSptThresholdContainerXml{Entries: objs}
	}
	if s.SsmAddressSpace != nil {
		var objs []multicastSsmAddressSpaceXml
		for _, elt := range s.SsmAddressSpace {
			var obj multicastSsmAddressSpaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SsmAddressSpace = &multicastSsmAddressSpaceContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o multicastXml) UnmarshalToObject() (*Multicast, error) {
	var interfaceGroupVal []MulticastInterfaceGroup
	if o.InterfaceGroup != nil {
		for _, elt := range o.InterfaceGroup.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceGroupVal = append(interfaceGroupVal, *obj)
		}
	}
	var rpVal *MulticastRp
	if o.Rp != nil {
		obj, err := o.Rp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		rpVal = obj
	}
	var sptThresholdVal []MulticastSptThreshold
	if o.SptThreshold != nil {
		for _, elt := range o.SptThreshold.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			sptThresholdVal = append(sptThresholdVal, *obj)
		}
	}
	var ssmAddressSpaceVal []MulticastSsmAddressSpace
	if o.SsmAddressSpace != nil {
		for _, elt := range o.SsmAddressSpace.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ssmAddressSpaceVal = append(ssmAddressSpaceVal, *obj)
		}
	}

	result := &Multicast{
		Enable:          util.AsBool(o.Enable, nil),
		InterfaceGroup:  interfaceGroupVal,
		RouteAgeoutTime: o.RouteAgeoutTime,
		Rp:              rpVal,
		SptThreshold:    sptThresholdVal,
		SsmAddressSpace: ssmAddressSpaceVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupXml) MarshalFromObject(s MulticastInterfaceGroup) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.GroupPermission != nil {
		var obj multicastInterfaceGroupGroupPermissionXml
		obj.MarshalFromObject(*s.GroupPermission)
		o.GroupPermission = &obj
	}
	if s.Igmp != nil {
		var obj multicastInterfaceGroupIgmpXml
		obj.MarshalFromObject(*s.Igmp)
		o.Igmp = &obj
	}
	if s.Pim != nil {
		var obj multicastInterfaceGroupPimXml
		obj.MarshalFromObject(*s.Pim)
		o.Pim = &obj
	}
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupXml) UnmarshalToObject() (*MulticastInterfaceGroup, error) {
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var groupPermissionVal *MulticastInterfaceGroupGroupPermission
	if o.GroupPermission != nil {
		obj, err := o.GroupPermission.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		groupPermissionVal = obj
	}
	var igmpVal *MulticastInterfaceGroupIgmp
	if o.Igmp != nil {
		obj, err := o.Igmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		igmpVal = obj
	}
	var pimVal *MulticastInterfaceGroupPim
	if o.Pim != nil {
		obj, err := o.Pim.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pimVal = obj
	}

	result := &MulticastInterfaceGroup{
		Name:            o.Name,
		Description:     o.Description,
		Interface:       interfaceVal,
		GroupPermission: groupPermissionVal,
		Igmp:            igmpVal,
		Pim:             pimVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupGroupPermissionXml) MarshalFromObject(s MulticastInterfaceGroupGroupPermission) {
	if s.AnySourceMulticast != nil {
		var objs []multicastInterfaceGroupGroupPermissionAnySourceMulticastXml
		for _, elt := range s.AnySourceMulticast {
			var obj multicastInterfaceGroupGroupPermissionAnySourceMulticastXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AnySourceMulticast = &multicastInterfaceGroupGroupPermissionAnySourceMulticastContainerXml{Entries: objs}
	}
	if s.SourceSpecificMulticast != nil {
		var objs []multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml
		for _, elt := range s.SourceSpecificMulticast {
			var obj multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SourceSpecificMulticast = &multicastInterfaceGroupGroupPermissionSourceSpecificMulticastContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupGroupPermissionXml) UnmarshalToObject() (*MulticastInterfaceGroupGroupPermission, error) {
	var anySourceMulticastVal []MulticastInterfaceGroupGroupPermissionAnySourceMulticast
	if o.AnySourceMulticast != nil {
		for _, elt := range o.AnySourceMulticast.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			anySourceMulticastVal = append(anySourceMulticastVal, *obj)
		}
	}
	var sourceSpecificMulticastVal []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast
	if o.SourceSpecificMulticast != nil {
		for _, elt := range o.SourceSpecificMulticast.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			sourceSpecificMulticastVal = append(sourceSpecificMulticastVal, *obj)
		}
	}

	result := &MulticastInterfaceGroupGroupPermission{
		AnySourceMulticast:      anySourceMulticastVal,
		SourceSpecificMulticast: sourceSpecificMulticastVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupGroupPermissionAnySourceMulticastXml) MarshalFromObject(s MulticastInterfaceGroupGroupPermissionAnySourceMulticast) {
	o.Name = s.Name
	o.GroupAddress = s.GroupAddress
	o.Included = util.YesNo(s.Included, nil)
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupGroupPermissionAnySourceMulticastXml) UnmarshalToObject() (*MulticastInterfaceGroupGroupPermissionAnySourceMulticast, error) {

	result := &MulticastInterfaceGroupGroupPermissionAnySourceMulticast{
		Name:         o.Name,
		GroupAddress: o.GroupAddress,
		Included:     util.AsBool(o.Included, nil),
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml) MarshalFromObject(s MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast) {
	o.Name = s.Name
	o.GroupAddress = s.GroupAddress
	o.SourceAddress = s.SourceAddress
	o.Included = util.YesNo(s.Included, nil)
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupGroupPermissionSourceSpecificMulticastXml) UnmarshalToObject() (*MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast, error) {

	result := &MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast{
		Name:          o.Name,
		GroupAddress:  o.GroupAddress,
		SourceAddress: o.SourceAddress,
		Included:      util.AsBool(o.Included, nil),
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupIgmpXml) MarshalFromObject(s MulticastInterfaceGroupIgmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Version = s.Version
	o.MaxQueryResponseTime = s.MaxQueryResponseTime
	o.QueryInterval = s.QueryInterval
	o.LastMemberQueryInterval = s.LastMemberQueryInterval
	o.ImmediateLeave = util.YesNo(s.ImmediateLeave, nil)
	o.Robustness = s.Robustness
	o.MaxGroups = s.MaxGroups
	o.MaxSources = s.MaxSources
	o.RouterAlertPolicing = util.YesNo(s.RouterAlertPolicing, nil)
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupIgmpXml) UnmarshalToObject() (*MulticastInterfaceGroupIgmp, error) {

	result := &MulticastInterfaceGroupIgmp{
		Enable:                  util.AsBool(o.Enable, nil),
		Version:                 o.Version,
		MaxQueryResponseTime:    o.MaxQueryResponseTime,
		QueryInterval:           o.QueryInterval,
		LastMemberQueryInterval: o.LastMemberQueryInterval,
		ImmediateLeave:          util.AsBool(o.ImmediateLeave, nil),
		Robustness:              o.Robustness,
		MaxGroups:               o.MaxGroups,
		MaxSources:              o.MaxSources,
		RouterAlertPolicing:     util.AsBool(o.RouterAlertPolicing, nil),
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupPimXml) MarshalFromObject(s MulticastInterfaceGroupPim) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.AssertInterval = s.AssertInterval
	o.HelloInterval = s.HelloInterval
	o.JoinPruneInterval = s.JoinPruneInterval
	o.DrPriority = s.DrPriority
	o.BsrBorder = util.YesNo(s.BsrBorder, nil)
	if s.AllowedNeighbors != nil {
		var objs []multicastInterfaceGroupPimAllowedNeighborsXml
		for _, elt := range s.AllowedNeighbors {
			var obj multicastInterfaceGroupPimAllowedNeighborsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AllowedNeighbors = &multicastInterfaceGroupPimAllowedNeighborsContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupPimXml) UnmarshalToObject() (*MulticastInterfaceGroupPim, error) {
	var allowedNeighborsVal []MulticastInterfaceGroupPimAllowedNeighbors
	if o.AllowedNeighbors != nil {
		for _, elt := range o.AllowedNeighbors.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			allowedNeighborsVal = append(allowedNeighborsVal, *obj)
		}
	}

	result := &MulticastInterfaceGroupPim{
		Enable:            util.AsBool(o.Enable, nil),
		AssertInterval:    o.AssertInterval,
		HelloInterval:     o.HelloInterval,
		JoinPruneInterval: o.JoinPruneInterval,
		DrPriority:        o.DrPriority,
		BsrBorder:         util.AsBool(o.BsrBorder, nil),
		AllowedNeighbors:  allowedNeighborsVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *multicastInterfaceGroupPimAllowedNeighborsXml) MarshalFromObject(s MulticastInterfaceGroupPimAllowedNeighbors) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o multicastInterfaceGroupPimAllowedNeighborsXml) UnmarshalToObject() (*MulticastInterfaceGroupPimAllowedNeighbors, error) {

	result := &MulticastInterfaceGroupPimAllowedNeighbors{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *multicastRpXml) MarshalFromObject(s MulticastRp) {
	if s.ExternalRp != nil {
		var objs []multicastRpExternalRpXml
		for _, elt := range s.ExternalRp {
			var obj multicastRpExternalRpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExternalRp = &multicastRpExternalRpContainerXml{Entries: objs}
	}
	if s.LocalRp != nil {
		var obj multicastRpLocalRpXml
		obj.MarshalFromObject(*s.LocalRp)
		o.LocalRp = &obj
	}
	o.Misc = s.Misc
}

func (o multicastRpXml) UnmarshalToObject() (*MulticastRp, error) {
	var externalRpVal []MulticastRpExternalRp
	if o.ExternalRp != nil {
		for _, elt := range o.ExternalRp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			externalRpVal = append(externalRpVal, *obj)
		}
	}
	var localRpVal *MulticastRpLocalRp
	if o.LocalRp != nil {
		obj, err := o.LocalRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localRpVal = obj
	}

	result := &MulticastRp{
		ExternalRp: externalRpVal,
		LocalRp:    localRpVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *multicastRpExternalRpXml) MarshalFromObject(s MulticastRpExternalRp) {
	o.Name = s.Name
	if s.GroupAddresses != nil {
		o.GroupAddresses = util.StrToMem(s.GroupAddresses)
	}
	o.Override = util.YesNo(s.Override, nil)
	o.Misc = s.Misc
}

func (o multicastRpExternalRpXml) UnmarshalToObject() (*MulticastRpExternalRp, error) {
	var groupAddressesVal []string
	if o.GroupAddresses != nil {
		groupAddressesVal = util.MemToStr(o.GroupAddresses)
	}

	result := &MulticastRpExternalRp{
		Name:           o.Name,
		GroupAddresses: groupAddressesVal,
		Override:       util.AsBool(o.Override, nil),
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *multicastRpLocalRpXml) MarshalFromObject(s MulticastRpLocalRp) {
	if s.CandidateRp != nil {
		var obj multicastRpLocalRpCandidateRpXml
		obj.MarshalFromObject(*s.CandidateRp)
		o.CandidateRp = &obj
	}
	if s.StaticRp != nil {
		var obj multicastRpLocalRpStaticRpXml
		obj.MarshalFromObject(*s.StaticRp)
		o.StaticRp = &obj
	}
	o.Misc = s.Misc
}

func (o multicastRpLocalRpXml) UnmarshalToObject() (*MulticastRpLocalRp, error) {
	var candidateRpVal *MulticastRpLocalRpCandidateRp
	if o.CandidateRp != nil {
		obj, err := o.CandidateRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		candidateRpVal = obj
	}
	var staticRpVal *MulticastRpLocalRpStaticRp
	if o.StaticRp != nil {
		obj, err := o.StaticRp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticRpVal = obj
	}

	result := &MulticastRpLocalRp{
		CandidateRp: candidateRpVal,
		StaticRp:    staticRpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *multicastRpLocalRpCandidateRpXml) MarshalFromObject(s MulticastRpLocalRpCandidateRp) {
	o.Address = s.Address
	o.AdvertisementInterval = s.AdvertisementInterval
	if s.GroupAddresses != nil {
		o.GroupAddresses = util.StrToMem(s.GroupAddresses)
	}
	o.Interface = s.Interface
	o.Priority = s.Priority
	o.Misc = s.Misc
}

func (o multicastRpLocalRpCandidateRpXml) UnmarshalToObject() (*MulticastRpLocalRpCandidateRp, error) {
	var groupAddressesVal []string
	if o.GroupAddresses != nil {
		groupAddressesVal = util.MemToStr(o.GroupAddresses)
	}

	result := &MulticastRpLocalRpCandidateRp{
		Address:               o.Address,
		AdvertisementInterval: o.AdvertisementInterval,
		GroupAddresses:        groupAddressesVal,
		Interface:             o.Interface,
		Priority:              o.Priority,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *multicastRpLocalRpStaticRpXml) MarshalFromObject(s MulticastRpLocalRpStaticRp) {
	o.Address = s.Address
	if s.GroupAddresses != nil {
		o.GroupAddresses = util.StrToMem(s.GroupAddresses)
	}
	o.Interface = s.Interface
	o.Override = util.YesNo(s.Override, nil)
	o.Misc = s.Misc
}

func (o multicastRpLocalRpStaticRpXml) UnmarshalToObject() (*MulticastRpLocalRpStaticRp, error) {
	var groupAddressesVal []string
	if o.GroupAddresses != nil {
		groupAddressesVal = util.MemToStr(o.GroupAddresses)
	}

	result := &MulticastRpLocalRpStaticRp{
		Address:        o.Address,
		GroupAddresses: groupAddressesVal,
		Interface:      o.Interface,
		Override:       util.AsBool(o.Override, nil),
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *multicastSptThresholdXml) MarshalFromObject(s MulticastSptThreshold) {
	o.Name = s.Name
	o.Threshold = s.Threshold
	o.Misc = s.Misc
}

func (o multicastSptThresholdXml) UnmarshalToObject() (*MulticastSptThreshold, error) {

	result := &MulticastSptThreshold{
		Name:      o.Name,
		Threshold: o.Threshold,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *multicastSsmAddressSpaceXml) MarshalFromObject(s MulticastSsmAddressSpace) {
	o.Name = s.Name
	o.GroupAddress = s.GroupAddress
	o.Included = util.YesNo(s.Included, nil)
	o.Misc = s.Misc
}

func (o multicastSsmAddressSpaceXml) UnmarshalToObject() (*MulticastSsmAddressSpace, error) {

	result := &MulticastSsmAddressSpace{
		Name:         o.Name,
		GroupAddress: o.GroupAddress,
		Included:     util.AsBool(o.Included, nil),
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *protocolXml) MarshalFromObject(s Protocol) {
	if s.Bgp != nil {
		var obj protocolBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	if s.Ospf != nil {
		var obj protocolOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Ospfv3 != nil {
		var obj protocolOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.RedistProfile != nil {
		var objs []protocolRedistProfileXml
		for _, elt := range s.RedistProfile {
			var obj protocolRedistProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RedistProfile = &protocolRedistProfileContainerXml{Entries: objs}
	}
	if s.RedistProfileIpv6 != nil {
		var objs []protocolRedistProfileIpv6Xml
		for _, elt := range s.RedistProfileIpv6 {
			var obj protocolRedistProfileIpv6Xml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RedistProfileIpv6 = &protocolRedistProfileIpv6ContainerXml{Entries: objs}
	}
	if s.Rip != nil {
		var obj protocolRipXml
		obj.MarshalFromObject(*s.Rip)
		o.Rip = &obj
	}
	o.Misc = s.Misc
}

func (o protocolXml) UnmarshalToObject() (*Protocol, error) {
	var bgpVal *ProtocolBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}
	var ospfVal *ProtocolOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var ospfv3Val *ProtocolOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var redistProfileVal []ProtocolRedistProfile
	if o.RedistProfile != nil {
		for _, elt := range o.RedistProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			redistProfileVal = append(redistProfileVal, *obj)
		}
	}
	var redistProfileIpv6Val []ProtocolRedistProfileIpv6
	if o.RedistProfileIpv6 != nil {
		for _, elt := range o.RedistProfileIpv6.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			redistProfileIpv6Val = append(redistProfileIpv6Val, *obj)
		}
	}
	var ripVal *ProtocolRip
	if o.Rip != nil {
		obj, err := o.Rip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ripVal = obj
	}

	result := &Protocol{
		Bgp:               bgpVal,
		Ospf:              ospfVal,
		Ospfv3:            ospfv3Val,
		RedistProfile:     redistProfileVal,
		RedistProfileIpv6: redistProfileIpv6Val,
		Rip:               ripVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpXml) MarshalFromObject(s ProtocolBgp) {
	o.AllowRedistDefaultRoute = util.YesNo(s.AllowRedistDefaultRoute, nil)
	if s.AuthProfile != nil {
		var objs []protocolBgpAuthProfileXml
		for _, elt := range s.AuthProfile {
			var obj protocolBgpAuthProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AuthProfile = &protocolBgpAuthProfileContainerXml{Entries: objs}
	}
	if s.DampeningProfile != nil {
		var objs []protocolBgpDampeningProfileXml
		for _, elt := range s.DampeningProfile {
			var obj protocolBgpDampeningProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DampeningProfile = &protocolBgpDampeningProfileContainerXml{Entries: objs}
	}
	o.EcmpMultiAs = util.YesNo(s.EcmpMultiAs, nil)
	o.Enable = util.YesNo(s.Enable, nil)
	o.EnforceFirstAs = util.YesNo(s.EnforceFirstAs, nil)
	if s.GlobalBfd != nil {
		var obj protocolBgpGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	o.InstallRoute = util.YesNo(s.InstallRoute, nil)
	o.LocalAs = s.LocalAs
	if s.PeerGroup != nil {
		var objs []protocolBgpPeerGroupXml
		for _, elt := range s.PeerGroup {
			var obj protocolBgpPeerGroupXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.PeerGroup = &protocolBgpPeerGroupContainerXml{Entries: objs}
	}
	if s.Policy != nil {
		var obj protocolBgpPolicyXml
		obj.MarshalFromObject(*s.Policy)
		o.Policy = &obj
	}
	if s.RedistRules != nil {
		var objs []protocolBgpRedistRulesXml
		for _, elt := range s.RedistRules {
			var obj protocolBgpRedistRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RedistRules = &protocolBgpRedistRulesContainerXml{Entries: objs}
	}
	o.RejectDefaultRoute = util.YesNo(s.RejectDefaultRoute, nil)
	o.RouterId = s.RouterId
	if s.RoutingOptions != nil {
		var obj protocolBgpRoutingOptionsXml
		obj.MarshalFromObject(*s.RoutingOptions)
		o.RoutingOptions = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpXml) UnmarshalToObject() (*ProtocolBgp, error) {
	var authProfileVal []ProtocolBgpAuthProfile
	if o.AuthProfile != nil {
		for _, elt := range o.AuthProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			authProfileVal = append(authProfileVal, *obj)
		}
	}
	var dampeningProfileVal []ProtocolBgpDampeningProfile
	if o.DampeningProfile != nil {
		for _, elt := range o.DampeningProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			dampeningProfileVal = append(dampeningProfileVal, *obj)
		}
	}
	var globalBfdVal *ProtocolBgpGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var peerGroupVal []ProtocolBgpPeerGroup
	if o.PeerGroup != nil {
		for _, elt := range o.PeerGroup.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerGroupVal = append(peerGroupVal, *obj)
		}
	}
	var policyVal *ProtocolBgpPolicy
	if o.Policy != nil {
		obj, err := o.Policy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policyVal = obj
	}
	var redistRulesVal []ProtocolBgpRedistRules
	if o.RedistRules != nil {
		for _, elt := range o.RedistRules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			redistRulesVal = append(redistRulesVal, *obj)
		}
	}
	var routingOptionsVal *ProtocolBgpRoutingOptions
	if o.RoutingOptions != nil {
		obj, err := o.RoutingOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingOptionsVal = obj
	}

	result := &ProtocolBgp{
		AllowRedistDefaultRoute: util.AsBool(o.AllowRedistDefaultRoute, nil),
		AuthProfile:             authProfileVal,
		DampeningProfile:        dampeningProfileVal,
		EcmpMultiAs:             util.AsBool(o.EcmpMultiAs, nil),
		Enable:                  util.AsBool(o.Enable, nil),
		EnforceFirstAs:          util.AsBool(o.EnforceFirstAs, nil),
		GlobalBfd:               globalBfdVal,
		InstallRoute:            util.AsBool(o.InstallRoute, nil),
		LocalAs:                 o.LocalAs,
		PeerGroup:               peerGroupVal,
		Policy:                  policyVal,
		RedistRules:             redistRulesVal,
		RejectDefaultRoute:      util.AsBool(o.RejectDefaultRoute, nil),
		RouterId:                o.RouterId,
		RoutingOptions:          routingOptionsVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpAuthProfileXml) MarshalFromObject(s ProtocolBgpAuthProfile) {
	o.Name = s.Name
	o.Secret = s.Secret
	o.Misc = s.Misc
}

func (o protocolBgpAuthProfileXml) UnmarshalToObject() (*ProtocolBgpAuthProfile, error) {

	result := &ProtocolBgpAuthProfile{
		Name:   o.Name,
		Secret: o.Secret,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpDampeningProfileXml) MarshalFromObject(s ProtocolBgpDampeningProfile) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Cutoff = s.Cutoff
	o.Reuse = s.Reuse
	o.MaxHoldTime = s.MaxHoldTime
	o.DecayHalfLifeReachable = s.DecayHalfLifeReachable
	o.DecayHalfLifeUnreachable = s.DecayHalfLifeUnreachable
	o.Misc = s.Misc
}

func (o protocolBgpDampeningProfileXml) UnmarshalToObject() (*ProtocolBgpDampeningProfile, error) {

	result := &ProtocolBgpDampeningProfile{
		Name:                     o.Name,
		Enable:                   util.AsBool(o.Enable, nil),
		Cutoff:                   o.Cutoff,
		Reuse:                    o.Reuse,
		MaxHoldTime:              o.MaxHoldTime,
		DecayHalfLifeReachable:   o.DecayHalfLifeReachable,
		DecayHalfLifeUnreachable: o.DecayHalfLifeUnreachable,
		Misc:                     o.Misc,
	}
	return result, nil
}
func (o *protocolBgpGlobalBfdXml) MarshalFromObject(s ProtocolBgpGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolBgpGlobalBfdXml) UnmarshalToObject() (*ProtocolBgpGlobalBfd, error) {

	result := &ProtocolBgpGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupXml) MarshalFromObject(s ProtocolBgpPeerGroup) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.AggregatedConfedAsPath = util.YesNo(s.AggregatedConfedAsPath, nil)
	o.SoftResetWithStoredInfo = util.YesNo(s.SoftResetWithStoredInfo, nil)
	if s.Type != nil {
		var obj protocolBgpPeerGroupTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Peer != nil {
		var objs []protocolBgpPeerGroupPeerXml
		for _, elt := range s.Peer {
			var obj protocolBgpPeerGroupPeerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Peer = &protocolBgpPeerGroupPeerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupXml) UnmarshalToObject() (*ProtocolBgpPeerGroup, error) {
	var typeVal *ProtocolBgpPeerGroupType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var peerVal []ProtocolBgpPeerGroupPeer
	if o.Peer != nil {
		for _, elt := range o.Peer.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			peerVal = append(peerVal, *obj)
		}
	}

	result := &ProtocolBgpPeerGroup{
		Name:                    o.Name,
		Enable:                  util.AsBool(o.Enable, nil),
		AggregatedConfedAsPath:  util.AsBool(o.AggregatedConfedAsPath, nil),
		SoftResetWithStoredInfo: util.AsBool(o.SoftResetWithStoredInfo, nil),
		Type:                    typeVal,
		Peer:                    peerVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupTypeXml) MarshalFromObject(s ProtocolBgpPeerGroupType) {
	if s.Ibgp != nil {
		var obj protocolBgpPeerGroupTypeIbgpXml
		obj.MarshalFromObject(*s.Ibgp)
		o.Ibgp = &obj
	}
	if s.EbgpConfed != nil {
		var obj protocolBgpPeerGroupTypeEbgpConfedXml
		obj.MarshalFromObject(*s.EbgpConfed)
		o.EbgpConfed = &obj
	}
	if s.IbgpConfed != nil {
		var obj protocolBgpPeerGroupTypeIbgpConfedXml
		obj.MarshalFromObject(*s.IbgpConfed)
		o.IbgpConfed = &obj
	}
	if s.Ebgp != nil {
		var obj protocolBgpPeerGroupTypeEbgpXml
		obj.MarshalFromObject(*s.Ebgp)
		o.Ebgp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupTypeXml) UnmarshalToObject() (*ProtocolBgpPeerGroupType, error) {
	var ibgpVal *ProtocolBgpPeerGroupTypeIbgp
	if o.Ibgp != nil {
		obj, err := o.Ibgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ibgpVal = obj
	}
	var ebgpConfedVal *ProtocolBgpPeerGroupTypeEbgpConfed
	if o.EbgpConfed != nil {
		obj, err := o.EbgpConfed.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ebgpConfedVal = obj
	}
	var ibgpConfedVal *ProtocolBgpPeerGroupTypeIbgpConfed
	if o.IbgpConfed != nil {
		obj, err := o.IbgpConfed.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ibgpConfedVal = obj
	}
	var ebgpVal *ProtocolBgpPeerGroupTypeEbgp
	if o.Ebgp != nil {
		obj, err := o.Ebgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ebgpVal = obj
	}

	result := &ProtocolBgpPeerGroupType{
		Ibgp:       ibgpVal,
		EbgpConfed: ebgpConfedVal,
		IbgpConfed: ibgpConfedVal,
		Ebgp:       ebgpVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupTypeIbgpXml) MarshalFromObject(s ProtocolBgpPeerGroupTypeIbgp) {
	o.ExportNexthop = s.ExportNexthop
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupTypeIbgpXml) UnmarshalToObject() (*ProtocolBgpPeerGroupTypeIbgp, error) {

	result := &ProtocolBgpPeerGroupTypeIbgp{
		ExportNexthop: o.ExportNexthop,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupTypeEbgpConfedXml) MarshalFromObject(s ProtocolBgpPeerGroupTypeEbgpConfed) {
	o.ExportNexthop = s.ExportNexthop
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupTypeEbgpConfedXml) UnmarshalToObject() (*ProtocolBgpPeerGroupTypeEbgpConfed, error) {

	result := &ProtocolBgpPeerGroupTypeEbgpConfed{
		ExportNexthop: o.ExportNexthop,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupTypeIbgpConfedXml) MarshalFromObject(s ProtocolBgpPeerGroupTypeIbgpConfed) {
	o.ExportNexthop = s.ExportNexthop
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupTypeIbgpConfedXml) UnmarshalToObject() (*ProtocolBgpPeerGroupTypeIbgpConfed, error) {

	result := &ProtocolBgpPeerGroupTypeIbgpConfed{
		ExportNexthop: o.ExportNexthop,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupTypeEbgpXml) MarshalFromObject(s ProtocolBgpPeerGroupTypeEbgp) {
	o.ImportNexthop = s.ImportNexthop
	o.ExportNexthop = s.ExportNexthop
	o.RemovePrivateAs = util.YesNo(s.RemovePrivateAs, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupTypeEbgpXml) UnmarshalToObject() (*ProtocolBgpPeerGroupTypeEbgp, error) {

	result := &ProtocolBgpPeerGroupTypeEbgp{
		ImportNexthop:   o.ImportNexthop,
		ExportNexthop:   o.ExportNexthop,
		RemovePrivateAs: util.AsBool(o.RemovePrivateAs, nil),
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerXml) MarshalFromObject(s ProtocolBgpPeerGroupPeer) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.PeerAs = s.PeerAs
	o.EnableMpBgp = util.YesNo(s.EnableMpBgp, nil)
	o.AddressFamilyIdentifier = s.AddressFamilyIdentifier
	o.EnableSenderSideLoopDetection = util.YesNo(s.EnableSenderSideLoopDetection, nil)
	o.ReflectorClient = s.ReflectorClient
	o.PeeringType = s.PeeringType
	o.MaxPrefixes = s.MaxPrefixes
	if s.SubsequentAddressFamilyIdentifier != nil {
		var obj protocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml
		obj.MarshalFromObject(*s.SubsequentAddressFamilyIdentifier)
		o.SubsequentAddressFamilyIdentifier = &obj
	}
	if s.LocalAddress != nil {
		var obj protocolBgpPeerGroupPeerLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj protocolBgpPeerGroupPeerPeerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	if s.ConnectionOptions != nil {
		var obj protocolBgpPeerGroupPeerConnectionOptionsXml
		obj.MarshalFromObject(*s.ConnectionOptions)
		o.ConnectionOptions = &obj
	}
	if s.Bfd != nil {
		var obj protocolBgpPeerGroupPeerBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeer, error) {
	var subsequentAddressFamilyIdentifierVal *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier
	if o.SubsequentAddressFamilyIdentifier != nil {
		obj, err := o.SubsequentAddressFamilyIdentifier.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		subsequentAddressFamilyIdentifierVal = obj
	}
	var localAddressVal *ProtocolBgpPeerGroupPeerLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *ProtocolBgpPeerGroupPeerPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var connectionOptionsVal *ProtocolBgpPeerGroupPeerConnectionOptions
	if o.ConnectionOptions != nil {
		obj, err := o.ConnectionOptions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		connectionOptionsVal = obj
	}
	var bfdVal *ProtocolBgpPeerGroupPeerBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolBgpPeerGroupPeer{
		Name:                              o.Name,
		Enable:                            util.AsBool(o.Enable, nil),
		PeerAs:                            o.PeerAs,
		EnableMpBgp:                       util.AsBool(o.EnableMpBgp, nil),
		AddressFamilyIdentifier:           o.AddressFamilyIdentifier,
		EnableSenderSideLoopDetection:     util.AsBool(o.EnableSenderSideLoopDetection, nil),
		ReflectorClient:                   o.ReflectorClient,
		PeeringType:                       o.PeeringType,
		MaxPrefixes:                       o.MaxPrefixes,
		SubsequentAddressFamilyIdentifier: subsequentAddressFamilyIdentifierVal,
		LocalAddress:                      localAddressVal,
		PeerAddress:                       peerAddressVal,
		ConnectionOptions:                 connectionOptionsVal,
		Bfd:                               bfdVal,
		Misc:                              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier) {
	o.Unicast = util.YesNo(s.Unicast, nil)
	o.Multicast = util.YesNo(s.Multicast, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier, error) {

	result := &ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier{
		Unicast:   util.AsBool(o.Unicast, nil),
		Multicast: util.AsBool(o.Multicast, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerLocalAddressXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerLocalAddress) {
	o.Interface = s.Interface
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerLocalAddressXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerLocalAddress, error) {

	result := &ProtocolBgpPeerGroupPeerLocalAddress{
		Interface: o.Interface,
		Ip:        o.Ip,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerPeerAddressXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerPeerAddress) {
	o.Ip = s.Ip
	o.Fqdn = s.Fqdn
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerPeerAddressXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerPeerAddress, error) {

	result := &ProtocolBgpPeerGroupPeerPeerAddress{
		Ip:   o.Ip,
		Fqdn: o.Fqdn,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerConnectionOptionsXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerConnectionOptions) {
	o.Authentication = s.Authentication
	o.KeepAliveInterval = s.KeepAliveInterval
	o.MinRouteAdvInterval = s.MinRouteAdvInterval
	o.Multihop = s.Multihop
	o.OpenDelayTime = s.OpenDelayTime
	o.HoldTime = s.HoldTime
	o.IdleHoldTime = s.IdleHoldTime
	if s.IncomingBgpConnection != nil {
		var obj protocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml
		obj.MarshalFromObject(*s.IncomingBgpConnection)
		o.IncomingBgpConnection = &obj
	}
	if s.OutgoingBgpConnection != nil {
		var obj protocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml
		obj.MarshalFromObject(*s.OutgoingBgpConnection)
		o.OutgoingBgpConnection = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerConnectionOptionsXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerConnectionOptions, error) {
	var incomingBgpConnectionVal *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection
	if o.IncomingBgpConnection != nil {
		obj, err := o.IncomingBgpConnection.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		incomingBgpConnectionVal = obj
	}
	var outgoingBgpConnectionVal *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection
	if o.OutgoingBgpConnection != nil {
		obj, err := o.OutgoingBgpConnection.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		outgoingBgpConnectionVal = obj
	}

	result := &ProtocolBgpPeerGroupPeerConnectionOptions{
		Authentication:        o.Authentication,
		KeepAliveInterval:     o.KeepAliveInterval,
		MinRouteAdvInterval:   o.MinRouteAdvInterval,
		Multihop:              o.Multihop,
		OpenDelayTime:         o.OpenDelayTime,
		HoldTime:              o.HoldTime,
		IdleHoldTime:          o.IdleHoldTime,
		IncomingBgpConnection: incomingBgpConnectionVal,
		OutgoingBgpConnection: outgoingBgpConnectionVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection) {
	o.RemotePort = s.RemotePort
	o.Allow = util.YesNo(s.Allow, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection, error) {

	result := &ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection{
		RemotePort: o.RemotePort,
		Allow:      util.AsBool(o.Allow, nil),
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection) {
	o.LocalPort = s.LocalPort
	o.Allow = util.YesNo(s.Allow, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection, error) {

	result := &ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection{
		LocalPort: o.LocalPort,
		Allow:     util.AsBool(o.Allow, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPeerGroupPeerBfdXml) MarshalFromObject(s ProtocolBgpPeerGroupPeerBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolBgpPeerGroupPeerBfdXml) UnmarshalToObject() (*ProtocolBgpPeerGroupPeerBfd, error) {

	result := &ProtocolBgpPeerGroupPeerBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyXml) MarshalFromObject(s ProtocolBgpPolicy) {
	if s.Aggregation != nil {
		var obj protocolBgpPolicyAggregationXml
		obj.MarshalFromObject(*s.Aggregation)
		o.Aggregation = &obj
	}
	if s.ConditionalAdvertisement != nil {
		var obj protocolBgpPolicyConditionalAdvertisementXml
		obj.MarshalFromObject(*s.ConditionalAdvertisement)
		o.ConditionalAdvertisement = &obj
	}
	if s.Export != nil {
		var obj protocolBgpPolicyExportXml
		obj.MarshalFromObject(*s.Export)
		o.Export = &obj
	}
	if s.Import != nil {
		var obj protocolBgpPolicyImportXml
		obj.MarshalFromObject(*s.Import)
		o.Import = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyXml) UnmarshalToObject() (*ProtocolBgpPolicy, error) {
	var aggregationVal *ProtocolBgpPolicyAggregation
	if o.Aggregation != nil {
		obj, err := o.Aggregation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregationVal = obj
	}
	var conditionalAdvertisementVal *ProtocolBgpPolicyConditionalAdvertisement
	if o.ConditionalAdvertisement != nil {
		obj, err := o.ConditionalAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		conditionalAdvertisementVal = obj
	}
	var exportVal *ProtocolBgpPolicyExport
	if o.Export != nil {
		obj, err := o.Export.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		exportVal = obj
	}
	var importVal *ProtocolBgpPolicyImport
	if o.Import != nil {
		obj, err := o.Import.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		importVal = obj
	}

	result := &ProtocolBgpPolicy{
		Aggregation:              aggregationVal,
		ConditionalAdvertisement: conditionalAdvertisementVal,
		Export:                   exportVal,
		Import:                   importVal,
		Misc:                     o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationXml) MarshalFromObject(s ProtocolBgpPolicyAggregation) {
	if s.Address != nil {
		var objs []protocolBgpPolicyAggregationAddressXml
		for _, elt := range s.Address {
			var obj protocolBgpPolicyAggregationAddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &protocolBgpPolicyAggregationAddressContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregation, error) {
	var addressVal []ProtocolBgpPolicyAggregationAddress
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyAggregation{
		Address: addressVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddress) {
	o.Name = s.Name
	o.Prefix = s.Prefix
	o.Enable = util.YesNo(s.Enable, nil)
	o.Summary = util.YesNo(s.Summary, nil)
	o.AsSet = util.YesNo(s.AsSet, nil)
	if s.AggregateRouteAttributes != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesXml
		obj.MarshalFromObject(*s.AggregateRouteAttributes)
		o.AggregateRouteAttributes = &obj
	}
	if s.SuppressFilters != nil {
		var objs []protocolBgpPolicyAggregationAddressSuppressFiltersXml
		for _, elt := range s.SuppressFilters {
			var obj protocolBgpPolicyAggregationAddressSuppressFiltersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.SuppressFilters = &protocolBgpPolicyAggregationAddressSuppressFiltersContainerXml{Entries: objs}
	}
	if s.AdvertiseFilters != nil {
		var objs []protocolBgpPolicyAggregationAddressAdvertiseFiltersXml
		for _, elt := range s.AdvertiseFilters {
			var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AdvertiseFilters = &protocolBgpPolicyAggregationAddressAdvertiseFiltersContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddress, error) {
	var aggregateRouteAttributesVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes
	if o.AggregateRouteAttributes != nil {
		obj, err := o.AggregateRouteAttributes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregateRouteAttributesVal = obj
	}
	var suppressFiltersVal []ProtocolBgpPolicyAggregationAddressSuppressFilters
	if o.SuppressFilters != nil {
		for _, elt := range o.SuppressFilters.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suppressFiltersVal = append(suppressFiltersVal, *obj)
		}
	}
	var advertiseFiltersVal []ProtocolBgpPolicyAggregationAddressAdvertiseFilters
	if o.AdvertiseFilters != nil {
		for _, elt := range o.AdvertiseFilters.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			advertiseFiltersVal = append(advertiseFiltersVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyAggregationAddress{
		Name:                     o.Name,
		Prefix:                   o.Prefix,
		Enable:                   util.AsBool(o.Enable, nil),
		Summary:                  util.AsBool(o.Summary, nil),
		AsSet:                    util.AsBool(o.AsSet, nil),
		AggregateRouteAttributes: aggregateRouteAttributesVal,
		SuppressFilters:          suppressFiltersVal,
		AdvertiseFilters:         advertiseFiltersVal,
		Misc:                     o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes) {
	o.LocalPreference = s.LocalPreference
	o.Med = s.Med
	o.Weight = s.Weight
	o.Nexthop = s.Nexthop
	o.Origin = s.Origin
	o.AsPathLimit = s.AsPathLimit
	if s.AsPath != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes, error) {
	var asPathVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes{
		LocalPreference:   o.LocalPreference,
		Med:               o.Med,
		Weight:            o.Weight,
		Nexthop:           o.Nexthop,
		Origin:            o.Origin,
		AsPathLimit:       o.AsPathLimit,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath) {
	if s.None != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	o.Prepend = s.Prepend
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath, error) {
	var noneVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath{
		None:    noneVal,
		Prepend: o.Prepend,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone, error) {

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity, error) {
	var noneVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone, error) {

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity, error) {
	var noneVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone, error) {

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFilters) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Match != nil {
		var obj protocolBgpPolicyAggregationAddressSuppressFiltersMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFilters, error) {
	var matchVal *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressSuppressFilters{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		Match:  matchVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersMatchXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix) {
	o.Name = s.Name
	o.Exact = util.YesNo(s.Exact, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix{
		Name:  o.Name,
		Exact: util.AsBool(o.Exact, nil),
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath, error) {

	result := &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity, error) {

	result := &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFilters) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Match != nil {
		var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFilters, error) {
	var matchVal *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFilters{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		Match:  matchVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix) {
	o.Name = s.Name
	o.Exact = util.YesNo(s.Exact, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix{
		Name:  o.Name,
		Exact: util.AsBool(o.Exact, nil),
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath, error) {

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity, error) {

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisement) {
	if s.Policy != nil {
		var objs []protocolBgpPolicyConditionalAdvertisementPolicyXml
		for _, elt := range s.Policy {
			var obj protocolBgpPolicyConditionalAdvertisementPolicyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Policy = &protocolBgpPolicyConditionalAdvertisementPolicyContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisement, error) {
	var policyVal []ProtocolBgpPolicyConditionalAdvertisementPolicy
	if o.Policy != nil {
		for _, elt := range o.Policy.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			policyVal = append(policyVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyConditionalAdvertisement{
		Policy: policyVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicy) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.UsedBy != nil {
		o.UsedBy = util.StrToMem(s.UsedBy)
	}
	if s.NonExistFilters != nil {
		var objs []protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml
		for _, elt := range s.NonExistFilters {
			var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NonExistFilters = &protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersContainerXml{Entries: objs}
	}
	if s.AdvertiseFilters != nil {
		var objs []protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml
		for _, elt := range s.AdvertiseFilters {
			var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AdvertiseFilters = &protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicy, error) {
	var usedByVal []string
	if o.UsedBy != nil {
		usedByVal = util.MemToStr(o.UsedBy)
	}
	var nonExistFiltersVal []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters
	if o.NonExistFilters != nil {
		for _, elt := range o.NonExistFilters.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nonExistFiltersVal = append(nonExistFiltersVal, *obj)
		}
	}
	var advertiseFiltersVal []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters
	if o.AdvertiseFilters != nil {
		for _, elt := range o.AdvertiseFilters.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			advertiseFiltersVal = append(advertiseFiltersVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicy{
		Name:             o.Name,
		Enable:           util.AsBool(o.Enable, nil),
		UsedBy:           usedByVal,
		NonExistFilters:  nonExistFiltersVal,
		AdvertiseFilters: advertiseFiltersVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Match != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters, error) {
	var matchVal *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		Match:  matchVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Match != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters, error) {
	var matchVal *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		Match:  matchVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportXml) MarshalFromObject(s ProtocolBgpPolicyExport) {
	if s.Rules != nil {
		var objs []protocolBgpPolicyExportRulesXml
		for _, elt := range s.Rules {
			var obj protocolBgpPolicyExportRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Rules = &protocolBgpPolicyExportRulesContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportXml) UnmarshalToObject() (*ProtocolBgpPolicyExport, error) {
	var rulesVal []ProtocolBgpPolicyExportRules
	if o.Rules != nil {
		for _, elt := range o.Rules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rulesVal = append(rulesVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyExport{
		Rules: rulesVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesXml) MarshalFromObject(s ProtocolBgpPolicyExportRules) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.UsedBy != nil {
		o.UsedBy = util.StrToMem(s.UsedBy)
	}
	if s.Match != nil {
		var obj protocolBgpPolicyExportRulesMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Action != nil {
		var obj protocolBgpPolicyExportRulesActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRules, error) {
	var usedByVal []string
	if o.UsedBy != nil {
		usedByVal = util.MemToStr(o.UsedBy)
	}
	var matchVal *ProtocolBgpPolicyExportRulesMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var actionVal *ProtocolBgpPolicyExportRulesAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &ProtocolBgpPolicyExportRules{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		UsedBy: usedByVal,
		Match:  matchVal,
		Action: actionVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesMatchXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyExportRulesMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyExportRulesMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyExportRulesMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyExportRulesMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyExportRulesMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyExportRulesMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyExportRulesMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyExportRulesMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyExportRulesMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyExportRulesMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyExportRulesMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesMatchAddressPrefix) {
	o.Name = s.Name
	o.Exact = util.YesNo(s.Exact, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyExportRulesMatchAddressPrefix{
		Name:  o.Name,
		Exact: util.AsBool(o.Exact, nil),
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesMatchAsPath, error) {

	result := &ProtocolBgpPolicyExportRulesMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesMatchCommunity, error) {

	result := &ProtocolBgpPolicyExportRulesMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyExportRulesMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesAction) {
	if s.Deny != nil {
		var obj protocolBgpPolicyExportRulesActionDenyXml
		obj.MarshalFromObject(*s.Deny)
		o.Deny = &obj
	}
	if s.Allow != nil {
		var obj protocolBgpPolicyExportRulesActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesAction, error) {
	var denyVal *ProtocolBgpPolicyExportRulesActionDeny
	if o.Deny != nil {
		obj, err := o.Deny.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		denyVal = obj
	}
	var allowVal *ProtocolBgpPolicyExportRulesActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}

	result := &ProtocolBgpPolicyExportRulesAction{
		Deny:  denyVal,
		Allow: allowVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionDenyXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionDeny) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionDenyXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionDeny, error) {

	result := &ProtocolBgpPolicyExportRulesActionDeny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllow) {
	if s.Update != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateXml
		obj.MarshalFromObject(*s.Update)
		o.Update = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllow, error) {
	var updateVal *ProtocolBgpPolicyExportRulesActionAllowUpdate
	if o.Update != nil {
		obj, err := o.Update.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		updateVal = obj
	}

	result := &ProtocolBgpPolicyExportRulesActionAllow{
		Update: updateVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdate) {
	o.LocalPreference = s.LocalPreference
	o.Med = s.Med
	o.Nexthop = s.Nexthop
	o.Origin = s.Origin
	o.AsPathLimit = s.AsPathLimit
	if s.AsPath != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdate, error) {
	var asPathVal *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdate{
		LocalPreference:   o.LocalPreference,
		Med:               o.Med,
		Nexthop:           o.Nexthop,
		Origin:            o.Origin,
		AsPathLimit:       o.AsPathLimit,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateAsPathXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath) {
	if s.None != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Remove != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml
		obj.MarshalFromObject(*s.Remove)
		o.Remove = &obj
	}
	o.Prepend = s.Prepend
	o.RemoveAndPrepend = s.RemoveAndPrepend
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath, error) {
	var noneVal *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeVal *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove
	if o.Remove != nil {
		obj, err := o.Remove.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeVal = obj
	}

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath{
		None:             noneVal,
		Remove:           removeVal,
		Prepend:          o.Prepend,
		RemoveAndPrepend: o.RemoveAndPrepend,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateCommunityXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity, error) {
	var noneVal *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity, error) {
	var noneVal *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportXml) MarshalFromObject(s ProtocolBgpPolicyImport) {
	if s.Rules != nil {
		var objs []protocolBgpPolicyImportRulesXml
		for _, elt := range s.Rules {
			var obj protocolBgpPolicyImportRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Rules = &protocolBgpPolicyImportRulesContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportXml) UnmarshalToObject() (*ProtocolBgpPolicyImport, error) {
	var rulesVal []ProtocolBgpPolicyImportRules
	if o.Rules != nil {
		for _, elt := range o.Rules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rulesVal = append(rulesVal, *obj)
		}
	}

	result := &ProtocolBgpPolicyImport{
		Rules: rulesVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesXml) MarshalFromObject(s ProtocolBgpPolicyImportRules) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	if s.UsedBy != nil {
		o.UsedBy = util.StrToMem(s.UsedBy)
	}
	if s.Match != nil {
		var obj protocolBgpPolicyImportRulesMatchXml
		obj.MarshalFromObject(*s.Match)
		o.Match = &obj
	}
	if s.Action != nil {
		var obj protocolBgpPolicyImportRulesActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRules, error) {
	var usedByVal []string
	if o.UsedBy != nil {
		usedByVal = util.MemToStr(o.UsedBy)
	}
	var matchVal *ProtocolBgpPolicyImportRulesMatch
	if o.Match != nil {
		obj, err := o.Match.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchVal = obj
	}
	var actionVal *ProtocolBgpPolicyImportRulesAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &ProtocolBgpPolicyImportRules{
		Name:   o.Name,
		Enable: util.AsBool(o.Enable, nil),
		UsedBy: usedByVal,
		Match:  matchVal,
		Action: actionVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesMatchXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesMatch) {
	o.RouteTable = s.RouteTable
	o.Med = s.Med
	if s.AddressPrefix != nil {
		var objs []protocolBgpPolicyImportRulesMatchAddressPrefixXml
		for _, elt := range s.AddressPrefix {
			var obj protocolBgpPolicyImportRulesMatchAddressPrefixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AddressPrefix = &protocolBgpPolicyImportRulesMatchAddressPrefixContainerXml{Entries: objs}
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.FromPeer != nil {
		o.FromPeer = util.StrToMem(s.FromPeer)
	}
	if s.AsPath != nil {
		var obj protocolBgpPolicyImportRulesMatchAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyImportRulesMatchCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyImportRulesMatchExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesMatchXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesMatch, error) {
	var addressPrefixVal []ProtocolBgpPolicyImportRulesMatchAddressPrefix
	if o.AddressPrefix != nil {
		for _, elt := range o.AddressPrefix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressPrefixVal = append(addressPrefixVal, *obj)
		}
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var fromPeerVal []string
	if o.FromPeer != nil {
		fromPeerVal = util.MemToStr(o.FromPeer)
	}
	var asPathVal *ProtocolBgpPolicyImportRulesMatchAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyImportRulesMatchCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyImportRulesMatchExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyImportRulesMatch{
		RouteTable:        o.RouteTable,
		Med:               o.Med,
		AddressPrefix:     addressPrefixVal,
		Nexthop:           nexthopVal,
		FromPeer:          fromPeerVal,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesMatchAddressPrefixXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesMatchAddressPrefix) {
	o.Name = s.Name
	o.Exact = util.YesNo(s.Exact, nil)
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesMatchAddressPrefixXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesMatchAddressPrefix, error) {

	result := &ProtocolBgpPolicyImportRulesMatchAddressPrefix{
		Name:  o.Name,
		Exact: util.AsBool(o.Exact, nil),
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesMatchAsPathXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesMatchAsPath) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesMatchAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesMatchAsPath, error) {

	result := &ProtocolBgpPolicyImportRulesMatchAsPath{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesMatchCommunityXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesMatchCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesMatchCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesMatchCommunity, error) {

	result := &ProtocolBgpPolicyImportRulesMatchCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesMatchExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesMatchExtendedCommunity) {
	o.Regex = s.Regex
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesMatchExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesMatchExtendedCommunity, error) {

	result := &ProtocolBgpPolicyImportRulesMatchExtendedCommunity{
		Regex: o.Regex,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesAction) {
	if s.Deny != nil {
		var obj protocolBgpPolicyImportRulesActionDenyXml
		obj.MarshalFromObject(*s.Deny)
		o.Deny = &obj
	}
	if s.Allow != nil {
		var obj protocolBgpPolicyImportRulesActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesAction, error) {
	var denyVal *ProtocolBgpPolicyImportRulesActionDeny
	if o.Deny != nil {
		obj, err := o.Deny.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		denyVal = obj
	}
	var allowVal *ProtocolBgpPolicyImportRulesActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}

	result := &ProtocolBgpPolicyImportRulesAction{
		Deny:  denyVal,
		Allow: allowVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionDenyXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionDeny) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionDenyXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionDeny, error) {

	result := &ProtocolBgpPolicyImportRulesActionDeny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllow) {
	o.Dampening = s.Dampening
	if s.Update != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateXml
		obj.MarshalFromObject(*s.Update)
		o.Update = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllow, error) {
	var updateVal *ProtocolBgpPolicyImportRulesActionAllowUpdate
	if o.Update != nil {
		obj, err := o.Update.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		updateVal = obj
	}

	result := &ProtocolBgpPolicyImportRulesActionAllow{
		Dampening: o.Dampening,
		Update:    updateVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdate) {
	o.LocalPreference = s.LocalPreference
	o.Med = s.Med
	o.Weight = s.Weight
	o.Nexthop = s.Nexthop
	o.Origin = s.Origin
	o.AsPathLimit = s.AsPathLimit
	if s.AsPath != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateAsPathXml
		obj.MarshalFromObject(*s.AsPath)
		o.AsPath = &obj
	}
	if s.Community != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateCommunityXml
		obj.MarshalFromObject(*s.Community)
		o.Community = &obj
	}
	if s.ExtendedCommunity != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml
		obj.MarshalFromObject(*s.ExtendedCommunity)
		o.ExtendedCommunity = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdate, error) {
	var asPathVal *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath
	if o.AsPath != nil {
		obj, err := o.AsPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		asPathVal = obj
	}
	var communityVal *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity
	if o.Community != nil {
		obj, err := o.Community.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		communityVal = obj
	}
	var extendedCommunityVal *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity
	if o.ExtendedCommunity != nil {
		obj, err := o.ExtendedCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedCommunityVal = obj
	}

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdate{
		LocalPreference:   o.LocalPreference,
		Med:               o.Med,
		Weight:            o.Weight,
		Nexthop:           o.Nexthop,
		Origin:            o.Origin,
		AsPathLimit:       o.AsPathLimit,
		AsPath:            asPathVal,
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateAsPathXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath) {
	if s.None != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Remove != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml
		obj.MarshalFromObject(*s.Remove)
		o.Remove = &obj
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateAsPathXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath, error) {
	var noneVal *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeVal *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove
	if o.Remove != nil {
		obj, err := o.Remove.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeVal = obj
	}

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath{
		None:   noneVal,
		Remove: removeVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateCommunityXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity, error) {
	var noneVal *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity) {
	if s.None != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RemoveAll != nil {
		var obj protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml
		obj.MarshalFromObject(*s.RemoveAll)
		o.RemoveAll = &obj
	}
	o.RemoveRegex = s.RemoveRegex
	if s.Append != nil {
		o.Append = util.StrToMem(s.Append)
	}
	if s.Overwrite != nil {
		o.Overwrite = util.StrToMem(s.Overwrite)
	}
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity, error) {
	var noneVal *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var removeAllVal *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll
	if o.RemoveAll != nil {
		obj, err := o.RemoveAll.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removeAllVal = obj
	}
	var appendVal []string
	if o.Append != nil {
		appendVal = util.MemToStr(o.Append)
	}
	var overwriteVal []string
	if o.Overwrite != nil {
		overwriteVal = util.MemToStr(o.Overwrite)
	}

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity{
		None:        noneVal,
		RemoveAll:   removeAllVal,
		RemoveRegex: o.RemoveRegex,
		Append:      appendVal,
		Overwrite:   overwriteVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml) MarshalFromObject(s ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll) {
	o.Misc = s.Misc
}

func (o protocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml) UnmarshalToObject() (*ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll, error) {

	result := &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolBgpRedistRulesXml) MarshalFromObject(s ProtocolBgpRedistRules) {
	o.Name = s.Name
	o.AddressFamilyIdentifier = s.AddressFamilyIdentifier
	o.RouteTable = s.RouteTable
	o.Enable = util.YesNo(s.Enable, nil)
	o.SetOrigin = s.SetOrigin
	o.SetMed = s.SetMed
	o.SetLocalPreference = s.SetLocalPreference
	o.SetAsPathLimit = s.SetAsPathLimit
	o.Metric = s.Metric
	if s.SetCommunity != nil {
		o.SetCommunity = util.StrToMem(s.SetCommunity)
	}
	if s.SetExtendedCommunity != nil {
		o.SetExtendedCommunity = util.StrToMem(s.SetExtendedCommunity)
	}
	o.Misc = s.Misc
}

func (o protocolBgpRedistRulesXml) UnmarshalToObject() (*ProtocolBgpRedistRules, error) {
	var setCommunityVal []string
	if o.SetCommunity != nil {
		setCommunityVal = util.MemToStr(o.SetCommunity)
	}
	var setExtendedCommunityVal []string
	if o.SetExtendedCommunity != nil {
		setExtendedCommunityVal = util.MemToStr(o.SetExtendedCommunity)
	}

	result := &ProtocolBgpRedistRules{
		Name:                    o.Name,
		AddressFamilyIdentifier: o.AddressFamilyIdentifier,
		RouteTable:              o.RouteTable,
		Enable:                  util.AsBool(o.Enable, nil),
		SetOrigin:               o.SetOrigin,
		SetMed:                  o.SetMed,
		SetLocalPreference:      o.SetLocalPreference,
		SetAsPathLimit:          o.SetAsPathLimit,
		Metric:                  o.Metric,
		SetCommunity:            setCommunityVal,
		SetExtendedCommunity:    setExtendedCommunityVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolBgpRoutingOptionsXml) MarshalFromObject(s ProtocolBgpRoutingOptions) {
	if s.Aggregate != nil {
		var obj protocolBgpRoutingOptionsAggregateXml
		obj.MarshalFromObject(*s.Aggregate)
		o.Aggregate = &obj
	}
	o.AsFormat = s.AsFormat
	o.ConfederationMemberAs = s.ConfederationMemberAs
	o.DefaultLocalPreference = s.DefaultLocalPreference
	if s.GracefulRestart != nil {
		var obj protocolBgpRoutingOptionsGracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	if s.Med != nil {
		var obj protocolBgpRoutingOptionsMedXml
		obj.MarshalFromObject(*s.Med)
		o.Med = &obj
	}
	o.ReflectorClusterId = s.ReflectorClusterId
	o.Misc = s.Misc
}

func (o protocolBgpRoutingOptionsXml) UnmarshalToObject() (*ProtocolBgpRoutingOptions, error) {
	var aggregateVal *ProtocolBgpRoutingOptionsAggregate
	if o.Aggregate != nil {
		obj, err := o.Aggregate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		aggregateVal = obj
	}
	var gracefulRestartVal *ProtocolBgpRoutingOptionsGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var medVal *ProtocolBgpRoutingOptionsMed
	if o.Med != nil {
		obj, err := o.Med.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		medVal = obj
	}

	result := &ProtocolBgpRoutingOptions{
		Aggregate:              aggregateVal,
		AsFormat:               o.AsFormat,
		ConfederationMemberAs:  o.ConfederationMemberAs,
		DefaultLocalPreference: o.DefaultLocalPreference,
		GracefulRestart:        gracefulRestartVal,
		Med:                    medVal,
		ReflectorClusterId:     o.ReflectorClusterId,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *protocolBgpRoutingOptionsAggregateXml) MarshalFromObject(s ProtocolBgpRoutingOptionsAggregate) {
	o.AggregateMed = util.YesNo(s.AggregateMed, nil)
	o.Misc = s.Misc
}

func (o protocolBgpRoutingOptionsAggregateXml) UnmarshalToObject() (*ProtocolBgpRoutingOptionsAggregate, error) {

	result := &ProtocolBgpRoutingOptionsAggregate{
		AggregateMed: util.AsBool(o.AggregateMed, nil),
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *protocolBgpRoutingOptionsGracefulRestartXml) MarshalFromObject(s ProtocolBgpRoutingOptionsGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.LocalRestartTime = s.LocalRestartTime
	o.MaxPeerRestartTime = s.MaxPeerRestartTime
	o.StaleRouteTime = s.StaleRouteTime
	o.Misc = s.Misc
}

func (o protocolBgpRoutingOptionsGracefulRestartXml) UnmarshalToObject() (*ProtocolBgpRoutingOptionsGracefulRestart, error) {

	result := &ProtocolBgpRoutingOptionsGracefulRestart{
		Enable:             util.AsBool(o.Enable, nil),
		LocalRestartTime:   o.LocalRestartTime,
		MaxPeerRestartTime: o.MaxPeerRestartTime,
		StaleRouteTime:     o.StaleRouteTime,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *protocolBgpRoutingOptionsMedXml) MarshalFromObject(s ProtocolBgpRoutingOptionsMed) {
	o.AlwaysCompareMed = util.YesNo(s.AlwaysCompareMed, nil)
	o.DeterministicMedComparison = util.YesNo(s.DeterministicMedComparison, nil)
	o.Misc = s.Misc
}

func (o protocolBgpRoutingOptionsMedXml) UnmarshalToObject() (*ProtocolBgpRoutingOptionsMed, error) {

	result := &ProtocolBgpRoutingOptionsMed{
		AlwaysCompareMed:           util.AsBool(o.AlwaysCompareMed, nil),
		DeterministicMedComparison: util.AsBool(o.DeterministicMedComparison, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *protocolOspfXml) MarshalFromObject(s ProtocolOspf) {
	o.AllowRedistDefaultRoute = util.YesNo(s.AllowRedistDefaultRoute, nil)
	if s.Area != nil {
		var objs []protocolOspfAreaXml
		for _, elt := range s.Area {
			var obj protocolOspfAreaXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &protocolOspfAreaContainerXml{Entries: objs}
	}
	if s.AuthProfile != nil {
		var objs []protocolOspfAuthProfileXml
		for _, elt := range s.AuthProfile {
			var obj protocolOspfAuthProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AuthProfile = &protocolOspfAuthProfileContainerXml{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.ExportRules != nil {
		var objs []protocolOspfExportRulesXml
		for _, elt := range s.ExportRules {
			var obj protocolOspfExportRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExportRules = &protocolOspfExportRulesContainerXml{Entries: objs}
	}
	if s.GlobalBfd != nil {
		var obj protocolOspfGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj protocolOspfGracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	o.RejectDefaultRoute = util.YesNo(s.RejectDefaultRoute, nil)
	o.Rfc1583 = util.YesNo(s.Rfc1583, nil)
	o.RouterId = s.RouterId
	if s.Timers != nil {
		var obj protocolOspfTimersXml
		obj.MarshalFromObject(*s.Timers)
		o.Timers = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfXml) UnmarshalToObject() (*ProtocolOspf, error) {
	var areaVal []ProtocolOspfArea
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}
	var authProfileVal []ProtocolOspfAuthProfile
	if o.AuthProfile != nil {
		for _, elt := range o.AuthProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			authProfileVal = append(authProfileVal, *obj)
		}
	}
	var exportRulesVal []ProtocolOspfExportRules
	if o.ExportRules != nil {
		for _, elt := range o.ExportRules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			exportRulesVal = append(exportRulesVal, *obj)
		}
	}
	var globalBfdVal *ProtocolOspfGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *ProtocolOspfGracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var timersVal *ProtocolOspfTimers
	if o.Timers != nil {
		obj, err := o.Timers.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		timersVal = obj
	}

	result := &ProtocolOspf{
		AllowRedistDefaultRoute: util.AsBool(o.AllowRedistDefaultRoute, nil),
		Area:                    areaVal,
		AuthProfile:             authProfileVal,
		Enable:                  util.AsBool(o.Enable, nil),
		ExportRules:             exportRulesVal,
		GlobalBfd:               globalBfdVal,
		GracefulRestart:         gracefulRestartVal,
		RejectDefaultRoute:      util.AsBool(o.RejectDefaultRoute, nil),
		Rfc1583:                 util.AsBool(o.Rfc1583, nil),
		RouterId:                o.RouterId,
		Timers:                  timersVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaXml) MarshalFromObject(s ProtocolOspfArea) {
	o.Name = s.Name
	if s.Type != nil {
		var obj protocolOspfAreaTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []protocolOspfAreaRangeXml
		for _, elt := range s.Range {
			var obj protocolOspfAreaRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &protocolOspfAreaRangeContainerXml{Entries: objs}
	}
	if s.Interface != nil {
		var objs []protocolOspfAreaInterfaceXml
		for _, elt := range s.Interface {
			var obj protocolOspfAreaInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &protocolOspfAreaInterfaceContainerXml{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []protocolOspfAreaVirtualLinkXml
		for _, elt := range s.VirtualLink {
			var obj protocolOspfAreaVirtualLinkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &protocolOspfAreaVirtualLinkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaXml) UnmarshalToObject() (*ProtocolOspfArea, error) {
	var typeVal *ProtocolOspfAreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []ProtocolOspfAreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []ProtocolOspfAreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []ProtocolOspfAreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &ProtocolOspfArea{
		Name:        o.Name,
		Type:        typeVal,
		Range:       rangeVal,
		Interface:   interfaceVal,
		VirtualLink: virtualLinkVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeXml) MarshalFromObject(s ProtocolOspfAreaType) {
	if s.Normal != nil {
		var obj protocolOspfAreaTypeNormalXml
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj protocolOspfAreaTypeStubXml
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj protocolOspfAreaTypeNssaXml
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeXml) UnmarshalToObject() (*ProtocolOspfAreaType, error) {
	var normalVal *ProtocolOspfAreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *ProtocolOspfAreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *ProtocolOspfAreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &ProtocolOspfAreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNormalXml) MarshalFromObject(s ProtocolOspfAreaTypeNormal) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNormalXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNormal, error) {

	result := &ProtocolOspfAreaTypeNormal{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeStubXml) MarshalFromObject(s ProtocolOspfAreaTypeStub) {
	o.AcceptSummary = util.YesNo(s.AcceptSummary, nil)
	if s.DefaultRoute != nil {
		var obj protocolOspfAreaTypeStubDefaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeStubXml) UnmarshalToObject() (*ProtocolOspfAreaTypeStub, error) {
	var defaultRouteVal *ProtocolOspfAreaTypeStubDefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}

	result := &ProtocolOspfAreaTypeStub{
		AcceptSummary: util.AsBool(o.AcceptSummary, nil),
		DefaultRoute:  defaultRouteVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeStubDefaultRouteXml) MarshalFromObject(s ProtocolOspfAreaTypeStubDefaultRoute) {
	if s.Disable != nil {
		var obj protocolOspfAreaTypeStubDefaultRouteDisableXml
		obj.MarshalFromObject(*s.Disable)
		o.Disable = &obj
	}
	if s.Advertise != nil {
		var obj protocolOspfAreaTypeStubDefaultRouteAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeStubDefaultRouteXml) UnmarshalToObject() (*ProtocolOspfAreaTypeStubDefaultRoute, error) {
	var disableVal *ProtocolOspfAreaTypeStubDefaultRouteDisable
	if o.Disable != nil {
		obj, err := o.Disable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disableVal = obj
	}
	var advertiseVal *ProtocolOspfAreaTypeStubDefaultRouteAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &ProtocolOspfAreaTypeStubDefaultRoute{
		Disable:   disableVal,
		Advertise: advertiseVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeStubDefaultRouteDisableXml) MarshalFromObject(s ProtocolOspfAreaTypeStubDefaultRouteDisable) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeStubDefaultRouteDisableXml) UnmarshalToObject() (*ProtocolOspfAreaTypeStubDefaultRouteDisable, error) {

	result := &ProtocolOspfAreaTypeStubDefaultRouteDisable{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeStubDefaultRouteAdvertiseXml) MarshalFromObject(s ProtocolOspfAreaTypeStubDefaultRouteAdvertise) {
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeStubDefaultRouteAdvertiseXml) UnmarshalToObject() (*ProtocolOspfAreaTypeStubDefaultRouteAdvertise, error) {

	result := &ProtocolOspfAreaTypeStubDefaultRouteAdvertise{
		Metric: o.Metric,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaXml) MarshalFromObject(s ProtocolOspfAreaTypeNssa) {
	o.AcceptSummary = util.YesNo(s.AcceptSummary, nil)
	if s.DefaultRoute != nil {
		var obj protocolOspfAreaTypeNssaDefaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	if s.NssaExtRange != nil {
		var objs []protocolOspfAreaTypeNssaNssaExtRangeXml
		for _, elt := range s.NssaExtRange {
			var obj protocolOspfAreaTypeNssaNssaExtRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &protocolOspfAreaTypeNssaNssaExtRangeContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssa, error) {
	var defaultRouteVal *ProtocolOspfAreaTypeNssaDefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}
	var nssaExtRangeVal []ProtocolOspfAreaTypeNssaNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &ProtocolOspfAreaTypeNssa{
		AcceptSummary: util.AsBool(o.AcceptSummary, nil),
		DefaultRoute:  defaultRouteVal,
		NssaExtRange:  nssaExtRangeVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaDefaultRouteXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaDefaultRoute) {
	if s.Disable != nil {
		var obj protocolOspfAreaTypeNssaDefaultRouteDisableXml
		obj.MarshalFromObject(*s.Disable)
		o.Disable = &obj
	}
	if s.Advertise != nil {
		var obj protocolOspfAreaTypeNssaDefaultRouteAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaDefaultRouteXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaDefaultRoute, error) {
	var disableVal *ProtocolOspfAreaTypeNssaDefaultRouteDisable
	if o.Disable != nil {
		obj, err := o.Disable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disableVal = obj
	}
	var advertiseVal *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &ProtocolOspfAreaTypeNssaDefaultRoute{
		Disable:   disableVal,
		Advertise: advertiseVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaDefaultRouteDisableXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaDefaultRouteDisable) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaDefaultRouteDisableXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaDefaultRouteDisable, error) {

	result := &ProtocolOspfAreaTypeNssaDefaultRouteDisable{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaDefaultRouteAdvertiseXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaDefaultRouteAdvertise) {
	o.Metric = s.Metric
	o.Type = s.Type
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaDefaultRouteAdvertiseXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaDefaultRouteAdvertise, error) {

	result := &ProtocolOspfAreaTypeNssaDefaultRouteAdvertise{
		Metric: o.Metric,
		Type:   o.Type,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaNssaExtRangeXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaNssaExtRange) {
	o.Name = s.Name
	if s.Advertise != nil {
		var obj protocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	if s.Suppress != nil {
		var obj protocolOspfAreaTypeNssaNssaExtRangeSuppressXml
		obj.MarshalFromObject(*s.Suppress)
		o.Suppress = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaNssaExtRangeXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaNssaExtRange, error) {
	var advertiseVal *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}
	var suppressVal *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress
	if o.Suppress != nil {
		obj, err := o.Suppress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		suppressVal = obj
	}

	result := &ProtocolOspfAreaTypeNssaNssaExtRange{
		Name:      o.Name,
		Advertise: advertiseVal,
		Suppress:  suppressVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise, error) {

	result := &ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaTypeNssaNssaExtRangeSuppressXml) MarshalFromObject(s ProtocolOspfAreaTypeNssaNssaExtRangeSuppress) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaTypeNssaNssaExtRangeSuppressXml) UnmarshalToObject() (*ProtocolOspfAreaTypeNssaNssaExtRangeSuppress, error) {

	result := &ProtocolOspfAreaTypeNssaNssaExtRangeSuppress{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaRangeXml) MarshalFromObject(s ProtocolOspfAreaRange) {
	o.Name = s.Name
	if s.Advertise != nil {
		var obj protocolOspfAreaRangeAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	if s.Suppress != nil {
		var obj protocolOspfAreaRangeSuppressXml
		obj.MarshalFromObject(*s.Suppress)
		o.Suppress = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaRangeXml) UnmarshalToObject() (*ProtocolOspfAreaRange, error) {
	var advertiseVal *ProtocolOspfAreaRangeAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}
	var suppressVal *ProtocolOspfAreaRangeSuppress
	if o.Suppress != nil {
		obj, err := o.Suppress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		suppressVal = obj
	}

	result := &ProtocolOspfAreaRange{
		Name:      o.Name,
		Advertise: advertiseVal,
		Suppress:  suppressVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaRangeAdvertiseXml) MarshalFromObject(s ProtocolOspfAreaRangeAdvertise) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaRangeAdvertiseXml) UnmarshalToObject() (*ProtocolOspfAreaRangeAdvertise, error) {

	result := &ProtocolOspfAreaRangeAdvertise{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaRangeSuppressXml) MarshalFromObject(s ProtocolOspfAreaRangeSuppress) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaRangeSuppressXml) UnmarshalToObject() (*ProtocolOspfAreaRangeSuppress, error) {

	result := &ProtocolOspfAreaRangeSuppress{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceXml) MarshalFromObject(s ProtocolOspfAreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Passive = util.YesNo(s.Passive, nil)
	o.Metric = s.Metric
	o.Priority = s.Priority
	o.HelloInterval = s.HelloInterval
	o.DeadCounts = s.DeadCounts
	o.RetransmitInterval = s.RetransmitInterval
	o.TransitDelay = s.TransitDelay
	o.Authentication = s.Authentication
	o.GrDelay = s.GrDelay
	if s.LinkType != nil {
		var obj protocolOspfAreaInterfaceLinkTypeXml
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Neighbor != nil {
		var objs []protocolOspfAreaInterfaceNeighborXml
		for _, elt := range s.Neighbor {
			var obj protocolOspfAreaInterfaceNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &protocolOspfAreaInterfaceNeighborContainerXml{Entries: objs}
	}
	if s.Bfd != nil {
		var obj protocolOspfAreaInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceXml) UnmarshalToObject() (*ProtocolOspfAreaInterface, error) {
	var linkTypeVal *ProtocolOspfAreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var neighborVal []ProtocolOspfAreaInterfaceNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var bfdVal *ProtocolOspfAreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolOspfAreaInterface{
		Name:               o.Name,
		Enable:             util.AsBool(o.Enable, nil),
		Passive:            util.AsBool(o.Passive, nil),
		Metric:             o.Metric,
		Priority:           o.Priority,
		HelloInterval:      o.HelloInterval,
		DeadCounts:         o.DeadCounts,
		RetransmitInterval: o.RetransmitInterval,
		TransitDelay:       o.TransitDelay,
		Authentication:     o.Authentication,
		GrDelay:            o.GrDelay,
		LinkType:           linkTypeVal,
		Neighbor:           neighborVal,
		Bfd:                bfdVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceLinkTypeXml) MarshalFromObject(s ProtocolOspfAreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj protocolOspfAreaInterfaceLinkTypeBroadcastXml
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj protocolOspfAreaInterfaceLinkTypeP2pXml
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj protocolOspfAreaInterfaceLinkTypeP2mpXml
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceLinkTypeXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceLinkType, error) {
	var broadcastVal *ProtocolOspfAreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *ProtocolOspfAreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *ProtocolOspfAreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &ProtocolOspfAreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceLinkTypeBroadcastXml) MarshalFromObject(s ProtocolOspfAreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceLinkTypeBroadcastXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceLinkTypeBroadcast, error) {

	result := &ProtocolOspfAreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceLinkTypeP2pXml) MarshalFromObject(s ProtocolOspfAreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceLinkTypeP2pXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceLinkTypeP2p, error) {

	result := &ProtocolOspfAreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceLinkTypeP2mpXml) MarshalFromObject(s ProtocolOspfAreaInterfaceLinkTypeP2mp) {
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceLinkTypeP2mpXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceLinkTypeP2mp, error) {

	result := &ProtocolOspfAreaInterfaceLinkTypeP2mp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceNeighborXml) MarshalFromObject(s ProtocolOspfAreaInterfaceNeighbor) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceNeighborXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceNeighbor, error) {

	result := &ProtocolOspfAreaInterfaceNeighbor{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaInterfaceBfdXml) MarshalFromObject(s ProtocolOspfAreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfAreaInterfaceBfdXml) UnmarshalToObject() (*ProtocolOspfAreaInterfaceBfd, error) {

	result := &ProtocolOspfAreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaVirtualLinkXml) MarshalFromObject(s ProtocolOspfAreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.HelloInterval = s.HelloInterval
	o.DeadCounts = s.DeadCounts
	o.RetransmitInterval = s.RetransmitInterval
	o.TransitDelay = s.TransitDelay
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj protocolOspfAreaVirtualLinkBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfAreaVirtualLinkXml) UnmarshalToObject() (*ProtocolOspfAreaVirtualLink, error) {
	var bfdVal *ProtocolOspfAreaVirtualLinkBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolOspfAreaVirtualLink{
		Name:               o.Name,
		NeighborId:         o.NeighborId,
		TransitAreaId:      o.TransitAreaId,
		Enable:             util.AsBool(o.Enable, nil),
		HelloInterval:      o.HelloInterval,
		DeadCounts:         o.DeadCounts,
		RetransmitInterval: o.RetransmitInterval,
		TransitDelay:       o.TransitDelay,
		Authentication:     o.Authentication,
		Bfd:                bfdVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAreaVirtualLinkBfdXml) MarshalFromObject(s ProtocolOspfAreaVirtualLinkBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfAreaVirtualLinkBfdXml) UnmarshalToObject() (*ProtocolOspfAreaVirtualLinkBfd, error) {

	result := &ProtocolOspfAreaVirtualLinkBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAuthProfileXml) MarshalFromObject(s ProtocolOspfAuthProfile) {
	o.Name = s.Name
	o.Password = s.Password
	if s.Md5 != nil {
		var objs []protocolOspfAuthProfileMd5Xml
		for _, elt := range s.Md5 {
			var obj protocolOspfAuthProfileMd5Xml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Md5 = &protocolOspfAuthProfileMd5ContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolOspfAuthProfileXml) UnmarshalToObject() (*ProtocolOspfAuthProfile, error) {
	var md5Val []ProtocolOspfAuthProfileMd5
	if o.Md5 != nil {
		for _, elt := range o.Md5.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			md5Val = append(md5Val, *obj)
		}
	}

	result := &ProtocolOspfAuthProfile{
		Name:     o.Name,
		Password: o.Password,
		Md5:      md5Val,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolOspfAuthProfileMd5Xml) MarshalFromObject(s ProtocolOspfAuthProfileMd5) {
	o.Name = s.Name
	o.Key = s.Key
	o.Preferred = util.YesNo(s.Preferred, nil)
	o.Misc = s.Misc
}

func (o protocolOspfAuthProfileMd5Xml) UnmarshalToObject() (*ProtocolOspfAuthProfileMd5, error) {

	result := &ProtocolOspfAuthProfileMd5{
		Name:      o.Name,
		Key:       o.Key,
		Preferred: util.AsBool(o.Preferred, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfExportRulesXml) MarshalFromObject(s ProtocolOspfExportRules) {
	o.Name = s.Name
	o.NewPathType = s.NewPathType
	o.NewTag = s.NewTag
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolOspfExportRulesXml) UnmarshalToObject() (*ProtocolOspfExportRules, error) {

	result := &ProtocolOspfExportRules{
		Name:        o.Name,
		NewPathType: o.NewPathType,
		NewTag:      o.NewTag,
		Metric:      o.Metric,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolOspfGlobalBfdXml) MarshalFromObject(s ProtocolOspfGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfGlobalBfdXml) UnmarshalToObject() (*ProtocolOspfGlobalBfd, error) {

	result := &ProtocolOspfGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfGracefulRestartXml) MarshalFromObject(s ProtocolOspfGracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.Misc = s.Misc
}

func (o protocolOspfGracefulRestartXml) UnmarshalToObject() (*ProtocolOspfGracefulRestart, error) {

	result := &ProtocolOspfGracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfTimersXml) MarshalFromObject(s ProtocolOspfTimers) {
	o.LsaInterval = s.LsaInterval
	o.SpfCalculationDelay = s.SpfCalculationDelay
	o.Misc = s.Misc
}

func (o protocolOspfTimersXml) UnmarshalToObject() (*ProtocolOspfTimers, error) {

	result := &ProtocolOspfTimers{
		LsaInterval:         o.LsaInterval,
		SpfCalculationDelay: o.SpfCalculationDelay,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3Xml) MarshalFromObject(s ProtocolOspfv3) {
	o.AllowRedistDefaultRoute = util.YesNo(s.AllowRedistDefaultRoute, nil)
	if s.Area != nil {
		var objs []protocolOspfv3AreaXml
		for _, elt := range s.Area {
			var obj protocolOspfv3AreaXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Area = &protocolOspfv3AreaContainerXml{Entries: objs}
	}
	if s.AuthProfile != nil {
		var objs []protocolOspfv3AuthProfileXml
		for _, elt := range s.AuthProfile {
			var obj protocolOspfv3AuthProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AuthProfile = &protocolOspfv3AuthProfileContainerXml{Entries: objs}
	}
	o.DisableTransitTraffic = util.YesNo(s.DisableTransitTraffic, nil)
	o.Enable = util.YesNo(s.Enable, nil)
	if s.ExportRules != nil {
		var objs []protocolOspfv3ExportRulesXml
		for _, elt := range s.ExportRules {
			var obj protocolOspfv3ExportRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExportRules = &protocolOspfv3ExportRulesContainerXml{Entries: objs}
	}
	if s.GlobalBfd != nil {
		var obj protocolOspfv3GlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.GracefulRestart != nil {
		var obj protocolOspfv3GracefulRestartXml
		obj.MarshalFromObject(*s.GracefulRestart)
		o.GracefulRestart = &obj
	}
	o.RejectDefaultRoute = util.YesNo(s.RejectDefaultRoute, nil)
	o.RouterId = s.RouterId
	if s.Timers != nil {
		var obj protocolOspfv3TimersXml
		obj.MarshalFromObject(*s.Timers)
		o.Timers = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3Xml) UnmarshalToObject() (*ProtocolOspfv3, error) {
	var areaVal []ProtocolOspfv3Area
	if o.Area != nil {
		for _, elt := range o.Area.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			areaVal = append(areaVal, *obj)
		}
	}
	var authProfileVal []ProtocolOspfv3AuthProfile
	if o.AuthProfile != nil {
		for _, elt := range o.AuthProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			authProfileVal = append(authProfileVal, *obj)
		}
	}
	var exportRulesVal []ProtocolOspfv3ExportRules
	if o.ExportRules != nil {
		for _, elt := range o.ExportRules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			exportRulesVal = append(exportRulesVal, *obj)
		}
	}
	var globalBfdVal *ProtocolOspfv3GlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var gracefulRestartVal *ProtocolOspfv3GracefulRestart
	if o.GracefulRestart != nil {
		obj, err := o.GracefulRestart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gracefulRestartVal = obj
	}
	var timersVal *ProtocolOspfv3Timers
	if o.Timers != nil {
		obj, err := o.Timers.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		timersVal = obj
	}

	result := &ProtocolOspfv3{
		AllowRedistDefaultRoute: util.AsBool(o.AllowRedistDefaultRoute, nil),
		Area:                    areaVal,
		AuthProfile:             authProfileVal,
		DisableTransitTraffic:   util.AsBool(o.DisableTransitTraffic, nil),
		Enable:                  util.AsBool(o.Enable, nil),
		ExportRules:             exportRulesVal,
		GlobalBfd:               globalBfdVal,
		GracefulRestart:         gracefulRestartVal,
		RejectDefaultRoute:      util.AsBool(o.RejectDefaultRoute, nil),
		RouterId:                o.RouterId,
		Timers:                  timersVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaXml) MarshalFromObject(s ProtocolOspfv3Area) {
	o.Name = s.Name
	o.Authentication = s.Authentication
	if s.Type != nil {
		var obj protocolOspfv3AreaTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	if s.Range != nil {
		var objs []protocolOspfv3AreaRangeXml
		for _, elt := range s.Range {
			var obj protocolOspfv3AreaRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Range = &protocolOspfv3AreaRangeContainerXml{Entries: objs}
	}
	if s.Interface != nil {
		var objs []protocolOspfv3AreaInterfaceXml
		for _, elt := range s.Interface {
			var obj protocolOspfv3AreaInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &protocolOspfv3AreaInterfaceContainerXml{Entries: objs}
	}
	if s.VirtualLink != nil {
		var objs []protocolOspfv3AreaVirtualLinkXml
		for _, elt := range s.VirtualLink {
			var obj protocolOspfv3AreaVirtualLinkXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.VirtualLink = &protocolOspfv3AreaVirtualLinkContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaXml) UnmarshalToObject() (*ProtocolOspfv3Area, error) {
	var typeVal *ProtocolOspfv3AreaType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}
	var rangeVal []ProtocolOspfv3AreaRange
	if o.Range != nil {
		for _, elt := range o.Range.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rangeVal = append(rangeVal, *obj)
		}
	}
	var interfaceVal []ProtocolOspfv3AreaInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var virtualLinkVal []ProtocolOspfv3AreaVirtualLink
	if o.VirtualLink != nil {
		for _, elt := range o.VirtualLink.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			virtualLinkVal = append(virtualLinkVal, *obj)
		}
	}

	result := &ProtocolOspfv3Area{
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
func (o *protocolOspfv3AreaTypeXml) MarshalFromObject(s ProtocolOspfv3AreaType) {
	if s.Normal != nil {
		var obj protocolOspfv3AreaTypeNormalXml
		obj.MarshalFromObject(*s.Normal)
		o.Normal = &obj
	}
	if s.Stub != nil {
		var obj protocolOspfv3AreaTypeStubXml
		obj.MarshalFromObject(*s.Stub)
		o.Stub = &obj
	}
	if s.Nssa != nil {
		var obj protocolOspfv3AreaTypeNssaXml
		obj.MarshalFromObject(*s.Nssa)
		o.Nssa = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeXml) UnmarshalToObject() (*ProtocolOspfv3AreaType, error) {
	var normalVal *ProtocolOspfv3AreaTypeNormal
	if o.Normal != nil {
		obj, err := o.Normal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		normalVal = obj
	}
	var stubVal *ProtocolOspfv3AreaTypeStub
	if o.Stub != nil {
		obj, err := o.Stub.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		stubVal = obj
	}
	var nssaVal *ProtocolOspfv3AreaTypeNssa
	if o.Nssa != nil {
		obj, err := o.Nssa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nssaVal = obj
	}

	result := &ProtocolOspfv3AreaType{
		Normal: normalVal,
		Stub:   stubVal,
		Nssa:   nssaVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNormalXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNormal) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNormalXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNormal, error) {

	result := &ProtocolOspfv3AreaTypeNormal{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeStubXml) MarshalFromObject(s ProtocolOspfv3AreaTypeStub) {
	o.AcceptSummary = util.YesNo(s.AcceptSummary, nil)
	if s.DefaultRoute != nil {
		var obj protocolOspfv3AreaTypeStubDefaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeStubXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeStub, error) {
	var defaultRouteVal *ProtocolOspfv3AreaTypeStubDefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}

	result := &ProtocolOspfv3AreaTypeStub{
		AcceptSummary: util.AsBool(o.AcceptSummary, nil),
		DefaultRoute:  defaultRouteVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeStubDefaultRouteXml) MarshalFromObject(s ProtocolOspfv3AreaTypeStubDefaultRoute) {
	if s.Disable != nil {
		var obj protocolOspfv3AreaTypeStubDefaultRouteDisableXml
		obj.MarshalFromObject(*s.Disable)
		o.Disable = &obj
	}
	if s.Advertise != nil {
		var obj protocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeStubDefaultRouteXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeStubDefaultRoute, error) {
	var disableVal *ProtocolOspfv3AreaTypeStubDefaultRouteDisable
	if o.Disable != nil {
		obj, err := o.Disable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disableVal = obj
	}
	var advertiseVal *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &ProtocolOspfv3AreaTypeStubDefaultRoute{
		Disable:   disableVal,
		Advertise: advertiseVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeStubDefaultRouteDisableXml) MarshalFromObject(s ProtocolOspfv3AreaTypeStubDefaultRouteDisable) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeStubDefaultRouteDisableXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeStubDefaultRouteDisable, error) {

	result := &ProtocolOspfv3AreaTypeStubDefaultRouteDisable{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml) MarshalFromObject(s ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise) {
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise, error) {

	result := &ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise{
		Metric: o.Metric,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssa) {
	o.AcceptSummary = util.YesNo(s.AcceptSummary, nil)
	if s.DefaultRoute != nil {
		var obj protocolOspfv3AreaTypeNssaDefaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	if s.NssaExtRange != nil {
		var objs []protocolOspfv3AreaTypeNssaNssaExtRangeXml
		for _, elt := range s.NssaExtRange {
			var obj protocolOspfv3AreaTypeNssaNssaExtRangeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NssaExtRange = &protocolOspfv3AreaTypeNssaNssaExtRangeContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssa, error) {
	var defaultRouteVal *ProtocolOspfv3AreaTypeNssaDefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}
	var nssaExtRangeVal []ProtocolOspfv3AreaTypeNssaNssaExtRange
	if o.NssaExtRange != nil {
		for _, elt := range o.NssaExtRange.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nssaExtRangeVal = append(nssaExtRangeVal, *obj)
		}
	}

	result := &ProtocolOspfv3AreaTypeNssa{
		AcceptSummary: util.AsBool(o.AcceptSummary, nil),
		DefaultRoute:  defaultRouteVal,
		NssaExtRange:  nssaExtRangeVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaDefaultRouteXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaDefaultRoute) {
	if s.Disable != nil {
		var obj protocolOspfv3AreaTypeNssaDefaultRouteDisableXml
		obj.MarshalFromObject(*s.Disable)
		o.Disable = &obj
	}
	if s.Advertise != nil {
		var obj protocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaDefaultRouteXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaDefaultRoute, error) {
	var disableVal *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable
	if o.Disable != nil {
		obj, err := o.Disable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disableVal = obj
	}
	var advertiseVal *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &ProtocolOspfv3AreaTypeNssaDefaultRoute{
		Disable:   disableVal,
		Advertise: advertiseVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaDefaultRouteDisableXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaDefaultRouteDisable) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaDefaultRouteDisableXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaDefaultRouteDisable, error) {

	result := &ProtocolOspfv3AreaTypeNssaDefaultRouteDisable{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise) {
	o.Metric = s.Metric
	o.Type = s.Type
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise, error) {

	result := &ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise{
		Metric: o.Metric,
		Type:   o.Type,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaNssaExtRangeXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaNssaExtRange) {
	o.Name = s.Name
	if s.Advertise != nil {
		var obj protocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	if s.Suppress != nil {
		var obj protocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml
		obj.MarshalFromObject(*s.Suppress)
		o.Suppress = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaNssaExtRangeXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaNssaExtRange, error) {
	var advertiseVal *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}
	var suppressVal *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress
	if o.Suppress != nil {
		obj, err := o.Suppress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		suppressVal = obj
	}

	result := &ProtocolOspfv3AreaTypeNssaNssaExtRange{
		Name:      o.Name,
		Advertise: advertiseVal,
		Suppress:  suppressVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise, error) {

	result := &ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml) MarshalFromObject(s ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml) UnmarshalToObject() (*ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress, error) {

	result := &ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaRangeXml) MarshalFromObject(s ProtocolOspfv3AreaRange) {
	o.Name = s.Name
	if s.Advertise != nil {
		var obj protocolOspfv3AreaRangeAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	if s.Suppress != nil {
		var obj protocolOspfv3AreaRangeSuppressXml
		obj.MarshalFromObject(*s.Suppress)
		o.Suppress = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaRangeXml) UnmarshalToObject() (*ProtocolOspfv3AreaRange, error) {
	var advertiseVal *ProtocolOspfv3AreaRangeAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}
	var suppressVal *ProtocolOspfv3AreaRangeSuppress
	if o.Suppress != nil {
		obj, err := o.Suppress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		suppressVal = obj
	}

	result := &ProtocolOspfv3AreaRange{
		Name:      o.Name,
		Advertise: advertiseVal,
		Suppress:  suppressVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaRangeAdvertiseXml) MarshalFromObject(s ProtocolOspfv3AreaRangeAdvertise) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaRangeAdvertiseXml) UnmarshalToObject() (*ProtocolOspfv3AreaRangeAdvertise, error) {

	result := &ProtocolOspfv3AreaRangeAdvertise{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaRangeSuppressXml) MarshalFromObject(s ProtocolOspfv3AreaRangeSuppress) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaRangeSuppressXml) UnmarshalToObject() (*ProtocolOspfv3AreaRangeSuppress, error) {

	result := &ProtocolOspfv3AreaRangeSuppress{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceXml) MarshalFromObject(s ProtocolOspfv3AreaInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.Passive = util.YesNo(s.Passive, nil)
	o.Metric = s.Metric
	o.Priority = s.Priority
	o.HelloInterval = s.HelloInterval
	o.DeadCounts = s.DeadCounts
	o.RetransmitInterval = s.RetransmitInterval
	o.TransitDelay = s.TransitDelay
	o.Authentication = s.Authentication
	o.GrDelay = s.GrDelay
	if s.LinkType != nil {
		var obj protocolOspfv3AreaInterfaceLinkTypeXml
		obj.MarshalFromObject(*s.LinkType)
		o.LinkType = &obj
	}
	if s.Neighbor != nil {
		var objs []protocolOspfv3AreaInterfaceNeighborXml
		for _, elt := range s.Neighbor {
			var obj protocolOspfv3AreaInterfaceNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &protocolOspfv3AreaInterfaceNeighborContainerXml{Entries: objs}
	}
	if s.Bfd != nil {
		var obj protocolOspfv3AreaInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterface, error) {
	var linkTypeVal *ProtocolOspfv3AreaInterfaceLinkType
	if o.LinkType != nil {
		obj, err := o.LinkType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linkTypeVal = obj
	}
	var neighborVal []ProtocolOspfv3AreaInterfaceNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var bfdVal *ProtocolOspfv3AreaInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolOspfv3AreaInterface{
		Name:               o.Name,
		Enable:             util.AsBool(o.Enable, nil),
		InstanceId:         o.InstanceId,
		Passive:            util.AsBool(o.Passive, nil),
		Metric:             o.Metric,
		Priority:           o.Priority,
		HelloInterval:      o.HelloInterval,
		DeadCounts:         o.DeadCounts,
		RetransmitInterval: o.RetransmitInterval,
		TransitDelay:       o.TransitDelay,
		Authentication:     o.Authentication,
		GrDelay:            o.GrDelay,
		LinkType:           linkTypeVal,
		Neighbor:           neighborVal,
		Bfd:                bfdVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceLinkTypeXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceLinkType) {
	if s.Broadcast != nil {
		var obj protocolOspfv3AreaInterfaceLinkTypeBroadcastXml
		obj.MarshalFromObject(*s.Broadcast)
		o.Broadcast = &obj
	}
	if s.P2p != nil {
		var obj protocolOspfv3AreaInterfaceLinkTypeP2pXml
		obj.MarshalFromObject(*s.P2p)
		o.P2p = &obj
	}
	if s.P2mp != nil {
		var obj protocolOspfv3AreaInterfaceLinkTypeP2mpXml
		obj.MarshalFromObject(*s.P2mp)
		o.P2mp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceLinkTypeXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceLinkType, error) {
	var broadcastVal *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast
	if o.Broadcast != nil {
		obj, err := o.Broadcast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		broadcastVal = obj
	}
	var p2pVal *ProtocolOspfv3AreaInterfaceLinkTypeP2p
	if o.P2p != nil {
		obj, err := o.P2p.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2pVal = obj
	}
	var p2mpVal *ProtocolOspfv3AreaInterfaceLinkTypeP2mp
	if o.P2mp != nil {
		obj, err := o.P2mp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		p2mpVal = obj
	}

	result := &ProtocolOspfv3AreaInterfaceLinkType{
		Broadcast: broadcastVal,
		P2p:       p2pVal,
		P2mp:      p2mpVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceLinkTypeBroadcastXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceLinkTypeBroadcast) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceLinkTypeBroadcastXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceLinkTypeBroadcast, error) {

	result := &ProtocolOspfv3AreaInterfaceLinkTypeBroadcast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceLinkTypeP2pXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceLinkTypeP2p) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceLinkTypeP2pXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceLinkTypeP2p, error) {

	result := &ProtocolOspfv3AreaInterfaceLinkTypeP2p{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceLinkTypeP2mpXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceLinkTypeP2mp) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceLinkTypeP2mpXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceLinkTypeP2mp, error) {

	result := &ProtocolOspfv3AreaInterfaceLinkTypeP2mp{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceNeighborXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceNeighbor) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceNeighborXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceNeighbor, error) {

	result := &ProtocolOspfv3AreaInterfaceNeighbor{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaInterfaceBfdXml) MarshalFromObject(s ProtocolOspfv3AreaInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaInterfaceBfdXml) UnmarshalToObject() (*ProtocolOspfv3AreaInterfaceBfd, error) {

	result := &ProtocolOspfv3AreaInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaVirtualLinkXml) MarshalFromObject(s ProtocolOspfv3AreaVirtualLink) {
	o.Name = s.Name
	o.NeighborId = s.NeighborId
	o.TransitAreaId = s.TransitAreaId
	o.Enable = util.YesNo(s.Enable, nil)
	o.InstanceId = s.InstanceId
	o.HelloInterval = s.HelloInterval
	o.DeadCounts = s.DeadCounts
	o.RetransmitInterval = s.RetransmitInterval
	o.TransitDelay = s.TransitDelay
	o.Authentication = s.Authentication
	if s.Bfd != nil {
		var obj protocolOspfv3AreaVirtualLinkBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaVirtualLinkXml) UnmarshalToObject() (*ProtocolOspfv3AreaVirtualLink, error) {
	var bfdVal *ProtocolOspfv3AreaVirtualLinkBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolOspfv3AreaVirtualLink{
		Name:               o.Name,
		NeighborId:         o.NeighborId,
		TransitAreaId:      o.TransitAreaId,
		Enable:             util.AsBool(o.Enable, nil),
		InstanceId:         o.InstanceId,
		HelloInterval:      o.HelloInterval,
		DeadCounts:         o.DeadCounts,
		RetransmitInterval: o.RetransmitInterval,
		TransitDelay:       o.TransitDelay,
		Authentication:     o.Authentication,
		Bfd:                bfdVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AreaVirtualLinkBfdXml) MarshalFromObject(s ProtocolOspfv3AreaVirtualLinkBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfv3AreaVirtualLinkBfdXml) UnmarshalToObject() (*ProtocolOspfv3AreaVirtualLinkBfd, error) {

	result := &ProtocolOspfv3AreaVirtualLinkBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileXml) MarshalFromObject(s ProtocolOspfv3AuthProfile) {
	o.Name = s.Name
	o.Spi = s.Spi
	if s.Esp != nil {
		var obj protocolOspfv3AuthProfileEspXml
		obj.MarshalFromObject(*s.Esp)
		o.Esp = &obj
	}
	if s.Ah != nil {
		var obj protocolOspfv3AuthProfileAhXml
		obj.MarshalFromObject(*s.Ah)
		o.Ah = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfile, error) {
	var espVal *ProtocolOspfv3AuthProfileEsp
	if o.Esp != nil {
		obj, err := o.Esp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		espVal = obj
	}
	var ahVal *ProtocolOspfv3AuthProfileAh
	if o.Ah != nil {
		obj, err := o.Ah.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ahVal = obj
	}

	result := &ProtocolOspfv3AuthProfile{
		Name: o.Name,
		Spi:  o.Spi,
		Esp:  espVal,
		Ah:   ahVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspXml) MarshalFromObject(s ProtocolOspfv3AuthProfileEsp) {
	if s.Authentication != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationXml
		obj.MarshalFromObject(*s.Authentication)
		o.Authentication = &obj
	}
	if s.Encryption != nil {
		var obj protocolOspfv3AuthProfileEspEncryptionXml
		obj.MarshalFromObject(*s.Encryption)
		o.Encryption = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEsp, error) {
	var authenticationVal *ProtocolOspfv3AuthProfileEspAuthentication
	if o.Authentication != nil {
		obj, err := o.Authentication.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationVal = obj
	}
	var encryptionVal *ProtocolOspfv3AuthProfileEspEncryption
	if o.Encryption != nil {
		obj, err := o.Encryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		encryptionVal = obj
	}

	result := &ProtocolOspfv3AuthProfileEsp{
		Authentication: authenticationVal,
		Encryption:     encryptionVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationXml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthentication) {
	if s.Md5 != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	if s.None != nil {
		var obj protocolOspfv3AuthProfileEspAuthenticationNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthentication, error) {
	var md5Val *ProtocolOspfv3AuthProfileEspAuthenticationMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *ProtocolOspfv3AuthProfileEspAuthenticationSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ProtocolOspfv3AuthProfileEspAuthenticationSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ProtocolOspfv3AuthProfileEspAuthenticationSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ProtocolOspfv3AuthProfileEspAuthenticationSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}
	var noneVal *ProtocolOspfv3AuthProfileEspAuthenticationNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}

	result := &ProtocolOspfv3AuthProfileEspAuthentication{
		Md5:    md5Val,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		None:   noneVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationMd5Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationMd5Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationMd5, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationSha1Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationSha1Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationSha1, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationSha256Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationSha256Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationSha256, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationSha384Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationSha384Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationSha384, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationSha512Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationSha512Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationSha512, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspAuthenticationNoneXml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspAuthenticationNone) {
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspAuthenticationNoneXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspAuthenticationNone, error) {

	result := &ProtocolOspfv3AuthProfileEspAuthenticationNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileEspEncryptionXml) MarshalFromObject(s ProtocolOspfv3AuthProfileEspEncryption) {
	o.Algorithm = s.Algorithm
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileEspEncryptionXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileEspEncryption, error) {

	result := &ProtocolOspfv3AuthProfileEspEncryption{
		Algorithm: o.Algorithm,
		Key:       o.Key,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhXml) MarshalFromObject(s ProtocolOspfv3AuthProfileAh) {
	if s.Md5 != nil {
		var obj protocolOspfv3AuthProfileAhMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj protocolOspfv3AuthProfileAhSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj protocolOspfv3AuthProfileAhSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj protocolOspfv3AuthProfileAhSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj protocolOspfv3AuthProfileAhSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhXml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAh, error) {
	var md5Val *ProtocolOspfv3AuthProfileAhMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *ProtocolOspfv3AuthProfileAhSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ProtocolOspfv3AuthProfileAhSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ProtocolOspfv3AuthProfileAhSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ProtocolOspfv3AuthProfileAhSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &ProtocolOspfv3AuthProfileAh{
		Md5:    md5Val,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhMd5Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileAhMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhMd5Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAhMd5, error) {

	result := &ProtocolOspfv3AuthProfileAhMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhSha1Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileAhSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhSha1Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAhSha1, error) {

	result := &ProtocolOspfv3AuthProfileAhSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhSha256Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileAhSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhSha256Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAhSha256, error) {

	result := &ProtocolOspfv3AuthProfileAhSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhSha384Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileAhSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhSha384Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAhSha384, error) {

	result := &ProtocolOspfv3AuthProfileAhSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3AuthProfileAhSha512Xml) MarshalFromObject(s ProtocolOspfv3AuthProfileAhSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o protocolOspfv3AuthProfileAhSha512Xml) UnmarshalToObject() (*ProtocolOspfv3AuthProfileAhSha512, error) {

	result := &ProtocolOspfv3AuthProfileAhSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3ExportRulesXml) MarshalFromObject(s ProtocolOspfv3ExportRules) {
	o.Name = s.Name
	o.NewPathType = s.NewPathType
	o.NewTag = s.NewTag
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolOspfv3ExportRulesXml) UnmarshalToObject() (*ProtocolOspfv3ExportRules, error) {

	result := &ProtocolOspfv3ExportRules{
		Name:        o.Name,
		NewPathType: o.NewPathType,
		NewTag:      o.NewTag,
		Metric:      o.Metric,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3GlobalBfdXml) MarshalFromObject(s ProtocolOspfv3GlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolOspfv3GlobalBfdXml) UnmarshalToObject() (*ProtocolOspfv3GlobalBfd, error) {

	result := &ProtocolOspfv3GlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3GracefulRestartXml) MarshalFromObject(s ProtocolOspfv3GracefulRestart) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GracePeriod = s.GracePeriod
	o.HelperEnable = util.YesNo(s.HelperEnable, nil)
	o.MaxNeighborRestartTime = s.MaxNeighborRestartTime
	o.StrictLSAChecking = util.YesNo(s.StrictLSAChecking, nil)
	o.Misc = s.Misc
}

func (o protocolOspfv3GracefulRestartXml) UnmarshalToObject() (*ProtocolOspfv3GracefulRestart, error) {

	result := &ProtocolOspfv3GracefulRestart{
		Enable:                 util.AsBool(o.Enable, nil),
		GracePeriod:            o.GracePeriod,
		HelperEnable:           util.AsBool(o.HelperEnable, nil),
		MaxNeighborRestartTime: o.MaxNeighborRestartTime,
		StrictLSAChecking:      util.AsBool(o.StrictLSAChecking, nil),
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *protocolOspfv3TimersXml) MarshalFromObject(s ProtocolOspfv3Timers) {
	o.LsaInterval = s.LsaInterval
	o.SpfCalculationDelay = s.SpfCalculationDelay
	o.Misc = s.Misc
}

func (o protocolOspfv3TimersXml) UnmarshalToObject() (*ProtocolOspfv3Timers, error) {

	result := &ProtocolOspfv3Timers{
		LsaInterval:         o.LsaInterval,
		SpfCalculationDelay: o.SpfCalculationDelay,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileXml) MarshalFromObject(s ProtocolRedistProfile) {
	o.Name = s.Name
	o.Priority = s.Priority
	if s.Filter != nil {
		var obj protocolRedistProfileFilterXml
		obj.MarshalFromObject(*s.Filter)
		o.Filter = &obj
	}
	if s.Action != nil {
		var obj protocolRedistProfileActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileXml) UnmarshalToObject() (*ProtocolRedistProfile, error) {
	var filterVal *ProtocolRedistProfileFilter
	if o.Filter != nil {
		obj, err := o.Filter.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filterVal = obj
	}
	var actionVal *ProtocolRedistProfileAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &ProtocolRedistProfile{
		Name:     o.Name,
		Priority: o.Priority,
		Filter:   filterVal,
		Action:   actionVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileFilterXml) MarshalFromObject(s ProtocolRedistProfileFilter) {
	if s.Type != nil {
		o.Type = util.StrToMem(s.Type)
	}
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.Destination != nil {
		o.Destination = util.StrToMem(s.Destination)
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.Ospf != nil {
		var obj protocolRedistProfileFilterOspfXml
		obj.MarshalFromObject(*s.Ospf)
		o.Ospf = &obj
	}
	if s.Bgp != nil {
		var obj protocolRedistProfileFilterBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileFilterXml) UnmarshalToObject() (*ProtocolRedistProfileFilter, error) {
	var typeVal []string
	if o.Type != nil {
		typeVal = util.MemToStr(o.Type)
	}
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var destinationVal []string
	if o.Destination != nil {
		destinationVal = util.MemToStr(o.Destination)
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var ospfVal *ProtocolRedistProfileFilterOspf
	if o.Ospf != nil {
		obj, err := o.Ospf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfVal = obj
	}
	var bgpVal *ProtocolRedistProfileFilterBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}

	result := &ProtocolRedistProfileFilter{
		Type:        typeVal,
		Interface:   interfaceVal,
		Destination: destinationVal,
		Nexthop:     nexthopVal,
		Ospf:        ospfVal,
		Bgp:         bgpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileFilterOspfXml) MarshalFromObject(s ProtocolRedistProfileFilterOspf) {
	if s.PathType != nil {
		o.PathType = util.StrToMem(s.PathType)
	}
	if s.Area != nil {
		o.Area = util.StrToMem(s.Area)
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileFilterOspfXml) UnmarshalToObject() (*ProtocolRedistProfileFilterOspf, error) {
	var pathTypeVal []string
	if o.PathType != nil {
		pathTypeVal = util.MemToStr(o.PathType)
	}
	var areaVal []string
	if o.Area != nil {
		areaVal = util.MemToStr(o.Area)
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}

	result := &ProtocolRedistProfileFilterOspf{
		PathType: pathTypeVal,
		Area:     areaVal,
		Tag:      tagVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileFilterBgpXml) MarshalFromObject(s ProtocolRedistProfileFilterBgp) {
	if s.Community != nil {
		o.Community = util.StrToMem(s.Community)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileFilterBgpXml) UnmarshalToObject() (*ProtocolRedistProfileFilterBgp, error) {
	var communityVal []string
	if o.Community != nil {
		communityVal = util.MemToStr(o.Community)
	}
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &ProtocolRedistProfileFilterBgp{
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileActionXml) MarshalFromObject(s ProtocolRedistProfileAction) {
	if s.NoRedist != nil {
		var obj protocolRedistProfileActionNoRedistXml
		obj.MarshalFromObject(*s.NoRedist)
		o.NoRedist = &obj
	}
	if s.Redist != nil {
		var obj protocolRedistProfileActionRedistXml
		obj.MarshalFromObject(*s.Redist)
		o.Redist = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileActionXml) UnmarshalToObject() (*ProtocolRedistProfileAction, error) {
	var noRedistVal *ProtocolRedistProfileActionNoRedist
	if o.NoRedist != nil {
		obj, err := o.NoRedist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noRedistVal = obj
	}
	var redistVal *ProtocolRedistProfileActionRedist
	if o.Redist != nil {
		obj, err := o.Redist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redistVal = obj
	}

	result := &ProtocolRedistProfileAction{
		NoRedist: noRedistVal,
		Redist:   redistVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileActionNoRedistXml) MarshalFromObject(s ProtocolRedistProfileActionNoRedist) {
	o.Misc = s.Misc
}

func (o protocolRedistProfileActionNoRedistXml) UnmarshalToObject() (*ProtocolRedistProfileActionNoRedist, error) {

	result := &ProtocolRedistProfileActionNoRedist{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileActionRedistXml) MarshalFromObject(s ProtocolRedistProfileActionRedist) {
	o.Misc = s.Misc
}

func (o protocolRedistProfileActionRedistXml) UnmarshalToObject() (*ProtocolRedistProfileActionRedist, error) {

	result := &ProtocolRedistProfileActionRedist{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6Xml) MarshalFromObject(s ProtocolRedistProfileIpv6) {
	o.Name = s.Name
	o.Priority = s.Priority
	if s.Filter != nil {
		var obj protocolRedistProfileIpv6FilterXml
		obj.MarshalFromObject(*s.Filter)
		o.Filter = &obj
	}
	if s.Action != nil {
		var obj protocolRedistProfileIpv6ActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6Xml) UnmarshalToObject() (*ProtocolRedistProfileIpv6, error) {
	var filterVal *ProtocolRedistProfileIpv6Filter
	if o.Filter != nil {
		obj, err := o.Filter.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		filterVal = obj
	}
	var actionVal *ProtocolRedistProfileIpv6Action
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &ProtocolRedistProfileIpv6{
		Name:     o.Name,
		Priority: o.Priority,
		Filter:   filterVal,
		Action:   actionVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6FilterXml) MarshalFromObject(s ProtocolRedistProfileIpv6Filter) {
	if s.Type != nil {
		o.Type = util.StrToMem(s.Type)
	}
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.Destination != nil {
		o.Destination = util.StrToMem(s.Destination)
	}
	if s.Nexthop != nil {
		o.Nexthop = util.StrToMem(s.Nexthop)
	}
	if s.Ospfv3 != nil {
		var obj protocolRedistProfileIpv6FilterOspfv3Xml
		obj.MarshalFromObject(*s.Ospfv3)
		o.Ospfv3 = &obj
	}
	if s.Bgp != nil {
		var obj protocolRedistProfileIpv6FilterBgpXml
		obj.MarshalFromObject(*s.Bgp)
		o.Bgp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6FilterXml) UnmarshalToObject() (*ProtocolRedistProfileIpv6Filter, error) {
	var typeVal []string
	if o.Type != nil {
		typeVal = util.MemToStr(o.Type)
	}
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var destinationVal []string
	if o.Destination != nil {
		destinationVal = util.MemToStr(o.Destination)
	}
	var nexthopVal []string
	if o.Nexthop != nil {
		nexthopVal = util.MemToStr(o.Nexthop)
	}
	var ospfv3Val *ProtocolRedistProfileIpv6FilterOspfv3
	if o.Ospfv3 != nil {
		obj, err := o.Ospfv3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ospfv3Val = obj
	}
	var bgpVal *ProtocolRedistProfileIpv6FilterBgp
	if o.Bgp != nil {
		obj, err := o.Bgp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bgpVal = obj
	}

	result := &ProtocolRedistProfileIpv6Filter{
		Type:        typeVal,
		Interface:   interfaceVal,
		Destination: destinationVal,
		Nexthop:     nexthopVal,
		Ospfv3:      ospfv3Val,
		Bgp:         bgpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6FilterOspfv3Xml) MarshalFromObject(s ProtocolRedistProfileIpv6FilterOspfv3) {
	if s.PathType != nil {
		o.PathType = util.StrToMem(s.PathType)
	}
	if s.Area != nil {
		o.Area = util.StrToMem(s.Area)
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6FilterOspfv3Xml) UnmarshalToObject() (*ProtocolRedistProfileIpv6FilterOspfv3, error) {
	var pathTypeVal []string
	if o.PathType != nil {
		pathTypeVal = util.MemToStr(o.PathType)
	}
	var areaVal []string
	if o.Area != nil {
		areaVal = util.MemToStr(o.Area)
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}

	result := &ProtocolRedistProfileIpv6FilterOspfv3{
		PathType: pathTypeVal,
		Area:     areaVal,
		Tag:      tagVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6FilterBgpXml) MarshalFromObject(s ProtocolRedistProfileIpv6FilterBgp) {
	if s.Community != nil {
		o.Community = util.StrToMem(s.Community)
	}
	if s.ExtendedCommunity != nil {
		o.ExtendedCommunity = util.StrToMem(s.ExtendedCommunity)
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6FilterBgpXml) UnmarshalToObject() (*ProtocolRedistProfileIpv6FilterBgp, error) {
	var communityVal []string
	if o.Community != nil {
		communityVal = util.MemToStr(o.Community)
	}
	var extendedCommunityVal []string
	if o.ExtendedCommunity != nil {
		extendedCommunityVal = util.MemToStr(o.ExtendedCommunity)
	}

	result := &ProtocolRedistProfileIpv6FilterBgp{
		Community:         communityVal,
		ExtendedCommunity: extendedCommunityVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6ActionXml) MarshalFromObject(s ProtocolRedistProfileIpv6Action) {
	if s.NoRedist != nil {
		var obj protocolRedistProfileIpv6ActionNoRedistXml
		obj.MarshalFromObject(*s.NoRedist)
		o.NoRedist = &obj
	}
	if s.Redist != nil {
		var obj protocolRedistProfileIpv6ActionRedistXml
		obj.MarshalFromObject(*s.Redist)
		o.Redist = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6ActionXml) UnmarshalToObject() (*ProtocolRedistProfileIpv6Action, error) {
	var noRedistVal *ProtocolRedistProfileIpv6ActionNoRedist
	if o.NoRedist != nil {
		obj, err := o.NoRedist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noRedistVal = obj
	}
	var redistVal *ProtocolRedistProfileIpv6ActionRedist
	if o.Redist != nil {
		obj, err := o.Redist.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redistVal = obj
	}

	result := &ProtocolRedistProfileIpv6Action{
		NoRedist: noRedistVal,
		Redist:   redistVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6ActionNoRedistXml) MarshalFromObject(s ProtocolRedistProfileIpv6ActionNoRedist) {
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6ActionNoRedistXml) UnmarshalToObject() (*ProtocolRedistProfileIpv6ActionNoRedist, error) {

	result := &ProtocolRedistProfileIpv6ActionNoRedist{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolRedistProfileIpv6ActionRedistXml) MarshalFromObject(s ProtocolRedistProfileIpv6ActionRedist) {
	o.Misc = s.Misc
}

func (o protocolRedistProfileIpv6ActionRedistXml) UnmarshalToObject() (*ProtocolRedistProfileIpv6ActionRedist, error) {

	result := &ProtocolRedistProfileIpv6ActionRedist{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolRipXml) MarshalFromObject(s ProtocolRip) {
	o.AllowRedistDefaultRoute = util.YesNo(s.AllowRedistDefaultRoute, nil)
	if s.AuthProfile != nil {
		var objs []protocolRipAuthProfileXml
		for _, elt := range s.AuthProfile {
			var obj protocolRipAuthProfileXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AuthProfile = &protocolRipAuthProfileContainerXml{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.ExportRules != nil {
		var objs []protocolRipExportRulesXml
		for _, elt := range s.ExportRules {
			var obj protocolRipExportRulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExportRules = &protocolRipExportRulesContainerXml{Entries: objs}
	}
	if s.GlobalBfd != nil {
		var obj protocolRipGlobalBfdXml
		obj.MarshalFromObject(*s.GlobalBfd)
		o.GlobalBfd = &obj
	}
	if s.Interface != nil {
		var objs []protocolRipInterfaceXml
		for _, elt := range s.Interface {
			var obj protocolRipInterfaceXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Interface = &protocolRipInterfaceContainerXml{Entries: objs}
	}
	o.RejectDefaultRoute = util.YesNo(s.RejectDefaultRoute, nil)
	if s.Timers != nil {
		var obj protocolRipTimersXml
		obj.MarshalFromObject(*s.Timers)
		o.Timers = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRipXml) UnmarshalToObject() (*ProtocolRip, error) {
	var authProfileVal []ProtocolRipAuthProfile
	if o.AuthProfile != nil {
		for _, elt := range o.AuthProfile.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			authProfileVal = append(authProfileVal, *obj)
		}
	}
	var exportRulesVal []ProtocolRipExportRules
	if o.ExportRules != nil {
		for _, elt := range o.ExportRules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			exportRulesVal = append(exportRulesVal, *obj)
		}
	}
	var globalBfdVal *ProtocolRipGlobalBfd
	if o.GlobalBfd != nil {
		obj, err := o.GlobalBfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalBfdVal = obj
	}
	var interfaceVal []ProtocolRipInterface
	if o.Interface != nil {
		for _, elt := range o.Interface.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			interfaceVal = append(interfaceVal, *obj)
		}
	}
	var timersVal *ProtocolRipTimers
	if o.Timers != nil {
		obj, err := o.Timers.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		timersVal = obj
	}

	result := &ProtocolRip{
		AllowRedistDefaultRoute: util.AsBool(o.AllowRedistDefaultRoute, nil),
		AuthProfile:             authProfileVal,
		Enable:                  util.AsBool(o.Enable, nil),
		ExportRules:             exportRulesVal,
		GlobalBfd:               globalBfdVal,
		Interface:               interfaceVal,
		RejectDefaultRoute:      util.AsBool(o.RejectDefaultRoute, nil),
		Timers:                  timersVal,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *protocolRipAuthProfileXml) MarshalFromObject(s ProtocolRipAuthProfile) {
	o.Name = s.Name
	o.Password = s.Password
	if s.Md5 != nil {
		var objs []protocolRipAuthProfileMd5Xml
		for _, elt := range s.Md5 {
			var obj protocolRipAuthProfileMd5Xml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Md5 = &protocolRipAuthProfileMd5ContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o protocolRipAuthProfileXml) UnmarshalToObject() (*ProtocolRipAuthProfile, error) {
	var md5Val []ProtocolRipAuthProfileMd5
	if o.Md5 != nil {
		for _, elt := range o.Md5.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			md5Val = append(md5Val, *obj)
		}
	}

	result := &ProtocolRipAuthProfile{
		Name:     o.Name,
		Password: o.Password,
		Md5:      md5Val,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolRipAuthProfileMd5Xml) MarshalFromObject(s ProtocolRipAuthProfileMd5) {
	o.Name = s.Name
	o.Key = s.Key
	o.Preferred = util.YesNo(s.Preferred, nil)
	o.Misc = s.Misc
}

func (o protocolRipAuthProfileMd5Xml) UnmarshalToObject() (*ProtocolRipAuthProfileMd5, error) {

	result := &ProtocolRipAuthProfileMd5{
		Name:      o.Name,
		Key:       o.Key,
		Preferred: util.AsBool(o.Preferred, nil),
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolRipExportRulesXml) MarshalFromObject(s ProtocolRipExportRules) {
	o.Name = s.Name
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolRipExportRulesXml) UnmarshalToObject() (*ProtocolRipExportRules, error) {

	result := &ProtocolRipExportRules{
		Name:   o.Name,
		Metric: o.Metric,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolRipGlobalBfdXml) MarshalFromObject(s ProtocolRipGlobalBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolRipGlobalBfdXml) UnmarshalToObject() (*ProtocolRipGlobalBfd, error) {

	result := &ProtocolRipGlobalBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolRipInterfaceXml) MarshalFromObject(s ProtocolRipInterface) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Authentication = s.Authentication
	o.Mode = s.Mode
	if s.DefaultRoute != nil {
		var obj protocolRipInterfaceDefaultRouteXml
		obj.MarshalFromObject(*s.DefaultRoute)
		o.DefaultRoute = &obj
	}
	if s.Bfd != nil {
		var obj protocolRipInterfaceBfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRipInterfaceXml) UnmarshalToObject() (*ProtocolRipInterface, error) {
	var defaultRouteVal *ProtocolRipInterfaceDefaultRoute
	if o.DefaultRoute != nil {
		obj, err := o.DefaultRoute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultRouteVal = obj
	}
	var bfdVal *ProtocolRipInterfaceBfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}

	result := &ProtocolRipInterface{
		Name:           o.Name,
		Enable:         util.AsBool(o.Enable, nil),
		Authentication: o.Authentication,
		Mode:           o.Mode,
		DefaultRoute:   defaultRouteVal,
		Bfd:            bfdVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *protocolRipInterfaceDefaultRouteXml) MarshalFromObject(s ProtocolRipInterfaceDefaultRoute) {
	if s.Disable != nil {
		var obj protocolRipInterfaceDefaultRouteDisableXml
		obj.MarshalFromObject(*s.Disable)
		o.Disable = &obj
	}
	if s.Advertise != nil {
		var obj protocolRipInterfaceDefaultRouteAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o protocolRipInterfaceDefaultRouteXml) UnmarshalToObject() (*ProtocolRipInterfaceDefaultRoute, error) {
	var disableVal *ProtocolRipInterfaceDefaultRouteDisable
	if o.Disable != nil {
		obj, err := o.Disable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disableVal = obj
	}
	var advertiseVal *ProtocolRipInterfaceDefaultRouteAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &ProtocolRipInterfaceDefaultRoute{
		Disable:   disableVal,
		Advertise: advertiseVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *protocolRipInterfaceDefaultRouteDisableXml) MarshalFromObject(s ProtocolRipInterfaceDefaultRouteDisable) {
	o.Misc = s.Misc
}

func (o protocolRipInterfaceDefaultRouteDisableXml) UnmarshalToObject() (*ProtocolRipInterfaceDefaultRouteDisable, error) {

	result := &ProtocolRipInterfaceDefaultRouteDisable{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolRipInterfaceDefaultRouteAdvertiseXml) MarshalFromObject(s ProtocolRipInterfaceDefaultRouteAdvertise) {
	o.Metric = s.Metric
	o.Misc = s.Misc
}

func (o protocolRipInterfaceDefaultRouteAdvertiseXml) UnmarshalToObject() (*ProtocolRipInterfaceDefaultRouteAdvertise, error) {

	result := &ProtocolRipInterfaceDefaultRouteAdvertise{
		Metric: o.Metric,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolRipInterfaceBfdXml) MarshalFromObject(s ProtocolRipInterfaceBfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o protocolRipInterfaceBfdXml) UnmarshalToObject() (*ProtocolRipInterfaceBfd, error) {

	result := &ProtocolRipInterfaceBfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolRipTimersXml) MarshalFromObject(s ProtocolRipTimers) {
	o.DeleteIntervals = s.DeleteIntervals
	o.ExpireIntervals = s.ExpireIntervals
	o.IntervalSeconds = s.IntervalSeconds
	o.UpdateIntervals = s.UpdateIntervals
	o.Misc = s.Misc
}

func (o protocolRipTimersXml) UnmarshalToObject() (*ProtocolRipTimers, error) {

	result := &ProtocolRipTimers{
		DeleteIntervals: o.DeleteIntervals,
		ExpireIntervals: o.ExpireIntervals,
		IntervalSeconds: o.IntervalSeconds,
		UpdateIntervals: o.UpdateIntervals,
		Misc:            o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "admin_dists" || v == "AdminDists" {
		return e.AdminDists, nil
	}
	if v == "ecmp" || v == "Ecmp" {
		return e.Ecmp, nil
	}
	if v == "interface" || v == "Interface" {
		return e.Interface, nil
	}
	if v == "interface|LENGTH" || v == "Interface|LENGTH" {
		return int64(len(e.Interface)), nil
	}
	if v == "multicast" || v == "Multicast" {
		return e.Multicast, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
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
	if !o.AdminDists.matches(other.AdminDists) {
		return false
	}
	if !o.Ecmp.matches(other.Ecmp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !o.Multicast.matches(other.Multicast) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}

	return true
}

func (o *AdminDists) matches(other *AdminDists) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Ebgp, other.Ebgp) {
		return false
	}
	if !util.Ints64Match(o.Ibgp, other.Ibgp) {
		return false
	}
	if !util.Ints64Match(o.OspfExt, other.OspfExt) {
		return false
	}
	if !util.Ints64Match(o.OspfInt, other.OspfInt) {
		return false
	}
	if !util.Ints64Match(o.Ospfv3Ext, other.Ospfv3Ext) {
		return false
	}
	if !util.Ints64Match(o.Ospfv3Int, other.Ospfv3Int) {
		return false
	}
	if !util.Ints64Match(o.Rip, other.Rip) {
		return false
	}
	if !util.Ints64Match(o.Static, other.Static) {
		return false
	}
	if !util.Ints64Match(o.StaticIpv6, other.StaticIpv6) {
		return false
	}

	return true
}

func (o *Ecmp) matches(other *Ecmp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Algorithm.matches(other.Algorithm) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.MaxPath, other.MaxPath) {
		return false
	}
	if !util.BoolsMatch(o.StrictSourcePath, other.StrictSourcePath) {
		return false
	}
	if !util.BoolsMatch(o.SymmetricReturn, other.SymmetricReturn) {
		return false
	}

	return true
}

func (o *EcmpAlgorithm) matches(other *EcmpAlgorithm) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.BalancedRoundRobin.matches(other.BalancedRoundRobin) {
		return false
	}
	if !o.IpHash.matches(other.IpHash) {
		return false
	}
	if !o.IpModulo.matches(other.IpModulo) {
		return false
	}
	if !o.WeightedRoundRobin.matches(other.WeightedRoundRobin) {
		return false
	}

	return true
}

func (o *EcmpAlgorithmBalancedRoundRobin) matches(other *EcmpAlgorithmBalancedRoundRobin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *EcmpAlgorithmIpHash) matches(other *EcmpAlgorithmIpHash) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.HashSeed, other.HashSeed) {
		return false
	}
	if !util.BoolsMatch(o.SrcOnly, other.SrcOnly) {
		return false
	}
	if !util.BoolsMatch(o.UsePort, other.UsePort) {
		return false
	}

	return true
}

func (o *EcmpAlgorithmIpModulo) matches(other *EcmpAlgorithmIpModulo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *EcmpAlgorithmWeightedRoundRobin) matches(other *EcmpAlgorithmWeightedRoundRobin) bool {
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

func (o *EcmpAlgorithmWeightedRoundRobinInterface) matches(other *EcmpAlgorithmWeightedRoundRobinInterface) bool {
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

func (o *Multicast) matches(other *Multicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.InterfaceGroup) != len(other.InterfaceGroup) {
		return false
	}
	for idx := range o.InterfaceGroup {
		if !o.InterfaceGroup[idx].matches(&other.InterfaceGroup[idx]) {
			return false
		}
	}
	if !util.Ints64Match(o.RouteAgeoutTime, other.RouteAgeoutTime) {
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
	if len(o.SsmAddressSpace) != len(other.SsmAddressSpace) {
		return false
	}
	for idx := range o.SsmAddressSpace {
		if !o.SsmAddressSpace[idx].matches(&other.SsmAddressSpace[idx]) {
			return false
		}
	}

	return true
}

func (o *MulticastInterfaceGroup) matches(other *MulticastInterfaceGroup) bool {
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
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !o.GroupPermission.matches(other.GroupPermission) {
		return false
	}
	if !o.Igmp.matches(other.Igmp) {
		return false
	}
	if !o.Pim.matches(other.Pim) {
		return false
	}

	return true
}

func (o *MulticastInterfaceGroupGroupPermission) matches(other *MulticastInterfaceGroupGroupPermission) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.AnySourceMulticast) != len(other.AnySourceMulticast) {
		return false
	}
	for idx := range o.AnySourceMulticast {
		if !o.AnySourceMulticast[idx].matches(&other.AnySourceMulticast[idx]) {
			return false
		}
	}
	if len(o.SourceSpecificMulticast) != len(other.SourceSpecificMulticast) {
		return false
	}
	for idx := range o.SourceSpecificMulticast {
		if !o.SourceSpecificMulticast[idx].matches(&other.SourceSpecificMulticast[idx]) {
			return false
		}
	}

	return true
}

func (o *MulticastInterfaceGroupGroupPermissionAnySourceMulticast) matches(other *MulticastInterfaceGroupGroupPermissionAnySourceMulticast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.GroupAddress, other.GroupAddress) {
		return false
	}
	if !util.BoolsMatch(o.Included, other.Included) {
		return false
	}

	return true
}

func (o *MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast) matches(other *MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.GroupAddress, other.GroupAddress) {
		return false
	}
	if !util.StringsMatch(o.SourceAddress, other.SourceAddress) {
		return false
	}
	if !util.BoolsMatch(o.Included, other.Included) {
		return false
	}

	return true
}

func (o *MulticastInterfaceGroupIgmp) matches(other *MulticastInterfaceGroupIgmp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Version, other.Version) {
		return false
	}
	if !util.FloatsMatch(o.MaxQueryResponseTime, other.MaxQueryResponseTime) {
		return false
	}
	if !util.Ints64Match(o.QueryInterval, other.QueryInterval) {
		return false
	}
	if !util.FloatsMatch(o.LastMemberQueryInterval, other.LastMemberQueryInterval) {
		return false
	}
	if !util.BoolsMatch(o.ImmediateLeave, other.ImmediateLeave) {
		return false
	}
	if !util.StringsMatch(o.Robustness, other.Robustness) {
		return false
	}
	if !util.StringsMatch(o.MaxGroups, other.MaxGroups) {
		return false
	}
	if !util.StringsMatch(o.MaxSources, other.MaxSources) {
		return false
	}
	if !util.BoolsMatch(o.RouterAlertPolicing, other.RouterAlertPolicing) {
		return false
	}

	return true
}

func (o *MulticastInterfaceGroupPim) matches(other *MulticastInterfaceGroupPim) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
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
	if !util.Ints64Match(o.DrPriority, other.DrPriority) {
		return false
	}
	if !util.BoolsMatch(o.BsrBorder, other.BsrBorder) {
		return false
	}
	if len(o.AllowedNeighbors) != len(other.AllowedNeighbors) {
		return false
	}
	for idx := range o.AllowedNeighbors {
		if !o.AllowedNeighbors[idx].matches(&other.AllowedNeighbors[idx]) {
			return false
		}
	}

	return true
}

func (o *MulticastInterfaceGroupPimAllowedNeighbors) matches(other *MulticastInterfaceGroupPimAllowedNeighbors) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *MulticastRp) matches(other *MulticastRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
	if !o.LocalRp.matches(other.LocalRp) {
		return false
	}

	return true
}

func (o *MulticastRpExternalRp) matches(other *MulticastRpExternalRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.GroupAddresses, other.GroupAddresses) {
		return false
	}
	if !util.BoolsMatch(o.Override, other.Override) {
		return false
	}

	return true
}

func (o *MulticastRpLocalRp) matches(other *MulticastRpLocalRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.CandidateRp.matches(other.CandidateRp) {
		return false
	}
	if !o.StaticRp.matches(other.StaticRp) {
		return false
	}

	return true
}

func (o *MulticastRpLocalRpCandidateRp) matches(other *MulticastRpLocalRpCandidateRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.Ints64Match(o.AdvertisementInterval, other.AdvertisementInterval) {
		return false
	}
	if !util.OrderedListsMatch[string](o.GroupAddresses, other.GroupAddresses) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}

	return true
}

func (o *MulticastRpLocalRpStaticRp) matches(other *MulticastRpLocalRpStaticRp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.OrderedListsMatch[string](o.GroupAddresses, other.GroupAddresses) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.BoolsMatch(o.Override, other.Override) {
		return false
	}

	return true
}

func (o *MulticastSptThreshold) matches(other *MulticastSptThreshold) bool {
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

func (o *MulticastSsmAddressSpace) matches(other *MulticastSsmAddressSpace) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.GroupAddress, other.GroupAddress) {
		return false
	}
	if !util.BoolsMatch(o.Included, other.Included) {
		return false
	}

	return true
}

func (o *Protocol) matches(other *Protocol) bool {
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
	if len(o.RedistProfile) != len(other.RedistProfile) {
		return false
	}
	for idx := range o.RedistProfile {
		if !o.RedistProfile[idx].matches(&other.RedistProfile[idx]) {
			return false
		}
	}
	if len(o.RedistProfileIpv6) != len(other.RedistProfileIpv6) {
		return false
	}
	for idx := range o.RedistProfileIpv6 {
		if !o.RedistProfileIpv6[idx].matches(&other.RedistProfileIpv6[idx]) {
			return false
		}
	}
	if !o.Rip.matches(other.Rip) {
		return false
	}

	return true
}

func (o *ProtocolBgp) matches(other *ProtocolBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowRedistDefaultRoute, other.AllowRedistDefaultRoute) {
		return false
	}
	if len(o.AuthProfile) != len(other.AuthProfile) {
		return false
	}
	for idx := range o.AuthProfile {
		if !o.AuthProfile[idx].matches(&other.AuthProfile[idx]) {
			return false
		}
	}
	if len(o.DampeningProfile) != len(other.DampeningProfile) {
		return false
	}
	for idx := range o.DampeningProfile {
		if !o.DampeningProfile[idx].matches(&other.DampeningProfile[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.EcmpMultiAs, other.EcmpMultiAs) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.EnforceFirstAs, other.EnforceFirstAs) {
		return false
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !util.BoolsMatch(o.InstallRoute, other.InstallRoute) {
		return false
	}
	if !util.StringsMatch(o.LocalAs, other.LocalAs) {
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
	if !o.Policy.matches(other.Policy) {
		return false
	}
	if len(o.RedistRules) != len(other.RedistRules) {
		return false
	}
	for idx := range o.RedistRules {
		if !o.RedistRules[idx].matches(&other.RedistRules[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.RejectDefaultRoute, other.RejectDefaultRoute) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !o.RoutingOptions.matches(other.RoutingOptions) {
		return false
	}

	return true
}

func (o *ProtocolBgpAuthProfile) matches(other *ProtocolBgpAuthProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Secret, other.Secret) {
		return false
	}

	return true
}

func (o *ProtocolBgpDampeningProfile) matches(other *ProtocolBgpDampeningProfile) bool {
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
	if !util.FloatsMatch(o.Cutoff, other.Cutoff) {
		return false
	}
	if !util.FloatsMatch(o.Reuse, other.Reuse) {
		return false
	}
	if !util.Ints64Match(o.MaxHoldTime, other.MaxHoldTime) {
		return false
	}
	if !util.Ints64Match(o.DecayHalfLifeReachable, other.DecayHalfLifeReachable) {
		return false
	}
	if !util.Ints64Match(o.DecayHalfLifeUnreachable, other.DecayHalfLifeUnreachable) {
		return false
	}

	return true
}

func (o *ProtocolBgpGlobalBfd) matches(other *ProtocolBgpGlobalBfd) bool {
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

func (o *ProtocolBgpPeerGroup) matches(other *ProtocolBgpPeerGroup) bool {
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
	if !util.BoolsMatch(o.AggregatedConfedAsPath, other.AggregatedConfedAsPath) {
		return false
	}
	if !util.BoolsMatch(o.SoftResetWithStoredInfo, other.SoftResetWithStoredInfo) {
		return false
	}
	if !o.Type.matches(other.Type) {
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

func (o *ProtocolBgpPeerGroupType) matches(other *ProtocolBgpPeerGroupType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ibgp.matches(other.Ibgp) {
		return false
	}
	if !o.EbgpConfed.matches(other.EbgpConfed) {
		return false
	}
	if !o.IbgpConfed.matches(other.IbgpConfed) {
		return false
	}
	if !o.Ebgp.matches(other.Ebgp) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupTypeIbgp) matches(other *ProtocolBgpPeerGroupTypeIbgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ExportNexthop, other.ExportNexthop) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupTypeEbgpConfed) matches(other *ProtocolBgpPeerGroupTypeEbgpConfed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ExportNexthop, other.ExportNexthop) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupTypeIbgpConfed) matches(other *ProtocolBgpPeerGroupTypeIbgpConfed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ExportNexthop, other.ExportNexthop) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupTypeEbgp) matches(other *ProtocolBgpPeerGroupTypeEbgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ImportNexthop, other.ImportNexthop) {
		return false
	}
	if !util.StringsMatch(o.ExportNexthop, other.ExportNexthop) {
		return false
	}
	if !util.BoolsMatch(o.RemovePrivateAs, other.RemovePrivateAs) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupPeer) matches(other *ProtocolBgpPeerGroupPeer) bool {
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
	if !util.BoolsMatch(o.EnableMpBgp, other.EnableMpBgp) {
		return false
	}
	if !util.StringsMatch(o.AddressFamilyIdentifier, other.AddressFamilyIdentifier) {
		return false
	}
	if !util.BoolsMatch(o.EnableSenderSideLoopDetection, other.EnableSenderSideLoopDetection) {
		return false
	}
	if !util.StringsMatch(o.ReflectorClient, other.ReflectorClient) {
		return false
	}
	if !util.StringsMatch(o.PeeringType, other.PeeringType) {
		return false
	}
	if !util.StringsMatch(o.MaxPrefixes, other.MaxPrefixes) {
		return false
	}
	if !o.SubsequentAddressFamilyIdentifier.matches(other.SubsequentAddressFamilyIdentifier) {
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

func (o *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier) matches(other *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Unicast, other.Unicast) {
		return false
	}
	if !util.BoolsMatch(o.Multicast, other.Multicast) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupPeerLocalAddress) matches(other *ProtocolBgpPeerGroupPeerLocalAddress) bool {
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

func (o *ProtocolBgpPeerGroupPeerPeerAddress) matches(other *ProtocolBgpPeerGroupPeerPeerAddress) bool {
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

func (o *ProtocolBgpPeerGroupPeerConnectionOptions) matches(other *ProtocolBgpPeerGroupPeerConnectionOptions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.KeepAliveInterval, other.KeepAliveInterval) {
		return false
	}
	if !util.Ints64Match(o.MinRouteAdvInterval, other.MinRouteAdvInterval) {
		return false
	}
	if !util.Ints64Match(o.Multihop, other.Multihop) {
		return false
	}
	if !util.Ints64Match(o.OpenDelayTime, other.OpenDelayTime) {
		return false
	}
	if !util.StringsMatch(o.HoldTime, other.HoldTime) {
		return false
	}
	if !util.Ints64Match(o.IdleHoldTime, other.IdleHoldTime) {
		return false
	}
	if !o.IncomingBgpConnection.matches(other.IncomingBgpConnection) {
		return false
	}
	if !o.OutgoingBgpConnection.matches(other.OutgoingBgpConnection) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection) matches(other *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.RemotePort, other.RemotePort) {
		return false
	}
	if !util.BoolsMatch(o.Allow, other.Allow) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection) matches(other *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPort, other.LocalPort) {
		return false
	}
	if !util.BoolsMatch(o.Allow, other.Allow) {
		return false
	}

	return true
}

func (o *ProtocolBgpPeerGroupPeerBfd) matches(other *ProtocolBgpPeerGroupPeerBfd) bool {
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

func (o *ProtocolBgpPolicy) matches(other *ProtocolBgpPolicy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Aggregation.matches(other.Aggregation) {
		return false
	}
	if !o.ConditionalAdvertisement.matches(other.ConditionalAdvertisement) {
		return false
	}
	if !o.Export.matches(other.Export) {
		return false
	}
	if !o.Import.matches(other.Import) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregation) matches(other *ProtocolBgpPolicyAggregation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Address) != len(other.Address) {
		return false
	}
	for idx := range o.Address {
		if !o.Address[idx].matches(&other.Address[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddress) matches(other *ProtocolBgpPolicyAggregationAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Prefix, other.Prefix) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.Summary, other.Summary) {
		return false
	}
	if !util.BoolsMatch(o.AsSet, other.AsSet) {
		return false
	}
	if !o.AggregateRouteAttributes.matches(other.AggregateRouteAttributes) {
		return false
	}
	if len(o.SuppressFilters) != len(other.SuppressFilters) {
		return false
	}
	for idx := range o.SuppressFilters {
		if !o.SuppressFilters[idx].matches(&other.SuppressFilters[idx]) {
			return false
		}
	}
	if len(o.AdvertiseFilters) != len(other.AdvertiseFilters) {
		return false
	}
	for idx := range o.AdvertiseFilters {
		if !o.AdvertiseFilters[idx].matches(&other.AdvertiseFilters[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPreference, other.LocalPreference) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if !util.Ints64Match(o.Weight, other.Weight) {
		return false
	}
	if !util.StringsMatch(o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.StringsMatch(o.Origin, other.Origin) {
		return false
	}
	if !util.Ints64Match(o.AsPathLimit, other.AsPathLimit) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !util.Ints64Match(o.Prepend, other.Prepend) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll) matches(other *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFilters) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFilters) bool {
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
	if !o.Match.matches(other.Match) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Exact, other.Exact) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity) matches(other *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFilters) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFilters) bool {
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
	if !o.Match.matches(other.Match) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Exact, other.Exact) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity) matches(other *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisement) matches(other *ProtocolBgpPolicyConditionalAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Policy) != len(other.Policy) {
		return false
	}
	for idx := range o.Policy {
		if !o.Policy[idx].matches(&other.Policy[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicy) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicy) bool {
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
	if !util.OrderedListsMatch[string](o.UsedBy, other.UsedBy) {
		return false
	}
	if len(o.NonExistFilters) != len(other.NonExistFilters) {
		return false
	}
	for idx := range o.NonExistFilters {
		if !o.NonExistFilters[idx].matches(&other.NonExistFilters[idx]) {
			return false
		}
	}
	if len(o.AdvertiseFilters) != len(other.AdvertiseFilters) {
		return false
	}
	for idx := range o.AdvertiseFilters {
		if !o.AdvertiseFilters[idx].matches(&other.AdvertiseFilters[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters) bool {
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
	if !o.Match.matches(other.Match) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters) bool {
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
	if !o.Match.matches(other.Match) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity) matches(other *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExport) matches(other *ProtocolBgpPolicyExport) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Rules) != len(other.Rules) {
		return false
	}
	for idx := range o.Rules {
		if !o.Rules[idx].matches(&other.Rules[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyExportRules) matches(other *ProtocolBgpPolicyExportRules) bool {
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
	if !util.OrderedListsMatch[string](o.UsedBy, other.UsedBy) {
		return false
	}
	if !o.Match.matches(other.Match) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesMatch) matches(other *ProtocolBgpPolicyExportRulesMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesMatchAddressPrefix) matches(other *ProtocolBgpPolicyExportRulesMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Exact, other.Exact) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesMatchAsPath) matches(other *ProtocolBgpPolicyExportRulesMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesMatchCommunity) matches(other *ProtocolBgpPolicyExportRulesMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesMatchExtendedCommunity) matches(other *ProtocolBgpPolicyExportRulesMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesAction) matches(other *ProtocolBgpPolicyExportRulesAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Deny.matches(other.Deny) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionDeny) matches(other *ProtocolBgpPolicyExportRulesActionDeny) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllow) matches(other *ProtocolBgpPolicyExportRulesActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Update.matches(other.Update) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdate) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPreference, other.LocalPreference) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if !util.StringsMatch(o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.StringsMatch(o.Origin, other.Origin) {
		return false
	}
	if !util.Ints64Match(o.AsPathLimit, other.AsPathLimit) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Remove.matches(other.Remove) {
		return false
	}
	if !util.Ints64Match(o.Prepend, other.Prepend) {
		return false
	}
	if !util.Ints64Match(o.RemoveAndPrepend, other.RemoveAndPrepend) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll) matches(other *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImport) matches(other *ProtocolBgpPolicyImport) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Rules) != len(other.Rules) {
		return false
	}
	for idx := range o.Rules {
		if !o.Rules[idx].matches(&other.Rules[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolBgpPolicyImportRules) matches(other *ProtocolBgpPolicyImportRules) bool {
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
	if !util.OrderedListsMatch[string](o.UsedBy, other.UsedBy) {
		return false
	}
	if !o.Match.matches(other.Match) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesMatch) matches(other *ProtocolBgpPolicyImportRulesMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if len(o.AddressPrefix) != len(other.AddressPrefix) {
		return false
	}
	for idx := range o.AddressPrefix {
		if !o.AddressPrefix[idx].matches(&other.AddressPrefix[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FromPeer, other.FromPeer) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesMatchAddressPrefix) matches(other *ProtocolBgpPolicyImportRulesMatchAddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Exact, other.Exact) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesMatchAsPath) matches(other *ProtocolBgpPolicyImportRulesMatchAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesMatchCommunity) matches(other *ProtocolBgpPolicyImportRulesMatchCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesMatchExtendedCommunity) matches(other *ProtocolBgpPolicyImportRulesMatchExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Regex, other.Regex) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesAction) matches(other *ProtocolBgpPolicyImportRulesAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Deny.matches(other.Deny) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionDeny) matches(other *ProtocolBgpPolicyImportRulesActionDeny) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllow) matches(other *ProtocolBgpPolicyImportRulesActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Dampening, other.Dampening) {
		return false
	}
	if !o.Update.matches(other.Update) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdate) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPreference, other.LocalPreference) {
		return false
	}
	if !util.Ints64Match(o.Med, other.Med) {
		return false
	}
	if !util.Ints64Match(o.Weight, other.Weight) {
		return false
	}
	if !util.StringsMatch(o.Nexthop, other.Nexthop) {
		return false
	}
	if !util.StringsMatch(o.Origin, other.Origin) {
		return false
	}
	if !util.Ints64Match(o.AsPathLimit, other.AsPathLimit) {
		return false
	}
	if !o.AsPath.matches(other.AsPath) {
		return false
	}
	if !o.Community.matches(other.Community) {
		return false
	}
	if !o.ExtendedCommunity.matches(other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Remove.matches(other.Remove) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RemoveAll.matches(other.RemoveAll) {
		return false
	}
	if !util.StringsMatch(o.RemoveRegex, other.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Append, other.Append) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Overwrite, other.Overwrite) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll) matches(other *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolBgpRedistRules) matches(other *ProtocolBgpRedistRules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.AddressFamilyIdentifier, other.AddressFamilyIdentifier) {
		return false
	}
	if !util.StringsMatch(o.RouteTable, other.RouteTable) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.SetOrigin, other.SetOrigin) {
		return false
	}
	if !util.Ints64Match(o.SetMed, other.SetMed) {
		return false
	}
	if !util.Ints64Match(o.SetLocalPreference, other.SetLocalPreference) {
		return false
	}
	if !util.Ints64Match(o.SetAsPathLimit, other.SetAsPathLimit) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SetCommunity, other.SetCommunity) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SetExtendedCommunity, other.SetExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolBgpRoutingOptions) matches(other *ProtocolBgpRoutingOptions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Aggregate.matches(other.Aggregate) {
		return false
	}
	if !util.StringsMatch(o.AsFormat, other.AsFormat) {
		return false
	}
	if !util.StringsMatch(o.ConfederationMemberAs, other.ConfederationMemberAs) {
		return false
	}
	if !util.Ints64Match(o.DefaultLocalPreference, other.DefaultLocalPreference) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if !o.Med.matches(other.Med) {
		return false
	}
	if !util.StringsMatch(o.ReflectorClusterId, other.ReflectorClusterId) {
		return false
	}

	return true
}

func (o *ProtocolBgpRoutingOptionsAggregate) matches(other *ProtocolBgpRoutingOptionsAggregate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AggregateMed, other.AggregateMed) {
		return false
	}

	return true
}

func (o *ProtocolBgpRoutingOptionsGracefulRestart) matches(other *ProtocolBgpRoutingOptionsGracefulRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.LocalRestartTime, other.LocalRestartTime) {
		return false
	}
	if !util.Ints64Match(o.MaxPeerRestartTime, other.MaxPeerRestartTime) {
		return false
	}
	if !util.Ints64Match(o.StaleRouteTime, other.StaleRouteTime) {
		return false
	}

	return true
}

func (o *ProtocolBgpRoutingOptionsMed) matches(other *ProtocolBgpRoutingOptionsMed) bool {
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

func (o *ProtocolOspf) matches(other *ProtocolOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowRedistDefaultRoute, other.AllowRedistDefaultRoute) {
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
	if len(o.AuthProfile) != len(other.AuthProfile) {
		return false
	}
	for idx := range o.AuthProfile {
		if !o.AuthProfile[idx].matches(&other.AuthProfile[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.ExportRules) != len(other.ExportRules) {
		return false
	}
	for idx := range o.ExportRules {
		if !o.ExportRules[idx].matches(&other.ExportRules[idx]) {
			return false
		}
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if !util.BoolsMatch(o.RejectDefaultRoute, other.RejectDefaultRoute) {
		return false
	}
	if !util.BoolsMatch(o.Rfc1583, other.Rfc1583) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !o.Timers.matches(other.Timers) {
		return false
	}

	return true
}

func (o *ProtocolOspfArea) matches(other *ProtocolOspfArea) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
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

func (o *ProtocolOspfAreaType) matches(other *ProtocolOspfAreaType) bool {
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

func (o *ProtocolOspfAreaTypeNormal) matches(other *ProtocolOspfAreaTypeNormal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeStub) matches(other *ProtocolOspfAreaTypeStub) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AcceptSummary, other.AcceptSummary) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeStubDefaultRoute) matches(other *ProtocolOspfAreaTypeStubDefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disable.matches(other.Disable) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeStubDefaultRouteDisable) matches(other *ProtocolOspfAreaTypeStubDefaultRouteDisable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeStubDefaultRouteAdvertise) matches(other *ProtocolOspfAreaTypeStubDefaultRouteAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssa) matches(other *ProtocolOspfAreaTypeNssa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AcceptSummary, other.AcceptSummary) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
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

func (o *ProtocolOspfAreaTypeNssaDefaultRoute) matches(other *ProtocolOspfAreaTypeNssaDefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disable.matches(other.Disable) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssaDefaultRouteDisable) matches(other *ProtocolOspfAreaTypeNssaDefaultRouteDisable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise) matches(other *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssaNssaExtRange) matches(other *ProtocolOspfAreaTypeNssaNssaExtRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}
	if !o.Suppress.matches(other.Suppress) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise) matches(other *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress) matches(other *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaRange) matches(other *ProtocolOspfAreaRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}
	if !o.Suppress.matches(other.Suppress) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaRangeAdvertise) matches(other *ProtocolOspfAreaRangeAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaRangeSuppress) matches(other *ProtocolOspfAreaRangeSuppress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterface) matches(other *ProtocolOspfAreaInterface) bool {
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
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.DeadCounts, other.DeadCounts) {
		return false
	}
	if !util.Ints64Match(o.RetransmitInterval, other.RetransmitInterval) {
		return false
	}
	if !util.Ints64Match(o.TransitDelay, other.TransitDelay) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.Ints64Match(o.GrDelay, other.GrDelay) {
		return false
	}
	if !o.LinkType.matches(other.LinkType) {
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
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterfaceLinkType) matches(other *ProtocolOspfAreaInterfaceLinkType) bool {
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

func (o *ProtocolOspfAreaInterfaceLinkTypeBroadcast) matches(other *ProtocolOspfAreaInterfaceLinkTypeBroadcast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterfaceLinkTypeP2p) matches(other *ProtocolOspfAreaInterfaceLinkTypeP2p) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterfaceLinkTypeP2mp) matches(other *ProtocolOspfAreaInterfaceLinkTypeP2mp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterfaceNeighbor) matches(other *ProtocolOspfAreaInterfaceNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *ProtocolOspfAreaInterfaceBfd) matches(other *ProtocolOspfAreaInterfaceBfd) bool {
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

func (o *ProtocolOspfAreaVirtualLink) matches(other *ProtocolOspfAreaVirtualLink) bool {
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
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.DeadCounts, other.DeadCounts) {
		return false
	}
	if !util.Ints64Match(o.RetransmitInterval, other.RetransmitInterval) {
		return false
	}
	if !util.Ints64Match(o.TransitDelay, other.TransitDelay) {
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

func (o *ProtocolOspfAreaVirtualLinkBfd) matches(other *ProtocolOspfAreaVirtualLinkBfd) bool {
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

func (o *ProtocolOspfAuthProfile) matches(other *ProtocolOspfAuthProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if len(o.Md5) != len(other.Md5) {
		return false
	}
	for idx := range o.Md5 {
		if !o.Md5[idx].matches(&other.Md5[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolOspfAuthProfileMd5) matches(other *ProtocolOspfAuthProfileMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}
	if !util.BoolsMatch(o.Preferred, other.Preferred) {
		return false
	}

	return true
}

func (o *ProtocolOspfExportRules) matches(other *ProtocolOspfExportRules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.NewPathType, other.NewPathType) {
		return false
	}
	if !util.StringsMatch(o.NewTag, other.NewTag) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolOspfGlobalBfd) matches(other *ProtocolOspfGlobalBfd) bool {
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

func (o *ProtocolOspfGracefulRestart) matches(other *ProtocolOspfGracefulRestart) bool {
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
	if !util.Ints64Match(o.MaxNeighborRestartTime, other.MaxNeighborRestartTime) {
		return false
	}
	if !util.BoolsMatch(o.StrictLSAChecking, other.StrictLSAChecking) {
		return false
	}

	return true
}

func (o *ProtocolOspfTimers) matches(other *ProtocolOspfTimers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.FloatsMatch(o.LsaInterval, other.LsaInterval) {
		return false
	}
	if !util.FloatsMatch(o.SpfCalculationDelay, other.SpfCalculationDelay) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3) matches(other *ProtocolOspfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowRedistDefaultRoute, other.AllowRedistDefaultRoute) {
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
	if len(o.AuthProfile) != len(other.AuthProfile) {
		return false
	}
	for idx := range o.AuthProfile {
		if !o.AuthProfile[idx].matches(&other.AuthProfile[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.DisableTransitTraffic, other.DisableTransitTraffic) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.ExportRules) != len(other.ExportRules) {
		return false
	}
	for idx := range o.ExportRules {
		if !o.ExportRules[idx].matches(&other.ExportRules[idx]) {
			return false
		}
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
		return false
	}
	if !o.GracefulRestart.matches(other.GracefulRestart) {
		return false
	}
	if !util.BoolsMatch(o.RejectDefaultRoute, other.RejectDefaultRoute) {
		return false
	}
	if !util.StringsMatch(o.RouterId, other.RouterId) {
		return false
	}
	if !o.Timers.matches(other.Timers) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3Area) matches(other *ProtocolOspfv3Area) bool {
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

func (o *ProtocolOspfv3AreaType) matches(other *ProtocolOspfv3AreaType) bool {
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

func (o *ProtocolOspfv3AreaTypeNormal) matches(other *ProtocolOspfv3AreaTypeNormal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeStub) matches(other *ProtocolOspfv3AreaTypeStub) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AcceptSummary, other.AcceptSummary) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeStubDefaultRoute) matches(other *ProtocolOspfv3AreaTypeStubDefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disable.matches(other.Disable) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeStubDefaultRouteDisable) matches(other *ProtocolOspfv3AreaTypeStubDefaultRouteDisable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise) matches(other *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssa) matches(other *ProtocolOspfv3AreaTypeNssa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AcceptSummary, other.AcceptSummary) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
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

func (o *ProtocolOspfv3AreaTypeNssaDefaultRoute) matches(other *ProtocolOspfv3AreaTypeNssaDefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disable.matches(other.Disable) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable) matches(other *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise) matches(other *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssaNssaExtRange) matches(other *ProtocolOspfv3AreaTypeNssaNssaExtRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}
	if !o.Suppress.matches(other.Suppress) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise) matches(other *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress) matches(other *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaRange) matches(other *ProtocolOspfv3AreaRange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}
	if !o.Suppress.matches(other.Suppress) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaRangeAdvertise) matches(other *ProtocolOspfv3AreaRangeAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaRangeSuppress) matches(other *ProtocolOspfv3AreaRangeSuppress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterface) matches(other *ProtocolOspfv3AreaInterface) bool {
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
	if !util.Ints64Match(o.InstanceId, other.InstanceId) {
		return false
	}
	if !util.BoolsMatch(o.Passive, other.Passive) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.DeadCounts, other.DeadCounts) {
		return false
	}
	if !util.Ints64Match(o.RetransmitInterval, other.RetransmitInterval) {
		return false
	}
	if !util.Ints64Match(o.TransitDelay, other.TransitDelay) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.Ints64Match(o.GrDelay, other.GrDelay) {
		return false
	}
	if !o.LinkType.matches(other.LinkType) {
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
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterfaceLinkType) matches(other *ProtocolOspfv3AreaInterfaceLinkType) bool {
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

func (o *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast) matches(other *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterfaceLinkTypeP2p) matches(other *ProtocolOspfv3AreaInterfaceLinkTypeP2p) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterfaceLinkTypeP2mp) matches(other *ProtocolOspfv3AreaInterfaceLinkTypeP2mp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterfaceNeighbor) matches(other *ProtocolOspfv3AreaInterfaceNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AreaInterfaceBfd) matches(other *ProtocolOspfv3AreaInterfaceBfd) bool {
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

func (o *ProtocolOspfv3AreaVirtualLink) matches(other *ProtocolOspfv3AreaVirtualLink) bool {
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
	if !util.Ints64Match(o.HelloInterval, other.HelloInterval) {
		return false
	}
	if !util.Ints64Match(o.DeadCounts, other.DeadCounts) {
		return false
	}
	if !util.Ints64Match(o.RetransmitInterval, other.RetransmitInterval) {
		return false
	}
	if !util.Ints64Match(o.TransitDelay, other.TransitDelay) {
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

func (o *ProtocolOspfv3AreaVirtualLinkBfd) matches(other *ProtocolOspfv3AreaVirtualLinkBfd) bool {
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

func (o *ProtocolOspfv3AuthProfile) matches(other *ProtocolOspfv3AuthProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Spi, other.Spi) {
		return false
	}
	if !o.Esp.matches(other.Esp) {
		return false
	}
	if !o.Ah.matches(other.Ah) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEsp) matches(other *ProtocolOspfv3AuthProfileEsp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Authentication.matches(other.Authentication) {
		return false
	}
	if !o.Encryption.matches(other.Encryption) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthentication) matches(other *ProtocolOspfv3AuthProfileEspAuthentication) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationMd5) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationSha1) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationSha256) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationSha384) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationSha512) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspAuthenticationNone) matches(other *ProtocolOspfv3AuthProfileEspAuthenticationNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileEspEncryption) matches(other *ProtocolOspfv3AuthProfileEspEncryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Algorithm, other.Algorithm) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAh) matches(other *ProtocolOspfv3AuthProfileAh) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAhMd5) matches(other *ProtocolOspfv3AuthProfileAhMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAhSha1) matches(other *ProtocolOspfv3AuthProfileAhSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAhSha256) matches(other *ProtocolOspfv3AuthProfileAhSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAhSha384) matches(other *ProtocolOspfv3AuthProfileAhSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3AuthProfileAhSha512) matches(other *ProtocolOspfv3AuthProfileAhSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3ExportRules) matches(other *ProtocolOspfv3ExportRules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.NewPathType, other.NewPathType) {
		return false
	}
	if !util.StringsMatch(o.NewTag, other.NewTag) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3GlobalBfd) matches(other *ProtocolOspfv3GlobalBfd) bool {
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

func (o *ProtocolOspfv3GracefulRestart) matches(other *ProtocolOspfv3GracefulRestart) bool {
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
	if !util.Ints64Match(o.MaxNeighborRestartTime, other.MaxNeighborRestartTime) {
		return false
	}
	if !util.BoolsMatch(o.StrictLSAChecking, other.StrictLSAChecking) {
		return false
	}

	return true
}

func (o *ProtocolOspfv3Timers) matches(other *ProtocolOspfv3Timers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.FloatsMatch(o.LsaInterval, other.LsaInterval) {
		return false
	}
	if !util.FloatsMatch(o.SpfCalculationDelay, other.SpfCalculationDelay) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfile) matches(other *ProtocolRedistProfile) bool {
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
	if !o.Filter.matches(other.Filter) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileFilter) matches(other *ProtocolRedistProfileFilter) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Type, other.Type) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Destination, other.Destination) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !o.Ospf.matches(other.Ospf) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileFilterOspf) matches(other *ProtocolRedistProfileFilterOspf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.PathType, other.PathType) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Area, other.Area) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileFilterBgp) matches(other *ProtocolRedistProfileFilterBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Community, other.Community) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileAction) matches(other *ProtocolRedistProfileAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.NoRedist.matches(other.NoRedist) {
		return false
	}
	if !o.Redist.matches(other.Redist) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileActionNoRedist) matches(other *ProtocolRedistProfileActionNoRedist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileActionRedist) matches(other *ProtocolRedistProfileActionRedist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6) matches(other *ProtocolRedistProfileIpv6) bool {
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
	if !o.Filter.matches(other.Filter) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6Filter) matches(other *ProtocolRedistProfileIpv6Filter) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Type, other.Type) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Destination, other.Destination) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Nexthop, other.Nexthop) {
		return false
	}
	if !o.Ospfv3.matches(other.Ospfv3) {
		return false
	}
	if !o.Bgp.matches(other.Bgp) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6FilterOspfv3) matches(other *ProtocolRedistProfileIpv6FilterOspfv3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.PathType, other.PathType) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Area, other.Area) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6FilterBgp) matches(other *ProtocolRedistProfileIpv6FilterBgp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Community, other.Community) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExtendedCommunity, other.ExtendedCommunity) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6Action) matches(other *ProtocolRedistProfileIpv6Action) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.NoRedist.matches(other.NoRedist) {
		return false
	}
	if !o.Redist.matches(other.Redist) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6ActionNoRedist) matches(other *ProtocolRedistProfileIpv6ActionNoRedist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolRedistProfileIpv6ActionRedist) matches(other *ProtocolRedistProfileIpv6ActionRedist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolRip) matches(other *ProtocolRip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowRedistDefaultRoute, other.AllowRedistDefaultRoute) {
		return false
	}
	if len(o.AuthProfile) != len(other.AuthProfile) {
		return false
	}
	for idx := range o.AuthProfile {
		if !o.AuthProfile[idx].matches(&other.AuthProfile[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.ExportRules) != len(other.ExportRules) {
		return false
	}
	for idx := range o.ExportRules {
		if !o.ExportRules[idx].matches(&other.ExportRules[idx]) {
			return false
		}
	}
	if !o.GlobalBfd.matches(other.GlobalBfd) {
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
	if !util.BoolsMatch(o.RejectDefaultRoute, other.RejectDefaultRoute) {
		return false
	}
	if !o.Timers.matches(other.Timers) {
		return false
	}

	return true
}

func (o *ProtocolRipAuthProfile) matches(other *ProtocolRipAuthProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if len(o.Md5) != len(other.Md5) {
		return false
	}
	for idx := range o.Md5 {
		if !o.Md5[idx].matches(&other.Md5[idx]) {
			return false
		}
	}

	return true
}

func (o *ProtocolRipAuthProfileMd5) matches(other *ProtocolRipAuthProfileMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}
	if !util.BoolsMatch(o.Preferred, other.Preferred) {
		return false
	}

	return true
}

func (o *ProtocolRipExportRules) matches(other *ProtocolRipExportRules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolRipGlobalBfd) matches(other *ProtocolRipGlobalBfd) bool {
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

func (o *ProtocolRipInterface) matches(other *ProtocolRipInterface) bool {
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
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !o.DefaultRoute.matches(other.DefaultRoute) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}

	return true
}

func (o *ProtocolRipInterfaceDefaultRoute) matches(other *ProtocolRipInterfaceDefaultRoute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disable.matches(other.Disable) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *ProtocolRipInterfaceDefaultRouteDisable) matches(other *ProtocolRipInterfaceDefaultRouteDisable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolRipInterfaceDefaultRouteAdvertise) matches(other *ProtocolRipInterfaceDefaultRouteAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}

	return true
}

func (o *ProtocolRipInterfaceBfd) matches(other *ProtocolRipInterfaceBfd) bool {
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

func (o *ProtocolRipTimers) matches(other *ProtocolRipTimers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.DeleteIntervals, other.DeleteIntervals) {
		return false
	}
	if !util.Ints64Match(o.ExpireIntervals, other.ExpireIntervals) {
		return false
	}
	if !util.Ints64Match(o.IntervalSeconds, other.IntervalSeconds) {
		return false
	}
	if !util.Ints64Match(o.UpdateIntervals, other.UpdateIntervals) {
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
