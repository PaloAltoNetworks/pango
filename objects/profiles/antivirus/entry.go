package antivirus

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
	Suffix = []string{"virus"}
)

type Entry struct {
	Name                       string
	Application                []Application
	Decoder                    []Decoder
	Description                *string
	DisableOverride            *string
	MlavEngineFilebasedEnabled []MlavEngineFilebasedEnabled
	MlavException              []MlavException
	PacketCapture              *bool
	ThreatException            []ThreatException
	WfrtHoldMode               *bool

	Misc map[string][]generic.Xml
}

type Application struct {
	Name   string
	Action *string
}
type Decoder struct {
	Name           string
	Action         *string
	WildfireAction *string
	MlavAction     *string
}
type MlavEngineFilebasedEnabled struct {
	Name             string
	MlavPolicyAction *string
}
type MlavException struct {
	Name        string
	Filename    *string
	Description *string
}
type ThreatException struct {
	Name string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}
type entryXml struct {
	XMLName                    xml.Name                        `xml:"entry"`
	Name                       string                          `xml:"name,attr"`
	Application                []ApplicationXml                `xml:"application>entry,omitempty"`
	Decoder                    []DecoderXml                    `xml:"decoder>entry,omitempty"`
	Description                *string                         `xml:"description,omitempty"`
	DisableOverride            *string                         `xml:"disable-override,omitempty"`
	MlavEngineFilebasedEnabled []MlavEngineFilebasedEnabledXml `xml:"mlav-engine-filebased-enabled>entry,omitempty"`
	MlavException              []MlavExceptionXml              `xml:"mlav-exception>entry,omitempty"`
	PacketCapture              *string                         `xml:"packet-capture,omitempty"`
	ThreatException            []ThreatExceptionXml            `xml:"threat-exception>entry,omitempty"`
	WfrtHoldMode               *string                         `xml:"wfrt-hold-mode,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name                               `xml:"entry"`
	Name                       string                                 `xml:"name,attr"`
	Application                []ApplicationXml_11_0_2                `xml:"application>entry,omitempty"`
	Decoder                    []DecoderXml_11_0_2                    `xml:"decoder>entry,omitempty"`
	Description                *string                                `xml:"description,omitempty"`
	DisableOverride            *string                                `xml:"disable-override,omitempty"`
	MlavEngineFilebasedEnabled []MlavEngineFilebasedEnabledXml_11_0_2 `xml:"mlav-engine-filebased-enabled>entry,omitempty"`
	MlavException              []MlavExceptionXml_11_0_2              `xml:"mlav-exception>entry,omitempty"`
	PacketCapture              *string                                `xml:"packet-capture,omitempty"`
	ThreatException            []ThreatExceptionXml_11_0_2            `xml:"threat-exception>entry,omitempty"`
	WfrtHoldMode               *string                                `xml:"wfrt-hold-mode,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ApplicationXml struct {
	Action *string `xml:"action,omitempty"`
	Name   string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type DecoderXml struct {
	Action         *string `xml:"action,omitempty"`
	MlavAction     *string `xml:"mlav-action,omitempty"`
	Name           string  `xml:"name,attr"`
	WildfireAction *string `xml:"wildfire-action,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MlavEngineFilebasedEnabledXml struct {
	MlavPolicyAction *string `xml:"mlav-policy-action,omitempty"`
	Name             string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MlavExceptionXml struct {
	Description *string `xml:"description,omitempty"`
	Filename    *string `xml:"filename,omitempty"`
	Name        string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionXml struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ApplicationXml_11_0_2 struct {
	Action *string `xml:"action,omitempty"`
	Name   string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type DecoderXml_11_0_2 struct {
	Action         *string `xml:"action,omitempty"`
	MlavAction     *string `xml:"mlav-action,omitempty"`
	Name           string  `xml:"name,attr"`
	WildfireAction *string `xml:"wildfire-action,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type MlavEngineFilebasedEnabledXml_11_0_2 struct {
	MlavPolicyAction *string `xml:"mlav-policy-action,omitempty"`
	Name             string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MlavExceptionXml_11_0_2 struct {
	Description *string `xml:"description,omitempty"`
	Filename    *string `xml:"filename,omitempty"`
	Name        string  `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionXml_11_0_2 struct {
	Name string `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "application" || v == "Application" {
		return e.Application, nil
	}
	if v == "application|LENGTH" || v == "Application|LENGTH" {
		return int64(len(e.Application)), nil
	}
	if v == "decoder" || v == "Decoder" {
		return e.Decoder, nil
	}
	if v == "decoder|LENGTH" || v == "Decoder|LENGTH" {
		return int64(len(e.Decoder)), nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "mlav_engine_filebased_enabled" || v == "MlavEngineFilebasedEnabled" {
		return e.MlavEngineFilebasedEnabled, nil
	}
	if v == "mlav_engine_filebased_enabled|LENGTH" || v == "MlavEngineFilebasedEnabled|LENGTH" {
		return int64(len(e.MlavEngineFilebasedEnabled)), nil
	}
	if v == "mlav_exception" || v == "MlavException" {
		return e.MlavException, nil
	}
	if v == "mlav_exception|LENGTH" || v == "MlavException|LENGTH" {
		return int64(len(e.MlavException)), nil
	}
	if v == "packet_capture" || v == "PacketCapture" {
		return e.PacketCapture, nil
	}
	if v == "threat_exception" || v == "ThreatException" {
		return e.ThreatException, nil
	}
	if v == "threat_exception|LENGTH" || v == "ThreatException|LENGTH" {
		return int64(len(e.ThreatException)), nil
	}
	if v == "wfrt_hold_mode" || v == "WfrtHoldMode" {
		return e.WfrtHoldMode, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedApplicationCol []ApplicationXml
	if o.Application != nil {
		nestedApplicationCol = []ApplicationXml{}
		for _, oApplication := range o.Application {
			nestedApplication := ApplicationXml{}
			if _, ok := o.Misc["Application"]; ok {
				nestedApplication.Misc = o.Misc["Application"]
			}
			if oApplication.Name != "" {
				nestedApplication.Name = oApplication.Name
			}
			if oApplication.Action != nil {
				nestedApplication.Action = oApplication.Action
			}
			nestedApplicationCol = append(nestedApplicationCol, nestedApplication)
		}
		entry.Application = nestedApplicationCol
	}

	var nestedDecoderCol []DecoderXml
	if o.Decoder != nil {
		nestedDecoderCol = []DecoderXml{}
		for _, oDecoder := range o.Decoder {
			nestedDecoder := DecoderXml{}
			if _, ok := o.Misc["Decoder"]; ok {
				nestedDecoder.Misc = o.Misc["Decoder"]
			}
			if oDecoder.Name != "" {
				nestedDecoder.Name = oDecoder.Name
			}
			if oDecoder.Action != nil {
				nestedDecoder.Action = oDecoder.Action
			}
			if oDecoder.WildfireAction != nil {
				nestedDecoder.WildfireAction = oDecoder.WildfireAction
			}
			if oDecoder.MlavAction != nil {
				nestedDecoder.MlavAction = oDecoder.MlavAction
			}
			nestedDecoderCol = append(nestedDecoderCol, nestedDecoder)
		}
		entry.Decoder = nestedDecoderCol
	}

	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	var nestedMlavEngineFilebasedEnabledCol []MlavEngineFilebasedEnabledXml
	if o.MlavEngineFilebasedEnabled != nil {
		nestedMlavEngineFilebasedEnabledCol = []MlavEngineFilebasedEnabledXml{}
		for _, oMlavEngineFilebasedEnabled := range o.MlavEngineFilebasedEnabled {
			nestedMlavEngineFilebasedEnabled := MlavEngineFilebasedEnabledXml{}
			if _, ok := o.Misc["MlavEngineFilebasedEnabled"]; ok {
				nestedMlavEngineFilebasedEnabled.Misc = o.Misc["MlavEngineFilebasedEnabled"]
			}
			if oMlavEngineFilebasedEnabled.Name != "" {
				nestedMlavEngineFilebasedEnabled.Name = oMlavEngineFilebasedEnabled.Name
			}
			if oMlavEngineFilebasedEnabled.MlavPolicyAction != nil {
				nestedMlavEngineFilebasedEnabled.MlavPolicyAction = oMlavEngineFilebasedEnabled.MlavPolicyAction
			}
			nestedMlavEngineFilebasedEnabledCol = append(nestedMlavEngineFilebasedEnabledCol, nestedMlavEngineFilebasedEnabled)
		}
		entry.MlavEngineFilebasedEnabled = nestedMlavEngineFilebasedEnabledCol
	}

	var nestedMlavExceptionCol []MlavExceptionXml
	if o.MlavException != nil {
		nestedMlavExceptionCol = []MlavExceptionXml{}
		for _, oMlavException := range o.MlavException {
			nestedMlavException := MlavExceptionXml{}
			if _, ok := o.Misc["MlavException"]; ok {
				nestedMlavException.Misc = o.Misc["MlavException"]
			}
			if oMlavException.Name != "" {
				nestedMlavException.Name = oMlavException.Name
			}
			if oMlavException.Filename != nil {
				nestedMlavException.Filename = oMlavException.Filename
			}
			if oMlavException.Description != nil {
				nestedMlavException.Description = oMlavException.Description
			}
			nestedMlavExceptionCol = append(nestedMlavExceptionCol, nestedMlavException)
		}
		entry.MlavException = nestedMlavExceptionCol
	}

	entry.PacketCapture = util.YesNo(o.PacketCapture, nil)
	var nestedThreatExceptionCol []ThreatExceptionXml
	if o.ThreatException != nil {
		nestedThreatExceptionCol = []ThreatExceptionXml{}
		for _, oThreatException := range o.ThreatException {
			nestedThreatException := ThreatExceptionXml{}
			if _, ok := o.Misc["ThreatException"]; ok {
				nestedThreatException.Misc = o.Misc["ThreatException"]
			}
			if oThreatException.Name != "" {
				nestedThreatException.Name = oThreatException.Name
			}
			nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
		}
		entry.ThreatException = nestedThreatExceptionCol
	}

	entry.WfrtHoldMode = util.YesNo(o.WfrtHoldMode, nil)

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	var nestedApplicationCol []ApplicationXml_11_0_2
	if o.Application != nil {
		nestedApplicationCol = []ApplicationXml_11_0_2{}
		for _, oApplication := range o.Application {
			nestedApplication := ApplicationXml_11_0_2{}
			if _, ok := o.Misc["Application"]; ok {
				nestedApplication.Misc = o.Misc["Application"]
			}
			if oApplication.Name != "" {
				nestedApplication.Name = oApplication.Name
			}
			if oApplication.Action != nil {
				nestedApplication.Action = oApplication.Action
			}
			nestedApplicationCol = append(nestedApplicationCol, nestedApplication)
		}
		entry.Application = nestedApplicationCol
	}

	var nestedDecoderCol []DecoderXml_11_0_2
	if o.Decoder != nil {
		nestedDecoderCol = []DecoderXml_11_0_2{}
		for _, oDecoder := range o.Decoder {
			nestedDecoder := DecoderXml_11_0_2{}
			if _, ok := o.Misc["Decoder"]; ok {
				nestedDecoder.Misc = o.Misc["Decoder"]
			}
			if oDecoder.Name != "" {
				nestedDecoder.Name = oDecoder.Name
			}
			if oDecoder.Action != nil {
				nestedDecoder.Action = oDecoder.Action
			}
			if oDecoder.WildfireAction != nil {
				nestedDecoder.WildfireAction = oDecoder.WildfireAction
			}
			if oDecoder.MlavAction != nil {
				nestedDecoder.MlavAction = oDecoder.MlavAction
			}
			nestedDecoderCol = append(nestedDecoderCol, nestedDecoder)
		}
		entry.Decoder = nestedDecoderCol
	}

	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	var nestedMlavEngineFilebasedEnabledCol []MlavEngineFilebasedEnabledXml_11_0_2
	if o.MlavEngineFilebasedEnabled != nil {
		nestedMlavEngineFilebasedEnabledCol = []MlavEngineFilebasedEnabledXml_11_0_2{}
		for _, oMlavEngineFilebasedEnabled := range o.MlavEngineFilebasedEnabled {
			nestedMlavEngineFilebasedEnabled := MlavEngineFilebasedEnabledXml_11_0_2{}
			if _, ok := o.Misc["MlavEngineFilebasedEnabled"]; ok {
				nestedMlavEngineFilebasedEnabled.Misc = o.Misc["MlavEngineFilebasedEnabled"]
			}
			if oMlavEngineFilebasedEnabled.Name != "" {
				nestedMlavEngineFilebasedEnabled.Name = oMlavEngineFilebasedEnabled.Name
			}
			if oMlavEngineFilebasedEnabled.MlavPolicyAction != nil {
				nestedMlavEngineFilebasedEnabled.MlavPolicyAction = oMlavEngineFilebasedEnabled.MlavPolicyAction
			}
			nestedMlavEngineFilebasedEnabledCol = append(nestedMlavEngineFilebasedEnabledCol, nestedMlavEngineFilebasedEnabled)
		}
		entry.MlavEngineFilebasedEnabled = nestedMlavEngineFilebasedEnabledCol
	}

	var nestedMlavExceptionCol []MlavExceptionXml_11_0_2
	if o.MlavException != nil {
		nestedMlavExceptionCol = []MlavExceptionXml_11_0_2{}
		for _, oMlavException := range o.MlavException {
			nestedMlavException := MlavExceptionXml_11_0_2{}
			if _, ok := o.Misc["MlavException"]; ok {
				nestedMlavException.Misc = o.Misc["MlavException"]
			}
			if oMlavException.Name != "" {
				nestedMlavException.Name = oMlavException.Name
			}
			if oMlavException.Filename != nil {
				nestedMlavException.Filename = oMlavException.Filename
			}
			if oMlavException.Description != nil {
				nestedMlavException.Description = oMlavException.Description
			}
			nestedMlavExceptionCol = append(nestedMlavExceptionCol, nestedMlavException)
		}
		entry.MlavException = nestedMlavExceptionCol
	}

	entry.PacketCapture = util.YesNo(o.PacketCapture, nil)
	var nestedThreatExceptionCol []ThreatExceptionXml_11_0_2
	if o.ThreatException != nil {
		nestedThreatExceptionCol = []ThreatExceptionXml_11_0_2{}
		for _, oThreatException := range o.ThreatException {
			nestedThreatException := ThreatExceptionXml_11_0_2{}
			if _, ok := o.Misc["ThreatException"]; ok {
				nestedThreatException.Misc = o.Misc["ThreatException"]
			}
			if oThreatException.Name != "" {
				nestedThreatException.Name = oThreatException.Name
			}
			nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
		}
		entry.ThreatException = nestedThreatExceptionCol
	}

	entry.WfrtHoldMode = util.YesNo(o.WfrtHoldMode, nil)

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
		var nestedApplicationCol []Application
		if o.Application != nil {
			nestedApplicationCol = []Application{}
			for _, oApplication := range o.Application {
				nestedApplication := Application{}
				if oApplication.Misc != nil {
					entry.Misc["Application"] = oApplication.Misc
				}
				if oApplication.Name != "" {
					nestedApplication.Name = oApplication.Name
				}
				if oApplication.Action != nil {
					nestedApplication.Action = oApplication.Action
				}
				nestedApplicationCol = append(nestedApplicationCol, nestedApplication)
			}
			entry.Application = nestedApplicationCol
		}

		var nestedDecoderCol []Decoder
		if o.Decoder != nil {
			nestedDecoderCol = []Decoder{}
			for _, oDecoder := range o.Decoder {
				nestedDecoder := Decoder{}
				if oDecoder.Misc != nil {
					entry.Misc["Decoder"] = oDecoder.Misc
				}
				if oDecoder.Name != "" {
					nestedDecoder.Name = oDecoder.Name
				}
				if oDecoder.Action != nil {
					nestedDecoder.Action = oDecoder.Action
				}
				if oDecoder.WildfireAction != nil {
					nestedDecoder.WildfireAction = oDecoder.WildfireAction
				}
				if oDecoder.MlavAction != nil {
					nestedDecoder.MlavAction = oDecoder.MlavAction
				}
				nestedDecoderCol = append(nestedDecoderCol, nestedDecoder)
			}
			entry.Decoder = nestedDecoderCol
		}

		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		var nestedMlavEngineFilebasedEnabledCol []MlavEngineFilebasedEnabled
		if o.MlavEngineFilebasedEnabled != nil {
			nestedMlavEngineFilebasedEnabledCol = []MlavEngineFilebasedEnabled{}
			for _, oMlavEngineFilebasedEnabled := range o.MlavEngineFilebasedEnabled {
				nestedMlavEngineFilebasedEnabled := MlavEngineFilebasedEnabled{}
				if oMlavEngineFilebasedEnabled.Misc != nil {
					entry.Misc["MlavEngineFilebasedEnabled"] = oMlavEngineFilebasedEnabled.Misc
				}
				if oMlavEngineFilebasedEnabled.Name != "" {
					nestedMlavEngineFilebasedEnabled.Name = oMlavEngineFilebasedEnabled.Name
				}
				if oMlavEngineFilebasedEnabled.MlavPolicyAction != nil {
					nestedMlavEngineFilebasedEnabled.MlavPolicyAction = oMlavEngineFilebasedEnabled.MlavPolicyAction
				}
				nestedMlavEngineFilebasedEnabledCol = append(nestedMlavEngineFilebasedEnabledCol, nestedMlavEngineFilebasedEnabled)
			}
			entry.MlavEngineFilebasedEnabled = nestedMlavEngineFilebasedEnabledCol
		}

		var nestedMlavExceptionCol []MlavException
		if o.MlavException != nil {
			nestedMlavExceptionCol = []MlavException{}
			for _, oMlavException := range o.MlavException {
				nestedMlavException := MlavException{}
				if oMlavException.Misc != nil {
					entry.Misc["MlavException"] = oMlavException.Misc
				}
				if oMlavException.Name != "" {
					nestedMlavException.Name = oMlavException.Name
				}
				if oMlavException.Filename != nil {
					nestedMlavException.Filename = oMlavException.Filename
				}
				if oMlavException.Description != nil {
					nestedMlavException.Description = oMlavException.Description
				}
				nestedMlavExceptionCol = append(nestedMlavExceptionCol, nestedMlavException)
			}
			entry.MlavException = nestedMlavExceptionCol
		}

		entry.PacketCapture = util.AsBool(o.PacketCapture, nil)
		var nestedThreatExceptionCol []ThreatException
		if o.ThreatException != nil {
			nestedThreatExceptionCol = []ThreatException{}
			for _, oThreatException := range o.ThreatException {
				nestedThreatException := ThreatException{}
				if oThreatException.Misc != nil {
					entry.Misc["ThreatException"] = oThreatException.Misc
				}
				if oThreatException.Name != "" {
					nestedThreatException.Name = oThreatException.Name
				}
				nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
			}
			entry.ThreatException = nestedThreatExceptionCol
		}

		entry.WfrtHoldMode = util.AsBool(o.WfrtHoldMode, nil)

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}
func (c *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedApplicationCol []Application
		if o.Application != nil {
			nestedApplicationCol = []Application{}
			for _, oApplication := range o.Application {
				nestedApplication := Application{}
				if oApplication.Misc != nil {
					entry.Misc["Application"] = oApplication.Misc
				}
				if oApplication.Name != "" {
					nestedApplication.Name = oApplication.Name
				}
				if oApplication.Action != nil {
					nestedApplication.Action = oApplication.Action
				}
				nestedApplicationCol = append(nestedApplicationCol, nestedApplication)
			}
			entry.Application = nestedApplicationCol
		}

		var nestedDecoderCol []Decoder
		if o.Decoder != nil {
			nestedDecoderCol = []Decoder{}
			for _, oDecoder := range o.Decoder {
				nestedDecoder := Decoder{}
				if oDecoder.Misc != nil {
					entry.Misc["Decoder"] = oDecoder.Misc
				}
				if oDecoder.Name != "" {
					nestedDecoder.Name = oDecoder.Name
				}
				if oDecoder.Action != nil {
					nestedDecoder.Action = oDecoder.Action
				}
				if oDecoder.WildfireAction != nil {
					nestedDecoder.WildfireAction = oDecoder.WildfireAction
				}
				if oDecoder.MlavAction != nil {
					nestedDecoder.MlavAction = oDecoder.MlavAction
				}
				nestedDecoderCol = append(nestedDecoderCol, nestedDecoder)
			}
			entry.Decoder = nestedDecoderCol
		}

		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		var nestedMlavEngineFilebasedEnabledCol []MlavEngineFilebasedEnabled
		if o.MlavEngineFilebasedEnabled != nil {
			nestedMlavEngineFilebasedEnabledCol = []MlavEngineFilebasedEnabled{}
			for _, oMlavEngineFilebasedEnabled := range o.MlavEngineFilebasedEnabled {
				nestedMlavEngineFilebasedEnabled := MlavEngineFilebasedEnabled{}
				if oMlavEngineFilebasedEnabled.Misc != nil {
					entry.Misc["MlavEngineFilebasedEnabled"] = oMlavEngineFilebasedEnabled.Misc
				}
				if oMlavEngineFilebasedEnabled.Name != "" {
					nestedMlavEngineFilebasedEnabled.Name = oMlavEngineFilebasedEnabled.Name
				}
				if oMlavEngineFilebasedEnabled.MlavPolicyAction != nil {
					nestedMlavEngineFilebasedEnabled.MlavPolicyAction = oMlavEngineFilebasedEnabled.MlavPolicyAction
				}
				nestedMlavEngineFilebasedEnabledCol = append(nestedMlavEngineFilebasedEnabledCol, nestedMlavEngineFilebasedEnabled)
			}
			entry.MlavEngineFilebasedEnabled = nestedMlavEngineFilebasedEnabledCol
		}

		var nestedMlavExceptionCol []MlavException
		if o.MlavException != nil {
			nestedMlavExceptionCol = []MlavException{}
			for _, oMlavException := range o.MlavException {
				nestedMlavException := MlavException{}
				if oMlavException.Misc != nil {
					entry.Misc["MlavException"] = oMlavException.Misc
				}
				if oMlavException.Name != "" {
					nestedMlavException.Name = oMlavException.Name
				}
				if oMlavException.Filename != nil {
					nestedMlavException.Filename = oMlavException.Filename
				}
				if oMlavException.Description != nil {
					nestedMlavException.Description = oMlavException.Description
				}
				nestedMlavExceptionCol = append(nestedMlavExceptionCol, nestedMlavException)
			}
			entry.MlavException = nestedMlavExceptionCol
		}

		entry.PacketCapture = util.AsBool(o.PacketCapture, nil)
		var nestedThreatExceptionCol []ThreatException
		if o.ThreatException != nil {
			nestedThreatExceptionCol = []ThreatException{}
			for _, oThreatException := range o.ThreatException {
				nestedThreatException := ThreatException{}
				if oThreatException.Misc != nil {
					entry.Misc["ThreatException"] = oThreatException.Misc
				}
				if oThreatException.Name != "" {
					nestedThreatException.Name = oThreatException.Name
				}
				nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
			}
			entry.ThreatException = nestedThreatExceptionCol
		}

		entry.WfrtHoldMode = util.AsBool(o.WfrtHoldMode, nil)

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
	if !matchApplication(a.Application, b.Application) {
		return false
	}
	if !matchDecoder(a.Decoder, b.Decoder) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !matchMlavEngineFilebasedEnabled(a.MlavEngineFilebasedEnabled, b.MlavEngineFilebasedEnabled) {
		return false
	}
	if !matchMlavException(a.MlavException, b.MlavException) {
		return false
	}
	if !util.BoolsMatch(a.PacketCapture, b.PacketCapture) {
		return false
	}
	if !matchThreatException(a.ThreatException, b.ThreatException) {
		return false
	}
	if !util.BoolsMatch(a.WfrtHoldMode, b.WfrtHoldMode) {
		return false
	}

	return true
}

func matchApplication(a []Application, b []Application) bool {
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
			if !util.StringsMatch(a.Action, b.Action) {
				return false
			}
		}
	}
	return true
}
func matchDecoder(a []Decoder, b []Decoder) bool {
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
			if !util.StringsMatch(a.Action, b.Action) {
				return false
			}
			if !util.StringsMatch(a.WildfireAction, b.WildfireAction) {
				return false
			}
			if !util.StringsMatch(a.MlavAction, b.MlavAction) {
				return false
			}
		}
	}
	return true
}
func matchMlavEngineFilebasedEnabled(a []MlavEngineFilebasedEnabled, b []MlavEngineFilebasedEnabled) bool {
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
			if !util.StringsMatch(a.MlavPolicyAction, b.MlavPolicyAction) {
				return false
			}
		}
	}
	return true
}
func matchMlavException(a []MlavException, b []MlavException) bool {
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
			if !util.StringsMatch(a.Filename, b.Filename) {
				return false
			}
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
		}
	}
	return true
}
func matchThreatException(a []ThreatException, b []ThreatException) bool {
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
