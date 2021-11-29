package router

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
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
			if tc.doDefaults {
				tc.conf.Defaults()
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

func TestPanoAllDistancesDefault(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := PanoramaNamespace(mc)
	mc.AddResp(`
<entry name="foo">
    <interface>
        <member>ethernet1/1</member>
    </interface>
</entry>
`)
	ans, err := ns.Get("myTmpl", "", "foo")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	} else if len(ans.Interfaces) != 1 {
		t.Errorf("Interfaces is len %d, not 1", len(ans.Interfaces))
	} else if ans.Interfaces[0] != "ethernet1/1" {
		t.Errorf("entry.Interfaces[0] = %q, not \"ethernet1/1\"", ans.Interfaces[0])
	} else {
		if ans.StaticDist != 10 || ans.StaticIpv6Dist != 10 || ans.OspfIntDist != 30 || ans.OspfExtDist != 110 || ans.Ospfv3IntDist != 30 || ans.Ospfv3ExtDist != 110 || ans.IbgpDist != 200 || ans.EbgpDist != 20 || ans.RipDist != 120 {
			t.Errorf("ans defaults are wrong: %#v", ans)
		}
	}
}

func TestPanoPartialDistancesWithDefaults(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := PanoramaNamespace(mc)
	mc.AddResp(`
<entry name="foo">
    <interface>
        <member>ethernet1/1</member>
    </interface>
    <admin-dists>
        <static>1</static>
        <ospf-int>2</ospf-int>
        <ospfv3-int>3</ospfv3-int>
        <ibgp>4</ibgp>
        <rip>5</rip>
    </admin-dists>
</entry>
`)
	ans, err := ns.Get("myTmpl", "", "foo")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	} else if len(ans.Interfaces) != 1 {
		t.Errorf("Interfaces is len %d, not 1", len(ans.Interfaces))
	} else if ans.Interfaces[0] != "ethernet1/1" {
		t.Errorf("entry.Interfaces[0] = %q, not \"ethernet1/1\"", ans.Interfaces[0])
	} else {
		if ans.StaticDist != 1 || ans.StaticIpv6Dist != 10 || ans.OspfIntDist != 2 || ans.OspfExtDist != 110 || ans.Ospfv3IntDist != 3 || ans.Ospfv3ExtDist != 110 || ans.IbgpDist != 4 || ans.EbgpDist != 20 || ans.RipDist != 5 {
			t.Errorf("ans defaults are wrong: %#v", ans)
		}
	}
}
