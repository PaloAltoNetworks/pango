package andcond

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
		{"v1 no raw", version.Number{7, 1, 0, ""}, Entry{
			Name: "v1",
		}},
		{"v1 with raw", version.Number{7, 1, 0, ""}, Entry{
			Name: "v1",
			raw: map[string]string{
				"sigs": "or sig config",
			},
		}},
	}
}
