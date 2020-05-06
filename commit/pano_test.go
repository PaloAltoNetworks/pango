package commit

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestPanoNormalCommit(t *testing.T) {
	expected := `<commit><description>Hello</description></commit>`

	c := PanoramaCommit{
		Description: "Hello",
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoForcePartialCommit(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>forced partial commit</description>",
		"<force>",
		"<partial>",
		"<device-group><member>dg1</member></device-group>",
		"<shared-object>excluded</shared-object>",
		"</partial>",
		"</force>",
		"</commit>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommit{
		Description:          "forced partial commit",
		DeviceGroups:         []string{"dg1"},
		ExcludeSharedObjects: true,
		Force:                true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoFullCommit(t *testing.T) {
	s := []string{
		"<commit>",
		"<description>full</description>",
		"<partial>",
		"<admin><member>admin1</member><member>admin2</member></admin>",
		"<device-group><member>dg1</member></device-group>",
		"<template><member>tmpl1</member></template>",
		"<template-stack><member>ts1</member></template-stack>",
		"<wildfire-appliance><member>wfa1</member></wildfire-appliance>",
		"<wildfire-appliance-cluster><member>wfc1</member></wildfire-appliance-cluster>",
		"<log-collector><member>lc1</member></log-collector>",
		"<log-collector-group><member>lcg1</member></log-collector-group>",
		"<device-and-network>excluded</device-and-network>",
		"<shared-object>excluded</shared-object>",
		"</partial>",
		"</commit>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommit{
		Description:             "full",
		Admins:                  []string{"admin1", "admin2"},
		DeviceGroups:            []string{"dg1"},
		Templates:               []string{"tmpl1"},
		TemplateStacks:          []string{"ts1"},
		WildfireAppliances:      []string{"wfa1"},
		WildfireClusters:        []string{"wfc1"},
		LogCollectors:           []string{"lc1"},
		LogCollectorGroups:      []string{"lcg1"},
		ExcludeDeviceAndNetwork: true,
		ExcludeSharedObjects:    true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAction(t *testing.T) {
	c := PanoramaCommit{}

	if c.Action() != "" {
		t.Errorf("Action is %q instead of empty string", c.Action())
	}
}

func TestPanoBlankCommitAll(t *testing.T) {
	s := []string{
		"<commit-all>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllDeviceGroup(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<shared-policy>",
		"<device-group>",
		"<entry name=\"foo\"><devices><entry name=\"d1\"></entry></devices></entry>",
		"</device-group>",
		"<description>device group</description>",
		"<include-template>yes</include-template>",
		"<force-template-values>no</force-template-values>",
		"</shared-policy>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:            TypeDeviceGroup,
		Name:            "foo",
		Description:     "device group",
		Devices:         []string{"d1"},
		IncludeTemplate: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllTemplate(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<template>",
		"<name>tmpl1</name>",
		"<description>template</description>",
		"<device><member>12321</member></device>",
		"<force-template-values>yes</force-template-values>",
		"</template>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:                TypeTemplate,
		Name:                "tmpl1",
		Description:         "template",
		Devices:             []string{"12321"},
		ForceTemplateValues: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllTemplateStack(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<template-stack>",
		"<name>ts1</name>",
		"<description>template stack</description>",
		"<device><member>12321</member></device>",
		"<force-template-values>yes</force-template-values>",
		"</template-stack>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:                TypeTemplateStack,
		Name:                "ts1",
		Description:         "template stack",
		Devices:             []string{"12321"},
		ForceTemplateValues: true,
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllLogCollectorGroup(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<log-collector-config>",
		"<log-collector-group>lcg1</log-collector-group>",
		"<description>log collector</description>",
		"</log-collector-config>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:        TypeLogCollectorGroup,
		Name:        "lcg1",
		Description: "log collector",
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllWildfireAppliance(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<wildfire-appliance-config>",
		"<description>wf check</description>",
		"<wildfire-appliance>wfa1</wildfire-appliance>",
		"</wildfire-appliance-config>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:        TypeWildfireAppliance,
		Name:        "wfa1",
		Description: "wf check",
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllWildfireCluster(t *testing.T) {
	s := []string{
		"<commit-all>",
		"<wildfire-appliance-config>",
		"<description>wf check</description>",
		"<wildfire-appliance-cluster>wfc1</wildfire-appliance-cluster>",
		"</wildfire-appliance-config>",
		"</commit-all>",
	}
	expected := strings.Join(s, "")

	c := PanoramaCommitAll{
		Type:        TypeWildfireCluster,
		Name:        "wfc1",
		Description: "wf check",
	}

	b, _ := xml.Marshal(c.Element())
	if expected != string(b) {
		t.Errorf("Expected(%s) got(%s)", expected, b)
	}
}

func TestPanoCommitAllAction(t *testing.T) {
	c := PanoramaCommitAll{}
	expected := "all"

	if c.Action() != expected {
		t.Errorf("Commit all action should be %q, not %q", expected, c.Action())
	}
}
