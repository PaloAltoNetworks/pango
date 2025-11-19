package globalprotect

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
	suffix = []string{"log-settings", "globalprotect", "match-list", "$name"}
)

type Entry struct {
	Name           string
	Actions        []Actions
	Description    *string
	Filter         *string
	Quarantine     *bool
	SendEmail      []string
	SendHttp       []string
	SendSnmptrap   []string
	SendSyslog     []string
	SendToPanorama *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Actions struct {
	Name           string
	Type           *ActionsType
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsType struct {
	Integration    *ActionsTypeIntegration
	Tagging        *ActionsTypeTagging
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeIntegration struct {
	Action         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeTagging struct {
	Target         *string
	Action         *string
	Timeout        *int64
	Registration   *ActionsTypeTaggingRegistration
	Tags           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeTaggingRegistration struct {
	Localhost      *ActionsTypeTaggingRegistrationLocalhost
	Panorama       *ActionsTypeTaggingRegistrationPanorama
	Remote         *ActionsTypeTaggingRegistrationRemote
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeTaggingRegistrationLocalhost struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeTaggingRegistrationPanorama struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionsTypeTaggingRegistrationRemote struct {
	HttpProfile    *string
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
	XMLName        xml.Name             `xml:"entry"`
	Name           string               `xml:"name,attr"`
	Actions        *actionsContainerXml `xml:"actions,omitempty"`
	Description    *string              `xml:"description,omitempty"`
	Filter         *string              `xml:"filter,omitempty"`
	Quarantine     *string              `xml:"quarantine,omitempty"`
	SendEmail      *util.MemberType     `xml:"send-email,omitempty"`
	SendHttp       *util.MemberType     `xml:"send-http,omitempty"`
	SendSnmptrap   *util.MemberType     `xml:"send-snmptrap,omitempty"`
	SendSyslog     *util.MemberType     `xml:"send-syslog,omitempty"`
	SendToPanorama *string              `xml:"send-to-panorama,omitempty"`
	Misc           []generic.Xml        `xml:",any"`
	MiscAttributes []xml.Attr           `xml:",any,attr"`
}
type actionsContainerXml struct {
	Entries []actionsXml `xml:"entry"`
}
type actionsXml struct {
	XMLName        xml.Name        `xml:"entry"`
	Name           string          `xml:"name,attr"`
	Type           *actionsTypeXml `xml:"type,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type actionsTypeXml struct {
	Integration    *actionsTypeIntegrationXml `xml:"integration,omitempty"`
	Tagging        *actionsTypeTaggingXml     `xml:"tagging,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type actionsTypeIntegrationXml struct {
	Action         *string       `xml:"action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type actionsTypeTaggingXml struct {
	Target         *string                            `xml:"target,omitempty"`
	Action         *string                            `xml:"action,omitempty"`
	Timeout        *int64                             `xml:"timeout,omitempty"`
	Registration   *actionsTypeTaggingRegistrationXml `xml:"registration,omitempty"`
	Tags           *util.MemberType                   `xml:"tags,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type actionsTypeTaggingRegistrationXml struct {
	Localhost      *actionsTypeTaggingRegistrationLocalhostXml `xml:"localhost,omitempty"`
	Panorama       *actionsTypeTaggingRegistrationPanoramaXml  `xml:"panorama,omitempty"`
	Remote         *actionsTypeTaggingRegistrationRemoteXml    `xml:"remote,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type actionsTypeTaggingRegistrationLocalhostXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type actionsTypeTaggingRegistrationPanoramaXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type actionsTypeTaggingRegistrationRemoteXml struct {
	HttpProfile    *string       `xml:"http-profile,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Actions != nil {
		var objs []actionsXml
		for _, elt := range s.Actions {
			var obj actionsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Actions = &actionsContainerXml{Entries: objs}
	}
	o.Description = s.Description
	o.Filter = s.Filter
	o.Quarantine = util.YesNo(s.Quarantine, nil)
	if s.SendEmail != nil {
		o.SendEmail = util.StrToMem(s.SendEmail)
	}
	if s.SendHttp != nil {
		o.SendHttp = util.StrToMem(s.SendHttp)
	}
	if s.SendSnmptrap != nil {
		o.SendSnmptrap = util.StrToMem(s.SendSnmptrap)
	}
	if s.SendSyslog != nil {
		o.SendSyslog = util.StrToMem(s.SendSyslog)
	}
	o.SendToPanorama = util.YesNo(s.SendToPanorama, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var actionsVal []Actions
	if o.Actions != nil {
		for _, elt := range o.Actions.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			actionsVal = append(actionsVal, *obj)
		}
	}
	var sendEmailVal []string
	if o.SendEmail != nil {
		sendEmailVal = util.MemToStr(o.SendEmail)
	}
	var sendHttpVal []string
	if o.SendHttp != nil {
		sendHttpVal = util.MemToStr(o.SendHttp)
	}
	var sendSnmptrapVal []string
	if o.SendSnmptrap != nil {
		sendSnmptrapVal = util.MemToStr(o.SendSnmptrap)
	}
	var sendSyslogVal []string
	if o.SendSyslog != nil {
		sendSyslogVal = util.MemToStr(o.SendSyslog)
	}

	result := &Entry{
		Name:           o.Name,
		Actions:        actionsVal,
		Description:    o.Description,
		Filter:         o.Filter,
		Quarantine:     util.AsBool(o.Quarantine, nil),
		SendEmail:      sendEmailVal,
		SendHttp:       sendHttpVal,
		SendSnmptrap:   sendSnmptrapVal,
		SendSyslog:     sendSyslogVal,
		SendToPanorama: util.AsBool(o.SendToPanorama, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsXml) MarshalFromObject(s Actions) {
	o.Name = s.Name
	if s.Type != nil {
		var obj actionsTypeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsXml) UnmarshalToObject() (*Actions, error) {
	var typeVal *ActionsType
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Actions{
		Name:           o.Name,
		Type:           typeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeXml) MarshalFromObject(s ActionsType) {
	if s.Integration != nil {
		var obj actionsTypeIntegrationXml
		obj.MarshalFromObject(*s.Integration)
		o.Integration = &obj
	}
	if s.Tagging != nil {
		var obj actionsTypeTaggingXml
		obj.MarshalFromObject(*s.Tagging)
		o.Tagging = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeXml) UnmarshalToObject() (*ActionsType, error) {
	var integrationVal *ActionsTypeIntegration
	if o.Integration != nil {
		obj, err := o.Integration.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		integrationVal = obj
	}
	var taggingVal *ActionsTypeTagging
	if o.Tagging != nil {
		obj, err := o.Tagging.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		taggingVal = obj
	}

	result := &ActionsType{
		Integration:    integrationVal,
		Tagging:        taggingVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeIntegrationXml) MarshalFromObject(s ActionsTypeIntegration) {
	o.Action = s.Action
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeIntegrationXml) UnmarshalToObject() (*ActionsTypeIntegration, error) {

	result := &ActionsTypeIntegration{
		Action:         o.Action,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeTaggingXml) MarshalFromObject(s ActionsTypeTagging) {
	o.Target = s.Target
	o.Action = s.Action
	o.Timeout = s.Timeout
	if s.Registration != nil {
		var obj actionsTypeTaggingRegistrationXml
		obj.MarshalFromObject(*s.Registration)
		o.Registration = &obj
	}
	if s.Tags != nil {
		o.Tags = util.StrToMem(s.Tags)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeTaggingXml) UnmarshalToObject() (*ActionsTypeTagging, error) {
	var registrationVal *ActionsTypeTaggingRegistration
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

	result := &ActionsTypeTagging{
		Target:         o.Target,
		Action:         o.Action,
		Timeout:        o.Timeout,
		Registration:   registrationVal,
		Tags:           tagsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeTaggingRegistrationXml) MarshalFromObject(s ActionsTypeTaggingRegistration) {
	if s.Localhost != nil {
		var obj actionsTypeTaggingRegistrationLocalhostXml
		obj.MarshalFromObject(*s.Localhost)
		o.Localhost = &obj
	}
	if s.Panorama != nil {
		var obj actionsTypeTaggingRegistrationPanoramaXml
		obj.MarshalFromObject(*s.Panorama)
		o.Panorama = &obj
	}
	if s.Remote != nil {
		var obj actionsTypeTaggingRegistrationRemoteXml
		obj.MarshalFromObject(*s.Remote)
		o.Remote = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeTaggingRegistrationXml) UnmarshalToObject() (*ActionsTypeTaggingRegistration, error) {
	var localhostVal *ActionsTypeTaggingRegistrationLocalhost
	if o.Localhost != nil {
		obj, err := o.Localhost.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localhostVal = obj
	}
	var panoramaVal *ActionsTypeTaggingRegistrationPanorama
	if o.Panorama != nil {
		obj, err := o.Panorama.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		panoramaVal = obj
	}
	var remoteVal *ActionsTypeTaggingRegistrationRemote
	if o.Remote != nil {
		obj, err := o.Remote.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		remoteVal = obj
	}

	result := &ActionsTypeTaggingRegistration{
		Localhost:      localhostVal,
		Panorama:       panoramaVal,
		Remote:         remoteVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeTaggingRegistrationLocalhostXml) MarshalFromObject(s ActionsTypeTaggingRegistrationLocalhost) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeTaggingRegistrationLocalhostXml) UnmarshalToObject() (*ActionsTypeTaggingRegistrationLocalhost, error) {

	result := &ActionsTypeTaggingRegistrationLocalhost{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeTaggingRegistrationPanoramaXml) MarshalFromObject(s ActionsTypeTaggingRegistrationPanorama) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeTaggingRegistrationPanoramaXml) UnmarshalToObject() (*ActionsTypeTaggingRegistrationPanorama, error) {

	result := &ActionsTypeTaggingRegistrationPanorama{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionsTypeTaggingRegistrationRemoteXml) MarshalFromObject(s ActionsTypeTaggingRegistrationRemote) {
	o.HttpProfile = s.HttpProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionsTypeTaggingRegistrationRemoteXml) UnmarshalToObject() (*ActionsTypeTaggingRegistrationRemote, error) {

	result := &ActionsTypeTaggingRegistrationRemote{
		HttpProfile:    o.HttpProfile,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "actions" || v == "Actions" {
		return e.Actions, nil
	}
	if v == "actions|LENGTH" || v == "Actions|LENGTH" {
		return int64(len(e.Actions)), nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "filter" || v == "Filter" {
		return e.Filter, nil
	}
	if v == "quarantine" || v == "Quarantine" {
		return e.Quarantine, nil
	}
	if v == "send_email" || v == "SendEmail" {
		return e.SendEmail, nil
	}
	if v == "send_email|LENGTH" || v == "SendEmail|LENGTH" {
		return int64(len(e.SendEmail)), nil
	}
	if v == "send_http" || v == "SendHttp" {
		return e.SendHttp, nil
	}
	if v == "send_http|LENGTH" || v == "SendHttp|LENGTH" {
		return int64(len(e.SendHttp)), nil
	}
	if v == "send_snmptrap" || v == "SendSnmptrap" {
		return e.SendSnmptrap, nil
	}
	if v == "send_snmptrap|LENGTH" || v == "SendSnmptrap|LENGTH" {
		return int64(len(e.SendSnmptrap)), nil
	}
	if v == "send_syslog" || v == "SendSyslog" {
		return e.SendSyslog, nil
	}
	if v == "send_syslog|LENGTH" || v == "SendSyslog|LENGTH" {
		return int64(len(e.SendSyslog)), nil
	}
	if v == "send_to_panorama" || v == "SendToPanorama" {
		return e.SendToPanorama, nil
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
	if len(o.Actions) != len(other.Actions) {
		return false
	}
	for idx := range o.Actions {
		if !o.Actions[idx].matches(&other.Actions[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.Filter, other.Filter) {
		return false
	}
	if !util.BoolsMatch(o.Quarantine, other.Quarantine) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendEmail, other.SendEmail) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendHttp, other.SendHttp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSnmptrap, other.SendSnmptrap) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SendSyslog, other.SendSyslog) {
		return false
	}
	if !util.BoolsMatch(o.SendToPanorama, other.SendToPanorama) {
		return false
	}

	return true
}

func (o *Actions) matches(other *Actions) bool {
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

func (o *ActionsType) matches(other *ActionsType) bool {
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

func (o *ActionsTypeIntegration) matches(other *ActionsTypeIntegration) bool {
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

func (o *ActionsTypeTagging) matches(other *ActionsTypeTagging) bool {
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

func (o *ActionsTypeTaggingRegistration) matches(other *ActionsTypeTaggingRegistration) bool {
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

func (o *ActionsTypeTaggingRegistrationLocalhost) matches(other *ActionsTypeTaggingRegistrationLocalhost) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ActionsTypeTaggingRegistrationPanorama) matches(other *ActionsTypeTaggingRegistrationPanorama) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ActionsTypeTaggingRegistrationRemote) matches(other *ActionsTypeTaggingRegistrationRemote) bool {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
