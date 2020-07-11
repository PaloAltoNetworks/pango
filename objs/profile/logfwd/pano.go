package logfwd

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// PanoLogFwd is the client.Objects.LogForwardingProfile namespace.
type PanoLogFwd struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *PanoLogFwd) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *PanoLogFwd) ShowList(dg string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(dg, nil), result)
}

// GetList performs GET to retrieve a list of values.
func (c *PanoLogFwd) GetList(dg string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(dg, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *PanoLogFwd) Get(dg, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(dg, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// Show performs SHOW to retrieve information for the given uid.
func (c *PanoLogFwd) Show(dg, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(dg, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve all objects.
func (c *PanoLogFwd) GetAll(dg string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(dg, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// ShowAll performs SHOW to retrieve all objects.
func (c *PanoLogFwd) ShowAll(dg string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(dg, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *PanoLogFwd) Set(dg string, e ...Entry) error {
	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	// Build up the struct.
	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Name)
	}
	path := c.xpath(dg, names)

	return c.ns.Set(names, path, data)
}

// Edit performs EDIT to create / update one object.
func (c *PanoLogFwd) Edit(dg string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(dg, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// SetWithoutSubconfig performs a DELETE to remove any subconfig
// before performing a SET to create an object.
func (c *PanoLogFwd) SetWithoutSubconfig(dg string, e Entry) error {
	c.con.LogAction("(delete) %s subconfig for %s", singular, e.Name)

	path := c.xpath(dg, []string{e.Name})

	path = append(path, "description")
	_, _ = c.con.Delete(path, nil, nil)

	path[len(path)-1] = "match-list"
	_, _ = c.con.Delete(path, nil, nil)

	return c.Set(dg, e)
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *PanoLogFwd) Delete(dg string, e ...interface{}) error {
	names := make([]string, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names[i] = v
		case Entry:
			names[i] = v.Name
		default:
			return fmt.Errorf("Unknown type sent to delete: %s", v)
		}
	}
	path := c.xpath(dg, names)

	return c.ns.Delete(names, path)
}

/** Internal functions for this namespace struct **/

func (c *PanoLogFwd) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *PanoLogFwd) xpath(dg string, vals []string) []string {
	if dg == "" {
		dg = "shared"
	}

	ans := make([]string, 0, 8)
	ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
	ans = append(ans,
		"log-settings",
		"profiles",
		util.AsEntryXpath(vals),
	)

	return ans
}
