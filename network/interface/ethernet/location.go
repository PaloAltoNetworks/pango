package ethernet

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type ImportLocation interface {
	XpathForLocation(version.Number, util.ILocation) ([]string, error)
	MarshalPangoXML([]string) (string, error)
	UnmarshalPangoXML([]byte) ([]string, error)
}
type Layer3TemplateType int

const (
	layer3TemplateVirtualRouter Layer3TemplateType = iota
	layer3TemplateLogicalRouter Layer3TemplateType = iota
	layer3TemplateVsys          Layer3TemplateType = iota
	layer3TemplateZone          Layer3TemplateType = iota
)

type Layer3TemplateImportLocation struct {
	typ           Layer3TemplateType
	virtualRouter *Layer3TemplateVirtualRouterImportLocation
	logicalRouter *Layer3TemplateLogicalRouterImportLocation
	vsys          *Layer3TemplateVsysImportLocation
	zone          *Layer3TemplateZoneImportLocation
}

type Layer3TemplateVirtualRouterImportLocation struct {
	xpath  []string
	vsys   string
	router string
}

type Layer3TemplateVirtualRouterImportLocationSpec struct {
	Vsys   string
	Router string
}

func NewLayer3TemplateVirtualRouterImportLocation(spec Layer3TemplateVirtualRouterImportLocationSpec) *Layer3TemplateImportLocation {
	location := &Layer3TemplateVirtualRouterImportLocation{
		vsys:   spec.Vsys,
		router: spec.Router,
	}

	return &Layer3TemplateImportLocation{
		typ:           layer3TemplateVirtualRouter,
		virtualRouter: location,
	}
}

func (o *Layer3TemplateVirtualRouterImportLocation) XpathForLocation(vn version.Number, loc util.ILocation) ([]string, error) {
	ans, err := loc.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	importAns := []string{
		"network",
		"virtual-router",
		util.AsEntryXpath([]string{o.router}),
		"interface",
	}

	return append(ans, importAns...), nil
}

