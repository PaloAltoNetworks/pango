package vlan

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
)

// Entry is a normalized, version independent representation of
// a VLAN interface.
type Entry struct {
	Name                   string
	Comment                string
	NetflowProfile         string
	StaticIps              []string // ordered
	EnableDhcp             bool
	CreateDhcpDefaultRoute bool
	DhcpDefaultRouteMetric int
	ManagementProfile      string
	Mtu                    int
	AdjustTcpMss           bool
	Ipv4MssAdjust          int
	Ipv6MssAdjust          int

	raw map[string]string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Comment = s.Comment
	o.NetflowProfile = s.NetflowProfile
	o.StaticIps = s.StaticIps
	o.EnableDhcp = s.EnableDhcp
	o.CreateDhcpDefaultRoute = s.CreateDhcpDefaultRoute
	o.DhcpDefaultRouteMetric = s.DhcpDefaultRouteMetric
	o.ManagementProfile = s.ManagementProfile
	o.Mtu = s.Mtu
	o.AdjustTcpMss = s.AdjustTcpMss
	o.Ipv4MssAdjust = s.Ipv4MssAdjust
	o.Ipv6MssAdjust = s.Ipv6MssAdjust
}

/** Structs / functions for this namespace. **/

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:              o.Name,
		Comment:           o.Comment,
		NetflowProfile:    o.NetflowProfile,
		StaticIps:         util.EntToStr(o.StaticIps),
		Mtu:               int(o.Mtu),
		ManagementProfile: o.ManagementProfile,
		AdjustTcpMss:      util.AsBool(o.AdjustTcpMss),
	}
	if o.Dhcp != nil {
		ans.EnableDhcp = util.AsBool(o.Dhcp.Enable)
		ans.CreateDhcpDefaultRoute = util.AsBool(o.Dhcp.CreateDefaultRoute)
		ans.DhcpDefaultRouteMetric = int(o.Dhcp.Metric)
	}

	ans.raw = make(map[string]string)
	if o.Ipv6 != nil {
		ans.raw["ipv6"] = util.CleanRawXml(o.Ipv6.Text)
	}
	if o.Arp != nil {
		ans.raw["arp"] = util.CleanRawXml(o.Arp.Text)
	}
	if o.NdpProxy != nil {
		ans.raw["ndp"] = util.CleanRawXml(o.NdpProxy.Text)
	}
	if len(ans.raw) == 0 {
		ans.raw = nil
	}

	return ans
}

type entry_v1 struct {
	XMLName           xml.Name        `xml:"entry"`
	Name              string          `xml:"name,attr"`
	Comment           string          `xml:"comment,omitempty"`
	NetflowProfile    string          `xml:"netflow-profile,omitempty"`
	StaticIps         *util.EntryType `xml:"ip"`
	Dhcp              *dhcpSettings   `xml:"dhcp-client"`
	Mtu               int             `xml:"mtu,omitempty"`
	ManagementProfile string          `xml:"interface-management-profile,omitempty"`
	AdjustTcpMss      string          `xml:"adjust-tcp-mss"`

	Ipv6     *util.RawXml `xml:"ipv6"`
	Arp      *util.RawXml `xml:"arp"`
	NdpProxy *util.RawXml `xml:"ndp-proxy"`
}

type dhcpSettings struct {
	Enable             string `xml:"enable"`
	CreateDefaultRoute string `xml:"create-default-route"`
	Metric             int    `xml:"default-route-metric,omitempty"`
}

type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o *container_v2) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:              o.Name,
		Comment:           o.Comment,
		NetflowProfile:    o.NetflowProfile,
		StaticIps:         util.EntToStr(o.StaticIps),
		Mtu:               int(o.Mtu),
		ManagementProfile: o.ManagementProfile,
		AdjustTcpMss:      util.AsBool(o.AdjustTcpMss),
		Ipv4MssAdjust:     int(o.Ipv4MssAdjust),
		Ipv6MssAdjust:     int(o.Ipv6MssAdjust),
	}
	if o.Dhcp != nil {
		ans.EnableDhcp = util.AsBool(o.Dhcp.Enable)
		ans.CreateDhcpDefaultRoute = util.AsBool(o.Dhcp.CreateDefaultRoute)
		ans.DhcpDefaultRouteMetric = int(o.Dhcp.Metric)
	}

	ans.raw = make(map[string]string)
	if o.Ipv6 != nil {
		ans.raw["ipv6"] = util.CleanRawXml(o.Ipv6.Text)
	}
	if o.Arp != nil {
		ans.raw["arp"] = util.CleanRawXml(o.Arp.Text)
	}
	if o.NdpProxy != nil {
		ans.raw["ndp"] = util.CleanRawXml(o.NdpProxy.Text)
	}
	if len(ans.raw) == 0 {
		ans.raw = nil
	}

	return ans
}

