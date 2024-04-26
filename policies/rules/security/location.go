package security

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Location struct {
	Vsys        *VsysLocation        `json:"vsys,omitempty"`
	Shared      *SharedLocation      `json:"shared"`
	DeviceGroup *DeviceGroupLocation `json:"device_group,omitempty"`
}

func (o Location) IsValid() error {
	count := 0

	if o.Vsys != nil {
		if o.Vsys.Name == "" {
			return fmt.Errorf("vsys.name is unspecified")
		}
		if o.Vsys.NgfwDevice == "" {
			return fmt.Errorf("vsys.ngfw_device is unspecified")
		}
		count++
	}

	if o.Shared != nil {
		if o.Shared.Rulebase == "" {
			return fmt.Errorf("shared.rulebase is unspecified")
		}
		count++
	}

	if o.DeviceGroup != nil {
		if o.DeviceGroup.Name == "" {
			return fmt.Errorf("device_group.name is unspecified")
		}
		if o.DeviceGroup.PanoramaDevice == "" {
			return fmt.Errorf("device_group.panorama_device is unspecified")
		}
		if o.DeviceGroup.Rulebase == "" {
			return fmt.Errorf("device_group.rulebase is unspecified")
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

func (o Location) Xpath(vn version.Number, name, uuid string) ([]string, error) {
	var ans []string

	switch {
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Name == "" {
			return nil, fmt.Errorf("Name is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Vsys.NgfwDevice}),
			"vsys",
			util.AsEntryXpath([]string{o.Vsys.Name}),
			"rulebase",
		}
	case o.Shared != nil:
		if o.Shared.Rulebase == "" {
			return nil, fmt.Errorf("Rulebase is unspecified")
		}
		ans = []string{
			"config",
			"shared",
			o.Shared.Rulebase,
		}
	case o.DeviceGroup != nil:
		if o.DeviceGroup.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.DeviceGroup.Name == "" {
			return nil, fmt.Errorf("Name is unspecified")
		}
		if o.DeviceGroup.Rulebase == "" {
			return nil, fmt.Errorf("Rulebase is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.DeviceGroup.PanoramaDevice}),
			"device-group",
			util.AsEntryXpath([]string{o.DeviceGroup.Name}),
			o.DeviceGroup.Rulebase,
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	ans = append(ans, Suffix...)
	if uuid != "" {
		ans = append(ans, util.AsUuidXpath(uuid))
	} else {
		ans = append(ans, util.AsEntryXpath([]string{name}))
	}

	return ans, nil
}

type VsysLocation struct {
	NgfwDevice string `json:"ngfw_device"`
	Name       string `json:"name"`
}

type SharedLocation struct {
	Rulebase string `json:"rulebase"`
}

type DeviceGroupLocation struct {
	PanoramaDevice string `json:"panorama_device"`
	Rulebase       string `json:"rulebase"`
	Name           string `json:"name"`
}
