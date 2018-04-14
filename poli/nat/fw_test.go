package nat

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        conf Entry
    }{
        {"dst only", "", Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            SourceZones: []string{"zone1", "zone2"},
            DestinationZone: "zone3",
            ToInterface: "ethernet1/7",
            Service: "myService",
            SourceAddresses: []string{"any"},
            DestinationAddresses: []string{"any"},
            SatType: None,
            DatAddress: "10.5.1.1",
            DatPort: 1234,
            Tags: []string{"tag1", "tag2"},
        }},
        {"dynamic ip and port with translated address", "", Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            SourceZones: []string{"zone3"},
            DestinationZone: "zone4",
            ToInterface: "any",
            Service: "any",
            SourceAddresses: []string{"any"},
            DestinationAddresses: []string{"any"},
            SatType: DynamicIpAndPort,
            SatAddressType: TranslatedAddress,
            SatTranslatedAddresses: []string{"10.3.1.1", "10.3.2.1"},
        }},
        {"dynamic ip with interface address fallback", "vsys2", Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            SourceZones: []string{"zone3"},
            DestinationZone: "zone4",
            ToInterface: "any",
            Service: "any",
            SourceAddresses: []string{"any"},
            DestinationAddresses: []string{"any"},
            SatType: DynamicIp,
            SatTranslatedAddresses: []string{"10.5.5.5", "10.6.6.6"},
            SatFallbackType: InterfaceAddress,
            SatFallbackInterface: "ethernet1/7",
            SatFallbackIpType: Ip,
            SatFallbackIpAddress: "10.10.10.10",
        }},
        {"static ip with target", "vsys3", Entry{
            Name: "nat policy",
            Description: "my nat policy",
            Type: "ipv4",
            SourceZones: []string{"zone3"},
            DestinationZone: "zone4",
            ToInterface: "any",
            Service: "any",
            SourceAddresses: []string{"any"},
            DestinationAddresses: []string{"any"},
            SatType: StaticIp,
            SatStaticTranslatedAddress: "10.5.5.5",
            SatStaticBiDirectional: true,
            Targets: []string{"fw1", "fw2"},
            NegateTarget: true,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwNat{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.vsys, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.vsys, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
