package zone

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	Suffix = []string{"zone"}
)

type Entry struct {
	Name                       string
	DeviceAcl                  *DeviceAcl
	EnableDeviceIdentification *bool
	EnableUserIdentification   *bool
	Network                    *Network
	UserAcl                    *UserAcl

	Misc map[string][]generic.Xml
}

type DeviceAcl struct {
	ExcludeList []string
	IncludeList []string
}
type Network struct {
	EnablePacketBufferProtection *bool
	LogSetting                   *string
	NetInspection                *bool
	ZoneProtectionProfile        *string
	External                     []string
	Layer2                       []string
	Layer3                       []string
	Tap                          []string
	Tunnel                       *NetworkTunnel
	VirtualWire                  []string
}
type NetworkTunnel struct {
}
type UserAcl struct {
	ExcludeList []string
	IncludeList []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}
type entryXml struct {
	XMLName                    xml.Name      `xml:"entry"`
	Name                       string        `xml:"name,attr"`
	DeviceAcl                  *DeviceAclXml `xml:"device-acl,omitempty"`
	EnableDeviceIdentification *string       `xml:"enable-device-identification,omitempty"`
	EnableUserIdentification   *string       `xml:"enable-user-identification,omitempty"`
	Network                    *NetworkXml   `xml:"network,omitempty"`
	UserAcl                    *UserAclXml   `xml:"user-acl,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name             `xml:"entry"`
	Name                       string               `xml:"name,attr"`
	DeviceAcl                  *DeviceAclXml_11_0_2 `xml:"device-acl,omitempty"`
	EnableDeviceIdentification *string              `xml:"enable-device-identification,omitempty"`
	EnableUserIdentification   *string              `xml:"enable-user-identification,omitempty"`
	Network                    *NetworkXml_11_0_2   `xml:"network,omitempty"`
	UserAcl                    *UserAclXml_11_0_2   `xml:"user-acl,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DeviceAclXml struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NetworkXml struct {
	EnablePacketBufferProtection *string           `xml:"enable-packet-buffer-protection,omitempty"`
	LogSetting                   *string           `xml:"log-setting,omitempty"`
	NetInspection                *string           `xml:"net-inspection,omitempty"`
	ZoneProtectionProfile        *string           `xml:"zone-protection-profile,omitempty"`
	External                     *util.MemberType  `xml:"external,omitempty"`
	Layer2                       *util.MemberType  `xml:"layer2,omitempty"`
	Layer3                       *util.MemberType  `xml:"layer3,omitempty"`
	Tap                          *util.MemberType  `xml:"tap,omitempty"`
	Tunnel                       *NetworkTunnelXml `xml:"tunnel,omitempty"`
	VirtualWire                  *util.MemberType  `xml:"virtual-wire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NetworkTunnelXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UserAclXml struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DeviceAclXml_11_0_2 struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NetworkXml_11_0_2 struct {
	EnablePacketBufferProtection *string                  `xml:"enable-packet-buffer-protection,omitempty"`
	LogSetting                   *string                  `xml:"log-setting,omitempty"`
	NetInspection                *string                  `xml:"net-inspection,omitempty"`
	ZoneProtectionProfile        *string                  `xml:"zone-protection-profile,omitempty"`
	External                     *util.MemberType         `xml:"external,omitempty"`
	Layer2                       *util.MemberType         `xml:"layer2,omitempty"`
	Layer3                       *util.MemberType         `xml:"layer3,omitempty"`
	Tap                          *util.MemberType         `xml:"tap,omitempty"`
	Tunnel                       *NetworkTunnelXml_11_0_2 `xml:"tunnel,omitempty"`
	VirtualWire                  *util.MemberType         `xml:"virtual-wire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NetworkTunnelXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type UserAclXml_11_0_2 struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "device_acl" || v == "DeviceAcl" {
		return e.DeviceAcl, nil
	}
	if v == "enable_device_identification" || v == "EnableDeviceIdentification" {
		return e.EnableDeviceIdentification, nil
	}
	if v == "enable_user_identification" || v == "EnableUserIdentification" {
		return e.EnableUserIdentification, nil
	}
	if v == "network" || v == "Network" {
		return e.Network, nil
	}
	if v == "user_acl" || v == "UserAcl" {
		return e.UserAcl, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedDeviceAcl *DeviceAclXml
	if o.DeviceAcl != nil {
		nestedDeviceAcl = &DeviceAclXml{}
		if _, ok := o.Misc["DeviceAcl"]; ok {
			nestedDeviceAcl.Misc = o.Misc["DeviceAcl"]
		}
		if o.DeviceAcl.ExcludeList != nil {
			nestedDeviceAcl.ExcludeList = util.StrToMem(o.DeviceAcl.ExcludeList)
		}
		if o.DeviceAcl.IncludeList != nil {
			nestedDeviceAcl.IncludeList = util.StrToMem(o.DeviceAcl.IncludeList)
		}
	}
	entry.DeviceAcl = nestedDeviceAcl

	entry.EnableDeviceIdentification = util.YesNo(o.EnableDeviceIdentification, nil)
	entry.EnableUserIdentification = util.YesNo(o.EnableUserIdentification, nil)
	var nestedNetwork *NetworkXml
	if o.Network != nil {
		nestedNetwork = &NetworkXml{}
		if _, ok := o.Misc["Network"]; ok {
			nestedNetwork.Misc = o.Misc["Network"]
		}
		if o.Network.EnablePacketBufferProtection != nil {
			nestedNetwork.EnablePacketBufferProtection = util.YesNo(o.Network.EnablePacketBufferProtection, nil)
		}
		if o.Network.LogSetting != nil {
			nestedNetwork.LogSetting = o.Network.LogSetting
		}
		if o.Network.ZoneProtectionProfile != nil {
			nestedNetwork.ZoneProtectionProfile = o.Network.ZoneProtectionProfile
		}
		if o.Network.NetInspection != nil {
			nestedNetwork.NetInspection = util.YesNo(o.Network.NetInspection, nil)
		}
		if o.Network.Tap != nil {
			nestedNetwork.Tap = util.StrToMem(o.Network.Tap)
		}
		if o.Network.Tunnel != nil {
			nestedNetwork.Tunnel = &NetworkTunnelXml{}
			if _, ok := o.Misc["NetworkTunnel"]; ok {
				nestedNetwork.Tunnel.Misc = o.Misc["NetworkTunnel"]
			}
		}
		if o.Network.VirtualWire != nil {
			nestedNetwork.VirtualWire = util.StrToMem(o.Network.VirtualWire)
		}
		if o.Network.External != nil {
			nestedNetwork.External = util.StrToMem(o.Network.External)
		}
		if o.Network.Layer2 != nil {
			nestedNetwork.Layer2 = util.StrToMem(o.Network.Layer2)
		}
		if o.Network.Layer3 != nil {
			nestedNetwork.Layer3 = util.StrToMem(o.Network.Layer3)
		}
	}
	entry.Network = nestedNetwork

	var nestedUserAcl *UserAclXml
	if o.UserAcl != nil {
		nestedUserAcl = &UserAclXml{}
		if _, ok := o.Misc["UserAcl"]; ok {
			nestedUserAcl.Misc = o.Misc["UserAcl"]
		}
		if o.UserAcl.ExcludeList != nil {
			nestedUserAcl.ExcludeList = util.StrToMem(o.UserAcl.ExcludeList)
		}
		if o.UserAcl.IncludeList != nil {
			nestedUserAcl.IncludeList = util.StrToMem(o.UserAcl.IncludeList)
		}
	}
	entry.UserAcl = nestedUserAcl

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	var nestedDeviceAcl *DeviceAclXml_11_0_2
	if o.DeviceAcl != nil {
		nestedDeviceAcl = &DeviceAclXml_11_0_2{}
		if _, ok := o.Misc["DeviceAcl"]; ok {
			nestedDeviceAcl.Misc = o.Misc["DeviceAcl"]
		}
		if o.DeviceAcl.IncludeList != nil {
			nestedDeviceAcl.IncludeList = util.StrToMem(o.DeviceAcl.IncludeList)
		}
		if o.DeviceAcl.ExcludeList != nil {
			nestedDeviceAcl.ExcludeList = util.StrToMem(o.DeviceAcl.ExcludeList)
		}
	}
	entry.DeviceAcl = nestedDeviceAcl

	entry.EnableDeviceIdentification = util.YesNo(o.EnableDeviceIdentification, nil)
	entry.EnableUserIdentification = util.YesNo(o.EnableUserIdentification, nil)
	var nestedNetwork *NetworkXml_11_0_2
	if o.Network != nil {
		nestedNetwork = &NetworkXml_11_0_2{}
		if _, ok := o.Misc["Network"]; ok {
			nestedNetwork.Misc = o.Misc["Network"]
		}
		if o.Network.EnablePacketBufferProtection != nil {
			nestedNetwork.EnablePacketBufferProtection = util.YesNo(o.Network.EnablePacketBufferProtection, nil)
		}
		if o.Network.LogSetting != nil {
			nestedNetwork.LogSetting = o.Network.LogSetting
		}
		if o.Network.ZoneProtectionProfile != nil {
			nestedNetwork.ZoneProtectionProfile = o.Network.ZoneProtectionProfile
		}
		if o.Network.NetInspection != nil {
			nestedNetwork.NetInspection = util.YesNo(o.Network.NetInspection, nil)
		}
		if o.Network.VirtualWire != nil {
			nestedNetwork.VirtualWire = util.StrToMem(o.Network.VirtualWire)
		}
		if o.Network.External != nil {
			nestedNetwork.External = util.StrToMem(o.Network.External)
		}
		if o.Network.Layer2 != nil {
			nestedNetwork.Layer2 = util.StrToMem(o.Network.Layer2)
		}
		if o.Network.Layer3 != nil {
			nestedNetwork.Layer3 = util.StrToMem(o.Network.Layer3)
		}
		if o.Network.Tap != nil {
			nestedNetwork.Tap = util.StrToMem(o.Network.Tap)
		}
		if o.Network.Tunnel != nil {
			nestedNetwork.Tunnel = &NetworkTunnelXml_11_0_2{}
			if _, ok := o.Misc["NetworkTunnel"]; ok {
				nestedNetwork.Tunnel.Misc = o.Misc["NetworkTunnel"]
			}
		}
	}
	entry.Network = nestedNetwork

	var nestedUserAcl *UserAclXml_11_0_2
	if o.UserAcl != nil {
		nestedUserAcl = &UserAclXml_11_0_2{}
		if _, ok := o.Misc["UserAcl"]; ok {
			nestedUserAcl.Misc = o.Misc["UserAcl"]
		}
		if o.UserAcl.ExcludeList != nil {
			nestedUserAcl.ExcludeList = util.StrToMem(o.UserAcl.ExcludeList)
		}
		if o.UserAcl.IncludeList != nil {
			nestedUserAcl.IncludeList = util.StrToMem(o.UserAcl.IncludeList)
		}
	}
	entry.UserAcl = nestedUserAcl

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}
func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedDeviceAcl *DeviceAcl
		if o.DeviceAcl != nil {
			nestedDeviceAcl = &DeviceAcl{}
			if o.DeviceAcl.Misc != nil {
				entry.Misc["DeviceAcl"] = o.DeviceAcl.Misc
			}
			if o.DeviceAcl.ExcludeList != nil {
				nestedDeviceAcl.ExcludeList = util.MemToStr(o.DeviceAcl.ExcludeList)
			}
			if o.DeviceAcl.IncludeList != nil {
				nestedDeviceAcl.IncludeList = util.MemToStr(o.DeviceAcl.IncludeList)
			}
		}
		entry.DeviceAcl = nestedDeviceAcl

		entry.EnableDeviceIdentification = util.AsBool(o.EnableDeviceIdentification, nil)
		entry.EnableUserIdentification = util.AsBool(o.EnableUserIdentification, nil)
		var nestedNetwork *Network
		if o.Network != nil {
			nestedNetwork = &Network{}
			if o.Network.Misc != nil {
				entry.Misc["Network"] = o.Network.Misc
			}
			if o.Network.ZoneProtectionProfile != nil {
				nestedNetwork.ZoneProtectionProfile = o.Network.ZoneProtectionProfile
			}
			if o.Network.NetInspection != nil {
				nestedNetwork.NetInspection = util.AsBool(o.Network.NetInspection, nil)
			}
			if o.Network.EnablePacketBufferProtection != nil {
				nestedNetwork.EnablePacketBufferProtection = util.AsBool(o.Network.EnablePacketBufferProtection, nil)
			}
			if o.Network.LogSetting != nil {
				nestedNetwork.LogSetting = o.Network.LogSetting
			}
			if o.Network.Tunnel != nil {
				nestedNetwork.Tunnel = &NetworkTunnel{}
				if o.Network.Tunnel.Misc != nil {
					entry.Misc["NetworkTunnel"] = o.Network.Tunnel.Misc
				}
			}
			if o.Network.VirtualWire != nil {
				nestedNetwork.VirtualWire = util.MemToStr(o.Network.VirtualWire)
			}
			if o.Network.External != nil {
				nestedNetwork.External = util.MemToStr(o.Network.External)
			}
			if o.Network.Layer2 != nil {
				nestedNetwork.Layer2 = util.MemToStr(o.Network.Layer2)
			}
			if o.Network.Layer3 != nil {
				nestedNetwork.Layer3 = util.MemToStr(o.Network.Layer3)
			}
			if o.Network.Tap != nil {
				nestedNetwork.Tap = util.MemToStr(o.Network.Tap)
			}
		}
		entry.Network = nestedNetwork

		var nestedUserAcl *UserAcl
		if o.UserAcl != nil {
			nestedUserAcl = &UserAcl{}
			if o.UserAcl.Misc != nil {
				entry.Misc["UserAcl"] = o.UserAcl.Misc
			}
			if o.UserAcl.ExcludeList != nil {
				nestedUserAcl.ExcludeList = util.MemToStr(o.UserAcl.ExcludeList)
			}
			if o.UserAcl.IncludeList != nil {
				nestedUserAcl.IncludeList = util.MemToStr(o.UserAcl.IncludeList)
			}
		}
		entry.UserAcl = nestedUserAcl

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}
func (c *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedDeviceAcl *DeviceAcl
		if o.DeviceAcl != nil {
			nestedDeviceAcl = &DeviceAcl{}
			if o.DeviceAcl.Misc != nil {
				entry.Misc["DeviceAcl"] = o.DeviceAcl.Misc
			}
			if o.DeviceAcl.ExcludeList != nil {
				nestedDeviceAcl.ExcludeList = util.MemToStr(o.DeviceAcl.ExcludeList)
			}
			if o.DeviceAcl.IncludeList != nil {
				nestedDeviceAcl.IncludeList = util.MemToStr(o.DeviceAcl.IncludeList)
			}
		}
		entry.DeviceAcl = nestedDeviceAcl

		entry.EnableDeviceIdentification = util.AsBool(o.EnableDeviceIdentification, nil)
		entry.EnableUserIdentification = util.AsBool(o.EnableUserIdentification, nil)
		var nestedNetwork *Network
		if o.Network != nil {
			nestedNetwork = &Network{}
			if o.Network.Misc != nil {
				entry.Misc["Network"] = o.Network.Misc
			}
			if o.Network.EnablePacketBufferProtection != nil {
				nestedNetwork.EnablePacketBufferProtection = util.AsBool(o.Network.EnablePacketBufferProtection, nil)
			}
			if o.Network.LogSetting != nil {
				nestedNetwork.LogSetting = o.Network.LogSetting
			}
			if o.Network.ZoneProtectionProfile != nil {
				nestedNetwork.ZoneProtectionProfile = o.Network.ZoneProtectionProfile
			}
			if o.Network.NetInspection != nil {
				nestedNetwork.NetInspection = util.AsBool(o.Network.NetInspection, nil)
			}
			if o.Network.Tap != nil {
				nestedNetwork.Tap = util.MemToStr(o.Network.Tap)
			}
			if o.Network.Tunnel != nil {
				nestedNetwork.Tunnel = &NetworkTunnel{}
				if o.Network.Tunnel.Misc != nil {
					entry.Misc["NetworkTunnel"] = o.Network.Tunnel.Misc
				}
			}
			if o.Network.VirtualWire != nil {
				nestedNetwork.VirtualWire = util.MemToStr(o.Network.VirtualWire)
			}
			if o.Network.External != nil {
				nestedNetwork.External = util.MemToStr(o.Network.External)
			}
			if o.Network.Layer2 != nil {
				nestedNetwork.Layer2 = util.MemToStr(o.Network.Layer2)
			}
			if o.Network.Layer3 != nil {
				nestedNetwork.Layer3 = util.MemToStr(o.Network.Layer3)
			}
		}
		entry.Network = nestedNetwork

		var nestedUserAcl *UserAcl
		if o.UserAcl != nil {
			nestedUserAcl = &UserAcl{}
			if o.UserAcl.Misc != nil {
				entry.Misc["UserAcl"] = o.UserAcl.Misc
			}
			if o.UserAcl.ExcludeList != nil {
				nestedUserAcl.ExcludeList = util.MemToStr(o.UserAcl.ExcludeList)
			}
			if o.UserAcl.IncludeList != nil {
				nestedUserAcl.IncludeList = util.MemToStr(o.UserAcl.IncludeList)
			}
		}
		entry.UserAcl = nestedUserAcl

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !matchDeviceAcl(a.DeviceAcl, b.DeviceAcl) {
		return false
	}
	if !util.BoolsMatch(a.EnableDeviceIdentification, b.EnableDeviceIdentification) {
		return false
	}
	if !util.BoolsMatch(a.EnableUserIdentification, b.EnableUserIdentification) {
		return false
	}
	if !matchNetwork(a.Network, b.Network) {
		return false
	}
	if !matchUserAcl(a.UserAcl, b.UserAcl) {
		return false
	}

	return true
}