type entry_v2 struct {
	XMLName           xml.Name        `xml:"entry"`
	Name              string          `xml:"name,attr"`
	Comment           string          `xml:"comment,omitempty"`
	NetflowProfile    string          `xml:"netflow-profile,omitempty"`
	StaticIps         *util.EntryType `xml:"ip"`
	Dhcp              *dhcpSettings   `xml:"dhcp-client"`
	Mtu               int             `xml:"mtu,omitempty"`
	ManagementProfile string          `xml:"interface-management-profile,omitempty"`
	AdjustTcpMss      string          `xml:"adjust-tcp-mss>enable"`
	Ipv4MssAdjust     int             `xml:"adjust-tcp-mss>ipv4-mss-adjustment,omitempty"`
	Ipv6MssAdjust     int             `xml:"adjust-tcp-mss>ipv6-mss-adjustment,omitempty"`

	Ipv6     *util.RawXml `xml:"ipv6"`
	Arp      *util.RawXml `xml:"arp"`
	NdpProxy *util.RawXml `xml:"ndp-proxy"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:              e.Name,
		Comment:           e.Comment,
		NetflowProfile:    e.NetflowProfile,
		StaticIps:         util.StrToEnt(e.StaticIps),
		Mtu:               e.Mtu,
		ManagementProfile: e.ManagementProfile,
		AdjustTcpMss:      util.YesNo(e.AdjustTcpMss),
	}

	if e.EnableDhcp || e.CreateDhcpDefaultRoute || e.DhcpDefaultRouteMetric != 0 {
		v := dhcpSettings{
			Enable:             util.YesNo(e.EnableDhcp),
			CreateDefaultRoute: util.YesNo(e.CreateDhcpDefaultRoute),
			Metric:             e.DhcpDefaultRouteMetric,
		}
		ans.Dhcp = &v
	}

	if text, ok := e.raw["ipv6"]; ok {
		ans.Ipv6 = &util.RawXml{text}
	}
	if text, ok := e.raw["arp"]; ok {
		ans.Arp = &util.RawXml{text}
	}
	if text, ok := e.raw["ndp"]; ok {
		ans.NdpProxy = &util.RawXml{text}
	}

	return ans
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:              e.Name,
		Comment:           e.Comment,
		NetflowProfile:    e.NetflowProfile,
		StaticIps:         util.StrToEnt(e.StaticIps),
		Mtu:               e.Mtu,
		ManagementProfile: e.ManagementProfile,
		AdjustTcpMss:      util.YesNo(e.AdjustTcpMss),
		Ipv4MssAdjust:     e.Ipv4MssAdjust,
		Ipv6MssAdjust:     e.Ipv6MssAdjust,
	}

	if e.EnableDhcp || e.CreateDhcpDefaultRoute || e.DhcpDefaultRouteMetric != 0 {
		v := dhcpSettings{
			Enable:             util.YesNo(e.EnableDhcp),
			CreateDefaultRoute: util.YesNo(e.CreateDhcpDefaultRoute),
			Metric:             e.DhcpDefaultRouteMetric,
		}
		ans.Dhcp = &v
	}

	if text, ok := e.raw["ipv6"]; ok {
		ans.Ipv6 = &util.RawXml{text}
	}
	if text, ok := e.raw["arp"]; ok {
		ans.Arp = &util.RawXml{text}
	}
	if text, ok := e.raw["ndp"]; ok {
		ans.NdpProxy = &util.RawXml{text}
	}

	return ans
}
