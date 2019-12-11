package srvcgrp

import (
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestPanoNormalization(t *testing.T) {
	testCases := []struct {
		desc string
		dg   string
		conf Entry
	}{
		{"test no services", "", Entry{
			Name: "one",
			Tags: []string{"one", "two"},
		}},
		{"test one service", "dg1", Entry{
			Name:     "two",
			Services: []string{"svc1"},
			Tags:     []string{"single"},
		}},
		{"test two services", "dg2", Entry{
			Name:     "three",
			Services: []string{"svc1", "svc2"},
		}},
	}

	mc := &testdata.MockClient{}
	ns := &PanoSrvcGrp{}
	ns.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
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
