package security

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
	Suffix = []string{"security", "rules"}
)

// Entry is the normalized object.
//
// Targets is a map where the key is the serial number of the target device and
// the value is a list of specific vsys on that device.  The list of vsys is
// nil if all vsys on that device should be included or if the device is a
// virtual firewall (and thus only has vsys1).
type Entry struct {
	Name                            string
	Uuid                            string
	SourceZones                     []string // unordered
	DestinationZones                []string // unordered
	SourceAddresses                 []string // unordered
	SourceUsers                     []string // unordered
	DestinationAddresses            []string // unordered
	Services                        []string // unordered
	Categories                      []string // unordered
	Applications                    []string // unordered
	SourceDevices                   []string // unordered
	DestinationDevices              []string // unordered
	Schedule                        *string
	Tags                            []string // ordered
	NegateSource                    *bool
	NegateDestination               *bool
	Disabled                        *bool
	Description                     *string
	GroupTag                        *string
	Action                          string
	IcmpUnreachable                 *bool
	Type                            *string
	DisableServerResponseInspection *bool
	LogSetting                      *string
	LogStart                        *bool
	LogEnd                          *bool
	ProfileSettings                 *ProfileSettingsObject
	Qos                             *QosObject
	Targets                         map[string][]string
	NegateTarget                    *bool
	DisableInspect                  *bool // 10.2.4, 10.2.0? : disable-inspect

	Misc map[string][]generic.Xml
}

type ProfileSettingsObject struct {
	Groups   []string
	Profiles *ProfilesObject
}

type ProfilesObject struct {
	UrlFilteringProfiles     []string
	DataFilteringProfiles    []string
	FileBlockingProfiles     []string
	WildfireAnalysisProfiles []string
	AntiVirusProfiles        []string
	AntiSpywareProfiles      []string
	VulnerabilityProfiles    []string
}

