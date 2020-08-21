package router

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// PanoRouter is the client.Network.VirtualRouter namespace.
type PanoRouter struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *PanoRouter) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

/*
SetInterface performs a SET to add an interface to a virtual router.

The virtual router can be either a string or an Entry object.
*/
func (c *PanoRouter) SetInterface(tmpl, ts string, vr interface{}, iface string) error {
	var name string

	if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	switch v := vr.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s set interface: %s", singular, v)
	}

	c.con.LogAction("(set) interface for %s %q: %s", singular, name, iface)

	path := c.xpath(tmpl, ts, []string{name})
	path = append(path, "interface")

	_, err := c.con.Set(path, util.Member{Value: iface}, nil, nil)
	return err
}

/*
DeleteInterface performs a DELETE to remove an interface from a virtual router.

The virtual router can be either a string or an Entry object.
*/
func (c *PanoRouter) DeleteInterface(tmpl, ts string, vr interface{}, iface string) error {
	var name string

	if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	switch v := vr.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s delete interface: %s", singular, v)
	}

	c.con.LogAction("(delete) interface for %s %q: %s", singular, name, iface)

	path := c.xpath(tmpl, ts, []string{name})
	path = append(path, "interface", util.AsMemberXpath([]string{iface}))

	_, err := c.con.Delete(path, nil, nil)
	return err
}

// ShowList performs SHOW to retrieve a list of virtual routers.
func (c *PanoRouter) ShowList(tmpl, ts string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(tmpl, ts, nil), result)
}

// GetList performs GET to retrieve a list of virtual routers.
func (c *PanoRouter) GetList(tmpl, ts string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(tmpl, ts, nil), result)
}

// Get performs GET to retrieve information for the given virtual router.
func (c *PanoRouter) Get(tmpl, ts, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(tmpl, ts, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *PanoRouter) GetAll(tmpl, ts string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(tmpl, ts, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given virtual router.
func (c *PanoRouter) Show(tmpl, ts, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(tmpl, ts, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *PanoRouter) ShowAll(tmpl, ts string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(tmpl, ts, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more virtual routers.
//
// Specify a non-empty vsys to import the virtual routers into the given vsys
// after creating, allowing the vsys to use them.
func (c *PanoRouter) Set(tmpl, ts, vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	} else if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Name)
	}
	path := c.xpath(tmpl, ts, names)

	if err = c.ns.Set(names, path, data); err != nil {
		return err
	}

	// Remove the virtual routers from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, tmpl, ts, names); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VirtualRouterImport, tmpl, ts, vsys, names)
}

// Edit performs EDIT to create / update a virtual router.
//
// Specify a non-empty vsys to import the virtual router into the given vsys
// after creating, allowing the vsys to use them.
func (c *PanoRouter) Edit(tmpl, ts, vsys string, e Entry) error {
	var err error

	if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	_, fn := c.versioning()
	path := c.xpath(tmpl, ts, []string{e.Name})
	data := fn(e)

	if err = c.ns.Edit(e.Name, path, data); err != nil {
		return err
	}

	// Remove the virtual routers from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, tmpl, ts, []string{e.Name}); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VirtualRouterImport, tmpl, ts, vsys, []string{e.Name})
}

// Delete removes the given virtual routers from the firewall.
//
// Virtual routers can be a string or an Entry object.
func (c *PanoRouter) Delete(tmpl, ts string, e ...interface{}) error {
	if len(e) == 0 {
		return nil
	} else if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
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
	path := c.xpath(tmpl, ts, names)

	// Unimport virtual routers.
	if err := c.con.VsysUnimport(util.VirtualRouterImport, tmpl, ts, names); err != nil {
		return err
	}

	return c.ns.Delete(names, path)
}

// CleanupDefault clears the `default` route configuration instead of deleting
// it outright.  This involves unimporting the route "default" from the given
// vsys, then performing an `EDIT` with an empty router.Entry object.
func (c *PanoRouter) CleanupDefault(tmpl, ts string) error {
	var err error

	if tmpl == "" || ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	c.con.LogAction("(action) cleaning up default route")

	// Unimport the default virtual router.
	if err = c.con.VsysUnimport(util.VirtualRouterImport, tmpl, ts, []string{"default"}); err != nil {
		return err
	}

	// Cleanup the interfaces the virtual router refers to.
	info := Entry{Name: "default"}
	if err = c.Edit(tmpl, ts, "", info); err != nil {
		return err
	}

	return nil
}

/** Internal functions for this namespace struct **/

func (c *PanoRouter) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *PanoRouter) xpath(tmpl, ts string, vals []string) []string {
	ans := make([]string, 0, 11)
	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"virtual-router",
		util.AsEntryXpath(vals),
	)

	return ans
}
