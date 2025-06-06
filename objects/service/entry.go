package service

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"service", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Protocol        *Protocol
	Tag             []string
	Misc            []generic.Xml
}
type Protocol struct {
	Tcp  *ProtocolTcp
	Udp  *ProtocolUdp
	Misc []generic.Xml
}
type ProtocolTcp struct {
	Override   *ProtocolTcpOverride
	Port       *string
	SourcePort *string
	Misc       []generic.Xml
}
type ProtocolTcpOverride struct {
	HalfcloseTimeout *int64
	Timeout          *int64
	TimewaitTimeout  *int64
	Misc             []generic.Xml
}
type ProtocolUdp struct {
	Override   *ProtocolUdpOverride
	Port       *string
	SourcePort *string
	Misc       []generic.Xml
}
type ProtocolUdpOverride struct {
	Timeout *int64
	Misc    []generic.Xml
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
	Description     *string          `xml:"description,omitempty"`
	DisableOverride *string          `xml:"disable-override,omitempty"`
	Protocol        *protocolXml     `xml:"protocol,omitempty"`
	Tag             *util.MemberType `xml:"tag,omitempty"`
	Misc            []generic.Xml    `xml:",any"`
}
type protocolXml struct {
	Tcp  *protocolTcpXml `xml:"tcp,omitempty"`
	Udp  *protocolUdpXml `xml:"udp,omitempty"`
	Misc []generic.Xml   `xml:",any"`
}
type protocolTcpXml struct {
	Override   *protocolTcpOverrideXml `xml:"override>yes,omitempty"`
	Port       *string                 `xml:"port,omitempty"`
	SourcePort *string                 `xml:"source-port,omitempty"`
	Misc       []generic.Xml           `xml:",any"`
}
type protocolTcpOverrideXml struct {
	HalfcloseTimeout *int64        `xml:"halfclose-timeout,omitempty"`
	Timeout          *int64        `xml:"timeout,omitempty"`
	TimewaitTimeout  *int64        `xml:"timewait-timeout,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type protocolUdpXml struct {
	Override   *protocolUdpOverrideXml `xml:"override>yes,omitempty"`
	Port       *string                 `xml:"port,omitempty"`
	SourcePort *string                 `xml:"source-port,omitempty"`
	Misc       []generic.Xml           `xml:",any"`
}
type protocolUdpOverrideXml struct {
	Timeout *int64        `xml:"timeout,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.Protocol != nil {
		var obj protocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var protocolVal *Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		DisableOverride: o.DisableOverride,
		Protocol:        protocolVal,
		Tag:             tagVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *protocolXml) MarshalFromObject(s Protocol) {
	if s.Tcp != nil {
		var obj protocolTcpXml
		obj.MarshalFromObject(*s.Tcp)
		o.Tcp = &obj
	}
	if s.Udp != nil {
		var obj protocolUdpXml
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
}

func (o protocolXml) UnmarshalToObject() (*Protocol, error) {
	var tcpVal *ProtocolTcp
	if o.Tcp != nil {
		obj, err := o.Tcp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpVal = obj
	}
	var udpVal *ProtocolUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &Protocol{
		Tcp:  tcpVal,
		Udp:  udpVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *protocolTcpXml) MarshalFromObject(s ProtocolTcp) {
	if s.Override != nil {
		var obj protocolTcpOverrideXml
		obj.MarshalFromObject(*s.Override)
		o.Override = &obj
	}
	o.Port = s.Port
	o.SourcePort = s.SourcePort
	o.Misc = s.Misc
}

func (o protocolTcpXml) UnmarshalToObject() (*ProtocolTcp, error) {
	var overrideVal *ProtocolTcpOverride
	if o.Override != nil {
		obj, err := o.Override.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		overrideVal = obj
	}

	result := &ProtocolTcp{
		Override:   overrideVal,
		Port:       o.Port,
		SourcePort: o.SourcePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *protocolTcpOverrideXml) MarshalFromObject(s ProtocolTcpOverride) {
	o.HalfcloseTimeout = s.HalfcloseTimeout
	o.Timeout = s.Timeout
	o.TimewaitTimeout = s.TimewaitTimeout
	o.Misc = s.Misc
}

func (o protocolTcpOverrideXml) UnmarshalToObject() (*ProtocolTcpOverride, error) {

	result := &ProtocolTcpOverride{
		HalfcloseTimeout: o.HalfcloseTimeout,
		Timeout:          o.Timeout,
		TimewaitTimeout:  o.TimewaitTimeout,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolUdpXml) MarshalFromObject(s ProtocolUdp) {
	if s.Override != nil {
		var obj protocolUdpOverrideXml
		obj.MarshalFromObject(*s.Override)
		o.Override = &obj
	}
	o.Port = s.Port
	o.SourcePort = s.SourcePort
	o.Misc = s.Misc
}

func (o protocolUdpXml) UnmarshalToObject() (*ProtocolUdp, error) {
	var overrideVal *ProtocolUdpOverride
	if o.Override != nil {
		obj, err := o.Override.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		overrideVal = obj
	}

	result := &ProtocolUdp{
		Override:   overrideVal,
		Port:       o.Port,
		SourcePort: o.SourcePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *protocolUdpOverrideXml) MarshalFromObject(s ProtocolUdpOverride) {
	o.Timeout = s.Timeout
	o.Misc = s.Misc
}

func (o protocolUdpOverrideXml) UnmarshalToObject() (*ProtocolUdpOverride, error) {

	result := &ProtocolUdpOverride{
		Timeout: o.Timeout,
		Misc:    o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}

	return true
}

func (o *Protocol) matches(other *Protocol) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Tcp.matches(other.Tcp) {
		return false
	}
	if !o.Udp.matches(other.Udp) {
		return false
	}

	return true
}

func (o *ProtocolTcp) matches(other *ProtocolTcp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Override.matches(other.Override) {
		return false
	}
	if !util.StringsMatch(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.SourcePort, other.SourcePort) {
		return false
	}

	return true
}

func (o *ProtocolTcpOverride) matches(other *ProtocolTcpOverride) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.HalfcloseTimeout, other.HalfcloseTimeout) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !util.Ints64Match(o.TimewaitTimeout, other.TimewaitTimeout) {
		return false
	}

	return true
}

func (o *ProtocolUdp) matches(other *ProtocolUdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Override.matches(other.Override) {
		return false
	}
	if !util.StringsMatch(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.SourcePort, other.SourcePort) {
		return false
	}

	return true
}

func (o *ProtocolUdpOverride) matches(other *ProtocolUdpOverride) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
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
