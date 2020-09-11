package zone

type testCase struct {
	desc string
	vsys string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"empty zone", "", Entry{
			Name: "one",
			Mode: "layer3",
		}},
		{"layer3 zone", "vsys1", Entry{
			Name:        "two",
			Mode:        "layer3",
			Interfaces:  []string{"ethernet1/1", "ethernet1/2"},
			ZoneProfile: "profile1",
			LogSetting:  "setting1",
			IncludeAcls: []string{"10.1.2.0/24"},
		}},
		{"layer2 zone", "vsys2", Entry{
			Name:        "three",
			Mode:        "layer2",
			Interfaces:  []string{"ethernet1/3", "ethernet1/4"},
			ExcludeAcls: []string{"10.100.1.0/24"},
		}},
		{"vwire zone", "vsys3", Entry{
			Name:        "four",
			Mode:        "virtual-wire",
			Interfaces:  []string{"ethernet1/5", "ethernet1/6"},
			IncludeAcls: []string{"10.1.3.0/24"},
		}},
		{"tap zone", "vsys4", Entry{
			Name:        "five",
			Mode:        "external",
			Interfaces:  []string{"ethernet1/7", "ethernet1/8"},
			ExcludeAcls: []string{"10.100.2.0/24"},
		}},
	}
}
