package pnrm

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    o := &Pnrm{}
    o.Initialize(mc)

    if o.DeviceGroup == nil {
        t.Fail()
    }
}

