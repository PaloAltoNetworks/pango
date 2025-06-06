package nat

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
	suffix = []string{"nat", "rules", "$name"}
)

type Entry struct {
	Name                          string
	ActiveActiveDeviceBinding     *string
	Description                   *string
	Destination                   []string
	Disabled                      *bool
	From                          []string
	GroupTag                      *string
	NatType                       *string
	Service                       *string
	Source                        []string
	SourceTranslation             *SourceTranslation
	Tag                           []string
	Target                        *Target
	To                            []string
	ToInterface                   *string
	Uuid                          *string
	DestinationTranslation        *DestinationTranslation
	DynamicDestinationTranslation *DynamicDestinationTranslation
	Misc                          []generic.Xml
}
type SourceTranslation struct {
	DynamicIp        *SourceTranslationDynamicIp
	DynamicIpAndPort *SourceTranslationDynamicIpAndPort
	StaticIp         *SourceTranslationStaticIp
	Misc             []generic.Xml
}
type SourceTranslationDynamicIp struct {
	Fallback          *SourceTranslationDynamicIpFallback
	TranslatedAddress []string
	Misc              []generic.Xml
}
type SourceTranslationDynamicIpFallback struct {
	InterfaceAddress  *SourceTranslationDynamicIpFallbackInterfaceAddress
	TranslatedAddress []string
	Misc              []generic.Xml
}
type SourceTranslationDynamicIpFallbackInterfaceAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
	Misc       []generic.Xml
}
type SourceTranslationDynamicIpAndPort struct {
	InterfaceAddress  *SourceTranslationDynamicIpAndPortInterfaceAddress
	TranslatedAddress []string
	Misc              []generic.Xml
}
type SourceTranslationDynamicIpAndPortInterfaceAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
	Misc       []generic.Xml
}
type SourceTranslationStaticIp struct {
	BiDirectional     *string
	TranslatedAddress *string
	Misc              []generic.Xml
}
type Target struct {
	Devices []TargetDevices
	Negate  *bool
	Tags    []string
	Misc    []generic.Xml
}
type TargetDevices struct {
	Name string
	Vsys []TargetDevicesVsys
	Misc []generic.Xml
}
type TargetDevicesVsys struct {
	Name string
	Misc []generic.Xml
}
type DestinationTranslation struct {
	DnsRewrite        *DestinationTranslationDnsRewrite
	TranslatedAddress *string
	TranslatedPort    *int64
	Misc              []generic.Xml
}
type DestinationTranslationDnsRewrite struct {
	Direction *string
	Misc      []generic.Xml
}
type DynamicDestinationTranslation struct {
	Distribution      *string
	TranslatedAddress *string
	TranslatedPort    *int64
	Misc              []generic.Xml
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
	XMLName                       xml.Name                          `xml:"entry"`
	Name                          string                            `xml:"name,attr"`
	ActiveActiveDeviceBinding     *string                           `xml:"active-active-device-binding,omitempty"`
	Description                   *string                           `xml:"description,omitempty"`
	Destination                   *util.MemberType                  `xml:"destination,omitempty"`
	Disabled                      *string                           `xml:"disabled,omitempty"`
	From                          *util.MemberType                  `xml:"from,omitempty"`
	GroupTag                      *string                           `xml:"group-tag,omitempty"`
	NatType                       *string                           `xml:"nat-type,omitempty"`
	Service                       *string                           `xml:"service,omitempty"`
	Source                        *util.MemberType                  `xml:"source,omitempty"`
	SourceTranslation             *sourceTranslationXml             `xml:"source-translation,omitempty"`
	Tag                           *util.MemberType                  `xml:"tag,omitempty"`
	Target                        *targetXml                        `xml:"target,omitempty"`
	To                            *util.MemberType                  `xml:"to,omitempty"`
	ToInterface                   *string                           `xml:"to-interface,omitempty"`
	Uuid                          *string                           `xml:"uuid,attr,omitempty"`
	DestinationTranslation        *destinationTranslationXml        `xml:"destination-translation,omitempty"`
	DynamicDestinationTranslation *dynamicDestinationTranslationXml `xml:"dynamic-destination-translation,omitempty"`
	Misc                          []generic.Xml                     `xml:",any"`
}
type sourceTranslationXml struct {
	DynamicIp        *sourceTranslationDynamicIpXml        `xml:"dynamic-ip,omitempty"`
	DynamicIpAndPort *sourceTranslationDynamicIpAndPortXml `xml:"dynamic-ip-and-port,omitempty"`
	StaticIp         *sourceTranslationStaticIpXml         `xml:"static-ip,omitempty"`
	Misc             []generic.Xml                         `xml:",any"`
}
type sourceTranslationDynamicIpXml struct {
	Fallback          *sourceTranslationDynamicIpFallbackXml `xml:"fallback,omitempty"`
	TranslatedAddress *util.MemberType                       `xml:"translated-address,omitempty"`
	Misc              []generic.Xml                          `xml:",any"`
}
type sourceTranslationDynamicIpFallbackXml struct {
	InterfaceAddress  *sourceTranslationDynamicIpFallbackInterfaceAddressXml `xml:"interface-address,omitempty"`
	TranslatedAddress *util.MemberType                                       `xml:"translated-address,omitempty"`
	Misc              []generic.Xml                                          `xml:",any"`
}
type sourceTranslationDynamicIpFallbackInterfaceAddressXml struct {
	Interface  *string       `xml:"interface,omitempty"`
	FloatingIp *string       `xml:"floating-ip,omitempty"`
	Ip         *string       `xml:"ip,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type sourceTranslationDynamicIpAndPortXml struct {
	InterfaceAddress  *sourceTranslationDynamicIpAndPortInterfaceAddressXml `xml:"interface-address,omitempty"`
	TranslatedAddress *util.MemberType                                      `xml:"translated-address,omitempty"`
	Misc              []generic.Xml                                         `xml:",any"`
}
type sourceTranslationDynamicIpAndPortInterfaceAddressXml struct {
	Interface  *string       `xml:"interface,omitempty"`
	FloatingIp *string       `xml:"floating-ip,omitempty"`
	Ip         *string       `xml:"ip,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type sourceTranslationStaticIpXml struct {
	BiDirectional     *string       `xml:"bi-directional,omitempty"`
	TranslatedAddress *string       `xml:"translated-address,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type targetXml struct {
	Devices *targetDevicesContainerXml `xml:"devices,omitempty"`
	Negate  *string                    `xml:"negate,omitempty"`
	Tags    *util.MemberType           `xml:"tags,omitempty"`
	Misc    []generic.Xml              `xml:",any"`
}
type targetDevicesContainerXml struct {
	Entries []targetDevicesXml `xml:"entry"`
}
type targetDevicesXml struct {
	XMLName xml.Name                       `xml:"entry"`
	Name    string                         `xml:"name,attr"`
	Vsys    *targetDevicesVsysContainerXml `xml:"vsys,omitempty"`
	Misc    []generic.Xml                  `xml:",any"`
}
type targetDevicesVsysContainerXml struct {
	Entries []targetDevicesVsysXml `xml:"entry"`
}
type targetDevicesVsysXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type destinationTranslationXml struct {
	DnsRewrite        *destinationTranslationDnsRewriteXml `xml:"dns-rewrite,omitempty"`
	TranslatedAddress *string                              `xml:"translated-address,omitempty"`
	TranslatedPort    *int64                               `xml:"translated-port,omitempty"`
	Misc              []generic.Xml                        `xml:",any"`
}
type destinationTranslationDnsRewriteXml struct {
	Direction *string       `xml:"direction,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type dynamicDestinationTranslationXml struct {
	Distribution      *string       `xml:"distribution,omitempty"`
	TranslatedAddress *string       `xml:"translated-address,omitempty"`
	TranslatedPort    *int64        `xml:"translated-port,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.ActiveActiveDeviceBinding = s.ActiveActiveDeviceBinding
	o.Description = s.Description
	if s.Destination != nil {
		o.Destination = util.StrToMem(s.Destination)
	}
	o.Disabled = util.YesNo(s.Disabled, nil)
	if s.From != nil {
		o.From = util.StrToMem(s.From)
	}
	o.GroupTag = s.GroupTag
	o.NatType = s.NatType
	o.Service = s.Service
	if s.Source != nil {
		o.Source = util.StrToMem(s.Source)
	}
	if s.SourceTranslation != nil {
		var obj sourceTranslationXml
		obj.MarshalFromObject(*s.SourceTranslation)
		o.SourceTranslation = &obj
	}
	if s.Tag != nil {
		o.Tag = util.StrToMem(s.Tag)
	}
	if s.Target != nil {
		var obj targetXml
		obj.MarshalFromObject(*s.Target)
		o.Target = &obj
	}
	if s.To != nil {
		o.To = util.StrToMem(s.To)
	}
	o.ToInterface = s.ToInterface
	o.Uuid = s.Uuid
	if s.DestinationTranslation != nil {
		var obj destinationTranslationXml
		obj.MarshalFromObject(*s.DestinationTranslation)
		o.DestinationTranslation = &obj
	}
	if s.DynamicDestinationTranslation != nil {
		var obj dynamicDestinationTranslationXml
		obj.MarshalFromObject(*s.DynamicDestinationTranslation)
		o.DynamicDestinationTranslation = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var destinationVal []string
	if o.Destination != nil {
		destinationVal = util.MemToStr(o.Destination)
	}
	var fromVal []string
	if o.From != nil {
		fromVal = util.MemToStr(o.From)
	}
	var sourceVal []string
	if o.Source != nil {
		sourceVal = util.MemToStr(o.Source)
	}
	var sourceTranslationVal *SourceTranslation
	if o.SourceTranslation != nil {
		obj, err := o.SourceTranslation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceTranslationVal = obj
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
	var toVal []string
	if o.To != nil {
		toVal = util.MemToStr(o.To)
	}
	var destinationTranslationVal *DestinationTranslation
	if o.DestinationTranslation != nil {
		obj, err := o.DestinationTranslation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		destinationTranslationVal = obj
	}
	var dynamicDestinationTranslationVal *DynamicDestinationTranslation
	if o.DynamicDestinationTranslation != nil {
		obj, err := o.DynamicDestinationTranslation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicDestinationTranslationVal = obj
	}

	result := &Entry{
		Name:                          o.Name,
		ActiveActiveDeviceBinding:     o.ActiveActiveDeviceBinding,
		Description:                   o.Description,
		Destination:                   destinationVal,
		Disabled:                      util.AsBool(o.Disabled, nil),
		From:                          fromVal,
		GroupTag:                      o.GroupTag,
		NatType:                       o.NatType,
		Service:                       o.Service,
		Source:                        sourceVal,
		SourceTranslation:             sourceTranslationVal,
		Tag:                           tagVal,
		Target:                        targetVal,
		To:                            toVal,
		ToInterface:                   o.ToInterface,
		Uuid:                          o.Uuid,
		DestinationTranslation:        destinationTranslationVal,
		DynamicDestinationTranslation: dynamicDestinationTranslationVal,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationXml) MarshalFromObject(s SourceTranslation) {
	if s.DynamicIp != nil {
		var obj sourceTranslationDynamicIpXml
		obj.MarshalFromObject(*s.DynamicIp)
		o.DynamicIp = &obj
	}
	if s.DynamicIpAndPort != nil {
		var obj sourceTranslationDynamicIpAndPortXml
		obj.MarshalFromObject(*s.DynamicIpAndPort)
		o.DynamicIpAndPort = &obj
	}
	if s.StaticIp != nil {
		var obj sourceTranslationStaticIpXml
		obj.MarshalFromObject(*s.StaticIp)
		o.StaticIp = &obj
	}
	o.Misc = s.Misc
}

func (o sourceTranslationXml) UnmarshalToObject() (*SourceTranslation, error) {
	var dynamicIpVal *SourceTranslationDynamicIp
	if o.DynamicIp != nil {
		obj, err := o.DynamicIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIpVal = obj
	}
	var dynamicIpAndPortVal *SourceTranslationDynamicIpAndPort
	if o.DynamicIpAndPort != nil {
		obj, err := o.DynamicIpAndPort.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicIpAndPortVal = obj
	}
	var staticIpVal *SourceTranslationStaticIp
	if o.StaticIp != nil {
		obj, err := o.StaticIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		staticIpVal = obj
	}

	result := &SourceTranslation{
		DynamicIp:        dynamicIpVal,
		DynamicIpAndPort: dynamicIpAndPortVal,
		StaticIp:         staticIpVal,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationDynamicIpXml) MarshalFromObject(s SourceTranslationDynamicIp) {
	if s.Fallback != nil {
		var obj sourceTranslationDynamicIpFallbackXml
		obj.MarshalFromObject(*s.Fallback)
		o.Fallback = &obj
	}
	if s.TranslatedAddress != nil {
		o.TranslatedAddress = util.StrToMem(s.TranslatedAddress)
	}
	o.Misc = s.Misc
}

func (o sourceTranslationDynamicIpXml) UnmarshalToObject() (*SourceTranslationDynamicIp, error) {
	var fallbackVal *SourceTranslationDynamicIpFallback
	if o.Fallback != nil {
		obj, err := o.Fallback.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fallbackVal = obj
	}
	var translatedAddressVal []string
	if o.TranslatedAddress != nil {
		translatedAddressVal = util.MemToStr(o.TranslatedAddress)
	}

	result := &SourceTranslationDynamicIp{
		Fallback:          fallbackVal,
		TranslatedAddress: translatedAddressVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationDynamicIpFallbackXml) MarshalFromObject(s SourceTranslationDynamicIpFallback) {
	if s.InterfaceAddress != nil {
		var obj sourceTranslationDynamicIpFallbackInterfaceAddressXml
		obj.MarshalFromObject(*s.InterfaceAddress)
		o.InterfaceAddress = &obj
	}
	if s.TranslatedAddress != nil {
		o.TranslatedAddress = util.StrToMem(s.TranslatedAddress)
	}
	o.Misc = s.Misc
}

func (o sourceTranslationDynamicIpFallbackXml) UnmarshalToObject() (*SourceTranslationDynamicIpFallback, error) {
	var interfaceAddressVal *SourceTranslationDynamicIpFallbackInterfaceAddress
	if o.InterfaceAddress != nil {
		obj, err := o.InterfaceAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceAddressVal = obj
	}
	var translatedAddressVal []string
	if o.TranslatedAddress != nil {
		translatedAddressVal = util.MemToStr(o.TranslatedAddress)
	}

	result := &SourceTranslationDynamicIpFallback{
		InterfaceAddress:  interfaceAddressVal,
		TranslatedAddress: translatedAddressVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationDynamicIpFallbackInterfaceAddressXml) MarshalFromObject(s SourceTranslationDynamicIpFallbackInterfaceAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o sourceTranslationDynamicIpFallbackInterfaceAddressXml) UnmarshalToObject() (*SourceTranslationDynamicIpFallbackInterfaceAddress, error) {

	result := &SourceTranslationDynamicIpFallbackInterfaceAddress{
		Interface:  o.Interface,
		FloatingIp: o.FloatingIp,
		Ip:         o.Ip,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationDynamicIpAndPortXml) MarshalFromObject(s SourceTranslationDynamicIpAndPort) {
	if s.InterfaceAddress != nil {
		var obj sourceTranslationDynamicIpAndPortInterfaceAddressXml
		obj.MarshalFromObject(*s.InterfaceAddress)
		o.InterfaceAddress = &obj
	}
	if s.TranslatedAddress != nil {
		o.TranslatedAddress = util.StrToMem(s.TranslatedAddress)
	}
	o.Misc = s.Misc
}

func (o sourceTranslationDynamicIpAndPortXml) UnmarshalToObject() (*SourceTranslationDynamicIpAndPort, error) {
	var interfaceAddressVal *SourceTranslationDynamicIpAndPortInterfaceAddress
	if o.InterfaceAddress != nil {
		obj, err := o.InterfaceAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		interfaceAddressVal = obj
	}
	var translatedAddressVal []string
	if o.TranslatedAddress != nil {
		translatedAddressVal = util.MemToStr(o.TranslatedAddress)
	}

	result := &SourceTranslationDynamicIpAndPort{
		InterfaceAddress:  interfaceAddressVal,
		TranslatedAddress: translatedAddressVal,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationDynamicIpAndPortInterfaceAddressXml) MarshalFromObject(s SourceTranslationDynamicIpAndPortInterfaceAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o sourceTranslationDynamicIpAndPortInterfaceAddressXml) UnmarshalToObject() (*SourceTranslationDynamicIpAndPortInterfaceAddress, error) {

	result := &SourceTranslationDynamicIpAndPortInterfaceAddress{
		Interface:  o.Interface,
		FloatingIp: o.FloatingIp,
		Ip:         o.Ip,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *sourceTranslationStaticIpXml) MarshalFromObject(s SourceTranslationStaticIp) {
	o.BiDirectional = s.BiDirectional
	o.TranslatedAddress = s.TranslatedAddress
	o.Misc = s.Misc
}

func (o sourceTranslationStaticIpXml) UnmarshalToObject() (*SourceTranslationStaticIp, error) {

	result := &SourceTranslationStaticIp{
		BiDirectional:     o.BiDirectional,
		TranslatedAddress: o.TranslatedAddress,
		Misc:              o.Misc,
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
		Devices: devicesVal,
		Negate:  util.AsBool(o.Negate, nil),
		Tags:    tagsVal,
		Misc:    o.Misc,
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
		Name: o.Name,
		Vsys: vsysVal,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *targetDevicesVsysXml) MarshalFromObject(s TargetDevicesVsys) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o targetDevicesVsysXml) UnmarshalToObject() (*TargetDevicesVsys, error) {

	result := &TargetDevicesVsys{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *destinationTranslationXml) MarshalFromObject(s DestinationTranslation) {
	if s.DnsRewrite != nil {
		var obj destinationTranslationDnsRewriteXml
		obj.MarshalFromObject(*s.DnsRewrite)
		o.DnsRewrite = &obj
	}
	o.TranslatedAddress = s.TranslatedAddress
	o.TranslatedPort = s.TranslatedPort
	o.Misc = s.Misc
}

func (o destinationTranslationXml) UnmarshalToObject() (*DestinationTranslation, error) {
	var dnsRewriteVal *DestinationTranslationDnsRewrite
	if o.DnsRewrite != nil {
		obj, err := o.DnsRewrite.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsRewriteVal = obj
	}

	result := &DestinationTranslation{
		DnsRewrite:        dnsRewriteVal,
		TranslatedAddress: o.TranslatedAddress,
		TranslatedPort:    o.TranslatedPort,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *destinationTranslationDnsRewriteXml) MarshalFromObject(s DestinationTranslationDnsRewrite) {
	o.Direction = s.Direction
	o.Misc = s.Misc
}

func (o destinationTranslationDnsRewriteXml) UnmarshalToObject() (*DestinationTranslationDnsRewrite, error) {

	result := &DestinationTranslationDnsRewrite{
		Direction: o.Direction,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *dynamicDestinationTranslationXml) MarshalFromObject(s DynamicDestinationTranslation) {
	o.Distribution = s.Distribution
	o.TranslatedAddress = s.TranslatedAddress
	o.TranslatedPort = s.TranslatedPort
	o.Misc = s.Misc
}

func (o dynamicDestinationTranslationXml) UnmarshalToObject() (*DynamicDestinationTranslation, error) {

	result := &DynamicDestinationTranslation{
		Distribution:      o.Distribution,
		TranslatedAddress: o.TranslatedAddress,
		TranslatedPort:    o.TranslatedPort,
		Misc:              o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "active_active_device_binding" || v == "ActiveActiveDeviceBinding" {
		return e.ActiveActiveDeviceBinding, nil
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
	if v == "from" || v == "From" {
		return e.From, nil
	}
	if v == "from|LENGTH" || v == "From|LENGTH" {
		return int64(len(e.From)), nil
	}
	if v == "group_tag" || v == "GroupTag" {
		return e.GroupTag, nil
	}
	if v == "nat_type" || v == "NatType" {
		return e.NatType, nil
	}
	if v == "service" || v == "Service" {
		return e.Service, nil
	}
	if v == "source" || v == "Source" {
		return e.Source, nil
	}
	if v == "source|LENGTH" || v == "Source|LENGTH" {
		return int64(len(e.Source)), nil
	}
	if v == "source_translation" || v == "SourceTranslation" {
		return e.SourceTranslation, nil
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
	if v == "to" || v == "To" {
		return e.To, nil
	}
	if v == "to|LENGTH" || v == "To|LENGTH" {
		return int64(len(e.To)), nil
	}
	if v == "to_interface" || v == "ToInterface" {
		return e.ToInterface, nil
	}
	if v == "uuid" || v == "Uuid" {
		return e.Uuid, nil
	}
	if v == "destination_translation" || v == "DestinationTranslation" {
		return e.DestinationTranslation, nil
	}
	if v == "dynamic_destination_translation" || v == "DynamicDestinationTranslation" {
		return e.DynamicDestinationTranslation, nil
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
	if !util.StringsMatch(o.ActiveActiveDeviceBinding, other.ActiveActiveDeviceBinding) {
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
	if !util.OrderedListsMatch[string](o.From, other.From) {
		return false
	}
	if !util.StringsMatch(o.GroupTag, other.GroupTag) {
		return false
	}
	if !util.StringsMatch(o.NatType, other.NatType) {
		return false
	}
	if !util.StringsMatch(o.Service, other.Service) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Source, other.Source) {
		return false
	}
	if !o.SourceTranslation.matches(other.SourceTranslation) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tag, other.Tag) {
		return false
	}
	if !o.Target.matches(other.Target) {
		return false
	}
	if !util.OrderedListsMatch[string](o.To, other.To) {
		return false
	}
	if !util.StringsMatch(o.ToInterface, other.ToInterface) {
		return false
	}
	if !util.StringsMatch(o.Uuid, other.Uuid) {
		return false
	}
	if !o.DestinationTranslation.matches(other.DestinationTranslation) {
		return false
	}
	if !o.DynamicDestinationTranslation.matches(other.DynamicDestinationTranslation) {
		return false
	}

	return true
}

func (o *SourceTranslation) matches(other *SourceTranslation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.DynamicIp.matches(other.DynamicIp) {
		return false
	}
	if !o.DynamicIpAndPort.matches(other.DynamicIpAndPort) {
		return false
	}
	if !o.StaticIp.matches(other.StaticIp) {
		return false
	}

	return true
}

func (o *SourceTranslationDynamicIp) matches(other *SourceTranslationDynamicIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Fallback.matches(other.Fallback) {
		return false
	}
	if !util.OrderedListsMatch[string](o.TranslatedAddress, other.TranslatedAddress) {
		return false
	}

	return true
}

func (o *SourceTranslationDynamicIpFallback) matches(other *SourceTranslationDynamicIpFallback) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.InterfaceAddress.matches(other.InterfaceAddress) {
		return false
	}
	if !util.OrderedListsMatch[string](o.TranslatedAddress, other.TranslatedAddress) {
		return false
	}

	return true
}

func (o *SourceTranslationDynamicIpFallbackInterfaceAddress) matches(other *SourceTranslationDynamicIpFallbackInterfaceAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.FloatingIp, other.FloatingIp) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *SourceTranslationDynamicIpAndPort) matches(other *SourceTranslationDynamicIpAndPort) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.InterfaceAddress.matches(other.InterfaceAddress) {
		return false
	}
	if !util.OrderedListsMatch[string](o.TranslatedAddress, other.TranslatedAddress) {
		return false
	}

	return true
}

func (o *SourceTranslationDynamicIpAndPortInterfaceAddress) matches(other *SourceTranslationDynamicIpAndPortInterfaceAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.FloatingIp, other.FloatingIp) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *SourceTranslationStaticIp) matches(other *SourceTranslationStaticIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.BiDirectional, other.BiDirectional) {
		return false
	}
	if !util.StringsMatch(o.TranslatedAddress, other.TranslatedAddress) {
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

func (o *DestinationTranslation) matches(other *DestinationTranslation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.DnsRewrite.matches(other.DnsRewrite) {
		return false
	}
	if !util.StringsMatch(o.TranslatedAddress, other.TranslatedAddress) {
		return false
	}
	if !util.Ints64Match(o.TranslatedPort, other.TranslatedPort) {
		return false
	}

	return true
}

func (o *DestinationTranslationDnsRewrite) matches(other *DestinationTranslationDnsRewrite) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Direction, other.Direction) {
		return false
	}

	return true
}

func (o *DynamicDestinationTranslation) matches(other *DynamicDestinationTranslation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Distribution, other.Distribution) {
		return false
	}
	if !util.StringsMatch(o.TranslatedAddress, other.TranslatedAddress) {
		return false
	}
	if !util.Ints64Match(o.TranslatedPort, other.TranslatedPort) {
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
func (o *Entry) EntryUuid() *string {
	return o.Uuid
}

func (o *Entry) SetEntryUuid(uuid *string) {
	o.Uuid = uuid
}
