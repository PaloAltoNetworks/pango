package addr

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	vsys    string
	dg      string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 test ip netmask", "", "", version.Number{7, 1, 0, ""}, Entry{
			Name:        "one",
			Value:       "10.1.1.0/24",
			Type:        IpNetmask,
			Description: "my description",
			Tags:        []string{"tag1", "tag2"},
		}},
		{"v1 test ip range", "vsys2", "dg1", version.Number{7, 1, 0, ""}, Entry{
			Name:        "two",
			Value:       "10.1.1.1-10.1.1.254",
			Type:        IpRange,
			Description: "my description",
			Tags:        []string{"tag3", "tag4"},
		}},
		{"v1 test fqdn", "vsys3", "my device group", version.Number{7, 1, 0, ""}, Entry{
			Name:        "three",
			Value:       "example.com",
			Type:        Fqdn,
			Description: "my description",
		}},
		{"v2 test ip netmask", "", "", version.Number{9, 0, 0, ""}, Entry{
			Name:        "one",
			Value:       "10.1.1.0/24",
			Type:        IpNetmask,
			Description: "my description",
			Tags:        []string{"tag1", "tag2"},
		}},
		{"v2 test ip range", "vsys2", "dg2", version.Number{9, 0, 0, ""}, Entry{
			Name:        "two",
			Value:       "10.1.1.1-10.1.1.254",
			Type:        IpRange,
			Description: "my description",
			Tags:        []string{"tag3", "tag4"},
		}},
		{"v2 test fqdn", "vsys3", "my device group", version.Number{9, 0, 0, ""}, Entry{
			Name:        "three",
			Value:       "example.com",
			Type:        Fqdn,
			Description: "my description",
		}},
		{"v2 test ip wildcard", "vsys4", "dg4", version.Number{9, 0, 0, ""}, Entry{
			Name:        "four",
			Value:       "10.20.1.0/0.0.248.255",
			Type:        IpWildcard,
			Description: "my description",
		}},
	}
}
