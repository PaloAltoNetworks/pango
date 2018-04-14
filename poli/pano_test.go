package poli

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestPanoInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    p := &PanoPoli{}
    p.Initialize(mc)

    if p.Nat == nil || p.Security == nil {
        t.Fail()
    }
}

