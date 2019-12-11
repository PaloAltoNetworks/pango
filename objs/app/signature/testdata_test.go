package signature

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
		{"v1 no raw no order", version.Number{7, 1, 0, ""}, Entry{
			Name:      "v1",
			Comment:   "my comment",
			Scope:     ScopeSession,
			OrderFree: false,
		}},
		{"v1 raw no order", version.Number{7, 1, 0, ""}, Entry{
			Name:      "v1",
			Comment:   "my comment",
			Scope:     ScopeSession,
			OrderFree: false,
			raw: map[string]string{
				"sigs": "and sig config",
			},
		}},
		{"v1 no raw ordered", version.Number{7, 1, 0, ""}, Entry{
			Name:      "v1",
			Scope:     ScopeTransaction,
			OrderFree: true,
		}},
	}
}
