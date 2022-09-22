package dg

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
			mc.Version = tc.version
			mc.Reset()
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

func TestDevices(t *testing.T) {
	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	mc.Version = version.Number{10, 0, 0, ""}
	mc.AddResp(`
<entry name="foo">
    <description>war</description>
    <devices>
        <entry name="first"/>
        <entry name="second"/>
        <entry name="third">
            <vsys>
                <entry name="vsys2"/>
                <entry name="vsys3"/>
            </vsys>
        </entry>
        <entry name="fourth"/>
    </devices>
</entry>`)

	obj, err := ns.Get("foo")
	if err != nil {
		t.Fatalf("Error in get: %s", err)
	}
	if obj.Name != "foo" {
		t.Fatalf("Name is not foo, but %q", obj.Name)
	}
	if obj.Description != "war" {
		t.Fatalf("Description is %q, not war", obj.Description)
	}
	if len(obj.Devices) != 4 {
		t.Fatalf("Devices is len %d, not 4", len(obj.Devices))
	}

	val, ok := obj.Devices["first"]
	if !ok {
		t.Fatalf("first not in devices")
	} else if val != nil {
		t.Fatalf("first val is not nil")
	}

	val, ok = obj.Devices["second"]
	if !ok {
		t.Fatalf("second not in devices")
	} else if val != nil {
		t.Fatalf("second val is not nil")
	}

	val, ok = obj.Devices["third"]
	if !ok {
		t.Fatalf("third not in devices")
	} else if len(val) != 2 {
		t.Fatalf("third vsys list is %#v not [vsys2, vsys3]", val)
	} else if val[0] != "vsys2" && val[1] != "vsys2" {
		t.Errorf("vsys2 not in val: %#v", val)
	} else if val[0] != "vsys3" && val[1] != "vsys3" {
		t.Errorf("vsys3 not in val: %#v", val)
	}

	val, ok = obj.Devices["fourth"]
	if !ok {
		t.Fatalf("fourth not in devices")
	} else if val != nil {
		t.Fatalf("fourth val is not nil")
	}

}
