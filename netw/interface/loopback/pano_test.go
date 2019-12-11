package loopback

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		version    version.Number
		tmpl       string
		vsys       string
		importVsys string
		imports    []string
		conf       Entry
	}{
		{version.Number{5, 0, 0, ""}, "one", "vsys2", "vsys2", []string{"loopback.1"}, Entry{
			Name:              "loopback.1",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			NetflowProfile:    "some profile",
			Comment:           "v1 basic",
		}},
		{version.Number{5, 0, 0, ""}, "two", "vsys3", "vsys3", []string{"loopback.2"}, Entry{
			Name:      "loopback.2",
			StaticIps: []string{"10.3.1.1/24"},
			raw: map[string]string{
				"ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
			},
			Comment: "v1 raw no import",
		}},
		{version.Number{8, 0, 0, ""}, "three", "vsys2", "vsys2", []string{"loopback.3"}, Entry{
			Name:              "loopback.3",
			StaticIps:         []string{"10.1.1.1/24", "10.2.1.1/24"},
			ManagementProfile: "enable ping",
			Mtu:               1500,
			AdjustTcpMss:      true,
			Ipv4MssAdjust:     12,
			Ipv6MssAdjust:     14,
			NetflowProfile:    "some profile",
			Comment:           "v2 basic",
		}},
		{version.Number{8, 0, 0, ""}, "four", "vsys3", "vsys3", []string{"loopback.4"}, Entry{
			Name:          "loopback.4",
			StaticIps:     []string{"10.3.1.1/24"},
			Ipv6MssAdjust: 1024,
			raw: map[string]string{
				"ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
			},
			Comment: "v2 raw no import",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoLoopback{}
	ns.Initialize(mc)

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
