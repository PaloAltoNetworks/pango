package matchlist

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
		{"basic without raw", version.Number{8, 0, 0, ""}, Entry{
			Name:           "t1",
			Description:    "foobar",
			LogType:        LogTypeTunnel,
			Filter:         "my filter",
			SnmpProfiles:   []string{"prof1", "prof2"},
			SyslogProfiles: []string{"syslog1", "syslog2"},
		}},
		{"basic without raw", version.Number{8, 0, 0, ""}, Entry{
			Name:           "t1",
			LogType:        LogTypeAuth,
			Filter:         "my other filter",
			SendToPanorama: true,
			EmailProfiles:  []string{"email1", "email2"},
			HttpProfiles:   []string{"http1", "http3"},
			raw: map[string]string{
				"act": "my various actions",
			},
		}},
	}
}
