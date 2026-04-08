package spyware

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
	suffix = []string{"threats", "spyware", "$name"}
)

type Entry struct {
	Name            string
	Bugtraq         []string
	Comment         *string
	Cve             []string
	DefaultAction   *DefaultAction
	Direction       *string
	DisableOverride *string
	Reference       []string
	Severity        *string
	Signature       *Signature
	Threatname      *string
	Vendor          []string
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type DefaultAction struct {
	Alert          *DefaultActionAlert
	Allow          *DefaultActionAllow
	BlockIp        *DefaultActionBlockIp
	Drop           *DefaultActionDrop
	ResetBoth      *DefaultActionResetBoth
	ResetClient    *DefaultActionResetClient
	ResetServer    *DefaultActionResetServer
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionAlert struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionAllow struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionBlockIp struct {
	Duration       *int64
	TrackBy        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionDrop struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionResetBoth struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionResetClient struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultActionResetServer struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Signature struct {
	Combination    *SignatureCombination
	Standard       []SignatureStandard
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureCombination struct {
	AndCondition   []SignatureCombinationAndCondition
	OrderFree      *bool
	TimeAttribute  *SignatureCombinationTimeAttribute
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureCombinationAndCondition struct {
	Name           string
	OrCondition    []SignatureCombinationAndConditionOrCondition
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureCombinationAndConditionOrCondition struct {
	Name           string
	ThreatId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureCombinationTimeAttribute struct {
	Interval       *int64
	Threshold      *int64
	TrackBy        *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandard struct {
	Name           string
	Comment        *string
	Scope          *string
	OrderFree      *bool
	AndCondition   []SignatureStandardAndCondition
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndCondition struct {
	Name           string
	OrCondition    []SignatureStandardAndConditionOrCondition
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrCondition struct {
	Name           string
	Operator       *SignatureStandardAndConditionOrConditionOperator
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperator struct {
	LessThan       *SignatureStandardAndConditionOrConditionOperatorLessThan
	EqualTo        *SignatureStandardAndConditionOrConditionOperatorEqualTo
	GreaterThan    *SignatureStandardAndConditionOrConditionOperatorGreaterThan
	PatternMatch   *SignatureStandardAndConditionOrConditionOperatorPatternMatch
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorLessThan struct {
	Value          *int64
	Context        *string
	Qualifier      []SignatureStandardAndConditionOrConditionOperatorLessThanQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorLessThanQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorEqualTo struct {
	Value          *int64
	Negate         *bool
	Context        *string
	Qualifier      []SignatureStandardAndConditionOrConditionOperatorEqualToQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorEqualToQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorGreaterThan struct {
	Value          *int64
	Context        *string
	Qualifier      []SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier struct {
	Name           string
	Value          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorPatternMatch struct {
	Context        *string
	Pattern        *string
	Negate         *bool
	Qualifier      []SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier struct {
	Name           string
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
	XMLName         xml.Name          `xml:"entry"`
	Name            string            `xml:"name,attr"`
	Bugtraq         *util.MemberType  `xml:"bugtraq,omitempty"`
	Comment         *string           `xml:"comment,omitempty"`
	Cve             *util.MemberType  `xml:"cve,omitempty"`
	DefaultAction   *defaultActionXml `xml:"default-action,omitempty"`
	Direction       *string           `xml:"direction,omitempty"`
	DisableOverride *string           `xml:"disable-override,omitempty"`
	Reference       *util.MemberType  `xml:"reference,omitempty"`
	Severity        *string           `xml:"severity,omitempty"`
	Signature       *signatureXml     `xml:"signature,omitempty"`
	Threatname      *string           `xml:"threatname,omitempty"`
	Vendor          *util.MemberType  `xml:"vendor,omitempty"`
	Misc            []generic.Xml     `xml:",any"`
	MiscAttributes  []xml.Attr        `xml:",any,attr"`
}
type defaultActionXml struct {
	Alert          *defaultActionAlertXml       `xml:"alert,omitempty"`
	Allow          *defaultActionAllowXml       `xml:"allow,omitempty"`
	BlockIp        *defaultActionBlockIpXml     `xml:"block-ip,omitempty"`
	Drop           *defaultActionDropXml        `xml:"drop,omitempty"`
	ResetBoth      *defaultActionResetBothXml   `xml:"reset-both,omitempty"`
	ResetClient    *defaultActionResetClientXml `xml:"reset-client,omitempty"`
	ResetServer    *defaultActionResetServerXml `xml:"reset-server,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type defaultActionAlertXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionAllowXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionBlockIpXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	TrackBy        *string       `xml:"track-by,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionDropXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionResetBothXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionResetClientXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultActionResetServerXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureXml struct {
	Combination    *signatureCombinationXml       `xml:"combination,omitempty"`
	Standard       *signatureStandardContainerXml `xml:"standard,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type signatureCombinationXml struct {
	AndCondition   *signatureCombinationAndConditionContainerXml `xml:"and-condition,omitempty"`
	OrderFree      *string                                       `xml:"order-free,omitempty"`
	TimeAttribute  *signatureCombinationTimeAttributeXml         `xml:"time-attribute,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type signatureCombinationAndConditionContainerXml struct {
	Entries []signatureCombinationAndConditionXml `xml:"entry"`
}
type signatureCombinationAndConditionXml struct {
	XMLName        xml.Name                                                 `xml:"entry"`
	Name           string                                                   `xml:"name,attr"`
	OrCondition    *signatureCombinationAndConditionOrConditionContainerXml `xml:"or-condition,omitempty"`
	Misc           []generic.Xml                                            `xml:",any"`
	MiscAttributes []xml.Attr                                               `xml:",any,attr"`
}
type signatureCombinationAndConditionOrConditionContainerXml struct {
	Entries []signatureCombinationAndConditionOrConditionXml `xml:"entry"`
}
type signatureCombinationAndConditionOrConditionXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	ThreatId       *string       `xml:"threat-id,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureCombinationTimeAttributeXml struct {
	Interval       *int64        `xml:"interval,omitempty"`
	Threshold      *int64        `xml:"threshold,omitempty"`
	TrackBy        *string       `xml:"track-by,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureStandardContainerXml struct {
	Entries []signatureStandardXml `xml:"entry"`
}
type signatureStandardXml struct {
	XMLName        xml.Name                                   `xml:"entry"`
	Name           string                                     `xml:"name,attr"`
	Comment        *string                                    `xml:"comment,omitempty"`
	Scope          *string                                    `xml:"scope,omitempty"`
	OrderFree      *string                                    `xml:"order-free,omitempty"`
	AndCondition   *signatureStandardAndConditionContainerXml `xml:"and-condition,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type signatureStandardAndConditionContainerXml struct {
	Entries []signatureStandardAndConditionXml `xml:"entry"`
}
type signatureStandardAndConditionXml struct {
	XMLName        xml.Name                                              `xml:"entry"`
	Name           string                                                `xml:"name,attr"`
	OrCondition    *signatureStandardAndConditionOrConditionContainerXml `xml:"or-condition,omitempty"`
	Misc           []generic.Xml                                         `xml:",any"`
	MiscAttributes []xml.Attr                                            `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionContainerXml struct {
	Entries []signatureStandardAndConditionOrConditionXml `xml:"entry"`
}
type signatureStandardAndConditionOrConditionXml struct {
	XMLName        xml.Name                                             `xml:"entry"`
	Name           string                                               `xml:"name,attr"`
	Operator       *signatureStandardAndConditionOrConditionOperatorXml `xml:"operator,omitempty"`
	Misc           []generic.Xml                                        `xml:",any"`
	MiscAttributes []xml.Attr                                           `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorXml struct {
	LessThan       *signatureStandardAndConditionOrConditionOperatorLessThanXml     `xml:"less-than,omitempty"`
	EqualTo        *signatureStandardAndConditionOrConditionOperatorEqualToXml      `xml:"equal-to,omitempty"`
	GreaterThan    *signatureStandardAndConditionOrConditionOperatorGreaterThanXml  `xml:"greater-than,omitempty"`
	PatternMatch   *signatureStandardAndConditionOrConditionOperatorPatternMatchXml `xml:"pattern-match,omitempty"`
	Misc           []generic.Xml                                                    `xml:",any"`
	MiscAttributes []xml.Attr                                                       `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorLessThanXml struct {
	Value          *int64                                                                         `xml:"value,omitempty"`
	Context        *string                                                                        `xml:"context,omitempty"`
	Qualifier      *signatureStandardAndConditionOrConditionOperatorLessThanQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                                     `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorLessThanQualifierContainerXml struct {
	Entries []signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml `xml:"entry"`
}
type signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorEqualToXml struct {
	Value          *int64                                                                        `xml:"value,omitempty"`
	Negate         *string                                                                       `xml:"negate,omitempty"`
	Context        *string                                                                       `xml:"context,omitempty"`
	Qualifier      *signatureStandardAndConditionOrConditionOperatorEqualToQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                                 `xml:",any"`
	MiscAttributes []xml.Attr                                                                    `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorEqualToQualifierContainerXml struct {
	Entries []signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml `xml:"entry"`
}
type signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorGreaterThanXml struct {
	Value          *int64                                                                            `xml:"value,omitempty"`
	Context        *string                                                                           `xml:"context,omitempty"`
	Qualifier      *signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                                     `xml:",any"`
	MiscAttributes []xml.Attr                                                                        `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierContainerXml struct {
	Entries []signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml `xml:"entry"`
}
type signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorPatternMatchXml struct {
	Context        *string                                                                            `xml:"context,omitempty"`
	Pattern        *string                                                                            `xml:"pattern,omitempty"`
	Negate         *string                                                                            `xml:"negate,omitempty"`
	Qualifier      *signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierContainerXml `xml:"qualifier,omitempty"`
	Misc           []generic.Xml                                                                      `xml:",any"`
	MiscAttributes []xml.Attr                                                                         `xml:",any,attr"`
}
type signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierContainerXml struct {
	Entries []signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml `xml:"entry"`
}
type signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Bugtraq != nil {
		o.Bugtraq = util.StrToMem(s.Bugtraq)
	}
	o.Comment = s.Comment
	if s.Cve != nil {
		o.Cve = util.StrToMem(s.Cve)
	}
	if s.DefaultAction != nil {
		var obj defaultActionXml
		obj.MarshalFromObject(*s.DefaultAction)
		o.DefaultAction = &obj
	}
	o.Direction = s.Direction
	o.DisableOverride = s.DisableOverride
	if s.Reference != nil {
		o.Reference = util.StrToMem(s.Reference)
	}
	o.Severity = s.Severity
	if s.Signature != nil {
		var obj signatureXml
		obj.MarshalFromObject(*s.Signature)
		o.Signature = &obj
	}
	o.Threatname = s.Threatname
	if s.Vendor != nil {
		o.Vendor = util.StrToMem(s.Vendor)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var bugtraqVal []string
	if o.Bugtraq != nil {
		bugtraqVal = util.MemToStr(o.Bugtraq)
	}
	var cveVal []string
	if o.Cve != nil {
		cveVal = util.MemToStr(o.Cve)
	}
	var defaultActionVal *DefaultAction
	if o.DefaultAction != nil {
		obj, err := o.DefaultAction.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultActionVal = obj
	}
	var referenceVal []string
	if o.Reference != nil {
		referenceVal = util.MemToStr(o.Reference)
	}
	var signatureVal *Signature
	if o.Signature != nil {
		obj, err := o.Signature.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		signatureVal = obj
	}
	var vendorVal []string
	if o.Vendor != nil {
		vendorVal = util.MemToStr(o.Vendor)
	}

	result := &Entry{
		Name:            o.Name,
		Bugtraq:         bugtraqVal,
		Comment:         o.Comment,
		Cve:             cveVal,
		DefaultAction:   defaultActionVal,
		Direction:       o.Direction,
		DisableOverride: o.DisableOverride,
		Reference:       referenceVal,
		Severity:        o.Severity,
		Signature:       signatureVal,
		Threatname:      o.Threatname,
		Vendor:          vendorVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionXml) MarshalFromObject(s DefaultAction) {
	if s.Alert != nil {
		var obj defaultActionAlertXml
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Allow != nil {
		var obj defaultActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.BlockIp != nil {
		var obj defaultActionBlockIpXml
		obj.MarshalFromObject(*s.BlockIp)
		o.BlockIp = &obj
	}
	if s.Drop != nil {
		var obj defaultActionDropXml
		obj.MarshalFromObject(*s.Drop)
		o.Drop = &obj
	}
	if s.ResetBoth != nil {
		var obj defaultActionResetBothXml
		obj.MarshalFromObject(*s.ResetBoth)
		o.ResetBoth = &obj
	}
	if s.ResetClient != nil {
		var obj defaultActionResetClientXml
		obj.MarshalFromObject(*s.ResetClient)
		o.ResetClient = &obj
	}
	if s.ResetServer != nil {
		var obj defaultActionResetServerXml
		obj.MarshalFromObject(*s.ResetServer)
		o.ResetServer = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionXml) UnmarshalToObject() (*DefaultAction, error) {
	var alertVal *DefaultActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var allowVal *DefaultActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var blockIpVal *DefaultActionBlockIp
	if o.BlockIp != nil {
		obj, err := o.BlockIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockIpVal = obj
	}
	var dropVal *DefaultActionDrop
	if o.Drop != nil {
		obj, err := o.Drop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dropVal = obj
	}
	var resetBothVal *DefaultActionResetBoth
	if o.ResetBoth != nil {
		obj, err := o.ResetBoth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetBothVal = obj
	}
	var resetClientVal *DefaultActionResetClient
	if o.ResetClient != nil {
		obj, err := o.ResetClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetClientVal = obj
	}
	var resetServerVal *DefaultActionResetServer
	if o.ResetServer != nil {
		obj, err := o.ResetServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetServerVal = obj
	}

	result := &DefaultAction{
		Alert:          alertVal,
		Allow:          allowVal,
		BlockIp:        blockIpVal,
		Drop:           dropVal,
		ResetBoth:      resetBothVal,
		ResetClient:    resetClientVal,
		ResetServer:    resetServerVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionAlertXml) MarshalFromObject(s DefaultActionAlert) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionAlertXml) UnmarshalToObject() (*DefaultActionAlert, error) {

	result := &DefaultActionAlert{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionAllowXml) MarshalFromObject(s DefaultActionAllow) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionAllowXml) UnmarshalToObject() (*DefaultActionAllow, error) {

	result := &DefaultActionAllow{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionBlockIpXml) MarshalFromObject(s DefaultActionBlockIp) {
	o.Duration = s.Duration
	o.TrackBy = s.TrackBy
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionBlockIpXml) UnmarshalToObject() (*DefaultActionBlockIp, error) {

	result := &DefaultActionBlockIp{
		Duration:       o.Duration,
		TrackBy:        o.TrackBy,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionDropXml) MarshalFromObject(s DefaultActionDrop) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionDropXml) UnmarshalToObject() (*DefaultActionDrop, error) {

	result := &DefaultActionDrop{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionResetBothXml) MarshalFromObject(s DefaultActionResetBoth) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionResetBothXml) UnmarshalToObject() (*DefaultActionResetBoth, error) {

	result := &DefaultActionResetBoth{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionResetClientXml) MarshalFromObject(s DefaultActionResetClient) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionResetClientXml) UnmarshalToObject() (*DefaultActionResetClient, error) {

	result := &DefaultActionResetClient{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultActionResetServerXml) MarshalFromObject(s DefaultActionResetServer) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultActionResetServerXml) UnmarshalToObject() (*DefaultActionResetServer, error) {

	result := &DefaultActionResetServer{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureXml) MarshalFromObject(s Signature) {
	if s.Combination != nil {
		var obj signatureCombinationXml
		obj.MarshalFromObject(*s.Combination)
		o.Combination = &obj
	}
	if s.Standard != nil {
		var objs []signatureStandardXml
		for _, elt := range s.Standard {
			var obj signatureStandardXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Standard = &signatureStandardContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureXml) UnmarshalToObject() (*Signature, error) {
	var combinationVal *SignatureCombination
	if o.Combination != nil {
		obj, err := o.Combination.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		combinationVal = obj
	}
	var standardVal []SignatureStandard
	if o.Standard != nil {
		for _, elt := range o.Standard.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			standardVal = append(standardVal, *obj)
		}
	}

	result := &Signature{
		Combination:    combinationVal,
		Standard:       standardVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureCombinationXml) MarshalFromObject(s SignatureCombination) {
	if s.AndCondition != nil {
		var objs []signatureCombinationAndConditionXml
		for _, elt := range s.AndCondition {
			var obj signatureCombinationAndConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AndCondition = &signatureCombinationAndConditionContainerXml{Entries: objs}
	}
	o.OrderFree = util.YesNo(s.OrderFree, nil)
	if s.TimeAttribute != nil {
		var obj signatureCombinationTimeAttributeXml
		obj.MarshalFromObject(*s.TimeAttribute)
		o.TimeAttribute = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureCombinationXml) UnmarshalToObject() (*SignatureCombination, error) {
	var andConditionVal []SignatureCombinationAndCondition
	if o.AndCondition != nil {
		for _, elt := range o.AndCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			andConditionVal = append(andConditionVal, *obj)
		}
	}
	var timeAttributeVal *SignatureCombinationTimeAttribute
	if o.TimeAttribute != nil {
		obj, err := o.TimeAttribute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		timeAttributeVal = obj
	}

	result := &SignatureCombination{
		AndCondition:   andConditionVal,
		OrderFree:      util.AsBool(o.OrderFree, nil),
		TimeAttribute:  timeAttributeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureCombinationAndConditionXml) MarshalFromObject(s SignatureCombinationAndCondition) {
	o.Name = s.Name
	if s.OrCondition != nil {
		var objs []signatureCombinationAndConditionOrConditionXml
		for _, elt := range s.OrCondition {
			var obj signatureCombinationAndConditionOrConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.OrCondition = &signatureCombinationAndConditionOrConditionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureCombinationAndConditionXml) UnmarshalToObject() (*SignatureCombinationAndCondition, error) {
	var orConditionVal []SignatureCombinationAndConditionOrCondition
	if o.OrCondition != nil {
		for _, elt := range o.OrCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			orConditionVal = append(orConditionVal, *obj)
		}
	}

	result := &SignatureCombinationAndCondition{
		Name:           o.Name,
		OrCondition:    orConditionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureCombinationAndConditionOrConditionXml) MarshalFromObject(s SignatureCombinationAndConditionOrCondition) {
	o.Name = s.Name
	o.ThreatId = s.ThreatId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureCombinationAndConditionOrConditionXml) UnmarshalToObject() (*SignatureCombinationAndConditionOrCondition, error) {

	result := &SignatureCombinationAndConditionOrCondition{
		Name:           o.Name,
		ThreatId:       o.ThreatId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureCombinationTimeAttributeXml) MarshalFromObject(s SignatureCombinationTimeAttribute) {
	o.Interval = s.Interval
	o.Threshold = s.Threshold
	o.TrackBy = s.TrackBy
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureCombinationTimeAttributeXml) UnmarshalToObject() (*SignatureCombinationTimeAttribute, error) {

	result := &SignatureCombinationTimeAttribute{
		Interval:       o.Interval,
		Threshold:      o.Threshold,
		TrackBy:        o.TrackBy,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardXml) MarshalFromObject(s SignatureStandard) {
	o.Name = s.Name
	o.Comment = s.Comment
	o.Scope = s.Scope
	o.OrderFree = util.YesNo(s.OrderFree, nil)
	if s.AndCondition != nil {
		var objs []signatureStandardAndConditionXml
		for _, elt := range s.AndCondition {
			var obj signatureStandardAndConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AndCondition = &signatureStandardAndConditionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardXml) UnmarshalToObject() (*SignatureStandard, error) {
	var andConditionVal []SignatureStandardAndCondition
	if o.AndCondition != nil {
		for _, elt := range o.AndCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			andConditionVal = append(andConditionVal, *obj)
		}
	}

	result := &SignatureStandard{
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
func (o *signatureStandardAndConditionXml) MarshalFromObject(s SignatureStandardAndCondition) {
	o.Name = s.Name
	if s.OrCondition != nil {
		var objs []signatureStandardAndConditionOrConditionXml
		for _, elt := range s.OrCondition {
			var obj signatureStandardAndConditionOrConditionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.OrCondition = &signatureStandardAndConditionOrConditionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionXml) UnmarshalToObject() (*SignatureStandardAndCondition, error) {
	var orConditionVal []SignatureStandardAndConditionOrCondition
	if o.OrCondition != nil {
		for _, elt := range o.OrCondition.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			orConditionVal = append(orConditionVal, *obj)
		}
	}

	result := &SignatureStandardAndCondition{
		Name:           o.Name,
		OrCondition:    orConditionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionXml) MarshalFromObject(s SignatureStandardAndConditionOrCondition) {
	o.Name = s.Name
	if s.Operator != nil {
		var obj signatureStandardAndConditionOrConditionOperatorXml
		obj.MarshalFromObject(*s.Operator)
		o.Operator = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionXml) UnmarshalToObject() (*SignatureStandardAndConditionOrCondition, error) {
	var operatorVal *SignatureStandardAndConditionOrConditionOperator
	if o.Operator != nil {
		obj, err := o.Operator.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operatorVal = obj
	}

	result := &SignatureStandardAndConditionOrCondition{
		Name:           o.Name,
		Operator:       operatorVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperator) {
	if s.LessThan != nil {
		var obj signatureStandardAndConditionOrConditionOperatorLessThanXml
		obj.MarshalFromObject(*s.LessThan)
		o.LessThan = &obj
	}
	if s.EqualTo != nil {
		var obj signatureStandardAndConditionOrConditionOperatorEqualToXml
		obj.MarshalFromObject(*s.EqualTo)
		o.EqualTo = &obj
	}
	if s.GreaterThan != nil {
		var obj signatureStandardAndConditionOrConditionOperatorGreaterThanXml
		obj.MarshalFromObject(*s.GreaterThan)
		o.GreaterThan = &obj
	}
	if s.PatternMatch != nil {
		var obj signatureStandardAndConditionOrConditionOperatorPatternMatchXml
		obj.MarshalFromObject(*s.PatternMatch)
		o.PatternMatch = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperator, error) {
	var lessThanVal *SignatureStandardAndConditionOrConditionOperatorLessThan
	if o.LessThan != nil {
		obj, err := o.LessThan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lessThanVal = obj
	}
	var equalToVal *SignatureStandardAndConditionOrConditionOperatorEqualTo
	if o.EqualTo != nil {
		obj, err := o.EqualTo.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		equalToVal = obj
	}
	var greaterThanVal *SignatureStandardAndConditionOrConditionOperatorGreaterThan
	if o.GreaterThan != nil {
		obj, err := o.GreaterThan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		greaterThanVal = obj
	}
	var patternMatchVal *SignatureStandardAndConditionOrConditionOperatorPatternMatch
	if o.PatternMatch != nil {
		obj, err := o.PatternMatch.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		patternMatchVal = obj
	}

	result := &SignatureStandardAndConditionOrConditionOperator{
		LessThan:       lessThanVal,
		EqualTo:        equalToVal,
		GreaterThan:    greaterThanVal,
		PatternMatch:   patternMatchVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorLessThanXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorLessThan) {
	o.Value = s.Value
	o.Context = s.Context
	if s.Qualifier != nil {
		var objs []signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureStandardAndConditionOrConditionOperatorLessThanQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorLessThanXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorLessThan, error) {
	var qualifierVal []SignatureStandardAndConditionOrConditionOperatorLessThanQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureStandardAndConditionOrConditionOperatorLessThan{
		Value:          o.Value,
		Context:        o.Context,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorLessThanQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorLessThanQualifierXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorLessThanQualifier, error) {

	result := &SignatureStandardAndConditionOrConditionOperatorLessThanQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorEqualToXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorEqualTo) {
	o.Value = s.Value
	o.Negate = util.YesNo(s.Negate, nil)
	o.Context = s.Context
	if s.Qualifier != nil {
		var objs []signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureStandardAndConditionOrConditionOperatorEqualToQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorEqualToXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorEqualTo, error) {
	var qualifierVal []SignatureStandardAndConditionOrConditionOperatorEqualToQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureStandardAndConditionOrConditionOperatorEqualTo{
		Value:          o.Value,
		Negate:         util.AsBool(o.Negate, nil),
		Context:        o.Context,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorEqualToQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorEqualToQualifierXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorEqualToQualifier, error) {

	result := &SignatureStandardAndConditionOrConditionOperatorEqualToQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorGreaterThanXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorGreaterThan) {
	o.Value = s.Value
	o.Context = s.Context
	if s.Qualifier != nil {
		var objs []signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorGreaterThanXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorGreaterThan, error) {
	var qualifierVal []SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureStandardAndConditionOrConditionOperatorGreaterThan{
		Value:          o.Value,
		Context:        o.Context,
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorGreaterThanQualifierXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier, error) {

	result := &SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier{
		Name:           o.Name,
		Value:          o.Value,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorPatternMatchXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorPatternMatch) {
	o.Context = s.Context
	o.Pattern = s.Pattern
	o.Negate = util.YesNo(s.Negate, nil)
	if s.Qualifier != nil {
		var objs []signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml
		for _, elt := range s.Qualifier {
			var obj signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Qualifier = &signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorPatternMatchXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorPatternMatch, error) {
	var qualifierVal []SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier
	if o.Qualifier != nil {
		for _, elt := range o.Qualifier.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			qualifierVal = append(qualifierVal, *obj)
		}
	}

	result := &SignatureStandardAndConditionOrConditionOperatorPatternMatch{
		Context:        o.Context,
		Pattern:        o.Pattern,
		Negate:         util.AsBool(o.Negate, nil),
		Qualifier:      qualifierVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml) MarshalFromObject(s SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier) {
	o.Name = s.Name
	o.Value = s.Value
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o signatureStandardAndConditionOrConditionOperatorPatternMatchQualifierXml) UnmarshalToObject() (*SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier, error) {

	result := &SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier{
		Name:           o.Name,
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
	if v == "bugtraq" || v == "Bugtraq" {
		return e.Bugtraq, nil
	}
	if v == "bugtraq|LENGTH" || v == "Bugtraq|LENGTH" {
		return int64(len(e.Bugtraq)), nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "cve" || v == "Cve" {
		return e.Cve, nil
	}
	if v == "cve|LENGTH" || v == "Cve|LENGTH" {
		return int64(len(e.Cve)), nil
	}
	if v == "default_action" || v == "DefaultAction" {
		return e.DefaultAction, nil
	}
	if v == "direction" || v == "Direction" {
		return e.Direction, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "reference" || v == "Reference" {
		return e.Reference, nil
	}
	if v == "reference|LENGTH" || v == "Reference|LENGTH" {
		return int64(len(e.Reference)), nil
	}
	if v == "severity" || v == "Severity" {
		return e.Severity, nil
	}
	if v == "signature" || v == "Signature" {
		return e.Signature, nil
	}
	if v == "threatname" || v == "Threatname" {
		return e.Threatname, nil
	}
	if v == "vendor" || v == "Vendor" {
		return e.Vendor, nil
	}
	if v == "vendor|LENGTH" || v == "Vendor|LENGTH" {
		return int64(len(e.Vendor)), nil
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
	if !util.OrderedListsMatch[string](o.Bugtraq, other.Bugtraq) {
		return false
	}
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Cve, other.Cve) {
		return false
	}
	if !o.DefaultAction.matches(other.DefaultAction) {
		return false
	}
	if !util.StringsMatch(o.Direction, other.Direction) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Reference, other.Reference) {
		return false
	}
	if !util.StringsMatch(o.Severity, other.Severity) {
		return false
	}
	if !o.Signature.matches(other.Signature) {
		return false
	}
	if !util.StringsMatch(o.Threatname, other.Threatname) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vendor, other.Vendor) {
		return false
	}

	return true
}

func (o *DefaultAction) matches(other *DefaultAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Alert.matches(other.Alert) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}
	if !o.BlockIp.matches(other.BlockIp) {
		return false
	}
	if !o.Drop.matches(other.Drop) {
		return false
	}
	if !o.ResetBoth.matches(other.ResetBoth) {
		return false
	}
	if !o.ResetClient.matches(other.ResetClient) {
		return false
	}
	if !o.ResetServer.matches(other.ResetServer) {
		return false
	}

	return true
}

func (o *DefaultActionAlert) matches(other *DefaultActionAlert) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *DefaultActionAllow) matches(other *DefaultActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *DefaultActionBlockIp) matches(other *DefaultActionBlockIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}
	if !util.StringsMatch(o.TrackBy, other.TrackBy) {
		return false
	}

	return true
}

func (o *DefaultActionDrop) matches(other *DefaultActionDrop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *DefaultActionResetBoth) matches(other *DefaultActionResetBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *DefaultActionResetClient) matches(other *DefaultActionResetClient) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *DefaultActionResetServer) matches(other *DefaultActionResetServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
	if !o.Combination.matches(other.Combination) {
		return false
	}
	if len(o.Standard) != len(other.Standard) {
		return false
	}
	for idx := range o.Standard {
		if !o.Standard[idx].matches(&other.Standard[idx]) {
			return false
		}
	}

	return true
}

func (o *SignatureCombination) matches(other *SignatureCombination) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
	if !util.BoolsMatch(o.OrderFree, other.OrderFree) {
		return false
	}
	if !o.TimeAttribute.matches(other.TimeAttribute) {
		return false
	}

	return true
}

func (o *SignatureCombinationAndCondition) matches(other *SignatureCombinationAndCondition) bool {
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

func (o *SignatureCombinationAndConditionOrCondition) matches(other *SignatureCombinationAndConditionOrCondition) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.ThreatId, other.ThreatId) {
		return false
	}

	return true
}

func (o *SignatureCombinationTimeAttribute) matches(other *SignatureCombinationTimeAttribute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}
	if !util.StringsMatch(o.TrackBy, other.TrackBy) {
		return false
	}

	return true
}

func (o *SignatureStandard) matches(other *SignatureStandard) bool {
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

func (o *SignatureStandardAndCondition) matches(other *SignatureStandardAndCondition) bool {
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

func (o *SignatureStandardAndConditionOrCondition) matches(other *SignatureStandardAndConditionOrCondition) bool {
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

func (o *SignatureStandardAndConditionOrConditionOperator) matches(other *SignatureStandardAndConditionOrConditionOperator) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.LessThan.matches(other.LessThan) {
		return false
	}
	if !o.EqualTo.matches(other.EqualTo) {
		return false
	}
	if !o.GreaterThan.matches(other.GreaterThan) {
		return false
	}
	if !o.PatternMatch.matches(other.PatternMatch) {
		return false
	}

	return true
}

func (o *SignatureStandardAndConditionOrConditionOperatorLessThan) matches(other *SignatureStandardAndConditionOrConditionOperatorLessThan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
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

func (o *SignatureStandardAndConditionOrConditionOperatorLessThanQualifier) matches(other *SignatureStandardAndConditionOrConditionOperatorLessThanQualifier) bool {
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

func (o *SignatureStandardAndConditionOrConditionOperatorEqualTo) matches(other *SignatureStandardAndConditionOrConditionOperatorEqualTo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
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

func (o *SignatureStandardAndConditionOrConditionOperatorEqualToQualifier) matches(other *SignatureStandardAndConditionOrConditionOperatorEqualToQualifier) bool {
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

func (o *SignatureStandardAndConditionOrConditionOperatorGreaterThan) matches(other *SignatureStandardAndConditionOrConditionOperatorGreaterThan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Value, other.Value) {
		return false
	}
	if !util.StringsMatch(o.Context, other.Context) {
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

func (o *SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier) matches(other *SignatureStandardAndConditionOrConditionOperatorGreaterThanQualifier) bool {
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

func (o *SignatureStandardAndConditionOrConditionOperatorPatternMatch) matches(other *SignatureStandardAndConditionOrConditionOperatorPatternMatch) bool {
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
	if !util.BoolsMatch(o.Negate, other.Negate) {
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

func (o *SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier) matches(other *SignatureStandardAndConditionOrConditionOperatorPatternMatchQualifier) bool {
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
