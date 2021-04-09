package srvcgrp

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test no services", Entry{
			Name: "one",
			Tags: []string{"one", "two"},
		}},
		{"test one service", Entry{
			Name:     "two",
			Services: []string{"svc1"},
			Tags:     []string{"single"},
		}},
		{"test two services", Entry{
			Name:     "three",
			Services: []string{"svc1", "svc2"},
		}},
	}
}
