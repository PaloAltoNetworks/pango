package addr

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        dg string
        conf Entry
    }{
        {"test ip netmask", "", Entry{
            Name: "one",
            Value: "10.1.1.0/24",
            Type: IpNetmask,
            Description: "my description",
            Tags: []string{"tag1", "tag2"},
        }},
        {"test ip range", "dg1", Entry{
            Name: "two",
            Value: "10.1.1.1-10.1.1.254",
            Type: IpRange,
            Description: "my description",
            Tags: []string{"tag3", "tag4"},
        }},
        {"test fqdn", "dg2", Entry{
            Name: "three",
            Value: "example.com",
            Type: Fqdn,
            Description: "my description",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoAddr{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.dg, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.dg, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}