func (o *Layer3TemplateVirtualRouterImportLocation) MarshalPangoXML(interfaces []string) (string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type request struct {
		XMLName xml.Name `xml:"interface"`
		Members []member `xml:"member"`
	}

	var members []member
	for _, elt := range interfaces {
		members = append(members, member{Name: elt})
	}

	expected := request{
		Members: members,
	}
	bytes, err := xml.Marshal(expected)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (o *Layer3TemplateVirtualRouterImportLocation) UnmarshalPangoXML(bytes []byte) ([]string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type response struct {
		Members []member `xml:"result>interface>member"`
	}

	var existing response
	err := xml.Unmarshal(bytes, &existing)
	if err != nil {
		return nil, err
	}

	var interfaces []string
	for _, elt := range existing.Members {
		interfaces = append(interfaces, elt.Name)
	}

	return interfaces, nil
}

type Layer3TemplateLogicalRouterImportLocation struct {
	xpath  []string
	vsys   string
	router string
	vrf    string
}

type Layer3TemplateLogicalRouterImportLocationSpec struct {
	Vsys   string
	Router string
	Vrf    string
}

func NewLayer3TemplateLogicalRouterImportLocation(spec Layer3TemplateLogicalRouterImportLocationSpec) *Layer3TemplateImportLocation {
	location := &Layer3TemplateLogicalRouterImportLocation{
		vsys:   spec.Vsys,
		router: spec.Router,
		vrf:    spec.Vrf,
	}

	return &Layer3TemplateImportLocation{
		typ:           layer3TemplateLogicalRouter,
		logicalRouter: location,
	}
}

func (o *Layer3TemplateLogicalRouterImportLocation) XpathForLocation(vn version.Number, loc util.ILocation) ([]string, error) {
	ans, err := loc.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	importAns := []string{
		"network",
		"logical-router",
		util.AsEntryXpath([]string{o.router}),
		"vrf",
		util.AsEntryXpath([]string{o.vrf}),
		"interface",
	}

	return append(ans, importAns...), nil
}

func (o *Layer3TemplateLogicalRouterImportLocation) MarshalPangoXML(interfaces []string) (string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type request struct {
		XMLName xml.Name `xml:"interface"`
		Members []member `xml:"member"`
	}

	var members []member
	for _, elt := range interfaces {
		members = append(members, member{Name: elt})
	}

	expected := request{
		Members: members,
	}
	bytes, err := xml.Marshal(expected)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (o *Layer3TemplateLogicalRouterImportLocation) UnmarshalPangoXML(bytes []byte) ([]string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type response struct {
		Members []member `xml:"result>interface>member"`
	}

	var existing response
	err := xml.Unmarshal(bytes, &existing)
	if err != nil {
		return nil, err
	}

	var interfaces []string
	for _, elt := range existing.Members {
		interfaces = append(interfaces, elt.Name)
	}

	return interfaces, nil
}

type Layer3TemplateVsysImportLocation struct {
	xpath []string
	vsys  string
}

type Layer3TemplateVsysImportLocationSpec struct {
	Vsys string
}

func NewLayer3TemplateVsysImportLocation(spec Layer3TemplateVsysImportLocationSpec) *Layer3TemplateImportLocation {
	location := &Layer3TemplateVsysImportLocation{
		vsys: spec.Vsys,
	}

	return &Layer3TemplateImportLocation{
		typ:  layer3TemplateVsys,
		vsys: location,
	}
}

func (o *Layer3TemplateVsysImportLocation) XpathForLocation(vn version.Number, loc util.ILocation) ([]string, error) {
	ans, err := loc.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	importAns := []string{
		"vsys",
		util.AsEntryXpath([]string{o.vsys}),
		"import",
		"network",
		"interface",
	}

	return append(ans, importAns...), nil
}

func (o *Layer3TemplateVsysImportLocation) MarshalPangoXML(interfaces []string) (string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type request struct {
		XMLName xml.Name `xml:"interface"`
		Members []member `xml:"member"`
	}

	var members []member
	for _, elt := range interfaces {
		members = append(members, member{Name: elt})
	}

	expected := request{
		Members: members,
	}
	bytes, err := xml.Marshal(expected)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (o *Layer3TemplateVsysImportLocation) UnmarshalPangoXML(bytes []byte) ([]string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type response struct {
		Members []member `xml:"result>interface>member"`
	}

	var existing response
	err := xml.Unmarshal(bytes, &existing)
	if err != nil {
		return nil, err
	}

	var interfaces []string
	for _, elt := range existing.Members {
		interfaces = append(interfaces, elt.Name)
	}

	return interfaces, nil
}

type Layer3TemplateZoneImportLocation struct {
	xpath []string
	vsys  string
	zone  string
}

type Layer3TemplateZoneImportLocationSpec struct {
	Vsys string
	Zone string
}

func NewLayer3TemplateZoneImportLocation(spec Layer3TemplateZoneImportLocationSpec) *Layer3TemplateImportLocation {
	location := &Layer3TemplateZoneImportLocation{
		vsys: spec.Vsys,
		zone: spec.Zone,
	}

	return &Layer3TemplateImportLocation{
		typ:  layer3TemplateZone,
		zone: location,
	}
}

func (o *Layer3TemplateZoneImportLocation) XpathForLocation(vn version.Number, loc util.ILocation) ([]string, error) {
	ans, err := loc.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	importAns := []string{
		"vsys",
		util.AsEntryXpath([]string{o.vsys}),
		"zone",
		util.AsEntryXpath([]string{o.zone}),
		"network",
		"layer3",
	}

	return append(ans, importAns...), nil
}

func (o *Layer3TemplateZoneImportLocation) MarshalPangoXML(interfaces []string) (string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type request struct {
		XMLName xml.Name `xml:"layer3"`
		Members []member `xml:"member"`
	}

	var members []member
	for _, elt := range interfaces {
		members = append(members, member{Name: elt})
	}

	expected := request{
		Members: members,
	}
	bytes, err := xml.Marshal(expected)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (o *Layer3TemplateZoneImportLocation) UnmarshalPangoXML(bytes []byte) ([]string, error) {
	type member struct {
		Name string `xml:",chardata"`
	}

	type response struct {
		Members []member `xml:"result>layer3>member"`
	}

	var existing response
	err := xml.Unmarshal(bytes, &existing)
	if err != nil {
		return nil, err
	}

	var interfaces []string
	for _, elt := range existing.Members {
		interfaces = append(interfaces, elt.Name)
	}

	return interfaces, nil
}

func (o *Layer3TemplateImportLocation) MarshalPangoXML(interfaces []string) (string, error) {
	switch o.typ {
	case layer3TemplateVirtualRouter:
		return o.virtualRouter.MarshalPangoXML(interfaces)
	case layer3TemplateLogicalRouter:
		return o.logicalRouter.MarshalPangoXML(interfaces)
	case layer3TemplateVsys:
		return o.vsys.MarshalPangoXML(interfaces)
	case layer3TemplateZone:
		return o.zone.MarshalPangoXML(interfaces)
	default:
		return "", fmt.Errorf("invalid import location")
	}
}

func (o *Layer3TemplateImportLocation) UnmarshalPangoXML(bytes []byte) ([]string, error) {
	switch o.typ {
	case layer3TemplateVirtualRouter:
		return o.virtualRouter.UnmarshalPangoXML(bytes)
	case layer3TemplateLogicalRouter:
		return o.logicalRouter.UnmarshalPangoXML(bytes)
	case layer3TemplateVsys:
		return o.vsys.UnmarshalPangoXML(bytes)
	case layer3TemplateZone:
		return o.zone.UnmarshalPangoXML(bytes)
	default:
		return nil, fmt.Errorf("invalid import location")
	}
}

func (o *Layer3TemplateImportLocation) XpathForLocation(vn version.Number, loc util.ILocation) ([]string, error) {
	switch o.typ {
	case layer3TemplateVirtualRouter:
		return o.virtualRouter.XpathForLocation(vn, loc)
	case layer3TemplateLogicalRouter:
		return o.logicalRouter.XpathForLocation(vn, loc)
	case layer3TemplateVsys:
		return o.vsys.XpathForLocation(vn, loc)
	case layer3TemplateZone:
		return o.zone.XpathForLocation(vn, loc)
	default:
		return nil, fmt.Errorf("invalid import location")
	}
}

type Location struct {
	Shared        *SharedLocation        `json:"shared"`
	Template      *TemplateLocation      `json:"template,omitempty"`
	TemplateStack *TemplateStackLocation `json:"template_stack,omitempty"`
	Ngfw          *NgfwLocation          `json:"ngfw,omitempty"`
}

type SharedLocation struct {
}
type TemplateLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	Template       string `json:"template"`
}
type TemplateStackLocation struct {
	NgfwDevice     string `json:"ngfw_device"`
	PanoramaDevice string `json:"panorama_device"`
	TemplateStack  string `json:"template_stack"`
}
type NgfwLocation struct {
	NgfwDevice string `json:"ngfw_device"`
}

func NewSharedLocation() *Location {
	return &Location{Shared: &SharedLocation{},
	}
}
func NewTemplateLocation() *Location {
	return &Location{Template: &TemplateLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		Template:       "",
	},
	}
}
func NewTemplateStackLocation() *Location {
	return &Location{TemplateStack: &TemplateStackLocation{
		NgfwDevice:     "localhost.localdomain",
		PanoramaDevice: "localhost.localdomain",
		TemplateStack:  "",
	},
	}
}
func NewNgfwLocation() *Location {
	return &Location{Ngfw: &NgfwLocation{
		NgfwDevice: "localhost.localdomain",
	},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Shared != nil:
		count++
	case o.Template != nil:
		if o.Template.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Template.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.Template.Template == "" {
			return fmt.Errorf("Template is unspecified")
		}
		count++
	case o.TemplateStack != nil:
		if o.TemplateStack.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateStack.PanoramaDevice == "" {
			return fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStack.TemplateStack == "" {
			return fmt.Errorf("TemplateStack is unspecified")
		}
		count++
	case o.Ngfw != nil:
		if o.Ngfw.NgfwDevice == "" {
			return fmt.Errorf("NgfwDevice is unspecified")
		}
		count++
	}

	if count == 0 {
		return fmt.Errorf("no path specified")
	}

	if count > 1 {
		return fmt.Errorf("multiple paths specified: only one should be specified")
	}

	return nil
}

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.Shared != nil:
		ans = []string{
			"config",
			"shared",
		}
	case o.Template != nil:
		if o.Template.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.Template.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.Template.Template == "" {
			return nil, fmt.Errorf("Template is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Template.PanoramaDevice}),
			"template",
			util.AsEntryXpath([]string{o.Template.Template}),
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Template.NgfwDevice}),
		}
	case o.TemplateStack != nil:
		if o.TemplateStack.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		if o.TemplateStack.PanoramaDevice == "" {
			return nil, fmt.Errorf("PanoramaDevice is unspecified")
		}
		if o.TemplateStack.TemplateStack == "" {
			return nil, fmt.Errorf("TemplateStack is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateStack.PanoramaDevice}),
			"template-stack",
			util.AsEntryXpath([]string{o.TemplateStack.TemplateStack}),
			"config",
			"devices",
			util.AsEntryXpath([]string{o.TemplateStack.NgfwDevice}),
		}
	case o.Ngfw != nil:
		if o.Ngfw.NgfwDevice == "" {
			return nil, fmt.Errorf("NgfwDevice is unspecified")
		}
		ans = []string{
			"config",
			"devices",
			util.AsEntryXpath([]string{o.Ngfw.NgfwDevice}),
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}
func (o Location) XpathWithEntryName(vn version.Number, name string) ([]string, error) {

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}
	ans = append(ans, Suffix...)
	ans = append(ans, util.AsEntryXpath([]string{name}))

	return ans, nil
}
func (o Location) XpathWithUuid(vn version.Number, uuid string) ([]string, error) {

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}
	ans = append(ans, Suffix...)
	ans = append(ans, util.AsUuidXpath(uuid))

	return ans, nil
}
