package iface

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
		{"v1 area interface broadcast", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Enable:             true,
			Passive:            true,
			LinkType:           LinkTypeBroadcast,
			Metric:             10,
			Priority:           1,
			HelloInterval:      10,
			DeadCounts:         5,
			RetransmitInterval: 10,
			TransitDelay:       1,
			GraceRestartDelay:  1,
			AuthProfile:        "auth profile",
			BfdProfile:         "bfd profile",
		}},
		{"v1 area interface p2p", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t2",
			Enable:             true,
			Passive:            true,
			LinkType:           LinkTypePointToPoint,
			Metric:             10,
			Priority:           1,
			HelloInterval:      10,
			DeadCounts:         5,
			RetransmitInterval: 10,
			TransitDelay:       1,
			GraceRestartDelay:  1,
			AuthProfile:        "auth profile",
			BfdProfile:         "bfd profile",
		}},
		{"v1 area interface p2mp", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t3",
			Enable:             true,
			Passive:            true,
			LinkType:           LinkTypePointToMultiPoint,
			Metric:             10,
			Priority:           1,
			HelloInterval:      10,
			DeadCounts:         5,
			RetransmitInterval: 10,
			TransitDelay:       1,
			GraceRestartDelay:  1,
			AuthProfile:        "auth profile",
			Neighbors: []string{
				"10.1.1.1",
				"192.1.1.1",
			},
			BfdProfile: "bfd profile",
		}},
	}
}
