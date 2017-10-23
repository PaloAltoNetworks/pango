package nat

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
    "github.com/PaloAltoNetworks/pango/util"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        base string
        conf Entry
    }{
        {"dst only", "", util.Rulebase, Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            From: []string{"zone1", "zone2"},
            To: []string{"zone3"},
            ToInterface: "ethernet1/7",
            Service: "myService",
            Source: []string{"any"},
            Destination: []string{"any"},
            SatType: None,
            DatAddress: "10.5.1.1",
            DatPort: 1234,
            Tag: []string{"tag1", "tag2"},
        }},
        {"dynamic ip and port with translated address", "", util.Rulebase, Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            From: []string{"zone3"},
            To: []string{"zone4"},
            ToInterface: "any",
            Service: "any",
            Source: []string{"any"},
            Destination: []string{"any"},
            SatType: DynamicIpAndPort,
            SatAddressType: TranslatedAddress,
            SatTranslatedAddress: []string{"10.3.1.1", "10.3.2.1"},
        }},
        {"dynamic ip with interface address fallback", "vsys2", util.PreRulebase, Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            From: []string{"zone3"},
            To: []string{"zone4"},
            ToInterface: "any",
            Service: "any",
            Source: []string{"any"},
            Destination: []string{"any"},
            SatType: DynamicIp,
            SatTranslatedAddress: []string{"10.5.5.5", "10.6.6.6"},
            SatFallbackType: InterfaceAddress,
            SatFallbackInterface: "ethernet1/7",
            SatFallbackIpType: Ip,
            SatFallbackIpAddress: "10.10.10.10",
        }},
        {"static ip with target", "vsys3", util.PostRulebase, Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            From: []string{"zone3"},
            To: []string{"zone4"},
            ToInterface: "any",
            Service: "any",
            Source: []string{"any"},
            Destination: []string{"any"},
            SatType: StaticIp,
            SatStaticTranslatedAddress: "10.5.5.5",
            SatStaticBiDirectional: true,
            Target: []string{"fw1", "fw2"},
            NegateTarget: true,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &Nat{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.vsys, tc.base, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.vsys, tc.base, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
