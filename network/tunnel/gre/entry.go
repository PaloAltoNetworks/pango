package gre

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
	suffix = []string{"network", "tunnel", "gre", "$name"}
)

type Entry struct {
	Name            string
	CopyTos         *bool
	Disabled        *bool
	Erspan          *bool
	KeepAlive       *KeepAlive
	LocalAddress    *LocalAddress
	PeerAddress     *PeerAddress
	Ttl             *int64
	TunnelInterface *string
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type KeepAlive struct {
	Enable         *bool
	HoldTimer      *int64
	Interval       *int64
	Retry          *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type LocalAddress struct {
	Interface      *string
	FloatingIp     *string
	Ip             *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PeerAddress struct {
	Ip             *string
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
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	CopyTos         *string          `xml:"copy-tos,omitempty"`
	Disabled        *string          `xml:"disabled,omitempty"`
	Erspan          *string          `xml:"erspan,omitempty"`
	KeepAlive       *keepAliveXml    `xml:"keep-alive,omitempty"`
	LocalAddress    *localAddressXml `xml:"local-address,omitempty"`
	PeerAddress     *peerAddressXml  `xml:"peer-address,omitempty"`
	Ttl             *int64           `xml:"ttl,omitempty"`
	TunnelInterface *string          `xml:"tunnel-interface,omitempty"`
	Misc            []generic.Xml    `xml:",any"`
	MiscAttributes  []xml.Attr       `xml:",any,attr"`
}
type keepAliveXml struct {
	Enable         *string       `xml:"enable,omitempty"`
	HoldTimer      *int64        `xml:"hold-timer,omitempty"`
	Interval       *int64        `xml:"interval,omitempty"`
	Retry          *int64        `xml:"retry,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type localAddressXml struct {
	Interface      *string       `xml:"interface,omitempty"`
	FloatingIp     *string       `xml:"floating-ip,omitempty"`
	Ip             *string       `xml:"ip,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type peerAddressXml struct {
	Ip             *string       `xml:"ip,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.CopyTos = util.YesNo(s.CopyTos, nil)
	o.Disabled = util.YesNo(s.Disabled, nil)
	o.Erspan = util.YesNo(s.Erspan, nil)
	if s.KeepAlive != nil {
		var obj keepAliveXml
		obj.MarshalFromObject(*s.KeepAlive)
		o.KeepAlive = &obj
	}
	if s.LocalAddress != nil {
		var obj localAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.PeerAddress != nil {
		var obj peerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	o.Ttl = s.Ttl
	o.TunnelInterface = s.TunnelInterface
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var keepAliveVal *KeepAlive
	if o.KeepAlive != nil {
		obj, err := o.KeepAlive.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		keepAliveVal = obj
	}
	var localAddressVal *LocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *PeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		CopyTos:         util.AsBool(o.CopyTos, nil),
		Disabled:        util.AsBool(o.Disabled, nil),
		Erspan:          util.AsBool(o.Erspan, nil),
		KeepAlive:       keepAliveVal,
		LocalAddress:    localAddressVal,
		PeerAddress:     peerAddressVal,
		Ttl:             o.Ttl,
		TunnelInterface: o.TunnelInterface,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *keepAliveXml) MarshalFromObject(s KeepAlive) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.HoldTimer = s.HoldTimer
	o.Interval = s.Interval
	o.Retry = s.Retry
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o keepAliveXml) UnmarshalToObject() (*KeepAlive, error) {

	result := &KeepAlive{
		Enable:         util.AsBool(o.Enable, nil),
		HoldTimer:      o.HoldTimer,
		Interval:       o.Interval,
		Retry:          o.Retry,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressXml) MarshalFromObject(s LocalAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressXml) UnmarshalToObject() (*LocalAddress, error) {

	result := &LocalAddress{
		Interface:      o.Interface,
		FloatingIp:     o.FloatingIp,
		Ip:             o.Ip,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *peerAddressXml) MarshalFromObject(s PeerAddress) {
	o.Ip = s.Ip
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o peerAddressXml) UnmarshalToObject() (*PeerAddress, error) {

	result := &PeerAddress{
		Ip:             o.Ip,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "copy_tos" || v == "CopyTos" {
		return e.CopyTos, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "erspan" || v == "Erspan" {
		return e.Erspan, nil
	}
	if v == "keep_alive" || v == "KeepAlive" {
		return e.KeepAlive, nil
	}
	if v == "local_address" || v == "LocalAddress" {
		return e.LocalAddress, nil
	}
	if v == "peer_address" || v == "PeerAddress" {
		return e.PeerAddress, nil
	}
	if v == "ttl" || v == "Ttl" {
		return e.Ttl, nil
	}
	if v == "tunnel_interface" || v == "TunnelInterface" {
		return e.TunnelInterface, nil
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
	if !util.BoolsMatch(o.CopyTos, other.CopyTos) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !util.BoolsMatch(o.Erspan, other.Erspan) {
		return false
	}
	if !o.KeepAlive.matches(other.KeepAlive) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !o.PeerAddress.matches(other.PeerAddress) {
		return false
	}
	if !util.Ints64Match(o.Ttl, other.Ttl) {
		return false
	}
	if !util.StringsMatch(o.TunnelInterface, other.TunnelInterface) {
		return false
	}

	return true
}

func (o *KeepAlive) matches(other *KeepAlive) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.HoldTimer, other.HoldTimer) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Retry, other.Retry) {
		return false
	}

	return true
}

func (o *LocalAddress) matches(other *LocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.FloatingIp, other.FloatingIp) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *PeerAddress) matches(other *PeerAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
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
