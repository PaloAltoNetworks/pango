Palo Alto Networks xapi
=======================

Package xapi is a golang cross version mechanism for interacting with Palo Alto Networks devices (including physical and virtualized Next-generation Firewalls and Panorama).

To start, create a client connection with the desired parameters and then initialize the connection:

```go
package main

import (
    "log"
    "github.com/PaloAltoNetworks/xapi"
)

func main() {
    var err error

    c := &xapi.Firewall{Client: xapi.Client{
        Hostname: "127.0.0.1",
        Username: "admin",
        Password: "admin",
        Logging: xapi.LogAction | xapi.LogOp,
    }}
    if err = c.Initialize(); err != nil {
        log.Printf("Failed to initialize client: %s", err)
        return
    }
}
```

Initializing the connection creates the API key (if it was not already specified), then performs `show system info` to get the PANOS version.  Once the client connection is created, you can query and configure the Palo Alto Networks device from the functions inside the various namespaces of the client connection.  Namespaces correspond to the various configuration areas available in the GUI.  For example:

```go
    err = c.Network.EthernetInterface.Set(...)
    myPolicies, err := c.Policies.Security.GetList(...)
```

Generally speaking, there are the following functions inside each namespace:

    * `GetList`
    * `ShowList`
    * `Get`
    * `Show`
    * `Set`
    * `Edit`
    * `Delete`

These functions correspond with PANOS `Get`, `Show`, `Set`, `Edit`, and `Delete` API calls.  `Get()`, `Set()`, and `Edit()` take and return normalized, version independent objects.  These version safe objects are typically named `Entry`, which corresponds to how the object is placed in the PANOS XPATH.  For any version safe object, attempting to configure a parameter that your PANOS doesn't support will be safely ignored in the resultant XML sent to the firewall / Panorama.


Using `Edit` Functions
======================

The PANOS XML API `Edit` command can be used to both create as well as update existing config, however it can also truncate config for the given XPATH.  Due to this, if you want to use `Edit()`, you need to make sure that you perform either a `Get()` or a `Show()` first, make your modification, then invoke `Edit()` using that object.  If you don't do this, you will truncate any sub config.
