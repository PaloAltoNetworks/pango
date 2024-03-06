package generic

import (
	"encoding/xml"
	"strings"
)

// Xml is a generic catch-all for parsing XML returned from PAN-OS.
type Xml struct {
	XMLName         xml.Name
	Name            *string `xml:"name,attr,omitempty"`
	Uuid            *string `xml:"uuid,attr,omitempty"`
	DetailedVersion *string `xml:"detail-version,attr,omitempty"`
	Text            []byte  `xml:",chardata"`
	Nodes           []Xml   `xml:",any"`

	// TrimmedText contains the trimmed value of Text.  Note that since this could
	// very well be trimming legitimate spacing that the text field would otherwise
	// contain, refering to this field for anything other than debugging purposes is
	// probably not a good idea.
	TrimmedText *string `xml:"-"`
}

func (x *Xml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local Xml
	var ans local
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}

	if len(ans.Text) != 0 {
		v := strings.TrimSpace(string(ans.Text))
		if v != "" {
			ans.TrimmedText = &v
		}
	}

	*x = Xml(ans)
	return nil
}
