package group

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        conf Entry
    }{
        {"ebgp check one", Entry{
            Name: "ebgpOne",
            Enable: true,
            AggregatedConfedAsPath: false,
            SoftResetWithStoredInfo: true,
            Type: TypeEbgp,
            ExportNextHop: NextHopResolve,
            ImportNextHop: NextHopUsePeer,
            RemovePrivateAs: false,
        }},
        {"ebgp check two with raw info", Entry{
            Name: "ebgpTwo",
            Enable: false,
            AggregatedConfedAsPath: true,
            SoftResetWithStoredInfo: false,
            Type: TypeEbgp,
            ExportNextHop: NextHopUseSelf,
            ImportNextHop: NextHopOriginal,
            RemovePrivateAs: true,
            raw: map[string] string{
                "peer": "peer info",
            },
        }},
        {"ebgp confed", Entry{
            Name: "ebgpConfed",
            Enable: true,
            AggregatedConfedAsPath: true,
            SoftResetWithStoredInfo: true,
            Type: TypeEbgpConfed,
            ExportNextHop: NextHopOriginal,
        }},
        {"ibgp", Entry{
            Name: "ibgp",
            Enable: true,
            AggregatedConfedAsPath: true,
            SoftResetWithStoredInfo: true,
            Type: TypeIbgp,
            ExportNextHop: NextHopOriginal,
        }},
        {"ibgp confed", Entry{
            Name: "ibgpConfed",
            Enable: true,
            AggregatedConfedAsPath: true,
            SoftResetWithStoredInfo: true,
            Type: TypeIbgpConfed,
            ExportNextHop: NextHopOriginal,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwGroup{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set("vr", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("vr", tc.conf.Name)
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
