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
	Suffix = []string{"network", "profiles", "interface-management-profile"}
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

	Misc map[string][]generic.Xml
}

type PermittedIp struct {
	Name string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                 xml.Name         `xml:"entry"`
	Name                    string           `xml:"name,attr"`
	Http                    *string          `xml:"http,omitempty"`
	HttpOcsp                *string          `xml:"http-ocsp,omitempty"`
	Https                   *string          `xml:"https,omitempty"`
	PermittedIp             []PermittedIpXml `xml:"permitted-ip>entry,omitempty"`
	Ping                    *string          `xml:"ping,omitempty"`
	ResponsePages           *string          `xml:"response-pages,omitempty"`
	Snmp                    *string          `xml:"snmp,omitempty"`
	Ssh                     *string          `xml:"ssh,omitempty"`
	Telnet                  *string          `xml:"telnet,omitempty"`
	UseridService           *string          `xml:"userid-service,omitempty"`
	UseridSyslogListenerSsl *string          `xml:"userid-syslog-listener-ssl,omitempty"`
	UseridSyslogListenerUdp *string          `xml:"userid-syslog-listener-udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type PermittedIpXml struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Http = util.YesNo(o.Http, nil)
	entry.HttpOcsp = util.YesNo(o.HttpOcsp, nil)
	entry.Https = util.YesNo(o.Https, nil)
	var nestedPermittedIpCol []PermittedIpXml
	if o.PermittedIp != nil {
		nestedPermittedIpCol = []PermittedIpXml{}
		for _, oPermittedIp := range o.PermittedIp {
			nestedPermittedIp := PermittedIpXml{}
			if _, ok := o.Misc["PermittedIp"]; ok {
				nestedPermittedIp.Misc = o.Misc["PermittedIp"]
			}
			if oPermittedIp.Name != "" {
				nestedPermittedIp.Name = oPermittedIp.Name
			}
			nestedPermittedIpCol = append(nestedPermittedIpCol, nestedPermittedIp)
		}
		entry.PermittedIp = nestedPermittedIpCol
	}

	entry.Ping = util.YesNo(o.Ping, nil)
	entry.ResponsePages = util.YesNo(o.ResponsePages, nil)
	entry.Snmp = util.YesNo(o.Snmp, nil)
	entry.Ssh = util.YesNo(o.Ssh, nil)
	entry.Telnet = util.YesNo(o.Telnet, nil)
	entry.UseridService = util.YesNo(o.UseridService, nil)
	entry.UseridSyslogListenerSsl = util.YesNo(o.UseridSyslogListenerSsl, nil)
	entry.UseridSyslogListenerUdp = util.YesNo(o.UseridSyslogListenerUdp, nil)

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
		entry.Http = util.AsBool(o.Http, nil)
		entry.HttpOcsp = util.AsBool(o.HttpOcsp, nil)
		entry.Https = util.AsBool(o.Https, nil)
		var nestedPermittedIpCol []PermittedIp
		if o.PermittedIp != nil {
			nestedPermittedIpCol = []PermittedIp{}
			for _, oPermittedIp := range o.PermittedIp {
				nestedPermittedIp := PermittedIp{}
				if oPermittedIp.Misc != nil {
					entry.Misc["PermittedIp"] = oPermittedIp.Misc
				}
				if oPermittedIp.Name != "" {
					nestedPermittedIp.Name = oPermittedIp.Name
				}
				nestedPermittedIpCol = append(nestedPermittedIpCol, nestedPermittedIp)
			}
			entry.PermittedIp = nestedPermittedIpCol
		}

		entry.Ping = util.AsBool(o.Ping, nil)
		entry.ResponsePages = util.AsBool(o.ResponsePages, nil)
		entry.Snmp = util.AsBool(o.Snmp, nil)
		entry.Ssh = util.AsBool(o.Ssh, nil)
		entry.Telnet = util.AsBool(o.Telnet, nil)
		entry.UseridService = util.AsBool(o.UseridService, nil)
		entry.UseridSyslogListenerSsl = util.AsBool(o.UseridSyslogListenerSsl, nil)
		entry.UseridSyslogListenerUdp = util.AsBool(o.UseridSyslogListenerUdp, nil)

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
	if !util.BoolsMatch(a.Http, b.Http) {
		return false
	}
	if !util.BoolsMatch(a.HttpOcsp, b.HttpOcsp) {
		return false
	}
	if !util.BoolsMatch(a.Https, b.Https) {
		return false
	}
	if !matchPermittedIp(a.PermittedIp, b.PermittedIp) {
		return false
	}
	if !util.BoolsMatch(a.Ping, b.Ping) {
		return false
	}
	if !util.BoolsMatch(a.ResponsePages, b.ResponsePages) {
		return false
	}
	if !util.BoolsMatch(a.Snmp, b.Snmp) {
		return false
	}
	if !util.BoolsMatch(a.Ssh, b.Ssh) {
		return false
	}
	if !util.BoolsMatch(a.Telnet, b.Telnet) {
		return false
	}
	if !util.BoolsMatch(a.UseridService, b.UseridService) {
		return false
	}
	if !util.BoolsMatch(a.UseridSyslogListenerSsl, b.UseridSyslogListenerSsl) {
		return false
	}
	if !util.BoolsMatch(a.UseridSyslogListenerUdp, b.UseridSyslogListenerUdp) {
		return false
	}

	return true
}

func matchPermittedIp(a []PermittedIp, b []PermittedIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
