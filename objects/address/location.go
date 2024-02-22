package address

import (
    "fmt"

    "github.com/PaloAltoNetworks/pango/errors"
    "github.com/PaloAltoNetworks/pango/util"
    "github.com/PaloAltoNetworks/pango/version"
)

type Location struct {
    Shared bool
    Vsys *VsysLocation
    DeviceGroup *DeviceGroupLocation
    FromPanorama bool
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
        if o.DeviceGroup.PanoramaDevice == "" {
            return nil, fmt.Errorf("PanoramaDevice is unspecified")
        }
        if o.DeviceGroup.DeviceGroup == "" {
            return nil, fmt.Errorf("DeviceGroup is unspecified")
        }
        ans = []string{
            "config",
            "devices",
            util.AsEntryXpath([]string{o.DeviceGroup.PanoramaDevice}),
            "device-group",
            util.AsEntryXpath([]string{o.DeviceGroup.DeviceGroup}),
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
    NgfwDevice string
    Vsys string
}

type DeviceGroupLocation struct {
    PanoramaDevice string
    DeviceGroup string
}
