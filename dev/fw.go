package dev

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/dev/general"
	"github.com/PaloAltoNetworks/pango/dev/ha"
	halink "github.com/PaloAltoNetworks/pango/dev/ha/monitor/link"
	hapath "github.com/PaloAltoNetworks/pango/dev/ha/monitor/path"
	"github.com/PaloAltoNetworks/pango/dev/profile/certificate"
	"github.com/PaloAltoNetworks/pango/dev/profile/email"
	"github.com/PaloAltoNetworks/pango/dev/profile/http"
	"github.com/PaloAltoNetworks/pango/dev/profile/snmp"
	"github.com/PaloAltoNetworks/pango/dev/profile/syslog"
	"github.com/PaloAltoNetworks/pango/dev/ssldecrypt"
	"github.com/PaloAltoNetworks/pango/dev/telemetry"
)

// FwDev is the client.Device namespace.
type FwDev struct {
	CertificateProfile  *certificate.Firewall
	EmailServerProfile  *email.Firewall
	GeneralSettings     *general.Firewall
	HaConfig            *ha.Firewall
	HaLinkMonitorGroup  *halink.Firewall
	HaPathMonitorGroup  *hapath.Firewall
	HttpServerProfile   *http.Firewall
	SslDecrypt          *ssldecrypt.Firewall
	SnmpServerProfile   *snmp.Firewall
	SyslogServerProfile *syslog.Firewall
	Telemetry           *telemetry.Firewall
}

// Initialize is invoked on client.Initialize().
func (c *FwDev) Initialize(i util.XapiClient) {
	c.CertificateProfile = certificate.FirewallNamespace(i)
	c.EmailServerProfile = email.FirewallNamespace(i)
	c.GeneralSettings = general.FirewallNamespace(i)
	c.HaConfig = ha.FirewallNamespace(i)
	c.HaLinkMonitorGroup = halink.FirewallNamespace(i)
	c.HaPathMonitorGroup = hapath.FirewallNamespace(i)
	c.HttpServerProfile = http.FirewallNamespace(i)
	c.SslDecrypt = ssldecrypt.FirewallNamespace(i)
	c.SnmpServerProfile = snmp.FirewallNamespace(i)
	c.SyslogServerProfile = syslog.FirewallNamespace(i)
	c.Telemetry = telemetry.FirewallNamespace(i)
}
