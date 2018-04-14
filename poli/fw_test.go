package poli

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    p := &FwPoli{}
    p.Initialize(mc)

    if p.Nat == nil || p.Security == nil {
        t.Fail()
    }
}

