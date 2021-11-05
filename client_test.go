package pango

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/PaloAltoNetworks/pango/commit"
	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/testdata"
)

func init() {
	log.SetFlags(0)
}

// rl restores the logger output to the default log location.
func rl() {
	log.SetOutput(os.Stderr)
}

func TestClientStringerHidesPasswords(t *testing.T) {
	c := &Client{
		Hostname: "foo",
		Username: "admin",
		Password: "secret",
	}

	if strings.Index(c.String(), "secret") != -1 {
		t.Fail()
	}
}

func TestClientStringerHidesApiKey(t *testing.T) {
	c := &Client{
		Hostname: "foo",
		Username: "admin",
		ApiKey:   "secret",
	}

	if strings.Index(c.String(), "secret") != -1 {
		t.Fail()
	}
}

func TestLogActionDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogQuiet,
	}
	c.LogAction("no")
	if buf.String() != "" {
		t.Fail()
	}
}

func TestLogActionEnabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogAction,
	}
	c.LogAction("yes")
	s := buf.String()
	if s != "yes\n" {
		t.Fail()
	}
}

func TestLogQueryDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogQuiet,
	}
	c.LogQuery("no")
	if buf.String() != "" {
		t.Fail()
	}
}

func TestLogQueryEnabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogQuery,
	}
	c.LogQuery("yes")
	s := buf.String()
	if s != "yes\n" {
		t.Fail()
	}
}

func TestLogOpDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogQuiet,
	}
	c.LogOp("no")
	if buf.String() != "" {
		t.Fail()
	}
}

func TestLogOpEnabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogOp,
	}
	c.LogOp("yes")
	s := buf.String()
	if s != "yes\n" {
		t.Fail()
	}
}

func TestLogUidDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogQuiet,
	}
	c.LogUid("no")
	if buf.String() != "" {
		t.Fail()
	}
}

func TestLogUidEnabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Logging: LogUid,
	}
	c.LogUid("yes")
	s := buf.String()
	if s != "yes\n" {
		t.Fail()
	}
}

func TestRetrieveApiKey(t *testing.T) {
	c := &Client{}
	c.rb = [][]byte{
		[]byte(testdata.ApiKeyXml),
	}
	if err := c.Initialize(); err != nil {
		t.Errorf("Initialize failed: %s", err)
		return
	}

	err := c.RetrieveApiKey()
	if err != nil {
		t.Errorf("Failed to retrieve api key: %s", err)
	} else if c.ApiKey != "secret" {
		t.Fail()
	}
}

func TestAsStringWithString(t *testing.T) {
	expected := "foobar"
	val, err := asString(expected, true)

	if err != nil {
		t.Errorf("Failed asString: %s", err)
	} else if val != expected {
		t.Errorf("Expected:%s got:%s", expected, val)
	}
}

type StringerCheck struct{}

func (o StringerCheck) String() string { return "hello" }

func TestAsStringWithStringer(t *testing.T) {
	x := StringerCheck{}
	expected := x.String()

	val, err := asString(x, true)
	if err != nil {
		t.Errorf("Failed asString: %s", err)
	} else if val != expected {
		t.Errorf("Expected:%s got:%s", expected, val)
	}
}

func TestAsStringWithElementer(t *testing.T) {
	x := commit.FirewallCommit{
		Description: "hello",
	}

	expected, err := xml.Marshal(x.Element())
	if err != nil {
		t.Errorf("Failed to marshal firewall commit: %s", err)
	}

	val, err := asString(x, true)
	if err != nil {
		t.Errorf("Failed asString: %s", err)
	} else if val != string(expected) {
		t.Errorf("Expected:%s got:%s", expected, val)
	}
}

func TestAsStringWithNil(t *testing.T) {
	if _, err := asString(nil, true); err == nil {
		t.Errorf("asString() returned no error on nil input")
	}
}

func TestLogSendBasic(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		ApiKey:  "secret",
		Logging: LogSend,
	}
	values := url.Values{
		"type":     []string{"keygen"},
		"user":     []string{"foo"},
		"password": []string{"bar"},
		"key":      []string{c.ApiKey},
	}
	c.logSend(values)

	s := buf.String()
	if s == "" {
		t.Fail()
	} else if strings.Contains(s, c.ApiKey) {
		t.Errorf("API key was not masked")
	} else {
		for _, k := range []string{"type", "keygen", "user", "foo", "password", "bar"} {
			if !strings.Contains(s, k) {
				t.Errorf("%s is not present in basic logsend", k)
			}
		}
	}
}

func TestLogSendOsxCurlWithoutPersonalDataWithoutElement(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Hostname: "127.0.0.1",
		ApiKey:   "secret",
		Protocol: "https",
		Logging:  LogOsxCurl,
	}
	values := url.Values{
		"type":   []string{"config"},
		"action": []string{"delete"},
		"xpath":  []string{"mypathhere"},
		"key":    []string{"APIKEY"},
	}
	expected := fmt.Sprintf("curl -k 'https://HOST/api?%s'\n", values.Encode())
	values.Set("key", c.ApiKey)
	c.logSend(values)

	s := buf.String()
	if s == "" {
		t.Fail()
	} else if strings.Contains(s, c.ApiKey) {
		t.Errorf("API key was included in the output")
	} else if expected != s {
		t.Errorf("expected: %q\ngot: %q", expected, s)
	}
}

