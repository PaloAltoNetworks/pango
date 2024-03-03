package xmlapi

import (
	"fmt"
	"net/url"
)

// Op performs a type=op XML API action.
type Op struct {
	// The comamnd to execute.  This can be a properly formatting XML string,
	// a fmt.Stringer, or a struct that can mer marshalled into XML.
	Command any

	// The vsys to send the command to.
	Vsys string

	// The target, copied from the client connection.
	Target string

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Op) AsUrlValues() (url.Values, error) {
	if c.Command == nil {
		return nil, fmt.Errorf("no command specified")
	}

	data := url.Values{}
	data.Set("type", "op")

	if err := addToData("cmd", c.Command, true, &data); err != nil {
		return nil, err
	}

	if c.Vsys != "" {
		data.Set("vsys", c.Vsys)
	}

	if c.Target != "" {
		data.Set("target", c.Target)
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
