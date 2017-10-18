package srvc

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
        {"tcp service no source port no tag", "", Entry{
            Name: "tcp1",
            Description: "my service",
            Protocol: "tcp",
            DestinationPort: "1234",
        }},
        {"tcp service no source port with tag", "", Entry{
            Name: "tcp2",
            Description: "my service",
            Protocol: "tcp",
            DestinationPort: "1234",
            Tag: []string{"tag1", "tag2"},
        }},
        {"tcp service with source port no tag", "", Entry{
            Name: "tcp3",
            Description: "my service",
            Protocol: "tcp",
            SourcePort: "1025",
            DestinationPort: "1234",
        }},
        {"tcp service with source port with tag", "", Entry{
            Name: "tcp4",
            Description: "my service",
            Protocol: "tcp",
            SourcePort: "1025",
            DestinationPort: "1234",
            Tag: []string{"tag1", "tag2"},
        }},
        {"udp service no source port no tag", "", Entry{
            Name: "udp1",
            Description: "my service",
            Protocol: "udp",
            DestinationPort: "1234",
        }},
        {"udp service no source port with tag", "", Entry{
            Name: "udp2",
            Description: "my service",
            Protocol: "udp",
            DestinationPort: "1234",
            Tag: []string{"tag1", "tag2"},
        }},
        {"udp service with source port no tag", "", Entry{
            Name: "udp3",
            Description: "my service",
            Protocol: "udp",
            SourcePort: "1025",
            DestinationPort: "1234",
        }},
        {"udp service with source port with tag", "", Entry{
            Name: "udp4",
            Description: "my service",
            Protocol: "udp",
            SourcePort: "1025",
            DestinationPort: "1234",
            Tag: []string{"tag1", "tag2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &Srvc{}
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

