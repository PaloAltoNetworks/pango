package defaultsecurity

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
	suffix = []string{"default-security-rules", "rules", "$name"}
)

type Entry struct {
	Name            string
	Action          *string
	GroupTag        *string
	IcmpUnreachable *bool
	LogEnd          *bool
	LogSetting      *string
	LogStart        *bool
	ProfileSetting  *ProfileSetting
	Tag             []string
	Uuid            *string
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type ProfileSetting struct {
	Group          []string
	Profiles       *ProfileSettingProfiles
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	MiscAttributes   []xml.Attr
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
	XMLName         xml.Name           `xml:"entry"`
	Name            string             `xml:"name,attr"`
	Action          *string            `xml:"action,omitempty"`
	GroupTag        *string            `xml:"group-tag,omitempty"`
	IcmpUnreachable *string            `xml:"icmp-unreachable,omitempty"`
	LogEnd          *string            `xml:"log-end,omitempty"`
	LogSetting      *string            `xml:"log-setting,omitempty"`
	LogStart        *string            `xml:"log-start,omitempty"`
	ProfileSetting  *profileSettingXml `xml:"profile-setting,omitempty"`
	Tag             *util.MemberType   `xml:"tag,omitempty"`
	Uuid            *string            `xml:"uuid,attr,omitempty"`
	Misc            []generic.Xml      `xml:",any"`
	MiscAttributes  []xml.Attr         `xml:",any,attr"`
}
type profileSettingXml struct {
	Group          *util.MemberType           `xml:"group,omitempty"`
	Profiles       *profileSettingProfilesXml `xml:"profiles,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
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
	MiscAttributes   []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Action = s.Action
	o.GroupTag = s.GroupTag
	o.IcmpUnreachable = util.YesNo(s.IcmpUnreachable, nil)
	o.LogEnd = util.YesNo(s.LogEnd, nil)
	o.LogSetting = s.LogSetting
	o.LogStart = util.YesNo(s.LogStart, nil)
	if s.ProfileSetting != nil {
		var obj profileSettingXml
		obj.MarshalFromObject(*s.ProfileSetting)
		o.ProfileSetting = &obj
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	o.Uuid = s.Uuid
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var profileSettingVal *ProfileSetting
	if o.ProfileSetting != nil {
		obj, err := o.ProfileSetting.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		profileSettingVal = obj
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}

	result := &Entry{
		Name:            o.Name,
		Action:          o.Action,
		GroupTag:        o.GroupTag,
		IcmpUnreachable: util.AsBool(o.IcmpUnreachable, nil),
		LogEnd:          util.AsBool(o.LogEnd, nil),
		LogSetting:      o.LogSetting,
		LogStart:        util.AsBool(o.LogStart, nil),
		ProfileSetting:  profileSettingVal,
		Tag:             tagVal,
		Uuid:            o.Uuid,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
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
	o.MiscAttributes = s.MiscAttributes
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
		Group:          groupVal,
		Profiles:       profilesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
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
	o.MiscAttributes = s.MiscAttributes
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
		MiscAttributes:   o.MiscAttributes,
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
	if v == "profile_setting" || v == "ProfileSetting" {
		return e.ProfileSetting, nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
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
	if !o.ProfileSetting.matches(other.ProfileSetting) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
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
