package vlan

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.vsys, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
				if tc.vsys != mc.Vsys {
					t.Errorf("vsys: %s != %s", tc.vsys, mc.Vsys)
				}
				if !reflect.DeepEqual(tc.imports, mc.Imports) {
					t.Errorf("imports: %#v != %#v", tc.imports, mc.Imports)
				}
			}
		})
	}
}
