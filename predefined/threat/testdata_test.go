package threat

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	tt      string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 vulnerability basic", Vulnerability, version.Number{6, 1, 0, ""}, Entry{
			Name:       "12345",
			ThreatName: "Fake threat",
		}},
		{"v1 phone home basic", PhoneHome, version.Number{6, 1, 0, ""}, Entry{
			Name:       "12345",
			ThreatName: "Fake threat",
		}},
	}
}
