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
	Name        string
	Description *string
	Tags        []string // ordered
	IpNetmask   *string
	IpRange     *string
	Fqdn        *string
	IpWildcard  *string // PAN-OS 9.0

	Misc map[string][]generic.Xml
}

func (e *Entry) CopyMiscFrom(v *Entry) {
	if v == nil || len(v.Misc) == 0 {
		return
	}

	e.Misc = make(map[string][]generic.Xml)
	for key := range v.Misc {
		e.Misc[key] = append([]generic.Xml(nil), v.Misc[key]...)
	}
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" {
		return e.Name, nil
	}
	if v == "description" {
		return e.Description, nil
	}
	if v == "tags" {
		return e.Tags, nil
	}
	if v == "tags|LENGTH" {
		return int64(len(e.Tags)), nil
	}
	if v == "ip_netmask" {
		return e.IpNetmask, nil
	}
	if v == "ip_range" {
		return e.IpRange, nil
	}
	if v == "fqdn" {
		return e.Fqdn, nil
	}
	if v == "ip_wildcard" {
		return e.IpWildcard, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return Entry1Specify, &Entry1Container{}, nil
}

func Entry1Specify(o Entry) (any, error) {
	ans := Entry1{}
	ans.Name = o.Name
	ans.Description = o.Description
	ans.Tags = util.StrToMem(o.Tags)
	ans.IpNetmask = o.IpNetmask
	ans.IpRange = o.IpRange
	ans.Fqdn = o.Fqdn
	ans.IpWildcard = o.IpWildcard

	ans.Misc = o.Misc[fmt.Sprintf("%s\n%s", "Entry", o.Name)]

	return ans, nil
}

type Entry1Container struct {
	Answer []Entry1 `xml:"entry"`
}

func (c *Entry1Container) Normalize() ([]Entry, error) {
	ans := make([]Entry, 0, len(c.Answer))
	for _, var0 := range c.Answer {
		var1 := Entry{
			Misc: make(map[string][]generic.Xml),
		}
		var1.Name = var0.Name
		var1.Description = var0.Description
		var1.IpNetmask = var0.IpNetmask
		var1.IpRange = var0.IpRange
		var1.Fqdn = var0.Fqdn
		var1.IpWildcard = var0.IpWildcard

		var1.Misc[fmt.Sprintf("%s\n%s", "Entry", var0.Name)] = var0.Misc

		ans = append(ans, var1)
	}

	return ans, nil
}

type Entry1 struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	IpNetmask   *string          `xml:"ip-netmask"`
	IpRange     *string          `xml:"ip-range"`
	Fqdn        *string          `xml:"fqdn"`
	IpWildcard  *string          `xml:"ip-wildcard"`
	Description *string          `xml:"description,omitempty"`
	Tags        *util.MemberType `xml:"tag"`

	Misc []generic.Xml `xml:",any"`
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.

	if !util.OptionalStringsMatch(a.Description, b.Description) {
		return false
	}

	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}

	if !util.OptionalStringsMatch(a.IpNetmask, b.IpNetmask) {
		return false
	}

	if !util.OptionalStringsMatch(a.IpRange, b.IpRange) {
		return false
	}

	if !util.OptionalStringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}

	if !util.OptionalStringsMatch(a.IpWildcard, b.IpWildcard) {
		return false
	}

	return true
}
