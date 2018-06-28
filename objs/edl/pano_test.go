package edl

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
        dg string
        conf Entry
    }{
        {"v1 ip hourly with at", version.Number{6, 1, 0, ""}, "", Entry{
            Name: "one",
            Type: TypeIp,
            Description: "first",
            Source: "http://1.1.1.1",
            Repeat: RepeatHourly,
            RepeatAt: "15",
        }},
        {"v1 ip hourly without at", version.Number{6, 1, 0, ""}, "w", Entry{
            Name: "two",
            Type: TypeIp,
            Description: "second",
            Source: "http://2.2.2.2",
            Repeat: RepeatHourly,
        }},
        {"v1 domain daily", version.Number{6, 1, 0, ""}, "x", Entry{
            Name: "three",
            Type: TypeDomain,
            Description: "third",
            Source: "http://3.3.3.3",
            Repeat: RepeatDaily,
            RepeatAt: "12:34",
        }},
        {"v1 url weekly", version.Number{6, 1, 0, ""}, "y", Entry{
            Name: "four",
            Type: TypeUrl,
            Description: "fourth",
            Source: "http://4.4.4.4",
            Repeat: RepeatWeekly,
            RepeatAt: "actually invalid",
            RepeatDayOfWeek: "tuesday",
        }},
        {"v1 url monthly", version.Number{6, 1, 0, ""}, "z", Entry{
            Name: "five",
            Type: TypeUrl,
            Description: "fifth",
            Source: "http://5.5.5.5",
            Repeat: RepeatMonthly,
            RepeatAt: "also invalid",
            RepeatDayOfMonth: 7,
        }},
        {"v2 ip 5min", version.Number{8, 0, 0, ""}, "", Entry{
            Name: "one",
            Type: TypeIp,
            Description: "first",
            Source: "http://1.1.1.1",
            Exceptions: []string{"10.1.1.1", "10.1.1.2"},
            Repeat: RepeatEveryFiveMinutes,
        }},
        {"v2 ip hourly", version.Number{8, 0, 0, ""}, "w", Entry{
            Name: "two",
            Type: TypeIp,
            Description: "second",
            CertificateProfile: "my cert profile",
            Username: "jack",
            Password: "burton",
            Source: "http://2.2.2.2",
            Repeat: RepeatHourly,
        }},
        {"v2 domain daily", version.Number{8, 0, 0, ""}, "x", Entry{
            Name: "three",
            Type: TypeDomain,
            Description: "third",
            Source: "http://3.3.3.3",
            Exceptions: []string{"paloaltonetworks.com"},
            Repeat: RepeatDaily,
            RepeatAt: "13",
        }},
        {"v2 url weekly", version.Number{8, 0, 0, ""}, "y", Entry{
            Name: "four",
            Type: TypeUrl,
            Description: "fourth",
            Source: "http://4.4.4.4",
            CertificateProfile: "another profile",
            Repeat: RepeatWeekly,
            RepeatAt: "actually invalid",
            RepeatDayOfWeek: "tuesday",
        }},
        {"v2 url monthly", version.Number{8, 0, 0, ""}, "z", Entry{
            Name: "five",
            Type: TypeUrl,
            Description: "fifth",
            Source: "http://5.5.5.5",
            Username: "kung",
            Password: "fu",
            Repeat: RepeatMonthly,
            RepeatAt: "also invalid",
            RepeatDayOfMonth: 7,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoEdl{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.Version = tc.version
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

