package redist

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
		{"v1 disable", version.Number{7, 0, 0, ""}, Entry{
			Name:                 "one",
			Enable:               false,
			Metric:               42,
			SetOrigin:            "origin",
			SetMed:               "med",
			SetLocalPreference:   "localpref",
			SetAsPathLimit:       1,
			SetCommunity:         []string{"com1", "com2"},
			SetExtendedCommunity: []string{"ec1", "ec2"},
		}},
		{"v1 enable", version.Number{7, 0, 0, ""}, Entry{
			Name:                 "two",
			Enable:               true,
			Metric:               42,
			SetOrigin:            "origin",
			SetMed:               "med",
			SetLocalPreference:   "localpref",
			SetAsPathLimit:       2,
			SetCommunity:         []string{"com1", "com2"},
			SetExtendedCommunity: []string{"ec1", "ec2"},
		}},
		{"v2 disable", version.Number{8, 0, 0, ""}, Entry{
			Name:                 "three",
			Enable:               false,
			AddressFamily:        AddressFamilyIpv4,
			Metric:               42,
			SetOrigin:            "origin",
			SetMed:               "med",
			SetLocalPreference:   "localpref",
			SetAsPathLimit:       3,
			SetCommunity:         []string{"com1", "com2"},
			SetExtendedCommunity: []string{"ec1", "ec2"},
		}},
		{"v2 enable", version.Number{8, 0, 0, ""}, Entry{
			Name:                 "four",
			Enable:               true,
			AddressFamily:        AddressFamilyIpv6,
			RouteTable:           RouteTableBoth,
			Metric:               42,
			SetOrigin:            "origin",
			SetMed:               "med",
			SetLocalPreference:   "localpref",
			SetAsPathLimit:       4,
			SetCommunity:         []string{"com1", "com2"},
			SetExtendedCommunity: []string{"ec1", "ec2"},
		}},
	}
}
