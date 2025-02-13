package application

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
	Suffix = []string{}
)

type Entry struct {
	Name                   string
	AbleToTransferFile     *bool
	AlgDisableCapability   *string
	Category               *string
	ConsumeBigBandwidth    *bool
	DataIdent              *bool
	Default                *Default
	Description            *string
	DisableOverride        *string
	EvasiveBehavior        *bool
	FileTypeIdent          *bool
	HasKnownVulnerability  *bool
	NoAppidCaching         *bool
	ParentApp              *string
	PervasiveUse           *bool
	ProneToMisuse          *bool
	Risk                   *int64
	Signature              []Signature
	Subcategory            *string
	TcpHalfClosedTimeout   *int64
	TcpTimeWaitTimeout     *int64
	TcpTimeout             *int64
	Technology             *string
	Timeout                *int64
	TunnelApplications     *bool
	TunnelOtherApplication *bool
	UdpTimeout             *int64
	UsedByMalware          *bool
	VirusIdent             *bool

	Misc map[string][]generic.Xml
}

type Default struct {
	IdentByIcmpType   *DefaultIdentByIcmpType
	IdentByIcmp6Type  *DefaultIdentByIcmp6Type
	IdentByIpProtocol *string
	Port              []string
}
type DefaultIdentByIcmp6Type struct {
	Code *string
	Type *string
}
type DefaultIdentByIcmpType struct {
	Code *string
	Type *string
}
type Signature struct {
	AndCondition []SignatureAndCondition
	Comment      *string
	Name         string
	OrderFree    *bool
	Scope        *string
}
type SignatureAndCondition struct {
	Name        string
	OrCondition []SignatureAndConditionOrCondition
}
type SignatureAndConditionOrCondition struct {
	Name     string
	Operator *SignatureAndConditionOrConditionOperator
}
type SignatureAndConditionOrConditionOperator struct {
	EqualTo      *SignatureAndConditionOrConditionOperatorEqualTo
	GreaterThan  *SignatureAndConditionOrConditionOperatorGreaterThan
	LessThan     *SignatureAndConditionOrConditionOperatorLessThan
	PatternMatch *SignatureAndConditionOrConditionOperatorPatternMatch
}
type SignatureAndConditionOrConditionOperatorEqualTo struct {
	Context  *string
	Mask     *string
	Position *string
	Value    *string
}
type SignatureAndConditionOrConditionOperatorGreaterThan struct {
	Context   *string
	Qualifier []SignatureAndConditionOrConditionOperatorGreaterThanQualifier
	Value     *int64
}
type SignatureAndConditionOrConditionOperatorGreaterThanQualifier struct {
	Name  string
	Value *string
}
type SignatureAndConditionOrConditionOperatorLessThan struct {
	Context   *string
	Qualifier []SignatureAndConditionOrConditionOperatorLessThanQualifier
	Value     *int64
}
type SignatureAndConditionOrConditionOperatorLessThanQualifier struct {
	Name  string
	Value *string
}
type SignatureAndConditionOrConditionOperatorPatternMatch struct {
	Context   *string
	Pattern   *string
	Qualifier []SignatureAndConditionOrConditionOperatorPatternMatchQualifier
}
type SignatureAndConditionOrConditionOperatorPatternMatchQualifier struct {
	Name  string
	Value *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                xml.Name       `xml:"entry"`
	Name                   string         `xml:"name,attr"`
	AbleToTransferFile     *string        `xml:"able-to-transfer-file,omitempty"`
	AlgDisableCapability   *string        `xml:"alg-disable-capability,omitempty"`
	Category               *string        `xml:"category,omitempty"`
	ConsumeBigBandwidth    *string        `xml:"consume-big-bandwidth,omitempty"`
	DataIdent              *string        `xml:"data-ident,omitempty"`
	Default                *DefaultXml    `xml:"default,omitempty"`
	Description            *string        `xml:"description,omitempty"`
	DisableOverride        *string        `xml:"disable-override,omitempty"`
	EvasiveBehavior        *string        `xml:"evasive-behavior,omitempty"`
	FileTypeIdent          *string        `xml:"file-type-ident,omitempty"`
	HasKnownVulnerability  *string        `xml:"has-known-vulnerability,omitempty"`
	NoAppidCaching         *string        `xml:"no-appid-caching,omitempty"`
	ParentApp              *string        `xml:"parent-app,omitempty"`
	PervasiveUse           *string        `xml:"pervasive-use,omitempty"`
	ProneToMisuse          *string        `xml:"prone-to-misuse,omitempty"`
	Risk                   *int64         `xml:"risk,omitempty"`
	Signature              []SignatureXml `xml:"signature>entry,omitempty"`
	Subcategory            *string        `xml:"subcategory,omitempty"`
	TcpHalfClosedTimeout   *int64         `xml:"tcp-half-closed-timeout,omitempty"`
	TcpTimeWaitTimeout     *int64         `xml:"tcp-time-wait-timeout,omitempty"`
	TcpTimeout             *int64         `xml:"tcp-timeout,omitempty"`
	Technology             *string        `xml:"technology,omitempty"`
	Timeout                *int64         `xml:"timeout,omitempty"`
	TunnelApplications     *string        `xml:"tunnel-applications,omitempty"`
	TunnelOtherApplication *string        `xml:"tunnel-other-application,omitempty"`
	UdpTimeout             *int64         `xml:"udp-timeout,omitempty"`
	UsedByMalware          *string        `xml:"used-by-malware,omitempty"`
	VirusIdent             *string        `xml:"virus-ident,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DefaultXml struct {
	IdentByIcmpType   *DefaultIdentByIcmpTypeXml  `xml:"ident-by-icmp-type,omitempty"`
	IdentByIcmp6Type  *DefaultIdentByIcmp6TypeXml `xml:"ident-by-icmp6-type,omitempty"`
	IdentByIpProtocol *string                     `xml:"ident-by-ip-protocol,omitempty"`
	Port              *util.MemberType            `xml:"port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DefaultIdentByIcmp6TypeXml struct {
	Code *string `xml:"code,omitempty"`
	Type *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DefaultIdentByIcmpTypeXml struct {
	Code *string `xml:"code,omitempty"`
	Type *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureXml struct {
	AndCondition []SignatureAndConditionXml `xml:"and-condition>entry,omitempty"`
	Comment      *string                    `xml:"comment,omitempty"`
	XMLName      xml.Name                   `xml:"entry"`
	Name         string                     `xml:"name,attr"`
	OrderFree    *string                    `xml:"order-free,omitempty"`
	Scope        *string                    `xml:"scope,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionXml struct {
	XMLName     xml.Name                              `xml:"entry"`
	Name        string                                `xml:"name,attr"`
	OrCondition []SignatureAndConditionOrConditionXml `xml:"or-condition>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionXml struct {
	XMLName  xml.Name                                     `xml:"entry"`
	Name     string                                       `xml:"name,attr"`
	Operator *SignatureAndConditionOrConditionOperatorXml `xml:"operator,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorXml struct {
	EqualTo      *SignatureAndConditionOrConditionOperatorEqualToXml      `xml:"equal-to,omitempty"`
	GreaterThan  *SignatureAndConditionOrConditionOperatorGreaterThanXml  `xml:"greater-than,omitempty"`
	LessThan     *SignatureAndConditionOrConditionOperatorLessThanXml     `xml:"less-than,omitempty"`
	PatternMatch *SignatureAndConditionOrConditionOperatorPatternMatchXml `xml:"pattern-match,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorEqualToXml struct {
	Context  *string `xml:"context,omitempty"`
	Mask     *string `xml:"mask,omitempty"`
	Position *string `xml:"position,omitempty"`
	Value    *string `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorGreaterThanXml struct {
	Context   *string                                                           `xml:"context,omitempty"`
	Qualifier []SignatureAndConditionOrConditionOperatorGreaterThanQualifierXml `xml:"qualifier>entry,omitempty"`
	Value     *int64                                                            `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorGreaterThanQualifierXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   *string  `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorLessThanXml struct {
	Context   *string                                                        `xml:"context,omitempty"`
	Qualifier []SignatureAndConditionOrConditionOperatorLessThanQualifierXml `xml:"qualifier>entry,omitempty"`
	Value     *int64                                                         `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorLessThanQualifierXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   *string  `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorPatternMatchXml struct {
	Context   *string                                                            `xml:"context,omitempty"`
	Pattern   *string                                                            `xml:"pattern,omitempty"`
	Qualifier []SignatureAndConditionOrConditionOperatorPatternMatchQualifierXml `xml:"qualifier>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type SignatureAndConditionOrConditionOperatorPatternMatchQualifierXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Value   *string  `xml:"value,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "able_to_transfer_file" || v == "AbleToTransferFile" {
		return e.AbleToTransferFile, nil
	}
	if v == "alg_disable_capability" || v == "AlgDisableCapability" {
		return e.AlgDisableCapability, nil
	}
	if v == "category" || v == "Category" {
		return e.Category, nil
	}
	if v == "consume_big_bandwidth" || v == "ConsumeBigBandwidth" {
		return e.ConsumeBigBandwidth, nil
	}
	if v == "data_ident" || v == "DataIdent" {
		return e.DataIdent, nil
	}
	if v == "default" || v == "Default" {
		return e.Default, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "evasive_behavior" || v == "EvasiveBehavior" {
		return e.EvasiveBehavior, nil
	}
	if v == "file_type_ident" || v == "FileTypeIdent" {
		return e.FileTypeIdent, nil
	}
	if v == "has_known_vulnerability" || v == "HasKnownVulnerability" {
		return e.HasKnownVulnerability, nil
	}
	if v == "no_appid_caching" || v == "NoAppidCaching" {
		return e.NoAppidCaching, nil
	}
	if v == "parent_app" || v == "ParentApp" {
		return e.ParentApp, nil
	}
	if v == "pervasive_use" || v == "PervasiveUse" {
		return e.PervasiveUse, nil
	}
	if v == "prone_to_misuse" || v == "ProneToMisuse" {
		return e.ProneToMisuse, nil
	}
	if v == "risk" || v == "Risk" {
		return e.Risk, nil
	}
	if v == "signature" || v == "Signature" {
		return e.Signature, nil
	}
	if v == "signature|LENGTH" || v == "Signature|LENGTH" {
		return int64(len(e.Signature)), nil
	}
	if v == "subcategory" || v == "Subcategory" {
		return e.Subcategory, nil
	}
	if v == "tcp_half_closed_timeout" || v == "TcpHalfClosedTimeout" {
		return e.TcpHalfClosedTimeout, nil
	}
	if v == "tcp_time_wait_timeout" || v == "TcpTimeWaitTimeout" {
		return e.TcpTimeWaitTimeout, nil
	}
	if v == "tcp_timeout" || v == "TcpTimeout" {
		return e.TcpTimeout, nil
	}
	if v == "technology" || v == "Technology" {
		return e.Technology, nil
	}
	if v == "timeout" || v == "Timeout" {
		return e.Timeout, nil
	}
	if v == "tunnel_applications" || v == "TunnelApplications" {
		return e.TunnelApplications, nil
	}
	if v == "tunnel_other_application" || v == "TunnelOtherApplication" {
		return e.TunnelOtherApplication, nil
	}
	if v == "udp_timeout" || v == "UdpTimeout" {
		return e.UdpTimeout, nil
	}
	if v == "used_by_malware" || v == "UsedByMalware" {
		return e.UsedByMalware, nil
	}
	if v == "virus_ident" || v == "VirusIdent" {
		return e.VirusIdent, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.AbleToTransferFile = util.YesNo(o.AbleToTransferFile, nil)
	entry.AlgDisableCapability = o.AlgDisableCapability
	entry.Category = o.Category
	entry.ConsumeBigBandwidth = util.YesNo(o.ConsumeBigBandwidth, nil)
	entry.DataIdent = util.YesNo(o.DataIdent, nil)
	var nestedDefault *DefaultXml
	if o.Default != nil {
		nestedDefault = &DefaultXml{}
		if _, ok := o.Misc["Default"]; ok {
			nestedDefault.Misc = o.Misc["Default"]
		}
		if o.Default.Port != nil {
			nestedDefault.Port = util.StrToMem(o.Default.Port)
		}
		if o.Default.IdentByIcmpType != nil {
			nestedDefault.IdentByIcmpType = &DefaultIdentByIcmpTypeXml{}
			if _, ok := o.Misc["DefaultIdentByIcmpType"]; ok {
				nestedDefault.IdentByIcmpType.Misc = o.Misc["DefaultIdentByIcmpType"]
			}
			if o.Default.IdentByIcmpType.Code != nil {
				nestedDefault.IdentByIcmpType.Code = o.Default.IdentByIcmpType.Code
			}
			if o.Default.IdentByIcmpType.Type != nil {
				nestedDefault.IdentByIcmpType.Type = o.Default.IdentByIcmpType.Type
			}
		}
		if o.Default.IdentByIcmp6Type != nil {
			nestedDefault.IdentByIcmp6Type = &DefaultIdentByIcmp6TypeXml{}
			if _, ok := o.Misc["DefaultIdentByIcmp6Type"]; ok {
				nestedDefault.IdentByIcmp6Type.Misc = o.Misc["DefaultIdentByIcmp6Type"]
			}
			if o.Default.IdentByIcmp6Type.Type != nil {
				nestedDefault.IdentByIcmp6Type.Type = o.Default.IdentByIcmp6Type.Type
			}
			if o.Default.IdentByIcmp6Type.Code != nil {
				nestedDefault.IdentByIcmp6Type.Code = o.Default.IdentByIcmp6Type.Code
			}
		}
		if o.Default.IdentByIpProtocol != nil {
			nestedDefault.IdentByIpProtocol = o.Default.IdentByIpProtocol
		}
	}
	entry.Default = nestedDefault

	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.EvasiveBehavior = util.YesNo(o.EvasiveBehavior, nil)
	entry.FileTypeIdent = util.YesNo(o.FileTypeIdent, nil)
	entry.HasKnownVulnerability = util.YesNo(o.HasKnownVulnerability, nil)
	entry.NoAppidCaching = util.YesNo(o.NoAppidCaching, nil)
	entry.ParentApp = o.ParentApp
	entry.PervasiveUse = util.YesNo(o.PervasiveUse, nil)
	entry.ProneToMisuse = util.YesNo(o.ProneToMisuse, nil)
	entry.Risk = o.Risk
	var nestedSignatureCol []SignatureXml
	if o.Signature != nil {
		nestedSignatureCol = []SignatureXml{}
		for _, oSignature := range o.Signature {
			nestedSignature := SignatureXml{}
			if _, ok := o.Misc["Signature"]; ok {
				nestedSignature.Misc = o.Misc["Signature"]
			}
			if oSignature.Comment != nil {
				nestedSignature.Comment = oSignature.Comment
			}
			if oSignature.Scope != nil {
				nestedSignature.Scope = oSignature.Scope
			}
			if oSignature.OrderFree != nil {
				nestedSignature.OrderFree = util.YesNo(oSignature.OrderFree, nil)
			}
			if oSignature.AndCondition != nil {
				nestedSignature.AndCondition = []SignatureAndConditionXml{}
				for _, oSignatureAndCondition := range oSignature.AndCondition {
					nestedSignatureAndCondition := SignatureAndConditionXml{}
					if _, ok := o.Misc["SignatureAndCondition"]; ok {
						nestedSignatureAndCondition.Misc = o.Misc["SignatureAndCondition"]
					}
					if oSignatureAndCondition.OrCondition != nil {
						nestedSignatureAndCondition.OrCondition = []SignatureAndConditionOrConditionXml{}
						for _, oSignatureAndConditionOrCondition := range oSignatureAndCondition.OrCondition {
							nestedSignatureAndConditionOrCondition := SignatureAndConditionOrConditionXml{}
							if _, ok := o.Misc["SignatureAndConditionOrCondition"]; ok {
								nestedSignatureAndConditionOrCondition.Misc = o.Misc["SignatureAndConditionOrCondition"]
							}
							if oSignatureAndConditionOrCondition.Operator != nil {
								nestedSignatureAndConditionOrCondition.Operator = &SignatureAndConditionOrConditionOperatorXml{}
								if _, ok := o.Misc["SignatureAndConditionOrConditionOperator"]; ok {
									nestedSignatureAndConditionOrCondition.Operator.Misc = o.Misc["SignatureAndConditionOrConditionOperator"]
								}
								if oSignatureAndConditionOrCondition.Operator.PatternMatch != nil {
									nestedSignatureAndConditionOrCondition.Operator.PatternMatch = &SignatureAndConditionOrConditionOperatorPatternMatchXml{}
									if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorPatternMatch"]; ok {
										nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Misc = o.Misc["SignatureAndConditionOrConditionOperatorPatternMatch"]
									}
									if oSignatureAndConditionOrCondition.Operator.PatternMatch.Context != nil {
										nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Context = oSignatureAndConditionOrCondition.Operator.PatternMatch.Context
									}
									if oSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern != nil {
										nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern = oSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern
									}
									if oSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier != nil {
										nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier = []SignatureAndConditionOrConditionOperatorPatternMatchQualifierXml{}
										for _, oSignatureAndConditionOrConditionOperatorPatternMatchQualifier := range oSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier {
											nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier := SignatureAndConditionOrConditionOperatorPatternMatchQualifierXml{}
											if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorPatternMatchQualifier"]; ok {
												nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Misc = o.Misc["SignatureAndConditionOrConditionOperatorPatternMatchQualifier"]
											}
											if oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value != nil {
												nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value = oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value
											}
											if oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name != "" {
												nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name = oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name
											}
											nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier, nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier)
										}
									}
								}
								if oSignatureAndConditionOrCondition.Operator.GreaterThan != nil {
									nestedSignatureAndConditionOrCondition.Operator.GreaterThan = &SignatureAndConditionOrConditionOperatorGreaterThanXml{}
									if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorGreaterThan"]; ok {
										nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Misc = o.Misc["SignatureAndConditionOrConditionOperatorGreaterThan"]
									}
									if oSignatureAndConditionOrCondition.Operator.GreaterThan.Context != nil {
										nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Context = oSignatureAndConditionOrCondition.Operator.GreaterThan.Context
									}
									if oSignatureAndConditionOrCondition.Operator.GreaterThan.Value != nil {
										nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Value = oSignatureAndConditionOrCondition.Operator.GreaterThan.Value
									}
									if oSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier != nil {
										nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier = []SignatureAndConditionOrConditionOperatorGreaterThanQualifierXml{}
										for _, oSignatureAndConditionOrConditionOperatorGreaterThanQualifier := range oSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier {
											nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier := SignatureAndConditionOrConditionOperatorGreaterThanQualifierXml{}
											if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorGreaterThanQualifier"]; ok {
												nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Misc = o.Misc["SignatureAndConditionOrConditionOperatorGreaterThanQualifier"]
											}
											if oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value != nil {
												nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value = oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value
											}
											if oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name != "" {
												nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name = oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name
											}
											nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier, nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier)
										}
									}
								}
								if oSignatureAndConditionOrCondition.Operator.LessThan != nil {
									nestedSignatureAndConditionOrCondition.Operator.LessThan = &SignatureAndConditionOrConditionOperatorLessThanXml{}
									if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorLessThan"]; ok {
										nestedSignatureAndConditionOrCondition.Operator.LessThan.Misc = o.Misc["SignatureAndConditionOrConditionOperatorLessThan"]
									}
									if oSignatureAndConditionOrCondition.Operator.LessThan.Context != nil {
										nestedSignatureAndConditionOrCondition.Operator.LessThan.Context = oSignatureAndConditionOrCondition.Operator.LessThan.Context
									}
									if oSignatureAndConditionOrCondition.Operator.LessThan.Value != nil {
										nestedSignatureAndConditionOrCondition.Operator.LessThan.Value = oSignatureAndConditionOrCondition.Operator.LessThan.Value
									}
									if oSignatureAndConditionOrCondition.Operator.LessThan.Qualifier != nil {
										nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier = []SignatureAndConditionOrConditionOperatorLessThanQualifierXml{}
										for _, oSignatureAndConditionOrConditionOperatorLessThanQualifier := range oSignatureAndConditionOrCondition.Operator.LessThan.Qualifier {
											nestedSignatureAndConditionOrConditionOperatorLessThanQualifier := SignatureAndConditionOrConditionOperatorLessThanQualifierXml{}
											if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorLessThanQualifier"]; ok {
												nestedSignatureAndConditionOrConditionOperatorLessThanQualifier.Misc = o.Misc["SignatureAndConditionOrConditionOperatorLessThanQualifier"]
											}
											if oSignatureAndConditionOrConditionOperatorLessThanQualifier.Value != nil {
												nestedSignatureAndConditionOrConditionOperatorLessThanQualifier.Value = oSignatureAndConditionOrConditionOperatorLessThanQualifier.Value
											}
											if oSignatureAndConditionOrConditionOperatorLessThanQualifier.Name != "" {
												nestedSignatureAndConditionOrConditionOperatorLessThanQualifier.Name = oSignatureAndConditionOrConditionOperatorLessThanQualifier.Name
											}
											nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier, nestedSignatureAndConditionOrConditionOperatorLessThanQualifier)
										}
									}
								}
								if oSignatureAndConditionOrCondition.Operator.EqualTo != nil {
									nestedSignatureAndConditionOrCondition.Operator.EqualTo = &SignatureAndConditionOrConditionOperatorEqualToXml{}
									if _, ok := o.Misc["SignatureAndConditionOrConditionOperatorEqualTo"]; ok {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo.Misc = o.Misc["SignatureAndConditionOrConditionOperatorEqualTo"]
									}
									if oSignatureAndConditionOrCondition.Operator.EqualTo.Position != nil {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo.Position = oSignatureAndConditionOrCondition.Operator.EqualTo.Position
									}
									if oSignatureAndConditionOrCondition.Operator.EqualTo.Mask != nil {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo.Mask = oSignatureAndConditionOrCondition.Operator.EqualTo.Mask
									}
									if oSignatureAndConditionOrCondition.Operator.EqualTo.Value != nil {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo.Value = oSignatureAndConditionOrCondition.Operator.EqualTo.Value
									}
									if oSignatureAndConditionOrCondition.Operator.EqualTo.Context != nil {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo.Context = oSignatureAndConditionOrCondition.Operator.EqualTo.Context
									}
								}
							}
							if oSignatureAndConditionOrCondition.Name != "" {
								nestedSignatureAndConditionOrCondition.Name = oSignatureAndConditionOrCondition.Name
							}
							nestedSignatureAndCondition.OrCondition = append(nestedSignatureAndCondition.OrCondition, nestedSignatureAndConditionOrCondition)
						}
					}
					if oSignatureAndCondition.Name != "" {
						nestedSignatureAndCondition.Name = oSignatureAndCondition.Name
					}
					nestedSignature.AndCondition = append(nestedSignature.AndCondition, nestedSignatureAndCondition)
				}
			}
			if oSignature.Name != "" {
				nestedSignature.Name = oSignature.Name
			}
			nestedSignatureCol = append(nestedSignatureCol, nestedSignature)
		}
		entry.Signature = nestedSignatureCol
	}

	entry.Subcategory = o.Subcategory
	entry.TcpHalfClosedTimeout = o.TcpHalfClosedTimeout
	entry.TcpTimeWaitTimeout = o.TcpTimeWaitTimeout
	entry.TcpTimeout = o.TcpTimeout
	entry.Technology = o.Technology
	entry.Timeout = o.Timeout
	entry.TunnelApplications = util.YesNo(o.TunnelApplications, nil)
	entry.TunnelOtherApplication = util.YesNo(o.TunnelOtherApplication, nil)
	entry.UdpTimeout = o.UdpTimeout
	entry.UsedByMalware = util.YesNo(o.UsedByMalware, nil)
	entry.VirusIdent = util.YesNo(o.VirusIdent, nil)

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
		entry.AbleToTransferFile = util.AsBool(o.AbleToTransferFile, nil)
		entry.AlgDisableCapability = o.AlgDisableCapability
		entry.Category = o.Category
		entry.ConsumeBigBandwidth = util.AsBool(o.ConsumeBigBandwidth, nil)
		entry.DataIdent = util.AsBool(o.DataIdent, nil)
		var nestedDefault *Default
		if o.Default != nil {
			nestedDefault = &Default{}
			if o.Default.Misc != nil {
				entry.Misc["Default"] = o.Default.Misc
			}
			if o.Default.IdentByIcmp6Type != nil {
				nestedDefault.IdentByIcmp6Type = &DefaultIdentByIcmp6Type{}
				if o.Default.IdentByIcmp6Type.Misc != nil {
					entry.Misc["DefaultIdentByIcmp6Type"] = o.Default.IdentByIcmp6Type.Misc
				}
				if o.Default.IdentByIcmp6Type.Code != nil {
					nestedDefault.IdentByIcmp6Type.Code = o.Default.IdentByIcmp6Type.Code
				}
				if o.Default.IdentByIcmp6Type.Type != nil {
					nestedDefault.IdentByIcmp6Type.Type = o.Default.IdentByIcmp6Type.Type
				}
			}
			if o.Default.IdentByIpProtocol != nil {
				nestedDefault.IdentByIpProtocol = o.Default.IdentByIpProtocol
			}
			if o.Default.Port != nil {
				nestedDefault.Port = util.MemToStr(o.Default.Port)
			}
			if o.Default.IdentByIcmpType != nil {
				nestedDefault.IdentByIcmpType = &DefaultIdentByIcmpType{}
				if o.Default.IdentByIcmpType.Misc != nil {
					entry.Misc["DefaultIdentByIcmpType"] = o.Default.IdentByIcmpType.Misc
				}
				if o.Default.IdentByIcmpType.Code != nil {
					nestedDefault.IdentByIcmpType.Code = o.Default.IdentByIcmpType.Code
				}
				if o.Default.IdentByIcmpType.Type != nil {
					nestedDefault.IdentByIcmpType.Type = o.Default.IdentByIcmpType.Type
				}
			}
		}
		entry.Default = nestedDefault

		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.EvasiveBehavior = util.AsBool(o.EvasiveBehavior, nil)
		entry.FileTypeIdent = util.AsBool(o.FileTypeIdent, nil)
		entry.HasKnownVulnerability = util.AsBool(o.HasKnownVulnerability, nil)
		entry.NoAppidCaching = util.AsBool(o.NoAppidCaching, nil)
		entry.ParentApp = o.ParentApp
		entry.PervasiveUse = util.AsBool(o.PervasiveUse, nil)
		entry.ProneToMisuse = util.AsBool(o.ProneToMisuse, nil)
		entry.Risk = o.Risk
		var nestedSignatureCol []Signature
		if o.Signature != nil {
			nestedSignatureCol = []Signature{}
			for _, oSignature := range o.Signature {
				nestedSignature := Signature{}
				if oSignature.Misc != nil {
					entry.Misc["Signature"] = oSignature.Misc
				}
				if oSignature.Name != "" {
					nestedSignature.Name = oSignature.Name
				}
				if oSignature.Comment != nil {
					nestedSignature.Comment = oSignature.Comment
				}
				if oSignature.Scope != nil {
					nestedSignature.Scope = oSignature.Scope
				}
				if oSignature.OrderFree != nil {
					nestedSignature.OrderFree = util.AsBool(oSignature.OrderFree, nil)
				}
				if oSignature.AndCondition != nil {
					nestedSignature.AndCondition = []SignatureAndCondition{}
					for _, oSignatureAndCondition := range oSignature.AndCondition {
						nestedSignatureAndCondition := SignatureAndCondition{}
						if oSignatureAndCondition.Misc != nil {
							entry.Misc["SignatureAndCondition"] = oSignatureAndCondition.Misc
						}
						if oSignatureAndCondition.OrCondition != nil {
							nestedSignatureAndCondition.OrCondition = []SignatureAndConditionOrCondition{}
							for _, oSignatureAndConditionOrCondition := range oSignatureAndCondition.OrCondition {
								nestedSignatureAndConditionOrCondition := SignatureAndConditionOrCondition{}
								if oSignatureAndConditionOrCondition.Misc != nil {
									entry.Misc["SignatureAndConditionOrCondition"] = oSignatureAndConditionOrCondition.Misc
								}
								if oSignatureAndConditionOrCondition.Operator != nil {
									nestedSignatureAndConditionOrCondition.Operator = &SignatureAndConditionOrConditionOperator{}
									if oSignatureAndConditionOrCondition.Operator.Misc != nil {
										entry.Misc["SignatureAndConditionOrConditionOperator"] = oSignatureAndConditionOrCondition.Operator.Misc
									}
									if oSignatureAndConditionOrCondition.Operator.PatternMatch != nil {
										nestedSignatureAndConditionOrCondition.Operator.PatternMatch = &SignatureAndConditionOrConditionOperatorPatternMatch{}
										if oSignatureAndConditionOrCondition.Operator.PatternMatch.Misc != nil {
											entry.Misc["SignatureAndConditionOrConditionOperatorPatternMatch"] = oSignatureAndConditionOrCondition.Operator.PatternMatch.Misc
										}
										if oSignatureAndConditionOrCondition.Operator.PatternMatch.Context != nil {
											nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Context = oSignatureAndConditionOrCondition.Operator.PatternMatch.Context
										}
										if oSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern != nil {
											nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern = oSignatureAndConditionOrCondition.Operator.PatternMatch.Pattern
										}
										if oSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier != nil {
											nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier = []SignatureAndConditionOrConditionOperatorPatternMatchQualifier{}
											for _, oSignatureAndConditionOrConditionOperatorPatternMatchQualifier := range oSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier {
												nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier := SignatureAndConditionOrConditionOperatorPatternMatchQualifier{}
												if oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Misc != nil {
													entry.Misc["SignatureAndConditionOrConditionOperatorPatternMatchQualifier"] = oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Misc
												}
												if oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value != nil {
													nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value = oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Value
												}
												if oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name != "" {
													nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name = oSignatureAndConditionOrConditionOperatorPatternMatchQualifier.Name
												}
												nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.PatternMatch.Qualifier, nestedSignatureAndConditionOrConditionOperatorPatternMatchQualifier)
											}
										}
									}
									if oSignatureAndConditionOrCondition.Operator.GreaterThan != nil {
										nestedSignatureAndConditionOrCondition.Operator.GreaterThan = &SignatureAndConditionOrConditionOperatorGreaterThan{}
										if oSignatureAndConditionOrCondition.Operator.GreaterThan.Misc != nil {
											entry.Misc["SignatureAndConditionOrConditionOperatorGreaterThan"] = oSignatureAndConditionOrCondition.Operator.GreaterThan.Misc
										}
										if oSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier != nil {
											nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier = []SignatureAndConditionOrConditionOperatorGreaterThanQualifier{}
											for _, oSignatureAndConditionOrConditionOperatorGreaterThanQualifier := range oSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier {
												nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier := SignatureAndConditionOrConditionOperatorGreaterThanQualifier{}
												if oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Misc != nil {
													entry.Misc["SignatureAndConditionOrConditionOperatorGreaterThanQualifier"] = oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Misc
												}
												if oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value != nil {
													nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value = oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Value
												}
												if oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name != "" {
													nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name = oSignatureAndConditionOrConditionOperatorGreaterThanQualifier.Name
												}
												nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Qualifier, nestedSignatureAndConditionOrConditionOperatorGreaterThanQualifier)
											}
										}
										if oSignatureAndConditionOrCondition.Operator.GreaterThan.Context != nil {
											nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Context = oSignatureAndConditionOrCondition.Operator.GreaterThan.Context
										}
										if oSignatureAndConditionOrCondition.Operator.GreaterThan.Value != nil {
											nestedSignatureAndConditionOrCondition.Operator.GreaterThan.Value = oSignatureAndConditionOrCondition.Operator.GreaterThan.Value
										}
									}
									if oSignatureAndConditionOrCondition.Operator.LessThan != nil {
										nestedSignatureAndConditionOrCondition.Operator.LessThan = &SignatureAndConditionOrConditionOperatorLessThan{}
										if oSignatureAndConditionOrCondition.Operator.LessThan.Misc != nil {
											entry.Misc["SignatureAndConditionOrConditionOperatorLessThan"] = oSignatureAndConditionOrCondition.Operator.LessThan.Misc
										}
										if oSignatureAndConditionOrCondition.Operator.LessThan.Context != nil {
											nestedSignatureAndConditionOrCondition.Operator.LessThan.Context = oSignatureAndConditionOrCondition.Operator.LessThan.Context
										}
										if oSignatureAndConditionOrCondition.Operator.LessThan.Value != nil {
											nestedSignatureAndConditionOrCondition.Operator.LessThan.Value = oSignatureAndConditionOrCondition.Operator.LessThan.Value
										}
										if oSignatureAndConditionOrCondition.Operator.LessThan.Qualifier != nil {
											nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier = []SignatureAndConditionOrConditionOperatorLessThanQualifier{}
											for _, oSignatureAndConditionOrConditionOperatorLessThanQualifier := range oSignatureAndConditionOrCondition.Operator.LessThan.Qualifier {
												nestedSignatureAndConditionOrConditionOperatorLessThanQualifier := SignatureAndConditionOrConditionOperatorLessThanQualifier{}
												if oSignatureAndConditionOrConditionOperatorLessThanQualifier.Misc != nil {
													entry.Misc["SignatureAndConditionOrConditionOperatorLessThanQualifier"] = oSignatureAndConditionOrConditionOperatorLessThanQualifier.Misc
												}
												if oSignatureAndConditionOrConditionOperatorLessThanQualifier.Name != "" {
													nestedSignatureAndConditionOrConditionOperatorLessThanQualifier.Name = oSignatureAndConditionOrConditionOperatorLessThanQualifier.Name
												}
												if oSignatureAndConditionOrConditionOperatorLessThanQualifier.Value != nil {
													nestedSignatureAndConditionOrConditionOperatorLessThanQualifier.Value = oSignatureAndConditionOrConditionOperatorLessThanQualifier.Value
												}
												nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier = append(nestedSignatureAndConditionOrCondition.Operator.LessThan.Qualifier, nestedSignatureAndConditionOrConditionOperatorLessThanQualifier)
											}
										}
									}
									if oSignatureAndConditionOrCondition.Operator.EqualTo != nil {
										nestedSignatureAndConditionOrCondition.Operator.EqualTo = &SignatureAndConditionOrConditionOperatorEqualTo{}
										if oSignatureAndConditionOrCondition.Operator.EqualTo.Misc != nil {
											entry.Misc["SignatureAndConditionOrConditionOperatorEqualTo"] = oSignatureAndConditionOrCondition.Operator.EqualTo.Misc
										}
										if oSignatureAndConditionOrCondition.Operator.EqualTo.Context != nil {
											nestedSignatureAndConditionOrCondition.Operator.EqualTo.Context = oSignatureAndConditionOrCondition.Operator.EqualTo.Context
										}
										if oSignatureAndConditionOrCondition.Operator.EqualTo.Position != nil {
											nestedSignatureAndConditionOrCondition.Operator.EqualTo.Position = oSignatureAndConditionOrCondition.Operator.EqualTo.Position
										}
										if oSignatureAndConditionOrCondition.Operator.EqualTo.Mask != nil {
											nestedSignatureAndConditionOrCondition.Operator.EqualTo.Mask = oSignatureAndConditionOrCondition.Operator.EqualTo.Mask
										}
										if oSignatureAndConditionOrCondition.Operator.EqualTo.Value != nil {
											nestedSignatureAndConditionOrCondition.Operator.EqualTo.Value = oSignatureAndConditionOrCondition.Operator.EqualTo.Value
										}
									}
								}
								if oSignatureAndConditionOrCondition.Name != "" {
									nestedSignatureAndConditionOrCondition.Name = oSignatureAndConditionOrCondition.Name
								}
								nestedSignatureAndCondition.OrCondition = append(nestedSignatureAndCondition.OrCondition, nestedSignatureAndConditionOrCondition)
							}
						}
						if oSignatureAndCondition.Name != "" {
							nestedSignatureAndCondition.Name = oSignatureAndCondition.Name
						}
						nestedSignature.AndCondition = append(nestedSignature.AndCondition, nestedSignatureAndCondition)
					}
				}
				nestedSignatureCol = append(nestedSignatureCol, nestedSignature)
			}
			entry.Signature = nestedSignatureCol
		}

		entry.Subcategory = o.Subcategory
		entry.TcpHalfClosedTimeout = o.TcpHalfClosedTimeout
		entry.TcpTimeWaitTimeout = o.TcpTimeWaitTimeout
		entry.TcpTimeout = o.TcpTimeout
		entry.Technology = o.Technology
		entry.Timeout = o.Timeout
		entry.TunnelApplications = util.AsBool(o.TunnelApplications, nil)
		entry.TunnelOtherApplication = util.AsBool(o.TunnelOtherApplication, nil)
		entry.UdpTimeout = o.UdpTimeout
		entry.UsedByMalware = util.AsBool(o.UsedByMalware, nil)
		entry.VirusIdent = util.AsBool(o.VirusIdent, nil)

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
	if !util.BoolsMatch(a.AbleToTransferFile, b.AbleToTransferFile) {
		return false
	}
	if !util.StringsMatch(a.AlgDisableCapability, b.AlgDisableCapability) {
		return false
	}
	if !util.StringsMatch(a.Category, b.Category) {
		return false
	}
	if !util.BoolsMatch(a.ConsumeBigBandwidth, b.ConsumeBigBandwidth) {
		return false
	}
	if !util.BoolsMatch(a.DataIdent, b.DataIdent) {
		return false
	}
	if !matchDefault(a.Default, b.Default) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(a.EvasiveBehavior, b.EvasiveBehavior) {
		return false
	}
	if !util.BoolsMatch(a.FileTypeIdent, b.FileTypeIdent) {
		return false
	}
	if !util.BoolsMatch(a.HasKnownVulnerability, b.HasKnownVulnerability) {
		return false
	}
	if !util.BoolsMatch(a.NoAppidCaching, b.NoAppidCaching) {
		return false
	}
	if !util.StringsMatch(a.ParentApp, b.ParentApp) {
		return false
	}
	if !util.BoolsMatch(a.PervasiveUse, b.PervasiveUse) {
		return false
	}
	if !util.BoolsMatch(a.ProneToMisuse, b.ProneToMisuse) {
		return false
	}
	if !util.Ints64Match(a.Risk, b.Risk) {
		return false
	}
	if !matchSignature(a.Signature, b.Signature) {
		return false
	}
	if !util.StringsMatch(a.Subcategory, b.Subcategory) {
		return false
	}
	if !util.Ints64Match(a.TcpHalfClosedTimeout, b.TcpHalfClosedTimeout) {
		return false
	}
	if !util.Ints64Match(a.TcpTimeWaitTimeout, b.TcpTimeWaitTimeout) {
		return false
	}
	if !util.Ints64Match(a.TcpTimeout, b.TcpTimeout) {
		return false
	}
	if !util.StringsMatch(a.Technology, b.Technology) {
		return false
	}
	if !util.Ints64Match(a.Timeout, b.Timeout) {
		return false
	}
	if !util.BoolsMatch(a.TunnelApplications, b.TunnelApplications) {
		return false
	}
	if !util.BoolsMatch(a.TunnelOtherApplication, b.TunnelOtherApplication) {
		return false
	}
	if !util.Ints64Match(a.UdpTimeout, b.UdpTimeout) {
		return false
	}
	if !util.BoolsMatch(a.UsedByMalware, b.UsedByMalware) {
		return false
	}
	if !util.BoolsMatch(a.VirusIdent, b.VirusIdent) {
		return false
	}

	return true
}

