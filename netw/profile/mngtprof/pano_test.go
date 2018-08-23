package mngtprof

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        tmpl string
        conf Entry
    }{
        {"test1", "a", Entry{
            Name: "test1",
            Ping: true,
            Telnet: true,
            Ssh: true,
            Https: true,
            UseridSyslogListenerSsl: true,
        }},
        {"test2", "b", Entry{
            Name: "test2",
            Http: true,
            HttpOcsp: true,
            Snmp: true,
            ResponsePages: true,
            UseridService: true,
            UseridSyslogListenerUdp: true,
        }},
        {"test3", "c", Entry{
            Name: "test3",
            Ping: true,
            PermittedIps: []string{"10.1.1.1", "10.2.2.2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &PanoMngtProf{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.AddResp("")
            err := ns.Set(tc.tmpl, "", tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.tmpl, "", tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
