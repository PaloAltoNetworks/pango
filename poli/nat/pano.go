package nat

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// PanoNat is the client.Policies.Nat namespace.
type PanoNat struct {
	con util.XapiClient
}

// Initialize is invoed by client.Initialize().
func (c *PanoNat) Initialize(con util.XapiClient) {
	c.con = con
}

// GetList performs GET to retrieve a list of NAT policies.
func (c *PanoNat) GetList(dg, base string) ([]string, error) {
	c.con.LogQuery("(get) list of %s", plural)
	path := c.xpath(dg, base, nil)
	return c.con.EntryListUsing(c.con.Get, path[:len(path)-1])
}

// GetFullList performs GET to retreive a list of security policies
func (c *PanoNat) GetFullList(dg, base string) (EntryContainerMulti, error) {
	c.con.LogQuery("(get) list of service objects")
	path := c.xpath(dg, base, nil)
	ans := EntryContainerMulti{}
	err := c.con.EntryObjectsUsing(c.con.Get, path[:len(path)-1], &ans)
	if err != nil {
		return EntryContainerMulti{}, err
	}
	return ans, nil
}

// ShowList performs SHOW to retrieve a list of security policies.
func (c *PanoNat) ShowList(dg, base string) ([]string, error) {
	c.con.LogQuery("(show) list of security policies")
	path := c.xpath(dg, base, nil)
	return c.con.EntryListUsing(c.con.Show, path[:len(path)-1])
}

// ShowFullList performs SHOW to retreive a list of security policies
func (c *PanoNat) ShowFullList(dg, base string) (EntryContainerMulti, error) {
	c.con.LogQuery("(show) list of service objects")
	path := c.xpath(dg, base, nil)
	ans := EntryContainerMulti{}
	err := c.con.EntryObjectsUsing(c.con.Get, path[:len(path)-1], &ans)
	if err != nil {
		return EntryContainerMulti{}, err
	}
	return ans, nil
}

// Get performs GET to retrieve information for the given NAT policy.
func (c *PanoNat) Get(dg, base, name string) (Entry, error) {
	c.con.LogQuery("(get) %s %q", singular, name)
	return c.details(c.con.Get, dg, base, name)
}

// Get performs SHOW to retrieve information for the given NAT policy.
func (c *PanoNat) Show(dg, base, name string) (Entry, error) {
	c.con.LogQuery("(show) %s %q", singular, name)
	return c.details(c.con.Show, dg, base, name)
}

// Set performs SET to create / update one or more NAT policies.
func (c *PanoNat) Set(dg, base string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	} else {
		// Make sure rule names are unique.
		m := make(map[string]int)
		for i := range e {
			m[e[i].Name] = m[e[i].Name] + 1
			if m[e[i].Name] > 1 {
				return fmt.Errorf("NAT rule is defined multiple times: %s", e[i].Name)
			}
		}
	}

	_, fn := c.versioning()
	names := make([]string, len(e))

	// Build up the struct with the given configs.
	d := util.BulkElement{XMLName: xml.Name{Local: "rules"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
	}
	c.con.LogAction("(set) %s: %v", plural, names)

	// Set xpath.
	path := c.xpath(dg, base, names)
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the NAT policies.
	_, err = c.con.Set(path, d.Config(), nil, nil)

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
	var err error

	_, fn := c.versioning()

	c.con.LogAction("(edit) %s %q", singular, e.Name)

	// Set xpath.
	path := c.xpath(dg, base, []string{e.Name})

	// Edit the NAT policy.
	_, err = c.con.Edit(path, fn(e), nil, nil)
	return err
}

// Delete removes the given NAT policies.
//
// NAT policies can be either a string or an Entry object.
func (c *PanoNat) Delete(dg, base string, e ...interface{}) error {
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
	c.con.LogAction("(delete) %s: %v", plural, names)

	path := c.xpath(dg, base, names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

// MoveGroup moves a logical group of NAT rules somewhere in relation
// to another rule.
func (c *PanoNat) MoveGroup(dg, base string, mvt int, rule string, e ...Entry) error {
	var err error

	c.con.LogAction("(move) %s group", singular)

	if len(e) < 1 {
		return fmt.Errorf("Requires at least one rule")
	}

	path := c.xpath(dg, base, []string{e[0].Name})
	list, err := c.GetList(dg, base)
	if err != nil {
		return err
	}

	// Set the first entity's position.
	if err = c.con.PositionFirstEntity(mvt, rule, e[0].Name, path, list); err != nil {
		return err
	}

	// Move all the rest under it.
	li := len(path) - 1
	for i := 1; i < len(e); i++ {
		path[li] = util.AsEntryXpath([]string{e[i].Name})
		if _, err = c.con.Move(path, "after", e[i-1].Name, nil, nil); err != nil {
			return err
		}
	}

	return nil
}

/** Internal functions for the Zone struct **/

func (c *PanoNat) versioning() (normalizer, func(Entry) interface{}) {
	v := c.con.Versioning()

	if v.Gte(version.Number{8, 1, 0, ""}) {
		return &container_v2{}, specify_v2
	} else {
		return &container_v1{}, specify_v1
	}
}

func (c *PanoNat) details(fn util.Retriever, dg, base, name string) (Entry, error) {
	path := c.xpath(dg, base, []string{name})
	obj, _ := c.versioning()
	if _, err := fn(path, nil, obj); err != nil {
		return Entry{}, err
	}
	ans := obj.Normalize()

	return ans, nil
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
