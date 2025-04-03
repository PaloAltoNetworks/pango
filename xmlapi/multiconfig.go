package xmlapi

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

func NewChunkedMultiConfig(capacity int, batchSize int) *MultiConfig {
	return &MultiConfig{
		BatchSize:  batchSize,
		Operations: make([]MultiConfigOperation, 0, capacity),
	}
}

// Returns a new struct for performing multi-config operations with.
func NewMultiConfig(capacity int) *MultiConfig {
	return &MultiConfig{
		Operations: make([]MultiConfigOperation, 0, capacity),
	}
}

// MultiConfig contains individual operations of a larger multi-config API call.
type MultiConfig struct {
	XMLName    xml.Name `xml:"multi-configure-request"`
	BatchSize  int      `xml:"-"`
	Operations []MultiConfigOperation
}

// IncrementalIds assigns incremental IDs to a multi-config call.  Any individual
// operation that already has an ID will be left as-is.
func (m *MultiConfig) IncrementalIds() {
	for i := range m.Operations {
		if m.Operations[i].Id == "" {
			m.Operations[i].Id = fmt.Sprintf("%d", i+1)
		}
	}
}

// Add adds a Config operation to the multi-config request.
func (m *MultiConfig) Add(c *Config) {
	if c == nil {
		return
	}

	o := MultiConfigOperation{
		XMLName:     xml.Name{Local: c.Action},
		Xpath:       c.Xpath,
		Where:       c.Where,
		Destination: c.Destination,
		NewName:     c.NewName,
	}

	if c.Element != nil {
		o.Data = c.Element
	}

	m.Operations = append(m.Operations, o)
}

// MultiConfigOperation is an individual operation in a MultiConfigure instance.
type MultiConfigOperation struct {
	XMLName xml.Name
	Id      string `xml:"id,attr,omitempty"`

	Xpath string `xml:"xpath,attr"`
	Data  interface{}

	// action=move.
	Where       string `xml:"where,attr,omitempty"`
	Destination string `xml:"dst,attr,omitempty"`

	// action=rename.
	NewName string `xml:"newname,attr,omitempty"`
}

// NewMultiConfigResponse creates a multi-config response from the given text.
func NewMultiConfigResponse(text []byte) (*MultiConfigResponse, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("no data received in the multi-config response")
	}

	ans := MultiConfigResponse{raw: text}
	err := xml.Unmarshal(text, &ans)

	return &ans, err
}

// MultiConfigResponse is a struct to handle the response from multi-config
// commands.
type MultiConfigResponse struct {
	XMLName xml.Name                     `xml:"response"`
	Status  string                       `xml:"status,attr"`
	Code    int                          `xml:"code,attr"`
	Results []MultiConfigResponseElement `xml:"response"`

	raw []byte `xml:"-"`
}

// Ok returns if there was an error or not.
func (m *MultiConfigResponse) Ok() bool {
	return m.Status == "success"
}

// Error returns the error if there was one.
func (m *MultiConfigResponse) Error() string {
	if len(m.Results) == 0 {
		if err := errors.Parse(m.raw); err != nil {
			return err.Error()
		}
		return "unknown multi-config error format"
	}

	r := m.Results[len(m.Results)-1]
	if r.Ok() {
		return ""
	}

	return r.Message()
}

// MultiConfigResponseElement is a single response from a multi-config request.
type MultiConfigResponseElement struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    int      `xml:"code,attr"`
	Id      string   `xml:"id,attr,omitempty"`
	Msg     McreMsg  `xml:"msg"`
}

type McreMsg struct {
	Line    []CdataText `xml:"line"`
	Message string      `xml:",chardata"`
}

func (m *MultiConfigResponseElement) Ok() bool {
	return m.Status == "success"
}

func (m *MultiConfigResponseElement) Message() string {
	if len(m.Msg.Line) > 0 {
		var b strings.Builder
		for i := range m.Msg.Line {
			if i != 0 {
				b.WriteString(" | ")
			}
			b.WriteString(strings.TrimSpace(m.Msg.Line[i].Text))
		}
		return b.String()
	}

	return m.Msg.Message
}

type ChunkedMultiConfigResponse struct {
	Data                []byte
	HttpResponse        *http.Response
	MultiConfigResponse *MultiConfigResponse
}
