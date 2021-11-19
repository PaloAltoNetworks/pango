package certificate

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
		{"v1 bare", version.Number{7, 1, 0, ""}, Entry{
			Name:       "foo",
			CommonName: "common",
			Status:     StatusValid,
		}},
		{"v1 basic with optional args", version.Number{7, 1, 0, ""}, Entry{
			Name:           "foo",
			CommonName:     "common",
			Algorithm:      "rsa",
			Ca:             false,
			NotValidAfter:  "not valid after",
			NotValidBefore: "not valid before",
			ExpiryEpoch:    "expiryepoch",
			Subject:        "subject",
			SubjectHash:    "subject hash",
			Issuer:         "issuer",
			IssuerHash:     "issuer hash",
			Status:         StatusRevoked,
		}},
		{"v1 with csr", version.Number{7, 1, 0, ""}, Entry{
			Name:       "foo",
			CommonName: "common",
			Csr:        "csr",
			Status:     StatusValid,
		}},
		{"v1 with public key", version.Number{7, 1, 0, ""}, Entry{
			Name:       "foo",
			CommonName: "common",
			PublicKey:  "public",
			Status:     StatusRevoked,
		}},
		{"v1 with private key", version.Number{7, 1, 0, ""}, Entry{
			Name:       "foo",
			CommonName: "common",
			PrivateKey: "private key",
			Status:     StatusValid,
		}},
		{"v1 with private key on hsm", version.Number{7, 1, 0, ""}, Entry{
			Name:            "foo",
			CommonName:      "common",
			PrivateKeyOnHsm: true,
			Status:          StatusRevoked,
		}},
	}
}
