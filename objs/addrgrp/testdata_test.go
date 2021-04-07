package addrgrp

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test static no tags", Entry{
			Name:            "one",
			Description:     "my description",
			StaticAddresses: []string{"adr1", "adr2"},
		}},
		{"test static with tags", Entry{
			Name:            "one",
			Description:     "my description",
			StaticAddresses: []string{"adr1", "adr2"},
			Tags:            []string{"tag1", "tag2"},
		}},
		{"test dynamic no tags", Entry{
			Name:         "one",
			Description:  "my description",
			DynamicMatch: "'tag1' or 'tag2' and 'tag3'",
		}},
		{"test dynamic with tags", Entry{
			Name:         "one",
			Description:  "my description",
			DynamicMatch: "'tag1' or 'tag2' and 'tag3'",
			Tags:         []string{"tag1", "tag2"},
		}},
	}
}
