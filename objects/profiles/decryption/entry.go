package decryption

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
	suffix = []string{"profiles", "decryption", "$name"}
)

type Entry struct {
	Name                string
	DisableOverride     *string
	ForwardedOnly       *bool
	Interface           *string
	SshProxy            *SshProxy
	SslForwardProxy     *SslForwardProxy
	SslInboundProxy     *SslInboundProxy
	SslNoProxy          *SslNoProxy
	SslProtocolSettings *SslProtocolSettings
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type SshProxy struct {
	BlockIfNoResource       *bool
	BlockSshErrors          *bool
	BlockUnsupportedAlg     *bool
	BlockUnsupportedVersion *bool
	Misc                    []generic.Xml
	MiscAttributes          []xml.Attr
}
type SslForwardProxy struct {
	AutoIncludeAltname            *bool
	BlockClientCert               *bool
	BlockExpiredCertificate       *bool
	BlockIfHsmUnavailable         *bool
	BlockIfNoResource             *bool
	BlockIfSniMismatch            *bool
	BlockTimeoutCert              *bool
	BlockTls13DowngradeNoResource *bool
	BlockUnknownCert              *bool
	BlockUnsupportedCipher        *bool
	BlockUnsupportedVersion       *bool
	BlockUntrustedIssuer          *bool
	RestrictCertExts              *bool
	StripAlpn                     *bool
	Misc                          []generic.Xml
	MiscAttributes                []xml.Attr
}
type SslInboundProxy struct {
	BlockIfHsmUnavailable         *bool
	BlockIfNoResource             *bool
	BlockTls13DowngradeNoResource *bool
	BlockUnsupportedCipher        *bool
	BlockUnsupportedVersion       *bool
	Misc                          []generic.Xml
	MiscAttributes                []xml.Attr
}
type SslNoProxy struct {
	BlockExpiredCertificate *bool
	BlockUntrustedIssuer    *bool
	Misc                    []generic.Xml
	MiscAttributes          []xml.Attr
}
type SslProtocolSettings struct {
	AuthAlgoMd5             *bool
	AuthAlgoSha1            *bool
	AuthAlgoSha256          *bool
	AuthAlgoSha384          *bool
	EncAlgo3des             *bool
	EncAlgoAes128Cbc        *bool
	EncAlgoAes128Gcm        *bool
	EncAlgoAes256Cbc        *bool
	EncAlgoAes256Gcm        *bool
	EncAlgoChacha20Poly1305 *bool
	EncAlgoRc4              *bool
	KeyxchgAlgoDhe          *bool
	KeyxchgAlgoEcdhe        *bool
	KeyxchgAlgoRsa          *bool
	MaxVersion              *string
	MinVersion              *string
	Misc                    []generic.Xml
	MiscAttributes          []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
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
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName             xml.Name                `xml:"entry"`
	Name                string                  `xml:"name,attr"`
	DisableOverride     *string                 `xml:"disable-override,omitempty"`
	ForwardedOnly       *string                 `xml:"forwarded-only,omitempty"`
	Interface           *string                 `xml:"interface,omitempty"`
	SshProxy            *sshProxyXml            `xml:"ssh-proxy,omitempty"`
	SslForwardProxy     *sslForwardProxyXml     `xml:"ssl-forward-proxy,omitempty"`
	SslInboundProxy     *sslInboundProxyXml     `xml:"ssl-inbound-proxy,omitempty"`
	SslNoProxy          *sslNoProxyXml          `xml:"ssl-no-proxy,omitempty"`
	SslProtocolSettings *sslProtocolSettingsXml `xml:"ssl-protocol-settings,omitempty"`
	Misc                []generic.Xml           `xml:",any"`
	MiscAttributes      []xml.Attr              `xml:",any,attr"`
}
type sshProxyXml struct {
	BlockIfNoResource       *string       `xml:"block-if-no-resource,omitempty"`
	BlockSshErrors          *string       `xml:"block-ssh-errors,omitempty"`
	BlockUnsupportedAlg     *string       `xml:"block-unsupported-alg,omitempty"`
	BlockUnsupportedVersion *string       `xml:"block-unsupported-version,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type sslForwardProxyXml struct {
	AutoIncludeAltname            *string       `xml:"auto-include-altname,omitempty"`
	BlockClientCert               *string       `xml:"block-client-cert,omitempty"`
	BlockExpiredCertificate       *string       `xml:"block-expired-certificate,omitempty"`
	BlockIfHsmUnavailable         *string       `xml:"block-if-hsm-unavailable,omitempty"`
	BlockIfNoResource             *string       `xml:"block-if-no-resource,omitempty"`
	BlockIfSniMismatch            *string       `xml:"block-if-sni-mismatch,omitempty"`
	BlockTimeoutCert              *string       `xml:"block-timeout-cert,omitempty"`
	BlockTls13DowngradeNoResource *string       `xml:"block-tls13-downgrade-no-resource,omitempty"`
	BlockUnknownCert              *string       `xml:"block-unknown-cert,omitempty"`
	BlockUnsupportedCipher        *string       `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion       *string       `xml:"block-unsupported-version,omitempty"`
	BlockUntrustedIssuer          *string       `xml:"block-untrusted-issuer,omitempty"`
	RestrictCertExts              *string       `xml:"restrict-cert-exts,omitempty"`
	StripAlpn                     *string       `xml:"strip-alpn,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
	MiscAttributes                []xml.Attr    `xml:",any,attr"`
}
type sslInboundProxyXml struct {
	BlockIfHsmUnavailable         *string       `xml:"block-if-hsm-unavailable,omitempty"`
	BlockIfNoResource             *string       `xml:"block-if-no-resource,omitempty"`
	BlockTls13DowngradeNoResource *string       `xml:"block-tls13-downgrade-no-resource,omitempty"`
	BlockUnsupportedCipher        *string       `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion       *string       `xml:"block-unsupported-version,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
	MiscAttributes                []xml.Attr    `xml:",any,attr"`
}
type sslNoProxyXml struct {
	BlockExpiredCertificate *string       `xml:"block-expired-certificate,omitempty"`
	BlockUntrustedIssuer    *string       `xml:"block-untrusted-issuer,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type sslProtocolSettingsXml struct {
	AuthAlgoMd5             *string       `xml:"auth-algo-md5,omitempty"`
	AuthAlgoSha1            *string       `xml:"auth-algo-sha1,omitempty"`
	AuthAlgoSha256          *string       `xml:"auth-algo-sha256,omitempty"`
	AuthAlgoSha384          *string       `xml:"auth-algo-sha384,omitempty"`
	EncAlgo3des             *string       `xml:"enc-algo-3des,omitempty"`
	EncAlgoAes128Cbc        *string       `xml:"enc-algo-aes-128-cbc,omitempty"`
	EncAlgoAes128Gcm        *string       `xml:"enc-algo-aes-128-gcm,omitempty"`
	EncAlgoAes256Cbc        *string       `xml:"enc-algo-aes-256-cbc,omitempty"`
	EncAlgoAes256Gcm        *string       `xml:"enc-algo-aes-256-gcm,omitempty"`
	EncAlgoChacha20Poly1305 *string       `xml:"enc-algo-chacha20-poly1305,omitempty"`
	EncAlgoRc4              *string       `xml:"enc-algo-rc4,omitempty"`
	KeyxchgAlgoDhe          *string       `xml:"keyxchg-algo-dhe,omitempty"`
	KeyxchgAlgoEcdhe        *string       `xml:"keyxchg-algo-ecdhe,omitempty"`
	KeyxchgAlgoRsa          *string       `xml:"keyxchg-algo-rsa,omitempty"`
	MaxVersion              *string       `xml:"max-version,omitempty"`
	MinVersion              *string       `xml:"min-version,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type entryXml_11_0_2 struct {
	XMLName             xml.Name                       `xml:"entry"`
	Name                string                         `xml:"name,attr"`
	DisableOverride     *string                        `xml:"disable-override,omitempty"`
	ForwardedOnly       *string                        `xml:"forwarded-only,omitempty"`
	Interface           *string                        `xml:"interface,omitempty"`
	SshProxy            *sshProxyXml_11_0_2            `xml:"ssh-proxy,omitempty"`
	SslForwardProxy     *sslForwardProxyXml_11_0_2     `xml:"ssl-forward-proxy,omitempty"`
	SslInboundProxy     *sslInboundProxyXml_11_0_2     `xml:"ssl-inbound-proxy,omitempty"`
	SslNoProxy          *sslNoProxyXml_11_0_2          `xml:"ssl-no-proxy,omitempty"`
	SslProtocolSettings *sslProtocolSettingsXml_11_0_2 `xml:"ssl-protocol-settings,omitempty"`
	Misc                []generic.Xml                  `xml:",any"`
	MiscAttributes      []xml.Attr                     `xml:",any,attr"`
}
type sshProxyXml_11_0_2 struct {
	BlockIfNoResource       *string       `xml:"block-if-no-resource,omitempty"`
	BlockSshErrors          *string       `xml:"block-ssh-errors,omitempty"`
	BlockUnsupportedAlg     *string       `xml:"block-unsupported-alg,omitempty"`
	BlockUnsupportedVersion *string       `xml:"block-unsupported-version,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type sslForwardProxyXml_11_0_2 struct {
	AutoIncludeAltname            *string       `xml:"auto-include-altname,omitempty"`
	BlockClientCert               *string       `xml:"block-client-cert,omitempty"`
	BlockExpiredCertificate       *string       `xml:"block-expired-certificate,omitempty"`
	BlockIfHsmUnavailable         *string       `xml:"block-if-hsm-unavailable,omitempty"`
	BlockIfNoResource             *string       `xml:"block-if-no-resource,omitempty"`
	BlockIfSniMismatch            *string       `xml:"block-if-sni-mismatch,omitempty"`
	BlockTimeoutCert              *string       `xml:"block-timeout-cert,omitempty"`
	BlockTls13DowngradeNoResource *string       `xml:"block-tls13-downgrade-no-resource,omitempty"`
	BlockUnknownCert              *string       `xml:"block-unknown-cert,omitempty"`
	BlockUnsupportedCipher        *string       `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion       *string       `xml:"block-unsupported-version,omitempty"`
	BlockUntrustedIssuer          *string       `xml:"block-untrusted-issuer,omitempty"`
	RestrictCertExts              *string       `xml:"restrict-cert-exts,omitempty"`
	StripAlpn                     *string       `xml:"strip-alpn,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
	MiscAttributes                []xml.Attr    `xml:",any,attr"`
}
type sslInboundProxyXml_11_0_2 struct {
	BlockIfHsmUnavailable         *string       `xml:"block-if-hsm-unavailable,omitempty"`
	BlockIfNoResource             *string       `xml:"block-if-no-resource,omitempty"`
	BlockTls13DowngradeNoResource *string       `xml:"block-tls13-downgrade-no-resource,omitempty"`
	BlockUnsupportedCipher        *string       `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion       *string       `xml:"block-unsupported-version,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
	MiscAttributes                []xml.Attr    `xml:",any,attr"`
}
type sslNoProxyXml_11_0_2 struct {
	BlockExpiredCertificate *string       `xml:"block-expired-certificate,omitempty"`
	BlockUntrustedIssuer    *string       `xml:"block-untrusted-issuer,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type sslProtocolSettingsXml_11_0_2 struct {
	AuthAlgoMd5             *string       `xml:"auth-algo-md5,omitempty"`
	AuthAlgoSha1            *string       `xml:"auth-algo-sha1,omitempty"`
	AuthAlgoSha256          *string       `xml:"auth-algo-sha256,omitempty"`
	AuthAlgoSha384          *string       `xml:"auth-algo-sha384,omitempty"`
	EncAlgo3des             *string       `xml:"enc-algo-3des,omitempty"`
	EncAlgoAes128Cbc        *string       `xml:"enc-algo-aes-128-cbc,omitempty"`
	EncAlgoAes128Gcm        *string       `xml:"enc-algo-aes-128-gcm,omitempty"`
	EncAlgoAes256Cbc        *string       `xml:"enc-algo-aes-256-cbc,omitempty"`
	EncAlgoAes256Gcm        *string       `xml:"enc-algo-aes-256-gcm,omitempty"`
	EncAlgoChacha20Poly1305 *string       `xml:"enc-algo-chacha20-poly1305,omitempty"`
	EncAlgoRc4              *string       `xml:"enc-algo-rc4,omitempty"`
	KeyxchgAlgoDhe          *string       `xml:"keyxchg-algo-dhe,omitempty"`
	KeyxchgAlgoEcdhe        *string       `xml:"keyxchg-algo-ecdhe,omitempty"`
	KeyxchgAlgoRsa          *string       `xml:"keyxchg-algo-rsa,omitempty"`
	MaxVersion              *string       `xml:"max-version,omitempty"`
	MinVersion              *string       `xml:"min-version,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DisableOverride = s.DisableOverride
	o.ForwardedOnly = util.YesNo(s.ForwardedOnly, nil)
	o.Interface = s.Interface
	if s.SshProxy != nil {
		var obj sshProxyXml
		obj.MarshalFromObject(*s.SshProxy)
		o.SshProxy = &obj
	}
	if s.SslForwardProxy != nil {
		var obj sslForwardProxyXml
		obj.MarshalFromObject(*s.SslForwardProxy)
		o.SslForwardProxy = &obj
	}
	if s.SslInboundProxy != nil {
		var obj sslInboundProxyXml
		obj.MarshalFromObject(*s.SslInboundProxy)
		o.SslInboundProxy = &obj
	}
	if s.SslNoProxy != nil {
		var obj sslNoProxyXml
		obj.MarshalFromObject(*s.SslNoProxy)
		o.SslNoProxy = &obj
	}
	if s.SslProtocolSettings != nil {
		var obj sslProtocolSettingsXml
		obj.MarshalFromObject(*s.SslProtocolSettings)
		o.SslProtocolSettings = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var sshProxyVal *SshProxy
	if o.SshProxy != nil {
		obj, err := o.SshProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sshProxyVal = obj
	}
	var sslForwardProxyVal *SslForwardProxy
	if o.SslForwardProxy != nil {
		obj, err := o.SslForwardProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslForwardProxyVal = obj
	}
	var sslInboundProxyVal *SslInboundProxy
	if o.SslInboundProxy != nil {
		obj, err := o.SslInboundProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslInboundProxyVal = obj
	}
	var sslNoProxyVal *SslNoProxy
	if o.SslNoProxy != nil {
		obj, err := o.SslNoProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslNoProxyVal = obj
	}
	var sslProtocolSettingsVal *SslProtocolSettings
	if o.SslProtocolSettings != nil {
		obj, err := o.SslProtocolSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslProtocolSettingsVal = obj
	}

	result := &Entry{
		Name:                o.Name,
		DisableOverride:     o.DisableOverride,
		ForwardedOnly:       util.AsBool(o.ForwardedOnly, nil),
		Interface:           o.Interface,
		SshProxy:            sshProxyVal,
		SslForwardProxy:     sslForwardProxyVal,
		SslInboundProxy:     sslInboundProxyVal,
		SslNoProxy:          sslNoProxyVal,
		SslProtocolSettings: sslProtocolSettingsVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *sshProxyXml) MarshalFromObject(s SshProxy) {
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockSshErrors = util.YesNo(s.BlockSshErrors, nil)
	o.BlockUnsupportedAlg = util.YesNo(s.BlockUnsupportedAlg, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sshProxyXml) UnmarshalToObject() (*SshProxy, error) {

	result := &SshProxy{
		BlockIfNoResource:       util.AsBool(o.BlockIfNoResource, nil),
		BlockSshErrors:          util.AsBool(o.BlockSshErrors, nil),
		BlockUnsupportedAlg:     util.AsBool(o.BlockUnsupportedAlg, nil),
		BlockUnsupportedVersion: util.AsBool(o.BlockUnsupportedVersion, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *sslForwardProxyXml) MarshalFromObject(s SslForwardProxy) {
	o.AutoIncludeAltname = util.YesNo(s.AutoIncludeAltname, nil)
	o.BlockClientCert = util.YesNo(s.BlockClientCert, nil)
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockIfHsmUnavailable = util.YesNo(s.BlockIfHsmUnavailable, nil)
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockIfSniMismatch = util.YesNo(s.BlockIfSniMismatch, nil)
	o.BlockTimeoutCert = util.YesNo(s.BlockTimeoutCert, nil)
	o.BlockTls13DowngradeNoResource = util.YesNo(s.BlockTls13DowngradeNoResource, nil)
	o.BlockUnknownCert = util.YesNo(s.BlockUnknownCert, nil)
	o.BlockUnsupportedCipher = util.YesNo(s.BlockUnsupportedCipher, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.BlockUntrustedIssuer = util.YesNo(s.BlockUntrustedIssuer, nil)
	o.RestrictCertExts = util.YesNo(s.RestrictCertExts, nil)
	o.StripAlpn = util.YesNo(s.StripAlpn, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslForwardProxyXml) UnmarshalToObject() (*SslForwardProxy, error) {

	result := &SslForwardProxy{
		AutoIncludeAltname:            util.AsBool(o.AutoIncludeAltname, nil),
		BlockClientCert:               util.AsBool(o.BlockClientCert, nil),
		BlockExpiredCertificate:       util.AsBool(o.BlockExpiredCertificate, nil),
		BlockIfHsmUnavailable:         util.AsBool(o.BlockIfHsmUnavailable, nil),
		BlockIfNoResource:             util.AsBool(o.BlockIfNoResource, nil),
		BlockIfSniMismatch:            util.AsBool(o.BlockIfSniMismatch, nil),
		BlockTimeoutCert:              util.AsBool(o.BlockTimeoutCert, nil),
		BlockTls13DowngradeNoResource: util.AsBool(o.BlockTls13DowngradeNoResource, nil),
		BlockUnknownCert:              util.AsBool(o.BlockUnknownCert, nil),
		BlockUnsupportedCipher:        util.AsBool(o.BlockUnsupportedCipher, nil),
		BlockUnsupportedVersion:       util.AsBool(o.BlockUnsupportedVersion, nil),
		BlockUntrustedIssuer:          util.AsBool(o.BlockUntrustedIssuer, nil),
		RestrictCertExts:              util.AsBool(o.RestrictCertExts, nil),
		StripAlpn:                     util.AsBool(o.StripAlpn, nil),
		Misc:                          o.Misc,
		MiscAttributes:                o.MiscAttributes,
	}
	return result, nil
}
func (o *sslInboundProxyXml) MarshalFromObject(s SslInboundProxy) {
	o.BlockIfHsmUnavailable = util.YesNo(s.BlockIfHsmUnavailable, nil)
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockTls13DowngradeNoResource = util.YesNo(s.BlockTls13DowngradeNoResource, nil)
	o.BlockUnsupportedCipher = util.YesNo(s.BlockUnsupportedCipher, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslInboundProxyXml) UnmarshalToObject() (*SslInboundProxy, error) {

	result := &SslInboundProxy{
		BlockIfHsmUnavailable:         util.AsBool(o.BlockIfHsmUnavailable, nil),
		BlockIfNoResource:             util.AsBool(o.BlockIfNoResource, nil),
		BlockTls13DowngradeNoResource: util.AsBool(o.BlockTls13DowngradeNoResource, nil),
		BlockUnsupportedCipher:        util.AsBool(o.BlockUnsupportedCipher, nil),
		BlockUnsupportedVersion:       util.AsBool(o.BlockUnsupportedVersion, nil),
		Misc:                          o.Misc,
		MiscAttributes:                o.MiscAttributes,
	}
	return result, nil
}
func (o *sslNoProxyXml) MarshalFromObject(s SslNoProxy) {
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockUntrustedIssuer = util.YesNo(s.BlockUntrustedIssuer, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslNoProxyXml) UnmarshalToObject() (*SslNoProxy, error) {

	result := &SslNoProxy{
		BlockExpiredCertificate: util.AsBool(o.BlockExpiredCertificate, nil),
		BlockUntrustedIssuer:    util.AsBool(o.BlockUntrustedIssuer, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *sslProtocolSettingsXml) MarshalFromObject(s SslProtocolSettings) {
	o.AuthAlgoMd5 = util.YesNo(s.AuthAlgoMd5, nil)
	o.AuthAlgoSha1 = util.YesNo(s.AuthAlgoSha1, nil)
	o.AuthAlgoSha256 = util.YesNo(s.AuthAlgoSha256, nil)
	o.AuthAlgoSha384 = util.YesNo(s.AuthAlgoSha384, nil)
	o.EncAlgo3des = util.YesNo(s.EncAlgo3des, nil)
	o.EncAlgoAes128Cbc = util.YesNo(s.EncAlgoAes128Cbc, nil)
	o.EncAlgoAes128Gcm = util.YesNo(s.EncAlgoAes128Gcm, nil)
	o.EncAlgoAes256Cbc = util.YesNo(s.EncAlgoAes256Cbc, nil)
	o.EncAlgoAes256Gcm = util.YesNo(s.EncAlgoAes256Gcm, nil)
	o.EncAlgoChacha20Poly1305 = util.YesNo(s.EncAlgoChacha20Poly1305, nil)
	o.EncAlgoRc4 = util.YesNo(s.EncAlgoRc4, nil)
	o.KeyxchgAlgoDhe = util.YesNo(s.KeyxchgAlgoDhe, nil)
	o.KeyxchgAlgoEcdhe = util.YesNo(s.KeyxchgAlgoEcdhe, nil)
	o.KeyxchgAlgoRsa = util.YesNo(s.KeyxchgAlgoRsa, nil)
	o.MaxVersion = s.MaxVersion
	o.MinVersion = s.MinVersion
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslProtocolSettingsXml) UnmarshalToObject() (*SslProtocolSettings, error) {

	result := &SslProtocolSettings{
		AuthAlgoMd5:             util.AsBool(o.AuthAlgoMd5, nil),
		AuthAlgoSha1:            util.AsBool(o.AuthAlgoSha1, nil),
		AuthAlgoSha256:          util.AsBool(o.AuthAlgoSha256, nil),
		AuthAlgoSha384:          util.AsBool(o.AuthAlgoSha384, nil),
		EncAlgo3des:             util.AsBool(o.EncAlgo3des, nil),
		EncAlgoAes128Cbc:        util.AsBool(o.EncAlgoAes128Cbc, nil),
		EncAlgoAes128Gcm:        util.AsBool(o.EncAlgoAes128Gcm, nil),
		EncAlgoAes256Cbc:        util.AsBool(o.EncAlgoAes256Cbc, nil),
		EncAlgoAes256Gcm:        util.AsBool(o.EncAlgoAes256Gcm, nil),
		EncAlgoChacha20Poly1305: util.AsBool(o.EncAlgoChacha20Poly1305, nil),
		EncAlgoRc4:              util.AsBool(o.EncAlgoRc4, nil),
		KeyxchgAlgoDhe:          util.AsBool(o.KeyxchgAlgoDhe, nil),
		KeyxchgAlgoEcdhe:        util.AsBool(o.KeyxchgAlgoEcdhe, nil),
		KeyxchgAlgoRsa:          util.AsBool(o.KeyxchgAlgoRsa, nil),
		MaxVersion:              o.MaxVersion,
		MinVersion:              o.MinVersion,
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DisableOverride = s.DisableOverride
	o.ForwardedOnly = util.YesNo(s.ForwardedOnly, nil)
	o.Interface = s.Interface
	if s.SshProxy != nil {
		var obj sshProxyXml_11_0_2
		obj.MarshalFromObject(*s.SshProxy)
		o.SshProxy = &obj
	}
	if s.SslForwardProxy != nil {
		var obj sslForwardProxyXml_11_0_2
		obj.MarshalFromObject(*s.SslForwardProxy)
		o.SslForwardProxy = &obj
	}
	if s.SslInboundProxy != nil {
		var obj sslInboundProxyXml_11_0_2
		obj.MarshalFromObject(*s.SslInboundProxy)
		o.SslInboundProxy = &obj
	}
	if s.SslNoProxy != nil {
		var obj sslNoProxyXml_11_0_2
		obj.MarshalFromObject(*s.SslNoProxy)
		o.SslNoProxy = &obj
	}
	if s.SslProtocolSettings != nil {
		var obj sslProtocolSettingsXml_11_0_2
		obj.MarshalFromObject(*s.SslProtocolSettings)
		o.SslProtocolSettings = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var sshProxyVal *SshProxy
	if o.SshProxy != nil {
		obj, err := o.SshProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sshProxyVal = obj
	}
	var sslForwardProxyVal *SslForwardProxy
	if o.SslForwardProxy != nil {
		obj, err := o.SslForwardProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslForwardProxyVal = obj
	}
	var sslInboundProxyVal *SslInboundProxy
	if o.SslInboundProxy != nil {
		obj, err := o.SslInboundProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslInboundProxyVal = obj
	}
	var sslNoProxyVal *SslNoProxy
	if o.SslNoProxy != nil {
		obj, err := o.SslNoProxy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslNoProxyVal = obj
	}
	var sslProtocolSettingsVal *SslProtocolSettings
	if o.SslProtocolSettings != nil {
		obj, err := o.SslProtocolSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslProtocolSettingsVal = obj
	}

	result := &Entry{
		Name:                o.Name,
		DisableOverride:     o.DisableOverride,
		ForwardedOnly:       util.AsBool(o.ForwardedOnly, nil),
		Interface:           o.Interface,
		SshProxy:            sshProxyVal,
		SslForwardProxy:     sslForwardProxyVal,
		SslInboundProxy:     sslInboundProxyVal,
		SslNoProxy:          sslNoProxyVal,
		SslProtocolSettings: sslProtocolSettingsVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *sshProxyXml_11_0_2) MarshalFromObject(s SshProxy) {
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockSshErrors = util.YesNo(s.BlockSshErrors, nil)
	o.BlockUnsupportedAlg = util.YesNo(s.BlockUnsupportedAlg, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sshProxyXml_11_0_2) UnmarshalToObject() (*SshProxy, error) {

	result := &SshProxy{
		BlockIfNoResource:       util.AsBool(o.BlockIfNoResource, nil),
		BlockSshErrors:          util.AsBool(o.BlockSshErrors, nil),
		BlockUnsupportedAlg:     util.AsBool(o.BlockUnsupportedAlg, nil),
		BlockUnsupportedVersion: util.AsBool(o.BlockUnsupportedVersion, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *sslForwardProxyXml_11_0_2) MarshalFromObject(s SslForwardProxy) {
	o.AutoIncludeAltname = util.YesNo(s.AutoIncludeAltname, nil)
	o.BlockClientCert = util.YesNo(s.BlockClientCert, nil)
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockIfHsmUnavailable = util.YesNo(s.BlockIfHsmUnavailable, nil)
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockIfSniMismatch = util.YesNo(s.BlockIfSniMismatch, nil)
	o.BlockTimeoutCert = util.YesNo(s.BlockTimeoutCert, nil)
	o.BlockTls13DowngradeNoResource = util.YesNo(s.BlockTls13DowngradeNoResource, nil)
	o.BlockUnknownCert = util.YesNo(s.BlockUnknownCert, nil)
	o.BlockUnsupportedCipher = util.YesNo(s.BlockUnsupportedCipher, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.BlockUntrustedIssuer = util.YesNo(s.BlockUntrustedIssuer, nil)
	o.RestrictCertExts = util.YesNo(s.RestrictCertExts, nil)
	o.StripAlpn = util.YesNo(s.StripAlpn, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslForwardProxyXml_11_0_2) UnmarshalToObject() (*SslForwardProxy, error) {

	result := &SslForwardProxy{
		AutoIncludeAltname:            util.AsBool(o.AutoIncludeAltname, nil),
		BlockClientCert:               util.AsBool(o.BlockClientCert, nil),
		BlockExpiredCertificate:       util.AsBool(o.BlockExpiredCertificate, nil),
		BlockIfHsmUnavailable:         util.AsBool(o.BlockIfHsmUnavailable, nil),
		BlockIfNoResource:             util.AsBool(o.BlockIfNoResource, nil),
		BlockIfSniMismatch:            util.AsBool(o.BlockIfSniMismatch, nil),
		BlockTimeoutCert:              util.AsBool(o.BlockTimeoutCert, nil),
		BlockTls13DowngradeNoResource: util.AsBool(o.BlockTls13DowngradeNoResource, nil),
		BlockUnknownCert:              util.AsBool(o.BlockUnknownCert, nil),
		BlockUnsupportedCipher:        util.AsBool(o.BlockUnsupportedCipher, nil),
		BlockUnsupportedVersion:       util.AsBool(o.BlockUnsupportedVersion, nil),
		BlockUntrustedIssuer:          util.AsBool(o.BlockUntrustedIssuer, nil),
		RestrictCertExts:              util.AsBool(o.RestrictCertExts, nil),
		StripAlpn:                     util.AsBool(o.StripAlpn, nil),
		Misc:                          o.Misc,
		MiscAttributes:                o.MiscAttributes,
	}
	return result, nil
}
func (o *sslInboundProxyXml_11_0_2) MarshalFromObject(s SslInboundProxy) {
	o.BlockIfHsmUnavailable = util.YesNo(s.BlockIfHsmUnavailable, nil)
	o.BlockIfNoResource = util.YesNo(s.BlockIfNoResource, nil)
	o.BlockTls13DowngradeNoResource = util.YesNo(s.BlockTls13DowngradeNoResource, nil)
	o.BlockUnsupportedCipher = util.YesNo(s.BlockUnsupportedCipher, nil)
	o.BlockUnsupportedVersion = util.YesNo(s.BlockUnsupportedVersion, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslInboundProxyXml_11_0_2) UnmarshalToObject() (*SslInboundProxy, error) {

	result := &SslInboundProxy{
		BlockIfHsmUnavailable:         util.AsBool(o.BlockIfHsmUnavailable, nil),
		BlockIfNoResource:             util.AsBool(o.BlockIfNoResource, nil),
		BlockTls13DowngradeNoResource: util.AsBool(o.BlockTls13DowngradeNoResource, nil),
		BlockUnsupportedCipher:        util.AsBool(o.BlockUnsupportedCipher, nil),
		BlockUnsupportedVersion:       util.AsBool(o.BlockUnsupportedVersion, nil),
		Misc:                          o.Misc,
		MiscAttributes:                o.MiscAttributes,
	}
	return result, nil
}
func (o *sslNoProxyXml_11_0_2) MarshalFromObject(s SslNoProxy) {
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockUntrustedIssuer = util.YesNo(s.BlockUntrustedIssuer, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslNoProxyXml_11_0_2) UnmarshalToObject() (*SslNoProxy, error) {

	result := &SslNoProxy{
		BlockExpiredCertificate: util.AsBool(o.BlockExpiredCertificate, nil),
		BlockUntrustedIssuer:    util.AsBool(o.BlockUntrustedIssuer, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *sslProtocolSettingsXml_11_0_2) MarshalFromObject(s SslProtocolSettings) {
	o.AuthAlgoMd5 = util.YesNo(s.AuthAlgoMd5, nil)
	o.AuthAlgoSha1 = util.YesNo(s.AuthAlgoSha1, nil)
	o.AuthAlgoSha256 = util.YesNo(s.AuthAlgoSha256, nil)
	o.AuthAlgoSha384 = util.YesNo(s.AuthAlgoSha384, nil)
	o.EncAlgo3des = util.YesNo(s.EncAlgo3des, nil)
	o.EncAlgoAes128Cbc = util.YesNo(s.EncAlgoAes128Cbc, nil)
	o.EncAlgoAes128Gcm = util.YesNo(s.EncAlgoAes128Gcm, nil)
	o.EncAlgoAes256Cbc = util.YesNo(s.EncAlgoAes256Cbc, nil)
	o.EncAlgoAes256Gcm = util.YesNo(s.EncAlgoAes256Gcm, nil)
	o.EncAlgoChacha20Poly1305 = util.YesNo(s.EncAlgoChacha20Poly1305, nil)
	o.EncAlgoRc4 = util.YesNo(s.EncAlgoRc4, nil)
	o.KeyxchgAlgoDhe = util.YesNo(s.KeyxchgAlgoDhe, nil)
	o.KeyxchgAlgoEcdhe = util.YesNo(s.KeyxchgAlgoEcdhe, nil)
	o.KeyxchgAlgoRsa = util.YesNo(s.KeyxchgAlgoRsa, nil)
	o.MaxVersion = s.MaxVersion
	o.MinVersion = s.MinVersion
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o sslProtocolSettingsXml_11_0_2) UnmarshalToObject() (*SslProtocolSettings, error) {

	result := &SslProtocolSettings{
		AuthAlgoMd5:             util.AsBool(o.AuthAlgoMd5, nil),
		AuthAlgoSha1:            util.AsBool(o.AuthAlgoSha1, nil),
		AuthAlgoSha256:          util.AsBool(o.AuthAlgoSha256, nil),
		AuthAlgoSha384:          util.AsBool(o.AuthAlgoSha384, nil),
		EncAlgo3des:             util.AsBool(o.EncAlgo3des, nil),
		EncAlgoAes128Cbc:        util.AsBool(o.EncAlgoAes128Cbc, nil),
		EncAlgoAes128Gcm:        util.AsBool(o.EncAlgoAes128Gcm, nil),
		EncAlgoAes256Cbc:        util.AsBool(o.EncAlgoAes256Cbc, nil),
		EncAlgoAes256Gcm:        util.AsBool(o.EncAlgoAes256Gcm, nil),
		EncAlgoChacha20Poly1305: util.AsBool(o.EncAlgoChacha20Poly1305, nil),
		EncAlgoRc4:              util.AsBool(o.EncAlgoRc4, nil),
		KeyxchgAlgoDhe:          util.AsBool(o.KeyxchgAlgoDhe, nil),
		KeyxchgAlgoEcdhe:        util.AsBool(o.KeyxchgAlgoEcdhe, nil),
		KeyxchgAlgoRsa:          util.AsBool(o.KeyxchgAlgoRsa, nil),
		MaxVersion:              o.MaxVersion,
		MinVersion:              o.MinVersion,
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "forwarded_only" || v == "ForwardedOnly" {
		return e.ForwardedOnly, nil
	}
	if v == "interface" || v == "Interface" {
		return e.Interface, nil
	}
	if v == "ssh_proxy" || v == "SshProxy" {
		return e.SshProxy, nil
	}
	if v == "ssl_forward_proxy" || v == "SslForwardProxy" {
		return e.SslForwardProxy, nil
	}
	if v == "ssl_inbound_proxy" || v == "SslInboundProxy" {
		return e.SslInboundProxy, nil
	}
	if v == "ssl_no_proxy" || v == "SslNoProxy" {
		return e.SslNoProxy, nil
	}
	if v == "ssl_protocol_settings" || v == "SslProtocolSettings" {
		return e.SslProtocolSettings, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

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
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.BoolsMatch(o.ForwardedOnly, other.ForwardedOnly) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !o.SshProxy.matches(other.SshProxy) {
		return false
	}
	if !o.SslForwardProxy.matches(other.SslForwardProxy) {
		return false
	}
	if !o.SslInboundProxy.matches(other.SslInboundProxy) {
		return false
	}
	if !o.SslNoProxy.matches(other.SslNoProxy) {
		return false
	}
	if !o.SslProtocolSettings.matches(other.SslProtocolSettings) {
		return false
	}

	return true
}

func (o *SshProxy) matches(other *SshProxy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfNoResource, other.BlockIfNoResource) {
		return false
	}
	if !util.BoolsMatch(o.BlockSshErrors, other.BlockSshErrors) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedAlg, other.BlockUnsupportedAlg) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedVersion, other.BlockUnsupportedVersion) {
		return false
	}

	return true
}

func (o *SslForwardProxy) matches(other *SslForwardProxy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AutoIncludeAltname, other.AutoIncludeAltname) {
		return false
	}
	if !util.BoolsMatch(o.BlockClientCert, other.BlockClientCert) {
		return false
	}
	if !util.BoolsMatch(o.BlockExpiredCertificate, other.BlockExpiredCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfHsmUnavailable, other.BlockIfHsmUnavailable) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfNoResource, other.BlockIfNoResource) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfSniMismatch, other.BlockIfSniMismatch) {
		return false
	}
	if !util.BoolsMatch(o.BlockTimeoutCert, other.BlockTimeoutCert) {
		return false
	}
	if !util.BoolsMatch(o.BlockTls13DowngradeNoResource, other.BlockTls13DowngradeNoResource) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnknownCert, other.BlockUnknownCert) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedCipher, other.BlockUnsupportedCipher) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedVersion, other.BlockUnsupportedVersion) {
		return false
	}
	if !util.BoolsMatch(o.BlockUntrustedIssuer, other.BlockUntrustedIssuer) {
		return false
	}
	if !util.BoolsMatch(o.RestrictCertExts, other.RestrictCertExts) {
		return false
	}
	if !util.BoolsMatch(o.StripAlpn, other.StripAlpn) {
		return false
	}

	return true
}

func (o *SslInboundProxy) matches(other *SslInboundProxy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfHsmUnavailable, other.BlockIfHsmUnavailable) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfNoResource, other.BlockIfNoResource) {
		return false
	}
	if !util.BoolsMatch(o.BlockTls13DowngradeNoResource, other.BlockTls13DowngradeNoResource) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedCipher, other.BlockUnsupportedCipher) {
		return false
	}
	if !util.BoolsMatch(o.BlockUnsupportedVersion, other.BlockUnsupportedVersion) {
		return false
	}

	return true
}

func (o *SslNoProxy) matches(other *SslNoProxy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.BlockExpiredCertificate, other.BlockExpiredCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockUntrustedIssuer, other.BlockUntrustedIssuer) {
		return false
	}

	return true
}

func (o *SslProtocolSettings) matches(other *SslProtocolSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.AuthAlgoMd5, other.AuthAlgoMd5) {
		return false
	}
	if !util.BoolsMatch(o.AuthAlgoSha1, other.AuthAlgoSha1) {
		return false
	}
	if !util.BoolsMatch(o.AuthAlgoSha256, other.AuthAlgoSha256) {
		return false
	}
	if !util.BoolsMatch(o.AuthAlgoSha384, other.AuthAlgoSha384) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgo3des, other.EncAlgo3des) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoAes128Cbc, other.EncAlgoAes128Cbc) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoAes128Gcm, other.EncAlgoAes128Gcm) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoAes256Cbc, other.EncAlgoAes256Cbc) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoAes256Gcm, other.EncAlgoAes256Gcm) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoChacha20Poly1305, other.EncAlgoChacha20Poly1305) {
		return false
	}
	if !util.BoolsMatch(o.EncAlgoRc4, other.EncAlgoRc4) {
		return false
	}
	if !util.BoolsMatch(o.KeyxchgAlgoDhe, other.KeyxchgAlgoDhe) {
		return false
	}
	if !util.BoolsMatch(o.KeyxchgAlgoEcdhe, other.KeyxchgAlgoEcdhe) {
		return false
	}
	if !util.BoolsMatch(o.KeyxchgAlgoRsa, other.KeyxchgAlgoRsa) {
		return false
	}
	if !util.StringsMatch(o.MaxVersion, other.MaxVersion) {
		return false
	}
	if !util.StringsMatch(o.MinVersion, other.MinVersion) {
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
