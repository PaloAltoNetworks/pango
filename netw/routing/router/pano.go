package router

import (
    "fmt"
    "encoding/xml"

    "github.com/PaloAltoNetworks/pango/util"
)


// PanoRouter is the client.Network.VirtualRouter namespace.
type PanoRouter struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *PanoRouter) Initialize(con util.XapiClient) {
    c.con = con
}

// ShowList performs SHOW to retrieve a list of virtual routers.
func (c *PanoRouter) ShowList(tmpl string) ([]string, error) {
    c.con.LogQuery("(show) list of virtual routeres")
    path := c.xpath(tmpl, nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// GetList performs GET to retrieve a list of virtual routers.
func (c *PanoRouter) GetList(tmpl string) ([]string, error) {
    c.con.LogQuery("(get) list of virtual routers")
    path := c.xpath(tmpl, nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given virtual router.
func (c *PanoRouter) Get(tmpl, name string) (Entry, error) {
    c.con.LogQuery("(get) virtual router %q", name)
    return c.details(c.con.Get, tmpl, name)
}

// Show performs SHOW to retrieve information for the given virtual router.
func (c *PanoRouter) Show(tmpl, name string) (Entry, error) {
    c.con.LogQuery("(show) virtual router %q", name)
    return c.details(c.con.Show, tmpl, name)
}

// Set performs SET to create / update one or more virtual routers.
//
// Specify a non-empty vsys to import the virtual routers into the given vsys
// after creating, allowing the vsys to use them.
func (c *PanoRouter) Set(tmpl, vsys string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if tmpl == "" {
        return fmt.Errorf("tmpl must be specified")
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given router configs.
    d := util.BulkElement{XMLName: xml.Name{Local: "virtual-router"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) virtual routers: %v", names)

    // Set xpath.
    path := c.xpath(tmpl, names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the virtual routers.
    if _, err = c.con.Set(path, d.Config(), nil, nil); err != nil {
        return err
    }

    // Remove the virtual routers from any vsys they're currently in.
    if err = c.con.UnimportVirtualRouters(tmpl, names); err != nil {
        return err
    }

    // Perform vsys import next.
    return c.con.ImportVirtualRouters(tmpl, vsys, names)
}

// Edit performs EDIT to create / update a virtual router.
//
// Specify a non-empty vsys to import the virtual router into the given vsys
// after creating, allowing the vsys to use them.
func (c *PanoRouter) Edit(tmpl, vsys string, e Entry) error {
    var err error

    if tmpl == "" {
        return fmt.Errorf("tmpl must be specified")
    }

    _, fn := c.versioning()

    c.con.LogAction("(edit) virtual router %q", e.Name)

    // Set xpath.
    path := c.xpath(tmpl, []string{e.Name})

    // Edit the virtual router.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    if err != nil {
        return err
    }

    // Remove the virtual routers from any vsys they're currently in.
    if err = c.con.UnimportVirtualRouters(tmpl, []string{e.Name}); err != nil {
        return err
    }

    // Perform vsys import next.
    return c.con.ImportVirtualRouters(tmpl, vsys, []string{e.Name})
}

// Delete removes the given virtual routers from the firewall.
//
// Virtual routers can be a string or an Entry object.
func (c *PanoRouter) Delete(tmpl string, e ...interface{}) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if tmpl == "" {
        return fmt.Errorf("tmpl must be specified")
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
    c.con.LogAction("(delete) virtual routers: %v", names)

    // Unimport virtual routers.
    err = c.con.UnimportVirtualRouters(tmpl, names)
    if err != nil {
        return err
    }

    // Remove virtual routers next.
    path := c.xpath(tmpl, names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for this namespace struct **/

func (c *PanoRouter) versioning() (normalizer, func(Entry) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *PanoRouter) details(fn util.Retriever, tmpl, name string) (Entry, error) {
    path := c.xpath(tmpl, []string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *PanoRouter) xpath(tmpl string, vals []string) []string {
    ans := make([]string, 0, 11)
    ans = append(ans, util.TemplateXpathPrefix(tmpl)...)
    ans = append(ans,
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "network",
        "virtual-router",
        util.AsEntryXpath(vals),
    )

    return ans
}
