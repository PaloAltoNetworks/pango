package poli

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/poli/nat"
	"github.com/PaloAltoNetworks/pango/poli/pbf"
	"github.com/PaloAltoNetworks/pango/poli/security"
)

// Firewall is the client.Policies namespace.
type Firewall struct {
	Nat                   *nat.Firewall
	PolicyBasedForwarding *pbf.Firewall
	Security              *security.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Nat:                   nat.FirewallNamespace(x),
		PolicyBasedForwarding: pbf.FirewallNamespace(x),
		Security:              security.FirewallNamespace(x),
	}
}
