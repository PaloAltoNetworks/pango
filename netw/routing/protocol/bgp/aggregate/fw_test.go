package aggregate

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		conf Entry
	}{
		{"minimal", Entry{
			Name:   "one",
			Enable: true,
			Prefix: "prefix",
			raw: map[string]string{
				"sf": "suppress filters",
				"af": "advertise filters",
			},
		}},
		{"community none", Entry{
			Name:                  "two",
			Prefix:                "prefix",
			AsSet:                 false,
			Enable:                false,
			Summary:               true,
			LocalPreference:       "local pref",
			Med:                   "med",
			Weight:                5,
			NextHop:               "nexthop",
			Origin:                "origin",
			AsPathLimit:           6,
			AsPathType:            AsPathTypeNone,
			CommunityType:         CommunityTypeNone,
			ExtendedCommunityType: CommunityTypeRemoveAll,
			raw: map[string]string{
				"af": "advertise filters",
			},
		}},
		{"community removeall", Entry{
			Name:                  "three",
			Prefix:                "prefix",
			AsSet:                 true,
			Enable:                false,
			Summary:               false,
			LocalPreference:       "local pref",
			Med:                   "med",
			Weight:                5,
			NextHop:               "nexthop",
			Origin:                "origin",
			AsPathLimit:           6,
			AsPathType:            AsPathTypePrepend,
			AsPathValue:           "1",
			CommunityType:         CommunityTypeRemoveAll,
			ExtendedCommunityType: CommunityTypeNone,
			raw: map[string]string{
				"sf": "suppress filters",
			},
		}},
		{"community remove regex", Entry{
			Name:                   "four",
			Prefix:                 "prefix",
			AsSet:                  true,
			Enable:                 true,
			Summary:                true,
			LocalPreference:        "local pref",
			Med:                    "med",
			Weight:                 5,
			NextHop:                "nexthop",
			Origin:                 "origin",
			AsPathLimit:            6,
			AsPathType:             AsPathTypeNone,
			CommunityType:          CommunityTypeRemoveRegex,
			CommunityValue:         "remove regex",
			ExtendedCommunityType:  CommunityTypeAppend,
			ExtendedCommunityValue: "append",
		}},
		{"community append", Entry{
			Name:                   "five",
			Prefix:                 "prefix",
			AsSet:                  false,
			Enable:                 false,
			Summary:                false,
			LocalPreference:        "local pref",
			Med:                    "med",
			Weight:                 5,
			NextHop:                "nexthop",
			Origin:                 "origin",
			AsPathLimit:            6,
			AsPathType:             AsPathTypeNone,
			CommunityType:          CommunityTypeAppend,
			CommunityValue:         "append",
			ExtendedCommunityType:  CommunityTypeOverwrite,
			ExtendedCommunityValue: "overwrite",
		}},
		{"community overwrite", Entry{
			Name:                   "six",
			Prefix:                 "prefix",
			AsSet:                  false,
			Enable:                 false,
			Summary:                false,
			LocalPreference:        "local pref",
			Med:                    "med",
			Weight:                 5,
			NextHop:                "nexthop",
			Origin:                 "origin",
			AsPathLimit:            6,
			AsPathType:             AsPathTypeNone,
			CommunityType:          CommunityTypeOverwrite,
			CommunityValue:         "Overwrite",
			ExtendedCommunityType:  CommunityTypeRemoveRegex,
			ExtendedCommunityValue: "remove regex",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwAggregate{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set("vr", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("vr", tc.conf.Name)
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
