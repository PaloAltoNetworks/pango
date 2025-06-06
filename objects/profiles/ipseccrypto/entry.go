package ipseccrypto

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
	suffix = []string{"network", "ike", "crypto-profiles", "ipsec-crypto-profiles", "$name"}
)

type Entry struct {
	Name     string
	DhGroup  *string
	Lifesize *Lifesize
	Lifetime *Lifetime
	Ah       *Ah
	Esp      *Esp
	Misc     []generic.Xml
}
type Lifesize struct {
	Gb   *int64
	Kb   *int64
	Mb   *int64
	Tb   *int64
	Misc []generic.Xml
}
type Lifetime struct {
	Days    *int64
	Hours   *int64
	Minutes *int64
	Seconds *int64
	Misc    []generic.Xml
}
type Ah struct {
	Authentication []string
	Misc           []generic.Xml
}
type Esp struct {
	Authentication []string
	Encryption     []string
	Misc           []generic.Xml
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
	XMLName  xml.Name      `xml:"entry"`
	Name     string        `xml:"name,attr"`
	DhGroup  *string       `xml:"dh-group,omitempty"`
	Lifesize *lifesizeXml  `xml:"lifesize,omitempty"`
	Lifetime *lifetimeXml  `xml:"lifetime,omitempty"`
	Ah       *ahXml        `xml:"ah,omitempty"`
	Esp      *espXml       `xml:"esp,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type lifesizeXml struct {
	Gb   *int64        `xml:"gb,omitempty"`
	Kb   *int64        `xml:"kb,omitempty"`
	Mb   *int64        `xml:"mb,omitempty"`
	Tb   *int64        `xml:"tb,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type lifetimeXml struct {
	Days    *int64        `xml:"days,omitempty"`
	Hours   *int64        `xml:"hours,omitempty"`
	Minutes *int64        `xml:"minutes,omitempty"`
	Seconds *int64        `xml:"seconds,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type ahXml struct {
	Authentication *util.MemberType `xml:"authentication,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
}
type espXml struct {
	Authentication *util.MemberType `xml:"authentication,omitempty"`
	Encryption     *util.MemberType `xml:"encryption,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DhGroup = s.DhGroup
	if s.Lifesize != nil {
		var obj lifesizeXml
		obj.MarshalFromObject(*s.Lifesize)
		o.Lifesize = &obj
	}
	if s.Lifetime != nil {
		var obj lifetimeXml
		obj.MarshalFromObject(*s.Lifetime)
		o.Lifetime = &obj
	}
	if s.Ah != nil {
		var obj ahXml
		obj.MarshalFromObject(*s.Ah)
		o.Ah = &obj
	}
	if s.Esp != nil {
		var obj espXml
		obj.MarshalFromObject(*s.Esp)
		o.Esp = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var lifesizeVal *Lifesize
	if o.Lifesize != nil {
		obj, err := o.Lifesize.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lifesizeVal = obj
	}
	var lifetimeVal *Lifetime
	if o.Lifetime != nil {
		obj, err := o.Lifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lifetimeVal = obj
	}
	var ahVal *Ah
	if o.Ah != nil {
		obj, err := o.Ah.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ahVal = obj
	}
	var espVal *Esp
	if o.Esp != nil {
		obj, err := o.Esp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		espVal = obj
	}

	result := &Entry{
		Name:     o.Name,
		DhGroup:  o.DhGroup,
		Lifesize: lifesizeVal,
		Lifetime: lifetimeVal,
		Ah:       ahVal,
		Esp:      espVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *lifesizeXml) MarshalFromObject(s Lifesize) {
	o.Gb = s.Gb
	o.Kb = s.Kb
	o.Mb = s.Mb
	o.Tb = s.Tb
	o.Misc = s.Misc
}

func (o lifesizeXml) UnmarshalToObject() (*Lifesize, error) {

	result := &Lifesize{
		Gb:   o.Gb,
		Kb:   o.Kb,
		Mb:   o.Mb,
		Tb:   o.Tb,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *lifetimeXml) MarshalFromObject(s Lifetime) {
	o.Days = s.Days
	o.Hours = s.Hours
	o.Minutes = s.Minutes
	o.Seconds = s.Seconds
	o.Misc = s.Misc
}

func (o lifetimeXml) UnmarshalToObject() (*Lifetime, error) {

	result := &Lifetime{
		Days:    o.Days,
		Hours:   o.Hours,
		Minutes: o.Minutes,
		Seconds: o.Seconds,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *ahXml) MarshalFromObject(s Ah) {
	if s.Authentication != nil {
		o.Authentication = util.StrToMem(s.Authentication)
	}
	o.Misc = s.Misc
}

func (o ahXml) UnmarshalToObject() (*Ah, error) {
	var authenticationVal []string
	if o.Authentication != nil {
		authenticationVal = util.MemToStr(o.Authentication)
	}

	result := &Ah{
		Authentication: authenticationVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *espXml) MarshalFromObject(s Esp) {
	if s.Authentication != nil {
		o.Authentication = util.StrToMem(s.Authentication)
	}
	if s.Encryption != nil {
		o.Encryption = util.StrToMem(s.Encryption)
	}
	o.Misc = s.Misc
}

func (o espXml) UnmarshalToObject() (*Esp, error) {
	var authenticationVal []string
	if o.Authentication != nil {
		authenticationVal = util.MemToStr(o.Authentication)
	}
	var encryptionVal []string
	if o.Encryption != nil {
		encryptionVal = util.MemToStr(o.Encryption)
	}

	result := &Esp{
		Authentication: authenticationVal,
		Encryption:     encryptionVal,
		Misc:           o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "dh_group" || v == "DhGroup" {
		return e.DhGroup, nil
	}
	if v == "lifesize" || v == "Lifesize" {
		return e.Lifesize, nil
	}
	if v == "lifetime" || v == "Lifetime" {
		return e.Lifetime, nil
	}
	if v == "ah" || v == "Ah" {
		return e.Ah, nil
	}
	if v == "esp" || v == "Esp" {
		return e.Esp, nil
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
	if !util.StringsMatch(o.DhGroup, other.DhGroup) {
		return false
	}
	if !o.Lifesize.matches(other.Lifesize) {
		return false
	}
	if !o.Lifetime.matches(other.Lifetime) {
		return false
	}
	if !o.Ah.matches(other.Ah) {
		return false
	}
	if !o.Esp.matches(other.Esp) {
		return false
	}

	return true
}

func (o *Lifesize) matches(other *Lifesize) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Gb, other.Gb) {
		return false
	}
	if !util.Ints64Match(o.Kb, other.Kb) {
		return false
	}
	if !util.Ints64Match(o.Mb, other.Mb) {
		return false
	}
	if !util.Ints64Match(o.Tb, other.Tb) {
		return false
	}

	return true
}

func (o *Lifetime) matches(other *Lifetime) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Days, other.Days) {
		return false
	}
	if !util.Ints64Match(o.Hours, other.Hours) {
		return false
	}
	if !util.Ints64Match(o.Minutes, other.Minutes) {
		return false
	}
	if !util.Ints64Match(o.Seconds, other.Seconds) {
		return false
	}

	return true
}

func (o *Ah) matches(other *Ah) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Authentication, other.Authentication) {
		return false
	}

	return true
}

func (o *Esp) matches(other *Esp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Authentication, other.Authentication) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Encryption, other.Encryption) {
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
