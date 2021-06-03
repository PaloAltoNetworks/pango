package predefined

import (
	"github.com/PaloAltoNetworks/pango/objs/app"
	dlpft "github.com/PaloAltoNetworks/pango/predefined/dlp/filetype"
	tdbft "github.com/PaloAltoNetworks/pango/predefined/tdb/filetype"
	"github.com/PaloAltoNetworks/pango/predefined/threat"
	"github.com/PaloAltoNetworks/pango/util"
)

type Firewall struct {
	Application *app.Predefined
	DlpFileType *dlpft.Firewall
	TdbFileType *tdbft.Firewall
	Threat      *threat.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Application: app.PredefinedNamespace(x),
		DlpFileType: dlpft.FirewallNamespace(x),
		TdbFileType: tdbft.FirewallNamespace(x),
		Threat:      threat.FirewallNamespace(x),
	}
}
