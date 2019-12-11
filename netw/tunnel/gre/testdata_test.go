package gre

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
		{"v1 ip no keep alive", version.Number{9, 0, 0, ""}, Entry{
			Name:              "v1",
			Interface:         "ethernet1/1",
			LocalAddressType:  LocalAddressTypeIp,
			LocalAddressValue: "192.168.1.1",
			PeerAddress:       "10.1.1.1",
			TunnelInterface:   "tunnel.1",
			Ttl:               50,
			CopyTos:           true,
			Disabled:          false,
		}},
		{"v1 floating ip no keep alive", version.Number{9, 0, 0, ""}, Entry{
			Name:              "v1",
			Interface:         "ethernet1/1",
			LocalAddressType:  LocalAddressTypeFloatingIp,
			LocalAddressValue: "192.168.1.1",
			PeerAddress:       "10.1.1.1",
			TunnelInterface:   "tunnel.1",
			Ttl:               60,
			CopyTos:           false,
			Disabled:          true,
		}},
		{"v1 ip with keep alive", version.Number{9, 0, 0, ""}, Entry{
			Name:               "v1",
			Interface:          "ethernet1/1",
			LocalAddressType:   LocalAddressTypeIp,
			LocalAddressValue:  "192.168.1.1",
			PeerAddress:        "10.1.1.1",
			TunnelInterface:    "tunnel.1",
			Ttl:                70,
			EnableKeepAlive:    true,
			KeepAliveInterval:  2,
			KeepAliveRetry:     3,
			KeepAliveHoldTimer: 4,
		}},
	}
}
