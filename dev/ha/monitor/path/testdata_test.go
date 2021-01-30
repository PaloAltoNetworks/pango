package path

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	gType   string
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 VirtualRouter 0 groups", version.Number{10, 0, 0, ""}, VirtualRouter,
			Entry{
				Name:             "t1",
				Enable:           true,
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
			}},
		{"v1 VirtualRouter 1 group", version.Number{10, 0, 0, ""}, VirtualRouter,
			Entry{
				Name:             "t2",
				Enable:           true,
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
				DstIpGroups: []DstIpGroup{
					DstIpGroup{
						Name:             "group1",
						Enable:           false,
						FailureCondition: FailureConditionAny,
						DstIps: []string{
							"ip 1",
						},
					},
				},
			}},
		{"v1 VirtualRouter 2 groups", version.Number{10, 0, 0, ""}, VirtualRouter,
			Entry{
				Name:             "t3",
				Enable:           true,
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
				DstIpGroups: []DstIpGroup{
					DstIpGroup{
						Name:             "group1",
						Enable:           false,
						FailureCondition: FailureConditionAny,
						DstIps: []string{
							"ip 1",
						},
					},
					DstIpGroup{
						Name:   "group2",
						Enable: true,
						DstIps: []string{
							"ip 1",
							"ip 2",
						},
					},
				},
			}},
		{"v1 Vlan 2 groups", version.Number{10, 0, 0, ""}, Vlan,
			Entry{
				Name:             "t4",
				Enable:           true,
				SrcIp:            "src ip",
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
				DstIpGroups: []DstIpGroup{
					DstIpGroup{
						Name:             "group1",
						Enable:           false,
						FailureCondition: FailureConditionAny,
						DstIps: []string{
							"ip 1",
						},
					},
					DstIpGroup{
						Name:   "group2",
						Enable: true,
						DstIps: []string{
							"ip 1",
							"ip 2",
						},
					},
				},
			}},
		{"v1 VirtualWire 2 groups", version.Number{10, 0, 0, ""}, VirtualWire,
			Entry{
				Name:             "t5",
				Enable:           true,
				SrcIp:            "src ip",
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
				DstIpGroups: []DstIpGroup{
					DstIpGroup{
						Name:             "group1",
						Enable:           false,
						FailureCondition: FailureConditionAny,
						DstIps: []string{
							"ip 1",
						},
					},
					DstIpGroup{
						Name:   "group2",
						Enable: true,
						DstIps: []string{
							"ip 1",
							"ip 2",
						},
					},
				},
			}},
		{"v1 LogicalRouter 2 groups", version.Number{10, 0, 0, ""}, LogicalRouter,
			Entry{
				Name:             "t6",
				Enable:           true,
				FailureCondition: FailureConditionAll,
				PingInterval:     50000,
				PingCount:        3,
				DstIpGroups: []DstIpGroup{
					DstIpGroup{
						Name:             "group1",
						Enable:           false,
						FailureCondition: FailureConditionAny,
						DstIps: []string{
							"ip 1",
						},
					},
					DstIpGroup{
						Name:   "group2",
						Enable: true,
						DstIps: []string{
							"ip 1",
							"ip 2",
						},
					},
				},
			}},
	}
}
