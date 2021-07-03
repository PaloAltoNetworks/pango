package manager

import (
	"github.com/PaloAltoNetworks/pango/plugin"
)

type tc struct {
	desc   string
	plugin plugin.Info
	conf   Entry
}

func getTests() []tc {
	return []tc{
		{
			"basic project",
			plugin.Info{Name: "sw_fw_license", Installed: "yes", Version: "1.0.0"},
			Entry{
				Name:                "licmanager",
				Description:         "some desc",
				DeviceGroup:         "dg",
				TemplateStack:       "ts",
				BootstrapDefinition: "bsdef",
				AutoDeactivate:      5,
			},
		},
	}
}
