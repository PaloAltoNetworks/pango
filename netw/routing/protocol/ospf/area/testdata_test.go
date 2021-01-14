package area

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
		{"v1 area normal", version.Number{10, 0, 0, ""}, Entry{
			Name: "t1",
			Type: TypeNormal,
		}},
		{"v1 area normal ranges", version.Number{10, 0, 0, ""}, Entry{
			Name: "t2",
			Type: TypeNormal,
			Ranges: []Range{
				Range{
					Network: "10.0.1.0/24",
					Action:  ActionAdvertise,
				},
				Range{
					Network: "10.0.2.0/24",
					Action:  ActionSuppress,
				},
			},
		}},
		{"v1 area stub", version.Number{10, 0, 0, ""}, Entry{
			Name:                  "t3",
			Type:                  TypeStub,
			AcceptSummary:         true,
			DefaultRouteAdvertise: true,
			AdvertiseMetric:       10,
		}},
		{"v1 area nssa", version.Number{10, 0, 0, ""}, Entry{
			Name:                  "t4",
			Type:                  TypeNssa,
			AcceptSummary:         true,
			DefaultRouteAdvertise: true,
			AdvertiseMetric:       10,
			AdvertiseType:         "advertise type",
		}},
		{"v1 area nssa ext ranges", version.Number{10, 0, 0, ""}, Entry{
			Name:                  "t5",
			Type:                  TypeNssa,
			AcceptSummary:         true,
			DefaultRouteAdvertise: true,
			AdvertiseMetric:       10,
			AdvertiseType:         "advertise type",
			ExtRanges: []Range{
				Range{
					Network: "10.0.1.0/24",
					Action:  ActionAdvertise,
				},
				Range{
					Network: "10.0.2.0/24",
					Action:  ActionSuppress,
				},
			},
		}},
	}
}
