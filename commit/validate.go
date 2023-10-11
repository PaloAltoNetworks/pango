package commit

import (
	"encoding/xml"
	"github.com/PaloAltoNetworks/pango/util"
)

// PanoramaValidate is a Panorama Valid structure contained all option for Client.ValidateConfig().
type PanoramaValidate struct {
	Admins                    []string
	DeviceGroups              []string
	LogCollectors             []string
	LogCollectorGroups        []string
	Templates                 []string
	TemplateStacks            []string
	WildfireAppliances        []string
	WildfireApplianceClusters []string
	ExcludeDeviceAndNetwork   bool
	ExcludeSharedObjects      bool
	Full                      bool
}

// PanoramaValidateAll it is a shell struct for doing validate all changes in Panorama.
type PanoramaValidateAll struct{}

// Element this function return an interface ready to be marshalled into XML for validate.
func (o PanoramaValidate) Element() interface{} {
	ans := panoValidate{}

	var p *panoValidatePartial
	if len(o.Admins) > 0 || len(o.DeviceGroups) > 0 || len(o.Templates) > 0 ||
		len(o.TemplateStacks) > 0 || len(o.LogCollectors) > 0 || len(o.LogCollectorGroups) > 0 {
		p = &panoValidatePartial{
			Admins:             util.StrToMem(o.Admins),
			DeviceGroups:       util.StrToMem(o.DeviceGroups),
			Templates:          util.StrToMem(o.Templates),
			TemplateStacks:     util.StrToMem(o.TemplateStacks),
			WildfireAppliances: util.StrToMem(o.WildfireAppliances),
			WildfireClusters:   util.StrToMem(o.WildfireApplianceClusters),
			LogCollectors:      util.StrToMem(o.LogCollectors),
			LogCollectorGroups: util.StrToMem(o.LogCollectorGroups),
		}

		if o.ExcludeSharedObjects {
			p.ExcludeSharedObjects = "excluded"
		}

		if o.ExcludeDeviceAndNetwork {
			p.ExcludeDeviceAndNetwork = "excluded"
		}
	} else {
		ans.Full = &panoValidateFull{}
	}

	ans.Partial = p

	return ans
}

// panoValidate main XML structure for validate, it contains main 'validate' option and regarded for the action
// it can be either 'full' or 'partial'
type panoValidate struct {
	XMLName xml.Name             `xml:"validate"`
	Full    *panoValidateFull    `xml:"full"`
	Partial *panoValidatePartial `xml:"partial"`
}

type panoValidateFull struct{}

type panoValidatePartial struct {
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
