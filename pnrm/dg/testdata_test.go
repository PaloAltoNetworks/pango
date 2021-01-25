package dg

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
		{"v1 basic", version.Number{6, 1, 0, ""}, Entry{
			Name:        "one",
			Description: "my description",
			Devices: map[string][]string{
				"01234": []string{"vsys1", "vsys2"},
				"234":   nil,
			},
		}},
		{"v1 with raw", version.Number{6, 1, 0, ""}, Entry{
			Name:        "two",
			Description: "raw fields",
			raw: map[string]string{
				"address":              "address",
				"addressGroup":         "addressGroup",
				"application":          "application",
				"applicationFilter":    "applicationFilter",
				"applicationGroup":     "applicationGroup",
				"applicationStatus":    "applicationStatus",
				"applicationTag":       "applicationTag",
				"authenticationObject": "authenticationObject",
				"authorizationCode":    "authorizationCode",
				"dynamicUserGroup":     "dynamicUserGroup",
				"emailScheduler":       "emailScheduler",
				"edl":                  "edl",
				"logSettings":          "logSettings",
				"masterDevice":         "masterDevice",
				"pdfSummaryReport":     "pdfSummaryReport",
				"postRulebase":         "postRulebase",
				"preRulebase":          "preRulebase",
				"profileGroup":         "profileGroup",
				"profiles":             "profiles",
				"region":               "region",
				"reportGroup":          "reportGroup",
				"reports":              "reports",
				"schedule":             "schedule",
				"service":              "service",
				"serviceGroup":         "serviceGroup",
				"tag":                  "tag",
				"threats":              "threats",
				"toSwVersion":          "toSwVersion",
			},
		}},
	}
}
