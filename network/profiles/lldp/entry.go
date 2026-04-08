package lldp

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
	suffix = []string{"network", "profiles", "lldp-profile", "$name"}
)

type Entry struct {
	Name                   string
	Mode                   *string
	OptionTlvs             *OptionTlvs
	SnmpSyslogNotification *bool
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
}
type OptionTlvs struct {
	ManagementAddress  *OptionTlvsManagementAddress
	PortDescription    *bool
	SystemCapabilities *bool
	SystemDescription  *bool
	SystemName         *bool
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type OptionTlvsManagementAddress struct {
	Enabled        *bool
	Iplist         []OptionTlvsManagementAddressIplist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type OptionTlvsManagementAddressIplist struct {
	Name           string
	Interface      *string
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
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

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName                xml.Name       `xml:"entry"`
	Name                   string         `xml:"name,attr"`
	Mode                   *string        `xml:"mode,omitempty"`
	OptionTlvs             *optionTlvsXml `xml:"option-tlvs,omitempty"`
	SnmpSyslogNotification *string        `xml:"snmp-syslog-notification,omitempty"`
	Misc                   []generic.Xml  `xml:",any"`
	MiscAttributes         []xml.Attr     `xml:",any,attr"`
}
type optionTlvsXml struct {
	ManagementAddress  *optionTlvsManagementAddressXml `xml:"management-address,omitempty"`
	PortDescription    *string                         `xml:"port-description,omitempty"`
	SystemCapabilities *string                         `xml:"system-capabilities,omitempty"`
	SystemDescription  *string                         `xml:"system-description,omitempty"`
	SystemName         *string                         `xml:"system-name,omitempty"`
	Misc               []generic.Xml                   `xml:",any"`
	MiscAttributes     []xml.Attr                      `xml:",any,attr"`
}
type optionTlvsManagementAddressXml struct {
	Enabled        *string                                        `xml:"enabled,omitempty"`
	Iplist         *optionTlvsManagementAddressIplistContainerXml `xml:"iplist,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type optionTlvsManagementAddressIplistContainerXml struct {
	Entries []optionTlvsManagementAddressIplistXml `xml:"entry"`
}
type optionTlvsManagementAddressIplistXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Interface      *string       `xml:"interface,omitempty"`
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Mode = s.Mode
	if s.OptionTlvs != nil {
		var obj optionTlvsXml
		obj.MarshalFromObject(*s.OptionTlvs)
		o.OptionTlvs = &obj
	}
	o.SnmpSyslogNotification = util.YesNo(s.SnmpSyslogNotification, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var optionTlvsVal *OptionTlvs
	if o.OptionTlvs != nil {
		obj, err := o.OptionTlvs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		optionTlvsVal = obj
	}

	result := &Entry{
		Name:                   o.Name,
		Mode:                   o.Mode,
		OptionTlvs:             optionTlvsVal,
		SnmpSyslogNotification: util.AsBool(o.SnmpSyslogNotification, nil),
		Misc:                   o.Misc,
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *optionTlvsXml) MarshalFromObject(s OptionTlvs) {
	if s.ManagementAddress != nil {
		var obj optionTlvsManagementAddressXml
		obj.MarshalFromObject(*s.ManagementAddress)
		o.ManagementAddress = &obj
	}
	o.PortDescription = util.YesNo(s.PortDescription, nil)
	o.SystemCapabilities = util.YesNo(s.SystemCapabilities, nil)
	o.SystemDescription = util.YesNo(s.SystemDescription, nil)
	o.SystemName = util.YesNo(s.SystemName, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o optionTlvsXml) UnmarshalToObject() (*OptionTlvs, error) {
	var managementAddressVal *OptionTlvsManagementAddress
	if o.ManagementAddress != nil {
		obj, err := o.ManagementAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		managementAddressVal = obj
	}

	result := &OptionTlvs{
		ManagementAddress:  managementAddressVal,
		PortDescription:    util.AsBool(o.PortDescription, nil),
		SystemCapabilities: util.AsBool(o.SystemCapabilities, nil),
		SystemDescription:  util.AsBool(o.SystemDescription, nil),
		SystemName:         util.AsBool(o.SystemName, nil),
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *optionTlvsManagementAddressXml) MarshalFromObject(s OptionTlvsManagementAddress) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Iplist != nil {
		var objs []optionTlvsManagementAddressIplistXml
		for _, elt := range s.Iplist {
			var obj optionTlvsManagementAddressIplistXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Iplist = &optionTlvsManagementAddressIplistContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o optionTlvsManagementAddressXml) UnmarshalToObject() (*OptionTlvsManagementAddress, error) {
	var iplistVal []OptionTlvsManagementAddressIplist
	if o.Iplist != nil {
		for _, elt := range o.Iplist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			iplistVal = append(iplistVal, *obj)
		}
	}

	result := &OptionTlvsManagementAddress{
		Enabled:        util.AsBool(o.Enabled, nil),
		Iplist:         iplistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *optionTlvsManagementAddressIplistXml) MarshalFromObject(s OptionTlvsManagementAddressIplist) {
	o.Name = s.Name
	o.Interface = s.Interface
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o optionTlvsManagementAddressIplistXml) UnmarshalToObject() (*OptionTlvsManagementAddressIplist, error) {

	result := &OptionTlvsManagementAddressIplist{
		Name:           o.Name,
		Interface:      o.Interface,
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "mode" || v == "Mode" {
		return e.Mode, nil
	}
	if v == "option_tlvs" || v == "OptionTlvs" {
		return e.OptionTlvs, nil
	}
	if v == "snmp_syslog_notification" || v == "SnmpSyslogNotification" {
		return e.SnmpSyslogNotification, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

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
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !o.OptionTlvs.matches(other.OptionTlvs) {
		return false
	}
	if !util.BoolsMatch(o.SnmpSyslogNotification, other.SnmpSyslogNotification) {
		return false
	}

	return true
}

func (o *OptionTlvs) matches(other *OptionTlvs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.ManagementAddress.matches(other.ManagementAddress) {
		return false
	}
	if !util.BoolsMatch(o.PortDescription, other.PortDescription) {
		return false
	}
	if !util.BoolsMatch(o.SystemCapabilities, other.SystemCapabilities) {
		return false
	}
	if !util.BoolsMatch(o.SystemDescription, other.SystemDescription) {
		return false
	}
	if !util.BoolsMatch(o.SystemName, other.SystemName) {
		return false
	}

	return true
}

func (o *OptionTlvsManagementAddress) matches(other *OptionTlvsManagementAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if len(o.Iplist) != len(other.Iplist) {
		return false
	}
	for idx := range o.Iplist {
		if !o.Iplist[idx].matches(&other.Iplist[idx]) {
			return false
		}
	}

	return true
}

func (o *OptionTlvsManagementAddressIplist) matches(other *OptionTlvsManagementAddressIplist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
