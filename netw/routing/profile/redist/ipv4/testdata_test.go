package ipv4

type testCase struct {
	desc string
	vr   string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"barebones redist", "v1", Entry{
			Name:     "one",
			Priority: 1,
			Action:   ActionRedist,
		}},
		{"barebones no redist", "v1", Entry{
			Name:     "two",
			Priority: 2,
			Action:   ActionNoRedist,
		}},
		{"redist static", "v1", Entry{
			Name:     "three",
			Priority: 3,
			Action:   ActionRedist,
			Types:    []string{TypeStatic},
		}},
		{"redist rip and connect", "v1", Entry{
			Name:     "four",
			Priority: 4,
			Action:   ActionRedist,
			Types:    []string{TypeRip, TypeConnect},
		}},
		{"no redist with ospf", "v1", Entry{
			Name:          "five",
			Priority:      5,
			Action:        ActionNoRedist,
			Types:         []string{TypeOspf},
			OspfPathTypes: []string{OspfPathTypeExt1, OspfPathTypeExt2},
			OspfAreas:     []string{"10.1.7.1", "10.1.7.2"},
			OspfTags:      []string{"10.1.7.3", "10.1.7.4"},
		}},
		{"no redist with bgp", "v1", Entry{
			Name:                   "six",
			Priority:               6,
			Action:                 ActionNoRedist,
			Types:                  []string{TypeBgp},
			BgpCommunities:         []string{"com1", "com2"},
			BgpExtendedCommunities: []string{"excom1", "excom2"},
		}},
	}
}