func matchDefaultIdentByIcmpType(a *DefaultIdentByIcmpType, b *DefaultIdentByIcmpType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Code, b.Code) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	return true
}
func matchDefaultIdentByIcmp6Type(a *DefaultIdentByIcmp6Type, b *DefaultIdentByIcmp6Type) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Code, b.Code) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	return true
}
func matchDefault(a *Default, b *Default) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchDefaultIdentByIcmpType(a.IdentByIcmpType, b.IdentByIcmpType) {
		return false
	}
	if !matchDefaultIdentByIcmp6Type(a.IdentByIcmp6Type, b.IdentByIcmp6Type) {
		return false
	}
	if !util.StringsMatch(a.IdentByIpProtocol, b.IdentByIpProtocol) {
		return false
	}
	if !util.OrderedListsMatch(a.Port, b.Port) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorEqualTo(a *SignatureAndConditionOrConditionOperatorEqualTo, b *SignatureAndConditionOrConditionOperatorEqualTo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Context, b.Context) {
		return false
	}
	if !util.StringsMatch(a.Position, b.Position) {
		return false
	}
	if !util.StringsMatch(a.Mask, b.Mask) {
		return false
	}
	if !util.StringsMatch(a.Value, b.Value) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorPatternMatchQualifier(a []SignatureAndConditionOrConditionOperatorPatternMatchQualifier, b []SignatureAndConditionOrConditionOperatorPatternMatchQualifier) bool {
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
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
		}
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorPatternMatch(a *SignatureAndConditionOrConditionOperatorPatternMatch, b *SignatureAndConditionOrConditionOperatorPatternMatch) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Context, b.Context) {
		return false
	}
	if !util.StringsMatch(a.Pattern, b.Pattern) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorPatternMatchQualifier(a.Qualifier, b.Qualifier) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorGreaterThanQualifier(a []SignatureAndConditionOrConditionOperatorGreaterThanQualifier, b []SignatureAndConditionOrConditionOperatorGreaterThanQualifier) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorGreaterThan(a *SignatureAndConditionOrConditionOperatorGreaterThan, b *SignatureAndConditionOrConditionOperatorGreaterThan) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Context, b.Context) {
		return false
	}
	if !util.Ints64Match(a.Value, b.Value) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorGreaterThanQualifier(a.Qualifier, b.Qualifier) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorLessThanQualifier(a []SignatureAndConditionOrConditionOperatorLessThanQualifier, b []SignatureAndConditionOrConditionOperatorLessThanQualifier) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Value, b.Value) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchSignatureAndConditionOrConditionOperatorLessThan(a *SignatureAndConditionOrConditionOperatorLessThan, b *SignatureAndConditionOrConditionOperatorLessThan) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Context, b.Context) {
		return false
	}
	if !util.Ints64Match(a.Value, b.Value) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorLessThanQualifier(a.Qualifier, b.Qualifier) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrConditionOperator(a *SignatureAndConditionOrConditionOperator, b *SignatureAndConditionOrConditionOperator) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchSignatureAndConditionOrConditionOperatorEqualTo(a.EqualTo, b.EqualTo) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorPatternMatch(a.PatternMatch, b.PatternMatch) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorGreaterThan(a.GreaterThan, b.GreaterThan) {
		return false
	}
	if !matchSignatureAndConditionOrConditionOperatorLessThan(a.LessThan, b.LessThan) {
		return false
	}
	return true
}
func matchSignatureAndConditionOrCondition(a []SignatureAndConditionOrCondition, b []SignatureAndConditionOrCondition) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchSignatureAndConditionOrConditionOperator(a.Operator, b.Operator) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchSignatureAndCondition(a []SignatureAndCondition, b []SignatureAndCondition) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchSignatureAndConditionOrCondition(a.OrCondition, b.OrCondition) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchSignature(a []Signature, b []Signature) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Comment, b.Comment) {
				return false
			}
			if !util.StringsMatch(a.Scope, b.Scope) {
				return false
			}
			if !util.BoolsMatch(a.OrderFree, b.OrderFree) {
				return false
			}
			if !matchSignatureAndCondition(a.AndCondition, b.AndCondition) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
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
