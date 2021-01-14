package vlink

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
		{"v1 area virtual link", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Enable:             true,
			NeighborId:         "neighbor id",
			TransitAreaId:      "transit area id",
			HelloInterval:      2600,
			DeadCounts:         20,
			RetransmitInterval: 2600,
			TransitDelay:       2600,
			AuthProfile:        "auth profile",
			BfdProfile:         "bfd profile",
		}},
	}
}
