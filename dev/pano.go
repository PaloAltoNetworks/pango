package dev

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/dev/profile/certificate"
	"github.com/PaloAltoNetworks/pango/dev/profile/email"
	"github.com/PaloAltoNetworks/pango/dev/profile/http"
	"github.com/PaloAltoNetworks/pango/dev/profile/snmp"
	"github.com/PaloAltoNetworks/pango/dev/profile/syslog"
	"github.com/PaloAltoNetworks/pango/dev/ssldecrypt"
)

// Panorama is the client.Device namespace.
type Panorama struct {
	CertificateProfile  *certificate.Panorama
	EmailServerProfile  *email.Panorama
	HttpServerProfile   *http.Panorama
	SslDecrypt          *ssldecrypt.Panorama
	SnmpServerProfile   *snmp.Panorama
	SyslogServerProfile *syslog.Panorama
}

// Initialize is invoked on client.Initialize().
func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		CertificateProfile:  certificate.PanoramaNamespace(x),
		EmailServerProfile:  email.PanoramaNamespace(x),
		HttpServerProfile:   http.PanoramaNamespace(x),
		SslDecrypt:          ssldecrypt.PanoramaNamespace(x),
		SnmpServerProfile:   snmp.PanoramaNamespace(x),
		SyslogServerProfile: syslog.PanoramaNamespace(x),
	}
}
