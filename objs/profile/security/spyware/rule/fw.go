package rule

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/namespace"
	"github.com/PaloAltoNetworks/pango/util"
)

// Firewall is the client.Objects.AntiSpywareRule namespace.
type Firewall struct {
	ns *namespace.Standard
}

// GetList performs GET to retrieve a list of all objects.
func (c *Firewall) GetList(vsys, prof string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(vsys, prof), ans)
}

// ShowList performs SHOW to retrieve a list of all objects.
func (c *Firewall) ShowList(vsys, prof string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(vsys, prof), ans)
}

// Get performs GET to retrieve information for the given object.
func (c *Firewall) Get(vsys, prof, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(vsys, prof), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve information for the given object.
func (c *Firewall) Show(vsys, prof, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(vsys, prof), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Firewall) GetAll(vsys, prof string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(vsys, prof), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve information for all objects.
func (c *Firewall) ShowAll(vsys, prof string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(vsys, prof), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Firewall) Set(vsys, prof string, e ...Entry) error {
	return c.ns.Set(c.pather(vsys, prof), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Firewall) Edit(vsys, prof string, e Entry) error {
	return c.ns.Edit(c.pather(vsys, prof), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Firewall) Delete(vsys, prof string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(vsys, prof), names, nErr)
}

// MoveGroup moves a logical group of anti spyware rules somewhere
// in relation to another rule.
//
// The `movement` param should be one of the Move constants in the util
// package.
//
// The `rule` param is the other rule the `movement` param is referencing.  If
// this is an empty string, then the first policy in the group isn't moved
// anywhere, but all other policies will still be moved to be grouped with the
// first one.
func (c *Firewall) MoveGroup(vsys, prof string, movement int, rule string, e ...Entry) error {
	lister := func() ([]string, error) {
		return c.GetList(vsys, prof)
	}

	ei := make([]interface{}, 0, len(e))
	for i := range e {
		ei = append(ei, e[i])
	}
	names, _ := toNames(ei)

	return c.ns.MoveGroup(c.pather(vsys, prof), lister, movement, rule, names)
}

func (c *Firewall) pather(vsys, prof string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(vsys, prof, v)
	}
}

func (c *Firewall) xpath(vsys, prof string, vals []string) ([]string, error) {
	if prof == "" {
		return nil, fmt.Errorf("prof must be specified")
	}

	ans := make([]string, 0, 10)
	ans = append(ans, util.VsysXpathPrefix(vsys)...)
	ans = append(ans,
		"profiles",
		"spyware",
		util.AsEntryXpath([]string{prof}),
		"rules",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Firewall) container() normalizer {
	return container(c.ns.Client.Versioning())
}
