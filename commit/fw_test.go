package commit

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFwNormalCommit(t *testing.T) {
	expected := `<commit><description>Hello</description></commit>`

	c := FirewallCommit{
		Description: "Hello",
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPartialCommitWithAdmins(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>example</description>",
		"<partial><admin>",
		"<member>a1</member>",
		"<member>a2</member>",
		"</admin></partial>",
		"</commit>",
	}

	expected := strings.Join(s, "")

	c := FirewallCommit{
		Description: "example",
		Admins:      []string{"a1", "a2"},
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPartialExcludeDeviceAndNetwork(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>example</description>",
		"<partial>",
		"<device-and-network>excluded</device-and-network>",
		"</partial>",
		"</commit>",
	}

	expected := strings.Join(s, "")

	c := FirewallCommit{
		Description:             "example",
		ExcludeDeviceAndNetwork: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPartialExcludePolicyAndObject(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>example</description>",
		"<partial>",
		"<policy-and-objects>excluded</policy-and-objects>",
		"</partial>",
		"</commit>",
	}

	expected := strings.Join(s, "")

	c := FirewallCommit{
		Description:             "example",
		ExcludePolicyAndObjects: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPartialExcludeShared(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>example</description>",
		"<partial>",
		"<shared-object>excluded</shared-object>",
		"</partial>",
		"</commit>",
	}

	expected := strings.Join(s, "")

	c := FirewallCommit{
		Description:          "example",
		ExcludeSharedObjects: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestFullPartial(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>check</description>",
		"<partial>",
		"<admin>",
		"<member>user1</member>",
		"<member>user2</member>",
		"</admin>",
		"<device-and-network>excluded</device-and-network>",
		"<shared-object>excluded</shared-object>",
		"<policy-and-objects>excluded</policy-and-objects>",
		"</partial>",
		"</commit>",
	}

	expected := strings.Join(s, "")

	c := FirewallCommit{
		Description:             "check",
		Admins:                  []string{"user1", "user2"},
		ExcludeDeviceAndNetwork: true,
		ExcludeSharedObjects:    true,
		ExcludePolicyAndObjects: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestFirewallAction(t *testing.T) {
	c := FirewallCommit{}

	if c.Action() != "" {
		t.Errorf("Action is %q and not empty string", c.Action())
	}
}
