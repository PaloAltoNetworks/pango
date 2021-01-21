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
			Name: "pdf",
			Properties: []Property{
				Property{
					Name:  "panav-rsp-pdf-dlp-title",
					Label: "Title",
				},
				Property{
					Name:  "panav-rsp-pdf-dlp-author",
					Label: "Author",
				},
			},
		}},
	}
}
