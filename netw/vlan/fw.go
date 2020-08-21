package vlan

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// FwVlan is the client.Network.Vlan namespace.
type FwVlan struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

/*
SetInterface performs a SET to add an interface to a VLAN.

The VLAN can be either a string or an Entry object.
The iface variable is the interface.
The rmMacs and addMacs params are MAC addresses to remove/add that
will reference the iface interface.
*/
func (c *FwVlan) SetInterface(vlan interface{}, iface string, rmMacs, addMacs []string) error {
	var (
		name string
		err  error
	)

	switch v := vlan.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s set interface: %s", singular, v)
	}

	c.con.LogAction("(set) interface for %s %q: %s", singular, name, iface)

	basePath := c.xpath([]string{name})
	iPath := append(basePath, "interface")

	if _, err = c.con.Set(iPath, util.Member{Value: iface}, nil, nil); err != nil {
		return err
	}

	if len(rmMacs) > 0 {
		c.con.LogAction("(delete) removing %q mac addresses: %#v", name, rmMacs)
		rPath := append(basePath, "mac", util.AsEntryXpath(rmMacs))
		if _, err = c.con.Delete(rPath, nil, nil); err != nil {
			return err
		}
	}

	if len(addMacs) > 0 {
		c.con.LogAction("(set) adding %q mac addresses: %#v", name, addMacs)
		d := util.BulkElement{XMLName: xml.Name{Local: "mac"}}
		for i := range addMacs {
			d.Data = append(d.Data, macList{Mac: addMacs[i], Interface: iface})
		}
		aPath := make([]string, 0, len(basePath)+1)
		aPath = append(aPath, basePath...)
		if len(addMacs) == 1 {
			aPath = append(aPath, "mac")
		}
		if _, err = c.con.Set(aPath, d.Config(), nil, nil); err != nil {
			return err
		}
	}

	return nil
}

/*
DeleteInterface performs a DELETE to remove an interface from a VLAN.

The VLAN can be either a string or an Entry object.

All MAC addresses referencing this interface are deleted.
*/
func (c *FwVlan) DeleteInterface(vlan interface{}, iface string) error {
	var (
		name string
		err  error
	)

	switch v := vlan.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s delete interface: %s", singular, v)
	}

	o, err := c.Get(name)
	if err != nil {
		return err
	}
	rmMacs := make([]string, 0, len(o.StaticMacs))
	for k, v := range o.StaticMacs {
		if v == iface {
			rmMacs = append(rmMacs, k)
		}
	}

	c.con.LogAction("(delete) interface for %s %q: %s", singular, name, iface)

	basePath := c.xpath([]string{name})
	mPath := append(basePath, "mac", util.AsEntryXpath(rmMacs))
	iPath := append(basePath, "interface", util.AsMemberXpath([]string{iface}))

	if len(rmMacs) > 0 {
		c.con.LogAction("(delete) removing %q mac addresses: %#v", iface, rmMacs)
		if _, err = c.con.Delete(mPath, nil, nil); err != nil {
			return err
		}
	}

	_, err = c.con.Delete(iPath, nil, nil)
	return err
}

// Initialize is invoked by client.Initialize().
func (c *FwVlan) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of VLANs.
func (c *FwVlan) ShowList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(nil), result)
}

// GetList performs GET to retrieve a list of VLANs.
func (c *FwVlan) GetList() ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(nil), result)
}

// Get performs GET to retrieve information for the given VLAN.
func (c *FwVlan) Get(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *FwVlan) GetAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given VLAN.
func (c *FwVlan) Show(name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath([]string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs GET to retrieve information for all objects.
func (c *FwVlan) ShowAll() ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more VLANs.
//
// Specify a non-empty vsys to import the VLAN(s) into the given vsys
// after creating, allowing the vsys to use them.
func (c *FwVlan) Set(vsys string, e ...Entry) error {
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

	// Remove the VLANs from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VlanImport, "", "", names); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VlanImport, "", "", vsys, names)
}

// Edit performs EDIT to create / update a VLAN.
//
// Specify a non-empty vsys to import the VLAN into the given vsys
// after creating, allowing the vsys to use it.
func (c *FwVlan) Edit(vsys string, e Entry) error {
	var err error

	_, fn := c.versioning()
	path := c.xpath([]string{e.Name})
	data := fn(e)

	if err = c.ns.Edit(e.Name, path, data); err != nil {
		return err
	}

	// Remove the VLANs from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.VlanImport, "", "", []string{e.Name}); err != nil {
		return err
	}

	// Perform vsys import next.
	return c.con.VsysImport(util.VlanImport, "", "", vsys, []string{e.Name})
}

// Delete removes the given VLAN(s) from the firewall.
//
// VLANs can be a string or an Entry object.
func (c *FwVlan) Delete(e ...interface{}) error {
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
	path := c.xpath(names)

	// Unimport VLANs.
	if err := c.con.VsysUnimport(util.VlanImport, "", "", names); err != nil {
		return err
	}

	return c.ns.Delete(names, path)
}

/** Internal functions for this namespace struct **/

func (c *FwVlan) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *FwVlan) xpath(vals []string) []string {
	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"vlan",
		util.AsEntryXpath(vals),
	}
}
