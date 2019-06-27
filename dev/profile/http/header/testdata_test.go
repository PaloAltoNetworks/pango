package header

import (
    "github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
    desc string
    version version.Number
    conf Entry
}

func getTests() []tc {
    return []tc{
        {"basic", version.Number{7, 1, 0, ""}, Entry{
            Name: "Content-Type",
            Value: "application/json",
        }},
    }
}
