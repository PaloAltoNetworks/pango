package nat

import (
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// PanoNat is the client.Policies.Nat namespace.
type PanoNat struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoed by client.Initialize().
func (c *PanoNat) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// GetList performs GET to retrieve a list of NAT policies.
func (c *PanoNat) GetList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(dg, base, nil), result)
}

// ShowList performs SHOW to retrieve a list of NAT policies.
func (c *PanoNat) ShowList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(dg, base, nil), result)
}

// Get performs GET to retrieve information for the given NAT policy.
func (c *PanoNat) Get(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs a GET to retrieve all objects.
func (c *PanoNat) GetAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given NAT policy.
func (c *PanoNat) Show(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs a SHOW to retrieve all objects.
func (c *PanoNat) ShowAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more NAT policies.
func (c *PanoNat) Set(dg, base string, e ...Entry) error {
	var err error

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

// Edit performs EDIT to create / update a NAT policy.
func (c *PanoNat) Edit(dg, base string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(dg, base, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// Delete removes the given NAT policies.
//
// NAT policies can be either a string or an Entry object.
func (c *PanoNat) Delete(dg, base string, e ...interface{}) error {
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

	path := c.xpath(dg, base, names)
	return c.ns.Delete(names, path)
}

// MoveGroup moves a logical group of rules somewhere in relation
// to another rule.
//
// The `movement` param should be one of the Move constants in the util
// package.
//
// The `rule` param is the other rule the `movement` param is referencing.  If
// this is an empty string, then the first rule in the group isn't moved
// anywhere, but all other rules will still be moved to be grouped with the
// first one.
func (c *PanoNat) MoveGroup(dg, base string, movement int, rule string, e ...Entry) error {
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

/** Internal functions **/

func (c *PanoNat) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *PanoNat) xpath(dg, base string, vals []string) []string {
	if dg == "" {
		dg = "shared"
	}
	if base == "" {
		base = util.PreRulebase
	}

	if dg == "shared" {
		return []string{
			"config",
			"shared",
			base,
			"nat",
			"rules",
			util.AsEntryXpath(vals),
		}
	}

	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"device-group",
		util.AsEntryXpath([]string{dg}),
		base,
		"nat",
		"rules",
		util.AsEntryXpath(vals),
	}
}
