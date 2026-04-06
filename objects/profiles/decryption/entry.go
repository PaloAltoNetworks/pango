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
	Suffix = []string{"decryption"}
)

// Entry represents a PAN-OS decryption profile.
type Entry struct {
	Name                string
	Description         *string
	SslForwardProxy     *SslForwardProxy
	SslInboundInspection *SslInboundInspection
	SslNoProxy          *SslNoProxy
	SslProtocolSettings *SslProtocolSettings

	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type SslForwardProxy struct {
	AutoIncludeAltname             *bool
	BlockClientCert                *bool
	BlockExpiredCertificate        *bool
	BlockTimeoutCert               *bool
	BlockTls13DowngradeNoResource  *bool
	BlockUnknownCert               *bool
	BlockUnsupportedCipher         *bool
	BlockUnsupportedVersion        *bool
	BlockUntrustedIssuer           *bool
	RestrictCertExts               *bool
	StripAlpn                      *bool
}

type SslInboundInspection struct {
	BlockIfHsmUnavailable   *bool
	BlockIfNoResource       *bool
	BlockUnsupportedCipher  *bool
	BlockUnsupportedVersion *bool
}

type SslNoProxy struct {
	BlockClientCert         *bool
	BlockExpiredCertificate *bool
	BlockTimeoutCert        *bool
	BlockUnknownCert        *bool
	BlockUnsupportedCipher  *bool
	BlockUnsupportedVersion *bool
	BlockUntrustedIssuer    *bool
}

type SslProtocolSettings struct {
	AuthAlgoMd5        *bool
	AuthAlgoSha1       *bool
	AuthAlgoSha256     *bool
	AuthAlgoSha384     *bool
	EncAlgo3des        *bool
	EncAlgoAes128Cbc   *bool
	EncAlgoAes128Gcm   *bool
	EncAlgoAes256Cbc   *bool
	EncAlgoAes256Gcm   *bool
	EncAlgoRc4         *bool
	KeyxchgAlgoDhe     *bool
	KeyxchgAlgoEcdhe   *bool
	KeyxchgAlgoRsa     *bool
	MaxVersion         *string
	MinVersion         *string
}

// EntryObject interface methods.

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

// XML types.

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName              xml.Name                  `xml:"entry"`
	Name                 string                    `xml:"name,attr"`
	Description          *string                   `xml:"description,omitempty"`
	SslForwardProxy      *sslForwardProxyXml       `xml:"ssl-forward-proxy,omitempty"`
	SslInboundInspection *sslInboundInspectionXml  `xml:"ssl-inbound-inspection,omitempty"`
	SslNoProxy           *sslNoProxyXml            `xml:"ssl-no-proxy,omitempty"`
	SslProtocolSettings  *sslProtocolSettingsXml   `xml:"ssl-protocol-settings,omitempty"`
	Misc                 []generic.Xml             `xml:",any"`
	MiscAttributes       []xml.Attr                `xml:",any,attr"`
}

type sslForwardProxyXml struct {
	AutoIncludeAltname            *string `xml:"auto-include-altname,omitempty"`
	BlockClientCert               *string `xml:"block-client-cert,omitempty"`
	BlockExpiredCertificate       *string `xml:"block-expired-certificate,omitempty"`
	BlockTimeoutCert              *string `xml:"block-timeout-cert,omitempty"`
	BlockTls13DowngradeNoResource *string `xml:"block-tls13-downgrade-no-resource,omitempty"`
	BlockUnknownCert              *string `xml:"block-unknown-cert,omitempty"`
	BlockUnsupportedCipher        *string `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion       *string `xml:"block-unsupported-version,omitempty"`
	BlockUntrustedIssuer          *string `xml:"block-untrusted-issuer,omitempty"`
	RestrictCertExts              *string `xml:"restrict-cert-exts,omitempty"`
	StripAlpn                     *string `xml:"strip-alpn,omitempty"`
}

type sslInboundInspectionXml struct {
	BlockIfHsmUnavailable   *string `xml:"block-if-hsm-unavailable,omitempty"`
	BlockIfNoResource       *string `xml:"block-if-no-resource,omitempty"`
	BlockUnsupportedCipher  *string `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion *string `xml:"block-unsupported-version,omitempty"`
}

