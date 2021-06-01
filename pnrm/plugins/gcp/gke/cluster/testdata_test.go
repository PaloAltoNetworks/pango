package cluster

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
		{"basic check", plugin.Info{Name: "gcp", Installed: "yes", Version: "1.0.0"}, Entry{
			Name:              "v1",
			GcpZone:           "us-west1",
			ClusterCredential: "my creds",
		}},
	}
}
