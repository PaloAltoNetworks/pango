package cloudwatch

import (
	"github.com/PaloAltoNetworks/pango/plugin"
)

type tc struct {
	desc   string
	plugin plugin.Info
	conf   Config
}

func getTests() []tc {
	return []tc{
		{
			"basic config",
			plugin.Info{Name: "vm_series", Installed: "yes", Version: "1.0.0"},
			Config{
				Enabled:        true,
				Namespace:      "ns",
				UpdateInterval: 42,
			},
		},
	}
}
