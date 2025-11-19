package proxy

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	LcaasUseProxy       *bool
	SecureProxyPassword *string
	SecureProxyPort     *int64
	SecureProxyServer   *string
	SecureProxyUser     *string
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
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
	XMLName             xml.Name      `xml:"system"`
	LcaasUseProxy       *string       `xml:"lcaas-use-proxy,omitempty"`
	SecureProxyPassword *string       `xml:"secure-proxy-password,omitempty"`
	SecureProxyPort     *int64        `xml:"secure-proxy-port,omitempty"`
	SecureProxyServer   *string       `xml:"secure-proxy-server,omitempty"`
	SecureProxyUser     *string       `xml:"secure-proxy-user,omitempty"`
	Misc                []generic.Xml `xml:",any"`
	MiscAttributes      []xml.Attr    `xml:",any,attr"`
}

func (o *configXml) MarshalFromObject(s Config) {
	o.LcaasUseProxy = util.YesNo(s.LcaasUseProxy, nil)
	o.SecureProxyPassword = s.SecureProxyPassword
	o.SecureProxyPort = s.SecureProxyPort
	o.SecureProxyServer = s.SecureProxyServer
	o.SecureProxyUser = s.SecureProxyUser
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configXml) UnmarshalToObject() (*Config, error) {

	result := &Config{
		LcaasUseProxy:       util.AsBool(o.LcaasUseProxy, nil),
		SecureProxyPassword: o.SecureProxyPassword,
		SecureProxyPort:     o.SecureProxyPort,
		SecureProxyServer:   o.SecureProxyServer,
		SecureProxyUser:     o.SecureProxyUser,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
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
	if !util.BoolsMatch(o.LcaasUseProxy, other.LcaasUseProxy) {
		return false
	}
	if !util.StringsMatch(o.SecureProxyPassword, other.SecureProxyPassword) {
		return false
	}
	if !util.Ints64Match(o.SecureProxyPort, other.SecureProxyPort) {
		return false
	}
	if !util.StringsMatch(o.SecureProxyServer, other.SecureProxyServer) {
		return false
	}
	if !util.StringsMatch(o.SecureProxyUser, other.SecureProxyUser) {
		return false
	}

	return true
}
