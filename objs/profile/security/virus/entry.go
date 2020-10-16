package virus

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a log forwarding profile.
//
// PAN-OS 8.0+.
type Entry struct {
	Name                  string
	Description           string
	PacketCapture         bool
	Decoders              []Decoder
	ApplicationExceptions []ApplicationException
	ThreatExceptions      []string
}

type Decoder struct {
	Name           string
	Action         string
	WildfireAction string
}

type ApplicationException struct {
	Application string
	Action      string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.PacketCapture = s.PacketCapture
	o.Decoders = s.Decoders
	o.ApplicationExceptions = s.ApplicationExceptions
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
		Name:             o.Name,
		Description:      o.Description,
		PacketCapture:    util.AsBool(o.PacketCapture),
		ThreatExceptions: util.EntToStr(o.ThreatExceptions),
	}

	if o.Decoder != nil {
		data := make([]Decoder, 0, len(o.Decoder.Entries))
		for _, d := range o.Decoder.Entries {
			data = append(data, Decoder{
				Name:           d.Name,
				Action:         d.Action,
				WildfireAction: d.WildfireAction,
			})
		}

		ans.Decoders = data
	}

	if o.Application != nil {
		data := make([]ApplicationException, 0, len(o.Application.Entries))
		for _, d := range o.Application.Entries {
			data = append(data, ApplicationException{
				Application: d.Application,
				Action:      d.Action,
			})
		}

		ans.ApplicationExceptions = data
	}

	return ans
}

type entry_v1 struct {
	XMLName          xml.Name        `xml:"entry"`
	Name             string          `xml:"name,attr"`
	Description      string          `xml:"description,omitempty"`
	PacketCapture    string          `xml:"packet-capture"`
	Decoder          *decoder        `xml:"decoder"`
	Application      *application    `xml:"application"`
	ThreatExceptions *util.EntryType `xml:"threat-exception"`
}

type decoder struct {
	Entries []decoderEntry `xml:"entry"`
}

type decoderEntry struct {
	Name           string `xml:"name,attr"`
	Action         string `xml:"action,omitempty"`
	WildfireAction string `xml:"wildfire-action,omitempty"`
}

type application struct {
	Entries []applicationEntry `xml:"entry"`
}

type applicationEntry struct {
	Application string `xml:"name,attr"`
	Action      string `xml:"action,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:             e.Name,
		Description:      e.Description,
		PacketCapture:    util.YesNo(e.PacketCapture),
		ThreatExceptions: util.StrToEnt(e.ThreatExceptions),
	}

	if len(e.Decoders) > 0 {
		data := make([]decoderEntry, 0, len(e.Decoders))
		for _, d := range e.Decoders {
			data = append(data, decoderEntry{
				Name:           d.Name,
				Action:         d.Action,
				WildfireAction: d.WildfireAction,
			})
		}

		ans.Decoder = &decoder{Entries: data}
	}

	if len(e.ApplicationExceptions) > 0 {
		data := make([]applicationEntry, 0, len(e.ApplicationExceptions))
		for _, d := range e.ApplicationExceptions {
			data = append(data, applicationEntry{
				Application: d.Application,
				Action:      d.Action,
			})
		}

		ans.Application = &application{Entries: data}
	}

	return ans
}
