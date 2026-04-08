package authentication

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
	suffix = []string{"authentication", "rules", "$name"}
)

type Entry struct {
	Name                      string
	AuthenticationEnforcement *string
	Category                  []string
	Description               *string
	Destination               []string
	DestinationHip            []string
	Disabled                  *bool
	From                      []string
	GroupTag                  *string
	LogAuthenticationTimeout  *bool
	LogSetting                *string
	NegateDestination         *bool
	NegateSource              *bool
	Service                   []string
	Source                    []string
	SourceHip                 []string
	SourceUser                []string
	Tag                       []string
	Target                    *Target
	Timeout                   *int64
	To                        []string
	Uuid                      *string
	Misc                      []generic.Xml
	MiscAttributes            []xml.Attr
}
type Target struct {
	Devices        []TargetDevices
	Negate         *bool
	Tags           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TargetDevices struct {
	Name           string
	Vsys           []TargetDevicesVsys
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TargetDevicesVsys struct {
	Name           string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName                   xml.Name         `xml:"entry"`
	Name                      string           `xml:"name,attr"`
	AuthenticationEnforcement *string          `xml:"authentication-enforcement,omitempty"`
	Category                  *util.MemberType `xml:"category,omitempty"`
	Description               *string          `xml:"description,omitempty"`
	Destination               *util.MemberType `xml:"destination,omitempty"`
	DestinationHip            *util.MemberType `xml:"destination-hip,omitempty"`
	Disabled                  *string          `xml:"disabled,omitempty"`
	From                      *util.MemberType `xml:"from,omitempty"`
	GroupTag                  *string          `xml:"group-tag,omitempty"`
	LogAuthenticationTimeout  *string          `xml:"log-authentication-timeout,omitempty"`
	LogSetting                *string          `xml:"log-setting,omitempty"`
	NegateDestination         *string          `xml:"negate-destination,omitempty"`
	NegateSource              *string          `xml:"negate-source,omitempty"`
	Service                   *util.MemberType `xml:"service,omitempty"`
	Source                    *util.MemberType `xml:"source,omitempty"`
	SourceHip                 *util.MemberType `xml:"source-hip,omitempty"`
	SourceUser                *util.MemberType `xml:"source-user,omitempty"`
	Tag                       *util.MemberType `xml:"tag,omitempty"`
	Target                    *targetXml       `xml:"target,omitempty"`
	Timeout                   *int64           `xml:"timeout,omitempty"`
	To                        *util.MemberType `xml:"to,omitempty"`
	Uuid                      *string          `xml:"uuid,attr,omitempty"`
	Misc                      []generic.Xml    `xml:",any"`
	MiscAttributes            []xml.Attr       `xml:",any,attr"`
}
type targetXml struct {
	Devices        *targetDevicesContainerXml `xml:"devices,omitempty"`
	Negate         *string                    `xml:"negate,omitempty"`
	Tags           *util.MemberType           `xml:"tags,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type targetDevicesContainerXml struct {
	Entries []targetDevicesXml `xml:"entry"`
}
type targetDevicesXml struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Vsys           *targetDevicesVsysContainerXml `xml:"vsys,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type targetDevicesVsysContainerXml struct {
	Entries []targetDevicesVsysXml `xml:"entry"`
}
type targetDevicesVsysXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AuthenticationEnforcement = s.AuthenticationEnforcement
	if s.Category != nil {
		o.Category = util.StrToMem(s.Category)
	}
	o.Description = s.Description
	if s.Destination != nil {
		o.Destination = util.StrToMem(s.Destination)
	}
	if s.DestinationHip != nil {
		o.DestinationHip = util.StrToMem(s.DestinationHip)
	}
	o.Disabled = util.YesNo(s.Disabled, nil)
	if s.From != nil {
		o.From = util.StrToMem(s.From)
	}
	o.GroupTag = s.GroupTag
	o.LogAuthenticationTimeout = util.YesNo(s.LogAuthenticationTimeout, nil)
	o.LogSetting = s.LogSetting
	o.NegateDestination = util.YesNo(s.NegateDestination, nil)
	o.NegateSource = util.YesNo(s.NegateSource, nil)
	if s.Service != nil {
		o.Service = util.StrToMem(s.Service)
	}
	if s.Source != nil {
		o.Source = util.StrToMem(s.Source)
	}
	if s.SourceHip != nil {
		o.SourceHip = util.StrToMem(s.SourceHip)
	}
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	if s.Target != nil {
		var obj targetXml
		obj.MarshalFromObject(*s.Target)
		o.Target = &obj
	}
	o.Timeout = s.Timeout
	if s.To != nil {
		o.To = util.StrToMem(s.To)
	}
	o.Uuid = s.Uuid
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var categoryVal []string
	if o.Category != nil {
		categoryVal = util.MemToStr(o.Category)
	}
	var destinationVal []string
	if o.Destination != nil {
		destinationVal = util.MemToStr(o.Destination)
	}
	var destinationHipVal []string
	if o.DestinationHip != nil {
		destinationHipVal = util.MemToStr(o.DestinationHip)
	}
	var fromVal []string
	if o.From != nil {
		fromVal = util.MemToStr(o.From)
	}
	var serviceVal []string
	if o.Service != nil {
		serviceVal = util.MemToStr(o.Service)
	}
	var sourceVal []string
	if o.Source != nil {
		sourceVal = util.MemToStr(o.Source)
	}
	var sourceHipVal []string
	if o.SourceHip != nil {
		sourceHipVal = util.MemToStr(o.SourceHip)
	}
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}
	var targetVal *Target
	if o.Target != nil {
		obj, err := o.Target.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		targetVal = obj
	}
	var toVal []string
	if o.To != nil {
		toVal = util.MemToStr(o.To)
	}

	result := &Entry{
		Name:                      o.Name,
		AuthenticationEnforcement: o.AuthenticationEnforcement,
		Category:                  categoryVal,
		Description:               o.Description,
		Destination:               destinationVal,
		DestinationHip:            destinationHipVal,
		Disabled:                  util.AsBool(o.Disabled, nil),
		From:                      fromVal,
		GroupTag:                  o.GroupTag,
		LogAuthenticationTimeout:  util.AsBool(o.LogAuthenticationTimeout, nil),
		LogSetting:                o.LogSetting,
		NegateDestination:         util.AsBool(o.NegateDestination, nil),
		NegateSource:              util.AsBool(o.NegateSource, nil),
		Service:                   serviceVal,
		Source:                    sourceVal,
		SourceHip:                 sourceHipVal,
		SourceUser:                sourceUserVal,
		Tag:                       tagVal,
		Target:                    targetVal,
		Timeout:                   o.Timeout,
		To:                        toVal,
		Uuid:                      o.Uuid,
		Misc:                      o.Misc,
		MiscAttributes:            o.MiscAttributes,
	}
	return result, nil
}
func (o *targetXml) MarshalFromObject(s Target) {
	if s.Devices != nil {
		var objs []targetDevicesXml
		for _, elt := range s.Devices {
			var obj targetDevicesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Devices = &targetDevicesContainerXml{Entries: objs}
	}
	o.Negate = util.YesNo(s.Negate, nil)
	if s.Tags != nil {
		o.Tags = util.StrToMem(s.Tags)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetXml) UnmarshalToObject() (*Target, error) {
	var devicesVal []TargetDevices
	if o.Devices != nil {
		for _, elt := range o.Devices.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			devicesVal = append(devicesVal, *obj)
		}
	}
	var tagsVal []string
	if o.Tags != nil {
		tagsVal = util.MemToStr(o.Tags)
	}

	result := &Target{
		Devices:        devicesVal,
		Negate:         util.AsBool(o.Negate, nil),
		Tags:           tagsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *targetDevicesXml) MarshalFromObject(s TargetDevices) {
	o.Name = s.Name
	if s.Vsys != nil {
		var objs []targetDevicesVsysXml
		for _, elt := range s.Vsys {
			var obj targetDevicesVsysXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vsys = &targetDevicesVsysContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetDevicesXml) UnmarshalToObject() (*TargetDevices, error) {
	var vsysVal []TargetDevicesVsys
	if o.Vsys != nil {
		for _, elt := range o.Vsys.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vsysVal = append(vsysVal, *obj)
		}
	}

	result := &TargetDevices{
		Name:           o.Name,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *targetDevicesVsysXml) MarshalFromObject(s TargetDevicesVsys) {
	o.Name = s.Name
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetDevicesVsysXml) UnmarshalToObject() (*TargetDevicesVsys, error) {

	result := &TargetDevicesVsys{
		Name:           o.Name,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authentication_enforcement" || v == "AuthenticationEnforcement" {
		return e.AuthenticationEnforcement, nil
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
	if v == "log_authentication_timeout" || v == "LogAuthenticationTimeout" {
		return e.LogAuthenticationTimeout, nil
	}
	if v == "log_setting" || v == "LogSetting" {
		return e.LogSetting, nil
	}
	if v == "negate_destination" || v == "NegateDestination" {
		return e.NegateDestination, nil
	}
	if v == "negate_source" || v == "NegateSource" {
		return e.NegateSource, nil
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
	if v == "timeout" || v == "Timeout" {
		return e.Timeout, nil
	}
	if v == "to" || v == "To" {
		return e.To, nil
	}
	if v == "to|LENGTH" || v == "To|LENGTH" {
		return int64(len(e.To)), nil
	}
	if v == "uuid" || v == "Uuid" {
		return e.Uuid, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationEnforcement, other.AuthenticationEnforcement) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Category, other.Category) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Destination, other.Destination) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DestinationHip, other.DestinationHip) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !util.OrderedListsMatch[string](o.From, other.From) {
		return false
	}
	if !util.StringsMatch(o.GroupTag, other.GroupTag) {
		return false
	}
	if !util.BoolsMatch(o.LogAuthenticationTimeout, other.LogAuthenticationTimeout) {
		return false
	}
	if !util.StringsMatch(o.LogSetting, other.LogSetting) {
		return false
	}
	if !util.BoolsMatch(o.NegateDestination, other.NegateDestination) {
		return false
	}
	if !util.BoolsMatch(o.NegateSource, other.NegateSource) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Service, other.Service) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Source, other.Source) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceHip, other.SourceHip) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceUser, other.SourceUser) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}
	if !o.Target.matches(other.Target) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !util.OrderedListsMatch[string](o.To, other.To) {
		return false
	}
	if !util.StringsMatch(o.Uuid, other.Uuid) {
		return false
	}

	return true
}

func (o *Target) matches(other *Target) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Devices) != len(other.Devices) {
		return false
	}
	for idx := range o.Devices {
		if !o.Devices[idx].matches(&other.Devices[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tags, other.Tags) {
		return false
	}

	return true
}

func (o *TargetDevices) matches(other *TargetDevices) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.Vsys) != len(other.Vsys) {
		return false
	}
	for idx := range o.Vsys {
		if !o.Vsys[idx].matches(&other.Vsys[idx]) {
			return false
		}
	}

	return true
}

func (o *TargetDevicesVsys) matches(other *TargetDevicesVsys) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
func (o *Entry) EntryUuid() *string {
	return o.Uuid
}

func (o *Entry) SetEntryUuid(uuid *string) {
	o.Uuid = uuid
}
