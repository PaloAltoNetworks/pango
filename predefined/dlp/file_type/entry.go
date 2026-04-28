package file_type

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
	suffix = []string{"file-type", "$name"}
)

type Entry struct {
	Name           string
	FileProperty   []FileProperty
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type FileProperty struct {
	Name           string
	Label          *string
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
	XMLName        xml.Name                  `xml:"entry"`
	Name           string                    `xml:"name,attr"`
	FileProperty   *filePropertyContainerXml `xml:"file-property,omitempty"`
	Misc           []generic.Xml             `xml:",any"`
	MiscAttributes []xml.Attr                `xml:",any,attr"`
}
type filePropertyContainerXml struct {
	Entries []filePropertyXml `xml:"entry"`
}
type filePropertyXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Label          *string       `xml:"label,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.FileProperty != nil {
		var objs []filePropertyXml
		for _, elt := range s.FileProperty {
			var obj filePropertyXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.FileProperty = &filePropertyContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var filePropertyVal []FileProperty
	if o.FileProperty != nil {
		for _, elt := range o.FileProperty.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			filePropertyVal = append(filePropertyVal, *obj)
		}
	}

	result := &Entry{
		Name:           o.Name,
		FileProperty:   filePropertyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *filePropertyXml) MarshalFromObject(s FileProperty) {
	o.Name = s.Name
	o.Label = s.Label
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o filePropertyXml) UnmarshalToObject() (*FileProperty, error) {

	result := &FileProperty{
		Name:           o.Name,
		Label:          o.Label,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "file_property" || v == "FileProperty" {
		return e.FileProperty, nil
	}
	if v == "file_property|LENGTH" || v == "FileProperty|LENGTH" {
		return int64(len(e.FileProperty)), nil
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
	if len(o.FileProperty) != len(other.FileProperty) {
		return false
	}
	for idx := range o.FileProperty {
		if !o.FileProperty[idx].matches(&other.FileProperty[idx]) {
			return false
		}
	}

	return true
}

func (o *FileProperty) matches(other *FileProperty) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Label, other.Label) {
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
