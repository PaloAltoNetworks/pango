package arp

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// PanoArp is the client.Policies.Arp namespace.
type PanoArp struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoed by client.Initialize().
func (c *PanoArp) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// GetList performs GET to retrieve a list object names.
func (c *PanoArp) GetList(tmpl, ts, iType, iName, subName string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(tmpl, ts, iType, iName, subName, nil), result)
}

// ShowList performs SHOW to retrieve a list of object names.
func (c *PanoArp) ShowList(tmpl, ts, iType, iName, subName string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(tmpl, ts, iType, iName, subName, nil), result)
}

// Get performs GET to retrieve information for the given object.
func (c *PanoArp) Get(tmpl, ts, iType, iName, subName, ip string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(tmpl, ts, iType, iName, subName, []string{ip}), ip, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs a GET to retrieve information for all objects.
func (c *PanoArp) GetAll(tmpl, ts, iType, iName, subName string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(tmpl, ts, iType, iName, subName, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given object.
func (c *PanoArp) Show(tmpl, ts, iType, iName, subName, ip string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(tmpl, ts, iType, iName, subName, []string{ip}), ip, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs a SHOW to retrieve information for all objects.
func (c *PanoArp) ShowAll(tmpl, ts, iType, iName, subName string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(tmpl, ts, iType, iName, subName, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *PanoArp) Set(tmpl, ts, iType, iName, subName string, e ...Entry) error {
	var err error

	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, 0, len(e))

	for i := range e {
		data = append(data, fn(e[i]))
		names = append(names, e[i].Ip)
	}
	path := c.xpath(tmpl, ts, iType, iName, subName, names)

	err = c.ns.Set(names, path, data)

	return err
}

// Edit performs EDIT to create / update an object.
func (c *PanoArp) Edit(tmpl, ts, iType, iName, subName string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(tmpl, ts, iType, iName, subName, []string{e.Ip})
	data := fn(e)

	return c.ns.Edit(e.Ip, path, data)
}

// Delete removes the given objects.
//
// Objects can be either a string or an Entry object.
func (c *PanoArp) Delete(tmpl, ts, iType, iName, subName string, e ...interface{}) error {
	names := make([]string, 0, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names = append(names, v)
		case Entry:
			names = append(names, v.Ip)
		default:
			return fmt.Errorf("Unsupported type to delete: %s", v)
		}
	}

	path := c.xpath(tmpl, ts, iType, iName, subName, names)
	return c.ns.Delete(names, path)
}

/** Internal functions for the PanoArp struct **/

func (c *PanoArp) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *PanoArp) xpath(tmpl, ts, iType, iName, subName string, vals []string) []string {
	ans := make([]string, 0, 17)

	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"network",
		"interface",
		iType,
	)

	if iType != TypeVlan {
		ans = append(ans, util.AsEntryXpath([]string{iName}), "layer3")
	}

	if subName != "" {
		ans = append(ans, "units", util.AsEntryXpath([]string{subName}))
	}

	ans = append(ans, "arp", util.AsEntryXpath(vals))

	return ans
}
