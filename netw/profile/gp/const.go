package gp

// Valid values for Entry.Encryption.
const (
	EncryptionAes128Cbc = "aes-128-cbc"
	EncryptionAes128Gcm = "aes-128-gcm"
	EncryptionAes256Gcm = "aes-256-gcm"
)

const (
	AuthenticationSha1 = "sha1"
)

const (
	singular = "globalprotect ipsec crypto profile"
	plural   = "globalprotect ipsec crypto profiles"
)
