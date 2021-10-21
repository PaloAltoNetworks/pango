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
		{"v2c check", version.Number{7, 0, 0, ""}, Entry{
			Name: "v2c",
			V2cServers: []V2cServer{
				{
					Name:      "t1",
					Manager:   "snmp.example.com",
					Community: "public",
				},
			},
		}},
		{"v3 check", version.Number{7, 0, 0, ""}, Entry{
			Name: "v3",
			V3Servers: []V3Server{
				{
					Name:         "t1",
					Manager:      "snmp.example.com",
					User:         "jdoe",
					EngineId:     "0A",
					AuthPassword: "auth",
					PrivPassword: "priv",
				},
			},
		}},
	}
}
