package urlfiltering

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"profiles", "url-filtering", "$name"}
)

type Entry struct {
	Name                  string
	Alert                 []string
	Allow                 []string
	Block                 []string
	CloudInlineCat        *bool
	Continue              []string
	CredentialEnforcement *CredentialEnforcement
	Description           *string
	DisableOverride       *string
	EnableContainerPage   *bool
	HttpHeaderInsertion   []HttpHeaderInsertion
	LocalInlineCat        *bool
	LogContainerPageOnly  *bool
	LogHttpHdrReferer     *bool
	LogHttpHdrUserAgent   *bool
	LogHttpHdrXff         *bool
	MlavCategoryException []string
	Override              []string
	SafeSearchEnforcement *bool
	Misc                  []generic.Xml
}
type CredentialEnforcement struct {
	Alert       []string
	Allow       []string
	Block       []string
	Continue    []string
	LogSeverity *string
	Mode        *CredentialEnforcementMode
	Misc        []generic.Xml
}
type CredentialEnforcementMode struct {
	Disabled          *CredentialEnforcementModeDisabled
	DomainCredentials *CredentialEnforcementModeDomainCredentials
	GroupMapping      *string
	IpUser            *CredentialEnforcementModeIpUser
	Misc              []generic.Xml
}
type CredentialEnforcementModeDisabled struct {
	Misc []generic.Xml
}
type CredentialEnforcementModeDomainCredentials struct {
	Misc []generic.Xml
}
type CredentialEnforcementModeIpUser struct {
	Misc []generic.Xml
}
type HttpHeaderInsertion struct {
	Name            string
	DisableOverride *string
	Type            []HttpHeaderInsertionType
	Misc            []generic.Xml
}
type HttpHeaderInsertionType struct {
	Name    string
	Headers []HttpHeaderInsertionTypeHeaders
	Domains []string
	Misc    []generic.Xml
}
type HttpHeaderInsertionTypeHeaders struct {
	Name   string
	Header *string
	Value  *string
	Log    *bool
	Misc   []generic.Xml
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
	XMLName               xml.Name                         `xml:"entry"`
	Name                  string                           `xml:"name,attr"`
	Alert                 *util.MemberType                 `xml:"alert,omitempty"`
	Allow                 *util.MemberType                 `xml:"allow,omitempty"`
	Block                 *util.MemberType                 `xml:"block,omitempty"`
	CloudInlineCat        *string                          `xml:"cloud-inline-cat,omitempty"`
	Continue              *util.MemberType                 `xml:"continue,omitempty"`
	CredentialEnforcement *credentialEnforcementXml        `xml:"credential-enforcement,omitempty"`
	Description           *string                          `xml:"description,omitempty"`
	DisableOverride       *string                          `xml:"disable-override,omitempty"`
	EnableContainerPage   *string                          `xml:"enable-container-page,omitempty"`
	HttpHeaderInsertion   *httpHeaderInsertionContainerXml `xml:"http-header-insertion,omitempty"`
	LocalInlineCat        *string                          `xml:"local-inline-cat,omitempty"`
	LogContainerPageOnly  *string                          `xml:"log-container-page-only,omitempty"`
	LogHttpHdrReferer     *string                          `xml:"log-http-hdr-referer,omitempty"`
	LogHttpHdrUserAgent   *string                          `xml:"log-http-hdr-user-agent,omitempty"`
	LogHttpHdrXff         *string                          `xml:"log-http-hdr-xff,omitempty"`
	MlavCategoryException *util.MemberType                 `xml:"mlav-category-exception,omitempty"`
	Override              *util.MemberType                 `xml:"override,omitempty"`
	SafeSearchEnforcement *string                          `xml:"safe-search-enforcement,omitempty"`
	Misc                  []generic.Xml                    `xml:",any"`
}
type credentialEnforcementXml struct {
	Alert       *util.MemberType              `xml:"alert,omitempty"`
	Allow       *util.MemberType              `xml:"allow,omitempty"`
	Block       *util.MemberType              `xml:"block,omitempty"`
	Continue    *util.MemberType              `xml:"continue,omitempty"`
	LogSeverity *string                       `xml:"log-severity,omitempty"`
	Mode        *credentialEnforcementModeXml `xml:"mode,omitempty"`
	Misc        []generic.Xml                 `xml:",any"`
}
type credentialEnforcementModeXml struct {
	Disabled          *credentialEnforcementModeDisabledXml          `xml:"disabled,omitempty"`
	DomainCredentials *credentialEnforcementModeDomainCredentialsXml `xml:"domain-credentials,omitempty"`
	GroupMapping      *string                                        `xml:"group-mapping,omitempty"`
	IpUser            *credentialEnforcementModeIpUserXml            `xml:"ip-user,omitempty"`
	Misc              []generic.Xml                                  `xml:",any"`
}
type credentialEnforcementModeDisabledXml struct {
	Misc []generic.Xml `xml:",any"`
}
type credentialEnforcementModeDomainCredentialsXml struct {
	Misc []generic.Xml `xml:",any"`
}
type credentialEnforcementModeIpUserXml struct {
	Misc []generic.Xml `xml:",any"`
}
type httpHeaderInsertionContainerXml struct {
	Entries []httpHeaderInsertionXml `xml:"entry"`
}
type httpHeaderInsertionXml struct {
	XMLName         xml.Name                             `xml:"entry"`
	Name            string                               `xml:"name,attr"`
	DisableOverride *string                              `xml:"disable-override,omitempty"`
	Type            *httpHeaderInsertionTypeContainerXml `xml:"type,omitempty"`
	Misc            []generic.Xml                        `xml:",any"`
}
type httpHeaderInsertionTypeContainerXml struct {
	Entries []httpHeaderInsertionTypeXml `xml:"entry"`
}
type httpHeaderInsertionTypeXml struct {
	XMLName xml.Name                                    `xml:"entry"`
	Name    string                                      `xml:"name,attr"`
	Headers *httpHeaderInsertionTypeHeadersContainerXml `xml:"headers,omitempty"`
	Domains *util.MemberType                            `xml:"domains,omitempty"`
	Misc    []generic.Xml                               `xml:",any"`
}
type httpHeaderInsertionTypeHeadersContainerXml struct {
	Entries []httpHeaderInsertionTypeHeadersXml `xml:"entry"`
}
type httpHeaderInsertionTypeHeadersXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Header  *string       `xml:"header,omitempty"`
	Value   *string       `xml:"value,omitempty"`
	Log     *string       `xml:"log,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Alert != nil {
		o.Alert = util.StrToMem(s.Alert)
	}
	if s.Allow != nil {
		o.Allow = util.StrToMem(s.Allow)
	}
	if s.Block != nil {
		o.Block = util.StrToMem(s.Block)
	}
	o.CloudInlineCat = util.YesNo(s.CloudInlineCat, nil)
	if s.Continue != nil {
		o.Continue = util.StrToMem(s.Continue)
	}
	if s.CredentialEnforcement != nil {
		var obj credentialEnforcementXml
		obj.MarshalFromObject(*s.CredentialEnforcement)
		o.CredentialEnforcement = &obj
	}
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	o.EnableContainerPage = util.YesNo(s.EnableContainerPage, nil)
	if s.HttpHeaderInsertion != nil {
		var objs []httpHeaderInsertionXml
		for _, elt := range s.HttpHeaderInsertion {
			var obj httpHeaderInsertionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.HttpHeaderInsertion = &httpHeaderInsertionContainerXml{Entries: objs}
	}
	o.LocalInlineCat = util.YesNo(s.LocalInlineCat, nil)
	o.LogContainerPageOnly = util.YesNo(s.LogContainerPageOnly, nil)
	o.LogHttpHdrReferer = util.YesNo(s.LogHttpHdrReferer, nil)
	o.LogHttpHdrUserAgent = util.YesNo(s.LogHttpHdrUserAgent, nil)
	o.LogHttpHdrXff = util.YesNo(s.LogHttpHdrXff, nil)
	if s.MlavCategoryException != nil {
		o.MlavCategoryException = util.StrToMem(s.MlavCategoryException)
	}
	if s.Override != nil {
		o.Override = util.StrToMem(s.Override)
	}
	o.SafeSearchEnforcement = util.YesNo(s.SafeSearchEnforcement, nil)
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var alertVal []string
	if o.Alert != nil {
		alertVal = util.MemToStr(o.Alert)
	}
	var allowVal []string
	if o.Allow != nil {
		allowVal = util.MemToStr(o.Allow)
	}
	var blockVal []string
	if o.Block != nil {
		blockVal = util.MemToStr(o.Block)
	}
	var continueVal []string
	if o.Continue != nil {
		continueVal = util.MemToStr(o.Continue)
	}
	var credentialEnforcementVal *CredentialEnforcement
	if o.CredentialEnforcement != nil {
		obj, err := o.CredentialEnforcement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		credentialEnforcementVal = obj
	}
	var httpHeaderInsertionVal []HttpHeaderInsertion
	if o.HttpHeaderInsertion != nil {
		for _, elt := range o.HttpHeaderInsertion.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			httpHeaderInsertionVal = append(httpHeaderInsertionVal, *obj)
		}
	}
	var mlavCategoryExceptionVal []string
	if o.MlavCategoryException != nil {
		mlavCategoryExceptionVal = util.MemToStr(o.MlavCategoryException)
	}
	var overrideVal []string
	if o.Override != nil {
		overrideVal = util.MemToStr(o.Override)
	}

	result := &Entry{
		Name:                  o.Name,
		Alert:                 alertVal,
		Allow:                 allowVal,
		Block:                 blockVal,
		CloudInlineCat:        util.AsBool(o.CloudInlineCat, nil),
		Continue:              continueVal,
		CredentialEnforcement: credentialEnforcementVal,
		Description:           o.Description,
		DisableOverride:       o.DisableOverride,
		EnableContainerPage:   util.AsBool(o.EnableContainerPage, nil),
		HttpHeaderInsertion:   httpHeaderInsertionVal,
		LocalInlineCat:        util.AsBool(o.LocalInlineCat, nil),
		LogContainerPageOnly:  util.AsBool(o.LogContainerPageOnly, nil),
		LogHttpHdrReferer:     util.AsBool(o.LogHttpHdrReferer, nil),
		LogHttpHdrUserAgent:   util.AsBool(o.LogHttpHdrUserAgent, nil),
		LogHttpHdrXff:         util.AsBool(o.LogHttpHdrXff, nil),
		MlavCategoryException: mlavCategoryExceptionVal,
		Override:              overrideVal,
		SafeSearchEnforcement: util.AsBool(o.SafeSearchEnforcement, nil),
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *credentialEnforcementXml) MarshalFromObject(s CredentialEnforcement) {
	if s.Alert != nil {
		o.Alert = util.StrToMem(s.Alert)
	}
	if s.Allow != nil {
		o.Allow = util.StrToMem(s.Allow)
	}
	if s.Block != nil {
		o.Block = util.StrToMem(s.Block)
	}
	if s.Continue != nil {
		o.Continue = util.StrToMem(s.Continue)
	}
	o.LogSeverity = s.LogSeverity
	if s.Mode != nil {
		var obj credentialEnforcementModeXml
		obj.MarshalFromObject(*s.Mode)
		o.Mode = &obj
	}
	o.Misc = s.Misc
}

func (o credentialEnforcementXml) UnmarshalToObject() (*CredentialEnforcement, error) {
	var alertVal []string
	if o.Alert != nil {
		alertVal = util.MemToStr(o.Alert)
	}
	var allowVal []string
	if o.Allow != nil {
		allowVal = util.MemToStr(o.Allow)
	}
	var blockVal []string
	if o.Block != nil {
		blockVal = util.MemToStr(o.Block)
	}
	var continueVal []string
	if o.Continue != nil {
		continueVal = util.MemToStr(o.Continue)
	}
	var modeVal *CredentialEnforcementMode
	if o.Mode != nil {
		obj, err := o.Mode.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		modeVal = obj
	}

	result := &CredentialEnforcement{
		Alert:       alertVal,
		Allow:       allowVal,
		Block:       blockVal,
		Continue:    continueVal,
		LogSeverity: o.LogSeverity,
		Mode:        modeVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *credentialEnforcementModeXml) MarshalFromObject(s CredentialEnforcementMode) {
	if s.Disabled != nil {
		var obj credentialEnforcementModeDisabledXml
		obj.MarshalFromObject(*s.Disabled)
		o.Disabled = &obj
	}
	if s.DomainCredentials != nil {
		var obj credentialEnforcementModeDomainCredentialsXml
		obj.MarshalFromObject(*s.DomainCredentials)
		o.DomainCredentials = &obj
	}
	o.GroupMapping = s.GroupMapping
	if s.IpUser != nil {
		var obj credentialEnforcementModeIpUserXml
		obj.MarshalFromObject(*s.IpUser)
		o.IpUser = &obj
	}
	o.Misc = s.Misc
}

func (o credentialEnforcementModeXml) UnmarshalToObject() (*CredentialEnforcementMode, error) {
	var disabledVal *CredentialEnforcementModeDisabled
	if o.Disabled != nil {
		obj, err := o.Disabled.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		disabledVal = obj
	}
	var domainCredentialsVal *CredentialEnforcementModeDomainCredentials
	if o.DomainCredentials != nil {
		obj, err := o.DomainCredentials.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		domainCredentialsVal = obj
	}
	var ipUserVal *CredentialEnforcementModeIpUser
	if o.IpUser != nil {
		obj, err := o.IpUser.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipUserVal = obj
	}

	result := &CredentialEnforcementMode{
		Disabled:          disabledVal,
		DomainCredentials: domainCredentialsVal,
		GroupMapping:      o.GroupMapping,
		IpUser:            ipUserVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *credentialEnforcementModeDisabledXml) MarshalFromObject(s CredentialEnforcementModeDisabled) {
	o.Misc = s.Misc
}

func (o credentialEnforcementModeDisabledXml) UnmarshalToObject() (*CredentialEnforcementModeDisabled, error) {

	result := &CredentialEnforcementModeDisabled{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *credentialEnforcementModeDomainCredentialsXml) MarshalFromObject(s CredentialEnforcementModeDomainCredentials) {
	o.Misc = s.Misc
}

func (o credentialEnforcementModeDomainCredentialsXml) UnmarshalToObject() (*CredentialEnforcementModeDomainCredentials, error) {

	result := &CredentialEnforcementModeDomainCredentials{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *credentialEnforcementModeIpUserXml) MarshalFromObject(s CredentialEnforcementModeIpUser) {
	o.Misc = s.Misc
}

func (o credentialEnforcementModeIpUserXml) UnmarshalToObject() (*CredentialEnforcementModeIpUser, error) {

	result := &CredentialEnforcementModeIpUser{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *httpHeaderInsertionXml) MarshalFromObject(s HttpHeaderInsertion) {
	o.Name = s.Name
	o.DisableOverride = s.DisableOverride
	if s.Type != nil {
		var objs []httpHeaderInsertionTypeXml
		for _, elt := range s.Type {
			var obj httpHeaderInsertionTypeXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Type = &httpHeaderInsertionTypeContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o httpHeaderInsertionXml) UnmarshalToObject() (*HttpHeaderInsertion, error) {
	var typeVal []HttpHeaderInsertionType
	if o.Type != nil {
		for _, elt := range o.Type.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			typeVal = append(typeVal, *obj)
		}
	}

	result := &HttpHeaderInsertion{
		Name:            o.Name,
		DisableOverride: o.DisableOverride,
		Type:            typeVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *httpHeaderInsertionTypeXml) MarshalFromObject(s HttpHeaderInsertionType) {
	o.Name = s.Name
	if s.Headers != nil {
		var objs []httpHeaderInsertionTypeHeadersXml
		for _, elt := range s.Headers {
			var obj httpHeaderInsertionTypeHeadersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Headers = &httpHeaderInsertionTypeHeadersContainerXml{Entries: objs}
	}
	if s.Domains != nil {
		o.Domains = util.StrToMem(s.Domains)
	}
	o.Misc = s.Misc
}

func (o httpHeaderInsertionTypeXml) UnmarshalToObject() (*HttpHeaderInsertionType, error) {
	var headersVal []HttpHeaderInsertionTypeHeaders
	if o.Headers != nil {
		for _, elt := range o.Headers.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			headersVal = append(headersVal, *obj)
		}
	}
	var domainsVal []string
	if o.Domains != nil {
		domainsVal = util.MemToStr(o.Domains)
	}

	result := &HttpHeaderInsertionType{
		Name:    o.Name,
		Headers: headersVal,
		Domains: domainsVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *httpHeaderInsertionTypeHeadersXml) MarshalFromObject(s HttpHeaderInsertionTypeHeaders) {
	o.Name = s.Name
	o.Header = s.Header
	o.Value = s.Value
	o.Log = util.YesNo(s.Log, nil)
	o.Misc = s.Misc
}

func (o httpHeaderInsertionTypeHeadersXml) UnmarshalToObject() (*HttpHeaderInsertionTypeHeaders, error) {

	result := &HttpHeaderInsertionTypeHeaders{
		Name:   o.Name,
		Header: o.Header,
		Value:  o.Value,
		Log:    util.AsBool(o.Log, nil),
		Misc:   o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "alert" || v == "Alert" {
		return e.Alert, nil
	}
	if v == "alert|LENGTH" || v == "Alert|LENGTH" {
		return int64(len(e.Alert)), nil
	}
	if v == "allow" || v == "Allow" {
		return e.Allow, nil
	}
	if v == "allow|LENGTH" || v == "Allow|LENGTH" {
		return int64(len(e.Allow)), nil
	}
	if v == "block" || v == "Block" {
		return e.Block, nil
	}
	if v == "block|LENGTH" || v == "Block|LENGTH" {
		return int64(len(e.Block)), nil
	}
	if v == "cloud_inline_cat" || v == "CloudInlineCat" {
		return e.CloudInlineCat, nil
	}
	if v == "continue" || v == "Continue" {
		return e.Continue, nil
	}
	if v == "continue|LENGTH" || v == "Continue|LENGTH" {
		return int64(len(e.Continue)), nil
	}
	if v == "credential_enforcement" || v == "CredentialEnforcement" {
		return e.CredentialEnforcement, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "enable_container_page" || v == "EnableContainerPage" {
		return e.EnableContainerPage, nil
	}
	if v == "http_header_insertion" || v == "HttpHeaderInsertion" {
		return e.HttpHeaderInsertion, nil
	}
	if v == "http_header_insertion|LENGTH" || v == "HttpHeaderInsertion|LENGTH" {
		return int64(len(e.HttpHeaderInsertion)), nil
	}
	if v == "local_inline_cat" || v == "LocalInlineCat" {
		return e.LocalInlineCat, nil
	}
	if v == "log_container_page_only" || v == "LogContainerPageOnly" {
		return e.LogContainerPageOnly, nil
	}
	if v == "log_http_hdr_referer" || v == "LogHttpHdrReferer" {
		return e.LogHttpHdrReferer, nil
	}
	if v == "log_http_hdr_user_agent" || v == "LogHttpHdrUserAgent" {
		return e.LogHttpHdrUserAgent, nil
	}
	if v == "log_http_hdr_xff" || v == "LogHttpHdrXff" {
		return e.LogHttpHdrXff, nil
	}
	if v == "mlav_category_exception" || v == "MlavCategoryException" {
		return e.MlavCategoryException, nil
	}
	if v == "mlav_category_exception|LENGTH" || v == "MlavCategoryException|LENGTH" {
		return int64(len(e.MlavCategoryException)), nil
	}
	if v == "override" || v == "Override" {
		return e.Override, nil
	}
	if v == "override|LENGTH" || v == "Override|LENGTH" {
		return int64(len(e.Override)), nil
	}
	if v == "safe_search_enforcement" || v == "SafeSearchEnforcement" {
		return e.SafeSearchEnforcement, nil
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
	if !util.OrderedListsMatch[string](o.Alert, other.Alert) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Allow, other.Allow) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Block, other.Block) {
		return false
	}
	if !util.BoolsMatch(o.CloudInlineCat, other.CloudInlineCat) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Continue, other.Continue) {
		return false
	}
	if !o.CredentialEnforcement.matches(other.CredentialEnforcement) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(o.EnableContainerPage, other.EnableContainerPage) {
		return false
	}
	if len(o.HttpHeaderInsertion) != len(other.HttpHeaderInsertion) {
		return false
	}
	for idx := range o.HttpHeaderInsertion {
		if !o.HttpHeaderInsertion[idx].matches(&other.HttpHeaderInsertion[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.LocalInlineCat, other.LocalInlineCat) {
		return false
	}
	if !util.BoolsMatch(o.LogContainerPageOnly, other.LogContainerPageOnly) {
		return false
	}
	if !util.BoolsMatch(o.LogHttpHdrReferer, other.LogHttpHdrReferer) {
		return false
	}
	if !util.BoolsMatch(o.LogHttpHdrUserAgent, other.LogHttpHdrUserAgent) {
		return false
	}
	if !util.BoolsMatch(o.LogHttpHdrXff, other.LogHttpHdrXff) {
		return false
	}
	if !util.OrderedListsMatch[string](o.MlavCategoryException, other.MlavCategoryException) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Override, other.Override) {
		return false
	}
	if !util.BoolsMatch(o.SafeSearchEnforcement, other.SafeSearchEnforcement) {
		return false
	}

	return true
}

func (o *CredentialEnforcement) matches(other *CredentialEnforcement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Alert, other.Alert) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Allow, other.Allow) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Block, other.Block) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Continue, other.Continue) {
		return false
	}
	if !util.StringsMatch(o.LogSeverity, other.LogSeverity) {
		return false
	}
	if !o.Mode.matches(other.Mode) {
		return false
	}

	return true
}

func (o *CredentialEnforcementMode) matches(other *CredentialEnforcementMode) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Disabled.matches(other.Disabled) {
		return false
	}
	if !o.DomainCredentials.matches(other.DomainCredentials) {
		return false
	}
	if !util.StringsMatch(o.GroupMapping, other.GroupMapping) {
		return false
	}
	if !o.IpUser.matches(other.IpUser) {
		return false
	}

	return true
}

func (o *CredentialEnforcementModeDisabled) matches(other *CredentialEnforcementModeDisabled) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *CredentialEnforcementModeDomainCredentials) matches(other *CredentialEnforcementModeDomainCredentials) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *CredentialEnforcementModeIpUser) matches(other *CredentialEnforcementModeIpUser) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *HttpHeaderInsertion) matches(other *HttpHeaderInsertion) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if len(o.Type) != len(other.Type) {
		return false
	}
	for idx := range o.Type {
		if !o.Type[idx].matches(&other.Type[idx]) {
			return false
		}
	}

	return true
}

func (o *HttpHeaderInsertionType) matches(other *HttpHeaderInsertionType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.Headers) != len(other.Headers) {
		return false
	}
	for idx := range o.Headers {
		if !o.Headers[idx].matches(&other.Headers[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.Domains, other.Domains) {
		return false
	}

	return true
}

func (o *HttpHeaderInsertionTypeHeaders) matches(other *HttpHeaderInsertionTypeHeaders) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Header, other.Header) {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}
	if !util.BoolsMatch(o.Log, other.Log) {
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
