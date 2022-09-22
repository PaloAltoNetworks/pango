package ldap

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// radius profile.
//
// PAN-OS 7.0+
type Entry struct {
	Name                    string
	AdminUseOnly            bool
	LdapType                string
	Ssl                     bool
	VerifyServerCertificate bool
	Disabled                bool
	BaseDn                  string
	BindDn                  string
	Password                string
	BindTimeout             int
	SearchTimeout           int
	RetryInterval           int
	Servers                 []Server
}

type Server struct {
	Name   string
	Server string
	Port   int
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.AdminUseOnly = s.AdminUseOnly
	o.LdapType = s.LdapType
	o.Ssl = s.Ssl
	o.VerifyServerCertificate = s.VerifyServerCertificate
	o.Disabled = s.Disabled
	o.BaseDn = s.BaseDn
	o.BindDn = s.BindDn
	o.Password = s.Password
	o.BindTimeout = s.BindTimeout
	o.SearchTimeout = s.SearchTimeout
	o.RetryInterval = s.RetryInterval

	if s.Servers == nil {
		o.Servers = nil
	} else {
		o.Servers = make([]Server, 0, len(s.Servers))
		for _, x := range s.Servers {
			o.Servers = append(o.Servers, Server{
				Name:   x.Name,
				Server: x.Server,
				Port:   x.Port,
			})
		}
	}
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

type entry_v1 struct {
	XMLName                 xml.Name `xml:"entry"`
	Name                    string   `xml:"name,attr"`
	AdminUseOnly            string   `xml:"admin-use-only"`
	LdapType                string   `xml:"ldap-type,omitempty"`
	Ssl                     string   `xml:"ssl,omitempty"`
	VerifyServerCertificate string   `xml:"verify-server-certificate"`
	Disabled                string   `xml:"disabled"`
	BaseDn                  string   `xml:"base,omitempty"`
	BindDn                  string   `xml:"bind-dn,omitempty"`
	Password                string   `xml:"bind-password,omitempty"`
	BindTimeout             int      `xml:"bind-timelimit,omitempty"`
	SearchTimeout           int      `xml:"timelimit,omitempty"`
	RetryInterval           int      `xml:"retry-interval,omitempty"`
	Servers                 *servers `xml:"server"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{
		LdapType:      LdapTypeOther,
		Ssl:           util.YesNo(true),
		BindTimeout:   30,
		SearchTimeout: 30,
		RetryInterval: 60,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v1(ans)
	return nil
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:                    o.Name,
		AdminUseOnly:            util.AsBool(o.AdminUseOnly),
		LdapType:                o.LdapType,
		Ssl:                     util.AsBool(o.Ssl),
		VerifyServerCertificate: util.AsBool(o.VerifyServerCertificate),
		Disabled:                util.AsBool(o.Disabled),
		BaseDn:                  o.BaseDn,
		BindDn:                  o.BindDn,
		Password:                o.Password,
		BindTimeout:             o.BindTimeout,
		SearchTimeout:           o.SearchTimeout,
		RetryInterval:           o.RetryInterval,
	}

	if o.Servers != nil && len(o.Servers.Entries) > 0 {
		listing := make([]Server, 0, len(o.Servers.Entries))
		for _, x := range o.Servers.Entries {
			listing = append(listing, Server{
				Name:   x.Name,
				Server: x.Server,
				Port:   x.Port,
			})
		}
		ans.Servers = listing
	}

	return ans
}

type servers struct {
	Entries []serverEntry `xml:"entry"`
}

type serverEntry struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Server  string   `xml:"address"`
	Port    int      `xml:"port,omitempty"`
}

func (e *serverEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local serverEntry
	ans := local{
		Port: 389,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = serverEntry(ans)
	return nil
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:                    e.Name,
		AdminUseOnly:            util.YesNo(e.AdminUseOnly),
		LdapType:                e.LdapType,
		VerifyServerCertificate: util.YesNo(e.VerifyServerCertificate),
		Disabled:                util.YesNo(e.Disabled),
		BaseDn:                  e.BaseDn,
		BindDn:                  e.BindDn,
		Password:                e.Password,
		BindTimeout:             e.BindTimeout,
		SearchTimeout:           e.SearchTimeout,
		RetryInterval:           e.RetryInterval,
	}

	if !e.Ssl {
		ans.Ssl = util.YesNo(false)
	}

	if len(e.Servers) > 0 {
		listing := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			listing = append(listing, serverEntry{
				Name:   x.Name,
				Server: x.Server,
				Port:   x.Port,
			})
		}
		ans.Servers = &servers{Entries: listing}
	}

	return ans
}
