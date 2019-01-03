package nonexist

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
    "github.com/PaloAltoNetworks/pango/version"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        version version.Number
        conf Entry
    }{
        {"v1 basic", version.Number{7, 0, 0, ""}, Entry{
            Name: "one",
            Enable: false,
        }},
        {"v1 full", version.Number{7, 0, 0, ""}, Entry{
            Name: "two",
            Enable: true,
            AsPathRegex: "as path regex",
            CommunityRegex: "community regex",
            ExtendedCommunityRegex: "extended community regex",
            Med: "match med",
            AddressPrefix: []string{"5.5.0.0", "6.6.0.0"},
            NextHop: []string{"nh1", "nh2"},
            FromPeer: []string{"fp1", "fp2"},
        }},
        {"v2 basic", version.Number{8, 0, 0, ""}, Entry{
            Name: "three",
            Enable: false,
        }},
        {"v2 full", version.Number{8, 0, 0, ""}, Entry{
            Name: "four",
            Enable: true,
            AsPathRegex: "as path regex",
            CommunityRegex: "community regex",
            ExtendedCommunityRegex: "extended community regex",
            Med: "match med",
            RouteTable: RouteTableUnicast,
            AddressPrefix: []string{"5.5.0.0", "6.6.0.0"},
            NextHop: []string{"nh1", "nh2"},
            FromPeer: []string{"fp1", "fp2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwNonExist{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.Version = tc.version
            mc.AddResp("")
            err := ns.Set("vr", "ca", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("vr", "ca", tc.conf.Name)
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
