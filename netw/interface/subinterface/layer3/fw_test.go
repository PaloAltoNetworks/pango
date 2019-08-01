package layer3

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/version"
    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        version version.Number
        vsys string
        importVsys string
        imports []string
        conf Entry
    }{
        {version.Number{7, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.2"}, Entry{
            Name: "ethernet1/1.2",
            Tag: 2,
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            Ipv6Enabled: true,
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            Ipv4MssAdjust: 5,
            Ipv6MssAdjust: 6,
            NetflowProfile: "some profile",
            Comment: "v1 basic",
        }},
        {version.Number{7, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.3"}, Entry{
            Name: "ethernet1/1.3",
            Tag: 3,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp without default route",
        }},
        {version.Number{7, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.4"}, Entry{
            Name: "ethernet1/1.4",
            Tag: 4,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp with default route",
        }},
        {version.Number{7, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.5"}, Entry{
            Name: "ethernet1/1.5",
            Tag: 5,
            Ipv6Enabled: true,
            Ipv6InterfaceId: "v6-interface-id",
            ManagementProfile: "block everything",
            Mtu: 2600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            Comment: "v1 dhcp with raw",
            raw: map[string] string{
                "arp": "arp entries",
                "ndp": "ndp info",
                "v6adr": "ipv6 addresses",
                "v6nbr": "ipv6 neighbor info",
            },
        }},
        {version.Number{8, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.2"}, Entry{
            Name: "ethernet1/1.2",
            Tag: 2,
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            Ipv6Enabled: true,
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            Ipv4MssAdjust: 5,
            Ipv6MssAdjust: 6,
            NetflowProfile: "some profile",
            Comment: "v1 basic",
        }},
        {version.Number{8, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.3"}, Entry{
            Name: "ethernet1/1.3",
            Tag: 3,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp without default route",
        }},
        {version.Number{8, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.4"}, Entry{
            Name: "ethernet1/1.4",
            Tag: 4,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp with default route",
        }},
        {version.Number{8, 1, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.5"}, Entry{
            Name: "ethernet1/1.5",
            Tag: 5,
            DecryptForward: true,
            Ipv6Enabled: true,
            Ipv6InterfaceId: "v6-interface-id",
            ManagementProfile: "block everything",
            Mtu: 2600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            Comment: "v1 dhcp with raw and decrypt forward",
            raw: map[string] string{
                "arp": "arp entries",
                "ndp": "ndp info",
                "v6adr": "ipv6 addresses",
                "v6nbr": "ipv6 neighbor info",
            },
        }},
        {version.Number{9, 0, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.2"}, Entry{
            Name: "ethernet1/1.2",
            Tag: 2,
            StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
            Ipv6Enabled: true,
            ManagementProfile: "enable ping",
            Mtu: 1500,
            AdjustTcpMss: true,
            Ipv4MssAdjust: 5,
            Ipv6MssAdjust: 6,
            NetflowProfile: "some profile",
            Comment: "v1 basic",
        }},
        {version.Number{9, 0, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.3"}, Entry{
            Name: "ethernet1/1.3",
            Tag: 3,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp without default route",
        }},
        {version.Number{9, 0, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.4"}, Entry{
            Name: "ethernet1/1.4",
            Tag: 4,
            ManagementProfile: "enable ping",
            Mtu: 1600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            DhcpDefaultRouteMetric: 5,
            NetflowProfile: "some profile",
            Comment: "v1 dhcp with default route",
        }},
        {version.Number{9, 0, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.5"}, Entry{
            Name: "ethernet1/1.5",
            Tag: 5,
            DecryptForward: true,
            Ipv6Enabled: true,
            Ipv6InterfaceId: "v6-interface-id",
            ManagementProfile: "block everything",
            Mtu: 2600,
            EnableDhcp: true,
            CreateDhcpDefaultRoute: true,
            Comment: "v1 dhcp with raw and decrypt forward",
            raw: map[string] string{
                "arp": "arp entries",
                "ndp": "ndp info",
                "v6adr": "ipv6 addresses",
                "v6nbr": "ipv6 neighbor info",
            },
        }},
        {version.Number{9, 0, 0, ""}, "vsys1", "vsys1", []string{"ethernet1/1.6"}, Entry{
            Name: "ethernet1/1.6",
            Tag: 6,
            ManagementProfile: "allow ssh",
            Mtu: 1100,
            EnableDhcp: true,
            DhcpSendHostnameEnable: true,
            DhcpSendHostnameValue: "foo.bar.com",
            NetflowProfile: "some profile",
            Comment: "v1 dhcp with send hostname",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwLayer3{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.conf.Comment, func(t *testing.T) {
            var err error
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err = ns.Set(EthernetInterface, "ethernet1/1", tc.vsys, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(EthernetInterface, "ethernet1/1", tc.conf.Name)
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
