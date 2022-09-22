package kerberos

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 no servers", version.Number{7, 0, 0, ""}, Entry{
			Name:         "basicNoServers",
			AdminUseOnly: true,
		}},
		{"v1 with one server", version.Number{7, 0, 0, ""}, Entry{
			Name:         "withServer",
			AdminUseOnly: false,
			Servers: []Server{
				{
					Name:   "first",
					Server: "first.example.com",
					Port:   88,
				},
			},
		}},
		{"v1 with multiple server", version.Number{7, 0, 0, ""}, Entry{
			Name:         "withServers",
			AdminUseOnly: true,
			Servers: []Server{
				{
					Name:   "first",
					Server: "first.example.com",
					Port:   88,
				},
				{
					Name:   "second",
					Server: "second.example.com",
					Port:   1234,
				},
			},
		}},
	}
}
