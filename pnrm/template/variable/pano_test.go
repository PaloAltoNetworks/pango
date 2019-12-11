package variable

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		conf Entry
	}{
		{"ip netmask test", Entry{
			Name:  "$ipnm1",
			Type:  TypeIpNetmask,
			Value: "10.1.1.1/24",
		}},
		{"ip range test", Entry{
			Name:  "$range1",
			Type:  TypeIpRange,
			Value: "10.1.1.1-10.1.1.255",
		}},
		{"fqdn test", Entry{
			Name:  "$fqdn1",
			Type:  TypeFqdn,
			Value: "example.com",
		}},
		{"group id test", Entry{
			Name:  "$grp1",
			Type:  TypeGroupId,
			Value: "42",
		}},
		{"interface test", Entry{
			Name:  "$iface1",
			Type:  TypeInterface,
			Value: "ethernet1/1",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &Variable{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Reset()
			mc.AddResp("")
			err := ns.Set("my template", "", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("my template", "", tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
