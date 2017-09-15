// Package objs is the client.Objects namespace.
package objs


import (
    "github.com/PaloAltoNetworks/xapi/util"

    "github.com/PaloAltoNetworks/xapi/objs/addr"
    "github.com/PaloAltoNetworks/xapi/objs/srvc"
)


// Objs is the client.Network namespace.
type Objs struct {
    Address *addr.Addr
    Services *srvc.Srvc
}

// Initialize is invoked on client.Initialize().
func (c *Objs) Initialize(i util.XapiClient) {
    c.Address = &addr.Addr{}
    c.Address.Initialize(i)

    c.Services = &srvc.Srvc{}
    c.Services.Initialize(i)
}
