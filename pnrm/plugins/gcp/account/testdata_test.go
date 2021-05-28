package account

import (
	"github.com/PaloAltoNetworks/pango/plugin"
)

type tc struct {
	desc   string
	plugin plugin.Info
	conf   Entry
}

func getTests() []tc {
	return []tc{
		{"basic project", plugin.Info{Name: "gcp", Installed: "yes", Version: "1.0.0"}, Entry{
			Name:                         "creds",
			ProjectId:                    "my projectid",
			Description:                  "some desc",
			ServiceAccountCredentialType: Project,
			CredentialFile:               "file contents",
		}},
	}
}
