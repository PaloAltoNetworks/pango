package poli

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/poli/nat"
	"github.com/PaloAltoNetworks/pango/poli/pbf"
	"github.com/PaloAltoNetworks/pango/poli/security"
)

// Panorama is the client.Policies namespace.
type Panorama struct {
	Nat                   *nat.Panorama
	PolicyBasedForwarding *pbf.Panorama
	Security              *security.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		Nat:                   nat.PanoramaNamespace(x),
		PolicyBasedForwarding: pbf.PanoramaNamespace(x),
		Security:              security.PanoramaNamespace(x),
	}
}
