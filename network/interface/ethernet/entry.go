package ethernet

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
	Suffix = []string{"network", "interface", "ethernet"}
)

type Entry struct {
	Name       string
	Comment    *string
	LinkDuplex *string
	LinkSpeed  *string
	LinkState  *string
	Poe        *Poe
	Ha         *Ha
	Layer3     *Layer3
	Tap        *Tap

	Misc map[string][]generic.Xml
}

type Ha struct {
}
type Layer3 struct {
	AdjustTcpMss               *Layer3AdjustTcpMss
	Arp                        []Layer3Arp
	Bonjour                    *Layer3Bonjour
	DhcpClient                 *Layer3DhcpClient
	InterfaceManagementProfile *string
	Ips                        []Layer3Ips
	Ipv6                       *Layer3Ipv6
	Lldp                       *Layer3Lldp
	Mtu                        *int64
	NdpProxy                   *bool
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
	Enable *bool
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
type Layer3Ips struct {
	Name         string
	SdwanGateway *string
}
type Layer3Ipv6 struct {
	Addresses         []Layer3Ipv6Addresses
	DnsServer         *Layer3Ipv6DnsServer
	Enabled           *bool
	InterfaceId       *string
	NeighborDiscovery *Layer3Ipv6NeighborDiscovery
}
type Layer3Ipv6Addresses struct {
	Advertise         *Layer3Ipv6AddressesAdvertise
	Anycast           *string
	EnableOnInterface *bool
	Name              string
	Prefix            *string
}
type Layer3Ipv6AddressesAdvertise struct {
	AutoConfigFlag    *bool
	Enable            *bool
	OnlinkFlag        *bool
	PreferredLifetime *string
	ValidLifetime     *string
}
type Layer3Ipv6DnsServer struct {
	DnsSupport *Layer3Ipv6DnsServerDnsSupport
	Enable     *bool
	Source     *Layer3Ipv6DnsServerSource
}
type Layer3Ipv6DnsServerDnsSupport struct {
	Enable *bool
	Server []Layer3Ipv6DnsServerDnsSupportServer
	Suffix []Layer3Ipv6DnsServerDnsSupportSuffix
}
type Layer3Ipv6DnsServerDnsSupportServer struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6DnsServerDnsSupportSuffix struct {
	Lifetime *int64
	Name     string
}
type Layer3Ipv6DnsServerSource struct {
	Dhcpv6 *Layer3Ipv6DnsServerSourceDhcpv6
	Manual *Layer3Ipv6DnsServerSourceManual
}
type Layer3Ipv6DnsServerSourceDhcpv6 struct {
	PrefixPool *string
}
type Layer3Ipv6DnsServerSourceManual struct {
	Suffix []Layer3Ipv6DnsServerSourceManualSuffix
}
type Layer3Ipv6DnsServerSourceManualSuffix struct {
	Lifetime *int64
	Name     string
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
type Layer3Lldp struct {
	Enable  *bool
	Profile *string
}
type Layer3SdwanLinkSettings struct {
	Enable                *bool
	SdwanInterfaceProfile *string
	UpstreamNat           *Layer3SdwanLinkSettingsUpstreamNat
}
type Layer3SdwanLinkSettingsUpstreamNat struct {
	Enable   *bool
	StaticIp *string
}
type Poe struct {
	Enabled       *bool
	ReservedPower *int64
}
type Tap struct {
	NetflowProfile *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName    xml.Name `xml:"entry"`
	Name       string   `xml:"name,attr"`
	Comment    *string  `xml:"comment,omitempty"`
	LinkDuplex *string  `xml:"link-duplex,omitempty"`
	LinkSpeed  *string  `xml:"link-speed,omitempty"`
	LinkState  *string  `xml:"link-state,omitempty"`
	Poe        *PoeXml
	Ha         *HaXml     `xml:"ha,omitempty"`
	Layer3     *Layer3Xml `xml:"layer3,omitempty"`
	Tap        *TapXml    `xml:"tap,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type HaXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Layer3Xml struct {
	AdjustTcpMss               *Layer3AdjustTcpMssXml      `xml:"adjust-tcp-mss,omitempty"`
	Arp                        []Layer3ArpXml              `xml:"arp>entry,omitempty"`
	Bonjour                    *Layer3BonjourXml           `xml:"bonjour,omitempty"`
	DhcpClient                 *Layer3DhcpClientXml        `xml:"dhcp-client,omitempty"`
	InterfaceManagementProfile *string                     `xml:"interface-management-profile,omitempty"`
	Ips                        []Layer3IpsXml              `xml:"ip>entry,omitempty"`
	Ipv6                       *Layer3Ipv6Xml              `xml:"ipv6,omitempty"`
	Lldp                       *Layer3LldpXml              `xml:"lldp,omitempty"`
	Mtu                        *int64                      `xml:"mtu,omitempty"`
	NdpProxy                   *string                     `xml:"ndp-proxy>enabled,omitempty"`
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
	Enable *string

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
type Layer3IpsXml struct {
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	SdwanGateway *string  `xml:"sdwan-gateway,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6Xml struct {
	Addresses         []Layer3Ipv6AddressesXml        `xml:"address>entry,omitempty"`
	DnsServer         *Layer3Ipv6DnsServerXml         `xml:"dns-server,omitempty"`
	Enabled           *string                         `xml:"enabled,omitempty"`
	InterfaceId       *string                         `xml:"interface-id,omitempty"`
	NeighborDiscovery *Layer3Ipv6NeighborDiscoveryXml `xml:"neighbor-discovery,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressesXml struct {
	Advertise         *Layer3Ipv6AddressesAdvertiseXml `xml:"advertise,omitempty"`
	Anycast           *string                          `xml:"anycast,omitempty"`
	EnableOnInterface *string                          `xml:"enable-on-interface,omitempty"`
	XMLName           xml.Name                         `xml:"entry"`
	Name              string                           `xml:"name,attr"`
	Prefix            *string                          `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6AddressesAdvertiseXml struct {
	AutoConfigFlag    *string `xml:"auto-config-flag,omitempty"`
	Enable            *string `xml:"enable,omitempty"`
	OnlinkFlag        *string `xml:"onlink-flag,omitempty"`
	PreferredLifetime *string `xml:"preferred-lifetime,omitempty"`
	ValidLifetime     *string `xml:"valid-lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerXml struct {
	DnsSupport *Layer3Ipv6DnsServerDnsSupportXml `xml:"dns-support,omitempty"`
	Enable     *string                           `xml:"enable,omitempty"`
	Source     *Layer3Ipv6DnsServerSourceXml     `xml:"source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerDnsSupportXml struct {
	Enable *string
	Server []Layer3Ipv6DnsServerDnsSupportServerXml `xml:"server>entry,omitempty"`
	Suffix []Layer3Ipv6DnsServerDnsSupportSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerDnsSupportServerXml struct {
	Lifetime *int64
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerDnsSupportSuffixXml struct {
	Lifetime *int64
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerSourceXml struct {
	Dhcpv6 *Layer3Ipv6DnsServerSourceDhcpv6Xml `xml:"dhcpv6,omitempty"`
	Manual *Layer3Ipv6DnsServerSourceManualXml `xml:"manual,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerSourceDhcpv6Xml struct {
	PrefixPool *string `xml:"prefix-pool,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerSourceManualXml struct {
	Suffix []Layer3Ipv6DnsServerSourceManualSuffixXml `xml:"suffix>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3Ipv6DnsServerSourceManualSuffixXml struct {
	Lifetime *int64   `xml:"lifetime,omitempty"`
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`

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
type Layer3LldpXml struct {
	Enable  *string `xml:"enable,omitempty"`
	Profile *string `xml:"profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsXml struct {
	Enable                *string                                `xml:"enable,omitempty"`
	SdwanInterfaceProfile *string                                `xml:"sdwan-interface-profile,omitempty"`
	UpstreamNat           *Layer3SdwanLinkSettingsUpstreamNatXml `xml:"upstream-nat,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Layer3SdwanLinkSettingsUpstreamNatXml struct {
	Enable   *string `xml:"enable,omitempty"`
	StaticIp *string `xml:"static-ip>ip-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type PoeXml struct {
	Enabled       *string `xml:"poe-enabled,omitempty"`
	ReservedPower *int64  `xml:"poe-rsvd-pwr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TapXml struct {
	NetflowProfile *string `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "link_duplex" || v == "LinkDuplex" {
		return e.LinkDuplex, nil
	}
	if v == "link_speed" || v == "LinkSpeed" {
		return e.LinkSpeed, nil
	}
	if v == "link_state" || v == "LinkState" {
		return e.LinkState, nil
	}
	if v == "poe" || v == "Poe" {
		return e.Poe, nil
	}
	if v == "ha" || v == "Ha" {
		return e.Ha, nil
	}
	if v == "layer3" || v == "Layer3" {
		return e.Layer3, nil
	}
	if v == "tap" || v == "Tap" {
		return e.Tap, nil
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
	entry.LinkDuplex = o.LinkDuplex
	entry.LinkSpeed = o.LinkSpeed
	entry.LinkState = o.LinkState
	var nestedPoe *PoeXml
	if o.Poe != nil {
		nestedPoe = &PoeXml{}
		if _, ok := o.Misc["Poe"]; ok {
			nestedPoe.Misc = o.Misc["Poe"]
		}
		if o.Poe.ReservedPower != nil {
			nestedPoe.ReservedPower = o.Poe.ReservedPower
		}
		if o.Poe.Enabled != nil {
			nestedPoe.Enabled = util.YesNo(o.Poe.Enabled, nil)
		}
	}
	entry.Poe = nestedPoe

	var nestedHa *HaXml
	if o.Ha != nil {
		nestedHa = &HaXml{}
		if _, ok := o.Misc["Ha"]; ok {
			nestedHa.Misc = o.Misc["Ha"]
		}
	}
	entry.Ha = nestedHa

	var nestedLayer3 *Layer3Xml
	if o.Layer3 != nil {
		nestedLayer3 = &Layer3Xml{}
		if _, ok := o.Misc["Layer3"]; ok {
			nestedLayer3.Misc = o.Misc["Layer3"]
		}
		if o.Layer3.Bonjour != nil {
			nestedLayer3.Bonjour = &Layer3BonjourXml{}
			if _, ok := o.Misc["Layer3Bonjour"]; ok {
				nestedLayer3.Bonjour.Misc = o.Misc["Layer3Bonjour"]
			}
			if o.Layer3.Bonjour.Enable != nil {
				nestedLayer3.Bonjour.Enable = util.YesNo(o.Layer3.Bonjour.Enable, nil)
			}
		}
		if o.Layer3.UntaggedSubInterface != nil {
			nestedLayer3.UntaggedSubInterface = util.YesNo(o.Layer3.UntaggedSubInterface, nil)
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
		if o.Layer3.NdpProxy != nil {
			nestedLayer3.NdpProxy = util.YesNo(o.Layer3.NdpProxy, nil)
		}
		if o.Layer3.Lldp != nil {
			nestedLayer3.Lldp = &Layer3LldpXml{}
			if _, ok := o.Misc["Layer3Lldp"]; ok {
				nestedLayer3.Lldp.Misc = o.Misc["Layer3Lldp"]
			}
			if o.Layer3.Lldp.Enable != nil {
				nestedLayer3.Lldp.Enable = util.YesNo(o.Layer3.Lldp.Enable, nil)
			}
			if o.Layer3.Lldp.Profile != nil {
				nestedLayer3.Lldp.Profile = o.Layer3.Lldp.Profile
			}
		}
		if o.Layer3.DhcpClient != nil {
			nestedLayer3.DhcpClient = &Layer3DhcpClientXml{}
			if _, ok := o.Misc["Layer3DhcpClient"]; ok {
				nestedLayer3.DhcpClient.Misc = o.Misc["Layer3DhcpClient"]
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
			if o.Layer3.DhcpClient.Enable != nil {
				nestedLayer3.DhcpClient.Enable = util.YesNo(o.Layer3.DhcpClient.Enable, nil)
			}
			if o.Layer3.DhcpClient.CreateDefaultRoute != nil {
				nestedLayer3.DhcpClient.CreateDefaultRoute = util.YesNo(o.Layer3.DhcpClient.CreateDefaultRoute, nil)
			}
			if o.Layer3.DhcpClient.DefaultRouteMetric != nil {
				nestedLayer3.DhcpClient.DefaultRouteMetric = o.Layer3.DhcpClient.DefaultRouteMetric
			}
		}
		if o.Layer3.NetflowProfile != nil {
			nestedLayer3.NetflowProfile = o.Layer3.NetflowProfile
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
				if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp != nil {
					nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp
				}
			}
		}
		if o.Layer3.AdjustTcpMss != nil {
			nestedLayer3.AdjustTcpMss = &Layer3AdjustTcpMssXml{}
			if _, ok := o.Misc["Layer3AdjustTcpMss"]; ok {
				nestedLayer3.AdjustTcpMss.Misc = o.Misc["Layer3AdjustTcpMss"]
			}
			if o.Layer3.AdjustTcpMss.Ipv6MssAdjustment != nil {
				nestedLayer3.AdjustTcpMss.Ipv6MssAdjustment = o.Layer3.AdjustTcpMss.Ipv6MssAdjustment
			}
			if o.Layer3.AdjustTcpMss.Enable != nil {
				nestedLayer3.AdjustTcpMss.Enable = util.YesNo(o.Layer3.AdjustTcpMss.Enable, nil)
			}
			if o.Layer3.AdjustTcpMss.Ipv4MssAdjustment != nil {
				nestedLayer3.AdjustTcpMss.Ipv4MssAdjustment = o.Layer3.AdjustTcpMss.Ipv4MssAdjustment
			}
		}
		if o.Layer3.Mtu != nil {
			nestedLayer3.Mtu = o.Layer3.Mtu
		}
		if o.Layer3.Ips != nil {
			nestedLayer3.Ips = []Layer3IpsXml{}
			for _, oLayer3Ips := range o.Layer3.Ips {
				nestedLayer3Ips := Layer3IpsXml{}
				if _, ok := o.Misc["Layer3Ips"]; ok {
					nestedLayer3Ips.Misc = o.Misc["Layer3Ips"]
				}
				if oLayer3Ips.SdwanGateway != nil {
					nestedLayer3Ips.SdwanGateway = oLayer3Ips.SdwanGateway
				}
				if oLayer3Ips.Name != "" {
					nestedLayer3Ips.Name = oLayer3Ips.Name
				}
				nestedLayer3.Ips = append(nestedLayer3.Ips, nestedLayer3Ips)
			}
		}
		if o.Layer3.Ipv6 != nil {
			nestedLayer3.Ipv6 = &Layer3Ipv6Xml{}
			if _, ok := o.Misc["Layer3Ipv6"]; ok {
				nestedLayer3.Ipv6.Misc = o.Misc["Layer3Ipv6"]
			}
			if o.Layer3.Ipv6.Enabled != nil {
				nestedLayer3.Ipv6.Enabled = util.YesNo(o.Layer3.Ipv6.Enabled, nil)
			}
			if o.Layer3.Ipv6.InterfaceId != nil {
				nestedLayer3.Ipv6.InterfaceId = o.Layer3.Ipv6.InterfaceId
			}
			if o.Layer3.Ipv6.Addresses != nil {
				nestedLayer3.Ipv6.Addresses = []Layer3Ipv6AddressesXml{}
				for _, oLayer3Ipv6Addresses := range o.Layer3.Ipv6.Addresses {
					nestedLayer3Ipv6Addresses := Layer3Ipv6AddressesXml{}
					if _, ok := o.Misc["Layer3Ipv6Addresses"]; ok {
						nestedLayer3Ipv6Addresses.Misc = o.Misc["Layer3Ipv6Addresses"]
					}
					if oLayer3Ipv6Addresses.Name != "" {
						nestedLayer3Ipv6Addresses.Name = oLayer3Ipv6Addresses.Name
					}
					if oLayer3Ipv6Addresses.EnableOnInterface != nil {
						nestedLayer3Ipv6Addresses.EnableOnInterface = util.YesNo(oLayer3Ipv6Addresses.EnableOnInterface, nil)
					}
					if oLayer3Ipv6Addresses.Prefix != nil {
						nestedLayer3Ipv6Addresses.Prefix = oLayer3Ipv6Addresses.Prefix
					}
					if oLayer3Ipv6Addresses.Anycast != nil {
						nestedLayer3Ipv6Addresses.Anycast = oLayer3Ipv6Addresses.Anycast
					}
					if oLayer3Ipv6Addresses.Advertise != nil {
						nestedLayer3Ipv6Addresses.Advertise = &Layer3Ipv6AddressesAdvertiseXml{}
						if _, ok := o.Misc["Layer3Ipv6AddressesAdvertise"]; ok {
							nestedLayer3Ipv6Addresses.Advertise.Misc = o.Misc["Layer3Ipv6AddressesAdvertise"]
						}
						if oLayer3Ipv6Addresses.Advertise.OnlinkFlag != nil {
							nestedLayer3Ipv6Addresses.Advertise.OnlinkFlag = util.YesNo(oLayer3Ipv6Addresses.Advertise.OnlinkFlag, nil)
						}
						if oLayer3Ipv6Addresses.Advertise.AutoConfigFlag != nil {
							nestedLayer3Ipv6Addresses.Advertise.AutoConfigFlag = util.YesNo(oLayer3Ipv6Addresses.Advertise.AutoConfigFlag, nil)
						}
						if oLayer3Ipv6Addresses.Advertise.Enable != nil {
							nestedLayer3Ipv6Addresses.Advertise.Enable = util.YesNo(oLayer3Ipv6Addresses.Advertise.Enable, nil)
						}
						if oLayer3Ipv6Addresses.Advertise.ValidLifetime != nil {
							nestedLayer3Ipv6Addresses.Advertise.ValidLifetime = oLayer3Ipv6Addresses.Advertise.ValidLifetime
						}
						if oLayer3Ipv6Addresses.Advertise.PreferredLifetime != nil {
							nestedLayer3Ipv6Addresses.Advertise.PreferredLifetime = oLayer3Ipv6Addresses.Advertise.PreferredLifetime
						}
					}
					nestedLayer3.Ipv6.Addresses = append(nestedLayer3.Ipv6.Addresses, nestedLayer3Ipv6Addresses)
				}
			}
			if o.Layer3.Ipv6.NeighborDiscovery != nil {
				nestedLayer3.Ipv6.NeighborDiscovery = &Layer3Ipv6NeighborDiscoveryXml{}
				if _, ok := o.Misc["Layer3Ipv6NeighborDiscovery"]; ok {
					nestedLayer3.Ipv6.NeighborDiscovery.Misc = o.Misc["Layer3Ipv6NeighborDiscovery"]
				}
				if o.Layer3.Ipv6.NeighborDiscovery.NsInterval != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.NeighborDiscovery.NsInterval
				}
				if o.Layer3.Ipv6.NeighborDiscovery.ReachableTime != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.ReachableTime = o.Layer3.Ipv6.NeighborDiscovery.ReachableTime
				}
				if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement = &Layer3Ipv6NeighborDiscoveryRouterAdvertisementXml{}
					if _, ok := o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisement"]; ok {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Misc = o.Misc["Layer3Ipv6NeighborDiscoveryRouterAdvertisement"]
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Enable, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
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
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
					}
					if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
					}
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
				if o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.EnableNdpMonitor = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
				}
				if o.Layer3.Ipv6.NeighborDiscovery.EnableDad != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.EnableDad = util.YesNo(o.Layer3.Ipv6.NeighborDiscovery.EnableDad, nil)
				}
				if o.Layer3.Ipv6.NeighborDiscovery.DadAttempts != nil {
					nestedLayer3.Ipv6.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.NeighborDiscovery.DadAttempts
				}
			}
			if o.Layer3.Ipv6.DnsServer != nil {
				nestedLayer3.Ipv6.DnsServer = &Layer3Ipv6DnsServerXml{}
				if _, ok := o.Misc["Layer3Ipv6DnsServer"]; ok {
					nestedLayer3.Ipv6.DnsServer.Misc = o.Misc["Layer3Ipv6DnsServer"]
				}
				if o.Layer3.Ipv6.DnsServer.Enable != nil {
					nestedLayer3.Ipv6.DnsServer.Enable = util.YesNo(o.Layer3.Ipv6.DnsServer.Enable, nil)
				}
				if o.Layer3.Ipv6.DnsServer.Source != nil {
					nestedLayer3.Ipv6.DnsServer.Source = &Layer3Ipv6DnsServerSourceXml{}
					if _, ok := o.Misc["Layer3Ipv6DnsServerSource"]; ok {
						nestedLayer3.Ipv6.DnsServer.Source.Misc = o.Misc["Layer3Ipv6DnsServerSource"]
					}
					if o.Layer3.Ipv6.DnsServer.Source.Dhcpv6 != nil {
						nestedLayer3.Ipv6.DnsServer.Source.Dhcpv6 = &Layer3Ipv6DnsServerSourceDhcpv6Xml{}
						if _, ok := o.Misc["Layer3Ipv6DnsServerSourceDhcpv6"]; ok {
							nestedLayer3.Ipv6.DnsServer.Source.Dhcpv6.Misc = o.Misc["Layer3Ipv6DnsServerSourceDhcpv6"]
						}
						if o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool != nil {
							nestedLayer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool
						}
					}
					if o.Layer3.Ipv6.DnsServer.Source.Manual != nil {
						nestedLayer3.Ipv6.DnsServer.Source.Manual = &Layer3Ipv6DnsServerSourceManualXml{}
						if _, ok := o.Misc["Layer3Ipv6DnsServerSourceManual"]; ok {
							nestedLayer3.Ipv6.DnsServer.Source.Manual.Misc = o.Misc["Layer3Ipv6DnsServerSourceManual"]
						}
						if o.Layer3.Ipv6.DnsServer.Source.Manual.Suffix != nil {
							nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix = []Layer3Ipv6DnsServerSourceManualSuffixXml{}
							for _, oLayer3Ipv6DnsServerSourceManualSuffix := range o.Layer3.Ipv6.DnsServer.Source.Manual.Suffix {
								nestedLayer3Ipv6DnsServerSourceManualSuffix := Layer3Ipv6DnsServerSourceManualSuffixXml{}
								if _, ok := o.Misc["Layer3Ipv6DnsServerSourceManualSuffix"]; ok {
									nestedLayer3Ipv6DnsServerSourceManualSuffix.Misc = o.Misc["Layer3Ipv6DnsServerSourceManualSuffix"]
								}
								if oLayer3Ipv6DnsServerSourceManualSuffix.Lifetime != nil {
									nestedLayer3Ipv6DnsServerSourceManualSuffix.Lifetime = oLayer3Ipv6DnsServerSourceManualSuffix.Lifetime
								}
								if oLayer3Ipv6DnsServerSourceManualSuffix.Name != "" {
									nestedLayer3Ipv6DnsServerSourceManualSuffix.Name = oLayer3Ipv6DnsServerSourceManualSuffix.Name
								}
								nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix = append(nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix, nestedLayer3Ipv6DnsServerSourceManualSuffix)
							}
						}
					}
				}
				if o.Layer3.Ipv6.DnsServer.DnsSupport != nil {
					nestedLayer3.Ipv6.DnsServer.DnsSupport = &Layer3Ipv6DnsServerDnsSupportXml{}
					if _, ok := o.Misc["Layer3Ipv6DnsServerDnsSupport"]; ok {
						nestedLayer3.Ipv6.DnsServer.DnsSupport.Misc = o.Misc["Layer3Ipv6DnsServerDnsSupport"]
					}
					if o.Layer3.Ipv6.DnsServer.DnsSupport.Server != nil {
						nestedLayer3.Ipv6.DnsServer.DnsSupport.Server = []Layer3Ipv6DnsServerDnsSupportServerXml{}
						for _, oLayer3Ipv6DnsServerDnsSupportServer := range o.Layer3.Ipv6.DnsServer.DnsSupport.Server {
							nestedLayer3Ipv6DnsServerDnsSupportServer := Layer3Ipv6DnsServerDnsSupportServerXml{}
							if _, ok := o.Misc["Layer3Ipv6DnsServerDnsSupportServer"]; ok {
								nestedLayer3Ipv6DnsServerDnsSupportServer.Misc = o.Misc["Layer3Ipv6DnsServerDnsSupportServer"]
							}
							if oLayer3Ipv6DnsServerDnsSupportServer.Name != "" {
								nestedLayer3Ipv6DnsServerDnsSupportServer.Name = oLayer3Ipv6DnsServerDnsSupportServer.Name
							}
							if oLayer3Ipv6DnsServerDnsSupportServer.Lifetime != nil {
								nestedLayer3Ipv6DnsServerDnsSupportServer.Lifetime = oLayer3Ipv6DnsServerDnsSupportServer.Lifetime
							}
							nestedLayer3.Ipv6.DnsServer.DnsSupport.Server = append(nestedLayer3.Ipv6.DnsServer.DnsSupport.Server, nestedLayer3Ipv6DnsServerDnsSupportServer)
						}
					}
					if o.Layer3.Ipv6.DnsServer.DnsSupport.Suffix != nil {
						nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix = []Layer3Ipv6DnsServerDnsSupportSuffixXml{}
						for _, oLayer3Ipv6DnsServerDnsSupportSuffix := range o.Layer3.Ipv6.DnsServer.DnsSupport.Suffix {
							nestedLayer3Ipv6DnsServerDnsSupportSuffix := Layer3Ipv6DnsServerDnsSupportSuffixXml{}
							if _, ok := o.Misc["Layer3Ipv6DnsServerDnsSupportSuffix"]; ok {
								nestedLayer3Ipv6DnsServerDnsSupportSuffix.Misc = o.Misc["Layer3Ipv6DnsServerDnsSupportSuffix"]
							}
							if oLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime != nil {
								nestedLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime = oLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime
							}
							if oLayer3Ipv6DnsServerDnsSupportSuffix.Name != "" {
								nestedLayer3Ipv6DnsServerDnsSupportSuffix.Name = oLayer3Ipv6DnsServerDnsSupportSuffix.Name
							}
							nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix = append(nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix, nestedLayer3Ipv6DnsServerDnsSupportSuffix)
						}
					}
					if o.Layer3.Ipv6.DnsServer.DnsSupport.Enable != nil {
						nestedLayer3.Ipv6.DnsServer.DnsSupport.Enable = util.YesNo(o.Layer3.Ipv6.DnsServer.DnsSupport.Enable, nil)
					}
				}
			}
		}
		if o.Layer3.InterfaceManagementProfile != nil {
			nestedLayer3.InterfaceManagementProfile = o.Layer3.InterfaceManagementProfile
		}
	}
	entry.Layer3 = nestedLayer3

	var nestedTap *TapXml
	if o.Tap != nil {
		nestedTap = &TapXml{}
		if _, ok := o.Misc["Tap"]; ok {
			nestedTap.Misc = o.Misc["Tap"]
		}
		if o.Tap.NetflowProfile != nil {
			nestedTap.NetflowProfile = o.Tap.NetflowProfile
		}
	}
	entry.Tap = nestedTap

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
		entry.LinkDuplex = o.LinkDuplex
		entry.LinkSpeed = o.LinkSpeed
		entry.LinkState = o.LinkState
		var nestedPoe *Poe
		if o.Poe != nil {
			nestedPoe = &Poe{}
			if o.Poe.Misc != nil {
				entry.Misc["Poe"] = o.Poe.Misc
			}
			if o.Poe.ReservedPower != nil {
				nestedPoe.ReservedPower = o.Poe.ReservedPower
			}
			if o.Poe.Enabled != nil {
				nestedPoe.Enabled = util.AsBool(o.Poe.Enabled, nil)
			}
		}
		entry.Poe = nestedPoe

		var nestedHa *Ha
		if o.Ha != nil {
			nestedHa = &Ha{}
			if o.Ha.Misc != nil {
				entry.Misc["Ha"] = o.Ha.Misc
			}
		}
		entry.Ha = nestedHa

		var nestedLayer3 *Layer3
		if o.Layer3 != nil {
			nestedLayer3 = &Layer3{}
			if o.Layer3.Misc != nil {
				entry.Misc["Layer3"] = o.Layer3.Misc
			}
			if o.Layer3.Ipv6 != nil {
				nestedLayer3.Ipv6 = &Layer3Ipv6{}
				if o.Layer3.Ipv6.Misc != nil {
					entry.Misc["Layer3Ipv6"] = o.Layer3.Ipv6.Misc
				}
				if o.Layer3.Ipv6.NeighborDiscovery != nil {
					nestedLayer3.Ipv6.NeighborDiscovery = &Layer3Ipv6NeighborDiscovery{}
					if o.Layer3.Ipv6.NeighborDiscovery.Misc != nil {
						entry.Misc["Layer3Ipv6NeighborDiscovery"] = o.Layer3.Ipv6.NeighborDiscovery.Misc
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
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.MaxInterval
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.LinkMtu
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.Lifetime
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
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.HopLimit
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference = o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.RouterPreference
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.ManagedFlag, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.OtherFlag, nil)
						}
						if o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck != nil {
							nestedLayer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.RouterAdvertisement.EnableConsistencyCheck, nil)
						}
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
					if o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.EnableNdpMonitor = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.EnableNdpMonitor, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.EnableDad != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.EnableDad = util.AsBool(o.Layer3.Ipv6.NeighborDiscovery.EnableDad, nil)
					}
					if o.Layer3.Ipv6.NeighborDiscovery.DadAttempts != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.DadAttempts = o.Layer3.Ipv6.NeighborDiscovery.DadAttempts
					}
					if o.Layer3.Ipv6.NeighborDiscovery.NsInterval != nil {
						nestedLayer3.Ipv6.NeighborDiscovery.NsInterval = o.Layer3.Ipv6.NeighborDiscovery.NsInterval
					}
				}
				if o.Layer3.Ipv6.DnsServer != nil {
					nestedLayer3.Ipv6.DnsServer = &Layer3Ipv6DnsServer{}
					if o.Layer3.Ipv6.DnsServer.Misc != nil {
						entry.Misc["Layer3Ipv6DnsServer"] = o.Layer3.Ipv6.DnsServer.Misc
					}
					if o.Layer3.Ipv6.DnsServer.Enable != nil {
						nestedLayer3.Ipv6.DnsServer.Enable = util.AsBool(o.Layer3.Ipv6.DnsServer.Enable, nil)
					}
					if o.Layer3.Ipv6.DnsServer.Source != nil {
						nestedLayer3.Ipv6.DnsServer.Source = &Layer3Ipv6DnsServerSource{}
						if o.Layer3.Ipv6.DnsServer.Source.Misc != nil {
							entry.Misc["Layer3Ipv6DnsServerSource"] = o.Layer3.Ipv6.DnsServer.Source.Misc
						}
						if o.Layer3.Ipv6.DnsServer.Source.Dhcpv6 != nil {
							nestedLayer3.Ipv6.DnsServer.Source.Dhcpv6 = &Layer3Ipv6DnsServerSourceDhcpv6{}
							if o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.Misc != nil {
								entry.Misc["Layer3Ipv6DnsServerSourceDhcpv6"] = o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.Misc
							}
							if o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool != nil {
								nestedLayer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool = o.Layer3.Ipv6.DnsServer.Source.Dhcpv6.PrefixPool
							}
						}
						if o.Layer3.Ipv6.DnsServer.Source.Manual != nil {
							nestedLayer3.Ipv6.DnsServer.Source.Manual = &Layer3Ipv6DnsServerSourceManual{}
							if o.Layer3.Ipv6.DnsServer.Source.Manual.Misc != nil {
								entry.Misc["Layer3Ipv6DnsServerSourceManual"] = o.Layer3.Ipv6.DnsServer.Source.Manual.Misc
							}
							if o.Layer3.Ipv6.DnsServer.Source.Manual.Suffix != nil {
								nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix = []Layer3Ipv6DnsServerSourceManualSuffix{}
								for _, oLayer3Ipv6DnsServerSourceManualSuffix := range o.Layer3.Ipv6.DnsServer.Source.Manual.Suffix {
									nestedLayer3Ipv6DnsServerSourceManualSuffix := Layer3Ipv6DnsServerSourceManualSuffix{}
									if oLayer3Ipv6DnsServerSourceManualSuffix.Misc != nil {
										entry.Misc["Layer3Ipv6DnsServerSourceManualSuffix"] = oLayer3Ipv6DnsServerSourceManualSuffix.Misc
									}
									if oLayer3Ipv6DnsServerSourceManualSuffix.Lifetime != nil {
										nestedLayer3Ipv6DnsServerSourceManualSuffix.Lifetime = oLayer3Ipv6DnsServerSourceManualSuffix.Lifetime
									}
									if oLayer3Ipv6DnsServerSourceManualSuffix.Name != "" {
										nestedLayer3Ipv6DnsServerSourceManualSuffix.Name = oLayer3Ipv6DnsServerSourceManualSuffix.Name
									}
									nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix = append(nestedLayer3.Ipv6.DnsServer.Source.Manual.Suffix, nestedLayer3Ipv6DnsServerSourceManualSuffix)
								}
							}
						}
					}
					if o.Layer3.Ipv6.DnsServer.DnsSupport != nil {
						nestedLayer3.Ipv6.DnsServer.DnsSupport = &Layer3Ipv6DnsServerDnsSupport{}
						if o.Layer3.Ipv6.DnsServer.DnsSupport.Misc != nil {
							entry.Misc["Layer3Ipv6DnsServerDnsSupport"] = o.Layer3.Ipv6.DnsServer.DnsSupport.Misc
						}
						if o.Layer3.Ipv6.DnsServer.DnsSupport.Suffix != nil {
							nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix = []Layer3Ipv6DnsServerDnsSupportSuffix{}
							for _, oLayer3Ipv6DnsServerDnsSupportSuffix := range o.Layer3.Ipv6.DnsServer.DnsSupport.Suffix {
								nestedLayer3Ipv6DnsServerDnsSupportSuffix := Layer3Ipv6DnsServerDnsSupportSuffix{}
								if oLayer3Ipv6DnsServerDnsSupportSuffix.Misc != nil {
									entry.Misc["Layer3Ipv6DnsServerDnsSupportSuffix"] = oLayer3Ipv6DnsServerDnsSupportSuffix.Misc
								}
								if oLayer3Ipv6DnsServerDnsSupportSuffix.Name != "" {
									nestedLayer3Ipv6DnsServerDnsSupportSuffix.Name = oLayer3Ipv6DnsServerDnsSupportSuffix.Name
								}
								if oLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime != nil {
									nestedLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime = oLayer3Ipv6DnsServerDnsSupportSuffix.Lifetime
								}
								nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix = append(nestedLayer3.Ipv6.DnsServer.DnsSupport.Suffix, nestedLayer3Ipv6DnsServerDnsSupportSuffix)
							}
						}
						if o.Layer3.Ipv6.DnsServer.DnsSupport.Enable != nil {
							nestedLayer3.Ipv6.DnsServer.DnsSupport.Enable = util.AsBool(o.Layer3.Ipv6.DnsServer.DnsSupport.Enable, nil)
						}
						if o.Layer3.Ipv6.DnsServer.DnsSupport.Server != nil {
							nestedLayer3.Ipv6.DnsServer.DnsSupport.Server = []Layer3Ipv6DnsServerDnsSupportServer{}
							for _, oLayer3Ipv6DnsServerDnsSupportServer := range o.Layer3.Ipv6.DnsServer.DnsSupport.Server {
								nestedLayer3Ipv6DnsServerDnsSupportServer := Layer3Ipv6DnsServerDnsSupportServer{}
								if oLayer3Ipv6DnsServerDnsSupportServer.Misc != nil {
									entry.Misc["Layer3Ipv6DnsServerDnsSupportServer"] = oLayer3Ipv6DnsServerDnsSupportServer.Misc
								}
								if oLayer3Ipv6DnsServerDnsSupportServer.Lifetime != nil {
									nestedLayer3Ipv6DnsServerDnsSupportServer.Lifetime = oLayer3Ipv6DnsServerDnsSupportServer.Lifetime
								}
								if oLayer3Ipv6DnsServerDnsSupportServer.Name != "" {
									nestedLayer3Ipv6DnsServerDnsSupportServer.Name = oLayer3Ipv6DnsServerDnsSupportServer.Name
								}
								nestedLayer3.Ipv6.DnsServer.DnsSupport.Server = append(nestedLayer3.Ipv6.DnsServer.DnsSupport.Server, nestedLayer3Ipv6DnsServerDnsSupportServer)
							}
						}
					}
				}
				if o.Layer3.Ipv6.Enabled != nil {
					nestedLayer3.Ipv6.Enabled = util.AsBool(o.Layer3.Ipv6.Enabled, nil)
				}
				if o.Layer3.Ipv6.InterfaceId != nil {
					nestedLayer3.Ipv6.InterfaceId = o.Layer3.Ipv6.InterfaceId
				}
				if o.Layer3.Ipv6.Addresses != nil {
					nestedLayer3.Ipv6.Addresses = []Layer3Ipv6Addresses{}
					for _, oLayer3Ipv6Addresses := range o.Layer3.Ipv6.Addresses {
						nestedLayer3Ipv6Addresses := Layer3Ipv6Addresses{}
						if oLayer3Ipv6Addresses.Misc != nil {
							entry.Misc["Layer3Ipv6Addresses"] = oLayer3Ipv6Addresses.Misc
						}
						if oLayer3Ipv6Addresses.EnableOnInterface != nil {
							nestedLayer3Ipv6Addresses.EnableOnInterface = util.AsBool(oLayer3Ipv6Addresses.EnableOnInterface, nil)
						}
						if oLayer3Ipv6Addresses.Prefix != nil {
							nestedLayer3Ipv6Addresses.Prefix = oLayer3Ipv6Addresses.Prefix
						}
						if oLayer3Ipv6Addresses.Anycast != nil {
							nestedLayer3Ipv6Addresses.Anycast = oLayer3Ipv6Addresses.Anycast
						}
						if oLayer3Ipv6Addresses.Advertise != nil {
							nestedLayer3Ipv6Addresses.Advertise = &Layer3Ipv6AddressesAdvertise{}
							if oLayer3Ipv6Addresses.Advertise.Misc != nil {
								entry.Misc["Layer3Ipv6AddressesAdvertise"] = oLayer3Ipv6Addresses.Advertise.Misc
							}
							if oLayer3Ipv6Addresses.Advertise.OnlinkFlag != nil {
								nestedLayer3Ipv6Addresses.Advertise.OnlinkFlag = util.AsBool(oLayer3Ipv6Addresses.Advertise.OnlinkFlag, nil)
							}
							if oLayer3Ipv6Addresses.Advertise.AutoConfigFlag != nil {
								nestedLayer3Ipv6Addresses.Advertise.AutoConfigFlag = util.AsBool(oLayer3Ipv6Addresses.Advertise.AutoConfigFlag, nil)
							}
							if oLayer3Ipv6Addresses.Advertise.Enable != nil {
								nestedLayer3Ipv6Addresses.Advertise.Enable = util.AsBool(oLayer3Ipv6Addresses.Advertise.Enable, nil)
							}
							if oLayer3Ipv6Addresses.Advertise.ValidLifetime != nil {
								nestedLayer3Ipv6Addresses.Advertise.ValidLifetime = oLayer3Ipv6Addresses.Advertise.ValidLifetime
							}
							if oLayer3Ipv6Addresses.Advertise.PreferredLifetime != nil {
								nestedLayer3Ipv6Addresses.Advertise.PreferredLifetime = oLayer3Ipv6Addresses.Advertise.PreferredLifetime
							}
						}
						if oLayer3Ipv6Addresses.Name != "" {
							nestedLayer3Ipv6Addresses.Name = oLayer3Ipv6Addresses.Name
						}
						nestedLayer3.Ipv6.Addresses = append(nestedLayer3.Ipv6.Addresses, nestedLayer3Ipv6Addresses)
					}
				}
			}
			if o.Layer3.InterfaceManagementProfile != nil {
				nestedLayer3.InterfaceManagementProfile = o.Layer3.InterfaceManagementProfile
			}
			if o.Layer3.NetflowProfile != nil {
				nestedLayer3.NetflowProfile = o.Layer3.NetflowProfile
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
					if o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp != nil {
						nestedLayer3.SdwanLinkSettings.UpstreamNat.StaticIp = o.Layer3.SdwanLinkSettings.UpstreamNat.StaticIp
					}
				}
			}
			if o.Layer3.AdjustTcpMss != nil {
				nestedLayer3.AdjustTcpMss = &Layer3AdjustTcpMss{}
				if o.Layer3.AdjustTcpMss.Misc != nil {
					entry.Misc["Layer3AdjustTcpMss"] = o.Layer3.AdjustTcpMss.Misc
				}
				if o.Layer3.AdjustTcpMss.Ipv4MssAdjustment != nil {
					nestedLayer3.AdjustTcpMss.Ipv4MssAdjustment = o.Layer3.AdjustTcpMss.Ipv4MssAdjustment
				}
				if o.Layer3.AdjustTcpMss.Ipv6MssAdjustment != nil {
					nestedLayer3.AdjustTcpMss.Ipv6MssAdjustment = o.Layer3.AdjustTcpMss.Ipv6MssAdjustment
				}
				if o.Layer3.AdjustTcpMss.Enable != nil {
					nestedLayer3.AdjustTcpMss.Enable = util.AsBool(o.Layer3.AdjustTcpMss.Enable, nil)
				}
			}
			if o.Layer3.Mtu != nil {
				nestedLayer3.Mtu = o.Layer3.Mtu
			}
			if o.Layer3.Ips != nil {
				nestedLayer3.Ips = []Layer3Ips{}
				for _, oLayer3Ips := range o.Layer3.Ips {
					nestedLayer3Ips := Layer3Ips{}
					if oLayer3Ips.Misc != nil {
						entry.Misc["Layer3Ips"] = oLayer3Ips.Misc
					}
					if oLayer3Ips.SdwanGateway != nil {
						nestedLayer3Ips.SdwanGateway = oLayer3Ips.SdwanGateway
					}
					if oLayer3Ips.Name != "" {
						nestedLayer3Ips.Name = oLayer3Ips.Name
					}
					nestedLayer3.Ips = append(nestedLayer3.Ips, nestedLayer3Ips)
				}
			}
			if o.Layer3.Lldp != nil {
				nestedLayer3.Lldp = &Layer3Lldp{}
				if o.Layer3.Lldp.Misc != nil {
					entry.Misc["Layer3Lldp"] = o.Layer3.Lldp.Misc
				}
				if o.Layer3.Lldp.Enable != nil {
					nestedLayer3.Lldp.Enable = util.AsBool(o.Layer3.Lldp.Enable, nil)
				}
				if o.Layer3.Lldp.Profile != nil {
					nestedLayer3.Lldp.Profile = o.Layer3.Lldp.Profile
				}
			}
			if o.Layer3.DhcpClient != nil {
				nestedLayer3.DhcpClient = &Layer3DhcpClient{}
				if o.Layer3.DhcpClient.Misc != nil {
					entry.Misc["Layer3DhcpClient"] = o.Layer3.DhcpClient.Misc
				}
				if o.Layer3.DhcpClient.DefaultRouteMetric != nil {
					nestedLayer3.DhcpClient.DefaultRouteMetric = o.Layer3.DhcpClient.DefaultRouteMetric
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
				if o.Layer3.DhcpClient.Enable != nil {
					nestedLayer3.DhcpClient.Enable = util.AsBool(o.Layer3.DhcpClient.Enable, nil)
				}
				if o.Layer3.DhcpClient.CreateDefaultRoute != nil {
					nestedLayer3.DhcpClient.CreateDefaultRoute = util.AsBool(o.Layer3.DhcpClient.CreateDefaultRoute, nil)
				}
			}
			if o.Layer3.Bonjour != nil {
				nestedLayer3.Bonjour = &Layer3Bonjour{}
				if o.Layer3.Bonjour.Misc != nil {
					entry.Misc["Layer3Bonjour"] = o.Layer3.Bonjour.Misc
				}
				if o.Layer3.Bonjour.Enable != nil {
					nestedLayer3.Bonjour.Enable = util.AsBool(o.Layer3.Bonjour.Enable, nil)
				}
			}
			if o.Layer3.UntaggedSubInterface != nil {
				nestedLayer3.UntaggedSubInterface = util.AsBool(o.Layer3.UntaggedSubInterface, nil)
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
			if o.Layer3.NdpProxy != nil {
				nestedLayer3.NdpProxy = util.AsBool(o.Layer3.NdpProxy, nil)
			}
		}
		entry.Layer3 = nestedLayer3

		var nestedTap *Tap
		if o.Tap != nil {
			nestedTap = &Tap{}
			if o.Tap.Misc != nil {
				entry.Misc["Tap"] = o.Tap.Misc
			}
			if o.Tap.NetflowProfile != nil {
				nestedTap.NetflowProfile = o.Tap.NetflowProfile
			}
		}
		entry.Tap = nestedTap

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
	if !util.StringsMatch(a.LinkDuplex, b.LinkDuplex) {
		return false
	}
	if !util.StringsMatch(a.LinkSpeed, b.LinkSpeed) {
		return false
	}
	if !util.StringsMatch(a.LinkState, b.LinkState) {
		return false
	}
	if !matchPoe(a.Poe, b.Poe) {
		return false
	}
	if !matchHa(a.Ha, b.Ha) {
		return false
	}
	if !matchLayer3(a.Layer3, b.Layer3) {
		return false
	}
	if !matchTap(a.Tap, b.Tap) {
		return false
	}

	return true
}

func matchPoe(a *Poe, b *Poe) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !util.Ints64Match(a.ReservedPower, b.ReservedPower) {
		return false
	}
	return true
}
func matchTap(a *Tap, b *Tap) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
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
func matchLayer3Lldp(a *Layer3Lldp, b *Layer3Lldp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
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
	if !util.Ints64Match(a.DefaultRouteMetric, b.DefaultRouteMetric) {
		return false
	}
	if !matchLayer3DhcpClientSendHostname(a.SendHostname, b.SendHostname) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.CreateDefaultRoute, b.CreateDefaultRoute) {
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
	return true
}
func matchLayer3Ips(a []Layer3Ips, b []Layer3Ips) bool {
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
func matchLayer3Ipv6AddressesAdvertise(a *Layer3Ipv6AddressesAdvertise, b *Layer3Ipv6AddressesAdvertise) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
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
	if !util.BoolsMatch(a.OnlinkFlag, b.OnlinkFlag) {
		return false
	}
	return true
}
func matchLayer3Ipv6Addresses(a []Layer3Ipv6Addresses, b []Layer3Ipv6Addresses) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Anycast, b.Anycast) {
				return false
			}
			if !matchLayer3Ipv6AddressesAdvertise(a.Advertise, b.Advertise) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
				return false
			}
			if !util.StringsMatch(a.Prefix, b.Prefix) {
				return false
			}
		}
	}
	return true
}
func matchLayer3Ipv6NeighborDiscoveryRouterAdvertisement(a *Layer3Ipv6NeighborDiscoveryRouterAdvertisement, b *Layer3Ipv6NeighborDiscoveryRouterAdvertisement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.MaxInterval, b.MaxInterval) {
		return false
	}
	if !util.StringsMatch(a.LinkMtu, b.LinkMtu) {
		return false
	}
	if !util.Ints64Match(a.Lifetime, b.Lifetime) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !util.StringsMatch(a.RetransmissionTimer, b.RetransmissionTimer) {
		return false
	}
	if !util.StringsMatch(a.HopLimit, b.HopLimit) {
		return false
	}
	if !util.StringsMatch(a.RouterPreference, b.RouterPreference) {
		return false
	}
	if !util.BoolsMatch(a.ManagedFlag, b.ManagedFlag) {
		return false
	}
	if !util.BoolsMatch(a.OtherFlag, b.OtherFlag) {
		return false
	}
	if !util.BoolsMatch(a.EnableConsistencyCheck, b.EnableConsistencyCheck) {
		return false
	}
	if !util.Ints64Match(a.MinInterval, b.MinInterval) {
		return false
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
func matchLayer3Ipv6NeighborDiscovery(a *Layer3Ipv6NeighborDiscovery, b *Layer3Ipv6NeighborDiscovery) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.ReachableTime, b.ReachableTime) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryRouterAdvertisement(a.RouterAdvertisement, b.RouterAdvertisement) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscoveryNeighbor(a.Neighbor, b.Neighbor) {
		return false
	}
	if !util.BoolsMatch(a.EnableNdpMonitor, b.EnableNdpMonitor) {
		return false
	}
	if !util.BoolsMatch(a.EnableDad, b.EnableDad) {
		return false
	}
	if !util.Ints64Match(a.DadAttempts, b.DadAttempts) {
		return false
	}
	if !util.Ints64Match(a.NsInterval, b.NsInterval) {
		return false
	}
	return true
}
func matchLayer3Ipv6DnsServerSourceDhcpv6(a *Layer3Ipv6DnsServerSourceDhcpv6, b *Layer3Ipv6DnsServerSourceDhcpv6) bool {
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
func matchLayer3Ipv6DnsServerSourceManualSuffix(a []Layer3Ipv6DnsServerSourceManualSuffix, b []Layer3Ipv6DnsServerSourceManualSuffix) bool {
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
func matchLayer3Ipv6DnsServerSourceManual(a *Layer3Ipv6DnsServerSourceManual, b *Layer3Ipv6DnsServerSourceManual) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DnsServerSourceManualSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchLayer3Ipv6DnsServerSource(a *Layer3Ipv6DnsServerSource, b *Layer3Ipv6DnsServerSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3Ipv6DnsServerSourceDhcpv6(a.Dhcpv6, b.Dhcpv6) {
		return false
	}
	if !matchLayer3Ipv6DnsServerSourceManual(a.Manual, b.Manual) {
		return false
	}
	return true
}
func matchLayer3Ipv6DnsServerDnsSupportServer(a []Layer3Ipv6DnsServerDnsSupportServer, b []Layer3Ipv6DnsServerDnsSupportServer) bool {
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
func matchLayer3Ipv6DnsServerDnsSupportSuffix(a []Layer3Ipv6DnsServerDnsSupportSuffix, b []Layer3Ipv6DnsServerDnsSupportSuffix) bool {
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
func matchLayer3Ipv6DnsServerDnsSupport(a *Layer3Ipv6DnsServerDnsSupport, b *Layer3Ipv6DnsServerDnsSupport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6DnsServerDnsSupportServer(a.Server, b.Server) {
		return false
	}
	if !matchLayer3Ipv6DnsServerDnsSupportSuffix(a.Suffix, b.Suffix) {
		return false
	}
	return true
}
func matchLayer3Ipv6DnsServer(a *Layer3Ipv6DnsServer, b *Layer3Ipv6DnsServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !matchLayer3Ipv6DnsServerSource(a.Source, b.Source) {
		return false
	}
	if !matchLayer3Ipv6DnsServerDnsSupport(a.DnsSupport, b.DnsSupport) {
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
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !util.StringsMatch(a.InterfaceId, b.InterfaceId) {
		return false
	}
	if !matchLayer3Ipv6Addresses(a.Addresses, b.Addresses) {
		return false
	}
	if !matchLayer3Ipv6NeighborDiscovery(a.NeighborDiscovery, b.NeighborDiscovery) {
		return false
	}
	if !matchLayer3Ipv6DnsServer(a.DnsServer, b.DnsServer) {
		return false
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
	if !util.StringsMatch(a.StaticIp, b.StaticIp) {
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
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaceProfile, b.SdwanInterfaceProfile) {
		return false
	}
	if !matchLayer3SdwanLinkSettingsUpstreamNat(a.UpstreamNat, b.UpstreamNat) {
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
func matchLayer3(a *Layer3, b *Layer3) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchLayer3SdwanLinkSettings(a.SdwanLinkSettings, b.SdwanLinkSettings) {
		return false
	}
	if !matchLayer3AdjustTcpMss(a.AdjustTcpMss, b.AdjustTcpMss) {
		return false
	}
	if !util.Ints64Match(a.Mtu, b.Mtu) {
		return false
	}
	if !matchLayer3Ips(a.Ips, b.Ips) {
		return false
	}
	if !matchLayer3Ipv6(a.Ipv6, b.Ipv6) {
		return false
	}
	if !util.StringsMatch(a.InterfaceManagementProfile, b.InterfaceManagementProfile) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
		return false
	}
	if !util.BoolsMatch(a.UntaggedSubInterface, b.UntaggedSubInterface) {
		return false
	}
	if !matchLayer3Arp(a.Arp, b.Arp) {
		return false
	}
	if !util.BoolsMatch(a.NdpProxy, b.NdpProxy) {
		return false
	}
	if !matchLayer3Lldp(a.Lldp, b.Lldp) {
		return false
	}
	if !matchLayer3DhcpClient(a.DhcpClient, b.DhcpClient) {
		return false
	}
	if !matchLayer3Bonjour(a.Bonjour, b.Bonjour) {
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
