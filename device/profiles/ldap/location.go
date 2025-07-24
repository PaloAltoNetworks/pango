package ldap

import (
	"fmt"
	"strings"

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
	Panorama          *PanoramaLocation          `json:"panorama"`
	Shared            *SharedLocation            `json:"shared"`
	Vsys              *VsysLocation              `json:"vsys,omitempty"`
	Template          *TemplateLocation          `json:"template,omitempty"`
	TemplateVsys      *TemplateVsysLocation      `json:"template_vsys,omitempty"`
	TemplateStack     *TemplateStackLocation     `json:"template_stack,omitempty"`
	TemplateStackVsys *TemplateStackVsysLocation `json:"template_stack_vsys,omitempty"`
}

type PanoramaLocation struct {
}
type SharedLocation struct {
}
type VsysLocation struct {
	NgfwDevice string `json:"ngfw_device"`
	Vsys       string `json:"vsys"`
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

func NewPanoramaLocation() *Location {
	return &Location{Panorama: &PanoramaLocation{},
	}
}
func NewSharedLocation() *Location {
	return &Location{Shared: &SharedLocation{},
	}
}
func NewVsysLocation() *Location {
	return &Location{Vsys: &VsysLocation{
		NgfwDevice: "localhost.localdomain",
		Vsys:       "vsys1",
	},
	}
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

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Panorama != nil:
		count++
	case o.Shared != nil:
		count++
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Vsys == "" {
			return fmt.Errorf("Vsys is unspecified")
		}
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
	}

	if count == 0 {
		return fmt.Errorf("no path specified")
	}

	if count > 1 {
		return fmt.Errorf("multiple paths specified: only one should be specified")
	}

	return nil
}

func (o Location) LocationFilter() *string {

	return nil
}

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.Panorama != nil:
		ans = []string{
			"config",
			"panorama",
		}
	case o.Shared != nil:
		ans = []string{
			"config",
			"shared",
		}
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Vsys == "" {
			return nil, fmt.Errorf("Vsys is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath(o.Vsys.NgfwDevice),
			"vsys",
			util.AsEntryXpath(o.Vsys.Vsys),
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
			util.AsEntryXpath(o.Template.PanoramaDevice),
			"template",
			util.AsEntryXpath(o.Template.Template),
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
			util.AsEntryXpath(o.TemplateVsys.PanoramaDevice),
			"template",
			util.AsEntryXpath(o.TemplateVsys.Template),
			"config",
			"devices",
			util.AsEntryXpath(o.TemplateVsys.NgfwDevice),
			"vsys",
			util.AsEntryXpath(o.TemplateVsys.Vsys),
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
			util.AsEntryXpath(o.TemplateStack.PanoramaDevice),
			"template-stack",
			util.AsEntryXpath(o.TemplateStack.TemplateStack),
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
			util.AsEntryXpath(o.TemplateStackVsys.PanoramaDevice),
			"template-stack",
			util.AsEntryXpath(o.TemplateStackVsys.TemplateStack),
			"config",
			"devices",
			util.AsEntryXpath(o.TemplateStackVsys.NgfwDevice),
			"vsys",
			util.AsEntryXpath(o.TemplateStackVsys.Vsys),
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}

func (o Location) XpathWithComponents(vn version.Number, components ...string) ([]string, error) {
	if len(components) != 1 {
		return nil, fmt.Errorf("invalid number of arguments for XpathWithComponents() call")
	}

	{
		component := components[0]
		if component != "entry" {
			if !strings.HasPrefix(component, "entry[@name=\"]") && !strings.HasPrefix(component, "entry[@name='") {
				return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Name must be formatted as entry: %s", component))
			}

			if !strings.HasSuffix(component, "\"]") && !strings.HasSuffix(component, "']") {
				return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Name must be formatted as entry: %s", component))
			}
		}
	}

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	ans = append(ans, "server-profile")
	ans = append(ans, "ldap")
	ans = append(ans, components[0])

	return ans, nil
}
