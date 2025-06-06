package logforwarding

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
	suffix = []string{"log-settings", "profiles", "$name"}
)

type Entry struct {
	Name                       string
	Description                *string
	DisableOverride            *string
	EnhancedApplicationLogging *bool
	MatchList                  []MatchList
	Misc                       []generic.Xml
}
type MatchList struct {
	Name           string
	ActionDesc     *string
	LogType        *string
	Filter         *string
	SendToPanorama *bool
	Quarantine     *bool
	SendSnmptrap   []string
	SendEmail      []string
	SendSyslog     []string
	SendHttp       []string
	Actions        []MatchListActions
	Misc           []generic.Xml
}
type MatchListActions struct {
	Name string
	Type *MatchListActionsType
	Misc []generic.Xml
}
type MatchListActionsType struct {
	Integration *MatchListActionsTypeIntegration
	Tagging     *MatchListActionsTypeTagging
	Misc        []generic.Xml
}
type MatchListActionsTypeIntegration struct {
	Action *string
	Misc   []generic.Xml
}
type MatchListActionsTypeTagging struct {
	Target       *string
	Action       *string
	Timeout      *int64
	Registration *MatchListActionsTypeTaggingRegistration
	Tags         []string
	Misc         []generic.Xml
}
type MatchListActionsTypeTaggingRegistration struct {
	Localhost *MatchListActionsTypeTaggingRegistrationLocalhost
	Panorama  *MatchListActionsTypeTaggingRegistrationPanorama
	Remote    *MatchListActionsTypeTaggingRegistrationRemote
	Misc      []generic.Xml
}
type MatchListActionsTypeTaggingRegistrationLocalhost struct {
	Misc []generic.Xml
}
type MatchListActionsTypeTaggingRegistrationPanorama struct {
	Misc []generic.Xml
}
type MatchListActionsTypeTaggingRegistrationRemote struct {
	HttpProfile *string
	Misc        []generic.Xml
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
	XMLName                    xml.Name               `xml:"entry"`
	Name                       string                 `xml:"name,attr"`
	Description                *string                `xml:"description,omitempty"`
	DisableOverride            *string                `xml:"disable-override,omitempty"`
	EnhancedApplicationLogging *string                `xml:"enhanced-application-logging,omitempty"`
	MatchList                  *matchListContainerXml `xml:"match-list,omitempty"`
	Misc                       []generic.Xml          `xml:",any"`
}
type matchListContainerXml struct {
	Entries []matchListXml `xml:"entry"`
}
type matchListXml struct {
	XMLName        xml.Name                      `xml:"entry"`
	Name           string                        `xml:"name,attr"`
	ActionDesc     *string                       `xml:"action-desc,omitempty"`
	LogType        *string                       `xml:"log-type,omitempty"`
	Filter         *string                       `xml:"filter,omitempty"`
	SendToPanorama *string                       `xml:"send-to-panorama,omitempty"`
	Quarantine     *string                       `xml:"quarantine,omitempty"`
	SendSnmptrap   *util.MemberType              `xml:"send-snmptrap,omitempty"`
	SendEmail      *util.MemberType              `xml:"send-email,omitempty"`
	SendSyslog     *util.MemberType              `xml:"send-syslog,omitempty"`
	SendHttp       *util.MemberType              `xml:"send-http,omitempty"`
	Actions        *matchListActionsContainerXml `xml:"actions,omitempty"`
	Misc           []generic.Xml                 `xml:",any"`
}
type matchListActionsContainerXml struct {
	Entries []matchListActionsXml `xml:"entry"`
}
type matchListActionsXml struct {
	XMLName xml.Name                 `xml:"entry"`
	Name    string                   `xml:"name,attr"`
	Type    *matchListActionsTypeXml `xml:"type,omitempty"`
	Misc    []generic.Xml            `xml:",any"`
}
type matchListActionsTypeXml struct {
	Integration *matchListActionsTypeIntegrationXml `xml:"integration,omitempty"`
	Tagging     *matchListActionsTypeTaggingXml     `xml:"tagging,omitempty"`
	Misc        []generic.Xml                       `xml:",any"`
}
type matchListActionsTypeIntegrationXml struct {
	Action *string       `xml:"action,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type matchListActionsTypeTaggingXml struct {
	Target       *string                                     `xml:"target,omitempty"`
	Action       *string                                     `xml:"action,omitempty"`
	Timeout      *int64                                      `xml:"timeout,omitempty"`
	Registration *matchListActionsTypeTaggingRegistrationXml `xml:"registration,omitempty"`
	Tags         *util.MemberType                            `xml:"tags,omitempty"`
	Misc         []generic.Xml                               `xml:",any"`
}
type matchListActionsTypeTaggingRegistrationXml struct {
	Localhost *matchListActionsTypeTaggingRegistrationLocalhostXml `xml:"localhost,omitempty"`
	Panorama  *matchListActionsTypeTaggingRegistrationPanoramaXml  `xml:"panorama,omitempty"`
	Remote    *matchListActionsTypeTaggingRegistrationRemoteXml    `xml:"remote,omitempty"`
	Misc      []generic.Xml                                        `xml:",any"`
}
type matchListActionsTypeTaggingRegistrationLocalhostXml struct {
	Misc []generic.Xml `xml:",any"`
}
type matchListActionsTypeTaggingRegistrationPanoramaXml struct {
	Misc []generic.Xml `xml:",any"`
}
type matchListActionsTypeTaggingRegistrationRemoteXml struct {
	HttpProfile *string       `xml:"http-profile,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	o.EnhancedApplicationLogging = util.YesNo(s.EnhancedApplicationLogging, nil)
	if s.MatchList != nil {
		var objs []matchListXml
		for _, elt := range s.MatchList {
			var obj matchListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MatchList = &matchListContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var matchListVal []MatchList
	if o.MatchList != nil {
		for _, elt := range o.MatchList.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			matchListVal = append(matchListVal, *obj)
		}
	}

	result := &Entry{
		Name:                       o.Name,
		Description:                o.Description,
		DisableOverride:            o.DisableOverride,
		EnhancedApplicationLogging: util.AsBool(o.EnhancedApplicationLogging, nil),
		MatchList:                  matchListVal,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *matchListXml) MarshalFromObject(s MatchList) {
	o.Name = s.Name
	o.ActionDesc = s.ActionDesc
	o.LogType = s.LogType
	o.Filter = s.Filter
	o.SendToPanorama = util.YesNo(s.SendToPanorama, nil)
	o.Quarantine = util.YesNo(s.Quarantine, nil)
	if s.SendSnmptrap != nil {
		o.SendSnmptrap = util.StrToMem(s.SendSnmptrap)
	}
	if s.SendEmail != nil {
		o.SendEmail = util.StrToMem(s.SendEmail)
	}
	if s.SendSyslog != nil {
		o.SendSyslog = util.StrToMem(s.SendSyslog)
	}
	if s.SendHttp != nil {
		o.SendHttp = util.StrToMem(s.SendHttp)
	}
	if s.Actions != nil {
		var objs []matchListActionsXml
		for _, elt := range s.Actions {
			var obj matchListActionsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Actions = &matchListActionsContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o matchListXml) UnmarshalToObject() (*MatchList, error) {
	var sendSnmptrapVal []string
	if o.SendSnmptrap != nil {
		sendSnmptrapVal = util.MemToStr(o.SendSnmptrap)
	}
	var sendEmailVal []string
	if o.SendEmail != nil {
		sendEmailVal = util.MemToStr(o.SendEmail)
	}
	var sendSyslogVal []string
	if o.SendSyslog != nil {
		sendSyslogVal = util.MemToStr(o.SendSyslog)
	}
	var sendHttpVal []string
	if o.SendHttp != nil {
		sendHttpVal = util.MemToStr(o.SendHttp)
	}
	var actionsVal []MatchListActions
	if o.Actions != nil {
		for _, elt := range o.Actions.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			actionsVal = append(actionsVal, *obj)
		}
	}

	result := &MatchList{
		Name:           o.Name,
		ActionDesc:     o.ActionDesc,
		LogType:        o.LogType,
		Filter:         o.Filter,
		SendToPanorama: util.AsBool(o.SendToPanorama, nil),
		Quarantine:     util.AsBool(o.Quarantine, nil),
		SendSnmptrap:   sendSnmptrapVal,
		SendEmail:      sendEmailVal,
		SendSyslog:     sendSyslogVal,
		SendHttp:       sendHttpVal,
		Actions:        actionsVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *matchListActionsXml) MarshalFromObject(s MatchListActions) {
	o.Name = s.Name
	if s.Type != nil {
		var obj matchListActionsTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o matchListActionsXml) UnmarshalToObject() (*MatchListActions, error) {
	var typeVal *MatchListActionsType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &MatchListActions{
		Name: o.Name,
		Type: typeVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeXml) MarshalFromObject(s MatchListActionsType) {
	if s.Integration != nil {
		var obj matchListActionsTypeIntegrationXml
		obj.MarshalFromObject(*s.Integration)
		o.Integration = &obj
	}
	if s.Tagging != nil {
		var obj matchListActionsTypeTaggingXml
		obj.MarshalFromObject(*s.Tagging)
		o.Tagging = &obj
	}
	o.Misc = s.Misc
}

func (o matchListActionsTypeXml) UnmarshalToObject() (*MatchListActionsType, error) {
	var integrationVal *MatchListActionsTypeIntegration
	if o.Integration != nil {
		obj, err := o.Integration.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		integrationVal = obj
	}
	var taggingVal *MatchListActionsTypeTagging
	if o.Tagging != nil {
		obj, err := o.Tagging.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		taggingVal = obj
	}

	result := &MatchListActionsType{
		Integration: integrationVal,
		Tagging:     taggingVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeIntegrationXml) MarshalFromObject(s MatchListActionsTypeIntegration) {
	o.Action = s.Action
	o.Misc = s.Misc
}

func (o matchListActionsTypeIntegrationXml) UnmarshalToObject() (*MatchListActionsTypeIntegration, error) {

	result := &MatchListActionsTypeIntegration{
		Action: o.Action,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeTaggingXml) MarshalFromObject(s MatchListActionsTypeTagging) {
	o.Target = s.Target
	o.Action = s.Action
	o.Timeout = s.Timeout
	if s.Registration != nil {
		var obj matchListActionsTypeTaggingRegistrationXml
		obj.MarshalFromObject(*s.Registration)
		o.Registration = &obj
	}
	if s.Tags != nil {
		o.Tags = util.StrToMem(s.Tags)
	}
	o.Misc = s.Misc
}

func (o matchListActionsTypeTaggingXml) UnmarshalToObject() (*MatchListActionsTypeTagging, error) {
	var registrationVal *MatchListActionsTypeTaggingRegistration
	if o.Registration != nil {
		obj, err := o.Registration.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		registrationVal = obj
	}
	var tagsVal []string
	if o.Tags != nil {
		tagsVal = util.MemToStr(o.Tags)
	}

	result := &MatchListActionsTypeTagging{
		Target:       o.Target,
		Action:       o.Action,
		Timeout:      o.Timeout,
		Registration: registrationVal,
		Tags:         tagsVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeTaggingRegistrationXml) MarshalFromObject(s MatchListActionsTypeTaggingRegistration) {
	if s.Localhost != nil {
		var obj matchListActionsTypeTaggingRegistrationLocalhostXml
		obj.MarshalFromObject(*s.Localhost)
		o.Localhost = &obj
	}
	if s.Panorama != nil {
		var obj matchListActionsTypeTaggingRegistrationPanoramaXml
		obj.MarshalFromObject(*s.Panorama)
		o.Panorama = &obj
	}
	if s.Remote != nil {
		var obj matchListActionsTypeTaggingRegistrationRemoteXml
		obj.MarshalFromObject(*s.Remote)
		o.Remote = &obj
	}
	o.Misc = s.Misc
}

func (o matchListActionsTypeTaggingRegistrationXml) UnmarshalToObject() (*MatchListActionsTypeTaggingRegistration, error) {
	var localhostVal *MatchListActionsTypeTaggingRegistrationLocalhost
	if o.Localhost != nil {
		obj, err := o.Localhost.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localhostVal = obj
	}
	var panoramaVal *MatchListActionsTypeTaggingRegistrationPanorama
	if o.Panorama != nil {
		obj, err := o.Panorama.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		panoramaVal = obj
	}
	var remoteVal *MatchListActionsTypeTaggingRegistrationRemote
	if o.Remote != nil {
		obj, err := o.Remote.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		remoteVal = obj
	}

	result := &MatchListActionsTypeTaggingRegistration{
		Localhost: localhostVal,
		Panorama:  panoramaVal,
		Remote:    remoteVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeTaggingRegistrationLocalhostXml) MarshalFromObject(s MatchListActionsTypeTaggingRegistrationLocalhost) {
	o.Misc = s.Misc
}

func (o matchListActionsTypeTaggingRegistrationLocalhostXml) UnmarshalToObject() (*MatchListActionsTypeTaggingRegistrationLocalhost, error) {

	result := &MatchListActionsTypeTaggingRegistrationLocalhost{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeTaggingRegistrationPanoramaXml) MarshalFromObject(s MatchListActionsTypeTaggingRegistrationPanorama) {
	o.Misc = s.Misc
}

func (o matchListActionsTypeTaggingRegistrationPanoramaXml) UnmarshalToObject() (*MatchListActionsTypeTaggingRegistrationPanorama, error) {

	result := &MatchListActionsTypeTaggingRegistrationPanorama{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *matchListActionsTypeTaggingRegistrationRemoteXml) MarshalFromObject(s MatchListActionsTypeTaggingRegistrationRemote) {
	o.HttpProfile = s.HttpProfile
	o.Misc = s.Misc
}

func (o matchListActionsTypeTaggingRegistrationRemoteXml) UnmarshalToObject() (*MatchListActionsTypeTaggingRegistrationRemote, error) {

	result := &MatchListActionsTypeTaggingRegistrationRemote{
		HttpProfile: o.HttpProfile,
		Misc:        o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "enhanced_application_logging" || v == "EnhancedApplicationLogging" {
		return e.EnhancedApplicationLogging, nil
	}
	if v == "match_list" || v == "MatchList" {
		return e.MatchList, nil
	}
	if v == "match_list|LENGTH" || v == "MatchList|LENGTH" {
		return int64(len(e.MatchList)), nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(o.EnhancedApplicationLogging, other.EnhancedApplicationLogging) {
		return false
	}
	if len(o.MatchList) != len(other.MatchList) {
		return false
	}
	for idx := range o.MatchList {
		if !o.MatchList[idx].matches(&other.MatchList[idx]) {
			return false
		}
	}

	return true
}

func (o *MatchList) matches(other *MatchList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.ActionDesc, other.ActionDesc) {
		return false
	}
	if !util.StringsMatch(o.LogType, other.LogType) {
		return false
	}
	if !util.StringsMatch(o.Filter, other.Filter) {
		return false
	}
	if !util.BoolsMatch(o.SendToPanorama, other.SendToPanorama) {
		return false
	}
	if !util.BoolsMatch(o.Quarantine, other.Quarantine) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSnmptrap, other.SendSnmptrap) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendEmail, other.SendEmail) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSyslog, other.SendSyslog) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendHttp, other.SendHttp) {
		return false
	}
	if len(o.Actions) != len(other.Actions) {
		return false
	}
	for idx := range o.Actions {
		if !o.Actions[idx].matches(&other.Actions[idx]) {
			return false
		}
	}

	return true
}

func (o *MatchListActions) matches(other *MatchListActions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}

	return true
}

func (o *MatchListActionsType) matches(other *MatchListActionsType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Integration.matches(other.Integration) {
		return false
	}
	if !o.Tagging.matches(other.Tagging) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeIntegration) matches(other *MatchListActionsTypeIntegration) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeTagging) matches(other *MatchListActionsTypeTagging) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Target, other.Target) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !o.Registration.matches(other.Registration) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tags, other.Tags) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeTaggingRegistration) matches(other *MatchListActionsTypeTaggingRegistration) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Localhost.matches(other.Localhost) {
		return false
	}
	if !o.Panorama.matches(other.Panorama) {
		return false
	}
	if !o.Remote.matches(other.Remote) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeTaggingRegistrationLocalhost) matches(other *MatchListActionsTypeTaggingRegistrationLocalhost) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeTaggingRegistrationPanorama) matches(other *MatchListActionsTypeTaggingRegistrationPanorama) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *MatchListActionsTypeTaggingRegistrationRemote) matches(other *MatchListActionsTypeTaggingRegistrationRemote) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.HttpProfile, other.HttpProfile) {
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
