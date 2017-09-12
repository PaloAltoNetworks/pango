=======================
Palo Alto Networks xapi
=======================

Package xapi is a cross version mechanism for interacting with Palo Alto
Networks devices (including physical and virtualized Next-generation Firewalls
and Panorama).

To start, create a client connection with the desired parameters and then
initialize the connection:

    c := xapi.Client{
        Hostname: "127.0.0.1",
        Username: "admin",
        Password: "admin",
    }
    err := c.Initialize()
    if err != nil {
        log.Printf("Failed to initialize client: %s", err)
        return
    }

Initializing the connection creates the API key (if it was not already
specified), then performs "show system info" to get the PANOS version.  Once
the client connection is created, you can query and configure the Palo
Alto Networks device from the functions inside the various namespaces of the
client connection.  Namespaces correspond to the various configuration areas
available in the GUI.  For example:

    err = c.Network.EthernetInterface.Set(...)
    myPolicies, err := c.Policies.Security.GetList(...)

Generally speaking, there are the following functions inside each namespace:

* GetList
* ShowList
* Get
* Show
* Set
* Delete

These functions correspond with PANOS GET, SHOW, SET, and DELETE API calls.  The
Get and Set functions take and return normalized, version independent objects.
These version safe objects are typically named Entry, which corresponds to
how the object is placed in the PANOS XPATH.  For any version safe object,
attempting to configure a parameter that your PANOS doesn't support will be
safely ignored in the resultant XML sent to the firewall / Panorama.

Those more familiar with PANOS XAPI may notice the lack of Edit above.  Due
to singular focus of functions in package xapi and how Edit truncates config,
there is no namespace support for Edit calls.  You can, however, still invoke
an Edit if desired, but creating the XPATH and XML document is up to you:

    c.Edit(...)
