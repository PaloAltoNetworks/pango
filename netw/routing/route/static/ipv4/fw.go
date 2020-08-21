package ipv4

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// FwIpv4 is the client.Network.StaticRoute namespace.
type FwIpv4 struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *FwIpv4) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of IPv4 routes.
func (c *FwIpv4) ShowList(vr string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(vr, nil), result)
}

// GetList performs GET to retrieve a list of IPv4 routes.
func (c *FwIpv4) GetList(vr string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(vr, nil), result)
}

// Get performs GET to retrieve information for the given IPv4 route.
func (c *FwIpv4) Get(vr, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(vr, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs a GET to retrieve information for all objects.
func (c *FwIpv4) GetAll(vr string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(vr, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given IPv4 route.
func (c *FwIpv4) Show(vr, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(vr, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs a SHOW to retrieve information for all objects.
func (c *FwIpv4) ShowAll(vr string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(vr, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more IPv4 routes.
func (c *FwIpv4) Set(vr string, e ...Entry) error {
	if vr == "" {
		return fmt.Errorf("vr must be specified")
	}

	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Name)
	}
	path := c.xpath(vr, names)

	return c.ns.Set(names, path, data)
}

// Edit performs EDIT to create / update an IPv4 route.
func (c *FwIpv4) Edit(vr string, e Entry) error {
	if vr == "" {
		return fmt.Errorf("vr must be specified")
	}

	_, fn := c.versioning()
	path := c.xpath(vr, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// Delete removes the given IPv4 routes.
//
// IPv4 routes can be a string or an Entry object.
func (c *FwIpv4) Delete(vr string, e ...interface{}) error {
	if vr == "" {
		return fmt.Errorf("vr must be specified")
	}

	names := make([]string, 0, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names = append(names, v)
		case Entry:
			names = append(names, v.Name)
		default:
			return fmt.Errorf("Unknown type sent to delete: %s", v)
		}
	}
	path := c.xpath(vr, names)

	return c.ns.Delete(names, path)
}

/** Internal functions for this namespace struct **/

func (c *FwIpv4) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{8, 0, 0, ""}) {
		return &container_v3{}, specify_v3
	} else if v.Gte(version.Number{7, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *FwIpv4) xpath(vr string, vals []string) []string {
	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"virtual-router",
		util.AsEntryXpath([]string{vr}),
		"routing-table",
		"ip",
		"static-route",
		util.AsEntryXpath(vals),
	}
}
