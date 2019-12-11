package dampening

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		conf Entry
	}{
		{"with enable", Entry{
			Name:                     "with enable",
			Enable:                   true,
			Cutoff:                   1.25,
			Reuse:                    0.5,
			MaxHoldTime:              900,
			DecayHalfLifeReachable:   300,
			DecayHalfLifeUnreachable: 900,
		}},
		{"without enable", Entry{
			Name:                     "without enable",
			Enable:                   false,
			Cutoff:                   2,
			Reuse:                    500,
			MaxHoldTime:              600,
			DecayHalfLifeReachable:   700,
			DecayHalfLifeUnreachable: 800,
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoDampening{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set("tmpl", "", "vr", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("tmpl", "", "vr", tc.conf.Name)
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
