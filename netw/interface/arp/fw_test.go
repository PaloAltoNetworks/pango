package arp

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		iType   string
		iName   string
		subName string
		conf    Entry
	}{
		{"eth no sub", TypeEthernet, "ethernet1/5", "", Entry{
			Ip:         "10.1.1.1",
			MacAddress: "00:30:48:52:ab:c1",
		}},
		{"eth with sub", TypeEthernet, "ethernet1/5", "ethernet1/5.7", Entry{
			Ip:         "10.1.1.2",
			MacAddress: "00:30:48:52:ab:c2",
		}},
		{"agg no sub", TypeAggregate, "ae5", "", Entry{
			Ip:         "10.1.1.3",
			MacAddress: "00:30:48:52:ab:c3",
		}},
		{"agg with sub", TypeAggregate, "ae5", "ae5.4", Entry{
			Ip:         "10.1.1.4",
			MacAddress: "00:30:48:52:ab:c4",
		}},
		{"vlan", TypeVlan, "", "vlan.4", Entry{
			Ip:         "10.1.1.5",
			MacAddress: "00:30:48:52:ab:c5",
			Interface:  "vlan.2",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwArp{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.iType, tc.iName, tc.subName, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.iType, tc.iName, tc.subName, tc.conf.Ip)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
