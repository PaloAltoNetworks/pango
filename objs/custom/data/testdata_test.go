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
		{"v1 predefined", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			Type:        TypePredefined,
			PredefinedPatterns: []PredefinedPattern{
				PredefinedPattern{
					Name:      "social-security-numbers",
					FileTypes: []string{"any"},
				},
			},
		}},
		{"v1 regex", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			Type:        TypeRegex,
			Regexes: []Regex{
				Regex{
					Name:      "pat1",
					FileTypes: []string{"any"},
					Regex:     "*\\.gif",
				},
			},
		}},
		{"v1 file properties", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			Type:        TypeFileProperties,
			FileProperties: []FileProperty{
				FileProperty{
					Name:          "pat1",
					FileType:      "xlsx",
					FileProperty:  "panav-rsp-office-dlp-category",
					PropertyValue: "John",
				},
			},
		}},
	}
}
