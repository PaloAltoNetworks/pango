package dosprotection

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
	suffix = []string{"profiles", "dos-protection", "$name"}
)

type Entry struct {
	Name            string
	Description     *string
	DisableOverride *string
	Flood           *Flood
	Resource        *Resource
	Type            *string
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type Flood struct {
	Icmp           *FloodIcmp
	Icmpv6         *FloodIcmpv6
	OtherIp        *FloodOtherIp
	TcpSyn         *FloodTcpSyn
	Udp            *FloodUdp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmp struct {
	Enable         *bool
	Red            *FloodIcmpRed
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmpRed struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodIcmpRedBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmpRedBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmpv6 struct {
	Enable         *bool
	Red            *FloodIcmpv6Red
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmpv6Red struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodIcmpv6RedBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodIcmpv6RedBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodOtherIp struct {
	Enable         *bool
	Red            *FloodOtherIpRed
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodOtherIpRed struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodOtherIpRedBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodOtherIpRedBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodTcpSyn struct {
	Enable         *bool
	Red            *FloodTcpSynRed
	SynCookies     *FloodTcpSynSynCookies
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodTcpSynRed struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodTcpSynRedBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodTcpSynRedBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodTcpSynSynCookies struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodTcpSynSynCookiesBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodTcpSynSynCookiesBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodUdp struct {
	Enable         *bool
	Red            *FloodUdpRed
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodUdpRed struct {
	ActivateRate   *int64
	AlarmRate      *int64
	Block          *FloodUdpRedBlock
	MaximalRate    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FloodUdpRedBlock struct {
	Duration       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Resource struct {
	Sessions       *ResourceSessions
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ResourceSessions struct {
	Enabled            *bool
	MaxConcurrentLimit *int64
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
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
	XMLName         xml.Name      `xml:"entry"`
	Name            string        `xml:"name,attr"`
	Description     *string       `xml:"description,omitempty"`
	DisableOverride *string       `xml:"disable-override,omitempty"`
	Flood           *floodXml     `xml:"flood,omitempty"`
	Resource        *resourceXml  `xml:"resource,omitempty"`
	Type            *string       `xml:"type,omitempty"`
	Misc            []generic.Xml `xml:",any"`
	MiscAttributes  []xml.Attr    `xml:",any,attr"`
}
type floodXml struct {
	Icmp           *floodIcmpXml    `xml:"icmp,omitempty"`
	Icmpv6         *floodIcmpv6Xml  `xml:"icmpv6,omitempty"`
	OtherIp        *floodOtherIpXml `xml:"other-ip,omitempty"`
	TcpSyn         *floodTcpSynXml  `xml:"tcp-syn,omitempty"`
	Udp            *floodUdpXml     `xml:"udp,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type floodIcmpXml struct {
	Enable         *string          `xml:"enable,omitempty"`
	Red            *floodIcmpRedXml `xml:"red,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type floodIcmpRedXml struct {
	ActivateRate   *int64                `xml:"activate-rate,omitempty"`
	AlarmRate      *int64                `xml:"alarm-rate,omitempty"`
	Block          *floodIcmpRedBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64                `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
	MiscAttributes []xml.Attr            `xml:",any,attr"`
}
type floodIcmpRedBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type floodIcmpv6Xml struct {
	Enable         *string            `xml:"enable,omitempty"`
	Red            *floodIcmpv6RedXml `xml:"red,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
	MiscAttributes []xml.Attr         `xml:",any,attr"`
}
type floodIcmpv6RedXml struct {
	ActivateRate   *int64                  `xml:"activate-rate,omitempty"`
	AlarmRate      *int64                  `xml:"alarm-rate,omitempty"`
	Block          *floodIcmpv6RedBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64                  `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type floodIcmpv6RedBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type floodOtherIpXml struct {
	Enable         *string             `xml:"enable,omitempty"`
	Red            *floodOtherIpRedXml `xml:"red,omitempty"`
	Misc           []generic.Xml       `xml:",any"`
	MiscAttributes []xml.Attr          `xml:",any,attr"`
}
type floodOtherIpRedXml struct {
	ActivateRate   *int64                   `xml:"activate-rate,omitempty"`
	AlarmRate      *int64                   `xml:"alarm-rate,omitempty"`
	Block          *floodOtherIpRedBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64                   `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type floodOtherIpRedBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type floodTcpSynXml struct {
	Enable         *string                   `xml:"enable,omitempty"`
	Red            *floodTcpSynRedXml        `xml:"red,omitempty"`
	SynCookies     *floodTcpSynSynCookiesXml `xml:"syn-cookies,omitempty"`
	Misc           []generic.Xml             `xml:",any"`
	MiscAttributes []xml.Attr                `xml:",any,attr"`
}
type floodTcpSynRedXml struct {
	ActivateRate   *int64                  `xml:"activate-rate,omitempty"`
	AlarmRate      *int64                  `xml:"alarm-rate,omitempty"`
	Block          *floodTcpSynRedBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64                  `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type floodTcpSynRedBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type floodTcpSynSynCookiesXml struct {
	ActivateRate   *int64                         `xml:"activate-rate,omitempty"`
	AlarmRate      *int64                         `xml:"alarm-rate,omitempty"`
	Block          *floodTcpSynSynCookiesBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64                         `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type floodTcpSynSynCookiesBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type floodUdpXml struct {
	Enable         *string         `xml:"enable,omitempty"`
	Red            *floodUdpRedXml `xml:"red,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type floodUdpRedXml struct {
	ActivateRate   *int64               `xml:"activate-rate,omitempty"`
	AlarmRate      *int64               `xml:"alarm-rate,omitempty"`
	Block          *floodUdpRedBlockXml `xml:"block,omitempty"`
	MaximalRate    *int64               `xml:"maximal-rate,omitempty"`
	Misc           []generic.Xml        `xml:",any"`
	MiscAttributes []xml.Attr           `xml:",any,attr"`
}
type floodUdpRedBlockXml struct {
	Duration       *int64        `xml:"duration,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type resourceXml struct {
	Sessions       *resourceSessionsXml `xml:"sessions,omitempty"`
	Misc           []generic.Xml        `xml:",any"`
	MiscAttributes []xml.Attr           `xml:",any,attr"`
}
type resourceSessionsXml struct {
	Enabled            *string       `xml:"enabled,omitempty"`
	MaxConcurrentLimit *int64        `xml:"max-concurrent-limit,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.Flood != nil {
		var obj floodXml
		obj.MarshalFromObject(*s.Flood)
		o.Flood = &obj
	}
	if s.Resource != nil {
		var obj resourceXml
		obj.MarshalFromObject(*s.Resource)
		o.Resource = &obj
	}
	o.Type = s.Type
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var floodVal *Flood
	if o.Flood != nil {
		obj, err := o.Flood.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floodVal = obj
	}
	var resourceVal *Resource
	if o.Resource != nil {
		obj, err := o.Resource.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resourceVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Description:     o.Description,
		DisableOverride: o.DisableOverride,
		Flood:           floodVal,
		Resource:        resourceVal,
		Type:            o.Type,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *floodXml) MarshalFromObject(s Flood) {
	if s.Icmp != nil {
		var obj floodIcmpXml
		obj.MarshalFromObject(*s.Icmp)
		o.Icmp = &obj
	}
	if s.Icmpv6 != nil {
		var obj floodIcmpv6Xml
		obj.MarshalFromObject(*s.Icmpv6)
		o.Icmpv6 = &obj
	}
	if s.OtherIp != nil {
		var obj floodOtherIpXml
		obj.MarshalFromObject(*s.OtherIp)
		o.OtherIp = &obj
	}
	if s.TcpSyn != nil {
		var obj floodTcpSynXml
		obj.MarshalFromObject(*s.TcpSyn)
		o.TcpSyn = &obj
	}
	if s.Udp != nil {
		var obj floodUdpXml
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodXml) UnmarshalToObject() (*Flood, error) {
	var icmpVal *FloodIcmp
	if o.Icmp != nil {
		obj, err := o.Icmp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		icmpVal = obj
	}
	var icmpv6Val *FloodIcmpv6
	if o.Icmpv6 != nil {
		obj, err := o.Icmpv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		icmpv6Val = obj
	}
	var otherIpVal *FloodOtherIp
	if o.OtherIp != nil {
		obj, err := o.OtherIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		otherIpVal = obj
	}
	var tcpSynVal *FloodTcpSyn
	if o.TcpSyn != nil {
		obj, err := o.TcpSyn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpSynVal = obj
	}
	var udpVal *FloodUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &Flood{
		Icmp:           icmpVal,
		Icmpv6:         icmpv6Val,
		OtherIp:        otherIpVal,
		TcpSyn:         tcpSynVal,
		Udp:            udpVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpXml) MarshalFromObject(s FloodIcmp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Red != nil {
		var obj floodIcmpRedXml
		obj.MarshalFromObject(*s.Red)
		o.Red = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpXml) UnmarshalToObject() (*FloodIcmp, error) {
	var redVal *FloodIcmpRed
	if o.Red != nil {
		obj, err := o.Red.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redVal = obj
	}

	result := &FloodIcmp{
		Enable:         util.AsBool(o.Enable, nil),
		Red:            redVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpRedXml) MarshalFromObject(s FloodIcmpRed) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodIcmpRedBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpRedXml) UnmarshalToObject() (*FloodIcmpRed, error) {
	var blockVal *FloodIcmpRedBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodIcmpRed{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpRedBlockXml) MarshalFromObject(s FloodIcmpRedBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpRedBlockXml) UnmarshalToObject() (*FloodIcmpRedBlock, error) {

	result := &FloodIcmpRedBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpv6Xml) MarshalFromObject(s FloodIcmpv6) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Red != nil {
		var obj floodIcmpv6RedXml
		obj.MarshalFromObject(*s.Red)
		o.Red = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpv6Xml) UnmarshalToObject() (*FloodIcmpv6, error) {
	var redVal *FloodIcmpv6Red
	if o.Red != nil {
		obj, err := o.Red.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redVal = obj
	}

	result := &FloodIcmpv6{
		Enable:         util.AsBool(o.Enable, nil),
		Red:            redVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpv6RedXml) MarshalFromObject(s FloodIcmpv6Red) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodIcmpv6RedBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpv6RedXml) UnmarshalToObject() (*FloodIcmpv6Red, error) {
	var blockVal *FloodIcmpv6RedBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodIcmpv6Red{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodIcmpv6RedBlockXml) MarshalFromObject(s FloodIcmpv6RedBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodIcmpv6RedBlockXml) UnmarshalToObject() (*FloodIcmpv6RedBlock, error) {

	result := &FloodIcmpv6RedBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodOtherIpXml) MarshalFromObject(s FloodOtherIp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Red != nil {
		var obj floodOtherIpRedXml
		obj.MarshalFromObject(*s.Red)
		o.Red = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodOtherIpXml) UnmarshalToObject() (*FloodOtherIp, error) {
	var redVal *FloodOtherIpRed
	if o.Red != nil {
		obj, err := o.Red.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redVal = obj
	}

	result := &FloodOtherIp{
		Enable:         util.AsBool(o.Enable, nil),
		Red:            redVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodOtherIpRedXml) MarshalFromObject(s FloodOtherIpRed) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodOtherIpRedBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodOtherIpRedXml) UnmarshalToObject() (*FloodOtherIpRed, error) {
	var blockVal *FloodOtherIpRedBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodOtherIpRed{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodOtherIpRedBlockXml) MarshalFromObject(s FloodOtherIpRedBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodOtherIpRedBlockXml) UnmarshalToObject() (*FloodOtherIpRedBlock, error) {

	result := &FloodOtherIpRedBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodTcpSynXml) MarshalFromObject(s FloodTcpSyn) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Red != nil {
		var obj floodTcpSynRedXml
		obj.MarshalFromObject(*s.Red)
		o.Red = &obj
	}
	if s.SynCookies != nil {
		var obj floodTcpSynSynCookiesXml
		obj.MarshalFromObject(*s.SynCookies)
		o.SynCookies = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodTcpSynXml) UnmarshalToObject() (*FloodTcpSyn, error) {
	var redVal *FloodTcpSynRed
	if o.Red != nil {
		obj, err := o.Red.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redVal = obj
	}
	var synCookiesVal *FloodTcpSynSynCookies
	if o.SynCookies != nil {
		obj, err := o.SynCookies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		synCookiesVal = obj
	}

	result := &FloodTcpSyn{
		Enable:         util.AsBool(o.Enable, nil),
		Red:            redVal,
		SynCookies:     synCookiesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodTcpSynRedXml) MarshalFromObject(s FloodTcpSynRed) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodTcpSynRedBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodTcpSynRedXml) UnmarshalToObject() (*FloodTcpSynRed, error) {
	var blockVal *FloodTcpSynRedBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodTcpSynRed{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodTcpSynRedBlockXml) MarshalFromObject(s FloodTcpSynRedBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodTcpSynRedBlockXml) UnmarshalToObject() (*FloodTcpSynRedBlock, error) {

	result := &FloodTcpSynRedBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodTcpSynSynCookiesXml) MarshalFromObject(s FloodTcpSynSynCookies) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodTcpSynSynCookiesBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodTcpSynSynCookiesXml) UnmarshalToObject() (*FloodTcpSynSynCookies, error) {
	var blockVal *FloodTcpSynSynCookiesBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodTcpSynSynCookies{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodTcpSynSynCookiesBlockXml) MarshalFromObject(s FloodTcpSynSynCookiesBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodTcpSynSynCookiesBlockXml) UnmarshalToObject() (*FloodTcpSynSynCookiesBlock, error) {

	result := &FloodTcpSynSynCookiesBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodUdpXml) MarshalFromObject(s FloodUdp) {
	o.Enable = util.YesNo(s.Enable, nil)
	if s.Red != nil {
		var obj floodUdpRedXml
		obj.MarshalFromObject(*s.Red)
		o.Red = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodUdpXml) UnmarshalToObject() (*FloodUdp, error) {
	var redVal *FloodUdpRed
	if o.Red != nil {
		obj, err := o.Red.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		redVal = obj
	}

	result := &FloodUdp{
		Enable:         util.AsBool(o.Enable, nil),
		Red:            redVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodUdpRedXml) MarshalFromObject(s FloodUdpRed) {
	o.ActivateRate = s.ActivateRate
	o.AlarmRate = s.AlarmRate
	if s.Block != nil {
		var obj floodUdpRedBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	o.MaximalRate = s.MaximalRate
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodUdpRedXml) UnmarshalToObject() (*FloodUdpRed, error) {
	var blockVal *FloodUdpRedBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}

	result := &FloodUdpRed{
		ActivateRate:   o.ActivateRate,
		AlarmRate:      o.AlarmRate,
		Block:          blockVal,
		MaximalRate:    o.MaximalRate,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *floodUdpRedBlockXml) MarshalFromObject(s FloodUdpRedBlock) {
	o.Duration = s.Duration
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o floodUdpRedBlockXml) UnmarshalToObject() (*FloodUdpRedBlock, error) {

	result := &FloodUdpRedBlock{
		Duration:       o.Duration,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *resourceXml) MarshalFromObject(s Resource) {
	if s.Sessions != nil {
		var obj resourceSessionsXml
		obj.MarshalFromObject(*s.Sessions)
		o.Sessions = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o resourceXml) UnmarshalToObject() (*Resource, error) {
	var sessionsVal *ResourceSessions
	if o.Sessions != nil {
		obj, err := o.Sessions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sessionsVal = obj
	}

	result := &Resource{
		Sessions:       sessionsVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *resourceSessionsXml) MarshalFromObject(s ResourceSessions) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.MaxConcurrentLimit = s.MaxConcurrentLimit
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o resourceSessionsXml) UnmarshalToObject() (*ResourceSessions, error) {

	result := &ResourceSessions{
		Enabled:            util.AsBool(o.Enabled, nil),
		MaxConcurrentLimit: o.MaxConcurrentLimit,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
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
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "flood" || v == "Flood" {
		return e.Flood, nil
	}
	if v == "resource" || v == "Resource" {
		return e.Resource, nil
	}
	if v == "type" || v == "Type" {
		return e.Type, nil
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
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !o.Flood.matches(other.Flood) {
		return false
	}
	if !o.Resource.matches(other.Resource) {
		return false
	}
	if !util.StringsMatch(o.Type, other.Type) {
		return false
	}

	return true
}

func (o *Flood) matches(other *Flood) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Icmp.matches(other.Icmp) {
		return false
	}
	if !o.Icmpv6.matches(other.Icmpv6) {
		return false
	}
	if !o.OtherIp.matches(other.OtherIp) {
		return false
	}
	if !o.TcpSyn.matches(other.TcpSyn) {
		return false
	}
	if !o.Udp.matches(other.Udp) {
		return false
	}

	return true
}

func (o *FloodIcmp) matches(other *FloodIcmp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Red.matches(other.Red) {
		return false
	}

	return true
}

func (o *FloodIcmpRed) matches(other *FloodIcmpRed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodIcmpRedBlock) matches(other *FloodIcmpRedBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *FloodIcmpv6) matches(other *FloodIcmpv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Red.matches(other.Red) {
		return false
	}

	return true
}

func (o *FloodIcmpv6Red) matches(other *FloodIcmpv6Red) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodIcmpv6RedBlock) matches(other *FloodIcmpv6RedBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *FloodOtherIp) matches(other *FloodOtherIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Red.matches(other.Red) {
		return false
	}

	return true
}

func (o *FloodOtherIpRed) matches(other *FloodOtherIpRed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodOtherIpRedBlock) matches(other *FloodOtherIpRedBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *FloodTcpSyn) matches(other *FloodTcpSyn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Red.matches(other.Red) {
		return false
	}
	if !o.SynCookies.matches(other.SynCookies) {
		return false
	}

	return true
}

func (o *FloodTcpSynRed) matches(other *FloodTcpSynRed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodTcpSynRedBlock) matches(other *FloodTcpSynRedBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *FloodTcpSynSynCookies) matches(other *FloodTcpSynSynCookies) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodTcpSynSynCookiesBlock) matches(other *FloodTcpSynSynCookiesBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *FloodUdp) matches(other *FloodUdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.Red.matches(other.Red) {
		return false
	}

	return true
}

func (o *FloodUdpRed) matches(other *FloodUdpRed) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ActivateRate, other.ActivateRate) {
		return false
	}
	if !util.Ints64Match(o.AlarmRate, other.AlarmRate) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !util.Ints64Match(o.MaximalRate, other.MaximalRate) {
		return false
	}

	return true
}

func (o *FloodUdpRedBlock) matches(other *FloodUdpRedBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *Resource) matches(other *Resource) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Sessions.matches(other.Sessions) {
		return false
	}

	return true
}

func (o *ResourceSessions) matches(other *ResourceSessions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.Ints64Match(o.MaxConcurrentLimit, other.MaxConcurrentLimit) {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
