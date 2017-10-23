package general

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        d string
        c Config
    }{
        {"test1", Config{
            Hostname: "host1",
            IpAddress: "10.1.1.2",
            Netmask: "255.255.255.0",
            Gateway: "10.1.1.1",
            Timezone: "US/Pacific",
            DnsPrimary: "10.1.1.1",
            DnsSecondary: "10.1.1.50",
            NtpPrimaryAddress: "10.1.1.1",
            NtpPrimaryAuthType: NoAuth,
            NtpSecondaryAddress: "10.1.1.51",
            NtpSecondaryAuthType: SymmetricKeyAuth,
            NtpSecondaryKeyId: 1,
            NtpSecondaryAlgorithm: Sha1,
            NtpSecondaryAuthKey: "secret",
        }},
        {"test2", Config{
            IpAddress: "10.2.1.2",
            Netmask: "255.255.0.0",
            Gateway: "10.2.1.1",
            Timezone: "UTC",
            DnsPrimary: "10.2.1.5",
            NtpPrimaryAddress: "10.2.5.7",
            NtpPrimaryAuthType: SymmetricKeyAuth,
            NtpPrimaryKeyId: 5,
            NtpPrimaryAlgorithm: Md5,
            NtpPrimaryAuthKey: "password",
            NtpSecondaryAddress: "10.2.5.8",
            NtpSecondaryAuthType: AutokeyAuth,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &General{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.d, func(t *testing.T) {
            var err error
            mc.AddResp("")
            err = ns.Set(tc.c)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get()
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.c, r) {
                    t.Errorf("%#v != %#v", tc.c, r)
                }
            }
        })
    }
}
