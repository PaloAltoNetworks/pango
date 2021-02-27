package ipsec

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
		{"v1 esp aes128 aes192 sec kb", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test1",
			Protocol:       ProtocolEsp,
			Encryption:     []string{EncryptionAes128, EncryptionAes192},
			Authentication: []string{"md5", "none"},
			DhGroup:        "no-pfs",
			LifetimeType:   TimeSeconds,
			LifetimeValue:  1,
			LifesizeType:   SizeKb,
			LifesizeValue:  11,
		}},
		{"v1 esp 3des min mb", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test2",
			Protocol:       ProtocolEsp,
			Encryption:     []string{Encryption3des},
			Authentication: []string{"sha1"},
			DhGroup:        "group1",
			LifetimeType:   TimeMinutes,
			LifetimeValue:  2,
			LifesizeType:   SizeMb,
			LifesizeValue:  22,
		}},
		{"v1 esp aes256 hour gb", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test3",
			Protocol:       ProtocolEsp,
			Encryption:     []string{EncryptionAes256},
			Authentication: []string{"md5", "sha384"},
			DhGroup:        "group2",
			LifetimeType:   TimeHours,
			LifetimeValue:  3,
			LifesizeType:   SizeGb,
			LifesizeValue:  33,
		}},
		{"v1 ah both aes256 day tb", version.Number{6, 1, 0, ""}, Entry{
			Name:           "test4",
			Protocol:       ProtocolAh,
			Authentication: []string{"md5", "sha1", "sha512"},
			DhGroup:        "group5",
			LifetimeType:   TimeDays,
			LifetimeValue:  4,
			LifesizeType:   SizeTb,
			LifesizeValue:  44,
		}},
		{"v2 esp aes128 aes192 sec kb", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test1",
			Protocol:       ProtocolEsp,
			Encryption:     []string{EncryptionAes128, EncryptionAes192},
			Authentication: []string{"md5", "none"},
			DhGroup:        "no-pfs",
			LifetimeType:   TimeSeconds,
			LifetimeValue:  1,
			LifesizeType:   SizeKb,
			LifesizeValue:  11,
		}},
		{"v2 esp 3des min mb", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test2",
			Protocol:       ProtocolEsp,
			Encryption:     []string{Encryption3des},
			Authentication: []string{"sha1"},
			DhGroup:        "group1",
			LifetimeType:   TimeMinutes,
			LifetimeValue:  2,
			LifesizeType:   SizeMb,
			LifesizeValue:  22,
		}},
		{"v2 esp aes256 hour gb", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test3",
			Protocol:       ProtocolEsp,
			Encryption:     []string{EncryptionAes256},
			Authentication: []string{"md5", "sha384"},
			DhGroup:        "group2",
			LifetimeType:   TimeHours,
			LifetimeValue:  3,
			LifesizeType:   SizeGb,
			LifesizeValue:  33,
		}},
		{"v2 ah both aes256 day tb", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test4",
			Protocol:       ProtocolAh,
			Authentication: []string{"md5", "sha1", "sha512"},
			DhGroup:        "group5",
			LifetimeType:   TimeDays,
			LifetimeValue:  4,
			LifesizeType:   SizeTb,
			LifesizeValue:  44,
		}},
		{"v2 esp new stuff", version.Number{7, 0, 0, ""}, Entry{
			Name:           "test5",
			Protocol:       ProtocolEsp,
			Encryption:     []string{EncryptionDes, EncryptionAes128Gcm, EncryptionAes256},
			Authentication: []string{"none"},
			DhGroup:        "group19",
			LifetimeType:   TimeHours,
			LifetimeValue:  5,
		}},
	}
}
