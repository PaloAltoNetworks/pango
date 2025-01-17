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
	Ip                         []Ip
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
type Ip struct {
	Name string
}
type Ipv6 struct {
	Address     []Ipv6Address
	Enabled     *bool
	InterfaceId *string
}
type Ipv6Address struct {
	Anycast           *Ipv6AddressAnycast
	EnableOnInterface *bool
	Name              string
	Prefix            *Ipv6AddressPrefix
}
type Ipv6AddressAnycast struct {
}
type Ipv6AddressPrefix struct {
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
	Ip                         []IpXml          `xml:"ip>entry,omitempty"`
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
type IpXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6Xml struct {
	Address     []Ipv6AddressXml `xml:"address>entry,omitempty"`
	Enabled     *string          `xml:"enabled,omitempty"`
	InterfaceId *string          `xml:"interface-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressXml struct {
	Anycast           *Ipv6AddressAnycastXml `xml:"anycast,omitempty"`
	EnableOnInterface *string                `xml:"enable-on-interface,omitempty"`
	XMLName           xml.Name               `xml:"entry"`
	Name              string                 `xml:"name,attr"`
	Prefix            *Ipv6AddressPrefixXml  `xml:"prefix,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressAnycastXml struct {
	Misc []generic.Xml `xml:",any"`
}
type Ipv6AddressPrefixXml struct {
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
	if v == "ip" || v == "Ip" {
		return e.Ip, nil
	}
	if v == "ip|LENGTH" || v == "Ip|LENGTH" {
		return int64(len(e.Ip)), nil
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
	var nestedIpCol []IpXml
	if o.Ip != nil {
		nestedIpCol = []IpXml{}
		for _, oIp := range o.Ip {
			nestedIp := IpXml{}
			if _, ok := o.Misc["Ip"]; ok {
				nestedIp.Misc = o.Misc["Ip"]
			}
			if oIp.Name != "" {
				nestedIp.Name = oIp.Name
			}
			nestedIpCol = append(nestedIpCol, nestedIp)
		}
		entry.Ip = nestedIpCol
	}

	var nestedIpv6 *Ipv6Xml
	if o.Ipv6 != nil {
		nestedIpv6 = &Ipv6Xml{}
		if _, ok := o.Misc["Ipv6"]; ok {
			nestedIpv6.Misc = o.Misc["Ipv6"]
		}
		if o.Ipv6.InterfaceId != nil {
			nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
		}
		if o.Ipv6.Address != nil {
			nestedIpv6.Address = []Ipv6AddressXml{}
			for _, oIpv6Address := range o.Ipv6.Address {
				nestedIpv6Address := Ipv6AddressXml{}
				if _, ok := o.Misc["Ipv6Address"]; ok {
					nestedIpv6Address.Misc = o.Misc["Ipv6Address"]
				}
				if oIpv6Address.Anycast != nil {
					nestedIpv6Address.Anycast = &Ipv6AddressAnycastXml{}
					if _, ok := o.Misc["Ipv6AddressAnycast"]; ok {
						nestedIpv6Address.Anycast.Misc = o.Misc["Ipv6AddressAnycast"]
					}
				}
				if oIpv6Address.Name != "" {
					nestedIpv6Address.Name = oIpv6Address.Name
				}
				if oIpv6Address.EnableOnInterface != nil {
					nestedIpv6Address.EnableOnInterface = util.YesNo(oIpv6Address.EnableOnInterface, nil)
				}
				if oIpv6Address.Prefix != nil {
					nestedIpv6Address.Prefix = &Ipv6AddressPrefixXml{}
					if _, ok := o.Misc["Ipv6AddressPrefix"]; ok {
						nestedIpv6Address.Prefix.Misc = o.Misc["Ipv6AddressPrefix"]
					}
				}
				nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
			}
		}
		if o.Ipv6.Enabled != nil {
			nestedIpv6.Enabled = util.YesNo(o.Ipv6.Enabled, nil)
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
		var nestedIpCol []Ip
		if o.Ip != nil {
			nestedIpCol = []Ip{}
			for _, oIp := range o.Ip {
				nestedIp := Ip{}
				if oIp.Misc != nil {
					entry.Misc["Ip"] = oIp.Misc
				}
				if oIp.Name != "" {
					nestedIp.Name = oIp.Name
				}
				nestedIpCol = append(nestedIpCol, nestedIp)
			}
			entry.Ip = nestedIpCol
		}

		var nestedIpv6 *Ipv6
		if o.Ipv6 != nil {
			nestedIpv6 = &Ipv6{}
			if o.Ipv6.Misc != nil {
				entry.Misc["Ipv6"] = o.Ipv6.Misc
			}
			if o.Ipv6.Address != nil {
				nestedIpv6.Address = []Ipv6Address{}
				for _, oIpv6Address := range o.Ipv6.Address {
					nestedIpv6Address := Ipv6Address{}
					if oIpv6Address.Misc != nil {
						entry.Misc["Ipv6Address"] = oIpv6Address.Misc
					}
					if oIpv6Address.EnableOnInterface != nil {
						nestedIpv6Address.EnableOnInterface = util.AsBool(oIpv6Address.EnableOnInterface, nil)
					}
					if oIpv6Address.Prefix != nil {
						nestedIpv6Address.Prefix = &Ipv6AddressPrefix{}
						if oIpv6Address.Prefix.Misc != nil {
							entry.Misc["Ipv6AddressPrefix"] = oIpv6Address.Prefix.Misc
						}
					}
					if oIpv6Address.Anycast != nil {
						nestedIpv6Address.Anycast = &Ipv6AddressAnycast{}
						if oIpv6Address.Anycast.Misc != nil {
							entry.Misc["Ipv6AddressAnycast"] = oIpv6Address.Anycast.Misc
						}
					}
					if oIpv6Address.Name != "" {
						nestedIpv6Address.Name = oIpv6Address.Name
					}
					nestedIpv6.Address = append(nestedIpv6.Address, nestedIpv6Address)
				}
			}
			if o.Ipv6.Enabled != nil {
				nestedIpv6.Enabled = util.AsBool(o.Ipv6.Enabled, nil)
			}
			if o.Ipv6.InterfaceId != nil {
				nestedIpv6.InterfaceId = o.Ipv6.InterfaceId
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
	if !matchIp(a.Ip, b.Ip) {
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

func matchAdjustTcpMss(a *AdjustTcpMss, b *AdjustTcpMss) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Ipv4MssAdjustment, b.Ipv4MssAdjustment) {
		return false
	}
	if !util.Ints64Match(a.Ipv6MssAdjustment, b.Ipv6MssAdjustment) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchIp(a []Ip, b []Ip) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchIpv6AddressPrefix(a *Ipv6AddressPrefix, b *Ipv6AddressPrefix) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6AddressAnycast(a *Ipv6AddressAnycast, b *Ipv6AddressAnycast) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchIpv6Address(a []Ipv6Address, b []Ipv6Address) bool {
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
			if !matchIpv6AddressPrefix(a.Prefix, b.Prefix) {
				return false
			}
			if !matchIpv6AddressAnycast(a.Anycast, b.Anycast) {
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
	if !matchIpv6Address(a.Address, b.Address) {
		return false
	}
	if !util.BoolsMatch(a.Enabled, b.Enabled) {
		return false
	}
	if !util.StringsMatch(a.InterfaceId, b.InterfaceId) {
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
