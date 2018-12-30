package bgp

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
        conf Config
    }{
        {"v1 bgp no raw", version.Number{6, 1, 0, ""}, Config{
            Enable: true,
            RouterId: "router id",
            AsNumber: "as number",
            RejectDefaultRoute: false,
            InstallRoute: true,
            AggregateMed: false,
            DefaultLocalPreference: "def local pref",
            AsFormat: AsFormat2Byte,
            AlwaysCompareMed: true,
            DeterministicMedComparison: false,
            EcmpMultiAs: true,
            EnforceFirstAs: false,
            EnableGracefulRestart: true,
            StaleRouteTime: 2,
            LocalRestartTime: 3,
            MaxPeerRestartTime: 4,
            ReflectorClusterId: "reflector cluster id",
            ConfederationMemberAs: "conf mem as",
            AllowRedistributeDefaultRoute: false,
        }},
        {"v1 bgp with raw", version.Number{6, 1, 0, ""}, Config{
            Enable: true,
            RouterId: "router id",
            AsNumber: "as number",
            RejectDefaultRoute: false,
            InstallRoute: true,
            AggregateMed: false,
            DefaultLocalPreference: "def local pref",
            AsFormat: AsFormat4Byte,
            AlwaysCompareMed: true,
            DeterministicMedComparison: false,
            EcmpMultiAs: true,
            EnforceFirstAs: false,
            EnableGracefulRestart: true,
            StaleRouteTime: 2,
            LocalRestartTime: 3,
            MaxPeerRestartTime: 4,
            ReflectorClusterId: "reflector cluster id",
            ConfederationMemberAs: "conf mem as",
            AllowRedistributeDefaultRoute: false,
            raw: map[string] string{
                "orf": "outbound route filter",
                "ap": "auth profile",
                "dp": "dampening profile",
                "pg": "peer group",
                "poli": "policy",
                "rr": "redist rules",
            },
        }},
        {"v2 bgp no raw", version.Number{7, 1, 0, ""}, Config{
            Enable: true,
            RouterId: "router id",
            AsNumber: "as number",
            BfdProfile: "None",
            RejectDefaultRoute: false,
            InstallRoute: true,
            AggregateMed: false,
            DefaultLocalPreference: "def local pref",
            AsFormat: AsFormat2Byte,
            AlwaysCompareMed: true,
            DeterministicMedComparison: false,
            EcmpMultiAs: true,
            EnforceFirstAs: false,
            EnableGracefulRestart: true,
            StaleRouteTime: 2,
            LocalRestartTime: 3,
            MaxPeerRestartTime: 4,
            ReflectorClusterId: "reflector cluster id",
            ConfederationMemberAs: "conf mem as",
            AllowRedistributeDefaultRoute: false,
        }},
        {"v2 bgp with raw", version.Number{7, 1, 0, ""}, Config{
            Enable: true,
            RouterId: "router id",
            AsNumber: "as number",
            BfdProfile: "bfd profile",
            RejectDefaultRoute: false,
            InstallRoute: true,
            AggregateMed: false,
            DefaultLocalPreference: "def local pref",
            AsFormat: AsFormat4Byte,
            AlwaysCompareMed: true,
            DeterministicMedComparison: false,
            EcmpMultiAs: true,
            EnforceFirstAs: false,
            EnableGracefulRestart: true,
            StaleRouteTime: 2,
            LocalRestartTime: 3,
            MaxPeerRestartTime: 4,
            ReflectorClusterId: "reflector cluster id",
            ConfederationMemberAs: "conf mem as",
            AllowRedistributeDefaultRoute: false,
            raw: map[string] string{
                "orf": "outbound route filter",
                "ap": "auth profile",
                "dp": "dampening profile",
                "pg": "peer group",
                "poli": "policy",
                "rr": "redist rules",
            },
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoBgp{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Version = tc.version
            mc.Reset()
            mc.AddResp("")
            err := ns.Set("my template", "", "mockVr", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("my template", "", "mockVr")
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
