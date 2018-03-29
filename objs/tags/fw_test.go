package tags

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
        {"test with all fields", "", Entry{
            Name: "one",
            Color: "color1",
            Comment: "first test",
        }},
        {"test no color", "", Entry{
            Name: "two",
            Comment: "second test",
        }},
        {"test no comment", "vsys1", Entry{
            Name: "three",
            Color: "color3",
        }},
        {"test no color or comment", "vsys2", Entry{
            Name: "four",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwTags{}
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

