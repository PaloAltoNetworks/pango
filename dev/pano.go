package dev

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/dev/profile/certificate"
	"github.com/PaloAltoNetworks/pango/dev/profile/email"
	"github.com/PaloAltoNetworks/pango/dev/profile/http"
	"github.com/PaloAltoNetworks/pango/dev/profile/snmp"
	"github.com/PaloAltoNetworks/pango/dev/profile/syslog"
	syslogsrv "github.com/PaloAltoNetworks/pango/dev/profile/syslog/server"
	"github.com/PaloAltoNetworks/pango/dev/ssldecrypt"
)

// PanoDev is the client.Device namespace.
type PanoDev struct {
	CertificateProfile  *certificate.Panorama
	EmailServerProfile  *email.Panorama
	HttpServerProfile   *http.Panorama
	SslDecrypt          *ssldecrypt.Panorama
	SnmpServerProfile   *snmp.Panorama
	SyslogServer        *syslogsrv.PanoServer
	SyslogServerProfile *syslog.PanoSyslog
}

// Initialize is invoked on client.Initialize().
func (c *PanoDev) Initialize(i util.XapiClient) {
	c.CertificateProfile = certificate.PanoramaNamespace(i)
	c.EmailServerProfile = email.PanoramaNamespace(i)
	c.HttpServerProfile = http.PanoramaNamespace(i)
	c.SslDecrypt = ssldecrypt.PanoramaNamespace(i)
	c.SnmpServerProfile = snmp.PanoramaNamespace(i)

	c.SyslogServer = &syslogsrv.PanoServer{}
	c.SyslogServer.Initialize(i)

	c.SyslogServerProfile = &syslog.PanoSyslog{}
	c.SyslogServerProfile.Initialize(i)
}
