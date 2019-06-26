package server

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
        {"v1 basic", version.Number{7, 1, 0, ""}, Entry{
            Name: "t1",
            Server: "syslog.server.int",
            Transport: TransportSsl,
            Port: 42,
            SyslogFormat: SyslogFormatIetf,
            Facility: FacilityLocal4,
        }},
    }
}
