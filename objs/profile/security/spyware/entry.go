package spyware

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// anti-spyware security profile.
//
// PAN-OS 8.0+
type Entry struct {
	Name                string
	Description         string
	PacketCapture       string // 8.x only.
	BotnetLists         []BotnetList
	DnsCategories       []DnsCategory // 10.0
	WhiteLists          []WhiteList   // 10.0
	SinkholeIpv4Address string
	SinkholeIpv6Address string
	ThreatExceptions    []string

	raw map[string]string
}

type BotnetList struct {
	Name          string
	Action        string
	PacketCapture string // 9.0+
}

type DnsCategory struct {
	Name          string
	Action        string
	LogLevel      string
	PacketCapture string
}

type WhiteList struct {
	Name        string
	Description string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.PacketCapture = s.PacketCapture
	o.BotnetLists = s.BotnetLists
	o.SinkholeIpv4Address = s.SinkholeIpv4Address
	o.SinkholeIpv6Address = s.SinkholeIpv6Address
	o.ThreatExceptions = s.ThreatExceptions
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

// 8.0
type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:        o.Name,
		Description: o.Description,
	}

	if o.Botnet != nil {
		ans.PacketCapture = o.Botnet.PacketCapture
		ans.ThreatExceptions = util.EntToStr(o.Botnet.ThreatExceptions)

		if o.Botnet.Lists != nil {
			lists := make([]BotnetList, 0, len(o.Botnet.Lists.Entries))
			for _, x := range o.Botnet.Lists.Entries {
				val := BotnetList{
					Name: x.Name,
				}

				switch {
				case x.Action.Alert != nil:
					val.Action = ActionAlert
				case x.Action.Allow != nil:
					val.Action = ActionAllow
				case x.Action.Block != nil:
					val.Action = ActionBlock
				case x.Action.Sinkhole != nil:
					val.Action = ActionSinkhole
				}

				lists = append(lists, val)
			}

			ans.BotnetLists = lists
		}

		if o.Botnet.Sinkhole != nil {
			ans.SinkholeIpv4Address = o.Botnet.Sinkhole.SinkholeIpv4Address
			ans.SinkholeIpv6Address = o.Botnet.Sinkhole.SinkholeIpv6Address
		}
	}

	raw := make(map[string]string)
	if o.Rules != nil {
		raw["rules"] = util.CleanRawXml(o.Rules.Text)
	}
	if o.ThreatException != nil {
		raw["threat"] = util.CleanRawXml(o.ThreatException.Text)
	}
	if len(raw) > 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v1 struct {
	XMLName         xml.Name     `xml:"entry"`
	Name            string       `xml:"name,attr"`
	Description     string       `xml:"description,omitempty"`
	Botnet          *botnet_v1   `xml:"botnet-domain"`
	Rules           *util.RawXml `xml:"rules"`
	ThreatException *util.RawXml `xml:"threat-exception"`
}

type botnet_v1 struct {
	Lists            *bnLists_v1     `xml:"lists"`
	Sinkhole         *sinkhole       `xml:"sinkhole"`
	PacketCapture    string          `xml:"packet-capture,omitempty"`
	ThreatExceptions *util.EntryType `xml:"threat-exception"`
}

type bnLists_v1 struct {
	Entries []bnlEntry_v1 `xml:"entry"`
}

type bnlEntry_v1 struct {
	Name   string `xml:"name,attr"`
	Action action `xml:"action"`
}

type action struct {
	Alert    *string `xml:"alert"`
	Allow    *string `xml:"allow"`
	Block    *string `xml:"block"`
	Sinkhole *string `xml:"sinkhole"`
}

