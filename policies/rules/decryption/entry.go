package decryption

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
	Name                string
	Action              *string
	Category            []string
	Description         *string
	Destination         []string
	DestinationHip      []string
	Disabled            *bool
	From                []string
	GroupTag            *string
	LogFail             *bool
	LogSetting          *string
	LogSuccess          *bool
	NegateDestination   *bool
	NegateSource        *bool
	PacketBrokerProfile *string
	Profile             *string
	Service             []string
	Source              []string
	SourceHip           []string
	SourceUser          []string
	Tag                 []string
	Target              *Target
	To                  []string
	Type                *Type
	Uuid                *string

	Misc map[string][]generic.Xml
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
type Type struct {
	SshProxy             *TypeSshProxy
	SslForwardProxy      *TypeSslForwardProxy
	SslInboundInspection *TypeSslInboundInspection
}
type TypeSshProxy struct {
}
type TypeSslForwardProxy struct {
}
type TypeSslInboundInspection struct {
	Certificates []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName             xml.Name         `xml:"entry"`
	Name                string           `xml:"name,attr"`
	Action              *string          `xml:"action,omitempty"`
	Category            *util.MemberType `xml:"category,omitempty"`
	Description         *string          `xml:"description,omitempty"`
	Destination         *util.MemberType `xml:"destination,omitempty"`
	DestinationHip      *util.MemberType `xml:"destination-hip,omitempty"`
	Disabled            *string          `xml:"disabled,omitempty"`
	From                *util.MemberType `xml:"from,omitempty"`
	GroupTag            *string          `xml:"group-tag,omitempty"`
	LogFail             *string          `xml:"log-fail,omitempty"`
	LogSetting          *string          `xml:"log-setting,omitempty"`
	LogSuccess          *string          `xml:"log-success,omitempty"`
	NegateDestination   *string          `xml:"negate-destination,omitempty"`
	NegateSource        *string          `xml:"negate-source,omitempty"`
	PacketBrokerProfile *string          `xml:"packet-broker-profile,omitempty"`
	Profile             *string          `xml:"profile,omitempty"`
	Service             *util.MemberType `xml:"service,omitempty"`
	Source              *util.MemberType `xml:"source,omitempty"`
	SourceHip           *util.MemberType `xml:"source-hip,omitempty"`
	SourceUser          *util.MemberType `xml:"source-user,omitempty"`
	Tag                 *util.MemberType `xml:"tag,omitempty"`
	Target              *TargetXml       `xml:"target,omitempty"`
	To                  *util.MemberType `xml:"to,omitempty"`
	Type                *TypeXml         `xml:"type,omitempty"`
	Uuid                *string          `xml:"uuid,attr,omitempty"`

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
type TypeXml struct {
	SshProxy             *TypeSshProxyXml             `xml:"ssh-proxy,omitempty"`
	SslForwardProxy      *TypeSslForwardProxyXml      `xml:"ssl-forward-proxy,omitempty"`
	SslInboundInspection *TypeSslInboundInspectionXml `xml:"ssl-inbound-inspection,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeSshProxyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeSslForwardProxyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeSslInboundInspectionXml struct {
	Certificates *util.MemberType `xml:"certificates,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "action" || v == "Action" {
		return e.Action, nil
	}
	if v == "category" || v == "Category" {
		return e.Category, nil
	}
	if v == "category|LENGTH" || v == "Category|LENGTH" {
		return int64(len(e.Category)), nil
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
	if v == "destination_hip" || v == "DestinationHip" {
		return e.DestinationHip, nil
	}
	if v == "destination_hip|LENGTH" || v == "DestinationHip|LENGTH" {
		return int64(len(e.DestinationHip)), nil
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
	if v == "log_fail" || v == "LogFail" {
		return e.LogFail, nil
	}
	if v == "log_setting" || v == "LogSetting" {
		return e.LogSetting, nil
	}
	if v == "log_success" || v == "LogSuccess" {
		return e.LogSuccess, nil
	}
	if v == "negate_destination" || v == "NegateDestination" {
		return e.NegateDestination, nil
	}
	if v == "negate_source" || v == "NegateSource" {
		return e.NegateSource, nil
	}
	if v == "packet_broker_profile" || v == "PacketBrokerProfile" {
		return e.PacketBrokerProfile, nil
	}
	if v == "profile" || v == "Profile" {
		return e.Profile, nil
	}
	if v == "service" || v == "Service" {
		return e.Service, nil
	}
	if v == "service|LENGTH" || v == "Service|LENGTH" {
		return int64(len(e.Service)), nil
	}
	if v == "source" || v == "Source" {
		return e.Source, nil
	}
	if v == "source|LENGTH" || v == "Source|LENGTH" {
		return int64(len(e.Source)), nil
	}
	if v == "source_hip" || v == "SourceHip" {
		return e.SourceHip, nil
	}
	if v == "source_hip|LENGTH" || v == "SourceHip|LENGTH" {
		return int64(len(e.SourceHip)), nil
	}
	if v == "source_user" || v == "SourceUser" {
		return e.SourceUser, nil
	}
	if v == "source_user|LENGTH" || v == "SourceUser|LENGTH" {
		return int64(len(e.SourceUser)), nil
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
	if v == "type" || v == "Type" {
		return e.Type, nil
	}
	if v == "uuid" || v == "Uuid" {
		return e.Uuid, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Action = o.Action
	entry.Category = util.StrToMem(o.Category)
	entry.Description = o.Description
	entry.Destination = util.StrToMem(o.Destination)
	entry.DestinationHip = util.StrToMem(o.DestinationHip)
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.From = util.StrToMem(o.From)
	entry.GroupTag = o.GroupTag
	entry.LogFail = util.YesNo(o.LogFail, nil)
	entry.LogSetting = o.LogSetting
	entry.LogSuccess = util.YesNo(o.LogSuccess, nil)
	entry.NegateDestination = util.YesNo(o.NegateDestination, nil)
	entry.NegateSource = util.YesNo(o.NegateSource, nil)
	entry.PacketBrokerProfile = o.PacketBrokerProfile
	entry.Profile = o.Profile
	entry.Service = util.StrToMem(o.Service)
	entry.Source = util.StrToMem(o.Source)
	entry.SourceHip = util.StrToMem(o.SourceHip)
	entry.SourceUser = util.StrToMem(o.SourceUser)
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
				if oTargetDevices.Name != "" {
					nestedTargetDevices.Name = oTargetDevices.Name
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
	var nestedType *TypeXml
	if o.Type != nil {
		nestedType = &TypeXml{}
		if _, ok := o.Misc["Type"]; ok {
			nestedType.Misc = o.Misc["Type"]
		}
		if o.Type.SslInboundInspection != nil {
			nestedType.SslInboundInspection = &TypeSslInboundInspectionXml{}
			if _, ok := o.Misc["TypeSslInboundInspection"]; ok {
				nestedType.SslInboundInspection.Misc = o.Misc["TypeSslInboundInspection"]
			}
			if o.Type.SslInboundInspection.Certificates != nil {
				nestedType.SslInboundInspection.Certificates = util.StrToMem(o.Type.SslInboundInspection.Certificates)
			}
		}
		if o.Type.SshProxy != nil {
			nestedType.SshProxy = &TypeSshProxyXml{}
			if _, ok := o.Misc["TypeSshProxy"]; ok {
				nestedType.SshProxy.Misc = o.Misc["TypeSshProxy"]
			}
		}
		if o.Type.SslForwardProxy != nil {
			nestedType.SslForwardProxy = &TypeSslForwardProxyXml{}
			if _, ok := o.Misc["TypeSslForwardProxy"]; ok {
				nestedType.SslForwardProxy.Misc = o.Misc["TypeSslForwardProxy"]
			}
		}
	}
	entry.Type = nestedType

	entry.Uuid = o.Uuid

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
		entry.Action = o.Action
		entry.Category = util.MemToStr(o.Category)
		entry.Description = o.Description
		entry.Destination = util.MemToStr(o.Destination)
		entry.DestinationHip = util.MemToStr(o.DestinationHip)
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.From = util.MemToStr(o.From)
		entry.GroupTag = o.GroupTag
		entry.LogFail = util.AsBool(o.LogFail, nil)
		entry.LogSetting = o.LogSetting
		entry.LogSuccess = util.AsBool(o.LogSuccess, nil)
		entry.NegateDestination = util.AsBool(o.NegateDestination, nil)
		entry.NegateSource = util.AsBool(o.NegateSource, nil)
		entry.PacketBrokerProfile = o.PacketBrokerProfile
		entry.Profile = o.Profile
		entry.Service = util.MemToStr(o.Service)
		entry.Source = util.MemToStr(o.Source)
		entry.SourceHip = util.MemToStr(o.SourceHip)
		entry.SourceUser = util.MemToStr(o.SourceUser)
		entry.Tag = util.MemToStr(o.Tag)
		var nestedTarget *Target
		if o.Target != nil {
			nestedTarget = &Target{}
			if o.Target.Misc != nil {
				entry.Misc["Target"] = o.Target.Misc
			}
			if o.Target.Negate != nil {
				nestedTarget.Negate = util.AsBool(o.Target.Negate, nil)
			}
			if o.Target.Tags != nil {
				nestedTarget.Tags = util.MemToStr(o.Target.Tags)
			}
			if o.Target.Devices != nil {
				nestedTarget.Devices = []TargetDevices{}
				for _, oTargetDevices := range o.Target.Devices {
					nestedTargetDevices := TargetDevices{}
					if oTargetDevices.Misc != nil {
						entry.Misc["TargetDevices"] = oTargetDevices.Misc
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
					if oTargetDevices.Name != "" {
						nestedTargetDevices.Name = oTargetDevices.Name
					}
					nestedTarget.Devices = append(nestedTarget.Devices, nestedTargetDevices)
				}
			}
		}
		entry.Target = nestedTarget

		entry.To = util.MemToStr(o.To)
		var nestedType *Type
		if o.Type != nil {
			nestedType = &Type{}
			if o.Type.Misc != nil {
				entry.Misc["Type"] = o.Type.Misc
			}
			if o.Type.SshProxy != nil {
				nestedType.SshProxy = &TypeSshProxy{}
				if o.Type.SshProxy.Misc != nil {
					entry.Misc["TypeSshProxy"] = o.Type.SshProxy.Misc
				}
			}
			if o.Type.SslForwardProxy != nil {
				nestedType.SslForwardProxy = &TypeSslForwardProxy{}
				if o.Type.SslForwardProxy.Misc != nil {
					entry.Misc["TypeSslForwardProxy"] = o.Type.SslForwardProxy.Misc
				}
			}
			if o.Type.SslInboundInspection != nil {
				nestedType.SslInboundInspection = &TypeSslInboundInspection{}
				if o.Type.SslInboundInspection.Misc != nil {
					entry.Misc["TypeSslInboundInspection"] = o.Type.SslInboundInspection.Misc
				}
				if o.Type.SslInboundInspection.Certificates != nil {
					nestedType.SslInboundInspection.Certificates = util.MemToStr(o.Type.SslInboundInspection.Certificates)
				}
			}
		}
		entry.Type = nestedType

		entry.Uuid = o.Uuid

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
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.OrderedListsMatch(a.Category, b.Category) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.Destination, b.Destination) {
		return false
	}
	if !util.OrderedListsMatch(a.DestinationHip, b.DestinationHip) {
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
	if !util.BoolsMatch(a.LogFail, b.LogFail) {
		return false
	}
	if !util.StringsMatch(a.LogSetting, b.LogSetting) {
		return false
	}
	if !util.BoolsMatch(a.LogSuccess, b.LogSuccess) {
		return false
	}
	if !util.BoolsMatch(a.NegateDestination, b.NegateDestination) {
		return false
	}
	if !util.BoolsMatch(a.NegateSource, b.NegateSource) {
		return false
	}
	if !util.StringsMatch(a.PacketBrokerProfile, b.PacketBrokerProfile) {
		return false
	}
	if !util.StringsMatch(a.Profile, b.Profile) {
		return false
	}
	if !util.OrderedListsMatch(a.Service, b.Service) {
		return false
	}
	if !util.OrderedListsMatch(a.Source, b.Source) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceHip, b.SourceHip) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceUser, b.SourceUser) {
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
	if !matchType(a.Type, b.Type) {
		return false
	}
	if !util.StringsMatch(a.Uuid, b.Uuid) {
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
	if !util.BoolsMatch(a.Negate, b.Negate) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}
	if !matchTargetDevices(a.Devices, b.Devices) {
		return false
	}
	return true
}
func matchTypeSshProxy(a *TypeSshProxy, b *TypeSshProxy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeSslForwardProxy(a *TypeSslForwardProxy, b *TypeSslForwardProxy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeSslInboundInspection(a *TypeSslInboundInspection, b *TypeSslInboundInspection) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Certificates, b.Certificates) {
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
	if !matchTypeSslInboundInspection(a.SslInboundInspection, b.SslInboundInspection) {
		return false
	}
	if !matchTypeSshProxy(a.SshProxy, b.SshProxy) {
		return false
	}
	if !matchTypeSslForwardProxy(a.SslForwardProxy, b.SslForwardProxy) {
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
