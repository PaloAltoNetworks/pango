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
	LogSetting                   []string
	ZoneProtectionProfile        []string
	Layer2                       []string
	Layer3                       []string
	Tap                          []string
	VirtualWire                  []string
}
type UserAcl struct {
	ExcludeList []string
	IncludeList []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
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

type DeviceAclXml struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NetworkXml struct {
	EnablePacketBufferProtection *string          `xml:"enable-packet-buffer-protection,omitempty"`
	LogSetting                   *util.MemberType `xml:"log-setting,omitempty"`
	ZoneProtectionProfile        *util.MemberType `xml:"zone-protection-profile,omitempty"`
	Layer2                       *util.MemberType `xml:"layer2,omitempty"`
	Layer3                       *util.MemberType `xml:"layer3,omitempty"`
	Tap                          *util.MemberType `xml:"tap,omitempty"`
	VirtualWire                  *util.MemberType `xml:"virtual-wire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UserAclXml struct {
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
	var nestedNetwork *NetworkXml
	if o.Network != nil {
		nestedNetwork = &NetworkXml{}
		if _, ok := o.Misc["Network"]; ok {
			nestedNetwork.Misc = o.Misc["Network"]
		}
		if o.Network.EnablePacketBufferProtection != nil {
			nestedNetwork.EnablePacketBufferProtection = util.YesNo(o.Network.EnablePacketBufferProtection, nil)
		}
		if o.Network.ZoneProtectionProfile != nil {
			nestedNetwork.ZoneProtectionProfile = util.StrToMem(o.Network.ZoneProtectionProfile)
		}
		if o.Network.LogSetting != nil {
			nestedNetwork.LogSetting = util.StrToMem(o.Network.LogSetting)
		}
		if o.Network.Layer3 != nil {
			nestedNetwork.Layer3 = util.StrToMem(o.Network.Layer3)
		}
		if o.Network.Layer2 != nil {
			nestedNetwork.Layer2 = util.StrToMem(o.Network.Layer2)
		}
		if o.Network.VirtualWire != nil {
			nestedNetwork.VirtualWire = util.StrToMem(o.Network.VirtualWire)
		}
		if o.Network.Tap != nil {
			nestedNetwork.Tap = util.StrToMem(o.Network.Tap)
		}
	}
	entry.Network = nestedNetwork

	var nestedUserAcl *UserAclXml
	if o.UserAcl != nil {
		nestedUserAcl = &UserAclXml{}
		if _, ok := o.Misc["UserAcl"]; ok {
			nestedUserAcl.Misc = o.Misc["UserAcl"]
		}
		if o.UserAcl.IncludeList != nil {
			nestedUserAcl.IncludeList = util.StrToMem(o.UserAcl.IncludeList)
		}
		if o.UserAcl.ExcludeList != nil {
			nestedUserAcl.ExcludeList = util.StrToMem(o.UserAcl.ExcludeList)
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
			if o.DeviceAcl.IncludeList != nil {
				nestedDeviceAcl.IncludeList = util.MemToStr(o.DeviceAcl.IncludeList)
			}
			if o.DeviceAcl.ExcludeList != nil {
				nestedDeviceAcl.ExcludeList = util.MemToStr(o.DeviceAcl.ExcludeList)
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
			if o.Network.ZoneProtectionProfile != nil {
				nestedNetwork.ZoneProtectionProfile = util.MemToStr(o.Network.ZoneProtectionProfile)
			}
			if o.Network.LogSetting != nil {
				nestedNetwork.LogSetting = util.MemToStr(o.Network.LogSetting)
			}
			if o.Network.Layer3 != nil {
				nestedNetwork.Layer3 = util.MemToStr(o.Network.Layer3)
			}
			if o.Network.Layer2 != nil {
				nestedNetwork.Layer2 = util.MemToStr(o.Network.Layer2)
			}
			if o.Network.VirtualWire != nil {
				nestedNetwork.VirtualWire = util.MemToStr(o.Network.VirtualWire)
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
			if o.UserAcl.IncludeList != nil {
				nestedUserAcl.IncludeList = util.MemToStr(o.UserAcl.IncludeList)
			}
			if o.UserAcl.ExcludeList != nil {
				nestedUserAcl.ExcludeList = util.MemToStr(o.UserAcl.ExcludeList)
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

func matchUserAcl(a *UserAcl, b *UserAcl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.IncludeList, b.IncludeList) {
		return false
	}
	if !util.OrderedListsMatch(a.ExcludeList, b.ExcludeList) {
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
	if !util.OrderedListsMatch(a.IncludeList, b.IncludeList) {
		return false
	}
	if !util.OrderedListsMatch(a.ExcludeList, b.ExcludeList) {
		return false
	}
	return true
}
func matchNetwork(a *Network, b *Network) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.EnablePacketBufferProtection, b.EnablePacketBufferProtection) {
		return false
	}
	if !util.OrderedListsMatch(a.ZoneProtectionProfile, b.ZoneProtectionProfile) {
		return false
	}
	if !util.OrderedListsMatch(a.LogSetting, b.LogSetting) {
		return false
	}
	if !util.OrderedListsMatch(a.Tap, b.Tap) {
		return false
	}
	if !util.OrderedListsMatch(a.Layer3, b.Layer3) {
		return false
	}
	if !util.OrderedListsMatch(a.Layer2, b.Layer2) {
		return false
	}
	if !util.OrderedListsMatch(a.VirtualWire, b.VirtualWire) {
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
