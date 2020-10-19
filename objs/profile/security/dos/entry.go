package dos

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// DOS protection security profile.
type Entry struct {
	Name        string
	Description string

	EnableSyn        bool
	SynAction        string
	SynAlarmRate     int
	SynActivateRate  int
	SynMaxRate       int
	SynBlockDuration int

	EnableUdp        bool
	UdpAlarmRate     int
	UdpActivateRate  int
	UdpMaxRate       int
	UdpBlockDuration int

	EnableIcmp        bool
	IcmpAlarmRate     int
	IcmpActivateRate  int
	IcmpMaxRate       int
	IcmpBlockDuration int

	EnableIcmpv6        bool
	Icmpv6AlarmRate     int
	Icmpv6ActivateRate  int
	Icmpv6MaxRate       int
	Icmpv6BlockDuration int

	EnableOther        bool
	OtherAlarmRate     int
	OtherActivateRate  int
	OtherMaxRate       int
	OtherBlockDuration int

	EnableSessionsProtections bool
	MaxConcurrentSessions     int
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.EnableSyn = s.EnableSyn
	o.SynAction = s.SynAction
	o.SynAlarmRate = s.SynAlarmRate
	o.SynActivateRate = s.SynActivateRate
	o.SynMaxRate = s.SynMaxRate
	o.SynBlockDuration = s.SynBlockDuration
	o.EnableUdp = s.EnableUdp
	o.UdpAlarmRate = s.UdpAlarmRate
	o.UdpActivateRate = s.UdpActivateRate
	o.UdpMaxRate = s.UdpMaxRate
	o.UdpBlockDuration = s.UdpBlockDuration
	o.EnableIcmp = s.EnableIcmp
	o.IcmpAlarmRate = s.IcmpAlarmRate
	o.IcmpActivateRate = s.IcmpActivateRate
	o.IcmpMaxRate = s.IcmpMaxRate
	o.IcmpBlockDuration = s.IcmpBlockDuration
	o.EnableIcmpv6 = s.EnableIcmpv6
	o.Icmpv6AlarmRate = s.Icmpv6AlarmRate
	o.Icmpv6ActivateRate = s.Icmpv6ActivateRate
	o.Icmpv6MaxRate = s.Icmpv6MaxRate
	o.Icmpv6BlockDuration = s.Icmpv6BlockDuration
	o.EnableOther = s.EnableOther
	o.OtherAlarmRate = s.OtherAlarmRate
	o.OtherActivateRate = s.OtherActivateRate
	o.OtherMaxRate = s.OtherMaxRate
	o.OtherBlockDuration = s.OtherBlockDuration
	o.EnableSessionsProtections = s.EnableSessionsProtections
	o.MaxConcurrentSessions = s.MaxConcurrentSessions
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
		Name:        o.Name,
		Description: o.Description,
	}

	if o.Flood != nil {
		if o.Flood.Syn != nil {
			ans.EnableSyn = util.AsBool(o.Flood.Syn.EnableSyn)
			if o.Flood.Syn.Red != nil {
				ans.SynAction = SynActionRed
				ans.SynAlarmRate = o.Flood.Syn.Red.AlarmRate
				ans.SynActivateRate = o.Flood.Syn.Red.ActivateRate
				ans.SynMaxRate = o.Flood.Syn.Red.MaxRate
				if o.Flood.Syn.Red.Block != nil {
					ans.SynBlockDuration = o.Flood.Syn.Red.Block.BlockDuration
				}
			} else if o.Flood.Syn.Cookies != nil {
				ans.SynAction = SynActionCookies
				ans.SynAlarmRate = o.Flood.Syn.Cookies.AlarmRate
				ans.SynActivateRate = o.Flood.Syn.Cookies.ActivateRate
				ans.SynMaxRate = o.Flood.Syn.Cookies.MaxRate
				if o.Flood.Syn.Cookies.Block != nil {
					ans.SynBlockDuration = o.Flood.Syn.Cookies.Block.BlockDuration
				}
			}
		}

		if o.Flood.Udp != nil {
			ans.EnableUdp = util.AsBool(o.Flood.Udp.Enable)
			if o.Flood.Udp.Red != nil {
				ans.UdpAlarmRate = o.Flood.Udp.Red.AlarmRate
				ans.UdpActivateRate = o.Flood.Udp.Red.ActivateRate
				ans.UdpMaxRate = o.Flood.Udp.Red.MaxRate
				if o.Flood.Udp.Red.Block != nil {
					ans.UdpBlockDuration = o.Flood.Udp.Red.Block.BlockDuration
				}
			}
		}

		if o.Flood.Icmp != nil {
			ans.EnableIcmp = util.AsBool(o.Flood.Icmp.Enable)
			if o.Flood.Icmp.Red != nil {
				ans.IcmpAlarmRate = o.Flood.Icmp.Red.AlarmRate
				ans.IcmpActivateRate = o.Flood.Icmp.Red.ActivateRate
				ans.IcmpMaxRate = o.Flood.Icmp.Red.MaxRate
				if o.Flood.Icmp.Red.Block != nil {
					ans.IcmpBlockDuration = o.Flood.Icmp.Red.Block.BlockDuration
				}
			}
		}

		if o.Flood.Icmpv6 != nil {
			ans.EnableIcmpv6 = util.AsBool(o.Flood.Icmpv6.Enable)
			if o.Flood.Icmpv6.Red != nil {
				ans.Icmpv6AlarmRate = o.Flood.Icmpv6.Red.AlarmRate
				ans.Icmpv6ActivateRate = o.Flood.Icmpv6.Red.ActivateRate
				ans.Icmpv6MaxRate = o.Flood.Icmpv6.Red.MaxRate
				if o.Flood.Icmpv6.Red.Block != nil {
					ans.Icmpv6BlockDuration = o.Flood.Icmpv6.Red.Block.BlockDuration
				}
			}
		}

		if o.Flood.Other != nil {
			ans.EnableOther = util.AsBool(o.Flood.Other.Enable)
			if o.Flood.Other.Red != nil {
				ans.OtherAlarmRate = o.Flood.Other.Red.AlarmRate
				ans.OtherActivateRate = o.Flood.Other.Red.ActivateRate
				ans.OtherMaxRate = o.Flood.Other.Red.MaxRate
				if o.Flood.Other.Red.Block != nil {
					ans.OtherBlockDuration = o.Flood.Other.Red.Block.BlockDuration
				}
			}
		}
	}

	if o.Resource != nil {
		if o.Resource.Sess != nil {
			ans.EnableSessionsProtections = util.AsBool(o.Resource.Sess.EnableSessionsProtections)
			ans.MaxConcurrentSessions = o.Resource.Sess.MaxConcurrentSessions
		}
	}

	return ans
}

