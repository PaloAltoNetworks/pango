package netw

import (
	"github.com/PaloAltoNetworks/pango/netw/ikegw"
	aggeth "github.com/PaloAltoNetworks/pango/netw/interface/aggregate"
	"github.com/PaloAltoNetworks/pango/netw/interface/arp"
	"github.com/PaloAltoNetworks/pango/netw/interface/eth"
	ipv6a "github.com/PaloAltoNetworks/pango/netw/interface/ipv6/address"
	ipv6n "github.com/PaloAltoNetworks/pango/netw/interface/ipv6/neighbor"
	"github.com/PaloAltoNetworks/pango/netw/interface/loopback"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer2"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer3"
	"github.com/PaloAltoNetworks/pango/netw/interface/tunnel"
	vli "github.com/PaloAltoNetworks/pango/netw/interface/vlan"
	"github.com/PaloAltoNetworks/pango/netw/ipsectunnel"
	tpiv4 "github.com/PaloAltoNetworks/pango/netw/ipsectunnel/proxyid/ipv4"
	"github.com/PaloAltoNetworks/pango/netw/profile/bfd"
	"github.com/PaloAltoNetworks/pango/netw/profile/ike"
	"github.com/PaloAltoNetworks/pango/netw/profile/ipsec"
	"github.com/PaloAltoNetworks/pango/netw/profile/mngtprof"
	"github.com/PaloAltoNetworks/pango/netw/profile/monitor"
	redist4 "github.com/PaloAltoNetworks/pango/netw/routing/profile/redist/ipv4"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/aggregate"
	agaf "github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/aggregate/filter/advertise"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/aggregate/filter/suppress"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/conadv"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/conadv/filter/advertise"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/conadv/filter/nonexist"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/exp"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/imp"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/peer"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/peer/group"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/profile/auth"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/profile/dampening"
	bgpredist "github.com/PaloAltoNetworks/pango/netw/routing/protocol/bgp/redist"
	"github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf"
	ospfarea "github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf/area"
	ospfint "github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf/area/iface"
	ospfvlink "github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf/area/vlink"
	ospfexp "github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf/exp"
	ospfauth "github.com/PaloAltoNetworks/pango/netw/routing/protocol/ospf/profile/auth"
	"github.com/PaloAltoNetworks/pango/netw/routing/route/static/ipv4"
	ipv6sr "github.com/PaloAltoNetworks/pango/netw/routing/route/static/ipv6"
	"github.com/PaloAltoNetworks/pango/netw/routing/router"
	"github.com/PaloAltoNetworks/pango/netw/tunnel/gre"
	"github.com/PaloAltoNetworks/pango/netw/vlan"
	"github.com/PaloAltoNetworks/pango/netw/zone"
	"github.com/PaloAltoNetworks/pango/util"
)

// Netw is the client.Network namespace.
type FwNetw struct {
	AggregateInterface       *aggeth.Firewall
	Arp                      *arp.Firewall
	BfdProfile               *bfd.FwBfd
	BgpAggregate             *aggregate.FwAggregate
	BgpAggAdvertiseFilter    *agaf.FwAdvertise
	BgpAggSuppressFilter     *suppress.FwSuppress
	BgpAuthProfile           *auth.FwAuth
	BgpConAdvAdvertiseFilter *advertise.FwAdvertise
	BgpConAdvNonExistFilter  *nonexist.FwNonExist
	BgpConditionalAdv        *conadv.FwConAdv
	BgpConfig                *bgp.Firewall
	BgpDampeningProfile      *dampening.FwDampening
	BgpExport                *exp.FwExp
	BgpImport                *imp.FwImp
	BgpPeer                  *peer.FwPeer
	BgpPeerGroup             *group.FwGroup
	BgpRedistRule            *bgpredist.FwRedist
	EthernetInterface        *eth.Firewall
	GreTunnel                *gre.FwGre
	IkeCryptoProfile         *ike.FwIke
	IkeGateway               *ikegw.FwIkeGw
	IpsecCryptoProfile       *ipsec.FwIpsec
	IpsecTunnel              *ipsectunnel.FwIpsecTunnel
	IpsecTunnelProxyId       *tpiv4.FwIpv4
	Ipv6Address              *ipv6a.Firewall
	Ipv6NeighborDiscovery    *ipv6n.Firewall
	Ipv6StaticRoute          *ipv6sr.Firewall
	Layer2Subinterface       *layer2.Firewall
	Layer3Subinterface       *layer3.Firewall
	LoopbackInterface        *loopback.Firewall
	ManagementProfile        *mngtprof.FwMngtProf
	MonitorProfile           *monitor.FwMonitor
	OspfArea                 *ospfarea.Firewall
	OspfAreaInterface        *ospfint.Firewall
	OspfAreaVirtualLink      *ospfvlink.Firewall
	OspfAuthProfile          *ospfauth.Firewall
	OspfConfig               *ospf.Firewall
	OspfExport               *ospfexp.Firewall
	RedistributionProfile    *redist4.FwIpv4
	StaticRoute              *ipv4.Firewall
	TunnelInterface          *tunnel.Firewall
	VirtualRouter            *router.Firewall
	Vlan                     *vlan.Firewall
	VlanInterface            *vli.Firewall
	Zone                     *zone.Firewall
}

