package vlan

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		vsys    string
		imports []string
		conf    Entry
	}{
		{"empty vlan", "", []string{"one"}, Entry{
			Name: "one",
		}},
		{"full vlan spec", "vsys2", []string{"two"}, Entry{
			Name:          "two",
			VlanInterface: "ethernet1/3",
			Interfaces:    []string{"ethernet1/1", "ethernet1/2"},
			StaticMacs: map[string]string{
				"00:30:48:52:ab:cd": "ethernet1/1",
				"00:30:48:52:11:22": "ethernet1/2",
			},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwVlan{}
	ns.Initialize(mc)

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
