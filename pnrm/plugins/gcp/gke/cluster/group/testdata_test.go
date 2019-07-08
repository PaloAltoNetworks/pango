package group

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
        {"basic check", version.Number{9, 0, 0, ""}, Entry{
            Name: "v1",
            Description: "test",
            DeviceGroup: "dev group",
            TemplateStack: "my-stack",
            GcpProjectCredential: "my creds",
        }},
    }
}
