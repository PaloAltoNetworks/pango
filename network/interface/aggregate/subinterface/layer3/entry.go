package layer3

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
	suffix = []string{"network", "interface", "aggregate-ethernet", "$parent", "layer3", "units", "$name"}
)

type Entry struct {
	Name                       string
	AdjustTcpMss               *AdjustTcpMss
	Arp                        []Arp
	Bonjour                    *Bonjour
	Comment                    *string
	DdnsConfig                 *DdnsConfig
	DecryptForward             *bool
	DfIgnore                   *bool
	DhcpClient                 *DhcpClient
	InterfaceManagementProfile *string
	Ip                         []Ip
	Ipv6                       *Ipv6
	Mtu                        *int64
	NdpProxy                   *NdpProxy
	NetflowProfile             *string
	SdwanLinkSettings          *SdwanLinkSettings
	Tag                        *int64
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}
type AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Arp struct {
	Name           string
	HwAddress      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Bonjour struct {
	Enable         *bool
	GroupId        *int64
	TtlCheck       *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type DdnsConfigDdnsVendorConfig struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DhcpClient struct {
	CreateDefaultRoute *bool
	DefaultRouteMetric *int64
	Enable             *bool
	SendHostname       *DhcpClientSendHostname
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type DhcpClientSendHostname struct {
	Enable         *bool
	Hostname       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ip struct {
	Name           string
	SdwanGateway   *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6 struct {
	Address           []Ipv6Address
	DhcpClient        *Ipv6DhcpClient
	Enabled           *bool
	Inherited         *Ipv6Inherited
	InterfaceId       *string
	NeighborDiscovery *Ipv6NeighborDiscovery
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6Address struct {
	Name              string
	EnableOnInterface *bool
	Prefix            *Ipv6AddressPrefix
	Anycast           *Ipv6AddressAnycast
	Advertise         *Ipv6AddressAdvertise
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6AddressPrefix struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6AddressAnycast struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6AddressAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6DhcpClient struct {
	AcceptRaRoute      *bool
	DefaultRouteMetric *int64
	Enable             *bool
	NeighborDiscovery  *Ipv6DhcpClientNeighborDiscovery
	Preference         *string
	PrefixDelegation   *Ipv6DhcpClientPrefixDelegation
	V6Options          *Ipv6DhcpClientV6Options
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
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
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsServer struct {
	Enable         *bool
	Source         *Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSource struct {
	Dhcpv6         *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	Manual         *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual struct {
	Server         []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffix struct {
	Enable         *bool
	Source         *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6         *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual         *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix         []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientNeighborDiscoveryNeighbor struct {
	Name           string
	HwAddress      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientPrefixDelegation struct {
	Enable         *Ipv6DhcpClientPrefixDelegationEnable
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientPrefixDelegationEnable struct {
	No             *Ipv6DhcpClientPrefixDelegationEnableNo
	Yes            *Ipv6DhcpClientPrefixDelegationEnableYes
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientPrefixDelegationEnableNo struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientPrefixDelegationEnableYes struct {
	PfxPoolName    *string
	PrefixLen      *int64
	PrefixLenHint  *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientV6Options struct {
	DuidType            *string
	Enable              *Ipv6DhcpClientV6OptionsEnable
	RapidCommit         *bool
	SupportSrvrReconfig *bool
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type Ipv6DhcpClientV6OptionsEnable struct {
	No             *Ipv6DhcpClientV6OptionsEnableNo
	Yes            *Ipv6DhcpClientV6OptionsEnableYes
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientV6OptionsEnableNo struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6DhcpClientV6OptionsEnableYes struct {
	NonTempAddr    *bool
	TempAddr       *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6Inherited struct {
	AssignAddr        []Ipv6InheritedAssignAddr
	Enable            *bool
	NeighborDiscovery *Ipv6InheritedNeighborDiscovery
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6InheritedAssignAddr struct {
	Name           string
	Type           *Ipv6InheritedAssignAddrType
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrType struct {
	Gua            *Ipv6InheritedAssignAddrTypeGua
	Ula            *Ipv6InheritedAssignAddrTypeUla
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrTypeGua struct {
	EnableOnInterface *bool
	PrefixPool        *string
	PoolType          *Ipv6InheritedAssignAddrTypeGuaPoolType
	Advertise         *Ipv6InheritedAssignAddrTypeGuaAdvertise
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6InheritedAssignAddrTypeGuaPoolType struct {
	Dynamic        *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	DynamicId      *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId struct {
	Identifier     *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrTypeGuaAdvertise struct {
	Enable         *bool
	OnlinkFlag     *bool
	AutoConfigFlag *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedAssignAddrTypeUla struct {
	EnableOnInterface *bool
	Address           *string
	Prefix            *bool
	Anycast           *bool
	Advertise         *Ipv6InheritedAssignAddrTypeUlaAdvertise
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6InheritedAssignAddrTypeUlaAdvertise struct {
	Enable            *bool
	ValidLifetime     *string
	PreferredLifetime *string
	OnlinkFlag        *bool
	AutoConfigFlag    *bool
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
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
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsServer struct {
	Enable         *bool
	Source         *Ipv6InheritedNeighborDiscoveryDnsServerSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsServerSource struct {
	Dhcpv6         *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	Manual         *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6 struct {
	PrefixPool     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManual struct {
	Server         []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsSuffix struct {
	Enable         *bool
	Source         *Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSource struct {
	Dhcpv6         *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	Manual         *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6 struct {
	PrefixPool     *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual struct {
	Suffix         []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6InheritedNeighborDiscoveryNeighbor struct {
	Name           string
	HwAddress      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
}
type Ipv6NeighborDiscovery struct {
	DadAttempts         *int64
	EnableDad           *bool
	EnableNdpMonitor    *bool
	Neighbor            []Ipv6NeighborDiscoveryNeighbor
	NsInterval          *int64
	ReachableTime       *int64
	RouterAdvertisement *Ipv6NeighborDiscoveryRouterAdvertisement
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type Ipv6NeighborDiscoveryNeighbor struct {
	Name           string
	HwAddress      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport struct {
	Enable         *bool
	Server         []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	Suffix         []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix struct {
	Name           string
	Lifetime       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NdpProxy struct {
	Address        []NdpProxyAddress
	Enabled        *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NdpProxyAddress struct {
	Name           string
	Negate         *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SdwanLinkSettings struct {
	Enable                *bool
	SdwanInterfaceProfile *string
	UpstreamNat           *SdwanLinkSettingsUpstreamNat
	Misc                  []generic.Xml
	MiscAttributes        []xml.Attr
}
type SdwanLinkSettingsUpstreamNat struct {
	Enable         *bool
	Ddns           *SdwanLinkSettingsUpstreamNatDdns
	StaticIp       *SdwanLinkSettingsUpstreamNatStaticIp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SdwanLinkSettingsUpstreamNatDdns struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SdwanLinkSettingsUpstreamNatStaticIp struct {
	Fqdn           *string
	IpAddress      *string
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
	XMLName                    xml.Name              `xml:"entry"`
	Name                       string                `xml:"name,attr"`
	AdjustTcpMss               *adjustTcpMssXml      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        *arpContainerXml      `xml:"arp,omitempty"`
	Bonjour                    *bonjourXml           `xml:"bonjour,omitempty"`
	Comment                    *string               `xml:"comment,omitempty"`
	DdnsConfig                 *ddnsConfigXml        `xml:"ddns-config,omitempty"`
	DecryptForward             *string               `xml:"decrypt-forward,omitempty"`
	DfIgnore                   *string               `xml:"df-ignore,omitempty"`
	DhcpClient                 *dhcpClientXml        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string               `xml:"interface-management-profile,omitempty"`
	Ip                         *ipContainerXml       `xml:"ip,omitempty"`
	Ipv6                       *ipv6Xml              `xml:"ipv6,omitempty"`
	Mtu                        *int64                `xml:"mtu,omitempty"`
	NdpProxy                   *ndpProxyXml          `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string               `xml:"netflow-profile,omitempty"`
	SdwanLinkSettings          *sdwanLinkSettingsXml `xml:"sdwan-link-settings,omitempty"`
	Tag                        *int64                `xml:"tag,omitempty"`
	Misc                       []generic.Xml         `xml:",any"`
	MiscAttributes             []xml.Attr            `xml:",any,attr"`
}
type adjustTcpMssXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64        `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64        `xml:"ipv6-mss-adjustment,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type arpContainerXml struct {
	Entries []arpXml `xml:"entry"`
}
type arpXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bonjourXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	GroupId        *int64        `xml:"group-id,omitempty"`
	TtlCheck       *string       `xml:"ttl-check,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ddnsConfigXml struct {
	DdnsCertProfile    *string                                 `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                                 `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                                 `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                        `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                        `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                  `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                                 `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   *ddnsConfigDdnsVendorConfigContainerXml `xml:"ddns-vendor-config,omitempty"`
	Misc               []generic.Xml                           `xml:",any"`
	MiscAttributes     []xml.Attr                              `xml:",any,attr"`
}
type ddnsConfigDdnsVendorConfigContainerXml struct {
	Entries []ddnsConfigDdnsVendorConfigXml `xml:"entry"`
}
type ddnsConfigDdnsVendorConfigXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type dhcpClientXml struct {
	CreateDefaultRoute *string                    `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                     `xml:"default-route-metric,omitempty"`
	Enable             *string                    `xml:"enable,omitempty"`
	SendHostname       *dhcpClientSendHostnameXml `xml:"send-hostname,omitempty"`
	Misc               []generic.Xml              `xml:",any"`
	MiscAttributes     []xml.Attr                 `xml:",any,attr"`
}
type dhcpClientSendHostnameXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	Hostname       *string       `xml:"hostname,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipContainerXml struct {
	Entries []ipXml `xml:"entry"`
}
type ipXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	SdwanGateway   *string       `xml:"sdwan-gateway,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml struct {
	Address           *ipv6AddressContainerXml  `xml:"address,omitempty"`
	DhcpClient        *ipv6DhcpClientXml        `xml:"dhcp-client,omitempty"`
	Enabled           *string                   `xml:"enabled,omitempty"`
	Inherited         *ipv6InheritedXml         `xml:"inherited,omitempty"`
	InterfaceId       *string                   `xml:"interface-id,omitempty"`
	NeighborDiscovery *ipv6NeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml             `xml:",any"`
	MiscAttributes    []xml.Attr                `xml:",any,attr"`
}
type ipv6AddressContainerXml struct {
	Entries []ipv6AddressXml `xml:"entry"`
}
type ipv6AddressXml struct {
	XMLName           xml.Name                 `xml:"entry"`
	Name              string                   `xml:"name,attr"`
	EnableOnInterface *string                  `xml:"enable-on-interface,omitempty"`
	Prefix            *ipv6AddressPrefixXml    `xml:"prefix,omitempty"`
	Anycast           *ipv6AddressAnycastXml   `xml:"anycast,omitempty"`
	Advertise         *ipv6AddressAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml            `xml:",any"`
	MiscAttributes    []xml.Attr               `xml:",any,attr"`
}
type ipv6AddressPrefixXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6AddressAnycastXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6AddressAdvertiseXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientXml struct {
	AcceptRaRoute      *string                             `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                              `xml:"default-route-metric,omitempty"`
	Enable             *string                             `xml:"enable,omitempty"`
	NeighborDiscovery  *ipv6DhcpClientNeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`
	Preference         *string                             `xml:"preference,omitempty"`
	PrefixDelegation   *ipv6DhcpClientPrefixDelegationXml  `xml:"prefix-delegation,omitempty"`
	V6Options          *ipv6DhcpClientV6OptionsXml         `xml:"v6-options,omitempty"`
	Misc               []generic.Xml                       `xml:",any"`
	MiscAttributes     []xml.Attr                          `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryXml struct {
	DadAttempts      *int64                                               `xml:"dad-attempts,omitempty"`
	DnsServer        *ipv6DhcpClientNeighborDiscoveryDnsServerXml         `xml:"dns-server,omitempty"`
	DnsSuffix        *ipv6DhcpClientNeighborDiscoveryDnsSuffixXml         `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                              `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                              `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         *ipv6DhcpClientNeighborDiscoveryNeighborContainerXml `xml:"neighbor,omitempty"`
	NsInterval       *int64                                               `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                               `xml:"reachable-time,omitempty"`
	Misc             []generic.Xml                                        `xml:",any"`
	MiscAttributes   []xml.Attr                                           `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerXml struct {
	Enable         *string                                            `xml:"enable,omitempty"`
	Source         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml struct {
	Server         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml                                                           `xml:",any"`
	MiscAttributes []xml.Attr                                                              `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml struct {
	Entries []ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixXml struct {
	Enable         *string                                            `xml:"enable,omitempty"`
	Source         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                           `xml:",any"`
	MiscAttributes []xml.Attr                                                              `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml struct {
	Entries []ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryNeighborContainerXml struct {
	Entries []ipv6DhcpClientNeighborDiscoveryNeighborXml `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryNeighborXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationXml struct {
	Enable         *ipv6DhcpClientPrefixDelegationEnableXml `xml:"enable,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableXml struct {
	No             *ipv6DhcpClientPrefixDelegationEnableNoXml  `xml:"no,omitempty"`
	Yes            *ipv6DhcpClientPrefixDelegationEnableYesXml `xml:"yes,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableNoXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableYesXml struct {
	PfxPoolName    *string       `xml:"pfx-pool-name,omitempty"`
	PrefixLen      *int64        `xml:"prefix-len,omitempty"`
	PrefixLenHint  *string       `xml:"prefix-len-hint,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsXml struct {
	DuidType            *string                           `xml:"duid-type,omitempty"`
	Enable              *ipv6DhcpClientV6OptionsEnableXml `xml:"enable,omitempty"`
	RapidCommit         *string                           `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                           `xml:"support-srvr-reconfig,omitempty"`
	Misc                []generic.Xml                     `xml:",any"`
	MiscAttributes      []xml.Attr                        `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableXml struct {
	No             *ipv6DhcpClientV6OptionsEnableNoXml  `xml:"no,omitempty"`
	Yes            *ipv6DhcpClientV6OptionsEnableYesXml `xml:"yes,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableNoXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableYesXml struct {
	NonTempAddr    *string       `xml:"non-temp-addr,omitempty"`
	TempAddr       *string       `xml:"temp-addr,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedXml struct {
	AssignAddr        *ipv6InheritedAssignAddrContainerXml `xml:"assign-addr,omitempty"`
	Enable            *string                              `xml:"enable,omitempty"`
	NeighborDiscovery *ipv6InheritedNeighborDiscoveryXml   `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                        `xml:",any"`
	MiscAttributes    []xml.Attr                           `xml:",any,attr"`
}
type ipv6InheritedAssignAddrContainerXml struct {
	Entries []ipv6InheritedAssignAddrXml `xml:"entry"`
}
type ipv6InheritedAssignAddrXml struct {
	XMLName        xml.Name                        `xml:"entry"`
	Name           string                          `xml:"name,attr"`
	Type           *ipv6InheritedAssignAddrTypeXml `xml:"type,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeXml struct {
	Gua            *ipv6InheritedAssignAddrTypeGuaXml `xml:"gua,omitempty"`
	Ula            *ipv6InheritedAssignAddrTypeUlaXml `xml:"ula,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaXml struct {
	EnableOnInterface *string                                     `xml:"enable-on-interface,omitempty"`
	PrefixPool        *string                                     `xml:"prefix-pool,omitempty"`
	PoolType          *ipv6InheritedAssignAddrTypeGuaPoolTypeXml  `xml:"pool-type,omitempty"`
	Advertise         *ipv6InheritedAssignAddrTypeGuaAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml                               `xml:",any"`
	MiscAttributes    []xml.Attr                                  `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeXml struct {
	Dynamic        *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml   `xml:"dynamic,omitempty"`
	DynamicId      *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml `xml:"dynamic-id,omitempty"`
	Misc           []generic.Xml                                       `xml:",any"`
	MiscAttributes []xml.Attr                                          `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml struct {
	Identifier     *int64        `xml:"identifier,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaAdvertiseXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	OnlinkFlag     *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag *string       `xml:"auto-config-flag,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeUlaXml struct {
	EnableOnInterface *string                                     `xml:"enable-on-interface,omitempty"`
	Address           *string                                     `xml:"address,omitempty"`
	Prefix            *string                                     `xml:"prefix,omitempty"`
	Anycast           *string                                     `xml:"anycast,omitempty"`
	Advertise         *ipv6InheritedAssignAddrTypeUlaAdvertiseXml `xml:"advertise,omitempty"`
	Misc              []generic.Xml                               `xml:",any"`
	MiscAttributes    []xml.Attr                                  `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeUlaAdvertiseXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryXml struct {
	DadAttempts         *int64                                                `xml:"dad-attempts,omitempty"`
	DnsServer           *ipv6InheritedNeighborDiscoveryDnsServerXml           `xml:"dns-server,omitempty"`
	DnsSuffix           *ipv6InheritedNeighborDiscoveryDnsSuffixXml           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                               `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                               `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *ipv6InheritedNeighborDiscoveryNeighborContainerXml   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                                `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                `xml:"reachable-time,omitempty"`
	RouterAdvertisement *ipv6InheritedNeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                         `xml:",any"`
	MiscAttributes      []xml.Attr                                            `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerXml struct {
	Enable         *string                                           `xml:"enable,omitempty"`
	Source         *ipv6InheritedNeighborDiscoveryDnsServerSourceXml `xml:"source,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceXml struct {
	Dhcpv6         *ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual         *ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml `xml:"manual,omitempty"`
	Misc           []generic.Xml                                           `xml:",any"`
	MiscAttributes []xml.Attr                                              `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml struct {
	PrefixPool     *string       `xml:"prefix-pool,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml struct {
	Server         *ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml                                                          `xml:",any"`
	MiscAttributes []xml.Attr                                                             `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml struct {
	Entries []ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixXml struct {
	Enable         *string                                           `xml:"enable,omitempty"`
	Source         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml `xml:"source,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml struct {
	Dhcpv6         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml `xml:"manual,omitempty"`
	Misc           []generic.Xml                                           `xml:",any"`
	MiscAttributes []xml.Attr                                              `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml struct {
	PrefixPool     *string       `xml:"prefix-pool,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml struct {
	Suffix         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                          `xml:",any"`
	MiscAttributes []xml.Attr                                                             `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml struct {
	Entries []ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryNeighborContainerXml struct {
	Entries []ipv6InheritedNeighborDiscoveryNeighborXml `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryNeighborXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryRouterAdvertisementXml struct {
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
	MiscAttributes         []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryXml struct {
	DadAttempts         *int64                                       `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                      `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                      `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *ipv6NeighborDiscoveryNeighborContainerXml   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                       `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                       `xml:"reachable-time,omitempty"`
	RouterAdvertisement *ipv6NeighborDiscoveryRouterAdvertisementXml `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                `xml:",any"`
	MiscAttributes      []xml.Attr                                   `xml:",any,attr"`
}
type ipv6NeighborDiscoveryNeighborContainerXml struct {
	Entries []ipv6NeighborDiscoveryNeighborXml `xml:"entry"`
}
type ipv6NeighborDiscoveryNeighborXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementXml struct {
	DnsSupport             *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml `xml:"dns-support,omitempty"`
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
	Misc                   []generic.Xml                                          `xml:",any"`
	MiscAttributes         []xml.Attr                                             `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml struct {
	Enable         *string                                                               `xml:"enable,omitempty"`
	Server         *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml `xml:"server,omitempty"`
	Suffix         *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                         `xml:",any"`
	MiscAttributes []xml.Attr                                                            `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml struct {
	Entries []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml `xml:"entry"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml struct {
	Entries []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml `xml:"entry"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ndpProxyXml struct {
	Address        *ndpProxyAddressContainerXml `xml:"address,omitempty"`
	Enabled        *string                      `xml:"enabled,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type ndpProxyAddressContainerXml struct {
	Entries []ndpProxyAddressXml `xml:"entry"`
}
type ndpProxyAddressXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Negate         *string       `xml:"negate,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type sdwanLinkSettingsXml struct {
	Enable                *string                          `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                          `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *sdwanLinkSettingsUpstreamNatXml `xml:"upstream-nat,omitempty"`
	Misc                  []generic.Xml                    `xml:",any"`
	MiscAttributes        []xml.Attr                       `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatXml struct {
	Enable         *string                                  `xml:"enable,omitempty"`
	Ddns           *sdwanLinkSettingsUpstreamNatDdnsXml     `xml:"ddns,omitempty"`
	StaticIp       *sdwanLinkSettingsUpstreamNatStaticIpXml `xml:"static-ip,omitempty"`
	Misc           []generic.Xml                            `xml:",any"`
	MiscAttributes []xml.Attr                               `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatDdnsXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatStaticIpXml struct {
	Fqdn           *string       `xml:"fqdn,omitempty"`
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name                     `xml:"entry"`
	Name                       string                       `xml:"name,attr"`
	AdjustTcpMss               *adjustTcpMssXml_11_0_2      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        *arpContainerXml_11_0_2      `xml:"arp,omitempty"`
	Bonjour                    *bonjourXml_11_0_2           `xml:"bonjour,omitempty"`
	Comment                    *string                      `xml:"comment,omitempty"`
	DdnsConfig                 *ddnsConfigXml_11_0_2        `xml:"ddns-config,omitempty"`
	DecryptForward             *string                      `xml:"decrypt-forward,omitempty"`
	DfIgnore                   *string                      `xml:"df-ignore,omitempty"`
	DhcpClient                 *dhcpClientXml_11_0_2        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                      `xml:"interface-management-profile,omitempty"`
	Ip                         *ipContainerXml_11_0_2       `xml:"ip,omitempty"`
	Ipv6                       *ipv6Xml_11_0_2              `xml:"ipv6,omitempty"`
	Mtu                        *int64                       `xml:"mtu,omitempty"`
	NdpProxy                   *ndpProxyXml_11_0_2          `xml:"ndp-proxy,omitempty"`
	NetflowProfile             *string                      `xml:"netflow-profile,omitempty"`
	SdwanLinkSettings          *sdwanLinkSettingsXml_11_0_2 `xml:"sdwan-link-settings,omitempty"`
	Tag                        *int64                       `xml:"tag,omitempty"`
	Misc                       []generic.Xml                `xml:",any"`
	MiscAttributes             []xml.Attr                   `xml:",any,attr"`
}
type adjustTcpMssXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64        `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64        `xml:"ipv6-mss-adjustment,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type arpContainerXml_11_0_2 struct {
	Entries []arpXml_11_0_2 `xml:"entry"`
}
type arpXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type bonjourXml_11_0_2 struct {
	Enable         *string       `xml:"enable,omitempty"`
	GroupId        *int64        `xml:"group-id,omitempty"`
	TtlCheck       *string       `xml:"ttl-check,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ddnsConfigXml_11_0_2 struct {
	DdnsCertProfile    *string                                        `xml:"ddns-cert-profile,omitempty"`
	DdnsEnabled        *string                                        `xml:"ddns-enabled,omitempty"`
	DdnsHostname       *string                                        `xml:"ddns-hostname,omitempty"`
	DdnsIp             *util.MemberType                               `xml:"ddns-ip,omitempty"`
	DdnsIpv6           *util.MemberType                               `xml:"ddns-ipv6,omitempty"`
	DdnsUpdateInterval *int64                                         `xml:"ddns-update-interval,omitempty"`
	DdnsVendor         *string                                        `xml:"ddns-vendor,omitempty"`
	DdnsVendorConfig   *ddnsConfigDdnsVendorConfigContainerXml_11_0_2 `xml:"ddns-vendor-config,omitempty"`
	Misc               []generic.Xml                                  `xml:",any"`
	MiscAttributes     []xml.Attr                                     `xml:",any,attr"`
}
type ddnsConfigDdnsVendorConfigContainerXml_11_0_2 struct {
	Entries []ddnsConfigDdnsVendorConfigXml_11_0_2 `xml:"entry"`
}
type ddnsConfigDdnsVendorConfigXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type dhcpClientXml_11_0_2 struct {
	CreateDefaultRoute *string                           `xml:"create-default-route,omitempty"`
	DefaultRouteMetric *int64                            `xml:"default-route-metric,omitempty"`
	Enable             *string                           `xml:"enable,omitempty"`
	SendHostname       *dhcpClientSendHostnameXml_11_0_2 `xml:"send-hostname,omitempty"`
	Misc               []generic.Xml                     `xml:",any"`
	MiscAttributes     []xml.Attr                        `xml:",any,attr"`
}
type dhcpClientSendHostnameXml_11_0_2 struct {
	Enable         *string       `xml:"enable,omitempty"`
	Hostname       *string       `xml:"hostname,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipContainerXml_11_0_2 struct {
	Entries []ipXml_11_0_2 `xml:"entry"`
}
type ipXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	SdwanGateway   *string       `xml:"sdwan-gateway,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml_11_0_2 struct {
	Address           *ipv6AddressContainerXml_11_0_2  `xml:"address,omitempty"`
	DhcpClient        *ipv6DhcpClientXml_11_0_2        `xml:"dhcp-client,omitempty"`
	Enabled           *string                          `xml:"enabled,omitempty"`
	Inherited         *ipv6InheritedXml_11_0_2         `xml:"inherited,omitempty"`
	InterfaceId       *string                          `xml:"interface-id,omitempty"`
	NeighborDiscovery *ipv6NeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                    `xml:",any"`
	MiscAttributes    []xml.Attr                       `xml:",any,attr"`
}
type ipv6AddressContainerXml_11_0_2 struct {
	Entries []ipv6AddressXml_11_0_2 `xml:"entry"`
}
type ipv6AddressXml_11_0_2 struct {
	XMLName           xml.Name                        `xml:"entry"`
	Name              string                          `xml:"name,attr"`
	EnableOnInterface *string                         `xml:"enable-on-interface,omitempty"`
	Prefix            *ipv6AddressPrefixXml_11_0_2    `xml:"prefix,omitempty"`
	Anycast           *ipv6AddressAnycastXml_11_0_2   `xml:"anycast,omitempty"`
	Advertise         *ipv6AddressAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                   `xml:",any"`
	MiscAttributes    []xml.Attr                      `xml:",any,attr"`
}
type ipv6AddressPrefixXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6AddressAnycastXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6AddressAdvertiseXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientXml_11_0_2 struct {
	AcceptRaRoute      *string                                    `xml:"accept-ra-route,omitempty"`
	DefaultRouteMetric *int64                                     `xml:"default-route-metric,omitempty"`
	Enable             *string                                    `xml:"enable,omitempty"`
	NeighborDiscovery  *ipv6DhcpClientNeighborDiscoveryXml_11_0_2 `xml:"neighbor-discovery,omitempty"`
	Preference         *string                                    `xml:"preference,omitempty"`
	PrefixDelegation   *ipv6DhcpClientPrefixDelegationXml_11_0_2  `xml:"prefix-delegation,omitempty"`
	V6Options          *ipv6DhcpClientV6OptionsXml_11_0_2         `xml:"v6-options,omitempty"`
	Misc               []generic.Xml                              `xml:",any"`
	MiscAttributes     []xml.Attr                                 `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts      *int64                                                      `xml:"dad-attempts,omitempty"`
	DnsServer        *ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2         `xml:"dns-server,omitempty"`
	DnsSuffix        *ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2         `xml:"dns-suffix,omitempty"`
	EnableDad        *string                                                     `xml:"enable-dad,omitempty"`
	EnableNdpMonitor *string                                                     `xml:"enable-ndp-monitor,omitempty"`
	Neighbor         *ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2 `xml:"neighbor,omitempty"`
	NsInterval       *int64                                                      `xml:"ns-interval,omitempty"`
	ReachableTime    *int64                                                      `xml:"reachable-time,omitempty"`
	Misc             []generic.Xml                                               `xml:",any"`
	MiscAttributes   []xml.Attr                                                  `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable         *string                                                   `xml:"enable,omitempty"`
	Source         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc           []generic.Xml                                             `xml:",any"`
	MiscAttributes []xml.Attr                                                `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc           []generic.Xml                                                   `xml:",any"`
	MiscAttributes []xml.Attr                                                      `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server         *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Misc           []generic.Xml                                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                                     `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 struct {
	Entries []ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable         *string                                                   `xml:"enable,omitempty"`
	Source         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc           []generic.Xml                                             `xml:",any"`
	MiscAttributes []xml.Attr                                                `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc           []generic.Xml                                                   `xml:",any"`
	MiscAttributes []xml.Attr                                                      `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix         *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                                     `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 struct {
	Entries []ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationXml_11_0_2 struct {
	Enable         *ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 `xml:"enable,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableXml_11_0_2 struct {
	No             *ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes            *ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 `xml:"yes,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2 struct {
	PfxPoolName    *string       `xml:"pfx-pool-name,omitempty"`
	PrefixLen      *int64        `xml:"prefix-len,omitempty"`
	PrefixLenHint  *string       `xml:"prefix-len-hint,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsXml_11_0_2 struct {
	DuidType            *string                                  `xml:"duid-type,omitempty"`
	Enable              *ipv6DhcpClientV6OptionsEnableXml_11_0_2 `xml:"enable,omitempty"`
	RapidCommit         *string                                  `xml:"rapid-commit,omitempty"`
	SupportSrvrReconfig *string                                  `xml:"support-srvr-reconfig,omitempty"`
	Misc                []generic.Xml                            `xml:",any"`
	MiscAttributes      []xml.Attr                               `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableXml_11_0_2 struct {
	No             *ipv6DhcpClientV6OptionsEnableNoXml_11_0_2  `xml:"no,omitempty"`
	Yes            *ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 `xml:"yes,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableNoXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6DhcpClientV6OptionsEnableYesXml_11_0_2 struct {
	NonTempAddr    *string       `xml:"non-temp-addr,omitempty"`
	TempAddr       *string       `xml:"temp-addr,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedXml_11_0_2 struct {
	AssignAddr        *ipv6InheritedAssignAddrContainerXml_11_0_2 `xml:"assign-addr,omitempty"`
	Enable            *string                                     `xml:"enable,omitempty"`
	NeighborDiscovery *ipv6InheritedNeighborDiscoveryXml_11_0_2   `xml:"neighbor-discovery,omitempty"`
	Misc              []generic.Xml                               `xml:",any"`
	MiscAttributes    []xml.Attr                                  `xml:",any,attr"`
}
type ipv6InheritedAssignAddrContainerXml_11_0_2 struct {
	Entries []ipv6InheritedAssignAddrXml_11_0_2 `xml:"entry"`
}
type ipv6InheritedAssignAddrXml_11_0_2 struct {
	XMLName        xml.Name                               `xml:"entry"`
	Name           string                                 `xml:"name,attr"`
	Type           *ipv6InheritedAssignAddrTypeXml_11_0_2 `xml:"type,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeXml_11_0_2 struct {
	Gua            *ipv6InheritedAssignAddrTypeGuaXml_11_0_2 `xml:"gua,omitempty"`
	Ula            *ipv6InheritedAssignAddrTypeUlaXml_11_0_2 `xml:"ula,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
	MiscAttributes []xml.Attr                                `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaXml_11_0_2 struct {
	EnableOnInterface *string                                            `xml:"enable-on-interface,omitempty"`
	PrefixPool        *string                                            `xml:"prefix-pool,omitempty"`
	PoolType          *ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2  `xml:"pool-type,omitempty"`
	Advertise         *ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                      `xml:",any"`
	MiscAttributes    []xml.Attr                                         `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2 struct {
	Dynamic        *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2   `xml:"dynamic,omitempty"`
	DynamicId      *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 `xml:"dynamic-id,omitempty"`
	Misc           []generic.Xml                                              `xml:",any"`
	MiscAttributes []xml.Attr                                                 `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2 struct {
	Identifier     *int64        `xml:"identifier,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2 struct {
	Enable         *string       `xml:"enable,omitempty"`
	OnlinkFlag     *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag *string       `xml:"auto-config-flag,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeUlaXml_11_0_2 struct {
	EnableOnInterface *string                                            `xml:"enable-on-interface,omitempty"`
	Address           *string                                            `xml:"address,omitempty"`
	Prefix            *string                                            `xml:"prefix,omitempty"`
	Anycast           *string                                            `xml:"anycast,omitempty"`
	Advertise         *ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 `xml:"advertise,omitempty"`
	Misc              []generic.Xml                                      `xml:",any"`
	MiscAttributes    []xml.Attr                                         `xml:",any,attr"`
}
type ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2 struct {
	Enable            *string       `xml:"enable,omitempty"`
	ValidLifetime     *string       `xml:"valid-lifetime,omitempty"`
	PreferredLifetime *string       `xml:"preferred-lifetime,omitempty"`
	OnlinkFlag        *string       `xml:"onlink-flag,omitempty"`
	AutoConfigFlag    *string       `xml:"auto-config-flag,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                                       `xml:"dad-attempts,omitempty"`
	DnsServer           *ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2           `xml:"dns-server,omitempty"`
	DnsSuffix           *ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2           `xml:"dns-suffix,omitempty"`
	EnableDad           *string                                                      `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                                      `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                                       `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                                       `xml:"reachable-time,omitempty"`
	RouterAdvertisement *ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                                `xml:",any"`
	MiscAttributes      []xml.Attr                                                   `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2 struct {
	Enable         *string                                                  `xml:"enable,omitempty"`
	Source         *ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2 struct {
	Dhcpv6         *ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual         *ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc           []generic.Xml                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                     `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool     *string       `xml:"prefix-pool,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2 struct {
	Server         *ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Misc           []generic.Xml                                                                 `xml:",any"`
	MiscAttributes []xml.Attr                                                                    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2 struct {
	Entries []ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2 struct {
	Enable         *string                                                  `xml:"enable,omitempty"`
	Source         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 `xml:"source,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2 struct {
	Dhcpv6         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 `xml:"dhcpv6,omitempty"`
	Manual         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 `xml:"manual,omitempty"`
	Misc           []generic.Xml                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                     `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2 struct {
	PrefixPool     *string       `xml:"prefix-pool,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2 struct {
	Suffix         *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                                 `xml:",any"`
	MiscAttributes []xml.Attr                                                                    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2 struct {
	Entries []ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
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
	MiscAttributes         []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryXml_11_0_2 struct {
	DadAttempts         *int64                                              `xml:"dad-attempts,omitempty"`
	EnableDad           *string                                             `xml:"enable-dad,omitempty"`
	EnableNdpMonitor    *string                                             `xml:"enable-ndp-monitor,omitempty"`
	Neighbor            *ipv6NeighborDiscoveryNeighborContainerXml_11_0_2   `xml:"neighbor,omitempty"`
	NsInterval          *int64                                              `xml:"ns-interval,omitempty"`
	ReachableTime       *int64                                              `xml:"reachable-time,omitempty"`
	RouterAdvertisement *ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 `xml:"router-advertisement,omitempty"`
	Misc                []generic.Xml                                       `xml:",any"`
	MiscAttributes      []xml.Attr                                          `xml:",any,attr"`
}
type ipv6NeighborDiscoveryNeighborContainerXml_11_0_2 struct {
	Entries []ipv6NeighborDiscoveryNeighborXml_11_0_2 `xml:"entry"`
}
type ipv6NeighborDiscoveryNeighborXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	HwAddress      *string       `xml:"hw-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2 struct {
	DnsSupport             *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 `xml:"dns-support,omitempty"`
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
	Misc                   []generic.Xml                                                 `xml:",any"`
	MiscAttributes         []xml.Attr                                                    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2 struct {
	Enable         *string                                                                      `xml:"enable,omitempty"`
	Server         *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2 `xml:"server,omitempty"`
	Suffix         *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2 `xml:"suffix,omitempty"`
	Misc           []generic.Xml                                                                `xml:",any"`
	MiscAttributes []xml.Attr                                                                   `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2 struct {
	Entries []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 `xml:"entry"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2 struct {
	Entries []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 `xml:"entry"`
}
type ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Lifetime       *int64        `xml:"lifetime,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ndpProxyXml_11_0_2 struct {
	Address        *ndpProxyAddressContainerXml_11_0_2 `xml:"address,omitempty"`
	Enabled        *string                             `xml:"enabled,omitempty"`
	Misc           []generic.Xml                       `xml:",any"`
	MiscAttributes []xml.Attr                          `xml:",any,attr"`
}
type ndpProxyAddressContainerXml_11_0_2 struct {
	Entries []ndpProxyAddressXml_11_0_2 `xml:"entry"`
}
type ndpProxyAddressXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Negate         *string       `xml:"negate,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type sdwanLinkSettingsXml_11_0_2 struct {
	Enable                *string                                 `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                                 `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *sdwanLinkSettingsUpstreamNatXml_11_0_2 `xml:"upstream-nat,omitempty"`
	Misc                  []generic.Xml                           `xml:",any"`
	MiscAttributes        []xml.Attr                              `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatXml_11_0_2 struct {
	Enable         *string                                         `xml:"enable,omitempty"`
	Ddns           *sdwanLinkSettingsUpstreamNatDdnsXml_11_0_2     `xml:"ddns,omitempty"`
	StaticIp       *sdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2 `xml:"static-ip,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatDdnsXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type sdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2 struct {
	Fqdn           *string       `xml:"fqdn,omitempty"`
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AdjustTcpMss != nil {
		var obj adjustTcpMssXml
		obj.MarshalFromObject(*s.AdjustTcpMss)
		o.AdjustTcpMss = &obj
	}
	if s.Arp != nil {
		var objs []arpXml
		for _, elt := range s.Arp {
			var obj arpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Arp = &arpContainerXml{Entries: objs}
	}
	if s.Bonjour != nil {
		var obj bonjourXml
		obj.MarshalFromObject(*s.Bonjour)
		o.Bonjour = &obj
	}
	o.Comment = s.Comment
	if s.DdnsConfig != nil {
		var obj ddnsConfigXml
		obj.MarshalFromObject(*s.DdnsConfig)
		o.DdnsConfig = &obj
	}
	o.DecryptForward = util.YesNo(s.DecryptForward, nil)
	o.DfIgnore = util.YesNo(s.DfIgnore, nil)
	if s.DhcpClient != nil {
		var obj dhcpClientXml
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.InterfaceManagementProfile = s.InterfaceManagementProfile
	if s.Ip != nil {
		var objs []ipXml
		for _, elt := range s.Ip {
			var obj ipXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ip = &ipContainerXml{Entries: objs}
	}
	if s.Ipv6 != nil {
		var obj ipv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Mtu = s.Mtu
	if s.NdpProxy != nil {
		var obj ndpProxyXml
		obj.MarshalFromObject(*s.NdpProxy)
		o.NdpProxy = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	if s.SdwanLinkSettings != nil {
		var obj sdwanLinkSettingsXml
		obj.MarshalFromObject(*s.SdwanLinkSettings)
		o.SdwanLinkSettings = &obj
	}
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var adjustTcpMssVal *AdjustTcpMss
	if o.AdjustTcpMss != nil {
		obj, err := o.AdjustTcpMss.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adjustTcpMssVal = obj
	}
	var arpVal []Arp
	if o.Arp != nil {
		for _, elt := range o.Arp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			arpVal = append(arpVal, *obj)
		}
	}
	var bonjourVal *Bonjour
	if o.Bonjour != nil {
		obj, err := o.Bonjour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bonjourVal = obj
	}
	var ddnsConfigVal *DdnsConfig
	if o.DdnsConfig != nil {
		obj, err := o.DdnsConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsConfigVal = obj
	}
	var dhcpClientVal *DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var ipVal []Ip
	if o.Ip != nil {
		for _, elt := range o.Ip.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipVal = append(ipVal, *obj)
		}
	}
	var ipv6Val *Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var ndpProxyVal *NdpProxy
	if o.NdpProxy != nil {
		obj, err := o.NdpProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ndpProxyVal = obj
	}
	var sdwanLinkSettingsVal *SdwanLinkSettings
	if o.SdwanLinkSettings != nil {
		obj, err := o.SdwanLinkSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanLinkSettingsVal = obj
	}

	result := &Entry{
		Name:                       o.Name,
		AdjustTcpMss:               adjustTcpMssVal,
		Arp:                        arpVal,
		Bonjour:                    bonjourVal,
		Comment:                    o.Comment,
		DdnsConfig:                 ddnsConfigVal,
		DecryptForward:             util.AsBool(o.DecryptForward, nil),
		DfIgnore:                   util.AsBool(o.DfIgnore, nil),
		DhcpClient:                 dhcpClientVal,
		InterfaceManagementProfile: o.InterfaceManagementProfile,
		Ip:                         ipVal,
		Ipv6:                       ipv6Val,
		Mtu:                        o.Mtu,
		NdpProxy:                   ndpProxyVal,
		NetflowProfile:             o.NetflowProfile,
		SdwanLinkSettings:          sdwanLinkSettingsVal,
		Tag:                        o.Tag,
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *adjustTcpMssXml) MarshalFromObject(s AdjustTcpMss) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Ipv4MssAdjustment = s.Ipv4MssAdjustment
	o.Ipv6MssAdjustment = s.Ipv6MssAdjustment
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o adjustTcpMssXml) UnmarshalToObject() (*AdjustTcpMss, error) {

	result := &AdjustTcpMss{
		Enable:            util.AsBool(o.Enable, nil),
		Ipv4MssAdjustment: o.Ipv4MssAdjustment,
		Ipv6MssAdjustment: o.Ipv6MssAdjustment,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *arpXml) MarshalFromObject(s Arp) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o arpXml) UnmarshalToObject() (*Arp, error) {

	result := &Arp{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bonjourXml) MarshalFromObject(s Bonjour) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GroupId = s.GroupId
	o.TtlCheck = util.YesNo(s.TtlCheck, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bonjourXml) UnmarshalToObject() (*Bonjour, error) {

	result := &Bonjour{
		Enable:         util.AsBool(o.Enable, nil),
		GroupId:        o.GroupId,
		TtlCheck:       util.AsBool(o.TtlCheck, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ddnsConfigXml) MarshalFromObject(s DdnsConfig) {
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
		var objs []ddnsConfigDdnsVendorConfigXml
		for _, elt := range s.DdnsVendorConfig {
			var obj ddnsConfigDdnsVendorConfigXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DdnsVendorConfig = &ddnsConfigDdnsVendorConfigContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ddnsConfigXml) UnmarshalToObject() (*DdnsConfig, error) {
	var ddnsIpVal []string
	if o.DdnsIp != nil {
		ddnsIpVal = util.MemToStr(o.DdnsIp)
	}
	var ddnsIpv6Val []string
	if o.DdnsIpv6 != nil {
		ddnsIpv6Val = util.MemToStr(o.DdnsIpv6)
	}
	var ddnsVendorConfigVal []DdnsConfigDdnsVendorConfig
	if o.DdnsVendorConfig != nil {
		for _, elt := range o.DdnsVendorConfig.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ddnsVendorConfigVal = append(ddnsVendorConfigVal, *obj)
		}
	}

	result := &DdnsConfig{
		DdnsCertProfile:    o.DdnsCertProfile,
		DdnsEnabled:        util.AsBool(o.DdnsEnabled, nil),
		DdnsHostname:       o.DdnsHostname,
		DdnsIp:             ddnsIpVal,
		DdnsIpv6:           ddnsIpv6Val,
		DdnsUpdateInterval: o.DdnsUpdateInterval,
		DdnsVendor:         o.DdnsVendor,
		DdnsVendorConfig:   ddnsVendorConfigVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ddnsConfigDdnsVendorConfigXml) MarshalFromObject(s DdnsConfigDdnsVendorConfig) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ddnsConfigDdnsVendorConfigXml) UnmarshalToObject() (*DdnsConfigDdnsVendorConfig, error) {

	result := &DdnsConfigDdnsVendorConfig{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *dhcpClientXml) MarshalFromObject(s DhcpClient) {
	o.CreateDefaultRoute = util.YesNo(s.CreateDefaultRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.SendHostname != nil {
		var obj dhcpClientSendHostnameXml
		obj.MarshalFromObject(*s.SendHostname)
		o.SendHostname = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o dhcpClientXml) UnmarshalToObject() (*DhcpClient, error) {
	var sendHostnameVal *DhcpClientSendHostname
	if o.SendHostname != nil {
		obj, err := o.SendHostname.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendHostnameVal = obj
	}

	result := &DhcpClient{
		CreateDefaultRoute: util.AsBool(o.CreateDefaultRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		SendHostname:       sendHostnameVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *dhcpClientSendHostnameXml) MarshalFromObject(s DhcpClientSendHostname) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Hostname = s.Hostname
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o dhcpClientSendHostnameXml) UnmarshalToObject() (*DhcpClientSendHostname, error) {

	result := &DhcpClientSendHostname{
		Enable:         util.AsBool(o.Enable, nil),
		Hostname:       o.Hostname,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipXml) MarshalFromObject(s Ip) {
	o.Name = s.Name
	o.SdwanGateway = s.SdwanGateway
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipXml) UnmarshalToObject() (*Ip, error) {

	result := &Ip{
		Name:           o.Name,
		SdwanGateway:   o.SdwanGateway,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6Xml) MarshalFromObject(s Ipv6) {
	if s.Address != nil {
		var objs []ipv6AddressXml
		for _, elt := range s.Address {
			var obj ipv6AddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &ipv6AddressContainerXml{Entries: objs}
	}
	if s.DhcpClient != nil {
		var obj ipv6DhcpClientXml
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Inherited != nil {
		var obj ipv6InheritedXml
		obj.MarshalFromObject(*s.Inherited)
		o.Inherited = &obj
	}
	o.InterfaceId = s.InterfaceId
	if s.NeighborDiscovery != nil {
		var obj ipv6NeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6Xml) UnmarshalToObject() (*Ipv6, error) {
	var addressVal []Ipv6Address
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}
	var dhcpClientVal *Ipv6DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var inheritedVal *Ipv6Inherited
	if o.Inherited != nil {
		obj, err := o.Inherited.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritedVal = obj
	}
	var neighborDiscoveryVal *Ipv6NeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Ipv6{
		Address:           addressVal,
		DhcpClient:        dhcpClientVal,
		Enabled:           util.AsBool(o.Enabled, nil),
		Inherited:         inheritedVal,
		InterfaceId:       o.InterfaceId,
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressXml) MarshalFromObject(s Ipv6Address) {
	o.Name = s.Name
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	if s.Prefix != nil {
		var obj ipv6AddressPrefixXml
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	if s.Anycast != nil {
		var obj ipv6AddressAnycastXml
		obj.MarshalFromObject(*s.Anycast)
		o.Anycast = &obj
	}
	if s.Advertise != nil {
		var obj ipv6AddressAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressXml) UnmarshalToObject() (*Ipv6Address, error) {
	var prefixVal *Ipv6AddressPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}
	var anycastVal *Ipv6AddressAnycast
	if o.Anycast != nil {
		obj, err := o.Anycast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anycastVal = obj
	}
	var advertiseVal *Ipv6AddressAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6Address{
		Name:              o.Name,
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Prefix:            prefixVal,
		Anycast:           anycastVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressPrefixXml) MarshalFromObject(s Ipv6AddressPrefix) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressPrefixXml) UnmarshalToObject() (*Ipv6AddressPrefix, error) {

	result := &Ipv6AddressPrefix{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressAnycastXml) MarshalFromObject(s Ipv6AddressAnycast) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressAnycastXml) UnmarshalToObject() (*Ipv6AddressAnycast, error) {

	result := &Ipv6AddressAnycast{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressAdvertiseXml) MarshalFromObject(s Ipv6AddressAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressAdvertiseXml) UnmarshalToObject() (*Ipv6AddressAdvertise, error) {

	result := &Ipv6AddressAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientXml) MarshalFromObject(s Ipv6DhcpClient) {
	o.AcceptRaRoute = util.YesNo(s.AcceptRaRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj ipv6DhcpClientNeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Preference = s.Preference
	if s.PrefixDelegation != nil {
		var obj ipv6DhcpClientPrefixDelegationXml
		obj.MarshalFromObject(*s.PrefixDelegation)
		o.PrefixDelegation = &obj
	}
	if s.V6Options != nil {
		var obj ipv6DhcpClientV6OptionsXml
		obj.MarshalFromObject(*s.V6Options)
		o.V6Options = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientXml) UnmarshalToObject() (*Ipv6DhcpClient, error) {
	var neighborDiscoveryVal *Ipv6DhcpClientNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}
	var prefixDelegationVal *Ipv6DhcpClientPrefixDelegation
	if o.PrefixDelegation != nil {
		obj, err := o.PrefixDelegation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixDelegationVal = obj
	}
	var v6OptionsVal *Ipv6DhcpClientV6Options
	if o.V6Options != nil {
		obj, err := o.V6Options.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v6OptionsVal = obj
	}

	result := &Ipv6DhcpClient{
		AcceptRaRoute:      util.AsBool(o.AcceptRaRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		NeighborDiscovery:  neighborDiscoveryVal,
		Preference:         o.Preference,
		PrefixDelegation:   prefixDelegationVal,
		V6Options:          v6OptionsVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerXml
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixXml
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj ipv6DhcpClientNeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6DhcpClientNeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscovery, error) {
	var dnsServerVal *Ipv6DhcpClientNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Ipv6DhcpClientNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscovery{
		DadAttempts:      o.DadAttempts,
		DnsServer:        dnsServerVal,
		DnsSuffix:        dnsSuffixVal,
		EnableDad:        util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor: util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:         neighborVal,
		NsInterval:       o.NsInterval,
		ReachableTime:    o.ReachableTime,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServer, error) {
	var sourceVal *Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServer{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml
		for _, elt := range s.Server {
			var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffix{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml
		for _, elt := range s.Suffix {
			var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryNeighborXml) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryNeighborXml) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryNeighbor, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationXml) MarshalFromObject(s Ipv6DhcpClientPrefixDelegation) {
	if s.Enable != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableXml
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationXml) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegation, error) {
	var enableVal *Ipv6DhcpClientPrefixDelegationEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Ipv6DhcpClientPrefixDelegation{
		Enable:         enableVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableXml) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnable) {
	if s.No != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableXml) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnable, error) {
	var noVal *Ipv6DhcpClientPrefixDelegationEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Ipv6DhcpClientPrefixDelegationEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Ipv6DhcpClientPrefixDelegationEnable{
		No:             noVal,
		Yes:            yesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableNoXml) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnableNo) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableNoXml) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnableNo, error) {

	result := &Ipv6DhcpClientPrefixDelegationEnableNo{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableYesXml) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnableYes) {
	o.PfxPoolName = s.PfxPoolName
	o.PrefixLen = s.PrefixLen
	o.PrefixLenHint = util.YesNo(s.PrefixLenHint, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableYesXml) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnableYes, error) {

	result := &Ipv6DhcpClientPrefixDelegationEnableYes{
		PfxPoolName:    o.PfxPoolName,
		PrefixLen:      o.PrefixLen,
		PrefixLenHint:  util.AsBool(o.PrefixLenHint, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsXml) MarshalFromObject(s Ipv6DhcpClientV6Options) {
	o.DuidType = s.DuidType
	if s.Enable != nil {
		var obj ipv6DhcpClientV6OptionsEnableXml
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.RapidCommit = util.YesNo(s.RapidCommit, nil)
	o.SupportSrvrReconfig = util.YesNo(s.SupportSrvrReconfig, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsXml) UnmarshalToObject() (*Ipv6DhcpClientV6Options, error) {
	var enableVal *Ipv6DhcpClientV6OptionsEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Ipv6DhcpClientV6Options{
		DuidType:            o.DuidType,
		Enable:              enableVal,
		RapidCommit:         util.AsBool(o.RapidCommit, nil),
		SupportSrvrReconfig: util.AsBool(o.SupportSrvrReconfig, nil),
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableXml) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnable) {
	if s.No != nil {
		var obj ipv6DhcpClientV6OptionsEnableNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj ipv6DhcpClientV6OptionsEnableYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableXml) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnable, error) {
	var noVal *Ipv6DhcpClientV6OptionsEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Ipv6DhcpClientV6OptionsEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Ipv6DhcpClientV6OptionsEnable{
		No:             noVal,
		Yes:            yesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableNoXml) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnableNo) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableNoXml) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnableNo, error) {

	result := &Ipv6DhcpClientV6OptionsEnableNo{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableYesXml) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnableYes) {
	o.NonTempAddr = util.YesNo(s.NonTempAddr, nil)
	o.TempAddr = util.YesNo(s.TempAddr, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableYesXml) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnableYes, error) {

	result := &Ipv6DhcpClientV6OptionsEnableYes{
		NonTempAddr:    util.AsBool(o.NonTempAddr, nil),
		TempAddr:       util.AsBool(o.TempAddr, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedXml) MarshalFromObject(s Ipv6Inherited) {
	if s.AssignAddr != nil {
		var objs []ipv6InheritedAssignAddrXml
		for _, elt := range s.AssignAddr {
			var obj ipv6InheritedAssignAddrXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AssignAddr = &ipv6InheritedAssignAddrContainerXml{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj ipv6InheritedNeighborDiscoveryXml
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedXml) UnmarshalToObject() (*Ipv6Inherited, error) {
	var assignAddrVal []Ipv6InheritedAssignAddr
	if o.AssignAddr != nil {
		for _, elt := range o.AssignAddr.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			assignAddrVal = append(assignAddrVal, *obj)
		}
	}
	var neighborDiscoveryVal *Ipv6InheritedNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Ipv6Inherited{
		AssignAddr:        assignAddrVal,
		Enable:            util.AsBool(o.Enable, nil),
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrXml) MarshalFromObject(s Ipv6InheritedAssignAddr) {
	o.Name = s.Name
	if s.Type != nil {
		var obj ipv6InheritedAssignAddrTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrXml) UnmarshalToObject() (*Ipv6InheritedAssignAddr, error) {
	var typeVal *Ipv6InheritedAssignAddrType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Ipv6InheritedAssignAddr{
		Name:           o.Name,
		Type:           typeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeXml) MarshalFromObject(s Ipv6InheritedAssignAddrType) {
	if s.Gua != nil {
		var obj ipv6InheritedAssignAddrTypeGuaXml
		obj.MarshalFromObject(*s.Gua)
		o.Gua = &obj
	}
	if s.Ula != nil {
		var obj ipv6InheritedAssignAddrTypeUlaXml
		obj.MarshalFromObject(*s.Ula)
		o.Ula = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrType, error) {
	var guaVal *Ipv6InheritedAssignAddrTypeGua
	if o.Gua != nil {
		obj, err := o.Gua.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		guaVal = obj
	}
	var ulaVal *Ipv6InheritedAssignAddrTypeUla
	if o.Ula != nil {
		obj, err := o.Ula.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ulaVal = obj
	}

	result := &Ipv6InheritedAssignAddrType{
		Gua:            guaVal,
		Ula:            ulaVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGua) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.PrefixPool = s.PrefixPool
	if s.PoolType != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeXml
		obj.MarshalFromObject(*s.PoolType)
		o.PoolType = &obj
	}
	if s.Advertise != nil {
		var obj ipv6InheritedAssignAddrTypeGuaAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGua, error) {
	var poolTypeVal *Ipv6InheritedAssignAddrTypeGuaPoolType
	if o.PoolType != nil {
		obj, err := o.PoolType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		poolTypeVal = obj
	}
	var advertiseVal *Ipv6InheritedAssignAddrTypeGuaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeGua{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		PrefixPool:        o.PrefixPool,
		PoolType:          poolTypeVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolType) {
	if s.Dynamic != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.DynamicId != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml
		obj.MarshalFromObject(*s.DynamicId)
		o.DynamicId = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolType, error) {
	var dynamicVal *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var dynamicIdVal *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	if o.DynamicId != nil {
		obj, err := o.DynamicId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIdVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeGuaPoolType{
		Dynamic:        dynamicVal,
		DynamicId:      dynamicIdVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) {
	o.Identifier = s.Identifier
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{
		Identifier:     o.Identifier,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaAdvertiseXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaAdvertiseXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaAdvertise, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaAdvertise{
		Enable:         util.AsBool(o.Enable, nil),
		OnlinkFlag:     util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag: util.AsBool(o.AutoConfigFlag, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeUlaXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeUla) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.Address = s.Address
	o.Prefix = util.YesNo(s.Prefix, nil)
	o.Anycast = util.YesNo(s.Anycast, nil)
	if s.Advertise != nil {
		var obj ipv6InheritedAssignAddrTypeUlaAdvertiseXml
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeUlaXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeUla, error) {
	var advertiseVal *Ipv6InheritedAssignAddrTypeUlaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeUla{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Address:           o.Address,
		Prefix:            util.AsBool(o.Prefix, nil),
		Anycast:           util.AsBool(o.Anycast, nil),
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeUlaAdvertiseXml) MarshalFromObject(s Ipv6InheritedAssignAddrTypeUlaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeUlaAdvertiseXml) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeUlaAdvertise, error) {

	result := &Ipv6InheritedAssignAddrTypeUlaAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryXml) MarshalFromObject(s Ipv6InheritedNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerXml
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixXml
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6InheritedNeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj ipv6InheritedNeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6InheritedNeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj ipv6InheritedNeighborDiscoveryRouterAdvertisementXml
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscovery, error) {
	var dnsServerVal *Ipv6InheritedNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Ipv6InheritedNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Ipv6InheritedNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Ipv6InheritedNeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Ipv6InheritedNeighborDiscovery{
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
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServer, error) {
	var sourceVal *Ipv6InheritedNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServer{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{
		PrefixPool:     o.PrefixPool,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml
		for _, elt := range s.Server {
			var obj ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffix{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{
		PrefixPool:     o.PrefixPool,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml
		for _, elt := range s.Suffix {
			var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryNeighborXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryNeighborXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryNeighbor, error) {

	result := &Ipv6InheritedNeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryRouterAdvertisementXml) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryRouterAdvertisement) {
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
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryRouterAdvertisementXml) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryRouterAdvertisement, error) {

	result := &Ipv6InheritedNeighborDiscoveryRouterAdvertisement{
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
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryXml) MarshalFromObject(s Ipv6NeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6NeighborDiscoveryNeighborXml
		for _, elt := range s.Neighbor {
			var obj ipv6NeighborDiscoveryNeighborXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6NeighborDiscoveryNeighborContainerXml{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj ipv6NeighborDiscoveryRouterAdvertisementXml
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryXml) UnmarshalToObject() (*Ipv6NeighborDiscovery, error) {
	var neighborVal []Ipv6NeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Ipv6NeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Ipv6NeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryNeighborXml) MarshalFromObject(s Ipv6NeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryNeighborXml) UnmarshalToObject() (*Ipv6NeighborDiscoveryNeighbor, error) {

	result := &Ipv6NeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementXml) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisement) {
	if s.DnsSupport != nil {
		var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml
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
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementXml) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisement, error) {
	var dnsSupportVal *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
	if o.DnsSupport != nil {
		obj, err := o.DnsSupport.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSupportVal = obj
	}

	result := &Ipv6NeighborDiscoveryRouterAdvertisement{
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
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Server != nil {
		var objs []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml
		for _, elt := range s.Server {
			var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml{Entries: objs}
	}
	if s.Suffix != nil {
		var objs []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml
		for _, elt := range s.Suffix {
			var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, error) {
	var serverVal []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}
	var suffixVal []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{
		Enable:         util.AsBool(o.Enable, nil),
		Server:         serverVal,
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, error) {

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, error) {

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ndpProxyXml) MarshalFromObject(s NdpProxy) {
	if s.Address != nil {
		var objs []ndpProxyAddressXml
		for _, elt := range s.Address {
			var obj ndpProxyAddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &ndpProxyAddressContainerXml{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ndpProxyXml) UnmarshalToObject() (*NdpProxy, error) {
	var addressVal []NdpProxyAddress
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &NdpProxy{
		Address:        addressVal,
		Enabled:        util.AsBool(o.Enabled, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ndpProxyAddressXml) MarshalFromObject(s NdpProxyAddress) {
	o.Name = s.Name
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ndpProxyAddressXml) UnmarshalToObject() (*NdpProxyAddress, error) {

	result := &NdpProxyAddress{
		Name:           o.Name,
		Negate:         util.AsBool(o.Negate, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsXml) MarshalFromObject(s SdwanLinkSettings) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	if s.UpstreamNat != nil {
		var obj sdwanLinkSettingsUpstreamNatXml
		obj.MarshalFromObject(*s.UpstreamNat)
		o.UpstreamNat = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsXml) UnmarshalToObject() (*SdwanLinkSettings, error) {
	var upstreamNatVal *SdwanLinkSettingsUpstreamNat
	if o.UpstreamNat != nil {
		obj, err := o.UpstreamNat.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		upstreamNatVal = obj
	}

	result := &SdwanLinkSettings{
		Enable:                util.AsBool(o.Enable, nil),
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		UpstreamNat:           upstreamNatVal,
		Misc:                  o.Misc,
		MiscAttributes:        o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatXml) MarshalFromObject(s SdwanLinkSettingsUpstreamNat) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Ddns != nil {
		var obj sdwanLinkSettingsUpstreamNatDdnsXml
		obj.MarshalFromObject(*s.Ddns)
		o.Ddns = &obj
	}
	if s.StaticIp != nil {
		var obj sdwanLinkSettingsUpstreamNatStaticIpXml
		obj.MarshalFromObject(*s.StaticIp)
		o.StaticIp = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatXml) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNat, error) {
	var ddnsVal *SdwanLinkSettingsUpstreamNatDdns
	if o.Ddns != nil {
		obj, err := o.Ddns.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsVal = obj
	}
	var staticIpVal *SdwanLinkSettingsUpstreamNatStaticIp
	if o.StaticIp != nil {
		obj, err := o.StaticIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticIpVal = obj
	}

	result := &SdwanLinkSettingsUpstreamNat{
		Enable:         util.AsBool(o.Enable, nil),
		Ddns:           ddnsVal,
		StaticIp:       staticIpVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatDdnsXml) MarshalFromObject(s SdwanLinkSettingsUpstreamNatDdns) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatDdnsXml) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNatDdns, error) {

	result := &SdwanLinkSettingsUpstreamNatDdns{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatStaticIpXml) MarshalFromObject(s SdwanLinkSettingsUpstreamNatStaticIp) {
	o.Fqdn = s.Fqdn
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatStaticIpXml) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNatStaticIp, error) {

	result := &SdwanLinkSettingsUpstreamNatStaticIp{
		Fqdn:           o.Fqdn,
		IpAddress:      o.IpAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AdjustTcpMss != nil {
		var obj adjustTcpMssXml_11_0_2
		obj.MarshalFromObject(*s.AdjustTcpMss)
		o.AdjustTcpMss = &obj
	}
	if s.Arp != nil {
		var objs []arpXml_11_0_2
		for _, elt := range s.Arp {
			var obj arpXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Arp = &arpContainerXml_11_0_2{Entries: objs}
	}
	if s.Bonjour != nil {
		var obj bonjourXml_11_0_2
		obj.MarshalFromObject(*s.Bonjour)
		o.Bonjour = &obj
	}
	o.Comment = s.Comment
	if s.DdnsConfig != nil {
		var obj ddnsConfigXml_11_0_2
		obj.MarshalFromObject(*s.DdnsConfig)
		o.DdnsConfig = &obj
	}
	o.DecryptForward = util.YesNo(s.DecryptForward, nil)
	o.DfIgnore = util.YesNo(s.DfIgnore, nil)
	if s.DhcpClient != nil {
		var obj dhcpClientXml_11_0_2
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.InterfaceManagementProfile = s.InterfaceManagementProfile
	if s.Ip != nil {
		var objs []ipXml_11_0_2
		for _, elt := range s.Ip {
			var obj ipXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ip = &ipContainerXml_11_0_2{Entries: objs}
	}
	if s.Ipv6 != nil {
		var obj ipv6Xml_11_0_2
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Mtu = s.Mtu
	if s.NdpProxy != nil {
		var obj ndpProxyXml_11_0_2
		obj.MarshalFromObject(*s.NdpProxy)
		o.NdpProxy = &obj
	}
	o.NetflowProfile = s.NetflowProfile
	if s.SdwanLinkSettings != nil {
		var obj sdwanLinkSettingsXml_11_0_2
		obj.MarshalFromObject(*s.SdwanLinkSettings)
		o.SdwanLinkSettings = &obj
	}
	o.Tag = s.Tag
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var adjustTcpMssVal *AdjustTcpMss
	if o.AdjustTcpMss != nil {
		obj, err := o.AdjustTcpMss.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adjustTcpMssVal = obj
	}
	var arpVal []Arp
	if o.Arp != nil {
		for _, elt := range o.Arp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			arpVal = append(arpVal, *obj)
		}
	}
	var bonjourVal *Bonjour
	if o.Bonjour != nil {
		obj, err := o.Bonjour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bonjourVal = obj
	}
	var ddnsConfigVal *DdnsConfig
	if o.DdnsConfig != nil {
		obj, err := o.DdnsConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsConfigVal = obj
	}
	var dhcpClientVal *DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var ipVal []Ip
	if o.Ip != nil {
		for _, elt := range o.Ip.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipVal = append(ipVal, *obj)
		}
	}
	var ipv6Val *Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}
	var ndpProxyVal *NdpProxy
	if o.NdpProxy != nil {
		obj, err := o.NdpProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ndpProxyVal = obj
	}
	var sdwanLinkSettingsVal *SdwanLinkSettings
	if o.SdwanLinkSettings != nil {
		obj, err := o.SdwanLinkSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanLinkSettingsVal = obj
	}

	result := &Entry{
		Name:                       o.Name,
		AdjustTcpMss:               adjustTcpMssVal,
		Arp:                        arpVal,
		Bonjour:                    bonjourVal,
		Comment:                    o.Comment,
		DdnsConfig:                 ddnsConfigVal,
		DecryptForward:             util.AsBool(o.DecryptForward, nil),
		DfIgnore:                   util.AsBool(o.DfIgnore, nil),
		DhcpClient:                 dhcpClientVal,
		InterfaceManagementProfile: o.InterfaceManagementProfile,
		Ip:                         ipVal,
		Ipv6:                       ipv6Val,
		Mtu:                        o.Mtu,
		NdpProxy:                   ndpProxyVal,
		NetflowProfile:             o.NetflowProfile,
		SdwanLinkSettings:          sdwanLinkSettingsVal,
		Tag:                        o.Tag,
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *adjustTcpMssXml_11_0_2) MarshalFromObject(s AdjustTcpMss) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Ipv4MssAdjustment = s.Ipv4MssAdjustment
	o.Ipv6MssAdjustment = s.Ipv6MssAdjustment
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o adjustTcpMssXml_11_0_2) UnmarshalToObject() (*AdjustTcpMss, error) {

	result := &AdjustTcpMss{
		Enable:            util.AsBool(o.Enable, nil),
		Ipv4MssAdjustment: o.Ipv4MssAdjustment,
		Ipv6MssAdjustment: o.Ipv6MssAdjustment,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *arpXml_11_0_2) MarshalFromObject(s Arp) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o arpXml_11_0_2) UnmarshalToObject() (*Arp, error) {

	result := &Arp{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *bonjourXml_11_0_2) MarshalFromObject(s Bonjour) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.GroupId = s.GroupId
	o.TtlCheck = util.YesNo(s.TtlCheck, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o bonjourXml_11_0_2) UnmarshalToObject() (*Bonjour, error) {

	result := &Bonjour{
		Enable:         util.AsBool(o.Enable, nil),
		GroupId:        o.GroupId,
		TtlCheck:       util.AsBool(o.TtlCheck, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ddnsConfigXml_11_0_2) MarshalFromObject(s DdnsConfig) {
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
		var objs []ddnsConfigDdnsVendorConfigXml_11_0_2
		for _, elt := range s.DdnsVendorConfig {
			var obj ddnsConfigDdnsVendorConfigXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DdnsVendorConfig = &ddnsConfigDdnsVendorConfigContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ddnsConfigXml_11_0_2) UnmarshalToObject() (*DdnsConfig, error) {
	var ddnsIpVal []string
	if o.DdnsIp != nil {
		ddnsIpVal = util.MemToStr(o.DdnsIp)
	}
	var ddnsIpv6Val []string
	if o.DdnsIpv6 != nil {
		ddnsIpv6Val = util.MemToStr(o.DdnsIpv6)
	}
	var ddnsVendorConfigVal []DdnsConfigDdnsVendorConfig
	if o.DdnsVendorConfig != nil {
		for _, elt := range o.DdnsVendorConfig.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ddnsVendorConfigVal = append(ddnsVendorConfigVal, *obj)
		}
	}

	result := &DdnsConfig{
		DdnsCertProfile:    o.DdnsCertProfile,
		DdnsEnabled:        util.AsBool(o.DdnsEnabled, nil),
		DdnsHostname:       o.DdnsHostname,
		DdnsIp:             ddnsIpVal,
		DdnsIpv6:           ddnsIpv6Val,
		DdnsUpdateInterval: o.DdnsUpdateInterval,
		DdnsVendor:         o.DdnsVendor,
		DdnsVendorConfig:   ddnsVendorConfigVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ddnsConfigDdnsVendorConfigXml_11_0_2) MarshalFromObject(s DdnsConfigDdnsVendorConfig) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ddnsConfigDdnsVendorConfigXml_11_0_2) UnmarshalToObject() (*DdnsConfigDdnsVendorConfig, error) {

	result := &DdnsConfigDdnsVendorConfig{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *dhcpClientXml_11_0_2) MarshalFromObject(s DhcpClient) {
	o.CreateDefaultRoute = util.YesNo(s.CreateDefaultRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.SendHostname != nil {
		var obj dhcpClientSendHostnameXml_11_0_2
		obj.MarshalFromObject(*s.SendHostname)
		o.SendHostname = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o dhcpClientXml_11_0_2) UnmarshalToObject() (*DhcpClient, error) {
	var sendHostnameVal *DhcpClientSendHostname
	if o.SendHostname != nil {
		obj, err := o.SendHostname.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendHostnameVal = obj
	}

	result := &DhcpClient{
		CreateDefaultRoute: util.AsBool(o.CreateDefaultRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		SendHostname:       sendHostnameVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *dhcpClientSendHostnameXml_11_0_2) MarshalFromObject(s DhcpClientSendHostname) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Hostname = s.Hostname
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o dhcpClientSendHostnameXml_11_0_2) UnmarshalToObject() (*DhcpClientSendHostname, error) {

	result := &DhcpClientSendHostname{
		Enable:         util.AsBool(o.Enable, nil),
		Hostname:       o.Hostname,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipXml_11_0_2) MarshalFromObject(s Ip) {
	o.Name = s.Name
	o.SdwanGateway = s.SdwanGateway
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipXml_11_0_2) UnmarshalToObject() (*Ip, error) {

	result := &Ip{
		Name:           o.Name,
		SdwanGateway:   o.SdwanGateway,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6Xml_11_0_2) MarshalFromObject(s Ipv6) {
	if s.Address != nil {
		var objs []ipv6AddressXml_11_0_2
		for _, elt := range s.Address {
			var obj ipv6AddressXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &ipv6AddressContainerXml_11_0_2{Entries: objs}
	}
	if s.DhcpClient != nil {
		var obj ipv6DhcpClientXml_11_0_2
		obj.MarshalFromObject(*s.DhcpClient)
		o.DhcpClient = &obj
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Inherited != nil {
		var obj ipv6InheritedXml_11_0_2
		obj.MarshalFromObject(*s.Inherited)
		o.Inherited = &obj
	}
	o.InterfaceId = s.InterfaceId
	if s.NeighborDiscovery != nil {
		var obj ipv6NeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6Xml_11_0_2) UnmarshalToObject() (*Ipv6, error) {
	var addressVal []Ipv6Address
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}
	var dhcpClientVal *Ipv6DhcpClient
	if o.DhcpClient != nil {
		obj, err := o.DhcpClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpClientVal = obj
	}
	var inheritedVal *Ipv6Inherited
	if o.Inherited != nil {
		obj, err := o.Inherited.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritedVal = obj
	}
	var neighborDiscoveryVal *Ipv6NeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Ipv6{
		Address:           addressVal,
		DhcpClient:        dhcpClientVal,
		Enabled:           util.AsBool(o.Enabled, nil),
		Inherited:         inheritedVal,
		InterfaceId:       o.InterfaceId,
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressXml_11_0_2) MarshalFromObject(s Ipv6Address) {
	o.Name = s.Name
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	if s.Prefix != nil {
		var obj ipv6AddressPrefixXml_11_0_2
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	if s.Anycast != nil {
		var obj ipv6AddressAnycastXml_11_0_2
		obj.MarshalFromObject(*s.Anycast)
		o.Anycast = &obj
	}
	if s.Advertise != nil {
		var obj ipv6AddressAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressXml_11_0_2) UnmarshalToObject() (*Ipv6Address, error) {
	var prefixVal *Ipv6AddressPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}
	var anycastVal *Ipv6AddressAnycast
	if o.Anycast != nil {
		obj, err := o.Anycast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anycastVal = obj
	}
	var advertiseVal *Ipv6AddressAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6Address{
		Name:              o.Name,
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Prefix:            prefixVal,
		Anycast:           anycastVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressPrefixXml_11_0_2) MarshalFromObject(s Ipv6AddressPrefix) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressPrefixXml_11_0_2) UnmarshalToObject() (*Ipv6AddressPrefix, error) {

	result := &Ipv6AddressPrefix{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressAnycastXml_11_0_2) MarshalFromObject(s Ipv6AddressAnycast) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressAnycastXml_11_0_2) UnmarshalToObject() (*Ipv6AddressAnycast, error) {

	result := &Ipv6AddressAnycast{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressAdvertiseXml_11_0_2) MarshalFromObject(s Ipv6AddressAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressAdvertiseXml_11_0_2) UnmarshalToObject() (*Ipv6AddressAdvertise, error) {

	result := &Ipv6AddressAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientXml_11_0_2) MarshalFromObject(s Ipv6DhcpClient) {
	o.AcceptRaRoute = util.YesNo(s.AcceptRaRoute, nil)
	o.DefaultRouteMetric = s.DefaultRouteMetric
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj ipv6DhcpClientNeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Preference = s.Preference
	if s.PrefixDelegation != nil {
		var obj ipv6DhcpClientPrefixDelegationXml_11_0_2
		obj.MarshalFromObject(*s.PrefixDelegation)
		o.PrefixDelegation = &obj
	}
	if s.V6Options != nil {
		var obj ipv6DhcpClientV6OptionsXml_11_0_2
		obj.MarshalFromObject(*s.V6Options)
		o.V6Options = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClient, error) {
	var neighborDiscoveryVal *Ipv6DhcpClientNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}
	var prefixDelegationVal *Ipv6DhcpClientPrefixDelegation
	if o.PrefixDelegation != nil {
		obj, err := o.PrefixDelegation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixDelegationVal = obj
	}
	var v6OptionsVal *Ipv6DhcpClientV6Options
	if o.V6Options != nil {
		obj, err := o.V6Options.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v6OptionsVal = obj
	}

	result := &Ipv6DhcpClient{
		AcceptRaRoute:      util.AsBool(o.AcceptRaRoute, nil),
		DefaultRouteMetric: o.DefaultRouteMetric,
		Enable:             util.AsBool(o.Enable, nil),
		NeighborDiscovery:  neighborDiscoveryVal,
		Preference:         o.Preference,
		PrefixDelegation:   prefixDelegationVal,
		V6Options:          v6OptionsVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6DhcpClientNeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscovery, error) {
	var dnsServerVal *Ipv6DhcpClientNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Ipv6DhcpClientNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscovery{
		DadAttempts:      o.DadAttempts,
		DnsServer:        dnsServerVal,
		DnsSuffix:        dnsSuffixVal,
		EnableDad:        util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor: util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:         neighborVal,
		NsInterval:       o.NsInterval,
		ReachableTime:    o.ReachableTime,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServer, error) {
	var sourceVal *Ipv6DhcpClientNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServer{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
		for _, elt := range s.Server {
			var obj ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffix{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual{
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientNeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientNeighborDiscoveryNeighbor, error) {

	result := &Ipv6DhcpClientNeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientPrefixDelegation) {
	if s.Enable != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableXml_11_0_2
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegation, error) {
	var enableVal *Ipv6DhcpClientPrefixDelegationEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Ipv6DhcpClientPrefixDelegation{
		Enable:         enableVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnable) {
	if s.No != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnable, error) {
	var noVal *Ipv6DhcpClientPrefixDelegationEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Ipv6DhcpClientPrefixDelegationEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Ipv6DhcpClientPrefixDelegationEnable{
		No:             noVal,
		Yes:            yesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnableNo) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableNoXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnableNo, error) {

	result := &Ipv6DhcpClientPrefixDelegationEnableNo{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientPrefixDelegationEnableYes) {
	o.PfxPoolName = s.PfxPoolName
	o.PrefixLen = s.PrefixLen
	o.PrefixLenHint = util.YesNo(s.PrefixLenHint, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientPrefixDelegationEnableYesXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientPrefixDelegationEnableYes, error) {

	result := &Ipv6DhcpClientPrefixDelegationEnableYes{
		PfxPoolName:    o.PfxPoolName,
		PrefixLen:      o.PrefixLen,
		PrefixLenHint:  util.AsBool(o.PrefixLenHint, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientV6Options) {
	o.DuidType = s.DuidType
	if s.Enable != nil {
		var obj ipv6DhcpClientV6OptionsEnableXml_11_0_2
		obj.MarshalFromObject(*s.Enable)
		o.Enable = &obj
	}
	o.RapidCommit = util.YesNo(s.RapidCommit, nil)
	o.SupportSrvrReconfig = util.YesNo(s.SupportSrvrReconfig, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientV6Options, error) {
	var enableVal *Ipv6DhcpClientV6OptionsEnable
	if o.Enable != nil {
		obj, err := o.Enable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enableVal = obj
	}

	result := &Ipv6DhcpClientV6Options{
		DuidType:            o.DuidType,
		Enable:              enableVal,
		RapidCommit:         util.AsBool(o.RapidCommit, nil),
		SupportSrvrReconfig: util.AsBool(o.SupportSrvrReconfig, nil),
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnable) {
	if s.No != nil {
		var obj ipv6DhcpClientV6OptionsEnableNoXml_11_0_2
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj ipv6DhcpClientV6OptionsEnableYesXml_11_0_2
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnable, error) {
	var noVal *Ipv6DhcpClientV6OptionsEnableNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *Ipv6DhcpClientV6OptionsEnableYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &Ipv6DhcpClientV6OptionsEnable{
		No:             noVal,
		Yes:            yesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableNoXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnableNo) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableNoXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnableNo, error) {

	result := &Ipv6DhcpClientV6OptionsEnableNo{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6DhcpClientV6OptionsEnableYesXml_11_0_2) MarshalFromObject(s Ipv6DhcpClientV6OptionsEnableYes) {
	o.NonTempAddr = util.YesNo(s.NonTempAddr, nil)
	o.TempAddr = util.YesNo(s.TempAddr, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6DhcpClientV6OptionsEnableYesXml_11_0_2) UnmarshalToObject() (*Ipv6DhcpClientV6OptionsEnableYes, error) {

	result := &Ipv6DhcpClientV6OptionsEnableYes{
		NonTempAddr:    util.AsBool(o.NonTempAddr, nil),
		TempAddr:       util.AsBool(o.TempAddr, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedXml_11_0_2) MarshalFromObject(s Ipv6Inherited) {
	if s.AssignAddr != nil {
		var objs []ipv6InheritedAssignAddrXml_11_0_2
		for _, elt := range s.AssignAddr {
			var obj ipv6InheritedAssignAddrXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AssignAddr = &ipv6InheritedAssignAddrContainerXml_11_0_2{Entries: objs}
	}
	o.Enable = util.YesNo(s.Enable, nil)
	if s.NeighborDiscovery != nil {
		var obj ipv6InheritedNeighborDiscoveryXml_11_0_2
		obj.MarshalFromObject(*s.NeighborDiscovery)
		o.NeighborDiscovery = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedXml_11_0_2) UnmarshalToObject() (*Ipv6Inherited, error) {
	var assignAddrVal []Ipv6InheritedAssignAddr
	if o.AssignAddr != nil {
		for _, elt := range o.AssignAddr.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			assignAddrVal = append(assignAddrVal, *obj)
		}
	}
	var neighborDiscoveryVal *Ipv6InheritedNeighborDiscovery
	if o.NeighborDiscovery != nil {
		obj, err := o.NeighborDiscovery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		neighborDiscoveryVal = obj
	}

	result := &Ipv6Inherited{
		AssignAddr:        assignAddrVal,
		Enable:            util.AsBool(o.Enable, nil),
		NeighborDiscovery: neighborDiscoveryVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddr) {
	o.Name = s.Name
	if s.Type != nil {
		var obj ipv6InheritedAssignAddrTypeXml_11_0_2
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddr, error) {
	var typeVal *Ipv6InheritedAssignAddrType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Ipv6InheritedAssignAddr{
		Name:           o.Name,
		Type:           typeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrType) {
	if s.Gua != nil {
		var obj ipv6InheritedAssignAddrTypeGuaXml_11_0_2
		obj.MarshalFromObject(*s.Gua)
		o.Gua = &obj
	}
	if s.Ula != nil {
		var obj ipv6InheritedAssignAddrTypeUlaXml_11_0_2
		obj.MarshalFromObject(*s.Ula)
		o.Ula = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrType, error) {
	var guaVal *Ipv6InheritedAssignAddrTypeGua
	if o.Gua != nil {
		obj, err := o.Gua.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		guaVal = obj
	}
	var ulaVal *Ipv6InheritedAssignAddrTypeUla
	if o.Ula != nil {
		obj, err := o.Ula.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ulaVal = obj
	}

	result := &Ipv6InheritedAssignAddrType{
		Gua:            guaVal,
		Ula:            ulaVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGua) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.PrefixPool = s.PrefixPool
	if s.PoolType != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2
		obj.MarshalFromObject(*s.PoolType)
		o.PoolType = &obj
	}
	if s.Advertise != nil {
		var obj ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGua, error) {
	var poolTypeVal *Ipv6InheritedAssignAddrTypeGuaPoolType
	if o.PoolType != nil {
		obj, err := o.PoolType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		poolTypeVal = obj
	}
	var advertiseVal *Ipv6InheritedAssignAddrTypeGuaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeGua{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		PrefixPool:        o.PrefixPool,
		PoolType:          poolTypeVal,
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolType) {
	if s.Dynamic != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	if s.DynamicId != nil {
		var obj ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2
		obj.MarshalFromObject(*s.DynamicId)
		o.DynamicId = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolType, error) {
	var dynamicVal *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}
	var dynamicIdVal *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId
	if o.DynamicId != nil {
		obj, err := o.DynamicId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIdVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeGuaPoolType{
		Dynamic:        dynamicVal,
		DynamicId:      dynamicIdVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) {
	o.Identifier = s.Identifier
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicIdXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId{
		Identifier:     o.Identifier,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeGuaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeGuaAdvertiseXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeGuaAdvertise, error) {

	result := &Ipv6InheritedAssignAddrTypeGuaAdvertise{
		Enable:         util.AsBool(o.Enable, nil),
		OnlinkFlag:     util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag: util.AsBool(o.AutoConfigFlag, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeUlaXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeUla) {
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	o.Address = s.Address
	o.Prefix = util.YesNo(s.Prefix, nil)
	o.Anycast = util.YesNo(s.Anycast, nil)
	if s.Advertise != nil {
		var obj ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2
		obj.MarshalFromObject(*s.Advertise)
		o.Advertise = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeUlaXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeUla, error) {
	var advertiseVal *Ipv6InheritedAssignAddrTypeUlaAdvertise
	if o.Advertise != nil {
		obj, err := o.Advertise.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		advertiseVal = obj
	}

	result := &Ipv6InheritedAssignAddrTypeUla{
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Address:           o.Address,
		Prefix:            util.AsBool(o.Prefix, nil),
		Anycast:           util.AsBool(o.Anycast, nil),
		Advertise:         advertiseVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2) MarshalFromObject(s Ipv6InheritedAssignAddrTypeUlaAdvertise) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.ValidLifetime = s.ValidLifetime
	o.PreferredLifetime = s.PreferredLifetime
	o.OnlinkFlag = util.YesNo(s.OnlinkFlag, nil)
	o.AutoConfigFlag = util.YesNo(s.AutoConfigFlag, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedAssignAddrTypeUlaAdvertiseXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedAssignAddrTypeUlaAdvertise, error) {

	result := &Ipv6InheritedAssignAddrTypeUlaAdvertise{
		Enable:            util.AsBool(o.Enable, nil),
		ValidLifetime:     o.ValidLifetime,
		PreferredLifetime: o.PreferredLifetime,
		OnlinkFlag:        util.AsBool(o.OnlinkFlag, nil),
		AutoConfigFlag:    util.AsBool(o.AutoConfigFlag, nil),
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	if s.DnsServer != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2
		obj.MarshalFromObject(*s.DnsServer)
		o.DnsServer = &obj
	}
	if s.DnsSuffix != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2
		obj.MarshalFromObject(*s.DnsSuffix)
		o.DnsSuffix = &obj
	}
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6InheritedNeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscovery, error) {
	var dnsServerVal *Ipv6InheritedNeighborDiscoveryDnsServer
	if o.DnsServer != nil {
		obj, err := o.DnsServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsServerVal = obj
	}
	var dnsSuffixVal *Ipv6InheritedNeighborDiscoveryDnsSuffix
	if o.DnsSuffix != nil {
		obj, err := o.DnsSuffix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSuffixVal = obj
	}
	var neighborVal []Ipv6InheritedNeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Ipv6InheritedNeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Ipv6InheritedNeighborDiscovery{
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
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServer) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServer, error) {
	var sourceVal *Ipv6InheritedNeighborDiscoveryDnsServerSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServer{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSource, error) {
	var dhcpv6Val *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6{
		PrefixPool:     o.PrefixPool,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) {
	if s.Server != nil {
		var objs []ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
		for _, elt := range s.Server {
			var obj ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceManualXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceManual, error) {
	var serverVal []Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceManual{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsServerSourceManualServerXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffix) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Source != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2
		obj.MarshalFromObject(*s.Source)
		o.Source = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffix, error) {
	var sourceVal *Ipv6InheritedNeighborDiscoveryDnsSuffixSource
	if o.Source != nil {
		obj, err := o.Source.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffix{
		Enable:         util.AsBool(o.Enable, nil),
		Source:         sourceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSource) {
	if s.Dhcpv6 != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2
		obj.MarshalFromObject(*s.Dhcpv6)
		o.Dhcpv6 = &obj
	}
	if s.Manual != nil {
		var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2
		obj.MarshalFromObject(*s.Manual)
		o.Manual = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSource, error) {
	var dhcpv6Val *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6
	if o.Dhcpv6 != nil {
		obj, err := o.Dhcpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dhcpv6Val = obj
	}
	var manualVal *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual
	if o.Manual != nil {
		obj, err := o.Manual.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualVal = obj
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSource{
		Dhcpv6:         dhcpv6Val,
		Manual:         manualVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) {
	o.PrefixPool = s.PrefixPool
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6Xml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6{
		PrefixPool:     o.PrefixPool,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) {
	if s.Suffix != nil {
		var objs []ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual, error) {
	var suffixVal []Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual{
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffixXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix, error) {

	result := &Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryNeighbor, error) {

	result := &Ipv6InheritedNeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2) MarshalFromObject(s Ipv6InheritedNeighborDiscoveryRouterAdvertisement) {
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
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6InheritedNeighborDiscoveryRouterAdvertisementXml_11_0_2) UnmarshalToObject() (*Ipv6InheritedNeighborDiscoveryRouterAdvertisement, error) {

	result := &Ipv6InheritedNeighborDiscoveryRouterAdvertisement{
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
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscovery) {
	o.DadAttempts = s.DadAttempts
	o.EnableDad = util.YesNo(s.EnableDad, nil)
	o.EnableNdpMonitor = util.YesNo(s.EnableNdpMonitor, nil)
	if s.Neighbor != nil {
		var objs []ipv6NeighborDiscoveryNeighborXml_11_0_2
		for _, elt := range s.Neighbor {
			var obj ipv6NeighborDiscoveryNeighborXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Neighbor = &ipv6NeighborDiscoveryNeighborContainerXml_11_0_2{Entries: objs}
	}
	o.NsInterval = s.NsInterval
	o.ReachableTime = s.ReachableTime
	if s.RouterAdvertisement != nil {
		var obj ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2
		obj.MarshalFromObject(*s.RouterAdvertisement)
		o.RouterAdvertisement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscovery, error) {
	var neighborVal []Ipv6NeighborDiscoveryNeighbor
	if o.Neighbor != nil {
		for _, elt := range o.Neighbor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			neighborVal = append(neighborVal, *obj)
		}
	}
	var routerAdvertisementVal *Ipv6NeighborDiscoveryRouterAdvertisement
	if o.RouterAdvertisement != nil {
		obj, err := o.RouterAdvertisement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routerAdvertisementVal = obj
	}

	result := &Ipv6NeighborDiscovery{
		DadAttempts:         o.DadAttempts,
		EnableDad:           util.AsBool(o.EnableDad, nil),
		EnableNdpMonitor:    util.AsBool(o.EnableNdpMonitor, nil),
		Neighbor:            neighborVal,
		NsInterval:          o.NsInterval,
		ReachableTime:       o.ReachableTime,
		RouterAdvertisement: routerAdvertisementVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryNeighborXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscoveryNeighbor) {
	o.Name = s.Name
	o.HwAddress = s.HwAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryNeighborXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscoveryNeighbor, error) {

	result := &Ipv6NeighborDiscoveryNeighbor{
		Name:           o.Name,
		HwAddress:      o.HwAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisement) {
	if s.DnsSupport != nil {
		var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2
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
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisement, error) {
	var dnsSupportVal *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport
	if o.DnsSupport != nil {
		obj, err := o.DnsSupport.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSupportVal = obj
	}

	result := &Ipv6NeighborDiscoveryRouterAdvertisement{
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
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Server != nil {
		var objs []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2
		for _, elt := range s.Server {
			var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerContainerXml_11_0_2{Entries: objs}
	}
	if s.Suffix != nil {
		var objs []ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2
		for _, elt := range s.Suffix {
			var obj ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Suffix = &ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport, error) {
	var serverVal []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}
	var suffixVal []Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix
	if o.Suffix != nil {
		for _, elt := range o.Suffix.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			suffixVal = append(suffixVal, *obj)
		}
	}

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport{
		Enable:         util.AsBool(o.Enable, nil),
		Server:         serverVal,
		Suffix:         suffixVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServerXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer, error) {

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2) MarshalFromObject(s Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) {
	o.Name = s.Name
	o.Lifetime = s.Lifetime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffixXml_11_0_2) UnmarshalToObject() (*Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix, error) {

	result := &Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix{
		Name:           o.Name,
		Lifetime:       o.Lifetime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ndpProxyXml_11_0_2) MarshalFromObject(s NdpProxy) {
	if s.Address != nil {
		var objs []ndpProxyAddressXml_11_0_2
		for _, elt := range s.Address {
			var obj ndpProxyAddressXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &ndpProxyAddressContainerXml_11_0_2{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ndpProxyXml_11_0_2) UnmarshalToObject() (*NdpProxy, error) {
	var addressVal []NdpProxyAddress
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &NdpProxy{
		Address:        addressVal,
		Enabled:        util.AsBool(o.Enabled, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ndpProxyAddressXml_11_0_2) MarshalFromObject(s NdpProxyAddress) {
	o.Name = s.Name
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ndpProxyAddressXml_11_0_2) UnmarshalToObject() (*NdpProxyAddress, error) {

	result := &NdpProxyAddress{
		Name:           o.Name,
		Negate:         util.AsBool(o.Negate, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsXml_11_0_2) MarshalFromObject(s SdwanLinkSettings) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	if s.UpstreamNat != nil {
		var obj sdwanLinkSettingsUpstreamNatXml_11_0_2
		obj.MarshalFromObject(*s.UpstreamNat)
		o.UpstreamNat = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsXml_11_0_2) UnmarshalToObject() (*SdwanLinkSettings, error) {
	var upstreamNatVal *SdwanLinkSettingsUpstreamNat
	if o.UpstreamNat != nil {
		obj, err := o.UpstreamNat.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		upstreamNatVal = obj
	}

	result := &SdwanLinkSettings{
		Enable:                util.AsBool(o.Enable, nil),
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		UpstreamNat:           upstreamNatVal,
		Misc:                  o.Misc,
		MiscAttributes:        o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatXml_11_0_2) MarshalFromObject(s SdwanLinkSettingsUpstreamNat) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Ddns != nil {
		var obj sdwanLinkSettingsUpstreamNatDdnsXml_11_0_2
		obj.MarshalFromObject(*s.Ddns)
		o.Ddns = &obj
	}
	if s.StaticIp != nil {
		var obj sdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2
		obj.MarshalFromObject(*s.StaticIp)
		o.StaticIp = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatXml_11_0_2) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNat, error) {
	var ddnsVal *SdwanLinkSettingsUpstreamNatDdns
	if o.Ddns != nil {
		obj, err := o.Ddns.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ddnsVal = obj
	}
	var staticIpVal *SdwanLinkSettingsUpstreamNatStaticIp
	if o.StaticIp != nil {
		obj, err := o.StaticIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticIpVal = obj
	}

	result := &SdwanLinkSettingsUpstreamNat{
		Enable:         util.AsBool(o.Enable, nil),
		Ddns:           ddnsVal,
		StaticIp:       staticIpVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatDdnsXml_11_0_2) MarshalFromObject(s SdwanLinkSettingsUpstreamNatDdns) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatDdnsXml_11_0_2) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNatDdns, error) {

	result := &SdwanLinkSettingsUpstreamNatDdns{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *sdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2) MarshalFromObject(s SdwanLinkSettingsUpstreamNatStaticIp) {
	o.Fqdn = s.Fqdn
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sdwanLinkSettingsUpstreamNatStaticIpXml_11_0_2) UnmarshalToObject() (*SdwanLinkSettingsUpstreamNatStaticIp, error) {

	result := &SdwanLinkSettingsUpstreamNatStaticIp{
		Fqdn:           o.Fqdn,
		IpAddress:      o.IpAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
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
	if v == "decrypt_forward" || v == "DecryptForward" {
		return e.DecryptForward, nil
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
	if v == "sdwan_link_settings" || v == "SdwanLinkSettings" {
		return e.SdwanLinkSettings, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
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
	if !util.StringsMatch(o.Comment, other.Comment) {
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
	if !util.Ints64Match(o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *AdjustTcpMss) matches(other *AdjustTcpMss) bool {
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

func (o *Arp) matches(other *Arp) bool {
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

func (o *Bonjour) matches(other *Bonjour) bool {
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

func (o *DdnsConfig) matches(other *DdnsConfig) bool {
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

func (o *DdnsConfigDdnsVendorConfig) matches(other *DdnsConfigDdnsVendorConfig) bool {
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

func (o *DhcpClient) matches(other *DhcpClient) bool {
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

func (o *DhcpClientSendHostname) matches(other *DhcpClientSendHostname) bool {
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

func (o *Ip) matches(other *Ip) bool {
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

func (o *Ipv6) matches(other *Ipv6) bool {
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

func (o *Ipv6Address) matches(other *Ipv6Address) bool {
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

func (o *Ipv6AddressPrefix) matches(other *Ipv6AddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6AddressAnycast) matches(other *Ipv6AddressAnycast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6AddressAdvertise) matches(other *Ipv6AddressAdvertise) bool {
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

func (o *Ipv6DhcpClient) matches(other *Ipv6DhcpClient) bool {
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

func (o *Ipv6DhcpClientNeighborDiscovery) matches(other *Ipv6DhcpClientNeighborDiscovery) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsServer) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsServer) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsServerSource) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsServerSource) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManual) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsServerSourceManualServer) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsSuffix) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsSuffix) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSource) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManual) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) matches(other *Ipv6DhcpClientNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
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

func (o *Ipv6DhcpClientNeighborDiscoveryNeighbor) matches(other *Ipv6DhcpClientNeighborDiscoveryNeighbor) bool {
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

func (o *Ipv6DhcpClientPrefixDelegation) matches(other *Ipv6DhcpClientPrefixDelegation) bool {
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

func (o *Ipv6DhcpClientPrefixDelegationEnable) matches(other *Ipv6DhcpClientPrefixDelegationEnable) bool {
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

func (o *Ipv6DhcpClientPrefixDelegationEnableNo) matches(other *Ipv6DhcpClientPrefixDelegationEnableNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6DhcpClientPrefixDelegationEnableYes) matches(other *Ipv6DhcpClientPrefixDelegationEnableYes) bool {
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

func (o *Ipv6DhcpClientV6Options) matches(other *Ipv6DhcpClientV6Options) bool {
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

func (o *Ipv6DhcpClientV6OptionsEnable) matches(other *Ipv6DhcpClientV6OptionsEnable) bool {
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

func (o *Ipv6DhcpClientV6OptionsEnableNo) matches(other *Ipv6DhcpClientV6OptionsEnableNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6DhcpClientV6OptionsEnableYes) matches(other *Ipv6DhcpClientV6OptionsEnableYes) bool {
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

func (o *Ipv6Inherited) matches(other *Ipv6Inherited) bool {
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

func (o *Ipv6InheritedAssignAddr) matches(other *Ipv6InheritedAssignAddr) bool {
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

func (o *Ipv6InheritedAssignAddrType) matches(other *Ipv6InheritedAssignAddrType) bool {
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

func (o *Ipv6InheritedAssignAddrTypeGua) matches(other *Ipv6InheritedAssignAddrTypeGua) bool {
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

func (o *Ipv6InheritedAssignAddrTypeGuaPoolType) matches(other *Ipv6InheritedAssignAddrTypeGuaPoolType) bool {
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

func (o *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) matches(other *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) matches(other *Ipv6InheritedAssignAddrTypeGuaPoolTypeDynamicId) bool {
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

func (o *Ipv6InheritedAssignAddrTypeGuaAdvertise) matches(other *Ipv6InheritedAssignAddrTypeGuaAdvertise) bool {
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

func (o *Ipv6InheritedAssignAddrTypeUla) matches(other *Ipv6InheritedAssignAddrTypeUla) bool {
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

func (o *Ipv6InheritedAssignAddrTypeUlaAdvertise) matches(other *Ipv6InheritedAssignAddrTypeUlaAdvertise) bool {
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

func (o *Ipv6InheritedNeighborDiscovery) matches(other *Ipv6InheritedNeighborDiscovery) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsServer) matches(other *Ipv6InheritedNeighborDiscoveryDnsServer) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsServerSource) matches(other *Ipv6InheritedNeighborDiscoveryDnsServerSource) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) matches(other *Ipv6InheritedNeighborDiscoveryDnsServerSourceDhcpv6) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) matches(other *Ipv6InheritedNeighborDiscoveryDnsServerSourceManual) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) matches(other *Ipv6InheritedNeighborDiscoveryDnsServerSourceManualServer) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsSuffix) matches(other *Ipv6InheritedNeighborDiscoveryDnsSuffix) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsSuffixSource) matches(other *Ipv6InheritedNeighborDiscoveryDnsSuffixSource) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) matches(other *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceDhcpv6) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) matches(other *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManual) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) matches(other *Ipv6InheritedNeighborDiscoveryDnsSuffixSourceManualSuffix) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryNeighbor) matches(other *Ipv6InheritedNeighborDiscoveryNeighbor) bool {
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

func (o *Ipv6InheritedNeighborDiscoveryRouterAdvertisement) matches(other *Ipv6InheritedNeighborDiscoveryRouterAdvertisement) bool {
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

func (o *Ipv6NeighborDiscovery) matches(other *Ipv6NeighborDiscovery) bool {
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

func (o *Ipv6NeighborDiscoveryNeighbor) matches(other *Ipv6NeighborDiscoveryNeighbor) bool {
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

func (o *Ipv6NeighborDiscoveryRouterAdvertisement) matches(other *Ipv6NeighborDiscoveryRouterAdvertisement) bool {
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

func (o *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) matches(other *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupport) bool {
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

func (o *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) matches(other *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportServer) bool {
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

func (o *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) matches(other *Ipv6NeighborDiscoveryRouterAdvertisementDnsSupportSuffix) bool {
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

func (o *NdpProxy) matches(other *NdpProxy) bool {
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

func (o *NdpProxyAddress) matches(other *NdpProxyAddress) bool {
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

func (o *SdwanLinkSettings) matches(other *SdwanLinkSettings) bool {
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

func (o *SdwanLinkSettingsUpstreamNat) matches(other *SdwanLinkSettingsUpstreamNat) bool {
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

func (o *SdwanLinkSettingsUpstreamNatDdns) matches(other *SdwanLinkSettingsUpstreamNatDdns) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *SdwanLinkSettingsUpstreamNatStaticIp) matches(other *SdwanLinkSettingsUpstreamNatStaticIp) bool {
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
