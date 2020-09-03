package layer2

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	version    version.Number
	tmpl       string
	vsys       string
	importVsys string
	imports    []string
	conf       Entry
}

func getTests() []testCase {
	return []testCase{
		{version.Number{7, 1, 0, ""}, "tmpl1", "vsys1", "vsys1", []string{"ethernet1/1.2"}, Entry{
			Name:           "ethernet1/1.2",
			Tag:            2,
			NetflowProfile: "netflow profile",
			Comment:        "v1 basic",
		}},
	}
}
