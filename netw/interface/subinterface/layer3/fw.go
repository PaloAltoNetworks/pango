package layer3

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// FwLayer3 is the client.Network.Layer3Subinterface namespace.
type FwLayer3 struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *FwLayer3) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *FwLayer3) ShowList(iType, eth string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(iType, eth, nil), result)
}

// GetList performs GET to retrieve a list of values.
func (c *FwLayer3) GetList(iType, eth string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(iType, eth, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *FwLayer3) Get(iType, eth, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(iType, eth, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information on all objects.
func (c *FwLayer3) GetAll(iType, eth string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(iType, eth, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given uid.
func (c *FwLayer3) Show(iType, eth, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(iType, eth, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information on all objects.
func (c *FwLayer3) ShowAll(iType, eth string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(iType, eth, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *FwLayer3) Set(iType, eth, vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	} else if iType == "" {
		return fmt.Errorf("iType must be specified")
	} else if eth == "" {
		return fmt.Errorf("eth must be specified")
	}

	_, fn := c.versioning()
	names := make([]string, len(e))

	// Build up the struct.
	d := util.BulkElement{XMLName: xml.Name{Local: "units"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
	}
	c.con.LogAction("(set) %s: %v", plural, names)

	// Set xpath.
	path := c.xpath(iType, eth, names)
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the objects.
	_, err = c.con.Set(path, d.Config(), nil, nil)
	if err != nil {
		return err
	}

	// Remove from any vsys it's currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, "", "", names); err != nil {
		return err
	}

	// Perform vsys import.
	return c.con.VsysImport(util.InterfaceImport, "", "", vsys, names)
}

// Edit performs EDIT to create / update one object.
func (c *FwLayer3) Edit(iType, eth, vsys string, e Entry) error {
	var err error

	if iType == "" {
		return fmt.Errorf("iType must be specified")
	} else if eth == "" {
		return fmt.Errorf("eth must be specified")
	}

	_, fn := c.versioning()

	c.con.LogAction("(edit) %s %q", singular, e.Name)

	// Set xpath.
	path := c.xpath(iType, eth, []string{e.Name})

	// Edit the object.
	if _, err = c.con.Edit(path, fn(e), nil, nil); err != nil {
		return err
	}

	// Remove from any vsys it's currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, "", "", []string{e.Name}); err != nil {
		return err
	}

	// Perform vsys import.
	return c.con.VsysImport(util.InterfaceImport, "", "", vsys, []string{e.Name})
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *FwLayer3) Delete(iType, eth string, e ...interface{}) error {
	var err error

	if len(e) == 0 {
		return nil
	} else if iType == "" {
		return fmt.Errorf("iType must be specified")
	} else if eth == "" {
		return fmt.Errorf("eth must be specified")
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

	// Remove the objects.
	path := c.xpath(iType, eth, names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

/** Internal functions for this namespace struct **/

func (c *FwLayer3) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{9, 0, 0, ""}) {
		return &container_v3{}, specify_v3
	} else if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *FwLayer3) xpath(iType, eth string, vals []string) []string {
	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"interface",
		iType,
		util.AsEntryXpath([]string{eth}),
		"layer3",
		"units",
		util.AsEntryXpath(vals),
	}
}
