package telemetry

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	return []testCase{
		{"all no", version.Number{8, 0, 0, ""}, Config{}},
		{"all yes", version.Number{8, 0, 0, ""}, Config{
			ApplicationReports:             true,
			ThreatPreventionReports:        true,
			UrlReports:                     true,
			FileTypeIdentificationReports:  true,
			ThreatPreventionData:           true,
			ThreatPreventionPacketCaptures: true,
			ProductUsageStats:              true,
			PassiveDnsMonitoring:           true,
		}},
		{"mix", version.Number{8, 0, 0, ""}, Config{
			ApplicationReports:   true,
			UrlReports:           true,
			ThreatPreventionData: true,
			ProductUsageStats:    true,
		}},
	}
}
