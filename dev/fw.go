package dev

import (
    "github.com/PaloAltoNetworks/pango/util"

    "github.com/PaloAltoNetworks/pango/dev/general"
    "github.com/PaloAltoNetworks/pango/dev/profile/snmp"
    "github.com/PaloAltoNetworks/pango/dev/profile/snmp/v2c"
    "github.com/PaloAltoNetworks/pango/dev/profile/snmp/v3"
    "github.com/PaloAltoNetworks/pango/dev/telemetry"
)


// FwDev is the client.Device namespace.
type FwDev struct {
    GeneralSettings *general.FwGeneral
    SnmpServerProfile *snmp.FwSnmp
    SnmpV2cServer *v2c.FwV2c
    SnmpV3Server *v3.FwV3
    Telemetry *telemetry.FwTelemetry
}

// Initialize is invoked on client.Initialize().
func (c *FwDev) Initialize(i util.XapiClient) {
    c.GeneralSettings = &general.FwGeneral{}
    c.GeneralSettings.Initialize(i)

    c.SnmpServerProfile = &snmp.FwSnmp{}
    c.SnmpServerProfile.Initialize(i)

    c.SnmpV2cServer = &v2c.FwV2c{}
    c.SnmpV2cServer.Initialize(i)

    c.SnmpV3Server = &v3.FwV3{}
    c.SnmpV3Server.Initialize(i)

    c.Telemetry = &telemetry.FwTelemetry{}
    c.Telemetry.Initialize(i)
}
