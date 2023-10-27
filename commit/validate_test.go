package commit

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestPanoramaValidate_ValidateElement(t *testing.T) {
	testS := []string{
		"<validate>",
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
		"</validate>",
	}

	expected := strings.Join(testS, "")

	validateStruct := PanoramaValidate{
		Admins:                    []string{"admin1", "admin2"},
		DeviceGroups:              []string{"dg1"},
		LogCollectors:             []string{"lc1"},
		LogCollectorGroups:        []string{"lcg1"},
		Templates:                 []string{"tmpl1"},
		TemplateStacks:            []string{"ts1"},
		WildfireAppliances:        []string{"wfa1"},
		WildfireApplianceClusters: []string{"wfc1"},
		ExcludeDeviceAndNetwork:   true,
		ExcludeSharedObjects:      true,
	}

	marshV, _ := xml.Marshal(validateStruct.Element())
	if expected != string(marshV) {
		t.Errorf("Expected(%s) got(%s)", expected, marshV)
	}
}

func TestPanoramaValidateAll_ValidateAllElement(t *testing.T) {
	testS := []string{
		"<validate>",
		"<full>",
		"</full>",
		"</validate>",
	}

	expected := strings.Join(testS, "")
	validateStruct := PanoramaValidate{}

	marshV, _ := xml.Marshal(validateStruct.Element())
	if expected != string(marshV) {
		t.Errorf("Expected(%s) got(%s)", expected, marshV)
	}
}
