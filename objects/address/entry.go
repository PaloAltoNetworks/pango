package address

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
	Suffix = []string{"address"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Tag             []string
	Fqdn            *string
	IpNetmask       *string
	IpRange         *string
	IpWildcard      *string

	Misc map[string][]generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	Description     *string          `xml:"description,omitempty"`
	DisableOverride *string          `xml:"disable-override,omitempty"`
	Tag             *util.MemberType `xml:"tag,omitempty"`
	Fqdn            *string          `xml:"fqdn,omitempty"`
	IpNetmask       *string          `xml:"ip-netmask,omitempty"`
	IpRange         *string          `xml:"ip-range,omitempty"`
	IpWildcard      *string          `xml:"ip-wildcard,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
	}
	if v == "fqdn" || v == "Fqdn" {
		return e.Fqdn, nil
	}
	if v == "ip_netmask" || v == "IpNetmask" {
		return e.IpNetmask, nil
	}
	if v == "ip_range" || v == "IpRange" {
		return e.IpRange, nil
	}
	if v == "ip_wildcard" || v == "IpWildcard" {
		return e.IpWildcard, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.Tag = util.StrToMem(o.Tag)
	entry.Fqdn = o.Fqdn
	entry.IpNetmask = o.IpNetmask
	entry.IpRange = o.IpRange
	entry.IpWildcard = o.IpWildcard

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
		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.Tag = util.MemToStr(o.Tag)
		entry.Fqdn = o.Fqdn
		entry.IpNetmask = o.IpNetmask
		entry.IpRange = o.IpRange
		entry.IpWildcard = o.IpWildcard

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
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch(a.Tag, b.Tag) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.IpNetmask, b.IpNetmask) {
		return false
	}
	if !util.StringsMatch(a.IpRange, b.IpRange) {
		return false
	}
	if !util.StringsMatch(a.IpWildcard, b.IpWildcard) {
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
