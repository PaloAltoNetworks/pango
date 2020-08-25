package eth

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	version    version.Number
	vsys       string
	importVsys string
	imports    []string
	conf       Entry
}

func getTests() []testCase {
	return []testCase{
		{version.Number{5, 0, 0, ""}, "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
			Name:              "ethernet1/1",
			Mode:              "layer3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			Ipv6Enabled:       true,
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			NetflowProfile:    "some profile",
			LinkSpeed:         "auto",
			LinkDuplex:        "auto",
			LinkState:         "up",
			Comment:           "v1 basic l3",
		}},
		{version.Number{5, 0, 0, ""}, "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
			Name:                   "ethernet1/2",
			Mode:                   "layer3",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			Comment:                "v1 dhcp",
		}},
		{version.Number{5, 0, 0, ""}, "vsys4", "vsys4", []string{}, Entry{
			Name:    "ethernet1/3",
			Mode:    "ha",
			Comment: "v1 ha no import",
		}},
		{version.Number{5, 0, 0, ""}, "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
			Name: "ethernet1/4",
			Mode: "layer3",
			raw: map[string]string{
				"arp":            "<arp>raw arp</arp>",
				"v6adr":          "<address>raw ipv6 addresses</address>",
				"v6nd":           "ipv6 neighbor info",
				"l3subinterface": "<units>raw l3 subinterfaces</units>",
			},
			Comment: "v1 layer3 with raw config",
		}},
		{version.Number{5, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
			Name: "ethernet1/5",
			Mode: "layer2",
			raw: map[string]string{
				"l2subinterface": "<units>raw l2 subinterfaces</units>",
			},
			Comment: "v1 layer2 with raw config",
		}},
		{version.Number{5, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/6"}, Entry{
			Name: "ethernet1/6",
			Mode: "virtual-wire",
			raw: map[string]string{
				"vwsub": "virtual wire subinterfaces",
			},
		}},
		{version.Number{5, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/7"}, Entry{
			Name: "ethernet1/7",
			Mode: "tap",
		}},
		{version.Number{8, 0, 0, ""}, "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
			Name:              "ethernet1/1",
			Mode:              "layer3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			Ipv6Enabled:       true,
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     42,
			Ipv6MssAdjust:     84,
			NetflowProfile:    "some profile",
			LinkSpeed:         "auto",
			LinkDuplex:        "auto",
			LinkState:         "up",
			Comment:           "v2 basic l3",
		}},
		{version.Number{8, 0, 0, ""}, "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
			Name:                   "ethernet1/2",
			Mode:                   "layer3",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			Comment:                "v2 dhcp",
		}},
		{version.Number{8, 0, 0, ""}, "vsys4", "vsys4", []string{}, Entry{
			Name:    "ethernet1/3",
			Mode:    "ha",
			Comment: "v2 ha no import",
		}},
		{version.Number{8, 0, 0, ""}, "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
			Name: "ethernet1/4",
			Mode: "layer3",
			raw: map[string]string{
				"arp":            "<arp>raw arp</arp>",
				"v6adr":          "<address>raw ipv6 addresses</address>",
				"v6nd":           "ipv6 neighbor info",
				"pppoe":          "pppoe info",
				"ndp":            "ndp proxy info",
				"l3subinterface": "<units>raw l3 subinterfaces</units>",
			},
			Comment: "v2 layer3 with raw config",
		}},
		{version.Number{8, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
			Name: "ethernet1/5",
			Mode: "layer2",
			raw: map[string]string{
				"l2subinterface": "<units>raw l2 subinterfaces</units>",
			},
			Comment: "v2 layer2 with raw config",
		}},
		{version.Number{8, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/6"}, Entry{
			Name: "ethernet1/6",
			Mode: "virtual-wire",
			raw: map[string]string{
				"vwsub": "virtual wire subinterfaces",
			},
		}},
		{version.Number{8, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/7"}, Entry{
			Name: "ethernet1/7",
			Mode: "tap",
		}},
		{version.Number{8, 1, 0, ""}, "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
			Name:              "ethernet1/1",
			Mode:              "layer3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			Ipv6Enabled:       true,
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     42,
			Ipv6MssAdjust:     84,
			NetflowProfile:    "some profile",
			LinkSpeed:         "auto",
			LinkDuplex:        "auto",
			LinkState:         "up",
			DecryptForward:    true,
			RxPolicingRate:    88,
			TxPolicingRate:    99,
			Comment:           "v3 basic l3",
		}},
		{version.Number{8, 1, 0, ""}, "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
			Name:                   "ethernet1/2",
			Mode:                   "layer3",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			Comment:                "v3 dhcp",
		}},
		{version.Number{8, 1, 0, ""}, "vsys4", "vsys4", []string{}, Entry{
			Name:    "ethernet1/3",
			Mode:    "ha",
			Comment: "v3 ha no import",
		}},
		{version.Number{8, 1, 0, ""}, "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
			Name: "ethernet1/4",
			Mode: "layer3",
			raw: map[string]string{
				"arp":            "<arp>raw arp</arp>",
				"v6adr":          "<address>raw ipv6 addresses</address>",
				"v6nd":           "ipv6 neighbor info",
				"pppoe":          "pppoe info",
				"ndp":            "ndp proxy info",
				"l3subinterface": "<units>raw l3 subinterfaces</units>",
			},
			Comment: "v3 layer3 with raw config",
		}},
		{version.Number{8, 1, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
			Name: "ethernet1/5",
			Mode: "layer2",
			raw: map[string]string{
				"l2subinterface": "<units>raw l2 subinterfaces</units>",
			},
			Comment: "v3 layer2 with raw config",
		}},
		{version.Number{8, 1, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/6"}, Entry{
			Name: "ethernet1/6",
			Mode: "virtual-wire",
			raw: map[string]string{
				"vwsub": "virtual wire subinterfaces",
			},
		}},
		{version.Number{8, 1, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/7"}, Entry{
			Name: "ethernet1/7",
			Mode: "tap",
		}},
		{version.Number{9, 0, 0, ""}, "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
			Name:              "ethernet1/1",
			Mode:              "layer3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			Ipv6Enabled:       true,
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     42,
			Ipv6MssAdjust:     84,
			NetflowProfile:    "some profile",
			LinkSpeed:         "auto",
			LinkDuplex:        "auto",
			LinkState:         "up",
			DecryptForward:    true,
			RxPolicingRate:    88,
			TxPolicingRate:    99,
			Comment:           "v4 basic l3",
		}},
		{version.Number{9, 0, 0, ""}, "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
			Name:                   "ethernet1/2",
			Mode:                   "layer3",
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 12,
			DhcpSendHostnameEnable: true,
			DhcpSendHostnameValue:  "foo-hostname",
			Comment:                "v4 dhcp",
		}},
		{version.Number{9, 0, 0, ""}, "vsys4", "vsys4", []string{}, Entry{
			Name:    "ethernet1/3",
			Mode:    "ha",
			Comment: "v4 ha no import",
		}},
		{version.Number{9, 0, 0, ""}, "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
			Name: "ethernet1/4",
			Mode: "layer3",
			raw: map[string]string{
				"arp":            "<arp>raw arp</arp>",
				"v6adr":          "<address>raw ipv6 addresses</address>",
				"v6nd":           "ipv6 neighbor info",
				"pppoe":          "pppoe info",
				"ndp":            "ndp proxy info",
				"l3subinterface": "<units>raw l3 subinterfaces</units>",
				"v6client":       "ipv6 client info",
				"ddns":           "ddns config",
			},
			Comment: "v4 layer3 with raw config",
		}},
		{version.Number{9, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
			Name: "ethernet1/5",
			Mode: "layer2",
			raw: map[string]string{
				"l2subinterface": "<units>raw l2 subinterfaces</units>",
			},
			Comment: "v4 layer2 with raw config",
		}},
		{version.Number{9, 0, 0, ""}, "vsys7", "vsys7", []string{}, Entry{
			Name:           "ethernet1/7",
			Mode:           "aggregate-group",
			AggregateGroup: "ae1",
		}},
		{version.Number{9, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/6"}, Entry{
			Name: "ethernet1/6",
			Mode: "virtual-wire",
			raw: map[string]string{
				"vwsub": "virtual wire subinterfaces",
			},
		}},
		{version.Number{9, 0, 0, ""}, "vsys6", "vsys6", []string{"ethernet1/7"}, Entry{
			Name: "ethernet1/7",
			Mode: "tap",
		}},
	}
}
