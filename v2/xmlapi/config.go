package xmlapi

import (
	"fmt"
	"net/url"
)

// Config performs a type=config XML API action.
type Config struct {
	// The Action.  Can be things like, "get", "show", "set", "edit", "move", or "rename".
	// or "multi-config".
	Action string

	// The xpath the config operation is operating on, used by almost every type=config
	// command (not multi-config).
	Xpath string

	// For "set" and "edit" API calls, this should be a struct that can be turned into
	// a string by XML.Marshal(), an io.Stringer, or already a string.
	Element any

	// For "move" API calls, Where should be the location keyword while Destination
	// is the other name if Where is "before" or "after".
	Where       string
	Destination string

	// For "rename" API calls, the new name.
	NewName string

	// For "multi-config", if you want strict transactional support or not.
	StrictTransactional bool

	// The target, copied from the client connection.
	Target string

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Config) AsUrlValues() (url.Values, error) {
	if c.Action == "" {
		return nil, fmt.Errorf("no action specified")
	}

	data := url.Values{}
	data.Set("type", "config")
	data.Set("action", c.Action)

	if c.Xpath != "" {
		data.Set("xpath", c.Xpath)
	}

	if c.Element != nil {
		if err := addToData("element", c.Element, true, &data); err != nil {
			return nil, err
		}
	}

	if c.Where != "" {
		data.Set("where", c.Where)
	}

	if c.Destination != "" {
		data.Set("dst", c.Destination)
	}

	if c.NewName != "" {
		data.Set("newname", c.NewName)
	}

	if c.StrictTransactional {
		data.Set("strict-transactional", "yes")
	}

	if c.Target != "" {
		data.Set("target", c.Target)
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
