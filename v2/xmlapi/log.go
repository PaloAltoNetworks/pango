package xmlapi

import (
	"fmt"
	"net/url"
	"strconv"
)

// Log performs a type=log XML API action.
type Log struct {
	LogType string

	Action string

	Query string

	Direction string

	Nlogs int

	Skip int

	JobId uint

	// Any extra params that need to be specified should be put here.
	Extras url.Values
}

// AsUrlValues returns the full API command as url.Values to send to PAN-OS.
func (c *Log) AsUrlValues() (url.Values, error) {
	data := url.Values{}
	data.Set("type", "log")

	if c.LogType != "" {
		data.Set("log-type", c.LogType)
	}

	if c.Action != "" {
		data.Set("action", c.Action)
	}

	if c.Query != "" {
		data.Set("query", c.Query)
	}

	if c.Direction != "" {
		data.Set("dir", c.Direction)
	}

	if c.Nlogs != 0 {
		data.Set("nlogs", strconv.Itoa(c.Nlogs))
	}

	if c.Skip != 0 {
		data.Set("skip", strconv.Itoa(c.Skip))
	}

	if c.JobId != 0 {
		data.Set("job-id", fmt.Sprintf("%d", c.JobId))
	}

	for key := range c.Extras {
		data.Set(key, c.Extras.Get(key))
	}

	return data, nil
}
