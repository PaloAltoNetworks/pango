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
	suffix = []string{"application", "$name"}
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
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
}
type Default struct {
	IdentByIcmpType   *DefaultIdentByIcmpType
	IdentByIcmp6Type  *DefaultIdentByIcmp6Type
	IdentByIpProtocol *string
	Port              []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type DefaultIdentByIcmpType struct {
	Code           *string
	Type           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultIdentByIcmp6Type struct {
	Code           *string
	Type           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Signature struct {
	Name           string
	Comment        *string
	Scope          *string
	OrderFree      *bool
	AndCondition   []SignatureAndCondition
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndCondition struct {
	Name           string
	OrCondition    []SignatureAndConditionOrCondition
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrCondition struct {
	Name           string
	Operator       *SignatureAndConditionOrConditionOperator
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperator struct {
	PatternMatch   *SignatureAndConditionOrConditionOperatorPatternMatch
	GreaterThan    *SignatureAndConditionOrConditionOperatorGreaterThan
	LessThan       *SignatureAndConditionOrConditionOperatorLessThan
	EqualTo        *SignatureAndConditionOrConditionOperatorEqualTo
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorPatternMatch struct {
	Context        *string
	Pattern        *string
	Qualifier      []SignatureAndConditionOrConditionOperatorPatternMatchQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorPatternMatchQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorGreaterThan struct {
	Context        *string
	Value          *int64
	Qualifier      []SignatureAndConditionOrConditionOperatorGreaterThanQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorGreaterThanQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorLessThan struct {
	Context        *string
	Value          *int64
	Qualifier      []SignatureAndConditionOrConditionOperatorLessThanQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorLessThanQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureAndConditionOrConditionOperatorEqualTo struct {
	Context        *string
	Position       *string
	Mask           *string
	Value          *string
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
	XMLName                xml.Name               `xml:"entry"`
	Name                   string                 `xml:"name,attr"`
	AbleToTransferFile     *string                `xml:"able-to-transfer-file,omitempty"`
	AlgDisableCapability   *string                `xml:"alg-disable-capability,omitempty"`
	Category               *string                `xml:"category,omitempty"`
	ConsumeBigBandwidth    *string                `xml:"consume-big-bandwidth,omitempty"`
	DataIdent              *string                `xml:"data-ident,omitempty"`
	Default                *defaultXml            `xml:"default,omitempty"`
	Description            *string                `xml:"description,omitempty"`
	DisableOverride        *string                `xml:"disable-override,omitempty"`
	EvasiveBehavior        *string                `xml:"evasive-behavior,omitempty"`
	FileTypeIdent          *string                `xml:"file-type-ident,omitempty"`
	HasKnownVulnerability  *string                `xml:"has-known-vulnerability,omitempty"`
	NoAppidCaching         *string                `xml:"no-appid-caching,omitempty"`
	ParentApp              *string                `xml:"parent-app,omitempty"`
	PervasiveUse           *string                `xml:"pervasive-use,omitempty"`
	ProneToMisuse          *string                `xml:"prone-to-misuse,omitempty"`
	Risk                   *int64                 `xml:"risk,omitempty"`
	Signature              *signatureContainerXml `xml:"signature,omitempty"`
	Subcategory            *string                `xml:"subcategory,omitempty"`
	TcpHalfClosedTimeout   *int64                 `xml:"tcp-half-closed-timeout,omitempty"`
	TcpTimeWaitTimeout     *int64                 `xml:"tcp-time-wait-timeout,omitempty"`
	TcpTimeout             *int64                 `xml:"tcp-timeout,omitempty"`
	Technology             *string                `xml:"technology,omitempty"`
	Timeout                *int64                 `xml:"timeout,omitempty"`
	TunnelApplications     *string                `xml:"tunnel-applications,omitempty"`
	TunnelOtherApplication *string                `xml:"tunnel-other-application,omitempty"`
	UdpTimeout             *int64                 `xml:"udp-timeout,omitempty"`
	UsedByMalware          *string                `xml:"used-by-malware,omitempty"`
	VirusIdent             *string                `xml:"virus-ident,omitempty"`
	Misc                   []generic.Xml          `xml:",any"`
	MiscAttributes         []xml.Attr             `xml:",any,attr"`
}
type defaultXml struct {
	IdentByIcmpType   *defaultIdentByIcmpTypeXml  `xml:"ident-by-icmp-type,omitempty"`
	IdentByIcmp6Type  *defaultIdentByIcmp6TypeXml `xml:"ident-by-icmp6-type,omitempty"`
	IdentByIpProtocol *string                     `xml:"ident-by-ip-protocol,omitempty"`
	Port              *util.MemberType            `xml:"port,omitempty"`
	Misc              []generic.Xml               `xml:",any"`
	MiscAttributes    []xml.Attr                  `xml:",any,attr"`
}
type defaultIdentByIcmpTypeXml struct {
	Code           *string       `xml:"code,omitempty"`
	Type           *string       `xml:"type,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultIdentByIcmp6TypeXml struct {
	Code           *string       `xml:"code,omitempty"`
	Type           *string       `xml:"type,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureContainerXml struct {
	Entries []signatureXml `xml:"entry"`
}
type signatureXml struct {
	XMLName        xml.Name                           `xml:"entry"`
	Name           string                             `xml:"name,attr"`
	Comment        *string                            `xml:"comment,omitempty"`
	Scope          *string                            `xml:"scope,omitempty"`
	OrderFree      *string                            `xml:"order-free,omitempty"`
	AndCondition   *signatureAndConditionContainerXml `xml:"and-condition,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type signatureAndConditionContainerXml struct {
	Entries []signatureAndConditionXml `xml:"entry"`
}
type signatureAndConditionXml struct {
	XMLName        xml.Name                                      `xml:"entry"`
	Name           string                                        `xml:"name,attr"`
	OrCondition    *signatureAndConditionOrConditionContainerXml `xml:"or-condition,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type signatureAndConditionOrConditionContainerXml struct {
	Entries []signatureAndConditionOrConditionXml `xml:"entry"`
}
type signatureAndConditionOrConditionXml struct {
	XMLName        xml.Name                                     `xml:"entry"`
	Name           string                                       `xml:"name,attr"`
	Operator       *signatureAndConditionOrConditionOperatorXml `xml:"operator,omitempty"`
	Misc           []generic.Xml                                `xml:",any"`
	MiscAttributes []xml.Attr                                   `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorXml struct {
	PatternMatch   *signatureAndConditionOrConditionOperatorPatternMatchXml `xml:"pattern-match,omitempty"`
	GreaterThan    *signatureAndConditionOrConditionOperatorGreaterThanXml  `xml:"greater-than,omitempty"`
	LessThan       *signatureAndConditionOrConditionOperatorLessThanXml     `xml:"less-than,omitempty"`
	EqualTo        *signatureAndConditionOrConditionOperatorEqualToXml      `xml:"equal-to,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorPatternMatchXml struct {
	Context        *string                                                                    `xml:"context,omitempty"`
	Pattern        *string                                                                    `xml:"pattern,omitempty"`
	Qualifier      *signatureAndConditionOrConditionOperatorPatternMatchQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                              `xml:",any"`
	MiscAttributes []xml.Attr                                                                 `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorPatternMatchQualifierContainerXml struct {
	Entries []signatureAndConditionOrConditionOperatorPatternMatchQualifierXml `xml:"entry"`
}
type signatureAndConditionOrConditionOperatorPatternMatchQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorGreaterThanXml struct {
	Context        *string                                                                   `xml:"context,omitempty"`
	Value          *int64                                                                    `xml:"value,omitempty"`
	Qualifier      *signatureAndConditionOrConditionOperatorGreaterThanQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                             `xml:",any"`
	MiscAttributes []xml.Attr                                                                `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorGreaterThanQualifierContainerXml struct {
	Entries []signatureAndConditionOrConditionOperatorGreaterThanQualifierXml `xml:"entry"`
}
type signatureAndConditionOrConditionOperatorGreaterThanQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorLessThanXml struct {
	Context        *string                                                                `xml:"context,omitempty"`
	Value          *int64                                                                 `xml:"value,omitempty"`
	Qualifier      *signatureAndConditionOrConditionOperatorLessThanQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                          `xml:",any"`
	MiscAttributes []xml.Attr                                                             `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorLessThanQualifierContainerXml struct {
	Entries []signatureAndConditionOrConditionOperatorLessThanQualifierXml `xml:"entry"`
}
type signatureAndConditionOrConditionOperatorLessThanQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureAndConditionOrConditionOperatorEqualToXml struct {
	Context        *string       `xml:"context,omitempty"`
	Position       *string       `xml:"position,omitempty"`
	Mask           *string       `xml:"mask,omitempty"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AbleToTransferFile = util.YesNo(s.AbleToTransferFile, nil)
	o.AlgDisableCapability = s.AlgDisableCapability
	o.Category = s.Category
	o.ConsumeBigBandwidth = util.YesNo(s.ConsumeBigBandwidth, nil)
	o.DataIdent = util.YesNo(s.DataIdent, nil)
	if s.Default != nil {
		var obj defaultXml
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	o.EvasiveBehavior = util.YesNo(s.EvasiveBehavior, nil)
	o.FileTypeIdent = util.YesNo(s.FileTypeIdent, nil)
	o.HasKnownVulnerability = util.YesNo(s.HasKnownVulnerability, nil)
	o.NoAppidCaching = util.YesNo(s.NoAppidCaching, nil)
	o.ParentApp = s.ParentApp
	o.PervasiveUse = util.YesNo(s.PervasiveUse, nil)
	o.ProneToMisuse = util.YesNo(s.ProneToMisuse, nil)
	o.Risk = s.Risk
	if s.Signature != nil {
		var objs []signatureXml
		for _, elt := range s.Signature {
			var obj signatureXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Signature = &signatureContainerXml{Entries: objs}
	}
	o.Subcategory = s.Subcategory
	o.TcpHalfClosedTimeout = s.TcpHalfClosedTimeout
	o.TcpTimeWaitTimeout = s.TcpTimeWaitTimeout
	o.TcpTimeout = s.TcpTimeout
	o.Technology = s.Technology
	o.Timeout = s.Timeout
	o.TunnelApplications = util.YesNo(s.TunnelApplications, nil)
	o.TunnelOtherApplication = util.YesNo(s.TunnelOtherApplication, nil)
	o.UdpTimeout = s.UdpTimeout
	o.UsedByMalware = util.YesNo(s.UsedByMalware, nil)
	o.VirusIdent = util.YesNo(s.VirusIdent, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var defaultVal *Default
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var signatureVal []Signature
	if o.Signature != nil {
		for _, elt := range o.Signature.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			signatureVal = append(signatureVal, *obj)
		}
	}

	result := &Entry{
		Name:                   o.Name,
		AbleToTransferFile:     util.AsBool(o.AbleToTransferFile, nil),
		AlgDisableCapability:   o.AlgDisableCapability,
		Category:               o.Category,
		ConsumeBigBandwidth:    util.AsBool(o.ConsumeBigBandwidth, nil),
		DataIdent:              util.AsBool(o.DataIdent, nil),
		Default:                defaultVal,
		Description:            o.Description,
		DisableOverride:        o.DisableOverride,
		EvasiveBehavior:        util.AsBool(o.EvasiveBehavior, nil),
		FileTypeIdent:          util.AsBool(o.FileTypeIdent, nil),
		HasKnownVulnerability:  util.AsBool(o.HasKnownVulnerability, nil),
		NoAppidCaching:         util.AsBool(o.NoAppidCaching, nil),
		ParentApp:              o.ParentApp,
		PervasiveUse:           util.AsBool(o.PervasiveUse, nil),
		ProneToMisuse:          util.AsBool(o.ProneToMisuse, nil),
		Risk:                   o.Risk,
		Signature:              signatureVal,
		Subcategory:            o.Subcategory,
		TcpHalfClosedTimeout:   o.TcpHalfClosedTimeout,
		TcpTimeWaitTimeout:     o.TcpTimeWaitTimeout,
		TcpTimeout:             o.TcpTimeout,
		Technology:             o.Technology,
		Timeout:                o.Timeout,
		TunnelApplications:     util.AsBool(o.TunnelApplications, nil),
		TunnelOtherApplication: util.AsBool(o.TunnelOtherApplication, nil),
		UdpTimeout:             o.UdpTimeout,
		UsedByMalware:          util.AsBool(o.UsedByMalware, nil),
		VirusIdent:             util.AsBool(o.VirusIdent, nil),
		Misc:                   o.Misc,
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultXml) MarshalFromObject(s Default) {
	if s.IdentByIcmpType != nil {
		var obj defaultIdentByIcmpTypeXml
		obj.MarshalFromObject(*s.IdentByIcmpType)
		o.IdentByIcmpType = &obj
	}
	if s.IdentByIcmp6Type != nil {
		var obj defaultIdentByIcmp6TypeXml
		obj.MarshalFromObject(*s.IdentByIcmp6Type)
		o.IdentByIcmp6Type = &obj
	}
	o.IdentByIpProtocol = s.IdentByIpProtocol
	if s.Port != nil {
		o.Port = util.StrToMem(s.Port)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultXml) UnmarshalToObject() (*Default, error) {
	var identByIcmpTypeVal *DefaultIdentByIcmpType
	if o.IdentByIcmpType != nil {
		obj, err := o.IdentByIcmpType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		identByIcmpTypeVal = obj
	}
	var identByIcmp6TypeVal *DefaultIdentByIcmp6Type
	if o.IdentByIcmp6Type != nil {
		obj, err := o.IdentByIcmp6Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		identByIcmp6TypeVal = obj
	}
	var portVal []string
	if o.Port != nil {
		portVal = util.MemToStr(o.Port)
	}

	result := &Default{
		IdentByIcmpType:   identByIcmpTypeVal,
		IdentByIcmp6Type:  identByIcmp6TypeVal,
		IdentByIpProtocol: o.IdentByIpProtocol,
		Port:              portVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultIdentByIcmpTypeXml) MarshalFromObject(s DefaultIdentByIcmpType) {
	o.Code = s.Code
	o.Type = s.Type
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultIdentByIcmpTypeXml) UnmarshalToObject() (*DefaultIdentByIcmpType, error) {

	result := &DefaultIdentByIcmpType{
		Code:           o.Code,
		Type:           o.Type,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultIdentByIcmp6TypeXml) MarshalFromObject(s DefaultIdentByIcmp6Type) {
	o.Code = s.Code
	o.Type = s.Type
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultIdentByIcmp6TypeXml) UnmarshalToObject() (*DefaultIdentByIcmp6Type, error) {

	result := &DefaultIdentByIcmp6Type{
		Code:           o.Code,
		Type:           o.Type,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureXml) MarshalFromObject(s Signature) {
	o.Name = s.Name
	o.Comment = s.Comment
	o.Scope = s.Scope
	o.OrderFree = util.YesNo(s.OrderFree, nil)
	if s.AndCondition != nil {
		var objs []signatureAndConditionXml
		for _, elt := range s.AndCondition {
			var obj signatureAndConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AndCondition = &signatureAndConditionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureXml) UnmarshalToObject() (*Signature, error) {
	var andConditionVal []SignatureAndCondition
	if o.AndCondition != nil {
		for _, elt := range o.AndCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			andConditionVal = append(andConditionVal, *obj)
		}
	}

	result := &Signature{
		Name:           o.Name,
		Comment:        o.Comment,
		Scope:          o.Scope,
		OrderFree:      util.AsBool(o.OrderFree, nil),
		AndCondition:   andConditionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionXml) MarshalFromObject(s SignatureAndCondition) {
	o.Name = s.Name
	if s.OrCondition != nil {
		var objs []signatureAndConditionOrConditionXml
		for _, elt := range s.OrCondition {
			var obj signatureAndConditionOrConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.OrCondition = &signatureAndConditionOrConditionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionXml) UnmarshalToObject() (*SignatureAndCondition, error) {
	var orConditionVal []SignatureAndConditionOrCondition
	if o.OrCondition != nil {
		for _, elt := range o.OrCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			orConditionVal = append(orConditionVal, *obj)
		}
	}

	result := &SignatureAndCondition{
		Name:           o.Name,
		OrCondition:    orConditionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionXml) MarshalFromObject(s SignatureAndConditionOrCondition) {
	o.Name = s.Name
	if s.Operator != nil {
		var obj signatureAndConditionOrConditionOperatorXml
		obj.MarshalFromObject(*s.Operator)
		o.Operator = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionXml) UnmarshalToObject() (*SignatureAndConditionOrCondition, error) {
	var operatorVal *SignatureAndConditionOrConditionOperator
	if o.Operator != nil {
		obj, err := o.Operator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operatorVal = obj
	}

	result := &SignatureAndConditionOrCondition{
		Name:           o.Name,
		Operator:       operatorVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorXml) MarshalFromObject(s SignatureAndConditionOrConditionOperator) {
	if s.PatternMatch != nil {
		var obj signatureAndConditionOrConditionOperatorPatternMatchXml
		obj.MarshalFromObject(*s.PatternMatch)
		o.PatternMatch = &obj
	}
	if s.GreaterThan != nil {
		var obj signatureAndConditionOrConditionOperatorGreaterThanXml
		obj.MarshalFromObject(*s.GreaterThan)
		o.GreaterThan = &obj
	}
	if s.LessThan != nil {
		var obj signatureAndConditionOrConditionOperatorLessThanXml
		obj.MarshalFromObject(*s.LessThan)
		o.LessThan = &obj
	}
	if s.EqualTo != nil {
		var obj signatureAndConditionOrConditionOperatorEqualToXml
		obj.MarshalFromObject(*s.EqualTo)
		o.EqualTo = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperator, error) {
	var patternMatchVal *SignatureAndConditionOrConditionOperatorPatternMatch
	if o.PatternMatch != nil {
		obj, err := o.PatternMatch.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		patternMatchVal = obj
	}
	var greaterThanVal *SignatureAndConditionOrConditionOperatorGreaterThan
	if o.GreaterThan != nil {
		obj, err := o.GreaterThan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		greaterThanVal = obj
	}
	var lessThanVal *SignatureAndConditionOrConditionOperatorLessThan
	if o.LessThan != nil {
		obj, err := o.LessThan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lessThanVal = obj
	}
	var equalToVal *SignatureAndConditionOrConditionOperatorEqualTo
	if o.EqualTo != nil {
		obj, err := o.EqualTo.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		equalToVal = obj
	}

	result := &SignatureAndConditionOrConditionOperator{
		PatternMatch:   patternMatchVal,
		GreaterThan:    greaterThanVal,
		LessThan:       lessThanVal,
		EqualTo:        equalToVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorPatternMatchXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorPatternMatch) {
	o.Context = s.Context
	o.Pattern = s.Pattern
	if s.Qualifier != nil {
		var objs []signatureAndConditionOrConditionOperatorPatternMatchQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureAndConditionOrConditionOperatorPatternMatchQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureAndConditionOrConditionOperatorPatternMatchQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorPatternMatchXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorPatternMatch, error) {
	var qualifierVal []SignatureAndConditionOrConditionOperatorPatternMatchQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureAndConditionOrConditionOperatorPatternMatch{
		Context:        o.Context,
		Pattern:        o.Pattern,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorPatternMatchQualifierXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorPatternMatchQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorPatternMatchQualifierXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorPatternMatchQualifier, error) {

	result := &SignatureAndConditionOrConditionOperatorPatternMatchQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorGreaterThanXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorGreaterThan) {
	o.Context = s.Context
	o.Value = s.Value
	if s.Qualifier != nil {
		var objs []signatureAndConditionOrConditionOperatorGreaterThanQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureAndConditionOrConditionOperatorGreaterThanQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureAndConditionOrConditionOperatorGreaterThanQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorGreaterThanXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorGreaterThan, error) {
	var qualifierVal []SignatureAndConditionOrConditionOperatorGreaterThanQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureAndConditionOrConditionOperatorGreaterThan{
		Context:        o.Context,
		Value:          o.Value,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorGreaterThanQualifierXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorGreaterThanQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorGreaterThanQualifierXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorGreaterThanQualifier, error) {

	result := &SignatureAndConditionOrConditionOperatorGreaterThanQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorLessThanXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorLessThan) {
	o.Context = s.Context
	o.Value = s.Value
	if s.Qualifier != nil {
		var objs []signatureAndConditionOrConditionOperatorLessThanQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureAndConditionOrConditionOperatorLessThanQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureAndConditionOrConditionOperatorLessThanQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorLessThanXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorLessThan, error) {
	var qualifierVal []SignatureAndConditionOrConditionOperatorLessThanQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureAndConditionOrConditionOperatorLessThan{
		Context:        o.Context,
		Value:          o.Value,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorLessThanQualifierXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorLessThanQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorLessThanQualifierXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorLessThanQualifier, error) {

	result := &SignatureAndConditionOrConditionOperatorLessThanQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureAndConditionOrConditionOperatorEqualToXml) MarshalFromObject(s SignatureAndConditionOrConditionOperatorEqualTo) {
	o.Context = s.Context
	o.Position = s.Position
	o.Mask = s.Mask
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureAndConditionOrConditionOperatorEqualToXml) UnmarshalToObject() (*SignatureAndConditionOrConditionOperatorEqualTo, error) {

	result := &SignatureAndConditionOrConditionOperatorEqualTo{
		Context:        o.Context,
		Position:       o.Position,
		Mask:           o.Mask,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
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
	if !util.BoolsMatch(o.AbleToTransferFile, other.AbleToTransferFile) {
		return false
	}
	if !util.StringsMatch(o.AlgDisableCapability, other.AlgDisableCapability) {
		return false
	}
	if !util.StringsMatch(o.Category, other.Category) {
		return false
	}
	if !util.BoolsMatch(o.ConsumeBigBandwidth, other.ConsumeBigBandwidth) {
		return false
	}
	if !util.BoolsMatch(o.DataIdent, other.DataIdent) {
		return false
	}
	if !o.Default.matches(other.Default) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(o.EvasiveBehavior, other.EvasiveBehavior) {
		return false
	}
	if !util.BoolsMatch(o.FileTypeIdent, other.FileTypeIdent) {
		return false
	}
	if !util.BoolsMatch(o.HasKnownVulnerability, other.HasKnownVulnerability) {
		return false
	}
	if !util.BoolsMatch(o.NoAppidCaching, other.NoAppidCaching) {
		return false
	}
	if !util.StringsMatch(o.ParentApp, other.ParentApp) {
		return false
	}
	if !util.BoolsMatch(o.PervasiveUse, other.PervasiveUse) {
		return false
	}
	if !util.BoolsMatch(o.ProneToMisuse, other.ProneToMisuse) {
		return false
	}
	if !util.Ints64Match(o.Risk, other.Risk) {
		return false
	}
	if len(o.Signature) != len(other.Signature) {
		return false
	}
	for idx := range o.Signature {
		if !o.Signature[idx].matches(&other.Signature[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.Subcategory, other.Subcategory) {
		return false
	}
	if !util.Ints64Match(o.TcpHalfClosedTimeout, other.TcpHalfClosedTimeout) {
		return false
	}
	if !util.Ints64Match(o.TcpTimeWaitTimeout, other.TcpTimeWaitTimeout) {
		return false
	}
	if !util.Ints64Match(o.TcpTimeout, other.TcpTimeout) {
		return false
	}
	if !util.StringsMatch(o.Technology, other.Technology) {
		return false
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
		return false
	}
	if !util.BoolsMatch(o.TunnelApplications, other.TunnelApplications) {
		return false
	}
	if !util.BoolsMatch(o.TunnelOtherApplication, other.TunnelOtherApplication) {
		return false
	}
	if !util.Ints64Match(o.UdpTimeout, other.UdpTimeout) {
		return false
	}
	if !util.BoolsMatch(o.UsedByMalware, other.UsedByMalware) {
		return false
	}
	if !util.BoolsMatch(o.VirusIdent, other.VirusIdent) {
		return false
	}

	return true
}

func (o *Default) matches(other *Default) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.IdentByIcmpType.matches(other.IdentByIcmpType) {
		return false
	}
	if !o.IdentByIcmp6Type.matches(other.IdentByIcmp6Type) {
		return false
	}
	if !util.StringsMatch(o.IdentByIpProtocol, other.IdentByIpProtocol) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Port, other.Port) {
		return false
	}

	return true
}

func (o *DefaultIdentByIcmpType) matches(other *DefaultIdentByIcmpType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Code, other.Code) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *DefaultIdentByIcmp6Type) matches(other *DefaultIdentByIcmp6Type) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Code, other.Code) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *Signature) matches(other *Signature) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !util.StringsMatch(o.Scope, other.Scope) {
		return false
	}
	if !util.BoolsMatch(o.OrderFree, other.OrderFree) {
		return false
	}
	if len(o.AndCondition) != len(other.AndCondition) {
		return false
	}
	for idx := range o.AndCondition {
		if !o.AndCondition[idx].matches(&other.AndCondition[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureAndCondition) matches(other *SignatureAndCondition) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.OrCondition) != len(other.OrCondition) {
		return false
	}
	for idx := range o.OrCondition {
		if !o.OrCondition[idx].matches(&other.OrCondition[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureAndConditionOrCondition) matches(other *SignatureAndConditionOrCondition) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Operator.matches(other.Operator) {
		return false
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperator) matches(other *SignatureAndConditionOrConditionOperator) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.PatternMatch.matches(other.PatternMatch) {
		return false
	}
	if !o.GreaterThan.matches(other.GreaterThan) {
		return false
	}
	if !o.LessThan.matches(other.LessThan) {
		return false
	}
	if !o.EqualTo.matches(other.EqualTo) {
		return false
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorPatternMatch) matches(other *SignatureAndConditionOrConditionOperatorPatternMatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
		return false
	}
	if !util.StringsMatch(o.Pattern, other.Pattern) {
		return false
	}
	if len(o.Qualifier) != len(other.Qualifier) {
		return false
	}
	for idx := range o.Qualifier {
		if !o.Qualifier[idx].matches(&other.Qualifier[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorPatternMatchQualifier) matches(other *SignatureAndConditionOrConditionOperatorPatternMatchQualifier) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorGreaterThan) matches(other *SignatureAndConditionOrConditionOperatorGreaterThan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if len(o.Qualifier) != len(other.Qualifier) {
		return false
	}
	for idx := range o.Qualifier {
		if !o.Qualifier[idx].matches(&other.Qualifier[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorGreaterThanQualifier) matches(other *SignatureAndConditionOrConditionOperatorGreaterThanQualifier) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorLessThan) matches(other *SignatureAndConditionOrConditionOperatorLessThan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if len(o.Qualifier) != len(other.Qualifier) {
		return false
	}
	for idx := range o.Qualifier {
		if !o.Qualifier[idx].matches(&other.Qualifier[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorLessThanQualifier) matches(other *SignatureAndConditionOrConditionOperatorLessThanQualifier) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}

	return true
}

func (o *SignatureAndConditionOrConditionOperatorEqualTo) matches(other *SignatureAndConditionOrConditionOperatorEqualTo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
		return false
	}
	if !util.StringsMatch(o.Position, other.Position) {
		return false
	}
	if !util.StringsMatch(o.Mask, other.Mask) {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
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
