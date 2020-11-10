package vlan

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	version    version.Number
	tmpl       string
	vsys       string
	importVsys string
	imports    []string
	conf       Entry
}

func getTests() []testCase {
	return []testCase{
		{version.Number{5, 0, 0, ""}, "tmpl", "vsys2", "vsys2", []string{"vlan.1"}, Entry{
			Name:              "vlan.1",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			NetflowProfile:    "some profile",
			Comment:           "v1 basic",
		}},
		{version.Number{5, 0, 0, ""}, "tmpl", "vsys3", "vsys3", []string{"vlan.2"}, Entry{
			Name:                   "vlan.2",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			Comment:                "v1 dhcp",
		}},
		{version.Number{5, 0, 0, ""}, "", "", "", []string{"vlan.3"}, Entry{
			Name:       "vlan.3",
			EnableDhcp: true,
			raw: map[string]string{
				"arp": "<arp>raw arp</arp>",
				"ndp": "<ndp-proxy>raw ndp</ndp-proxy>",
			},
			Comment: "v1 raw no import",
		}},
		{version.Number{5, 0, 0, ""}, "", "", "", []string{"vlan.4"}, Entry{
			Name:       "vlan.4",
			EnableDhcp: true,
			raw: map[string]string{
				"v6a": "ipv6 addresses",
				"v6n": "ipv6 neighbor discovery config",
			},
			Comment: "v1 ipv6 raw no import",
		}},
		{version.Number{5, 0, 0, ""}, "", "", "", []string{"vlan.5"}, Entry{
			Name:            "vlan.5",
			EnableIpv6:      true,
			Ipv6InterfaceId: "myId",
			Comment:         "v1 ipv6 only",
		}},
		{version.Number{8, 0, 0, ""}, "tmpl", "vsys2", "vsys2", []string{"vlan.4"}, Entry{
			Name:              "vlan.4",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     12,
			Ipv6MssAdjust:     14,
			NetflowProfile:    "some profile",
			Comment:           "v2 basic",
		}},
		{version.Number{8, 0, 0, ""}, "tmpl", "vsys3", "vsys3", []string{"vlan.5"}, Entry{
			Name:                   "vlan.5",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			Comment:                "v2 dhcp",
		}},
		{version.Number{8, 0, 0, ""}, "", "", "", []string{"vlan.6"}, Entry{
			Name:       "vlan.6",
			EnableDhcp: true,
			raw: map[string]string{
				"arp": "<arp>raw arp</arp>",
				"ndp": "<ndp-proxy>raw ndp</ndp-proxy>",
			},
			Comment: "v2 raw no import",
		}},
		{version.Number{8, 0, 0, ""}, "", "", "", []string{"vlan.4"}, Entry{
			Name:       "vlan.4",
			EnableDhcp: true,
			raw: map[string]string{
				"v6a": "ipv6 addresses",
				"v6n": "ipv6 neighbor discovery config",
			},
			Comment: "v2 ipv6 raw no import",
		}},
		{version.Number{8, 0, 0, ""}, "", "", "", []string{"vlan.5"}, Entry{
			Name:            "vlan.5",
			EnableIpv6:      true,
			Ipv6InterfaceId: "myId",
			Comment:         "v2 ipv6 only",
		}},
	}
}
