package tunnel

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// FwTunnel is the client.Network.TunnelInterface namespace.
type FwTunnel struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *FwTunnel) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of tunnel interfaces.
func (c *FwTunnel) ShowList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(nil), result)
}

// GetList performs GET to retrieve a list of tunnel interfaces.
func (c *FwTunnel) GetList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(nil), result)
}

// Get performs GET to retrieve information for the given tunnel interface.
func (c *FwTunnel) Get(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *FwTunnel) GetAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given tunnel interface.
func (c *FwTunnel) Show(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *FwTunnel) ShowAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more tunnel interfaces.
//
// Specifying a non-empty vsys will import the interfaces into that vsys,
// allowing the vsys to use them.
func (c *FwTunnel) Set(vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	}

	_, fn := c.versioning()
	names := make([]string, len(e))

	// Build up the struct with the given interface configs.
	d := util.BulkElement{XMLName: xml.Name{Local: "units"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
	}
	c.con.LogAction("(set) %s: %v", plural, names)

	// Set xpath.
	path := c.xpath(names)
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the interfaces.
	_, err = c.con.Set(path, d.Config(), nil, nil)
	if err != nil {
		return err
	}

	// Remove the interfaces from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, "", "", names); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.InterfaceImport, "", "", vsys, names)
}

// Edit performs EDIT to create / update the specified tunnel interface.
//
// Specifying a non-empty vsys will import the interface into that vsys,
// allowing the vsys to use it.
func (c *FwTunnel) Edit(vsys string, e Entry) error {
	var err error

	_, fn := c.versioning()

	c.con.LogAction("(edit) %s: %q", singular, e.Name)

	// Set xpath.
	path := c.xpath([]string{e.Name})

	// Edit the interface.
	_, err = c.con.Edit(path, fn(e), nil, nil)
	if err != nil {
		return err
	}

	// Remove the interface from any vsys it's currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, "", "", []string{e.Name}); err != nil {
		return err
	}

	// Import the interface.
	return c.con.VsysImport(util.InterfaceImport, "", "", vsys, []string{e.Name})
}

// Delete removes the given tunnel interface(s) from the firewall.
//
// Interfaces can be a string or an Entry object.
func (c *FwTunnel) Delete(e ...interface{}) error {
	var err error

	if len(e) == 0 {
		return nil
	}

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
	c.con.LogAction("(delete) %s: %v", plural, names)

	// Unimport interfaces.
	if err = c.con.VsysUnimport(util.InterfaceImport, "", "", names); err != nil {
		return err
	}

	// Remove interfaces next.
	path := c.xpath(names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

/** Internal functions for this namespace struct **/

func (c *FwTunnel) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *FwTunnel) xpath(vals []string) []string {
	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"interface",
		"tunnel",
		"units",
		util.AsEntryXpath(vals),
	}
}
