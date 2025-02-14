package nat

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
	Suffix = []string{}
)

type Entry struct {
	Name                          string
	ActiveActiveDeviceBinding     *string
	Description                   *string
	Destination                   []string
	Disabled                      *bool
	From                          []string
	GroupTag                      *string
	NatType                       *string
	Service                       *string
	Source                        []string
	SourceTranslation             *SourceTranslation
	Tag                           []string
	Target                        *Target
	To                            []string
	ToInterface                   *string
	Uuid                          *string
	DestinationTranslation        *DestinationTranslation
	DynamicDestinationTranslation *DynamicDestinationTranslation

	Misc map[string][]generic.Xml
}

type DestinationTranslation struct {
	DnsRewrite        *DestinationTranslationDnsRewrite
	TranslatedAddress *string
	TranslatedPort    *int64
}
type DestinationTranslationDnsRewrite struct {
	Direction *string
}
type DynamicDestinationTranslation struct {
	Distribution      *string
	TranslatedAddress *string
	TranslatedPort    *int64
}
type SourceTranslation struct {
	DynamicIp        *SourceTranslationDynamicIp
	DynamicIpAndPort *SourceTranslationDynamicIpAndPort
	StaticIp         *SourceTranslationStaticIp
}
type SourceTranslationDynamicIp struct {
	Fallback          *SourceTranslationDynamicIpFallback
	TranslatedAddress []string
}
type SourceTranslationDynamicIpAndPort struct {
	InterfaceAddress  *SourceTranslationDynamicIpAndPortInterfaceAddress
	TranslatedAddress []string
}
type SourceTranslationDynamicIpAndPortInterfaceAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
}
type SourceTranslationDynamicIpFallback struct {
	InterfaceAddress  *SourceTranslationDynamicIpFallbackInterfaceAddress
	TranslatedAddress []string
}
type SourceTranslationDynamicIpFallbackInterfaceAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
}
type SourceTranslationStaticIp struct {
	BiDirectional     *string
	TranslatedAddress *string
}
type Target struct {
	Devices []TargetDevices
	Negate  *bool
	Tags    []string
}
type TargetDevices struct {
	Name string
	Vsys []TargetDevicesVsys
}
type TargetDevicesVsys struct {
	Name string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                       xml.Name                          `xml:"entry"`
	Name                          string                            `xml:"name,attr"`
	ActiveActiveDeviceBinding     *string                           `xml:"active-active-device-binding,omitempty"`
	Description                   *string                           `xml:"description,omitempty"`
	Destination                   *util.MemberType                  `xml:"destination,omitempty"`
	Disabled                      *string                           `xml:"disabled,omitempty"`
	From                          *util.MemberType                  `xml:"from,omitempty"`
	GroupTag                      *string                           `xml:"group-tag,omitempty"`
	NatType                       *string                           `xml:"nat-type,omitempty"`
	Service                       *string                           `xml:"service,omitempty"`
	Source                        *util.MemberType                  `xml:"source,omitempty"`
	SourceTranslation             *SourceTranslationXml             `xml:"source-translation,omitempty"`
	Tag                           *util.MemberType                  `xml:"tag,omitempty"`
	Target                        *TargetXml                        `xml:"target,omitempty"`
	To                            *util.MemberType                  `xml:"to,omitempty"`
	ToInterface                   *string                           `xml:"to-interface,omitempty"`
	Uuid                          *string                           `xml:"uuid,attr,omitempty"`
	DestinationTranslation        *DestinationTranslationXml        `xml:"destination-translation,omitempty"`
	DynamicDestinationTranslation *DynamicDestinationTranslationXml `xml:"dynamic-destination-translation,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DestinationTranslationXml struct {
	DnsRewrite        *DestinationTranslationDnsRewriteXml `xml:"dns-rewrite,omitempty"`
	TranslatedAddress *string                              `xml:"translated-address,omitempty"`
	TranslatedPort    *int64                               `xml:"translated-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DestinationTranslationDnsRewriteXml struct {
	Direction *string `xml:"direction,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DynamicDestinationTranslationXml struct {
	Distribution      *string `xml:"distribution,omitempty"`
	TranslatedAddress *string `xml:"translated-address,omitempty"`
	TranslatedPort    *int64  `xml:"translated-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationXml struct {
	DynamicIp        *SourceTranslationDynamicIpXml        `xml:"dynamic-ip,omitempty"`
	DynamicIpAndPort *SourceTranslationDynamicIpAndPortXml `xml:"dynamic-ip-and-port,omitempty"`
	StaticIp         *SourceTranslationStaticIpXml         `xml:"static-ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationDynamicIpXml struct {
	Fallback          *SourceTranslationDynamicIpFallbackXml `xml:"fallback,omitempty"`
	TranslatedAddress *util.MemberType                       `xml:"translated-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationDynamicIpAndPortXml struct {
	InterfaceAddress  *SourceTranslationDynamicIpAndPortInterfaceAddressXml `xml:"interface-address,omitempty"`
	TranslatedAddress *util.MemberType                                      `xml:"translated-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationDynamicIpAndPortInterfaceAddressXml struct {
	Interface  *string `xml:"interface,omitempty"`
	FloatingIp *string `xml:"floating-ip,omitempty"`
	Ip         *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationDynamicIpFallbackXml struct {
	InterfaceAddress  *SourceTranslationDynamicIpFallbackInterfaceAddressXml `xml:"interface-address,omitempty"`
	TranslatedAddress *util.MemberType                                       `xml:"translated-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationDynamicIpFallbackInterfaceAddressXml struct {
	Interface  *string `xml:"interface,omitempty"`
	FloatingIp *string `xml:"floating-ip,omitempty"`
	Ip         *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SourceTranslationStaticIpXml struct {
	BiDirectional     *string `xml:"bi-directional,omitempty"`
	TranslatedAddress *string `xml:"translated-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TargetXml struct {
	Devices []TargetDevicesXml `xml:"devices>entry,omitempty"`
	Negate  *string            `xml:"negate,omitempty"`
	Tags    *util.MemberType   `xml:"tags,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TargetDevicesXml struct {
	XMLName xml.Name               `xml:"entry"`
	Name    string                 `xml:"name,attr"`
	Vsys    []TargetDevicesVsysXml `xml:"vsys>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TargetDevicesVsysXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "active_active_device_binding" || v == "ActiveActiveDeviceBinding" {
		return e.ActiveActiveDeviceBinding, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "destination" || v == "Destination" {
		return e.Destination, nil
	}
	if v == "destination|LENGTH" || v == "Destination|LENGTH" {
		return int64(len(e.Destination)), nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "from" || v == "From" {
		return e.From, nil
	}
	if v == "from|LENGTH" || v == "From|LENGTH" {
		return int64(len(e.From)), nil
	}
	if v == "group_tag" || v == "GroupTag" {
		return e.GroupTag, nil
	}
	if v == "nat_type" || v == "NatType" {
		return e.NatType, nil
	}
	if v == "service" || v == "Service" {
		return e.Service, nil
	}
	if v == "source" || v == "Source" {
		return e.Source, nil
	}
	if v == "source|LENGTH" || v == "Source|LENGTH" {
		return int64(len(e.Source)), nil
	}
	if v == "source_translation" || v == "SourceTranslation" {
		return e.SourceTranslation, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
	}
	if v == "target" || v == "Target" {
		return e.Target, nil
	}
	if v == "to" || v == "To" {
		return e.To, nil
	}
	if v == "to|LENGTH" || v == "To|LENGTH" {
		return int64(len(e.To)), nil
	}
	if v == "to_interface" || v == "ToInterface" {
		return e.ToInterface, nil
	}
	if v == "uuid" || v == "Uuid" {
		return e.Uuid, nil
	}
	if v == "destination_translation" || v == "DestinationTranslation" {
		return e.DestinationTranslation, nil
	}
	if v == "dynamic_destination_translation" || v == "DynamicDestinationTranslation" {
		return e.DynamicDestinationTranslation, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.ActiveActiveDeviceBinding = o.ActiveActiveDeviceBinding
	entry.Description = o.Description
	entry.Destination = util.StrToMem(o.Destination)
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.From = util.StrToMem(o.From)
	entry.GroupTag = o.GroupTag
	entry.NatType = o.NatType
	entry.Service = o.Service
	entry.Source = util.StrToMem(o.Source)
	var nestedSourceTranslation *SourceTranslationXml
	if o.SourceTranslation != nil {
		nestedSourceTranslation = &SourceTranslationXml{}
		if _, ok := o.Misc["SourceTranslation"]; ok {
			nestedSourceTranslation.Misc = o.Misc["SourceTranslation"]
		}
		if o.SourceTranslation.DynamicIp != nil {
			nestedSourceTranslation.DynamicIp = &SourceTranslationDynamicIpXml{}
			if _, ok := o.Misc["SourceTranslationDynamicIp"]; ok {
				nestedSourceTranslation.DynamicIp.Misc = o.Misc["SourceTranslationDynamicIp"]
			}
			if o.SourceTranslation.DynamicIp.Fallback != nil {
				nestedSourceTranslation.DynamicIp.Fallback = &SourceTranslationDynamicIpFallbackXml{}
				if _, ok := o.Misc["SourceTranslationDynamicIpFallback"]; ok {
					nestedSourceTranslation.DynamicIp.Fallback.Misc = o.Misc["SourceTranslationDynamicIpFallback"]
				}
				if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress != nil {
					nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress = &SourceTranslationDynamicIpFallbackInterfaceAddressXml{}
					if _, ok := o.Misc["SourceTranslationDynamicIpFallbackInterfaceAddress"]; ok {
						nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.Misc = o.Misc["SourceTranslationDynamicIpFallbackInterfaceAddress"]
					}
					if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface != nil {
						nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface
					}
					if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp != nil {
						nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp
					}
					if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip != nil {
						nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip
					}
				}
				if o.SourceTranslation.DynamicIp.Fallback.TranslatedAddress != nil {
					nestedSourceTranslation.DynamicIp.Fallback.TranslatedAddress = util.StrToMem(o.SourceTranslation.DynamicIp.Fallback.TranslatedAddress)
				}
			}
			if o.SourceTranslation.DynamicIp.TranslatedAddress != nil {
				nestedSourceTranslation.DynamicIp.TranslatedAddress = util.StrToMem(o.SourceTranslation.DynamicIp.TranslatedAddress)
			}
		}
		if o.SourceTranslation.DynamicIpAndPort != nil {
			nestedSourceTranslation.DynamicIpAndPort = &SourceTranslationDynamicIpAndPortXml{}
			if _, ok := o.Misc["SourceTranslationDynamicIpAndPort"]; ok {
				nestedSourceTranslation.DynamicIpAndPort.Misc = o.Misc["SourceTranslationDynamicIpAndPort"]
			}
			if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress != nil {
				nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress = &SourceTranslationDynamicIpAndPortInterfaceAddressXml{}
				if _, ok := o.Misc["SourceTranslationDynamicIpAndPortInterfaceAddress"]; ok {
					nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.Misc = o.Misc["SourceTranslationDynamicIpAndPortInterfaceAddress"]
				}
				if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface != nil {
					nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface
				}
				if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip != nil {
					nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip
				}
				if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp != nil {
					nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp
				}
			}
			if o.SourceTranslation.DynamicIpAndPort.TranslatedAddress != nil {
				nestedSourceTranslation.DynamicIpAndPort.TranslatedAddress = util.StrToMem(o.SourceTranslation.DynamicIpAndPort.TranslatedAddress)
			}
		}
		if o.SourceTranslation.StaticIp != nil {
			nestedSourceTranslation.StaticIp = &SourceTranslationStaticIpXml{}
			if _, ok := o.Misc["SourceTranslationStaticIp"]; ok {
				nestedSourceTranslation.StaticIp.Misc = o.Misc["SourceTranslationStaticIp"]
			}
			if o.SourceTranslation.StaticIp.BiDirectional != nil {
				nestedSourceTranslation.StaticIp.BiDirectional = o.SourceTranslation.StaticIp.BiDirectional
			}
			if o.SourceTranslation.StaticIp.TranslatedAddress != nil {
				nestedSourceTranslation.StaticIp.TranslatedAddress = o.SourceTranslation.StaticIp.TranslatedAddress
			}
		}
	}
	entry.SourceTranslation = nestedSourceTranslation

	entry.Tag = util.StrToMem(o.Tag)
	var nestedTarget *TargetXml
	if o.Target != nil {
		nestedTarget = &TargetXml{}
		if _, ok := o.Misc["Target"]; ok {
			nestedTarget.Misc = o.Misc["Target"]
		}
		if o.Target.Devices != nil {
			nestedTarget.Devices = []TargetDevicesXml{}
			for _, oTargetDevices := range o.Target.Devices {
				nestedTargetDevices := TargetDevicesXml{}
				if _, ok := o.Misc["TargetDevices"]; ok {
					nestedTargetDevices.Misc = o.Misc["TargetDevices"]
				}
				if oTargetDevices.Vsys != nil {
					nestedTargetDevices.Vsys = []TargetDevicesVsysXml{}
					for _, oTargetDevicesVsys := range oTargetDevices.Vsys {
						nestedTargetDevicesVsys := TargetDevicesVsysXml{}
						if _, ok := o.Misc["TargetDevicesVsys"]; ok {
							nestedTargetDevicesVsys.Misc = o.Misc["TargetDevicesVsys"]
						}
						if oTargetDevicesVsys.Name != "" {
							nestedTargetDevicesVsys.Name = oTargetDevicesVsys.Name
						}
						nestedTargetDevices.Vsys = append(nestedTargetDevices.Vsys, nestedTargetDevicesVsys)
					}
				}
				if oTargetDevices.Name != "" {
					nestedTargetDevices.Name = oTargetDevices.Name
				}
				nestedTarget.Devices = append(nestedTarget.Devices, nestedTargetDevices)
			}
		}
		if o.Target.Negate != nil {
			nestedTarget.Negate = util.YesNo(o.Target.Negate, nil)
		}
		if o.Target.Tags != nil {
			nestedTarget.Tags = util.StrToMem(o.Target.Tags)
		}
	}
	entry.Target = nestedTarget

	entry.To = util.StrToMem(o.To)
	entry.ToInterface = o.ToInterface
	entry.Uuid = o.Uuid
	var nestedDestinationTranslation *DestinationTranslationXml
	if o.DestinationTranslation != nil {
		nestedDestinationTranslation = &DestinationTranslationXml{}
		if _, ok := o.Misc["DestinationTranslation"]; ok {
			nestedDestinationTranslation.Misc = o.Misc["DestinationTranslation"]
		}
		if o.DestinationTranslation.TranslatedPort != nil {
			nestedDestinationTranslation.TranslatedPort = o.DestinationTranslation.TranslatedPort
		}
		if o.DestinationTranslation.DnsRewrite != nil {
			nestedDestinationTranslation.DnsRewrite = &DestinationTranslationDnsRewriteXml{}
			if _, ok := o.Misc["DestinationTranslationDnsRewrite"]; ok {
				nestedDestinationTranslation.DnsRewrite.Misc = o.Misc["DestinationTranslationDnsRewrite"]
			}
			if o.DestinationTranslation.DnsRewrite.Direction != nil {
				nestedDestinationTranslation.DnsRewrite.Direction = o.DestinationTranslation.DnsRewrite.Direction
			}
		}
		if o.DestinationTranslation.TranslatedAddress != nil {
			nestedDestinationTranslation.TranslatedAddress = o.DestinationTranslation.TranslatedAddress
		}
	}
	entry.DestinationTranslation = nestedDestinationTranslation

	var nestedDynamicDestinationTranslation *DynamicDestinationTranslationXml
	if o.DynamicDestinationTranslation != nil {
		nestedDynamicDestinationTranslation = &DynamicDestinationTranslationXml{}
		if _, ok := o.Misc["DynamicDestinationTranslation"]; ok {
			nestedDynamicDestinationTranslation.Misc = o.Misc["DynamicDestinationTranslation"]
		}
		if o.DynamicDestinationTranslation.TranslatedAddress != nil {
			nestedDynamicDestinationTranslation.TranslatedAddress = o.DynamicDestinationTranslation.TranslatedAddress
		}
		if o.DynamicDestinationTranslation.TranslatedPort != nil {
			nestedDynamicDestinationTranslation.TranslatedPort = o.DynamicDestinationTranslation.TranslatedPort
		}
		if o.DynamicDestinationTranslation.Distribution != nil {
			nestedDynamicDestinationTranslation.Distribution = o.DynamicDestinationTranslation.Distribution
		}
	}
	entry.DynamicDestinationTranslation = nestedDynamicDestinationTranslation

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
		entry.ActiveActiveDeviceBinding = o.ActiveActiveDeviceBinding
		entry.Description = o.Description
		entry.Destination = util.MemToStr(o.Destination)
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.From = util.MemToStr(o.From)
		entry.GroupTag = o.GroupTag
		entry.NatType = o.NatType
		entry.Service = o.Service
		entry.Source = util.MemToStr(o.Source)
		var nestedSourceTranslation *SourceTranslation
		if o.SourceTranslation != nil {
			nestedSourceTranslation = &SourceTranslation{}
			if o.SourceTranslation.Misc != nil {
				entry.Misc["SourceTranslation"] = o.SourceTranslation.Misc
			}
			if o.SourceTranslation.DynamicIp != nil {
				nestedSourceTranslation.DynamicIp = &SourceTranslationDynamicIp{}
				if o.SourceTranslation.DynamicIp.Misc != nil {
					entry.Misc["SourceTranslationDynamicIp"] = o.SourceTranslation.DynamicIp.Misc
				}
				if o.SourceTranslation.DynamicIp.Fallback != nil {
					nestedSourceTranslation.DynamicIp.Fallback = &SourceTranslationDynamicIpFallback{}
					if o.SourceTranslation.DynamicIp.Fallback.Misc != nil {
						entry.Misc["SourceTranslationDynamicIpFallback"] = o.SourceTranslation.DynamicIp.Fallback.Misc
					}
					if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress != nil {
						nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress = &SourceTranslationDynamicIpFallbackInterfaceAddress{}
						if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Misc != nil {
							entry.Misc["SourceTranslationDynamicIpFallbackInterfaceAddress"] = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Misc
						}
						if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface != nil {
							nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Interface
						}
						if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp != nil {
							nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.FloatingIp
						}
						if o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip != nil {
							nestedSourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip = o.SourceTranslation.DynamicIp.Fallback.InterfaceAddress.Ip
						}
					}
					if o.SourceTranslation.DynamicIp.Fallback.TranslatedAddress != nil {
						nestedSourceTranslation.DynamicIp.Fallback.TranslatedAddress = util.MemToStr(o.SourceTranslation.DynamicIp.Fallback.TranslatedAddress)
					}
				}
				if o.SourceTranslation.DynamicIp.TranslatedAddress != nil {
					nestedSourceTranslation.DynamicIp.TranslatedAddress = util.MemToStr(o.SourceTranslation.DynamicIp.TranslatedAddress)
				}
			}
			if o.SourceTranslation.DynamicIpAndPort != nil {
				nestedSourceTranslation.DynamicIpAndPort = &SourceTranslationDynamicIpAndPort{}
				if o.SourceTranslation.DynamicIpAndPort.Misc != nil {
					entry.Misc["SourceTranslationDynamicIpAndPort"] = o.SourceTranslation.DynamicIpAndPort.Misc
				}
				if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress != nil {
					nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress = &SourceTranslationDynamicIpAndPortInterfaceAddress{}
					if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Misc != nil {
						entry.Misc["SourceTranslationDynamicIpAndPortInterfaceAddress"] = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Misc
					}
					if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface != nil {
						nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Interface
					}
					if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp != nil {
						nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.FloatingIp
					}
					if o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip != nil {
						nestedSourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip = o.SourceTranslation.DynamicIpAndPort.InterfaceAddress.Ip
					}
				}
				if o.SourceTranslation.DynamicIpAndPort.TranslatedAddress != nil {
					nestedSourceTranslation.DynamicIpAndPort.TranslatedAddress = util.MemToStr(o.SourceTranslation.DynamicIpAndPort.TranslatedAddress)
				}
			}
			if o.SourceTranslation.StaticIp != nil {
				nestedSourceTranslation.StaticIp = &SourceTranslationStaticIp{}
				if o.SourceTranslation.StaticIp.Misc != nil {
					entry.Misc["SourceTranslationStaticIp"] = o.SourceTranslation.StaticIp.Misc
				}
				if o.SourceTranslation.StaticIp.BiDirectional != nil {
					nestedSourceTranslation.StaticIp.BiDirectional = o.SourceTranslation.StaticIp.BiDirectional
				}
				if o.SourceTranslation.StaticIp.TranslatedAddress != nil {
					nestedSourceTranslation.StaticIp.TranslatedAddress = o.SourceTranslation.StaticIp.TranslatedAddress
				}
			}
		}
		entry.SourceTranslation = nestedSourceTranslation

		entry.Tag = util.MemToStr(o.Tag)
		var nestedTarget *Target
		if o.Target != nil {
			nestedTarget = &Target{}
			if o.Target.Misc != nil {
				entry.Misc["Target"] = o.Target.Misc
			}
			if o.Target.Devices != nil {
				nestedTarget.Devices = []TargetDevices{}
				for _, oTargetDevices := range o.Target.Devices {
					nestedTargetDevices := TargetDevices{}
					if oTargetDevices.Misc != nil {
						entry.Misc["TargetDevices"] = oTargetDevices.Misc
					}
					if oTargetDevices.Name != "" {
						nestedTargetDevices.Name = oTargetDevices.Name
					}
					if oTargetDevices.Vsys != nil {
						nestedTargetDevices.Vsys = []TargetDevicesVsys{}
						for _, oTargetDevicesVsys := range oTargetDevices.Vsys {
							nestedTargetDevicesVsys := TargetDevicesVsys{}
							if oTargetDevicesVsys.Misc != nil {
								entry.Misc["TargetDevicesVsys"] = oTargetDevicesVsys.Misc
							}
							if oTargetDevicesVsys.Name != "" {
								nestedTargetDevicesVsys.Name = oTargetDevicesVsys.Name
							}
							nestedTargetDevices.Vsys = append(nestedTargetDevices.Vsys, nestedTargetDevicesVsys)
						}
					}
					nestedTarget.Devices = append(nestedTarget.Devices, nestedTargetDevices)
				}
			}
			if o.Target.Negate != nil {
				nestedTarget.Negate = util.AsBool(o.Target.Negate, nil)
			}
			if o.Target.Tags != nil {
				nestedTarget.Tags = util.MemToStr(o.Target.Tags)
			}
		}
		entry.Target = nestedTarget

		entry.To = util.MemToStr(o.To)
		entry.ToInterface = o.ToInterface
		entry.Uuid = o.Uuid
		var nestedDestinationTranslation *DestinationTranslation
		if o.DestinationTranslation != nil {
			nestedDestinationTranslation = &DestinationTranslation{}
			if o.DestinationTranslation.Misc != nil {
				entry.Misc["DestinationTranslation"] = o.DestinationTranslation.Misc
			}
			if o.DestinationTranslation.TranslatedPort != nil {
				nestedDestinationTranslation.TranslatedPort = o.DestinationTranslation.TranslatedPort
			}
			if o.DestinationTranslation.DnsRewrite != nil {
				nestedDestinationTranslation.DnsRewrite = &DestinationTranslationDnsRewrite{}
				if o.DestinationTranslation.DnsRewrite.Misc != nil {
					entry.Misc["DestinationTranslationDnsRewrite"] = o.DestinationTranslation.DnsRewrite.Misc
				}
				if o.DestinationTranslation.DnsRewrite.Direction != nil {
					nestedDestinationTranslation.DnsRewrite.Direction = o.DestinationTranslation.DnsRewrite.Direction
				}
			}
			if o.DestinationTranslation.TranslatedAddress != nil {
				nestedDestinationTranslation.TranslatedAddress = o.DestinationTranslation.TranslatedAddress
			}
		}
		entry.DestinationTranslation = nestedDestinationTranslation

		var nestedDynamicDestinationTranslation *DynamicDestinationTranslation
		if o.DynamicDestinationTranslation != nil {
			nestedDynamicDestinationTranslation = &DynamicDestinationTranslation{}
			if o.DynamicDestinationTranslation.Misc != nil {
				entry.Misc["DynamicDestinationTranslation"] = o.DynamicDestinationTranslation.Misc
			}
			if o.DynamicDestinationTranslation.Distribution != nil {
				nestedDynamicDestinationTranslation.Distribution = o.DynamicDestinationTranslation.Distribution
			}
			if o.DynamicDestinationTranslation.TranslatedAddress != nil {
				nestedDynamicDestinationTranslation.TranslatedAddress = o.DynamicDestinationTranslation.TranslatedAddress
			}
			if o.DynamicDestinationTranslation.TranslatedPort != nil {
				nestedDynamicDestinationTranslation.TranslatedPort = o.DynamicDestinationTranslation.TranslatedPort
			}
		}
		entry.DynamicDestinationTranslation = nestedDynamicDestinationTranslation

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
	if !util.StringsMatch(a.ActiveActiveDeviceBinding, b.ActiveActiveDeviceBinding) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.Destination, b.Destination) {
		return false
	}
	if !util.BoolsMatch(a.Disabled, b.Disabled) {
		return false
	}
	if !util.OrderedListsMatch(a.From, b.From) {
		return false
	}
	if !util.StringsMatch(a.GroupTag, b.GroupTag) {
		return false
	}
	if !util.StringsMatch(a.NatType, b.NatType) {
		return false
	}
	if !util.StringsMatch(a.Service, b.Service) {
		return false
	}
	if !util.OrderedListsMatch(a.Source, b.Source) {
		return false
	}
	if !matchSourceTranslation(a.SourceTranslation, b.SourceTranslation) {
		return false
	}
	if !util.OrderedListsMatch(a.Tag, b.Tag) {
		return false
	}
	if !matchTarget(a.Target, b.Target) {
		return false
	}
	if !util.OrderedListsMatch(a.To, b.To) {
		return false
	}
	if !util.StringsMatch(a.ToInterface, b.ToInterface) {
		return false
	}
	if !util.StringsMatch(a.Uuid, b.Uuid) {
		return false
	}
	if !matchDestinationTranslation(a.DestinationTranslation, b.DestinationTranslation) {
		return false
	}
	if !matchDynamicDestinationTranslation(a.DynamicDestinationTranslation, b.DynamicDestinationTranslation) {
		return false
	}

	return true
}

func matchSourceTranslationDynamicIpFallbackInterfaceAddress(a *SourceTranslationDynamicIpFallbackInterfaceAddress, b *SourceTranslationDynamicIpFallbackInterfaceAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.FloatingIp, b.FloatingIp) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchSourceTranslationDynamicIpFallback(a *SourceTranslationDynamicIpFallback, b *SourceTranslationDynamicIpFallback) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchSourceTranslationDynamicIpFallbackInterfaceAddress(a.InterfaceAddress, b.InterfaceAddress) {
		return false
	}
	if !util.OrderedListsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	return true
}
func matchSourceTranslationDynamicIp(a *SourceTranslationDynamicIp, b *SourceTranslationDynamicIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchSourceTranslationDynamicIpFallback(a.Fallback, b.Fallback) {
		return false
	}
	if !util.OrderedListsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	return true
}
func matchSourceTranslationDynamicIpAndPortInterfaceAddress(a *SourceTranslationDynamicIpAndPortInterfaceAddress, b *SourceTranslationDynamicIpAndPortInterfaceAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.FloatingIp, b.FloatingIp) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchSourceTranslationDynamicIpAndPort(a *SourceTranslationDynamicIpAndPort, b *SourceTranslationDynamicIpAndPort) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchSourceTranslationDynamicIpAndPortInterfaceAddress(a.InterfaceAddress, b.InterfaceAddress) {
		return false
	}
	if !util.OrderedListsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	return true
}
func matchSourceTranslationStaticIp(a *SourceTranslationStaticIp, b *SourceTranslationStaticIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.BiDirectional, b.BiDirectional) {
		return false
	}
	if !util.StringsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	return true
}
func matchSourceTranslation(a *SourceTranslation, b *SourceTranslation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchSourceTranslationDynamicIp(a.DynamicIp, b.DynamicIp) {
		return false
	}
	if !matchSourceTranslationDynamicIpAndPort(a.DynamicIpAndPort, b.DynamicIpAndPort) {
		return false
	}
	if !matchSourceTranslationStaticIp(a.StaticIp, b.StaticIp) {
		return false
	}
	return true
}
func matchTargetDevicesVsys(a []TargetDevicesVsys, b []TargetDevicesVsys) bool {
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
func matchTargetDevices(a []TargetDevices, b []TargetDevices) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchTargetDevicesVsys(a.Vsys, b.Vsys) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchTarget(a *Target, b *Target) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTargetDevices(a.Devices, b.Devices) {
		return false
	}
	if !util.BoolsMatch(a.Negate, b.Negate) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}
	return true
}
func matchDestinationTranslationDnsRewrite(a *DestinationTranslationDnsRewrite, b *DestinationTranslationDnsRewrite) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Direction, b.Direction) {
		return false
	}
	return true
}
func matchDestinationTranslation(a *DestinationTranslation, b *DestinationTranslation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	if !util.Ints64Match(a.TranslatedPort, b.TranslatedPort) {
		return false
	}
	if !matchDestinationTranslationDnsRewrite(a.DnsRewrite, b.DnsRewrite) {
		return false
	}
	return true
}
func matchDynamicDestinationTranslation(a *DynamicDestinationTranslation, b *DynamicDestinationTranslation) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Distribution, b.Distribution) {
		return false
	}
	if !util.StringsMatch(a.TranslatedAddress, b.TranslatedAddress) {
		return false
	}
	if !util.Ints64Match(a.TranslatedPort, b.TranslatedPort) {
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
func (o *Entry) EntryUuid() *string {
	return o.Uuid
}

func (o *Entry) SetEntryUuid(uuid *string) {
	o.Uuid = uuid
}
