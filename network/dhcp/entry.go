package dhcp

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
	suffix = []string{"network", "dhcp", "interface", "$name"}
)

type Entry struct {
	Name           string
	Relay          *Relay
	Server         *Server
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Relay struct {
	Ip             *RelayIp
	Ipv6           *RelayIpv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RelayIp struct {
	Enabled        *bool
	Server         []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RelayIpv6 struct {
	Enabled        *bool
	Server         []RelayIpv6Server
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RelayIpv6Server struct {
	Name           string
	Interface      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Server struct {
	IpPool         []string
	Mode           *string
	Option         *ServerOption
	ProbeIp        *bool
	Reserved       []ServerReserved
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOption struct {
	Dns            *ServerOptionDns
	DnsSuffix      *string
	Gateway        *string
	Inheritance    *ServerOptionInheritance
	Lease          *ServerOptionLease
	Nis            *ServerOptionNis
	Ntp            *ServerOptionNtp
	Pop3Server     *string
	SmtpServer     *string
	SubnetMask     *string
	UserDefined    []ServerOptionUserDefined
	Wins           *ServerOptionWins
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionDns struct {
	Primary        *string
	Secondary      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionInheritance struct {
	Source         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionLease struct {
	Timeout        *int64
	Unlimited      *ServerOptionLeaseUnlimited
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionLeaseUnlimited struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionNis struct {
	Primary        *string
	Secondary      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionNtp struct {
	Primary        *string
	Secondary      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerOptionUserDefined struct {
	Name                  string
	Code                  *int64
	VendorClassIdentifier *string
	Inherited             *bool
	Ip                    []string
	Ascii                 []string
	Hex                   []string
	Misc                  []generic.Xml
	MiscAttributes        []xml.Attr
}
type ServerOptionWins struct {
	Primary        *string
	Secondary      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ServerReserved struct {
	Name           string
	Mac            *string
	Description    *string
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Relay          *relayXml     `xml:"relay,omitempty"`
	Server         *serverXml    `xml:"server,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type relayXml struct {
	Ip             *relayIpXml   `xml:"ip,omitempty"`
	Ipv6           *relayIpv6Xml `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type relayIpXml struct {
	Enabled        *string          `xml:"enabled,omitempty"`
	Server         *util.MemberType `xml:"server,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type relayIpv6Xml struct {
	Enabled        *string                      `xml:"enabled,omitempty"`
	Server         *relayIpv6ServerContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type relayIpv6ServerContainerXml struct {
	Entries []relayIpv6ServerXml `xml:"entry"`
}
type relayIpv6ServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Interface      *string       `xml:"interface,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverXml struct {
	IpPool         *util.MemberType            `xml:"ip-pool,omitempty"`
	Mode           *string                     `xml:"mode,omitempty"`
	Option         *serverOptionXml            `xml:"option,omitempty"`
	ProbeIp        *string                     `xml:"probe-ip,omitempty"`
	Reserved       *serverReservedContainerXml `xml:"reserved,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type serverOptionXml struct {
	Dns            *serverOptionDnsXml                  `xml:"dns,omitempty"`
	DnsSuffix      *string                              `xml:"dns-suffix,omitempty"`
	Gateway        *string                              `xml:"gateway,omitempty"`
	Inheritance    *serverOptionInheritanceXml          `xml:"inheritance,omitempty"`
	Lease          *serverOptionLeaseXml                `xml:"lease,omitempty"`
	Nis            *serverOptionNisXml                  `xml:"nis,omitempty"`
	Ntp            *serverOptionNtpXml                  `xml:"ntp,omitempty"`
	Pop3Server     *string                              `xml:"pop3-server,omitempty"`
	SmtpServer     *string                              `xml:"smtp-server,omitempty"`
	SubnetMask     *string                              `xml:"subnet-mask,omitempty"`
	UserDefined    *serverOptionUserDefinedContainerXml `xml:"user-defined,omitempty"`
	Wins           *serverOptionWinsXml                 `xml:"wins,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type serverOptionDnsXml struct {
	Primary        *string       `xml:"primary,omitempty"`
	Secondary      *string       `xml:"secondary,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverOptionInheritanceXml struct {
	Source         *string       `xml:"source,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverOptionLeaseXml struct {
	Timeout        *int64                         `xml:"timeout,omitempty"`
	Unlimited      *serverOptionLeaseUnlimitedXml `xml:"unlimited,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type serverOptionLeaseUnlimitedXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverOptionNisXml struct {
	Primary        *string       `xml:"primary,omitempty"`
	Secondary      *string       `xml:"secondary,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverOptionNtpXml struct {
	Primary        *string       `xml:"primary,omitempty"`
	Secondary      *string       `xml:"secondary,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverOptionUserDefinedContainerXml struct {
	Entries []serverOptionUserDefinedXml `xml:"entry"`
}
type serverOptionUserDefinedXml struct {
	XMLName               xml.Name         `xml:"entry"`
	Name                  string           `xml:"name,attr"`
	Code                  *int64           `xml:"code,omitempty"`
	VendorClassIdentifier *string          `xml:"vendor-class-identifier,omitempty"`
	Inherited             *string          `xml:"inherited,omitempty"`
	Ip                    *util.MemberType `xml:"ip,omitempty"`
	Ascii                 *util.MemberType `xml:"ascii,omitempty"`
	Hex                   *util.MemberType `xml:"hex,omitempty"`
	Misc                  []generic.Xml    `xml:",any"`
	MiscAttributes        []xml.Attr       `xml:",any,attr"`
}
type serverOptionWinsXml struct {
	Primary        *string       `xml:"primary,omitempty"`
	Secondary      *string       `xml:"secondary,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type serverReservedContainerXml struct {
	Entries []serverReservedXml `xml:"entry"`
}
type serverReservedXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Mac            *string       `xml:"mac,omitempty"`
	Description    *string       `xml:"description,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Relay != nil {
		var obj relayXml
		obj.MarshalFromObject(*s.Relay)
		o.Relay = &obj
	}
	if s.Server != nil {
		var obj serverXml
		obj.MarshalFromObject(*s.Server)
		o.Server = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var relayVal *Relay
	if o.Relay != nil {
		obj, err := o.Relay.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		relayVal = obj
	}
	var serverVal *Server
	if o.Server != nil {
		obj, err := o.Server.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Relay:          relayVal,
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *relayXml) MarshalFromObject(s Relay) {
	if s.Ip != nil {
		var obj relayIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	if s.Ipv6 != nil {
		var obj relayIpv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o relayXml) UnmarshalToObject() (*Relay, error) {
	var ipVal *RelayIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}
	var ipv6Val *RelayIpv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &Relay{
		Ip:             ipVal,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *relayIpXml) MarshalFromObject(s RelayIp) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Server != nil {
		o.Server = util.StrToMem(s.Server)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o relayIpXml) UnmarshalToObject() (*RelayIp, error) {
	var serverVal []string
	if o.Server != nil {
		serverVal = util.MemToStr(o.Server)
	}

	result := &RelayIp{
		Enabled:        util.AsBool(o.Enabled, nil),
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *relayIpv6Xml) MarshalFromObject(s RelayIpv6) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Server != nil {
		var objs []relayIpv6ServerXml
		for _, elt := range s.Server {
			var obj relayIpv6ServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &relayIpv6ServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o relayIpv6Xml) UnmarshalToObject() (*RelayIpv6, error) {
	var serverVal []RelayIpv6Server
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &RelayIpv6{
		Enabled:        util.AsBool(o.Enabled, nil),
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *relayIpv6ServerXml) MarshalFromObject(s RelayIpv6Server) {
	o.Name = s.Name
	o.Interface = s.Interface
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o relayIpv6ServerXml) UnmarshalToObject() (*RelayIpv6Server, error) {

	result := &RelayIpv6Server{
		Name:           o.Name,
		Interface:      o.Interface,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverXml) MarshalFromObject(s Server) {
	if s.IpPool != nil {
		o.IpPool = util.StrToMem(s.IpPool)
	}
	o.Mode = s.Mode
	if s.Option != nil {
		var obj serverOptionXml
		obj.MarshalFromObject(*s.Option)
		o.Option = &obj
	}
	o.ProbeIp = util.YesNo(s.ProbeIp, nil)
	if s.Reserved != nil {
		var objs []serverReservedXml
		for _, elt := range s.Reserved {
			var obj serverReservedXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Reserved = &serverReservedContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverXml) UnmarshalToObject() (*Server, error) {
	var ipPoolVal []string
	if o.IpPool != nil {
		ipPoolVal = util.MemToStr(o.IpPool)
	}
	var optionVal *ServerOption
	if o.Option != nil {
		obj, err := o.Option.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		optionVal = obj
	}
	var reservedVal []ServerReserved
	if o.Reserved != nil {
		for _, elt := range o.Reserved.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			reservedVal = append(reservedVal, *obj)
		}
	}

	result := &Server{
		IpPool:         ipPoolVal,
		Mode:           o.Mode,
		Option:         optionVal,
		ProbeIp:        util.AsBool(o.ProbeIp, nil),
		Reserved:       reservedVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionXml) MarshalFromObject(s ServerOption) {
	if s.Dns != nil {
		var obj serverOptionDnsXml
		obj.MarshalFromObject(*s.Dns)
		o.Dns = &obj
	}
	o.DnsSuffix = s.DnsSuffix
	o.Gateway = s.Gateway
	if s.Inheritance != nil {
		var obj serverOptionInheritanceXml
		obj.MarshalFromObject(*s.Inheritance)
		o.Inheritance = &obj
	}
	if s.Lease != nil {
		var obj serverOptionLeaseXml
		obj.MarshalFromObject(*s.Lease)
		o.Lease = &obj
	}
	if s.Nis != nil {
		var obj serverOptionNisXml
		obj.MarshalFromObject(*s.Nis)
		o.Nis = &obj
	}
	if s.Ntp != nil {
		var obj serverOptionNtpXml
		obj.MarshalFromObject(*s.Ntp)
		o.Ntp = &obj
	}
	o.Pop3Server = s.Pop3Server
	o.SmtpServer = s.SmtpServer
	o.SubnetMask = s.SubnetMask
	if s.UserDefined != nil {
		var objs []serverOptionUserDefinedXml
		for _, elt := range s.UserDefined {
			var obj serverOptionUserDefinedXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.UserDefined = &serverOptionUserDefinedContainerXml{Entries: objs}
	}
	if s.Wins != nil {
		var obj serverOptionWinsXml
		obj.MarshalFromObject(*s.Wins)
		o.Wins = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionXml) UnmarshalToObject() (*ServerOption, error) {
	var dnsVal *ServerOptionDns
	if o.Dns != nil {
		obj, err := o.Dns.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsVal = obj
	}
	var inheritanceVal *ServerOptionInheritance
	if o.Inheritance != nil {
		obj, err := o.Inheritance.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritanceVal = obj
	}
	var leaseVal *ServerOptionLease
	if o.Lease != nil {
		obj, err := o.Lease.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		leaseVal = obj
	}
	var nisVal *ServerOptionNis
	if o.Nis != nil {
		obj, err := o.Nis.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nisVal = obj
	}
	var ntpVal *ServerOptionNtp
	if o.Ntp != nil {
		obj, err := o.Ntp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ntpVal = obj
	}
	var userDefinedVal []ServerOptionUserDefined
	if o.UserDefined != nil {
		for _, elt := range o.UserDefined.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			userDefinedVal = append(userDefinedVal, *obj)
		}
	}
	var winsVal *ServerOptionWins
	if o.Wins != nil {
		obj, err := o.Wins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		winsVal = obj
	}

	result := &ServerOption{
		Dns:            dnsVal,
		DnsSuffix:      o.DnsSuffix,
		Gateway:        o.Gateway,
		Inheritance:    inheritanceVal,
		Lease:          leaseVal,
		Nis:            nisVal,
		Ntp:            ntpVal,
		Pop3Server:     o.Pop3Server,
		SmtpServer:     o.SmtpServer,
		SubnetMask:     o.SubnetMask,
		UserDefined:    userDefinedVal,
		Wins:           winsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionDnsXml) MarshalFromObject(s ServerOptionDns) {
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionDnsXml) UnmarshalToObject() (*ServerOptionDns, error) {

	result := &ServerOptionDns{
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionInheritanceXml) MarshalFromObject(s ServerOptionInheritance) {
	o.Source = s.Source
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionInheritanceXml) UnmarshalToObject() (*ServerOptionInheritance, error) {

	result := &ServerOptionInheritance{
		Source:         o.Source,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionLeaseXml) MarshalFromObject(s ServerOptionLease) {
	o.Timeout = s.Timeout
	if s.Unlimited != nil {
		var obj serverOptionLeaseUnlimitedXml
		obj.MarshalFromObject(*s.Unlimited)
		o.Unlimited = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionLeaseXml) UnmarshalToObject() (*ServerOptionLease, error) {
	var unlimitedVal *ServerOptionLeaseUnlimited
	if o.Unlimited != nil {
		obj, err := o.Unlimited.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unlimitedVal = obj
	}

	result := &ServerOptionLease{
		Timeout:        o.Timeout,
		Unlimited:      unlimitedVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionLeaseUnlimitedXml) MarshalFromObject(s ServerOptionLeaseUnlimited) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionLeaseUnlimitedXml) UnmarshalToObject() (*ServerOptionLeaseUnlimited, error) {

	result := &ServerOptionLeaseUnlimited{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionNisXml) MarshalFromObject(s ServerOptionNis) {
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionNisXml) UnmarshalToObject() (*ServerOptionNis, error) {

	result := &ServerOptionNis{
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionNtpXml) MarshalFromObject(s ServerOptionNtp) {
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionNtpXml) UnmarshalToObject() (*ServerOptionNtp, error) {

	result := &ServerOptionNtp{
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionUserDefinedXml) MarshalFromObject(s ServerOptionUserDefined) {
	o.Name = s.Name
	o.Code = s.Code
	o.VendorClassIdentifier = s.VendorClassIdentifier
	o.Inherited = util.YesNo(s.Inherited, nil)
	if s.Ip != nil {
		o.Ip = util.StrToMem(s.Ip)
	}
	if s.Ascii != nil {
		o.Ascii = util.StrToMem(s.Ascii)
	}
	if s.Hex != nil {
		o.Hex = util.StrToMem(s.Hex)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionUserDefinedXml) UnmarshalToObject() (*ServerOptionUserDefined, error) {
	var ipVal []string
	if o.Ip != nil {
		ipVal = util.MemToStr(o.Ip)
	}
	var asciiVal []string
	if o.Ascii != nil {
		asciiVal = util.MemToStr(o.Ascii)
	}
	var hexVal []string
	if o.Hex != nil {
		hexVal = util.MemToStr(o.Hex)
	}

	result := &ServerOptionUserDefined{
		Name:                  o.Name,
		Code:                  o.Code,
		VendorClassIdentifier: o.VendorClassIdentifier,
		Inherited:             util.AsBool(o.Inherited, nil),
		Ip:                    ipVal,
		Ascii:                 asciiVal,
		Hex:                   hexVal,
		Misc:                  o.Misc,
		MiscAttributes:        o.MiscAttributes,
	}
	return result, nil
}
func (o *serverOptionWinsXml) MarshalFromObject(s ServerOptionWins) {
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverOptionWinsXml) UnmarshalToObject() (*ServerOptionWins, error) {

	result := &ServerOptionWins{
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *serverReservedXml) MarshalFromObject(s ServerReserved) {
	o.Name = s.Name
	o.Mac = s.Mac
	o.Description = s.Description
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverReservedXml) UnmarshalToObject() (*ServerReserved, error) {

	result := &ServerReserved{
		Name:           o.Name,
		Mac:            o.Mac,
		Description:    o.Description,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "relay" || v == "Relay" {
		return e.Relay, nil
	}
	if v == "server" || v == "Server" {
		return e.Server, nil
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
	if !o.Relay.matches(other.Relay) {
		return false
	}
	if !o.Server.matches(other.Server) {
		return false
	}

	return true
}

func (o *Relay) matches(other *Relay) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *RelayIp) matches(other *RelayIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Server, other.Server) {
		return false
	}

	return true
}

func (o *RelayIpv6) matches(other *RelayIpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if len(o.Server) != len(other.Server) {
		return false
	}
	for idx := range o.Server {
		if !o.Server[idx].matches(&other.Server[idx]) {
			return false
		}
	}

	return true
}

func (o *RelayIpv6Server) matches(other *RelayIpv6Server) bool {
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

	return true
}

func (o *Server) matches(other *Server) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IpPool, other.IpPool) {
		return false
	}
	if !util.StringsMatch(o.Mode, other.Mode) {
		return false
	}
	if !o.Option.matches(other.Option) {
		return false
	}
	if !util.BoolsMatch(o.ProbeIp, other.ProbeIp) {
		return false
	}
	if len(o.Reserved) != len(other.Reserved) {
		return false
	}
	for idx := range o.Reserved {
		if !o.Reserved[idx].matches(&other.Reserved[idx]) {
			return false
		}
	}

	return true
}

func (o *ServerOption) matches(other *ServerOption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dns.matches(other.Dns) {
		return false
	}
	if !util.StringsMatch(o.DnsSuffix, other.DnsSuffix) {
		return false
	}
	if !util.StringsMatch(o.Gateway, other.Gateway) {
		return false
	}
	if !o.Inheritance.matches(other.Inheritance) {
		return false
	}
	if !o.Lease.matches(other.Lease) {
		return false
	}
	if !o.Nis.matches(other.Nis) {
		return false
	}
	if !o.Ntp.matches(other.Ntp) {
		return false
	}
	if !util.StringsMatch(o.Pop3Server, other.Pop3Server) {
		return false
	}
	if !util.StringsMatch(o.SmtpServer, other.SmtpServer) {
		return false
	}
	if !util.StringsMatch(o.SubnetMask, other.SubnetMask) {
		return false
	}
	if len(o.UserDefined) != len(other.UserDefined) {
		return false
	}
	for idx := range o.UserDefined {
		if !o.UserDefined[idx].matches(&other.UserDefined[idx]) {
			return false
		}
	}
	if !o.Wins.matches(other.Wins) {
		return false
	}

	return true
}

func (o *ServerOptionDns) matches(other *ServerOptionDns) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}

func (o *ServerOptionInheritance) matches(other *ServerOptionInheritance) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Source, other.Source) {
		return false
	}

	return true
}

func (o *ServerOptionLease) matches(other *ServerOptionLease) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !o.Unlimited.matches(other.Unlimited) {
		return false
	}

	return true
}

func (o *ServerOptionLeaseUnlimited) matches(other *ServerOptionLeaseUnlimited) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ServerOptionNis) matches(other *ServerOptionNis) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}

func (o *ServerOptionNtp) matches(other *ServerOptionNtp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}

func (o *ServerOptionUserDefined) matches(other *ServerOptionUserDefined) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.Ints64Match(o.Code, other.Code) {
		return false
	}
	if !util.StringsMatch(o.VendorClassIdentifier, other.VendorClassIdentifier) {
		return false
	}
	if !util.BoolsMatch(o.Inherited, other.Inherited) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Ip, other.Ip) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Ascii, other.Ascii) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Hex, other.Hex) {
		return false
	}

	return true
}

func (o *ServerOptionWins) matches(other *ServerOptionWins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}

func (o *ServerReserved) matches(other *ServerReserved) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Mac, other.Mac) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
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
