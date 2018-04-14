package pango

import (
    "encoding/xml"

    // Various namespace imports.
    "github.com/PaloAltoNetworks/pango/netw"
    "github.com/PaloAltoNetworks/pango/dev"
    "github.com/PaloAltoNetworks/pango/poli"
    "github.com/PaloAltoNetworks/pango/objs"
    "github.com/PaloAltoNetworks/pango/licen"
    "github.com/PaloAltoNetworks/pango/userid"
)


// Firewall is a firewall specific client, providing version safe functions
// for the PAN-OS Xpath API methods.  After creating the object, invoke
// Initialize() to prepare it for use.
//
// It has the following namespaces:
//      * Network
//      * Device
//      * Policies
//      * Objects
//      * Licensing
//      * UserId
type Firewall struct {
    Client

    // Namespaces
    Network *netw.Netw
    Device *dev.Dev
    Policies *poli.FwPoli
    Objects *objs.FwObjs
    Licensing *licen.Licen
    UserId *userid.UserId
}

// Initialize does some initial setup of the Firewall connection, retrieves
// the API key if it was not already present, then performs "show system
// info" to get the PAN-OS version.  The full results are saved into the
// client's SystemInfo map.
//
// If not specified, the following is assumed:
//  * Protocol: https
//  * Port: (unspecified)
//  * Timeout: 10
//  * Logging: LogAction | LogUid
func (c *Firewall) Initialize() error {
    if len(c.rb) == 0 {
        var e error

        if e = c.initCon(); e != nil {
            return e
        } else if e = c.initApiKey(); e != nil {
            return e
        } else if e = c.initSystemInfo(); e != nil {
            return e
        }
    } else {
        c.Hostname = "localhost"
        c.ApiKey = "password"
    }
    c.initNamespaces()

    return nil
}

// Commit performs a Firewall commit.
//
// Param desc is the optional commit description message you want associated
// with the commit.
//
// Params dan and pao are advanced options for doing partial commits.  Setting
// param dan to false excludes the Device and Network configuration, while
// setting param pao to false excludes the Policy and Object configuration.
//
// Param force is if you want to force a commit even if no changes are
// required.
//
// Param sync should be true if you want this function to block until the
// commit job completes.
//
// Commits result in a job being submitted to the backend.  The job ID and
// if an error was encountered or not are returned from this function.
func (c *Firewall) Commit(desc string, dan, pao, force, sync bool) (uint, error) {
    c.LogAction("(commit) %q", desc)

    req := fwCommit{Description: desc}
    if !dan || !pao {
        req.Partial = &fwCommitPartial{}
        if !dan {
            req.Partial.Dan = "excluded"
        }
        if !pao {
            req.Partial.Pao = "excluded"
        }
    }
    if force {
        req.Force = ""
    }

    job, _, err := c.CommitConfig(req, "", nil)
    if err != nil || !sync || job == 0 {
        return job, err
    }

    return job, c.WaitForJob(job, nil)
}

/** Private functions **/

func (c *Firewall) initNamespaces() {
    c.Network = &netw.Netw{}
    c.Network.Initialize(c)

    c.Device = &dev.Dev{}
    c.Device.Initialize(c)

    c.Policies = &poli.FwPoli{}
    c.Policies.Initialize(c)

    c.Objects = &objs.FwObjs{}
    c.Objects.Initialize(c)

    c.Licensing = &licen.Licen{}
    c.Licensing.Initialize(c)

    c.UserId = &userid.UserId{}
    c.UserId.Initialize(c)
}

/** Internal structs / functions **/

type fwCommit struct {
    XMLName xml.Name `xml:"commit"`
    Description string `xml:"description,omitempty"`
    Partial *fwCommitPartial `xml:"partial"`
    Force interface{} `xml:"force"`
}

type fwCommitPartial struct {
    Dan string `xml:"device-and-network,omitempty"`
    Pao string `xml:"policy-and-objects,omitempty"`
}
