package loopback

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
		{version.Number{5, 0, 0, ""}, "one", "vsys2", "vsys2", []string{"loopback.1"}, Entry{
			Name:              "loopback.1",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			NetflowProfile:    "some profile",
			Comment:           "v1 basic",
		}},
		{version.Number{5, 0, 0, ""}, "two", "vsys3", "vsys3", []string{"loopback.2"}, Entry{
			Name:      "loopback.2",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"v6a": "ipv6 addresses",
			},
			Comment: "v1 raw no import",
		}},
		{version.Number{5, 0, 0, ""}, "", "", "", []string{"loopback.2"}, Entry{
			Name:      "loopback.2",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"v6a": "ipv6 addresses",
			},
			Comment: "v1 tmpl with raw",
		}},
		{version.Number{5, 0, 0, ""}, "", "", "", []string{"loopback.2"}, Entry{
			Name:            "loopback.2",
			StaticIps:       []string{"10.3.1.1/24"},
			EnableIpv6:      true,
			Ipv6InterfaceId: "someId",
			Comment:         "v1 with ipv6 config",
		}},
		{version.Number{8, 0, 0, ""}, "three", "vsys2", "vsys2", []string{"loopback.3"}, Entry{
			Name:              "loopback.3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     12,
			Ipv6MssAdjust:     14,
			NetflowProfile:    "some profile",
			Comment:           "v2 basic",
		}},
		{version.Number{8, 0, 0, ""}, "", "", "", []string{"loopback.4"}, Entry{
			Name:          "loopback.4",
			StaticIps:     []string{"10.3.1.1/24"},
			Ipv6MssAdjust: 1024,
			raw: map[string]string{
				"v6a": "ipv6 addresses",
			},
			Comment: "v2 raw no import",
		}},
		{version.Number{8, 0, 0, ""}, "four", "vsys3", "vsys3", []string{"loopback.4"}, Entry{
			Name:          "loopback.4",
			StaticIps:     []string{"10.3.1.1/24"},
			Ipv6MssAdjust: 1024,
			raw: map[string]string{
				"v6a": "ipv6 addresses",
			},
			Comment: "v2 tmpl with raw",
		}},
		{version.Number{8, 0, 0, ""}, "", "", "", []string{"loopback.2"}, Entry{
			Name:            "loopback.2",
			StaticIps:       []string{"10.3.1.1/24"},
			EnableIpv6:      true,
			Ipv6InterfaceId: "someId",
			Comment:         "v2 with ipv6 config",
		}},
	}
}
