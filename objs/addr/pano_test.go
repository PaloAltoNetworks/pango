package addr

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc    string
		dg      string
		version version.Number
		conf    Entry
	}{
		{"v1 test ip netmask", "", version.Number{7, 1, 0, ""}, Entry{
			Name:        "one",
			Value:       "10.1.1.0/24",
			Type:        IpNetmask,
			Description: "my description",
			Tags:        []string{"tag1", "tag2"},
		}},
		{"v1 test ip range", "dg1", version.Number{7, 1, 0, ""}, Entry{
			Name:        "two",
			Value:       "10.1.1.1-10.1.1.254",
			Type:        IpRange,
			Description: "my description",
			Tags:        []string{"tag3", "tag4"},
		}},
		{"v1 test fqdn", "dg2", version.Number{7, 1, 0, ""}, Entry{
			Name:        "three",
			Value:       "example.com",
			Type:        Fqdn,
			Description: "my description",
		}},
		{"v2 test ip netmask", "", version.Number{9, 0, 0, ""}, Entry{
			Name:        "one",
			Value:       "10.1.1.0/24",
			Type:        IpNetmask,
			Description: "my description",
			Tags:        []string{"tag1", "tag2"},
		}},
		{"v2 test ip range", "dg1", version.Number{9, 0, 0, ""}, Entry{
			Name:        "two",
			Value:       "10.1.1.1-10.1.1.254",
			Type:        IpRange,
			Description: "my description",
			Tags:        []string{"tag3", "tag4"},
		}},
		{"v2 test fqdn", "dg2", version.Number{9, 0, 0, ""}, Entry{
			Name:        "three",
			Value:       "example.com",
			Type:        Fqdn,
			Description: "my description",
		}},
		{"v2 test ip wildcard", "dg3", version.Number{9, 0, 0, ""}, Entry{
			Name:        "four",
			Value:       "10.20.1.0/0.0.248.255",
			Type:        IpWildcard,
			Description: "my description",
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoAddr{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set(tc.dg, tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get(tc.dg, tc.conf.Name)
				if err != nil {
					t.Errorf("Error in get: %s", err)
				} else if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}
