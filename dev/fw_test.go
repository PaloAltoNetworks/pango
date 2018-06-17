package dev

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    d := FwDev{}
    d.Initialize(mc)

    if d.GeneralSettings == nil {
        t.Fail()
    } else if d.Telemetry == nil {
        t.Fail()
    }
}
