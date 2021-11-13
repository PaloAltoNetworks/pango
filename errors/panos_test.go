package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetSingularMissingObjectIsError(t *testing.T) {
	data := `<response status="success" code="7"><result/></response>`

	err := Parse([]byte(data))
	if err == nil {
		t.Errorf("Error is nil")
	} else {
		e2, ok := err.(Panos)
		if !ok {
			t.Errorf("Not a panos error")
		} else if !e2.ObjectNotFound() {
			t.Errorf("Not an object not found error")
		}
	}
}

func TestShowSingularMissingObjectIsError(t *testing.T) {
	data := `<response status="error"><msg><line>No such node</line></msg></response>`

	err := Parse([]byte(data))
	if err == nil {
		t.Errorf("Error is nil")
	} else {
		e2, ok := err.(Panos)
		if !ok {
			t.Errorf("Not a panos error")
		} else if e2.Msg != "No such node" {
			t.Errorf("Incorrect msg: %s", e2.Msg)
		}
	}
}

func TestMultilineErrorMessage(t *testing.T) {
	expected := "HTTP method must be GET"
	data := fmt.Sprintf(`<response status="error" code="12"><msg><line><![CDATA[ tf123456 -> server -> first server  is invalid. %s, Username/Password must not be empty when Tag Distribution is chosen]]></line><line><![CDATA[ tf123456 -> server is invalid]]></line></msg></response>`, expected)

	err := Parse([]byte(data))
	if err == nil {
		t.Errorf("Error is nil")
	} else {
		e2, ok := err.(Panos)
		if !ok {
			t.Errorf("Not a panos error")
		} else if !strings.Contains(e2.Msg, expected) {
			t.Errorf("Does not contain the expected substring: %s", e2.Msg)
		}
	}
}

func TestFailedExportErrorMessage(t *testing.T) {
	expected := `Parameter "format" is required while exporting certificate`
	data := `<response status = 'error' code = '400'><result><msg>Parameter &quot;format&quot; is required while exporting certificate</msg></result></response>`

	err := Parse([]byte(data))
	if err == nil {
		t.Errorf("Error is nil")
	} else {
		e2, ok := err.(Panos)
		if !ok {
			t.Errorf("Not a pnos error")
		} else if !strings.Contains(e2.Msg, expected) {
			t.Errorf("Does not contain the expected substring: %s", e2.Msg)
		}
	}
}
