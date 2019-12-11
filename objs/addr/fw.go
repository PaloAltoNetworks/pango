package addr

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// FwAddr is a namespace struct, included as part of pango.Firewall.
type FwAddr struct {
	con util.XapiClient
}

// Initialize is invoked when Initialize on the pango.Client is called.
func (c *FwAddr) Initialize(con util.XapiClient) {
	c.con = con
}

// GetList performs GET to retrieve a list of address objects.
func (c *FwAddr) GetList(vsys string) ([]string, error) {
	c.con.LogQuery("(get) list of address objects")
	path := c.xpath(vsys, nil)
	return c.con.EntryListUsing(c.con.Get, path[:len(path)-1])
}

// ShowList performs SHOW to retrieve a list of address objects.
func (c *FwAddr) ShowList(vsys string) ([]string, error) {
	c.con.LogQuery("(show) list of address objects")
	path := c.xpath(vsys, nil)
	return c.con.EntryListUsing(c.con.Show, path[:len(path)-1])
}

// Get performs GET to retrieve information for the given address object.
func (c *FwAddr) Get(vsys, name string) (Entry, error) {
	c.con.LogQuery("(get) address object %q", name)
	listing, err := c.details(c.con.Get, vsys, name)
	if err == nil && len(listing) > 0 {
		return listing[0], nil
	}
	return Entry{}, err
}

// GetAll performs a GET to retrieve objects.
func (c *FwAddr) GetAll(vsys string) ([]Entry, error) {
	c.con.LogQuery("(get) all address objects")
	return c.details(c.con.Get, vsys, "")
}

// Show performs SHOW to retrieve information for the given address object.
func (c *FwAddr) Show(vsys, name string) (Entry, error) {
	c.con.LogQuery("(show) address object %q", name)
	listing, err := c.details(c.con.Show, vsys, name)
	if err == nil && len(listing) > 0 {
		return listing[0], nil
	}
	return Entry{}, err
}

// ShowAll performs SHOW to retrieve all objects.
func (c *FwAddr) ShowAll(vsys string) ([]Entry, error) {
	c.con.LogQuery("(show) all address objects")
	return c.details(c.con.Show, vsys, "")
}

// Set performs SET to create / update one or more address objects.
func (c *FwAddr) Set(vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	}

	_, fn := c.versioning()
	names := make([]string, len(e))

	// Build up the struct with the given configs.
	d := util.BulkElement{XMLName: xml.Name{Local: "address"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
	}
	c.con.LogAction("(set) address objects: %v", names)

	// Set xpath.
	path := c.xpath(vsys, names)
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the objects.
	_, err = c.con.Set(path, d.Config(), nil, nil)
	return err
}

// Edit performs EDIT to create / update an address object.
func (c *FwAddr) Edit(vsys string, e Entry) error {
	var err error

	_, fn := c.versioning()

	c.con.LogAction("(edit) address object %q", e.Name)

	// Set xpath.
	path := c.xpath(vsys, []string{e.Name})

	// Create the objects.
	_, err = c.con.Edit(path, fn(e), nil, nil)
	return err
}

// Delete removes the given address objects from the firewall.
//
// Address objects can be either a string or an Entry object.
func (c *FwAddr) Delete(vsys string, e ...interface{}) error {
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
			return fmt.Errorf("Unsupported type to delete: %s", v)
		}
	}
	c.con.LogAction("(delete) address objects: %v", names)

	path := c.xpath(vsys, names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

/** Internal functions for the FwAddr struct **/

func (c *FwAddr) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{9, 0, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *FwAddr) details(fn util.Retriever, vsys, name string) ([]Entry, error) {
	path := c.xpath(vsys, []string{name})
	obj, _ := c.versioning()
	_, err := fn(path, nil, obj)
	if err != nil {
		return nil, err
	}
	ans := obj.Normalize()

	return ans, nil
}

func (c *FwAddr) xpath(vsys string, vals []string) []string {
	if vsys == "" {
		vsys = "vsys1"
	}

	if vsys == "shared" {
		return []string{
			"config",
			"shared",
			"address",
			util.AsEntryXpath(vals),
		}
	}

	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"vsys",
		util.AsEntryXpath([]string{vsys}),
		"address",
		util.AsEntryXpath(vals),
	}
}
