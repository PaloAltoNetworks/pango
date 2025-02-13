package urlfiltering

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
	Suffix = []string{"profiles", "url-filtering"}
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

	Misc map[string][]generic.Xml
}

type CredentialEnforcement struct {
	Alert       []string
	Allow       []string
	Block       []string
	Continue    []string
	LogSeverity *string
	Mode        *CredentialEnforcementMode
}
type CredentialEnforcementMode struct {
	Disabled          *CredentialEnforcementModeDisabled
	DomainCredentials *CredentialEnforcementModeDomainCredentials
	GroupMapping      *string
	IpUser            *CredentialEnforcementModeIpUser
}
type CredentialEnforcementModeDisabled struct {
}
type CredentialEnforcementModeDomainCredentials struct {
}
type CredentialEnforcementModeIpUser struct {
}
type HttpHeaderInsertion struct {
	DisableOverride *string
	Name            string
	Type            []HttpHeaderInsertionType
}
type HttpHeaderInsertionType struct {
	Domains []string
	Headers []HttpHeaderInsertionTypeHeaders
	Name    string
}
type HttpHeaderInsertionTypeHeaders struct {
	Header *string
	Log    *bool
	Name   string
	Value  *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName               xml.Name                  `xml:"entry"`
	Name                  string                    `xml:"name,attr"`
	Alert                 *util.MemberType          `xml:"alert,omitempty"`
	Allow                 *util.MemberType          `xml:"allow,omitempty"`
	Block                 *util.MemberType          `xml:"block,omitempty"`
	CloudInlineCat        *string                   `xml:"cloud-inline-cat,omitempty"`
	Continue              *util.MemberType          `xml:"continue,omitempty"`
	CredentialEnforcement *CredentialEnforcementXml `xml:"credential-enforcement,omitempty"`
	Description           *string                   `xml:"description,omitempty"`
	DisableOverride       *string                   `xml:"disable-override,omitempty"`
	EnableContainerPage   *string                   `xml:"enable-container-page,omitempty"`
	HttpHeaderInsertion   []HttpHeaderInsertionXml  `xml:"http-header-insertion>entry,omitempty"`
	LocalInlineCat        *string                   `xml:"local-inline-cat,omitempty"`
	LogContainerPageOnly  *string                   `xml:"log-container-page-only,omitempty"`
	LogHttpHdrReferer     *string                   `xml:"log-http-hdr-referer,omitempty"`
	LogHttpHdrUserAgent   *string                   `xml:"log-http-hdr-user-agent,omitempty"`
	LogHttpHdrXff         *string                   `xml:"log-http-hdr-xff,omitempty"`
	MlavCategoryException *util.MemberType          `xml:"mlav-category-exception,omitempty"`
	Override              *util.MemberType          `xml:"override,omitempty"`
	SafeSearchEnforcement *string                   `xml:"safe-search-enforcement,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type CredentialEnforcementXml struct {
	Alert       *util.MemberType              `xml:"alert,omitempty"`
	Allow       *util.MemberType              `xml:"allow,omitempty"`
	Block       *util.MemberType              `xml:"block,omitempty"`
	Continue    *util.MemberType              `xml:"continue,omitempty"`
	LogSeverity *string                       `xml:"log-severity,omitempty"`
	Mode        *CredentialEnforcementModeXml `xml:"mode,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type CredentialEnforcementModeXml struct {
	Disabled          *CredentialEnforcementModeDisabledXml          `xml:"disabled,omitempty"`
	DomainCredentials *CredentialEnforcementModeDomainCredentialsXml `xml:"domain-credentials,omitempty"`
	GroupMapping      *string                                        `xml:"group-mapping,omitempty"`
	IpUser            *CredentialEnforcementModeIpUserXml            `xml:"ip-user,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type CredentialEnforcementModeDisabledXml struct {
	Misc []generic.Xml `xml:",any"`
}
type CredentialEnforcementModeDomainCredentialsXml struct {
	Misc []generic.Xml `xml:",any"`
}
type CredentialEnforcementModeIpUserXml struct {
	Misc []generic.Xml `xml:",any"`
}
type HttpHeaderInsertionXml struct {
	DisableOverride *string                      `xml:"disable-override,omitempty"`
	XMLName         xml.Name                     `xml:"entry"`
	Name            string                       `xml:"name,attr"`
	Type            []HttpHeaderInsertionTypeXml `xml:"type>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type HttpHeaderInsertionTypeXml struct {
	Domains *util.MemberType                    `xml:"domains,omitempty"`
	Headers []HttpHeaderInsertionTypeHeadersXml `xml:"headers>entry,omitempty"`
	XMLName xml.Name                            `xml:"entry"`
	Name    string                              `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type HttpHeaderInsertionTypeHeadersXml struct {
	Header  *string  `xml:"header,omitempty"`
	Log     *string  `xml:"log,omitempty"`
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   *string  `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Alert = util.StrToMem(o.Alert)
	entry.Allow = util.StrToMem(o.Allow)
	entry.Block = util.StrToMem(o.Block)
	entry.CloudInlineCat = util.YesNo(o.CloudInlineCat, nil)
	entry.Continue = util.StrToMem(o.Continue)
	var nestedCredentialEnforcement *CredentialEnforcementXml
	if o.CredentialEnforcement != nil {
		nestedCredentialEnforcement = &CredentialEnforcementXml{}
		if _, ok := o.Misc["CredentialEnforcement"]; ok {
			nestedCredentialEnforcement.Misc = o.Misc["CredentialEnforcement"]
		}
		if o.CredentialEnforcement.LogSeverity != nil {
			nestedCredentialEnforcement.LogSeverity = o.CredentialEnforcement.LogSeverity
		}
		if o.CredentialEnforcement.Mode != nil {
			nestedCredentialEnforcement.Mode = &CredentialEnforcementModeXml{}
			if _, ok := o.Misc["CredentialEnforcementMode"]; ok {
				nestedCredentialEnforcement.Mode.Misc = o.Misc["CredentialEnforcementMode"]
			}
			if o.CredentialEnforcement.Mode.Disabled != nil {
				nestedCredentialEnforcement.Mode.Disabled = &CredentialEnforcementModeDisabledXml{}
				if _, ok := o.Misc["CredentialEnforcementModeDisabled"]; ok {
					nestedCredentialEnforcement.Mode.Disabled.Misc = o.Misc["CredentialEnforcementModeDisabled"]
				}
			}
			if o.CredentialEnforcement.Mode.DomainCredentials != nil {
				nestedCredentialEnforcement.Mode.DomainCredentials = &CredentialEnforcementModeDomainCredentialsXml{}
				if _, ok := o.Misc["CredentialEnforcementModeDomainCredentials"]; ok {
					nestedCredentialEnforcement.Mode.DomainCredentials.Misc = o.Misc["CredentialEnforcementModeDomainCredentials"]
				}
			}
			if o.CredentialEnforcement.Mode.GroupMapping != nil {
				nestedCredentialEnforcement.Mode.GroupMapping = o.CredentialEnforcement.Mode.GroupMapping
			}
			if o.CredentialEnforcement.Mode.IpUser != nil {
				nestedCredentialEnforcement.Mode.IpUser = &CredentialEnforcementModeIpUserXml{}
				if _, ok := o.Misc["CredentialEnforcementModeIpUser"]; ok {
					nestedCredentialEnforcement.Mode.IpUser.Misc = o.Misc["CredentialEnforcementModeIpUser"]
				}
			}
		}
		if o.CredentialEnforcement.Alert != nil {
			nestedCredentialEnforcement.Alert = util.StrToMem(o.CredentialEnforcement.Alert)
		}
		if o.CredentialEnforcement.Allow != nil {
			nestedCredentialEnforcement.Allow = util.StrToMem(o.CredentialEnforcement.Allow)
		}
		if o.CredentialEnforcement.Block != nil {
			nestedCredentialEnforcement.Block = util.StrToMem(o.CredentialEnforcement.Block)
		}
		if o.CredentialEnforcement.Continue != nil {
			nestedCredentialEnforcement.Continue = util.StrToMem(o.CredentialEnforcement.Continue)
		}
	}
	entry.CredentialEnforcement = nestedCredentialEnforcement

	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.EnableContainerPage = util.YesNo(o.EnableContainerPage, nil)
	var nestedHttpHeaderInsertionCol []HttpHeaderInsertionXml
	if o.HttpHeaderInsertion != nil {
		nestedHttpHeaderInsertionCol = []HttpHeaderInsertionXml{}
		for _, oHttpHeaderInsertion := range o.HttpHeaderInsertion {
			nestedHttpHeaderInsertion := HttpHeaderInsertionXml{}
			if _, ok := o.Misc["HttpHeaderInsertion"]; ok {
				nestedHttpHeaderInsertion.Misc = o.Misc["HttpHeaderInsertion"]
			}
			if oHttpHeaderInsertion.DisableOverride != nil {
				nestedHttpHeaderInsertion.DisableOverride = oHttpHeaderInsertion.DisableOverride
			}
			if oHttpHeaderInsertion.Type != nil {
				nestedHttpHeaderInsertion.Type = []HttpHeaderInsertionTypeXml{}
				for _, oHttpHeaderInsertionType := range oHttpHeaderInsertion.Type {
					nestedHttpHeaderInsertionType := HttpHeaderInsertionTypeXml{}
					if _, ok := o.Misc["HttpHeaderInsertionType"]; ok {
						nestedHttpHeaderInsertionType.Misc = o.Misc["HttpHeaderInsertionType"]
					}
					if oHttpHeaderInsertionType.Headers != nil {
						nestedHttpHeaderInsertionType.Headers = []HttpHeaderInsertionTypeHeadersXml{}
						for _, oHttpHeaderInsertionTypeHeaders := range oHttpHeaderInsertionType.Headers {
							nestedHttpHeaderInsertionTypeHeaders := HttpHeaderInsertionTypeHeadersXml{}
							if _, ok := o.Misc["HttpHeaderInsertionTypeHeaders"]; ok {
								nestedHttpHeaderInsertionTypeHeaders.Misc = o.Misc["HttpHeaderInsertionTypeHeaders"]
							}
							if oHttpHeaderInsertionTypeHeaders.Header != nil {
								nestedHttpHeaderInsertionTypeHeaders.Header = oHttpHeaderInsertionTypeHeaders.Header
							}
							if oHttpHeaderInsertionTypeHeaders.Value != nil {
								nestedHttpHeaderInsertionTypeHeaders.Value = oHttpHeaderInsertionTypeHeaders.Value
							}
							if oHttpHeaderInsertionTypeHeaders.Log != nil {
								nestedHttpHeaderInsertionTypeHeaders.Log = util.YesNo(oHttpHeaderInsertionTypeHeaders.Log, nil)
							}
							if oHttpHeaderInsertionTypeHeaders.Name != "" {
								nestedHttpHeaderInsertionTypeHeaders.Name = oHttpHeaderInsertionTypeHeaders.Name
							}
							nestedHttpHeaderInsertionType.Headers = append(nestedHttpHeaderInsertionType.Headers, nestedHttpHeaderInsertionTypeHeaders)
						}
					}
					if oHttpHeaderInsertionType.Domains != nil {
						nestedHttpHeaderInsertionType.Domains = util.StrToMem(oHttpHeaderInsertionType.Domains)
					}
					if oHttpHeaderInsertionType.Name != "" {
						nestedHttpHeaderInsertionType.Name = oHttpHeaderInsertionType.Name
					}
					nestedHttpHeaderInsertion.Type = append(nestedHttpHeaderInsertion.Type, nestedHttpHeaderInsertionType)
				}
			}
			if oHttpHeaderInsertion.Name != "" {
				nestedHttpHeaderInsertion.Name = oHttpHeaderInsertion.Name
			}
			nestedHttpHeaderInsertionCol = append(nestedHttpHeaderInsertionCol, nestedHttpHeaderInsertion)
		}
		entry.HttpHeaderInsertion = nestedHttpHeaderInsertionCol
	}

	entry.LocalInlineCat = util.YesNo(o.LocalInlineCat, nil)
	entry.LogContainerPageOnly = util.YesNo(o.LogContainerPageOnly, nil)
	entry.LogHttpHdrReferer = util.YesNo(o.LogHttpHdrReferer, nil)
	entry.LogHttpHdrUserAgent = util.YesNo(o.LogHttpHdrUserAgent, nil)
	entry.LogHttpHdrXff = util.YesNo(o.LogHttpHdrXff, nil)
	entry.MlavCategoryException = util.StrToMem(o.MlavCategoryException)
	entry.Override = util.StrToMem(o.Override)
	entry.SafeSearchEnforcement = util.YesNo(o.SafeSearchEnforcement, nil)

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
		entry.Alert = util.MemToStr(o.Alert)
		entry.Allow = util.MemToStr(o.Allow)
		entry.Block = util.MemToStr(o.Block)
		entry.CloudInlineCat = util.AsBool(o.CloudInlineCat, nil)
		entry.Continue = util.MemToStr(o.Continue)
		var nestedCredentialEnforcement *CredentialEnforcement
		if o.CredentialEnforcement != nil {
			nestedCredentialEnforcement = &CredentialEnforcement{}
			if o.CredentialEnforcement.Misc != nil {
				entry.Misc["CredentialEnforcement"] = o.CredentialEnforcement.Misc
			}
			if o.CredentialEnforcement.Alert != nil {
				nestedCredentialEnforcement.Alert = util.MemToStr(o.CredentialEnforcement.Alert)
			}
			if o.CredentialEnforcement.Allow != nil {
				nestedCredentialEnforcement.Allow = util.MemToStr(o.CredentialEnforcement.Allow)
			}
			if o.CredentialEnforcement.Block != nil {
				nestedCredentialEnforcement.Block = util.MemToStr(o.CredentialEnforcement.Block)
			}
			if o.CredentialEnforcement.Continue != nil {
				nestedCredentialEnforcement.Continue = util.MemToStr(o.CredentialEnforcement.Continue)
			}
			if o.CredentialEnforcement.LogSeverity != nil {
				nestedCredentialEnforcement.LogSeverity = o.CredentialEnforcement.LogSeverity
			}
			if o.CredentialEnforcement.Mode != nil {
				nestedCredentialEnforcement.Mode = &CredentialEnforcementMode{}
				if o.CredentialEnforcement.Mode.Misc != nil {
					entry.Misc["CredentialEnforcementMode"] = o.CredentialEnforcement.Mode.Misc
				}
				if o.CredentialEnforcement.Mode.Disabled != nil {
					nestedCredentialEnforcement.Mode.Disabled = &CredentialEnforcementModeDisabled{}
					if o.CredentialEnforcement.Mode.Disabled.Misc != nil {
						entry.Misc["CredentialEnforcementModeDisabled"] = o.CredentialEnforcement.Mode.Disabled.Misc
					}
				}
				if o.CredentialEnforcement.Mode.DomainCredentials != nil {
					nestedCredentialEnforcement.Mode.DomainCredentials = &CredentialEnforcementModeDomainCredentials{}
					if o.CredentialEnforcement.Mode.DomainCredentials.Misc != nil {
						entry.Misc["CredentialEnforcementModeDomainCredentials"] = o.CredentialEnforcement.Mode.DomainCredentials.Misc
					}
				}
				if o.CredentialEnforcement.Mode.GroupMapping != nil {
					nestedCredentialEnforcement.Mode.GroupMapping = o.CredentialEnforcement.Mode.GroupMapping
				}
				if o.CredentialEnforcement.Mode.IpUser != nil {
					nestedCredentialEnforcement.Mode.IpUser = &CredentialEnforcementModeIpUser{}
					if o.CredentialEnforcement.Mode.IpUser.Misc != nil {
						entry.Misc["CredentialEnforcementModeIpUser"] = o.CredentialEnforcement.Mode.IpUser.Misc
					}
				}
			}
		}
		entry.CredentialEnforcement = nestedCredentialEnforcement

		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.EnableContainerPage = util.AsBool(o.EnableContainerPage, nil)
		var nestedHttpHeaderInsertionCol []HttpHeaderInsertion
		if o.HttpHeaderInsertion != nil {
			nestedHttpHeaderInsertionCol = []HttpHeaderInsertion{}
			for _, oHttpHeaderInsertion := range o.HttpHeaderInsertion {
				nestedHttpHeaderInsertion := HttpHeaderInsertion{}
				if oHttpHeaderInsertion.Misc != nil {
					entry.Misc["HttpHeaderInsertion"] = oHttpHeaderInsertion.Misc
				}
				if oHttpHeaderInsertion.DisableOverride != nil {
					nestedHttpHeaderInsertion.DisableOverride = oHttpHeaderInsertion.DisableOverride
				}
				if oHttpHeaderInsertion.Type != nil {
					nestedHttpHeaderInsertion.Type = []HttpHeaderInsertionType{}
					for _, oHttpHeaderInsertionType := range oHttpHeaderInsertion.Type {
						nestedHttpHeaderInsertionType := HttpHeaderInsertionType{}
						if oHttpHeaderInsertionType.Misc != nil {
							entry.Misc["HttpHeaderInsertionType"] = oHttpHeaderInsertionType.Misc
						}
						if oHttpHeaderInsertionType.Headers != nil {
							nestedHttpHeaderInsertionType.Headers = []HttpHeaderInsertionTypeHeaders{}
							for _, oHttpHeaderInsertionTypeHeaders := range oHttpHeaderInsertionType.Headers {
								nestedHttpHeaderInsertionTypeHeaders := HttpHeaderInsertionTypeHeaders{}
								if oHttpHeaderInsertionTypeHeaders.Misc != nil {
									entry.Misc["HttpHeaderInsertionTypeHeaders"] = oHttpHeaderInsertionTypeHeaders.Misc
								}
								if oHttpHeaderInsertionTypeHeaders.Header != nil {
									nestedHttpHeaderInsertionTypeHeaders.Header = oHttpHeaderInsertionTypeHeaders.Header
								}
								if oHttpHeaderInsertionTypeHeaders.Value != nil {
									nestedHttpHeaderInsertionTypeHeaders.Value = oHttpHeaderInsertionTypeHeaders.Value
								}
								if oHttpHeaderInsertionTypeHeaders.Log != nil {
									nestedHttpHeaderInsertionTypeHeaders.Log = util.AsBool(oHttpHeaderInsertionTypeHeaders.Log, nil)
								}
								if oHttpHeaderInsertionTypeHeaders.Name != "" {
									nestedHttpHeaderInsertionTypeHeaders.Name = oHttpHeaderInsertionTypeHeaders.Name
								}
								nestedHttpHeaderInsertionType.Headers = append(nestedHttpHeaderInsertionType.Headers, nestedHttpHeaderInsertionTypeHeaders)
							}
						}
						if oHttpHeaderInsertionType.Domains != nil {
							nestedHttpHeaderInsertionType.Domains = util.MemToStr(oHttpHeaderInsertionType.Domains)
						}
						if oHttpHeaderInsertionType.Name != "" {
							nestedHttpHeaderInsertionType.Name = oHttpHeaderInsertionType.Name
						}
						nestedHttpHeaderInsertion.Type = append(nestedHttpHeaderInsertion.Type, nestedHttpHeaderInsertionType)
					}
				}
				if oHttpHeaderInsertion.Name != "" {
					nestedHttpHeaderInsertion.Name = oHttpHeaderInsertion.Name
				}
				nestedHttpHeaderInsertionCol = append(nestedHttpHeaderInsertionCol, nestedHttpHeaderInsertion)
			}
			entry.HttpHeaderInsertion = nestedHttpHeaderInsertionCol
		}

		entry.LocalInlineCat = util.AsBool(o.LocalInlineCat, nil)
		entry.LogContainerPageOnly = util.AsBool(o.LogContainerPageOnly, nil)
		entry.LogHttpHdrReferer = util.AsBool(o.LogHttpHdrReferer, nil)
		entry.LogHttpHdrUserAgent = util.AsBool(o.LogHttpHdrUserAgent, nil)
		entry.LogHttpHdrXff = util.AsBool(o.LogHttpHdrXff, nil)
		entry.MlavCategoryException = util.MemToStr(o.MlavCategoryException)
		entry.Override = util.MemToStr(o.Override)
		entry.SafeSearchEnforcement = util.AsBool(o.SafeSearchEnforcement, nil)

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
	if !util.OrderedListsMatch(a.Alert, b.Alert) {
		return false
	}
	if !util.OrderedListsMatch(a.Allow, b.Allow) {
		return false
	}
	if !util.OrderedListsMatch(a.Block, b.Block) {
		return false
	}
	if !util.BoolsMatch(a.CloudInlineCat, b.CloudInlineCat) {
		return false
	}
	if !util.OrderedListsMatch(a.Continue, b.Continue) {
		return false
	}
	if !matchCredentialEnforcement(a.CredentialEnforcement, b.CredentialEnforcement) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(a.EnableContainerPage, b.EnableContainerPage) {
		return false
	}
	if !matchHttpHeaderInsertion(a.HttpHeaderInsertion, b.HttpHeaderInsertion) {
		return false
	}
	if !util.BoolsMatch(a.LocalInlineCat, b.LocalInlineCat) {
		return false
	}
	if !util.BoolsMatch(a.LogContainerPageOnly, b.LogContainerPageOnly) {
		return false
	}
	if !util.BoolsMatch(a.LogHttpHdrReferer, b.LogHttpHdrReferer) {
		return false
	}
	if !util.BoolsMatch(a.LogHttpHdrUserAgent, b.LogHttpHdrUserAgent) {
		return false
	}
	if !util.BoolsMatch(a.LogHttpHdrXff, b.LogHttpHdrXff) {
		return false
	}
	if !util.OrderedListsMatch(a.MlavCategoryException, b.MlavCategoryException) {
		return false
	}
	if !util.OrderedListsMatch(a.Override, b.Override) {
		return false
	}
	if !util.BoolsMatch(a.SafeSearchEnforcement, b.SafeSearchEnforcement) {
		return false
	}

	return true
}

