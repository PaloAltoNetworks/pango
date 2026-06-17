package snmptrap

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
	suffix = []string{"log-settings", "snmptrap", "$name"}
)

type Entry struct {
	Name           string
	Version        *Version
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Version struct {
	V2c            *VersionV2c
	V3             *VersionV3
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type VersionV2c struct {
	Server         []VersionV2cServer
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type VersionV2cServer struct {
	Name           string
	Manager        *string
	Community      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type VersionV3 struct {
	Server         []VersionV3Server
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type VersionV3Server struct {
	Name           string
	Manager        *string
	User           *string
	Engineid       *string
	Authpwd        *string
	Privpwd        *string
	Authproto      *string
	Privproto      *string
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
	Version        *versionXml   `xml:"version,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type versionXml struct {
	V2c            *versionV2cXml `xml:"v2c,omitempty"`
	V3             *versionV3Xml  `xml:"v3,omitempty"`
	Misc           []generic.Xml  `xml:",any"`
	MiscAttributes []xml.Attr     `xml:",any,attr"`
}
type versionV2cXml struct {
	Server         *versionV2cServerContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml                 `xml:",any"`
	MiscAttributes []xml.Attr                    `xml:",any,attr"`
}
type versionV2cServerContainerXml struct {
	Entries []versionV2cServerXml `xml:"entry"`
}
type versionV2cServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Manager        *string       `xml:"manager,omitempty"`
	Community      *string       `xml:"community,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type versionV3Xml struct {
	Server         *versionV3ServerContainerXml `xml:"server,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type versionV3ServerContainerXml struct {
	Entries []versionV3ServerXml `xml:"entry"`
}
type versionV3ServerXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Manager        *string       `xml:"manager,omitempty"`
	User           *string       `xml:"user,omitempty"`
	Engineid       *string       `xml:"engineid,omitempty"`
	Authpwd        *string       `xml:"authpwd,omitempty"`
	Privpwd        *string       `xml:"privpwd,omitempty"`
	Authproto      *string       `xml:"authproto,omitempty"`
	Privproto      *string       `xml:"privproto,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Version != nil {
		var obj versionXml
		obj.MarshalFromObject(*s.Version)
		o.Version = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var versionVal *Version
	if o.Version != nil {
		obj, err := o.Version.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		versionVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Version:        versionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *versionXml) MarshalFromObject(s Version) {
	if s.V2c != nil {
		var obj versionV2cXml
		obj.MarshalFromObject(*s.V2c)
		o.V2c = &obj
	}
	if s.V3 != nil {
		var obj versionV3Xml
		obj.MarshalFromObject(*s.V3)
		o.V3 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o versionXml) UnmarshalToObject() (*Version, error) {
	var v2cVal *VersionV2c
	if o.V2c != nil {
		obj, err := o.V2c.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v2cVal = obj
	}
	var v3Val *VersionV3
	if o.V3 != nil {
		obj, err := o.V3.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		v3Val = obj
	}

	result := &Version{
		V2c:            v2cVal,
		V3:             v3Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *versionV2cXml) MarshalFromObject(s VersionV2c) {
	if s.Server != nil {
		var objs []versionV2cServerXml
		for _, elt := range s.Server {
			var obj versionV2cServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &versionV2cServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o versionV2cXml) UnmarshalToObject() (*VersionV2c, error) {
	var serverVal []VersionV2cServer
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &VersionV2c{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *versionV2cServerXml) MarshalFromObject(s VersionV2cServer) {
	o.Name = s.Name
	o.Manager = s.Manager
	o.Community = s.Community
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o versionV2cServerXml) UnmarshalToObject() (*VersionV2cServer, error) {

	result := &VersionV2cServer{
		Name:           o.Name,
		Manager:        o.Manager,
		Community:      o.Community,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *versionV3Xml) MarshalFromObject(s VersionV3) {
	if s.Server != nil {
		var objs []versionV3ServerXml
		for _, elt := range s.Server {
			var obj versionV3ServerXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &versionV3ServerContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o versionV3Xml) UnmarshalToObject() (*VersionV3, error) {
	var serverVal []VersionV3Server
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &VersionV3{
		Server:         serverVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *versionV3ServerXml) MarshalFromObject(s VersionV3Server) {
	o.Name = s.Name
	o.Manager = s.Manager
	o.User = s.User
	o.Engineid = s.Engineid
	o.Authpwd = s.Authpwd
	o.Privpwd = s.Privpwd
	o.Authproto = s.Authproto
	o.Privproto = s.Privproto
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o versionV3ServerXml) UnmarshalToObject() (*VersionV3Server, error) {

	result := &VersionV3Server{
		Name:           o.Name,
		Manager:        o.Manager,
		User:           o.User,
		Engineid:       o.Engineid,
		Authpwd:        o.Authpwd,
		Privpwd:        o.Privpwd,
		Authproto:      o.Authproto,
		Privproto:      o.Privproto,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "version" || v == "Version" {
		return e.Version, nil
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
	if !o.Version.matches(other.Version) {
		return false
	}

	return true
}

func (o *Version) matches(other *Version) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.V2c.matches(other.V2c) {
		return false
	}
	if !o.V3.matches(other.V3) {
		return false
	}

	return true
}

func (o *VersionV2c) matches(other *VersionV2c) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *VersionV2cServer) matches(other *VersionV2cServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Manager, other.Manager) {
		return false
	}
	if !util.StringsMatch(o.Community, other.Community) {
		return false
	}

	return true
}

func (o *VersionV3) matches(other *VersionV3) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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

func (o *VersionV3Server) matches(other *VersionV3Server) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Manager, other.Manager) {
		return false
	}
	if !util.StringsMatch(o.User, other.User) {
		return false
	}
	if !util.StringsMatch(o.Engineid, other.Engineid) {
		return false
	}
	if !util.StringsMatch(o.Authpwd, other.Authpwd) {
		return false
	}
	if !util.StringsMatch(o.Privpwd, other.Privpwd) {
		return false
	}
	if !util.StringsMatch(o.Authproto, other.Authproto) {
		return false
	}
	if !util.StringsMatch(o.Privproto, other.Privproto) {
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
