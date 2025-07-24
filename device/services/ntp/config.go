package ntp

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	NtpServers     *NtpServers
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServers struct {
	PrimaryNtpServer   *NtpServersPrimaryNtpServer
	SecondaryNtpServer *NtpServersSecondaryNtpServer
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type NtpServersPrimaryNtpServer struct {
	AuthenticationType *NtpServersPrimaryNtpServerAuthenticationType
	NtpServerAddress   *string
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationType struct {
	Autokey        *NtpServersPrimaryNtpServerAuthenticationTypeAutokey
	None           *NtpServersPrimaryNtpServerAuthenticationTypeNone
	SymmetricKey   *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeAutokey struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeNone struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey struct {
	Algorithm      *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	KeyId          *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm struct {
	Md5            *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	Sha1           *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 struct {
	AuthenticationKey *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1 struct {
	AuthenticationKey *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type NtpServersSecondaryNtpServer struct {
	AuthenticationType *NtpServersSecondaryNtpServerAuthenticationType
	NtpServerAddress   *string
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationType struct {
	Autokey        *NtpServersSecondaryNtpServerAuthenticationTypeAutokey
	None           *NtpServersSecondaryNtpServerAuthenticationTypeNone
	SymmetricKey   *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeAutokey struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeNone struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey struct {
	Algorithm      *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	KeyId          *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm struct {
	Md5            *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	Sha1           *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 struct {
	AuthenticationKey *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1 struct {
	AuthenticationKey *string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}

type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}

func (o *configXmlContainer) Normalize() ([]*Config, error) {
	entries := make([]*Config, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyConfig(source *Config) (any, error) {
	var obj configXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type configXml struct {
	XMLName        xml.Name       `xml:"system"`
	NtpServers     *ntpServersXml `xml:"ntp-servers,omitempty"`
	Misc           []generic.Xml  `xml:",any"`
	MiscAttributes []xml.Attr     `xml:",any,attr"`
}
type ntpServersXml struct {
	PrimaryNtpServer   *ntpServersPrimaryNtpServerXml   `xml:"primary-ntp-server,omitempty"`
	SecondaryNtpServer *ntpServersSecondaryNtpServerXml `xml:"secondary-ntp-server,omitempty"`
	Misc               []generic.Xml                    `xml:",any"`
	MiscAttributes     []xml.Attr                       `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerXml struct {
	AuthenticationType *ntpServersPrimaryNtpServerAuthenticationTypeXml `xml:"authentication-type,omitempty"`
	NtpServerAddress   *string                                          `xml:"ntp-server-address,omitempty"`
	Misc               []generic.Xml                                    `xml:",any"`
	MiscAttributes     []xml.Attr                                       `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeXml struct {
	Autokey        *ntpServersPrimaryNtpServerAuthenticationTypeAutokeyXml      `xml:"autokey,omitempty"`
	None           *ntpServersPrimaryNtpServerAuthenticationTypeNoneXml         `xml:"none,omitempty"`
	SymmetricKey   *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`
	Misc           []generic.Xml                                                `xml:",any"`
	MiscAttributes []xml.Attr                                                   `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeAutokeyXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeNoneXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	Algorithm      *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml `xml:"algorithm,omitempty"`
	KeyId          *int64                                                                `xml:"key-id,omitempty"`
	Misc           []generic.Xml                                                         `xml:",any"`
	MiscAttributes []xml.Attr                                                            `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml struct {
	Md5            *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml  `xml:"md5,omitempty"`
	Sha1           *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml `xml:"sha1,omitempty"`
	Misc           []generic.Xml                                                             `xml:",any"`
	MiscAttributes []xml.Attr                                                                `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml struct {
	AuthenticationKey *string       `xml:"authentication-key,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml struct {
	AuthenticationKey *string       `xml:"authentication-key,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerXml struct {
	AuthenticationType *ntpServersSecondaryNtpServerAuthenticationTypeXml `xml:"authentication-type,omitempty"`
	NtpServerAddress   *string                                            `xml:"ntp-server-address,omitempty"`
	Misc               []generic.Xml                                      `xml:",any"`
	MiscAttributes     []xml.Attr                                         `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeXml struct {
	Autokey        *ntpServersSecondaryNtpServerAuthenticationTypeAutokeyXml      `xml:"autokey,omitempty"`
	None           *ntpServersSecondaryNtpServerAuthenticationTypeNoneXml         `xml:"none,omitempty"`
	SymmetricKey   *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`
	Misc           []generic.Xml                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                     `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeAutokeyXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeNoneXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	Algorithm      *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml `xml:"algorithm,omitempty"`
	KeyId          *int64                                                                  `xml:"key-id,omitempty"`
	Misc           []generic.Xml                                                           `xml:",any"`
	MiscAttributes []xml.Attr                                                              `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml struct {
	Md5            *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml  `xml:"md5,omitempty"`
	Sha1           *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml `xml:"sha1,omitempty"`
	Misc           []generic.Xml                                                               `xml:",any"`
	MiscAttributes []xml.Attr                                                                  `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml struct {
	AuthenticationKey *string       `xml:"authentication-key,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml struct {
	AuthenticationKey *string       `xml:"authentication-key,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}

func (o *configXml) MarshalFromObject(s Config) {
	if s.NtpServers != nil {
		var obj ntpServersXml
		obj.MarshalFromObject(*s.NtpServers)
		o.NtpServers = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var ntpServersVal *NtpServers
	if o.NtpServers != nil {
		obj, err := o.NtpServers.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ntpServersVal = obj
	}

	result := &Config{
		NtpServers:     ntpServersVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersXml) MarshalFromObject(s NtpServers) {
	if s.PrimaryNtpServer != nil {
		var obj ntpServersPrimaryNtpServerXml
		obj.MarshalFromObject(*s.PrimaryNtpServer)
		o.PrimaryNtpServer = &obj
	}
	if s.SecondaryNtpServer != nil {
		var obj ntpServersSecondaryNtpServerXml
		obj.MarshalFromObject(*s.SecondaryNtpServer)
		o.SecondaryNtpServer = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersXml) UnmarshalToObject() (*NtpServers, error) {
	var primaryNtpServerVal *NtpServersPrimaryNtpServer
	if o.PrimaryNtpServer != nil {
		obj, err := o.PrimaryNtpServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		primaryNtpServerVal = obj
	}
	var secondaryNtpServerVal *NtpServersSecondaryNtpServer
	if o.SecondaryNtpServer != nil {
		obj, err := o.SecondaryNtpServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		secondaryNtpServerVal = obj
	}

	result := &NtpServers{
		PrimaryNtpServer:   primaryNtpServerVal,
		SecondaryNtpServer: secondaryNtpServerVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerXml) MarshalFromObject(s NtpServersPrimaryNtpServer) {
	if s.AuthenticationType != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeXml
		obj.MarshalFromObject(*s.AuthenticationType)
		o.AuthenticationType = &obj
	}
	o.NtpServerAddress = s.NtpServerAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerXml) UnmarshalToObject() (*NtpServersPrimaryNtpServer, error) {
	var authenticationTypeVal *NtpServersPrimaryNtpServerAuthenticationType
	if o.AuthenticationType != nil {
		obj, err := o.AuthenticationType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationTypeVal = obj
	}

	result := &NtpServersPrimaryNtpServer{
		AuthenticationType: authenticationTypeVal,
		NtpServerAddress:   o.NtpServerAddress,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeXml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationType) {
	if s.Autokey != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeAutokeyXml
		obj.MarshalFromObject(*s.Autokey)
		o.Autokey = &obj
	}
	if s.None != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.SymmetricKey != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml
		obj.MarshalFromObject(*s.SymmetricKey)
		o.SymmetricKey = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeXml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationType, error) {
	var autokeyVal *NtpServersPrimaryNtpServerAuthenticationTypeAutokey
	if o.Autokey != nil {
		obj, err := o.Autokey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		autokeyVal = obj
	}
	var noneVal *NtpServersPrimaryNtpServerAuthenticationTypeNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var symmetricKeyVal *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
	if o.SymmetricKey != nil {
		obj, err := o.SymmetricKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		symmetricKeyVal = obj
	}

	result := &NtpServersPrimaryNtpServerAuthenticationType{
		Autokey:        autokeyVal,
		None:           noneVal,
		SymmetricKey:   symmetricKeyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeAutokeyXml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeAutokey) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeAutokeyXml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeAutokey, error) {

	result := &NtpServersPrimaryNtpServerAuthenticationTypeAutokey{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeNoneXml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeNone) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeNoneXml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeNone, error) {

	result := &NtpServersPrimaryNtpServerAuthenticationTypeNone{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) {
	if s.Algorithm != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml
		obj.MarshalFromObject(*s.Algorithm)
		o.Algorithm = &obj
	}
	o.KeyId = s.KeyId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, error) {
	var algorithmVal *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	if o.Algorithm != nil {
		obj, err := o.Algorithm.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		algorithmVal = obj
	}

	result := &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{
		Algorithm:      algorithmVal,
		KeyId:          o.KeyId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) {
	if s.Md5 != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, error) {
	var md5Val *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}

	result := &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{
		Md5:            md5Val,
		Sha1:           sha1Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) {
	o.AuthenticationKey = s.AuthenticationKey
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, error) {

	result := &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{
		AuthenticationKey: o.AuthenticationKey,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml) MarshalFromObject(s NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) {
	o.AuthenticationKey = s.AuthenticationKey
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml) UnmarshalToObject() (*NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1, error) {

	result := &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1{
		AuthenticationKey: o.AuthenticationKey,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerXml) MarshalFromObject(s NtpServersSecondaryNtpServer) {
	if s.AuthenticationType != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeXml
		obj.MarshalFromObject(*s.AuthenticationType)
		o.AuthenticationType = &obj
	}
	o.NtpServerAddress = s.NtpServerAddress
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerXml) UnmarshalToObject() (*NtpServersSecondaryNtpServer, error) {
	var authenticationTypeVal *NtpServersSecondaryNtpServerAuthenticationType
	if o.AuthenticationType != nil {
		obj, err := o.AuthenticationType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationTypeVal = obj
	}

	result := &NtpServersSecondaryNtpServer{
		AuthenticationType: authenticationTypeVal,
		NtpServerAddress:   o.NtpServerAddress,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeXml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationType) {
	if s.Autokey != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeAutokeyXml
		obj.MarshalFromObject(*s.Autokey)
		o.Autokey = &obj
	}
	if s.None != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.SymmetricKey != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml
		obj.MarshalFromObject(*s.SymmetricKey)
		o.SymmetricKey = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeXml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationType, error) {
	var autokeyVal *NtpServersSecondaryNtpServerAuthenticationTypeAutokey
	if o.Autokey != nil {
		obj, err := o.Autokey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		autokeyVal = obj
	}
	var noneVal *NtpServersSecondaryNtpServerAuthenticationTypeNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var symmetricKeyVal *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey
	if o.SymmetricKey != nil {
		obj, err := o.SymmetricKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		symmetricKeyVal = obj
	}

	result := &NtpServersSecondaryNtpServerAuthenticationType{
		Autokey:        autokeyVal,
		None:           noneVal,
		SymmetricKey:   symmetricKeyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeAutokeyXml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeAutokey) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeAutokeyXml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeAutokey, error) {

	result := &NtpServersSecondaryNtpServerAuthenticationTypeAutokey{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeNoneXml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeNone) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeNoneXml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeNone, error) {

	result := &NtpServersSecondaryNtpServerAuthenticationTypeNone{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey) {
	if s.Algorithm != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml
		obj.MarshalFromObject(*s.Algorithm)
		o.Algorithm = &obj
	}
	o.KeyId = s.KeyId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey, error) {
	var algorithmVal *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	if o.Algorithm != nil {
		obj, err := o.Algorithm.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		algorithmVal = obj
	}

	result := &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey{
		Algorithm:      algorithmVal,
		KeyId:          o.KeyId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) {
	if s.Md5 != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, error) {
	var md5Val *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}

	result := &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{
		Md5:            md5Val,
		Sha1:           sha1Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) {
	o.AuthenticationKey = s.AuthenticationKey
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, error) {

	result := &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{
		AuthenticationKey: o.AuthenticationKey,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml) MarshalFromObject(s NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) {
	o.AuthenticationKey = s.AuthenticationKey
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ntpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml) UnmarshalToObject() (*NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1, error) {

	result := &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1{
		AuthenticationKey: o.AuthenticationKey,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func SpecMatches(a, b *Config) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Config) matches(other *Config) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.NtpServers.matches(other.NtpServers) {
		return false
	}

	return true
}

func (o *NtpServers) matches(other *NtpServers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.PrimaryNtpServer.matches(other.PrimaryNtpServer) {
		return false
	}
	if !o.SecondaryNtpServer.matches(other.SecondaryNtpServer) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServer) matches(other *NtpServersPrimaryNtpServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AuthenticationType.matches(other.AuthenticationType) {
		return false
	}
	if !util.StringsMatch(o.NtpServerAddress, other.NtpServerAddress) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationType) matches(other *NtpServersPrimaryNtpServerAuthenticationType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Autokey.matches(other.Autokey) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.SymmetricKey.matches(other.SymmetricKey) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeAutokey) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeAutokey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeNone) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Algorithm.matches(other.Algorithm) {
		return false
	}
	if !util.Ints64Match(o.KeyId, other.KeyId) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationKey, other.AuthenticationKey) {
		return false
	}

	return true
}

func (o *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) matches(other *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationKey, other.AuthenticationKey) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServer) matches(other *NtpServersSecondaryNtpServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AuthenticationType.matches(other.AuthenticationType) {
		return false
	}
	if !util.StringsMatch(o.NtpServerAddress, other.NtpServerAddress) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationType) matches(other *NtpServersSecondaryNtpServerAuthenticationType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Autokey.matches(other.Autokey) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.SymmetricKey.matches(other.SymmetricKey) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeAutokey) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeAutokey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeNone) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Algorithm.matches(other.Algorithm) {
		return false
	}
	if !util.Ints64Match(o.KeyId, other.KeyId) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationKey, other.AuthenticationKey) {
		return false
	}

	return true
}

func (o *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) matches(other *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationKey, other.AuthenticationKey) {
		return false
	}

	return true
}
