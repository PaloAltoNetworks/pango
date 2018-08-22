package dg

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        conf Entry
    }{
        {"test one", Entry{
            Name: "one",
            Description: "my description",
        }},
        {"test two", Entry{
            Name: "two",
            Description: "with devices",
            Devices: map[string] []string{
                "001234": nil,
                "998765": {"vsys3", "vsys5"},
            },
        }},
        {"test three", Entry{
            Name: "three",
            Description: "with one device",
            Devices: map[string] []string{
                "001234": nil,
            },
        }},
        {"test four", Entry{
            Name: "four",
            Description: "all the raws",
            raw: map[string] string{
                "ao": "ao",
                "ag": "ag",
                "app": "app",
                "afil": "afil",
                "appg": "appg",
                "as": "as",
                "aobj": "aobj",
                "acode": "acode",
                "email": "email",
                "edl": "edl",
                "ls": "ls",
                "master": "master",
                "pdf": "pdf",
                "postrb": "postrb",
                "prerb": "prerb",
                "profg": "profg",
                "profs": "profs",
                "reg": "reg",
                "repg": "repg",
                "rep": "rep",
                "schd": "schd",
                "srv": "srv",
                "srvg": "srvg",
                "tag": "tag",
                "thr": "thr",
                "tsv": "tsv",
            },
        }},
    }

    mc := &testdata.MockClient{}
    ns := &Dg{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}

