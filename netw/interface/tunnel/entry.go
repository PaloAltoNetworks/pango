package tunnel

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of
// a VLAN interface.
type Entry struct {
	Name              string
	Comment           string
	NetflowProfile    string
	StaticIps         []string // ordered
	ManagementProfile string
	Mtu               int

	raw map[string]string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Comment = s.Comment
	o.NetflowProfile = s.NetflowProfile
	o.StaticIps = s.StaticIps
	o.ManagementProfile = s.ManagementProfile
	o.Mtu = s.Mtu
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, string, interface{}) {
	_, fn := versioning(v)
	return o.Name, o.Name, fn(o)
}

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
	}

	ans.raw = make(map[string]string)
	if o.Ipv6 != nil {
		ans.raw["ipv6"] = util.CleanRawXml(o.Ipv6.Text)
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
	Mtu               int             `xml:"mtu,omitempty"`
	ManagementProfile string          `xml:"interface-management-profile,omitempty"`

	Ipv6 *util.RawXml `xml:"ipv6"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:              e.Name,
		Comment:           e.Comment,
		NetflowProfile:    e.NetflowProfile,
		StaticIps:         util.StrToEnt(e.StaticIps),
		Mtu:               e.Mtu,
		ManagementProfile: e.ManagementProfile,
	}

	if text, ok := e.raw["ipv6"]; ok {
		ans.Ipv6 = &util.RawXml{text}
	}

	return ans
}
