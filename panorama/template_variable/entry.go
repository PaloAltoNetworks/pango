package template_variable

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
	Suffix = []string{"variable"}
)

type Entry struct {
	Name        string
	Description *string
	Type        *Type

	Misc map[string][]generic.Xml
}

type Type struct {
	AsNumber       *string
	DeviceId       *string
	DevicePriority *string
	EgressMax      *string
	Fqdn           *string
	GroupId        *string
	Interface      *string
	IpNetmask      *string
	IpRange        *string
	LinkTag        *string
	QosProfile     *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Description *string  `xml:"description,omitempty"`
	Type        *TypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type TypeXml struct {
	AsNumber       *string `xml:"as-number,omitempty"`
	DeviceId       *string `xml:"device-id,omitempty"`
	DevicePriority *string `xml:"device-priority,omitempty"`
	EgressMax      *string `xml:"egress-max,omitempty"`
	Fqdn           *string `xml:"fqdn,omitempty"`
	GroupId        *string `xml:"group-id,omitempty"`
	Interface      *string `xml:"interface,omitempty"`
	IpNetmask      *string `xml:"ip-netmask,omitempty"`
	IpRange        *string `xml:"ip-range,omitempty"`
	LinkTag        *string `xml:"link-tag,omitempty"`
	QosProfile     *string `xml:"qos-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "type" || v == "Type" {
		return e.Type, nil
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
	var nestedType *TypeXml
	if o.Type != nil {
		nestedType = &TypeXml{}
		if _, ok := o.Misc["Type"]; ok {
			nestedType.Misc = o.Misc["Type"]
		}
		if o.Type.Fqdn != nil {
			nestedType.Fqdn = o.Type.Fqdn
		}
		if o.Type.DevicePriority != nil {
			nestedType.DevicePriority = o.Type.DevicePriority
		}
		if o.Type.DeviceId != nil {
			nestedType.DeviceId = o.Type.DeviceId
		}
		if o.Type.Interface != nil {
			nestedType.Interface = o.Type.Interface
		}
		if o.Type.EgressMax != nil {
			nestedType.EgressMax = o.Type.EgressMax
		}
		if o.Type.IpRange != nil {
			nestedType.IpRange = o.Type.IpRange
		}
		if o.Type.GroupId != nil {
			nestedType.GroupId = o.Type.GroupId
		}
		if o.Type.AsNumber != nil {
			nestedType.AsNumber = o.Type.AsNumber
		}
		if o.Type.QosProfile != nil {
			nestedType.QosProfile = o.Type.QosProfile
		}
		if o.Type.LinkTag != nil {
			nestedType.LinkTag = o.Type.LinkTag
		}
		if o.Type.IpNetmask != nil {
			nestedType.IpNetmask = o.Type.IpNetmask
		}
	}
	entry.Type = nestedType

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
		var nestedType *Type
		if o.Type != nil {
			nestedType = &Type{}
			if o.Type.Misc != nil {
				entry.Misc["Type"] = o.Type.Misc
			}
			if o.Type.IpNetmask != nil {
				nestedType.IpNetmask = o.Type.IpNetmask
			}
			if o.Type.GroupId != nil {
				nestedType.GroupId = o.Type.GroupId
			}
			if o.Type.AsNumber != nil {
				nestedType.AsNumber = o.Type.AsNumber
			}
			if o.Type.QosProfile != nil {
				nestedType.QosProfile = o.Type.QosProfile
			}
			if o.Type.LinkTag != nil {
				nestedType.LinkTag = o.Type.LinkTag
			}
			if o.Type.IpRange != nil {
				nestedType.IpRange = o.Type.IpRange
			}
			if o.Type.Fqdn != nil {
				nestedType.Fqdn = o.Type.Fqdn
			}
			if o.Type.DevicePriority != nil {
				nestedType.DevicePriority = o.Type.DevicePriority
			}
			if o.Type.DeviceId != nil {
				nestedType.DeviceId = o.Type.DeviceId
			}
			if o.Type.Interface != nil {
				nestedType.Interface = o.Type.Interface
			}
			if o.Type.EgressMax != nil {
				nestedType.EgressMax = o.Type.EgressMax
			}
		}
		entry.Type = nestedType

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
	if !matchType(a.Type, b.Type) {
		return false
	}

	return true
}

func matchType(a *Type, b *Type) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.IpNetmask, b.IpNetmask) {
		return false
	}
	if !util.StringsMatch(a.GroupId, b.GroupId) {
		return false
	}
	if !util.StringsMatch(a.AsNumber, b.AsNumber) {
		return false
	}
	if !util.StringsMatch(a.QosProfile, b.QosProfile) {
		return false
	}
	if !util.StringsMatch(a.LinkTag, b.LinkTag) {
		return false
	}
	if !util.StringsMatch(a.IpRange, b.IpRange) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.DevicePriority, b.DevicePriority) {
		return false
	}
	if !util.StringsMatch(a.DeviceId, b.DeviceId) {
		return false
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.EgressMax, b.EgressMax) {
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
