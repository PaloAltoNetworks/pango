package pbf

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
	suffix = []string{"pbf", "rules", "$name"}
)

type Entry struct {
	Name                      string
	Action                    *Action
	ActiveActiveDeviceBinding *string
	Application               []string
	Description               *string
	Destination               []string
	Disabled                  *bool
	EnforceSymmetricReturn    *EnforceSymmetricReturn
	From                      *From
	GroupTag                  *string
	NegateDestination         *bool
	NegateSource              *bool
	Schedule                  *string
	Service                   []string
	Source                    []string
	SourceUser                []string
	Tag                       []string
	Target                    *Target
	Uuid                      *string
	Misc                      []generic.Xml
	MiscAttributes            []xml.Attr
}
type Action struct {
	Discard        *ActionDiscard
	Forward        *ActionForward
	ForwardToVsys  *string
	NoPbf          *ActionNoPbf
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionDiscard struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionForward struct {
	EgressInterface *string
	Monitor         *ActionForwardMonitor
	Nexthop         *ActionForwardNexthop
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type ActionForwardMonitor struct {
	DisableIfUnreachable *bool
	IpAddress            *string
	Profile              *string
	Misc                 []generic.Xml
	MiscAttributes       []xml.Attr
}
type ActionForwardNexthop struct {
	Fqdn           *string
	IpAddress      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ActionNoPbf struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EnforceSymmetricReturn struct {
	Enabled            *bool
	NexthopAddressList []EnforceSymmetricReturnNexthopAddressList
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type EnforceSymmetricReturnNexthopAddressList struct {
	Name           string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type From struct {
	Interface      []string
	Zone           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Target struct {
	Devices        []TargetDevices
	Negate         *bool
	Tags           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TargetDevices struct {
	Name           string
	Vsys           []TargetDevicesVsys
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TargetDevicesVsys struct {
	Name           string
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
	XMLName                   xml.Name                   `xml:"entry"`
	Name                      string                     `xml:"name,attr"`
	Action                    *actionXml                 `xml:"action,omitempty"`
	ActiveActiveDeviceBinding *string                    `xml:"active-active-device-binding,omitempty"`
	Application               *util.MemberType           `xml:"application,omitempty"`
	Description               *string                    `xml:"description,omitempty"`
	Destination               *util.MemberType           `xml:"destination,omitempty"`
	Disabled                  *string                    `xml:"disabled,omitempty"`
	EnforceSymmetricReturn    *enforceSymmetricReturnXml `xml:"enforce-symmetric-return,omitempty"`
	From                      *fromXml                   `xml:"from,omitempty"`
	GroupTag                  *string                    `xml:"group-tag,omitempty"`
	NegateDestination         *string                    `xml:"negate-destination,omitempty"`
	NegateSource              *string                    `xml:"negate-source,omitempty"`
	Schedule                  *string                    `xml:"schedule,omitempty"`
	Service                   *util.MemberType           `xml:"service,omitempty"`
	Source                    *util.MemberType           `xml:"source,omitempty"`
	SourceUser                *util.MemberType           `xml:"source-user,omitempty"`
	Tag                       *util.MemberType           `xml:"tag,omitempty"`
	Target                    *targetXml                 `xml:"target,omitempty"`
	Uuid                      *string                    `xml:"uuid,attr,omitempty"`
	Misc                      []generic.Xml              `xml:",any"`
	MiscAttributes            []xml.Attr                 `xml:",any,attr"`
}
type actionXml struct {
	Discard        *actionDiscardXml `xml:"discard,omitempty"`
	Forward        *actionForwardXml `xml:"forward,omitempty"`
	ForwardToVsys  *string           `xml:"forward-to-vsys,omitempty"`
	NoPbf          *actionNoPbfXml   `xml:"no-pbf,omitempty"`
	Misc           []generic.Xml     `xml:",any"`
	MiscAttributes []xml.Attr        `xml:",any,attr"`
}
type actionDiscardXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type actionForwardXml struct {
	EgressInterface *string                  `xml:"egress-interface,omitempty"`
	Monitor         *actionForwardMonitorXml `xml:"monitor,omitempty"`
	Nexthop         *actionForwardNexthopXml `xml:"nexthop,omitempty"`
	Misc            []generic.Xml            `xml:",any"`
	MiscAttributes  []xml.Attr               `xml:",any,attr"`
}
type actionForwardMonitorXml struct {
	DisableIfUnreachable *string       `xml:"disable-if-unreachable,omitempty"`
	IpAddress            *string       `xml:"ip-address,omitempty"`
	Profile              *string       `xml:"profile,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
	MiscAttributes       []xml.Attr    `xml:",any,attr"`
}
type actionForwardNexthopXml struct {
	Fqdn           *string       `xml:"fqdn,omitempty"`
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type actionNoPbfXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type enforceSymmetricReturnXml struct {
	Enabled            *string                                               `xml:"enabled,omitempty"`
	NexthopAddressList *enforceSymmetricReturnNexthopAddressListContainerXml `xml:"nexthop-address-list,omitempty"`
	Misc               []generic.Xml                                         `xml:",any"`
	MiscAttributes     []xml.Attr                                            `xml:",any,attr"`
}
type enforceSymmetricReturnNexthopAddressListContainerXml struct {
	Entries []enforceSymmetricReturnNexthopAddressListXml `xml:"entry"`
}
type enforceSymmetricReturnNexthopAddressListXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type fromXml struct {
	Interface      *util.MemberType `xml:"interface,omitempty"`
	Zone           *util.MemberType `xml:"zone,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type targetXml struct {
	Devices        *targetDevicesContainerXml `xml:"devices,omitempty"`
	Negate         *string                    `xml:"negate,omitempty"`
	Tags           *util.MemberType           `xml:"tags,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type targetDevicesContainerXml struct {
	Entries []targetDevicesXml `xml:"entry"`
}
type targetDevicesXml struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Vsys           *targetDevicesVsysContainerXml `xml:"vsys,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type targetDevicesVsysContainerXml struct {
	Entries []targetDevicesVsysXml `xml:"entry"`
}
type targetDevicesVsysXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Action != nil {
		var obj actionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.ActiveActiveDeviceBinding = s.ActiveActiveDeviceBinding
	if s.Application != nil {
		o.Application = util.StrToMem(s.Application)
	}
	o.Description = s.Description
	if s.Destination != nil {
		o.Destination = util.StrToMem(s.Destination)
	}
	o.Disabled = util.YesNo(s.Disabled, nil)
	if s.EnforceSymmetricReturn != nil {
		var obj enforceSymmetricReturnXml
		obj.MarshalFromObject(*s.EnforceSymmetricReturn)
		o.EnforceSymmetricReturn = &obj
	}
	if s.From != nil {
		var obj fromXml
		obj.MarshalFromObject(*s.From)
		o.From = &obj
	}
	o.GroupTag = s.GroupTag
	o.NegateDestination = util.YesNo(s.NegateDestination, nil)
	o.NegateSource = util.YesNo(s.NegateSource, nil)
	o.Schedule = s.Schedule
	if s.Service != nil {
		o.Service = util.StrToMem(s.Service)
	}
	if s.Source != nil {
		o.Source = util.StrToMem(s.Source)
	}
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	if s.Target != nil {
		var obj targetXml
		obj.MarshalFromObject(*s.Target)
		o.Target = &obj
	}
	o.Uuid = s.Uuid
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var actionVal *Action
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}
	var applicationVal []string
	if o.Application != nil {
		applicationVal = util.MemToStr(o.Application)
	}
	var destinationVal []string
	if o.Destination != nil {
		destinationVal = util.MemToStr(o.Destination)
	}
	var enforceSymmetricReturnVal *EnforceSymmetricReturn
	if o.EnforceSymmetricReturn != nil {
		obj, err := o.EnforceSymmetricReturn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		enforceSymmetricReturnVal = obj
	}
	var fromVal *From
	if o.From != nil {
		obj, err := o.From.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fromVal = obj
	}
	var serviceVal []string
	if o.Service != nil {
		serviceVal = util.MemToStr(o.Service)
	}
	var sourceVal []string
	if o.Source != nil {
		sourceVal = util.MemToStr(o.Source)
	}
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var tagVal []string
	if o.Tag != nil {
		tagVal = util.MemToStr(o.Tag)
	}
	var targetVal *Target
	if o.Target != nil {
		obj, err := o.Target.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		targetVal = obj
	}

	result := &Entry{
		Name:                      o.Name,
		Action:                    actionVal,
		ActiveActiveDeviceBinding: o.ActiveActiveDeviceBinding,
		Application:               applicationVal,
		Description:               o.Description,
		Destination:               destinationVal,
		Disabled:                  util.AsBool(o.Disabled, nil),
		EnforceSymmetricReturn:    enforceSymmetricReturnVal,
		From:                      fromVal,
		GroupTag:                  o.GroupTag,
		NegateDestination:         util.AsBool(o.NegateDestination, nil),
		NegateSource:              util.AsBool(o.NegateSource, nil),
		Schedule:                  o.Schedule,
		Service:                   serviceVal,
		Source:                    sourceVal,
		SourceUser:                sourceUserVal,
		Tag:                       tagVal,
		Target:                    targetVal,
		Uuid:                      o.Uuid,
		Misc:                      o.Misc,
		MiscAttributes:            o.MiscAttributes,
	}
	return result, nil
}
func (o *actionXml) MarshalFromObject(s Action) {
	if s.Discard != nil {
		var obj actionDiscardXml
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	if s.Forward != nil {
		var obj actionForwardXml
		obj.MarshalFromObject(*s.Forward)
		o.Forward = &obj
	}
	o.ForwardToVsys = s.ForwardToVsys
	if s.NoPbf != nil {
		var obj actionNoPbfXml
		obj.MarshalFromObject(*s.NoPbf)
		o.NoPbf = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionXml) UnmarshalToObject() (*Action, error) {
	var discardVal *ActionDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}
	var forwardVal *ActionForward
	if o.Forward != nil {
		obj, err := o.Forward.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		forwardVal = obj
	}
	var noPbfVal *ActionNoPbf
	if o.NoPbf != nil {
		obj, err := o.NoPbf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noPbfVal = obj
	}

	result := &Action{
		Discard:        discardVal,
		Forward:        forwardVal,
		ForwardToVsys:  o.ForwardToVsys,
		NoPbf:          noPbfVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionDiscardXml) MarshalFromObject(s ActionDiscard) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionDiscardXml) UnmarshalToObject() (*ActionDiscard, error) {

	result := &ActionDiscard{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionForwardXml) MarshalFromObject(s ActionForward) {
	o.EgressInterface = s.EgressInterface
	if s.Monitor != nil {
		var obj actionForwardMonitorXml
		obj.MarshalFromObject(*s.Monitor)
		o.Monitor = &obj
	}
	if s.Nexthop != nil {
		var obj actionForwardNexthopXml
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionForwardXml) UnmarshalToObject() (*ActionForward, error) {
	var monitorVal *ActionForwardMonitor
	if o.Monitor != nil {
		obj, err := o.Monitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monitorVal = obj
	}
	var nexthopVal *ActionForwardNexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}

	result := &ActionForward{
		EgressInterface: o.EgressInterface,
		Monitor:         monitorVal,
		Nexthop:         nexthopVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *actionForwardMonitorXml) MarshalFromObject(s ActionForwardMonitor) {
	o.DisableIfUnreachable = util.YesNo(s.DisableIfUnreachable, nil)
	o.IpAddress = s.IpAddress
	o.Profile = s.Profile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionForwardMonitorXml) UnmarshalToObject() (*ActionForwardMonitor, error) {

	result := &ActionForwardMonitor{
		DisableIfUnreachable: util.AsBool(o.DisableIfUnreachable, nil),
		IpAddress:            o.IpAddress,
		Profile:              o.Profile,
		Misc:                 o.Misc,
		MiscAttributes:       o.MiscAttributes,
	}
	return result, nil
}
func (o *actionForwardNexthopXml) MarshalFromObject(s ActionForwardNexthop) {
	o.Fqdn = s.Fqdn
	o.IpAddress = s.IpAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionForwardNexthopXml) UnmarshalToObject() (*ActionForwardNexthop, error) {

	result := &ActionForwardNexthop{
		Fqdn:           o.Fqdn,
		IpAddress:      o.IpAddress,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *actionNoPbfXml) MarshalFromObject(s ActionNoPbf) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o actionNoPbfXml) UnmarshalToObject() (*ActionNoPbf, error) {

	result := &ActionNoPbf{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *enforceSymmetricReturnXml) MarshalFromObject(s EnforceSymmetricReturn) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.NexthopAddressList != nil {
		var objs []enforceSymmetricReturnNexthopAddressListXml
		for _, elt := range s.NexthopAddressList {
			var obj enforceSymmetricReturnNexthopAddressListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.NexthopAddressList = &enforceSymmetricReturnNexthopAddressListContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o enforceSymmetricReturnXml) UnmarshalToObject() (*EnforceSymmetricReturn, error) {
	var nexthopAddressListVal []EnforceSymmetricReturnNexthopAddressList
	if o.NexthopAddressList != nil {
		for _, elt := range o.NexthopAddressList.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			nexthopAddressListVal = append(nexthopAddressListVal, *obj)
		}
	}

	result := &EnforceSymmetricReturn{
		Enabled:            util.AsBool(o.Enabled, nil),
		NexthopAddressList: nexthopAddressListVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *enforceSymmetricReturnNexthopAddressListXml) MarshalFromObject(s EnforceSymmetricReturnNexthopAddressList) {
	o.Name = s.Name
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o enforceSymmetricReturnNexthopAddressListXml) UnmarshalToObject() (*EnforceSymmetricReturnNexthopAddressList, error) {

	result := &EnforceSymmetricReturnNexthopAddressList{
		Name:           o.Name,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *fromXml) MarshalFromObject(s From) {
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.Zone != nil {
		o.Zone = util.StrToMem(s.Zone)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o fromXml) UnmarshalToObject() (*From, error) {
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var zoneVal []string
	if o.Zone != nil {
		zoneVal = util.MemToStr(o.Zone)
	}

	result := &From{
		Interface:      interfaceVal,
		Zone:           zoneVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *targetXml) MarshalFromObject(s Target) {
	if s.Devices != nil {
		var objs []targetDevicesXml
		for _, elt := range s.Devices {
			var obj targetDevicesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Devices = &targetDevicesContainerXml{Entries: objs}
	}
	o.Negate = util.YesNo(s.Negate, nil)
	if s.Tags != nil {
		o.Tags = util.StrToMem(s.Tags)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetXml) UnmarshalToObject() (*Target, error) {
	var devicesVal []TargetDevices
	if o.Devices != nil {
		for _, elt := range o.Devices.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			devicesVal = append(devicesVal, *obj)
		}
	}
	var tagsVal []string
	if o.Tags != nil {
		tagsVal = util.MemToStr(o.Tags)
	}

	result := &Target{
		Devices:        devicesVal,
		Negate:         util.AsBool(o.Negate, nil),
		Tags:           tagsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *targetDevicesXml) MarshalFromObject(s TargetDevices) {
	o.Name = s.Name
	if s.Vsys != nil {
		var objs []targetDevicesVsysXml
		for _, elt := range s.Vsys {
			var obj targetDevicesVsysXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vsys = &targetDevicesVsysContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetDevicesXml) UnmarshalToObject() (*TargetDevices, error) {
	var vsysVal []TargetDevicesVsys
	if o.Vsys != nil {
		for _, elt := range o.Vsys.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vsysVal = append(vsysVal, *obj)
		}
	}

	result := &TargetDevices{
		Name:           o.Name,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *targetDevicesVsysXml) MarshalFromObject(s TargetDevicesVsys) {
	o.Name = s.Name
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o targetDevicesVsysXml) UnmarshalToObject() (*TargetDevicesVsys, error) {

	result := &TargetDevicesVsys{
		Name:           o.Name,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
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
	if v == "active_active_device_binding" || v == "ActiveActiveDeviceBinding" {
		return e.ActiveActiveDeviceBinding, nil
	}
	if v == "application" || v == "Application" {
		return e.Application, nil
	}
	if v == "application|LENGTH" || v == "Application|LENGTH" {
		return int64(len(e.Application)), nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "destination" || v == "Destination" {
		return e.Destination, nil
	}
	if v == "destination|LENGTH" || v == "Destination|LENGTH" {
		return int64(len(e.Destination)), nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "enforce_symmetric_return" || v == "EnforceSymmetricReturn" {
		return e.EnforceSymmetricReturn, nil
	}
	if v == "from" || v == "From" {
		return e.From, nil
	}
	if v == "group_tag" || v == "GroupTag" {
		return e.GroupTag, nil
	}
	if v == "negate_destination" || v == "NegateDestination" {
		return e.NegateDestination, nil
	}
	if v == "negate_source" || v == "NegateSource" {
		return e.NegateSource, nil
	}
	if v == "schedule" || v == "Schedule" {
		return e.Schedule, nil
	}
	if v == "service" || v == "Service" {
		return e.Service, nil
	}
	if v == "service|LENGTH" || v == "Service|LENGTH" {
		return int64(len(e.Service)), nil
	}
	if v == "source" || v == "Source" {
		return e.Source, nil
	}
	if v == "source|LENGTH" || v == "Source|LENGTH" {
		return int64(len(e.Source)), nil
	}
	if v == "source_user" || v == "SourceUser" {
		return e.SourceUser, nil
	}
	if v == "source_user|LENGTH" || v == "SourceUser|LENGTH" {
		return int64(len(e.SourceUser)), nil
	}
	if v == "tag" || v == "Tag" {
		return e.Tag, nil
	}
	if v == "tag|LENGTH" || v == "Tag|LENGTH" {
		return int64(len(e.Tag)), nil
	}
	if v == "target" || v == "Target" {
		return e.Target, nil
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
	if !o.Action.matches(other.Action) {
		return false
	}
	if !util.StringsMatch(o.ActiveActiveDeviceBinding, other.ActiveActiveDeviceBinding) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Application, other.Application) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Destination, other.Destination) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !o.EnforceSymmetricReturn.matches(other.EnforceSymmetricReturn) {
		return false
	}
	if !o.From.matches(other.From) {
		return false
	}
	if !util.StringsMatch(o.GroupTag, other.GroupTag) {
		return false
	}
	if !util.BoolsMatch(o.NegateDestination, other.NegateDestination) {
		return false
	}
	if !util.BoolsMatch(o.NegateSource, other.NegateSource) {
		return false
	}
	if !util.StringsMatch(o.Schedule, other.Schedule) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Service, other.Service) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Source, other.Source) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceUser, other.SourceUser) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}
	if !o.Target.matches(other.Target) {
		return false
	}
	if !util.StringsMatch(o.Uuid, other.Uuid) {
		return false
	}

	return true
}

func (o *Action) matches(other *Action) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Discard.matches(other.Discard) {
		return false
	}
	if !o.Forward.matches(other.Forward) {
		return false
	}
	if !util.StringsMatch(o.ForwardToVsys, other.ForwardToVsys) {
		return false
	}
	if !o.NoPbf.matches(other.NoPbf) {
		return false
	}

	return true
}

func (o *ActionDiscard) matches(other *ActionDiscard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ActionForward) matches(other *ActionForward) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EgressInterface, other.EgressInterface) {
		return false
	}
	if !o.Monitor.matches(other.Monitor) {
		return false
	}
	if !o.Nexthop.matches(other.Nexthop) {
		return false
	}

	return true
}

func (o *ActionForwardMonitor) matches(other *ActionForwardMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.DisableIfUnreachable, other.DisableIfUnreachable) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *ActionForwardNexthop) matches(other *ActionForwardNexthop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}

	return true
}

func (o *ActionNoPbf) matches(other *ActionNoPbf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *EnforceSymmetricReturn) matches(other *EnforceSymmetricReturn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if len(o.NexthopAddressList) != len(other.NexthopAddressList) {
		return false
	}
	for idx := range o.NexthopAddressList {
		if !o.NexthopAddressList[idx].matches(&other.NexthopAddressList[idx]) {
			return false
		}
	}

	return true
}

func (o *EnforceSymmetricReturnNexthopAddressList) matches(other *EnforceSymmetricReturnNexthopAddressList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *From) matches(other *From) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Zone, other.Zone) {
		return false
	}

	return true
}

func (o *Target) matches(other *Target) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Devices) != len(other.Devices) {
		return false
	}
	for idx := range o.Devices {
		if !o.Devices[idx].matches(&other.Devices[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tags, other.Tags) {
		return false
	}

	return true
}

func (o *TargetDevices) matches(other *TargetDevices) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.Vsys) != len(other.Vsys) {
		return false
	}
	for idx := range o.Vsys {
		if !o.Vsys[idx].matches(&other.Vsys[idx]) {
			return false
		}
	}

	return true
}

func (o *TargetDevicesVsys) matches(other *TargetDevicesVsys) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
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
