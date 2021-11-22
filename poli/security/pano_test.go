package security

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			if tc.base == "" {
				t.Skip()
			}
			mc.AddResp("")
			err := ns.Set(tc.vsys, tc.base, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.vsys, tc.base, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}

func TestPanoLogEndMissingIsTrue(t *testing.T) {
	versions := []version.Number{
		{6, 1, 0, ""},
		{9, 0, 0, ""},
		{10, 0, 0, ""},
	}

	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	mc.AddResp(`
<entry name="foo">
    <description>Hello world</description>
</entry>
`)
	for _, v := range versions {
		t.Run(fmt.Sprintf("version %s", v), func(t *testing.T) {
			mc.Version = v
			r, err := ns.Get("dg", util.PreRulebase, "foo")
			if err != nil {
				t.Errorf("Error in get: %s", err)
			} else if r.Description != "Hello world" {
				t.Errorf("Description is wrong: %s", r.Description)
			} else if !r.LogEnd {
				t.Errorf("LogEnd is not enabled")
			}
		})
	}
}
