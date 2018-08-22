package stack

import (
    "fmt"
    "encoding/xml"

    "github.com/PaloAltoNetworks/pango/util"
)

// Stack is the client.Panorama.TemplateStack namespace.
type Stack struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *Stack) Initialize(con util.XapiClient) {
    c.con = con
}

/*
SetDeviceVsys performs a SET to add specific vsys from a device to
template stack st.

If you want all vsys to be included, or the device is a virtual firewall, then
leave the vsys list empty.

The template stack can be either a string or an Entry object.
*/
func (c *Stack) SetDeviceVsys(st interface{}, d string, vsys []string) error {
    var name string

    switch v := st.(type) {
    case string:
        name = v
    case Entry:
        name = v.Name
    default:
        return fmt.Errorf("Unknown type sent to set device vsys: %s", v)
    }

    c.con.LogAction("(set) device vsys in template stack: %s", name)

    m := util.MapToVsysEnt(map[string] []string{d: vsys})
    path := c.xpath([]string{name})
    path = append(path, "devices")

    _, err := c.con.Set(path, m.Entries[0], nil, nil)
    return err
}

/*
EditDeviceVsys performs an EDIT to add specific vsys from a device to
template stack st.

If you want all vsys to be included, or the device is a virtual firewall, then
leave the vsys list empty.

The template stack can be either a string or an Entry object.
*/
func (c *Stack) EditDeviceVsys(st interface{}, d string, vsys []string) error {
    var name string

    switch v := st.(type) {
    case string:
        name = v
    case Entry:
        name = v.Name
    default:
        return fmt.Errorf("Unknown type sent to edit device vsys: %s", v)
    }

    c.con.LogAction("(edit) device vsys in template stack: %s", name)

    m := util.MapToVsysEnt(map[string] []string{d: vsys})
    path := c.xpath([]string{name})
    path = append(path, "devices", util.AsEntryXpath([]string{d}))

    _, err := c.con.Edit(path, m.Entries[0], nil, nil)
    return err
}

/*
DeleteDeviceVsys performs a DELETE to remove specific vsys from device d from
template stack st.

If you want all vsys to be removed, or the device is a virtual firewall, then
leave the vsys list empty.

The template stack can be either a string or an Entry object.
*/
func (c *Stack) DeleteDeviceVsys(st interface{}, d string, vsys []string) error {
    var name string

    switch v := st.(type) {
    case string:
        name = v
    case Entry:
        name = v.Name
    default:
        return fmt.Errorf("Unknown type sent to delete device vsys: %s", v)
    }

    c.con.LogAction("(delete) device vsys from template stack: %s", name)

    path := make([]string, 0, 9)
    path = append(path, c.xpath([]string{name})...)
    path = append(path, "devices", util.AsEntryXpath([]string{d}))
    if len(vsys) > 0 {
        path = append(path, "vsys", util.AsEntryXpath(vsys))
    }

    _, err := c.con.Delete(path, nil, nil)
    return err
}

// ShowList performs SHOW to retrieve a list of template stacks.
func (c *Stack) ShowList() ([]string, error) {
    c.con.LogQuery("(show) list of template stacks")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// GetList performs GET to retrieve a list of template stacks.
func (c *Stack) GetList() ([]string, error) {
    c.con.LogQuery("(get) list of template stacks")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given template stack.
func (c *Stack) Get(name string) (Entry, error) {
    c.con.LogQuery("(get) template stack %q", name)
    return c.details(c.con.Get, name)
}

// Show performs SHOW to retrieve information for the given template stack.
func (c *Stack) Show(name string) (Entry, error) {
    c.con.LogQuery("(show) template stack %q", name)
    return c.details(c.con.Show, name)
}

// Set performs SET to create / update one or more template stacks.
func (c *Stack) Set(e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given configs.
    d := util.BulkElement{XMLName: xml.Name{Local: "template-stack"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) template stacks: %v", names)

    // Set xpath.
    path := c.xpath(names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the template stacks.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit performs EDIT to create / update a template stack.
func (c *Stack) Edit(e Entry) error {
    var err error

    _, fn := c.versioning()

    c.con.LogAction("(edit) template stack %q", e.Name)

    // Set xpath.
    path := c.xpath([]string{e.Name})

    // Edit the template stack.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given template stacks from the firewall.
//
// Objects can be a string or an Entry object.
func (c *Stack) Delete(e ...interface{}) error {
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
            return fmt.Errorf("Unknown type sent to delete: %s", v)
        }
    }
    c.con.LogAction("(delete) template stacks: %v", names)

    // Remove the template stacks.
    path := c.xpath(names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for this namespace struct **/

func (c *Stack) versioning() (normalizer, func(Entry) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *Stack) details(fn util.Retriever, name string) (Entry, error) {
    path := c.xpath([]string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *Stack) xpath(vals []string) []string {
    return []string{
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "template-stack",
        util.AsEntryXpath(vals),
    }
}
