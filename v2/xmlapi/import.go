package xmlapi

import (
	"net/url"
)

// Import performs a type=import XML API action.
//
// Param Category is the category.
type Import struct {
	Category string

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Import) AsUrlValues() (url.Values, error) {
	data := url.Values{}
	data.Set("type", "import")

	if c.Category != "" {
		data.Set("category", c.Category)
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
