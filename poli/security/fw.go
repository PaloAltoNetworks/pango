package security

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/util"
)

// FwSecurity is the client.Policies.Security namespace.
type FwSecurity struct {
	con util.XapiClient
}

// Initialize is invoed by client.Initialize().
func (c *FwSecurity) Initialize(con util.XapiClient) {
	c.con = con
}

// GetList performs GET to retrieve a list of security policies.
func (c *FwSecurity) GetList(vsys string) ([]string, error) {
	c.con.LogQuery("(get) list of security policies")
	path := c.xpath(vsys, nil)
	return c.con.EntryListUsing(c.con.Get, path[:len(path)-1])
}

// ShowList performs SHOW to retrieve a list of security policies.
func (c *FwSecurity) ShowList(vsys string) ([]string, error) {
	c.con.LogQuery("(show) list of security policies")
	path := c.xpath(vsys, nil)
	return c.con.EntryListUsing(c.con.Show, path[:len(path)-1])
}

// Get performs GET to retrieve information for the given security policy.
func (c *FwSecurity) Get(vsys, name string) (Entry, error) {
	c.con.LogQuery("(get) security policy %q", name)
	listing, err := c.details(c.con.Get, vsys, name)
	if err == nil && len(listing) > 0 {
		return listing[0], nil
	}
	return Entry{}, err
}

// GetAll performs a GET to retrieve information for all security policies.
func (c *FwSecurity) GetAll(vsys string) ([]Entry, error) {
	c.con.LogQuery("(get) all security policies")
	return c.details(c.con.Get, vsys, "")
}

// Get performs SHOW to retrieve information for the given security policy.
func (c *FwSecurity) Show(vsys, name string) (Entry, error) {
	c.con.LogQuery("(show) security policy %q", name)
	listing, err := c.details(c.con.Show, vsys, name)
	if err == nil && len(listing) > 0 {
		return listing[0], nil
	}
	return Entry{}, err
}

// ShowAll performs a GET to retrieve information for all security policies.
func (c *FwSecurity) ShowAll(vsys string) ([]Entry, error) {
	c.con.LogQuery("(show) all security policies")
	return c.details(c.con.Show, vsys, "")
}

