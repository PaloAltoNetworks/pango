package ipv6

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
		{"v1 route no nexthop", version.Number{5, 0, 0, ""}, Entry{
			Name:          "one",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			BfdProfile:    "myBfdProfile",
			AdminDistance: 101,
			Metric:        111,
			RouteTable:    RouteTableNoInstall,
		}},
		{"v1 route discard nexthop", version.Number{5, 0, 0, ""}, Entry{
			Name:          "two",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			Type:          NextHopDiscard,
			AdminDistance: 101,
		}},
		{"v1 route ip address nexthop", version.Number{5, 0, 0, ""}, Entry{
			Name:        "three",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			BfdProfile:  "myBfdProfile",
			Type:        NextHopIpv6Address,
			NextHop:     "ipv6 address",
			Metric:      111,
		}},
		{"v1 route next vr nexthop", version.Number{5, 0, 0, ""}, Entry{
			Name:        "four",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			Type:        NextHopNextVr,
			NextHop:     "default",
		}},
		{"v2 route no nexthop", version.Number{8, 0, 0, ""}, Entry{
			Name:          "one",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			BfdProfile:    "myBfdProfile",
			AdminDistance: 101,
			Metric:        111,
			RouteTable:    RouteTableNoInstall,
		}},
		{"v2 route discard nexthop", version.Number{8, 0, 0, ""}, Entry{
			Name:          "two",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			Type:          NextHopDiscard,
			AdminDistance: 101,
			RouteTable:    RouteTableUnicast,
		}},
		{"v2 route ip address nexthop", version.Number{8, 0, 0, ""}, Entry{
			Name:        "three",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			BfdProfile:  "myBfdProfile",
			Type:        NextHopIpv6Address,
			NextHop:     "ipv6 address",
			Metric:      111,
		}},
		{"v2 route next vr nexthop", version.Number{8, 0, 0, ""}, Entry{
			Name:        "four",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			Type:        NextHopNextVr,
			NextHop:     "default",
		}},
		{"v2 next vr with monitor", version.Number{8, 0, 0, ""}, Entry{
			Name:               "five",
			Destination:        "0.0.0.0/0",
			BfdProfile:         "myBfdProfile",
			Type:               NextHopNextVr,
			NextHop:            "default",
			EnablePathMonitor:  true,
			PmFailureCondition: "any",
			PmHoldTime:         5,
		}},
		{"v2 discard with dest monitor", version.Number{8, 0, 0, ""}, Entry{
			Name:        "six",
			Destination: "0.0.0.0/0",
			Type:        NextHopIpv6Address,
			NextHop:     "ipv6 address",
			MonitorDestinations: []MonitorDestination{
				MonitorDestination{
					Name:          "first",
					Enable:        true,
					SourceIp:      "1.2.3.4",
					DestinationIp: "5.6.7.8",
					PingInterval:  5,
					PingCount:     6,
				},
				MonitorDestination{
					Name:          "second",
					SourceIp:      "10.1.1.1",
					DestinationIp: "10.2.2.2",
				},
			},
		}},
		{"v3 route no nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:          "one",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			BfdProfile:    "myBfdProfile",
			AdminDistance: 101,
			Metric:        111,
			RouteTable:    RouteTableNoInstall,
		}},
		{"v3 route discard nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:          "two",
			Destination:   "0.0.0.0/0",
			Interface:     "ethernet1/1",
			Type:          NextHopDiscard,
			AdminDistance: 101,
			RouteTable:    RouteTableUnicast,
		}},
		{"v3 route ip address nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:        "three",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			BfdProfile:  "myBfdProfile",
			Type:        NextHopIpv6Address,
			NextHop:     "ipv6 address",
			Metric:      111,
		}},
		{"v3 route next vr nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:        "four",
			Destination: "0.0.0.0/0",
			Interface:   "ethernet1/1",
			Type:        NextHopNextVr,
			NextHop:     "default",
		}},
		{"v3 next vr with monitor", version.Number{9, 0, 0, ""}, Entry{
			Name:               "five",
			Destination:        "0.0.0.0/0",
			BfdProfile:         "myBfdProfile",
			Type:               NextHopNextVr,
			NextHop:            "default",
			EnablePathMonitor:  true,
			PmFailureCondition: "any",
			PmHoldTime:         5,
		}},
		{"v3 discard with dest monitor", version.Number{9, 0, 0, ""}, Entry{
			Name:        "six",
			Destination: "0.0.0.0/0",
			Type:        NextHopIpv6Address,
			NextHop:     "ipv6 address",
			MonitorDestinations: []MonitorDestination{
				MonitorDestination{
					Name:          "first",
					Enable:        true,
					SourceIp:      "1.2.3.4",
					DestinationIp: "5.6.7.8",
					PingInterval:  5,
					PingCount:     6,
				},
				MonitorDestination{
					Name:          "second",
					SourceIp:      "10.1.1.1",
					DestinationIp: "10.2.2.2",
				},
			},
		}},
		{"v3 fqdn", version.Number{9, 0, 0, ""}, Entry{
			Name:       "seven",
			Type:       NextHopFqdn,
			NextHop:    "example.com",
			BfdProfile: "myBfdProfile",
		}},
	}
}
