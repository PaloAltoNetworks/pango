package vminfosource

import (
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

func TestAwsVpcDefaults(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<entry name="jack">
    <AWS-VPC>
        <description>hi</description>
    </AWS-VPC>
</entry>
`)
	ans, err := ns.Get("vsys1", "jack")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	} else if ans.AwsVpc == nil {
		t.Errorf("AwsVpc is nil")
	} else if ans.AwsVpc.Description != "hi" {
		t.Errorf("AwsVpc description is %q, not \"hi\"", ans.AwsVpc.Description)
	} else if ans.AwsVpc.UpdateInterval != 60 {
		t.Errorf("AwsVpc update interval is %d", ans.AwsVpc.UpdateInterval)
	} else if ans.AwsVpc.Timeout != 2 {
		t.Errorf("AwsVpc timeout is %d", ans.AwsVpc.Timeout)
	}
}

func TestEsxiDefaults(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<entry name="jack">
    <VMware-ESXi>
        <description>hi</description>
    </VMware-ESXi>
</entry>
`)
	ans, err := ns.Get("vsys1", "jack")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	} else if ans.Esxi == nil {
		t.Errorf("Esxi is nil")
	} else if ans.Esxi.Description != "hi" {
		t.Errorf("Esxi description is %q, not \"hi\"", ans.Esxi.Description)
	} else if ans.Esxi.Port != 443 {
		t.Errorf("Esxi port is %d", ans.Esxi.Port)
	} else if ans.Esxi.UpdateInterval != 5 {
		t.Errorf("Esxi update interval is %d", ans.Esxi.UpdateInterval)
	} else if ans.Esxi.Timeout != 2 {
		t.Errorf("Esxi timeout is %d", ans.Esxi.Timeout)
	}
}

func TestVcenterDefaults(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<entry name="jack">
    <VMware-vCenter>
        <description>hi</description>
    </VMware-vCenter>
</entry>
`)
	ans, err := ns.Get("vsys1", "jack")
	if err != nil {
		t.Errorf("Error in get: %s", err)
	} else if ans.Vcenter == nil {
		t.Errorf("Vcenter is nil")
	} else if ans.Vcenter.Description != "hi" {
		t.Errorf("Vcenter description is %q, not \"hi\"", ans.Vcenter.Description)
	} else if ans.Vcenter.Port != 443 {
		t.Errorf("Vcenter port is %d", ans.Vcenter.Port)
	} else if ans.Vcenter.UpdateInterval != 5 {
		t.Errorf("Vcenter update interval is %d", ans.Vcenter.UpdateInterval)
	} else if ans.Vcenter.Timeout != 2 {
		t.Errorf("Vcenter timeout is %d", ans.Vcenter.Timeout)
	}
}
