package authprofile

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
	suffix = []string{"authentication-profile", "$name"}
)

type Entry struct {
	Name             string
	AllowList        []string
	Lockout          *Lockout
	Method           *Method
	MultiFactorAuth  *MultiFactorAuth
	SingleSignOn     *SingleSignOn
	UserDomain       *string
	UsernameModifier *string
	Misc             []generic.Xml
	MiscAttributes   []xml.Attr
}
type Lockout struct {
	FailedAttempts *int64
	LockoutTime    *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Method struct {
	Cloud          *MethodCloud
	Kerberos       *MethodKerberos
	Ldap           *MethodLdap
	LocalDatabase  *MethodLocalDatabase
	None           *MethodNone
	Radius         *MethodRadius
	SamlIdp        *MethodSamlIdp
	Tacplus        *MethodTacplus
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodCloud struct {
	ClockSkew      *int64
	Region         *MethodCloudRegion
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodCloudRegion struct {
	RegionId       *string
	Tenant         *MethodCloudRegionTenant
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodCloudRegionTenant struct {
	Profile        *MethodCloudRegionTenantProfile
	TenantId       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodCloudRegionTenantProfile struct {
	Mfa            *MethodCloudRegionTenantProfileMfa
	ProfileId      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodCloudRegionTenantProfileMfa struct {
	ForceMfa       *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodKerberos struct {
	Realm          *string
	ServerProfile  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodLdap struct {
	LoginAttribute *string
	PasswdExpDays  *int64
	ServerProfile  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodLocalDatabase struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodNone struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodRadius struct {
	Checkgroup     *bool
	ServerProfile  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MethodSamlIdp struct {
	AttributeNameAccessDomain *string
	AttributeNameAdminRole    *string
	AttributeNameUsergroup    *string
	AttributeNameUsername     *string
	CertificateProfile        *string
	EnableSingleLogout        *bool
	RequestSigningCertificate *string
	ServerProfile             *string
	Misc                      []generic.Xml
	MiscAttributes            []xml.Attr
}
type MethodTacplus struct {
	Checkgroup     *bool
	ServerProfile  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type MultiFactorAuth struct {
	Factors        []string
	MfaEnable      *bool
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type SingleSignOn struct {
	KerberosKeytab   *string
	Realm            *string
	ServicePrincipal *string
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
	XMLName          xml.Name            `xml:"entry"`
	Name             string              `xml:"name,attr"`
	AllowList        *util.MemberType    `xml:"allow-list,omitempty"`
	Lockout          *lockoutXml         `xml:"lockout,omitempty"`
	Method           *methodXml          `xml:"method,omitempty"`
	MultiFactorAuth  *multiFactorAuthXml `xml:"multi-factor-auth,omitempty"`
	SingleSignOn     *singleSignOnXml    `xml:"single-sign-on,omitempty"`
	UserDomain       *string             `xml:"user-domain,omitempty"`
	UsernameModifier *string             `xml:"username-modifier,omitempty"`
	Misc             []generic.Xml       `xml:",any"`
	MiscAttributes   []xml.Attr          `xml:",any,attr"`
}
type lockoutXml struct {
	FailedAttempts *int64        `xml:"failed-attempts,omitempty"`
	LockoutTime    *int64        `xml:"lockout-time,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodXml struct {
	Cloud          *methodCloudXml         `xml:"cloud,omitempty"`
	Kerberos       *methodKerberosXml      `xml:"kerberos,omitempty"`
	Ldap           *methodLdapXml          `xml:"ldap,omitempty"`
	LocalDatabase  *methodLocalDatabaseXml `xml:"local-database,omitempty"`
	None           *methodNoneXml          `xml:"none,omitempty"`
	Radius         *methodRadiusXml        `xml:"radius,omitempty"`
	SamlIdp        *methodSamlIdpXml       `xml:"saml-idp,omitempty"`
	Tacplus        *methodTacplusXml       `xml:"tacplus,omitempty"`
	Misc           []generic.Xml           `xml:",any"`
	MiscAttributes []xml.Attr              `xml:",any,attr"`
}
type methodCloudXml struct {
	ClockSkew      *int64                `xml:"clock-skew,omitempty"`
	Region         *methodCloudRegionXml `xml:"region,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
	MiscAttributes []xml.Attr            `xml:",any,attr"`
}
type methodCloudRegionXml struct {
	RegionId       *string                     `xml:"region_id,omitempty"`
	Tenant         *methodCloudRegionTenantXml `xml:"tenant,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type methodCloudRegionTenantXml struct {
	Profile        *methodCloudRegionTenantProfileXml `xml:"profile,omitempty"`
	TenantId       *string                            `xml:"tenant_id,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type methodCloudRegionTenantProfileXml struct {
	Mfa            *methodCloudRegionTenantProfileMfaXml `xml:"mfa,omitempty"`
	ProfileId      *string                               `xml:"profile_id,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
	MiscAttributes []xml.Attr                            `xml:",any,attr"`
}
type methodCloudRegionTenantProfileMfaXml struct {
	ForceMfa       *string       `xml:"force-mfa,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodKerberosXml struct {
	Realm          *string       `xml:"realm,omitempty"`
	ServerProfile  *string       `xml:"server-profile,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodLdapXml struct {
	LoginAttribute *string       `xml:"login-attribute,omitempty"`
	PasswdExpDays  *int64        `xml:"passwd-exp-days,omitempty"`
	ServerProfile  *string       `xml:"server-profile,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodLocalDatabaseXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodNoneXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodRadiusXml struct {
	Checkgroup     *string       `xml:"checkgroup,omitempty"`
	ServerProfile  *string       `xml:"server-profile,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type methodSamlIdpXml struct {
	AttributeNameAccessDomain *string       `xml:"attribute-name-access-domain,omitempty"`
	AttributeNameAdminRole    *string       `xml:"attribute-name-admin-role,omitempty"`
	AttributeNameUsergroup    *string       `xml:"attribute-name-usergroup,omitempty"`
	AttributeNameUsername     *string       `xml:"attribute-name-username,omitempty"`
	CertificateProfile        *string       `xml:"certificate-profile,omitempty"`
	EnableSingleLogout        *string       `xml:"enable-single-logout,omitempty"`
	RequestSigningCertificate *string       `xml:"request-signing-certificate,omitempty"`
	ServerProfile             *string       `xml:"server-profile,omitempty"`
	Misc                      []generic.Xml `xml:",any"`
	MiscAttributes            []xml.Attr    `xml:",any,attr"`
}
type methodTacplusXml struct {
	Checkgroup     *string       `xml:"checkgroup,omitempty"`
	ServerProfile  *string       `xml:"server-profile,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type multiFactorAuthXml struct {
	Factors        *util.MemberType `xml:"factors,omitempty"`
	MfaEnable      *string          `xml:"mfa-enable,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type singleSignOnXml struct {
	KerberosKeytab   *string       `xml:"kerberos-keytab,omitempty"`
	Realm            *string       `xml:"realm,omitempty"`
	ServicePrincipal *string       `xml:"service-principal,omitempty"`
	Misc             []generic.Xml `xml:",any"`
	MiscAttributes   []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.AllowList != nil {
		o.AllowList = util.StrToMem(s.AllowList)
	}
	if s.Lockout != nil {
		var obj lockoutXml
		obj.MarshalFromObject(*s.Lockout)
		o.Lockout = &obj
	}
	if s.Method != nil {
		var obj methodXml
		obj.MarshalFromObject(*s.Method)
		o.Method = &obj
	}
	if s.MultiFactorAuth != nil {
		var obj multiFactorAuthXml
		obj.MarshalFromObject(*s.MultiFactorAuth)
		o.MultiFactorAuth = &obj
	}
	if s.SingleSignOn != nil {
		var obj singleSignOnXml
		obj.MarshalFromObject(*s.SingleSignOn)
		o.SingleSignOn = &obj
	}
	o.UserDomain = s.UserDomain
	o.UsernameModifier = s.UsernameModifier
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var allowListVal []string
	if o.AllowList != nil {
		allowListVal = util.MemToStr(o.AllowList)
	}
	var lockoutVal *Lockout
	if o.Lockout != nil {
		obj, err := o.Lockout.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		lockoutVal = obj
	}
	var methodVal *Method
	if o.Method != nil {
		obj, err := o.Method.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		methodVal = obj
	}
	var multiFactorAuthVal *MultiFactorAuth
	if o.MultiFactorAuth != nil {
		obj, err := o.MultiFactorAuth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multiFactorAuthVal = obj
	}
	var singleSignOnVal *SingleSignOn
	if o.SingleSignOn != nil {
		obj, err := o.SingleSignOn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		singleSignOnVal = obj
	}

	result := &Entry{
		Name:             o.Name,
		AllowList:        allowListVal,
		Lockout:          lockoutVal,
		Method:           methodVal,
		MultiFactorAuth:  multiFactorAuthVal,
		SingleSignOn:     singleSignOnVal,
		UserDomain:       o.UserDomain,
		UsernameModifier: o.UsernameModifier,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}
func (o *lockoutXml) MarshalFromObject(s Lockout) {
	o.FailedAttempts = s.FailedAttempts
	o.LockoutTime = s.LockoutTime
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o lockoutXml) UnmarshalToObject() (*Lockout, error) {

	result := &Lockout{
		FailedAttempts: o.FailedAttempts,
		LockoutTime:    o.LockoutTime,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodXml) MarshalFromObject(s Method) {
	if s.Cloud != nil {
		var obj methodCloudXml
		obj.MarshalFromObject(*s.Cloud)
		o.Cloud = &obj
	}
	if s.Kerberos != nil {
		var obj methodKerberosXml
		obj.MarshalFromObject(*s.Kerberos)
		o.Kerberos = &obj
	}
	if s.Ldap != nil {
		var obj methodLdapXml
		obj.MarshalFromObject(*s.Ldap)
		o.Ldap = &obj
	}
	if s.LocalDatabase != nil {
		var obj methodLocalDatabaseXml
		obj.MarshalFromObject(*s.LocalDatabase)
		o.LocalDatabase = &obj
	}
	if s.None != nil {
		var obj methodNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Radius != nil {
		var obj methodRadiusXml
		obj.MarshalFromObject(*s.Radius)
		o.Radius = &obj
	}
	if s.SamlIdp != nil {
		var obj methodSamlIdpXml
		obj.MarshalFromObject(*s.SamlIdp)
		o.SamlIdp = &obj
	}
	if s.Tacplus != nil {
		var obj methodTacplusXml
		obj.MarshalFromObject(*s.Tacplus)
		o.Tacplus = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodXml) UnmarshalToObject() (*Method, error) {
	var cloudVal *MethodCloud
	if o.Cloud != nil {
		obj, err := o.Cloud.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cloudVal = obj
	}
	var kerberosVal *MethodKerberos
	if o.Kerberos != nil {
		obj, err := o.Kerberos.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		kerberosVal = obj
	}
	var ldapVal *MethodLdap
	if o.Ldap != nil {
		obj, err := o.Ldap.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ldapVal = obj
	}
	var localDatabaseVal *MethodLocalDatabase
	if o.LocalDatabase != nil {
		obj, err := o.LocalDatabase.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localDatabaseVal = obj
	}
	var noneVal *MethodNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var radiusVal *MethodRadius
	if o.Radius != nil {
		obj, err := o.Radius.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		radiusVal = obj
	}
	var samlIdpVal *MethodSamlIdp
	if o.SamlIdp != nil {
		obj, err := o.SamlIdp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		samlIdpVal = obj
	}
	var tacplusVal *MethodTacplus
	if o.Tacplus != nil {
		obj, err := o.Tacplus.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tacplusVal = obj
	}

	result := &Method{
		Cloud:          cloudVal,
		Kerberos:       kerberosVal,
		Ldap:           ldapVal,
		LocalDatabase:  localDatabaseVal,
		None:           noneVal,
		Radius:         radiusVal,
		SamlIdp:        samlIdpVal,
		Tacplus:        tacplusVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodCloudXml) MarshalFromObject(s MethodCloud) {
	o.ClockSkew = s.ClockSkew
	if s.Region != nil {
		var obj methodCloudRegionXml
		obj.MarshalFromObject(*s.Region)
		o.Region = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodCloudXml) UnmarshalToObject() (*MethodCloud, error) {
	var regionVal *MethodCloudRegion
	if o.Region != nil {
		obj, err := o.Region.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		regionVal = obj
	}

	result := &MethodCloud{
		ClockSkew:      o.ClockSkew,
		Region:         regionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodCloudRegionXml) MarshalFromObject(s MethodCloudRegion) {
	o.RegionId = s.RegionId
	if s.Tenant != nil {
		var obj methodCloudRegionTenantXml
		obj.MarshalFromObject(*s.Tenant)
		o.Tenant = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodCloudRegionXml) UnmarshalToObject() (*MethodCloudRegion, error) {
	var tenantVal *MethodCloudRegionTenant
	if o.Tenant != nil {
		obj, err := o.Tenant.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tenantVal = obj
	}

	result := &MethodCloudRegion{
		RegionId:       o.RegionId,
		Tenant:         tenantVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodCloudRegionTenantXml) MarshalFromObject(s MethodCloudRegionTenant) {
	if s.Profile != nil {
		var obj methodCloudRegionTenantProfileXml
		obj.MarshalFromObject(*s.Profile)
		o.Profile = &obj
	}
	o.TenantId = s.TenantId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodCloudRegionTenantXml) UnmarshalToObject() (*MethodCloudRegionTenant, error) {
	var profileVal *MethodCloudRegionTenantProfile
	if o.Profile != nil {
		obj, err := o.Profile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		profileVal = obj
	}

	result := &MethodCloudRegionTenant{
		Profile:        profileVal,
		TenantId:       o.TenantId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodCloudRegionTenantProfileXml) MarshalFromObject(s MethodCloudRegionTenantProfile) {
	if s.Mfa != nil {
		var obj methodCloudRegionTenantProfileMfaXml
		obj.MarshalFromObject(*s.Mfa)
		o.Mfa = &obj
	}
	o.ProfileId = s.ProfileId
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodCloudRegionTenantProfileXml) UnmarshalToObject() (*MethodCloudRegionTenantProfile, error) {
	var mfaVal *MethodCloudRegionTenantProfileMfa
	if o.Mfa != nil {
		obj, err := o.Mfa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		mfaVal = obj
	}

	result := &MethodCloudRegionTenantProfile{
		Mfa:            mfaVal,
		ProfileId:      o.ProfileId,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodCloudRegionTenantProfileMfaXml) MarshalFromObject(s MethodCloudRegionTenantProfileMfa) {
	o.ForceMfa = s.ForceMfa
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodCloudRegionTenantProfileMfaXml) UnmarshalToObject() (*MethodCloudRegionTenantProfileMfa, error) {

	result := &MethodCloudRegionTenantProfileMfa{
		ForceMfa:       o.ForceMfa,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodKerberosXml) MarshalFromObject(s MethodKerberos) {
	o.Realm = s.Realm
	o.ServerProfile = s.ServerProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodKerberosXml) UnmarshalToObject() (*MethodKerberos, error) {

	result := &MethodKerberos{
		Realm:          o.Realm,
		ServerProfile:  o.ServerProfile,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodLdapXml) MarshalFromObject(s MethodLdap) {
	o.LoginAttribute = s.LoginAttribute
	o.PasswdExpDays = s.PasswdExpDays
	o.ServerProfile = s.ServerProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodLdapXml) UnmarshalToObject() (*MethodLdap, error) {

	result := &MethodLdap{
		LoginAttribute: o.LoginAttribute,
		PasswdExpDays:  o.PasswdExpDays,
		ServerProfile:  o.ServerProfile,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodLocalDatabaseXml) MarshalFromObject(s MethodLocalDatabase) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodLocalDatabaseXml) UnmarshalToObject() (*MethodLocalDatabase, error) {

	result := &MethodLocalDatabase{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodNoneXml) MarshalFromObject(s MethodNone) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodNoneXml) UnmarshalToObject() (*MethodNone, error) {

	result := &MethodNone{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodRadiusXml) MarshalFromObject(s MethodRadius) {
	o.Checkgroup = util.YesNo(s.Checkgroup, nil)
	o.ServerProfile = s.ServerProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodRadiusXml) UnmarshalToObject() (*MethodRadius, error) {

	result := &MethodRadius{
		Checkgroup:     util.AsBool(o.Checkgroup, nil),
		ServerProfile:  o.ServerProfile,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *methodSamlIdpXml) MarshalFromObject(s MethodSamlIdp) {
	o.AttributeNameAccessDomain = s.AttributeNameAccessDomain
	o.AttributeNameAdminRole = s.AttributeNameAdminRole
	o.AttributeNameUsergroup = s.AttributeNameUsergroup
	o.AttributeNameUsername = s.AttributeNameUsername
	o.CertificateProfile = s.CertificateProfile
	o.EnableSingleLogout = util.YesNo(s.EnableSingleLogout, nil)
	o.RequestSigningCertificate = s.RequestSigningCertificate
	o.ServerProfile = s.ServerProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodSamlIdpXml) UnmarshalToObject() (*MethodSamlIdp, error) {

	result := &MethodSamlIdp{
		AttributeNameAccessDomain: o.AttributeNameAccessDomain,
		AttributeNameAdminRole:    o.AttributeNameAdminRole,
		AttributeNameUsergroup:    o.AttributeNameUsergroup,
		AttributeNameUsername:     o.AttributeNameUsername,
		CertificateProfile:        o.CertificateProfile,
		EnableSingleLogout:        util.AsBool(o.EnableSingleLogout, nil),
		RequestSigningCertificate: o.RequestSigningCertificate,
		ServerProfile:             o.ServerProfile,
		Misc:                      o.Misc,
		MiscAttributes:            o.MiscAttributes,
	}
	return result, nil
}
func (o *methodTacplusXml) MarshalFromObject(s MethodTacplus) {
	o.Checkgroup = util.YesNo(s.Checkgroup, nil)
	o.ServerProfile = s.ServerProfile
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o methodTacplusXml) UnmarshalToObject() (*MethodTacplus, error) {

	result := &MethodTacplus{
		Checkgroup:     util.AsBool(o.Checkgroup, nil),
		ServerProfile:  o.ServerProfile,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *multiFactorAuthXml) MarshalFromObject(s MultiFactorAuth) {
	if s.Factors != nil {
		o.Factors = util.StrToMem(s.Factors)
	}
	o.MfaEnable = util.YesNo(s.MfaEnable, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o multiFactorAuthXml) UnmarshalToObject() (*MultiFactorAuth, error) {
	var factorsVal []string
	if o.Factors != nil {
		factorsVal = util.MemToStr(o.Factors)
	}

	result := &MultiFactorAuth{
		Factors:        factorsVal,
		MfaEnable:      util.AsBool(o.MfaEnable, nil),
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *singleSignOnXml) MarshalFromObject(s SingleSignOn) {
	o.KerberosKeytab = s.KerberosKeytab
	o.Realm = s.Realm
	o.ServicePrincipal = s.ServicePrincipal
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o singleSignOnXml) UnmarshalToObject() (*SingleSignOn, error) {

	result := &SingleSignOn{
		KerberosKeytab:   o.KerberosKeytab,
		Realm:            o.Realm,
		ServicePrincipal: o.ServicePrincipal,
		Misc:             o.Misc,
		MiscAttributes:   o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "allow_list" || v == "AllowList" {
		return e.AllowList, nil
	}
	if v == "allow_list|LENGTH" || v == "AllowList|LENGTH" {
		return int64(len(e.AllowList)), nil
	}
	if v == "lockout" || v == "Lockout" {
		return e.Lockout, nil
	}
	if v == "method" || v == "Method" {
		return e.Method, nil
	}
	if v == "multi_factor_auth" || v == "MultiFactorAuth" {
		return e.MultiFactorAuth, nil
	}
	if v == "single_sign_on" || v == "SingleSignOn" {
		return e.SingleSignOn, nil
	}
	if v == "user_domain" || v == "UserDomain" {
		return e.UserDomain, nil
	}
	if v == "username_modifier" || v == "UsernameModifier" {
		return e.UsernameModifier, nil
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
	if !util.OrderedListsMatch[string](o.AllowList, other.AllowList) {
		return false
	}
	if !o.Lockout.matches(other.Lockout) {
		return false
	}
	if !o.Method.matches(other.Method) {
		return false
	}
	if !o.MultiFactorAuth.matches(other.MultiFactorAuth) {
		return false
	}
	if !o.SingleSignOn.matches(other.SingleSignOn) {
		return false
	}
	if !util.StringsMatch(o.UserDomain, other.UserDomain) {
		return false
	}
	if !util.StringsMatch(o.UsernameModifier, other.UsernameModifier) {
		return false
	}

	return true
}

func (o *Lockout) matches(other *Lockout) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.FailedAttempts, other.FailedAttempts) {
		return false
	}
	if !util.Ints64Match(o.LockoutTime, other.LockoutTime) {
		return false
	}

	return true
}

func (o *Method) matches(other *Method) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Cloud.matches(other.Cloud) {
		return false
	}
	if !o.Kerberos.matches(other.Kerberos) {
		return false
	}
	if !o.Ldap.matches(other.Ldap) {
		return false
	}
	if !o.LocalDatabase.matches(other.LocalDatabase) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Radius.matches(other.Radius) {
		return false
	}
	if !o.SamlIdp.matches(other.SamlIdp) {
		return false
	}
	if !o.Tacplus.matches(other.Tacplus) {
		return false
	}

	return true
}

func (o *MethodCloud) matches(other *MethodCloud) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ClockSkew, other.ClockSkew) {
		return false
	}
	if !o.Region.matches(other.Region) {
		return false
	}

	return true
}

func (o *MethodCloudRegion) matches(other *MethodCloudRegion) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.RegionId, other.RegionId) {
		return false
	}
	if !o.Tenant.matches(other.Tenant) {
		return false
	}

	return true
}

func (o *MethodCloudRegionTenant) matches(other *MethodCloudRegionTenant) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Profile.matches(other.Profile) {
		return false
	}
	if !util.StringsMatch(o.TenantId, other.TenantId) {
		return false
	}

	return true
}

func (o *MethodCloudRegionTenantProfile) matches(other *MethodCloudRegionTenantProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Mfa.matches(other.Mfa) {
		return false
	}
	if !util.StringsMatch(o.ProfileId, other.ProfileId) {
		return false
	}

	return true
}

func (o *MethodCloudRegionTenantProfileMfa) matches(other *MethodCloudRegionTenantProfileMfa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ForceMfa, other.ForceMfa) {
		return false
	}

	return true
}

func (o *MethodKerberos) matches(other *MethodKerberos) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Realm, other.Realm) {
		return false
	}
	if !util.StringsMatch(o.ServerProfile, other.ServerProfile) {
		return false
	}

	return true
}

func (o *MethodLdap) matches(other *MethodLdap) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.LoginAttribute, other.LoginAttribute) {
		return false
	}
	if !util.Ints64Match(o.PasswdExpDays, other.PasswdExpDays) {
		return false
	}
	if !util.StringsMatch(o.ServerProfile, other.ServerProfile) {
		return false
	}

	return true
}

func (o *MethodLocalDatabase) matches(other *MethodLocalDatabase) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *MethodNone) matches(other *MethodNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *MethodRadius) matches(other *MethodRadius) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Checkgroup, other.Checkgroup) {
		return false
	}
	if !util.StringsMatch(o.ServerProfile, other.ServerProfile) {
		return false
	}

	return true
}

func (o *MethodSamlIdp) matches(other *MethodSamlIdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameAccessDomain, other.AttributeNameAccessDomain) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameAdminRole, other.AttributeNameAdminRole) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameUsergroup, other.AttributeNameUsergroup) {
		return false
	}
	if !util.StringsMatch(o.AttributeNameUsername, other.AttributeNameUsername) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.BoolsMatch(o.EnableSingleLogout, other.EnableSingleLogout) {
		return false
	}
	if !util.StringsMatch(o.RequestSigningCertificate, other.RequestSigningCertificate) {
		return false
	}
	if !util.StringsMatch(o.ServerProfile, other.ServerProfile) {
		return false
	}

	return true
}

func (o *MethodTacplus) matches(other *MethodTacplus) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Checkgroup, other.Checkgroup) {
		return false
	}
	if !util.StringsMatch(o.ServerProfile, other.ServerProfile) {
		return false
	}

	return true
}

func (o *MultiFactorAuth) matches(other *MultiFactorAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Factors, other.Factors) {
		return false
	}
	if !util.BoolsMatch(o.MfaEnable, other.MfaEnable) {
		return false
	}

	return true
}

func (o *SingleSignOn) matches(other *SingleSignOn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.KerberosKeytab, other.KerberosKeytab) {
		return false
	}
	if !util.StringsMatch(o.Realm, other.Realm) {
		return false
	}
	if !util.StringsMatch(o.ServicePrincipal, other.ServicePrincipal) {
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
