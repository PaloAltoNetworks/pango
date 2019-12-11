package peer

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		version version.Number
		conf    Entry
	}{
		{"v1", version.Number{6, 1, 0, ""}, Entry{
			Name:                          "v1",
			Enable:                        true,
			PeerAs:                        "peer as",
			LocalAddressInterface:         "ethernet1/1",
			LocalAddressIp:                "2.2.2.2",
			PeerAddressIp:                 "3.3.3.3",
			ReflectorClient:               ReflectorClientNonClient,
			PeeringType:                   PeeringTypeBilateral,
			MaxPrefixes:                   MaxPrefixesUnlimited,
			AuthProfile:                   "my auth profile",
			KeepAliveInterval:             30,
			MultiHop:                      2,
			OpenDelayTime:                 3,
			HoldTime:                      4,
			IdleHoldTime:                  5,
			AllowIncomingConnections:      true,
			IncomingConnectionsRemotePort: 1010,
			AllowOutgoingConnections:      true,
			OutgoingConnectionsLocalPort:  2020,
		}},
		{"v2", version.Number{7, 1, 0, ""}, Entry{
			Name:                          "v2",
			Enable:                        false,
			PeerAs:                        "peer as",
			LocalAddressInterface:         "ethernet1/2",
			LocalAddressIp:                "2.2.2.2",
			PeerAddressIp:                 "3.3.3.3",
			ReflectorClient:               ReflectorClientClient,
			PeeringType:                   PeeringTypeBilateral,
			MaxPrefixes:                   MaxPrefixesUnlimited,
			AuthProfile:                   "my auth profile",
			KeepAliveInterval:             30,
			MultiHop:                      2,
			OpenDelayTime:                 3,
			HoldTime:                      4,
			IdleHoldTime:                  5,
			AllowIncomingConnections:      true,
			IncomingConnectionsRemotePort: 1010,
			AllowOutgoingConnections:      true,
			OutgoingConnectionsLocalPort:  2020,
			BfdProfile:                    BfdProfileInherit,
		}},
		{"v3", version.Number{8, 0, 0, ""}, Entry{
			Name:                           "v3",
			Enable:                         true,
			PeerAs:                         "peer as",
			LocalAddressInterface:          "ethernet1/2",
			LocalAddressIp:                 "2.2.2.2",
			PeerAddressIp:                  "3.3.3.3",
			ReflectorClient:                ReflectorClientMeshedClient,
			PeeringType:                    PeeringTypeUnspecified,
			MaxPrefixes:                    MaxPrefixesUnlimited,
			AuthProfile:                    "my auth profile",
			KeepAliveInterval:              30,
			MultiHop:                       2,
			OpenDelayTime:                  3,
			HoldTime:                       4,
			IdleHoldTime:                   5,
			AllowIncomingConnections:       true,
			IncomingConnectionsRemotePort:  1010,
			AllowOutgoingConnections:       true,
			OutgoingConnectionsLocalPort:   2020,
			BfdProfile:                     BfdProfileNone,
			EnableMpBgp:                    true,
			AddressFamilyType:              AddressFamilyTypeIpv4,
			SubsequentAddressFamilyUnicast: true,
			EnableSenderSideLoopDetection:  true,
		}},
		{"v4", version.Number{8, 1, 0, ""}, Entry{
			Name:                             "v4",
			Enable:                           false,
			PeerAs:                           "peer as",
			LocalAddressInterface:            "ethernet1/3",
			LocalAddressIp:                   "2.2.2.2",
			PeerAddressIp:                    "3.3.3.3",
			ReflectorClient:                  ReflectorClientMeshedClient,
			PeeringType:                      PeeringTypeUnspecified,
			MaxPrefixes:                      "42",
			AuthProfile:                      "my auth profile",
			KeepAliveInterval:                30,
			MultiHop:                         2,
			OpenDelayTime:                    3,
			HoldTime:                         4,
			IdleHoldTime:                     5,
			AllowIncomingConnections:         false,
			IncomingConnectionsRemotePort:    1010,
			AllowOutgoingConnections:         false,
			OutgoingConnectionsLocalPort:     2020,
			BfdProfile:                       "my bfd profile",
			EnableMpBgp:                      false,
			AddressFamilyType:                AddressFamilyTypeIpv6,
			SubsequentAddressFamilyMulticast: true,
			EnableSenderSideLoopDetection:    false,
			MinRouteAdvertisementInterval:    77,
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoPeer{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set("tmpl", "", "vr", "pg", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("tmpl", "", "vr", "pg", tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
