package data

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a
// custom data pattern object.
type Entry struct {
	Name               string
	Description        string
	Type               string
	PredefinedPatterns []PredefinedPattern
	Regexes            []Regex
	FileProperties     []FileProperty
}

type PredefinedPattern struct {
	Name      string
	FileTypes []string
}

type Regex struct {
	Name      string
	FileTypes []string
	Regex     string
}

type FileProperty struct {
	Name          string
	FileType      string
	FileProperty  string
	PropertyValue string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	o.Type = s.Type
	o.PredefinedPatterns = s.PredefinedPatterns
	o.Regexes = s.Regexes
	o.FileProperties = s.FileProperties
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

	switch {
	case o.Type.Predefined != nil:
		ans.Type = TypePredefined
		if o.Type.Predefined.Pattern != nil {
			data := make([]PredefinedPattern, 0, len(o.Type.Predefined.Pattern.Entries))
			for _, d := range o.Type.Predefined.Pattern.Entries {
				data = append(data, PredefinedPattern{
					Name:      d.Name,
					FileTypes: util.MemToStr(d.FileTypes),
				})
			}

			ans.PredefinedPatterns = data
		}
	case o.Type.Regex != nil:
		ans.Type = TypeRegex
		if o.Type.Regex.Pattern != nil {
			data := make([]Regex, 0, len(o.Type.Regex.Pattern.Entries))
			for _, d := range o.Type.Regex.Pattern.Entries {
				data = append(data, Regex{
					Name:      d.Name,
					FileTypes: util.MemToStr(d.FileTypes),
					Regex:     d.Regex,
				})
			}

			ans.Regexes = data
		}
	case o.Type.FileProperties != nil:
		ans.Type = TypeFileProperties
		if o.Type.FileProperties.Pattern != nil {
			data := make([]FileProperty, 0, len(o.Type.FileProperties.Pattern.Entries))
			for _, d := range o.Type.FileProperties.Pattern.Entries {
				data = append(data, FileProperty{
					Name:          d.Name,
					FileType:      d.FileType,
					FileProperty:  d.FileProperty,
					PropertyValue: d.PropertyValue,
				})
			}

			ans.FileProperties = data
		}
	}

	return ans
}

type entry_v1 struct {
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,omitempty"`
	Type        dpType   `xml:"pattern-type"`
}

type dpType struct {
	Predefined     *predefined `xml:"predefined"`
	Regex          *regex      `xml:"regex"`
	FileProperties *fp         `xml:"file-properties"`
}

type predefined struct {
	Pattern *predefPat `xml:"pattern"`
}

type predefPat struct {
	Entries []predefEntry `xml:"entry"`
}

type predefEntry struct {
	Name      string           `xml:"name,attr"`
	FileTypes *util.MemberType `xml:"file-type"`
}

type regex struct {
	Pattern *regexPat `xml:"pattern"`
}

type regexPat struct {
	Entries []regexEntry `xml:"entry"`
}

type regexEntry struct {
	Name      string           `xml:"name,attr"`
	FileTypes *util.MemberType `xml:"file-type"`
	Regex     string           `xml:"regex"`
}

type fp struct {
	Pattern *fpPat `xml:"pattern"`
}

type fpPat struct {
	Entries []fpEntry `xml:"entry"`
}

type fpEntry struct {
	Name          string `xml:"name,attr"`
	FileType      string `xml:"file-type"`
	FileProperty  string `xml:"file-property"`
	PropertyValue string `xml:"property-value"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:        e.Name,
		Description: e.Description,
	}

	switch e.Type {
	case TypePredefined:
		ans.Type.Predefined = &predefined{}
		if len(e.PredefinedPatterns) > 0 {
			data := make([]predefEntry, 0, len(e.PredefinedPatterns))
			for _, d := range e.PredefinedPatterns {
				data = append(data, predefEntry{
					Name:      d.Name,
					FileTypes: util.StrToMem(d.FileTypes),
				})
			}

			ans.Type.Predefined.Pattern = &predefPat{Entries: data}
		}
	case TypeRegex:
		ans.Type.Regex = &regex{}
		if len(e.Regexes) > 0 {
			data := make([]regexEntry, 0, len(e.Regexes))
			for _, d := range e.Regexes {
				data = append(data, regexEntry{
					Name:      d.Name,
					FileTypes: util.StrToMem(d.FileTypes),
					Regex:     d.Regex,
				})
			}

			ans.Type.Regex.Pattern = &regexPat{Entries: data}
		}
	case TypeFileProperties:
		ans.Type.FileProperties = &fp{}
		if len(e.FileProperties) > 0 {
			data := make([]fpEntry, 0, len(e.FileProperties))
			for _, d := range e.FileProperties {
				data = append(data, fpEntry{
					Name:          d.Name,
					FileType:      d.FileType,
					FileProperty:  d.FileProperty,
					PropertyValue: d.PropertyValue,
				})
			}

			ans.Type.FileProperties.Pattern = &fpPat{Entries: data}
		}
	}

	return ans
}
