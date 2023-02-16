package security

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestFwNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.vsys, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.vsys, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}

func TestLogEndMissingIsTrue(t *testing.T) {
	versions := []version.Number{
		{6, 1, 0, ""},
		{9, 0, 0, ""},
		{10, 0, 0, ""},
	}

	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	mc.AddResp(`
<entry name="foo">
    <description>Hello world</description>
</entry>
`)
	for _, v := range versions {
		t.Run(fmt.Sprintf("version %s", v), func(t *testing.T) {
			mc.Version = v
			r, err := ns.Get("vsys1", "foo")
			if err != nil {
				t.Errorf("Error in get: %s", err)
			} else if r.Description != "Hello world" {
				t.Errorf("Description is wrong: %s", r.Description)
			} else if !r.LogEnd {
				t.Errorf("LogEnd is not enabled")
			}
		})
	}
}

func TestHipProfilesIsAbsent(t *testing.T) {
	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	mc.Version = version.Number{10, 1, 5, ""}
	mc.AddResp("")

	elm := Entry{
		Name:                 "rule1",
		Type:                 "universal",
		SourceZones:          []string{"sz1", "sz2"},
		SourceAddresses:      []string{"sa1", "sa2"},
		SourceUsers:          []string{"su1", "su2"},
		DestinationZones:     []string{"dz1", "dz2"},
		HipProfiles:          []string{"hip1", "hip2"},
		DestinationAddresses: []string{"da1", "da2"},
		Applications:         []string{"app1"},
		Services:             []string{"s2", "s1"},
		Categories:           []string{"cat1"},
		Action:               "allow",
		LogEnd:               true,
		SourceDevices:        []string{"src2", "src1"},
		DestinationDevices:   []string{"dstDev"},
	}

	err := ns.Set("vsys1", elm)
	if err != nil {
		t.Fatalf("Failed set: %s", err)
	}
	mc.AddResp(mc.Elm)
	r, err := ns.Get("vsys1", elm.Name)
	if err != nil {
		t.Fatalf("Failed get: %s", err)
	}
	if len(r.HipProfiles) != 0 {
		t.Fatalf("HipProfiles has data and shouldn't")
	}
}
