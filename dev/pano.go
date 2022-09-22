package dev

import (
	"github.com/PaloAltoNetworks/pango/util"

	cert "github.com/PaloAltoNetworks/pango/dev/certificate"
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
	"github.com/PaloAltoNetworks/pango/dev/vminfosource"
)

// Panorama is the client.Device namespace.
type Panorama struct {
	AuthenticationProfile *authentication.Panorama
	Certificate           *cert.Panorama
	CertificateProfile    *certificate.Panorama
	EmailServerProfile    *email.Panorama
	HttpServerProfile     *http.Panorama
	KerberosProfile       *kerberos.Panorama
	LdapProfile           *ldap.Panorama
	LocalUserDbGroup      *group.Panorama
	LocalUserDbUser       *user.Panorama
	RadiusProfile         *radius.Panorama
	SamlProfile           *saml.Panorama
	SslTlsServiceProfile  *ssltls.Panorama
	SslDecrypt            *ssldecrypt.Panorama
	SnmpServerProfile     *snmp.Panorama
	SyslogServerProfile   *syslog.Panorama
	TacacsPlusProfile     *tacplus.Panorama
	VmInfoSource          *vminfosource.Panorama
}

// Initialize is invoked on client.Initialize().
func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		AuthenticationProfile: authentication.PanoramaNamespace(x),
		Certificate:           cert.PanoramaNamespace(x),
		CertificateProfile:    certificate.PanoramaNamespace(x),
		EmailServerProfile:    email.PanoramaNamespace(x),
		HttpServerProfile:     http.PanoramaNamespace(x),
		KerberosProfile:       kerberos.PanoramaNamespace(x),
		LdapProfile:           ldap.PanoramaNamespace(x),
		LocalUserDbGroup:      group.PanoramaNamespace(x),
		LocalUserDbUser:       user.PanoramaNamespace(x),
		RadiusProfile:         radius.PanoramaNamespace(x),
		SamlProfile:           saml.PanoramaNamespace(x),
		SslTlsServiceProfile:  ssltls.PanoramaNamespace(x),
		SslDecrypt:            ssldecrypt.PanoramaNamespace(x),
		SnmpServerProfile:     snmp.PanoramaNamespace(x),
		SyslogServerProfile:   syslog.PanoramaNamespace(x),
		TacacsPlusProfile:     tacplus.PanoramaNamespace(x),
		VmInfoSource:          vminfosource.PanoramaNamespace(x),
	}
}
