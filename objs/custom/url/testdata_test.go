package url

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 basic", version.Number{7, 0, 0, ""}, Entry{
			Name:        "one",
			Description: "my description",
			Sites:       []string{"palo", "alto", "networks"},
		}},
		{"v2 basic", version.Number{9, 0, 0, ""}, Entry{
			Name:        "one",
			Description: "my description",
			Sites:       []string{"palo", "alto", "networks"},
			Type:        "URL List",
		}},
	}
}
