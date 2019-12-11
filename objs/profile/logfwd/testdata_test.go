package logfwd

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
		{"version1 basic", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"version1 with raw data", version.Number{8, 0, 0, ""}, Entry{
			Name: "t2",
			raw: map[string]string{
				"ml": "match list data",
			},
		}},
		{"version2 basic no enhanced logging", version.Number{8, 1, 0, ""}, Entry{
			Name:        "t3",
			Description: "foobar",
		}},
		{"version2 basic enhanced logging", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t4",
			Description:     "foobar",
			EnhancedLogging: true,
		}},
		{"version2 with raw data", version.Number{8, 1, 0, ""}, Entry{
			Name: "t5",
			raw: map[string]string{
				"ml": "match list data",
			},
		}},
	}
}
