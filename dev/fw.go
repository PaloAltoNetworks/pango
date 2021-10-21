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

// Firewall is the client.Device namespace.
type Firewall struct {
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

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		CertificateProfile:  certificate.FirewallNamespace(x),
		EmailServerProfile:  email.FirewallNamespace(x),
		GeneralSettings:     general.FirewallNamespace(x),
		HaConfig:            ha.FirewallNamespace(x),
		HaLinkMonitorGroup:  halink.FirewallNamespace(x),
		HaPathMonitorGroup:  hapath.FirewallNamespace(x),
		HttpServerProfile:   http.FirewallNamespace(x),
		SslDecrypt:          ssldecrypt.FirewallNamespace(x),
		SnmpServerProfile:   snmp.FirewallNamespace(x),
		SyslogServerProfile: syslog.FirewallNamespace(x),
		Telemetry:           telemetry.FirewallNamespace(x),
	}
}
