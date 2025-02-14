package ikecrypto

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
	Suffix = []string{"network", "ike", "crypto-profiles", "ike-crypto-profiles"}
)

type Entry struct {
	Name                   string
	AuthenticationMultiple *int64
	DhGroup                []string
	Encryption             []string
	Hash                   []string
	Lifetime               *Lifetime

	Misc map[string][]generic.Xml
}

type Lifetime struct {
	Days    *int64
	Hours   *int64
	Minutes *int64
	Seconds *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	AuthenticationMultiple *int64           `xml:"authentication-multiple,omitempty"`
	DhGroup                *util.MemberType `xml:"dh-group,omitempty"`
	Encryption             *util.MemberType `xml:"encryption,omitempty"`
	Hash                   *util.MemberType `xml:"hash,omitempty"`
	Lifetime               *LifetimeXml     `xml:"lifetime,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type LifetimeXml struct {
	Days    *int64 `xml:"days,omitempty"`
	Hours   *int64 `xml:"hours,omitempty"`
	Minutes *int64 `xml:"minutes,omitempty"`
	Seconds *int64 `xml:"seconds,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.AuthenticationMultiple = o.AuthenticationMultiple
	entry.DhGroup = util.StrToMem(o.DhGroup)
	entry.Encryption = util.StrToMem(o.Encryption)
	entry.Hash = util.StrToMem(o.Hash)
	var nestedLifetime *LifetimeXml
	if o.Lifetime != nil {
		nestedLifetime = &LifetimeXml{}
		if _, ok := o.Misc["Lifetime"]; ok {
			nestedLifetime.Misc = o.Misc["Lifetime"]
		}
		if o.Lifetime.Minutes != nil {
			nestedLifetime.Minutes = o.Lifetime.Minutes
		}
		if o.Lifetime.Seconds != nil {
			nestedLifetime.Seconds = o.Lifetime.Seconds
		}
		if o.Lifetime.Days != nil {
			nestedLifetime.Days = o.Lifetime.Days
		}
		if o.Lifetime.Hours != nil {
			nestedLifetime.Hours = o.Lifetime.Hours
		}
	}
	entry.Lifetime = nestedLifetime

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		entry.AuthenticationMultiple = o.AuthenticationMultiple
		entry.DhGroup = util.MemToStr(o.DhGroup)
		entry.Encryption = util.MemToStr(o.Encryption)
		entry.Hash = util.MemToStr(o.Hash)
		var nestedLifetime *Lifetime
		if o.Lifetime != nil {
			nestedLifetime = &Lifetime{}
			if o.Lifetime.Misc != nil {
				entry.Misc["Lifetime"] = o.Lifetime.Misc
			}
			if o.Lifetime.Days != nil {
				nestedLifetime.Days = o.Lifetime.Days
			}
			if o.Lifetime.Hours != nil {
				nestedLifetime.Hours = o.Lifetime.Hours
			}
			if o.Lifetime.Minutes != nil {
				nestedLifetime.Minutes = o.Lifetime.Minutes
			}
			if o.Lifetime.Seconds != nil {
				nestedLifetime.Seconds = o.Lifetime.Seconds
			}
		}
		entry.Lifetime = nestedLifetime

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !util.Ints64Match(a.AuthenticationMultiple, b.AuthenticationMultiple) {
		return false
	}
	if !util.OrderedListsMatch(a.DhGroup, b.DhGroup) {
		return false
	}
	if !util.OrderedListsMatch(a.Encryption, b.Encryption) {
		return false
	}
	if !util.OrderedListsMatch(a.Hash, b.Hash) {
		return false
	}
	if !matchLifetime(a.Lifetime, b.Lifetime) {
		return false
	}

	return true
}

func matchLifetime(a *Lifetime, b *Lifetime) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Days, b.Days) {
		return false
	}
	if !util.Ints64Match(a.Hours, b.Hours) {
		return false
	}
	if !util.Ints64Match(a.Minutes, b.Minutes) {
		return false
	}
	if !util.Ints64Match(a.Seconds, b.Seconds) {
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
