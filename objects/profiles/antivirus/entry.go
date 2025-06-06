package antivirus

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"profiles", "virus", "$name"}
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
	Misc                       []generic.Xml
}
type Application struct {
	Name   string
	Action *string
	Misc   []generic.Xml
}
type Decoder struct {
	Name           string
	Action         *string
	WildfireAction *string
	MlavAction     *string
	Misc           []generic.Xml
}
type MlavEngineFilebasedEnabled struct {
	Name             string
	MlavPolicyAction *string
	Misc             []generic.Xml
}
type MlavException struct {
	Name        string
	Filename    *string
	Description *string
	Misc        []generic.Xml
}
type ThreatException struct {
	Name string
	Misc []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
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
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName                    xml.Name                                `xml:"entry"`
	Name                       string                                  `xml:"name,attr"`
	Application                *applicationContainerXml                `xml:"application,omitempty"`
	Decoder                    *decoderContainerXml                    `xml:"decoder,omitempty"`
	Description                *string                                 `xml:"description,omitempty"`
	DisableOverride            *string                                 `xml:"disable-override,omitempty"`
	MlavEngineFilebasedEnabled *mlavEngineFilebasedEnabledContainerXml `xml:"mlav-engine-filebased-enabled,omitempty"`
	MlavException              *mlavExceptionContainerXml              `xml:"mlav-exception,omitempty"`
	PacketCapture              *string                                 `xml:"packet-capture,omitempty"`
	ThreatException            *threatExceptionContainerXml            `xml:"threat-exception,omitempty"`
	WfrtHoldMode               *string                                 `xml:"wfrt-hold-mode,omitempty"`
	Misc                       []generic.Xml                           `xml:",any"`
}
type applicationContainerXml struct {
	Entries []applicationXml `xml:"entry"`
}
type applicationXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Action  *string       `xml:"action,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type decoderContainerXml struct {
	Entries []decoderXml `xml:"entry"`
}
type decoderXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Action         *string       `xml:"action,omitempty"`
	WildfireAction *string       `xml:"wildfire-action,omitempty"`
	MlavAction     *string       `xml:"mlav-action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type mlavEngineFilebasedEnabledContainerXml struct {
	Entries []mlavEngineFilebasedEnabledXml `xml:"entry"`
}
type mlavEngineFilebasedEnabledXml struct {
	XMLName          xml.Name      `xml:"entry"`
	Name             string        `xml:"name,attr"`
	MlavPolicyAction *string       `xml:"mlav-policy-action,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type mlavExceptionContainerXml struct {
	Entries []mlavExceptionXml `xml:"entry"`
}
type mlavExceptionXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Filename    *string       `xml:"filename,omitempty"`
	Description *string       `xml:"description,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type threatExceptionContainerXml struct {
	Entries []threatExceptionXml `xml:"entry"`
}
type threatExceptionXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                    xml.Name                                       `xml:"entry"`
	Name                       string                                         `xml:"name,attr"`
	Application                *applicationContainerXml_11_0_2                `xml:"application,omitempty"`
	Decoder                    *decoderContainerXml_11_0_2                    `xml:"decoder,omitempty"`
	Description                *string                                        `xml:"description,omitempty"`
	DisableOverride            *string                                        `xml:"disable-override,omitempty"`
	MlavEngineFilebasedEnabled *mlavEngineFilebasedEnabledContainerXml_11_0_2 `xml:"mlav-engine-filebased-enabled,omitempty"`
	MlavException              *mlavExceptionContainerXml_11_0_2              `xml:"mlav-exception,omitempty"`
	PacketCapture              *string                                        `xml:"packet-capture,omitempty"`
	ThreatException            *threatExceptionContainerXml_11_0_2            `xml:"threat-exception,omitempty"`
	WfrtHoldMode               *string                                        `xml:"wfrt-hold-mode,omitempty"`
	Misc                       []generic.Xml                                  `xml:",any"`
}
type applicationContainerXml_11_0_2 struct {
	Entries []applicationXml_11_0_2 `xml:"entry"`
}
type applicationXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Action  *string       `xml:"action,omitempty"`
	Misc    []generic.Xml `xml:",any"`
}
type decoderContainerXml_11_0_2 struct {
	Entries []decoderXml_11_0_2 `xml:"entry"`
}
type decoderXml_11_0_2 struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Action         *string       `xml:"action,omitempty"`
	WildfireAction *string       `xml:"wildfire-action,omitempty"`
	MlavAction     *string       `xml:"mlav-action,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type mlavEngineFilebasedEnabledContainerXml_11_0_2 struct {
	Entries []mlavEngineFilebasedEnabledXml_11_0_2 `xml:"entry"`
}
type mlavEngineFilebasedEnabledXml_11_0_2 struct {
	XMLName          xml.Name      `xml:"entry"`
	Name             string        `xml:"name,attr"`
	MlavPolicyAction *string       `xml:"mlav-policy-action,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type mlavExceptionContainerXml_11_0_2 struct {
	Entries []mlavExceptionXml_11_0_2 `xml:"entry"`
}
type mlavExceptionXml_11_0_2 struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Filename    *string       `xml:"filename,omitempty"`
	Description *string       `xml:"description,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type threatExceptionContainerXml_11_0_2 struct {
	Entries []threatExceptionXml_11_0_2 `xml:"entry"`
}
type threatExceptionXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Application != nil {
		var objs []applicationXml
		for _, elt := range s.Application {
			var obj applicationXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Application = &applicationContainerXml{Entries: objs}
	}
	if s.Decoder != nil {
		var objs []decoderXml
		for _, elt := range s.Decoder {
			var obj decoderXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Decoder = &decoderContainerXml{Entries: objs}
	}
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.MlavEngineFilebasedEnabled != nil {
		var objs []mlavEngineFilebasedEnabledXml
		for _, elt := range s.MlavEngineFilebasedEnabled {
			var obj mlavEngineFilebasedEnabledXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MlavEngineFilebasedEnabled = &mlavEngineFilebasedEnabledContainerXml{Entries: objs}
	}
	if s.MlavException != nil {
		var objs []mlavExceptionXml
		for _, elt := range s.MlavException {
			var obj mlavExceptionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MlavException = &mlavExceptionContainerXml{Entries: objs}
	}
	o.PacketCapture = util.YesNo(s.PacketCapture, nil)
	if s.ThreatException != nil {
		var objs []threatExceptionXml
		for _, elt := range s.ThreatException {
			var obj threatExceptionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &threatExceptionContainerXml{Entries: objs}
	}
	o.WfrtHoldMode = util.YesNo(s.WfrtHoldMode, nil)
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var applicationVal []Application
	if o.Application != nil {
		for _, elt := range o.Application.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			applicationVal = append(applicationVal, *obj)
		}
	}
	var decoderVal []Decoder
	if o.Decoder != nil {
		for _, elt := range o.Decoder.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			decoderVal = append(decoderVal, *obj)
		}
	}
	var mlavEngineFilebasedEnabledVal []MlavEngineFilebasedEnabled
	if o.MlavEngineFilebasedEnabled != nil {
		for _, elt := range o.MlavEngineFilebasedEnabled.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			mlavEngineFilebasedEnabledVal = append(mlavEngineFilebasedEnabledVal, *obj)
		}
	}
	var mlavExceptionVal []MlavException
	if o.MlavException != nil {
		for _, elt := range o.MlavException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			mlavExceptionVal = append(mlavExceptionVal, *obj)
		}
	}
	var threatExceptionVal []ThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}

	result := &Entry{
		Name:                       o.Name,
		Application:                applicationVal,
		Decoder:                    decoderVal,
		Description:                o.Description,
		DisableOverride:            o.DisableOverride,
		MlavEngineFilebasedEnabled: mlavEngineFilebasedEnabledVal,
		MlavException:              mlavExceptionVal,
		PacketCapture:              util.AsBool(o.PacketCapture, nil),
		ThreatException:            threatExceptionVal,
		WfrtHoldMode:               util.AsBool(o.WfrtHoldMode, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *applicationXml) MarshalFromObject(s Application) {
	o.Name = s.Name
	o.Action = s.Action
	o.Misc = s.Misc
}

func (o applicationXml) UnmarshalToObject() (*Application, error) {

	result := &Application{
		Name:   o.Name,
		Action: o.Action,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *decoderXml) MarshalFromObject(s Decoder) {
	o.Name = s.Name
	o.Action = s.Action
	o.WildfireAction = s.WildfireAction
	o.MlavAction = s.MlavAction
	o.Misc = s.Misc
}

func (o decoderXml) UnmarshalToObject() (*Decoder, error) {

	result := &Decoder{
		Name:           o.Name,
		Action:         o.Action,
		WildfireAction: o.WildfireAction,
		MlavAction:     o.MlavAction,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *mlavEngineFilebasedEnabledXml) MarshalFromObject(s MlavEngineFilebasedEnabled) {
	o.Name = s.Name
	o.MlavPolicyAction = s.MlavPolicyAction
	o.Misc = s.Misc
}

func (o mlavEngineFilebasedEnabledXml) UnmarshalToObject() (*MlavEngineFilebasedEnabled, error) {

	result := &MlavEngineFilebasedEnabled{
		Name:             o.Name,
		MlavPolicyAction: o.MlavPolicyAction,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *mlavExceptionXml) MarshalFromObject(s MlavException) {
	o.Name = s.Name
	o.Filename = s.Filename
	o.Description = s.Description
	o.Misc = s.Misc
}

func (o mlavExceptionXml) UnmarshalToObject() (*MlavException, error) {

	result := &MlavException{
		Name:        o.Name,
		Filename:    o.Filename,
		Description: o.Description,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *threatExceptionXml) MarshalFromObject(s ThreatException) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o threatExceptionXml) UnmarshalToObject() (*ThreatException, error) {

	result := &ThreatException{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Application != nil {
		var objs []applicationXml_11_0_2
		for _, elt := range s.Application {
			var obj applicationXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Application = &applicationContainerXml_11_0_2{Entries: objs}
	}
	if s.Decoder != nil {
		var objs []decoderXml_11_0_2
		for _, elt := range s.Decoder {
			var obj decoderXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Decoder = &decoderContainerXml_11_0_2{Entries: objs}
	}
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.MlavEngineFilebasedEnabled != nil {
		var objs []mlavEngineFilebasedEnabledXml_11_0_2
		for _, elt := range s.MlavEngineFilebasedEnabled {
			var obj mlavEngineFilebasedEnabledXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MlavEngineFilebasedEnabled = &mlavEngineFilebasedEnabledContainerXml_11_0_2{Entries: objs}
	}
	if s.MlavException != nil {
		var objs []mlavExceptionXml_11_0_2
		for _, elt := range s.MlavException {
			var obj mlavExceptionXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MlavException = &mlavExceptionContainerXml_11_0_2{Entries: objs}
	}
	o.PacketCapture = util.YesNo(s.PacketCapture, nil)
	if s.ThreatException != nil {
		var objs []threatExceptionXml_11_0_2
		for _, elt := range s.ThreatException {
			var obj threatExceptionXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &threatExceptionContainerXml_11_0_2{Entries: objs}
	}
	o.WfrtHoldMode = util.YesNo(s.WfrtHoldMode, nil)
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var applicationVal []Application
	if o.Application != nil {
		for _, elt := range o.Application.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			applicationVal = append(applicationVal, *obj)
		}
	}
	var decoderVal []Decoder
	if o.Decoder != nil {
		for _, elt := range o.Decoder.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			decoderVal = append(decoderVal, *obj)
		}
	}
	var mlavEngineFilebasedEnabledVal []MlavEngineFilebasedEnabled
	if o.MlavEngineFilebasedEnabled != nil {
		for _, elt := range o.MlavEngineFilebasedEnabled.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			mlavEngineFilebasedEnabledVal = append(mlavEngineFilebasedEnabledVal, *obj)
		}
	}
	var mlavExceptionVal []MlavException
	if o.MlavException != nil {
		for _, elt := range o.MlavException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			mlavExceptionVal = append(mlavExceptionVal, *obj)
		}
	}
	var threatExceptionVal []ThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}

	result := &Entry{
		Name:                       o.Name,
		Application:                applicationVal,
		Decoder:                    decoderVal,
		Description:                o.Description,
		DisableOverride:            o.DisableOverride,
		MlavEngineFilebasedEnabled: mlavEngineFilebasedEnabledVal,
		MlavException:              mlavExceptionVal,
		PacketCapture:              util.AsBool(o.PacketCapture, nil),
		ThreatException:            threatExceptionVal,
		WfrtHoldMode:               util.AsBool(o.WfrtHoldMode, nil),
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *applicationXml_11_0_2) MarshalFromObject(s Application) {
	o.Name = s.Name
	o.Action = s.Action
	o.Misc = s.Misc
}

func (o applicationXml_11_0_2) UnmarshalToObject() (*Application, error) {

	result := &Application{
		Name:   o.Name,
		Action: o.Action,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *decoderXml_11_0_2) MarshalFromObject(s Decoder) {
	o.Name = s.Name
	o.Action = s.Action
	o.WildfireAction = s.WildfireAction
	o.MlavAction = s.MlavAction
	o.Misc = s.Misc
}

func (o decoderXml_11_0_2) UnmarshalToObject() (*Decoder, error) {

	result := &Decoder{
		Name:           o.Name,
		Action:         o.Action,
		WildfireAction: o.WildfireAction,
		MlavAction:     o.MlavAction,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *mlavEngineFilebasedEnabledXml_11_0_2) MarshalFromObject(s MlavEngineFilebasedEnabled) {
	o.Name = s.Name
	o.MlavPolicyAction = s.MlavPolicyAction
	o.Misc = s.Misc
}

func (o mlavEngineFilebasedEnabledXml_11_0_2) UnmarshalToObject() (*MlavEngineFilebasedEnabled, error) {

	result := &MlavEngineFilebasedEnabled{
		Name:             o.Name,
		MlavPolicyAction: o.MlavPolicyAction,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *mlavExceptionXml_11_0_2) MarshalFromObject(s MlavException) {
	o.Name = s.Name
	o.Filename = s.Filename
	o.Description = s.Description
	o.Misc = s.Misc
}

func (o mlavExceptionXml_11_0_2) UnmarshalToObject() (*MlavException, error) {

	result := &MlavException{
		Name:        o.Name,
		Filename:    o.Filename,
		Description: o.Description,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *threatExceptionXml_11_0_2) MarshalFromObject(s ThreatException) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o threatExceptionXml_11_0_2) UnmarshalToObject() (*ThreatException, error) {

	result := &ThreatException{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
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
	if len(o.Application) != len(other.Application) {
		return false
	}
	for idx := range o.Application {
		if !o.Application[idx].matches(&other.Application[idx]) {
			return false
		}
	}
	if len(o.Decoder) != len(other.Decoder) {
		return false
	}
	for idx := range o.Decoder {
		if !o.Decoder[idx].matches(&other.Decoder[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if len(o.MlavEngineFilebasedEnabled) != len(other.MlavEngineFilebasedEnabled) {
		return false
	}
	for idx := range o.MlavEngineFilebasedEnabled {
		if !o.MlavEngineFilebasedEnabled[idx].matches(&other.MlavEngineFilebasedEnabled[idx]) {
			return false
		}
	}
	if len(o.MlavException) != len(other.MlavException) {
		return false
	}
	for idx := range o.MlavException {
		if !o.MlavException[idx].matches(&other.MlavException[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}
	if len(o.ThreatException) != len(other.ThreatException) {
		return false
	}
	for idx := range o.ThreatException {
		if !o.ThreatException[idx].matches(&other.ThreatException[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.WfrtHoldMode, other.WfrtHoldMode) {
		return false
	}

	return true
}

func (o *Application) matches(other *Application) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}

	return true
}

func (o *Decoder) matches(other *Decoder) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.WildfireAction, other.WildfireAction) {
		return false
	}
	if !util.StringsMatch(o.MlavAction, other.MlavAction) {
		return false
	}

	return true
}

func (o *MlavEngineFilebasedEnabled) matches(other *MlavEngineFilebasedEnabled) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.MlavPolicyAction, other.MlavPolicyAction) {
		return false
	}

	return true
}

func (o *MlavException) matches(other *MlavException) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Filename, other.Filename) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}

	return true
}

func (o *ThreatException) matches(other *ThreatException) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
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