type QosObject struct {
	IpDscp                   *string
	IpPrecedence             *string
	FollowClientToServerFlow any
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
	if v == "uuid" || v == "Uuid" {
		return obj.Uuid, nil
	}
	if v == "source_zones" || v == "SourceZones" {
		return obj.SourceZones, nil
	}
	if v == "source_zones|LENGTH" || v == "SourceZones|LENGTH" {
		return len(obj.SourceZones), nil
	}
	if v == "destination_zones" || v == "DestinationZones" {
		return obj.DestinationZones, nil
	}
	if v == "destination_zones|LENGTH" || v == "DestinationZones|LENGTH" {
		return len(obj.DestinationZones), nil
	}
	if v == "source_addresses" || v == "SourceAddresses" {
		return obj.SourceAddresses, nil
	}
	if v == "source_addresses|LENGTH" || v == "SourceAddresses|LENGTH" {
		return len(obj.SourceAddresses), nil
	}
	if v == "source_users" || v == "SourceUsers" {
		return obj.SourceUsers, nil
	}
	if v == "source_users|LENGTH" || v == "SourceUsers|LENGTH" {
		return len(obj.SourceUsers), nil
	}
	if v == "destination_addresses" || v == "DestinationAddresses" {
		return obj.DestinationAddresses, nil
	}
	if v == "destination_addresses|LENGTH" || v == "DestinationAddresses|LENGTH" {
		return len(obj.DestinationAddresses), nil
	}
	if v == "services" || v == "Services" {
		return obj.Services, nil
	}
	if v == "services|LENGTH" || v == "Services|LENGTH" {
		return len(obj.Services), nil
	}
	if v == "categories" || v == "Categories" {
		return obj.Categories, nil
	}
	if v == "categories|LENGTH" || v == "Categories|LENGTH" {
		return len(obj.Categories), nil
	}
	if v == "applications" || v == "Applications" {
		return obj.Applications, nil
	}
	if v == "applications|LENGTH" || v == "Applications|LENGTH" {
		return len(obj.Applications), nil
	}
	if v == "source_devices" || v == "SourceDevices" {
		return obj.SourceDevices, nil
	}
	if v == "source_devices|LENGTH" || v == "SourceDevices|LENGTH" {
		return len(obj.SourceDevices), nil
	}
	if v == "destination_devices" || v == "DestinationDevices" {
		return obj.DestinationDevices, nil
	}
	if v == "destination_devices|LENGTH" || v == "DestinationDevices|LENGTH" {
		return len(obj.DestinationDevices), nil
	}
	if v == "schedule" || v == "Schedule" {
		return obj.Schedule, nil
	}
	if v == "tags" || v == "Tags" {
		return obj.Tags, nil
	}
	if v == "tags|LENGTH" || v == "Tags|LENGTH" {
		return len(obj.Tags), nil
	}
	if v == "negate_source" || v == "NegateSource" {
		return obj.NegateSource, nil
	}
	if v == "negate_destination" || v == "NegateDestination" {
		return obj.NegateDestination, nil
	}
	if v == "disabled" || v == "Disabled" {
		return obj.Disabled, nil
	}
	if v == "description" || v == "Description" {
		return obj.Description, nil
	}
	if v == "group_tag" || v == "GroupTag" {
		return obj.GroupTag, nil
	}
	if v == "action" || v == "Action" {
		return obj.Action, nil
	}
	if v == "icmp_unreachable" || v == "IcmpUnreachable" {
		return obj.IcmpUnreachable, nil
	}
	if v == "type" || v == "Type" {
		return obj.Type, nil
	}
	if v == "disable_server_response_inspection" || v == "DisableServerResponseInspection" {
		return obj.DisableServerResponseInspection, nil
	}
	if v == "log_setting" || v == "LogSetting" {
		return obj.LogSetting, nil
	}
	if v == "log_start" || v == "LogStart" {
		return obj.LogStart, nil
	}
	if v == "log_end" || v == "LogEnd" {
		return obj.LogEnd, nil
	}
	if v == "profile_settings" || v == "ProfileSettings" {
		return obj.ProfileSettings != nil, nil
	}
	if strings.HasPrefix(v, "profile_settings.") || strings.HasPrefix(v, "ProfileSettings.") {
		if obj.ProfileSettings == nil {
			return nil, fmt.Errorf("profile_settings is nil")
		}
		obj := obj.ProfileSettings
		for _, chk := range []string{"profile_settings.", "ProfileSettings."} {
			if strings.HasPrefix(v, chk) {
				v = v[len(chk):]
				break
			}
		}

		if v == "groups" || v == "Groups" {
			return obj.Groups, nil
		}
		if v == "groups|LENGTH" || v == "Groups|LENGTH" {
			return len(obj.Groups), nil
		}
		if v == "profiles" || v == "Profiles" {
			return obj.Profiles != nil, nil
		}
		if strings.HasPrefix(v, "profiles.") || strings.HasPrefix(v, "Profiles.") {
			if obj.Profiles == nil {
				return nil, fmt.Errorf("profiles is nil")
			}
			obj := obj.Profiles
			for _, chk := range []string{"profiles.", "Profiles."} {
				if strings.HasPrefix(v, chk) {
					v = v[len(chk):]
					break
				}
			}

			if v == "url_filtering_profiles" || v == "UrlFilteringProfiles" {
				return obj.UrlFilteringProfiles, nil
			}
			if v == "url_filtering_profiles|LENGTH" || v == "UrlFilteringProfiles|LENGTH" {
				return len(obj.UrlFilteringProfiles), nil
			}
			if v == "data_filtering_profiles" || v == "DataFilteringProfiles" {
				return obj.DataFilteringProfiles, nil
			}
			if v == "data_filtering_profiles|LENGTH" || v == "DataFilteringProfiles|LENGTH" {
				return len(obj.DataFilteringProfiles), nil
			}
			if v == "file_blocking_profiles" || v == "FileBlockingProfiles" {
				return obj.FileBlockingProfiles, nil
			}
			if v == "file_blocking_profiles|LENGTH" || v == "FileBlockingProfiles|LENGTH" {
				return len(obj.FileBlockingProfiles), nil
			}
			if v == "wildfire_analysis_profiles" || v == "WildfireAnalysisProfiles" {
				return obj.WildfireAnalysisProfiles, nil
			}
			if v == "wildfire_analysis_profiles|LENGTH" || v == "WildfireAnalysisProfiles|LENGTH" {
				return len(obj.WildfireAnalysisProfiles), nil
			}
			if v == "anti_virus_profiles" || v == "AntiVirusProfiles" {
				return obj.AntiVirusProfiles, nil
			}
			if v == "anti_virus_profiles|LENGTH" || v == "AntiVirusProfiles|LENGTH" {
				return len(obj.AntiVirusProfiles), nil
			}
			if v == "anti_spyware_profiles" || v == "AntiSpywareProfiles" {
				return obj.AntiSpywareProfiles, nil
			}
			if v == "anti_spyware_profiles|LENGTH" || v == "AntiSpywareProfiles|LENGTH" {
				return len(obj.AntiSpywareProfiles), nil
			}
			if v == "vulnerability_profiles" || v == "VulnerabilityProfiles" {
				return obj.VulnerabilityProfiles, nil
			}
			if v == "vulnerability_profiles|LENGTH" || v == "VulnerabilityProfiles|LENGTH" {
				return len(obj.VulnerabilityProfiles), nil
			}
		}
	}
	if v == "qos" || v == "Qos" {
		return obj.Qos != nil, nil
	}
	if strings.HasPrefix(v, "qos.") || strings.HasPrefix(v, "Qos.") {
		if obj.Qos == nil {
			return nil, fmt.Errorf("qos is nil")
		}
		obj := obj.Qos
		for _, chk := range []string{"qos.", "Qos."} {
			if strings.HasPrefix(v, chk) {
				v = v[len(chk):]
				break
			}
		}

		if v == "ip_dscp" || v == "IpDscp" {
			return obj.IpDscp, nil
		}
		if v == "ip_precedence" || v == "IpPrecedence" {
			return obj.IpPrecedence, nil
		}
		if v == "follow_client_to_server_flow" || v == "FollowClientToServerFlow" {
			return obj.FollowClientToServerFlow != nil, nil
		}
	}
	if v == "negate_target" || v == "NegateTarget" {
		return e.NegateTarget, nil
	}
	if v == "disable_inspect" || v == "DisableInspect" {
		return e.DisableInspect, nil
	}

	return nil, fmt.Errorf("unknown field: %s", value)
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	/*
	   if vn.Gte(version.Number{10, 2, 0, ""}) {
	       return Entry2Specify, &Entry2Container{}, nil
	   }
	*/

	return Entry1Specify, &Entry1Container{}, nil
}

