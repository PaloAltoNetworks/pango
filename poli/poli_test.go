package poli

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    p := &Poli{}
    p.Initialize(mc)

    if p.Nat == nil || p.Security == nil {
        t.Fail()
    }
}

