package aggregate

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
	suffix = []string{"network", "interface", "aggregate-ethernet", "$name"}
)

type Entry struct {
	Name          string
	Comment       *string
	DecryptMirror *DecryptMirror
	Ha            *Ha
	Layer2        *Layer2
	Layer3        *Layer3
	VirtualWire   *VirtualWire
	Misc          []generic.Xml
}
type DecryptMirror struct {
	Misc []generic.Xml
}
type Ha struct {
	Lacp *HaLacp
	Misc []generic.Xml
}
type HaLacp struct {
	Enable           *bool
	FastFailover     *bool
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
	Misc             []generic.Xml
}
type Layer2 struct {
	Lacp           *Layer2Lacp
	Lldp           *Layer2Lldp
	NetflowProfile *string
	Misc           []generic.Xml
}
type Layer2Lacp struct {
	Enable           *bool
	FastFailover     *bool
	HighAvailability *Layer2LacpHighAvailability
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
	Misc             []generic.Xml
}
type Layer2LacpHighAvailability struct {
	PassivePreNegotiation *bool
	Misc                  []generic.Xml
}
type Layer2Lldp struct {
	Enable           *bool
	HighAvailability *Layer2LldpHighAvailability
	Profile          *string
	Misc             []generic.Xml
}
type Layer2LldpHighAvailability struct {
	PassivePreNegotiation *bool
	Misc                  []generic.Xml
}
type Layer3 struct {
	AdjustTcpMss               *Layer3AdjustTcpMss
	Arp                        []Layer3Arp
	Bonjour                    *Layer3Bonjour
	DdnsConfig                 *Layer3DdnsConfig
	DecryptForward             *bool
	DfIgnore                   *bool
	DhcpClient                 *Layer3DhcpClient
	InterfaceManagementProfile *string
	Ip                         []Layer3Ip
	Ipv6                       *Layer3Ipv6
	Lacp                       *Layer3Lacp
	Lldp                       *Layer3Lldp
	Mtu                        *int64
	NdpProxy                   *Layer3NdpProxy
	NetflowProfile             *string
	SdwanLinkSettings          *Layer3SdwanLinkSettings
	UntaggedSubInterface       *bool
	Misc                       []generic.Xml
}
type Layer3AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
	Misc              []generic.Xml
}
type Layer3Arp struct {
	Name      string
	HwAddress *string
	Misc      []generic.Xml
}
type Layer3Bonjour struct {
	Enable   *bool
	GroupId  *int64
	TtlCheck *bool
	Misc     []generic.Xml
}
type Layer3DdnsConfig struct {
	DdnsCertProfile    *string
	DdnsEnabled        *bool
	DdnsHostname       *string
	DdnsIp             []string
	DdnsIpv6           []string
	DdnsUpdateInterval *int64
	DdnsVendor         *string
	DdnsVendorConfig   []Layer3DdnsConfigDdnsVendorConfig
	Misc               []generic.Xml
}
type Layer3DdnsConfigDdnsVendorConfig struct {
	Name  string
	Value *string
	Misc  []generic.Xml
}
type Layer3DhcpClient struct {
	CreateDefaultRoute *bool
	DefaultRouteMetric *int64
	Enable             *bool
	SendHostname       *Layer3DhcpClientSendHostname
	Misc               []generic.Xml
}
type Layer3DhcpClientSendHostname struct {
	Enable   *bool
	Hostname *string
	Misc     []generic.Xml
}
type Layer3Ip struct {
	Name         string
	SdwanGateway *string
	Misc         []generic.Xml
}
type Layer3Ipv6 struct {
	Address           []Layer3Ipv6Address
	DhcpClient        *Layer3Ipv6DhcpClient
	Enabled           *bool
	Inherited         *Layer3Ipv6Inherited
	InterfaceId       *string
	NeighborDiscovery *Layer3Ipv6NeighborDiscovery
	Misc              []generic.Xml
}
type Layer3Ipv6Address struct {
	Name              string
	EnableOnInterface *bool
	Prefix            *Layer3Ipv6AddressPrefix
	Anycast           *Layer3Ipv6AddressAnycast
	Advertise         *Layer3Ipv6AddressAdvertise
	Misc              []generic.Xml
}
type Layer3Ipv6AddressPrefix struct {
	Misc []generic.Xml
}
type Layer3Ipv6AddressAnycast struct {
	Misc []generic.Xml
}
type Layer3Ipv6AddressAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
	Misc              []generic.Xml
}
type Layer3Ipv6DhcpClient struct {
	AcceptRaRoute      *bool
	DefaultRouteMetric *int64
	Enable             *bool
	NeighborDiscovery  *Layer3Ipv6DhcpClientNeighborDiscovery
	Preference         *string
	PrefixDelegation   *Layer3Ipv6DhcpClientPrefixDelegation
	V6Options          *Layer3Ipv6DhcpClientV6Options
	Misc               []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscovery struct {
	DadAttempts      *int64
	DnsServer        *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer
	DnsSuffix        *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	EnableDad        *bool
	EnableNdpMonitor *bool
	Neighbor         []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor
	NsInterval       *int64
	ReachableTime    *int64
	Misc             []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual struct {
	Server []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
	Misc      []generic.Xml
}
type Layer3Ipv6DhcpClientPrefixDelegation struct {
	Enable *Layer3Ipv6DhcpClientPrefixDelegationEnable
	Misc   []generic.Xml
}
type Layer3Ipv6DhcpClientPrefixDelegationEnable struct {
	No   *Layer3Ipv6DhcpClientPrefixDelegationEnableNo
	Yes  *Layer3Ipv6DhcpClientPrefixDelegationEnableYes
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableNo struct {
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableYes struct {
	PfxPoolName   *string
	PrefixLen     *int64
	PrefixLenHint *bool
	Misc          []generic.Xml
}
type Layer3Ipv6DhcpClientV6Options struct {
	DuidType            *string
	Enable              *Layer3Ipv6DhcpClientV6OptionsEnable
	RapidCommit         *bool
	SupportSrvrReconfig *bool
	Misc                []generic.Xml
}
type Layer3Ipv6DhcpClientV6OptionsEnable struct {
	No   *Layer3Ipv6DhcpClientV6OptionsEnableNo
	Yes  *Layer3Ipv6DhcpClientV6OptionsEnableYes
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientV6OptionsEnableNo struct {
	Misc []generic.Xml
}
type Layer3Ipv6DhcpClientV6OptionsEnableYes struct {
	NonTempAddr *bool
	TempAddr    *bool
	Misc        []generic.Xml
}
type Layer3Ipv6Inherited struct {
	AssignAddr        []Layer3Ipv6InheritedAssignAddr
	Enable            *bool
	NeighborDiscovery *Layer3Ipv6InheritedNeighborDiscovery
	Misc              []generic.Xml
}
type Layer3Ipv6InheritedAssignAddr struct {
	Name string
	Type *Layer3Ipv6InheritedAssignAddrType
	Misc []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrType struct {
	Gua  *Layer3Ipv6InheritedAssignAddrTypeGua
	Ula  *Layer3Ipv6InheritedAssignAddrTypeUla
	Misc []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeGua struct {
	EnableOnInterface *bool
	PrefixPool        *string
	PoolType          *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise
	Misc              []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolType struct {
	Dynamic   *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	DynamicId *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	Misc      []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic struct {
	Misc []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId struct {
	Identifier *int64
	Misc       []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise struct {
	Enable         *bool
	OnlinkFlag     *bool
	AutoConfigFlag *bool
	Misc           []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeUla struct {
	EnableOnInterface *bool
	Address           *string
	Prefix            *bool
	Anycast           *bool
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise
	Misc              []generic.Xml
}
type Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
	Misc              []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscovery struct {
	DadAttempts         *int64
	DnsServer           *Layer3Ipv6InheritedNeighborDiscoveryDnsServer
	DnsSuffix           *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Layer3Ipv6InheritedNeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement
	Misc                []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	PrefixPool *string
	Misc       []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual struct {
	Server []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	PrefixPool *string
	Misc       []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	Misc   []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
	Misc      []generic.Xml
}
type Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement struct {
	Enable                 *bool
	EnableConsistencyCheck *bool
	HopLimit               *string
	Lifetime               *int64
	LinkMtu                *string
	ManagedFlag            *bool
	MaxInterval            *int64
	MinInterval            *int64
	OtherFlag              *bool
	ReachableTime          *string
	RetransmissionTimer    *string
	RouterPreference       *string
	Misc                   []generic.Xml
}
type Layer3Ipv6NeighborDiscovery struct {
	DadAttempts         *int64
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Layer3Ipv6NeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Layer3Ipv6NeighborDiscoveryRouterAdvertisement
	Misc                []generic.Xml
}
type Layer3Ipv6NeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
	Misc      []generic.Xml
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisement struct {
	DnsSupport             *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
	Enable                 *bool
	EnableConsistencyCheck *bool
	HopLimit               *string
	Lifetime               *int64
	LinkMtu                *string
	ManagedFlag            *bool
	MaxInterval            *int64
	MinInterval            *int64
	OtherFlag              *bool
	ReachableTime          *string
	RetransmissionTimer    *string
	RouterPreference       *string
	Misc                   []generic.Xml
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport struct {
	Enable *bool
	Server []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	Suffix []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	Misc   []generic.Xml
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix struct {
	Name     string
	Lifetime *int64
	Misc     []generic.Xml
}
type Layer3Lacp struct {
	Enable           *bool
	FastFailover     *bool
	HighAvailability *Layer3LacpHighAvailability
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
	Misc             []generic.Xml
}
type Layer3LacpHighAvailability struct {
	PassivePreNegotiation *bool
	Misc                  []generic.Xml
}
type Layer3Lldp struct {
	Enable           *bool
	HighAvailability *Layer3LldpHighAvailability
	Profile          *string
	Misc             []generic.Xml
}
type Layer3LldpHighAvailability struct {
	PassivePreNegotiation *bool
	Misc                  []generic.Xml
}
type Layer3NdpProxy struct {
	Address []Layer3NdpProxyAddress
	Enabled *bool
	Misc    []generic.Xml
}
type Layer3NdpProxyAddress struct {
	Name   string
	Negate *bool
	Misc   []generic.Xml
}
type Layer3SdwanLinkSettings struct {
	Enable                *bool
	SdwanInterfaceProfile *string
	UpstreamNat           *Layer3SdwanLinkSettingsUpstreamNat
	Misc                  []generic.Xml
}
type Layer3SdwanLinkSettingsUpstreamNat struct {
	Enable   *bool
	Ddns     *Layer3SdwanLinkSettingsUpstreamNatDdns
	StaticIp *Layer3SdwanLinkSettingsUpstreamNatStaticIp
	Misc     []generic.Xml
}
type Layer3SdwanLinkSettingsUpstreamNatDdns struct {
	Misc []generic.Xml
}
type Layer3SdwanLinkSettingsUpstreamNatStaticIp struct {
	Fqdn      *string
	IpAddress *string
	Misc      []generic.Xml
}
type VirtualWire struct {
	Lldp           *VirtualWireLldp
	NetflowProfile *string
	Misc           []generic.Xml
}
type VirtualWireLldp struct {
	Enable           *bool
	HighAvailability *VirtualWireLldpHighAvailability
	Profile          *string
	Misc             []generic.Xml
}
type VirtualWireLldpHighAvailability struct {
	PassivePreNegotiation *bool
	Misc                  []generic.Xml
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
	XMLName       xml.Name          `xml:"entry"`
	Name          string            `xml:"name,attr"`
	Comment       *string           `xml:"comment,omitempty"`
	DecryptMirror *decryptMirrorXml `xml:"decrypt-mirror,omitempty"`
	Ha            *haXml            `xml:"ha,omitempty"`
	Layer2        *layer2Xml        `xml:"layer2,omitempty"`
	Layer3        *layer3Xml        `xml:"layer3,omitempty"`
	VirtualWire   *virtualWireXml   `xml:"virtual-wire,omitempty"`
	Misc          []generic.Xml     `xml:",any"`
}
type decryptMirrorXml struct {
	Misc []generic.Xml `xml:",any"`
}
type haXml struct {
	Lacp *haLacpXml    `xml:"lacp,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type haLacpXml struct {
	Enable           *string       `xml:"enable,omitempty"`
	FastFailover     *string       `xml:"fast-failover,omitempty"`
	MaxPorts         *int64        `xml:"max-ports,omitempty"`
	Mode             *string       `xml:"mode,omitempty"`
	SystemPriority   *int64        `xml:"system-priority,omitempty"`
	TransmissionRate *string       `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type layer2Xml struct {
	Lacp           *layer2LacpXml `xml:"lacp,omitempty"`
	Lldp           *layer2LldpXml `xml:"lldp,omitempty"`
	NetflowProfile *string        `xml:"netflow-profile,omitempty"`
	Misc           []generic.Xml  `xml:",any"`
}
type layer2LacpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	FastFailover     *string                        `xml:"fast-failover,omitempty"`
	HighAvailability *layer2LacpHighAvailabilityXml `xml:"high-availability,omitempty"`
	MaxPorts         *int64                         `xml:"max-ports,omitempty"`
	Mode             *string                        `xml:"mode,omitempty"`
	SystemPriority   *int64                         `xml:"system-priority,omitempty"`
	TransmissionRate *string                        `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml                  `xml:",any"`
}
type layer2LacpHighAvailabilityXml struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer2LldpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	HighAvailability *layer2LldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                        `xml:"profile,omitempty"`
	Misc             []generic.Xml                  `xml:",any"`
}
type layer2LldpHighAvailabilityXml struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3Xml struct {
	AdjustTcpMss               *layer3AdjustTcpMssXml      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        *layer3ArpContainerXml      `xml:"arp,omitempty"`
	Bonjour                    *layer3BonjourXml           `xml:"bonjour,omitempty"`
	DdnsConfig                 *layer3DdnsConfigXml        `xml:"ddns-config,omitempty"`
	DecryptForward             *string                     `xml:"decrypt-forward,omitempty"`
	DfIgnore                   *string                     `xml:"df-ignore,omitempty"`
	DhcpClient                 *layer3DhcpClientXml        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                     `xml:"interface-management-profile,omitempty"`
	Ip                         *layer3IpContainerXml       `xml:"ip,omitempty"`
	Ipv6                       *layer3Ipv6Xml              `xml:"ipv6,omitempty"`
	Lacp                       *layer3LacpXml              `xml:"lacp,omitempty"`
	Lldp                       *layer3LldpXml              `xml:"lldp,omitempty"`
	Mtu                        *int64                      `xml:"mtu,omitempty"`
	NdpProxy                   *layer3NdpProxyXml          `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string                     `xml:"netflow-profile,omitempty"`
	SdwanLinkSettings          *layer3SdwanLinkSettingsXml `xml:"sdwan-link-settings,omitempty"`
	UntaggedSubInterface       *string                     `xml:"untagged-sub-interface,omitempty"`
	Misc                       []generic.Xml               `xml:",any"`
}
type layer3AdjustTcpMssXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64        `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64        `xml:"ipv6-mss-adjustment,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3ArpContainerXml struct {
	Entries []layer3ArpXml `xml:"entry"`
}
type layer3ArpXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3BonjourXml struct {
	Enable   *string       `xml:"enable,omitempty"`
	GroupId  *int64        `xml:"group-id,omitempty"`
	TtlCheck *string       `xml:"ttl-check,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3DdnsConfigXml struct {
	DdnsCertProfile    *string                                       `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                                       `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                                       `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                              `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                              `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                        `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                                       `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   *layer3DdnsConfigDdnsVendorConfigContainerXml `xml:"ddns-vendor-config,omitempty"`
	Misc               []generic.Xml                                 `xml:",any"`
}
type layer3DdnsConfigDdnsVendorConfigContainerXml struct {
	Entries []layer3DdnsConfigDdnsVendorConfigXml `xml:"entry"`
}
type layer3DdnsConfigDdnsVendorConfigXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Value   *string       `xml:"value,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type layer3DhcpClientXml struct {
	CreateDefaultRoute *string                          `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                           `xml:"default-route-metric,omitempty"`
	Enable             *string                          `xml:"enable,omitempty"`
	SendHostname       *layer3DhcpClientSendHostnameXml `xml:"send-hostname,omitempty"`
	Misc               []generic.Xml                    `xml:",any"`
}
type layer3DhcpClientSendHostnameXml struct {
	Enable   *string       `xml:"enable,omitempty"`
	Hostname *string       `xml:"hostname,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3IpContainerXml struct {
	Entries []layer3IpXml `xml:"entry"`
}
type layer3IpXml struct {
	XMLName      xml.Name      `xml:"entry"`
	Name         string        `xml:"name,attr"`
	SdwanGateway *string       `xml:"sdwan-gateway,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type layer3Ipv6Xml struct {
	Address           *layer3Ipv6AddressContainerXml  `xml:"address,omitempty"`
	DhcpClient        *layer3Ipv6DhcpClientXml        `xml:"dhcp-client,omitempty"`
	Enabled           *string                         `xml:"enabled,omitempty"`
	Inherited         *layer3Ipv6InheritedXml         `xml:"inherited,omitempty"`
	InterfaceId       *string                         `xml:"interface-id,omitempty"`
	NeighborDiscovery *layer3Ipv6NeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                   `xml:",any"`
}
type layer3Ipv6AddressContainerXml struct {
	Entries []layer3Ipv6AddressXml `xml:"entry"`
}
type layer3Ipv6AddressXml struct {
	XMLName           xml.Name                       `xml:"entry"`
	Name              string                         `xml:"name,attr"`
	EnableOnInterface *string                        `xml:"enable-on-interface,omitempty"`
	Prefix            *layer3Ipv6AddressPrefixXml    `xml:"prefix,omitempty"`
	Anycast           *layer3Ipv6AddressAnycastXml   `xml:"anycast,omitempty"`
	Advertise         *layer3Ipv6AddressAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml                  `xml:",any"`
}
type layer3Ipv6AddressPrefixXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6AddressAnycastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6AddressAdvertiseXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientXml struct {
	AcceptRaRoute      *string                                   `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                                    `xml:"default-route-metric,omitempty"`
	Enable             *string                                   `xml:"enable,omitempty"`
	NeighborDiscovery  *layer3Ipv6DhcpClientNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Preference         *string                                   `xml:"preference,omitempty"`
	PrefixDelegation   *layer3Ipv6DhcpClientPrefixDelegationXml  `xml:"prefix-delegation,omitempty"`
	V6Options          *layer3Ipv6DhcpClientV6OptionsXml         `xml:"v6-options,omitempty"`
	Misc               []generic.Xml                             `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryXml struct {
	DadAttempts      *int64                                                     `xml:"dad-attempts,omitempty"`
	DnsServer        *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml         `xml:"dns-server,omitempty"`
	DnsSuffix        *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml         `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                                    `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                                    `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         *layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml `xml:"neighbor,omitempty"`
	NsInterval       *int64                                                     `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                                     `xml:"reachable-time,omitempty"`
	Misc             []generic.Xml                                              `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`
	Misc   []generic.Xml                                            `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                  `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml struct {
	Server *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml `xml:"server,omitempty"`
	Misc   []generic.Xml                                                                 `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`
	Misc   []generic.Xml                                            `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                  `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                                 `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationXml struct {
	Enable *layer3Ipv6DhcpClientPrefixDelegationEnableXml `xml:"enable,omitempty"`
	Misc   []generic.Xml                                  `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableXml struct {
	No   *layer3Ipv6DhcpClientPrefixDelegationEnableNoXml  `xml:"no,omitempty"`
	Yes  *layer3Ipv6DhcpClientPrefixDelegationEnableYesXml `xml:"yes,omitempty"`
	Misc []generic.Xml                                     `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableYesXml struct {
	PfxPoolName   *string       `xml:"pfx-pool-name,omitempty"`
	PrefixLen     *int64        `xml:"prefix-len,omitempty"`
	PrefixLenHint *string       `xml:"prefix-len-hint,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsXml struct {
	DuidType            *string                                 `xml:"duid-type,omitempty"`
	Enable              *layer3Ipv6DhcpClientV6OptionsEnableXml `xml:"enable,omitempty"`
	RapidCommit         *string                                 `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                                 `xml:"support-srvr-reconfig,omitempty"`
	Misc                []generic.Xml                           `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableXml struct {
	No   *layer3Ipv6DhcpClientV6OptionsEnableNoXml  `xml:"no,omitempty"`
	Yes  *layer3Ipv6DhcpClientV6OptionsEnableYesXml `xml:"yes,omitempty"`
	Misc []generic.Xml                              `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableYesXml struct {
	NonTempAddr *string       `xml:"non-temp-addr,omitempty"`
	TempAddr    *string       `xml:"temp-addr,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedXml struct {
	AssignAddr        *layer3Ipv6InheritedAssignAddrContainerXml `xml:"assign-addr,omitempty"`
	Enable            *string                                    `xml:"enable,omitempty"`
	NeighborDiscovery *layer3Ipv6InheritedNeighborDiscoveryXml   `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                              `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrContainerXml struct {
	Entries []layer3Ipv6InheritedAssignAddrXml `xml:"entry"`
}
type layer3Ipv6InheritedAssignAddrXml struct {
	XMLName xml.Name                              `xml:"entry"`
	Name    string                                `xml:"name,attr"`
	Type    *layer3Ipv6InheritedAssignAddrTypeXml `xml:"type,omitempty"`
	Misc    []generic.Xml                         `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeXml struct {
	Gua  *layer3Ipv6InheritedAssignAddrTypeGuaXml `xml:"gua,omitempty"`
	Ula  *layer3Ipv6InheritedAssignAddrTypeUlaXml `xml:"ula,omitempty"`
	Misc []generic.Xml                            `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaXml struct {
	EnableOnInterface *string                                           `xml:"enable-on-interface,omitempty"`
	PrefixPool        *string                                           `xml:"prefix-pool,omitempty"`
	PoolType          *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml  `xml:"pool-type,omitempty"`
	Advertise         *layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                     `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml struct {
	Dynamic   *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml   `xml:"dynamic,omitempty"`
	DynamicId *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml `xml:"dynamic-id,omitempty"`
	Misc      []generic.Xml                                             `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml struct {
	Identifier *int64        `xml:"identifier,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	OnlinkFlag     *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag *string       `xml:"auto-config-flag,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeUlaXml struct {
	EnableOnInterface *string                                           `xml:"enable-on-interface,omitempty"`
	Address           *string                                           `xml:"address,omitempty"`
	Prefix            *string                                           `xml:"prefix,omitempty"`
	Anycast           *string                                           `xml:"anycast,omitempty"`
	Advertise         *layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                     `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryXml struct {
	DadAttempts         *int64                                                      `xml:"dad-attempts,omitempty"`
	DnsServer           *layer3Ipv6InheritedNeighborDiscoveryDnsServerXml           `xml:"dns-server,omitempty"`
	DnsSuffix           *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                                     `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                     `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                                      `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                      `xml:"reachable-time,omitempty"`
	RouterAdvertisement *layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                               `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerXml struct {
	Enable *string                                                 `xml:"enable,omitempty"`
	Source *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`
	Misc   []generic.Xml                                           `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                 `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	PrefixPool *string       `xml:"prefix-pool,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml struct {
	Server *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml `xml:"server,omitempty"`
	Misc   []generic.Xml                                                                `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                                 `xml:"enable,omitempty"`
	Source *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`
	Misc   []generic.Xml                                           `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                 `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	PrefixPool *string       `xml:"prefix-pool,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                                `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryNeighborXml `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryNeighborXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml struct {
	Enable                 *string       `xml:"enable,omitempty"`
	EnableConsistencyCheck *string       `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string       `xml:"hop-limit,omitempty"`
	Lifetime               *int64        `xml:"lifetime,omitempty"`
	LinkMtu                *string       `xml:"link-mtu,omitempty"`
	ManagedFlag            *string       `xml:"managed-flag,omitempty"`
	MaxInterval            *int64        `xml:"max-interval,omitempty"`
	MinInterval            *int64        `xml:"min-interval,omitempty"`
	OtherFlag              *string       `xml:"other-flag,omitempty"`
	ReachableTime          *string       `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string       `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string       `xml:"router-preference,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryXml struct {
	DadAttempts         *int64                                             `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                            `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                            `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *layer3Ipv6NeighborDiscoveryNeighborContainerXml   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                             `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                             `xml:"reachable-time,omitempty"`
	RouterAdvertisement *layer3Ipv6NeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                      `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryNeighborContainerXml struct {
	Entries []layer3Ipv6NeighborDiscoveryNeighborXml `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryNeighborXml struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementXml struct {
	DnsSupport             *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml `xml:"dns-support,omitempty"`
	Enable                 *string                                                      `xml:"enable,omitempty"`
	EnableConsistencyCheck *string                                                      `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string                                                      `xml:"hop-limit,omitempty"`
	Lifetime               *int64                                                       `xml:"lifetime,omitempty"`
	LinkMtu                *string                                                      `xml:"link-mtu,omitempty"`
	ManagedFlag            *string                                                      `xml:"managed-flag,omitempty"`
	MaxInterval            *int64                                                       `xml:"max-interval,omitempty"`
	MinInterval            *int64                                                       `xml:"min-interval,omitempty"`
	OtherFlag              *string                                                      `xml:"other-flag,omitempty"`
	ReachableTime          *string                                                      `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string                                                      `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string                                                      `xml:"router-preference,omitempty"`
	Misc                   []generic.Xml                                                `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml struct {
	Enable *string                                                                     `xml:"enable,omitempty"`
	Server *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml `xml:"server,omitempty"`
	Suffix *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                               `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml struct {
	Entries []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml struct {
	Entries []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3LacpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	FastFailover     *string                        `xml:"fast-failover,omitempty"`
	HighAvailability *layer3LacpHighAvailabilityXml `xml:"high-availability,omitempty"`
	MaxPorts         *int64                         `xml:"max-ports,omitempty"`
	Mode             *string                        `xml:"mode,omitempty"`
	SystemPriority   *int64                         `xml:"system-priority,omitempty"`
	TransmissionRate *string                        `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml                  `xml:",any"`
}
type layer3LacpHighAvailabilityXml struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3LldpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	HighAvailability *layer3LldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                        `xml:"profile,omitempty"`
	Misc             []generic.Xml                  `xml:",any"`
}
type layer3LldpHighAvailabilityXml struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3NdpProxyXml struct {
	Address *layer3NdpProxyAddressContainerXml `xml:"address,omitempty"`
	Enabled *string                            `xml:"enabled,omitempty"`
	Misc    []generic.Xml                      `xml:",any"`
}
type layer3NdpProxyAddressContainerXml struct {
	Entries []layer3NdpProxyAddressXml `xml:"entry"`
}
type layer3NdpProxyAddressXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Negate  *string       `xml:"negate,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type layer3SdwanLinkSettingsXml struct {
	Enable                *string                                `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                                `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *layer3SdwanLinkSettingsUpstreamNatXml `xml:"upstream-nat,omitempty"`
	Misc                  []generic.Xml                          `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatXml struct {
	Enable   *string                                        `xml:"enable,omitempty"`
	Ddns     *layer3SdwanLinkSettingsUpstreamNatDdnsXml     `xml:"ddns,omitempty"`
	StaticIp *layer3SdwanLinkSettingsUpstreamNatStaticIpXml `xml:"static-ip,omitempty"`
	Misc     []generic.Xml                                  `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatDdnsXml struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatStaticIpXml struct {
	Fqdn      *string       `xml:"fqdn,omitempty"`
	IpAddress *string       `xml:"ip-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type virtualWireXml struct {
	Lldp           *virtualWireLldpXml `xml:"lldp,omitempty"`
	NetflowProfile *string             `xml:"netflow-profile,omitempty"`
	Misc           []generic.Xml       `xml:",any"`
}
type virtualWireLldpXml struct {
	Enable           *string                             `xml:"enable,omitempty"`
	HighAvailability *virtualWireLldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                             `xml:"profile,omitempty"`
	Misc             []generic.Xml                       `xml:",any"`
}
type virtualWireLldpHighAvailabilityXml struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName       xml.Name                 `xml:"entry"`
	Name          string                   `xml:"name,attr"`
	Comment       *string                  `xml:"comment,omitempty"`
	DecryptMirror *decryptMirrorXml_11_0_2 `xml:"decrypt-mirror,omitempty"`
	Ha            *haXml_11_0_2            `xml:"ha,omitempty"`
	Layer2        *layer2Xml_11_0_2        `xml:"layer2,omitempty"`
	Layer3        *layer3Xml_11_0_2        `xml:"layer3,omitempty"`
	VirtualWire   *virtualWireXml_11_0_2   `xml:"virtual-wire,omitempty"`
	Misc          []generic.Xml            `xml:",any"`
}
type decryptMirrorXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type haXml_11_0_2 struct {
	Lacp *haLacpXml_11_0_2 `xml:"lacp,omitempty"`
	Misc []generic.Xml     `xml:",any"`
}
type haLacpXml_11_0_2 struct {
	Enable           *string       `xml:"enable,omitempty"`
	FastFailover     *string       `xml:"fast-failover,omitempty"`
	MaxPorts         *int64        `xml:"max-ports,omitempty"`
	Mode             *string       `xml:"mode,omitempty"`
	SystemPriority   *int64        `xml:"system-priority,omitempty"`
	TransmissionRate *string       `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type layer2Xml_11_0_2 struct {
	Lacp           *layer2LacpXml_11_0_2 `xml:"lacp,omitempty"`
	Lldp           *layer2LldpXml_11_0_2 `xml:"lldp,omitempty"`
	NetflowProfile *string               `xml:"netflow-profile,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
}
type layer2LacpXml_11_0_2 struct {
	Enable           *string                               `xml:"enable,omitempty"`
	FastFailover     *string                               `xml:"fast-failover,omitempty"`
	HighAvailability *layer2LacpHighAvailabilityXml_11_0_2 `xml:"high-availability,omitempty"`
	MaxPorts         *int64                                `xml:"max-ports,omitempty"`
	Mode             *string                               `xml:"mode,omitempty"`
	SystemPriority   *int64                                `xml:"system-priority,omitempty"`
	TransmissionRate *string                               `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml                         `xml:",any"`
}
type layer2LacpHighAvailabilityXml_11_0_2 struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer2LldpXml_11_0_2 struct {
	Enable           *string                               `xml:"enable,omitempty"`
	HighAvailability *layer2LldpHighAvailabilityXml_11_0_2 `xml:"high-availability,omitempty"`
	Profile          *string                               `xml:"profile,omitempty"`
	Misc             []generic.Xml                         `xml:",any"`
}
type layer2LldpHighAvailabilityXml_11_0_2 struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3Xml_11_0_2 struct {
	AdjustTcpMss               *layer3AdjustTcpMssXml_11_0_2      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        *layer3ArpContainerXml_11_0_2      `xml:"arp,omitempty"`
	Bonjour                    *layer3BonjourXml_11_0_2           `xml:"bonjour,omitempty"`
	DdnsConfig                 *layer3DdnsConfigXml_11_0_2        `xml:"ddns-config,omitempty"`
	DecryptForward             *string                            `xml:"decrypt-forward,omitempty"`
	DfIgnore                   *string                            `xml:"df-ignore,omitempty"`
	DhcpClient                 *layer3DhcpClientXml_11_0_2        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                            `xml:"interface-management-profile,omitempty"`
	Ip                         *layer3IpContainerXml_11_0_2       `xml:"ip,omitempty"`
	Ipv6                       *layer3Ipv6Xml_11_0_2              `xml:"ipv6,omitempty"`
	Lacp                       *layer3LacpXml_11_0_2              `xml:"lacp,omitempty"`
	Lldp                       *layer3LldpXml_11_0_2              `xml:"lldp,omitempty"`
	Mtu                        *int64                             `xml:"mtu,omitempty"`
	NdpProxy                   *layer3NdpProxyXml_11_0_2          `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string                            `xml:"netflow-profile,omitempty"`
	SdwanLinkSettings          *layer3SdwanLinkSettingsXml_11_0_2 `xml:"sdwan-link-settings,omitempty"`
	UntaggedSubInterface       *string                            `xml:"untagged-sub-interface,omitempty"`
	Misc                       []generic.Xml                      `xml:",any"`
}
type layer3AdjustTcpMssXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64        `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64        `xml:"ipv6-mss-adjustment,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3ArpContainerXml_11_0_2 struct {
	Entries []layer3ArpXml_11_0_2 `xml:"entry"`
}
type layer3ArpXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3BonjourXml_11_0_2 struct {
	Enable   *string       `xml:"enable,omitempty"`
	GroupId  *int64        `xml:"group-id,omitempty"`
	TtlCheck *string       `xml:"ttl-check,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3DdnsConfigXml_11_0_2 struct {
	DdnsCertProfile    *string                                              `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                                              `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                                              `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                                     `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                                     `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                               `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                                              `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   *layer3DdnsConfigDdnsVendorConfigContainerXml_11_0_2 `xml:"ddns-vendor-config,omitempty"`
	Misc               []generic.Xml                                        `xml:",any"`
}
type layer3DdnsConfigDdnsVendorConfigContainerXml_11_0_2 struct {
	Entries []layer3DdnsConfigDdnsVendorConfigXml_11_0_2 `xml:"entry"`
}
type layer3DdnsConfigDdnsVendorConfigXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Value   *string       `xml:"value,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type layer3DhcpClientXml_11_0_2 struct {
	CreateDefaultRoute *string                                 `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                                  `xml:"default-route-metric,omitempty"`
	Enable             *string                                 `xml:"enable,omitempty"`
	SendHostname       *layer3DhcpClientSendHostnameXml_11_0_2 `xml:"send-hostname,omitempty"`
	Misc               []generic.Xml                           `xml:",any"`
}
type layer3DhcpClientSendHostnameXml_11_0_2 struct {
	Enable   *string       `xml:"enable,omitempty"`
	Hostname *string       `xml:"hostname,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3IpContainerXml_11_0_2 struct {
	Entries []layer3IpXml_11_0_2 `xml:"entry"`
}
type layer3IpXml_11_0_2 struct {
	XMLName      xml.Name      `xml:"entry"`
	Name         string        `xml:"name,attr"`
	SdwanGateway *string       `xml:"sdwan-gateway,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type layer3Ipv6Xml_11_0_2 struct {
	Address           *layer3Ipv6AddressContainerXml_11_0_2  `xml:"address,omitempty"`
	DhcpClient        *layer3Ipv6DhcpClientXml_11_0_2        `xml:"dhcp-client,omitempty"`
	Enabled           *string                                `xml:"enabled,omitempty"`
	Inherited         *layer3Ipv6InheritedXml_11_0_2         `xml:"inherited,omitempty"`
	InterfaceId       *string                                `xml:"interface-id,omitempty"`
	NeighborDiscovery *layer3Ipv6NeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                          `xml:",any"`
}
type layer3Ipv6AddressContainerXml_11_0_2 struct {
	Entries []layer3Ipv6AddressXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6AddressXml_11_0_2 struct {
	XMLName           xml.Name                              `xml:"entry"`
	Name              string                                `xml:"name,attr"`
	EnableOnInterface *string                               `xml:"enable-on-interface,omitempty"`
	Prefix            *layer3Ipv6AddressPrefixXml_11_0_2    `xml:"prefix,omitempty"`
	Anycast           *layer3Ipv6AddressAnycastXml_11_0_2   `xml:"anycast,omitempty"`
	Advertise         *layer3Ipv6AddressAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                         `xml:",any"`
}
type layer3Ipv6AddressPrefixXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6AddressAnycastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6AddressAdvertiseXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientXml_11_0_2 struct {
	AcceptRaRoute      *string                                          `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                                           `xml:"default-route-metric,omitempty"`
	Enable             *string                                          `xml:"enable,omitempty"`
	NeighborDiscovery  *layer3Ipv6DhcpClientNeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`
	Preference         *string                                          `xml:"preference,omitempty"`
	PrefixDelegation   *layer3Ipv6DhcpClientPrefixDelegationXml_11_0_2  `xml:"prefix-delegation,omitempty"`
	V6Options          *layer3Ipv6DhcpClientV6OptionsXml_11_0_2         `xml:"v6-options,omitempty"`
	Misc               []generic.Xml                                    `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts      *int64                                                            `xml:"dad-attempts,omitempty"`
	DnsServer        *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2         `xml:"dns-server,omitempty"`
	DnsSuffix        *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2         `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                                           `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                                           `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         *layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2 `xml:"neighbor,omitempty"`
	NsInterval       *int64                                                            `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                                            `xml:"reachable-time,omitempty"`
	Misc             []generic.Xml                                                     `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable *string                                                         `xml:"enable,omitempty"`
	Source *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc   []generic.Xml                                                   `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6 *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                         `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Misc   []generic.Xml                                                                        `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable *string                                                         `xml:"enable,omitempty"`
	Source *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc   []generic.Xml                                                   `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6 *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                         `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                                        `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationXml_11_0_2 struct {
	Enable *layer3Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 `xml:"enable,omitempty"`
	Misc   []generic.Xml                                         `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 struct {
	No   *layer3Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes  *layer3Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 `xml:"yes,omitempty"`
	Misc []generic.Xml                                            `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 struct {
	PfxPoolName   *string       `xml:"pfx-pool-name,omitempty"`
	PrefixLen     *int64        `xml:"prefix-len,omitempty"`
	PrefixLenHint *string       `xml:"prefix-len-hint,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsXml_11_0_2 struct {
	DuidType            *string                                        `xml:"duid-type,omitempty"`
	Enable              *layer3Ipv6DhcpClientV6OptionsEnableXml_11_0_2 `xml:"enable,omitempty"`
	RapidCommit         *string                                        `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                                        `xml:"support-srvr-reconfig,omitempty"`
	Misc                []generic.Xml                                  `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableXml_11_0_2 struct {
	No   *layer3Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes  *layer3Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 `xml:"yes,omitempty"`
	Misc []generic.Xml                                     `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 struct {
	NonTempAddr *string       `xml:"non-temp-addr,omitempty"`
	TempAddr    *string       `xml:"temp-addr,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedXml_11_0_2 struct {
	AssignAddr        *layer3Ipv6InheritedAssignAddrContainerXml_11_0_2 `xml:"assign-addr,omitempty"`
	Enable            *string                                           `xml:"enable,omitempty"`
	NeighborDiscovery *layer3Ipv6InheritedNeighborDiscoveryXml_11_0_2   `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                                     `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrContainerXml_11_0_2 struct {
	Entries []layer3Ipv6InheritedAssignAddrXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6InheritedAssignAddrXml_11_0_2 struct {
	XMLName xml.Name                                     `xml:"entry"`
	Name    string                                       `xml:"name,attr"`
	Type    *layer3Ipv6InheritedAssignAddrTypeXml_11_0_2 `xml:"type,omitempty"`
	Misc    []generic.Xml                                `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeXml_11_0_2 struct {
	Gua  *layer3Ipv6InheritedAssignAddrTypeGuaXml_11_0_2 `xml:"gua,omitempty"`
	Ula  *layer3Ipv6InheritedAssignAddrTypeUlaXml_11_0_2 `xml:"ula,omitempty"`
	Misc []generic.Xml                                   `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaXml_11_0_2 struct {
	EnableOnInterface *string                                                  `xml:"enable-on-interface,omitempty"`
	PrefixPool        *string                                                  `xml:"prefix-pool,omitempty"`
	PoolType          *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2  `xml:"pool-type,omitempty"`
	Advertise         *layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                            `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2 struct {
	Dynamic   *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2   `xml:"dynamic,omitempty"`
	DynamicId *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 `xml:"dynamic-id,omitempty"`
	Misc      []generic.Xml                                                    `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 struct {
	Identifier *int64        `xml:"identifier,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 struct {
	Enable         *string       `xml:"enable,omitempty"`
	OnlinkFlag     *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag *string       `xml:"auto-config-flag,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeUlaXml_11_0_2 struct {
	EnableOnInterface *string                                                  `xml:"enable-on-interface,omitempty"`
	Address           *string                                                  `xml:"address,omitempty"`
	Prefix            *string                                                  `xml:"prefix,omitempty"`
	Anycast           *string                                                  `xml:"anycast,omitempty"`
	Advertise         *layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                            `xml:",any"`
}
type layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                                             `xml:"dad-attempts,omitempty"`
	DnsServer           *layer3Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2           `xml:"dns-server,omitempty"`
	DnsSuffix           *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                                            `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                            `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                                             `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                             `xml:"reachable-time,omitempty"`
	RouterAdvertisement *layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                                      `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable *string                                                        `xml:"enable,omitempty"`
	Source *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc   []generic.Xml                                                  `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6 *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                        `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool *string       `xml:"prefix-pool,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Misc   []generic.Xml                                                                       `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable *string                                                        `xml:"enable,omitempty"`
	Source *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc   []generic.Xml                                                  `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6 *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc   []generic.Xml                                                        `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool *string       `xml:"prefix-pool,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                                       `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
	Enable                 *string       `xml:"enable,omitempty"`
	EnableConsistencyCheck *string       `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string       `xml:"hop-limit,omitempty"`
	Lifetime               *int64        `xml:"lifetime,omitempty"`
	LinkMtu                *string       `xml:"link-mtu,omitempty"`
	ManagedFlag            *string       `xml:"managed-flag,omitempty"`
	MaxInterval            *int64        `xml:"max-interval,omitempty"`
	MinInterval            *int64        `xml:"min-interval,omitempty"`
	OtherFlag              *string       `xml:"other-flag,omitempty"`
	ReachableTime          *string       `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string       `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string       `xml:"router-preference,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                                    `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                                   `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                   `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *layer3Ipv6NeighborDiscoveryNeighborContainerXml_11_0_2   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                                    `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                    `xml:"reachable-time,omitempty"`
	RouterAdvertisement *layer3Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                             `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName   xml.Name      `xml:"entry"`
	Name      string        `xml:"name,attr"`
	HwAddress *string       `xml:"hw-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
	DnsSupport             *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 `xml:"dns-support,omitempty"`
	Enable                 *string                                                             `xml:"enable,omitempty"`
	EnableConsistencyCheck *string                                                             `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string                                                             `xml:"hop-limit,omitempty"`
	Lifetime               *int64                                                              `xml:"lifetime,omitempty"`
	LinkMtu                *string                                                             `xml:"link-mtu,omitempty"`
	ManagedFlag            *string                                                             `xml:"managed-flag,omitempty"`
	MaxInterval            *int64                                                              `xml:"max-interval,omitempty"`
	MinInterval            *int64                                                              `xml:"min-interval,omitempty"`
	OtherFlag              *string                                                             `xml:"other-flag,omitempty"`
	ReachableTime          *string                                                             `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string                                                             `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string                                                             `xml:"router-preference,omitempty"`
	Misc                   []generic.Xml                                                       `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 struct {
	Enable *string                                                                            `xml:"enable,omitempty"`
	Server *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Suffix *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc   []generic.Xml                                                                      `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2 struct {
	Entries []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2 struct {
	Entries []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 `xml:"entry"`
}
type layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 struct {
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	Lifetime *int64        `xml:"lifetime,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type layer3LacpXml_11_0_2 struct {
	Enable           *string                               `xml:"enable,omitempty"`
	FastFailover     *string                               `xml:"fast-failover,omitempty"`
	HighAvailability *layer3LacpHighAvailabilityXml_11_0_2 `xml:"high-availability,omitempty"`
	MaxPorts         *int64                                `xml:"max-ports,omitempty"`
	Mode             *string                               `xml:"mode,omitempty"`
	SystemPriority   *int64                                `xml:"system-priority,omitempty"`
	TransmissionRate *string                               `xml:"transmission-rate,omitempty"`
	Misc             []generic.Xml                         `xml:",any"`
}
type layer3LacpHighAvailabilityXml_11_0_2 struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3LldpXml_11_0_2 struct {
	Enable           *string                               `xml:"enable,omitempty"`
	HighAvailability *layer3LldpHighAvailabilityXml_11_0_2 `xml:"high-availability,omitempty"`
	Profile          *string                               `xml:"profile,omitempty"`
	Misc             []generic.Xml                         `xml:",any"`
}
type layer3LldpHighAvailabilityXml_11_0_2 struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type layer3NdpProxyXml_11_0_2 struct {
	Address *layer3NdpProxyAddressContainerXml_11_0_2 `xml:"address,omitempty"`
	Enabled *string                                   `xml:"enabled,omitempty"`
	Misc    []generic.Xml                             `xml:",any"`
}
type layer3NdpProxyAddressContainerXml_11_0_2 struct {
	Entries []layer3NdpProxyAddressXml_11_0_2 `xml:"entry"`
}
type layer3NdpProxyAddressXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Negate  *string       `xml:"negate,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type layer3SdwanLinkSettingsXml_11_0_2 struct {
	Enable                *string                                       `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                                       `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *layer3SdwanLinkSettingsUpstreamNatXml_11_0_2 `xml:"upstream-nat,omitempty"`
	Misc                  []generic.Xml                                 `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatXml_11_0_2 struct {
	Enable   *string                                               `xml:"enable,omitempty"`
	Ddns     *layer3SdwanLinkSettingsUpstreamNatDdnsXml_11_0_2     `xml:"ddns,omitempty"`
	StaticIp *layer3SdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2 `xml:"static-ip,omitempty"`
	Misc     []generic.Xml                                         `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatDdnsXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type layer3SdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2 struct {
	Fqdn      *string       `xml:"fqdn,omitempty"`
	IpAddress *string       `xml:"ip-address,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type virtualWireXml_11_0_2 struct {
	Lldp           *virtualWireLldpXml_11_0_2 `xml:"lldp,omitempty"`
	NetflowProfile *string                    `xml:"netflow-profile,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
}
type virtualWireLldpXml_11_0_2 struct {
	Enable           *string                                    `xml:"enable,omitempty"`
	HighAvailability *virtualWireLldpHighAvailabilityXml_11_0_2 `xml:"high-availability,omitempty"`
	Profile          *string                                    `xml:"profile,omitempty"`
	Misc             []generic.Xml                              `xml:",any"`
}
type virtualWireLldpHighAvailabilityXml_11_0_2 struct {
	PassivePreNegotiation *string       `xml:"passive-pre-negotiation,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Comment = s.Comment
	if s.DecryptMirror != nil {
		var obj decryptMirrorXml
		obj.MarshalFromObject(*s.DecryptMirror)
		o.DecryptMirror = &obj
	}
	if s.Ha != nil {
		var obj haXml
		obj.MarshalFromObject(*s.Ha)
		o.Ha = &obj
	}
	if s.Layer2 != nil {
		var obj layer2Xml
		obj.MarshalFromObject(*s.Layer2)
		o.Layer2 = &obj
	}
	if s.Layer3 != nil {
		var obj layer3Xml
		obj.MarshalFromObject(*s.Layer3)
		o.Layer3 = &obj
	}
	if s.VirtualWire != nil {
		var obj virtualWireXml
		obj.MarshalFromObject(*s.VirtualWire)
		o.VirtualWire = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var decryptMirrorVal *DecryptMirror
	if o.DecryptMirror != nil {
		obj, err := o.DecryptMirror.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptMirrorVal = obj
	}
	var haVal *Ha
	if o.Ha != nil {
		obj, err := o.Ha.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		haVal = obj
	}
	var layer2Val *Layer2
	if o.Layer2 != nil {
		obj, err := o.Layer2.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		layer2Val = obj
	}
	var layer3Val *Layer3
	if o.Layer3 != nil {
		obj, err := o.Layer3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		layer3Val = obj
	}
	var virtualWireVal *VirtualWire
	if o.VirtualWire != nil {
		obj, err := o.VirtualWire.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		virtualWireVal = obj
	}

	result := &Entry{
		Name:          o.Name,
		Comment:       o.Comment,
		DecryptMirror: decryptMirrorVal,
		Ha:            haVal,
		Layer2:        layer2Val,
		Layer3:        layer3Val,
		VirtualWire:   virtualWireVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *decryptMirrorXml) MarshalFromObject(s DecryptMirror) {
	o.Misc = s.Misc
}

func (o decryptMirrorXml) UnmarshalToObject() (*DecryptMirror, error) {

	result := &DecryptMirror{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *haXml) MarshalFromObject(s Ha) {
	if s.Lacp != nil {
		var obj haLacpXml
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	o.Misc = s.Misc
}

func (o haXml) UnmarshalToObject() (*Ha, error) {
	var lacpVal *HaLacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}

	result := &Ha{
		Lacp: lacpVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *haLacpXml) MarshalFromObject(s HaLacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o haLacpXml) UnmarshalToObject() (*HaLacp, error) {

	result := &HaLacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2Xml) MarshalFromObject(s Layer2) {
	if s.Lacp != nil {
		var obj layer2LacpXml
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	if s.Lldp != nil {
		var obj layer2LldpXml
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	o.Misc = s.Misc
}

func (o layer2Xml) UnmarshalToObject() (*Layer2, error) {
	var lacpVal *Layer2Lacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}
	var lldpVal *Layer2Lldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}

	result := &Layer2{
		Lacp:           lacpVal,
		Lldp:           lldpVal,
		NetflowProfile: o.NetflowProfile,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *layer2LacpXml) MarshalFromObject(s Layer2Lacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	if s.HighAvailability != nil {
		var obj layer2LacpHighAvailabilityXml
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o layer2LacpXml) UnmarshalToObject() (*Layer2Lacp, error) {
	var highAvailabilityVal *Layer2LacpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer2Lacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		HighAvailability: highAvailabilityVal,
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2LacpHighAvailabilityXml) MarshalFromObject(s Layer2LacpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer2LacpHighAvailabilityXml) UnmarshalToObject() (*Layer2LacpHighAvailability, error) {

	result := &Layer2LacpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer2LldpXml) MarshalFromObject(s Layer2Lldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj layer2LldpHighAvailabilityXml
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o layer2LldpXml) UnmarshalToObject() (*Layer2Lldp, error) {
	var highAvailabilityVal *Layer2LldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer2Lldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2LldpHighAvailabilityXml) MarshalFromObject(s Layer2LldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer2LldpHighAvailabilityXml) UnmarshalToObject() (*Layer2LldpHighAvailability, error) {

	result := &Layer2LldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3Xml) MarshalFromObject(s Layer3) {
	if s.AdjustTcpMss != nil {
		var obj layer3AdjustTcpMssXml
		obj.MarshalFromObject(*s.AdjustTcpMss)
		o.AdjustTcpMss = &obj
	}
	if s.Arp != nil {
		var objs []layer3ArpXml
		for _, elt := range s.Arp {
			var obj layer3ArpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Arp = &layer3ArpContainerXml{Entries: objs}
	}
	if s.Bonjour != nil {
		var obj layer3BonjourXml
		obj.MarshalFromObject(*s.Bonjour)
		o.Bonjour = &obj
	}
	if s.DdnsConfig != nil {
		var obj layer3DdnsConfigXml
		obj.MarshalFromObject(*s.DdnsConfig)
		o.DdnsConfig = &obj
	}
	o.DecryptForward = util.YesNo(s.DecryptForward, nil)
	o.DfIgnore = util.YesNo(s.DfIgnore, nil)
	if s.DhcpClient != nil {
		var obj layer3DhcpClientXml
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.InterfaceManagementProfile = s.InterfaceManagementProfile
	if s.Ip != nil {
		var objs []layer3IpXml
		for _, elt := range s.Ip {
			var obj layer3IpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ip = &layer3IpContainerXml{Entries: objs}
	}
	if s.Ipv6 != nil {
		var obj layer3Ipv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	if s.Lacp != nil {
		var obj layer3LacpXml
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	if s.Lldp != nil {
		var obj layer3LldpXml
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.Mtu = s.Mtu
	if s.NdpProxy != nil {
		var obj layer3NdpProxyXml
		obj.MarshalFromObject(*s.NdpProxy)
		o.NdpProxy = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	if s.SdwanLinkSettings != nil {
		var obj layer3SdwanLinkSettingsXml
		obj.MarshalFromObject(*s.SdwanLinkSettings)
		o.SdwanLinkSettings = &obj
	}
	o.UntaggedSubInterface = util.YesNo(s.UntaggedSubInterface, nil)
	o.Misc = s.Misc
}

func (o layer3Xml) UnmarshalToObject() (*Layer3, error) {
	var adjustTcpMssVal *Layer3AdjustTcpMss
	if o.AdjustTcpMss != nil {
		obj, err := o.AdjustTcpMss.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adjustTcpMssVal = obj
	}
	var arpVal []Layer3Arp
	if o.Arp != nil {
		for _, elt := range o.Arp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			arpVal = append(arpVal, *obj)
		}
	}
	var bonjourVal *Layer3Bonjour
	if o.Bonjour != nil {
		obj, err := o.Bonjour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bonjourVal = obj
	}
	var ddnsConfigVal *Layer3DdnsConfig
	if o.DdnsConfig != nil {
		obj, err := o.DdnsConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsConfigVal = obj
	}
	var dhcpClientVal *Layer3DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var ipVal []Layer3Ip
	if o.Ip != nil {
		for _, elt := range o.Ip.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipVal = append(ipVal, *obj)
		}
	}
	var ipv6Val *Layer3Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var lacpVal *Layer3Lacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}
	var lldpVal *Layer3Lldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}
	var ndpProxyVal *Layer3NdpProxy
	if o.NdpProxy != nil {
		obj, err := o.NdpProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ndpProxyVal = obj
	}
	var sdwanLinkSettingsVal *Layer3SdwanLinkSettings
	if o.SdwanLinkSettings != nil {
		obj, err := o.SdwanLinkSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanLinkSettingsVal = obj
	}

	result := &Layer3{
		AdjustTcpMss:               adjustTcpMssVal,
		Arp:                        arpVal,
		Bonjour:                    bonjourVal,
		DdnsConfig:                 ddnsConfigVal,
		DecryptForward:             util.AsBool(o.DecryptForward, nil),
		DfIgnore:                   util.AsBool(o.DfIgnore, nil),
		DhcpClient:                 dhcpClientVal,
		InterfaceManagementProfile: o.InterfaceManagementProfile,
		Ip:                         ipVal,
		Ipv6:                       ipv6Val,
		Lacp:                       lacpVal,
		Lldp:                       lldpVal,
		Mtu:                        o.Mtu,
		NdpProxy:                   ndpProxyVal,
		NetflowProfile:             o.NetflowProfile,
		SdwanLinkSettings:          sdwanLinkSettingsVal,
		UntaggedSubInterface:       util.AsBool(o.UntaggedSubInterface, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *layer3AdjustTcpMssXml) MarshalFromObject(s Layer3AdjustTcpMss) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Ipv4MssAdjustment = s.Ipv4MssAdjustment
	o.Ipv6MssAdjustment = s.Ipv6MssAdjustment
	o.Misc = s.Misc
}

func (o layer3AdjustTcpMssXml) UnmarshalToObject() (*Layer3AdjustTcpMss, error) {

	result := &Layer3AdjustTcpMss{
		Enable:            util.AsBool(o.Enable, nil),
		Ipv4MssAdjustment: o.Ipv4MssAdjustment,
		Ipv6MssAdjustment: o.Ipv6MssAdjustment,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3ArpXml) MarshalFromObject(s Layer3Arp) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3ArpXml) UnmarshalToObject() (*Layer3Arp, error) {

	result := &Layer3Arp{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3BonjourXml) MarshalFromObject(s Layer3Bonjour) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GroupId = s.GroupId
	o.TtlCheck = util.YesNo(s.TtlCheck, nil)
	o.Misc = s.Misc
}

func (o layer3BonjourXml) UnmarshalToObject() (*Layer3Bonjour, error) {

	result := &Layer3Bonjour{
		Enable:   util.AsBool(o.Enable, nil),
		GroupId:  o.GroupId,
		TtlCheck: util.AsBool(o.TtlCheck, nil),
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3DdnsConfigXml) MarshalFromObject(s Layer3DdnsConfig) {
	o.DdnsCertProfile = s.DdnsCertProfile
	o.DdnsEnabled = util.YesNo(s.DdnsEnabled, nil)
	o.DdnsHostname = s.DdnsHostname
	if s.DdnsIp != nil {
		o.DdnsIp = util.StrToMem(s.DdnsIp)
	}
	if s.DdnsIpv6 != nil {
		o.DdnsIpv6 = util.StrToMem(s.DdnsIpv6)
	}
	o.DdnsUpdateInterval = s.DdnsUpdateInterval
	o.DdnsVendor = s.DdnsVendor
	if s.DdnsVendorConfig != nil {
		var objs []layer3DdnsConfigDdnsVendorConfigXml
		for _, elt := range s.DdnsVendorConfig {
			var obj layer3DdnsConfigDdnsVendorConfigXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DdnsVendorConfig = &layer3DdnsConfigDdnsVendorConfigContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3DdnsConfigXml) UnmarshalToObject() (*Layer3DdnsConfig, error) {
	var ddnsIpVal []string
	if o.DdnsIp != nil {
		ddnsIpVal = util.MemToStr(o.DdnsIp)
	}
	var ddnsIpv6Val []string
	if o.DdnsIpv6 != nil {
		ddnsIpv6Val = util.MemToStr(o.DdnsIpv6)
	}
	var ddnsVendorConfigVal []Layer3DdnsConfigDdnsVendorConfig
	if o.DdnsVendorConfig != nil {
		for _, elt := range o.DdnsVendorConfig.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ddnsVendorConfigVal = append(ddnsVendorConfigVal, *obj)
		}
	}

	result := &Layer3DdnsConfig{
		DdnsCertProfile:    o.DdnsCertProfile,
		DdnsEnabled:        util.AsBool(o.DdnsEnabled, nil),
		DdnsHostname:       o.DdnsHostname,
		DdnsIp:             ddnsIpVal,
		DdnsIpv6:           ddnsIpv6Val,
		DdnsUpdateInterval: o.DdnsUpdateInterval,
		DdnsVendor:         o.DdnsVendor,
		DdnsVendorConfig:   ddnsVendorConfigVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3DdnsConfigDdnsVendorConfigXml) MarshalFromObject(s Layer3DdnsConfigDdnsVendorConfig) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
}

func (o layer3DdnsConfigDdnsVendorConfigXml) UnmarshalToObject() (*Layer3DdnsConfigDdnsVendorConfig, error) {

	result := &Layer3DdnsConfigDdnsVendorConfig{
		Name:  o.Name,
		Value: o.Value,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *layer3DhcpClientXml) MarshalFromObject(s Layer3DhcpClient) {
	o.CreateDefaultRoute = util.YesNo(s.CreateDefaultRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.SendHostname != nil {
		var obj layer3DhcpClientSendHostnameXml
		obj.MarshalFromObject(*s.SendHostname)
		o.SendHostname = &obj
	}
	o.Misc = s.Misc
}

func (o layer3DhcpClientXml) UnmarshalToObject() (*Layer3DhcpClient, error) {
	var sendHostnameVal *Layer3DhcpClientSendHostname
	if o.SendHostname != nil {
		obj, err := o.SendHostname.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendHostnameVal = obj
	}

	result := &Layer3DhcpClient{
		CreateDefaultRoute: util.AsBool(o.CreateDefaultRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		SendHostname:       sendHostnameVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3DhcpClientSendHostnameXml) MarshalFromObject(s Layer3DhcpClientSendHostname) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Hostname = s.Hostname
	o.Misc = s.Misc
}

func (o layer3DhcpClientSendHostnameXml) UnmarshalToObject() (*Layer3DhcpClientSendHostname, error) {

	result := &Layer3DhcpClientSendHostname{
		Enable:   util.AsBool(o.Enable, nil),
		Hostname: o.Hostname,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3IpXml) MarshalFromObject(s Layer3Ip) {
	o.Name = s.Name
	o.SdwanGateway = s.SdwanGateway
	o.Misc = s.Misc
}

func (o layer3IpXml) UnmarshalToObject() (*Layer3Ip, error) {

	result := &Layer3Ip{
		Name:         o.Name,
		SdwanGateway: o.SdwanGateway,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6Xml) MarshalFromObject(s Layer3Ipv6) {
	if s.Address != nil {
		var objs []layer3Ipv6AddressXml
		for _, elt := range s.Address {
			var obj layer3Ipv6AddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &layer3Ipv6AddressContainerXml{Entries: objs}
	}
	if s.DhcpClient != nil {
		var obj layer3Ipv6DhcpClientXml
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Inherited != nil {
		var obj layer3Ipv6InheritedXml
		obj.MarshalFromObject(*s.Inherited)
		o.Inherited = &obj
	}
	o.InterfaceId = s.InterfaceId
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6NeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6Xml) UnmarshalToObject() (*Layer3Ipv6, error) {
	var addressVal []Layer3Ipv6Address
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}
	var dhcpClientVal *Layer3Ipv6DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var inheritedVal *Layer3Ipv6Inherited
	if o.Inherited != nil {
		obj, err := o.Inherited.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritedVal = obj
	}
	var neighborDiscoveryVal *Layer3Ipv6NeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Layer3Ipv6{
		Address:           addressVal,
		DhcpClient:        dhcpClientVal,
		Enabled:           util.AsBool(o.Enabled, nil),
		Inherited:         inheritedVal,
		InterfaceId:       o.InterfaceId,
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressXml) MarshalFromObject(s Layer3Ipv6Address) {
	o.Name = s.Name
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	if s.Prefix != nil {
		var obj layer3Ipv6AddressPrefixXml
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	if s.Anycast != nil {
		var obj layer3Ipv6AddressAnycastXml
		obj.MarshalFromObject(*s.Anycast)
		o.Anycast = &obj
	}
	if s.Advertise != nil {
		var obj layer3Ipv6AddressAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressXml) UnmarshalToObject() (*Layer3Ipv6Address, error) {
	var prefixVal *Layer3Ipv6AddressPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}
	var anycastVal *Layer3Ipv6AddressAnycast
	if o.Anycast != nil {
		obj, err := o.Anycast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anycastVal = obj
	}
	var advertiseVal *Layer3Ipv6AddressAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6Address{
		Name:              o.Name,
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Prefix:            prefixVal,
		Anycast:           anycastVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressPrefixXml) MarshalFromObject(s Layer3Ipv6AddressPrefix) {
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressPrefixXml) UnmarshalToObject() (*Layer3Ipv6AddressPrefix, error) {

	result := &Layer3Ipv6AddressPrefix{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressAnycastXml) MarshalFromObject(s Layer3Ipv6AddressAnycast) {
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressAnycastXml) UnmarshalToObject() (*Layer3Ipv6AddressAnycast, error) {

	result := &Layer3Ipv6AddressAnycast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressAdvertiseXml) MarshalFromObject(s Layer3Ipv6AddressAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressAdvertiseXml) UnmarshalToObject() (*Layer3Ipv6AddressAdvertise, error) {

	result := &Layer3Ipv6AddressAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientXml) MarshalFromObject(s Layer3Ipv6DhcpClient) {
	o.AcceptRaRoute = util.YesNo(s.AcceptRaRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Preference = s.Preference
	if s.PrefixDelegation != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationXml
		obj.MarshalFromObject(*s.PrefixDelegation)
		o.PrefixDelegation = &obj
	}
	if s.V6Options != nil {
		var obj layer3Ipv6DhcpClientV6OptionsXml
		obj.MarshalFromObject(*s.V6Options)
		o.V6Options = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientXml) UnmarshalToObject() (*Layer3Ipv6DhcpClient, error) {
	var neighborDiscoveryVal *Layer3Ipv6DhcpClientNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}
	var prefixDelegationVal *Layer3Ipv6DhcpClientPrefixDelegation
	if o.PrefixDelegation != nil {
		obj, err := o.PrefixDelegation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixDelegationVal = obj
	}
	var v6OptionsVal *Layer3Ipv6DhcpClientV6Options
	if o.V6Options != nil {
		obj, err := o.V6Options.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v6OptionsVal = obj
	}

	result := &Layer3Ipv6DhcpClient{
		AcceptRaRoute:      util.AsBool(o.AcceptRaRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		NeighborDiscovery:  neighborDiscoveryVal,
		Preference:         o.Preference,
		PrefixDelegation:   prefixDelegationVal,
		V6Options:          v6OptionsVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscovery, error) {
	var dnsServerVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscovery{
		DadAttempts:      o.DadAttempts,
		DnsServer:        dnsServerVal,
		DnsSuffix:        dnsSuffixVal,
		EnableDad:        util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor: util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:         neighborVal,
		NsInterval:       o.NsInterval,
		ReachableTime:    o.ReachableTime,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer, error) {
	var sourceVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml
		for _, elt := range s.Server {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{
		Server: serverVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml
		for _, elt := range s.Suffix {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationXml) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegation) {
	if s.Enable != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableXml
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegation, error) {
	var enableVal *Layer3Ipv6DhcpClientPrefixDelegationEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Layer3Ipv6DhcpClientPrefixDelegation{
		Enable: enableVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableXml) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnable) {
	if s.No != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnable, error) {
	var noVal *Layer3Ipv6DhcpClientPrefixDelegationEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Layer3Ipv6DhcpClientPrefixDelegationEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnable{
		No:   noVal,
		Yes:  yesVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableNoXml) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnableNo) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableNoXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnableNo, error) {

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnableNo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableYesXml) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnableYes) {
	o.PfxPoolName = s.PfxPoolName
	o.PrefixLen = s.PrefixLen
	o.PrefixLenHint = util.YesNo(s.PrefixLenHint, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableYesXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnableYes, error) {

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnableYes{
		PfxPoolName:   o.PfxPoolName,
		PrefixLen:     o.PrefixLen,
		PrefixLenHint: util.AsBool(o.PrefixLenHint, nil),
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsXml) MarshalFromObject(s Layer3Ipv6DhcpClientV6Options) {
	o.DuidType = s.DuidType
	if s.Enable != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableXml
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.RapidCommit = util.YesNo(s.RapidCommit, nil)
	o.SupportSrvrReconfig = util.YesNo(s.SupportSrvrReconfig, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6Options, error) {
	var enableVal *Layer3Ipv6DhcpClientV6OptionsEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Layer3Ipv6DhcpClientV6Options{
		DuidType:            o.DuidType,
		Enable:              enableVal,
		RapidCommit:         util.AsBool(o.RapidCommit, nil),
		SupportSrvrReconfig: util.AsBool(o.SupportSrvrReconfig, nil),
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableXml) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnable) {
	if s.No != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnable, error) {
	var noVal *Layer3Ipv6DhcpClientV6OptionsEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Layer3Ipv6DhcpClientV6OptionsEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Layer3Ipv6DhcpClientV6OptionsEnable{
		No:   noVal,
		Yes:  yesVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableNoXml) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnableNo) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableNoXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnableNo, error) {

	result := &Layer3Ipv6DhcpClientV6OptionsEnableNo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableYesXml) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnableYes) {
	o.NonTempAddr = util.YesNo(s.NonTempAddr, nil)
	o.TempAddr = util.YesNo(s.TempAddr, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableYesXml) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnableYes, error) {

	result := &Layer3Ipv6DhcpClientV6OptionsEnableYes{
		NonTempAddr: util.AsBool(o.NonTempAddr, nil),
		TempAddr:    util.AsBool(o.TempAddr, nil),
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedXml) MarshalFromObject(s Layer3Ipv6Inherited) {
	if s.AssignAddr != nil {
		var objs []layer3Ipv6InheritedAssignAddrXml
		for _, elt := range s.AssignAddr {
			var obj layer3Ipv6InheritedAssignAddrXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AssignAddr = &layer3Ipv6InheritedAssignAddrContainerXml{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedXml) UnmarshalToObject() (*Layer3Ipv6Inherited, error) {
	var assignAddrVal []Layer3Ipv6InheritedAssignAddr
	if o.AssignAddr != nil {
		for _, elt := range o.AssignAddr.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			assignAddrVal = append(assignAddrVal, *obj)
		}
	}
	var neighborDiscoveryVal *Layer3Ipv6InheritedNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Layer3Ipv6Inherited{
		AssignAddr:        assignAddrVal,
		Enable:            util.AsBool(o.Enable, nil),
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddr) {
	o.Name = s.Name
	if s.Type != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddr, error) {
	var typeVal *Layer3Ipv6InheritedAssignAddrType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddr{
		Name: o.Name,
		Type: typeVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrType) {
	if s.Gua != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaXml
		obj.MarshalFromObject(*s.Gua)
		o.Gua = &obj
	}
	if s.Ula != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeUlaXml
		obj.MarshalFromObject(*s.Ula)
		o.Ula = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrType, error) {
	var guaVal *Layer3Ipv6InheritedAssignAddrTypeGua
	if o.Gua != nil {
		obj, err := o.Gua.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		guaVal = obj
	}
	var ulaVal *Layer3Ipv6InheritedAssignAddrTypeUla
	if o.Ula != nil {
		obj, err := o.Ula.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ulaVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrType{
		Gua:  guaVal,
		Ula:  ulaVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGua) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.PrefixPool = s.PrefixPool
	if s.PoolType != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml
		obj.MarshalFromObject(*s.PoolType)
		o.PoolType = &obj
	}
	if s.Advertise != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGua, error) {
	var poolTypeVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType
	if o.PoolType != nil {
		obj, err := o.PoolType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		poolTypeVal = obj
	}
	var advertiseVal *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeGua{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		PrefixPool:        o.PrefixPool,
		PoolType:          poolTypeVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolType) {
	if s.Dynamic != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.DynamicId != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml
		obj.MarshalFromObject(*s.DynamicId)
		o.DynamicId = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolType, error) {
	var dynamicVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var dynamicIdVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	if o.DynamicId != nil {
		obj, err := o.DynamicId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIdVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolType{
		Dynamic:   dynamicVal,
		DynamicId: dynamicIdVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) {
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) {
	o.Identifier = s.Identifier
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{
		Identifier: o.Identifier,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise{
		Enable:         util.AsBool(o.Enable, nil),
		OnlinkFlag:     util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag: util.AsBool(o.AutoConfigFlag, nil),
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeUlaXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeUla) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.Address = s.Address
	o.Prefix = util.YesNo(s.Prefix, nil)
	o.Anycast = util.YesNo(s.Anycast, nil)
	if s.Advertise != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeUlaXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeUla, error) {
	var advertiseVal *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeUla{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Address:           o.Address,
		Prefix:            util.AsBool(o.Prefix, nil),
		Anycast:           util.AsBool(o.Anycast, nil),
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerXml
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6InheritedNeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscovery, error) {
	var dnsServerVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Layer3Ipv6InheritedNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		DnsServer:           dnsServerVal,
		DnsSuffix:           dnsSuffixVal,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServer, error) {
	var sourceVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServer{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{
		PrefixPool: o.PrefixPool,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml
		for _, elt := range s.Server {
			var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{
		Server: serverVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{
		PrefixPool: o.PrefixPool,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml
		for _, elt := range s.Suffix {
			var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryNeighborXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryNeighborXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.EnableConsistencyCheck = util.YesNo(s.EnableConsistencyCheck, nil)
	o.HopLimit = s.HopLimit
	o.Lifetime = s.Lifetime
	o.LinkMtu = s.LinkMtu
	o.ManagedFlag = util.YesNo(s.ManagedFlag, nil)
	o.MaxInterval = s.MaxInterval
	o.MinInterval = s.MinInterval
	o.OtherFlag = util.YesNo(s.OtherFlag, nil)
	o.ReachableTime = s.ReachableTime
	o.RetransmissionTimer = s.RetransmissionTimer
	o.RouterPreference = s.RouterPreference
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement{
		Enable:                 util.AsBool(o.Enable, nil),
		EnableConsistencyCheck: util.AsBool(o.EnableConsistencyCheck, nil),
		HopLimit:               o.HopLimit,
		Lifetime:               o.Lifetime,
		LinkMtu:                o.LinkMtu,
		ManagedFlag:            util.AsBool(o.ManagedFlag, nil),
		MaxInterval:            o.MaxInterval,
		MinInterval:            o.MinInterval,
		OtherFlag:              util.AsBool(o.OtherFlag, nil),
		ReachableTime:          o.ReachableTime,
		RetransmissionTimer:    o.RetransmissionTimer,
		RouterPreference:       o.RouterPreference,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryXml) MarshalFromObject(s Layer3Ipv6NeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6NeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6NeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6NeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementXml
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscovery, error) {
	var neighborVal []Layer3Ipv6NeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Layer3Ipv6NeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Layer3Ipv6NeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryNeighborXml) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryNeighborXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6NeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementXml) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisement) {
	if s.DnsSupport != nil {
		var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml
		obj.MarshalFromObject(*s.DnsSupport)
		o.DnsSupport = &obj
	}
	o.Enable = util.YesNo(s.Enable, nil)
	o.EnableConsistencyCheck = util.YesNo(s.EnableConsistencyCheck, nil)
	o.HopLimit = s.HopLimit
	o.Lifetime = s.Lifetime
	o.LinkMtu = s.LinkMtu
	o.ManagedFlag = util.YesNo(s.ManagedFlag, nil)
	o.MaxInterval = s.MaxInterval
	o.MinInterval = s.MinInterval
	o.OtherFlag = util.YesNo(s.OtherFlag, nil)
	o.ReachableTime = s.ReachableTime
	o.RetransmissionTimer = s.RetransmissionTimer
	o.RouterPreference = s.RouterPreference
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisement, error) {
	var dnsSupportVal *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
	if o.DnsSupport != nil {
		obj, err := o.DnsSupport.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSupportVal = obj
	}

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisement{
		DnsSupport:             dnsSupportVal,
		Enable:                 util.AsBool(o.Enable, nil),
		EnableConsistencyCheck: util.AsBool(o.EnableConsistencyCheck, nil),
		HopLimit:               o.HopLimit,
		Lifetime:               o.Lifetime,
		LinkMtu:                o.LinkMtu,
		ManagedFlag:            util.AsBool(o.ManagedFlag, nil),
		MaxInterval:            o.MaxInterval,
		MinInterval:            o.MinInterval,
		OtherFlag:              util.AsBool(o.OtherFlag, nil),
		ReachableTime:          o.ReachableTime,
		RetransmissionTimer:    o.RetransmissionTimer,
		RouterPreference:       o.RouterPreference,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Server != nil {
		var objs []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml
		for _, elt := range s.Server {
			var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml{Entries: objs}
	}
	if s.Suffix != nil {
		var objs []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml
		for _, elt := range s.Suffix {
			var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, error) {
	var serverVal []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}
	var suffixVal []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{
		Enable: util.AsBool(o.Enable, nil),
		Server: serverVal,
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, error) {

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, error) {

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3LacpXml) MarshalFromObject(s Layer3Lacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	if s.HighAvailability != nil {
		var obj layer3LacpHighAvailabilityXml
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o layer3LacpXml) UnmarshalToObject() (*Layer3Lacp, error) {
	var highAvailabilityVal *Layer3LacpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer3Lacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		HighAvailability: highAvailabilityVal,
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3LacpHighAvailabilityXml) MarshalFromObject(s Layer3LacpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer3LacpHighAvailabilityXml) UnmarshalToObject() (*Layer3LacpHighAvailability, error) {

	result := &Layer3LacpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3LldpXml) MarshalFromObject(s Layer3Lldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj layer3LldpHighAvailabilityXml
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o layer3LldpXml) UnmarshalToObject() (*Layer3Lldp, error) {
	var highAvailabilityVal *Layer3LldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer3Lldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3LldpHighAvailabilityXml) MarshalFromObject(s Layer3LldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer3LldpHighAvailabilityXml) UnmarshalToObject() (*Layer3LldpHighAvailability, error) {

	result := &Layer3LldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3NdpProxyXml) MarshalFromObject(s Layer3NdpProxy) {
	if s.Address != nil {
		var objs []layer3NdpProxyAddressXml
		for _, elt := range s.Address {
			var obj layer3NdpProxyAddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &layer3NdpProxyAddressContainerXml{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.Misc = s.Misc
}

func (o layer3NdpProxyXml) UnmarshalToObject() (*Layer3NdpProxy, error) {
	var addressVal []Layer3NdpProxyAddress
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &Layer3NdpProxy{
		Address: addressVal,
		Enabled: util.AsBool(o.Enabled, nil),
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *layer3NdpProxyAddressXml) MarshalFromObject(s Layer3NdpProxyAddress) {
	o.Name = s.Name
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
}

func (o layer3NdpProxyAddressXml) UnmarshalToObject() (*Layer3NdpProxyAddress, error) {

	result := &Layer3NdpProxyAddress{
		Name:   o.Name,
		Negate: util.AsBool(o.Negate, nil),
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsXml) MarshalFromObject(s Layer3SdwanLinkSettings) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	if s.UpstreamNat != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatXml
		obj.MarshalFromObject(*s.UpstreamNat)
		o.UpstreamNat = &obj
	}
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsXml) UnmarshalToObject() (*Layer3SdwanLinkSettings, error) {
	var upstreamNatVal *Layer3SdwanLinkSettingsUpstreamNat
	if o.UpstreamNat != nil {
		obj, err := o.UpstreamNat.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		upstreamNatVal = obj
	}

	result := &Layer3SdwanLinkSettings{
		Enable:                util.AsBool(o.Enable, nil),
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		UpstreamNat:           upstreamNatVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatXml) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNat) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Ddns != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatDdnsXml
		obj.MarshalFromObject(*s.Ddns)
		o.Ddns = &obj
	}
	if s.StaticIp != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatStaticIpXml
		obj.MarshalFromObject(*s.StaticIp)
		o.StaticIp = &obj
	}
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatXml) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNat, error) {
	var ddnsVal *Layer3SdwanLinkSettingsUpstreamNatDdns
	if o.Ddns != nil {
		obj, err := o.Ddns.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsVal = obj
	}
	var staticIpVal *Layer3SdwanLinkSettingsUpstreamNatStaticIp
	if o.StaticIp != nil {
		obj, err := o.StaticIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticIpVal = obj
	}

	result := &Layer3SdwanLinkSettingsUpstreamNat{
		Enable:   util.AsBool(o.Enable, nil),
		Ddns:     ddnsVal,
		StaticIp: staticIpVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatDdnsXml) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNatDdns) {
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatDdnsXml) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNatDdns, error) {

	result := &Layer3SdwanLinkSettingsUpstreamNatDdns{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatStaticIpXml) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNatStaticIp) {
	o.Fqdn = s.Fqdn
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatStaticIpXml) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNatStaticIp, error) {

	result := &Layer3SdwanLinkSettingsUpstreamNatStaticIp{
		Fqdn:      o.Fqdn,
		IpAddress: o.IpAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *virtualWireXml) MarshalFromObject(s VirtualWire) {
	if s.Lldp != nil {
		var obj virtualWireLldpXml
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	o.Misc = s.Misc
}

func (o virtualWireXml) UnmarshalToObject() (*VirtualWire, error) {
	var lldpVal *VirtualWireLldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}

	result := &VirtualWire{
		Lldp:           lldpVal,
		NetflowProfile: o.NetflowProfile,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *virtualWireLldpXml) MarshalFromObject(s VirtualWireLldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj virtualWireLldpHighAvailabilityXml
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o virtualWireLldpXml) UnmarshalToObject() (*VirtualWireLldp, error) {
	var highAvailabilityVal *VirtualWireLldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &VirtualWireLldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *virtualWireLldpHighAvailabilityXml) MarshalFromObject(s VirtualWireLldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o virtualWireLldpHighAvailabilityXml) UnmarshalToObject() (*VirtualWireLldpHighAvailability, error) {

	result := &VirtualWireLldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Comment = s.Comment
	if s.DecryptMirror != nil {
		var obj decryptMirrorXml_11_0_2
		obj.MarshalFromObject(*s.DecryptMirror)
		o.DecryptMirror = &obj
	}
	if s.Ha != nil {
		var obj haXml_11_0_2
		obj.MarshalFromObject(*s.Ha)
		o.Ha = &obj
	}
	if s.Layer2 != nil {
		var obj layer2Xml_11_0_2
		obj.MarshalFromObject(*s.Layer2)
		o.Layer2 = &obj
	}
	if s.Layer3 != nil {
		var obj layer3Xml_11_0_2
		obj.MarshalFromObject(*s.Layer3)
		o.Layer3 = &obj
	}
	if s.VirtualWire != nil {
		var obj virtualWireXml_11_0_2
		obj.MarshalFromObject(*s.VirtualWire)
		o.VirtualWire = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var decryptMirrorVal *DecryptMirror
	if o.DecryptMirror != nil {
		obj, err := o.DecryptMirror.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptMirrorVal = obj
	}
	var haVal *Ha
	if o.Ha != nil {
		obj, err := o.Ha.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		haVal = obj
	}
	var layer2Val *Layer2
	if o.Layer2 != nil {
		obj, err := o.Layer2.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		layer2Val = obj
	}
	var layer3Val *Layer3
	if o.Layer3 != nil {
		obj, err := o.Layer3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		layer3Val = obj
	}
	var virtualWireVal *VirtualWire
	if o.VirtualWire != nil {
		obj, err := o.VirtualWire.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		virtualWireVal = obj
	}

	result := &Entry{
		Name:          o.Name,
		Comment:       o.Comment,
		DecryptMirror: decryptMirrorVal,
		Ha:            haVal,
		Layer2:        layer2Val,
		Layer3:        layer3Val,
		VirtualWire:   virtualWireVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *decryptMirrorXml_11_0_2) MarshalFromObject(s DecryptMirror) {
	o.Misc = s.Misc
}

func (o decryptMirrorXml_11_0_2) UnmarshalToObject() (*DecryptMirror, error) {

	result := &DecryptMirror{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *haXml_11_0_2) MarshalFromObject(s Ha) {
	if s.Lacp != nil {
		var obj haLacpXml_11_0_2
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	o.Misc = s.Misc
}

func (o haXml_11_0_2) UnmarshalToObject() (*Ha, error) {
	var lacpVal *HaLacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}

	result := &Ha{
		Lacp: lacpVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *haLacpXml_11_0_2) MarshalFromObject(s HaLacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o haLacpXml_11_0_2) UnmarshalToObject() (*HaLacp, error) {

	result := &HaLacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2Xml_11_0_2) MarshalFromObject(s Layer2) {
	if s.Lacp != nil {
		var obj layer2LacpXml_11_0_2
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	if s.Lldp != nil {
		var obj layer2LldpXml_11_0_2
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	o.Misc = s.Misc
}

func (o layer2Xml_11_0_2) UnmarshalToObject() (*Layer2, error) {
	var lacpVal *Layer2Lacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}
	var lldpVal *Layer2Lldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}

	result := &Layer2{
		Lacp:           lacpVal,
		Lldp:           lldpVal,
		NetflowProfile: o.NetflowProfile,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *layer2LacpXml_11_0_2) MarshalFromObject(s Layer2Lacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	if s.HighAvailability != nil {
		var obj layer2LacpHighAvailabilityXml_11_0_2
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o layer2LacpXml_11_0_2) UnmarshalToObject() (*Layer2Lacp, error) {
	var highAvailabilityVal *Layer2LacpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer2Lacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		HighAvailability: highAvailabilityVal,
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2LacpHighAvailabilityXml_11_0_2) MarshalFromObject(s Layer2LacpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer2LacpHighAvailabilityXml_11_0_2) UnmarshalToObject() (*Layer2LacpHighAvailability, error) {

	result := &Layer2LacpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer2LldpXml_11_0_2) MarshalFromObject(s Layer2Lldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj layer2LldpHighAvailabilityXml_11_0_2
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o layer2LldpXml_11_0_2) UnmarshalToObject() (*Layer2Lldp, error) {
	var highAvailabilityVal *Layer2LldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer2Lldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer2LldpHighAvailabilityXml_11_0_2) MarshalFromObject(s Layer2LldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer2LldpHighAvailabilityXml_11_0_2) UnmarshalToObject() (*Layer2LldpHighAvailability, error) {

	result := &Layer2LldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3Xml_11_0_2) MarshalFromObject(s Layer3) {
	if s.AdjustTcpMss != nil {
		var obj layer3AdjustTcpMssXml_11_0_2
		obj.MarshalFromObject(*s.AdjustTcpMss)
		o.AdjustTcpMss = &obj
	}
	if s.Arp != nil {
		var objs []layer3ArpXml_11_0_2
		for _, elt := range s.Arp {
			var obj layer3ArpXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Arp = &layer3ArpContainerXml_11_0_2{Entries: objs}
	}
	if s.Bonjour != nil {
		var obj layer3BonjourXml_11_0_2
		obj.MarshalFromObject(*s.Bonjour)
		o.Bonjour = &obj
	}
	if s.DdnsConfig != nil {
		var obj layer3DdnsConfigXml_11_0_2
		obj.MarshalFromObject(*s.DdnsConfig)
		o.DdnsConfig = &obj
	}
	o.DecryptForward = util.YesNo(s.DecryptForward, nil)
	o.DfIgnore = util.YesNo(s.DfIgnore, nil)
	if s.DhcpClient != nil {
		var obj layer3DhcpClientXml_11_0_2
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.InterfaceManagementProfile = s.InterfaceManagementProfile
	if s.Ip != nil {
		var objs []layer3IpXml_11_0_2
		for _, elt := range s.Ip {
			var obj layer3IpXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ip = &layer3IpContainerXml_11_0_2{Entries: objs}
	}
	if s.Ipv6 != nil {
		var obj layer3Ipv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	if s.Lacp != nil {
		var obj layer3LacpXml_11_0_2
		obj.MarshalFromObject(*s.Lacp)
		o.Lacp = &obj
	}
	if s.Lldp != nil {
		var obj layer3LldpXml_11_0_2
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.Mtu = s.Mtu
	if s.NdpProxy != nil {
		var obj layer3NdpProxyXml_11_0_2
		obj.MarshalFromObject(*s.NdpProxy)
		o.NdpProxy = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	if s.SdwanLinkSettings != nil {
		var obj layer3SdwanLinkSettingsXml_11_0_2
		obj.MarshalFromObject(*s.SdwanLinkSettings)
		o.SdwanLinkSettings = &obj
	}
	o.UntaggedSubInterface = util.YesNo(s.UntaggedSubInterface, nil)
	o.Misc = s.Misc
}

func (o layer3Xml_11_0_2) UnmarshalToObject() (*Layer3, error) {
	var adjustTcpMssVal *Layer3AdjustTcpMss
	if o.AdjustTcpMss != nil {
		obj, err := o.AdjustTcpMss.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adjustTcpMssVal = obj
	}
	var arpVal []Layer3Arp
	if o.Arp != nil {
		for _, elt := range o.Arp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			arpVal = append(arpVal, *obj)
		}
	}
	var bonjourVal *Layer3Bonjour
	if o.Bonjour != nil {
		obj, err := o.Bonjour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bonjourVal = obj
	}
	var ddnsConfigVal *Layer3DdnsConfig
	if o.DdnsConfig != nil {
		obj, err := o.DdnsConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsConfigVal = obj
	}
	var dhcpClientVal *Layer3DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var ipVal []Layer3Ip
	if o.Ip != nil {
		for _, elt := range o.Ip.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipVal = append(ipVal, *obj)
		}
	}
	var ipv6Val *Layer3Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var lacpVal *Layer3Lacp
	if o.Lacp != nil {
		obj, err := o.Lacp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lacpVal = obj
	}
	var lldpVal *Layer3Lldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}
	var ndpProxyVal *Layer3NdpProxy
	if o.NdpProxy != nil {
		obj, err := o.NdpProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ndpProxyVal = obj
	}
	var sdwanLinkSettingsVal *Layer3SdwanLinkSettings
	if o.SdwanLinkSettings != nil {
		obj, err := o.SdwanLinkSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanLinkSettingsVal = obj
	}

	result := &Layer3{
		AdjustTcpMss:               adjustTcpMssVal,
		Arp:                        arpVal,
		Bonjour:                    bonjourVal,
		DdnsConfig:                 ddnsConfigVal,
		DecryptForward:             util.AsBool(o.DecryptForward, nil),
		DfIgnore:                   util.AsBool(o.DfIgnore, nil),
		DhcpClient:                 dhcpClientVal,
		InterfaceManagementProfile: o.InterfaceManagementProfile,
		Ip:                         ipVal,
		Ipv6:                       ipv6Val,
		Lacp:                       lacpVal,
		Lldp:                       lldpVal,
		Mtu:                        o.Mtu,
		NdpProxy:                   ndpProxyVal,
		NetflowProfile:             o.NetflowProfile,
		SdwanLinkSettings:          sdwanLinkSettingsVal,
		UntaggedSubInterface:       util.AsBool(o.UntaggedSubInterface, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *layer3AdjustTcpMssXml_11_0_2) MarshalFromObject(s Layer3AdjustTcpMss) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Ipv4MssAdjustment = s.Ipv4MssAdjustment
	o.Ipv6MssAdjustment = s.Ipv6MssAdjustment
	o.Misc = s.Misc
}

func (o layer3AdjustTcpMssXml_11_0_2) UnmarshalToObject() (*Layer3AdjustTcpMss, error) {

	result := &Layer3AdjustTcpMss{
		Enable:            util.AsBool(o.Enable, nil),
		Ipv4MssAdjustment: o.Ipv4MssAdjustment,
		Ipv6MssAdjustment: o.Ipv6MssAdjustment,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3ArpXml_11_0_2) MarshalFromObject(s Layer3Arp) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3ArpXml_11_0_2) UnmarshalToObject() (*Layer3Arp, error) {

	result := &Layer3Arp{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3BonjourXml_11_0_2) MarshalFromObject(s Layer3Bonjour) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GroupId = s.GroupId
	o.TtlCheck = util.YesNo(s.TtlCheck, nil)
	o.Misc = s.Misc
}

func (o layer3BonjourXml_11_0_2) UnmarshalToObject() (*Layer3Bonjour, error) {

	result := &Layer3Bonjour{
		Enable:   util.AsBool(o.Enable, nil),
		GroupId:  o.GroupId,
		TtlCheck: util.AsBool(o.TtlCheck, nil),
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3DdnsConfigXml_11_0_2) MarshalFromObject(s Layer3DdnsConfig) {
	o.DdnsCertProfile = s.DdnsCertProfile
	o.DdnsEnabled = util.YesNo(s.DdnsEnabled, nil)
	o.DdnsHostname = s.DdnsHostname
	if s.DdnsIp != nil {
		o.DdnsIp = util.StrToMem(s.DdnsIp)
	}
	if s.DdnsIpv6 != nil {
		o.DdnsIpv6 = util.StrToMem(s.DdnsIpv6)
	}
	o.DdnsUpdateInterval = s.DdnsUpdateInterval
	o.DdnsVendor = s.DdnsVendor
	if s.DdnsVendorConfig != nil {
		var objs []layer3DdnsConfigDdnsVendorConfigXml_11_0_2
		for _, elt := range s.DdnsVendorConfig {
			var obj layer3DdnsConfigDdnsVendorConfigXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DdnsVendorConfig = &layer3DdnsConfigDdnsVendorConfigContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3DdnsConfigXml_11_0_2) UnmarshalToObject() (*Layer3DdnsConfig, error) {
	var ddnsIpVal []string
	if o.DdnsIp != nil {
		ddnsIpVal = util.MemToStr(o.DdnsIp)
	}
	var ddnsIpv6Val []string
	if o.DdnsIpv6 != nil {
		ddnsIpv6Val = util.MemToStr(o.DdnsIpv6)
	}
	var ddnsVendorConfigVal []Layer3DdnsConfigDdnsVendorConfig
	if o.DdnsVendorConfig != nil {
		for _, elt := range o.DdnsVendorConfig.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ddnsVendorConfigVal = append(ddnsVendorConfigVal, *obj)
		}
	}

	result := &Layer3DdnsConfig{
		DdnsCertProfile:    o.DdnsCertProfile,
		DdnsEnabled:        util.AsBool(o.DdnsEnabled, nil),
		DdnsHostname:       o.DdnsHostname,
		DdnsIp:             ddnsIpVal,
		DdnsIpv6:           ddnsIpv6Val,
		DdnsUpdateInterval: o.DdnsUpdateInterval,
		DdnsVendor:         o.DdnsVendor,
		DdnsVendorConfig:   ddnsVendorConfigVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3DdnsConfigDdnsVendorConfigXml_11_0_2) MarshalFromObject(s Layer3DdnsConfigDdnsVendorConfig) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
}

func (o layer3DdnsConfigDdnsVendorConfigXml_11_0_2) UnmarshalToObject() (*Layer3DdnsConfigDdnsVendorConfig, error) {

	result := &Layer3DdnsConfigDdnsVendorConfig{
		Name:  o.Name,
		Value: o.Value,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *layer3DhcpClientXml_11_0_2) MarshalFromObject(s Layer3DhcpClient) {
	o.CreateDefaultRoute = util.YesNo(s.CreateDefaultRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.SendHostname != nil {
		var obj layer3DhcpClientSendHostnameXml_11_0_2
		obj.MarshalFromObject(*s.SendHostname)
		o.SendHostname = &obj
	}
	o.Misc = s.Misc
}

func (o layer3DhcpClientXml_11_0_2) UnmarshalToObject() (*Layer3DhcpClient, error) {
	var sendHostnameVal *Layer3DhcpClientSendHostname
	if o.SendHostname != nil {
		obj, err := o.SendHostname.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendHostnameVal = obj
	}

	result := &Layer3DhcpClient{
		CreateDefaultRoute: util.AsBool(o.CreateDefaultRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		SendHostname:       sendHostnameVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3DhcpClientSendHostnameXml_11_0_2) MarshalFromObject(s Layer3DhcpClientSendHostname) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Hostname = s.Hostname
	o.Misc = s.Misc
}

func (o layer3DhcpClientSendHostnameXml_11_0_2) UnmarshalToObject() (*Layer3DhcpClientSendHostname, error) {

	result := &Layer3DhcpClientSendHostname{
		Enable:   util.AsBool(o.Enable, nil),
		Hostname: o.Hostname,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3IpXml_11_0_2) MarshalFromObject(s Layer3Ip) {
	o.Name = s.Name
	o.SdwanGateway = s.SdwanGateway
	o.Misc = s.Misc
}

func (o layer3IpXml_11_0_2) UnmarshalToObject() (*Layer3Ip, error) {

	result := &Layer3Ip{
		Name:         o.Name,
		SdwanGateway: o.SdwanGateway,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6Xml_11_0_2) MarshalFromObject(s Layer3Ipv6) {
	if s.Address != nil {
		var objs []layer3Ipv6AddressXml_11_0_2
		for _, elt := range s.Address {
			var obj layer3Ipv6AddressXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &layer3Ipv6AddressContainerXml_11_0_2{Entries: objs}
	}
	if s.DhcpClient != nil {
		var obj layer3Ipv6DhcpClientXml_11_0_2
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Inherited != nil {
		var obj layer3Ipv6InheritedXml_11_0_2
		obj.MarshalFromObject(*s.Inherited)
		o.Inherited = &obj
	}
	o.InterfaceId = s.InterfaceId
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6NeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6Xml_11_0_2) UnmarshalToObject() (*Layer3Ipv6, error) {
	var addressVal []Layer3Ipv6Address
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}
	var dhcpClientVal *Layer3Ipv6DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var inheritedVal *Layer3Ipv6Inherited
	if o.Inherited != nil {
		obj, err := o.Inherited.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritedVal = obj
	}
	var neighborDiscoveryVal *Layer3Ipv6NeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Layer3Ipv6{
		Address:           addressVal,
		DhcpClient:        dhcpClientVal,
		Enabled:           util.AsBool(o.Enabled, nil),
		Inherited:         inheritedVal,
		InterfaceId:       o.InterfaceId,
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressXml_11_0_2) MarshalFromObject(s Layer3Ipv6Address) {
	o.Name = s.Name
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	if s.Prefix != nil {
		var obj layer3Ipv6AddressPrefixXml_11_0_2
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	if s.Anycast != nil {
		var obj layer3Ipv6AddressAnycastXml_11_0_2
		obj.MarshalFromObject(*s.Anycast)
		o.Anycast = &obj
	}
	if s.Advertise != nil {
		var obj layer3Ipv6AddressAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6Address, error) {
	var prefixVal *Layer3Ipv6AddressPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}
	var anycastVal *Layer3Ipv6AddressAnycast
	if o.Anycast != nil {
		obj, err := o.Anycast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anycastVal = obj
	}
	var advertiseVal *Layer3Ipv6AddressAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6Address{
		Name:              o.Name,
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Prefix:            prefixVal,
		Anycast:           anycastVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressPrefixXml_11_0_2) MarshalFromObject(s Layer3Ipv6AddressPrefix) {
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressPrefixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6AddressPrefix, error) {

	result := &Layer3Ipv6AddressPrefix{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressAnycastXml_11_0_2) MarshalFromObject(s Layer3Ipv6AddressAnycast) {
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressAnycastXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6AddressAnycast, error) {

	result := &Layer3Ipv6AddressAnycast{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6AddressAdvertiseXml_11_0_2) MarshalFromObject(s Layer3Ipv6AddressAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6AddressAdvertiseXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6AddressAdvertise, error) {

	result := &Layer3Ipv6AddressAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClient) {
	o.AcceptRaRoute = util.YesNo(s.AcceptRaRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Preference = s.Preference
	if s.PrefixDelegation != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationXml_11_0_2
		obj.MarshalFromObject(*s.PrefixDelegation)
		o.PrefixDelegation = &obj
	}
	if s.V6Options != nil {
		var obj layer3Ipv6DhcpClientV6OptionsXml_11_0_2
		obj.MarshalFromObject(*s.V6Options)
		o.V6Options = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClient, error) {
	var neighborDiscoveryVal *Layer3Ipv6DhcpClientNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}
	var prefixDelegationVal *Layer3Ipv6DhcpClientPrefixDelegation
	if o.PrefixDelegation != nil {
		obj, err := o.PrefixDelegation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixDelegationVal = obj
	}
	var v6OptionsVal *Layer3Ipv6DhcpClientV6Options
	if o.V6Options != nil {
		obj, err := o.V6Options.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v6OptionsVal = obj
	}

	result := &Layer3Ipv6DhcpClient{
		AcceptRaRoute:      util.AsBool(o.AcceptRaRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		NeighborDiscovery:  neighborDiscoveryVal,
		Preference:         o.Preference,
		PrefixDelegation:   prefixDelegationVal,
		V6Options:          v6OptionsVal,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscovery, error) {
	var dnsServerVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscovery{
		DadAttempts:      o.DadAttempts,
		DnsServer:        dnsServerVal,
		DnsSuffix:        dnsSuffixVal,
		EnableDad:        util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor: util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:         neighborVal,
		NsInterval:       o.NsInterval,
		ReachableTime:    o.ReachableTime,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer, error) {
	var sourceVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
		for _, elt := range s.Server {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{
		Server: serverVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegation) {
	if s.Enable != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegation, error) {
	var enableVal *Layer3Ipv6DhcpClientPrefixDelegationEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Layer3Ipv6DhcpClientPrefixDelegation{
		Enable: enableVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnable) {
	if s.No != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj layer3Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnable, error) {
	var noVal *Layer3Ipv6DhcpClientPrefixDelegationEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Layer3Ipv6DhcpClientPrefixDelegationEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnable{
		No:   noVal,
		Yes:  yesVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnableNo) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnableNo, error) {

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnableNo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientPrefixDelegationEnableYes) {
	o.PfxPoolName = s.PfxPoolName
	o.PrefixLen = s.PrefixLen
	o.PrefixLenHint = util.YesNo(s.PrefixLenHint, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientPrefixDelegationEnableYes, error) {

	result := &Layer3Ipv6DhcpClientPrefixDelegationEnableYes{
		PfxPoolName:   o.PfxPoolName,
		PrefixLen:     o.PrefixLen,
		PrefixLenHint: util.AsBool(o.PrefixLenHint, nil),
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientV6Options) {
	o.DuidType = s.DuidType
	if s.Enable != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableXml_11_0_2
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.RapidCommit = util.YesNo(s.RapidCommit, nil)
	o.SupportSrvrReconfig = util.YesNo(s.SupportSrvrReconfig, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6Options, error) {
	var enableVal *Layer3Ipv6DhcpClientV6OptionsEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Layer3Ipv6DhcpClientV6Options{
		DuidType:            o.DuidType,
		Enable:              enableVal,
		RapidCommit:         util.AsBool(o.RapidCommit, nil),
		SupportSrvrReconfig: util.AsBool(o.SupportSrvrReconfig, nil),
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnable) {
	if s.No != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj layer3Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnable, error) {
	var noVal *Layer3Ipv6DhcpClientV6OptionsEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Layer3Ipv6DhcpClientV6OptionsEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Layer3Ipv6DhcpClientV6OptionsEnable{
		No:   noVal,
		Yes:  yesVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnableNo) {
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnableNo, error) {

	result := &Layer3Ipv6DhcpClientV6OptionsEnableNo{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2) MarshalFromObject(s Layer3Ipv6DhcpClientV6OptionsEnableYes) {
	o.NonTempAddr = util.YesNo(s.NonTempAddr, nil)
	o.TempAddr = util.YesNo(s.TempAddr, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6DhcpClientV6OptionsEnableYes, error) {

	result := &Layer3Ipv6DhcpClientV6OptionsEnableYes{
		NonTempAddr: util.AsBool(o.NonTempAddr, nil),
		TempAddr:    util.AsBool(o.TempAddr, nil),
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedXml_11_0_2) MarshalFromObject(s Layer3Ipv6Inherited) {
	if s.AssignAddr != nil {
		var objs []layer3Ipv6InheritedAssignAddrXml_11_0_2
		for _, elt := range s.AssignAddr {
			var obj layer3Ipv6InheritedAssignAddrXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AssignAddr = &layer3Ipv6InheritedAssignAddrContainerXml_11_0_2{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6Inherited, error) {
	var assignAddrVal []Layer3Ipv6InheritedAssignAddr
	if o.AssignAddr != nil {
		for _, elt := range o.AssignAddr.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			assignAddrVal = append(assignAddrVal, *obj)
		}
	}
	var neighborDiscoveryVal *Layer3Ipv6InheritedNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Layer3Ipv6Inherited{
		AssignAddr:        assignAddrVal,
		Enable:            util.AsBool(o.Enable, nil),
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddr) {
	o.Name = s.Name
	if s.Type != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddr, error) {
	var typeVal *Layer3Ipv6InheritedAssignAddrType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddr{
		Name: o.Name,
		Type: typeVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrType) {
	if s.Gua != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaXml_11_0_2
		obj.MarshalFromObject(*s.Gua)
		o.Gua = &obj
	}
	if s.Ula != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeUlaXml_11_0_2
		obj.MarshalFromObject(*s.Ula)
		o.Ula = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrType, error) {
	var guaVal *Layer3Ipv6InheritedAssignAddrTypeGua
	if o.Gua != nil {
		obj, err := o.Gua.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		guaVal = obj
	}
	var ulaVal *Layer3Ipv6InheritedAssignAddrTypeUla
	if o.Ula != nil {
		obj, err := o.Ula.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ulaVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrType{
		Gua:  guaVal,
		Ula:  ulaVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGua) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.PrefixPool = s.PrefixPool
	if s.PoolType != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2
		obj.MarshalFromObject(*s.PoolType)
		o.PoolType = &obj
	}
	if s.Advertise != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGua, error) {
	var poolTypeVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType
	if o.PoolType != nil {
		obj, err := o.PoolType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		poolTypeVal = obj
	}
	var advertiseVal *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeGua{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		PrefixPool:        o.PrefixPool,
		PoolType:          poolTypeVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolType) {
	if s.Dynamic != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.DynamicId != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2
		obj.MarshalFromObject(*s.DynamicId)
		o.DynamicId = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolType, error) {
	var dynamicVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var dynamicIdVal *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	if o.DynamicId != nil {
		obj, err := o.DynamicId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIdVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolType{
		Dynamic:   dynamicVal,
		DynamicId: dynamicIdVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) {
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) {
	o.Identifier = s.Identifier
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{
		Identifier: o.Identifier,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise{
		Enable:         util.AsBool(o.Enable, nil),
		OnlinkFlag:     util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag: util.AsBool(o.AutoConfigFlag, nil),
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeUlaXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeUla) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.Address = s.Address
	o.Prefix = util.YesNo(s.Prefix, nil)
	o.Anycast = util.YesNo(s.Anycast, nil)
	if s.Advertise != nil {
		var obj layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeUlaXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeUla, error) {
	var advertiseVal *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Layer3Ipv6InheritedAssignAddrTypeUla{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Address:           o.Address,
		Prefix:            util.AsBool(o.Prefix, nil),
		Anycast:           util.AsBool(o.Anycast, nil),
		Advertise:         advertiseVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise, error) {

	result := &Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscovery, error) {
	var dnsServerVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Layer3Ipv6InheritedNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		DnsServer:           dnsServerVal,
		DnsSuffix:           dnsSuffixVal,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServer, error) {
	var sourceVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServer{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{
		PrefixPool: o.PrefixPool,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
		for _, elt := range s.Server {
			var obj layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{
		Server: serverVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix{
		Enable: util.AsBool(o.Enable, nil),
		Source: sourceVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource{
		Dhcpv6: dhcpv6Val,
		Manual: manualVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{
		PrefixPool: o.PrefixPool,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2) MarshalFromObject(s Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.EnableConsistencyCheck = util.YesNo(s.EnableConsistencyCheck, nil)
	o.HopLimit = s.HopLimit
	o.Lifetime = s.Lifetime
	o.LinkMtu = s.LinkMtu
	o.ManagedFlag = util.YesNo(s.ManagedFlag, nil)
	o.MaxInterval = s.MaxInterval
	o.MinInterval = s.MinInterval
	o.OtherFlag = util.YesNo(s.OtherFlag, nil)
	o.ReachableTime = s.ReachableTime
	o.RetransmissionTimer = s.RetransmissionTimer
	o.RouterPreference = s.RouterPreference
	o.Misc = s.Misc
}

func (o layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement, error) {

	result := &Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement{
		Enable:                 util.AsBool(o.Enable, nil),
		EnableConsistencyCheck: util.AsBool(o.EnableConsistencyCheck, nil),
		HopLimit:               o.HopLimit,
		Lifetime:               o.Lifetime,
		LinkMtu:                o.LinkMtu,
		ManagedFlag:            util.AsBool(o.ManagedFlag, nil),
		MaxInterval:            o.MaxInterval,
		MinInterval:            o.MinInterval,
		OtherFlag:              util.AsBool(o.OtherFlag, nil),
		ReachableTime:          o.ReachableTime,
		RetransmissionTimer:    o.RetransmissionTimer,
		RouterPreference:       o.RouterPreference,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &layer3Ipv6NeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscovery, error) {
	var neighborVal []Layer3Ipv6NeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Layer3Ipv6NeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Layer3Ipv6NeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryNeighbor, error) {

	result := &Layer3Ipv6NeighborDiscoveryNeighbor{
		Name:      o.Name,
		HwAddress: o.HwAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisement) {
	if s.DnsSupport != nil {
		var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2
		obj.MarshalFromObject(*s.DnsSupport)
		o.DnsSupport = &obj
	}
	o.Enable = util.YesNo(s.Enable, nil)
	o.EnableConsistencyCheck = util.YesNo(s.EnableConsistencyCheck, nil)
	o.HopLimit = s.HopLimit
	o.Lifetime = s.Lifetime
	o.LinkMtu = s.LinkMtu
	o.ManagedFlag = util.YesNo(s.ManagedFlag, nil)
	o.MaxInterval = s.MaxInterval
	o.MinInterval = s.MinInterval
	o.OtherFlag = util.YesNo(s.OtherFlag, nil)
	o.ReachableTime = s.ReachableTime
	o.RetransmissionTimer = s.RetransmissionTimer
	o.RouterPreference = s.RouterPreference
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisement, error) {
	var dnsSupportVal *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
	if o.DnsSupport != nil {
		obj, err := o.DnsSupport.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSupportVal = obj
	}

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisement{
		DnsSupport:             dnsSupportVal,
		Enable:                 util.AsBool(o.Enable, nil),
		EnableConsistencyCheck: util.AsBool(o.EnableConsistencyCheck, nil),
		HopLimit:               o.HopLimit,
		Lifetime:               o.Lifetime,
		LinkMtu:                o.LinkMtu,
		ManagedFlag:            util.AsBool(o.ManagedFlag, nil),
		MaxInterval:            o.MaxInterval,
		MinInterval:            o.MinInterval,
		OtherFlag:              util.AsBool(o.OtherFlag, nil),
		ReachableTime:          o.ReachableTime,
		RetransmissionTimer:    o.RetransmissionTimer,
		RouterPreference:       o.RouterPreference,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Server != nil {
		var objs []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2
		for _, elt := range s.Server {
			var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2{Entries: objs}
	}
	if s.Suffix != nil {
		var objs []layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, error) {
	var serverVal []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}
	var suffixVal []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{
		Enable: util.AsBool(o.Enable, nil),
		Server: serverVal,
		Suffix: suffixVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, error) {

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2) MarshalFromObject(s Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
}

func (o layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2) UnmarshalToObject() (*Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, error) {

	result := &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{
		Name:     o.Name,
		Lifetime: o.Lifetime,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3LacpXml_11_0_2) MarshalFromObject(s Layer3Lacp) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FastFailover = util.YesNo(s.FastFailover, nil)
	if s.HighAvailability != nil {
		var obj layer3LacpHighAvailabilityXml_11_0_2
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.MaxPorts = s.MaxPorts
	o.Mode = s.Mode
	o.SystemPriority = s.SystemPriority
	o.TransmissionRate = s.TransmissionRate
	o.Misc = s.Misc
}

func (o layer3LacpXml_11_0_2) UnmarshalToObject() (*Layer3Lacp, error) {
	var highAvailabilityVal *Layer3LacpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer3Lacp{
		Enable:           util.AsBool(o.Enable, nil),
		FastFailover:     util.AsBool(o.FastFailover, nil),
		HighAvailability: highAvailabilityVal,
		MaxPorts:         o.MaxPorts,
		Mode:             o.Mode,
		SystemPriority:   o.SystemPriority,
		TransmissionRate: o.TransmissionRate,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3LacpHighAvailabilityXml_11_0_2) MarshalFromObject(s Layer3LacpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer3LacpHighAvailabilityXml_11_0_2) UnmarshalToObject() (*Layer3LacpHighAvailability, error) {

	result := &Layer3LacpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3LldpXml_11_0_2) MarshalFromObject(s Layer3Lldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj layer3LldpHighAvailabilityXml_11_0_2
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o layer3LldpXml_11_0_2) UnmarshalToObject() (*Layer3Lldp, error) {
	var highAvailabilityVal *Layer3LldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &Layer3Lldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *layer3LldpHighAvailabilityXml_11_0_2) MarshalFromObject(s Layer3LldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o layer3LldpHighAvailabilityXml_11_0_2) UnmarshalToObject() (*Layer3LldpHighAvailability, error) {

	result := &Layer3LldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3NdpProxyXml_11_0_2) MarshalFromObject(s Layer3NdpProxy) {
	if s.Address != nil {
		var objs []layer3NdpProxyAddressXml_11_0_2
		for _, elt := range s.Address {
			var obj layer3NdpProxyAddressXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &layer3NdpProxyAddressContainerXml_11_0_2{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.Misc = s.Misc
}

func (o layer3NdpProxyXml_11_0_2) UnmarshalToObject() (*Layer3NdpProxy, error) {
	var addressVal []Layer3NdpProxyAddress
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &Layer3NdpProxy{
		Address: addressVal,
		Enabled: util.AsBool(o.Enabled, nil),
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *layer3NdpProxyAddressXml_11_0_2) MarshalFromObject(s Layer3NdpProxyAddress) {
	o.Name = s.Name
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
}

func (o layer3NdpProxyAddressXml_11_0_2) UnmarshalToObject() (*Layer3NdpProxyAddress, error) {

	result := &Layer3NdpProxyAddress{
		Name:   o.Name,
		Negate: util.AsBool(o.Negate, nil),
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsXml_11_0_2) MarshalFromObject(s Layer3SdwanLinkSettings) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	if s.UpstreamNat != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatXml_11_0_2
		obj.MarshalFromObject(*s.UpstreamNat)
		o.UpstreamNat = &obj
	}
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsXml_11_0_2) UnmarshalToObject() (*Layer3SdwanLinkSettings, error) {
	var upstreamNatVal *Layer3SdwanLinkSettingsUpstreamNat
	if o.UpstreamNat != nil {
		obj, err := o.UpstreamNat.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		upstreamNatVal = obj
	}

	result := &Layer3SdwanLinkSettings{
		Enable:                util.AsBool(o.Enable, nil),
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		UpstreamNat:           upstreamNatVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatXml_11_0_2) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNat) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Ddns != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatDdnsXml_11_0_2
		obj.MarshalFromObject(*s.Ddns)
		o.Ddns = &obj
	}
	if s.StaticIp != nil {
		var obj layer3SdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2
		obj.MarshalFromObject(*s.StaticIp)
		o.StaticIp = &obj
	}
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatXml_11_0_2) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNat, error) {
	var ddnsVal *Layer3SdwanLinkSettingsUpstreamNatDdns
	if o.Ddns != nil {
		obj, err := o.Ddns.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsVal = obj
	}
	var staticIpVal *Layer3SdwanLinkSettingsUpstreamNatStaticIp
	if o.StaticIp != nil {
		obj, err := o.StaticIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticIpVal = obj
	}

	result := &Layer3SdwanLinkSettingsUpstreamNat{
		Enable:   util.AsBool(o.Enable, nil),
		Ddns:     ddnsVal,
		StaticIp: staticIpVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatDdnsXml_11_0_2) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNatDdns) {
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatDdnsXml_11_0_2) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNatDdns, error) {

	result := &Layer3SdwanLinkSettingsUpstreamNatDdns{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *layer3SdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2) MarshalFromObject(s Layer3SdwanLinkSettingsUpstreamNatStaticIp) {
	o.Fqdn = s.Fqdn
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
}

func (o layer3SdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2) UnmarshalToObject() (*Layer3SdwanLinkSettingsUpstreamNatStaticIp, error) {

	result := &Layer3SdwanLinkSettingsUpstreamNatStaticIp{
		Fqdn:      o.Fqdn,
		IpAddress: o.IpAddress,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *virtualWireXml_11_0_2) MarshalFromObject(s VirtualWire) {
	if s.Lldp != nil {
		var obj virtualWireLldpXml_11_0_2
		obj.MarshalFromObject(*s.Lldp)
		o.Lldp = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	o.Misc = s.Misc
}

func (o virtualWireXml_11_0_2) UnmarshalToObject() (*VirtualWire, error) {
	var lldpVal *VirtualWireLldp
	if o.Lldp != nil {
		obj, err := o.Lldp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lldpVal = obj
	}

	result := &VirtualWire{
		Lldp:           lldpVal,
		NetflowProfile: o.NetflowProfile,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *virtualWireLldpXml_11_0_2) MarshalFromObject(s VirtualWireLldp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.HighAvailability != nil {
		var obj virtualWireLldpHighAvailabilityXml_11_0_2
		obj.MarshalFromObject(*s.HighAvailability)
		o.HighAvailability = &obj
	}
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o virtualWireLldpXml_11_0_2) UnmarshalToObject() (*VirtualWireLldp, error) {
	var highAvailabilityVal *VirtualWireLldpHighAvailability
	if o.HighAvailability != nil {
		obj, err := o.HighAvailability.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		highAvailabilityVal = obj
	}

	result := &VirtualWireLldp{
		Enable:           util.AsBool(o.Enable, nil),
		HighAvailability: highAvailabilityVal,
		Profile:          o.Profile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *virtualWireLldpHighAvailabilityXml_11_0_2) MarshalFromObject(s VirtualWireLldpHighAvailability) {
	o.PassivePreNegotiation = util.YesNo(s.PassivePreNegotiation, nil)
	o.Misc = s.Misc
}

func (o virtualWireLldpHighAvailabilityXml_11_0_2) UnmarshalToObject() (*VirtualWireLldpHighAvailability, error) {

	result := &VirtualWireLldpHighAvailability{
		PassivePreNegotiation: util.AsBool(o.PassivePreNegotiation, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "decrypt_mirror" || v == "DecryptMirror" {
		return e.DecryptMirror, nil
	}
	if v == "ha" || v == "Ha" {
		return e.Ha, nil
	}
	if v == "layer2" || v == "Layer2" {
		return e.Layer2, nil
	}
	if v == "layer3" || v == "Layer3" {
		return e.Layer3, nil
	}
	if v == "virtual_wire" || v == "VirtualWire" {
		return e.VirtualWire, nil
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
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !o.DecryptMirror.matches(other.DecryptMirror) {
		return false
	}
	if !o.Ha.matches(other.Ha) {
		return false
	}
	if !o.Layer2.matches(other.Layer2) {
		return false
	}
	if !o.Layer3.matches(other.Layer3) {
		return false
	}
	if !o.VirtualWire.matches(other.VirtualWire) {
		return false
	}

	return true
}

func (o *DecryptMirror) matches(other *DecryptMirror) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ha) matches(other *Ha) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Lacp.matches(other.Lacp) {
		return false
	}

	return true
}

func (o *HaLacp) matches(other *HaLacp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.FastFailover, other.FastFailover) {
		return false
	}
	if !util.Ints64Match(o.MaxPorts, other.MaxPorts) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !util.Ints64Match(o.SystemPriority, other.SystemPriority) {
		return false
	}
	if !util.StringsMatch(o.TransmissionRate, other.TransmissionRate) {
		return false
	}

	return true
}

func (o *Layer2) matches(other *Layer2) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Lacp.matches(other.Lacp) {
		return false
	}
	if !o.Lldp.matches(other.Lldp) {
		return false
	}
	if !util.StringsMatch(o.NetflowProfile, other.NetflowProfile) {
		return false
	}

	return true
}

func (o *Layer2Lacp) matches(other *Layer2Lacp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.FastFailover, other.FastFailover) {
		return false
	}
	if !o.HighAvailability.matches(other.HighAvailability) {
		return false
	}
	if !util.Ints64Match(o.MaxPorts, other.MaxPorts) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !util.Ints64Match(o.SystemPriority, other.SystemPriority) {
		return false
	}
	if !util.StringsMatch(o.TransmissionRate, other.TransmissionRate) {
		return false
	}

	return true
}

func (o *Layer2LacpHighAvailability) matches(other *Layer2LacpHighAvailability) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.PassivePreNegotiation, other.PassivePreNegotiation) {
		return false
	}

	return true
}

func (o *Layer2Lldp) matches(other *Layer2Lldp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.HighAvailability.matches(other.HighAvailability) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *Layer2LldpHighAvailability) matches(other *Layer2LldpHighAvailability) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.PassivePreNegotiation, other.PassivePreNegotiation) {
		return false
	}

	return true
}

func (o *Layer3) matches(other *Layer3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AdjustTcpMss.matches(other.AdjustTcpMss) {
		return false
	}
	if len(o.Arp) != len(other.Arp) {
		return false
	}
	for idx := range o.Arp {
		if !o.Arp[idx].matches(&other.Arp[idx]) {
			return false
		}
	}
	if !o.Bonjour.matches(other.Bonjour) {
		return false
	}
	if !o.DdnsConfig.matches(other.DdnsConfig) {
		return false
	}
	if !util.BoolsMatch(o.DecryptForward, other.DecryptForward) {
		return false
	}
	if !util.BoolsMatch(o.DfIgnore, other.DfIgnore) {
		return false
	}
	if !o.DhcpClient.matches(other.DhcpClient) {
		return false
	}
	if !util.StringsMatch(o.InterfaceManagementProfile, other.InterfaceManagementProfile) {
		return false
	}
	if len(o.Ip) != len(other.Ip) {
		return false
	}
	for idx := range o.Ip {
		if !o.Ip[idx].matches(&other.Ip[idx]) {
			return false
		}
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}
	if !o.Lacp.matches(other.Lacp) {
		return false
	}
	if !o.Lldp.matches(other.Lldp) {
		return false
	}
	if !util.Ints64Match(o.Mtu, other.Mtu) {
		return false
	}
	if !o.NdpProxy.matches(other.NdpProxy) {
		return false
	}
	if !util.StringsMatch(o.NetflowProfile, other.NetflowProfile) {
		return false
	}
	if !o.SdwanLinkSettings.matches(other.SdwanLinkSettings) {
		return false
	}
	if !util.BoolsMatch(o.UntaggedSubInterface, other.UntaggedSubInterface) {
		return false
	}

	return true
}

func (o *Layer3AdjustTcpMss) matches(other *Layer3AdjustTcpMss) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Ipv4MssAdjustment, other.Ipv4MssAdjustment) {
		return false
	}
	if !util.Ints64Match(o.Ipv6MssAdjustment, other.Ipv6MssAdjustment) {
		return false
	}

	return true
}

func (o *Layer3Arp) matches(other *Layer3Arp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.HwAddress, other.HwAddress) {
		return false
	}

	return true
}

func (o *Layer3Bonjour) matches(other *Layer3Bonjour) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.GroupId, other.GroupId) {
		return false
	}
	if !util.BoolsMatch(o.TtlCheck, other.TtlCheck) {
		return false
	}

	return true
}

func (o *Layer3DdnsConfig) matches(other *Layer3DdnsConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DdnsCertProfile, other.DdnsCertProfile) {
		return false
	}
	if !util.BoolsMatch(o.DdnsEnabled, other.DdnsEnabled) {
		return false
	}
	if !util.StringsMatch(o.DdnsHostname, other.DdnsHostname) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DdnsIp, other.DdnsIp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DdnsIpv6, other.DdnsIpv6) {
		return false
	}
	if !util.Ints64Match(o.DdnsUpdateInterval, other.DdnsUpdateInterval) {
		return false
	}
	if !util.StringsMatch(o.DdnsVendor, other.DdnsVendor) {
		return false
	}
	if len(o.DdnsVendorConfig) != len(other.DdnsVendorConfig) {
		return false
	}
	for idx := range o.DdnsVendorConfig {
		if !o.DdnsVendorConfig[idx].matches(&other.DdnsVendorConfig[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3DdnsConfigDdnsVendorConfig) matches(other *Layer3DdnsConfigDdnsVendorConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}

	return true
}

func (o *Layer3DhcpClient) matches(other *Layer3DhcpClient) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.CreateDefaultRoute, other.CreateDefaultRoute) {
		return false
	}
	if !util.Ints64Match(o.DefaultRouteMetric, other.DefaultRouteMetric) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.SendHostname.matches(other.SendHostname) {
		return false
	}

	return true
}

func (o *Layer3DhcpClientSendHostname) matches(other *Layer3DhcpClientSendHostname) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}

	return true
}

func (o *Layer3Ip) matches(other *Layer3Ip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.SdwanGateway, other.SdwanGateway) {
		return false
	}

	return true
}

func (o *Layer3Ipv6) matches(other *Layer3Ipv6) bool {
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
	if !o.DhcpClient.matches(other.DhcpClient) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !o.Inherited.matches(other.Inherited) {
		return false
	}
	if !util.StringsMatch(o.InterfaceId, other.InterfaceId) {
		return false
	}
	if !o.NeighborDiscovery.matches(other.NeighborDiscovery) {
		return false
	}

	return true
}

func (o *Layer3Ipv6Address) matches(other *Layer3Ipv6Address) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.EnableOnInterface, other.EnableOnInterface) {
		return false
	}
	if !o.Prefix.matches(other.Prefix) {
		return false
	}
	if !o.Anycast.matches(other.Anycast) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *Layer3Ipv6AddressPrefix) matches(other *Layer3Ipv6AddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6AddressAnycast) matches(other *Layer3Ipv6AddressAnycast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6AddressAdvertise) matches(other *Layer3Ipv6AddressAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.ValidLifetime, other.ValidLifetime) {
		return false
	}
	if !util.StringsMatch(o.PreferredLifetime, other.PreferredLifetime) {
		return false
	}
	if !util.BoolsMatch(o.OnlinkFlag, other.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(o.AutoConfigFlag, other.AutoConfigFlag) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClient) matches(other *Layer3Ipv6DhcpClient) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AcceptRaRoute, other.AcceptRaRoute) {
		return false
	}
	if !util.Ints64Match(o.DefaultRouteMetric, other.DefaultRouteMetric) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.NeighborDiscovery.matches(other.NeighborDiscovery) {
		return false
	}
	if !util.StringsMatch(o.Preference, other.Preference) {
		return false
	}
	if !o.PrefixDelegation.matches(other.PrefixDelegation) {
		return false
	}
	if !o.V6Options.matches(other.V6Options) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscovery) matches(other *Layer3Ipv6DhcpClientNeighborDiscovery) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.DadAttempts, other.DadAttempts) {
		return false
	}
	if !o.DnsServer.matches(other.DnsServer) {
		return false
	}
	if !o.DnsSuffix.matches(other.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(o.EnableDad, other.EnableDad) {
		return false
	}
	if !util.BoolsMatch(o.EnableNdpMonitor, other.EnableNdpMonitor) {
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
	if !util.Ints64Match(o.NsInterval, other.NsInterval) {
		return false
	}
	if !util.Ints64Match(o.ReachableTime, other.ReachableTime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Source.matches(other.Source) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dhcpv6.matches(other.Dhcpv6) {
		return false
	}
	if !o.Manual.matches(other.Manual) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Server) != len(other.Server) {
		return false
	}
	for idx := range o.Server {
		if !o.Server[idx].matches(&other.Server[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Source.matches(other.Source) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dhcpv6.matches(other.Dhcpv6) {
		return false
	}
	if !o.Manual.matches(other.Manual) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Suffix) != len(other.Suffix) {
		return false
	}
	for idx := range o.Suffix {
		if !o.Suffix[idx].matches(&other.Suffix[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor) matches(other *Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.HwAddress, other.HwAddress) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientPrefixDelegation) matches(other *Layer3Ipv6DhcpClientPrefixDelegation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Enable.matches(other.Enable) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientPrefixDelegationEnable) matches(other *Layer3Ipv6DhcpClientPrefixDelegationEnable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.No.matches(other.No) {
		return false
	}
	if !o.Yes.matches(other.Yes) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientPrefixDelegationEnableNo) matches(other *Layer3Ipv6DhcpClientPrefixDelegationEnableNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientPrefixDelegationEnableYes) matches(other *Layer3Ipv6DhcpClientPrefixDelegationEnableYes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.PfxPoolName, other.PfxPoolName) {
		return false
	}
	if !util.Ints64Match(o.PrefixLen, other.PrefixLen) {
		return false
	}
	if !util.BoolsMatch(o.PrefixLenHint, other.PrefixLenHint) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientV6Options) matches(other *Layer3Ipv6DhcpClientV6Options) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DuidType, other.DuidType) {
		return false
	}
	if !o.Enable.matches(other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.RapidCommit, other.RapidCommit) {
		return false
	}
	if !util.BoolsMatch(o.SupportSrvrReconfig, other.SupportSrvrReconfig) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientV6OptionsEnable) matches(other *Layer3Ipv6DhcpClientV6OptionsEnable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.No.matches(other.No) {
		return false
	}
	if !o.Yes.matches(other.Yes) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientV6OptionsEnableNo) matches(other *Layer3Ipv6DhcpClientV6OptionsEnableNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6DhcpClientV6OptionsEnableYes) matches(other *Layer3Ipv6DhcpClientV6OptionsEnableYes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.NonTempAddr, other.NonTempAddr) {
		return false
	}
	if !util.BoolsMatch(o.TempAddr, other.TempAddr) {
		return false
	}

	return true
}

func (o *Layer3Ipv6Inherited) matches(other *Layer3Ipv6Inherited) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.AssignAddr) != len(other.AssignAddr) {
		return false
	}
	for idx := range o.AssignAddr {
		if !o.AssignAddr[idx].matches(&other.AssignAddr[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.NeighborDiscovery.matches(other.NeighborDiscovery) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddr) matches(other *Layer3Ipv6InheritedAssignAddr) bool {
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

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrType) matches(other *Layer3Ipv6InheritedAssignAddrType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Gua.matches(other.Gua) {
		return false
	}
	if !o.Ula.matches(other.Ula) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeGua) matches(other *Layer3Ipv6InheritedAssignAddrTypeGua) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.EnableOnInterface, other.EnableOnInterface) {
		return false
	}
	if !util.StringsMatch(o.PrefixPool, other.PrefixPool) {
		return false
	}
	if !o.PoolType.matches(other.PoolType) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType) matches(other *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dynamic.matches(other.Dynamic) {
		return false
	}
	if !o.DynamicId.matches(other.DynamicId) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) matches(other *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) matches(other *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Identifier, other.Identifier) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise) matches(other *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.OnlinkFlag, other.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(o.AutoConfigFlag, other.AutoConfigFlag) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeUla) matches(other *Layer3Ipv6InheritedAssignAddrTypeUla) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.EnableOnInterface, other.EnableOnInterface) {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.BoolsMatch(o.Prefix, other.Prefix) {
		return false
	}
	if !util.BoolsMatch(o.Anycast, other.Anycast) {
		return false
	}
	if !o.Advertise.matches(other.Advertise) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise) matches(other *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.ValidLifetime, other.ValidLifetime) {
		return false
	}
	if !util.StringsMatch(o.PreferredLifetime, other.PreferredLifetime) {
		return false
	}
	if !util.BoolsMatch(o.OnlinkFlag, other.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(o.AutoConfigFlag, other.AutoConfigFlag) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscovery) matches(other *Layer3Ipv6InheritedNeighborDiscovery) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.DadAttempts, other.DadAttempts) {
		return false
	}
	if !o.DnsServer.matches(other.DnsServer) {
		return false
	}
	if !o.DnsSuffix.matches(other.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(o.EnableDad, other.EnableDad) {
		return false
	}
	if !util.BoolsMatch(o.EnableNdpMonitor, other.EnableNdpMonitor) {
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
	if !util.Ints64Match(o.NsInterval, other.NsInterval) {
		return false
	}
	if !util.Ints64Match(o.ReachableTime, other.ReachableTime) {
		return false
	}
	if !o.RouterAdvertisement.matches(other.RouterAdvertisement) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsServer) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Source.matches(other.Source) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dhcpv6.matches(other.Dhcpv6) {
		return false
	}
	if !o.Manual.matches(other.Manual) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.PrefixPool, other.PrefixPool) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Server) != len(other.Server) {
		return false
	}
	for idx := range o.Server {
		if !o.Server[idx].matches(&other.Server[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Source.matches(other.Source) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dhcpv6.matches(other.Dhcpv6) {
		return false
	}
	if !o.Manual.matches(other.Manual) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.PrefixPool, other.PrefixPool) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Suffix) != len(other.Suffix) {
		return false
	}
	for idx := range o.Suffix {
		if !o.Suffix[idx].matches(&other.Suffix[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) matches(other *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryNeighbor) matches(other *Layer3Ipv6InheritedNeighborDiscoveryNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.HwAddress, other.HwAddress) {
		return false
	}

	return true
}

func (o *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement) matches(other *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.EnableConsistencyCheck, other.EnableConsistencyCheck) {
		return false
	}
	if !util.StringsMatch(o.HopLimit, other.HopLimit) {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}
	if !util.StringsMatch(o.LinkMtu, other.LinkMtu) {
		return false
	}
	if !util.BoolsMatch(o.ManagedFlag, other.ManagedFlag) {
		return false
	}
	if !util.Ints64Match(o.MaxInterval, other.MaxInterval) {
		return false
	}
	if !util.Ints64Match(o.MinInterval, other.MinInterval) {
		return false
	}
	if !util.BoolsMatch(o.OtherFlag, other.OtherFlag) {
		return false
	}
	if !util.StringsMatch(o.ReachableTime, other.ReachableTime) {
		return false
	}
	if !util.StringsMatch(o.RetransmissionTimer, other.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(o.RouterPreference, other.RouterPreference) {
		return false
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscovery) matches(other *Layer3Ipv6NeighborDiscovery) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.DadAttempts, other.DadAttempts) {
		return false
	}
	if !util.BoolsMatch(o.EnableDad, other.EnableDad) {
		return false
	}
	if !util.BoolsMatch(o.EnableNdpMonitor, other.EnableNdpMonitor) {
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
	if !util.Ints64Match(o.NsInterval, other.NsInterval) {
		return false
	}
	if !util.Ints64Match(o.ReachableTime, other.ReachableTime) {
		return false
	}
	if !o.RouterAdvertisement.matches(other.RouterAdvertisement) {
		return false
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscoveryNeighbor) matches(other *Layer3Ipv6NeighborDiscoveryNeighbor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.HwAddress, other.HwAddress) {
		return false
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscoveryRouterAdvertisement) matches(other *Layer3Ipv6NeighborDiscoveryRouterAdvertisement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.DnsSupport.matches(other.DnsSupport) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.EnableConsistencyCheck, other.EnableConsistencyCheck) {
		return false
	}
	if !util.StringsMatch(o.HopLimit, other.HopLimit) {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}
	if !util.StringsMatch(o.LinkMtu, other.LinkMtu) {
		return false
	}
	if !util.BoolsMatch(o.ManagedFlag, other.ManagedFlag) {
		return false
	}
	if !util.Ints64Match(o.MaxInterval, other.MaxInterval) {
		return false
	}
	if !util.Ints64Match(o.MinInterval, other.MinInterval) {
		return false
	}
	if !util.BoolsMatch(o.OtherFlag, other.OtherFlag) {
		return false
	}
	if !util.StringsMatch(o.ReachableTime, other.ReachableTime) {
		return false
	}
	if !util.StringsMatch(o.RetransmissionTimer, other.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(o.RouterPreference, other.RouterPreference) {
		return false
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) matches(other *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if len(o.Server) != len(other.Server) {
		return false
	}
	for idx := range o.Server {
		if !o.Server[idx].matches(&other.Server[idx]) {
			return false
		}
	}
	if len(o.Suffix) != len(other.Suffix) {
		return false
	}
	for idx := range o.Suffix {
		if !o.Suffix[idx].matches(&other.Suffix[idx]) {
			return false
		}
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) matches(other *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) matches(other *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Lifetime, other.Lifetime) {
		return false
	}

	return true
}

func (o *Layer3Lacp) matches(other *Layer3Lacp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.BoolsMatch(o.FastFailover, other.FastFailover) {
		return false
	}
	if !o.HighAvailability.matches(other.HighAvailability) {
		return false
	}
	if !util.Ints64Match(o.MaxPorts, other.MaxPorts) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !util.Ints64Match(o.SystemPriority, other.SystemPriority) {
		return false
	}
	if !util.StringsMatch(o.TransmissionRate, other.TransmissionRate) {
		return false
	}

	return true
}

func (o *Layer3LacpHighAvailability) matches(other *Layer3LacpHighAvailability) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.PassivePreNegotiation, other.PassivePreNegotiation) {
		return false
	}

	return true
}

func (o *Layer3Lldp) matches(other *Layer3Lldp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.HighAvailability.matches(other.HighAvailability) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *Layer3LldpHighAvailability) matches(other *Layer3LldpHighAvailability) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.PassivePreNegotiation, other.PassivePreNegotiation) {
		return false
	}

	return true
}

func (o *Layer3NdpProxy) matches(other *Layer3NdpProxy) bool {
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
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}

	return true
}

func (o *Layer3NdpProxyAddress) matches(other *Layer3NdpProxyAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}

	return true
}

func (o *Layer3SdwanLinkSettings) matches(other *Layer3SdwanLinkSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaceProfile, other.SdwanInterfaceProfile) {
		return false
	}
	if !o.UpstreamNat.matches(other.UpstreamNat) {
		return false
	}

	return true
}

func (o *Layer3SdwanLinkSettingsUpstreamNat) matches(other *Layer3SdwanLinkSettingsUpstreamNat) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Ddns.matches(other.Ddns) {
		return false
	}
	if !o.StaticIp.matches(other.StaticIp) {
		return false
	}

	return true
}

func (o *Layer3SdwanLinkSettingsUpstreamNatDdns) matches(other *Layer3SdwanLinkSettingsUpstreamNatDdns) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Layer3SdwanLinkSettingsUpstreamNatStaticIp) matches(other *Layer3SdwanLinkSettingsUpstreamNatStaticIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}

	return true
}

func (o *VirtualWire) matches(other *VirtualWire) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Lldp.matches(other.Lldp) {
		return false
	}
	if !util.StringsMatch(o.NetflowProfile, other.NetflowProfile) {
		return false
	}

	return true
}

func (o *VirtualWireLldp) matches(other *VirtualWireLldp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.HighAvailability.matches(other.HighAvailability) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *VirtualWireLldpHighAvailability) matches(other *VirtualWireLldpHighAvailability) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.PassivePreNegotiation, other.PassivePreNegotiation) {
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
