package layer3

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
	Shared        *SharedLocation        `json:"shared"`
	Template      *TemplateLocation      `json:"template,omitempty"`
	TemplateStack *TemplateStackLocation `json:"template_stack,omitempty"`
	Ngfw          *NgfwLocation          `json:"ngfw,omitempty"`
}

type SharedLocation struct {
}
type TemplateLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	Template       string `json:"template"`
}
type TemplateStackLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	TemplateStack  string `json:"template_stack"`
}
type NgfwLocation struct {
	NgfwDevice string `json:"ngfw_device"`
}

func NewSharedLocation() *Location {
	return &Location{Shared: &SharedLocation{},
	}
}
func NewTemplateLocation() *Location {
	return &Location{Template: &TemplateLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		Template:       "",
	},
	}
}
func NewTemplateStackLocation() *Location {
	return &Location{TemplateStack: &TemplateStackLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		TemplateStack:  "",
	},
	}
}
func NewNgfwLocation() *Location {
	return &Location{Ngfw: &NgfwLocation{
		NgfwDevice: "localhost.localdomain",
	},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Shared != nil:
		count++
	case o.Template != nil:
		if o.Template.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Template.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.Template.Template == "" {
			return fmt.Errorf("Template is unspecified")
		}
		count++
	case o.TemplateStack != nil:
		if o.TemplateStack.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateStack.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStack.TemplateStack == "" {
			return fmt.Errorf("TemplateStack is unspecified")
		}
		count++
	case o.Ngfw != nil:
		if o.Ngfw.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
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

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.Shared != nil:
		ans = []string{
			"config",
			"shared",
		}
	case o.Template != nil:
		if o.Template.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
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
			"devices",
			util.AsEntryXpath(o.Template.NgfwDevice),
		}
	case o.TemplateStack != nil:
		if o.TemplateStack.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
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
			"devices",
			util.AsEntryXpath(o.TemplateStack.NgfwDevice),
		}
	case o.Ngfw != nil:
		if o.Ngfw.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath(o.Ngfw.NgfwDevice),
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}

func (o Location) XpathWithComponents(vn version.Number, components ...string) ([]string, error) {
	if len(components) != 2 {
		return nil, fmt.Errorf("invalid number of arguments for XpathWithComponents() call")
	}

	{
		component := components[0]
		if !strings.HasPrefix(component, "entry[@name=\"]") && !strings.HasPrefix(component, "entry[@name='") {
			return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Parent must be formatted as entry: %s", component))
		}

		if !strings.HasSuffix(component, "\"]") && !strings.HasSuffix(component, "']") {
			return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Parent must be formatted as entry: %s", component))
		}
	}
	{
		component := components[1]
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

	ans = append(ans, "network")
	ans = append(ans, "interface")
	ans = append(ans, "aggregate-ethernet")
	ans = append(ans, components[0])
	ans = append(ans, "layer3")
	ans = append(ans, "units")
	ans = append(ans, components[1])

	return ans, nil
}