func TestLogSendOsxCurlVerifyCertificate(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Hostname:          "127.0.0.1",
		ApiKey:            "secret",
		Protocol:          "https",
		VerifyCertificate: true,
		Logging:           LogOsxCurl,
	}
	values := url.Values{
		"type":   []string{"config"},
		"action": []string{"delete"},
		"xpath":  []string{"mypathhere"},
		"key":    []string{c.ApiKey},
	}
	c.logSend(values)

	s := buf.String()
	if s == "" {
		t.Fail()
	} else if strings.Contains(s, c.ApiKey) {
		t.Errorf("API key was included in the output")
	} else if strings.Contains(s, " -k ") {
		t.Errorf("Insecure flag was incorrectly present")
	}
}

func TestLogSendOsxCurlWithoutPersonalDataWithElement(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Hostname: "127.0.0.1",
		ApiKey:   "secret",
		Protocol: "https",
		Logging:  LogOsxCurl,
		Headers: map[string]string{
			"X-Header-1": "value",
		},
	}
	element := "<my><xml/></my>"
	values := url.Values{
		"type":   []string{"config"},
		"action": []string{"set"},
		"xpath":  []string{"mypathhere"},
		"key":    []string{"APIKEY"},
	}
	lines := []string{
		fmt.Sprintf("curl -k --data-urlencode element@element.xml 'https://HOST/api?%s'\n", values.Encode()),
		"element.xml:\n",
		fmt.Sprintf("%s\n", element),
	}
	expected := strings.Join(lines, "")
	values.Set("element", element)
	values.Set("key", c.ApiKey)
	c.logSend(values)

	s := buf.String()
	if s == "" {
		t.Fail()
	} else if strings.Contains(s, c.ApiKey) {
		t.Errorf("API key was included in the output")
	} else if strings.Contains(s, " --header ") {
		t.Errorf("Header was incorrectly included")
	} else if expected != s {
		t.Errorf("expected: %q\ngot: %q", expected, s)
	}
}

func TestLogSendOsxCurlWithPersonalDataHeadersWithValues(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := &Client{
		Hostname: "127.0.0.1",
		ApiKey:   "secret",
		Protocol: "https",
		Logging:  LogOsxCurl | LogCurlWithPersonalData,
		Headers: map[string]string{
			"X-Header-1": "wu tang clan",
			"X-Do-Stuff": "daft punk",
		},
	}
	values := url.Values{
		"type":   []string{"config"},
		"action": []string{"delete"},
		"xpath":  []string{"mypathhere"},
		"key":    []string{c.ApiKey},
	}
	c.logSend(values)

	s := buf.String()

	if s == "" {
		t.Fail()
	} else {
		regular := []string{
			" -k ",
			c.Hostname,
			c.ApiKey,
		}
		expected := make([]string, 0, len(regular)+len(c.Headers))
		for _, k := range regular {
			expected = append(expected, k)
		}
		for k, v := range c.Headers {
			expected = append(expected, fmt.Sprintf(" --header '%s: %s'", k, v))
		}
		for _, v := range expected {
			if !strings.Contains(s, v) {
				t.Errorf("Could not find %q in output: %s", v, s)
			}
		}
	}
}

func TestGetOnMissingObjectReturnsObjectNotFoundError(t *testing.T) {
	resp := `<response status="success" code="7"><result/></response>`

	c := &Client{}

	err := c.endCommunication([]byte(resp), nil)
	if err == nil {
		t.Errorf("Error is nil")
	}
	e2, ok := err.(errors.Panos)
	if !ok {
		t.Errorf("Error is not an errors.Panos")
	}
	if !e2.ObjectNotFound() {
		t.Errorf("Is not an object not found error")
	}
}

func TestShowOnMissingObjectReturnsNoSuchNode(t *testing.T) {
	expected := "No such node"
	resp := `<response status="error"><msg><line>No such node</line></msg></response>`

	c := &Client{}

	err := c.endCommunication([]byte(resp), nil)
	if err == nil {
		t.Errorf("Error is nil")
	}
	e2, ok := err.(errors.Panos)
	if !ok {
		t.Errorf("Error is not an errors.Panos")
	}
	if e2.Msg != "No such node" {
		t.Errorf("Error is %q, not %q", e2.Msg, expected)
	}
}

func TestMultilineErrorMessage(t *testing.T) {
	expected := "HTTP method must be GET"
	resp := fmt.Sprintf(`<response status="error" code="12"><msg><line><![CDATA[ tf123456 -> server -> first server  is invalid. %s, Username/Password must not be empty when Tag Distribution is chosen]]></line><line><![CDATA[ tf123456 -> server is invalid]]></line></msg></response>`, expected)

	c := &Client{}

	err := c.endCommunication([]byte(resp), nil)
	if err == nil {
		t.Errorf("Error is nil")
	}
	e2, ok := err.(errors.Panos)
	if !ok {
		t.Errorf("Error is not an errors.Panos")
	}
	if !strings.Contains(e2.Msg, expected) {
		t.Errorf("Error does not have the %q substring in the first error line", expected)
	}
}
