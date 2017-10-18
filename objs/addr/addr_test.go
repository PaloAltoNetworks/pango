package addr

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/xapi/testdata"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        conf Entry
    }{
        {"test ip netmask", "", Entry{
            Name: "one",
            Value: "10.1.1.0/24",
            Type: IpNetmask,
            Description: "my description",
            Tag: []string{"tag1", "tag2"},
        }},
        {"test ip range", "vsys2", Entry{
            Name: "two",
            Value: "10.1.1.1-10.1.1.254",
            Type: IpRange,
            Description: "my description",
            Tag: []string{"tag3", "tag4"},
        }},
        {"test fqdn", "vsys3", Entry{
            Name: "three",
            Value: "example.com",
            Type: Fqdn,
            Description: "my description",
        }},
    }

    mc := &testdata.MockClient{}
    ns := &Addr{}
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

