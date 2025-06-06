package errors

import (
	"encoding/xml"
	stderr "errors"
	"fmt"
	"strings"
)

var ObjectExists = stderr.New("object already exists")
var InvalidFilterError = stderr.New("filter is improperly formatted")
var NameNotSpecifiedError = stderr.New("name is not specified")
var NoLocationSpecifiedError = stderr.New("no location specified")
var RelativePositionWithRemoveEverythingElseError = stderr.New("cannot do relative positioning when removing all other rules")
var UnrecognizedOperatorError = stderr.New("unsupported filter operator")
var UnsupportedFilterTypeError = stderr.New("unsupported type for filtering")
var UuidNotSpecifiedError = stderr.New("uuid is not specified")
var UnsupportedMethodError = stderr.New("method is not supported")

// Panos is an error returned from PAN-OS.
//
// The error contains both the error message and the code returned from PAN-OS.
type Panos struct {
	Msg  string
	Code int
}

// Error returns the error message.
func (e Panos) Error() string {
	return e.Msg
}

// ObjectNotFound returns true if this is an object not found error.
func (e Panos) ObjectNotFound() bool {
	return e.Code == 7
}

func IsObjectNotFound(e error) bool {
	e2, ok := e.(Panos)
	if ok && e2.ObjectNotFound() {
		return true
	}

	return false
}

// ObjectNotFound returns an object not found error.
func ObjectNotFound() Panos {
	return Panos{
		Msg:  "Object not found",
		Code: 7,
	}
}

// Parse attempts to parse an error from the given XML response.
func Parse(body []byte) error {
	var e errorCheck

	_ = xml.Unmarshal(body, &e)
	if e.Failed() {
		return Panos{
			Msg:  e.Message(),
			Code: e.Code,
		}
	}

	return nil
}

type errorCheck struct {
	XMLName   xml.Name       `xml:"response"`
	Status    string         `xml:"status,attr"`
	Code      int            `xml:"code,attr"`
	Msg       *errorCheckMsg `xml:"msg"`
	ResultMsg *string        `xml:"result>msg"`
}

type errorCheckMsg struct {
	Line    []errLine `xml:"line"`
	Message string    `xml:",chardata"`
}

type errLine struct {
	Text string `xml:",cdata"`
}

func (e *errorCheck) Failed() bool {
	if e.Status == "failed" || e.Status == "error" {
		return true
	} else if e.Code == 0 || e.Code == 19 || e.Code == 20 {
		return false
	}

	return true
}

func (e *errorCheck) Message() string {
	if e.Msg != nil {
		if len(e.Msg.Line) > 0 {
			var b strings.Builder
			for i := range e.Msg.Line {
				if i != 0 {
					b.WriteString(" | ")
				}
				b.WriteString(strings.TrimSpace(e.Msg.Line[i].Text))
			}
			return b.String()
		}

		if e.Msg.Message != "" {
			return e.Msg.Message
		}
	}

	if e.ResultMsg != nil {
		return *e.ResultMsg
	}

	return e.CodeError()
}

func (e *errorCheck) CodeError() string {
	switch e.Code {
	case 1:
		return "Unknown command"
	case 2, 3, 4, 5, 11:
		return fmt.Sprintf("Internal error (%d) encountered", e.Code)
	case 6:
		return "Bad Xpath"
	case 7:
		return "Object not found"
	case 8:
		return "Object not unique"
	case 10:
		return "Reference count not zero"
	case 12:
		return "Invalid object"
	case 14:
		return "Operation not possible"
	case 15:
		return "Operation denied"
	case 16:
		return "Unauthorized"
	case 17:
		return "Invalid command"
	case 18:
		return "Malformed command"
	case 0, 19, 20:
		return ""
	case 22:
		return "Session timed out"
	default:
		return fmt.Sprintf("(%d) Unknown failure code, operation failed", e.Code)
	}
}

type InvalidXpathComponentError struct {
	Message string
}

func (o InvalidXpathComponentError) Error() string {
	return o.Message
}

func NewInvalidXpathComponentError(message string) *InvalidXpathComponentError {
	return &InvalidXpathComponentError{
		Message: message,
	}
}
