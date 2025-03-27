package certificate

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type ImportLocation interface {
	XpathForLocation(version.Number, util.ILocation) ([]string, error)
	MarshalPangoXML([]string) (string, error)
	UnmarshalPangoXML([]byte) ([]string, error)
}

type Location struct {
	Panorama          bool                       `json:"panorama"`
	Template          *TemplateLocation          `json:"template,omitempty"`
	TemplateVsys      *TemplateVsysLocation      `json:"template_vsys,omitempty"`
	TemplateStack     *TemplateStackLocation     `json:"template_stack,omitempty"`
	TemplateStackVsys *TemplateStackVsysLocation `json:"template_stack_vsys,omitempty"`
	Shared            bool                       `json:"shared"`
}

type TemplateLocation struct {
	PanoramaDevice string `json:"panorama_device"`
	Template       string `json:"template"`
}

type TemplateVsysLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	Template       string `json:"template"`
	Vsys           string `json:"vsys"`
}

type TemplateStackLocation struct {
	PanoramaDevice string `json:"panorama_device"`
	TemplateStack  string `json:"template_stack"`
}

type TemplateStackVsysLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	TemplateStack  string `json:"template_stack"`
	Vsys           string `json:"vsys"`
}

func NewTemplateLocation() *Location {
	return &Location{Template: &TemplateLocation{
		PanoramaDevice: "localhost.localdomain",
		Template:       "",
	},
	}
}
func NewTemplateVsysLocation() *Location {
	return &Location{TemplateVsys: &TemplateVsysLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		Template:       "",
		Vsys:           "vsys1",
	},
	}
}
func NewTemplateStackLocation() *Location {
	return &Location{TemplateStack: &TemplateStackLocation{
		PanoramaDevice: "localhost.localdomain",
		TemplateStack:  "",
	},
	}
}
func NewTemplateStackVsysLocation() *Location {
	return &Location{TemplateStackVsys: &TemplateStackVsysLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		TemplateStack:  "",
		Vsys:           "vsys1",
	},
	}
}
func NewSharedLocation() *Location {
	return &Location{
		Shared: true,
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Panorama:
		count++
	case o.Template != nil:
		if o.Template.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.Template.Template == "" {
			return fmt.Errorf("Template is unspecified")
		}
		count++
	case o.TemplateVsys != nil:
		if o.TemplateVsys.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateVsys.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateVsys.Template == "" {
			return fmt.Errorf("Template is unspecified")
		}
		if o.TemplateVsys.Vsys == "" {
			return fmt.Errorf("Vsys is unspecified")
		}
		count++
	case o.TemplateStack != nil:
		if o.TemplateStack.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStack.TemplateStack == "" {
			return fmt.Errorf("TemplateStack is unspecified")
		}
		count++
	case o.TemplateStackVsys != nil:
		if o.TemplateStackVsys.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateStackVsys.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStackVsys.TemplateStack == "" {
			return fmt.Errorf("TemplateStack is unspecified")
		}
		if o.TemplateStackVsys.Vsys == "" {
			return fmt.Errorf("Vsys is unspecified")
		}
		count++
	case o.Shared:
		count++
	}

	if count == 0 {
		return fmt.Errorf("no path specified")
	}

	if count > 1 {
		return fmt.Errorf("multiple paths specified: only one should be specified")
	}

	return nil
}

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.Panorama:
		ans = []string{
			"config",
			"panorama",
		}
	case o.Template != nil:
		if o.Template.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.Template.Template == "" {
			return nil, fmt.Errorf("Template is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Template.PanoramaDevice}),
			"template",
			util.AsEntryXpath([]string{o.Template.Template}),
			"config",
			"shared",
		}
	case o.TemplateVsys != nil:
		if o.TemplateVsys.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateVsys.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateVsys.Template == "" {
			return nil, fmt.Errorf("Template is unspecified")
		}
		if o.TemplateVsys.Vsys == "" {
			return nil, fmt.Errorf("Vsys is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateVsys.PanoramaDevice}),
			"template",
			util.AsEntryXpath([]string{o.TemplateVsys.Template}),
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateVsys.NgfwDevice}),
			"vsys",
			util.AsEntryXpath([]string{o.TemplateVsys.Vsys}),
		}
	case o.TemplateStack != nil:
		if o.TemplateStack.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStack.TemplateStack == "" {
			return nil, fmt.Errorf("TemplateStack is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateStack.PanoramaDevice}),
			"template-stack",
			util.AsEntryXpath([]string{o.TemplateStack.TemplateStack}),
			"config",
			"shared",
		}
	case o.TemplateStackVsys != nil:
		if o.TemplateStackVsys.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateStackVsys.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStackVsys.TemplateStack == "" {
			return nil, fmt.Errorf("TemplateStack is unspecified")
		}
		if o.TemplateStackVsys.Vsys == "" {
			return nil, fmt.Errorf("Vsys is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateStackVsys.PanoramaDevice}),
			"template-stack",
			util.AsEntryXpath([]string{o.TemplateStackVsys.TemplateStack}),
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateStackVsys.NgfwDevice}),
			"vsys",
			util.AsEntryXpath([]string{o.TemplateStackVsys.Vsys}),
		}
	case o.Shared:
		ans = []string{
			"config",
			"shared",
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}
func (o Location) XpathWithEntryName(vn version.Number, name string) ([]string, error) {

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}
	ans = append(ans, Suffix...)
	ans = append(ans, util.AsEntryXpath([]string{name}))

	return ans, nil
}
func (o Location) XpathWithUuid(vn version.Number, uuid string) ([]string, error) {

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}
	ans = append(ans, Suffix...)
	ans = append(ans, util.AsUuidXpath(uuid))

	return ans, nil
}
