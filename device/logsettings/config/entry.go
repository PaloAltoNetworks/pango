package config

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
	suffix = []string{"log-settings", "config", "match-list", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	Filter         *string
	SendEmail      []string
	SendHttp       []string
	SendSnmptrap   []string
	SendSyslog     []string
	SendToPanorama *bool
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
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Description    *string          `xml:"description,omitempty"`
	Filter         *string          `xml:"filter,omitempty"`
	SendEmail      *util.MemberType `xml:"send-email,omitempty"`
	SendHttp       *util.MemberType `xml:"send-http,omitempty"`
	SendSnmptrap   *util.MemberType `xml:"send-snmptrap,omitempty"`
	SendSyslog     *util.MemberType `xml:"send-syslog,omitempty"`
	SendToPanorama *string          `xml:"send-to-panorama,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.Filter = s.Filter
	if s.SendEmail != nil {
		o.SendEmail = util.StrToMem(s.SendEmail)
	}
	if s.SendHttp != nil {
		o.SendHttp = util.StrToMem(s.SendHttp)
	}
	if s.SendSnmptrap != nil {
		o.SendSnmptrap = util.StrToMem(s.SendSnmptrap)
	}
	if s.SendSyslog != nil {
		o.SendSyslog = util.StrToMem(s.SendSyslog)
	}
	o.SendToPanorama = util.YesNo(s.SendToPanorama, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var sendEmailVal []string
	if o.SendEmail != nil {
		sendEmailVal = util.MemToStr(o.SendEmail)
	}
	var sendHttpVal []string
	if o.SendHttp != nil {
		sendHttpVal = util.MemToStr(o.SendHttp)
	}
	var sendSnmptrapVal []string
	if o.SendSnmptrap != nil {
		sendSnmptrapVal = util.MemToStr(o.SendSnmptrap)
	}
	var sendSyslogVal []string
	if o.SendSyslog != nil {
		sendSyslogVal = util.MemToStr(o.SendSyslog)
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		Filter:         o.Filter,
		SendEmail:      sendEmailVal,
		SendHttp:       sendHttpVal,
		SendSnmptrap:   sendSnmptrapVal,
		SendSyslog:     sendSyslogVal,
		SendToPanorama: util.AsBool(o.SendToPanorama, nil),
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
	if v == "filter" || v == "Filter" {
		return e.Filter, nil
	}
	if v == "send_email" || v == "SendEmail" {
		return e.SendEmail, nil
	}
	if v == "send_email|LENGTH" || v == "SendEmail|LENGTH" {
		return int64(len(e.SendEmail)), nil
	}
	if v == "send_http" || v == "SendHttp" {
		return e.SendHttp, nil
	}
	if v == "send_http|LENGTH" || v == "SendHttp|LENGTH" {
		return int64(len(e.SendHttp)), nil
	}
	if v == "send_snmptrap" || v == "SendSnmptrap" {
		return e.SendSnmptrap, nil
	}
	if v == "send_snmptrap|LENGTH" || v == "SendSnmptrap|LENGTH" {
		return int64(len(e.SendSnmptrap)), nil
	}
	if v == "send_syslog" || v == "SendSyslog" {
		return e.SendSyslog, nil
	}
	if v == "send_syslog|LENGTH" || v == "SendSyslog|LENGTH" {
		return int64(len(e.SendSyslog)), nil
	}
	if v == "send_to_panorama" || v == "SendToPanorama" {
		return e.SendToPanorama, nil
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
	if !util.StringsMatch(o.Filter, other.Filter) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendEmail, other.SendEmail) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendHttp, other.SendHttp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSnmptrap, other.SendSnmptrap) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSyslog, other.SendSyslog) {
		return false
	}
	if !util.BoolsMatch(o.SendToPanorama, other.SendToPanorama) {
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
