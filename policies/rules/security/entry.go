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
	Suffix = []string{"security", "rules"}
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

	Misc map[string][]generic.Xml
}

type ProfileSetting struct {
	Group    []string
	Profiles *ProfileSettingProfiles
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
}
type Qos struct {
	Marking *QosMarking
}
type QosMarking struct {
	FollowC2sFlow *QosMarkingFollowC2sFlow
	IpDscp        *string
	IpPrecedence  *string
}
type QosMarkingFollowC2sFlow struct {
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
	ProfileSetting                  *ProfileSettingXml `xml:"profile-setting,omitempty"`
	Qos                             *QosXml            `xml:"qos,omitempty"`
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
	Target                          *TargetXml         `xml:"target,omitempty"`
	To                              *util.MemberType   `xml:"to,omitempty"`
	Uuid                            *string            `xml:"uuid,attr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProfileSettingXml struct {
	Group    *util.MemberType           `xml:"group,omitempty"`
	Profiles *ProfileSettingProfilesXml `xml:"profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProfileSettingProfilesXml struct {
	DataFiltering    *util.MemberType `xml:"data-filtering,omitempty"`
	FileBlocking     *util.MemberType `xml:"file-blocking,omitempty"`
	Gtp              *util.MemberType `xml:"gtp,omitempty"`
	Sctp             *util.MemberType `xml:"sctp,omitempty"`
	Spyware          *util.MemberType `xml:"spyware,omitempty"`
	UrlFiltering     *util.MemberType `xml:"url-filtering,omitempty"`
	Virus            *util.MemberType `xml:"virus,omitempty"`
	Vulnerability    *util.MemberType `xml:"vulnerability,omitempty"`
	WildfireAnalysis *util.MemberType `xml:"wildfire-analysis,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type QosXml struct {
	Marking *QosMarkingXml `xml:"marking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type QosMarkingXml struct {
	FollowC2sFlow *QosMarkingFollowC2sFlowXml `xml:"follow-c2s-flow,omitempty"`
	IpDscp        *string                     `xml:"ip-dscp,omitempty"`
	IpPrecedence  *string                     `xml:"ip-precedence,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type QosMarkingFollowC2sFlowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TargetXml struct {
	Devices []TargetDevicesXml `xml:"devices>entry,omitempty"`
	Negate  *string            `xml:"negate,omitempty"`
	Tags    *util.MemberType   `xml:"tags,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TargetDevicesXml struct {
	Name string                 `xml:"name,attr"`
	Vsys []TargetDevicesVsysXml `xml:"vsys>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TargetDevicesVsysXml struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
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
	if v == "disable_server_response_inspection" || v == "DisableServerResponseInspection" {
		return e.DisableServerResponseInspection, nil
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
	entry.Application = util.StrToMem(o.Application)
	entry.Category = util.StrToMem(o.Category)
	entry.Description = o.Description
	entry.Destination = util.StrToMem(o.Destination)
	entry.DestinationHip = util.StrToMem(o.DestinationHip)
	entry.DisableInspect = util.YesNo(o.DisableInspect, nil)
	entry.DisableServerResponseInspection = util.YesNo(o.DisableServerResponseInspection, nil)
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.From = util.StrToMem(o.From)
	entry.GroupTag = o.GroupTag
	entry.IcmpUnreachable = util.YesNo(o.IcmpUnreachable, nil)
	entry.LogEnd = util.YesNo(o.LogEnd, nil)
	entry.LogSetting = o.LogSetting
	entry.LogStart = util.YesNo(o.LogStart, nil)
	entry.NegateDestination = util.YesNo(o.NegateDestination, nil)
	entry.NegateSource = util.YesNo(o.NegateSource, nil)
	var nestedProfileSetting *ProfileSettingXml
	if o.ProfileSetting != nil {
		nestedProfileSetting = &ProfileSettingXml{}
		if _, ok := o.Misc["ProfileSetting"]; ok {
			nestedProfileSetting.Misc = o.Misc["ProfileSetting"]
		}
		if o.ProfileSetting.Group != nil {
			nestedProfileSetting.Group = util.StrToMem(o.ProfileSetting.Group)
		}
		if o.ProfileSetting.Profiles != nil {
			nestedProfileSetting.Profiles = &ProfileSettingProfilesXml{}
			if _, ok := o.Misc["ProfileSettingProfiles"]; ok {
				nestedProfileSetting.Profiles.Misc = o.Misc["ProfileSettingProfiles"]
			}
			if o.ProfileSetting.Profiles.DataFiltering != nil {
				nestedProfileSetting.Profiles.DataFiltering = util.StrToMem(o.ProfileSetting.Profiles.DataFiltering)
			}
			if o.ProfileSetting.Profiles.FileBlocking != nil {
				nestedProfileSetting.Profiles.FileBlocking = util.StrToMem(o.ProfileSetting.Profiles.FileBlocking)
			}
			if o.ProfileSetting.Profiles.Gtp != nil {
				nestedProfileSetting.Profiles.Gtp = util.StrToMem(o.ProfileSetting.Profiles.Gtp)
			}
			if o.ProfileSetting.Profiles.Sctp != nil {
				nestedProfileSetting.Profiles.Sctp = util.StrToMem(o.ProfileSetting.Profiles.Sctp)
			}
			if o.ProfileSetting.Profiles.Spyware != nil {
				nestedProfileSetting.Profiles.Spyware = util.StrToMem(o.ProfileSetting.Profiles.Spyware)
			}
			if o.ProfileSetting.Profiles.UrlFiltering != nil {
				nestedProfileSetting.Profiles.UrlFiltering = util.StrToMem(o.ProfileSetting.Profiles.UrlFiltering)
			}
			if o.ProfileSetting.Profiles.Virus != nil {
				nestedProfileSetting.Profiles.Virus = util.StrToMem(o.ProfileSetting.Profiles.Virus)
			}
			if o.ProfileSetting.Profiles.Vulnerability != nil {
				nestedProfileSetting.Profiles.Vulnerability = util.StrToMem(o.ProfileSetting.Profiles.Vulnerability)
			}
			if o.ProfileSetting.Profiles.WildfireAnalysis != nil {
				nestedProfileSetting.Profiles.WildfireAnalysis = util.StrToMem(o.ProfileSetting.Profiles.WildfireAnalysis)
			}
		}
	}
	entry.ProfileSetting = nestedProfileSetting

	var nestedQos *QosXml
	if o.Qos != nil {
		nestedQos = &QosXml{}
		if _, ok := o.Misc["Qos"]; ok {
			nestedQos.Misc = o.Misc["Qos"]
		}
		if o.Qos.Marking != nil {
			nestedQos.Marking = &QosMarkingXml{}
			if _, ok := o.Misc["QosMarking"]; ok {
				nestedQos.Marking.Misc = o.Misc["QosMarking"]
			}
			if o.Qos.Marking.FollowC2sFlow != nil {
				nestedQos.Marking.FollowC2sFlow = &QosMarkingFollowC2sFlowXml{}
				if _, ok := o.Misc["QosMarkingFollowC2sFlow"]; ok {
					nestedQos.Marking.FollowC2sFlow.Misc = o.Misc["QosMarkingFollowC2sFlow"]
				}
			}
			if o.Qos.Marking.IpDscp != nil {
				nestedQos.Marking.IpDscp = o.Qos.Marking.IpDscp
			}
			if o.Qos.Marking.IpPrecedence != nil {
				nestedQos.Marking.IpPrecedence = o.Qos.Marking.IpPrecedence
			}
		}
	}
	entry.Qos = nestedQos

	entry.RuleType = o.RuleType
	entry.Schedule = o.Schedule
	entry.Service = util.StrToMem(o.Service)
	entry.Source = util.StrToMem(o.Source)
	entry.SourceHip = util.StrToMem(o.SourceHip)
	entry.SourceImei = util.StrToMem(o.SourceImei)
	entry.SourceImsi = util.StrToMem(o.SourceImsi)
	entry.SourceNwSlice = util.StrToMem(o.SourceNwSlice)
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
		entry.Application = util.MemToStr(o.Application)
		entry.Category = util.MemToStr(o.Category)
		entry.Description = o.Description
		entry.Destination = util.MemToStr(o.Destination)
		entry.DestinationHip = util.MemToStr(o.DestinationHip)
		entry.DisableInspect = util.AsBool(o.DisableInspect, nil)
		entry.DisableServerResponseInspection = util.AsBool(o.DisableServerResponseInspection, nil)
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.From = util.MemToStr(o.From)
		entry.GroupTag = o.GroupTag
		entry.IcmpUnreachable = util.AsBool(o.IcmpUnreachable, nil)
		entry.LogEnd = util.AsBool(o.LogEnd, nil)
		entry.LogSetting = o.LogSetting
		entry.LogStart = util.AsBool(o.LogStart, nil)
		entry.NegateDestination = util.AsBool(o.NegateDestination, nil)
		entry.NegateSource = util.AsBool(o.NegateSource, nil)
		var nestedProfileSetting *ProfileSetting
		if o.ProfileSetting != nil {
			nestedProfileSetting = &ProfileSetting{}
			if o.ProfileSetting.Misc != nil {
				entry.Misc["ProfileSetting"] = o.ProfileSetting.Misc
			}
			if o.ProfileSetting.Group != nil {
				nestedProfileSetting.Group = util.MemToStr(o.ProfileSetting.Group)
			}
			if o.ProfileSetting.Profiles != nil {
				nestedProfileSetting.Profiles = &ProfileSettingProfiles{}
				if o.ProfileSetting.Profiles.Misc != nil {
					entry.Misc["ProfileSettingProfiles"] = o.ProfileSetting.Profiles.Misc
				}
				if o.ProfileSetting.Profiles.DataFiltering != nil {
					nestedProfileSetting.Profiles.DataFiltering = util.MemToStr(o.ProfileSetting.Profiles.DataFiltering)
				}
				if o.ProfileSetting.Profiles.FileBlocking != nil {
					nestedProfileSetting.Profiles.FileBlocking = util.MemToStr(o.ProfileSetting.Profiles.FileBlocking)
				}
				if o.ProfileSetting.Profiles.Gtp != nil {
					nestedProfileSetting.Profiles.Gtp = util.MemToStr(o.ProfileSetting.Profiles.Gtp)
				}
				if o.ProfileSetting.Profiles.Sctp != nil {
					nestedProfileSetting.Profiles.Sctp = util.MemToStr(o.ProfileSetting.Profiles.Sctp)
				}
				if o.ProfileSetting.Profiles.Spyware != nil {
					nestedProfileSetting.Profiles.Spyware = util.MemToStr(o.ProfileSetting.Profiles.Spyware)
				}
				if o.ProfileSetting.Profiles.UrlFiltering != nil {
					nestedProfileSetting.Profiles.UrlFiltering = util.MemToStr(o.ProfileSetting.Profiles.UrlFiltering)
				}
				if o.ProfileSetting.Profiles.Virus != nil {
					nestedProfileSetting.Profiles.Virus = util.MemToStr(o.ProfileSetting.Profiles.Virus)
				}
				if o.ProfileSetting.Profiles.Vulnerability != nil {
					nestedProfileSetting.Profiles.Vulnerability = util.MemToStr(o.ProfileSetting.Profiles.Vulnerability)
				}
				if o.ProfileSetting.Profiles.WildfireAnalysis != nil {
					nestedProfileSetting.Profiles.WildfireAnalysis = util.MemToStr(o.ProfileSetting.Profiles.WildfireAnalysis)
				}
			}
		}
		entry.ProfileSetting = nestedProfileSetting

		var nestedQos *Qos
		if o.Qos != nil {
			nestedQos = &Qos{}
			if o.Qos.Misc != nil {
				entry.Misc["Qos"] = o.Qos.Misc
			}
			if o.Qos.Marking != nil {
				nestedQos.Marking = &QosMarking{}
				if o.Qos.Marking.Misc != nil {
					entry.Misc["QosMarking"] = o.Qos.Marking.Misc
				}
				if o.Qos.Marking.FollowC2sFlow != nil {
					nestedQos.Marking.FollowC2sFlow = &QosMarkingFollowC2sFlow{}
					if o.Qos.Marking.FollowC2sFlow.Misc != nil {
						entry.Misc["QosMarkingFollowC2sFlow"] = o.Qos.Marking.FollowC2sFlow.Misc
					}
				}
				if o.Qos.Marking.IpDscp != nil {
					nestedQos.Marking.IpDscp = o.Qos.Marking.IpDscp
				}
				if o.Qos.Marking.IpPrecedence != nil {
					nestedQos.Marking.IpPrecedence = o.Qos.Marking.IpPrecedence
				}
			}
		}
		entry.Qos = nestedQos

		entry.RuleType = o.RuleType
		entry.Schedule = o.Schedule
		entry.Service = util.MemToStr(o.Service)
		entry.Source = util.MemToStr(o.Source)
		entry.SourceHip = util.MemToStr(o.SourceHip)
		entry.SourceImei = util.MemToStr(o.SourceImei)
		entry.SourceImsi = util.MemToStr(o.SourceImsi)
		entry.SourceNwSlice = util.MemToStr(o.SourceNwSlice)
		entry.SourceUser = util.MemToStr(o.SourceUser)
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
	if !util.OrderedListsMatch(a.Application, b.Application) {
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
	if !util.BoolsMatch(a.DisableInspect, b.DisableInspect) {
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
	if !util.BoolsMatch(a.IcmpUnreachable, b.IcmpUnreachable) {
		return false
	}
	if !util.BoolsMatch(a.LogEnd, b.LogEnd) {
		return false
	}
	if !util.StringsMatch(a.LogSetting, b.LogSetting) {
		return false
	}
	if !util.BoolsMatch(a.LogStart, b.LogStart) {
		return false
	}
	if !util.BoolsMatch(a.NegateDestination, b.NegateDestination) {
		return false
	}
	if !util.BoolsMatch(a.NegateSource, b.NegateSource) {
		return false
	}
	if !matchProfileSetting(a.ProfileSetting, b.ProfileSetting) {
		return false
	}
	if !matchQos(a.Qos, b.Qos) {
		return false
	}
	if !util.StringsMatch(a.RuleType, b.RuleType) {
		return false
	}
	if !util.StringsMatch(a.Schedule, b.Schedule) {
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
	if !util.OrderedListsMatch(a.SourceImei, b.SourceImei) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceImsi, b.SourceImsi) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceNwSlice, b.SourceNwSlice) {
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
	if !util.StringsMatch(a.Uuid, b.Uuid) {
		return false
	}
	if !util.BoolsMatch(a.DisableServerResponseInspection, b.DisableServerResponseInspection) {
		return false
	}

	return true
}

func matchProfileSettingProfiles(a *ProfileSettingProfiles, b *ProfileSettingProfiles) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.OrderedListsMatch(a.FileBlocking, b.FileBlocking) {
		return false
	}
	if !util.OrderedListsMatch(a.Gtp, b.Gtp) {
		return false
	}
	if !util.OrderedListsMatch(a.Sctp, b.Sctp) {
		return false
	}
	if !util.OrderedListsMatch(a.Spyware, b.Spyware) {
		return false
	}
	if !util.OrderedListsMatch(a.UrlFiltering, b.UrlFiltering) {
		return false
	}
	if !util.OrderedListsMatch(a.Virus, b.Virus) {
		return false
	}
	if !util.OrderedListsMatch(a.Vulnerability, b.Vulnerability) {
		return false
	}
	if !util.OrderedListsMatch(a.WildfireAnalysis, b.WildfireAnalysis) {
		return false
	}
	return true
}
func matchProfileSetting(a *ProfileSetting, b *ProfileSetting) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Group, b.Group) {
		return false
	}
	if !matchProfileSettingProfiles(a.Profiles, b.Profiles) {
		return false
	}
	return true
}
func matchQosMarkingFollowC2sFlow(a *QosMarkingFollowC2sFlow, b *QosMarkingFollowC2sFlow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchQosMarking(a *QosMarking, b *QosMarking) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchQosMarkingFollowC2sFlow(a.FollowC2sFlow, b.FollowC2sFlow) {
		return false
	}
	if !util.StringsMatch(a.IpDscp, b.IpDscp) {
		return false
	}
	if !util.StringsMatch(a.IpPrecedence, b.IpPrecedence) {
		return false
	}
	return true
}
func matchQos(a *Qos, b *Qos) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchQosMarking(a.Marking, b.Marking) {
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
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !matchTargetDevicesVsys(a.Vsys, b.Vsys) {
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
