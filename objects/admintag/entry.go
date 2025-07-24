package admintag

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
	suffix = []string{"tag", "$name"}
)

type Entry struct {
	Name            string
	Color           *string
	Comments        *string
	DisableOverride *string
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
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
	Color           *string       `xml:"color,omitempty"`
	Comments        *string       `xml:"comments,omitempty"`
	DisableOverride *string       `xml:"disable-override,omitempty"`
	Misc            []generic.Xml `xml:",any"`
	MiscAttributes  []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Color = s.Color
	o.Comments = s.Comments
	o.DisableOverride = s.DisableOverride
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {

	result := &Entry{
		Name:            o.Name,
		Color:           o.Color,
		Comments:        o.Comments,
		DisableOverride: o.DisableOverride,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}

const (
	ColorAzureBlue     = "color24"
	ColorBlack         = "color14"
	ColorBlueGray      = "color12"
	ColorBlueViolet    = "color30"
	ColorBrown         = "color16"
	ColorBurntSienna   = "color41"
	ColorCeruleanBlue  = "color25"
	ColorChestnut      = "color42"
	ColorCobaltBlue    = "color28"
	ColorCopper        = "color4"
	ColorCyan          = "color9"
	ColorForestGreen   = "color22"
	ColorGold          = "color15"
	ColorGray          = "color7"
	ColorGreen         = "color2"
	ColorLavender      = "color33"
	ColorLightGray     = "color10"
	ColorLightGreen    = "color8"
	ColorLime          = "color13"
	ColorMagenta       = "color38"
	ColorMahogany      = "color40"
	ColorMaroon        = "color19"
	ColorMediumBlue    = "color27"
	ColorMediumRose    = "color32"
	ColorMediumViolet  = "color31"
	ColorMidnightBlue  = "color26"
	ColorOlive         = "color17"
	ColorOrange        = "color5"
	ColorOrchid        = "color34"
	ColorPeach         = "color36"
	ColorPurple        = "color6"
	ColorRed           = "color1"
	ColorRedViolet     = "color39"
	ColorRedOrange     = "color20"
	ColorSalmon        = "color37"
	ColorThistle       = "color35"
	ColorTurquoiseBlue = "color23"
	ColorVioletBlue    = "color29"
	ColorYellow        = "color3"
	ColorYellowOrange  = "color21"
)

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "color" || v == "Color" {
		return e.Color, nil
	}
	if v == "comments" || v == "Comments" {
		return e.Comments, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
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
	if !util.StringsMatch(o.Color, other.Color) {
		return false
	}
	if !util.StringsMatch(o.Comments, other.Comments) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
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
