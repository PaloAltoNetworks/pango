package radius

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
	suffix = []string{"server-profile", "radius", "$name"}
)

type Entry struct {
	Name           string
	Protocol       *Protocol
	Retries        *int64
	Server         []Server
	Timeout        *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Protocol struct {
	Chap           *ProtocolChap
	EapTtlsWithPap *ProtocolEapTtlsWithPap
	Pap            *ProtocolPap
	PeapMschapv2   *ProtocolPeapMschapv2
	PeapWithGtc    *ProtocolPeapWithGtc
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ProtocolChap struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ProtocolEapTtlsWithPap struct {
	AnonOuterId       *bool
	RadiusCertProfile *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type ProtocolPap struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ProtocolPeapMschapv2 struct {
	AllowPwdChange    *bool
	AnonOuterId       *bool
	RadiusCertProfile *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type ProtocolPeapWithGtc struct {
	AnonOuterId       *bool
	RadiusCertProfile *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type Server struct {
	Name           string
	IpAddress      *string
	Secret         *string
	Port           *int64
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
	XMLName        xml.Name            `xml:"entry"`
	Name           string              `xml:"name,attr"`
	Protocol       *protocolXml        `xml:"protocol,omitempty"`
	Retries        *int64              `xml:"retries,omitempty"`
	Server         *serverContainerXml `xml:"server,omitempty"`
	Timeout        *int64              `xml:"timeout,omitempty"`
	Misc           []generic.Xml       `xml:",any"`
	MiscAttributes []xml.Attr          `xml:",any,attr"`
}
type protocolXml struct {
	Chap           *protocolChapXml           `xml:"CHAP,omitempty"`
	EapTtlsWithPap *protocolEapTtlsWithPapXml `xml:"EAP-TTLS-with-PAP,omitempty"`
	Pap            *protocolPapXml            `xml:"PAP,omitempty"`
	PeapMschapv2   *protocolPeapMschapv2Xml   `xml:"PEAP-MSCHAPv2,omitempty"`
	PeapWithGtc    *protocolPeapWithGtcXml    `xml:"PEAP-with-GTC,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type protocolChapXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type protocolEapTtlsWithPapXml struct {
	AnonOuterId       *string       `xml:"anon-outer-id,omitempty"`
	RadiusCertProfile *string       `xml:"radius-cert-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type protocolPapXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type protocolPeapMschapv2Xml struct {
	AllowPwdChange    *string       `xml:"allow-pwd-change,omitempty"`
	AnonOuterId       *string       `xml:"anon-outer-id,omitempty"`
	RadiusCertProfile *string       `xml:"radius-cert-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type protocolPeapWithGtcXml struct {
	AnonOuterId       *string       `xml:"anon-outer-id,omitempty"`
	RadiusCertProfile *string       `xml:"radius-cert-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type serverContainerXml struct {
	Entries []serverXml `xml:"entry"`
}
type serverXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Secret         *string       `xml:"secret,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Protocol != nil {
		var obj protocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Retries = s.Retries
	if s.Server != nil {
		var objs []serverXml
		for _, elt := range s.Server {
			var obj serverXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Server = &serverContainerXml{Entries: objs}
	}
	o.Timeout = s.Timeout
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var protocolVal *Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}
	var serverVal []Server
	if o.Server != nil {
		for _, elt := range o.Server.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			serverVal = append(serverVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		Protocol:       protocolVal,
		Retries:        o.Retries,
		Server:         serverVal,
		Timeout:        o.Timeout,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolXml) MarshalFromObject(s Protocol) {
	if s.Chap != nil {
		var obj protocolChapXml
		obj.MarshalFromObject(*s.Chap)
		o.Chap = &obj
	}
	if s.EapTtlsWithPap != nil {
		var obj protocolEapTtlsWithPapXml
		obj.MarshalFromObject(*s.EapTtlsWithPap)
		o.EapTtlsWithPap = &obj
	}
	if s.Pap != nil {
		var obj protocolPapXml
		obj.MarshalFromObject(*s.Pap)
		o.Pap = &obj
	}
	if s.PeapMschapv2 != nil {
		var obj protocolPeapMschapv2Xml
		obj.MarshalFromObject(*s.PeapMschapv2)
		o.PeapMschapv2 = &obj
	}
	if s.PeapWithGtc != nil {
		var obj protocolPeapWithGtcXml
		obj.MarshalFromObject(*s.PeapWithGtc)
		o.PeapWithGtc = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolXml) UnmarshalToObject() (*Protocol, error) {
	var chapVal *ProtocolChap
	if o.Chap != nil {
		obj, err := o.Chap.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		chapVal = obj
	}
	var eapTtlsWithPapVal *ProtocolEapTtlsWithPap
	if o.EapTtlsWithPap != nil {
		obj, err := o.EapTtlsWithPap.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		eapTtlsWithPapVal = obj
	}
	var papVal *ProtocolPap
	if o.Pap != nil {
		obj, err := o.Pap.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		papVal = obj
	}
	var peapMschapv2Val *ProtocolPeapMschapv2
	if o.PeapMschapv2 != nil {
		obj, err := o.PeapMschapv2.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peapMschapv2Val = obj
	}
	var peapWithGtcVal *ProtocolPeapWithGtc
	if o.PeapWithGtc != nil {
		obj, err := o.PeapWithGtc.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peapWithGtcVal = obj
	}

	result := &Protocol{
		Chap:           chapVal,
		EapTtlsWithPap: eapTtlsWithPapVal,
		Pap:            papVal,
		PeapMschapv2:   peapMschapv2Val,
		PeapWithGtc:    peapWithGtcVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolChapXml) MarshalFromObject(s ProtocolChap) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolChapXml) UnmarshalToObject() (*ProtocolChap, error) {

	result := &ProtocolChap{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolEapTtlsWithPapXml) MarshalFromObject(s ProtocolEapTtlsWithPap) {
	o.AnonOuterId = util.YesNo(s.AnonOuterId, nil)
	o.RadiusCertProfile = s.RadiusCertProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolEapTtlsWithPapXml) UnmarshalToObject() (*ProtocolEapTtlsWithPap, error) {

	result := &ProtocolEapTtlsWithPap{
		AnonOuterId:       util.AsBool(o.AnonOuterId, nil),
		RadiusCertProfile: o.RadiusCertProfile,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolPapXml) MarshalFromObject(s ProtocolPap) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolPapXml) UnmarshalToObject() (*ProtocolPap, error) {

	result := &ProtocolPap{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolPeapMschapv2Xml) MarshalFromObject(s ProtocolPeapMschapv2) {
	o.AllowPwdChange = util.YesNo(s.AllowPwdChange, nil)
	o.AnonOuterId = util.YesNo(s.AnonOuterId, nil)
	o.RadiusCertProfile = s.RadiusCertProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolPeapMschapv2Xml) UnmarshalToObject() (*ProtocolPeapMschapv2, error) {

	result := &ProtocolPeapMschapv2{
		AllowPwdChange:    util.AsBool(o.AllowPwdChange, nil),
		AnonOuterId:       util.AsBool(o.AnonOuterId, nil),
		RadiusCertProfile: o.RadiusCertProfile,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *protocolPeapWithGtcXml) MarshalFromObject(s ProtocolPeapWithGtc) {
	o.AnonOuterId = util.YesNo(s.AnonOuterId, nil)
	o.RadiusCertProfile = s.RadiusCertProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o protocolPeapWithGtcXml) UnmarshalToObject() (*ProtocolPeapWithGtc, error) {

	result := &ProtocolPeapWithGtc{
		AnonOuterId:       util.AsBool(o.AnonOuterId, nil),
		RadiusCertProfile: o.RadiusCertProfile,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *serverXml) MarshalFromObject(s Server) {
	o.Name = s.Name
	o.IpAddress = s.IpAddress
	o.Secret = s.Secret
	o.Port = s.Port
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o serverXml) UnmarshalToObject() (*Server, error) {

	result := &Server{
		Name:           o.Name,
		IpAddress:      o.IpAddress,
		Secret:         o.Secret,
		Port:           o.Port,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "retries" || v == "Retries" {
		return e.Retries, nil
	}
	if v == "server" || v == "Server" {
		return e.Server, nil
	}
	if v == "server|LENGTH" || v == "Server|LENGTH" {
		return int64(len(e.Server)), nil
	}
	if v == "timeout" || v == "Timeout" {
		return e.Timeout, nil
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
	if !o.Protocol.matches(other.Protocol) {
		return false
	}
	if !util.Ints64Match(o.Retries, other.Retries) {
		return false
	}
	if len(o.Server) != len(other.Server) {
		return false
	}
	for idx := range o.Server {
		if !o.Server[idx].matches(&other.Server[idx]) {
			return false
		}
	}
	if !util.Ints64Match(o.Timeout, other.Timeout) {
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
	if !o.Chap.matches(other.Chap) {
		return false
	}
	if !o.EapTtlsWithPap.matches(other.EapTtlsWithPap) {
		return false
	}
	if !o.Pap.matches(other.Pap) {
		return false
	}
	if !o.PeapMschapv2.matches(other.PeapMschapv2) {
		return false
	}
	if !o.PeapWithGtc.matches(other.PeapWithGtc) {
		return false
	}

	return true
}

func (o *ProtocolChap) matches(other *ProtocolChap) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolEapTtlsWithPap) matches(other *ProtocolEapTtlsWithPap) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AnonOuterId, other.AnonOuterId) {
		return false
	}
	if !util.StringsMatch(o.RadiusCertProfile, other.RadiusCertProfile) {
		return false
	}

	return true
}

func (o *ProtocolPap) matches(other *ProtocolPap) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ProtocolPeapMschapv2) matches(other *ProtocolPeapMschapv2) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AllowPwdChange, other.AllowPwdChange) {
		return false
	}
	if !util.BoolsMatch(o.AnonOuterId, other.AnonOuterId) {
		return false
	}
	if !util.StringsMatch(o.RadiusCertProfile, other.RadiusCertProfile) {
		return false
	}

	return true
}

func (o *ProtocolPeapWithGtc) matches(other *ProtocolPeapWithGtc) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AnonOuterId, other.AnonOuterId) {
		return false
	}
	if !util.StringsMatch(o.RadiusCertProfile, other.RadiusCertProfile) {
		return false
	}

	return true
}

func (o *Server) matches(other *Server) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}
	if !util.StringsMatch(o.Secret, other.Secret) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
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
