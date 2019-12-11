package tunnel

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		version    version.Number
		vsys       string
		importVsys string
		imports    []string
		conf       Entry
	}{
		{version.Number{5, 0, 0, ""}, "vsys2", "vsys2", []string{"tunnel.1"}, Entry{
			Name:              "tunnel.1",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			NetflowProfile:    "some profile",
			Comment:           "v1 basic",
		}},
		{version.Number{5, 0, 0, ""}, "", "", []string{"tunnel.2"}, Entry{
			Name:      "tunnel.2",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
			},
			Comment: "v1 raw no import",
		}},
		{version.Number{8, 0, 0, ""}, "vsys2", "vsys2", []string{"tunnel.3"}, Entry{
			Name:              "tunnel.3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			NetflowProfile:    "some profile",
			Comment:           "v2 basic",
		}},
		{version.Number{8, 0, 0, ""}, "", "", []string{"tunnel.4"}, Entry{
			Name:      "tunnel.4",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
			},
			Comment: "v2 raw no import",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwTunnel{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.conf.Comment, func(t *testing.T) {
			var err error
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err = ns.Set(tc.vsys, tc.conf)
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
