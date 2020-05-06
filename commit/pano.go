package commit

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
)

// PanoramaCommit is a normalized Panorama commit.
//
// Admins is the list of admins whose changes should be committed.
type PanoramaCommit struct {
	Description             string
	Admins                  []string
	DeviceGroups            []string
	Templates               []string
	TemplateStacks          []string
	WildfireAppliances      []string
	WildfireClusters        []string
	LogCollectors           []string
	LogCollectorGroups      []string
	ExcludeDeviceAndNetwork bool
	ExcludeSharedObjects    bool
	Force                   bool
}

// Action returns a commit action of an empty string.
func (o PanoramaCommit) Action() string { return "" }

// Element returns an interface to be marshalled to perform the specified commit.
func (o PanoramaCommit) Element() interface{} {
	ans := panoCommit{
		Description: o.Description,
	}

	var p *panoPartialCommit
	if len(o.Admins) > 0 || len(o.DeviceGroups) > 0 ||
		len(o.WildfireAppliances) > 0 || len(o.WildfireClusters) > 0 ||
		len(o.LogCollectors) > 0 || len(o.LogCollectorGroups) > 0 ||
		len(o.Templates) > 0 || len(o.TemplateStacks) > 0 ||
		o.ExcludeDeviceAndNetwork || o.ExcludeSharedObjects {
		p = &panoPartialCommit{
			Admins:             util.StrToMem(o.Admins),
			DeviceGroups:       util.StrToMem(o.DeviceGroups),
			Templates:          util.StrToMem(o.Templates),
			TemplateStacks:     util.StrToMem(o.TemplateStacks),
			WildfireAppliances: util.StrToMem(o.WildfireAppliances),
			WildfireClusters:   util.StrToMem(o.WildfireClusters),
			LogCollectors:      util.StrToMem(o.LogCollectors),
			LogCollectorGroups: util.StrToMem(o.LogCollectorGroups),
		}

		if o.ExcludeDeviceAndNetwork {
			p.ExcludeDeviceAndNetwork = "excluded"
		}

		if o.ExcludeSharedObjects {
			p.ExcludeSharedObjects = "excluded"
		}
	}

	if o.Force {
		ans.Force = &panoForce{
			Partial: p,
		}
	} else {
		ans.Partial = p
	}

	return ans
}

type panoCommit struct {
	XMLName     xml.Name           `xml:"commit"`
	Description string             `xml:"description,omitempty"`
	Force       *panoForce         `xml:"force"`
	Partial     *panoPartialCommit `xml:"partial"`
}

type panoForce struct {
	Partial *panoPartialCommit `xml:"partial"`
}

type panoPartialCommit struct {
	Admins                  *util.MemberType `xml:"admin"`
	DeviceGroups            *util.MemberType `xml:"device-group"`
	Templates               *util.MemberType `xml:"template"`
	TemplateStacks          *util.MemberType `xml:"template-stack"`
	WildfireAppliances      *util.MemberType `xml:"wildfire-appliance"`
	WildfireClusters        *util.MemberType `xml:"wildfire-appliance-cluster"`
	LogCollectors           *util.MemberType `xml:"log-collector"`
	LogCollectorGroups      *util.MemberType `xml:"log-collector-group"`
	ExcludeDeviceAndNetwork string           `xml:"device-and-network,omitempty"`
	ExcludeSharedObjects    string           `xml:"shared-object,omitempty"`
}

/*
PanoramaCommitAll is a normalized Panorama commit-all.

Depending on the type of commit specified, only certain parameters of this
object are used.

Regardless of the type given, all commits will use the Description  and Name params.

TypeDeviceGroup types uses Devices, IncludeTemplate, and ForceTemplateValues.

TypeTemplate and TypeTemplateStack uses Devices and ForceTemplateValues.
*/
type PanoramaCommitAll struct {
	Type                string
	Name                string
	Description         string
	IncludeTemplate     bool
	ForceTemplateValues bool
	Devices             []string
}

// Action returns a commit action of "all".
func (o PanoramaCommitAll) Action() string { return "all" }

// Element returns an interface to be marshalled to perform the specified commit.
func (o PanoramaCommitAll) Element() interface{} {
	var ans panoCommitAll

	switch o.Type {
	case TypeDeviceGroup:
		ans = panoCommitAll{
			DeviceGroup: &pcaDeviceGroup{
				DgInfo: pcaDgInfo{
					Entry: pcaDgInfoEntry{
						Name:    o.Name,
						Devices: util.StrToEnt(o.Devices),
					},
				},
				Description:         o.Description,
				IncludeTemplate:     util.YesNo(o.IncludeTemplate),
				ForceTemplateValues: util.YesNo(o.ForceTemplateValues),
			},
		}
	case TypeTemplate:
		ans = panoCommitAll{
			Template: &pcaTemplate{
				Name:                o.Name,
				Description:         o.Description,
				ForceTemplateValues: util.YesNo(o.ForceTemplateValues),
				Devices:             util.StrToMem(o.Devices),
			},
		}
	case TypeTemplateStack:
		ans = panoCommitAll{
			TemplateStack: &pcaTemplate{
				Name:                o.Name,
				Description:         o.Description,
				ForceTemplateValues: util.YesNo(o.ForceTemplateValues),
				Devices:             util.StrToMem(o.Devices),
			},
		}
	case TypeLogCollectorGroup:
		ans = panoCommitAll{
			LogCollectorGroup: &pcaLogCollectorGroup{
				Name:        o.Name,
				Description: o.Description,
			},
		}
	case TypeWildfireAppliance:
		ans = panoCommitAll{
			Wildfire: &pcaWildfire{
				Appliance:   o.Name,
				Description: o.Description,
			},
		}
	case TypeWildfireCluster:
		ans = panoCommitAll{
			Wildfire: &pcaWildfire{
				Cluster:     o.Name,
				Description: o.Description,
			},
		}
	}

	return ans
}

type panoCommitAll struct {
	XMLName           xml.Name              `xml:"commit-all"`
	DeviceGroup       *pcaDeviceGroup       `xml:"shared-policy"`
	Template          *pcaTemplate          `xml:"template"`
	TemplateStack     *pcaTemplate          `xml:"template-stack"`
	LogCollectorGroup *pcaLogCollectorGroup `xml:"log-collector-config"`
	Wildfire          *pcaWildfire          `xml:"wildfire-appliance-config"`
}

type pcaDeviceGroup struct {
	DgInfo              pcaDgInfo `xml:"device-group"`
	Description         string    `xml:"description,omitempty"`
	IncludeTemplate     string    `xml:"include-template"`
	ForceTemplateValues string    `xml:"force-template-values"`
}

type pcaDgInfo struct {
	Entry pcaDgInfoEntry `xml:"entry"`
}

type pcaDgInfoEntry struct {
	Name    string          `xml:"name,attr"`
	Devices *util.EntryType `xml:"devices"`
}

type pcaTemplate struct {
	Name                string           `xml:"name"`
	Description         string           `xml:"description,omitempty"`
	Devices             *util.MemberType `xml:"device"`
	ForceTemplateValues string           `xml:"force-template-values"`
}

type pcaLogCollectorGroup struct {
	Name        string `xml:"log-collector-group"`
	Description string `xml:"description,omitempty"`
}

type pcaWildfire struct {
	Description string `xml:"description,omitempty"`
	Appliance   string `xml:"wildfire-appliance,omitempty"`
	Cluster     string `xml:"wildfire-appliance-cluster,omitempty"`
}
