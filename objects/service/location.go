package service

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Location struct {
	Shared       bool                 `json:"shared"`
	Vsys         *VsysLocation        `json:"vsys,omitempty"`
	DeviceGroup  *DeviceGroupLocation `json:"device_group,omitempty"`
	FromPanorama bool                 `json:"from_panorama"`
}

func (o Location) IsValid() error {
	count := 0

	if o.Shared {
		count++
	}

	if o.Vsys != nil {
		if o.Vsys.Name == "" {
			return fmt.Errorf("vsys.name is unspecified")
		}
		if o.Vsys.NgfwDevice == "" {
			return fmt.Errorf("vsys.ngfw_device is unspecified")
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
		count++
	}

	if o.FromPanorama {
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

func (o Location) Xpath(vn version.Number, name string) ([]string, error) {
	var ans []string

	switch {
	case o.Shared:
		ans = []string{
			"config",
			"shared",
		}
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
		}
	case o.DeviceGroup != nil:
		if o.DeviceGroup.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.DeviceGroup.Name == "" {
			return nil, fmt.Errorf("Name is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.DeviceGroup.PanoramaDevice}),
			"device-group",
			util.AsEntryXpath([]string{o.DeviceGroup.Name}),
		}
	case o.FromPanorama:
		ans = []string{"config", "panorama"}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	ans = append(ans, Suffix...)
	ans = append(ans, util.AsEntryXpath([]string{name}))

	return ans, nil
}

type VsysLocation struct {
	NgfwDevice string `json:"ngfw_device"`
	Name       string `json:"name"`
}

type DeviceGroupLocation struct {
	PanoramaDevice string `json:"panorama_device"`
	Name           string `json:"name"`
}
