package server

import (
    "github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
    desc string
    version version.Number
    conf Entry
}

func getTests() []tc {
    return []tc{
        {"v1 basic", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            Address: "addy",
            Protocol: ProtocolHttps,
            Port: 443,
            HttpMethod: "POST",
            Username: "someuser",
            Password: "somepasswd",
        }},
        {"v2 basic", version.Number{9, 0, 0, ""}, Entry{
            Name: "t1",
            Address: "http.example.com",
            Protocol: ProtocolHttp,
            Port: 443,
            HttpMethod: "POST",
            Username: "someuser",
            Password: "somepasswd",
        }},
        {"v2 with tls", version.Number{9, 0, 0, ""}, Entry{
            Name: "t1",
            Address: "http.example.com",
            Protocol: ProtocolHttp,
            Port: 443,
            HttpMethod: "POST",
            Username: "someuser",
            Password: "somepasswd",
            TlsVersion: "1.2",
            CertificateProfile: "cert-profile",
        }},
    }
}
