// Package netw is the client.Network namespace.
package netw


import (
    "github.com/PaloAltoNetworks/xapi/util"

    "github.com/PaloAltoNetworks/xapi/netw/eth"
    "github.com/PaloAltoNetworks/xapi/netw/zone"
    "github.com/PaloAltoNetworks/xapi/netw/mngtprof"
)


// Netw is the client.Network namespace.
type Netw struct {
    EthernetInterface *eth.Eth
    Zone *zone.Zone
    ManagementProfile *mngtprof.MngtProf
}

// Initialize is invoked on client.Initialize().
func (c *Netw) Initialize(i util.XapiClient) {
    c.EthernetInterface = &eth.Eth{}
    c.EthernetInterface.Initialize(i)

    c.Zone = &zone.Zone{}
    c.Zone.Initialize(i)

    c.ManagementProfile = &mngtprof.MngtProf{}
    c.ManagementProfile.Initialize(i)
}
