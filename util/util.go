// Package util contains various shared structs and functions used across
// the xapi package.
package util


import (
    "encoding/xml"
    "fmt"
    "regexp"
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
    LogUid(string, ...interface{})
    Op(interface{}, string, interface{}, interface{}) ([]byte, error)
    Show(interface{}, interface{}, interface{}) ([]byte, error)
    Get(interface{}, interface{}, interface{}) ([]byte, error)
    Delete(interface{}, interface{}, interface{}) ([]byte, error)
    Set(interface{}, interface{}, interface{}, interface{}) ([]byte, error)
    Edit(interface{}, interface{}, interface{}, interface{}) ([]byte, error)
    Move(interface{}, string, string, interface{}, interface{}) ([]byte, error)
    Uid(interface{}, string, interface{}, interface{}) ([]byte, error)
    EntryListUsing(func(interface{}, interface{}, interface{}) ([]byte, error), []string) ([]string, error)
    MemberListUsing(func(interface{}, interface{}, interface{}) ([]byte, error), []string) ([]string, error)
    RequestPasswordHash(string) (string, error)
    ImportInterfaces(string, []string) error
    UnimportInterfaces(string, []string) error
    ImportVlans(string, []string) error
    UnimportVlans(string, []string) error
    ImportVirtualRouters(string, []string) error
    UnimportVirtualRouters(string, []string) error
    WaitForJob(uint, interface{}) error
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
    Member []string `xml:"member"`
}

// MemToStr takes a pointer of a Member object and returns a list of strings.
func MemToStr(e *Member) []string {
    if e == nil {
        return nil
    }

    return e.Member
}

// StrToMem takes a list of strings and returns a list of Member objects.
func StrToMem(e []string) *Member {
    if e == nil {
        return nil
    }

    return &Member{e}
}

// Entry defines an entry config node used for sending and receiving XML
// from PANOS.
type Entry struct {
    Entry []innerEntry `xml:"entry"`
}

type innerEntry struct {
    Name string `xml:"name,attr"`
}

// EntToStr takes a list of Entry objects and returns a list of strings.
func EntToStr(e *Entry) []string {
    if e == nil {
        return nil
    }

    m := make([]string, len(e.Entry))
    for i := range e.Entry {
        m[i] = e.Entry[i].Name
    }

    return m
}

// StrToEnt takes a list of strings and returns a list of Entry objects.
func StrToEnt(e []string) *Entry {
    if e == nil {
        return nil
    }

    m := make([]innerEntry, len(e))
    for i := range e {
        m[i] = innerEntry{e[i]}
    }

    return &Entry{m}
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

// Lock represents either a config lock or a commit lock.
type Lock struct {
    XMLName xml.Name `xml:"entry"`
    Owner string `xml:"name,attr"`
    Name string `xml:"name"`
    Type string `xml:"type"`
    LoggedIn string `xml:"loggedin"`
    Comment CdataText `xml:"comment"`
}

type CdataText struct {
    Text string `xml:",cdata"`
}

type RawXml struct {
    Text string `xml:",innerxml"`
}

func CleanRawXml(v string) string {
    re := regexp.MustCompile(` admin="\S+" dirtyId="\d+" time="\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}"`)
    return re.ReplaceAllString(v, "")
}

// JobResponse parses a XML response that includes a job ID.
type JobResponse struct {
    XMLName xml.Name `xml:"response"`
    Id uint `xml:"result>job"`
}

// BasicJob is a struct for parsing minimal information about a submitted
// job to PANOS.
type BasicJob struct {
    XMLName xml.Name `xml:"response"`
    Result string `xml:"result>job>result"`
    Progress uint `xml:"result>job>progress"`
    Details []string `xml:"result>job>details>line"`
}
