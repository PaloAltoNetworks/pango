package srvcgrp

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        conf Entry
    }{
        {"test no services", "", Entry{
            Name: "one",
            Tag: []string{"one", "two"},
        }},
        {"test one service", "", Entry{
            Name: "two",
            Services: []string{"svc1"},
            Tag: []string{"single"},
        }},
        {"test two services", "", Entry{
            Name: "three",
            Services: []string{"svc1", "svc2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &SrvcGrp{}
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

