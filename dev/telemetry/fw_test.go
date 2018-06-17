package telemetry

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        conf Settings
    }{
        {"all no", Settings{
        }},
        {"all yes", Settings{
            ApplicationReports: true,
            ThreatPreventionReports: true,
            UrlReports: true,
            FileTypeIdentificationReports: true,
            ThreatPreventionData: true,
            ThreatPreventionPacketCaptures: true,
            ProductUsageStats: true,
            PassiveDnsMonitoring: true,
        }},
        {"mix", Settings{
            ApplicationReports: true,
            UrlReports: true,
            ThreatPreventionData: true,
            ProductUsageStats: true,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwTelemetry{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get()
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}