type sinkhole struct {
	SinkholeIpv4Address string `xml:"ipv4-address,omitempty"`
	SinkholeIpv6Address string `xml:"ipv6-address,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:        e.Name,
		Description: e.Description,
	}

	if e.PacketCapture != "" || len(e.ThreatExceptions) > 0 || len(e.BotnetLists) > 0 || e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" {
		spec := botnet_v1{
			PacketCapture:    e.PacketCapture,
			ThreatExceptions: util.StrToEnt(e.ThreatExceptions),
		}

		if len(e.BotnetLists) > 0 {
			list := make([]bnlEntry_v1, 0, len(e.BotnetLists))
			for _, x := range e.BotnetLists {
				val := bnlEntry_v1{
					Name: x.Name,
				}

				s := ""
				switch x.Action {
				case ActionAlert:
					val.Action.Alert = &s
				case ActionAllow:
					val.Action.Allow = &s
				case ActionBlock:
					val.Action.Block = &s
				case ActionSinkhole:
					val.Action.Sinkhole = &s
				}

				list = append(list, val)
			}

			spec.Lists = &bnLists_v1{Entries: list}
		}

		if e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" {
			spec.Sinkhole = &sinkhole{
				SinkholeIpv4Address: e.SinkholeIpv4Address,
				SinkholeIpv6Address: e.SinkholeIpv6Address,
			}
		}

		ans.Botnet = &spec
	}

	if text, ok := e.raw["rules"]; ok {
		ans.Rules = &util.RawXml{text}
	}

	if text, ok := e.raw["threat"]; ok {
		ans.ThreatException = &util.RawXml{text}
	}

	return ans
}

// 9.0
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o *container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:        o.Name,
		Description: o.Description,
	}

	if o.Botnet != nil {
		ans.ThreatExceptions = util.EntToStr(o.Botnet.ThreatExceptions)

		if o.Botnet.Lists != nil {
			lists := make([]BotnetList, 0, len(o.Botnet.Lists.Entries))
			for _, x := range o.Botnet.Lists.Entries {
				val := BotnetList{
					Name:          x.Name,
					PacketCapture: x.PacketCapture,
				}

				switch {
				case x.Action.Alert != nil:
					val.Action = ActionAlert
				case x.Action.Allow != nil:
					val.Action = ActionAllow
				case x.Action.Block != nil:
					val.Action = ActionBlock
				case x.Action.Sinkhole != nil:
					val.Action = ActionSinkhole
				}

				lists = append(lists, val)
			}

			ans.BotnetLists = lists
		}

		if o.Botnet.Sinkhole != nil {
			ans.SinkholeIpv4Address = o.Botnet.Sinkhole.SinkholeIpv4Address
			ans.SinkholeIpv6Address = o.Botnet.Sinkhole.SinkholeIpv6Address
		}
	}

	raw := make(map[string]string)
	if o.Rules != nil {
		raw["rules"] = util.CleanRawXml(o.Rules.Text)
	}
	if o.ThreatException != nil {
		raw["threat"] = util.CleanRawXml(o.ThreatException.Text)
	}
	if len(raw) > 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v2 struct {
	XMLName         xml.Name     `xml:"entry"`
	Name            string       `xml:"name,attr"`
	Description     string       `xml:"description,omitempty"`
	Botnet          *botnet_v2   `xml:"botnet-domain"`
	Rules           *util.RawXml `xml:"rules"`
	ThreatException *util.RawXml `xml:"threat-exception"`
}

type botnet_v2 struct {
	Lists            *bnLists_v2     `xml:"lists"`
	Sinkhole         *sinkhole       `xml:"sinkhole"`
	ThreatExceptions *util.EntryType `xml:"threat-exception"`
}

type bnLists_v2 struct {
	Entries []bnlEntry_v2 `xml:"entry"`
}

type bnlEntry_v2 struct {
	Name          string `xml:"name,attr"`
	Action        action `xml:"action"`
	PacketCapture string `xml:"packet-capture,omitempty"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:        e.Name,
		Description: e.Description,
	}

	if len(e.ThreatExceptions) > 0 || len(e.BotnetLists) > 0 || e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" {
		spec := botnet_v2{
			ThreatExceptions: util.StrToEnt(e.ThreatExceptions),
		}

		if len(e.BotnetLists) > 0 {
			list := make([]bnlEntry_v2, 0, len(e.BotnetLists))
			for _, x := range e.BotnetLists {
				val := bnlEntry_v2{
					Name:          x.Name,
					PacketCapture: x.PacketCapture,
				}

				s := ""
				switch x.Action {
				case ActionAlert:
					val.Action.Alert = &s
				case ActionAllow:
					val.Action.Allow = &s
				case ActionBlock:
					val.Action.Block = &s
				case ActionSinkhole:
					val.Action.Sinkhole = &s
				}

				list = append(list, val)
			}

			spec.Lists = &bnLists_v2{Entries: list}
		}

		if e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" {
			spec.Sinkhole = &sinkhole{
				SinkholeIpv4Address: e.SinkholeIpv4Address,
				SinkholeIpv6Address: e.SinkholeIpv6Address,
			}
		}

		ans.Botnet = &spec
	}

	if text, ok := e.raw["rules"]; ok {
		ans.Rules = &util.RawXml{text}
	}

	if text, ok := e.raw["threat"]; ok {
		ans.ThreatException = &util.RawXml{text}
	}

	return ans
}

