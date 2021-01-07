package predefined

import (
	"github.com/PaloAltoNetworks/pango/predefined/threat"
	"github.com/PaloAltoNetworks/pango/util"
)

type Panorama struct {
	Threat *threat.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		Threat: threat.PanoramaNamespace(x),
	}
}
