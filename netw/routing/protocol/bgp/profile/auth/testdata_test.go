package auth

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"no secret", Entry{
			Name: "one",
		}},
		{"with secret", Entry{
			Name:   "two",
			Secret: "secret",
		}},
	}
}
