package radius

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 no servers", version.Number{7, 0, 0, ""}, Entry{
			Name:    "basicNoServers",
			Timeout: 5,
			Retries: 4,
		}},
		{"v1 with servers", version.Number{7, 0, 0, ""}, Entry{
			Name:         "basicWithServers",
			AdminUseOnly: true,
			Timeout:      5,
			Retries:      4,
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
			},
		}},
		{"v2 chap protocol", version.Number{8, 0, 0, ""}, Entry{
			Name:     "chapProtocol",
			Timeout:  5,
			Retries:  4,
			Protocol: Protocol{Chap: true},
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
				{
					Name:   "srvr2",
					Server: "second.example.com",
					Secret: "wish",
					Port:   1234,
				},
			},
		}},
		{"v2 pap protocol", version.Number{8, 0, 0, ""}, Entry{
			Name:     "papProtocol",
			Timeout:  5,
			Retries:  4,
			Protocol: Protocol{Pap: true},
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
			},
		}},
		{"v2 auto protocol", version.Number{8, 0, 0, ""}, Entry{
			Name:     "autoProtocol",
			Timeout:  2,
			Retries:  3,
			Protocol: Protocol{Auto: true},
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
			},
		}},
		{"v3 chap protocol", version.Number{8, 1, 0, ""}, Entry{
			Name:     "chapProtocol",
			Timeout:  5,
			Retries:  4,
			Protocol: Protocol{Chap: true},
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
				{
					Name:   "srvr2",
					Server: "second.example.com",
					Secret: "wish",
					Port:   1234,
				},
			},
		}},
		{"v3 pap protocol", version.Number{8, 1, 0, ""}, Entry{
			Name:     "papProtocol",
			Timeout:  5,
			Retries:  4,
			Protocol: Protocol{Pap: true},
			Servers: []Server{
				{
					Name:   "srvr1",
					Server: "example.com",
					Secret: "drowssap",
					Port:   1812,
				},
			},
		}},
		{"v3 peapmschapv2 with ident anon", version.Number{8, 1, 0, ""}, Entry{
			Name:    "peapmschapv2",
			Timeout: 5,
			Retries: 4,
			Protocol: Protocol{PeapMschapv2: &PeapMschapv2{
				MakeOuterIdentityAnonymous: true,
				AllowExpiredPasswordChange: false,
				CertificateProfile:         "myCert",
			}},
		}},
		{"v3 peapmschapv2 with passwd change", version.Number{8, 1, 0, ""}, Entry{
			Name:    "peapmschapv2",
			Timeout: 5,
			Retries: 4,
			Protocol: Protocol{PeapMschapv2: &PeapMschapv2{
				MakeOuterIdentityAnonymous: false,
				AllowExpiredPasswordChange: true,
				CertificateProfile:         "myCert",
			}},
		}},
		{"v3 peapwithgtc", version.Number{8, 1, 0, ""}, Entry{
			Name:    "peapgtc",
			Timeout: 5,
			Retries: 4,
			Protocol: Protocol{PeapWithGtc: &PeapWithGtc{
				MakeOuterIdentityAnonymous: true,
				CertificateProfile:         "myCert",
			}},
		}},
		{"v3 eapttlswithpap", version.Number{8, 1, 0, ""}, Entry{
			Name:    "eapttlswithpap",
			Timeout: 5,
			Retries: 4,
			Protocol: Protocol{EapTtlsWithPap: &EapTtlsWithPap{
				MakeOuterIdentityAnonymous: true,
				CertificateProfile:         "myCert",
			}},
		}},
	}
}
