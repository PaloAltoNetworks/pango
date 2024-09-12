package dns

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	DnsSetting      *DnsSetting
	FqdnRefreshTime *int64

	Misc map[string][]generic.Xml
}
type DnsSetting struct {
	Servers *DnsSettingServers
}
type DnsSettingServers struct {
	Primary   *string
	Secondary *string
}
type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}
type configXml struct {
	XMLName         xml.Name       `xml:"system"`
	DnsSetting      *DnsSettingXml `xml:"dns-setting,omitempty"`
	FqdnRefreshTime *int64         `xml:"fqdn-refresh-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DnsSettingXml struct {
	Servers *DnsSettingServersXml `xml:"servers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DnsSettingServersXml struct {
	Primary   *string `xml:"primary,omitempty"`
	Secondary *string `xml:"secondary,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func specifyConfig(o *Config) (any, error) {
	config := configXml{}
	var nestedDnsSetting *DnsSettingXml
	if o.DnsSetting != nil {
		nestedDnsSetting = &DnsSettingXml{}
		if _, ok := o.Misc["DnsSetting"]; ok {
			nestedDnsSetting.Misc = o.Misc["DnsSetting"]
		}
		if o.DnsSetting.Servers != nil {
			nestedDnsSetting.Servers = &DnsSettingServersXml{}
			if _, ok := o.Misc["DnsSettingServers"]; ok {
				nestedDnsSetting.Servers.Misc = o.Misc["DnsSettingServers"]
			}
			if o.DnsSetting.Servers.Primary != nil {
				nestedDnsSetting.Servers.Primary = o.DnsSetting.Servers.Primary
			}
			if o.DnsSetting.Servers.Secondary != nil {
				nestedDnsSetting.Servers.Secondary = o.DnsSetting.Servers.Secondary
			}
		}
	}
	config.DnsSetting = nestedDnsSetting

	config.FqdnRefreshTime = o.FqdnRefreshTime

	config.Misc = o.Misc["Config"]

	return config, nil
}
func (c *configXmlContainer) Normalize() ([]*Config, error) {
	configList := make([]*Config, 0, len(c.Answer))
	for _, o := range c.Answer {
		config := &Config{
			Misc: make(map[string][]generic.Xml),
		}
		var nestedDnsSetting *DnsSetting
		if o.DnsSetting != nil {
			nestedDnsSetting = &DnsSetting{}
			if o.DnsSetting.Misc != nil {
				config.Misc["DnsSetting"] = o.DnsSetting.Misc
			}
			if o.DnsSetting.Servers != nil {
				nestedDnsSetting.Servers = &DnsSettingServers{}
				if o.DnsSetting.Servers.Misc != nil {
					config.Misc["DnsSettingServers"] = o.DnsSetting.Servers.Misc
				}
				if o.DnsSetting.Servers.Secondary != nil {
					nestedDnsSetting.Servers.Secondary = o.DnsSetting.Servers.Secondary
				}
				if o.DnsSetting.Servers.Primary != nil {
					nestedDnsSetting.Servers.Primary = o.DnsSetting.Servers.Primary
				}
			}
		}
		config.DnsSetting = nestedDnsSetting

		config.FqdnRefreshTime = o.FqdnRefreshTime

		config.Misc["Config"] = o.Misc

		configList = append(configList, config)
	}

	return configList, nil
}

func SpecMatches(a, b *Config) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !matchDnsSetting(a.DnsSetting, b.DnsSetting) {
		return false
	}
	if !util.Ints64Match(a.FqdnRefreshTime, b.FqdnRefreshTime) {
		return false
	}

	return true
}

func matchDnsSettingServers(a *DnsSettingServers, b *DnsSettingServers) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Primary, b.Primary) {
		return false
	}
	if !util.StringsMatch(a.Secondary, b.Secondary) {
		return false
	}
	return true
}
func matchDnsSetting(a *DnsSetting, b *DnsSetting) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchDnsSettingServers(a.Servers, b.Servers) {
		return false
	}
	return true
}
