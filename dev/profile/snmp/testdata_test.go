package snmp

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v2c check", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			SnmpVersion: SnmpVersionV2c,
		}},
		{"v2c check with servers", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t2",
			SnmpVersion: SnmpVersionV2c,
			raw: map[string]string{
				"v2c": "v2c servers",
			},
		}},
		{"v3 check", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t3",
			SnmpVersion: SnmpVersionV3,
		}},
		{"v3 check with servers", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t4",
			SnmpVersion: SnmpVersionV3,
			raw: map[string]string{
				"v3": "v3 servers",
			},
		}},
	}
}
