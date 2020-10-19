package wildfire

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
		{"v1 no rules", version.Number{7, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"v1 with rules", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t2",
			Description: "second",
			Rules: []Rule{
				Rule{
					Name:         "rule1",
					Applications: []string{"app1", "app2"},
					FileTypes:    []string{"7z"},
					Direction:    DirectionDownload,
					Analysis:     AnalysisPublicCloud,
				},
				Rule{
					Name:         "rule2",
					Applications: []string{"app3"},
					FileTypes:    []string{"ogg", "gif"},
					Direction:    DirectionUpload,
					Analysis:     AnalysisPrivateCloud,
				},
			},
		}},
	}
}
