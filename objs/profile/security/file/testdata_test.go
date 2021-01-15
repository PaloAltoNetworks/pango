package file

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
					Action:       ActionAlert,
				},
				Rule{
					Name:         "rule2",
					Applications: []string{"app1"},
					FileTypes:    []string{"mpeg2", "mpeg3", "mpeg4"},
					Direction:    DirectionUpload,
					Action:       ActionBlock,
				},
				Rule{
					Name:         "rule3",
					Applications: []string{"app3"},
					FileTypes:    []string{"mp3", "ogg"},
					Direction:    DirectionBoth,
					Action:       ActionContinue,
				},
				Rule{
					Name:         "rule4",
					Applications: []string{"app4", "app5"},
					FileTypes:    []string{"tar", "gz"},
					Direction:    DirectionBoth,
					Action:       ActionForward,
				},
				Rule{
					Name:         "rule5",
					Applications: []string{"app4", "app5"},
					FileTypes:    []string{"jpeg", "jpg"},
					Direction:    DirectionUpload,
					Action:       ActionContinueAndForward,
				},
			},
		}},
	}
}
