package xmlapi

import (
	"fmt"
	"net/url"
)

type CommitAction interface {
	Action() string
	Element() any
}

// Commit performs a type=commit XML API action.
type Commit struct {
	Command CommitAction

	Target string

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Commit) AsUrlValues() (url.Values, error) {
	if c.Command == nil {
		return nil, fmt.Errorf("no command specified")
	}

	data := url.Values{}
	data.Set("type", "commit")

	if action := c.Command.Action(); action != "" {
		data.Set("action", action)
	}

	if err := addToData("cmd", c.Command, true, &data); err != nil {
		return nil, err
	}

	if c.Target != "" {
		data.Set("target", c.Target)
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
