package gateway

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
	suffix = []string{"network", "ike", "gateway", "$name"}
)

type Entry struct {
	Name           string
	Authentication *Authentication
	Comment        *string
	Disabled       *bool
	Ipv6           *bool
	LocalAddress   *LocalAddress
	LocalId        *LocalId
	PeerAddress    *PeerAddress
	PeerId         *PeerId
	Protocol       *Protocol
	ProtocolCommon *ProtocolCommon
	Misc           []generic.Xml
}
type Authentication struct {
	Certificate  *AuthenticationCertificate
	PreSharedKey *AuthenticationPreSharedKey
	Misc         []generic.Xml
}
type AuthenticationCertificate struct {
	AllowIdPayloadMismatch     *bool
	CertificateProfile         *string
	LocalCertificate           *AuthenticationCertificateLocalCertificate
	StrictValidationRevocation *bool
	UseManagementAsSource      *bool
	Misc                       []generic.Xml
}
type AuthenticationCertificateLocalCertificate struct {
	HashAndUrl *AuthenticationCertificateLocalCertificateHashAndUrl
	Name       *string
	Misc       []generic.Xml
}
type AuthenticationCertificateLocalCertificateHashAndUrl struct {
	BaseUrl *string
	Enable  *bool
	Misc    []generic.Xml
}
type AuthenticationPreSharedKey struct {
	Key  *string
	Misc []generic.Xml
}
type LocalAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
	Misc       []generic.Xml
}
type LocalId struct {
	Id   *string
	Type *string
	Misc []generic.Xml
}
type PeerAddress struct {
	Dynamic *PeerAddressDynamic
	Fqdn    *string
	Ip      *string
	Misc    []generic.Xml
}
type PeerAddressDynamic struct {
	Misc []generic.Xml
}
type PeerId struct {
	Id       *string
	Matching *string
	Type     *string
	Misc     []generic.Xml
}
type Protocol struct {
	Ikev1   *ProtocolIkev1
	Ikev2   *ProtocolIkev2
	Version *string
	Misc    []generic.Xml
}
type ProtocolIkev1 struct {
	Dpd              *ProtocolIkev1Dpd
	ExchangeMode     *string
	IkeCryptoProfile *string
	Misc             []generic.Xml
}
type ProtocolIkev1Dpd struct {
	Enable   *bool
	Interval *int64
	Retry    *int64
	Misc     []generic.Xml
}
type ProtocolIkev2 struct {
	Dpd              *ProtocolIkev2Dpd
	IkeCryptoProfile *string
	RequireCookie    *bool
	Misc             []generic.Xml
}
type ProtocolIkev2Dpd struct {
	Enable   *bool
	Interval *int64
	Misc     []generic.Xml
}
type ProtocolCommon struct {
	Fragmentation *ProtocolCommonFragmentation
	NatTraversal  *ProtocolCommonNatTraversal
	PassiveMode   *bool
	Misc          []generic.Xml
}
type ProtocolCommonFragmentation struct {
	Enable *bool
	Misc   []generic.Xml
}
type ProtocolCommonNatTraversal struct {
	Enable            *bool
	KeepAliveInterval *int64
	UdpChecksumEnable *bool
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
	XMLName        xml.Name           `xml:"entry"`
	Name           string             `xml:"name,attr"`
	Authentication *authenticationXml `xml:"authentication,omitempty"`
	Comment        *string            `xml:"comment,omitempty"`
	Disabled       *string            `xml:"disabled,omitempty"`
	Ipv6           *string            `xml:"ipv6,omitempty"`
	LocalAddress   *localAddressXml   `xml:"local-address,omitempty"`
	LocalId        *localIdXml        `xml:"local-id,omitempty"`
	PeerAddress    *peerAddressXml    `xml:"peer-address,omitempty"`
	PeerId         *peerIdXml         `xml:"peer-id,omitempty"`
	Protocol       *protocolXml       `xml:"protocol,omitempty"`
	ProtocolCommon *protocolCommonXml `xml:"protocol-common,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
}
type authenticationXml struct {
	Certificate  *authenticationCertificateXml  `xml:"certificate,omitempty"`
	PreSharedKey *authenticationPreSharedKeyXml `xml:"pre-shared-key,omitempty"`
	Misc         []generic.Xml                  `xml:",any"`
}
type authenticationCertificateXml struct {
	AllowIdPayloadMismatch     *string                                       `xml:"allow-id-payload-mismatch,omitempty"`
	CertificateProfile         *string                                       `xml:"certificate-profile,omitempty"`
	LocalCertificate           *authenticationCertificateLocalCertificateXml `xml:"local-certificate,omitempty"`
	StrictValidationRevocation *string                                       `xml:"strict-validation-revocation,omitempty"`
	UseManagementAsSource      *string                                       `xml:"use-management-as-source,omitempty"`
	Misc                       []generic.Xml                                 `xml:",any"`
}
type authenticationCertificateLocalCertificateXml struct {
	HashAndUrl *authenticationCertificateLocalCertificateHashAndUrlXml `xml:"hash-and-url,omitempty"`
	Name       *string                                                 `xml:"name,attr,omitempty"`
	Misc       []generic.Xml                                           `xml:",any"`
}
type authenticationCertificateLocalCertificateHashAndUrlXml struct {
	BaseUrl *string       `xml:"base-url,omitempty"`
	Enable  *string       `xml:"enable,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type authenticationPreSharedKeyXml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type localAddressXml struct {
	Interface  *string       `xml:"interface,omitempty"`
	FloatingIp *string       `xml:"floating-ip,omitempty"`
	Ip         *string       `xml:"ip,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type localIdXml struct {
	Id   *string       `xml:"id,omitempty"`
	Type *string       `xml:"type,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type peerAddressXml struct {
	Dynamic *peerAddressDynamicXml `xml:"dynamic,omitempty"`
	Fqdn    *string                `xml:"fqdn,omitempty"`
	Ip      *string                `xml:"ip,omitempty"`
	Misc    []generic.Xml          `xml:",any"`
}
type peerAddressDynamicXml struct {
	Misc []generic.Xml `xml:",any"`
}
type peerIdXml struct {
	Id       *string       `xml:"id,omitempty"`
	Matching *string       `xml:"matching,omitempty"`
	Type     *string       `xml:"type,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type protocolXml struct {
	Ikev1   *protocolIkev1Xml `xml:"ikev1,omitempty"`
	Ikev2   *protocolIkev2Xml `xml:"ikev2,omitempty"`
	Version *string           `xml:"version,omitempty"`
	Misc    []generic.Xml     `xml:",any"`
}
type protocolIkev1Xml struct {
	Dpd              *protocolIkev1DpdXml `xml:"dpd,omitempty"`
	ExchangeMode     *string              `xml:"exchange-mode,omitempty"`
	IkeCryptoProfile *string              `xml:"ike-crypto-profile,omitempty"`
	Misc             []generic.Xml        `xml:",any"`
}
type protocolIkev1DpdXml struct {
	Enable   *string       `xml:"enable,omitempty"`
	Interval *int64        `xml:"interval,omitempty"`
	Retry    *int64        `xml:"retry,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type protocolIkev2Xml struct {
	Dpd              *protocolIkev2DpdXml `xml:"dpd,omitempty"`
	IkeCryptoProfile *string              `xml:"ike-crypto-profile,omitempty"`
	RequireCookie    *string              `xml:"require-cookie,omitempty"`
	Misc             []generic.Xml        `xml:",any"`
}
type protocolIkev2DpdXml struct {
	Enable   *string       `xml:"enable,omitempty"`
	Interval *int64        `xml:"interval,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type protocolCommonXml struct {
	Fragmentation *protocolCommonFragmentationXml `xml:"fragmentation,omitempty"`
	NatTraversal  *protocolCommonNatTraversalXml  `xml:"nat-traversal,omitempty"`
	PassiveMode   *string                         `xml:"passive-mode,omitempty"`
	Misc          []generic.Xml                   `xml:",any"`
}
type protocolCommonFragmentationXml struct {
	Enable *string       `xml:"enable,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type protocolCommonNatTraversalXml struct {
	Enable            *string       `xml:"enable,omitempty"`
	KeepAliveInterval *int64        `xml:"keep-alive-interval,omitempty"`
	UdpChecksumEnable *string       `xml:"udp-checksum-enable,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Authentication != nil {
		var obj authenticationXml
		obj.MarshalFromObject(*s.Authentication)
		o.Authentication = &obj
	}
	o.Comment = s.Comment
	o.Disabled = util.YesNo(s.Disabled, nil)
	o.Ipv6 = util.YesNo(s.Ipv6, nil)
	if s.LocalAddress != nil {
		var obj localAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	if s.LocalId != nil {
		var obj localIdXml
		obj.MarshalFromObject(*s.LocalId)
		o.LocalId = &obj
	}
	if s.PeerAddress != nil {
		var obj peerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	if s.PeerId != nil {
		var obj peerIdXml
		obj.MarshalFromObject(*s.PeerId)
		o.PeerId = &obj
	}
	if s.Protocol != nil {
		var obj protocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	if s.ProtocolCommon != nil {
		var obj protocolCommonXml
		obj.MarshalFromObject(*s.ProtocolCommon)
		o.ProtocolCommon = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var authenticationVal *Authentication
	if o.Authentication != nil {
		obj, err := o.Authentication.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationVal = obj
	}
	var localAddressVal *LocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var localIdVal *LocalId
	if o.LocalId != nil {
		obj, err := o.LocalId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localIdVal = obj
	}
	var peerAddressVal *PeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var peerIdVal *PeerId
	if o.PeerId != nil {
		obj, err := o.PeerId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerIdVal = obj
	}
	var protocolVal *Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}
	var protocolCommonVal *ProtocolCommon
	if o.ProtocolCommon != nil {
		obj, err := o.ProtocolCommon.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolCommonVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Authentication: authenticationVal,
		Comment:        o.Comment,
		Disabled:       util.AsBool(o.Disabled, nil),
		Ipv6:           util.AsBool(o.Ipv6, nil),
		LocalAddress:   localAddressVal,
		LocalId:        localIdVal,
		PeerAddress:    peerAddressVal,
		PeerId:         peerIdVal,
		Protocol:       protocolVal,
		ProtocolCommon: protocolCommonVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *authenticationXml) MarshalFromObject(s Authentication) {
	if s.Certificate != nil {
		var obj authenticationCertificateXml
		obj.MarshalFromObject(*s.Certificate)
		o.Certificate = &obj
	}
	if s.PreSharedKey != nil {
		var obj authenticationPreSharedKeyXml
		obj.MarshalFromObject(*s.PreSharedKey)
		o.PreSharedKey = &obj
	}
	o.Misc = s.Misc
}

func (o authenticationXml) UnmarshalToObject() (*Authentication, error) {
	var certificateVal *AuthenticationCertificate
	if o.Certificate != nil {
		obj, err := o.Certificate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateVal = obj
	}
	var preSharedKeyVal *AuthenticationPreSharedKey
	if o.PreSharedKey != nil {
		obj, err := o.PreSharedKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		preSharedKeyVal = obj
	}

	result := &Authentication{
		Certificate:  certificateVal,
		PreSharedKey: preSharedKeyVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *authenticationCertificateXml) MarshalFromObject(s AuthenticationCertificate) {
	o.AllowIdPayloadMismatch = util.YesNo(s.AllowIdPayloadMismatch, nil)
	o.CertificateProfile = s.CertificateProfile
	if s.LocalCertificate != nil {
		var obj authenticationCertificateLocalCertificateXml
		obj.MarshalFromObject(*s.LocalCertificate)
		o.LocalCertificate = &obj
	}
	o.StrictValidationRevocation = util.YesNo(s.StrictValidationRevocation, nil)
	o.UseManagementAsSource = util.YesNo(s.UseManagementAsSource, nil)
	o.Misc = s.Misc
}

func (o authenticationCertificateXml) UnmarshalToObject() (*AuthenticationCertificate, error) {
	var localCertificateVal *AuthenticationCertificateLocalCertificate
	if o.LocalCertificate != nil {
		obj, err := o.LocalCertificate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localCertificateVal = obj
	}

	result := &AuthenticationCertificate{
		AllowIdPayloadMismatch:     util.AsBool(o.AllowIdPayloadMismatch, nil),
		CertificateProfile:         o.CertificateProfile,
		LocalCertificate:           localCertificateVal,
		StrictValidationRevocation: util.AsBool(o.StrictValidationRevocation, nil),
		UseManagementAsSource:      util.AsBool(o.UseManagementAsSource, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *authenticationCertificateLocalCertificateXml) MarshalFromObject(s AuthenticationCertificateLocalCertificate) {
	if s.HashAndUrl != nil {
		var obj authenticationCertificateLocalCertificateHashAndUrlXml
		obj.MarshalFromObject(*s.HashAndUrl)
		o.HashAndUrl = &obj
	}
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o authenticationCertificateLocalCertificateXml) UnmarshalToObject() (*AuthenticationCertificateLocalCertificate, error) {
	var hashAndUrlVal *AuthenticationCertificateLocalCertificateHashAndUrl
	if o.HashAndUrl != nil {
		obj, err := o.HashAndUrl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hashAndUrlVal = obj
	}

	result := &AuthenticationCertificateLocalCertificate{
		HashAndUrl: hashAndUrlVal,
		Name:       o.Name,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *authenticationCertificateLocalCertificateHashAndUrlXml) MarshalFromObject(s AuthenticationCertificateLocalCertificateHashAndUrl) {
	o.BaseUrl = s.BaseUrl
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
}

func (o authenticationCertificateLocalCertificateHashAndUrlXml) UnmarshalToObject() (*AuthenticationCertificateLocalCertificateHashAndUrl, error) {

	result := &AuthenticationCertificateLocalCertificateHashAndUrl{
		BaseUrl: o.BaseUrl,
		Enable:  util.AsBool(o.Enable, nil),
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *authenticationPreSharedKeyXml) MarshalFromObject(s AuthenticationPreSharedKey) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o authenticationPreSharedKeyXml) UnmarshalToObject() (*AuthenticationPreSharedKey, error) {

	result := &AuthenticationPreSharedKey{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *localAddressXml) MarshalFromObject(s LocalAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o localAddressXml) UnmarshalToObject() (*LocalAddress, error) {

	result := &LocalAddress{
		Interface:  o.Interface,
		FloatingIp: o.FloatingIp,
		Ip:         o.Ip,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *localIdXml) MarshalFromObject(s LocalId) {
	o.Id = s.Id
	o.Type = s.Type
	o.Misc = s.Misc
}

func (o localIdXml) UnmarshalToObject() (*LocalId, error) {

	result := &LocalId{
		Id:   o.Id,
		Type: o.Type,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *peerAddressXml) MarshalFromObject(s PeerAddress) {
	if s.Dynamic != nil {
		var obj peerAddressDynamicXml
		obj.MarshalFromObject(*s.Dynamic)
		o.Dynamic = &obj
	}
	o.Fqdn = s.Fqdn
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o peerAddressXml) UnmarshalToObject() (*PeerAddress, error) {
	var dynamicVal *PeerAddressDynamic
	if o.Dynamic != nil {
		obj, err := o.Dynamic.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dynamicVal = obj
	}

	result := &PeerAddress{
		Dynamic: dynamicVal,
		Fqdn:    o.Fqdn,
		Ip:      o.Ip,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *peerAddressDynamicXml) MarshalFromObject(s PeerAddressDynamic) {
	o.Misc = s.Misc
}

func (o peerAddressDynamicXml) UnmarshalToObject() (*PeerAddressDynamic, error) {

	result := &PeerAddressDynamic{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *peerIdXml) MarshalFromObject(s PeerId) {
	o.Id = s.Id
	o.Matching = s.Matching
	o.Type = s.Type
	o.Misc = s.Misc
}

func (o peerIdXml) UnmarshalToObject() (*PeerId, error) {

	result := &PeerId{
		Id:       o.Id,
		Matching: o.Matching,
		Type:     o.Type,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolXml) MarshalFromObject(s Protocol) {
	if s.Ikev1 != nil {
		var obj protocolIkev1Xml
		obj.MarshalFromObject(*s.Ikev1)
		o.Ikev1 = &obj
	}
	if s.Ikev2 != nil {
		var obj protocolIkev2Xml
		obj.MarshalFromObject(*s.Ikev2)
		o.Ikev2 = &obj
	}
	o.Version = s.Version
	o.Misc = s.Misc
}

func (o protocolXml) UnmarshalToObject() (*Protocol, error) {
	var ikev1Val *ProtocolIkev1
	if o.Ikev1 != nil {
		obj, err := o.Ikev1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ikev1Val = obj
	}
	var ikev2Val *ProtocolIkev2
	if o.Ikev2 != nil {
		obj, err := o.Ikev2.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ikev2Val = obj
	}

	result := &Protocol{
		Ikev1:   ikev1Val,
		Ikev2:   ikev2Val,
		Version: o.Version,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *protocolIkev1Xml) MarshalFromObject(s ProtocolIkev1) {
	if s.Dpd != nil {
		var obj protocolIkev1DpdXml
		obj.MarshalFromObject(*s.Dpd)
		o.Dpd = &obj
	}
	o.ExchangeMode = s.ExchangeMode
	o.IkeCryptoProfile = s.IkeCryptoProfile
	o.Misc = s.Misc
}

func (o protocolIkev1Xml) UnmarshalToObject() (*ProtocolIkev1, error) {
	var dpdVal *ProtocolIkev1Dpd
	if o.Dpd != nil {
		obj, err := o.Dpd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dpdVal = obj
	}

	result := &ProtocolIkev1{
		Dpd:              dpdVal,
		ExchangeMode:     o.ExchangeMode,
		IkeCryptoProfile: o.IkeCryptoProfile,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolIkev1DpdXml) MarshalFromObject(s ProtocolIkev1Dpd) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Interval = s.Interval
	o.Retry = s.Retry
	o.Misc = s.Misc
}

func (o protocolIkev1DpdXml) UnmarshalToObject() (*ProtocolIkev1Dpd, error) {

	result := &ProtocolIkev1Dpd{
		Enable:   util.AsBool(o.Enable, nil),
		Interval: o.Interval,
		Retry:    o.Retry,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolIkev2Xml) MarshalFromObject(s ProtocolIkev2) {
	if s.Dpd != nil {
		var obj protocolIkev2DpdXml
		obj.MarshalFromObject(*s.Dpd)
		o.Dpd = &obj
	}
	o.IkeCryptoProfile = s.IkeCryptoProfile
	o.RequireCookie = util.YesNo(s.RequireCookie, nil)
	o.Misc = s.Misc
}

func (o protocolIkev2Xml) UnmarshalToObject() (*ProtocolIkev2, error) {
	var dpdVal *ProtocolIkev2Dpd
	if o.Dpd != nil {
		obj, err := o.Dpd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dpdVal = obj
	}

	result := &ProtocolIkev2{
		Dpd:              dpdVal,
		IkeCryptoProfile: o.IkeCryptoProfile,
		RequireCookie:    util.AsBool(o.RequireCookie, nil),
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *protocolIkev2DpdXml) MarshalFromObject(s ProtocolIkev2Dpd) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Interval = s.Interval
	o.Misc = s.Misc
}

func (o protocolIkev2DpdXml) UnmarshalToObject() (*ProtocolIkev2Dpd, error) {

	result := &ProtocolIkev2Dpd{
		Enable:   util.AsBool(o.Enable, nil),
		Interval: o.Interval,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *protocolCommonXml) MarshalFromObject(s ProtocolCommon) {
	if s.Fragmentation != nil {
		var obj protocolCommonFragmentationXml
		obj.MarshalFromObject(*s.Fragmentation)
		o.Fragmentation = &obj
	}
	if s.NatTraversal != nil {
		var obj protocolCommonNatTraversalXml
		obj.MarshalFromObject(*s.NatTraversal)
		o.NatTraversal = &obj
	}
	o.PassiveMode = util.YesNo(s.PassiveMode, nil)
	o.Misc = s.Misc
}

func (o protocolCommonXml) UnmarshalToObject() (*ProtocolCommon, error) {
	var fragmentationVal *ProtocolCommonFragmentation
	if o.Fragmentation != nil {
		obj, err := o.Fragmentation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fragmentationVal = obj
	}
	var natTraversalVal *ProtocolCommonNatTraversal
	if o.NatTraversal != nil {
		obj, err := o.NatTraversal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		natTraversalVal = obj
	}

	result := &ProtocolCommon{
		Fragmentation: fragmentationVal,
		NatTraversal:  natTraversalVal,
		PassiveMode:   util.AsBool(o.PassiveMode, nil),
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *protocolCommonFragmentationXml) MarshalFromObject(s ProtocolCommonFragmentation) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
}

func (o protocolCommonFragmentationXml) UnmarshalToObject() (*ProtocolCommonFragmentation, error) {

	result := &ProtocolCommonFragmentation{
		Enable: util.AsBool(o.Enable, nil),
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *protocolCommonNatTraversalXml) MarshalFromObject(s ProtocolCommonNatTraversal) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.KeepAliveInterval = s.KeepAliveInterval
	o.UdpChecksumEnable = util.YesNo(s.UdpChecksumEnable, nil)
	o.Misc = s.Misc
}

func (o protocolCommonNatTraversalXml) UnmarshalToObject() (*ProtocolCommonNatTraversal, error) {

	result := &ProtocolCommonNatTraversal{
		Enable:            util.AsBool(o.Enable, nil),
		KeepAliveInterval: o.KeepAliveInterval,
		UdpChecksumEnable: util.AsBool(o.UdpChecksumEnable, nil),
		Misc:              o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authentication" || v == "Authentication" {
		return e.Authentication, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "local_address" || v == "LocalAddress" {
		return e.LocalAddress, nil
	}
	if v == "local_id" || v == "LocalId" {
		return e.LocalId, nil
	}
	if v == "peer_address" || v == "PeerAddress" {
		return e.PeerAddress, nil
	}
	if v == "peer_id" || v == "PeerId" {
		return e.PeerId, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "protocol_common" || v == "ProtocolCommon" {
		return e.ProtocolCommon, nil
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
	if !o.Authentication.matches(other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !util.BoolsMatch(o.Ipv6, other.Ipv6) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !o.LocalId.matches(other.LocalId) {
		return false
	}
	if !o.PeerAddress.matches(other.PeerAddress) {
		return false
	}
	if !o.PeerId.matches(other.PeerId) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}
	if !o.ProtocolCommon.matches(other.ProtocolCommon) {
		return false
	}

	return true
}

func (o *Authentication) matches(other *Authentication) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Certificate.matches(other.Certificate) {
		return false
	}
	if !o.PreSharedKey.matches(other.PreSharedKey) {
		return false
	}

	return true
}

func (o *AuthenticationCertificate) matches(other *AuthenticationCertificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowIdPayloadMismatch, other.AllowIdPayloadMismatch) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !o.LocalCertificate.matches(other.LocalCertificate) {
		return false
	}
	if !util.BoolsMatch(o.StrictValidationRevocation, other.StrictValidationRevocation) {
		return false
	}
	if !util.BoolsMatch(o.UseManagementAsSource, other.UseManagementAsSource) {
		return false
	}

	return true
}

func (o *AuthenticationCertificateLocalCertificate) matches(other *AuthenticationCertificateLocalCertificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.HashAndUrl.matches(other.HashAndUrl) {
		return false
	}
	if !util.StringsMatch(o.Name, other.Name) {
		return false
	}

	return true
}

func (o *AuthenticationCertificateLocalCertificateHashAndUrl) matches(other *AuthenticationCertificateLocalCertificateHashAndUrl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.BaseUrl, other.BaseUrl) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}

	return true
}

func (o *AuthenticationPreSharedKey) matches(other *AuthenticationPreSharedKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *LocalAddress) matches(other *LocalAddress) bool {
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

func (o *LocalId) matches(other *LocalId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Id, other.Id) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *PeerAddress) matches(other *PeerAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dynamic.matches(other.Dynamic) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *PeerAddressDynamic) matches(other *PeerAddressDynamic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *PeerId) matches(other *PeerId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Id, other.Id) {
		return false
	}
	if !util.StringsMatch(o.Matching, other.Matching) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *Protocol) matches(other *Protocol) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Ikev1.matches(other.Ikev1) {
		return false
	}
	if !o.Ikev2.matches(other.Ikev2) {
		return false
	}
	if !util.StringsMatch(o.Version, other.Version) {
		return false
	}

	return true
}

func (o *ProtocolIkev1) matches(other *ProtocolIkev1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dpd.matches(other.Dpd) {
		return false
	}
	if !util.StringsMatch(o.ExchangeMode, other.ExchangeMode) {
		return false
	}
	if !util.StringsMatch(o.IkeCryptoProfile, other.IkeCryptoProfile) {
		return false
	}

	return true
}

func (o *ProtocolIkev1Dpd) matches(other *ProtocolIkev1Dpd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Retry, other.Retry) {
		return false
	}

	return true
}

func (o *ProtocolIkev2) matches(other *ProtocolIkev2) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Dpd.matches(other.Dpd) {
		return false
	}
	if !util.StringsMatch(o.IkeCryptoProfile, other.IkeCryptoProfile) {
		return false
	}
	if !util.BoolsMatch(o.RequireCookie, other.RequireCookie) {
		return false
	}

	return true
}

func (o *ProtocolIkev2Dpd) matches(other *ProtocolIkev2Dpd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}

	return true
}

func (o *ProtocolCommon) matches(other *ProtocolCommon) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Fragmentation.matches(other.Fragmentation) {
		return false
	}
	if !o.NatTraversal.matches(other.NatTraversal) {
		return false
	}
	if !util.BoolsMatch(o.PassiveMode, other.PassiveMode) {
		return false
	}

	return true
}

func (o *ProtocolCommonFragmentation) matches(other *ProtocolCommonFragmentation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}

	return true
}

func (o *ProtocolCommonNatTraversal) matches(other *ProtocolCommonNatTraversal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.Ints64Match(o.KeepAliveInterval, other.KeepAliveInterval) {
		return false
	}
	if !util.BoolsMatch(o.UdpChecksumEnable, other.UdpChecksumEnable) {
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
