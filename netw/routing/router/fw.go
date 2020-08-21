package router

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// FwRouter is the client.Network.VirtualRouter namespace.
type FwRouter struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *FwRouter) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

/*
SetInterface performs a SET to add an interface to a virtual router.

The virtual router can be either a string or an Entry object.
*/
func (c *FwRouter) SetInterface(vr interface{}, iface string) error {
	var name string

	switch v := vr.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s set interface: %s", singular, v)
	}

	c.con.LogAction("(set) interface for %s %q: %s", singular, name, iface)

	path := c.xpath([]string{name})
	path = append(path, "interface")

	_, err := c.con.Set(path, util.Member{Value: iface}, nil, nil)
	return err
}

/*
DeleteInterface performs a DELETE to remove an interface from a virtual router.

The virtual router can be either a string or an Entry object.
*/
func (c *FwRouter) DeleteInterface(vr interface{}, iface string) error {
	var name string

	switch v := vr.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s delete interface: %s", singular, v)
	}

	c.con.LogAction("(delete) interface for %s %q: %s", singular, name, iface)

	path := c.xpath([]string{name})
	path = append(path, "interface", util.AsMemberXpath([]string{iface}))

	_, err := c.con.Delete(path, nil, nil)
	return err
}

// ShowList performs SHOW to retrieve a list of virtual routers.
func (c *FwRouter) ShowList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(nil), result)
}

// GetList performs GET to retrieve a list of virtual routers.
func (c *FwRouter) GetList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(nil), result)
}

// Get performs GET to retrieve information for the given virtual router.
func (c *FwRouter) Get(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *FwRouter) GetAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given virtual router.
func (c *FwRouter) Show(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *FwRouter) ShowAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more virtual routers.
//
// Specify a non-empty vsys to import the virtual routers into the given vsys
// after creating, allowing the vsys to use them.
func (c *FwRouter) Set(vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	}

	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Name)
	}
	path := c.xpath(names)

	if err = c.ns.Set(names, path, data); err != nil {
		return err
	}

	// Remove the virtual routers from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, "", "", names); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VirtualRouterImport, "", "", vsys, names)
}

// Edit performs EDIT to create / update a virtual router.
//
// Specify a non-empty vsys to import the virtual router into the given vsys
// after creating, allowing the vsys to use them.
func (c *FwRouter) Edit(vsys string, e Entry) error {
	var err error

	_, fn := c.versioning()
	path := c.xpath([]string{e.Name})
	data := fn(e)

	if err = c.ns.Edit(e.Name, path, data); err != nil {
		return err
	}

	// Remove the virtual routers from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, "", "", []string{e.Name}); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VirtualRouterImport, "", "", vsys, []string{e.Name})
}

// Delete removes the given virtual routers from the firewall.
//
// Virtual routers can be a string or an Entry object.
func (c *FwRouter) Delete(e ...interface{}) error {
	names := make([]string, 0, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names = append(names, v)
		case Entry:
			names = append(names, v.Name)
		default:
			return fmt.Errorf("Unknown type to delete: %s", v)
		}
	}
	path := c.xpath(names)

	// Unimport virtual routers.
	if err := c.con.VsysUnimport(util.VirtualRouterImport, "", "", names); err != nil {
		return err
	}

	return c.ns.Delete(names, path)
}

// CleanupDefault clears the `default` route configuration instead of deleting
// it outright.  This involves unimporting the route "default" from the given
// vsys, then performing an `EDIT` with an empty router.Entry object.
func (c *FwRouter) CleanupDefault() error {
	var err error

	c.con.LogAction("(action) cleaning up default route")

	// Unimport the default virtual router.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, "", "", []string{"default"}); err != nil {
		return err
	}

	// Cleanup the interfaces the virtual router refers to.
	info := Entry{Name: "default"}
	return c.Edit("", info)
}

/** Internal functions for this namespace struct **/

func (c *FwRouter) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *FwRouter) xpath(vals []string) []string {
	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"virtual-router",
		util.AsEntryXpath(vals),
	}
}
