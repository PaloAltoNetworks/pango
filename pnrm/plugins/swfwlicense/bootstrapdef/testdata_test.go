package bootstrapdef

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
			Entry{Name: "bootstrapdef",
				Description: "some desc",
				Authcode:    "abcdefg",
			},
		},
	}
}
