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
		{"v1 first half", version.Number{6, 0, 0, ""}, Entry{
			Name:                 "foo",
			AntivirusProfile:     "av",
			AntiSpywareProfile:   "spyware",
			VulnerabilityProfile: "vuln",
		}},
		{"v1 second half", version.Number{6, 0, 0, ""}, Entry{
			Name:                 "foo",
			UrlFilteringProfile:  "url",
			FileBlockingProfile:  "file",
			DataFilteringProfile: "data",
		}},
		{"v2 first half", version.Number{7, 0, 0, ""}, Entry{
			Name:                 "foo",
			AntivirusProfile:     "av",
			AntiSpywareProfile:   "spyware",
			VulnerabilityProfile: "vuln",
		}},
		{"v2 second half", version.Number{7, 0, 0, ""}, Entry{
			Name:                 "foo",
			UrlFilteringProfile:  "url",
			FileBlockingProfile:  "file",
			DataFilteringProfile: "data",
		}},
		{"v2 wildfire", version.Number{7, 0, 0, ""}, Entry{
			Name:                    "foo",
			WildfireAnalysisProfile: "wildfire",
		}},
		{"v3 first half", version.Number{8, 0, 0, ""}, Entry{
			Name:                 "foo",
			AntivirusProfile:     "av",
			AntiSpywareProfile:   "spyware",
			VulnerabilityProfile: "vuln",
		}},
		{"v3 second half", version.Number{8, 0, 0, ""}, Entry{
			Name:                 "foo",
			UrlFilteringProfile:  "url",
			FileBlockingProfile:  "file",
			DataFilteringProfile: "data",
		}},
		{"v3 wildfire", version.Number{8, 0, 0, ""}, Entry{
			Name:                    "foo",
			WildfireAnalysisProfile: "wildfire",
		}},
		{"v3 gtp", version.Number{8, 0, 0, ""}, Entry{
			Name:       "foo",
			GtpProfile: "wildfire",
		}},
		{"v4 first half", version.Number{9, 0, 0, ""}, Entry{
			Name:                 "foo",
			AntivirusProfile:     "av",
			AntiSpywareProfile:   "spyware",
			VulnerabilityProfile: "vuln",
		}},
		{"v4 second half", version.Number{9, 0, 0, ""}, Entry{
			Name:                 "foo",
			UrlFilteringProfile:  "url",
			FileBlockingProfile:  "file",
			DataFilteringProfile: "data",
		}},
		{"v4 wildfire", version.Number{9, 0, 0, ""}, Entry{
			Name:                    "foo",
			WildfireAnalysisProfile: "wildfire",
		}},
		{"v4 gtp", version.Number{9, 0, 0, ""}, Entry{
			Name:       "foo",
			GtpProfile: "wildfire",
		}},
		{"v4 sctp", version.Number{9, 0, 0, ""}, Entry{
			Name:        "foo",
			SctpProfile: "wildfire",
		}},
	}
}
