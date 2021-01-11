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
		{"v1 with rules", version.Number{8, 0, 0, ""}, Entry{
			Name:        "with rules",
			Description: "blah",
			Rules: []Rule{
				Rule{
					Name:          "t1",
					ThreatName:    "any",
					Category:      "my category",
					Severities:    []string{"high", "low"},
					PacketCapture: Disable,
					Action:        ActionDefault,
				},
				Rule{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
				},
				Rule{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Rule{
					Name:   "t1",
					Action: ActionDrop,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Rule{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
				},
			},
		}},
		{"v1 with exceptions", version.Number{8, 0, 0, ""}, Entry{
			Name:        "with exceptions",
			Description: "with exceptions",
			Exceptions: []Exception{
				Exception{
					Name:          "t1",
					PacketCapture: Disable,
					Action:        ActionDefault,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Exception{
					Name:   "t1",
					Action: ActionDrop,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Exception{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
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
		{"v2 with rules", version.Number{9, 0, 0, ""}, Entry{
			Name:        "with rules",
			Description: "blah",
			Rules: []Rule{
				Rule{
					Name:          "t1",
					ThreatName:    "any",
					Category:      "my category",
					Severities:    []string{"high", "low"},
					PacketCapture: Disable,
					Action:        ActionDefault,
				},
				Rule{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
				},
				Rule{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Rule{
					Name:   "t1",
					Action: ActionDrop,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Rule{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
				},
			},
		}},
		{"v2 with exceptions", version.Number{9, 0, 0, ""}, Entry{
			Name:        "with exceptions",
			Description: "with exceptions",
			Exceptions: []Exception{
				Exception{
					Name:          "t1",
					PacketCapture: Disable,
					Action:        ActionDefault,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Exception{
					Name:   "t1",
					Action: ActionDrop,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Exception{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
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
		{"v3 with rules", version.Number{10, 0, 0, ""}, Entry{
			Name:        "with rules",
			Description: "blah",
			Rules: []Rule{
				Rule{
					Name:          "t1",
					ThreatName:    "any",
					Category:      "my category",
					Severities:    []string{"high", "low"},
					PacketCapture: Disable,
					Action:        ActionDefault,
				},
				Rule{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
				},
				Rule{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Rule{
					Name:   "t1",
					Action: ActionDrop,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Rule{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Rule{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
				},
			},
		}},
		{"v3 with exceptions", version.Number{10, 0, 0, ""}, Entry{
			Name:        "with exceptions",
			Description: "with exceptions",
			Exceptions: []Exception{
				Exception{
					Name:          "t1",
					PacketCapture: Disable,
					Action:        ActionDefault,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: SinglePacket,
					Action:        ActionAllow,
					ExemptIps:     []string{"192.168.50.1", "192.168.50.2"},
				},
				Exception{
					Name:          "t1",
					PacketCapture: ExtendedCapture,
					Action:        ActionAlert,
				},
				Exception{
					Name:   "t1",
					Action: ActionDrop,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetClient,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetServer,
				},
				Exception{
					Name:   "t1",
					Action: ActionResetBoth,
				},
				Exception{
					Name:            "t1",
					Action:          ActionBlockIp,
					BlockIpTrackBy:  TrackBySourceAndDestination,
					BlockIpDuration: 42,
				},
			},
		}},
	}
}
