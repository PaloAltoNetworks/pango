package tags

type testCase struct {
	desc string
	vsys string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test with all fields", "", Entry{
			Name:    "one",
			Color:   "color1",
			Comment: "first test",
		}},
		{"test no color", "", Entry{
			Name:    "two",
			Comment: "second test",
		}},
		{"test no comment", "vsys1", Entry{
			Name:  "three",
			Color: "color3",
		}},
		{"test no color or comment", "vsys2", Entry{
			Name: "four",
		}},
	}
}
