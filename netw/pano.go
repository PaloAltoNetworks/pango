package netw

import (
	"github.com/PaloAltoNetworks/pango/netw/ikegw"
	aggeth "github.com/PaloAltoNetworks/pango/netw/interface/aggregate"
	"github.com/PaloAltoNetworks/pango/netw/interface/arp"
	"github.com/PaloAltoNetworks/pango/netw/interface/eth"
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
	"github.com/PaloAltoNetworks/pango/netw/routing/route/static/ipv4"
	"github.com/PaloAltoNetworks/pango/netw/routing/router"
	"github.com/PaloAltoNetworks/pango/netw/tunnel/gre"
	"github.com/PaloAltoNetworks/pango/netw/vlan"
	"github.com/PaloAltoNetworks/pango/netw/zone"
	"github.com/PaloAltoNetworks/pango/util"
)

// PanoNetw is the client.Network namespace.
type PanoNetw struct {
	AggregateInterface       *aggeth.Panorama
	Arp                      *arp.Panorama
	BfdProfile               *bfd.PanoBfd
	BgpAggregate             *aggregate.PanoAggregate
	BgpAggAdvertiseFilter    *agaf.PanoAdvertise
	BgpAggSuppressFilter     *suppress.PanoSuppress
	BgpAuthProfile           *auth.PanoAuth
	BgpConAdvAdvertiseFilter *advertise.PanoAdvertise
	BgpConAdvNonExistFilter  *nonexist.PanoNonExist
	BgpConditionalAdv        *conadv.PanoConAdv
	BgpConfig                *bgp.Panorama
	BgpDampeningProfile      *dampening.PanoDampening
	BgpExport                *exp.PanoExp
	BgpImport                *imp.PanoImp
	BgpPeer                  *peer.PanoPeer
	BgpPeerGroup             *group.PanoGroup
	BgpRedistRule            *bgpredist.PanoRedist
	EthernetInterface        *eth.Panorama
	GreTunnel                *gre.PanoGre
	IkeCryptoProfile         *ike.PanoIke
	IkeGateway               *ikegw.PanoIkeGw
	IpsecCryptoProfile       *ipsec.PanoIpsec
	IpsecTunnel              *ipsectunnel.PanoIpsecTunnel
	IpsecTunnelProxyId       *tpiv4.PanoIpv4
	Layer2Subinterface       *layer2.Panorama
	Layer3Subinterface       *layer3.Panorama
	LoopbackInterface        *loopback.Panorama
	ManagementProfile        *mngtprof.PanoMngtProf
	MonitorProfile           *monitor.PanoMonitor
	RedistributionProfile    *redist4.PanoIpv4
	StaticRoute              *ipv4.Panorama
	TunnelInterface          *tunnel.Panorama
	VirtualRouter            *router.Panorama
	Vlan                     *vlan.Panorama
	VlanInterface            *vli.Panorama
	Zone                     *zone.Panorama
}

// Initialize is invoked on client.Initialize().
func (c *PanoNetw) Initialize(i util.XapiClient) {
	c.AggregateInterface = aggeth.PanoramaNamespace(i)

	c.Arp = arp.PanoramaNamespace(i)

	c.BfdProfile = &bfd.PanoBfd{}
	c.BfdProfile.Initialize(i)

	c.BgpAggregate = &aggregate.PanoAggregate{}
	c.BgpAggregate.Initialize(i)

	c.BgpAggAdvertiseFilter = &agaf.PanoAdvertise{}
	c.BgpAggAdvertiseFilter.Initialize(i)

	c.BgpAggSuppressFilter = &suppress.PanoSuppress{}
	c.BgpAggSuppressFilter.Initialize(i)

	c.BgpAuthProfile = &auth.PanoAuth{}
	c.BgpAuthProfile.Initialize(i)

	c.BgpConAdvAdvertiseFilter = &advertise.PanoAdvertise{}
	c.BgpConAdvAdvertiseFilter.Initialize(i)

	c.BgpConAdvNonExistFilter = &nonexist.PanoNonExist{}
	c.BgpConAdvNonExistFilter.Initialize(i)

	c.BgpConditionalAdv = &conadv.PanoConAdv{}
	c.BgpConditionalAdv.Initialize(i)

	c.BgpConfig = bgp.PanoramaNamespace(i)

	c.BgpDampeningProfile = &dampening.PanoDampening{}
	c.BgpDampeningProfile.Initialize(i)

	c.BgpExport = &exp.PanoExp{}
	c.BgpExport.Initialize(i)

	c.BgpImport = &imp.PanoImp{}
	c.BgpImport.Initialize(i)

	c.BgpPeer = &peer.PanoPeer{}
	c.BgpPeer.Initialize(i)

	c.BgpPeerGroup = &group.PanoGroup{}
	c.BgpPeerGroup.Initialize(i)

	c.BgpRedistRule = &bgpredist.PanoRedist{}
	c.BgpRedistRule.Initialize(i)

	c.EthernetInterface = eth.PanoramaNamespace(i)

	c.GreTunnel = &gre.PanoGre{}
	c.GreTunnel.Initialize(i)

	c.IkeCryptoProfile = &ike.PanoIke{}
	c.IkeCryptoProfile.Initialize(i)

	c.IkeGateway = &ikegw.PanoIkeGw{}
	c.IkeGateway.Initialize(i)

	c.IpsecCryptoProfile = &ipsec.PanoIpsec{}
	c.IpsecCryptoProfile.Initialize(i)

	c.IpsecTunnel = &ipsectunnel.PanoIpsecTunnel{}
	c.IpsecTunnel.Initialize(i)

	c.IpsecTunnelProxyId = &tpiv4.PanoIpv4{}
	c.IpsecTunnelProxyId.Initialize(i)

	c.Layer2Subinterface = layer2.PanoramaNamespace(i)
	c.Layer3Subinterface = layer3.PanoramaNamespace(i)
	c.LoopbackInterface = loopback.PanoramaNamespace(i)

	c.ManagementProfile = &mngtprof.PanoMngtProf{}
	c.ManagementProfile.Initialize(i)

	c.MonitorProfile = &monitor.PanoMonitor{}
	c.MonitorProfile.Initialize(i)

	c.RedistributionProfile = &redist4.PanoIpv4{}
	c.RedistributionProfile.Initialize(i)

	c.StaticRoute = ipv4.PanoramaNamespace(i)
	c.TunnelInterface = tunnel.PanoramaNamespace(i)
	c.VirtualRouter = router.PanoramaNamespace(i)
	c.Vlan = vlan.PanoramaNamespace(i)
	c.VlanInterface = vli.PanoramaNamespace(i)
	c.Zone = zone.PanoramaNamespace(i)
}
