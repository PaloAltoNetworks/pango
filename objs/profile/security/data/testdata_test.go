package data

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
		{"v1 no rules", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			DataCapture: true,
		}},
		{"v1 with rules", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Rules: []Rule{
				Rule{
					Name:           "rule0",
					Applications:   []string{"aim-mail"},
					FileTypes:      []string{"text/html", "rtf"},
					Direction:      DirectionDownload,
					AlertThreshold: 42,
					BlockThreshold: 53,
					DataPattern:    "myObj",
				},
			},
		}},
		{"v2 no rules", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"v2 with rules", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			DataCapture: true,
			Rules: []Rule{
				Rule{
					Name:           "rule0",
					Applications:   []string{"aim-mail"},
					FileTypes:      []string{"text/html", "rtf"},
					Direction:      DirectionDownload,
					AlertThreshold: 42,
					BlockThreshold: 53,
					DataPattern:    "myObj",
					LogSeverity:    "high",
				},
			},
		}},
	}
}
