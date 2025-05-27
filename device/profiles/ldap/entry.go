package ldap

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
	Suffix = []string{"server-profile", "ldap"}
)

type Entry struct {
	Name                    string
	Base                    *string
	BindDn                  *string
	BindPassword            *string
	BindTimelimit           *int64
	Disabled                *bool
	LdapType                *string
	RetryInterval           *int64
	Server                  []Server
	Ssl                     *bool
	Timelimit               *int64
	VerifyServerCertificate *bool

	Misc map[string][]generic.Xml
}

type Server struct {
	Name    string
	Address *string
	Port    *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                 xml.Name    `xml:"entry"`
	Name                    string      `xml:"name,attr"`
	Base                    *string     `xml:"base,omitempty"`
	BindDn                  *string     `xml:"bind-dn,omitempty"`
	BindPassword            *string     `xml:"bind-password,omitempty"`
	BindTimelimit           *int64      `xml:"bind-timelimit,omitempty"`
	Disabled                *string     `xml:"disabled,omitempty"`
	LdapType                *string     `xml:"ldap-type,omitempty"`
	RetryInterval           *int64      `xml:"retry-interval,omitempty"`
	Server                  []ServerXml `xml:"server>entry,omitempty"`
	Ssl                     *string     `xml:"ssl,omitempty"`
	Timelimit               *int64      `xml:"timelimit,omitempty"`
	VerifyServerCertificate *string     `xml:"verify-server-certificate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ServerXml struct {
	Address *string `xml:"address,omitempty"`
	Name    string  `xml:"name,attr"`
	Port    *int64  `xml:"port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "base" || v == "Base" {
		return e.Base, nil
	}
	if v == "bind_dn" || v == "BindDn" {
		return e.BindDn, nil
	}
	if v == "bind_password" || v == "BindPassword" {
		return e.BindPassword, nil
	}
	if v == "bind_timelimit" || v == "BindTimelimit" {
		return e.BindTimelimit, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "ldap_type" || v == "LdapType" {
		return e.LdapType, nil
	}
	if v == "retry_interval" || v == "RetryInterval" {
		return e.RetryInterval, nil
	}
	if v == "server" || v == "Server" {
		return e.Server, nil
	}
	if v == "server|LENGTH" || v == "Server|LENGTH" {
		return int64(len(e.Server)), nil
	}
	if v == "ssl" || v == "Ssl" {
		return e.Ssl, nil
	}
	if v == "timelimit" || v == "Timelimit" {
		return e.Timelimit, nil
	}
	if v == "verify_server_certificate" || v == "VerifyServerCertificate" {
		return e.VerifyServerCertificate, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Base = o.Base
	entry.BindDn = o.BindDn
	entry.BindPassword = o.BindPassword
	entry.BindTimelimit = o.BindTimelimit
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.LdapType = o.LdapType
	entry.RetryInterval = o.RetryInterval
	var nestedServerCol []ServerXml
	if o.Server != nil {
		nestedServerCol = []ServerXml{}
		for _, oServer := range o.Server {
			nestedServer := ServerXml{}
			if _, ok := o.Misc["Server"]; ok {
				nestedServer.Misc = o.Misc["Server"]
			}
			if oServer.Name != "" {
				nestedServer.Name = oServer.Name
			}
			if oServer.Address != nil {
				nestedServer.Address = oServer.Address
			}
			if oServer.Port != nil {
				nestedServer.Port = oServer.Port
			}
			nestedServerCol = append(nestedServerCol, nestedServer)
		}
		entry.Server = nestedServerCol
	}

	entry.Ssl = util.YesNo(o.Ssl, nil)
	entry.Timelimit = o.Timelimit
	entry.VerifyServerCertificate = util.YesNo(o.VerifyServerCertificate, nil)

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
		entry.Base = o.Base
		entry.BindDn = o.BindDn
		entry.BindPassword = o.BindPassword
		entry.BindTimelimit = o.BindTimelimit
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.LdapType = o.LdapType
		entry.RetryInterval = o.RetryInterval
		var nestedServerCol []Server
		if o.Server != nil {
			nestedServerCol = []Server{}
			for _, oServer := range o.Server {
				nestedServer := Server{}
				if oServer.Misc != nil {
					entry.Misc["Server"] = oServer.Misc
				}
				if oServer.Name != "" {
					nestedServer.Name = oServer.Name
				}
				if oServer.Address != nil {
					nestedServer.Address = oServer.Address
				}
				if oServer.Port != nil {
					nestedServer.Port = oServer.Port
				}
				nestedServerCol = append(nestedServerCol, nestedServer)
			}
			entry.Server = nestedServerCol
		}

		entry.Ssl = util.AsBool(o.Ssl, nil)
		entry.Timelimit = o.Timelimit
		entry.VerifyServerCertificate = util.AsBool(o.VerifyServerCertificate, nil)

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
	if !util.StringsMatch(a.Base, b.Base) {
		return false
	}
	if !util.StringsMatch(a.BindDn, b.BindDn) {
		return false
	}
	if !util.StringsMatch(a.BindPassword, b.BindPassword) {
		return false
	}
	if !util.Ints64Match(a.BindTimelimit, b.BindTimelimit) {
		return false
	}
	if !util.BoolsMatch(a.Disabled, b.Disabled) {
		return false
	}
	if !util.StringsMatch(a.LdapType, b.LdapType) {
		return false
	}
	if !util.Ints64Match(a.RetryInterval, b.RetryInterval) {
		return false
	}
	if !matchServer(a.Server, b.Server) {
		return false
	}
	if !util.BoolsMatch(a.Ssl, b.Ssl) {
		return false
	}
	if !util.Ints64Match(a.Timelimit, b.Timelimit) {
		return false
	}
	if !util.BoolsMatch(a.VerifyServerCertificate, b.VerifyServerCertificate) {
		return false
	}

	return true
}

func matchServer(a []Server, b []Server) bool {
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
			if !util.StringsMatch(a.Address, b.Address) {
				return false
			}
			if !util.Ints64Match(a.Port, b.Port) {
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
