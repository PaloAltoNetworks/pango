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

type Member struct {
    Text string `xml:"member"`
}

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

type Entry struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
}

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

type XapiContainer interface {
    Normalize() (interface{})
}

type PanosEncapsulation interface {
    Load(map[int] []string, bool) error
    Dump() (map[int] []string)
    Element() (interface{})
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


// Gms retrieves a string from the given map, if present, else return the
// given default value.
func Gms(m map[int] []string, key int, dv string) string {
    if val, ok := m[key]; ok {
        return val[0]
    }
    return dv
}