// Initialize is invoked on client.Initialize().
func (c *FwNetw) Initialize(i util.XapiClient) {
	c.AggregateInterface = aggeth.FirewallNamespace(i)

	c.Arp = arp.FirewallNamespace(i)

	c.BfdProfile = &bfd.FwBfd{}
	c.BfdProfile.Initialize(i)

	c.BgpAggregate = &aggregate.FwAggregate{}
	c.BgpAggregate.Initialize(i)

	c.BgpAggAdvertiseFilter = &agaf.FwAdvertise{}
	c.BgpAggAdvertiseFilter.Initialize(i)

	c.BgpAggSuppressFilter = &suppress.FwSuppress{}
	c.BgpAggSuppressFilter.Initialize(i)

	c.BgpAuthProfile = &auth.FwAuth{}
	c.BgpAuthProfile.Initialize(i)

	c.BgpConAdvAdvertiseFilter = &advertise.FwAdvertise{}
	c.BgpConAdvAdvertiseFilter.Initialize(i)

	c.BgpConAdvNonExistFilter = &nonexist.FwNonExist{}
	c.BgpConAdvNonExistFilter.Initialize(i)

	c.BgpConditionalAdv = &conadv.FwConAdv{}
	c.BgpConditionalAdv.Initialize(i)

	c.BgpConfig = bgp.FirewallNamespace(i)

	c.BgpDampeningProfile = &dampening.FwDampening{}
	c.BgpDampeningProfile.Initialize(i)

	c.BgpExport = &exp.FwExp{}
	c.BgpExport.Initialize(i)

	c.BgpImport = &imp.FwImp{}
	c.BgpImport.Initialize(i)

	c.BgpPeer = &peer.FwPeer{}
	c.BgpPeer.Initialize(i)

	c.BgpPeerGroup = &group.FwGroup{}
	c.BgpPeerGroup.Initialize(i)

	c.BgpRedistRule = &bgpredist.FwRedist{}
	c.BgpRedistRule.Initialize(i)

	c.EthernetInterface = eth.FirewallNamespace(i)

	c.GreTunnel = &gre.FwGre{}
	c.GreTunnel.Initialize(i)

	c.IkeCryptoProfile = &ike.FwIke{}
	c.IkeCryptoProfile.Initialize(i)

	c.IkeGateway = &ikegw.FwIkeGw{}
	c.IkeGateway.Initialize(i)

	c.IpsecCryptoProfile = &ipsec.FwIpsec{}
	c.IpsecCryptoProfile.Initialize(i)

	c.IpsecTunnel = &ipsectunnel.FwIpsecTunnel{}
	c.IpsecTunnel.Initialize(i)

	c.IpsecTunnelProxyId = &tpiv4.FwIpv4{}
	c.IpsecTunnelProxyId.Initialize(i)

	c.Ipv6Address = ipv6a.FirewallNamespace(i)
	c.Ipv6NeighborDiscovery = ipv6n.FirewallNamespace(i)
	c.Ipv6StaticRoute = ipv6sr.FirewallNamespace(i)
	c.Layer2Subinterface = layer2.FirewallNamespace(i)
	c.Layer3Subinterface = layer3.FirewallNamespace(i)
	c.LoopbackInterface = loopback.FirewallNamespace(i)

	c.ManagementProfile = &mngtprof.FwMngtProf{}
	c.ManagementProfile.Initialize(i)

	c.MonitorProfile = &monitor.FwMonitor{}
	c.MonitorProfile.Initialize(i)

	c.OspfArea = ospfarea.FirewallNamespace(i)

	c.OspfAreaInterface = ospfint.FirewallNamespace(i)

	c.OspfAreaVirtualLink = ospfvlink.FirewallNamespace(i)

	c.OspfAuthProfile = ospfauth.FirewallNamespace(i)

	c.OspfConfig = ospf.FirewallNamespace(i)

	c.OspfExport = ospfexp.FirewallNamespace(i)

	c.RedistributionProfile = &redist4.FwIpv4{}
	c.RedistributionProfile.Initialize(i)

	c.StaticRoute = ipv4.FirewallNamespace(i)
	c.TunnelInterface = tunnel.FirewallNamespace(i)
	c.VirtualRouter = router.FirewallNamespace(i)
	c.Vlan = vlan.FirewallNamespace(i)
	c.VlanInterface = vli.FirewallNamespace(i)
	c.Zone = zone.FirewallNamespace(i)
}
