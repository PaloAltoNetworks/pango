package ntp

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
	System        *SystemLocation        `json:"system,omitempty"`
	Template      *TemplateLocation      `json:"template,omitempty"`
	TemplateStack *TemplateStackLocation `json:"template_stack,omitempty"`
}

type SystemLocation struct {
	Device string `json:"device"`
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

func NewSystemLocation() *Location {
	return &Location{System: &SystemLocation{
		Device: "localhost.localdomain",
	},
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

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.System != nil:
		if o.System.Device == "" {
			return fmt.Errorf("Device is unspecified")
		}
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
	case o.System != nil:
		if o.System.Device == "" {
			return nil, fmt.Errorf("Device is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath(o.System.Device),
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
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}

func (o Location) XpathWithComponents(vn version.Number, components ...string) ([]string, error) {
	if len(components) != 0 {
		return nil, fmt.Errorf("invalid number of arguments for XpathWithComponents() call")
	}

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	ans = append(ans, "deviceconfig")
	ans = append(ans, "system")

	return ans, nil
}
