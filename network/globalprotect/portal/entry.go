package portal

import (
	"encoding/xml"
	"errors"
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
	suffix = []string{"global-protect", "global-protect-portal", "$name"}
)

type Entry struct {
	Name            string
	ClientConfig    *ClientConfig
	ClientlessVpn   *ClientlessVpn
	PortalConfig    *PortalConfig
	SatelliteConfig *SatelliteConfig
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type ClientConfig struct {
	AgentUserOverrideKey *string
	Configs              []ClientConfigConfigs
	RootCa               []ClientConfigRootCa
	Misc                 []generic.Xml
	MiscAttributes       []xml.Attr
}
type ClientConfigConfigs struct {
	Name                             string
	SaveUserCredentials              *string
	Portal2fa                        *bool
	InternalGateway2fa               *bool
	AutoDiscoveryExternalGateway2fa  *bool
	ManualOnlyGateway2fa             *bool
	RefreshConfig                    *bool
	MdmAddress                       *string
	MdmEnrollmentPort                *string
	SourceUser                       []string
	ThirdPartyVpnClients             []string
	Os                               []string
	Certificate                      *ClientConfigConfigsCertificate
	CustomChecks                     *ClientConfigConfigsCustomChecks
	Gateways                         *ClientConfigConfigsGateways
	InternalHostDetection            *ClientConfigConfigsInternalHostDetection
	InternalHostDetectionV6          *ClientConfigConfigsInternalHostDetectionV6
	AgentUi                          *ClientConfigConfigsAgentUi
	HipCollection                    *ClientConfigConfigsHipCollection
	AgentConfig                      *ClientConfigConfigsAgentConfig
	GpAppConfig                      *ClientConfigConfigsGpAppConfig
	AuthenticationOverride           *ClientConfigConfigsAuthenticationOverride
	MachineAccountExistsWithSerialno *ClientConfigConfigsMachineAccountExistsWithSerialno
	ClientCertificate                *ClientConfigConfigsClientCertificate
	Misc                             []generic.Xml
	MiscAttributes                   []xml.Attr
}
type ClientConfigConfigsCertificate struct {
	Criteria       *ClientConfigConfigsCertificateCriteria
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsCertificateCriteria struct {
	CertificateProfile *string
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type ClientConfigConfigsCustomChecks struct {
	Criteria       *ClientConfigConfigsCustomChecksCriteria
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsCustomChecksCriteria struct {
	RegistryKey    []ClientConfigConfigsCustomChecksCriteriaRegistryKey
	Plist          []ClientConfigConfigsCustomChecksCriteriaPlist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsCustomChecksCriteriaRegistryKey struct {
	Name             string
	DefaultValueData *string
	Negate           *bool
	RegistryValue    []ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
}
type ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue struct {
	Name           string
	ValueData      *string
	Negate         *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsCustomChecksCriteriaPlist struct {
	Name           string
	Negate         *bool
	Key            []ClientConfigConfigsCustomChecksCriteriaPlistKey
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsCustomChecksCriteriaPlistKey struct {
	Name           string
	Value          *string
	Negate         *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGateways struct {
	Internal       *ClientConfigConfigsGatewaysInternal
	External       *ClientConfigConfigsGatewaysExternal
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysInternal struct {
	List           []ClientConfigConfigsGatewaysInternalList
	DhcpOptionCode []int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysInternalList struct {
	Name           string
	SourceIp       []string
	Fqdn           *string
	Ip             *ClientConfigConfigsGatewaysInternalListIp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysInternalListIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysExternal struct {
	CutoffTime     *int64
	List           []ClientConfigConfigsGatewaysExternalList
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysExternalList struct {
	Name           string
	PriorityRule   []ClientConfigConfigsGatewaysExternalListPriorityRule
	Manual         *bool
	Fqdn           *string
	Ip             *ClientConfigConfigsGatewaysExternalListIp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysExternalListPriorityRule struct {
	Name           string
	Priority       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGatewaysExternalListIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsInternalHostDetection struct {
	IpAddress      *string
	Hostname       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsInternalHostDetectionV6 struct {
	IpAddress      *string
	Hostname       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsAgentUi struct {
	Passcode                 *string
	UninstallPassword        *string
	AgentUserOverrideTimeout *int64
	MaxAgentUserOverrides    *int64
	WelcomePage              *ClientConfigConfigsAgentUiWelcomePage
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type ClientConfigConfigsAgentUiWelcomePage struct {
	Page           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollection struct {
	CertificateProfile *string
	MaxWaitTime        *int64
	CollectHipData     *bool
	Exclusion          *ClientConfigConfigsHipCollectionExclusion
	CustomChecks       *ClientConfigConfigsHipCollectionCustomChecks
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type ClientConfigConfigsHipCollectionExclusion struct {
	Category       []ClientConfigConfigsHipCollectionExclusionCategory
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionExclusionCategory struct {
	Name           string
	Vendor         []ClientConfigConfigsHipCollectionExclusionCategoryVendor
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionExclusionCategoryVendor struct {
	Name           string
	Product        []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecks struct {
	Windows        *ClientConfigConfigsHipCollectionCustomChecksWindows
	MacOs          *ClientConfigConfigsHipCollectionCustomChecksMacOs
	Linux          *ClientConfigConfigsHipCollectionCustomChecksLinux
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecksWindows struct {
	RegistryKey    []ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey
	ProcessList    []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey struct {
	Name           string
	RegistryValue  []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecksMacOs struct {
	Plist          []ClientConfigConfigsHipCollectionCustomChecksMacOsPlist
	ProcessList    []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecksMacOsPlist struct {
	Name           string
	Key            []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsHipCollectionCustomChecksLinux struct {
	ProcessList    []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsAgentConfig struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGpAppConfig struct {
	Config         []ClientConfigConfigsGpAppConfigConfig
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsGpAppConfigConfig struct {
	Name           string
	Value          []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsAuthenticationOverride struct {
	GenerateCookie           *bool
	CookieEncryptDecryptCert *string
	AcceptCookie             *ClientConfigConfigsAuthenticationOverrideAcceptCookie
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type ClientConfigConfigsAuthenticationOverrideAcceptCookie struct {
	CookieLifetime *ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime struct {
	LifetimeInDays    *int64
	LifetimeInHours   *int64
	LifetimeInMinutes *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type ClientConfigConfigsMachineAccountExistsWithSerialno struct {
	No             *ClientConfigConfigsMachineAccountExistsWithSerialnoNo
	Yes            *ClientConfigConfigsMachineAccountExistsWithSerialnoYes
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsMachineAccountExistsWithSerialnoNo struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsMachineAccountExistsWithSerialnoYes struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigConfigsClientCertificate struct {
	Local          *string
	Scep           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientConfigRootCa struct {
	Name               string
	InstallInCertStore *bool
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type ClientlessVpn struct {
	AppsToUserMapping        []ClientlessVpnAppsToUserMapping
	CryptoSettings           *ClientlessVpnCryptoSettings
	DnsProxy                 *string
	Hostname                 *string
	InactivityLogout         *ClientlessVpnInactivityLogout
	LoginLifetime            *ClientlessVpnLoginLifetime
	MaxUser                  *int64
	ProxyServerSetting       []ClientlessVpnProxyServerSetting
	RewriteExcludeDomainList []string
	SecurityZone             *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type ClientlessVpnAppsToUserMapping struct {
	Name                                  string
	SourceUser                            []string
	Applications                          []string
	EnableCustomAppURLAddressBar          *bool
	DisplayGlobalProtectAgentDownloadLink *bool
	Misc                                  []generic.Xml
	MiscAttributes                        []xml.Attr
}
type ClientlessVpnCryptoSettings struct {
	ServerCertVerification *ClientlessVpnCryptoSettingsServerCertVerification
	SslProtocol            *ClientlessVpnCryptoSettingsSslProtocol
	Misc                   []generic.Xml
	MiscAttributes         []xml.Attr
}
type ClientlessVpnCryptoSettingsServerCertVerification struct {
	BlockExpiredCertificate *bool
	BlockTimeoutCert        *bool
	BlockUnknownCert        *bool
	BlockUntrustedIssuer    *bool
	Misc                    []generic.Xml
	MiscAttributes          []xml.Attr
}
type ClientlessVpnCryptoSettingsSslProtocol struct {
	AuthAlgoMd5      *bool
	AuthAlgoSha1     *bool
	AuthAlgoSha256   *bool
	AuthAlgoSha384   *bool
	EncAlgo3des      *bool
	EncAlgoAes128Cbc *bool
	EncAlgoAes128Gcm *bool
	EncAlgoAes256Cbc *bool
	EncAlgoAes256Gcm *bool
	EncAlgoRc4       *bool
	KeyxchgAlgoDhe   *bool
	KeyxchgAlgoEcdhe *bool
	KeyxchgAlgoRsa   *bool
	MaxVersion       *string
	MinVersion       *string
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
}
type ClientlessVpnInactivityLogout struct {
	Hours          *int64
	Minutes        *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientlessVpnLoginLifetime struct {
	Hours          *int64
	Minutes        *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientlessVpnProxyServerSetting struct {
	Name           string
	Domains        []string
	UseProxy       *bool
	ProxyServer    *ClientlessVpnProxyServerSettingProxyServer
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ClientlessVpnProxyServerSettingProxyServer struct {
	Server         *string
	Port           *int64
	User           *string
	Password       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfig struct {
	CertificateProfile   *string
	ClientAuth           []PortalConfigClientAuth
	ConfigSelection      *PortalConfigConfigSelection
	CustomHelpPage       *string
	CustomHomePage       *string
	CustomLoginPage      *string
	LocalAddress         *PortalConfigLocalAddress
	LogFail              *bool
	LogSetting           *string
	LogSuccess           *bool
	SslTlsServiceProfile *string
	Misc                 []generic.Xml
	MiscAttributes       []xml.Attr
}
type PortalConfigClientAuth struct {
	Name                               string
	Os                                 *string
	AuthenticationProfile              *string
	AutoRetrievePasscode               *bool
	UsernameLabel                      *string
	PasswordLabel                      *string
	AuthenticationMessage              *string
	UserCredentialOrClientCertRequired *string
	Misc                               []generic.Xml
	MiscAttributes                     []xml.Attr
}
type PortalConfigConfigSelection struct {
	CertificateProfile *string
	CustomChecks       *PortalConfigConfigSelectionCustomChecks
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type PortalConfigConfigSelectionCustomChecks struct {
	MacOs          *PortalConfigConfigSelectionCustomChecksMacOs
	Windows        *PortalConfigConfigSelectionCustomChecksWindows
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigConfigSelectionCustomChecksMacOs struct {
	Plist          []PortalConfigConfigSelectionCustomChecksMacOsPlist
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigConfigSelectionCustomChecksMacOsPlist struct {
	Name           string
	Key            []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigConfigSelectionCustomChecksWindows struct {
	RegistryKey    []PortalConfigConfigSelectionCustomChecksWindowsRegistryKey
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigConfigSelectionCustomChecksWindowsRegistryKey struct {
	Name           string
	RegistryValue  []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigLocalAddress struct {
	Interface       *string
	IpAddressFamily *string
	FloatingIp      *PortalConfigLocalAddressFloatingIp
	Ip              *PortalConfigLocalAddressIp
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type PortalConfigLocalAddressFloatingIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PortalConfigLocalAddressIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SatelliteConfig struct {
	ClientCertificate *SatelliteConfigClientCertificate
	Configs           []SatelliteConfigConfigs
	RootCa            []string
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type SatelliteConfigClientCertificate struct {
	Local          *SatelliteConfigClientCertificateLocal
	Scep           *SatelliteConfigClientCertificateScep
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SatelliteConfigClientCertificateLocal struct {
	CertificateLifeTime      *int64
	CertificateRenewalPeriod *int64
	IssuingCertificate       *string
	OcspResponder            *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type SatelliteConfigClientCertificateScep struct {
	CertificateRenewalPeriod *int64
	Scep                     *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type SatelliteConfigConfigs struct {
	Name                  string
	Devices               []string
	SourceUser            []string
	Gateways              []SatelliteConfigConfigsGateways
	ConfigRefreshInterval *int64
	Misc                  []generic.Xml
	MiscAttributes        []xml.Attr
}
type SatelliteConfigConfigsGateways struct {
	Name           string
	Ipv6Preferred  *bool
	Priority       *int64
	Fqdn           *string
	Ip             *SatelliteConfigConfigsGatewaysIp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SatelliteConfigConfigsGatewaysIp struct {
	Ipv4           *string
	Ipv6           *string
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
	XMLName         xml.Name            `xml:"entry"`
	Name            string              `xml:"name,attr"`
	ClientConfig    *clientConfigXml    `xml:"client-config,omitempty"`
	ClientlessVpn   *clientlessVpnXml   `xml:"clientless-vpn,omitempty"`
	PortalConfig    *portalConfigXml    `xml:"portal-config,omitempty"`
	SatelliteConfig *satelliteConfigXml `xml:"satellite-config,omitempty"`
	Misc            []generic.Xml       `xml:",any"`
	MiscAttributes  []xml.Attr          `xml:",any,attr"`
}
type clientConfigXml struct {
	AgentUserOverrideKey *string                          `xml:"agent-user-override-key,omitempty"`
	Configs              *clientConfigConfigsContainerXml `xml:"configs,omitempty"`
	RootCa               *clientConfigRootCaContainerXml  `xml:"root-ca,omitempty"`
	Misc                 []generic.Xml                    `xml:",any"`
	MiscAttributes       []xml.Attr                       `xml:",any,attr"`
}
type clientConfigConfigsContainerXml struct {
	Entries []clientConfigConfigsXml `xml:"entry"`
}
type clientConfigConfigsXml struct {
	XMLName                          xml.Name                                                `xml:"entry"`
	Name                             string                                                  `xml:"name,attr"`
	SaveUserCredentials              *string                                                 `xml:"save-user-credentials,omitempty"`
	Portal2fa                        *string                                                 `xml:"portal-2fa,omitempty"`
	InternalGateway2fa               *string                                                 `xml:"internal-gateway-2fa,omitempty"`
	AutoDiscoveryExternalGateway2fa  *string                                                 `xml:"auto-discovery-external-gateway-2fa,omitempty"`
	ManualOnlyGateway2fa             *string                                                 `xml:"manual-only-gateway-2fa,omitempty"`
	RefreshConfig                    *string                                                 `xml:"refresh-config,omitempty"`
	MdmAddress                       *string                                                 `xml:"mdm-address,omitempty"`
	MdmEnrollmentPort                *string                                                 `xml:"mdm-enrollment-port,omitempty"`
	SourceUser                       *util.MemberType                                        `xml:"source-user,omitempty"`
	ThirdPartyVpnClients             *util.MemberType                                        `xml:"third-party-vpn-clients,omitempty"`
	Os                               *util.MemberType                                        `xml:"os,omitempty"`
	Certificate                      *clientConfigConfigsCertificateXml                      `xml:"certificate,omitempty"`
	CustomChecks                     *clientConfigConfigsCustomChecksXml                     `xml:"custom-checks,omitempty"`
	Gateways                         *clientConfigConfigsGatewaysXml                         `xml:"gateways,omitempty"`
	InternalHostDetection            *clientConfigConfigsInternalHostDetectionXml            `xml:"internal-host-detection,omitempty"`
	InternalHostDetectionV6          *clientConfigConfigsInternalHostDetectionV6Xml          `xml:"internal-host-detection-v6,omitempty"`
	AgentUi                          *clientConfigConfigsAgentUiXml                          `xml:"agent-ui,omitempty"`
	HipCollection                    *clientConfigConfigsHipCollectionXml                    `xml:"hip-collection,omitempty"`
	AgentConfig                      *clientConfigConfigsAgentConfigXml                      `xml:"agent-config,omitempty"`
	GpAppConfig                      *clientConfigConfigsGpAppConfigXml                      `xml:"gp-app-config,omitempty"`
	AuthenticationOverride           *clientConfigConfigsAuthenticationOverrideXml           `xml:"authentication-override,omitempty"`
	MachineAccountExistsWithSerialno *clientConfigConfigsMachineAccountExistsWithSerialnoXml `xml:"machine-account-exists-with-serialno,omitempty"`
	ClientCertificate                *clientConfigConfigsClientCertificateXml                `xml:"client-certificate,omitempty"`
	Misc                             []generic.Xml                                           `xml:",any"`
	MiscAttributes                   []xml.Attr                                              `xml:",any,attr"`
}
type clientConfigConfigsCertificateXml struct {
	Criteria       *clientConfigConfigsCertificateCriteriaXml `xml:"criteria,omitempty"`
	Misc           []generic.Xml                              `xml:",any"`
	MiscAttributes []xml.Attr                                 `xml:",any,attr"`
}
type clientConfigConfigsCertificateCriteriaXml struct {
	CertificateProfile *string       `xml:"certificate-profile,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksXml struct {
	Criteria       *clientConfigConfigsCustomChecksCriteriaXml `xml:"criteria,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksCriteriaXml struct {
	RegistryKey    *clientConfigConfigsCustomChecksCriteriaRegistryKeyContainerXml `xml:"registry-key,omitempty"`
	Plist          *clientConfigConfigsCustomChecksCriteriaPlistContainerXml       `xml:"plist,omitempty"`
	Misc           []generic.Xml                                                   `xml:",any"`
	MiscAttributes []xml.Attr                                                      `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksCriteriaRegistryKeyContainerXml struct {
	Entries []clientConfigConfigsCustomChecksCriteriaRegistryKeyXml `xml:"entry"`
}
type clientConfigConfigsCustomChecksCriteriaRegistryKeyXml struct {
	XMLName          xml.Name                                                                     `xml:"entry"`
	Name             string                                                                       `xml:"name,attr"`
	DefaultValueData *string                                                                      `xml:"default-value-data,omitempty"`
	Negate           *string                                                                      `xml:"negate,omitempty"`
	RegistryValue    *clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueContainerXml `xml:"registry-value,omitempty"`
	Misc             []generic.Xml                                                                `xml:",any"`
	MiscAttributes   []xml.Attr                                                                   `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueContainerXml struct {
	Entries []clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml `xml:"entry"`
}
type clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	ValueData      *string       `xml:"value-data,omitempty"`
	Negate         *string       `xml:"negate,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksCriteriaPlistContainerXml struct {
	Entries []clientConfigConfigsCustomChecksCriteriaPlistXml `xml:"entry"`
}
type clientConfigConfigsCustomChecksCriteriaPlistXml struct {
	XMLName        xml.Name                                                     `xml:"entry"`
	Name           string                                                       `xml:"name,attr"`
	Negate         *string                                                      `xml:"negate,omitempty"`
	Key            *clientConfigConfigsCustomChecksCriteriaPlistKeyContainerXml `xml:"key,omitempty"`
	Misc           []generic.Xml                                                `xml:",any"`
	MiscAttributes []xml.Attr                                                   `xml:",any,attr"`
}
type clientConfigConfigsCustomChecksCriteriaPlistKeyContainerXml struct {
	Entries []clientConfigConfigsCustomChecksCriteriaPlistKeyXml `xml:"entry"`
}
type clientConfigConfigsCustomChecksCriteriaPlistKeyXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Value          *string       `xml:"value,omitempty"`
	Negate         *string       `xml:"negate,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsGatewaysXml struct {
	Internal       *clientConfigConfigsGatewaysInternalXml `xml:"internal,omitempty"`
	External       *clientConfigConfigsGatewaysExternalXml `xml:"external,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type clientConfigConfigsGatewaysInternalXml struct {
	List           *clientConfigConfigsGatewaysInternalListContainerXml `xml:"list,omitempty"`
	DhcpOptionCode *util.MemberType                                     `xml:"dhcp-option-code,omitempty"`
	Misc           []generic.Xml                                        `xml:",any"`
	MiscAttributes []xml.Attr                                           `xml:",any,attr"`
}
type clientConfigConfigsGatewaysInternalListContainerXml struct {
	Entries []clientConfigConfigsGatewaysInternalListXml `xml:"entry"`
}
type clientConfigConfigsGatewaysInternalListXml struct {
	XMLName        xml.Name                                      `xml:"entry"`
	Name           string                                        `xml:"name,attr"`
	SourceIp       *util.MemberType                              `xml:"source-ip,omitempty"`
	Fqdn           *string                                       `xml:"fqdn,omitempty"`
	Ip             *clientConfigConfigsGatewaysInternalListIpXml `xml:"ip,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type clientConfigConfigsGatewaysInternalListIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsGatewaysExternalXml struct {
	CutoffTime     *int64                                               `xml:"cutoff-time,omitempty"`
	List           *clientConfigConfigsGatewaysExternalListContainerXml `xml:"list,omitempty"`
	Misc           []generic.Xml                                        `xml:",any"`
	MiscAttributes []xml.Attr                                           `xml:",any,attr"`
}
type clientConfigConfigsGatewaysExternalListContainerXml struct {
	Entries []clientConfigConfigsGatewaysExternalListXml `xml:"entry"`
}
type clientConfigConfigsGatewaysExternalListXml struct {
	XMLName        xml.Name                                                         `xml:"entry"`
	Name           string                                                           `xml:"name,attr"`
	PriorityRule   *clientConfigConfigsGatewaysExternalListPriorityRuleContainerXml `xml:"priority-rule,omitempty"`
	Manual         *string                                                          `xml:"manual,omitempty"`
	Fqdn           *string                                                          `xml:"fqdn,omitempty"`
	Ip             *clientConfigConfigsGatewaysExternalListIpXml                    `xml:"ip,omitempty"`
	Misc           []generic.Xml                                                    `xml:",any"`
	MiscAttributes []xml.Attr                                                       `xml:",any,attr"`
}
type clientConfigConfigsGatewaysExternalListPriorityRuleContainerXml struct {
	Entries []clientConfigConfigsGatewaysExternalListPriorityRuleXml `xml:"entry"`
}
type clientConfigConfigsGatewaysExternalListPriorityRuleXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Priority       *string       `xml:"priority,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsGatewaysExternalListIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsInternalHostDetectionXml struct {
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Hostname       *string       `xml:"hostname,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsInternalHostDetectionV6Xml struct {
	IpAddress      *string       `xml:"ip-address,omitempty"`
	Hostname       *string       `xml:"hostname,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsAgentUiXml struct {
	Passcode                 *string                                   `xml:"passcode,omitempty"`
	UninstallPassword        *string                                   `xml:"uninstall-password,omitempty"`
	AgentUserOverrideTimeout *int64                                    `xml:"agent-user-override-timeout,omitempty"`
	MaxAgentUserOverrides    *int64                                    `xml:"max-agent-user-overrides,omitempty"`
	WelcomePage              *clientConfigConfigsAgentUiWelcomePageXml `xml:"welcome-page,omitempty"`
	Misc                     []generic.Xml                             `xml:",any"`
	MiscAttributes           []xml.Attr                                `xml:",any,attr"`
}
type clientConfigConfigsAgentUiWelcomePageXml struct {
	Page           *string       `xml:"page,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionXml struct {
	CertificateProfile *string                                          `xml:"certificate-profile,omitempty"`
	MaxWaitTime        *int64                                           `xml:"max-wait-time,omitempty"`
	CollectHipData     *string                                          `xml:"collect-hip-data,omitempty"`
	Exclusion          *clientConfigConfigsHipCollectionExclusionXml    `xml:"exclusion,omitempty"`
	CustomChecks       *clientConfigConfigsHipCollectionCustomChecksXml `xml:"custom-checks,omitempty"`
	Misc               []generic.Xml                                    `xml:",any"`
	MiscAttributes     []xml.Attr                                       `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionExclusionXml struct {
	Category       *clientConfigConfigsHipCollectionExclusionCategoryContainerXml `xml:"category,omitempty"`
	Misc           []generic.Xml                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                     `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionExclusionCategoryContainerXml struct {
	Entries []clientConfigConfigsHipCollectionExclusionCategoryXml `xml:"entry"`
}
type clientConfigConfigsHipCollectionExclusionCategoryXml struct {
	XMLName        xml.Name                                                             `xml:"entry"`
	Name           string                                                               `xml:"name,attr"`
	Vendor         *clientConfigConfigsHipCollectionExclusionCategoryVendorContainerXml `xml:"vendor,omitempty"`
	Misc           []generic.Xml                                                        `xml:",any"`
	MiscAttributes []xml.Attr                                                           `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionExclusionCategoryVendorContainerXml struct {
	Entries []clientConfigConfigsHipCollectionExclusionCategoryVendorXml `xml:"entry"`
}
type clientConfigConfigsHipCollectionExclusionCategoryVendorXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Product        *util.MemberType `xml:"product,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksXml struct {
	Windows        *clientConfigConfigsHipCollectionCustomChecksWindowsXml `xml:"windows,omitempty"`
	MacOs          *clientConfigConfigsHipCollectionCustomChecksMacOsXml   `xml:"mac-os,omitempty"`
	Linux          *clientConfigConfigsHipCollectionCustomChecksLinuxXml   `xml:"linux,omitempty"`
	Misc           []generic.Xml                                           `xml:",any"`
	MiscAttributes []xml.Attr                                              `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksWindowsXml struct {
	RegistryKey    *clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyContainerXml `xml:"registry-key,omitempty"`
	ProcessList    *util.MemberType                                                            `xml:"process-list,omitempty"`
	Misc           []generic.Xml                                                               `xml:",any"`
	MiscAttributes []xml.Attr                                                                  `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyContainerXml struct {
	Entries []clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml `xml:"entry"`
}
type clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	RegistryValue  *util.MemberType `xml:"registry-value,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksMacOsXml struct {
	Plist          *clientConfigConfigsHipCollectionCustomChecksMacOsPlistContainerXml `xml:"plist,omitempty"`
	ProcessList    *util.MemberType                                                    `xml:"process-list,omitempty"`
	Misc           []generic.Xml                                                       `xml:",any"`
	MiscAttributes []xml.Attr                                                          `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksMacOsPlistContainerXml struct {
	Entries []clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml `xml:"entry"`
}
type clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Key            *util.MemberType `xml:"key,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type clientConfigConfigsHipCollectionCustomChecksLinuxXml struct {
	ProcessList    *util.MemberType `xml:"process-list,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type clientConfigConfigsAgentConfigXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsGpAppConfigXml struct {
	Config         *clientConfigConfigsGpAppConfigConfigContainerXml `xml:"config,omitempty"`
	Misc           []generic.Xml                                     `xml:",any"`
	MiscAttributes []xml.Attr                                        `xml:",any,attr"`
}
type clientConfigConfigsGpAppConfigConfigContainerXml struct {
	Entries []clientConfigConfigsGpAppConfigConfigXml `xml:"entry"`
}
type clientConfigConfigsGpAppConfigConfigXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Value          *util.MemberType `xml:"value,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type clientConfigConfigsAuthenticationOverrideXml struct {
	GenerateCookie           *string                                                   `xml:"generate-cookie,omitempty"`
	CookieEncryptDecryptCert *string                                                   `xml:"cookie-encrypt-decrypt-cert,omitempty"`
	AcceptCookie             *clientConfigConfigsAuthenticationOverrideAcceptCookieXml `xml:"accept-cookie,omitempty"`
	Misc                     []generic.Xml                                             `xml:",any"`
	MiscAttributes           []xml.Attr                                                `xml:",any,attr"`
}
type clientConfigConfigsAuthenticationOverrideAcceptCookieXml struct {
	CookieLifetime *clientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml `xml:"cookie-lifetime,omitempty"`
	Misc           []generic.Xml                                                           `xml:",any"`
	MiscAttributes []xml.Attr                                                              `xml:",any,attr"`
}
type clientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml struct {
	LifetimeInDays    *int64        `xml:"lifetime-in-days,omitempty"`
	LifetimeInHours   *int64        `xml:"lifetime-in-hours,omitempty"`
	LifetimeInMinutes *int64        `xml:"lifetime-in-minutes,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsMachineAccountExistsWithSerialnoXml struct {
	No             *clientConfigConfigsMachineAccountExistsWithSerialnoNoXml  `xml:"no,omitempty"`
	Yes            *clientConfigConfigsMachineAccountExistsWithSerialnoYesXml `xml:"yes,omitempty"`
	Misc           []generic.Xml                                              `xml:",any"`
	MiscAttributes []xml.Attr                                                 `xml:",any,attr"`
}
type clientConfigConfigsMachineAccountExistsWithSerialnoNoXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsMachineAccountExistsWithSerialnoYesXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigConfigsClientCertificateXml struct {
	Local          *string       `xml:"local,omitempty"`
	Scep           *string       `xml:"scep,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientConfigRootCaContainerXml struct {
	Entries []clientConfigRootCaXml `xml:"entry"`
}
type clientConfigRootCaXml struct {
	XMLName            xml.Name      `xml:"entry"`
	Name               string        `xml:"name,attr"`
	InstallInCertStore *string       `xml:"install-in-cert-store,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type clientlessVpnXml struct {
	AppsToUserMapping        *clientlessVpnAppsToUserMappingContainerXml  `xml:"apps-to-user-mapping,omitempty"`
	CryptoSettings           *clientlessVpnCryptoSettingsXml              `xml:"crypto-settings,omitempty"`
	DnsProxy                 *string                                      `xml:"dns-proxy,omitempty"`
	Hostname                 *string                                      `xml:"hostname,omitempty"`
	InactivityLogout         *clientlessVpnInactivityLogoutXml            `xml:"inactivity-logout,omitempty"`
	LoginLifetime            *clientlessVpnLoginLifetimeXml               `xml:"login-lifetime,omitempty"`
	MaxUser                  *int64                                       `xml:"max-user,omitempty"`
	ProxyServerSetting       *clientlessVpnProxyServerSettingContainerXml `xml:"proxy-server-setting,omitempty"`
	RewriteExcludeDomainList *util.MemberType                             `xml:"rewrite-exclude-domain-list,omitempty"`
	SecurityZone             *string                                      `xml:"security-zone,omitempty"`
	Misc                     []generic.Xml                                `xml:",any"`
	MiscAttributes           []xml.Attr                                   `xml:",any,attr"`
}
type clientlessVpnAppsToUserMappingContainerXml struct {
	Entries []clientlessVpnAppsToUserMappingXml `xml:"entry"`
}
type clientlessVpnAppsToUserMappingXml struct {
	XMLName                               xml.Name         `xml:"entry"`
	Name                                  string           `xml:"name,attr"`
	SourceUser                            *util.MemberType `xml:"source-user,omitempty"`
	Applications                          *util.MemberType `xml:"applications,omitempty"`
	EnableCustomAppURLAddressBar          *string          `xml:"enable-custom-app-URL-address-bar,omitempty"`
	DisplayGlobalProtectAgentDownloadLink *string          `xml:"display-global-protect-agent-download-link,omitempty"`
	Misc                                  []generic.Xml    `xml:",any"`
	MiscAttributes                        []xml.Attr       `xml:",any,attr"`
}
type clientlessVpnCryptoSettingsXml struct {
	ServerCertVerification *clientlessVpnCryptoSettingsServerCertVerificationXml `xml:"server-cert-verification,omitempty"`
	SslProtocol            *clientlessVpnCryptoSettingsSslProtocolXml            `xml:"ssl-protocol,omitempty"`
	Misc                   []generic.Xml                                         `xml:",any"`
	MiscAttributes         []xml.Attr                                            `xml:",any,attr"`
}
type clientlessVpnCryptoSettingsServerCertVerificationXml struct {
	BlockExpiredCertificate *string       `xml:"block-expired-certificate,omitempty"`
	BlockTimeoutCert        *string       `xml:"block-timeout-cert,omitempty"`
	BlockUnknownCert        *string       `xml:"block-unknown-cert,omitempty"`
	BlockUntrustedIssuer    *string       `xml:"block-untrusted-issuer,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
	MiscAttributes          []xml.Attr    `xml:",any,attr"`
}
type clientlessVpnCryptoSettingsSslProtocolXml struct {
	AuthAlgoMd5      *string       `xml:"auth-algo-md5,omitempty"`
	AuthAlgoSha1     *string       `xml:"auth-algo-sha1,omitempty"`
	AuthAlgoSha256   *string       `xml:"auth-algo-sha256,omitempty"`
	AuthAlgoSha384   *string       `xml:"auth-algo-sha384,omitempty"`
	EncAlgo3des      *string       `xml:"enc-algo-3des,omitempty"`
	EncAlgoAes128Cbc *string       `xml:"enc-algo-aes-128-cbc,omitempty"`
	EncAlgoAes128Gcm *string       `xml:"enc-algo-aes-128-gcm,omitempty"`
	EncAlgoAes256Cbc *string       `xml:"enc-algo-aes-256-cbc,omitempty"`
	EncAlgoAes256Gcm *string       `xml:"enc-algo-aes-256-gcm,omitempty"`
	EncAlgoRc4       *string       `xml:"enc-algo-rc4,omitempty"`
	KeyxchgAlgoDhe   *string       `xml:"keyxchg-algo-dhe,omitempty"`
	KeyxchgAlgoEcdhe *string       `xml:"keyxchg-algo-ecdhe,omitempty"`
	KeyxchgAlgoRsa   *string       `xml:"keyxchg-algo-rsa,omitempty"`
	MaxVersion       *string       `xml:"max-version,omitempty"`
	MinVersion       *string       `xml:"min-version,omitempty"`
	Misc             []generic.Xml `xml:",any"`
	MiscAttributes   []xml.Attr    `xml:",any,attr"`
}
type clientlessVpnInactivityLogoutXml struct {
	Hours          *int64        `xml:"hours,omitempty"`
	Minutes        *int64        `xml:"minutes,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientlessVpnLoginLifetimeXml struct {
	Hours          *int64        `xml:"hours,omitempty"`
	Minutes        *int64        `xml:"minutes,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type clientlessVpnProxyServerSettingContainerXml struct {
	Entries []clientlessVpnProxyServerSettingXml `xml:"entry"`
}
type clientlessVpnProxyServerSettingXml struct {
	XMLName        xml.Name                                       `xml:"entry"`
	Name           string                                         `xml:"name,attr"`
	Domains        *util.MemberType                               `xml:"domains,omitempty"`
	UseProxy       *string                                        `xml:"use-proxy,omitempty"`
	ProxyServer    *clientlessVpnProxyServerSettingProxyServerXml `xml:"proxy-server,omitempty"`
	Misc           []generic.Xml                                  `xml:",any"`
	MiscAttributes []xml.Attr                                     `xml:",any,attr"`
}
type clientlessVpnProxyServerSettingProxyServerXml struct {
	Server         *string       `xml:"server,omitempty"`
	Port           *int64        `xml:"port,omitempty"`
	User           *string       `xml:"user,omitempty"`
	Password       *string       `xml:"password,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type portalConfigXml struct {
	CertificateProfile   *string                             `xml:"certificate-profile,omitempty"`
	ClientAuth           *portalConfigClientAuthContainerXml `xml:"client-auth,omitempty"`
	ConfigSelection      *portalConfigConfigSelectionXml     `xml:"config-selection,omitempty"`
	CustomHelpPage       *string                             `xml:"custom-help-page,omitempty"`
	CustomHomePage       *string                             `xml:"custom-home-page,omitempty"`
	CustomLoginPage      *string                             `xml:"custom-login-page,omitempty"`
	LocalAddress         *portalConfigLocalAddressXml        `xml:"local-address,omitempty"`
	LogFail              *string                             `xml:"log-fail,omitempty"`
	LogSetting           *string                             `xml:"log-setting,omitempty"`
	LogSuccess           *string                             `xml:"log-success,omitempty"`
	SslTlsServiceProfile *string                             `xml:"ssl-tls-service-profile,omitempty"`
	Misc                 []generic.Xml                       `xml:",any"`
	MiscAttributes       []xml.Attr                          `xml:",any,attr"`
}
type portalConfigClientAuthContainerXml struct {
	Entries []portalConfigClientAuthXml `xml:"entry"`
}
type portalConfigClientAuthXml struct {
	XMLName                            xml.Name      `xml:"entry"`
	Name                               string        `xml:"name,attr"`
	Os                                 *string       `xml:"os,omitempty"`
	AuthenticationProfile              *string       `xml:"authentication-profile,omitempty"`
	AutoRetrievePasscode               *string       `xml:"auto-retrieve-passcode,omitempty"`
	UsernameLabel                      *string       `xml:"username-label,omitempty"`
	PasswordLabel                      *string       `xml:"password-label,omitempty"`
	AuthenticationMessage              *string       `xml:"authentication-message,omitempty"`
	UserCredentialOrClientCertRequired *string       `xml:"user-credential-or-client-cert-required,omitempty"`
	Misc                               []generic.Xml `xml:",any"`
	MiscAttributes                     []xml.Attr    `xml:",any,attr"`
}
type portalConfigConfigSelectionXml struct {
	CertificateProfile *string                                     `xml:"certificate-profile,omitempty"`
	CustomChecks       *portalConfigConfigSelectionCustomChecksXml `xml:"custom-checks,omitempty"`
	Misc               []generic.Xml                               `xml:",any"`
	MiscAttributes     []xml.Attr                                  `xml:",any,attr"`
}
type portalConfigConfigSelectionCustomChecksXml struct {
	MacOs          *portalConfigConfigSelectionCustomChecksMacOsXml   `xml:"mac-os,omitempty"`
	Windows        *portalConfigConfigSelectionCustomChecksWindowsXml `xml:"windows,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type portalConfigConfigSelectionCustomChecksMacOsXml struct {
	Plist          *portalConfigConfigSelectionCustomChecksMacOsPlistContainerXml `xml:"plist,omitempty"`
	Misc           []generic.Xml                                                  `xml:",any"`
	MiscAttributes []xml.Attr                                                     `xml:",any,attr"`
}
type portalConfigConfigSelectionCustomChecksMacOsPlistContainerXml struct {
	Entries []portalConfigConfigSelectionCustomChecksMacOsPlistXml `xml:"entry"`
}
type portalConfigConfigSelectionCustomChecksMacOsPlistXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Key            *util.MemberType `xml:"key,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type portalConfigConfigSelectionCustomChecksWindowsXml struct {
	RegistryKey    *portalConfigConfigSelectionCustomChecksWindowsRegistryKeyContainerXml `xml:"registry-key,omitempty"`
	Misc           []generic.Xml                                                          `xml:",any"`
	MiscAttributes []xml.Attr                                                             `xml:",any,attr"`
}
type portalConfigConfigSelectionCustomChecksWindowsRegistryKeyContainerXml struct {
	Entries []portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml `xml:"entry"`
}
type portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	RegistryValue  *util.MemberType `xml:"registry-value,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type portalConfigLocalAddressXml struct {
	Interface       *string                                `xml:"interface,omitempty"`
	IpAddressFamily *string                                `xml:"ip-address-family,omitempty"`
	FloatingIp      *portalConfigLocalAddressFloatingIpXml `xml:"floating-ip,omitempty"`
	Ip              *portalConfigLocalAddressIpXml         `xml:"ip,omitempty"`
	Misc            []generic.Xml                          `xml:",any"`
	MiscAttributes  []xml.Attr                             `xml:",any,attr"`
}
type portalConfigLocalAddressFloatingIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type portalConfigLocalAddressIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type satelliteConfigXml struct {
	ClientCertificate *satelliteConfigClientCertificateXml `xml:"client-certificate,omitempty"`
	Configs           *satelliteConfigConfigsContainerXml  `xml:"configs,omitempty"`
	RootCa            *util.MemberType                     `xml:"root-ca,omitempty"`
	Misc              []generic.Xml                        `xml:",any"`
	MiscAttributes    []xml.Attr                           `xml:",any,attr"`
}
type satelliteConfigClientCertificateXml struct {
	Local          *satelliteConfigClientCertificateLocalXml `xml:"local,omitempty"`
	Scep           *satelliteConfigClientCertificateScepXml  `xml:"scep,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
	MiscAttributes []xml.Attr                                `xml:",any,attr"`
}
type satelliteConfigClientCertificateLocalXml struct {
	CertificateLifeTime      *int64        `xml:"certificate-life-time,omitempty"`
	CertificateRenewalPeriod *int64        `xml:"certificate-renewal-period,omitempty"`
	IssuingCertificate       *string       `xml:"issuing-certificate,omitempty"`
	OcspResponder            *string       `xml:"ocsp-responder,omitempty"`
	Misc                     []generic.Xml `xml:",any"`
	MiscAttributes           []xml.Attr    `xml:",any,attr"`
}
type satelliteConfigClientCertificateScepXml struct {
	CertificateRenewalPeriod *int64        `xml:"certificate-renewal-period,omitempty"`
	Scep                     *string       `xml:"scep,omitempty"`
	Misc                     []generic.Xml `xml:",any"`
	MiscAttributes           []xml.Attr    `xml:",any,attr"`
}
type satelliteConfigConfigsContainerXml struct {
	Entries []satelliteConfigConfigsXml `xml:"entry"`
}
type satelliteConfigConfigsXml struct {
	XMLName               xml.Name                                    `xml:"entry"`
	Name                  string                                      `xml:"name,attr"`
	Devices               *util.MemberType                            `xml:"devices,omitempty"`
	SourceUser            *util.MemberType                            `xml:"source-user,omitempty"`
	Gateways              *satelliteConfigConfigsGatewaysContainerXml `xml:"gateways,omitempty"`
	ConfigRefreshInterval *int64                                      `xml:"config-refresh-interval,omitempty"`
	Misc                  []generic.Xml                               `xml:",any"`
	MiscAttributes        []xml.Attr                                  `xml:",any,attr"`
}
type satelliteConfigConfigsGatewaysContainerXml struct {
	Entries []satelliteConfigConfigsGatewaysXml `xml:"entry"`
}
type satelliteConfigConfigsGatewaysXml struct {
	XMLName        xml.Name                             `xml:"entry"`
	Name           string                               `xml:"name,attr"`
	Ipv6Preferred  *string                              `xml:"ipv6-preferred,omitempty"`
	Priority       *int64                               `xml:"priority,omitempty"`
	Fqdn           *string                              `xml:"fqdn,omitempty"`
	Ip             *satelliteConfigConfigsGatewaysIpXml `xml:"ip,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type satelliteConfigConfigsGatewaysIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.ClientConfig != nil {
		var obj clientConfigXml
		obj.MarshalFromObject(*s.ClientConfig)
		o.ClientConfig = &obj
	}
	if s.ClientlessVpn != nil {
		var obj clientlessVpnXml
		obj.MarshalFromObject(*s.ClientlessVpn)
		o.ClientlessVpn = &obj
	}
	if s.PortalConfig != nil {
		var obj portalConfigXml
		obj.MarshalFromObject(*s.PortalConfig)
		o.PortalConfig = &obj
	}
	if s.SatelliteConfig != nil {
		var obj satelliteConfigXml
		obj.MarshalFromObject(*s.SatelliteConfig)
		o.SatelliteConfig = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var clientConfigVal *ClientConfig
	if o.ClientConfig != nil {
		obj, err := o.ClientConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		clientConfigVal = obj
	}
	var clientlessVpnVal *ClientlessVpn
	if o.ClientlessVpn != nil {
		obj, err := o.ClientlessVpn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		clientlessVpnVal = obj
	}
	var portalConfigVal *PortalConfig
	if o.PortalConfig != nil {
		obj, err := o.PortalConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		portalConfigVal = obj
	}
	var satelliteConfigVal *SatelliteConfig
	if o.SatelliteConfig != nil {
		obj, err := o.SatelliteConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		satelliteConfigVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		ClientConfig:    clientConfigVal,
		ClientlessVpn:   clientlessVpnVal,
		PortalConfig:    portalConfigVal,
		SatelliteConfig: satelliteConfigVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigXml) MarshalFromObject(s ClientConfig) {
	o.AgentUserOverrideKey = s.AgentUserOverrideKey
	if s.Configs != nil {
		var objs []clientConfigConfigsXml
		for _, elt := range s.Configs {
			var obj clientConfigConfigsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Configs = &clientConfigConfigsContainerXml{Entries: objs}
	}
	if s.RootCa != nil {
		var objs []clientConfigRootCaXml
		for _, elt := range s.RootCa {
			var obj clientConfigRootCaXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RootCa = &clientConfigRootCaContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigXml) UnmarshalToObject() (*ClientConfig, error) {
	var configsVal []ClientConfigConfigs
	if o.Configs != nil {
		for _, elt := range o.Configs.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			configsVal = append(configsVal, *obj)
		}
	}
	var rootCaVal []ClientConfigRootCa
	if o.RootCa != nil {
		for _, elt := range o.RootCa.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rootCaVal = append(rootCaVal, *obj)
		}
	}

	result := &ClientConfig{
		AgentUserOverrideKey: o.AgentUserOverrideKey,
		Configs:              configsVal,
		RootCa:               rootCaVal,
		Misc:                 o.Misc,
		MiscAttributes:       o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsXml) MarshalFromObject(s ClientConfigConfigs) {
	o.Name = s.Name
	o.SaveUserCredentials = s.SaveUserCredentials
	o.Portal2fa = util.YesNo(s.Portal2fa, nil)
	o.InternalGateway2fa = util.YesNo(s.InternalGateway2fa, nil)
	o.AutoDiscoveryExternalGateway2fa = util.YesNo(s.AutoDiscoveryExternalGateway2fa, nil)
	o.ManualOnlyGateway2fa = util.YesNo(s.ManualOnlyGateway2fa, nil)
	o.RefreshConfig = util.YesNo(s.RefreshConfig, nil)
	o.MdmAddress = s.MdmAddress
	o.MdmEnrollmentPort = s.MdmEnrollmentPort
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.ThirdPartyVpnClients != nil {
		o.ThirdPartyVpnClients = util.StrToMem(s.ThirdPartyVpnClients)
	}
	if s.Os != nil {
		o.Os = util.StrToMem(s.Os)
	}
	if s.Certificate != nil {
		var obj clientConfigConfigsCertificateXml
		obj.MarshalFromObject(*s.Certificate)
		o.Certificate = &obj
	}
	if s.CustomChecks != nil {
		var obj clientConfigConfigsCustomChecksXml
		obj.MarshalFromObject(*s.CustomChecks)
		o.CustomChecks = &obj
	}
	if s.Gateways != nil {
		var obj clientConfigConfigsGatewaysXml
		obj.MarshalFromObject(*s.Gateways)
		o.Gateways = &obj
	}
	if s.InternalHostDetection != nil {
		var obj clientConfigConfigsInternalHostDetectionXml
		obj.MarshalFromObject(*s.InternalHostDetection)
		o.InternalHostDetection = &obj
	}
	if s.InternalHostDetectionV6 != nil {
		var obj clientConfigConfigsInternalHostDetectionV6Xml
		obj.MarshalFromObject(*s.InternalHostDetectionV6)
		o.InternalHostDetectionV6 = &obj
	}
	if s.AgentUi != nil {
		var obj clientConfigConfigsAgentUiXml
		obj.MarshalFromObject(*s.AgentUi)
		o.AgentUi = &obj
	}
	if s.HipCollection != nil {
		var obj clientConfigConfigsHipCollectionXml
		obj.MarshalFromObject(*s.HipCollection)
		o.HipCollection = &obj
	}
	if s.AgentConfig != nil {
		var obj clientConfigConfigsAgentConfigXml
		obj.MarshalFromObject(*s.AgentConfig)
		o.AgentConfig = &obj
	}
	if s.GpAppConfig != nil {
		var obj clientConfigConfigsGpAppConfigXml
		obj.MarshalFromObject(*s.GpAppConfig)
		o.GpAppConfig = &obj
	}
	if s.AuthenticationOverride != nil {
		var obj clientConfigConfigsAuthenticationOverrideXml
		obj.MarshalFromObject(*s.AuthenticationOverride)
		o.AuthenticationOverride = &obj
	}
	if s.MachineAccountExistsWithSerialno != nil {
		var obj clientConfigConfigsMachineAccountExistsWithSerialnoXml
		obj.MarshalFromObject(*s.MachineAccountExistsWithSerialno)
		o.MachineAccountExistsWithSerialno = &obj
	}
	if s.ClientCertificate != nil {
		var obj clientConfigConfigsClientCertificateXml
		obj.MarshalFromObject(*s.ClientCertificate)
		o.ClientCertificate = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsXml) UnmarshalToObject() (*ClientConfigConfigs, error) {
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var thirdPartyVpnClientsVal []string
	if o.ThirdPartyVpnClients != nil {
		thirdPartyVpnClientsVal = util.MemToStr(o.ThirdPartyVpnClients)
	}
	var osVal []string
	if o.Os != nil {
		osVal = util.MemToStr(o.Os)
	}
	var certificateVal *ClientConfigConfigsCertificate
	if o.Certificate != nil {
		obj, err := o.Certificate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateVal = obj
	}
	var customChecksVal *ClientConfigConfigsCustomChecks
	if o.CustomChecks != nil {
		obj, err := o.CustomChecks.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customChecksVal = obj
	}
	var gatewaysVal *ClientConfigConfigsGateways
	if o.Gateways != nil {
		obj, err := o.Gateways.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gatewaysVal = obj
	}
	var internalHostDetectionVal *ClientConfigConfigsInternalHostDetection
	if o.InternalHostDetection != nil {
		obj, err := o.InternalHostDetection.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		internalHostDetectionVal = obj
	}
	var internalHostDetectionV6Val *ClientConfigConfigsInternalHostDetectionV6
	if o.InternalHostDetectionV6 != nil {
		obj, err := o.InternalHostDetectionV6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		internalHostDetectionV6Val = obj
	}
	var agentUiVal *ClientConfigConfigsAgentUi
	if o.AgentUi != nil {
		obj, err := o.AgentUi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		agentUiVal = obj
	}
	var hipCollectionVal *ClientConfigConfigsHipCollection
	if o.HipCollection != nil {
		obj, err := o.HipCollection.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hipCollectionVal = obj
	}
	var agentConfigVal *ClientConfigConfigsAgentConfig
	if o.AgentConfig != nil {
		obj, err := o.AgentConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		agentConfigVal = obj
	}
	var gpAppConfigVal *ClientConfigConfigsGpAppConfig
	if o.GpAppConfig != nil {
		obj, err := o.GpAppConfig.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		gpAppConfigVal = obj
	}
	var authenticationOverrideVal *ClientConfigConfigsAuthenticationOverride
	if o.AuthenticationOverride != nil {
		obj, err := o.AuthenticationOverride.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationOverrideVal = obj
	}
	var machineAccountExistsWithSerialnoVal *ClientConfigConfigsMachineAccountExistsWithSerialno
	if o.MachineAccountExistsWithSerialno != nil {
		obj, err := o.MachineAccountExistsWithSerialno.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		machineAccountExistsWithSerialnoVal = obj
	}
	var clientCertificateVal *ClientConfigConfigsClientCertificate
	if o.ClientCertificate != nil {
		obj, err := o.ClientCertificate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		clientCertificateVal = obj
	}

	result := &ClientConfigConfigs{
		Name:                             o.Name,
		SaveUserCredentials:              o.SaveUserCredentials,
		Portal2fa:                        util.AsBool(o.Portal2fa, nil),
		InternalGateway2fa:               util.AsBool(o.InternalGateway2fa, nil),
		AutoDiscoveryExternalGateway2fa:  util.AsBool(o.AutoDiscoveryExternalGateway2fa, nil),
		ManualOnlyGateway2fa:             util.AsBool(o.ManualOnlyGateway2fa, nil),
		RefreshConfig:                    util.AsBool(o.RefreshConfig, nil),
		MdmAddress:                       o.MdmAddress,
		MdmEnrollmentPort:                o.MdmEnrollmentPort,
		SourceUser:                       sourceUserVal,
		ThirdPartyVpnClients:             thirdPartyVpnClientsVal,
		Os:                               osVal,
		Certificate:                      certificateVal,
		CustomChecks:                     customChecksVal,
		Gateways:                         gatewaysVal,
		InternalHostDetection:            internalHostDetectionVal,
		InternalHostDetectionV6:          internalHostDetectionV6Val,
		AgentUi:                          agentUiVal,
		HipCollection:                    hipCollectionVal,
		AgentConfig:                      agentConfigVal,
		GpAppConfig:                      gpAppConfigVal,
		AuthenticationOverride:           authenticationOverrideVal,
		MachineAccountExistsWithSerialno: machineAccountExistsWithSerialnoVal,
		ClientCertificate:                clientCertificateVal,
		Misc:                             o.Misc,
		MiscAttributes:                   o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCertificateXml) MarshalFromObject(s ClientConfigConfigsCertificate) {
	if s.Criteria != nil {
		var obj clientConfigConfigsCertificateCriteriaXml
		obj.MarshalFromObject(*s.Criteria)
		o.Criteria = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCertificateXml) UnmarshalToObject() (*ClientConfigConfigsCertificate, error) {
	var criteriaVal *ClientConfigConfigsCertificateCriteria
	if o.Criteria != nil {
		obj, err := o.Criteria.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		criteriaVal = obj
	}

	result := &ClientConfigConfigsCertificate{
		Criteria:       criteriaVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCertificateCriteriaXml) MarshalFromObject(s ClientConfigConfigsCertificateCriteria) {
	o.CertificateProfile = s.CertificateProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCertificateCriteriaXml) UnmarshalToObject() (*ClientConfigConfigsCertificateCriteria, error) {

	result := &ClientConfigConfigsCertificateCriteria{
		CertificateProfile: o.CertificateProfile,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksXml) MarshalFromObject(s ClientConfigConfigsCustomChecks) {
	if s.Criteria != nil {
		var obj clientConfigConfigsCustomChecksCriteriaXml
		obj.MarshalFromObject(*s.Criteria)
		o.Criteria = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecks, error) {
	var criteriaVal *ClientConfigConfigsCustomChecksCriteria
	if o.Criteria != nil {
		obj, err := o.Criteria.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		criteriaVal = obj
	}

	result := &ClientConfigConfigsCustomChecks{
		Criteria:       criteriaVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksCriteriaXml) MarshalFromObject(s ClientConfigConfigsCustomChecksCriteria) {
	if s.RegistryKey != nil {
		var objs []clientConfigConfigsCustomChecksCriteriaRegistryKeyXml
		for _, elt := range s.RegistryKey {
			var obj clientConfigConfigsCustomChecksCriteriaRegistryKeyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RegistryKey = &clientConfigConfigsCustomChecksCriteriaRegistryKeyContainerXml{Entries: objs}
	}
	if s.Plist != nil {
		var objs []clientConfigConfigsCustomChecksCriteriaPlistXml
		for _, elt := range s.Plist {
			var obj clientConfigConfigsCustomChecksCriteriaPlistXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Plist = &clientConfigConfigsCustomChecksCriteriaPlistContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksCriteriaXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecksCriteria, error) {
	var registryKeyVal []ClientConfigConfigsCustomChecksCriteriaRegistryKey
	if o.RegistryKey != nil {
		for _, elt := range o.RegistryKey.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			registryKeyVal = append(registryKeyVal, *obj)
		}
	}
	var plistVal []ClientConfigConfigsCustomChecksCriteriaPlist
	if o.Plist != nil {
		for _, elt := range o.Plist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			plistVal = append(plistVal, *obj)
		}
	}

	result := &ClientConfigConfigsCustomChecksCriteria{
		RegistryKey:    registryKeyVal,
		Plist:          plistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksCriteriaRegistryKeyXml) MarshalFromObject(s ClientConfigConfigsCustomChecksCriteriaRegistryKey) {
	o.Name = s.Name
	o.DefaultValueData = s.DefaultValueData
	o.Negate = util.YesNo(s.Negate, nil)
	if s.RegistryValue != nil {
		var objs []clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml
		for _, elt := range s.RegistryValue {
			var obj clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RegistryValue = &clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksCriteriaRegistryKeyXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecksCriteriaRegistryKey, error) {
	var registryValueVal []ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue
	if o.RegistryValue != nil {
		for _, elt := range o.RegistryValue.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			registryValueVal = append(registryValueVal, *obj)
		}
	}

	result := &ClientConfigConfigsCustomChecksCriteriaRegistryKey{
		Name:             o.Name,
		DefaultValueData: o.DefaultValueData,
		Negate:           util.AsBool(o.Negate, nil),
		RegistryValue:    registryValueVal,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml) MarshalFromObject(s ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue) {
	o.Name = s.Name
	o.ValueData = s.ValueData
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValueXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue, error) {

	result := &ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue{
		Name:           o.Name,
		ValueData:      o.ValueData,
		Negate:         util.AsBool(o.Negate, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksCriteriaPlistXml) MarshalFromObject(s ClientConfigConfigsCustomChecksCriteriaPlist) {
	o.Name = s.Name
	o.Negate = util.YesNo(s.Negate, nil)
	if s.Key != nil {
		var objs []clientConfigConfigsCustomChecksCriteriaPlistKeyXml
		for _, elt := range s.Key {
			var obj clientConfigConfigsCustomChecksCriteriaPlistKeyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Key = &clientConfigConfigsCustomChecksCriteriaPlistKeyContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksCriteriaPlistXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecksCriteriaPlist, error) {
	var keyVal []ClientConfigConfigsCustomChecksCriteriaPlistKey
	if o.Key != nil {
		for _, elt := range o.Key.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			keyVal = append(keyVal, *obj)
		}
	}

	result := &ClientConfigConfigsCustomChecksCriteriaPlist{
		Name:           o.Name,
		Negate:         util.AsBool(o.Negate, nil),
		Key:            keyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsCustomChecksCriteriaPlistKeyXml) MarshalFromObject(s ClientConfigConfigsCustomChecksCriteriaPlistKey) {
	o.Name = s.Name
	o.Value = s.Value
	o.Negate = util.YesNo(s.Negate, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsCustomChecksCriteriaPlistKeyXml) UnmarshalToObject() (*ClientConfigConfigsCustomChecksCriteriaPlistKey, error) {

	result := &ClientConfigConfigsCustomChecksCriteriaPlistKey{
		Name:           o.Name,
		Value:          o.Value,
		Negate:         util.AsBool(o.Negate, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysXml) MarshalFromObject(s ClientConfigConfigsGateways) {
	if s.Internal != nil {
		var obj clientConfigConfigsGatewaysInternalXml
		obj.MarshalFromObject(*s.Internal)
		o.Internal = &obj
	}
	if s.External != nil {
		var obj clientConfigConfigsGatewaysExternalXml
		obj.MarshalFromObject(*s.External)
		o.External = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysXml) UnmarshalToObject() (*ClientConfigConfigsGateways, error) {
	var internalVal *ClientConfigConfigsGatewaysInternal
	if o.Internal != nil {
		obj, err := o.Internal.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		internalVal = obj
	}
	var externalVal *ClientConfigConfigsGatewaysExternal
	if o.External != nil {
		obj, err := o.External.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		externalVal = obj
	}

	result := &ClientConfigConfigsGateways{
		Internal:       internalVal,
		External:       externalVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysInternalXml) MarshalFromObject(s ClientConfigConfigsGatewaysInternal) {
	if s.List != nil {
		var objs []clientConfigConfigsGatewaysInternalListXml
		for _, elt := range s.List {
			var obj clientConfigConfigsGatewaysInternalListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &clientConfigConfigsGatewaysInternalListContainerXml{Entries: objs}
	}
	if s.DhcpOptionCode != nil {
		o.DhcpOptionCode = util.Int64ToMem(s.DhcpOptionCode)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysInternalXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysInternal, error) {
	var listVal []ClientConfigConfigsGatewaysInternalList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}
	var dhcpOptionCodeVal []int64
	if o.DhcpOptionCode != nil {
		var err error
		dhcpOptionCodeVal, err = util.MemToInt64(o.DhcpOptionCode)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}

	result := &ClientConfigConfigsGatewaysInternal{
		List:           listVal,
		DhcpOptionCode: dhcpOptionCodeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysInternalListXml) MarshalFromObject(s ClientConfigConfigsGatewaysInternalList) {
	o.Name = s.Name
	if s.SourceIp != nil {
		o.SourceIp = util.StrToMem(s.SourceIp)
	}
	o.Fqdn = s.Fqdn
	if s.Ip != nil {
		var obj clientConfigConfigsGatewaysInternalListIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysInternalListXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysInternalList, error) {
	var sourceIpVal []string
	if o.SourceIp != nil {
		sourceIpVal = util.MemToStr(o.SourceIp)
	}
	var ipVal *ClientConfigConfigsGatewaysInternalListIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &ClientConfigConfigsGatewaysInternalList{
		Name:           o.Name,
		SourceIp:       sourceIpVal,
		Fqdn:           o.Fqdn,
		Ip:             ipVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysInternalListIpXml) MarshalFromObject(s ClientConfigConfigsGatewaysInternalListIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysInternalListIpXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysInternalListIp, error) {

	result := &ClientConfigConfigsGatewaysInternalListIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysExternalXml) MarshalFromObject(s ClientConfigConfigsGatewaysExternal) {
	o.CutoffTime = s.CutoffTime
	if s.List != nil {
		var objs []clientConfigConfigsGatewaysExternalListXml
		for _, elt := range s.List {
			var obj clientConfigConfigsGatewaysExternalListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &clientConfigConfigsGatewaysExternalListContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysExternalXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysExternal, error) {
	var listVal []ClientConfigConfigsGatewaysExternalList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}

	result := &ClientConfigConfigsGatewaysExternal{
		CutoffTime:     o.CutoffTime,
		List:           listVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysExternalListXml) MarshalFromObject(s ClientConfigConfigsGatewaysExternalList) {
	o.Name = s.Name
	if s.PriorityRule != nil {
		var objs []clientConfigConfigsGatewaysExternalListPriorityRuleXml
		for _, elt := range s.PriorityRule {
			var obj clientConfigConfigsGatewaysExternalListPriorityRuleXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.PriorityRule = &clientConfigConfigsGatewaysExternalListPriorityRuleContainerXml{Entries: objs}
	}
	o.Manual = util.YesNo(s.Manual, nil)
	o.Fqdn = s.Fqdn
	if s.Ip != nil {
		var obj clientConfigConfigsGatewaysExternalListIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysExternalListXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysExternalList, error) {
	var priorityRuleVal []ClientConfigConfigsGatewaysExternalListPriorityRule
	if o.PriorityRule != nil {
		for _, elt := range o.PriorityRule.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			priorityRuleVal = append(priorityRuleVal, *obj)
		}
	}
	var ipVal *ClientConfigConfigsGatewaysExternalListIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &ClientConfigConfigsGatewaysExternalList{
		Name:           o.Name,
		PriorityRule:   priorityRuleVal,
		Manual:         util.AsBool(o.Manual, nil),
		Fqdn:           o.Fqdn,
		Ip:             ipVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysExternalListPriorityRuleXml) MarshalFromObject(s ClientConfigConfigsGatewaysExternalListPriorityRule) {
	o.Name = s.Name
	o.Priority = s.Priority
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysExternalListPriorityRuleXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysExternalListPriorityRule, error) {

	result := &ClientConfigConfigsGatewaysExternalListPriorityRule{
		Name:           o.Name,
		Priority:       o.Priority,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGatewaysExternalListIpXml) MarshalFromObject(s ClientConfigConfigsGatewaysExternalListIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGatewaysExternalListIpXml) UnmarshalToObject() (*ClientConfigConfigsGatewaysExternalListIp, error) {

	result := &ClientConfigConfigsGatewaysExternalListIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsInternalHostDetectionXml) MarshalFromObject(s ClientConfigConfigsInternalHostDetection) {
	o.IpAddress = s.IpAddress
	o.Hostname = s.Hostname
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsInternalHostDetectionXml) UnmarshalToObject() (*ClientConfigConfigsInternalHostDetection, error) {

	result := &ClientConfigConfigsInternalHostDetection{
		IpAddress:      o.IpAddress,
		Hostname:       o.Hostname,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsInternalHostDetectionV6Xml) MarshalFromObject(s ClientConfigConfigsInternalHostDetectionV6) {
	o.IpAddress = s.IpAddress
	o.Hostname = s.Hostname
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsInternalHostDetectionV6Xml) UnmarshalToObject() (*ClientConfigConfigsInternalHostDetectionV6, error) {

	result := &ClientConfigConfigsInternalHostDetectionV6{
		IpAddress:      o.IpAddress,
		Hostname:       o.Hostname,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAgentUiXml) MarshalFromObject(s ClientConfigConfigsAgentUi) {
	o.Passcode = s.Passcode
	o.UninstallPassword = s.UninstallPassword
	o.AgentUserOverrideTimeout = s.AgentUserOverrideTimeout
	o.MaxAgentUserOverrides = s.MaxAgentUserOverrides
	if s.WelcomePage != nil {
		var obj clientConfigConfigsAgentUiWelcomePageXml
		obj.MarshalFromObject(*s.WelcomePage)
		o.WelcomePage = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAgentUiXml) UnmarshalToObject() (*ClientConfigConfigsAgentUi, error) {
	var welcomePageVal *ClientConfigConfigsAgentUiWelcomePage
	if o.WelcomePage != nil {
		obj, err := o.WelcomePage.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		welcomePageVal = obj
	}

	result := &ClientConfigConfigsAgentUi{
		Passcode:                 o.Passcode,
		UninstallPassword:        o.UninstallPassword,
		AgentUserOverrideTimeout: o.AgentUserOverrideTimeout,
		MaxAgentUserOverrides:    o.MaxAgentUserOverrides,
		WelcomePage:              welcomePageVal,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAgentUiWelcomePageXml) MarshalFromObject(s ClientConfigConfigsAgentUiWelcomePage) {
	o.Page = s.Page
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAgentUiWelcomePageXml) UnmarshalToObject() (*ClientConfigConfigsAgentUiWelcomePage, error) {

	result := &ClientConfigConfigsAgentUiWelcomePage{
		Page:           o.Page,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionXml) MarshalFromObject(s ClientConfigConfigsHipCollection) {
	o.CertificateProfile = s.CertificateProfile
	o.MaxWaitTime = s.MaxWaitTime
	o.CollectHipData = util.YesNo(s.CollectHipData, nil)
	if s.Exclusion != nil {
		var obj clientConfigConfigsHipCollectionExclusionXml
		obj.MarshalFromObject(*s.Exclusion)
		o.Exclusion = &obj
	}
	if s.CustomChecks != nil {
		var obj clientConfigConfigsHipCollectionCustomChecksXml
		obj.MarshalFromObject(*s.CustomChecks)
		o.CustomChecks = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionXml) UnmarshalToObject() (*ClientConfigConfigsHipCollection, error) {
	var exclusionVal *ClientConfigConfigsHipCollectionExclusion
	if o.Exclusion != nil {
		obj, err := o.Exclusion.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		exclusionVal = obj
	}
	var customChecksVal *ClientConfigConfigsHipCollectionCustomChecks
	if o.CustomChecks != nil {
		obj, err := o.CustomChecks.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customChecksVal = obj
	}

	result := &ClientConfigConfigsHipCollection{
		CertificateProfile: o.CertificateProfile,
		MaxWaitTime:        o.MaxWaitTime,
		CollectHipData:     util.AsBool(o.CollectHipData, nil),
		Exclusion:          exclusionVal,
		CustomChecks:       customChecksVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionExclusionXml) MarshalFromObject(s ClientConfigConfigsHipCollectionExclusion) {
	if s.Category != nil {
		var objs []clientConfigConfigsHipCollectionExclusionCategoryXml
		for _, elt := range s.Category {
			var obj clientConfigConfigsHipCollectionExclusionCategoryXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Category = &clientConfigConfigsHipCollectionExclusionCategoryContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionExclusionXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionExclusion, error) {
	var categoryVal []ClientConfigConfigsHipCollectionExclusionCategory
	if o.Category != nil {
		for _, elt := range o.Category.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			categoryVal = append(categoryVal, *obj)
		}
	}

	result := &ClientConfigConfigsHipCollectionExclusion{
		Category:       categoryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionExclusionCategoryXml) MarshalFromObject(s ClientConfigConfigsHipCollectionExclusionCategory) {
	o.Name = s.Name
	if s.Vendor != nil {
		var objs []clientConfigConfigsHipCollectionExclusionCategoryVendorXml
		for _, elt := range s.Vendor {
			var obj clientConfigConfigsHipCollectionExclusionCategoryVendorXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vendor = &clientConfigConfigsHipCollectionExclusionCategoryVendorContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionExclusionCategoryXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionExclusionCategory, error) {
	var vendorVal []ClientConfigConfigsHipCollectionExclusionCategoryVendor
	if o.Vendor != nil {
		for _, elt := range o.Vendor.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vendorVal = append(vendorVal, *obj)
		}
	}

	result := &ClientConfigConfigsHipCollectionExclusionCategory{
		Name:           o.Name,
		Vendor:         vendorVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionExclusionCategoryVendorXml) MarshalFromObject(s ClientConfigConfigsHipCollectionExclusionCategoryVendor) {
	o.Name = s.Name
	if s.Product != nil {
		o.Product = util.StrToMem(s.Product)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionExclusionCategoryVendorXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionExclusionCategoryVendor, error) {
	var productVal []string
	if o.Product != nil {
		productVal = util.MemToStr(o.Product)
	}

	result := &ClientConfigConfigsHipCollectionExclusionCategoryVendor{
		Name:           o.Name,
		Product:        productVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecks) {
	if s.Windows != nil {
		var obj clientConfigConfigsHipCollectionCustomChecksWindowsXml
		obj.MarshalFromObject(*s.Windows)
		o.Windows = &obj
	}
	if s.MacOs != nil {
		var obj clientConfigConfigsHipCollectionCustomChecksMacOsXml
		obj.MarshalFromObject(*s.MacOs)
		o.MacOs = &obj
	}
	if s.Linux != nil {
		var obj clientConfigConfigsHipCollectionCustomChecksLinuxXml
		obj.MarshalFromObject(*s.Linux)
		o.Linux = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecks, error) {
	var windowsVal *ClientConfigConfigsHipCollectionCustomChecksWindows
	if o.Windows != nil {
		obj, err := o.Windows.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		windowsVal = obj
	}
	var macOsVal *ClientConfigConfigsHipCollectionCustomChecksMacOs
	if o.MacOs != nil {
		obj, err := o.MacOs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		macOsVal = obj
	}
	var linuxVal *ClientConfigConfigsHipCollectionCustomChecksLinux
	if o.Linux != nil {
		obj, err := o.Linux.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		linuxVal = obj
	}

	result := &ClientConfigConfigsHipCollectionCustomChecks{
		Windows:        windowsVal,
		MacOs:          macOsVal,
		Linux:          linuxVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksWindowsXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecksWindows) {
	if s.RegistryKey != nil {
		var objs []clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml
		for _, elt := range s.RegistryKey {
			var obj clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RegistryKey = &clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyContainerXml{Entries: objs}
	}
	if s.ProcessList != nil {
		o.ProcessList = util.StrToMem(s.ProcessList)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksWindowsXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecksWindows, error) {
	var registryKeyVal []ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey
	if o.RegistryKey != nil {
		for _, elt := range o.RegistryKey.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			registryKeyVal = append(registryKeyVal, *obj)
		}
	}
	var processListVal []string
	if o.ProcessList != nil {
		processListVal = util.MemToStr(o.ProcessList)
	}

	result := &ClientConfigConfigsHipCollectionCustomChecksWindows{
		RegistryKey:    registryKeyVal,
		ProcessList:    processListVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey) {
	o.Name = s.Name
	if s.RegistryValue != nil {
		o.RegistryValue = util.StrToMem(s.RegistryValue)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksWindowsRegistryKeyXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey, error) {
	var registryValueVal []string
	if o.RegistryValue != nil {
		registryValueVal = util.MemToStr(o.RegistryValue)
	}

	result := &ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey{
		Name:           o.Name,
		RegistryValue:  registryValueVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksMacOsXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecksMacOs) {
	if s.Plist != nil {
		var objs []clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml
		for _, elt := range s.Plist {
			var obj clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Plist = &clientConfigConfigsHipCollectionCustomChecksMacOsPlistContainerXml{Entries: objs}
	}
	if s.ProcessList != nil {
		o.ProcessList = util.StrToMem(s.ProcessList)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksMacOsXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecksMacOs, error) {
	var plistVal []ClientConfigConfigsHipCollectionCustomChecksMacOsPlist
	if o.Plist != nil {
		for _, elt := range o.Plist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			plistVal = append(plistVal, *obj)
		}
	}
	var processListVal []string
	if o.ProcessList != nil {
		processListVal = util.MemToStr(o.ProcessList)
	}

	result := &ClientConfigConfigsHipCollectionCustomChecksMacOs{
		Plist:          plistVal,
		ProcessList:    processListVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecksMacOsPlist) {
	o.Name = s.Name
	if s.Key != nil {
		o.Key = util.StrToMem(s.Key)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksMacOsPlistXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecksMacOsPlist, error) {
	var keyVal []string
	if o.Key != nil {
		keyVal = util.MemToStr(o.Key)
	}

	result := &ClientConfigConfigsHipCollectionCustomChecksMacOsPlist{
		Name:           o.Name,
		Key:            keyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsHipCollectionCustomChecksLinuxXml) MarshalFromObject(s ClientConfigConfigsHipCollectionCustomChecksLinux) {
	if s.ProcessList != nil {
		o.ProcessList = util.StrToMem(s.ProcessList)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsHipCollectionCustomChecksLinuxXml) UnmarshalToObject() (*ClientConfigConfigsHipCollectionCustomChecksLinux, error) {
	var processListVal []string
	if o.ProcessList != nil {
		processListVal = util.MemToStr(o.ProcessList)
	}

	result := &ClientConfigConfigsHipCollectionCustomChecksLinux{
		ProcessList:    processListVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAgentConfigXml) MarshalFromObject(s ClientConfigConfigsAgentConfig) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAgentConfigXml) UnmarshalToObject() (*ClientConfigConfigsAgentConfig, error) {

	result := &ClientConfigConfigsAgentConfig{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGpAppConfigXml) MarshalFromObject(s ClientConfigConfigsGpAppConfig) {
	if s.Config != nil {
		var objs []clientConfigConfigsGpAppConfigConfigXml
		for _, elt := range s.Config {
			var obj clientConfigConfigsGpAppConfigConfigXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Config = &clientConfigConfigsGpAppConfigConfigContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGpAppConfigXml) UnmarshalToObject() (*ClientConfigConfigsGpAppConfig, error) {
	var configVal []ClientConfigConfigsGpAppConfigConfig
	if o.Config != nil {
		for _, elt := range o.Config.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			configVal = append(configVal, *obj)
		}
	}

	result := &ClientConfigConfigsGpAppConfig{
		Config:         configVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsGpAppConfigConfigXml) MarshalFromObject(s ClientConfigConfigsGpAppConfigConfig) {
	o.Name = s.Name
	if s.Value != nil {
		o.Value = util.StrToMem(s.Value)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsGpAppConfigConfigXml) UnmarshalToObject() (*ClientConfigConfigsGpAppConfigConfig, error) {
	var valueVal []string
	if o.Value != nil {
		valueVal = util.MemToStr(o.Value)
	}

	result := &ClientConfigConfigsGpAppConfigConfig{
		Name:           o.Name,
		Value:          valueVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAuthenticationOverrideXml) MarshalFromObject(s ClientConfigConfigsAuthenticationOverride) {
	o.GenerateCookie = util.YesNo(s.GenerateCookie, nil)
	o.CookieEncryptDecryptCert = s.CookieEncryptDecryptCert
	if s.AcceptCookie != nil {
		var obj clientConfigConfigsAuthenticationOverrideAcceptCookieXml
		obj.MarshalFromObject(*s.AcceptCookie)
		o.AcceptCookie = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAuthenticationOverrideXml) UnmarshalToObject() (*ClientConfigConfigsAuthenticationOverride, error) {
	var acceptCookieVal *ClientConfigConfigsAuthenticationOverrideAcceptCookie
	if o.AcceptCookie != nil {
		obj, err := o.AcceptCookie.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		acceptCookieVal = obj
	}

	result := &ClientConfigConfigsAuthenticationOverride{
		GenerateCookie:           util.AsBool(o.GenerateCookie, nil),
		CookieEncryptDecryptCert: o.CookieEncryptDecryptCert,
		AcceptCookie:             acceptCookieVal,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAuthenticationOverrideAcceptCookieXml) MarshalFromObject(s ClientConfigConfigsAuthenticationOverrideAcceptCookie) {
	if s.CookieLifetime != nil {
		var obj clientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml
		obj.MarshalFromObject(*s.CookieLifetime)
		o.CookieLifetime = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAuthenticationOverrideAcceptCookieXml) UnmarshalToObject() (*ClientConfigConfigsAuthenticationOverrideAcceptCookie, error) {
	var cookieLifetimeVal *ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime
	if o.CookieLifetime != nil {
		obj, err := o.CookieLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cookieLifetimeVal = obj
	}

	result := &ClientConfigConfigsAuthenticationOverrideAcceptCookie{
		CookieLifetime: cookieLifetimeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml) MarshalFromObject(s ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime) {
	o.LifetimeInDays = s.LifetimeInDays
	o.LifetimeInHours = s.LifetimeInHours
	o.LifetimeInMinutes = s.LifetimeInMinutes
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml) UnmarshalToObject() (*ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime, error) {

	result := &ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime{
		LifetimeInDays:    o.LifetimeInDays,
		LifetimeInHours:   o.LifetimeInHours,
		LifetimeInMinutes: o.LifetimeInMinutes,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsMachineAccountExistsWithSerialnoXml) MarshalFromObject(s ClientConfigConfigsMachineAccountExistsWithSerialno) {
	if s.No != nil {
		var obj clientConfigConfigsMachineAccountExistsWithSerialnoNoXml
		obj.MarshalFromObject(*s.No)
		o.No = &obj
	}
	if s.Yes != nil {
		var obj clientConfigConfigsMachineAccountExistsWithSerialnoYesXml
		obj.MarshalFromObject(*s.Yes)
		o.Yes = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsMachineAccountExistsWithSerialnoXml) UnmarshalToObject() (*ClientConfigConfigsMachineAccountExistsWithSerialno, error) {
	var noVal *ClientConfigConfigsMachineAccountExistsWithSerialnoNo
	if o.No != nil {
		obj, err := o.No.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noVal = obj
	}
	var yesVal *ClientConfigConfigsMachineAccountExistsWithSerialnoYes
	if o.Yes != nil {
		obj, err := o.Yes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		yesVal = obj
	}

	result := &ClientConfigConfigsMachineAccountExistsWithSerialno{
		No:             noVal,
		Yes:            yesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsMachineAccountExistsWithSerialnoNoXml) MarshalFromObject(s ClientConfigConfigsMachineAccountExistsWithSerialnoNo) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsMachineAccountExistsWithSerialnoNoXml) UnmarshalToObject() (*ClientConfigConfigsMachineAccountExistsWithSerialnoNo, error) {

	result := &ClientConfigConfigsMachineAccountExistsWithSerialnoNo{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsMachineAccountExistsWithSerialnoYesXml) MarshalFromObject(s ClientConfigConfigsMachineAccountExistsWithSerialnoYes) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsMachineAccountExistsWithSerialnoYesXml) UnmarshalToObject() (*ClientConfigConfigsMachineAccountExistsWithSerialnoYes, error) {

	result := &ClientConfigConfigsMachineAccountExistsWithSerialnoYes{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigConfigsClientCertificateXml) MarshalFromObject(s ClientConfigConfigsClientCertificate) {
	o.Local = s.Local
	o.Scep = s.Scep
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigConfigsClientCertificateXml) UnmarshalToObject() (*ClientConfigConfigsClientCertificate, error) {

	result := &ClientConfigConfigsClientCertificate{
		Local:          o.Local,
		Scep:           o.Scep,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientConfigRootCaXml) MarshalFromObject(s ClientConfigRootCa) {
	o.Name = s.Name
	o.InstallInCertStore = util.YesNo(s.InstallInCertStore, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientConfigRootCaXml) UnmarshalToObject() (*ClientConfigRootCa, error) {

	result := &ClientConfigRootCa{
		Name:               o.Name,
		InstallInCertStore: util.AsBool(o.InstallInCertStore, nil),
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnXml) MarshalFromObject(s ClientlessVpn) {
	if s.AppsToUserMapping != nil {
		var objs []clientlessVpnAppsToUserMappingXml
		for _, elt := range s.AppsToUserMapping {
			var obj clientlessVpnAppsToUserMappingXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.AppsToUserMapping = &clientlessVpnAppsToUserMappingContainerXml{Entries: objs}
	}
	if s.CryptoSettings != nil {
		var obj clientlessVpnCryptoSettingsXml
		obj.MarshalFromObject(*s.CryptoSettings)
		o.CryptoSettings = &obj
	}
	o.DnsProxy = s.DnsProxy
	o.Hostname = s.Hostname
	if s.InactivityLogout != nil {
		var obj clientlessVpnInactivityLogoutXml
		obj.MarshalFromObject(*s.InactivityLogout)
		o.InactivityLogout = &obj
	}
	if s.LoginLifetime != nil {
		var obj clientlessVpnLoginLifetimeXml
		obj.MarshalFromObject(*s.LoginLifetime)
		o.LoginLifetime = &obj
	}
	o.MaxUser = s.MaxUser
	if s.ProxyServerSetting != nil {
		var objs []clientlessVpnProxyServerSettingXml
		for _, elt := range s.ProxyServerSetting {
			var obj clientlessVpnProxyServerSettingXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ProxyServerSetting = &clientlessVpnProxyServerSettingContainerXml{Entries: objs}
	}
	if s.RewriteExcludeDomainList != nil {
		o.RewriteExcludeDomainList = util.StrToMem(s.RewriteExcludeDomainList)
	}
	o.SecurityZone = s.SecurityZone
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnXml) UnmarshalToObject() (*ClientlessVpn, error) {
	var appsToUserMappingVal []ClientlessVpnAppsToUserMapping
	if o.AppsToUserMapping != nil {
		for _, elt := range o.AppsToUserMapping.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			appsToUserMappingVal = append(appsToUserMappingVal, *obj)
		}
	}
	var cryptoSettingsVal *ClientlessVpnCryptoSettings
	if o.CryptoSettings != nil {
		obj, err := o.CryptoSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cryptoSettingsVal = obj
	}
	var inactivityLogoutVal *ClientlessVpnInactivityLogout
	if o.InactivityLogout != nil {
		obj, err := o.InactivityLogout.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inactivityLogoutVal = obj
	}
	var loginLifetimeVal *ClientlessVpnLoginLifetime
	if o.LoginLifetime != nil {
		obj, err := o.LoginLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		loginLifetimeVal = obj
	}
	var proxyServerSettingVal []ClientlessVpnProxyServerSetting
	if o.ProxyServerSetting != nil {
		for _, elt := range o.ProxyServerSetting.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			proxyServerSettingVal = append(proxyServerSettingVal, *obj)
		}
	}
	var rewriteExcludeDomainListVal []string
	if o.RewriteExcludeDomainList != nil {
		rewriteExcludeDomainListVal = util.MemToStr(o.RewriteExcludeDomainList)
	}

	result := &ClientlessVpn{
		AppsToUserMapping:        appsToUserMappingVal,
		CryptoSettings:           cryptoSettingsVal,
		DnsProxy:                 o.DnsProxy,
		Hostname:                 o.Hostname,
		InactivityLogout:         inactivityLogoutVal,
		LoginLifetime:            loginLifetimeVal,
		MaxUser:                  o.MaxUser,
		ProxyServerSetting:       proxyServerSettingVal,
		RewriteExcludeDomainList: rewriteExcludeDomainListVal,
		SecurityZone:             o.SecurityZone,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnAppsToUserMappingXml) MarshalFromObject(s ClientlessVpnAppsToUserMapping) {
	o.Name = s.Name
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Applications != nil {
		o.Applications = util.StrToMem(s.Applications)
	}
	o.EnableCustomAppURLAddressBar = util.YesNo(s.EnableCustomAppURLAddressBar, nil)
	o.DisplayGlobalProtectAgentDownloadLink = util.YesNo(s.DisplayGlobalProtectAgentDownloadLink, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnAppsToUserMappingXml) UnmarshalToObject() (*ClientlessVpnAppsToUserMapping, error) {
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var applicationsVal []string
	if o.Applications != nil {
		applicationsVal = util.MemToStr(o.Applications)
	}

	result := &ClientlessVpnAppsToUserMapping{
		Name:                                  o.Name,
		SourceUser:                            sourceUserVal,
		Applications:                          applicationsVal,
		EnableCustomAppURLAddressBar:          util.AsBool(o.EnableCustomAppURLAddressBar, nil),
		DisplayGlobalProtectAgentDownloadLink: util.AsBool(o.DisplayGlobalProtectAgentDownloadLink, nil),
		Misc:                                  o.Misc,
		MiscAttributes:                        o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnCryptoSettingsXml) MarshalFromObject(s ClientlessVpnCryptoSettings) {
	if s.ServerCertVerification != nil {
		var obj clientlessVpnCryptoSettingsServerCertVerificationXml
		obj.MarshalFromObject(*s.ServerCertVerification)
		o.ServerCertVerification = &obj
	}
	if s.SslProtocol != nil {
		var obj clientlessVpnCryptoSettingsSslProtocolXml
		obj.MarshalFromObject(*s.SslProtocol)
		o.SslProtocol = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnCryptoSettingsXml) UnmarshalToObject() (*ClientlessVpnCryptoSettings, error) {
	var serverCertVerificationVal *ClientlessVpnCryptoSettingsServerCertVerification
	if o.ServerCertVerification != nil {
		obj, err := o.ServerCertVerification.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverCertVerificationVal = obj
	}
	var sslProtocolVal *ClientlessVpnCryptoSettingsSslProtocol
	if o.SslProtocol != nil {
		obj, err := o.SslProtocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sslProtocolVal = obj
	}

	result := &ClientlessVpnCryptoSettings{
		ServerCertVerification: serverCertVerificationVal,
		SslProtocol:            sslProtocolVal,
		Misc:                   o.Misc,
		MiscAttributes:         o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnCryptoSettingsServerCertVerificationXml) MarshalFromObject(s ClientlessVpnCryptoSettingsServerCertVerification) {
	o.BlockExpiredCertificate = util.YesNo(s.BlockExpiredCertificate, nil)
	o.BlockTimeoutCert = util.YesNo(s.BlockTimeoutCert, nil)
	o.BlockUnknownCert = util.YesNo(s.BlockUnknownCert, nil)
	o.BlockUntrustedIssuer = util.YesNo(s.BlockUntrustedIssuer, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnCryptoSettingsServerCertVerificationXml) UnmarshalToObject() (*ClientlessVpnCryptoSettingsServerCertVerification, error) {

	result := &ClientlessVpnCryptoSettingsServerCertVerification{
		BlockExpiredCertificate: util.AsBool(o.BlockExpiredCertificate, nil),
		BlockTimeoutCert:        util.AsBool(o.BlockTimeoutCert, nil),
		BlockUnknownCert:        util.AsBool(o.BlockUnknownCert, nil),
		BlockUntrustedIssuer:    util.AsBool(o.BlockUntrustedIssuer, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnCryptoSettingsSslProtocolXml) MarshalFromObject(s ClientlessVpnCryptoSettingsSslProtocol) {
	o.AuthAlgoMd5 = util.YesNo(s.AuthAlgoMd5, nil)
	o.AuthAlgoSha1 = util.YesNo(s.AuthAlgoSha1, nil)
	o.AuthAlgoSha256 = util.YesNo(s.AuthAlgoSha256, nil)
	o.AuthAlgoSha384 = util.YesNo(s.AuthAlgoSha384, nil)
	o.EncAlgo3des = util.YesNo(s.EncAlgo3des, nil)
	o.EncAlgoAes128Cbc = util.YesNo(s.EncAlgoAes128Cbc, nil)
	o.EncAlgoAes128Gcm = util.YesNo(s.EncAlgoAes128Gcm, nil)
	o.EncAlgoAes256Cbc = util.YesNo(s.EncAlgoAes256Cbc, nil)
	o.EncAlgoAes256Gcm = util.YesNo(s.EncAlgoAes256Gcm, nil)
	o.EncAlgoRc4 = util.YesNo(s.EncAlgoRc4, nil)
	o.KeyxchgAlgoDhe = util.YesNo(s.KeyxchgAlgoDhe, nil)
	o.KeyxchgAlgoEcdhe = util.YesNo(s.KeyxchgAlgoEcdhe, nil)
	o.KeyxchgAlgoRsa = util.YesNo(s.KeyxchgAlgoRsa, nil)
	o.MaxVersion = s.MaxVersion
	o.MinVersion = s.MinVersion
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnCryptoSettingsSslProtocolXml) UnmarshalToObject() (*ClientlessVpnCryptoSettingsSslProtocol, error) {

	result := &ClientlessVpnCryptoSettingsSslProtocol{
		AuthAlgoMd5:      util.AsBool(o.AuthAlgoMd5, nil),
		AuthAlgoSha1:     util.AsBool(o.AuthAlgoSha1, nil),
		AuthAlgoSha256:   util.AsBool(o.AuthAlgoSha256, nil),
		AuthAlgoSha384:   util.AsBool(o.AuthAlgoSha384, nil),
		EncAlgo3des:      util.AsBool(o.EncAlgo3des, nil),
		EncAlgoAes128Cbc: util.AsBool(o.EncAlgoAes128Cbc, nil),
		EncAlgoAes128Gcm: util.AsBool(o.EncAlgoAes128Gcm, nil),
		EncAlgoAes256Cbc: util.AsBool(o.EncAlgoAes256Cbc, nil),
		EncAlgoAes256Gcm: util.AsBool(o.EncAlgoAes256Gcm, nil),
		EncAlgoRc4:       util.AsBool(o.EncAlgoRc4, nil),
		KeyxchgAlgoDhe:   util.AsBool(o.KeyxchgAlgoDhe, nil),
		KeyxchgAlgoEcdhe: util.AsBool(o.KeyxchgAlgoEcdhe, nil),
		KeyxchgAlgoRsa:   util.AsBool(o.KeyxchgAlgoRsa, nil),
		MaxVersion:       o.MaxVersion,
		MinVersion:       o.MinVersion,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnInactivityLogoutXml) MarshalFromObject(s ClientlessVpnInactivityLogout) {
	o.Hours = s.Hours
	o.Minutes = s.Minutes
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnInactivityLogoutXml) UnmarshalToObject() (*ClientlessVpnInactivityLogout, error) {

	result := &ClientlessVpnInactivityLogout{
		Hours:          o.Hours,
		Minutes:        o.Minutes,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnLoginLifetimeXml) MarshalFromObject(s ClientlessVpnLoginLifetime) {
	o.Hours = s.Hours
	o.Minutes = s.Minutes
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnLoginLifetimeXml) UnmarshalToObject() (*ClientlessVpnLoginLifetime, error) {

	result := &ClientlessVpnLoginLifetime{
		Hours:          o.Hours,
		Minutes:        o.Minutes,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnProxyServerSettingXml) MarshalFromObject(s ClientlessVpnProxyServerSetting) {
	o.Name = s.Name
	if s.Domains != nil {
		o.Domains = util.StrToMem(s.Domains)
	}
	o.UseProxy = util.YesNo(s.UseProxy, nil)
	if s.ProxyServer != nil {
		var obj clientlessVpnProxyServerSettingProxyServerXml
		obj.MarshalFromObject(*s.ProxyServer)
		o.ProxyServer = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnProxyServerSettingXml) UnmarshalToObject() (*ClientlessVpnProxyServerSetting, error) {
	var domainsVal []string
	if o.Domains != nil {
		domainsVal = util.MemToStr(o.Domains)
	}
	var proxyServerVal *ClientlessVpnProxyServerSettingProxyServer
	if o.ProxyServer != nil {
		obj, err := o.ProxyServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		proxyServerVal = obj
	}

	result := &ClientlessVpnProxyServerSetting{
		Name:           o.Name,
		Domains:        domainsVal,
		UseProxy:       util.AsBool(o.UseProxy, nil),
		ProxyServer:    proxyServerVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *clientlessVpnProxyServerSettingProxyServerXml) MarshalFromObject(s ClientlessVpnProxyServerSettingProxyServer) {
	o.Server = s.Server
	o.Port = s.Port
	o.User = s.User
	o.Password = s.Password
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o clientlessVpnProxyServerSettingProxyServerXml) UnmarshalToObject() (*ClientlessVpnProxyServerSettingProxyServer, error) {

	result := &ClientlessVpnProxyServerSettingProxyServer{
		Server:         o.Server,
		Port:           o.Port,
		User:           o.User,
		Password:       o.Password,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigXml) MarshalFromObject(s PortalConfig) {
	o.CertificateProfile = s.CertificateProfile
	if s.ClientAuth != nil {
		var objs []portalConfigClientAuthXml
		for _, elt := range s.ClientAuth {
			var obj portalConfigClientAuthXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ClientAuth = &portalConfigClientAuthContainerXml{Entries: objs}
	}
	if s.ConfigSelection != nil {
		var obj portalConfigConfigSelectionXml
		obj.MarshalFromObject(*s.ConfigSelection)
		o.ConfigSelection = &obj
	}
	o.CustomHelpPage = s.CustomHelpPage
	o.CustomHomePage = s.CustomHomePage
	o.CustomLoginPage = s.CustomLoginPage
	if s.LocalAddress != nil {
		var obj portalConfigLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.LogFail = util.YesNo(s.LogFail, nil)
	o.LogSetting = s.LogSetting
	o.LogSuccess = util.YesNo(s.LogSuccess, nil)
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigXml) UnmarshalToObject() (*PortalConfig, error) {
	var clientAuthVal []PortalConfigClientAuth
	if o.ClientAuth != nil {
		for _, elt := range o.ClientAuth.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			clientAuthVal = append(clientAuthVal, *obj)
		}
	}
	var configSelectionVal *PortalConfigConfigSelection
	if o.ConfigSelection != nil {
		obj, err := o.ConfigSelection.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		configSelectionVal = obj
	}
	var localAddressVal *PortalConfigLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}

	result := &PortalConfig{
		CertificateProfile:   o.CertificateProfile,
		ClientAuth:           clientAuthVal,
		ConfigSelection:      configSelectionVal,
		CustomHelpPage:       o.CustomHelpPage,
		CustomHomePage:       o.CustomHomePage,
		CustomLoginPage:      o.CustomLoginPage,
		LocalAddress:         localAddressVal,
		LogFail:              util.AsBool(o.LogFail, nil),
		LogSetting:           o.LogSetting,
		LogSuccess:           util.AsBool(o.LogSuccess, nil),
		SslTlsServiceProfile: o.SslTlsServiceProfile,
		Misc:                 o.Misc,
		MiscAttributes:       o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigClientAuthXml) MarshalFromObject(s PortalConfigClientAuth) {
	o.Name = s.Name
	o.Os = s.Os
	o.AuthenticationProfile = s.AuthenticationProfile
	o.AutoRetrievePasscode = util.YesNo(s.AutoRetrievePasscode, nil)
	o.UsernameLabel = s.UsernameLabel
	o.PasswordLabel = s.PasswordLabel
	o.AuthenticationMessage = s.AuthenticationMessage
	o.UserCredentialOrClientCertRequired = s.UserCredentialOrClientCertRequired
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigClientAuthXml) UnmarshalToObject() (*PortalConfigClientAuth, error) {

	result := &PortalConfigClientAuth{
		Name:                               o.Name,
		Os:                                 o.Os,
		AuthenticationProfile:              o.AuthenticationProfile,
		AutoRetrievePasscode:               util.AsBool(o.AutoRetrievePasscode, nil),
		UsernameLabel:                      o.UsernameLabel,
		PasswordLabel:                      o.PasswordLabel,
		AuthenticationMessage:              o.AuthenticationMessage,
		UserCredentialOrClientCertRequired: o.UserCredentialOrClientCertRequired,
		Misc:                               o.Misc,
		MiscAttributes:                     o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionXml) MarshalFromObject(s PortalConfigConfigSelection) {
	o.CertificateProfile = s.CertificateProfile
	if s.CustomChecks != nil {
		var obj portalConfigConfigSelectionCustomChecksXml
		obj.MarshalFromObject(*s.CustomChecks)
		o.CustomChecks = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionXml) UnmarshalToObject() (*PortalConfigConfigSelection, error) {
	var customChecksVal *PortalConfigConfigSelectionCustomChecks
	if o.CustomChecks != nil {
		obj, err := o.CustomChecks.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customChecksVal = obj
	}

	result := &PortalConfigConfigSelection{
		CertificateProfile: o.CertificateProfile,
		CustomChecks:       customChecksVal,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionCustomChecksXml) MarshalFromObject(s PortalConfigConfigSelectionCustomChecks) {
	if s.MacOs != nil {
		var obj portalConfigConfigSelectionCustomChecksMacOsXml
		obj.MarshalFromObject(*s.MacOs)
		o.MacOs = &obj
	}
	if s.Windows != nil {
		var obj portalConfigConfigSelectionCustomChecksWindowsXml
		obj.MarshalFromObject(*s.Windows)
		o.Windows = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionCustomChecksXml) UnmarshalToObject() (*PortalConfigConfigSelectionCustomChecks, error) {
	var macOsVal *PortalConfigConfigSelectionCustomChecksMacOs
	if o.MacOs != nil {
		obj, err := o.MacOs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		macOsVal = obj
	}
	var windowsVal *PortalConfigConfigSelectionCustomChecksWindows
	if o.Windows != nil {
		obj, err := o.Windows.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		windowsVal = obj
	}

	result := &PortalConfigConfigSelectionCustomChecks{
		MacOs:          macOsVal,
		Windows:        windowsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionCustomChecksMacOsXml) MarshalFromObject(s PortalConfigConfigSelectionCustomChecksMacOs) {
	if s.Plist != nil {
		var objs []portalConfigConfigSelectionCustomChecksMacOsPlistXml
		for _, elt := range s.Plist {
			var obj portalConfigConfigSelectionCustomChecksMacOsPlistXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Plist = &portalConfigConfigSelectionCustomChecksMacOsPlistContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionCustomChecksMacOsXml) UnmarshalToObject() (*PortalConfigConfigSelectionCustomChecksMacOs, error) {
	var plistVal []PortalConfigConfigSelectionCustomChecksMacOsPlist
	if o.Plist != nil {
		for _, elt := range o.Plist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			plistVal = append(plistVal, *obj)
		}
	}

	result := &PortalConfigConfigSelectionCustomChecksMacOs{
		Plist:          plistVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionCustomChecksMacOsPlistXml) MarshalFromObject(s PortalConfigConfigSelectionCustomChecksMacOsPlist) {
	o.Name = s.Name
	if s.Key != nil {
		o.Key = util.StrToMem(s.Key)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionCustomChecksMacOsPlistXml) UnmarshalToObject() (*PortalConfigConfigSelectionCustomChecksMacOsPlist, error) {
	var keyVal []string
	if o.Key != nil {
		keyVal = util.MemToStr(o.Key)
	}

	result := &PortalConfigConfigSelectionCustomChecksMacOsPlist{
		Name:           o.Name,
		Key:            keyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionCustomChecksWindowsXml) MarshalFromObject(s PortalConfigConfigSelectionCustomChecksWindows) {
	if s.RegistryKey != nil {
		var objs []portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml
		for _, elt := range s.RegistryKey {
			var obj portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RegistryKey = &portalConfigConfigSelectionCustomChecksWindowsRegistryKeyContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionCustomChecksWindowsXml) UnmarshalToObject() (*PortalConfigConfigSelectionCustomChecksWindows, error) {
	var registryKeyVal []PortalConfigConfigSelectionCustomChecksWindowsRegistryKey
	if o.RegistryKey != nil {
		for _, elt := range o.RegistryKey.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			registryKeyVal = append(registryKeyVal, *obj)
		}
	}

	result := &PortalConfigConfigSelectionCustomChecksWindows{
		RegistryKey:    registryKeyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml) MarshalFromObject(s PortalConfigConfigSelectionCustomChecksWindowsRegistryKey) {
	o.Name = s.Name
	if s.RegistryValue != nil {
		o.RegistryValue = util.StrToMem(s.RegistryValue)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigConfigSelectionCustomChecksWindowsRegistryKeyXml) UnmarshalToObject() (*PortalConfigConfigSelectionCustomChecksWindowsRegistryKey, error) {
	var registryValueVal []string
	if o.RegistryValue != nil {
		registryValueVal = util.MemToStr(o.RegistryValue)
	}

	result := &PortalConfigConfigSelectionCustomChecksWindowsRegistryKey{
		Name:           o.Name,
		RegistryValue:  registryValueVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigLocalAddressXml) MarshalFromObject(s PortalConfigLocalAddress) {
	o.Interface = s.Interface
	o.IpAddressFamily = s.IpAddressFamily
	if s.FloatingIp != nil {
		var obj portalConfigLocalAddressFloatingIpXml
		obj.MarshalFromObject(*s.FloatingIp)
		o.FloatingIp = &obj
	}
	if s.Ip != nil {
		var obj portalConfigLocalAddressIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigLocalAddressXml) UnmarshalToObject() (*PortalConfigLocalAddress, error) {
	var floatingIpVal *PortalConfigLocalAddressFloatingIp
	if o.FloatingIp != nil {
		obj, err := o.FloatingIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floatingIpVal = obj
	}
	var ipVal *PortalConfigLocalAddressIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &PortalConfigLocalAddress{
		Interface:       o.Interface,
		IpAddressFamily: o.IpAddressFamily,
		FloatingIp:      floatingIpVal,
		Ip:              ipVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigLocalAddressFloatingIpXml) MarshalFromObject(s PortalConfigLocalAddressFloatingIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigLocalAddressFloatingIpXml) UnmarshalToObject() (*PortalConfigLocalAddressFloatingIp, error) {

	result := &PortalConfigLocalAddressFloatingIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *portalConfigLocalAddressIpXml) MarshalFromObject(s PortalConfigLocalAddressIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o portalConfigLocalAddressIpXml) UnmarshalToObject() (*PortalConfigLocalAddressIp, error) {

	result := &PortalConfigLocalAddressIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigXml) MarshalFromObject(s SatelliteConfig) {
	if s.ClientCertificate != nil {
		var obj satelliteConfigClientCertificateXml
		obj.MarshalFromObject(*s.ClientCertificate)
		o.ClientCertificate = &obj
	}
	if s.Configs != nil {
		var objs []satelliteConfigConfigsXml
		for _, elt := range s.Configs {
			var obj satelliteConfigConfigsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Configs = &satelliteConfigConfigsContainerXml{Entries: objs}
	}
	if s.RootCa != nil {
		o.RootCa = util.StrToMem(s.RootCa)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigXml) UnmarshalToObject() (*SatelliteConfig, error) {
	var clientCertificateVal *SatelliteConfigClientCertificate
	if o.ClientCertificate != nil {
		obj, err := o.ClientCertificate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		clientCertificateVal = obj
	}
	var configsVal []SatelliteConfigConfigs
	if o.Configs != nil {
		for _, elt := range o.Configs.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			configsVal = append(configsVal, *obj)
		}
	}
	var rootCaVal []string
	if o.RootCa != nil {
		rootCaVal = util.MemToStr(o.RootCa)
	}

	result := &SatelliteConfig{
		ClientCertificate: clientCertificateVal,
		Configs:           configsVal,
		RootCa:            rootCaVal,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigClientCertificateXml) MarshalFromObject(s SatelliteConfigClientCertificate) {
	if s.Local != nil {
		var obj satelliteConfigClientCertificateLocalXml
		obj.MarshalFromObject(*s.Local)
		o.Local = &obj
	}
	if s.Scep != nil {
		var obj satelliteConfigClientCertificateScepXml
		obj.MarshalFromObject(*s.Scep)
		o.Scep = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigClientCertificateXml) UnmarshalToObject() (*SatelliteConfigClientCertificate, error) {
	var localVal *SatelliteConfigClientCertificateLocal
	if o.Local != nil {
		obj, err := o.Local.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localVal = obj
	}
	var scepVal *SatelliteConfigClientCertificateScep
	if o.Scep != nil {
		obj, err := o.Scep.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		scepVal = obj
	}

	result := &SatelliteConfigClientCertificate{
		Local:          localVal,
		Scep:           scepVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigClientCertificateLocalXml) MarshalFromObject(s SatelliteConfigClientCertificateLocal) {
	o.CertificateLifeTime = s.CertificateLifeTime
	o.CertificateRenewalPeriod = s.CertificateRenewalPeriod
	o.IssuingCertificate = s.IssuingCertificate
	o.OcspResponder = s.OcspResponder
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigClientCertificateLocalXml) UnmarshalToObject() (*SatelliteConfigClientCertificateLocal, error) {

	result := &SatelliteConfigClientCertificateLocal{
		CertificateLifeTime:      o.CertificateLifeTime,
		CertificateRenewalPeriod: o.CertificateRenewalPeriod,
		IssuingCertificate:       o.IssuingCertificate,
		OcspResponder:            o.OcspResponder,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigClientCertificateScepXml) MarshalFromObject(s SatelliteConfigClientCertificateScep) {
	o.CertificateRenewalPeriod = s.CertificateRenewalPeriod
	o.Scep = s.Scep
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigClientCertificateScepXml) UnmarshalToObject() (*SatelliteConfigClientCertificateScep, error) {

	result := &SatelliteConfigClientCertificateScep{
		CertificateRenewalPeriod: o.CertificateRenewalPeriod,
		Scep:                     o.Scep,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigConfigsXml) MarshalFromObject(s SatelliteConfigConfigs) {
	o.Name = s.Name
	if s.Devices != nil {
		o.Devices = util.StrToMem(s.Devices)
	}
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Gateways != nil {
		var objs []satelliteConfigConfigsGatewaysXml
		for _, elt := range s.Gateways {
			var obj satelliteConfigConfigsGatewaysXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Gateways = &satelliteConfigConfigsGatewaysContainerXml{Entries: objs}
	}
	o.ConfigRefreshInterval = s.ConfigRefreshInterval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigConfigsXml) UnmarshalToObject() (*SatelliteConfigConfigs, error) {
	var devicesVal []string
	if o.Devices != nil {
		devicesVal = util.MemToStr(o.Devices)
	}
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var gatewaysVal []SatelliteConfigConfigsGateways
	if o.Gateways != nil {
		for _, elt := range o.Gateways.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			gatewaysVal = append(gatewaysVal, *obj)
		}
	}

	result := &SatelliteConfigConfigs{
		Name:                  o.Name,
		Devices:               devicesVal,
		SourceUser:            sourceUserVal,
		Gateways:              gatewaysVal,
		ConfigRefreshInterval: o.ConfigRefreshInterval,
		Misc:                  o.Misc,
		MiscAttributes:        o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigConfigsGatewaysXml) MarshalFromObject(s SatelliteConfigConfigsGateways) {
	o.Name = s.Name
	o.Ipv6Preferred = util.YesNo(s.Ipv6Preferred, nil)
	o.Priority = s.Priority
	o.Fqdn = s.Fqdn
	if s.Ip != nil {
		var obj satelliteConfigConfigsGatewaysIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigConfigsGatewaysXml) UnmarshalToObject() (*SatelliteConfigConfigsGateways, error) {
	var ipVal *SatelliteConfigConfigsGatewaysIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &SatelliteConfigConfigsGateways{
		Name:           o.Name,
		Ipv6Preferred:  util.AsBool(o.Ipv6Preferred, nil),
		Priority:       o.Priority,
		Fqdn:           o.Fqdn,
		Ip:             ipVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *satelliteConfigConfigsGatewaysIpXml) MarshalFromObject(s SatelliteConfigConfigsGatewaysIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o satelliteConfigConfigsGatewaysIpXml) UnmarshalToObject() (*SatelliteConfigConfigsGatewaysIp, error) {

	result := &SatelliteConfigConfigsGatewaysIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "client_config" || v == "ClientConfig" {
		return e.ClientConfig, nil
	}
	if v == "clientless_vpn" || v == "ClientlessVpn" {
		return e.ClientlessVpn, nil
	}
	if v == "portal_config" || v == "PortalConfig" {
		return e.PortalConfig, nil
	}
	if v == "satellite_config" || v == "SatelliteConfig" {
		return e.SatelliteConfig, nil
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
	if !o.ClientConfig.matches(other.ClientConfig) {
		return false
	}
	if !o.ClientlessVpn.matches(other.ClientlessVpn) {
		return false
	}
	if !o.PortalConfig.matches(other.PortalConfig) {
		return false
	}
	if !o.SatelliteConfig.matches(other.SatelliteConfig) {
		return false
	}

	return true
}

func (o *ClientConfig) matches(other *ClientConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AgentUserOverrideKey, other.AgentUserOverrideKey) {
		return false
	}
	if len(o.Configs) != len(other.Configs) {
		return false
	}
	for idx := range o.Configs {
		if !o.Configs[idx].matches(&other.Configs[idx]) {
			return false
		}
	}
	if len(o.RootCa) != len(other.RootCa) {
		return false
	}
	for idx := range o.RootCa {
		if !o.RootCa[idx].matches(&other.RootCa[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigs) matches(other *ClientConfigConfigs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.SaveUserCredentials, other.SaveUserCredentials) {
		return false
	}
	if !util.BoolsMatch(o.Portal2fa, other.Portal2fa) {
		return false
	}
	if !util.BoolsMatch(o.InternalGateway2fa, other.InternalGateway2fa) {
		return false
	}
	if !util.BoolsMatch(o.AutoDiscoveryExternalGateway2fa, other.AutoDiscoveryExternalGateway2fa) {
		return false
	}
	if !util.BoolsMatch(o.ManualOnlyGateway2fa, other.ManualOnlyGateway2fa) {
		return false
	}
	if !util.BoolsMatch(o.RefreshConfig, other.RefreshConfig) {
		return false
	}
	if !util.StringsMatch(o.MdmAddress, other.MdmAddress) {
		return false
	}
	if !util.StringsMatch(o.MdmEnrollmentPort, other.MdmEnrollmentPort) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceUser, other.SourceUser) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ThirdPartyVpnClients, other.ThirdPartyVpnClients) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Os, other.Os) {
		return false
	}
	if !o.Certificate.matches(other.Certificate) {
		return false
	}
	if !o.CustomChecks.matches(other.CustomChecks) {
		return false
	}
	if !o.Gateways.matches(other.Gateways) {
		return false
	}
	if !o.InternalHostDetection.matches(other.InternalHostDetection) {
		return false
	}
	if !o.InternalHostDetectionV6.matches(other.InternalHostDetectionV6) {
		return false
	}
	if !o.AgentUi.matches(other.AgentUi) {
		return false
	}
	if !o.HipCollection.matches(other.HipCollection) {
		return false
	}
	if !o.AgentConfig.matches(other.AgentConfig) {
		return false
	}
	if !o.GpAppConfig.matches(other.GpAppConfig) {
		return false
	}
	if !o.AuthenticationOverride.matches(other.AuthenticationOverride) {
		return false
	}
	if !o.MachineAccountExistsWithSerialno.matches(other.MachineAccountExistsWithSerialno) {
		return false
	}
	if !o.ClientCertificate.matches(other.ClientCertificate) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsCertificate) matches(other *ClientConfigConfigsCertificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Criteria.matches(other.Criteria) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsCertificateCriteria) matches(other *ClientConfigConfigsCertificateCriteria) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsCustomChecks) matches(other *ClientConfigConfigsCustomChecks) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Criteria.matches(other.Criteria) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsCustomChecksCriteria) matches(other *ClientConfigConfigsCustomChecksCriteria) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.RegistryKey) != len(other.RegistryKey) {
		return false
	}
	for idx := range o.RegistryKey {
		if !o.RegistryKey[idx].matches(&other.RegistryKey[idx]) {
			return false
		}
	}
	if len(o.Plist) != len(other.Plist) {
		return false
	}
	for idx := range o.Plist {
		if !o.Plist[idx].matches(&other.Plist[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsCustomChecksCriteriaRegistryKey) matches(other *ClientConfigConfigsCustomChecksCriteriaRegistryKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.DefaultValueData, other.DefaultValueData) {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}
	if len(o.RegistryValue) != len(other.RegistryValue) {
		return false
	}
	for idx := range o.RegistryValue {
		if !o.RegistryValue[idx].matches(&other.RegistryValue[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue) matches(other *ClientConfigConfigsCustomChecksCriteriaRegistryKeyRegistryValue) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.ValueData, other.ValueData) {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsCustomChecksCriteriaPlist) matches(other *ClientConfigConfigsCustomChecksCriteriaPlist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}
	if len(o.Key) != len(other.Key) {
		return false
	}
	for idx := range o.Key {
		if !o.Key[idx].matches(&other.Key[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsCustomChecksCriteriaPlistKey) matches(other *ClientConfigConfigsCustomChecksCriteriaPlistKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Value, other.Value) {
		return false
	}
	if !util.BoolsMatch(o.Negate, other.Negate) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGateways) matches(other *ClientConfigConfigsGateways) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Internal.matches(other.Internal) {
		return false
	}
	if !o.External.matches(other.External) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysInternal) matches(other *ClientConfigConfigsGatewaysInternal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.List) != len(other.List) {
		return false
	}
	for idx := range o.List {
		if !o.List[idx].matches(&other.List[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[int64](o.DhcpOptionCode, other.DhcpOptionCode) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysInternalList) matches(other *ClientConfigConfigsGatewaysInternalList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceIp, other.SourceIp) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysInternalListIp) matches(other *ClientConfigConfigsGatewaysInternalListIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysExternal) matches(other *ClientConfigConfigsGatewaysExternal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.CutoffTime, other.CutoffTime) {
		return false
	}
	if len(o.List) != len(other.List) {
		return false
	}
	for idx := range o.List {
		if !o.List[idx].matches(&other.List[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsGatewaysExternalList) matches(other *ClientConfigConfigsGatewaysExternalList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.PriorityRule) != len(other.PriorityRule) {
		return false
	}
	for idx := range o.PriorityRule {
		if !o.PriorityRule[idx].matches(&other.PriorityRule[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Manual, other.Manual) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysExternalListPriorityRule) matches(other *ClientConfigConfigsGatewaysExternalListPriorityRule) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Priority, other.Priority) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGatewaysExternalListIp) matches(other *ClientConfigConfigsGatewaysExternalListIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsInternalHostDetection) matches(other *ClientConfigConfigsInternalHostDetection) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsInternalHostDetectionV6) matches(other *ClientConfigConfigsInternalHostDetectionV6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.IpAddress, other.IpAddress) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAgentUi) matches(other *ClientConfigConfigsAgentUi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Passcode, other.Passcode) {
		return false
	}
	if !util.StringsMatch(o.UninstallPassword, other.UninstallPassword) {
		return false
	}
	if !util.Ints64Match(o.AgentUserOverrideTimeout, other.AgentUserOverrideTimeout) {
		return false
	}
	if !util.Ints64Match(o.MaxAgentUserOverrides, other.MaxAgentUserOverrides) {
		return false
	}
	if !o.WelcomePage.matches(other.WelcomePage) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAgentUiWelcomePage) matches(other *ClientConfigConfigsAgentUiWelcomePage) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Page, other.Page) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollection) matches(other *ClientConfigConfigsHipCollection) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.Ints64Match(o.MaxWaitTime, other.MaxWaitTime) {
		return false
	}
	if !util.BoolsMatch(o.CollectHipData, other.CollectHipData) {
		return false
	}
	if !o.Exclusion.matches(other.Exclusion) {
		return false
	}
	if !o.CustomChecks.matches(other.CustomChecks) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionExclusion) matches(other *ClientConfigConfigsHipCollectionExclusion) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Category) != len(other.Category) {
		return false
	}
	for idx := range o.Category {
		if !o.Category[idx].matches(&other.Category[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionExclusionCategory) matches(other *ClientConfigConfigsHipCollectionExclusionCategory) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.Vendor) != len(other.Vendor) {
		return false
	}
	for idx := range o.Vendor {
		if !o.Vendor[idx].matches(&other.Vendor[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionExclusionCategoryVendor) matches(other *ClientConfigConfigsHipCollectionExclusionCategoryVendor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Product, other.Product) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecks) matches(other *ClientConfigConfigsHipCollectionCustomChecks) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Windows.matches(other.Windows) {
		return false
	}
	if !o.MacOs.matches(other.MacOs) {
		return false
	}
	if !o.Linux.matches(other.Linux) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecksWindows) matches(other *ClientConfigConfigsHipCollectionCustomChecksWindows) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.RegistryKey) != len(other.RegistryKey) {
		return false
	}
	for idx := range o.RegistryKey {
		if !o.RegistryKey[idx].matches(&other.RegistryKey[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.ProcessList, other.ProcessList) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey) matches(other *ClientConfigConfigsHipCollectionCustomChecksWindowsRegistryKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegistryValue, other.RegistryValue) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecksMacOs) matches(other *ClientConfigConfigsHipCollectionCustomChecksMacOs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Plist) != len(other.Plist) {
		return false
	}
	for idx := range o.Plist {
		if !o.Plist[idx].matches(&other.Plist[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.ProcessList, other.ProcessList) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecksMacOsPlist) matches(other *ClientConfigConfigsHipCollectionCustomChecksMacOsPlist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsHipCollectionCustomChecksLinux) matches(other *ClientConfigConfigsHipCollectionCustomChecksLinux) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ProcessList, other.ProcessList) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAgentConfig) matches(other *ClientConfigConfigsAgentConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsGpAppConfig) matches(other *ClientConfigConfigsGpAppConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Config) != len(other.Config) {
		return false
	}
	for idx := range o.Config {
		if !o.Config[idx].matches(&other.Config[idx]) {
			return false
		}
	}

	return true
}

func (o *ClientConfigConfigsGpAppConfigConfig) matches(other *ClientConfigConfigsGpAppConfigConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Value, other.Value) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAuthenticationOverride) matches(other *ClientConfigConfigsAuthenticationOverride) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.GenerateCookie, other.GenerateCookie) {
		return false
	}
	if !util.StringsMatch(o.CookieEncryptDecryptCert, other.CookieEncryptDecryptCert) {
		return false
	}
	if !o.AcceptCookie.matches(other.AcceptCookie) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAuthenticationOverrideAcceptCookie) matches(other *ClientConfigConfigsAuthenticationOverrideAcceptCookie) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.CookieLifetime.matches(other.CookieLifetime) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime) matches(other *ClientConfigConfigsAuthenticationOverrideAcceptCookieCookieLifetime) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LifetimeInDays, other.LifetimeInDays) {
		return false
	}
	if !util.Ints64Match(o.LifetimeInHours, other.LifetimeInHours) {
		return false
	}
	if !util.Ints64Match(o.LifetimeInMinutes, other.LifetimeInMinutes) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsMachineAccountExistsWithSerialno) matches(other *ClientConfigConfigsMachineAccountExistsWithSerialno) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.No.matches(other.No) {
		return false
	}
	if !o.Yes.matches(other.Yes) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsMachineAccountExistsWithSerialnoNo) matches(other *ClientConfigConfigsMachineAccountExistsWithSerialnoNo) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsMachineAccountExistsWithSerialnoYes) matches(other *ClientConfigConfigsMachineAccountExistsWithSerialnoYes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ClientConfigConfigsClientCertificate) matches(other *ClientConfigConfigsClientCertificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Local, other.Local) {
		return false
	}
	if !util.StringsMatch(o.Scep, other.Scep) {
		return false
	}

	return true
}

func (o *ClientConfigRootCa) matches(other *ClientConfigRootCa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.InstallInCertStore, other.InstallInCertStore) {
		return false
	}

	return true
}

func (o *ClientlessVpn) matches(other *ClientlessVpn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.AppsToUserMapping) != len(other.AppsToUserMapping) {
		return false
	}
	for idx := range o.AppsToUserMapping {
		if !o.AppsToUserMapping[idx].matches(&other.AppsToUserMapping[idx]) {
			return false
		}
	}
	if !o.CryptoSettings.matches(other.CryptoSettings) {
		return false
	}
	if !util.StringsMatch(o.DnsProxy, other.DnsProxy) {
		return false
	}
	if !util.StringsMatch(o.Hostname, other.Hostname) {
		return false
	}
	if !o.InactivityLogout.matches(other.InactivityLogout) {
		return false
	}
	if !o.LoginLifetime.matches(other.LoginLifetime) {
		return false
	}
	if !util.Ints64Match(o.MaxUser, other.MaxUser) {
		return false
	}
	if len(o.ProxyServerSetting) != len(other.ProxyServerSetting) {
		return false
	}
	for idx := range o.ProxyServerSetting {
		if !o.ProxyServerSetting[idx].matches(&other.ProxyServerSetting[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.RewriteExcludeDomainList, other.RewriteExcludeDomainList) {
		return false
	}
	if !util.StringsMatch(o.SecurityZone, other.SecurityZone) {
		return false
	}

	return true
}

func (o *ClientlessVpnAppsToUserMapping) matches(other *ClientlessVpnAppsToUserMapping) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceUser, other.SourceUser) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Applications, other.Applications) {
		return false
	}
	if !util.BoolsMatch(o.EnableCustomAppURLAddressBar, other.EnableCustomAppURLAddressBar) {
		return false
	}
	if !util.BoolsMatch(o.DisplayGlobalProtectAgentDownloadLink, other.DisplayGlobalProtectAgentDownloadLink) {
		return false
	}

	return true
}

func (o *ClientlessVpnCryptoSettings) matches(other *ClientlessVpnCryptoSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.ServerCertVerification.matches(other.ServerCertVerification) {
		return false
	}
	if !o.SslProtocol.matches(other.SslProtocol) {
		return false
	}

	return true
}

func (o *ClientlessVpnCryptoSettingsServerCertVerification) matches(other *ClientlessVpnCryptoSettingsServerCertVerification) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
	if !util.BoolsMatch(o.BlockUntrustedIssuer, other.BlockUntrustedIssuer) {
		return false
	}

	return true
}

func (o *ClientlessVpnCryptoSettingsSslProtocol) matches(other *ClientlessVpnCryptoSettingsSslProtocol) bool {
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

func (o *ClientlessVpnInactivityLogout) matches(other *ClientlessVpnInactivityLogout) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Hours, other.Hours) {
		return false
	}
	if !util.Ints64Match(o.Minutes, other.Minutes) {
		return false
	}

	return true
}

func (o *ClientlessVpnLoginLifetime) matches(other *ClientlessVpnLoginLifetime) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Hours, other.Hours) {
		return false
	}
	if !util.Ints64Match(o.Minutes, other.Minutes) {
		return false
	}

	return true
}

func (o *ClientlessVpnProxyServerSetting) matches(other *ClientlessVpnProxyServerSetting) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Domains, other.Domains) {
		return false
	}
	if !util.BoolsMatch(o.UseProxy, other.UseProxy) {
		return false
	}
	if !o.ProxyServer.matches(other.ProxyServer) {
		return false
	}

	return true
}

func (o *ClientlessVpnProxyServerSettingProxyServer) matches(other *ClientlessVpnProxyServerSettingProxyServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Server, other.Server) {
		return false
	}
	if !util.Ints64Match(o.Port, other.Port) {
		return false
	}
	if !util.StringsMatch(o.User, other.User) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}

	return true
}

func (o *PortalConfig) matches(other *PortalConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if len(o.ClientAuth) != len(other.ClientAuth) {
		return false
	}
	for idx := range o.ClientAuth {
		if !o.ClientAuth[idx].matches(&other.ClientAuth[idx]) {
			return false
		}
	}
	if !o.ConfigSelection.matches(other.ConfigSelection) {
		return false
	}
	if !util.StringsMatch(o.CustomHelpPage, other.CustomHelpPage) {
		return false
	}
	if !util.StringsMatch(o.CustomHomePage, other.CustomHomePage) {
		return false
	}
	if !util.StringsMatch(o.CustomLoginPage, other.CustomLoginPage) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !util.BoolsMatch(o.LogFail, other.LogFail) {
		return false
	}
	if !util.StringsMatch(o.LogSetting, other.LogSetting) {
		return false
	}
	if !util.BoolsMatch(o.LogSuccess, other.LogSuccess) {
		return false
	}
	if !util.StringsMatch(o.SslTlsServiceProfile, other.SslTlsServiceProfile) {
		return false
	}

	return true
}

func (o *PortalConfigClientAuth) matches(other *PortalConfigClientAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Os, other.Os) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationProfile, other.AuthenticationProfile) {
		return false
	}
	if !util.BoolsMatch(o.AutoRetrievePasscode, other.AutoRetrievePasscode) {
		return false
	}
	if !util.StringsMatch(o.UsernameLabel, other.UsernameLabel) {
		return false
	}
	if !util.StringsMatch(o.PasswordLabel, other.PasswordLabel) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationMessage, other.AuthenticationMessage) {
		return false
	}
	if !util.StringsMatch(o.UserCredentialOrClientCertRequired, other.UserCredentialOrClientCertRequired) {
		return false
	}

	return true
}

func (o *PortalConfigConfigSelection) matches(other *PortalConfigConfigSelection) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !o.CustomChecks.matches(other.CustomChecks) {
		return false
	}

	return true
}

func (o *PortalConfigConfigSelectionCustomChecks) matches(other *PortalConfigConfigSelectionCustomChecks) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.MacOs.matches(other.MacOs) {
		return false
	}
	if !o.Windows.matches(other.Windows) {
		return false
	}

	return true
}

func (o *PortalConfigConfigSelectionCustomChecksMacOs) matches(other *PortalConfigConfigSelectionCustomChecksMacOs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Plist) != len(other.Plist) {
		return false
	}
	for idx := range o.Plist {
		if !o.Plist[idx].matches(&other.Plist[idx]) {
			return false
		}
	}

	return true
}

func (o *PortalConfigConfigSelectionCustomChecksMacOsPlist) matches(other *PortalConfigConfigSelectionCustomChecksMacOsPlist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Key, other.Key) {
		return false
	}

	return true
}

func (o *PortalConfigConfigSelectionCustomChecksWindows) matches(other *PortalConfigConfigSelectionCustomChecksWindows) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.RegistryKey) != len(other.RegistryKey) {
		return false
	}
	for idx := range o.RegistryKey {
		if !o.RegistryKey[idx].matches(&other.RegistryKey[idx]) {
			return false
		}
	}

	return true
}

func (o *PortalConfigConfigSelectionCustomChecksWindowsRegistryKey) matches(other *PortalConfigConfigSelectionCustomChecksWindowsRegistryKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.RegistryValue, other.RegistryValue) {
		return false
	}

	return true
}

func (o *PortalConfigLocalAddress) matches(other *PortalConfigLocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.IpAddressFamily, other.IpAddressFamily) {
		return false
	}
	if !o.FloatingIp.matches(other.FloatingIp) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}

	return true
}

func (o *PortalConfigLocalAddressFloatingIp) matches(other *PortalConfigLocalAddressFloatingIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *PortalConfigLocalAddressIp) matches(other *PortalConfigLocalAddressIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *SatelliteConfig) matches(other *SatelliteConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.ClientCertificate.matches(other.ClientCertificate) {
		return false
	}
	if len(o.Configs) != len(other.Configs) {
		return false
	}
	for idx := range o.Configs {
		if !o.Configs[idx].matches(&other.Configs[idx]) {
			return false
		}
	}
	if !util.OrderedListsMatch[string](o.RootCa, other.RootCa) {
		return false
	}

	return true
}

func (o *SatelliteConfigClientCertificate) matches(other *SatelliteConfigClientCertificate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Local.matches(other.Local) {
		return false
	}
	if !o.Scep.matches(other.Scep) {
		return false
	}

	return true
}

func (o *SatelliteConfigClientCertificateLocal) matches(other *SatelliteConfigClientCertificateLocal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.CertificateLifeTime, other.CertificateLifeTime) {
		return false
	}
	if !util.Ints64Match(o.CertificateRenewalPeriod, other.CertificateRenewalPeriod) {
		return false
	}
	if !util.StringsMatch(o.IssuingCertificate, other.IssuingCertificate) {
		return false
	}
	if !util.StringsMatch(o.OcspResponder, other.OcspResponder) {
		return false
	}

	return true
}

func (o *SatelliteConfigClientCertificateScep) matches(other *SatelliteConfigClientCertificateScep) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.CertificateRenewalPeriod, other.CertificateRenewalPeriod) {
		return false
	}
	if !util.StringsMatch(o.Scep, other.Scep) {
		return false
	}

	return true
}

func (o *SatelliteConfigConfigs) matches(other *SatelliteConfigConfigs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Devices, other.Devices) {
		return false
	}
	if !util.OrderedListsMatch[string](o.SourceUser, other.SourceUser) {
		return false
	}
	if len(o.Gateways) != len(other.Gateways) {
		return false
	}
	for idx := range o.Gateways {
		if !o.Gateways[idx].matches(&other.Gateways[idx]) {
			return false
		}
	}
	if !util.Ints64Match(o.ConfigRefreshInterval, other.ConfigRefreshInterval) {
		return false
	}

	return true
}

func (o *SatelliteConfigConfigsGateways) matches(other *SatelliteConfigConfigsGateways) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Ipv6Preferred, other.Ipv6Preferred) {
		return false
	}
	if !util.Ints64Match(o.Priority, other.Priority) {
		return false
	}
	if !util.StringsMatch(o.Fqdn, other.Fqdn) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}

	return true
}

func (o *SatelliteConfigConfigsGatewaysIp) matches(other *SatelliteConfigConfigsGatewaysIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
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
