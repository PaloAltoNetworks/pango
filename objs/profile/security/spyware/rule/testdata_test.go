package rule

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
		{"v1 default", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t1",
			ThreatName:    "any",
			Category:      "my category",
			Severities:    []string{"high", "low"},
			PacketCapture: Disable,
			Action:        ActionDefault,
		}},
		{"v1 allow", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t1",
			PacketCapture: SinglePacket,
			Action:        ActionAllow,
		}},
		{"v1 alert", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t1",
			PacketCapture: ExtendedCapture,
			Action:        ActionAlert,
		}},
		{"v1 drop", version.Number{7, 0, 0, ""}, Entry{
			Name:   "t1",
			Action: ActionDrop,
		}},
		{"v1 reset client", version.Number{7, 0, 0, ""}, Entry{
			Name:   "t1",
			Action: ActionResetClient,
		}},
		{"v1 reset server", version.Number{7, 0, 0, ""}, Entry{
			Name:   "t1",
			Action: ActionResetServer,
		}},
		{"v1 reset both", version.Number{7, 0, 0, ""}, Entry{
			Name:   "t1",
			Action: ActionResetBoth,
		}},
		{"v1 block ip", version.Number{7, 0, 0, ""}, Entry{
			Name:            "t1",
			Action:          ActionBlockIp,
			BlockIpTrackBy:  TrackBySourceAndDestination,
			BlockIpDuration: 42,
		}},
	}
}
