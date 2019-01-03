package conadv

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
        {"minimal", Entry{
            Name: "one",
            Enable: false,
        }},
        {"standard", Entry{
            Name: "two",
            Enable: true,
            UsedBy: []string{"one", "two"},
        }},
        {"with raw", Entry{
            Name: "three",
            Enable: true,
            UsedBy: []string{"one", "two"},
            raw: map[string] string{
                "af": "advertisement filters",
                "nf": "non exist filters",
            },
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwConAdv{}
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
