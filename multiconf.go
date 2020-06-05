package pango

import (
	"encoding/xml"
	"strings"

	"github.com/PaloAltoNetworks/pango/util"
)

type MultiConfigure struct {
	XMLName xml.Name                `xml:"multi-configure-request"`
	Reqs    []MultiConfigureRequest `xml:"request"`
}

func (m *MultiConfigure) IncrementalIds() {
	for i := range m.Reqs {
		m.Reqs[i].Id = i + 1
	}
}

type MultiConfigureRequest struct {
	Id      int    `xml:"id,attr,omitempty"`
	Command string `xml:"cmd,attr"`
	Xpath   string `xml:"obj,attr"`
	Data    interface{}
}

type MultiConfigureResponse struct {
	XMLName xml.Name                     `xml:"response"`
	Status  string                       `xml:"status,attr"`
	Code    int                          `xml:"code,attr"`
	Results []MultiConfigResponseElement `xml:"response"`
}

func (m *MultiConfigureResponse) Ok() bool {
	return m.Status == "success"
}

func (m *MultiConfigureResponse) Error() string {
	r := m.Results[len(m.Results)-1]
	if r.Ok() {
		return ""
	}

	return r.Message()
}

type MultiConfigResponseElement struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    int      `xml:"code,attr"`
	Id      int      `xml:"id,attr,omitempty"`
	Msg     McreMsg  `xml:"msg"`
}

type McreMsg struct {
	Line    *util.CdataText `xml:"line"`
	Message string          `xml:",chardata"`
}

func (m *MultiConfigResponseElement) Ok() bool {
	return m.Status == "success"
}

func (m *MultiConfigResponseElement) Message() string {
	if m.Msg.Line != nil {
		return strings.TrimSpace(m.Msg.Line.Text)
	}

	return m.Msg.Message
}
