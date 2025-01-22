package virtual_router

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
	Suffix = []string{"network", "virtual-router"}
)

type Entry struct {
	Name         string
	AdminDists   *AdminDists
	Ecmp         *Ecmp
	Interface    []string
	Multicast    *Multicast
	Protocol     *Protocol
	RoutingTable *RoutingTable

	Misc map[string][]generic.Xml
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
}
type Ecmp struct {
	Algorithm        *EcmpAlgorithm
	Enable           *bool
	MaxPath          *int64
	StrictSourcePath *bool
	SymmetricReturn  *bool
}
type EcmpAlgorithm struct {
	BalancedRoundRobin *EcmpAlgorithmBalancedRoundRobin
	IpHash             *EcmpAlgorithmIpHash
	IpModulo           *EcmpAlgorithmIpModulo
	WeightedRoundRobin *EcmpAlgorithmWeightedRoundRobin
}
type EcmpAlgorithmBalancedRoundRobin struct {
}
type EcmpAlgorithmIpHash struct {
	HashSeed *int64
	SrcOnly  *bool
	UsePort  *bool
}
type EcmpAlgorithmIpModulo struct {
}
type EcmpAlgorithmWeightedRoundRobin struct {
	Interface []EcmpAlgorithmWeightedRoundRobinInterface
}
type EcmpAlgorithmWeightedRoundRobinInterface struct {
	Name   string
	Weight *int64
}
type Multicast struct {
	Enable          *bool
	InterfaceGroup  []MulticastInterfaceGroup
	RouteAgeoutTime *int64
	Rp              *MulticastRp
	SptThreshold    []MulticastSptThreshold
	SsmAddressSpace []MulticastSsmAddressSpace
}
type MulticastInterfaceGroup struct {
	Description     *string
	GroupPermission *MulticastInterfaceGroupGroupPermission
	Igmp            *MulticastInterfaceGroupIgmp
	Interface       []string
	Name            string
	Pim             *MulticastInterfaceGroupPim
}
type MulticastInterfaceGroupGroupPermission struct {
	AnySourceMulticast      []MulticastInterfaceGroupGroupPermissionAnySourceMulticast
	SourceSpecificMulticast []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast
}
type MulticastInterfaceGroupGroupPermissionAnySourceMulticast struct {
	GroupAddress *string
	Included     *bool
	Name         string
}
type MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast struct {
	GroupAddress  *string
	Included      *bool
	Name          string
	SourceAddress *string
}
type MulticastInterfaceGroupIgmp struct {
	Enable                  *bool
	ImmediateLeave          *bool
	LastMemberQueryInterval *float64
	MaxGroups               *string
	MaxQueryResponseTime    *float64
	MaxSources              *string
	QueryInterval           *int64
	Robustness              *string
	RouterAlertPolicing     *bool
	Version                 *string
}
type MulticastInterfaceGroupPim struct {
	AllowedNeighbors  []MulticastInterfaceGroupPimAllowedNeighbors
	AssertInterval    *int64
	BsrBorder         *bool
	DrPriority        *int64
	Enable            *bool
	HelloInterval     *int64
	JoinPruneInterval *int64
}
type MulticastInterfaceGroupPimAllowedNeighbors struct {
	Name string
}
type MulticastRp struct {
	ExternalRp []MulticastRpExternalRp
	LocalRp    *MulticastRpLocalRp
}
type MulticastRpExternalRp struct {
	GroupAddresses []string
	Name           string
	Override       *bool
}
type MulticastRpLocalRp struct {
	CandidateRp *MulticastRpLocalRpCandidateRp
	StaticRp    *MulticastRpLocalRpStaticRp
}
type MulticastRpLocalRpCandidateRp struct {
	Address               *string
	AdvertisementInterval *int64
	GroupAddresses        []string
	Interface             *string
	Priority              *int64
}
type MulticastRpLocalRpStaticRp struct {
	Address        *string
	GroupAddresses []string
	Interface      *string
	Override       *bool
}
type MulticastSptThreshold struct {
	Name      string
	Threshold *string
}
type MulticastSsmAddressSpace struct {
	GroupAddress *string
	Included     *bool
	Name         string
}
type Protocol struct {
	Bgp               *ProtocolBgp
	Ospf              *ProtocolOspf
	Ospfv3            *ProtocolOspfv3
	RedistProfile     []ProtocolRedistProfile
	RedistProfileIpv6 []ProtocolRedistProfileIpv6
	Rip               *ProtocolRip
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
}
type ProtocolBgpAuthProfile struct {
	Name   string
	Secret *string
}
type ProtocolBgpDampeningProfile struct {
	Cutoff                   *float64
	DecayHalfLifeReachable   *int64
	DecayHalfLifeUnreachable *int64
	Enable                   *bool
	MaxHoldTime              *int64
	Name                     string
	Reuse                    *float64
}
type ProtocolBgpGlobalBfd struct {
	Profile *string
}
type ProtocolBgpPeerGroup struct {
	AggregatedConfedAsPath  *bool
	Enable                  *bool
	Name                    string
	Peer                    []ProtocolBgpPeerGroupPeer
	SoftResetWithStoredInfo *bool
	Type                    *ProtocolBgpPeerGroupType
}
type ProtocolBgpPeerGroupPeer struct {
	AddressFamilyIdentifier           *string
	Bfd                               *ProtocolBgpPeerGroupPeerBfd
	ConnectionOptions                 *ProtocolBgpPeerGroupPeerConnectionOptions
	Enable                            *bool
	EnableMpBgp                       *bool
	EnableSenderSideLoopDetection     *bool
	LocalAddress                      *ProtocolBgpPeerGroupPeerLocalAddress
	MaxPrefixes                       *string
	Name                              string
	PeerAddress                       *ProtocolBgpPeerGroupPeerPeerAddress
	PeerAs                            *string
	PeeringType                       *string
	ReflectorClient                   *string
	SubsequentAddressFamilyIdentifier *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier
}
type ProtocolBgpPeerGroupPeerBfd struct {
	Profile *string
}
type ProtocolBgpPeerGroupPeerConnectionOptions struct {
	Authentication        *string
	HoldTime              *string
	IdleHoldTime          *int64
	IncomingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection
	KeepAliveInterval     *string
	MinRouteAdvInterval   *int64
	Multihop              *int64
	OpenDelayTime         *int64
	OutgoingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection
}
type ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection struct {
	Allow      *bool
	RemotePort *int64
}
type ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection struct {
	Allow     *bool
	LocalPort *int64
}
type ProtocolBgpPeerGroupPeerLocalAddress struct {
	Interface *string
	Ip        *string
}
type ProtocolBgpPeerGroupPeerPeerAddress struct {
	Fqdn *string
	Ip   *string
}
type ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier struct {
	Multicast *bool
	Unicast   *bool
}
type ProtocolBgpPeerGroupType struct {
	Ebgp       *ProtocolBgpPeerGroupTypeEbgp
	EbgpConfed *ProtocolBgpPeerGroupTypeEbgpConfed
	Ibgp       *ProtocolBgpPeerGroupTypeIbgp
	IbgpConfed *ProtocolBgpPeerGroupTypeIbgpConfed
}
type ProtocolBgpPeerGroupTypeEbgp struct {
	ExportNexthop   *string
	ImportNexthop   *string
	RemovePrivateAs *bool
}
type ProtocolBgpPeerGroupTypeEbgpConfed struct {
	ExportNexthop *string
}
type ProtocolBgpPeerGroupTypeIbgp struct {
	ExportNexthop *string
}
type ProtocolBgpPeerGroupTypeIbgpConfed struct {
	ExportNexthop *string
}
type ProtocolBgpPolicy struct {
	Aggregation              *ProtocolBgpPolicyAggregation
	ConditionalAdvertisement *ProtocolBgpPolicyConditionalAdvertisement
	Export                   *ProtocolBgpPolicyExport
	Import                   *ProtocolBgpPolicyImport
}
type ProtocolBgpPolicyAggregation struct {
	Address []ProtocolBgpPolicyAggregationAddress
}
type ProtocolBgpPolicyAggregationAddress struct {
	AdvertiseFilters         []ProtocolBgpPolicyAggregationAddressAdvertiseFilters
	AggregateRouteAttributes *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes
	AsSet                    *bool
	Enable                   *bool
	Name                     string
	Prefix                   *string
	Summary                  *bool
	SuppressFilters          []ProtocolBgpPolicyAggregationAddressSuppressFilters
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFilters struct {
	Enable *bool
	Match  *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch
	Name   string
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch struct {
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath
	Community         *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix struct {
	Exact *bool
	Name  string
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes struct {
	AsPath            *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath
	AsPathLimit       *int64
	Community         *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity
	LocalPreference   *int64
	Med               *int64
	Nexthop           *string
	Origin            *string
	Weight            *int64
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath struct {
	None    *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone
	Prepend *int64
	Remove  *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone struct {
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove struct {
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone struct {
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll struct {
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone struct {
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll struct {
}
type ProtocolBgpPolicyAggregationAddressSuppressFilters struct {
	Enable *bool
	Match  *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch
	Name   string
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch struct {
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath
	Community         *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix struct {
	Exact *bool
	Name  string
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisement struct {
	Policy []ProtocolBgpPolicyConditionalAdvertisementPolicy
}
type ProtocolBgpPolicyConditionalAdvertisementPolicy struct {
	AdvertiseFilters []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters
	Enable           *bool
	Name             string
	NonExistFilters  []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters
	UsedBy           []string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters struct {
	Enable *bool
	Match  *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch
	Name   string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch struct {
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix struct {
	Name string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters struct {
	Enable *bool
	Match  *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch
	Name   string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch struct {
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix struct {
	Name string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyExport struct {
	Rules []ProtocolBgpPolicyExportRules
}
type ProtocolBgpPolicyExportRules struct {
	Action *ProtocolBgpPolicyExportRulesAction
	Enable *bool
	Match  *ProtocolBgpPolicyExportRulesMatch
	Name   string
	UsedBy []string
}
type ProtocolBgpPolicyExportRulesAction struct {
	Allow *ProtocolBgpPolicyExportRulesActionAllow
	Deny  *ProtocolBgpPolicyExportRulesActionDeny
}
type ProtocolBgpPolicyExportRulesActionAllow struct {
	Update *ProtocolBgpPolicyExportRulesActionAllowUpdate
}
type ProtocolBgpPolicyExportRulesActionAllowUpdate struct {
	AsPath            *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath
	AsPathLimit       *int64
	Community         *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity
	ExtendedCommunity *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity
	LocalPreference   *int64
	Med               *int64
	Nexthop           *string
	Origin            *string
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath struct {
	None             *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone
	Prepend          *int64
	Remove           *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove
	RemoveAndPrepend *int64
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone struct {
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove struct {
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone struct {
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll struct {
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone struct {
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll struct {
}
type ProtocolBgpPolicyExportRulesActionDeny struct {
}
type ProtocolBgpPolicyExportRulesMatch struct {
	AddressPrefix     []ProtocolBgpPolicyExportRulesMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyExportRulesMatchAsPath
	Community         *ProtocolBgpPolicyExportRulesMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyExportRulesMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyExportRulesMatchAddressPrefix struct {
	Exact *bool
	Name  string
}
type ProtocolBgpPolicyExportRulesMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyExportRulesMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyExportRulesMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyImport struct {
	Rules []ProtocolBgpPolicyImportRules
}
type ProtocolBgpPolicyImportRules struct {
	Action *ProtocolBgpPolicyImportRulesAction
	Enable *bool
	Match  *ProtocolBgpPolicyImportRulesMatch
	Name   string
	UsedBy []string
}
type ProtocolBgpPolicyImportRulesAction struct {
	Allow *ProtocolBgpPolicyImportRulesActionAllow
	Deny  *ProtocolBgpPolicyImportRulesActionDeny
}
type ProtocolBgpPolicyImportRulesActionAllow struct {
	Dampening *string
	Update    *ProtocolBgpPolicyImportRulesActionAllowUpdate
}
type ProtocolBgpPolicyImportRulesActionAllowUpdate struct {
	AsPath            *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath
	AsPathLimit       *int64
	Community         *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity
	ExtendedCommunity *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity
	LocalPreference   *int64
	Med               *int64
	Nexthop           *string
	Origin            *string
	Weight            *int64
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath struct {
	None   *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone
	Remove *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone struct {
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove struct {
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone struct {
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll struct {
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity struct {
	Append      []string
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone
	Overwrite   []string
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll
	RemoveRegex *string
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone struct {
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll struct {
}
type ProtocolBgpPolicyImportRulesActionDeny struct {
}
type ProtocolBgpPolicyImportRulesMatch struct {
	AddressPrefix     []ProtocolBgpPolicyImportRulesMatchAddressPrefix
	AsPath            *ProtocolBgpPolicyImportRulesMatchAsPath
	Community         *ProtocolBgpPolicyImportRulesMatchCommunity
	ExtendedCommunity *ProtocolBgpPolicyImportRulesMatchExtendedCommunity
	FromPeer          []string
	Med               *int64
	Nexthop           []string
	RouteTable        *string
}
type ProtocolBgpPolicyImportRulesMatchAddressPrefix struct {
	Exact *bool
	Name  string
}
type ProtocolBgpPolicyImportRulesMatchAsPath struct {
	Regex *string
}
type ProtocolBgpPolicyImportRulesMatchCommunity struct {
	Regex *string
}
type ProtocolBgpPolicyImportRulesMatchExtendedCommunity struct {
	Regex *string
}
type ProtocolBgpRedistRules struct {
	AddressFamilyIdentifier *string
	Enable                  *bool
	Metric                  *int64
	Name                    string
	RouteTable              *string
	SetAsPathLimit          *int64
	SetCommunity            []string
	SetExtendedCommunity    []string
	SetLocalPreference      *int64
	SetMed                  *int64
	SetOrigin               *string
}
type ProtocolBgpRoutingOptions struct {
	Aggregate              *ProtocolBgpRoutingOptionsAggregate
	AsFormat               *string
	ConfederationMemberAs  *string
	DefaultLocalPreference *int64
	GracefulRestart        *ProtocolBgpRoutingOptionsGracefulRestart
	Med                    *ProtocolBgpRoutingOptionsMed
	ReflectorClusterId     *string
}
type ProtocolBgpRoutingOptionsAggregate struct {
	AggregateMed *bool
}
type ProtocolBgpRoutingOptionsGracefulRestart struct {
	Enable             *bool
	LocalRestartTime   *int64
	MaxPeerRestartTime *int64
	StaleRouteTime     *int64
}
type ProtocolBgpRoutingOptionsMed struct {
	AlwaysCompareMed           *bool
	DeterministicMedComparison *bool
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
}
type ProtocolOspfArea struct {
	Interface   []ProtocolOspfAreaInterface
	Name        string
	Range       []ProtocolOspfAreaRange
	Type        *ProtocolOspfAreaType
	VirtualLink []ProtocolOspfAreaVirtualLink
}
type ProtocolOspfAreaInterface struct {
	Authentication     *string
	Bfd                *ProtocolOspfAreaInterfaceBfd
	DeadCounts         *int64
	Enable             *bool
	GrDelay            *int64
	HelloInterval      *int64
	LinkType           *ProtocolOspfAreaInterfaceLinkType
	Metric             *int64
	Name               string
	Neighbor           []ProtocolOspfAreaInterfaceNeighbor
	Passive            *bool
	Priority           *int64
	RetransmitInterval *int64
	TransitDelay       *int64
}
type ProtocolOspfAreaInterfaceBfd struct {
	Profile *string
}
type ProtocolOspfAreaInterfaceLinkType struct {
	Broadcast *ProtocolOspfAreaInterfaceLinkTypeBroadcast
	P2mp      *ProtocolOspfAreaInterfaceLinkTypeP2mp
	P2p       *ProtocolOspfAreaInterfaceLinkTypeP2p
}
type ProtocolOspfAreaInterfaceLinkTypeBroadcast struct {
}
type ProtocolOspfAreaInterfaceLinkTypeP2mp struct {
}
type ProtocolOspfAreaInterfaceLinkTypeP2p struct {
}
type ProtocolOspfAreaInterfaceNeighbor struct {
	Name string
}
type ProtocolOspfAreaRange struct {
	Name      string
	Advertise *ProtocolOspfAreaRangeAdvertise
	Suppress  *ProtocolOspfAreaRangeSuppress
}
type ProtocolOspfAreaRangeAdvertise struct {
}
type ProtocolOspfAreaRangeSuppress struct {
}
type ProtocolOspfAreaType struct {
	Normal *ProtocolOspfAreaTypeNormal
	Nssa   *ProtocolOspfAreaTypeNssa
	Stub   *ProtocolOspfAreaTypeStub
}
type ProtocolOspfAreaTypeNormal struct {
}
type ProtocolOspfAreaTypeNssa struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfAreaTypeNssaDefaultRoute
	NssaExtRange  []ProtocolOspfAreaTypeNssaNssaExtRange
}
type ProtocolOspfAreaTypeNssaDefaultRoute struct {
	Advertise *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise
	Disable   *ProtocolOspfAreaTypeNssaDefaultRouteDisable
}
type ProtocolOspfAreaTypeNssaDefaultRouteAdvertise struct {
	Metric *int64
	Type   *string
}
type ProtocolOspfAreaTypeNssaDefaultRouteDisable struct {
}
type ProtocolOspfAreaTypeNssaNssaExtRange struct {
	Name      string
	Advertise *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise
	Suppress  *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress
}
type ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise struct {
}
type ProtocolOspfAreaTypeNssaNssaExtRangeSuppress struct {
}
type ProtocolOspfAreaTypeStub struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfAreaTypeStubDefaultRoute
}
type ProtocolOspfAreaTypeStubDefaultRoute struct {
	Advertise *ProtocolOspfAreaTypeStubDefaultRouteAdvertise
	Disable   *ProtocolOspfAreaTypeStubDefaultRouteDisable
}
type ProtocolOspfAreaTypeStubDefaultRouteAdvertise struct {
	Metric *int64
}
type ProtocolOspfAreaTypeStubDefaultRouteDisable struct {
}
type ProtocolOspfAreaVirtualLink struct {
	Authentication     *string
	Bfd                *ProtocolOspfAreaVirtualLinkBfd
	DeadCounts         *int64
	Enable             *bool
	HelloInterval      *int64
	Name               string
	NeighborId         *string
	RetransmitInterval *int64
	TransitAreaId      *string
	TransitDelay       *int64
}
type ProtocolOspfAreaVirtualLinkBfd struct {
	Profile *string
}
type ProtocolOspfAuthProfile struct {
	Name     string
	Md5      []ProtocolOspfAuthProfileMd5
	Password *string
}
type ProtocolOspfAuthProfileMd5 struct {
	Key       *string
	Name      string
	Preferred *bool
}
type ProtocolOspfExportRules struct {
	Metric      *int64
	Name        string
	NewPathType *string
	NewTag      *string
}
type ProtocolOspfGlobalBfd struct {
	Profile *string
}
type ProtocolOspfGracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	MaxNeighborRestartTime *int64
	StrictLSAChecking      *bool
}
type ProtocolOspfTimers struct {
	LsaInterval         *float64
	SpfCalculationDelay *float64
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
}
type ProtocolOspfv3Area struct {
	Authentication *string
	Interface      []ProtocolOspfv3AreaInterface
	Name           string
	Range          []ProtocolOspfv3AreaRange
	Type           *ProtocolOspfv3AreaType
	VirtualLink    []ProtocolOspfv3AreaVirtualLink
}
type ProtocolOspfv3AreaInterface struct {
	Authentication     *string
	Bfd                *ProtocolOspfv3AreaInterfaceBfd
	DeadCounts         *int64
	Enable             *bool
	GrDelay            *int64
	HelloInterval      *int64
	InstanceId         *int64
	LinkType           *ProtocolOspfv3AreaInterfaceLinkType
	Metric             *int64
	Name               string
	Neighbor           []ProtocolOspfv3AreaInterfaceNeighbor
	Passive            *bool
	Priority           *int64
	RetransmitInterval *int64
	TransitDelay       *int64
}
type ProtocolOspfv3AreaInterfaceBfd struct {
	Profile *string
}
type ProtocolOspfv3AreaInterfaceLinkType struct {
	Broadcast *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast
	P2mp      *ProtocolOspfv3AreaInterfaceLinkTypeP2mp
	P2p       *ProtocolOspfv3AreaInterfaceLinkTypeP2p
}
type ProtocolOspfv3AreaInterfaceLinkTypeBroadcast struct {
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2mp struct {
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2p struct {
}
type ProtocolOspfv3AreaInterfaceNeighbor struct {
	Name string
}
type ProtocolOspfv3AreaRange struct {
	Name      string
	Advertise *ProtocolOspfv3AreaRangeAdvertise
	Suppress  *ProtocolOspfv3AreaRangeSuppress
}
type ProtocolOspfv3AreaRangeAdvertise struct {
}
type ProtocolOspfv3AreaRangeSuppress struct {
}
type ProtocolOspfv3AreaType struct {
	Normal *ProtocolOspfv3AreaTypeNormal
	Nssa   *ProtocolOspfv3AreaTypeNssa
	Stub   *ProtocolOspfv3AreaTypeStub
}
type ProtocolOspfv3AreaTypeNormal struct {
}
type ProtocolOspfv3AreaTypeNssa struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfv3AreaTypeNssaDefaultRoute
	NssaExtRange  []ProtocolOspfv3AreaTypeNssaNssaExtRange
}
type ProtocolOspfv3AreaTypeNssaDefaultRoute struct {
	Advertise *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise
	Disable   *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise struct {
	Metric *int64
	Type   *string
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteDisable struct {
}
type ProtocolOspfv3AreaTypeNssaNssaExtRange struct {
	Name      string
	Advertise *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise
	Suppress  *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise struct {
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress struct {
}
type ProtocolOspfv3AreaTypeStub struct {
	AcceptSummary *bool
	DefaultRoute  *ProtocolOspfv3AreaTypeStubDefaultRoute
}
type ProtocolOspfv3AreaTypeStubDefaultRoute struct {
	Advertise *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise
	Disable   *ProtocolOspfv3AreaTypeStubDefaultRouteDisable
}
type ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise struct {
	Metric *int64
}
type ProtocolOspfv3AreaTypeStubDefaultRouteDisable struct {
}
type ProtocolOspfv3AreaVirtualLink struct {
	Authentication     *string
	Bfd                *ProtocolOspfv3AreaVirtualLinkBfd
	DeadCounts         *int64
	Enable             *bool
	HelloInterval      *int64
	InstanceId         *int64
	Name               string
	NeighborId         *string
	RetransmitInterval *int64
	TransitAreaId      *string
	TransitDelay       *int64
}
type ProtocolOspfv3AreaVirtualLinkBfd struct {
	Profile *string
}
type ProtocolOspfv3AuthProfile struct {
	Name string
	Spi  *string
	Ah   *ProtocolOspfv3AuthProfileAh
	Esp  *ProtocolOspfv3AuthProfileEsp
}
type ProtocolOspfv3AuthProfileAh struct {
	Md5    *ProtocolOspfv3AuthProfileAhMd5
	Sha1   *ProtocolOspfv3AuthProfileAhSha1
	Sha256 *ProtocolOspfv3AuthProfileAhSha256
	Sha384 *ProtocolOspfv3AuthProfileAhSha384
	Sha512 *ProtocolOspfv3AuthProfileAhSha512
}
type ProtocolOspfv3AuthProfileAhMd5 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileAhSha1 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileAhSha256 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileAhSha384 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileAhSha512 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEsp struct {
	Authentication *ProtocolOspfv3AuthProfileEspAuthentication
	Encryption     *ProtocolOspfv3AuthProfileEspEncryption
}
type ProtocolOspfv3AuthProfileEspAuthentication struct {
	Md5    *ProtocolOspfv3AuthProfileEspAuthenticationMd5
	None   *ProtocolOspfv3AuthProfileEspAuthenticationNone
	Sha1   *ProtocolOspfv3AuthProfileEspAuthenticationSha1
	Sha256 *ProtocolOspfv3AuthProfileEspAuthenticationSha256
	Sha384 *ProtocolOspfv3AuthProfileEspAuthenticationSha384
	Sha512 *ProtocolOspfv3AuthProfileEspAuthenticationSha512
}
type ProtocolOspfv3AuthProfileEspAuthenticationMd5 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEspAuthenticationNone struct {
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha1 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha256 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha384 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha512 struct {
	Key *string
}
type ProtocolOspfv3AuthProfileEspEncryption struct {
	Algorithm *string
	Key       *string
}
type ProtocolOspfv3ExportRules struct {
	Metric      *int64
	Name        string
	NewPathType *string
	NewTag      *string
}
type ProtocolOspfv3GlobalBfd struct {
	Profile *string
}
type ProtocolOspfv3GracefulRestart struct {
	Enable                 *bool
	GracePeriod            *int64
	HelperEnable           *bool
	MaxNeighborRestartTime *int64
	StrictLSAChecking      *bool
}
type ProtocolOspfv3Timers struct {
	LsaInterval         *float64
	SpfCalculationDelay *float64
}
type ProtocolRedistProfile struct {
	Action   *ProtocolRedistProfileAction
	Filter   *ProtocolRedistProfileFilter
	Name     string
	Priority *int64
}
type ProtocolRedistProfileAction struct {
	NoRedist *ProtocolRedistProfileActionNoRedist
	Redist   *ProtocolRedistProfileActionRedist
}
type ProtocolRedistProfileActionNoRedist struct {
}
type ProtocolRedistProfileActionRedist struct {
}
type ProtocolRedistProfileFilter struct {
	Bgp         *ProtocolRedistProfileFilterBgp
	Destination []string
	Interface   []string
	Nexthop     []string
	Ospf        *ProtocolRedistProfileFilterOspf
	Type        []string
}
type ProtocolRedistProfileFilterBgp struct {
	Community         []string
	ExtendedCommunity []string
}
type ProtocolRedistProfileFilterOspf struct {
	Area     []string
	PathType []string
	Tag      []string
}
type ProtocolRedistProfileIpv6 struct {
	Action   *ProtocolRedistProfileIpv6Action
	Filter   *ProtocolRedistProfileIpv6Filter
	Name     string
	Priority *int64
}
type ProtocolRedistProfileIpv6Action struct {
	NoRedist *ProtocolRedistProfileIpv6ActionNoRedist
	Redist   *ProtocolRedistProfileIpv6ActionRedist
}
type ProtocolRedistProfileIpv6ActionNoRedist struct {
}
type ProtocolRedistProfileIpv6ActionRedist struct {
}
type ProtocolRedistProfileIpv6Filter struct {
	Bgp         *ProtocolRedistProfileIpv6FilterBgp
	Destination []string
	Interface   []string
	Nexthop     []string
	Ospfv3      *ProtocolRedistProfileIpv6FilterOspfv3
	Type        []string
}
type ProtocolRedistProfileIpv6FilterBgp struct {
	Community         []string
	ExtendedCommunity []string
}
type ProtocolRedistProfileIpv6FilterOspfv3 struct {
	Area     []string
	PathType []string
	Tag      []string
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
}
type ProtocolRipAuthProfile struct {
	Name     string
	Md5      []ProtocolRipAuthProfileMd5
	Password *string
}
type ProtocolRipAuthProfileMd5 struct {
	Key       *string
	Name      string
	Preferred *bool
}
type ProtocolRipExportRules struct {
	Metric *int64
	Name   string
}
type ProtocolRipGlobalBfd struct {
	Profile *string
}
type ProtocolRipInterface struct {
	Authentication *string
	Bfd            *ProtocolRipInterfaceBfd
	DefaultRoute   *ProtocolRipInterfaceDefaultRoute
	Enable         *bool
	Mode           *string
	Name           string
}
type ProtocolRipInterfaceBfd struct {
	Profile *string
}
type ProtocolRipInterfaceDefaultRoute struct {
	Advertise *ProtocolRipInterfaceDefaultRouteAdvertise
	Disable   *ProtocolRipInterfaceDefaultRouteDisable
}
type ProtocolRipInterfaceDefaultRouteAdvertise struct {
	Metric *int64
}
type ProtocolRipInterfaceDefaultRouteDisable struct {
}
type ProtocolRipTimers struct {
	DeleteIntervals *int64
	ExpireIntervals *int64
	IntervalSeconds *int64
	UpdateIntervals *int64
}
type RoutingTable struct {
	Ip   *RoutingTableIp
	Ipv6 *RoutingTableIpv6
}
type RoutingTableIp struct {
	StaticRoute []RoutingTableIpStaticRoute
}
type RoutingTableIpStaticRoute struct {
	AdminDist   *int64
	Bfd         *RoutingTableIpStaticRouteBfd
	Destination *string
	Interface   *string
	Metric      *int64
	Name        string
	Nexthop     *RoutingTableIpStaticRouteNexthop
	PathMonitor *RoutingTableIpStaticRoutePathMonitor
	RouteTable  *RoutingTableIpStaticRouteRouteTable
}
type RoutingTableIpStaticRouteBfd struct {
	Profile *string
}
type RoutingTableIpStaticRouteNexthop struct {
	Discard   *RoutingTableIpStaticRouteNexthopDiscard
	Fqdn      *string
	IpAddress *string
	NextVr    *string
	Receive   *RoutingTableIpStaticRouteNexthopReceive
}
type RoutingTableIpStaticRouteNexthopDiscard struct {
}
type RoutingTableIpStaticRouteNexthopReceive struct {
}
type RoutingTableIpStaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []RoutingTableIpStaticRoutePathMonitorMonitorDestinations
}
type RoutingTableIpStaticRoutePathMonitorMonitorDestinations struct {
	Count       *int64
	Destination *string
	Enable      *bool
	Interval    *int64
	Name        string
	Source      *string
}
type RoutingTableIpStaticRouteRouteTable struct {
	Both      *RoutingTableIpStaticRouteRouteTableBoth
	Multicast *RoutingTableIpStaticRouteRouteTableMulticast
	NoInstall *RoutingTableIpStaticRouteRouteTableNoInstall
	Unicast   *RoutingTableIpStaticRouteRouteTableUnicast
}
type RoutingTableIpStaticRouteRouteTableBoth struct {
}
type RoutingTableIpStaticRouteRouteTableMulticast struct {
}
type RoutingTableIpStaticRouteRouteTableNoInstall struct {
}
type RoutingTableIpStaticRouteRouteTableUnicast struct {
}
type RoutingTableIpv6 struct {
	StaticRoute []RoutingTableIpv6StaticRoute
}
type RoutingTableIpv6StaticRoute struct {
	AdminDist   *int64
	Bfd         *RoutingTableIpv6StaticRouteBfd
	Destination *string
	Interface   *string
	Metric      *int64
	Name        string
	Nexthop     *RoutingTableIpv6StaticRouteNexthop
	Option      *RoutingTableIpv6StaticRouteOption
	PathMonitor *RoutingTableIpv6StaticRoutePathMonitor
	RouteTable  *RoutingTableIpv6StaticRouteRouteTable
}
type RoutingTableIpv6StaticRouteBfd struct {
	Profile *string
}
type RoutingTableIpv6StaticRouteNexthop struct {
	Discard     *RoutingTableIpv6StaticRouteNexthopDiscard
	Ipv6Address *string
	NextVr      *string
	Receive     *RoutingTableIpv6StaticRouteNexthopReceive
}
type RoutingTableIpv6StaticRouteNexthopDiscard struct {
}
type RoutingTableIpv6StaticRouteNexthopReceive struct {
}
type RoutingTableIpv6StaticRouteOption struct {
}
type RoutingTableIpv6StaticRoutePathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations
}
type RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations struct {
	Count       *int64
	Destination *string
	Enable      *bool
	Interval    *int64
	Name        string
	Source      *string
}
type RoutingTableIpv6StaticRouteRouteTable struct {
	NoInstall *RoutingTableIpv6StaticRouteRouteTableNoInstall
	Unicast   *RoutingTableIpv6StaticRouteRouteTableUnicast
}
type RoutingTableIpv6StaticRouteRouteTableNoInstall struct {
}
type RoutingTableIpv6StaticRouteRouteTableUnicast struct {
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName      xml.Name         `xml:"entry"`
	Name         string           `xml:"name,attr"`
	AdminDists   *AdminDistsXml   `xml:"admin-dists,omitempty"`
	Ecmp         *EcmpXml         `xml:"ecmp,omitempty"`
	Interface    *util.MemberType `xml:"interface,omitempty"`
	Multicast    *MulticastXml    `xml:"multicast,omitempty"`
	Protocol     *ProtocolXml     `xml:"protocol,omitempty"`
	RoutingTable *RoutingTableXml `xml:"routing-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AdminDistsXml struct {
	Ebgp       *int64 `xml:"ebgp,omitempty"`
	Ibgp       *int64 `xml:"ibgp,omitempty"`
	OspfExt    *int64 `xml:"ospf-ext,omitempty"`
	OspfInt    *int64 `xml:"ospf-int,omitempty"`
	Ospfv3Ext  *int64 `xml:"ospfv3-ext,omitempty"`
	Ospfv3Int  *int64 `xml:"ospfv3-int,omitempty"`
	Rip        *int64 `xml:"rip,omitempty"`
	Static     *int64 `xml:"static,omitempty"`
	StaticIpv6 *int64 `xml:"static-ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpXml struct {
	Algorithm        *EcmpAlgorithmXml `xml:"algorithm,omitempty"`
	Enable           *string           `xml:"enable,omitempty"`
	MaxPath          *int64            `xml:"max-path,omitempty"`
	StrictSourcePath *string           `xml:"strict-source-path,omitempty"`
	SymmetricReturn  *string           `xml:"symmetric-return,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmXml struct {
	BalancedRoundRobin *EcmpAlgorithmBalancedRoundRobinXml `xml:"balanced-round-robin,omitempty"`
	IpHash             *EcmpAlgorithmIpHashXml             `xml:"ip-hash,omitempty"`
	IpModulo           *EcmpAlgorithmIpModuloXml           `xml:"ip-modulo,omitempty"`
	WeightedRoundRobin *EcmpAlgorithmWeightedRoundRobinXml `xml:"weighted-round-robin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmBalancedRoundRobinXml struct {
	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmIpHashXml struct {
	HashSeed *int64  `xml:"hash-seed,omitempty"`
	SrcOnly  *string `xml:"src-only,omitempty"`
	UsePort  *string `xml:"use-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmIpModuloXml struct {
	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmWeightedRoundRobinXml struct {
	Interface []EcmpAlgorithmWeightedRoundRobinInterfaceXml `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmWeightedRoundRobinInterfaceXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Weight  *int64   `xml:"weight,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastXml struct {
	Enable          *string                       `xml:"enable,omitempty"`
	InterfaceGroup  []MulticastInterfaceGroupXml  `xml:"interface-group>entry,omitempty"`
	RouteAgeoutTime *int64                        `xml:"route-ageout-time,omitempty"`
	Rp              *MulticastRpXml               `xml:"rp,omitempty"`
	SptThreshold    []MulticastSptThresholdXml    `xml:"spt-threshold>entry,omitempty"`
	SsmAddressSpace []MulticastSsmAddressSpaceXml `xml:"ssm-address-space>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupXml struct {
	Description     *string                                    `xml:"description,omitempty"`
	GroupPermission *MulticastInterfaceGroupGroupPermissionXml `xml:"group-permission,omitempty"`
	Igmp            *MulticastInterfaceGroupIgmpXml            `xml:"igmp,omitempty"`
	Interface       *util.MemberType                           `xml:"interface,omitempty"`
	XMLName         xml.Name                                   `xml:"entry"`
	Name            string                                     `xml:"name,attr"`
	Pim             *MulticastInterfaceGroupPimXml             `xml:"pim,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupGroupPermissionXml struct {
	AnySourceMulticast      []MulticastInterfaceGroupGroupPermissionAnySourceMulticastXml      `xml:"any-source-multicast>entry,omitempty"`
	SourceSpecificMulticast []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticastXml `xml:"source-specific-multicast>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupGroupPermissionAnySourceMulticastXml struct {
	GroupAddress *string  `xml:"group-address,omitempty"`
	Included     *string  `xml:"included,omitempty"`
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupGroupPermissionSourceSpecificMulticastXml struct {
	GroupAddress  *string  `xml:"group-address,omitempty"`
	Included      *string  `xml:"included,omitempty"`
	XMLName       xml.Name `xml:"entry"`
	Name          string   `xml:"name,attr"`
	SourceAddress *string  `xml:"source-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupIgmpXml struct {
	Enable                  *string  `xml:"enable,omitempty"`
	ImmediateLeave          *string  `xml:"immediate-leave,omitempty"`
	LastMemberQueryInterval *float64 `xml:"last-member-query-interval,omitempty"`
	MaxGroups               *string  `xml:"max-groups,omitempty"`
	MaxQueryResponseTime    *float64 `xml:"max-query-response-time,omitempty"`
	MaxSources              *string  `xml:"max-sources,omitempty"`
	QueryInterval           *int64   `xml:"query-interval,omitempty"`
	Robustness              *string  `xml:"robustness,omitempty"`
	RouterAlertPolicing     *string  `xml:"router-alert-policing,omitempty"`
	Version                 *string  `xml:"version,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupPimXml struct {
	AllowedNeighbors  []MulticastInterfaceGroupPimAllowedNeighborsXml `xml:"allowed-neighbors>entry,omitempty"`
	AssertInterval    *int64                                          `xml:"assert-interval,omitempty"`
	BsrBorder         *string                                         `xml:"bsr-border,omitempty"`
	DrPriority        *int64                                          `xml:"dr-priority,omitempty"`
	Enable            *string                                         `xml:"enable,omitempty"`
	HelloInterval     *int64                                          `xml:"hello-interval,omitempty"`
	JoinPruneInterval *int64                                          `xml:"join-prune-interval,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastInterfaceGroupPimAllowedNeighborsXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastRpXml struct {
	ExternalRp []MulticastRpExternalRpXml `xml:"external-rp>entry,omitempty"`
	LocalRp    *MulticastRpLocalRpXml     `xml:"local-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastRpExternalRpXml struct {
	GroupAddresses *util.MemberType `xml:"group-addresses,omitempty"`
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Override       *string          `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastRpLocalRpXml struct {
	CandidateRp *MulticastRpLocalRpCandidateRpXml `xml:"candidate-rp,omitempty"`
	StaticRp    *MulticastRpLocalRpStaticRpXml    `xml:"static-rp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastRpLocalRpCandidateRpXml struct {
	Address               *string          `xml:"address,omitempty"`
	AdvertisementInterval *int64           `xml:"advertisement-interval,omitempty"`
	GroupAddresses        *util.MemberType `xml:"group-addresses,omitempty"`
	Interface             *string          `xml:"interface,omitempty"`
	Priority              *int64           `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastRpLocalRpStaticRpXml struct {
	Address        *string          `xml:"address,omitempty"`
	GroupAddresses *util.MemberType `xml:"group-addresses,omitempty"`
	Interface      *string          `xml:"interface,omitempty"`
	Override       *string          `xml:"override,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastSptThresholdXml struct {
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`
	Threshold *string  `xml:"threshold,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MulticastSsmAddressSpaceXml struct {
	GroupAddress *string  `xml:"group-address,omitempty"`
	Included     *string  `xml:"included,omitempty"`
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolXml struct {
	Bgp               *ProtocolBgpXml                `xml:"bgp,omitempty"`
	Ospf              *ProtocolOspfXml               `xml:"ospf,omitempty"`
	Ospfv3            *ProtocolOspfv3Xml             `xml:"ospfv3,omitempty"`
	RedistProfile     []ProtocolRedistProfileXml     `xml:"redist-profile>entry,omitempty"`
	RedistProfileIpv6 []ProtocolRedistProfileIpv6Xml `xml:"redist-profile-ipv6>entry,omitempty"`
	Rip               *ProtocolRipXml                `xml:"rip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpXml struct {
	AllowRedistDefaultRoute *string                          `xml:"allow-redist-default-route,omitempty"`
	AuthProfile             []ProtocolBgpAuthProfileXml      `xml:"auth-profile>entry,omitempty"`
	DampeningProfile        []ProtocolBgpDampeningProfileXml `xml:"dampening-profile>entry,omitempty"`
	EcmpMultiAs             *string                          `xml:"ecmp-multi-as,omitempty"`
	Enable                  *string                          `xml:"enable,omitempty"`
	EnforceFirstAs          *string                          `xml:"enforce-first-as,omitempty"`
	GlobalBfd               *ProtocolBgpGlobalBfdXml         `xml:"global-bfd,omitempty"`
	InstallRoute            *string                          `xml:"install-route,omitempty"`
	LocalAs                 *string                          `xml:"local-as,omitempty"`
	PeerGroup               []ProtocolBgpPeerGroupXml        `xml:"peer-group>entry,omitempty"`
	Policy                  *ProtocolBgpPolicyXml            `xml:"policy,omitempty"`
	RedistRules             []ProtocolBgpRedistRulesXml      `xml:"redist-rules>entry,omitempty"`
	RejectDefaultRoute      *string                          `xml:"reject-default-route,omitempty"`
	RouterId                *string                          `xml:"router-id,omitempty"`
	RoutingOptions          *ProtocolBgpRoutingOptionsXml    `xml:"routing-options,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpAuthProfileXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Secret  *string  `xml:"secret,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpDampeningProfileXml struct {
	Cutoff                   *float64 `xml:"cutoff,omitempty"`
	DecayHalfLifeReachable   *int64   `xml:"decay-half-life-reachable,omitempty"`
	DecayHalfLifeUnreachable *int64   `xml:"decay-half-life-unreachable,omitempty"`
	Enable                   *string  `xml:"enable,omitempty"`
	MaxHoldTime              *int64   `xml:"max-hold-time,omitempty"`
	XMLName                  xml.Name `xml:"entry"`
	Name                     string   `xml:"name,attr"`
	Reuse                    *float64 `xml:"reuse,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupXml struct {
	AggregatedConfedAsPath  *string                       `xml:"aggregated-confed-as-path,omitempty"`
	Enable                  *string                       `xml:"enable,omitempty"`
	XMLName                 xml.Name                      `xml:"entry"`
	Name                    string                        `xml:"name,attr"`
	Peer                    []ProtocolBgpPeerGroupPeerXml `xml:"peer>entry,omitempty"`
	SoftResetWithStoredInfo *string                       `xml:"soft-reset-with-stored-info,omitempty"`
	Type                    *ProtocolBgpPeerGroupTypeXml  `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerXml struct {
	AddressFamilyIdentifier           *string                                                       `xml:"address-family-identifier,omitempty"`
	Bfd                               *ProtocolBgpPeerGroupPeerBfdXml                               `xml:"bfd,omitempty"`
	ConnectionOptions                 *ProtocolBgpPeerGroupPeerConnectionOptionsXml                 `xml:"connection-options,omitempty"`
	Enable                            *string                                                       `xml:"enable,omitempty"`
	EnableMpBgp                       *string                                                       `xml:"enable-mp-bgp,omitempty"`
	EnableSenderSideLoopDetection     *string                                                       `xml:"enable-sender-side-loop-detection,omitempty"`
	LocalAddress                      *ProtocolBgpPeerGroupPeerLocalAddressXml                      `xml:"local-address,omitempty"`
	MaxPrefixes                       *string                                                       `xml:"max-prefixes,omitempty"`
	XMLName                           xml.Name                                                      `xml:"entry"`
	Name                              string                                                        `xml:"name,attr"`
	PeerAddress                       *ProtocolBgpPeerGroupPeerPeerAddressXml                       `xml:"peer-address,omitempty"`
	PeerAs                            *string                                                       `xml:"peer-as,omitempty"`
	PeeringType                       *string                                                       `xml:"peering-type,omitempty"`
	ReflectorClient                   *string                                                       `xml:"reflector-client,omitempty"`
	SubsequentAddressFamilyIdentifier *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml `xml:"subsequent-address-family-identifier,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerConnectionOptionsXml struct {
	Authentication        *string                                                            `xml:"authentication,omitempty"`
	HoldTime              *string                                                            `xml:"hold-time,omitempty"`
	IdleHoldTime          *int64                                                             `xml:"idle-hold-time,omitempty"`
	IncomingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml `xml:"incoming-bgp-connection,omitempty"`
	KeepAliveInterval     *string                                                            `xml:"keep-alive-interval,omitempty"`
	MinRouteAdvInterval   *int64                                                             `xml:"min-route-adv-interval,omitempty"`
	Multihop              *int64                                                             `xml:"multihop,omitempty"`
	OpenDelayTime         *int64                                                             `xml:"open-delay-time,omitempty"`
	OutgoingBgpConnection *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml `xml:"outgoing-bgp-connection,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml struct {
	Allow      *string `xml:"allow,omitempty"`
	RemotePort *int64  `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml struct {
	Allow     *string `xml:"allow,omitempty"`
	LocalPort *int64  `xml:"local-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerLocalAddressXml struct {
	Interface *string `xml:"interface,omitempty"`
	Ip        *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerPeerAddressXml struct {
	Fqdn *string `xml:"fqdn,omitempty"`
	Ip   *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml struct {
	Multicast *string `xml:"multicast,omitempty"`
	Unicast   *string `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupTypeXml struct {
	Ebgp       *ProtocolBgpPeerGroupTypeEbgpXml       `xml:"ebgp,omitempty"`
	EbgpConfed *ProtocolBgpPeerGroupTypeEbgpConfedXml `xml:"ebgp-confed,omitempty"`
	Ibgp       *ProtocolBgpPeerGroupTypeIbgpXml       `xml:"ibgp,omitempty"`
	IbgpConfed *ProtocolBgpPeerGroupTypeIbgpConfedXml `xml:"ibgp-confed,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupTypeEbgpXml struct {
	ExportNexthop   *string `xml:"export-nexthop,omitempty"`
	ImportNexthop   *string `xml:"import-nexthop,omitempty"`
	RemovePrivateAs *string `xml:"remove-private-as,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupTypeEbgpConfedXml struct {
	ExportNexthop *string `xml:"export-nexthop,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupTypeIbgpXml struct {
	ExportNexthop *string `xml:"export-nexthop,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPeerGroupTypeIbgpConfedXml struct {
	ExportNexthop *string `xml:"export-nexthop,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyXml struct {
	Aggregation              *ProtocolBgpPolicyAggregationXml              `xml:"aggregation,omitempty"`
	ConditionalAdvertisement *ProtocolBgpPolicyConditionalAdvertisementXml `xml:"conditional-advertisement,omitempty"`
	Export                   *ProtocolBgpPolicyExportXml                   `xml:"export,omitempty"`
	Import                   *ProtocolBgpPolicyImportXml                   `xml:"import,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationXml struct {
	Address []ProtocolBgpPolicyAggregationAddressXml `xml:"address>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressXml struct {
	AdvertiseFilters         []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersXml        `xml:"advertise-filters>entry,omitempty"`
	AggregateRouteAttributes *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesXml `xml:"aggregate-route-attributes,omitempty"`
	AsSet                    *string                                                         `xml:"as-set,omitempty"`
	Enable                   *string                                                         `xml:"enable,omitempty"`
	XMLName                  xml.Name                                                        `xml:"entry"`
	Name                     string                                                          `xml:"name,attr"`
	Prefix                   *string                                                         `xml:"prefix,omitempty"`
	Summary                  *string                                                         `xml:"summary,omitempty"`
	SuppressFilters          []ProtocolBgpPolicyAggregationAddressSuppressFiltersXml         `xml:"suppress-filters>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersXml struct {
	Enable  *string                                                      `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml `xml:"match,omitempty"`
	XMLName xml.Name                                                     `xml:"entry"`
	Name    string                                                       `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                                              `xml:"from-peer,omitempty"`
	Med               *int64                                                                        `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                                              `xml:"nexthop,omitempty"`
	RouteTable        *string                                                                       `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml struct {
	Exact   *string  `xml:"exact,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesXml struct {
	AsPath            *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml            `xml:"as-path,omitempty"`
	AsPathLimit       *int64                                                                           `xml:"as-path-limit,omitempty"`
	Community         *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml `xml:"extended-community,omitempty"`
	LocalPreference   *int64                                                                           `xml:"local-preference,omitempty"`
	Med               *int64                                                                           `xml:"med,omitempty"`
	Nexthop           *string                                                                          `xml:"nexthop,omitempty"`
	Origin            *string                                                                          `xml:"origin,omitempty"`
	Weight            *int64                                                                           `xml:"weight,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml struct {
	None    *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml   `xml:"none,omitempty"`
	Prepend *int64                                                                      `xml:"prepend,omitempty"`
	Remove  *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemoveXml `xml:"remove,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemoveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml struct {
	Append      *util.MemberType                                                                  `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                                  `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                           `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml struct {
	Append      *util.MemberType                                                                          `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                                          `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                                   `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersXml struct {
	Enable  *string                                                     `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchXml `xml:"match,omitempty"`
	XMLName xml.Name                                                    `xml:"entry"`
	Name    string                                                      `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                                             `xml:"from-peer,omitempty"`
	Med               *int64                                                                       `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                                             `xml:"nexthop,omitempty"`
	RouteTable        *string                                                                      `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml struct {
	Exact   *string  `xml:"exact,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementXml struct {
	Policy []ProtocolBgpPolicyConditionalAdvertisementPolicyXml `xml:"policy>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyXml struct {
	AdvertiseFilters []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml `xml:"advertise-filters>entry,omitempty"`
	Enable           *string                                                              `xml:"enable,omitempty"`
	XMLName          xml.Name                                                             `xml:"entry"`
	Name             string                                                               `xml:"name,attr"`
	NonExistFilters  []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml  `xml:"non-exist-filters>entry,omitempty"`
	UsedBy           *util.MemberType                                                     `xml:"used-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml struct {
	Enable  *string                                                                  `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml `xml:"match,omitempty"`
	XMLName xml.Name                                                                 `xml:"entry"`
	Name    string                                                                   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                                                          `xml:"from-peer,omitempty"`
	Med               *int64                                                                                    `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                                                          `xml:"nexthop,omitempty"`
	RouteTable        *string                                                                                   `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml struct {
	Enable  *string                                                                 `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml `xml:"match,omitempty"`
	XMLName xml.Name                                                                `xml:"entry"`
	Name    string                                                                  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                                                         `xml:"from-peer,omitempty"`
	Med               *int64                                                                                   `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                                                         `xml:"nexthop,omitempty"`
	RouteTable        *string                                                                                  `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportXml struct {
	Rules []ProtocolBgpPolicyExportRulesXml `xml:"rules>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesXml struct {
	Action  *ProtocolBgpPolicyExportRulesActionXml `xml:"action,omitempty"`
	Enable  *string                                `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyExportRulesMatchXml  `xml:"match,omitempty"`
	XMLName xml.Name                               `xml:"entry"`
	Name    string                                 `xml:"name,attr"`
	UsedBy  *util.MemberType                       `xml:"used-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionXml struct {
	Allow *ProtocolBgpPolicyExportRulesActionAllowXml `xml:"allow,omitempty"`
	Deny  *ProtocolBgpPolicyExportRulesActionDenyXml  `xml:"deny,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowXml struct {
	Update *ProtocolBgpPolicyExportRulesActionAllowUpdateXml `xml:"update,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateXml struct {
	AsPath            *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathXml            `xml:"as-path,omitempty"`
	AsPathLimit       *int64                                                             `xml:"as-path-limit,omitempty"`
	Community         *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml `xml:"extended-community,omitempty"`
	LocalPreference   *int64                                                             `xml:"local-preference,omitempty"`
	Med               *int64                                                             `xml:"med,omitempty"`
	Nexthop           *string                                                            `xml:"nexthop,omitempty"`
	Origin            *string                                                            `xml:"origin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathXml struct {
	None             *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml   `xml:"none,omitempty"`
	Prepend          *int64                                                        `xml:"prepend,omitempty"`
	Remove           *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml `xml:"remove,omitempty"`
	RemoveAndPrepend *int64                                                        `xml:"remove-and-prepend,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityXml struct {
	Append      *util.MemberType                                                    `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                    `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                             `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml struct {
	Append      *util.MemberType                                                            `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                            `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                     `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesActionDenyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyExportRulesMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyExportRulesMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyExportRulesMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyExportRulesMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                       `xml:"from-peer,omitempty"`
	Med               *int64                                                 `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                       `xml:"nexthop,omitempty"`
	RouteTable        *string                                                `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesMatchAddressPrefixXml struct {
	Exact   *string  `xml:"exact,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyExportRulesMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportXml struct {
	Rules []ProtocolBgpPolicyImportRulesXml `xml:"rules>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesXml struct {
	Action  *ProtocolBgpPolicyImportRulesActionXml `xml:"action,omitempty"`
	Enable  *string                                `xml:"enable,omitempty"`
	Match   *ProtocolBgpPolicyImportRulesMatchXml  `xml:"match,omitempty"`
	XMLName xml.Name                               `xml:"entry"`
	Name    string                                 `xml:"name,attr"`
	UsedBy  *util.MemberType                       `xml:"used-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionXml struct {
	Allow *ProtocolBgpPolicyImportRulesActionAllowXml `xml:"allow,omitempty"`
	Deny  *ProtocolBgpPolicyImportRulesActionDenyXml  `xml:"deny,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowXml struct {
	Dampening *string                                           `xml:"dampening,omitempty"`
	Update    *ProtocolBgpPolicyImportRulesActionAllowUpdateXml `xml:"update,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateXml struct {
	AsPath            *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathXml            `xml:"as-path,omitempty"`
	AsPathLimit       *int64                                                             `xml:"as-path-limit,omitempty"`
	Community         *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml `xml:"extended-community,omitempty"`
	LocalPreference   *int64                                                             `xml:"local-preference,omitempty"`
	Med               *int64                                                             `xml:"med,omitempty"`
	Nexthop           *string                                                            `xml:"nexthop,omitempty"`
	Origin            *string                                                            `xml:"origin,omitempty"`
	Weight            *int64                                                             `xml:"weight,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathXml struct {
	None   *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml   `xml:"none,omitempty"`
	Remove *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml `xml:"remove,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityXml struct {
	Append      *util.MemberType                                                    `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                    `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                             `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml struct {
	Append      *util.MemberType                                                            `xml:"append,omitempty"`
	None        *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml      `xml:"none,omitempty"`
	Overwrite   *util.MemberType                                                            `xml:"overwrite,omitempty"`
	RemoveAll   *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml `xml:"remove-all,omitempty"`
	RemoveRegex *string                                                                     `xml:"remove-regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesActionDenyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesMatchXml struct {
	AddressPrefix     []ProtocolBgpPolicyImportRulesMatchAddressPrefixXml    `xml:"address-prefix>entry,omitempty"`
	AsPath            *ProtocolBgpPolicyImportRulesMatchAsPathXml            `xml:"as-path,omitempty"`
	Community         *ProtocolBgpPolicyImportRulesMatchCommunityXml         `xml:"community,omitempty"`
	ExtendedCommunity *ProtocolBgpPolicyImportRulesMatchExtendedCommunityXml `xml:"extended-community,omitempty"`
	FromPeer          *util.MemberType                                       `xml:"from-peer,omitempty"`
	Med               *int64                                                 `xml:"med,omitempty"`
	Nexthop           *util.MemberType                                       `xml:"nexthop,omitempty"`
	RouteTable        *string                                                `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesMatchAddressPrefixXml struct {
	Exact   *string  `xml:"exact,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesMatchAsPathXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesMatchCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpPolicyImportRulesMatchExtendedCommunityXml struct {
	Regex *string `xml:"regex,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpRedistRulesXml struct {
	AddressFamilyIdentifier *string          `xml:"address-family-identifier,omitempty"`
	Enable                  *string          `xml:"enable,omitempty"`
	Metric                  *int64           `xml:"metric,omitempty"`
	XMLName                 xml.Name         `xml:"entry"`
	Name                    string           `xml:"name,attr"`
	RouteTable              *string          `xml:"route-table,omitempty"`
	SetAsPathLimit          *int64           `xml:"set-as-path-limit,omitempty"`
	SetCommunity            *util.MemberType `xml:"set-community,omitempty"`
	SetExtendedCommunity    *util.MemberType `xml:"set-extended-community,omitempty"`
	SetLocalPreference      *int64           `xml:"set-local-preference,omitempty"`
	SetMed                  *int64           `xml:"set-med,omitempty"`
	SetOrigin               *string          `xml:"set-origin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpRoutingOptionsXml struct {
	Aggregate              *ProtocolBgpRoutingOptionsAggregateXml       `xml:"aggregate,omitempty"`
	AsFormat               *string                                      `xml:"as-format,omitempty"`
	ConfederationMemberAs  *string                                      `xml:"confederation-member-as,omitempty"`
	DefaultLocalPreference *int64                                       `xml:"default-local-preference,omitempty"`
	GracefulRestart        *ProtocolBgpRoutingOptionsGracefulRestartXml `xml:"graceful-restart,omitempty"`
	Med                    *ProtocolBgpRoutingOptionsMedXml             `xml:"med,omitempty"`
	ReflectorClusterId     *string                                      `xml:"reflector-cluster-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpRoutingOptionsAggregateXml struct {
	AggregateMed *string `xml:"aggregate-med,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpRoutingOptionsGracefulRestartXml struct {
	Enable             *string `xml:"enable,omitempty"`
	LocalRestartTime   *int64  `xml:"local-restart-time,omitempty"`
	MaxPeerRestartTime *int64  `xml:"max-peer-restart-time,omitempty"`
	StaleRouteTime     *int64  `xml:"stale-route-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpRoutingOptionsMedXml struct {
	AlwaysCompareMed           *string `xml:"always-compare-med,omitempty"`
	DeterministicMedComparison *string `xml:"deterministic-med-comparison,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfXml struct {
	AllowRedistDefaultRoute *string                         `xml:"allow-redist-default-route,omitempty"`
	Area                    []ProtocolOspfAreaXml           `xml:"area>entry,omitempty"`
	AuthProfile             []ProtocolOspfAuthProfileXml    `xml:"auth-profile>entry,omitempty"`
	Enable                  *string                         `xml:"enable,omitempty"`
	ExportRules             []ProtocolOspfExportRulesXml    `xml:"export-rules>entry,omitempty"`
	GlobalBfd               *ProtocolOspfGlobalBfdXml       `xml:"global-bfd,omitempty"`
	GracefulRestart         *ProtocolOspfGracefulRestartXml `xml:"graceful-restart,omitempty"`
	RejectDefaultRoute      *string                         `xml:"reject-default-route,omitempty"`
	Rfc1583                 *string                         `xml:"rfc1583,omitempty"`
	RouterId                *string                         `xml:"router-id,omitempty"`
	Timers                  *ProtocolOspfTimersXml          `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaXml struct {
	Interface   []ProtocolOspfAreaInterfaceXml   `xml:"interface>entry,omitempty"`
	XMLName     xml.Name                         `xml:"entry"`
	Name        string                           `xml:"name,attr"`
	Range       []ProtocolOspfAreaRangeXml       `xml:"range>entry,omitempty"`
	Type        *ProtocolOspfAreaTypeXml         `xml:"type,omitempty"`
	VirtualLink []ProtocolOspfAreaVirtualLinkXml `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceXml struct {
	Authentication     *string                                `xml:"authentication,omitempty"`
	Bfd                *ProtocolOspfAreaInterfaceBfdXml       `xml:"bfd,omitempty"`
	DeadCounts         *int64                                 `xml:"dead-counts,omitempty"`
	Enable             *string                                `xml:"enable,omitempty"`
	GrDelay            *int64                                 `xml:"gr-delay,omitempty"`
	HelloInterval      *int64                                 `xml:"hello-interval,omitempty"`
	LinkType           *ProtocolOspfAreaInterfaceLinkTypeXml  `xml:"link-type,omitempty"`
	Metric             *int64                                 `xml:"metric,omitempty"`
	XMLName            xml.Name                               `xml:"entry"`
	Name               string                                 `xml:"name,attr"`
	Neighbor           []ProtocolOspfAreaInterfaceNeighborXml `xml:"neighbor>entry,omitempty"`
	Passive            *string                                `xml:"passive,omitempty"`
	Priority           *int64                                 `xml:"priority,omitempty"`
	RetransmitInterval *int64                                 `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                                 `xml:"transit-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceLinkTypeXml struct {
	Broadcast *ProtocolOspfAreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2mp      *ProtocolOspfAreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	P2p       *ProtocolOspfAreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceLinkTypeP2mpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaInterfaceNeighborXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaRangeXml struct {
	XMLName   xml.Name                           `xml:"entry"`
	Name      string                             `xml:"name,attr"`
	Advertise *ProtocolOspfAreaRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *ProtocolOspfAreaRangeSuppressXml  `xml:"suppress,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeXml struct {
	Normal *ProtocolOspfAreaTypeNormalXml `xml:"normal,omitempty"`
	Nssa   *ProtocolOspfAreaTypeNssaXml   `xml:"nssa,omitempty"`
	Stub   *ProtocolOspfAreaTypeStubXml   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNormalXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaXml struct {
	AcceptSummary *string                                   `xml:"accept-summary,omitempty"`
	DefaultRoute  *ProtocolOspfAreaTypeNssaDefaultRouteXml  `xml:"default-route,omitempty"`
	NssaExtRange  []ProtocolOspfAreaTypeNssaNssaExtRangeXml `xml:"nssa-ext-range>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaDefaultRouteXml struct {
	Advertise *ProtocolOspfAreaTypeNssaDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Disable   *ProtocolOspfAreaTypeNssaDefaultRouteDisableXml   `xml:"disable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaDefaultRouteAdvertiseXml struct {
	Metric *int64  `xml:"metric,omitempty"`
	Type   *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaNssaExtRangeXml struct {
	XMLName   xml.Name                                          `xml:"entry"`
	Name      string                                            `xml:"name,attr"`
	Advertise *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *ProtocolOspfAreaTypeNssaNssaExtRangeSuppressXml  `xml:"suppress,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeNssaNssaExtRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeStubXml struct {
	AcceptSummary *string                                  `xml:"accept-summary,omitempty"`
	DefaultRoute  *ProtocolOspfAreaTypeStubDefaultRouteXml `xml:"default-route,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeStubDefaultRouteXml struct {
	Advertise *ProtocolOspfAreaTypeStubDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Disable   *ProtocolOspfAreaTypeStubDefaultRouteDisableXml   `xml:"disable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeStubDefaultRouteAdvertiseXml struct {
	Metric *int64 `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaTypeStubDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaVirtualLinkXml struct {
	Authentication     *string                            `xml:"authentication,omitempty"`
	Bfd                *ProtocolOspfAreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	DeadCounts         *int64                             `xml:"dead-counts,omitempty"`
	Enable             *string                            `xml:"enable,omitempty"`
	HelloInterval      *int64                             `xml:"hello-interval,omitempty"`
	XMLName            xml.Name                           `xml:"entry"`
	Name               string                             `xml:"name,attr"`
	NeighborId         *string                            `xml:"neighbor-id,omitempty"`
	RetransmitInterval *int64                             `xml:"retransmit-interval,omitempty"`
	TransitAreaId      *string                            `xml:"transit-area-id,omitempty"`
	TransitDelay       *int64                             `xml:"transit-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAreaVirtualLinkBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAuthProfileXml struct {
	XMLName  xml.Name                        `xml:"entry"`
	Name     string                          `xml:"name,attr"`
	Md5      []ProtocolOspfAuthProfileMd5Xml `xml:"md5>entry,omitempty"`
	Password *string                         `xml:"password,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfAuthProfileMd5Xml struct {
	Key       *string  `xml:"key,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`
	Preferred *string  `xml:"preferred,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfExportRulesXml struct {
	Metric      *int64   `xml:"metric,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	NewPathType *string  `xml:"new-path-type,omitempty"`
	NewTag      *string  `xml:"new-tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfGracefulRestartXml struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfTimersXml struct {
	LsaInterval         *float64 `xml:"lsa-interval,omitempty"`
	SpfCalculationDelay *float64 `xml:"spf-calculation-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3Xml struct {
	AllowRedistDefaultRoute *string                           `xml:"allow-redist-default-route,omitempty"`
	Area                    []ProtocolOspfv3AreaXml           `xml:"area>entry,omitempty"`
	AuthProfile             []ProtocolOspfv3AuthProfileXml    `xml:"auth-profile>entry,omitempty"`
	DisableTransitTraffic   *string                           `xml:"disable-transit-traffic,omitempty"`
	Enable                  *string                           `xml:"enable,omitempty"`
	ExportRules             []ProtocolOspfv3ExportRulesXml    `xml:"export-rules>entry,omitempty"`
	GlobalBfd               *ProtocolOspfv3GlobalBfdXml       `xml:"global-bfd,omitempty"`
	GracefulRestart         *ProtocolOspfv3GracefulRestartXml `xml:"graceful-restart,omitempty"`
	RejectDefaultRoute      *string                           `xml:"reject-default-route,omitempty"`
	RouterId                *string                           `xml:"router-id,omitempty"`
	Timers                  *ProtocolOspfv3TimersXml          `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaXml struct {
	Authentication *string                            `xml:"authentication,omitempty"`
	Interface      []ProtocolOspfv3AreaInterfaceXml   `xml:"interface>entry,omitempty"`
	XMLName        xml.Name                           `xml:"entry"`
	Name           string                             `xml:"name,attr"`
	Range          []ProtocolOspfv3AreaRangeXml       `xml:"range>entry,omitempty"`
	Type           *ProtocolOspfv3AreaTypeXml         `xml:"type,omitempty"`
	VirtualLink    []ProtocolOspfv3AreaVirtualLinkXml `xml:"virtual-link>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceXml struct {
	Authentication     *string                                  `xml:"authentication,omitempty"`
	Bfd                *ProtocolOspfv3AreaInterfaceBfdXml       `xml:"bfd,omitempty"`
	DeadCounts         *int64                                   `xml:"dead-counts,omitempty"`
	Enable             *string                                  `xml:"enable,omitempty"`
	GrDelay            *int64                                   `xml:"gr-delay,omitempty"`
	HelloInterval      *int64                                   `xml:"hello-interval,omitempty"`
	InstanceId         *int64                                   `xml:"instance-id,omitempty"`
	LinkType           *ProtocolOspfv3AreaInterfaceLinkTypeXml  `xml:"link-type,omitempty"`
	Metric             *int64                                   `xml:"metric,omitempty"`
	XMLName            xml.Name                                 `xml:"entry"`
	Name               string                                   `xml:"name,attr"`
	Neighbor           []ProtocolOspfv3AreaInterfaceNeighborXml `xml:"neighbor>entry,omitempty"`
	Passive            *string                                  `xml:"passive,omitempty"`
	Priority           *int64                                   `xml:"priority,omitempty"`
	RetransmitInterval *int64                                   `xml:"retransmit-interval,omitempty"`
	TransitDelay       *int64                                   `xml:"transit-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceLinkTypeXml struct {
	Broadcast *ProtocolOspfv3AreaInterfaceLinkTypeBroadcastXml `xml:"broadcast,omitempty"`
	P2mp      *ProtocolOspfv3AreaInterfaceLinkTypeP2mpXml      `xml:"p2mp,omitempty"`
	P2p       *ProtocolOspfv3AreaInterfaceLinkTypeP2pXml       `xml:"p2p,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceLinkTypeBroadcastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2mpXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceLinkTypeP2pXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaInterfaceNeighborXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaRangeXml struct {
	XMLName   xml.Name                             `xml:"entry"`
	Name      string                               `xml:"name,attr"`
	Advertise *ProtocolOspfv3AreaRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *ProtocolOspfv3AreaRangeSuppressXml  `xml:"suppress,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeXml struct {
	Normal *ProtocolOspfv3AreaTypeNormalXml `xml:"normal,omitempty"`
	Nssa   *ProtocolOspfv3AreaTypeNssaXml   `xml:"nssa,omitempty"`
	Stub   *ProtocolOspfv3AreaTypeStubXml   `xml:"stub,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNormalXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaXml struct {
	AcceptSummary *string                                     `xml:"accept-summary,omitempty"`
	DefaultRoute  *ProtocolOspfv3AreaTypeNssaDefaultRouteXml  `xml:"default-route,omitempty"`
	NssaExtRange  []ProtocolOspfv3AreaTypeNssaNssaExtRangeXml `xml:"nssa-ext-range>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteXml struct {
	Advertise *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Disable   *ProtocolOspfv3AreaTypeNssaDefaultRouteDisableXml   `xml:"disable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml struct {
	Metric *int64  `xml:"metric,omitempty"`
	Type   *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeXml struct {
	XMLName   xml.Name                                            `xml:"entry"`
	Name      string                                              `xml:"name,attr"`
	Advertise *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml `xml:"advertise,omitempty"`
	Suppress  *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml  `xml:"suppress,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeStubXml struct {
	AcceptSummary *string                                    `xml:"accept-summary,omitempty"`
	DefaultRoute  *ProtocolOspfv3AreaTypeStubDefaultRouteXml `xml:"default-route,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeStubDefaultRouteXml struct {
	Advertise *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Disable   *ProtocolOspfv3AreaTypeStubDefaultRouteDisableXml   `xml:"disable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml struct {
	Metric *int64 `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaTypeStubDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaVirtualLinkXml struct {
	Authentication     *string                              `xml:"authentication,omitempty"`
	Bfd                *ProtocolOspfv3AreaVirtualLinkBfdXml `xml:"bfd,omitempty"`
	DeadCounts         *int64                               `xml:"dead-counts,omitempty"`
	Enable             *string                              `xml:"enable,omitempty"`
	HelloInterval      *int64                               `xml:"hello-interval,omitempty"`
	InstanceId         *int64                               `xml:"instance-id,omitempty"`
	XMLName            xml.Name                             `xml:"entry"`
	Name               string                               `xml:"name,attr"`
	NeighborId         *string                              `xml:"neighbor-id,omitempty"`
	RetransmitInterval *int64                               `xml:"retransmit-interval,omitempty"`
	TransitAreaId      *string                              `xml:"transit-area-id,omitempty"`
	TransitDelay       *int64                               `xml:"transit-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AreaVirtualLinkBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileXml struct {
	XMLName xml.Name                         `xml:"entry"`
	Name    string                           `xml:"name,attr"`
	Spi     *string                          `xml:"spi,omitempty"`
	Ah      *ProtocolOspfv3AuthProfileAhXml  `xml:"ah,omitempty"`
	Esp     *ProtocolOspfv3AuthProfileEspXml `xml:"esp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhXml struct {
	Md5    *ProtocolOspfv3AuthProfileAhMd5Xml    `xml:"md5,omitempty"`
	Sha1   *ProtocolOspfv3AuthProfileAhSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *ProtocolOspfv3AuthProfileAhSha256Xml `xml:"sha256,omitempty"`
	Sha384 *ProtocolOspfv3AuthProfileAhSha384Xml `xml:"sha384,omitempty"`
	Sha512 *ProtocolOspfv3AuthProfileAhSha512Xml `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhMd5Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhSha1Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhSha256Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhSha384Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileAhSha512Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspXml struct {
	Authentication *ProtocolOspfv3AuthProfileEspAuthenticationXml `xml:"authentication,omitempty"`
	Encryption     *ProtocolOspfv3AuthProfileEspEncryptionXml     `xml:"encryption,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationXml struct {
	Md5    *ProtocolOspfv3AuthProfileEspAuthenticationMd5Xml    `xml:"md5,omitempty"`
	None   *ProtocolOspfv3AuthProfileEspAuthenticationNoneXml   `xml:"none,omitempty"`
	Sha1   *ProtocolOspfv3AuthProfileEspAuthenticationSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *ProtocolOspfv3AuthProfileEspAuthenticationSha256Xml `xml:"sha256,omitempty"`
	Sha384 *ProtocolOspfv3AuthProfileEspAuthenticationSha384Xml `xml:"sha384,omitempty"`
	Sha512 *ProtocolOspfv3AuthProfileEspAuthenticationSha512Xml `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationMd5Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha1Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha256Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha384Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspAuthenticationSha512Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3AuthProfileEspEncryptionXml struct {
	Algorithm *string `xml:"algorithm,omitempty"`
	Key       *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3ExportRulesXml struct {
	Metric      *int64   `xml:"metric,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	NewPathType *string  `xml:"new-path-type,omitempty"`
	NewTag      *string  `xml:"new-tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3GlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3GracefulRestartXml struct {
	Enable                 *string `xml:"enable,omitempty"`
	GracePeriod            *int64  `xml:"grace-period,omitempty"`
	HelperEnable           *string `xml:"helper-enable,omitempty"`
	MaxNeighborRestartTime *int64  `xml:"max-neighbor-restart-time,omitempty"`
	StrictLSAChecking      *string `xml:"strict-LSA-checking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3TimersXml struct {
	LsaInterval         *float64 `xml:"lsa-interval,omitempty"`
	SpfCalculationDelay *float64 `xml:"spf-calculation-delay,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileXml struct {
	Action   *ProtocolRedistProfileActionXml `xml:"action,omitempty"`
	Filter   *ProtocolRedistProfileFilterXml `xml:"filter,omitempty"`
	XMLName  xml.Name                        `xml:"entry"`
	Name     string                          `xml:"name,attr"`
	Priority *int64                          `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileActionXml struct {
	NoRedist *ProtocolRedistProfileActionNoRedistXml `xml:"no-redist,omitempty"`
	Redist   *ProtocolRedistProfileActionRedistXml   `xml:"redist,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileActionNoRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileActionRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileFilterXml struct {
	Bgp         *ProtocolRedistProfileFilterBgpXml  `xml:"bgp,omitempty"`
	Destination *util.MemberType                    `xml:"destination,omitempty"`
	Interface   *util.MemberType                    `xml:"interface,omitempty"`
	Nexthop     *util.MemberType                    `xml:"nexthop,omitempty"`
	Ospf        *ProtocolRedistProfileFilterOspfXml `xml:"ospf,omitempty"`
	Type        *util.MemberType                    `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileFilterBgpXml struct {
	Community         *util.MemberType `xml:"community,omitempty"`
	ExtendedCommunity *util.MemberType `xml:"extended-community,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileFilterOspfXml struct {
	Area     *util.MemberType `xml:"area,omitempty"`
	PathType *util.MemberType `xml:"path-type,omitempty"`
	Tag      *util.MemberType `xml:"tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6Xml struct {
	Action   *ProtocolRedistProfileIpv6ActionXml `xml:"action,omitempty"`
	Filter   *ProtocolRedistProfileIpv6FilterXml `xml:"filter,omitempty"`
	XMLName  xml.Name                            `xml:"entry"`
	Name     string                              `xml:"name,attr"`
	Priority *int64                              `xml:"priority,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6ActionXml struct {
	NoRedist *ProtocolRedistProfileIpv6ActionNoRedistXml `xml:"no-redist,omitempty"`
	Redist   *ProtocolRedistProfileIpv6ActionRedistXml   `xml:"redist,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6ActionNoRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6ActionRedistXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6FilterXml struct {
	Bgp         *ProtocolRedistProfileIpv6FilterBgpXml    `xml:"bgp,omitempty"`
	Destination *util.MemberType                          `xml:"destination,omitempty"`
	Interface   *util.MemberType                          `xml:"interface,omitempty"`
	Nexthop     *util.MemberType                          `xml:"nexthop,omitempty"`
	Ospfv3      *ProtocolRedistProfileIpv6FilterOspfv3Xml `xml:"ospfv3,omitempty"`
	Type        *util.MemberType                          `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6FilterBgpXml struct {
	Community         *util.MemberType `xml:"community,omitempty"`
	ExtendedCommunity *util.MemberType `xml:"extended-community,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRedistProfileIpv6FilterOspfv3Xml struct {
	Area     *util.MemberType `xml:"area,omitempty"`
	PathType *util.MemberType `xml:"path-type,omitempty"`
	Tag      *util.MemberType `xml:"tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipXml struct {
	AllowRedistDefaultRoute *string                     `xml:"allow-redist-default-route,omitempty"`
	AuthProfile             []ProtocolRipAuthProfileXml `xml:"auth-profile>entry,omitempty"`
	Enable                  *string                     `xml:"enable,omitempty"`
	ExportRules             []ProtocolRipExportRulesXml `xml:"export-rules>entry,omitempty"`
	GlobalBfd               *ProtocolRipGlobalBfdXml    `xml:"global-bfd,omitempty"`
	Interface               []ProtocolRipInterfaceXml   `xml:"interface>entry,omitempty"`
	RejectDefaultRoute      *string                     `xml:"reject-default-route,omitempty"`
	Timers                  *ProtocolRipTimersXml       `xml:"timers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipAuthProfileXml struct {
	XMLName  xml.Name                       `xml:"entry"`
	Name     string                         `xml:"name,attr"`
	Md5      []ProtocolRipAuthProfileMd5Xml `xml:"md5>entry,omitempty"`
	Password *string                        `xml:"password,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipAuthProfileMd5Xml struct {
	Key       *string  `xml:"key,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`
	Preferred *string  `xml:"preferred,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipExportRulesXml struct {
	Metric  *int64   `xml:"metric,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipGlobalBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipInterfaceXml struct {
	Authentication *string                              `xml:"authentication,omitempty"`
	Bfd            *ProtocolRipInterfaceBfdXml          `xml:"bfd,omitempty"`
	DefaultRoute   *ProtocolRipInterfaceDefaultRouteXml `xml:"default-route,omitempty"`
	Enable         *string                              `xml:"enable,omitempty"`
	Mode           *string                              `xml:"mode,omitempty"`
	XMLName        xml.Name                             `xml:"entry"`
	Name           string                               `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipInterfaceBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipInterfaceDefaultRouteXml struct {
	Advertise *ProtocolRipInterfaceDefaultRouteAdvertiseXml `xml:"advertise,omitempty"`
	Disable   *ProtocolRipInterfaceDefaultRouteDisableXml   `xml:"disable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipInterfaceDefaultRouteAdvertiseXml struct {
	Metric *int64 `xml:"metric,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipInterfaceDefaultRouteDisableXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipTimersXml struct {
	DeleteIntervals *int64 `xml:"delete-intervals,omitempty"`
	ExpireIntervals *int64 `xml:"expire-intervals,omitempty"`
	IntervalSeconds *int64 `xml:"interval-seconds,omitempty"`
	UpdateIntervals *int64 `xml:"update-intervals,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableXml struct {
	Ip   *RoutingTableIpXml   `xml:"ip,omitempty"`
	Ipv6 *RoutingTableIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpXml struct {
	StaticRoute []RoutingTableIpStaticRouteXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteXml struct {
	AdminDist   *int64                                   `xml:"admin-dist,omitempty"`
	Bfd         *RoutingTableIpStaticRouteBfdXml         `xml:"bfd,omitempty"`
	Destination *string                                  `xml:"destination,omitempty"`
	Interface   *string                                  `xml:"interface,omitempty"`
	Metric      *int64                                   `xml:"metric,omitempty"`
	XMLName     xml.Name                                 `xml:"entry"`
	Name        string                                   `xml:"name,attr"`
	Nexthop     *RoutingTableIpStaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	PathMonitor *RoutingTableIpStaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`
	RouteTable  *RoutingTableIpStaticRouteRouteTableXml  `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteNexthopXml struct {
	Discard   *RoutingTableIpStaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	Fqdn      *string                                     `xml:"fqdn,omitempty"`
	IpAddress *string                                     `xml:"ip-address,omitempty"`
	NextVr    *string                                     `xml:"next-vr,omitempty"`
	Receive   *RoutingTableIpStaticRouteNexthopReceiveXml `xml:"receive,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteNexthopReceiveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRoutePathMonitorXml struct {
	Enable              *string                                                      `xml:"enable,omitempty"`
	FailureCondition    *string                                                      `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                       `xml:"hold-time,omitempty"`
	MonitorDestinations []RoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml struct {
	Count       *int64   `xml:"count,omitempty"`
	Destination *string  `xml:"destination,omitempty"`
	Enable      *string  `xml:"enable,omitempty"`
	Interval    *int64   `xml:"interval,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Source      *string  `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteRouteTableXml struct {
	Both      *RoutingTableIpStaticRouteRouteTableBothXml      `xml:"both,omitempty"`
	Multicast *RoutingTableIpStaticRouteRouteTableMulticastXml `xml:"multicast,omitempty"`
	NoInstall *RoutingTableIpStaticRouteRouteTableNoInstallXml `xml:"no-install,omitempty"`
	Unicast   *RoutingTableIpStaticRouteRouteTableUnicastXml   `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteRouteTableBothXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteRouteTableMulticastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteRouteTableNoInstallXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRouteRouteTableUnicastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6Xml struct {
	StaticRoute []RoutingTableIpv6StaticRouteXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteXml struct {
	AdminDist   *int64                                     `xml:"admin-dist,omitempty"`
	Bfd         *RoutingTableIpv6StaticRouteBfdXml         `xml:"bfd,omitempty"`
	Destination *string                                    `xml:"destination,omitempty"`
	Interface   *string                                    `xml:"interface,omitempty"`
	Metric      *int64                                     `xml:"metric,omitempty"`
	XMLName     xml.Name                                   `xml:"entry"`
	Name        string                                     `xml:"name,attr"`
	Nexthop     *RoutingTableIpv6StaticRouteNexthopXml     `xml:"nexthop,omitempty"`
	Option      *RoutingTableIpv6StaticRouteOptionXml      `xml:"option,omitempty"`
	PathMonitor *RoutingTableIpv6StaticRoutePathMonitorXml `xml:"path-monitor,omitempty"`
	RouteTable  *RoutingTableIpv6StaticRouteRouteTableXml  `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteBfdXml struct {
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteNexthopXml struct {
	Discard     *RoutingTableIpv6StaticRouteNexthopDiscardXml `xml:"discard,omitempty"`
	Ipv6Address *string                                       `xml:"ipv6-address,omitempty"`
	NextVr      *string                                       `xml:"next-vr,omitempty"`
	Receive     *RoutingTableIpv6StaticRouteNexthopReceiveXml `xml:"receive,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteNexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteNexthopReceiveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteOptionXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRoutePathMonitorXml struct {
	Enable              *string                                                        `xml:"enable,omitempty"`
	FailureCondition    *string                                                        `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                                         `xml:"hold-time,omitempty"`
	MonitorDestinations []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml `xml:"monitor-destinations>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml struct {
	Count       *int64   `xml:"count,omitempty"`
	Destination *string  `xml:"destination,omitempty"`
	Enable      *string  `xml:"enable,omitempty"`
	Interval    *int64   `xml:"interval,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Source      *string  `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteRouteTableXml struct {
	NoInstall *RoutingTableIpv6StaticRouteRouteTableNoInstallXml `xml:"no-install,omitempty"`
	Unicast   *RoutingTableIpv6StaticRouteRouteTableUnicastXml   `xml:"unicast,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteRouteTableNoInstallXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRouteRouteTableUnicastXml struct {
	Misc []generic.Xml `xml:",any"`
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
	if v == "routing_table" || v == "RoutingTable" {
		return e.RoutingTable, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedAdminDists *AdminDistsXml
	if o.AdminDists != nil {
		nestedAdminDists = &AdminDistsXml{}
		if _, ok := o.Misc["AdminDists"]; ok {
			nestedAdminDists.Misc = o.Misc["AdminDists"]
		}
		if o.AdminDists.OspfExt != nil {
			nestedAdminDists.OspfExt = o.AdminDists.OspfExt
		}
		if o.AdminDists.Ospfv3Ext != nil {
			nestedAdminDists.Ospfv3Ext = o.AdminDists.Ospfv3Ext
		}
		if o.AdminDists.Ospfv3Int != nil {
			nestedAdminDists.Ospfv3Int = o.AdminDists.Ospfv3Int
		}
		if o.AdminDists.Rip != nil {
			nestedAdminDists.Rip = o.AdminDists.Rip
		}
		if o.AdminDists.Ebgp != nil {
			nestedAdminDists.Ebgp = o.AdminDists.Ebgp
		}
		if o.AdminDists.Ibgp != nil {
			nestedAdminDists.Ibgp = o.AdminDists.Ibgp
		}
		if o.AdminDists.OspfInt != nil {
			nestedAdminDists.OspfInt = o.AdminDists.OspfInt
		}
		if o.AdminDists.Static != nil {
			nestedAdminDists.Static = o.AdminDists.Static
		}
		if o.AdminDists.StaticIpv6 != nil {
			nestedAdminDists.StaticIpv6 = o.AdminDists.StaticIpv6
		}
	}
	entry.AdminDists = nestedAdminDists

	var nestedEcmp *EcmpXml
	if o.Ecmp != nil {
		nestedEcmp = &EcmpXml{}
		if _, ok := o.Misc["Ecmp"]; ok {
			nestedEcmp.Misc = o.Misc["Ecmp"]
		}
		if o.Ecmp.Algorithm != nil {
			nestedEcmp.Algorithm = &EcmpAlgorithmXml{}
			if _, ok := o.Misc["EcmpAlgorithm"]; ok {
				nestedEcmp.Algorithm.Misc = o.Misc["EcmpAlgorithm"]
			}
			if o.Ecmp.Algorithm.BalancedRoundRobin != nil {
				nestedEcmp.Algorithm.BalancedRoundRobin = &EcmpAlgorithmBalancedRoundRobinXml{}
				if _, ok := o.Misc["EcmpAlgorithmBalancedRoundRobin"]; ok {
					nestedEcmp.Algorithm.BalancedRoundRobin.Misc = o.Misc["EcmpAlgorithmBalancedRoundRobin"]
				}
			}
			if o.Ecmp.Algorithm.IpHash != nil {
				nestedEcmp.Algorithm.IpHash = &EcmpAlgorithmIpHashXml{}
				if _, ok := o.Misc["EcmpAlgorithmIpHash"]; ok {
					nestedEcmp.Algorithm.IpHash.Misc = o.Misc["EcmpAlgorithmIpHash"]
				}
				if o.Ecmp.Algorithm.IpHash.HashSeed != nil {
					nestedEcmp.Algorithm.IpHash.HashSeed = o.Ecmp.Algorithm.IpHash.HashSeed
				}
				if o.Ecmp.Algorithm.IpHash.SrcOnly != nil {
					nestedEcmp.Algorithm.IpHash.SrcOnly = util.YesNo(o.Ecmp.Algorithm.IpHash.SrcOnly, nil)
				}
				if o.Ecmp.Algorithm.IpHash.UsePort != nil {
					nestedEcmp.Algorithm.IpHash.UsePort = util.YesNo(o.Ecmp.Algorithm.IpHash.UsePort, nil)
				}
			}
			if o.Ecmp.Algorithm.IpModulo != nil {
				nestedEcmp.Algorithm.IpModulo = &EcmpAlgorithmIpModuloXml{}
				if _, ok := o.Misc["EcmpAlgorithmIpModulo"]; ok {
					nestedEcmp.Algorithm.IpModulo.Misc = o.Misc["EcmpAlgorithmIpModulo"]
				}
			}
			if o.Ecmp.Algorithm.WeightedRoundRobin != nil {
				nestedEcmp.Algorithm.WeightedRoundRobin = &EcmpAlgorithmWeightedRoundRobinXml{}
				if _, ok := o.Misc["EcmpAlgorithmWeightedRoundRobin"]; ok {
					nestedEcmp.Algorithm.WeightedRoundRobin.Misc = o.Misc["EcmpAlgorithmWeightedRoundRobin"]
				}
				if o.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
					nestedEcmp.Algorithm.WeightedRoundRobin.Interface = []EcmpAlgorithmWeightedRoundRobinInterfaceXml{}
					for _, oEcmpAlgorithmWeightedRoundRobinInterface := range o.Ecmp.Algorithm.WeightedRoundRobin.Interface {
						nestedEcmpAlgorithmWeightedRoundRobinInterface := EcmpAlgorithmWeightedRoundRobinInterfaceXml{}
						if _, ok := o.Misc["EcmpAlgorithmWeightedRoundRobinInterface"]; ok {
							nestedEcmpAlgorithmWeightedRoundRobinInterface.Misc = o.Misc["EcmpAlgorithmWeightedRoundRobinInterface"]
						}
						if oEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
							nestedEcmpAlgorithmWeightedRoundRobinInterface.Name = oEcmpAlgorithmWeightedRoundRobinInterface.Name
						}
						if oEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
							nestedEcmpAlgorithmWeightedRoundRobinInterface.Weight = oEcmpAlgorithmWeightedRoundRobinInterface.Weight
						}
						nestedEcmp.Algorithm.WeightedRoundRobin.Interface = append(nestedEcmp.Algorithm.WeightedRoundRobin.Interface, nestedEcmpAlgorithmWeightedRoundRobinInterface)
					}
				}
			}
		}
		if o.Ecmp.Enable != nil {
			nestedEcmp.Enable = util.YesNo(o.Ecmp.Enable, nil)
		}
		if o.Ecmp.MaxPath != nil {
			nestedEcmp.MaxPath = o.Ecmp.MaxPath
		}
		if o.Ecmp.StrictSourcePath != nil {
			nestedEcmp.StrictSourcePath = util.YesNo(o.Ecmp.StrictSourcePath, nil)
		}
		if o.Ecmp.SymmetricReturn != nil {
			nestedEcmp.SymmetricReturn = util.YesNo(o.Ecmp.SymmetricReturn, nil)
		}
	}
	entry.Ecmp = nestedEcmp

	entry.Interface = util.StrToMem(o.Interface)
	var nestedMulticast *MulticastXml
	if o.Multicast != nil {
		nestedMulticast = &MulticastXml{}
		if _, ok := o.Misc["Multicast"]; ok {
			nestedMulticast.Misc = o.Misc["Multicast"]
		}
		if o.Multicast.InterfaceGroup != nil {
			nestedMulticast.InterfaceGroup = []MulticastInterfaceGroupXml{}
			for _, oMulticastInterfaceGroup := range o.Multicast.InterfaceGroup {
				nestedMulticastInterfaceGroup := MulticastInterfaceGroupXml{}
				if _, ok := o.Misc["MulticastInterfaceGroup"]; ok {
					nestedMulticastInterfaceGroup.Misc = o.Misc["MulticastInterfaceGroup"]
				}
				if oMulticastInterfaceGroup.GroupPermission != nil {
					nestedMulticastInterfaceGroup.GroupPermission = &MulticastInterfaceGroupGroupPermissionXml{}
					if _, ok := o.Misc["MulticastInterfaceGroupGroupPermission"]; ok {
						nestedMulticastInterfaceGroup.GroupPermission.Misc = o.Misc["MulticastInterfaceGroupGroupPermission"]
					}
					if oMulticastInterfaceGroup.GroupPermission.AnySourceMulticast != nil {
						nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast = []MulticastInterfaceGroupGroupPermissionAnySourceMulticastXml{}
						for _, oMulticastInterfaceGroupGroupPermissionAnySourceMulticast := range oMulticastInterfaceGroup.GroupPermission.AnySourceMulticast {
							nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast := MulticastInterfaceGroupGroupPermissionAnySourceMulticastXml{}
							if _, ok := o.Misc["MulticastInterfaceGroupGroupPermissionAnySourceMulticast"]; ok {
								nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Misc = o.Misc["MulticastInterfaceGroupGroupPermissionAnySourceMulticast"]
							}
							if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress != nil {
								nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress = oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress
							}
							if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included != nil {
								nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included = util.YesNo(oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included, nil)
							}
							if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name != "" {
								nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name = oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name
							}
							nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast = append(nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast, nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast)
						}
					}
					if oMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast != nil {
						nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast = []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticastXml{}
						for _, oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast := range oMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast {
							nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast := MulticastInterfaceGroupGroupPermissionSourceSpecificMulticastXml{}
							if _, ok := o.Misc["MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast"]; ok {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Misc = o.Misc["MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast"]
							}
							if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress != nil {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress
							}
							if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress != nil {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress
							}
							if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included != nil {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included = util.YesNo(oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included, nil)
							}
							if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name != "" {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name
							}
							nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast = append(nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast, nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast)
						}
					}
				}
				if oMulticastInterfaceGroup.Igmp != nil {
					nestedMulticastInterfaceGroup.Igmp = &MulticastInterfaceGroupIgmpXml{}
					if _, ok := o.Misc["MulticastInterfaceGroupIgmp"]; ok {
						nestedMulticastInterfaceGroup.Igmp.Misc = o.Misc["MulticastInterfaceGroupIgmp"]
					}
					if oMulticastInterfaceGroup.Igmp.RouterAlertPolicing != nil {
						nestedMulticastInterfaceGroup.Igmp.RouterAlertPolicing = util.YesNo(oMulticastInterfaceGroup.Igmp.RouterAlertPolicing, nil)
					}
					if oMulticastInterfaceGroup.Igmp.Version != nil {
						nestedMulticastInterfaceGroup.Igmp.Version = oMulticastInterfaceGroup.Igmp.Version
					}
					if oMulticastInterfaceGroup.Igmp.MaxQueryResponseTime != nil {
						nestedMulticastInterfaceGroup.Igmp.MaxQueryResponseTime = oMulticastInterfaceGroup.Igmp.MaxQueryResponseTime
					}
					if oMulticastInterfaceGroup.Igmp.MaxGroups != nil {
						nestedMulticastInterfaceGroup.Igmp.MaxGroups = oMulticastInterfaceGroup.Igmp.MaxGroups
					}
					if oMulticastInterfaceGroup.Igmp.ImmediateLeave != nil {
						nestedMulticastInterfaceGroup.Igmp.ImmediateLeave = util.YesNo(oMulticastInterfaceGroup.Igmp.ImmediateLeave, nil)
					}
					if oMulticastInterfaceGroup.Igmp.Robustness != nil {
						nestedMulticastInterfaceGroup.Igmp.Robustness = oMulticastInterfaceGroup.Igmp.Robustness
					}
					if oMulticastInterfaceGroup.Igmp.MaxSources != nil {
						nestedMulticastInterfaceGroup.Igmp.MaxSources = oMulticastInterfaceGroup.Igmp.MaxSources
					}
					if oMulticastInterfaceGroup.Igmp.Enable != nil {
						nestedMulticastInterfaceGroup.Igmp.Enable = util.YesNo(oMulticastInterfaceGroup.Igmp.Enable, nil)
					}
					if oMulticastInterfaceGroup.Igmp.QueryInterval != nil {
						nestedMulticastInterfaceGroup.Igmp.QueryInterval = oMulticastInterfaceGroup.Igmp.QueryInterval
					}
					if oMulticastInterfaceGroup.Igmp.LastMemberQueryInterval != nil {
						nestedMulticastInterfaceGroup.Igmp.LastMemberQueryInterval = oMulticastInterfaceGroup.Igmp.LastMemberQueryInterval
					}
				}
				if oMulticastInterfaceGroup.Pim != nil {
					nestedMulticastInterfaceGroup.Pim = &MulticastInterfaceGroupPimXml{}
					if _, ok := o.Misc["MulticastInterfaceGroupPim"]; ok {
						nestedMulticastInterfaceGroup.Pim.Misc = o.Misc["MulticastInterfaceGroupPim"]
					}
					if oMulticastInterfaceGroup.Pim.Enable != nil {
						nestedMulticastInterfaceGroup.Pim.Enable = util.YesNo(oMulticastInterfaceGroup.Pim.Enable, nil)
					}
					if oMulticastInterfaceGroup.Pim.AssertInterval != nil {
						nestedMulticastInterfaceGroup.Pim.AssertInterval = oMulticastInterfaceGroup.Pim.AssertInterval
					}
					if oMulticastInterfaceGroup.Pim.HelloInterval != nil {
						nestedMulticastInterfaceGroup.Pim.HelloInterval = oMulticastInterfaceGroup.Pim.HelloInterval
					}
					if oMulticastInterfaceGroup.Pim.JoinPruneInterval != nil {
						nestedMulticastInterfaceGroup.Pim.JoinPruneInterval = oMulticastInterfaceGroup.Pim.JoinPruneInterval
					}
					if oMulticastInterfaceGroup.Pim.DrPriority != nil {
						nestedMulticastInterfaceGroup.Pim.DrPriority = oMulticastInterfaceGroup.Pim.DrPriority
					}
					if oMulticastInterfaceGroup.Pim.BsrBorder != nil {
						nestedMulticastInterfaceGroup.Pim.BsrBorder = util.YesNo(oMulticastInterfaceGroup.Pim.BsrBorder, nil)
					}
					if oMulticastInterfaceGroup.Pim.AllowedNeighbors != nil {
						nestedMulticastInterfaceGroup.Pim.AllowedNeighbors = []MulticastInterfaceGroupPimAllowedNeighborsXml{}
						for _, oMulticastInterfaceGroupPimAllowedNeighbors := range oMulticastInterfaceGroup.Pim.AllowedNeighbors {
							nestedMulticastInterfaceGroupPimAllowedNeighbors := MulticastInterfaceGroupPimAllowedNeighborsXml{}
							if _, ok := o.Misc["MulticastInterfaceGroupPimAllowedNeighbors"]; ok {
								nestedMulticastInterfaceGroupPimAllowedNeighbors.Misc = o.Misc["MulticastInterfaceGroupPimAllowedNeighbors"]
							}
							if oMulticastInterfaceGroupPimAllowedNeighbors.Name != "" {
								nestedMulticastInterfaceGroupPimAllowedNeighbors.Name = oMulticastInterfaceGroupPimAllowedNeighbors.Name
							}
							nestedMulticastInterfaceGroup.Pim.AllowedNeighbors = append(nestedMulticastInterfaceGroup.Pim.AllowedNeighbors, nestedMulticastInterfaceGroupPimAllowedNeighbors)
						}
					}
				}
				if oMulticastInterfaceGroup.Name != "" {
					nestedMulticastInterfaceGroup.Name = oMulticastInterfaceGroup.Name
				}
				if oMulticastInterfaceGroup.Description != nil {
					nestedMulticastInterfaceGroup.Description = oMulticastInterfaceGroup.Description
				}
				if oMulticastInterfaceGroup.Interface != nil {
					nestedMulticastInterfaceGroup.Interface = util.StrToMem(oMulticastInterfaceGroup.Interface)
				}
				nestedMulticast.InterfaceGroup = append(nestedMulticast.InterfaceGroup, nestedMulticastInterfaceGroup)
			}
		}
		if o.Multicast.RouteAgeoutTime != nil {
			nestedMulticast.RouteAgeoutTime = o.Multicast.RouteAgeoutTime
		}
		if o.Multicast.Rp != nil {
			nestedMulticast.Rp = &MulticastRpXml{}
			if _, ok := o.Misc["MulticastRp"]; ok {
				nestedMulticast.Rp.Misc = o.Misc["MulticastRp"]
			}
			if o.Multicast.Rp.ExternalRp != nil {
				nestedMulticast.Rp.ExternalRp = []MulticastRpExternalRpXml{}
				for _, oMulticastRpExternalRp := range o.Multicast.Rp.ExternalRp {
					nestedMulticastRpExternalRp := MulticastRpExternalRpXml{}
					if _, ok := o.Misc["MulticastRpExternalRp"]; ok {
						nestedMulticastRpExternalRp.Misc = o.Misc["MulticastRpExternalRp"]
					}
					if oMulticastRpExternalRp.GroupAddresses != nil {
						nestedMulticastRpExternalRp.GroupAddresses = util.StrToMem(oMulticastRpExternalRp.GroupAddresses)
					}
					if oMulticastRpExternalRp.Override != nil {
						nestedMulticastRpExternalRp.Override = util.YesNo(oMulticastRpExternalRp.Override, nil)
					}
					if oMulticastRpExternalRp.Name != "" {
						nestedMulticastRpExternalRp.Name = oMulticastRpExternalRp.Name
					}
					nestedMulticast.Rp.ExternalRp = append(nestedMulticast.Rp.ExternalRp, nestedMulticastRpExternalRp)
				}
			}
			if o.Multicast.Rp.LocalRp != nil {
				nestedMulticast.Rp.LocalRp = &MulticastRpLocalRpXml{}
				if _, ok := o.Misc["MulticastRpLocalRp"]; ok {
					nestedMulticast.Rp.LocalRp.Misc = o.Misc["MulticastRpLocalRp"]
				}
				if o.Multicast.Rp.LocalRp.CandidateRp != nil {
					nestedMulticast.Rp.LocalRp.CandidateRp = &MulticastRpLocalRpCandidateRpXml{}
					if _, ok := o.Misc["MulticastRpLocalRpCandidateRp"]; ok {
						nestedMulticast.Rp.LocalRp.CandidateRp.Misc = o.Misc["MulticastRpLocalRpCandidateRp"]
					}
					if o.Multicast.Rp.LocalRp.CandidateRp.Address != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp.Address = o.Multicast.Rp.LocalRp.CandidateRp.Address
					}
					if o.Multicast.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp.AdvertisementInterval = o.Multicast.Rp.LocalRp.CandidateRp.AdvertisementInterval
					}
					if o.Multicast.Rp.LocalRp.CandidateRp.GroupAddresses != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp.GroupAddresses = util.StrToMem(o.Multicast.Rp.LocalRp.CandidateRp.GroupAddresses)
					}
					if o.Multicast.Rp.LocalRp.CandidateRp.Interface != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp.Interface = o.Multicast.Rp.LocalRp.CandidateRp.Interface
					}
					if o.Multicast.Rp.LocalRp.CandidateRp.Priority != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp.Priority = o.Multicast.Rp.LocalRp.CandidateRp.Priority
					}
				}
				if o.Multicast.Rp.LocalRp.StaticRp != nil {
					nestedMulticast.Rp.LocalRp.StaticRp = &MulticastRpLocalRpStaticRpXml{}
					if _, ok := o.Misc["MulticastRpLocalRpStaticRp"]; ok {
						nestedMulticast.Rp.LocalRp.StaticRp.Misc = o.Misc["MulticastRpLocalRpStaticRp"]
					}
					if o.Multicast.Rp.LocalRp.StaticRp.Override != nil {
						nestedMulticast.Rp.LocalRp.StaticRp.Override = util.YesNo(o.Multicast.Rp.LocalRp.StaticRp.Override, nil)
					}
					if o.Multicast.Rp.LocalRp.StaticRp.Address != nil {
						nestedMulticast.Rp.LocalRp.StaticRp.Address = o.Multicast.Rp.LocalRp.StaticRp.Address
					}
					if o.Multicast.Rp.LocalRp.StaticRp.GroupAddresses != nil {
						nestedMulticast.Rp.LocalRp.StaticRp.GroupAddresses = util.StrToMem(o.Multicast.Rp.LocalRp.StaticRp.GroupAddresses)
					}
					if o.Multicast.Rp.LocalRp.StaticRp.Interface != nil {
						nestedMulticast.Rp.LocalRp.StaticRp.Interface = o.Multicast.Rp.LocalRp.StaticRp.Interface
					}
				}
			}
		}
		if o.Multicast.SptThreshold != nil {
			nestedMulticast.SptThreshold = []MulticastSptThresholdXml{}
			for _, oMulticastSptThreshold := range o.Multicast.SptThreshold {
				nestedMulticastSptThreshold := MulticastSptThresholdXml{}
				if _, ok := o.Misc["MulticastSptThreshold"]; ok {
					nestedMulticastSptThreshold.Misc = o.Misc["MulticastSptThreshold"]
				}
				if oMulticastSptThreshold.Threshold != nil {
					nestedMulticastSptThreshold.Threshold = oMulticastSptThreshold.Threshold
				}
				if oMulticastSptThreshold.Name != "" {
					nestedMulticastSptThreshold.Name = oMulticastSptThreshold.Name
				}
				nestedMulticast.SptThreshold = append(nestedMulticast.SptThreshold, nestedMulticastSptThreshold)
			}
		}
		if o.Multicast.SsmAddressSpace != nil {
			nestedMulticast.SsmAddressSpace = []MulticastSsmAddressSpaceXml{}
			for _, oMulticastSsmAddressSpace := range o.Multicast.SsmAddressSpace {
				nestedMulticastSsmAddressSpace := MulticastSsmAddressSpaceXml{}
				if _, ok := o.Misc["MulticastSsmAddressSpace"]; ok {
					nestedMulticastSsmAddressSpace.Misc = o.Misc["MulticastSsmAddressSpace"]
				}
				if oMulticastSsmAddressSpace.Included != nil {
					nestedMulticastSsmAddressSpace.Included = util.YesNo(oMulticastSsmAddressSpace.Included, nil)
				}
				if oMulticastSsmAddressSpace.Name != "" {
					nestedMulticastSsmAddressSpace.Name = oMulticastSsmAddressSpace.Name
				}
				if oMulticastSsmAddressSpace.GroupAddress != nil {
					nestedMulticastSsmAddressSpace.GroupAddress = oMulticastSsmAddressSpace.GroupAddress
				}
				nestedMulticast.SsmAddressSpace = append(nestedMulticast.SsmAddressSpace, nestedMulticastSsmAddressSpace)
			}
		}
		if o.Multicast.Enable != nil {
			nestedMulticast.Enable = util.YesNo(o.Multicast.Enable, nil)
		}
	}
	entry.Multicast = nestedMulticast

	var nestedProtocol *ProtocolXml
	if o.Protocol != nil {
		nestedProtocol = &ProtocolXml{}
		if _, ok := o.Misc["Protocol"]; ok {
			nestedProtocol.Misc = o.Misc["Protocol"]
		}
		if o.Protocol.RedistProfileIpv6 != nil {
			nestedProtocol.RedistProfileIpv6 = []ProtocolRedistProfileIpv6Xml{}
			for _, oProtocolRedistProfileIpv6 := range o.Protocol.RedistProfileIpv6 {
				nestedProtocolRedistProfileIpv6 := ProtocolRedistProfileIpv6Xml{}
				if _, ok := o.Misc["ProtocolRedistProfileIpv6"]; ok {
					nestedProtocolRedistProfileIpv6.Misc = o.Misc["ProtocolRedistProfileIpv6"]
				}
				if oProtocolRedistProfileIpv6.Priority != nil {
					nestedProtocolRedistProfileIpv6.Priority = oProtocolRedistProfileIpv6.Priority
				}
				if oProtocolRedistProfileIpv6.Filter != nil {
					nestedProtocolRedistProfileIpv6.Filter = &ProtocolRedistProfileIpv6FilterXml{}
					if _, ok := o.Misc["ProtocolRedistProfileIpv6Filter"]; ok {
						nestedProtocolRedistProfileIpv6.Filter.Misc = o.Misc["ProtocolRedistProfileIpv6Filter"]
					}
					if oProtocolRedistProfileIpv6.Filter.Bgp != nil {
						nestedProtocolRedistProfileIpv6.Filter.Bgp = &ProtocolRedistProfileIpv6FilterBgpXml{}
						if _, ok := o.Misc["ProtocolRedistProfileIpv6FilterBgp"]; ok {
							nestedProtocolRedistProfileIpv6.Filter.Bgp.Misc = o.Misc["ProtocolRedistProfileIpv6FilterBgp"]
						}
						if oProtocolRedistProfileIpv6.Filter.Bgp.Community != nil {
							nestedProtocolRedistProfileIpv6.Filter.Bgp.Community = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Bgp.Community)
						}
						if oProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity != nil {
							nestedProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity)
						}
					}
					if oProtocolRedistProfileIpv6.Filter.Type != nil {
						nestedProtocolRedistProfileIpv6.Filter.Type = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Type)
					}
					if oProtocolRedistProfileIpv6.Filter.Interface != nil {
						nestedProtocolRedistProfileIpv6.Filter.Interface = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Interface)
					}
					if oProtocolRedistProfileIpv6.Filter.Destination != nil {
						nestedProtocolRedistProfileIpv6.Filter.Destination = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Destination)
					}
					if oProtocolRedistProfileIpv6.Filter.Nexthop != nil {
						nestedProtocolRedistProfileIpv6.Filter.Nexthop = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Nexthop)
					}
					if oProtocolRedistProfileIpv6.Filter.Ospfv3 != nil {
						nestedProtocolRedistProfileIpv6.Filter.Ospfv3 = &ProtocolRedistProfileIpv6FilterOspfv3Xml{}
						if _, ok := o.Misc["ProtocolRedistProfileIpv6FilterOspfv3"]; ok {
							nestedProtocolRedistProfileIpv6.Filter.Ospfv3.Misc = o.Misc["ProtocolRedistProfileIpv6FilterOspfv3"]
						}
						if oProtocolRedistProfileIpv6.Filter.Ospfv3.Area != nil {
							nestedProtocolRedistProfileIpv6.Filter.Ospfv3.Area = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Ospfv3.Area)
						}
						if oProtocolRedistProfileIpv6.Filter.Ospfv3.Tag != nil {
							nestedProtocolRedistProfileIpv6.Filter.Ospfv3.Tag = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Ospfv3.Tag)
						}
						if oProtocolRedistProfileIpv6.Filter.Ospfv3.PathType != nil {
							nestedProtocolRedistProfileIpv6.Filter.Ospfv3.PathType = util.StrToMem(oProtocolRedistProfileIpv6.Filter.Ospfv3.PathType)
						}
					}
				}
				if oProtocolRedistProfileIpv6.Action != nil {
					nestedProtocolRedistProfileIpv6.Action = &ProtocolRedistProfileIpv6ActionXml{}
					if _, ok := o.Misc["ProtocolRedistProfileIpv6Action"]; ok {
						nestedProtocolRedistProfileIpv6.Action.Misc = o.Misc["ProtocolRedistProfileIpv6Action"]
					}
					if oProtocolRedistProfileIpv6.Action.NoRedist != nil {
						nestedProtocolRedistProfileIpv6.Action.NoRedist = &ProtocolRedistProfileIpv6ActionNoRedistXml{}
						if _, ok := o.Misc["ProtocolRedistProfileIpv6ActionNoRedist"]; ok {
							nestedProtocolRedistProfileIpv6.Action.NoRedist.Misc = o.Misc["ProtocolRedistProfileIpv6ActionNoRedist"]
						}
					}
					if oProtocolRedistProfileIpv6.Action.Redist != nil {
						nestedProtocolRedistProfileIpv6.Action.Redist = &ProtocolRedistProfileIpv6ActionRedistXml{}
						if _, ok := o.Misc["ProtocolRedistProfileIpv6ActionRedist"]; ok {
							nestedProtocolRedistProfileIpv6.Action.Redist.Misc = o.Misc["ProtocolRedistProfileIpv6ActionRedist"]
						}
					}
				}
				if oProtocolRedistProfileIpv6.Name != "" {
					nestedProtocolRedistProfileIpv6.Name = oProtocolRedistProfileIpv6.Name
				}
				nestedProtocol.RedistProfileIpv6 = append(nestedProtocol.RedistProfileIpv6, nestedProtocolRedistProfileIpv6)
			}
		}
		if o.Protocol.Rip != nil {
			nestedProtocol.Rip = &ProtocolRipXml{}
			if _, ok := o.Misc["ProtocolRip"]; ok {
				nestedProtocol.Rip.Misc = o.Misc["ProtocolRip"]
			}
			if o.Protocol.Rip.GlobalBfd != nil {
				nestedProtocol.Rip.GlobalBfd = &ProtocolRipGlobalBfdXml{}
				if _, ok := o.Misc["ProtocolRipGlobalBfd"]; ok {
					nestedProtocol.Rip.GlobalBfd.Misc = o.Misc["ProtocolRipGlobalBfd"]
				}
				if o.Protocol.Rip.GlobalBfd.Profile != nil {
					nestedProtocol.Rip.GlobalBfd.Profile = o.Protocol.Rip.GlobalBfd.Profile
				}
			}
			if o.Protocol.Rip.Interface != nil {
				nestedProtocol.Rip.Interface = []ProtocolRipInterfaceXml{}
				for _, oProtocolRipInterface := range o.Protocol.Rip.Interface {
					nestedProtocolRipInterface := ProtocolRipInterfaceXml{}
					if _, ok := o.Misc["ProtocolRipInterface"]; ok {
						nestedProtocolRipInterface.Misc = o.Misc["ProtocolRipInterface"]
					}
					if oProtocolRipInterface.Enable != nil {
						nestedProtocolRipInterface.Enable = util.YesNo(oProtocolRipInterface.Enable, nil)
					}
					if oProtocolRipInterface.Authentication != nil {
						nestedProtocolRipInterface.Authentication = oProtocolRipInterface.Authentication
					}
					if oProtocolRipInterface.Mode != nil {
						nestedProtocolRipInterface.Mode = oProtocolRipInterface.Mode
					}
					if oProtocolRipInterface.DefaultRoute != nil {
						nestedProtocolRipInterface.DefaultRoute = &ProtocolRipInterfaceDefaultRouteXml{}
						if _, ok := o.Misc["ProtocolRipInterfaceDefaultRoute"]; ok {
							nestedProtocolRipInterface.DefaultRoute.Misc = o.Misc["ProtocolRipInterfaceDefaultRoute"]
						}
						if oProtocolRipInterface.DefaultRoute.Disable != nil {
							nestedProtocolRipInterface.DefaultRoute.Disable = &ProtocolRipInterfaceDefaultRouteDisableXml{}
							if _, ok := o.Misc["ProtocolRipInterfaceDefaultRouteDisable"]; ok {
								nestedProtocolRipInterface.DefaultRoute.Disable.Misc = o.Misc["ProtocolRipInterfaceDefaultRouteDisable"]
							}
						}
						if oProtocolRipInterface.DefaultRoute.Advertise != nil {
							nestedProtocolRipInterface.DefaultRoute.Advertise = &ProtocolRipInterfaceDefaultRouteAdvertiseXml{}
							if _, ok := o.Misc["ProtocolRipInterfaceDefaultRouteAdvertise"]; ok {
								nestedProtocolRipInterface.DefaultRoute.Advertise.Misc = o.Misc["ProtocolRipInterfaceDefaultRouteAdvertise"]
							}
							if oProtocolRipInterface.DefaultRoute.Advertise.Metric != nil {
								nestedProtocolRipInterface.DefaultRoute.Advertise.Metric = oProtocolRipInterface.DefaultRoute.Advertise.Metric
							}
						}
					}
					if oProtocolRipInterface.Bfd != nil {
						nestedProtocolRipInterface.Bfd = &ProtocolRipInterfaceBfdXml{}
						if _, ok := o.Misc["ProtocolRipInterfaceBfd"]; ok {
							nestedProtocolRipInterface.Bfd.Misc = o.Misc["ProtocolRipInterfaceBfd"]
						}
						if oProtocolRipInterface.Bfd.Profile != nil {
							nestedProtocolRipInterface.Bfd.Profile = oProtocolRipInterface.Bfd.Profile
						}
					}
					if oProtocolRipInterface.Name != "" {
						nestedProtocolRipInterface.Name = oProtocolRipInterface.Name
					}
					nestedProtocol.Rip.Interface = append(nestedProtocol.Rip.Interface, nestedProtocolRipInterface)
				}
			}
			if o.Protocol.Rip.RejectDefaultRoute != nil {
				nestedProtocol.Rip.RejectDefaultRoute = util.YesNo(o.Protocol.Rip.RejectDefaultRoute, nil)
			}
			if o.Protocol.Rip.Timers != nil {
				nestedProtocol.Rip.Timers = &ProtocolRipTimersXml{}
				if _, ok := o.Misc["ProtocolRipTimers"]; ok {
					nestedProtocol.Rip.Timers.Misc = o.Misc["ProtocolRipTimers"]
				}
				if o.Protocol.Rip.Timers.DeleteIntervals != nil {
					nestedProtocol.Rip.Timers.DeleteIntervals = o.Protocol.Rip.Timers.DeleteIntervals
				}
				if o.Protocol.Rip.Timers.ExpireIntervals != nil {
					nestedProtocol.Rip.Timers.ExpireIntervals = o.Protocol.Rip.Timers.ExpireIntervals
				}
				if o.Protocol.Rip.Timers.IntervalSeconds != nil {
					nestedProtocol.Rip.Timers.IntervalSeconds = o.Protocol.Rip.Timers.IntervalSeconds
				}
				if o.Protocol.Rip.Timers.UpdateIntervals != nil {
					nestedProtocol.Rip.Timers.UpdateIntervals = o.Protocol.Rip.Timers.UpdateIntervals
				}
			}
			if o.Protocol.Rip.AllowRedistDefaultRoute != nil {
				nestedProtocol.Rip.AllowRedistDefaultRoute = util.YesNo(o.Protocol.Rip.AllowRedistDefaultRoute, nil)
			}
			if o.Protocol.Rip.AuthProfile != nil {
				nestedProtocol.Rip.AuthProfile = []ProtocolRipAuthProfileXml{}
				for _, oProtocolRipAuthProfile := range o.Protocol.Rip.AuthProfile {
					nestedProtocolRipAuthProfile := ProtocolRipAuthProfileXml{}
					if _, ok := o.Misc["ProtocolRipAuthProfile"]; ok {
						nestedProtocolRipAuthProfile.Misc = o.Misc["ProtocolRipAuthProfile"]
					}
					if oProtocolRipAuthProfile.Name != "" {
						nestedProtocolRipAuthProfile.Name = oProtocolRipAuthProfile.Name
					}
					if oProtocolRipAuthProfile.Password != nil {
						nestedProtocolRipAuthProfile.Password = oProtocolRipAuthProfile.Password
					}
					if oProtocolRipAuthProfile.Md5 != nil {
						nestedProtocolRipAuthProfile.Md5 = []ProtocolRipAuthProfileMd5Xml{}
						for _, oProtocolRipAuthProfileMd5 := range oProtocolRipAuthProfile.Md5 {
							nestedProtocolRipAuthProfileMd5 := ProtocolRipAuthProfileMd5Xml{}
							if _, ok := o.Misc["ProtocolRipAuthProfileMd5"]; ok {
								nestedProtocolRipAuthProfileMd5.Misc = o.Misc["ProtocolRipAuthProfileMd5"]
							}
							if oProtocolRipAuthProfileMd5.Key != nil {
								nestedProtocolRipAuthProfileMd5.Key = oProtocolRipAuthProfileMd5.Key
							}
							if oProtocolRipAuthProfileMd5.Preferred != nil {
								nestedProtocolRipAuthProfileMd5.Preferred = util.YesNo(oProtocolRipAuthProfileMd5.Preferred, nil)
							}
							if oProtocolRipAuthProfileMd5.Name != "" {
								nestedProtocolRipAuthProfileMd5.Name = oProtocolRipAuthProfileMd5.Name
							}
							nestedProtocolRipAuthProfile.Md5 = append(nestedProtocolRipAuthProfile.Md5, nestedProtocolRipAuthProfileMd5)
						}
					}
					nestedProtocol.Rip.AuthProfile = append(nestedProtocol.Rip.AuthProfile, nestedProtocolRipAuthProfile)
				}
			}
			if o.Protocol.Rip.Enable != nil {
				nestedProtocol.Rip.Enable = util.YesNo(o.Protocol.Rip.Enable, nil)
			}
			if o.Protocol.Rip.ExportRules != nil {
				nestedProtocol.Rip.ExportRules = []ProtocolRipExportRulesXml{}
				for _, oProtocolRipExportRules := range o.Protocol.Rip.ExportRules {
					nestedProtocolRipExportRules := ProtocolRipExportRulesXml{}
					if _, ok := o.Misc["ProtocolRipExportRules"]; ok {
						nestedProtocolRipExportRules.Misc = o.Misc["ProtocolRipExportRules"]
					}
					if oProtocolRipExportRules.Metric != nil {
						nestedProtocolRipExportRules.Metric = oProtocolRipExportRules.Metric
					}
					if oProtocolRipExportRules.Name != "" {
						nestedProtocolRipExportRules.Name = oProtocolRipExportRules.Name
					}
					nestedProtocol.Rip.ExportRules = append(nestedProtocol.Rip.ExportRules, nestedProtocolRipExportRules)
				}
			}
		}
		if o.Protocol.Bgp != nil {
			nestedProtocol.Bgp = &ProtocolBgpXml{}
			if _, ok := o.Misc["ProtocolBgp"]; ok {
				nestedProtocol.Bgp.Misc = o.Misc["ProtocolBgp"]
			}
			if o.Protocol.Bgp.RedistRules != nil {
				nestedProtocol.Bgp.RedistRules = []ProtocolBgpRedistRulesXml{}
				for _, oProtocolBgpRedistRules := range o.Protocol.Bgp.RedistRules {
					nestedProtocolBgpRedistRules := ProtocolBgpRedistRulesXml{}
					if _, ok := o.Misc["ProtocolBgpRedistRules"]; ok {
						nestedProtocolBgpRedistRules.Misc = o.Misc["ProtocolBgpRedistRules"]
					}
					if oProtocolBgpRedistRules.Name != "" {
						nestedProtocolBgpRedistRules.Name = oProtocolBgpRedistRules.Name
					}
					if oProtocolBgpRedistRules.AddressFamilyIdentifier != nil {
						nestedProtocolBgpRedistRules.AddressFamilyIdentifier = oProtocolBgpRedistRules.AddressFamilyIdentifier
					}
					if oProtocolBgpRedistRules.Enable != nil {
						nestedProtocolBgpRedistRules.Enable = util.YesNo(oProtocolBgpRedistRules.Enable, nil)
					}
					if oProtocolBgpRedistRules.SetLocalPreference != nil {
						nestedProtocolBgpRedistRules.SetLocalPreference = oProtocolBgpRedistRules.SetLocalPreference
					}
					if oProtocolBgpRedistRules.SetExtendedCommunity != nil {
						nestedProtocolBgpRedistRules.SetExtendedCommunity = util.StrToMem(oProtocolBgpRedistRules.SetExtendedCommunity)
					}
					if oProtocolBgpRedistRules.Metric != nil {
						nestedProtocolBgpRedistRules.Metric = oProtocolBgpRedistRules.Metric
					}
					if oProtocolBgpRedistRules.SetCommunity != nil {
						nestedProtocolBgpRedistRules.SetCommunity = util.StrToMem(oProtocolBgpRedistRules.SetCommunity)
					}
					if oProtocolBgpRedistRules.RouteTable != nil {
						nestedProtocolBgpRedistRules.RouteTable = oProtocolBgpRedistRules.RouteTable
					}
					if oProtocolBgpRedistRules.SetOrigin != nil {
						nestedProtocolBgpRedistRules.SetOrigin = oProtocolBgpRedistRules.SetOrigin
					}
					if oProtocolBgpRedistRules.SetMed != nil {
						nestedProtocolBgpRedistRules.SetMed = oProtocolBgpRedistRules.SetMed
					}
					if oProtocolBgpRedistRules.SetAsPathLimit != nil {
						nestedProtocolBgpRedistRules.SetAsPathLimit = oProtocolBgpRedistRules.SetAsPathLimit
					}
					nestedProtocol.Bgp.RedistRules = append(nestedProtocol.Bgp.RedistRules, nestedProtocolBgpRedistRules)
				}
			}
			if o.Protocol.Bgp.RouterId != nil {
				nestedProtocol.Bgp.RouterId = o.Protocol.Bgp.RouterId
			}
			if o.Protocol.Bgp.AllowRedistDefaultRoute != nil {
				nestedProtocol.Bgp.AllowRedistDefaultRoute = util.YesNo(o.Protocol.Bgp.AllowRedistDefaultRoute, nil)
			}
			if o.Protocol.Bgp.EnforceFirstAs != nil {
				nestedProtocol.Bgp.EnforceFirstAs = util.YesNo(o.Protocol.Bgp.EnforceFirstAs, nil)
			}
			if o.Protocol.Bgp.GlobalBfd != nil {
				nestedProtocol.Bgp.GlobalBfd = &ProtocolBgpGlobalBfdXml{}
				if _, ok := o.Misc["ProtocolBgpGlobalBfd"]; ok {
					nestedProtocol.Bgp.GlobalBfd.Misc = o.Misc["ProtocolBgpGlobalBfd"]
				}
				if o.Protocol.Bgp.GlobalBfd.Profile != nil {
					nestedProtocol.Bgp.GlobalBfd.Profile = o.Protocol.Bgp.GlobalBfd.Profile
				}
			}
			if o.Protocol.Bgp.LocalAs != nil {
				nestedProtocol.Bgp.LocalAs = o.Protocol.Bgp.LocalAs
			}
			if o.Protocol.Bgp.RoutingOptions != nil {
				nestedProtocol.Bgp.RoutingOptions = &ProtocolBgpRoutingOptionsXml{}
				if _, ok := o.Misc["ProtocolBgpRoutingOptions"]; ok {
					nestedProtocol.Bgp.RoutingOptions.Misc = o.Misc["ProtocolBgpRoutingOptions"]
				}
				if o.Protocol.Bgp.RoutingOptions.ReflectorClusterId != nil {
					nestedProtocol.Bgp.RoutingOptions.ReflectorClusterId = o.Protocol.Bgp.RoutingOptions.ReflectorClusterId
				}
				if o.Protocol.Bgp.RoutingOptions.Aggregate != nil {
					nestedProtocol.Bgp.RoutingOptions.Aggregate = &ProtocolBgpRoutingOptionsAggregateXml{}
					if _, ok := o.Misc["ProtocolBgpRoutingOptionsAggregate"]; ok {
						nestedProtocol.Bgp.RoutingOptions.Aggregate.Misc = o.Misc["ProtocolBgpRoutingOptionsAggregate"]
					}
					if o.Protocol.Bgp.RoutingOptions.Aggregate.AggregateMed != nil {
						nestedProtocol.Bgp.RoutingOptions.Aggregate.AggregateMed = util.YesNo(o.Protocol.Bgp.RoutingOptions.Aggregate.AggregateMed, nil)
					}
				}
				if o.Protocol.Bgp.RoutingOptions.AsFormat != nil {
					nestedProtocol.Bgp.RoutingOptions.AsFormat = o.Protocol.Bgp.RoutingOptions.AsFormat
				}
				if o.Protocol.Bgp.RoutingOptions.ConfederationMemberAs != nil {
					nestedProtocol.Bgp.RoutingOptions.ConfederationMemberAs = o.Protocol.Bgp.RoutingOptions.ConfederationMemberAs
				}
				if o.Protocol.Bgp.RoutingOptions.DefaultLocalPreference != nil {
					nestedProtocol.Bgp.RoutingOptions.DefaultLocalPreference = o.Protocol.Bgp.RoutingOptions.DefaultLocalPreference
				}
				if o.Protocol.Bgp.RoutingOptions.GracefulRestart != nil {
					nestedProtocol.Bgp.RoutingOptions.GracefulRestart = &ProtocolBgpRoutingOptionsGracefulRestartXml{}
					if _, ok := o.Misc["ProtocolBgpRoutingOptionsGracefulRestart"]; ok {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart.Misc = o.Misc["ProtocolBgpRoutingOptionsGracefulRestart"]
					}
					if o.Protocol.Bgp.RoutingOptions.GracefulRestart.Enable != nil {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart.Enable = util.YesNo(o.Protocol.Bgp.RoutingOptions.GracefulRestart.Enable, nil)
					}
					if o.Protocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime != nil {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime
					}
					if o.Protocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime != nil {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime
					}
					if o.Protocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime != nil {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime
					}
				}
				if o.Protocol.Bgp.RoutingOptions.Med != nil {
					nestedProtocol.Bgp.RoutingOptions.Med = &ProtocolBgpRoutingOptionsMedXml{}
					if _, ok := o.Misc["ProtocolBgpRoutingOptionsMed"]; ok {
						nestedProtocol.Bgp.RoutingOptions.Med.Misc = o.Misc["ProtocolBgpRoutingOptionsMed"]
					}
					if o.Protocol.Bgp.RoutingOptions.Med.AlwaysCompareMed != nil {
						nestedProtocol.Bgp.RoutingOptions.Med.AlwaysCompareMed = util.YesNo(o.Protocol.Bgp.RoutingOptions.Med.AlwaysCompareMed, nil)
					}
					if o.Protocol.Bgp.RoutingOptions.Med.DeterministicMedComparison != nil {
						nestedProtocol.Bgp.RoutingOptions.Med.DeterministicMedComparison = util.YesNo(o.Protocol.Bgp.RoutingOptions.Med.DeterministicMedComparison, nil)
					}
				}
			}
			if o.Protocol.Bgp.AuthProfile != nil {
				nestedProtocol.Bgp.AuthProfile = []ProtocolBgpAuthProfileXml{}
				for _, oProtocolBgpAuthProfile := range o.Protocol.Bgp.AuthProfile {
					nestedProtocolBgpAuthProfile := ProtocolBgpAuthProfileXml{}
					if _, ok := o.Misc["ProtocolBgpAuthProfile"]; ok {
						nestedProtocolBgpAuthProfile.Misc = o.Misc["ProtocolBgpAuthProfile"]
					}
					if oProtocolBgpAuthProfile.Secret != nil {
						nestedProtocolBgpAuthProfile.Secret = oProtocolBgpAuthProfile.Secret
					}
					if oProtocolBgpAuthProfile.Name != "" {
						nestedProtocolBgpAuthProfile.Name = oProtocolBgpAuthProfile.Name
					}
					nestedProtocol.Bgp.AuthProfile = append(nestedProtocol.Bgp.AuthProfile, nestedProtocolBgpAuthProfile)
				}
			}
			if o.Protocol.Bgp.Enable != nil {
				nestedProtocol.Bgp.Enable = util.YesNo(o.Protocol.Bgp.Enable, nil)
			}
			if o.Protocol.Bgp.InstallRoute != nil {
				nestedProtocol.Bgp.InstallRoute = util.YesNo(o.Protocol.Bgp.InstallRoute, nil)
			}
			if o.Protocol.Bgp.Policy != nil {
				nestedProtocol.Bgp.Policy = &ProtocolBgpPolicyXml{}
				if _, ok := o.Misc["ProtocolBgpPolicy"]; ok {
					nestedProtocol.Bgp.Policy.Misc = o.Misc["ProtocolBgpPolicy"]
				}
				if o.Protocol.Bgp.Policy.Aggregation != nil {
					nestedProtocol.Bgp.Policy.Aggregation = &ProtocolBgpPolicyAggregationXml{}
					if _, ok := o.Misc["ProtocolBgpPolicyAggregation"]; ok {
						nestedProtocol.Bgp.Policy.Aggregation.Misc = o.Misc["ProtocolBgpPolicyAggregation"]
					}
					if o.Protocol.Bgp.Policy.Aggregation.Address != nil {
						nestedProtocol.Bgp.Policy.Aggregation.Address = []ProtocolBgpPolicyAggregationAddressXml{}
						for _, oProtocolBgpPolicyAggregationAddress := range o.Protocol.Bgp.Policy.Aggregation.Address {
							nestedProtocolBgpPolicyAggregationAddress := ProtocolBgpPolicyAggregationAddressXml{}
							if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddress"]; ok {
								nestedProtocolBgpPolicyAggregationAddress.Misc = o.Misc["ProtocolBgpPolicyAggregationAddress"]
							}
							if oProtocolBgpPolicyAggregationAddress.Prefix != nil {
								nestedProtocolBgpPolicyAggregationAddress.Prefix = oProtocolBgpPolicyAggregationAddress.Prefix
							}
							if oProtocolBgpPolicyAggregationAddress.Enable != nil {
								nestedProtocolBgpPolicyAggregationAddress.Enable = util.YesNo(oProtocolBgpPolicyAggregationAddress.Enable, nil)
							}
							if oProtocolBgpPolicyAggregationAddress.Summary != nil {
								nestedProtocolBgpPolicyAggregationAddress.Summary = util.YesNo(oProtocolBgpPolicyAggregationAddress.Summary, nil)
							}
							if oProtocolBgpPolicyAggregationAddress.AsSet != nil {
								nestedProtocolBgpPolicyAggregationAddress.AsSet = util.YesNo(oProtocolBgpPolicyAggregationAddress.AsSet, nil)
							}
							if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes != nil {
								nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesXml{}
								if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes"]; ok {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes"]
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath"]; ok {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath"]
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNoneXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone"]
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemoveXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove"]
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend
									}
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity"]; ok {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity"]
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAllXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll"]
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append = util.StrToMem(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append)
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite = util.StrToMem(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite)
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNoneXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone"]
										}
									}
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity"]; ok {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity"]
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNoneXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone"]
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAllXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll"]; ok {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll"]
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append = util.StrToMem(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append)
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite = util.StrToMem(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite)
									}
								}
							}
							if oProtocolBgpPolicyAggregationAddress.SuppressFilters != nil {
								nestedProtocolBgpPolicyAggregationAddress.SuppressFilters = []ProtocolBgpPolicyAggregationAddressSuppressFiltersXml{}
								for _, oProtocolBgpPolicyAggregationAddressSuppressFilters := range oProtocolBgpPolicyAggregationAddress.SuppressFilters {
									nestedProtocolBgpPolicyAggregationAddressSuppressFilters := ProtocolBgpPolicyAggregationAddressSuppressFiltersXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFilters"]; ok {
										nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFilters"]
									}
									if oProtocolBgpPolicyAggregationAddressSuppressFilters.Enable != nil {
										nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Enable = util.YesNo(oProtocolBgpPolicyAggregationAddressSuppressFilters.Enable, nil)
									}
									if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match != nil {
										nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch"]; ok {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch"]
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix = []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml{}
											for _, oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix := range oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix {
												nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix := ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefixXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix"]; ok {
													nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix"]
												}
												if oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact != nil {
													nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact = util.YesNo(oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact, nil)
												}
												if oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name != "" {
													nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name
												}
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix, nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix)
											}
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop)
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer)
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath"]; ok {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath"]
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity"]; ok {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity"]
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity"]
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable
										}
									}
									if oProtocolBgpPolicyAggregationAddressSuppressFilters.Name != "" {
										nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Name = oProtocolBgpPolicyAggregationAddressSuppressFilters.Name
									}
									nestedProtocolBgpPolicyAggregationAddress.SuppressFilters = append(nestedProtocolBgpPolicyAggregationAddress.SuppressFilters, nestedProtocolBgpPolicyAggregationAddressSuppressFilters)
								}
							}
							if oProtocolBgpPolicyAggregationAddress.AdvertiseFilters != nil {
								nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters = []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersXml{}
								for _, oProtocolBgpPolicyAggregationAddressAdvertiseFilters := range oProtocolBgpPolicyAggregationAddress.AdvertiseFilters {
									nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters := ProtocolBgpPolicyAggregationAddressAdvertiseFiltersXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFilters"]; ok {
										nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFilters"]
									}
									if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name != "" {
										nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name
									}
									if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable != nil {
										nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable = util.YesNo(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable, nil)
									}
									if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match != nil {
										nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch"]; ok {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch"]
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix = []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml{}
											for _, oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix := range oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix := ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefixXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix"]; ok {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix"]
												}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name != "" {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name
												}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact != nil {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact = util.YesNo(oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact, nil)
												}
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix, nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix)
											}
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop)
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer)
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath"]; ok {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath"]
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity"]; ok {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity"]
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity"]
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex
											}
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable
										}
									}
									nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters = append(nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters, nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters)
								}
							}
							if oProtocolBgpPolicyAggregationAddress.Name != "" {
								nestedProtocolBgpPolicyAggregationAddress.Name = oProtocolBgpPolicyAggregationAddress.Name
							}
							nestedProtocol.Bgp.Policy.Aggregation.Address = append(nestedProtocol.Bgp.Policy.Aggregation.Address, nestedProtocolBgpPolicyAggregationAddress)
						}
					}
				}
				if o.Protocol.Bgp.Policy.ConditionalAdvertisement != nil {
					nestedProtocol.Bgp.Policy.ConditionalAdvertisement = &ProtocolBgpPolicyConditionalAdvertisementXml{}
					if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisement"]; ok {
						nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisement"]
					}
					if o.Protocol.Bgp.Policy.ConditionalAdvertisement.Policy != nil {
						nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy = []ProtocolBgpPolicyConditionalAdvertisementPolicyXml{}
						for _, oProtocolBgpPolicyConditionalAdvertisementPolicy := range o.Protocol.Bgp.Policy.ConditionalAdvertisement.Policy {
							nestedProtocolBgpPolicyConditionalAdvertisementPolicy := ProtocolBgpPolicyConditionalAdvertisementPolicyXml{}
							if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicy"]; ok {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicy"]
							}
							if oProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters != nil {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters = []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml{}
								for _, oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters := range oProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters := ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters"]; ok {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters"]
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable != nil {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable = util.YesNo(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable, nil)
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match != nil {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch"]; ok {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch"]
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix = []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml{}
											for _, oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix := range oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix := ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefixXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix"]; ok {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix"]
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name != "" {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name
												}
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix, nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix)
											}
										}
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name != "" {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name
									}
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters, nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters)
								}
							}
							if oProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters != nil {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters = []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml{}
								for _, oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters := range oProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters := ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters"]; ok {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters"]
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable != nil {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable = util.YesNo(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable, nil)
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match != nil {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch"]; ok {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch"]
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix = []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml{}
											for _, oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix := range oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix := ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefixXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix"]; ok {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix"]
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name != "" {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name
												}
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix, nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix)
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity"]
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med
										}
									}
									if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name != "" {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name
									}
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters, nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters)
								}
							}
							if oProtocolBgpPolicyConditionalAdvertisementPolicy.Name != "" {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.Name = oProtocolBgpPolicyConditionalAdvertisementPolicy.Name
							}
							if oProtocolBgpPolicyConditionalAdvertisementPolicy.Enable != nil {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.Enable = util.YesNo(oProtocolBgpPolicyConditionalAdvertisementPolicy.Enable, nil)
							}
							if oProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy != nil {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy = util.StrToMem(oProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy)
							}
							nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy = append(nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy, nestedProtocolBgpPolicyConditionalAdvertisementPolicy)
						}
					}
				}
				if o.Protocol.Bgp.Policy.Export != nil {
					nestedProtocol.Bgp.Policy.Export = &ProtocolBgpPolicyExportXml{}
					if _, ok := o.Misc["ProtocolBgpPolicyExport"]; ok {
						nestedProtocol.Bgp.Policy.Export.Misc = o.Misc["ProtocolBgpPolicyExport"]
					}
					if o.Protocol.Bgp.Policy.Export.Rules != nil {
						nestedProtocol.Bgp.Policy.Export.Rules = []ProtocolBgpPolicyExportRulesXml{}
						for _, oProtocolBgpPolicyExportRules := range o.Protocol.Bgp.Policy.Export.Rules {
							nestedProtocolBgpPolicyExportRules := ProtocolBgpPolicyExportRulesXml{}
							if _, ok := o.Misc["ProtocolBgpPolicyExportRules"]; ok {
								nestedProtocolBgpPolicyExportRules.Misc = o.Misc["ProtocolBgpPolicyExportRules"]
							}
							if oProtocolBgpPolicyExportRules.Name != "" {
								nestedProtocolBgpPolicyExportRules.Name = oProtocolBgpPolicyExportRules.Name
							}
							if oProtocolBgpPolicyExportRules.Enable != nil {
								nestedProtocolBgpPolicyExportRules.Enable = util.YesNo(oProtocolBgpPolicyExportRules.Enable, nil)
							}
							if oProtocolBgpPolicyExportRules.UsedBy != nil {
								nestedProtocolBgpPolicyExportRules.UsedBy = util.StrToMem(oProtocolBgpPolicyExportRules.UsedBy)
							}
							if oProtocolBgpPolicyExportRules.Match != nil {
								nestedProtocolBgpPolicyExportRules.Match = &ProtocolBgpPolicyExportRulesMatchXml{}
								if _, ok := o.Misc["ProtocolBgpPolicyExportRulesMatch"]; ok {
									nestedProtocolBgpPolicyExportRules.Match.Misc = o.Misc["ProtocolBgpPolicyExportRulesMatch"]
								}
								if oProtocolBgpPolicyExportRules.Match.FromPeer != nil {
									nestedProtocolBgpPolicyExportRules.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyExportRules.Match.FromPeer)
								}
								if oProtocolBgpPolicyExportRules.Match.AsPath != nil {
									nestedProtocolBgpPolicyExportRules.Match.AsPath = &ProtocolBgpPolicyExportRulesMatchAsPathXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyExportRulesMatchAsPath"]; ok {
										nestedProtocolBgpPolicyExportRules.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyExportRulesMatchAsPath"]
									}
									if oProtocolBgpPolicyExportRules.Match.AsPath.Regex != nil {
										nestedProtocolBgpPolicyExportRules.Match.AsPath.Regex = oProtocolBgpPolicyExportRules.Match.AsPath.Regex
									}
								}
								if oProtocolBgpPolicyExportRules.Match.Community != nil {
									nestedProtocolBgpPolicyExportRules.Match.Community = &ProtocolBgpPolicyExportRulesMatchCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyExportRulesMatchCommunity"]; ok {
										nestedProtocolBgpPolicyExportRules.Match.Community.Misc = o.Misc["ProtocolBgpPolicyExportRulesMatchCommunity"]
									}
									if oProtocolBgpPolicyExportRules.Match.Community.Regex != nil {
										nestedProtocolBgpPolicyExportRules.Match.Community.Regex = oProtocolBgpPolicyExportRules.Match.Community.Regex
									}
								}
								if oProtocolBgpPolicyExportRules.Match.ExtendedCommunity != nil {
									nestedProtocolBgpPolicyExportRules.Match.ExtendedCommunity = &ProtocolBgpPolicyExportRulesMatchExtendedCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyExportRulesMatchExtendedCommunity"]; ok {
										nestedProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyExportRulesMatchExtendedCommunity"]
									}
									if oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex != nil {
										nestedProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex
									}
								}
								if oProtocolBgpPolicyExportRules.Match.RouteTable != nil {
									nestedProtocolBgpPolicyExportRules.Match.RouteTable = oProtocolBgpPolicyExportRules.Match.RouteTable
								}
								if oProtocolBgpPolicyExportRules.Match.Med != nil {
									nestedProtocolBgpPolicyExportRules.Match.Med = oProtocolBgpPolicyExportRules.Match.Med
								}
								if oProtocolBgpPolicyExportRules.Match.AddressPrefix != nil {
									nestedProtocolBgpPolicyExportRules.Match.AddressPrefix = []ProtocolBgpPolicyExportRulesMatchAddressPrefixXml{}
									for _, oProtocolBgpPolicyExportRulesMatchAddressPrefix := range oProtocolBgpPolicyExportRules.Match.AddressPrefix {
										nestedProtocolBgpPolicyExportRulesMatchAddressPrefix := ProtocolBgpPolicyExportRulesMatchAddressPrefixXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyExportRulesMatchAddressPrefix"]; ok {
											nestedProtocolBgpPolicyExportRulesMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyExportRulesMatchAddressPrefix"]
										}
										if oProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact != nil {
											nestedProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact = util.YesNo(oProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact, nil)
										}
										if oProtocolBgpPolicyExportRulesMatchAddressPrefix.Name != "" {
											nestedProtocolBgpPolicyExportRulesMatchAddressPrefix.Name = oProtocolBgpPolicyExportRulesMatchAddressPrefix.Name
										}
										nestedProtocolBgpPolicyExportRules.Match.AddressPrefix = append(nestedProtocolBgpPolicyExportRules.Match.AddressPrefix, nestedProtocolBgpPolicyExportRulesMatchAddressPrefix)
									}
								}
								if oProtocolBgpPolicyExportRules.Match.Nexthop != nil {
									nestedProtocolBgpPolicyExportRules.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyExportRules.Match.Nexthop)
								}
							}
							if oProtocolBgpPolicyExportRules.Action != nil {
								nestedProtocolBgpPolicyExportRules.Action = &ProtocolBgpPolicyExportRulesActionXml{}
								if _, ok := o.Misc["ProtocolBgpPolicyExportRulesAction"]; ok {
									nestedProtocolBgpPolicyExportRules.Action.Misc = o.Misc["ProtocolBgpPolicyExportRulesAction"]
								}
								if oProtocolBgpPolicyExportRules.Action.Allow != nil {
									nestedProtocolBgpPolicyExportRules.Action.Allow = &ProtocolBgpPolicyExportRulesActionAllowXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllow"]; ok {
										nestedProtocolBgpPolicyExportRules.Action.Allow.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllow"]
									}
									if oProtocolBgpPolicyExportRules.Action.Allow.Update != nil {
										nestedProtocolBgpPolicyExportRules.Action.Allow.Update = &ProtocolBgpPolicyExportRulesActionAllowUpdateXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdate"]; ok {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdate"]
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath"]; ok {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath"]
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemoveXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend
											}
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity"]; ok {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity"]
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite = util.StrToMem(oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite)
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAllXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex = oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append = util.StrToMem(oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append)
											}
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity"]
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append = util.StrToMem(oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append)
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite = util.StrToMem(oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite)
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAllXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll"]; ok {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll"]
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex
											}
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference = oProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.Med != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Med = oProtocolBgpPolicyExportRules.Action.Allow.Update.Med
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop = oProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.Origin != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Origin = oProtocolBgpPolicyExportRules.Action.Allow.Update.Origin
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit
										}
									}
								}
								if oProtocolBgpPolicyExportRules.Action.Deny != nil {
									nestedProtocolBgpPolicyExportRules.Action.Deny = &ProtocolBgpPolicyExportRulesActionDenyXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyExportRulesActionDeny"]; ok {
										nestedProtocolBgpPolicyExportRules.Action.Deny.Misc = o.Misc["ProtocolBgpPolicyExportRulesActionDeny"]
									}
								}
							}
							nestedProtocol.Bgp.Policy.Export.Rules = append(nestedProtocol.Bgp.Policy.Export.Rules, nestedProtocolBgpPolicyExportRules)
						}
					}
				}
				if o.Protocol.Bgp.Policy.Import != nil {
					nestedProtocol.Bgp.Policy.Import = &ProtocolBgpPolicyImportXml{}
					if _, ok := o.Misc["ProtocolBgpPolicyImport"]; ok {
						nestedProtocol.Bgp.Policy.Import.Misc = o.Misc["ProtocolBgpPolicyImport"]
					}
					if o.Protocol.Bgp.Policy.Import.Rules != nil {
						nestedProtocol.Bgp.Policy.Import.Rules = []ProtocolBgpPolicyImportRulesXml{}
						for _, oProtocolBgpPolicyImportRules := range o.Protocol.Bgp.Policy.Import.Rules {
							nestedProtocolBgpPolicyImportRules := ProtocolBgpPolicyImportRulesXml{}
							if _, ok := o.Misc["ProtocolBgpPolicyImportRules"]; ok {
								nestedProtocolBgpPolicyImportRules.Misc = o.Misc["ProtocolBgpPolicyImportRules"]
							}
							if oProtocolBgpPolicyImportRules.Enable != nil {
								nestedProtocolBgpPolicyImportRules.Enable = util.YesNo(oProtocolBgpPolicyImportRules.Enable, nil)
							}
							if oProtocolBgpPolicyImportRules.UsedBy != nil {
								nestedProtocolBgpPolicyImportRules.UsedBy = util.StrToMem(oProtocolBgpPolicyImportRules.UsedBy)
							}
							if oProtocolBgpPolicyImportRules.Match != nil {
								nestedProtocolBgpPolicyImportRules.Match = &ProtocolBgpPolicyImportRulesMatchXml{}
								if _, ok := o.Misc["ProtocolBgpPolicyImportRulesMatch"]; ok {
									nestedProtocolBgpPolicyImportRules.Match.Misc = o.Misc["ProtocolBgpPolicyImportRulesMatch"]
								}
								if oProtocolBgpPolicyImportRules.Match.RouteTable != nil {
									nestedProtocolBgpPolicyImportRules.Match.RouteTable = oProtocolBgpPolicyImportRules.Match.RouteTable
								}
								if oProtocolBgpPolicyImportRules.Match.Med != nil {
									nestedProtocolBgpPolicyImportRules.Match.Med = oProtocolBgpPolicyImportRules.Match.Med
								}
								if oProtocolBgpPolicyImportRules.Match.AddressPrefix != nil {
									nestedProtocolBgpPolicyImportRules.Match.AddressPrefix = []ProtocolBgpPolicyImportRulesMatchAddressPrefixXml{}
									for _, oProtocolBgpPolicyImportRulesMatchAddressPrefix := range oProtocolBgpPolicyImportRules.Match.AddressPrefix {
										nestedProtocolBgpPolicyImportRulesMatchAddressPrefix := ProtocolBgpPolicyImportRulesMatchAddressPrefixXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyImportRulesMatchAddressPrefix"]; ok {
											nestedProtocolBgpPolicyImportRulesMatchAddressPrefix.Misc = o.Misc["ProtocolBgpPolicyImportRulesMatchAddressPrefix"]
										}
										if oProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact != nil {
											nestedProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact = util.YesNo(oProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact, nil)
										}
										if oProtocolBgpPolicyImportRulesMatchAddressPrefix.Name != "" {
											nestedProtocolBgpPolicyImportRulesMatchAddressPrefix.Name = oProtocolBgpPolicyImportRulesMatchAddressPrefix.Name
										}
										nestedProtocolBgpPolicyImportRules.Match.AddressPrefix = append(nestedProtocolBgpPolicyImportRules.Match.AddressPrefix, nestedProtocolBgpPolicyImportRulesMatchAddressPrefix)
									}
								}
								if oProtocolBgpPolicyImportRules.Match.Nexthop != nil {
									nestedProtocolBgpPolicyImportRules.Match.Nexthop = util.StrToMem(oProtocolBgpPolicyImportRules.Match.Nexthop)
								}
								if oProtocolBgpPolicyImportRules.Match.FromPeer != nil {
									nestedProtocolBgpPolicyImportRules.Match.FromPeer = util.StrToMem(oProtocolBgpPolicyImportRules.Match.FromPeer)
								}
								if oProtocolBgpPolicyImportRules.Match.AsPath != nil {
									nestedProtocolBgpPolicyImportRules.Match.AsPath = &ProtocolBgpPolicyImportRulesMatchAsPathXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyImportRulesMatchAsPath"]; ok {
										nestedProtocolBgpPolicyImportRules.Match.AsPath.Misc = o.Misc["ProtocolBgpPolicyImportRulesMatchAsPath"]
									}
									if oProtocolBgpPolicyImportRules.Match.AsPath.Regex != nil {
										nestedProtocolBgpPolicyImportRules.Match.AsPath.Regex = oProtocolBgpPolicyImportRules.Match.AsPath.Regex
									}
								}
								if oProtocolBgpPolicyImportRules.Match.Community != nil {
									nestedProtocolBgpPolicyImportRules.Match.Community = &ProtocolBgpPolicyImportRulesMatchCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyImportRulesMatchCommunity"]; ok {
										nestedProtocolBgpPolicyImportRules.Match.Community.Misc = o.Misc["ProtocolBgpPolicyImportRulesMatchCommunity"]
									}
									if oProtocolBgpPolicyImportRules.Match.Community.Regex != nil {
										nestedProtocolBgpPolicyImportRules.Match.Community.Regex = oProtocolBgpPolicyImportRules.Match.Community.Regex
									}
								}
								if oProtocolBgpPolicyImportRules.Match.ExtendedCommunity != nil {
									nestedProtocolBgpPolicyImportRules.Match.ExtendedCommunity = &ProtocolBgpPolicyImportRulesMatchExtendedCommunityXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyImportRulesMatchExtendedCommunity"]; ok {
										nestedProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyImportRulesMatchExtendedCommunity"]
									}
									if oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex != nil {
										nestedProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex
									}
								}
							}
							if oProtocolBgpPolicyImportRules.Action != nil {
								nestedProtocolBgpPolicyImportRules.Action = &ProtocolBgpPolicyImportRulesActionXml{}
								if _, ok := o.Misc["ProtocolBgpPolicyImportRulesAction"]; ok {
									nestedProtocolBgpPolicyImportRules.Action.Misc = o.Misc["ProtocolBgpPolicyImportRulesAction"]
								}
								if oProtocolBgpPolicyImportRules.Action.Deny != nil {
									nestedProtocolBgpPolicyImportRules.Action.Deny = &ProtocolBgpPolicyImportRulesActionDenyXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionDeny"]; ok {
										nestedProtocolBgpPolicyImportRules.Action.Deny.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionDeny"]
									}
								}
								if oProtocolBgpPolicyImportRules.Action.Allow != nil {
									nestedProtocolBgpPolicyImportRules.Action.Allow = &ProtocolBgpPolicyImportRulesActionAllowXml{}
									if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllow"]; ok {
										nestedProtocolBgpPolicyImportRules.Action.Allow.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllow"]
									}
									if oProtocolBgpPolicyImportRules.Action.Allow.Dampening != nil {
										nestedProtocolBgpPolicyImportRules.Action.Allow.Dampening = oProtocolBgpPolicyImportRules.Action.Allow.Dampening
									}
									if oProtocolBgpPolicyImportRules.Action.Allow.Update != nil {
										nestedProtocolBgpPolicyImportRules.Action.Allow.Update = &ProtocolBgpPolicyImportRulesActionAllowUpdateXml{}
										if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdate"]; ok {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdate"]
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.Weight != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Weight = oProtocolBgpPolicyImportRules.Action.Allow.Update.Weight
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.Origin != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Origin = oProtocolBgpPolicyImportRules.Action.Allow.Update.Origin
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit = oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity"]; ok {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity"]
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone"]
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAllXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll"]
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append = util.StrToMem(oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append)
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite = util.StrToMem(oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite)
											}
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference = oProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.Med != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Med = oProtocolBgpPolicyImportRules.Action.Allow.Update.Med
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop = oProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath"]; ok {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath"]
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone"]
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemoveXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove"]
												}
											}
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityXml{}
											if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity"]; ok {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity"]
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite = util.StrToMem(oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite)
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNoneXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone"]
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAllXml{}
												if _, ok := o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll"]; ok {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll.Misc = o.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll"]
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex = oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append = util.StrToMem(oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append)
											}
										}
									}
								}
							}
							if oProtocolBgpPolicyImportRules.Name != "" {
								nestedProtocolBgpPolicyImportRules.Name = oProtocolBgpPolicyImportRules.Name
							}
							nestedProtocol.Bgp.Policy.Import.Rules = append(nestedProtocol.Bgp.Policy.Import.Rules, nestedProtocolBgpPolicyImportRules)
						}
					}
				}
			}
			if o.Protocol.Bgp.DampeningProfile != nil {
				nestedProtocol.Bgp.DampeningProfile = []ProtocolBgpDampeningProfileXml{}
				for _, oProtocolBgpDampeningProfile := range o.Protocol.Bgp.DampeningProfile {
					nestedProtocolBgpDampeningProfile := ProtocolBgpDampeningProfileXml{}
					if _, ok := o.Misc["ProtocolBgpDampeningProfile"]; ok {
						nestedProtocolBgpDampeningProfile.Misc = o.Misc["ProtocolBgpDampeningProfile"]
					}
					if oProtocolBgpDampeningProfile.MaxHoldTime != nil {
						nestedProtocolBgpDampeningProfile.MaxHoldTime = oProtocolBgpDampeningProfile.MaxHoldTime
					}
					if oProtocolBgpDampeningProfile.DecayHalfLifeReachable != nil {
						nestedProtocolBgpDampeningProfile.DecayHalfLifeReachable = oProtocolBgpDampeningProfile.DecayHalfLifeReachable
					}
					if oProtocolBgpDampeningProfile.DecayHalfLifeUnreachable != nil {
						nestedProtocolBgpDampeningProfile.DecayHalfLifeUnreachable = oProtocolBgpDampeningProfile.DecayHalfLifeUnreachable
					}
					if oProtocolBgpDampeningProfile.Name != "" {
						nestedProtocolBgpDampeningProfile.Name = oProtocolBgpDampeningProfile.Name
					}
					if oProtocolBgpDampeningProfile.Enable != nil {
						nestedProtocolBgpDampeningProfile.Enable = util.YesNo(oProtocolBgpDampeningProfile.Enable, nil)
					}
					if oProtocolBgpDampeningProfile.Cutoff != nil {
						nestedProtocolBgpDampeningProfile.Cutoff = oProtocolBgpDampeningProfile.Cutoff
					}
					if oProtocolBgpDampeningProfile.Reuse != nil {
						nestedProtocolBgpDampeningProfile.Reuse = oProtocolBgpDampeningProfile.Reuse
					}
					nestedProtocol.Bgp.DampeningProfile = append(nestedProtocol.Bgp.DampeningProfile, nestedProtocolBgpDampeningProfile)
				}
			}
			if o.Protocol.Bgp.EcmpMultiAs != nil {
				nestedProtocol.Bgp.EcmpMultiAs = util.YesNo(o.Protocol.Bgp.EcmpMultiAs, nil)
			}
			if o.Protocol.Bgp.PeerGroup != nil {
				nestedProtocol.Bgp.PeerGroup = []ProtocolBgpPeerGroupXml{}
				for _, oProtocolBgpPeerGroup := range o.Protocol.Bgp.PeerGroup {
					nestedProtocolBgpPeerGroup := ProtocolBgpPeerGroupXml{}
					if _, ok := o.Misc["ProtocolBgpPeerGroup"]; ok {
						nestedProtocolBgpPeerGroup.Misc = o.Misc["ProtocolBgpPeerGroup"]
					}
					if oProtocolBgpPeerGroup.Enable != nil {
						nestedProtocolBgpPeerGroup.Enable = util.YesNo(oProtocolBgpPeerGroup.Enable, nil)
					}
					if oProtocolBgpPeerGroup.AggregatedConfedAsPath != nil {
						nestedProtocolBgpPeerGroup.AggregatedConfedAsPath = util.YesNo(oProtocolBgpPeerGroup.AggregatedConfedAsPath, nil)
					}
					if oProtocolBgpPeerGroup.SoftResetWithStoredInfo != nil {
						nestedProtocolBgpPeerGroup.SoftResetWithStoredInfo = util.YesNo(oProtocolBgpPeerGroup.SoftResetWithStoredInfo, nil)
					}
					if oProtocolBgpPeerGroup.Type != nil {
						nestedProtocolBgpPeerGroup.Type = &ProtocolBgpPeerGroupTypeXml{}
						if _, ok := o.Misc["ProtocolBgpPeerGroupType"]; ok {
							nestedProtocolBgpPeerGroup.Type.Misc = o.Misc["ProtocolBgpPeerGroupType"]
						}
						if oProtocolBgpPeerGroup.Type.EbgpConfed != nil {
							nestedProtocolBgpPeerGroup.Type.EbgpConfed = &ProtocolBgpPeerGroupTypeEbgpConfedXml{}
							if _, ok := o.Misc["ProtocolBgpPeerGroupTypeEbgpConfed"]; ok {
								nestedProtocolBgpPeerGroup.Type.EbgpConfed.Misc = o.Misc["ProtocolBgpPeerGroupTypeEbgpConfed"]
							}
							if oProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop != nil {
								nestedProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop = oProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop
							}
						}
						if oProtocolBgpPeerGroup.Type.IbgpConfed != nil {
							nestedProtocolBgpPeerGroup.Type.IbgpConfed = &ProtocolBgpPeerGroupTypeIbgpConfedXml{}
							if _, ok := o.Misc["ProtocolBgpPeerGroupTypeIbgpConfed"]; ok {
								nestedProtocolBgpPeerGroup.Type.IbgpConfed.Misc = o.Misc["ProtocolBgpPeerGroupTypeIbgpConfed"]
							}
							if oProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop != nil {
								nestedProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop = oProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop
							}
						}
						if oProtocolBgpPeerGroup.Type.Ebgp != nil {
							nestedProtocolBgpPeerGroup.Type.Ebgp = &ProtocolBgpPeerGroupTypeEbgpXml{}
							if _, ok := o.Misc["ProtocolBgpPeerGroupTypeEbgp"]; ok {
								nestedProtocolBgpPeerGroup.Type.Ebgp.Misc = o.Misc["ProtocolBgpPeerGroupTypeEbgp"]
							}
							if oProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop != nil {
								nestedProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop = oProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop
							}
							if oProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop != nil {
								nestedProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop = oProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop
							}
							if oProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs != nil {
								nestedProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs = util.YesNo(oProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs, nil)
							}
						}
						if oProtocolBgpPeerGroup.Type.Ibgp != nil {
							nestedProtocolBgpPeerGroup.Type.Ibgp = &ProtocolBgpPeerGroupTypeIbgpXml{}
							if _, ok := o.Misc["ProtocolBgpPeerGroupTypeIbgp"]; ok {
								nestedProtocolBgpPeerGroup.Type.Ibgp.Misc = o.Misc["ProtocolBgpPeerGroupTypeIbgp"]
							}
							if oProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop != nil {
								nestedProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop = oProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop
							}
						}
					}
					if oProtocolBgpPeerGroup.Peer != nil {
						nestedProtocolBgpPeerGroup.Peer = []ProtocolBgpPeerGroupPeerXml{}
						for _, oProtocolBgpPeerGroupPeer := range oProtocolBgpPeerGroup.Peer {
							nestedProtocolBgpPeerGroupPeer := ProtocolBgpPeerGroupPeerXml{}
							if _, ok := o.Misc["ProtocolBgpPeerGroupPeer"]; ok {
								nestedProtocolBgpPeerGroupPeer.Misc = o.Misc["ProtocolBgpPeerGroupPeer"]
							}
							if oProtocolBgpPeerGroupPeer.MaxPrefixes != nil {
								nestedProtocolBgpPeerGroupPeer.MaxPrefixes = oProtocolBgpPeerGroupPeer.MaxPrefixes
							}
							if oProtocolBgpPeerGroupPeer.Name != "" {
								nestedProtocolBgpPeerGroupPeer.Name = oProtocolBgpPeerGroupPeer.Name
							}
							if oProtocolBgpPeerGroupPeer.AddressFamilyIdentifier != nil {
								nestedProtocolBgpPeerGroupPeer.AddressFamilyIdentifier = oProtocolBgpPeerGroupPeer.AddressFamilyIdentifier
							}
							if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier != nil {
								nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier = &ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifierXml{}
								if _, ok := o.Misc["ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier"]; ok {
									nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Misc = o.Misc["ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier"]
								}
								if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast != nil {
									nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast = util.YesNo(oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast, nil)
								}
								if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast != nil {
									nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast = util.YesNo(oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast, nil)
								}
							}
							if oProtocolBgpPeerGroupPeer.Bfd != nil {
								nestedProtocolBgpPeerGroupPeer.Bfd = &ProtocolBgpPeerGroupPeerBfdXml{}
								if _, ok := o.Misc["ProtocolBgpPeerGroupPeerBfd"]; ok {
									nestedProtocolBgpPeerGroupPeer.Bfd.Misc = o.Misc["ProtocolBgpPeerGroupPeerBfd"]
								}
								if oProtocolBgpPeerGroupPeer.Bfd.Profile != nil {
									nestedProtocolBgpPeerGroupPeer.Bfd.Profile = oProtocolBgpPeerGroupPeer.Bfd.Profile
								}
							}
							if oProtocolBgpPeerGroupPeer.EnableMpBgp != nil {
								nestedProtocolBgpPeerGroupPeer.EnableMpBgp = util.YesNo(oProtocolBgpPeerGroupPeer.EnableMpBgp, nil)
							}
							if oProtocolBgpPeerGroupPeer.PeerAs != nil {
								nestedProtocolBgpPeerGroupPeer.PeerAs = oProtocolBgpPeerGroupPeer.PeerAs
							}
							if oProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
								nestedProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.YesNo(oProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
							}
							if oProtocolBgpPeerGroupPeer.PeeringType != nil {
								nestedProtocolBgpPeerGroupPeer.PeeringType = oProtocolBgpPeerGroupPeer.PeeringType
							}
							if oProtocolBgpPeerGroupPeer.LocalAddress != nil {
								nestedProtocolBgpPeerGroupPeer.LocalAddress = &ProtocolBgpPeerGroupPeerLocalAddressXml{}
								if _, ok := o.Misc["ProtocolBgpPeerGroupPeerLocalAddress"]; ok {
									nestedProtocolBgpPeerGroupPeer.LocalAddress.Misc = o.Misc["ProtocolBgpPeerGroupPeerLocalAddress"]
								}
								if oProtocolBgpPeerGroupPeer.LocalAddress.Interface != nil {
									nestedProtocolBgpPeerGroupPeer.LocalAddress.Interface = oProtocolBgpPeerGroupPeer.LocalAddress.Interface
								}
								if oProtocolBgpPeerGroupPeer.LocalAddress.Ip != nil {
									nestedProtocolBgpPeerGroupPeer.LocalAddress.Ip = oProtocolBgpPeerGroupPeer.LocalAddress.Ip
								}
							}
							if oProtocolBgpPeerGroupPeer.ConnectionOptions != nil {
								nestedProtocolBgpPeerGroupPeer.ConnectionOptions = &ProtocolBgpPeerGroupPeerConnectionOptionsXml{}
								if _, ok := o.Misc["ProtocolBgpPeerGroupPeerConnectionOptions"]; ok {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.Misc = o.Misc["ProtocolBgpPeerGroupPeerConnectionOptions"]
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop = oProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication = oProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval = oProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection = &ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnectionXml{}
									if _, ok := o.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection"]; ok {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Misc = o.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection"]
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort = oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow = util.YesNo(oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow, nil)
									}
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection = &ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnectionXml{}
									if _, ok := o.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection"]; ok {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Misc = o.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection"]
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort = oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow = util.YesNo(oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow, nil)
									}
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval = oProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime
								}
							}
							if oProtocolBgpPeerGroupPeer.Enable != nil {
								nestedProtocolBgpPeerGroupPeer.Enable = util.YesNo(oProtocolBgpPeerGroupPeer.Enable, nil)
							}
							if oProtocolBgpPeerGroupPeer.PeerAddress != nil {
								nestedProtocolBgpPeerGroupPeer.PeerAddress = &ProtocolBgpPeerGroupPeerPeerAddressXml{}
								if _, ok := o.Misc["ProtocolBgpPeerGroupPeerPeerAddress"]; ok {
									nestedProtocolBgpPeerGroupPeer.PeerAddress.Misc = o.Misc["ProtocolBgpPeerGroupPeerPeerAddress"]
								}
								if oProtocolBgpPeerGroupPeer.PeerAddress.Ip != nil {
									nestedProtocolBgpPeerGroupPeer.PeerAddress.Ip = oProtocolBgpPeerGroupPeer.PeerAddress.Ip
								}
								if oProtocolBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
									nestedProtocolBgpPeerGroupPeer.PeerAddress.Fqdn = oProtocolBgpPeerGroupPeer.PeerAddress.Fqdn
								}
							}
							if oProtocolBgpPeerGroupPeer.ReflectorClient != nil {
								nestedProtocolBgpPeerGroupPeer.ReflectorClient = oProtocolBgpPeerGroupPeer.ReflectorClient
							}
							nestedProtocolBgpPeerGroup.Peer = append(nestedProtocolBgpPeerGroup.Peer, nestedProtocolBgpPeerGroupPeer)
						}
					}
					if oProtocolBgpPeerGroup.Name != "" {
						nestedProtocolBgpPeerGroup.Name = oProtocolBgpPeerGroup.Name
					}
					nestedProtocol.Bgp.PeerGroup = append(nestedProtocol.Bgp.PeerGroup, nestedProtocolBgpPeerGroup)
				}
			}
			if o.Protocol.Bgp.RejectDefaultRoute != nil {
				nestedProtocol.Bgp.RejectDefaultRoute = util.YesNo(o.Protocol.Bgp.RejectDefaultRoute, nil)
			}
		}
		if o.Protocol.Ospf != nil {
			nestedProtocol.Ospf = &ProtocolOspfXml{}
			if _, ok := o.Misc["ProtocolOspf"]; ok {
				nestedProtocol.Ospf.Misc = o.Misc["ProtocolOspf"]
			}
			if o.Protocol.Ospf.AllowRedistDefaultRoute != nil {
				nestedProtocol.Ospf.AllowRedistDefaultRoute = util.YesNo(o.Protocol.Ospf.AllowRedistDefaultRoute, nil)
			}
			if o.Protocol.Ospf.Enable != nil {
				nestedProtocol.Ospf.Enable = util.YesNo(o.Protocol.Ospf.Enable, nil)
			}
			if o.Protocol.Ospf.GracefulRestart != nil {
				nestedProtocol.Ospf.GracefulRestart = &ProtocolOspfGracefulRestartXml{}
				if _, ok := o.Misc["ProtocolOspfGracefulRestart"]; ok {
					nestedProtocol.Ospf.GracefulRestart.Misc = o.Misc["ProtocolOspfGracefulRestart"]
				}
				if o.Protocol.Ospf.GracefulRestart.Enable != nil {
					nestedProtocol.Ospf.GracefulRestart.Enable = util.YesNo(o.Protocol.Ospf.GracefulRestart.Enable, nil)
				}
				if o.Protocol.Ospf.GracefulRestart.GracePeriod != nil {
					nestedProtocol.Ospf.GracefulRestart.GracePeriod = o.Protocol.Ospf.GracefulRestart.GracePeriod
				}
				if o.Protocol.Ospf.GracefulRestart.HelperEnable != nil {
					nestedProtocol.Ospf.GracefulRestart.HelperEnable = util.YesNo(o.Protocol.Ospf.GracefulRestart.HelperEnable, nil)
				}
				if o.Protocol.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
					nestedProtocol.Ospf.GracefulRestart.MaxNeighborRestartTime = o.Protocol.Ospf.GracefulRestart.MaxNeighborRestartTime
				}
				if o.Protocol.Ospf.GracefulRestart.StrictLSAChecking != nil {
					nestedProtocol.Ospf.GracefulRestart.StrictLSAChecking = util.YesNo(o.Protocol.Ospf.GracefulRestart.StrictLSAChecking, nil)
				}
			}
			if o.Protocol.Ospf.RejectDefaultRoute != nil {
				nestedProtocol.Ospf.RejectDefaultRoute = util.YesNo(o.Protocol.Ospf.RejectDefaultRoute, nil)
			}
			if o.Protocol.Ospf.Rfc1583 != nil {
				nestedProtocol.Ospf.Rfc1583 = util.YesNo(o.Protocol.Ospf.Rfc1583, nil)
			}
			if o.Protocol.Ospf.Timers != nil {
				nestedProtocol.Ospf.Timers = &ProtocolOspfTimersXml{}
				if _, ok := o.Misc["ProtocolOspfTimers"]; ok {
					nestedProtocol.Ospf.Timers.Misc = o.Misc["ProtocolOspfTimers"]
				}
				if o.Protocol.Ospf.Timers.LsaInterval != nil {
					nestedProtocol.Ospf.Timers.LsaInterval = o.Protocol.Ospf.Timers.LsaInterval
				}
				if o.Protocol.Ospf.Timers.SpfCalculationDelay != nil {
					nestedProtocol.Ospf.Timers.SpfCalculationDelay = o.Protocol.Ospf.Timers.SpfCalculationDelay
				}
			}
			if o.Protocol.Ospf.Area != nil {
				nestedProtocol.Ospf.Area = []ProtocolOspfAreaXml{}
				for _, oProtocolOspfArea := range o.Protocol.Ospf.Area {
					nestedProtocolOspfArea := ProtocolOspfAreaXml{}
					if _, ok := o.Misc["ProtocolOspfArea"]; ok {
						nestedProtocolOspfArea.Misc = o.Misc["ProtocolOspfArea"]
					}
					if oProtocolOspfArea.VirtualLink != nil {
						nestedProtocolOspfArea.VirtualLink = []ProtocolOspfAreaVirtualLinkXml{}
						for _, oProtocolOspfAreaVirtualLink := range oProtocolOspfArea.VirtualLink {
							nestedProtocolOspfAreaVirtualLink := ProtocolOspfAreaVirtualLinkXml{}
							if _, ok := o.Misc["ProtocolOspfAreaVirtualLink"]; ok {
								nestedProtocolOspfAreaVirtualLink.Misc = o.Misc["ProtocolOspfAreaVirtualLink"]
							}
							if oProtocolOspfAreaVirtualLink.RetransmitInterval != nil {
								nestedProtocolOspfAreaVirtualLink.RetransmitInterval = oProtocolOspfAreaVirtualLink.RetransmitInterval
							}
							if oProtocolOspfAreaVirtualLink.TransitDelay != nil {
								nestedProtocolOspfAreaVirtualLink.TransitDelay = oProtocolOspfAreaVirtualLink.TransitDelay
							}
							if oProtocolOspfAreaVirtualLink.Authentication != nil {
								nestedProtocolOspfAreaVirtualLink.Authentication = oProtocolOspfAreaVirtualLink.Authentication
							}
							if oProtocolOspfAreaVirtualLink.Name != "" {
								nestedProtocolOspfAreaVirtualLink.Name = oProtocolOspfAreaVirtualLink.Name
							}
							if oProtocolOspfAreaVirtualLink.NeighborId != nil {
								nestedProtocolOspfAreaVirtualLink.NeighborId = oProtocolOspfAreaVirtualLink.NeighborId
							}
							if oProtocolOspfAreaVirtualLink.TransitAreaId != nil {
								nestedProtocolOspfAreaVirtualLink.TransitAreaId = oProtocolOspfAreaVirtualLink.TransitAreaId
							}
							if oProtocolOspfAreaVirtualLink.DeadCounts != nil {
								nestedProtocolOspfAreaVirtualLink.DeadCounts = oProtocolOspfAreaVirtualLink.DeadCounts
							}
							if oProtocolOspfAreaVirtualLink.Enable != nil {
								nestedProtocolOspfAreaVirtualLink.Enable = util.YesNo(oProtocolOspfAreaVirtualLink.Enable, nil)
							}
							if oProtocolOspfAreaVirtualLink.HelloInterval != nil {
								nestedProtocolOspfAreaVirtualLink.HelloInterval = oProtocolOspfAreaVirtualLink.HelloInterval
							}
							if oProtocolOspfAreaVirtualLink.Bfd != nil {
								nestedProtocolOspfAreaVirtualLink.Bfd = &ProtocolOspfAreaVirtualLinkBfdXml{}
								if _, ok := o.Misc["ProtocolOspfAreaVirtualLinkBfd"]; ok {
									nestedProtocolOspfAreaVirtualLink.Bfd.Misc = o.Misc["ProtocolOspfAreaVirtualLinkBfd"]
								}
								if oProtocolOspfAreaVirtualLink.Bfd.Profile != nil {
									nestedProtocolOspfAreaVirtualLink.Bfd.Profile = oProtocolOspfAreaVirtualLink.Bfd.Profile
								}
							}
							nestedProtocolOspfArea.VirtualLink = append(nestedProtocolOspfArea.VirtualLink, nestedProtocolOspfAreaVirtualLink)
						}
					}
					if oProtocolOspfArea.Name != "" {
						nestedProtocolOspfArea.Name = oProtocolOspfArea.Name
					}
					if oProtocolOspfArea.Type != nil {
						nestedProtocolOspfArea.Type = &ProtocolOspfAreaTypeXml{}
						if _, ok := o.Misc["ProtocolOspfAreaType"]; ok {
							nestedProtocolOspfArea.Type.Misc = o.Misc["ProtocolOspfAreaType"]
						}
						if oProtocolOspfArea.Type.Normal != nil {
							nestedProtocolOspfArea.Type.Normal = &ProtocolOspfAreaTypeNormalXml{}
							if _, ok := o.Misc["ProtocolOspfAreaTypeNormal"]; ok {
								nestedProtocolOspfArea.Type.Normal.Misc = o.Misc["ProtocolOspfAreaTypeNormal"]
							}
						}
						if oProtocolOspfArea.Type.Stub != nil {
							nestedProtocolOspfArea.Type.Stub = &ProtocolOspfAreaTypeStubXml{}
							if _, ok := o.Misc["ProtocolOspfAreaTypeStub"]; ok {
								nestedProtocolOspfArea.Type.Stub.Misc = o.Misc["ProtocolOspfAreaTypeStub"]
							}
							if oProtocolOspfArea.Type.Stub.DefaultRoute != nil {
								nestedProtocolOspfArea.Type.Stub.DefaultRoute = &ProtocolOspfAreaTypeStubDefaultRouteXml{}
								if _, ok := o.Misc["ProtocolOspfAreaTypeStubDefaultRoute"]; ok {
									nestedProtocolOspfArea.Type.Stub.DefaultRoute.Misc = o.Misc["ProtocolOspfAreaTypeStubDefaultRoute"]
								}
								if oProtocolOspfArea.Type.Stub.DefaultRoute.Disable != nil {
									nestedProtocolOspfArea.Type.Stub.DefaultRoute.Disable = &ProtocolOspfAreaTypeStubDefaultRouteDisableXml{}
									if _, ok := o.Misc["ProtocolOspfAreaTypeStubDefaultRouteDisable"]; ok {
										nestedProtocolOspfArea.Type.Stub.DefaultRoute.Disable.Misc = o.Misc["ProtocolOspfAreaTypeStubDefaultRouteDisable"]
									}
								}
								if oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise != nil {
									nestedProtocolOspfArea.Type.Stub.DefaultRoute.Advertise = &ProtocolOspfAreaTypeStubDefaultRouteAdvertiseXml{}
									if _, ok := o.Misc["ProtocolOspfAreaTypeStubDefaultRouteAdvertise"]; ok {
										nestedProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Misc = o.Misc["ProtocolOspfAreaTypeStubDefaultRouteAdvertise"]
									}
									if oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric != nil {
										nestedProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric = oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric
									}
								}
							}
							if oProtocolOspfArea.Type.Stub.AcceptSummary != nil {
								nestedProtocolOspfArea.Type.Stub.AcceptSummary = util.YesNo(oProtocolOspfArea.Type.Stub.AcceptSummary, nil)
							}
						}
						if oProtocolOspfArea.Type.Nssa != nil {
							nestedProtocolOspfArea.Type.Nssa = &ProtocolOspfAreaTypeNssaXml{}
							if _, ok := o.Misc["ProtocolOspfAreaTypeNssa"]; ok {
								nestedProtocolOspfArea.Type.Nssa.Misc = o.Misc["ProtocolOspfAreaTypeNssa"]
							}
							if oProtocolOspfArea.Type.Nssa.AcceptSummary != nil {
								nestedProtocolOspfArea.Type.Nssa.AcceptSummary = util.YesNo(oProtocolOspfArea.Type.Nssa.AcceptSummary, nil)
							}
							if oProtocolOspfArea.Type.Nssa.DefaultRoute != nil {
								nestedProtocolOspfArea.Type.Nssa.DefaultRoute = &ProtocolOspfAreaTypeNssaDefaultRouteXml{}
								if _, ok := o.Misc["ProtocolOspfAreaTypeNssaDefaultRoute"]; ok {
									nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Misc = o.Misc["ProtocolOspfAreaTypeNssaDefaultRoute"]
								}
								if oProtocolOspfArea.Type.Nssa.DefaultRoute.Disable != nil {
									nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Disable = &ProtocolOspfAreaTypeNssaDefaultRouteDisableXml{}
									if _, ok := o.Misc["ProtocolOspfAreaTypeNssaDefaultRouteDisable"]; ok {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Disable.Misc = o.Misc["ProtocolOspfAreaTypeNssaDefaultRouteDisable"]
									}
								}
								if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise != nil {
									nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise = &ProtocolOspfAreaTypeNssaDefaultRouteAdvertiseXml{}
									if _, ok := o.Misc["ProtocolOspfAreaTypeNssaDefaultRouteAdvertise"]; ok {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Misc = o.Misc["ProtocolOspfAreaTypeNssaDefaultRouteAdvertise"]
									}
									if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric != nil {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric = oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric
									}
									if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type != nil {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type = oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type
									}
								}
							}
							if oProtocolOspfArea.Type.Nssa.NssaExtRange != nil {
								nestedProtocolOspfArea.Type.Nssa.NssaExtRange = []ProtocolOspfAreaTypeNssaNssaExtRangeXml{}
								for _, oProtocolOspfAreaTypeNssaNssaExtRange := range oProtocolOspfArea.Type.Nssa.NssaExtRange {
									nestedProtocolOspfAreaTypeNssaNssaExtRange := ProtocolOspfAreaTypeNssaNssaExtRangeXml{}
									if _, ok := o.Misc["ProtocolOspfAreaTypeNssaNssaExtRange"]; ok {
										nestedProtocolOspfAreaTypeNssaNssaExtRange.Misc = o.Misc["ProtocolOspfAreaTypeNssaNssaExtRange"]
									}
									if oProtocolOspfAreaTypeNssaNssaExtRange.Name != "" {
										nestedProtocolOspfAreaTypeNssaNssaExtRange.Name = oProtocolOspfAreaTypeNssaNssaExtRange.Name
									}
									if oProtocolOspfAreaTypeNssaNssaExtRange.Advertise != nil {
										nestedProtocolOspfAreaTypeNssaNssaExtRange.Advertise = &ProtocolOspfAreaTypeNssaNssaExtRangeAdvertiseXml{}
										if _, ok := o.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise"]; ok {
											nestedProtocolOspfAreaTypeNssaNssaExtRange.Advertise.Misc = o.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise"]
										}
									}
									if oProtocolOspfAreaTypeNssaNssaExtRange.Suppress != nil {
										nestedProtocolOspfAreaTypeNssaNssaExtRange.Suppress = &ProtocolOspfAreaTypeNssaNssaExtRangeSuppressXml{}
										if _, ok := o.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeSuppress"]; ok {
											nestedProtocolOspfAreaTypeNssaNssaExtRange.Suppress.Misc = o.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeSuppress"]
										}
									}
									nestedProtocolOspfArea.Type.Nssa.NssaExtRange = append(nestedProtocolOspfArea.Type.Nssa.NssaExtRange, nestedProtocolOspfAreaTypeNssaNssaExtRange)
								}
							}
						}
					}
					if oProtocolOspfArea.Range != nil {
						nestedProtocolOspfArea.Range = []ProtocolOspfAreaRangeXml{}
						for _, oProtocolOspfAreaRange := range oProtocolOspfArea.Range {
							nestedProtocolOspfAreaRange := ProtocolOspfAreaRangeXml{}
							if _, ok := o.Misc["ProtocolOspfAreaRange"]; ok {
								nestedProtocolOspfAreaRange.Misc = o.Misc["ProtocolOspfAreaRange"]
							}
							if oProtocolOspfAreaRange.Name != "" {
								nestedProtocolOspfAreaRange.Name = oProtocolOspfAreaRange.Name
							}
							if oProtocolOspfAreaRange.Advertise != nil {
								nestedProtocolOspfAreaRange.Advertise = &ProtocolOspfAreaRangeAdvertiseXml{}
								if _, ok := o.Misc["ProtocolOspfAreaRangeAdvertise"]; ok {
									nestedProtocolOspfAreaRange.Advertise.Misc = o.Misc["ProtocolOspfAreaRangeAdvertise"]
								}
							}
							if oProtocolOspfAreaRange.Suppress != nil {
								nestedProtocolOspfAreaRange.Suppress = &ProtocolOspfAreaRangeSuppressXml{}
								if _, ok := o.Misc["ProtocolOspfAreaRangeSuppress"]; ok {
									nestedProtocolOspfAreaRange.Suppress.Misc = o.Misc["ProtocolOspfAreaRangeSuppress"]
								}
							}
							nestedProtocolOspfArea.Range = append(nestedProtocolOspfArea.Range, nestedProtocolOspfAreaRange)
						}
					}
					if oProtocolOspfArea.Interface != nil {
						nestedProtocolOspfArea.Interface = []ProtocolOspfAreaInterfaceXml{}
						for _, oProtocolOspfAreaInterface := range oProtocolOspfArea.Interface {
							nestedProtocolOspfAreaInterface := ProtocolOspfAreaInterfaceXml{}
							if _, ok := o.Misc["ProtocolOspfAreaInterface"]; ok {
								nestedProtocolOspfAreaInterface.Misc = o.Misc["ProtocolOspfAreaInterface"]
							}
							if oProtocolOspfAreaInterface.Metric != nil {
								nestedProtocolOspfAreaInterface.Metric = oProtocolOspfAreaInterface.Metric
							}
							if oProtocolOspfAreaInterface.Neighbor != nil {
								nestedProtocolOspfAreaInterface.Neighbor = []ProtocolOspfAreaInterfaceNeighborXml{}
								for _, oProtocolOspfAreaInterfaceNeighbor := range oProtocolOspfAreaInterface.Neighbor {
									nestedProtocolOspfAreaInterfaceNeighbor := ProtocolOspfAreaInterfaceNeighborXml{}
									if _, ok := o.Misc["ProtocolOspfAreaInterfaceNeighbor"]; ok {
										nestedProtocolOspfAreaInterfaceNeighbor.Misc = o.Misc["ProtocolOspfAreaInterfaceNeighbor"]
									}
									if oProtocolOspfAreaInterfaceNeighbor.Name != "" {
										nestedProtocolOspfAreaInterfaceNeighbor.Name = oProtocolOspfAreaInterfaceNeighbor.Name
									}
									nestedProtocolOspfAreaInterface.Neighbor = append(nestedProtocolOspfAreaInterface.Neighbor, nestedProtocolOspfAreaInterfaceNeighbor)
								}
							}
							if oProtocolOspfAreaInterface.LinkType != nil {
								nestedProtocolOspfAreaInterface.LinkType = &ProtocolOspfAreaInterfaceLinkTypeXml{}
								if _, ok := o.Misc["ProtocolOspfAreaInterfaceLinkType"]; ok {
									nestedProtocolOspfAreaInterface.LinkType.Misc = o.Misc["ProtocolOspfAreaInterfaceLinkType"]
								}
								if oProtocolOspfAreaInterface.LinkType.Broadcast != nil {
									nestedProtocolOspfAreaInterface.LinkType.Broadcast = &ProtocolOspfAreaInterfaceLinkTypeBroadcastXml{}
									if _, ok := o.Misc["ProtocolOspfAreaInterfaceLinkTypeBroadcast"]; ok {
										nestedProtocolOspfAreaInterface.LinkType.Broadcast.Misc = o.Misc["ProtocolOspfAreaInterfaceLinkTypeBroadcast"]
									}
								}
								if oProtocolOspfAreaInterface.LinkType.P2p != nil {
									nestedProtocolOspfAreaInterface.LinkType.P2p = &ProtocolOspfAreaInterfaceLinkTypeP2pXml{}
									if _, ok := o.Misc["ProtocolOspfAreaInterfaceLinkTypeP2p"]; ok {
										nestedProtocolOspfAreaInterface.LinkType.P2p.Misc = o.Misc["ProtocolOspfAreaInterfaceLinkTypeP2p"]
									}
								}
								if oProtocolOspfAreaInterface.LinkType.P2mp != nil {
									nestedProtocolOspfAreaInterface.LinkType.P2mp = &ProtocolOspfAreaInterfaceLinkTypeP2mpXml{}
									if _, ok := o.Misc["ProtocolOspfAreaInterfaceLinkTypeP2mp"]; ok {
										nestedProtocolOspfAreaInterface.LinkType.P2mp.Misc = o.Misc["ProtocolOspfAreaInterfaceLinkTypeP2mp"]
									}
								}
							}
							if oProtocolOspfAreaInterface.Enable != nil {
								nestedProtocolOspfAreaInterface.Enable = util.YesNo(oProtocolOspfAreaInterface.Enable, nil)
							}
							if oProtocolOspfAreaInterface.Passive != nil {
								nestedProtocolOspfAreaInterface.Passive = util.YesNo(oProtocolOspfAreaInterface.Passive, nil)
							}
							if oProtocolOspfAreaInterface.HelloInterval != nil {
								nestedProtocolOspfAreaInterface.HelloInterval = oProtocolOspfAreaInterface.HelloInterval
							}
							if oProtocolOspfAreaInterface.DeadCounts != nil {
								nestedProtocolOspfAreaInterface.DeadCounts = oProtocolOspfAreaInterface.DeadCounts
							}
							if oProtocolOspfAreaInterface.TransitDelay != nil {
								nestedProtocolOspfAreaInterface.TransitDelay = oProtocolOspfAreaInterface.TransitDelay
							}
							if oProtocolOspfAreaInterface.GrDelay != nil {
								nestedProtocolOspfAreaInterface.GrDelay = oProtocolOspfAreaInterface.GrDelay
							}
							if oProtocolOspfAreaInterface.Bfd != nil {
								nestedProtocolOspfAreaInterface.Bfd = &ProtocolOspfAreaInterfaceBfdXml{}
								if _, ok := o.Misc["ProtocolOspfAreaInterfaceBfd"]; ok {
									nestedProtocolOspfAreaInterface.Bfd.Misc = o.Misc["ProtocolOspfAreaInterfaceBfd"]
								}
								if oProtocolOspfAreaInterface.Bfd.Profile != nil {
									nestedProtocolOspfAreaInterface.Bfd.Profile = oProtocolOspfAreaInterface.Bfd.Profile
								}
							}
							if oProtocolOspfAreaInterface.Name != "" {
								nestedProtocolOspfAreaInterface.Name = oProtocolOspfAreaInterface.Name
							}
							if oProtocolOspfAreaInterface.Priority != nil {
								nestedProtocolOspfAreaInterface.Priority = oProtocolOspfAreaInterface.Priority
							}
							if oProtocolOspfAreaInterface.RetransmitInterval != nil {
								nestedProtocolOspfAreaInterface.RetransmitInterval = oProtocolOspfAreaInterface.RetransmitInterval
							}
							if oProtocolOspfAreaInterface.Authentication != nil {
								nestedProtocolOspfAreaInterface.Authentication = oProtocolOspfAreaInterface.Authentication
							}
							nestedProtocolOspfArea.Interface = append(nestedProtocolOspfArea.Interface, nestedProtocolOspfAreaInterface)
						}
					}
					nestedProtocol.Ospf.Area = append(nestedProtocol.Ospf.Area, nestedProtocolOspfArea)
				}
			}
			if o.Protocol.Ospf.AuthProfile != nil {
				nestedProtocol.Ospf.AuthProfile = []ProtocolOspfAuthProfileXml{}
				for _, oProtocolOspfAuthProfile := range o.Protocol.Ospf.AuthProfile {
					nestedProtocolOspfAuthProfile := ProtocolOspfAuthProfileXml{}
					if _, ok := o.Misc["ProtocolOspfAuthProfile"]; ok {
						nestedProtocolOspfAuthProfile.Misc = o.Misc["ProtocolOspfAuthProfile"]
					}
					if oProtocolOspfAuthProfile.Name != "" {
						nestedProtocolOspfAuthProfile.Name = oProtocolOspfAuthProfile.Name
					}
					if oProtocolOspfAuthProfile.Password != nil {
						nestedProtocolOspfAuthProfile.Password = oProtocolOspfAuthProfile.Password
					}
					if oProtocolOspfAuthProfile.Md5 != nil {
						nestedProtocolOspfAuthProfile.Md5 = []ProtocolOspfAuthProfileMd5Xml{}
						for _, oProtocolOspfAuthProfileMd5 := range oProtocolOspfAuthProfile.Md5 {
							nestedProtocolOspfAuthProfileMd5 := ProtocolOspfAuthProfileMd5Xml{}
							if _, ok := o.Misc["ProtocolOspfAuthProfileMd5"]; ok {
								nestedProtocolOspfAuthProfileMd5.Misc = o.Misc["ProtocolOspfAuthProfileMd5"]
							}
							if oProtocolOspfAuthProfileMd5.Preferred != nil {
								nestedProtocolOspfAuthProfileMd5.Preferred = util.YesNo(oProtocolOspfAuthProfileMd5.Preferred, nil)
							}
							if oProtocolOspfAuthProfileMd5.Name != "" {
								nestedProtocolOspfAuthProfileMd5.Name = oProtocolOspfAuthProfileMd5.Name
							}
							if oProtocolOspfAuthProfileMd5.Key != nil {
								nestedProtocolOspfAuthProfileMd5.Key = oProtocolOspfAuthProfileMd5.Key
							}
							nestedProtocolOspfAuthProfile.Md5 = append(nestedProtocolOspfAuthProfile.Md5, nestedProtocolOspfAuthProfileMd5)
						}
					}
					nestedProtocol.Ospf.AuthProfile = append(nestedProtocol.Ospf.AuthProfile, nestedProtocolOspfAuthProfile)
				}
			}
			if o.Protocol.Ospf.ExportRules != nil {
				nestedProtocol.Ospf.ExportRules = []ProtocolOspfExportRulesXml{}
				for _, oProtocolOspfExportRules := range o.Protocol.Ospf.ExportRules {
					nestedProtocolOspfExportRules := ProtocolOspfExportRulesXml{}
					if _, ok := o.Misc["ProtocolOspfExportRules"]; ok {
						nestedProtocolOspfExportRules.Misc = o.Misc["ProtocolOspfExportRules"]
					}
					if oProtocolOspfExportRules.NewTag != nil {
						nestedProtocolOspfExportRules.NewTag = oProtocolOspfExportRules.NewTag
					}
					if oProtocolOspfExportRules.Metric != nil {
						nestedProtocolOspfExportRules.Metric = oProtocolOspfExportRules.Metric
					}
					if oProtocolOspfExportRules.Name != "" {
						nestedProtocolOspfExportRules.Name = oProtocolOspfExportRules.Name
					}
					if oProtocolOspfExportRules.NewPathType != nil {
						nestedProtocolOspfExportRules.NewPathType = oProtocolOspfExportRules.NewPathType
					}
					nestedProtocol.Ospf.ExportRules = append(nestedProtocol.Ospf.ExportRules, nestedProtocolOspfExportRules)
				}
			}
			if o.Protocol.Ospf.GlobalBfd != nil {
				nestedProtocol.Ospf.GlobalBfd = &ProtocolOspfGlobalBfdXml{}
				if _, ok := o.Misc["ProtocolOspfGlobalBfd"]; ok {
					nestedProtocol.Ospf.GlobalBfd.Misc = o.Misc["ProtocolOspfGlobalBfd"]
				}
				if o.Protocol.Ospf.GlobalBfd.Profile != nil {
					nestedProtocol.Ospf.GlobalBfd.Profile = o.Protocol.Ospf.GlobalBfd.Profile
				}
			}
			if o.Protocol.Ospf.RouterId != nil {
				nestedProtocol.Ospf.RouterId = o.Protocol.Ospf.RouterId
			}
		}
		if o.Protocol.Ospfv3 != nil {
			nestedProtocol.Ospfv3 = &ProtocolOspfv3Xml{}
			if _, ok := o.Misc["ProtocolOspfv3"]; ok {
				nestedProtocol.Ospfv3.Misc = o.Misc["ProtocolOspfv3"]
			}
			if o.Protocol.Ospfv3.Area != nil {
				nestedProtocol.Ospfv3.Area = []ProtocolOspfv3AreaXml{}
				for _, oProtocolOspfv3Area := range o.Protocol.Ospfv3.Area {
					nestedProtocolOspfv3Area := ProtocolOspfv3AreaXml{}
					if _, ok := o.Misc["ProtocolOspfv3Area"]; ok {
						nestedProtocolOspfv3Area.Misc = o.Misc["ProtocolOspfv3Area"]
					}
					if oProtocolOspfv3Area.Range != nil {
						nestedProtocolOspfv3Area.Range = []ProtocolOspfv3AreaRangeXml{}
						for _, oProtocolOspfv3AreaRange := range oProtocolOspfv3Area.Range {
							nestedProtocolOspfv3AreaRange := ProtocolOspfv3AreaRangeXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaRange"]; ok {
								nestedProtocolOspfv3AreaRange.Misc = o.Misc["ProtocolOspfv3AreaRange"]
							}
							if oProtocolOspfv3AreaRange.Name != "" {
								nestedProtocolOspfv3AreaRange.Name = oProtocolOspfv3AreaRange.Name
							}
							if oProtocolOspfv3AreaRange.Advertise != nil {
								nestedProtocolOspfv3AreaRange.Advertise = &ProtocolOspfv3AreaRangeAdvertiseXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaRangeAdvertise"]; ok {
									nestedProtocolOspfv3AreaRange.Advertise.Misc = o.Misc["ProtocolOspfv3AreaRangeAdvertise"]
								}
							}
							if oProtocolOspfv3AreaRange.Suppress != nil {
								nestedProtocolOspfv3AreaRange.Suppress = &ProtocolOspfv3AreaRangeSuppressXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaRangeSuppress"]; ok {
									nestedProtocolOspfv3AreaRange.Suppress.Misc = o.Misc["ProtocolOspfv3AreaRangeSuppress"]
								}
							}
							nestedProtocolOspfv3Area.Range = append(nestedProtocolOspfv3Area.Range, nestedProtocolOspfv3AreaRange)
						}
					}
					if oProtocolOspfv3Area.Interface != nil {
						nestedProtocolOspfv3Area.Interface = []ProtocolOspfv3AreaInterfaceXml{}
						for _, oProtocolOspfv3AreaInterface := range oProtocolOspfv3Area.Interface {
							nestedProtocolOspfv3AreaInterface := ProtocolOspfv3AreaInterfaceXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaInterface"]; ok {
								nestedProtocolOspfv3AreaInterface.Misc = o.Misc["ProtocolOspfv3AreaInterface"]
							}
							if oProtocolOspfv3AreaInterface.LinkType != nil {
								nestedProtocolOspfv3AreaInterface.LinkType = &ProtocolOspfv3AreaInterfaceLinkTypeXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceLinkType"]; ok {
									nestedProtocolOspfv3AreaInterface.LinkType.Misc = o.Misc["ProtocolOspfv3AreaInterfaceLinkType"]
								}
								if oProtocolOspfv3AreaInterface.LinkType.Broadcast != nil {
									nestedProtocolOspfv3AreaInterface.LinkType.Broadcast = &ProtocolOspfv3AreaInterfaceLinkTypeBroadcastXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeBroadcast"]; ok {
										nestedProtocolOspfv3AreaInterface.LinkType.Broadcast.Misc = o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeBroadcast"]
									}
								}
								if oProtocolOspfv3AreaInterface.LinkType.P2p != nil {
									nestedProtocolOspfv3AreaInterface.LinkType.P2p = &ProtocolOspfv3AreaInterfaceLinkTypeP2pXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2p"]; ok {
										nestedProtocolOspfv3AreaInterface.LinkType.P2p.Misc = o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2p"]
									}
								}
								if oProtocolOspfv3AreaInterface.LinkType.P2mp != nil {
									nestedProtocolOspfv3AreaInterface.LinkType.P2mp = &ProtocolOspfv3AreaInterfaceLinkTypeP2mpXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2mp"]; ok {
										nestedProtocolOspfv3AreaInterface.LinkType.P2mp.Misc = o.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2mp"]
									}
								}
							}
							if oProtocolOspfv3AreaInterface.Enable != nil {
								nestedProtocolOspfv3AreaInterface.Enable = util.YesNo(oProtocolOspfv3AreaInterface.Enable, nil)
							}
							if oProtocolOspfv3AreaInterface.Passive != nil {
								nestedProtocolOspfv3AreaInterface.Passive = util.YesNo(oProtocolOspfv3AreaInterface.Passive, nil)
							}
							if oProtocolOspfv3AreaInterface.GrDelay != nil {
								nestedProtocolOspfv3AreaInterface.GrDelay = oProtocolOspfv3AreaInterface.GrDelay
							}
							if oProtocolOspfv3AreaInterface.TransitDelay != nil {
								nestedProtocolOspfv3AreaInterface.TransitDelay = oProtocolOspfv3AreaInterface.TransitDelay
							}
							if oProtocolOspfv3AreaInterface.Name != "" {
								nestedProtocolOspfv3AreaInterface.Name = oProtocolOspfv3AreaInterface.Name
							}
							if oProtocolOspfv3AreaInterface.InstanceId != nil {
								nestedProtocolOspfv3AreaInterface.InstanceId = oProtocolOspfv3AreaInterface.InstanceId
							}
							if oProtocolOspfv3AreaInterface.Metric != nil {
								nestedProtocolOspfv3AreaInterface.Metric = oProtocolOspfv3AreaInterface.Metric
							}
							if oProtocolOspfv3AreaInterface.HelloInterval != nil {
								nestedProtocolOspfv3AreaInterface.HelloInterval = oProtocolOspfv3AreaInterface.HelloInterval
							}
							if oProtocolOspfv3AreaInterface.Priority != nil {
								nestedProtocolOspfv3AreaInterface.Priority = oProtocolOspfv3AreaInterface.Priority
							}
							if oProtocolOspfv3AreaInterface.Neighbor != nil {
								nestedProtocolOspfv3AreaInterface.Neighbor = []ProtocolOspfv3AreaInterfaceNeighborXml{}
								for _, oProtocolOspfv3AreaInterfaceNeighbor := range oProtocolOspfv3AreaInterface.Neighbor {
									nestedProtocolOspfv3AreaInterfaceNeighbor := ProtocolOspfv3AreaInterfaceNeighborXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceNeighbor"]; ok {
										nestedProtocolOspfv3AreaInterfaceNeighbor.Misc = o.Misc["ProtocolOspfv3AreaInterfaceNeighbor"]
									}
									if oProtocolOspfv3AreaInterfaceNeighbor.Name != "" {
										nestedProtocolOspfv3AreaInterfaceNeighbor.Name = oProtocolOspfv3AreaInterfaceNeighbor.Name
									}
									nestedProtocolOspfv3AreaInterface.Neighbor = append(nestedProtocolOspfv3AreaInterface.Neighbor, nestedProtocolOspfv3AreaInterfaceNeighbor)
								}
							}
							if oProtocolOspfv3AreaInterface.Bfd != nil {
								nestedProtocolOspfv3AreaInterface.Bfd = &ProtocolOspfv3AreaInterfaceBfdXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaInterfaceBfd"]; ok {
									nestedProtocolOspfv3AreaInterface.Bfd.Misc = o.Misc["ProtocolOspfv3AreaInterfaceBfd"]
								}
								if oProtocolOspfv3AreaInterface.Bfd.Profile != nil {
									nestedProtocolOspfv3AreaInterface.Bfd.Profile = oProtocolOspfv3AreaInterface.Bfd.Profile
								}
							}
							if oProtocolOspfv3AreaInterface.DeadCounts != nil {
								nestedProtocolOspfv3AreaInterface.DeadCounts = oProtocolOspfv3AreaInterface.DeadCounts
							}
							if oProtocolOspfv3AreaInterface.RetransmitInterval != nil {
								nestedProtocolOspfv3AreaInterface.RetransmitInterval = oProtocolOspfv3AreaInterface.RetransmitInterval
							}
							if oProtocolOspfv3AreaInterface.Authentication != nil {
								nestedProtocolOspfv3AreaInterface.Authentication = oProtocolOspfv3AreaInterface.Authentication
							}
							nestedProtocolOspfv3Area.Interface = append(nestedProtocolOspfv3Area.Interface, nestedProtocolOspfv3AreaInterface)
						}
					}
					if oProtocolOspfv3Area.VirtualLink != nil {
						nestedProtocolOspfv3Area.VirtualLink = []ProtocolOspfv3AreaVirtualLinkXml{}
						for _, oProtocolOspfv3AreaVirtualLink := range oProtocolOspfv3Area.VirtualLink {
							nestedProtocolOspfv3AreaVirtualLink := ProtocolOspfv3AreaVirtualLinkXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaVirtualLink"]; ok {
								nestedProtocolOspfv3AreaVirtualLink.Misc = o.Misc["ProtocolOspfv3AreaVirtualLink"]
							}
							if oProtocolOspfv3AreaVirtualLink.DeadCounts != nil {
								nestedProtocolOspfv3AreaVirtualLink.DeadCounts = oProtocolOspfv3AreaVirtualLink.DeadCounts
							}
							if oProtocolOspfv3AreaVirtualLink.RetransmitInterval != nil {
								nestedProtocolOspfv3AreaVirtualLink.RetransmitInterval = oProtocolOspfv3AreaVirtualLink.RetransmitInterval
							}
							if oProtocolOspfv3AreaVirtualLink.Authentication != nil {
								nestedProtocolOspfv3AreaVirtualLink.Authentication = oProtocolOspfv3AreaVirtualLink.Authentication
							}
							if oProtocolOspfv3AreaVirtualLink.Name != "" {
								nestedProtocolOspfv3AreaVirtualLink.Name = oProtocolOspfv3AreaVirtualLink.Name
							}
							if oProtocolOspfv3AreaVirtualLink.TransitAreaId != nil {
								nestedProtocolOspfv3AreaVirtualLink.TransitAreaId = oProtocolOspfv3AreaVirtualLink.TransitAreaId
							}
							if oProtocolOspfv3AreaVirtualLink.InstanceId != nil {
								nestedProtocolOspfv3AreaVirtualLink.InstanceId = oProtocolOspfv3AreaVirtualLink.InstanceId
							}
							if oProtocolOspfv3AreaVirtualLink.HelloInterval != nil {
								nestedProtocolOspfv3AreaVirtualLink.HelloInterval = oProtocolOspfv3AreaVirtualLink.HelloInterval
							}
							if oProtocolOspfv3AreaVirtualLink.TransitDelay != nil {
								nestedProtocolOspfv3AreaVirtualLink.TransitDelay = oProtocolOspfv3AreaVirtualLink.TransitDelay
							}
							if oProtocolOspfv3AreaVirtualLink.Bfd != nil {
								nestedProtocolOspfv3AreaVirtualLink.Bfd = &ProtocolOspfv3AreaVirtualLinkBfdXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaVirtualLinkBfd"]; ok {
									nestedProtocolOspfv3AreaVirtualLink.Bfd.Misc = o.Misc["ProtocolOspfv3AreaVirtualLinkBfd"]
								}
								if oProtocolOspfv3AreaVirtualLink.Bfd.Profile != nil {
									nestedProtocolOspfv3AreaVirtualLink.Bfd.Profile = oProtocolOspfv3AreaVirtualLink.Bfd.Profile
								}
							}
							if oProtocolOspfv3AreaVirtualLink.NeighborId != nil {
								nestedProtocolOspfv3AreaVirtualLink.NeighborId = oProtocolOspfv3AreaVirtualLink.NeighborId
							}
							if oProtocolOspfv3AreaVirtualLink.Enable != nil {
								nestedProtocolOspfv3AreaVirtualLink.Enable = util.YesNo(oProtocolOspfv3AreaVirtualLink.Enable, nil)
							}
							nestedProtocolOspfv3Area.VirtualLink = append(nestedProtocolOspfv3Area.VirtualLink, nestedProtocolOspfv3AreaVirtualLink)
						}
					}
					if oProtocolOspfv3Area.Name != "" {
						nestedProtocolOspfv3Area.Name = oProtocolOspfv3Area.Name
					}
					if oProtocolOspfv3Area.Authentication != nil {
						nestedProtocolOspfv3Area.Authentication = oProtocolOspfv3Area.Authentication
					}
					if oProtocolOspfv3Area.Type != nil {
						nestedProtocolOspfv3Area.Type = &ProtocolOspfv3AreaTypeXml{}
						if _, ok := o.Misc["ProtocolOspfv3AreaType"]; ok {
							nestedProtocolOspfv3Area.Type.Misc = o.Misc["ProtocolOspfv3AreaType"]
						}
						if oProtocolOspfv3Area.Type.Normal != nil {
							nestedProtocolOspfv3Area.Type.Normal = &ProtocolOspfv3AreaTypeNormalXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaTypeNormal"]; ok {
								nestedProtocolOspfv3Area.Type.Normal.Misc = o.Misc["ProtocolOspfv3AreaTypeNormal"]
							}
						}
						if oProtocolOspfv3Area.Type.Stub != nil {
							nestedProtocolOspfv3Area.Type.Stub = &ProtocolOspfv3AreaTypeStubXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaTypeStub"]; ok {
								nestedProtocolOspfv3Area.Type.Stub.Misc = o.Misc["ProtocolOspfv3AreaTypeStub"]
							}
							if oProtocolOspfv3Area.Type.Stub.DefaultRoute != nil {
								nestedProtocolOspfv3Area.Type.Stub.DefaultRoute = &ProtocolOspfv3AreaTypeStubDefaultRouteXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaTypeStubDefaultRoute"]; ok {
									nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Misc = o.Misc["ProtocolOspfv3AreaTypeStubDefaultRoute"]
								}
								if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise != nil {
									nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise = &ProtocolOspfv3AreaTypeStubDefaultRouteAdvertiseXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise"]; ok {
										nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Misc = o.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise"]
									}
									if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric != nil {
										nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric = oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric
									}
								}
								if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable != nil {
									nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable = &ProtocolOspfv3AreaTypeStubDefaultRouteDisableXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteDisable"]; ok {
										nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable.Misc = o.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteDisable"]
									}
								}
							}
							if oProtocolOspfv3Area.Type.Stub.AcceptSummary != nil {
								nestedProtocolOspfv3Area.Type.Stub.AcceptSummary = util.YesNo(oProtocolOspfv3Area.Type.Stub.AcceptSummary, nil)
							}
						}
						if oProtocolOspfv3Area.Type.Nssa != nil {
							nestedProtocolOspfv3Area.Type.Nssa = &ProtocolOspfv3AreaTypeNssaXml{}
							if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssa"]; ok {
								nestedProtocolOspfv3Area.Type.Nssa.Misc = o.Misc["ProtocolOspfv3AreaTypeNssa"]
							}
							if oProtocolOspfv3Area.Type.Nssa.AcceptSummary != nil {
								nestedProtocolOspfv3Area.Type.Nssa.AcceptSummary = util.YesNo(oProtocolOspfv3Area.Type.Nssa.AcceptSummary, nil)
							}
							if oProtocolOspfv3Area.Type.Nssa.DefaultRoute != nil {
								nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute = &ProtocolOspfv3AreaTypeNssaDefaultRouteXml{}
								if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRoute"]; ok {
									nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRoute"]
								}
								if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable != nil {
									nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable = &ProtocolOspfv3AreaTypeNssaDefaultRouteDisableXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteDisable"]; ok {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteDisable"]
									}
								}
								if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise != nil {
									nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise = &ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertiseXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise"]; ok {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise"]
									}
									if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type != nil {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type
									}
									if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric != nil {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric
									}
								}
							}
							if oProtocolOspfv3Area.Type.Nssa.NssaExtRange != nil {
								nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange = []ProtocolOspfv3AreaTypeNssaNssaExtRangeXml{}
								for _, oProtocolOspfv3AreaTypeNssaNssaExtRange := range oProtocolOspfv3Area.Type.Nssa.NssaExtRange {
									nestedProtocolOspfv3AreaTypeNssaNssaExtRange := ProtocolOspfv3AreaTypeNssaNssaExtRangeXml{}
									if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRange"]; ok {
										nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRange"]
									}
									if oProtocolOspfv3AreaTypeNssaNssaExtRange.Name != "" {
										nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Name = oProtocolOspfv3AreaTypeNssaNssaExtRange.Name
									}
									if oProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise != nil {
										nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise = &ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertiseXml{}
										if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise"]; ok {
											nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise"]
										}
									}
									if oProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress != nil {
										nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress = &ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppressXml{}
										if _, ok := o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress"]; ok {
											nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress.Misc = o.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress"]
										}
									}
									nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange = append(nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange, nestedProtocolOspfv3AreaTypeNssaNssaExtRange)
								}
							}
						}
					}
					nestedProtocol.Ospfv3.Area = append(nestedProtocol.Ospfv3.Area, nestedProtocolOspfv3Area)
				}
			}
			if o.Protocol.Ospfv3.DisableTransitTraffic != nil {
				nestedProtocol.Ospfv3.DisableTransitTraffic = util.YesNo(o.Protocol.Ospfv3.DisableTransitTraffic, nil)
			}
			if o.Protocol.Ospfv3.Enable != nil {
				nestedProtocol.Ospfv3.Enable = util.YesNo(o.Protocol.Ospfv3.Enable, nil)
			}
			if o.Protocol.Ospfv3.ExportRules != nil {
				nestedProtocol.Ospfv3.ExportRules = []ProtocolOspfv3ExportRulesXml{}
				for _, oProtocolOspfv3ExportRules := range o.Protocol.Ospfv3.ExportRules {
					nestedProtocolOspfv3ExportRules := ProtocolOspfv3ExportRulesXml{}
					if _, ok := o.Misc["ProtocolOspfv3ExportRules"]; ok {
						nestedProtocolOspfv3ExportRules.Misc = o.Misc["ProtocolOspfv3ExportRules"]
					}
					if oProtocolOspfv3ExportRules.NewPathType != nil {
						nestedProtocolOspfv3ExportRules.NewPathType = oProtocolOspfv3ExportRules.NewPathType
					}
					if oProtocolOspfv3ExportRules.NewTag != nil {
						nestedProtocolOspfv3ExportRules.NewTag = oProtocolOspfv3ExportRules.NewTag
					}
					if oProtocolOspfv3ExportRules.Metric != nil {
						nestedProtocolOspfv3ExportRules.Metric = oProtocolOspfv3ExportRules.Metric
					}
					if oProtocolOspfv3ExportRules.Name != "" {
						nestedProtocolOspfv3ExportRules.Name = oProtocolOspfv3ExportRules.Name
					}
					nestedProtocol.Ospfv3.ExportRules = append(nestedProtocol.Ospfv3.ExportRules, nestedProtocolOspfv3ExportRules)
				}
			}
			if o.Protocol.Ospfv3.GlobalBfd != nil {
				nestedProtocol.Ospfv3.GlobalBfd = &ProtocolOspfv3GlobalBfdXml{}
				if _, ok := o.Misc["ProtocolOspfv3GlobalBfd"]; ok {
					nestedProtocol.Ospfv3.GlobalBfd.Misc = o.Misc["ProtocolOspfv3GlobalBfd"]
				}
				if o.Protocol.Ospfv3.GlobalBfd.Profile != nil {
					nestedProtocol.Ospfv3.GlobalBfd.Profile = o.Protocol.Ospfv3.GlobalBfd.Profile
				}
			}
			if o.Protocol.Ospfv3.RejectDefaultRoute != nil {
				nestedProtocol.Ospfv3.RejectDefaultRoute = util.YesNo(o.Protocol.Ospfv3.RejectDefaultRoute, nil)
			}
			if o.Protocol.Ospfv3.AllowRedistDefaultRoute != nil {
				nestedProtocol.Ospfv3.AllowRedistDefaultRoute = util.YesNo(o.Protocol.Ospfv3.AllowRedistDefaultRoute, nil)
			}
			if o.Protocol.Ospfv3.AuthProfile != nil {
				nestedProtocol.Ospfv3.AuthProfile = []ProtocolOspfv3AuthProfileXml{}
				for _, oProtocolOspfv3AuthProfile := range o.Protocol.Ospfv3.AuthProfile {
					nestedProtocolOspfv3AuthProfile := ProtocolOspfv3AuthProfileXml{}
					if _, ok := o.Misc["ProtocolOspfv3AuthProfile"]; ok {
						nestedProtocolOspfv3AuthProfile.Misc = o.Misc["ProtocolOspfv3AuthProfile"]
					}
					if oProtocolOspfv3AuthProfile.Spi != nil {
						nestedProtocolOspfv3AuthProfile.Spi = oProtocolOspfv3AuthProfile.Spi
					}
					if oProtocolOspfv3AuthProfile.Name != "" {
						nestedProtocolOspfv3AuthProfile.Name = oProtocolOspfv3AuthProfile.Name
					}
					if oProtocolOspfv3AuthProfile.Esp != nil {
						nestedProtocolOspfv3AuthProfile.Esp = &ProtocolOspfv3AuthProfileEspXml{}
						if _, ok := o.Misc["ProtocolOspfv3AuthProfileEsp"]; ok {
							nestedProtocolOspfv3AuthProfile.Esp.Misc = o.Misc["ProtocolOspfv3AuthProfileEsp"]
						}
						if oProtocolOspfv3AuthProfile.Esp.Authentication != nil {
							nestedProtocolOspfv3AuthProfile.Esp.Authentication = &ProtocolOspfv3AuthProfileEspAuthenticationXml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthentication"]; ok {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthentication"]
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.Md5 != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Md5 = &ProtocolOspfv3AuthProfileEspAuthenticationMd5Xml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationMd5"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationMd5"]
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1 != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha1 = &ProtocolOspfv3AuthProfileEspAuthenticationSha1Xml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha1"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha1"]
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256 != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha256 = &ProtocolOspfv3AuthProfileEspAuthenticationSha256Xml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha256"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha256"]
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384 != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha384 = &ProtocolOspfv3AuthProfileEspAuthenticationSha384Xml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha384"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha384"]
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512 != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha512 = &ProtocolOspfv3AuthProfileEspAuthenticationSha512Xml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha512"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha512"]
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication.None != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication.None = &ProtocolOspfv3AuthProfileEspAuthenticationNoneXml{}
								if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationNone"]; ok {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.None.Misc = o.Misc["ProtocolOspfv3AuthProfileEspAuthenticationNone"]
								}
							}
						}
						if oProtocolOspfv3AuthProfile.Esp.Encryption != nil {
							nestedProtocolOspfv3AuthProfile.Esp.Encryption = &ProtocolOspfv3AuthProfileEspEncryptionXml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileEspEncryption"]; ok {
								nestedProtocolOspfv3AuthProfile.Esp.Encryption.Misc = o.Misc["ProtocolOspfv3AuthProfileEspEncryption"]
							}
							if oProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm = oProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm
							}
							if oProtocolOspfv3AuthProfile.Esp.Encryption.Key != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Encryption.Key = oProtocolOspfv3AuthProfile.Esp.Encryption.Key
							}
						}
					}
					if oProtocolOspfv3AuthProfile.Ah != nil {
						nestedProtocolOspfv3AuthProfile.Ah = &ProtocolOspfv3AuthProfileAhXml{}
						if _, ok := o.Misc["ProtocolOspfv3AuthProfileAh"]; ok {
							nestedProtocolOspfv3AuthProfile.Ah.Misc = o.Misc["ProtocolOspfv3AuthProfileAh"]
						}
						if oProtocolOspfv3AuthProfile.Ah.Sha1 != nil {
							nestedProtocolOspfv3AuthProfile.Ah.Sha1 = &ProtocolOspfv3AuthProfileAhSha1Xml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileAhSha1"]; ok {
								nestedProtocolOspfv3AuthProfile.Ah.Sha1.Misc = o.Misc["ProtocolOspfv3AuthProfileAhSha1"]
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha1.Key != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha1.Key = oProtocolOspfv3AuthProfile.Ah.Sha1.Key
							}
						}
						if oProtocolOspfv3AuthProfile.Ah.Sha256 != nil {
							nestedProtocolOspfv3AuthProfile.Ah.Sha256 = &ProtocolOspfv3AuthProfileAhSha256Xml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileAhSha256"]; ok {
								nestedProtocolOspfv3AuthProfile.Ah.Sha256.Misc = o.Misc["ProtocolOspfv3AuthProfileAhSha256"]
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha256.Key != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha256.Key = oProtocolOspfv3AuthProfile.Ah.Sha256.Key
							}
						}
						if oProtocolOspfv3AuthProfile.Ah.Sha384 != nil {
							nestedProtocolOspfv3AuthProfile.Ah.Sha384 = &ProtocolOspfv3AuthProfileAhSha384Xml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileAhSha384"]; ok {
								nestedProtocolOspfv3AuthProfile.Ah.Sha384.Misc = o.Misc["ProtocolOspfv3AuthProfileAhSha384"]
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha384.Key != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha384.Key = oProtocolOspfv3AuthProfile.Ah.Sha384.Key
							}
						}
						if oProtocolOspfv3AuthProfile.Ah.Sha512 != nil {
							nestedProtocolOspfv3AuthProfile.Ah.Sha512 = &ProtocolOspfv3AuthProfileAhSha512Xml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileAhSha512"]; ok {
								nestedProtocolOspfv3AuthProfile.Ah.Sha512.Misc = o.Misc["ProtocolOspfv3AuthProfileAhSha512"]
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha512.Key != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha512.Key = oProtocolOspfv3AuthProfile.Ah.Sha512.Key
							}
						}
						if oProtocolOspfv3AuthProfile.Ah.Md5 != nil {
							nestedProtocolOspfv3AuthProfile.Ah.Md5 = &ProtocolOspfv3AuthProfileAhMd5Xml{}
							if _, ok := o.Misc["ProtocolOspfv3AuthProfileAhMd5"]; ok {
								nestedProtocolOspfv3AuthProfile.Ah.Md5.Misc = o.Misc["ProtocolOspfv3AuthProfileAhMd5"]
							}
							if oProtocolOspfv3AuthProfile.Ah.Md5.Key != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Md5.Key = oProtocolOspfv3AuthProfile.Ah.Md5.Key
							}
						}
					}
					nestedProtocol.Ospfv3.AuthProfile = append(nestedProtocol.Ospfv3.AuthProfile, nestedProtocolOspfv3AuthProfile)
				}
			}
			if o.Protocol.Ospfv3.GracefulRestart != nil {
				nestedProtocol.Ospfv3.GracefulRestart = &ProtocolOspfv3GracefulRestartXml{}
				if _, ok := o.Misc["ProtocolOspfv3GracefulRestart"]; ok {
					nestedProtocol.Ospfv3.GracefulRestart.Misc = o.Misc["ProtocolOspfv3GracefulRestart"]
				}
				if o.Protocol.Ospfv3.GracefulRestart.Enable != nil {
					nestedProtocol.Ospfv3.GracefulRestart.Enable = util.YesNo(o.Protocol.Ospfv3.GracefulRestart.Enable, nil)
				}
				if o.Protocol.Ospfv3.GracefulRestart.GracePeriod != nil {
					nestedProtocol.Ospfv3.GracefulRestart.GracePeriod = o.Protocol.Ospfv3.GracefulRestart.GracePeriod
				}
				if o.Protocol.Ospfv3.GracefulRestart.HelperEnable != nil {
					nestedProtocol.Ospfv3.GracefulRestart.HelperEnable = util.YesNo(o.Protocol.Ospfv3.GracefulRestart.HelperEnable, nil)
				}
				if o.Protocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
					nestedProtocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime = o.Protocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime
				}
				if o.Protocol.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
					nestedProtocol.Ospfv3.GracefulRestart.StrictLSAChecking = util.YesNo(o.Protocol.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
				}
			}
			if o.Protocol.Ospfv3.RouterId != nil {
				nestedProtocol.Ospfv3.RouterId = o.Protocol.Ospfv3.RouterId
			}
			if o.Protocol.Ospfv3.Timers != nil {
				nestedProtocol.Ospfv3.Timers = &ProtocolOspfv3TimersXml{}
				if _, ok := o.Misc["ProtocolOspfv3Timers"]; ok {
					nestedProtocol.Ospfv3.Timers.Misc = o.Misc["ProtocolOspfv3Timers"]
				}
				if o.Protocol.Ospfv3.Timers.LsaInterval != nil {
					nestedProtocol.Ospfv3.Timers.LsaInterval = o.Protocol.Ospfv3.Timers.LsaInterval
				}
				if o.Protocol.Ospfv3.Timers.SpfCalculationDelay != nil {
					nestedProtocol.Ospfv3.Timers.SpfCalculationDelay = o.Protocol.Ospfv3.Timers.SpfCalculationDelay
				}
			}
		}
		if o.Protocol.RedistProfile != nil {
			nestedProtocol.RedistProfile = []ProtocolRedistProfileXml{}
			for _, oProtocolRedistProfile := range o.Protocol.RedistProfile {
				nestedProtocolRedistProfile := ProtocolRedistProfileXml{}
				if _, ok := o.Misc["ProtocolRedistProfile"]; ok {
					nestedProtocolRedistProfile.Misc = o.Misc["ProtocolRedistProfile"]
				}
				if oProtocolRedistProfile.Priority != nil {
					nestedProtocolRedistProfile.Priority = oProtocolRedistProfile.Priority
				}
				if oProtocolRedistProfile.Filter != nil {
					nestedProtocolRedistProfile.Filter = &ProtocolRedistProfileFilterXml{}
					if _, ok := o.Misc["ProtocolRedistProfileFilter"]; ok {
						nestedProtocolRedistProfile.Filter.Misc = o.Misc["ProtocolRedistProfileFilter"]
					}
					if oProtocolRedistProfile.Filter.Ospf != nil {
						nestedProtocolRedistProfile.Filter.Ospf = &ProtocolRedistProfileFilterOspfXml{}
						if _, ok := o.Misc["ProtocolRedistProfileFilterOspf"]; ok {
							nestedProtocolRedistProfile.Filter.Ospf.Misc = o.Misc["ProtocolRedistProfileFilterOspf"]
						}
						if oProtocolRedistProfile.Filter.Ospf.PathType != nil {
							nestedProtocolRedistProfile.Filter.Ospf.PathType = util.StrToMem(oProtocolRedistProfile.Filter.Ospf.PathType)
						}
						if oProtocolRedistProfile.Filter.Ospf.Area != nil {
							nestedProtocolRedistProfile.Filter.Ospf.Area = util.StrToMem(oProtocolRedistProfile.Filter.Ospf.Area)
						}
						if oProtocolRedistProfile.Filter.Ospf.Tag != nil {
							nestedProtocolRedistProfile.Filter.Ospf.Tag = util.StrToMem(oProtocolRedistProfile.Filter.Ospf.Tag)
						}
					}
					if oProtocolRedistProfile.Filter.Bgp != nil {
						nestedProtocolRedistProfile.Filter.Bgp = &ProtocolRedistProfileFilterBgpXml{}
						if _, ok := o.Misc["ProtocolRedistProfileFilterBgp"]; ok {
							nestedProtocolRedistProfile.Filter.Bgp.Misc = o.Misc["ProtocolRedistProfileFilterBgp"]
						}
						if oProtocolRedistProfile.Filter.Bgp.Community != nil {
							nestedProtocolRedistProfile.Filter.Bgp.Community = util.StrToMem(oProtocolRedistProfile.Filter.Bgp.Community)
						}
						if oProtocolRedistProfile.Filter.Bgp.ExtendedCommunity != nil {
							nestedProtocolRedistProfile.Filter.Bgp.ExtendedCommunity = util.StrToMem(oProtocolRedistProfile.Filter.Bgp.ExtendedCommunity)
						}
					}
					if oProtocolRedistProfile.Filter.Type != nil {
						nestedProtocolRedistProfile.Filter.Type = util.StrToMem(oProtocolRedistProfile.Filter.Type)
					}
					if oProtocolRedistProfile.Filter.Interface != nil {
						nestedProtocolRedistProfile.Filter.Interface = util.StrToMem(oProtocolRedistProfile.Filter.Interface)
					}
					if oProtocolRedistProfile.Filter.Destination != nil {
						nestedProtocolRedistProfile.Filter.Destination = util.StrToMem(oProtocolRedistProfile.Filter.Destination)
					}
					if oProtocolRedistProfile.Filter.Nexthop != nil {
						nestedProtocolRedistProfile.Filter.Nexthop = util.StrToMem(oProtocolRedistProfile.Filter.Nexthop)
					}
				}
				if oProtocolRedistProfile.Action != nil {
					nestedProtocolRedistProfile.Action = &ProtocolRedistProfileActionXml{}
					if _, ok := o.Misc["ProtocolRedistProfileAction"]; ok {
						nestedProtocolRedistProfile.Action.Misc = o.Misc["ProtocolRedistProfileAction"]
					}
					if oProtocolRedistProfile.Action.NoRedist != nil {
						nestedProtocolRedistProfile.Action.NoRedist = &ProtocolRedistProfileActionNoRedistXml{}
						if _, ok := o.Misc["ProtocolRedistProfileActionNoRedist"]; ok {
							nestedProtocolRedistProfile.Action.NoRedist.Misc = o.Misc["ProtocolRedistProfileActionNoRedist"]
						}
					}
					if oProtocolRedistProfile.Action.Redist != nil {
						nestedProtocolRedistProfile.Action.Redist = &ProtocolRedistProfileActionRedistXml{}
						if _, ok := o.Misc["ProtocolRedistProfileActionRedist"]; ok {
							nestedProtocolRedistProfile.Action.Redist.Misc = o.Misc["ProtocolRedistProfileActionRedist"]
						}
					}
				}
				if oProtocolRedistProfile.Name != "" {
					nestedProtocolRedistProfile.Name = oProtocolRedistProfile.Name
				}
				nestedProtocol.RedistProfile = append(nestedProtocol.RedistProfile, nestedProtocolRedistProfile)
			}
		}
	}
	entry.Protocol = nestedProtocol

	var nestedRoutingTable *RoutingTableXml
	if o.RoutingTable != nil {
		nestedRoutingTable = &RoutingTableXml{}
		if _, ok := o.Misc["RoutingTable"]; ok {
			nestedRoutingTable.Misc = o.Misc["RoutingTable"]
		}
		if o.RoutingTable.Ip != nil {
			nestedRoutingTable.Ip = &RoutingTableIpXml{}
			if _, ok := o.Misc["RoutingTableIp"]; ok {
				nestedRoutingTable.Ip.Misc = o.Misc["RoutingTableIp"]
			}
			if o.RoutingTable.Ip.StaticRoute != nil {
				nestedRoutingTable.Ip.StaticRoute = []RoutingTableIpStaticRouteXml{}
				for _, oRoutingTableIpStaticRoute := range o.RoutingTable.Ip.StaticRoute {
					nestedRoutingTableIpStaticRoute := RoutingTableIpStaticRouteXml{}
					if _, ok := o.Misc["RoutingTableIpStaticRoute"]; ok {
						nestedRoutingTableIpStaticRoute.Misc = o.Misc["RoutingTableIpStaticRoute"]
					}
					if oRoutingTableIpStaticRoute.RouteTable != nil {
						nestedRoutingTableIpStaticRoute.RouteTable = &RoutingTableIpStaticRouteRouteTableXml{}
						if _, ok := o.Misc["RoutingTableIpStaticRouteRouteTable"]; ok {
							nestedRoutingTableIpStaticRoute.RouteTable.Misc = o.Misc["RoutingTableIpStaticRouteRouteTable"]
						}
						if oRoutingTableIpStaticRoute.RouteTable.NoInstall != nil {
							nestedRoutingTableIpStaticRoute.RouteTable.NoInstall = &RoutingTableIpStaticRouteRouteTableNoInstallXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteRouteTableNoInstall"]; ok {
								nestedRoutingTableIpStaticRoute.RouteTable.NoInstall.Misc = o.Misc["RoutingTableIpStaticRouteRouteTableNoInstall"]
							}
						}
						if oRoutingTableIpStaticRoute.RouteTable.Unicast != nil {
							nestedRoutingTableIpStaticRoute.RouteTable.Unicast = &RoutingTableIpStaticRouteRouteTableUnicastXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteRouteTableUnicast"]; ok {
								nestedRoutingTableIpStaticRoute.RouteTable.Unicast.Misc = o.Misc["RoutingTableIpStaticRouteRouteTableUnicast"]
							}
						}
						if oRoutingTableIpStaticRoute.RouteTable.Multicast != nil {
							nestedRoutingTableIpStaticRoute.RouteTable.Multicast = &RoutingTableIpStaticRouteRouteTableMulticastXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteRouteTableMulticast"]; ok {
								nestedRoutingTableIpStaticRoute.RouteTable.Multicast.Misc = o.Misc["RoutingTableIpStaticRouteRouteTableMulticast"]
							}
						}
						if oRoutingTableIpStaticRoute.RouteTable.Both != nil {
							nestedRoutingTableIpStaticRoute.RouteTable.Both = &RoutingTableIpStaticRouteRouteTableBothXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteRouteTableBoth"]; ok {
								nestedRoutingTableIpStaticRoute.RouteTable.Both.Misc = o.Misc["RoutingTableIpStaticRouteRouteTableBoth"]
							}
						}
					}
					if oRoutingTableIpStaticRoute.Interface != nil {
						nestedRoutingTableIpStaticRoute.Interface = oRoutingTableIpStaticRoute.Interface
					}
					if oRoutingTableIpStaticRoute.AdminDist != nil {
						nestedRoutingTableIpStaticRoute.AdminDist = oRoutingTableIpStaticRoute.AdminDist
					}
					if oRoutingTableIpStaticRoute.Nexthop != nil {
						nestedRoutingTableIpStaticRoute.Nexthop = &RoutingTableIpStaticRouteNexthopXml{}
						if _, ok := o.Misc["RoutingTableIpStaticRouteNexthop"]; ok {
							nestedRoutingTableIpStaticRoute.Nexthop.Misc = o.Misc["RoutingTableIpStaticRouteNexthop"]
						}
						if oRoutingTableIpStaticRoute.Nexthop.NextVr != nil {
							nestedRoutingTableIpStaticRoute.Nexthop.NextVr = oRoutingTableIpStaticRoute.Nexthop.NextVr
						}
						if oRoutingTableIpStaticRoute.Nexthop.Receive != nil {
							nestedRoutingTableIpStaticRoute.Nexthop.Receive = &RoutingTableIpStaticRouteNexthopReceiveXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteNexthopReceive"]; ok {
								nestedRoutingTableIpStaticRoute.Nexthop.Receive.Misc = o.Misc["RoutingTableIpStaticRouteNexthopReceive"]
							}
						}
						if oRoutingTableIpStaticRoute.Nexthop.Discard != nil {
							nestedRoutingTableIpStaticRoute.Nexthop.Discard = &RoutingTableIpStaticRouteNexthopDiscardXml{}
							if _, ok := o.Misc["RoutingTableIpStaticRouteNexthopDiscard"]; ok {
								nestedRoutingTableIpStaticRoute.Nexthop.Discard.Misc = o.Misc["RoutingTableIpStaticRouteNexthopDiscard"]
							}
						}
						if oRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
							nestedRoutingTableIpStaticRoute.Nexthop.IpAddress = oRoutingTableIpStaticRoute.Nexthop.IpAddress
						}
						if oRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
							nestedRoutingTableIpStaticRoute.Nexthop.Fqdn = oRoutingTableIpStaticRoute.Nexthop.Fqdn
						}
					}
					if oRoutingTableIpStaticRoute.Bfd != nil {
						nestedRoutingTableIpStaticRoute.Bfd = &RoutingTableIpStaticRouteBfdXml{}
						if _, ok := o.Misc["RoutingTableIpStaticRouteBfd"]; ok {
							nestedRoutingTableIpStaticRoute.Bfd.Misc = o.Misc["RoutingTableIpStaticRouteBfd"]
						}
						if oRoutingTableIpStaticRoute.Bfd.Profile != nil {
							nestedRoutingTableIpStaticRoute.Bfd.Profile = oRoutingTableIpStaticRoute.Bfd.Profile
						}
					}
					if oRoutingTableIpStaticRoute.PathMonitor != nil {
						nestedRoutingTableIpStaticRoute.PathMonitor = &RoutingTableIpStaticRoutePathMonitorXml{}
						if _, ok := o.Misc["RoutingTableIpStaticRoutePathMonitor"]; ok {
							nestedRoutingTableIpStaticRoute.PathMonitor.Misc = o.Misc["RoutingTableIpStaticRoutePathMonitor"]
						}
						if oRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
							nestedRoutingTableIpStaticRoute.PathMonitor.HoldTime = oRoutingTableIpStaticRoute.PathMonitor.HoldTime
						}
						if oRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
							nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []RoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml{}
							for _, oRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
								nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations := RoutingTableIpStaticRoutePathMonitorMonitorDestinationsXml{}
								if _, ok := o.Misc["RoutingTableIpStaticRoutePathMonitorMonitorDestinations"]; ok {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["RoutingTableIpStaticRoutePathMonitorMonitorDestinations"]
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
								}
								if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
								}
								nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
							}
						}
						if oRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
							nestedRoutingTableIpStaticRoute.PathMonitor.Enable = util.YesNo(oRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
						}
						if oRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
							nestedRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oRoutingTableIpStaticRoute.PathMonitor.FailureCondition
						}
					}
					if oRoutingTableIpStaticRoute.Name != "" {
						nestedRoutingTableIpStaticRoute.Name = oRoutingTableIpStaticRoute.Name
					}
					if oRoutingTableIpStaticRoute.Destination != nil {
						nestedRoutingTableIpStaticRoute.Destination = oRoutingTableIpStaticRoute.Destination
					}
					if oRoutingTableIpStaticRoute.Metric != nil {
						nestedRoutingTableIpStaticRoute.Metric = oRoutingTableIpStaticRoute.Metric
					}
					nestedRoutingTable.Ip.StaticRoute = append(nestedRoutingTable.Ip.StaticRoute, nestedRoutingTableIpStaticRoute)
				}
			}
		}
		if o.RoutingTable.Ipv6 != nil {
			nestedRoutingTable.Ipv6 = &RoutingTableIpv6Xml{}
			if _, ok := o.Misc["RoutingTableIpv6"]; ok {
				nestedRoutingTable.Ipv6.Misc = o.Misc["RoutingTableIpv6"]
			}
			if o.RoutingTable.Ipv6.StaticRoute != nil {
				nestedRoutingTable.Ipv6.StaticRoute = []RoutingTableIpv6StaticRouteXml{}
				for _, oRoutingTableIpv6StaticRoute := range o.RoutingTable.Ipv6.StaticRoute {
					nestedRoutingTableIpv6StaticRoute := RoutingTableIpv6StaticRouteXml{}
					if _, ok := o.Misc["RoutingTableIpv6StaticRoute"]; ok {
						nestedRoutingTableIpv6StaticRoute.Misc = o.Misc["RoutingTableIpv6StaticRoute"]
					}
					if oRoutingTableIpv6StaticRoute.AdminDist != nil {
						nestedRoutingTableIpv6StaticRoute.AdminDist = oRoutingTableIpv6StaticRoute.AdminDist
					}
					if oRoutingTableIpv6StaticRoute.Metric != nil {
						nestedRoutingTableIpv6StaticRoute.Metric = oRoutingTableIpv6StaticRoute.Metric
					}
					if oRoutingTableIpv6StaticRoute.Nexthop != nil {
						nestedRoutingTableIpv6StaticRoute.Nexthop = &RoutingTableIpv6StaticRouteNexthopXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRouteNexthop"]; ok {
							nestedRoutingTableIpv6StaticRoute.Nexthop.Misc = o.Misc["RoutingTableIpv6StaticRouteNexthop"]
						}
						if oRoutingTableIpv6StaticRoute.Nexthop.Receive != nil {
							nestedRoutingTableIpv6StaticRoute.Nexthop.Receive = &RoutingTableIpv6StaticRouteNexthopReceiveXml{}
							if _, ok := o.Misc["RoutingTableIpv6StaticRouteNexthopReceive"]; ok {
								nestedRoutingTableIpv6StaticRoute.Nexthop.Receive.Misc = o.Misc["RoutingTableIpv6StaticRouteNexthopReceive"]
							}
						}
						if oRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
							nestedRoutingTableIpv6StaticRoute.Nexthop.Discard = &RoutingTableIpv6StaticRouteNexthopDiscardXml{}
							if _, ok := o.Misc["RoutingTableIpv6StaticRouteNexthopDiscard"]; ok {
								nestedRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc = o.Misc["RoutingTableIpv6StaticRouteNexthopDiscard"]
							}
						}
						if oRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
							nestedRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
						}
						if oRoutingTableIpv6StaticRoute.Nexthop.NextVr != nil {
							nestedRoutingTableIpv6StaticRoute.Nexthop.NextVr = oRoutingTableIpv6StaticRoute.Nexthop.NextVr
						}
					}
					if oRoutingTableIpv6StaticRoute.Option != nil {
						nestedRoutingTableIpv6StaticRoute.Option = &RoutingTableIpv6StaticRouteOptionXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRouteOption"]; ok {
							nestedRoutingTableIpv6StaticRoute.Option.Misc = o.Misc["RoutingTableIpv6StaticRouteOption"]
						}
					}
					if oRoutingTableIpv6StaticRoute.RouteTable != nil {
						nestedRoutingTableIpv6StaticRoute.RouteTable = &RoutingTableIpv6StaticRouteRouteTableXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRouteRouteTable"]; ok {
							nestedRoutingTableIpv6StaticRoute.RouteTable.Misc = o.Misc["RoutingTableIpv6StaticRouteRouteTable"]
						}
						if oRoutingTableIpv6StaticRoute.RouteTable.Unicast != nil {
							nestedRoutingTableIpv6StaticRoute.RouteTable.Unicast = &RoutingTableIpv6StaticRouteRouteTableUnicastXml{}
							if _, ok := o.Misc["RoutingTableIpv6StaticRouteRouteTableUnicast"]; ok {
								nestedRoutingTableIpv6StaticRoute.RouteTable.Unicast.Misc = o.Misc["RoutingTableIpv6StaticRouteRouteTableUnicast"]
							}
						}
						if oRoutingTableIpv6StaticRoute.RouteTable.NoInstall != nil {
							nestedRoutingTableIpv6StaticRoute.RouteTable.NoInstall = &RoutingTableIpv6StaticRouteRouteTableNoInstallXml{}
							if _, ok := o.Misc["RoutingTableIpv6StaticRouteRouteTableNoInstall"]; ok {
								nestedRoutingTableIpv6StaticRoute.RouteTable.NoInstall.Misc = o.Misc["RoutingTableIpv6StaticRouteRouteTableNoInstall"]
							}
						}
					}
					if oRoutingTableIpv6StaticRoute.PathMonitor != nil {
						nestedRoutingTableIpv6StaticRoute.PathMonitor = &RoutingTableIpv6StaticRoutePathMonitorXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRoutePathMonitor"]; ok {
							nestedRoutingTableIpv6StaticRoute.PathMonitor.Misc = o.Misc["RoutingTableIpv6StaticRoutePathMonitor"]
						}
						if oRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
							nestedRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.YesNo(oRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
						}
						if oRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
							nestedRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
						}
						if oRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
							nestedRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
						}
						if oRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
							nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml{}
							for _, oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
								nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := RoutingTableIpv6StaticRoutePathMonitorMonitorDestinationsXml{}
								if _, ok := o.Misc["RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]; ok {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc = o.Misc["RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"]
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.YesNo(oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
								}
								if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
								}
								nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
							}
						}
					}
					if oRoutingTableIpv6StaticRoute.Destination != nil {
						nestedRoutingTableIpv6StaticRoute.Destination = oRoutingTableIpv6StaticRoute.Destination
					}
					if oRoutingTableIpv6StaticRoute.Interface != nil {
						nestedRoutingTableIpv6StaticRoute.Interface = oRoutingTableIpv6StaticRoute.Interface
					}
					if oRoutingTableIpv6StaticRoute.Bfd != nil {
						nestedRoutingTableIpv6StaticRoute.Bfd = &RoutingTableIpv6StaticRouteBfdXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRouteBfd"]; ok {
							nestedRoutingTableIpv6StaticRoute.Bfd.Misc = o.Misc["RoutingTableIpv6StaticRouteBfd"]
						}
						if oRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
							nestedRoutingTableIpv6StaticRoute.Bfd.Profile = oRoutingTableIpv6StaticRoute.Bfd.Profile
						}
					}
					if oRoutingTableIpv6StaticRoute.Name != "" {
						nestedRoutingTableIpv6StaticRoute.Name = oRoutingTableIpv6StaticRoute.Name
					}
					nestedRoutingTable.Ipv6.StaticRoute = append(nestedRoutingTable.Ipv6.StaticRoute, nestedRoutingTableIpv6StaticRoute)
				}
			}
		}
	}
	entry.RoutingTable = nestedRoutingTable

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
		var nestedAdminDists *AdminDists
		if o.AdminDists != nil {
			nestedAdminDists = &AdminDists{}
			if o.AdminDists.Misc != nil {
				entry.Misc["AdminDists"] = o.AdminDists.Misc
			}
			if o.AdminDists.OspfExt != nil {
				nestedAdminDists.OspfExt = o.AdminDists.OspfExt
			}
			if o.AdminDists.Ospfv3Ext != nil {
				nestedAdminDists.Ospfv3Ext = o.AdminDists.Ospfv3Ext
			}
			if o.AdminDists.Ospfv3Int != nil {
				nestedAdminDists.Ospfv3Int = o.AdminDists.Ospfv3Int
			}
			if o.AdminDists.Rip != nil {
				nestedAdminDists.Rip = o.AdminDists.Rip
			}
			if o.AdminDists.Ebgp != nil {
				nestedAdminDists.Ebgp = o.AdminDists.Ebgp
			}
			if o.AdminDists.Ibgp != nil {
				nestedAdminDists.Ibgp = o.AdminDists.Ibgp
			}
			if o.AdminDists.OspfInt != nil {
				nestedAdminDists.OspfInt = o.AdminDists.OspfInt
			}
			if o.AdminDists.Static != nil {
				nestedAdminDists.Static = o.AdminDists.Static
			}
			if o.AdminDists.StaticIpv6 != nil {
				nestedAdminDists.StaticIpv6 = o.AdminDists.StaticIpv6
			}
		}
		entry.AdminDists = nestedAdminDists

		var nestedEcmp *Ecmp
		if o.Ecmp != nil {
			nestedEcmp = &Ecmp{}
			if o.Ecmp.Misc != nil {
				entry.Misc["Ecmp"] = o.Ecmp.Misc
			}
			if o.Ecmp.Algorithm != nil {
				nestedEcmp.Algorithm = &EcmpAlgorithm{}
				if o.Ecmp.Algorithm.Misc != nil {
					entry.Misc["EcmpAlgorithm"] = o.Ecmp.Algorithm.Misc
				}
				if o.Ecmp.Algorithm.IpModulo != nil {
					nestedEcmp.Algorithm.IpModulo = &EcmpAlgorithmIpModulo{}
					if o.Ecmp.Algorithm.IpModulo.Misc != nil {
						entry.Misc["EcmpAlgorithmIpModulo"] = o.Ecmp.Algorithm.IpModulo.Misc
					}
				}
				if o.Ecmp.Algorithm.WeightedRoundRobin != nil {
					nestedEcmp.Algorithm.WeightedRoundRobin = &EcmpAlgorithmWeightedRoundRobin{}
					if o.Ecmp.Algorithm.WeightedRoundRobin.Misc != nil {
						entry.Misc["EcmpAlgorithmWeightedRoundRobin"] = o.Ecmp.Algorithm.WeightedRoundRobin.Misc
					}
					if o.Ecmp.Algorithm.WeightedRoundRobin.Interface != nil {
						nestedEcmp.Algorithm.WeightedRoundRobin.Interface = []EcmpAlgorithmWeightedRoundRobinInterface{}
						for _, oEcmpAlgorithmWeightedRoundRobinInterface := range o.Ecmp.Algorithm.WeightedRoundRobin.Interface {
							nestedEcmpAlgorithmWeightedRoundRobinInterface := EcmpAlgorithmWeightedRoundRobinInterface{}
							if oEcmpAlgorithmWeightedRoundRobinInterface.Misc != nil {
								entry.Misc["EcmpAlgorithmWeightedRoundRobinInterface"] = oEcmpAlgorithmWeightedRoundRobinInterface.Misc
							}
							if oEcmpAlgorithmWeightedRoundRobinInterface.Weight != nil {
								nestedEcmpAlgorithmWeightedRoundRobinInterface.Weight = oEcmpAlgorithmWeightedRoundRobinInterface.Weight
							}
							if oEcmpAlgorithmWeightedRoundRobinInterface.Name != "" {
								nestedEcmpAlgorithmWeightedRoundRobinInterface.Name = oEcmpAlgorithmWeightedRoundRobinInterface.Name
							}
							nestedEcmp.Algorithm.WeightedRoundRobin.Interface = append(nestedEcmp.Algorithm.WeightedRoundRobin.Interface, nestedEcmpAlgorithmWeightedRoundRobinInterface)
						}
					}
				}
				if o.Ecmp.Algorithm.BalancedRoundRobin != nil {
					nestedEcmp.Algorithm.BalancedRoundRobin = &EcmpAlgorithmBalancedRoundRobin{}
					if o.Ecmp.Algorithm.BalancedRoundRobin.Misc != nil {
						entry.Misc["EcmpAlgorithmBalancedRoundRobin"] = o.Ecmp.Algorithm.BalancedRoundRobin.Misc
					}
				}
				if o.Ecmp.Algorithm.IpHash != nil {
					nestedEcmp.Algorithm.IpHash = &EcmpAlgorithmIpHash{}
					if o.Ecmp.Algorithm.IpHash.Misc != nil {
						entry.Misc["EcmpAlgorithmIpHash"] = o.Ecmp.Algorithm.IpHash.Misc
					}
					if o.Ecmp.Algorithm.IpHash.HashSeed != nil {
						nestedEcmp.Algorithm.IpHash.HashSeed = o.Ecmp.Algorithm.IpHash.HashSeed
					}
					if o.Ecmp.Algorithm.IpHash.SrcOnly != nil {
						nestedEcmp.Algorithm.IpHash.SrcOnly = util.AsBool(o.Ecmp.Algorithm.IpHash.SrcOnly, nil)
					}
					if o.Ecmp.Algorithm.IpHash.UsePort != nil {
						nestedEcmp.Algorithm.IpHash.UsePort = util.AsBool(o.Ecmp.Algorithm.IpHash.UsePort, nil)
					}
				}
			}
			if o.Ecmp.Enable != nil {
				nestedEcmp.Enable = util.AsBool(o.Ecmp.Enable, nil)
			}
			if o.Ecmp.MaxPath != nil {
				nestedEcmp.MaxPath = o.Ecmp.MaxPath
			}
			if o.Ecmp.StrictSourcePath != nil {
				nestedEcmp.StrictSourcePath = util.AsBool(o.Ecmp.StrictSourcePath, nil)
			}
			if o.Ecmp.SymmetricReturn != nil {
				nestedEcmp.SymmetricReturn = util.AsBool(o.Ecmp.SymmetricReturn, nil)
			}
		}
		entry.Ecmp = nestedEcmp

		entry.Interface = util.MemToStr(o.Interface)
		var nestedMulticast *Multicast
		if o.Multicast != nil {
			nestedMulticast = &Multicast{}
			if o.Multicast.Misc != nil {
				entry.Misc["Multicast"] = o.Multicast.Misc
			}
			if o.Multicast.SsmAddressSpace != nil {
				nestedMulticast.SsmAddressSpace = []MulticastSsmAddressSpace{}
				for _, oMulticastSsmAddressSpace := range o.Multicast.SsmAddressSpace {
					nestedMulticastSsmAddressSpace := MulticastSsmAddressSpace{}
					if oMulticastSsmAddressSpace.Misc != nil {
						entry.Misc["MulticastSsmAddressSpace"] = oMulticastSsmAddressSpace.Misc
					}
					if oMulticastSsmAddressSpace.GroupAddress != nil {
						nestedMulticastSsmAddressSpace.GroupAddress = oMulticastSsmAddressSpace.GroupAddress
					}
					if oMulticastSsmAddressSpace.Included != nil {
						nestedMulticastSsmAddressSpace.Included = util.AsBool(oMulticastSsmAddressSpace.Included, nil)
					}
					if oMulticastSsmAddressSpace.Name != "" {
						nestedMulticastSsmAddressSpace.Name = oMulticastSsmAddressSpace.Name
					}
					nestedMulticast.SsmAddressSpace = append(nestedMulticast.SsmAddressSpace, nestedMulticastSsmAddressSpace)
				}
			}
			if o.Multicast.Enable != nil {
				nestedMulticast.Enable = util.AsBool(o.Multicast.Enable, nil)
			}
			if o.Multicast.InterfaceGroup != nil {
				nestedMulticast.InterfaceGroup = []MulticastInterfaceGroup{}
				for _, oMulticastInterfaceGroup := range o.Multicast.InterfaceGroup {
					nestedMulticastInterfaceGroup := MulticastInterfaceGroup{}
					if oMulticastInterfaceGroup.Misc != nil {
						entry.Misc["MulticastInterfaceGroup"] = oMulticastInterfaceGroup.Misc
					}
					if oMulticastInterfaceGroup.Igmp != nil {
						nestedMulticastInterfaceGroup.Igmp = &MulticastInterfaceGroupIgmp{}
						if oMulticastInterfaceGroup.Igmp.Misc != nil {
							entry.Misc["MulticastInterfaceGroupIgmp"] = oMulticastInterfaceGroup.Igmp.Misc
						}
						if oMulticastInterfaceGroup.Igmp.Enable != nil {
							nestedMulticastInterfaceGroup.Igmp.Enable = util.AsBool(oMulticastInterfaceGroup.Igmp.Enable, nil)
						}
						if oMulticastInterfaceGroup.Igmp.QueryInterval != nil {
							nestedMulticastInterfaceGroup.Igmp.QueryInterval = oMulticastInterfaceGroup.Igmp.QueryInterval
						}
						if oMulticastInterfaceGroup.Igmp.LastMemberQueryInterval != nil {
							nestedMulticastInterfaceGroup.Igmp.LastMemberQueryInterval = oMulticastInterfaceGroup.Igmp.LastMemberQueryInterval
						}
						if oMulticastInterfaceGroup.Igmp.ImmediateLeave != nil {
							nestedMulticastInterfaceGroup.Igmp.ImmediateLeave = util.AsBool(oMulticastInterfaceGroup.Igmp.ImmediateLeave, nil)
						}
						if oMulticastInterfaceGroup.Igmp.Robustness != nil {
							nestedMulticastInterfaceGroup.Igmp.Robustness = oMulticastInterfaceGroup.Igmp.Robustness
						}
						if oMulticastInterfaceGroup.Igmp.MaxSources != nil {
							nestedMulticastInterfaceGroup.Igmp.MaxSources = oMulticastInterfaceGroup.Igmp.MaxSources
						}
						if oMulticastInterfaceGroup.Igmp.Version != nil {
							nestedMulticastInterfaceGroup.Igmp.Version = oMulticastInterfaceGroup.Igmp.Version
						}
						if oMulticastInterfaceGroup.Igmp.MaxQueryResponseTime != nil {
							nestedMulticastInterfaceGroup.Igmp.MaxQueryResponseTime = oMulticastInterfaceGroup.Igmp.MaxQueryResponseTime
						}
						if oMulticastInterfaceGroup.Igmp.MaxGroups != nil {
							nestedMulticastInterfaceGroup.Igmp.MaxGroups = oMulticastInterfaceGroup.Igmp.MaxGroups
						}
						if oMulticastInterfaceGroup.Igmp.RouterAlertPolicing != nil {
							nestedMulticastInterfaceGroup.Igmp.RouterAlertPolicing = util.AsBool(oMulticastInterfaceGroup.Igmp.RouterAlertPolicing, nil)
						}
					}
					if oMulticastInterfaceGroup.Pim != nil {
						nestedMulticastInterfaceGroup.Pim = &MulticastInterfaceGroupPim{}
						if oMulticastInterfaceGroup.Pim.Misc != nil {
							entry.Misc["MulticastInterfaceGroupPim"] = oMulticastInterfaceGroup.Pim.Misc
						}
						if oMulticastInterfaceGroup.Pim.BsrBorder != nil {
							nestedMulticastInterfaceGroup.Pim.BsrBorder = util.AsBool(oMulticastInterfaceGroup.Pim.BsrBorder, nil)
						}
						if oMulticastInterfaceGroup.Pim.AllowedNeighbors != nil {
							nestedMulticastInterfaceGroup.Pim.AllowedNeighbors = []MulticastInterfaceGroupPimAllowedNeighbors{}
							for _, oMulticastInterfaceGroupPimAllowedNeighbors := range oMulticastInterfaceGroup.Pim.AllowedNeighbors {
								nestedMulticastInterfaceGroupPimAllowedNeighbors := MulticastInterfaceGroupPimAllowedNeighbors{}
								if oMulticastInterfaceGroupPimAllowedNeighbors.Misc != nil {
									entry.Misc["MulticastInterfaceGroupPimAllowedNeighbors"] = oMulticastInterfaceGroupPimAllowedNeighbors.Misc
								}
								if oMulticastInterfaceGroupPimAllowedNeighbors.Name != "" {
									nestedMulticastInterfaceGroupPimAllowedNeighbors.Name = oMulticastInterfaceGroupPimAllowedNeighbors.Name
								}
								nestedMulticastInterfaceGroup.Pim.AllowedNeighbors = append(nestedMulticastInterfaceGroup.Pim.AllowedNeighbors, nestedMulticastInterfaceGroupPimAllowedNeighbors)
							}
						}
						if oMulticastInterfaceGroup.Pim.Enable != nil {
							nestedMulticastInterfaceGroup.Pim.Enable = util.AsBool(oMulticastInterfaceGroup.Pim.Enable, nil)
						}
						if oMulticastInterfaceGroup.Pim.AssertInterval != nil {
							nestedMulticastInterfaceGroup.Pim.AssertInterval = oMulticastInterfaceGroup.Pim.AssertInterval
						}
						if oMulticastInterfaceGroup.Pim.HelloInterval != nil {
							nestedMulticastInterfaceGroup.Pim.HelloInterval = oMulticastInterfaceGroup.Pim.HelloInterval
						}
						if oMulticastInterfaceGroup.Pim.JoinPruneInterval != nil {
							nestedMulticastInterfaceGroup.Pim.JoinPruneInterval = oMulticastInterfaceGroup.Pim.JoinPruneInterval
						}
						if oMulticastInterfaceGroup.Pim.DrPriority != nil {
							nestedMulticastInterfaceGroup.Pim.DrPriority = oMulticastInterfaceGroup.Pim.DrPriority
						}
					}
					if oMulticastInterfaceGroup.Name != "" {
						nestedMulticastInterfaceGroup.Name = oMulticastInterfaceGroup.Name
					}
					if oMulticastInterfaceGroup.Description != nil {
						nestedMulticastInterfaceGroup.Description = oMulticastInterfaceGroup.Description
					}
					if oMulticastInterfaceGroup.Interface != nil {
						nestedMulticastInterfaceGroup.Interface = util.MemToStr(oMulticastInterfaceGroup.Interface)
					}
					if oMulticastInterfaceGroup.GroupPermission != nil {
						nestedMulticastInterfaceGroup.GroupPermission = &MulticastInterfaceGroupGroupPermission{}
						if oMulticastInterfaceGroup.GroupPermission.Misc != nil {
							entry.Misc["MulticastInterfaceGroupGroupPermission"] = oMulticastInterfaceGroup.GroupPermission.Misc
						}
						if oMulticastInterfaceGroup.GroupPermission.AnySourceMulticast != nil {
							nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast = []MulticastInterfaceGroupGroupPermissionAnySourceMulticast{}
							for _, oMulticastInterfaceGroupGroupPermissionAnySourceMulticast := range oMulticastInterfaceGroup.GroupPermission.AnySourceMulticast {
								nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast := MulticastInterfaceGroupGroupPermissionAnySourceMulticast{}
								if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Misc != nil {
									entry.Misc["MulticastInterfaceGroupGroupPermissionAnySourceMulticast"] = oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Misc
								}
								if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name != "" {
									nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name = oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Name
								}
								if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress != nil {
									nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress = oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.GroupAddress
								}
								if oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included != nil {
									nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included = util.AsBool(oMulticastInterfaceGroupGroupPermissionAnySourceMulticast.Included, nil)
								}
								nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast = append(nestedMulticastInterfaceGroup.GroupPermission.AnySourceMulticast, nestedMulticastInterfaceGroupGroupPermissionAnySourceMulticast)
							}
						}
						if oMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast != nil {
							nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast = []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast{}
							for _, oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast := range oMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast {
								nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast := MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast{}
								if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Misc != nil {
									entry.Misc["MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast"] = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Misc
								}
								if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress != nil {
									nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.SourceAddress
								}
								if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included != nil {
									nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included = util.AsBool(oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Included, nil)
								}
								if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name != "" {
									nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.Name
								}
								if oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress != nil {
									nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress = oMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast.GroupAddress
								}
								nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast = append(nestedMulticastInterfaceGroup.GroupPermission.SourceSpecificMulticast, nestedMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast)
							}
						}
					}
					nestedMulticast.InterfaceGroup = append(nestedMulticast.InterfaceGroup, nestedMulticastInterfaceGroup)
				}
			}
			if o.Multicast.RouteAgeoutTime != nil {
				nestedMulticast.RouteAgeoutTime = o.Multicast.RouteAgeoutTime
			}
			if o.Multicast.Rp != nil {
				nestedMulticast.Rp = &MulticastRp{}
				if o.Multicast.Rp.Misc != nil {
					entry.Misc["MulticastRp"] = o.Multicast.Rp.Misc
				}
				if o.Multicast.Rp.ExternalRp != nil {
					nestedMulticast.Rp.ExternalRp = []MulticastRpExternalRp{}
					for _, oMulticastRpExternalRp := range o.Multicast.Rp.ExternalRp {
						nestedMulticastRpExternalRp := MulticastRpExternalRp{}
						if oMulticastRpExternalRp.Misc != nil {
							entry.Misc["MulticastRpExternalRp"] = oMulticastRpExternalRp.Misc
						}
						if oMulticastRpExternalRp.GroupAddresses != nil {
							nestedMulticastRpExternalRp.GroupAddresses = util.MemToStr(oMulticastRpExternalRp.GroupAddresses)
						}
						if oMulticastRpExternalRp.Override != nil {
							nestedMulticastRpExternalRp.Override = util.AsBool(oMulticastRpExternalRp.Override, nil)
						}
						if oMulticastRpExternalRp.Name != "" {
							nestedMulticastRpExternalRp.Name = oMulticastRpExternalRp.Name
						}
						nestedMulticast.Rp.ExternalRp = append(nestedMulticast.Rp.ExternalRp, nestedMulticastRpExternalRp)
					}
				}
				if o.Multicast.Rp.LocalRp != nil {
					nestedMulticast.Rp.LocalRp = &MulticastRpLocalRp{}
					if o.Multicast.Rp.LocalRp.Misc != nil {
						entry.Misc["MulticastRpLocalRp"] = o.Multicast.Rp.LocalRp.Misc
					}
					if o.Multicast.Rp.LocalRp.CandidateRp != nil {
						nestedMulticast.Rp.LocalRp.CandidateRp = &MulticastRpLocalRpCandidateRp{}
						if o.Multicast.Rp.LocalRp.CandidateRp.Misc != nil {
							entry.Misc["MulticastRpLocalRpCandidateRp"] = o.Multicast.Rp.LocalRp.CandidateRp.Misc
						}
						if o.Multicast.Rp.LocalRp.CandidateRp.GroupAddresses != nil {
							nestedMulticast.Rp.LocalRp.CandidateRp.GroupAddresses = util.MemToStr(o.Multicast.Rp.LocalRp.CandidateRp.GroupAddresses)
						}
						if o.Multicast.Rp.LocalRp.CandidateRp.Interface != nil {
							nestedMulticast.Rp.LocalRp.CandidateRp.Interface = o.Multicast.Rp.LocalRp.CandidateRp.Interface
						}
						if o.Multicast.Rp.LocalRp.CandidateRp.Priority != nil {
							nestedMulticast.Rp.LocalRp.CandidateRp.Priority = o.Multicast.Rp.LocalRp.CandidateRp.Priority
						}
						if o.Multicast.Rp.LocalRp.CandidateRp.Address != nil {
							nestedMulticast.Rp.LocalRp.CandidateRp.Address = o.Multicast.Rp.LocalRp.CandidateRp.Address
						}
						if o.Multicast.Rp.LocalRp.CandidateRp.AdvertisementInterval != nil {
							nestedMulticast.Rp.LocalRp.CandidateRp.AdvertisementInterval = o.Multicast.Rp.LocalRp.CandidateRp.AdvertisementInterval
						}
					}
					if o.Multicast.Rp.LocalRp.StaticRp != nil {
						nestedMulticast.Rp.LocalRp.StaticRp = &MulticastRpLocalRpStaticRp{}
						if o.Multicast.Rp.LocalRp.StaticRp.Misc != nil {
							entry.Misc["MulticastRpLocalRpStaticRp"] = o.Multicast.Rp.LocalRp.StaticRp.Misc
						}
						if o.Multicast.Rp.LocalRp.StaticRp.Address != nil {
							nestedMulticast.Rp.LocalRp.StaticRp.Address = o.Multicast.Rp.LocalRp.StaticRp.Address
						}
						if o.Multicast.Rp.LocalRp.StaticRp.GroupAddresses != nil {
							nestedMulticast.Rp.LocalRp.StaticRp.GroupAddresses = util.MemToStr(o.Multicast.Rp.LocalRp.StaticRp.GroupAddresses)
						}
						if o.Multicast.Rp.LocalRp.StaticRp.Interface != nil {
							nestedMulticast.Rp.LocalRp.StaticRp.Interface = o.Multicast.Rp.LocalRp.StaticRp.Interface
						}
						if o.Multicast.Rp.LocalRp.StaticRp.Override != nil {
							nestedMulticast.Rp.LocalRp.StaticRp.Override = util.AsBool(o.Multicast.Rp.LocalRp.StaticRp.Override, nil)
						}
					}
				}
			}
			if o.Multicast.SptThreshold != nil {
				nestedMulticast.SptThreshold = []MulticastSptThreshold{}
				for _, oMulticastSptThreshold := range o.Multicast.SptThreshold {
					nestedMulticastSptThreshold := MulticastSptThreshold{}
					if oMulticastSptThreshold.Misc != nil {
						entry.Misc["MulticastSptThreshold"] = oMulticastSptThreshold.Misc
					}
					if oMulticastSptThreshold.Threshold != nil {
						nestedMulticastSptThreshold.Threshold = oMulticastSptThreshold.Threshold
					}
					if oMulticastSptThreshold.Name != "" {
						nestedMulticastSptThreshold.Name = oMulticastSptThreshold.Name
					}
					nestedMulticast.SptThreshold = append(nestedMulticast.SptThreshold, nestedMulticastSptThreshold)
				}
			}
		}
		entry.Multicast = nestedMulticast

		var nestedProtocol *Protocol
		if o.Protocol != nil {
			nestedProtocol = &Protocol{}
			if o.Protocol.Misc != nil {
				entry.Misc["Protocol"] = o.Protocol.Misc
			}
			if o.Protocol.RedistProfile != nil {
				nestedProtocol.RedistProfile = []ProtocolRedistProfile{}
				for _, oProtocolRedistProfile := range o.Protocol.RedistProfile {
					nestedProtocolRedistProfile := ProtocolRedistProfile{}
					if oProtocolRedistProfile.Misc != nil {
						entry.Misc["ProtocolRedistProfile"] = oProtocolRedistProfile.Misc
					}
					if oProtocolRedistProfile.Filter != nil {
						nestedProtocolRedistProfile.Filter = &ProtocolRedistProfileFilter{}
						if oProtocolRedistProfile.Filter.Misc != nil {
							entry.Misc["ProtocolRedistProfileFilter"] = oProtocolRedistProfile.Filter.Misc
						}
						if oProtocolRedistProfile.Filter.Type != nil {
							nestedProtocolRedistProfile.Filter.Type = util.MemToStr(oProtocolRedistProfile.Filter.Type)
						}
						if oProtocolRedistProfile.Filter.Interface != nil {
							nestedProtocolRedistProfile.Filter.Interface = util.MemToStr(oProtocolRedistProfile.Filter.Interface)
						}
						if oProtocolRedistProfile.Filter.Destination != nil {
							nestedProtocolRedistProfile.Filter.Destination = util.MemToStr(oProtocolRedistProfile.Filter.Destination)
						}
						if oProtocolRedistProfile.Filter.Nexthop != nil {
							nestedProtocolRedistProfile.Filter.Nexthop = util.MemToStr(oProtocolRedistProfile.Filter.Nexthop)
						}
						if oProtocolRedistProfile.Filter.Ospf != nil {
							nestedProtocolRedistProfile.Filter.Ospf = &ProtocolRedistProfileFilterOspf{}
							if oProtocolRedistProfile.Filter.Ospf.Misc != nil {
								entry.Misc["ProtocolRedistProfileFilterOspf"] = oProtocolRedistProfile.Filter.Ospf.Misc
							}
							if oProtocolRedistProfile.Filter.Ospf.Tag != nil {
								nestedProtocolRedistProfile.Filter.Ospf.Tag = util.MemToStr(oProtocolRedistProfile.Filter.Ospf.Tag)
							}
							if oProtocolRedistProfile.Filter.Ospf.PathType != nil {
								nestedProtocolRedistProfile.Filter.Ospf.PathType = util.MemToStr(oProtocolRedistProfile.Filter.Ospf.PathType)
							}
							if oProtocolRedistProfile.Filter.Ospf.Area != nil {
								nestedProtocolRedistProfile.Filter.Ospf.Area = util.MemToStr(oProtocolRedistProfile.Filter.Ospf.Area)
							}
						}
						if oProtocolRedistProfile.Filter.Bgp != nil {
							nestedProtocolRedistProfile.Filter.Bgp = &ProtocolRedistProfileFilterBgp{}
							if oProtocolRedistProfile.Filter.Bgp.Misc != nil {
								entry.Misc["ProtocolRedistProfileFilterBgp"] = oProtocolRedistProfile.Filter.Bgp.Misc
							}
							if oProtocolRedistProfile.Filter.Bgp.Community != nil {
								nestedProtocolRedistProfile.Filter.Bgp.Community = util.MemToStr(oProtocolRedistProfile.Filter.Bgp.Community)
							}
							if oProtocolRedistProfile.Filter.Bgp.ExtendedCommunity != nil {
								nestedProtocolRedistProfile.Filter.Bgp.ExtendedCommunity = util.MemToStr(oProtocolRedistProfile.Filter.Bgp.ExtendedCommunity)
							}
						}
					}
					if oProtocolRedistProfile.Action != nil {
						nestedProtocolRedistProfile.Action = &ProtocolRedistProfileAction{}
						if oProtocolRedistProfile.Action.Misc != nil {
							entry.Misc["ProtocolRedistProfileAction"] = oProtocolRedistProfile.Action.Misc
						}
						if oProtocolRedistProfile.Action.Redist != nil {
							nestedProtocolRedistProfile.Action.Redist = &ProtocolRedistProfileActionRedist{}
							if oProtocolRedistProfile.Action.Redist.Misc != nil {
								entry.Misc["ProtocolRedistProfileActionRedist"] = oProtocolRedistProfile.Action.Redist.Misc
							}
						}
						if oProtocolRedistProfile.Action.NoRedist != nil {
							nestedProtocolRedistProfile.Action.NoRedist = &ProtocolRedistProfileActionNoRedist{}
							if oProtocolRedistProfile.Action.NoRedist.Misc != nil {
								entry.Misc["ProtocolRedistProfileActionNoRedist"] = oProtocolRedistProfile.Action.NoRedist.Misc
							}
						}
					}
					if oProtocolRedistProfile.Name != "" {
						nestedProtocolRedistProfile.Name = oProtocolRedistProfile.Name
					}
					if oProtocolRedistProfile.Priority != nil {
						nestedProtocolRedistProfile.Priority = oProtocolRedistProfile.Priority
					}
					nestedProtocol.RedistProfile = append(nestedProtocol.RedistProfile, nestedProtocolRedistProfile)
				}
			}
			if o.Protocol.RedistProfileIpv6 != nil {
				nestedProtocol.RedistProfileIpv6 = []ProtocolRedistProfileIpv6{}
				for _, oProtocolRedistProfileIpv6 := range o.Protocol.RedistProfileIpv6 {
					nestedProtocolRedistProfileIpv6 := ProtocolRedistProfileIpv6{}
					if oProtocolRedistProfileIpv6.Misc != nil {
						entry.Misc["ProtocolRedistProfileIpv6"] = oProtocolRedistProfileIpv6.Misc
					}
					if oProtocolRedistProfileIpv6.Priority != nil {
						nestedProtocolRedistProfileIpv6.Priority = oProtocolRedistProfileIpv6.Priority
					}
					if oProtocolRedistProfileIpv6.Filter != nil {
						nestedProtocolRedistProfileIpv6.Filter = &ProtocolRedistProfileIpv6Filter{}
						if oProtocolRedistProfileIpv6.Filter.Misc != nil {
							entry.Misc["ProtocolRedistProfileIpv6Filter"] = oProtocolRedistProfileIpv6.Filter.Misc
						}
						if oProtocolRedistProfileIpv6.Filter.Nexthop != nil {
							nestedProtocolRedistProfileIpv6.Filter.Nexthop = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Nexthop)
						}
						if oProtocolRedistProfileIpv6.Filter.Ospfv3 != nil {
							nestedProtocolRedistProfileIpv6.Filter.Ospfv3 = &ProtocolRedistProfileIpv6FilterOspfv3{}
							if oProtocolRedistProfileIpv6.Filter.Ospfv3.Misc != nil {
								entry.Misc["ProtocolRedistProfileIpv6FilterOspfv3"] = oProtocolRedistProfileIpv6.Filter.Ospfv3.Misc
							}
							if oProtocolRedistProfileIpv6.Filter.Ospfv3.PathType != nil {
								nestedProtocolRedistProfileIpv6.Filter.Ospfv3.PathType = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Ospfv3.PathType)
							}
							if oProtocolRedistProfileIpv6.Filter.Ospfv3.Area != nil {
								nestedProtocolRedistProfileIpv6.Filter.Ospfv3.Area = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Ospfv3.Area)
							}
							if oProtocolRedistProfileIpv6.Filter.Ospfv3.Tag != nil {
								nestedProtocolRedistProfileIpv6.Filter.Ospfv3.Tag = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Ospfv3.Tag)
							}
						}
						if oProtocolRedistProfileIpv6.Filter.Bgp != nil {
							nestedProtocolRedistProfileIpv6.Filter.Bgp = &ProtocolRedistProfileIpv6FilterBgp{}
							if oProtocolRedistProfileIpv6.Filter.Bgp.Misc != nil {
								entry.Misc["ProtocolRedistProfileIpv6FilterBgp"] = oProtocolRedistProfileIpv6.Filter.Bgp.Misc
							}
							if oProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity != nil {
								nestedProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Bgp.ExtendedCommunity)
							}
							if oProtocolRedistProfileIpv6.Filter.Bgp.Community != nil {
								nestedProtocolRedistProfileIpv6.Filter.Bgp.Community = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Bgp.Community)
							}
						}
						if oProtocolRedistProfileIpv6.Filter.Type != nil {
							nestedProtocolRedistProfileIpv6.Filter.Type = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Type)
						}
						if oProtocolRedistProfileIpv6.Filter.Interface != nil {
							nestedProtocolRedistProfileIpv6.Filter.Interface = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Interface)
						}
						if oProtocolRedistProfileIpv6.Filter.Destination != nil {
							nestedProtocolRedistProfileIpv6.Filter.Destination = util.MemToStr(oProtocolRedistProfileIpv6.Filter.Destination)
						}
					}
					if oProtocolRedistProfileIpv6.Action != nil {
						nestedProtocolRedistProfileIpv6.Action = &ProtocolRedistProfileIpv6Action{}
						if oProtocolRedistProfileIpv6.Action.Misc != nil {
							entry.Misc["ProtocolRedistProfileIpv6Action"] = oProtocolRedistProfileIpv6.Action.Misc
						}
						if oProtocolRedistProfileIpv6.Action.NoRedist != nil {
							nestedProtocolRedistProfileIpv6.Action.NoRedist = &ProtocolRedistProfileIpv6ActionNoRedist{}
							if oProtocolRedistProfileIpv6.Action.NoRedist.Misc != nil {
								entry.Misc["ProtocolRedistProfileIpv6ActionNoRedist"] = oProtocolRedistProfileIpv6.Action.NoRedist.Misc
							}
						}
						if oProtocolRedistProfileIpv6.Action.Redist != nil {
							nestedProtocolRedistProfileIpv6.Action.Redist = &ProtocolRedistProfileIpv6ActionRedist{}
							if oProtocolRedistProfileIpv6.Action.Redist.Misc != nil {
								entry.Misc["ProtocolRedistProfileIpv6ActionRedist"] = oProtocolRedistProfileIpv6.Action.Redist.Misc
							}
						}
					}
					if oProtocolRedistProfileIpv6.Name != "" {
						nestedProtocolRedistProfileIpv6.Name = oProtocolRedistProfileIpv6.Name
					}
					nestedProtocol.RedistProfileIpv6 = append(nestedProtocol.RedistProfileIpv6, nestedProtocolRedistProfileIpv6)
				}
			}
			if o.Protocol.Rip != nil {
				nestedProtocol.Rip = &ProtocolRip{}
				if o.Protocol.Rip.Misc != nil {
					entry.Misc["ProtocolRip"] = o.Protocol.Rip.Misc
				}
				if o.Protocol.Rip.ExportRules != nil {
					nestedProtocol.Rip.ExportRules = []ProtocolRipExportRules{}
					for _, oProtocolRipExportRules := range o.Protocol.Rip.ExportRules {
						nestedProtocolRipExportRules := ProtocolRipExportRules{}
						if oProtocolRipExportRules.Misc != nil {
							entry.Misc["ProtocolRipExportRules"] = oProtocolRipExportRules.Misc
						}
						if oProtocolRipExportRules.Metric != nil {
							nestedProtocolRipExportRules.Metric = oProtocolRipExportRules.Metric
						}
						if oProtocolRipExportRules.Name != "" {
							nestedProtocolRipExportRules.Name = oProtocolRipExportRules.Name
						}
						nestedProtocol.Rip.ExportRules = append(nestedProtocol.Rip.ExportRules, nestedProtocolRipExportRules)
					}
				}
				if o.Protocol.Rip.GlobalBfd != nil {
					nestedProtocol.Rip.GlobalBfd = &ProtocolRipGlobalBfd{}
					if o.Protocol.Rip.GlobalBfd.Misc != nil {
						entry.Misc["ProtocolRipGlobalBfd"] = o.Protocol.Rip.GlobalBfd.Misc
					}
					if o.Protocol.Rip.GlobalBfd.Profile != nil {
						nestedProtocol.Rip.GlobalBfd.Profile = o.Protocol.Rip.GlobalBfd.Profile
					}
				}
				if o.Protocol.Rip.Interface != nil {
					nestedProtocol.Rip.Interface = []ProtocolRipInterface{}
					for _, oProtocolRipInterface := range o.Protocol.Rip.Interface {
						nestedProtocolRipInterface := ProtocolRipInterface{}
						if oProtocolRipInterface.Misc != nil {
							entry.Misc["ProtocolRipInterface"] = oProtocolRipInterface.Misc
						}
						if oProtocolRipInterface.Enable != nil {
							nestedProtocolRipInterface.Enable = util.AsBool(oProtocolRipInterface.Enable, nil)
						}
						if oProtocolRipInterface.Authentication != nil {
							nestedProtocolRipInterface.Authentication = oProtocolRipInterface.Authentication
						}
						if oProtocolRipInterface.Mode != nil {
							nestedProtocolRipInterface.Mode = oProtocolRipInterface.Mode
						}
						if oProtocolRipInterface.DefaultRoute != nil {
							nestedProtocolRipInterface.DefaultRoute = &ProtocolRipInterfaceDefaultRoute{}
							if oProtocolRipInterface.DefaultRoute.Misc != nil {
								entry.Misc["ProtocolRipInterfaceDefaultRoute"] = oProtocolRipInterface.DefaultRoute.Misc
							}
							if oProtocolRipInterface.DefaultRoute.Disable != nil {
								nestedProtocolRipInterface.DefaultRoute.Disable = &ProtocolRipInterfaceDefaultRouteDisable{}
								if oProtocolRipInterface.DefaultRoute.Disable.Misc != nil {
									entry.Misc["ProtocolRipInterfaceDefaultRouteDisable"] = oProtocolRipInterface.DefaultRoute.Disable.Misc
								}
							}
							if oProtocolRipInterface.DefaultRoute.Advertise != nil {
								nestedProtocolRipInterface.DefaultRoute.Advertise = &ProtocolRipInterfaceDefaultRouteAdvertise{}
								if oProtocolRipInterface.DefaultRoute.Advertise.Misc != nil {
									entry.Misc["ProtocolRipInterfaceDefaultRouteAdvertise"] = oProtocolRipInterface.DefaultRoute.Advertise.Misc
								}
								if oProtocolRipInterface.DefaultRoute.Advertise.Metric != nil {
									nestedProtocolRipInterface.DefaultRoute.Advertise.Metric = oProtocolRipInterface.DefaultRoute.Advertise.Metric
								}
							}
						}
						if oProtocolRipInterface.Bfd != nil {
							nestedProtocolRipInterface.Bfd = &ProtocolRipInterfaceBfd{}
							if oProtocolRipInterface.Bfd.Misc != nil {
								entry.Misc["ProtocolRipInterfaceBfd"] = oProtocolRipInterface.Bfd.Misc
							}
							if oProtocolRipInterface.Bfd.Profile != nil {
								nestedProtocolRipInterface.Bfd.Profile = oProtocolRipInterface.Bfd.Profile
							}
						}
						if oProtocolRipInterface.Name != "" {
							nestedProtocolRipInterface.Name = oProtocolRipInterface.Name
						}
						nestedProtocol.Rip.Interface = append(nestedProtocol.Rip.Interface, nestedProtocolRipInterface)
					}
				}
				if o.Protocol.Rip.RejectDefaultRoute != nil {
					nestedProtocol.Rip.RejectDefaultRoute = util.AsBool(o.Protocol.Rip.RejectDefaultRoute, nil)
				}
				if o.Protocol.Rip.Timers != nil {
					nestedProtocol.Rip.Timers = &ProtocolRipTimers{}
					if o.Protocol.Rip.Timers.Misc != nil {
						entry.Misc["ProtocolRipTimers"] = o.Protocol.Rip.Timers.Misc
					}
					if o.Protocol.Rip.Timers.IntervalSeconds != nil {
						nestedProtocol.Rip.Timers.IntervalSeconds = o.Protocol.Rip.Timers.IntervalSeconds
					}
					if o.Protocol.Rip.Timers.UpdateIntervals != nil {
						nestedProtocol.Rip.Timers.UpdateIntervals = o.Protocol.Rip.Timers.UpdateIntervals
					}
					if o.Protocol.Rip.Timers.DeleteIntervals != nil {
						nestedProtocol.Rip.Timers.DeleteIntervals = o.Protocol.Rip.Timers.DeleteIntervals
					}
					if o.Protocol.Rip.Timers.ExpireIntervals != nil {
						nestedProtocol.Rip.Timers.ExpireIntervals = o.Protocol.Rip.Timers.ExpireIntervals
					}
				}
				if o.Protocol.Rip.AllowRedistDefaultRoute != nil {
					nestedProtocol.Rip.AllowRedistDefaultRoute = util.AsBool(o.Protocol.Rip.AllowRedistDefaultRoute, nil)
				}
				if o.Protocol.Rip.AuthProfile != nil {
					nestedProtocol.Rip.AuthProfile = []ProtocolRipAuthProfile{}
					for _, oProtocolRipAuthProfile := range o.Protocol.Rip.AuthProfile {
						nestedProtocolRipAuthProfile := ProtocolRipAuthProfile{}
						if oProtocolRipAuthProfile.Misc != nil {
							entry.Misc["ProtocolRipAuthProfile"] = oProtocolRipAuthProfile.Misc
						}
						if oProtocolRipAuthProfile.Name != "" {
							nestedProtocolRipAuthProfile.Name = oProtocolRipAuthProfile.Name
						}
						if oProtocolRipAuthProfile.Password != nil {
							nestedProtocolRipAuthProfile.Password = oProtocolRipAuthProfile.Password
						}
						if oProtocolRipAuthProfile.Md5 != nil {
							nestedProtocolRipAuthProfile.Md5 = []ProtocolRipAuthProfileMd5{}
							for _, oProtocolRipAuthProfileMd5 := range oProtocolRipAuthProfile.Md5 {
								nestedProtocolRipAuthProfileMd5 := ProtocolRipAuthProfileMd5{}
								if oProtocolRipAuthProfileMd5.Misc != nil {
									entry.Misc["ProtocolRipAuthProfileMd5"] = oProtocolRipAuthProfileMd5.Misc
								}
								if oProtocolRipAuthProfileMd5.Preferred != nil {
									nestedProtocolRipAuthProfileMd5.Preferred = util.AsBool(oProtocolRipAuthProfileMd5.Preferred, nil)
								}
								if oProtocolRipAuthProfileMd5.Name != "" {
									nestedProtocolRipAuthProfileMd5.Name = oProtocolRipAuthProfileMd5.Name
								}
								if oProtocolRipAuthProfileMd5.Key != nil {
									nestedProtocolRipAuthProfileMd5.Key = oProtocolRipAuthProfileMd5.Key
								}
								nestedProtocolRipAuthProfile.Md5 = append(nestedProtocolRipAuthProfile.Md5, nestedProtocolRipAuthProfileMd5)
							}
						}
						nestedProtocol.Rip.AuthProfile = append(nestedProtocol.Rip.AuthProfile, nestedProtocolRipAuthProfile)
					}
				}
				if o.Protocol.Rip.Enable != nil {
					nestedProtocol.Rip.Enable = util.AsBool(o.Protocol.Rip.Enable, nil)
				}
			}
			if o.Protocol.Bgp != nil {
				nestedProtocol.Bgp = &ProtocolBgp{}
				if o.Protocol.Bgp.Misc != nil {
					entry.Misc["ProtocolBgp"] = o.Protocol.Bgp.Misc
				}
				if o.Protocol.Bgp.PeerGroup != nil {
					nestedProtocol.Bgp.PeerGroup = []ProtocolBgpPeerGroup{}
					for _, oProtocolBgpPeerGroup := range o.Protocol.Bgp.PeerGroup {
						nestedProtocolBgpPeerGroup := ProtocolBgpPeerGroup{}
						if oProtocolBgpPeerGroup.Misc != nil {
							entry.Misc["ProtocolBgpPeerGroup"] = oProtocolBgpPeerGroup.Misc
						}
						if oProtocolBgpPeerGroup.Enable != nil {
							nestedProtocolBgpPeerGroup.Enable = util.AsBool(oProtocolBgpPeerGroup.Enable, nil)
						}
						if oProtocolBgpPeerGroup.AggregatedConfedAsPath != nil {
							nestedProtocolBgpPeerGroup.AggregatedConfedAsPath = util.AsBool(oProtocolBgpPeerGroup.AggregatedConfedAsPath, nil)
						}
						if oProtocolBgpPeerGroup.SoftResetWithStoredInfo != nil {
							nestedProtocolBgpPeerGroup.SoftResetWithStoredInfo = util.AsBool(oProtocolBgpPeerGroup.SoftResetWithStoredInfo, nil)
						}
						if oProtocolBgpPeerGroup.Type != nil {
							nestedProtocolBgpPeerGroup.Type = &ProtocolBgpPeerGroupType{}
							if oProtocolBgpPeerGroup.Type.Misc != nil {
								entry.Misc["ProtocolBgpPeerGroupType"] = oProtocolBgpPeerGroup.Type.Misc
							}
							if oProtocolBgpPeerGroup.Type.Ebgp != nil {
								nestedProtocolBgpPeerGroup.Type.Ebgp = &ProtocolBgpPeerGroupTypeEbgp{}
								if oProtocolBgpPeerGroup.Type.Ebgp.Misc != nil {
									entry.Misc["ProtocolBgpPeerGroupTypeEbgp"] = oProtocolBgpPeerGroup.Type.Ebgp.Misc
								}
								if oProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop != nil {
									nestedProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop = oProtocolBgpPeerGroup.Type.Ebgp.ImportNexthop
								}
								if oProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop != nil {
									nestedProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop = oProtocolBgpPeerGroup.Type.Ebgp.ExportNexthop
								}
								if oProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs != nil {
									nestedProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs = util.AsBool(oProtocolBgpPeerGroup.Type.Ebgp.RemovePrivateAs, nil)
								}
							}
							if oProtocolBgpPeerGroup.Type.Ibgp != nil {
								nestedProtocolBgpPeerGroup.Type.Ibgp = &ProtocolBgpPeerGroupTypeIbgp{}
								if oProtocolBgpPeerGroup.Type.Ibgp.Misc != nil {
									entry.Misc["ProtocolBgpPeerGroupTypeIbgp"] = oProtocolBgpPeerGroup.Type.Ibgp.Misc
								}
								if oProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop != nil {
									nestedProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop = oProtocolBgpPeerGroup.Type.Ibgp.ExportNexthop
								}
							}
							if oProtocolBgpPeerGroup.Type.EbgpConfed != nil {
								nestedProtocolBgpPeerGroup.Type.EbgpConfed = &ProtocolBgpPeerGroupTypeEbgpConfed{}
								if oProtocolBgpPeerGroup.Type.EbgpConfed.Misc != nil {
									entry.Misc["ProtocolBgpPeerGroupTypeEbgpConfed"] = oProtocolBgpPeerGroup.Type.EbgpConfed.Misc
								}
								if oProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop != nil {
									nestedProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop = oProtocolBgpPeerGroup.Type.EbgpConfed.ExportNexthop
								}
							}
							if oProtocolBgpPeerGroup.Type.IbgpConfed != nil {
								nestedProtocolBgpPeerGroup.Type.IbgpConfed = &ProtocolBgpPeerGroupTypeIbgpConfed{}
								if oProtocolBgpPeerGroup.Type.IbgpConfed.Misc != nil {
									entry.Misc["ProtocolBgpPeerGroupTypeIbgpConfed"] = oProtocolBgpPeerGroup.Type.IbgpConfed.Misc
								}
								if oProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop != nil {
									nestedProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop = oProtocolBgpPeerGroup.Type.IbgpConfed.ExportNexthop
								}
							}
						}
						if oProtocolBgpPeerGroup.Peer != nil {
							nestedProtocolBgpPeerGroup.Peer = []ProtocolBgpPeerGroupPeer{}
							for _, oProtocolBgpPeerGroupPeer := range oProtocolBgpPeerGroup.Peer {
								nestedProtocolBgpPeerGroupPeer := ProtocolBgpPeerGroupPeer{}
								if oProtocolBgpPeerGroupPeer.Misc != nil {
									entry.Misc["ProtocolBgpPeerGroupPeer"] = oProtocolBgpPeerGroupPeer.Misc
								}
								if oProtocolBgpPeerGroupPeer.Bfd != nil {
									nestedProtocolBgpPeerGroupPeer.Bfd = &ProtocolBgpPeerGroupPeerBfd{}
									if oProtocolBgpPeerGroupPeer.Bfd.Misc != nil {
										entry.Misc["ProtocolBgpPeerGroupPeerBfd"] = oProtocolBgpPeerGroupPeer.Bfd.Misc
									}
									if oProtocolBgpPeerGroupPeer.Bfd.Profile != nil {
										nestedProtocolBgpPeerGroupPeer.Bfd.Profile = oProtocolBgpPeerGroupPeer.Bfd.Profile
									}
								}
								if oProtocolBgpPeerGroupPeer.EnableMpBgp != nil {
									nestedProtocolBgpPeerGroupPeer.EnableMpBgp = util.AsBool(oProtocolBgpPeerGroupPeer.EnableMpBgp, nil)
								}
								if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier != nil {
									nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier = &ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier{}
									if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Misc != nil {
										entry.Misc["ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier"] = oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Misc
									}
									if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast != nil {
										nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast = util.AsBool(oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Unicast, nil)
									}
									if oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast != nil {
										nestedProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast = util.AsBool(oProtocolBgpPeerGroupPeer.SubsequentAddressFamilyIdentifier.Multicast, nil)
									}
								}
								if oProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection != nil {
									nestedProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection = util.AsBool(oProtocolBgpPeerGroupPeer.EnableSenderSideLoopDetection, nil)
								}
								if oProtocolBgpPeerGroupPeer.PeeringType != nil {
									nestedProtocolBgpPeerGroupPeer.PeeringType = oProtocolBgpPeerGroupPeer.PeeringType
								}
								if oProtocolBgpPeerGroupPeer.LocalAddress != nil {
									nestedProtocolBgpPeerGroupPeer.LocalAddress = &ProtocolBgpPeerGroupPeerLocalAddress{}
									if oProtocolBgpPeerGroupPeer.LocalAddress.Misc != nil {
										entry.Misc["ProtocolBgpPeerGroupPeerLocalAddress"] = oProtocolBgpPeerGroupPeer.LocalAddress.Misc
									}
									if oProtocolBgpPeerGroupPeer.LocalAddress.Ip != nil {
										nestedProtocolBgpPeerGroupPeer.LocalAddress.Ip = oProtocolBgpPeerGroupPeer.LocalAddress.Ip
									}
									if oProtocolBgpPeerGroupPeer.LocalAddress.Interface != nil {
										nestedProtocolBgpPeerGroupPeer.LocalAddress.Interface = oProtocolBgpPeerGroupPeer.LocalAddress.Interface
									}
								}
								if oProtocolBgpPeerGroupPeer.ConnectionOptions != nil {
									nestedProtocolBgpPeerGroupPeer.ConnectionOptions = &ProtocolBgpPeerGroupPeerConnectionOptions{}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.Misc != nil {
										entry.Misc["ProtocolBgpPeerGroupPeerConnectionOptions"] = oProtocolBgpPeerGroupPeer.ConnectionOptions.Misc
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection = &ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection{}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Misc != nil {
											entry.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection"] = oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Misc
										}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort != nil {
											nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort = oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.RemotePort
										}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow != nil {
											nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow = util.AsBool(oProtocolBgpPeerGroupPeer.ConnectionOptions.IncomingBgpConnection.Allow, nil)
										}
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection = &ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection{}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Misc != nil {
											entry.Misc["ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection"] = oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Misc
										}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort != nil {
											nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort = oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.LocalPort
										}
										if oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow != nil {
											nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow = util.AsBool(oProtocolBgpPeerGroupPeer.ConnectionOptions.OutgoingBgpConnection.Allow, nil)
										}
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval = oProtocolBgpPeerGroupPeer.ConnectionOptions.MinRouteAdvInterval
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.IdleHoldTime
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop = oProtocolBgpPeerGroupPeer.ConnectionOptions.Multihop
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.OpenDelayTime
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime = oProtocolBgpPeerGroupPeer.ConnectionOptions.HoldTime
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication = oProtocolBgpPeerGroupPeer.ConnectionOptions.Authentication
									}
									if oProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval != nil {
										nestedProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval = oProtocolBgpPeerGroupPeer.ConnectionOptions.KeepAliveInterval
									}
								}
								if oProtocolBgpPeerGroupPeer.Enable != nil {
									nestedProtocolBgpPeerGroupPeer.Enable = util.AsBool(oProtocolBgpPeerGroupPeer.Enable, nil)
								}
								if oProtocolBgpPeerGroupPeer.PeerAs != nil {
									nestedProtocolBgpPeerGroupPeer.PeerAs = oProtocolBgpPeerGroupPeer.PeerAs
								}
								if oProtocolBgpPeerGroupPeer.ReflectorClient != nil {
									nestedProtocolBgpPeerGroupPeer.ReflectorClient = oProtocolBgpPeerGroupPeer.ReflectorClient
								}
								if oProtocolBgpPeerGroupPeer.PeerAddress != nil {
									nestedProtocolBgpPeerGroupPeer.PeerAddress = &ProtocolBgpPeerGroupPeerPeerAddress{}
									if oProtocolBgpPeerGroupPeer.PeerAddress.Misc != nil {
										entry.Misc["ProtocolBgpPeerGroupPeerPeerAddress"] = oProtocolBgpPeerGroupPeer.PeerAddress.Misc
									}
									if oProtocolBgpPeerGroupPeer.PeerAddress.Ip != nil {
										nestedProtocolBgpPeerGroupPeer.PeerAddress.Ip = oProtocolBgpPeerGroupPeer.PeerAddress.Ip
									}
									if oProtocolBgpPeerGroupPeer.PeerAddress.Fqdn != nil {
										nestedProtocolBgpPeerGroupPeer.PeerAddress.Fqdn = oProtocolBgpPeerGroupPeer.PeerAddress.Fqdn
									}
								}
								if oProtocolBgpPeerGroupPeer.Name != "" {
									nestedProtocolBgpPeerGroupPeer.Name = oProtocolBgpPeerGroupPeer.Name
								}
								if oProtocolBgpPeerGroupPeer.AddressFamilyIdentifier != nil {
									nestedProtocolBgpPeerGroupPeer.AddressFamilyIdentifier = oProtocolBgpPeerGroupPeer.AddressFamilyIdentifier
								}
								if oProtocolBgpPeerGroupPeer.MaxPrefixes != nil {
									nestedProtocolBgpPeerGroupPeer.MaxPrefixes = oProtocolBgpPeerGroupPeer.MaxPrefixes
								}
								nestedProtocolBgpPeerGroup.Peer = append(nestedProtocolBgpPeerGroup.Peer, nestedProtocolBgpPeerGroupPeer)
							}
						}
						if oProtocolBgpPeerGroup.Name != "" {
							nestedProtocolBgpPeerGroup.Name = oProtocolBgpPeerGroup.Name
						}
						nestedProtocol.Bgp.PeerGroup = append(nestedProtocol.Bgp.PeerGroup, nestedProtocolBgpPeerGroup)
					}
				}
				if o.Protocol.Bgp.RejectDefaultRoute != nil {
					nestedProtocol.Bgp.RejectDefaultRoute = util.AsBool(o.Protocol.Bgp.RejectDefaultRoute, nil)
				}
				if o.Protocol.Bgp.DampeningProfile != nil {
					nestedProtocol.Bgp.DampeningProfile = []ProtocolBgpDampeningProfile{}
					for _, oProtocolBgpDampeningProfile := range o.Protocol.Bgp.DampeningProfile {
						nestedProtocolBgpDampeningProfile := ProtocolBgpDampeningProfile{}
						if oProtocolBgpDampeningProfile.Misc != nil {
							entry.Misc["ProtocolBgpDampeningProfile"] = oProtocolBgpDampeningProfile.Misc
						}
						if oProtocolBgpDampeningProfile.Cutoff != nil {
							nestedProtocolBgpDampeningProfile.Cutoff = oProtocolBgpDampeningProfile.Cutoff
						}
						if oProtocolBgpDampeningProfile.Reuse != nil {
							nestedProtocolBgpDampeningProfile.Reuse = oProtocolBgpDampeningProfile.Reuse
						}
						if oProtocolBgpDampeningProfile.MaxHoldTime != nil {
							nestedProtocolBgpDampeningProfile.MaxHoldTime = oProtocolBgpDampeningProfile.MaxHoldTime
						}
						if oProtocolBgpDampeningProfile.DecayHalfLifeReachable != nil {
							nestedProtocolBgpDampeningProfile.DecayHalfLifeReachable = oProtocolBgpDampeningProfile.DecayHalfLifeReachable
						}
						if oProtocolBgpDampeningProfile.DecayHalfLifeUnreachable != nil {
							nestedProtocolBgpDampeningProfile.DecayHalfLifeUnreachable = oProtocolBgpDampeningProfile.DecayHalfLifeUnreachable
						}
						if oProtocolBgpDampeningProfile.Name != "" {
							nestedProtocolBgpDampeningProfile.Name = oProtocolBgpDampeningProfile.Name
						}
						if oProtocolBgpDampeningProfile.Enable != nil {
							nestedProtocolBgpDampeningProfile.Enable = util.AsBool(oProtocolBgpDampeningProfile.Enable, nil)
						}
						nestedProtocol.Bgp.DampeningProfile = append(nestedProtocol.Bgp.DampeningProfile, nestedProtocolBgpDampeningProfile)
					}
				}
				if o.Protocol.Bgp.EcmpMultiAs != nil {
					nestedProtocol.Bgp.EcmpMultiAs = util.AsBool(o.Protocol.Bgp.EcmpMultiAs, nil)
				}
				if o.Protocol.Bgp.GlobalBfd != nil {
					nestedProtocol.Bgp.GlobalBfd = &ProtocolBgpGlobalBfd{}
					if o.Protocol.Bgp.GlobalBfd.Misc != nil {
						entry.Misc["ProtocolBgpGlobalBfd"] = o.Protocol.Bgp.GlobalBfd.Misc
					}
					if o.Protocol.Bgp.GlobalBfd.Profile != nil {
						nestedProtocol.Bgp.GlobalBfd.Profile = o.Protocol.Bgp.GlobalBfd.Profile
					}
				}
				if o.Protocol.Bgp.LocalAs != nil {
					nestedProtocol.Bgp.LocalAs = o.Protocol.Bgp.LocalAs
				}
				if o.Protocol.Bgp.RedistRules != nil {
					nestedProtocol.Bgp.RedistRules = []ProtocolBgpRedistRules{}
					for _, oProtocolBgpRedistRules := range o.Protocol.Bgp.RedistRules {
						nestedProtocolBgpRedistRules := ProtocolBgpRedistRules{}
						if oProtocolBgpRedistRules.Misc != nil {
							entry.Misc["ProtocolBgpRedistRules"] = oProtocolBgpRedistRules.Misc
						}
						if oProtocolBgpRedistRules.RouteTable != nil {
							nestedProtocolBgpRedistRules.RouteTable = oProtocolBgpRedistRules.RouteTable
						}
						if oProtocolBgpRedistRules.SetOrigin != nil {
							nestedProtocolBgpRedistRules.SetOrigin = oProtocolBgpRedistRules.SetOrigin
						}
						if oProtocolBgpRedistRules.SetMed != nil {
							nestedProtocolBgpRedistRules.SetMed = oProtocolBgpRedistRules.SetMed
						}
						if oProtocolBgpRedistRules.SetAsPathLimit != nil {
							nestedProtocolBgpRedistRules.SetAsPathLimit = oProtocolBgpRedistRules.SetAsPathLimit
						}
						if oProtocolBgpRedistRules.Metric != nil {
							nestedProtocolBgpRedistRules.Metric = oProtocolBgpRedistRules.Metric
						}
						if oProtocolBgpRedistRules.SetCommunity != nil {
							nestedProtocolBgpRedistRules.SetCommunity = util.MemToStr(oProtocolBgpRedistRules.SetCommunity)
						}
						if oProtocolBgpRedistRules.AddressFamilyIdentifier != nil {
							nestedProtocolBgpRedistRules.AddressFamilyIdentifier = oProtocolBgpRedistRules.AddressFamilyIdentifier
						}
						if oProtocolBgpRedistRules.Enable != nil {
							nestedProtocolBgpRedistRules.Enable = util.AsBool(oProtocolBgpRedistRules.Enable, nil)
						}
						if oProtocolBgpRedistRules.SetLocalPreference != nil {
							nestedProtocolBgpRedistRules.SetLocalPreference = oProtocolBgpRedistRules.SetLocalPreference
						}
						if oProtocolBgpRedistRules.SetExtendedCommunity != nil {
							nestedProtocolBgpRedistRules.SetExtendedCommunity = util.MemToStr(oProtocolBgpRedistRules.SetExtendedCommunity)
						}
						if oProtocolBgpRedistRules.Name != "" {
							nestedProtocolBgpRedistRules.Name = oProtocolBgpRedistRules.Name
						}
						nestedProtocol.Bgp.RedistRules = append(nestedProtocol.Bgp.RedistRules, nestedProtocolBgpRedistRules)
					}
				}
				if o.Protocol.Bgp.RouterId != nil {
					nestedProtocol.Bgp.RouterId = o.Protocol.Bgp.RouterId
				}
				if o.Protocol.Bgp.AllowRedistDefaultRoute != nil {
					nestedProtocol.Bgp.AllowRedistDefaultRoute = util.AsBool(o.Protocol.Bgp.AllowRedistDefaultRoute, nil)
				}
				if o.Protocol.Bgp.EnforceFirstAs != nil {
					nestedProtocol.Bgp.EnforceFirstAs = util.AsBool(o.Protocol.Bgp.EnforceFirstAs, nil)
				}
				if o.Protocol.Bgp.RoutingOptions != nil {
					nestedProtocol.Bgp.RoutingOptions = &ProtocolBgpRoutingOptions{}
					if o.Protocol.Bgp.RoutingOptions.Misc != nil {
						entry.Misc["ProtocolBgpRoutingOptions"] = o.Protocol.Bgp.RoutingOptions.Misc
					}
					if o.Protocol.Bgp.RoutingOptions.ConfederationMemberAs != nil {
						nestedProtocol.Bgp.RoutingOptions.ConfederationMemberAs = o.Protocol.Bgp.RoutingOptions.ConfederationMemberAs
					}
					if o.Protocol.Bgp.RoutingOptions.DefaultLocalPreference != nil {
						nestedProtocol.Bgp.RoutingOptions.DefaultLocalPreference = o.Protocol.Bgp.RoutingOptions.DefaultLocalPreference
					}
					if o.Protocol.Bgp.RoutingOptions.GracefulRestart != nil {
						nestedProtocol.Bgp.RoutingOptions.GracefulRestart = &ProtocolBgpRoutingOptionsGracefulRestart{}
						if o.Protocol.Bgp.RoutingOptions.GracefulRestart.Misc != nil {
							entry.Misc["ProtocolBgpRoutingOptionsGracefulRestart"] = o.Protocol.Bgp.RoutingOptions.GracefulRestart.Misc
						}
						if o.Protocol.Bgp.RoutingOptions.GracefulRestart.Enable != nil {
							nestedProtocol.Bgp.RoutingOptions.GracefulRestart.Enable = util.AsBool(o.Protocol.Bgp.RoutingOptions.GracefulRestart.Enable, nil)
						}
						if o.Protocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime != nil {
							nestedProtocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.LocalRestartTime
						}
						if o.Protocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime != nil {
							nestedProtocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.MaxPeerRestartTime
						}
						if o.Protocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime != nil {
							nestedProtocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime = o.Protocol.Bgp.RoutingOptions.GracefulRestart.StaleRouteTime
						}
					}
					if o.Protocol.Bgp.RoutingOptions.Med != nil {
						nestedProtocol.Bgp.RoutingOptions.Med = &ProtocolBgpRoutingOptionsMed{}
						if o.Protocol.Bgp.RoutingOptions.Med.Misc != nil {
							entry.Misc["ProtocolBgpRoutingOptionsMed"] = o.Protocol.Bgp.RoutingOptions.Med.Misc
						}
						if o.Protocol.Bgp.RoutingOptions.Med.AlwaysCompareMed != nil {
							nestedProtocol.Bgp.RoutingOptions.Med.AlwaysCompareMed = util.AsBool(o.Protocol.Bgp.RoutingOptions.Med.AlwaysCompareMed, nil)
						}
						if o.Protocol.Bgp.RoutingOptions.Med.DeterministicMedComparison != nil {
							nestedProtocol.Bgp.RoutingOptions.Med.DeterministicMedComparison = util.AsBool(o.Protocol.Bgp.RoutingOptions.Med.DeterministicMedComparison, nil)
						}
					}
					if o.Protocol.Bgp.RoutingOptions.ReflectorClusterId != nil {
						nestedProtocol.Bgp.RoutingOptions.ReflectorClusterId = o.Protocol.Bgp.RoutingOptions.ReflectorClusterId
					}
					if o.Protocol.Bgp.RoutingOptions.Aggregate != nil {
						nestedProtocol.Bgp.RoutingOptions.Aggregate = &ProtocolBgpRoutingOptionsAggregate{}
						if o.Protocol.Bgp.RoutingOptions.Aggregate.Misc != nil {
							entry.Misc["ProtocolBgpRoutingOptionsAggregate"] = o.Protocol.Bgp.RoutingOptions.Aggregate.Misc
						}
						if o.Protocol.Bgp.RoutingOptions.Aggregate.AggregateMed != nil {
							nestedProtocol.Bgp.RoutingOptions.Aggregate.AggregateMed = util.AsBool(o.Protocol.Bgp.RoutingOptions.Aggregate.AggregateMed, nil)
						}
					}
					if o.Protocol.Bgp.RoutingOptions.AsFormat != nil {
						nestedProtocol.Bgp.RoutingOptions.AsFormat = o.Protocol.Bgp.RoutingOptions.AsFormat
					}
				}
				if o.Protocol.Bgp.InstallRoute != nil {
					nestedProtocol.Bgp.InstallRoute = util.AsBool(o.Protocol.Bgp.InstallRoute, nil)
				}
				if o.Protocol.Bgp.Policy != nil {
					nestedProtocol.Bgp.Policy = &ProtocolBgpPolicy{}
					if o.Protocol.Bgp.Policy.Misc != nil {
						entry.Misc["ProtocolBgpPolicy"] = o.Protocol.Bgp.Policy.Misc
					}
					if o.Protocol.Bgp.Policy.Aggregation != nil {
						nestedProtocol.Bgp.Policy.Aggregation = &ProtocolBgpPolicyAggregation{}
						if o.Protocol.Bgp.Policy.Aggregation.Misc != nil {
							entry.Misc["ProtocolBgpPolicyAggregation"] = o.Protocol.Bgp.Policy.Aggregation.Misc
						}
						if o.Protocol.Bgp.Policy.Aggregation.Address != nil {
							nestedProtocol.Bgp.Policy.Aggregation.Address = []ProtocolBgpPolicyAggregationAddress{}
							for _, oProtocolBgpPolicyAggregationAddress := range o.Protocol.Bgp.Policy.Aggregation.Address {
								nestedProtocolBgpPolicyAggregationAddress := ProtocolBgpPolicyAggregationAddress{}
								if oProtocolBgpPolicyAggregationAddress.Misc != nil {
									entry.Misc["ProtocolBgpPolicyAggregationAddress"] = oProtocolBgpPolicyAggregationAddress.Misc
								}
								if oProtocolBgpPolicyAggregationAddress.Enable != nil {
									nestedProtocolBgpPolicyAggregationAddress.Enable = util.AsBool(oProtocolBgpPolicyAggregationAddress.Enable, nil)
								}
								if oProtocolBgpPolicyAggregationAddress.Summary != nil {
									nestedProtocolBgpPolicyAggregationAddress.Summary = util.AsBool(oProtocolBgpPolicyAggregationAddress.Summary, nil)
								}
								if oProtocolBgpPolicyAggregationAddress.AsSet != nil {
									nestedProtocolBgpPolicyAggregationAddress.AsSet = util.AsBool(oProtocolBgpPolicyAggregationAddress.AsSet, nil)
								}
								if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes != nil {
									nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes{}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Misc != nil {
										entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Misc
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Nexthop
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Origin
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath{}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Misc != nil {
											entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Misc
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Remove.Misc
											}
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.Prepend
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPath.None.Misc
											}
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity{}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Misc != nil {
											entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Misc
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.None.Misc
											}
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveAll.Misc
											}
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.RemoveRegex
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append = util.MemToStr(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Append)
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite = util.MemToStr(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Community.Overwrite)
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.LocalPreference
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Weight
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity{}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Misc != nil {
											entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Misc
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveRegex
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append = util.MemToStr(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Append)
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite = util.MemToStr(oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.Overwrite)
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.None.Misc
											}
										}
										if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll != nil {
											nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll{}
											if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll"] = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.ExtendedCommunity.RemoveAll.Misc
											}
										}
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.Med
									}
									if oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit != nil {
										nestedProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit = oProtocolBgpPolicyAggregationAddress.AggregateRouteAttributes.AsPathLimit
									}
								}
								if oProtocolBgpPolicyAggregationAddress.SuppressFilters != nil {
									nestedProtocolBgpPolicyAggregationAddress.SuppressFilters = []ProtocolBgpPolicyAggregationAddressSuppressFilters{}
									for _, oProtocolBgpPolicyAggregationAddressSuppressFilters := range oProtocolBgpPolicyAggregationAddress.SuppressFilters {
										nestedProtocolBgpPolicyAggregationAddressSuppressFilters := ProtocolBgpPolicyAggregationAddressSuppressFilters{}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Misc != nil {
											entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFilters"] = oProtocolBgpPolicyAggregationAddressSuppressFilters.Misc
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Enable != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Enable = util.AsBool(oProtocolBgpPolicyAggregationAddressSuppressFilters.Enable, nil)
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match != nil {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch{}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch"] = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Misc
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity{}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity"] = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.ExtendedCommunity.Regex
												}
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.RouteTable
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Med
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix = []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix{}
												for _, oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix := range oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix {
													nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix := ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix{}
													if oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Misc != nil {
														entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix"] = oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Misc
													}
													if oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact != nil {
														nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact = util.AsBool(oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Exact, nil)
													}
													if oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name != "" {
														nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix.Name
													}
													nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AddressPrefix, nestedProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix)
												}
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Nexthop)
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.FromPeer)
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath{}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath"] = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Misc
												}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.AsPath.Regex
												}
											}
											if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community != nil {
												nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community = &ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity{}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity"] = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Misc
												}
												if oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex = oProtocolBgpPolicyAggregationAddressSuppressFilters.Match.Community.Regex
												}
											}
										}
										if oProtocolBgpPolicyAggregationAddressSuppressFilters.Name != "" {
											nestedProtocolBgpPolicyAggregationAddressSuppressFilters.Name = oProtocolBgpPolicyAggregationAddressSuppressFilters.Name
										}
										nestedProtocolBgpPolicyAggregationAddress.SuppressFilters = append(nestedProtocolBgpPolicyAggregationAddress.SuppressFilters, nestedProtocolBgpPolicyAggregationAddressSuppressFilters)
									}
								}
								if oProtocolBgpPolicyAggregationAddress.AdvertiseFilters != nil {
									nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters = []ProtocolBgpPolicyAggregationAddressAdvertiseFilters{}
									for _, oProtocolBgpPolicyAggregationAddressAdvertiseFilters := range oProtocolBgpPolicyAggregationAddress.AdvertiseFilters {
										nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters := ProtocolBgpPolicyAggregationAddressAdvertiseFilters{}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Misc != nil {
											entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFilters"] = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Misc
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable = util.AsBool(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Enable, nil)
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match != nil {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch{}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Misc != nil {
												entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch"] = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Misc
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath{}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath"] = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Misc
												}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AsPath.Regex
												}
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity{}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity"] = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Misc
												}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Community.Regex
												}
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity{}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity"] = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex != nil {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.ExtendedCommunity.Regex
												}
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.RouteTable
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Med
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix = []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix{}
												for _, oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix := range oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix {
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix := ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix{}
													if oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Misc != nil {
														entry.Misc["ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix"] = oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Misc
													}
													if oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact != nil {
														nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact = util.AsBool(oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Exact, nil)
													}
													if oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name != "" {
														nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix.Name
													}
													nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.AddressPrefix, nestedProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix)
												}
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.Nexthop)
											}
											if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer != nil {
												nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Match.FromPeer)
											}
										}
										if oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name != "" {
											nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name = oProtocolBgpPolicyAggregationAddressAdvertiseFilters.Name
										}
										nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters = append(nestedProtocolBgpPolicyAggregationAddress.AdvertiseFilters, nestedProtocolBgpPolicyAggregationAddressAdvertiseFilters)
									}
								}
								if oProtocolBgpPolicyAggregationAddress.Name != "" {
									nestedProtocolBgpPolicyAggregationAddress.Name = oProtocolBgpPolicyAggregationAddress.Name
								}
								if oProtocolBgpPolicyAggregationAddress.Prefix != nil {
									nestedProtocolBgpPolicyAggregationAddress.Prefix = oProtocolBgpPolicyAggregationAddress.Prefix
								}
								nestedProtocol.Bgp.Policy.Aggregation.Address = append(nestedProtocol.Bgp.Policy.Aggregation.Address, nestedProtocolBgpPolicyAggregationAddress)
							}
						}
					}
					if o.Protocol.Bgp.Policy.ConditionalAdvertisement != nil {
						nestedProtocol.Bgp.Policy.ConditionalAdvertisement = &ProtocolBgpPolicyConditionalAdvertisement{}
						if o.Protocol.Bgp.Policy.ConditionalAdvertisement.Misc != nil {
							entry.Misc["ProtocolBgpPolicyConditionalAdvertisement"] = o.Protocol.Bgp.Policy.ConditionalAdvertisement.Misc
						}
						if o.Protocol.Bgp.Policy.ConditionalAdvertisement.Policy != nil {
							nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy = []ProtocolBgpPolicyConditionalAdvertisementPolicy{}
							for _, oProtocolBgpPolicyConditionalAdvertisementPolicy := range o.Protocol.Bgp.Policy.ConditionalAdvertisement.Policy {
								nestedProtocolBgpPolicyConditionalAdvertisementPolicy := ProtocolBgpPolicyConditionalAdvertisementPolicy{}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.Misc != nil {
									entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicy"] = oProtocolBgpPolicyConditionalAdvertisementPolicy.Misc
								}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.Enable != nil {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.Enable = util.AsBool(oProtocolBgpPolicyConditionalAdvertisementPolicy.Enable, nil)
								}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy != nil {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy = util.MemToStr(oProtocolBgpPolicyConditionalAdvertisementPolicy.UsedBy)
								}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters != nil {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters = []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters{}
									for _, oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters := range oProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters := ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters{}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Misc != nil {
											entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Misc
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable = util.AsBool(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Enable, nil)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch{}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Misc != nil {
												entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Misc
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Nexthop)
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.FromPeer)
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AsPath.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Community.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.ExtendedCommunity.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.RouteTable
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.Med
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix = []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix{}
												for _, oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix := range oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix := ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix{}
													if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Misc != nil {
														entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix"] = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Misc
													}
													if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name != "" {
														nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix.Name
													}
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Match.AddressPrefix, nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix)
												}
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name != "" {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters.Name
										}
										nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicy.NonExistFilters, nestedProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters)
									}
								}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters != nil {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters = []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters{}
									for _, oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters := range oProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters {
										nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters := ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters{}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Misc != nil {
											entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Misc
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable = util.AsBool(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Enable, nil)
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match != nil {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch{}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Misc != nil {
												entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Misc
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AsPath.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Community.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity = &ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity{}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex != nil {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.ExtendedCommunity.Regex
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.RouteTable
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Med
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix = []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix{}
												for _, oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix := range oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix {
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix := ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix{}
													if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Misc != nil {
														entry.Misc["ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix"] = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Misc
													}
													if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name != "" {
														nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix.Name
													}
													nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.AddressPrefix, nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix)
												}
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.Nexthop)
											}
											if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer != nil {
												nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Match.FromPeer)
											}
										}
										if oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name != "" {
											nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name = oProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters.Name
										}
										nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters = append(nestedProtocolBgpPolicyConditionalAdvertisementPolicy.AdvertiseFilters, nestedProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters)
									}
								}
								if oProtocolBgpPolicyConditionalAdvertisementPolicy.Name != "" {
									nestedProtocolBgpPolicyConditionalAdvertisementPolicy.Name = oProtocolBgpPolicyConditionalAdvertisementPolicy.Name
								}
								nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy = append(nestedProtocol.Bgp.Policy.ConditionalAdvertisement.Policy, nestedProtocolBgpPolicyConditionalAdvertisementPolicy)
							}
						}
					}
					if o.Protocol.Bgp.Policy.Export != nil {
						nestedProtocol.Bgp.Policy.Export = &ProtocolBgpPolicyExport{}
						if o.Protocol.Bgp.Policy.Export.Misc != nil {
							entry.Misc["ProtocolBgpPolicyExport"] = o.Protocol.Bgp.Policy.Export.Misc
						}
						if o.Protocol.Bgp.Policy.Export.Rules != nil {
							nestedProtocol.Bgp.Policy.Export.Rules = []ProtocolBgpPolicyExportRules{}
							for _, oProtocolBgpPolicyExportRules := range o.Protocol.Bgp.Policy.Export.Rules {
								nestedProtocolBgpPolicyExportRules := ProtocolBgpPolicyExportRules{}
								if oProtocolBgpPolicyExportRules.Misc != nil {
									entry.Misc["ProtocolBgpPolicyExportRules"] = oProtocolBgpPolicyExportRules.Misc
								}
								if oProtocolBgpPolicyExportRules.UsedBy != nil {
									nestedProtocolBgpPolicyExportRules.UsedBy = util.MemToStr(oProtocolBgpPolicyExportRules.UsedBy)
								}
								if oProtocolBgpPolicyExportRules.Match != nil {
									nestedProtocolBgpPolicyExportRules.Match = &ProtocolBgpPolicyExportRulesMatch{}
									if oProtocolBgpPolicyExportRules.Match.Misc != nil {
										entry.Misc["ProtocolBgpPolicyExportRulesMatch"] = oProtocolBgpPolicyExportRules.Match.Misc
									}
									if oProtocolBgpPolicyExportRules.Match.RouteTable != nil {
										nestedProtocolBgpPolicyExportRules.Match.RouteTable = oProtocolBgpPolicyExportRules.Match.RouteTable
									}
									if oProtocolBgpPolicyExportRules.Match.Med != nil {
										nestedProtocolBgpPolicyExportRules.Match.Med = oProtocolBgpPolicyExportRules.Match.Med
									}
									if oProtocolBgpPolicyExportRules.Match.AddressPrefix != nil {
										nestedProtocolBgpPolicyExportRules.Match.AddressPrefix = []ProtocolBgpPolicyExportRulesMatchAddressPrefix{}
										for _, oProtocolBgpPolicyExportRulesMatchAddressPrefix := range oProtocolBgpPolicyExportRules.Match.AddressPrefix {
											nestedProtocolBgpPolicyExportRulesMatchAddressPrefix := ProtocolBgpPolicyExportRulesMatchAddressPrefix{}
											if oProtocolBgpPolicyExportRulesMatchAddressPrefix.Misc != nil {
												entry.Misc["ProtocolBgpPolicyExportRulesMatchAddressPrefix"] = oProtocolBgpPolicyExportRulesMatchAddressPrefix.Misc
											}
											if oProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact != nil {
												nestedProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact = util.AsBool(oProtocolBgpPolicyExportRulesMatchAddressPrefix.Exact, nil)
											}
											if oProtocolBgpPolicyExportRulesMatchAddressPrefix.Name != "" {
												nestedProtocolBgpPolicyExportRulesMatchAddressPrefix.Name = oProtocolBgpPolicyExportRulesMatchAddressPrefix.Name
											}
											nestedProtocolBgpPolicyExportRules.Match.AddressPrefix = append(nestedProtocolBgpPolicyExportRules.Match.AddressPrefix, nestedProtocolBgpPolicyExportRulesMatchAddressPrefix)
										}
									}
									if oProtocolBgpPolicyExportRules.Match.Nexthop != nil {
										nestedProtocolBgpPolicyExportRules.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyExportRules.Match.Nexthop)
									}
									if oProtocolBgpPolicyExportRules.Match.FromPeer != nil {
										nestedProtocolBgpPolicyExportRules.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyExportRules.Match.FromPeer)
									}
									if oProtocolBgpPolicyExportRules.Match.AsPath != nil {
										nestedProtocolBgpPolicyExportRules.Match.AsPath = &ProtocolBgpPolicyExportRulesMatchAsPath{}
										if oProtocolBgpPolicyExportRules.Match.AsPath.Misc != nil {
											entry.Misc["ProtocolBgpPolicyExportRulesMatchAsPath"] = oProtocolBgpPolicyExportRules.Match.AsPath.Misc
										}
										if oProtocolBgpPolicyExportRules.Match.AsPath.Regex != nil {
											nestedProtocolBgpPolicyExportRules.Match.AsPath.Regex = oProtocolBgpPolicyExportRules.Match.AsPath.Regex
										}
									}
									if oProtocolBgpPolicyExportRules.Match.Community != nil {
										nestedProtocolBgpPolicyExportRules.Match.Community = &ProtocolBgpPolicyExportRulesMatchCommunity{}
										if oProtocolBgpPolicyExportRules.Match.Community.Misc != nil {
											entry.Misc["ProtocolBgpPolicyExportRulesMatchCommunity"] = oProtocolBgpPolicyExportRules.Match.Community.Misc
										}
										if oProtocolBgpPolicyExportRules.Match.Community.Regex != nil {
											nestedProtocolBgpPolicyExportRules.Match.Community.Regex = oProtocolBgpPolicyExportRules.Match.Community.Regex
										}
									}
									if oProtocolBgpPolicyExportRules.Match.ExtendedCommunity != nil {
										nestedProtocolBgpPolicyExportRules.Match.ExtendedCommunity = &ProtocolBgpPolicyExportRulesMatchExtendedCommunity{}
										if oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Misc != nil {
											entry.Misc["ProtocolBgpPolicyExportRulesMatchExtendedCommunity"] = oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Misc
										}
										if oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex != nil {
											nestedProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyExportRules.Match.ExtendedCommunity.Regex
										}
									}
								}
								if oProtocolBgpPolicyExportRules.Action != nil {
									nestedProtocolBgpPolicyExportRules.Action = &ProtocolBgpPolicyExportRulesAction{}
									if oProtocolBgpPolicyExportRules.Action.Misc != nil {
										entry.Misc["ProtocolBgpPolicyExportRulesAction"] = oProtocolBgpPolicyExportRules.Action.Misc
									}
									if oProtocolBgpPolicyExportRules.Action.Deny != nil {
										nestedProtocolBgpPolicyExportRules.Action.Deny = &ProtocolBgpPolicyExportRulesActionDeny{}
										if oProtocolBgpPolicyExportRules.Action.Deny.Misc != nil {
											entry.Misc["ProtocolBgpPolicyExportRulesActionDeny"] = oProtocolBgpPolicyExportRules.Action.Deny.Misc
										}
									}
									if oProtocolBgpPolicyExportRules.Action.Allow != nil {
										nestedProtocolBgpPolicyExportRules.Action.Allow = &ProtocolBgpPolicyExportRulesActionAllow{}
										if oProtocolBgpPolicyExportRules.Action.Allow.Misc != nil {
											entry.Misc["ProtocolBgpPolicyExportRulesActionAllow"] = oProtocolBgpPolicyExportRules.Action.Allow.Misc
										}
										if oProtocolBgpPolicyExportRules.Action.Allow.Update != nil {
											nestedProtocolBgpPolicyExportRules.Action.Allow.Update = &ProtocolBgpPolicyExportRulesActionAllowUpdate{}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Misc != nil {
												entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdate"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.Misc
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPathLimit
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath{}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Misc
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Remove.Misc
													}
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.Prepend
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.RemoveAndPrepend
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.AsPath.None.Misc
													}
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity{}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Misc
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite = util.MemToStr(oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Overwrite)
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.None.Misc
													}
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll = &ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveAll.Misc
													}
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex = oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.RemoveRegex
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append = util.MemToStr(oProtocolBgpPolicyExportRules.Action.Allow.Update.Community.Append)
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity{}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc
													}
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append = util.MemToStr(oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Append)
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite = util.MemToStr(oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.Overwrite)
												}
												if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None != nil {
													nestedProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None = &ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone{}
													if oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone"] = oProtocolBgpPolicyExportRules.Action.Allow.Update.ExtendedCommunity.None.Misc
													}
												}
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference = oProtocolBgpPolicyExportRules.Action.Allow.Update.LocalPreference
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Med != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Med = oProtocolBgpPolicyExportRules.Action.Allow.Update.Med
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop = oProtocolBgpPolicyExportRules.Action.Allow.Update.Nexthop
											}
											if oProtocolBgpPolicyExportRules.Action.Allow.Update.Origin != nil {
												nestedProtocolBgpPolicyExportRules.Action.Allow.Update.Origin = oProtocolBgpPolicyExportRules.Action.Allow.Update.Origin
											}
										}
									}
								}
								if oProtocolBgpPolicyExportRules.Name != "" {
									nestedProtocolBgpPolicyExportRules.Name = oProtocolBgpPolicyExportRules.Name
								}
								if oProtocolBgpPolicyExportRules.Enable != nil {
									nestedProtocolBgpPolicyExportRules.Enable = util.AsBool(oProtocolBgpPolicyExportRules.Enable, nil)
								}
								nestedProtocol.Bgp.Policy.Export.Rules = append(nestedProtocol.Bgp.Policy.Export.Rules, nestedProtocolBgpPolicyExportRules)
							}
						}
					}
					if o.Protocol.Bgp.Policy.Import != nil {
						nestedProtocol.Bgp.Policy.Import = &ProtocolBgpPolicyImport{}
						if o.Protocol.Bgp.Policy.Import.Misc != nil {
							entry.Misc["ProtocolBgpPolicyImport"] = o.Protocol.Bgp.Policy.Import.Misc
						}
						if o.Protocol.Bgp.Policy.Import.Rules != nil {
							nestedProtocol.Bgp.Policy.Import.Rules = []ProtocolBgpPolicyImportRules{}
							for _, oProtocolBgpPolicyImportRules := range o.Protocol.Bgp.Policy.Import.Rules {
								nestedProtocolBgpPolicyImportRules := ProtocolBgpPolicyImportRules{}
								if oProtocolBgpPolicyImportRules.Misc != nil {
									entry.Misc["ProtocolBgpPolicyImportRules"] = oProtocolBgpPolicyImportRules.Misc
								}
								if oProtocolBgpPolicyImportRules.Enable != nil {
									nestedProtocolBgpPolicyImportRules.Enable = util.AsBool(oProtocolBgpPolicyImportRules.Enable, nil)
								}
								if oProtocolBgpPolicyImportRules.UsedBy != nil {
									nestedProtocolBgpPolicyImportRules.UsedBy = util.MemToStr(oProtocolBgpPolicyImportRules.UsedBy)
								}
								if oProtocolBgpPolicyImportRules.Match != nil {
									nestedProtocolBgpPolicyImportRules.Match = &ProtocolBgpPolicyImportRulesMatch{}
									if oProtocolBgpPolicyImportRules.Match.Misc != nil {
										entry.Misc["ProtocolBgpPolicyImportRulesMatch"] = oProtocolBgpPolicyImportRules.Match.Misc
									}
									if oProtocolBgpPolicyImportRules.Match.FromPeer != nil {
										nestedProtocolBgpPolicyImportRules.Match.FromPeer = util.MemToStr(oProtocolBgpPolicyImportRules.Match.FromPeer)
									}
									if oProtocolBgpPolicyImportRules.Match.AsPath != nil {
										nestedProtocolBgpPolicyImportRules.Match.AsPath = &ProtocolBgpPolicyImportRulesMatchAsPath{}
										if oProtocolBgpPolicyImportRules.Match.AsPath.Misc != nil {
											entry.Misc["ProtocolBgpPolicyImportRulesMatchAsPath"] = oProtocolBgpPolicyImportRules.Match.AsPath.Misc
										}
										if oProtocolBgpPolicyImportRules.Match.AsPath.Regex != nil {
											nestedProtocolBgpPolicyImportRules.Match.AsPath.Regex = oProtocolBgpPolicyImportRules.Match.AsPath.Regex
										}
									}
									if oProtocolBgpPolicyImportRules.Match.Community != nil {
										nestedProtocolBgpPolicyImportRules.Match.Community = &ProtocolBgpPolicyImportRulesMatchCommunity{}
										if oProtocolBgpPolicyImportRules.Match.Community.Misc != nil {
											entry.Misc["ProtocolBgpPolicyImportRulesMatchCommunity"] = oProtocolBgpPolicyImportRules.Match.Community.Misc
										}
										if oProtocolBgpPolicyImportRules.Match.Community.Regex != nil {
											nestedProtocolBgpPolicyImportRules.Match.Community.Regex = oProtocolBgpPolicyImportRules.Match.Community.Regex
										}
									}
									if oProtocolBgpPolicyImportRules.Match.ExtendedCommunity != nil {
										nestedProtocolBgpPolicyImportRules.Match.ExtendedCommunity = &ProtocolBgpPolicyImportRulesMatchExtendedCommunity{}
										if oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Misc != nil {
											entry.Misc["ProtocolBgpPolicyImportRulesMatchExtendedCommunity"] = oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Misc
										}
										if oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex != nil {
											nestedProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex = oProtocolBgpPolicyImportRules.Match.ExtendedCommunity.Regex
										}
									}
									if oProtocolBgpPolicyImportRules.Match.RouteTable != nil {
										nestedProtocolBgpPolicyImportRules.Match.RouteTable = oProtocolBgpPolicyImportRules.Match.RouteTable
									}
									if oProtocolBgpPolicyImportRules.Match.Med != nil {
										nestedProtocolBgpPolicyImportRules.Match.Med = oProtocolBgpPolicyImportRules.Match.Med
									}
									if oProtocolBgpPolicyImportRules.Match.AddressPrefix != nil {
										nestedProtocolBgpPolicyImportRules.Match.AddressPrefix = []ProtocolBgpPolicyImportRulesMatchAddressPrefix{}
										for _, oProtocolBgpPolicyImportRulesMatchAddressPrefix := range oProtocolBgpPolicyImportRules.Match.AddressPrefix {
											nestedProtocolBgpPolicyImportRulesMatchAddressPrefix := ProtocolBgpPolicyImportRulesMatchAddressPrefix{}
											if oProtocolBgpPolicyImportRulesMatchAddressPrefix.Misc != nil {
												entry.Misc["ProtocolBgpPolicyImportRulesMatchAddressPrefix"] = oProtocolBgpPolicyImportRulesMatchAddressPrefix.Misc
											}
											if oProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact != nil {
												nestedProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact = util.AsBool(oProtocolBgpPolicyImportRulesMatchAddressPrefix.Exact, nil)
											}
											if oProtocolBgpPolicyImportRulesMatchAddressPrefix.Name != "" {
												nestedProtocolBgpPolicyImportRulesMatchAddressPrefix.Name = oProtocolBgpPolicyImportRulesMatchAddressPrefix.Name
											}
											nestedProtocolBgpPolicyImportRules.Match.AddressPrefix = append(nestedProtocolBgpPolicyImportRules.Match.AddressPrefix, nestedProtocolBgpPolicyImportRulesMatchAddressPrefix)
										}
									}
									if oProtocolBgpPolicyImportRules.Match.Nexthop != nil {
										nestedProtocolBgpPolicyImportRules.Match.Nexthop = util.MemToStr(oProtocolBgpPolicyImportRules.Match.Nexthop)
									}
								}
								if oProtocolBgpPolicyImportRules.Action != nil {
									nestedProtocolBgpPolicyImportRules.Action = &ProtocolBgpPolicyImportRulesAction{}
									if oProtocolBgpPolicyImportRules.Action.Misc != nil {
										entry.Misc["ProtocolBgpPolicyImportRulesAction"] = oProtocolBgpPolicyImportRules.Action.Misc
									}
									if oProtocolBgpPolicyImportRules.Action.Deny != nil {
										nestedProtocolBgpPolicyImportRules.Action.Deny = &ProtocolBgpPolicyImportRulesActionDeny{}
										if oProtocolBgpPolicyImportRules.Action.Deny.Misc != nil {
											entry.Misc["ProtocolBgpPolicyImportRulesActionDeny"] = oProtocolBgpPolicyImportRules.Action.Deny.Misc
										}
									}
									if oProtocolBgpPolicyImportRules.Action.Allow != nil {
										nestedProtocolBgpPolicyImportRules.Action.Allow = &ProtocolBgpPolicyImportRulesActionAllow{}
										if oProtocolBgpPolicyImportRules.Action.Allow.Misc != nil {
											entry.Misc["ProtocolBgpPolicyImportRulesActionAllow"] = oProtocolBgpPolicyImportRules.Action.Allow.Misc
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Update != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Update = &ProtocolBgpPolicyImportRulesActionAllowUpdate{}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Misc != nil {
												entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdate"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.Misc
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Weight != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Weight = oProtocolBgpPolicyImportRules.Action.Allow.Update.Weight
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Origin != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Origin = oProtocolBgpPolicyImportRules.Action.Allow.Update.Origin
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit = oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPathLimit
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity{}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Misc != nil {
													entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Misc
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveAll.Misc
													}
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex = oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.RemoveRegex
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append = util.MemToStr(oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Append)
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite = util.MemToStr(oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.Overwrite)
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.ExtendedCommunity.None.Misc
													}
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference = oProtocolBgpPolicyImportRules.Action.Allow.Update.LocalPreference
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Med != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Med = oProtocolBgpPolicyImportRules.Action.Allow.Update.Med
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop = oProtocolBgpPolicyImportRules.Action.Allow.Update.Nexthop
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath{}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Misc != nil {
													entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Misc
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.None.Misc
													}
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove = &ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.AsPath.Remove.Misc
													}
												}
											}
											if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community != nil {
												nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity{}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Misc != nil {
													entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Misc
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveAll.Misc
													}
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex = oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.RemoveRegex
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append = util.MemToStr(oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Append)
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite = util.MemToStr(oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.Overwrite)
												}
												if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None != nil {
													nestedProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None = &ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone{}
													if oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None.Misc != nil {
														entry.Misc["ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone"] = oProtocolBgpPolicyImportRules.Action.Allow.Update.Community.None.Misc
													}
												}
											}
										}
										if oProtocolBgpPolicyImportRules.Action.Allow.Dampening != nil {
											nestedProtocolBgpPolicyImportRules.Action.Allow.Dampening = oProtocolBgpPolicyImportRules.Action.Allow.Dampening
										}
									}
								}
								if oProtocolBgpPolicyImportRules.Name != "" {
									nestedProtocolBgpPolicyImportRules.Name = oProtocolBgpPolicyImportRules.Name
								}
								nestedProtocol.Bgp.Policy.Import.Rules = append(nestedProtocol.Bgp.Policy.Import.Rules, nestedProtocolBgpPolicyImportRules)
							}
						}
					}
				}
				if o.Protocol.Bgp.AuthProfile != nil {
					nestedProtocol.Bgp.AuthProfile = []ProtocolBgpAuthProfile{}
					for _, oProtocolBgpAuthProfile := range o.Protocol.Bgp.AuthProfile {
						nestedProtocolBgpAuthProfile := ProtocolBgpAuthProfile{}
						if oProtocolBgpAuthProfile.Misc != nil {
							entry.Misc["ProtocolBgpAuthProfile"] = oProtocolBgpAuthProfile.Misc
						}
						if oProtocolBgpAuthProfile.Secret != nil {
							nestedProtocolBgpAuthProfile.Secret = oProtocolBgpAuthProfile.Secret
						}
						if oProtocolBgpAuthProfile.Name != "" {
							nestedProtocolBgpAuthProfile.Name = oProtocolBgpAuthProfile.Name
						}
						nestedProtocol.Bgp.AuthProfile = append(nestedProtocol.Bgp.AuthProfile, nestedProtocolBgpAuthProfile)
					}
				}
				if o.Protocol.Bgp.Enable != nil {
					nestedProtocol.Bgp.Enable = util.AsBool(o.Protocol.Bgp.Enable, nil)
				}
			}
			if o.Protocol.Ospf != nil {
				nestedProtocol.Ospf = &ProtocolOspf{}
				if o.Protocol.Ospf.Misc != nil {
					entry.Misc["ProtocolOspf"] = o.Protocol.Ospf.Misc
				}
				if o.Protocol.Ospf.Area != nil {
					nestedProtocol.Ospf.Area = []ProtocolOspfArea{}
					for _, oProtocolOspfArea := range o.Protocol.Ospf.Area {
						nestedProtocolOspfArea := ProtocolOspfArea{}
						if oProtocolOspfArea.Misc != nil {
							entry.Misc["ProtocolOspfArea"] = oProtocolOspfArea.Misc
						}
						if oProtocolOspfArea.Type != nil {
							nestedProtocolOspfArea.Type = &ProtocolOspfAreaType{}
							if oProtocolOspfArea.Type.Misc != nil {
								entry.Misc["ProtocolOspfAreaType"] = oProtocolOspfArea.Type.Misc
							}
							if oProtocolOspfArea.Type.Normal != nil {
								nestedProtocolOspfArea.Type.Normal = &ProtocolOspfAreaTypeNormal{}
								if oProtocolOspfArea.Type.Normal.Misc != nil {
									entry.Misc["ProtocolOspfAreaTypeNormal"] = oProtocolOspfArea.Type.Normal.Misc
								}
							}
							if oProtocolOspfArea.Type.Stub != nil {
								nestedProtocolOspfArea.Type.Stub = &ProtocolOspfAreaTypeStub{}
								if oProtocolOspfArea.Type.Stub.Misc != nil {
									entry.Misc["ProtocolOspfAreaTypeStub"] = oProtocolOspfArea.Type.Stub.Misc
								}
								if oProtocolOspfArea.Type.Stub.AcceptSummary != nil {
									nestedProtocolOspfArea.Type.Stub.AcceptSummary = util.AsBool(oProtocolOspfArea.Type.Stub.AcceptSummary, nil)
								}
								if oProtocolOspfArea.Type.Stub.DefaultRoute != nil {
									nestedProtocolOspfArea.Type.Stub.DefaultRoute = &ProtocolOspfAreaTypeStubDefaultRoute{}
									if oProtocolOspfArea.Type.Stub.DefaultRoute.Misc != nil {
										entry.Misc["ProtocolOspfAreaTypeStubDefaultRoute"] = oProtocolOspfArea.Type.Stub.DefaultRoute.Misc
									}
									if oProtocolOspfArea.Type.Stub.DefaultRoute.Disable != nil {
										nestedProtocolOspfArea.Type.Stub.DefaultRoute.Disable = &ProtocolOspfAreaTypeStubDefaultRouteDisable{}
										if oProtocolOspfArea.Type.Stub.DefaultRoute.Disable.Misc != nil {
											entry.Misc["ProtocolOspfAreaTypeStubDefaultRouteDisable"] = oProtocolOspfArea.Type.Stub.DefaultRoute.Disable.Misc
										}
									}
									if oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise != nil {
										nestedProtocolOspfArea.Type.Stub.DefaultRoute.Advertise = &ProtocolOspfAreaTypeStubDefaultRouteAdvertise{}
										if oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Misc != nil {
											entry.Misc["ProtocolOspfAreaTypeStubDefaultRouteAdvertise"] = oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Misc
										}
										if oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric != nil {
											nestedProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric = oProtocolOspfArea.Type.Stub.DefaultRoute.Advertise.Metric
										}
									}
								}
							}
							if oProtocolOspfArea.Type.Nssa != nil {
								nestedProtocolOspfArea.Type.Nssa = &ProtocolOspfAreaTypeNssa{}
								if oProtocolOspfArea.Type.Nssa.Misc != nil {
									entry.Misc["ProtocolOspfAreaTypeNssa"] = oProtocolOspfArea.Type.Nssa.Misc
								}
								if oProtocolOspfArea.Type.Nssa.DefaultRoute != nil {
									nestedProtocolOspfArea.Type.Nssa.DefaultRoute = &ProtocolOspfAreaTypeNssaDefaultRoute{}
									if oProtocolOspfArea.Type.Nssa.DefaultRoute.Misc != nil {
										entry.Misc["ProtocolOspfAreaTypeNssaDefaultRoute"] = oProtocolOspfArea.Type.Nssa.DefaultRoute.Misc
									}
									if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise != nil {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise = &ProtocolOspfAreaTypeNssaDefaultRouteAdvertise{}
										if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Misc != nil {
											entry.Misc["ProtocolOspfAreaTypeNssaDefaultRouteAdvertise"] = oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Misc
										}
										if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric != nil {
											nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric = oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Metric
										}
										if oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type != nil {
											nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type = oProtocolOspfArea.Type.Nssa.DefaultRoute.Advertise.Type
										}
									}
									if oProtocolOspfArea.Type.Nssa.DefaultRoute.Disable != nil {
										nestedProtocolOspfArea.Type.Nssa.DefaultRoute.Disable = &ProtocolOspfAreaTypeNssaDefaultRouteDisable{}
										if oProtocolOspfArea.Type.Nssa.DefaultRoute.Disable.Misc != nil {
											entry.Misc["ProtocolOspfAreaTypeNssaDefaultRouteDisable"] = oProtocolOspfArea.Type.Nssa.DefaultRoute.Disable.Misc
										}
									}
								}
								if oProtocolOspfArea.Type.Nssa.NssaExtRange != nil {
									nestedProtocolOspfArea.Type.Nssa.NssaExtRange = []ProtocolOspfAreaTypeNssaNssaExtRange{}
									for _, oProtocolOspfAreaTypeNssaNssaExtRange := range oProtocolOspfArea.Type.Nssa.NssaExtRange {
										nestedProtocolOspfAreaTypeNssaNssaExtRange := ProtocolOspfAreaTypeNssaNssaExtRange{}
										if oProtocolOspfAreaTypeNssaNssaExtRange.Misc != nil {
											entry.Misc["ProtocolOspfAreaTypeNssaNssaExtRange"] = oProtocolOspfAreaTypeNssaNssaExtRange.Misc
										}
										if oProtocolOspfAreaTypeNssaNssaExtRange.Name != "" {
											nestedProtocolOspfAreaTypeNssaNssaExtRange.Name = oProtocolOspfAreaTypeNssaNssaExtRange.Name
										}
										if oProtocolOspfAreaTypeNssaNssaExtRange.Advertise != nil {
											nestedProtocolOspfAreaTypeNssaNssaExtRange.Advertise = &ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise{}
											if oProtocolOspfAreaTypeNssaNssaExtRange.Advertise.Misc != nil {
												entry.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise"] = oProtocolOspfAreaTypeNssaNssaExtRange.Advertise.Misc
											}
										}
										if oProtocolOspfAreaTypeNssaNssaExtRange.Suppress != nil {
											nestedProtocolOspfAreaTypeNssaNssaExtRange.Suppress = &ProtocolOspfAreaTypeNssaNssaExtRangeSuppress{}
											if oProtocolOspfAreaTypeNssaNssaExtRange.Suppress.Misc != nil {
												entry.Misc["ProtocolOspfAreaTypeNssaNssaExtRangeSuppress"] = oProtocolOspfAreaTypeNssaNssaExtRange.Suppress.Misc
											}
										}
										nestedProtocolOspfArea.Type.Nssa.NssaExtRange = append(nestedProtocolOspfArea.Type.Nssa.NssaExtRange, nestedProtocolOspfAreaTypeNssaNssaExtRange)
									}
								}
								if oProtocolOspfArea.Type.Nssa.AcceptSummary != nil {
									nestedProtocolOspfArea.Type.Nssa.AcceptSummary = util.AsBool(oProtocolOspfArea.Type.Nssa.AcceptSummary, nil)
								}
							}
						}
						if oProtocolOspfArea.Range != nil {
							nestedProtocolOspfArea.Range = []ProtocolOspfAreaRange{}
							for _, oProtocolOspfAreaRange := range oProtocolOspfArea.Range {
								nestedProtocolOspfAreaRange := ProtocolOspfAreaRange{}
								if oProtocolOspfAreaRange.Misc != nil {
									entry.Misc["ProtocolOspfAreaRange"] = oProtocolOspfAreaRange.Misc
								}
								if oProtocolOspfAreaRange.Name != "" {
									nestedProtocolOspfAreaRange.Name = oProtocolOspfAreaRange.Name
								}
								if oProtocolOspfAreaRange.Advertise != nil {
									nestedProtocolOspfAreaRange.Advertise = &ProtocolOspfAreaRangeAdvertise{}
									if oProtocolOspfAreaRange.Advertise.Misc != nil {
										entry.Misc["ProtocolOspfAreaRangeAdvertise"] = oProtocolOspfAreaRange.Advertise.Misc
									}
								}
								if oProtocolOspfAreaRange.Suppress != nil {
									nestedProtocolOspfAreaRange.Suppress = &ProtocolOspfAreaRangeSuppress{}
									if oProtocolOspfAreaRange.Suppress.Misc != nil {
										entry.Misc["ProtocolOspfAreaRangeSuppress"] = oProtocolOspfAreaRange.Suppress.Misc
									}
								}
								nestedProtocolOspfArea.Range = append(nestedProtocolOspfArea.Range, nestedProtocolOspfAreaRange)
							}
						}
						if oProtocolOspfArea.Interface != nil {
							nestedProtocolOspfArea.Interface = []ProtocolOspfAreaInterface{}
							for _, oProtocolOspfAreaInterface := range oProtocolOspfArea.Interface {
								nestedProtocolOspfAreaInterface := ProtocolOspfAreaInterface{}
								if oProtocolOspfAreaInterface.Misc != nil {
									entry.Misc["ProtocolOspfAreaInterface"] = oProtocolOspfAreaInterface.Misc
								}
								if oProtocolOspfAreaInterface.GrDelay != nil {
									nestedProtocolOspfAreaInterface.GrDelay = oProtocolOspfAreaInterface.GrDelay
								}
								if oProtocolOspfAreaInterface.Bfd != nil {
									nestedProtocolOspfAreaInterface.Bfd = &ProtocolOspfAreaInterfaceBfd{}
									if oProtocolOspfAreaInterface.Bfd.Misc != nil {
										entry.Misc["ProtocolOspfAreaInterfaceBfd"] = oProtocolOspfAreaInterface.Bfd.Misc
									}
									if oProtocolOspfAreaInterface.Bfd.Profile != nil {
										nestedProtocolOspfAreaInterface.Bfd.Profile = oProtocolOspfAreaInterface.Bfd.Profile
									}
								}
								if oProtocolOspfAreaInterface.Name != "" {
									nestedProtocolOspfAreaInterface.Name = oProtocolOspfAreaInterface.Name
								}
								if oProtocolOspfAreaInterface.Priority != nil {
									nestedProtocolOspfAreaInterface.Priority = oProtocolOspfAreaInterface.Priority
								}
								if oProtocolOspfAreaInterface.RetransmitInterval != nil {
									nestedProtocolOspfAreaInterface.RetransmitInterval = oProtocolOspfAreaInterface.RetransmitInterval
								}
								if oProtocolOspfAreaInterface.Authentication != nil {
									nestedProtocolOspfAreaInterface.Authentication = oProtocolOspfAreaInterface.Authentication
								}
								if oProtocolOspfAreaInterface.Metric != nil {
									nestedProtocolOspfAreaInterface.Metric = oProtocolOspfAreaInterface.Metric
								}
								if oProtocolOspfAreaInterface.Neighbor != nil {
									nestedProtocolOspfAreaInterface.Neighbor = []ProtocolOspfAreaInterfaceNeighbor{}
									for _, oProtocolOspfAreaInterfaceNeighbor := range oProtocolOspfAreaInterface.Neighbor {
										nestedProtocolOspfAreaInterfaceNeighbor := ProtocolOspfAreaInterfaceNeighbor{}
										if oProtocolOspfAreaInterfaceNeighbor.Misc != nil {
											entry.Misc["ProtocolOspfAreaInterfaceNeighbor"] = oProtocolOspfAreaInterfaceNeighbor.Misc
										}
										if oProtocolOspfAreaInterfaceNeighbor.Name != "" {
											nestedProtocolOspfAreaInterfaceNeighbor.Name = oProtocolOspfAreaInterfaceNeighbor.Name
										}
										nestedProtocolOspfAreaInterface.Neighbor = append(nestedProtocolOspfAreaInterface.Neighbor, nestedProtocolOspfAreaInterfaceNeighbor)
									}
								}
								if oProtocolOspfAreaInterface.LinkType != nil {
									nestedProtocolOspfAreaInterface.LinkType = &ProtocolOspfAreaInterfaceLinkType{}
									if oProtocolOspfAreaInterface.LinkType.Misc != nil {
										entry.Misc["ProtocolOspfAreaInterfaceLinkType"] = oProtocolOspfAreaInterface.LinkType.Misc
									}
									if oProtocolOspfAreaInterface.LinkType.Broadcast != nil {
										nestedProtocolOspfAreaInterface.LinkType.Broadcast = &ProtocolOspfAreaInterfaceLinkTypeBroadcast{}
										if oProtocolOspfAreaInterface.LinkType.Broadcast.Misc != nil {
											entry.Misc["ProtocolOspfAreaInterfaceLinkTypeBroadcast"] = oProtocolOspfAreaInterface.LinkType.Broadcast.Misc
										}
									}
									if oProtocolOspfAreaInterface.LinkType.P2p != nil {
										nestedProtocolOspfAreaInterface.LinkType.P2p = &ProtocolOspfAreaInterfaceLinkTypeP2p{}
										if oProtocolOspfAreaInterface.LinkType.P2p.Misc != nil {
											entry.Misc["ProtocolOspfAreaInterfaceLinkTypeP2p"] = oProtocolOspfAreaInterface.LinkType.P2p.Misc
										}
									}
									if oProtocolOspfAreaInterface.LinkType.P2mp != nil {
										nestedProtocolOspfAreaInterface.LinkType.P2mp = &ProtocolOspfAreaInterfaceLinkTypeP2mp{}
										if oProtocolOspfAreaInterface.LinkType.P2mp.Misc != nil {
											entry.Misc["ProtocolOspfAreaInterfaceLinkTypeP2mp"] = oProtocolOspfAreaInterface.LinkType.P2mp.Misc
										}
									}
								}
								if oProtocolOspfAreaInterface.Enable != nil {
									nestedProtocolOspfAreaInterface.Enable = util.AsBool(oProtocolOspfAreaInterface.Enable, nil)
								}
								if oProtocolOspfAreaInterface.Passive != nil {
									nestedProtocolOspfAreaInterface.Passive = util.AsBool(oProtocolOspfAreaInterface.Passive, nil)
								}
								if oProtocolOspfAreaInterface.HelloInterval != nil {
									nestedProtocolOspfAreaInterface.HelloInterval = oProtocolOspfAreaInterface.HelloInterval
								}
								if oProtocolOspfAreaInterface.DeadCounts != nil {
									nestedProtocolOspfAreaInterface.DeadCounts = oProtocolOspfAreaInterface.DeadCounts
								}
								if oProtocolOspfAreaInterface.TransitDelay != nil {
									nestedProtocolOspfAreaInterface.TransitDelay = oProtocolOspfAreaInterface.TransitDelay
								}
								nestedProtocolOspfArea.Interface = append(nestedProtocolOspfArea.Interface, nestedProtocolOspfAreaInterface)
							}
						}
						if oProtocolOspfArea.VirtualLink != nil {
							nestedProtocolOspfArea.VirtualLink = []ProtocolOspfAreaVirtualLink{}
							for _, oProtocolOspfAreaVirtualLink := range oProtocolOspfArea.VirtualLink {
								nestedProtocolOspfAreaVirtualLink := ProtocolOspfAreaVirtualLink{}
								if oProtocolOspfAreaVirtualLink.Misc != nil {
									entry.Misc["ProtocolOspfAreaVirtualLink"] = oProtocolOspfAreaVirtualLink.Misc
								}
								if oProtocolOspfAreaVirtualLink.TransitDelay != nil {
									nestedProtocolOspfAreaVirtualLink.TransitDelay = oProtocolOspfAreaVirtualLink.TransitDelay
								}
								if oProtocolOspfAreaVirtualLink.Authentication != nil {
									nestedProtocolOspfAreaVirtualLink.Authentication = oProtocolOspfAreaVirtualLink.Authentication
								}
								if oProtocolOspfAreaVirtualLink.Name != "" {
									nestedProtocolOspfAreaVirtualLink.Name = oProtocolOspfAreaVirtualLink.Name
								}
								if oProtocolOspfAreaVirtualLink.NeighborId != nil {
									nestedProtocolOspfAreaVirtualLink.NeighborId = oProtocolOspfAreaVirtualLink.NeighborId
								}
								if oProtocolOspfAreaVirtualLink.TransitAreaId != nil {
									nestedProtocolOspfAreaVirtualLink.TransitAreaId = oProtocolOspfAreaVirtualLink.TransitAreaId
								}
								if oProtocolOspfAreaVirtualLink.DeadCounts != nil {
									nestedProtocolOspfAreaVirtualLink.DeadCounts = oProtocolOspfAreaVirtualLink.DeadCounts
								}
								if oProtocolOspfAreaVirtualLink.RetransmitInterval != nil {
									nestedProtocolOspfAreaVirtualLink.RetransmitInterval = oProtocolOspfAreaVirtualLink.RetransmitInterval
								}
								if oProtocolOspfAreaVirtualLink.Enable != nil {
									nestedProtocolOspfAreaVirtualLink.Enable = util.AsBool(oProtocolOspfAreaVirtualLink.Enable, nil)
								}
								if oProtocolOspfAreaVirtualLink.HelloInterval != nil {
									nestedProtocolOspfAreaVirtualLink.HelloInterval = oProtocolOspfAreaVirtualLink.HelloInterval
								}
								if oProtocolOspfAreaVirtualLink.Bfd != nil {
									nestedProtocolOspfAreaVirtualLink.Bfd = &ProtocolOspfAreaVirtualLinkBfd{}
									if oProtocolOspfAreaVirtualLink.Bfd.Misc != nil {
										entry.Misc["ProtocolOspfAreaVirtualLinkBfd"] = oProtocolOspfAreaVirtualLink.Bfd.Misc
									}
									if oProtocolOspfAreaVirtualLink.Bfd.Profile != nil {
										nestedProtocolOspfAreaVirtualLink.Bfd.Profile = oProtocolOspfAreaVirtualLink.Bfd.Profile
									}
								}
								nestedProtocolOspfArea.VirtualLink = append(nestedProtocolOspfArea.VirtualLink, nestedProtocolOspfAreaVirtualLink)
							}
						}
						if oProtocolOspfArea.Name != "" {
							nestedProtocolOspfArea.Name = oProtocolOspfArea.Name
						}
						nestedProtocol.Ospf.Area = append(nestedProtocol.Ospf.Area, nestedProtocolOspfArea)
					}
				}
				if o.Protocol.Ospf.AuthProfile != nil {
					nestedProtocol.Ospf.AuthProfile = []ProtocolOspfAuthProfile{}
					for _, oProtocolOspfAuthProfile := range o.Protocol.Ospf.AuthProfile {
						nestedProtocolOspfAuthProfile := ProtocolOspfAuthProfile{}
						if oProtocolOspfAuthProfile.Misc != nil {
							entry.Misc["ProtocolOspfAuthProfile"] = oProtocolOspfAuthProfile.Misc
						}
						if oProtocolOspfAuthProfile.Name != "" {
							nestedProtocolOspfAuthProfile.Name = oProtocolOspfAuthProfile.Name
						}
						if oProtocolOspfAuthProfile.Password != nil {
							nestedProtocolOspfAuthProfile.Password = oProtocolOspfAuthProfile.Password
						}
						if oProtocolOspfAuthProfile.Md5 != nil {
							nestedProtocolOspfAuthProfile.Md5 = []ProtocolOspfAuthProfileMd5{}
							for _, oProtocolOspfAuthProfileMd5 := range oProtocolOspfAuthProfile.Md5 {
								nestedProtocolOspfAuthProfileMd5 := ProtocolOspfAuthProfileMd5{}
								if oProtocolOspfAuthProfileMd5.Misc != nil {
									entry.Misc["ProtocolOspfAuthProfileMd5"] = oProtocolOspfAuthProfileMd5.Misc
								}
								if oProtocolOspfAuthProfileMd5.Key != nil {
									nestedProtocolOspfAuthProfileMd5.Key = oProtocolOspfAuthProfileMd5.Key
								}
								if oProtocolOspfAuthProfileMd5.Preferred != nil {
									nestedProtocolOspfAuthProfileMd5.Preferred = util.AsBool(oProtocolOspfAuthProfileMd5.Preferred, nil)
								}
								if oProtocolOspfAuthProfileMd5.Name != "" {
									nestedProtocolOspfAuthProfileMd5.Name = oProtocolOspfAuthProfileMd5.Name
								}
								nestedProtocolOspfAuthProfile.Md5 = append(nestedProtocolOspfAuthProfile.Md5, nestedProtocolOspfAuthProfileMd5)
							}
						}
						nestedProtocol.Ospf.AuthProfile = append(nestedProtocol.Ospf.AuthProfile, nestedProtocolOspfAuthProfile)
					}
				}
				if o.Protocol.Ospf.ExportRules != nil {
					nestedProtocol.Ospf.ExportRules = []ProtocolOspfExportRules{}
					for _, oProtocolOspfExportRules := range o.Protocol.Ospf.ExportRules {
						nestedProtocolOspfExportRules := ProtocolOspfExportRules{}
						if oProtocolOspfExportRules.Misc != nil {
							entry.Misc["ProtocolOspfExportRules"] = oProtocolOspfExportRules.Misc
						}
						if oProtocolOspfExportRules.NewTag != nil {
							nestedProtocolOspfExportRules.NewTag = oProtocolOspfExportRules.NewTag
						}
						if oProtocolOspfExportRules.Metric != nil {
							nestedProtocolOspfExportRules.Metric = oProtocolOspfExportRules.Metric
						}
						if oProtocolOspfExportRules.Name != "" {
							nestedProtocolOspfExportRules.Name = oProtocolOspfExportRules.Name
						}
						if oProtocolOspfExportRules.NewPathType != nil {
							nestedProtocolOspfExportRules.NewPathType = oProtocolOspfExportRules.NewPathType
						}
						nestedProtocol.Ospf.ExportRules = append(nestedProtocol.Ospf.ExportRules, nestedProtocolOspfExportRules)
					}
				}
				if o.Protocol.Ospf.GlobalBfd != nil {
					nestedProtocol.Ospf.GlobalBfd = &ProtocolOspfGlobalBfd{}
					if o.Protocol.Ospf.GlobalBfd.Misc != nil {
						entry.Misc["ProtocolOspfGlobalBfd"] = o.Protocol.Ospf.GlobalBfd.Misc
					}
					if o.Protocol.Ospf.GlobalBfd.Profile != nil {
						nestedProtocol.Ospf.GlobalBfd.Profile = o.Protocol.Ospf.GlobalBfd.Profile
					}
				}
				if o.Protocol.Ospf.RouterId != nil {
					nestedProtocol.Ospf.RouterId = o.Protocol.Ospf.RouterId
				}
				if o.Protocol.Ospf.Timers != nil {
					nestedProtocol.Ospf.Timers = &ProtocolOspfTimers{}
					if o.Protocol.Ospf.Timers.Misc != nil {
						entry.Misc["ProtocolOspfTimers"] = o.Protocol.Ospf.Timers.Misc
					}
					if o.Protocol.Ospf.Timers.LsaInterval != nil {
						nestedProtocol.Ospf.Timers.LsaInterval = o.Protocol.Ospf.Timers.LsaInterval
					}
					if o.Protocol.Ospf.Timers.SpfCalculationDelay != nil {
						nestedProtocol.Ospf.Timers.SpfCalculationDelay = o.Protocol.Ospf.Timers.SpfCalculationDelay
					}
				}
				if o.Protocol.Ospf.AllowRedistDefaultRoute != nil {
					nestedProtocol.Ospf.AllowRedistDefaultRoute = util.AsBool(o.Protocol.Ospf.AllowRedistDefaultRoute, nil)
				}
				if o.Protocol.Ospf.Enable != nil {
					nestedProtocol.Ospf.Enable = util.AsBool(o.Protocol.Ospf.Enable, nil)
				}
				if o.Protocol.Ospf.GracefulRestart != nil {
					nestedProtocol.Ospf.GracefulRestart = &ProtocolOspfGracefulRestart{}
					if o.Protocol.Ospf.GracefulRestart.Misc != nil {
						entry.Misc["ProtocolOspfGracefulRestart"] = o.Protocol.Ospf.GracefulRestart.Misc
					}
					if o.Protocol.Ospf.GracefulRestart.Enable != nil {
						nestedProtocol.Ospf.GracefulRestart.Enable = util.AsBool(o.Protocol.Ospf.GracefulRestart.Enable, nil)
					}
					if o.Protocol.Ospf.GracefulRestart.GracePeriod != nil {
						nestedProtocol.Ospf.GracefulRestart.GracePeriod = o.Protocol.Ospf.GracefulRestart.GracePeriod
					}
					if o.Protocol.Ospf.GracefulRestart.HelperEnable != nil {
						nestedProtocol.Ospf.GracefulRestart.HelperEnable = util.AsBool(o.Protocol.Ospf.GracefulRestart.HelperEnable, nil)
					}
					if o.Protocol.Ospf.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedProtocol.Ospf.GracefulRestart.MaxNeighborRestartTime = o.Protocol.Ospf.GracefulRestart.MaxNeighborRestartTime
					}
					if o.Protocol.Ospf.GracefulRestart.StrictLSAChecking != nil {
						nestedProtocol.Ospf.GracefulRestart.StrictLSAChecking = util.AsBool(o.Protocol.Ospf.GracefulRestart.StrictLSAChecking, nil)
					}
				}
				if o.Protocol.Ospf.RejectDefaultRoute != nil {
					nestedProtocol.Ospf.RejectDefaultRoute = util.AsBool(o.Protocol.Ospf.RejectDefaultRoute, nil)
				}
				if o.Protocol.Ospf.Rfc1583 != nil {
					nestedProtocol.Ospf.Rfc1583 = util.AsBool(o.Protocol.Ospf.Rfc1583, nil)
				}
			}
			if o.Protocol.Ospfv3 != nil {
				nestedProtocol.Ospfv3 = &ProtocolOspfv3{}
				if o.Protocol.Ospfv3.Misc != nil {
					entry.Misc["ProtocolOspfv3"] = o.Protocol.Ospfv3.Misc
				}
				if o.Protocol.Ospfv3.Area != nil {
					nestedProtocol.Ospfv3.Area = []ProtocolOspfv3Area{}
					for _, oProtocolOspfv3Area := range o.Protocol.Ospfv3.Area {
						nestedProtocolOspfv3Area := ProtocolOspfv3Area{}
						if oProtocolOspfv3Area.Misc != nil {
							entry.Misc["ProtocolOspfv3Area"] = oProtocolOspfv3Area.Misc
						}
						if oProtocolOspfv3Area.Authentication != nil {
							nestedProtocolOspfv3Area.Authentication = oProtocolOspfv3Area.Authentication
						}
						if oProtocolOspfv3Area.Type != nil {
							nestedProtocolOspfv3Area.Type = &ProtocolOspfv3AreaType{}
							if oProtocolOspfv3Area.Type.Misc != nil {
								entry.Misc["ProtocolOspfv3AreaType"] = oProtocolOspfv3Area.Type.Misc
							}
							if oProtocolOspfv3Area.Type.Normal != nil {
								nestedProtocolOspfv3Area.Type.Normal = &ProtocolOspfv3AreaTypeNormal{}
								if oProtocolOspfv3Area.Type.Normal.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaTypeNormal"] = oProtocolOspfv3Area.Type.Normal.Misc
								}
							}
							if oProtocolOspfv3Area.Type.Stub != nil {
								nestedProtocolOspfv3Area.Type.Stub = &ProtocolOspfv3AreaTypeStub{}
								if oProtocolOspfv3Area.Type.Stub.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaTypeStub"] = oProtocolOspfv3Area.Type.Stub.Misc
								}
								if oProtocolOspfv3Area.Type.Stub.AcceptSummary != nil {
									nestedProtocolOspfv3Area.Type.Stub.AcceptSummary = util.AsBool(oProtocolOspfv3Area.Type.Stub.AcceptSummary, nil)
								}
								if oProtocolOspfv3Area.Type.Stub.DefaultRoute != nil {
									nestedProtocolOspfv3Area.Type.Stub.DefaultRoute = &ProtocolOspfv3AreaTypeStubDefaultRoute{}
									if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaTypeStubDefaultRoute"] = oProtocolOspfv3Area.Type.Stub.DefaultRoute.Misc
									}
									if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise != nil {
										nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise = &ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise{}
										if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise"] = oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Misc
										}
										if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric != nil {
											nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric = oProtocolOspfv3Area.Type.Stub.DefaultRoute.Advertise.Metric
										}
									}
									if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable != nil {
										nestedProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable = &ProtocolOspfv3AreaTypeStubDefaultRouteDisable{}
										if oProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaTypeStubDefaultRouteDisable"] = oProtocolOspfv3Area.Type.Stub.DefaultRoute.Disable.Misc
										}
									}
								}
							}
							if oProtocolOspfv3Area.Type.Nssa != nil {
								nestedProtocolOspfv3Area.Type.Nssa = &ProtocolOspfv3AreaTypeNssa{}
								if oProtocolOspfv3Area.Type.Nssa.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaTypeNssa"] = oProtocolOspfv3Area.Type.Nssa.Misc
								}
								if oProtocolOspfv3Area.Type.Nssa.AcceptSummary != nil {
									nestedProtocolOspfv3Area.Type.Nssa.AcceptSummary = util.AsBool(oProtocolOspfv3Area.Type.Nssa.AcceptSummary, nil)
								}
								if oProtocolOspfv3Area.Type.Nssa.DefaultRoute != nil {
									nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute = &ProtocolOspfv3AreaTypeNssaDefaultRoute{}
									if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaTypeNssaDefaultRoute"] = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Misc
									}
									if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable != nil {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable = &ProtocolOspfv3AreaTypeNssaDefaultRouteDisable{}
										if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteDisable"] = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Disable.Misc
										}
									}
									if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise != nil {
										nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise = &ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise{}
										if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise"] = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Misc
										}
										if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric != nil {
											nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Metric
										}
										if oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type != nil {
											nestedProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type = oProtocolOspfv3Area.Type.Nssa.DefaultRoute.Advertise.Type
										}
									}
								}
								if oProtocolOspfv3Area.Type.Nssa.NssaExtRange != nil {
									nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange = []ProtocolOspfv3AreaTypeNssaNssaExtRange{}
									for _, oProtocolOspfv3AreaTypeNssaNssaExtRange := range oProtocolOspfv3Area.Type.Nssa.NssaExtRange {
										nestedProtocolOspfv3AreaTypeNssaNssaExtRange := ProtocolOspfv3AreaTypeNssaNssaExtRange{}
										if oProtocolOspfv3AreaTypeNssaNssaExtRange.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRange"] = oProtocolOspfv3AreaTypeNssaNssaExtRange.Misc
										}
										if oProtocolOspfv3AreaTypeNssaNssaExtRange.Name != "" {
											nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Name = oProtocolOspfv3AreaTypeNssaNssaExtRange.Name
										}
										if oProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise != nil {
											nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise = &ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise{}
											if oProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise.Misc != nil {
												entry.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise"] = oProtocolOspfv3AreaTypeNssaNssaExtRange.Advertise.Misc
											}
										}
										if oProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress != nil {
											nestedProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress = &ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress{}
											if oProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress.Misc != nil {
												entry.Misc["ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress"] = oProtocolOspfv3AreaTypeNssaNssaExtRange.Suppress.Misc
											}
										}
										nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange = append(nestedProtocolOspfv3Area.Type.Nssa.NssaExtRange, nestedProtocolOspfv3AreaTypeNssaNssaExtRange)
									}
								}
							}
						}
						if oProtocolOspfv3Area.Range != nil {
							nestedProtocolOspfv3Area.Range = []ProtocolOspfv3AreaRange{}
							for _, oProtocolOspfv3AreaRange := range oProtocolOspfv3Area.Range {
								nestedProtocolOspfv3AreaRange := ProtocolOspfv3AreaRange{}
								if oProtocolOspfv3AreaRange.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaRange"] = oProtocolOspfv3AreaRange.Misc
								}
								if oProtocolOspfv3AreaRange.Name != "" {
									nestedProtocolOspfv3AreaRange.Name = oProtocolOspfv3AreaRange.Name
								}
								if oProtocolOspfv3AreaRange.Advertise != nil {
									nestedProtocolOspfv3AreaRange.Advertise = &ProtocolOspfv3AreaRangeAdvertise{}
									if oProtocolOspfv3AreaRange.Advertise.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaRangeAdvertise"] = oProtocolOspfv3AreaRange.Advertise.Misc
									}
								}
								if oProtocolOspfv3AreaRange.Suppress != nil {
									nestedProtocolOspfv3AreaRange.Suppress = &ProtocolOspfv3AreaRangeSuppress{}
									if oProtocolOspfv3AreaRange.Suppress.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaRangeSuppress"] = oProtocolOspfv3AreaRange.Suppress.Misc
									}
								}
								nestedProtocolOspfv3Area.Range = append(nestedProtocolOspfv3Area.Range, nestedProtocolOspfv3AreaRange)
							}
						}
						if oProtocolOspfv3Area.Interface != nil {
							nestedProtocolOspfv3Area.Interface = []ProtocolOspfv3AreaInterface{}
							for _, oProtocolOspfv3AreaInterface := range oProtocolOspfv3Area.Interface {
								nestedProtocolOspfv3AreaInterface := ProtocolOspfv3AreaInterface{}
								if oProtocolOspfv3AreaInterface.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaInterface"] = oProtocolOspfv3AreaInterface.Misc
								}
								if oProtocolOspfv3AreaInterface.Priority != nil {
									nestedProtocolOspfv3AreaInterface.Priority = oProtocolOspfv3AreaInterface.Priority
								}
								if oProtocolOspfv3AreaInterface.RetransmitInterval != nil {
									nestedProtocolOspfv3AreaInterface.RetransmitInterval = oProtocolOspfv3AreaInterface.RetransmitInterval
								}
								if oProtocolOspfv3AreaInterface.Authentication != nil {
									nestedProtocolOspfv3AreaInterface.Authentication = oProtocolOspfv3AreaInterface.Authentication
								}
								if oProtocolOspfv3AreaInterface.Neighbor != nil {
									nestedProtocolOspfv3AreaInterface.Neighbor = []ProtocolOspfv3AreaInterfaceNeighbor{}
									for _, oProtocolOspfv3AreaInterfaceNeighbor := range oProtocolOspfv3AreaInterface.Neighbor {
										nestedProtocolOspfv3AreaInterfaceNeighbor := ProtocolOspfv3AreaInterfaceNeighbor{}
										if oProtocolOspfv3AreaInterfaceNeighbor.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaInterfaceNeighbor"] = oProtocolOspfv3AreaInterfaceNeighbor.Misc
										}
										if oProtocolOspfv3AreaInterfaceNeighbor.Name != "" {
											nestedProtocolOspfv3AreaInterfaceNeighbor.Name = oProtocolOspfv3AreaInterfaceNeighbor.Name
										}
										nestedProtocolOspfv3AreaInterface.Neighbor = append(nestedProtocolOspfv3AreaInterface.Neighbor, nestedProtocolOspfv3AreaInterfaceNeighbor)
									}
								}
								if oProtocolOspfv3AreaInterface.Bfd != nil {
									nestedProtocolOspfv3AreaInterface.Bfd = &ProtocolOspfv3AreaInterfaceBfd{}
									if oProtocolOspfv3AreaInterface.Bfd.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaInterfaceBfd"] = oProtocolOspfv3AreaInterface.Bfd.Misc
									}
									if oProtocolOspfv3AreaInterface.Bfd.Profile != nil {
										nestedProtocolOspfv3AreaInterface.Bfd.Profile = oProtocolOspfv3AreaInterface.Bfd.Profile
									}
								}
								if oProtocolOspfv3AreaInterface.DeadCounts != nil {
									nestedProtocolOspfv3AreaInterface.DeadCounts = oProtocolOspfv3AreaInterface.DeadCounts
								}
								if oProtocolOspfv3AreaInterface.Passive != nil {
									nestedProtocolOspfv3AreaInterface.Passive = util.AsBool(oProtocolOspfv3AreaInterface.Passive, nil)
								}
								if oProtocolOspfv3AreaInterface.GrDelay != nil {
									nestedProtocolOspfv3AreaInterface.GrDelay = oProtocolOspfv3AreaInterface.GrDelay
								}
								if oProtocolOspfv3AreaInterface.LinkType != nil {
									nestedProtocolOspfv3AreaInterface.LinkType = &ProtocolOspfv3AreaInterfaceLinkType{}
									if oProtocolOspfv3AreaInterface.LinkType.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaInterfaceLinkType"] = oProtocolOspfv3AreaInterface.LinkType.Misc
									}
									if oProtocolOspfv3AreaInterface.LinkType.Broadcast != nil {
										nestedProtocolOspfv3AreaInterface.LinkType.Broadcast = &ProtocolOspfv3AreaInterfaceLinkTypeBroadcast{}
										if oProtocolOspfv3AreaInterface.LinkType.Broadcast.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaInterfaceLinkTypeBroadcast"] = oProtocolOspfv3AreaInterface.LinkType.Broadcast.Misc
										}
									}
									if oProtocolOspfv3AreaInterface.LinkType.P2p != nil {
										nestedProtocolOspfv3AreaInterface.LinkType.P2p = &ProtocolOspfv3AreaInterfaceLinkTypeP2p{}
										if oProtocolOspfv3AreaInterface.LinkType.P2p.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2p"] = oProtocolOspfv3AreaInterface.LinkType.P2p.Misc
										}
									}
									if oProtocolOspfv3AreaInterface.LinkType.P2mp != nil {
										nestedProtocolOspfv3AreaInterface.LinkType.P2mp = &ProtocolOspfv3AreaInterfaceLinkTypeP2mp{}
										if oProtocolOspfv3AreaInterface.LinkType.P2mp.Misc != nil {
											entry.Misc["ProtocolOspfv3AreaInterfaceLinkTypeP2mp"] = oProtocolOspfv3AreaInterface.LinkType.P2mp.Misc
										}
									}
								}
								if oProtocolOspfv3AreaInterface.Enable != nil {
									nestedProtocolOspfv3AreaInterface.Enable = util.AsBool(oProtocolOspfv3AreaInterface.Enable, nil)
								}
								if oProtocolOspfv3AreaInterface.Metric != nil {
									nestedProtocolOspfv3AreaInterface.Metric = oProtocolOspfv3AreaInterface.Metric
								}
								if oProtocolOspfv3AreaInterface.HelloInterval != nil {
									nestedProtocolOspfv3AreaInterface.HelloInterval = oProtocolOspfv3AreaInterface.HelloInterval
								}
								if oProtocolOspfv3AreaInterface.TransitDelay != nil {
									nestedProtocolOspfv3AreaInterface.TransitDelay = oProtocolOspfv3AreaInterface.TransitDelay
								}
								if oProtocolOspfv3AreaInterface.Name != "" {
									nestedProtocolOspfv3AreaInterface.Name = oProtocolOspfv3AreaInterface.Name
								}
								if oProtocolOspfv3AreaInterface.InstanceId != nil {
									nestedProtocolOspfv3AreaInterface.InstanceId = oProtocolOspfv3AreaInterface.InstanceId
								}
								nestedProtocolOspfv3Area.Interface = append(nestedProtocolOspfv3Area.Interface, nestedProtocolOspfv3AreaInterface)
							}
						}
						if oProtocolOspfv3Area.VirtualLink != nil {
							nestedProtocolOspfv3Area.VirtualLink = []ProtocolOspfv3AreaVirtualLink{}
							for _, oProtocolOspfv3AreaVirtualLink := range oProtocolOspfv3Area.VirtualLink {
								nestedProtocolOspfv3AreaVirtualLink := ProtocolOspfv3AreaVirtualLink{}
								if oProtocolOspfv3AreaVirtualLink.Misc != nil {
									entry.Misc["ProtocolOspfv3AreaVirtualLink"] = oProtocolOspfv3AreaVirtualLink.Misc
								}
								if oProtocolOspfv3AreaVirtualLink.TransitDelay != nil {
									nestedProtocolOspfv3AreaVirtualLink.TransitDelay = oProtocolOspfv3AreaVirtualLink.TransitDelay
								}
								if oProtocolOspfv3AreaVirtualLink.Bfd != nil {
									nestedProtocolOspfv3AreaVirtualLink.Bfd = &ProtocolOspfv3AreaVirtualLinkBfd{}
									if oProtocolOspfv3AreaVirtualLink.Bfd.Misc != nil {
										entry.Misc["ProtocolOspfv3AreaVirtualLinkBfd"] = oProtocolOspfv3AreaVirtualLink.Bfd.Misc
									}
									if oProtocolOspfv3AreaVirtualLink.Bfd.Profile != nil {
										nestedProtocolOspfv3AreaVirtualLink.Bfd.Profile = oProtocolOspfv3AreaVirtualLink.Bfd.Profile
									}
								}
								if oProtocolOspfv3AreaVirtualLink.NeighborId != nil {
									nestedProtocolOspfv3AreaVirtualLink.NeighborId = oProtocolOspfv3AreaVirtualLink.NeighborId
								}
								if oProtocolOspfv3AreaVirtualLink.Enable != nil {
									nestedProtocolOspfv3AreaVirtualLink.Enable = util.AsBool(oProtocolOspfv3AreaVirtualLink.Enable, nil)
								}
								if oProtocolOspfv3AreaVirtualLink.HelloInterval != nil {
									nestedProtocolOspfv3AreaVirtualLink.HelloInterval = oProtocolOspfv3AreaVirtualLink.HelloInterval
								}
								if oProtocolOspfv3AreaVirtualLink.RetransmitInterval != nil {
									nestedProtocolOspfv3AreaVirtualLink.RetransmitInterval = oProtocolOspfv3AreaVirtualLink.RetransmitInterval
								}
								if oProtocolOspfv3AreaVirtualLink.Authentication != nil {
									nestedProtocolOspfv3AreaVirtualLink.Authentication = oProtocolOspfv3AreaVirtualLink.Authentication
								}
								if oProtocolOspfv3AreaVirtualLink.Name != "" {
									nestedProtocolOspfv3AreaVirtualLink.Name = oProtocolOspfv3AreaVirtualLink.Name
								}
								if oProtocolOspfv3AreaVirtualLink.TransitAreaId != nil {
									nestedProtocolOspfv3AreaVirtualLink.TransitAreaId = oProtocolOspfv3AreaVirtualLink.TransitAreaId
								}
								if oProtocolOspfv3AreaVirtualLink.InstanceId != nil {
									nestedProtocolOspfv3AreaVirtualLink.InstanceId = oProtocolOspfv3AreaVirtualLink.InstanceId
								}
								if oProtocolOspfv3AreaVirtualLink.DeadCounts != nil {
									nestedProtocolOspfv3AreaVirtualLink.DeadCounts = oProtocolOspfv3AreaVirtualLink.DeadCounts
								}
								nestedProtocolOspfv3Area.VirtualLink = append(nestedProtocolOspfv3Area.VirtualLink, nestedProtocolOspfv3AreaVirtualLink)
							}
						}
						if oProtocolOspfv3Area.Name != "" {
							nestedProtocolOspfv3Area.Name = oProtocolOspfv3Area.Name
						}
						nestedProtocol.Ospfv3.Area = append(nestedProtocol.Ospfv3.Area, nestedProtocolOspfv3Area)
					}
				}
				if o.Protocol.Ospfv3.DisableTransitTraffic != nil {
					nestedProtocol.Ospfv3.DisableTransitTraffic = util.AsBool(o.Protocol.Ospfv3.DisableTransitTraffic, nil)
				}
				if o.Protocol.Ospfv3.Enable != nil {
					nestedProtocol.Ospfv3.Enable = util.AsBool(o.Protocol.Ospfv3.Enable, nil)
				}
				if o.Protocol.Ospfv3.ExportRules != nil {
					nestedProtocol.Ospfv3.ExportRules = []ProtocolOspfv3ExportRules{}
					for _, oProtocolOspfv3ExportRules := range o.Protocol.Ospfv3.ExportRules {
						nestedProtocolOspfv3ExportRules := ProtocolOspfv3ExportRules{}
						if oProtocolOspfv3ExportRules.Misc != nil {
							entry.Misc["ProtocolOspfv3ExportRules"] = oProtocolOspfv3ExportRules.Misc
						}
						if oProtocolOspfv3ExportRules.NewPathType != nil {
							nestedProtocolOspfv3ExportRules.NewPathType = oProtocolOspfv3ExportRules.NewPathType
						}
						if oProtocolOspfv3ExportRules.NewTag != nil {
							nestedProtocolOspfv3ExportRules.NewTag = oProtocolOspfv3ExportRules.NewTag
						}
						if oProtocolOspfv3ExportRules.Metric != nil {
							nestedProtocolOspfv3ExportRules.Metric = oProtocolOspfv3ExportRules.Metric
						}
						if oProtocolOspfv3ExportRules.Name != "" {
							nestedProtocolOspfv3ExportRules.Name = oProtocolOspfv3ExportRules.Name
						}
						nestedProtocol.Ospfv3.ExportRules = append(nestedProtocol.Ospfv3.ExportRules, nestedProtocolOspfv3ExportRules)
					}
				}
				if o.Protocol.Ospfv3.GlobalBfd != nil {
					nestedProtocol.Ospfv3.GlobalBfd = &ProtocolOspfv3GlobalBfd{}
					if o.Protocol.Ospfv3.GlobalBfd.Misc != nil {
						entry.Misc["ProtocolOspfv3GlobalBfd"] = o.Protocol.Ospfv3.GlobalBfd.Misc
					}
					if o.Protocol.Ospfv3.GlobalBfd.Profile != nil {
						nestedProtocol.Ospfv3.GlobalBfd.Profile = o.Protocol.Ospfv3.GlobalBfd.Profile
					}
				}
				if o.Protocol.Ospfv3.RejectDefaultRoute != nil {
					nestedProtocol.Ospfv3.RejectDefaultRoute = util.AsBool(o.Protocol.Ospfv3.RejectDefaultRoute, nil)
				}
				if o.Protocol.Ospfv3.AllowRedistDefaultRoute != nil {
					nestedProtocol.Ospfv3.AllowRedistDefaultRoute = util.AsBool(o.Protocol.Ospfv3.AllowRedistDefaultRoute, nil)
				}
				if o.Protocol.Ospfv3.AuthProfile != nil {
					nestedProtocol.Ospfv3.AuthProfile = []ProtocolOspfv3AuthProfile{}
					for _, oProtocolOspfv3AuthProfile := range o.Protocol.Ospfv3.AuthProfile {
						nestedProtocolOspfv3AuthProfile := ProtocolOspfv3AuthProfile{}
						if oProtocolOspfv3AuthProfile.Misc != nil {
							entry.Misc["ProtocolOspfv3AuthProfile"] = oProtocolOspfv3AuthProfile.Misc
						}
						if oProtocolOspfv3AuthProfile.Spi != nil {
							nestedProtocolOspfv3AuthProfile.Spi = oProtocolOspfv3AuthProfile.Spi
						}
						if oProtocolOspfv3AuthProfile.Name != "" {
							nestedProtocolOspfv3AuthProfile.Name = oProtocolOspfv3AuthProfile.Name
						}
						if oProtocolOspfv3AuthProfile.Ah != nil {
							nestedProtocolOspfv3AuthProfile.Ah = &ProtocolOspfv3AuthProfileAh{}
							if oProtocolOspfv3AuthProfile.Ah.Misc != nil {
								entry.Misc["ProtocolOspfv3AuthProfileAh"] = oProtocolOspfv3AuthProfile.Ah.Misc
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha1 != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha1 = &ProtocolOspfv3AuthProfileAhSha1{}
								if oProtocolOspfv3AuthProfile.Ah.Sha1.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileAhSha1"] = oProtocolOspfv3AuthProfile.Ah.Sha1.Misc
								}
								if oProtocolOspfv3AuthProfile.Ah.Sha1.Key != nil {
									nestedProtocolOspfv3AuthProfile.Ah.Sha1.Key = oProtocolOspfv3AuthProfile.Ah.Sha1.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha256 != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha256 = &ProtocolOspfv3AuthProfileAhSha256{}
								if oProtocolOspfv3AuthProfile.Ah.Sha256.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileAhSha256"] = oProtocolOspfv3AuthProfile.Ah.Sha256.Misc
								}
								if oProtocolOspfv3AuthProfile.Ah.Sha256.Key != nil {
									nestedProtocolOspfv3AuthProfile.Ah.Sha256.Key = oProtocolOspfv3AuthProfile.Ah.Sha256.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha384 != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha384 = &ProtocolOspfv3AuthProfileAhSha384{}
								if oProtocolOspfv3AuthProfile.Ah.Sha384.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileAhSha384"] = oProtocolOspfv3AuthProfile.Ah.Sha384.Misc
								}
								if oProtocolOspfv3AuthProfile.Ah.Sha384.Key != nil {
									nestedProtocolOspfv3AuthProfile.Ah.Sha384.Key = oProtocolOspfv3AuthProfile.Ah.Sha384.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Ah.Sha512 != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Sha512 = &ProtocolOspfv3AuthProfileAhSha512{}
								if oProtocolOspfv3AuthProfile.Ah.Sha512.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileAhSha512"] = oProtocolOspfv3AuthProfile.Ah.Sha512.Misc
								}
								if oProtocolOspfv3AuthProfile.Ah.Sha512.Key != nil {
									nestedProtocolOspfv3AuthProfile.Ah.Sha512.Key = oProtocolOspfv3AuthProfile.Ah.Sha512.Key
								}
							}
							if oProtocolOspfv3AuthProfile.Ah.Md5 != nil {
								nestedProtocolOspfv3AuthProfile.Ah.Md5 = &ProtocolOspfv3AuthProfileAhMd5{}
								if oProtocolOspfv3AuthProfile.Ah.Md5.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileAhMd5"] = oProtocolOspfv3AuthProfile.Ah.Md5.Misc
								}
								if oProtocolOspfv3AuthProfile.Ah.Md5.Key != nil {
									nestedProtocolOspfv3AuthProfile.Ah.Md5.Key = oProtocolOspfv3AuthProfile.Ah.Md5.Key
								}
							}
						}
						if oProtocolOspfv3AuthProfile.Esp != nil {
							nestedProtocolOspfv3AuthProfile.Esp = &ProtocolOspfv3AuthProfileEsp{}
							if oProtocolOspfv3AuthProfile.Esp.Misc != nil {
								entry.Misc["ProtocolOspfv3AuthProfileEsp"] = oProtocolOspfv3AuthProfile.Esp.Misc
							}
							if oProtocolOspfv3AuthProfile.Esp.Authentication != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Authentication = &ProtocolOspfv3AuthProfileEspAuthentication{}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileEspAuthentication"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Misc
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.None != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.None = &ProtocolOspfv3AuthProfileEspAuthenticationNone{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.None.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationNone"] = oProtocolOspfv3AuthProfile.Esp.Authentication.None.Misc
									}
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Md5 != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Md5 = &ProtocolOspfv3AuthProfileEspAuthenticationMd5{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationMd5"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Misc
									}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key != nil {
										nestedProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Md5.Key
									}
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1 != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha1 = &ProtocolOspfv3AuthProfileEspAuthenticationSha1{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha1"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Misc
									}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key != nil {
										nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha1.Key
									}
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256 != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha256 = &ProtocolOspfv3AuthProfileEspAuthenticationSha256{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha256"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Misc
									}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key != nil {
										nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha256.Key
									}
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384 != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha384 = &ProtocolOspfv3AuthProfileEspAuthenticationSha384{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha384"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Misc
									}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key != nil {
										nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha384.Key
									}
								}
								if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512 != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha512 = &ProtocolOspfv3AuthProfileEspAuthenticationSha512{}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Misc != nil {
										entry.Misc["ProtocolOspfv3AuthProfileEspAuthenticationSha512"] = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Misc
									}
									if oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key != nil {
										nestedProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key = oProtocolOspfv3AuthProfile.Esp.Authentication.Sha512.Key
									}
								}
							}
							if oProtocolOspfv3AuthProfile.Esp.Encryption != nil {
								nestedProtocolOspfv3AuthProfile.Esp.Encryption = &ProtocolOspfv3AuthProfileEspEncryption{}
								if oProtocolOspfv3AuthProfile.Esp.Encryption.Misc != nil {
									entry.Misc["ProtocolOspfv3AuthProfileEspEncryption"] = oProtocolOspfv3AuthProfile.Esp.Encryption.Misc
								}
								if oProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm = oProtocolOspfv3AuthProfile.Esp.Encryption.Algorithm
								}
								if oProtocolOspfv3AuthProfile.Esp.Encryption.Key != nil {
									nestedProtocolOspfv3AuthProfile.Esp.Encryption.Key = oProtocolOspfv3AuthProfile.Esp.Encryption.Key
								}
							}
						}
						nestedProtocol.Ospfv3.AuthProfile = append(nestedProtocol.Ospfv3.AuthProfile, nestedProtocolOspfv3AuthProfile)
					}
				}
				if o.Protocol.Ospfv3.GracefulRestart != nil {
					nestedProtocol.Ospfv3.GracefulRestart = &ProtocolOspfv3GracefulRestart{}
					if o.Protocol.Ospfv3.GracefulRestart.Misc != nil {
						entry.Misc["ProtocolOspfv3GracefulRestart"] = o.Protocol.Ospfv3.GracefulRestart.Misc
					}
					if o.Protocol.Ospfv3.GracefulRestart.Enable != nil {
						nestedProtocol.Ospfv3.GracefulRestart.Enable = util.AsBool(o.Protocol.Ospfv3.GracefulRestart.Enable, nil)
					}
					if o.Protocol.Ospfv3.GracefulRestart.GracePeriod != nil {
						nestedProtocol.Ospfv3.GracefulRestart.GracePeriod = o.Protocol.Ospfv3.GracefulRestart.GracePeriod
					}
					if o.Protocol.Ospfv3.GracefulRestart.HelperEnable != nil {
						nestedProtocol.Ospfv3.GracefulRestart.HelperEnable = util.AsBool(o.Protocol.Ospfv3.GracefulRestart.HelperEnable, nil)
					}
					if o.Protocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime != nil {
						nestedProtocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime = o.Protocol.Ospfv3.GracefulRestart.MaxNeighborRestartTime
					}
					if o.Protocol.Ospfv3.GracefulRestart.StrictLSAChecking != nil {
						nestedProtocol.Ospfv3.GracefulRestart.StrictLSAChecking = util.AsBool(o.Protocol.Ospfv3.GracefulRestart.StrictLSAChecking, nil)
					}
				}
				if o.Protocol.Ospfv3.RouterId != nil {
					nestedProtocol.Ospfv3.RouterId = o.Protocol.Ospfv3.RouterId
				}
				if o.Protocol.Ospfv3.Timers != nil {
					nestedProtocol.Ospfv3.Timers = &ProtocolOspfv3Timers{}
					if o.Protocol.Ospfv3.Timers.Misc != nil {
						entry.Misc["ProtocolOspfv3Timers"] = o.Protocol.Ospfv3.Timers.Misc
					}
					if o.Protocol.Ospfv3.Timers.LsaInterval != nil {
						nestedProtocol.Ospfv3.Timers.LsaInterval = o.Protocol.Ospfv3.Timers.LsaInterval
					}
					if o.Protocol.Ospfv3.Timers.SpfCalculationDelay != nil {
						nestedProtocol.Ospfv3.Timers.SpfCalculationDelay = o.Protocol.Ospfv3.Timers.SpfCalculationDelay
					}
				}
			}
		}
		entry.Protocol = nestedProtocol

		var nestedRoutingTable *RoutingTable
		if o.RoutingTable != nil {
			nestedRoutingTable = &RoutingTable{}
			if o.RoutingTable.Misc != nil {
				entry.Misc["RoutingTable"] = o.RoutingTable.Misc
			}
			if o.RoutingTable.Ipv6 != nil {
				nestedRoutingTable.Ipv6 = &RoutingTableIpv6{}
				if o.RoutingTable.Ipv6.Misc != nil {
					entry.Misc["RoutingTableIpv6"] = o.RoutingTable.Ipv6.Misc
				}
				if o.RoutingTable.Ipv6.StaticRoute != nil {
					nestedRoutingTable.Ipv6.StaticRoute = []RoutingTableIpv6StaticRoute{}
					for _, oRoutingTableIpv6StaticRoute := range o.RoutingTable.Ipv6.StaticRoute {
						nestedRoutingTableIpv6StaticRoute := RoutingTableIpv6StaticRoute{}
						if oRoutingTableIpv6StaticRoute.Misc != nil {
							entry.Misc["RoutingTableIpv6StaticRoute"] = oRoutingTableIpv6StaticRoute.Misc
						}
						if oRoutingTableIpv6StaticRoute.AdminDist != nil {
							nestedRoutingTableIpv6StaticRoute.AdminDist = oRoutingTableIpv6StaticRoute.AdminDist
						}
						if oRoutingTableIpv6StaticRoute.Metric != nil {
							nestedRoutingTableIpv6StaticRoute.Metric = oRoutingTableIpv6StaticRoute.Metric
						}
						if oRoutingTableIpv6StaticRoute.Nexthop != nil {
							nestedRoutingTableIpv6StaticRoute.Nexthop = &RoutingTableIpv6StaticRouteNexthop{}
							if oRoutingTableIpv6StaticRoute.Nexthop.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRouteNexthop"] = oRoutingTableIpv6StaticRoute.Nexthop.Misc
							}
							if oRoutingTableIpv6StaticRoute.Nexthop.Receive != nil {
								nestedRoutingTableIpv6StaticRoute.Nexthop.Receive = &RoutingTableIpv6StaticRouteNexthopReceive{}
								if oRoutingTableIpv6StaticRoute.Nexthop.Receive.Misc != nil {
									entry.Misc["RoutingTableIpv6StaticRouteNexthopReceive"] = oRoutingTableIpv6StaticRoute.Nexthop.Receive.Misc
								}
							}
							if oRoutingTableIpv6StaticRoute.Nexthop.Discard != nil {
								nestedRoutingTableIpv6StaticRoute.Nexthop.Discard = &RoutingTableIpv6StaticRouteNexthopDiscard{}
								if oRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc != nil {
									entry.Misc["RoutingTableIpv6StaticRouteNexthopDiscard"] = oRoutingTableIpv6StaticRoute.Nexthop.Discard.Misc
								}
							}
							if oRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address != nil {
								nestedRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address = oRoutingTableIpv6StaticRoute.Nexthop.Ipv6Address
							}
							if oRoutingTableIpv6StaticRoute.Nexthop.NextVr != nil {
								nestedRoutingTableIpv6StaticRoute.Nexthop.NextVr = oRoutingTableIpv6StaticRoute.Nexthop.NextVr
							}
						}
						if oRoutingTableIpv6StaticRoute.Option != nil {
							nestedRoutingTableIpv6StaticRoute.Option = &RoutingTableIpv6StaticRouteOption{}
							if oRoutingTableIpv6StaticRoute.Option.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRouteOption"] = oRoutingTableIpv6StaticRoute.Option.Misc
							}
						}
						if oRoutingTableIpv6StaticRoute.RouteTable != nil {
							nestedRoutingTableIpv6StaticRoute.RouteTable = &RoutingTableIpv6StaticRouteRouteTable{}
							if oRoutingTableIpv6StaticRoute.RouteTable.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRouteRouteTable"] = oRoutingTableIpv6StaticRoute.RouteTable.Misc
							}
							if oRoutingTableIpv6StaticRoute.RouteTable.Unicast != nil {
								nestedRoutingTableIpv6StaticRoute.RouteTable.Unicast = &RoutingTableIpv6StaticRouteRouteTableUnicast{}
								if oRoutingTableIpv6StaticRoute.RouteTable.Unicast.Misc != nil {
									entry.Misc["RoutingTableIpv6StaticRouteRouteTableUnicast"] = oRoutingTableIpv6StaticRoute.RouteTable.Unicast.Misc
								}
							}
							if oRoutingTableIpv6StaticRoute.RouteTable.NoInstall != nil {
								nestedRoutingTableIpv6StaticRoute.RouteTable.NoInstall = &RoutingTableIpv6StaticRouteRouteTableNoInstall{}
								if oRoutingTableIpv6StaticRoute.RouteTable.NoInstall.Misc != nil {
									entry.Misc["RoutingTableIpv6StaticRouteRouteTableNoInstall"] = oRoutingTableIpv6StaticRoute.RouteTable.NoInstall.Misc
								}
							}
						}
						if oRoutingTableIpv6StaticRoute.PathMonitor != nil {
							nestedRoutingTableIpv6StaticRoute.PathMonitor = &RoutingTableIpv6StaticRoutePathMonitor{}
							if oRoutingTableIpv6StaticRoute.PathMonitor.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRoutePathMonitor"] = oRoutingTableIpv6StaticRoute.PathMonitor.Misc
							}
							if oRoutingTableIpv6StaticRoute.PathMonitor.HoldTime != nil {
								nestedRoutingTableIpv6StaticRoute.PathMonitor.HoldTime = oRoutingTableIpv6StaticRoute.PathMonitor.HoldTime
							}
							if oRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations != nil {
								nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
								for _, oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := range oRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations {
									nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations := RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations{}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc != nil {
										entry.Misc["RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations"] = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Misc
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count != nil {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Count
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name != "" {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Name
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable != nil {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Enable, nil)
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source != nil {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Source
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination != nil {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Destination
									}
									if oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval != nil {
										nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval = oRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations.Interval
									}
									nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations = append(nestedRoutingTableIpv6StaticRoute.PathMonitor.MonitorDestinations, nestedRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations)
								}
							}
							if oRoutingTableIpv6StaticRoute.PathMonitor.Enable != nil {
								nestedRoutingTableIpv6StaticRoute.PathMonitor.Enable = util.AsBool(oRoutingTableIpv6StaticRoute.PathMonitor.Enable, nil)
							}
							if oRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition != nil {
								nestedRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition = oRoutingTableIpv6StaticRoute.PathMonitor.FailureCondition
							}
						}
						if oRoutingTableIpv6StaticRoute.Destination != nil {
							nestedRoutingTableIpv6StaticRoute.Destination = oRoutingTableIpv6StaticRoute.Destination
						}
						if oRoutingTableIpv6StaticRoute.Interface != nil {
							nestedRoutingTableIpv6StaticRoute.Interface = oRoutingTableIpv6StaticRoute.Interface
						}
						if oRoutingTableIpv6StaticRoute.Bfd != nil {
							nestedRoutingTableIpv6StaticRoute.Bfd = &RoutingTableIpv6StaticRouteBfd{}
							if oRoutingTableIpv6StaticRoute.Bfd.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRouteBfd"] = oRoutingTableIpv6StaticRoute.Bfd.Misc
							}
							if oRoutingTableIpv6StaticRoute.Bfd.Profile != nil {
								nestedRoutingTableIpv6StaticRoute.Bfd.Profile = oRoutingTableIpv6StaticRoute.Bfd.Profile
							}
						}
						if oRoutingTableIpv6StaticRoute.Name != "" {
							nestedRoutingTableIpv6StaticRoute.Name = oRoutingTableIpv6StaticRoute.Name
						}
						nestedRoutingTable.Ipv6.StaticRoute = append(nestedRoutingTable.Ipv6.StaticRoute, nestedRoutingTableIpv6StaticRoute)
					}
				}
			}
			if o.RoutingTable.Ip != nil {
				nestedRoutingTable.Ip = &RoutingTableIp{}
				if o.RoutingTable.Ip.Misc != nil {
					entry.Misc["RoutingTableIp"] = o.RoutingTable.Ip.Misc
				}
				if o.RoutingTable.Ip.StaticRoute != nil {
					nestedRoutingTable.Ip.StaticRoute = []RoutingTableIpStaticRoute{}
					for _, oRoutingTableIpStaticRoute := range o.RoutingTable.Ip.StaticRoute {
						nestedRoutingTableIpStaticRoute := RoutingTableIpStaticRoute{}
						if oRoutingTableIpStaticRoute.Misc != nil {
							entry.Misc["RoutingTableIpStaticRoute"] = oRoutingTableIpStaticRoute.Misc
						}
						if oRoutingTableIpStaticRoute.PathMonitor != nil {
							nestedRoutingTableIpStaticRoute.PathMonitor = &RoutingTableIpStaticRoutePathMonitor{}
							if oRoutingTableIpStaticRoute.PathMonitor.Misc != nil {
								entry.Misc["RoutingTableIpStaticRoutePathMonitor"] = oRoutingTableIpStaticRoute.PathMonitor.Misc
							}
							if oRoutingTableIpStaticRoute.PathMonitor.FailureCondition != nil {
								nestedRoutingTableIpStaticRoute.PathMonitor.FailureCondition = oRoutingTableIpStaticRoute.PathMonitor.FailureCondition
							}
							if oRoutingTableIpStaticRoute.PathMonitor.HoldTime != nil {
								nestedRoutingTableIpStaticRoute.PathMonitor.HoldTime = oRoutingTableIpStaticRoute.PathMonitor.HoldTime
							}
							if oRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations != nil {
								nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = []RoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
								for _, oRoutingTableIpStaticRoutePathMonitorMonitorDestinations := range oRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations {
									nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations := RoutingTableIpStaticRoutePathMonitorMonitorDestinations{}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc != nil {
										entry.Misc["RoutingTableIpStaticRoutePathMonitorMonitorDestinations"] = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Misc
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source != nil {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Source
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination != nil {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Destination
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval != nil {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Interval
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count != nil {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Count
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name != "" {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name = oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Name
									}
									if oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable != nil {
										nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable = util.AsBool(oRoutingTableIpStaticRoutePathMonitorMonitorDestinations.Enable, nil)
									}
									nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations = append(nestedRoutingTableIpStaticRoute.PathMonitor.MonitorDestinations, nestedRoutingTableIpStaticRoutePathMonitorMonitorDestinations)
								}
							}
							if oRoutingTableIpStaticRoute.PathMonitor.Enable != nil {
								nestedRoutingTableIpStaticRoute.PathMonitor.Enable = util.AsBool(oRoutingTableIpStaticRoute.PathMonitor.Enable, nil)
							}
						}
						if oRoutingTableIpStaticRoute.Name != "" {
							nestedRoutingTableIpStaticRoute.Name = oRoutingTableIpStaticRoute.Name
						}
						if oRoutingTableIpStaticRoute.Destination != nil {
							nestedRoutingTableIpStaticRoute.Destination = oRoutingTableIpStaticRoute.Destination
						}
						if oRoutingTableIpStaticRoute.Metric != nil {
							nestedRoutingTableIpStaticRoute.Metric = oRoutingTableIpStaticRoute.Metric
						}
						if oRoutingTableIpStaticRoute.Nexthop != nil {
							nestedRoutingTableIpStaticRoute.Nexthop = &RoutingTableIpStaticRouteNexthop{}
							if oRoutingTableIpStaticRoute.Nexthop.Misc != nil {
								entry.Misc["RoutingTableIpStaticRouteNexthop"] = oRoutingTableIpStaticRoute.Nexthop.Misc
							}
							if oRoutingTableIpStaticRoute.Nexthop.Receive != nil {
								nestedRoutingTableIpStaticRoute.Nexthop.Receive = &RoutingTableIpStaticRouteNexthopReceive{}
								if oRoutingTableIpStaticRoute.Nexthop.Receive.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteNexthopReceive"] = oRoutingTableIpStaticRoute.Nexthop.Receive.Misc
								}
							}
							if oRoutingTableIpStaticRoute.Nexthop.Discard != nil {
								nestedRoutingTableIpStaticRoute.Nexthop.Discard = &RoutingTableIpStaticRouteNexthopDiscard{}
								if oRoutingTableIpStaticRoute.Nexthop.Discard.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteNexthopDiscard"] = oRoutingTableIpStaticRoute.Nexthop.Discard.Misc
								}
							}
							if oRoutingTableIpStaticRoute.Nexthop.IpAddress != nil {
								nestedRoutingTableIpStaticRoute.Nexthop.IpAddress = oRoutingTableIpStaticRoute.Nexthop.IpAddress
							}
							if oRoutingTableIpStaticRoute.Nexthop.Fqdn != nil {
								nestedRoutingTableIpStaticRoute.Nexthop.Fqdn = oRoutingTableIpStaticRoute.Nexthop.Fqdn
							}
							if oRoutingTableIpStaticRoute.Nexthop.NextVr != nil {
								nestedRoutingTableIpStaticRoute.Nexthop.NextVr = oRoutingTableIpStaticRoute.Nexthop.NextVr
							}
						}
						if oRoutingTableIpStaticRoute.Bfd != nil {
							nestedRoutingTableIpStaticRoute.Bfd = &RoutingTableIpStaticRouteBfd{}
							if oRoutingTableIpStaticRoute.Bfd.Misc != nil {
								entry.Misc["RoutingTableIpStaticRouteBfd"] = oRoutingTableIpStaticRoute.Bfd.Misc
							}
							if oRoutingTableIpStaticRoute.Bfd.Profile != nil {
								nestedRoutingTableIpStaticRoute.Bfd.Profile = oRoutingTableIpStaticRoute.Bfd.Profile
							}
						}
						if oRoutingTableIpStaticRoute.Interface != nil {
							nestedRoutingTableIpStaticRoute.Interface = oRoutingTableIpStaticRoute.Interface
						}
						if oRoutingTableIpStaticRoute.AdminDist != nil {
							nestedRoutingTableIpStaticRoute.AdminDist = oRoutingTableIpStaticRoute.AdminDist
						}
						if oRoutingTableIpStaticRoute.RouteTable != nil {
							nestedRoutingTableIpStaticRoute.RouteTable = &RoutingTableIpStaticRouteRouteTable{}
							if oRoutingTableIpStaticRoute.RouteTable.Misc != nil {
								entry.Misc["RoutingTableIpStaticRouteRouteTable"] = oRoutingTableIpStaticRoute.RouteTable.Misc
							}
							if oRoutingTableIpStaticRoute.RouteTable.Unicast != nil {
								nestedRoutingTableIpStaticRoute.RouteTable.Unicast = &RoutingTableIpStaticRouteRouteTableUnicast{}
								if oRoutingTableIpStaticRoute.RouteTable.Unicast.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteRouteTableUnicast"] = oRoutingTableIpStaticRoute.RouteTable.Unicast.Misc
								}
							}
							if oRoutingTableIpStaticRoute.RouteTable.Multicast != nil {
								nestedRoutingTableIpStaticRoute.RouteTable.Multicast = &RoutingTableIpStaticRouteRouteTableMulticast{}
								if oRoutingTableIpStaticRoute.RouteTable.Multicast.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteRouteTableMulticast"] = oRoutingTableIpStaticRoute.RouteTable.Multicast.Misc
								}
							}
							if oRoutingTableIpStaticRoute.RouteTable.Both != nil {
								nestedRoutingTableIpStaticRoute.RouteTable.Both = &RoutingTableIpStaticRouteRouteTableBoth{}
								if oRoutingTableIpStaticRoute.RouteTable.Both.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteRouteTableBoth"] = oRoutingTableIpStaticRoute.RouteTable.Both.Misc
								}
							}
							if oRoutingTableIpStaticRoute.RouteTable.NoInstall != nil {
								nestedRoutingTableIpStaticRoute.RouteTable.NoInstall = &RoutingTableIpStaticRouteRouteTableNoInstall{}
								if oRoutingTableIpStaticRoute.RouteTable.NoInstall.Misc != nil {
									entry.Misc["RoutingTableIpStaticRouteRouteTableNoInstall"] = oRoutingTableIpStaticRoute.RouteTable.NoInstall.Misc
								}
							}
						}
						nestedRoutingTable.Ip.StaticRoute = append(nestedRoutingTable.Ip.StaticRoute, nestedRoutingTableIpStaticRoute)
					}
				}
			}
		}
		entry.RoutingTable = nestedRoutingTable

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
	if !matchAdminDists(a.AdminDists, b.AdminDists) {
		return false
	}
	if !matchEcmp(a.Ecmp, b.Ecmp) {
		return false
	}
	if !util.OrderedListsMatch(a.Interface, b.Interface) {
		return false
	}
	if !matchMulticast(a.Multicast, b.Multicast) {
		return false
	}
	if !matchProtocol(a.Protocol, b.Protocol) {
		return false
	}
	if !matchRoutingTable(a.RoutingTable, b.RoutingTable) {
		return false
	}

	return true
}

func matchMulticastInterfaceGroupPimAllowedNeighbors(a []MulticastInterfaceGroupPimAllowedNeighbors, b []MulticastInterfaceGroupPimAllowedNeighbors) bool {
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
		}
	}
	return true
}
func matchMulticastInterfaceGroupPim(a *MulticastInterfaceGroupPim, b *MulticastInterfaceGroupPim) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.JoinPruneInterval, b.JoinPruneInterval) {
		return false
	}
	if !util.Ints64Match(a.DrPriority, b.DrPriority) {
		return false
	}
	if !util.BoolsMatch(a.BsrBorder, b.BsrBorder) {
		return false
	}
	if !matchMulticastInterfaceGroupPimAllowedNeighbors(a.AllowedNeighbors, b.AllowedNeighbors) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.AssertInterval, b.AssertInterval) {
		return false
	}
	if !util.Ints64Match(a.HelloInterval, b.HelloInterval) {
		return false
	}
	return true
}
func matchMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast(a []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast, b []MulticastInterfaceGroupGroupPermissionSourceSpecificMulticast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.SourceAddress, b.SourceAddress) {
				return false
			}
			if !util.BoolsMatch(a.Included, b.Included) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.GroupAddress, b.GroupAddress) {
				return false
			}
		}
	}
	return true
}
func matchMulticastInterfaceGroupGroupPermissionAnySourceMulticast(a []MulticastInterfaceGroupGroupPermissionAnySourceMulticast, b []MulticastInterfaceGroupGroupPermissionAnySourceMulticast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.GroupAddress, b.GroupAddress) {
				return false
			}
			if !util.BoolsMatch(a.Included, b.Included) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchMulticastInterfaceGroupGroupPermission(a *MulticastInterfaceGroupGroupPermission, b *MulticastInterfaceGroupGroupPermission) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchMulticastInterfaceGroupGroupPermissionAnySourceMulticast(a.AnySourceMulticast, b.AnySourceMulticast) {
		return false
	}
	if !matchMulticastInterfaceGroupGroupPermissionSourceSpecificMulticast(a.SourceSpecificMulticast, b.SourceSpecificMulticast) {
		return false
	}
	return true
}
func matchMulticastInterfaceGroupIgmp(a *MulticastInterfaceGroupIgmp, b *MulticastInterfaceGroupIgmp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.MaxSources, b.MaxSources) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.QueryInterval, b.QueryInterval) {
		return false
	}
	if !util.FloatsMatch(a.LastMemberQueryInterval, b.LastMemberQueryInterval) {
		return false
	}
	if !util.BoolsMatch(a.ImmediateLeave, b.ImmediateLeave) {
		return false
	}
	if !util.StringsMatch(a.Robustness, b.Robustness) {
		return false
	}
	if !util.StringsMatch(a.Version, b.Version) {
		return false
	}
	if !util.FloatsMatch(a.MaxQueryResponseTime, b.MaxQueryResponseTime) {
		return false
	}
	if !util.StringsMatch(a.MaxGroups, b.MaxGroups) {
		return false
	}
	if !util.BoolsMatch(a.RouterAlertPolicing, b.RouterAlertPolicing) {
		return false
	}
	return true
}
func matchMulticastInterfaceGroup(a []MulticastInterfaceGroup, b []MulticastInterfaceGroup) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchMulticastInterfaceGroupGroupPermission(a.GroupPermission, b.GroupPermission) {
				return false
			}
			if !matchMulticastInterfaceGroupIgmp(a.Igmp, b.Igmp) {
				return false
			}
			if !matchMulticastInterfaceGroupPim(a.Pim, b.Pim) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
			if !util.OrderedListsMatch(a.Interface, b.Interface) {
				return false
			}
		}
	}
	return true
}
func matchMulticastRpExternalRp(a []MulticastRpExternalRp, b []MulticastRpExternalRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.OrderedListsMatch(a.GroupAddresses, b.GroupAddresses) {
				return false
			}
			if !util.BoolsMatch(a.Override, b.Override) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchMulticastRpLocalRpCandidateRp(a *MulticastRpLocalRpCandidateRp, b *MulticastRpLocalRpCandidateRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Address, b.Address) {
		return false
	}
	if !util.Ints64Match(a.AdvertisementInterval, b.AdvertisementInterval) {
		return false
	}
	if !util.OrderedListsMatch(a.GroupAddresses, b.GroupAddresses) {
		return false
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.Ints64Match(a.Priority, b.Priority) {
		return false
	}
	return true
}
func matchMulticastRpLocalRpStaticRp(a *MulticastRpLocalRpStaticRp, b *MulticastRpLocalRpStaticRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Address, b.Address) {
		return false
	}
	if !util.OrderedListsMatch(a.GroupAddresses, b.GroupAddresses) {
		return false
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.BoolsMatch(a.Override, b.Override) {
		return false
	}
	return true
}
func matchMulticastRpLocalRp(a *MulticastRpLocalRp, b *MulticastRpLocalRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchMulticastRpLocalRpCandidateRp(a.CandidateRp, b.CandidateRp) {
		return false
	}
	if !matchMulticastRpLocalRpStaticRp(a.StaticRp, b.StaticRp) {
		return false
	}
	return true
}
func matchMulticastRp(a *MulticastRp, b *MulticastRp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchMulticastRpExternalRp(a.ExternalRp, b.ExternalRp) {
		return false
	}
	if !matchMulticastRpLocalRp(a.LocalRp, b.LocalRp) {
		return false
	}
	return true
}
func matchMulticastSptThreshold(a []MulticastSptThreshold, b []MulticastSptThreshold) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Threshold, b.Threshold) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchMulticastSsmAddressSpace(a []MulticastSsmAddressSpace, b []MulticastSsmAddressSpace) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.GroupAddress, b.GroupAddress) {
				return false
			}
			if !util.BoolsMatch(a.Included, b.Included) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchMulticast(a *Multicast, b *Multicast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchMulticastInterfaceGroup(a.InterfaceGroup, b.InterfaceGroup) {
		return false
	}
	if !util.Ints64Match(a.RouteAgeoutTime, b.RouteAgeoutTime) {
		return false
	}
	if !matchMulticastRp(a.Rp, b.Rp) {
		return false
	}
	if !matchMulticastSptThreshold(a.SptThreshold, b.SptThreshold) {
		return false
	}
	if !matchMulticastSsmAddressSpace(a.SsmAddressSpace, b.SsmAddressSpace) {
		return false
	}
	return true
}
func matchProtocolBgpAuthProfile(a []ProtocolBgpAuthProfile, b []ProtocolBgpAuthProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Secret, b.Secret) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPathRemove(a.Remove, b.Remove) {
		return false
	}
	if !util.Ints64Match(a.Prepend, b.Prepend) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributes(a *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes, b *ProtocolBgpPolicyAggregationAddressAggregateRouteAttributes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Weight, b.Weight) {
		return false
	}
	if !util.StringsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.StringsMatch(a.Origin, b.Origin) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesCommunity(a.Community, b.Community) {
		return false
	}
	if !util.Ints64Match(a.LocalPreference, b.LocalPreference) {
		return false
	}
	if !util.Ints64Match(a.AsPathLimit, b.AsPathLimit) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributesExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity(a *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity, b *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity(a *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity, b *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix(a []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix, b []ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Exact, b.Exact) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath(a *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath, b *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatch(a *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch, b *ProtocolBgpPolicyAggregationAddressSuppressFiltersMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressSuppressFilters(a []ProtocolBgpPolicyAggregationAddressSuppressFilters, b []ProtocolBgpPolicyAggregationAddressSuppressFilters) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !matchProtocolBgpPolicyAggregationAddressSuppressFiltersMatch(a.Match, b.Match) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix(a []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix, b []ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Exact, b.Exact) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath(a *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath, b *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity(a *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity, b *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity(a *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity, b *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch(a *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch, b *ProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddressAdvertiseFilters(a []ProtocolBgpPolicyAggregationAddressAdvertiseFilters, b []ProtocolBgpPolicyAggregationAddressAdvertiseFilters) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !matchProtocolBgpPolicyAggregationAddressAdvertiseFiltersMatch(a.Match, b.Match) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregationAddress(a []ProtocolBgpPolicyAggregationAddress, b []ProtocolBgpPolicyAggregationAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Summary, b.Summary) {
				return false
			}
			if !util.BoolsMatch(a.AsSet, b.AsSet) {
				return false
			}
			if !matchProtocolBgpPolicyAggregationAddressAggregateRouteAttributes(a.AggregateRouteAttributes, b.AggregateRouteAttributes) {
				return false
			}
			if !matchProtocolBgpPolicyAggregationAddressSuppressFilters(a.SuppressFilters, b.SuppressFilters) {
				return false
			}
			if !matchProtocolBgpPolicyAggregationAddressAdvertiseFilters(a.AdvertiseFilters, b.AdvertiseFilters) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Prefix, b.Prefix) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyAggregation(a *ProtocolBgpPolicyAggregation, b *ProtocolBgpPolicyAggregation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregationAddress(a.Address, b.Address) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix(a []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix, b []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix) bool {
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
		}
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath(a *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath, b *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity(a *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity, b *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity(a *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity, b *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch(a *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch, b *ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatchCommunity(a.Community, b.Community) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters(a []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters, b []ProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFiltersMatch(a.Match, b.Match) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix(a []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix, b []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix) bool {
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
		}
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath(a *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath, b *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity(a *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity, b *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity(a *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity, b *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch(a *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch, b *ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters(a []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters, b []ProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFiltersMatch(a.Match, b.Match) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisementPolicy(a []ProtocolBgpPolicyConditionalAdvertisementPolicy, b []ProtocolBgpPolicyConditionalAdvertisementPolicy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.OrderedListsMatch(a.UsedBy, b.UsedBy) {
				return false
			}
			if !matchProtocolBgpPolicyConditionalAdvertisementPolicyNonExistFilters(a.NonExistFilters, b.NonExistFilters) {
				return false
			}
			if !matchProtocolBgpPolicyConditionalAdvertisementPolicyAdvertiseFilters(a.AdvertiseFilters, b.AdvertiseFilters) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyConditionalAdvertisement(a *ProtocolBgpPolicyConditionalAdvertisement, b *ProtocolBgpPolicyConditionalAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyConditionalAdvertisementPolicy(a.Policy, b.Policy) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesMatchCommunity(a *ProtocolBgpPolicyExportRulesMatchCommunity, b *ProtocolBgpPolicyExportRulesMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesMatchExtendedCommunity(a *ProtocolBgpPolicyExportRulesMatchExtendedCommunity, b *ProtocolBgpPolicyExportRulesMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesMatchAddressPrefix(a []ProtocolBgpPolicyExportRulesMatchAddressPrefix, b []ProtocolBgpPolicyExportRulesMatchAddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Exact, b.Exact) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyExportRulesMatchAsPath(a *ProtocolBgpPolicyExportRulesMatchAsPath, b *ProtocolBgpPolicyExportRulesMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesMatch(a *ProtocolBgpPolicyExportRulesMatch, b *ProtocolBgpPolicyExportRulesMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesMatchCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionDeny(a *ProtocolBgpPolicyExportRulesActionDeny, b *ProtocolBgpPolicyExportRulesActionDeny) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone(a *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone, b *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove(a *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove, b *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPath(a *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath, b *ProtocolBgpPolicyExportRulesActionAllowUpdateAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPathNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPathRemove(a.Remove, b.Remove) {
		return false
	}
	if !util.Ints64Match(a.Prepend, b.Prepend) {
		return false
	}
	if !util.Ints64Match(a.RemoveAndPrepend, b.RemoveAndPrepend) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone(a *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone, b *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll(a *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll, b *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunity(a *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity, b *ProtocolBgpPolicyExportRulesActionAllowUpdateCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone(a *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone, b *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll(a *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll, b *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity(a *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity, b *ProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllowUpdate(a *ProtocolBgpPolicyExportRulesActionAllowUpdate, b *ProtocolBgpPolicyExportRulesActionAllowUpdate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.LocalPreference, b.LocalPreference) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !util.StringsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !util.StringsMatch(a.Origin, b.Origin) {
		return false
	}
	if !util.Ints64Match(a.AsPathLimit, b.AsPathLimit) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdateExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesActionAllow(a *ProtocolBgpPolicyExportRulesActionAllow, b *ProtocolBgpPolicyExportRulesActionAllow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesActionAllowUpdate(a.Update, b.Update) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRulesAction(a *ProtocolBgpPolicyExportRulesAction, b *ProtocolBgpPolicyExportRulesAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRulesActionDeny(a.Deny, b.Deny) {
		return false
	}
	if !matchProtocolBgpPolicyExportRulesActionAllow(a.Allow, b.Allow) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyExportRules(a []ProtocolBgpPolicyExportRules, b []ProtocolBgpPolicyExportRules) bool {
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
			if !util.OrderedListsMatch(a.UsedBy, b.UsedBy) {
				return false
			}
			if !matchProtocolBgpPolicyExportRulesMatch(a.Match, b.Match) {
				return false
			}
			if !matchProtocolBgpPolicyExportRulesAction(a.Action, b.Action) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyExport(a *ProtocolBgpPolicyExport, b *ProtocolBgpPolicyExport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyExportRules(a.Rules, b.Rules) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesMatchAddressPrefix(a []ProtocolBgpPolicyImportRulesMatchAddressPrefix, b []ProtocolBgpPolicyImportRulesMatchAddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Exact, b.Exact) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyImportRulesMatchAsPath(a *ProtocolBgpPolicyImportRulesMatchAsPath, b *ProtocolBgpPolicyImportRulesMatchAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesMatchCommunity(a *ProtocolBgpPolicyImportRulesMatchCommunity, b *ProtocolBgpPolicyImportRulesMatchCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesMatchExtendedCommunity(a *ProtocolBgpPolicyImportRulesMatchExtendedCommunity, b *ProtocolBgpPolicyImportRulesMatchExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Regex, b.Regex) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesMatch(a *ProtocolBgpPolicyImportRulesMatch, b *ProtocolBgpPolicyImportRulesMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.FromPeer, b.FromPeer) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesMatchAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesMatchCommunity(a.Community, b.Community) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesMatchExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.StringsMatch(a.RouteTable, b.RouteTable) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesMatchAddressPrefix(a.AddressPrefix, b.AddressPrefix) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionDeny(a *ProtocolBgpPolicyImportRulesActionDeny, b *ProtocolBgpPolicyImportRulesActionDeny) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone(a *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone, b *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll(a *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll, b *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity(a *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity, b *ProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone(a *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone, b *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove(a *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove, b *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPath(a *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath, b *ProtocolBgpPolicyImportRulesActionAllowUpdateAsPath) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPathNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPathRemove(a.Remove, b.Remove) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone(a *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone, b *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll(a *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll, b *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunity(a *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity, b *ProtocolBgpPolicyImportRulesActionAllowUpdateCommunity) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunityNone(a.None, b.None) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunityRemoveAll(a.RemoveAll, b.RemoveAll) {
		return false
	}
	if !util.StringsMatch(a.RemoveRegex, b.RemoveRegex) {
		return false
	}
	if !util.OrderedListsMatch(a.Append, b.Append) {
		return false
	}
	if !util.OrderedListsMatch(a.Overwrite, b.Overwrite) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllowUpdate(a *ProtocolBgpPolicyImportRulesActionAllowUpdate, b *ProtocolBgpPolicyImportRulesActionAllowUpdate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Weight, b.Weight) {
		return false
	}
	if !util.StringsMatch(a.Origin, b.Origin) {
		return false
	}
	if !util.Ints64Match(a.AsPathLimit, b.AsPathLimit) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateExtendedCommunity(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	if !util.Ints64Match(a.LocalPreference, b.LocalPreference) {
		return false
	}
	if !util.Ints64Match(a.Med, b.Med) {
		return false
	}
	if !util.StringsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateAsPath(a.AsPath, b.AsPath) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdateCommunity(a.Community, b.Community) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesActionAllow(a *ProtocolBgpPolicyImportRulesActionAllow, b *ProtocolBgpPolicyImportRulesActionAllow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Dampening, b.Dampening) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllowUpdate(a.Update, b.Update) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRulesAction(a *ProtocolBgpPolicyImportRulesAction, b *ProtocolBgpPolicyImportRulesAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyImportRulesActionDeny(a.Deny, b.Deny) {
		return false
	}
	if !matchProtocolBgpPolicyImportRulesActionAllow(a.Allow, b.Allow) {
		return false
	}
	return true
}
func matchProtocolBgpPolicyImportRules(a []ProtocolBgpPolicyImportRules, b []ProtocolBgpPolicyImportRules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.OrderedListsMatch(a.UsedBy, b.UsedBy) {
				return false
			}
			if !matchProtocolBgpPolicyImportRulesMatch(a.Match, b.Match) {
				return false
			}
			if !matchProtocolBgpPolicyImportRulesAction(a.Action, b.Action) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPolicyImport(a *ProtocolBgpPolicyImport, b *ProtocolBgpPolicyImport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyImportRules(a.Rules, b.Rules) {
		return false
	}
	return true
}
func matchProtocolBgpPolicy(a *ProtocolBgpPolicy, b *ProtocolBgpPolicy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPolicyAggregation(a.Aggregation, b.Aggregation) {
		return false
	}
	if !matchProtocolBgpPolicyConditionalAdvertisement(a.ConditionalAdvertisement, b.ConditionalAdvertisement) {
		return false
	}
	if !matchProtocolBgpPolicyExport(a.Export, b.Export) {
		return false
	}
	if !matchProtocolBgpPolicyImport(a.Import, b.Import) {
		return false
	}
	return true
}
func matchProtocolBgpDampeningProfile(a []ProtocolBgpDampeningProfile, b []ProtocolBgpDampeningProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.FloatsMatch(a.Cutoff, b.Cutoff) {
				return false
			}
			if !util.FloatsMatch(a.Reuse, b.Reuse) {
				return false
			}
			if !util.Ints64Match(a.MaxHoldTime, b.MaxHoldTime) {
				return false
			}
			if !util.Ints64Match(a.DecayHalfLifeReachable, b.DecayHalfLifeReachable) {
				return false
			}
			if !util.Ints64Match(a.DecayHalfLifeUnreachable, b.DecayHalfLifeUnreachable) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPeerGroupPeerPeerAddress(a *ProtocolBgpPeerGroupPeerPeerAddress, b *ProtocolBgpPeerGroupPeerPeerAddress) bool {
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
func matchProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier(a *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier, b *ProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Unicast, b.Unicast) {
		return false
	}
	if !util.BoolsMatch(a.Multicast, b.Multicast) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupPeerBfd(a *ProtocolBgpPeerGroupPeerBfd, b *ProtocolBgpPeerGroupPeerBfd) bool {
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
func matchProtocolBgpPeerGroupPeerLocalAddress(a *ProtocolBgpPeerGroupPeerLocalAddress, b *ProtocolBgpPeerGroupPeerLocalAddress) bool {
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
func matchProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection(a *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection, b *ProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.RemotePort, b.RemotePort) {
		return false
	}
	if !util.BoolsMatch(a.Allow, b.Allow) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection(a *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection, b *ProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.LocalPort, b.LocalPort) {
		return false
	}
	if !util.BoolsMatch(a.Allow, b.Allow) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupPeerConnectionOptions(a *ProtocolBgpPeerGroupPeerConnectionOptions, b *ProtocolBgpPeerGroupPeerConnectionOptions) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.KeepAliveInterval, b.KeepAliveInterval) {
		return false
	}
	if !util.Ints64Match(a.Multihop, b.Multihop) {
		return false
	}
	if !util.Ints64Match(a.OpenDelayTime, b.OpenDelayTime) {
		return false
	}
	if !util.StringsMatch(a.HoldTime, b.HoldTime) {
		return false
	}
	if !util.Ints64Match(a.MinRouteAdvInterval, b.MinRouteAdvInterval) {
		return false
	}
	if !util.Ints64Match(a.IdleHoldTime, b.IdleHoldTime) {
		return false
	}
	if !matchProtocolBgpPeerGroupPeerConnectionOptionsIncomingBgpConnection(a.IncomingBgpConnection, b.IncomingBgpConnection) {
		return false
	}
	if !matchProtocolBgpPeerGroupPeerConnectionOptionsOutgoingBgpConnection(a.OutgoingBgpConnection, b.OutgoingBgpConnection) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupPeer(a []ProtocolBgpPeerGroupPeer, b []ProtocolBgpPeerGroupPeer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.ReflectorClient, b.ReflectorClient) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeerPeerAddress(a.PeerAddress, b.PeerAddress) {
				return false
			}
			if !util.StringsMatch(a.AddressFamilyIdentifier, b.AddressFamilyIdentifier) {
				return false
			}
			if !util.StringsMatch(a.MaxPrefixes, b.MaxPrefixes) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.EnableMpBgp, b.EnableMpBgp) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeerSubsequentAddressFamilyIdentifier(a.SubsequentAddressFamilyIdentifier, b.SubsequentAddressFamilyIdentifier) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeerBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeerConnectionOptions(a.ConnectionOptions, b.ConnectionOptions) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.PeerAs, b.PeerAs) {
				return false
			}
			if !util.BoolsMatch(a.EnableSenderSideLoopDetection, b.EnableSenderSideLoopDetection) {
				return false
			}
			if !util.StringsMatch(a.PeeringType, b.PeeringType) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeerLocalAddress(a.LocalAddress, b.LocalAddress) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpPeerGroupTypeIbgpConfed(a *ProtocolBgpPeerGroupTypeIbgpConfed, b *ProtocolBgpPeerGroupTypeIbgpConfed) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ExportNexthop, b.ExportNexthop) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupTypeEbgp(a *ProtocolBgpPeerGroupTypeEbgp, b *ProtocolBgpPeerGroupTypeEbgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.RemovePrivateAs, b.RemovePrivateAs) {
		return false
	}
	if !util.StringsMatch(a.ImportNexthop, b.ImportNexthop) {
		return false
	}
	if !util.StringsMatch(a.ExportNexthop, b.ExportNexthop) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupTypeIbgp(a *ProtocolBgpPeerGroupTypeIbgp, b *ProtocolBgpPeerGroupTypeIbgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ExportNexthop, b.ExportNexthop) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupTypeEbgpConfed(a *ProtocolBgpPeerGroupTypeEbgpConfed, b *ProtocolBgpPeerGroupTypeEbgpConfed) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ExportNexthop, b.ExportNexthop) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroupType(a *ProtocolBgpPeerGroupType, b *ProtocolBgpPeerGroupType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpPeerGroupTypeEbgpConfed(a.EbgpConfed, b.EbgpConfed) {
		return false
	}
	if !matchProtocolBgpPeerGroupTypeIbgpConfed(a.IbgpConfed, b.IbgpConfed) {
		return false
	}
	if !matchProtocolBgpPeerGroupTypeEbgp(a.Ebgp, b.Ebgp) {
		return false
	}
	if !matchProtocolBgpPeerGroupTypeIbgp(a.Ibgp, b.Ibgp) {
		return false
	}
	return true
}
func matchProtocolBgpPeerGroup(a []ProtocolBgpPeerGroup, b []ProtocolBgpPeerGroup) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.SoftResetWithStoredInfo, b.SoftResetWithStoredInfo) {
				return false
			}
			if !matchProtocolBgpPeerGroupType(a.Type, b.Type) {
				return false
			}
			if !matchProtocolBgpPeerGroupPeer(a.Peer, b.Peer) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.AggregatedConfedAsPath, b.AggregatedConfedAsPath) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpRedistRules(a []ProtocolBgpRedistRules, b []ProtocolBgpRedistRules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.RouteTable, b.RouteTable) {
				return false
			}
			if !util.StringsMatch(a.SetOrigin, b.SetOrigin) {
				return false
			}
			if !util.Ints64Match(a.SetMed, b.SetMed) {
				return false
			}
			if !util.Ints64Match(a.SetAsPathLimit, b.SetAsPathLimit) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.OrderedListsMatch(a.SetCommunity, b.SetCommunity) {
				return false
			}
			if !util.StringsMatch(a.AddressFamilyIdentifier, b.AddressFamilyIdentifier) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.Ints64Match(a.SetLocalPreference, b.SetLocalPreference) {
				return false
			}
			if !util.OrderedListsMatch(a.SetExtendedCommunity, b.SetExtendedCommunity) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolBgpGlobalBfd(a *ProtocolBgpGlobalBfd, b *ProtocolBgpGlobalBfd) bool {
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
func matchProtocolBgpRoutingOptionsMed(a *ProtocolBgpRoutingOptionsMed, b *ProtocolBgpRoutingOptionsMed) bool {
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
func matchProtocolBgpRoutingOptionsAggregate(a *ProtocolBgpRoutingOptionsAggregate, b *ProtocolBgpRoutingOptionsAggregate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AggregateMed, b.AggregateMed) {
		return false
	}
	return true
}
func matchProtocolBgpRoutingOptionsGracefulRestart(a *ProtocolBgpRoutingOptionsGracefulRestart, b *ProtocolBgpRoutingOptionsGracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.LocalRestartTime, b.LocalRestartTime) {
		return false
	}
	if !util.Ints64Match(a.MaxPeerRestartTime, b.MaxPeerRestartTime) {
		return false
	}
	if !util.Ints64Match(a.StaleRouteTime, b.StaleRouteTime) {
		return false
	}
	return true
}
func matchProtocolBgpRoutingOptions(a *ProtocolBgpRoutingOptions, b *ProtocolBgpRoutingOptions) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpRoutingOptionsMed(a.Med, b.Med) {
		return false
	}
	if !util.StringsMatch(a.ReflectorClusterId, b.ReflectorClusterId) {
		return false
	}
	if !matchProtocolBgpRoutingOptionsAggregate(a.Aggregate, b.Aggregate) {
		return false
	}
	if !util.StringsMatch(a.AsFormat, b.AsFormat) {
		return false
	}
	if !util.StringsMatch(a.ConfederationMemberAs, b.ConfederationMemberAs) {
		return false
	}
	if !util.Ints64Match(a.DefaultLocalPreference, b.DefaultLocalPreference) {
		return false
	}
	if !matchProtocolBgpRoutingOptionsGracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	return true
}
func matchProtocolBgp(a *ProtocolBgp, b *ProtocolBgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgpDampeningProfile(a.DampeningProfile, b.DampeningProfile) {
		return false
	}
	if !util.BoolsMatch(a.EcmpMultiAs, b.EcmpMultiAs) {
		return false
	}
	if !matchProtocolBgpPeerGroup(a.PeerGroup, b.PeerGroup) {
		return false
	}
	if !util.BoolsMatch(a.RejectDefaultRoute, b.RejectDefaultRoute) {
		return false
	}
	if !util.BoolsMatch(a.AllowRedistDefaultRoute, b.AllowRedistDefaultRoute) {
		return false
	}
	if !util.BoolsMatch(a.EnforceFirstAs, b.EnforceFirstAs) {
		return false
	}
	if !matchProtocolBgpGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !util.StringsMatch(a.LocalAs, b.LocalAs) {
		return false
	}
	if !matchProtocolBgpRedistRules(a.RedistRules, b.RedistRules) {
		return false
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !matchProtocolBgpRoutingOptions(a.RoutingOptions, b.RoutingOptions) {
		return false
	}
	if !matchProtocolBgpAuthProfile(a.AuthProfile, b.AuthProfile) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.InstallRoute, b.InstallRoute) {
		return false
	}
	if !matchProtocolBgpPolicy(a.Policy, b.Policy) {
		return false
	}
	return true
}
func matchProtocolOspfGlobalBfd(a *ProtocolOspfGlobalBfd, b *ProtocolOspfGlobalBfd) bool {
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
func matchProtocolOspfAreaInterfaceNeighbor(a []ProtocolOspfAreaInterfaceNeighbor, b []ProtocolOspfAreaInterfaceNeighbor) bool {
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
		}
	}
	return true
}
func matchProtocolOspfAreaInterfaceLinkTypeBroadcast(a *ProtocolOspfAreaInterfaceLinkTypeBroadcast, b *ProtocolOspfAreaInterfaceLinkTypeBroadcast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaInterfaceLinkTypeP2p(a *ProtocolOspfAreaInterfaceLinkTypeP2p, b *ProtocolOspfAreaInterfaceLinkTypeP2p) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaInterfaceLinkTypeP2mp(a *ProtocolOspfAreaInterfaceLinkTypeP2mp, b *ProtocolOspfAreaInterfaceLinkTypeP2mp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaInterfaceLinkType(a *ProtocolOspfAreaInterfaceLinkType, b *ProtocolOspfAreaInterfaceLinkType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfAreaInterfaceLinkTypeBroadcast(a.Broadcast, b.Broadcast) {
		return false
	}
	if !matchProtocolOspfAreaInterfaceLinkTypeP2p(a.P2p, b.P2p) {
		return false
	}
	if !matchProtocolOspfAreaInterfaceLinkTypeP2mp(a.P2mp, b.P2mp) {
		return false
	}
	return true
}
func matchProtocolOspfAreaInterfaceBfd(a *ProtocolOspfAreaInterfaceBfd, b *ProtocolOspfAreaInterfaceBfd) bool {
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
func matchProtocolOspfAreaInterface(a []ProtocolOspfAreaInterface, b []ProtocolOspfAreaInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchProtocolOspfAreaInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
			if !util.Ints64Match(a.RetransmitInterval, b.RetransmitInterval) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.Ints64Match(a.GrDelay, b.GrDelay) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !matchProtocolOspfAreaInterfaceNeighbor(a.Neighbor, b.Neighbor) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.Passive, b.Passive) {
				return false
			}
			if !util.Ints64Match(a.HelloInterval, b.HelloInterval) {
				return false
			}
			if !matchProtocolOspfAreaInterfaceLinkType(a.LinkType, b.LinkType) {
				return false
			}
			if !util.Ints64Match(a.DeadCounts, b.DeadCounts) {
				return false
			}
			if !util.Ints64Match(a.TransitDelay, b.TransitDelay) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfAreaVirtualLinkBfd(a *ProtocolOspfAreaVirtualLinkBfd, b *ProtocolOspfAreaVirtualLinkBfd) bool {
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
func matchProtocolOspfAreaVirtualLink(a []ProtocolOspfAreaVirtualLink, b []ProtocolOspfAreaVirtualLink) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.Ints64Match(a.HelloInterval, b.HelloInterval) {
				return false
			}
			if !matchProtocolOspfAreaVirtualLinkBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.Ints64Match(a.TransitDelay, b.TransitDelay) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.NeighborId, b.NeighborId) {
				return false
			}
			if !util.StringsMatch(a.TransitAreaId, b.TransitAreaId) {
				return false
			}
			if !util.Ints64Match(a.DeadCounts, b.DeadCounts) {
				return false
			}
			if !util.Ints64Match(a.RetransmitInterval, b.RetransmitInterval) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfAreaTypeNormal(a *ProtocolOspfAreaTypeNormal, b *ProtocolOspfAreaTypeNormal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaTypeStubDefaultRouteDisable(a *ProtocolOspfAreaTypeStubDefaultRouteDisable, b *ProtocolOspfAreaTypeStubDefaultRouteDisable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaTypeStubDefaultRouteAdvertise(a *ProtocolOspfAreaTypeStubDefaultRouteAdvertise, b *ProtocolOspfAreaTypeStubDefaultRouteAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	return true
}
func matchProtocolOspfAreaTypeStubDefaultRoute(a *ProtocolOspfAreaTypeStubDefaultRoute, b *ProtocolOspfAreaTypeStubDefaultRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfAreaTypeStubDefaultRouteDisable(a.Disable, b.Disable) {
		return false
	}
	if !matchProtocolOspfAreaTypeStubDefaultRouteAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchProtocolOspfAreaTypeStub(a *ProtocolOspfAreaTypeStub, b *ProtocolOspfAreaTypeStub) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AcceptSummary, b.AcceptSummary) {
		return false
	}
	if !matchProtocolOspfAreaTypeStubDefaultRoute(a.DefaultRoute, b.DefaultRoute) {
		return false
	}
	return true
}
func matchProtocolOspfAreaTypeNssaDefaultRouteDisable(a *ProtocolOspfAreaTypeNssaDefaultRouteDisable, b *ProtocolOspfAreaTypeNssaDefaultRouteDisable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaTypeNssaDefaultRouteAdvertise(a *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise, b *ProtocolOspfAreaTypeNssaDefaultRouteAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	return true
}
func matchProtocolOspfAreaTypeNssaDefaultRoute(a *ProtocolOspfAreaTypeNssaDefaultRoute, b *ProtocolOspfAreaTypeNssaDefaultRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfAreaTypeNssaDefaultRouteDisable(a.Disable, b.Disable) {
		return false
	}
	if !matchProtocolOspfAreaTypeNssaDefaultRouteAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchProtocolOspfAreaTypeNssaNssaExtRangeAdvertise(a *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise, b *ProtocolOspfAreaTypeNssaNssaExtRangeAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaTypeNssaNssaExtRangeSuppress(a *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress, b *ProtocolOspfAreaTypeNssaNssaExtRangeSuppress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaTypeNssaNssaExtRange(a []ProtocolOspfAreaTypeNssaNssaExtRange, b []ProtocolOspfAreaTypeNssaNssaExtRange) bool {
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
			if !matchProtocolOspfAreaTypeNssaNssaExtRangeAdvertise(a.Advertise, b.Advertise) {
				return false
			}
			if !matchProtocolOspfAreaTypeNssaNssaExtRangeSuppress(a.Suppress, b.Suppress) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfAreaTypeNssa(a *ProtocolOspfAreaTypeNssa, b *ProtocolOspfAreaTypeNssa) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AcceptSummary, b.AcceptSummary) {
		return false
	}
	if !matchProtocolOspfAreaTypeNssaDefaultRoute(a.DefaultRoute, b.DefaultRoute) {
		return false
	}
	if !matchProtocolOspfAreaTypeNssaNssaExtRange(a.NssaExtRange, b.NssaExtRange) {
		return false
	}
	return true
}
func matchProtocolOspfAreaType(a *ProtocolOspfAreaType, b *ProtocolOspfAreaType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfAreaTypeNormal(a.Normal, b.Normal) {
		return false
	}
	if !matchProtocolOspfAreaTypeStub(a.Stub, b.Stub) {
		return false
	}
	if !matchProtocolOspfAreaTypeNssa(a.Nssa, b.Nssa) {
		return false
	}
	return true
}
func matchProtocolOspfAreaRangeAdvertise(a *ProtocolOspfAreaRangeAdvertise, b *ProtocolOspfAreaRangeAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaRangeSuppress(a *ProtocolOspfAreaRangeSuppress, b *ProtocolOspfAreaRangeSuppress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfAreaRange(a []ProtocolOspfAreaRange, b []ProtocolOspfAreaRange) bool {
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
			if !matchProtocolOspfAreaRangeSuppress(a.Suppress, b.Suppress) {
				return false
			}
			if !matchProtocolOspfAreaRangeAdvertise(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfArea(a []ProtocolOspfArea, b []ProtocolOspfArea) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchProtocolOspfAreaType(a.Type, b.Type) {
				return false
			}
			if !matchProtocolOspfAreaRange(a.Range, b.Range) {
				return false
			}
			if !matchProtocolOspfAreaInterface(a.Interface, b.Interface) {
				return false
			}
			if !matchProtocolOspfAreaVirtualLink(a.VirtualLink, b.VirtualLink) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfAuthProfileMd5(a []ProtocolOspfAuthProfileMd5, b []ProtocolOspfAuthProfileMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Key, b.Key) {
				return false
			}
			if !util.BoolsMatch(a.Preferred, b.Preferred) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfAuthProfile(a []ProtocolOspfAuthProfile, b []ProtocolOspfAuthProfile) bool {
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
			if !util.StringsMatch(a.Password, b.Password) {
				return false
			}
			if !matchProtocolOspfAuthProfileMd5(a.Md5, b.Md5) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfExportRules(a []ProtocolOspfExportRules, b []ProtocolOspfExportRules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.NewPathType, b.NewPathType) {
				return false
			}
			if !util.StringsMatch(a.NewTag, b.NewTag) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfTimers(a *ProtocolOspfTimers, b *ProtocolOspfTimers) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.FloatsMatch(a.LsaInterval, b.LsaInterval) {
		return false
	}
	if !util.FloatsMatch(a.SpfCalculationDelay, b.SpfCalculationDelay) {
		return false
	}
	return true
}
func matchProtocolOspfGracefulRestart(a *ProtocolOspfGracefulRestart, b *ProtocolOspfGracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.MaxNeighborRestartTime, b.MaxNeighborRestartTime) {
		return false
	}
	if !util.BoolsMatch(a.StrictLSAChecking, b.StrictLSAChecking) {
		return false
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
	return true
}
func matchProtocolOspf(a *ProtocolOspf, b *ProtocolOspf) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfExportRules(a.ExportRules, b.ExportRules) {
		return false
	}
	if !matchProtocolOspfGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !matchProtocolOspfArea(a.Area, b.Area) {
		return false
	}
	if !matchProtocolOspfAuthProfile(a.AuthProfile, b.AuthProfile) {
		return false
	}
	if !matchProtocolOspfGracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	if !util.BoolsMatch(a.RejectDefaultRoute, b.RejectDefaultRoute) {
		return false
	}
	if !util.BoolsMatch(a.Rfc1583, b.Rfc1583) {
		return false
	}
	if !matchProtocolOspfTimers(a.Timers, b.Timers) {
		return false
	}
	if !util.BoolsMatch(a.AllowRedistDefaultRoute, b.AllowRedistDefaultRoute) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchProtocolOspfv3ExportRules(a []ProtocolOspfv3ExportRules, b []ProtocolOspfv3ExportRules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.NewPathType, b.NewPathType) {
				return false
			}
			if !util.StringsMatch(a.NewTag, b.NewTag) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3GlobalBfd(a *ProtocolOspfv3GlobalBfd, b *ProtocolOspfv3GlobalBfd) bool {
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
func matchProtocolOspfv3AreaRangeAdvertise(a *ProtocolOspfv3AreaRangeAdvertise, b *ProtocolOspfv3AreaRangeAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaRangeSuppress(a *ProtocolOspfv3AreaRangeSuppress, b *ProtocolOspfv3AreaRangeSuppress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaRange(a []ProtocolOspfv3AreaRange, b []ProtocolOspfv3AreaRange) bool {
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
			if !matchProtocolOspfv3AreaRangeAdvertise(a.Advertise, b.Advertise) {
				return false
			}
			if !matchProtocolOspfv3AreaRangeSuppress(a.Suppress, b.Suppress) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3AreaInterfaceNeighbor(a []ProtocolOspfv3AreaInterfaceNeighbor, b []ProtocolOspfv3AreaInterfaceNeighbor) bool {
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
		}
	}
	return true
}
func matchProtocolOspfv3AreaInterfaceBfd(a *ProtocolOspfv3AreaInterfaceBfd, b *ProtocolOspfv3AreaInterfaceBfd) bool {
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
func matchProtocolOspfv3AreaInterfaceLinkTypeBroadcast(a *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast, b *ProtocolOspfv3AreaInterfaceLinkTypeBroadcast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaInterfaceLinkTypeP2p(a *ProtocolOspfv3AreaInterfaceLinkTypeP2p, b *ProtocolOspfv3AreaInterfaceLinkTypeP2p) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaInterfaceLinkTypeP2mp(a *ProtocolOspfv3AreaInterfaceLinkTypeP2mp, b *ProtocolOspfv3AreaInterfaceLinkTypeP2mp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaInterfaceLinkType(a *ProtocolOspfv3AreaInterfaceLinkType, b *ProtocolOspfv3AreaInterfaceLinkType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AreaInterfaceLinkTypeBroadcast(a.Broadcast, b.Broadcast) {
		return false
	}
	if !matchProtocolOspfv3AreaInterfaceLinkTypeP2p(a.P2p, b.P2p) {
		return false
	}
	if !matchProtocolOspfv3AreaInterfaceLinkTypeP2mp(a.P2mp, b.P2mp) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaInterface(a []ProtocolOspfv3AreaInterface, b []ProtocolOspfv3AreaInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.DeadCounts, b.DeadCounts) {
				return false
			}
			if !util.Ints64Match(a.RetransmitInterval, b.RetransmitInterval) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchProtocolOspfv3AreaInterfaceNeighbor(a.Neighbor, b.Neighbor) {
				return false
			}
			if !matchProtocolOspfv3AreaInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.BoolsMatch(a.Passive, b.Passive) {
				return false
			}
			if !util.Ints64Match(a.GrDelay, b.GrDelay) {
				return false
			}
			if !matchProtocolOspfv3AreaInterfaceLinkType(a.LinkType, b.LinkType) {
				return false
			}
			if !util.Ints64Match(a.InstanceId, b.InstanceId) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.Ints64Match(a.HelloInterval, b.HelloInterval) {
				return false
			}
			if !util.Ints64Match(a.TransitDelay, b.TransitDelay) {
				return false
			}
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
func matchProtocolOspfv3AreaVirtualLinkBfd(a *ProtocolOspfv3AreaVirtualLinkBfd, b *ProtocolOspfv3AreaVirtualLinkBfd) bool {
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
func matchProtocolOspfv3AreaVirtualLink(a []ProtocolOspfv3AreaVirtualLink, b []ProtocolOspfv3AreaVirtualLink) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchProtocolOspfv3AreaVirtualLinkBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.StringsMatch(a.NeighborId, b.NeighborId) {
				return false
			}
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.Ints64Match(a.HelloInterval, b.HelloInterval) {
				return false
			}
			if !util.Ints64Match(a.TransitDelay, b.TransitDelay) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.TransitAreaId, b.TransitAreaId) {
				return false
			}
			if !util.Ints64Match(a.InstanceId, b.InstanceId) {
				return false
			}
			if !util.Ints64Match(a.DeadCounts, b.DeadCounts) {
				return false
			}
			if !util.Ints64Match(a.RetransmitInterval, b.RetransmitInterval) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3AreaTypeNormal(a *ProtocolOspfv3AreaTypeNormal, b *ProtocolOspfv3AreaTypeNormal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaTypeStubDefaultRouteDisable(a *ProtocolOspfv3AreaTypeStubDefaultRouteDisable, b *ProtocolOspfv3AreaTypeStubDefaultRouteDisable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaTypeStubDefaultRouteAdvertise(a *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise, b *ProtocolOspfv3AreaTypeStubDefaultRouteAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaTypeStubDefaultRoute(a *ProtocolOspfv3AreaTypeStubDefaultRoute, b *ProtocolOspfv3AreaTypeStubDefaultRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AreaTypeStubDefaultRouteDisable(a.Disable, b.Disable) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeStubDefaultRouteAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaTypeStub(a *ProtocolOspfv3AreaTypeStub, b *ProtocolOspfv3AreaTypeStub) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AcceptSummary, b.AcceptSummary) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeStubDefaultRoute(a.DefaultRoute, b.DefaultRoute) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaDefaultRouteDisable(a *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable, b *ProtocolOspfv3AreaTypeNssaDefaultRouteDisable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise(a *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise, b *ProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaDefaultRoute(a *ProtocolOspfv3AreaTypeNssaDefaultRoute, b *ProtocolOspfv3AreaTypeNssaDefaultRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AreaTypeNssaDefaultRouteDisable(a.Disable, b.Disable) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeNssaDefaultRouteAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise(a *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise, b *ProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress(a *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress, b *ProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssaNssaExtRange(a []ProtocolOspfv3AreaTypeNssaNssaExtRange, b []ProtocolOspfv3AreaTypeNssaNssaExtRange) bool {
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
			if !matchProtocolOspfv3AreaTypeNssaNssaExtRangeSuppress(a.Suppress, b.Suppress) {
				return false
			}
			if !matchProtocolOspfv3AreaTypeNssaNssaExtRangeAdvertise(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3AreaTypeNssa(a *ProtocolOspfv3AreaTypeNssa, b *ProtocolOspfv3AreaTypeNssa) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AreaTypeNssaNssaExtRange(a.NssaExtRange, b.NssaExtRange) {
		return false
	}
	if !util.BoolsMatch(a.AcceptSummary, b.AcceptSummary) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeNssaDefaultRoute(a.DefaultRoute, b.DefaultRoute) {
		return false
	}
	return true
}
func matchProtocolOspfv3AreaType(a *ProtocolOspfv3AreaType, b *ProtocolOspfv3AreaType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AreaTypeNormal(a.Normal, b.Normal) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeStub(a.Stub, b.Stub) {
		return false
	}
	if !matchProtocolOspfv3AreaTypeNssa(a.Nssa, b.Nssa) {
		return false
	}
	return true
}
func matchProtocolOspfv3Area(a []ProtocolOspfv3Area, b []ProtocolOspfv3Area) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !matchProtocolOspfv3AreaType(a.Type, b.Type) {
				return false
			}
			if !matchProtocolOspfv3AreaRange(a.Range, b.Range) {
				return false
			}
			if !matchProtocolOspfv3AreaInterface(a.Interface, b.Interface) {
				return false
			}
			if !matchProtocolOspfv3AreaVirtualLink(a.VirtualLink, b.VirtualLink) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationSha1(a *ProtocolOspfv3AuthProfileEspAuthenticationSha1, b *ProtocolOspfv3AuthProfileEspAuthenticationSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationSha256(a *ProtocolOspfv3AuthProfileEspAuthenticationSha256, b *ProtocolOspfv3AuthProfileEspAuthenticationSha256) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationSha384(a *ProtocolOspfv3AuthProfileEspAuthenticationSha384, b *ProtocolOspfv3AuthProfileEspAuthenticationSha384) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationSha512(a *ProtocolOspfv3AuthProfileEspAuthenticationSha512, b *ProtocolOspfv3AuthProfileEspAuthenticationSha512) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationNone(a *ProtocolOspfv3AuthProfileEspAuthenticationNone, b *ProtocolOspfv3AuthProfileEspAuthenticationNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthenticationMd5(a *ProtocolOspfv3AuthProfileEspAuthenticationMd5, b *ProtocolOspfv3AuthProfileEspAuthenticationMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspAuthentication(a *ProtocolOspfv3AuthProfileEspAuthentication, b *ProtocolOspfv3AuthProfileEspAuthentication) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationSha384(a.Sha384, b.Sha384) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationSha512(a.Sha512, b.Sha512) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationNone(a.None, b.None) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationSha1(a.Sha1, b.Sha1) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspAuthenticationSha256(a.Sha256, b.Sha256) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEspEncryption(a *ProtocolOspfv3AuthProfileEspEncryption, b *ProtocolOspfv3AuthProfileEspEncryption) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileEsp(a *ProtocolOspfv3AuthProfileEsp, b *ProtocolOspfv3AuthProfileEsp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AuthProfileEspAuthentication(a.Authentication, b.Authentication) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileEspEncryption(a.Encryption, b.Encryption) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAhSha1(a *ProtocolOspfv3AuthProfileAhSha1, b *ProtocolOspfv3AuthProfileAhSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAhSha256(a *ProtocolOspfv3AuthProfileAhSha256, b *ProtocolOspfv3AuthProfileAhSha256) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAhSha384(a *ProtocolOspfv3AuthProfileAhSha384, b *ProtocolOspfv3AuthProfileAhSha384) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAhSha512(a *ProtocolOspfv3AuthProfileAhSha512, b *ProtocolOspfv3AuthProfileAhSha512) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAhMd5(a *ProtocolOspfv3AuthProfileAhMd5, b *ProtocolOspfv3AuthProfileAhMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfileAh(a *ProtocolOspfv3AuthProfileAh, b *ProtocolOspfv3AuthProfileAh) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolOspfv3AuthProfileAhSha512(a.Sha512, b.Sha512) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileAhMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileAhSha1(a.Sha1, b.Sha1) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileAhSha256(a.Sha256, b.Sha256) {
		return false
	}
	if !matchProtocolOspfv3AuthProfileAhSha384(a.Sha384, b.Sha384) {
		return false
	}
	return true
}
func matchProtocolOspfv3AuthProfile(a []ProtocolOspfv3AuthProfile, b []ProtocolOspfv3AuthProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Spi, b.Spi) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !matchProtocolOspfv3AuthProfileEsp(a.Esp, b.Esp) {
				return false
			}
			if !matchProtocolOspfv3AuthProfileAh(a.Ah, b.Ah) {
				return false
			}
		}
	}
	return true
}
func matchProtocolOspfv3GracefulRestart(a *ProtocolOspfv3GracefulRestart, b *ProtocolOspfv3GracefulRestart) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.StrictLSAChecking, b.StrictLSAChecking) {
		return false
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
	if !util.Ints64Match(a.MaxNeighborRestartTime, b.MaxNeighborRestartTime) {
		return false
	}
	return true
}
func matchProtocolOspfv3Timers(a *ProtocolOspfv3Timers, b *ProtocolOspfv3Timers) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.FloatsMatch(a.LsaInterval, b.LsaInterval) {
		return false
	}
	if !util.FloatsMatch(a.SpfCalculationDelay, b.SpfCalculationDelay) {
		return false
	}
	return true
}
func matchProtocolOspfv3(a *ProtocolOspfv3, b *ProtocolOspfv3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.AllowRedistDefaultRoute, b.AllowRedistDefaultRoute) {
		return false
	}
	if !matchProtocolOspfv3AuthProfile(a.AuthProfile, b.AuthProfile) {
		return false
	}
	if !matchProtocolOspfv3GracefulRestart(a.GracefulRestart, b.GracefulRestart) {
		return false
	}
	if !util.StringsMatch(a.RouterId, b.RouterId) {
		return false
	}
	if !matchProtocolOspfv3Timers(a.Timers, b.Timers) {
		return false
	}
	if !util.BoolsMatch(a.RejectDefaultRoute, b.RejectDefaultRoute) {
		return false
	}
	if !matchProtocolOspfv3Area(a.Area, b.Area) {
		return false
	}
	if !util.BoolsMatch(a.DisableTransitTraffic, b.DisableTransitTraffic) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchProtocolOspfv3ExportRules(a.ExportRules, b.ExportRules) {
		return false
	}
	if !matchProtocolOspfv3GlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	return true
}
func matchProtocolRedistProfileFilterBgp(a *ProtocolRedistProfileFilterBgp, b *ProtocolRedistProfileFilterBgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Community, b.Community) {
		return false
	}
	if !util.OrderedListsMatch(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	return true
}
func matchProtocolRedistProfileFilterOspf(a *ProtocolRedistProfileFilterOspf, b *ProtocolRedistProfileFilterOspf) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.PathType, b.PathType) {
		return false
	}
	if !util.OrderedListsMatch(a.Area, b.Area) {
		return false
	}
	if !util.OrderedListsMatch(a.Tag, b.Tag) {
		return false
	}
	return true
}
func matchProtocolRedistProfileFilter(a *ProtocolRedistProfileFilter, b *ProtocolRedistProfileFilter) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Type, b.Type) {
		return false
	}
	if !util.OrderedListsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.OrderedListsMatch(a.Destination, b.Destination) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !matchProtocolRedistProfileFilterOspf(a.Ospf, b.Ospf) {
		return false
	}
	if !matchProtocolRedistProfileFilterBgp(a.Bgp, b.Bgp) {
		return false
	}
	return true
}
func matchProtocolRedistProfileActionNoRedist(a *ProtocolRedistProfileActionNoRedist, b *ProtocolRedistProfileActionNoRedist) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolRedistProfileActionRedist(a *ProtocolRedistProfileActionRedist, b *ProtocolRedistProfileActionRedist) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolRedistProfileAction(a *ProtocolRedistProfileAction, b *ProtocolRedistProfileAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolRedistProfileActionNoRedist(a.NoRedist, b.NoRedist) {
		return false
	}
	if !matchProtocolRedistProfileActionRedist(a.Redist, b.Redist) {
		return false
	}
	return true
}
func matchProtocolRedistProfile(a []ProtocolRedistProfile, b []ProtocolRedistProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
			if !matchProtocolRedistProfileFilter(a.Filter, b.Filter) {
				return false
			}
			if !matchProtocolRedistProfileAction(a.Action, b.Action) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRedistProfileIpv6FilterOspfv3(a *ProtocolRedistProfileIpv6FilterOspfv3, b *ProtocolRedistProfileIpv6FilterOspfv3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.PathType, b.PathType) {
		return false
	}
	if !util.OrderedListsMatch(a.Area, b.Area) {
		return false
	}
	if !util.OrderedListsMatch(a.Tag, b.Tag) {
		return false
	}
	return true
}
func matchProtocolRedistProfileIpv6FilterBgp(a *ProtocolRedistProfileIpv6FilterBgp, b *ProtocolRedistProfileIpv6FilterBgp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Community, b.Community) {
		return false
	}
	if !util.OrderedListsMatch(a.ExtendedCommunity, b.ExtendedCommunity) {
		return false
	}
	return true
}
func matchProtocolRedistProfileIpv6Filter(a *ProtocolRedistProfileIpv6Filter, b *ProtocolRedistProfileIpv6Filter) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Type, b.Type) {
		return false
	}
	if !util.OrderedListsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.OrderedListsMatch(a.Destination, b.Destination) {
		return false
	}
	if !util.OrderedListsMatch(a.Nexthop, b.Nexthop) {
		return false
	}
	if !matchProtocolRedistProfileIpv6FilterOspfv3(a.Ospfv3, b.Ospfv3) {
		return false
	}
	if !matchProtocolRedistProfileIpv6FilterBgp(a.Bgp, b.Bgp) {
		return false
	}
	return true
}
func matchProtocolRedistProfileIpv6ActionNoRedist(a *ProtocolRedistProfileIpv6ActionNoRedist, b *ProtocolRedistProfileIpv6ActionNoRedist) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolRedistProfileIpv6ActionRedist(a *ProtocolRedistProfileIpv6ActionRedist, b *ProtocolRedistProfileIpv6ActionRedist) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolRedistProfileIpv6Action(a *ProtocolRedistProfileIpv6Action, b *ProtocolRedistProfileIpv6Action) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolRedistProfileIpv6ActionNoRedist(a.NoRedist, b.NoRedist) {
		return false
	}
	if !matchProtocolRedistProfileIpv6ActionRedist(a.Redist, b.Redist) {
		return false
	}
	return true
}
func matchProtocolRedistProfileIpv6(a []ProtocolRedistProfileIpv6, b []ProtocolRedistProfileIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Priority, b.Priority) {
				return false
			}
			if !matchProtocolRedistProfileIpv6Filter(a.Filter, b.Filter) {
				return false
			}
			if !matchProtocolRedistProfileIpv6Action(a.Action, b.Action) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRipExportRules(a []ProtocolRipExportRules, b []ProtocolRipExportRules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRipGlobalBfd(a *ProtocolRipGlobalBfd, b *ProtocolRipGlobalBfd) bool {
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
func matchProtocolRipInterfaceDefaultRouteAdvertise(a *ProtocolRipInterfaceDefaultRouteAdvertise, b *ProtocolRipInterfaceDefaultRouteAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Metric, b.Metric) {
		return false
	}
	return true
}
func matchProtocolRipInterfaceDefaultRouteDisable(a *ProtocolRipInterfaceDefaultRouteDisable, b *ProtocolRipInterfaceDefaultRouteDisable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchProtocolRipInterfaceDefaultRoute(a *ProtocolRipInterfaceDefaultRoute, b *ProtocolRipInterfaceDefaultRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolRipInterfaceDefaultRouteAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	if !matchProtocolRipInterfaceDefaultRouteDisable(a.Disable, b.Disable) {
		return false
	}
	return true
}
func matchProtocolRipInterfaceBfd(a *ProtocolRipInterfaceBfd, b *ProtocolRipInterfaceBfd) bool {
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
func matchProtocolRipInterface(a []ProtocolRipInterface, b []ProtocolRipInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Enable, b.Enable) {
				return false
			}
			if !util.StringsMatch(a.Authentication, b.Authentication) {
				return false
			}
			if !util.StringsMatch(a.Mode, b.Mode) {
				return false
			}
			if !matchProtocolRipInterfaceDefaultRoute(a.DefaultRoute, b.DefaultRoute) {
				return false
			}
			if !matchProtocolRipInterfaceBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRipTimers(a *ProtocolRipTimers, b *ProtocolRipTimers) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.ExpireIntervals, b.ExpireIntervals) {
		return false
	}
	if !util.Ints64Match(a.IntervalSeconds, b.IntervalSeconds) {
		return false
	}
	if !util.Ints64Match(a.UpdateIntervals, b.UpdateIntervals) {
		return false
	}
	if !util.Ints64Match(a.DeleteIntervals, b.DeleteIntervals) {
		return false
	}
	return true
}
func matchProtocolRipAuthProfileMd5(a []ProtocolRipAuthProfileMd5, b []ProtocolRipAuthProfileMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Key, b.Key) {
				return false
			}
			if !util.BoolsMatch(a.Preferred, b.Preferred) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRipAuthProfile(a []ProtocolRipAuthProfile, b []ProtocolRipAuthProfile) bool {
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
			if !util.StringsMatch(a.Password, b.Password) {
				return false
			}
			if !matchProtocolRipAuthProfileMd5(a.Md5, b.Md5) {
				return false
			}
		}
	}
	return true
}
func matchProtocolRip(a *ProtocolRip, b *ProtocolRip) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchProtocolRipExportRules(a.ExportRules, b.ExportRules) {
		return false
	}
	if !matchProtocolRipGlobalBfd(a.GlobalBfd, b.GlobalBfd) {
		return false
	}
	if !matchProtocolRipInterface(a.Interface, b.Interface) {
		return false
	}
	if !util.BoolsMatch(a.RejectDefaultRoute, b.RejectDefaultRoute) {
		return false
	}
	if !matchProtocolRipTimers(a.Timers, b.Timers) {
		return false
	}
	if !util.BoolsMatch(a.AllowRedistDefaultRoute, b.AllowRedistDefaultRoute) {
		return false
	}
	if !matchProtocolRipAuthProfile(a.AuthProfile, b.AuthProfile) {
		return false
	}
	return true
}
func matchProtocol(a *Protocol, b *Protocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgp(a.Bgp, b.Bgp) {
		return false
	}
	if !matchProtocolOspf(a.Ospf, b.Ospf) {
		return false
	}
	if !matchProtocolOspfv3(a.Ospfv3, b.Ospfv3) {
		return false
	}
	if !matchProtocolRedistProfile(a.RedistProfile, b.RedistProfile) {
		return false
	}
	if !matchProtocolRedistProfileIpv6(a.RedistProfileIpv6, b.RedistProfileIpv6) {
		return false
	}
	if !matchProtocolRip(a.Rip, b.Rip) {
		return false
	}
	return true
}
func matchRoutingTableIpStaticRouteRouteTableUnicast(a *RoutingTableIpStaticRouteRouteTableUnicast, b *RoutingTableIpStaticRouteRouteTableUnicast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteRouteTableMulticast(a *RoutingTableIpStaticRouteRouteTableMulticast, b *RoutingTableIpStaticRouteRouteTableMulticast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteRouteTableBoth(a *RoutingTableIpStaticRouteRouteTableBoth, b *RoutingTableIpStaticRouteRouteTableBoth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteRouteTableNoInstall(a *RoutingTableIpStaticRouteRouteTableNoInstall, b *RoutingTableIpStaticRouteRouteTableNoInstall) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteRouteTable(a *RoutingTableIpStaticRouteRouteTable, b *RoutingTableIpStaticRouteRouteTable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpStaticRouteRouteTableUnicast(a.Unicast, b.Unicast) {
		return false
	}
	if !matchRoutingTableIpStaticRouteRouteTableMulticast(a.Multicast, b.Multicast) {
		return false
	}
	if !matchRoutingTableIpStaticRouteRouteTableBoth(a.Both, b.Both) {
		return false
	}
	if !matchRoutingTableIpStaticRouteRouteTableNoInstall(a.NoInstall, b.NoInstall) {
		return false
	}
	return true
}
func matchRoutingTableIpStaticRouteNexthopReceive(a *RoutingTableIpStaticRouteNexthopReceive, b *RoutingTableIpStaticRouteNexthopReceive) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteNexthopDiscard(a *RoutingTableIpStaticRouteNexthopDiscard, b *RoutingTableIpStaticRouteNexthopDiscard) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpStaticRouteNexthop(a *RoutingTableIpStaticRouteNexthop, b *RoutingTableIpStaticRouteNexthop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.IpAddress, b.IpAddress) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.NextVr, b.NextVr) {
		return false
	}
	if !matchRoutingTableIpStaticRouteNexthopReceive(a.Receive, b.Receive) {
		return false
	}
	if !matchRoutingTableIpStaticRouteNexthopDiscard(a.Discard, b.Discard) {
		return false
	}
	return true
}
func matchRoutingTableIpStaticRouteBfd(a *RoutingTableIpStaticRouteBfd, b *RoutingTableIpStaticRouteBfd) bool {
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
func matchRoutingTableIpStaticRoutePathMonitorMonitorDestinations(a []RoutingTableIpStaticRoutePathMonitorMonitorDestinations, b []RoutingTableIpStaticRoutePathMonitorMonitorDestinations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
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
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchRoutingTableIpStaticRoutePathMonitor(a *RoutingTableIpStaticRoutePathMonitor, b *RoutingTableIpStaticRoutePathMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.HoldTime, b.HoldTime) {
		return false
	}
	if !matchRoutingTableIpStaticRoutePathMonitorMonitorDestinations(a.MonitorDestinations, b.MonitorDestinations) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.FailureCondition, b.FailureCondition) {
		return false
	}
	return true
}
func matchRoutingTableIpStaticRoute(a []RoutingTableIpStaticRoute, b []RoutingTableIpStaticRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !matchRoutingTableIpStaticRouteRouteTable(a.RouteTable, b.RouteTable) {
				return false
			}
			if !matchRoutingTableIpStaticRouteBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !matchRoutingTableIpStaticRoutePathMonitor(a.PathMonitor, b.PathMonitor) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !matchRoutingTableIpStaticRouteNexthop(a.Nexthop, b.Nexthop) {
				return false
			}
		}
	}
	return true
}
func matchRoutingTableIp(a *RoutingTableIp, b *RoutingTableIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpStaticRoute(a.StaticRoute, b.StaticRoute) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRouteBfd(a *RoutingTableIpv6StaticRouteBfd, b *RoutingTableIpv6StaticRouteBfd) bool {
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
func matchRoutingTableIpv6StaticRouteNexthopReceive(a *RoutingTableIpv6StaticRouteNexthopReceive, b *RoutingTableIpv6StaticRouteNexthopReceive) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpv6StaticRouteNexthopDiscard(a *RoutingTableIpv6StaticRouteNexthopDiscard, b *RoutingTableIpv6StaticRouteNexthopDiscard) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpv6StaticRouteNexthop(a *RoutingTableIpv6StaticRouteNexthop, b *RoutingTableIpv6StaticRouteNexthop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpv6StaticRouteNexthopReceive(a.Receive, b.Receive) {
		return false
	}
	if !matchRoutingTableIpv6StaticRouteNexthopDiscard(a.Discard, b.Discard) {
		return false
	}
	if !util.StringsMatch(a.Ipv6Address, b.Ipv6Address) {
		return false
	}
	if !util.StringsMatch(a.NextVr, b.NextVr) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRouteOption(a *RoutingTableIpv6StaticRouteOption, b *RoutingTableIpv6StaticRouteOption) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpv6StaticRouteRouteTableUnicast(a *RoutingTableIpv6StaticRouteRouteTableUnicast, b *RoutingTableIpv6StaticRouteRouteTableUnicast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpv6StaticRouteRouteTableNoInstall(a *RoutingTableIpv6StaticRouteRouteTableNoInstall, b *RoutingTableIpv6StaticRouteRouteTableNoInstall) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRoutingTableIpv6StaticRouteRouteTable(a *RoutingTableIpv6StaticRouteRouteTable, b *RoutingTableIpv6StaticRouteRouteTable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpv6StaticRouteRouteTableUnicast(a.Unicast, b.Unicast) {
		return false
	}
	if !matchRoutingTableIpv6StaticRouteRouteTableNoInstall(a.NoInstall, b.NoInstall) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations(a []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations, b []RoutingTableIpv6StaticRoutePathMonitorMonitorDestinations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Interval, b.Interval) {
				return false
			}
			if !util.Ints64Match(a.Count, b.Count) {
				return false
			}
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
		}
	}
	return true
}
func matchRoutingTableIpv6StaticRoutePathMonitor(a *RoutingTableIpv6StaticRoutePathMonitor, b *RoutingTableIpv6StaticRoutePathMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.FailureCondition, b.FailureCondition) {
		return false
	}
	if !util.Ints64Match(a.HoldTime, b.HoldTime) {
		return false
	}
	if !matchRoutingTableIpv6StaticRoutePathMonitorMonitorDestinations(a.MonitorDestinations, b.MonitorDestinations) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRoute(a []RoutingTableIpv6StaticRoute, b []RoutingTableIpv6StaticRoute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchRoutingTableIpv6StaticRouteNexthop(a.Nexthop, b.Nexthop) {
				return false
			}
			if !matchRoutingTableIpv6StaticRouteOption(a.Option, b.Option) {
				return false
			}
			if !matchRoutingTableIpv6StaticRouteRouteTable(a.RouteTable, b.RouteTable) {
				return false
			}
			if !matchRoutingTableIpv6StaticRoutePathMonitor(a.PathMonitor, b.PathMonitor) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !matchRoutingTableIpv6StaticRouteBfd(a.Bfd, b.Bfd) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
		}
	}
	return true
}
func matchRoutingTableIpv6(a *RoutingTableIpv6, b *RoutingTableIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpv6StaticRoute(a.StaticRoute, b.StaticRoute) {
		return false
	}
	return true
}
func matchRoutingTable(a *RoutingTable, b *RoutingTable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIp(a.Ip, b.Ip) {
		return false
	}
	if !matchRoutingTableIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchAdminDists(a *AdminDists, b *AdminDists) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.OspfExt, b.OspfExt) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Ext, b.Ospfv3Ext) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Int, b.Ospfv3Int) {
		return false
	}
	if !util.Ints64Match(a.Rip, b.Rip) {
		return false
	}
	if !util.Ints64Match(a.StaticIpv6, b.StaticIpv6) {
		return false
	}
	if !util.Ints64Match(a.Ebgp, b.Ebgp) {
		return false
	}
	if !util.Ints64Match(a.Ibgp, b.Ibgp) {
		return false
	}
	if !util.Ints64Match(a.OspfInt, b.OspfInt) {
		return false
	}
	if !util.Ints64Match(a.Static, b.Static) {
		return false
	}
	return true
}
func matchEcmpAlgorithmWeightedRoundRobinInterface(a []EcmpAlgorithmWeightedRoundRobinInterface, b []EcmpAlgorithmWeightedRoundRobinInterface) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Weight, b.Weight) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchEcmpAlgorithmWeightedRoundRobin(a *EcmpAlgorithmWeightedRoundRobin, b *EcmpAlgorithmWeightedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchEcmpAlgorithmWeightedRoundRobinInterface(a.Interface, b.Interface) {
		return false
	}
	return true
}
func matchEcmpAlgorithmBalancedRoundRobin(a *EcmpAlgorithmBalancedRoundRobin, b *EcmpAlgorithmBalancedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchEcmpAlgorithmIpHash(a *EcmpAlgorithmIpHash, b *EcmpAlgorithmIpHash) bool {
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
func matchEcmpAlgorithmIpModulo(a *EcmpAlgorithmIpModulo, b *EcmpAlgorithmIpModulo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchEcmpAlgorithm(a *EcmpAlgorithm, b *EcmpAlgorithm) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchEcmpAlgorithmBalancedRoundRobin(a.BalancedRoundRobin, b.BalancedRoundRobin) {
		return false
	}
	if !matchEcmpAlgorithmIpHash(a.IpHash, b.IpHash) {
		return false
	}
	if !matchEcmpAlgorithmIpModulo(a.IpModulo, b.IpModulo) {
		return false
	}
	if !matchEcmpAlgorithmWeightedRoundRobin(a.WeightedRoundRobin, b.WeightedRoundRobin) {
		return false
	}
	return true
}
func matchEcmp(a *Ecmp, b *Ecmp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.MaxPath, b.MaxPath) {
		return false
	}
	if !util.BoolsMatch(a.StrictSourcePath, b.StrictSourcePath) {
		return false
	}
	if !util.BoolsMatch(a.SymmetricReturn, b.SymmetricReturn) {
		return false
	}
	if !matchEcmpAlgorithm(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
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
