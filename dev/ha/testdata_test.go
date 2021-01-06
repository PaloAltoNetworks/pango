package ha

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	return []testCase{
		{"v1 ha active passive no raw", version.Number{10, 0, 0, ""}, Config{
			Enable:                 true,
			GroupId:                3,
			Description:            "description",
			Mode:                   ModeActivePassive,
			PeerHa1IpAddress:       "peer ha1 ip",
			BackupPeerHa1IpAddress: "backup peer ha1 ip",
			ConfigSyncEnable:       true,

			Ha1: &Ha1Interface{
				Port:             "ha1 port",
				IpAddress:        "ha1 ip",
				Netmask:          "ha1 netmask",
				Gateway:          "ha1 gateway",
				EncryptionEnable: false,
				MonitorHoldTime:  1000,
			},
			Ha1Backup: &Ha1BackupInterface{
				Port:      "ha1 backup port",
				IpAddress: "ha1 backup ip",
				Netmask:   "ha1 backup netmask",
				Gateway:   "ha1 backup gateway",
			},
			Ha2: &Ha2Interface{
				Port:      "ha2 port",
				IpAddress: "ha2 ip",
				Netmask:   "ha2 netmask",
				Gateway:   "ha2 gateway",
			},
			Ha2Backup: &Ha2Interface{
				Port:      "ha2 backup port",
				IpAddress: "ha2 backup ip",
				Netmask:   "ha2 backup netmask",
				Gateway:   "ha2 backup gateway",
			},
			Ha3: &Ha3Interface{
				Port: "ha3 port",
			},
			Ha4: &Ha4Interface{
				Port:      "ha4 port",
				IpAddress: "ha4 ip",
				Netmask:   "ha4 netmask",
			},
			Ha4Backup: &Ha4Interface{
				Port:      "ha4 backup port",
				IpAddress: "ha4 backup ip",
				Netmask:   "ha4 backup netmask",
			},

			ApPassiveLinkState:        ApPassiveLinkStateAuto,
			ApMonitorFailHoldDownTime: 60,

			ElectionDevicePriority:  "100",
			ElectionPreemptive:      true,
			ElectionHeartBeatBackup: true,

			ElectionTimersMode:                          ElectionTimersModeAdvanced,
			ElectionTimersAdvPromotionHoldTime:          "1",
			ElectionTimersAdvHelloInterval:              8000,
			ElectionTimersAdvHeartBeatInterval:          1000,
			ElectionTimersAdvFlapMax:                    "1",
			ElectionTimersAdvPreemptionHoldTime:         1,
			ElectionTimersAdvMonitorFailHoldUpTime:      "0",
			ElectionTimersAdvAdditionalMasterHoldUpTime: "0",

			Ha2StateSyncEnable:             true,
			Ha2StateSyncTransport:          Ha2StateSyncTransportUdp,
			Ha2StateSyncKeepAliveEnable:    true,
			Ha2StateSyncKeepAliveAction:    Ha2StateSyncKeepAliveActionLogOnly,
			Ha2StateSyncKeepAliveThreshold: 5000,

			LinkMonitorEnable:           true,
			LinkMonitorFailureCondition: LinkMonitorFailureConditionAny,
		}},
		{"v1 ha active passive with raw", version.Number{10, 0, 0, ""}, Config{
			Enable:                 true,
			GroupId:                3,
			Description:            "description",
			Mode:                   ModeActivePassive,
			PeerHa1IpAddress:       "peer ha1 ip",
			BackupPeerHa1IpAddress: "backup peer ha1 ip",
			ConfigSyncEnable:       true,
			Ha1: &Ha1Interface{
				Port:             "ha1 port",
				IpAddress:        "ha1 ip",
				Netmask:          "ha1 netmask",
				Gateway:          "ha1 gateway",
				EncryptionEnable: false,
				MonitorHoldTime:  1000,
			},
			Ha1Backup: &Ha1BackupInterface{
				Port:      "ha1 backup port",
				IpAddress: "ha1 backup ip",
				Netmask:   "ha1 backup netmask",
				Gateway:   "ha1 backup gateway",
			},
			Ha2: &Ha2Interface{
				Port:      "ha2 port",
				IpAddress: "ha2 ip",
				Netmask:   "ha2 netmask",
				Gateway:   "ha2 gateway",
			},
			Ha2Backup: &Ha2Interface{
				Port:      "ha2 backup port",
				IpAddress: "ha2 backup ip",
				Netmask:   "ha2 backup netmask",
				Gateway:   "ha2 backup gateway",
			},
			Ha3: &Ha3Interface{
				Port: "ha3 port",
			},
			Ha4: &Ha4Interface{
				Port:      "ha4 port",
				IpAddress: "ha4 ip",
				Netmask:   "ha4 netmask",
			},
			Ha4Backup: &Ha4Interface{
				Port:      "ha4 backup port",
				IpAddress: "ha4 backup ip",
				Netmask:   "ha4 backup netmask",
			},

			ApPassiveLinkState:        ApPassiveLinkStateAuto,
			ApMonitorFailHoldDownTime: 60,

			ElectionDevicePriority:  "100",
			ElectionPreemptive:      true,
			ElectionHeartBeatBackup: true,

			ElectionTimersMode:                          ElectionTimersModeAdvanced,
			ElectionTimersAdvPromotionHoldTime:          "1",
			ElectionTimersAdvHelloInterval:              8000,
			ElectionTimersAdvHeartBeatInterval:          1000,
			ElectionTimersAdvFlapMax:                    "1",
			ElectionTimersAdvPreemptionHoldTime:         1,
			ElectionTimersAdvMonitorFailHoldUpTime:      "0",
			ElectionTimersAdvAdditionalMasterHoldUpTime: "0",

			Ha2StateSyncEnable:             true,
			Ha2StateSyncTransport:          Ha2StateSyncTransportUdp,
			Ha2StateSyncKeepAliveEnable:    true,
			Ha2StateSyncKeepAliveAction:    Ha2StateSyncKeepAliveActionLogOnly,
			Ha2StateSyncKeepAliveThreshold: 5000,

			LinkMonitorEnable:           true,
			LinkMonitorFailureCondition: LinkMonitorFailureConditionAny,

			raw: map[string]string{
				"linkgroup":   "link group",
				"cluster":     "cluster",
				"pathmonitor": "path monitor",
			},
		}},
	}
}
