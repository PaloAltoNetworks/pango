package interface_management

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
	suffix = []string{"network", "profiles", "interface-management-profile", "$name"}
)

type Entry struct {
	Name                    string
	Http                    *bool
	HttpOcsp                *bool
	Https                   *bool
	PermittedIp             []PermittedIp
	Ping                    *bool
	ResponsePages           *bool
	Snmp                    *bool
	Ssh                     *bool
	Telnet                  *bool
	UseridService           *bool
	UseridSyslogListenerSsl *bool
	UseridSyslogListenerUdp *bool
	Misc                    []generic.Xml
}
type PermittedIp struct {
	Name string
	Misc []generic.Xml
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
	XMLName                 xml.Name                 `xml:"entry"`
	Name                    string                   `xml:"name,attr"`
	Http                    *string                  `xml:"http,omitempty"`
	HttpOcsp                *string                  `xml:"http-ocsp,omitempty"`
	Https                   *string                  `xml:"https,omitempty"`
	PermittedIp             *permittedIpContainerXml `xml:"permitted-ip,omitempty"`
	Ping                    *string                  `xml:"ping,omitempty"`
	ResponsePages           *string                  `xml:"response-pages,omitempty"`
	Snmp                    *string                  `xml:"snmp,omitempty"`
	Ssh                     *string                  `xml:"ssh,omitempty"`
	Telnet                  *string                  `xml:"telnet,omitempty"`
	UseridService           *string                  `xml:"userid-service,omitempty"`
	UseridSyslogListenerSsl *string                  `xml:"userid-syslog-listener-ssl,omitempty"`
	UseridSyslogListenerUdp *string                  `xml:"userid-syslog-listener-udp,omitempty"`
	Misc                    []generic.Xml            `xml:",any"`
}
type permittedIpContainerXml struct {
	Entries []permittedIpXml `xml:"entry"`
}
type permittedIpXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Http = util.YesNo(s.Http, nil)
	o.HttpOcsp = util.YesNo(s.HttpOcsp, nil)
	o.Https = util.YesNo(s.Https, nil)
	if s.PermittedIp != nil {
		var objs []permittedIpXml
		for _, elt := range s.PermittedIp {
			var obj permittedIpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.PermittedIp = &permittedIpContainerXml{Entries: objs}
	}
	o.Ping = util.YesNo(s.Ping, nil)
	o.ResponsePages = util.YesNo(s.ResponsePages, nil)
	o.Snmp = util.YesNo(s.Snmp, nil)
	o.Ssh = util.YesNo(s.Ssh, nil)
	o.Telnet = util.YesNo(s.Telnet, nil)
	o.UseridService = util.YesNo(s.UseridService, nil)
	o.UseridSyslogListenerSsl = util.YesNo(s.UseridSyslogListenerSsl, nil)
	o.UseridSyslogListenerUdp = util.YesNo(s.UseridSyslogListenerUdp, nil)
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var permittedIpVal []PermittedIp
	if o.PermittedIp != nil {
		for _, elt := range o.PermittedIp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			permittedIpVal = append(permittedIpVal, *obj)
		}
	}

	result := &Entry{
		Name:                    o.Name,
		Http:                    util.AsBool(o.Http, nil),
		HttpOcsp:                util.AsBool(o.HttpOcsp, nil),
		Https:                   util.AsBool(o.Https, nil),
		PermittedIp:             permittedIpVal,
		Ping:                    util.AsBool(o.Ping, nil),
		ResponsePages:           util.AsBool(o.ResponsePages, nil),
		Snmp:                    util.AsBool(o.Snmp, nil),
		Ssh:                     util.AsBool(o.Ssh, nil),
		Telnet:                  util.AsBool(o.Telnet, nil),
		UseridService:           util.AsBool(o.UseridService, nil),
		UseridSyslogListenerSsl: util.AsBool(o.UseridSyslogListenerSsl, nil),
		UseridSyslogListenerUdp: util.AsBool(o.UseridSyslogListenerUdp, nil),
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *permittedIpXml) MarshalFromObject(s PermittedIp) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o permittedIpXml) UnmarshalToObject() (*PermittedIp, error) {

	result := &PermittedIp{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "http" || v == "Http" {
		return e.Http, nil
	}
	if v == "http_ocsp" || v == "HttpOcsp" {
		return e.HttpOcsp, nil
	}
	if v == "https" || v == "Https" {
		return e.Https, nil
	}
	if v == "permitted_ip" || v == "PermittedIp" {
		return e.PermittedIp, nil
	}
	if v == "permitted_ip|LENGTH" || v == "PermittedIp|LENGTH" {
		return int64(len(e.PermittedIp)), nil
	}
	if v == "ping" || v == "Ping" {
		return e.Ping, nil
	}
	if v == "response_pages" || v == "ResponsePages" {
		return e.ResponsePages, nil
	}
	if v == "snmp" || v == "Snmp" {
		return e.Snmp, nil
	}
	if v == "ssh" || v == "Ssh" {
		return e.Ssh, nil
	}
	if v == "telnet" || v == "Telnet" {
		return e.Telnet, nil
	}
	if v == "userid_service" || v == "UseridService" {
		return e.UseridService, nil
	}
	if v == "userid_syslog_listener_ssl" || v == "UseridSyslogListenerSsl" {
		return e.UseridSyslogListenerSsl, nil
	}
	if v == "userid_syslog_listener_udp" || v == "UseridSyslogListenerUdp" {
		return e.UseridSyslogListenerUdp, nil
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
	if !util.BoolsMatch(o.Http, other.Http) {
		return false
	}
	if !util.BoolsMatch(o.HttpOcsp, other.HttpOcsp) {
		return false
	}
	if !util.BoolsMatch(o.Https, other.Https) {
		return false
	}
	if len(o.PermittedIp) != len(other.PermittedIp) {
		return false
	}
	for idx := range o.PermittedIp {
		if !o.PermittedIp[idx].matches(&other.PermittedIp[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Ping, other.Ping) {
		return false
	}
	if !util.BoolsMatch(o.ResponsePages, other.ResponsePages) {
		return false
	}
	if !util.BoolsMatch(o.Snmp, other.Snmp) {
		return false
	}
	if !util.BoolsMatch(o.Ssh, other.Ssh) {
		return false
	}
	if !util.BoolsMatch(o.Telnet, other.Telnet) {
		return false
	}
	if !util.BoolsMatch(o.UseridService, other.UseridService) {
		return false
	}
	if !util.BoolsMatch(o.UseridSyslogListenerSsl, other.UseridSyslogListenerSsl) {
		return false
	}
	if !util.BoolsMatch(o.UseridSyslogListenerUdp, other.UseridSyslogListenerUdp) {
		return false
	}

	return true
}

func (o *PermittedIp) matches(other *PermittedIp) bool {
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

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
