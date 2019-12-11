package tags

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		dg   string
		conf Entry
	}{
		{"test with all fields", "", Entry{
			Name:    "one",
			Color:   "color1",
			Comment: "first test",
		}},
		{"test no color", "", Entry{
			Name:    "two",
			Comment: "second test",
		}},
		{"test no comment", "dg1", Entry{
			Name:  "three",
			Color: "color3",
		}},
		{"test no color or comment", "dg2", Entry{
			Name: "four",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoTags{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.dg, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.dg, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
