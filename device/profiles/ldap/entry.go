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
	suffix = []string{"server-profile", "ldap", "$name"}
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
	Misc                    []generic.Xml
}
type Server struct {
	Name    string
	Address *string
	Port    *int64
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
	XMLName                 xml.Name            `xml:"entry"`
	Name                    string              `xml:"name,attr"`
	Base                    *string             `xml:"base,omitempty"`
	BindDn                  *string             `xml:"bind-dn,omitempty"`
	BindPassword            *string             `xml:"bind-password,omitempty"`
	BindTimelimit           *int64              `xml:"bind-timelimit,omitempty"`
	Disabled                *string             `xml:"disabled,omitempty"`
	LdapType                *string             `xml:"ldap-type,omitempty"`
	RetryInterval           *int64              `xml:"retry-interval,omitempty"`
	Server                  *serverContainerXml `xml:"server,omitempty"`
	Ssl                     *string             `xml:"ssl,omitempty"`
	Timelimit               *int64              `xml:"timelimit,omitempty"`
	VerifyServerCertificate *string             `xml:"verify-server-certificate,omitempty"`
	Misc                    []generic.Xml       `xml:",any"`
}
type serverContainerXml struct {
	Entries []serverXml `xml:"entry"`
}
type serverXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Address *string       `xml:"address,omitempty"`
	Port    *int64        `xml:"port,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Base = s.Base
	o.BindDn = s.BindDn
	o.BindPassword = s.BindPassword
	o.BindTimelimit = s.BindTimelimit
	o.Disabled = util.YesNo(s.Disabled, nil)
	o.LdapType = s.LdapType
	o.RetryInterval = s.RetryInterval
	if s.Server != nil {
		var objs []serverXml
		for _, elt := range s.Server {
			var obj serverXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &serverContainerXml{Entries: objs}
	}
	o.Ssl = util.YesNo(s.Ssl, nil)
	o.Timelimit = s.Timelimit
	o.VerifyServerCertificate = util.YesNo(s.VerifyServerCertificate, nil)
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var serverVal []Server
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Entry{
		Name:                    o.Name,
		Base:                    o.Base,
		BindDn:                  o.BindDn,
		BindPassword:            o.BindPassword,
		BindTimelimit:           o.BindTimelimit,
		Disabled:                util.AsBool(o.Disabled, nil),
		LdapType:                o.LdapType,
		RetryInterval:           o.RetryInterval,
		Server:                  serverVal,
		Ssl:                     util.AsBool(o.Ssl, nil),
		Timelimit:               o.Timelimit,
		VerifyServerCertificate: util.AsBool(o.VerifyServerCertificate, nil),
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *serverXml) MarshalFromObject(s Server) {
	o.Name = s.Name
	o.Address = s.Address
	o.Port = s.Port
	o.Misc = s.Misc
}

func (o serverXml) UnmarshalToObject() (*Server, error) {

	result := &Server{
		Name:    o.Name,
		Address: o.Address,
		Port:    o.Port,
		Misc:    o.Misc,
	}
	return result, nil
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
	if !util.StringsMatch(o.Base, other.Base) {
		return false
	}
	if !util.StringsMatch(o.BindDn, other.BindDn) {
		return false
	}
	if !util.StringsMatch(o.BindPassword, other.BindPassword) {
		return false
	}
	if !util.Ints64Match(o.BindTimelimit, other.BindTimelimit) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !util.StringsMatch(o.LdapType, other.LdapType) {
		return false
	}
	if !util.Ints64Match(o.RetryInterval, other.RetryInterval) {
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
	if !util.BoolsMatch(o.Ssl, other.Ssl) {
		return false
	}
	if !util.Ints64Match(o.Timelimit, other.Timelimit) {
		return false
	}
	if !util.BoolsMatch(o.VerifyServerCertificate, other.VerifyServerCertificate) {
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
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Address, other.Address) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
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