type sslNoProxyXml struct {
	BlockClientCert         *string `xml:"block-client-cert,omitempty"`
	BlockExpiredCertificate *string `xml:"block-expired-certificate,omitempty"`
	BlockTimeoutCert        *string `xml:"block-timeout-cert,omitempty"`
	BlockUnknownCert        *string `xml:"block-unknown-cert,omitempty"`
	BlockUnsupportedCipher  *string `xml:"block-unsupported-cipher,omitempty"`
	BlockUnsupportedVersion *string `xml:"block-unsupported-version,omitempty"`
	BlockUntrustedIssuer    *string `xml:"block-untrusted-issuer,omitempty"`
}

type sslProtocolSettingsXml struct {
	AuthAlgoMd5        *string `xml:"auth-algo-md5,omitempty"`
	AuthAlgoSha1       *string `xml:"auth-algo-sha1,omitempty"`
	AuthAlgoSha256     *string `xml:"auth-algo-sha256,omitempty"`
	AuthAlgoSha384     *string `xml:"auth-algo-sha384,omitempty"`
	EncAlgo3des        *string `xml:"enc-algo-3des,omitempty"`
	EncAlgoAes128Cbc   *string `xml:"enc-algo-aes-128-cbc,omitempty"`
	EncAlgoAes128Gcm   *string `xml:"enc-algo-aes-128-gcm,omitempty"`
	EncAlgoAes256Cbc   *string `xml:"enc-algo-aes-256-cbc,omitempty"`
	EncAlgoAes256Gcm   *string `xml:"enc-algo-aes-256-gcm,omitempty"`
	EncAlgoRc4         *string `xml:"enc-algo-rc4,omitempty"`
	KeyxchgAlgoDhe     *string `xml:"keyxchg-algo-dhe,omitempty"`
	KeyxchgAlgoEcdhe   *string `xml:"keyxchg-algo-ecdhe,omitempty"`
	KeyxchgAlgoRsa     *string `xml:"keyxchg-algo-rsa,omitempty"`
	MaxVersion         *string `xml:"max-version,omitempty"`
	MinVersion         *string `xml:"min-version,omitempty"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.unmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}
	return entries, nil
}

func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{
		Name:           o.Name,
		Description:    o.Description,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}

	if o.SslForwardProxy != nil {
		entry.SslForwardProxy = &sslForwardProxyXml{
			AutoIncludeAltname:            util.YesNo(o.SslForwardProxy.AutoIncludeAltname, nil),
			BlockClientCert:               util.YesNo(o.SslForwardProxy.BlockClientCert, nil),
			BlockExpiredCertificate:       util.YesNo(o.SslForwardProxy.BlockExpiredCertificate, nil),
			BlockTimeoutCert:              util.YesNo(o.SslForwardProxy.BlockTimeoutCert, nil),
			BlockTls13DowngradeNoResource: util.YesNo(o.SslForwardProxy.BlockTls13DowngradeNoResource, nil),
			BlockUnknownCert:              util.YesNo(o.SslForwardProxy.BlockUnknownCert, nil),
			BlockUnsupportedCipher:        util.YesNo(o.SslForwardProxy.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion:       util.YesNo(o.SslForwardProxy.BlockUnsupportedVersion, nil),
			BlockUntrustedIssuer:          util.YesNo(o.SslForwardProxy.BlockUntrustedIssuer, nil),
			RestrictCertExts:              util.YesNo(o.SslForwardProxy.RestrictCertExts, nil),
			StripAlpn:                     util.YesNo(o.SslForwardProxy.StripAlpn, nil),
		}
	}

	if o.SslInboundInspection != nil {
		entry.SslInboundInspection = &sslInboundInspectionXml{
			BlockIfHsmUnavailable:   util.YesNo(o.SslInboundInspection.BlockIfHsmUnavailable, nil),
			BlockIfNoResource:       util.YesNo(o.SslInboundInspection.BlockIfNoResource, nil),
			BlockUnsupportedCipher:  util.YesNo(o.SslInboundInspection.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion: util.YesNo(o.SslInboundInspection.BlockUnsupportedVersion, nil),
		}
	}

	if o.SslNoProxy != nil {
		entry.SslNoProxy = &sslNoProxyXml{
			BlockClientCert:         util.YesNo(o.SslNoProxy.BlockClientCert, nil),
			BlockExpiredCertificate: util.YesNo(o.SslNoProxy.BlockExpiredCertificate, nil),
			BlockTimeoutCert:        util.YesNo(o.SslNoProxy.BlockTimeoutCert, nil),
			BlockUnknownCert:        util.YesNo(o.SslNoProxy.BlockUnknownCert, nil),
			BlockUnsupportedCipher:  util.YesNo(o.SslNoProxy.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion: util.YesNo(o.SslNoProxy.BlockUnsupportedVersion, nil),
			BlockUntrustedIssuer:    util.YesNo(o.SslNoProxy.BlockUntrustedIssuer, nil),
		}
	}

	if o.SslProtocolSettings != nil {
		entry.SslProtocolSettings = &sslProtocolSettingsXml{
			AuthAlgoMd5:        util.YesNo(o.SslProtocolSettings.AuthAlgoMd5, nil),
			AuthAlgoSha1:       util.YesNo(o.SslProtocolSettings.AuthAlgoSha1, nil),
			AuthAlgoSha256:     util.YesNo(o.SslProtocolSettings.AuthAlgoSha256, nil),
			AuthAlgoSha384:     util.YesNo(o.SslProtocolSettings.AuthAlgoSha384, nil),
			EncAlgo3des:        util.YesNo(o.SslProtocolSettings.EncAlgo3des, nil),
			EncAlgoAes128Cbc:   util.YesNo(o.SslProtocolSettings.EncAlgoAes128Cbc, nil),
			EncAlgoAes128Gcm:   util.YesNo(o.SslProtocolSettings.EncAlgoAes128Gcm, nil),
			EncAlgoAes256Cbc:   util.YesNo(o.SslProtocolSettings.EncAlgoAes256Cbc, nil),
			EncAlgoAes256Gcm:   util.YesNo(o.SslProtocolSettings.EncAlgoAes256Gcm, nil),
			EncAlgoRc4:         util.YesNo(o.SslProtocolSettings.EncAlgoRc4, nil),
			KeyxchgAlgoDhe:     util.YesNo(o.SslProtocolSettings.KeyxchgAlgoDhe, nil),
			KeyxchgAlgoEcdhe:   util.YesNo(o.SslProtocolSettings.KeyxchgAlgoEcdhe, nil),
			KeyxchgAlgoRsa:     util.YesNo(o.SslProtocolSettings.KeyxchgAlgoRsa, nil),
			MaxVersion:         o.SslProtocolSettings.MaxVersion,
			MinVersion:         o.SslProtocolSettings.MinVersion,
		}
	}

	return entry, nil
}

func (o entryXml) unmarshalToObject() (*Entry, error) {
	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}

	if o.SslForwardProxy != nil {
		result.SslForwardProxy = &SslForwardProxy{
			AutoIncludeAltname:            util.AsBool(o.SslForwardProxy.AutoIncludeAltname, nil),
			BlockClientCert:               util.AsBool(o.SslForwardProxy.BlockClientCert, nil),
			BlockExpiredCertificate:       util.AsBool(o.SslForwardProxy.BlockExpiredCertificate, nil),
			BlockTimeoutCert:              util.AsBool(o.SslForwardProxy.BlockTimeoutCert, nil),
			BlockTls13DowngradeNoResource: util.AsBool(o.SslForwardProxy.BlockTls13DowngradeNoResource, nil),
			BlockUnknownCert:              util.AsBool(o.SslForwardProxy.BlockUnknownCert, nil),
			BlockUnsupportedCipher:        util.AsBool(o.SslForwardProxy.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion:       util.AsBool(o.SslForwardProxy.BlockUnsupportedVersion, nil),
			BlockUntrustedIssuer:          util.AsBool(o.SslForwardProxy.BlockUntrustedIssuer, nil),
			RestrictCertExts:              util.AsBool(o.SslForwardProxy.RestrictCertExts, nil),
			StripAlpn:                     util.AsBool(o.SslForwardProxy.StripAlpn, nil),
		}
	}

	if o.SslInboundInspection != nil {
		result.SslInboundInspection = &SslInboundInspection{
			BlockIfHsmUnavailable:   util.AsBool(o.SslInboundInspection.BlockIfHsmUnavailable, nil),
			BlockIfNoResource:       util.AsBool(o.SslInboundInspection.BlockIfNoResource, nil),
			BlockUnsupportedCipher:  util.AsBool(o.SslInboundInspection.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion: util.AsBool(o.SslInboundInspection.BlockUnsupportedVersion, nil),
		}
	}

	if o.SslNoProxy != nil {
		result.SslNoProxy = &SslNoProxy{
			BlockClientCert:         util.AsBool(o.SslNoProxy.BlockClientCert, nil),
			BlockExpiredCertificate: util.AsBool(o.SslNoProxy.BlockExpiredCertificate, nil),
			BlockTimeoutCert:        util.AsBool(o.SslNoProxy.BlockTimeoutCert, nil),
			BlockUnknownCert:        util.AsBool(o.SslNoProxy.BlockUnknownCert, nil),
			BlockUnsupportedCipher:  util.AsBool(o.SslNoProxy.BlockUnsupportedCipher, nil),
			BlockUnsupportedVersion: util.AsBool(o.SslNoProxy.BlockUnsupportedVersion, nil),
			BlockUntrustedIssuer:    util.AsBool(o.SslNoProxy.BlockUntrustedIssuer, nil),
		}
	}

	if o.SslProtocolSettings != nil {
		result.SslProtocolSettings = &SslProtocolSettings{
			AuthAlgoMd5:        util.AsBool(o.SslProtocolSettings.AuthAlgoMd5, nil),
			AuthAlgoSha1:       util.AsBool(o.SslProtocolSettings.AuthAlgoSha1, nil),
			AuthAlgoSha256:     util.AsBool(o.SslProtocolSettings.AuthAlgoSha256, nil),
			AuthAlgoSha384:     util.AsBool(o.SslProtocolSettings.AuthAlgoSha384, nil),
			EncAlgo3des:        util.AsBool(o.SslProtocolSettings.EncAlgo3des, nil),
			EncAlgoAes128Cbc:   util.AsBool(o.SslProtocolSettings.EncAlgoAes128Cbc, nil),
			EncAlgoAes128Gcm:   util.AsBool(o.SslProtocolSettings.EncAlgoAes128Gcm, nil),
			EncAlgoAes256Cbc:   util.AsBool(o.SslProtocolSettings.EncAlgoAes256Cbc, nil),
			EncAlgoAes256Gcm:   util.AsBool(o.SslProtocolSettings.EncAlgoAes256Gcm, nil),
			EncAlgoRc4:         util.AsBool(o.SslProtocolSettings.EncAlgoRc4, nil),
			KeyxchgAlgoDhe:     util.AsBool(o.SslProtocolSettings.KeyxchgAlgoDhe, nil),
			KeyxchgAlgoEcdhe:   util.AsBool(o.SslProtocolSettings.KeyxchgAlgoEcdhe, nil),
			KeyxchgAlgoRsa:     util.AsBool(o.SslProtocolSettings.KeyxchgAlgoRsa, nil),
			MaxVersion:         o.SslProtocolSettings.MaxVersion,
			MinVersion:         o.SslProtocolSettings.MinVersion,
		}
	}

	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "ssl_forward_proxy" || v == "SslForwardProxy" {
		return e.SslForwardProxy, nil
	}
	if v == "ssl_inbound_inspection" || v == "SslInboundInspection" {
		return e.SslInboundInspection, nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !o.SslForwardProxy.matches(other.SslForwardProxy) {
		return false
	}
	if !o.SslInboundInspection.matches(other.SslInboundInspection) {
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

func (o *SslForwardProxy) matches(other *SslForwardProxy) bool {
	if o == nil && other == nil {
		return true
	}
	if (o == nil) != (other == nil) {
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

func (o *SslInboundInspection) matches(other *SslInboundInspection) bool {
	if o == nil && other == nil {
		return true
	}
	if (o == nil) != (other == nil) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfHsmUnavailable, other.BlockIfHsmUnavailable) {
		return false
	}
	if !util.BoolsMatch(o.BlockIfNoResource, other.BlockIfNoResource) {
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
	if (o == nil) != (other == nil) {
		return false
	}
	if !util.BoolsMatch(o.BlockClientCert, other.BlockClientCert) {
		return false
	}
	if !util.BoolsMatch(o.BlockExpiredCertificate, other.BlockExpiredCertificate) {
		return false
	}
	if !util.BoolsMatch(o.BlockTimeoutCert, other.BlockTimeoutCert) {
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
	return true
}

func (o *SslProtocolSettings) matches(other *SslProtocolSettings) bool {
	if o == nil && other == nil {
		return true
	}
	if (o == nil) != (other == nil) {
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
