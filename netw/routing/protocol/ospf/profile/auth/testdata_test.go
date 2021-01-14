package auth

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 auth profile password", version.Number{10, 0, 0, ""}, Entry{
			Name:     "t1",
			AuthType: AuthTypePassword,
			Password: "password",
		}},
		{"v1 auth profile md5", version.Number{10, 0, 0, ""}, Entry{
			Name:     "t2",
			AuthType: AuthTypeMd5,
			Md5Keys: []Md5Key{
				Md5Key{
					KeyId:     10,
					Key:       "key 10",
					Preferred: true,
				},
				Md5Key{
					KeyId:     20,
					Key:       "key 20",
					Preferred: false,
				},
			},
		}},
	}
}