func Entry1Specify(o Entry) (any, error) {
	ans := Entry1{Misc: o.Misc["Entry"]}
	ans.Name = o.Name
	ans.Uuid = o.Uuid
	ans.SourceZones = util.StrToMem(o.SourceZones)
	ans.DestinationZones = util.StrToMem(o.DestinationZones)
	ans.SourceAddresses = util.StrToMem(o.SourceAddresses)
	ans.SourceUsers = util.StrToMem(o.SourceUsers)
	ans.DestinationAddresses = util.StrToMem(o.DestinationAddresses)
	ans.Services = util.StrToMem(o.Services)
	ans.Categories = util.StrToMem(o.Categories)
	ans.Applications = util.StrToMem(o.Applications)
	ans.SourceDevices = util.StrToMem(o.SourceDevices)
	ans.DestinationDevices = util.StrToMem(o.DestinationDevices)
	ans.Schedule = o.Schedule
	ans.Tags = util.StrToMem(o.Tags)
	if o.NegateSource != nil {
		ans.NegateSource = util.String(util.YesNo(*o.NegateSource))
	}
	if o.NegateDestination != nil {
		ans.NegateDestination = util.String(util.YesNo(*o.NegateDestination))
	}
	if o.Disabled != nil {
		ans.Disabled = util.String(util.YesNo(*o.Disabled))
	}
	ans.Description = o.Description
	ans.GroupTag = o.GroupTag
	ans.Action = o.Action
	if o.IcmpUnreachable != nil {
		ans.IcmpUnreachable = util.String(util.YesNo(*o.IcmpUnreachable))
	}
	ans.Type = o.Type
	if o.DisableServerResponseInspection != nil {
		ans.Options = &secOptions{Misc: o.Misc["secOptions"]}
		ans.Options.DisableServerResponseInspection = util.String(util.YesNo(*o.DisableServerResponseInspection))
	}
	ans.LogSetting = o.LogSetting
	if o.LogStart != nil {
		ans.LogStart = util.String(util.YesNo(*o.LogStart))
	}
	if o.LogEnd != nil {
		ans.LogEnd = util.String(util.YesNo(*o.LogEnd))
	}
	if o.ProfileSettings != nil {
		ans.ProfileSettings = &profileSettings{Misc: o.Misc["profileSettings"]}
		ans.ProfileSettings.Groups = util.StrToMem(o.ProfileSettings.Groups)
		if o.ProfileSettings.Profiles != nil {
			ans.ProfileSettings.Profiles = &profileSettingsProfiles{Misc: o.Misc["profileSettingsProfiles"]}
			ans.ProfileSettings.Profiles.UrlFilteringProfiles = util.StrToMem(o.ProfileSettings.Profiles.UrlFilteringProfiles)
			ans.ProfileSettings.Profiles.DataFilteringProfiles = util.StrToMem(o.ProfileSettings.Profiles.DataFilteringProfiles)
			ans.ProfileSettings.Profiles.FileBlockingProfiles = util.StrToMem(o.ProfileSettings.Profiles.FileBlockingProfiles)
			ans.ProfileSettings.Profiles.WildfireAnalysisProfiles = util.StrToMem(o.ProfileSettings.Profiles.WildfireAnalysisProfiles)
			ans.ProfileSettings.Profiles.AntiVirusProfiles = util.StrToMem(o.ProfileSettings.Profiles.AntiVirusProfiles)
			ans.ProfileSettings.Profiles.AntiSpywareProfiles = util.StrToMem(o.ProfileSettings.Profiles.AntiSpywareProfiles)
			ans.ProfileSettings.Profiles.VulnerabilityProfiles = util.StrToMem(o.ProfileSettings.Profiles.VulnerabilityProfiles)
		}
	}
	if o.Qos != nil {
		ans.Qos = &qos{Misc: o.Misc["qos"]}
		ans.Qos.Marking = &qosMarking{Misc: o.Misc["qosMarking"]}
		ans.Qos.Marking.IpDscp = o.Qos.IpDscp
		ans.Qos.Marking.IpPrecedence = o.Qos.IpPrecedence
		if o.Qos.FollowClientToServerFlow != nil {
			ans.Qos.Marking.FollowClientToServerFlow = ""
		}
	}

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
		var1.Misc["Entry"] = var0.Misc
		var1.Name = var0.Name
		var1.Uuid = var0.Uuid
		var1.SourceZones = util.MemToStr(var0.SourceZones)
		var1.SourceUsers = util.MemToStr(var0.SourceUsers)
		var1.SourceAddresses = util.MemToStr(var0.SourceAddresses)
		var1.DestinationZones = util.MemToStr(var0.DestinationZones)
		var1.DestinationAddresses = util.MemToStr(var0.DestinationAddresses)
		var1.Services = util.MemToStr(var0.Services)
		var1.Categories = util.MemToStr(var0.Categories)
		var1.Applications = util.MemToStr(var0.Applications)
		var1.SourceDevices = util.MemToStr(var0.SourceDevices)
		var1.DestinationDevices = util.MemToStr(var0.DestinationDevices)
		var1.Schedule = var0.Schedule
		var1.Tags = util.MemToStr(var0.Tags)
		if var0.NegateSource != nil {
			var5 := util.AsBool(*var0.NegateSource)
			var1.NegateSource = &var5
		}
		if var0.NegateDestination != nil {
			var5 := util.AsBool(*var0.NegateDestination)
			var1.NegateDestination = &var5
		}
		if var0.Disabled != nil {
			var5 := util.AsBool(*var0.Disabled)
			var1.Disabled = &var5
		}
		var1.Description = var0.Description
		var1.GroupTag = var0.GroupTag
		var1.Action = var0.Action
		if var0.IcmpUnreachable != nil {
			var5 := util.AsBool(*var0.IcmpUnreachable)
			var1.IcmpUnreachable = &var5
		}
		var1.Type = var0.Type
		if var0.Options != nil {
			var1.Misc["secOptions"] = var0.Options.Misc
			if var0.Options.DisableServerResponseInspection != nil {
				var5 := util.AsBool(*var0.Options.DisableServerResponseInspection)
				var1.DisableServerResponseInspection = &var5
			}
		}
		var1.LogSetting = var0.LogSetting
		if var0.LogStart != nil {
			var5 := util.AsBool(*var0.LogStart)
			var1.LogStart = &var5
		}
		if var0.LogEnd != nil {
			var5 := util.AsBool(*var0.LogEnd)
			var1.LogEnd = &var5
		}
		if var0.ProfileSettings != nil {
			var1.ProfileSettings = &ProfileSettingsObject{}
			var1.Misc["profileSettings"] = var0.ProfileSettings.Misc
			var1.ProfileSettings.Groups = util.MemToStr(var0.ProfileSettings.Groups)
			if var0.ProfileSettings.Profiles != nil {
				var1.ProfileSettings.Profiles = &ProfilesObject{}
				var1.Misc["profileSettingsProfiles"] = var0.ProfileSettings.Profiles.Misc
				var1.ProfileSettings.Profiles.UrlFilteringProfiles = util.MemToStr(var0.ProfileSettings.Profiles.UrlFilteringProfiles)
				var1.ProfileSettings.Profiles.DataFilteringProfiles = util.MemToStr(var0.ProfileSettings.Profiles.DataFilteringProfiles)
				var1.ProfileSettings.Profiles.FileBlockingProfiles = util.MemToStr(var0.ProfileSettings.Profiles.FileBlockingProfiles)
				var1.ProfileSettings.Profiles.WildfireAnalysisProfiles = util.MemToStr(var0.ProfileSettings.Profiles.WildfireAnalysisProfiles)
				var1.ProfileSettings.Profiles.AntiVirusProfiles = util.MemToStr(var0.ProfileSettings.Profiles.AntiVirusProfiles)
				var1.ProfileSettings.Profiles.AntiSpywareProfiles = util.MemToStr(var0.ProfileSettings.Profiles.AntiSpywareProfiles)
				var1.ProfileSettings.Profiles.VulnerabilityProfiles = util.MemToStr(var0.ProfileSettings.Profiles.VulnerabilityProfiles)
			}
		}
		// qos
		if var0.Qos != nil {
			var1.Misc["qos"] = var0.Qos.Misc
			if var0.Qos.Marking != nil {
				var1.Qos = &QosObject{}
				var1.Misc["qosMarking"] = var0.Qos.Marking.Misc
				var1.Qos.IpDscp = var0.Qos.Marking.IpDscp
				var1.Qos.IpPrecedence = var0.Qos.Marking.IpPrecedence
				var1.Qos.FollowClientToServerFlow = var0.Qos.Marking.FollowClientToServerFlow
			}
		}

		ans = append(ans, var1)
	}

	return ans, nil
}

