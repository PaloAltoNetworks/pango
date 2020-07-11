package arp

type tc struct {
	desc    string
	iType   string
	iName   string
	subName string
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"eth no sub", TypeEthernet, "ethernet1/5", "", Entry{
			Ip:         "10.1.1.1",
			MacAddress: "00:30:48:52:ab:c1",
		}},
		{"eth with sub", TypeEthernet, "ethernet1/5", "ethernet1/5.7", Entry{
			Ip:         "10.1.1.2",
			MacAddress: "00:30:48:52:ab:c2",
		}},
		{"agg no sub", TypeAggregate, "ae5", "", Entry{
			Ip:         "10.1.1.3",
			MacAddress: "00:30:48:52:ab:c3",
		}},
		{"agg with sub", TypeAggregate, "ae5", "ae5.4", Entry{
			Ip:         "10.1.1.4",
			MacAddress: "00:30:48:52:ab:c4",
		}},
		{"vlan", TypeVlan, "", "vlan.4", Entry{
			Ip:         "10.1.1.5",
			MacAddress: "00:30:48:52:ab:c5",
			Interface:  "vlan.2",
		}},
	}
}
