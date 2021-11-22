package security

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestFwNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.vsys, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.vsys, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}

func TestLogEndMissingIsTrue(t *testing.T) {
	versions := []version.Number{
		{6, 1, 0, ""},
		{9, 0, 0, ""},
		{10, 0, 0, ""},
	}

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	mc.AddResp(`
<entry name="foo">
    <description>Hello world</description>
</entry>
`)
	for _, v := range versions {
		t.Run(fmt.Sprintf("version %s", v), func(t *testing.T) {
			mc.Version = v
			r, err := ns.Get("vsys1", "foo")
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
