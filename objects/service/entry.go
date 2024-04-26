package service

import (
	"encoding/xml"
	"fmt"
	"strings"

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
	Tags        []string // ordered
	Protocol    ProtocolObject

	Misc map[string][]generic.Xml
}

type ProtocolObject struct {
	Tcp  *TcpObject
	Udp  *UdpObject
	Sctp *SctpObject
}

type TcpObject struct {
	DestinationPort string
	SourcePort      *string
	Override        *TcpOverrideObject
}

type TcpOverrideObject struct {
	No  bool
	Yes *YesTcpOverrideObject
}

type YesTcpOverrideObject struct {
	Timeout           *int64
	HalfClosedTimeout *int64
	TimeWaitTimeout   *int64
}

type UdpObject struct {
	DestinationPort string
	SourcePort      *string
	Override        *UdpOverrideObject
}

type UdpOverrideObject struct {
	No  bool
	Yes *YesUdpOverrideObject
}

type YesUdpOverrideObject struct {
	Timeout *int64
}

type SctpObject struct {
	DestinationPort string
	SourcePort      *string
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

func (e *Entry) Field(value string) (any, error) {
	obj := e
	v := value

	if v == "name" || v == "Name" {
		return obj.Name, nil
	}
	if v == "description" || v == "Description" {
		return obj.Description, nil
	}
	if v == "tags" || v == "Tags" {
		return obj.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return int64(len(obj.Tags)), nil
	}
	if v == "protocol" || v == "Protocol" {
		return true, nil
	}
	if strings.HasPrefix(v, "protocol.") || strings.HasPrefix(v, "Protocol.") {
		obj := obj.Protocol
		for _, chk := range []string{"protocol.", "Protocol."} {
			if strings.HasPrefix(v, chk) {
				v = v[len(chk):]
				break
			}
		}

		if v == "tcp" || v == "Tcp" {
			return obj.Tcp != nil, nil
		}
		if strings.HasPrefix(v, "tcp.") || strings.HasPrefix(v, "Tcp.") {
			if obj.Tcp == nil {
				return nil, fmt.Errorf("protocol.tcp is nil")
			}
			obj := obj.Tcp
			for _, chk := range []string{"tcp.", "Tcp."} {
				if strings.HasPrefix(v, chk) {
					v = v[len(chk):]
					break
				}
			}

			if v == "destination_port" || v == "DestinationPort" {
				return obj.DestinationPort, nil
			}
			if v == "source_port" || v == "SourcePort" {
				return obj.SourcePort, nil
			}
			if v == "override" || v == "Override" {
				return obj.Override != nil, nil
			}
			if strings.HasPrefix(v, "override.") || strings.HasPrefix(v, "Override.") {
				if obj.Override == nil {
					return nil, fmt.Errorf("protocol.tcp.override is nil")
				}
				obj := obj.Override
				for _, chk := range []string{"override.", "Override."} {
					if strings.HasPrefix(v, chk) {
						v = v[len(chk):]
						break
					}
				}

				if v == "no" || v == "No" {
					return obj.No, nil
				}
				if v == "yes" || v == "Yes" {
					return obj.Yes != nil, nil
				}
				if strings.HasPrefix(v, "yes.") || strings.HasPrefix(v, "Yes.") {
					if obj.Yes == nil {
						return nil, fmt.Errorf("protocol.tcp.override.yes is nil")
					}
					obj := obj.Yes
					for _, chk := range []string{"yes.", "Yes."} {
						if strings.HasPrefix(v, chk) {
							v = v[len(chk):]
							break
						}
					}

					if v == "timeout" || v == "Timeout" {
						return obj.Timeout, nil
					}
					if v == "half_closed_timeout" || v == "HalfClosedTimeout" {
						return obj.HalfClosedTimeout, nil
					}
					if v == "time_wait_timeout" || v == "TimeWaitTimeout" {
						return obj.TimeWaitTimeout, nil
					}
				}
			}
		}
		if v == "udp" || v == "Udp" {
			return obj.Udp != nil, nil
		}
		if strings.HasPrefix(v, "udp.") || strings.HasPrefix(v, "Udp.") {
			if obj.Udp == nil {
				return nil, fmt.Errorf("protocol.udp is nil")
			}
			obj := obj.Udp
			for _, chk := range []string{"udp.", "Udp."} {
				if strings.HasPrefix(v, chk) {
					v = v[len(chk):]
					break
				}
			}

			if v == "destination_port" || v == "DestinationPort" {
				return obj.DestinationPort, nil
			}
			if v == "source_port" || v == "SourcePort" {
				return obj.SourcePort, nil
			}
			if v == "override" || v == "Override" {
				return obj.Override != nil, nil
			}
			if strings.HasPrefix(v, "override.") || strings.HasPrefix(v, "Override.") {
				if obj.Override == nil {
					return nil, fmt.Errorf("protocol.udp.override is nil")
				}
				obj := obj.Override
				for _, chk := range []string{"override.", "Override."} {
					if strings.HasPrefix(v, chk) {
						v = v[len(chk):]
						break
					}
				}

				if v == "no" || v == "No" {
					return obj.No, nil
				}
				if v == "yes" || v == "Yes" {
					return obj.Yes != nil, nil
				}
				if strings.HasPrefix(v, "yes.") || strings.HasPrefix(v, "Yes.") {
					if obj.Yes == nil {
						return nil, fmt.Errorf("protocol.udp.override.yes is nil")
					}
					obj := obj.Yes
					for _, chk := range []string{"yes.", "Yes."} {
						if strings.HasPrefix(v, chk) {
							v = v[len(v):]
							break
						}
					}

					if v == "timeout" || v == "Timeout" {
						return obj.Timeout, nil
					}
				}
			}
		}
		if v == "sctp" || v == "Sctp" {
			return obj.Sctp != nil, nil
		}
		if strings.HasPrefix(v, "sctp.") || strings.HasPrefix(v, "Sctp.") {
			if obj.Sctp == nil {
				return nil, fmt.Errorf("protocol.sctp is nil")
			}
			obj := obj.Sctp
			for _, chk := range []string{"sctp.", "Sctp."} {
				if strings.HasPrefix(v, chk) {
					v = v[len(chk):]
					break
				}
			}

			if v == "destination_port" || v == "DestinationPort" {
				return obj.DestinationPort, nil
			}
			if v == "source_port" || v == "SourcePort" {
				return obj.SourcePort, nil
			}
		}
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

	if o.Protocol.Tcp != nil {
		ans.Protocol.Tcp = &tcpProtocol1{}
		ans.Protocol.Tcp.DestinationPort = o.Protocol.Tcp.DestinationPort
		ans.Protocol.Tcp.SourcePort = o.Protocol.Tcp.SourcePort
		if o.Protocol.Tcp.Override != nil {
			ans.Protocol.Tcp.Override = &tcpProtocolOverride1{}
			if o.Protocol.Tcp.Override.No {
				ans.Protocol.Tcp.Override.No = util.String("")
			}
			if o.Protocol.Tcp.Override.Yes != nil {
				ans.Protocol.Tcp.Override.Yes = &tcpProtocolOverrideYes1{}
				ans.Protocol.Tcp.Override.Yes.Timeout = o.Protocol.Tcp.Override.Yes.Timeout
				ans.Protocol.Tcp.Override.Yes.HalfClosedTimeout = o.Protocol.Tcp.Override.Yes.HalfClosedTimeout
				ans.Protocol.Tcp.Override.Yes.TimeWaitTimeout = o.Protocol.Tcp.Override.Yes.TimeWaitTimeout

				ans.Protocol.Tcp.Override.Yes.Misc = o.Misc["tcpProtocolOverrideYes"]
			}

			ans.Protocol.Tcp.Override.Misc = o.Misc["tcpProtocolOverride"]
		}

		ans.Protocol.Tcp.Misc = o.Misc["tcpProtocol"]
	}
	if o.Protocol.Udp != nil {
		ans.Protocol.Udp = &udpProtocol1{}
		ans.Protocol.Udp.DestinationPort = o.Protocol.Udp.DestinationPort
		ans.Protocol.Udp.SourcePort = o.Protocol.Udp.SourcePort
		if o.Protocol.Udp.Override != nil {
			ans.Protocol.Udp.Override = &udpProtocolOverride1{}
			if o.Protocol.Udp.Override.No {
				ans.Protocol.Udp.Override.No = util.String("")
			}
			if o.Protocol.Udp.Override.Yes != nil {
				ans.Protocol.Udp.Override.Yes = &udpProtocolOverrideYes1{}
				ans.Protocol.Udp.Override.Yes.Timeout = o.Protocol.Udp.Override.Yes.Timeout

				ans.Protocol.Udp.Override.Yes.Misc = o.Misc["udpProtocolOverrideYes"]
			}

			ans.Protocol.Udp.Override.Misc = o.Misc["udpProtocolOverride"]
		}

		ans.Protocol.Udp.Misc = o.Misc["udpProtocol"]
	}
	if o.Protocol.Sctp != nil {
		ans.Protocol.Sctp = &sctpProtocol1{}
		ans.Protocol.Sctp.DestinationPort = o.Protocol.Sctp.DestinationPort
		ans.Protocol.Sctp.SourcePort = o.Protocol.Sctp.SourcePort

		ans.Protocol.Sctp.Misc = o.Misc["sctpProtocol"]
	}

	ans.Misc = o.Misc["entry"]

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
		var1.Tags = util.MemToStr(var0.Tags)
		if var0.Protocol.Tcp != nil {
			var1.Protocol.Tcp = &TcpObject{}
			var1.Protocol.Tcp.DestinationPort = var0.Protocol.Tcp.DestinationPort
			var1.Protocol.Tcp.SourcePort = var0.Protocol.Tcp.SourcePort
			if var0.Protocol.Tcp.Override != nil {
				var1.Protocol.Tcp.Override = &TcpOverrideObject{}
				if var0.Protocol.Tcp.Override.No != nil {
					var1.Protocol.Tcp.Override.No = true
				}
				if var0.Protocol.Tcp.Override.Yes != nil {
					var1.Protocol.Tcp.Override.Yes = &YesTcpOverrideObject{}
					var1.Protocol.Tcp.Override.Yes.Timeout = var0.Protocol.Tcp.Override.Yes.Timeout
					var1.Protocol.Tcp.Override.Yes.HalfClosedTimeout = var0.Protocol.Tcp.Override.Yes.HalfClosedTimeout
					var1.Protocol.Tcp.Override.Yes.TimeWaitTimeout = var0.Protocol.Tcp.Override.Yes.TimeWaitTimeout

					var1.Misc["tcpProtocolOverrideYes"] = var0.Protocol.Tcp.Override.Yes.Misc
				}

				var1.Misc["tcpProtocolOverride"] = var0.Protocol.Tcp.Override.Misc
			}

			var1.Misc["tcpProtocol"] = var0.Protocol.Tcp.Misc
		}
		if var0.Protocol.Udp != nil {
			var1.Protocol.Udp = &UdpObject{}
			var1.Protocol.Udp.DestinationPort = var0.Protocol.Udp.DestinationPort
			var1.Protocol.Udp.SourcePort = var0.Protocol.Udp.SourcePort
			if var0.Protocol.Udp.Override != nil {
				var1.Protocol.Udp.Override = &UdpOverrideObject{}
				if var0.Protocol.Udp.Override.No != nil {
					var1.Protocol.Udp.Override.No = true
				}
				if var0.Protocol.Udp.Override.Yes != nil {
					var1.Protocol.Udp.Override.Yes = &YesUdpOverrideObject{}
					var1.Protocol.Udp.Override.Yes.Timeout = var0.Protocol.Udp.Override.Yes.Timeout

					var1.Misc["udpProtocolOverrideYes"] = var0.Protocol.Udp.Override.Yes.Misc
				}

				var1.Misc["udpProtocolOverride"] = var0.Protocol.Udp.Override.Misc
			}

			var1.Misc["udpProtocol"] = var0.Protocol.Udp.Misc
		}
		if var0.Protocol.Sctp != nil {
			var1.Protocol.Sctp = &SctpObject{}
			var1.Protocol.Sctp.DestinationPort = var0.Protocol.Sctp.DestinationPort
			var1.Protocol.Sctp.SourcePort = var0.Protocol.Sctp.SourcePort

			var1.Misc["sctpProtocol"] = var0.Protocol.Sctp.Misc
		}

		var1.Misc["entry"] = var0.Misc

		ans = append(ans, var1)
	}

	return ans, nil
}

type Entry1 struct {
	XMLName     xml.Name         `xml:"entry"`
	Name        string           `xml:"name,attr"`
	Description *string          `xml:"description,omitempty"`
	Tags        *util.MemberType `xml:"tag"`
	Protocol    protocol1        `xml:"protocol"`

	Misc []generic.Xml `xml:",any"`
}

type protocol1 struct {
	Tcp  *tcpProtocol1  `xml:"tcp"`
	Udp  *udpProtocol1  `xml:"udp"`
	Sctp *sctpProtocol1 `xml:"sctp"`

	Misc []generic.Xml
}

type tcpProtocol1 struct {
	DestinationPort string                `xml:"port"`
	SourcePort      *string               `xml:"source-port"`
	Override        *tcpProtocolOverride1 `xml:"override"`

	Misc []generic.Xml
}

type tcpProtocolOverride1 struct {
	No  *string                  `xml:"no,omitempty"`
	Yes *tcpProtocolOverrideYes1 `xml:"yes"`

	Misc []generic.Xml
}

type tcpProtocolOverrideYes1 struct {
	Timeout           *int64 `xml:"timeout"`
	HalfClosedTimeout *int64 `xml:"halfclose-timeout"`
	TimeWaitTimeout   *int64 `xml:"timewait-timeout"`

	Misc []generic.Xml
}

type udpProtocol1 struct {
	DestinationPort string                `xml:"port"`
	SourcePort      *string               `xml:"source-port"`
	Override        *udpProtocolOverride1 `xml:"override"`

	Misc []generic.Xml
}

type udpProtocolOverride1 struct {
	No  *string                  `xml:"no,omitempty"`
	Yes *udpProtocolOverrideYes1 `xml:"yes"`

	Misc []generic.Xml
}

type udpProtocolOverrideYes1 struct {
	Timeout *int64 `xml:"timeout"`

	Misc []generic.Xml
}

type sctpProtocol1 struct {
	DestinationPort string  `xml:"port"`
	SourcePort      *string `xml:"source-port"`

	Misc []generic.Xml
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

	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}

	if a.Protocol.Tcp == nil && b.Protocol.Tcp == nil {
	} else if a.Protocol.Tcp == nil || b.Protocol.Tcp == nil {
		return false
	} else {
		if a.Protocol.Tcp.DestinationPort != b.Protocol.Tcp.DestinationPort {
			return false
		}
		if !util.StringsMatch(a.Protocol.Tcp.SourcePort, b.Protocol.Tcp.SourcePort) {
			return false
		}
		if a.Protocol.Tcp.Override == nil && b.Protocol.Tcp.Override == nil {
		} else if a.Protocol.Tcp.Override == nil || b.Protocol.Tcp.Override == nil {
			return false
		} else {
			if a.Protocol.Tcp.Override.No != b.Protocol.Tcp.Override.No {
				return false
			}
			if a.Protocol.Tcp.Override.Yes == nil && b.Protocol.Tcp.Override.Yes == nil {
			} else if a.Protocol.Tcp.Override.Yes == nil || b.Protocol.Tcp.Override.Yes == nil {
				return false
			} else {
				if !util.IntsMatch(a.Protocol.Tcp.Override.Yes.Timeout, b.Protocol.Tcp.Override.Yes.Timeout) {
					return false
				}
				if !util.IntsMatch(a.Protocol.Tcp.Override.Yes.HalfClosedTimeout, b.Protocol.Tcp.Override.Yes.HalfClosedTimeout) {
					return false
				}
				if !util.IntsMatch(a.Protocol.Tcp.Override.Yes.TimeWaitTimeout, b.Protocol.Tcp.Override.Yes.TimeWaitTimeout) {
					return false
				}
			}
		}
	}

	if a.Protocol.Udp == nil && b.Protocol.Udp == nil {
	} else if a.Protocol.Udp == nil || b.Protocol.Udp == nil {
		return false
	} else {
		if a.Protocol.Udp.DestinationPort != b.Protocol.Udp.DestinationPort {
			return false
		}
		if !util.StringsMatch(a.Protocol.Udp.SourcePort, b.Protocol.Udp.SourcePort) {
			return false
		}
		if a.Protocol.Udp.Override == nil && b.Protocol.Udp.Override == nil {
		} else if a.Protocol.Udp.Override == nil || b.Protocol.Udp.Override == nil {
			return false
		} else {
			if a.Protocol.Udp.Override.No != b.Protocol.Udp.Override.No {
				return false
			}
			if a.Protocol.Udp.Override.Yes == nil && b.Protocol.Udp.Override.Yes == nil {
			} else if a.Protocol.Udp.Override.Yes == nil || b.Protocol.Udp.Override.Yes == nil {
				return false
			} else {
				if !util.IntsMatch(a.Protocol.Tcp.Override.Yes.Timeout, b.Protocol.Tcp.Override.Yes.Timeout) {
					return false
				}
			}
		}
	}

	if a.Protocol.Sctp == nil && b.Protocol.Sctp == nil {
	} else if a.Protocol.Sctp == nil || b.Protocol.Sctp == nil {
		return false
	} else {
		if a.Protocol.Sctp.DestinationPort != b.Protocol.Sctp.DestinationPort {
			return false
		}
		if !util.StringsMatch(a.Protocol.Sctp.SourcePort, b.Protocol.Sctp.SourcePort) {
			return false
		}
	}

	return true
}
