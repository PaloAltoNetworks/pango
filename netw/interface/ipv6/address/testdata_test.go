package address

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
		{"v1 without neighbor config", version.Number{6, 1, 0, ""}, Entry{
			Name:              "::1",
			Enabled:           true,
			InterfaceIdAsHost: false,
			Anycast:           true,
		}},
		{"v1 with neighbor config pt1", version.Number{6, 1, 0, ""}, Entry{
			Name:                "::1",
			Enabled:             false,
			InterfaceIdAsHost:   true,
			Anycast:             false,
			EnableRa:            true,
			RaValidLifetime:     "1234",
			RaPreferredLifetime: "infinity",
			RaOnLink:            true,
			RaAutonomous:        false,
		}},
		{"v1 with neighbor config pt2", version.Number{6, 1, 0, ""}, Entry{
			Name:                "::1",
			Enabled:             false,
			InterfaceIdAsHost:   true,
			Anycast:             false,
			EnableRa:            false,
			RaValidLifetime:     "infinity",
			RaPreferredLifetime: "4321",
			RaOnLink:            false,
			RaAutonomous:        true,
		}},
	}
}
