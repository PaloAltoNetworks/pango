package pango

import (
	"bytes"
	"encoding/xml"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PaloAltoNetworks/pango/commit"
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
