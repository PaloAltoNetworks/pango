package zone

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// FwZone is a namespace struct, included as part of pango.Client.
type FwZone struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked when Initialize on the pango.Client is called.
func (c *FwZone) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

/*
SetInterface performs a SET to add an interface to a zone.

The zone can be either a string or an Entry object.
*/
func (c *FwZone) SetInterface(vsys string, zone interface{}, mode, iface string) error {
	var name string

	switch v := zone.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s set interface: %s", singular, v)
	}

	c.con.LogAction("(set) %s interface: %s", singular, name)

	path := c.xpath(vsys, []string{name})
	path = append(path, "network", mode)

	_, err := c.con.Set(path, util.Member{Value: iface}, nil, nil)
	return err
}

/*
DeleteInterface performs a DELETE to remove the interface from the zone.

The zone can be either a string or an Entry object.
*/
func (c *FwZone) DeleteInterface(vsys string, zone interface{}, mode, iface string) error {
	var name string

	switch v := zone.(type) {
	case string:
		name = v
	case Entry:
		name = v.Name
	default:
		return fmt.Errorf("Unknown type sent to %s delete interface: %s", singular, v)
	}

	c.con.LogAction("(delete) %s interface: %s", singular, name)

	path := c.xpath(vsys, []string{name})
	path = append(path, "network", mode, util.AsMemberXpath([]string{iface}))

	_, err := c.con.Delete(path, nil, nil)
	return err
}

// GetList performs GET to retrieve a list of values.
func (c *FwZone) GetList(vsys string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(vsys, nil), result)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *FwZone) ShowList(vsys string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(vsys, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *FwZone) Get(vsys, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(vsys, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *FwZone) GetAll(vsys string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(vsys, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Get performs SHOW to retrieve information for the given uid.
func (c *FwZone) Show(vsys, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(vsys, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *FwZone) ShowAll(vsys string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(vsys, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *FwZone) Set(vsys string, e ...Entry) error {
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
func (c *FwZone) Edit(vsys string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(vsys, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// Delete removes the given objects.
//
// Objects can be either a string or an Entry object.
func (c *FwZone) Delete(vsys string, e ...interface{}) error {
	names := make([]string, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names[i] = v
		case Entry:
			names[i] = v.Name
		default:
			return fmt.Errorf("Unsupported type to delete: %s", v)
		}
	}

	path := c.xpath(vsys, names)
	return c.ns.Delete(names, path)
}

/** Internal functions for this namespace struct **/

func (c *FwZone) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *FwZone) xpath(vsys string, vals []string) []string {
	if vsys == "" {
		vsys = "vsys1"
	}

	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"vsys",
		util.AsEntryXpath([]string{vsys}),
		"zone",
		util.AsEntryXpath(vals),
	}
}
