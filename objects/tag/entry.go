package tag

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
	Suffix = []string{"tag"}
)

type Entry struct {
	Name     string
	Color    *string
	Comments *string

	Misc map[string][]generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`
	Color    *string  `xml:"color,omitempty"`
	Comments *string  `xml:"comments,omitempty"`

	Misc []generic.Xml `xml:",any"`
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

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyEntry, &entryXmlContainer{}, nil
}

func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}

	entry.Name = o.Name
	entry.Color = o.Color
	entry.Comments = o.Comments

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}
func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		entry.Color = o.Color
		entry.Comments = o.Comments

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !util.StringsMatch(a.Color, b.Color) {
		return false
	}
	if !util.StringsMatch(a.Comments, b.Comments) {
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
