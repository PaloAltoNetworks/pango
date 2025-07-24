package loopback

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
	suffix = []string{"network", "interface", "loopback", "units", "$name"}
)

type Entry struct {
	Name                       string
	AdjustTcpMss               *AdjustTcpMss
	Comment                    *string
	InterfaceManagementProfile *string
	Ip                         []Ip
	Ipv6                       *Ipv6
	Mtu                        *int64
	NetflowProfile             *string
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}
type AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ip struct {
	Name           string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6 struct {
	Address        []Ipv6Address
	Enabled        *bool
	InterfaceId    *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6Address struct {
	Name              string
	EnableOnInterface *bool
	Prefix            *Ipv6AddressPrefix
	Anycast           *Ipv6AddressAnycast
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Ipv6AddressPrefix struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6AddressAnycast struct {
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
	XMLName                    xml.Name         `xml:"entry"`
	Name                       string           `xml:"name,attr"`
	AdjustTcpMss               *adjustTcpMssXml `xml:"adjust-tcp-mss,omitempty"`
	Comment                    *string          `xml:"comment,omitempty"`
	InterfaceManagementProfile *string          `xml:"interface-management-profile,omitempty"`
	Ip                         *ipContainerXml  `xml:"ip,omitempty"`
	Ipv6                       *ipv6Xml         `xml:"ipv6,omitempty"`
	Mtu                        *int64           `xml:"mtu,omitempty"`
	NetflowProfile             *string          `xml:"netflow-profile,omitempty"`
	Misc                       []generic.Xml    `xml:",any"`
	MiscAttributes             []xml.Attr       `xml:",any,attr"`
}
type adjustTcpMssXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64        `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64        `xml:"ipv6-mss-adjustment,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ipContainerXml struct {
	Entries []ipXml `xml:"entry"`
}
type ipXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml struct {
	Address        *ipv6AddressContainerXml `xml:"address,omitempty"`
	Enabled        *string                  `xml:"enabled,omitempty"`
	InterfaceId    *string                  `xml:"interface-id,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type ipv6AddressContainerXml struct {
	Entries []ipv6AddressXml `xml:"entry"`
}
type ipv6AddressXml struct {
	XMLName           xml.Name               `xml:"entry"`
	Name              string                 `xml:"name,attr"`
	EnableOnInterface *string                `xml:"enable-on-interface,omitempty"`
	Prefix            *ipv6AddressPrefixXml  `xml:"prefix,omitempty"`
	Anycast           *ipv6AddressAnycastXml `xml:"anycast,omitempty"`
	Misc              []generic.Xml          `xml:",any"`
	MiscAttributes    []xml.Attr             `xml:",any,attr"`
}
type ipv6AddressPrefixXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6AddressAnycastXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AdjustTcpMss != nil {
		var obj adjustTcpMssXml
		obj.MarshalFromObject(*s.AdjustTcpMss)
		o.AdjustTcpMss = &obj
	}
	o.Comment = s.Comment
	o.InterfaceManagementProfile = s.InterfaceManagementProfile
	if s.Ip != nil {
		var objs []ipXml
		for _, elt := range s.Ip {
			var obj ipXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Ip = &ipContainerXml{Entries: objs}
	}
	if s.Ipv6 != nil {
		var obj ipv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Mtu = s.Mtu
	o.NetflowProfile = s.NetflowProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var adjustTcpMssVal *AdjustTcpMss
	if o.AdjustTcpMss != nil {
		obj, err := o.AdjustTcpMss.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		adjustTcpMssVal = obj
	}
	var ipVal []Ip
	if o.Ip != nil {
		for _, elt := range o.Ip.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ipVal = append(ipVal, *obj)
		}
	}
	var ipv6Val *Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &Entry{
		Name:                       o.Name,
		AdjustTcpMss:               adjustTcpMssVal,
		Comment:                    o.Comment,
		InterfaceManagementProfile: o.InterfaceManagementProfile,
		Ip:                         ipVal,
		Ipv6:                       ipv6Val,
		Mtu:                        o.Mtu,
		NetflowProfile:             o.NetflowProfile,
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *adjustTcpMssXml) MarshalFromObject(s AdjustTcpMss) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Ipv4MssAdjustment = s.Ipv4MssAdjustment
	o.Ipv6MssAdjustment = s.Ipv6MssAdjustment
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o adjustTcpMssXml) UnmarshalToObject() (*AdjustTcpMss, error) {

	result := &AdjustTcpMss{
		Enable:            util.AsBool(o.Enable, nil),
		Ipv4MssAdjustment: o.Ipv4MssAdjustment,
		Ipv6MssAdjustment: o.Ipv6MssAdjustment,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipXml) MarshalFromObject(s Ip) {
	o.Name = s.Name
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipXml) UnmarshalToObject() (*Ip, error) {

	result := &Ip{
		Name:           o.Name,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6Xml) MarshalFromObject(s Ipv6) {
	if s.Address != nil {
		var objs []ipv6AddressXml
		for _, elt := range s.Address {
			var obj ipv6AddressXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Address = &ipv6AddressContainerXml{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.InterfaceId = s.InterfaceId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6Xml) UnmarshalToObject() (*Ipv6, error) {
	var addressVal []Ipv6Address
	if o.Address != nil {
		for _, elt := range o.Address.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			addressVal = append(addressVal, *obj)
		}
	}

	result := &Ipv6{
		Address:        addressVal,
		Enabled:        util.AsBool(o.Enabled, nil),
		InterfaceId:    o.InterfaceId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressXml) MarshalFromObject(s Ipv6Address) {
	o.Name = s.Name
	o.EnableOnInterface = util.YesNo(s.EnableOnInterface, nil)
	if s.Prefix != nil {
		var obj ipv6AddressPrefixXml
		obj.MarshalFromObject(*s.Prefix)
		o.Prefix = &obj
	}
	if s.Anycast != nil {
		var obj ipv6AddressAnycastXml
		obj.MarshalFromObject(*s.Anycast)
		o.Anycast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressXml) UnmarshalToObject() (*Ipv6Address, error) {
	var prefixVal *Ipv6AddressPrefix
	if o.Prefix != nil {
		obj, err := o.Prefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		prefixVal = obj
	}
	var anycastVal *Ipv6AddressAnycast
	if o.Anycast != nil {
		obj, err := o.Anycast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anycastVal = obj
	}

	result := &Ipv6Address{
		Name:              o.Name,
		EnableOnInterface: util.AsBool(o.EnableOnInterface, nil),
		Prefix:            prefixVal,
		Anycast:           anycastVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressPrefixXml) MarshalFromObject(s Ipv6AddressPrefix) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressPrefixXml) UnmarshalToObject() (*Ipv6AddressPrefix, error) {

	result := &Ipv6AddressPrefix{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6AddressAnycastXml) MarshalFromObject(s Ipv6AddressAnycast) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6AddressAnycastXml) UnmarshalToObject() (*Ipv6AddressAnycast, error) {

	result := &Ipv6AddressAnycast{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "adjust_tcp_mss" || v == "AdjustTcpMss" {
		return e.AdjustTcpMss, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "interface_management_profile" || v == "InterfaceManagementProfile" {
		return e.InterfaceManagementProfile, nil
	}
	if v == "ip" || v == "Ip" {
		return e.Ip, nil
	}
	if v == "ip|LENGTH" || v == "Ip|LENGTH" {
		return int64(len(e.Ip)), nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "mtu" || v == "Mtu" {
		return e.Mtu, nil
	}
	if v == "netflow_profile" || v == "NetflowProfile" {
		return e.NetflowProfile, nil
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
	if !o.AdjustTcpMss.matches(other.AdjustTcpMss) {
		return false
	}
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !util.StringsMatch(o.InterfaceManagementProfile, other.InterfaceManagementProfile) {
		return false
	}
	if len(o.Ip) != len(other.Ip) {
		return false
	}
	for idx := range o.Ip {
		if !o.Ip[idx].matches(&other.Ip[idx]) {
			return false
		}
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}
	if !util.Ints64Match(o.Mtu, other.Mtu) {
		return false
	}
	if !util.StringsMatch(o.NetflowProfile, other.NetflowProfile) {
		return false
	}

	return true
}

func (o *AdjustTcpMss) matches(other *AdjustTcpMss) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Ipv4MssAdjustment, other.Ipv4MssAdjustment) {
		return false
	}
	if !util.Ints64Match(o.Ipv6MssAdjustment, other.Ipv6MssAdjustment) {
		return false
	}

	return true
}

func (o *Ip) matches(other *Ip) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *Ipv6) matches(other *Ipv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Address) != len(other.Address) {
		return false
	}
	for idx := range o.Address {
		if !o.Address[idx].matches(&other.Address[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.StringsMatch(o.InterfaceId, other.InterfaceId) {
		return false
	}

	return true
}

func (o *Ipv6Address) matches(other *Ipv6Address) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.EnableOnInterface, other.EnableOnInterface) {
		return false
	}
	if !o.Prefix.matches(other.Prefix) {
		return false
	}
	if !o.Anycast.matches(other.Anycast) {
		return false
	}

	return true
}

func (o *Ipv6AddressPrefix) matches(other *Ipv6AddressPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6AddressAnycast) matches(other *Ipv6AddressAnycast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
