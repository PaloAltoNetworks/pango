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
	Name        string
	Description *string
	Protocol    *Protocol
	Tags        []string

	Misc map[string][]generic.Xml
}

type Protocol struct {
	Tcp *ProtocolTcp
	Udp *ProtocolUdp
}
type ProtocolTcp struct {
	DestinationPort *int64
	Override        *ProtocolTcpOverride
	SourcePort      *int64
}
type ProtocolTcpOverride struct {
	HalfcloseTimeout *int64
	Timeout          *int64
	TimewaitTimeout  *int64
}
type ProtocolUdp struct {
	DestinationPort *int64
	Override        *ProtocolUdpOverride
	SourcePort      *int64
}
type ProtocolUdpOverride struct {
	No  *string
	Yes *ProtocolUdpOverrideYes
}
type ProtocolUdpOverrideYes struct {
	Timeout *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	Description *string          `xml:"description,omitempty"`
	Protocol    *ProtocolXml     `xml:"protocol,omitempty"`
	Tags        *util.MemberType `xml:"tag,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type ProtocolXml struct {
	Tcp *ProtocolTcpXml `xml:"tcp,omitempty"`
	Udp *ProtocolUdpXml `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolTcpXml struct {
	DestinationPort *int64                  `xml:"port,omitempty"`
	Override        *ProtocolTcpOverrideXml `xml:"override>yes,omitempty"`
	SourcePort      *int64                  `xml:"source-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolTcpOverrideXml struct {
	HalfcloseTimeout *int64 `xml:"halfclose-timeout,omitempty"`
	Timeout          *int64 `xml:"timeout,omitempty"`
	TimewaitTimeout  *int64 `xml:"timewait-timeout,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolUdpXml struct {
	DestinationPort *int64                  `xml:"port,omitempty"`
	Override        *ProtocolUdpOverrideXml `xml:"override,omitempty"`
	SourcePort      *int64                  `xml:"source-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolUdpOverrideXml struct {
	No  *string                    `xml:"no,omitempty"`
	Yes *ProtocolUdpOverrideYesXml `xml:"yes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolUdpOverrideYesXml struct {
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
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "tags" || v == "Tags" {
		return e.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return int64(len(e.Tags)), nil
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
			if o.Protocol.Tcp.DestinationPort != nil {
				nestedProtocol.Tcp.DestinationPort = o.Protocol.Tcp.DestinationPort
			}
			if o.Protocol.Tcp.SourcePort != nil {
				nestedProtocol.Tcp.SourcePort = o.Protocol.Tcp.SourcePort
			}
			if o.Protocol.Tcp.Override != nil {
				nestedProtocol.Tcp.Override = &ProtocolTcpOverrideXml{}
				if _, ok := o.Misc["ProtocolTcpOverride"]; ok {
					nestedProtocol.Tcp.Override.Misc = o.Misc["ProtocolTcpOverride"]
				}
				if o.Protocol.Tcp.Override.Timeout != nil {
					nestedProtocol.Tcp.Override.Timeout = o.Protocol.Tcp.Override.Timeout
				}
				if o.Protocol.Tcp.Override.HalfcloseTimeout != nil {
					nestedProtocol.Tcp.Override.HalfcloseTimeout = o.Protocol.Tcp.Override.HalfcloseTimeout
				}
				if o.Protocol.Tcp.Override.TimewaitTimeout != nil {
					nestedProtocol.Tcp.Override.TimewaitTimeout = o.Protocol.Tcp.Override.TimewaitTimeout
				}
			}
		}
		if o.Protocol.Udp != nil {
			nestedProtocol.Udp = &ProtocolUdpXml{}
			if _, ok := o.Misc["ProtocolUdp"]; ok {
				nestedProtocol.Udp.Misc = o.Misc["ProtocolUdp"]
			}
			if o.Protocol.Udp.DestinationPort != nil {
				nestedProtocol.Udp.DestinationPort = o.Protocol.Udp.DestinationPort
			}
			if o.Protocol.Udp.SourcePort != nil {
				nestedProtocol.Udp.SourcePort = o.Protocol.Udp.SourcePort
			}
			if o.Protocol.Udp.Override != nil {
				nestedProtocol.Udp.Override = &ProtocolUdpOverrideXml{}
				if _, ok := o.Misc["ProtocolUdpOverride"]; ok {
					nestedProtocol.Udp.Override.Misc = o.Misc["ProtocolUdpOverride"]
				}
				if o.Protocol.Udp.Override.Yes != nil {
					nestedProtocol.Udp.Override.Yes = &ProtocolUdpOverrideYesXml{}
					if _, ok := o.Misc["ProtocolUdpOverrideYes"]; ok {
						nestedProtocol.Udp.Override.Yes.Misc = o.Misc["ProtocolUdpOverrideYes"]
					}
					if o.Protocol.Udp.Override.Yes.Timeout != nil {
						nestedProtocol.Udp.Override.Yes.Timeout = o.Protocol.Udp.Override.Yes.Timeout
					}
				}
				if o.Protocol.Udp.Override.No != nil {
					nestedProtocol.Udp.Override.No = o.Protocol.Udp.Override.No
				}
			}
		}
	}
	entry.Protocol = nestedProtocol

	entry.Tags = util.StrToMem(o.Tags)

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
		var nestedProtocol *Protocol
		if o.Protocol != nil {
			nestedProtocol = &Protocol{}
			if o.Protocol.Misc != nil {
				entry.Misc["Protocol"] = o.Protocol.Misc
			}
			if o.Protocol.Udp != nil {
				nestedProtocol.Udp = &ProtocolUdp{}
				if o.Protocol.Udp.Misc != nil {
					entry.Misc["ProtocolUdp"] = o.Protocol.Udp.Misc
				}
				if o.Protocol.Udp.DestinationPort != nil {
					nestedProtocol.Udp.DestinationPort = o.Protocol.Udp.DestinationPort
				}
				if o.Protocol.Udp.SourcePort != nil {
					nestedProtocol.Udp.SourcePort = o.Protocol.Udp.SourcePort
				}
				if o.Protocol.Udp.Override != nil {
					nestedProtocol.Udp.Override = &ProtocolUdpOverride{}
					if o.Protocol.Udp.Override.Misc != nil {
						entry.Misc["ProtocolUdpOverride"] = o.Protocol.Udp.Override.Misc
					}
					if o.Protocol.Udp.Override.Yes != nil {
						nestedProtocol.Udp.Override.Yes = &ProtocolUdpOverrideYes{}
						if o.Protocol.Udp.Override.Yes.Misc != nil {
							entry.Misc["ProtocolUdpOverrideYes"] = o.Protocol.Udp.Override.Yes.Misc
						}
						if o.Protocol.Udp.Override.Yes.Timeout != nil {
							nestedProtocol.Udp.Override.Yes.Timeout = o.Protocol.Udp.Override.Yes.Timeout
						}
					}
					if o.Protocol.Udp.Override.No != nil {
						nestedProtocol.Udp.Override.No = o.Protocol.Udp.Override.No
					}
				}
			}
			if o.Protocol.Tcp != nil {
				nestedProtocol.Tcp = &ProtocolTcp{}
				if o.Protocol.Tcp.Misc != nil {
					entry.Misc["ProtocolTcp"] = o.Protocol.Tcp.Misc
				}
				if o.Protocol.Tcp.DestinationPort != nil {
					nestedProtocol.Tcp.DestinationPort = o.Protocol.Tcp.DestinationPort
				}
				if o.Protocol.Tcp.SourcePort != nil {
					nestedProtocol.Tcp.SourcePort = o.Protocol.Tcp.SourcePort
				}
				if o.Protocol.Tcp.Override != nil {
					nestedProtocol.Tcp.Override = &ProtocolTcpOverride{}
					if o.Protocol.Tcp.Override.Misc != nil {
						entry.Misc["ProtocolTcpOverride"] = o.Protocol.Tcp.Override.Misc
					}
					if o.Protocol.Tcp.Override.Timeout != nil {
						nestedProtocol.Tcp.Override.Timeout = o.Protocol.Tcp.Override.Timeout
					}
					if o.Protocol.Tcp.Override.HalfcloseTimeout != nil {
						nestedProtocol.Tcp.Override.HalfcloseTimeout = o.Protocol.Tcp.Override.HalfcloseTimeout
					}
					if o.Protocol.Tcp.Override.TimewaitTimeout != nil {
						nestedProtocol.Tcp.Override.TimewaitTimeout = o.Protocol.Tcp.Override.TimewaitTimeout
					}
				}
			}
		}
		entry.Protocol = nestedProtocol

		entry.Tags = util.MemToStr(o.Tags)

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
	if !matchProtocol(a.Protocol, b.Protocol) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
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
	if !util.Ints64Match(a.Timeout, b.Timeout) {
		return false
	}
	if !util.Ints64Match(a.HalfcloseTimeout, b.HalfcloseTimeout) {
		return false
	}
	if !util.Ints64Match(a.TimewaitTimeout, b.TimewaitTimeout) {
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
	if !util.Ints64Match(a.DestinationPort, b.DestinationPort) {
		return false
	}
	if !util.Ints64Match(a.SourcePort, b.SourcePort) {
		return false
	}
	if !matchProtocolTcpOverride(a.Override, b.Override) {
		return false
	}
	return true
}
func matchProtocolUdpOverrideYes(a *ProtocolUdpOverrideYes, b *ProtocolUdpOverrideYes) bool {
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
func matchProtocolUdpOverride(a *ProtocolUdpOverride, b *ProtocolUdpOverride) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolUdpOverrideYes(a.Yes, b.Yes) {
		return false
	}
	if !util.StringsMatch(a.No, b.No) {
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
	if !util.Ints64Match(a.DestinationPort, b.DestinationPort) {
		return false
	}
	if !util.Ints64Match(a.SourcePort, b.SourcePort) {
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
