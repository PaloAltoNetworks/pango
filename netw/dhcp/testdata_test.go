package dhcp

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 relay ipv4", version.Number{10, 0, 0, ""}, Entry{
			Name: "t1",
			Relay: &Relay{
				Ipv4Enabled: true,
				Ipv4Servers: []string{
					"server 1",
					"server 2",
				},
			},
		}},
		{"v1 relay ipv6", version.Number{10, 0, 0, ""}, Entry{
			Name: "t2",
			Relay: &Relay{
				Ipv6Enabled: true,
				Ipv6Servers: []Ipv6Server{
					Ipv6Server{
						Server:    "server 1",
						Interface: "interface 1",
					},
					Ipv6Server{
						Server: "server 2",
					},
				},
			},
		}},
		{"v1 relay ipv4+ipv6", version.Number{10, 0, 0, ""}, Entry{
			Name: "t3",
			Relay: &Relay{
				Ipv4Enabled: true,
				Ipv4Servers: []string{
					"server 1",
					"server 2",
				},
				Ipv6Enabled: true,
				Ipv6Servers: []Ipv6Server{
					Ipv6Server{
						Server:    "server 1",
						Interface: "interface 1",
					},
					Ipv6Server{
						Server: "server 2",
					},
				},
			},
		}},
		{"v1 relay ipv4+ipv6 with raw", version.Number{10, 0, 0, ""}, Entry{
			Name: "t4",
			Relay: &Relay{
				Ipv4Enabled: true,
				Ipv4Servers: []string{
					"server 1",
					"server 2",
				},
				Ipv6Enabled: true,
				Ipv6Servers: []Ipv6Server{
					Ipv6Server{
						Server:    "server 1",
						Interface: "interface 1",
					},
					Ipv6Server{
						Server: "server 2",
					},
				},
			},
			raw: map[string]string{
				"server": "dhcp server",
			},
		}},
	}
}
