package vlan

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/version"
    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        version version.Number
        tmpl string
        vsys string
        importVsys string
        imports []string
        conf Entry
    }{
        {version.Number{5, 0, 0, ""}, "one", "vsys2", "vsys2", []string{"vlan.1"}, Entry{
            Name: "vlan.1",
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            NetflowProfile: "some profile",
            Comment: "v1 basic",
            raw: map[string] string{
                "arp": "<arp>raw arp</arp>",
                "ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
                "ndp": "<ndp-proxy>raw ndp</ndp-proxy>",
            },
        }},
        {version.Number{5, 0, 0, ""}, "two", "vsys3", "vsys3", []string{"vlan.2"}, Entry{
            Name: "vlan.2",
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 12,
            Comment: "v1 dhcp",
        }},
        {version.Number{8, 0, 0, ""}, "four", "vsys2", "vsys2", []string{"vlan.4"}, Entry{
            Name: "vlan.4",
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            Ipv4MssAdjust: 12,
            Ipv6MssAdjust: 14,
            NetflowProfile: "some profile",
            Comment: "v2 basic",
            raw: map[string] string{
                "arp": "<arp>raw arp</arp>",
                "ipv6": "<ipv6>raw ipv6 addresses</ipv6>",
                "ndp": "<ndp-proxy>raw ndp</ndp-proxy>",
            },
        }},
        {version.Number{8, 0, 0, ""}, "five", "vsys3", "vsys3", []string{"vlan.5"}, Entry{
            Name: "vlan.5",
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 12,
            Comment: "v2 dhcp",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoVlan{}
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
