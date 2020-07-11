package security

import (
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// PanoSecurity is the client.Policies.Security namespace.
type PanoSecurity struct {
	con util.XapiClient
	ns  *namespace.Namespace
}

// Initialize is invoed by client.Initialize().
func (c *PanoSecurity) Initialize(con util.XapiClient) {
	c.con = con
	c.ns = namespace.New(singular, plural, con)
}

// GetList performs GET to retrieve a list of object names.
func (c *PanoSecurity) GetList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Get, c.xpath(dg, base, nil), result)
}

// ShowList performs SHOW to retrieve a list of object names.
func (c *PanoSecurity) ShowList(dg, base string) ([]string, error) {
	result, _ := c.versioning()
	return c.ns.Listing(util.Show, c.xpath(dg, base, nil), result)
}

// Get performs GET to retrieve information for the given object.
func (c *PanoSecurity) Get(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Get, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// GetAll performs GET to retrieve all objects.
func (c *PanoSecurity) GetAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Get, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Show performs SHOW to retrieve information for the given object.
func (c *PanoSecurity) Show(dg, base, name string) (Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Object(util.Show, c.xpath(dg, base, []string{name}), name, result); err != nil {
		return Entry{}, err
	}

	return result.Normalize()[0], nil
}

// ShowAll performs SHOW to retrieve all objects.
func (c *PanoSecurity) ShowAll(dg, base string) ([]Entry, error) {
	result, _ := c.versioning()
	if err := c.ns.Objects(util.Show, c.xpath(dg, base, nil), result); err != nil {
		return nil, err
	}

	return result.Normalize(), nil
}

// Set performs SET to create / update one or more security policies.
func (c *PanoSecurity) Set(dg, base string, e ...Entry) error {
	var err error

	_, fn := c.versioning()
	data := make([]interface{}, 0, len(e))
	names := make([]string, len(e))

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

// VerifiableSet behaves like Set(), except policies with LogEnd as true
// will first be created with LogEnd as false, and then a second Set() is
// performed which will do LogEnd as true.  This is due to the unique
// combination of being a boolean value that is true by default, the XML
// returned from querying the rule details will omit the LogEnd setting,
// which will be interpreted as false, when in fact it is true.  We can
// get around this by setting the value to a non-standard value, then back
// again, in which case it will properly show up in the returned XML.
func (c *PanoSecurity) VerifiableSet(dg, base string, e ...Entry) error {
	c.con.LogAction("(set) performing verifiable set")
	again := make([]Entry, 0, len(e))

	for i := range e {
		if e[i].LogEnd {
			again = append(again, e[i])
			e[i].LogEnd = false
		}
	}

	if err := c.Set(dg, base, e...); err != nil {
		return err
	}

	if len(again) == 0 {
		return nil
	}

	return c.Set(dg, base, again...)
}

// Edit performs EDIT to create / update a security policy.
func (c *PanoSecurity) Edit(dg, base string, e Entry) error {
	_, fn := c.versioning()
	path := c.xpath(dg, base, []string{e.Name})
	data := fn(e)

	return c.ns.Edit(e.Name, path, data)
}

// VerifiableEdit behaves like Edit(), except policies with LogEnd as true
// will first be created with LogEnd as false, and then a second Set() is
// performed which will do LogEnd as true.  This is due to the unique
// combination of being a boolean value that is true by default, the XML
// returned from querying the rule details will omit the LogEnd setting,
// which will be interpreted as false, when in fact it is true.  We can
// get around this by setting the value to a non-standard value, then back
// again, in which case it will properly show up in the returned XML.
func (c *PanoSecurity) VerifiableEdit(dg, base string, e ...Entry) error {
	var err error

	c.con.LogAction("(edit) performing verifiable edit")
	again := make([]Entry, 0, len(e))

	for i := range e {
		if e[i].LogEnd {
			again = append(again, e[i])
			e[i].LogEnd = false
		}
		if err = c.Edit(dg, base, e[i]); err != nil {
			return err
		}
	}

	if len(again) == 0 {
		return nil
	}

	// It's ok to do a SET following an EDIT because we are guaranteed
	// to not have stray or conflicting config, so use SET since it
	// supports bulk operations.
	return c.Set(dg, base, again...)
}

// Delete removes the given security policies.
//
// Security policies can be either a string or an Entry object.
func (c *PanoSecurity) Delete(dg, base string, e ...interface{}) error {
	names := make([]string, 0, len(e))
	for i := range e {
		switch v := e[i].(type) {
		case string:
			names = append(names, v)
		case Entry:
			names = append(names, v.Name)
		default:
			return fmt.Errorf("Unsupported type to delete: %s", v)
		}
	}

	path := c.xpath(dg, base, names)
	return c.ns.Delete(names, path)
}

// DeleteAll removes all security policies from the specified dg / rulebase.
func (c *PanoSecurity) DeleteAll(dg, base string) error {
	c.con.LogAction("(delete) all security policies")
	list, err := c.GetList(dg, base)
	if err != nil || len(list) == 0 {
		return err
	}
	li := make([]interface{}, len(list))
	for i := range list {
		li[i] = list[i]
	}
	return c.Delete(dg, base, li...)
}

// MoveGroup moves a logical group of security policies somewhere in relation
// to another security policy.
//
// The `movement` param should be one of the Move constants in the util
// package.
//
// The `rule` param is the other rule the `movement` param is referencing.  If
// this is an empty string, then the first policy in the group isn't moved
// anywhere, but all other policies will still be moved to be grouped with the
// first one.
func (c *PanoSecurity) MoveGroup(dg, base string, movement int, rule string, e ...Entry) error {
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

/** Internal functions for the PanoSecurity struct **/

func (c *PanoSecurity) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *PanoSecurity) xpath(dg, base string, vals []string) []string {
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
			"security",
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
		"security",
		"rules",
		util.AsEntryXpath(vals),
	}
}
