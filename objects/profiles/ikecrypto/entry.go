package ikecrypto

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"network", "ike", "crypto-profiles", "ike-crypto-profiles", "$name"}
)

type Entry struct {
	Name                   string
	AuthenticationMultiple *int64
	DhGroup                []string
	Encryption             []string
	Hash                   []string
	Lifetime               *Lifetime
	Misc                   []generic.Xml
}
type Lifetime struct {
	Days    *int64
	Hours   *int64
	Minutes *int64
	Seconds *int64
	Misc    []generic.Xml
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
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	AuthenticationMultiple *int64           `xml:"authentication-multiple,omitempty"`
	DhGroup                *util.MemberType `xml:"dh-group,omitempty"`
	Encryption             *util.MemberType `xml:"encryption,omitempty"`
	Hash                   *util.MemberType `xml:"hash,omitempty"`
	Lifetime               *lifetimeXml     `xml:"lifetime,omitempty"`
	Misc                   []generic.Xml    `xml:",any"`
}
type lifetimeXml struct {
	Days    *int64        `xml:"days,omitempty"`
	Hours   *int64        `xml:"hours,omitempty"`
	Minutes *int64        `xml:"minutes,omitempty"`
	Seconds *int64        `xml:"seconds,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AuthenticationMultiple = s.AuthenticationMultiple
	if s.DhGroup != nil {
		o.DhGroup = util.StrToMem(s.DhGroup)
	}
	if s.Encryption != nil {
		o.Encryption = util.StrToMem(s.Encryption)
	}
	if s.Hash != nil {
		o.Hash = util.StrToMem(s.Hash)
	}
	if s.Lifetime != nil {
		var obj lifetimeXml
		obj.MarshalFromObject(*s.Lifetime)
		o.Lifetime = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var dhGroupVal []string
	if o.DhGroup != nil {
		dhGroupVal = util.MemToStr(o.DhGroup)
	}
	var encryptionVal []string
	if o.Encryption != nil {
		encryptionVal = util.MemToStr(o.Encryption)
	}
	var hashVal []string
	if o.Hash != nil {
		hashVal = util.MemToStr(o.Hash)
	}
	var lifetimeVal *Lifetime
	if o.Lifetime != nil {
		obj, err := o.Lifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lifetimeVal = obj
	}

	result := &Entry{
		Name:                   o.Name,
		AuthenticationMultiple: o.AuthenticationMultiple,
		DhGroup:                dhGroupVal,
		Encryption:             encryptionVal,
		Hash:                   hashVal,
		Lifetime:               lifetimeVal,
		Misc:                   o.Misc,
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

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authentication_multiple" || v == "AuthenticationMultiple" {
		return e.AuthenticationMultiple, nil
	}
	if v == "dh_group" || v == "DhGroup" {
		return e.DhGroup, nil
	}
	if v == "dh_group|LENGTH" || v == "DhGroup|LENGTH" {
		return int64(len(e.DhGroup)), nil
	}
	if v == "encryption" || v == "Encryption" {
		return e.Encryption, nil
	}
	if v == "encryption|LENGTH" || v == "Encryption|LENGTH" {
		return int64(len(e.Encryption)), nil
	}
	if v == "hash" || v == "Hash" {
		return e.Hash, nil
	}
	if v == "hash|LENGTH" || v == "Hash|LENGTH" {
		return int64(len(e.Hash)), nil
	}
	if v == "lifetime" || v == "Lifetime" {
		return e.Lifetime, nil
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
	if !util.Ints64Match(o.AuthenticationMultiple, other.AuthenticationMultiple) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DhGroup, other.DhGroup) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Encryption, other.Encryption) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Hash, other.Hash) {
		return false
	}
	if !o.Lifetime.matches(other.Lifetime) {
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

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
