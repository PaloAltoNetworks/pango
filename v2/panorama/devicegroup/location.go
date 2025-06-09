package devicegroup

import (
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/v2/errors"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

type ImportLocation interface {
	XpathForLocation(version.Number, util.ILocation) ([]string, error)
	MarshalPangoXML([]string) (string, error)
	UnmarshalPangoXML([]byte) ([]string, error)
}

type Location struct {
	Panorama *PanoramaLocation `json:"panorama,omitempty"`
}

type PanoramaLocation struct {
	PanoramaDevice string `json:"panorama_device"`
}

func NewPanoramaLocation() *Location {
	return &Location{Panorama: &PanoramaLocation{
		PanoramaDevice: "localhost.localdomain",
	},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Panorama != nil:
		if o.Panorama.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
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
	case o.Panorama != nil:
		if o.Panorama.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath(o.Panorama.PanoramaDevice),
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

	ans = append(ans, "device-group")
	ans = append(ans, components[0])

	return ans, nil
}
