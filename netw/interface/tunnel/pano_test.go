package tunnel

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
		t.Run(tc.conf.Comment, func(t *testing.T) {
			var err error
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err = ns.Set(tc.tmpl, "", tc.vsys, tc.conf)
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
				if mc.Template != tc.tmpl {
					t.Errorf("template: %q != %q", tc.tmpl, mc.Template)
				}
				if tc.importVsys != mc.Vsys {
					t.Errorf("vsys: %q != %q", tc.importVsys, mc.Vsys)
				}
				if !reflect.DeepEqual(tc.imports, mc.Imports) {
					t.Errorf("imports: %#v != %#v", tc.imports, mc.Imports)
				}
			}
		})
	}
}
