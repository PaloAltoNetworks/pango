package filetype

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
		{"v1 no ident", version.Number{6, 1, 0, ""}, Entry{
			Name:          "pdf",
			Id:            1,
			FileTypeIdent: true,
			ThreatName:    "Adobe Portable Document Format (PDF)",
			FullName:      "Portable Document Format",
		}},
		{"v1 with ident", version.Number{6, 1, 0, ""}, Entry{
			Name:          "doc",
			Id:            5,
			DataIdent:     true,
			FileTypeIdent: true,
			ThreatName:    "Microsoft Word DOC",
			FullName:      "Microsoft Word 97-2004",
		}},
		{"v1 no file type ident or threat name", version.Number{6, 1, 0, ""}, Entry{
			Name:      "text/html",
			Id:        35,
			DataIdent: true,
			FullName:  "TEXT/HTML",
		}},
	}
}
