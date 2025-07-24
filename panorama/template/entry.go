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
	suffix = []string{"template", "$name"}
)

type Entry struct {
	Name           string
	Description    *string
	DefaultVsys    *string
	Config         *Config
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Config struct {
	Devices        []ConfigDevices
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConfigDevices struct {
	Name           string
	Vsys           []ConfigDevicesVsys
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConfigDevicesVsys struct {
	Name           string
	Import         *ConfigDevicesVsysImport
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConfigDevicesVsysImport struct {
	Network        *ConfigDevicesVsysImportNetwork
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ConfigDevicesVsysImportNetwork struct {
	Interfaces     []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Description    *string       `xml:"description,omitempty"`
	DefaultVsys    *string       `xml:"settings>default-vsys,omitempty"`
	Config         *configXml    `xml:"config,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type configXml struct {
	Devices        *configDevicesContainerXml `xml:"devices,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type configDevicesContainerXml struct {
	Entries []configDevicesXml `xml:"entry"`
}
type configDevicesXml struct {
	XMLName        xml.Name                       `xml:"entry"`
	Name           string                         `xml:"name,attr"`
	Vsys           *configDevicesVsysContainerXml `xml:"vsys,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type configDevicesVsysContainerXml struct {
	Entries []configDevicesVsysXml `xml:"entry"`
}
type configDevicesVsysXml struct {
	XMLName        xml.Name                    `xml:"entry"`
	Name           string                      `xml:"name,attr"`
	Import         *configDevicesVsysImportXml `xml:"import,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type configDevicesVsysImportXml struct {
	Network        *configDevicesVsysImportNetworkXml `xml:"network,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type configDevicesVsysImportNetworkXml struct {
	Interfaces     *util.MemberType `xml:"interface,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	o.DefaultVsys = s.DefaultVsys
	if s.Config != nil {
		var obj configXml
		obj.MarshalFromObject(*s.Config)
		o.Config = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var configVal *Config
	if o.Config != nil {
		obj, err := o.Config.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		configVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Description:    o.Description,
		DefaultVsys:    o.DefaultVsys,
		Config:         configVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *configXml) MarshalFromObject(s Config) {
	if s.Devices != nil {
		var objs []configDevicesXml
		for _, elt := range s.Devices {
			var obj configDevicesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Devices = &configDevicesContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var devicesVal []ConfigDevices
	if o.Devices != nil {
		for _, elt := range o.Devices.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			devicesVal = append(devicesVal, *obj)
		}
	}

	result := &Config{
		Devices:        devicesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *configDevicesXml) MarshalFromObject(s ConfigDevices) {
	o.Name = s.Name
	if s.Vsys != nil {
		var objs []configDevicesVsysXml
		for _, elt := range s.Vsys {
			var obj configDevicesVsysXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vsys = &configDevicesVsysContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configDevicesXml) UnmarshalToObject() (*ConfigDevices, error) {
	var vsysVal []ConfigDevicesVsys
	if o.Vsys != nil {
		for _, elt := range o.Vsys.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vsysVal = append(vsysVal, *obj)
		}
	}

	result := &ConfigDevices{
		Name:           o.Name,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *configDevicesVsysXml) MarshalFromObject(s ConfigDevicesVsys) {
	o.Name = s.Name
	if s.Import != nil {
		var obj configDevicesVsysImportXml
		obj.MarshalFromObject(*s.Import)
		o.Import = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configDevicesVsysXml) UnmarshalToObject() (*ConfigDevicesVsys, error) {
	var importVal *ConfigDevicesVsysImport
	if o.Import != nil {
		obj, err := o.Import.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		importVal = obj
	}

	result := &ConfigDevicesVsys{
		Name:           o.Name,
		Import:         importVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *configDevicesVsysImportXml) MarshalFromObject(s ConfigDevicesVsysImport) {
	if s.Network != nil {
		var obj configDevicesVsysImportNetworkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configDevicesVsysImportXml) UnmarshalToObject() (*ConfigDevicesVsysImport, error) {
	var networkVal *ConfigDevicesVsysImportNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}

	result := &ConfigDevicesVsysImport{
		Network:        networkVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *configDevicesVsysImportNetworkXml) MarshalFromObject(s ConfigDevicesVsysImportNetwork) {
	if s.Interfaces != nil {
		o.Interfaces = util.StrToMem(s.Interfaces)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o configDevicesVsysImportNetworkXml) UnmarshalToObject() (*ConfigDevicesVsysImportNetwork, error) {
	var interfacesVal []string
	if o.Interfaces != nil {
		interfacesVal = util.MemToStr(o.Interfaces)
	}

	result := &ConfigDevicesVsysImportNetwork{
		Interfaces:     interfacesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "default_vsys" || v == "DefaultVsys" {
		return e.DefaultVsys, nil
	}
	if v == "config" || v == "Config" {
		return e.Config, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DefaultVsys, other.DefaultVsys) {
		return false
	}
	if !o.Config.matches(other.Config) {
		return false
	}

	return true
}

func (o *Config) matches(other *Config) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Devices) != len(other.Devices) {
		return false
	}
	for idx := range o.Devices {
		if !o.Devices[idx].matches(&other.Devices[idx]) {
			return false
		}
	}

	return true
}

func (o *ConfigDevices) matches(other *ConfigDevices) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if len(o.Vsys) != len(other.Vsys) {
		return false
	}
	for idx := range o.Vsys {
		if !o.Vsys[idx].matches(&other.Vsys[idx]) {
			return false
		}
	}

	return true
}

func (o *ConfigDevicesVsys) matches(other *ConfigDevicesVsys) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Import.matches(other.Import) {
		return false
	}

	return true
}

func (o *ConfigDevicesVsysImport) matches(other *ConfigDevicesVsysImport) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}

	return true
}

func (o *ConfigDevicesVsysImportNetwork) matches(other *ConfigDevicesVsysImportNetwork) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interfaces, other.Interfaces) {
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

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
