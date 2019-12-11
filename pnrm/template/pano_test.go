package template

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		version version.Number
		conf    Entry
	}{
		{"v1 test one", version.Number{6, 1, 0, ""}, Entry{
			Name:        "one",
			Description: "my description",
			raw: map[string]string{
				"conf": "conf",
			},
		}},
		{"v1 test two", version.Number{6, 1, 0, ""}, Entry{
			Name:        "two",
			Description: "with devices",
			MultiVsys:   true,
			Devices: map[string][]string{
				"001234": nil,
				"998765": {"vsys3", "vsys5"},
			},
			raw: map[string]string{
				"conf": "conf",
			},
		}},
		{"v1 test three", version.Number{6, 1, 0, ""}, Entry{
			Name:        "three",
			Description: "with one device",
			Mode:        "normal",
			Devices: map[string][]string{
				"001234": nil,
			},
			raw: map[string]string{
				"conf": "conf",
			},
		}},
		{"v1 test four", version.Number{6, 1, 0, ""}, Entry{
			Name:           "four",
			Description:    "with one device",
			Mode:           "normal",
			MultiVsys:      true,
			VpnDisableMode: true,
			Devices: map[string][]string{
				"001234": nil,
			},
			raw: map[string]string{
				"conf": "conf",
			},
		}},
		{"v3 raw test", version.Number{8, 1, 0, ""}, Entry{
			Name:        "v3 full",
			Description: "v3 desc",
			DefaultVsys: "vsys3",
			raw: map[string]string{
				"conf": "conf",
				"vars": "vars",
			},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &Template{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.Version = tc.version
			mc.AddResp("")
			err := ns.Set(tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
