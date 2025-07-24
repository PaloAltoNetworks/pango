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
	suffix = []string{"profile-group", "$name"}
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
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
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
	Misc             []generic.Xml    `xml:",any"`
	MiscAttributes   []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.DataFiltering != nil {
		o.DataFiltering = util.StrToMem(s.DataFiltering)
	}
	o.DisableOverride = s.DisableOverride
	if s.FileBlocking != nil {
		o.FileBlocking = util.StrToMem(s.FileBlocking)
	}
	if s.Gtp != nil {
		o.Gtp = util.StrToMem(s.Gtp)
	}
	if s.Sctp != nil {
		o.Sctp = util.StrToMem(s.Sctp)
	}
	if s.Spyware != nil {
		o.Spyware = util.StrToMem(s.Spyware)
	}
	if s.UrlFiltering != nil {
		o.UrlFiltering = util.StrToMem(s.UrlFiltering)
	}
	if s.Virus != nil {
		o.Virus = util.StrToMem(s.Virus)
	}
	if s.Vulnerability != nil {
		o.Vulnerability = util.StrToMem(s.Vulnerability)
	}
	if s.WildfireAnalysis != nil {
		o.WildfireAnalysis = util.StrToMem(s.WildfireAnalysis)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var dataFilteringVal []string
	if o.DataFiltering != nil {
		dataFilteringVal = util.MemToStr(o.DataFiltering)
	}
	var fileBlockingVal []string
	if o.FileBlocking != nil {
		fileBlockingVal = util.MemToStr(o.FileBlocking)
	}
	var gtpVal []string
	if o.Gtp != nil {
		gtpVal = util.MemToStr(o.Gtp)
	}
	var sctpVal []string
	if o.Sctp != nil {
		sctpVal = util.MemToStr(o.Sctp)
	}
	var spywareVal []string
	if o.Spyware != nil {
		spywareVal = util.MemToStr(o.Spyware)
	}
	var urlFilteringVal []string
	if o.UrlFiltering != nil {
		urlFilteringVal = util.MemToStr(o.UrlFiltering)
	}
	var virusVal []string
	if o.Virus != nil {
		virusVal = util.MemToStr(o.Virus)
	}
	var vulnerabilityVal []string
	if o.Vulnerability != nil {
		vulnerabilityVal = util.MemToStr(o.Vulnerability)
	}
	var wildfireAnalysisVal []string
	if o.WildfireAnalysis != nil {
		wildfireAnalysisVal = util.MemToStr(o.WildfireAnalysis)
	}

	result := &Entry{
		Name:             o.Name,
		DataFiltering:    dataFilteringVal,
		DisableOverride:  o.DisableOverride,
		FileBlocking:     fileBlockingVal,
		Gtp:              gtpVal,
		Sctp:             sctpVal,
		Spyware:          spywareVal,
		UrlFiltering:     urlFilteringVal,
		Virus:            virusVal,
		Vulnerability:    vulnerabilityVal,
		WildfireAnalysis: wildfireAnalysisVal,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
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
	if !util.OrderedListsMatch[string](o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch[string](o.FileBlocking, other.FileBlocking) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Gtp, other.Gtp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Sctp, other.Sctp) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Spyware, other.Spyware) {
		return false
	}
	if !util.OrderedListsMatch[string](o.UrlFiltering, other.UrlFiltering) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Virus, other.Virus) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vulnerability, other.Vulnerability) {
		return false
	}
	if !util.OrderedListsMatch[string](o.WildfireAnalysis, other.WildfireAnalysis) {
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
