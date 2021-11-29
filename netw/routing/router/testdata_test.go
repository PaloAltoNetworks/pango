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
			Name:           "two",
			Interfaces:     []string{"ethernet1/3", "ethernet1/4"},
			StaticDist:     1,
			StaticIpv6Dist: 2,
			OspfIntDist:    3,
			OspfExtDist:    4,
			Ospfv3IntDist:  5,
			Ospfv3ExtDist:  6,
			IbgpDist:       7,
			EbgpDist:       8,
			RipDist:        9,
		}},
		{"with raw fields", "y", "vsys3", true, []string{"foo"}, Entry{
			Name: "foo",
			raw: map[string]string{
				"multicast": "<multicast><raw>multicast</raw></multicast>",
				"protocol":  "<protocol><some><proto>field</proto></some></protocol>",
				"routing":   "<routing-table><route1>something</route1><route2>b</route2></routing-table>",
			},
		}},
		{"ecmp ip modulo", "x", "vsys1", false, []string{"ecmp1"}, Entry{
			Name:                  "ecmp1",
			EnableEcmp:            true,
			EcmpMaxPath:           7,
			EcmpLoadBalanceMethod: EcmpLoadBalanceMethodIpModulo,
			StaticDist:            1,
			StaticIpv6Dist:        2,
			OspfIntDist:           3,
			OspfExtDist:           4,
			Ospfv3IntDist:         5,
			Ospfv3ExtDist:         6,
			IbgpDist:              7,
			EbgpDist:              8,
			RipDist:               9,
			raw: map[string]string{
				"multicast": "<multicast><raw>multicast</raw></multicast>",
			},
		}},
		{"ecmp ip hash", "x", "vsys1", false, []string{"ecmp2"}, Entry{
			Name:                  "ecmp2",
			EcmpSymmetricReturn:   true,
			EcmpMaxPath:           8,
			EcmpLoadBalanceMethod: EcmpLoadBalanceMethodIpHash,
			EcmpHashSourceOnly:    true,
			EcmpHashUsePort:       true,
			EcmpHashSeed:          3,
			StaticDist:            1,
			StaticIpv6Dist:        2,
			OspfIntDist:           3,
			OspfExtDist:           4,
			Ospfv3IntDist:         5,
			Ospfv3ExtDist:         6,
			IbgpDist:              7,
			EbgpDist:              8,
			RipDist:               9,
			raw: map[string]string{
				"protocol": "<protocol><some><proto>field</proto></some></protocol>",
			},
		}},
		{"ecmp wrr", "x", "vsys1", false, []string{"ecmp3"}, Entry{
			Name:                  "ecmp3",
			EcmpStrictSourcePath:  true,
			EcmpMaxPath:           9,
			EcmpLoadBalanceMethod: EcmpLoadBalanceMethodWeightedRoundRobin,
			EcmpWeightedRoundRobinInterfaces: map[string]int{
				"ethernet1/1": 1,
				"ethernet1/2": 3,
				"ethernet1/5": 8,
			},
			StaticDist:     1,
			StaticIpv6Dist: 2,
			OspfIntDist:    3,
			OspfExtDist:    4,
			Ospfv3IntDist:  5,
			Ospfv3ExtDist:  6,
			IbgpDist:       7,
			EbgpDist:       8,
			RipDist:        9,
			raw: map[string]string{
				"routing": "<routing-table><route1>something</route1><route2>b</route2></routing-table>",
			},
		}},
		{"ecmp brr", "x", "vsys1", false, []string{"ecmp4"}, Entry{
			Name:                  "ecmp4",
			EnableEcmp:            true,
			EcmpSymmetricReturn:   true,
			EcmpStrictSourcePath:  true,
			EcmpMaxPath:           7,
			EcmpLoadBalanceMethod: EcmpLoadBalanceMethodBalancedRoundRobin,
			StaticDist:            1,
			StaticIpv6Dist:        2,
			OspfIntDist:           3,
			OspfExtDist:           4,
			Ospfv3IntDist:         5,
			Ospfv3ExtDist:         6,
			IbgpDist:              7,
			EbgpDist:              8,
			RipDist:               9,
			raw: map[string]string{
				"multicast": "<multicast><raw>multicast</raw></multicast>",
				"protocol":  "<protocol><some><proto>field</proto></some></protocol>",
				"routing":   "<routing-table><route1>something</route1><route2>b</route2></routing-table>",
			},
		}},
	}
}
