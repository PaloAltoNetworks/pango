package dev

import (
	"github.com/PaloAltoNetworks/pango/util"

	cert "github.com/PaloAltoNetworks/pango/dev/certificate"
	"github.com/PaloAltoNetworks/pango/dev/general"
	"github.com/PaloAltoNetworks/pango/dev/ha"
	halink "github.com/PaloAltoNetworks/pango/dev/ha/monitor/link"
	hapath "github.com/PaloAltoNetworks/pango/dev/ha/monitor/path"
	"github.com/PaloAltoNetworks/pango/dev/localuserdb/group"
	"github.com/PaloAltoNetworks/pango/dev/localuserdb/user"
	"github.com/PaloAltoNetworks/pango/dev/profile/authentication"
	"github.com/PaloAltoNetworks/pango/dev/profile/certificate"
	"github.com/PaloAltoNetworks/pango/dev/profile/email"
	"github.com/PaloAltoNetworks/pango/dev/profile/http"
	"github.com/PaloAltoNetworks/pango/dev/profile/kerberos"
	"github.com/PaloAltoNetworks/pango/dev/profile/ldap"
	"github.com/PaloAltoNetworks/pango/dev/profile/radius"
	"github.com/PaloAltoNetworks/pango/dev/profile/saml"
	"github.com/PaloAltoNetworks/pango/dev/profile/snmp"
	"github.com/PaloAltoNetworks/pango/dev/profile/ssltls"
	"github.com/PaloAltoNetworks/pango/dev/profile/syslog"
	"github.com/PaloAltoNetworks/pango/dev/profile/tacplus"
	"github.com/PaloAltoNetworks/pango/dev/ssldecrypt"
	"github.com/PaloAltoNetworks/pango/dev/telemetry"
	"github.com/PaloAltoNetworks/pango/dev/vminfosource"
)

// Firewall is the client.Device namespace.
type Firewall struct {
	AuthenticationProfile *authentication.Firewall
	Certificate           *cert.Firewall
	CertificateProfile    *certificate.Firewall
	EmailServerProfile    *email.Firewall
	GeneralSettings       *general.Firewall
	HaConfig              *ha.Firewall
	HaLinkMonitorGroup    *halink.Firewall
	HaPathMonitorGroup    *hapath.Firewall
	HttpServerProfile     *http.Firewall
	KerberosProfile       *kerberos.Firewall
	LdapProfile           *ldap.Firewall
	LocalUserDbGroup      *group.Firewall
	LocalUserDbUser       *user.Firewall
	RadiusProfile         *radius.Firewall
	SamlProfile           *saml.Firewall
	SslDecrypt            *ssldecrypt.Firewall
	SslTlsServiceProfile  *ssltls.Firewall
	SnmpServerProfile     *snmp.Firewall
	SyslogServerProfile   *syslog.Firewall
	TacacsPlusProfile     *tacplus.Firewall
	Telemetry             *telemetry.Firewall
	VmInfoSource          *vminfosource.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		AuthenticationProfile: authentication.FirewallNamespace(x),
		Certificate:           cert.FirewallNamespace(x),
		CertificateProfile:    certificate.FirewallNamespace(x),
		EmailServerProfile:    email.FirewallNamespace(x),
		GeneralSettings:       general.FirewallNamespace(x),
		HaConfig:              ha.FirewallNamespace(x),
		HaLinkMonitorGroup:    halink.FirewallNamespace(x),
		HaPathMonitorGroup:    hapath.FirewallNamespace(x),
		HttpServerProfile:     http.FirewallNamespace(x),
		KerberosProfile:       kerberos.FirewallNamespace(x),
		LdapProfile:           ldap.FirewallNamespace(x),
		LocalUserDbGroup:      group.FirewallNamespace(x),
		LocalUserDbUser:       user.FirewallNamespace(x),
		RadiusProfile:         radius.FirewallNamespace(x),
		SamlProfile:           saml.FirewallNamespace(x),
		SslTlsServiceProfile:  ssltls.FirewallNamespace(x),
		SslDecrypt:            ssldecrypt.FirewallNamespace(x),
		SnmpServerProfile:     snmp.FirewallNamespace(x),
		SyslogServerProfile:   syslog.FirewallNamespace(x),
		TacacsPlusProfile:     tacplus.FirewallNamespace(x),
		Telemetry:             telemetry.FirewallNamespace(x),
		VmInfoSource:          vminfosource.FirewallNamespace(x),
	}
}
