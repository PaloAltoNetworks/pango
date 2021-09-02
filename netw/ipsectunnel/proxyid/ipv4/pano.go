package ipv4

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// Panorama is the client.Network.IpsecTunnelProxyId namespace.
type Panorama struct {
	ns *namespace.Standard
}

// GetList performs GET to retrieve a list of all objects.
func (c *Panorama) GetList(tmpl, ts, tun string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(tmpl, ts, tun), ans)
}

// ShowList performs SHOW to retrieve a list of all objects.
func (c *Panorama) ShowList(tmpl, ts, tun string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(tmpl, ts, tun), ans)
}

// Get performs GET to retrieve information for the given object.
func (c *Panorama) Get(tmpl, ts, tun, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(tmpl, ts, tun), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve information for the given object.
func (c *Panorama) Show(tmpl, ts, tun, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(tmpl, ts, tun), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(tmpl, ts, tun string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(tmpl, ts, tun), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *Panorama) ShowAll(tmpl, ts, tun string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(tmpl, ts, tun), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(tmpl, ts, tun string, e ...Entry) error {
	return c.ns.Set(c.pather(tmpl, ts, tun), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(tmpl, ts, tun string, e Entry) error {
	return c.ns.Edit(c.pather(tmpl, ts, tun), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(tmpl, ts, tun string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(tmpl, ts, tun), names, nErr)
}

// AllFromPanosConfig retrieves all objects stored in the retrieved config.
func (c *Panorama) AllFromPanosConfig(tmpl, ts, tun string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.FromPanosConfig(c.pather(tmpl, ts, tun), ans)
	return all(ans, err)
}

func (c *Panorama) pather(tmpl, ts, tun string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(tmpl, ts, tun, v)
	}
}

func (c *Panorama) xpath(tmpl, ts, tun string, vals []string) ([]string, error) {
	if tmpl == "" && ts == "" {
		return nil, fmt.Errorf("tmpl or ts must be specified")
	}
	if tun == "" {
		return nil, fmt.Errorf("tun must be specified")
	}

	ans := make([]string, 0, 15)
	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"tunnel",
		"ipsec",
		util.AsEntryXpath([]string{tun}),
		"auto-key",
		"proxy-id",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
