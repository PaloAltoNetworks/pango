package ospf

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

func versioning(v version.Number) (normalizer, func(Config) interface{}) {
	return &container_v1{}, specify_v1
}

func specifier(e Config) []namespace.Specifier {
	return []namespace.Specifier{e}
}

func container(v version.Number) normalizer {
	r, _ := versioning(v)
	return r
}

func first(ans normalizer, err error) (Config, error) {
	if err != nil {
		return Config{}, err
	}

	// TODO(shinmog): return the real ObjectNotFound error.
	list := ans.Normalize()
	if len(list) == 0 {
		return Config{}, fmt.Errorf("Object not found")
	}

	return list[0], nil
}

// FirewallNamespace returns an initialized namespace.
func FirewallNamespace(client util.XapiClient) *Firewall {
	return &Firewall{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Client:   client,
			},
		},
	}
}

// PanoramaNamespace returns an initialized namespace.
func PanoramaNamespace(client util.XapiClient) *Panorama {
	return &Panorama{
		ns: &namespace.Standard{
			Common: namespace.Common{
				Singular: singular,
				Client:   client,
			},
		},
	}
}
