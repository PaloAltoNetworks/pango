package iface

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestFwNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	vr := "mockVr"
	area := "area 0"
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(vr, area, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(vr, area, tc.conf.Name)
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

func TestDefaults(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<entry name="myinterface">
    <enable>yes</enable>
</entry>
`)
	ans, err := ns.Get("vr", "area", "myinterface")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	}
	if !ans.Enable {
		t.Errorf("Not enabled")
	}
	if ans.Metric != 10 {
		t.Errorf("Metric is not 10")
	}
	if ans.Priority != 1 {
		t.Errorf("Priority is not 1")
	}
	if ans.HelloInterval != 10 {
		t.Errorf("Hello interval is not 10")
	}
	if ans.DeadCounts != 4 {
		t.Errorf("Dead counts is not 4")
	}
	if ans.RetransmitInterval != 5 {
		t.Errorf("Retransmit interval is not 5")
	}
	if ans.TransitDelay != 1 {
		t.Errorf("Transit delay is not 1")
	}
	if ans.GraceRestartDelay != 10 {
		t.Errorf("Grace restart delay is not 10")
	}
}
