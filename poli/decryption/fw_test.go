package decryption

import (
	"reflect"
	"strings"
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
			err := ns.Set("vsys1", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("vsys1", tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}

func TestNotPresent(t *testing.T) {
	mc := &testdata.MockClient{}
	ns := FirewallNamespace(mc)

	mc.Version = version.Number{10, 2, 0, ""}
	mc.AddResp("")

	elm := Entry{
		Name:                   "rule1",
		Uuid:                   "uuid123",
		GroupTag:               "tag123",
		Description:            "blah",
		DestinationHips:        []string{"dst2", "dst1"},
		LogFailedTlsHandshakes: true,
		LogSetting:             "my log setting",
	}

	err := ns.Set("vsys1", elm)
	if err != nil {
		t.Fatalf("Failed set: %s", err)
	}

	if strings.Contains(mc.Elm, "ssl-inbound-inspection") {
		t.Fatalf("Contains ssl-inbound-inspection")
	}
}
