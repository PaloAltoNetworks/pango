package decryption

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of a decryption
// rule.
//
// Targets is a map where the key is the serial number of the target device and
// the value is a list of specific vsys on that device.  The list of vsys is
// nil if all vsys on that device should be included or if the device is a
// virtual firewall (and thus only has vsys1).
type Entry struct {
	Name                       string
	Description                string
	SourceZones                []string // unordered
	SourceAddresses            []string // unordered
	NegateSource               bool
	SourceUsers                []string // unordered
	DestinationZones           []string // unordered
	DestinationAddresses       []string // unordered
	NegateDestination          bool
	Tags                       []string // ordered
	Disabled                   bool
	Services                   []string // unordered
	UrlCategories              []string
	Action                     string
	DecryptionType             string
	SslCertificate             string
	DecryptionProfile          string
	Targets                    map[string][]string
	NegateTarget               bool
	ForwardingProfile          string   // PAN-OS 8.1+
	Uuid                       string   // PAN-OS 9.0+
	GroupTag                   string   // PAN-OS 9.0+
	SourceHips                 []string // PAN-OS 10.0+
	DestinationHips            []string // PAN-OS 10.0+
	LogSuccessfulTlsHandshakes bool     // PAN-OS 10.0+
	LogFailedTlsHandshakes     bool     // PAN-OS 10.0+
	LogSetting                 string   // PAN-OS 10.0+
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name and Uuid fields relate to the identify of this object, they are not copied.
func (o *Entry) Copy(s Entry) {
	o.Description = s.Description
	if s.SourceZones == nil {
		o.SourceZones = nil
	} else {
		o.SourceZones = make([]string, len(s.SourceZones))
		copy(o.SourceZones, s.SourceZones)
	}
	if s.SourceAddresses == nil {
		o.SourceAddresses = nil
	} else {
		o.SourceAddresses = make([]string, len(s.SourceAddresses))
		copy(o.SourceAddresses, s.SourceAddresses)
	}
	o.NegateSource = s.NegateSource
	if s.SourceUsers == nil {
		o.SourceUsers = nil
	} else {
		o.SourceUsers = make([]string, len(s.SourceUsers))
		copy(o.SourceUsers, s.SourceUsers)
	}
	if s.DestinationZones == nil {
		o.DestinationZones = nil
	} else {
		o.DestinationZones = make([]string, len(s.DestinationZones))
		copy(o.DestinationZones, s.DestinationZones)
	}
	if s.DestinationAddresses == nil {
		o.DestinationAddresses = nil
	} else {
		o.DestinationAddresses = make([]string, len(s.DestinationAddresses))
		copy(o.DestinationAddresses, s.DestinationAddresses)
	}
	o.NegateDestination = s.NegateDestination
	if s.Tags == nil {
		o.Tags = nil
	} else {
		s.Tags = make([]string, len(s.Tags))
		copy(o.Tags, s.Tags)
	}
	o.Disabled = s.Disabled
	if s.Services == nil {
		o.Services = nil
	} else {
		o.Services = make([]string, len(s.Services))
		copy(o.Services, s.Services)
	}
	if s.UrlCategories == nil {
		o.UrlCategories = nil
	} else {
		o.UrlCategories = make([]string, len(s.UrlCategories))
		copy(o.UrlCategories, s.UrlCategories)
	}
	o.Action = s.Action
	o.DecryptionType = s.DecryptionType
	o.SslCertificate = s.SslCertificate
	o.DecryptionProfile = s.DecryptionProfile
	o.ForwardingProfile = s.ForwardingProfile
	o.GroupTag = s.GroupTag
	if s.SourceHips == nil {
		o.SourceHips = nil
	} else {
		o.SourceHips = make([]string, len(s.SourceHips))
		copy(o.SourceHips, s.SourceHips)
	}
	if s.DestinationHips == nil {
		o.DestinationHips = nil
	} else {
		o.DestinationHips = make([]string, len(s.DestinationHips))
		copy(o.DestinationHips, s.DestinationHips)
	}
	o.LogSuccessfulTlsHandshakes = s.LogSuccessfulTlsHandshakes
	o.LogFailedTlsHandshakes = s.LogFailedTlsHandshakes
	o.LogSetting = s.LogSetting
	if s.Targets == nil {
		o.Targets = nil
	} else {
		o.Targets = make(map[string][]string)
		for key, oval := range s.Targets {
			if oval == nil {
				o.Targets[key] = nil
			} else {
				val := make([]string, len(oval))
				copy(val, oval)
				o.Targets[key] = val
			}
		}
	}
	o.NegateTarget = s.NegateTarget
}

/** Structs / functions for normalization. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:                 o.Name,
		Description:          o.Description,
		SourceZones:          util.MemToStr(o.SourceZones),
		SourceAddresses:      util.MemToStr(o.SourceAddresses),
		NegateSource:         util.AsBool(o.NegateSource),
		SourceUsers:          util.MemToStr(o.SourceUsers),
		DestinationZones:     util.MemToStr(o.DestinationZones),
		DestinationAddresses: util.MemToStr(o.DestinationAddresses),
		NegateDestination:    util.AsBool(o.NegateDestination),
		Tags:                 util.MemToStr(o.Tags),
		Disabled:             util.AsBool(o.Disabled),
		Services:             util.MemToStr(o.Services),
		UrlCategories:        util.MemToStr(o.UrlCategories),
		Action:               o.Action,
		DecryptionProfile:    o.DecryptionProfile,
	}

	switch {
	case o.Type.SslForwardProxy != nil:
		ans.DecryptionType = DecryptionTypeSslForwardProxy
	case o.Type.SshProxy != nil:
		ans.DecryptionType = DecryptionTypeSshProxy
	case o.Type.SslCertificate != "":
		ans.DecryptionType = DecryptionTypeSslInboundInspection
		ans.SslCertificate = o.Type.SslCertificate
	}

	if o.TargetInfo != nil {
		ans.NegateTarget = util.AsBool(o.TargetInfo.NegateTarget)
		ans.Targets = util.VsysEntToMap(o.TargetInfo.Targets)
	}

	return ans
}

type entry_v1 struct {
	XMLName              xml.Name         `xml:"entry"`
	Name                 string           `xml:"name,attr"`
	Description          string           `xml:"description,omitempty"`
	SourceZones          *util.MemberType `xml:"from"`
	SourceAddresses      *util.MemberType `xml:"source"`
	NegateSource         string           `xml:"negate-source"`
	SourceUsers          *util.MemberType `xml:"source-user"`
	DestinationZones     *util.MemberType `xml:"to"`
	DestinationAddresses *util.MemberType `xml:"destination"`
	NegateDestination    string           `xml:"negate-destination"`
	Tags                 *util.MemberType `xml:"tag"`
	Disabled             string           `xml:"disabled"`
	Services             *util.MemberType `xml:"service"`
	UrlCategories        *util.MemberType `xml:"category"`
	Action               string           `xml:"action"`
	Type                 dType            `xml:"type"`
	DecryptionProfile    string           `xml:"profile,omitempty"`
	TargetInfo           *targetInfo      `xml:"target"`
}

type dType struct {
	SslForwardProxy *string `xml:"ssl-forward-proxy"`
	SshProxy        *string `xml:"ssh-proxy"`
	SslCertificate  string  `xml:"ssl-inbound-inspection,omitempty"`
}

type targetInfo struct {
	Targets      *util.VsysEntryType `xml:"devices"`
	NegateTarget string              `xml:"negate,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:                 e.Name,
		Description:          e.Description,
		SourceZones:          util.StrToMem(e.SourceZones),
		SourceAddresses:      util.StrToMem(e.SourceAddresses),
		NegateSource:         util.YesNo(e.NegateSource),
		SourceUsers:          util.StrToMem(e.SourceUsers),
		DestinationZones:     util.StrToMem(e.DestinationZones),
		DestinationAddresses: util.StrToMem(e.DestinationAddresses),
		NegateDestination:    util.YesNo(e.NegateDestination),
		Tags:                 util.StrToMem(e.Tags),
		Disabled:             util.YesNo(e.Disabled),
		Services:             util.StrToMem(e.Services),
		UrlCategories:        util.StrToMem(e.UrlCategories),
		Action:               e.Action,
		DecryptionProfile:    e.DecryptionProfile,
	}

	s := ""
	switch e.DecryptionType {
	case DecryptionTypeSslForwardProxy:
		ans.Type.SslForwardProxy = &s
	case DecryptionTypeSshProxy:
		ans.Type.SshProxy = &s
	case DecryptionTypeSslInboundInspection:
		ans.Type.SslCertificate = e.SslCertificate
	}

	if e.Targets != nil || e.NegateTarget {
		nfo := &targetInfo{
			Targets:      util.MapToVsysEnt(e.Targets),
			NegateTarget: util.YesNo(e.NegateTarget),
		}
		ans.TargetInfo = nfo
	}

	return ans
}

// PAN-OS 8.1
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o *container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:                 o.Name,
		Description:          o.Description,
		SourceZones:          util.MemToStr(o.SourceZones),
		SourceAddresses:      util.MemToStr(o.SourceAddresses),
		NegateSource:         util.AsBool(o.NegateSource),
		SourceUsers:          util.MemToStr(o.SourceUsers),
		DestinationZones:     util.MemToStr(o.DestinationZones),
		DestinationAddresses: util.MemToStr(o.DestinationAddresses),
		NegateDestination:    util.AsBool(o.NegateDestination),
		Tags:                 util.MemToStr(o.Tags),
		Disabled:             util.AsBool(o.Disabled),
		Services:             util.MemToStr(o.Services),
		UrlCategories:        util.MemToStr(o.UrlCategories),
		Action:               o.Action,
		DecryptionProfile:    o.DecryptionProfile,
		ForwardingProfile:    o.ForwardingProfile,
	}

	switch {
	case o.Type.SslForwardProxy != nil:
		ans.DecryptionType = DecryptionTypeSslForwardProxy
	case o.Type.SshProxy != nil:
		ans.DecryptionType = DecryptionTypeSshProxy
	case o.Type.SslCertificate != "":
		ans.DecryptionType = DecryptionTypeSslInboundInspection
		ans.SslCertificate = o.Type.SslCertificate
	}

	if o.TargetInfo != nil {
		ans.NegateTarget = util.AsBool(o.TargetInfo.NegateTarget)
		ans.Targets = util.VsysEntToMap(o.TargetInfo.Targets)
	}

	return ans
}

type entry_v2 struct {
	XMLName              xml.Name         `xml:"entry"`
	Name                 string           `xml:"name,attr"`
	Description          string           `xml:"description,omitempty"`
	SourceZones          *util.MemberType `xml:"from"`
	SourceAddresses      *util.MemberType `xml:"source"`
	NegateSource         string           `xml:"negate-source"`
	SourceUsers          *util.MemberType `xml:"source-user"`
	DestinationZones     *util.MemberType `xml:"to"`
	DestinationAddresses *util.MemberType `xml:"destination"`
	NegateDestination    string           `xml:"negate-destination"`
	Tags                 *util.MemberType `xml:"tag"`
	Disabled             string           `xml:"disabled"`
	Services             *util.MemberType `xml:"service"`
	UrlCategories        *util.MemberType `xml:"category"`
	Action               string           `xml:"action"`
	Type                 dType            `xml:"type"`
	DecryptionProfile    string           `xml:"profile,omitempty"`
	ForwardingProfile    string           `xml:"forwarding-profile,omitempty"`
	TargetInfo           *targetInfo      `xml:"target"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:                 e.Name,
		Description:          e.Description,
		SourceZones:          util.StrToMem(e.SourceZones),
		SourceAddresses:      util.StrToMem(e.SourceAddresses),
		NegateSource:         util.YesNo(e.NegateSource),
		SourceUsers:          util.StrToMem(e.SourceUsers),
		DestinationZones:     util.StrToMem(e.DestinationZones),
		DestinationAddresses: util.StrToMem(e.DestinationAddresses),
		NegateDestination:    util.YesNo(e.NegateDestination),
		Tags:                 util.StrToMem(e.Tags),
		Disabled:             util.YesNo(e.Disabled),
		Services:             util.StrToMem(e.Services),
		UrlCategories:        util.StrToMem(e.UrlCategories),
		Action:               e.Action,
		DecryptionProfile:    e.DecryptionProfile,
		ForwardingProfile:    e.ForwardingProfile,
	}

	s := ""
	switch e.DecryptionType {
	case DecryptionTypeSslForwardProxy:
		ans.Type.SslForwardProxy = &s
	case DecryptionTypeSshProxy:
		ans.Type.SshProxy = &s
	case DecryptionTypeSslInboundInspection:
		ans.Type.SslCertificate = e.SslCertificate
	}

	if e.Targets != nil || e.NegateTarget {
		nfo := &targetInfo{
			Targets:      util.MapToVsysEnt(e.Targets),
			NegateTarget: util.YesNo(e.NegateTarget),
		}
		ans.TargetInfo = nfo
	}

	return ans
}

// PAN-OS 9.0
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o *container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v3) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:                 o.Name,
		Uuid:                 o.Uuid,
		Description:          o.Description,
		SourceZones:          util.MemToStr(o.SourceZones),
		SourceAddresses:      util.MemToStr(o.SourceAddresses),
		NegateSource:         util.AsBool(o.NegateSource),
		SourceUsers:          util.MemToStr(o.SourceUsers),
		DestinationZones:     util.MemToStr(o.DestinationZones),
		DestinationAddresses: util.MemToStr(o.DestinationAddresses),
		NegateDestination:    util.AsBool(o.NegateDestination),
		Tags:                 util.MemToStr(o.Tags),
		Disabled:             util.AsBool(o.Disabled),
		Services:             util.MemToStr(o.Services),
		UrlCategories:        util.MemToStr(o.UrlCategories),
		Action:               o.Action,
		DecryptionProfile:    o.DecryptionProfile,
		ForwardingProfile:    o.ForwardingProfile,
		GroupTag:             o.GroupTag,
	}

	switch {
	case o.Type.SslForwardProxy != nil:
		ans.DecryptionType = DecryptionTypeSslForwardProxy
	case o.Type.SshProxy != nil:
		ans.DecryptionType = DecryptionTypeSshProxy
	case o.Type.SslCertificate != "":
		ans.DecryptionType = DecryptionTypeSslInboundInspection
		ans.SslCertificate = o.Type.SslCertificate
	}

	if o.TargetInfo != nil {
		ans.NegateTarget = util.AsBool(o.TargetInfo.NegateTarget)
		ans.Targets = util.VsysEntToMap(o.TargetInfo.Targets)
	}

	return ans
}

type entry_v3 struct {
	XMLName              xml.Name         `xml:"entry"`
	Name                 string           `xml:"name,attr"`
	Uuid                 string           `xml:"uuid,attr,omitempty"`
	Description          string           `xml:"description,omitempty"`
	SourceZones          *util.MemberType `xml:"from"`
	SourceAddresses      *util.MemberType `xml:"source"`
	NegateSource         string           `xml:"negate-source"`
	SourceUsers          *util.MemberType `xml:"source-user"`
	DestinationZones     *util.MemberType `xml:"to"`
	DestinationAddresses *util.MemberType `xml:"destination"`
	NegateDestination    string           `xml:"negate-destination"`
	Tags                 *util.MemberType `xml:"tag"`
	Disabled             string           `xml:"disabled"`
	Services             *util.MemberType `xml:"service"`
	UrlCategories        *util.MemberType `xml:"category"`
	Action               string           `xml:"action"`
	Type                 dType            `xml:"type"`
	DecryptionProfile    string           `xml:"profile,omitempty"`
	ForwardingProfile    string           `xml:"forwarding-profile,omitempty"`
	GroupTag             string           `xml:"group-tag,omitempty"`
	TargetInfo           *targetInfo      `xml:"target"`
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:                 e.Name,
		Uuid:                 e.Uuid,
		Description:          e.Description,
		SourceZones:          util.StrToMem(e.SourceZones),
		SourceAddresses:      util.StrToMem(e.SourceAddresses),
		NegateSource:         util.YesNo(e.NegateSource),
		SourceUsers:          util.StrToMem(e.SourceUsers),
		DestinationZones:     util.StrToMem(e.DestinationZones),
		DestinationAddresses: util.StrToMem(e.DestinationAddresses),
		NegateDestination:    util.YesNo(e.NegateDestination),
		Tags:                 util.StrToMem(e.Tags),
		Disabled:             util.YesNo(e.Disabled),
		Services:             util.StrToMem(e.Services),
		UrlCategories:        util.StrToMem(e.UrlCategories),
		Action:               e.Action,
		DecryptionProfile:    e.DecryptionProfile,
		ForwardingProfile:    e.ForwardingProfile,
		GroupTag:             e.GroupTag,
	}

	s := ""
	switch e.DecryptionType {
	case DecryptionTypeSslForwardProxy:
		ans.Type.SslForwardProxy = &s
	case DecryptionTypeSshProxy:
		ans.Type.SshProxy = &s
	case DecryptionTypeSslInboundInspection:
		ans.Type.SslCertificate = e.SslCertificate
	}

	if e.Targets != nil || e.NegateTarget {
		nfo := &targetInfo{
			Targets:      util.MapToVsysEnt(e.Targets),
			NegateTarget: util.YesNo(e.NegateTarget),
		}
		ans.TargetInfo = nfo
	}

	return ans
}

// PAN-OS 10.0
type container_v4 struct {
	Answer []entry_v4 `xml:"entry"`
}

func (o *container_v4) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v4) Normalize() []Entry {
	arr := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		arr = append(arr, o.Answer[i].normalize())
	}
	return arr
}

func (o *entry_v4) normalize() Entry {
	ans := Entry{
		Name:                       o.Name,
		Uuid:                       o.Uuid,
		Description:                o.Description,
		SourceZones:                util.MemToStr(o.SourceZones),
		SourceAddresses:            util.MemToStr(o.SourceAddresses),
		NegateSource:               util.AsBool(o.NegateSource),
		SourceUsers:                util.MemToStr(o.SourceUsers),
		DestinationZones:           util.MemToStr(o.DestinationZones),
		DestinationAddresses:       util.MemToStr(o.DestinationAddresses),
		NegateDestination:          util.AsBool(o.NegateDestination),
		Tags:                       util.MemToStr(o.Tags),
		Disabled:                   util.AsBool(o.Disabled),
		Services:                   util.MemToStr(o.Services),
		UrlCategories:              util.MemToStr(o.UrlCategories),
		Action:                     o.Action,
		DecryptionProfile:          o.DecryptionProfile,
		ForwardingProfile:          o.ForwardingProfile,
		GroupTag:                   o.GroupTag,
		SourceHips:                 util.MemToStr(o.SourceHips),
		DestinationHips:            util.MemToStr(o.DestinationHips),
		LogSuccessfulTlsHandshakes: util.AsBool(o.LogSuccessfulTlsHandshakes),
		LogFailedTlsHandshakes:     util.AsBool(o.LogFailedTlsHandshakes),
		LogSetting:                 o.LogSetting,
	}

	switch {
	case o.Type.SslForwardProxy != nil:
		ans.DecryptionType = DecryptionTypeSslForwardProxy
	case o.Type.SshProxy != nil:
		ans.DecryptionType = DecryptionTypeSshProxy
	case o.Type.SslCertificate != "":
		ans.DecryptionType = DecryptionTypeSslInboundInspection
		ans.SslCertificate = o.Type.SslCertificate
	}

	if o.TargetInfo != nil {
		ans.NegateTarget = util.AsBool(o.TargetInfo.NegateTarget)
		ans.Targets = util.VsysEntToMap(o.TargetInfo.Targets)
	}

	return ans
}

type entry_v4 struct {
	XMLName                    xml.Name         `xml:"entry"`
	Name                       string           `xml:"name,attr"`
	Uuid                       string           `xml:"uuid,attr,omitempty"`
	Description                string           `xml:"description,omitempty"`
	SourceZones                *util.MemberType `xml:"from"`
	SourceAddresses            *util.MemberType `xml:"source"`
	NegateSource               string           `xml:"negate-source"`
	SourceUsers                *util.MemberType `xml:"source-user"`
	DestinationZones           *util.MemberType `xml:"to"`
	DestinationAddresses       *util.MemberType `xml:"destination"`
	NegateDestination          string           `xml:"negate-destination"`
	Tags                       *util.MemberType `xml:"tag"`
	Disabled                   string           `xml:"disabled"`
	Services                   *util.MemberType `xml:"service"`
	UrlCategories              *util.MemberType `xml:"category"`
	Action                     string           `xml:"action"`
	Type                       dType            `xml:"type"`
	DecryptionProfile          string           `xml:"profile,omitempty"`
	ForwardingProfile          string           `xml:"forwarding-profile,omitempty"`
	GroupTag                   string           `xml:"group-tag,omitempty"`
	SourceHips                 *util.MemberType `xml:"source-hip"`
	DestinationHips            *util.MemberType `xml:"destination-hip"`
	LogSuccessfulTlsHandshakes string           `xml:"log-success"`
	LogFailedTlsHandshakes     string           `xml:"log-fail"`
	LogSetting                 string           `xml:"log-setting,omitempty"`
	TargetInfo                 *targetInfo      `xml:"target"`
}

func (e *entry_v4) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v4
	ans := local{
		LogFailedTlsHandshakes: util.YesNo(true),
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = entry_v4(ans)
	return nil
}

func specify_v4(e Entry) interface{} {
	ans := entry_v4{
		Name:                       e.Name,
		Uuid:                       e.Uuid,
		Description:                e.Description,
		SourceZones:                util.StrToMem(e.SourceZones),
		SourceAddresses:            util.StrToMem(e.SourceAddresses),
		NegateSource:               util.YesNo(e.NegateSource),
		SourceUsers:                util.StrToMem(e.SourceUsers),
		DestinationZones:           util.StrToMem(e.DestinationZones),
		DestinationAddresses:       util.StrToMem(e.DestinationAddresses),
		NegateDestination:          util.YesNo(e.NegateDestination),
		Tags:                       util.StrToMem(e.Tags),
		Disabled:                   util.YesNo(e.Disabled),
		Services:                   util.StrToMem(e.Services),
		UrlCategories:              util.StrToMem(e.UrlCategories),
		Action:                     e.Action,
		DecryptionProfile:          e.DecryptionProfile,
		ForwardingProfile:          e.ForwardingProfile,
		GroupTag:                   e.GroupTag,
		SourceHips:                 util.StrToMem(e.SourceHips),
		DestinationHips:            util.StrToMem(e.DestinationHips),
		LogSuccessfulTlsHandshakes: util.YesNo(e.LogSuccessfulTlsHandshakes),
		LogFailedTlsHandshakes:     util.YesNo(e.LogFailedTlsHandshakes),
		LogSetting:                 e.LogSetting,
	}

	s := ""
	switch e.DecryptionType {
	case DecryptionTypeSslForwardProxy:
		ans.Type.SslForwardProxy = &s
	case DecryptionTypeSshProxy:
		ans.Type.SshProxy = &s
	case DecryptionTypeSslInboundInspection:
		ans.Type.SslCertificate = e.SslCertificate
	}

	if e.Targets != nil || e.NegateTarget {
		nfo := &targetInfo{
			Targets:      util.MapToVsysEnt(e.Targets),
			NegateTarget: util.YesNo(e.NegateTarget),
		}
		ans.TargetInfo = nfo
	}

	return ans
}
