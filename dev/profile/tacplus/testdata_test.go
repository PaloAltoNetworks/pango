package tacplus

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
			Name:                "basicNoServers",
			Timeout:             2,
			AdminUseOnly:        false,
			UseSingleConnection: true,
		}},
		{"v1 with servers", version.Number{7, 0, 0, ""}, Entry{
			Name:                "basicWithServers",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: false,
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   49,
				},
			},
		}},
		{"v2 no servers", version.Number{8, 0, 0, ""}, Entry{
			Name:                "basicNoServers",
			Timeout:             2,
			AdminUseOnly:        false,
			UseSingleConnection: true,
		}},
		{"v2 with servers", version.Number{8, 0, 0, ""}, Entry{
			Name:                "basicWithServers",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: false,
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   49,
				},
			},
		}},
		{"v2 chap", version.Number{8, 0, 0, ""}, Entry{
			Name:                "v2Chap",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: true,
			Protocol: Protocol{
				Chap: true,
			},
		}},
		{"v2 pap", version.Number{8, 0, 0, ""}, Entry{
			Name:                "v2Pap",
			Timeout:             3,
			AdminUseOnly:        false,
			UseSingleConnection: false,
			Protocol: Protocol{
				Pap: true,
			},
		}},
		{"v2 auto", version.Number{8, 0, 0, ""}, Entry{
			Name:                "v2Auto",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: false,
			Protocol: Protocol{
				Auto: true,
			},
		}},
		{"v3 no servers", version.Number{8, 1, 0, ""}, Entry{
			Name:                "basicNoServers",
			Timeout:             2,
			AdminUseOnly:        false,
			UseSingleConnection: true,
		}},
		{"v3 with servers", version.Number{8, 1, 0, ""}, Entry{
			Name:                "basicWithServers",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: false,
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   49,
				},
			},
		}},
		{"v3 chap", version.Number{8, 1, 0, ""}, Entry{
			Name:                "v3Chap",
			Timeout:             3,
			AdminUseOnly:        true,
			UseSingleConnection: true,
			Protocol: Protocol{
				Chap: true,
			},
		}},
		{"v3 pap", version.Number{8, 1, 0, ""}, Entry{
			Name:                "v3Pap",
			Timeout:             3,
			AdminUseOnly:        false,
			UseSingleConnection: false,
			Protocol: Protocol{
				Pap: true,
			},
		}},
	}
}
