package staticroute

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"network", "virtual-router", "$parent", "routing-table", "ipv6", "static-route", "$name"}
)

type Entry struct {
	Name        string
	AdminDist   *int64
	Bfd         *Bfd
	Destination *string
	Interface   *string
	Metric      *int64
	Nexthop     *Nexthop
	PathMonitor *PathMonitor
	RouteTable  *RouteTable
	Misc        []generic.Xml
}
type Bfd struct {
	Profile *string
	Misc    []generic.Xml
}
type Nexthop struct {
	Discard     *NexthopDiscard
	Ipv6Address *string
	NextVr      *string
	Receive     *NexthopReceive
	Misc        []generic.Xml
}
type NexthopDiscard struct {
	Misc []generic.Xml
}
type NexthopReceive struct {
	Misc []generic.Xml
}
type PathMonitor struct {
	Enable              *bool
	FailureCondition    *string
	HoldTime            *int64
	MonitorDestinations []PathMonitorMonitorDestinations
	Misc                []generic.Xml
}
type PathMonitorMonitorDestinations struct {
	Name        string
	Enable      *bool
	Source      *string
	Destination *string
	Interval    *int64
	Count       *int64
	Misc        []generic.Xml
}
type RouteTable struct {
	NoInstall *RouteTableNoInstall
	Unicast   *RouteTableUnicast
	Misc      []generic.Xml
}
type RouteTableNoInstall struct {
	Misc []generic.Xml
}
type RouteTableUnicast struct {
	Misc []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName     xml.Name        `xml:"entry"`
	Name        string          `xml:"name,attr"`
	AdminDist   *int64          `xml:"admin-dist,omitempty"`
	Bfd         *bfdXml         `xml:"bfd,omitempty"`
	Destination *string         `xml:"destination,omitempty"`
	Interface   *string         `xml:"interface,omitempty"`
	Metric      *int64          `xml:"metric,omitempty"`
	Nexthop     *nexthopXml     `xml:"nexthop,omitempty"`
	PathMonitor *pathMonitorXml `xml:"path-monitor,omitempty"`
	RouteTable  *routeTableXml  `xml:"route-table,omitempty"`
	Misc        []generic.Xml   `xml:",any"`
}
type bfdXml struct {
	Profile *string       `xml:"profile,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type nexthopXml struct {
	Discard     *nexthopDiscardXml `xml:"discard,omitempty"`
	Ipv6Address *string            `xml:"ipv6-address,omitempty"`
	NextVr      *string            `xml:"next-vr,omitempty"`
	Receive     *nexthopReceiveXml `xml:"receive,omitempty"`
	Misc        []generic.Xml      `xml:",any"`
}
type nexthopDiscardXml struct {
	Misc []generic.Xml `xml:",any"`
}
type nexthopReceiveXml struct {
	Misc []generic.Xml `xml:",any"`
}
type pathMonitorXml struct {
	Enable              *string                                     `xml:"enable,omitempty"`
	FailureCondition    *string                                     `xml:"failure-condition,omitempty"`
	HoldTime            *int64                                      `xml:"hold-time,omitempty"`
	MonitorDestinations *pathMonitorMonitorDestinationsContainerXml `xml:"monitor-destinations,omitempty"`
	Misc                []generic.Xml                               `xml:",any"`
}
type pathMonitorMonitorDestinationsContainerXml struct {
	Entries []pathMonitorMonitorDestinationsXml `xml:"entry"`
}
type pathMonitorMonitorDestinationsXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Enable      *string       `xml:"enable,omitempty"`
	Source      *string       `xml:"source,omitempty"`
	Destination *string       `xml:"destination,omitempty"`
	Interval    *int64        `xml:"interval,omitempty"`
	Count       *int64        `xml:"count,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type routeTableXml struct {
	NoInstall *routeTableNoInstallXml `xml:"no-install,omitempty"`
	Unicast   *routeTableUnicastXml   `xml:"unicast,omitempty"`
	Misc      []generic.Xml           `xml:",any"`
}
type routeTableNoInstallXml struct {
	Misc []generic.Xml `xml:",any"`
}
type routeTableUnicastXml struct {
	Misc []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AdminDist = s.AdminDist
	if s.Bfd != nil {
		var obj bfdXml
		obj.MarshalFromObject(*s.Bfd)
		o.Bfd = &obj
	}
	o.Destination = s.Destination
	o.Interface = s.Interface
	o.Metric = s.Metric
	if s.Nexthop != nil {
		var obj nexthopXml
		obj.MarshalFromObject(*s.Nexthop)
		o.Nexthop = &obj
	}
	if s.PathMonitor != nil {
		var obj pathMonitorXml
		obj.MarshalFromObject(*s.PathMonitor)
		o.PathMonitor = &obj
	}
	if s.RouteTable != nil {
		var obj routeTableXml
		obj.MarshalFromObject(*s.RouteTable)
		o.RouteTable = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var bfdVal *Bfd
	if o.Bfd != nil {
		obj, err := o.Bfd.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bfdVal = obj
	}
	var nexthopVal *Nexthop
	if o.Nexthop != nil {
		obj, err := o.Nexthop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nexthopVal = obj
	}
	var pathMonitorVal *PathMonitor
	if o.PathMonitor != nil {
		obj, err := o.PathMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pathMonitorVal = obj
	}
	var routeTableVal *RouteTable
	if o.RouteTable != nil {
		obj, err := o.RouteTable.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routeTableVal = obj
	}

	result := &Entry{
		Name:        o.Name,
		AdminDist:   o.AdminDist,
		Bfd:         bfdVal,
		Destination: o.Destination,
		Interface:   o.Interface,
		Metric:      o.Metric,
		Nexthop:     nexthopVal,
		PathMonitor: pathMonitorVal,
		RouteTable:  routeTableVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *bfdXml) MarshalFromObject(s Bfd) {
	o.Profile = s.Profile
	o.Misc = s.Misc
}

func (o bfdXml) UnmarshalToObject() (*Bfd, error) {

	result := &Bfd{
		Profile: o.Profile,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *nexthopXml) MarshalFromObject(s Nexthop) {
	if s.Discard != nil {
		var obj nexthopDiscardXml
		obj.MarshalFromObject(*s.Discard)
		o.Discard = &obj
	}
	o.Ipv6Address = s.Ipv6Address
	o.NextVr = s.NextVr
	if s.Receive != nil {
		var obj nexthopReceiveXml
		obj.MarshalFromObject(*s.Receive)
		o.Receive = &obj
	}
	o.Misc = s.Misc
}

func (o nexthopXml) UnmarshalToObject() (*Nexthop, error) {
	var discardVal *NexthopDiscard
	if o.Discard != nil {
		obj, err := o.Discard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		discardVal = obj
	}
	var receiveVal *NexthopReceive
	if o.Receive != nil {
		obj, err := o.Receive.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		receiveVal = obj
	}

	result := &Nexthop{
		Discard:     discardVal,
		Ipv6Address: o.Ipv6Address,
		NextVr:      o.NextVr,
		Receive:     receiveVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *nexthopDiscardXml) MarshalFromObject(s NexthopDiscard) {
	o.Misc = s.Misc
}

func (o nexthopDiscardXml) UnmarshalToObject() (*NexthopDiscard, error) {

	result := &NexthopDiscard{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *nexthopReceiveXml) MarshalFromObject(s NexthopReceive) {
	o.Misc = s.Misc
}

func (o nexthopReceiveXml) UnmarshalToObject() (*NexthopReceive, error) {

	result := &NexthopReceive{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *pathMonitorXml) MarshalFromObject(s PathMonitor) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.FailureCondition = s.FailureCondition
	o.HoldTime = s.HoldTime
	if s.MonitorDestinations != nil {
		var objs []pathMonitorMonitorDestinationsXml
		for _, elt := range s.MonitorDestinations {
			var obj pathMonitorMonitorDestinationsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MonitorDestinations = &pathMonitorMonitorDestinationsContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o pathMonitorXml) UnmarshalToObject() (*PathMonitor, error) {
	var monitorDestinationsVal []PathMonitorMonitorDestinations
	if o.MonitorDestinations != nil {
		for _, elt := range o.MonitorDestinations.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			monitorDestinationsVal = append(monitorDestinationsVal, *obj)
		}
	}

	result := &PathMonitor{
		Enable:              util.AsBool(o.Enable, nil),
		FailureCondition:    o.FailureCondition,
		HoldTime:            o.HoldTime,
		MonitorDestinations: monitorDestinationsVal,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *pathMonitorMonitorDestinationsXml) MarshalFromObject(s PathMonitorMonitorDestinations) {
	o.Name = s.Name
	o.Enable = util.YesNo(s.Enable, nil)
	o.Source = s.Source
	o.Destination = s.Destination
	o.Interval = s.Interval
	o.Count = s.Count
	o.Misc = s.Misc
}

func (o pathMonitorMonitorDestinationsXml) UnmarshalToObject() (*PathMonitorMonitorDestinations, error) {

	result := &PathMonitorMonitorDestinations{
		Name:        o.Name,
		Enable:      util.AsBool(o.Enable, nil),
		Source:      o.Source,
		Destination: o.Destination,
		Interval:    o.Interval,
		Count:       o.Count,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *routeTableXml) MarshalFromObject(s RouteTable) {
	if s.NoInstall != nil {
		var obj routeTableNoInstallXml
		obj.MarshalFromObject(*s.NoInstall)
		o.NoInstall = &obj
	}
	if s.Unicast != nil {
		var obj routeTableUnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
}

func (o routeTableXml) UnmarshalToObject() (*RouteTable, error) {
	var noInstallVal *RouteTableNoInstall
	if o.NoInstall != nil {
		obj, err := o.NoInstall.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noInstallVal = obj
	}
	var unicastVal *RouteTableUnicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &RouteTable{
		NoInstall: noInstallVal,
		Unicast:   unicastVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *routeTableNoInstallXml) MarshalFromObject(s RouteTableNoInstall) {
	o.Misc = s.Misc
}

func (o routeTableNoInstallXml) UnmarshalToObject() (*RouteTableNoInstall, error) {

	result := &RouteTableNoInstall{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *routeTableUnicastXml) MarshalFromObject(s RouteTableUnicast) {
	o.Misc = s.Misc
}

func (o routeTableUnicastXml) UnmarshalToObject() (*RouteTableUnicast, error) {

	result := &RouteTableUnicast{
		Misc: o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "admin_dist" || v == "AdminDist" {
		return e.AdminDist, nil
	}
	if v == "bfd" || v == "Bfd" {
		return e.Bfd, nil
	}
	if v == "destination" || v == "Destination" {
		return e.Destination, nil
	}
	if v == "interface" || v == "Interface" {
		return e.Interface, nil
	}
	if v == "metric" || v == "Metric" {
		return e.Metric, nil
	}
	if v == "nexthop" || v == "Nexthop" {
		return e.Nexthop, nil
	}
	if v == "path_monitor" || v == "PathMonitor" {
		return e.PathMonitor, nil
	}
	if v == "route_table" || v == "RouteTable" {
		return e.RouteTable, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.AdminDist, other.AdminDist) {
		return false
	}
	if !o.Bfd.matches(other.Bfd) {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.Ints64Match(o.Metric, other.Metric) {
		return false
	}
	if !o.Nexthop.matches(other.Nexthop) {
		return false
	}
	if !o.PathMonitor.matches(other.PathMonitor) {
		return false
	}
	if !o.RouteTable.matches(other.RouteTable) {
		return false
	}

	return true
}

func (o *Bfd) matches(other *Bfd) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}

	return true
}

func (o *Nexthop) matches(other *Nexthop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Discard.matches(other.Discard) {
		return false
	}
	if !util.StringsMatch(o.Ipv6Address, other.Ipv6Address) {
		return false
	}
	if !util.StringsMatch(o.NextVr, other.NextVr) {
		return false
	}
	if !o.Receive.matches(other.Receive) {
		return false
	}

	return true
}

func (o *NexthopDiscard) matches(other *NexthopDiscard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *NexthopReceive) matches(other *NexthopReceive) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *PathMonitor) matches(other *PathMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.FailureCondition, other.FailureCondition) {
		return false
	}
	if !util.Ints64Match(o.HoldTime, other.HoldTime) {
		return false
	}
	if len(o.MonitorDestinations) != len(other.MonitorDestinations) {
		return false
	}
	for idx := range o.MonitorDestinations {
		if !o.MonitorDestinations[idx].matches(&other.MonitorDestinations[idx]) {
			return false
		}
	}

	return true
}

func (o *PathMonitorMonitorDestinations) matches(other *PathMonitorMonitorDestinations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.Source, other.Source) {
		return false
	}
	if !util.StringsMatch(o.Destination, other.Destination) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}
	if !util.Ints64Match(o.Count, other.Count) {
		return false
	}

	return true
}

func (o *RouteTable) matches(other *RouteTable) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.NoInstall.matches(other.NoInstall) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *RouteTableNoInstall) matches(other *RouteTableNoInstall) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RouteTableUnicast) matches(other *RouteTableUnicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
