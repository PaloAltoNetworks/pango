// Package poli is the client.Policies namespace.
package poli

import (
    "github.com/PaloAltoNetworks/xapi/util"

    "github.com/PaloAltoNetworks/xapi/poli/security"
    "github.com/PaloAltoNetworks/xapi/poli/nat"
)


// Poli is the client.Policies namespace.
type Poli struct {
    Security *security.Security
    Nat *nat.Nat
}

// Initialize is invoked on client.Initialize().
func (c *Poli) Initialize(i util.XapiClient) {
    c.Security = &security.Security{}
    c.Security.Initialize(i)

    c.Nat = &nat.Nat{}
    c.Nat.Initialize(i)
}
