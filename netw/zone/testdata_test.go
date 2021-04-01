package zone

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	vsys    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 empty zone", "", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			Mode: ModeL3,
		}},
		{"v1 layer3 zone", "vsys1", version.Number{6, 1, 0, ""}, Entry{
			Name:         "two",
			Mode:         ModeL3,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/1", "ethernet1/2"},
			ZoneProfile:  "profile1",
			LogSetting:   "setting1",
			IncludeAcls:  []string{"10.1.2.0/24"},
		}},
		{"v1 layer2 zone", "vsys2", version.Number{6, 1, 0, ""}, Entry{
			Name:        "three",
			Mode:        ModeL2,
			Interfaces:  []string{"ethernet1/3", "ethernet1/4"},
			ExcludeAcls: []string{"10.100.1.0/24"},
		}},
		{"v1 vwire zone", "vsys3", version.Number{6, 1, 0, ""}, Entry{
			Name:         "four",
			Mode:         ModeVirtualWire,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/5", "ethernet1/6"},
			IncludeAcls:  []string{"10.1.3.0/24"},
		}},
		{"v1 external zone", "vsys4", version.Number{6, 1, 0, ""}, Entry{
			Name:        "five",
			Mode:        ModeExternal,
			Interfaces:  []string{"ext1", "ext2"},
			ExcludeAcls: []string{"10.100.2.0/24"},
		}},
		{"v1 tap zone", "vsys1", version.Number{6, 1, 0, ""}, Entry{
			Name:         "six",
			Mode:         ModeTap,
			EnableUserId: true,
			Interfaces:   []string{"int1", "int2"},
		}},
		{"v2 empty zone", "", version.Number{8, 0, 0, ""}, Entry{
			Name:                         "one",
			Mode:                         ModeL3,
			EnablePacketBufferProtection: true,
		}},
		{"v2 layer3 zone", "vsys1", version.Number{8, 0, 0, ""}, Entry{
			Name:         "two",
			Mode:         ModeL3,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/1", "ethernet1/2"},
			ZoneProfile:  "profile1",
			LogSetting:   "setting1",
			IncludeAcls:  []string{"10.1.2.0/24"},
		}},
		{"v2 layer2 zone", "vsys2", version.Number{8, 0, 0, ""}, Entry{
			Name:                         "three",
			Mode:                         ModeL2,
			EnablePacketBufferProtection: true,
			Interfaces:                   []string{"ethernet1/3", "ethernet1/4"},
			ExcludeAcls:                  []string{"10.100.1.0/24"},
		}},
		{"v2 vwire zone", "vsys3", version.Number{8, 0, 0, ""}, Entry{
			Name:         "four",
			Mode:         ModeVirtualWire,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/5", "ethernet1/6"},
			IncludeAcls:  []string{"10.1.3.0/24"},
		}},
		{"v2 external zone", "vsys4", version.Number{8, 0, 0, ""}, Entry{
			Name:                         "five",
			Mode:                         ModeExternal,
			EnablePacketBufferProtection: true,
			Interfaces:                   []string{"ext1", "ext2"},
			ExcludeAcls:                  []string{"10.100.2.0/24"},
		}},
		{"v2 tap zone", "vsys1", version.Number{8, 0, 0, ""}, Entry{
			Name:         "six",
			Mode:         ModeTap,
			EnableUserId: true,
			Interfaces:   []string{"int1", "int2"},
		}},
		{"v2 tunnel zone", "", version.Number{8, 0, 0, ""}, Entry{
			Name:                         "seven",
			Mode:                         ModeTunnel,
			EnablePacketBufferProtection: true,
		}},
		{"v3 empty zone", "", version.Number{10, 0, 0, ""}, Entry{
			Name:                         "one",
			Mode:                         ModeL3,
			EnablePacketBufferProtection: true,
		}},
		{"v3 layer3 zone", "vsys1", version.Number{10, 0, 0, ""}, Entry{
			Name:         "two",
			Mode:         ModeL3,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/1", "ethernet1/2"},
			ZoneProfile:  "profile1",
			LogSetting:   "setting1",
			IncludeAcls:  []string{"10.1.2.0/24"},
		}},
		{"v3 layer2 zone", "vsys2", version.Number{10, 0, 0, ""}, Entry{
			Name:                         "three",
			Mode:                         ModeL2,
			EnablePacketBufferProtection: true,
			Interfaces:                   []string{"ethernet1/3", "ethernet1/4"},
			ExcludeAcls:                  []string{"10.100.1.0/24"},
		}},
		{"v3 vwire zone", "vsys3", version.Number{10, 0, 0, ""}, Entry{
			Name:         "four",
			Mode:         ModeVirtualWire,
			EnableUserId: true,
			Interfaces:   []string{"ethernet1/5", "ethernet1/6"},
			IncludeAcls:  []string{"10.1.3.0/24"},
		}},
		{"v3 external zone", "vsys4", version.Number{10, 0, 0, ""}, Entry{
			Name:                         "five",
			Mode:                         ModeExternal,
			EnablePacketBufferProtection: true,
			Interfaces:                   []string{"ext1", "ext2"},
			ExcludeAcls:                  []string{"10.100.2.0/24"},
		}},
		{"v3 tap zone", "vsys1", version.Number{10, 0, 0, ""}, Entry{
			Name:         "six",
			Mode:         ModeTap,
			EnableUserId: true,
			Interfaces:   []string{"int1", "int2"},
		}},
		{"v3 tunnel zone", "", version.Number{10, 0, 0, ""}, Entry{
			Name:                         "seven",
			Mode:                         ModeTunnel,
			EnablePacketBufferProtection: true,
		}},
		{"v3 device id enabled no acls", "", version.Number{10, 0, 0, ""}, Entry{
			Name:                       "eight",
			Mode:                       ModeL3,
			EnableDeviceIdentification: true,
		}},
		{"v3 device id enabled include acls", "", version.Number{10, 0, 0, ""}, Entry{
			Name:                       "eight",
			Mode:                       ModeL3,
			EnableDeviceIdentification: true,
			DeviceIncludeAcls:          []string{"inc1", "inc2"},
		}},
		{"v3 device id enabled exclude acls", "", version.Number{10, 0, 0, ""}, Entry{
			Name:                       "eight",
			Mode:                       ModeL3,
			EnableDeviceIdentification: true,
			DeviceExcludeAcls:          []string{"exc1", "exc2"},
		}},
		{"v3 device id disabled all acls", "", version.Number{10, 0, 0, ""}, Entry{
			Name:              "eight",
			Mode:              ModeL3,
			DeviceIncludeAcls: []string{"inc1", "inc2"},
			DeviceExcludeAcls: []string{"exc1", "exc2"},
		}},
	}
}
