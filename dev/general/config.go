package general

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Config is a normalized, version independent representation of a device's
// general settings.
type Config struct {
	Hostname              string
	IpAddress             string
	Netmask               string
	Gateway               string
	Timezone              string
	Domain                string
	UpdateServer          string
	VerifyUpdateServer    bool
	LoginBanner           string
	PanoramaPrimary       string
	PanoramaSecondary     string
	ProxyServer           string
	ProxyPort             int
	ProxyUser             string
	ProxyPassword         string
	DnsPrimary            string
	DnsSecondary          string
	NtpPrimaryAddress     string
	NtpPrimaryAuthType    string
	NtpPrimaryKeyId       int
	NtpPrimaryAlgorithm   string
	NtpPrimaryAuthKey     string
	NtpSecondaryAddress   string
	NtpSecondaryAuthType  string
	NtpSecondaryKeyId     int
	NtpSecondaryAlgorithm string
	NtpSecondaryAuthKey   string

	raw map[string]string
}

// Defaults sets params with uninitialized values to their GUI default setting.
//
// The defaults are as follows:
//      * UpdateServer: updates.paloaltonetworks.com
func (o *Config) Defaults() {
	if o.UpdateServer == "" {
		o.UpdateServer = "updates.paloaltonetworks.com"
	}
}

// Merge copies non connectivity variables from source Config `s` to this
// object.  The fields that are not copied are as follows:
//
//      * IpAddress
//      * Netmask
//      * Gateway
func (o *Config) Merge(s Config) {
	if s.Hostname != "" {
		o.Hostname = s.Hostname
	}

	if s.Timezone != "" {
		o.Timezone = s.Timezone
	}

	if s.Domain != "" {
		o.Domain = s.Domain
	}

	if s.UpdateServer != "" {
		o.UpdateServer = s.UpdateServer
	}

	o.VerifyUpdateServer = s.VerifyUpdateServer

	if s.LoginBanner != "" {
		o.LoginBanner = s.LoginBanner
	}

	if s.PanoramaPrimary != "" {
		o.PanoramaPrimary = s.PanoramaPrimary
	}

	if s.PanoramaSecondary != "" {
		o.PanoramaSecondary = s.PanoramaSecondary
	}

	if s.DnsPrimary != "" {
		o.DnsPrimary = s.DnsPrimary
	}

	if s.DnsSecondary != "" {
		o.DnsSecondary = s.DnsSecondary
	}

	if s.NtpPrimaryAddress != "" {
		o.NtpPrimaryAddress = s.NtpPrimaryAddress
	}

	if s.NtpPrimaryAuthType != "" {
		o.NtpPrimaryAuthType = s.NtpPrimaryAuthType
	}

	if s.NtpPrimaryKeyId != 0 {
		o.NtpPrimaryKeyId = s.NtpPrimaryKeyId
	}

	if s.NtpPrimaryAlgorithm != "" {
		o.NtpPrimaryAlgorithm = s.NtpPrimaryAlgorithm
	}

	if s.NtpPrimaryAuthKey != "" {
		o.NtpPrimaryAuthKey = s.NtpPrimaryAuthKey
	}

	if s.NtpSecondaryAddress != "" {
		o.NtpSecondaryAddress = s.NtpSecondaryAddress
	}

	if s.NtpSecondaryAuthType != "" {
		o.NtpSecondaryAuthType = s.NtpSecondaryAuthType
	}

	if s.NtpSecondaryKeyId != 0 {
		o.NtpSecondaryKeyId = s.NtpSecondaryKeyId
	}

	if s.NtpSecondaryAlgorithm != "" {
		o.NtpSecondaryAlgorithm = s.NtpSecondaryAlgorithm
	}

	if s.NtpSecondaryAuthKey != "" {
		o.NtpSecondaryAuthKey = s.NtpSecondaryAuthKey
	}

	if s.ProxyServer != "" {
		o.ProxyServer = s.ProxyServer
	}

	if s.ProxyPort != 0 {
		o.ProxyPort = s.ProxyPort
	}

	if s.ProxyUser != "" {
		o.ProxyUser = s.ProxyUser
	}

	if s.ProxyPassword != "" {
		o.ProxyPassword = s.ProxyPassword
	}
}

