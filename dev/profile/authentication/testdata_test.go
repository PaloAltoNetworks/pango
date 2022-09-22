package authentication

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
		{"v1 none", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "none",
			LockoutFailedAttempts: "7",
			LockoutTime:           5,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				None: true,
			},
		}},
		{"v1 local database", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "local_database",
			LockoutFailedAttempts: "6",
			LockoutTime:           0,
			Type: AuthenticationType{
				LocalDatabase: true,
			},
		}},
		{"v1 radius", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "radius",
			LockoutFailedAttempts: "6",
			LockoutTime:           12,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Radius: &Radius{
					ServerProfile: "serverProfile",
				},
			},
		}},
		{"v1 ldap", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "ldap",
			LockoutFailedAttempts: "10",
			LockoutTime:           7,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Ldap: &Ldap{
					ServerProfile:         "myServerProfile",
					LoginAttribute:        "loginAttrib",
					PasswordExpiryWarning: "6",
				},
			},
		}},
		{"v1 kerberos", version.Number{6, 1, 0, ""}, Entry{
			Name:                  "kerberos",
			LockoutFailedAttempts: "20",
			LockoutTime:           10,
			Type: AuthenticationType{
				Kerberos: &Kerberos{
					ServerProfile: "serverProfile",
				},
			},
		}},
		{"v2 none", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "none",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "7",
			LockoutTime:           5,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				None: true,
			},
		}},
		{"v2 local database", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "local_database",
			UsernameModifier:      UsernameModifierInputDomain,
			UserDomain:            "myUserDomain",
			LockoutFailedAttempts: "6",
			LockoutTime:           0,
			Type: AuthenticationType{
				LocalDatabase: true,
			},
		}},
		{"v2 radius no group", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "radiusNoGroup",
			UsernameModifier:      UsernameModifierDomainInput,
			LockoutFailedAttempts: "6",
			LockoutTime:           12,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Radius: &Radius{
					ServerProfile: "serverProfile",
				},
			},
		}},
		{"v2 radius with group", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "radiusWithGroup",
			UsernameModifier:      UsernameModifierDomainInput,
			LockoutFailedAttempts: "6",
			LockoutTime:           12,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Radius: &Radius{
					ServerProfile:     "serverProfile",
					RetrieveUserGroup: true,
				},
			},
		}},
		{"v2 ldap", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "ldap",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "10",
			LockoutTime:           7,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Ldap: &Ldap{
					ServerProfile:         "myServerProfile",
					LoginAttribute:        "loginAttrib",
					PasswordExpiryWarning: "6",
				},
			},
		}},
		{"v2 kerberos", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "kerberos",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "20",
			LockoutTime:           10,
			Type: AuthenticationType{
				Kerberos: &Kerberos{
					ServerProfile: "serverProfile",
					Realm:         "myRealm",
				},
			},
		}},
		{"v2 tacacs plus", version.Number{7, 0, 0, ""}, Entry{
			Name:                  "tacacsPlus",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "2",
			LockoutTime:           4,
			Type: AuthenticationType{
				TacacsPlus: &TacacsPlus{
					ServerProfile: "myServerProfile",
				},
			},
		}},
		{"v3 none", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "none",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "7",
			LockoutTime:           5,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				None: true,
			},
		}},
		{"v3 none with mfa factors", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "none",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "7",
			LockoutTime:           5,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				None: true,
			},
			MultiFactorAuthentication: &MultiFactorAuthentication{
				Factors: []string{"mfab", "mfaa"},
			},
		}},
		{"v3 none with mfa enabled", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "none",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "7",
			LockoutTime:           5,
			Type: AuthenticationType{
				None: true,
			},
			MultiFactorAuthentication: &MultiFactorAuthentication{
				Enabled: true,
			},
		}},
		{"v3 local database", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "local_database",
			UsernameModifier:      UsernameModifierInputDomain,
			UserDomain:            "myUserDomain",
			LockoutFailedAttempts: "6",
			LockoutTime:           0,
			Type: AuthenticationType{
				LocalDatabase: true,
			},
		}},
		{"v3 radius no group", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "radiusNoGroup",
			UsernameModifier:      UsernameModifierDomainInput,
			LockoutFailedAttempts: "6",
			LockoutTime:           12,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Radius: &Radius{
					ServerProfile: "serverProfile",
				},
			},
		}},
		{"v3 radius with group", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "radiusWithGroup",
			UsernameModifier:      UsernameModifierDomainInput,
			LockoutFailedAttempts: "6",
			LockoutTime:           12,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Radius: &Radius{
					ServerProfile:     "serverProfile",
					RetrieveUserGroup: true,
				},
			},
		}},
		{"v3 ldap", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "ldap",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "10",
			LockoutTime:           7,
			AllowList:             []string{"allow1", "allow2"},
			Type: AuthenticationType{
				Ldap: &Ldap{
					ServerProfile:         "myServerProfile",
					LoginAttribute:        "loginAttrib",
					PasswordExpiryWarning: "6",
				},
			},
		}},
		{"v3 kerberos", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "kerberos",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "20",
			LockoutTime:           10,
			Type: AuthenticationType{
				Kerberos: &Kerberos{
					ServerProfile: "serverProfile",
					Realm:         "myRealm",
				},
			},
		}},
		{"v3 tacacs plus no group", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "tacacsPlus",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "2",
			LockoutTime:           4,
			Type: AuthenticationType{
				TacacsPlus: &TacacsPlus{
					ServerProfile: "myServerProfile",
				},
			},
		}},
		{"v3 tacacs plus with group", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "tacacsPlusWithGroup",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "2",
			LockoutTime:           4,
			Type: AuthenticationType{
				TacacsPlus: &TacacsPlus{
					ServerProfile:     "myServerProfile",
					RetrieveUserGroup: true,
				},
			},
		}},
		{"v3 saml no single logout", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "saml no logout",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "2",
			LockoutTime:           4,
			Type: AuthenticationType{
				Saml: &Saml{
					ServerProfile:             "saml server profile",
					EnableSingleLogout:        false,
					RequestSigningCertificate: "reqSigner",
					CertificateProfile:        "certProfile",
					UsernameAttribute:         UsernameAttributeDefault,
					UserGroupAttribute:        "ug attr",
					AdminRoleAttribute:        "ar attr",
					AccessDomainAttribute:     "ad attr",
				},
			},
		}},
		{"v3 saml with single logout", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "saml with logout",
			UsernameModifier:      UsernameModifierInput,
			LockoutFailedAttempts: "2",
			LockoutTime:           4,
			Type: AuthenticationType{
				Saml: &Saml{
					ServerProfile:             "saml server profile",
					EnableSingleLogout:        true,
					RequestSigningCertificate: "reqSigner",
					CertificateProfile:        "certProfile",
					UsernameAttribute:         UsernameAttributeDefault,
					UserGroupAttribute:        "ug attr",
					AdminRoleAttribute:        "ar attr",
					AccessDomainAttribute:     "ad attr",
				},
			},
		}},
	}
}
