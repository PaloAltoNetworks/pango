package file

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// file blocking security profile.
//
// PAN-OS 8.0+.
type Entry struct {
	Name        string
	Description string
	Rules       []Rule
}

type Rule struct {
	Name         string
	Applications []string
	FileTypes    []string
	Direction    string
	Action       string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.Rules = s.Rules
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

	if o.Rules != nil {
		rules := make([]Rule, 0, len(o.Rules.Entries))
		for _, v := range o.Rules.Entries {
			r := Rule{
				Name:         v.Name,
				Applications: util.MemToStr(v.Applications),
				FileTypes:    util.MemToStr(v.FileTypes),
				Direction:    v.Direction,
			}

			if v.Action != nil {
				if v.Action.Alert != nil {
					r.Action = ActionAlert
				} else if v.Action.Block != nil {
					r.Action = ActionBlock
				} else if v.Action.Continue != nil {
					r.Action = ActionContinue
				} else if v.Action.Forward != nil {
					r.Action = ActionForward
				} else if v.Action.ContinueAndForward != nil {
					r.Action = ActionContinueAndForward
				}
			}
			rules = append(rules, r)
		}
		ans.Rules = rules
	}

	return ans
}

type entry_v1 struct {
	XMLName     xml.Name  `xml:"entry"`
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description,omitempty"`
	Rules       *rules_v1 `xml:"rules"`
}

type rules_v1 struct {
	Entries []rule_v1 `xml:"entry"`
}

type rule_v1 struct {
	Name         string           `xml:"name,attr"`
	Applications *util.MemberType `xml:"application"`
	FileTypes    *util.MemberType `xml:"file-type"`
	Direction    string           `xml:"direction,omitempty"`
	Action       *action_v1       `xml:"action"`
}

type action_v1 struct {
	Alert              *string `xml:"alert"`
	Block              *string `xml:"block"`
	Continue           *string `xml:"continue"`
	Forward            *string `xml:"forward"`
	ContinueAndForward *string `xml:"continue-and-forward"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:        e.Name,
		Description: e.Description,
	}

	if len(e.Rules) > 0 {
		rules := make([]rule_v1, 0, len(e.Rules))
		for _, er := range e.Rules {
			r := rule_v1{
				Name:         er.Name,
				Applications: util.StrToMem(er.Applications),
				FileTypes:    util.StrToMem(er.FileTypes),
				Direction:    er.Direction,
			}

			s := ""
			switch er.Action {
			case ActionAlert:
				r.Action = &action_v1{
					Alert: &s,
				}
			case ActionBlock:
				r.Action = &action_v1{
					Block: &s,
				}
			case ActionContinue:
				r.Action = &action_v1{
					Continue: &s,
				}
			case ActionForward:
				r.Action = &action_v1{
					Forward: &s,
				}
			case ActionContinueAndForward:
				r.Action = &action_v1{
					ContinueAndForward: &s,
				}
			}
			rules = append(rules, r)
		}
		ans.Rules = &rules_v1{Entries: rules}
	}

	return ans
}

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

	if o.Rules != nil {
		rules := make([]Rule, 0, len(o.Rules.Entries))
		for _, v := range o.Rules.Entries {
			r := Rule{
				Name:         v.Name,
				Applications: util.MemToStr(v.Applications),
				FileTypes:    util.MemToStr(v.FileTypes),
				Direction:    v.Direction,
			}

			if v.Action != nil {
				if v.Action.Alert != nil {
					r.Action = ActionAlert
				} else if v.Action.Block != nil {
					r.Action = ActionBlock
				} else if v.Action.Continue != nil {
					r.Action = ActionContinue
				}
			}
			rules = append(rules, r)
		}
		ans.Rules = rules
	}
	return ans
}

type entry_v2 struct {
	XMLName     xml.Name  `xml:"entry"`
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description,omitempty"`
	Rules       *rules_v2 `xml:"rules"`
}

type rules_v2 struct {
	Entries []rule_v2 `xml:"entry"`
}

type rule_v2 struct {
	Name         string           `xml:"name,attr"`
	Applications *util.MemberType `xml:"application"`
	FileTypes    *util.MemberType `xml:"file-type"`
	Direction    string           `xml:"direction,omitempty"`
	Action       *action_v2       `xml:"action"`
}

type action_v2 struct {
	Alert    *string `xml:"alert"`
	Block    *string `xml:"block"`
	Continue *string `xml:"continue"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:        e.Name,
		Description: e.Description,
	}

	if len(e.Rules) > 0 {
		rules := make([]rule_v2, 0, len(e.Rules))
		for _, er := range e.Rules {
			r := rule_v2{
				Name:         er.Name,
				Applications: util.StrToMem(er.Applications),
				FileTypes:    util.StrToMem(er.FileTypes),
				Direction:    er.Direction,
			}

			s := ""
			switch er.Action {
			case ActionAlert:
				r.Action = &action_v2{
					Alert: &s,
				}
			case ActionBlock:
				r.Action = &action_v2{
					Block: &s,
				}
			case ActionContinue:
				r.Action = &action_v2{
					Continue: &s,
				}
			}
			rules = append(rules, r)
		}
		ans.Rules = &rules_v2{Entries: rules}
	}

	return ans
}
