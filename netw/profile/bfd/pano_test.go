package bfd

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        conf Entry
    }{
        {"test1", Entry{
            Name: "test1",
            Mode: ModeActive,
            MinimumTxInterval: 3,
            MinimumRxInterval: 4,
            DetectionMultiplier: 5,
            HoldTime: 6,
        }},
        {"test2", Entry{
            Name: "test2",
            MinimumRxTtl: 7,
        }},
        {"test3", Entry{
            Name: "test3",
            Mode: ModePassive,
            HoldTime: 42,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoBfd{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.AddResp("")
            err := ns.Set("tem1", "", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get("tem1", "", tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
