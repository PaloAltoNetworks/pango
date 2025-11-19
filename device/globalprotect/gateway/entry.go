package gateway

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
	suffix = []string{"global-protect", "global-protect-gateway", "$name"}
)

type Entry struct {
	Name                    string
	BlockQuarantinedDevices *bool
	CertificateProfile      *string
	ClientAuth              []ClientAuth
	HipNotification         []HipNotification
	LocalAddress            *LocalAddress
	LogFail                 *bool
	LogSetting              *string
	LogSuccess              *bool
	RemoteUserTunnel        *string
	RemoteUserTunnelConfigs []RemoteUserTunnelConfigs
	Roles                   []Roles
	SatelliteTunnel         *string
	SecurityRestrictions    *SecurityRestrictions
	SslTlsServiceProfile    *string
	TunnelMode              *bool
	Misc                    []generic.Xml
	MiscAttributes          []xml.Attr
}
type ClientAuth struct {
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
type HipNotification struct {
	Name            string
	MatchMessage    *HipNotificationMatchMessage
	NotMatchMessage *HipNotificationNotMatchMessage
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type HipNotificationMatchMessage struct {
	IncludeAppList     *bool
	ShowNotificationAs *string
	Message            *string
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type HipNotificationNotMatchMessage struct {
	ShowNotificationAs *string
	Message            *string
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type LocalAddress struct {
	Interface       *string
	IpAddressFamily *string
	FloatingIp      *LocalAddressFloatingIp
	Ip              *LocalAddressIp
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type LocalAddressFloatingIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type LocalAddressIp struct {
	Ipv4           *string
	Ipv6           *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigs struct {
	Name                         string
	SourceUser                   []string
	Os                           []string
	DnsServer                    []string
	DnsSuffix                    []string
	IpPool                       []string
	AuthenticationServerIpPool   []string
	AuthenticationOverride       *RemoteUserTunnelConfigsAuthenticationOverride
	SourceAddress                *RemoteUserTunnelConfigsSourceAddress
	SplitTunneling               *RemoteUserTunnelConfigsSplitTunneling
	NoDirectAccessToLocalNetwork *bool
	RetrieveFramedIpAddress      *bool
	Misc                         []generic.Xml
	MiscAttributes               []xml.Attr
}
type RemoteUserTunnelConfigsAuthenticationOverride struct {
	GenerateCookie           *bool
	CookieEncryptDecryptCert *string
	AcceptCookie             *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie struct {
	CookieLifetime *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime struct {
	LifetimeInDays    *int64
	LifetimeInHours   *int64
	LifetimeInMinutes *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type RemoteUserTunnelConfigsSourceAddress struct {
	Region         []string
	IpAddress      []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigsSplitTunneling struct {
	AccessRoute         []string
	ExcludeAccessRoute  []string
	IncludeApplications []string
	ExcludeApplications []string
	IncludeDomains      *RemoteUserTunnelConfigsSplitTunnelingIncludeDomains
	ExcludeDomains      *RemoteUserTunnelConfigsSplitTunnelingExcludeDomains
	Misc                []generic.Xml
	MiscAttributes      []xml.Attr
}
type RemoteUserTunnelConfigsSplitTunnelingIncludeDomains struct {
	List           []RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList struct {
	Name           string
	Ports          []int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigsSplitTunnelingExcludeDomains struct {
	List           []RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList struct {
	Name           string
	Ports          []int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Roles struct {
	Name                     string
	LoginLifetime            *RolesLoginLifetime
	InactivityLogout         *int64
	LifetimeNotifyPrior      *int64
	LifetimeNotifyMessage    *string
	InactivityNotifyPrior    *int64
	InactivityNotifyMessage  *string
	AdminLogoutNotify        *bool
	AdminLogoutNotifyMessage *string
	Misc                     []generic.Xml
	MiscAttributes           []xml.Attr
}
type RolesLoginLifetime struct {
	Minutes        *int64
	Hours          *int64
	Days           *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SecurityRestrictions struct {
	DisallowAutomaticRestoration *bool
	SourceIpEnforcement          *SecurityRestrictionsSourceIpEnforcement
	Misc                         []generic.Xml
	MiscAttributes               []xml.Attr
}
type SecurityRestrictionsSourceIpEnforcement struct {
	Enable         *bool
	Custom         *SecurityRestrictionsSourceIpEnforcementCustom
	Default        *SecurityRestrictionsSourceIpEnforcementDefault
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SecurityRestrictionsSourceIpEnforcementCustom struct {
	SourceIpv4Netmask *int64
	SourceIpv6Netmask *int64
	Misc              []generic.Xml
	MiscAttributes    []xml.Attr
}
type SecurityRestrictionsSourceIpEnforcementDefault struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	XMLName                 xml.Name                             `xml:"entry"`
	Name                    string                               `xml:"name,attr"`
	BlockQuarantinedDevices *string                              `xml:"block-quarantined-devices,omitempty"`
	CertificateProfile      *string                              `xml:"certificate-profile,omitempty"`
	ClientAuth              *clientAuthContainerXml              `xml:"client-auth,omitempty"`
	HipNotification         *hipNotificationContainerXml         `xml:"hip-notification,omitempty"`
	LocalAddress            *localAddressXml                     `xml:"local-address,omitempty"`
	LogFail                 *string                              `xml:"log-fail,omitempty"`
	LogSetting              *string                              `xml:"log-setting,omitempty"`
	LogSuccess              *string                              `xml:"log-success,omitempty"`
	RemoteUserTunnel        *string                              `xml:"remote-user-tunnel,omitempty"`
	RemoteUserTunnelConfigs *remoteUserTunnelConfigsContainerXml `xml:"remote-user-tunnel-configs,omitempty"`
	Roles                   *rolesContainerXml                   `xml:"roles,omitempty"`
	SatelliteTunnel         *string                              `xml:"satellite-tunnel,omitempty"`
	SecurityRestrictions    *securityRestrictionsXml             `xml:"security-restrictions,omitempty"`
	SslTlsServiceProfile    *string                              `xml:"ssl-tls-service-profile,omitempty"`
	TunnelMode              *string                              `xml:"tunnel-mode,omitempty"`
	Misc                    []generic.Xml                        `xml:",any"`
	MiscAttributes          []xml.Attr                           `xml:",any,attr"`
}
type clientAuthContainerXml struct {
	Entries []clientAuthXml `xml:"entry"`
}
type clientAuthXml struct {
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
type hipNotificationContainerXml struct {
	Entries []hipNotificationXml `xml:"entry"`
}
type hipNotificationXml struct {
	XMLName         xml.Name                           `xml:"entry"`
	Name            string                             `xml:"name,attr"`
	MatchMessage    *hipNotificationMatchMessageXml    `xml:"match-message,omitempty"`
	NotMatchMessage *hipNotificationNotMatchMessageXml `xml:"not-match-message,omitempty"`
	Misc            []generic.Xml                      `xml:",any"`
	MiscAttributes  []xml.Attr                         `xml:",any,attr"`
}
type hipNotificationMatchMessageXml struct {
	IncludeAppList     *string       `xml:"include-app-list,omitempty"`
	ShowNotificationAs *string       `xml:"show-notification-as,omitempty"`
	Message            *string       `xml:"message,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type hipNotificationNotMatchMessageXml struct {
	ShowNotificationAs *string       `xml:"show-notification-as,omitempty"`
	Message            *string       `xml:"message,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type localAddressXml struct {
	Interface       *string                    `xml:"interface,omitempty"`
	IpAddressFamily *string                    `xml:"ip-address-family,omitempty"`
	FloatingIp      *localAddressFloatingIpXml `xml:"floating-ip,omitempty"`
	Ip              *localAddressIpXml         `xml:"ip,omitempty"`
	Misc            []generic.Xml              `xml:",any"`
	MiscAttributes  []xml.Attr                 `xml:",any,attr"`
}
type localAddressFloatingIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type localAddressIpXml struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type remoteUserTunnelConfigsContainerXml struct {
	Entries []remoteUserTunnelConfigsXml `xml:"entry"`
}
type remoteUserTunnelConfigsXml struct {
	XMLName                      xml.Name                                          `xml:"entry"`
	Name                         string                                            `xml:"name,attr"`
	SourceUser                   *util.MemberType                                  `xml:"source-user,omitempty"`
	Os                           *util.MemberType                                  `xml:"os,omitempty"`
	DnsServer                    *util.MemberType                                  `xml:"dns-server,omitempty"`
	DnsSuffix                    *util.MemberType                                  `xml:"dns-suffix,omitempty"`
	IpPool                       *util.MemberType                                  `xml:"ip-pool,omitempty"`
	AuthenticationServerIpPool   *util.MemberType                                  `xml:"authentication-server-ip-pool,omitempty"`
	AuthenticationOverride       *remoteUserTunnelConfigsAuthenticationOverrideXml `xml:"authentication-override,omitempty"`
	SourceAddress                *remoteUserTunnelConfigsSourceAddressXml          `xml:"source-address,omitempty"`
	SplitTunneling               *remoteUserTunnelConfigsSplitTunnelingXml         `xml:"split-tunneling,omitempty"`
	NoDirectAccessToLocalNetwork *string                                           `xml:"no-direct-access-to-local-network,omitempty"`
	RetrieveFramedIpAddress      *string                                           `xml:"retrieve-framed-ip-address,omitempty"`
	Misc                         []generic.Xml                                     `xml:",any"`
	MiscAttributes               []xml.Attr                                        `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideXml struct {
	GenerateCookie           *string                                                       `xml:"generate-cookie,omitempty"`
	CookieEncryptDecryptCert *string                                                       `xml:"cookie-encrypt-decrypt-cert,omitempty"`
	AcceptCookie             *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml `xml:"accept-cookie,omitempty"`
	Misc                     []generic.Xml                                                 `xml:",any"`
	MiscAttributes           []xml.Attr                                                    `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml struct {
	CookieLifetime *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml `xml:"cookie-lifetime,omitempty"`
	Misc           []generic.Xml                                                               `xml:",any"`
	MiscAttributes []xml.Attr                                                                  `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml struct {
	LifetimeInDays    *int64        `xml:"lifetime-in-days,omitempty"`
	LifetimeInHours   *int64        `xml:"lifetime-in-hours,omitempty"`
	LifetimeInMinutes *int64        `xml:"lifetime-in-minutes,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type remoteUserTunnelConfigsSourceAddressXml struct {
	Region         *util.MemberType `xml:"region,omitempty"`
	IpAddress      *util.MemberType `xml:"ip-address,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingXml struct {
	AccessRoute         *util.MemberType                                        `xml:"access-route,omitempty"`
	ExcludeAccessRoute  *util.MemberType                                        `xml:"exclude-access-route,omitempty"`
	IncludeApplications *util.MemberType                                        `xml:"include-applications,omitempty"`
	ExcludeApplications *util.MemberType                                        `xml:"exclude-applications,omitempty"`
	IncludeDomains      *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml `xml:"include-domains,omitempty"`
	ExcludeDomains      *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml `xml:"exclude-domains,omitempty"`
	Misc                []generic.Xml                                           `xml:",any"`
	MiscAttributes      []xml.Attr                                              `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml struct {
	List           *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml `xml:"list,omitempty"`
	Misc           []generic.Xml                                                        `xml:",any"`
	MiscAttributes []xml.Attr                                                           `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml struct {
	Entries []remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml `xml:"entry"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Ports          *util.MemberType `xml:"ports,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml struct {
	List           *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml `xml:"list,omitempty"`
	Misc           []generic.Xml                                                        `xml:",any"`
	MiscAttributes []xml.Attr                                                           `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml struct {
	Entries []remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml `xml:"entry"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Ports          *util.MemberType `xml:"ports,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type rolesContainerXml struct {
	Entries []rolesXml `xml:"entry"`
}
type rolesXml struct {
	XMLName                  xml.Name               `xml:"entry"`
	Name                     string                 `xml:"name,attr"`
	LoginLifetime            *rolesLoginLifetimeXml `xml:"login-lifetime,omitempty"`
	InactivityLogout         *int64                 `xml:"inactivity-logout,omitempty"`
	LifetimeNotifyPrior      *int64                 `xml:"lifetime-notify-prior,omitempty"`
	LifetimeNotifyMessage    *string                `xml:"lifetime-notify-message,omitempty"`
	InactivityNotifyPrior    *int64                 `xml:"inactivity-notify-prior,omitempty"`
	InactivityNotifyMessage  *string                `xml:"inactivity-notify-message,omitempty"`
	AdminLogoutNotify        *string                `xml:"admin-logout-notify,omitempty"`
	AdminLogoutNotifyMessage *string                `xml:"admin-logout-notify-message,omitempty"`
	Misc                     []generic.Xml          `xml:",any"`
	MiscAttributes           []xml.Attr             `xml:",any,attr"`
}
type rolesLoginLifetimeXml struct {
	Minutes        *int64        `xml:"minutes,omitempty"`
	Hours          *int64        `xml:"hours,omitempty"`
	Days           *int64        `xml:"days,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type securityRestrictionsXml struct {
	DisallowAutomaticRestoration *string                                     `xml:"disallow-automatic-restoration,omitempty"`
	SourceIpEnforcement          *securityRestrictionsSourceIpEnforcementXml `xml:"source-ip-enforcement,omitempty"`
	Misc                         []generic.Xml                               `xml:",any"`
	MiscAttributes               []xml.Attr                                  `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementXml struct {
	Enable         *string                                            `xml:"enable,omitempty"`
	Custom         *securityRestrictionsSourceIpEnforcementCustomXml  `xml:"custom,omitempty"`
	Default        *securityRestrictionsSourceIpEnforcementDefaultXml `xml:"default,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementCustomXml struct {
	SourceIpv4Netmask *int64        `xml:"source-ipv4-netmask,omitempty"`
	SourceIpv6Netmask *int64        `xml:"source-ipv6-netmask,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementDefaultXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type entryXml_11_0_2 struct {
	XMLName                 xml.Name                                    `xml:"entry"`
	Name                    string                                      `xml:"name,attr"`
	BlockQuarantinedDevices *string                                     `xml:"block-quarantined-devices,omitempty"`
	CertificateProfile      *string                                     `xml:"certificate-profile,omitempty"`
	ClientAuth              *clientAuthContainerXml_11_0_2              `xml:"client-auth,omitempty"`
	HipNotification         *hipNotificationContainerXml_11_0_2         `xml:"hip-notification,omitempty"`
	LocalAddress            *localAddressXml_11_0_2                     `xml:"local-address,omitempty"`
	LogFail                 *string                                     `xml:"log-fail,omitempty"`
	LogSetting              *string                                     `xml:"log-setting,omitempty"`
	LogSuccess              *string                                     `xml:"log-success,omitempty"`
	RemoteUserTunnel        *string                                     `xml:"remote-user-tunnel,omitempty"`
	RemoteUserTunnelConfigs *remoteUserTunnelConfigsContainerXml_11_0_2 `xml:"remote-user-tunnel-configs,omitempty"`
	Roles                   *rolesContainerXml_11_0_2                   `xml:"roles,omitempty"`
	SatelliteTunnel         *string                                     `xml:"satellite-tunnel,omitempty"`
	SecurityRestrictions    *securityRestrictionsXml_11_0_2             `xml:"security-restrictions,omitempty"`
	SslTlsServiceProfile    *string                                     `xml:"ssl-tls-service-profile,omitempty"`
	TunnelMode              *string                                     `xml:"tunnel-mode,omitempty"`
	Misc                    []generic.Xml                               `xml:",any"`
	MiscAttributes          []xml.Attr                                  `xml:",any,attr"`
}
type clientAuthContainerXml_11_0_2 struct {
	Entries []clientAuthXml_11_0_2 `xml:"entry"`
}
type clientAuthXml_11_0_2 struct {
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
type hipNotificationContainerXml_11_0_2 struct {
	Entries []hipNotificationXml_11_0_2 `xml:"entry"`
}
type hipNotificationXml_11_0_2 struct {
	XMLName         xml.Name                                  `xml:"entry"`
	Name            string                                    `xml:"name,attr"`
	MatchMessage    *hipNotificationMatchMessageXml_11_0_2    `xml:"match-message,omitempty"`
	NotMatchMessage *hipNotificationNotMatchMessageXml_11_0_2 `xml:"not-match-message,omitempty"`
	Misc            []generic.Xml                             `xml:",any"`
	MiscAttributes  []xml.Attr                                `xml:",any,attr"`
}
type hipNotificationMatchMessageXml_11_0_2 struct {
	IncludeAppList     *string       `xml:"include-app-list,omitempty"`
	ShowNotificationAs *string       `xml:"show-notification-as,omitempty"`
	Message            *string       `xml:"message,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type hipNotificationNotMatchMessageXml_11_0_2 struct {
	ShowNotificationAs *string       `xml:"show-notification-as,omitempty"`
	Message            *string       `xml:"message,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type localAddressXml_11_0_2 struct {
	Interface       *string                           `xml:"interface,omitempty"`
	IpAddressFamily *string                           `xml:"ip-address-family,omitempty"`
	FloatingIp      *localAddressFloatingIpXml_11_0_2 `xml:"floating-ip,omitempty"`
	Ip              *localAddressIpXml_11_0_2         `xml:"ip,omitempty"`
	Misc            []generic.Xml                     `xml:",any"`
	MiscAttributes  []xml.Attr                        `xml:",any,attr"`
}
type localAddressFloatingIpXml_11_0_2 struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type localAddressIpXml_11_0_2 struct {
	Ipv4           *string       `xml:"ipv4,omitempty"`
	Ipv6           *string       `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type remoteUserTunnelConfigsContainerXml_11_0_2 struct {
	Entries []remoteUserTunnelConfigsXml_11_0_2 `xml:"entry"`
}
type remoteUserTunnelConfigsXml_11_0_2 struct {
	XMLName                      xml.Name                                                 `xml:"entry"`
	Name                         string                                                   `xml:"name,attr"`
	SourceUser                   *util.MemberType                                         `xml:"source-user,omitempty"`
	Os                           *util.MemberType                                         `xml:"os,omitempty"`
	DnsServer                    *util.MemberType                                         `xml:"dns-server,omitempty"`
	DnsSuffix                    *util.MemberType                                         `xml:"dns-suffix,omitempty"`
	IpPool                       *util.MemberType                                         `xml:"ip-pool,omitempty"`
	AuthenticationServerIpPool   *util.MemberType                                         `xml:"authentication-server-ip-pool,omitempty"`
	AuthenticationOverride       *remoteUserTunnelConfigsAuthenticationOverrideXml_11_0_2 `xml:"authentication-override,omitempty"`
	SourceAddress                *remoteUserTunnelConfigsSourceAddressXml_11_0_2          `xml:"source-address,omitempty"`
	SplitTunneling               *remoteUserTunnelConfigsSplitTunnelingXml_11_0_2         `xml:"split-tunneling,omitempty"`
	NoDirectAccessToLocalNetwork *string                                                  `xml:"no-direct-access-to-local-network,omitempty"`
	RetrieveFramedIpAddress      *string                                                  `xml:"retrieve-framed-ip-address,omitempty"`
	Misc                         []generic.Xml                                            `xml:",any"`
	MiscAttributes               []xml.Attr                                               `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideXml_11_0_2 struct {
	GenerateCookie           *string                                                              `xml:"generate-cookie,omitempty"`
	CookieEncryptDecryptCert *string                                                              `xml:"cookie-encrypt-decrypt-cert,omitempty"`
	AcceptCookie             *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml_11_0_2 `xml:"accept-cookie,omitempty"`
	Misc                     []generic.Xml                                                        `xml:",any"`
	MiscAttributes           []xml.Attr                                                           `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml_11_0_2 struct {
	CookieLifetime *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml_11_0_2 `xml:"cookie-lifetime,omitempty"`
	Misc           []generic.Xml                                                                      `xml:",any"`
	MiscAttributes []xml.Attr                                                                         `xml:",any,attr"`
}
type remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml_11_0_2 struct {
	LifetimeInDays    *int64        `xml:"lifetime-in-days,omitempty"`
	LifetimeInHours   *int64        `xml:"lifetime-in-hours,omitempty"`
	LifetimeInMinutes *int64        `xml:"lifetime-in-minutes,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type remoteUserTunnelConfigsSourceAddressXml_11_0_2 struct {
	Region         *util.MemberType `xml:"region,omitempty"`
	IpAddress      *util.MemberType `xml:"ip-address,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingXml_11_0_2 struct {
	AccessRoute         *util.MemberType                                               `xml:"access-route,omitempty"`
	ExcludeAccessRoute  *util.MemberType                                               `xml:"exclude-access-route,omitempty"`
	IncludeApplications *util.MemberType                                               `xml:"include-applications,omitempty"`
	ExcludeApplications *util.MemberType                                               `xml:"exclude-applications,omitempty"`
	IncludeDomains      *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml_11_0_2 `xml:"include-domains,omitempty"`
	ExcludeDomains      *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml_11_0_2 `xml:"exclude-domains,omitempty"`
	Misc                []generic.Xml                                                  `xml:",any"`
	MiscAttributes      []xml.Attr                                                     `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml_11_0_2 struct {
	List           *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml_11_0_2 `xml:"list,omitempty"`
	Misc           []generic.Xml                                                               `xml:",any"`
	MiscAttributes []xml.Attr                                                                  `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml_11_0_2 struct {
	Entries []remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2 `xml:"entry"`
}
type remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2 struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Ports          *util.MemberType `xml:"ports,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml_11_0_2 struct {
	List           *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml_11_0_2 `xml:"list,omitempty"`
	Misc           []generic.Xml                                                               `xml:",any"`
	MiscAttributes []xml.Attr                                                                  `xml:",any,attr"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml_11_0_2 struct {
	Entries []remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2 `xml:"entry"`
}
type remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2 struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Ports          *util.MemberType `xml:"ports,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type rolesContainerXml_11_0_2 struct {
	Entries []rolesXml_11_0_2 `xml:"entry"`
}
type rolesXml_11_0_2 struct {
	XMLName                  xml.Name                      `xml:"entry"`
	Name                     string                        `xml:"name,attr"`
	LoginLifetime            *rolesLoginLifetimeXml_11_0_2 `xml:"login-lifetime,omitempty"`
	InactivityLogout         *int64                        `xml:"inactivity-logout,omitempty"`
	LifetimeNotifyPrior      *int64                        `xml:"lifetime-notify-prior,omitempty"`
	LifetimeNotifyMessage    *string                       `xml:"lifetime-notify-message,omitempty"`
	InactivityNotifyPrior    *int64                        `xml:"inactivity-notify-prior,omitempty"`
	InactivityNotifyMessage  *string                       `xml:"inactivity-notify-message,omitempty"`
	AdminLogoutNotify        *string                       `xml:"admin-logout-notify,omitempty"`
	AdminLogoutNotifyMessage *string                       `xml:"admin-logout-notify-message,omitempty"`
	Misc                     []generic.Xml                 `xml:",any"`
	MiscAttributes           []xml.Attr                    `xml:",any,attr"`
}
type rolesLoginLifetimeXml_11_0_2 struct {
	Minutes        *int64        `xml:"minutes,omitempty"`
	Hours          *int64        `xml:"hours,omitempty"`
	Days           *int64        `xml:"days,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type securityRestrictionsXml_11_0_2 struct {
	DisallowAutomaticRestoration *string                                            `xml:"disallow-automatic-restoration,omitempty"`
	SourceIpEnforcement          *securityRestrictionsSourceIpEnforcementXml_11_0_2 `xml:"source-ip-enforcement,omitempty"`
	Misc                         []generic.Xml                                      `xml:",any"`
	MiscAttributes               []xml.Attr                                         `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementXml_11_0_2 struct {
	Enable         *string                                                   `xml:"enable,omitempty"`
	Custom         *securityRestrictionsSourceIpEnforcementCustomXml_11_0_2  `xml:"custom,omitempty"`
	Default        *securityRestrictionsSourceIpEnforcementDefaultXml_11_0_2 `xml:"default,omitempty"`
	Misc           []generic.Xml                                             `xml:",any"`
	MiscAttributes []xml.Attr                                                `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementCustomXml_11_0_2 struct {
	SourceIpv4Netmask *int64        `xml:"source-ipv4-netmask,omitempty"`
	SourceIpv6Netmask *int64        `xml:"source-ipv6-netmask,omitempty"`
	Misc              []generic.Xml `xml:",any"`
	MiscAttributes    []xml.Attr    `xml:",any,attr"`
}
type securityRestrictionsSourceIpEnforcementDefaultXml_11_0_2 struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.BlockQuarantinedDevices = util.YesNo(s.BlockQuarantinedDevices, nil)
	o.CertificateProfile = s.CertificateProfile
	if s.ClientAuth != nil {
		var objs []clientAuthXml
		for _, elt := range s.ClientAuth {
			var obj clientAuthXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ClientAuth = &clientAuthContainerXml{Entries: objs}
	}
	if s.HipNotification != nil {
		var objs []hipNotificationXml
		for _, elt := range s.HipNotification {
			var obj hipNotificationXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.HipNotification = &hipNotificationContainerXml{Entries: objs}
	}
	if s.LocalAddress != nil {
		var obj localAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.LogFail = util.YesNo(s.LogFail, nil)
	o.LogSetting = s.LogSetting
	o.LogSuccess = util.YesNo(s.LogSuccess, nil)
	o.RemoteUserTunnel = s.RemoteUserTunnel
	if s.RemoteUserTunnelConfigs != nil {
		var objs []remoteUserTunnelConfigsXml
		for _, elt := range s.RemoteUserTunnelConfigs {
			var obj remoteUserTunnelConfigsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RemoteUserTunnelConfigs = &remoteUserTunnelConfigsContainerXml{Entries: objs}
	}
	if s.Roles != nil {
		var objs []rolesXml
		for _, elt := range s.Roles {
			var obj rolesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Roles = &rolesContainerXml{Entries: objs}
	}
	o.SatelliteTunnel = s.SatelliteTunnel
	if s.SecurityRestrictions != nil {
		var obj securityRestrictionsXml
		obj.MarshalFromObject(*s.SecurityRestrictions)
		o.SecurityRestrictions = &obj
	}
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.TunnelMode = util.YesNo(s.TunnelMode, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var clientAuthVal []ClientAuth
	if o.ClientAuth != nil {
		for _, elt := range o.ClientAuth.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			clientAuthVal = append(clientAuthVal, *obj)
		}
	}
	var hipNotificationVal []HipNotification
	if o.HipNotification != nil {
		for _, elt := range o.HipNotification.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			hipNotificationVal = append(hipNotificationVal, *obj)
		}
	}
	var localAddressVal *LocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var remoteUserTunnelConfigsVal []RemoteUserTunnelConfigs
	if o.RemoteUserTunnelConfigs != nil {
		for _, elt := range o.RemoteUserTunnelConfigs.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			remoteUserTunnelConfigsVal = append(remoteUserTunnelConfigsVal, *obj)
		}
	}
	var rolesVal []Roles
	if o.Roles != nil {
		for _, elt := range o.Roles.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rolesVal = append(rolesVal, *obj)
		}
	}
	var securityRestrictionsVal *SecurityRestrictions
	if o.SecurityRestrictions != nil {
		obj, err := o.SecurityRestrictions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityRestrictionsVal = obj
	}

	result := &Entry{
		Name:                    o.Name,
		BlockQuarantinedDevices: util.AsBool(o.BlockQuarantinedDevices, nil),
		CertificateProfile:      o.CertificateProfile,
		ClientAuth:              clientAuthVal,
		HipNotification:         hipNotificationVal,
		LocalAddress:            localAddressVal,
		LogFail:                 util.AsBool(o.LogFail, nil),
		LogSetting:              o.LogSetting,
		LogSuccess:              util.AsBool(o.LogSuccess, nil),
		RemoteUserTunnel:        o.RemoteUserTunnel,
		RemoteUserTunnelConfigs: remoteUserTunnelConfigsVal,
		Roles:                   rolesVal,
		SatelliteTunnel:         o.SatelliteTunnel,
		SecurityRestrictions:    securityRestrictionsVal,
		SslTlsServiceProfile:    o.SslTlsServiceProfile,
		TunnelMode:              util.AsBool(o.TunnelMode, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *clientAuthXml) MarshalFromObject(s ClientAuth) {
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

func (o clientAuthXml) UnmarshalToObject() (*ClientAuth, error) {

	result := &ClientAuth{
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
func (o *hipNotificationXml) MarshalFromObject(s HipNotification) {
	o.Name = s.Name
	if s.MatchMessage != nil {
		var obj hipNotificationMatchMessageXml
		obj.MarshalFromObject(*s.MatchMessage)
		o.MatchMessage = &obj
	}
	if s.NotMatchMessage != nil {
		var obj hipNotificationNotMatchMessageXml
		obj.MarshalFromObject(*s.NotMatchMessage)
		o.NotMatchMessage = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationXml) UnmarshalToObject() (*HipNotification, error) {
	var matchMessageVal *HipNotificationMatchMessage
	if o.MatchMessage != nil {
		obj, err := o.MatchMessage.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchMessageVal = obj
	}
	var notMatchMessageVal *HipNotificationNotMatchMessage
	if o.NotMatchMessage != nil {
		obj, err := o.NotMatchMessage.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		notMatchMessageVal = obj
	}

	result := &HipNotification{
		Name:            o.Name,
		MatchMessage:    matchMessageVal,
		NotMatchMessage: notMatchMessageVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *hipNotificationMatchMessageXml) MarshalFromObject(s HipNotificationMatchMessage) {
	o.IncludeAppList = util.YesNo(s.IncludeAppList, nil)
	o.ShowNotificationAs = s.ShowNotificationAs
	o.Message = s.Message
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationMatchMessageXml) UnmarshalToObject() (*HipNotificationMatchMessage, error) {

	result := &HipNotificationMatchMessage{
		IncludeAppList:     util.AsBool(o.IncludeAppList, nil),
		ShowNotificationAs: o.ShowNotificationAs,
		Message:            o.Message,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *hipNotificationNotMatchMessageXml) MarshalFromObject(s HipNotificationNotMatchMessage) {
	o.ShowNotificationAs = s.ShowNotificationAs
	o.Message = s.Message
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationNotMatchMessageXml) UnmarshalToObject() (*HipNotificationNotMatchMessage, error) {

	result := &HipNotificationNotMatchMessage{
		ShowNotificationAs: o.ShowNotificationAs,
		Message:            o.Message,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressXml) MarshalFromObject(s LocalAddress) {
	o.Interface = s.Interface
	o.IpAddressFamily = s.IpAddressFamily
	if s.FloatingIp != nil {
		var obj localAddressFloatingIpXml
		obj.MarshalFromObject(*s.FloatingIp)
		o.FloatingIp = &obj
	}
	if s.Ip != nil {
		var obj localAddressIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressXml) UnmarshalToObject() (*LocalAddress, error) {
	var floatingIpVal *LocalAddressFloatingIp
	if o.FloatingIp != nil {
		obj, err := o.FloatingIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floatingIpVal = obj
	}
	var ipVal *LocalAddressIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &LocalAddress{
		Interface:       o.Interface,
		IpAddressFamily: o.IpAddressFamily,
		FloatingIp:      floatingIpVal,
		Ip:              ipVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressFloatingIpXml) MarshalFromObject(s LocalAddressFloatingIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressFloatingIpXml) UnmarshalToObject() (*LocalAddressFloatingIp, error) {

	result := &LocalAddressFloatingIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressIpXml) MarshalFromObject(s LocalAddressIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressIpXml) UnmarshalToObject() (*LocalAddressIp, error) {

	result := &LocalAddressIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsXml) MarshalFromObject(s RemoteUserTunnelConfigs) {
	o.Name = s.Name
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Os != nil {
		o.Os = util.StrToMem(s.Os)
	}
	if s.DnsServer != nil {
		o.DnsServer = util.StrToMem(s.DnsServer)
	}
	if s.DnsSuffix != nil {
		o.DnsSuffix = util.StrToMem(s.DnsSuffix)
	}
	if s.IpPool != nil {
		o.IpPool = util.StrToMem(s.IpPool)
	}
	if s.AuthenticationServerIpPool != nil {
		o.AuthenticationServerIpPool = util.StrToMem(s.AuthenticationServerIpPool)
	}
	if s.AuthenticationOverride != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideXml
		obj.MarshalFromObject(*s.AuthenticationOverride)
		o.AuthenticationOverride = &obj
	}
	if s.SourceAddress != nil {
		var obj remoteUserTunnelConfigsSourceAddressXml
		obj.MarshalFromObject(*s.SourceAddress)
		o.SourceAddress = &obj
	}
	if s.SplitTunneling != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingXml
		obj.MarshalFromObject(*s.SplitTunneling)
		o.SplitTunneling = &obj
	}
	o.NoDirectAccessToLocalNetwork = util.YesNo(s.NoDirectAccessToLocalNetwork, nil)
	o.RetrieveFramedIpAddress = util.YesNo(s.RetrieveFramedIpAddress, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsXml) UnmarshalToObject() (*RemoteUserTunnelConfigs, error) {
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var osVal []string
	if o.Os != nil {
		osVal = util.MemToStr(o.Os)
	}
	var dnsServerVal []string
	if o.DnsServer != nil {
		dnsServerVal = util.MemToStr(o.DnsServer)
	}
	var dnsSuffixVal []string
	if o.DnsSuffix != nil {
		dnsSuffixVal = util.MemToStr(o.DnsSuffix)
	}
	var ipPoolVal []string
	if o.IpPool != nil {
		ipPoolVal = util.MemToStr(o.IpPool)
	}
	var authenticationServerIpPoolVal []string
	if o.AuthenticationServerIpPool != nil {
		authenticationServerIpPoolVal = util.MemToStr(o.AuthenticationServerIpPool)
	}
	var authenticationOverrideVal *RemoteUserTunnelConfigsAuthenticationOverride
	if o.AuthenticationOverride != nil {
		obj, err := o.AuthenticationOverride.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationOverrideVal = obj
	}
	var sourceAddressVal *RemoteUserTunnelConfigsSourceAddress
	if o.SourceAddress != nil {
		obj, err := o.SourceAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceAddressVal = obj
	}
	var splitTunnelingVal *RemoteUserTunnelConfigsSplitTunneling
	if o.SplitTunneling != nil {
		obj, err := o.SplitTunneling.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		splitTunnelingVal = obj
	}

	result := &RemoteUserTunnelConfigs{
		Name:                         o.Name,
		SourceUser:                   sourceUserVal,
		Os:                           osVal,
		DnsServer:                    dnsServerVal,
		DnsSuffix:                    dnsSuffixVal,
		IpPool:                       ipPoolVal,
		AuthenticationServerIpPool:   authenticationServerIpPoolVal,
		AuthenticationOverride:       authenticationOverrideVal,
		SourceAddress:                sourceAddressVal,
		SplitTunneling:               splitTunnelingVal,
		NoDirectAccessToLocalNetwork: util.AsBool(o.NoDirectAccessToLocalNetwork, nil),
		RetrieveFramedIpAddress:      util.AsBool(o.RetrieveFramedIpAddress, nil),
		Misc:                         o.Misc,
		MiscAttributes:               o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideXml) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverride) {
	o.GenerateCookie = util.YesNo(s.GenerateCookie, nil)
	o.CookieEncryptDecryptCert = s.CookieEncryptDecryptCert
	if s.AcceptCookie != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml
		obj.MarshalFromObject(*s.AcceptCookie)
		o.AcceptCookie = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideXml) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverride, error) {
	var acceptCookieVal *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie
	if o.AcceptCookie != nil {
		obj, err := o.AcceptCookie.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		acceptCookieVal = obj
	}

	result := &RemoteUserTunnelConfigsAuthenticationOverride{
		GenerateCookie:           util.AsBool(o.GenerateCookie, nil),
		CookieEncryptDecryptCert: o.CookieEncryptDecryptCert,
		AcceptCookie:             acceptCookieVal,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie) {
	if s.CookieLifetime != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml
		obj.MarshalFromObject(*s.CookieLifetime)
		o.CookieLifetime = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie, error) {
	var cookieLifetimeVal *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime
	if o.CookieLifetime != nil {
		obj, err := o.CookieLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cookieLifetimeVal = obj
	}

	result := &RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie{
		CookieLifetime: cookieLifetimeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime) {
	o.LifetimeInDays = s.LifetimeInDays
	o.LifetimeInHours = s.LifetimeInHours
	o.LifetimeInMinutes = s.LifetimeInMinutes
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime, error) {

	result := &RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime{
		LifetimeInDays:    o.LifetimeInDays,
		LifetimeInHours:   o.LifetimeInHours,
		LifetimeInMinutes: o.LifetimeInMinutes,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSourceAddressXml) MarshalFromObject(s RemoteUserTunnelConfigsSourceAddress) {
	if s.Region != nil {
		o.Region = util.StrToMem(s.Region)
	}
	if s.IpAddress != nil {
		o.IpAddress = util.StrToMem(s.IpAddress)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSourceAddressXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSourceAddress, error) {
	var regionVal []string
	if o.Region != nil {
		regionVal = util.MemToStr(o.Region)
	}
	var ipAddressVal []string
	if o.IpAddress != nil {
		ipAddressVal = util.MemToStr(o.IpAddress)
	}

	result := &RemoteUserTunnelConfigsSourceAddress{
		Region:         regionVal,
		IpAddress:      ipAddressVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingXml) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunneling) {
	if s.AccessRoute != nil {
		o.AccessRoute = util.StrToMem(s.AccessRoute)
	}
	if s.ExcludeAccessRoute != nil {
		o.ExcludeAccessRoute = util.StrToMem(s.ExcludeAccessRoute)
	}
	if s.IncludeApplications != nil {
		o.IncludeApplications = util.StrToMem(s.IncludeApplications)
	}
	if s.ExcludeApplications != nil {
		o.ExcludeApplications = util.StrToMem(s.ExcludeApplications)
	}
	if s.IncludeDomains != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml
		obj.MarshalFromObject(*s.IncludeDomains)
		o.IncludeDomains = &obj
	}
	if s.ExcludeDomains != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml
		obj.MarshalFromObject(*s.ExcludeDomains)
		o.ExcludeDomains = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunneling, error) {
	var accessRouteVal []string
	if o.AccessRoute != nil {
		accessRouteVal = util.MemToStr(o.AccessRoute)
	}
	var excludeAccessRouteVal []string
	if o.ExcludeAccessRoute != nil {
		excludeAccessRouteVal = util.MemToStr(o.ExcludeAccessRoute)
	}
	var includeApplicationsVal []string
	if o.IncludeApplications != nil {
		includeApplicationsVal = util.MemToStr(o.IncludeApplications)
	}
	var excludeApplicationsVal []string
	if o.ExcludeApplications != nil {
		excludeApplicationsVal = util.MemToStr(o.ExcludeApplications)
	}
	var includeDomainsVal *RemoteUserTunnelConfigsSplitTunnelingIncludeDomains
	if o.IncludeDomains != nil {
		obj, err := o.IncludeDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		includeDomainsVal = obj
	}
	var excludeDomainsVal *RemoteUserTunnelConfigsSplitTunnelingExcludeDomains
	if o.ExcludeDomains != nil {
		obj, err := o.ExcludeDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		excludeDomainsVal = obj
	}

	result := &RemoteUserTunnelConfigsSplitTunneling{
		AccessRoute:         accessRouteVal,
		ExcludeAccessRoute:  excludeAccessRouteVal,
		IncludeApplications: includeApplicationsVal,
		ExcludeApplications: excludeApplicationsVal,
		IncludeDomains:      includeDomainsVal,
		ExcludeDomains:      excludeDomainsVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingIncludeDomains) {
	if s.List != nil {
		var objs []remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml
		for _, elt := range s.List {
			var obj remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingIncludeDomains, error) {
	var listVal []RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingIncludeDomains{
		List:           listVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList) {
	o.Name = s.Name
	if s.Ports != nil {
		o.Ports = util.Int64ToMem(s.Ports)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList, error) {
	var portsVal []int64
	if o.Ports != nil {
		var err error
		portsVal, err = util.MemToInt64(o.Ports)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList{
		Name:           o.Name,
		Ports:          portsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingExcludeDomains) {
	if s.List != nil {
		var objs []remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml
		for _, elt := range s.List {
			var obj remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingExcludeDomains, error) {
	var listVal []RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingExcludeDomains{
		List:           listVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList) {
	o.Name = s.Name
	if s.Ports != nil {
		o.Ports = util.Int64ToMem(s.Ports)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList, error) {
	var portsVal []int64
	if o.Ports != nil {
		var err error
		portsVal, err = util.MemToInt64(o.Ports)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList{
		Name:           o.Name,
		Ports:          portsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *rolesXml) MarshalFromObject(s Roles) {
	o.Name = s.Name
	if s.LoginLifetime != nil {
		var obj rolesLoginLifetimeXml
		obj.MarshalFromObject(*s.LoginLifetime)
		o.LoginLifetime = &obj
	}
	o.InactivityLogout = s.InactivityLogout
	o.LifetimeNotifyPrior = s.LifetimeNotifyPrior
	o.LifetimeNotifyMessage = s.LifetimeNotifyMessage
	o.InactivityNotifyPrior = s.InactivityNotifyPrior
	o.InactivityNotifyMessage = s.InactivityNotifyMessage
	o.AdminLogoutNotify = util.YesNo(s.AdminLogoutNotify, nil)
	o.AdminLogoutNotifyMessage = s.AdminLogoutNotifyMessage
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o rolesXml) UnmarshalToObject() (*Roles, error) {
	var loginLifetimeVal *RolesLoginLifetime
	if o.LoginLifetime != nil {
		obj, err := o.LoginLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		loginLifetimeVal = obj
	}

	result := &Roles{
		Name:                     o.Name,
		LoginLifetime:            loginLifetimeVal,
		InactivityLogout:         o.InactivityLogout,
		LifetimeNotifyPrior:      o.LifetimeNotifyPrior,
		LifetimeNotifyMessage:    o.LifetimeNotifyMessage,
		InactivityNotifyPrior:    o.InactivityNotifyPrior,
		InactivityNotifyMessage:  o.InactivityNotifyMessage,
		AdminLogoutNotify:        util.AsBool(o.AdminLogoutNotify, nil),
		AdminLogoutNotifyMessage: o.AdminLogoutNotifyMessage,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *rolesLoginLifetimeXml) MarshalFromObject(s RolesLoginLifetime) {
	o.Minutes = s.Minutes
	o.Hours = s.Hours
	o.Days = s.Days
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o rolesLoginLifetimeXml) UnmarshalToObject() (*RolesLoginLifetime, error) {

	result := &RolesLoginLifetime{
		Minutes:        o.Minutes,
		Hours:          o.Hours,
		Days:           o.Days,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsXml) MarshalFromObject(s SecurityRestrictions) {
	o.DisallowAutomaticRestoration = util.YesNo(s.DisallowAutomaticRestoration, nil)
	if s.SourceIpEnforcement != nil {
		var obj securityRestrictionsSourceIpEnforcementXml
		obj.MarshalFromObject(*s.SourceIpEnforcement)
		o.SourceIpEnforcement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsXml) UnmarshalToObject() (*SecurityRestrictions, error) {
	var sourceIpEnforcementVal *SecurityRestrictionsSourceIpEnforcement
	if o.SourceIpEnforcement != nil {
		obj, err := o.SourceIpEnforcement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceIpEnforcementVal = obj
	}

	result := &SecurityRestrictions{
		DisallowAutomaticRestoration: util.AsBool(o.DisallowAutomaticRestoration, nil),
		SourceIpEnforcement:          sourceIpEnforcementVal,
		Misc:                         o.Misc,
		MiscAttributes:               o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementXml) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcement) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Custom != nil {
		var obj securityRestrictionsSourceIpEnforcementCustomXml
		obj.MarshalFromObject(*s.Custom)
		o.Custom = &obj
	}
	if s.Default != nil {
		var obj securityRestrictionsSourceIpEnforcementDefaultXml
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementXml) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcement, error) {
	var customVal *SecurityRestrictionsSourceIpEnforcementCustom
	if o.Custom != nil {
		obj, err := o.Custom.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customVal = obj
	}
	var defaultVal *SecurityRestrictionsSourceIpEnforcementDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}

	result := &SecurityRestrictionsSourceIpEnforcement{
		Enable:         util.AsBool(o.Enable, nil),
		Custom:         customVal,
		Default:        defaultVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementCustomXml) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcementCustom) {
	o.SourceIpv4Netmask = s.SourceIpv4Netmask
	o.SourceIpv6Netmask = s.SourceIpv6Netmask
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementCustomXml) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcementCustom, error) {

	result := &SecurityRestrictionsSourceIpEnforcementCustom{
		SourceIpv4Netmask: o.SourceIpv4Netmask,
		SourceIpv6Netmask: o.SourceIpv6Netmask,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementDefaultXml) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcementDefault) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementDefaultXml) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcementDefault, error) {

	result := &SecurityRestrictionsSourceIpEnforcementDefault{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.BlockQuarantinedDevices = util.YesNo(s.BlockQuarantinedDevices, nil)
	o.CertificateProfile = s.CertificateProfile
	if s.ClientAuth != nil {
		var objs []clientAuthXml_11_0_2
		for _, elt := range s.ClientAuth {
			var obj clientAuthXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ClientAuth = &clientAuthContainerXml_11_0_2{Entries: objs}
	}
	if s.HipNotification != nil {
		var objs []hipNotificationXml_11_0_2
		for _, elt := range s.HipNotification {
			var obj hipNotificationXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.HipNotification = &hipNotificationContainerXml_11_0_2{Entries: objs}
	}
	if s.LocalAddress != nil {
		var obj localAddressXml_11_0_2
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.LogFail = util.YesNo(s.LogFail, nil)
	o.LogSetting = s.LogSetting
	o.LogSuccess = util.YesNo(s.LogSuccess, nil)
	o.RemoteUserTunnel = s.RemoteUserTunnel
	if s.RemoteUserTunnelConfigs != nil {
		var objs []remoteUserTunnelConfigsXml_11_0_2
		for _, elt := range s.RemoteUserTunnelConfigs {
			var obj remoteUserTunnelConfigsXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.RemoteUserTunnelConfigs = &remoteUserTunnelConfigsContainerXml_11_0_2{Entries: objs}
	}
	if s.Roles != nil {
		var objs []rolesXml_11_0_2
		for _, elt := range s.Roles {
			var obj rolesXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Roles = &rolesContainerXml_11_0_2{Entries: objs}
	}
	o.SatelliteTunnel = s.SatelliteTunnel
	if s.SecurityRestrictions != nil {
		var obj securityRestrictionsXml_11_0_2
		obj.MarshalFromObject(*s.SecurityRestrictions)
		o.SecurityRestrictions = &obj
	}
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.TunnelMode = util.YesNo(s.TunnelMode, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var clientAuthVal []ClientAuth
	if o.ClientAuth != nil {
		for _, elt := range o.ClientAuth.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			clientAuthVal = append(clientAuthVal, *obj)
		}
	}
	var hipNotificationVal []HipNotification
	if o.HipNotification != nil {
		for _, elt := range o.HipNotification.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			hipNotificationVal = append(hipNotificationVal, *obj)
		}
	}
	var localAddressVal *LocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var remoteUserTunnelConfigsVal []RemoteUserTunnelConfigs
	if o.RemoteUserTunnelConfigs != nil {
		for _, elt := range o.RemoteUserTunnelConfigs.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			remoteUserTunnelConfigsVal = append(remoteUserTunnelConfigsVal, *obj)
		}
	}
	var rolesVal []Roles
	if o.Roles != nil {
		for _, elt := range o.Roles.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rolesVal = append(rolesVal, *obj)
		}
	}
	var securityRestrictionsVal *SecurityRestrictions
	if o.SecurityRestrictions != nil {
		obj, err := o.SecurityRestrictions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityRestrictionsVal = obj
	}

	result := &Entry{
		Name:                    o.Name,
		BlockQuarantinedDevices: util.AsBool(o.BlockQuarantinedDevices, nil),
		CertificateProfile:      o.CertificateProfile,
		ClientAuth:              clientAuthVal,
		HipNotification:         hipNotificationVal,
		LocalAddress:            localAddressVal,
		LogFail:                 util.AsBool(o.LogFail, nil),
		LogSetting:              o.LogSetting,
		LogSuccess:              util.AsBool(o.LogSuccess, nil),
		RemoteUserTunnel:        o.RemoteUserTunnel,
		RemoteUserTunnelConfigs: remoteUserTunnelConfigsVal,
		Roles:                   rolesVal,
		SatelliteTunnel:         o.SatelliteTunnel,
		SecurityRestrictions:    securityRestrictionsVal,
		SslTlsServiceProfile:    o.SslTlsServiceProfile,
		TunnelMode:              util.AsBool(o.TunnelMode, nil),
		Misc:                    o.Misc,
		MiscAttributes:          o.MiscAttributes,
	}
	return result, nil
}
func (o *clientAuthXml_11_0_2) MarshalFromObject(s ClientAuth) {
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

func (o clientAuthXml_11_0_2) UnmarshalToObject() (*ClientAuth, error) {

	result := &ClientAuth{
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
func (o *hipNotificationXml_11_0_2) MarshalFromObject(s HipNotification) {
	o.Name = s.Name
	if s.MatchMessage != nil {
		var obj hipNotificationMatchMessageXml_11_0_2
		obj.MarshalFromObject(*s.MatchMessage)
		o.MatchMessage = &obj
	}
	if s.NotMatchMessage != nil {
		var obj hipNotificationNotMatchMessageXml_11_0_2
		obj.MarshalFromObject(*s.NotMatchMessage)
		o.NotMatchMessage = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationXml_11_0_2) UnmarshalToObject() (*HipNotification, error) {
	var matchMessageVal *HipNotificationMatchMessage
	if o.MatchMessage != nil {
		obj, err := o.MatchMessage.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		matchMessageVal = obj
	}
	var notMatchMessageVal *HipNotificationNotMatchMessage
	if o.NotMatchMessage != nil {
		obj, err := o.NotMatchMessage.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		notMatchMessageVal = obj
	}

	result := &HipNotification{
		Name:            o.Name,
		MatchMessage:    matchMessageVal,
		NotMatchMessage: notMatchMessageVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *hipNotificationMatchMessageXml_11_0_2) MarshalFromObject(s HipNotificationMatchMessage) {
	o.IncludeAppList = util.YesNo(s.IncludeAppList, nil)
	o.ShowNotificationAs = s.ShowNotificationAs
	o.Message = s.Message
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationMatchMessageXml_11_0_2) UnmarshalToObject() (*HipNotificationMatchMessage, error) {

	result := &HipNotificationMatchMessage{
		IncludeAppList:     util.AsBool(o.IncludeAppList, nil),
		ShowNotificationAs: o.ShowNotificationAs,
		Message:            o.Message,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *hipNotificationNotMatchMessageXml_11_0_2) MarshalFromObject(s HipNotificationNotMatchMessage) {
	o.ShowNotificationAs = s.ShowNotificationAs
	o.Message = s.Message
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o hipNotificationNotMatchMessageXml_11_0_2) UnmarshalToObject() (*HipNotificationNotMatchMessage, error) {

	result := &HipNotificationNotMatchMessage{
		ShowNotificationAs: o.ShowNotificationAs,
		Message:            o.Message,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressXml_11_0_2) MarshalFromObject(s LocalAddress) {
	o.Interface = s.Interface
	o.IpAddressFamily = s.IpAddressFamily
	if s.FloatingIp != nil {
		var obj localAddressFloatingIpXml_11_0_2
		obj.MarshalFromObject(*s.FloatingIp)
		o.FloatingIp = &obj
	}
	if s.Ip != nil {
		var obj localAddressIpXml_11_0_2
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressXml_11_0_2) UnmarshalToObject() (*LocalAddress, error) {
	var floatingIpVal *LocalAddressFloatingIp
	if o.FloatingIp != nil {
		obj, err := o.FloatingIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floatingIpVal = obj
	}
	var ipVal *LocalAddressIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &LocalAddress{
		Interface:       o.Interface,
		IpAddressFamily: o.IpAddressFamily,
		FloatingIp:      floatingIpVal,
		Ip:              ipVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressFloatingIpXml_11_0_2) MarshalFromObject(s LocalAddressFloatingIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressFloatingIpXml_11_0_2) UnmarshalToObject() (*LocalAddressFloatingIp, error) {

	result := &LocalAddressFloatingIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *localAddressIpXml_11_0_2) MarshalFromObject(s LocalAddressIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o localAddressIpXml_11_0_2) UnmarshalToObject() (*LocalAddressIp, error) {

	result := &LocalAddressIp{
		Ipv4:           o.Ipv4,
		Ipv6:           o.Ipv6,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigs) {
	o.Name = s.Name
	if s.SourceUser != nil {
		o.SourceUser = util.StrToMem(s.SourceUser)
	}
	if s.Os != nil {
		o.Os = util.StrToMem(s.Os)
	}
	if s.DnsServer != nil {
		o.DnsServer = util.StrToMem(s.DnsServer)
	}
	if s.DnsSuffix != nil {
		o.DnsSuffix = util.StrToMem(s.DnsSuffix)
	}
	if s.IpPool != nil {
		o.IpPool = util.StrToMem(s.IpPool)
	}
	if s.AuthenticationServerIpPool != nil {
		o.AuthenticationServerIpPool = util.StrToMem(s.AuthenticationServerIpPool)
	}
	if s.AuthenticationOverride != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideXml_11_0_2
		obj.MarshalFromObject(*s.AuthenticationOverride)
		o.AuthenticationOverride = &obj
	}
	if s.SourceAddress != nil {
		var obj remoteUserTunnelConfigsSourceAddressXml_11_0_2
		obj.MarshalFromObject(*s.SourceAddress)
		o.SourceAddress = &obj
	}
	if s.SplitTunneling != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingXml_11_0_2
		obj.MarshalFromObject(*s.SplitTunneling)
		o.SplitTunneling = &obj
	}
	o.NoDirectAccessToLocalNetwork = util.YesNo(s.NoDirectAccessToLocalNetwork, nil)
	o.RetrieveFramedIpAddress = util.YesNo(s.RetrieveFramedIpAddress, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigs, error) {
	var sourceUserVal []string
	if o.SourceUser != nil {
		sourceUserVal = util.MemToStr(o.SourceUser)
	}
	var osVal []string
	if o.Os != nil {
		osVal = util.MemToStr(o.Os)
	}
	var dnsServerVal []string
	if o.DnsServer != nil {
		dnsServerVal = util.MemToStr(o.DnsServer)
	}
	var dnsSuffixVal []string
	if o.DnsSuffix != nil {
		dnsSuffixVal = util.MemToStr(o.DnsSuffix)
	}
	var ipPoolVal []string
	if o.IpPool != nil {
		ipPoolVal = util.MemToStr(o.IpPool)
	}
	var authenticationServerIpPoolVal []string
	if o.AuthenticationServerIpPool != nil {
		authenticationServerIpPoolVal = util.MemToStr(o.AuthenticationServerIpPool)
	}
	var authenticationOverrideVal *RemoteUserTunnelConfigsAuthenticationOverride
	if o.AuthenticationOverride != nil {
		obj, err := o.AuthenticationOverride.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationOverrideVal = obj
	}
	var sourceAddressVal *RemoteUserTunnelConfigsSourceAddress
	if o.SourceAddress != nil {
		obj, err := o.SourceAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceAddressVal = obj
	}
	var splitTunnelingVal *RemoteUserTunnelConfigsSplitTunneling
	if o.SplitTunneling != nil {
		obj, err := o.SplitTunneling.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		splitTunnelingVal = obj
	}

	result := &RemoteUserTunnelConfigs{
		Name:                         o.Name,
		SourceUser:                   sourceUserVal,
		Os:                           osVal,
		DnsServer:                    dnsServerVal,
		DnsSuffix:                    dnsSuffixVal,
		IpPool:                       ipPoolVal,
		AuthenticationServerIpPool:   authenticationServerIpPoolVal,
		AuthenticationOverride:       authenticationOverrideVal,
		SourceAddress:                sourceAddressVal,
		SplitTunneling:               splitTunnelingVal,
		NoDirectAccessToLocalNetwork: util.AsBool(o.NoDirectAccessToLocalNetwork, nil),
		RetrieveFramedIpAddress:      util.AsBool(o.RetrieveFramedIpAddress, nil),
		Misc:                         o.Misc,
		MiscAttributes:               o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverride) {
	o.GenerateCookie = util.YesNo(s.GenerateCookie, nil)
	o.CookieEncryptDecryptCert = s.CookieEncryptDecryptCert
	if s.AcceptCookie != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml_11_0_2
		obj.MarshalFromObject(*s.AcceptCookie)
		o.AcceptCookie = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverride, error) {
	var acceptCookieVal *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie
	if o.AcceptCookie != nil {
		obj, err := o.AcceptCookie.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		acceptCookieVal = obj
	}

	result := &RemoteUserTunnelConfigsAuthenticationOverride{
		GenerateCookie:           util.AsBool(o.GenerateCookie, nil),
		CookieEncryptDecryptCert: o.CookieEncryptDecryptCert,
		AcceptCookie:             acceptCookieVal,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie) {
	if s.CookieLifetime != nil {
		var obj remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml_11_0_2
		obj.MarshalFromObject(*s.CookieLifetime)
		o.CookieLifetime = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie, error) {
	var cookieLifetimeVal *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime
	if o.CookieLifetime != nil {
		obj, err := o.CookieLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cookieLifetimeVal = obj
	}

	result := &RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie{
		CookieLifetime: cookieLifetimeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime) {
	o.LifetimeInDays = s.LifetimeInDays
	o.LifetimeInHours = s.LifetimeInHours
	o.LifetimeInMinutes = s.LifetimeInMinutes
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetimeXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime, error) {

	result := &RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime{
		LifetimeInDays:    o.LifetimeInDays,
		LifetimeInHours:   o.LifetimeInHours,
		LifetimeInMinutes: o.LifetimeInMinutes,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSourceAddressXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSourceAddress) {
	if s.Region != nil {
		o.Region = util.StrToMem(s.Region)
	}
	if s.IpAddress != nil {
		o.IpAddress = util.StrToMem(s.IpAddress)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSourceAddressXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSourceAddress, error) {
	var regionVal []string
	if o.Region != nil {
		regionVal = util.MemToStr(o.Region)
	}
	var ipAddressVal []string
	if o.IpAddress != nil {
		ipAddressVal = util.MemToStr(o.IpAddress)
	}

	result := &RemoteUserTunnelConfigsSourceAddress{
		Region:         regionVal,
		IpAddress:      ipAddressVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunneling) {
	if s.AccessRoute != nil {
		o.AccessRoute = util.StrToMem(s.AccessRoute)
	}
	if s.ExcludeAccessRoute != nil {
		o.ExcludeAccessRoute = util.StrToMem(s.ExcludeAccessRoute)
	}
	if s.IncludeApplications != nil {
		o.IncludeApplications = util.StrToMem(s.IncludeApplications)
	}
	if s.ExcludeApplications != nil {
		o.ExcludeApplications = util.StrToMem(s.ExcludeApplications)
	}
	if s.IncludeDomains != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml_11_0_2
		obj.MarshalFromObject(*s.IncludeDomains)
		o.IncludeDomains = &obj
	}
	if s.ExcludeDomains != nil {
		var obj remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml_11_0_2
		obj.MarshalFromObject(*s.ExcludeDomains)
		o.ExcludeDomains = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunneling, error) {
	var accessRouteVal []string
	if o.AccessRoute != nil {
		accessRouteVal = util.MemToStr(o.AccessRoute)
	}
	var excludeAccessRouteVal []string
	if o.ExcludeAccessRoute != nil {
		excludeAccessRouteVal = util.MemToStr(o.ExcludeAccessRoute)
	}
	var includeApplicationsVal []string
	if o.IncludeApplications != nil {
		includeApplicationsVal = util.MemToStr(o.IncludeApplications)
	}
	var excludeApplicationsVal []string
	if o.ExcludeApplications != nil {
		excludeApplicationsVal = util.MemToStr(o.ExcludeApplications)
	}
	var includeDomainsVal *RemoteUserTunnelConfigsSplitTunnelingIncludeDomains
	if o.IncludeDomains != nil {
		obj, err := o.IncludeDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		includeDomainsVal = obj
	}
	var excludeDomainsVal *RemoteUserTunnelConfigsSplitTunnelingExcludeDomains
	if o.ExcludeDomains != nil {
		obj, err := o.ExcludeDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		excludeDomainsVal = obj
	}

	result := &RemoteUserTunnelConfigsSplitTunneling{
		AccessRoute:         accessRouteVal,
		ExcludeAccessRoute:  excludeAccessRouteVal,
		IncludeApplications: includeApplicationsVal,
		ExcludeApplications: excludeApplicationsVal,
		IncludeDomains:      includeDomainsVal,
		ExcludeDomains:      excludeDomainsVal,
		Misc:                o.Misc,
		MiscAttributes:      o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingIncludeDomains) {
	if s.List != nil {
		var objs []remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2
		for _, elt := range s.List {
			var obj remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingIncludeDomainsXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingIncludeDomains, error) {
	var listVal []RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingIncludeDomains{
		List:           listVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList) {
	o.Name = s.Name
	if s.Ports != nil {
		o.Ports = util.Int64ToMem(s.Ports)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingIncludeDomainsListXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList, error) {
	var portsVal []int64
	if o.Ports != nil {
		var err error
		portsVal, err = util.MemToInt64(o.Ports)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList{
		Name:           o.Name,
		Ports:          portsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingExcludeDomains) {
	if s.List != nil {
		var objs []remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2
		for _, elt := range s.List {
			var obj remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.List = &remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingExcludeDomainsXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingExcludeDomains, error) {
	var listVal []RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList
	if o.List != nil {
		for _, elt := range o.List.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listVal = append(listVal, *obj)
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingExcludeDomains{
		List:           listVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2) MarshalFromObject(s RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList) {
	o.Name = s.Name
	if s.Ports != nil {
		o.Ports = util.Int64ToMem(s.Ports)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o remoteUserTunnelConfigsSplitTunnelingExcludeDomainsListXml_11_0_2) UnmarshalToObject() (*RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList, error) {
	var portsVal []int64
	if o.Ports != nil {
		var err error
		portsVal, err = util.MemToInt64(o.Ports)
		if err != nil && !errors.Is(err, util.ErrEmptyMemberList) {
			return nil, err
		}
	}

	result := &RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList{
		Name:           o.Name,
		Ports:          portsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *rolesXml_11_0_2) MarshalFromObject(s Roles) {
	o.Name = s.Name
	if s.LoginLifetime != nil {
		var obj rolesLoginLifetimeXml_11_0_2
		obj.MarshalFromObject(*s.LoginLifetime)
		o.LoginLifetime = &obj
	}
	o.InactivityLogout = s.InactivityLogout
	o.LifetimeNotifyPrior = s.LifetimeNotifyPrior
	o.LifetimeNotifyMessage = s.LifetimeNotifyMessage
	o.InactivityNotifyPrior = s.InactivityNotifyPrior
	o.InactivityNotifyMessage = s.InactivityNotifyMessage
	o.AdminLogoutNotify = util.YesNo(s.AdminLogoutNotify, nil)
	o.AdminLogoutNotifyMessage = s.AdminLogoutNotifyMessage
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o rolesXml_11_0_2) UnmarshalToObject() (*Roles, error) {
	var loginLifetimeVal *RolesLoginLifetime
	if o.LoginLifetime != nil {
		obj, err := o.LoginLifetime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		loginLifetimeVal = obj
	}

	result := &Roles{
		Name:                     o.Name,
		LoginLifetime:            loginLifetimeVal,
		InactivityLogout:         o.InactivityLogout,
		LifetimeNotifyPrior:      o.LifetimeNotifyPrior,
		LifetimeNotifyMessage:    o.LifetimeNotifyMessage,
		InactivityNotifyPrior:    o.InactivityNotifyPrior,
		InactivityNotifyMessage:  o.InactivityNotifyMessage,
		AdminLogoutNotify:        util.AsBool(o.AdminLogoutNotify, nil),
		AdminLogoutNotifyMessage: o.AdminLogoutNotifyMessage,
		Misc:                     o.Misc,
		MiscAttributes:           o.MiscAttributes,
	}
	return result, nil
}
func (o *rolesLoginLifetimeXml_11_0_2) MarshalFromObject(s RolesLoginLifetime) {
	o.Minutes = s.Minutes
	o.Hours = s.Hours
	o.Days = s.Days
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o rolesLoginLifetimeXml_11_0_2) UnmarshalToObject() (*RolesLoginLifetime, error) {

	result := &RolesLoginLifetime{
		Minutes:        o.Minutes,
		Hours:          o.Hours,
		Days:           o.Days,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsXml_11_0_2) MarshalFromObject(s SecurityRestrictions) {
	o.DisallowAutomaticRestoration = util.YesNo(s.DisallowAutomaticRestoration, nil)
	if s.SourceIpEnforcement != nil {
		var obj securityRestrictionsSourceIpEnforcementXml_11_0_2
		obj.MarshalFromObject(*s.SourceIpEnforcement)
		o.SourceIpEnforcement = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsXml_11_0_2) UnmarshalToObject() (*SecurityRestrictions, error) {
	var sourceIpEnforcementVal *SecurityRestrictionsSourceIpEnforcement
	if o.SourceIpEnforcement != nil {
		obj, err := o.SourceIpEnforcement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sourceIpEnforcementVal = obj
	}

	result := &SecurityRestrictions{
		DisallowAutomaticRestoration: util.AsBool(o.DisallowAutomaticRestoration, nil),
		SourceIpEnforcement:          sourceIpEnforcementVal,
		Misc:                         o.Misc,
		MiscAttributes:               o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementXml_11_0_2) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcement) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Custom != nil {
		var obj securityRestrictionsSourceIpEnforcementCustomXml_11_0_2
		obj.MarshalFromObject(*s.Custom)
		o.Custom = &obj
	}
	if s.Default != nil {
		var obj securityRestrictionsSourceIpEnforcementDefaultXml_11_0_2
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementXml_11_0_2) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcement, error) {
	var customVal *SecurityRestrictionsSourceIpEnforcementCustom
	if o.Custom != nil {
		obj, err := o.Custom.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customVal = obj
	}
	var defaultVal *SecurityRestrictionsSourceIpEnforcementDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}

	result := &SecurityRestrictionsSourceIpEnforcement{
		Enable:         util.AsBool(o.Enable, nil),
		Custom:         customVal,
		Default:        defaultVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementCustomXml_11_0_2) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcementCustom) {
	o.SourceIpv4Netmask = s.SourceIpv4Netmask
	o.SourceIpv6Netmask = s.SourceIpv6Netmask
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementCustomXml_11_0_2) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcementCustom, error) {

	result := &SecurityRestrictionsSourceIpEnforcementCustom{
		SourceIpv4Netmask: o.SourceIpv4Netmask,
		SourceIpv6Netmask: o.SourceIpv6Netmask,
		Misc:              o.Misc,
		MiscAttributes:    o.MiscAttributes,
	}
	return result, nil
}
func (o *securityRestrictionsSourceIpEnforcementDefaultXml_11_0_2) MarshalFromObject(s SecurityRestrictionsSourceIpEnforcementDefault) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o securityRestrictionsSourceIpEnforcementDefaultXml_11_0_2) UnmarshalToObject() (*SecurityRestrictionsSourceIpEnforcementDefault, error) {

	result := &SecurityRestrictionsSourceIpEnforcementDefault{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "block_quarantined_devices" || v == "BlockQuarantinedDevices" {
		return e.BlockQuarantinedDevices, nil
	}
	if v == "certificate_profile" || v == "CertificateProfile" {
		return e.CertificateProfile, nil
	}
	if v == "client_auth" || v == "ClientAuth" {
		return e.ClientAuth, nil
	}
	if v == "client_auth|LENGTH" || v == "ClientAuth|LENGTH" {
		return int64(len(e.ClientAuth)), nil
	}
	if v == "hip_notification" || v == "HipNotification" {
		return e.HipNotification, nil
	}
	if v == "hip_notification|LENGTH" || v == "HipNotification|LENGTH" {
		return int64(len(e.HipNotification)), nil
	}
	if v == "local_address" || v == "LocalAddress" {
		return e.LocalAddress, nil
	}
	if v == "log_fail" || v == "LogFail" {
		return e.LogFail, nil
	}
	if v == "log_setting" || v == "LogSetting" {
		return e.LogSetting, nil
	}
	if v == "log_success" || v == "LogSuccess" {
		return e.LogSuccess, nil
	}
	if v == "remote_user_tunnel" || v == "RemoteUserTunnel" {
		return e.RemoteUserTunnel, nil
	}
	if v == "remote_user_tunnel_configs" || v == "RemoteUserTunnelConfigs" {
		return e.RemoteUserTunnelConfigs, nil
	}
	if v == "remote_user_tunnel_configs|LENGTH" || v == "RemoteUserTunnelConfigs|LENGTH" {
		return int64(len(e.RemoteUserTunnelConfigs)), nil
	}
	if v == "roles" || v == "Roles" {
		return e.Roles, nil
	}
	if v == "roles|LENGTH" || v == "Roles|LENGTH" {
		return int64(len(e.Roles)), nil
	}
	if v == "satellite_tunnel" || v == "SatelliteTunnel" {
		return e.SatelliteTunnel, nil
	}
	if v == "security_restrictions" || v == "SecurityRestrictions" {
		return e.SecurityRestrictions, nil
	}
	if v == "ssl_tls_service_profile" || v == "SslTlsServiceProfile" {
		return e.SslTlsServiceProfile, nil
	}
	if v == "tunnel_mode" || v == "TunnelMode" {
		return e.TunnelMode, nil
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
	if !util.BoolsMatch(o.BlockQuarantinedDevices, other.BlockQuarantinedDevices) {
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
	if len(o.HipNotification) != len(other.HipNotification) {
		return false
	}
	for idx := range o.HipNotification {
		if !o.HipNotification[idx].matches(&other.HipNotification[idx]) {
			return false
		}
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
	if !util.StringsMatch(o.RemoteUserTunnel, other.RemoteUserTunnel) {
		return false
	}
	if len(o.RemoteUserTunnelConfigs) != len(other.RemoteUserTunnelConfigs) {
		return false
	}
	for idx := range o.RemoteUserTunnelConfigs {
		if !o.RemoteUserTunnelConfigs[idx].matches(&other.RemoteUserTunnelConfigs[idx]) {
			return false
		}
	}
	if len(o.Roles) != len(other.Roles) {
		return false
	}
	for idx := range o.Roles {
		if !o.Roles[idx].matches(&other.Roles[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.SatelliteTunnel, other.SatelliteTunnel) {
		return false
	}
	if !o.SecurityRestrictions.matches(other.SecurityRestrictions) {
		return false
	}
	if !util.StringsMatch(o.SslTlsServiceProfile, other.SslTlsServiceProfile) {
		return false
	}
	if !util.BoolsMatch(o.TunnelMode, other.TunnelMode) {
		return false
	}

	return true
}

func (o *ClientAuth) matches(other *ClientAuth) bool {
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

func (o *HipNotification) matches(other *HipNotification) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.MatchMessage.matches(other.MatchMessage) {
		return false
	}
	if !o.NotMatchMessage.matches(other.NotMatchMessage) {
		return false
	}

	return true
}

func (o *HipNotificationMatchMessage) matches(other *HipNotificationMatchMessage) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.IncludeAppList, other.IncludeAppList) {
		return false
	}
	if !util.StringsMatch(o.ShowNotificationAs, other.ShowNotificationAs) {
		return false
	}
	if !util.StringsMatch(o.Message, other.Message) {
		return false
	}

	return true
}

func (o *HipNotificationNotMatchMessage) matches(other *HipNotificationNotMatchMessage) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ShowNotificationAs, other.ShowNotificationAs) {
		return false
	}
	if !util.StringsMatch(o.Message, other.Message) {
		return false
	}

	return true
}

func (o *LocalAddress) matches(other *LocalAddress) bool {
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

func (o *LocalAddressFloatingIp) matches(other *LocalAddressFloatingIp) bool {
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

func (o *LocalAddressIp) matches(other *LocalAddressIp) bool {
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

func (o *RemoteUserTunnelConfigs) matches(other *RemoteUserTunnelConfigs) bool {
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
	if !util.OrderedListsMatch[string](o.Os, other.Os) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DnsServer, other.DnsServer) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DnsSuffix, other.DnsSuffix) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IpPool, other.IpPool) {
		return false
	}
	if !util.OrderedListsMatch[string](o.AuthenticationServerIpPool, other.AuthenticationServerIpPool) {
		return false
	}
	if !o.AuthenticationOverride.matches(other.AuthenticationOverride) {
		return false
	}
	if !o.SourceAddress.matches(other.SourceAddress) {
		return false
	}
	if !o.SplitTunneling.matches(other.SplitTunneling) {
		return false
	}
	if !util.BoolsMatch(o.NoDirectAccessToLocalNetwork, other.NoDirectAccessToLocalNetwork) {
		return false
	}
	if !util.BoolsMatch(o.RetrieveFramedIpAddress, other.RetrieveFramedIpAddress) {
		return false
	}

	return true
}

func (o *RemoteUserTunnelConfigsAuthenticationOverride) matches(other *RemoteUserTunnelConfigsAuthenticationOverride) bool {
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

func (o *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie) matches(other *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookie) bool {
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

func (o *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime) matches(other *RemoteUserTunnelConfigsAuthenticationOverrideAcceptCookieCookieLifetime) bool {
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

func (o *RemoteUserTunnelConfigsSourceAddress) matches(other *RemoteUserTunnelConfigsSourceAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Region, other.Region) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IpAddress, other.IpAddress) {
		return false
	}

	return true
}

func (o *RemoteUserTunnelConfigsSplitTunneling) matches(other *RemoteUserTunnelConfigsSplitTunneling) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.AccessRoute, other.AccessRoute) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExcludeAccessRoute, other.ExcludeAccessRoute) {
		return false
	}
	if !util.OrderedListsMatch[string](o.IncludeApplications, other.IncludeApplications) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExcludeApplications, other.ExcludeApplications) {
		return false
	}
	if !o.IncludeDomains.matches(other.IncludeDomains) {
		return false
	}
	if !o.ExcludeDomains.matches(other.ExcludeDomains) {
		return false
	}

	return true
}

func (o *RemoteUserTunnelConfigsSplitTunnelingIncludeDomains) matches(other *RemoteUserTunnelConfigsSplitTunnelingIncludeDomains) bool {
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

	return true
}

func (o *RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList) matches(other *RemoteUserTunnelConfigsSplitTunnelingIncludeDomainsList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[int64](o.Ports, other.Ports) {
		return false
	}

	return true
}

func (o *RemoteUserTunnelConfigsSplitTunnelingExcludeDomains) matches(other *RemoteUserTunnelConfigsSplitTunnelingExcludeDomains) bool {
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

	return true
}

func (o *RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList) matches(other *RemoteUserTunnelConfigsSplitTunnelingExcludeDomainsList) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[int64](o.Ports, other.Ports) {
		return false
	}

	return true
}

func (o *Roles) matches(other *Roles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.LoginLifetime.matches(other.LoginLifetime) {
		return false
	}
	if !util.Ints64Match(o.InactivityLogout, other.InactivityLogout) {
		return false
	}
	if !util.Ints64Match(o.LifetimeNotifyPrior, other.LifetimeNotifyPrior) {
		return false
	}
	if !util.StringsMatch(o.LifetimeNotifyMessage, other.LifetimeNotifyMessage) {
		return false
	}
	if !util.Ints64Match(o.InactivityNotifyPrior, other.InactivityNotifyPrior) {
		return false
	}
	if !util.StringsMatch(o.InactivityNotifyMessage, other.InactivityNotifyMessage) {
		return false
	}
	if !util.BoolsMatch(o.AdminLogoutNotify, other.AdminLogoutNotify) {
		return false
	}
	if !util.StringsMatch(o.AdminLogoutNotifyMessage, other.AdminLogoutNotifyMessage) {
		return false
	}

	return true
}

func (o *RolesLoginLifetime) matches(other *RolesLoginLifetime) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Minutes, other.Minutes) {
		return false
	}
	if !util.Ints64Match(o.Hours, other.Hours) {
		return false
	}
	if !util.Ints64Match(o.Days, other.Days) {
		return false
	}

	return true
}

func (o *SecurityRestrictions) matches(other *SecurityRestrictions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.DisallowAutomaticRestoration, other.DisallowAutomaticRestoration) {
		return false
	}
	if !o.SourceIpEnforcement.matches(other.SourceIpEnforcement) {
		return false
	}

	return true
}

func (o *SecurityRestrictionsSourceIpEnforcement) matches(other *SecurityRestrictionsSourceIpEnforcement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Custom.matches(other.Custom) {
		return false
	}
	if !o.Default.matches(other.Default) {
		return false
	}

	return true
}

func (o *SecurityRestrictionsSourceIpEnforcementCustom) matches(other *SecurityRestrictionsSourceIpEnforcementCustom) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.SourceIpv4Netmask, other.SourceIpv4Netmask) {
		return false
	}
	if !util.Ints64Match(o.SourceIpv6Netmask, other.SourceIpv6Netmask) {
		return false
	}

	return true
}

func (o *SecurityRestrictionsSourceIpEnforcementDefault) matches(other *SecurityRestrictionsSourceIpEnforcementDefault) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
