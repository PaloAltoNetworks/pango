package objs

import (
	"github.com/PaloAltoNetworks/pango/util"

	"github.com/PaloAltoNetworks/pango/objs/addr"
	"github.com/PaloAltoNetworks/pango/objs/addrgrp"
	"github.com/PaloAltoNetworks/pango/objs/app"
	appgrp "github.com/PaloAltoNetworks/pango/objs/app/group"
	"github.com/PaloAltoNetworks/pango/objs/app/signature"
	"github.com/PaloAltoNetworks/pango/objs/app/signature/andcond"
	"github.com/PaloAltoNetworks/pango/objs/app/signature/orcond"
	"github.com/PaloAltoNetworks/pango/objs/dug"
	"github.com/PaloAltoNetworks/pango/objs/edl"
	"github.com/PaloAltoNetworks/pango/objs/profile/logfwd"
	"github.com/PaloAltoNetworks/pango/objs/profile/logfwd/matchlist"
	"github.com/PaloAltoNetworks/pango/objs/profile/logfwd/matchlist/action"
	"github.com/PaloAltoNetworks/pango/objs/profile/security/virus"
	"github.com/PaloAltoNetworks/pango/objs/profile/security/vulnerability"
	vulnexcep "github.com/PaloAltoNetworks/pango/objs/profile/security/vulnerability/exception"
	vulnrule "github.com/PaloAltoNetworks/pango/objs/profile/security/vulnerability/rule"
	"github.com/PaloAltoNetworks/pango/objs/srvc"
	"github.com/PaloAltoNetworks/pango/objs/srvcgrp"
	"github.com/PaloAltoNetworks/pango/objs/tags"
)

// FwObjs is the client.Objects namespace.
type FwObjs struct {
	Address                             *addr.Firewall
	AddressGroup                        *addrgrp.FwAddrGrp
	AntivirusProfile                    *virus.Firewall
	Application                         *app.FwApp
	AppGroup                            *appgrp.FwGroup
	AppSignature                        *signature.FwSignature
	AppSigAndCond                       *andcond.FwAndCond
	AppSigOrCond                        *orcond.FwOrCond
	DynamicUserGroup                    *dug.Firewall
	Edl                                 *edl.FwEdl
	LogForwardingProfile                *logfwd.Firewall
	LogForwardingProfileMatchList       *matchlist.FwMatchList
	LogForwardingProfileMatchListAction *action.FwAction
	Services                            *srvc.FwSrvc
	ServiceGroup                        *srvcgrp.FwSrvcGrp
	Tags                                *tags.FwTags
	VulnerabilityProfile                *vulnerability.Firewall
	VulnerabilityProfileException       *vulnexcep.Firewall
	VulnerabilityProfileRule            *vulnrule.Firewall
}

// Initialize is invoked on client.Initialize().
func (c *FwObjs) Initialize(i util.XapiClient) {
	c.Address = addr.FirewallNamespace(i)

	c.AddressGroup = &addrgrp.FwAddrGrp{}
	c.AddressGroup.Initialize(i)

	c.AntivirusProfile = virus.FirewallNamespace(i)

	c.Application = &app.FwApp{}
	c.Application.Initialize(i)

	c.AppGroup = &appgrp.FwGroup{}
	c.AppGroup.Initialize(i)

	c.AppSignature = &signature.FwSignature{}
	c.AppSignature.Initialize(i)

	c.AppSigAndCond = &andcond.FwAndCond{}
	c.AppSigAndCond.Initialize(i)

	c.AppSigOrCond = &orcond.FwOrCond{}
	c.AppSigOrCond.Initialize(i)

	c.DynamicUserGroup = dug.FirewallNamespace(i)

	c.Edl = &edl.FwEdl{}
	c.Edl.Initialize(i)

	c.LogForwardingProfile = logfwd.FirewallNamespace(i)

	c.LogForwardingProfileMatchList = &matchlist.FwMatchList{}
	c.LogForwardingProfileMatchList.Initialize(i)

	c.LogForwardingProfileMatchListAction = &action.FwAction{}
	c.LogForwardingProfileMatchListAction.Initialize(i)

	c.Services = &srvc.FwSrvc{}
	c.Services.Initialize(i)

	c.ServiceGroup = &srvcgrp.FwSrvcGrp{}
	c.ServiceGroup.Initialize(i)

	c.Tags = &tags.FwTags{}
	c.Tags.Initialize(i)

	c.VulnerabilityProfile = vulnerability.FirewallNamespace(i)
	c.VulnerabilityProfileException = vulnexcep.FirewallNamespace(i)
	c.VulnerabilityProfileRule = vulnrule.FirewallNamespace(i)
}
