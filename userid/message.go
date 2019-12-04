package userid

import (
    "encoding/xml"
    "fmt"
    "strings"

    "github.com/PaloAltoNetworks/pango/util"
)

/*
Message is a user-id message to be sent to PAN-OS.

This can contain multiple actions to be performed, such as
logging in a user, tagging an IP, or setting group membership.
*/
type Message struct {
    Logins []Login
    Logouts []Logout
    TagIps []TagIp
    UntagIps []UntagIp
    Groups []Group
}

// Login logs a user in.
type Login struct {
    User string
    Ip string
}

// Logout logs a user out.
type Logout struct {
    User string
    Ip string
}

// TagIp assigns tags to the specified IP address.
type TagIp struct {
    Ip string
    Tags []string
}

// UntagIp removes tags from the specified IP address.
type UntagIp struct {
    Ip string
    Tags []string
}

/*
Group specifies a static user group's members.

Some care is needed when dealing with the groups.  You cannot add or
remove individual users from a group.  You have to specify the group in its
entirety each time you want to update group membership.
*/
type Group struct {
    Name string
    Users []string
}

type uid struct {
    XMLName xml.Name `xml:"uid-message"`
    Version string `xml:"version"`
    Type string `xml:"type"`
    Payload payload `xml:"payload"`
}

type payload struct {
    Login *loginLogoutSpec `xml:"login"`
    Logout *loginLogoutSpec `xml:"logout"`
    TagIp *tagUntagIpSpec `xml:"register"`
    UntagIp *tagUntagIpSpec `xml:"unregister"`
    Group *groupSpec `xml:"groups"`
}

type loginLogoutSpec struct {
    Entry []inOut `xml:"entry"`
}

type inOut struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    Ip string `xml:"ip,attr"`
}

type tagUntagIpSpec struct {
    Entry []ipTag `xml:"entry"`
}

type ipTag struct {
    XMLName xml.Name `xml:"entry"`
    Ip string `xml:"ip,attr"`
    Tags []string `xml:"tag>member"`
}

type groupSpec struct {
    Entries []groupDef `xml:"entry"`
}

type groupDef struct {
    Name string `xml:"name,attr"`
    Members util.EntryType `xml:"members"`
}

func encode(m *Message) (*uid, string) {
    var buf strings.Builder
    if m == nil {
        return nil, ""
    }

    msg := uid{
        Version: "1.0",
        Type: "update",
    }

    // Login users.
    if len(m.Logins) > 0 {
        buf.WriteString(fmt.Sprintf(" logins:%d", len(m.Logins)))
        msg.Payload.Login = &loginLogoutSpec{}
        msg.Payload.Login.Entry = make([]inOut, 0, len(m.Logins))
        for i := range m.Logins {
            x := inOut{
                Name: m.Logins[i].User,
                Ip: m.Logins[i].Ip,
            }
            msg.Payload.Login.Entry = append(msg.Payload.Login.Entry, x)
        }
    }

    // Logout users.
    if len(m.Logouts) > 0 {
        buf.WriteString(fmt.Sprintf(" logouts:%d", len(m.Logouts)))
        msg.Payload.Logout = &loginLogoutSpec{}
        msg.Payload.Logout.Entry = make([]inOut, 0, len(m.Logouts))
        for i := range m.Logouts {
            x := inOut{
                Name: m.Logouts[i].User,
                Ip: m.Logouts[i].Ip,
            }
            msg.Payload.Logout.Entry = append(msg.Payload.Logout.Entry, x)
        }
    }

    // Tag IPs.
    if len(m.TagIps) > 0 {
        buf.WriteString(fmt.Sprintf(" tagip:%d", len(m.TagIps)))
        msg.Payload.TagIp = &tagUntagIpSpec{}
        msg.Payload.TagIp.Entry = make([]ipTag, 0, len(m.TagIps))
        for i := range m.TagIps {
            x := ipTag{
                Ip: m.TagIps[i].Ip,
                Tags: m.TagIps[i].Tags,
            }
            msg.Payload.TagIp.Entry = append(msg.Payload.TagIp.Entry, x)
        }
    }

    // Remove tags from IPs.
    if len(m.UntagIps) > 0 {
        buf.WriteString(fmt.Sprintf(" untagip:%d", len(m.UntagIps)))
        msg.Payload.UntagIp = &tagUntagIpSpec{}
        msg.Payload.UntagIp.Entry = make([]ipTag, 0, len(m.UntagIps))
        for i := range m.UntagIps {
            x := ipTag{
                Ip: m.UntagIps[i].Ip,
                Tags: m.UntagIps[i].Tags,
            }
            msg.Payload.UntagIp.Entry = append(msg.Payload.UntagIp.Entry, x)
        }
    }

    // Specify a static group of users.
    if len(m.Groups) > 0 {
        buf.WriteString(fmt.Sprintf(" group:%d", len(m.Groups)))
        msg.Payload.Group = &groupSpec{}
        msg.Payload.Group.Entries = make([]groupDef, 0, len(m.Groups))
        for i := range m.Groups {
            members := util.StrToEnt(m.Groups[i].Users)
            if members == nil {
                members = &util.EntryType{}
            }
            x := groupDef{
                Name: m.Groups[i].Name,
                Members: *members,
            }
            msg.Payload.Group.Entries = append(msg.Payload.Group.Entries, x)
        }
    }

    if buf.Len() == 0 {
        return nil, ""
    }

    return &msg, buf.String()
}
