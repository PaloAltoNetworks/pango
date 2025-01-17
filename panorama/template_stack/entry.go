package template_stack

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
	Suffix = []string{"template-stack"}
)

type Entry struct {
	Name            string
	DefaultVsys     *string
	Description     *string
	Devices         []string
	Templates       []string
	UserGroupSource *UserGroupSource

	Misc map[string][]generic.Xml
}

type UserGroupSource struct {
	MasterDevice *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName         xml.Name            `xml:"entry"`
	Name            string              `xml:"name,attr"`
	DefaultVsys     *string             `xml:"settings>default-vsys,omitempty"`
	Description     *string             `xml:"description,omitempty"`
	Devices         *util.EntryType     `xml:"devices,omitempty"`
	Templates       *util.MemberType    `xml:"templates,omitempty"`
	UserGroupSource *UserGroupSourceXml `xml:"user-group-source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UserGroupSourceXml struct {
	MasterDevice *string `xml:"master-device,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "default_vsys" || v == "DefaultVsys" {
		return e.DefaultVsys, nil
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
	if v == "user_group_source" || v == "UserGroupSource" {
		return e.UserGroupSource, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.DefaultVsys = o.DefaultVsys
	entry.Description = o.Description
	entry.Devices = util.StrToEnt(o.Devices)
	entry.Templates = util.StrToMem(o.Templates)
	var nestedUserGroupSource *UserGroupSourceXml
	if o.UserGroupSource != nil {
		nestedUserGroupSource = &UserGroupSourceXml{}
		if _, ok := o.Misc["UserGroupSource"]; ok {
			nestedUserGroupSource.Misc = o.Misc["UserGroupSource"]
		}
		if o.UserGroupSource.MasterDevice != nil {
			nestedUserGroupSource.MasterDevice = o.UserGroupSource.MasterDevice
		}
	}
	entry.UserGroupSource = nestedUserGroupSource

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
		entry.DefaultVsys = o.DefaultVsys
		entry.Description = o.Description
		entry.Devices = util.EntToStr(o.Devices)
		entry.Templates = util.MemToStr(o.Templates)
		var nestedUserGroupSource *UserGroupSource
		if o.UserGroupSource != nil {
			nestedUserGroupSource = &UserGroupSource{}
			if o.UserGroupSource.Misc != nil {
				entry.Misc["UserGroupSource"] = o.UserGroupSource.Misc
			}
			if o.UserGroupSource.MasterDevice != nil {
				nestedUserGroupSource.MasterDevice = o.UserGroupSource.MasterDevice
			}
		}
		entry.UserGroupSource = nestedUserGroupSource

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
	if !util.StringsMatch(a.DefaultVsys, b.DefaultVsys) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.Devices, b.Devices) {
		return false
	}
	if !util.OrderedListsMatch(a.Templates, b.Templates) {
		return false
	}
	if !matchUserGroupSource(a.UserGroupSource, b.UserGroupSource) {
		return false
	}

	return true
}

func matchUserGroupSource(a *UserGroupSource, b *UserGroupSource) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.MasterDevice, b.MasterDevice) {
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
