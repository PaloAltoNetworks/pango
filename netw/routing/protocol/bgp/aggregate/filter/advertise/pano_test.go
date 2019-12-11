package advertise

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		version version.Number
		conf    Entry
	}{
		{"v1 minimal", version.Number{7, 0, 0, ""}, Entry{
			Name:   "minimal",
			Enable: false,
		}},
		{"v1 full", version.Number{7, 0, 0, ""}, Entry{
			Name:                   "foo",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			AddressPrefix: map[string]bool{
				"5.5.0.0": true,
				"6.6.0.0": false,
				"7.7.0.0": true,
				"8.8.0.0": false,
			},
			NextHop:  []string{"nh1", "nh2"},
			FromPeer: []string{"fp1", "fp2"},
		}},
		{"v2 minimal", version.Number{8, 0, 0, ""}, Entry{
			Name:   "minimal",
			Enable: false,
		}},
		{"v2 full", version.Number{8, 0, 0, ""}, Entry{
			Name:                   "foo",
			Enable:                 true,
			AsPathRegex:            "as path regex",
			CommunityRegex:         "community regex",
			ExtendedCommunityRegex: "extended community regex",
			Med:                    "match med",
			RouteTable:             RouteTableBoth,
			AddressPrefix: map[string]bool{
				"5.5.0.0": true,
				"6.6.0.0": false,
				"7.7.0.0": true,
				"8.8.0.0": false,
			},
			NextHop:  []string{"nh1", "nh2"},
			FromPeer: []string{"fp1", "fp2"},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoAdvertise{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.Version = tc.version
			mc.AddResp("")
			err := ns.Set("tmpl", "", "vr", "ag", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("tmpl", "", "vr", "ag", tc.conf.Name)
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
