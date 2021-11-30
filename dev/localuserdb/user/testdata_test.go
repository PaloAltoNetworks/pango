package user

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
		{"v1 test disabled", version.Number{6, 1, 0, ""}, Entry{
			Name:         "one",
			PasswordHash: "secret",
			Disabled:     true,
		}},
		{"v1 test enabled", version.Number{6, 1, 0, ""}, Entry{
			Name:         "two",
			PasswordHash: "secret",
			Disabled:     false,
		}},
	}
}
