package loopback

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
	Suffix = []string{"network", "interface", "loopback", "units"}
)

type Entry struct {
	Name                       string
	AdjustTcpMss               *AdjustTcpMss
	Comment                    *string
	InterfaceManagementProfile *string
	Ips                        []string
	Ipv6                       *Ipv6
	Mtu                        *int64
	NetflowProfile             *string

	Misc map[string][]generic.Xml
}

type AdjustTcpMss struct {
	Enable            *bool
	Ipv4MssAdjustment *int64
	Ipv6MssAdjustment *int64
}
type Ipv6 struct {
	Addresses []Ipv6Addresses
	Enabled   *bool
}
type Ipv6Addresses struct {
	EnableOnInterface *bool
	Name              string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                    xml.Name         `xml:"entry"`
	Name                       string           `xml:"name,attr"`
	AdjustTcpMss               *AdjustTcpMssXml `xml:"adjust-tcp-mss,omitempty"`
	Comment                    *string          `xml:"comment,omitempty"`
	InterfaceManagementProfile *string          `xml:"interface-management-profile,omitempty"`
	Ips                        *util.EntryType  `xml:"ip,omitempty"`
	Ipv6                       *Ipv6Xml         `xml:"ipv6,omitempty"`
	Mtu                        *int64           `xml:"mtu,omitempty"`
	NetflowProfile             *string          `xml:"netflow-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type AdjustTcpMssXml struct {
	Enable            *string `xml:"enable,omitempty"`
	Ipv4MssAdjustment *int64  `xml:"ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjustment *int64  `xml:"ipv6-mss-adjustment,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6Xml struct {
	Addresses []Ipv6AddressesXml `xml:"address>entry,omitempty"`
	Enabled   *string            `xml:"enabled,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressesXml struct {
	EnableOnInterface *string  `xml:"enable-on-interface,omitempty"`
	XMLName           xml.Name `xml:"entry"`
	Name              string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "adjust_tcp_mss" || v == "AdjustTcpMss" {
		return e.AdjustTcpMss, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "interface_management_profile" || v == "InterfaceManagementProfile" {
		return e.InterfaceManagementProfile, nil
	}
	if v == "ips" || v == "Ips" {
		return e.Ips, nil
	}
	if v == "ips|LENGTH" || v == "Ips|LENGTH" {
		return int64(len(e.Ips)), nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "mtu" || v == "Mtu" {
		return e.Mtu, nil
	}
	if v == "netflow_profile" || v == "NetflowProfile" {
		return e.NetflowProfile, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyEntry, &entryXmlContainer{}, nil
}

func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}

	entry.Name = o.Name
	var nestedAdjustTcpMss *AdjustTcpMssXml
	if o.AdjustTcpMss != nil {
		nestedAdjustTcpMss = &AdjustTcpMssXml{}
		if _, ok := o.Misc["AdjustTcpMss"]; ok {
			nestedAdjustTcpMss.Misc = o.Misc["AdjustTcpMss"]
		}
		if o.AdjustTcpMss.Enable != nil {
			nestedAdjustTcpMss.Enable = util.YesNo(o.AdjustTcpMss.Enable, nil)
		}
		if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
		}
		if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
			nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
		}
	}
	entry.AdjustTcpMss = nestedAdjustTcpMss

	entry.Comment = o.Comment
	entry.InterfaceManagementProfile = o.InterfaceManagementProfile
	entry.Ips = util.StrToEnt(o.Ips)
	var nestedIpv6 *Ipv6Xml
	if o.Ipv6 != nil {
		nestedIpv6 = &Ipv6Xml{}
		if _, ok := o.Misc["Ipv6"]; ok {
			nestedIpv6.Misc = o.Misc["Ipv6"]
		}
		if o.Ipv6.Enabled != nil {
			nestedIpv6.Enabled = util.YesNo(o.Ipv6.Enabled, nil)
		}
		if o.Ipv6.Addresses != nil {
			nestedIpv6.Addresses = []Ipv6AddressesXml{}
			for _, oIpv6Addresses := range o.Ipv6.Addresses {
				nestedIpv6Addresses := Ipv6AddressesXml{}
				if _, ok := o.Misc["Ipv6Addresses"]; ok {
					nestedIpv6Addresses.Misc = o.Misc["Ipv6Addresses"]
				}
				if oIpv6Addresses.EnableOnInterface != nil {
					nestedIpv6Addresses.EnableOnInterface = util.YesNo(oIpv6Addresses.EnableOnInterface, nil)
				}
				if oIpv6Addresses.Name != "" {
					nestedIpv6Addresses.Name = oIpv6Addresses.Name
				}
				nestedIpv6.Addresses = append(nestedIpv6.Addresses, nestedIpv6Addresses)
			}
		}
	}
	entry.Ipv6 = nestedIpv6

	entry.Mtu = o.Mtu
	entry.NetflowProfile = o.NetflowProfile

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
		var nestedAdjustTcpMss *AdjustTcpMss
		if o.AdjustTcpMss != nil {
			nestedAdjustTcpMss = &AdjustTcpMss{}
			if o.AdjustTcpMss.Misc != nil {
				entry.Misc["AdjustTcpMss"] = o.AdjustTcpMss.Misc
			}
			if o.AdjustTcpMss.Enable != nil {
				nestedAdjustTcpMss.Enable = util.AsBool(o.AdjustTcpMss.Enable, nil)
			}
			if o.AdjustTcpMss.Ipv4MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv4MssAdjustment = o.AdjustTcpMss.Ipv4MssAdjustment
			}
			if o.AdjustTcpMss.Ipv6MssAdjustment != nil {
				nestedAdjustTcpMss.Ipv6MssAdjustment = o.AdjustTcpMss.Ipv6MssAdjustment
			}
		}
		entry.AdjustTcpMss = nestedAdjustTcpMss

		entry.Comment = o.Comment
		entry.InterfaceManagementProfile = o.InterfaceManagementProfile
		entry.Ips = util.EntToStr(o.Ips)
		var nestedIpv6 *Ipv6
		if o.Ipv6 != nil {
			nestedIpv6 = &Ipv6{}
			if o.Ipv6.Misc != nil {
				entry.Misc["Ipv6"] = o.Ipv6.Misc
			}
			if o.Ipv6.Addresses != nil {
				nestedIpv6.Addresses = []Ipv6Addresses{}
				for _, oIpv6Addresses := range o.Ipv6.Addresses {
					nestedIpv6Addresses := Ipv6Addresses{}
					if oIpv6Addresses.Misc != nil {
						entry.Misc["Ipv6Addresses"] = oIpv6Addresses.Misc
					}
					if oIpv6Addresses.EnableOnInterface != nil {
						nestedIpv6Addresses.EnableOnInterface = util.AsBool(oIpv6Addresses.EnableOnInterface, nil)
					}
					if oIpv6Addresses.Name != "" {
						nestedIpv6Addresses.Name = oIpv6Addresses.Name
					}
					nestedIpv6.Addresses = append(nestedIpv6.Addresses, nestedIpv6Addresses)
				}
			}
			if o.Ipv6.Enabled != nil {
				nestedIpv6.Enabled = util.AsBool(o.Ipv6.Enabled, nil)
			}
		}
		entry.Ipv6 = nestedIpv6

		entry.Mtu = o.Mtu
		entry.NetflowProfile = o.NetflowProfile

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
	if !matchAdjustTcpMss(a.AdjustTcpMss, b.AdjustTcpMss) {
		return false
	}
	if !util.StringsMatch(a.Comment, b.Comment) {
		return false
	}
	if !util.StringsMatch(a.InterfaceManagementProfile, b.InterfaceManagementProfile) {
		return false
	}
	if !util.OrderedListsMatch(a.Ips, b.Ips) {
		return false
	}
	if !matchIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	if !util.Ints64Match(a.Mtu, b.Mtu) {
		return false
	}
	if !util.StringsMatch(a.NetflowProfile, b.NetflowProfile) {
		return false
	}

	return true
}

func matchIpv6Addresses(a []Ipv6Addresses, b []Ipv6Addresses) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.BoolsMatch(a.EnableOnInterface, b.EnableOnInterface) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchIpv6(a *Ipv6, b *Ipv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !matchIpv6Addresses(a.Addresses, b.Addresses) {
		return false
	}
	return true
}
func matchAdjustTcpMss(a *AdjustTcpMss, b *AdjustTcpMss) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.Ipv4MssAdjustment, b.Ipv4MssAdjustment) {
		return false
	}
	if !util.Ints64Match(a.Ipv6MssAdjustment, b.Ipv6MssAdjustment) {
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