// Set performs SET to create / update one or more security policies.
func (c *FwSecurity) Set(vsys string, e ...Entry) error {
	var err error

	if len(e) == 0 {
		return nil
	} else {
		// Make sure rule names are unique.
		m := make(map[string]int)
		for i := range e {
			m[e[i].Name] = m[e[i].Name] + 1
			if m[e[i].Name] > 1 {
				return fmt.Errorf("Security rule is defined multiple times: %s", e[i].Name)
			}
		}
	}

	_, fn := c.versioning()
	names := make([]string, len(e))

	// Build up the struct with the given security policy configs.
	d := util.BulkElement{XMLName: xml.Name{Local: "rules"}}
	for i := range e {
		d.Data = append(d.Data, fn(e[i]))
		names[i] = e[i].Name
	}
	c.con.LogAction("(set) security policies: %v", names)

	// Set xpath.
	path := c.xpath(vsys, names)
	if len(e) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	// Create the security policies.
	_, err = c.con.Set(path, d.Config(), nil, nil)

	// On error: find the rule that's causing the error if multiple rules
	// were given.
	if err != nil && strings.Contains(err.Error(), "rules is invalid") {
		for i := 0; i < len(e); i++ {
			if e2 := c.Set(vsys, e[i]); e2 != nil {
				return fmt.Errorf("Error with rule %d: %s", i+1, e2)
			} else {
				_ = c.Delete(vsys, e[i])
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
func (c *FwSecurity) VerifiableSet(vsys string, e ...Entry) error {
	c.con.LogAction("(set) performing verifiable set")
	again := make([]Entry, 0, len(e))

	for i := range e {
		if e[i].LogEnd {
			again = append(again, e[i])
			e[i].LogEnd = false
		}
	}

	if err := c.Set(vsys, e...); err != nil {
		return err
	}

	if len(again) == 0 {
		return nil
	}

	return c.Set(vsys, again...)
}

// Edit performs EDIT to create / update a security policy.
func (c *FwSecurity) Edit(vsys string, e Entry) error {
	var err error

	_, fn := c.versioning()

	c.con.LogAction("(edit) security policy %q", e.Name)

	// Set xpath.
	path := c.xpath(vsys, []string{e.Name})

	// Edit the security policy.
	_, err = c.con.Edit(path, fn(e), nil, nil)
	return err
}

// VerifiableEdit behaves like Edit(), except policies with LogEnd as true
// will first be created with LogEnd as false, and then a second Set() is
// performed which will do LogEnd as true.  This is due to the unique
// combination of being a boolean value that is true by default, the XML
// returned from querying the rule details will omit the LogEnd setting,
// which will be interpreted as false, when in fact it is true.  We can
// get around this by setting the value to a non-standard value, then back
// again, in which case it will properly show up in the returned XML.
func (c *FwSecurity) VerifiableEdit(vsys string, e ...Entry) error {
	var err error

	c.con.LogAction("(edit) performing verifiable edit")
	again := make([]Entry, 0, len(e))

	for i := range e {
		if e[i].LogEnd {
			again = append(again, e[i])
			e[i].LogEnd = false
		}
		if err = c.Edit(vsys, e[i]); err != nil {
			return err
		}
	}

	if len(again) == 0 {
		return nil
	}

	// It's ok to do a SET following an EDIT because we are guaranteed
	// to not have stray or conflicting config, so use SET since it
	// supports bulk operations.
	return c.Set(vsys, again...)
}

// Delete removes the given security policies.
//
// Security policies can be either a string or an Entry object.
func (c *FwSecurity) Delete(vsys string, e ...interface{}) error {
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
	c.con.LogAction("(delete) security policies: %v", names)

	path := c.xpath(vsys, names)
	_, err = c.con.Delete(path, nil, nil)
	return err
}

// DeleteAll removes all security policies from the specified vsys.
func (c *FwSecurity) DeleteAll(vsys string) error {
	c.con.LogAction("(delete) all security policies")
	list, err := c.GetList(vsys)
	if err != nil || len(list) == 0 {
		return err
	}
	li := make([]interface{}, len(list))
	for i := range list {
		li[i] = list[i]
	}
	return c.Delete(vsys, li...)
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
func (c *FwSecurity) MoveGroup(vsys string, movement int, rule string, e ...Entry) error {
	var err error

	fIdx := -1
	oIdx := -1

	c.con.LogAction("(move) security policy group")
	if len(e) < 1 {
		return fmt.Errorf("Requires at least one security policy")
	} else if rule == e[0].Name {
		return fmt.Errorf("Can't position %q in relation to itself", rule)
	} else if movement != util.MoveSkip && movement != util.MoveBefore && movement != util.MoveDirectlyBefore && movement != util.MoveAfter && movement != util.MoveDirectlyAfter && movement != util.MoveTop && movement != util.MoveBottom {
		return fmt.Errorf("Invalid position specified: %d", movement)
	} else if (movement == util.MoveBefore || movement == util.MoveDirectlyBefore || movement == util.MoveAfter || movement == util.MoveDirectlyAfter) && rule == "" {
		return fmt.Errorf("Specify 'rule' in order to perform relative group positioning")
	}
	path := c.xpath(vsys, []string{e[0].Name})

	if movement != util.MoveSkip {
		// Get the list of current security policies.
		curList, err := c.GetList(vsys)
		if err != nil {
			return err
		} else if len(curList) == 0 {
			return fmt.Errorf("No policies found")
		}

		switch movement {
		case util.MoveTop:
			_, em := c.con.Move(path, "top", "", nil, nil)
			if em != nil {
				if em.Error() != "already at the top" {
					err = em
				}
			}
		case util.MoveBottom:
			_, em := c.con.Move(path, "bottom", "", nil, nil)
			if em != nil {
				if em.Error() != "already at the bottom" {
					err = em
				}
			}
		default:
			// Find the indexes of the first security policy and the ref policy.
			for i, v := range curList {
				if v == e[0].Name {
					fIdx = i
				} else if v == rule {
					oIdx = i
				}
				if fIdx != -1 && oIdx != -1 {
					break
				}
			}

			// Sanity check:  both rules should be present.
			if fIdx == -1 {
				return fmt.Errorf("First security policy in group %q does not exist", e[0].Name)
			} else if oIdx == -1 {
				return fmt.Errorf("Reference security policy %q does not exist", rule)
			}

			// Perform the move of the first security policy, if needed.
			if (movement == util.MoveBefore && fIdx > oIdx) || (movement == util.MoveDirectlyBefore && fIdx+1 != oIdx) {
				_, err = c.con.Move(path, "before", rule, nil, nil)
			} else if (movement == util.MoveAfter && fIdx < oIdx) || (movement == util.MoveDirectlyAfter && fIdx != oIdx+1) {
				_, err = c.con.Move(path, "after", rule, nil, nil)
			}
		}

		// If we moved something, make sure it worked.
		if err != nil {
			return err
		}
	}

	// Now move all other rules under the first.
	li := len(path) - 1
	for i := 1; i < len(e); i++ {
		path[li] = util.AsEntryXpath([]string{e[i].Name})
		_, err = c.con.Move(path, "after", e[i-1].Name, nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

/** Internal functions for the FwSecurity struct **/

func (c *FwSecurity) versioning() (normalizer, func(Entry) interface{}) {
	return &container_v1{}, specify_v1
}

func (c *FwSecurity) details(fn util.Retriever, vsys, name string) ([]Entry, error) {
	path := c.xpath(vsys, []string{name})
	obj, _ := c.versioning()
	if _, err := fn(path, nil, obj); err != nil {
		return []Entry{}, err
	}
	ans := obj.Normalize()

	return ans, nil
}

func (c *FwSecurity) xpath(vsys string, vals []string) []string {
	if vsys == "" {
		vsys = "vsys1"
	}

	if vsys == "shared" {
		return []string{
			"config",
			"shared",
			"rulebase",
			"security",
			"rules",
			util.AsEntryXpath(vals),
		}
	}

	return []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"vsys",
		util.AsEntryXpath([]string{vsys}),
		"rulebase",
		"security",
		"rules",
		util.AsEntryXpath(vals),
	}
}
