package syslog

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
	suffix = []string{"log-settings", "syslog", "$name"}
)

type Entry struct {
	Name           string
	Format         *Format
	Server         []Server
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Format struct {
	Auth           *string
	Config         *string
	Correlation    *string
	Data           *string
	Decryption     *string
	Escaping       *FormatEscaping
	Globalprotect  *string
	Gtp            *string
	HipMatch       *string
	Iptag          *string
	Sctp           *string
	System         *string
	Threat         *string
	Traffic        *string
	Tunnel         *string
	Url            *string
	Userid         *string
	Wildfire       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FormatEscaping struct {
	EscapeCharacter   *string
	EscapedCharacters *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Server struct {
	Name           string
	Server         *string
	Transport      *string
	Port           *int64
	Format         *string
	Facility       *string
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
	XMLName        xml.Name            `xml:"entry"`
	Name           string              `xml:"name,attr"`
	Format         *formatXml          `xml:"format,omitempty"`
	Server         *serverContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml       `xml:",any"`
	MiscAttributes []xml.Attr          `xml:",any,attr"`
}
type formatXml struct {
	Auth           *string            `xml:"auth,omitempty"`
	Config         *string            `xml:"config,omitempty"`
	Correlation    *string            `xml:"correlation,omitempty"`
	Data           *string            `xml:"data,omitempty"`
	Decryption     *string            `xml:"decryption,omitempty"`
	Escaping       *formatEscapingXml `xml:"escaping,omitempty"`
	Globalprotect  *string            `xml:"globalprotect,omitempty"`
	Gtp            *string            `xml:"gtp,omitempty"`
	HipMatch       *string            `xml:"hip-match,omitempty"`
	Iptag          *string            `xml:"iptag,omitempty"`
	Sctp           *string            `xml:"sctp,omitempty"`
	System         *string            `xml:"system,omitempty"`
	Threat         *string            `xml:"threat,omitempty"`
	Traffic        *string            `xml:"traffic,omitempty"`
	Tunnel         *string            `xml:"tunnel,omitempty"`
	Url            *string            `xml:"url,omitempty"`
	Userid         *string            `xml:"userid,omitempty"`
	Wildfire       *string            `xml:"wildfire,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
	MiscAttributes []xml.Attr         `xml:",any,attr"`
}
type formatEscapingXml struct {
	EscapeCharacter   *string       `xml:"escape-character,omitempty"`
	EscapedCharacters *string       `xml:"escaped-characters,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type serverContainerXml struct {
	Entries []serverXml `xml:"entry"`
}
type serverXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Server         *string       `xml:"server,omitempty"`
	Transport      *string       `xml:"transport,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	Format         *string       `xml:"format,omitempty"`
	Facility       *string       `xml:"facility,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Format != nil {
		var obj formatXml
		obj.MarshalFromObject(*s.Format)
		o.Format = &obj
	}
	if s.Server != nil {
		var objs []serverXml
		for _, elt := range s.Server {
			var obj serverXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &serverContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var formatVal *Format
	if o.Format != nil {
		obj, err := o.Format.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		formatVal = obj
	}
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
		Name:           o.Name,
		Format:         formatVal,
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *formatXml) MarshalFromObject(s Format) {
	o.Auth = s.Auth
	o.Config = s.Config
	o.Correlation = s.Correlation
	o.Data = s.Data
	o.Decryption = s.Decryption
	if s.Escaping != nil {
		var obj formatEscapingXml
		obj.MarshalFromObject(*s.Escaping)
		o.Escaping = &obj
	}
	o.Globalprotect = s.Globalprotect
	o.Gtp = s.Gtp
	o.HipMatch = s.HipMatch
	o.Iptag = s.Iptag
	o.Sctp = s.Sctp
	o.System = s.System
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.Tunnel = s.Tunnel
	o.Url = s.Url
	o.Userid = s.Userid
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o formatXml) UnmarshalToObject() (*Format, error) {
	var escapingVal *FormatEscaping
	if o.Escaping != nil {
		obj, err := o.Escaping.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		escapingVal = obj
	}

	result := &Format{
		Auth:           o.Auth,
		Config:         o.Config,
		Correlation:    o.Correlation,
		Data:           o.Data,
		Decryption:     o.Decryption,
		Escaping:       escapingVal,
		Globalprotect:  o.Globalprotect,
		Gtp:            o.Gtp,
		HipMatch:       o.HipMatch,
		Iptag:          o.Iptag,
		Sctp:           o.Sctp,
		System:         o.System,
		Threat:         o.Threat,
		Traffic:        o.Traffic,
		Tunnel:         o.Tunnel,
		Url:            o.Url,
		Userid:         o.Userid,
		Wildfire:       o.Wildfire,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *formatEscapingXml) MarshalFromObject(s FormatEscaping) {
	o.EscapeCharacter = s.EscapeCharacter
	o.EscapedCharacters = s.EscapedCharacters
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o formatEscapingXml) UnmarshalToObject() (*FormatEscaping, error) {

	result := &FormatEscaping{
		EscapeCharacter:   o.EscapeCharacter,
		EscapedCharacters: o.EscapedCharacters,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *serverXml) MarshalFromObject(s Server) {
	o.Name = s.Name
	o.Server = s.Server
	o.Transport = s.Transport
	o.Port = s.Port
	o.Format = s.Format
	o.Facility = s.Facility
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverXml) UnmarshalToObject() (*Server, error) {

	result := &Server{
		Name:           o.Name,
		Server:         o.Server,
		Transport:      o.Transport,
		Port:           o.Port,
		Format:         o.Format,
		Facility:       o.Facility,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "format" || v == "Format" {
		return e.Format, nil
	}
	if v == "server" || v == "Server" {
		return e.Server, nil
	}
	if v == "server|LENGTH" || v == "Server|LENGTH" {
		return int64(len(e.Server)), nil
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
	if !o.Format.matches(other.Format) {
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

func (o *Format) matches(other *Format) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Auth, other.Auth) {
		return false
	}
	if !util.StringsMatch(o.Config, other.Config) {
		return false
	}
	if !util.StringsMatch(o.Correlation, other.Correlation) {
		return false
	}
	if !util.StringsMatch(o.Data, other.Data) {
		return false
	}
	if !util.StringsMatch(o.Decryption, other.Decryption) {
		return false
	}
	if !o.Escaping.matches(other.Escaping) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.Gtp, other.Gtp) {
		return false
	}
	if !util.StringsMatch(o.HipMatch, other.HipMatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.Sctp, other.Sctp) {
		return false
	}
	if !util.StringsMatch(o.System, other.System) {
		return false
	}
	if !util.StringsMatch(o.Threat, other.Threat) {
		return false
	}
	if !util.StringsMatch(o.Traffic, other.Traffic) {
		return false
	}
	if !util.StringsMatch(o.Tunnel, other.Tunnel) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}
	if !util.StringsMatch(o.Userid, other.Userid) {
		return false
	}
	if !util.StringsMatch(o.Wildfire, other.Wildfire) {
		return false
	}

	return true
}

func (o *FormatEscaping) matches(other *FormatEscaping) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EscapeCharacter, other.EscapeCharacter) {
		return false
	}
	if !util.StringsMatch(o.EscapedCharacters, other.EscapedCharacters) {
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
	if !util.StringsMatch(o.Server, other.Server) {
		return false
	}
	if !util.StringsMatch(o.Transport, other.Transport) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.Format, other.Format) {
		return false
	}
	if !util.StringsMatch(o.Facility, other.Facility) {
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
