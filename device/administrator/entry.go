package administrator

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
	suffix = []string{"users", "$name"}
)

type Entry struct {
	Name                  string
	AuthenticationProfile *string
	ClientCertificateOnly *bool
	PasswordProfile       *string
	Permissions           *Permissions
	Phash                 *string
	Preferences           *Preferences
	PublicKey             *string
	Misc                  []generic.Xml
	MiscAttributes        []xml.Attr
}
type Permissions struct {
	RoleBased      *PermissionsRoleBased
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PermissionsRoleBased struct {
	Custom         *PermissionsRoleBasedCustom
	Deviceadmin    []string
	Devicereader   []string
	PanoramaAdmin  *string
	Superreader    *string
	Superuser      *string
	Vsysadmin      []PermissionsRoleBasedVsysadmin
	Vsysreader     []PermissionsRoleBasedVsysreader
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PermissionsRoleBasedCustom struct {
	Profile        *string
	Vsys           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PermissionsRoleBasedVsysadmin struct {
	Name           string
	Vsys           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PermissionsRoleBasedVsysreader struct {
	Name           string
	Vsys           []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Preferences struct {
	DisableDns     *bool
	SavedLogQuery  *PreferencesSavedLogQuery
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQuery struct {
	Alarm          []PreferencesSavedLogQueryAlarm
	Auth           []PreferencesSavedLogQueryAuth
	Config         []PreferencesSavedLogQueryConfig
	Corr           []PreferencesSavedLogQueryCorr
	Data           []PreferencesSavedLogQueryData
	Decryption     []PreferencesSavedLogQueryDecryption
	Globalprotect  []PreferencesSavedLogQueryGlobalprotect
	Gtp            []PreferencesSavedLogQueryGtp
	Hipmatch       []PreferencesSavedLogQueryHipmatch
	System         []PreferencesSavedLogQuerySystem
	Threat         []PreferencesSavedLogQueryThreat
	Traffic        []PreferencesSavedLogQueryTraffic
	Tunnel         []PreferencesSavedLogQueryTunnel
	Unified        []PreferencesSavedLogQueryUnified
	Url            []PreferencesSavedLogQueryUrl
	Userid         []PreferencesSavedLogQueryUserid
	Wildfire       []PreferencesSavedLogQueryWildfire
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryAlarm struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryAuth struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryConfig struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryCorr struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryData struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryDecryption struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryGlobalprotect struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryGtp struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryHipmatch struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQuerySystem struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryThreat struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryTraffic struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryTunnel struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryUnified struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryUrl struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryUserid struct {
	Name           string
	Query          *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PreferencesSavedLogQueryWildfire struct {
	Name           string
	Query          *string
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
	XMLName               xml.Name        `xml:"entry"`
	Name                  string          `xml:"name,attr"`
	AuthenticationProfile *string         `xml:"authentication-profile,omitempty"`
	ClientCertificateOnly *string         `xml:"client-certificate-only,omitempty"`
	PasswordProfile       *string         `xml:"password-profile,omitempty"`
	Permissions           *permissionsXml `xml:"permissions,omitempty"`
	Phash                 *string         `xml:"phash,omitempty"`
	Preferences           *preferencesXml `xml:"preferences,omitempty"`
	PublicKey             *string         `xml:"public-key,omitempty"`
	Misc                  []generic.Xml   `xml:",any"`
	MiscAttributes        []xml.Attr      `xml:",any,attr"`
}
type permissionsXml struct {
	RoleBased      *permissionsRoleBasedXml `xml:"role-based,omitempty"`
	Misc           []generic.Xml            `xml:",any"`
	MiscAttributes []xml.Attr               `xml:",any,attr"`
}
type permissionsRoleBasedXml struct {
	Custom         *permissionsRoleBasedCustomXml              `xml:"custom,omitempty"`
	Deviceadmin    *util.MemberType                            `xml:"deviceadmin,omitempty"`
	Devicereader   *util.MemberType                            `xml:"devicereader,omitempty"`
	PanoramaAdmin  *string                                     `xml:"panorama-admin,omitempty"`
	Superreader    *string                                     `xml:"superreader,omitempty"`
	Superuser      *string                                     `xml:"superuser,omitempty"`
	Vsysadmin      *permissionsRoleBasedVsysadminContainerXml  `xml:"vsysadmin,omitempty"`
	Vsysreader     *permissionsRoleBasedVsysreaderContainerXml `xml:"vsysreader,omitempty"`
	Misc           []generic.Xml                               `xml:",any"`
	MiscAttributes []xml.Attr                                  `xml:",any,attr"`
}
type permissionsRoleBasedCustomXml struct {
	Profile        *string          `xml:"profile,omitempty"`
	Vsys           *util.MemberType `xml:"vsys,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type permissionsRoleBasedVsysadminContainerXml struct {
	Entries []permissionsRoleBasedVsysadminXml `xml:"entry"`
}
type permissionsRoleBasedVsysadminXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Vsys           *util.MemberType `xml:"vsys,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type permissionsRoleBasedVsysreaderContainerXml struct {
	Entries []permissionsRoleBasedVsysreaderXml `xml:"entry"`
}
type permissionsRoleBasedVsysreaderXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Vsys           *util.MemberType `xml:"vsys,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type preferencesXml struct {
	DisableDns     *string                      `xml:"disable-dns,omitempty"`
	SavedLogQuery  *preferencesSavedLogQueryXml `xml:"saved-log-query,omitempty"`
	Misc           []generic.Xml                `xml:",any"`
	MiscAttributes []xml.Attr                   `xml:",any,attr"`
}
type preferencesSavedLogQueryXml struct {
	Alarm          *preferencesSavedLogQueryAlarmContainerXml         `xml:"alarm,omitempty"`
	Auth           *preferencesSavedLogQueryAuthContainerXml          `xml:"auth,omitempty"`
	Config         *preferencesSavedLogQueryConfigContainerXml        `xml:"config,omitempty"`
	Corr           *preferencesSavedLogQueryCorrContainerXml          `xml:"corr,omitempty"`
	Data           *preferencesSavedLogQueryDataContainerXml          `xml:"data,omitempty"`
	Decryption     *preferencesSavedLogQueryDecryptionContainerXml    `xml:"decryption,omitempty"`
	Globalprotect  *preferencesSavedLogQueryGlobalprotectContainerXml `xml:"globalprotect,omitempty"`
	Gtp            *preferencesSavedLogQueryGtpContainerXml           `xml:"gtp,omitempty"`
	Hipmatch       *preferencesSavedLogQueryHipmatchContainerXml      `xml:"hipmatch,omitempty"`
	System         *preferencesSavedLogQuerySystemContainerXml        `xml:"system,omitempty"`
	Threat         *preferencesSavedLogQueryThreatContainerXml        `xml:"threat,omitempty"`
	Traffic        *preferencesSavedLogQueryTrafficContainerXml       `xml:"traffic,omitempty"`
	Tunnel         *preferencesSavedLogQueryTunnelContainerXml        `xml:"tunnel,omitempty"`
	Unified        *preferencesSavedLogQueryUnifiedContainerXml       `xml:"unified,omitempty"`
	Url            *preferencesSavedLogQueryUrlContainerXml           `xml:"url,omitempty"`
	Userid         *preferencesSavedLogQueryUseridContainerXml        `xml:"userid,omitempty"`
	Wildfire       *preferencesSavedLogQueryWildfireContainerXml      `xml:"wildfire,omitempty"`
	Misc           []generic.Xml                                      `xml:",any"`
	MiscAttributes []xml.Attr                                         `xml:",any,attr"`
}
type preferencesSavedLogQueryAlarmContainerXml struct {
	Entries []preferencesSavedLogQueryAlarmXml `xml:"entry"`
}
type preferencesSavedLogQueryAlarmXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryAuthContainerXml struct {
	Entries []preferencesSavedLogQueryAuthXml `xml:"entry"`
}
type preferencesSavedLogQueryAuthXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryConfigContainerXml struct {
	Entries []preferencesSavedLogQueryConfigXml `xml:"entry"`
}
type preferencesSavedLogQueryConfigXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryCorrContainerXml struct {
	Entries []preferencesSavedLogQueryCorrXml `xml:"entry"`
}
type preferencesSavedLogQueryCorrXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryDataContainerXml struct {
	Entries []preferencesSavedLogQueryDataXml `xml:"entry"`
}
type preferencesSavedLogQueryDataXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryDecryptionContainerXml struct {
	Entries []preferencesSavedLogQueryDecryptionXml `xml:"entry"`
}
type preferencesSavedLogQueryDecryptionXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryGlobalprotectContainerXml struct {
	Entries []preferencesSavedLogQueryGlobalprotectXml `xml:"entry"`
}
type preferencesSavedLogQueryGlobalprotectXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryGtpContainerXml struct {
	Entries []preferencesSavedLogQueryGtpXml `xml:"entry"`
}
type preferencesSavedLogQueryGtpXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryHipmatchContainerXml struct {
	Entries []preferencesSavedLogQueryHipmatchXml `xml:"entry"`
}
type preferencesSavedLogQueryHipmatchXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQuerySystemContainerXml struct {
	Entries []preferencesSavedLogQuerySystemXml `xml:"entry"`
}
type preferencesSavedLogQuerySystemXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryThreatContainerXml struct {
	Entries []preferencesSavedLogQueryThreatXml `xml:"entry"`
}
type preferencesSavedLogQueryThreatXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryTrafficContainerXml struct {
	Entries []preferencesSavedLogQueryTrafficXml `xml:"entry"`
}
type preferencesSavedLogQueryTrafficXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryTunnelContainerXml struct {
	Entries []preferencesSavedLogQueryTunnelXml `xml:"entry"`
}
type preferencesSavedLogQueryTunnelXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryUnifiedContainerXml struct {
	Entries []preferencesSavedLogQueryUnifiedXml `xml:"entry"`
}
type preferencesSavedLogQueryUnifiedXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryUrlContainerXml struct {
	Entries []preferencesSavedLogQueryUrlXml `xml:"entry"`
}
type preferencesSavedLogQueryUrlXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryUseridContainerXml struct {
	Entries []preferencesSavedLogQueryUseridXml `xml:"entry"`
}
type preferencesSavedLogQueryUseridXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type preferencesSavedLogQueryWildfireContainerXml struct {
	Entries []preferencesSavedLogQueryWildfireXml `xml:"entry"`
}
type preferencesSavedLogQueryWildfireXml struct {
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Query          *string       `xml:"query,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AuthenticationProfile = s.AuthenticationProfile
	o.ClientCertificateOnly = util.YesNo(s.ClientCertificateOnly, nil)
	o.PasswordProfile = s.PasswordProfile
	if s.Permissions != nil {
		var obj permissionsXml
		obj.MarshalFromObject(*s.Permissions)
		o.Permissions = &obj
	}
	o.Phash = s.Phash
	if s.Preferences != nil {
		var obj preferencesXml
		obj.MarshalFromObject(*s.Preferences)
		o.Preferences = &obj
	}
	o.PublicKey = s.PublicKey
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var permissionsVal *Permissions
	if o.Permissions != nil {
		obj, err := o.Permissions.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		permissionsVal = obj
	}
	var preferencesVal *Preferences
	if o.Preferences != nil {
		obj, err := o.Preferences.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		preferencesVal = obj
	}

	result := &Entry{
		Name:                  o.Name,
		AuthenticationProfile: o.AuthenticationProfile,
		ClientCertificateOnly: util.AsBool(o.ClientCertificateOnly, nil),
		PasswordProfile:       o.PasswordProfile,
		Permissions:           permissionsVal,
		Phash:                 o.Phash,
		Preferences:           preferencesVal,
		PublicKey:             o.PublicKey,
		Misc:                  o.Misc,
		MiscAttributes:        o.MiscAttributes,
	}
	return result, nil
}
func (o *permissionsXml) MarshalFromObject(s Permissions) {
	if s.RoleBased != nil {
		var obj permissionsRoleBasedXml
		obj.MarshalFromObject(*s.RoleBased)
		o.RoleBased = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o permissionsXml) UnmarshalToObject() (*Permissions, error) {
	var roleBasedVal *PermissionsRoleBased
	if o.RoleBased != nil {
		obj, err := o.RoleBased.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		roleBasedVal = obj
	}

	result := &Permissions{
		RoleBased:      roleBasedVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *permissionsRoleBasedXml) MarshalFromObject(s PermissionsRoleBased) {
	if s.Custom != nil {
		var obj permissionsRoleBasedCustomXml
		obj.MarshalFromObject(*s.Custom)
		o.Custom = &obj
	}
	if s.Deviceadmin != nil {
		o.Deviceadmin = util.StrToMem(s.Deviceadmin)
	}
	if s.Devicereader != nil {
		o.Devicereader = util.StrToMem(s.Devicereader)
	}
	o.PanoramaAdmin = s.PanoramaAdmin
	o.Superreader = s.Superreader
	o.Superuser = s.Superuser
	if s.Vsysadmin != nil {
		var objs []permissionsRoleBasedVsysadminXml
		for _, elt := range s.Vsysadmin {
			var obj permissionsRoleBasedVsysadminXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vsysadmin = &permissionsRoleBasedVsysadminContainerXml{Entries: objs}
	}
	if s.Vsysreader != nil {
		var objs []permissionsRoleBasedVsysreaderXml
		for _, elt := range s.Vsysreader {
			var obj permissionsRoleBasedVsysreaderXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Vsysreader = &permissionsRoleBasedVsysreaderContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o permissionsRoleBasedXml) UnmarshalToObject() (*PermissionsRoleBased, error) {
	var customVal *PermissionsRoleBasedCustom
	if o.Custom != nil {
		obj, err := o.Custom.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customVal = obj
	}
	var deviceadminVal []string
	if o.Deviceadmin != nil {
		deviceadminVal = util.MemToStr(o.Deviceadmin)
	}
	var devicereaderVal []string
	if o.Devicereader != nil {
		devicereaderVal = util.MemToStr(o.Devicereader)
	}
	var vsysadminVal []PermissionsRoleBasedVsysadmin
	if o.Vsysadmin != nil {
		for _, elt := range o.Vsysadmin.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vsysadminVal = append(vsysadminVal, *obj)
		}
	}
	var vsysreaderVal []PermissionsRoleBasedVsysreader
	if o.Vsysreader != nil {
		for _, elt := range o.Vsysreader.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			vsysreaderVal = append(vsysreaderVal, *obj)
		}
	}

	result := &PermissionsRoleBased{
		Custom:         customVal,
		Deviceadmin:    deviceadminVal,
		Devicereader:   devicereaderVal,
		PanoramaAdmin:  o.PanoramaAdmin,
		Superreader:    o.Superreader,
		Superuser:      o.Superuser,
		Vsysadmin:      vsysadminVal,
		Vsysreader:     vsysreaderVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *permissionsRoleBasedCustomXml) MarshalFromObject(s PermissionsRoleBasedCustom) {
	o.Profile = s.Profile
	if s.Vsys != nil {
		o.Vsys = util.StrToMem(s.Vsys)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o permissionsRoleBasedCustomXml) UnmarshalToObject() (*PermissionsRoleBasedCustom, error) {
	var vsysVal []string
	if o.Vsys != nil {
		vsysVal = util.MemToStr(o.Vsys)
	}

	result := &PermissionsRoleBasedCustom{
		Profile:        o.Profile,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *permissionsRoleBasedVsysadminXml) MarshalFromObject(s PermissionsRoleBasedVsysadmin) {
	o.Name = s.Name
	if s.Vsys != nil {
		o.Vsys = util.StrToMem(s.Vsys)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o permissionsRoleBasedVsysadminXml) UnmarshalToObject() (*PermissionsRoleBasedVsysadmin, error) {
	var vsysVal []string
	if o.Vsys != nil {
		vsysVal = util.MemToStr(o.Vsys)
	}

	result := &PermissionsRoleBasedVsysadmin{
		Name:           o.Name,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *permissionsRoleBasedVsysreaderXml) MarshalFromObject(s PermissionsRoleBasedVsysreader) {
	o.Name = s.Name
	if s.Vsys != nil {
		o.Vsys = util.StrToMem(s.Vsys)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o permissionsRoleBasedVsysreaderXml) UnmarshalToObject() (*PermissionsRoleBasedVsysreader, error) {
	var vsysVal []string
	if o.Vsys != nil {
		vsysVal = util.MemToStr(o.Vsys)
	}

	result := &PermissionsRoleBasedVsysreader{
		Name:           o.Name,
		Vsys:           vsysVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesXml) MarshalFromObject(s Preferences) {
	o.DisableDns = util.YesNo(s.DisableDns, nil)
	if s.SavedLogQuery != nil {
		var obj preferencesSavedLogQueryXml
		obj.MarshalFromObject(*s.SavedLogQuery)
		o.SavedLogQuery = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesXml) UnmarshalToObject() (*Preferences, error) {
	var savedLogQueryVal *PreferencesSavedLogQuery
	if o.SavedLogQuery != nil {
		obj, err := o.SavedLogQuery.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		savedLogQueryVal = obj
	}

	result := &Preferences{
		DisableDns:     util.AsBool(o.DisableDns, nil),
		SavedLogQuery:  savedLogQueryVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryXml) MarshalFromObject(s PreferencesSavedLogQuery) {
	if s.Alarm != nil {
		var objs []preferencesSavedLogQueryAlarmXml
		for _, elt := range s.Alarm {
			var obj preferencesSavedLogQueryAlarmXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Alarm = &preferencesSavedLogQueryAlarmContainerXml{Entries: objs}
	}
	if s.Auth != nil {
		var objs []preferencesSavedLogQueryAuthXml
		for _, elt := range s.Auth {
			var obj preferencesSavedLogQueryAuthXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Auth = &preferencesSavedLogQueryAuthContainerXml{Entries: objs}
	}
	if s.Config != nil {
		var objs []preferencesSavedLogQueryConfigXml
		for _, elt := range s.Config {
			var obj preferencesSavedLogQueryConfigXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Config = &preferencesSavedLogQueryConfigContainerXml{Entries: objs}
	}
	if s.Corr != nil {
		var objs []preferencesSavedLogQueryCorrXml
		for _, elt := range s.Corr {
			var obj preferencesSavedLogQueryCorrXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Corr = &preferencesSavedLogQueryCorrContainerXml{Entries: objs}
	}
	if s.Data != nil {
		var objs []preferencesSavedLogQueryDataXml
		for _, elt := range s.Data {
			var obj preferencesSavedLogQueryDataXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Data = &preferencesSavedLogQueryDataContainerXml{Entries: objs}
	}
	if s.Decryption != nil {
		var objs []preferencesSavedLogQueryDecryptionXml
		for _, elt := range s.Decryption {
			var obj preferencesSavedLogQueryDecryptionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Decryption = &preferencesSavedLogQueryDecryptionContainerXml{Entries: objs}
	}
	if s.Globalprotect != nil {
		var objs []preferencesSavedLogQueryGlobalprotectXml
		for _, elt := range s.Globalprotect {
			var obj preferencesSavedLogQueryGlobalprotectXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Globalprotect = &preferencesSavedLogQueryGlobalprotectContainerXml{Entries: objs}
	}
	if s.Gtp != nil {
		var objs []preferencesSavedLogQueryGtpXml
		for _, elt := range s.Gtp {
			var obj preferencesSavedLogQueryGtpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Gtp = &preferencesSavedLogQueryGtpContainerXml{Entries: objs}
	}
	if s.Hipmatch != nil {
		var objs []preferencesSavedLogQueryHipmatchXml
		for _, elt := range s.Hipmatch {
			var obj preferencesSavedLogQueryHipmatchXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Hipmatch = &preferencesSavedLogQueryHipmatchContainerXml{Entries: objs}
	}
	if s.System != nil {
		var objs []preferencesSavedLogQuerySystemXml
		for _, elt := range s.System {
			var obj preferencesSavedLogQuerySystemXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.System = &preferencesSavedLogQuerySystemContainerXml{Entries: objs}
	}
	if s.Threat != nil {
		var objs []preferencesSavedLogQueryThreatXml
		for _, elt := range s.Threat {
			var obj preferencesSavedLogQueryThreatXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Threat = &preferencesSavedLogQueryThreatContainerXml{Entries: objs}
	}
	if s.Traffic != nil {
		var objs []preferencesSavedLogQueryTrafficXml
		for _, elt := range s.Traffic {
			var obj preferencesSavedLogQueryTrafficXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Traffic = &preferencesSavedLogQueryTrafficContainerXml{Entries: objs}
	}
	if s.Tunnel != nil {
		var objs []preferencesSavedLogQueryTunnelXml
		for _, elt := range s.Tunnel {
			var obj preferencesSavedLogQueryTunnelXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Tunnel = &preferencesSavedLogQueryTunnelContainerXml{Entries: objs}
	}
	if s.Unified != nil {
		var objs []preferencesSavedLogQueryUnifiedXml
		for _, elt := range s.Unified {
			var obj preferencesSavedLogQueryUnifiedXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Unified = &preferencesSavedLogQueryUnifiedContainerXml{Entries: objs}
	}
	if s.Url != nil {
		var objs []preferencesSavedLogQueryUrlXml
		for _, elt := range s.Url {
			var obj preferencesSavedLogQueryUrlXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Url = &preferencesSavedLogQueryUrlContainerXml{Entries: objs}
	}
	if s.Userid != nil {
		var objs []preferencesSavedLogQueryUseridXml
		for _, elt := range s.Userid {
			var obj preferencesSavedLogQueryUseridXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Userid = &preferencesSavedLogQueryUseridContainerXml{Entries: objs}
	}
	if s.Wildfire != nil {
		var objs []preferencesSavedLogQueryWildfireXml
		for _, elt := range s.Wildfire {
			var obj preferencesSavedLogQueryWildfireXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Wildfire = &preferencesSavedLogQueryWildfireContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryXml) UnmarshalToObject() (*PreferencesSavedLogQuery, error) {
	var alarmVal []PreferencesSavedLogQueryAlarm
	if o.Alarm != nil {
		for _, elt := range o.Alarm.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			alarmVal = append(alarmVal, *obj)
		}
	}
	var authVal []PreferencesSavedLogQueryAuth
	if o.Auth != nil {
		for _, elt := range o.Auth.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			authVal = append(authVal, *obj)
		}
	}
	var configVal []PreferencesSavedLogQueryConfig
	if o.Config != nil {
		for _, elt := range o.Config.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			configVal = append(configVal, *obj)
		}
	}
	var corrVal []PreferencesSavedLogQueryCorr
	if o.Corr != nil {
		for _, elt := range o.Corr.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			corrVal = append(corrVal, *obj)
		}
	}
	var dataVal []PreferencesSavedLogQueryData
	if o.Data != nil {
		for _, elt := range o.Data.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			dataVal = append(dataVal, *obj)
		}
	}
	var decryptionVal []PreferencesSavedLogQueryDecryption
	if o.Decryption != nil {
		for _, elt := range o.Decryption.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			decryptionVal = append(decryptionVal, *obj)
		}
	}
	var globalprotectVal []PreferencesSavedLogQueryGlobalprotect
	if o.Globalprotect != nil {
		for _, elt := range o.Globalprotect.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			globalprotectVal = append(globalprotectVal, *obj)
		}
	}
	var gtpVal []PreferencesSavedLogQueryGtp
	if o.Gtp != nil {
		for _, elt := range o.Gtp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			gtpVal = append(gtpVal, *obj)
		}
	}
	var hipmatchVal []PreferencesSavedLogQueryHipmatch
	if o.Hipmatch != nil {
		for _, elt := range o.Hipmatch.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			hipmatchVal = append(hipmatchVal, *obj)
		}
	}
	var systemVal []PreferencesSavedLogQuerySystem
	if o.System != nil {
		for _, elt := range o.System.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			systemVal = append(systemVal, *obj)
		}
	}
	var threatVal []PreferencesSavedLogQueryThreat
	if o.Threat != nil {
		for _, elt := range o.Threat.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatVal = append(threatVal, *obj)
		}
	}
	var trafficVal []PreferencesSavedLogQueryTraffic
	if o.Traffic != nil {
		for _, elt := range o.Traffic.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			trafficVal = append(trafficVal, *obj)
		}
	}
	var tunnelVal []PreferencesSavedLogQueryTunnel
	if o.Tunnel != nil {
		for _, elt := range o.Tunnel.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			tunnelVal = append(tunnelVal, *obj)
		}
	}
	var unifiedVal []PreferencesSavedLogQueryUnified
	if o.Unified != nil {
		for _, elt := range o.Unified.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			unifiedVal = append(unifiedVal, *obj)
		}
	}
	var urlVal []PreferencesSavedLogQueryUrl
	if o.Url != nil {
		for _, elt := range o.Url.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			urlVal = append(urlVal, *obj)
		}
	}
	var useridVal []PreferencesSavedLogQueryUserid
	if o.Userid != nil {
		for _, elt := range o.Userid.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			useridVal = append(useridVal, *obj)
		}
	}
	var wildfireVal []PreferencesSavedLogQueryWildfire
	if o.Wildfire != nil {
		for _, elt := range o.Wildfire.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			wildfireVal = append(wildfireVal, *obj)
		}
	}

	result := &PreferencesSavedLogQuery{
		Alarm:          alarmVal,
		Auth:           authVal,
		Config:         configVal,
		Corr:           corrVal,
		Data:           dataVal,
		Decryption:     decryptionVal,
		Globalprotect:  globalprotectVal,
		Gtp:            gtpVal,
		Hipmatch:       hipmatchVal,
		System:         systemVal,
		Threat:         threatVal,
		Traffic:        trafficVal,
		Tunnel:         tunnelVal,
		Unified:        unifiedVal,
		Url:            urlVal,
		Userid:         useridVal,
		Wildfire:       wildfireVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryAlarmXml) MarshalFromObject(s PreferencesSavedLogQueryAlarm) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryAlarmXml) UnmarshalToObject() (*PreferencesSavedLogQueryAlarm, error) {

	result := &PreferencesSavedLogQueryAlarm{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryAuthXml) MarshalFromObject(s PreferencesSavedLogQueryAuth) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryAuthXml) UnmarshalToObject() (*PreferencesSavedLogQueryAuth, error) {

	result := &PreferencesSavedLogQueryAuth{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryConfigXml) MarshalFromObject(s PreferencesSavedLogQueryConfig) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryConfigXml) UnmarshalToObject() (*PreferencesSavedLogQueryConfig, error) {

	result := &PreferencesSavedLogQueryConfig{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryCorrXml) MarshalFromObject(s PreferencesSavedLogQueryCorr) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryCorrXml) UnmarshalToObject() (*PreferencesSavedLogQueryCorr, error) {

	result := &PreferencesSavedLogQueryCorr{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryDataXml) MarshalFromObject(s PreferencesSavedLogQueryData) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryDataXml) UnmarshalToObject() (*PreferencesSavedLogQueryData, error) {

	result := &PreferencesSavedLogQueryData{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryDecryptionXml) MarshalFromObject(s PreferencesSavedLogQueryDecryption) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryDecryptionXml) UnmarshalToObject() (*PreferencesSavedLogQueryDecryption, error) {

	result := &PreferencesSavedLogQueryDecryption{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryGlobalprotectXml) MarshalFromObject(s PreferencesSavedLogQueryGlobalprotect) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryGlobalprotectXml) UnmarshalToObject() (*PreferencesSavedLogQueryGlobalprotect, error) {

	result := &PreferencesSavedLogQueryGlobalprotect{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryGtpXml) MarshalFromObject(s PreferencesSavedLogQueryGtp) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryGtpXml) UnmarshalToObject() (*PreferencesSavedLogQueryGtp, error) {

	result := &PreferencesSavedLogQueryGtp{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryHipmatchXml) MarshalFromObject(s PreferencesSavedLogQueryHipmatch) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryHipmatchXml) UnmarshalToObject() (*PreferencesSavedLogQueryHipmatch, error) {

	result := &PreferencesSavedLogQueryHipmatch{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQuerySystemXml) MarshalFromObject(s PreferencesSavedLogQuerySystem) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQuerySystemXml) UnmarshalToObject() (*PreferencesSavedLogQuerySystem, error) {

	result := &PreferencesSavedLogQuerySystem{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryThreatXml) MarshalFromObject(s PreferencesSavedLogQueryThreat) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryThreatXml) UnmarshalToObject() (*PreferencesSavedLogQueryThreat, error) {

	result := &PreferencesSavedLogQueryThreat{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryTrafficXml) MarshalFromObject(s PreferencesSavedLogQueryTraffic) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryTrafficXml) UnmarshalToObject() (*PreferencesSavedLogQueryTraffic, error) {

	result := &PreferencesSavedLogQueryTraffic{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryTunnelXml) MarshalFromObject(s PreferencesSavedLogQueryTunnel) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryTunnelXml) UnmarshalToObject() (*PreferencesSavedLogQueryTunnel, error) {

	result := &PreferencesSavedLogQueryTunnel{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryUnifiedXml) MarshalFromObject(s PreferencesSavedLogQueryUnified) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryUnifiedXml) UnmarshalToObject() (*PreferencesSavedLogQueryUnified, error) {

	result := &PreferencesSavedLogQueryUnified{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryUrlXml) MarshalFromObject(s PreferencesSavedLogQueryUrl) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryUrlXml) UnmarshalToObject() (*PreferencesSavedLogQueryUrl, error) {

	result := &PreferencesSavedLogQueryUrl{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryUseridXml) MarshalFromObject(s PreferencesSavedLogQueryUserid) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryUseridXml) UnmarshalToObject() (*PreferencesSavedLogQueryUserid, error) {

	result := &PreferencesSavedLogQueryUserid{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *preferencesSavedLogQueryWildfireXml) MarshalFromObject(s PreferencesSavedLogQueryWildfire) {
	o.Name = s.Name
	o.Query = s.Query
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o preferencesSavedLogQueryWildfireXml) UnmarshalToObject() (*PreferencesSavedLogQueryWildfire, error) {

	result := &PreferencesSavedLogQueryWildfire{
		Name:           o.Name,
		Query:          o.Query,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authentication_profile" || v == "AuthenticationProfile" {
		return e.AuthenticationProfile, nil
	}
	if v == "client_certificate_only" || v == "ClientCertificateOnly" {
		return e.ClientCertificateOnly, nil
	}
	if v == "password_profile" || v == "PasswordProfile" {
		return e.PasswordProfile, nil
	}
	if v == "permissions" || v == "Permissions" {
		return e.Permissions, nil
	}
	if v == "phash" || v == "Phash" {
		return e.Phash, nil
	}
	if v == "preferences" || v == "Preferences" {
		return e.Preferences, nil
	}
	if v == "public_key" || v == "PublicKey" {
		return e.PublicKey, nil
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
	if !util.StringsMatch(o.AuthenticationProfile, other.AuthenticationProfile) {
		return false
	}
	if !util.BoolsMatch(o.ClientCertificateOnly, other.ClientCertificateOnly) {
		return false
	}
	if !util.StringsMatch(o.PasswordProfile, other.PasswordProfile) {
		return false
	}
	if !o.Permissions.matches(other.Permissions) {
		return false
	}
	if !util.StringsMatch(o.Phash, other.Phash) {
		return false
	}
	if !o.Preferences.matches(other.Preferences) {
		return false
	}
	if !util.StringsMatch(o.PublicKey, other.PublicKey) {
		return false
	}

	return true
}

func (o *Permissions) matches(other *Permissions) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.RoleBased.matches(other.RoleBased) {
		return false
	}

	return true
}

func (o *PermissionsRoleBased) matches(other *PermissionsRoleBased) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Custom.matches(other.Custom) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Deviceadmin, other.Deviceadmin) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Devicereader, other.Devicereader) {
		return false
	}
	if !util.StringsMatch(o.PanoramaAdmin, other.PanoramaAdmin) {
		return false
	}
	if !util.StringsMatch(o.Superreader, other.Superreader) {
		return false
	}
	if !util.StringsMatch(o.Superuser, other.Superuser) {
		return false
	}
	if len(o.Vsysadmin) != len(other.Vsysadmin) {
		return false
	}
	for idx := range o.Vsysadmin {
		if !o.Vsysadmin[idx].matches(&other.Vsysadmin[idx]) {
			return false
		}
	}
	if len(o.Vsysreader) != len(other.Vsysreader) {
		return false
	}
	for idx := range o.Vsysreader {
		if !o.Vsysreader[idx].matches(&other.Vsysreader[idx]) {
			return false
		}
	}

	return true
}

func (o *PermissionsRoleBasedCustom) matches(other *PermissionsRoleBasedCustom) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Profile, other.Profile) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vsys, other.Vsys) {
		return false
	}

	return true
}

func (o *PermissionsRoleBasedVsysadmin) matches(other *PermissionsRoleBasedVsysadmin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vsys, other.Vsys) {
		return false
	}

	return true
}

func (o *PermissionsRoleBasedVsysreader) matches(other *PermissionsRoleBasedVsysreader) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.OrderedListsMatch[string](o.Vsys, other.Vsys) {
		return false
	}

	return true
}

func (o *Preferences) matches(other *Preferences) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.DisableDns, other.DisableDns) {
		return false
	}
	if !o.SavedLogQuery.matches(other.SavedLogQuery) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQuery) matches(other *PreferencesSavedLogQuery) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.Alarm) != len(other.Alarm) {
		return false
	}
	for idx := range o.Alarm {
		if !o.Alarm[idx].matches(&other.Alarm[idx]) {
			return false
		}
	}
	if len(o.Auth) != len(other.Auth) {
		return false
	}
	for idx := range o.Auth {
		if !o.Auth[idx].matches(&other.Auth[idx]) {
			return false
		}
	}
	if len(o.Config) != len(other.Config) {
		return false
	}
	for idx := range o.Config {
		if !o.Config[idx].matches(&other.Config[idx]) {
			return false
		}
	}
	if len(o.Corr) != len(other.Corr) {
		return false
	}
	for idx := range o.Corr {
		if !o.Corr[idx].matches(&other.Corr[idx]) {
			return false
		}
	}
	if len(o.Data) != len(other.Data) {
		return false
	}
	for idx := range o.Data {
		if !o.Data[idx].matches(&other.Data[idx]) {
			return false
		}
	}
	if len(o.Decryption) != len(other.Decryption) {
		return false
	}
	for idx := range o.Decryption {
		if !o.Decryption[idx].matches(&other.Decryption[idx]) {
			return false
		}
	}
	if len(o.Globalprotect) != len(other.Globalprotect) {
		return false
	}
	for idx := range o.Globalprotect {
		if !o.Globalprotect[idx].matches(&other.Globalprotect[idx]) {
			return false
		}
	}
	if len(o.Gtp) != len(other.Gtp) {
		return false
	}
	for idx := range o.Gtp {
		if !o.Gtp[idx].matches(&other.Gtp[idx]) {
			return false
		}
	}
	if len(o.Hipmatch) != len(other.Hipmatch) {
		return false
	}
	for idx := range o.Hipmatch {
		if !o.Hipmatch[idx].matches(&other.Hipmatch[idx]) {
			return false
		}
	}
	if len(o.System) != len(other.System) {
		return false
	}
	for idx := range o.System {
		if !o.System[idx].matches(&other.System[idx]) {
			return false
		}
	}
	if len(o.Threat) != len(other.Threat) {
		return false
	}
	for idx := range o.Threat {
		if !o.Threat[idx].matches(&other.Threat[idx]) {
			return false
		}
	}
	if len(o.Traffic) != len(other.Traffic) {
		return false
	}
	for idx := range o.Traffic {
		if !o.Traffic[idx].matches(&other.Traffic[idx]) {
			return false
		}
	}
	if len(o.Tunnel) != len(other.Tunnel) {
		return false
	}
	for idx := range o.Tunnel {
		if !o.Tunnel[idx].matches(&other.Tunnel[idx]) {
			return false
		}
	}
	if len(o.Unified) != len(other.Unified) {
		return false
	}
	for idx := range o.Unified {
		if !o.Unified[idx].matches(&other.Unified[idx]) {
			return false
		}
	}
	if len(o.Url) != len(other.Url) {
		return false
	}
	for idx := range o.Url {
		if !o.Url[idx].matches(&other.Url[idx]) {
			return false
		}
	}
	if len(o.Userid) != len(other.Userid) {
		return false
	}
	for idx := range o.Userid {
		if !o.Userid[idx].matches(&other.Userid[idx]) {
			return false
		}
	}
	if len(o.Wildfire) != len(other.Wildfire) {
		return false
	}
	for idx := range o.Wildfire {
		if !o.Wildfire[idx].matches(&other.Wildfire[idx]) {
			return false
		}
	}

	return true
}

func (o *PreferencesSavedLogQueryAlarm) matches(other *PreferencesSavedLogQueryAlarm) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryAuth) matches(other *PreferencesSavedLogQueryAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryConfig) matches(other *PreferencesSavedLogQueryConfig) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryCorr) matches(other *PreferencesSavedLogQueryCorr) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryData) matches(other *PreferencesSavedLogQueryData) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryDecryption) matches(other *PreferencesSavedLogQueryDecryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryGlobalprotect) matches(other *PreferencesSavedLogQueryGlobalprotect) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryGtp) matches(other *PreferencesSavedLogQueryGtp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryHipmatch) matches(other *PreferencesSavedLogQueryHipmatch) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQuerySystem) matches(other *PreferencesSavedLogQuerySystem) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryThreat) matches(other *PreferencesSavedLogQueryThreat) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryTraffic) matches(other *PreferencesSavedLogQueryTraffic) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryTunnel) matches(other *PreferencesSavedLogQueryTunnel) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryUnified) matches(other *PreferencesSavedLogQueryUnified) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryUrl) matches(other *PreferencesSavedLogQueryUrl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryUserid) matches(other *PreferencesSavedLogQueryUserid) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
		return false
	}

	return true
}

func (o *PreferencesSavedLogQueryWildfire) matches(other *PreferencesSavedLogQueryWildfire) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Query, other.Query) {
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