func matchNetworkTunnel(a *NetworkTunnel, b *NetworkTunnel) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchNetwork(a *Network, b *Network) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ZoneProtectionProfile, b.ZoneProtectionProfile) {
		return false
	}
	if !util.BoolsMatch(a.NetInspection, b.NetInspection) {
		return false
	}
	if !util.BoolsMatch(a.EnablePacketBufferProtection, b.EnablePacketBufferProtection) {
		return false
	}
	if !util.StringsMatch(a.LogSetting, b.LogSetting) {
		return false
	}
	if !util.OrderedListsMatch(a.Tap, b.Tap) {
		return false
	}
	if !matchNetworkTunnel(a.Tunnel, b.Tunnel) {
		return false
	}
	if !util.OrderedListsMatch(a.VirtualWire, b.VirtualWire) {
		return false
	}
	if !util.OrderedListsMatch(a.External, b.External) {
		return false
	}
	if !util.OrderedListsMatch(a.Layer2, b.Layer2) {
		return false
	}
	if !util.OrderedListsMatch(a.Layer3, b.Layer3) {
		return false
	}
	return true
}
func matchUserAcl(a *UserAcl, b *UserAcl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.ExcludeList, b.ExcludeList) {
		return false
	}
	if !util.OrderedListsMatch(a.IncludeList, b.IncludeList) {
		return false
	}
	return true
}
func matchDeviceAcl(a *DeviceAcl, b *DeviceAcl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.ExcludeList, b.ExcludeList) {
		return false
	}
	if !util.OrderedListsMatch(a.IncludeList, b.IncludeList) {
		return false
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