func matchHttpHeaderInsertionTypeHeaders(a []HttpHeaderInsertionTypeHeaders, b []HttpHeaderInsertionTypeHeaders) bool {
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
			if !util.StringsMatch(a.Header, b.Header) {
				return false
			}
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
			if !util.BoolsMatch(a.Log, b.Log) {
				return false
			}
		}
	}
	return true
}
func matchHttpHeaderInsertionType(a []HttpHeaderInsertionType, b []HttpHeaderInsertionType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchHttpHeaderInsertionTypeHeaders(a.Headers, b.Headers) {
				return false
			}
			if !util.OrderedListsMatch(a.Domains, b.Domains) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchHttpHeaderInsertion(a []HttpHeaderInsertion, b []HttpHeaderInsertion) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
				return false
			}
			if !matchHttpHeaderInsertionType(a.Type, b.Type) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchCredentialEnforcementModeDomainCredentials(a *CredentialEnforcementModeDomainCredentials, b *CredentialEnforcementModeDomainCredentials) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchCredentialEnforcementModeIpUser(a *CredentialEnforcementModeIpUser, b *CredentialEnforcementModeIpUser) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchCredentialEnforcementModeDisabled(a *CredentialEnforcementModeDisabled, b *CredentialEnforcementModeDisabled) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchCredentialEnforcementMode(a *CredentialEnforcementMode, b *CredentialEnforcementMode) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchCredentialEnforcementModeDisabled(a.Disabled, b.Disabled) {
		return false
	}
	if !matchCredentialEnforcementModeDomainCredentials(a.DomainCredentials, b.DomainCredentials) {
		return false
	}
	if !util.StringsMatch(a.GroupMapping, b.GroupMapping) {
		return false
	}
	if !matchCredentialEnforcementModeIpUser(a.IpUser, b.IpUser) {
		return false
	}
	return true
}
func matchCredentialEnforcement(a *CredentialEnforcement, b *CredentialEnforcement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Alert, b.Alert) {
		return false
	}
	if !util.OrderedListsMatch(a.Allow, b.Allow) {
		return false
	}
	if !util.OrderedListsMatch(a.Block, b.Block) {
		return false
	}
	if !util.OrderedListsMatch(a.Continue, b.Continue) {
		return false
	}
	if !util.StringsMatch(a.LogSeverity, b.LogSeverity) {
		return false
	}
	if !matchCredentialEnforcementMode(a.Mode, b.Mode) {
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
