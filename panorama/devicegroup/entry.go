package devicegroup

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
	Suffix = []string{"device-group"}
)

type Entry struct {
	Name              string
	AuthorizationCode *string
	Description       *string
	Devices           []Devices
	Templates         []string

	Misc map[string][]generic.Xml
}

type Devices struct {
	Name string
	Vsys []string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName           xml.Name         `xml:"entry"`
	Name              string           `xml:"name,attr"`
	AuthorizationCode *string          `xml:"authorization-code,omitempty"`
	Description       *string          `xml:"description,omitempty"`
	Devices           []DevicesXml     `xml:"devices>entry,omitempty"`
	Templates         *util.MemberType `xml:"reference-templates,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type DevicesXml struct {
	XMLName xml.Name         `xml:"entry"`
	Name    string           `xml:"name,attr"`
	Vsys    *util.MemberType `xml:"vsys,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authorization_code" || v == "AuthorizationCode" {
		return e.AuthorizationCode, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "devices" || v == "Devices" {
		return e.Devices, nil
	}
	if v == "devices|LENGTH" || v == "Devices|LENGTH" {
		return int64(len(e.Devices)), nil
	}
	if v == "templates" || v == "Templates" {
		return e.Templates, nil
	}
	if v == "templates|LENGTH" || v == "Templates|LENGTH" {
		return int64(len(e.Templates)), nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.AuthorizationCode = o.AuthorizationCode
	entry.Description = o.Description
	var nestedDevicesCol []DevicesXml
	if o.Devices != nil {
		nestedDevicesCol = []DevicesXml{}
		for _, oDevices := range o.Devices {
			nestedDevices := DevicesXml{}
			if _, ok := o.Misc["Devices"]; ok {
				nestedDevices.Misc = o.Misc["Devices"]
			}
			if oDevices.Vsys != nil {
				nestedDevices.Vsys = util.StrToMem(oDevices.Vsys)
			}
			if oDevices.Name != "" {
				nestedDevices.Name = oDevices.Name
			}
			nestedDevicesCol = append(nestedDevicesCol, nestedDevices)
		}
		entry.Devices = nestedDevicesCol
	}

	entry.Templates = util.StrToMem(o.Templates)

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
		entry.AuthorizationCode = o.AuthorizationCode
		entry.Description = o.Description
		var nestedDevicesCol []Devices
		if o.Devices != nil {
			nestedDevicesCol = []Devices{}
			for _, oDevices := range o.Devices {
				nestedDevices := Devices{}
				if oDevices.Misc != nil {
					entry.Misc["Devices"] = oDevices.Misc
				}
				if oDevices.Vsys != nil {
					nestedDevices.Vsys = util.MemToStr(oDevices.Vsys)
				}
				if oDevices.Name != "" {
					nestedDevices.Name = oDevices.Name
				}
				nestedDevicesCol = append(nestedDevicesCol, nestedDevices)
			}
			entry.Devices = nestedDevicesCol
		}

		entry.Templates = util.MemToStr(o.Templates)

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
	if !util.StringsMatch(a.AuthorizationCode, b.AuthorizationCode) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !matchDevices(a.Devices, b.Devices) {
		return false
	}
	if !util.OrderedListsMatch(a.Templates, b.Templates) {
		return false
	}

	return true
}

func matchDevices(a []Devices, b []Devices) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.OrderedListsMatch(a.Vsys, b.Vsys) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
