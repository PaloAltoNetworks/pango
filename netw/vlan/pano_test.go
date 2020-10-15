package vlan

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestPanoNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.tmpl == "" {
				t.Skip()
			}
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.tmpl, "", tc.vsys, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.tmpl, "", tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
				if tc.tmpl != mc.Template {
					t.Errorf("template: %s != %s", tc.tmpl, mc.Template)
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
