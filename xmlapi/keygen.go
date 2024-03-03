package xmlapi

import (
	"net/url"
)

// KeyGen performs a type=keygen XML API action.
type KeyGen struct {
	// The username.
	Username string

	// The password.
	Password string

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *KeyGen) AsUrlValues() (url.Values, error) {
	data := url.Values{}
	data.Set("type", "keygen")
	data.Set("user", c.Username)
	data.Set("password", c.Password)

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
