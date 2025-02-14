package template

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	Suffix = []string{}
)

type Entry struct {
	Name        string
	Config      *Config
	DefaultVsys *string
	Description *string

	Misc map[string][]generic.Xml
}

type Config struct {
	Devices []ConfigDevices
}
type ConfigDevices struct {
	Name string
	Vsys []ConfigDevicesVsys
}
type ConfigDevicesVsys struct {
	Import *ConfigDevicesVsysImport
	Name   string
}
type ConfigDevicesVsysImport struct {
	Network *ConfigDevicesVsysImportNetwork
}
type ConfigDevicesVsysImportNetwork struct {
	Interfaces []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name   `xml:"entry"`
	Name        string     `xml:"name,attr"`
	Config      *ConfigXml `xml:"config,omitempty"`
	DefaultVsys *string    `xml:"settings>default-vsys,omitempty"`
	Description *string    `xml:"description,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ConfigXml struct {
	Devices []ConfigDevicesXml `xml:"devices>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ConfigDevicesXml struct {
	XMLName xml.Name               `xml:"entry"`
	Name    string                 `xml:"name,attr"`
	Vsys    []ConfigDevicesVsysXml `xml:"vsys>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ConfigDevicesVsysXml struct {
	Import  *ConfigDevicesVsysImportXml `xml:"import,omitempty"`
	XMLName xml.Name                    `xml:"entry"`
	Name    string                      `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ConfigDevicesVsysImportXml struct {
	Network *ConfigDevicesVsysImportNetworkXml `xml:"network,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ConfigDevicesVsysImportNetworkXml struct {
	Interfaces *util.MemberType `xml:"interface,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "config" || v == "Config" {
		return e.Config, nil
	}
	if v == "default_vsys" || v == "DefaultVsys" {
		return e.DefaultVsys, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedConfig *ConfigXml
	if o.Config != nil {
		nestedConfig = &ConfigXml{}
		if _, ok := o.Misc["Config"]; ok {
			nestedConfig.Misc = o.Misc["Config"]
		}
		if o.Config.Devices != nil {
			nestedConfig.Devices = []ConfigDevicesXml{}
			for _, oConfigDevices := range o.Config.Devices {
				nestedConfigDevices := ConfigDevicesXml{}
				if _, ok := o.Misc["ConfigDevices"]; ok {
					nestedConfigDevices.Misc = o.Misc["ConfigDevices"]
				}
				if oConfigDevices.Name != "" {
					nestedConfigDevices.Name = oConfigDevices.Name
				}
				if oConfigDevices.Vsys != nil {
					nestedConfigDevices.Vsys = []ConfigDevicesVsysXml{}
					for _, oConfigDevicesVsys := range oConfigDevices.Vsys {
						nestedConfigDevicesVsys := ConfigDevicesVsysXml{}
						if _, ok := o.Misc["ConfigDevicesVsys"]; ok {
							nestedConfigDevicesVsys.Misc = o.Misc["ConfigDevicesVsys"]
						}
						if oConfigDevicesVsys.Import != nil {
							nestedConfigDevicesVsys.Import = &ConfigDevicesVsysImportXml{}
							if _, ok := o.Misc["ConfigDevicesVsysImport"]; ok {
								nestedConfigDevicesVsys.Import.Misc = o.Misc["ConfigDevicesVsysImport"]
							}
							if oConfigDevicesVsys.Import.Network != nil {
								nestedConfigDevicesVsys.Import.Network = &ConfigDevicesVsysImportNetworkXml{}
								if _, ok := o.Misc["ConfigDevicesVsysImportNetwork"]; ok {
									nestedConfigDevicesVsys.Import.Network.Misc = o.Misc["ConfigDevicesVsysImportNetwork"]
								}
								if oConfigDevicesVsys.Import.Network.Interfaces != nil {
									nestedConfigDevicesVsys.Import.Network.Interfaces = util.StrToMem(oConfigDevicesVsys.Import.Network.Interfaces)
								}
							}
						}
						if oConfigDevicesVsys.Name != "" {
							nestedConfigDevicesVsys.Name = oConfigDevicesVsys.Name
						}
						nestedConfigDevices.Vsys = append(nestedConfigDevices.Vsys, nestedConfigDevicesVsys)
					}
				}
				nestedConfig.Devices = append(nestedConfig.Devices, nestedConfigDevices)
			}
		}
	}
	entry.Config = nestedConfig

	entry.DefaultVsys = o.DefaultVsys
	entry.Description = o.Description

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func (c *entryXmlContainer) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedConfig *Config
		if o.Config != nil {
			nestedConfig = &Config{}
			if o.Config.Misc != nil {
				entry.Misc["Config"] = o.Config.Misc
			}
			if o.Config.Devices != nil {
				nestedConfig.Devices = []ConfigDevices{}
				for _, oConfigDevices := range o.Config.Devices {
					nestedConfigDevices := ConfigDevices{}
					if oConfigDevices.Misc != nil {
						entry.Misc["ConfigDevices"] = oConfigDevices.Misc
					}
					if oConfigDevices.Vsys != nil {
						nestedConfigDevices.Vsys = []ConfigDevicesVsys{}
						for _, oConfigDevicesVsys := range oConfigDevices.Vsys {
							nestedConfigDevicesVsys := ConfigDevicesVsys{}
							if oConfigDevicesVsys.Misc != nil {
								entry.Misc["ConfigDevicesVsys"] = oConfigDevicesVsys.Misc
							}
							if oConfigDevicesVsys.Name != "" {
								nestedConfigDevicesVsys.Name = oConfigDevicesVsys.Name
							}
							if oConfigDevicesVsys.Import != nil {
								nestedConfigDevicesVsys.Import = &ConfigDevicesVsysImport{}
								if oConfigDevicesVsys.Import.Misc != nil {
									entry.Misc["ConfigDevicesVsysImport"] = oConfigDevicesVsys.Import.Misc
								}
								if oConfigDevicesVsys.Import.Network != nil {
									nestedConfigDevicesVsys.Import.Network = &ConfigDevicesVsysImportNetwork{}
									if oConfigDevicesVsys.Import.Network.Misc != nil {
										entry.Misc["ConfigDevicesVsysImportNetwork"] = oConfigDevicesVsys.Import.Network.Misc
									}
									if oConfigDevicesVsys.Import.Network.Interfaces != nil {
										nestedConfigDevicesVsys.Import.Network.Interfaces = util.MemToStr(oConfigDevicesVsys.Import.Network.Interfaces)
									}
								}
							}
							nestedConfigDevices.Vsys = append(nestedConfigDevices.Vsys, nestedConfigDevicesVsys)
						}
					}
					if oConfigDevices.Name != "" {
						nestedConfigDevices.Name = oConfigDevices.Name
					}
					nestedConfig.Devices = append(nestedConfig.Devices, nestedConfigDevices)
				}
			}
		}
		entry.Config = nestedConfig

		entry.DefaultVsys = o.DefaultVsys
		entry.Description = o.Description

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func SpecMatches(a, b *Entry) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !matchConfig(a.Config, b.Config) {
		return false
	}
	if !util.StringsMatch(a.DefaultVsys, b.DefaultVsys) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}

	return true
}

func matchConfigDevicesVsysImportNetwork(a *ConfigDevicesVsysImportNetwork, b *ConfigDevicesVsysImportNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.OrderedListsMatch(a.Interfaces, b.Interfaces) {
		return false
	}
	return true
}
func matchConfigDevicesVsysImport(a *ConfigDevicesVsysImport, b *ConfigDevicesVsysImport) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchConfigDevicesVsysImportNetwork(a.Network, b.Network) {
		return false
	}
	return true
}
func matchConfigDevicesVsys(a []ConfigDevicesVsys, b []ConfigDevicesVsys) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchConfigDevicesVsysImport(a.Import, b.Import) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchConfigDevices(a []ConfigDevices, b []ConfigDevices) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchConfigDevicesVsys(a.Vsys, b.Vsys) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchConfig(a *Config, b *Config) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchConfigDevices(a.Devices, b.Devices) {
		return false
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
