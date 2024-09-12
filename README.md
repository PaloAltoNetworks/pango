Palo Alto Networks pango
========================

> [!NOTE]  
> This package is auto-generated via [pan-os-codegen](https://github.com/PaloAltoNetworks/pan-os-codegen)

> [!CAUTION]
> This software is currently in alpha development stage. It is strongly recommended not to use this package in production environments. If you choose to use it for experimental or developmental purposes, please do so with caution.


<!-- [![GoDoc](https://godoc.org/github.com/PaloAltoNetworks/pango?status.svg)](https://godoc.org/github.com/PaloAltoNetworks/pango) -->
<!-- [![Build](https://github.com/PaloAltoNetworks/pango/workflows/Sanity%20Check/badge.svg?branch=main)](https://github.com/PaloAltoNetworks/pango/actions?query=workflow%3A%22Sanity+Check%22) -->

Package pango is a golang cross version mechanism for interacting with Palo Alto Networks devices (including physical and virtualized Next-generation Firewalls and Panorama).  Versioning support is in place for PANOS 10.1 and above.

Please refer to the godoc reference documentation above to get started.

Using pango
===========

To start, create a client connection with the desired parameters and then initialize the connection:

```go
package main

import (
    "log"
    "github.com/PaloAltoNetworks/pango"
)

func main() {
    var err error

    con = &sdk.Client{
        Hostname: "127.0.0.1",
        Username: "admin",
        Password: "admin",
    }

    if err := con.Setup(); err != nil {
        log.Printf("Failed to setup client: %s", err)
        return
    }

    if err := con.Initialize(ctx); err != nil {
        log.Printf("Failed to initialize client: %s", err)
        return
    }

    log.Printf("Initialize ok")

}
```
