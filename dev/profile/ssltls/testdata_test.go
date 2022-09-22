package ssltls

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
		{"v1 basic", version.Number{7, 0, 0, ""}, Entry{
			Name:        "basic",
			Certificate: "my cert",
			MinVersion:  Tls1_0,
			MaxVersion:  TlsMax,
		}},
		{"v2 yes", version.Number{8, 1, 0, ""}, Entry{
			Name:              "basic",
			Certificate:       "my cert",
			MinVersion:        Tls1_0,
			MaxVersion:        TlsMax,
			AllowAlgorithmRsa: true,
		}},
		{"v2 no", version.Number{8, 1, 0, ""}, Entry{
			Name:              "basic",
			Certificate:       "my cert",
			MinVersion:        Tls1_0,
			MaxVersion:        TlsMax,
			AllowAlgorithmRsa: false,
		}},
	}
}
