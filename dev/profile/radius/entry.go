package radius

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
	Name         string
	AdminUseOnly bool
	Timeout      int
	Retries      int
	Servers      []Server
	Protocol     Protocol // 8.0
}

type Server struct {
	Name   string
	Server string
	Secret string // encrypted
	Port   int
}

type Protocol struct {
	Chap           bool
	Pap            bool
	Auto           bool            // 8.0, removed in 8.1
	PeapMschapv2   *PeapMschapv2   // 8.1
	PeapWithGtc    *PeapWithGtc    // 8.1
	EapTtlsWithPap *EapTtlsWithPap // 8.1
}

type PeapMschapv2 struct {
	MakeOuterIdentityAnonymous bool
	AllowExpiredPasswordChange bool
	CertificateProfile         string
}

type PeapWithGtc struct {
	MakeOuterIdentityAnonymous bool
	CertificateProfile         string
}

type EapTtlsWithPap struct {
	MakeOuterIdentityAnonymous bool
	CertificateProfile         string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.AdminUseOnly = s.AdminUseOnly
	o.Timeout = s.Timeout
	o.Retries = s.Retries

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

	switch {
	case s.Protocol.Chap:
		o.Protocol = Protocol{Chap: true}
	case s.Protocol.Pap:
		o.Protocol = Protocol{Pap: true}
	case s.Protocol.Auto:
		o.Protocol = Protocol{Auto: true}
	case s.Protocol.PeapMschapv2 != nil:
		o.Protocol = Protocol{PeapMschapv2: &PeapMschapv2{
			MakeOuterIdentityAnonymous: s.Protocol.PeapMschapv2.MakeOuterIdentityAnonymous,
			AllowExpiredPasswordChange: s.Protocol.PeapMschapv2.AllowExpiredPasswordChange,
			CertificateProfile:         s.Protocol.PeapMschapv2.CertificateProfile,
		}}
	case s.Protocol.PeapWithGtc != nil:
		o.Protocol = Protocol{PeapWithGtc: &PeapWithGtc{
			MakeOuterIdentityAnonymous: s.Protocol.PeapWithGtc.MakeOuterIdentityAnonymous,
			CertificateProfile:         s.Protocol.PeapWithGtc.CertificateProfile,
		}}
	case s.Protocol.EapTtlsWithPap != nil:
		o.Protocol = Protocol{EapTtlsWithPap: &EapTtlsWithPap{
			MakeOuterIdentityAnonymous: s.Protocol.EapTtlsWithPap.MakeOuterIdentityAnonymous,
			CertificateProfile:         s.Protocol.EapTtlsWithPap.CertificateProfile,
		}}
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
	Timeout      int      `xml:"timeout,omitempty"`
	Retries      int      `xml:"retries,omitempty"`
	Servers      *servers `xml:"server"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{
		Timeout: 3,
		Retries: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v1(ans)
	return nil
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:         o.Name,
		AdminUseOnly: util.AsBool(o.AdminUseOnly),
		Timeout:      o.Timeout,
		Retries:      o.Retries,
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
	Server  string   `xml:"ip-address"`
	Secret  string   `xml:"secret"`
	Port    int      `xml:"port,omitempty"`
}

func (e *serverEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local serverEntry
	ans := local{
		Port: 1812,
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
		Timeout:      e.Timeout,
		Retries:      e.Retries,
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
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	AdminUseOnly string   `xml:"admin-use-only"`
	Timeout      int      `xml:"timeout,omitempty"`
	Retries      int      `xml:"retries,omitempty"`
	Servers      *servers `xml:"server"`
	Protocol     string   `xml:"protocol"`
}

func (e *entry_v2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v2
	ans := local{
		Timeout: 3,
		Retries: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v2(ans)
	return nil
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:         o.Name,
		AdminUseOnly: util.AsBool(o.AdminUseOnly),
		Timeout:      o.Timeout,
		Retries:      o.Retries,
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

	switch o.Protocol {
	case "CHAP":
		ans.Protocol = Protocol{Chap: true}
	case "PAP":
		ans.Protocol = Protocol{Pap: true}
	case "Auto":
		ans.Protocol = Protocol{Auto: true}
	}

	return ans
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:         e.Name,
		AdminUseOnly: util.YesNo(e.AdminUseOnly),
		Timeout:      e.Timeout,
		Retries:      e.Retries,
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

	switch {
	case e.Protocol.Chap:
		ans.Protocol = "CHAP"
	case e.Protocol.Pap:
		ans.Protocol = "PAP"
	case e.Protocol.Auto:
		ans.Protocol = "Auto"
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
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	AdminUseOnly string   `xml:"admin-use-only"`
	Timeout      int      `xml:"timeout,omitempty"`
	Retries      int      `xml:"retries,omitempty"`
	Servers      *servers `xml:"server"`
	Protocol     protocol `xml:"protocol"`
}

type protocol struct {
	Chap           *string         `xml:"CHAP"`
	Pap            *string         `xml:"PAP"`
	PeapMschapv2   *peapMschapv2   `xml:"PEAP-MSCHAPv2"`
	PeapWithGtc    *peapWithGtc    `xml:"PEAP-with-GTC"`
	EapTtlsWithPap *eapTtlsWithPap `xml:"EAP-TTLS-with-PAP"`
}

type peapMschapv2 struct {
	MakeOuterIdentityAnonymous string `xml:"anon-outer-id"`
	AllowExpiredPasswordChange string `xml:"allow-pwd-change"`
	CertificateProfile         string `xml:"radius-cert-profile"`
}

type peapWithGtc struct {
	MakeOuterIdentityAnonymous string `xml:"anon-outer-id"`
	CertificateProfile         string `xml:"radius-cert-profile"`
}

type eapTtlsWithPap struct {
	MakeOuterIdentityAnonymous string `xml:"anon-outer-id"`
	CertificateProfile         string `xml:"radius-cert-profile"`
}

func (e *entry_v3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v3
	ans := local{
		Timeout: 3,
		Retries: 3,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v3(ans)
	return nil
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:         o.Name,
		AdminUseOnly: util.AsBool(o.AdminUseOnly),
		Timeout:      o.Timeout,
		Retries:      o.Retries,
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

	switch {
	case o.Protocol.Chap != nil:
		ans.Protocol = Protocol{Chap: true}
	case o.Protocol.Pap != nil:
		ans.Protocol = Protocol{Pap: true}
	case o.Protocol.PeapMschapv2 != nil:
		ans.Protocol = Protocol{PeapMschapv2: &PeapMschapv2{
			MakeOuterIdentityAnonymous: util.AsBool(o.Protocol.PeapMschapv2.MakeOuterIdentityAnonymous),
			AllowExpiredPasswordChange: util.AsBool(o.Protocol.PeapMschapv2.AllowExpiredPasswordChange),
			CertificateProfile:         o.Protocol.PeapMschapv2.CertificateProfile,
		}}
	case o.Protocol.PeapWithGtc != nil:
		ans.Protocol = Protocol{PeapWithGtc: &PeapWithGtc{
			MakeOuterIdentityAnonymous: util.AsBool(o.Protocol.PeapWithGtc.MakeOuterIdentityAnonymous),
			CertificateProfile:         o.Protocol.PeapWithGtc.CertificateProfile,
		}}
	case o.Protocol.EapTtlsWithPap != nil:
		ans.Protocol = Protocol{EapTtlsWithPap: &EapTtlsWithPap{
			MakeOuterIdentityAnonymous: util.AsBool(o.Protocol.EapTtlsWithPap.MakeOuterIdentityAnonymous),
			CertificateProfile:         o.Protocol.EapTtlsWithPap.CertificateProfile,
		}}
	}

	return ans
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:         e.Name,
		AdminUseOnly: util.YesNo(e.AdminUseOnly),
		Timeout:      e.Timeout,
		Retries:      e.Retries,
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

	s := ""
	switch {
	case e.Protocol.Chap:
		ans.Protocol.Chap = &s
	case e.Protocol.Pap:
		ans.Protocol.Pap = &s
	case e.Protocol.PeapMschapv2 != nil:
		ans.Protocol.PeapMschapv2 = &peapMschapv2{
			MakeOuterIdentityAnonymous: util.YesNo(e.Protocol.PeapMschapv2.MakeOuterIdentityAnonymous),
			AllowExpiredPasswordChange: util.YesNo(e.Protocol.PeapMschapv2.AllowExpiredPasswordChange),
			CertificateProfile:         e.Protocol.PeapMschapv2.CertificateProfile,
		}
	case e.Protocol.PeapWithGtc != nil:
		ans.Protocol.PeapWithGtc = &peapWithGtc{
			MakeOuterIdentityAnonymous: util.YesNo(e.Protocol.PeapWithGtc.MakeOuterIdentityAnonymous),
			CertificateProfile:         e.Protocol.PeapWithGtc.CertificateProfile,
		}
	case e.Protocol.EapTtlsWithPap != nil:
		ans.Protocol.EapTtlsWithPap = &eapTtlsWithPap{
			MakeOuterIdentityAnonymous: util.YesNo(e.Protocol.EapTtlsWithPap.MakeOuterIdentityAnonymous),
			CertificateProfile:         e.Protocol.EapTtlsWithPap.CertificateProfile,
		}
	}

	return ans
}
