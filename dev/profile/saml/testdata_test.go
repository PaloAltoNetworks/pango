package saml

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
		{"v1 admin use only", version.Number{8, 0, 0, ""}, Entry{
			Name:                                "v1 admin use only",
			AdminUseOnly:                        true,
			IdentityProviderId:                  "idpId",
			IdentityProviderCertificate:         "idpCert",
			SsoUrl:                              "ssoUrl",
			SsoBinding:                          BindingPost,
			SloUrl:                              "sloUrl",
			SloBinding:                          BindingRedirect,
			ValidateIdentityProviderCertificate: false,
			SignSamlMessage:                     false,
			MaxClockSkew:                        60,
		}},
		{"v1 validate identity provider certificate", version.Number{8, 0, 0, ""}, Entry{
			Name:                                "v1 validate identity provider certificate",
			AdminUseOnly:                        false,
			IdentityProviderId:                  "idpId",
			IdentityProviderCertificate:         "idpCert",
			SsoUrl:                              "ssoUrl",
			SsoBinding:                          BindingRedirect,
			SloUrl:                              "sloUrl",
			SloBinding:                          BindingPost,
			ValidateIdentityProviderCertificate: true,
			SignSamlMessage:                     false,
			MaxClockSkew:                        50,
		}},
		{"v1 sign saml message", version.Number{8, 0, 0, ""}, Entry{
			Name:                                "v1 sign saml message",
			AdminUseOnly:                        false,
			IdentityProviderId:                  "idpId",
			IdentityProviderCertificate:         "idpCert",
			SsoUrl:                              "ssoUrl",
			SsoBinding:                          BindingPost,
			SloUrl:                              "sloUrl",
			SloBinding:                          BindingPost,
			ValidateIdentityProviderCertificate: false,
			SignSamlMessage:                     true,
			MaxClockSkew:                        40,
		}},
	}
}
