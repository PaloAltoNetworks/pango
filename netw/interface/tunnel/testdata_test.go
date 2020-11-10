package tunnel

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
		{version.Number{6, 1, 0, ""}, "one", "vsys2", "vsys2", []string{"tunnel.1"}, Entry{
			Name:              "tunnel.1",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			NetflowProfile:    "some profile",
			Comment:           "v1 basic",
		}},
		{version.Number{6, 1, 0, ""}, "two", "vsys3", "vsys3", []string{"tunnel.2"}, Entry{
			Name:      "tunnel.2",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"v6a": "ipv6 addresses",
			},
			Comment: "v1 raw no import",
		}},
		{version.Number{6, 1, 0, ""}, "two", "vsys3", "vsys3", []string{"tunnel.3"}, Entry{
			Name:            "tunnel.3",
			StaticIps:       []string{"10.3.1.1/24"},
			EnableIpv6:      true,
			Ipv6InterfaceId: "myId",
			Comment:         "v1 no import with ipv6 config",
		}},
	}
}
