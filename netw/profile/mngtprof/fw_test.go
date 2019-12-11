package mngtprof

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestFwNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		conf Entry
	}{
		{"test1", Entry{
			Name:                    "test1",
			Ping:                    true,
			Telnet:                  true,
			Ssh:                     true,
			Https:                   true,
			UseridSyslogListenerSsl: true,
		}},
		{"test2", Entry{
			Name:                    "test2",
			Http:                    true,
			HttpOcsp:                true,
			Snmp:                    true,
			ResponsePages:           true,
			UseridService:           true,
			UseridSyslogListenerUdp: true,
		}},
		{"test3", Entry{
			Name:         "test3",
			Ping:         true,
			PermittedIps: []string{"10.1.1.1", "10.2.2.2"},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &FwMngtProf{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
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
