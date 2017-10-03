// Package security is the client.Policies.Security namespace.
//
// Normalized object:  Entry
package security

import (
    "fmt"
    "encoding/xml"

    "github.com/PaloAltoNetworks/xapi/util"
)


// Entry is a normalized, version independent representation of a security
// rule.
type Entry struct {
    Name string
    Type string
    Description string
    Tags []string
    From []string
    To []string
    Source []string
    NegateSource bool
    SourceUser []string
    HipProfile []string
    Destination []string
    NegateDestination bool
    Application []string
    Service []string
    Category []string
    Action string
    LogSetting string
    LogStart bool
    LogEnd bool
    Disabled bool
    Schedule string
    IcmpUnreachable string
    DisableServerResponseInspection bool
    Group []string
    Target []string
    NegateTarget bool
    Virus []string
    Spyware []string
    Vulnerability []string
    UrlFiltering []string
    FileBlocking []string
    WildFireAnalysis []string
    DataFiltering []string
}

// Defaults sets params with uninitialized values to their GUI default setting.
//
// The defaults are as follows:
//      * Type: "universal"
//      * From: ["any"]
//      * To: ["any"]
//      * Source: ["any"]
//      * SourceUser: ["any"]
//      * HipProfile: ["any"]
//      * Destination: ["any"]
//      * Application: ["any"]
//      * Service: ["application-default"]
//      * Category: ["any"]
//      * LogEnd: true
func (e *Entry) Defaults() {
    if e.Type == "" {
        e.Type = "universal"
    }

    if len(e.From) == 0 {
        e.From = []string{"any"}
    }

    if len(e.To) == 0 {
        e.To = []string{"any"}
    }

    if len(e.Source) == 0 {
        e.Source = []string{"any"}
    }

    if len(e.SourceUser) == 0 {
        e.SourceUser = []string{"any"}
    }

    if len(e.HipProfile) == 0 {
        e.HipProfile = []string{"any"}
    }

    if len(e.Destination) == 0 {
        e.Destination = []string{"any"}
    }

    if len(e.Application) == 0 {
        e.Application = []string{"any"}
    }

    if len(e.Service) == 0 {
        e.Service = []string{"application-default"}
    }

    if len(e.Category) == 0 {
        e.Category = []string{"any"}
    }

    if !e.LogEnd {
        e.LogEnd = true
    }
}

// Security is the client.Policies.Security namespace.
type Security struct {
    con util.XapiClient
}

// Initialize is invoed by client.Initialize().
func (c *Security) Initialize(con util.XapiClient) {
    c.con = con
}

