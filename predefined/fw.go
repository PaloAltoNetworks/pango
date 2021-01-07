package predefined

import (
	"github.com/PaloAltoNetworks/pango/predefined/threat"
	"github.com/PaloAltoNetworks/pango/util"
)

type Firewall struct {
	Threat *threat.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Threat: threat.FirewallNamespace(x),
	}
}
