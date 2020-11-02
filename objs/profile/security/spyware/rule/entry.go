package rule

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// anti-spyware security rule.
//
// PAN-OS 7.0+
type Entry struct {
	Name            string
	ThreatName      string
	Category        string
	Severities      []string
	PacketCapture   string
	Action          string
	BlockIpTrackBy  string
	BlockIpDuration int
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.PacketCapture = s.PacketCapture
	o.ThreatName = s.ThreatName
	o.Category = s.Category
	o.Severities = s.Severities
	o.Action = s.Action
	o.BlockIpTrackBy = s.BlockIpTrackBy
	o.BlockIpDuration = s.BlockIpDuration
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
		Name:          o.Name,
		ThreatName:    o.ThreatName,
		Category:      o.Category,
		Severities:    util.MemToStr(o.Severities),
		PacketCapture: o.PacketCapture,
	}

	if o.Action != nil {
		switch {
		case o.Action.Default != nil:
			ans.Action = ActionDefault
		case o.Action.Allow != nil:
			ans.Action = ActionAllow
		case o.Action.Alert != nil:
			ans.Action = ActionAlert
		case o.Action.Drop != nil:
			ans.Action = ActionDrop
		case o.Action.ResetClient != nil:
			ans.Action = ActionResetClient
		case o.Action.ResetServer != nil:
			ans.Action = ActionResetServer
		case o.Action.ResetBoth != nil:
			ans.Action = ActionResetBoth
		case o.Action.BlockIp != nil:
			ans.Action = ActionBlockIp
			ans.BlockIpTrackBy = o.Action.BlockIp.TrackBy
			ans.BlockIpDuration = o.Action.BlockIp.Duration
		}
	}

	return ans
}

type entry_v1 struct {
	XMLName       xml.Name         `xml:"entry"`
	Name          string           `xml:"name,attr"`
	ThreatName    string           `xml:"threat-name,omitempty"`
	Category      string           `xml:"category"`
	Severities    *util.MemberType `xml:"severity"`
	PacketCapture string           `xml:"packet-capture,omitempty"`
	Action        *action          `xml:"action"`
}

type action struct {
	Default     *string  `xml:"default"`
	Allow       *string  `xml:"allow"`
	Alert       *string  `xml:"alert"`
	Drop        *string  `xml:"drop"`
	ResetClient *string  `xml:"reset-client"`
	ResetServer *string  `xml:"reset-server"`
	ResetBoth   *string  `xml:"reset-both"`
	BlockIp     *blockIp `xml:"block-ip"`
}

type blockIp struct {
	TrackBy  string `xml:"track-by"`
	Duration int    `xml:"duration"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:          e.Name,
		ThreatName:    e.ThreatName,
		Category:      e.Category,
		Severities:    util.StrToMem(e.Severities),
		PacketCapture: e.PacketCapture,
	}

	s := ""
	switch e.Action {
	case ActionDefault:
		ans.Action = &action{
			Default: &s,
		}
	case ActionAllow:
		ans.Action = &action{
			Allow: &s,
		}
	case ActionAlert:
		ans.Action = &action{
			Alert: &s,
		}
	case ActionDrop:
		ans.Action = &action{
			Drop: &s,
		}
	case ActionResetClient:
		ans.Action = &action{
			ResetClient: &s,
		}
	case ActionResetServer:
		ans.Action = &action{
			ResetServer: &s,
		}
	case ActionResetBoth:
		ans.Action = &action{
			ResetBoth: &s,
		}
	case ActionBlockIp:
		ans.Action = &action{
			BlockIp: &blockIp{
				TrackBy:  e.BlockIpTrackBy,
				Duration: e.BlockIpDuration,
			},
		}
	}

	return ans
}
