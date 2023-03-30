package ha

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// Panorama is the client.Device.GeneralSettings namespace.
type Panorama struct {
	ns *namespace.Standard
}

// Get performs GET to retrieve the device's general settings.
func (c *Panorama) Get(tmpl, ts, vsys string) (Config, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(tmpl, ts, vsys), "", ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve the device's general settings.
func (c *Panorama) Show(tmpl, ts, vsys string) (Config, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(tmpl, ts, vsys), "", ans)
	return first(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(tmpl, ts, vsys string, e Config) error {
	return c.ns.Set(c.pather(tmpl, ts, vsys), specifier(e))
}

// Edit performs EDIT to update the device's general settings.
func (c *Panorama) Edit(tmpl, ts, vsys string, e Config) error {
	return c.ns.Edit(c.pather(tmpl, ts, vsys), e)
}

func (c *Panorama) pather(tmpl, ts, vsys string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(tmpl, ts, vsys)
	}
}

func (c *Panorama) xpath(tmpl, ts, vsys string) ([]string, error) {
	if tmpl == "" && ts == "" {
		return nil, fmt.Errorf("tmpl or ts must be specified")
	}

	ans := make([]string, 0, 12)
	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"deviceconfig",
		"high-availability",
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
