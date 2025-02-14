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
	Suffix = []string{"network", "ike", "crypto-profiles", "ipsec-crypto-profiles"}
)

type Entry struct {
	Name     string
	DhGroup  *string
	Lifesize *Lifesize
	Lifetime *Lifetime
	Ah       *Ah
	Esp      *Esp

	Misc map[string][]generic.Xml
}

type Ah struct {
	Authentication []string
}
type Esp struct {
	Authentication []string
	Encryption     []string
}
type Lifesize struct {
	Gb *int64
	Kb *int64
	Mb *int64
	Tb *int64
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
	XMLName  xml.Name     `xml:"entry"`
	Name     string       `xml:"name,attr"`
	DhGroup  *string      `xml:"dh-group,omitempty"`
	Lifesize *LifesizeXml `xml:"lifesize,omitempty"`
	Lifetime *LifetimeXml `xml:"lifetime,omitempty"`
	Ah       *AhXml       `xml:"ah,omitempty"`
	Esp      *EspXml      `xml:"esp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AhXml struct {
	Authentication *util.MemberType `xml:"authentication,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EspXml struct {
	Authentication *util.MemberType `xml:"authentication,omitempty"`
	Encryption     *util.MemberType `xml:"encryption,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type LifesizeXml struct {
	Gb *int64 `xml:"gb,omitempty"`
	Kb *int64 `xml:"kb,omitempty"`
	Mb *int64 `xml:"mb,omitempty"`
	Tb *int64 `xml:"tb,omitempty"`

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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.DhGroup = o.DhGroup
	var nestedLifesize *LifesizeXml
	if o.Lifesize != nil {
		nestedLifesize = &LifesizeXml{}
		if _, ok := o.Misc["Lifesize"]; ok {
			nestedLifesize.Misc = o.Misc["Lifesize"]
		}
		if o.Lifesize.Tb != nil {
			nestedLifesize.Tb = o.Lifesize.Tb
		}
		if o.Lifesize.Gb != nil {
			nestedLifesize.Gb = o.Lifesize.Gb
		}
		if o.Lifesize.Kb != nil {
			nestedLifesize.Kb = o.Lifesize.Kb
		}
		if o.Lifesize.Mb != nil {
			nestedLifesize.Mb = o.Lifesize.Mb
		}
	}
	entry.Lifesize = nestedLifesize

	var nestedLifetime *LifetimeXml
	if o.Lifetime != nil {
		nestedLifetime = &LifetimeXml{}
		if _, ok := o.Misc["Lifetime"]; ok {
			nestedLifetime.Misc = o.Misc["Lifetime"]
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

	var nestedAh *AhXml
	if o.Ah != nil {
		nestedAh = &AhXml{}
		if _, ok := o.Misc["Ah"]; ok {
			nestedAh.Misc = o.Misc["Ah"]
		}
		if o.Ah.Authentication != nil {
			nestedAh.Authentication = util.StrToMem(o.Ah.Authentication)
		}
	}
	entry.Ah = nestedAh

	var nestedEsp *EspXml
	if o.Esp != nil {
		nestedEsp = &EspXml{}
		if _, ok := o.Misc["Esp"]; ok {
			nestedEsp.Misc = o.Misc["Esp"]
		}
		if o.Esp.Encryption != nil {
			nestedEsp.Encryption = util.StrToMem(o.Esp.Encryption)
		}
		if o.Esp.Authentication != nil {
			nestedEsp.Authentication = util.StrToMem(o.Esp.Authentication)
		}
	}
	entry.Esp = nestedEsp

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
		entry.DhGroup = o.DhGroup
		var nestedLifesize *Lifesize
		if o.Lifesize != nil {
			nestedLifesize = &Lifesize{}
			if o.Lifesize.Misc != nil {
				entry.Misc["Lifesize"] = o.Lifesize.Misc
			}
			if o.Lifesize.Gb != nil {
				nestedLifesize.Gb = o.Lifesize.Gb
			}
			if o.Lifesize.Kb != nil {
				nestedLifesize.Kb = o.Lifesize.Kb
			}
			if o.Lifesize.Mb != nil {
				nestedLifesize.Mb = o.Lifesize.Mb
			}
			if o.Lifesize.Tb != nil {
				nestedLifesize.Tb = o.Lifesize.Tb
			}
		}
		entry.Lifesize = nestedLifesize

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

		var nestedAh *Ah
		if o.Ah != nil {
			nestedAh = &Ah{}
			if o.Ah.Misc != nil {
				entry.Misc["Ah"] = o.Ah.Misc
			}
			if o.Ah.Authentication != nil {
				nestedAh.Authentication = util.MemToStr(o.Ah.Authentication)
			}
		}
		entry.Ah = nestedAh

		var nestedEsp *Esp
		if o.Esp != nil {
			nestedEsp = &Esp{}
			if o.Esp.Misc != nil {
				entry.Misc["Esp"] = o.Esp.Misc
			}
			if o.Esp.Authentication != nil {
				nestedEsp.Authentication = util.MemToStr(o.Esp.Authentication)
			}
			if o.Esp.Encryption != nil {
				nestedEsp.Encryption = util.MemToStr(o.Esp.Encryption)
			}
		}
		entry.Esp = nestedEsp

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
	if !util.StringsMatch(a.DhGroup, b.DhGroup) {
		return false
	}
	if !matchLifesize(a.Lifesize, b.Lifesize) {
		return false
	}
	if !matchLifetime(a.Lifetime, b.Lifetime) {
		return false
	}
	if !matchAh(a.Ah, b.Ah) {
		return false
	}
	if !matchEsp(a.Esp, b.Esp) {
		return false
	}

	return true
}

func matchLifesize(a *Lifesize, b *Lifesize) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Gb, b.Gb) {
		return false
	}
	if !util.Ints64Match(a.Kb, b.Kb) {
		return false
	}
	if !util.Ints64Match(a.Mb, b.Mb) {
		return false
	}
	if !util.Ints64Match(a.Tb, b.Tb) {
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
func matchAh(a *Ah, b *Ah) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Authentication, b.Authentication) {
		return false
	}
	return true
}
func matchEsp(a *Esp, b *Esp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.OrderedListsMatch(a.Encryption, b.Encryption) {
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