type entry_v1 struct {
	XMLName     xml.Name  `xml:"entry"`
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description,omitempty"`
	Flood       *flood    `xml:"flood"`
	Resource    *resource `xml:"resource"`
}

type flood struct {
	Syn    *syn    `xml:"tcp-syn"`
	Udp    *common `xml:"udp"`
	Icmp   *common `xml:"icmp"`
	Icmpv6 *common `xml:"icmpv6"`
	Other  *common `xml:"other-ip"`
}

type syn struct {
	EnableSyn string   `xml:"enable"`
	Red       *details `xml:"red"`
	Cookies   *details `xml:"syn-cookies"`
}

type details struct {
	AlarmRate    int    `xml:"alarm-rate"`
	ActivateRate int    `xml:"activate-rate"`
	MaxRate      int    `xml:"maximal-rate"`
	Block        *block `xml:"block"`
}

type block struct {
	BlockDuration int `xml:"duration,omitempty"`
}

type common struct {
	Enable string   `xml:"enable"`
	Red    *details `xml:"red"`
}

type resource struct {
	Sess *sess `xml:"sessions"`
}

type sess struct {
	EnableSessionsProtections string `xml:"enabled"`
	MaxConcurrentSessions     int    `xml:"max-concurrent-limit,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:        e.Name,
		Description: e.Description,
	}

	incSyn := e.SynAction != "" || e.SynAlarmRate != 0 || e.SynActivateRate != 0 || e.SynMaxRate != 0 || e.SynBlockDuration != 0
	incUdp := e.UdpAlarmRate != 0 || e.UdpActivateRate != 0 || e.UdpMaxRate != 0 || e.UdpBlockDuration != 0
	incIcmp := e.IcmpAlarmRate != 0 || e.IcmpActivateRate != 0 || e.IcmpMaxRate != 0 || e.IcmpBlockDuration != 0
	incIcmpv6 := e.Icmpv6AlarmRate != 0 || e.Icmpv6ActivateRate != 0 || e.Icmpv6MaxRate != 0 || e.Icmpv6BlockDuration != 0
	incOther := e.OtherAlarmRate != 0 || e.OtherActivateRate != 0 || e.OtherMaxRate != 0 || e.OtherBlockDuration != 0

	if e.EnableSyn || incSyn || e.EnableUdp || incUdp || e.EnableIcmp || incIcmp || e.EnableIcmpv6 || incIcmpv6 || e.EnableOther || incOther {
		ans.Flood = &flood{}
		if e.EnableSyn || incSyn {
			ans.Flood.Syn = &syn{
				EnableSyn: util.YesNo(e.EnableSyn),
			}

			if incSyn {
				switch e.SynAction {
				case SynActionRed:
					ans.Flood.Syn.Red = &details{
						AlarmRate:    e.SynAlarmRate,
						ActivateRate: e.SynActivateRate,
						MaxRate:      e.SynMaxRate,
					}
					if e.SynBlockDuration != 0 {
						ans.Flood.Syn.Red.Block = &block{
							BlockDuration: e.SynBlockDuration,
						}
					}
				case SynActionCookies:
					ans.Flood.Syn.Cookies = &details{
						AlarmRate:    e.SynAlarmRate,
						ActivateRate: e.SynActivateRate,
						MaxRate:      e.SynMaxRate,
					}
					if e.SynBlockDuration != 0 {
						ans.Flood.Syn.Cookies.Block = &block{
							BlockDuration: e.SynBlockDuration,
						}
					}
				}
			}
		}

		if e.EnableUdp || incUdp {
			ans.Flood.Udp = &common{
				Enable: util.YesNo(e.EnableUdp),
			}

			if incUdp {
				ans.Flood.Udp.Red = &details{
					AlarmRate:    e.UdpAlarmRate,
					ActivateRate: e.UdpActivateRate,
					MaxRate:      e.UdpMaxRate,
				}
				if e.UdpBlockDuration != 0 {
					ans.Flood.Udp.Red.Block = &block{
						BlockDuration: e.UdpBlockDuration,
					}
				}
			}
		}

		if e.EnableIcmp || incIcmp {
			ans.Flood.Icmp = &common{
				Enable: util.YesNo(e.EnableIcmp),
			}

			if incIcmp {
				ans.Flood.Icmp.Red = &details{
					AlarmRate:    e.IcmpAlarmRate,
					ActivateRate: e.IcmpActivateRate,
					MaxRate:      e.IcmpMaxRate,
				}
				if e.IcmpBlockDuration != 0 {
					ans.Flood.Icmp.Red.Block = &block{
						BlockDuration: e.IcmpBlockDuration,
					}
				}
			}
		}

		if e.EnableIcmpv6 || incIcmpv6 {
			ans.Flood.Icmpv6 = &common{
				Enable: util.YesNo(e.EnableIcmpv6),
			}

			if incIcmpv6 {
				ans.Flood.Icmpv6.Red = &details{
					AlarmRate:    e.Icmpv6AlarmRate,
					ActivateRate: e.Icmpv6ActivateRate,
					MaxRate:      e.Icmpv6MaxRate,
				}
				if e.Icmpv6BlockDuration != 0 {
					ans.Flood.Icmpv6.Red.Block = &block{
						BlockDuration: e.Icmpv6BlockDuration,
					}
				}
			}
		}

		if e.EnableOther || incOther {
			ans.Flood.Other = &common{
				Enable: util.YesNo(e.EnableOther),
			}

			if incOther {
				ans.Flood.Other.Red = &details{
					AlarmRate:    e.OtherAlarmRate,
					ActivateRate: e.OtherActivateRate,
					MaxRate:      e.OtherMaxRate,
				}
				if e.OtherBlockDuration != 0 {
					ans.Flood.Other.Red.Block = &block{
						BlockDuration: e.OtherBlockDuration,
					}
				}
			}
		}
	}

	if e.EnableSessionsProtections || e.MaxConcurrentSessions != 0 {
		ans.Resource = &resource{
			Sess: &sess{
				EnableSessionsProtections: util.YesNo(e.EnableSessionsProtections),
				MaxConcurrentSessions:     e.MaxConcurrentSessions,
			},
		}
	}

	return ans
}
