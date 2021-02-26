package ike

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
		{"v1 md5 aes128 1sec", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test1",
			DhGroup:        []string{"group1"},
			Encryption:     []string{EncryptionAes128},
			Authentication: []string{"md5"},
			LifetimeType:   TimeSeconds,
			LifetimeValue:  1,
		}},
		{"v1 sha1/256 aes128/192 2min", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test2",
			DhGroup:        []string{"group2"},
			Encryption:     []string{EncryptionAes128, EncryptionAes192},
			Authentication: []string{"sha1", "sha256"},
			LifetimeType:   TimeMinutes,
			LifetimeValue:  2,
		}},
		{"v1 md5/sha512 3des/aes256 3hr", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test3",
			DhGroup:        []string{"group5"},
			Encryption:     []string{Encryption3des, EncryptionAes128},
			Authentication: []string{"md5", "sha512"},
			LifetimeType:   TimeHours,
			LifetimeValue:  3,
		}},
		{"v1 sha384 aes192 4d", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test4",
			DhGroup:        []string{"group14"},
			Encryption:     []string{EncryptionAes192},
			Authentication: []string{"sha384"},
			LifetimeType:   TimeDays,
			LifetimeValue:  4,
		}},
		{"v2 md5 aes128 1sec", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test1",
			DhGroup:        []string{"group1"},
			Encryption:     []string{EncryptionAes128},
			Authentication: []string{"md5"},
			LifetimeType:   TimeSeconds,
			LifetimeValue:  1,
		}},
		{"v2 sha1/256 aes128/192 2min am5", version.Number{7, 0, 0, ""}, Entry{
			Name:                   "test2",
			DhGroup:                []string{"group2"},
			Encryption:             []string{EncryptionAes128, EncryptionAes192},
			Authentication:         []string{"sha1", "sha256"},
			LifetimeType:           TimeMinutes,
			LifetimeValue:          2,
			AuthenticationMultiple: 5,
		}},
		{"v2 md5/sha512 3des/aes256 3hr", version.Number{7, 1, 0, ""}, Entry{
			Name:           "test3",
			DhGroup:        []string{"group5"},
			Encryption:     []string{Encryption3des, EncryptionAes128},
			Authentication: []string{"md5", "sha512"},
			LifetimeType:   TimeHours,
			LifetimeValue:  3,
		}},
		{"v2 sha384 aes192 4d am7", version.Number{7, 1, 0, ""}, Entry{
			Name:                   "test4",
			DhGroup:                []string{"group14"},
			Encryption:             []string{EncryptionAes192},
			Authentication:         []string{"sha384"},
			LifetimeType:           TimeDays,
			LifetimeValue:          4,
			AuthenticationMultiple: 7,
		}},
	}
}