// 10.0
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o *container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v3) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:        o.Name,
		Description: o.Description,
	}

	if o.Botnet != nil {
		ans.ThreatExceptions = util.EntToStr(o.Botnet.ThreatExceptions)

		if o.Botnet.Lists != nil {
			lists := make([]BotnetList, 0, len(o.Botnet.Lists.Entries))
			for _, x := range o.Botnet.Lists.Entries {
				val := BotnetList{
					Name:          x.Name,
					PacketCapture: x.PacketCapture,
				}

				switch {
				case x.Action.Alert != nil:
					val.Action = ActionAlert
				case x.Action.Allow != nil:
					val.Action = ActionAllow
				case x.Action.Block != nil:
					val.Action = ActionBlock
				case x.Action.Sinkhole != nil:
					val.Action = ActionSinkhole
				}

				lists = append(lists, val)
			}

			ans.BotnetLists = lists
		}

		if o.Botnet.Dns != nil {
			list := make([]DnsCategory, 0, len(o.Botnet.Dns.Entries))
			for _, x := range o.Botnet.Dns.Entries {
				list = append(list, DnsCategory{
					Name:          x.Name,
					Action:        x.Action,
					LogLevel:      x.LogLevel,
					PacketCapture: x.PacketCapture,
				})
			}

			ans.DnsCategories = list
		}

		if o.Botnet.WhiteLists != nil {
			list := make([]WhiteList, 0, len(o.Botnet.WhiteLists.Entries))
			for _, x := range o.Botnet.WhiteLists.Entries {
				list = append(list, WhiteList{
					Name:        x.Name,
					Description: x.Description,
				})
			}

			ans.WhiteLists = list
		}

		if o.Botnet.Sinkhole != nil {
			ans.SinkholeIpv4Address = o.Botnet.Sinkhole.SinkholeIpv4Address
			ans.SinkholeIpv6Address = o.Botnet.Sinkhole.SinkholeIpv6Address
		}
	}

	raw := make(map[string]string)
	if o.Rules != nil {
		raw["rules"] = util.CleanRawXml(o.Rules.Text)
	}
	if o.ThreatException != nil {
		raw["threat"] = util.CleanRawXml(o.ThreatException.Text)
	}
	if len(raw) > 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v3 struct {
	XMLName         xml.Name     `xml:"entry"`
	Name            string       `xml:"name,attr"`
	Description     string       `xml:"description,omitempty"`
	Botnet          *botnet_v3   `xml:"botnet-domain"`
	Rules           *util.RawXml `xml:"rules"`
	ThreatException *util.RawXml `xml:"threat-exception"`
}

type botnet_v3 struct {
	Lists            *bnLists_v2     `xml:"lists"`
	Dns              *dns            `xml:"dns-security-categories"`
	WhiteLists       *whiteList      `xml:"whitelist"`
	Sinkhole         *sinkhole       `xml:"sinkhole"`
	ThreatExceptions *util.EntryType `xml:"threat-exception"`
}

type dns struct {
	Entries []dnsEntry `xml:"entry"`
}

type dnsEntry struct {
	Name          string `xml:"name,attr"`
	Action        string `xml:"action,omitempty"`
	LogLevel      string `xml:"log-level,omitempty"`
	PacketCapture string `xml:"packet-capture,omitempty"`
}

type whiteList struct {
	Entries []wlEntry `xml:"entry"`
}

type wlEntry struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,omitempty"`
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:        e.Name,
		Description: e.Description,
	}

	if len(e.ThreatExceptions) > 0 || len(e.BotnetLists) > 0 || e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" || len(e.DnsCategories) > 0 || len(e.WhiteLists) > 0 {
		spec := botnet_v3{
			ThreatExceptions: util.StrToEnt(e.ThreatExceptions),
		}

		if len(e.DnsCategories) > 0 {
			list := make([]dnsEntry, 0, len(e.DnsCategories))
			for _, x := range e.DnsCategories {
				list = append(list, dnsEntry{
					Name:          x.Name,
					Action:        x.Action,
					LogLevel:      x.LogLevel,
					PacketCapture: x.PacketCapture,
				})
			}

			spec.Dns = &dns{Entries: list}
		}

		if len(e.WhiteLists) > 0 {
			list := make([]wlEntry, 0, len(e.WhiteLists))
			for _, x := range e.WhiteLists {
				list = append(list, wlEntry{
					Name:        x.Name,
					Description: x.Description,
				})
			}

			spec.WhiteLists = &whiteList{Entries: list}
		}

		if len(e.BotnetLists) > 0 {
			list := make([]bnlEntry_v2, 0, len(e.BotnetLists))
			for _, x := range e.BotnetLists {
				val := bnlEntry_v2{
					Name:          x.Name,
					PacketCapture: x.PacketCapture,
				}

				s := ""
				switch x.Action {
				case ActionAlert:
					val.Action.Alert = &s
				case ActionAllow:
					val.Action.Allow = &s
				case ActionBlock:
					val.Action.Block = &s
				case ActionSinkhole:
					val.Action.Sinkhole = &s
				}

				list = append(list, val)
			}

			spec.Lists = &bnLists_v2{Entries: list}
		}

		if e.SinkholeIpv4Address != "" || e.SinkholeIpv6Address != "" {
			spec.Sinkhole = &sinkhole{
				SinkholeIpv4Address: e.SinkholeIpv4Address,
				SinkholeIpv6Address: e.SinkholeIpv6Address,
			}
		}

		ans.Botnet = &spec
	}

	if text, ok := e.raw["rules"]; ok {
		ans.Rules = &util.RawXml{text}
	}

	if text, ok := e.raw["threat"]; ok {
		ans.ThreatException = &util.RawXml{text}
	}

	return ans
}