type Entry1 struct {
	XMLName              xml.Name         `xml:"entry"`
	Name                 string           `xml:"name,attr"`
	Uuid                 string           `xml:"uuid,attr,omitempty"`
	SourceZones          *util.MemberType `xml:"from"`
	DestinationZones     *util.MemberType `xml:"to"`
	SourceAddresses      *util.MemberType `xml:"source"`
	SourceUsers          *util.MemberType `xml:"source-user"`
	DestinationAddresses *util.MemberType `xml:"destination"`
	Services             *util.MemberType `xml:"service"`
	Categories           *util.MemberType `xml:"category"`
	Applications         *util.MemberType `xml:"application"`
	SourceDevices        *util.MemberType `xml:"source-hip"`
	DestinationDevices   *util.MemberType `xml:"destination-hip"`
	Schedule             *string          `xml:"schedule,omitempty"`
	Tags                 *util.MemberType `xml:"tag"`
	NegateSource         *string          `xml:"negate-source,omitempty"`      // bool
	NegateDestination    *string          `xml:"negate-destination,omitempty"` // bool
	Disabled             *string          `xml:"disabled,omitempty"`           // bool
	Description          *string          `xml:"description,omitempty"`
	GroupTag             *string          `xml:"group-tag,omitempty"`
	Action               string           `xml:"action"`
	IcmpUnreachable      *string          `xml:"icmp-unreachable,omitempty"` // bool
	Type                 *string          `xml:"rule-type,omitempty"`
	Options              *secOptions      `xml:"option"`
	LogSetting           *string          `xml:"log-setting,omitempty"`
	LogStart             *string          `xml:"log-start,omitempty"` // bool
	LogEnd               *string          `xml:"log-end,omitempty"`   // bool, default is true
	ProfileSettings      *profileSettings `xml:"profile-setting,omitempty"`
	Qos                  *qos             `xml:"qos,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local Entry1
	yes := util.YesNo(true)
	ans := local{
		LogEnd: &yes,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = Entry1(ans)
	return nil
}

type secOptions struct {
	DisableServerResponseInspection *string `xml:"disable-server-response-inspection,omitempty"` // bool

	Misc []generic.Xml `xml:",any"`
}

type profileSettings struct { // One of.
	Groups   *util.MemberType         `xml:"group,omitempty"`
	Profiles *profileSettingsProfiles `xml:"profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type profileSettingsProfiles struct {
	UrlFilteringProfiles     *util.MemberType `xml:"url-filtering,omitempty"`
	DataFilteringProfiles    *util.MemberType `xml:"data-filtering,omitempty"`
	FileBlockingProfiles     *util.MemberType `xml:"file-blocking,omitempty"`
	WildfireAnalysisProfiles *util.MemberType `xml:"wildfire-analysis,omitempty"`
	AntiVirusProfiles        *util.MemberType `xml:"virus,omitempty"`
	AntiSpywareProfiles      *util.MemberType `xml:"spyware,omitempty"`
	VulnerabilityProfiles    *util.MemberType `xml:"vulnerability,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type qos struct {
	Marking *qosMarking `xml:"marking,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type qosMarking struct { // One of.
	IpDscp                   *string `xml:"ip-dscp,omitempty"`
	IpPrecedence             *string `xml:"ip-precedence,omitempty"`
	FollowClientToServerFlow any     `xml:"follow-c2s-flow,omitempty"` // exist bool

	Misc []generic.Xml `xml:",any"`
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	// omit compare: name
	// omit compare: uuid

	if !util.UnorderedListsMatch(a.SourceZones, b.SourceZones) {
		return false
	}

	if !util.UnorderedListsMatch(a.DestinationZones, b.DestinationZones) {
		return false
	}

	if !util.UnorderedListsMatch(a.SourceAddresses, b.SourceAddresses) {
		return false
	}

	if !util.UnorderedListsMatch(a.SourceUsers, b.SourceUsers) {
		return false
	}

	if !util.UnorderedListsMatch(a.DestinationAddresses, b.DestinationAddresses) {
		return false
	}

	if !util.UnorderedListsMatch(a.Services, b.Services) {
		return false
	}

	if !util.UnorderedListsMatch(a.Categories, b.Categories) {
		return false
	}

	if !util.UnorderedListsMatch(a.Applications, b.Applications) {
		return false
	}

	if !util.UnorderedListsMatch(a.SourceDevices, b.SourceDevices) {
		return false
	}

	if !util.UnorderedListsMatch(a.DestinationDevices, b.DestinationDevices) {
		return false
	}

	if !util.StringsMatch(a.Schedule, b.Schedule) {
		return false
	}

	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}

	if !util.BoolsMatch(a.NegateSource, b.NegateSource) {
		return false
	}

	if !util.BoolsMatch(a.NegateDestination, b.NegateDestination) {
		return false
	}

	if !util.BoolsMatch(a.Disabled, b.Disabled) {
		return false
	}

	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}

	if !util.StringsMatch(a.GroupTag, b.GroupTag) {
		return false
	}

	if a.Action != b.Action {
		return false
	}

	if !util.BoolsMatch(a.IcmpUnreachable, b.IcmpUnreachable) {
		return false
	}

	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}

	if !util.BoolsMatch(a.DisableServerResponseInspection, b.DisableServerResponseInspection) {
		return false
	}

	if !util.BoolsMatch(a.LogStart, b.LogStart) {
		return false
	}

	if !util.BoolsMatch(a.LogEnd, b.LogEnd) {
		return false
	}

	// profile settings
	if a.ProfileSettings == nil && b.ProfileSettings == nil {
	} else if a.ProfileSettings == nil || b.ProfileSettings == nil {
		return false
	} else {
		if !util.OrderedListsMatch(a.ProfileSettings.Groups, b.ProfileSettings.Groups) {
			return false
		}

		if a.ProfileSettings.Profiles == nil && b.ProfileSettings.Profiles == nil {
		} else if a.ProfileSettings.Profiles == nil || b.ProfileSettings.Profiles == nil {
			return false
		} else {
			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.UrlFilteringProfiles, b.ProfileSettings.Profiles.UrlFilteringProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.DataFilteringProfiles, b.ProfileSettings.Profiles.DataFilteringProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.FileBlockingProfiles, b.ProfileSettings.Profiles.FileBlockingProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.WildfireAnalysisProfiles, b.ProfileSettings.Profiles.WildfireAnalysisProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.AntiVirusProfiles, b.ProfileSettings.Profiles.AntiVirusProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.AntiSpywareProfiles, b.ProfileSettings.Profiles.AntiSpywareProfiles) {
				return false
			}

			if !util.OrderedListsMatch(a.ProfileSettings.Profiles.VulnerabilityProfiles, b.ProfileSettings.Profiles.VulnerabilityProfiles) {
				return false
			}
		}
	}

	if a.Qos == nil && b.Qos == nil {
	} else if a.Qos == nil || b.Qos == nil {
		return false
	} else {
		if !util.StringsMatch(a.Qos.IpDscp, b.Qos.IpDscp) {
			return false
		}
		if !util.StringsMatch(a.Qos.IpPrecedence, b.Qos.IpPrecedence) {
			return false
		}
		if !util.AnysMatch(a.Qos.FollowClientToServerFlow, b.Qos.FollowClientToServerFlow) {
			return false
		}
	}

	if !util.BoolsMatch(a.DisableInspect, b.DisableInspect) {
		return false
	}

	return true
}
