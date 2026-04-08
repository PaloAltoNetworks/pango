package tacacsplus

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
	suffix = []string{"server-profile", "tacplus", "$name"}
)

type Entry struct {
	Name                string
	Protocol            *string
	Server              []Server
	Timeout             *int64
	UseSingleConnection *bool
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type Server struct {
	Name           string
	Address        *string
	Secret         *string
	Port           *int64
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
	XMLName             xml.Name            `xml:"entry"`
	Name                string              `xml:"name,attr"`
	Protocol            *string             `xml:"protocol,omitempty"`
	Server              *serverContainerXml `xml:"server,omitempty"`
	Timeout             *int64              `xml:"timeout,omitempty"`
	UseSingleConnection *string             `xml:"use-single-connection,omitempty"`
	Misc                []generic.Xml       `xml:",any"`
	MiscAttributes      []xml.Attr          `xml:",any,attr"`
}
type serverContainerXml struct {
	Entries []serverXml `xml:"entry"`
}
type serverXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Address        *string       `xml:"address,omitempty"`
	Secret         *string       `xml:"secret,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Protocol = s.Protocol
	if s.Server != nil {
		var objs []serverXml
		for _, elt := range s.Server {
			var obj serverXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &serverContainerXml{Entries: objs}
	}
	o.Timeout = s.Timeout
	o.UseSingleConnection = util.YesNo(s.UseSingleConnection, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
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
		Name:                o.Name,
		Protocol:            o.Protocol,
		Server:              serverVal,
		Timeout:             o.Timeout,
		UseSingleConnection: util.AsBool(o.UseSingleConnection, nil),
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *serverXml) MarshalFromObject(s Server) {
	o.Name = s.Name
	o.Address = s.Address
	o.Secret = s.Secret
	o.Port = s.Port
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverXml) UnmarshalToObject() (*Server, error) {

	result := &Server{
		Name:           o.Name,
		Address:        o.Address,
		Secret:         o.Secret,
		Port:           o.Port,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "server" || v == "Server" {
		return e.Server, nil
	}
	if v == "server|LENGTH" || v == "Server|LENGTH" {
		return int64(len(e.Server)), nil
	}
	if v == "timeout" || v == "Timeout" {
		return e.Timeout, nil
	}
	if v == "use_single_connection" || v == "UseSingleConnection" {
		return e.UseSingleConnection, nil
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
	if !util.StringsMatch(o.Protocol, other.Protocol) {
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
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !util.BoolsMatch(o.UseSingleConnection, other.UseSingleConnection) {
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
	if !util.StringsMatch(o.Secret, other.Secret) {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
