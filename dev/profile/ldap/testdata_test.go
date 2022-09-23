package ldap

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
			Name:                    "v1 no servers",
			AdminUseOnly:            true,
			LdapType:                LdapTypeOther,
			Ssl:                     false,
			VerifyServerCertificate: true,
			Disabled:                false,
			BaseDn:                  "baseDn",
			BindDn:                  "bindDn",
			Password:                "secret",
			BindTimeout:             1,
			SearchTimeout:           2,
			RetryInterval:           3,
		}},
		{"v1 with servers", version.Number{7, 0, 0, ""}, Entry{
			Name:                    "v1 with servers",
			AdminUseOnly:            false,
			LdapType:                LdapTypeEdirectory,
			Ssl:                     true,
			VerifyServerCertificate: false,
			Disabled:                true,
			BaseDn:                  "baseDn",
			BindDn:                  "bindDn",
			Password:                "secret",
			BindTimeout:             10,
			SearchTimeout:           20,
			RetryInterval:           30,
			Servers: []Server{
				{
					Name:   "first",
					Server: "first.example.com",
					Port:   123,
				},
				{
					Name:   "second",
					Server: "second.example.com",
					Port:   389,
				},
			},
		}},
	}
}
