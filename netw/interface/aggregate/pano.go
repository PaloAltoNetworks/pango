package aggregate

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// PanoAggregate is the client.Network.AggregateInterface namespace.
type PanoAggregate struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *PanoAggregate) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *PanoAggregate) ShowList(tmpl, ts string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(tmpl, ts, nil), result)
}

// GetList performs GET to retrieve a list of values.
func (c *PanoAggregate) GetList(tmpl, ts string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(tmpl, ts, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *PanoAggregate) Get(tmpl, ts, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(tmpl, ts, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *PanoAggregate) GetAll(tmpl, ts string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(tmpl, ts, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given uid.
func (c *PanoAggregate) Show(tmpl, ts, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(tmpl, ts, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs GET to retrieve information for all objects.
func (c *PanoAggregate) ShowAll(tmpl, ts string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(tmpl, ts, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
//
// Specifying a non-empty vsys will import the interfaces into that vsys,
// allowing the vsys to use them, as long as the interface does not have a
// mode of "ha".  Interfaces that have that mode are omitted from this
// function's followup vsys import.
func (c *PanoAggregate) Set(tmpl, ts, vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	} else if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	_, fn := c.versioning()
	names := make([]string, len(e))
	n2 := make([]string, 0, len(e))

	// Build up the struct.
	d := util.BulkElement{XMLName: xml.Name{Local: "temp"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
		if e[i].Mode != ModeHa {
			n2 = append(n2, e[i].Name)
		}
	}
	c.con.LogAction("(set) %s: %v", plural, names)

	// Set xpath.
	path := c.xpath(tmpl, ts, names)
	d.XMLName = xml.Name{Local: path[len(path)-2]}
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the objects.
	if _, err = c.con.Set(path, d.Config(), nil, nil); err != nil {
		return err
	}

	// Remove the interfaces from any vsys they're currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, tmpl, ts, names); err != nil {
		return err
	}

	// Perform the vsys import.
	return c.con.VsysImport(util.InterfaceImport, tmpl, ts, vsys, n2)
}

// Edit performs EDIT to create / update one object.
func (c *PanoAggregate) Edit(tmpl, ts, vsys string, e Entry) error {
	var err error

	if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
	}

	_, fn := c.versioning()

	c.con.LogAction("(edit) %s %q", singular, e.Name)

	// Set xpath.
	path := c.xpath(tmpl, ts, []string{e.Name})

	// Edit the object.
	if _, err = c.con.Edit(path, fn(e), nil, nil); err != nil {
		return err
	}

	// Remove the interface from any vsys it's currently in.
	if err = c.con.VsysUnimport(util.InterfaceImport, tmpl, ts, []string{e.Name}); err != nil {
		return err
	}

	// Check if we should skip the import step.
	if e.Mode == ModeHa {
		return nil
	}

	// Import the interface.
	return c.con.VsysImport(util.InterfaceImport, tmpl, ts, vsys, []string{e.Name})
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *PanoAggregate) Delete(tmpl, ts string, e ...interface{}) error {
	var err error

	if len(e) == 0 {
		return nil
	} else if tmpl == "" && ts == "" {
		return fmt.Errorf("tmpl or ts must be specified")
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
	if err = c.con.VsysUnimport(util.InterfaceImport, tmpl, ts, names); err != nil {
		return err
	}

	// Remove the objects.
	path := c.xpath(tmpl, ts, names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

/** Internal functions for this namespace struct **/

func (c *PanoAggregate) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{9, 0, 0, ""}) {
		return &container_v3{}, specify_v3
	} else if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *PanoAggregate) xpath(tmpl, ts string, vals []string) []string {
	ans := make([]string, 0, 12)
	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"interface",
		"aggregate-ethernet",
		util.AsEntryXpath(vals),
	)

	return ans
}
