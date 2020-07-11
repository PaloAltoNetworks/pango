package logfwd

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// FwLogFwd is the client.Objects.LogForwardingProfile namespace.
type FwLogFwd struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *FwLogFwd) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *FwLogFwd) ShowList(vsys string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(vsys, nil), result)
}

// GetList performs GET to retrieve a list of values.
func (c *FwLogFwd) GetList(vsys string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(vsys, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *FwLogFwd) Get(vsys, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(vsys, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// Show performs SHOW to retrieve information for the given uid.
func (c *FwLogFwd) Show(vsys, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(vsys, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve all objects.
func (c *FwLogFwd) GetAll(vsys string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(vsys, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// ShowAll performs SHOW to retrieve all objects.
func (c *FwLogFwd) ShowAll(vsys string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(vsys, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *FwLogFwd) Set(vsys string, e ...Entry) error {
	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Name)
	}
	path := c.xpath(vsys, names)

	return c.ns.Set(names, path, data)
}

// Edit performs EDIT to create / update one object.
func (c *FwLogFwd) Edit(vsys string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(vsys, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// SetWithoutSubconfig performs a DELETE to remove any subconfig
// before performing a SET to create an object.
func (c *FwLogFwd) SetWithoutSubconfig(vsys string, e Entry) error {
	c.con.LogAction("(delete) %s subconfig for %s", singular, e.Name)

	path := c.xpath(vsys, []string{e.Name})

	path = append(path, "description")
	_, _ = c.con.Delete(path, nil, nil)

	path[len(path)-1] = "match-list"
	_, _ = c.con.Delete(path, nil, nil)

	return c.Set(vsys, e)
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *FwLogFwd) Delete(vsys string, e ...interface{}) error {
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

	path := c.xpath(vsys, names)
	return c.ns.Delete(names, path)
}

/** Internal functions for this namespace struct **/

func (c *FwLogFwd) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *FwLogFwd) xpath(vsys string, vals []string) []string {
	if vsys == "" {
		vsys = "shared"
	}

	ans := make([]string, 0, 8)
	ans = append(ans, util.VsysXpathPrefix(vsys)...)
	ans = append(ans,
		"log-settings",
		"profiles",
		util.AsEntryXpath(vals),
	)

	return ans
}
