package server

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
		{"v1 basic", version.Number{7, 1, 0, ""}, Entry{
			Name:         "t1",
			DisplayName:  "foo",
			From:         "src@example.com",
			To:           "dst@example.com",
			AlsoTo:       "also@example.com",
			EmailGateway: "smtp.example.com",
		}},
	}
}
