package dos

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
		{"v1 blank", version.Number{7, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"v1 with resource", version.Number{6, 1, 0, ""}, Entry{
			Name:                      "t2",
			Description:               "second",
			EnableSessionsProtections: true,
			MaxConcurrentSessions:     42000,
		}},
		{"v1 with syn red", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Syn: &SynProtection{
				Enable:        true,
				Action:        SynActionRed,
				AlarmRate:     1001,
				ActivateRate:  1002,
				MaxRate:       1003,
				BlockDuration: 300,
			},
		}},
		{"v1 with syn cookies", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "t3",
			MaxConcurrentSessions: 12345,
			Syn: &SynProtection{
				Action:       SynActionCookies,
				AlarmRate:    1001,
				ActivateRate: 1002,
				MaxRate:      1003,
			},
		}},
		{"v1 with udp red", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Udp: &Protection{
				Enable:        true,
				AlarmRate:     1001,
				ActivateRate:  1002,
				MaxRate:       1003,
				BlockDuration: 300,
			},
		}},
		{"v1 with icmp red", version.Number{6, 1, 0, ""}, Entry{
			Name:                      "t2",
			Description:               "second",
			EnableSessionsProtections: true,
			Icmp: &Protection{
				Enable:        true,
				AlarmRate:     1001,
				ActivateRate:  1002,
				MaxRate:       1003,
				BlockDuration: 300,
			},
		}},
		{"v1 with icmpv6 red", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Icmpv6: &Protection{
				Enable:        true,
				AlarmRate:     1001,
				ActivateRate:  1002,
				MaxRate:       1003,
				BlockDuration: 300,
			},
		}},
		{"v1 with other red", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Other: &Protection{
				Enable:        true,
				AlarmRate:     1001,
				ActivateRate:  1002,
				MaxRate:       1003,
				BlockDuration: 300,
			},
		}},
	}
}
