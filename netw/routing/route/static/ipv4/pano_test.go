package ipv4

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
    "github.com/PaloAltoNetworks/pango/version"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        version version.Number
        vr string
        conf Entry
    }{
        {"v1 route no nexthop", version.Number{5, 0, 0, ""}, "v1", Entry{
            Name: "one",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableNoInstall,
        }},
        {"v1 route discard nexthop", version.Number{5, 0, 0, ""}, "v1", Entry{
            Name: "two",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopDiscard,
            AdminDistance: 101,
            Metric: 111,
        }},
        {"v1 route ip address nexthop", version.Number{5, 0, 0, ""}, "v1", Entry{
            Name: "three",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopIpAddress,
            NextHop: "10.2.3.4",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableNoInstall,
        }},
        {"v1 route next vr nexthop", version.Number{5, 0, 0, ""}, "v1", Entry{
            Name: "four",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopNextVr,
            NextHop: "default",
            AdminDistance: 101,
            Metric: 111,
        }},
        {"v2 route no nexthop", version.Number{7, 1, 0, ""}, "v1", Entry{
            Name: "one",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableNoInstall,
        }},
        {"v2 route discard nexthop", version.Number{7, 1, 0, ""}, "v1", Entry{
            Name: "two",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopDiscard,
            AdminDistance: 101,
            Metric: 111,
            BfdProfile: "bfd two",
        }},
        {"v2 route ip address nexthop", version.Number{7, 1, 0, ""}, "v1", Entry{
            Name: "three",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopIpAddress,
            NextHop: "10.2.3.4",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableNoInstall,
            BfdProfile: "bfd three",
        }},
        {"v2 route next vr nexthop", version.Number{7, 1, 0, ""}, "v1", Entry{
            Name: "four",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopNextVr,
            NextHop: "default",
            AdminDistance: 101,
            Metric: 111,
        }},
        {"v3 route no nexthop", version.Number{8, 0, 0, ""}, "v1", Entry{
            Name: "one",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableNoInstall,
        }},
        {"v3 route discard nexthop", version.Number{8, 0, 0, ""}, "v1", Entry{
            Name: "two",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopDiscard,
            AdminDistance: 101,
            Metric: 111,
            BfdProfile: "bfd two",
            RouteTable: RouteTableUnicast,
        }},
        {"v3 route ip address nexthop", version.Number{8, 0, 0, ""}, "v1", Entry{
            Name: "three",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopIpAddress,
            NextHop: "10.2.3.4",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableMulticast,
            BfdProfile: "bfd three",
        }},
        {"v3 route next vr nexthop", version.Number{8, 0, 0, ""}, "v1", Entry{
            Name: "four",
            Destination: "0.0.0.0/0",
            Interface: "ethernet1/1",
            Type: NextHopNextVr,
            NextHop: "default",
            AdminDistance: 101,
            Metric: 111,
            RouteTable: RouteTableBoth,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoIpv4{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err := ns.Set("some template", "", tc.vr, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("some template", "", tc.vr, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                }
                if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
