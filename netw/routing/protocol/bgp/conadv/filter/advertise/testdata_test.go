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
		{"v1 basic", version.Number{7, 0, 0, ""}, Entry{
			Name:   "one",
			Enable: false,
		}},
		{"v1 full", version.Number{7, 0, 0, ""}, Entry{
			Name:                   "two",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			AddressPrefix:          []string{"5.5.0.0", "6.6.0.0"},
			NextHop:                []string{"nh1", "nh2"},
			FromPeer:               []string{"fp1", "fp2"},
		}},
		{"v2 basic", version.Number{8, 0, 0, ""}, Entry{
			Name:   "three",
			Enable: false,
		}},
		{"v2 full", version.Number{8, 0, 0, ""}, Entry{
			Name:                   "four",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			RouteTable:             RouteTableUnicast,
			AddressPrefix:          []string{"5.5.0.0", "6.6.0.0"},
			NextHop:                []string{"nh1", "nh2"},
			FromPeer:               []string{"fp1", "fp2"},
		}},
	}
}
