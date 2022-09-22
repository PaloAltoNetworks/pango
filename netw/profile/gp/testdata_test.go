package gp

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test1", Entry{
			Name: "test1",
			Encryptions: []string{
				EncryptionAes128Cbc,
				EncryptionAes128Gcm,
			},
			Authentications: []string{AuthenticationSha1},
		}},
		{"test2", Entry{
			Name: "test2",
			Encryptions: []string{
				EncryptionAes256Gcm,
			},
			Authentications: []string{AuthenticationSha1},
		}},
	}
}
