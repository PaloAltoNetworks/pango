package authprofile

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
	suffix = []string{"network", "routing-profile", "ospf", "auth-profile", "$name"}
)

type Entry struct {
	Name           string
	Md5            []Md5
	Password       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Md5 struct {
	Name           string
	Key            *string
	Preferred      *bool
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
	Md5            *md5ContainerXml `xml:"md5,omitempty"`
	Password       *string          `xml:"password,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type md5ContainerXml struct {
	Entries []md5Xml `xml:"entry"`
}
type md5Xml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Key            *string       `xml:"key,omitempty"`
	Preferred      *string       `xml:"preferred,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Md5 != nil {
		var objs []md5Xml
		for _, elt := range s.Md5 {
			var obj md5Xml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Md5 = &md5ContainerXml{Entries: objs}
	}
	o.Password = s.Password
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var md5Val []Md5
	if o.Md5 != nil {
		for _, elt := range o.Md5.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			md5Val = append(md5Val, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		Md5:            md5Val,
		Password:       o.Password,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *md5Xml) MarshalFromObject(s Md5) {
	o.Name = s.Name
	o.Key = s.Key
	o.Preferred = util.YesNo(s.Preferred, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o md5Xml) UnmarshalToObject() (*Md5, error) {

	result := &Md5{
		Name:           o.Name,
		Key:            o.Key,
		Preferred:      util.AsBool(o.Preferred, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "md5" || v == "Md5" {
		return e.Md5, nil
	}
	if v == "password" || v == "Password" {
		return e.Password, nil
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
	if len(o.Md5) != len(other.Md5) {
		return false
	}
	for idx := range o.Md5 {
		if !o.Md5[idx].matches(&other.Md5[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}

	return true
}

func (o *Md5) matches(other *Md5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}
	if !util.BoolsMatch(o.Preferred, other.Preferred) {
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
