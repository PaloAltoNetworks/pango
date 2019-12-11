package monitor

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
		{"v1 empty", version.Number{7, 1, 0, ""}, Entry{
			Name: "v1",
		}},
		{"v1 basic", version.Number{7, 1, 0, ""}, Entry{
			Name:      "v1",
			Interval:  50,
			Threshold: 7,
			Action:    ActionWaitRecover,
		}},
	}
}
