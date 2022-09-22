package tacplus

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// TACACS+ profile.
//
// PAN-OS 7.0+
type Entry struct {
	Name                string
	AdminUseOnly        bool
	Timeout             int
	UseSingleConnection bool
	Servers             []Server
	Protocol            Protocol // 8.0+
}

type Server struct {
	Name   string
	Server string
	Secret string // encrypted
	Port   int
}

type Protocol struct {
	Chap bool
	Pap  bool
	Auto bool // This option was removed in 8.1
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.AdminUseOnly = s.AdminUseOnly
	o.Timeout = s.Timeout
	o.Protocol = Protocol{
		Chap: s.Protocol.Chap,
		Pap:  s.Protocol.Pap,
		Auto: s.Protocol.Auto,
	}

	if s.Servers == nil {
		o.Servers = nil
	} else {
		o.Servers = make([]Server, 0, len(s.Servers))
		for _, x := range s.Servers {
			o.Servers = append(o.Servers, Server{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
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
	XMLName             xml.Name `xml:"entry"`
	Name                string   `xml:"name,attr"`
	AdminUseOnly        string   `xml:"admin-use-only"`
	Timeout             int      `xml:"timeout,omitempty"`
	UseSingleConnection string   `xml:"use-single-connection"`
	Servers             *servers `xml:"server"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{
		Timeout: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v1(ans)
	return nil
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:                o.Name,
		AdminUseOnly:        util.AsBool(o.AdminUseOnly),
		Timeout:             o.Timeout,
		UseSingleConnection: util.AsBool(o.UseSingleConnection),
	}

	if o.Servers != nil && len(o.Servers.Entries) > 0 {
		listing := make([]Server, 0, len(o.Servers.Entries))
		for _, x := range o.Servers.Entries {
			listing = append(listing, Server{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
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
	Secret  string   `xml:"secret"`
	Port    int      `xml:"port,omitempty"`
}

func (e *serverEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local serverEntry
	ans := local{
		Port: 49,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = serverEntry(ans)
	return nil
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:                e.Name,
		AdminUseOnly:        util.YesNo(e.AdminUseOnly),
		Timeout:             e.Timeout,
		UseSingleConnection: util.YesNo(e.UseSingleConnection),
	}

	if len(e.Servers) > 0 {
		listing := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			listing = append(listing, serverEntry{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
				Port:   x.Port,
			})
		}
		ans.Servers = &servers{Entries: listing}
	}

	return ans
}

// PAN-OS 8.0
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

type entry_v2 struct {
	XMLName             xml.Name `xml:"entry"`
	Name                string   `xml:"name,attr"`
	AdminUseOnly        string   `xml:"admin-use-only"`
	Timeout             int      `xml:"timeout,omitempty"`
	UseSingleConnection string   `xml:"use-single-connection"`
	Protocol            string   `xml:"protocol"`
	Servers             *servers `xml:"server"`
}

func (e *entry_v2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v2
	ans := local{
		Timeout: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v2(ans)
	return nil
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:                o.Name,
		AdminUseOnly:        util.AsBool(o.AdminUseOnly),
		Timeout:             o.Timeout,
		UseSingleConnection: util.AsBool(o.UseSingleConnection),
	}

	switch o.Protocol {
	case "CHAP":
		ans.Protocol.Chap = true
	case "PAP":
		ans.Protocol.Pap = true
	case "Auto":
		ans.Protocol.Auto = true
	}

	if o.Servers != nil && len(o.Servers.Entries) > 0 {
		listing := make([]Server, 0, len(o.Servers.Entries))
		for _, x := range o.Servers.Entries {
			listing = append(listing, Server{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
				Port:   x.Port,
			})
		}
		ans.Servers = listing
	}

	return ans
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:                e.Name,
		AdminUseOnly:        util.YesNo(e.AdminUseOnly),
		Timeout:             e.Timeout,
		UseSingleConnection: util.YesNo(e.UseSingleConnection),
	}

	switch {
	case e.Protocol.Chap:
		ans.Protocol = "CHAP"
	case e.Protocol.Pap:
		ans.Protocol = "PAP"
	case e.Protocol.Auto:
		ans.Protocol = "Auto"
	}

	if len(e.Servers) > 0 {
		listing := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			listing = append(listing, serverEntry{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
				Port:   x.Port,
			})
		}
		ans.Servers = &servers{Entries: listing}
	}

	return ans
}

// PAN-OS 8.1
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v3) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

type entry_v3 struct {
	XMLName             xml.Name `xml:"entry"`
	Name                string   `xml:"name,attr"`
	AdminUseOnly        string   `xml:"admin-use-only"`
	Timeout             int      `xml:"timeout,omitempty"`
	UseSingleConnection string   `xml:"use-single-connection"`
	Protocol            string   `xml:"protocol"`
	Servers             *servers `xml:"server"`
}

func (e *entry_v3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v3
	ans := local{
		Timeout: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v3(ans)
	return nil
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:                o.Name,
		AdminUseOnly:        util.AsBool(o.AdminUseOnly),
		Timeout:             o.Timeout,
		UseSingleConnection: util.AsBool(o.UseSingleConnection),
	}

	switch o.Protocol {
	case "CHAP":
		ans.Protocol.Chap = true
	case "PAP":
		ans.Protocol.Pap = true
	}

	if o.Servers != nil && len(o.Servers.Entries) > 0 {
		listing := make([]Server, 0, len(o.Servers.Entries))
		for _, x := range o.Servers.Entries {
			listing = append(listing, Server{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
				Port:   x.Port,
			})
		}
		ans.Servers = listing
	}

	return ans
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:                e.Name,
		AdminUseOnly:        util.YesNo(e.AdminUseOnly),
		Timeout:             e.Timeout,
		UseSingleConnection: util.YesNo(e.UseSingleConnection),
	}

	switch {
	case e.Protocol.Chap:
		ans.Protocol = "CHAP"
	case e.Protocol.Pap:
		ans.Protocol = "PAP"
	}

	if len(e.Servers) > 0 {
		listing := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			listing = append(listing, serverEntry{
				Name:   x.Name,
				Server: x.Server,
				Secret: x.Secret,
				Port:   x.Port,
			})
		}
		ans.Servers = &servers{Entries: listing}
	}

	return ans
}
