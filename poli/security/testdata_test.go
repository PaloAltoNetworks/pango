package security

import (
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	version    version.Number
	desc       string
	vsys       string
	base       string
	doDefaults bool
	conf       Entry
}

func getTests() []testCase {
	return []testCase{
		{version.Number{6, 1, 0, ""}, "basic rule", "", "", true, Entry{
			Name: "rule1",
		}},
		{version.Number{6, 1, 0, ""}, "prerulebase rule", "vsys2", util.PreRulebase, true, Entry{
			Name:     "rule2",
			LogEnd:   true,
			Disabled: true,
		}},
		{version.Number{6, 1, 0, ""}, "postrulebase rule with target", "vsys3", util.PostRulebase, true, Entry{
			Name: "rule3",
			Targets: map[string][]string{
				"fw1": nil,
				"fw2": {"vsys2", "vsys3"},
			},
			NegateTarget: true,
		}},
	}
}
