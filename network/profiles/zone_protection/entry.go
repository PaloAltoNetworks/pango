package zone_protection

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

// Entry represents a zone protection profile.
type Entry struct {
	Name                       string
	Description                *string
	Flood                      *Flood
	Scan                       []ScanEntry
	DiscardIpSpoof             *bool
	DiscardStrictSourceRouting *bool
	DiscardLooseSourceRouting  *bool
	DiscardMalformedOption     *bool
	RemoveTcpTimestamp         *bool
	DiscardIpFrag              *bool
	TcpSynWithData             *bool
	StripTcpFastOpenAndData    *bool
	StripMptcpOption           *string
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}

type Flood struct {
	Syn    *FloodSyn
	Icmp   *FloodProtocol
	Icmpv6 *FloodProtocol
	Udp    *FloodProtocol
	Other  *FloodProtocol
}

type FloodSyn struct {
	Enable     *bool
	Red        *FloodRates
	SynCookies *FloodRates
}

type FloodProtocol struct {
	Enable *bool
	Red    *FloodRates
}

type FloodRates struct {
	AlarmRate    *int64
	ActivateRate *int64
	MaximalRate  *int64
}

type ScanEntry struct {
	Name      string
	Action    ScanAction
	Interval  *int64
	Threshold *int64
}

type ScanAction struct {
	Alert   bool
	Block   bool
	BlockIp *ScanBlockIp
}

type ScanBlockIp struct {
	TrackBy  *string
	Duration *int64
}

// EntryObject interface methods.

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}

// XML types.

type entryXml struct {
	XMLName                    xml.Name       `xml:"entry"`
	Name                       string         `xml:"name,attr"`
	Description                *string        `xml:"description,omitempty"`
	Flood                      *floodXml      `xml:"flood,omitempty"`
	Scan                       *scanXml       `xml:"scan,omitempty"`
	DiscardIpSpoof             *string        `xml:"discard-ip-spoof,omitempty"`
	DiscardStrictSourceRouting *string        `xml:"discard-strict-source-routing,omitempty"`
	DiscardLooseSourceRouting  *string        `xml:"discard-loose-source-routing,omitempty"`
	DiscardMalformedOption     *string          `xml:"discard-malformed-option,omitempty"`
	RemoveTcpTimestamp         *string          `xml:"remove-tcp-timestamp,omitempty"`
	DiscardIpFrag              *string          `xml:"discard-ip-frag,omitempty"`
	PacketBased                *packetBasedXml  `xml:"packet-based,omitempty"`
	Misc                       []generic.Xml    `xml:",any"`
	MiscAttributes             []xml.Attr       `xml:",any,attr"`
}

type packetBasedXml struct {
	TcpSynWithData          *string `xml:"tcp-syn-with-data,omitempty"`
	StripTcpFastOpenAndData *string `xml:"strip-tcp-fast-open-and-data,omitempty"`
	StripMptcpOption        *string `xml:"strip-mptcp-option,omitempty"`
}

type floodXml struct {
	Syn    *floodSynXml      `xml:"tcp-syn,omitempty"`
	Icmp   *floodProtocolXml `xml:"icmp,omitempty"`
	Icmpv6 *floodProtocolXml `xml:"icmpv6,omitempty"`
	Udp    *floodProtocolXml `xml:"udp,omitempty"`
	Other  *floodProtocolXml `xml:"other-ip,omitempty"`
}

type floodSynXml struct {
	Enable     *string        `xml:"enable,omitempty"`
	Red        *floodRatesXml `xml:"red,omitempty"`
	SynCookies *floodRatesXml `xml:"syn-cookies,omitempty"`
}

type floodProtocolXml struct {
	Enable *string        `xml:"enable,omitempty"`
	Red    *floodRatesXml `xml:"red,omitempty"`
}

type floodRatesXml struct {
	AlarmRate    *int64 `xml:"alarm-rate,omitempty"`
	ActivateRate *int64 `xml:"activate-rate,omitempty"`
	MaximalRate  *int64 `xml:"maximal-rate,omitempty"`
}

type scanXml struct {
	Entries []scanEntryXml `xml:"entry"`
}

type scanEntryXml struct {
	XMLName   xml.Name       `xml:"entry"`
	Name      string         `xml:"name,attr"`
	Action    *scanActionXml `xml:"action,omitempty"`
	Interval  *int64         `xml:"interval,omitempty"`
	Threshold *int64         `xml:"threshold,omitempty"`
}

