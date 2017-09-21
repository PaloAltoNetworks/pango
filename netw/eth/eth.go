// Package eth is the client.Network.EthernetInterface namespace.
//
// Normalized object:  Entry
package eth

import (
    "fmt"
    "encoding/xml"

    "github.com/PaloAltoNetworks/xapi/util"
    "github.com/PaloAltoNetworks/xapi/version"
)


// Entry is a normalized, version independent representation of an ethernet
// interface.
type Entry struct {
    Name string
    Mode string
    StaticIps []string
    EnableDhcp bool
    CreateDhcpDefaultRoute bool
    DhcpDefaultRouteMetric int
    Ipv6Enabled bool
    ManagementProfile string
    Mtu int
    AdjustTcpMss bool
    NetflowProfile string
    LldpEnabled bool
    LldpProfile string
    LinkSpeed string
    LinkDuplex string
    LinkState string
    AggregateGroup string
    Comment string
    Ipv4MssAdjust int
    Ipv6MssAdjust int
}

// Eth is the client.Network.EthernetInterface namespace.
type Eth struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *Eth) Initialize(con util.XapiClient) {
    c.con = con
}

// ShowList performs SHOW to retrieve a list of ethernet interfaces.
func (c *Eth) ShowList() ([]string, error) {
    c.con.LogQuery("(show) list of ethernet interfaces")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// GetList performs GET to retrieve a list of ethernet interfaces.
func (c *Eth) GetList() ([]string, error) {
    c.con.LogQuery("(get) list of ethernet interfaces")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given ethernet interface.
func (c *Eth) Get(name string) (Entry, error) {
    c.con.LogQuery("(get) ethernet interface %q", name)
    return c.details(c.con.Get, name)
}

// Show performs SHOW to retrieve information for the given ethernet interface.
func (c *Eth) Show(name string) (Entry, error) {
    c.con.LogQuery("(show) ethernet interface %q", name)
    return c.details(c.con.Show, name)
}

// Set creates / updates one or more ethernet interfaces.
//
// Specify a non-empty vsys to import the interface(s) into the given vsys
// after creating, allowing the vsys to use them.
func (c *Eth) Set(vsys string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given interface configs.
    d := util.BulkElement{XMLName: xml.Name{Local: "ethernet"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) ethernet interfaces: %v", names)

    // Set xpath.
    path := c.xpath(names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the interfaces.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    if err != nil {
        return err
    }

    // Perform vsys import next.
    return c.con.ImportInterfaces(vsys, names)
}

// Delete removes the given interface(s) from the firewall.
//
// Specify a non-empty vsys to have this function remove the interface(s) from
// the vsys prior to deleting them.
//
// Interfaces can be a string or an eth.Entry object.
func (c *Eth) Delete(vsys string, e ...interface{}) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    names := make([]string, len(e))
    for i := range e {
        switch v := e[i].(type) {
        case string:
            names[i] = v
        case Entry:
            names[i] = v.Name
        default:
            return fmt.Errorf("Unknown type sent to delete: %s", v)
        }
    }
    c.con.LogAction("(delete) ethernet interface(s): %v", names)

    // Unimport interfaces from the given vsys.
    err = c.con.UnimportInterfaces(vsys, names)
    if err != nil {
        return err
    }

    // Remove interfaces next.
    path := c.xpath(names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for the Eth struct **/

func (c *Eth) versioning() (normalizer, func(Entry) (interface{})) {
    v := c.con.Versioning()

    if v.Gte(version.Number{7, 1, 0, ""}) {
        return &container_v2{}, specify_v2
    } else {
        return &container_v1{}, specify_v1
    }
}

func (c *Eth) details(fn func(interface{}, interface{}, interface{}) ([]byte, error), name string) (Entry, error) {
    path := c.xpath([]string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *Eth) xpath(vals []string) []string {
    return []string{
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "network",
        "interface",
        "ethernet",
        util.AsEntryXpath(vals),
    }
}

/** Structs / functions for this namespace. **/

type normalizer interface {
    Normalize() Entry
}

type container_v1 struct {
    Answer entry_v1 `xml:"result>entry"`
}

func (o *container_v1) Normalize() Entry {
    ans := Entry{
        Name: o.Answer.Name,
        LinkSpeed: o.Answer.LinkSpeed,
        LinkDuplex: o.Answer.LinkDuplex,
        LinkState: o.Answer.LinkState,
        Comment: o.Answer.Comment,
    }
    switch {
        case o.Answer.ModeL3 != nil:
            ans.Mode = "layer3"
            ans.Ipv6Enabled = util.AsBool(o.Answer.ModeL3.Ipv6Enabled)
            ans.ManagementProfile = o.Answer.ModeL3.ManagementProfile
            ans.Mtu = o.Answer.ModeL3.Mtu
            ans.NetflowProfile = o.Answer.ModeL3.NetflowProfile
            ans.AdjustTcpMss = util.AsBool(o.Answer.ModeL3.AdjustTcpMss)
            if o.Answer.ModeL3.StaticIps != nil {
                ans.StaticIps = make([]string, len(o.Answer.ModeL3.StaticIps))
                for i := range o.Answer.ModeL3.StaticIps {
                    ans.StaticIps[i] = o.Answer.ModeL3.StaticIps[i].IpAddress
                }
            }
            if o.Answer.ModeL3.Dhcp != nil {
                ans.EnableDhcp = util.AsBool(o.Answer.ModeL3.Dhcp.Enable)
                ans.CreateDhcpDefaultRoute = util.AsBool(o.Answer.ModeL3.Dhcp.CreateDefaultRoute)
                ans.DhcpDefaultRouteMetric = o.Answer.ModeL3.Dhcp.Metric
            }
        case o.Answer.ModeL2 != nil:
            ans.Mode = "layer2"
            ans.LldpEnabled = util.AsBool(o.Answer.ModeL2.LldpEnabled)
            ans.LldpProfile = o.Answer.ModeL2.LldpProfile
            ans.NetflowProfile = o.Answer.ModeL2.NetflowProfile
        case o.Answer.ModeVwire != nil:
            ans.Mode = "virtual-wire"
            ans.LldpEnabled = util.AsBool(o.Answer.ModeVwire.LldpEnabled)
            ans.LldpProfile = o.Answer.ModeVwire.LldpProfile
            ans.NetflowProfile = o.Answer.ModeVwire.NetflowProfile
        case o.Answer.TapMode != nil:
            ans.Mode = "tap"
        case o.Answer.HaMode != nil:
            ans.Mode = "ha"
        case o.Answer.DecryptMirrorMode != nil:
            ans.Mode = "decrypt-mirror"
        case o.Answer.AggregateGroupMode != nil:
            ans.Mode = "aggregate-group"
    }

    return ans
}

type entry_v1 struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    ModeL2 *otherMode `xml:"layer2"`
    ModeL3 *l3Mode_v1 `xml:"layer3"`
    ModeVwire *otherMode `xml:"virtual-wire"`
    TapMode *emptyMode `xml:"tap"`
    HaMode *emptyMode `xml:"ha"`
    DecryptMirrorMode *emptyMode `xml:"decrypt-mirror"`
    AggregateGroupMode *emptyMode `xml:"aggregate-group"`
    LinkSpeed string `xml:"link-speed,omitempty"`
    LinkDuplex string `xml:"link-duplex,omitempty"`
    LinkState string `xml:"link-state,omitempty"`
    Comment string `xml:"comment"`
}

type emptyMode struct {}

type otherMode struct {
    LldpEnabled string `xml:"lldp>enable"`
    LldpProfile string `xml:"lldp>profile"`
    NetflowProfile string `xml:"netflow-profile"`
}

type l3Mode_v1 struct {
    Ipv6Enabled string `xml:"ipv6>enabled"`
    ManagementProfile string `xml:"interface-management-profile"`
    Mtu int `xml:"mtu,omitempty"`
    NetflowProfile string `xml:"netflow-profile"`
    AdjustTcpMss string `xml:"adjust-tcp-mss"`
    StaticIps []staticIpv4 `xml:"ip"`
    Dhcp *dhcpSettings `xml:"dhcp-client"`
}

type staticIpv4 struct {
    XMLName xml.Name `xml:"entry`
    IpAddress string `xml:"name,attr"`
}

type dhcpSettings struct {
    Enable string `xml:"enable"`
    CreateDefaultRoute string `xml:"create-default-route"`
    Metric int `xml:"default-route-metric,omitempty"`
}

type container_v2 struct {
    Answer entry_v2 `xml:"result>entry"`
}

func (o *container_v2) Normalize() Entry {
    ans := Entry{
        Name: o.Answer.Name,
        LinkSpeed: o.Answer.LinkSpeed,
        LinkDuplex: o.Answer.LinkDuplex,
        LinkState: o.Answer.LinkState,
        Comment: o.Answer.Comment,
    }
    switch {
        case o.Answer.ModeL3 != nil:
            ans.Mode = "layer3"
            ans.Ipv6Enabled = util.AsBool(o.Answer.ModeL3.Ipv6Enabled)
            ans.ManagementProfile = o.Answer.ModeL3.ManagementProfile
            ans.Mtu = o.Answer.ModeL3.Mtu
            ans.NetflowProfile = o.Answer.ModeL3.NetflowProfile
            ans.AdjustTcpMss = util.AsBool(o.Answer.ModeL3.AdjustTcpMss)
            ans.Ipv4MssAdjust = o.Answer.ModeL3.Ipv4MssAdjust
            ans.Ipv6MssAdjust = o.Answer.ModeL3.Ipv6MssAdjust
            if o.Answer.ModeL3.StaticIps != nil {
                ans.StaticIps = make([]string, len(o.Answer.ModeL3.StaticIps))
                for i := range o.Answer.ModeL3.StaticIps {
                    ans.StaticIps[i] = o.Answer.ModeL3.StaticIps[i].IpAddress
                }
            }
            if o.Answer.ModeL3.Dhcp != nil {
                ans.EnableDhcp = util.AsBool(o.Answer.ModeL3.Dhcp.Enable)
                ans.CreateDhcpDefaultRoute = util.AsBool(o.Answer.ModeL3.Dhcp.CreateDefaultRoute)
                ans.DhcpDefaultRouteMetric = o.Answer.ModeL3.Dhcp.Metric
            }
        case o.Answer.ModeL2 != nil:
            ans.Mode = "layer2"
            ans.LldpEnabled = util.AsBool(o.Answer.ModeL2.LldpEnabled)
            ans.LldpProfile = o.Answer.ModeL2.LldpProfile
            ans.NetflowProfile = o.Answer.ModeL2.NetflowProfile
        case o.Answer.ModeVwire != nil:
            ans.Mode = "virtual-wire"
            ans.LldpEnabled = util.AsBool(o.Answer.ModeVwire.LldpEnabled)
            ans.LldpProfile = o.Answer.ModeVwire.LldpProfile
            ans.NetflowProfile = o.Answer.ModeVwire.NetflowProfile
        case o.Answer.TapMode != nil:
            ans.Mode = "tap"
        case o.Answer.HaMode != nil:
            ans.Mode = "ha"
        case o.Answer.DecryptMirrorMode != nil:
            ans.Mode = "decrypt-mirror"
        case o.Answer.AggregateGroupMode != nil:
            ans.Mode = "aggregate-group"
    }

    return ans
}

type entry_v2 struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    ModeL3 *l3Mode_v2 `xml:"layer3"`
    ModeL2 *otherMode `xml:"layer2"`
    ModeVwire *otherMode `xml:"virtual-wire"`
    TapMode *emptyMode `xml:"tap"`
    HaMode *emptyMode `xml:"ha"`
    DecryptMirrorMode *emptyMode `xml:"decrypt-mirror"`
    AggregateGroupMode *emptyMode `xml:"aggregate-group"`
    LinkSpeed string `xml:"link-speed,omitempty"`
    LinkDuplex string `xml:"link-duplex,omitempty"`
    LinkState string `xml:"link-state,omitempty"`
    Comment string `xml:"comment"`
}

type l3Mode_v2 struct {
    Ipv6Enabled string `xml:"ipv6>enabled"`
    ManagementProfile string `xml:"interface-management-profile"`
    Mtu int `xml:"mtu,omitempty"`
    NetflowProfile string `xml:"netflow-profile"`
    AdjustTcpMss string `xml:"adjust-tcp-mss>enable"`
    Ipv4MssAdjust int `xml:"adjust-tcp-mss>ipv4-mss-adjustment,omitempty"`
    Ipv6MssAdjust int `xml:"adjust-tcp-mss>ipv6-mss-adjustment,omitempty"`
    StaticIps []staticIpv4 `xml:"ip"`
    Dhcp *dhcpSettings `xml:"dhcp-client"`
}


func specify_v1(e Entry) interface{} {
    ans := entry_v1{
        Name: e.Name,
        LinkSpeed: e.LinkSpeed,
        LinkDuplex: e.LinkDuplex,
        LinkState: e.LinkState,
        Comment: e.Comment,
    }

    switch e.Mode {
    case "layer3":
        i := &l3Mode_v1{
            Ipv6Enabled: util.YesNo(e.Ipv6Enabled),
            ManagementProfile: e.ManagementProfile,
            Mtu: e.Mtu,
            NetflowProfile: e.NetflowProfile,
            AdjustTcpMss: util.YesNo(e.AdjustTcpMss),
        }
        if len(e.StaticIps) > 0 {
            nfo := make([]staticIpv4, len(e.StaticIps))
            for j := range e.StaticIps {
                nfo = append(nfo, staticIpv4{IpAddress: e.StaticIps[j]})
            }
            i.StaticIps = nfo
        }
        if e.EnableDhcp || e.CreateDhcpDefaultRoute || e.DhcpDefaultRouteMetric != 0 {
            i.Dhcp = &dhcpSettings{
                Enable: util.YesNo(e.EnableDhcp),
                CreateDefaultRoute: util.YesNo(e.CreateDhcpDefaultRoute),
                Metric: e.DhcpDefaultRouteMetric,
            }
        }
        ans.ModeL3 = i
    case "layer2":
        i := &otherMode{
            LldpEnabled: util.YesNo(e.LldpEnabled),
            LldpProfile: e.LldpProfile,
            NetflowProfile: e.NetflowProfile,
        }
        ans.ModeL2 = i
    case "virtual-wire":
        i := &otherMode{
            LldpEnabled: util.YesNo(e.LldpEnabled),
            LldpProfile: e.LldpProfile,
            NetflowProfile: e.NetflowProfile,
        }
        ans.ModeVwire = i
    case "tap":
        ans.TapMode = &emptyMode{}
    case "ha":
        ans.HaMode = &emptyMode{}
    case "decrypt-mirror":
        ans.DecryptMirrorMode = &emptyMode{}
    case "aggregate-group":
        ans.AggregateGroupMode = &emptyMode{}
    }

    return ans
}

func specify_v2(e Entry) interface{} {
    ans := entry_v2{
        Name: e.Name,
        LinkSpeed: e.LinkSpeed,
        LinkDuplex: e.LinkDuplex,
        LinkState: e.LinkState,
        Comment: e.Comment,
    }

    switch e.Mode {
    case "layer3":
        i := &l3Mode_v2{
            Ipv6Enabled: util.YesNo(e.Ipv6Enabled),
            ManagementProfile: e.ManagementProfile,
            Mtu: e.Mtu,
            NetflowProfile: e.NetflowProfile,
            AdjustTcpMss: util.YesNo(e.AdjustTcpMss),
            Ipv4MssAdjust: e.Ipv4MssAdjust,
            Ipv6MssAdjust: e.Ipv6MssAdjust,
        }
        if len(e.StaticIps) > 0 {
            nfo := make([]staticIpv4, len(e.StaticIps))
            for j := range e.StaticIps {
                nfo = append(nfo, staticIpv4{IpAddress: e.StaticIps[j]})
            }
            i.StaticIps = nfo
        } else if e.EnableDhcp || e.CreateDhcpDefaultRoute || e.DhcpDefaultRouteMetric != 0 {
            i.Dhcp = &dhcpSettings{
                Enable: util.YesNo(e.EnableDhcp),
                CreateDefaultRoute: util.YesNo(e.CreateDhcpDefaultRoute),
                Metric: e.DhcpDefaultRouteMetric,
            }
        }
        ans.ModeL3 = i
    case "layer2":
        i := &otherMode{
            LldpEnabled: util.YesNo(e.LldpEnabled),
            LldpProfile: e.LldpProfile,
            NetflowProfile: e.NetflowProfile,
        }
        ans.ModeL2 = i
    case "virtual-wire":
        i := &otherMode{
            LldpEnabled: util.YesNo(e.LldpEnabled),
            LldpProfile: e.LldpProfile,
            NetflowProfile: e.NetflowProfile,
        }
        ans.ModeVwire = i
    case "tap":
        ans.TapMode = &emptyMode{}
    case "ha":
        ans.HaMode = &emptyMode{}
    case "decrypt-mirror":
        ans.DecryptMirrorMode = &emptyMode{}
    case "aggregate-group":
        ans.AggregateGroupMode = &emptyMode{}
    }

    return ans
}

