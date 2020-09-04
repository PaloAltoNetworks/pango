package pbf

import (
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// PanoPbf is the client.Policies.PolicyBasedForwarding namespace.
type PanoPbf struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoked by client.Initialize().
func (c *PanoPbf) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// ShowList performs SHOW to retrieve a list of values.
func (c *PanoPbf) ShowList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(dg, base, nil), result)
}

// GetList performs GET to retrieve a list of values.
func (c *PanoPbf) GetList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(dg, base, nil), result)
}

// Get performs GET to retrieve information for the given uid.
func (c *PanoPbf) Get(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve information for all objects.
func (c *PanoPbf) GetAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given uid.
func (c *PanoPbf) Show(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *PanoPbf) ShowAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more objects.
func (c *PanoPbf) Set(dg, base string, e ...Entry) error {
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
	path := c.xpath(dg, base, names)

	err = c.ns.Set(names, path, data)

	// On error: find the rule that's causing the error if multiple rules
	// were given.
	if err != nil && strings.Contains(err.Error(), "rules is invalid") {
		for i := 0; i < len(e); i++ {
			if e2 := c.Set(dg, base, e[i]); e2 != nil {
				return fmt.Errorf("Error with rule %d: %s", i+1, e2)
			} else {
				_ = c.Delete(dg, base, e[i])
			}
		}

		// Couldn't find it, just return the original error.
		return err
	}

	return err
}

// Edit performs EDIT to create / update one object.
func (c *PanoPbf) Edit(dg, base string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(dg, base, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *PanoPbf) Delete(dg, base string, e ...interface{}) error {
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
	path := c.xpath(dg, base, names)

	return c.ns.Delete(names, path)
}

// MoveGroup moves a logical group of policy based forwarding rules
// somewhere in relation to another rule.
//
// The `movement` param should be one of the Move constants in the util
// package.
//
// The `rule` param is the other rule the `movement` param is referencing.  If
// this is an empty string, then the first policy in the group isn't moved
// anywhere, but all other policies will still be moved to be grouped with the
// first one.
func (c *PanoPbf) MoveGroup(dg, base string, movement int, rule string, e ...Entry) error {
	pather := func(v string) []string {
		return c.xpath(dg, base, []string{v})
	}

	lister := func() ([]string, error) {
		return c.GetList(dg, base)
	}

	names := make([]string, 0, len(e))
	for i := range e {
		names = append(names, e[i].Name)
	}

	return c.ns.MoveGroup(pather, lister, movement, rule, names)
}

/** Internal functions for this namespace struct **/

func (c *PanoPbf) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{9, 0, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *PanoPbf) xpath(dg, base string, vals []string) []string {
	ans := make([]string, 0, 9)
	ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
	ans = append(ans,
		base,
		"pbf",
		"rules",
		util.AsEntryXpath(vals),
	)

	return ans
}
