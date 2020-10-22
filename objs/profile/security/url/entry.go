package url

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// URL filtering security profile.
type Entry struct {
	Name                      string
	Description               string
	DynamicUrl                bool     // Removed in 9.0
	ExpiredLicenseAction      bool     // Removed in 9.0
	BlockListAction           string   // Removed in 9.0
	BlockList                 []string // Removed in 9.0
	AllowList                 []string // Removed in 9.0
	AllowCategories           []string
	AlertCategories           []string
	BlockCategories           []string
	ContinueCategories        []string
	OverrideCategories        []string
	TrackContainerPage        bool
	LogContainerPageOnly      bool
	SafeSearchEnforcement     bool
	LogHttpHeaderXff          bool
	LogHttpHeaderUserAgent    bool
	LogHttpHeaderReferer      bool
	CesMode                   string                 // 8.0
	CesModeGroupMapping       string                 // 8.0
	CesLogSeverity            string                 // 8.0
	CesAllowCategories        []string               // 8.0
	CesAlertCategories        []string               // 8.0
	CesBlockCategories        []string               // 8.0
	CesContinueCategories     []string               // 8.0
	HttpHeaderInsertions      []HttpHeaderInsertion  // 8.1
	MachineLearningModels     []MachineLearningModel // 10.0
	MachineLearningExceptions []string               // 10.0
}

type HttpHeaderInsertion struct {
	Name        string
	Type        string
	Domains     []string
	HttpHeaders []HttpHeader
}

// HttpHeader is an individual HTTP header.  In PAN-OS, the Name param is
// auto generated and look like "header-0", "header-1"..  If the Name param
// is an empty string, the name will be auto populated as appropriate.
type HttpHeader struct {
	Name   string
	Header string
	Value  string
	Log    bool
}

