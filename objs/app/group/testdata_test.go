package group

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
		{"v1 basic", version.Number{7, 1, 0, ""}, Entry{
			Name:         "group",
			Applications: []string{"app1", "app2"},
		}},
		{"v1 no apps", version.Number{7, 1, 0, ""}, Entry{
			Name: "group",
		}},
	}
}