// GetList performs GET to retrieve a list of security policies.
func (c *Security) GetList(vsys, base string) ([]string, error) {
    c.con.LogQuery("(get) list of security policies")
    path := c.xpath(vsys, base, nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// ShowList performs SHOW to retrieve a list of security policies.
func (c *Security) ShowList(vsys, base string) ([]string, error) {
    c.con.LogQuery("(show) list of security policies")
    path := c.xpath(vsys, base, nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given security policy.
func (c *Security) Get(vsys, base, name string) (Entry, error) {
    c.con.LogQuery("(get) security policy %q", name)
    return c.details(c.con.Get, vsys, base, name)
}

// Get performs SHOW to retrieve information for the given security policy.
func (c *Security) Show(vsys, base, name string) (Entry, error) {
    c.con.LogQuery("(show) security policy %q", name)
    return c.details(c.con.Show, vsys, base, name)
}

// Set creates / updates one or more security policies.
func (c *Security) Set(vsys, base string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
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
    path := c.xpath(vsys, base, names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the security policies.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit creates / updates a security policy.
func (c *Security) Edit(vsys, base string, e Entry) error {
    var err error

    _, fn := c.versioning()

    c.con.LogAction("(edit) security policy %q", e.Name)

    // Set xpath.
    path := c.xpath(vsys, base, []string{e.Name})

    // Edit the security policy.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given security policies.
//
// Security policies can be either a string or an Entry object.
func (c *Security) Delete(vsys, base string, e ...interface{}) error {
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

    path := c.xpath(vsys, base, names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for the Zone struct **/

func (c *Security) versioning() (normalizer, func(Entry) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *Security) details(fn func(interface{}, interface{}, interface{}) ([]byte, error), vsys, base, name string) (Entry, error) {
    path := c.xpath(vsys, base, []string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *Security) xpath(vsys, base string, vals []string) []string {
    if vsys == "" {
        vsys = "vsys1"
    }
    if base == "" {
        base = util.Rulebase
    }

    return []string{
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "vsys",
        util.AsEntryXpath([]string{vsys}),
        base,
        "security",
        "rules",
        util.AsEntryXpath(vals),
    }
}

/** Structs / functions for this namespace. **/

type normalizer interface {
    Normalize() Entry
}

type container_v1 struct {
    Answer entry_v1 `xml:"result>entry"`
}

func (o *container_v1) Normalize() Entry {
    ans := Entry{
        Name: o.Answer.Name,
        Type: o.Answer.Type,
        Description: o.Answer.Description,
        Tags: util.MemToStr(o.Answer.Tags),
        From: util.MemToStr(o.Answer.From),
        To: util.MemToStr(o.Answer.To),
        Source: util.MemToStr(o.Answer.Source),
        NegateSource: util.AsBool(o.Answer.NegateSource),
        SourceUser: util.MemToStr(o.Answer.SourceUser),
        HipProfile: util.MemToStr(o.Answer.HipProfile),
        Destination: util.MemToStr(o.Answer.Destination),
        NegateDestination: util.AsBool(o.Answer.NegateDestination),
        Application: util.MemToStr(o.Answer.Application),
        Service: util.MemToStr(o.Answer.Service),
        Category: util.MemToStr(o.Answer.Category),
        Action: o.Answer.Action,
        LogSetting: o.Answer.LogSetting,
        LogStart: util.AsBool(o.Answer.LogStart),
        LogEnd: util.AsBool(o.Answer.LogEnd),
        Disabled: util.AsBool(o.Answer.Disabled),
        Schedule: o.Answer.Schedule,
        IcmpUnreachable: o.Answer.IcmpUnreachable,
    }
    if o.Answer.Options != nil {
        ans.DisableServerResponseInspection = util.AsBool(o.Answer.Options.DisableServerResponseInspection)
    }
    if o.Answer.TargetInfo != nil {
        ans.NegateTarget = util.AsBool(o.Answer.TargetInfo.NegateTarget)
        ans.Target = util.EntToStr(o.Answer.TargetInfo.Target)
    }
    if o.Answer.ProfileSettings != nil {
        ans.Group = util.MemToStr(o.Answer.ProfileSettings.Group)
        if o.Answer.ProfileSettings.Profiles != nil {
            ans.Virus = util.MemToStr(o.Answer.ProfileSettings.Profiles.Virus)
            ans.Spyware = util.MemToStr(o.Answer.ProfileSettings.Profiles.Spyware)
            ans.Vulnerability = util.MemToStr(o.Answer.ProfileSettings.Profiles.Vulnerability)
            ans.UrlFiltering = util.MemToStr(o.Answer.ProfileSettings.Profiles.UrlFiltering)
            ans.FileBlocking = util.MemToStr(o.Answer.ProfileSettings.Profiles.FileBlocking)
            ans.WildFireAnalysis = util.MemToStr(o.Answer.ProfileSettings.Profiles.WildFireAnalysis)
            ans.DataFiltering = util.MemToStr(o.Answer.ProfileSettings.Profiles.DataFiltering)
        }
    }

    return ans
}

type entry_v1 struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    Type string `xml:"rule-type"`
    Description string `xml:"description"`
    Tags *util.Member `xml:"tag"`
    From *util.Member `xml:"from"`
    To *util.Member `xml:"to"`
    Source *util.Member `xml:"source"`
    NegateSource string `xml:"negate-source"`
    SourceUser *util.Member `xml:"source-user"`
    HipProfile *util.Member `xml:"hip-profiles"`
    Destination *util.Member `xml:"destination"`
    NegateDestination string `xml:"negate-destination"`
    Application *util.Member `xml:"application"`
    Service *util.Member `xml:"service"`
    Category *util.Member `xml:"category"`
    Action string `xml:"action"`
    LogSetting string `xml:"log-setting,omitempty"`
    LogStart string `xml:"log-start"`
    LogEnd string `xml:"log-end"`
    Disabled string `xml:"disabled"`
    Schedule string `xml:"schedule,omitempty"`
    IcmpUnreachable string `xml:"icmp-unreachable,omitempty"`
    Options *secOptions `xml:"option"`
    TargetInfo *targetInfo `xml:"target"`
    ProfileSettings *profileSettings `xml:"profile-setting"`
}

type secOptions struct {
    DisableServerResponseInspection string `xml:"disable-server-response-inspection,omitempty"`
}

type targetInfo struct {
    Target *util.Entry `xml:"devices"`
    NegateTarget string `xml:"negate,omitempty"`
}

type profileSettings struct {
    Group *util.Member `xml:"group"`
    Profiles *profileSettingsProfile `xml:"profiles"`
}

type profileSettingsProfile struct {
    Virus *util.Member `xml:"virus"`
    Spyware *util.Member `xml:"spyware"`
    Vulnerability *util.Member `xml:"vulnerability"`
    UrlFiltering *util.Member `xml:"url-filtering"`
    FileBlocking *util.Member `xml:"file-blocking"`
    WildFireAnalysis *util.Member `xml:"wildfire-analysis"`
    DataFiltering *util.Member `xml:"data-filtering"`
}

func specify_v1(e Entry) interface{} {
    ans := entry_v1{
        Name: e.Name,
        Type: e.Type,
        Description: e.Description,
        Tags: util.StrToMem(e.Tags),
        From: util.StrToMem(e.From),
        To: util.StrToMem(e.To),
        Source: util.StrToMem(e.Source),
        NegateSource: util.YesNo(e.NegateSource),
        SourceUser: util.StrToMem(e.SourceUser),
        HipProfile: util.StrToMem(e.HipProfile),
        Destination: util.StrToMem(e.Destination),
        NegateDestination: util.YesNo(e.NegateDestination),
        Application: util.StrToMem(e.Application),
        Service: util.StrToMem(e.Service),
        Category: util.StrToMem(e.Category),
        Action: e.Action,
        LogSetting: e.LogSetting,
        LogStart: util.YesNo(e.LogStart),
        LogEnd: util.YesNo(e.LogEnd),
        Disabled: util.YesNo(e.Disabled),
        Schedule: e.Schedule,
        IcmpUnreachable: e.IcmpUnreachable,
        Options: &secOptions{util.YesNo(e.DisableServerResponseInspection)},
    }
    if e.Target != nil || e.NegateTarget {
        nfo := &targetInfo{
            Target: util.StrToEnt(e.Target),
            NegateTarget: util.YesNo(e.NegateTarget),
        }
        ans.TargetInfo = nfo
    }
    gs := e.Virus != nil || e.Spyware != nil || e.Vulnerability != nil || e.UrlFiltering != nil || e.FileBlocking != nil || e.WildFireAnalysis != nil || e.DataFiltering != nil
    if e.Group != nil || gs {
        ps := &profileSettings{
            Group: util.StrToMem(e.Group),
        }
        if gs {
            ps.Profiles = &profileSettingsProfile{
                util.StrToMem(e.Virus),
                util.StrToMem(e.Spyware),
                util.StrToMem(e.Vulnerability),
                util.StrToMem(e.UrlFiltering),
                util.StrToMem(e.FileBlocking),
                util.StrToMem(e.WildFireAnalysis),
                util.StrToMem(e.DataFiltering),
            }
        }
        ans.ProfileSettings = ps
    }

    return ans
}
