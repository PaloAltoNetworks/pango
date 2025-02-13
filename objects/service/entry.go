package service

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
	Suffix = []string{"service"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Protocol        *Protocol
	Tag             []string

	Misc map[string][]generic.Xml
}

type Protocol struct {
	Tcp *ProtocolTcp
	Udp *ProtocolUdp
}
type ProtocolTcp struct {
	Override   *ProtocolTcpOverride
	Port       *string
	SourcePort *string
}
type ProtocolTcpOverride struct {
	HalfcloseTimeout *int64
	Timeout          *int64
	TimewaitTimeout  *int64
}
type ProtocolUdp struct {
	Override   *ProtocolUdpOverride
	Port       *string
	SourcePort *string
}
type ProtocolUdpOverride struct {
	Timeout *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	Description     *string          `xml:"description,omitempty"`
	DisableOverride *string          `xml:"disable-override,omitempty"`
	Protocol        *ProtocolXml     `xml:"protocol,omitempty"`
	Tag             *util.MemberType `xml:"tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolXml struct {
	Tcp *ProtocolTcpXml `xml:"tcp,omitempty"`
	Udp *ProtocolUdpXml `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolTcpXml struct {
	Override   *ProtocolTcpOverrideXml `xml:"override>yes,omitempty"`
	Port       *string                 `xml:"port,omitempty"`
	SourcePort *string                 `xml:"source-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolTcpOverrideXml struct {
	HalfcloseTimeout *int64 `xml:"halfclose-timeout,omitempty"`
	Timeout          *int64 `xml:"timeout,omitempty"`
	TimewaitTimeout  *int64 `xml:"timewait-timeout,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolUdpXml struct {
	Override   *ProtocolUdpOverrideXml `xml:"override>yes,omitempty"`
	Port       *string                 `xml:"port,omitempty"`
	SourcePort *string                 `xml:"source-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolUdpOverrideXml struct {
	Timeout *int64 `xml:"timeout,omitempty"`

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
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
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
	var nestedProtocol *ProtocolXml
	if o.Protocol != nil {
		nestedProtocol = &ProtocolXml{}
		if _, ok := o.Misc["Protocol"]; ok {
			nestedProtocol.Misc = o.Misc["Protocol"]
		}
		if o.Protocol.Tcp != nil {
			nestedProtocol.Tcp = &ProtocolTcpXml{}
			if _, ok := o.Misc["ProtocolTcp"]; ok {
				nestedProtocol.Tcp.Misc = o.Misc["ProtocolTcp"]
			}
			if o.Protocol.Tcp.Override != nil {
				nestedProtocol.Tcp.Override = &ProtocolTcpOverrideXml{}
				if _, ok := o.Misc["ProtocolTcpOverride"]; ok {
					nestedProtocol.Tcp.Override.Misc = o.Misc["ProtocolTcpOverride"]
				}
				if o.Protocol.Tcp.Override.HalfcloseTimeout != nil {
					nestedProtocol.Tcp.Override.HalfcloseTimeout = o.Protocol.Tcp.Override.HalfcloseTimeout
				}
				if o.Protocol.Tcp.Override.Timeout != nil {
					nestedProtocol.Tcp.Override.Timeout = o.Protocol.Tcp.Override.Timeout
				}
				if o.Protocol.Tcp.Override.TimewaitTimeout != nil {
					nestedProtocol.Tcp.Override.TimewaitTimeout = o.Protocol.Tcp.Override.TimewaitTimeout
				}
			}
			if o.Protocol.Tcp.Port != nil {
				nestedProtocol.Tcp.Port = o.Protocol.Tcp.Port
			}
			if o.Protocol.Tcp.SourcePort != nil {
				nestedProtocol.Tcp.SourcePort = o.Protocol.Tcp.SourcePort
			}
		}
		if o.Protocol.Udp != nil {
			nestedProtocol.Udp = &ProtocolUdpXml{}
			if _, ok := o.Misc["ProtocolUdp"]; ok {
				nestedProtocol.Udp.Misc = o.Misc["ProtocolUdp"]
			}
			if o.Protocol.Udp.Override != nil {
				nestedProtocol.Udp.Override = &ProtocolUdpOverrideXml{}
				if _, ok := o.Misc["ProtocolUdpOverride"]; ok {
					nestedProtocol.Udp.Override.Misc = o.Misc["ProtocolUdpOverride"]
				}
				if o.Protocol.Udp.Override.Timeout != nil {
					nestedProtocol.Udp.Override.Timeout = o.Protocol.Udp.Override.Timeout
				}
			}
			if o.Protocol.Udp.Port != nil {
				nestedProtocol.Udp.Port = o.Protocol.Udp.Port
			}
			if o.Protocol.Udp.SourcePort != nil {
				nestedProtocol.Udp.SourcePort = o.Protocol.Udp.SourcePort
			}
		}
	}
	entry.Protocol = nestedProtocol

	entry.Tag = util.StrToMem(o.Tag)

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
		var nestedProtocol *Protocol
		if o.Protocol != nil {
			nestedProtocol = &Protocol{}
			if o.Protocol.Misc != nil {
				entry.Misc["Protocol"] = o.Protocol.Misc
			}
			if o.Protocol.Tcp != nil {
				nestedProtocol.Tcp = &ProtocolTcp{}
				if o.Protocol.Tcp.Misc != nil {
					entry.Misc["ProtocolTcp"] = o.Protocol.Tcp.Misc
				}
				if o.Protocol.Tcp.Override != nil {
					nestedProtocol.Tcp.Override = &ProtocolTcpOverride{}
					if o.Protocol.Tcp.Override.Misc != nil {
						entry.Misc["ProtocolTcpOverride"] = o.Protocol.Tcp.Override.Misc
					}
					if o.Protocol.Tcp.Override.HalfcloseTimeout != nil {
						nestedProtocol.Tcp.Override.HalfcloseTimeout = o.Protocol.Tcp.Override.HalfcloseTimeout
					}
					if o.Protocol.Tcp.Override.Timeout != nil {
						nestedProtocol.Tcp.Override.Timeout = o.Protocol.Tcp.Override.Timeout
					}
					if o.Protocol.Tcp.Override.TimewaitTimeout != nil {
						nestedProtocol.Tcp.Override.TimewaitTimeout = o.Protocol.Tcp.Override.TimewaitTimeout
					}
				}
				if o.Protocol.Tcp.Port != nil {
					nestedProtocol.Tcp.Port = o.Protocol.Tcp.Port
				}
				if o.Protocol.Tcp.SourcePort != nil {
					nestedProtocol.Tcp.SourcePort = o.Protocol.Tcp.SourcePort
				}
			}
			if o.Protocol.Udp != nil {
				nestedProtocol.Udp = &ProtocolUdp{}
				if o.Protocol.Udp.Misc != nil {
					entry.Misc["ProtocolUdp"] = o.Protocol.Udp.Misc
				}
				if o.Protocol.Udp.Override != nil {
					nestedProtocol.Udp.Override = &ProtocolUdpOverride{}
					if o.Protocol.Udp.Override.Misc != nil {
						entry.Misc["ProtocolUdpOverride"] = o.Protocol.Udp.Override.Misc
					}
					if o.Protocol.Udp.Override.Timeout != nil {
						nestedProtocol.Udp.Override.Timeout = o.Protocol.Udp.Override.Timeout
					}
				}
				if o.Protocol.Udp.Port != nil {
					nestedProtocol.Udp.Port = o.Protocol.Udp.Port
				}
				if o.Protocol.Udp.SourcePort != nil {
					nestedProtocol.Udp.SourcePort = o.Protocol.Udp.SourcePort
				}
			}
		}
		entry.Protocol = nestedProtocol

		entry.Tag = util.MemToStr(o.Tag)

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
	if !matchProtocol(a.Protocol, b.Protocol) {
		return false
	}
	if !util.OrderedListsMatch(a.Tag, b.Tag) {
		return false
	}

	return true
}

func matchProtocolTcpOverride(a *ProtocolTcpOverride, b *ProtocolTcpOverride) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.TimewaitTimeout, b.TimewaitTimeout) {
		return false
	}
	if !util.Ints64Match(a.HalfcloseTimeout, b.HalfcloseTimeout) {
		return false
	}
	if !util.Ints64Match(a.Timeout, b.Timeout) {
		return false
	}
	return true
}
func matchProtocolTcp(a *ProtocolTcp, b *ProtocolTcp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolTcpOverride(a.Override, b.Override) {
		return false
	}
	if !util.StringsMatch(a.Port, b.Port) {
		return false
	}
	if !util.StringsMatch(a.SourcePort, b.SourcePort) {
		return false
	}
	return true
}
func matchProtocolUdpOverride(a *ProtocolUdpOverride, b *ProtocolUdpOverride) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Timeout, b.Timeout) {
		return false
	}
	return true
}
func matchProtocolUdp(a *ProtocolUdp, b *ProtocolUdp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolUdpOverride(a.Override, b.Override) {
		return false
	}
	if !util.StringsMatch(a.Port, b.Port) {
		return false
	}
	if !util.StringsMatch(a.SourcePort, b.SourcePort) {
		return false
	}
	return true
}
func matchProtocol(a *Protocol, b *Protocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolTcp(a.Tcp, b.Tcp) {
		return false
	}
	if !matchProtocolUdp(a.Udp, b.Udp) {
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
