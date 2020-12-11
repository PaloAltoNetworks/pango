package advertise

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
		{"v1 minimal", version.Number{7, 0, 0, ""}, Entry{
			Name:   "minimal",
			Enable: false,
		}},
		{"v1 full", version.Number{7, 0, 0, ""}, Entry{
			Name:                   "foo",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			AddressPrefix: map[string]bool{
				"5.5.0.0": true,
				"6.6.0.0": false,
				"7.7.0.0": true,
				"8.8.0.0": false,
			},
			NextHop:  []string{"nh1", "nh2"},
			FromPeer: []string{"fp1", "fp2"},
		}},
		{"v2 minimal", version.Number{8, 0, 0, ""}, Entry{
			Name:   "minimal",
			Enable: false,
		}},
		{"v2 full", version.Number{8, 0, 0, ""}, Entry{
			Name:                   "foo",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			RouteTable:             RouteTableBoth,
			AddressPrefix: map[string]bool{
				"5.5.0.0": true,
				"6.6.0.0": false,
				"7.7.0.0": true,
				"8.8.0.0": false,
			},
			NextHop:  []string{"nh1", "nh2"},
			FromPeer: []string{"fp1", "fp2"},
		}},
	}
}
