package vlan

type testCase struct {
	desc    string
	tmpl    string
	vsys    string
	imports []string
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"empty vlan", "", "", []string{"one"}, Entry{
			Name: "one",
		}},
		{"full vlan spec", "x", "vsys2", []string{"two"}, Entry{
			Name:          "two",
			VlanInterface: "ethernet1/3",
			Interfaces:    []string{"ethernet1/1", "ethernet1/2"},
			StaticMacs: map[string]string{
				"00:30:48:52:ab:cd": "ethernet1/1",
				"00:30:48:52:11:22": "ethernet1/2",
			},
		}},
	}
}
