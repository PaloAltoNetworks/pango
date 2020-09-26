package commit

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
)

// FirewallCommit is a normalized object for defining firewall commits.
//
// This is a commit type, designed to be passed in to Client.Commit().
//
// Admins is the list of admins whose changes should be committed.
type FirewallCommit struct {
	Description             string
	Admins                  []string
	ExcludeDeviceAndNetwork bool
	ExcludeSharedObjects    bool
	ExcludePolicyAndObjects bool
	Force                   bool
}

// Action returns a commit action of an empty string.
func (o FirewallCommit) Action() string { return "" }

// Element returns an interface to be marshalled to perform the specified commit.
func (o FirewallCommit) Element() interface{} {
	ans := fwCommit{
		Description: o.Description,
	}

	var p *fwPartialCommit
	if len(o.Admins) > 0 ||
		o.ExcludeDeviceAndNetwork || o.ExcludeSharedObjects || o.ExcludePolicyAndObjects {
		p = &fwPartialCommit{
			Admins: util.StrToMem(o.Admins),
		}

		if o.ExcludeDeviceAndNetwork {
			p.ExcludeDeviceAndNetwork = "excluded"
		}

		if o.ExcludeSharedObjects {
			p.ExcludeSharedObjects = "excluded"
		}

		if o.ExcludePolicyAndObjects {
			p.ExcludePolicyAndObjects = "excluded"
		}
	}

	if o.Force {
		ans.Force = &fwForce{
			Partial: p,
		}
	} else {
		ans.Partial = p
	}

	return ans
}

type fwCommit struct {
	XMLName     xml.Name         `xml:"commit"`
	Description string           `xml:"description,omitempty"`
	Force       *fwForce         `xml:"force"`
	Partial     *fwPartialCommit `xml:"partial"`
}

type fwForce struct {
	Partial *fwPartialCommit `xml:"partial"`
}

type fwPartialCommit struct {
	Admins                  *util.MemberType `xml:"admin"`
	ExcludeDeviceAndNetwork string           `xml:"device-and-network,omitempty"`
	ExcludeSharedObjects    string           `xml:"shared-object,omitempty"`
	ExcludePolicyAndObjects string           `xml:"policy-and-objects,omitempty"`
}
