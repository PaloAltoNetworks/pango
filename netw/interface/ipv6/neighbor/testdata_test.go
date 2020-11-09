package neighbor

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	return []testCase{
		{"v1 no route advertise", version.Number{8, 0, 0, ""}, Config{
			EnableNdpMonitor:                  true,
			DuplicateAddressDetectionAttempts: 7,
			NeighborSolicitationInterval:      8,
			ReachableTime:                     9,
		}},
		{"v1 no route advertise with neighbors", version.Number{8, 0, 0, ""}, Config{
			EnableDuplicateAddressDetection:   true,
			DuplicateAddressDetectionAttempts: 7,
			NeighborSolicitationInterval:      8,
			ReachableTime:                     9,
			Neighbors: []Neighbor{
				Neighbor{
					Name:       "first",
					MacAddress: "00:30:48:52:ab:bc",
				},
				Neighbor{
					Name:       "second",
					MacAddress: "00:30:48:52:ab:cd",
				},
			},
		}},
		{"v1 with ra", version.Number{8, 0, 0, ""}, Config{
			EnableRa:                 true,
			RaMaxInterval:            600,
			RaMinInterval:            200,
			RaManagedFlag:            true,
			RaOtherFlag:              false,
			RaLinkMtu:                "9216",
			RaReachableTime:          "0",
			RaRetransmissionTimer:    "1",
			RaHopLimit:               "2",
			RaLifetime:               2,
			RaRouterPreference:       RaRouterPreferenceHigh,
			RaEnableConsistencyCheck: true,
		}},
		{"v1 with dns servers", version.Number{8, 0, 0, ""}, Config{
			RaEnableDnsSupport: true,
			RaDnsServers: []RaDnsServer{
				RaDnsServer{
					Name:     "one",
					Lifetime: 4,
				},
				RaDnsServer{
					Name:     "two",
					Lifetime: 3600,
				},
			},
		}},
		{"v1 with dns servers", version.Number{8, 0, 0, ""}, Config{
			RaDnsSuffixes: []RaDnsSuffix{
				RaDnsSuffix{
					Name:     "alpha",
					Lifetime: 4,
				},
				RaDnsSuffix{
					Name:     "beta",
					Lifetime: 3600,
				},
			},
		}},
	}
}
