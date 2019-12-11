package stack

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
			Description: "with one device",
			Devices:     []string{"001234"},
		}},
		{"v1 test two", version.Number{6, 1, 0, ""}, Entry{
			Name:        "two",
			Description: "with multi device",
			Templates:   []string{"tem1", "tem2"},
			Devices:     []string{"001234", "123"},
		}},
		{"v1 test three", version.Number{6, 1, 0, ""}, Entry{
			Name:        "three",
			Description: "with multi device",
			Templates:   []string{"tem1", "tem2"},
			DefaultVsys: "myvsys",
			raw: map[string]string{
				"var":  "this is the var data",
				"conf": "this is conf data",
			},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &Stack{}
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
