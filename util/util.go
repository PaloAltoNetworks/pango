// Package util contains various shared structs and functions used across
// the xapi package.
package util


import (
    "encoding/xml"
    "fmt"
    "strings"

    "github.com/PaloAltoNetworks/xapi/version"
)

// Rulebase constants for various policies.
const (
    Rulebase = "rulebase"
    PreRulebase = "pre-rulebase"
    PostRulebase = "post-rulebase"
)

// XapiClient is the interface that describes an xapi.Client.
type XapiClient interface {
    String() string
    Versioning() version.Number
    RetrieveApiKey() error
    LogAction(string, ...interface{})
    LogQuery(string, ...interface{})
    LogOp(string, ...interface{})
    Op(interface{}, string, string, interface{}, interface{}) (*[]byte, error)
    Show(interface{}, interface{}, interface{}) (*[]byte, error)
    Get(interface{}, interface{}, interface{}) (*[]byte, error)
    Delete(interface{}, interface{}, interface{}) (*[]byte, error)
    Set(interface{}, interface{}, interface{}, interface{}) (*[]byte, error)
    Edit(interface{}, interface{}, interface{}, interface{}) (*[]byte, error)
    Move(interface{}, string, string, interface{}, interface{}) (*[]byte, error)
    EntryListUsing(func(interface{}, interface{}, interface{}) (*[]byte, error), []string) ([]string, error)
    MemberListUsing(func(interface{}, interface{}, interface{}) (*[]byte, error), []string) ([]string, error)
    RequestPasswordHash(string) (string, error)
    ImportInterfaces(string, []string) error
    UnimportInterfaces(string, []string) error
}

// BulkElement is a generic bulk container for bulk operations.
type BulkElement struct {
    XMLName xml.Name
    Data []interface{}
}

// Config returns an interface to be Marshaled.
func (o BulkElement) Config() interface{} {
    if len(o.Data) == 1 {
        return o.Data[0]
    }
    return o
}

// Member defines a member config node used for sending and receiving XML
// from PANOS.
type Member struct {
    Text string `xml:"member"`
}

// MemToStr takes a list of Member objects and returns a list of strings.
func MemToStr(e []Member) []string {
    if e == nil {
        return nil
    }

    m := make([]string, len(e))
    for i := range e {
        m[i] = e[i].Text
    }

    return m
}

// StrToMem takes a list of strings and returns a list of Member objects.
func StrToMem(e []string) []Member {
    if e == nil {
        return nil
    }

    m := make([]Member, len(e))
    for i := range e {
        m[i] = Member{e[i]}
    }

    return m
}

// Entry defines an entry config node used for sending and receiving XML
// from PANOS.
type Entry struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
}

// EntToStr takes a list of Entry objects and returns a list of strings.
func EntToStr(e []Entry) []string {
    if e == nil {
        return nil
    }

    m := make([]string, len(e))
    for i := range e {
        m[i] = e[i].Name
    }

    return m
}

// StrToEnt takes a list of strings and returns a list of Entry objects.
func StrToEnt(e []string) []Entry {
    if e == nil {
        return nil
    }

    m := make([]Entry, len(e))
    for i := range e {
        m[i] = Entry{Name: e[i]}
    }

    return m
}

// YesNo returns "yes" on true, "no" on false.
func YesNo(v bool) string {
    if v {
        return "yes"
    }
    return "no"
}

// AsBool returns true on yes, else false.
func AsBool(val string) bool {
    if val == "yes" {
        return true
    }
    return false
}

// AsXpath makes an xpath out of the given interface.
func AsXpath(i interface{}) string {
    switch val := i.(type) {
    case string:
        return val
    case []string:
        return fmt.Sprintf("/%s", strings.Join(val, "/"))
    default:
        return ""
    }
}

// AsEntryXpath returns the given values as an entry xpath segment.
func AsEntryXpath(vals []string) string {
    inner := make([]string, len(vals))
    for i := range inner {
        inner[i] = fmt.Sprintf("@name='%s'", vals[i])
    }

    return fmt.Sprintf("entry[%s]", strings.Join(inner, " or "))
}

// AsMemberXpath returns the given values as a member xpath segment.
func AsMemberXpath(vals []string) string {
    inner := make([]string, len(vals))
    for i := range inner {
        inner[i] = fmt.Sprintf("text()='%s'", vals[i])
    }

    return fmt.Sprintf("member[%s]", strings.Join(inner, " or "))
}

// License defines a license entry.
type License struct {
    XMLName xml.Name `xml:"entry"`
    Feature string `xml:"feature"`
    Description string `xml:"description"`
    Serial string `xml:"serial"`
    Issued string `xml:"issued"`
    Expires string `xml:"expires"`
    Expired string `xml:"expired"`
    AuthCode string `xml:"authcode"`
}
