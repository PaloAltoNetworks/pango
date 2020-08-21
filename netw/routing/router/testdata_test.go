package router

type testCase struct {
	desc       string
	tmpl       string
	vsys       string
	doDefaults bool
	imports    []string
	conf       Entry
}

func getTests() []testCase {
	return []testCase{
		{"with defaults and import", "x", "vsys2", true, []string{"one"}, Entry{
			Name:       "one",
			Interfaces: []string{"ethernet1/1", "ethernet1/2"},
		}},
		{"no defaults or import", "", "", false, []string{"two"}, Entry{
			Name:       "two",
			Interfaces: []string{"ethernet1/3", "ethernet1/4"},
		}},
		{"with raw fields", "y", "vsys3", true, []string{"three"}, Entry{
			Name: "three",
			raw: map[string]string{
				"ecmp":      "<ecmp>raw ecmp</ecmp>",
				"multicast": "<multicast><raw>multicast</raw></multicast>",
				"protocol":  "<protocol><some><proto>field</proto></some></protocol>",
				"routing":   "<routing-table><route1>something</route1><route2>b</route2></routing-table>",
			},
		}},
	}
}
