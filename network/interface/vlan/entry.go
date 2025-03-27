package vlan

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
	Suffix = []string{"network", "interface", "vlan", "units"}
)

type Entry struct {
	Name                       string
	AdjustTcpMss               *AdjustTcpMss
	Arp                        []Arp
	Bonjour                    *Bonjour
	Comment                    *string
	DdnsConfig                 *DdnsConfig
	DfIgnore                   *bool
	DhcpClient                 *DhcpClient
	InterfaceManagementProfile *string
	Ip                         []Ip
	Ipv6                       *Ipv6
	Mtu                        *int64
	NdpProxy                   *NdpProxy
	NetflowProfile             *string

	Misc map[string][]generic.Xml
}

type AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
}
type Arp struct {
	Name      string
	HwAddress *string
	Interface *string
}
type Bonjour struct {
	Enable   *bool
	GroupId  *int64
	TtlCheck *bool
}
type DdnsConfig struct {
	DdnsCertProfile    *string
	DdnsEnabled        *bool
	DdnsHostname       *string
	DdnsIp             []string
	DdnsIpv6           []string
	DdnsUpdateInterval *int64
	DdnsVendor         *string
	DdnsVendorConfig   []DdnsConfigDdnsVendorConfig
}
type DdnsConfigDdnsVendorConfig struct {
	Name  string
	Value *string
}
type DhcpClient struct {
	CreateDefaultRoute *bool
	DefaultRouteMetric *int64
	Enable             *bool
	SendHostname       *DhcpClientSendHostname
}
type DhcpClientSendHostname struct {
	Enable   *bool
	Hostname *string
}
type Ip struct {
	Name string
}
type Ipv6 struct {
	Address           []Ipv6Address
	Enabled           *bool
	InterfaceId       *string
	NeighborDiscovery *Ipv6NeighborDiscovery
	DhcpClient        *Ipv6DhcpClient
	Inherited         *Ipv6Inherited
}
type Ipv6Address struct {
	Name              string
	EnableOnInterface *bool
	Prefix            *Ipv6AddressPrefix
	Anycast           *Ipv6AddressAnycast
	Advertise         *Ipv6AddressAdvertise
}
type Ipv6AddressAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
}
type Ipv6AddressAnycast struct {
}
type Ipv6AddressPrefix struct {
}
type Ipv6DhcpClient struct {
	AcceptRaRoute      *bool
	DefaultRouteMetric *int64
	Enable             *bool
	NeighborDiscovery  *Ipv6DhcpClientNeighborDiscovery
	Preference         *string
	PrefixDelegation   *Ipv6DhcpClientPrefixDelegation
	V6Options          *Ipv6DhcpClientV6Options
}
type Ipv6DhcpClientNeighborDiscovery struct {
	DadAttempts      *int64
	DnsServer        *Ipv6DhcpClientNeighborDiscoveryDnsServer
	DnsSuffix        *Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	EnableDad        *bool
	EnableNdpMonitor *bool
	Neighbor         []Ipv6DhcpClientNeighborDiscoveryNeighbor
	NsInterval       *int64
	ReachableTime    *int64
}
type Ipv6DhcpClientNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Ipv6DhcpClientNeighborDiscoveryDnsServerSource
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6 struct {
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual struct {
	Server []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer struct {
	Name     string
	Lifetime *int64
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name     string
	Lifetime *int64
}
type Ipv6DhcpClientNeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
}
type Ipv6DhcpClientPrefixDelegation struct {
	Enable *Ipv6DhcpClientPrefixDelegationEnable
}
type Ipv6DhcpClientPrefixDelegationEnable struct {
	No  *Ipv6DhcpClientPrefixDelegationEnableNo
	Yes *Ipv6DhcpClientPrefixDelegationEnableYes
}
type Ipv6DhcpClientPrefixDelegationEnableNo struct {
}
type Ipv6DhcpClientPrefixDelegationEnableYes struct {
	PfxPoolName   *string
	PrefixLen     *int64
	PrefixLenHint *bool
}
type Ipv6DhcpClientV6Options struct {
	DuidType            *string
	Enable              *Ipv6DhcpClientV6OptionsEnable
	RapidCommit         *bool
	SupportSrvrReconfig *bool
}
type Ipv6DhcpClientV6OptionsEnable struct {
	No  *Ipv6DhcpClientV6OptionsEnableNo
	Yes *Ipv6DhcpClientV6OptionsEnableYes
}
type Ipv6DhcpClientV6OptionsEnableNo struct {
}
type Ipv6DhcpClientV6OptionsEnableYes struct {
	NonTempAddr *bool
	TempAddr    *bool
}
type Ipv6Inherited struct {
	AssignAddr        []Ipv6InheritedAssignAddr
	Enable            *bool
	NeighborDiscovery *Ipv6InheritedNeighborDiscovery
}
type Ipv6InheritedAssignAddr struct {
	Name string
	Type *Ipv6InheritedAssignAddrType
}
type Ipv6InheritedAssignAddrType struct {
	Gua *Ipv6InheritedAssignAddrTypeGua
	Ula *Ipv6InheritedAssignAddrTypeUla
}
type Ipv6InheritedAssignAddrTypeGua struct {
	EnableOnInterface *bool
	PrefixPool        *string
	PoolType          *Ipv6InheritedAssignAddrTypeGuaPoolType
	Advertise         *Ipv6InheritedAssignAddrTypeGuaAdvertise
}
type Ipv6InheritedAssignAddrTypeGuaAdvertise struct {
	Enable         *bool
	OnlinkFlag     *bool
	AutoConfigFlag *bool
}
type Ipv6InheritedAssignAddrTypeGuaPoolType struct {
	Dynamic   *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	DynamicId *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic struct {
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId struct {
	Identifier *int64
}
type Ipv6InheritedAssignAddrTypeUla struct {
	EnableOnInterface *bool
	Address           *string
	Prefix            *bool
	Anycast           *bool
	Advertise         *Ipv6InheritedAssignAddrTypeUlaAdvertise
}
type Ipv6InheritedAssignAddrTypeUlaAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
}
type Ipv6InheritedNeighborDiscovery struct {
	DadAttempts         *int64
	DnsServer           *Ipv6InheritedNeighborDiscoveryDnsServer
	DnsSuffix           *Ipv6InheritedNeighborDiscoveryDnsSuffix
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Ipv6InheritedNeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Ipv6InheritedNeighborDiscoveryRouterAdvertisement
}
type Ipv6InheritedNeighborDiscoveryDnsServer struct {
	Enable *bool
	Source *Ipv6InheritedNeighborDiscoveryDnsServerSource
}
type Ipv6InheritedNeighborDiscoveryDnsServerSource struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	Manual *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	PrefixPool *string
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManual struct {
	Server []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer struct {
	Name     string
	Lifetime *int64
}
type Ipv6InheritedNeighborDiscoveryDnsSuffix struct {
	Enable *bool
	Source *Ipv6InheritedNeighborDiscoveryDnsSuffixSource
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	PrefixPool *string
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name     string
	Lifetime *int64
}
type Ipv6InheritedNeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
}
type Ipv6InheritedNeighborDiscoveryRouterAdvertisement struct {
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
type Ipv6NeighborDiscovery struct {
	DadAttempts         *int64
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Ipv6NeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Ipv6NeighborDiscoveryRouterAdvertisement
}
type Ipv6NeighborDiscoveryNeighbor struct {
	Name      string
	HwAddress *string
}
type Ipv6NeighborDiscoveryRouterAdvertisement struct {
	DnsSupport             *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
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
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport struct {
	Enable *bool
	Server []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	Suffix []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer struct {
	Name     string
	Lifetime *int64
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix struct {
	Name     string
	Lifetime *int64
}
type NdpProxy struct {
	Address []NdpProxyAddress
	Enabled *bool
}
type NdpProxyAddress struct {
	Name   string
	Negate *bool
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}
type entryXml struct {
	XMLName                    xml.Name         `xml:"entry"`
	Name                       string           `xml:"name,attr"`
	AdjustTcpMss               *AdjustTcpMssXml `xml:"adjust-tcp-mss,omitempty"`
	Arp                        []ArpXml         `xml:"arp>entry,omitempty"`
	Bonjour                    *BonjourXml      `xml:"bonjour,omitempty"`
	Comment                    *string          `xml:"comment,omitempty"`
	DdnsConfig                 *DdnsConfigXml   `xml:"ddns-config,omitempty"`
	DfIgnore                   *string          `xml:"df-ignore,omitempty"`
	DhcpClient                 *DhcpClientXml   `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string          `xml:"interface-management-profile,omitempty"`
	Ip                         []IpXml          `xml:"ip>entry,omitempty"`
	Ipv6                       *Ipv6Xml         `xml:"ipv6,omitempty"`
	Mtu                        *int64           `xml:"mtu,omitempty"`
	NdpProxy                   *NdpProxyXml     `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string          `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name                `xml:"entry"`
	Name                       string                  `xml:"name,attr"`
	AdjustTcpMss               *AdjustTcpMssXml_11_0_2 `xml:"adjust-tcp-mss,omitempty"`
	Arp                        []ArpXml_11_0_2         `xml:"arp>entry,omitempty"`
	Bonjour                    *BonjourXml_11_0_2      `xml:"bonjour,omitempty"`
	Comment                    *string                 `xml:"comment,omitempty"`
	DdnsConfig                 *DdnsConfigXml_11_0_2   `xml:"ddns-config,omitempty"`
	DfIgnore                   *string                 `xml:"df-ignore,omitempty"`
	DhcpClient                 *DhcpClientXml_11_0_2   `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                 `xml:"interface-management-profile,omitempty"`
	Ip                         []IpXml_11_0_2          `xml:"ip>entry,omitempty"`
	Ipv6                       *Ipv6Xml_11_0_2         `xml:"ipv6,omitempty"`
	Mtu                        *int64                  `xml:"mtu,omitempty"`
	NdpProxy                   *NdpProxyXml_11_0_2     `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string                 `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AdjustTcpMssXml struct {
	Enable            *string `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64  `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64  `xml:"ipv6-mss-adjustment,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ArpXml struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Interface *string `xml:"interface,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type BonjourXml struct {
	Enable   *string `xml:"enable,omitempty"`
	GroupId  *int64  `xml:"group-id,omitempty"`
	TtlCheck *string `xml:"ttl-check,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DdnsConfigXml struct {
	DdnsCertProfile    *string                         `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                         `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                         `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                          `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                         `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   []DdnsConfigDdnsVendorConfigXml `xml:"ddns-vendor-config>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DdnsConfigDdnsVendorConfigXml struct {
	Name  string  `xml:"name,attr"`
	Value *string `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DhcpClientXml struct {
	CreateDefaultRoute *string                    `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                     `xml:"default-route-metric,omitempty"`
	Enable             *string                    `xml:"enable,omitempty"`
	SendHostname       *DhcpClientSendHostnameXml `xml:"send-hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DhcpClientSendHostnameXml struct {
	Enable   *string `xml:"enable,omitempty"`
	Hostname *string `xml:"hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type IpXml struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6Xml struct {
	Address           []Ipv6AddressXml          `xml:"address>entry,omitempty"`
	DhcpClient        *Ipv6DhcpClientXml        `xml:"dhcp-client,omitempty"`
	Enabled           *string                   `xml:"enabled,omitempty"`
	Inherited         *Ipv6InheritedXml         `xml:"inherited,omitempty"`
	InterfaceId       *string                   `xml:"interface-id,omitempty"`
	NeighborDiscovery *Ipv6NeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressXml struct {
	Advertise         *Ipv6AddressAdvertiseXml `xml:"advertise,omitempty"`
	Anycast           *Ipv6AddressAnycastXml   `xml:"anycast,omitempty"`
	EnableOnInterface *string                  `xml:"enable-on-interface,omitempty"`
	Name              string                   `xml:"name,attr"`
	Prefix            *Ipv6AddressPrefixXml    `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressAdvertiseXml struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressAnycastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressPrefixXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientXml struct {
	AcceptRaRoute      *string                             `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                              `xml:"default-route-metric,omitempty"`
	Enable             *string                             `xml:"enable,omitempty"`
	NeighborDiscovery  *Ipv6DhcpClientNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Preference         *string                             `xml:"preference,omitempty"`
	PrefixDelegation   *Ipv6DhcpClientPrefixDelegationXml  `xml:"prefix-delegation,omitempty"`
	V6Options          *Ipv6DhcpClientV6OptionsXml         `xml:"v6-options,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryXml struct {
	DadAttempts      *int64                                       `xml:"dad-attempts,omitempty"`
	DnsServer        *Ipv6DhcpClientNeighborDiscoveryDnsServerXml `xml:"dns-server,omitempty"`
	DnsSuffix        *Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                      `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                      `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         []Ipv6DhcpClientNeighborDiscoveryNeighborXml `xml:"neighbor>entry,omitempty"`
	NsInterval       *int64                                       `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                       `xml:"reachable-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerXml struct {
	Enable *string                                            `xml:"enable,omitempty"`
	Source *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml struct {
	Server []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                            `xml:"enable,omitempty"`
	Source *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryNeighborXml struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationXml struct {
	Enable *Ipv6DhcpClientPrefixDelegationEnableXml `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableXml struct {
	No  *Ipv6DhcpClientPrefixDelegationEnableNoXml  `xml:"no,omitempty"`
	Yes *Ipv6DhcpClientPrefixDelegationEnableYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableYesXml struct {
	PfxPoolName   *string `xml:"pfx-pool-name,omitempty"`
	PrefixLen     *int64  `xml:"prefix-len,omitempty"`
	PrefixLenHint *string `xml:"prefix-len-hint,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsXml struct {
	DuidType            *string                           `xml:"duid-type,omitempty"`
	Enable              *Ipv6DhcpClientV6OptionsEnableXml `xml:"enable,omitempty"`
	RapidCommit         *string                           `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                           `xml:"support-srvr-reconfig,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableXml struct {
	No  *Ipv6DhcpClientV6OptionsEnableNoXml  `xml:"no,omitempty"`
	Yes *Ipv6DhcpClientV6OptionsEnableYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableNoXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableYesXml struct {
	NonTempAddr *string `xml:"non-temp-addr,omitempty"`
	TempAddr    *string `xml:"temp-addr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedXml struct {
	AssignAddr        []Ipv6InheritedAssignAddrXml       `xml:"assign-addr>entry,omitempty"`
	Enable            *string                            `xml:"enable,omitempty"`
	NeighborDiscovery *Ipv6InheritedNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrXml struct {
	Name string                          `xml:"name,attr"`
	Type *Ipv6InheritedAssignAddrTypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeXml struct {
	Gua *Ipv6InheritedAssignAddrTypeGuaXml `xml:"gua,omitempty"`
	Ula *Ipv6InheritedAssignAddrTypeUlaXml `xml:"ula,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaXml struct {
	Advertise         *Ipv6InheritedAssignAddrTypeGuaAdvertiseXml `xml:"advertise,omitempty"`
	EnableOnInterface *string                                     `xml:"enable-on-interface,omitempty"`
	PoolType          *Ipv6InheritedAssignAddrTypeGuaPoolTypeXml  `xml:"pool-type,omitempty"`
	PrefixPool        *string                                     `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaAdvertiseXml struct {
	AutoConfigFlag *string `xml:"auto-config-flag,omitempty"`
	Enable         *string `xml:"enable,omitempty"`
	OnlinkFlag     *string `xml:"onlink-flag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeXml struct {
	Dynamic   *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml   `xml:"dynamic,omitempty"`
	DynamicId *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml `xml:"dynamic-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml struct {
	Identifier *int64 `xml:"identifier,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeUlaXml struct {
	Address           *string                                     `xml:"address,omitempty"`
	Advertise         *Ipv6InheritedAssignAddrTypeUlaAdvertiseXml `xml:"advertise,omitempty"`
	Anycast           *string                                     `xml:"anycast,omitempty"`
	EnableOnInterface *string                                     `xml:"enable-on-interface,omitempty"`
	Prefix            *string                                     `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeUlaAdvertiseXml struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryXml struct {
	DadAttempts         *int64                                                `xml:"dad-attempts,omitempty"`
	DnsServer           *Ipv6InheritedNeighborDiscoveryDnsServerXml           `xml:"dns-server,omitempty"`
	DnsSuffix           *Ipv6InheritedNeighborDiscoveryDnsSuffixXml           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                               `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                               `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Ipv6InheritedNeighborDiscoveryNeighborXml           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                                `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerXml struct {
	Enable *string                                           `xml:"enable,omitempty"`
	Source *Ipv6InheritedNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml struct {
	Server []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixXml struct {
	Enable *string                                           `xml:"enable,omitempty"`
	Source *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryNeighborXml struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml struct {
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
type Ipv6NeighborDiscoveryXml struct {
	DadAttempts         *int64                                       `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                      `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                      `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Ipv6NeighborDiscoveryNeighborXml           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                       `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                       `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Ipv6NeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryNeighborXml struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementXml struct {
	DnsSupport             *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml `xml:"dns-support,omitempty"`
	Enable                 *string                                                `xml:"enable,omitempty"`
	EnableConsistencyCheck *string                                                `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string                                                `xml:"hop-limit,omitempty"`
	Lifetime               *int64                                                 `xml:"lifetime,omitempty"`
	LinkMtu                *string                                                `xml:"link-mtu,omitempty"`
	ManagedFlag            *string                                                `xml:"managed-flag,omitempty"`
	MaxInterval            *int64                                                 `xml:"max-interval,omitempty"`
	MinInterval            *int64                                                 `xml:"min-interval,omitempty"`
	OtherFlag              *string                                                `xml:"other-flag,omitempty"`
	ReachableTime          *string                                                `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string                                                `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string                                                `xml:"router-preference,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml struct {
	Enable *string                                                       `xml:"enable,omitempty"`
	Server []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml `xml:"server>entry,omitempty"`
	Suffix []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type NdpProxyXml struct {
	Address []NdpProxyAddressXml `xml:"address>entry,omitempty"`
	Enabled *string              `xml:"enabled,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NdpProxyAddressXml struct {
	Name   string  `xml:"name,attr"`
	Negate *string `xml:"negate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AdjustTcpMssXml_11_0_2 struct {
	Enable            *string `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64  `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64  `xml:"ipv6-mss-adjustment,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ArpXml_11_0_2 struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Interface *string `xml:"interface,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type BonjourXml_11_0_2 struct {
	Enable   *string `xml:"enable,omitempty"`
	GroupId  *int64  `xml:"group-id,omitempty"`
	TtlCheck *string `xml:"ttl-check,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DdnsConfigXml_11_0_2 struct {
	DdnsCertProfile    *string                                `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                                `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                                `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                       `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                       `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                 `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                                `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   []DdnsConfigDdnsVendorConfigXml_11_0_2 `xml:"ddns-vendor-config>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DdnsConfigDdnsVendorConfigXml_11_0_2 struct {
	Name  string  `xml:"name,attr"`
	Value *string `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DhcpClientXml_11_0_2 struct {
	CreateDefaultRoute *string                           `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                            `xml:"default-route-metric,omitempty"`
	Enable             *string                           `xml:"enable,omitempty"`
	SendHostname       *DhcpClientSendHostnameXml_11_0_2 `xml:"send-hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DhcpClientSendHostnameXml_11_0_2 struct {
	Enable   *string `xml:"enable,omitempty"`
	Hostname *string `xml:"hostname,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type IpXml_11_0_2 struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6Xml_11_0_2 struct {
	Address           []Ipv6AddressXml_11_0_2          `xml:"address>entry,omitempty"`
	DhcpClient        *Ipv6DhcpClientXml_11_0_2        `xml:"dhcp-client,omitempty"`
	Enabled           *string                          `xml:"enabled,omitempty"`
	Inherited         *Ipv6InheritedXml_11_0_2         `xml:"inherited,omitempty"`
	InterfaceId       *string                          `xml:"interface-id,omitempty"`
	NeighborDiscovery *Ipv6NeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressXml_11_0_2 struct {
	Advertise         *Ipv6AddressAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Anycast           *Ipv6AddressAnycastXml_11_0_2   `xml:"anycast,omitempty"`
	EnableOnInterface *string                         `xml:"enable-on-interface,omitempty"`
	Name              string                          `xml:"name,attr"`
	Prefix            *Ipv6AddressPrefixXml_11_0_2    `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressAdvertiseXml_11_0_2 struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressAnycastXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressPrefixXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientXml_11_0_2 struct {
	AcceptRaRoute      *string                                    `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                                     `xml:"default-route-metric,omitempty"`
	Enable             *string                                    `xml:"enable,omitempty"`
	NeighborDiscovery  *Ipv6DhcpClientNeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`
	Preference         *string                                    `xml:"preference,omitempty"`
	PrefixDelegation   *Ipv6DhcpClientPrefixDelegationXml_11_0_2  `xml:"prefix-delegation,omitempty"`
	V6Options          *Ipv6DhcpClientV6OptionsXml_11_0_2         `xml:"v6-options,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts      *int64                                              `xml:"dad-attempts,omitempty"`
	DnsServer        *Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2 `xml:"dns-server,omitempty"`
	DnsSuffix        *Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2 `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                             `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                             `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         []Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 `xml:"neighbor>entry,omitempty"`
	NsInterval       *int64                                              `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                              `xml:"reachable-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable *string                                                   `xml:"enable,omitempty"`
	Source *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable *string                                                   `xml:"enable,omitempty"`
	Source *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6 *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationXml_11_0_2 struct {
	Enable *Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 struct {
	No  *Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes *Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 struct {
	PfxPoolName   *string `xml:"pfx-pool-name,omitempty"`
	PrefixLen     *int64  `xml:"prefix-len,omitempty"`
	PrefixLenHint *string `xml:"prefix-len-hint,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsXml_11_0_2 struct {
	DuidType            *string                                  `xml:"duid-type,omitempty"`
	Enable              *Ipv6DhcpClientV6OptionsEnableXml_11_0_2 `xml:"enable,omitempty"`
	RapidCommit         *string                                  `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                                  `xml:"support-srvr-reconfig,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableXml_11_0_2 struct {
	No  *Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes *Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 struct {
	NonTempAddr *string `xml:"non-temp-addr,omitempty"`
	TempAddr    *string `xml:"temp-addr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedXml_11_0_2 struct {
	AssignAddr        []Ipv6InheritedAssignAddrXml_11_0_2       `xml:"assign-addr>entry,omitempty"`
	Enable            *string                                   `xml:"enable,omitempty"`
	NeighborDiscovery *Ipv6InheritedNeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrXml_11_0_2 struct {
	Name string                                 `xml:"name,attr"`
	Type *Ipv6InheritedAssignAddrTypeXml_11_0_2 `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeXml_11_0_2 struct {
	Gua *Ipv6InheritedAssignAddrTypeGuaXml_11_0_2 `xml:"gua,omitempty"`
	Ula *Ipv6InheritedAssignAddrTypeUlaXml_11_0_2 `xml:"ula,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaXml_11_0_2 struct {
	Advertise         *Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	EnableOnInterface *string                                            `xml:"enable-on-interface,omitempty"`
	PoolType          *Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2  `xml:"pool-type,omitempty"`
	PrefixPool        *string                                            `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 struct {
	AutoConfigFlag *string `xml:"auto-config-flag,omitempty"`
	Enable         *string `xml:"enable,omitempty"`
	OnlinkFlag     *string `xml:"onlink-flag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2 struct {
	Dynamic   *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2   `xml:"dynamic,omitempty"`
	DynamicId *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 `xml:"dynamic-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 struct {
	Identifier *int64 `xml:"identifier,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeUlaXml_11_0_2 struct {
	Address           *string                                            `xml:"address,omitempty"`
	Advertise         *Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Anycast           *string                                            `xml:"anycast,omitempty"`
	EnableOnInterface *string                                            `xml:"enable-on-interface,omitempty"`
	Prefix            *string                                            `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                                       `xml:"dad-attempts,omitempty"`
	DnsServer           *Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2           `xml:"dns-server,omitempty"`
	DnsSuffix           *Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                                      `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                      `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                                       `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                       `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"server>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable *string                                                  `xml:"enable,omitempty"`
	Source *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6 *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2 struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
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
type Ipv6NeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                              `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                             `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                             `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            []Ipv6NeighborDiscoveryNeighborXml_11_0_2           `xml:"neighbor>entry,omitempty"`
	NsInterval          *int64                                              `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                              `xml:"reachable-time,omitempty"`
	RouterAdvertisement *Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryNeighborXml_11_0_2 struct {
	HwAddress *string `xml:"hw-address,omitempty"`
	Name      string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
	DnsSupport             *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 `xml:"dns-support,omitempty"`
	Enable                 *string                                                       `xml:"enable,omitempty"`
	EnableConsistencyCheck *string                                                       `xml:"enable-consistency-check,omitempty"`
	HopLimit               *string                                                       `xml:"hop-limit,omitempty"`
	Lifetime               *int64                                                        `xml:"lifetime,omitempty"`
	LinkMtu                *string                                                       `xml:"link-mtu,omitempty"`
	ManagedFlag            *string                                                       `xml:"managed-flag,omitempty"`
	MaxInterval            *int64                                                        `xml:"max-interval,omitempty"`
	MinInterval            *int64                                                        `xml:"min-interval,omitempty"`
	OtherFlag              *string                                                       `xml:"other-flag,omitempty"`
	ReachableTime          *string                                                       `xml:"reachable-time,omitempty"`
	RetransmissionTimer    *string                                                       `xml:"retransmission-timer,omitempty"`
	RouterPreference       *string                                                       `xml:"router-preference,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 struct {
	Enable *string                                                              `xml:"enable,omitempty"`
	Server []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 `xml:"server>entry,omitempty"`
	Suffix []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 struct {
	Lifetime *int64 `xml:"lifetime,omitempty"`
	Name     string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type NdpProxyXml_11_0_2 struct {
	Address []NdpProxyAddressXml_11_0_2 `xml:"address>entry,omitempty"`
	Enabled *string                     `xml:"enabled,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NdpProxyAddressXml_11_0_2 struct {
	Name   string  `xml:"name,attr"`
	Negate *string `xml:"negate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "adjust_tcp_mss" || v == "AdjustTcpMss" {
		return e.AdjustTcpMss, nil
	}
	if v == "arp" || v == "Arp" {
		return e.Arp, nil
	}
	if v == "arp|LENGTH" || v == "Arp|LENGTH" {
		return int64(len(e.Arp)), nil
	}
	if v == "bonjour" || v == "Bonjour" {
		return e.Bonjour, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "ddns_config" || v == "DdnsConfig" {
		return e.DdnsConfig, nil
	}
	if v == "df_ignore" || v == "DfIgnore" {
		return e.DfIgnore, nil
	}
	if v == "dhcp_client" || v == "DhcpClient" {
		return e.DhcpClient, nil
	}
	if v == "interface_management_profile" || v == "InterfaceManagementProfile" {
		return e.InterfaceManagementProfile, nil
	}
	if v == "ip" || v == "Ip" {
		return e.Ip, nil
	}
	if v == "ip|LENGTH" || v == "Ip|LENGTH" {
		return int64(len(e.Ip)), nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "mtu" || v == "Mtu" {
		return e.Mtu, nil
	}
	if v == "ndp_proxy" || v == "NdpProxy" {
		return e.NdpProxy, nil
	}
	if v == "netflow_profile" || v == "NetflowProfile" {
		return e.NetflowProfile, nil
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
	var nestedAdjustTcpMss *AdjustTcpMssXml
	if o.AdjustTcpMss != nil {
		nestedAdjustTcpMss = &AdjustTcpMssXml{}
		if _, ok := o.Misc["AdjustTcpMss"]; ok {
			nestedAdjustTcpMss.Misc = o.Misc["AdjustTcpMss"]
		}
		if o.AdjustTcpMss.Enable != nil {
			nestedAdjustTcpMss.Enable = util.YesNo(o.AdjustTcpMss.Enable, nil)
		}
		if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
		}
		if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
		}
	}
	entry.AdjustTcpMss = nestedAdjustTcpMss

	var nestedArpCol []ArpXml
	if o.Arp != nil {
		nestedArpCol = []ArpXml{}
		for _, oArp := range o.Arp {
			nestedArp := ArpXml{}
			if _, ok := o.Misc["Arp"]; ok {
				nestedArp.Misc = o.Misc["Arp"]
			}
			if oArp.Name != "" {
				nestedArp.Name = oArp.Name
			}
			if oArp.HwAddress != nil {
				nestedArp.HwAddress = oArp.HwAddress
			}
			if oArp.Interface != nil {
				nestedArp.Interface = oArp.Interface
			}
			nestedArpCol = append(nestedArpCol, nestedArp)
		}
		entry.Arp = nestedArpCol
	}

	var nestedBonjour *BonjourXml
	if o.Bonjour != nil {
		nestedBonjour = &BonjourXml{}
		if _, ok := o.Misc["Bonjour"]; ok {
			nestedBonjour.Misc = o.Misc["Bonjour"]
		}
		if o.Bonjour.Enable != nil {
			nestedBonjour.Enable = util.YesNo(o.Bonjour.Enable, nil)
		}
		if o.Bonjour.GroupId != nil {
			nestedBonjour.GroupId = o.Bonjour.GroupId
		}
		if o.Bonjour.TtlCheck != nil {
			nestedBonjour.TtlCheck = util.YesNo(o.Bonjour.TtlCheck, nil)
		}
	}
	entry.Bonjour = nestedBonjour

	entry.Comment = o.Comment
	var nestedDdnsConfig *DdnsConfigXml
	if o.DdnsConfig != nil {
		nestedDdnsConfig = &DdnsConfigXml{}
		if _, ok := o.Misc["DdnsConfig"]; ok {
			nestedDdnsConfig.Misc = o.Misc["DdnsConfig"]
		}
		if o.DdnsConfig.DdnsCertProfile != nil {
			nestedDdnsConfig.DdnsCertProfile = o.DdnsConfig.DdnsCertProfile
		}
		if o.DdnsConfig.DdnsEnabled != nil {
			nestedDdnsConfig.DdnsEnabled = util.YesNo(o.DdnsConfig.DdnsEnabled, nil)
		}
		if o.DdnsConfig.DdnsHostname != nil {
			nestedDdnsConfig.DdnsHostname = o.DdnsConfig.DdnsHostname
		}
		if o.DdnsConfig.DdnsIp != nil {
			nestedDdnsConfig.DdnsIp = util.StrToMem(o.DdnsConfig.DdnsIp)
		}
		if o.DdnsConfig.DdnsIpv6 != nil {
			nestedDdnsConfig.DdnsIpv6 = util.StrToMem(o.DdnsConfig.DdnsIpv6)
		}
		if o.DdnsConfig.DdnsUpdateInterval != nil {
			nestedDdnsConfig.DdnsUpdateInterval = o.DdnsConfig.DdnsUpdateInterval
		}
		if o.DdnsConfig.DdnsVendor != nil {
			nestedDdnsConfig.DdnsVendor = o.DdnsConfig.DdnsVendor
		}
		if o.DdnsConfig.DdnsVendorConfig != nil {
			nestedDdnsConfig.DdnsVendorConfig = []DdnsConfigDdnsVendorConfigXml{}
			for _, oDdnsConfigDdnsVendorConfig := range o.DdnsConfig.DdnsVendorConfig {
				nestedDdnsConfigDdnsVendorConfig := DdnsConfigDdnsVendorConfigXml{}
				if _, ok := o.Misc["DdnsConfigDdnsVendorConfig"]; ok {
					nestedDdnsConfigDdnsVendorConfig.Misc = o.Misc["DdnsConfigDdnsVendorConfig"]
				}
				if oDdnsConfigDdnsVendorConfig.Name != "" {
					nestedDdnsConfigDdnsVendorConfig.Name = oDdnsConfigDdnsVendorConfig.Name
				}
				if oDdnsConfigDdnsVendorConfig.Value != nil {
					nestedDdnsConfigDdnsVendorConfig.Value = oDdnsConfigDdnsVendorConfig.Value
				}
				nestedDdnsConfig.DdnsVendorConfig = append(nestedDdnsConfig.DdnsVendorConfig, nestedDdnsConfigDdnsVendorConfig)
			}
		}
	}
	entry.DdnsConfig = nestedDdnsConfig

	entry.DfIgnore = util.YesNo(o.DfIgnore, nil)
	var nestedDhcpClient *DhcpClientXml
	if o.DhcpClient != nil {
		nestedDhcpClient = &DhcpClientXml{}
		if _, ok := o.Misc["DhcpClient"]; ok {
			nestedDhcpClient.Misc = o.Misc["DhcpClient"]
		}
		if o.DhcpClient.CreateDefaultRoute != nil {
			nestedDhcpClient.CreateDefaultRoute = util.YesNo(o.DhcpClient.CreateDefaultRoute, nil)
		}
		if o.DhcpClient.DefaultRouteMetric != nil {
			nestedDhcpClient.DefaultRouteMetric = o.DhcpClient.DefaultRouteMetric
		}
		if o.DhcpClient.Enable != nil {
			nestedDhcpClient.Enable = util.YesNo(o.DhcpClient.Enable, nil)
		}
		if o.DhcpClient.SendHostname != nil {
			nestedDhcpClient.SendHostname = &DhcpClientSendHostnameXml{}
			if _, ok := o.Misc["DhcpClientSendHostname"]; ok {
				nestedDhcpClient.SendHostname.Misc = o.Misc["DhcpClientSendHostname"]
			}
			if o.DhcpClient.SendHostname.Enable != nil {
				nestedDhcpClient.SendHostname.Enable = util.YesNo(o.DhcpClient.SendHostname.Enable, nil)
			}
			if o.DhcpClient.SendHostname.Hostname != nil {
				nestedDhcpClient.SendHostname.Hostname = o.DhcpClient.SendHostname.Hostname
			}
		}
	}
	entry.DhcpClient = nestedDhcpClient

	entry.InterfaceManagementProfile = o.InterfaceManagementProfile
	var nestedIpCol []IpXml
	if o.Ip != nil {
		nestedIpCol = []IpXml{}
		for _, oIp := range o.Ip {
			nestedIp := IpXml{}
			if _, ok := o.Misc["Ip"]; ok {
				nestedIp.Misc = o.Misc["Ip"]
			}
			if oIp.Name != "" {
				nestedIp.Name = oIp.Name
			}
			nestedIpCol = append(nestedIpCol, nestedIp)
		}
		entry.Ip = nestedIpCol
	}

	var nestedIpv6 *Ipv6Xml
	if o.Ipv6 != nil {
		nestedIpv6 = &Ipv6Xml{}
		if _, ok := o.Misc["Ipv6"]; ok {
			nestedIpv6.Misc = o.Misc["Ipv6"]
		}
		if o.Ipv6.Address != nil {
			nestedIpv6.Address = []Ipv6AddressXml{}
			for _, oIpv6Address := range o.Ipv6.Address {
				nestedIpv6Address := Ipv6AddressXml{}
				if _, ok := o.Misc["Ipv6Address"]; ok {
					nestedIpv6Address.Misc = o.Misc["Ipv6Address"]
				}
				if oIpv6Address.Name != "" {
					nestedIpv6Address.Name = oIpv6Address.Name
				}
				if oIpv6Address.EnableOnInterface != nil {
					nestedIpv6Address.EnableOnInterface = util.YesNo(oIpv6Address.EnableOnInterface, nil)
				}
				if oIpv6Address.Prefix != nil {
					nestedIpv6Address.Prefix = &Ipv6AddressPrefixXml{}
					if _, ok := o.Misc["Ipv6AddressPrefix"]; ok {
						nestedIpv6Address.Prefix.Misc = o.Misc["Ipv6AddressPrefix"]
					}
				}
				if oIpv6Address.Anycast != nil {
					nestedIpv6Address.Anycast = &Ipv6AddressAnycastXml{}
					if _, ok := o.Misc["Ipv6AddressAnycast"]; ok {
						nestedIpv6Address.Anycast.Misc = o.Misc["Ipv6AddressAnycast"]
					}
				}
				if oIpv6Address.Advertise != nil {
					nestedIpv6Address.Advertise = &Ipv6AddressAdvertiseXml{}
					if _, ok := o.Misc["Ipv6AddressAdvertise"]; ok {
						nestedIpv6Address.Advertise.Misc = o.Misc["Ipv6AddressAdvertise"]
					}
					if oIpv6Address.Advertise.Enable != nil {
						nestedIpv6Address.Advertise.Enable = util.YesNo(oIpv6Address.Advertise.Enable, nil)
					}
					if oIpv6Address.Advertise.ValidLifetime != nil {
						nestedIpv6Address.Advertise.ValidLifetime = oIpv6Address.Advertise.ValidLifetime
					}
					if oIpv6Address.Advertise.PreferredLifetime != nil {
						nestedIpv6Address.Advertise.PreferredLifetime = oIpv6Address.Advertise.PreferredLifetime
					}
					if oIpv6Address.Advertise.OnlinkFlag != nil {
						nestedIpv6Address.Advertise.OnlinkFlag = util.YesNo(oIpv6Address.Advertise.OnlinkFlag, nil)
					}
					if oIpv6Address.Advertise.AutoConfigFlag != nil {
						nestedIpv6Address.Advertise.AutoConfigFlag = util.YesNo(oIpv6Address.Advertise.AutoConfigFlag, nil)
					}
				}
				nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
			}
		}
		if o.Ipv6.Enabled != nil {
			nestedIpv6.Enabled = util.YesNo(o.Ipv6.Enabled, nil)
		}
		if o.Ipv6.InterfaceId != nil {
			nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
		}
		if o.Ipv6.NeighborDiscovery != nil {
			nestedIpv6.NeighborDiscovery = &Ipv6NeighborDiscoveryXml{}
			if _, ok := o.Misc["Ipv6NeighborDiscovery"]; ok {
				nestedIpv6.NeighborDiscovery.Misc = o.Misc["Ipv6NeighborDiscovery"]
			}
			if o.Ipv6.NeighborDiscovery.DadAttempts != nil {
				nestedIpv6.NeighborDiscovery.DadAttempts = o.Ipv6.NeighborDiscovery.DadAttempts
			}
			if o.Ipv6.NeighborDiscovery.EnableDad != nil {
				nestedIpv6.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.NeighborDiscovery.EnableDad, nil)
			}
			if o.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
				nestedIpv6.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
			}
			if o.Ipv6.NeighborDiscovery.Neighbor != nil {
				nestedIpv6.NeighborDiscovery.Neighbor = []Ipv6NeighborDiscoveryNeighborXml{}
				for _, oIpv6NeighborDiscoveryNeighbor := range o.Ipv6.NeighborDiscovery.Neighbor {
					nestedIpv6NeighborDiscoveryNeighbor := Ipv6NeighborDiscoveryNeighborXml{}
					if _, ok := o.Misc["Ipv6NeighborDiscoveryNeighbor"]; ok {
						nestedIpv6NeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6NeighborDiscoveryNeighbor"]
					}
					if oIpv6NeighborDiscoveryNeighbor.Name != "" {
						nestedIpv6NeighborDiscoveryNeighbor.Name = oIpv6NeighborDiscoveryNeighbor.Name
					}
					if oIpv6NeighborDiscoveryNeighbor.HwAddress != nil {
						nestedIpv6NeighborDiscoveryNeighbor.HwAddress = oIpv6NeighborDiscoveryNeighbor.HwAddress
					}
					nestedIpv6.NeighborDiscovery.Neighbor = append(nestedIpv6.NeighborDiscovery.Neighbor, nestedIpv6NeighborDiscoveryNeighbor)
				}
			}
			if o.Ipv6.NeighborDiscovery.NsInterval != nil {
				nestedIpv6.NeighborDiscovery.NsInterval = o.Ipv6.NeighborDiscovery.NsInterval
			}
			if o.Ipv6.NeighborDiscovery.ReachableTime != nil {
				nestedIpv6.NeighborDiscovery.ReachableTime = o.Ipv6.NeighborDiscovery.ReachableTime
			}
			if o.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
				nestedIpv6.NeighborDiscovery.RouterAdvertisement = &Ipv6NeighborDiscoveryRouterAdvertisementXml{}
				if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"]; ok {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"]
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml{}
					if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]; ok {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml{}
						for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
							nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml{}
							if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]; ok {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
							}
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
						}
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml{}
						for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
							nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml{}
							if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]; ok {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
							}
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
						}
					}
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
				}
			}
		}
		if o.Ipv6.DhcpClient != nil {
			nestedIpv6.DhcpClient = &Ipv6DhcpClientXml{}
			if _, ok := o.Misc["Ipv6DhcpClient"]; ok {
				nestedIpv6.DhcpClient.Misc = o.Misc["Ipv6DhcpClient"]
			}
			if o.Ipv6.DhcpClient.AcceptRaRoute != nil {
				nestedIpv6.DhcpClient.AcceptRaRoute = util.YesNo(o.Ipv6.DhcpClient.AcceptRaRoute, nil)
			}
			if o.Ipv6.DhcpClient.DefaultRouteMetric != nil {
				nestedIpv6.DhcpClient.DefaultRouteMetric = o.Ipv6.DhcpClient.DefaultRouteMetric
			}
			if o.Ipv6.DhcpClient.Enable != nil {
				nestedIpv6.DhcpClient.Enable = util.YesNo(o.Ipv6.DhcpClient.Enable, nil)
			}
			if o.Ipv6.DhcpClient.NeighborDiscovery != nil {
				nestedIpv6.DhcpClient.NeighborDiscovery = &Ipv6DhcpClientNeighborDiscoveryXml{}
				if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscovery"]; ok {
					nestedIpv6.DhcpClient.NeighborDiscovery.Misc = o.Misc["Ipv6DhcpClientNeighborDiscovery"]
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer = &Ipv6DhcpClientNeighborDiscoveryDnsServerXml{}
					if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"]; ok {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"]
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]; ok {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]
							}
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml{}
								for _, oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
									nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml{}
									if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]; ok {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
									}
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
								}
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml{}
					if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]; ok {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]; ok {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]
							}
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
								for _, oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
									nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
									if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
									}
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
								}
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = []Ipv6DhcpClientNeighborDiscoveryNeighborXml{}
					for _, oIpv6DhcpClientNeighborDiscoveryNeighbor := range o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
						nestedIpv6DhcpClientNeighborDiscoveryNeighbor := Ipv6DhcpClientNeighborDiscoveryNeighborXml{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"]; ok {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"]
						}
						if oIpv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Name = oIpv6DhcpClientNeighborDiscoveryNeighbor.Name
						}
						if oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor, nestedIpv6DhcpClientNeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.NsInterval = o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
				}
			}
			if o.Ipv6.DhcpClient.Preference != nil {
				nestedIpv6.DhcpClient.Preference = o.Ipv6.DhcpClient.Preference
			}
			if o.Ipv6.DhcpClient.PrefixDelegation != nil {
				nestedIpv6.DhcpClient.PrefixDelegation = &Ipv6DhcpClientPrefixDelegationXml{}
				if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegation"]; ok {
					nestedIpv6.DhcpClient.PrefixDelegation.Misc = o.Misc["Ipv6DhcpClientPrefixDelegation"]
				}
				if o.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
					nestedIpv6.DhcpClient.PrefixDelegation.Enable = &Ipv6DhcpClientPrefixDelegationEnableXml{}
					if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnable"]; ok {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnable"]
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.No = &Ipv6DhcpClientPrefixDelegationEnableNoXml{}
						if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"]; ok {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.No.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"]
						}
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes = &Ipv6DhcpClientPrefixDelegationEnableYesXml{}
						if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"]; ok {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"]
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.YesNo(o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
						}
					}
				}
			}
			if o.Ipv6.DhcpClient.V6Options != nil {
				nestedIpv6.DhcpClient.V6Options = &Ipv6DhcpClientV6OptionsXml{}
				if _, ok := o.Misc["Ipv6DhcpClientV6Options"]; ok {
					nestedIpv6.DhcpClient.V6Options.Misc = o.Misc["Ipv6DhcpClientV6Options"]
				}
				if o.Ipv6.DhcpClient.V6Options.DuidType != nil {
					nestedIpv6.DhcpClient.V6Options.DuidType = o.Ipv6.DhcpClient.V6Options.DuidType
				}
				if o.Ipv6.DhcpClient.V6Options.Enable != nil {
					nestedIpv6.DhcpClient.V6Options.Enable = &Ipv6DhcpClientV6OptionsEnableXml{}
					if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnable"]; ok {
						nestedIpv6.DhcpClient.V6Options.Enable.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnable"]
					}
					if o.Ipv6.DhcpClient.V6Options.Enable.No != nil {
						nestedIpv6.DhcpClient.V6Options.Enable.No = &Ipv6DhcpClientV6OptionsEnableNoXml{}
						if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnableNo"]; ok {
							nestedIpv6.DhcpClient.V6Options.Enable.No.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnableNo"]
						}
					}
					if o.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
						nestedIpv6.DhcpClient.V6Options.Enable.Yes = &Ipv6DhcpClientV6OptionsEnableYesXml{}
						if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnableYes"]; ok {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnableYes"]
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.YesNo(o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.YesNo(o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
						}
					}
				}
				if o.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
					nestedIpv6.DhcpClient.V6Options.RapidCommit = util.YesNo(o.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
				}
				if o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
					nestedIpv6.DhcpClient.V6Options.SupportSrvrReconfig = util.YesNo(o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
				}
			}
		}
		if o.Ipv6.Inherited != nil {
			nestedIpv6.Inherited = &Ipv6InheritedXml{}
			if _, ok := o.Misc["Ipv6Inherited"]; ok {
				nestedIpv6.Inherited.Misc = o.Misc["Ipv6Inherited"]
			}
			if o.Ipv6.Inherited.AssignAddr != nil {
				nestedIpv6.Inherited.AssignAddr = []Ipv6InheritedAssignAddrXml{}
				for _, oIpv6InheritedAssignAddr := range o.Ipv6.Inherited.AssignAddr {
					nestedIpv6InheritedAssignAddr := Ipv6InheritedAssignAddrXml{}
					if _, ok := o.Misc["Ipv6InheritedAssignAddr"]; ok {
						nestedIpv6InheritedAssignAddr.Misc = o.Misc["Ipv6InheritedAssignAddr"]
					}
					if oIpv6InheritedAssignAddr.Name != "" {
						nestedIpv6InheritedAssignAddr.Name = oIpv6InheritedAssignAddr.Name
					}
					if oIpv6InheritedAssignAddr.Type != nil {
						nestedIpv6InheritedAssignAddr.Type = &Ipv6InheritedAssignAddrTypeXml{}
						if _, ok := o.Misc["Ipv6InheritedAssignAddrType"]; ok {
							nestedIpv6InheritedAssignAddr.Type.Misc = o.Misc["Ipv6InheritedAssignAddrType"]
						}
						if oIpv6InheritedAssignAddr.Type.Gua != nil {
							nestedIpv6InheritedAssignAddr.Type.Gua = &Ipv6InheritedAssignAddrTypeGuaXml{}
							if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGua"]; ok {
								nestedIpv6InheritedAssignAddr.Type.Gua.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGua"]
							}
							if oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.PrefixPool = oIpv6InheritedAssignAddr.Type.Gua.PrefixPool
							}
							if oIpv6InheritedAssignAddr.Type.Gua.PoolType != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.PoolType = &Ipv6InheritedAssignAddrTypeGuaPoolTypeXml{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"]
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml{}
									if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]; ok {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]
									}
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml{}
									if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]; ok {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
									}
								}
							}
							if oIpv6InheritedAssignAddr.Type.Gua.Advertise != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.Advertise = &Ipv6InheritedAssignAddrTypeGuaAdvertiseXml{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"]
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
								}
							}
						}
						if oIpv6InheritedAssignAddr.Type.Ula != nil {
							nestedIpv6InheritedAssignAddr.Type.Ula = &Ipv6InheritedAssignAddrTypeUlaXml{}
							if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeUla"]; ok {
								nestedIpv6InheritedAssignAddr.Type.Ula.Misc = o.Misc["Ipv6InheritedAssignAddrTypeUla"]
							}
							if oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Address != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Address = oIpv6InheritedAssignAddr.Type.Ula.Address
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Prefix != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Prefix = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Prefix, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Anycast != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Anycast = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Anycast, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Advertise != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Advertise = &Ipv6InheritedAssignAddrTypeUlaAdvertiseXml{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc = o.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"]
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
								}
							}
						}
					}
					nestedIpv6.Inherited.AssignAddr = append(nestedIpv6.Inherited.AssignAddr, nestedIpv6InheritedAssignAddr)
				}
			}
			if o.Ipv6.Inherited.Enable != nil {
				nestedIpv6.Inherited.Enable = util.YesNo(o.Ipv6.Inherited.Enable, nil)
			}
			if o.Ipv6.Inherited.NeighborDiscovery != nil {
				nestedIpv6.Inherited.NeighborDiscovery = &Ipv6InheritedNeighborDiscoveryXml{}
				if _, ok := o.Misc["Ipv6InheritedNeighborDiscovery"]; ok {
					nestedIpv6.Inherited.NeighborDiscovery.Misc = o.Misc["Ipv6InheritedNeighborDiscovery"]
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DadAttempts = o.Ipv6.Inherited.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DnsServer = &Ipv6InheritedNeighborDiscoveryDnsServerXml{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source = &Ipv6InheritedNeighborDiscoveryDnsServerSourceXml{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"]; ok {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"]
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
							}
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml{}
								for _, oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
									nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml{}
									if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]; ok {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]
									}
									if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
									}
									if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
									}
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
								}
							}
						}
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix = &Ipv6InheritedNeighborDiscoveryDnsSuffixXml{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]; ok {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
							}
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
								for _, oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
									nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml{}
									if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]
									}
									if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
									}
									if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
									}
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
								}
							}
						}
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.Inherited.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.Inherited.NeighborDiscovery.Neighbor = []Ipv6InheritedNeighborDiscoveryNeighborXml{}
					for _, oIpv6InheritedNeighborDiscoveryNeighbor := range o.Ipv6.Inherited.NeighborDiscovery.Neighbor {
						nestedIpv6InheritedNeighborDiscoveryNeighbor := Ipv6InheritedNeighborDiscoveryNeighborXml{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"]; ok {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"]
						}
						if oIpv6InheritedNeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.Name = oIpv6InheritedNeighborDiscoveryNeighbor.Name
						}
						if oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.HwAddress = oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.Inherited.NeighborDiscovery.Neighbor = append(nestedIpv6.Inherited.NeighborDiscovery.Neighbor, nestedIpv6InheritedNeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.Inherited.NeighborDiscovery.NsInterval = o.Ipv6.Inherited.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.Inherited.NeighborDiscovery.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.ReachableTime
				}
				if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
					nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
				}
			}
		}
	}
	entry.Ipv6 = nestedIpv6

	entry.Mtu = o.Mtu
	var nestedNdpProxy *NdpProxyXml
	if o.NdpProxy != nil {
		nestedNdpProxy = &NdpProxyXml{}
		if _, ok := o.Misc["NdpProxy"]; ok {
			nestedNdpProxy.Misc = o.Misc["NdpProxy"]
		}
		if o.NdpProxy.Address != nil {
			nestedNdpProxy.Address = []NdpProxyAddressXml{}
			for _, oNdpProxyAddress := range o.NdpProxy.Address {
				nestedNdpProxyAddress := NdpProxyAddressXml{}
				if _, ok := o.Misc["NdpProxyAddress"]; ok {
					nestedNdpProxyAddress.Misc = o.Misc["NdpProxyAddress"]
				}
				if oNdpProxyAddress.Name != "" {
					nestedNdpProxyAddress.Name = oNdpProxyAddress.Name
				}
				if oNdpProxyAddress.Negate != nil {
					nestedNdpProxyAddress.Negate = util.YesNo(oNdpProxyAddress.Negate, nil)
				}
				nestedNdpProxy.Address = append(nestedNdpProxy.Address, nestedNdpProxyAddress)
			}
		}
		if o.NdpProxy.Enabled != nil {
			nestedNdpProxy.Enabled = util.YesNo(o.NdpProxy.Enabled, nil)
		}
	}
	entry.NdpProxy = nestedNdpProxy

	entry.NetflowProfile = o.NetflowProfile

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	var nestedAdjustTcpMss *AdjustTcpMssXml_11_0_2
	if o.AdjustTcpMss != nil {
		nestedAdjustTcpMss = &AdjustTcpMssXml_11_0_2{}
		if _, ok := o.Misc["AdjustTcpMss"]; ok {
			nestedAdjustTcpMss.Misc = o.Misc["AdjustTcpMss"]
		}
		if o.AdjustTcpMss.Enable != nil {
			nestedAdjustTcpMss.Enable = util.YesNo(o.AdjustTcpMss.Enable, nil)
		}
		if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
		}
		if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
		}
	}
	entry.AdjustTcpMss = nestedAdjustTcpMss

	var nestedArpCol []ArpXml_11_0_2
	if o.Arp != nil {
		nestedArpCol = []ArpXml_11_0_2{}
		for _, oArp := range o.Arp {
			nestedArp := ArpXml_11_0_2{}
			if _, ok := o.Misc["Arp"]; ok {
				nestedArp.Misc = o.Misc["Arp"]
			}
			if oArp.Name != "" {
				nestedArp.Name = oArp.Name
			}
			if oArp.HwAddress != nil {
				nestedArp.HwAddress = oArp.HwAddress
			}
			if oArp.Interface != nil {
				nestedArp.Interface = oArp.Interface
			}
			nestedArpCol = append(nestedArpCol, nestedArp)
		}
		entry.Arp = nestedArpCol
	}

	var nestedBonjour *BonjourXml_11_0_2
	if o.Bonjour != nil {
		nestedBonjour = &BonjourXml_11_0_2{}
		if _, ok := o.Misc["Bonjour"]; ok {
			nestedBonjour.Misc = o.Misc["Bonjour"]
		}
		if o.Bonjour.Enable != nil {
			nestedBonjour.Enable = util.YesNo(o.Bonjour.Enable, nil)
		}
		if o.Bonjour.GroupId != nil {
			nestedBonjour.GroupId = o.Bonjour.GroupId
		}
		if o.Bonjour.TtlCheck != nil {
			nestedBonjour.TtlCheck = util.YesNo(o.Bonjour.TtlCheck, nil)
		}
	}
	entry.Bonjour = nestedBonjour

	entry.Comment = o.Comment
	var nestedDdnsConfig *DdnsConfigXml_11_0_2
	if o.DdnsConfig != nil {
		nestedDdnsConfig = &DdnsConfigXml_11_0_2{}
		if _, ok := o.Misc["DdnsConfig"]; ok {
			nestedDdnsConfig.Misc = o.Misc["DdnsConfig"]
		}
		if o.DdnsConfig.DdnsCertProfile != nil {
			nestedDdnsConfig.DdnsCertProfile = o.DdnsConfig.DdnsCertProfile
		}
		if o.DdnsConfig.DdnsEnabled != nil {
			nestedDdnsConfig.DdnsEnabled = util.YesNo(o.DdnsConfig.DdnsEnabled, nil)
		}
		if o.DdnsConfig.DdnsHostname != nil {
			nestedDdnsConfig.DdnsHostname = o.DdnsConfig.DdnsHostname
		}
		if o.DdnsConfig.DdnsIp != nil {
			nestedDdnsConfig.DdnsIp = util.StrToMem(o.DdnsConfig.DdnsIp)
		}
		if o.DdnsConfig.DdnsIpv6 != nil {
			nestedDdnsConfig.DdnsIpv6 = util.StrToMem(o.DdnsConfig.DdnsIpv6)
		}
		if o.DdnsConfig.DdnsUpdateInterval != nil {
			nestedDdnsConfig.DdnsUpdateInterval = o.DdnsConfig.DdnsUpdateInterval
		}
		if o.DdnsConfig.DdnsVendor != nil {
			nestedDdnsConfig.DdnsVendor = o.DdnsConfig.DdnsVendor
		}
		if o.DdnsConfig.DdnsVendorConfig != nil {
			nestedDdnsConfig.DdnsVendorConfig = []DdnsConfigDdnsVendorConfigXml_11_0_2{}
			for _, oDdnsConfigDdnsVendorConfig := range o.DdnsConfig.DdnsVendorConfig {
				nestedDdnsConfigDdnsVendorConfig := DdnsConfigDdnsVendorConfigXml_11_0_2{}
				if _, ok := o.Misc["DdnsConfigDdnsVendorConfig"]; ok {
					nestedDdnsConfigDdnsVendorConfig.Misc = o.Misc["DdnsConfigDdnsVendorConfig"]
				}
				if oDdnsConfigDdnsVendorConfig.Name != "" {
					nestedDdnsConfigDdnsVendorConfig.Name = oDdnsConfigDdnsVendorConfig.Name
				}
				if oDdnsConfigDdnsVendorConfig.Value != nil {
					nestedDdnsConfigDdnsVendorConfig.Value = oDdnsConfigDdnsVendorConfig.Value
				}
				nestedDdnsConfig.DdnsVendorConfig = append(nestedDdnsConfig.DdnsVendorConfig, nestedDdnsConfigDdnsVendorConfig)
			}
		}
	}
	entry.DdnsConfig = nestedDdnsConfig

	entry.DfIgnore = util.YesNo(o.DfIgnore, nil)
	var nestedDhcpClient *DhcpClientXml_11_0_2
	if o.DhcpClient != nil {
		nestedDhcpClient = &DhcpClientXml_11_0_2{}
		if _, ok := o.Misc["DhcpClient"]; ok {
			nestedDhcpClient.Misc = o.Misc["DhcpClient"]
		}
		if o.DhcpClient.CreateDefaultRoute != nil {
			nestedDhcpClient.CreateDefaultRoute = util.YesNo(o.DhcpClient.CreateDefaultRoute, nil)
		}
		if o.DhcpClient.DefaultRouteMetric != nil {
			nestedDhcpClient.DefaultRouteMetric = o.DhcpClient.DefaultRouteMetric
		}
		if o.DhcpClient.Enable != nil {
			nestedDhcpClient.Enable = util.YesNo(o.DhcpClient.Enable, nil)
		}
		if o.DhcpClient.SendHostname != nil {
			nestedDhcpClient.SendHostname = &DhcpClientSendHostnameXml_11_0_2{}
			if _, ok := o.Misc["DhcpClientSendHostname"]; ok {
				nestedDhcpClient.SendHostname.Misc = o.Misc["DhcpClientSendHostname"]
			}
			if o.DhcpClient.SendHostname.Enable != nil {
				nestedDhcpClient.SendHostname.Enable = util.YesNo(o.DhcpClient.SendHostname.Enable, nil)
			}
			if o.DhcpClient.SendHostname.Hostname != nil {
				nestedDhcpClient.SendHostname.Hostname = o.DhcpClient.SendHostname.Hostname
			}
		}
	}
	entry.DhcpClient = nestedDhcpClient

	entry.InterfaceManagementProfile = o.InterfaceManagementProfile
	var nestedIpCol []IpXml_11_0_2
	if o.Ip != nil {
		nestedIpCol = []IpXml_11_0_2{}
		for _, oIp := range o.Ip {
			nestedIp := IpXml_11_0_2{}
			if _, ok := o.Misc["Ip"]; ok {
				nestedIp.Misc = o.Misc["Ip"]
			}
			if oIp.Name != "" {
				nestedIp.Name = oIp.Name
			}
			nestedIpCol = append(nestedIpCol, nestedIp)
		}
		entry.Ip = nestedIpCol
	}

	var nestedIpv6 *Ipv6Xml_11_0_2
	if o.Ipv6 != nil {
		nestedIpv6 = &Ipv6Xml_11_0_2{}
		if _, ok := o.Misc["Ipv6"]; ok {
			nestedIpv6.Misc = o.Misc["Ipv6"]
		}
		if o.Ipv6.Address != nil {
			nestedIpv6.Address = []Ipv6AddressXml_11_0_2{}
			for _, oIpv6Address := range o.Ipv6.Address {
				nestedIpv6Address := Ipv6AddressXml_11_0_2{}
				if _, ok := o.Misc["Ipv6Address"]; ok {
					nestedIpv6Address.Misc = o.Misc["Ipv6Address"]
				}
				if oIpv6Address.Name != "" {
					nestedIpv6Address.Name = oIpv6Address.Name
				}
				if oIpv6Address.EnableOnInterface != nil {
					nestedIpv6Address.EnableOnInterface = util.YesNo(oIpv6Address.EnableOnInterface, nil)
				}
				if oIpv6Address.Prefix != nil {
					nestedIpv6Address.Prefix = &Ipv6AddressPrefixXml_11_0_2{}
					if _, ok := o.Misc["Ipv6AddressPrefix"]; ok {
						nestedIpv6Address.Prefix.Misc = o.Misc["Ipv6AddressPrefix"]
					}
				}
				if oIpv6Address.Anycast != nil {
					nestedIpv6Address.Anycast = &Ipv6AddressAnycastXml_11_0_2{}
					if _, ok := o.Misc["Ipv6AddressAnycast"]; ok {
						nestedIpv6Address.Anycast.Misc = o.Misc["Ipv6AddressAnycast"]
					}
				}
				if oIpv6Address.Advertise != nil {
					nestedIpv6Address.Advertise = &Ipv6AddressAdvertiseXml_11_0_2{}
					if _, ok := o.Misc["Ipv6AddressAdvertise"]; ok {
						nestedIpv6Address.Advertise.Misc = o.Misc["Ipv6AddressAdvertise"]
					}
					if oIpv6Address.Advertise.Enable != nil {
						nestedIpv6Address.Advertise.Enable = util.YesNo(oIpv6Address.Advertise.Enable, nil)
					}
					if oIpv6Address.Advertise.ValidLifetime != nil {
						nestedIpv6Address.Advertise.ValidLifetime = oIpv6Address.Advertise.ValidLifetime
					}
					if oIpv6Address.Advertise.PreferredLifetime != nil {
						nestedIpv6Address.Advertise.PreferredLifetime = oIpv6Address.Advertise.PreferredLifetime
					}
					if oIpv6Address.Advertise.OnlinkFlag != nil {
						nestedIpv6Address.Advertise.OnlinkFlag = util.YesNo(oIpv6Address.Advertise.OnlinkFlag, nil)
					}
					if oIpv6Address.Advertise.AutoConfigFlag != nil {
						nestedIpv6Address.Advertise.AutoConfigFlag = util.YesNo(oIpv6Address.Advertise.AutoConfigFlag, nil)
					}
				}
				nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
			}
		}
		if o.Ipv6.Enabled != nil {
			nestedIpv6.Enabled = util.YesNo(o.Ipv6.Enabled, nil)
		}
		if o.Ipv6.InterfaceId != nil {
			nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
		}
		if o.Ipv6.NeighborDiscovery != nil {
			nestedIpv6.NeighborDiscovery = &Ipv6NeighborDiscoveryXml_11_0_2{}
			if _, ok := o.Misc["Ipv6NeighborDiscovery"]; ok {
				nestedIpv6.NeighborDiscovery.Misc = o.Misc["Ipv6NeighborDiscovery"]
			}
			if o.Ipv6.NeighborDiscovery.DadAttempts != nil {
				nestedIpv6.NeighborDiscovery.DadAttempts = o.Ipv6.NeighborDiscovery.DadAttempts
			}
			if o.Ipv6.NeighborDiscovery.EnableDad != nil {
				nestedIpv6.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.NeighborDiscovery.EnableDad, nil)
			}
			if o.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
				nestedIpv6.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
			}
			if o.Ipv6.NeighborDiscovery.Neighbor != nil {
				nestedIpv6.NeighborDiscovery.Neighbor = []Ipv6NeighborDiscoveryNeighborXml_11_0_2{}
				for _, oIpv6NeighborDiscoveryNeighbor := range o.Ipv6.NeighborDiscovery.Neighbor {
					nestedIpv6NeighborDiscoveryNeighbor := Ipv6NeighborDiscoveryNeighborXml_11_0_2{}
					if _, ok := o.Misc["Ipv6NeighborDiscoveryNeighbor"]; ok {
						nestedIpv6NeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6NeighborDiscoveryNeighbor"]
					}
					if oIpv6NeighborDiscoveryNeighbor.Name != "" {
						nestedIpv6NeighborDiscoveryNeighbor.Name = oIpv6NeighborDiscoveryNeighbor.Name
					}
					if oIpv6NeighborDiscoveryNeighbor.HwAddress != nil {
						nestedIpv6NeighborDiscoveryNeighbor.HwAddress = oIpv6NeighborDiscoveryNeighbor.HwAddress
					}
					nestedIpv6.NeighborDiscovery.Neighbor = append(nestedIpv6.NeighborDiscovery.Neighbor, nestedIpv6NeighborDiscoveryNeighbor)
				}
			}
			if o.Ipv6.NeighborDiscovery.NsInterval != nil {
				nestedIpv6.NeighborDiscovery.NsInterval = o.Ipv6.NeighborDiscovery.NsInterval
			}
			if o.Ipv6.NeighborDiscovery.ReachableTime != nil {
				nestedIpv6.NeighborDiscovery.ReachableTime = o.Ipv6.NeighborDiscovery.ReachableTime
			}
			if o.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
				nestedIpv6.NeighborDiscovery.RouterAdvertisement = &Ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2{}
				if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"]; ok {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"]
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2{}
					if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]; ok {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"]
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2{}
						for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
							nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2{}
							if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]; ok {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"]
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
							}
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
						}
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2{}
						for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
							nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2{}
							if _, ok := o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]; ok {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc = o.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"]
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
							}
							if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
							}
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
						}
					}
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
				}
			}
		}
		if o.Ipv6.DhcpClient != nil {
			nestedIpv6.DhcpClient = &Ipv6DhcpClientXml_11_0_2{}
			if _, ok := o.Misc["Ipv6DhcpClient"]; ok {
				nestedIpv6.DhcpClient.Misc = o.Misc["Ipv6DhcpClient"]
			}
			if o.Ipv6.DhcpClient.AcceptRaRoute != nil {
				nestedIpv6.DhcpClient.AcceptRaRoute = util.YesNo(o.Ipv6.DhcpClient.AcceptRaRoute, nil)
			}
			if o.Ipv6.DhcpClient.DefaultRouteMetric != nil {
				nestedIpv6.DhcpClient.DefaultRouteMetric = o.Ipv6.DhcpClient.DefaultRouteMetric
			}
			if o.Ipv6.DhcpClient.Enable != nil {
				nestedIpv6.DhcpClient.Enable = util.YesNo(o.Ipv6.DhcpClient.Enable, nil)
			}
			if o.Ipv6.DhcpClient.NeighborDiscovery != nil {
				nestedIpv6.DhcpClient.NeighborDiscovery = &Ipv6DhcpClientNeighborDiscoveryXml_11_0_2{}
				if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscovery"]; ok {
					nestedIpv6.DhcpClient.NeighborDiscovery.Misc = o.Misc["Ipv6DhcpClientNeighborDiscovery"]
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer = &Ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2{}
					if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"]; ok {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"]
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]; ok {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"]
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"]
							}
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"]
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2{}
								for _, oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
									nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2{}
									if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]; ok {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"]
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
									}
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
								}
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2{}
					if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]; ok {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"]
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]; ok {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"]
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"]
							}
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2{}
							if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]; ok {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"]
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2{}
								for _, oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
									nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2{}
									if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"]
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
									}
									if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
									}
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
								}
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = []Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2{}
					for _, oIpv6DhcpClientNeighborDiscoveryNeighbor := range o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
						nestedIpv6DhcpClientNeighborDiscoveryNeighbor := Ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"]; ok {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"]
						}
						if oIpv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Name = oIpv6DhcpClientNeighborDiscoveryNeighbor.Name
						}
						if oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor, nestedIpv6DhcpClientNeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.NsInterval = o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
				}
			}
			if o.Ipv6.DhcpClient.Preference != nil {
				nestedIpv6.DhcpClient.Preference = o.Ipv6.DhcpClient.Preference
			}
			if o.Ipv6.DhcpClient.PrefixDelegation != nil {
				nestedIpv6.DhcpClient.PrefixDelegation = &Ipv6DhcpClientPrefixDelegationXml_11_0_2{}
				if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegation"]; ok {
					nestedIpv6.DhcpClient.PrefixDelegation.Misc = o.Misc["Ipv6DhcpClientPrefixDelegation"]
				}
				if o.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
					nestedIpv6.DhcpClient.PrefixDelegation.Enable = &Ipv6DhcpClientPrefixDelegationEnableXml_11_0_2{}
					if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnable"]; ok {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnable"]
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.No = &Ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"]; ok {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.No.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"]
						}
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes = &Ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"]; ok {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc = o.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"]
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.YesNo(o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
						}
					}
				}
			}
			if o.Ipv6.DhcpClient.V6Options != nil {
				nestedIpv6.DhcpClient.V6Options = &Ipv6DhcpClientV6OptionsXml_11_0_2{}
				if _, ok := o.Misc["Ipv6DhcpClientV6Options"]; ok {
					nestedIpv6.DhcpClient.V6Options.Misc = o.Misc["Ipv6DhcpClientV6Options"]
				}
				if o.Ipv6.DhcpClient.V6Options.DuidType != nil {
					nestedIpv6.DhcpClient.V6Options.DuidType = o.Ipv6.DhcpClient.V6Options.DuidType
				}
				if o.Ipv6.DhcpClient.V6Options.Enable != nil {
					nestedIpv6.DhcpClient.V6Options.Enable = &Ipv6DhcpClientV6OptionsEnableXml_11_0_2{}
					if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnable"]; ok {
						nestedIpv6.DhcpClient.V6Options.Enable.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnable"]
					}
					if o.Ipv6.DhcpClient.V6Options.Enable.No != nil {
						nestedIpv6.DhcpClient.V6Options.Enable.No = &Ipv6DhcpClientV6OptionsEnableNoXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnableNo"]; ok {
							nestedIpv6.DhcpClient.V6Options.Enable.No.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnableNo"]
						}
					}
					if o.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
						nestedIpv6.DhcpClient.V6Options.Enable.Yes = &Ipv6DhcpClientV6OptionsEnableYesXml_11_0_2{}
						if _, ok := o.Misc["Ipv6DhcpClientV6OptionsEnableYes"]; ok {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.Misc = o.Misc["Ipv6DhcpClientV6OptionsEnableYes"]
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.YesNo(o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.YesNo(o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
						}
					}
				}
				if o.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
					nestedIpv6.DhcpClient.V6Options.RapidCommit = util.YesNo(o.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
				}
				if o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
					nestedIpv6.DhcpClient.V6Options.SupportSrvrReconfig = util.YesNo(o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
				}
			}
		}
		if o.Ipv6.Inherited != nil {
			nestedIpv6.Inherited = &Ipv6InheritedXml_11_0_2{}
			if _, ok := o.Misc["Ipv6Inherited"]; ok {
				nestedIpv6.Inherited.Misc = o.Misc["Ipv6Inherited"]
			}
			if o.Ipv6.Inherited.AssignAddr != nil {
				nestedIpv6.Inherited.AssignAddr = []Ipv6InheritedAssignAddrXml_11_0_2{}
				for _, oIpv6InheritedAssignAddr := range o.Ipv6.Inherited.AssignAddr {
					nestedIpv6InheritedAssignAddr := Ipv6InheritedAssignAddrXml_11_0_2{}
					if _, ok := o.Misc["Ipv6InheritedAssignAddr"]; ok {
						nestedIpv6InheritedAssignAddr.Misc = o.Misc["Ipv6InheritedAssignAddr"]
					}
					if oIpv6InheritedAssignAddr.Name != "" {
						nestedIpv6InheritedAssignAddr.Name = oIpv6InheritedAssignAddr.Name
					}
					if oIpv6InheritedAssignAddr.Type != nil {
						nestedIpv6InheritedAssignAddr.Type = &Ipv6InheritedAssignAddrTypeXml_11_0_2{}
						if _, ok := o.Misc["Ipv6InheritedAssignAddrType"]; ok {
							nestedIpv6InheritedAssignAddr.Type.Misc = o.Misc["Ipv6InheritedAssignAddrType"]
						}
						if oIpv6InheritedAssignAddr.Type.Gua != nil {
							nestedIpv6InheritedAssignAddr.Type.Gua = &Ipv6InheritedAssignAddrTypeGuaXml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGua"]; ok {
								nestedIpv6InheritedAssignAddr.Type.Gua.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGua"]
							}
							if oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.PrefixPool = oIpv6InheritedAssignAddr.Type.Gua.PrefixPool
							}
							if oIpv6InheritedAssignAddr.Type.Gua.PoolType != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.PoolType = &Ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"]
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2{}
									if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]; ok {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"]
									}
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2{}
									if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]; ok {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"]
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
									}
								}
							}
							if oIpv6InheritedAssignAddr.Type.Gua.Advertise != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua.Advertise = &Ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc = o.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"]
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
								}
							}
						}
						if oIpv6InheritedAssignAddr.Type.Ula != nil {
							nestedIpv6InheritedAssignAddr.Type.Ula = &Ipv6InheritedAssignAddrTypeUlaXml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeUla"]; ok {
								nestedIpv6InheritedAssignAddr.Type.Ula.Misc = o.Misc["Ipv6InheritedAssignAddrTypeUla"]
							}
							if oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Address != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Address = oIpv6InheritedAssignAddr.Type.Ula.Address
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Prefix != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Prefix = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Prefix, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Anycast != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Anycast = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Anycast, nil)
							}
							if oIpv6InheritedAssignAddr.Type.Ula.Advertise != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula.Advertise = &Ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2{}
								if _, ok := o.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"]; ok {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc = o.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"]
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.YesNo(oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
								}
							}
						}
					}
					nestedIpv6.Inherited.AssignAddr = append(nestedIpv6.Inherited.AssignAddr, nestedIpv6InheritedAssignAddr)
				}
			}
			if o.Ipv6.Inherited.Enable != nil {
				nestedIpv6.Inherited.Enable = util.YesNo(o.Ipv6.Inherited.Enable, nil)
			}
			if o.Ipv6.Inherited.NeighborDiscovery != nil {
				nestedIpv6.Inherited.NeighborDiscovery = &Ipv6InheritedNeighborDiscoveryXml_11_0_2{}
				if _, ok := o.Misc["Ipv6InheritedNeighborDiscovery"]; ok {
					nestedIpv6.Inherited.NeighborDiscovery.Misc = o.Misc["Ipv6InheritedNeighborDiscovery"]
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DadAttempts = o.Ipv6.Inherited.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DnsServer = &Ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source = &Ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"]; ok {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"]
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
							}
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2{}
								for _, oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
									nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2{}
									if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]; ok {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"]
									}
									if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
									}
									if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
									}
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
								}
							}
						}
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
					nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix = &Ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]; ok {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"]
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
							}
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2{}
							if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]; ok {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"]
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2{}
								for _, oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
									nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2{}
									if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]; ok {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"]
									}
									if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
									}
									if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
									}
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
								}
							}
						}
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.Inherited.NeighborDiscovery.EnableDad = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.Inherited.NeighborDiscovery.Neighbor = []Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2{}
					for _, oIpv6InheritedNeighborDiscoveryNeighbor := range o.Ipv6.Inherited.NeighborDiscovery.Neighbor {
						nestedIpv6InheritedNeighborDiscoveryNeighbor := Ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2{}
						if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"]; ok {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"]
						}
						if oIpv6InheritedNeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.Name = oIpv6InheritedNeighborDiscoveryNeighbor.Name
						}
						if oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6InheritedNeighborDiscoveryNeighbor.HwAddress = oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.Inherited.NeighborDiscovery.Neighbor = append(nestedIpv6.Inherited.NeighborDiscovery.Neighbor, nestedIpv6InheritedNeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.Inherited.NeighborDiscovery.NsInterval = o.Ipv6.Inherited.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.Inherited.NeighborDiscovery.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.ReachableTime
				}
				if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
					nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2{}
					if _, ok := o.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]; ok {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"]
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
				}
			}
		}
	}
	entry.Ipv6 = nestedIpv6

	entry.Mtu = o.Mtu
	var nestedNdpProxy *NdpProxyXml_11_0_2
	if o.NdpProxy != nil {
		nestedNdpProxy = &NdpProxyXml_11_0_2{}
		if _, ok := o.Misc["NdpProxy"]; ok {
			nestedNdpProxy.Misc = o.Misc["NdpProxy"]
		}
		if o.NdpProxy.Address != nil {
			nestedNdpProxy.Address = []NdpProxyAddressXml_11_0_2{}
			for _, oNdpProxyAddress := range o.NdpProxy.Address {
				nestedNdpProxyAddress := NdpProxyAddressXml_11_0_2{}
				if _, ok := o.Misc["NdpProxyAddress"]; ok {
					nestedNdpProxyAddress.Misc = o.Misc["NdpProxyAddress"]
				}
				if oNdpProxyAddress.Name != "" {
					nestedNdpProxyAddress.Name = oNdpProxyAddress.Name
				}
				if oNdpProxyAddress.Negate != nil {
					nestedNdpProxyAddress.Negate = util.YesNo(oNdpProxyAddress.Negate, nil)
				}
				nestedNdpProxy.Address = append(nestedNdpProxy.Address, nestedNdpProxyAddress)
			}
		}
		if o.NdpProxy.Enabled != nil {
			nestedNdpProxy.Enabled = util.YesNo(o.NdpProxy.Enabled, nil)
		}
	}
	entry.NdpProxy = nestedNdpProxy

	entry.NetflowProfile = o.NetflowProfile

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
		var nestedAdjustTcpMss *AdjustTcpMss
		if o.AdjustTcpMss != nil {
			nestedAdjustTcpMss = &AdjustTcpMss{}
			if o.AdjustTcpMss.Misc != nil {
				entry.Misc["AdjustTcpMss"] = o.AdjustTcpMss.Misc
			}
			if o.AdjustTcpMss.Enable != nil {
				nestedAdjustTcpMss.Enable = util.AsBool(o.AdjustTcpMss.Enable, nil)
			}
			if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
			}
			if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
			}
		}
		entry.AdjustTcpMss = nestedAdjustTcpMss

		var nestedArpCol []Arp
		if o.Arp != nil {
			nestedArpCol = []Arp{}
			for _, oArp := range o.Arp {
				nestedArp := Arp{}
				if oArp.Misc != nil {
					entry.Misc["Arp"] = oArp.Misc
				}
				if oArp.Name != "" {
					nestedArp.Name = oArp.Name
				}
				if oArp.HwAddress != nil {
					nestedArp.HwAddress = oArp.HwAddress
				}
				if oArp.Interface != nil {
					nestedArp.Interface = oArp.Interface
				}
				nestedArpCol = append(nestedArpCol, nestedArp)
			}
			entry.Arp = nestedArpCol
		}

		var nestedBonjour *Bonjour
		if o.Bonjour != nil {
			nestedBonjour = &Bonjour{}
			if o.Bonjour.Misc != nil {
				entry.Misc["Bonjour"] = o.Bonjour.Misc
			}
			if o.Bonjour.Enable != nil {
				nestedBonjour.Enable = util.AsBool(o.Bonjour.Enable, nil)
			}
			if o.Bonjour.GroupId != nil {
				nestedBonjour.GroupId = o.Bonjour.GroupId
			}
			if o.Bonjour.TtlCheck != nil {
				nestedBonjour.TtlCheck = util.AsBool(o.Bonjour.TtlCheck, nil)
			}
		}
		entry.Bonjour = nestedBonjour

		entry.Comment = o.Comment
		var nestedDdnsConfig *DdnsConfig
		if o.DdnsConfig != nil {
			nestedDdnsConfig = &DdnsConfig{}
			if o.DdnsConfig.Misc != nil {
				entry.Misc["DdnsConfig"] = o.DdnsConfig.Misc
			}
			if o.DdnsConfig.DdnsCertProfile != nil {
				nestedDdnsConfig.DdnsCertProfile = o.DdnsConfig.DdnsCertProfile
			}
			if o.DdnsConfig.DdnsEnabled != nil {
				nestedDdnsConfig.DdnsEnabled = util.AsBool(o.DdnsConfig.DdnsEnabled, nil)
			}
			if o.DdnsConfig.DdnsHostname != nil {
				nestedDdnsConfig.DdnsHostname = o.DdnsConfig.DdnsHostname
			}
			if o.DdnsConfig.DdnsIp != nil {
				nestedDdnsConfig.DdnsIp = util.MemToStr(o.DdnsConfig.DdnsIp)
			}
			if o.DdnsConfig.DdnsIpv6 != nil {
				nestedDdnsConfig.DdnsIpv6 = util.MemToStr(o.DdnsConfig.DdnsIpv6)
			}
			if o.DdnsConfig.DdnsUpdateInterval != nil {
				nestedDdnsConfig.DdnsUpdateInterval = o.DdnsConfig.DdnsUpdateInterval
			}
			if o.DdnsConfig.DdnsVendor != nil {
				nestedDdnsConfig.DdnsVendor = o.DdnsConfig.DdnsVendor
			}
			if o.DdnsConfig.DdnsVendorConfig != nil {
				nestedDdnsConfig.DdnsVendorConfig = []DdnsConfigDdnsVendorConfig{}
				for _, oDdnsConfigDdnsVendorConfig := range o.DdnsConfig.DdnsVendorConfig {
					nestedDdnsConfigDdnsVendorConfig := DdnsConfigDdnsVendorConfig{}
					if oDdnsConfigDdnsVendorConfig.Misc != nil {
						entry.Misc["DdnsConfigDdnsVendorConfig"] = oDdnsConfigDdnsVendorConfig.Misc
					}
					if oDdnsConfigDdnsVendorConfig.Name != "" {
						nestedDdnsConfigDdnsVendorConfig.Name = oDdnsConfigDdnsVendorConfig.Name
					}
					if oDdnsConfigDdnsVendorConfig.Value != nil {
						nestedDdnsConfigDdnsVendorConfig.Value = oDdnsConfigDdnsVendorConfig.Value
					}
					nestedDdnsConfig.DdnsVendorConfig = append(nestedDdnsConfig.DdnsVendorConfig, nestedDdnsConfigDdnsVendorConfig)
				}
			}
		}
		entry.DdnsConfig = nestedDdnsConfig

		entry.DfIgnore = util.AsBool(o.DfIgnore, nil)
		var nestedDhcpClient *DhcpClient
		if o.DhcpClient != nil {
			nestedDhcpClient = &DhcpClient{}
			if o.DhcpClient.Misc != nil {
				entry.Misc["DhcpClient"] = o.DhcpClient.Misc
			}
			if o.DhcpClient.CreateDefaultRoute != nil {
				nestedDhcpClient.CreateDefaultRoute = util.AsBool(o.DhcpClient.CreateDefaultRoute, nil)
			}
			if o.DhcpClient.DefaultRouteMetric != nil {
				nestedDhcpClient.DefaultRouteMetric = o.DhcpClient.DefaultRouteMetric
			}
			if o.DhcpClient.Enable != nil {
				nestedDhcpClient.Enable = util.AsBool(o.DhcpClient.Enable, nil)
			}
			if o.DhcpClient.SendHostname != nil {
				nestedDhcpClient.SendHostname = &DhcpClientSendHostname{}
				if o.DhcpClient.SendHostname.Misc != nil {
					entry.Misc["DhcpClientSendHostname"] = o.DhcpClient.SendHostname.Misc
				}
				if o.DhcpClient.SendHostname.Enable != nil {
					nestedDhcpClient.SendHostname.Enable = util.AsBool(o.DhcpClient.SendHostname.Enable, nil)
				}
				if o.DhcpClient.SendHostname.Hostname != nil {
					nestedDhcpClient.SendHostname.Hostname = o.DhcpClient.SendHostname.Hostname
				}
			}
		}
		entry.DhcpClient = nestedDhcpClient

		entry.InterfaceManagementProfile = o.InterfaceManagementProfile
		var nestedIpCol []Ip
		if o.Ip != nil {
			nestedIpCol = []Ip{}
			for _, oIp := range o.Ip {
				nestedIp := Ip{}
				if oIp.Misc != nil {
					entry.Misc["Ip"] = oIp.Misc
				}
				if oIp.Name != "" {
					nestedIp.Name = oIp.Name
				}
				nestedIpCol = append(nestedIpCol, nestedIp)
			}
			entry.Ip = nestedIpCol
		}

		var nestedIpv6 *Ipv6
		if o.Ipv6 != nil {
			nestedIpv6 = &Ipv6{}
			if o.Ipv6.Misc != nil {
				entry.Misc["Ipv6"] = o.Ipv6.Misc
			}
			if o.Ipv6.Address != nil {
				nestedIpv6.Address = []Ipv6Address{}
				for _, oIpv6Address := range o.Ipv6.Address {
					nestedIpv6Address := Ipv6Address{}
					if oIpv6Address.Misc != nil {
						entry.Misc["Ipv6Address"] = oIpv6Address.Misc
					}
					if oIpv6Address.Name != "" {
						nestedIpv6Address.Name = oIpv6Address.Name
					}
					if oIpv6Address.EnableOnInterface != nil {
						nestedIpv6Address.EnableOnInterface = util.AsBool(oIpv6Address.EnableOnInterface, nil)
					}
					if oIpv6Address.Prefix != nil {
						nestedIpv6Address.Prefix = &Ipv6AddressPrefix{}
						if oIpv6Address.Prefix.Misc != nil {
							entry.Misc["Ipv6AddressPrefix"] = oIpv6Address.Prefix.Misc
						}
					}
					if oIpv6Address.Anycast != nil {
						nestedIpv6Address.Anycast = &Ipv6AddressAnycast{}
						if oIpv6Address.Anycast.Misc != nil {
							entry.Misc["Ipv6AddressAnycast"] = oIpv6Address.Anycast.Misc
						}
					}
					if oIpv6Address.Advertise != nil {
						nestedIpv6Address.Advertise = &Ipv6AddressAdvertise{}
						if oIpv6Address.Advertise.Misc != nil {
							entry.Misc["Ipv6AddressAdvertise"] = oIpv6Address.Advertise.Misc
						}
						if oIpv6Address.Advertise.Enable != nil {
							nestedIpv6Address.Advertise.Enable = util.AsBool(oIpv6Address.Advertise.Enable, nil)
						}
						if oIpv6Address.Advertise.ValidLifetime != nil {
							nestedIpv6Address.Advertise.ValidLifetime = oIpv6Address.Advertise.ValidLifetime
						}
						if oIpv6Address.Advertise.PreferredLifetime != nil {
							nestedIpv6Address.Advertise.PreferredLifetime = oIpv6Address.Advertise.PreferredLifetime
						}
						if oIpv6Address.Advertise.OnlinkFlag != nil {
							nestedIpv6Address.Advertise.OnlinkFlag = util.AsBool(oIpv6Address.Advertise.OnlinkFlag, nil)
						}
						if oIpv6Address.Advertise.AutoConfigFlag != nil {
							nestedIpv6Address.Advertise.AutoConfigFlag = util.AsBool(oIpv6Address.Advertise.AutoConfigFlag, nil)
						}
					}
					nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
				}
			}
			if o.Ipv6.Enabled != nil {
				nestedIpv6.Enabled = util.AsBool(o.Ipv6.Enabled, nil)
			}
			if o.Ipv6.InterfaceId != nil {
				nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
			}
			if o.Ipv6.NeighborDiscovery != nil {
				nestedIpv6.NeighborDiscovery = &Ipv6NeighborDiscovery{}
				if o.Ipv6.NeighborDiscovery.Misc != nil {
					entry.Misc["Ipv6NeighborDiscovery"] = o.Ipv6.NeighborDiscovery.Misc
				}
				if o.Ipv6.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.NeighborDiscovery.DadAttempts = o.Ipv6.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.NeighborDiscovery.Neighbor = []Ipv6NeighborDiscoveryNeighbor{}
					for _, oIpv6NeighborDiscoveryNeighbor := range o.Ipv6.NeighborDiscovery.Neighbor {
						nestedIpv6NeighborDiscoveryNeighbor := Ipv6NeighborDiscoveryNeighbor{}
						if oIpv6NeighborDiscoveryNeighbor.Misc != nil {
							entry.Misc["Ipv6NeighborDiscoveryNeighbor"] = oIpv6NeighborDiscoveryNeighbor.Misc
						}
						if oIpv6NeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6NeighborDiscoveryNeighbor.Name = oIpv6NeighborDiscoveryNeighbor.Name
						}
						if oIpv6NeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6NeighborDiscoveryNeighbor.HwAddress = oIpv6NeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.NeighborDiscovery.Neighbor = append(nestedIpv6.NeighborDiscovery.Neighbor, nestedIpv6NeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.NeighborDiscovery.NsInterval = o.Ipv6.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.NeighborDiscovery.ReachableTime = o.Ipv6.NeighborDiscovery.ReachableTime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement = &Ipv6NeighborDiscoveryRouterAdvertisement{}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc != nil {
						entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"] = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc != nil {
							entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"] = o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
							for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc != nil {
									entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"] = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
								}
								nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
							}
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
							for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc != nil {
									entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"] = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
								}
								nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
							}
						}
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
				}
			}
			if o.Ipv6.DhcpClient != nil {
				nestedIpv6.DhcpClient = &Ipv6DhcpClient{}
				if o.Ipv6.DhcpClient.Misc != nil {
					entry.Misc["Ipv6DhcpClient"] = o.Ipv6.DhcpClient.Misc
				}
				if o.Ipv6.DhcpClient.AcceptRaRoute != nil {
					nestedIpv6.DhcpClient.AcceptRaRoute = util.AsBool(o.Ipv6.DhcpClient.AcceptRaRoute, nil)
				}
				if o.Ipv6.DhcpClient.DefaultRouteMetric != nil {
					nestedIpv6.DhcpClient.DefaultRouteMetric = o.Ipv6.DhcpClient.DefaultRouteMetric
				}
				if o.Ipv6.DhcpClient.Enable != nil {
					nestedIpv6.DhcpClient.Enable = util.AsBool(o.Ipv6.DhcpClient.Enable, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery = &Ipv6DhcpClientNeighborDiscovery{}
					if o.Ipv6.DhcpClient.NeighborDiscovery.Misc != nil {
						entry.Misc["Ipv6DhcpClientNeighborDiscovery"] = o.Ipv6.DhcpClient.NeighborDiscovery.Misc
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer = &Ipv6DhcpClientNeighborDiscoveryDnsServer{}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc != nil {
							entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Ipv6DhcpClientNeighborDiscoveryDnsServerSource{}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
								}
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc
								}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
									for _, oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
											entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"] = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Ipv6DhcpClientNeighborDiscoveryDnsSuffix{}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc != nil {
							entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
								}
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
								}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
									for _, oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
											entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = []Ipv6DhcpClientNeighborDiscoveryNeighbor{}
						for _, oIpv6DhcpClientNeighborDiscoveryNeighbor := range o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor := Ipv6DhcpClientNeighborDiscoveryNeighbor{}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"] = oIpv6DhcpClientNeighborDiscoveryNeighbor.Misc
							}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
								nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Name = oIpv6DhcpClientNeighborDiscoveryNeighbor.Name
							}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
							}
							nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor, nestedIpv6DhcpClientNeighborDiscoveryNeighbor)
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.NsInterval = o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
					}
				}
				if o.Ipv6.DhcpClient.Preference != nil {
					nestedIpv6.DhcpClient.Preference = o.Ipv6.DhcpClient.Preference
				}
				if o.Ipv6.DhcpClient.PrefixDelegation != nil {
					nestedIpv6.DhcpClient.PrefixDelegation = &Ipv6DhcpClientPrefixDelegation{}
					if o.Ipv6.DhcpClient.PrefixDelegation.Misc != nil {
						entry.Misc["Ipv6DhcpClientPrefixDelegation"] = o.Ipv6.DhcpClient.PrefixDelegation.Misc
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable = &Ipv6DhcpClientPrefixDelegationEnable{}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc != nil {
							entry.Misc["Ipv6DhcpClientPrefixDelegationEnable"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.No = &Ipv6DhcpClientPrefixDelegationEnableNo{}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc != nil {
								entry.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc
							}
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes = &Ipv6DhcpClientPrefixDelegationEnableYes{}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc != nil {
								entry.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.AsBool(o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.V6Options != nil {
					nestedIpv6.DhcpClient.V6Options = &Ipv6DhcpClientV6Options{}
					if o.Ipv6.DhcpClient.V6Options.Misc != nil {
						entry.Misc["Ipv6DhcpClientV6Options"] = o.Ipv6.DhcpClient.V6Options.Misc
					}
					if o.Ipv6.DhcpClient.V6Options.DuidType != nil {
						nestedIpv6.DhcpClient.V6Options.DuidType = o.Ipv6.DhcpClient.V6Options.DuidType
					}
					if o.Ipv6.DhcpClient.V6Options.Enable != nil {
						nestedIpv6.DhcpClient.V6Options.Enable = &Ipv6DhcpClientV6OptionsEnable{}
						if o.Ipv6.DhcpClient.V6Options.Enable.Misc != nil {
							entry.Misc["Ipv6DhcpClientV6OptionsEnable"] = o.Ipv6.DhcpClient.V6Options.Enable.Misc
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.No != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.No = &Ipv6DhcpClientV6OptionsEnableNo{}
							if o.Ipv6.DhcpClient.V6Options.Enable.No.Misc != nil {
								entry.Misc["Ipv6DhcpClientV6OptionsEnableNo"] = o.Ipv6.DhcpClient.V6Options.Enable.No.Misc
							}
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes = &Ipv6DhcpClientV6OptionsEnableYes{}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc != nil {
								entry.Misc["Ipv6DhcpClientV6OptionsEnableYes"] = o.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc
							}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
								nestedIpv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.AsBool(o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
							}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
								nestedIpv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.AsBool(o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
							}
						}
					}
					if o.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
						nestedIpv6.DhcpClient.V6Options.RapidCommit = util.AsBool(o.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
					}
					if o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
						nestedIpv6.DhcpClient.V6Options.SupportSrvrReconfig = util.AsBool(o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
					}
				}
			}
			if o.Ipv6.Inherited != nil {
				nestedIpv6.Inherited = &Ipv6Inherited{}
				if o.Ipv6.Inherited.Misc != nil {
					entry.Misc["Ipv6Inherited"] = o.Ipv6.Inherited.Misc
				}
				if o.Ipv6.Inherited.AssignAddr != nil {
					nestedIpv6.Inherited.AssignAddr = []Ipv6InheritedAssignAddr{}
					for _, oIpv6InheritedAssignAddr := range o.Ipv6.Inherited.AssignAddr {
						nestedIpv6InheritedAssignAddr := Ipv6InheritedAssignAddr{}
						if oIpv6InheritedAssignAddr.Misc != nil {
							entry.Misc["Ipv6InheritedAssignAddr"] = oIpv6InheritedAssignAddr.Misc
						}
						if oIpv6InheritedAssignAddr.Name != "" {
							nestedIpv6InheritedAssignAddr.Name = oIpv6InheritedAssignAddr.Name
						}
						if oIpv6InheritedAssignAddr.Type != nil {
							nestedIpv6InheritedAssignAddr.Type = &Ipv6InheritedAssignAddrType{}
							if oIpv6InheritedAssignAddr.Type.Misc != nil {
								entry.Misc["Ipv6InheritedAssignAddrType"] = oIpv6InheritedAssignAddr.Type.Misc
							}
							if oIpv6InheritedAssignAddr.Type.Gua != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua = &Ipv6InheritedAssignAddrTypeGua{}
								if oIpv6InheritedAssignAddr.Type.Gua.Misc != nil {
									entry.Misc["Ipv6InheritedAssignAddrTypeGua"] = oIpv6InheritedAssignAddr.Type.Gua.Misc
								}
								if oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PrefixPool = oIpv6InheritedAssignAddr.Type.Gua.PrefixPool
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType = &Ipv6InheritedAssignAddrTypeGuaPoolType{}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc != nil {
											entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc
										}
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc != nil {
											entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc
										}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
											nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
										}
									}
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise = &Ipv6InheritedAssignAddrTypeGuaAdvertise{}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"] = oIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
									}
								}
							}
							if oIpv6InheritedAssignAddr.Type.Ula != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula = &Ipv6InheritedAssignAddrTypeUla{}
								if oIpv6InheritedAssignAddr.Type.Ula.Misc != nil {
									entry.Misc["Ipv6InheritedAssignAddrTypeUla"] = oIpv6InheritedAssignAddr.Type.Ula.Misc
								}
								if oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Address != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Address = oIpv6InheritedAssignAddr.Type.Ula.Address
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Prefix != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Prefix = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Prefix, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Anycast != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Anycast = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Anycast, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise = &Ipv6InheritedAssignAddrTypeUlaAdvertise{}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"] = oIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
									}
								}
							}
						}
						nestedIpv6.Inherited.AssignAddr = append(nestedIpv6.Inherited.AssignAddr, nestedIpv6InheritedAssignAddr)
					}
				}
				if o.Ipv6.Inherited.Enable != nil {
					nestedIpv6.Inherited.Enable = util.AsBool(o.Ipv6.Inherited.Enable, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery != nil {
					nestedIpv6.Inherited.NeighborDiscovery = &Ipv6InheritedNeighborDiscovery{}
					if o.Ipv6.Inherited.NeighborDiscovery.Misc != nil {
						entry.Misc["Ipv6InheritedNeighborDiscovery"] = o.Ipv6.Inherited.NeighborDiscovery.Misc
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DadAttempts = o.Ipv6.Inherited.NeighborDiscovery.DadAttempts
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer = &Ipv6InheritedNeighborDiscoveryDnsServer{}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source = &Ipv6InheritedNeighborDiscoveryDnsServerSource{}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
									for _, oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
											entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"] = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc
										}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix = &Ipv6InheritedNeighborDiscoveryDnsSuffix{}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Ipv6InheritedNeighborDiscoveryDnsSuffixSource{}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
									for _, oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
											entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
										}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
						nestedIpv6.Inherited.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedIpv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
						nestedIpv6.Inherited.NeighborDiscovery.Neighbor = []Ipv6InheritedNeighborDiscoveryNeighbor{}
						for _, oIpv6InheritedNeighborDiscoveryNeighbor := range o.Ipv6.Inherited.NeighborDiscovery.Neighbor {
							nestedIpv6InheritedNeighborDiscoveryNeighbor := Ipv6InheritedNeighborDiscoveryNeighbor{}
							if oIpv6InheritedNeighborDiscoveryNeighbor.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"] = oIpv6InheritedNeighborDiscoveryNeighbor.Misc
							}
							if oIpv6InheritedNeighborDiscoveryNeighbor.Name != "" {
								nestedIpv6InheritedNeighborDiscoveryNeighbor.Name = oIpv6InheritedNeighborDiscoveryNeighbor.Name
							}
							if oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedIpv6InheritedNeighborDiscoveryNeighbor.HwAddress = oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress
							}
							nestedIpv6.Inherited.NeighborDiscovery.Neighbor = append(nestedIpv6.Inherited.NeighborDiscovery.Neighbor, nestedIpv6InheritedNeighborDiscoveryNeighbor)
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.NsInterval = o.Ipv6.Inherited.NeighborDiscovery.NsInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.ReachableTime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Ipv6InheritedNeighborDiscoveryRouterAdvertisement{}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"] = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
						}
					}
				}
			}
		}
		entry.Ipv6 = nestedIpv6

		entry.Mtu = o.Mtu
		var nestedNdpProxy *NdpProxy
		if o.NdpProxy != nil {
			nestedNdpProxy = &NdpProxy{}
			if o.NdpProxy.Misc != nil {
				entry.Misc["NdpProxy"] = o.NdpProxy.Misc
			}
			if o.NdpProxy.Address != nil {
				nestedNdpProxy.Address = []NdpProxyAddress{}
				for _, oNdpProxyAddress := range o.NdpProxy.Address {
					nestedNdpProxyAddress := NdpProxyAddress{}
					if oNdpProxyAddress.Misc != nil {
						entry.Misc["NdpProxyAddress"] = oNdpProxyAddress.Misc
					}
					if oNdpProxyAddress.Name != "" {
						nestedNdpProxyAddress.Name = oNdpProxyAddress.Name
					}
					if oNdpProxyAddress.Negate != nil {
						nestedNdpProxyAddress.Negate = util.AsBool(oNdpProxyAddress.Negate, nil)
					}
					nestedNdpProxy.Address = append(nestedNdpProxy.Address, nestedNdpProxyAddress)
				}
			}
			if o.NdpProxy.Enabled != nil {
				nestedNdpProxy.Enabled = util.AsBool(o.NdpProxy.Enabled, nil)
			}
		}
		entry.NdpProxy = nestedNdpProxy

		entry.NetflowProfile = o.NetflowProfile

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
		var nestedAdjustTcpMss *AdjustTcpMss
		if o.AdjustTcpMss != nil {
			nestedAdjustTcpMss = &AdjustTcpMss{}
			if o.AdjustTcpMss.Misc != nil {
				entry.Misc["AdjustTcpMss"] = o.AdjustTcpMss.Misc
			}
			if o.AdjustTcpMss.Enable != nil {
				nestedAdjustTcpMss.Enable = util.AsBool(o.AdjustTcpMss.Enable, nil)
			}
			if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
			}
			if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
			}
		}
		entry.AdjustTcpMss = nestedAdjustTcpMss

		var nestedArpCol []Arp
		if o.Arp != nil {
			nestedArpCol = []Arp{}
			for _, oArp := range o.Arp {
				nestedArp := Arp{}
				if oArp.Misc != nil {
					entry.Misc["Arp"] = oArp.Misc
				}
				if oArp.Name != "" {
					nestedArp.Name = oArp.Name
				}
				if oArp.HwAddress != nil {
					nestedArp.HwAddress = oArp.HwAddress
				}
				if oArp.Interface != nil {
					nestedArp.Interface = oArp.Interface
				}
				nestedArpCol = append(nestedArpCol, nestedArp)
			}
			entry.Arp = nestedArpCol
		}

		var nestedBonjour *Bonjour
		if o.Bonjour != nil {
			nestedBonjour = &Bonjour{}
			if o.Bonjour.Misc != nil {
				entry.Misc["Bonjour"] = o.Bonjour.Misc
			}
			if o.Bonjour.Enable != nil {
				nestedBonjour.Enable = util.AsBool(o.Bonjour.Enable, nil)
			}
			if o.Bonjour.GroupId != nil {
				nestedBonjour.GroupId = o.Bonjour.GroupId
			}
			if o.Bonjour.TtlCheck != nil {
				nestedBonjour.TtlCheck = util.AsBool(o.Bonjour.TtlCheck, nil)
			}
		}
		entry.Bonjour = nestedBonjour

		entry.Comment = o.Comment
		var nestedDdnsConfig *DdnsConfig
		if o.DdnsConfig != nil {
			nestedDdnsConfig = &DdnsConfig{}
			if o.DdnsConfig.Misc != nil {
				entry.Misc["DdnsConfig"] = o.DdnsConfig.Misc
			}
			if o.DdnsConfig.DdnsCertProfile != nil {
				nestedDdnsConfig.DdnsCertProfile = o.DdnsConfig.DdnsCertProfile
			}
			if o.DdnsConfig.DdnsEnabled != nil {
				nestedDdnsConfig.DdnsEnabled = util.AsBool(o.DdnsConfig.DdnsEnabled, nil)
			}
			if o.DdnsConfig.DdnsHostname != nil {
				nestedDdnsConfig.DdnsHostname = o.DdnsConfig.DdnsHostname
			}
			if o.DdnsConfig.DdnsIp != nil {
				nestedDdnsConfig.DdnsIp = util.MemToStr(o.DdnsConfig.DdnsIp)
			}
			if o.DdnsConfig.DdnsIpv6 != nil {
				nestedDdnsConfig.DdnsIpv6 = util.MemToStr(o.DdnsConfig.DdnsIpv6)
			}
			if o.DdnsConfig.DdnsUpdateInterval != nil {
				nestedDdnsConfig.DdnsUpdateInterval = o.DdnsConfig.DdnsUpdateInterval
			}
			if o.DdnsConfig.DdnsVendor != nil {
				nestedDdnsConfig.DdnsVendor = o.DdnsConfig.DdnsVendor
			}
			if o.DdnsConfig.DdnsVendorConfig != nil {
				nestedDdnsConfig.DdnsVendorConfig = []DdnsConfigDdnsVendorConfig{}
				for _, oDdnsConfigDdnsVendorConfig := range o.DdnsConfig.DdnsVendorConfig {
					nestedDdnsConfigDdnsVendorConfig := DdnsConfigDdnsVendorConfig{}
					if oDdnsConfigDdnsVendorConfig.Misc != nil {
						entry.Misc["DdnsConfigDdnsVendorConfig"] = oDdnsConfigDdnsVendorConfig.Misc
					}
					if oDdnsConfigDdnsVendorConfig.Name != "" {
						nestedDdnsConfigDdnsVendorConfig.Name = oDdnsConfigDdnsVendorConfig.Name
					}
					if oDdnsConfigDdnsVendorConfig.Value != nil {
						nestedDdnsConfigDdnsVendorConfig.Value = oDdnsConfigDdnsVendorConfig.Value
					}
					nestedDdnsConfig.DdnsVendorConfig = append(nestedDdnsConfig.DdnsVendorConfig, nestedDdnsConfigDdnsVendorConfig)
				}
			}
		}
		entry.DdnsConfig = nestedDdnsConfig

		entry.DfIgnore = util.AsBool(o.DfIgnore, nil)
		var nestedDhcpClient *DhcpClient
		if o.DhcpClient != nil {
			nestedDhcpClient = &DhcpClient{}
			if o.DhcpClient.Misc != nil {
				entry.Misc["DhcpClient"] = o.DhcpClient.Misc
			}
			if o.DhcpClient.CreateDefaultRoute != nil {
				nestedDhcpClient.CreateDefaultRoute = util.AsBool(o.DhcpClient.CreateDefaultRoute, nil)
			}
			if o.DhcpClient.DefaultRouteMetric != nil {
				nestedDhcpClient.DefaultRouteMetric = o.DhcpClient.DefaultRouteMetric
			}
			if o.DhcpClient.Enable != nil {
				nestedDhcpClient.Enable = util.AsBool(o.DhcpClient.Enable, nil)
			}
			if o.DhcpClient.SendHostname != nil {
				nestedDhcpClient.SendHostname = &DhcpClientSendHostname{}
				if o.DhcpClient.SendHostname.Misc != nil {
					entry.Misc["DhcpClientSendHostname"] = o.DhcpClient.SendHostname.Misc
				}
				if o.DhcpClient.SendHostname.Enable != nil {
					nestedDhcpClient.SendHostname.Enable = util.AsBool(o.DhcpClient.SendHostname.Enable, nil)
				}
				if o.DhcpClient.SendHostname.Hostname != nil {
					nestedDhcpClient.SendHostname.Hostname = o.DhcpClient.SendHostname.Hostname
				}
			}
		}
		entry.DhcpClient = nestedDhcpClient

		entry.InterfaceManagementProfile = o.InterfaceManagementProfile
		var nestedIpCol []Ip
		if o.Ip != nil {
			nestedIpCol = []Ip{}
			for _, oIp := range o.Ip {
				nestedIp := Ip{}
				if oIp.Misc != nil {
					entry.Misc["Ip"] = oIp.Misc
				}
				if oIp.Name != "" {
					nestedIp.Name = oIp.Name
				}
				nestedIpCol = append(nestedIpCol, nestedIp)
			}
			entry.Ip = nestedIpCol
		}

		var nestedIpv6 *Ipv6
		if o.Ipv6 != nil {
			nestedIpv6 = &Ipv6{}
			if o.Ipv6.Misc != nil {
				entry.Misc["Ipv6"] = o.Ipv6.Misc
			}
			if o.Ipv6.Address != nil {
				nestedIpv6.Address = []Ipv6Address{}
				for _, oIpv6Address := range o.Ipv6.Address {
					nestedIpv6Address := Ipv6Address{}
					if oIpv6Address.Misc != nil {
						entry.Misc["Ipv6Address"] = oIpv6Address.Misc
					}
					if oIpv6Address.Name != "" {
						nestedIpv6Address.Name = oIpv6Address.Name
					}
					if oIpv6Address.EnableOnInterface != nil {
						nestedIpv6Address.EnableOnInterface = util.AsBool(oIpv6Address.EnableOnInterface, nil)
					}
					if oIpv6Address.Prefix != nil {
						nestedIpv6Address.Prefix = &Ipv6AddressPrefix{}
						if oIpv6Address.Prefix.Misc != nil {
							entry.Misc["Ipv6AddressPrefix"] = oIpv6Address.Prefix.Misc
						}
					}
					if oIpv6Address.Anycast != nil {
						nestedIpv6Address.Anycast = &Ipv6AddressAnycast{}
						if oIpv6Address.Anycast.Misc != nil {
							entry.Misc["Ipv6AddressAnycast"] = oIpv6Address.Anycast.Misc
						}
					}
					if oIpv6Address.Advertise != nil {
						nestedIpv6Address.Advertise = &Ipv6AddressAdvertise{}
						if oIpv6Address.Advertise.Misc != nil {
							entry.Misc["Ipv6AddressAdvertise"] = oIpv6Address.Advertise.Misc
						}
						if oIpv6Address.Advertise.Enable != nil {
							nestedIpv6Address.Advertise.Enable = util.AsBool(oIpv6Address.Advertise.Enable, nil)
						}
						if oIpv6Address.Advertise.ValidLifetime != nil {
							nestedIpv6Address.Advertise.ValidLifetime = oIpv6Address.Advertise.ValidLifetime
						}
						if oIpv6Address.Advertise.PreferredLifetime != nil {
							nestedIpv6Address.Advertise.PreferredLifetime = oIpv6Address.Advertise.PreferredLifetime
						}
						if oIpv6Address.Advertise.OnlinkFlag != nil {
							nestedIpv6Address.Advertise.OnlinkFlag = util.AsBool(oIpv6Address.Advertise.OnlinkFlag, nil)
						}
						if oIpv6Address.Advertise.AutoConfigFlag != nil {
							nestedIpv6Address.Advertise.AutoConfigFlag = util.AsBool(oIpv6Address.Advertise.AutoConfigFlag, nil)
						}
					}
					nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
				}
			}
			if o.Ipv6.Enabled != nil {
				nestedIpv6.Enabled = util.AsBool(o.Ipv6.Enabled, nil)
			}
			if o.Ipv6.InterfaceId != nil {
				nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
			}
			if o.Ipv6.NeighborDiscovery != nil {
				nestedIpv6.NeighborDiscovery = &Ipv6NeighborDiscovery{}
				if o.Ipv6.NeighborDiscovery.Misc != nil {
					entry.Misc["Ipv6NeighborDiscovery"] = o.Ipv6.NeighborDiscovery.Misc
				}
				if o.Ipv6.NeighborDiscovery.DadAttempts != nil {
					nestedIpv6.NeighborDiscovery.DadAttempts = o.Ipv6.NeighborDiscovery.DadAttempts
				}
				if o.Ipv6.NeighborDiscovery.EnableDad != nil {
					nestedIpv6.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.NeighborDiscovery.EnableDad, nil)
				}
				if o.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedIpv6.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Ipv6.NeighborDiscovery.Neighbor != nil {
					nestedIpv6.NeighborDiscovery.Neighbor = []Ipv6NeighborDiscoveryNeighbor{}
					for _, oIpv6NeighborDiscoveryNeighbor := range o.Ipv6.NeighborDiscovery.Neighbor {
						nestedIpv6NeighborDiscoveryNeighbor := Ipv6NeighborDiscoveryNeighbor{}
						if oIpv6NeighborDiscoveryNeighbor.Misc != nil {
							entry.Misc["Ipv6NeighborDiscoveryNeighbor"] = oIpv6NeighborDiscoveryNeighbor.Misc
						}
						if oIpv6NeighborDiscoveryNeighbor.Name != "" {
							nestedIpv6NeighborDiscoveryNeighbor.Name = oIpv6NeighborDiscoveryNeighbor.Name
						}
						if oIpv6NeighborDiscoveryNeighbor.HwAddress != nil {
							nestedIpv6NeighborDiscoveryNeighbor.HwAddress = oIpv6NeighborDiscoveryNeighbor.HwAddress
						}
						nestedIpv6.NeighborDiscovery.Neighbor = append(nestedIpv6.NeighborDiscovery.Neighbor, nestedIpv6NeighborDiscoveryNeighbor)
					}
				}
				if o.Ipv6.NeighborDiscovery.NsInterval != nil {
					nestedIpv6.NeighborDiscovery.NsInterval = o.Ipv6.NeighborDiscovery.NsInterval
				}
				if o.Ipv6.NeighborDiscovery.ReachableTime != nil {
					nestedIpv6.NeighborDiscovery.ReachableTime = o.Ipv6.NeighborDiscovery.ReachableTime
				}
				if o.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
					nestedIpv6.NeighborDiscovery.RouterAdvertisement = &Ipv6NeighborDiscoveryRouterAdvertisement{}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc != nil {
						entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisement"] = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport = &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc != nil {
							entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport"] = o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Misc
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Enable, nil)
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
							for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc != nil {
									entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer"] = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Misc
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name != "" {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Name
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime != nil {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer.Lifetime
								}
								nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Server, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer)
							}
						}
						if o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix != nil {
							nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
							for _, oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := range o.Ipv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix {
								nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix := Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc != nil {
									entry.Misc["Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix"] = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Misc
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name != "" {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Name
								}
								if oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime != nil {
									nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime = oIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix.Lifetime
								}
								nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix = append(nestedIpv6.NeighborDiscovery.RouterAdvertisement.DnsSupport.Suffix, nestedIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix)
							}
						}
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.NeighborDiscovery.RouterAdvertisement.MinInterval
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.NeighborDiscovery.RouterAdvertisement.ReachableTime
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
					}
					if o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedIpv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
				}
			}
			if o.Ipv6.DhcpClient != nil {
				nestedIpv6.DhcpClient = &Ipv6DhcpClient{}
				if o.Ipv6.DhcpClient.Misc != nil {
					entry.Misc["Ipv6DhcpClient"] = o.Ipv6.DhcpClient.Misc
				}
				if o.Ipv6.DhcpClient.AcceptRaRoute != nil {
					nestedIpv6.DhcpClient.AcceptRaRoute = util.AsBool(o.Ipv6.DhcpClient.AcceptRaRoute, nil)
				}
				if o.Ipv6.DhcpClient.DefaultRouteMetric != nil {
					nestedIpv6.DhcpClient.DefaultRouteMetric = o.Ipv6.DhcpClient.DefaultRouteMetric
				}
				if o.Ipv6.DhcpClient.Enable != nil {
					nestedIpv6.DhcpClient.Enable = util.AsBool(o.Ipv6.DhcpClient.Enable, nil)
				}
				if o.Ipv6.DhcpClient.NeighborDiscovery != nil {
					nestedIpv6.DhcpClient.NeighborDiscovery = &Ipv6DhcpClientNeighborDiscovery{}
					if o.Ipv6.DhcpClient.NeighborDiscovery.Misc != nil {
						entry.Misc["Ipv6DhcpClientNeighborDiscovery"] = o.Ipv6.DhcpClient.NeighborDiscovery.Misc
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DadAttempts = o.Ipv6.DhcpClient.NeighborDiscovery.DadAttempts
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer = &Ipv6DhcpClientNeighborDiscoveryDnsServer{}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc != nil {
							entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServer"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Misc
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source = &Ipv6DhcpClientNeighborDiscoveryDnsServerSource{}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSource"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Misc
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
								}
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Misc
								}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
									for _, oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer := Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
											entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer"] = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Misc
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix = &Ipv6DhcpClientNeighborDiscoveryDnsSuffix{}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc != nil {
							entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffix"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Misc
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
						if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Misc
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
								}
							}
							if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
									entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual"] = o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
								}
								if o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
									for _, oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
											entry.Misc["Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										if oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.DhcpClient.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.EnableDad, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.DhcpClient.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = []Ipv6DhcpClientNeighborDiscoveryNeighbor{}
						for _, oIpv6DhcpClientNeighborDiscoveryNeighbor := range o.Ipv6.DhcpClient.NeighborDiscovery.Neighbor {
							nestedIpv6DhcpClientNeighborDiscoveryNeighbor := Ipv6DhcpClientNeighborDiscoveryNeighbor{}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.Misc != nil {
								entry.Misc["Ipv6DhcpClientNeighborDiscoveryNeighbor"] = oIpv6DhcpClientNeighborDiscoveryNeighbor.Misc
							}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.Name != "" {
								nestedIpv6DhcpClientNeighborDiscoveryNeighbor.Name = oIpv6DhcpClientNeighborDiscoveryNeighbor.Name
							}
							if oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress = oIpv6DhcpClientNeighborDiscoveryNeighbor.HwAddress
							}
							nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor = append(nestedIpv6.DhcpClient.NeighborDiscovery.Neighbor, nestedIpv6DhcpClientNeighborDiscoveryNeighbor)
						}
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.NsInterval = o.Ipv6.DhcpClient.NeighborDiscovery.NsInterval
					}
					if o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime != nil {
						nestedIpv6.DhcpClient.NeighborDiscovery.ReachableTime = o.Ipv6.DhcpClient.NeighborDiscovery.ReachableTime
					}
				}
				if o.Ipv6.DhcpClient.Preference != nil {
					nestedIpv6.DhcpClient.Preference = o.Ipv6.DhcpClient.Preference
				}
				if o.Ipv6.DhcpClient.PrefixDelegation != nil {
					nestedIpv6.DhcpClient.PrefixDelegation = &Ipv6DhcpClientPrefixDelegation{}
					if o.Ipv6.DhcpClient.PrefixDelegation.Misc != nil {
						entry.Misc["Ipv6DhcpClientPrefixDelegation"] = o.Ipv6.DhcpClient.PrefixDelegation.Misc
					}
					if o.Ipv6.DhcpClient.PrefixDelegation.Enable != nil {
						nestedIpv6.DhcpClient.PrefixDelegation.Enable = &Ipv6DhcpClientPrefixDelegationEnable{}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc != nil {
							entry.Misc["Ipv6DhcpClientPrefixDelegationEnable"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Misc
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.No = &Ipv6DhcpClientPrefixDelegationEnableNo{}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc != nil {
								entry.Misc["Ipv6DhcpClientPrefixDelegationEnableNo"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.No.Misc
							}
						}
						if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes != nil {
							nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes = &Ipv6DhcpClientPrefixDelegationEnableYes{}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc != nil {
								entry.Misc["Ipv6DhcpClientPrefixDelegationEnableYes"] = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.Misc
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PfxPoolName
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen = o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLen
							}
							if o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint != nil {
								nestedIpv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint = util.AsBool(o.Ipv6.DhcpClient.PrefixDelegation.Enable.Yes.PrefixLenHint, nil)
							}
						}
					}
				}
				if o.Ipv6.DhcpClient.V6Options != nil {
					nestedIpv6.DhcpClient.V6Options = &Ipv6DhcpClientV6Options{}
					if o.Ipv6.DhcpClient.V6Options.Misc != nil {
						entry.Misc["Ipv6DhcpClientV6Options"] = o.Ipv6.DhcpClient.V6Options.Misc
					}
					if o.Ipv6.DhcpClient.V6Options.DuidType != nil {
						nestedIpv6.DhcpClient.V6Options.DuidType = o.Ipv6.DhcpClient.V6Options.DuidType
					}
					if o.Ipv6.DhcpClient.V6Options.Enable != nil {
						nestedIpv6.DhcpClient.V6Options.Enable = &Ipv6DhcpClientV6OptionsEnable{}
						if o.Ipv6.DhcpClient.V6Options.Enable.Misc != nil {
							entry.Misc["Ipv6DhcpClientV6OptionsEnable"] = o.Ipv6.DhcpClient.V6Options.Enable.Misc
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.No != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.No = &Ipv6DhcpClientV6OptionsEnableNo{}
							if o.Ipv6.DhcpClient.V6Options.Enable.No.Misc != nil {
								entry.Misc["Ipv6DhcpClientV6OptionsEnableNo"] = o.Ipv6.DhcpClient.V6Options.Enable.No.Misc
							}
						}
						if o.Ipv6.DhcpClient.V6Options.Enable.Yes != nil {
							nestedIpv6.DhcpClient.V6Options.Enable.Yes = &Ipv6DhcpClientV6OptionsEnableYes{}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc != nil {
								entry.Misc["Ipv6DhcpClientV6OptionsEnableYes"] = o.Ipv6.DhcpClient.V6Options.Enable.Yes.Misc
							}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr != nil {
								nestedIpv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr = util.AsBool(o.Ipv6.DhcpClient.V6Options.Enable.Yes.NonTempAddr, nil)
							}
							if o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr != nil {
								nestedIpv6.DhcpClient.V6Options.Enable.Yes.TempAddr = util.AsBool(o.Ipv6.DhcpClient.V6Options.Enable.Yes.TempAddr, nil)
							}
						}
					}
					if o.Ipv6.DhcpClient.V6Options.RapidCommit != nil {
						nestedIpv6.DhcpClient.V6Options.RapidCommit = util.AsBool(o.Ipv6.DhcpClient.V6Options.RapidCommit, nil)
					}
					if o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig != nil {
						nestedIpv6.DhcpClient.V6Options.SupportSrvrReconfig = util.AsBool(o.Ipv6.DhcpClient.V6Options.SupportSrvrReconfig, nil)
					}
				}
			}
			if o.Ipv6.Inherited != nil {
				nestedIpv6.Inherited = &Ipv6Inherited{}
				if o.Ipv6.Inherited.Misc != nil {
					entry.Misc["Ipv6Inherited"] = o.Ipv6.Inherited.Misc
				}
				if o.Ipv6.Inherited.AssignAddr != nil {
					nestedIpv6.Inherited.AssignAddr = []Ipv6InheritedAssignAddr{}
					for _, oIpv6InheritedAssignAddr := range o.Ipv6.Inherited.AssignAddr {
						nestedIpv6InheritedAssignAddr := Ipv6InheritedAssignAddr{}
						if oIpv6InheritedAssignAddr.Misc != nil {
							entry.Misc["Ipv6InheritedAssignAddr"] = oIpv6InheritedAssignAddr.Misc
						}
						if oIpv6InheritedAssignAddr.Name != "" {
							nestedIpv6InheritedAssignAddr.Name = oIpv6InheritedAssignAddr.Name
						}
						if oIpv6InheritedAssignAddr.Type != nil {
							nestedIpv6InheritedAssignAddr.Type = &Ipv6InheritedAssignAddrType{}
							if oIpv6InheritedAssignAddr.Type.Misc != nil {
								entry.Misc["Ipv6InheritedAssignAddrType"] = oIpv6InheritedAssignAddr.Type.Misc
							}
							if oIpv6InheritedAssignAddr.Type.Gua != nil {
								nestedIpv6InheritedAssignAddr.Type.Gua = &Ipv6InheritedAssignAddrTypeGua{}
								if oIpv6InheritedAssignAddr.Type.Gua.Misc != nil {
									entry.Misc["Ipv6InheritedAssignAddrTypeGua"] = oIpv6InheritedAssignAddr.Type.Gua.Misc
								}
								if oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.EnableOnInterface, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PrefixPool != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PrefixPool = oIpv6InheritedAssignAddr.Type.Gua.PrefixPool
								}
								if oIpv6InheritedAssignAddr.Type.Gua.PoolType != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.PoolType = &Ipv6InheritedAssignAddrTypeGuaPoolType{}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolType"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc != nil {
											entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.Dynamic.Misc
										}
									}
									if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId = &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc != nil {
											entry.Misc["Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId"] = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Misc
										}
										if oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier != nil {
											nestedIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier = oIpv6InheritedAssignAddr.Type.Gua.PoolType.DynamicId.Identifier
										}
									}
								}
								if oIpv6InheritedAssignAddr.Type.Gua.Advertise != nil {
									nestedIpv6InheritedAssignAddr.Type.Gua.Advertise = &Ipv6InheritedAssignAddrTypeGuaAdvertise{}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeGuaAdvertise"] = oIpv6InheritedAssignAddr.Type.Gua.Advertise.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.Enable, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.OnlinkFlag, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Gua.Advertise.AutoConfigFlag, nil)
									}
								}
							}
							if oIpv6InheritedAssignAddr.Type.Ula != nil {
								nestedIpv6InheritedAssignAddr.Type.Ula = &Ipv6InheritedAssignAddrTypeUla{}
								if oIpv6InheritedAssignAddr.Type.Ula.Misc != nil {
									entry.Misc["Ipv6InheritedAssignAddrTypeUla"] = oIpv6InheritedAssignAddr.Type.Ula.Misc
								}
								if oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.EnableOnInterface, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Address != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Address = oIpv6InheritedAssignAddr.Type.Ula.Address
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Prefix != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Prefix = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Prefix, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Anycast != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Anycast = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Anycast, nil)
								}
								if oIpv6InheritedAssignAddr.Type.Ula.Advertise != nil {
									nestedIpv6InheritedAssignAddr.Type.Ula.Advertise = &Ipv6InheritedAssignAddrTypeUlaAdvertise{}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc != nil {
										entry.Misc["Ipv6InheritedAssignAddrTypeUlaAdvertise"] = oIpv6InheritedAssignAddr.Type.Ula.Advertise.Misc
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.Enable, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.ValidLifetime
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime = oIpv6InheritedAssignAddr.Type.Ula.Advertise.PreferredLifetime
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.OnlinkFlag, nil)
									}
									if oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag != nil {
										nestedIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag = util.AsBool(oIpv6InheritedAssignAddr.Type.Ula.Advertise.AutoConfigFlag, nil)
									}
								}
							}
						}
						nestedIpv6.Inherited.AssignAddr = append(nestedIpv6.Inherited.AssignAddr, nestedIpv6InheritedAssignAddr)
					}
				}
				if o.Ipv6.Inherited.Enable != nil {
					nestedIpv6.Inherited.Enable = util.AsBool(o.Ipv6.Inherited.Enable, nil)
				}
				if o.Ipv6.Inherited.NeighborDiscovery != nil {
					nestedIpv6.Inherited.NeighborDiscovery = &Ipv6InheritedNeighborDiscovery{}
					if o.Ipv6.Inherited.NeighborDiscovery.Misc != nil {
						entry.Misc["Ipv6InheritedNeighborDiscovery"] = o.Ipv6.Inherited.NeighborDiscovery.Misc
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DadAttempts != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DadAttempts = o.Ipv6.Inherited.NeighborDiscovery.DadAttempts
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsServer != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsServer = &Ipv6InheritedNeighborDiscoveryDnsServer{}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServer"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source = &Ipv6InheritedNeighborDiscoveryDnsServerSource{}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSource"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Misc
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManual"] = o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
									for _, oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := range o.Ipv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server {
										nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer := Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc != nil {
											entry.Misc["Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer"] = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Misc
										}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name != "" {
											nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Name
										}
										if oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime != nil {
											nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime = oIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer.Lifetime
										}
										nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server = append(nestedIpv6.Inherited.NeighborDiscovery.DnsServer.Source.Manual.Server, nestedIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer)
									}
								}
							}
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix != nil {
						nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix = &Ipv6InheritedNeighborDiscoveryDnsSuffix{}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffix"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source != nil {
							nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source = &Ipv6InheritedNeighborDiscoveryDnsSuffixSource{}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSource"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Misc
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6 = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Dhcpv6.PrefixPool
								}
							}
							if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual != nil {
								nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual = &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc != nil {
									entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual"] = o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Misc
								}
								if o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix != nil {
									nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
									for _, oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := range o.Ipv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix {
										nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix := Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc != nil {
											entry.Misc["Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix"] = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Misc
										}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name != "" {
											nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Name
										}
										if oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime != nil {
											nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime = oIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix.Lifetime
										}
										nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix = append(nestedIpv6.Inherited.NeighborDiscovery.DnsSuffix.Source.Manual.Suffix, nestedIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix)
									}
								}
							}
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.EnableDad != nil {
						nestedIpv6.Inherited.NeighborDiscovery.EnableDad = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.EnableDad, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedIpv6.Inherited.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Ipv6.Inherited.NeighborDiscovery.Neighbor != nil {
						nestedIpv6.Inherited.NeighborDiscovery.Neighbor = []Ipv6InheritedNeighborDiscoveryNeighbor{}
						for _, oIpv6InheritedNeighborDiscoveryNeighbor := range o.Ipv6.Inherited.NeighborDiscovery.Neighbor {
							nestedIpv6InheritedNeighborDiscoveryNeighbor := Ipv6InheritedNeighborDiscoveryNeighbor{}
							if oIpv6InheritedNeighborDiscoveryNeighbor.Misc != nil {
								entry.Misc["Ipv6InheritedNeighborDiscoveryNeighbor"] = oIpv6InheritedNeighborDiscoveryNeighbor.Misc
							}
							if oIpv6InheritedNeighborDiscoveryNeighbor.Name != "" {
								nestedIpv6InheritedNeighborDiscoveryNeighbor.Name = oIpv6InheritedNeighborDiscoveryNeighbor.Name
							}
							if oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress != nil {
								nestedIpv6InheritedNeighborDiscoveryNeighbor.HwAddress = oIpv6InheritedNeighborDiscoveryNeighbor.HwAddress
							}
							nestedIpv6.Inherited.NeighborDiscovery.Neighbor = append(nestedIpv6.Inherited.NeighborDiscovery.Neighbor, nestedIpv6InheritedNeighborDiscoveryNeighbor)
						}
					}
					if o.Ipv6.Inherited.NeighborDiscovery.NsInterval != nil {
						nestedIpv6.Inherited.NeighborDiscovery.NsInterval = o.Ipv6.Inherited.NeighborDiscovery.NsInterval
					}
					if o.Ipv6.Inherited.NeighborDiscovery.ReachableTime != nil {
						nestedIpv6.Inherited.NeighborDiscovery.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.ReachableTime
					}
					if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement != nil {
						nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement = &Ipv6InheritedNeighborDiscoveryRouterAdvertisement{}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc != nil {
							entry.Misc["Ipv6InheritedNeighborDiscoveryRouterAdvertisement"] = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Misc
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Enable, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.HopLimit
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.Lifetime
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.LinkMtu
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MaxInterval
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.MinInterval
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.ReachableTime
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RetransmissionTimer
						}
						if o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
							nestedIpv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Ipv6.Inherited.NeighborDiscovery.RouterAdvertisement.RouterPreference
						}
					}
				}
			}
		}
		entry.Ipv6 = nestedIpv6

		entry.Mtu = o.Mtu
		var nestedNdpProxy *NdpProxy
		if o.NdpProxy != nil {
			nestedNdpProxy = &NdpProxy{}
			if o.NdpProxy.Misc != nil {
				entry.Misc["NdpProxy"] = o.NdpProxy.Misc
			}
			if o.NdpProxy.Address != nil {
				nestedNdpProxy.Address = []NdpProxyAddress{}
				for _, oNdpProxyAddress := range o.NdpProxy.Address {
					nestedNdpProxyAddress := NdpProxyAddress{}
					if oNdpProxyAddress.Misc != nil {
						entry.Misc["NdpProxyAddress"] = oNdpProxyAddress.Misc
					}
					if oNdpProxyAddress.Name != "" {
						nestedNdpProxyAddress.Name = oNdpProxyAddress.Name
					}
					if oNdpProxyAddress.Negate != nil {
						nestedNdpProxyAddress.Negate = util.AsBool(oNdpProxyAddress.Negate, nil)
					}
					nestedNdpProxy.Address = append(nestedNdpProxy.Address, nestedNdpProxyAddress)
				}
			}
			if o.NdpProxy.Enabled != nil {
				nestedNdpProxy.Enabled = util.AsBool(o.NdpProxy.Enabled, nil)
			}
		}
		entry.NdpProxy = nestedNdpProxy

		entry.NetflowProfile = o.NetflowProfile

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
	if !matchAdjustTcpMss(a.AdjustTcpMss, b.AdjustTcpMss) {
		return false
	}
	if !matchArp(a.Arp, b.Arp) {
		return false
	}
	if !matchBonjour(a.Bonjour, b.Bonjour) {
		return false
	}
	if !util.StringsMatch(a.Comment, b.Comment) {
		return false
	}
	if !matchDdnsConfig(a.DdnsConfig, b.DdnsConfig) {
		return false
	}
	if !util.BoolsMatch(a.DfIgnore, b.DfIgnore) {
		return false
	}
	if !matchDhcpClient(a.DhcpClient, b.DhcpClient) {
		return false
	}
	if !util.StringsMatch(a.InterfaceManagementProfile, b.InterfaceManagementProfile) {
		return false
	}
	if !matchIp(a.Ip, b.Ip) {
		return false
	}
	if !matchIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	if !util.Ints64Match(a.Mtu, b.Mtu) {
		return false
	}
	if !matchNdpProxy(a.NdpProxy, b.NdpProxy) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
		return false
	}

	return true
}

