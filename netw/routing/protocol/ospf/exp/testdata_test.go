package exp

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
		{"v1 export rule", version.Number{10, 0, 0, ""}, Entry{
			Name:     "t1",
			PathType: "path type",
			Tag:      "tag",
			Metric:   100,
		}},
	}
}
