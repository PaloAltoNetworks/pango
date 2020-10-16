package dug

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	vsys    string
	dg      string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 single with no tags", "", "", version.Number{9, 1, 0, ""}, Entry{
			Name:        "one",
			Description: "my description",
			Filter:      "'single'",
		}},
		{"v1 single with tags", "", "", version.Number{9, 1, 0, ""}, Entry{
			Name:        "two",
			Description: "my second description",
			Filter:      "'single' and 'foo'",
			Tags:        []string{"tag2", "tag1"},
		}},
	}
}
