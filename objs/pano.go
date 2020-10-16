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
	vulnrule "github.com/PaloAltoNetworks/pango/objs/profile/security/vulnerability/rule"
	"github.com/PaloAltoNetworks/pango/objs/srvc"
	"github.com/PaloAltoNetworks/pango/objs/srvcgrp"
	"github.com/PaloAltoNetworks/pango/objs/tags"
)

// PanoObjs is the client.Objects namespace.
type PanoObjs struct {
	Address                             *addr.Panorama
	AddressGroup                        *addrgrp.PanoAddrGrp
	AntivirusProfile                    *virus.Panorama
	Application                         *app.PanoApp
	AppGroup                            *appgrp.PanoGroup
	AppSignature                        *signature.PanoSignature
	AppSigAndCond                       *andcond.PanoAndCond
	AppSigOrCond                        *orcond.PanoOrCond
	DynamicUserGroup                    *dug.Panorama
	Edl                                 *edl.PanoEdl
	LogForwardingProfile                *logfwd.Panorama
	LogForwardingProfileMatchList       *matchlist.PanoMatchList
	LogForwardingProfileMatchListAction *action.PanoAction
	Services                            *srvc.PanoSrvc
	ServiceGroup                        *srvcgrp.PanoSrvcGrp
	Tags                                *tags.PanoTags
	VulnerabilityProfile                *vulnerability.Panorama
	VulnerabilityProfileRule            *vulnrule.Panorama
}

// Initialize is invoked on client.Initialize().
func (c *PanoObjs) Initialize(i util.XapiClient) {
	c.Address = addr.PanoramaNamespace(i)

	c.AddressGroup = &addrgrp.PanoAddrGrp{}
	c.AddressGroup.Initialize(i)

	c.AntivirusProfile = virus.PanoramaNamespace(i)

	c.Application = &app.PanoApp{}
	c.Application.Initialize(i)

	c.AppGroup = &appgrp.PanoGroup{}
	c.AppGroup.Initialize(i)

	c.AppSignature = &signature.PanoSignature{}
	c.AppSignature.Initialize(i)

	c.AppSigAndCond = &andcond.PanoAndCond{}
	c.AppSigAndCond.Initialize(i)

	c.AppSigOrCond = &orcond.PanoOrCond{}
	c.AppSigOrCond.Initialize(i)

	c.DynamicUserGroup = dug.PanoramaNamespace(i)

	c.Edl = &edl.PanoEdl{}
	c.Edl.Initialize(i)

	c.LogForwardingProfile = logfwd.PanoramaNamespace(i)

	c.LogForwardingProfileMatchList = &matchlist.PanoMatchList{}
	c.LogForwardingProfileMatchList.Initialize(i)

	c.LogForwardingProfileMatchListAction = &action.PanoAction{}
	c.LogForwardingProfileMatchListAction.Initialize(i)

	c.Services = &srvc.PanoSrvc{}
	c.Services.Initialize(i)

	c.ServiceGroup = &srvcgrp.PanoSrvcGrp{}
	c.ServiceGroup.Initialize(i)

	c.Tags = &tags.PanoTags{}
	c.Tags.Initialize(i)

	c.VulnerabilityProfile = vulnerability.PanoramaNamespace(i)
	c.VulnerabilityProfileRule = vulnrule.PanoramaNamespace(i)
}
