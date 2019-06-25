package dev

import (
    "github.com/PaloAltoNetworks/pango/util"

    "github.com/PaloAltoNetworks/pango/dev/profile/snmp"
    "github.com/PaloAltoNetworks/pango/dev/profile/snmp/v2c"
    "github.com/PaloAltoNetworks/pango/dev/profile/snmp/v3"
)


// PanoDev is the client.Device namespace.
type PanoDev struct {
    SnmpServerProfile *snmp.PanoSnmp
    SnmpV2cServer *v2c.PanoV2c
    SnmpV3Server *v3.PanoV3
}

// Initialize is invoked on client.Initialize().
func (c *PanoDev) Initialize(i util.XapiClient) {
    c.SnmpServerProfile = &snmp.PanoSnmp{}
    c.SnmpServerProfile.Initialize(i)

    c.SnmpV2cServer = &v2c.PanoV2c{}
    c.SnmpV2cServer.Initialize(i)

    c.SnmpV3Server = &v3.PanoV3{}
    c.SnmpV3Server.Initialize(i)
}
