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
	suffix = []string{"zone", "$name"}
)

type Entry struct {
	Name                       string
	DeviceAcl                  *DeviceAcl
	EnableDeviceIdentification *bool
	EnableUserIdentification   *bool
	Network                    *Network
	UserAcl                    *UserAcl
	Misc                       []generic.Xml
}
type DeviceAcl struct {
	ExcludeList []string
	IncludeList []string
	Misc        []generic.Xml
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
	Misc                         []generic.Xml
}
type NetworkTunnel struct {
	Misc []generic.Xml
}
type UserAcl struct {
	ExcludeList []string
	IncludeList []string
	Misc        []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName                    xml.Name      `xml:"entry"`
	Name                       string        `xml:"name,attr"`
	DeviceAcl                  *deviceAclXml `xml:"device-acl,omitempty"`
	EnableDeviceIdentification *string       `xml:"enable-device-identification,omitempty"`
	EnableUserIdentification   *string       `xml:"enable-user-identification,omitempty"`
	Network                    *networkXml   `xml:"network,omitempty"`
	UserAcl                    *userAclXml   `xml:"user-acl,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type deviceAclXml struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`
	Misc        []generic.Xml    `xml:",any"`
}
type networkXml struct {
	EnablePacketBufferProtection *string           `xml:"enable-packet-buffer-protection,omitempty"`
	LogSetting                   *string           `xml:"log-setting,omitempty"`
	NetInspection                *string           `xml:"net-inspection,omitempty"`
	ZoneProtectionProfile        *string           `xml:"zone-protection-profile,omitempty"`
	External                     *util.MemberType  `xml:"external,omitempty"`
	Layer2                       *util.MemberType  `xml:"layer2,omitempty"`
	Layer3                       *util.MemberType  `xml:"layer3,omitempty"`
	Tap                          *util.MemberType  `xml:"tap,omitempty"`
	Tunnel                       *networkTunnelXml `xml:"tunnel,omitempty"`
	VirtualWire                  *util.MemberType  `xml:"virtual-wire,omitempty"`
	Misc                         []generic.Xml     `xml:",any"`
}
type networkTunnelXml struct {
	Misc []generic.Xml `xml:",any"`
}
type userAclXml struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`
	Misc        []generic.Xml    `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name             `xml:"entry"`
	Name                       string               `xml:"name,attr"`
	DeviceAcl                  *deviceAclXml_11_0_2 `xml:"device-acl,omitempty"`
	EnableDeviceIdentification *string              `xml:"enable-device-identification,omitempty"`
	EnableUserIdentification   *string              `xml:"enable-user-identification,omitempty"`
	Network                    *networkXml_11_0_2   `xml:"network,omitempty"`
	UserAcl                    *userAclXml_11_0_2   `xml:"user-acl,omitempty"`
	Misc                       []generic.Xml        `xml:",any"`
}
type deviceAclXml_11_0_2 struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`
	Misc        []generic.Xml    `xml:",any"`
}
type networkXml_11_0_2 struct {
	EnablePacketBufferProtection *string                  `xml:"enable-packet-buffer-protection,omitempty"`
	LogSetting                   *string                  `xml:"log-setting,omitempty"`
	NetInspection                *string                  `xml:"net-inspection,omitempty"`
	ZoneProtectionProfile        *string                  `xml:"zone-protection-profile,omitempty"`
	External                     *util.MemberType         `xml:"external,omitempty"`
	Layer2                       *util.MemberType         `xml:"layer2,omitempty"`
	Layer3                       *util.MemberType         `xml:"layer3,omitempty"`
	Tap                          *util.MemberType         `xml:"tap,omitempty"`
	Tunnel                       *networkTunnelXml_11_0_2 `xml:"tunnel,omitempty"`
	VirtualWire                  *util.MemberType         `xml:"virtual-wire,omitempty"`
	Misc                         []generic.Xml            `xml:",any"`
}
type networkTunnelXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type userAclXml_11_0_2 struct {
	ExcludeList *util.MemberType `xml:"exclude-list,omitempty"`
	IncludeList *util.MemberType `xml:"include-list,omitempty"`
	Misc        []generic.Xml    `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.DeviceAcl != nil {
		var obj deviceAclXml
		obj.MarshalFromObject(*s.DeviceAcl)
		o.DeviceAcl = &obj
	}
	o.EnableDeviceIdentification = util.YesNo(s.EnableDeviceIdentification, nil)
	o.EnableUserIdentification = util.YesNo(s.EnableUserIdentification, nil)
	if s.Network != nil {
		var obj networkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.UserAcl != nil {
		var obj userAclXml
		obj.MarshalFromObject(*s.UserAcl)
		o.UserAcl = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var deviceAclVal *DeviceAcl
	if o.DeviceAcl != nil {
		obj, err := o.DeviceAcl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceAclVal = obj
	}
	var networkVal *Network
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var userAclVal *UserAcl
	if o.UserAcl != nil {
		obj, err := o.UserAcl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		userAclVal = obj
	}

	result := &Entry{
		Name:                       o.Name,
		DeviceAcl:                  deviceAclVal,
		EnableDeviceIdentification: util.AsBool(o.EnableDeviceIdentification, nil),
		EnableUserIdentification:   util.AsBool(o.EnableUserIdentification, nil),
		Network:                    networkVal,
		UserAcl:                    userAclVal,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *deviceAclXml) MarshalFromObject(s DeviceAcl) {
	if s.ExcludeList != nil {
		o.ExcludeList = util.StrToMem(s.ExcludeList)
	}
	if s.IncludeList != nil {
		o.IncludeList = util.StrToMem(s.IncludeList)
	}
	o.Misc = s.Misc
}

func (o deviceAclXml) UnmarshalToObject() (*DeviceAcl, error) {
	var excludeListVal []string
	if o.ExcludeList != nil {
		excludeListVal = util.MemToStr(o.ExcludeList)
	}
	var includeListVal []string
	if o.IncludeList != nil {
		includeListVal = util.MemToStr(o.IncludeList)
	}

	result := &DeviceAcl{
		ExcludeList: excludeListVal,
		IncludeList: includeListVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *networkXml) MarshalFromObject(s Network) {
	o.EnablePacketBufferProtection = util.YesNo(s.EnablePacketBufferProtection, nil)
	o.LogSetting = s.LogSetting
	o.NetInspection = util.YesNo(s.NetInspection, nil)
	o.ZoneProtectionProfile = s.ZoneProtectionProfile
	if s.External != nil {
		o.External = util.StrToMem(s.External)
	}
	if s.Layer2 != nil {
		o.Layer2 = util.StrToMem(s.Layer2)
	}
	if s.Layer3 != nil {
		o.Layer3 = util.StrToMem(s.Layer3)
	}
	if s.Tap != nil {
		o.Tap = util.StrToMem(s.Tap)
	}
	if s.Tunnel != nil {
		var obj networkTunnelXml
		obj.MarshalFromObject(*s.Tunnel)
		o.Tunnel = &obj
	}
	if s.VirtualWire != nil {
		o.VirtualWire = util.StrToMem(s.VirtualWire)
	}
	o.Misc = s.Misc
}

func (o networkXml) UnmarshalToObject() (*Network, error) {
	var externalVal []string
	if o.External != nil {
		externalVal = util.MemToStr(o.External)
	}
	var layer2Val []string
	if o.Layer2 != nil {
		layer2Val = util.MemToStr(o.Layer2)
	}
	var layer3Val []string
	if o.Layer3 != nil {
		layer3Val = util.MemToStr(o.Layer3)
	}
	var tapVal []string
	if o.Tap != nil {
		tapVal = util.MemToStr(o.Tap)
	}
	var tunnelVal *NetworkTunnel
	if o.Tunnel != nil {
		obj, err := o.Tunnel.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tunnelVal = obj
	}
	var virtualWireVal []string
	if o.VirtualWire != nil {
		virtualWireVal = util.MemToStr(o.VirtualWire)
	}

	result := &Network{
		EnablePacketBufferProtection: util.AsBool(o.EnablePacketBufferProtection, nil),
		LogSetting:                   o.LogSetting,
		NetInspection:                util.AsBool(o.NetInspection, nil),
		ZoneProtectionProfile:        o.ZoneProtectionProfile,
		External:                     externalVal,
		Layer2:                       layer2Val,
		Layer3:                       layer3Val,
		Tap:                          tapVal,
		Tunnel:                       tunnelVal,
		VirtualWire:                  virtualWireVal,
		Misc:                         o.Misc,
	}
	return result, nil
}
func (o *networkTunnelXml) MarshalFromObject(s NetworkTunnel) {
	o.Misc = s.Misc
}

func (o networkTunnelXml) UnmarshalToObject() (*NetworkTunnel, error) {

	result := &NetworkTunnel{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *userAclXml) MarshalFromObject(s UserAcl) {
	if s.ExcludeList != nil {
		o.ExcludeList = util.StrToMem(s.ExcludeList)
	}
	if s.IncludeList != nil {
		o.IncludeList = util.StrToMem(s.IncludeList)
	}
	o.Misc = s.Misc
}

func (o userAclXml) UnmarshalToObject() (*UserAcl, error) {
	var excludeListVal []string
	if o.ExcludeList != nil {
		excludeListVal = util.MemToStr(o.ExcludeList)
	}
	var includeListVal []string
	if o.IncludeList != nil {
		includeListVal = util.MemToStr(o.IncludeList)
	}

	result := &UserAcl{
		ExcludeList: excludeListVal,
		IncludeList: includeListVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.DeviceAcl != nil {
		var obj deviceAclXml_11_0_2
		obj.MarshalFromObject(*s.DeviceAcl)
		o.DeviceAcl = &obj
	}
	o.EnableDeviceIdentification = util.YesNo(s.EnableDeviceIdentification, nil)
	o.EnableUserIdentification = util.YesNo(s.EnableUserIdentification, nil)
	if s.Network != nil {
		var obj networkXml_11_0_2
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.UserAcl != nil {
		var obj userAclXml_11_0_2
		obj.MarshalFromObject(*s.UserAcl)
		o.UserAcl = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var deviceAclVal *DeviceAcl
	if o.DeviceAcl != nil {
		obj, err := o.DeviceAcl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceAclVal = obj
	}
	var networkVal *Network
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var userAclVal *UserAcl
	if o.UserAcl != nil {
		obj, err := o.UserAcl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		userAclVal = obj
	}

	result := &Entry{
		Name:                       o.Name,
		DeviceAcl:                  deviceAclVal,
		EnableDeviceIdentification: util.AsBool(o.EnableDeviceIdentification, nil),
		EnableUserIdentification:   util.AsBool(o.EnableUserIdentification, nil),
		Network:                    networkVal,
		UserAcl:                    userAclVal,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *deviceAclXml_11_0_2) MarshalFromObject(s DeviceAcl) {
	if s.ExcludeList != nil {
		o.ExcludeList = util.StrToMem(s.ExcludeList)
	}
	if s.IncludeList != nil {
		o.IncludeList = util.StrToMem(s.IncludeList)
	}
	o.Misc = s.Misc
}

func (o deviceAclXml_11_0_2) UnmarshalToObject() (*DeviceAcl, error) {
	var excludeListVal []string
	if o.ExcludeList != nil {
		excludeListVal = util.MemToStr(o.ExcludeList)
	}
	var includeListVal []string
	if o.IncludeList != nil {
		includeListVal = util.MemToStr(o.IncludeList)
	}

	result := &DeviceAcl{
		ExcludeList: excludeListVal,
		IncludeList: includeListVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *networkXml_11_0_2) MarshalFromObject(s Network) {
	o.EnablePacketBufferProtection = util.YesNo(s.EnablePacketBufferProtection, nil)
	o.LogSetting = s.LogSetting
	o.NetInspection = util.YesNo(s.NetInspection, nil)
	o.ZoneProtectionProfile = s.ZoneProtectionProfile
	if s.External != nil {
		o.External = util.StrToMem(s.External)
	}
	if s.Layer2 != nil {
		o.Layer2 = util.StrToMem(s.Layer2)
	}
	if s.Layer3 != nil {
		o.Layer3 = util.StrToMem(s.Layer3)
	}
	if s.Tap != nil {
		o.Tap = util.StrToMem(s.Tap)
	}
	if s.Tunnel != nil {
		var obj networkTunnelXml_11_0_2
		obj.MarshalFromObject(*s.Tunnel)
		o.Tunnel = &obj
	}
	if s.VirtualWire != nil {
		o.VirtualWire = util.StrToMem(s.VirtualWire)
	}
	o.Misc = s.Misc
}

func (o networkXml_11_0_2) UnmarshalToObject() (*Network, error) {
	var externalVal []string
	if o.External != nil {
		externalVal = util.MemToStr(o.External)
	}
	var layer2Val []string
	if o.Layer2 != nil {
		layer2Val = util.MemToStr(o.Layer2)
	}
	var layer3Val []string
	if o.Layer3 != nil {
		layer3Val = util.MemToStr(o.Layer3)
	}
	var tapVal []string
	if o.Tap != nil {
		tapVal = util.MemToStr(o.Tap)
	}
	var tunnelVal *NetworkTunnel
	if o.Tunnel != nil {
		obj, err := o.Tunnel.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tunnelVal = obj
	}
	var virtualWireVal []string
	if o.VirtualWire != nil {
		virtualWireVal = util.MemToStr(o.VirtualWire)
	}

	result := &Network{
		EnablePacketBufferProtection: util.AsBool(o.EnablePacketBufferProtection, nil),
		LogSetting:                   o.LogSetting,
		NetInspection:                util.AsBool(o.NetInspection, nil),
		ZoneProtectionProfile:        o.ZoneProtectionProfile,
		External:                     externalVal,
		Layer2:                       layer2Val,
		Layer3:                       layer3Val,
		Tap:                          tapVal,
		Tunnel:                       tunnelVal,
		VirtualWire:                  virtualWireVal,
		Misc:                         o.Misc,
	}
	return result, nil
}
func (o *networkTunnelXml_11_0_2) MarshalFromObject(s NetworkTunnel) {
	o.Misc = s.Misc
}

func (o networkTunnelXml_11_0_2) UnmarshalToObject() (*NetworkTunnel, error) {

	result := &NetworkTunnel{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *userAclXml_11_0_2) MarshalFromObject(s UserAcl) {
	if s.ExcludeList != nil {
		o.ExcludeList = util.StrToMem(s.ExcludeList)
	}
	if s.IncludeList != nil {
		o.IncludeList = util.StrToMem(s.IncludeList)
	}
	o.Misc = s.Misc
}

func (o userAclXml_11_0_2) UnmarshalToObject() (*UserAcl, error) {
	var excludeListVal []string
	if o.ExcludeList != nil {
		excludeListVal = util.MemToStr(o.ExcludeList)
	}
	var includeListVal []string
	if o.IncludeList != nil {
		includeListVal = util.MemToStr(o.IncludeList)
	}

	result := &UserAcl{
		ExcludeList: excludeListVal,
		IncludeList: includeListVal,
		Misc:        o.Misc,
	}
	return result, nil
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
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.DeviceAcl.matches(other.DeviceAcl) {
		return false
	}
	if !util.BoolsMatch(o.EnableDeviceIdentification, other.EnableDeviceIdentification) {
		return false
	}
	if !util.BoolsMatch(o.EnableUserIdentification, other.EnableUserIdentification) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}
	if !o.UserAcl.matches(other.UserAcl) {
		return false
	}

	return true
}

func (o *DeviceAcl) matches(other *DeviceAcl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExcludeList, other.ExcludeList) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IncludeList, other.IncludeList) {
		return false
	}

	return true
}

func (o *Network) matches(other *Network) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.EnablePacketBufferProtection, other.EnablePacketBufferProtection) {
		return false
	}
	if !util.StringsMatch(o.LogSetting, other.LogSetting) {
		return false
	}
	if !util.BoolsMatch(o.NetInspection, other.NetInspection) {
		return false
	}
	if !util.StringsMatch(o.ZoneProtectionProfile, other.ZoneProtectionProfile) {
		return false
	}
	if !util.OrderedListsMatch[string](o.External, other.External) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Layer2, other.Layer2) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Layer3, other.Layer3) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tap, other.Tap) {
		return false
	}
	if !o.Tunnel.matches(other.Tunnel) {
		return false
	}
	if !util.OrderedListsMatch[string](o.VirtualWire, other.VirtualWire) {
		return false
	}

	return true
}

func (o *NetworkTunnel) matches(other *NetworkTunnel) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UserAcl) matches(other *UserAcl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExcludeList, other.ExcludeList) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IncludeList, other.IncludeList) {
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
