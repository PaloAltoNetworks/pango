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
			Name:             "t2",
			Description:      "second",
			EnableSyn:        true,
			SynAction:        SynActionRed,
			SynAlarmRate:     1001,
			SynActivateRate:  1002,
			SynMaxRate:       1003,
			SynBlockDuration: 300,
		}},
		{"v1 with syn cookies", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "t3",
			EnableSyn:             true,
			SynAction:             SynActionCookies,
			SynAlarmRate:          1001,
			SynActivateRate:       1002,
			SynMaxRate:            1003,
			MaxConcurrentSessions: 12345,
		}},
		{"v1 with udp red", version.Number{6, 1, 0, ""}, Entry{
			Name:             "t2",
			Description:      "second",
			EnableSyn:        true,
			UdpAlarmRate:     1001,
			UdpActivateRate:  1002,
			UdpMaxRate:       1003,
			UdpBlockDuration: 300,
		}},
		{"v1 with icmp red", version.Number{6, 1, 0, ""}, Entry{
			Name:              "t2",
			Description:       "second",
			EnableSyn:         true,
			IcmpAlarmRate:     1001,
			IcmpActivateRate:  1002,
			IcmpMaxRate:       1003,
			IcmpBlockDuration: 300,
		}},
		{"v1 with icmpv6 red", version.Number{6, 1, 0, ""}, Entry{
			Name:                "t2",
			Description:         "second",
			EnableSyn:           true,
			Icmpv6AlarmRate:     1001,
			Icmpv6ActivateRate:  1002,
			Icmpv6MaxRate:       1003,
			Icmpv6BlockDuration: 300,
		}},
		{"v1 with other red", version.Number{6, 1, 0, ""}, Entry{
			Name:               "t2",
			Description:        "second",
			EnableSyn:          true,
			OtherAlarmRate:     1001,
			OtherActivateRate:  1002,
			OtherMaxRate:       1003,
			OtherBlockDuration: 300,
		}},
	}
}
