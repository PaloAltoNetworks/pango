package secgroup

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
	Name             string
	DataFiltering    []string
	DisableOverride  *string
	FileBlocking     []string
	Gtp              []string
	Sctp             []string
	Spyware          []string
	UrlFiltering     []string
	Virus            []string
	Vulnerability    []string
	WildfireAnalysis []string

	Misc map[string][]generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName          xml.Name         `xml:"entry"`
	Name             string           `xml:"name,attr"`
	DataFiltering    *util.MemberType `xml:"data-filtering,omitempty"`
	DisableOverride  *string          `xml:"disable-override,omitempty"`
	FileBlocking     *util.MemberType `xml:"file-blocking,omitempty"`
	Gtp              *util.MemberType `xml:"gtp,omitempty"`
	Sctp             *util.MemberType `xml:"sctp,omitempty"`
	Spyware          *util.MemberType `xml:"spyware,omitempty"`
	UrlFiltering     *util.MemberType `xml:"url-filtering,omitempty"`
	Virus            *util.MemberType `xml:"virus,omitempty"`
	Vulnerability    *util.MemberType `xml:"vulnerability,omitempty"`
	WildfireAnalysis *util.MemberType `xml:"wildfire-analysis,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "data_filtering" || v == "DataFiltering" {
		return e.DataFiltering, nil
	}
	if v == "data_filtering|LENGTH" || v == "DataFiltering|LENGTH" {
		return int64(len(e.DataFiltering)), nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "file_blocking" || v == "FileBlocking" {
		return e.FileBlocking, nil
	}
	if v == "file_blocking|LENGTH" || v == "FileBlocking|LENGTH" {
		return int64(len(e.FileBlocking)), nil
	}
	if v == "gtp" || v == "Gtp" {
		return e.Gtp, nil
	}
	if v == "gtp|LENGTH" || v == "Gtp|LENGTH" {
		return int64(len(e.Gtp)), nil
	}
	if v == "sctp" || v == "Sctp" {
		return e.Sctp, nil
	}
	if v == "sctp|LENGTH" || v == "Sctp|LENGTH" {
		return int64(len(e.Sctp)), nil
	}
	if v == "spyware" || v == "Spyware" {
		return e.Spyware, nil
	}
	if v == "spyware|LENGTH" || v == "Spyware|LENGTH" {
		return int64(len(e.Spyware)), nil
	}
	if v == "url_filtering" || v == "UrlFiltering" {
		return e.UrlFiltering, nil
	}
	if v == "url_filtering|LENGTH" || v == "UrlFiltering|LENGTH" {
		return int64(len(e.UrlFiltering)), nil
	}
	if v == "virus" || v == "Virus" {
		return e.Virus, nil
	}
	if v == "virus|LENGTH" || v == "Virus|LENGTH" {
		return int64(len(e.Virus)), nil
	}
	if v == "vulnerability" || v == "Vulnerability" {
		return e.Vulnerability, nil
	}
	if v == "vulnerability|LENGTH" || v == "Vulnerability|LENGTH" {
		return int64(len(e.Vulnerability)), nil
	}
	if v == "wildfire_analysis" || v == "WildfireAnalysis" {
		return e.WildfireAnalysis, nil
	}
	if v == "wildfire_analysis|LENGTH" || v == "WildfireAnalysis|LENGTH" {
		return int64(len(e.WildfireAnalysis)), nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.DataFiltering = util.StrToMem(o.DataFiltering)
	entry.DisableOverride = o.DisableOverride
	entry.FileBlocking = util.StrToMem(o.FileBlocking)
	entry.Gtp = util.StrToMem(o.Gtp)
	entry.Sctp = util.StrToMem(o.Sctp)
	entry.Spyware = util.StrToMem(o.Spyware)
	entry.UrlFiltering = util.StrToMem(o.UrlFiltering)
	entry.Virus = util.StrToMem(o.Virus)
	entry.Vulnerability = util.StrToMem(o.Vulnerability)
	entry.WildfireAnalysis = util.StrToMem(o.WildfireAnalysis)

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
		entry.DataFiltering = util.MemToStr(o.DataFiltering)
		entry.DisableOverride = o.DisableOverride
		entry.FileBlocking = util.MemToStr(o.FileBlocking)
		entry.Gtp = util.MemToStr(o.Gtp)
		entry.Sctp = util.MemToStr(o.Sctp)
		entry.Spyware = util.MemToStr(o.Spyware)
		entry.UrlFiltering = util.MemToStr(o.UrlFiltering)
		entry.Virus = util.MemToStr(o.Virus)
		entry.Vulnerability = util.MemToStr(o.Vulnerability)
		entry.WildfireAnalysis = util.MemToStr(o.WildfireAnalysis)

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
	if !util.OrderedListsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch(a.FileBlocking, b.FileBlocking) {
		return false
	}
	if !util.OrderedListsMatch(a.Gtp, b.Gtp) {
		return false
	}
	if !util.OrderedListsMatch(a.Sctp, b.Sctp) {
		return false
	}
	if !util.OrderedListsMatch(a.Spyware, b.Spyware) {
		return false
	}
	if !util.OrderedListsMatch(a.UrlFiltering, b.UrlFiltering) {
		return false
	}
	if !util.OrderedListsMatch(a.Virus, b.Virus) {
		return false
	}
	if !util.OrderedListsMatch(a.Vulnerability, b.Vulnerability) {
		return false
	}
	if !util.OrderedListsMatch(a.WildfireAnalysis, b.WildfireAnalysis) {
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
