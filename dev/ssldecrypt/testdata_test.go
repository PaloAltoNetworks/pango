package ssldecrypt

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	return []testCase{
		{"v1 basic", version.Number{8, 0, 0, ""}, Config{
			ForwardTrustCertificateRsa:            "rsa1",
			ForwardUntrustCertificateEcdsa:        "ecdsa1",
			RootCaExcludes:                        []string{"ex2", "ex1"},
			TrustedRootCas:                        []string{"root5", "root1"},
			DisabledPredefinedExcludeCertificates: []string{"dpec"},
		}},
		{"v1 with excludes", version.Number{8, 0, 0, ""}, Config{
			ForwardTrustCertificateEcdsa: "ecdsa1",
			ForwardUntrustCertificateRsa: "rsa1",
			SslDecryptExcludeCertificates: []SslDecryptExcludeCertificate{
				{
					"n1",
					"my description",
					true,
				},
				{
					"n2",
					"desc2",
					false,
				},
			},
		}},
	}
}