/** Structs / functions for this namespace. **/

func (o Config) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return "", fn(o)
}

type normalizer interface {
	Normalize() []Config
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"system"`
}

func (o *container_v1) Names() []string {
	return nil
}

func (o *container_v1) Normalize() []Config {
	ans := make([]Config, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v1) normalize() Config {
	ans := Config{
		Hostname:           o.Hostname,
		IpAddress:          o.IpAddress,
		Netmask:            o.Netmask,
		Gateway:            o.Gateway,
		Timezone:           o.Timezone,
		Domain:             o.Domain,
		UpdateServer:       o.UpdateServer,
		VerifyUpdateServer: util.AsBool(o.VerifyUpdateServer),
		LoginBanner:        o.LoginBanner,
		ProxyServer:        o.ProxyServer,
		ProxyPort:          o.ProxyPort,
		ProxyUser:          o.ProxyUser,
		ProxyPassword:      o.ProxyPassword,
	}
	if o.Panorama != nil {
		ans.PanoramaPrimary = o.Panorama.Primary
		ans.PanoramaSecondary = o.Panorama.Secondary
	}
	if o.Dns != nil {
		ans.DnsPrimary = o.Dns.Primary
		ans.DnsSecondary = o.Dns.Secondary
	}
	if o.Ntp != nil {
		if o.Ntp.Primary != nil {
			ans.NtpPrimaryAddress = o.Ntp.Primary.IpAddress
			switch {
			case o.Ntp.Primary.Auth.None != nil:
				ans.NtpPrimaryAuthType = NoAuth
			case o.Ntp.Primary.Auth.Autokey != nil:
				ans.NtpPrimaryAuthType = AutokeyAuth
			case o.Ntp.Primary.Auth.SymmetricKey != nil:
				ans.NtpPrimaryAuthType = SymmetricKeyAuth
				ans.NtpPrimaryKeyId = o.Ntp.Primary.Auth.SymmetricKey.KeyId
				switch {
				case o.Ntp.Primary.Auth.SymmetricKey.Algorithm.Sha1 != nil:
					ans.NtpPrimaryAlgorithm = Sha1
					ans.NtpPrimaryAuthKey = o.Ntp.Primary.Auth.SymmetricKey.Algorithm.Sha1.AuthenticationKey
				case o.Ntp.Primary.Auth.SymmetricKey.Algorithm.Md5 != nil:
					ans.NtpPrimaryAlgorithm = Md5
					ans.NtpPrimaryAuthKey = o.Ntp.Primary.Auth.SymmetricKey.Algorithm.Md5.AuthenticationKey
				}
			}
		}
		if o.Ntp.Secondary != nil {
			ans.NtpSecondaryAddress = o.Ntp.Secondary.IpAddress
			switch {
			case o.Ntp.Secondary.Auth.None != nil:
				ans.NtpSecondaryAuthType = NoAuth
			case o.Ntp.Secondary.Auth.Autokey != nil:
				ans.NtpSecondaryAuthType = AutokeyAuth
			case o.Ntp.Secondary.Auth.SymmetricKey != nil:
				ans.NtpSecondaryAuthType = SymmetricKeyAuth
				ans.NtpSecondaryKeyId = o.Ntp.Secondary.Auth.SymmetricKey.KeyId
				switch {
				case o.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Sha1 != nil:
					ans.NtpSecondaryAlgorithm = Sha1
					ans.NtpSecondaryAuthKey = o.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Sha1.AuthenticationKey
				case o.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Md5 != nil:
					ans.NtpSecondaryAlgorithm = Md5
					ans.NtpSecondaryAuthKey = o.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Md5.AuthenticationKey
				}
			}
		}
	}

	ans.raw = make(map[string]string)
	if o.AckLoginBanner != nil {
		ans.raw["alb"] = util.CleanRawXml(o.AckLoginBanner.Text)
	}
	if o.AuthenticationProfile != nil {
		ans.raw["ap"] = util.CleanRawXml(o.AuthenticationProfile.Text)
	}
	if o.CertificateProfile != nil {
		ans.raw["cp"] = util.CleanRawXml(o.CertificateProfile.Text)
	}
	if o.DomainLookupUrl != nil {
		ans.raw["dlu"] = util.CleanRawXml(o.DomainLookupUrl.Text)
	}
	if o.FqdnForceRefreshTime != nil {
		ans.raw["ffrt"] = util.CleanRawXml(o.FqdnForceRefreshTime.Text)
	}
	if o.FqdnRefreshTime != nil {
		ans.raw["frt"] = util.CleanRawXml(o.FqdnRefreshTime.Text)
	}
	if o.GeoLocation != nil {
		ans.raw["gl"] = util.CleanRawXml(o.GeoLocation.Text)
	}
	if o.HsmSettings != nil {
		ans.raw["hs"] = util.CleanRawXml(o.HsmSettings.Text)
	}
	if o.IpAddressLookupUrl != nil {
		ans.raw["ialu"] = util.CleanRawXml(o.IpAddressLookupUrl.Text)
	}
	if o.Ipv6Address != nil {
		ans.raw["i6a"] = util.CleanRawXml(o.Ipv6Address.Text)
	}
	if o.Ipv6DefaultGateway != nil {
		ans.raw["i6dg"] = util.CleanRawXml(o.Ipv6DefaultGateway.Text)
	}
	if o.Locale != nil {
		ans.raw["locale"] = util.CleanRawXml(o.Locale.Text)
	}
	if o.LogExportSchedule != nil {
		ans.raw["les"] = util.CleanRawXml(o.LogExportSchedule.Text)
	}
	if o.LogLink != nil {
		ans.raw["ll"] = util.CleanRawXml(o.LogLink.Text)
	}
	if o.MotdAndBanner != nil {
		ans.raw["mab"] = util.CleanRawXml(o.MotdAndBanner.Text)
	}
	if o.Mtu != nil {
		ans.raw["mtu"] = util.CleanRawXml(o.Mtu.Text)
	}
	if o.PermittedIp != nil {
		ans.raw["pi"] = util.CleanRawXml(o.PermittedIp.Text)
	}
	if o.Route != nil {
		ans.raw["route"] = util.CleanRawXml(o.Route.Text)
	}
	if o.Service != nil {
		ans.raw["service"] = util.CleanRawXml(o.Service.Text)
	}
	if o.SnmpSetting != nil {
		ans.raw["ss"] = util.CleanRawXml(o.SnmpSetting.Text)
	}
	if o.SpeedDuplex != nil {
		ans.raw["sd"] = util.CleanRawXml(o.SpeedDuplex.Text)
	}
	if o.SslTlsServiceProfile != nil {
		ans.raw["stsp"] = util.CleanRawXml(o.SslTlsServiceProfile.Text)
	}
	if o.SyslogCertificate != nil {
		ans.raw["sc"] = util.CleanRawXml(o.SyslogCertificate.Text)
	}
	if o.Type != nil {
		ans.raw["type"] = util.CleanRawXml(o.Type.Text)
	}
	if o.UpdateSchedule != nil {
		ans.raw["us"] = util.CleanRawXml(o.UpdateSchedule.Text)
	}
	if len(ans.raw) == 0 {
		ans.raw = nil
	}

	return ans
}

type entry_v1 struct {
	XMLName               xml.Name        `xml:"system"`
	Hostname              string          `xml:"hostname"`
	IpAddress             string          `xml:"ip-address,omitempty"`
	Netmask               string          `xml:"netmask,omitempty"`
	Gateway               string          `xml:"default-gateway,omitempty"`
	Timezone              string          `xml:"timezone,omitempty"`
	Domain                string          `xml:"domain,omitempty"`
	UpdateServer          string          `xml:"update-server,omitempty"`
	VerifyUpdateServer    string          `xml:"server-verification"`
	LoginBanner           string          `xml:"login-banner,omitempty"`
	ProxyServer           string          `xml:"secure-proxy-server,omitempty"`
	ProxyPort             int             `xml:"secure-proxy-port,omitempty"`
	ProxyUser             string          `xml:"secure-proxy-user,omitempty"`
	ProxyPassword         string          `xml:"secure-proxy-password,omitempty"`
	Panorama              *devicePanorama `xml:"panorama"`
	Dns                   *deviceDns      `xml:"dns-setting"`
	Ntp                   *deviceNtp      `xml:"ntp-servers"`
	AckLoginBanner        *util.RawXml    `xml:"ack-login-banner"`
	AuthenticationProfile *util.RawXml    `xml:"authentication-profile"`
	CertificateProfile    *util.RawXml    `xml:"certificate-profile"`
	DomainLookupUrl       *util.RawXml    `xml:"domain-lookup-url"`
	FqdnForceRefreshTime  *util.RawXml    `xml:"fqdn-forcerefresh-time"`
	FqdnRefreshTime       *util.RawXml    `xml:"fqdn-refresh-time"`
	GeoLocation           *util.RawXml    `xml:"geo-location"`
	HsmSettings           *util.RawXml    `xml:"hsm-settings"`
	IpAddressLookupUrl    *util.RawXml    `xml:"ip-address-lookup-url"`
	Ipv6Address           *util.RawXml    `xml:"ipv6-address"`
	Ipv6DefaultGateway    *util.RawXml    `xml:"ipv6-default-gateway"`
	Locale                *util.RawXml    `xml:"locale"`
	LogExportSchedule     *util.RawXml    `xml:"log-export-schedule"`
	LogLink               *util.RawXml    `xml:"log-link"`
	MotdAndBanner         *util.RawXml    `xml:"motd-and-banner"`
	Mtu                   *util.RawXml    `xml:"mtu"`
	PermittedIp           *util.RawXml    `xml:"permitted-ip"`
	Route                 *util.RawXml    `xml:"route"`
	Service               *util.RawXml    `xml:"service"`
	SnmpSetting           *util.RawXml    `xml:"snmp-setting"`
	SpeedDuplex           *util.RawXml    `xml:"speed-duplex"`
	SslTlsServiceProfile  *util.RawXml    `xml:"ssl-tls-service-profile"`
	SyslogCertificate     *util.RawXml    `xml:"syslog-certificate"`
	Type                  *util.RawXml    `xml:"type"`
	UpdateSchedule        *util.RawXml    `xml:"update-schedule"`
}

type devicePanorama struct {
	Primary   string
	Secondary string
}

type panoramaServers interface {
	getServers() (string, string)
}

func newDevicePanorama(ps panoramaServers) *devicePanorama {
	p, s := ps.getServers()
	return &devicePanorama{
		Primary:   p,
		Secondary: s,
	}
}

// PAN-OS 6.1+
type devicePanorama_v1 struct {
	Primary   string `xml:"local-panorama>panorama-server"`
	Secondary string `xml:"local-panorama>panorama-server-2"`
}

func (p *devicePanorama_v1) getServers() (string, string) {
	return p.Primary, p.Secondary
}

// PAN-OS 9.0+
type devicePanorama_v2 struct {
	Primary   string `xml:"cloud-service>local-panorama>panorama-server"`
	Secondary string `xml:"cloud-service>local-panorama>panorama-server-2"`
}

func (p *devicePanorama_v2) getServers() (string, string) {
	return p.Primary, p.Secondary
}

type deviceDns struct {
	Primary   string `xml:"servers>primary,omitempty"`
	Secondary string `xml:"servers>secondary,omitempty"`
}

type deviceNtp struct {
	Primary   *ntpConfig `xml:"primary-ntp-server"`
	Secondary *ntpConfig `xml:"secondary-ntp-server"`
}

type ntpConfig struct {
	IpAddress string  `xml:"ntp-server-address"`
	Auth      ntpAuth `xml:"authentication-type"`
}

type ntpAuth struct {
	None         *string `xml:"none"`
	Autokey      *string `xml:"autokey"`
	SymmetricKey *symKey `xml:"symmetric-key"`
}

type symKey struct {
	KeyId     int             `xml:"key-id"`
	Algorithm symKeyAlgorithm `xml:"algorithm"`
}

type symKeyAlgorithm struct {
	Sha1 *algorithmAuthKey `xml:"sha1"`
	Md5  *algorithmAuthKey `xml:"md5"`
}

type algorithmAuthKey struct {
	AuthenticationKey string `xml:"authentication-key"`
}

func specify_v1(c Config) interface{} {
	ans := entry_v1{
		Hostname:           c.Hostname,
		IpAddress:          c.IpAddress,
		Netmask:            c.Netmask,
		Gateway:            c.Gateway,
		Timezone:           c.Timezone,
		Domain:             c.Domain,
		UpdateServer:       c.UpdateServer,
		VerifyUpdateServer: util.YesNo(c.VerifyUpdateServer),
		LoginBanner:        c.LoginBanner,
		ProxyServer:        c.ProxyServer,
		ProxyPort:          c.ProxyPort,
		ProxyUser:          c.ProxyUser,
		ProxyPassword:      c.ProxyPassword,
	}
	// Panorama
	if c.PanoramaPrimary != "" || c.PanoramaSecondary != "" {
		ans.Panorama = newDevicePanorama(
			&devicePanorama_v1{
				c.PanoramaPrimary,
				c.PanoramaSecondary,
			},
		)
	}
	// DNS
	if c.DnsPrimary != "" || c.DnsSecondary != "" {
		ans.Dns = &deviceDns{
			c.DnsPrimary,
			c.DnsSecondary,
		}
	}
	// NTP
	ntp_config := &deviceNtp{}
	if c.NtpPrimaryAddress != "" || c.NtpPrimaryAuthType != "" {
		ntp_config.Primary = &ntpConfig{
			IpAddress: c.NtpPrimaryAddress,
		}
		var es string
		switch c.NtpPrimaryAuthType {
		case NoAuth:
			ntp_config.Primary.Auth.None = &es
		case AutokeyAuth:
			ntp_config.Primary.Auth.Autokey = &es
		case SymmetricKeyAuth:
			ntp_config.Primary.Auth.SymmetricKey = &symKey{
				KeyId: c.NtpPrimaryKeyId,
			}
			switch c.NtpPrimaryAlgorithm {
			case Sha1:
				ntp_config.Primary.Auth.SymmetricKey.Algorithm.Sha1 = &algorithmAuthKey{c.NtpPrimaryAuthKey}
			case Md5:
				ntp_config.Primary.Auth.SymmetricKey.Algorithm.Md5 = &algorithmAuthKey{c.NtpPrimaryAuthKey}
			}
		}
	}
	if c.NtpSecondaryAddress != "" || c.NtpSecondaryAuthType != "" {
		ntp_config.Secondary = &ntpConfig{
			IpAddress: c.NtpSecondaryAddress,
		}
		var es string
		switch c.NtpSecondaryAuthType {
		case NoAuth:
			ntp_config.Secondary.Auth.None = &es
		case AutokeyAuth:
			ntp_config.Secondary.Auth.Autokey = &es
		case SymmetricKeyAuth:
			ntp_config.Secondary.Auth.SymmetricKey = &symKey{
				KeyId: c.NtpSecondaryKeyId,
			}
			switch c.NtpSecondaryAlgorithm {
			case Sha1:
				ntp_config.Secondary.Auth.SymmetricKey.Algorithm.Sha1 = &algorithmAuthKey{c.NtpSecondaryAuthKey}
			case Md5:
				ntp_config.Secondary.Auth.SymmetricKey.Algorithm.Md5 = &algorithmAuthKey{c.NtpSecondaryAuthKey}
			}
		}
	}
	if ntp_config.Primary != nil || ntp_config.Secondary != nil {
		ans.Ntp = ntp_config
	}

	if text, present := c.raw["alb"]; present {
		ans.AckLoginBanner = &util.RawXml{text}
	}
	if text, present := c.raw["ap"]; present {
		ans.AuthenticationProfile = &util.RawXml{text}
	}
	if text, present := c.raw["cp"]; present {
		ans.CertificateProfile = &util.RawXml{text}
	}
	if text, present := c.raw["dlu"]; present {
		ans.DomainLookupUrl = &util.RawXml{text}
	}
	if text, present := c.raw["ffrt"]; present {
		ans.FqdnForceRefreshTime = &util.RawXml{text}
	}
	if text, present := c.raw["frt"]; present {
		ans.FqdnRefreshTime = &util.RawXml{text}
	}
	if text, present := c.raw["gl"]; present {
		ans.GeoLocation = &util.RawXml{text}
	}
	if text, present := c.raw["hs"]; present {
		ans.HsmSettings = &util.RawXml{text}
	}
	if text, present := c.raw["ialu"]; present {
		ans.IpAddressLookupUrl = &util.RawXml{text}
	}
	if text, present := c.raw["i6a"]; present {
		ans.Ipv6Address = &util.RawXml{text}
	}
	if text, present := c.raw["i6dg"]; present {
		ans.Ipv6DefaultGateway = &util.RawXml{text}
	}
	if text, present := c.raw["locale"]; present {
		ans.Locale = &util.RawXml{text}
	}
	if text, present := c.raw["les"]; present {
		ans.LogExportSchedule = &util.RawXml{text}
	}
	if text, present := c.raw["ll"]; present {
		ans.LogLink = &util.RawXml{text}
	}
	if text, present := c.raw["mab"]; present {
		ans.MotdAndBanner = &util.RawXml{text}
	}
	if text, present := c.raw["mtu"]; present {
		ans.Mtu = &util.RawXml{text}
	}
	if text, present := c.raw["pi"]; present {
		ans.PermittedIp = &util.RawXml{text}
	}
	if text, present := c.raw["route"]; present {
		ans.Route = &util.RawXml{text}
	}
	if text, present := c.raw["service"]; present {
		ans.Service = &util.RawXml{text}
	}
	if text, present := c.raw["ss"]; present {
		ans.SnmpSetting = &util.RawXml{text}
	}
	if text, present := c.raw["sd"]; present {
		ans.SpeedDuplex = &util.RawXml{text}
	}
	if text, present := c.raw["stsp"]; present {
		ans.SslTlsServiceProfile = &util.RawXml{text}
	}
	if text, present := c.raw["sc"]; present {
		ans.SyslogCertificate = &util.RawXml{text}
	}
	if text, present := c.raw["type"]; present {
		ans.Type = &util.RawXml{text}
	}
	if text, present := c.raw["us"]; present {
		ans.UpdateSchedule = &util.RawXml{text}
	}

	return ans
}

func specify_v2(c Config) interface{} {
  ans := entry_v1{
		Hostname:           c.Hostname,
		IpAddress:          c.IpAddress,
		Netmask:            c.Netmask,
		Gateway:            c.Gateway,
		Timezone:           c.Timezone,
		Domain:             c.Domain,
		UpdateServer:       c.UpdateServer,
		VerifyUpdateServer: util.YesNo(c.VerifyUpdateServer),
		LoginBanner:        c.LoginBanner,
		ProxyServer:        c.ProxyServer,
		ProxyPort:          c.ProxyPort,
		ProxyUser:          c.ProxyUser,
		ProxyPassword:      c.ProxyPassword,
	}
	// Panorama
	if c.PanoramaPrimary != "" || c.PanoramaSecondary != "" {
		ans.Panorama = newDevicePanorama(
			&devicePanorama_v2{
				c.PanoramaPrimary,
				c.PanoramaSecondary,
			},
		)
	}
	// DNS
	if c.DnsPrimary != "" || c.DnsSecondary != "" {
		ans.Dns = &deviceDns{
			c.DnsPrimary,
			c.DnsSecondary,
		}
	}
	// NTP
	ntp_config := &deviceNtp{}
	if c.NtpPrimaryAddress != "" || c.NtpPrimaryAuthType != "" {
		ntp_config.Primary = &ntpConfig{
			IpAddress: c.NtpPrimaryAddress,
		}
		var es string
		switch c.NtpPrimaryAuthType {
		case NoAuth:
			ntp_config.Primary.Auth.None = &es
		case AutokeyAuth:
			ntp_config.Primary.Auth.Autokey = &es
		case SymmetricKeyAuth:
			ntp_config.Primary.Auth.SymmetricKey = &symKey{
				KeyId: c.NtpPrimaryKeyId,
			}
			switch c.NtpPrimaryAlgorithm {
			case Sha1:
				ntp_config.Primary.Auth.SymmetricKey.Algorithm.Sha1 = &algorithmAuthKey{c.NtpPrimaryAuthKey}
			case Md5:
				ntp_config.Primary.Auth.SymmetricKey.Algorithm.Md5 = &algorithmAuthKey{c.NtpPrimaryAuthKey}
			}
		}
	}
	if c.NtpSecondaryAddress != "" || c.NtpSecondaryAuthType != "" {
		ntp_config.Secondary = &ntpConfig{
			IpAddress: c.NtpSecondaryAddress,
		}
		var es string
		switch c.NtpSecondaryAuthType {
		case NoAuth:
			ntp_config.Secondary.Auth.None = &es
		case AutokeyAuth:
			ntp_config.Secondary.Auth.Autokey = &es
		case SymmetricKeyAuth:
			ntp_config.Secondary.Auth.SymmetricKey = &symKey{
				KeyId: c.NtpSecondaryKeyId,
			}
			switch c.NtpSecondaryAlgorithm {
			case Sha1:
				ntp_config.Secondary.Auth.SymmetricKey.Algorithm.Sha1 = &algorithmAuthKey{c.NtpSecondaryAuthKey}
			case Md5:
				ntp_config.Secondary.Auth.SymmetricKey.Algorithm.Md5 = &algorithmAuthKey{c.NtpSecondaryAuthKey}
			}
		}
	}
	if ntp_config.Primary != nil || ntp_config.Secondary != nil {
		ans.Ntp = ntp_config
	}

	if text, present := c.raw["alb"]; present {
		ans.AckLoginBanner = &util.RawXml{text}
	}
	if text, present := c.raw["ap"]; present {
		ans.AuthenticationProfile = &util.RawXml{text}
	}
	if text, present := c.raw["cp"]; present {
		ans.CertificateProfile = &util.RawXml{text}
	}
	if text, present := c.raw["dlu"]; present {
		ans.DomainLookupUrl = &util.RawXml{text}
	}
	if text, present := c.raw["ffrt"]; present {
		ans.FqdnForceRefreshTime = &util.RawXml{text}
	}
	if text, present := c.raw["frt"]; present {
		ans.FqdnRefreshTime = &util.RawXml{text}
	}
	if text, present := c.raw["gl"]; present {
		ans.GeoLocation = &util.RawXml{text}
	}
	if text, present := c.raw["hs"]; present {
		ans.HsmSettings = &util.RawXml{text}
	}
	if text, present := c.raw["ialu"]; present {
		ans.IpAddressLookupUrl = &util.RawXml{text}
	}
	if text, present := c.raw["i6a"]; present {
		ans.Ipv6Address = &util.RawXml{text}
	}
	if text, present := c.raw["i6dg"]; present {
		ans.Ipv6DefaultGateway = &util.RawXml{text}
	}
	if text, present := c.raw["locale"]; present {
		ans.Locale = &util.RawXml{text}
	}
	if text, present := c.raw["les"]; present {
		ans.LogExportSchedule = &util.RawXml{text}
	}
	if text, present := c.raw["ll"]; present {
		ans.LogLink = &util.RawXml{text}
	}
	if text, present := c.raw["mab"]; present {
		ans.MotdAndBanner = &util.RawXml{text}
	}
	if text, present := c.raw["mtu"]; present {
		ans.Mtu = &util.RawXml{text}
	}
	if text, present := c.raw["pi"]; present {
		ans.PermittedIp = &util.RawXml{text}
	}
	if text, present := c.raw["route"]; present {
		ans.Route = &util.RawXml{text}
	}
	if text, present := c.raw["service"]; present {
		ans.Service = &util.RawXml{text}
	}
	if text, present := c.raw["ss"]; present {
		ans.SnmpSetting = &util.RawXml{text}
	}
	if text, present := c.raw["sd"]; present {
		ans.SpeedDuplex = &util.RawXml{text}
	}
	if text, present := c.raw["stsp"]; present {
		ans.SslTlsServiceProfile = &util.RawXml{text}
	}
	if text, present := c.raw["sc"]; present {
		ans.SyslogCertificate = &util.RawXml{text}
	}
	if text, present := c.raw["type"]; present {
		ans.Type = &util.RawXml{text}
	}
	if text, present := c.raw["us"]; present {
		ans.UpdateSchedule = &util.RawXml{text}
	}

	return ans
}
