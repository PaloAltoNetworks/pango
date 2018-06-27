package netw


import (
    "github.com/PaloAltoNetworks/pango/netw/interface/eth"
    "github.com/PaloAltoNetworks/pango/netw/interface/loopback"
    "github.com/PaloAltoNetworks/pango/netw/interface/tunnel"
    vli "github.com/PaloAltoNetworks/pango/netw/interface/vlan"
    "github.com/PaloAltoNetworks/pango/netw/routing/router"
    "github.com/PaloAltoNetworks/pango/netw/vlan"
    "github.com/PaloAltoNetworks/pango/util"
)


// PanoNetw is the client.Network namespace.
type PanoNetw struct {
    EthernetInterface *eth.PanoEth
    TunnelInterface *tunnel.PanoTunnel
    LoopbackInterface *loopback.PanoLoopback
    VirtualRouter *router.PanoRouter
    Vlan *vlan.PanoVlan
    VlanInterface *vli.PanoVlan
}

// Initialize is invoked on client.Initialize().
func (c *PanoNetw) Initialize(i util.XapiClient) {
    c.EthernetInterface = &eth.PanoEth{}
    c.EthernetInterface.Initialize(i)

    c.TunnelInterface = &tunnel.PanoTunnel{}
    c.TunnelInterface.Initialize(i)

    c.LoopbackInterface = &loopback.PanoLoopback{}
    c.LoopbackInterface.Initialize(i)

    c.VirtualRouter = &router.PanoRouter{}
    c.VirtualRouter.Initialize(i)

    c.Vlan = &vlan.PanoVlan{}
    c.Vlan.Initialize(i)

    c.VlanInterface = &vli.PanoVlan{}
    c.VlanInterface.Initialize(i)
}
