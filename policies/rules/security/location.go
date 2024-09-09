package security

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
	DeviceGroup      *DeviceGroupLocation      `json:"device_group,omitempty"`
	FromPanoramaVsys *FromPanoramaVsysLocation `json:"from_panorama_vsys,omitempty"`
	Shared           *SharedLocation           `json:"shared,omitempty"`
	Vsys             *VsysLocation             `json:"vsys,omitempty"`
}

type DeviceGroupLocation struct {
	DeviceGroup    string `json:"device_group"`
	PanoramaDevice string `json:"panorama_device"`
	Rulebase       string `json:"rulebase"`
}

type FromPanoramaVsysLocation struct {
	Vsys string `json:"vsys"`
}

type SharedLocation struct {
	Rulebase string `json:"rulebase"`
}

type VsysLocation struct {
	NgfwDevice string `json:"ngfw_device"`
	Rulebase   string `json:"rulebase"`
	Vsys       string `json:"vsys"`
}

func NewDeviceGroupLocation() *Location {
	return &Location{DeviceGroup: &DeviceGroupLocation{
		DeviceGroup:    "",
		PanoramaDevice: "localhost.localdomain",
		Rulebase:       "pre-rulebase",
	},
	}
}
func NewFromPanoramaVsysLocation() *Location {
	return &Location{FromPanoramaVsys: &FromPanoramaVsysLocation{
		Vsys: "vsys1",
	},
	}
}
func NewSharedLocation() *Location {
	return &Location{Shared: &SharedLocation{
		Rulebase: "pre-rulebase",
	},
	}
}
func NewVsysLocation() *Location {
	return &Location{Vsys: &VsysLocation{
		NgfwDevice: "localhost.localdomain",
		Rulebase:   "pre-rulebase",
		Vsys:       "vsys1",
	},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.DeviceGroup != nil:
		if o.DeviceGroup.DeviceGroup == "" {
			return fmt.Errorf("DeviceGroup is unspecified")
		}
		if o.DeviceGroup.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.DeviceGroup.Rulebase == "" {
			return fmt.Errorf("Rulebase is unspecified")
		}
		count++
	case o.FromPanoramaVsys != nil:
		if o.FromPanoramaVsys.Vsys == "" {
			return fmt.Errorf("Vsys is unspecified")
		}
		count++
	case o.Shared != nil:
		if o.Shared.Rulebase == "" {
			return fmt.Errorf("Rulebase is unspecified")
		}
		count++
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Rulebase == "" {
			return fmt.Errorf("Rulebase is unspecified")
		}
		if o.Vsys.Vsys == "" {
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

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.DeviceGroup != nil:
		if o.DeviceGroup.DeviceGroup == "" {
			return nil, fmt.Errorf("DeviceGroup is unspecified")
		}
		if o.DeviceGroup.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.DeviceGroup.Rulebase == "" {
			return nil, fmt.Errorf("Rulebase is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.DeviceGroup.PanoramaDevice}),
			"device-group",
			util.AsEntryXpath([]string{o.DeviceGroup.DeviceGroup}),
			o.DeviceGroup.Rulebase,
		}
	case o.FromPanoramaVsys != nil:
		if o.FromPanoramaVsys.Vsys == "" {
			return nil, fmt.Errorf("Vsys is unspecified")
		}
		ans = []string{
			"config",
			"panorama",
			"vsys",
			util.AsEntryXpath([]string{o.FromPanoramaVsys.Vsys}),
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
	case o.Vsys != nil:
		if o.Vsys.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Vsys.Rulebase == "" {
			return nil, fmt.Errorf("Rulebase is unspecified")
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
			"rulebase",
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
