package xmlapi

import (
	"encoding/xml"
)

type Generic struct {
	XMLName xml.Name
	Name    *string   `xml:"name,attr,omitempty"`
	Uuid    *string   `xml:"uuid,attr,omitempty"`
	Text    []byte    `xml:",innerxml"`
	Nodes   []Generic `xml:",any"`
}
