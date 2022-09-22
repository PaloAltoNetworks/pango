package kerberos

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// Kerberos profile.
//
// PAN-OS 7.0+
type Entry struct {
	Name         string
	AdminUseOnly bool
	Servers      []Server
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
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	AdminUseOnly string   `xml:"admin-use-only"`
	Servers      *servers `xml:"server"`
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:         o.Name,
		AdminUseOnly: util.AsBool(o.AdminUseOnly),
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
	Server  string   `xml:"host"`
	Port    int      `xml:"port,omitempty"`
}

func (e *serverEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local serverEntry
	ans := local{
		Port: 88,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = serverEntry(ans)
	return nil
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:         e.Name,
		AdminUseOnly: util.YesNo(e.AdminUseOnly),
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
