package objs


import (
    "github.com/PaloAltoNetworks/pango/util"

    "github.com/PaloAltoNetworks/pango/objs/addr"
    "github.com/PaloAltoNetworks/pango/objs/addrgrp"
    "github.com/PaloAltoNetworks/pango/objs/app"
    "github.com/PaloAltoNetworks/pango/objs/edl"
    "github.com/PaloAltoNetworks/pango/objs/profile/logfwd"
    "github.com/PaloAltoNetworks/pango/objs/profile/logfwd/matchlist"
    "github.com/PaloAltoNetworks/pango/objs/profile/logfwd/matchlist/action"
    "github.com/PaloAltoNetworks/pango/objs/srvc"
    "github.com/PaloAltoNetworks/pango/objs/srvcgrp"
    "github.com/PaloAltoNetworks/pango/objs/tags"
)


// PanoObjs is the client.Objects namespace.
type PanoObjs struct {
    Address *addr.PanoAddr
    AddressGroup *addrgrp.PanoAddrGrp
    Application *app.PanoApp
    Edl *edl.PanoEdl
    LogForwardingProfile *logfwd.PanoLogFwd
    LogForwardingProfileMatchList *matchlist.PanoMatchList
    LogForwardingProfileMatchListAction *action.PanoAction
    Services *srvc.PanoSrvc
    ServiceGroup *srvcgrp.PanoSrvcGrp
    Tags *tags.PanoTags
}

// Initialize is invoked on client.Initialize().
func (c *PanoObjs) Initialize(i util.XapiClient) {
    c.Address = &addr.PanoAddr{}
    c.Address.Initialize(i)

    c.AddressGroup = &addrgrp.PanoAddrGrp{}
    c.AddressGroup.Initialize(i)

    c.Application = &app.PanoApp{}
    c.Application.Initialize(i)

    c.Edl = &edl.PanoEdl{}
    c.Edl.Initialize(i)

    c.LogForwardingProfile = &logfwd.PanoLogFwd{}
    c.LogForwardingProfile.Initialize(i)

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
}
