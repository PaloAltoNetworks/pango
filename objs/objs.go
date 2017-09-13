// Package objs is the client.Objects namespace.
package objs


import (
    "github.com/PaloAltoNetworks/xapi/util"

    "github.com/PaloAltoNetworks/xapi/objs/addr"
)


// Objs is the client.Network namespace.
type Objs struct {
    Address *addr.Addr
}

// Initialize is invoked on client.Initialize().
func (c *Objs) Initialize(i util.XapiClient) {
    c.Address = &addr.Addr{}
    c.Address.Initialize(i)
}
