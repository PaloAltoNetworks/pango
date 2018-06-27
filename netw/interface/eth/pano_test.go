package eth

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
        {version.Number{5, 0, 0, ""}, "one", "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
            Name: "ethernet1/1",
            Mode: "layer3",
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            Ipv6Enabled: true,
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            NetflowProfile: "some profile",
            LinkSpeed: "auto",
            LinkDuplex: "auto",
            LinkState: "up",
            Comment: "v1 basic l3",
        }},
        {version.Number{5, 0, 0, ""}, "two", "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
            Name: "ethernet1/2",
            Mode: "layer3",
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 12,
            Comment: "v1 dhcp",
        }},
        {version.Number{5, 0, 0, ""}, "three", "vsys4", "vsys4", []string{}, Entry{
            Name: "ethernet1/3",
            Mode: "ha",
            Comment: "v1 ha no import",
        }},
        {version.Number{5, 0, 0, ""}, "four", "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
            Name: "ethernet1/4",
            Mode: "layer3",
            raw: map[string] string{
                "arp": "<arp>raw arp</arp>",
                "ipv6": "<address>raw ipv6 addresses</address>",
                "l3subinterface": "<units>raw l3 subinterfaces</units>",
            },
            Comment: "v1 layer3 with raw config",
        }},
        {version.Number{5, 0, 0, ""}, "five", "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
            Name: "ethernet1/5",
            Mode: "layer2",
            raw: map[string] string{
                "l2subinterface": "<units>raw l2 subinterfaces</units>",
            },
            Comment: "v1 layer2 with raw config",
        }},
        {version.Number{8, 0, 0, ""}, "six", "vsys2", "vsys2", []string{"ethernet1/1"}, Entry{
            Name: "ethernet1/1",
            Mode: "layer3",
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            Ipv6Enabled: true,
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            Ipv4MssAdjust: 42,
            Ipv6MssAdjust: 84,
            NetflowProfile: "some profile",
            LinkSpeed: "auto",
            LinkDuplex: "auto",
            LinkState: "up",
            Comment: "v2 basic l3",
        }},
        {version.Number{8, 0, 0, ""}, "seven", "vsys3", "vsys3", []string{"ethernet1/2"}, Entry{
            Name: "ethernet1/2",
            Mode: "layer3",
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 12,
            Comment: "v2 dhcp",
        }},
        {version.Number{8, 0, 0, ""}, "eight", "vsys4", "vsys4", []string{}, Entry{
            Name: "ethernet1/3",
            Mode: "ha",
            Comment: "v2 ha no import",
        }},
        {version.Number{8, 0, 0, ""}, "nine", "vsys5", "vsys5", []string{"ethernet1/4"}, Entry{
            Name: "ethernet1/4",
            Mode: "layer3",
            raw: map[string] string{
                "arp": "<arp>raw arp</arp>",
                "ipv6": "<address>raw ipv6 addresses</address>",
                "l3subinterface": "<units>raw l3 subinterfaces</units>",
            },
            Comment: "v2 layer3 with raw config",
        }},
        {version.Number{8, 0, 0, ""}, "ten", "vsys6", "vsys6", []string{"ethernet1/5"}, Entry{
            Name: "ethernet1/5",
            Mode: "layer2",
            raw: map[string] string{
                "l2subinterface": "<units>raw l2 subinterfaces</units>",
            },
            Comment: "v2 layer2 with raw config",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoEth{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.conf.Comment, func(t *testing.T) {
            var err error
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err = ns.Set(tc.tmpl, tc.vsys, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.tmpl, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                }
                if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
                if tc.tmpl != mc.Template {
                    t.Errorf("Template: %q != %q", tc.tmpl, mc.Template)
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
