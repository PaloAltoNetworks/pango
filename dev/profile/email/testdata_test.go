package email

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
        {"v1 config", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            Config: "custom format",
        }},
        {"v2 system", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            System: "custom format",
        }},
        {"v2 threat", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            Threat: "custom format",
        }},
        {"v2 traffic", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            Traffic: "custom format",
        }},
        {"v2 hip match", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            HipMatch: "custom format",
        }},
        {"v2 escaped characters", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            EscapedCharacters: "abc",
        }},
        {"v2 escape character", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            EscapeCharacter: "x",
        }},
    }
}
