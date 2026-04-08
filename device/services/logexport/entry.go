package logexport

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
	suffix = []string{"deviceconfig", "system", "log-export-schedule", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Enable         *bool
	LogType        *string
	Protocol       *Protocol
	StartTime      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Protocol struct {
	Ftp            *ProtocolFtp
	Scp            *ProtocolScp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ProtocolFtp struct {
	Hostname       *string
	PassiveMode    *bool
	Password       *string
	Path           *string
	Port           *int64
	Username       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ProtocolScp struct {
	Hostname       *string
	Password       *string
	Path           *string
	Port           *int64
	Username       *string
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	Enable         *string       `xml:"enable,omitempty"`
	LogType        *string       `xml:"log-type,omitempty"`
	Protocol       *protocolXml  `xml:"protocol,omitempty"`
	StartTime      *string       `xml:"start-time,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type protocolXml struct {
	Ftp            *protocolFtpXml `xml:"ftp,omitempty"`
	Scp            *protocolScpXml `xml:"scp,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type protocolFtpXml struct {
	Hostname       *string       `xml:"hostname,omitempty"`
	PassiveMode    *string       `xml:"passive-mode,omitempty"`
	Password       *string       `xml:"password,omitempty"`
	Path           *string       `xml:"path,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	Username       *string       `xml:"username,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type protocolScpXml struct {
	Hostname       *string       `xml:"hostname,omitempty"`
	Password       *string       `xml:"password,omitempty"`
	Path           *string       `xml:"path,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	Username       *string       `xml:"username,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.Enable = util.YesNo(s.Enable, nil)
	o.LogType = s.LogType
	if s.Protocol != nil {
		var obj protocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.StartTime = s.StartTime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var protocolVal *Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		Enable:         util.AsBool(o.Enable, nil),
		LogType:        o.LogType,
		Protocol:       protocolVal,
		StartTime:      o.StartTime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolXml) MarshalFromObject(s Protocol) {
	if s.Ftp != nil {
		var obj protocolFtpXml
		obj.MarshalFromObject(*s.Ftp)
		o.Ftp = &obj
	}
	if s.Scp != nil {
		var obj protocolScpXml
		obj.MarshalFromObject(*s.Scp)
		o.Scp = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolXml) UnmarshalToObject() (*Protocol, error) {
	var ftpVal *ProtocolFtp
	if o.Ftp != nil {
		obj, err := o.Ftp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ftpVal = obj
	}
	var scpVal *ProtocolScp
	if o.Scp != nil {
		obj, err := o.Scp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		scpVal = obj
	}

	result := &Protocol{
		Ftp:            ftpVal,
		Scp:            scpVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolFtpXml) MarshalFromObject(s ProtocolFtp) {
	o.Hostname = s.Hostname
	o.PassiveMode = util.YesNo(s.PassiveMode, nil)
	o.Password = s.Password
	o.Path = s.Path
	o.Port = s.Port
	o.Username = s.Username
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolFtpXml) UnmarshalToObject() (*ProtocolFtp, error) {

	result := &ProtocolFtp{
		Hostname:       o.Hostname,
		PassiveMode:    util.AsBool(o.PassiveMode, nil),
		Password:       o.Password,
		Path:           o.Path,
		Port:           o.Port,
		Username:       o.Username,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolScpXml) MarshalFromObject(s ProtocolScp) {
	o.Hostname = s.Hostname
	o.Password = s.Password
	o.Path = s.Path
	o.Port = s.Port
	o.Username = s.Username
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolScpXml) UnmarshalToObject() (*ProtocolScp, error) {

	result := &ProtocolScp{
		Hostname:       o.Hostname,
		Password:       o.Password,
		Path:           o.Path,
		Port:           o.Port,
		Username:       o.Username,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "enable" || v == "Enable" {
		return e.Enable, nil
	}
	if v == "log_type" || v == "LogType" {
		return e.LogType, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "start_time" || v == "StartTime" {
		return e.StartTime, nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.LogType, other.LogType) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}
	if !util.StringsMatch(o.StartTime, other.StartTime) {
		return false
	}

	return true
}

func (o *Protocol) matches(other *Protocol) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ftp.matches(other.Ftp) {
		return false
	}
	if !o.Scp.matches(other.Scp) {
		return false
	}

	return true
}

func (o *ProtocolFtp) matches(other *ProtocolFtp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}
	if !util.BoolsMatch(o.PassiveMode, other.PassiveMode) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Path, other.Path) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *ProtocolScp) matches(other *ProtocolScp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Path, other.Path) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
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
