package syslog

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
		{"v1 config", version.Number{7, 1, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
		}},
		{"v1 config with servers", version.Number{7, 1, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
			Servers: []Server{
				{
					Name:         "t1",
					Server:       "syslog.server.int",
					Transport:    TransportSsl,
					Port:         42,
					SyslogFormat: SyslogFormatIetf,
					Facility:     FacilityLocal4,
				},
			},
		}},
		{"v1 system", version.Number{7, 1, 0, ""}, Entry{
			Name:   "t1",
			System: "custom format",
		}},
		{"v1 threat", version.Number{7, 1, 0, ""}, Entry{
			Name:   "t1",
			Threat: "custom format",
		}},
		{"v1 traffic", version.Number{7, 1, 0, ""}, Entry{
			Name:    "t1",
			Traffic: "custom format",
		}},
		{"v1 hip match", version.Number{7, 1, 0, ""}, Entry{
			Name:     "t1",
			HipMatch: "custom format",
		}},
		{"v1 escaped characters", version.Number{7, 1, 0, ""}, Entry{
			Name:              "t1",
			EscapedCharacters: "abc",
		}},
		{"v1 escape character", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			EscapeCharacter: "x",
		}},
		{"v2 config", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
		}},
		{"v2 config with servers", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
			Servers: []Server{
				{
					Name:         "t1",
					Server:       "syslog.server.int",
					Transport:    TransportSsl,
					Port:         42,
					SyslogFormat: SyslogFormatIetf,
					Facility:     FacilityLocal4,
				},
			},
		}},
		{"v2 system", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			System: "custom format",
		}},
		{"v2 threat", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			Threat: "custom format",
		}},
		{"v2 traffic", version.Number{8, 0, 0, ""}, Entry{
			Name:    "t1",
			Traffic: "custom format",
		}},
		{"v2 hip match", version.Number{8, 0, 0, ""}, Entry{
			Name:     "t1",
			HipMatch: "custom format",
		}},
		{"v2 escaped characters", version.Number{8, 0, 0, ""}, Entry{
			Name:              "t1",
			EscapedCharacters: "abc",
		}},
		{"v2 escape character", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			EscapeCharacter: "x",
		}},
		{"v2 url", version.Number{8, 0, 0, ""}, Entry{
			Name: "t1",
			Url:  "custom format",
		}},
		{"v2 data", version.Number{8, 0, 0, ""}, Entry{
			Name: "t1",
			Data: "custom format",
		}},
		{"v2 wildfire", version.Number{8, 0, 0, ""}, Entry{
			Name:     "t1",
			Wildfire: "custom format",
		}},
		{"v2 tunnel", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			Tunnel: "custom format",
		}},
		{"v2 userid", version.Number{8, 0, 0, ""}, Entry{
			Name:   "t1",
			UserId: "custom format",
		}},
		{"v2 gtp", version.Number{8, 0, 0, ""}, Entry{
			Name: "t1",
			Gtp:  "custom format",
		}},
		{"v2 auth", version.Number{8, 0, 0, ""}, Entry{
			Name: "t1",
			Auth: "custom format",
		}},
		{"v3 config", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
		}},
		{"v3 config with servers", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
			Servers: []Server{
				{
					Name:         "t1",
					Server:       "syslog.server.int",
					Transport:    TransportSsl,
					Port:         42,
					SyslogFormat: SyslogFormatIetf,
					Facility:     FacilityLocal4,
				},
			},
		}},
		{"v3 system", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			System: "custom format",
		}},
		{"v3 threat", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			Threat: "custom format",
		}},
		{"v3 traffic", version.Number{8, 1, 0, ""}, Entry{
			Name:    "t1",
			Traffic: "custom format",
		}},
		{"v3 hip match", version.Number{8, 1, 0, ""}, Entry{
			Name:     "t1",
			HipMatch: "custom format",
		}},
		{"v3 escaped characters", version.Number{8, 1, 0, ""}, Entry{
			Name:              "t1",
			EscapedCharacters: "abc",
		}},
		{"v3 escape character", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			EscapeCharacter: "x",
		}},
		{"v3 url", version.Number{8, 1, 0, ""}, Entry{
			Name: "t1",
			Url:  "custom format",
		}},
		{"v3 data", version.Number{8, 1, 0, ""}, Entry{
			Name: "t1",
			Data: "custom format",
		}},
		{"v3 wildfire", version.Number{8, 1, 0, ""}, Entry{
			Name:     "t1",
			Wildfire: "custom format",
		}},
		{"v3 tunnel", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			Tunnel: "custom format",
		}},
		{"v3 userid", version.Number{8, 1, 0, ""}, Entry{
			Name:   "t1",
			UserId: "custom format",
		}},
		{"v3 gtp", version.Number{8, 1, 0, ""}, Entry{
			Name: "t1",
			Gtp:  "custom format",
		}},
		{"v3 auth", version.Number{8, 1, 0, ""}, Entry{
			Name: "t1",
			Auth: "custom format",
		}},
		{"v3 sctp", version.Number{8, 1, 0, ""}, Entry{
			Name: "t1",
			Sctp: "custom format",
		}},
		{"v4 config", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
		}},
		{"v4 config with servers", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			Config: "custom format",
			Servers: []Server{
				{
					Name:         "t1",
					Server:       "syslog.server.int",
					Transport:    TransportSsl,
					Port:         42,
					SyslogFormat: SyslogFormatIetf,
					Facility:     FacilityLocal4,
				},
			},
		}},
		{"v4 system", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			System: "custom format",
		}},
		{"v4 threat", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			Threat: "custom format",
		}},
		{"v4 traffic", version.Number{9, 0, 0, ""}, Entry{
			Name:    "t1",
			Traffic: "custom format",
		}},
		{"v4 hip match", version.Number{9, 0, 0, ""}, Entry{
			Name:     "t1",
			HipMatch: "custom format",
		}},
		{"v4 escaped characters", version.Number{9, 0, 0, ""}, Entry{
			Name:              "t1",
			EscapedCharacters: "abc",
		}},
		{"v4 escape character", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			EscapeCharacter: "x",
		}},
		{"v4 url", version.Number{9, 0, 0, ""}, Entry{
			Name: "t1",
			Url:  "custom format",
		}},
		{"v4 data", version.Number{9, 0, 0, ""}, Entry{
			Name: "t1",
			Data: "custom format",
		}},
		{"v4 wildfire", version.Number{9, 0, 0, ""}, Entry{
			Name:     "t1",
			Wildfire: "custom format",
		}},
		{"v4 tunnel", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			Tunnel: "custom format",
		}},
		{"v4 userid", version.Number{9, 0, 0, ""}, Entry{
			Name:   "t1",
			UserId: "custom format",
		}},
		{"v4 gtp", version.Number{9, 0, 0, ""}, Entry{
			Name: "t1",
			Gtp:  "custom format",
		}},
		{"v4 auth", version.Number{9, 0, 0, ""}, Entry{
			Name: "t1",
			Auth: "custom format",
		}},
		{"v4 sctp", version.Number{9, 0, 0, ""}, Entry{
			Name: "t1",
			Sctp: "custom format",
		}},
		{"v4 iptag", version.Number{9, 0, 0, ""}, Entry{
			Name:  "t1",
			Iptag: "custom format",
		}},
	}
}
