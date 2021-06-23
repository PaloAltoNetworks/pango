package vsys

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 empty", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
		}},
		{"v1 interfaces", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				Interfaces: []string{"ethernet1/1", "ethernet1/2"},
			},
		}},
		{"v1 virtual routers", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				VirtualRouters: []string{"default", "vr2"},
			},
		}},
		{"v1 virtual wires", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				VirtualWires: []string{"vwire1"},
			},
		}},
		{"v1 vlans", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				Vlans: []string{"vlan1", "vlan3", "vlan2"},
			},
		}},
		{"v2 empty", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
		}},
		{"v2 interfaces", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				Interfaces: []string{"ethernet1/1", "ethernet1/2"},
			},
		}},
		{"v2 virtual routers", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				VirtualRouters: []string{"default", "vr2"},
			},
		}},
		{"v2 virtual wires", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				VirtualWires: []string{"vwire1"},
			},
		}},
		{"v2 vlans", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				Vlans: []string{"vlan1", "vlan3", "vlan2"},
			},
		}},
		{"v2 logical routers", version.Number{10, 0, 0, ""}, Entry{
			Name: "one",
			NetworkImports: &NetworkImports{
				LogicalRouters: []string{"one", "two"},
			},
		}},
	}
}
