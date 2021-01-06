package link

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
		{"v1", version.Number{10, 0, 0, ""}, Entry{
			Name:             "v1",
			Enable:           true,
			FailureCondition: FailureConditionAll,
			Interfaces: []string{
				"interface 1",
				"interface 2",
			},
		}},
	}
}
