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
            Domain: "example.com",
            LoginBanner: "this is my banner",
            PanoramaPrimary: "pano1",
            PanoramaSecondary: "pano2",
            DnsPrimary: "10.1.1.1",
            DnsSecondary: "10.1.1.50",
            NtpPrimaryAddress: "10.1.1.1",
            NtpPrimaryAuthType: NoAuth,
            NtpSecondaryAddress: "10.1.1.51",
            NtpSecondaryAuthType: SymmetricKeyAuth,
            NtpSecondaryKeyId: 1,
            NtpSecondaryAlgorithm: Sha1,
            NtpSecondaryAuthKey: "secret",
            raw: map[string] string{
                "alb": "ack login banner",
                "ap": "auth profile",
                "cp": "certificate profile",
                "dlu": "domain lookup url",
                "ffrt": "fqdn force refresh time",
                "frt": "fqdn refresh time",
                "gl": "geo location",
                "hs": "hsm settings",
                "ialu": "ip address lookup url",
                "i6a": "ipv6 address",
                "i6dg": "ipv6 default gateway",
                "locale": "my locale",
                "les": "log export schedule",
                "ll": "log link",
                "mab": "motd and banner",
                "mtu": "mtu",
                "pi": "permitted ip",
                "route": "route",
                "sppassword": "secure proxy password",
                "spport": "secure proxy port",
                "sps": "secure proxy server",
                "spu": "secure proxy user",
                "service": "service list",
                "ss": "snmp setting",
                "sd": "speed duplex",
                "stsp": "ssl tls service profile",
                "sc": "syslog certificate",
                "type": "type",
                "us": "update schedule",
        }}},
        {"test2", Config{
            IpAddress: "10.2.1.2",
            Netmask: "255.255.0.0",
            Gateway: "10.2.1.1",
            Timezone: "UTC",
            UpdateServer: "updates.paloaltonetworks.com",
            VerifyUpdateServer: true,
            LoginBanner: "This is a secure system",
            PanoramaPrimary: "192.168.55.2",
            PanoramaSecondary: "192.168.55.3",
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
