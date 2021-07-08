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
		{version.Number{9, 0, 0, ""}, "basic rule", "", "", true, Entry{
			Name: "rule1",
		}},
		{version.Number{9, 0, 0, ""}, "rule with uuid", "", "", true, Entry{
			Name: "rule1",
			Uuid: "123-456-78",
		}},
		{version.Number{9, 0, 0, ""}, "prerulebase rule", "vsys2", util.PreRulebase, true, Entry{
			Name:     "rule2",
			LogEnd:   true,
			Disabled: true,
		}},
		{version.Number{9, 0, 0, ""}, "postrulebase rule with target", "vsys3", util.PostRulebase, true, Entry{
			Name: "rule3",
			Targets: map[string][]string{
				"fw1": nil,
				"fw2": {"vsys2", "vsys3"},
			},
			NegateTarget: true,
		}},
		{version.Number{10, 0, 0, ""}, "basic rule", "", "", true, Entry{
			Name: "rule1",
		}},
		{version.Number{10, 0, 0, ""}, "rule with uuid", "", "", true, Entry{
			Name: "rule1",
			Uuid: "123-456-78",
		}},
		{version.Number{10, 0, 0, ""}, "rule with source and dest devices", "", "", true, Entry{
			Name:               "rule1",
			SourceDevices:      []string{"src1", "wu tang"},
			DestinationDevices: []string{"dest1", "clan"},
		}},
		{version.Number{10, 0, 0, ""}, "prerulebase rule", "vsys2", util.PreRulebase, true, Entry{
			Name:     "rule2",
			LogEnd:   true,
			Disabled: true,
		}},
		{version.Number{10, 0, 0, ""}, "postrulebase rule with target", "vsys3", util.PostRulebase, true, Entry{
			Name: "rule3",
			Targets: map[string][]string{
				"fw1": nil,
				"fw2": {"vsys2", "vsys3"},
			},
			NegateTarget: true,
		}},
	}
}
