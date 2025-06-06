package general

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	Domain               *string
	GeoLocation          *GeoLocation
	Hostname             *string
	LoginBanner          *string
	SslTlsServiceProfile *string
	Timezone             *string
	Misc                 []generic.Xml
}
type GeoLocation struct {
	Latitude  *string
	Longitude *string
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
	XMLName              xml.Name        `xml:"system"`
	Domain               *string         `xml:"domain,omitempty"`
	GeoLocation          *geoLocationXml `xml:"geo-location,omitempty"`
	Hostname             *string         `xml:"hostname,omitempty"`
	LoginBanner          *string         `xml:"login-banner,omitempty"`
	SslTlsServiceProfile *string         `xml:"ssl-tls-service-profile,omitempty"`
	Timezone             *string         `xml:"timezone,omitempty"`
	Misc                 []generic.Xml   `xml:",any"`
}
type geoLocationXml struct {
	Latitude  *string       `xml:"latitude,omitempty"`
	Longitude *string       `xml:"longitude,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}

func (o *configXml) MarshalFromObject(s Config) {
	o.Domain = s.Domain
	if s.GeoLocation != nil {
		var obj geoLocationXml
		obj.MarshalFromObject(*s.GeoLocation)
		o.GeoLocation = &obj
	}
	o.Hostname = s.Hostname
	o.LoginBanner = s.LoginBanner
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Timezone = s.Timezone
	o.Misc = s.Misc
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var geoLocationVal *GeoLocation
	if o.GeoLocation != nil {
		obj, err := o.GeoLocation.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		geoLocationVal = obj
	}

	result := &Config{
		Domain:               o.Domain,
		GeoLocation:          geoLocationVal,
		Hostname:             o.Hostname,
		LoginBanner:          o.LoginBanner,
		SslTlsServiceProfile: o.SslTlsServiceProfile,
		Timezone:             o.Timezone,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *geoLocationXml) MarshalFromObject(s GeoLocation) {
	o.Latitude = s.Latitude
	o.Longitude = s.Longitude
	o.Misc = s.Misc
}

func (o geoLocationXml) UnmarshalToObject() (*GeoLocation, error) {

	result := &GeoLocation{
		Latitude:  o.Latitude,
		Longitude: o.Longitude,
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
	if !util.StringsMatch(o.Domain, other.Domain) {
		return false
	}
	if !o.GeoLocation.matches(other.GeoLocation) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}
	if !util.StringsMatch(o.LoginBanner, other.LoginBanner) {
		return false
	}
	if !util.StringsMatch(o.SslTlsServiceProfile, other.SslTlsServiceProfile) {
		return false
	}
	if !util.StringsMatch(o.Timezone, other.Timezone) {
		return false
	}

	return true
}

func (o *GeoLocation) matches(other *GeoLocation) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Latitude, other.Latitude) {
		return false
	}
	if !util.StringsMatch(o.Longitude, other.Longitude) {
		return false
	}

	return true
}
