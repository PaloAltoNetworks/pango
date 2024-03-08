package xmlapi

import (
	"fmt"
	"net/url"
)

// Export performs a type=export XML API action.
//
// Param Category is the category.
type Export struct {
	Category string

	Action string
	JobId  uint

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Export) AsUrlValues() (url.Values, error) {
	data := url.Values{}
	data.Set("type", "export")

	if c.Category != "" {
		data.Set("category", c.Category)
	}

	if c.Action != "" {
		data.Set("action", c.Action)
	}

	if c.JobId != 0 {
		data.Set("job-id", fmt.Sprintf("%d", c.JobId))
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