type MachineLearningModel struct {
	Model  string
	Action string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.DynamicUrl = s.DynamicUrl
	o.ExpiredLicenseAction = s.ExpiredLicenseAction
	o.BlockListAction = s.BlockListAction
	o.BlockList = s.BlockList
	o.AllowList = s.AllowList
	o.AllowCategories = s.AllowCategories
	o.AlertCategories = s.AlertCategories
	o.BlockCategories = s.BlockCategories
	o.ContinueCategories = s.ContinueCategories
	o.OverrideCategories = s.OverrideCategories
	o.TrackContainerPage = s.TrackContainerPage
	o.LogContainerPageOnly = s.LogContainerPageOnly
	o.SafeSearchEnforcement = s.SafeSearchEnforcement
	o.LogHttpHeaderXff = s.LogHttpHeaderXff
	o.LogHttpHeaderUserAgent = s.LogHttpHeaderUserAgent
	o.LogHttpHeaderReferer = s.LogHttpHeaderReferer
	o.CesMode = s.CesMode
	o.CesModeGroupMapping = s.CesModeGroupMapping
	o.CesLogSeverity = s.CesLogSeverity
	o.CesAllowCategories = s.CesAllowCategories
	o.CesAlertCategories = s.CesAlertCategories
	o.CesBlockCategories = s.CesBlockCategories
	o.CesContinueCategories = s.CesContinueCategories
	o.HttpHeaderInsertions = s.HttpHeaderInsertions
	o.MachineLearningModels = s.MachineLearningModels
	o.MachineLearningExceptions = s.MachineLearningExceptions
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

// 7.1 and lower.
type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:                   o.Name,
		Description:            o.Description,
		DynamicUrl:             util.AsBool(o.DynamicUrl),
		ExpiredLicenseAction:   util.AsBool(o.ExpiredLicenseAction),
		BlockListAction:        o.BlockListAction,
		BlockList:              util.MemToStr(o.BlockList),
		AllowList:              util.MemToStr(o.AllowList),
		AllowCategories:        util.MemToStr(o.AllowCategories),
		AlertCategories:        util.MemToStr(o.AlertCategories),
		BlockCategories:        util.MemToStr(o.BlockCategories),
		ContinueCategories:     util.MemToStr(o.ContinueCategories),
		OverrideCategories:     util.MemToStr(o.OverrideCategories),
		TrackContainerPage:     util.AsBool(o.TrackContainerPage),
		LogContainerPageOnly:   util.AsBool(o.LogContainerPageOnly),
		SafeSearchEnforcement:  util.AsBool(o.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.AsBool(o.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.AsBool(o.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.AsBool(o.LogHttpHeaderReferer),
	}

	return ans
}

type entry_v1 struct {
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	Description            string           `xml:"description,omitempty"`
	DynamicUrl             string           `xml:"dynamic-url"`
	ExpiredLicenseAction   string           `xml:"license-expired"`
	BlockListAction        string           `xml:"action"`
	BlockList              *util.MemberType `xml:"block-list"`
	AllowList              *util.MemberType `xml:"allow-list"`
	AllowCategories        *util.MemberType `xml:"allow"`
	AlertCategories        *util.MemberType `xml:"alert"`
	BlockCategories        *util.MemberType `xml:"block"`
	ContinueCategories     *util.MemberType `xml:"continue"`
	OverrideCategories     *util.MemberType `xml:"override"`
	TrackContainerPage     string           `xml:"enable-container-page"`
	LogContainerPageOnly   string           `xml:"log-container-page-only"`
	SafeSearchEnforcement  string           `xml:"safe-search-enforcement"`
	LogHttpHeaderXff       string           `xml:"log-http-hdr-xff"`
	LogHttpHeaderUserAgent string           `xml:"log-http-hdr-user-agent"`
	LogHttpHeaderReferer   string           `xml:"log-http-hdr-referer"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:                   e.Name,
		Description:            e.Description,
		DynamicUrl:             util.YesNo(e.DynamicUrl),
		ExpiredLicenseAction:   util.YesNo(e.ExpiredLicenseAction),
		BlockListAction:        e.BlockListAction,
		BlockList:              util.StrToMem(e.BlockList),
		AllowList:              util.StrToMem(e.AllowList),
		AllowCategories:        util.StrToMem(e.AllowCategories),
		AlertCategories:        util.StrToMem(e.AlertCategories),
		BlockCategories:        util.StrToMem(e.BlockCategories),
		ContinueCategories:     util.StrToMem(e.ContinueCategories),
		OverrideCategories:     util.StrToMem(e.OverrideCategories),
		TrackContainerPage:     util.YesNo(e.TrackContainerPage),
		LogContainerPageOnly:   util.YesNo(e.LogContainerPageOnly),
		SafeSearchEnforcement:  util.YesNo(e.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.YesNo(e.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.YesNo(e.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.YesNo(e.LogHttpHeaderReferer),
	}

	return ans
}

type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o *container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:                   o.Name,
		Description:            o.Description,
		DynamicUrl:             util.AsBool(o.DynamicUrl),
		ExpiredLicenseAction:   util.AsBool(o.ExpiredLicenseAction),
		BlockListAction:        o.BlockListAction,
		BlockList:              util.MemToStr(o.BlockList),
		AllowList:              util.MemToStr(o.AllowList),
		AllowCategories:        util.MemToStr(o.AllowCategories),
		AlertCategories:        util.MemToStr(o.AlertCategories),
		BlockCategories:        util.MemToStr(o.BlockCategories),
		ContinueCategories:     util.MemToStr(o.ContinueCategories),
		OverrideCategories:     util.MemToStr(o.OverrideCategories),
		TrackContainerPage:     util.AsBool(o.TrackContainerPage),
		LogContainerPageOnly:   util.AsBool(o.LogContainerPageOnly),
		SafeSearchEnforcement:  util.AsBool(o.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.AsBool(o.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.AsBool(o.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.AsBool(o.LogHttpHeaderReferer),
	}

	if o.Ces != nil {
		switch {
		case o.Ces.Mode.CesModeDisabled != nil:
			ans.CesMode = CesModeDisabled
		case o.Ces.Mode.CesModeIpUser != nil:
			ans.CesMode = CesModeIpUser
		case o.Ces.Mode.CesModeDomainCredentials != nil:
			ans.CesMode = CesModeDomainCredentials
		case o.Ces.Mode.CesModeGroupMapping != "":
			ans.CesMode = CesModeGroupMapping
			ans.CesModeGroupMapping = o.Ces.Mode.CesModeGroupMapping
		}

		ans.CesLogSeverity = o.Ces.CesLogSeverity
		ans.CesAllowCategories = util.MemToStr(o.Ces.CesAllowCategories)
		ans.CesAlertCategories = util.MemToStr(o.Ces.CesAlertCategories)
		ans.CesBlockCategories = util.MemToStr(o.Ces.CesBlockCategories)
		ans.CesContinueCategories = util.MemToStr(o.Ces.CesContinueCategories)
	}

	return ans
}

// 8.0
type entry_v2 struct {
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	Description            string           `xml:"description,omitempty"`
	DynamicUrl             string           `xml:"dynamic-url"`
	ExpiredLicenseAction   string           `xml:"license-expired"`
	BlockListAction        string           `xml:"action"`
	BlockList              *util.MemberType `xml:"block-list"`
	AllowList              *util.MemberType `xml:"allow-list"`
	AllowCategories        *util.MemberType `xml:"allow"`
	AlertCategories        *util.MemberType `xml:"alert"`
	BlockCategories        *util.MemberType `xml:"block"`
	ContinueCategories     *util.MemberType `xml:"continue"`
	OverrideCategories     *util.MemberType `xml:"override"`
	Ces                    *creds           `xml:"credential-enforcement"`
	TrackContainerPage     string           `xml:"enable-container-page"`
	LogContainerPageOnly   string           `xml:"log-container-page-only"`
	SafeSearchEnforcement  string           `xml:"safe-search-enforcement"`
	LogHttpHeaderXff       string           `xml:"log-http-hdr-xff"`
	LogHttpHeaderUserAgent string           `xml:"log-http-hdr-user-agent"`
	LogHttpHeaderReferer   string           `xml:"log-http-hdr-referer"`
}

type creds struct {
	Mode                  credMode         `xml:"mode"`
	CesLogSeverity        string           `xml:"log-severity"`
	CesAllowCategories    *util.MemberType `xml:"allow"`
	CesAlertCategories    *util.MemberType `xml:"alert"`
	CesBlockCategories    *util.MemberType `xml:"block"`
	CesContinueCategories *util.MemberType `xml:"continue"`
}

type credMode struct {
	CesModeDisabled          *string `xml:"disabled"`
	CesModeIpUser            *string `xml:"ip-user"`
	CesModeDomainCredentials *string `xml:"domain-credentials"`
	CesModeGroupMapping      string  `xml:"group-mapping,omitempty"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:                   e.Name,
		Description:            e.Description,
		DynamicUrl:             util.YesNo(e.DynamicUrl),
		ExpiredLicenseAction:   util.YesNo(e.ExpiredLicenseAction),
		BlockListAction:        e.BlockListAction,
		BlockList:              util.StrToMem(e.BlockList),
		AllowList:              util.StrToMem(e.AllowList),
		AllowCategories:        util.StrToMem(e.AllowCategories),
		AlertCategories:        util.StrToMem(e.AlertCategories),
		BlockCategories:        util.StrToMem(e.BlockCategories),
		ContinueCategories:     util.StrToMem(e.ContinueCategories),
		OverrideCategories:     util.StrToMem(e.OverrideCategories),
		TrackContainerPage:     util.YesNo(e.TrackContainerPage),
		LogContainerPageOnly:   util.YesNo(e.LogContainerPageOnly),
		SafeSearchEnforcement:  util.YesNo(e.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.YesNo(e.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.YesNo(e.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.YesNo(e.LogHttpHeaderReferer),
	}

	if e.CesMode != "" || e.CesLogSeverity != "" || len(e.CesAllowCategories) != 0 || len(e.CesAlertCategories) != 0 || len(e.CesBlockCategories) != 0 || len(e.CesContinueCategories) != 0 {
		s := ""
		var m credMode
		switch e.CesMode {
		case CesModeDisabled:
			m = credMode{
				CesModeDisabled: &s,
			}
		case CesModeIpUser:
			m = credMode{
				CesModeIpUser: &s,
			}
		case CesModeDomainCredentials:
			m = credMode{
				CesModeDomainCredentials: &s,
			}
		case CesModeGroupMapping:
			m = credMode{
				CesModeGroupMapping: e.CesModeGroupMapping,
			}
		}

		ans.Ces = &creds{
			Mode:                  m,
			CesLogSeverity:        e.CesLogSeverity,
			CesAllowCategories:    util.StrToMem(e.CesAllowCategories),
			CesAlertCategories:    util.StrToMem(e.CesAlertCategories),
			CesBlockCategories:    util.StrToMem(e.CesBlockCategories),
			CesContinueCategories: util.StrToMem(e.CesContinueCategories),
		}
	}

	return ans
}

// 8.1
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o *container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v3) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:                   o.Name,
		Description:            o.Description,
		DynamicUrl:             util.AsBool(o.DynamicUrl),
		ExpiredLicenseAction:   util.AsBool(o.ExpiredLicenseAction),
		BlockListAction:        o.BlockListAction,
		BlockList:              util.MemToStr(o.BlockList),
		AllowList:              util.MemToStr(o.AllowList),
		AllowCategories:        util.MemToStr(o.AllowCategories),
		AlertCategories:        util.MemToStr(o.AlertCategories),
		BlockCategories:        util.MemToStr(o.BlockCategories),
		ContinueCategories:     util.MemToStr(o.ContinueCategories),
		OverrideCategories:     util.MemToStr(o.OverrideCategories),
		TrackContainerPage:     util.AsBool(o.TrackContainerPage),
		LogContainerPageOnly:   util.AsBool(o.LogContainerPageOnly),
		SafeSearchEnforcement:  util.AsBool(o.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.AsBool(o.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.AsBool(o.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.AsBool(o.LogHttpHeaderReferer),
	}

	if o.Ces != nil {
		switch {
		case o.Ces.Mode.CesModeDisabled != nil:
			ans.CesMode = CesModeDisabled
		case o.Ces.Mode.CesModeIpUser != nil:
			ans.CesMode = CesModeIpUser
		case o.Ces.Mode.CesModeDomainCredentials != nil:
			ans.CesMode = CesModeDomainCredentials
		case o.Ces.Mode.CesModeGroupMapping != "":
			ans.CesMode = CesModeGroupMapping
			ans.CesModeGroupMapping = o.Ces.Mode.CesModeGroupMapping
		}

		ans.CesLogSeverity = o.Ces.CesLogSeverity
		ans.CesAllowCategories = util.MemToStr(o.Ces.CesAllowCategories)
		ans.CesAlertCategories = util.MemToStr(o.Ces.CesAlertCategories)
		ans.CesBlockCategories = util.MemToStr(o.Ces.CesBlockCategories)
		ans.CesContinueCategories = util.MemToStr(o.Ces.CesContinueCategories)
	}

	if o.Hhi != nil {
		ins := make([]HttpHeaderInsertion, 0, len(o.Hhi.Entries))
		for _, hhiObj := range o.Hhi.Entries {
			var headerList []HttpHeader
			if hhiObj.Types.Entry.Headers != nil {
				headerList = make([]HttpHeader, 0, len(hhiObj.Types.Entry.Headers.Entries))
				for _, hle := range hhiObj.Types.Entry.Headers.Entries {
					headerList = append(headerList, HttpHeader{
						Name:   hle.Name,
						Header: hle.Header,
						Value:  hle.Value,
						Log:    util.AsBool(hle.Log),
					})
				}
			}

			ins = append(ins, HttpHeaderInsertion{
				Name:        hhiObj.Name,
				Type:        hhiObj.Types.Entry.Type,
				Domains:     util.MemToStr(hhiObj.Types.Entry.Domains),
				HttpHeaders: headerList,
			})
		}

		ans.HttpHeaderInsertions = ins
	}

	return ans
}

type entry_v3 struct {
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	Description            string           `xml:"description,omitempty"`
	DynamicUrl             string           `xml:"dynamic-url"`
	ExpiredLicenseAction   string           `xml:"license-expired"`
	BlockListAction        string           `xml:"action"`
	BlockList              *util.MemberType `xml:"block-list"`
	AllowList              *util.MemberType `xml:"allow-list"`
	AllowCategories        *util.MemberType `xml:"allow"`
	AlertCategories        *util.MemberType `xml:"alert"`
	BlockCategories        *util.MemberType `xml:"block"`
	ContinueCategories     *util.MemberType `xml:"continue"`
	OverrideCategories     *util.MemberType `xml:"override"`
	Ces                    *creds           `xml:"credential-enforcement"`
	TrackContainerPage     string           `xml:"enable-container-page"`
	LogContainerPageOnly   string           `xml:"log-container-page-only"`
	SafeSearchEnforcement  string           `xml:"safe-search-enforcement"`
	LogHttpHeaderXff       string           `xml:"log-http-hdr-xff"`
	LogHttpHeaderUserAgent string           `xml:"log-http-hdr-user-agent"`
	LogHttpHeaderReferer   string           `xml:"log-http-hdr-referer"`
	Hhi                    *hhi             `xml:"http-header-insertion"`
}

type hhi struct {
	Entries []hhiEntry `xml:"entry"`
}

type hhiEntry struct {
	Name  string  `xml:"name,attr"`
	Types hhiType `xml:"type"`
}

type hhiType struct {
	Entry hhiTypeEntry `xml:"entry"`
}

type hhiTypeEntry struct {
	Type    string           `xml:"name"`
	Domains *util.MemberType `xml:"domains"`
	Headers *headers         `xml:"headers"`
}

type headers struct {
	Entries []headerEntry `xml:"entry"`
}

type headerEntry struct {
	Name   string `xml:"name,attr"`
	Header string `xml:"header"`
	Value  string `xml:"value"`
	Log    string `xml:"log"`
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:                   e.Name,
		Description:            e.Description,
		DynamicUrl:             util.YesNo(e.DynamicUrl),
		ExpiredLicenseAction:   util.YesNo(e.ExpiredLicenseAction),
		BlockListAction:        e.BlockListAction,
		BlockList:              util.StrToMem(e.BlockList),
		AllowList:              util.StrToMem(e.AllowList),
		AllowCategories:        util.StrToMem(e.AllowCategories),
		AlertCategories:        util.StrToMem(e.AlertCategories),
		BlockCategories:        util.StrToMem(e.BlockCategories),
		ContinueCategories:     util.StrToMem(e.ContinueCategories),
		OverrideCategories:     util.StrToMem(e.OverrideCategories),
		TrackContainerPage:     util.YesNo(e.TrackContainerPage),
		LogContainerPageOnly:   util.YesNo(e.LogContainerPageOnly),
		SafeSearchEnforcement:  util.YesNo(e.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.YesNo(e.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.YesNo(e.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.YesNo(e.LogHttpHeaderReferer),
	}

	if e.CesMode != "" || e.CesLogSeverity != "" || len(e.CesAllowCategories) != 0 || len(e.CesAlertCategories) != 0 || len(e.CesBlockCategories) != 0 || len(e.CesContinueCategories) != 0 {
		s := ""
		var m credMode
		switch e.CesMode {
		case CesModeDisabled:
			m = credMode{
				CesModeDisabled: &s,
			}
		case CesModeIpUser:
			m = credMode{
				CesModeIpUser: &s,
			}
		case CesModeDomainCredentials:
			m = credMode{
				CesModeDomainCredentials: &s,
			}
		case CesModeGroupMapping:
			m = credMode{
				CesModeGroupMapping: e.CesModeGroupMapping,
			}
		}

		ans.Ces = &creds{
			Mode:                  m,
			CesLogSeverity:        e.CesLogSeverity,
			CesAllowCategories:    util.StrToMem(e.CesAllowCategories),
			CesAlertCategories:    util.StrToMem(e.CesAlertCategories),
			CesBlockCategories:    util.StrToMem(e.CesBlockCategories),
			CesContinueCategories: util.StrToMem(e.CesContinueCategories),
		}
	}

	if len(e.HttpHeaderInsertions) > 0 {
		hhiEntries := make([]hhiEntry, 0, len(e.HttpHeaderInsertions))

		for _, hhiObject := range e.HttpHeaderInsertions {
			var headersInst *headers

			if len(hhiObject.HttpHeaders) > 0 {
				list := make([]headerEntry, 0, len(hhiObject.HttpHeaders))
				for i := range hhiObject.HttpHeaders {
					var name string
					if hhiObject.HttpHeaders[i].Name == "" {
						name = fmt.Sprintf("header-%d", i)
					} else {
						name = hhiObject.HttpHeaders[i].Name
					}
					list = append(list, headerEntry{
						Name:   name,
						Header: hhiObject.HttpHeaders[i].Header,
						Value:  hhiObject.HttpHeaders[i].Value,
						Log:    util.YesNo(hhiObject.HttpHeaders[i].Log),
					})
				}

				headersInst = &headers{
					Entries: list,
				}
			}

			hhiEntries = append(hhiEntries, hhiEntry{
				Name: hhiObject.Name,
				Types: hhiType{
					Entry: hhiTypeEntry{
						Type:    hhiObject.Type,
						Domains: util.StrToMem(hhiObject.Domains),
						Headers: headersInst,
					},
				},
			})
		}

		ans.Hhi = &hhi{
			Entries: hhiEntries,
		}
	}

	return ans
}

// 9.0
type container_v4 struct {
	Answer []entry_v4 `xml:"entry"`
}

func (o *container_v4) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v4) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v4) normalize() Entry {
	ans := Entry{
		Name:                   o.Name,
		Description:            o.Description,
		AllowCategories:        util.MemToStr(o.AllowCategories),
		AlertCategories:        util.MemToStr(o.AlertCategories),
		BlockCategories:        util.MemToStr(o.BlockCategories),
		ContinueCategories:     util.MemToStr(o.ContinueCategories),
		OverrideCategories:     util.MemToStr(o.OverrideCategories),
		TrackContainerPage:     util.AsBool(o.TrackContainerPage),
		LogContainerPageOnly:   util.AsBool(o.LogContainerPageOnly),
		SafeSearchEnforcement:  util.AsBool(o.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.AsBool(o.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.AsBool(o.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.AsBool(o.LogHttpHeaderReferer),
	}

	if o.Ces != nil {
		switch {
		case o.Ces.Mode.CesModeDisabled != nil:
			ans.CesMode = CesModeDisabled
		case o.Ces.Mode.CesModeIpUser != nil:
			ans.CesMode = CesModeIpUser
		case o.Ces.Mode.CesModeDomainCredentials != nil:
			ans.CesMode = CesModeDomainCredentials
		case o.Ces.Mode.CesModeGroupMapping != "":
			ans.CesMode = CesModeGroupMapping
			ans.CesModeGroupMapping = o.Ces.Mode.CesModeGroupMapping
		}

		ans.CesLogSeverity = o.Ces.CesLogSeverity
		ans.CesAllowCategories = util.MemToStr(o.Ces.CesAllowCategories)
		ans.CesAlertCategories = util.MemToStr(o.Ces.CesAlertCategories)
		ans.CesBlockCategories = util.MemToStr(o.Ces.CesBlockCategories)
		ans.CesContinueCategories = util.MemToStr(o.Ces.CesContinueCategories)
	}

	if o.Hhi != nil {
		ins := make([]HttpHeaderInsertion, 0, len(o.Hhi.Entries))
		for _, hhiObj := range o.Hhi.Entries {
			var headerList []HttpHeader
			if hhiObj.Types.Entry.Headers != nil {
				headerList = make([]HttpHeader, 0, len(hhiObj.Types.Entry.Headers.Entries))
				for _, hle := range hhiObj.Types.Entry.Headers.Entries {
					headerList = append(headerList, HttpHeader{
						Name:   hle.Name,
						Header: hle.Header,
						Value:  hle.Value,
						Log:    util.AsBool(hle.Log),
					})
				}
			}

			ins = append(ins, HttpHeaderInsertion{
				Name:        hhiObj.Name,
				Type:        hhiObj.Types.Entry.Type,
				Domains:     util.MemToStr(hhiObj.Types.Entry.Domains),
				HttpHeaders: headerList,
			})
		}

		ans.HttpHeaderInsertions = ins
	}

	return ans
}

type entry_v4 struct {
	XMLName                xml.Name         `xml:"entry"`
	Name                   string           `xml:"name,attr"`
	Description            string           `xml:"description,omitempty"`
	AllowCategories        *util.MemberType `xml:"allow"`
	AlertCategories        *util.MemberType `xml:"alert"`
	BlockCategories        *util.MemberType `xml:"block"`
	ContinueCategories     *util.MemberType `xml:"continue"`
	OverrideCategories     *util.MemberType `xml:"override"`
	Ces                    *creds           `xml:"credential-enforcement"`
	TrackContainerPage     string           `xml:"enable-container-page"`
	LogContainerPageOnly   string           `xml:"log-container-page-only"`
	SafeSearchEnforcement  string           `xml:"safe-search-enforcement"`
	LogHttpHeaderXff       string           `xml:"log-http-hdr-xff"`
	LogHttpHeaderUserAgent string           `xml:"log-http-hdr-user-agent"`
	LogHttpHeaderReferer   string           `xml:"log-http-hdr-referer"`
	Hhi                    *hhi             `xml:"http-header-insertion"`
}

func specify_v4(e Entry) interface{} {
	ans := entry_v4{
		Name:                   e.Name,
		Description:            e.Description,
		AllowCategories:        util.StrToMem(e.AllowCategories),
		AlertCategories:        util.StrToMem(e.AlertCategories),
		BlockCategories:        util.StrToMem(e.BlockCategories),
		ContinueCategories:     util.StrToMem(e.ContinueCategories),
		OverrideCategories:     util.StrToMem(e.OverrideCategories),
		TrackContainerPage:     util.YesNo(e.TrackContainerPage),
		LogContainerPageOnly:   util.YesNo(e.LogContainerPageOnly),
		SafeSearchEnforcement:  util.YesNo(e.SafeSearchEnforcement),
		LogHttpHeaderXff:       util.YesNo(e.LogHttpHeaderXff),
		LogHttpHeaderUserAgent: util.YesNo(e.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:   util.YesNo(e.LogHttpHeaderReferer),
	}

	if e.CesMode != "" || e.CesLogSeverity != "" || len(e.CesAllowCategories) != 0 || len(e.CesAlertCategories) != 0 || len(e.CesBlockCategories) != 0 || len(e.CesContinueCategories) != 0 {
		s := ""
		var m credMode
		switch e.CesMode {
		case CesModeDisabled:
			m = credMode{
				CesModeDisabled: &s,
			}
		case CesModeIpUser:
			m = credMode{
				CesModeIpUser: &s,
			}
		case CesModeDomainCredentials:
			m = credMode{
				CesModeDomainCredentials: &s,
			}
		case CesModeGroupMapping:
			m = credMode{
				CesModeGroupMapping: e.CesModeGroupMapping,
			}
		}

		ans.Ces = &creds{
			Mode:                  m,
			CesLogSeverity:        e.CesLogSeverity,
			CesAllowCategories:    util.StrToMem(e.CesAllowCategories),
			CesAlertCategories:    util.StrToMem(e.CesAlertCategories),
			CesBlockCategories:    util.StrToMem(e.CesBlockCategories),
			CesContinueCategories: util.StrToMem(e.CesContinueCategories),
		}
	}

	if len(e.HttpHeaderInsertions) > 0 {
		hhiEntries := make([]hhiEntry, 0, len(e.HttpHeaderInsertions))

		for _, hhiObject := range e.HttpHeaderInsertions {
			var headersInst *headers

			if len(hhiObject.HttpHeaders) > 0 {
				list := make([]headerEntry, 0, len(hhiObject.HttpHeaders))
				for i := range hhiObject.HttpHeaders {
					var name string
					if hhiObject.HttpHeaders[i].Name == "" {
						name = fmt.Sprintf("header-%d", i)
					} else {
						name = hhiObject.HttpHeaders[i].Name
					}
					list = append(list, headerEntry{
						Name:   name,
						Header: hhiObject.HttpHeaders[i].Header,
						Value:  hhiObject.HttpHeaders[i].Value,
						Log:    util.YesNo(hhiObject.HttpHeaders[i].Log),
					})
				}

				headersInst = &headers{
					Entries: list,
				}
			}

			hhiEntries = append(hhiEntries, hhiEntry{
				Name: hhiObject.Name,
				Types: hhiType{
					Entry: hhiTypeEntry{
						Type:    hhiObject.Type,
						Domains: util.StrToMem(hhiObject.Domains),
						Headers: headersInst,
					},
				},
			})
		}

		ans.Hhi = &hhi{
			Entries: hhiEntries,
		}
	}

	return ans
}

// 10.0
type container_v5 struct {
	Answer []entry_v5 `xml:"entry"`
}

func (o *container_v5) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v5) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v5) normalize() Entry {
	ans := Entry{
		Name:                      o.Name,
		Description:               o.Description,
		AllowCategories:           util.MemToStr(o.AllowCategories),
		AlertCategories:           util.MemToStr(o.AlertCategories),
		BlockCategories:           util.MemToStr(o.BlockCategories),
		ContinueCategories:        util.MemToStr(o.ContinueCategories),
		OverrideCategories:        util.MemToStr(o.OverrideCategories),
		TrackContainerPage:        util.AsBool(o.TrackContainerPage),
		LogContainerPageOnly:      util.AsBool(o.LogContainerPageOnly),
		SafeSearchEnforcement:     util.AsBool(o.SafeSearchEnforcement),
		LogHttpHeaderXff:          util.AsBool(o.LogHttpHeaderXff),
		LogHttpHeaderUserAgent:    util.AsBool(o.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:      util.AsBool(o.LogHttpHeaderReferer),
		MachineLearningExceptions: util.MemToStr(o.MachineLearningExceptions),
	}

	if o.Ces != nil {
		switch {
		case o.Ces.Mode.CesModeDisabled != nil:
			ans.CesMode = CesModeDisabled
		case o.Ces.Mode.CesModeIpUser != nil:
			ans.CesMode = CesModeIpUser
		case o.Ces.Mode.CesModeDomainCredentials != nil:
			ans.CesMode = CesModeDomainCredentials
		case o.Ces.Mode.CesModeGroupMapping != "":
			ans.CesMode = CesModeGroupMapping
			ans.CesModeGroupMapping = o.Ces.Mode.CesModeGroupMapping
		}

		ans.CesLogSeverity = o.Ces.CesLogSeverity
		ans.CesAllowCategories = util.MemToStr(o.Ces.CesAllowCategories)
		ans.CesAlertCategories = util.MemToStr(o.Ces.CesAlertCategories)
		ans.CesBlockCategories = util.MemToStr(o.Ces.CesBlockCategories)
		ans.CesContinueCategories = util.MemToStr(o.Ces.CesContinueCategories)
	}

	if o.Hhi != nil {
		ins := make([]HttpHeaderInsertion, 0, len(o.Hhi.Entries))
		for _, hhiObj := range o.Hhi.Entries {
			var headerList []HttpHeader
			if hhiObj.Types.Entry.Headers != nil {
				headerList = make([]HttpHeader, 0, len(hhiObj.Types.Entry.Headers.Entries))
				for _, hle := range hhiObj.Types.Entry.Headers.Entries {
					headerList = append(headerList, HttpHeader{
						Name:   hle.Name,
						Header: hle.Header,
						Value:  hle.Value,
						Log:    util.AsBool(hle.Log),
					})
				}
			}

			ins = append(ins, HttpHeaderInsertion{
				Name:        hhiObj.Name,
				Type:        hhiObj.Types.Entry.Type,
				Domains:     util.MemToStr(hhiObj.Types.Entry.Domains),
				HttpHeaders: headerList,
			})
		}

		ans.HttpHeaderInsertions = ins
	}

	if o.MlModels != nil {
		listing := make([]MachineLearningModel, 0, len(o.MlModels.Entries))
		for _, model := range o.MlModels.Entries {
			listing = append(listing, MachineLearningModel{
				Model:  model.Model,
				Action: model.Action,
			})
		}

		ans.MachineLearningModels = listing
	}

	return ans
}

type entry_v5 struct {
	XMLName                   xml.Name         `xml:"entry"`
	Name                      string           `xml:"name,attr"`
	Description               string           `xml:"description,omitempty"`
	AllowCategories           *util.MemberType `xml:"allow"`
	AlertCategories           *util.MemberType `xml:"alert"`
	BlockCategories           *util.MemberType `xml:"block"`
	ContinueCategories        *util.MemberType `xml:"continue"`
	OverrideCategories        *util.MemberType `xml:"override"`
	Ces                       *creds           `xml:"credential-enforcement"`
	TrackContainerPage        string           `xml:"enable-container-page"`
	LogContainerPageOnly      string           `xml:"log-container-page-only"`
	SafeSearchEnforcement     string           `xml:"safe-search-enforcement"`
	LogHttpHeaderXff          string           `xml:"log-http-hdr-xff"`
	LogHttpHeaderUserAgent    string           `xml:"log-http-hdr-user-agent"`
	LogHttpHeaderReferer      string           `xml:"log-http-hdr-referer"`
	Hhi                       *hhi             `xml:"http-header-insertion"`
	MachineLearningExceptions *util.MemberType `xml:"mlav-category-exception"`
	MlModels                  *mlmodels        `xml:"mlav-engine-urlbased-enabled"`
}

type mlmodels struct {
	Entries []mlmodel `xml:"entry"`
}

type mlmodel struct {
	Model  string `xml:"name,attr"`
	Action string `xml:"mlav-policy-action"`
}

func specify_v5(e Entry) interface{} {
	ans := entry_v5{
		Name:                      e.Name,
		Description:               e.Description,
		AllowCategories:           util.StrToMem(e.AllowCategories),
		AlertCategories:           util.StrToMem(e.AlertCategories),
		BlockCategories:           util.StrToMem(e.BlockCategories),
		ContinueCategories:        util.StrToMem(e.ContinueCategories),
		OverrideCategories:        util.StrToMem(e.OverrideCategories),
		TrackContainerPage:        util.YesNo(e.TrackContainerPage),
		LogContainerPageOnly:      util.YesNo(e.LogContainerPageOnly),
		SafeSearchEnforcement:     util.YesNo(e.SafeSearchEnforcement),
		LogHttpHeaderXff:          util.YesNo(e.LogHttpHeaderXff),
		LogHttpHeaderUserAgent:    util.YesNo(e.LogHttpHeaderUserAgent),
		LogHttpHeaderReferer:      util.YesNo(e.LogHttpHeaderReferer),
		MachineLearningExceptions: util.StrToMem(e.MachineLearningExceptions),
	}

	if e.CesMode != "" || e.CesLogSeverity != "" || len(e.CesAllowCategories) != 0 || len(e.CesAlertCategories) != 0 || len(e.CesBlockCategories) != 0 || len(e.CesContinueCategories) != 0 {
		s := ""
		var m credMode
		switch e.CesMode {
		case CesModeDisabled:
			m = credMode{
				CesModeDisabled: &s,
			}
		case CesModeIpUser:
			m = credMode{
				CesModeIpUser: &s,
			}
		case CesModeDomainCredentials:
			m = credMode{
				CesModeDomainCredentials: &s,
			}
		case CesModeGroupMapping:
			m = credMode{
				CesModeGroupMapping: e.CesModeGroupMapping,
			}
		}

		ans.Ces = &creds{
			Mode:                  m,
			CesLogSeverity:        e.CesLogSeverity,
			CesAllowCategories:    util.StrToMem(e.CesAllowCategories),
			CesAlertCategories:    util.StrToMem(e.CesAlertCategories),
			CesBlockCategories:    util.StrToMem(e.CesBlockCategories),
			CesContinueCategories: util.StrToMem(e.CesContinueCategories),
		}
	}

	if len(e.HttpHeaderInsertions) > 0 {
		hhiEntries := make([]hhiEntry, 0, len(e.HttpHeaderInsertions))

		for _, hhiObject := range e.HttpHeaderInsertions {
			var headersInst *headers

			if len(hhiObject.HttpHeaders) > 0 {
				list := make([]headerEntry, 0, len(hhiObject.HttpHeaders))
				for i := range hhiObject.HttpHeaders {
					var name string
					if hhiObject.HttpHeaders[i].Name == "" {
						name = fmt.Sprintf("header-%d", i)
					} else {
						name = hhiObject.HttpHeaders[i].Name
					}
					list = append(list, headerEntry{
						Name:   name,
						Header: hhiObject.HttpHeaders[i].Header,
						Value:  hhiObject.HttpHeaders[i].Value,
						Log:    util.YesNo(hhiObject.HttpHeaders[i].Log),
					})
				}

				headersInst = &headers{
					Entries: list,
				}
			}

			hhiEntries = append(hhiEntries, hhiEntry{
				Name: hhiObject.Name,
				Types: hhiType{
					Entry: hhiTypeEntry{
						Type:    hhiObject.Type,
						Domains: util.StrToMem(hhiObject.Domains),
						Headers: headersInst,
					},
				},
			})
		}

		ans.Hhi = &hhi{
			Entries: hhiEntries,
		}
	}

	if len(e.MachineLearningModels) > 0 {
		listing := make([]mlmodel, 0, len(e.MachineLearningModels))
		for _, model := range e.MachineLearningModels {
			listing = append(listing, mlmodel{
				Model:  model.Model,
				Action: model.Action,
			})
		}

		ans.MlModels = &mlmodels{
			Entries: listing,
		}
	}

	return ans
}
