package aggregate

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
	Suffix = []string{}
)

type Entry struct {
	Name          string
	Comment       *string
	DecryptMirror *DecryptMirror
	Ha            *Ha
	Layer2        *Layer2
	Layer3        *Layer3
	VirtualWire   *VirtualWire

	Misc map[string][]generic.Xml
}

type DecryptMirror struct {
}
type Ha struct {
	Lacp *HaLacp
}
type HaLacp struct {
	Enable           *bool
	FastFailover     *bool
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
}
type Layer2 struct {
	Lacp           *Layer2Lacp
	Lldp           *Layer2Lldp
	NetflowProfile *string
}
type Layer2Lacp struct {
	Enable           *bool
	FastFailover     *bool
	HighAvailability *Layer2LacpHighAvailability
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
}
type Layer2LacpHighAvailability struct {
	PassivePreNegotiation *bool
}
type Layer2Lldp struct {
	Enable           *bool
	HighAvailability *Layer2LldpHighAvailability
	Profile          *string
}
type Layer2LldpHighAvailability struct {
	PassivePreNegotiation *bool
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
}
type Layer3AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
}
type Layer3Arp struct {
	HwAddress *string
	Name      string
}
type Layer3Bonjour struct {
	Enable   *bool
	GroupId  *int64
	TtlCheck *bool
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
}
type Layer3DdnsConfigDdnsVendorConfig struct {
	Name  string
	Value *string
}
type Layer3DhcpClient struct {
	CreateDefaultRoute *bool
	DefaultRouteMetric *int64
	Enable             *bool
	SendHostname       *Layer3DhcpClientSendHostname
}
type Layer3DhcpClientSendHostname struct {
	Enable   *bool
	Hostname *string
}
type Layer3Ip struct {
	Name         string
	SdwanGateway *string
}
type Layer3Ipv6 struct {
	Address           []Layer3Ipv6Address
	DhcpClient        *Layer3Ipv6DhcpClient
	Enabled           *bool
	Inherited         *Layer3Ipv6Inherited
	InterfaceId       *string
	NeighborDiscovery *Layer3Ipv6NeighborDiscovery
}
type Layer3Ipv6Address struct {
	Advertise         *Layer3Ipv6AddressAdvertise
	Anycast           *Layer3Ipv6AddressAnycast
	EnableOnInterface *bool
	Name              string
	Prefix            *Layer3Ipv6AddressPrefix
}
type Layer3Ipv6AddressAdvertise struct {
	AutoConfigFlag    *bool
	Enable            *bool
	OnlinkFlag        *bool
	PreferredLifetime *string
	ValidLifetime     *string
}
type Layer3Ipv6AddressAnycast struct {
}
type Layer3Ipv6AddressPrefix struct {
}
type Layer3Ipv6DhcpClient struct {
	AcceptRaRoute      *bool
	DefaultRouteMetric *int64
	Enable             *bool
	NeighborDiscovery  *Layer3Ipv6DhcpClientNeighborDiscovery
	Preference         *string
	PrefixDelegation   *Layer3Ipv6DhcpClientPrefixDelegation
	V6Options          *Layer3Ipv6DhcpClientV6Options
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
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6 struct {
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual struct {
	Server []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor struct {
	HwAddress *string
	Name      string
}
type Layer3Ipv6DhcpClientPrefixDelegation struct {
	Enable *Layer3Ipv6DhcpClientPrefixDelegationEnable
}
type Layer3Ipv6DhcpClientPrefixDelegationEnable struct {
	No  *Layer3Ipv6DhcpClientPrefixDelegationEnableNo
	Yes *Layer3Ipv6DhcpClientPrefixDelegationEnableYes
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableNo struct {
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableYes struct {
	PfxPoolName   *string
	PrefixLen     *int64
	PrefixLenHint *bool
}
type Layer3Ipv6DhcpClientV6Options struct {
	DuidType            *string
	Enable              *Layer3Ipv6DhcpClientV6OptionsEnable
	RapidCommit         *bool
	SupportSrvrReconfig *bool
}
type Layer3Ipv6DhcpClientV6OptionsEnable struct {
	No  *Layer3Ipv6DhcpClientV6OptionsEnableNo
	Yes *Layer3Ipv6DhcpClientV6OptionsEnableYes
}
type Layer3Ipv6DhcpClientV6OptionsEnableNo struct {
}
type Layer3Ipv6DhcpClientV6OptionsEnableYes struct {
	NonTempAddr *bool
	TempAddr    *bool
}
type Layer3Ipv6Inherited struct {
	AssignAddr        []Layer3Ipv6InheritedAssignAddr
	Enable            *bool
	NeighborDiscovery *Layer3Ipv6InheritedNeighborDiscovery
}
type Layer3Ipv6InheritedAssignAddr struct {
	Name string
	Type *Layer3Ipv6InheritedAssignAddrType
}
type Layer3Ipv6InheritedAssignAddrType struct {
	Gua *Layer3Ipv6InheritedAssignAddrTypeGua
	Ula *Layer3Ipv6InheritedAssignAddrTypeUla
}
type Layer3Ipv6InheritedAssignAddrTypeGua struct {
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise
	EnableOnInterface *bool
	PoolType          *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType
	PrefixPool        *string
}
type Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise struct {
	AutoConfigFlag *bool
	Enable         *bool
	OnlinkFlag     *bool
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolType struct {
	Dynamic   *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	DynamicId *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic struct {
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId struct {
	Identifier *int64
}
type Layer3Ipv6InheritedAssignAddrTypeUla struct {
	Address           *string
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise
	Anycast           *bool
	EnableOnInterface *bool
	Prefix            *bool
}
type Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise struct {
	AutoConfigFlag    *bool
	Enable            *bool
	OnlinkFlag        *bool
	PreferredLifetime *string
	ValidLifetime     *string
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
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	PrefixPool *string
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual struct {
	Server []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	PrefixPool *string
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6InheritedNeighborDiscoveryNeighbor struct {
	HwAddress *string
	Name      string
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
}
type Layer3Ipv6NeighborDiscovery struct {
	DadAttempts         *int64
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Layer3Ipv6NeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Layer3Ipv6NeighborDiscoveryRouterAdvertisement
}
type Layer3Ipv6NeighborDiscoveryNeighbor struct {
	HwAddress *string
	Name      string
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
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport struct {
	Enable *bool
	Server []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	Suffix []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix struct {
	Lifetime *int64
	Name     string
}
type Layer3Lacp struct {
	Enable           *bool
	FastFailover     *bool
	HighAvailability *Layer3LacpHighAvailability
	MaxPorts         *int64
	Mode             *string
	SystemPriority   *int64
	TransmissionRate *string
}
type Layer3LacpHighAvailability struct {
	PassivePreNegotiation *bool
}
type Layer3Lldp struct {
	Enable           *bool
	HighAvailability *Layer3LldpHighAvailability
	Profile          *string
}
type Layer3LldpHighAvailability struct {
	PassivePreNegotiation *bool
}
type Layer3NdpProxy struct {
	Address []Layer3NdpProxyAddress
	Enabled *bool
}
type Layer3NdpProxyAddress struct {
	Name   string
	Negate *bool
}
type Layer3SdwanLinkSettings struct {
	Enable                *bool
	SdwanInterfaceProfile *string
	UpstreamNat           *Layer3SdwanLinkSettingsUpstreamNat
}
type Layer3SdwanLinkSettingsUpstreamNat struct {
	Enable   *bool
	Ddns     *Layer3SdwanLinkSettingsUpstreamNatDdns
	StaticIp *Layer3SdwanLinkSettingsUpstreamNatStaticIp
}
type Layer3SdwanLinkSettingsUpstreamNatDdns struct {
}
type Layer3SdwanLinkSettingsUpstreamNatStaticIp struct {
	Fqdn      *string
	IpAddress *string
}
type VirtualWire struct {
	Lldp           *VirtualWireLldp
	NetflowProfile *string
}
type VirtualWireLldp struct {
	Enable           *bool
	HighAvailability *VirtualWireLldpHighAvailability
	Profile          *string
}
type VirtualWireLldpHighAvailability struct {
	PassivePreNegotiation *bool
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName       xml.Name          `xml:"entry"`
	Name          string            `xml:"name,attr"`
	Comment       *string           `xml:"comment,omitempty"`
	DecryptMirror *DecryptMirrorXml `xml:"decrypt-mirror,omitempty"`
	Ha            *HaXml            `xml:"ha,omitempty"`
	Layer2        *Layer2Xml        `xml:"layer2,omitempty"`
	Layer3        *Layer3Xml        `xml:"layer3,omitempty"`
	VirtualWire   *VirtualWireXml   `xml:"virtual-wire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DecryptMirrorXml struct {
	Misc []generic.Xml `xml:",any"`
}
type HaXml struct {
	Lacp *HaLacpXml `xml:"lacp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type HaLacpXml struct {
	Enable           *string `xml:"enable,omitempty"`
	FastFailover     *string `xml:"fast-failover,omitempty"`
	MaxPorts         *int64  `xml:"max-ports,omitempty"`
	Mode             *string `xml:"mode,omitempty"`
	SystemPriority   *int64  `xml:"system-priority,omitempty"`
	TransmissionRate *string `xml:"transmission-rate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer2Xml struct {
	Lacp           *Layer2LacpXml `xml:"lacp,omitempty"`
	Lldp           *Layer2LldpXml `xml:"lldp,omitempty"`
	NetflowProfile *string        `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer2LacpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	FastFailover     *string                        `xml:"fast-failover,omitempty"`
	HighAvailability *Layer2LacpHighAvailabilityXml `xml:"high-availability,omitempty"`
	MaxPorts         *int64                         `xml:"max-ports,omitempty"`
	Mode             *string                        `xml:"mode,omitempty"`
	SystemPriority   *int64                         `xml:"system-priority,omitempty"`
	TransmissionRate *string                        `xml:"transmission-rate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer2LacpHighAvailabilityXml struct {
	PassivePreNegotiation *string `xml:"passive-pre-negotiation,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer2LldpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	HighAvailability *Layer2LldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                        `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer2LldpHighAvailabilityXml struct {
	PassivePreNegotiation *string `xml:"passive-pre-negotiation,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Xml struct {
	AdjustTcpMss               *Layer3AdjustTcpMssXml      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        []Layer3ArpXml              `xml:"arp>entry,omitempty"`
	Bonjour                    *Layer3BonjourXml           `xml:"bonjour,omitempty"`
	DdnsConfig                 *Layer3DdnsConfigXml        `xml:"ddns-config,omitempty"`
	DecryptForward             *string                     `xml:"decrypt-forward,omitempty"`
	DfIgnore                   *string                     `xml:"df-ignore,omitempty"`
	DhcpClient                 *Layer3DhcpClientXml        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                     `xml:"interface-management-profile,omitempty"`
	Ip                         []Layer3IpXml               `xml:"ip>entry,omitempty"`
	Ipv6                       *Layer3Ipv6Xml              `xml:"ipv6,omitempty"`
	Lacp                       *Layer3LacpXml              `xml:"lacp,omitempty"`
	Lldp                       *Layer3LldpXml              `xml:"lldp,omitempty"`
	Mtu                        *int64                      `xml:"mtu,omitempty"`
	NdpProxy                   *Layer3NdpProxyXml          `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string                     `xml:"netflow-profile,omitempty"`
	SdwanLinkSettings          *Layer3SdwanLinkSettingsXml `xml:"sdwan-link-settings,omitempty"`
	UntaggedSubInterface       *string                     `xml:"untagged-sub-interface,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3AdjustTcpMssXml struct {
	Enable            *string `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64  `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64  `xml:"ipv6-mss-adjustment,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3ArpXml struct {
	HwAddress *string  `xml:"hw-address,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3BonjourXml struct {
	Enable   *string `xml:"enable,omitempty"`
	GroupId  *int64  `xml:"group-id,omitempty"`
	TtlCheck *string `xml:"ttl-check,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3DdnsConfigXml struct {
	DdnsCertProfile    *string                               `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                               `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                               `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                      `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                      `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                               `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   []Layer3DdnsConfigDdnsVendorConfigXml `xml:"ddns-vendor-config>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3DdnsConfigDdnsVendorConfigXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   *string  `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3DhcpClientXml struct {
	CreateDefaultRoute *string                          `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                           `xml:"default-route-metric,omitempty"`
	Enable             *string                          `xml:"enable,omitempty"`
	SendHostname       *Layer3DhcpClientSendHostnameXml `xml:"send-hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3DhcpClientSendHostnameXml struct {
	Enable   *string `xml:"enable,omitempty"`
	Hostname *string `xml:"hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3IpXml struct {
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	SdwanGateway *string  `xml:"sdwan-gateway,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6Xml struct {
	Address           []Layer3Ipv6AddressXml          `xml:"address>entry,omitempty"`
	DhcpClient        *Layer3Ipv6DhcpClientXml        `xml:"dhcp-client,omitempty"`
	Enabled           *string                         `xml:"enabled,omitempty"`
	Inherited         *Layer3Ipv6InheritedXml         `xml:"inherited,omitempty"`
	InterfaceId       *string                         `xml:"interface-id,omitempty"`
	NeighborDiscovery *Layer3Ipv6NeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressXml struct {
	Advertise         *Layer3Ipv6AddressAdvertiseXml `xml:"advertise,omitempty"`
	Anycast           *Layer3Ipv6AddressAnycastXml   `xml:"anycast,omitempty"`
	EnableOnInterface *string                        `xml:"enable-on-interface,omitempty"`
	XMLName           xml.Name                       `xml:"entry"`
	Name              string                         `xml:"name,attr"`
	Prefix            *Layer3Ipv6AddressPrefixXml    `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressAdvertiseXml struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressAnycastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressPrefixXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientXml struct {
	AcceptRaRoute      *string                                   `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                                    `xml:"default-route-metric,omitempty"`
	Enable             *string                                   `xml:"enable,omitempty"`
	NeighborDiscovery  *Layer3Ipv6DhcpClientNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Preference         *string                                   `xml:"preference,omitempty"`
	PrefixDelegation   *Layer3Ipv6DhcpClientPrefixDelegationXml  `xml:"prefix-delegation,omitempty"`
	V6Options          *Layer3Ipv6DhcpClientV6OptionsXml         `xml:"v6-options,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryXml struct {
	DadAttempts      *int64                                             `xml:"dad-attempts,omitempty"`
	DnsServer        *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml `xml:"dns-server,omitempty"`
	DnsSuffix        *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                            `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                            `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         []Layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml `xml:"neighbor>entry,omitempty"`
	NsInterval       *int64                                             `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                             `xml:"reachable-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml struct {
	Server []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml struct {
	HwAddress *string  `xml:"hw-address,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientPrefixDelegationXml struct {
	Enable *Layer3Ipv6DhcpClientPrefixDelegationEnableXml `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableXml struct {
	No  *Layer3Ipv6DhcpClientPrefixDelegationEnableNoXml  `xml:"no,omitempty"`
	Yes *Layer3Ipv6DhcpClientPrefixDelegationEnableYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientPrefixDelegationEnableYesXml struct {
	PfxPoolName   *string `xml:"pfx-pool-name,omitempty"`
	PrefixLen     *int64  `xml:"prefix-len,omitempty"`
	PrefixLenHint *string `xml:"prefix-len-hint,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientV6OptionsXml struct {
	DuidType            *string                                 `xml:"duid-type,omitempty"`
	Enable              *Layer3Ipv6DhcpClientV6OptionsEnableXml `xml:"enable,omitempty"`
	RapidCommit         *string                                 `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                                 `xml:"support-srvr-reconfig,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientV6OptionsEnableXml struct {
	No  *Layer3Ipv6DhcpClientV6OptionsEnableNoXml  `xml:"no,omitempty"`
	Yes *Layer3Ipv6DhcpClientV6OptionsEnableYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientV6OptionsEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DhcpClientV6OptionsEnableYesXml struct {
	NonTempAddr *string `xml:"non-temp-addr,omitempty"`
	TempAddr    *string `xml:"temp-addr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedXml struct {
	AssignAddr        []Layer3Ipv6InheritedAssignAddrXml       `xml:"assign-addr>entry,omitempty"`
	Enable            *string                                  `xml:"enable,omitempty"`
	NeighborDiscovery *Layer3Ipv6InheritedNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrXml struct {
	XMLName xml.Name                              `xml:"entry"`
	Name    string                                `xml:"name,attr"`
	Type    *Layer3Ipv6InheritedAssignAddrTypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeXml struct {
	Gua *Layer3Ipv6InheritedAssignAddrTypeGuaXml `xml:"gua,omitempty"`
	Ula *Layer3Ipv6InheritedAssignAddrTypeUlaXml `xml:"ula,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeGuaXml struct {
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml `xml:"advertise,omitempty"`
	EnableOnInterface *string                                           `xml:"enable-on-interface,omitempty"`
	PoolType          *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml  `xml:"pool-type,omitempty"`
	PrefixPool        *string                                           `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml struct {
	AutoConfigFlag *string `xml:"auto-config-flag,omitempty"`
	Enable         *string `xml:"enable,omitempty"`
	OnlinkFlag     *string `xml:"onlink-flag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml struct {
	Dynamic   *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml   `xml:"dynamic,omitempty"`
	DynamicId *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml `xml:"dynamic-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml struct {
	Identifier *int64 `xml:"identifier,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeUlaXml struct {
	Address           *string                                           `xml:"address,omitempty"`
	Advertise         *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml `xml:"advertise,omitempty"`
	Anycast           *string                                           `xml:"anycast,omitempty"`
	EnableOnInterface *string                                           `xml:"enable-on-interface,omitempty"`
	Prefix            *string                                           `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryXml struct {
	DadAttempts         *int64                                                      `xml:"dad-attempts,omitempty"`
	DnsServer           *Layer3Ipv6InheritedNeighborDiscoveryDnsServerXml           `xml:"dns-server,omitempty"`
	DnsSuffix           *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                                     `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                     `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Layer3Ipv6InheritedNeighborDiscoveryNeighborXml           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                                      `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                      `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerXml struct {
	Enable *string                                                 `xml:"enable,omitempty"`
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml struct {
	Server []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                                 `xml:"enable,omitempty"`
	Source *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryNeighborXml struct {
	HwAddress *string  `xml:"hw-address,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml struct {
	Enable                 *string `xml:"enable,omitempty"`
	EnableConsistencyCheck *string `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string `xml:"hop-limit,omitempty"`
	Lifetime               *int64  `xml:"lifetime,omitempty"`
	LinkMtu                *string `xml:"link-mtu,omitempty"`
	ManagedFlag            *string `xml:"managed-flag,omitempty"`
	MaxInterval            *int64  `xml:"max-interval,omitempty"`
	MinInterval            *int64  `xml:"min-interval,omitempty"`
	OtherFlag              *string `xml:"other-flag,omitempty"`
	ReachableTime          *string `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string `xml:"router-preference,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryXml struct {
	DadAttempts         *int64                                             `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                            `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                            `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Layer3Ipv6NeighborDiscoveryNeighborXml           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                             `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                             `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Layer3Ipv6NeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryNeighborXml struct {
	HwAddress *string  `xml:"hw-address,omitempty"`
	XMLName   xml.Name `xml:"entry"`
	Name      string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementXml struct {
	DnsSupport             *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml `xml:"dns-support,omitempty"`
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

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml struct {
	Enable *string                                                             `xml:"enable,omitempty"`
	Server []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml `xml:"server>entry,omitempty"`
	Suffix []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3LacpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	FastFailover     *string                        `xml:"fast-failover,omitempty"`
	HighAvailability *Layer3LacpHighAvailabilityXml `xml:"high-availability,omitempty"`
	MaxPorts         *int64                         `xml:"max-ports,omitempty"`
	Mode             *string                        `xml:"mode,omitempty"`
	SystemPriority   *int64                         `xml:"system-priority,omitempty"`
	TransmissionRate *string                        `xml:"transmission-rate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3LacpHighAvailabilityXml struct {
	PassivePreNegotiation *string `xml:"passive-pre-negotiation,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3LldpXml struct {
	Enable           *string                        `xml:"enable,omitempty"`
	HighAvailability *Layer3LldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                        `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3LldpHighAvailabilityXml struct {
	PassivePreNegotiation *string `xml:"passive-pre-negotiation,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3NdpProxyXml struct {
	Address []Layer3NdpProxyAddressXml `xml:"address>entry,omitempty"`
	Enabled *string                    `xml:"enabled,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3NdpProxyAddressXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Negate  *string  `xml:"negate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsXml struct {
	Enable                *string                                `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                                `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *Layer3SdwanLinkSettingsUpstreamNatXml `xml:"upstream-nat,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsUpstreamNatXml struct {
	Enable   *string                                        `xml:"enable,omitempty"`
	Ddns     *Layer3SdwanLinkSettingsUpstreamNatDdnsXml     `xml:"ddns,omitempty"`
	StaticIp *Layer3SdwanLinkSettingsUpstreamNatStaticIpXml `xml:"static-ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsUpstreamNatDdnsXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsUpstreamNatStaticIpXml struct {
	Fqdn      *string `xml:"fqdn,omitempty"`
	IpAddress *string `xml:"ip-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VirtualWireXml struct {
	Lldp           *VirtualWireLldpXml `xml:"lldp,omitempty"`
	NetflowProfile *string             `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VirtualWireLldpXml struct {
	Enable           *string                             `xml:"enable,omitempty"`
	HighAvailability *VirtualWireLldpHighAvailabilityXml `xml:"high-availability,omitempty"`
	Profile          *string                             `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type VirtualWireLldpHighAvailabilityXml struct {
	PassivePreNegotiation *string `xml:"passive-pre-negotiation,omitempty"`

	Misc []generic.Xml `xml:",any"`
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

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Comment = o.Comment
	var nestedDecryptMirror *DecryptMirrorXml
	if o.DecryptMirror != nil {
		nestedDecryptMirror = &DecryptMirrorXml{}
		if _, ok := o.Misc["DecryptMirror"]; ok {
			nestedDecryptMirror.Misc = o.Misc["DecryptMirror"]
		}
	}
	entry.DecryptMirror = nestedDecryptMirror

	var nestedHa *HaXml
	if o.Ha != nil {
		nestedHa = &HaXml{}
		if _, ok := o.Misc["Ha"]; ok {
			nestedHa.Misc = o.Misc["Ha"]
		}
		if o.Ha.Lacp != nil {
			nestedHa.Lacp = &HaLacpXml{}
			if _, ok := o.Misc["HaLacp"]; ok {
				nestedHa.Lacp.Misc = o.Misc["HaLacp"]
			}
			if o.Ha.Lacp.Enable != nil {
				nestedHa.Lacp.Enable = util.YesNo(o.Ha.Lacp.Enable, nil)
			}
			if o.Ha.Lacp.FastFailover != nil {
				nestedHa.Lacp.FastFailover = util.YesNo(o.Ha.Lacp.FastFailover, nil)
			}
			if o.Ha.Lacp.MaxPorts != nil {
				nestedHa.Lacp.MaxPorts = o.Ha.Lacp.MaxPorts
			}
			if o.Ha.Lacp.Mode != nil {
				nestedHa.Lacp.Mode = o.Ha.Lacp.Mode
			}
			if o.Ha.Lacp.SystemPriority != nil {
				nestedHa.Lacp.SystemPriority = o.Ha.Lacp.SystemPriority
			}
			if o.Ha.Lacp.TransmissionRate != nil {
				nestedHa.Lacp.TransmissionRate = o.Ha.Lacp.TransmissionRate
			}
		}
	}
	entry.Ha = nestedHa

	var nestedLayer2 *Layer2Xml
	if o.Layer2 != nil {
		nestedLayer2 = &Layer2Xml{}
		if _, ok := o.Misc["Layer2"]; ok {
			nestedLayer2.Misc = o.Misc["Layer2"]
		}
		if o.Layer2.Lacp != nil {
			nestedLayer2.Lacp = &Layer2LacpXml{}
			if _, ok := o.Misc["Layer2Lacp"]; ok {
				nestedLayer2.Lacp.Misc = o.Misc["Layer2Lacp"]
			}
			if o.Layer2.Lacp.Mode != nil {
				nestedLayer2.Lacp.Mode = o.Layer2.Lacp.Mode
			}
			if o.Layer2.Lacp.SystemPriority != nil {
				nestedLayer2.Lacp.SystemPriority = o.Layer2.Lacp.SystemPriority
			}
			if o.Layer2.Lacp.TransmissionRate != nil {
				nestedLayer2.Lacp.TransmissionRate = o.Layer2.Lacp.TransmissionRate
			}
			if o.Layer2.Lacp.Enable != nil {
				nestedLayer2.Lacp.Enable = util.YesNo(o.Layer2.Lacp.Enable, nil)
			}
			if o.Layer2.Lacp.FastFailover != nil {
				nestedLayer2.Lacp.FastFailover = util.YesNo(o.Layer2.Lacp.FastFailover, nil)
			}
			if o.Layer2.Lacp.HighAvailability != nil {
				nestedLayer2.Lacp.HighAvailability = &Layer2LacpHighAvailabilityXml{}
				if _, ok := o.Misc["Layer2LacpHighAvailability"]; ok {
					nestedLayer2.Lacp.HighAvailability.Misc = o.Misc["Layer2LacpHighAvailability"]
				}
				if o.Layer2.Lacp.HighAvailability.PassivePreNegotiation != nil {
					nestedLayer2.Lacp.HighAvailability.PassivePreNegotiation = util.YesNo(o.Layer2.Lacp.HighAvailability.PassivePreNegotiation, nil)
				}
			}
			if o.Layer2.Lacp.MaxPorts != nil {
				nestedLayer2.Lacp.MaxPorts = o.Layer2.Lacp.MaxPorts
			}
		}
		if o.Layer2.Lldp != nil {
			nestedLayer2.Lldp = &Layer2LldpXml{}
			if _, ok := o.Misc["Layer2Lldp"]; ok {
				nestedLayer2.Lldp.Misc = o.Misc["Layer2Lldp"]
			}
			if o.Layer2.Lldp.Enable != nil {
				nestedLayer2.Lldp.Enable = util.YesNo(o.Layer2.Lldp.Enable, nil)
			}
			if o.Layer2.Lldp.HighAvailability != nil {
				nestedLayer2.Lldp.HighAvailability = &Layer2LldpHighAvailabilityXml{}
				if _, ok := o.Misc["Layer2LldpHighAvailability"]; ok {
					nestedLayer2.Lldp.HighAvailability.Misc = o.Misc["Layer2LldpHighAvailability"]
				}
				if o.Layer2.Lldp.HighAvailability.PassivePreNegotiation != nil {
					nestedLayer2.Lldp.HighAvailability.PassivePreNegotiation = util.YesNo(o.Layer2.Lldp.HighAvailability.PassivePreNegotiation, nil)
				}
			}
			if o.Layer2.Lldp.Profile != nil {
				nestedLayer2.Lldp.Profile = o.Layer2.Lldp.Profile
			}
		}
		if o.Layer2.NetflowProfile != nil {
			nestedLayer2.NetflowProfile = o.Layer2.NetflowProfile
		}
	}
	entry.Layer2 = nestedLayer2

	var nestedLayer3 *Layer3Xml
	if o.Layer3 != nil {
		nestedLayer3 = &Layer3Xml{}
		if _, ok := o.Misc["Layer3"]; ok {
			nestedLayer3.Misc = o.Misc["Layer3"]
		}
		if o.Layer3.UntaggedSubInterface != nil {
			nestedLayer3.UntaggedSubInterface = util.YesNo(o.Layer3.UntaggedSubInterface, nil)
		}
		if o.Layer3.DdnsConfig != nil {
			nestedLayer3.DdnsConfig = &Layer3DdnsConfigXml{}
			if _, ok := o.Misc["Layer3DdnsConfig"]; ok {
				nestedLayer3.DdnsConfig.Misc = o.Misc["Layer3DdnsConfig"]
			}
			if o.Layer3.DdnsConfig.DdnsIp != nil {
				nestedLayer3.DdnsConfig.DdnsIp = util.StrToMem(o.Layer3.DdnsConfig.DdnsIp)
			}
			if o.Layer3.DdnsConfig.DdnsIpv6 != nil {
				nestedLayer3.DdnsConfig.DdnsIpv6 = util.StrToMem(o.Layer3.DdnsConfig.DdnsIpv6)
			}
			if o.Layer3.DdnsConfig.DdnsUpdateInterval != nil {
				nestedLayer3.DdnsConfig.DdnsUpdateInterval = o.Layer3.DdnsConfig.DdnsUpdateInterval
			}
			if o.Layer3.DdnsConfig.DdnsVendor != nil {
				nestedLayer3.DdnsConfig.DdnsVendor = o.Layer3.DdnsConfig.DdnsVendor
			}
			if o.Layer3.DdnsConfig.DdnsVendorConfig != nil {
				nestedLayer3.DdnsConfig.DdnsVendorConfig = []Layer3DdnsConfigDdnsVendorConfigXml{}
				for _, oLayer3DdnsConfigDdnsVendorConfig := range o.Layer3.DdnsConfig.DdnsVendorConfig {
					nestedLayer3DdnsConfigDdnsVendorConfig := Layer3DdnsConfigDdnsVendorConfigXml{}
					if _, ok := o.Misc["Layer3DdnsConfigDdnsVendorConfig"]; ok {
						nestedLayer3DdnsConfigDdnsVendorConfig.Misc = o.Misc["Layer3DdnsConfigDdnsVendorConfig"]
					}
					if oLayer3DdnsConfigDdnsVendorConfig.Value != nil {
						nestedLayer3DdnsConfigDdnsVendorConfig.Value = oLayer3DdnsConfigDdnsVendorConfig.Value
					}
					if oLayer3DdnsConfigDdnsVendorConfig.Name != "" {
						nestedLayer3DdnsConfigDdnsVendorConfig.Name = oLayer3DdnsConfigDdnsVendorConfig.Name
					}
					nestedLayer3.DdnsConfig.DdnsVendorConfig = append(nestedLayer3.DdnsConfig.DdnsVendorConfig, nestedLayer3DdnsConfigDdnsVendorConfig)
				}
			}
			if o.Layer3.DdnsConfig.DdnsCertProfile != nil {
				nestedLayer3.DdnsConfig.DdnsCertProfile = o.Layer3.DdnsConfig.DdnsCertProfile
			}
			if o.Layer3.DdnsConfig.DdnsEnabled != nil {
				nestedLayer3.DdnsConfig.DdnsEnabled = util.YesNo(o.Layer3.DdnsConfig.DdnsEnabled, nil)
			}
			if o.Layer3.DdnsConfig.DdnsHostname != nil {
				nestedLayer3.DdnsConfig.DdnsHostname = o.Layer3.DdnsConfig.DdnsHostname
			}
		}
		if o.Layer3.DecryptForward != nil {
			nestedLayer3.DecryptForward = util.YesNo(o.Layer3.DecryptForward, nil)
		}
		if o.Layer3.DfIgnore != nil {
			nestedLayer3.DfIgnore = util.YesNo(o.Layer3.DfIgnore, nil)
		}
		if o.Layer3.DhcpClient != nil {
			nestedLayer3.DhcpClient = &Layer3DhcpClientXml{}
			if _, ok := o.Misc["Layer3DhcpClient"]; ok {
				nestedLayer3.DhcpClient.Misc = o.Misc["Layer3DhcpClient"]
			}
			if o.Layer3.DhcpClient.CreateDefaultRoute != nil {
				nestedLayer3.DhcpClient.CreateDefaultRoute = util.YesNo(o.Layer3.DhcpClient.CreateDefaultRoute, nil)
			}
			if o.Layer3.DhcpClient.DefaultRouteMetric != nil {
				nestedLayer3.DhcpClient.DefaultRouteMetric = o.Layer3.DhcpClient.DefaultRouteMetric
			}
			if o.Layer3.DhcpClient.Enable != nil {
				nestedLayer3.DhcpClient.Enable = util.YesNo(o.Layer3.DhcpClient.Enable, nil)
			}
			if o.Layer3.DhcpClient.SendHostname != nil {
				nestedLayer3.DhcpClient.SendHostname = &Layer3DhcpClientSendHostnameXml{}
				if _, ok := o.Misc["Layer3DhcpClientSendHostname"]; ok {
					nestedLayer3.DhcpClient.SendHostname.Misc = o.Misc["Layer3DhcpClientSendHostname"]
				}
				if o.Layer3.DhcpClient.SendHostname.Enable != nil {
					nestedLayer3.DhcpClient.SendHostname.Enable = util.YesNo(o.Layer3.DhcpClient.SendHostname.Enable, nil)
				}
				if o.Layer3.DhcpClient.SendHostname.Hostname != nil {
					nestedLayer3.DhcpClient.SendHostname.Hostname = o.Layer3.DhcpClient.SendHostname.Hostname
				}
			}
		}
		if o.Layer3.NetflowProfile != nil {
			nestedLayer3.NetflowProfile = o.Layer3.NetflowProfile
		}
		if o.Layer3.Bonjour != nil {
			nestedLayer3.Bonjour = &Layer3BonjourXml{}
			if _, ok := o.Misc["Layer3Bonjour"]; ok {
				nestedLayer3.Bonjour.Misc = o.Misc["Layer3Bonjour"]
			}
			if o.Layer3.Bonjour.Enable != nil {
				nestedLayer3.Bonjour.Enable = util.YesNo(o.Layer3.Bonjour.Enable, nil)
			}
			if o.Layer3.Bonjour.GroupId != nil {
				nestedLayer3.Bonjour.GroupId = o.Layer3.Bonjour.GroupId
			}
			if o.Layer3.Bonjour.TtlCheck != nil {
				nestedLayer3.Bonjour.TtlCheck = util.YesNo(o.Layer3.Bonjour.TtlCheck, nil)
			}
		}
		if o.Layer3.Ipv6 != nil {
			nestedLayer3.Ipv6 = &Layer3Ipv6Xml{}
			if _, ok := o.Misc["Layer3Ipv6"]; ok {
				nestedLayer3.Ipv6.Misc = o.Misc["Layer3Ipv6"]
			}
			if o.Layer3.Ipv6.Address != nil {
				nestedLayer3.Ipv6.Address = []Layer3Ipv6AddressXml{}
				for _, oLayer3Ipv6Address := range o.Layer3.Ipv6.Address {
					nestedLayer3Ipv6Address := Layer3Ipv6AddressXml{}
					if _, ok := o.Misc["Layer3Ipv6Address"]; ok {
						nestedLayer3Ipv6Address.Misc = o.Misc["Layer3Ipv6Address"]
					}
					if oLayer3Ipv6Address.EnableOnInterface != nil {
						nestedLayer3Ipv6Address.EnableOnInterface = util.YesNo(oLayer3Ipv6Address.EnableOnInterface, nil)
					}
					if oLayer3Ipv6Address.Prefix != nil {
						nestedLayer3Ipv6Address.Prefix = &Layer3Ipv6AddressPrefixXml{}
						if _, ok := o.Misc["Layer3Ipv6AddressPrefix"]; ok {
							nestedLayer3Ipv6Address.Prefix.Misc = o.Misc["Layer3Ipv6AddressPrefix"]
						}
					}
					if oLayer3Ipv6Address.Anycast != nil {
						nestedLayer3Ipv6Address.Anycast = &Layer3Ipv6AddressAnycastXml{}
						if _, ok := o.Misc["Layer3Ipv6AddressAnycast"]; ok {
							nestedLayer3Ipv6Address.Anycast.Misc = o.Misc["Layer3Ipv6AddressAnycast"]
						}
					}
					if oLayer3Ipv6Address.Advertise != nil {
						nestedLayer3Ipv6Address.Advertise = &Layer3Ipv6AddressAdvertiseXml{}
						if _, ok := o.Misc["Layer3Ipv6AddressAdvertise"]; ok {
							nestedLayer3Ipv6Address.Advertise.Misc = o.Misc["Layer3Ipv6AddressAdvertise"]
						}
						if oLayer3Ipv6Address.Advertise.Enable != nil {
							nestedLayer3Ipv6Address.Advertise.Enable = util.YesNo(oLayer3Ipv6Address.Advertise.Enable, nil)
						}
						if oLayer3Ipv6Address.Advertise.ValidLifetime != nil {
							nestedLayer3Ipv6Address.Advertise.ValidLifetime = oLayer3Ipv6Address.Advertise.ValidLifetime
						}
						if oLayer3Ipv6Address.Advertise.PreferredLifetime != nil {
							nestedLayer3Ipv6Address.Advertise.PreferredLifetime = oLayer3Ipv6Address.Advertise.PreferredLifetime
						}
						if oLayer3Ipv6Address.Advertise.OnlinkFlag != nil {
							nestedLayer3Ipv6Address.Advertise.OnlinkFlag = util.YesNo(oLayer3Ipv6Address.Advertise.OnlinkFlag, nil)
						}
						if oLayer3Ipv6Address.Advertise.AutoConfigFlag != nil {
							nestedLayer3Ipv6Address.Advertise.AutoConfigFlag = util.YesNo(oLayer3Ipv6Address.Advertise.AutoConfigFlag, nil)
						}
					}
					if oLayer3Ipv6Address.Name != "" {
						nestedLayer3Ipv6Address.Name = oLayer3Ipv6Address.Name
					}
					nestedLayer3.Ipv6.Address = append(nestedLayer3.Ipv6.Address, nestedLayer3Ipv6Address)
				}
			}
			if o.Layer3.Ipv6.Enabled != nil {
				nestedLayer3.Ipv6.Enabled = util.YesNo(o.Layer3.Ipv6.Enabled, nil)
			}
			if o.Layer3.Ipv6.InterfaceId != nil {
				nestedLayer3.Ipv6.InterfaceId = o.Layer3.Ipv6.InterfaceId
			}
			if o.Layer3.Ipv6.NeighborDiscovery != nil {
				nestedLayer3.Ipv6.NeighborDiscovery = &Layer3Ipv6NeighborDiscoveryXml{}
				if _, ok := o.Misc["Layer3Ipv6NeighborDiscovery"]; ok {
					nestedLayer3.Ipv6.NeighborDiscovery.Misc = o.Misc["Layer3Ipv6NeighborDiscovery"]
				}
				if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement = &Layer3Ipv6NeighborDiscoveryRouterAdvertisementXml{}
					if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisement"]; ok {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisement"]
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml{}
						if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]; ok {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml{}
							for _, oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
								nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml{}
								if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]; ok {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]
								}
								if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
								}
								if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
								}
								nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
							}
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml{}
							for _, oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
								nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml{}
								if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]; ok {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]
								}
								if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
								}
								if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
								}
								nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
							}
						}
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
				}
				if o.Layer3.Ipv6.NeighborDiscovery.DadAttempts != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.NeighborDiscovery.DadAttempts
				}
				if o.Layer3.Ipv6.NeighborDiscovery.EnableDad != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.EnableDad = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.EnableDad, nil)
				}
				if o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Layer3.Ipv6.NeighborDiscovery.Neighbor != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.Neighbor = []Layer3Ipv6NeighborDiscoveryNeighborXml{}
					for _, oLayer3Ipv6NeighborDiscoveryNeighbor := range o.Layer3.Ipv6.NeighborDiscovery.Neighbor {
						nestedLayer3Ipv6NeighborDiscoveryNeighbor := Layer3Ipv6NeighborDiscoveryNeighborXml{}
						if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryNeighbor"]; ok {
							nestedLayer3Ipv6NeighborDiscoveryNeighbor.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryNeighbor"]
						}
						if oLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress != nil {
							nestedLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress
						}
						if oLayer3Ipv6NeighborDiscoveryNeighbor.Name != "" {
							nestedLayer3Ipv6NeighborDiscoveryNeighbor.Name = oLayer3Ipv6NeighborDiscoveryNeighbor.Name
						}
						nestedLayer3.Ipv6.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.NeighborDiscovery.Neighbor, nestedLayer3Ipv6NeighborDiscoveryNeighbor)
					}
				}
				if o.Layer3.Ipv6.NeighborDiscovery.NsInterval != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.NeighborDiscovery.NsInterval
				}
				if o.Layer3.Ipv6.NeighborDiscovery.ReachableTime != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.NeighborDiscovery.ReachableTime
				}
			}
			if o.Layer3.Ipv6.DhcpClient != nil {
				nestedLayer3.Ipv6.DhcpClient = &Layer3Ipv6DhcpClientXml{}
				if _, ok := o.Misc["Layer3Ipv6DhcpClient"]; ok {
					nestedLayer3.Ipv6.DhcpClient.Misc = o.Misc["Layer3Ipv6DhcpClient"]
				}
				if o.Layer3.Ipv6.DhcpClient.DefaultRouteMetric != nil {
					nestedLayer3.Ipv6.DhcpClient.DefaultRouteMetric = o.Layer3.Ipv6.DhcpClient.DefaultRouteMetric
				}
				if o.Layer3.Ipv6.DhcpClient.Enable != nil {
					nestedLayer3.Ipv6.DhcpClient.Enable = util.YesNo(o.Layer3.Ipv6.DhcpClient.Enable, nil)
				}
				if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery != nil {
					nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery = &Layer3Ipv6DhcpClientNeighborDiscoveryXml{}
					if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscovery"]; ok {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscovery"]
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor = []Layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml{}
						for _, oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
							nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor := Layer3Ipv6DhcpClientNeighborDiscoveryNeighborXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor"]; ok {
								nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor"]
							}
							if oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
							}
							if oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
								nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name
							}
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor, nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor)
						}
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerXml{}
						if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer"]; ok {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer"]
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]; ok {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml{}
								if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]
								}
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml{}
								if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]; ok {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml{}
									for _, oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml{}
										if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]; ok {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]
										}
										if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml{}
						if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]; ok {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]; ok {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml{}
								if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]
								}
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml{}
								if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]; ok {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
									for _, oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
										if _, ok := o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]
										}
										if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad = util.YesNo(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
					}
				}
				if o.Layer3.Ipv6.DhcpClient.Preference != nil {
					nestedLayer3.Ipv6.DhcpClient.Preference = o.Layer3.Ipv6.DhcpClient.Preference
				}
				if o.Layer3.Ipv6.DhcpClient.PrefixDelegation != nil {
					nestedLayer3.Ipv6.DhcpClient.PrefixDelegation = &Layer3Ipv6DhcpClientPrefixDelegationXml{}
					if _, ok := o.Misc["Layer3Ipv6DhcpClientPrefixDelegation"]; ok {
						nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Misc = o.Misc["Layer3Ipv6DhcpClientPrefixDelegation"]
					}
					if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
						nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable = &Layer3Ipv6DhcpClientPrefixDelegationEnableXml{}
						if _, ok := o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnable"]; ok {
							nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc = o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnable"]
						}
						if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
							nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No = &Layer3Ipv6DhcpClientPrefixDelegationEnableNoXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableNo"]; ok {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc = o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableNo"]
							}
						}
						if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
							nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes = &Layer3Ipv6DhcpClientPrefixDelegationEnableYesXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableYes"]; ok {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc = o.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableYes"]
							}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
							}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
							}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.YesNo(o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
							}
						}
					}
				}
				if o.Layer3.Ipv6.DhcpClient.V6Options != nil {
					nestedLayer3.Ipv6.DhcpClient.V6Options = &Layer3Ipv6DhcpClientV6OptionsXml{}
					if _, ok := o.Misc["Layer3Ipv6DhcpClientV6Options"]; ok {
						nestedLayer3.Ipv6.DhcpClient.V6Options.Misc = o.Misc["Layer3Ipv6DhcpClientV6Options"]
					}
					if o.Layer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
						nestedLayer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig = util.YesNo(o.Layer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
					}
					if o.Layer3.Ipv6.DhcpClient.V6Options.DuidType != nil {
						nestedLayer3.Ipv6.DhcpClient.V6Options.DuidType = o.Layer3.Ipv6.DhcpClient.V6Options.DuidType
					}
					if o.Layer3.Ipv6.DhcpClient.V6Options.Enable != nil {
						nestedLayer3.Ipv6.DhcpClient.V6Options.Enable = &Layer3Ipv6DhcpClientV6OptionsEnableXml{}
						if _, ok := o.Misc["Layer3Ipv6DhcpClientV6OptionsEnable"]; ok {
							nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Misc = o.Misc["Layer3Ipv6DhcpClientV6OptionsEnable"]
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.No != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.No = &Layer3Ipv6DhcpClientV6OptionsEnableNoXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientV6OptionsEnableNo"]; ok {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.No.Misc = o.Misc["Layer3Ipv6DhcpClientV6OptionsEnableNo"]
							}
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes = &Layer3Ipv6DhcpClientV6OptionsEnableYesXml{}
							if _, ok := o.Misc["Layer3Ipv6DhcpClientV6OptionsEnableYes"]; ok {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc = o.Misc["Layer3Ipv6DhcpClientV6OptionsEnableYes"]
							}
							if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.YesNo(o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
							}
							if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.YesNo(o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
							}
						}
					}
					if o.Layer3.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
						nestedLayer3.Ipv6.DhcpClient.V6Options.RapidCommit = util.YesNo(o.Layer3.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
					}
				}
				if o.Layer3.Ipv6.DhcpClient.AcceptRaRoute != nil {
					nestedLayer3.Ipv6.DhcpClient.AcceptRaRoute = util.YesNo(o.Layer3.Ipv6.DhcpClient.AcceptRaRoute, nil)
				}
			}
			if o.Layer3.Ipv6.Inherited != nil {
				nestedLayer3.Ipv6.Inherited = &Layer3Ipv6InheritedXml{}
				if _, ok := o.Misc["Layer3Ipv6Inherited"]; ok {
					nestedLayer3.Ipv6.Inherited.Misc = o.Misc["Layer3Ipv6Inherited"]
				}
				if o.Layer3.Ipv6.Inherited.NeighborDiscovery != nil {
					nestedLayer3.Ipv6.Inherited.NeighborDiscovery = &Layer3Ipv6InheritedNeighborDiscoveryXml{}
					if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscovery"]; ok {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscovery"]
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.NsInterval
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor = []Layer3Ipv6InheritedNeighborDiscoveryNeighborXml{}
						for _, oLayer3Ipv6InheritedNeighborDiscoveryNeighbor := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.Neighbor {
							nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor := Layer3Ipv6InheritedNeighborDiscoveryNeighborXml{}
							if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryNeighbor"]; ok {
								nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryNeighbor"]
							}
							if oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress
							}
							if oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name != "" {
								nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name = oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name
							}
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor, nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor)
						}
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerXml{}
						if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServer"]; ok {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServer"]
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceXml{}
							if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource"]; ok {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource"]
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]; ok {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml{}
									for _, oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml{}
										if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]; ok {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]
										}
										if oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										if oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixXml{}
						if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix"]; ok {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix"]
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml{}
							if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]; ok {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]; ok {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
									for _, oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
										if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]
										}
										if oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										if oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.EnableDad = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml{}
						if _, ok := o.Misc["Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]; ok {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
						}
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts
					}
				}
				if o.Layer3.Ipv6.Inherited.AssignAddr != nil {
					nestedLayer3.Ipv6.Inherited.AssignAddr = []Layer3Ipv6InheritedAssignAddrXml{}
					for _, oLayer3Ipv6InheritedAssignAddr := range o.Layer3.Ipv6.Inherited.AssignAddr {
						nestedLayer3Ipv6InheritedAssignAddr := Layer3Ipv6InheritedAssignAddrXml{}
						if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddr"]; ok {
							nestedLayer3Ipv6InheritedAssignAddr.Misc = o.Misc["Layer3Ipv6InheritedAssignAddr"]
						}
						if oLayer3Ipv6InheritedAssignAddr.Type != nil {
							nestedLayer3Ipv6InheritedAssignAddr.Type = &Layer3Ipv6InheritedAssignAddrTypeXml{}
							if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrType"]; ok {
								nestedLayer3Ipv6InheritedAssignAddr.Type.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrType"]
							}
							if oLayer3Ipv6InheritedAssignAddr.Type.Gua != nil {
								nestedLayer3Ipv6InheritedAssignAddr.Type.Gua = &Layer3Ipv6InheritedAssignAddrTypeGuaXml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeGua"]; ok {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeGua"]
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeXml{}
									if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolType"]; ok {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolType"]
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml{}
										if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]; ok {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]
										}
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml{}
										if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]; ok {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
										}
									}
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise = &Layer3Ipv6InheritedAssignAddrTypeGuaAdvertiseXml{}
									if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise"]; ok {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise"]
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
									}
								}
							}
							if oLayer3Ipv6InheritedAssignAddr.Type.Ula != nil {
								nestedLayer3Ipv6InheritedAssignAddr.Type.Ula = &Layer3Ipv6InheritedAssignAddrTypeUlaXml{}
								if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeUla"]; ok {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeUla"]
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast, nil)
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise = &Layer3Ipv6InheritedAssignAddrTypeUlaAdvertiseXml{}
									if _, ok := o.Misc["Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise"]; ok {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Misc = o.Misc["Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise"]
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
									}
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Address != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Address = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Address
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix = util.YesNo(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix, nil)
								}
							}
						}
						if oLayer3Ipv6InheritedAssignAddr.Name != "" {
							nestedLayer3Ipv6InheritedAssignAddr.Name = oLayer3Ipv6InheritedAssignAddr.Name
						}
						nestedLayer3.Ipv6.Inherited.AssignAddr = append(nestedLayer3.Ipv6.Inherited.AssignAddr, nestedLayer3Ipv6InheritedAssignAddr)
					}
				}
				if o.Layer3.Ipv6.Inherited.Enable != nil {
					nestedLayer3.Ipv6.Inherited.Enable = util.YesNo(o.Layer3.Ipv6.Inherited.Enable, nil)
				}
			}
		}
		if o.Layer3.Lldp != nil {
			nestedLayer3.Lldp = &Layer3LldpXml{}
			if _, ok := o.Misc["Layer3Lldp"]; ok {
				nestedLayer3.Lldp.Misc = o.Misc["Layer3Lldp"]
			}
			if o.Layer3.Lldp.Enable != nil {
				nestedLayer3.Lldp.Enable = util.YesNo(o.Layer3.Lldp.Enable, nil)
			}
			if o.Layer3.Lldp.HighAvailability != nil {
				nestedLayer3.Lldp.HighAvailability = &Layer3LldpHighAvailabilityXml{}
				if _, ok := o.Misc["Layer3LldpHighAvailability"]; ok {
					nestedLayer3.Lldp.HighAvailability.Misc = o.Misc["Layer3LldpHighAvailability"]
				}
				if o.Layer3.Lldp.HighAvailability.PassivePreNegotiation != nil {
					nestedLayer3.Lldp.HighAvailability.PassivePreNegotiation = util.YesNo(o.Layer3.Lldp.HighAvailability.PassivePreNegotiation, nil)
				}
			}
			if o.Layer3.Lldp.Profile != nil {
				nestedLayer3.Lldp.Profile = o.Layer3.Lldp.Profile
			}
		}
		if o.Layer3.Arp != nil {
			nestedLayer3.Arp = []Layer3ArpXml{}
			for _, oLayer3Arp := range o.Layer3.Arp {
				nestedLayer3Arp := Layer3ArpXml{}
				if _, ok := o.Misc["Layer3Arp"]; ok {
					nestedLayer3Arp.Misc = o.Misc["Layer3Arp"]
				}
				if oLayer3Arp.HwAddress != nil {
					nestedLayer3Arp.HwAddress = oLayer3Arp.HwAddress
				}
				if oLayer3Arp.Name != "" {
					nestedLayer3Arp.Name = oLayer3Arp.Name
				}
				nestedLayer3.Arp = append(nestedLayer3.Arp, nestedLayer3Arp)
			}
		}
		if o.Layer3.InterfaceManagementProfile != nil {
			nestedLayer3.InterfaceManagementProfile = o.Layer3.InterfaceManagementProfile
		}
		if o.Layer3.Ip != nil {
			nestedLayer3.Ip = []Layer3IpXml{}
			for _, oLayer3Ip := range o.Layer3.Ip {
				nestedLayer3Ip := Layer3IpXml{}
				if _, ok := o.Misc["Layer3Ip"]; ok {
					nestedLayer3Ip.Misc = o.Misc["Layer3Ip"]
				}
				if oLayer3Ip.SdwanGateway != nil {
					nestedLayer3Ip.SdwanGateway = oLayer3Ip.SdwanGateway
				}
				if oLayer3Ip.Name != "" {
					nestedLayer3Ip.Name = oLayer3Ip.Name
				}
				nestedLayer3.Ip = append(nestedLayer3.Ip, nestedLayer3Ip)
			}
		}
		if o.Layer3.Lacp != nil {
			nestedLayer3.Lacp = &Layer3LacpXml{}
			if _, ok := o.Misc["Layer3Lacp"]; ok {
				nestedLayer3.Lacp.Misc = o.Misc["Layer3Lacp"]
			}
			if o.Layer3.Lacp.MaxPorts != nil {
				nestedLayer3.Lacp.MaxPorts = o.Layer3.Lacp.MaxPorts
			}
			if o.Layer3.Lacp.Mode != nil {
				nestedLayer3.Lacp.Mode = o.Layer3.Lacp.Mode
			}
			if o.Layer3.Lacp.SystemPriority != nil {
				nestedLayer3.Lacp.SystemPriority = o.Layer3.Lacp.SystemPriority
			}
			if o.Layer3.Lacp.TransmissionRate != nil {
				nestedLayer3.Lacp.TransmissionRate = o.Layer3.Lacp.TransmissionRate
			}
			if o.Layer3.Lacp.Enable != nil {
				nestedLayer3.Lacp.Enable = util.YesNo(o.Layer3.Lacp.Enable, nil)
			}
			if o.Layer3.Lacp.FastFailover != nil {
				nestedLayer3.Lacp.FastFailover = util.YesNo(o.Layer3.Lacp.FastFailover, nil)
			}
			if o.Layer3.Lacp.HighAvailability != nil {
				nestedLayer3.Lacp.HighAvailability = &Layer3LacpHighAvailabilityXml{}
				if _, ok := o.Misc["Layer3LacpHighAvailability"]; ok {
					nestedLayer3.Lacp.HighAvailability.Misc = o.Misc["Layer3LacpHighAvailability"]
				}
				if o.Layer3.Lacp.HighAvailability.PassivePreNegotiation != nil {
					nestedLayer3.Lacp.HighAvailability.PassivePreNegotiation = util.YesNo(o.Layer3.Lacp.HighAvailability.PassivePreNegotiation, nil)
				}
			}
		}
		if o.Layer3.Mtu != nil {
			nestedLayer3.Mtu = o.Layer3.Mtu
		}
		if o.Layer3.AdjustTcpMss != nil {
			nestedLayer3.AdjustTcpMss = &Layer3AdjustTcpMssXml{}
			if _, ok := o.Misc["Layer3AdjustTcpMss"]; ok {
				nestedLayer3.AdjustTcpMss.Misc = o.Misc["Layer3AdjustTcpMss"]
			}
			if o.Layer3.AdjustTcpMss.Enable != nil {
				nestedLayer3.AdjustTcpMss.Enable = util.YesNo(o.Layer3.AdjustTcpMss.Enable, nil)
			}
			if o.Layer3.AdjustTcpMss.Ipv4MssAdjustment != nil {
				nestedLayer3.AdjustTcpMss.Ipv4MssAdjustment = o.Layer3.AdjustTcpMss.Ipv4MssAdjustment
			}
			if o.Layer3.AdjustTcpMss.Ipv6MssAdjustment != nil {
				nestedLayer3.AdjustTcpMss.Ipv6MssAdjustment = o.Layer3.AdjustTcpMss.Ipv6MssAdjustment
			}
		}
		if o.Layer3.NdpProxy != nil {
			nestedLayer3.NdpProxy = &Layer3NdpProxyXml{}
			if _, ok := o.Misc["Layer3NdpProxy"]; ok {
				nestedLayer3.NdpProxy.Misc = o.Misc["Layer3NdpProxy"]
			}
			if o.Layer3.NdpProxy.Address != nil {
				nestedLayer3.NdpProxy.Address = []Layer3NdpProxyAddressXml{}
				for _, oLayer3NdpProxyAddress := range o.Layer3.NdpProxy.Address {
					nestedLayer3NdpProxyAddress := Layer3NdpProxyAddressXml{}
					if _, ok := o.Misc["Layer3NdpProxyAddress"]; ok {
						nestedLayer3NdpProxyAddress.Misc = o.Misc["Layer3NdpProxyAddress"]
					}
					if oLayer3NdpProxyAddress.Negate != nil {
						nestedLayer3NdpProxyAddress.Negate = util.YesNo(oLayer3NdpProxyAddress.Negate, nil)
					}
					if oLayer3NdpProxyAddress.Name != "" {
						nestedLayer3NdpProxyAddress.Name = oLayer3NdpProxyAddress.Name
					}
					nestedLayer3.NdpProxy.Address = append(nestedLayer3.NdpProxy.Address, nestedLayer3NdpProxyAddress)
				}
			}
			if o.Layer3.NdpProxy.Enabled != nil {
				nestedLayer3.NdpProxy.Enabled = util.YesNo(o.Layer3.NdpProxy.Enabled, nil)
			}
		}
		if o.Layer3.SdwanLinkSettings != nil {
			nestedLayer3.SdwanLinkSettings = &Layer3SdwanLinkSettingsXml{}
			if _, ok := o.Misc["Layer3SdwanLinkSettings"]; ok {
				nestedLayer3.SdwanLinkSettings.Misc = o.Misc["Layer3SdwanLinkSettings"]
			}
			if o.Layer3.SdwanLinkSettings.Enable != nil {
				nestedLayer3.SdwanLinkSettings.Enable = util.YesNo(o.Layer3.SdwanLinkSettings.Enable, nil)
			}
			if o.Layer3.SdwanLinkSettings.SdwanInterfaceProfile != nil {
				nestedLayer3.SdwanLinkSettings.SdwanInterfaceProfile = o.Layer3.SdwanLinkSettings.SdwanInterfaceProfile
			}
			if o.Layer3.SdwanLinkSettings.UpstreamNat != nil {
				nestedLayer3.SdwanLinkSettings.UpstreamNat = &Layer3SdwanLinkSettingsUpstreamNatXml{}
				if _, ok := o.Misc["Layer3SdwanLinkSettingsUpstreamNat"]; ok {
					nestedLayer3.SdwanLinkSettings.UpstreamNat.Misc = o.Misc["Layer3SdwanLinkSettingsUpstreamNat"]
				}
				if o.Layer3.SdwanLinkSettings.UpstreamNat.Enable != nil {
					nestedLayer3.SdwanLinkSettings.UpstreamNat.Enable = util.YesNo(o.Layer3.SdwanLinkSettings.UpstreamNat.Enable, nil)
				}
				if o.Layer3.SdwanLinkSettings.UpstreamNat.Ddns != nil {
					nestedLayer3.SdwanLinkSettings.UpstreamNat.Ddns = &Layer3SdwanLinkSettingsUpstreamNatDdnsXml{}
					if _, ok := o.Misc["Layer3SdwanLinkSettingsUpstreamNatDdns"]; ok {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.Ddns.Misc = o.Misc["Layer3SdwanLinkSettingsUpstreamNatDdns"]
					}
				}
				if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp != nil {
					nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp = &Layer3SdwanLinkSettingsUpstreamNatStaticIpXml{}
					if _, ok := o.Misc["Layer3SdwanLinkSettingsUpstreamNatStaticIp"]; ok {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp.Misc = o.Misc["Layer3SdwanLinkSettingsUpstreamNatStaticIp"]
					}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn
					}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress
					}
				}
			}
		}
	}
	entry.Layer3 = nestedLayer3

	var nestedVirtualWire *VirtualWireXml
	if o.VirtualWire != nil {
		nestedVirtualWire = &VirtualWireXml{}
		if _, ok := o.Misc["VirtualWire"]; ok {
			nestedVirtualWire.Misc = o.Misc["VirtualWire"]
		}
		if o.VirtualWire.Lldp != nil {
			nestedVirtualWire.Lldp = &VirtualWireLldpXml{}
			if _, ok := o.Misc["VirtualWireLldp"]; ok {
				nestedVirtualWire.Lldp.Misc = o.Misc["VirtualWireLldp"]
			}
			if o.VirtualWire.Lldp.Enable != nil {
				nestedVirtualWire.Lldp.Enable = util.YesNo(o.VirtualWire.Lldp.Enable, nil)
			}
			if o.VirtualWire.Lldp.HighAvailability != nil {
				nestedVirtualWire.Lldp.HighAvailability = &VirtualWireLldpHighAvailabilityXml{}
				if _, ok := o.Misc["VirtualWireLldpHighAvailability"]; ok {
					nestedVirtualWire.Lldp.HighAvailability.Misc = o.Misc["VirtualWireLldpHighAvailability"]
				}
				if o.VirtualWire.Lldp.HighAvailability.PassivePreNegotiation != nil {
					nestedVirtualWire.Lldp.HighAvailability.PassivePreNegotiation = util.YesNo(o.VirtualWire.Lldp.HighAvailability.PassivePreNegotiation, nil)
				}
			}
			if o.VirtualWire.Lldp.Profile != nil {
				nestedVirtualWire.Lldp.Profile = o.VirtualWire.Lldp.Profile
			}
		}
		if o.VirtualWire.NetflowProfile != nil {
			nestedVirtualWire.NetflowProfile = o.VirtualWire.NetflowProfile
		}
	}
	entry.VirtualWire = nestedVirtualWire

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
		entry.Comment = o.Comment
		var nestedDecryptMirror *DecryptMirror
		if o.DecryptMirror != nil {
			nestedDecryptMirror = &DecryptMirror{}
			if o.DecryptMirror.Misc != nil {
				entry.Misc["DecryptMirror"] = o.DecryptMirror.Misc
			}
		}
		entry.DecryptMirror = nestedDecryptMirror

		var nestedHa *Ha
		if o.Ha != nil {
			nestedHa = &Ha{}
			if o.Ha.Misc != nil {
				entry.Misc["Ha"] = o.Ha.Misc
			}
			if o.Ha.Lacp != nil {
				nestedHa.Lacp = &HaLacp{}
				if o.Ha.Lacp.Misc != nil {
					entry.Misc["HaLacp"] = o.Ha.Lacp.Misc
				}
				if o.Ha.Lacp.TransmissionRate != nil {
					nestedHa.Lacp.TransmissionRate = o.Ha.Lacp.TransmissionRate
				}
				if o.Ha.Lacp.Enable != nil {
					nestedHa.Lacp.Enable = util.AsBool(o.Ha.Lacp.Enable, nil)
				}
				if o.Ha.Lacp.FastFailover != nil {
					nestedHa.Lacp.FastFailover = util.AsBool(o.Ha.Lacp.FastFailover, nil)
				}
				if o.Ha.Lacp.MaxPorts != nil {
					nestedHa.Lacp.MaxPorts = o.Ha.Lacp.MaxPorts
				}
				if o.Ha.Lacp.Mode != nil {
					nestedHa.Lacp.Mode = o.Ha.Lacp.Mode
				}
				if o.Ha.Lacp.SystemPriority != nil {
					nestedHa.Lacp.SystemPriority = o.Ha.Lacp.SystemPriority
				}
			}
		}
		entry.Ha = nestedHa

		var nestedLayer2 *Layer2
		if o.Layer2 != nil {
			nestedLayer2 = &Layer2{}
			if o.Layer2.Misc != nil {
				entry.Misc["Layer2"] = o.Layer2.Misc
			}
			if o.Layer2.Lldp != nil {
				nestedLayer2.Lldp = &Layer2Lldp{}
				if o.Layer2.Lldp.Misc != nil {
					entry.Misc["Layer2Lldp"] = o.Layer2.Lldp.Misc
				}
				if o.Layer2.Lldp.Enable != nil {
					nestedLayer2.Lldp.Enable = util.AsBool(o.Layer2.Lldp.Enable, nil)
				}
				if o.Layer2.Lldp.HighAvailability != nil {
					nestedLayer2.Lldp.HighAvailability = &Layer2LldpHighAvailability{}
					if o.Layer2.Lldp.HighAvailability.Misc != nil {
						entry.Misc["Layer2LldpHighAvailability"] = o.Layer2.Lldp.HighAvailability.Misc
					}
					if o.Layer2.Lldp.HighAvailability.PassivePreNegotiation != nil {
						nestedLayer2.Lldp.HighAvailability.PassivePreNegotiation = util.AsBool(o.Layer2.Lldp.HighAvailability.PassivePreNegotiation, nil)
					}
				}
				if o.Layer2.Lldp.Profile != nil {
					nestedLayer2.Lldp.Profile = o.Layer2.Lldp.Profile
				}
			}
			if o.Layer2.NetflowProfile != nil {
				nestedLayer2.NetflowProfile = o.Layer2.NetflowProfile
			}
			if o.Layer2.Lacp != nil {
				nestedLayer2.Lacp = &Layer2Lacp{}
				if o.Layer2.Lacp.Misc != nil {
					entry.Misc["Layer2Lacp"] = o.Layer2.Lacp.Misc
				}
				if o.Layer2.Lacp.MaxPorts != nil {
					nestedLayer2.Lacp.MaxPorts = o.Layer2.Lacp.MaxPorts
				}
				if o.Layer2.Lacp.Mode != nil {
					nestedLayer2.Lacp.Mode = o.Layer2.Lacp.Mode
				}
				if o.Layer2.Lacp.SystemPriority != nil {
					nestedLayer2.Lacp.SystemPriority = o.Layer2.Lacp.SystemPriority
				}
				if o.Layer2.Lacp.TransmissionRate != nil {
					nestedLayer2.Lacp.TransmissionRate = o.Layer2.Lacp.TransmissionRate
				}
				if o.Layer2.Lacp.Enable != nil {
					nestedLayer2.Lacp.Enable = util.AsBool(o.Layer2.Lacp.Enable, nil)
				}
				if o.Layer2.Lacp.FastFailover != nil {
					nestedLayer2.Lacp.FastFailover = util.AsBool(o.Layer2.Lacp.FastFailover, nil)
				}
				if o.Layer2.Lacp.HighAvailability != nil {
					nestedLayer2.Lacp.HighAvailability = &Layer2LacpHighAvailability{}
					if o.Layer2.Lacp.HighAvailability.Misc != nil {
						entry.Misc["Layer2LacpHighAvailability"] = o.Layer2.Lacp.HighAvailability.Misc
					}
					if o.Layer2.Lacp.HighAvailability.PassivePreNegotiation != nil {
						nestedLayer2.Lacp.HighAvailability.PassivePreNegotiation = util.AsBool(o.Layer2.Lacp.HighAvailability.PassivePreNegotiation, nil)
					}
				}
			}
		}
		entry.Layer2 = nestedLayer2

		var nestedLayer3 *Layer3
		if o.Layer3 != nil {
			nestedLayer3 = &Layer3{}
			if o.Layer3.Misc != nil {
				entry.Misc["Layer3"] = o.Layer3.Misc
			}
			if o.Layer3.AdjustTcpMss != nil {
				nestedLayer3.AdjustTcpMss = &Layer3AdjustTcpMss{}
				if o.Layer3.AdjustTcpMss.Misc != nil {
					entry.Misc["Layer3AdjustTcpMss"] = o.Layer3.AdjustTcpMss.Misc
				}
				if o.Layer3.AdjustTcpMss.Ipv6MssAdjustment != nil {
					nestedLayer3.AdjustTcpMss.Ipv6MssAdjustment = o.Layer3.AdjustTcpMss.Ipv6MssAdjustment
				}
				if o.Layer3.AdjustTcpMss.Enable != nil {
					nestedLayer3.AdjustTcpMss.Enable = util.AsBool(o.Layer3.AdjustTcpMss.Enable, nil)
				}
				if o.Layer3.AdjustTcpMss.Ipv4MssAdjustment != nil {
					nestedLayer3.AdjustTcpMss.Ipv4MssAdjustment = o.Layer3.AdjustTcpMss.Ipv4MssAdjustment
				}
			}
			if o.Layer3.NdpProxy != nil {
				nestedLayer3.NdpProxy = &Layer3NdpProxy{}
				if o.Layer3.NdpProxy.Misc != nil {
					entry.Misc["Layer3NdpProxy"] = o.Layer3.NdpProxy.Misc
				}
				if o.Layer3.NdpProxy.Address != nil {
					nestedLayer3.NdpProxy.Address = []Layer3NdpProxyAddress{}
					for _, oLayer3NdpProxyAddress := range o.Layer3.NdpProxy.Address {
						nestedLayer3NdpProxyAddress := Layer3NdpProxyAddress{}
						if oLayer3NdpProxyAddress.Misc != nil {
							entry.Misc["Layer3NdpProxyAddress"] = oLayer3NdpProxyAddress.Misc
						}
						if oLayer3NdpProxyAddress.Negate != nil {
							nestedLayer3NdpProxyAddress.Negate = util.AsBool(oLayer3NdpProxyAddress.Negate, nil)
						}
						if oLayer3NdpProxyAddress.Name != "" {
							nestedLayer3NdpProxyAddress.Name = oLayer3NdpProxyAddress.Name
						}
						nestedLayer3.NdpProxy.Address = append(nestedLayer3.NdpProxy.Address, nestedLayer3NdpProxyAddress)
					}
				}
				if o.Layer3.NdpProxy.Enabled != nil {
					nestedLayer3.NdpProxy.Enabled = util.AsBool(o.Layer3.NdpProxy.Enabled, nil)
				}
			}
			if o.Layer3.SdwanLinkSettings != nil {
				nestedLayer3.SdwanLinkSettings = &Layer3SdwanLinkSettings{}
				if o.Layer3.SdwanLinkSettings.Misc != nil {
					entry.Misc["Layer3SdwanLinkSettings"] = o.Layer3.SdwanLinkSettings.Misc
				}
				if o.Layer3.SdwanLinkSettings.Enable != nil {
					nestedLayer3.SdwanLinkSettings.Enable = util.AsBool(o.Layer3.SdwanLinkSettings.Enable, nil)
				}
				if o.Layer3.SdwanLinkSettings.SdwanInterfaceProfile != nil {
					nestedLayer3.SdwanLinkSettings.SdwanInterfaceProfile = o.Layer3.SdwanLinkSettings.SdwanInterfaceProfile
				}
				if o.Layer3.SdwanLinkSettings.UpstreamNat != nil {
					nestedLayer3.SdwanLinkSettings.UpstreamNat = &Layer3SdwanLinkSettingsUpstreamNat{}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.Misc != nil {
						entry.Misc["Layer3SdwanLinkSettingsUpstreamNat"] = o.Layer3.SdwanLinkSettings.UpstreamNat.Misc
					}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.Enable != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.Enable = util.AsBool(o.Layer3.SdwanLinkSettings.UpstreamNat.Enable, nil)
					}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.Ddns != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.Ddns = &Layer3SdwanLinkSettingsUpstreamNatDdns{}
						if o.Layer3.SdwanLinkSettings.UpstreamNat.Ddns.Misc != nil {
							entry.Misc["Layer3SdwanLinkSettingsUpstreamNatDdns"] = o.Layer3.SdwanLinkSettings.UpstreamNat.Ddns.Misc
						}
					}
					if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp = &Layer3SdwanLinkSettingsUpstreamNatStaticIp{}
						if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Misc != nil {
							entry.Misc["Layer3SdwanLinkSettingsUpstreamNatStaticIp"] = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Misc
						}
						if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn != nil {
							nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.Fqdn
						}
						if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress != nil {
							nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp.IpAddress
						}
					}
				}
			}
			if o.Layer3.UntaggedSubInterface != nil {
				nestedLayer3.UntaggedSubInterface = util.AsBool(o.Layer3.UntaggedSubInterface, nil)
			}
			if o.Layer3.DdnsConfig != nil {
				nestedLayer3.DdnsConfig = &Layer3DdnsConfig{}
				if o.Layer3.DdnsConfig.Misc != nil {
					entry.Misc["Layer3DdnsConfig"] = o.Layer3.DdnsConfig.Misc
				}
				if o.Layer3.DdnsConfig.DdnsIp != nil {
					nestedLayer3.DdnsConfig.DdnsIp = util.MemToStr(o.Layer3.DdnsConfig.DdnsIp)
				}
				if o.Layer3.DdnsConfig.DdnsIpv6 != nil {
					nestedLayer3.DdnsConfig.DdnsIpv6 = util.MemToStr(o.Layer3.DdnsConfig.DdnsIpv6)
				}
				if o.Layer3.DdnsConfig.DdnsUpdateInterval != nil {
					nestedLayer3.DdnsConfig.DdnsUpdateInterval = o.Layer3.DdnsConfig.DdnsUpdateInterval
				}
				if o.Layer3.DdnsConfig.DdnsVendor != nil {
					nestedLayer3.DdnsConfig.DdnsVendor = o.Layer3.DdnsConfig.DdnsVendor
				}
				if o.Layer3.DdnsConfig.DdnsVendorConfig != nil {
					nestedLayer3.DdnsConfig.DdnsVendorConfig = []Layer3DdnsConfigDdnsVendorConfig{}
					for _, oLayer3DdnsConfigDdnsVendorConfig := range o.Layer3.DdnsConfig.DdnsVendorConfig {
						nestedLayer3DdnsConfigDdnsVendorConfig := Layer3DdnsConfigDdnsVendorConfig{}
						if oLayer3DdnsConfigDdnsVendorConfig.Misc != nil {
							entry.Misc["Layer3DdnsConfigDdnsVendorConfig"] = oLayer3DdnsConfigDdnsVendorConfig.Misc
						}
						if oLayer3DdnsConfigDdnsVendorConfig.Name != "" {
							nestedLayer3DdnsConfigDdnsVendorConfig.Name = oLayer3DdnsConfigDdnsVendorConfig.Name
						}
						if oLayer3DdnsConfigDdnsVendorConfig.Value != nil {
							nestedLayer3DdnsConfigDdnsVendorConfig.Value = oLayer3DdnsConfigDdnsVendorConfig.Value
						}
						nestedLayer3.DdnsConfig.DdnsVendorConfig = append(nestedLayer3.DdnsConfig.DdnsVendorConfig, nestedLayer3DdnsConfigDdnsVendorConfig)
					}
				}
				if o.Layer3.DdnsConfig.DdnsCertProfile != nil {
					nestedLayer3.DdnsConfig.DdnsCertProfile = o.Layer3.DdnsConfig.DdnsCertProfile
				}
				if o.Layer3.DdnsConfig.DdnsEnabled != nil {
					nestedLayer3.DdnsConfig.DdnsEnabled = util.AsBool(o.Layer3.DdnsConfig.DdnsEnabled, nil)
				}
				if o.Layer3.DdnsConfig.DdnsHostname != nil {
					nestedLayer3.DdnsConfig.DdnsHostname = o.Layer3.DdnsConfig.DdnsHostname
				}
			}
			if o.Layer3.DecryptForward != nil {
				nestedLayer3.DecryptForward = util.AsBool(o.Layer3.DecryptForward, nil)
			}
			if o.Layer3.DfIgnore != nil {
				nestedLayer3.DfIgnore = util.AsBool(o.Layer3.DfIgnore, nil)
			}
			if o.Layer3.DhcpClient != nil {
				nestedLayer3.DhcpClient = &Layer3DhcpClient{}
				if o.Layer3.DhcpClient.Misc != nil {
					entry.Misc["Layer3DhcpClient"] = o.Layer3.DhcpClient.Misc
				}
				if o.Layer3.DhcpClient.CreateDefaultRoute != nil {
					nestedLayer3.DhcpClient.CreateDefaultRoute = util.AsBool(o.Layer3.DhcpClient.CreateDefaultRoute, nil)
				}
				if o.Layer3.DhcpClient.DefaultRouteMetric != nil {
					nestedLayer3.DhcpClient.DefaultRouteMetric = o.Layer3.DhcpClient.DefaultRouteMetric
				}
				if o.Layer3.DhcpClient.Enable != nil {
					nestedLayer3.DhcpClient.Enable = util.AsBool(o.Layer3.DhcpClient.Enable, nil)
				}
				if o.Layer3.DhcpClient.SendHostname != nil {
					nestedLayer3.DhcpClient.SendHostname = &Layer3DhcpClientSendHostname{}
					if o.Layer3.DhcpClient.SendHostname.Misc != nil {
						entry.Misc["Layer3DhcpClientSendHostname"] = o.Layer3.DhcpClient.SendHostname.Misc
					}
					if o.Layer3.DhcpClient.SendHostname.Enable != nil {
						nestedLayer3.DhcpClient.SendHostname.Enable = util.AsBool(o.Layer3.DhcpClient.SendHostname.Enable, nil)
					}
					if o.Layer3.DhcpClient.SendHostname.Hostname != nil {
						nestedLayer3.DhcpClient.SendHostname.Hostname = o.Layer3.DhcpClient.SendHostname.Hostname
					}
				}
			}
			if o.Layer3.NetflowProfile != nil {
				nestedLayer3.NetflowProfile = o.Layer3.NetflowProfile
			}
			if o.Layer3.Bonjour != nil {
				nestedLayer3.Bonjour = &Layer3Bonjour{}
				if o.Layer3.Bonjour.Misc != nil {
					entry.Misc["Layer3Bonjour"] = o.Layer3.Bonjour.Misc
				}
				if o.Layer3.Bonjour.Enable != nil {
					nestedLayer3.Bonjour.Enable = util.AsBool(o.Layer3.Bonjour.Enable, nil)
				}
				if o.Layer3.Bonjour.GroupId != nil {
					nestedLayer3.Bonjour.GroupId = o.Layer3.Bonjour.GroupId
				}
				if o.Layer3.Bonjour.TtlCheck != nil {
					nestedLayer3.Bonjour.TtlCheck = util.AsBool(o.Layer3.Bonjour.TtlCheck, nil)
				}
			}
			if o.Layer3.Ipv6 != nil {
				nestedLayer3.Ipv6 = &Layer3Ipv6{}
				if o.Layer3.Ipv6.Misc != nil {
					entry.Misc["Layer3Ipv6"] = o.Layer3.Ipv6.Misc
				}
				if o.Layer3.Ipv6.Address != nil {
					nestedLayer3.Ipv6.Address = []Layer3Ipv6Address{}
					for _, oLayer3Ipv6Address := range o.Layer3.Ipv6.Address {
						nestedLayer3Ipv6Address := Layer3Ipv6Address{}
						if oLayer3Ipv6Address.Misc != nil {
							entry.Misc["Layer3Ipv6Address"] = oLayer3Ipv6Address.Misc
						}
						if oLayer3Ipv6Address.Anycast != nil {
							nestedLayer3Ipv6Address.Anycast = &Layer3Ipv6AddressAnycast{}
							if oLayer3Ipv6Address.Anycast.Misc != nil {
								entry.Misc["Layer3Ipv6AddressAnycast"] = oLayer3Ipv6Address.Anycast.Misc
							}
						}
						if oLayer3Ipv6Address.Advertise != nil {
							nestedLayer3Ipv6Address.Advertise = &Layer3Ipv6AddressAdvertise{}
							if oLayer3Ipv6Address.Advertise.Misc != nil {
								entry.Misc["Layer3Ipv6AddressAdvertise"] = oLayer3Ipv6Address.Advertise.Misc
							}
							if oLayer3Ipv6Address.Advertise.Enable != nil {
								nestedLayer3Ipv6Address.Advertise.Enable = util.AsBool(oLayer3Ipv6Address.Advertise.Enable, nil)
							}
							if oLayer3Ipv6Address.Advertise.ValidLifetime != nil {
								nestedLayer3Ipv6Address.Advertise.ValidLifetime = oLayer3Ipv6Address.Advertise.ValidLifetime
							}
							if oLayer3Ipv6Address.Advertise.PreferredLifetime != nil {
								nestedLayer3Ipv6Address.Advertise.PreferredLifetime = oLayer3Ipv6Address.Advertise.PreferredLifetime
							}
							if oLayer3Ipv6Address.Advertise.OnlinkFlag != nil {
								nestedLayer3Ipv6Address.Advertise.OnlinkFlag = util.AsBool(oLayer3Ipv6Address.Advertise.OnlinkFlag, nil)
							}
							if oLayer3Ipv6Address.Advertise.AutoConfigFlag != nil {
								nestedLayer3Ipv6Address.Advertise.AutoConfigFlag = util.AsBool(oLayer3Ipv6Address.Advertise.AutoConfigFlag, nil)
							}
						}
						if oLayer3Ipv6Address.Name != "" {
							nestedLayer3Ipv6Address.Name = oLayer3Ipv6Address.Name
						}
						if oLayer3Ipv6Address.EnableOnInterface != nil {
							nestedLayer3Ipv6Address.EnableOnInterface = util.AsBool(oLayer3Ipv6Address.EnableOnInterface, nil)
						}
						if oLayer3Ipv6Address.Prefix != nil {
							nestedLayer3Ipv6Address.Prefix = &Layer3Ipv6AddressPrefix{}
							if oLayer3Ipv6Address.Prefix.Misc != nil {
								entry.Misc["Layer3Ipv6AddressPrefix"] = oLayer3Ipv6Address.Prefix.Misc
							}
						}
						nestedLayer3.Ipv6.Address = append(nestedLayer3.Ipv6.Address, nestedLayer3Ipv6Address)
					}
				}
				if o.Layer3.Ipv6.Enabled != nil {
					nestedLayer3.Ipv6.Enabled = util.AsBool(o.Layer3.Ipv6.Enabled, nil)
				}
				if o.Layer3.Ipv6.InterfaceId != nil {
					nestedLayer3.Ipv6.InterfaceId = o.Layer3.Ipv6.InterfaceId
				}
				if o.Layer3.Ipv6.NeighborDiscovery != nil {
					nestedLayer3.Ipv6.NeighborDiscovery = &Layer3Ipv6NeighborDiscovery{}
					if o.Layer3.Ipv6.NeighborDiscovery.Misc != nil {
						entry.Misc["Layer3Ipv6NeighborDiscovery"] = o.Layer3.Ipv6.NeighborDiscovery.Misc
					}
					if o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.Neighbor != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.Neighbor = []Layer3Ipv6NeighborDiscoveryNeighbor{}
						for _, oLayer3Ipv6NeighborDiscoveryNeighbor := range o.Layer3.Ipv6.NeighborDiscovery.Neighbor {
							nestedLayer3Ipv6NeighborDiscoveryNeighbor := Layer3Ipv6NeighborDiscoveryNeighbor{}
							if oLayer3Ipv6NeighborDiscoveryNeighbor.Misc != nil {
								entry.Misc["Layer3Ipv6NeighborDiscoveryNeighbor"] = oLayer3Ipv6NeighborDiscoveryNeighbor.Misc
							}
							if oLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress != nil {
								nestedLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6NeighborDiscoveryNeighbor.HwAddress
							}
							if oLayer3Ipv6NeighborDiscoveryNeighbor.Name != "" {
								nestedLayer3Ipv6NeighborDiscoveryNeighbor.Name = oLayer3Ipv6NeighborDiscoveryNeighbor.Name
							}
							nestedLayer3.Ipv6.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.NeighborDiscovery.Neighbor, nestedLayer3Ipv6NeighborDiscoveryNeighbor)
						}
					}
					if o.Layer3.Ipv6.NeighborDiscovery.NsInterval != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.NeighborDiscovery.NsInterval
					}
					if o.Layer3.Ipv6.NeighborDiscovery.ReachableTime != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.NeighborDiscovery.ReachableTime
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement = &Layer3Ipv6NeighborDiscoveryRouterAdvertisement{}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc != nil {
							entry.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisement"] = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{}
							if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc != nil {
								entry.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"] = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc
							}
							if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
								nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
							}
							if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
								nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
								for _, oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc != nil {
										entry.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"] = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc
									}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
										nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
									}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
										nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
									}
									nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
								}
							}
							if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
								nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
								for _, oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
									nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc != nil {
										entry.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"] = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc
									}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
										nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
									}
									if oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
										nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
									}
									nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
								}
							}
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
						}
					}
					if o.Layer3.Ipv6.NeighborDiscovery.DadAttempts != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.NeighborDiscovery.DadAttempts
					}
					if o.Layer3.Ipv6.NeighborDiscovery.EnableDad != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.EnableDad = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.EnableDad, nil)
					}
				}
				if o.Layer3.Ipv6.DhcpClient != nil {
					nestedLayer3.Ipv6.DhcpClient = &Layer3Ipv6DhcpClient{}
					if o.Layer3.Ipv6.DhcpClient.Misc != nil {
						entry.Misc["Layer3Ipv6DhcpClient"] = o.Layer3.Ipv6.DhcpClient.Misc
					}
					if o.Layer3.Ipv6.DhcpClient.V6Options != nil {
						nestedLayer3.Ipv6.DhcpClient.V6Options = &Layer3Ipv6DhcpClientV6Options{}
						if o.Layer3.Ipv6.DhcpClient.V6Options.Misc != nil {
							entry.Misc["Layer3Ipv6DhcpClientV6Options"] = o.Layer3.Ipv6.DhcpClient.V6Options.Misc
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.RapidCommit = util.AsBool(o.Layer3.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig = util.AsBool(o.Layer3.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.DuidType != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.DuidType = o.Layer3.Ipv6.DhcpClient.V6Options.DuidType
						}
						if o.Layer3.Ipv6.DhcpClient.V6Options.Enable != nil {
							nestedLayer3.Ipv6.DhcpClient.V6Options.Enable = &Layer3Ipv6DhcpClientV6OptionsEnable{}
							if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Misc != nil {
								entry.Misc["Layer3Ipv6DhcpClientV6OptionsEnable"] = o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Misc
							}
							if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.No != nil {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.No = &Layer3Ipv6DhcpClientV6OptionsEnableNo{}
								if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.No.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientV6OptionsEnableNo"] = o.Layer3.Ipv6.DhcpClient.V6Options.Enable.No.Misc
								}
							}
							if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
								nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes = &Layer3Ipv6DhcpClientV6OptionsEnableYes{}
								if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientV6OptionsEnableYes"] = o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc
								}
								if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
									nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.AsBool(o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
								}
								if o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
									nestedLayer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.AsBool(o.Layer3.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
								}
							}
						}
					}
					if o.Layer3.Ipv6.DhcpClient.AcceptRaRoute != nil {
						nestedLayer3.Ipv6.DhcpClient.AcceptRaRoute = util.AsBool(o.Layer3.Ipv6.DhcpClient.AcceptRaRoute, nil)
					}
					if o.Layer3.Ipv6.DhcpClient.DefaultRouteMetric != nil {
						nestedLayer3.Ipv6.DhcpClient.DefaultRouteMetric = o.Layer3.Ipv6.DhcpClient.DefaultRouteMetric
					}
					if o.Layer3.Ipv6.DhcpClient.Enable != nil {
						nestedLayer3.Ipv6.DhcpClient.Enable = util.AsBool(o.Layer3.Ipv6.DhcpClient.Enable, nil)
					}
					if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery != nil {
						nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery = &Layer3Ipv6DhcpClientNeighborDiscovery{}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Misc != nil {
							entry.Misc["Layer3Ipv6DhcpClientNeighborDiscovery"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Misc
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor = []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor{}
							for _, oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
								nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor := Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor{}
								if oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor"] = oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Misc
								}
								if oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
									nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
								}
								if oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
									nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor.Name
								}
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.Neighbor, nestedLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor)
							}
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer{}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc != nil {
								entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource{}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
										entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
									}
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
										entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc
									}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
										nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
										for _, oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
												entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"] = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc
											}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
												nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
											}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
												nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
											}
											nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
										}
									}
								}
							}
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix{}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc != nil {
								entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
							}
							if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
								nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
										entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
									}
								}
								if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
									nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
										entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"] = o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
									}
									if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
										nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										for _, oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
											nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
												entry.Misc["Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
											}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
												nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
											}
											if oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
												nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
											}
											nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
										}
									}
								}
							}
						}
						if o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
							nestedLayer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad = util.AsBool(o.Layer3.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
						}
					}
					if o.Layer3.Ipv6.DhcpClient.Preference != nil {
						nestedLayer3.Ipv6.DhcpClient.Preference = o.Layer3.Ipv6.DhcpClient.Preference
					}
					if o.Layer3.Ipv6.DhcpClient.PrefixDelegation != nil {
						nestedLayer3.Ipv6.DhcpClient.PrefixDelegation = &Layer3Ipv6DhcpClientPrefixDelegation{}
						if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Misc != nil {
							entry.Misc["Layer3Ipv6DhcpClientPrefixDelegation"] = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Misc
						}
						if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
							nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable = &Layer3Ipv6DhcpClientPrefixDelegationEnable{}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc != nil {
								entry.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnable"] = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc
							}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No = &Layer3Ipv6DhcpClientPrefixDelegationEnableNo{}
								if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableNo"] = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc
								}
							}
							if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
								nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes = &Layer3Ipv6DhcpClientPrefixDelegationEnableYes{}
								if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc != nil {
									entry.Misc["Layer3Ipv6DhcpClientPrefixDelegationEnableYes"] = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc
								}
								if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
									nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
								}
								if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
									nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
								}
								if o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
									nestedLayer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.AsBool(o.Layer3.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
								}
							}
						}
					}
				}
				if o.Layer3.Ipv6.Inherited != nil {
					nestedLayer3.Ipv6.Inherited = &Layer3Ipv6Inherited{}
					if o.Layer3.Ipv6.Inherited.Misc != nil {
						entry.Misc["Layer3Ipv6Inherited"] = o.Layer3.Ipv6.Inherited.Misc
					}
					if o.Layer3.Ipv6.Inherited.AssignAddr != nil {
						nestedLayer3.Ipv6.Inherited.AssignAddr = []Layer3Ipv6InheritedAssignAddr{}
						for _, oLayer3Ipv6InheritedAssignAddr := range o.Layer3.Ipv6.Inherited.AssignAddr {
							nestedLayer3Ipv6InheritedAssignAddr := Layer3Ipv6InheritedAssignAddr{}
							if oLayer3Ipv6InheritedAssignAddr.Misc != nil {
								entry.Misc["Layer3Ipv6InheritedAssignAddr"] = oLayer3Ipv6InheritedAssignAddr.Misc
							}
							if oLayer3Ipv6InheritedAssignAddr.Type != nil {
								nestedLayer3Ipv6InheritedAssignAddr.Type = &Layer3Ipv6InheritedAssignAddrType{}
								if oLayer3Ipv6InheritedAssignAddr.Type.Misc != nil {
									entry.Misc["Layer3Ipv6InheritedAssignAddrType"] = oLayer3Ipv6InheritedAssignAddr.Type.Misc
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Gua != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Gua = &Layer3Ipv6InheritedAssignAddrTypeGua{}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedAssignAddrTypeGua"] = oLayer3Ipv6InheritedAssignAddr.Type.Gua.Misc
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PrefixPool
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolType{}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Misc != nil {
											entry.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolType"] = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Misc
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{}
											if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc != nil {
												entry.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"] = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc
											}
											if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
												nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
											}
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{}
											if oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc != nil {
												entry.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"] = oLayer3Ipv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc
											}
										}
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise = &Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise{}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Misc != nil {
											entry.Misc["Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise"] = oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Misc
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
										}
									}
								}
								if oLayer3Ipv6InheritedAssignAddr.Type.Ula != nil {
									nestedLayer3Ipv6InheritedAssignAddr.Type.Ula = &Layer3Ipv6InheritedAssignAddrTypeUla{}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedAssignAddrTypeUla"] = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Misc
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Address != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Address = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Address
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Prefix, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Anycast, nil)
									}
									if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise != nil {
										nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise = &Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise{}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Misc != nil {
											entry.Misc["Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise"] = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Misc
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
										}
										if oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
											nestedLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.AsBool(oLayer3Ipv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
										}
									}
								}
							}
							if oLayer3Ipv6InheritedAssignAddr.Name != "" {
								nestedLayer3Ipv6InheritedAssignAddr.Name = oLayer3Ipv6InheritedAssignAddr.Name
							}
							nestedLayer3.Ipv6.Inherited.AssignAddr = append(nestedLayer3.Ipv6.Inherited.AssignAddr, nestedLayer3Ipv6InheritedAssignAddr)
						}
					}
					if o.Layer3.Ipv6.Inherited.Enable != nil {
						nestedLayer3.Ipv6.Inherited.Enable = util.AsBool(o.Layer3.Ipv6.Inherited.Enable, nil)
					}
					if o.Layer3.Ipv6.Inherited.NeighborDiscovery != nil {
						nestedLayer3.Ipv6.Inherited.NeighborDiscovery = &Layer3Ipv6InheritedNeighborDiscovery{}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.Misc != nil {
							entry.Misc["Layer3Ipv6InheritedNeighborDiscovery"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.Misc
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor = []Layer3Ipv6InheritedNeighborDiscoveryNeighbor{}
							for _, oLayer3Ipv6InheritedNeighborDiscoveryNeighbor := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.Neighbor {
								nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor := Layer3Ipv6InheritedNeighborDiscoveryNeighbor{}
								if oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Misc != nil {
									entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryNeighbor"] = oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Misc
								}
								if oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
									nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress = oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.HwAddress
								}
								if oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name != "" {
									nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name = oLayer3Ipv6InheritedNeighborDiscoveryNeighbor.Name
								}
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.Neighbor, nestedLayer3Ipv6InheritedNeighborDiscoveryNeighbor)
							}
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.NsInterval
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DadAttempts
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer = &Layer3Ipv6InheritedNeighborDiscoveryDnsServer{}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc != nil {
								entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServer"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource{}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc != nil {
									entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc
									}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
										for _, oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
												entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"] = oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc
											}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
												nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
											}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
												nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
											}
											nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
										}
									}
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
									}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
									}
								}
							}
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix{}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc != nil {
								entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource{}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
									entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
									}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
									}
								}
								if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
									nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
										entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
									}
									if o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
										nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										for _, oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Layer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
											nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
												entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
											}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
												nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
											}
											if oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
												nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
											}
											nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedLayer3.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
										}
									}
								}
							}
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.EnableDad = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.ReachableTime
						}
						if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
							nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement{}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc != nil {
								entry.Misc["Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement"] = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
							}
							if o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
								nestedLayer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
							}
						}
					}
				}
			}
			if o.Layer3.Lldp != nil {
				nestedLayer3.Lldp = &Layer3Lldp{}
				if o.Layer3.Lldp.Misc != nil {
					entry.Misc["Layer3Lldp"] = o.Layer3.Lldp.Misc
				}
				if o.Layer3.Lldp.Profile != nil {
					nestedLayer3.Lldp.Profile = o.Layer3.Lldp.Profile
				}
				if o.Layer3.Lldp.Enable != nil {
					nestedLayer3.Lldp.Enable = util.AsBool(o.Layer3.Lldp.Enable, nil)
				}
				if o.Layer3.Lldp.HighAvailability != nil {
					nestedLayer3.Lldp.HighAvailability = &Layer3LldpHighAvailability{}
					if o.Layer3.Lldp.HighAvailability.Misc != nil {
						entry.Misc["Layer3LldpHighAvailability"] = o.Layer3.Lldp.HighAvailability.Misc
					}
					if o.Layer3.Lldp.HighAvailability.PassivePreNegotiation != nil {
						nestedLayer3.Lldp.HighAvailability.PassivePreNegotiation = util.AsBool(o.Layer3.Lldp.HighAvailability.PassivePreNegotiation, nil)
					}
				}
			}
			if o.Layer3.Arp != nil {
				nestedLayer3.Arp = []Layer3Arp{}
				for _, oLayer3Arp := range o.Layer3.Arp {
					nestedLayer3Arp := Layer3Arp{}
					if oLayer3Arp.Misc != nil {
						entry.Misc["Layer3Arp"] = oLayer3Arp.Misc
					}
					if oLayer3Arp.HwAddress != nil {
						nestedLayer3Arp.HwAddress = oLayer3Arp.HwAddress
					}
					if oLayer3Arp.Name != "" {
						nestedLayer3Arp.Name = oLayer3Arp.Name
					}
					nestedLayer3.Arp = append(nestedLayer3.Arp, nestedLayer3Arp)
				}
			}
			if o.Layer3.InterfaceManagementProfile != nil {
				nestedLayer3.InterfaceManagementProfile = o.Layer3.InterfaceManagementProfile
			}
			if o.Layer3.Ip != nil {
				nestedLayer3.Ip = []Layer3Ip{}
				for _, oLayer3Ip := range o.Layer3.Ip {
					nestedLayer3Ip := Layer3Ip{}
					if oLayer3Ip.Misc != nil {
						entry.Misc["Layer3Ip"] = oLayer3Ip.Misc
					}
					if oLayer3Ip.SdwanGateway != nil {
						nestedLayer3Ip.SdwanGateway = oLayer3Ip.SdwanGateway
					}
					if oLayer3Ip.Name != "" {
						nestedLayer3Ip.Name = oLayer3Ip.Name
					}
					nestedLayer3.Ip = append(nestedLayer3.Ip, nestedLayer3Ip)
				}
			}
			if o.Layer3.Lacp != nil {
				nestedLayer3.Lacp = &Layer3Lacp{}
				if o.Layer3.Lacp.Misc != nil {
					entry.Misc["Layer3Lacp"] = o.Layer3.Lacp.Misc
				}
				if o.Layer3.Lacp.Enable != nil {
					nestedLayer3.Lacp.Enable = util.AsBool(o.Layer3.Lacp.Enable, nil)
				}
				if o.Layer3.Lacp.FastFailover != nil {
					nestedLayer3.Lacp.FastFailover = util.AsBool(o.Layer3.Lacp.FastFailover, nil)
				}
				if o.Layer3.Lacp.HighAvailability != nil {
					nestedLayer3.Lacp.HighAvailability = &Layer3LacpHighAvailability{}
					if o.Layer3.Lacp.HighAvailability.Misc != nil {
						entry.Misc["Layer3LacpHighAvailability"] = o.Layer3.Lacp.HighAvailability.Misc
					}
					if o.Layer3.Lacp.HighAvailability.PassivePreNegotiation != nil {
						nestedLayer3.Lacp.HighAvailability.PassivePreNegotiation = util.AsBool(o.Layer3.Lacp.HighAvailability.PassivePreNegotiation, nil)
					}
				}
				if o.Layer3.Lacp.MaxPorts != nil {
					nestedLayer3.Lacp.MaxPorts = o.Layer3.Lacp.MaxPorts
				}
				if o.Layer3.Lacp.Mode != nil {
					nestedLayer3.Lacp.Mode = o.Layer3.Lacp.Mode
				}
				if o.Layer3.Lacp.SystemPriority != nil {
					nestedLayer3.Lacp.SystemPriority = o.Layer3.Lacp.SystemPriority
				}
				if o.Layer3.Lacp.TransmissionRate != nil {
					nestedLayer3.Lacp.TransmissionRate = o.Layer3.Lacp.TransmissionRate
				}
			}
			if o.Layer3.Mtu != nil {
				nestedLayer3.Mtu = o.Layer3.Mtu
			}
		}
		entry.Layer3 = nestedLayer3

		var nestedVirtualWire *VirtualWire
		if o.VirtualWire != nil {
			nestedVirtualWire = &VirtualWire{}
			if o.VirtualWire.Misc != nil {
				entry.Misc["VirtualWire"] = o.VirtualWire.Misc
			}
			if o.VirtualWire.Lldp != nil {
				nestedVirtualWire.Lldp = &VirtualWireLldp{}
				if o.VirtualWire.Lldp.Misc != nil {
					entry.Misc["VirtualWireLldp"] = o.VirtualWire.Lldp.Misc
				}
				if o.VirtualWire.Lldp.Enable != nil {
					nestedVirtualWire.Lldp.Enable = util.AsBool(o.VirtualWire.Lldp.Enable, nil)
				}
				if o.VirtualWire.Lldp.HighAvailability != nil {
					nestedVirtualWire.Lldp.HighAvailability = &VirtualWireLldpHighAvailability{}
					if o.VirtualWire.Lldp.HighAvailability.Misc != nil {
						entry.Misc["VirtualWireLldpHighAvailability"] = o.VirtualWire.Lldp.HighAvailability.Misc
					}
					if o.VirtualWire.Lldp.HighAvailability.PassivePreNegotiation != nil {
						nestedVirtualWire.Lldp.HighAvailability.PassivePreNegotiation = util.AsBool(o.VirtualWire.Lldp.HighAvailability.PassivePreNegotiation, nil)
					}
				}
				if o.VirtualWire.Lldp.Profile != nil {
					nestedVirtualWire.Lldp.Profile = o.VirtualWire.Lldp.Profile
				}
			}
			if o.VirtualWire.NetflowProfile != nil {
				nestedVirtualWire.NetflowProfile = o.VirtualWire.NetflowProfile
			}
		}
		entry.VirtualWire = nestedVirtualWire

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
	if !util.StringsMatch(a.Comment, b.Comment) {
		return false
	}
	if !matchDecryptMirror(a.DecryptMirror, b.DecryptMirror) {
		return false
	}
	if !matchHa(a.Ha, b.Ha) {
		return false
	}
	if !matchLayer2(a.Layer2, b.Layer2) {
		return false
	}
	if !matchLayer3(a.Layer3, b.Layer3) {
		return false
	}
	if !matchVirtualWire(a.VirtualWire, b.VirtualWire) {
		return false
	}

	return true
}

func matchDecryptMirror(a *DecryptMirror, b *DecryptMirror) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchHaLacp(a *HaLacp, b *HaLacp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.FastFailover, b.FastFailover) {
		return false
	}
	if !util.Ints64Match(a.MaxPorts, b.MaxPorts) {
		return false
	}
	if !util.StringsMatch(a.Mode, b.Mode) {
		return false
	}
	if !util.Ints64Match(a.SystemPriority, b.SystemPriority) {
		return false
	}
	if !util.StringsMatch(a.TransmissionRate, b.TransmissionRate) {
		return false
	}
	return true
}
func matchHa(a *Ha, b *Ha) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchHaLacp(a.Lacp, b.Lacp) {
		return false
	}
	return true
}
func matchLayer2LacpHighAvailability(a *Layer2LacpHighAvailability, b *Layer2LacpHighAvailability) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.PassivePreNegotiation, b.PassivePreNegotiation) {
		return false
	}
	return true
}
func matchLayer2Lacp(a *Layer2Lacp, b *Layer2Lacp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer2LacpHighAvailability(a.HighAvailability, b.HighAvailability) {
		return false
	}
	if !util.Ints64Match(a.MaxPorts, b.MaxPorts) {
		return false
	}
	if !util.StringsMatch(a.Mode, b.Mode) {
		return false
	}
	if !util.Ints64Match(a.SystemPriority, b.SystemPriority) {
		return false
	}
	if !util.StringsMatch(a.TransmissionRate, b.TransmissionRate) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.FastFailover, b.FastFailover) {
		return false
	}
	return true
}
func matchLayer2LldpHighAvailability(a *Layer2LldpHighAvailability, b *Layer2LldpHighAvailability) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.PassivePreNegotiation, b.PassivePreNegotiation) {
		return false
	}
	return true
}
func matchLayer2Lldp(a *Layer2Lldp, b *Layer2Lldp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer2LldpHighAvailability(a.HighAvailability, b.HighAvailability) {
		return false
	}
	return true
}
func matchLayer2(a *Layer2, b *Layer2) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer2Lldp(a.Lldp, b.Lldp) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
		return false
	}
	if !matchLayer2Lacp(a.Lacp, b.Lacp) {
		return false
	}
	return true
}
func matchLayer3Arp(a []Layer3Arp, b []Layer3Arp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ip(a []Layer3Ip, b []Layer3Ip) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.SdwanGateway, b.SdwanGateway) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3LacpHighAvailability(a *Layer3LacpHighAvailability, b *Layer3LacpHighAvailability) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.PassivePreNegotiation, b.PassivePreNegotiation) {
		return false
	}
	return true
}
func matchLayer3Lacp(a *Layer3Lacp, b *Layer3Lacp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.FastFailover, b.FastFailover) {
		return false
	}
	if !matchLayer3LacpHighAvailability(a.HighAvailability, b.HighAvailability) {
		return false
	}
	if !util.Ints64Match(a.MaxPorts, b.MaxPorts) {
		return false
	}
	if !util.StringsMatch(a.Mode, b.Mode) {
		return false
	}
	if !util.Ints64Match(a.SystemPriority, b.SystemPriority) {
		return false
	}
	if !util.StringsMatch(a.TransmissionRate, b.TransmissionRate) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchLayer3AdjustTcpMss(a *Layer3AdjustTcpMss, b *Layer3AdjustTcpMss) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.Ipv4MssAdjustment, b.Ipv4MssAdjustment) {
		return false
	}
	if !util.Ints64Match(a.Ipv6MssAdjustment, b.Ipv6MssAdjustment) {
		return false
	}
	return true
}
func matchLayer3NdpProxyAddress(a []Layer3NdpProxyAddress, b []Layer3NdpProxyAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.Negate, b.Negate) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3NdpProxy(a *Layer3NdpProxy, b *Layer3NdpProxy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3NdpProxyAddress(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	return true
}
func matchLayer3SdwanLinkSettingsUpstreamNatStaticIp(a *Layer3SdwanLinkSettingsUpstreamNatStaticIp, b *Layer3SdwanLinkSettingsUpstreamNatStaticIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.IpAddress, b.IpAddress) {
		return false
	}
	return true
}
func matchLayer3SdwanLinkSettingsUpstreamNatDdns(a *Layer3SdwanLinkSettingsUpstreamNatDdns, b *Layer3SdwanLinkSettingsUpstreamNatDdns) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3SdwanLinkSettingsUpstreamNat(a *Layer3SdwanLinkSettingsUpstreamNat, b *Layer3SdwanLinkSettingsUpstreamNat) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3SdwanLinkSettingsUpstreamNatDdns(a.Ddns, b.Ddns) {
		return false
	}
	if !matchLayer3SdwanLinkSettingsUpstreamNatStaticIp(a.StaticIp, b.StaticIp) {
		return false
	}
	return true
}
func matchLayer3SdwanLinkSettings(a *Layer3SdwanLinkSettings, b *Layer3SdwanLinkSettings) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SdwanInterfaceProfile, b.SdwanInterfaceProfile) {
		return false
	}
	if !matchLayer3SdwanLinkSettingsUpstreamNat(a.UpstreamNat, b.UpstreamNat) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchLayer3DdnsConfigDdnsVendorConfig(a []Layer3DdnsConfigDdnsVendorConfig, b []Layer3DdnsConfigDdnsVendorConfig) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3DdnsConfig(a *Layer3DdnsConfig, b *Layer3DdnsConfig) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DdnsHostname, b.DdnsHostname) {
		return false
	}
	if !util.OrderedListsMatch(a.DdnsIp, b.DdnsIp) {
		return false
	}
	if !util.OrderedListsMatch(a.DdnsIpv6, b.DdnsIpv6) {
		return false
	}
	if !util.Ints64Match(a.DdnsUpdateInterval, b.DdnsUpdateInterval) {
		return false
	}
	if !util.StringsMatch(a.DdnsVendor, b.DdnsVendor) {
		return false
	}
	if !matchLayer3DdnsConfigDdnsVendorConfig(a.DdnsVendorConfig, b.DdnsVendorConfig) {
		return false
	}
	if !util.StringsMatch(a.DdnsCertProfile, b.DdnsCertProfile) {
		return false
	}
	if !util.BoolsMatch(a.DdnsEnabled, b.DdnsEnabled) {
		return false
	}
	return true
}
func matchLayer3DhcpClientSendHostname(a *Layer3DhcpClientSendHostname, b *Layer3DhcpClientSendHostname) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.Hostname, b.Hostname) {
		return false
	}
	return true
}
func matchLayer3DhcpClient(a *Layer3DhcpClient, b *Layer3DhcpClient) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.CreateDefaultRoute, b.CreateDefaultRoute) {
		return false
	}
	if !util.Ints64Match(a.DefaultRouteMetric, b.DefaultRouteMetric) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3DhcpClientSendHostname(a.SendHostname, b.SendHostname) {
		return false
	}
	return true
}
func matchLayer3Bonjour(a *Layer3Bonjour, b *Layer3Bonjour) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.GroupId, b.GroupId) {
		return false
	}
	if !util.BoolsMatch(a.TtlCheck, b.TtlCheck) {
		return false
	}
	return true
}
func matchLayer3Ipv6AddressAnycast(a *Layer3Ipv6AddressAnycast, b *Layer3Ipv6AddressAnycast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6AddressAdvertise(a *Layer3Ipv6AddressAdvertise, b *Layer3Ipv6AddressAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.OnlinkFlag, b.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(a.AutoConfigFlag, b.AutoConfigFlag) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.ValidLifetime, b.ValidLifetime) {
		return false
	}
	if !util.StringsMatch(a.PreferredLifetime, b.PreferredLifetime) {
		return false
	}
	return true
}
func matchLayer3Ipv6AddressPrefix(a *Layer3Ipv6AddressPrefix, b *Layer3Ipv6AddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6Address(a []Layer3Ipv6Address, b []Layer3Ipv6Address) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
				return false
			}
			if !matchLayer3Ipv6AddressPrefix(a.Prefix, b.Prefix) {
				return false
			}
			if !matchLayer3Ipv6AddressAnycast(a.Anycast, b.Anycast) {
				return false
			}
			if !matchLayer3Ipv6AddressAdvertise(a.Advertise, b.Advertise) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryNeighbor(a []Layer3Ipv6NeighborDiscoveryNeighbor, b []Layer3Ipv6NeighborDiscoveryNeighbor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer(a []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, b []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) bool {
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
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix(a []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, b []Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport(a *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, b *Layer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer(a.Server, b.Server) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryRouterAdvertisement(a *Layer3Ipv6NeighborDiscoveryRouterAdvertisement, b *Layer3Ipv6NeighborDiscoveryRouterAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.RouterPreference, b.RouterPreference) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.EnableConsistencyCheck, b.EnableConsistencyCheck) {
		return false
	}
	if !util.StringsMatch(a.LinkMtu, b.LinkMtu) {
		return false
	}
	if !util.Ints64Match(a.MinInterval, b.MinInterval) {
		return false
	}
	if !util.StringsMatch(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.StringsMatch(a.RetransmissionTimer, b.RetransmissionTimer) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport(a.DnsSupport, b.DnsSupport) {
		return false
	}
	if !util.StringsMatch(a.HopLimit, b.HopLimit) {
		return false
	}
	if !util.Ints64Match(a.Lifetime, b.Lifetime) {
		return false
	}
	if !util.BoolsMatch(a.ManagedFlag, b.ManagedFlag) {
		return false
	}
	if !util.Ints64Match(a.MaxInterval, b.MaxInterval) {
		return false
	}
	if !util.BoolsMatch(a.OtherFlag, b.OtherFlag) {
		return false
	}
	return true
}
func matchLayer3Ipv6NeighborDiscovery(a *Layer3Ipv6NeighborDiscovery, b *Layer3Ipv6NeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6NeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryRouterAdvertisement(a.RouterAdvertisement, b.RouterAdvertisement) {
		return false
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientV6OptionsEnableNo(a *Layer3Ipv6DhcpClientV6OptionsEnableNo, b *Layer3Ipv6DhcpClientV6OptionsEnableNo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6DhcpClientV6OptionsEnableYes(a *Layer3Ipv6DhcpClientV6OptionsEnableYes, b *Layer3Ipv6DhcpClientV6OptionsEnableYes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.NonTempAddr, b.NonTempAddr) {
		return false
	}
	if !util.BoolsMatch(a.TempAddr, b.TempAddr) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientV6OptionsEnable(a *Layer3Ipv6DhcpClientV6OptionsEnable, b *Layer3Ipv6DhcpClientV6OptionsEnable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientV6OptionsEnableNo(a.No, b.No) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientV6OptionsEnableYes(a.Yes, b.Yes) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientV6Options(a *Layer3Ipv6DhcpClientV6Options, b *Layer3Ipv6DhcpClientV6Options) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DuidType, b.DuidType) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientV6OptionsEnable(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.RapidCommit, b.RapidCommit) {
		return false
	}
	if !util.BoolsMatch(a.SupportSrvrReconfig, b.SupportSrvrReconfig) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor(a []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor, b []Layer3Ipv6DhcpClientNeighborDiscoveryNeighbor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer(a []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, b []Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) bool {
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
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer(a.Server, b.Server) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServer(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServerSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix(a []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, b []Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix(a *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix, b *Layer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource(a.Source, b.Source) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientNeighborDiscovery(a *Layer3Ipv6DhcpClientNeighborDiscovery, b *Layer3Ipv6DhcpClientNeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsServer(a.DnsServer, b.DnsServer) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscoveryDnsSuffix(a.DnsSuffix, b.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientPrefixDelegationEnableNo(a *Layer3Ipv6DhcpClientPrefixDelegationEnableNo, b *Layer3Ipv6DhcpClientPrefixDelegationEnableNo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6DhcpClientPrefixDelegationEnableYes(a *Layer3Ipv6DhcpClientPrefixDelegationEnableYes, b *Layer3Ipv6DhcpClientPrefixDelegationEnableYes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.PrefixLen, b.PrefixLen) {
		return false
	}
	if !util.BoolsMatch(a.PrefixLenHint, b.PrefixLenHint) {
		return false
	}
	if !util.StringsMatch(a.PfxPoolName, b.PfxPoolName) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientPrefixDelegationEnable(a *Layer3Ipv6DhcpClientPrefixDelegationEnable, b *Layer3Ipv6DhcpClientPrefixDelegationEnable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientPrefixDelegationEnableNo(a.No, b.No) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientPrefixDelegationEnableYes(a.Yes, b.Yes) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClientPrefixDelegation(a *Layer3Ipv6DhcpClientPrefixDelegation, b *Layer3Ipv6DhcpClientPrefixDelegation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientPrefixDelegationEnable(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchLayer3Ipv6DhcpClient(a *Layer3Ipv6DhcpClient, b *Layer3Ipv6DhcpClient) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DhcpClientPrefixDelegation(a.PrefixDelegation, b.PrefixDelegation) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientV6Options(a.V6Options, b.V6Options) {
		return false
	}
	if !util.BoolsMatch(a.AcceptRaRoute, b.AcceptRaRoute) {
		return false
	}
	if !util.Ints64Match(a.DefaultRouteMetric, b.DefaultRouteMetric) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6DhcpClientNeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	if !util.StringsMatch(a.Preference, b.Preference) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryNeighbor(a []Layer3Ipv6InheritedNeighborDiscoveryNeighbor, b []Layer3Ipv6InheritedNeighborDiscoveryNeighbor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement(a *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement, b *Layer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.MaxInterval, b.MaxInterval) {
		return false
	}
	if !util.Ints64Match(a.MinInterval, b.MinInterval) {
		return false
	}
	if !util.StringsMatch(a.LinkMtu, b.LinkMtu) {
		return false
	}
	if !util.BoolsMatch(a.ManagedFlag, b.ManagedFlag) {
		return false
	}
	if !util.BoolsMatch(a.OtherFlag, b.OtherFlag) {
		return false
	}
	if !util.StringsMatch(a.RetransmissionTimer, b.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(a.RouterPreference, b.RouterPreference) {
		return false
	}
	if !util.BoolsMatch(a.EnableConsistencyCheck, b.EnableConsistencyCheck) {
		return false
	}
	if !util.StringsMatch(a.HopLimit, b.HopLimit) {
		return false
	}
	if !util.Ints64Match(a.Lifetime, b.Lifetime) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6(a *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, b *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.PrefixPool, b.PrefixPool) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer(a []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, b []Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual(a *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, b *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer(a.Server, b.Server) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSource(a *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource, b *Layer3Ipv6InheritedNeighborDiscoveryDnsServerSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsServer(a *Layer3Ipv6InheritedNeighborDiscoveryDnsServer, b *Layer3Ipv6InheritedNeighborDiscoveryDnsServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsServerSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6(a *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, b *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.PrefixPool, b.PrefixPool) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix(a []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, b []Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Lifetime, b.Lifetime) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual(a *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, b *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource(a *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource, b *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffix(a *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix, b *Layer3Ipv6InheritedNeighborDiscoveryDnsSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffixSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedNeighborDiscovery(a *Layer3Ipv6InheritedNeighborDiscovery, b *Layer3Ipv6InheritedNeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryRouterAdvertisement(a.RouterAdvertisement, b.RouterAdvertisement) {
		return false
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsServer(a.DnsServer, b.DnsServer) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryDnsSuffix(a.DnsSuffix, b.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeUlaAdvertise(a *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise, b *Layer3Ipv6InheritedAssignAddrTypeUlaAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.ValidLifetime, b.ValidLifetime) {
		return false
	}
	if !util.StringsMatch(a.PreferredLifetime, b.PreferredLifetime) {
		return false
	}
	if !util.BoolsMatch(a.OnlinkFlag, b.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(a.AutoConfigFlag, b.AutoConfigFlag) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeUla(a *Layer3Ipv6InheritedAssignAddrTypeUla, b *Layer3Ipv6InheritedAssignAddrTypeUla) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Prefix, b.Prefix) {
		return false
	}
	if !util.BoolsMatch(a.Anycast, b.Anycast) {
		return false
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeUlaAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic(a *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, b *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId(a *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, b *Layer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Identifier, b.Identifier) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolType(a *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType, b *Layer3Ipv6InheritedAssignAddrTypeGuaPoolType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic(a.Dynamic, b.Dynamic) {
		return false
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId(a.DynamicId, b.DynamicId) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeGuaAdvertise(a *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise, b *Layer3Ipv6InheritedAssignAddrTypeGuaAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.OnlinkFlag, b.OnlinkFlag) {
		return false
	}
	if !util.BoolsMatch(a.AutoConfigFlag, b.AutoConfigFlag) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrTypeGua(a *Layer3Ipv6InheritedAssignAddrTypeGua, b *Layer3Ipv6InheritedAssignAddrTypeGua) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
		return false
	}
	if !util.StringsMatch(a.PrefixPool, b.PrefixPool) {
		return false
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeGuaPoolType(a.PoolType, b.PoolType) {
		return false
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeGuaAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddrType(a *Layer3Ipv6InheritedAssignAddrType, b *Layer3Ipv6InheritedAssignAddrType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeUla(a.Ula, b.Ula) {
		return false
	}
	if !matchLayer3Ipv6InheritedAssignAddrTypeGua(a.Gua, b.Gua) {
		return false
	}
	return true
}
func matchLayer3Ipv6InheritedAssignAddr(a []Layer3Ipv6InheritedAssignAddr, b []Layer3Ipv6InheritedAssignAddr) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchLayer3Ipv6InheritedAssignAddrType(a.Type, b.Type) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6Inherited(a *Layer3Ipv6Inherited, b *Layer3Ipv6Inherited) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6InheritedAssignAddr(a.AssignAddr, b.AssignAddr) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6InheritedNeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	return true
}
func matchLayer3Ipv6(a *Layer3Ipv6, b *Layer3Ipv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6Address(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !util.StringsMatch(a.InterfaceId, b.InterfaceId) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	if !matchLayer3Ipv6DhcpClient(a.DhcpClient, b.DhcpClient) {
		return false
	}
	if !matchLayer3Ipv6Inherited(a.Inherited, b.Inherited) {
		return false
	}
	return true
}
func matchLayer3LldpHighAvailability(a *Layer3LldpHighAvailability, b *Layer3LldpHighAvailability) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.PassivePreNegotiation, b.PassivePreNegotiation) {
		return false
	}
	return true
}
func matchLayer3Lldp(a *Layer3Lldp, b *Layer3Lldp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3LldpHighAvailability(a.HighAvailability, b.HighAvailability) {
		return false
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchLayer3(a *Layer3, b *Layer3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.UntaggedSubInterface, b.UntaggedSubInterface) {
		return false
	}
	if !matchLayer3DdnsConfig(a.DdnsConfig, b.DdnsConfig) {
		return false
	}
	if !util.BoolsMatch(a.DecryptForward, b.DecryptForward) {
		return false
	}
	if !util.BoolsMatch(a.DfIgnore, b.DfIgnore) {
		return false
	}
	if !matchLayer3DhcpClient(a.DhcpClient, b.DhcpClient) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
		return false
	}
	if !matchLayer3Bonjour(a.Bonjour, b.Bonjour) {
		return false
	}
	if !matchLayer3Ipv6(a.Ipv6, b.Ipv6) {
		return false
	}
	if !matchLayer3Lldp(a.Lldp, b.Lldp) {
		return false
	}
	if !matchLayer3Arp(a.Arp, b.Arp) {
		return false
	}
	if !util.StringsMatch(a.InterfaceManagementProfile, b.InterfaceManagementProfile) {
		return false
	}
	if !matchLayer3Ip(a.Ip, b.Ip) {
		return false
	}
	if !matchLayer3Lacp(a.Lacp, b.Lacp) {
		return false
	}
	if !util.Ints64Match(a.Mtu, b.Mtu) {
		return false
	}
	if !matchLayer3AdjustTcpMss(a.AdjustTcpMss, b.AdjustTcpMss) {
		return false
	}
	if !matchLayer3NdpProxy(a.NdpProxy, b.NdpProxy) {
		return false
	}
	if !matchLayer3SdwanLinkSettings(a.SdwanLinkSettings, b.SdwanLinkSettings) {
		return false
	}
	return true
}
func matchVirtualWireLldpHighAvailability(a *VirtualWireLldpHighAvailability, b *VirtualWireLldpHighAvailability) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.PassivePreNegotiation, b.PassivePreNegotiation) {
		return false
	}
	return true
}
func matchVirtualWireLldp(a *VirtualWireLldp, b *VirtualWireLldp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchVirtualWireLldpHighAvailability(a.HighAvailability, b.HighAvailability) {
		return false
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	return true
}
func matchVirtualWire(a *VirtualWire, b *VirtualWire) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchVirtualWireLldp(a.Lldp, b.Lldp) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
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
