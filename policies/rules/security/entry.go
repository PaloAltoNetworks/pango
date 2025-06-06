package security

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
	suffix = []string{"security", "rules", "$name"}
)

type Entry struct {
	Name                            string
	Action                          *string
	Application                     []string
	Category                        []string
	Description                     *string
	Destination                     []string
	DestinationHip                  []string
	DisableInspect                  *bool
	DisableServerResponseInspection *bool
	Disabled                        *bool
	From                            []string
	GroupTag                        *string
	IcmpUnreachable                 *bool
	LogEnd                          *bool
	LogSetting                      *string
	LogStart                        *bool
	NegateDestination               *bool
	NegateSource                    *bool
	ProfileSetting                  *ProfileSetting
	Qos                             *Qos
	RuleType                        *string
	Schedule                        *string
	Service                         []string
	Source                          []string
	SourceHip                       []string
	SourceImei                      []string
	SourceImsi                      []string
	SourceNwSlice                   []string
	SourceUser                      []string
	Tag                             []string
	Target                          *Target
	To                              []string
	Uuid                            *string
	Misc                            []generic.Xml
}
type ProfileSetting struct {
	Group    []string
	Profiles *ProfileSettingProfiles
	Misc     []generic.Xml
}
type ProfileSettingProfiles struct {
	DataFiltering    []string
	FileBlocking     []string
	Gtp              []string
	Sctp             []string
	Spyware          []string
	UrlFiltering     []string
	Virus            []string
	Vulnerability    []string
	WildfireAnalysis []string
	Misc             []generic.Xml
}
type Qos struct {
	Marking *QosMarking
	Misc    []generic.Xml
}
type QosMarking struct {
	FollowC2sFlow *QosMarkingFollowC2sFlow
	IpDscp        *string
	IpPrecedence  *string
	Misc          []generic.Xml
}
type QosMarkingFollowC2sFlow struct {
	Misc []generic.Xml
}
type Target struct {
	Devices []TargetDevices
	Negate  *bool
	Tags    []string
	Misc    []generic.Xml
}
type TargetDevices struct {
	Name string
	Vsys []TargetDevicesVsys
	Misc []generic.Xml
}
type TargetDevicesVsys struct {
	Name string
	Misc []generic.Xml
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
	XMLName                         xml.Name           `xml:"entry"`
	Name                            string             `xml:"name,attr"`
	Action                          *string            `xml:"action,omitempty"`
	Application                     *util.MemberType   `xml:"application,omitempty"`
	Category                        *util.MemberType   `xml:"category,omitempty"`
	Description                     *string            `xml:"description,omitempty"`
	Destination                     *util.MemberType   `xml:"destination,omitempty"`
	DestinationHip                  *util.MemberType   `xml:"destination-hip,omitempty"`
	DisableInspect                  *string            `xml:"disable-inspect,omitempty"`
	DisableServerResponseInspection *string            `xml:"option>disable-server-response-inspection,omitempty"`
	Disabled                        *string            `xml:"disabled,omitempty"`
	From                            *util.MemberType   `xml:"from,omitempty"`
	GroupTag                        *string            `xml:"group-tag,omitempty"`
	IcmpUnreachable                 *string            `xml:"icmp-unreachable,omitempty"`
	LogEnd                          *string            `xml:"log-end,omitempty"`
	LogSetting                      *string            `xml:"log-setting,omitempty"`
	LogStart                        *string            `xml:"log-start,omitempty"`
	NegateDestination               *string            `xml:"negate-destination,omitempty"`
	NegateSource                    *string            `xml:"negate-source,omitempty"`
	ProfileSetting                  *profileSettingXml `xml:"profile-setting,omitempty"`
	Qos                             *qosXml            `xml:"qos,omitempty"`
	RuleType                        *string            `xml:"rule-type,omitempty"`
	Schedule                        *string            `xml:"schedule,omitempty"`
	Service                         *util.MemberType   `xml:"service,omitempty"`
	Source                          *util.MemberType   `xml:"source,omitempty"`
	SourceHip                       *util.MemberType   `xml:"source-hip,omitempty"`
	SourceImei                      *util.MemberType   `xml:"source-imei,omitempty"`
	SourceImsi                      *util.MemberType   `xml:"source-imsi,omitempty"`
	SourceNwSlice                   *util.MemberType   `xml:"source-nw-slice,omitempty"`
	SourceUser                      *util.MemberType   `xml:"source-user,omitempty"`
	Tag                             *util.MemberType   `xml:"tag,omitempty"`
	Target                          *targetXml         `xml:"target,omitempty"`
	To                              *util.MemberType   `xml:"to,omitempty"`
	Uuid                            *string            `xml:"uuid,attr,omitempty"`
	Misc                            []generic.Xml      `xml:",any"`
}
type profileSettingXml struct {
	Group    *util.MemberType           `xml:"group,omitempty"`
	Profiles *profileSettingProfilesXml `xml:"profiles,omitempty"`
	Misc     []generic.Xml              `xml:",any"`
}
type profileSettingProfilesXml struct {
	DataFiltering    *util.MemberType `xml:"data-filtering,omitempty"`
	FileBlocking     *util.MemberType `xml:"file-blocking,omitempty"`
	Gtp              *util.MemberType `xml:"gtp,omitempty"`
	Sctp             *util.MemberType `xml:"sctp,omitempty"`
	Spyware          *util.MemberType `xml:"spyware,omitempty"`
	UrlFiltering     *util.MemberType `xml:"url-filtering,omitempty"`
	Virus            *util.MemberType `xml:"virus,omitempty"`
	Vulnerability    *util.MemberType `xml:"vulnerability,omitempty"`
	WildfireAnalysis *util.MemberType `xml:"wildfire-analysis,omitempty"`
	Misc             []generic.Xml    `xml:",any"`
}
type qosXml struct {
	Marking *qosMarkingXml `xml:"marking,omitempty"`
	Misc    []generic.Xml  `xml:",any"`
}
type qosMarkingXml struct {
	FollowC2sFlow *qosMarkingFollowC2sFlowXml `xml:"follow-c2s-flow,omitempty"`
	IpDscp        *string                     `xml:"ip-dscp,omitempty"`
	IpPrecedence  *string                     `xml:"ip-precedence,omitempty"`
	Misc          []generic.Xml               `xml:",any"`
}
type qosMarkingFollowC2sFlowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type targetXml struct {
	Devices *targetDevicesContainerXml `xml:"devices,omitempty"`
	Negate  *string                    `xml:"negate,omitempty"`
	Tags    *util.MemberType           `xml:"tags,omitempty"`
	Misc    []generic.Xml              `xml:",any"`
}
type targetDevicesContainerXml struct {
	Entries []targetDevicesXml `xml:"entry"`
}
type targetDevicesXml struct {
	XMLName xml.Name                       `xml:"entry"`
	Name    string                         `xml:"name,attr"`
	Vsys    *targetDevicesVsysContainerXml `xml:"vsys,omitempty"`
	Misc    []generic.Xml                  `xml:",any"`
}
type targetDevicesVsysContainerXml struct {
	Entries []targetDevicesVsysXml `xml:"entry"`
}
type targetDevicesVsysXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Action = s.Action
	if s.Application != nil {
		o.Application = util.StrToMem(s.Application)
	}
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
	o.DisableInspect = util.YesNo(s.DisableInspect, nil)
	o.DisableServerResponseInspection = util.YesNo(s.DisableServerResponseInspection, nil)
	o.Disabled = util.YesNo(s.Disabled, nil)
	if s.From != nil {
		o.From = util.StrToMem(s.From)
	}
	o.GroupTag = s.GroupTag
	o.IcmpUnreachable = util.YesNo(s.IcmpUnreachable, nil)
	o.LogEnd = util.YesNo(s.LogEnd, nil)
	o.LogSetting = s.LogSetting
	o.LogStart = util.YesNo(s.LogStart, nil)
	o.NegateDestination = util.YesNo(s.NegateDestination, nil)
	o.NegateSource = util.YesNo(s.NegateSource, nil)
	if s.ProfileSetting != nil {
		var obj profileSettingXml
		obj.MarshalFromObject(*s.ProfileSetting)
		o.ProfileSetting = &obj
	}
	if s.Qos != nil {
		var obj qosXml
		obj.MarshalFromObject(*s.Qos)
		o.Qos = &obj
	}
	o.RuleType = s.RuleType
	o.Schedule = s.Schedule
	if s.Service != nil {
		o.Service = util.StrToMem(s.Service)
	}
	if s.Source != nil {
		o.Source = util.StrToMem(s.Source)
	}
	if s.SourceHip != nil {
		o.SourceHip = util.StrToMem(s.SourceHip)
	}
	if s.SourceImei != nil {
		o.SourceImei = util.StrToMem(s.SourceImei)
	}
	if s.SourceImsi != nil {
		o.SourceImsi = util.StrToMem(s.SourceImsi)
	}
	if s.SourceNwSlice != nil {
		o.SourceNwSlice = util.StrToMem(s.SourceNwSlice)
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
	if s.To != nil {
		o.To = util.StrToMem(s.To)
	}
	o.Uuid = s.Uuid
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var applicationVal []string
	if o.Application != nil {
		applicationVal = util.MemToStr(o.Application)
	}
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
	var profileSettingVal *ProfileSetting
	if o.ProfileSetting != nil {
		obj, err := o.ProfileSetting.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		profileSettingVal = obj
	}
	var qosVal *Qos
	if o.Qos != nil {
		obj, err := o.Qos.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		qosVal = obj
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
	var sourceImeiVal []string
	if o.SourceImei != nil {
		sourceImeiVal = util.MemToStr(o.SourceImei)
	}
	var sourceImsiVal []string
	if o.SourceImsi != nil {
		sourceImsiVal = util.MemToStr(o.SourceImsi)
	}
	var sourceNwSliceVal []string
	if o.SourceNwSlice != nil {
		sourceNwSliceVal = util.MemToStr(o.SourceNwSlice)
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
		Name:                            o.Name,
		Action:                          o.Action,
		Application:                     applicationVal,
		Category:                        categoryVal,
		Description:                     o.Description,
		Destination:                     destinationVal,
		DestinationHip:                  destinationHipVal,
		DisableInspect:                  util.AsBool(o.DisableInspect, nil),
		DisableServerResponseInspection: util.AsBool(o.DisableServerResponseInspection, nil),
		Disabled:                        util.AsBool(o.Disabled, nil),
		From:                            fromVal,
		GroupTag:                        o.GroupTag,
		IcmpUnreachable:                 util.AsBool(o.IcmpUnreachable, nil),
		LogEnd:                          util.AsBool(o.LogEnd, nil),
		LogSetting:                      o.LogSetting,
		LogStart:                        util.AsBool(o.LogStart, nil),
		NegateDestination:               util.AsBool(o.NegateDestination, nil),
		NegateSource:                    util.AsBool(o.NegateSource, nil),
		ProfileSetting:                  profileSettingVal,
		Qos:                             qosVal,
		RuleType:                        o.RuleType,
		Schedule:                        o.Schedule,
		Service:                         serviceVal,
		Source:                          sourceVal,
		SourceHip:                       sourceHipVal,
		SourceImei:                      sourceImeiVal,
		SourceImsi:                      sourceImsiVal,
		SourceNwSlice:                   sourceNwSliceVal,
		SourceUser:                      sourceUserVal,
		Tag:                             tagVal,
		Target:                          targetVal,
		To:                              toVal,
		Uuid:                            o.Uuid,
		Misc:                            o.Misc,
	}
	return result, nil
}
func (o *profileSettingXml) MarshalFromObject(s ProfileSetting) {
	if s.Group != nil {
		o.Group = util.StrToMem(s.Group)
	}
	if s.Profiles != nil {
		var obj profileSettingProfilesXml
		obj.MarshalFromObject(*s.Profiles)
		o.Profiles = &obj
	}
	o.Misc = s.Misc
}

func (o profileSettingXml) UnmarshalToObject() (*ProfileSetting, error) {
	var groupVal []string
	if o.Group != nil {
		groupVal = util.MemToStr(o.Group)
	}
	var profilesVal *ProfileSettingProfiles
	if o.Profiles != nil {
		obj, err := o.Profiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		profilesVal = obj
	}

	result := &ProfileSetting{
		Group:    groupVal,
		Profiles: profilesVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *profileSettingProfilesXml) MarshalFromObject(s ProfileSettingProfiles) {
	if s.DataFiltering != nil {
		o.DataFiltering = util.StrToMem(s.DataFiltering)
	}
	if s.FileBlocking != nil {
		o.FileBlocking = util.StrToMem(s.FileBlocking)
	}
	if s.Gtp != nil {
		o.Gtp = util.StrToMem(s.Gtp)
	}
	if s.Sctp != nil {
		o.Sctp = util.StrToMem(s.Sctp)
	}
	if s.Spyware != nil {
		o.Spyware = util.StrToMem(s.Spyware)
	}
	if s.UrlFiltering != nil {
		o.UrlFiltering = util.StrToMem(s.UrlFiltering)
	}
	if s.Virus != nil {
		o.Virus = util.StrToMem(s.Virus)
	}
	if s.Vulnerability != nil {
		o.Vulnerability = util.StrToMem(s.Vulnerability)
	}
	if s.WildfireAnalysis != nil {
		o.WildfireAnalysis = util.StrToMem(s.WildfireAnalysis)
	}
	o.Misc = s.Misc
}

func (o profileSettingProfilesXml) UnmarshalToObject() (*ProfileSettingProfiles, error) {
	var dataFilteringVal []string
	if o.DataFiltering != nil {
		dataFilteringVal = util.MemToStr(o.DataFiltering)
	}
	var fileBlockingVal []string
	if o.FileBlocking != nil {
		fileBlockingVal = util.MemToStr(o.FileBlocking)
	}
	var gtpVal []string
	if o.Gtp != nil {
		gtpVal = util.MemToStr(o.Gtp)
	}
	var sctpVal []string
	if o.Sctp != nil {
		sctpVal = util.MemToStr(o.Sctp)
	}
	var spywareVal []string
	if o.Spyware != nil {
		spywareVal = util.MemToStr(o.Spyware)
	}
	var urlFilteringVal []string
	if o.UrlFiltering != nil {
		urlFilteringVal = util.MemToStr(o.UrlFiltering)
	}
	var virusVal []string
	if o.Virus != nil {
		virusVal = util.MemToStr(o.Virus)
	}
	var vulnerabilityVal []string
	if o.Vulnerability != nil {
		vulnerabilityVal = util.MemToStr(o.Vulnerability)
	}
	var wildfireAnalysisVal []string
	if o.WildfireAnalysis != nil {
		wildfireAnalysisVal = util.MemToStr(o.WildfireAnalysis)
	}

	result := &ProfileSettingProfiles{
		DataFiltering:    dataFilteringVal,
		FileBlocking:     fileBlockingVal,
		Gtp:              gtpVal,
		Sctp:             sctpVal,
		Spyware:          spywareVal,
		UrlFiltering:     urlFilteringVal,
		Virus:            virusVal,
		Vulnerability:    vulnerabilityVal,
		WildfireAnalysis: wildfireAnalysisVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *qosXml) MarshalFromObject(s Qos) {
	if s.Marking != nil {
		var obj qosMarkingXml
		obj.MarshalFromObject(*s.Marking)
		o.Marking = &obj
	}
	o.Misc = s.Misc
}

func (o qosXml) UnmarshalToObject() (*Qos, error) {
	var markingVal *QosMarking
	if o.Marking != nil {
		obj, err := o.Marking.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		markingVal = obj
	}

	result := &Qos{
		Marking: markingVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *qosMarkingXml) MarshalFromObject(s QosMarking) {
	if s.FollowC2sFlow != nil {
		var obj qosMarkingFollowC2sFlowXml
		obj.MarshalFromObject(*s.FollowC2sFlow)
		o.FollowC2sFlow = &obj
	}
	o.IpDscp = s.IpDscp
	o.IpPrecedence = s.IpPrecedence
	o.Misc = s.Misc
}

func (o qosMarkingXml) UnmarshalToObject() (*QosMarking, error) {
	var followC2sFlowVal *QosMarkingFollowC2sFlow
	if o.FollowC2sFlow != nil {
		obj, err := o.FollowC2sFlow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		followC2sFlowVal = obj
	}

	result := &QosMarking{
		FollowC2sFlow: followC2sFlowVal,
		IpDscp:        o.IpDscp,
		IpPrecedence:  o.IpPrecedence,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *qosMarkingFollowC2sFlowXml) MarshalFromObject(s QosMarkingFollowC2sFlow) {
	o.Misc = s.Misc
}

func (o qosMarkingFollowC2sFlowXml) UnmarshalToObject() (*QosMarkingFollowC2sFlow, error) {

	result := &QosMarkingFollowC2sFlow{
		Misc: o.Misc,
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
		Devices: devicesVal,
		Negate:  util.AsBool(o.Negate, nil),
		Tags:    tagsVal,
		Misc:    o.Misc,
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
		Name: o.Name,
		Vsys: vsysVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *targetDevicesVsysXml) MarshalFromObject(s TargetDevicesVsys) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o targetDevicesVsysXml) UnmarshalToObject() (*TargetDevicesVsys, error) {

	result := &TargetDevicesVsys{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "action" || v == "Action" {
		return e.Action, nil
	}
	if v == "application" || v == "Application" {
		return e.Application, nil
	}
	if v == "application|LENGTH" || v == "Application|LENGTH" {
		return int64(len(e.Application)), nil
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
	if v == "disable_inspect" || v == "DisableInspect" {
		return e.DisableInspect, nil
	}
	if v == "disable_server_response_inspection" || v == "DisableServerResponseInspection" {
		return e.DisableServerResponseInspection, nil
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
	if v == "icmp_unreachable" || v == "IcmpUnreachable" {
		return e.IcmpUnreachable, nil
	}
	if v == "log_end" || v == "LogEnd" {
		return e.LogEnd, nil
	}
	if v == "log_setting" || v == "LogSetting" {
		return e.LogSetting, nil
	}
	if v == "log_start" || v == "LogStart" {
		return e.LogStart, nil
	}
	if v == "negate_destination" || v == "NegateDestination" {
		return e.NegateDestination, nil
	}
	if v == "negate_source" || v == "NegateSource" {
		return e.NegateSource, nil
	}
	if v == "profile_setting" || v == "ProfileSetting" {
		return e.ProfileSetting, nil
	}
	if v == "qos" || v == "Qos" {
		return e.Qos, nil
	}
	if v == "rule_type" || v == "RuleType" {
		return e.RuleType, nil
	}
	if v == "schedule" || v == "Schedule" {
		return e.Schedule, nil
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
	if v == "source_imei" || v == "SourceImei" {
		return e.SourceImei, nil
	}
	if v == "source_imei|LENGTH" || v == "SourceImei|LENGTH" {
		return int64(len(e.SourceImei)), nil
	}
	if v == "source_imsi" || v == "SourceImsi" {
		return e.SourceImsi, nil
	}
	if v == "source_imsi|LENGTH" || v == "SourceImsi|LENGTH" {
		return int64(len(e.SourceImsi)), nil
	}
	if v == "source_nw_slice" || v == "SourceNwSlice" {
		return e.SourceNwSlice, nil
	}
	if v == "source_nw_slice|LENGTH" || v == "SourceNwSlice|LENGTH" {
		return int64(len(e.SourceNwSlice)), nil
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
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Application, other.Application) {
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
	if !util.BoolsMatch(o.DisableInspect, other.DisableInspect) {
		return false
	}
	if !util.BoolsMatch(o.DisableServerResponseInspection, other.DisableServerResponseInspection) {
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
	if !util.BoolsMatch(o.IcmpUnreachable, other.IcmpUnreachable) {
		return false
	}
	if !util.BoolsMatch(o.LogEnd, other.LogEnd) {
		return false
	}
	if !util.StringsMatch(o.LogSetting, other.LogSetting) {
		return false
	}
	if !util.BoolsMatch(o.LogStart, other.LogStart) {
		return false
	}
	if !util.BoolsMatch(o.NegateDestination, other.NegateDestination) {
		return false
	}
	if !util.BoolsMatch(o.NegateSource, other.NegateSource) {
		return false
	}
	if !o.ProfileSetting.matches(other.ProfileSetting) {
		return false
	}
	if !o.Qos.matches(other.Qos) {
		return false
	}
	if !util.StringsMatch(o.RuleType, other.RuleType) {
		return false
	}
	if !util.StringsMatch(o.Schedule, other.Schedule) {
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
	if !util.OrderedListsMatch[string](o.SourceImei, other.SourceImei) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceImsi, other.SourceImsi) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceNwSlice, other.SourceNwSlice) {
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
	if !util.OrderedListsMatch[string](o.To, other.To) {
		return false
	}
	if !util.StringsMatch(o.Uuid, other.Uuid) {
		return false
	}

	return true
}

func (o *ProfileSetting) matches(other *ProfileSetting) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Group, other.Group) {
		return false
	}
	if !o.Profiles.matches(other.Profiles) {
		return false
	}

	return true
}

func (o *ProfileSettingProfiles) matches(other *ProfileSettingProfiles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FileBlocking, other.FileBlocking) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Gtp, other.Gtp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Sctp, other.Sctp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Spyware, other.Spyware) {
		return false
	}
	if !util.OrderedListsMatch[string](o.UrlFiltering, other.UrlFiltering) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Virus, other.Virus) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vulnerability, other.Vulnerability) {
		return false
	}
	if !util.OrderedListsMatch[string](o.WildfireAnalysis, other.WildfireAnalysis) {
		return false
	}

	return true
}

func (o *Qos) matches(other *Qos) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Marking.matches(other.Marking) {
		return false
	}

	return true
}

func (o *QosMarking) matches(other *QosMarking) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.FollowC2sFlow.matches(other.FollowC2sFlow) {
		return false
	}
	if !util.StringsMatch(o.IpDscp, other.IpDscp) {
		return false
	}
	if !util.StringsMatch(o.IpPrecedence, other.IpPrecedence) {
		return false
	}

	return true
}

func (o *QosMarkingFollowC2sFlow) matches(other *QosMarkingFollowC2sFlow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
func (o *Entry) EntryUuid() *string {
	return o.Uuid
}

func (o *Entry) SetEntryUuid(uuid *string) {
	o.Uuid = uuid
}
