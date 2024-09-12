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
	Applications                    []string
	Categories                      []string
	Description                     *string
	DestinationAddresses            []string
	DestinationHips                 []string
	DestinationZones                []string
	DisableServerResponseInspection *bool
	Disabled                        *bool
	IcmpUnreachable                 *bool
	LogEnd                          *bool
	LogSetting                      *string
	LogStart                        *bool
	NegateDestination               *bool
	NegateSource                    *bool
	ProfileSetting                  *ProfileSetting
	RuleType                        *string
	Services                        []string
	SourceAddresses                 []string
	SourceHips                      []string
	SourceUsers                     []string
	SourceZones                     []string
	Tags                            []string
	Uuid                            *string

	Misc map[string][]generic.Xml
}

type ProfileSetting struct {
	Group    *string
	Profiles *ProfileSettingProfiles
}
type ProfileSettingProfiles struct {
	DataFiltering    []string
	FileBlocking     []string
	Spyware          []string
	UrlFiltering     []string
	Virus            []string
	Vulnerability    []string
	WildfireAnalysis []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                         xml.Name           `xml:"entry"`
	Name                            string             `xml:"name,attr"`
	Action                          *string            `xml:"action,omitempty"`
	Applications                    *util.MemberType   `xml:"application,omitempty"`
	Categories                      *util.MemberType   `xml:"category,omitempty"`
	Description                     *string            `xml:"description,omitempty"`
	DestinationAddresses            *util.MemberType   `xml:"destination,omitempty"`
	DestinationHips                 *util.MemberType   `xml:"destination-hip,omitempty"`
	DestinationZones                *util.MemberType   `xml:"to,omitempty"`
	DisableServerResponseInspection *string            `xml:"option>disable-server-response-inspection,omitempty"`
	Disabled                        *string            `xml:"disabled,omitempty"`
	IcmpUnreachable                 *string            `xml:"icmp-unreachable,omitempty"`
	LogEnd                          *string            `xml:"log-end,omitempty"`
	LogSetting                      *string            `xml:"log-setting,omitempty"`
	LogStart                        *string            `xml:"log-start,omitempty"`
	NegateDestination               *string            `xml:"negate-destination,omitempty"`
	NegateSource                    *string            `xml:"negate-source,omitempty"`
	ProfileSetting                  *ProfileSettingXml `xml:"profile-setting,omitempty"`
	RuleType                        *string            `xml:"rule-type,omitempty"`
	Services                        *util.MemberType   `xml:"service,omitempty"`
	SourceAddresses                 *util.MemberType   `xml:"source,omitempty"`
	SourceHips                      *util.MemberType   `xml:"source-hip,omitempty"`
	SourceUsers                     *util.MemberType   `xml:"source-user,omitempty"`
	SourceZones                     *util.MemberType   `xml:"from,omitempty"`
	Tags                            *util.MemberType   `xml:"tag,omitempty"`
	Uuid                            *string            `xml:"uuid,attr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type ProfileSettingXml struct {
	Group    *string                    `xml:"group,omitempty"`
	Profiles *ProfileSettingProfilesXml `xml:"profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProfileSettingProfilesXml struct {
	DataFiltering    *util.MemberType `xml:"data-filtering,omitempty"`
	FileBlocking     *util.MemberType `xml:"file-blocking,omitempty"`
	Spyware          *util.MemberType `xml:"spyware,omitempty"`
	UrlFiltering     *util.MemberType `xml:"url-filtering,omitempty"`
	Virus            *util.MemberType `xml:"virus,omitempty"`
	Vulnerability    *util.MemberType `xml:"vulnerability,omitempty"`
	WildfireAnalysis *util.MemberType `xml:"wildfire-analysis,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "action" || v == "Action" {
		return e.Action, nil
	}
	if v == "applications" || v == "Applications" {
		return e.Applications, nil
	}
	if v == "applications|LENGTH" || v == "Applications|LENGTH" {
		return int64(len(e.Applications)), nil
	}
	if v == "categories" || v == "Categories" {
		return e.Categories, nil
	}
	if v == "categories|LENGTH" || v == "Categories|LENGTH" {
		return int64(len(e.Categories)), nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "destination_addresses" || v == "DestinationAddresses" {
		return e.DestinationAddresses, nil
	}
	if v == "destination_addresses|LENGTH" || v == "DestinationAddresses|LENGTH" {
		return int64(len(e.DestinationAddresses)), nil
	}
	if v == "destination_hips" || v == "DestinationHips" {
		return e.DestinationHips, nil
	}
	if v == "destination_hips|LENGTH" || v == "DestinationHips|LENGTH" {
		return int64(len(e.DestinationHips)), nil
	}
	if v == "destination_zones" || v == "DestinationZones" {
		return e.DestinationZones, nil
	}
	if v == "destination_zones|LENGTH" || v == "DestinationZones|LENGTH" {
		return int64(len(e.DestinationZones)), nil
	}
	if v == "disable_server_response_inspection" || v == "DisableServerResponseInspection" {
		return e.DisableServerResponseInspection, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
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
	if v == "rule_type" || v == "RuleType" {
		return e.RuleType, nil
	}
	if v == "services" || v == "Services" {
		return e.Services, nil
	}
	if v == "services|LENGTH" || v == "Services|LENGTH" {
		return int64(len(e.Services)), nil
	}
	if v == "source_addresses" || v == "SourceAddresses" {
		return e.SourceAddresses, nil
	}
	if v == "source_addresses|LENGTH" || v == "SourceAddresses|LENGTH" {
		return int64(len(e.SourceAddresses)), nil
	}
	if v == "source_hips" || v == "SourceHips" {
		return e.SourceHips, nil
	}
	if v == "source_hips|LENGTH" || v == "SourceHips|LENGTH" {
		return int64(len(e.SourceHips)), nil
	}
	if v == "source_users" || v == "SourceUsers" {
		return e.SourceUsers, nil
	}
	if v == "source_users|LENGTH" || v == "SourceUsers|LENGTH" {
		return int64(len(e.SourceUsers)), nil
	}
	if v == "source_zones" || v == "SourceZones" {
		return e.SourceZones, nil
	}
	if v == "source_zones|LENGTH" || v == "SourceZones|LENGTH" {
		return int64(len(e.SourceZones)), nil
	}
	if v == "tags" || v == "Tags" {
		return e.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return int64(len(e.Tags)), nil
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
	entry.Applications = util.StrToMem(o.Applications)
	entry.Categories = util.StrToMem(o.Categories)
	entry.Description = o.Description
	entry.DestinationAddresses = util.StrToMem(o.DestinationAddresses)
	entry.DestinationHips = util.StrToMem(o.DestinationHips)
	entry.DestinationZones = util.StrToMem(o.DestinationZones)
	entry.DisableServerResponseInspection = util.YesNo(o.DisableServerResponseInspection, nil)
	entry.Disabled = util.YesNo(o.Disabled, nil)
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
			nestedProfileSetting.Group = o.ProfileSetting.Group
		}
		if o.ProfileSetting.Profiles != nil {
			nestedProfileSetting.Profiles = &ProfileSettingProfilesXml{}
			if _, ok := o.Misc["ProfileSettingProfiles"]; ok {
				nestedProfileSetting.Profiles.Misc = o.Misc["ProfileSettingProfiles"]
			}
			if o.ProfileSetting.Profiles.WildfireAnalysis != nil {
				nestedProfileSetting.Profiles.WildfireAnalysis = util.StrToMem(o.ProfileSetting.Profiles.WildfireAnalysis)
			}
			if o.ProfileSetting.Profiles.DataFiltering != nil {
				nestedProfileSetting.Profiles.DataFiltering = util.StrToMem(o.ProfileSetting.Profiles.DataFiltering)
			}
			if o.ProfileSetting.Profiles.Virus != nil {
				nestedProfileSetting.Profiles.Virus = util.StrToMem(o.ProfileSetting.Profiles.Virus)
			}
			if o.ProfileSetting.Profiles.Spyware != nil {
				nestedProfileSetting.Profiles.Spyware = util.StrToMem(o.ProfileSetting.Profiles.Spyware)
			}
			if o.ProfileSetting.Profiles.Vulnerability != nil {
				nestedProfileSetting.Profiles.Vulnerability = util.StrToMem(o.ProfileSetting.Profiles.Vulnerability)
			}
			if o.ProfileSetting.Profiles.UrlFiltering != nil {
				nestedProfileSetting.Profiles.UrlFiltering = util.StrToMem(o.ProfileSetting.Profiles.UrlFiltering)
			}
			if o.ProfileSetting.Profiles.FileBlocking != nil {
				nestedProfileSetting.Profiles.FileBlocking = util.StrToMem(o.ProfileSetting.Profiles.FileBlocking)
			}
		}
	}
	entry.ProfileSetting = nestedProfileSetting

	entry.RuleType = o.RuleType
	entry.Services = util.StrToMem(o.Services)
	entry.SourceAddresses = util.StrToMem(o.SourceAddresses)
	entry.SourceHips = util.StrToMem(o.SourceHips)
	entry.SourceUsers = util.StrToMem(o.SourceUsers)
	entry.SourceZones = util.StrToMem(o.SourceZones)
	entry.Tags = util.StrToMem(o.Tags)
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
		entry.Applications = util.MemToStr(o.Applications)
		entry.Categories = util.MemToStr(o.Categories)
		entry.Description = o.Description
		entry.DestinationAddresses = util.MemToStr(o.DestinationAddresses)
		entry.DestinationHips = util.MemToStr(o.DestinationHips)
		entry.DestinationZones = util.MemToStr(o.DestinationZones)
		entry.DisableServerResponseInspection = util.AsBool(o.DisableServerResponseInspection, nil)
		entry.Disabled = util.AsBool(o.Disabled, nil)
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
				nestedProfileSetting.Group = o.ProfileSetting.Group
			}
			if o.ProfileSetting.Profiles != nil {
				nestedProfileSetting.Profiles = &ProfileSettingProfiles{}
				if o.ProfileSetting.Profiles.Misc != nil {
					entry.Misc["ProfileSettingProfiles"] = o.ProfileSetting.Profiles.Misc
				}
				if o.ProfileSetting.Profiles.Vulnerability != nil {
					nestedProfileSetting.Profiles.Vulnerability = util.MemToStr(o.ProfileSetting.Profiles.Vulnerability)
				}
				if o.ProfileSetting.Profiles.UrlFiltering != nil {
					nestedProfileSetting.Profiles.UrlFiltering = util.MemToStr(o.ProfileSetting.Profiles.UrlFiltering)
				}
				if o.ProfileSetting.Profiles.FileBlocking != nil {
					nestedProfileSetting.Profiles.FileBlocking = util.MemToStr(o.ProfileSetting.Profiles.FileBlocking)
				}
				if o.ProfileSetting.Profiles.WildfireAnalysis != nil {
					nestedProfileSetting.Profiles.WildfireAnalysis = util.MemToStr(o.ProfileSetting.Profiles.WildfireAnalysis)
				}
				if o.ProfileSetting.Profiles.DataFiltering != nil {
					nestedProfileSetting.Profiles.DataFiltering = util.MemToStr(o.ProfileSetting.Profiles.DataFiltering)
				}
				if o.ProfileSetting.Profiles.Virus != nil {
					nestedProfileSetting.Profiles.Virus = util.MemToStr(o.ProfileSetting.Profiles.Virus)
				}
				if o.ProfileSetting.Profiles.Spyware != nil {
					nestedProfileSetting.Profiles.Spyware = util.MemToStr(o.ProfileSetting.Profiles.Spyware)
				}
			}
		}
		entry.ProfileSetting = nestedProfileSetting

		entry.RuleType = o.RuleType
		entry.Services = util.MemToStr(o.Services)
		entry.SourceAddresses = util.MemToStr(o.SourceAddresses)
		entry.SourceHips = util.MemToStr(o.SourceHips)
		entry.SourceUsers = util.MemToStr(o.SourceUsers)
		entry.SourceZones = util.MemToStr(o.SourceZones)
		entry.Tags = util.MemToStr(o.Tags)
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
	if !util.OrderedListsMatch(a.Applications, b.Applications) {
		return false
	}
	if !util.OrderedListsMatch(a.Categories, b.Categories) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.DestinationAddresses, b.DestinationAddresses) {
		return false
	}
	if !util.OrderedListsMatch(a.DestinationHips, b.DestinationHips) {
		return false
	}
	if !util.OrderedListsMatch(a.DestinationZones, b.DestinationZones) {
		return false
	}
	if !util.BoolsMatch(a.DisableServerResponseInspection, b.DisableServerResponseInspection) {
		return false
	}
	if !util.BoolsMatch(a.Disabled, b.Disabled) {
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
	if !util.StringsMatch(a.RuleType, b.RuleType) {
		return false
	}
	if !util.OrderedListsMatch(a.Services, b.Services) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceAddresses, b.SourceAddresses) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceHips, b.SourceHips) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceUsers, b.SourceUsers) {
		return false
	}
	if !util.OrderedListsMatch(a.SourceZones, b.SourceZones) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.Uuid, b.Uuid) {
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
	if !util.OrderedListsMatch(a.Vulnerability, b.Vulnerability) {
		return false
	}
	if !util.OrderedListsMatch(a.UrlFiltering, b.UrlFiltering) {
		return false
	}
	if !util.OrderedListsMatch(a.FileBlocking, b.FileBlocking) {
		return false
	}
	if !util.OrderedListsMatch(a.WildfireAnalysis, b.WildfireAnalysis) {
		return false
	}
	if !util.OrderedListsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.OrderedListsMatch(a.Virus, b.Virus) {
		return false
	}
	if !util.OrderedListsMatch(a.Spyware, b.Spyware) {
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
	if !util.StringsMatch(a.Group, b.Group) {
		return false
	}
	if !matchProfileSettingProfiles(a.Profiles, b.Profiles) {
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
