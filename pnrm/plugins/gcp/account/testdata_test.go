package account

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
		{"basic project", version.Number{7, 1, 0, ""}, Entry{
			Name:                         "creds",
			ProjectId:                    "my projectid",
			Description:                  "some desc",
			ServiceAccountCredentialType: Project,
			CredentialFile:               "file contents",
		}},
	}
}
