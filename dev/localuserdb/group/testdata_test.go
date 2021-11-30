package group

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
		{"v1 empty group", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
		}},
		{"v1 group with users", version.Number{6, 1, 0, ""}, Entry{
			Name:  "two",
			Users: []string{"one", "two", "three"},
		}},
	}
}
