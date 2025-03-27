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
	Suffix = []string{"log-settings", "profiles"}
)

type Entry struct {
	Name                       string
	Description                *string
	DisableOverride            *string
	EnhancedApplicationLogging *bool
	MatchList                  []MatchList

	Misc map[string][]generic.Xml
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
}
type MatchListActions struct {
	Name string
	Type *MatchListActionsType
}
type MatchListActionsType struct {
	Integration *MatchListActionsTypeIntegration
	Tagging     *MatchListActionsTypeTagging
}
type MatchListActionsTypeIntegration struct {
	Action *string
}
type MatchListActionsTypeTagging struct {
	Target       *string
	Action       *string
	Timeout      *int64
	Registration *MatchListActionsTypeTaggingRegistration
	Tags         []string
}
type MatchListActionsTypeTaggingRegistration struct {
	Localhost *MatchListActionsTypeTaggingRegistrationLocalhost
	Panorama  *MatchListActionsTypeTaggingRegistrationPanorama
	Remote    *MatchListActionsTypeTaggingRegistrationRemote
}
type MatchListActionsTypeTaggingRegistrationLocalhost struct {
}
type MatchListActionsTypeTaggingRegistrationPanorama struct {
}
type MatchListActionsTypeTaggingRegistrationRemote struct {
	HttpProfile *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                    xml.Name       `xml:"entry"`
	Name                       string         `xml:"name,attr"`
	Description                *string        `xml:"description,omitempty"`
	DisableOverride            *string        `xml:"disable-override,omitempty"`
	EnhancedApplicationLogging *string        `xml:"enhanced-application-logging,omitempty"`
	MatchList                  []MatchListXml `xml:"match-list>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListXml struct {
	ActionDesc     *string               `xml:"action-desc,omitempty"`
	Actions        []MatchListActionsXml `xml:"actions>entry,omitempty"`
	Filter         *string               `xml:"filter,omitempty"`
	LogType        *string               `xml:"log-type,omitempty"`
	Name           string                `xml:"name,attr"`
	Quarantine     *string               `xml:"quarantine,omitempty"`
	SendEmail      *util.MemberType      `xml:"send-email,omitempty"`
	SendHttp       *util.MemberType      `xml:"send-http,omitempty"`
	SendSnmptrap   *util.MemberType      `xml:"send-snmptrap,omitempty"`
	SendSyslog     *util.MemberType      `xml:"send-syslog,omitempty"`
	SendToPanorama *string               `xml:"send-to-panorama,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsXml struct {
	Name string                   `xml:"name,attr"`
	Type *MatchListActionsTypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeXml struct {
	Integration *MatchListActionsTypeIntegrationXml `xml:"integration,omitempty"`
	Tagging     *MatchListActionsTypeTaggingXml     `xml:"tagging,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeIntegrationXml struct {
	Action *string `xml:"action,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeTaggingXml struct {
	Action       *string                                     `xml:"action,omitempty"`
	Registration *MatchListActionsTypeTaggingRegistrationXml `xml:"registration,omitempty"`
	Tags         *util.MemberType                            `xml:"tags,omitempty"`
	Target       *string                                     `xml:"target,omitempty"`
	Timeout      *int64                                      `xml:"timeout,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeTaggingRegistrationXml struct {
	Localhost *MatchListActionsTypeTaggingRegistrationLocalhostXml `xml:"localhost,omitempty"`
	Panorama  *MatchListActionsTypeTaggingRegistrationPanoramaXml  `xml:"panorama,omitempty"`
	Remote    *MatchListActionsTypeTaggingRegistrationRemoteXml    `xml:"remote,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeTaggingRegistrationLocalhostXml struct {
	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeTaggingRegistrationPanoramaXml struct {
	Misc []generic.Xml `xml:",any"`
}
type MatchListActionsTypeTaggingRegistrationRemoteXml struct {
	HttpProfile *string `xml:"http-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.EnhancedApplicationLogging = util.YesNo(o.EnhancedApplicationLogging, nil)
	var nestedMatchListCol []MatchListXml
	if o.MatchList != nil {
		nestedMatchListCol = []MatchListXml{}
		for _, oMatchList := range o.MatchList {
			nestedMatchList := MatchListXml{}
			if _, ok := o.Misc["MatchList"]; ok {
				nestedMatchList.Misc = o.Misc["MatchList"]
			}
			if oMatchList.Name != "" {
				nestedMatchList.Name = oMatchList.Name
			}
			if oMatchList.ActionDesc != nil {
				nestedMatchList.ActionDesc = oMatchList.ActionDesc
			}
			if oMatchList.LogType != nil {
				nestedMatchList.LogType = oMatchList.LogType
			}
			if oMatchList.Filter != nil {
				nestedMatchList.Filter = oMatchList.Filter
			}
			if oMatchList.SendToPanorama != nil {
				nestedMatchList.SendToPanorama = util.YesNo(oMatchList.SendToPanorama, nil)
			}
			if oMatchList.Quarantine != nil {
				nestedMatchList.Quarantine = util.YesNo(oMatchList.Quarantine, nil)
			}
			if oMatchList.SendSnmptrap != nil {
				nestedMatchList.SendSnmptrap = util.StrToMem(oMatchList.SendSnmptrap)
			}
			if oMatchList.SendEmail != nil {
				nestedMatchList.SendEmail = util.StrToMem(oMatchList.SendEmail)
			}
			if oMatchList.SendSyslog != nil {
				nestedMatchList.SendSyslog = util.StrToMem(oMatchList.SendSyslog)
			}
			if oMatchList.SendHttp != nil {
				nestedMatchList.SendHttp = util.StrToMem(oMatchList.SendHttp)
			}
			if oMatchList.Actions != nil {
				nestedMatchList.Actions = []MatchListActionsXml{}
				for _, oMatchListActions := range oMatchList.Actions {
					nestedMatchListActions := MatchListActionsXml{}
					if _, ok := o.Misc["MatchListActions"]; ok {
						nestedMatchListActions.Misc = o.Misc["MatchListActions"]
					}
					if oMatchListActions.Name != "" {
						nestedMatchListActions.Name = oMatchListActions.Name
					}
					if oMatchListActions.Type != nil {
						nestedMatchListActions.Type = &MatchListActionsTypeXml{}
						if _, ok := o.Misc["MatchListActionsType"]; ok {
							nestedMatchListActions.Type.Misc = o.Misc["MatchListActionsType"]
						}
						if oMatchListActions.Type.Integration != nil {
							nestedMatchListActions.Type.Integration = &MatchListActionsTypeIntegrationXml{}
							if _, ok := o.Misc["MatchListActionsTypeIntegration"]; ok {
								nestedMatchListActions.Type.Integration.Misc = o.Misc["MatchListActionsTypeIntegration"]
							}
							if oMatchListActions.Type.Integration.Action != nil {
								nestedMatchListActions.Type.Integration.Action = oMatchListActions.Type.Integration.Action
							}
						}
						if oMatchListActions.Type.Tagging != nil {
							nestedMatchListActions.Type.Tagging = &MatchListActionsTypeTaggingXml{}
							if _, ok := o.Misc["MatchListActionsTypeTagging"]; ok {
								nestedMatchListActions.Type.Tagging.Misc = o.Misc["MatchListActionsTypeTagging"]
							}
							if oMatchListActions.Type.Tagging.Target != nil {
								nestedMatchListActions.Type.Tagging.Target = oMatchListActions.Type.Tagging.Target
							}
							if oMatchListActions.Type.Tagging.Action != nil {
								nestedMatchListActions.Type.Tagging.Action = oMatchListActions.Type.Tagging.Action
							}
							if oMatchListActions.Type.Tagging.Timeout != nil {
								nestedMatchListActions.Type.Tagging.Timeout = oMatchListActions.Type.Tagging.Timeout
							}
							if oMatchListActions.Type.Tagging.Registration != nil {
								nestedMatchListActions.Type.Tagging.Registration = &MatchListActionsTypeTaggingRegistrationXml{}
								if _, ok := o.Misc["MatchListActionsTypeTaggingRegistration"]; ok {
									nestedMatchListActions.Type.Tagging.Registration.Misc = o.Misc["MatchListActionsTypeTaggingRegistration"]
								}
								if oMatchListActions.Type.Tagging.Registration.Localhost != nil {
									nestedMatchListActions.Type.Tagging.Registration.Localhost = &MatchListActionsTypeTaggingRegistrationLocalhostXml{}
									if _, ok := o.Misc["MatchListActionsTypeTaggingRegistrationLocalhost"]; ok {
										nestedMatchListActions.Type.Tagging.Registration.Localhost.Misc = o.Misc["MatchListActionsTypeTaggingRegistrationLocalhost"]
									}
								}
								if oMatchListActions.Type.Tagging.Registration.Panorama != nil {
									nestedMatchListActions.Type.Tagging.Registration.Panorama = &MatchListActionsTypeTaggingRegistrationPanoramaXml{}
									if _, ok := o.Misc["MatchListActionsTypeTaggingRegistrationPanorama"]; ok {
										nestedMatchListActions.Type.Tagging.Registration.Panorama.Misc = o.Misc["MatchListActionsTypeTaggingRegistrationPanorama"]
									}
								}
								if oMatchListActions.Type.Tagging.Registration.Remote != nil {
									nestedMatchListActions.Type.Tagging.Registration.Remote = &MatchListActionsTypeTaggingRegistrationRemoteXml{}
									if _, ok := o.Misc["MatchListActionsTypeTaggingRegistrationRemote"]; ok {
										nestedMatchListActions.Type.Tagging.Registration.Remote.Misc = o.Misc["MatchListActionsTypeTaggingRegistrationRemote"]
									}
									if oMatchListActions.Type.Tagging.Registration.Remote.HttpProfile != nil {
										nestedMatchListActions.Type.Tagging.Registration.Remote.HttpProfile = oMatchListActions.Type.Tagging.Registration.Remote.HttpProfile
									}
								}
							}
							if oMatchListActions.Type.Tagging.Tags != nil {
								nestedMatchListActions.Type.Tagging.Tags = util.StrToMem(oMatchListActions.Type.Tagging.Tags)
							}
						}
					}
					nestedMatchList.Actions = append(nestedMatchList.Actions, nestedMatchListActions)
				}
			}
			nestedMatchListCol = append(nestedMatchListCol, nestedMatchList)
		}
		entry.MatchList = nestedMatchListCol
	}

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
		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.EnhancedApplicationLogging = util.AsBool(o.EnhancedApplicationLogging, nil)
		var nestedMatchListCol []MatchList
		if o.MatchList != nil {
			nestedMatchListCol = []MatchList{}
			for _, oMatchList := range o.MatchList {
				nestedMatchList := MatchList{}
				if oMatchList.Misc != nil {
					entry.Misc["MatchList"] = oMatchList.Misc
				}
				if oMatchList.Name != "" {
					nestedMatchList.Name = oMatchList.Name
				}
				if oMatchList.ActionDesc != nil {
					nestedMatchList.ActionDesc = oMatchList.ActionDesc
				}
				if oMatchList.LogType != nil {
					nestedMatchList.LogType = oMatchList.LogType
				}
				if oMatchList.Filter != nil {
					nestedMatchList.Filter = oMatchList.Filter
				}
				if oMatchList.SendToPanorama != nil {
					nestedMatchList.SendToPanorama = util.AsBool(oMatchList.SendToPanorama, nil)
				}
				if oMatchList.Quarantine != nil {
					nestedMatchList.Quarantine = util.AsBool(oMatchList.Quarantine, nil)
				}
				if oMatchList.SendSnmptrap != nil {
					nestedMatchList.SendSnmptrap = util.MemToStr(oMatchList.SendSnmptrap)
				}
				if oMatchList.SendEmail != nil {
					nestedMatchList.SendEmail = util.MemToStr(oMatchList.SendEmail)
				}
				if oMatchList.SendSyslog != nil {
					nestedMatchList.SendSyslog = util.MemToStr(oMatchList.SendSyslog)
				}
				if oMatchList.SendHttp != nil {
					nestedMatchList.SendHttp = util.MemToStr(oMatchList.SendHttp)
				}
				if oMatchList.Actions != nil {
					nestedMatchList.Actions = []MatchListActions{}
					for _, oMatchListActions := range oMatchList.Actions {
						nestedMatchListActions := MatchListActions{}
						if oMatchListActions.Misc != nil {
							entry.Misc["MatchListActions"] = oMatchListActions.Misc
						}
						if oMatchListActions.Name != "" {
							nestedMatchListActions.Name = oMatchListActions.Name
						}
						if oMatchListActions.Type != nil {
							nestedMatchListActions.Type = &MatchListActionsType{}
							if oMatchListActions.Type.Misc != nil {
								entry.Misc["MatchListActionsType"] = oMatchListActions.Type.Misc
							}
							if oMatchListActions.Type.Integration != nil {
								nestedMatchListActions.Type.Integration = &MatchListActionsTypeIntegration{}
								if oMatchListActions.Type.Integration.Misc != nil {
									entry.Misc["MatchListActionsTypeIntegration"] = oMatchListActions.Type.Integration.Misc
								}
								if oMatchListActions.Type.Integration.Action != nil {
									nestedMatchListActions.Type.Integration.Action = oMatchListActions.Type.Integration.Action
								}
							}
							if oMatchListActions.Type.Tagging != nil {
								nestedMatchListActions.Type.Tagging = &MatchListActionsTypeTagging{}
								if oMatchListActions.Type.Tagging.Misc != nil {
									entry.Misc["MatchListActionsTypeTagging"] = oMatchListActions.Type.Tagging.Misc
								}
								if oMatchListActions.Type.Tagging.Target != nil {
									nestedMatchListActions.Type.Tagging.Target = oMatchListActions.Type.Tagging.Target
								}
								if oMatchListActions.Type.Tagging.Action != nil {
									nestedMatchListActions.Type.Tagging.Action = oMatchListActions.Type.Tagging.Action
								}
								if oMatchListActions.Type.Tagging.Timeout != nil {
									nestedMatchListActions.Type.Tagging.Timeout = oMatchListActions.Type.Tagging.Timeout
								}
								if oMatchListActions.Type.Tagging.Registration != nil {
									nestedMatchListActions.Type.Tagging.Registration = &MatchListActionsTypeTaggingRegistration{}
									if oMatchListActions.Type.Tagging.Registration.Misc != nil {
										entry.Misc["MatchListActionsTypeTaggingRegistration"] = oMatchListActions.Type.Tagging.Registration.Misc
									}
									if oMatchListActions.Type.Tagging.Registration.Localhost != nil {
										nestedMatchListActions.Type.Tagging.Registration.Localhost = &MatchListActionsTypeTaggingRegistrationLocalhost{}
										if oMatchListActions.Type.Tagging.Registration.Localhost.Misc != nil {
											entry.Misc["MatchListActionsTypeTaggingRegistrationLocalhost"] = oMatchListActions.Type.Tagging.Registration.Localhost.Misc
										}
									}
									if oMatchListActions.Type.Tagging.Registration.Panorama != nil {
										nestedMatchListActions.Type.Tagging.Registration.Panorama = &MatchListActionsTypeTaggingRegistrationPanorama{}
										if oMatchListActions.Type.Tagging.Registration.Panorama.Misc != nil {
											entry.Misc["MatchListActionsTypeTaggingRegistrationPanorama"] = oMatchListActions.Type.Tagging.Registration.Panorama.Misc
										}
									}
									if oMatchListActions.Type.Tagging.Registration.Remote != nil {
										nestedMatchListActions.Type.Tagging.Registration.Remote = &MatchListActionsTypeTaggingRegistrationRemote{}
										if oMatchListActions.Type.Tagging.Registration.Remote.Misc != nil {
											entry.Misc["MatchListActionsTypeTaggingRegistrationRemote"] = oMatchListActions.Type.Tagging.Registration.Remote.Misc
										}
										if oMatchListActions.Type.Tagging.Registration.Remote.HttpProfile != nil {
											nestedMatchListActions.Type.Tagging.Registration.Remote.HttpProfile = oMatchListActions.Type.Tagging.Registration.Remote.HttpProfile
										}
									}
								}
								if oMatchListActions.Type.Tagging.Tags != nil {
									nestedMatchListActions.Type.Tagging.Tags = util.MemToStr(oMatchListActions.Type.Tagging.Tags)
								}
							}
						}
						nestedMatchList.Actions = append(nestedMatchList.Actions, nestedMatchListActions)
					}
				}
				nestedMatchListCol = append(nestedMatchListCol, nestedMatchList)
			}
			entry.MatchList = nestedMatchListCol
		}

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
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(a.EnhancedApplicationLogging, b.EnhancedApplicationLogging) {
		return false
	}
	if !matchMatchList(a.MatchList, b.MatchList) {
		return false
	}

	return true
}

func matchMatchListActionsTypeIntegration(a *MatchListActionsTypeIntegration, b *MatchListActionsTypeIntegration) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	return true
}
func matchMatchListActionsTypeTaggingRegistrationLocalhost(a *MatchListActionsTypeTaggingRegistrationLocalhost, b *MatchListActionsTypeTaggingRegistrationLocalhost) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchMatchListActionsTypeTaggingRegistrationPanorama(a *MatchListActionsTypeTaggingRegistrationPanorama, b *MatchListActionsTypeTaggingRegistrationPanorama) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchMatchListActionsTypeTaggingRegistrationRemote(a *MatchListActionsTypeTaggingRegistrationRemote, b *MatchListActionsTypeTaggingRegistrationRemote) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.HttpProfile, b.HttpProfile) {
		return false
	}
	return true
}
func matchMatchListActionsTypeTaggingRegistration(a *MatchListActionsTypeTaggingRegistration, b *MatchListActionsTypeTaggingRegistration) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchMatchListActionsTypeTaggingRegistrationLocalhost(a.Localhost, b.Localhost) {
		return false
	}
	if !matchMatchListActionsTypeTaggingRegistrationPanorama(a.Panorama, b.Panorama) {
		return false
	}
	if !matchMatchListActionsTypeTaggingRegistrationRemote(a.Remote, b.Remote) {
		return false
	}
	return true
}
func matchMatchListActionsTypeTagging(a *MatchListActionsTypeTagging, b *MatchListActionsTypeTagging) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Target, b.Target) {
		return false
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.Timeout, b.Timeout) {
		return false
	}
	if !matchMatchListActionsTypeTaggingRegistration(a.Registration, b.Registration) {
		return false
	}
	if !util.OrderedListsMatch(a.Tags, b.Tags) {
		return false
	}
	return true
}
func matchMatchListActionsType(a *MatchListActionsType, b *MatchListActionsType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchMatchListActionsTypeIntegration(a.Integration, b.Integration) {
		return false
	}
	if !matchMatchListActionsTypeTagging(a.Tagging, b.Tagging) {
		return false
	}
	return true
}
func matchMatchListActions(a []MatchListActions, b []MatchListActions) bool {
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
			if !matchMatchListActionsType(a.Type, b.Type) {
				return false
			}
		}
	}
	return true
}
func matchMatchList(a []MatchList, b []MatchList) bool {
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
			if !util.StringsMatch(a.ActionDesc, b.ActionDesc) {
				return false
			}
			if !util.StringsMatch(a.LogType, b.LogType) {
				return false
			}
			if !util.StringsMatch(a.Filter, b.Filter) {
				return false
			}
			if !util.BoolsMatch(a.SendToPanorama, b.SendToPanorama) {
				return false
			}
			if !util.BoolsMatch(a.Quarantine, b.Quarantine) {
				return false
			}
			if !util.OrderedListsMatch(a.SendSnmptrap, b.SendSnmptrap) {
				return false
			}
			if !util.OrderedListsMatch(a.SendEmail, b.SendEmail) {
				return false
			}
			if !util.OrderedListsMatch(a.SendSyslog, b.SendSyslog) {
				return false
			}
			if !util.OrderedListsMatch(a.SendHttp, b.SendHttp) {
				return false
			}
			if !matchMatchListActions(a.Actions, b.Actions) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
