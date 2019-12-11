package ipv4

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		vr   string
		conf Entry
	}{
		{"barebones redist", "v1", Entry{
			Name:     "one",
			Priority: 1,
			Action:   ActionRedist,
		}},
		{"barebones no redist", "v1", Entry{
			Name:     "two",
			Priority: 2,
			Action:   ActionNoRedist,
		}},
		{"redist static", "v1", Entry{
			Name:     "three",
			Priority: 3,
			Action:   ActionRedist,
			Types:    []string{TypeStatic},
		}},
		{"redist rip and connect", "v1", Entry{
			Name:     "four",
			Priority: 4,
			Action:   ActionRedist,
			Types:    []string{TypeRip, TypeConnect},
		}},
		{"no redist with ospf", "v1", Entry{
			Name:          "five",
			Priority:      5,
			Action:        ActionNoRedist,
			Types:         []string{TypeOspf},
			OspfPathTypes: []string{OspfPathTypeExt1, OspfPathTypeExt2},
			OspfAreas:     []string{"10.1.7.1", "10.1.7.2"},
			OspfTags:      []string{"10.1.7.3", "10.1.7.4"},
		}},
		{"no redist with bgp", "v1", Entry{
			Name:                   "six",
			Priority:               6,
			Action:                 ActionNoRedist,
			Types:                  []string{TypeBgp},
			BgpCommunities:         []string{"com1", "com2"},
			BgpExtendedCommunities: []string{"excom1", "excom2"},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwIpv4{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.vr, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.vr, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
