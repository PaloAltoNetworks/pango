package conadv

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"minimal", Entry{
			Name:   "one",
			Enable: false,
		}},
		{"standard", Entry{
			Name:   "two",
			Enable: true,
			UsedBy: []string{"one", "two"},
		}},
		{"with raw", Entry{
			Name:   "three",
			Enable: true,
			UsedBy: []string{"one", "two"},
			raw: map[string]string{
				"af": "advertisement filters",
				"nf": "non exist filters",
			},
		}},
	}
}
