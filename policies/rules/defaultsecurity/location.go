package defaultsecurity

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
	Shared      *SharedLocation      `json:"shared"`
	DeviceGroup *DeviceGroupLocation `json:"device_group,omitempty"`
}

type SharedLocation struct {
}
type DeviceGroupLocation struct {
	DeviceGroup    string `json:"device_group"`
	PanoramaDevice string `json:"panorama_device"`
}

func NewSharedLocation() *Location {
	return &Location{Shared: &SharedLocation{},
	}
}
func NewDeviceGroupLocation() *Location {
	return &Location{DeviceGroup: &DeviceGroupLocation{
		DeviceGroup:    "",
		PanoramaDevice: "localhost.localdomain",
	},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Shared != nil:
		count++
	case o.DeviceGroup != nil:
		if o.DeviceGroup.DeviceGroup == "" {
			return fmt.Errorf("DeviceGroup is unspecified")
		}
		if o.DeviceGroup.PanoramaDevice == "" {
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

func (o Location) LocationFilter() *string {

	if o.DeviceGroup != nil {
		if o.DeviceGroup.DeviceGroup == "" {
			return nil
		} else {
			return &o.DeviceGroup.DeviceGroup
		}
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
			"post-rulebase",
		}
	case o.DeviceGroup != nil:
		if o.DeviceGroup.DeviceGroup == "" {
			return nil, fmt.Errorf("DeviceGroup is unspecified")
		}
		if o.DeviceGroup.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath(o.DeviceGroup.PanoramaDevice),
			"device-group",
			util.AsEntryXpath(o.DeviceGroup.DeviceGroup),
			"post-rulebase",
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

	ans = append(ans, "default-security-rules")
	ans = append(ans, "rules")
	ans = append(ans, components[0])

	return ans, nil
}
