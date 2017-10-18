package objs

import (
    "testing"

    "github.com/PaloAltoNetworks/xapi/testdata"
)


func TestInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    o := &Objs{}
    o.Initialize(mc)

    if o.Address == nil || o.AddressGroup == nil || o.Services == nil {
        t.Fail()
    }
}