func matchAdjustTcpMss(a *AdjustTcpMss, b *AdjustTcpMss) bool {
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
func matchArp(a []Arp, b []Arp) bool {
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
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
		}
	}
	return true
}
func matchBonjour(a *Bonjour, b *Bonjour) bool {
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
func matchDdnsConfigDdnsVendorConfig(a []DdnsConfigDdnsVendorConfig, b []DdnsConfigDdnsVendorConfig) bool {
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
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
		}
	}
	return true
}
func matchDdnsConfig(a *DdnsConfig, b *DdnsConfig) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DdnsCertProfile, b.DdnsCertProfile) {
		return false
	}
	if !util.BoolsMatch(a.DdnsEnabled, b.DdnsEnabled) {
		return false
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
	if !matchDdnsConfigDdnsVendorConfig(a.DdnsVendorConfig, b.DdnsVendorConfig) {
		return false
	}
	return true
}
func matchDhcpClientSendHostname(a *DhcpClientSendHostname, b *DhcpClientSendHostname) bool {
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
func matchDhcpClient(a *DhcpClient, b *DhcpClient) bool {
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
	if !matchDhcpClientSendHostname(a.SendHostname, b.SendHostname) {
		return false
	}
	return true
}
func matchIp(a []Ip, b []Ip) bool {
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
func matchIpv6AddressPrefix(a *Ipv6AddressPrefix, b *Ipv6AddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6AddressAnycast(a *Ipv6AddressAnycast, b *Ipv6AddressAnycast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6AddressAdvertise(a *Ipv6AddressAdvertise, b *Ipv6AddressAdvertise) bool {
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
func matchIpv6Address(a []Ipv6Address, b []Ipv6Address) bool {
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
			if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
				return false
			}
			if !matchIpv6AddressPrefix(a.Prefix, b.Prefix) {
				return false
			}
			if !matchIpv6AddressAnycast(a.Anycast, b.Anycast) {
				return false
			}
			if !matchIpv6AddressAdvertise(a.Advertise, b.Advertise) {
				return false
			}
		}
	}
	return true
}
func matchIpv6NeighborDiscoveryNeighbor(a []Ipv6NeighborDiscoveryNeighbor, b []Ipv6NeighborDiscoveryNeighbor) bool {
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
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
		}
	}
	return true
}
func matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer(a []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, b []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) bool {
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
func matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix(a []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, b []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) bool {
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
func matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupport(a *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, b *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupportServer(a.Server, b.Server) {
		return false
	}
	if !matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchIpv6NeighborDiscoveryRouterAdvertisement(a *Ipv6NeighborDiscoveryRouterAdvertisement, b *Ipv6NeighborDiscoveryRouterAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6NeighborDiscoveryRouterAdvertisementDnsSupport(a.DnsSupport, b.DnsSupport) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
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
	if !util.StringsMatch(a.LinkMtu, b.LinkMtu) {
		return false
	}
	if !util.BoolsMatch(a.ManagedFlag, b.ManagedFlag) {
		return false
	}
	if !util.Ints64Match(a.MaxInterval, b.MaxInterval) {
		return false
	}
	if !util.Ints64Match(a.MinInterval, b.MinInterval) {
		return false
	}
	if !util.BoolsMatch(a.OtherFlag, b.OtherFlag) {
		return false
	}
	if !util.StringsMatch(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.StringsMatch(a.RetransmissionTimer, b.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(a.RouterPreference, b.RouterPreference) {
		return false
	}
	return true
}
func matchIpv6NeighborDiscovery(a *Ipv6NeighborDiscovery, b *Ipv6NeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
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
	if !matchIpv6NeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !matchIpv6NeighborDiscoveryRouterAdvertisement(a.RouterAdvertisement, b.RouterAdvertisement) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6(a *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, b *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer(a []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, b []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) bool {
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
func matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceManual(a *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, b *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer(a.Server, b.Server) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsServerSource(a *Ipv6DhcpClientNeighborDiscoveryDnsServerSource, b *Ipv6DhcpClientNeighborDiscoveryDnsServerSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsServerSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsServer(a *Ipv6DhcpClientNeighborDiscoveryDnsServer, b *Ipv6DhcpClientNeighborDiscoveryDnsServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsServerSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6(a *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, b *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix(a []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, b []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
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
func matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual(a *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, b *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSource(a *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, b *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryDnsSuffix(a *Ipv6DhcpClientNeighborDiscoveryDnsSuffix, b *Ipv6DhcpClientNeighborDiscoveryDnsSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsSuffixSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscoveryNeighbor(a []Ipv6DhcpClientNeighborDiscoveryNeighbor, b []Ipv6DhcpClientNeighborDiscoveryNeighbor) bool {
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
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
		}
	}
	return true
}
func matchIpv6DhcpClientNeighborDiscovery(a *Ipv6DhcpClientNeighborDiscovery, b *Ipv6DhcpClientNeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsServer(a.DnsServer, b.DnsServer) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryDnsSuffix(a.DnsSuffix, b.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	if !matchIpv6DhcpClientNeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	return true
}
func matchIpv6DhcpClientPrefixDelegationEnableNo(a *Ipv6DhcpClientPrefixDelegationEnableNo, b *Ipv6DhcpClientPrefixDelegationEnableNo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6DhcpClientPrefixDelegationEnableYes(a *Ipv6DhcpClientPrefixDelegationEnableYes, b *Ipv6DhcpClientPrefixDelegationEnableYes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.PfxPoolName, b.PfxPoolName) {
		return false
	}
	if !util.Ints64Match(a.PrefixLen, b.PrefixLen) {
		return false
	}
	if !util.BoolsMatch(a.PrefixLenHint, b.PrefixLenHint) {
		return false
	}
	return true
}
func matchIpv6DhcpClientPrefixDelegationEnable(a *Ipv6DhcpClientPrefixDelegationEnable, b *Ipv6DhcpClientPrefixDelegationEnable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientPrefixDelegationEnableNo(a.No, b.No) {
		return false
	}
	if !matchIpv6DhcpClientPrefixDelegationEnableYes(a.Yes, b.Yes) {
		return false
	}
	return true
}
func matchIpv6DhcpClientPrefixDelegation(a *Ipv6DhcpClientPrefixDelegation, b *Ipv6DhcpClientPrefixDelegation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientPrefixDelegationEnable(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchIpv6DhcpClientV6OptionsEnableNo(a *Ipv6DhcpClientV6OptionsEnableNo, b *Ipv6DhcpClientV6OptionsEnableNo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6DhcpClientV6OptionsEnableYes(a *Ipv6DhcpClientV6OptionsEnableYes, b *Ipv6DhcpClientV6OptionsEnableYes) bool {
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
func matchIpv6DhcpClientV6OptionsEnable(a *Ipv6DhcpClientV6OptionsEnable, b *Ipv6DhcpClientV6OptionsEnable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6DhcpClientV6OptionsEnableNo(a.No, b.No) {
		return false
	}
	if !matchIpv6DhcpClientV6OptionsEnableYes(a.Yes, b.Yes) {
		return false
	}
	return true
}
func matchIpv6DhcpClientV6Options(a *Ipv6DhcpClientV6Options, b *Ipv6DhcpClientV6Options) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DuidType, b.DuidType) {
		return false
	}
	if !matchIpv6DhcpClientV6OptionsEnable(a.Enable, b.Enable) {
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
func matchIpv6DhcpClient(a *Ipv6DhcpClient, b *Ipv6DhcpClient) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
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
	if !matchIpv6DhcpClientNeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	if !util.StringsMatch(a.Preference, b.Preference) {
		return false
	}
	if !matchIpv6DhcpClientPrefixDelegation(a.PrefixDelegation, b.PrefixDelegation) {
		return false
	}
	if !matchIpv6DhcpClientV6Options(a.V6Options, b.V6Options) {
		return false
	}
	return true
}
func matchIpv6InheritedAssignAddrTypeGuaPoolTypeDynamic(a *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, b *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6InheritedAssignAddrTypeGuaPoolTypeDynamicId(a *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, b *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) bool {
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
func matchIpv6InheritedAssignAddrTypeGuaPoolType(a *Ipv6InheritedAssignAddrTypeGuaPoolType, b *Ipv6InheritedAssignAddrTypeGuaPoolType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedAssignAddrTypeGuaPoolTypeDynamic(a.Dynamic, b.Dynamic) {
		return false
	}
	if !matchIpv6InheritedAssignAddrTypeGuaPoolTypeDynamicId(a.DynamicId, b.DynamicId) {
		return false
	}
	return true
}
func matchIpv6InheritedAssignAddrTypeGuaAdvertise(a *Ipv6InheritedAssignAddrTypeGuaAdvertise, b *Ipv6InheritedAssignAddrTypeGuaAdvertise) bool {
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
func matchIpv6InheritedAssignAddrTypeGua(a *Ipv6InheritedAssignAddrTypeGua, b *Ipv6InheritedAssignAddrTypeGua) bool {
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
	if !matchIpv6InheritedAssignAddrTypeGuaPoolType(a.PoolType, b.PoolType) {
		return false
	}
	if !matchIpv6InheritedAssignAddrTypeGuaAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchIpv6InheritedAssignAddrTypeUlaAdvertise(a *Ipv6InheritedAssignAddrTypeUlaAdvertise, b *Ipv6InheritedAssignAddrTypeUlaAdvertise) bool {
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
func matchIpv6InheritedAssignAddrTypeUla(a *Ipv6InheritedAssignAddrTypeUla, b *Ipv6InheritedAssignAddrTypeUla) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
		return false
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
	if !matchIpv6InheritedAssignAddrTypeUlaAdvertise(a.Advertise, b.Advertise) {
		return false
	}
	return true
}
func matchIpv6InheritedAssignAddrType(a *Ipv6InheritedAssignAddrType, b *Ipv6InheritedAssignAddrType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedAssignAddrTypeGua(a.Gua, b.Gua) {
		return false
	}
	if !matchIpv6InheritedAssignAddrTypeUla(a.Ula, b.Ula) {
		return false
	}
	return true
}
func matchIpv6InheritedAssignAddr(a []Ipv6InheritedAssignAddr, b []Ipv6InheritedAssignAddr) bool {
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
			if !matchIpv6InheritedAssignAddrType(a.Type, b.Type) {
				return false
			}
		}
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6(a *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, b *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) bool {
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
func matchIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer(a []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, b []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) bool {
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
func matchIpv6InheritedNeighborDiscoveryDnsServerSourceManual(a *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, b *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsServerSourceManualServer(a.Server, b.Server) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsServerSource(a *Ipv6InheritedNeighborDiscoveryDnsServerSource, b *Ipv6InheritedNeighborDiscoveryDnsServerSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsServerSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsServer(a *Ipv6InheritedNeighborDiscoveryDnsServer, b *Ipv6InheritedNeighborDiscoveryDnsServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsServerSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6(a *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, b *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
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
func matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix(a []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, b []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
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
func matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceManual(a *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, b *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsSuffixSource(a *Ipv6InheritedNeighborDiscoveryDnsSuffixSource, b *Ipv6InheritedNeighborDiscoveryDnsSuffixSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsSuffixSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryDnsSuffix(a *Ipv6InheritedNeighborDiscoveryDnsSuffix, b *Ipv6InheritedNeighborDiscoveryDnsSuffix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsSuffixSource(a.Source, b.Source) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryNeighbor(a []Ipv6InheritedNeighborDiscoveryNeighbor, b []Ipv6InheritedNeighborDiscoveryNeighbor) bool {
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
			if !util.StringsMatch(a.HwAddress, b.HwAddress) {
				return false
			}
		}
	}
	return true
}
func matchIpv6InheritedNeighborDiscoveryRouterAdvertisement(a *Ipv6InheritedNeighborDiscoveryRouterAdvertisement, b *Ipv6InheritedNeighborDiscoveryRouterAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
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
	if !util.StringsMatch(a.LinkMtu, b.LinkMtu) {
		return false
	}
	if !util.BoolsMatch(a.ManagedFlag, b.ManagedFlag) {
		return false
	}
	if !util.Ints64Match(a.MaxInterval, b.MaxInterval) {
		return false
	}
	if !util.Ints64Match(a.MinInterval, b.MinInterval) {
		return false
	}
	if !util.BoolsMatch(a.OtherFlag, b.OtherFlag) {
		return false
	}
	if !util.StringsMatch(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.StringsMatch(a.RetransmissionTimer, b.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(a.RouterPreference, b.RouterPreference) {
		return false
	}
	return true
}
func matchIpv6InheritedNeighborDiscovery(a *Ipv6InheritedNeighborDiscovery, b *Ipv6InheritedNeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsServer(a.DnsServer, b.DnsServer) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryDnsSuffix(a.DnsSuffix, b.DnsSuffix) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscoveryRouterAdvertisement(a.RouterAdvertisement, b.RouterAdvertisement) {
		return false
	}
	return true
}
func matchIpv6Inherited(a *Ipv6Inherited, b *Ipv6Inherited) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6InheritedAssignAddr(a.AssignAddr, b.AssignAddr) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchIpv6InheritedNeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	return true
}
func matchIpv6(a *Ipv6, b *Ipv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchIpv6Address(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !util.StringsMatch(a.InterfaceId, b.InterfaceId) {
		return false
	}
	if !matchIpv6NeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	if !matchIpv6DhcpClient(a.DhcpClient, b.DhcpClient) {
		return false
	}
	if !matchIpv6Inherited(a.Inherited, b.Inherited) {
		return false
	}
	return true
}
func matchNdpProxyAddress(a []NdpProxyAddress, b []NdpProxyAddress) bool {
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
			if !util.BoolsMatch(a.Negate, b.Negate) {
				return false
			}
		}
	}
	return true
}
func matchNdpProxy(a *NdpProxy, b *NdpProxy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNdpProxyAddress(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
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
