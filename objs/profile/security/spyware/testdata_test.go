package spyware

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 only desc", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
		}},
		{"v1 with raw", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			raw: map[string]string{
				"rules":  "rules config",
				"threat": "threat exceptions",
			},
		}},
		{"v1 with packet capture", version.Number{8, 0, 0, ""}, Entry{
			Name:          "t1",
			Description:   "my desc",
			PacketCapture: SinglePacket,
		}},
		{"v1 with ipv4 sinkhole", version.Number{8, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv4Address: "127.0.0.1",
		}},
		{"v1 with ipv6 sinkhole", version.Number{8, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv6Address: "::1",
		}},
		{"v1 with threat exceptions", version.Number{8, 0, 0, ""}, Entry{
			Name:             "t1",
			Description:      "my desc",
			ThreatExceptions: []string{"threat1", "threat2"},
		}},
		{"v1 with botnet list", version.Number{8, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			BotnetLists: []BotnetList{
				BotnetList{
					Name:   "alert list",
					Action: ActionAlert,
				},
				BotnetList{
					Name:   "allow list",
					Action: ActionAllow,
				},
				BotnetList{
					Name:   "block list",
					Action: ActionBlock,
				},
				BotnetList{
					Name:   "sinkhole list",
					Action: ActionSinkhole,
				},
			},
		}},
		{"v2 only desc", version.Number{9, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
		}},
		{"v2 with raw", version.Number{9, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			raw: map[string]string{
				"rules":  "rules config",
				"threat": "threat exceptions",
			},
		}},
		{"v2 with ipv4 sinkhole", version.Number{9, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv4Address: "127.0.0.1",
		}},
		{"v2 with ipv6 sinkhole", version.Number{9, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv6Address: "::1",
		}},
		{"v2 with threat exceptions", version.Number{9, 0, 0, ""}, Entry{
			Name:             "t1",
			Description:      "my desc",
			ThreatExceptions: []string{"threat1", "threat2"},
		}},
		{"v2 with botnet list", version.Number{9, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			BotnetLists: []BotnetList{
				BotnetList{
					Name:          "alert list",
					Action:        ActionAlert,
					PacketCapture: Disable,
				},
				BotnetList{
					Name:          "allow list",
					Action:        ActionAllow,
					PacketCapture: SinglePacket,
				},
				BotnetList{
					Name:          "block list",
					Action:        ActionBlock,
					PacketCapture: ExtendedCapture,
				},
				BotnetList{
					Name:   "sinkhole list",
					Action: ActionSinkhole,
				},
			},
		}},
		{"v3 only desc", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
		}},
		{"v3 with raw", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			raw: map[string]string{
				"rules":  "rules config",
				"threat": "threat exceptions",
			},
		}},
		{"v3 with ipv4 sinkhole", version.Number{10, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv4Address: "127.0.0.1",
		}},
		{"v3 with ipv6 sinkhole", version.Number{10, 0, 0, ""}, Entry{
			Name:                "t1",
			Description:         "my desc",
			SinkholeIpv6Address: "::1",
		}},
		{"v3 with threat exceptions", version.Number{10, 0, 0, ""}, Entry{
			Name:             "t1",
			Description:      "my desc",
			ThreatExceptions: []string{"threat1", "threat2"},
		}},
		{"v3 with botnet list", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			BotnetLists: []BotnetList{
				BotnetList{
					Name:          "alert list",
					Action:        ActionAlert,
					PacketCapture: Disable,
				},
				BotnetList{
					Name:          "allow list",
					Action:        ActionAllow,
					PacketCapture: SinglePacket,
				},
				BotnetList{
					Name:          "block list",
					Action:        ActionBlock,
					PacketCapture: ExtendedCapture,
				},
				BotnetList{
					Name:   "sinkhole list",
					Action: ActionSinkhole,
				},
			},
		}},
		{"v3 with whitelists", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			WhiteLists: []WhiteList{
				WhiteList{
					Name:        "wl1",
					Description: "desc1",
				},
				WhiteList{
					Name:        "wl2",
					Description: "desc 2",
				},
			},
		}},
		{"v3 with dns categories", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "my desc",
			DnsCategories: []DnsCategory{
				DnsCategory{
					Name:          "cat1",
					Action:        ActionDefault,
					LogLevel:      LogLevelDefault,
					PacketCapture: SinglePacket,
				},
				DnsCategory{
					Name:          "cat2",
					Action:        ActionSinkhole,
					LogLevel:      LogLevelHigh,
					PacketCapture: Disable,
				},
			},
		}},
	}
}
