package netw

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    n := &Netw{}
    n.Initialize(mc)

    if n.EthernetInterface == nil || n.ManagementProfile == nil || n.Vlan == nil || n.Zone == nil || n.VirtualRouter == nil {
        t.Fail()
    }
}