type scanActionXml struct {
	Alert   *emptyXml       `xml:"alert"`
	Block   *emptyXml       `xml:"block"`
	BlockIp *scanBlockIpXml `xml:"block-ip"`
}

type emptyXml struct{}

type scanBlockIpXml struct {
	TrackBy  *string `xml:"track-by,omitempty"`
	Duration *int64  `xml:"duration,omitempty"`
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

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DiscardIpSpoof = util.YesNo(s.DiscardIpSpoof, nil)
	o.DiscardStrictSourceRouting = util.YesNo(s.DiscardStrictSourceRouting, nil)
	o.DiscardLooseSourceRouting = util.YesNo(s.DiscardLooseSourceRouting, nil)
	o.DiscardMalformedOption = util.YesNo(s.DiscardMalformedOption, nil)
	o.RemoveTcpTimestamp = util.YesNo(s.RemoveTcpTimestamp, nil)
	o.DiscardIpFrag = util.YesNo(s.DiscardIpFrag, nil)
	tcpSynWithData := util.YesNo(s.TcpSynWithData, nil)
	stripTcpFastOpenAndData := util.YesNo(s.StripTcpFastOpenAndData, nil)
	if tcpSynWithData != nil || stripTcpFastOpenAndData != nil || s.StripMptcpOption != nil {
		o.PacketBased = &packetBasedXml{
			TcpSynWithData:          tcpSynWithData,
			StripTcpFastOpenAndData: stripTcpFastOpenAndData,
			StripMptcpOption:        s.StripMptcpOption,
		}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes

	if len(s.Scan) > 0 {
		scan := &scanXml{}
		for _, se := range s.Scan {
			sxml := scanEntryXml{
				Name:      se.Name,
				Interval:  se.Interval,
				Threshold: se.Threshold,
				Action:    &scanActionXml{},
			}
			if se.Action.Alert {
				sxml.Action.Alert = &emptyXml{}
			}
			if se.Action.Block {
				sxml.Action.Block = &emptyXml{}
			}
			if se.Action.BlockIp != nil {
				sxml.Action.BlockIp = &scanBlockIpXml{
					TrackBy:  se.Action.BlockIp.TrackBy,
					Duration: se.Action.BlockIp.Duration,
				}
			}
			scan.Entries = append(scan.Entries, sxml)
		}
		o.Scan = scan
	}

	if s.Flood != nil {
		flood := &floodXml{}
		if s.Flood.Syn != nil {
			syn := &floodSynXml{
				Enable: util.YesNo(s.Flood.Syn.Enable, nil),
			}
			if s.Flood.Syn.Red != nil {
				syn.Red = &floodRatesXml{
					AlarmRate:    s.Flood.Syn.Red.AlarmRate,
					ActivateRate: s.Flood.Syn.Red.ActivateRate,
					MaximalRate:  s.Flood.Syn.Red.MaximalRate,
				}
			}
			if s.Flood.Syn.SynCookies != nil {
				syn.SynCookies = &floodRatesXml{
					AlarmRate:    s.Flood.Syn.SynCookies.AlarmRate,
					ActivateRate: s.Flood.Syn.SynCookies.ActivateRate,
					MaximalRate:  s.Flood.Syn.SynCookies.MaximalRate,
				}
			}
			flood.Syn = syn
		}
		if s.Flood.Icmp != nil {
			flood.Icmp = marshalFloodProtocol(s.Flood.Icmp)
		}
		if s.Flood.Icmpv6 != nil {
			flood.Icmpv6 = marshalFloodProtocol(s.Flood.Icmpv6)
		}
		if s.Flood.Udp != nil {
			flood.Udp = marshalFloodProtocol(s.Flood.Udp)
		}
		if s.Flood.Other != nil {
			flood.Other = marshalFloodProtocol(s.Flood.Other)
		}
		o.Flood = flood
	}
}

func marshalFloodProtocol(p *FloodProtocol) *floodProtocolXml {
	obj := &floodProtocolXml{
		Enable: util.YesNo(p.Enable, nil),
	}
	if p.Red != nil {
		obj.Red = &floodRatesXml{
			AlarmRate:    p.Red.AlarmRate,
			ActivateRate: p.Red.ActivateRate,
			MaximalRate:  p.Red.MaximalRate,
		}
	}
	return obj
}

func unmarshalFloodProtocol(p *floodProtocolXml) *FloodProtocol {
	if p == nil {
		return nil
	}
	obj := &FloodProtocol{
		Enable: util.AsBool(p.Enable, nil),
	}
	if p.Red != nil {
		obj.Red = &FloodRates{
			AlarmRate:    p.Red.AlarmRate,
			ActivateRate: p.Red.ActivateRate,
			MaximalRate:  p.Red.MaximalRate,
		}
	}
	return obj
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	result := &Entry{
		Name:                       o.Name,
		Description:                o.Description,
		DiscardIpSpoof:             util.AsBool(o.DiscardIpSpoof, nil),
		DiscardStrictSourceRouting: util.AsBool(o.DiscardStrictSourceRouting, nil),
		DiscardLooseSourceRouting:  util.AsBool(o.DiscardLooseSourceRouting, nil),
		DiscardMalformedOption:     util.AsBool(o.DiscardMalformedOption, nil),
		RemoveTcpTimestamp:         util.AsBool(o.RemoveTcpTimestamp, nil),
		DiscardIpFrag:              util.AsBool(o.DiscardIpFrag, nil),
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	if o.PacketBased != nil {
		result.TcpSynWithData = util.AsBool(o.PacketBased.TcpSynWithData, nil)
		result.StripTcpFastOpenAndData = util.AsBool(o.PacketBased.StripTcpFastOpenAndData, nil)
		result.StripMptcpOption = o.PacketBased.StripMptcpOption
	}

	if o.Scan != nil {
		for _, se := range o.Scan.Entries {
			entry := ScanEntry{
				Name:      se.Name,
				Interval:  se.Interval,
				Threshold: se.Threshold,
			}
			if se.Action != nil {
				entry.Action.Alert = se.Action.Alert != nil
				entry.Action.Block = se.Action.Block != nil
				if se.Action.BlockIp != nil {
					entry.Action.BlockIp = &ScanBlockIp{
						TrackBy:  se.Action.BlockIp.TrackBy,
						Duration: se.Action.BlockIp.Duration,
					}
				}
			}
			result.Scan = append(result.Scan, entry)
		}
	}

	if o.Flood != nil {
		flood := &Flood{}
		if o.Flood.Syn != nil {
			syn := &FloodSyn{
				Enable: util.AsBool(o.Flood.Syn.Enable, nil),
			}
			if o.Flood.Syn.Red != nil {
				syn.Red = &FloodRates{
					AlarmRate:    o.Flood.Syn.Red.AlarmRate,
					ActivateRate: o.Flood.Syn.Red.ActivateRate,
					MaximalRate:  o.Flood.Syn.Red.MaximalRate,
				}
			}
			if o.Flood.Syn.SynCookies != nil {
				syn.SynCookies = &FloodRates{
					AlarmRate:    o.Flood.Syn.SynCookies.AlarmRate,
					ActivateRate: o.Flood.Syn.SynCookies.ActivateRate,
					MaximalRate:  o.Flood.Syn.SynCookies.MaximalRate,
				}
			}
			flood.Syn = syn
		}
		flood.Icmp = unmarshalFloodProtocol(o.Flood.Icmp)
		flood.Icmpv6 = unmarshalFloodProtocol(o.Flood.Icmpv6)
		flood.Udp = unmarshalFloodProtocol(o.Flood.Udp)
		flood.Other = unmarshalFloodProtocol(o.Flood.Other)
		result.Flood = flood
	}

	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "flood" || v == "Flood" {
		return e.Flood, nil
	}
	if v == "discard_ip_spoof" || v == "DiscardIpSpoof" {
		return e.DiscardIpSpoof, nil
	}
	if v == "discard_strict_source_routing" || v == "DiscardStrictSourceRouting" {
		return e.DiscardStrictSourceRouting, nil
	}
	if v == "discard_loose_source_routing" || v == "DiscardLooseSourceRouting" {
		return e.DiscardLooseSourceRouting, nil
	}
	if v == "scan" || v == "Scan" {
		return e.Scan, nil
	}
	if v == "discard_malformed_option" || v == "DiscardMalformedOption" {
		return e.DiscardMalformedOption, nil
	}
	if v == "remove_tcp_timestamp" || v == "RemoveTcpTimestamp" {
		return e.RemoveTcpTimestamp, nil
	}
	if v == "discard_ip_frag" || v == "DiscardIpFrag" {
		return e.DiscardIpFrag, nil
	}
	if v == "tcp_syn_with_data" || v == "TcpSynWithData" {
		return e.TcpSynWithData, nil
	}
	if v == "strip_tcp_fast_open_and_data" || v == "StripTcpFastOpenAndData" {
		return e.StripTcpFastOpenAndData, nil
	}
	if v == "strip_mptcp_option" || v == "StripMptcpOption" {
		return e.StripMptcpOption, nil
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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !o.floodMatches(other) {
		return false
	}
	if !util.BoolsMatch(o.DiscardIpSpoof, other.DiscardIpSpoof) {
		return false
	}
	if !util.BoolsMatch(o.DiscardStrictSourceRouting, other.DiscardStrictSourceRouting) {
		return false
	}
	if !util.BoolsMatch(o.DiscardLooseSourceRouting, other.DiscardLooseSourceRouting) {
		return false
	}
	if !scanEntriesMatch(o.Scan, other.Scan) {
		return false
	}
	if !util.BoolsMatch(o.DiscardMalformedOption, other.DiscardMalformedOption) {
		return false
	}
	if !util.BoolsMatch(o.RemoveTcpTimestamp, other.RemoveTcpTimestamp) {
		return false
	}
	if !util.BoolsMatch(o.DiscardIpFrag, other.DiscardIpFrag) {
		return false
	}
	if !util.BoolsMatch(o.TcpSynWithData, other.TcpSynWithData) {
		return false
	}
	if !util.BoolsMatch(o.StripTcpFastOpenAndData, other.StripTcpFastOpenAndData) {
		return false
	}
	if !util.StringsMatch(o.StripMptcpOption, other.StripMptcpOption) {
		return false
	}
	return true
}

func (o *Entry) floodMatches(other *Entry) bool {
	if o.Flood == nil && other.Flood == nil {
		return true
	}
	if (o.Flood == nil) != (other.Flood == nil) {
		return false
	}
	if !floodSynMatches(o.Flood.Syn, other.Flood.Syn) {
		return false
	}
	if !floodProtocolMatches(o.Flood.Icmp, other.Flood.Icmp) {
		return false
	}
	if !floodProtocolMatches(o.Flood.Icmpv6, other.Flood.Icmpv6) {
		return false
	}
	if !floodProtocolMatches(o.Flood.Udp, other.Flood.Udp) {
		return false
	}
	if !floodProtocolMatches(o.Flood.Other, other.Flood.Other) {
		return false
	}
	return true
}

func floodSynMatches(a, b *FloodSyn) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil) != (b == nil) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !floodRatesMatch(a.Red, b.Red) {
		return false
	}
	if !floodRatesMatch(a.SynCookies, b.SynCookies) {
		return false
	}
	return true
}

func floodProtocolMatches(a, b *FloodProtocol) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil) != (b == nil) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !floodRatesMatch(a.Red, b.Red) {
		return false
	}
	return true
}

func scanEntriesMatch(a, b []ScanEntry) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name {
			return false
		}
		if a[i].Action.Alert != b[i].Action.Alert {
			return false
		}
		if a[i].Action.Block != b[i].Action.Block {
			return false
		}
		aBI, bBI := a[i].Action.BlockIp, b[i].Action.BlockIp
		if (aBI == nil) != (bBI == nil) {
			return false
		}
		if aBI != nil {
			if !util.StringsMatch(aBI.TrackBy, bBI.TrackBy) {
				return false
			}
			if !util.Ints64Match(aBI.Duration, bBI.Duration) {
				return false
			}
		}
		if !util.Ints64Match(a[i].Interval, b[i].Interval) {
			return false
		}
		if !util.Ints64Match(a[i].Threshold, b[i].Threshold) {
			return false
		}
	}
	return true
}

func floodRatesMatch(a, b *FloodRates) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil) != (b == nil) {
		return false
	}
	if !util.Ints64Match(a.AlarmRate, b.AlarmRate) {
		return false
	}
	if !util.Ints64Match(a.ActivateRate, b.ActivateRate) {
		return false
	}
	if !util.Ints64Match(a.MaximalRate, b.MaximalRate) {
		return false
	}
	return true
}
