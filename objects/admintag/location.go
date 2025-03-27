package admintag

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
	Shared      bool                 `json:"shared"`
	Vsys        *VsysLocation        `json:"vsys,omitempty"`
	DeviceGroup *DeviceGroupLocation `json:"device_group,omitempty"`
}

type VsysLocation struct {
	NgfwDevice string `json:"ngfw_device"`
	Vsys       string `json:"vsys"`
}

type DeviceGroupLocation struct {
	DeviceGroup    string `json:"device_group"`
	PanoramaDevice string `json:"panorama_device"`
}

func NewSharedLocation() *Location {
	return &Location{
		Shared: true,
	}
}
func NewVsysLocation() *Location {
	return &Location{Vsys: &VsysLocation{
		NgfwDevice: "localhost.localdomain",
		Vsys:       "vsys1",
	},
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
	case o.Shared:
		count++
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Vsys == "" {
			return fmt.Errorf("Vsys is unspecified")
		}
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

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

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
		if o.Vsys.Vsys == "" {
			return nil, fmt.Errorf("Vsys is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Vsys.NgfwDevice}),
			"vsys",
			util.AsEntryXpath([]string{o.Vsys.Vsys}),
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
			util.AsEntryXpath([]string{o.DeviceGroup.PanoramaDevice}),
			"device-group",
			util.AsEntryXpath([]string{o.DeviceGroup.DeviceGroup}),
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
