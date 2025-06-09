package dns

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

type Config struct {
	DnsSetting      *DnsSetting
	FqdnRefreshTime *int64
	Misc            []generic.Xml
}
type DnsSetting struct {
	Servers *DnsSettingServers
	Misc    []generic.Xml
}
type DnsSettingServers struct {
	Primary   *string
	Secondary *string
	Misc      []generic.Xml
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
	XMLName         xml.Name       `xml:"system"`
	DnsSetting      *dnsSettingXml `xml:"dns-setting,omitempty"`
	FqdnRefreshTime *int64         `xml:"fqdn-refresh-time,omitempty"`
	Misc            []generic.Xml  `xml:",any"`
}
type dnsSettingXml struct {
	Servers *dnsSettingServersXml `xml:"servers,omitempty"`
	Misc    []generic.Xml         `xml:",any"`
}
type dnsSettingServersXml struct {
	Primary   *string       `xml:"primary,omitempty"`
	Secondary *string       `xml:"secondary,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}

func (o *configXml) MarshalFromObject(s Config) {
	if s.DnsSetting != nil {
		var obj dnsSettingXml
		obj.MarshalFromObject(*s.DnsSetting)
		o.DnsSetting = &obj
	}
	o.FqdnRefreshTime = s.FqdnRefreshTime
	o.Misc = s.Misc
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var dnsSettingVal *DnsSetting
	if o.DnsSetting != nil {
		obj, err := o.DnsSetting.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dnsSettingVal = obj
	}

	result := &Config{
		DnsSetting:      dnsSettingVal,
		FqdnRefreshTime: o.FqdnRefreshTime,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *dnsSettingXml) MarshalFromObject(s DnsSetting) {
	if s.Servers != nil {
		var obj dnsSettingServersXml
		obj.MarshalFromObject(*s.Servers)
		o.Servers = &obj
	}
	o.Misc = s.Misc
}

func (o dnsSettingXml) UnmarshalToObject() (*DnsSetting, error) {
	var serversVal *DnsSettingServers
	if o.Servers != nil {
		obj, err := o.Servers.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serversVal = obj
	}

	result := &DnsSetting{
		Servers: serversVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *dnsSettingServersXml) MarshalFromObject(s DnsSettingServers) {
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
}

func (o dnsSettingServersXml) UnmarshalToObject() (*DnsSettingServers, error) {

	result := &DnsSettingServers{
		Primary:   o.Primary,
		Secondary: o.Secondary,
		Misc:      o.Misc,
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
	if !o.DnsSetting.matches(other.DnsSetting) {
		return false
	}
	if !util.Ints64Match(o.FqdnRefreshTime, other.FqdnRefreshTime) {
		return false
	}

	return true
}

func (o *DnsSetting) matches(other *DnsSetting) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Servers.matches(other.Servers) {
		return false
	}

	return true
}

func (o *DnsSettingServers) matches(other *DnsSettingServers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}
