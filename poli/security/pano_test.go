package security

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/util"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc       string
		vsys       string
		base       string
		doDefaults bool
		conf       Entry
	}{
		{"basic rule", "", "", true, Entry{
			Name: "rule1",
		}},
		{"prerulebase rule", "vsys2", util.PreRulebase, true, Entry{
			Name:     "rule2",
			LogEnd:   true,
			Disabled: true,
		}},
		{"postrulebase rule with target", "vsys3", util.PostRulebase, true, Entry{
			Name: "rule3",
			Targets: map[string][]string{
				"fw1": nil,
				"fw2": {"vsys2", "vsys3"},
			},
			NegateTarget: true,
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoSecurity{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			if tc.doDefaults {
				tc.conf.Defaults()
			}
			err := ns.Set(tc.vsys, tc.base, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.vsys, tc.base, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
