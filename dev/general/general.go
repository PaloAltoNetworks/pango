// Package general is the client.Device.GeneralSettings namespace.
//
// Normalized object: Config
package general

import (
    "encoding/xml"

    "github.com/PaloAltoNetworks/xapi/util"
)


// Constants for NTP auth types.
const (
    NoAuth = "none"
    AutokeyAuth = "autokey"
    SymmetricKeyAuth = "symmetric-key"
)

// Constants for NTP algorithms.
const (
    Sha1 = "sha1"
    Md5 = "md5"
)

// Config is a normalized, version independent representation of a device's
// general settings.
type Config struct {
    Hostname string
    IpAddress string
    Netmask string
    Gateway string
    Timezone string
    DnsPrimary string
    DnsSecondary string
    NtpPrimaryAddress string
    NtpPrimaryAuthType string
    NtpPrimaryKeyId int
    NtpPrimaryAlgorithm string
    NtpPrimaryAuthKey string
    NtpSecondaryAddress string
    NtpSecondaryAuthType string
    NtpSecondaryKeyId int
    NtpSecondaryAlgorithm string
    NtpSecondaryAuthKey string
}

// Gset is a namespace struct, included as part of xapi.Client.
type General struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *General) Initialize(con util.XapiClient) {
    c.con = con
}

// Show performs SHOW to retrieve the device's general settings.
func (c *General) Show() (Config, error) {
    c.con.LogQuery("(show) general settings")
    return c.details(c.con.Show)
}

// Get performs GET to retrieve the device's general settings.
func (c *General) Get() (Config, error) {
    c.con.LogQuery("(get) general settings")
    return c.details(c.con.Get)
}

// Set performs SET to create / update the device's general settings.
func (c *General) Set(e Config) error {
    var err error
    _, fn := c.versioning()
    c.con.LogAction("(set) general settings")

    path := c.xpath()
    path = path[:len(path) - 1]

    _, err = c.con.Set(path, fn(e), nil, nil)
    return err
}

/** Internal functions for the General struct **/

func (c *General) versioning() (normalizer, func(Config) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *General) details(fn func(interface{}, interface{}, interface{}) ([]byte, error)) (Config, error) {
    path := c.xpath()
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Config{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *General) xpath() []string {
    return []string{
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "deviceconfig",
        "system",
    }
}

/** Structs / functions for this namespace. **/

type normalizer interface {
    Normalize() Config
}

type container_v1 struct {
    Answer config_v1 `xml:"result>system"`
}

func (o *container_v1) Normalize() Config {
    ans := Config{
        Hostname: o.Answer.Hostname,
        IpAddress: o.Answer.IpAddress,
        Netmask: o.Answer.Netmask,
        Gateway: o.Answer.Gateway,
        Timezone: o.Answer.Timezone,
    }
    if o.Answer.Dns != nil {
        ans.DnsPrimary = o.Answer.Dns.Primary
        ans.DnsSecondary = o.Answer.Dns.Secondary
    }
    if o.Answer.Ntp != nil {
        if o.Answer.Ntp.Primary != nil {
            ans.NtpPrimaryAddress = o.Answer.Ntp.Primary.IpAddress
            switch {
            case o.Answer.Ntp.Primary.Auth.None != nil:
                ans.NtpPrimaryAuthType = NoAuth
            case o.Answer.Ntp.Primary.Auth.Autokey != nil:
                ans.NtpPrimaryAuthType = AutokeyAuth
            case o.Answer.Ntp.Primary.Auth.SymmetricKey != nil:
                ans.NtpPrimaryAuthType = SymmetricKeyAuth
                ans.NtpPrimaryKeyId = o.Answer.Ntp.Primary.Auth.SymmetricKey.KeyId
                switch {
                case o.Answer.Ntp.Primary.Auth.SymmetricKey.Algorithm.Sha1 != nil:
                    ans.NtpPrimaryAlgorithm = Sha1
                    ans.NtpPrimaryAuthKey = o.Answer.Ntp.Primary.Auth.SymmetricKey.Algorithm.Sha1.AuthenticationKey
                case o.Answer.Ntp.Primary.Auth.SymmetricKey.Algorithm.Md5 != nil:
                    ans.NtpPrimaryAlgorithm = Md5
                    ans.NtpPrimaryAuthKey = o.Answer.Ntp.Primary.Auth.SymmetricKey.Algorithm.Md5.AuthenticationKey
                }
            }
        }
        if o.Answer.Ntp.Secondary != nil {
            ans.NtpSecondaryAddress = o.Answer.Ntp.Secondary.IpAddress
            switch {
            case o.Answer.Ntp.Secondary.Auth.None != nil:
                ans.NtpSecondaryAuthType = NoAuth
            case o.Answer.Ntp.Secondary.Auth.Autokey != nil:
                ans.NtpSecondaryAuthType = AutokeyAuth
            case o.Answer.Ntp.Secondary.Auth.SymmetricKey != nil:
                ans.NtpSecondaryAuthType = SymmetricKeyAuth
                ans.NtpSecondaryKeyId = o.Answer.Ntp.Secondary.Auth.SymmetricKey.KeyId
                switch {
                case o.Answer.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Sha1 != nil:
                    ans.NtpSecondaryAlgorithm = Sha1
                    ans.NtpSecondaryAuthKey = o.Answer.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Sha1.AuthenticationKey
                case o.Answer.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Md5 != nil:
                    ans.NtpSecondaryAlgorithm = Md5
                    ans.NtpSecondaryAuthKey = o.Answer.Ntp.Secondary.Auth.SymmetricKey.Algorithm.Md5.AuthenticationKey
                }
            }
        }
    }

    return ans
}

type config_v1 struct {
    XMLName xml.Name `xml:"system"`
    Hostname string `xml:"hostname"`
    IpAddress string `xml:"ip-address"`
    Netmask string `xml:"netmask"`
    Gateway string `xml:"default-gateway"`
    Timezone string `xml:"timezone"`
    Dns *deviceDns `xml:"dns-setting"`
    Ntp *deviceNtp `xml:"ntp-servers"`
}

type deviceDns struct {
    Primary string `xml:"servers>primary"`
    Secondary string `xml:"servers>secondary"`
}

type deviceNtp struct {
    Primary *ntpConfig `xml:"primary-ntp-server"`
    Secondary *ntpConfig `xml:"secondary-ntp-server"`
}

type ntpConfig struct {
    IpAddress string `xml:"ntp-server-address"`
    Auth ntpAuth `xml:"authentication-type"`
}

type ntpAuth struct {
    None *string `xml:"none"`
    Autokey *string `xml:"autokey"`
    SymmetricKey *symKey `xml:"symmetric-key"`
}

type symKey struct {
    KeyId int `xml:"key-id"`
    Algorithm symKeyAlgorithm `xml:"algorithm"`
}

type symKeyAlgorithm struct {
    Sha1 *algorithmAuthKey `xml:"sha1"`
    Md5 *algorithmAuthKey `xml:"md5"`
}

type algorithmAuthKey struct {
    AuthenticationKey string `xml:"authentication-key"`
}

func specify_v1(c Config) interface{} {
    ans := config_v1{
        Hostname: c.Hostname,
        IpAddress: c.IpAddress,
        Netmask: c.Netmask,
        Gateway: c.Gateway,
        Timezone: c.Timezone,
    }
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

    return ans
}
