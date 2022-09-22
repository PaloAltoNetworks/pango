package authentication

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of an
// authentication profile.
type Entry struct {
	Name                      string
	LockoutFailedAttempts     string
	LockoutTime               int
	AllowList                 []string
	Type                      AuthenticationType
	UsernameModifier          string                     // 7.0+
	UserDomain                string                     // 7.0+
	SingleSignOn              *SingleSignOn              // 7.0+
	MultiFactorAuthentication *MultiFactorAuthentication // 8.0+
}

type AuthenticationType struct {
	None          bool
	LocalDatabase bool
	Radius        *Radius
	Ldap          *Ldap
	Kerberos      *Kerberos
	TacacsPlus    *TacacsPlus // 7.0+
	Saml          *Saml       // 8.0+
	//Cloud *Cloud // 10.1+
}

type Radius struct {
	ServerProfile     string
	RetrieveUserGroup bool // 7.0+
}

type Ldap struct {
	ServerProfile         string
	LoginAttribute        string
	PasswordExpiryWarning string // min 0 in pan-os 8.0+
}

type Kerberos struct {
	ServerProfile string
	Realm         string // 7.0+
}

type TacacsPlus struct {
	ServerProfile     string
	RetrieveUserGroup bool // 8.0+
}

type Saml struct {
	ServerProfile             string
	EnableSingleLogout        bool
	RequestSigningCertificate string
	CertificateProfile        string
	UsernameAttribute         string // default "username"
	UserGroupAttribute        string
	AdminRoleAttribute        string
	AccessDomainAttribute     string
}

type SingleSignOn struct {
	Realm            string
	ServicePrincipal string
	Keytab           string // encrypted
}

type MultiFactorAuthentication struct {
	Enabled bool
	Factors []string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.LockoutFailedAttempts = s.LockoutFailedAttempts
	o.LockoutTime = s.LockoutTime
	o.AllowList = util.CopyStringSlice(s.AllowList)
	o.UsernameModifier = s.UsernameModifier
	o.UserDomain = s.UserDomain

	// Single sign on.
	if s.SingleSignOn == nil {
		o.SingleSignOn = nil
	} else {
		o.SingleSignOn = &SingleSignOn{
			Realm:            s.SingleSignOn.Realm,
			ServicePrincipal: s.SingleSignOn.ServicePrincipal,
			Keytab:           s.SingleSignOn.Keytab,
		}
	}

	// Copy Entry.Type.
	switch {
	case s.Type.None:
		o.Type = AuthenticationType{None: true}
	case s.Type.LocalDatabase:
		o.Type = AuthenticationType{LocalDatabase: true}
	case s.Type.Radius != nil:
		o.Type = AuthenticationType{Radius: &Radius{
			ServerProfile:     s.Type.Radius.ServerProfile,
			RetrieveUserGroup: s.Type.Radius.RetrieveUserGroup,
		}}
	case s.Type.Ldap != nil:
		o.Type = AuthenticationType{Ldap: &Ldap{
			ServerProfile:         s.Type.Ldap.ServerProfile,
			LoginAttribute:        s.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: s.Type.Ldap.PasswordExpiryWarning,
		}}
	case s.Type.Kerberos != nil:
		o.Type = AuthenticationType{Kerberos: &Kerberos{
			ServerProfile: s.Type.Kerberos.ServerProfile,
			Realm:         s.Type.Kerberos.Realm,
		}}
	case s.Type.TacacsPlus != nil:
		o.Type = AuthenticationType{TacacsPlus: &TacacsPlus{
			ServerProfile:     s.Type.TacacsPlus.ServerProfile,
			RetrieveUserGroup: s.Type.TacacsPlus.RetrieveUserGroup,
		}}
	case s.Type.Saml != nil:
		o.Type = AuthenticationType{Saml: &Saml{
			ServerProfile:             s.Type.Saml.ServerProfile,
			EnableSingleLogout:        s.Type.Saml.EnableSingleLogout,
			RequestSigningCertificate: s.Type.Saml.RequestSigningCertificate,
			CertificateProfile:        s.Type.Saml.CertificateProfile,
			UsernameAttribute:         s.Type.Saml.UsernameAttribute,
			UserGroupAttribute:        s.Type.Saml.UserGroupAttribute,
			AdminRoleAttribute:        s.Type.Saml.AdminRoleAttribute,
			AccessDomainAttribute:     s.Type.Saml.AccessDomainAttribute,
		}}
	}

	// MFA.
	if s.MultiFactorAuthentication == nil {
		o.MultiFactorAuthentication = nil
	} else {
		o.MultiFactorAuthentication = &MultiFactorAuthentication{
			Enabled: s.MultiFactorAuthentication.Enabled,
			Factors: util.CopyStringSlice(s.MultiFactorAuthentication.Factors),
		}
	}
}

/** Structs / functions for this namespace. **/

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

func (o container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:      o.Name,
		AllowList: util.MemToStr(o.AllowList),
	}

	if o.Lockout != nil {
		ans.LockoutFailedAttempts = o.Lockout.LockoutFailedAttempts
		ans.LockoutTime = o.Lockout.LockoutTime
	}

	switch {
	case o.Type.None != nil:
		ans.Type.None = true
	case o.Type.LocalDatabase != nil:
		ans.Type.LocalDatabase = true
	case o.Type.Radius != nil:
		ans.Type.Radius = &Radius{
			ServerProfile: o.Type.Radius.ServerProfile,
		}
	case o.Type.Ldap != nil:
		ans.Type.Ldap = &Ldap{
			ServerProfile:         o.Type.Ldap.ServerProfile,
			LoginAttribute:        o.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: o.Type.Ldap.PasswordExpiryWarning,
		}
	case o.Type.Kerberos != nil:
		ans.Type.Kerberos = &Kerberos{
			ServerProfile: o.Type.Kerberos.ServerProfile,
		}
	}

	return ans
}

type entry_v1 struct {
	XMLName   xml.Name         `xml:"entry"`
	Name      string           `xml:"name,attr"`
	Lockout   *lockout         `xml:"lockout"`
	AllowList *util.MemberType `xml:"allow-list"`
	Type      method_v1        `xml:"method"`
}

func (e *entry_v1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v1
	ans := local{}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	if ans.Type.Ldap != nil && ans.Type.Ldap.PasswordExpiryWarning == "" {
		ans.Type.Ldap.PasswordExpiryWarning = "7"
	}
	*e = entry_v1(ans)
	return nil
}

type lockout struct {
	LockoutFailedAttempts string `xml:"failed-attempts,omitempty"`
	LockoutTime           int    `xml:"lockout-time"`
}

type method_v1 struct {
	None          *string      `xml:"none"`
	LocalDatabase *string      `xml:"local-database"`
	Radius        *radius_v1   `xml:"radius"`
	Ldap          *ldap        `xml:"ldap"`
	Kerberos      *kerberos_v1 `xml:"kerberos"`
}

type radius_v1 struct {
	ServerProfile string `xml:"server-profile"`
}

type ldap struct {
	ServerProfile         string `xml:"server-profile"`
	LoginAttribute        string `xml:"login-attribute,omitempty"`
	PasswordExpiryWarning string `xml:"passwd-exp-days,omitempty"`
}

type kerberos_v1 struct {
	ServerProfile string `xml:"server-profile"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:      e.Name,
		AllowList: util.StrToMem(e.AllowList),
	}

	if e.LockoutFailedAttempts != "" || e.LockoutTime > 0 {
		ans.Lockout = &lockout{
			LockoutFailedAttempts: e.LockoutFailedAttempts,
			LockoutTime:           e.LockoutTime,
		}
	}

	s := ""

	switch {
	case e.Type.None:
		ans.Type = method_v1{None: &s}
	case e.Type.LocalDatabase:
		ans.Type = method_v1{LocalDatabase: &s}
	case e.Type.Radius != nil:
		ans.Type = method_v1{Radius: &radius_v1{
			ServerProfile: e.Type.Radius.ServerProfile,
		}}
	case e.Type.Ldap != nil:
		ans.Type = method_v1{Ldap: &ldap{
			ServerProfile:         e.Type.Ldap.ServerProfile,
			LoginAttribute:        e.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: e.Type.Ldap.PasswordExpiryWarning,
		}}
	case e.Type.Kerberos != nil:
		ans.Type = method_v1{Kerberos: &kerberos_v1{
			ServerProfile: e.Type.Kerberos.ServerProfile,
		}}
	}

	return ans
}

// PAN-OS 7.0
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v2) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name:             o.Name,
		AllowList:        util.MemToStr(o.AllowList),
		UsernameModifier: o.UsernameModifier,
		UserDomain:       o.UserDomain,
	}

	if o.Lockout != nil {
		ans.LockoutFailedAttempts = o.Lockout.LockoutFailedAttempts
		ans.LockoutTime = o.Lockout.LockoutTime
	}

	if o.Sso != nil {
		ans.SingleSignOn = &SingleSignOn{
			Realm:            o.Sso.Realm,
			ServicePrincipal: o.Sso.ServicePrincipal,
			Keytab:           o.Sso.Keytab,
		}
	}

	switch {
	case o.Type.None != nil:
		ans.Type.None = true
	case o.Type.LocalDatabase != nil:
		ans.Type.LocalDatabase = true
	case o.Type.Radius != nil:
		ans.Type.Radius = &Radius{
			ServerProfile:     o.Type.Radius.ServerProfile,
			RetrieveUserGroup: util.AsBool(o.Type.Radius.RetrieveUserGroup),
		}
	case o.Type.Ldap != nil:
		ans.Type.Ldap = &Ldap{
			ServerProfile:         o.Type.Ldap.ServerProfile,
			LoginAttribute:        o.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: o.Type.Ldap.PasswordExpiryWarning,
		}
	case o.Type.Kerberos != nil:
		ans.Type.Kerberos = &Kerberos{
			ServerProfile: o.Type.Kerberos.ServerProfile,
			Realm:         o.Type.Kerberos.Realm,
		}
	case o.Type.TacacsPlus != nil:
		ans.Type.TacacsPlus = &TacacsPlus{
			ServerProfile: o.Type.TacacsPlus.ServerProfile,
		}
	}

	return ans
}

type entry_v2 struct {
	XMLName          xml.Name         `xml:"entry"`
	Name             string           `xml:"name,attr"`
	Lockout          *lockout         `xml:"lockout"`
	AllowList        *util.MemberType `xml:"allow-list"`
	Type             method_v2        `xml:"method"`
	UsernameModifier string           `xml:"username-modifier,omitempty"`
	UserDomain       string           `xml:"user-domain,omitempty"`
	Sso              *sso             `xml:"single-sign-on"`
}

type sso struct {
	Realm            string `xml:"realm,omitempty"`
	ServicePrincipal string `xml:"service-principal,omitempty"`
	Keytab           string `xml:"kerberos-keytab,omitempty"`
}

func (e *entry_v2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v2
	ans := local{
		UsernameModifier: UsernameModifierInput,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	if ans.Type.Ldap != nil && ans.Type.Ldap.PasswordExpiryWarning == "" {
		ans.Type.Ldap.PasswordExpiryWarning = "7"
	}
	*e = entry_v2(ans)
	return nil
}

type method_v2 struct {
	None          *string        `xml:"none"`
	LocalDatabase *string        `xml:"local-database"`
	Radius        *radius_v2     `xml:"radius"`
	Ldap          *ldap          `xml:"ldap"`
	Kerberos      *kerberos_v2   `xml:"kerberos"`
	TacacsPlus    *tacacsPlus_v1 `xml:"tacplus"`
}

type radius_v2 struct {
	ServerProfile     string `xml:"server-profile"`
	RetrieveUserGroup string `xml:"checkgroup"`
}

type kerberos_v2 struct {
	ServerProfile string `xml:"server-profile"`
	Realm         string `xml:"realm"`
}

type tacacsPlus_v1 struct {
	ServerProfile string `xml:"server-profile"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name:             e.Name,
		AllowList:        util.StrToMem(e.AllowList),
		UsernameModifier: e.UsernameModifier,
		UserDomain:       e.UserDomain,
	}

	if e.SingleSignOn != nil {
		ans.Sso = &sso{
			Realm:            e.SingleSignOn.Realm,
			ServicePrincipal: e.SingleSignOn.ServicePrincipal,
			Keytab:           e.SingleSignOn.Keytab,
		}
	}

	if e.LockoutFailedAttempts != "" || e.LockoutTime > 0 {
		ans.Lockout = &lockout{
			LockoutFailedAttempts: e.LockoutFailedAttempts,
			LockoutTime:           e.LockoutTime,
		}
	}

	s := ""

	switch {
	case e.Type.None:
		ans.Type = method_v2{None: &s}
	case e.Type.LocalDatabase:
		ans.Type = method_v2{LocalDatabase: &s}
	case e.Type.Radius != nil:
		ans.Type = method_v2{Radius: &radius_v2{
			ServerProfile:     e.Type.Radius.ServerProfile,
			RetrieveUserGroup: util.YesNo(e.Type.Radius.RetrieveUserGroup),
		}}
	case e.Type.Ldap != nil:
		ans.Type = method_v2{Ldap: &ldap{
			ServerProfile:         e.Type.Ldap.ServerProfile,
			LoginAttribute:        e.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: e.Type.Ldap.PasswordExpiryWarning,
		}}
	case e.Type.Kerberos != nil:
		ans.Type = method_v2{Kerberos: &kerberos_v2{
			ServerProfile: e.Type.Kerberos.ServerProfile,
			Realm:         e.Type.Kerberos.Realm,
		}}
	case e.Type.TacacsPlus != nil:
		ans.Type = method_v2{TacacsPlus: &tacacsPlus_v1{
			ServerProfile: e.Type.TacacsPlus.ServerProfile,
		}}
	}

	return ans
}

// PAN-OS 8.0
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v3) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name:             o.Name,
		AllowList:        util.MemToStr(o.AllowList),
		UsernameModifier: o.UsernameModifier,
		UserDomain:       o.UserDomain,
	}

	if o.Lockout != nil {
		ans.LockoutFailedAttempts = o.Lockout.LockoutFailedAttempts
		ans.LockoutTime = o.Lockout.LockoutTime
	}

	if o.Sso != nil {
		ans.SingleSignOn = &SingleSignOn{
			Realm:            o.Sso.Realm,
			ServicePrincipal: o.Sso.ServicePrincipal,
			Keytab:           o.Sso.Keytab,
		}
	}

	switch {
	case o.Type.None != nil:
		ans.Type.None = true
	case o.Type.LocalDatabase != nil:
		ans.Type.LocalDatabase = true
	case o.Type.Radius != nil:
		ans.Type.Radius = &Radius{
			ServerProfile:     o.Type.Radius.ServerProfile,
			RetrieveUserGroup: util.AsBool(o.Type.Radius.RetrieveUserGroup),
		}
	case o.Type.Ldap != nil:
		ans.Type.Ldap = &Ldap{
			ServerProfile:         o.Type.Ldap.ServerProfile,
			LoginAttribute:        o.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: o.Type.Ldap.PasswordExpiryWarning,
		}
	case o.Type.Kerberos != nil:
		ans.Type.Kerberos = &Kerberos{
			ServerProfile: o.Type.Kerberos.ServerProfile,
			Realm:         o.Type.Kerberos.Realm,
		}
	case o.Type.TacacsPlus != nil:
		ans.Type.TacacsPlus = &TacacsPlus{
			ServerProfile:     o.Type.TacacsPlus.ServerProfile,
			RetrieveUserGroup: util.AsBool(o.Type.TacacsPlus.RetrieveUserGroup),
		}
	case o.Type.Saml != nil:
		ans.Type.Saml = &Saml{
			ServerProfile:             o.Type.Saml.ServerProfile,
			EnableSingleLogout:        util.AsBool(o.Type.Saml.EnableSingleLogout),
			RequestSigningCertificate: o.Type.Saml.RequestSigningCertificate,
			CertificateProfile:        o.Type.Saml.CertificateProfile,
			UsernameAttribute:         o.Type.Saml.UsernameAttribute,
			UserGroupAttribute:        o.Type.Saml.UserGroupAttribute,
			AdminRoleAttribute:        o.Type.Saml.AdminRoleAttribute,
			AccessDomainAttribute:     o.Type.Saml.AccessDomainAttribute,
		}
	}

	if o.MultiFactorAuthentication != nil {
		ans.MultiFactorAuthentication = &MultiFactorAuthentication{
			Enabled: util.AsBool(o.MultiFactorAuthentication.Enabled),
			Factors: util.MemToStr(o.MultiFactorAuthentication.Factors),
		}
	}

	return ans
}

type entry_v3 struct {
	XMLName                   xml.Name         `xml:"entry"`
	Name                      string           `xml:"name,attr"`
	Lockout                   *lockout         `xml:"lockout"`
	AllowList                 *util.MemberType `xml:"allow-list"`
	Type                      method_v3        `xml:"method"`
	UsernameModifier          string           `xml:"username-modifier,omitempty"`
	UserDomain                string           `xml:"user-domain,omitempty"`
	Sso                       *sso             `xml:"single-sign-on"`
	MultiFactorAuthentication *mfa             `xml:"multi-factor-auth"`
}

func (e *entry_v3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local entry_v3
	ans := local{
		UsernameModifier: UsernameModifierInput,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	if ans.Type.Ldap != nil && ans.Type.Ldap.PasswordExpiryWarning == "" {
		ans.Type.Ldap.PasswordExpiryWarning = "7"
	}
	if ans.Type.Saml != nil && ans.Type.Saml.UsernameAttribute == "" {
		ans.Type.Saml.UsernameAttribute = UsernameAttributeDefault
	}
	*e = entry_v3(ans)
	return nil
}

type method_v3 struct {
	None          *string        `xml:"none"`
	LocalDatabase *string        `xml:"local-database"`
	Radius        *radius_v2     `xml:"radius"`
	Ldap          *ldap          `xml:"ldap"`
	Kerberos      *kerberos_v2   `xml:"kerberos"`
	TacacsPlus    *tacacsPlus_v2 `xml:"tacplus"`
	Saml          *saml          `xml:"saml-idp"`
}

type tacacsPlus_v2 struct {
	ServerProfile     string `xml:"server-profile"`
	RetrieveUserGroup string `xml:"checkgroup"`
}

type saml struct {
	ServerProfile             string `xml:"server-profile"`
	EnableSingleLogout        string `xml:"enable-single-logout"`
	RequestSigningCertificate string `xml:"request-signing-certificate,omitempty"`
	CertificateProfile        string `xml:"certificate-profile,omitempty"`
	UsernameAttribute         string `xml:"attribute-name-username"`
	UserGroupAttribute        string `xml:"attribute-name-usergroup,omitempty"`
	AdminRoleAttribute        string `xml:"attribute-name-admin-role,omitempty"`
	AccessDomainAttribute     string `xml:"attribute-name-access-domain,omitempty"`
}

type mfa struct {
	Enabled string           `xml:"mfa-enable"`
	Factors *util.MemberType `xml:"factors"`
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name:             e.Name,
		AllowList:        util.StrToMem(e.AllowList),
		UsernameModifier: e.UsernameModifier,
		UserDomain:       e.UserDomain,
	}

	if e.SingleSignOn != nil {
		ans.Sso = &sso{
			Realm:            e.SingleSignOn.Realm,
			ServicePrincipal: e.SingleSignOn.ServicePrincipal,
			Keytab:           e.SingleSignOn.Keytab,
		}
	}

	if e.LockoutFailedAttempts != "" || e.LockoutTime > 0 {
		ans.Lockout = &lockout{
			LockoutFailedAttempts: e.LockoutFailedAttempts,
			LockoutTime:           e.LockoutTime,
		}
	}

	s := ""

	switch {
	case e.Type.None:
		ans.Type = method_v3{None: &s}
	case e.Type.LocalDatabase:
		ans.Type = method_v3{LocalDatabase: &s}
	case e.Type.Radius != nil:
		ans.Type = method_v3{Radius: &radius_v2{
			ServerProfile:     e.Type.Radius.ServerProfile,
			RetrieveUserGroup: util.YesNo(e.Type.Radius.RetrieveUserGroup),
		}}
	case e.Type.Ldap != nil:
		ans.Type = method_v3{Ldap: &ldap{
			ServerProfile:         e.Type.Ldap.ServerProfile,
			LoginAttribute:        e.Type.Ldap.LoginAttribute,
			PasswordExpiryWarning: e.Type.Ldap.PasswordExpiryWarning,
		}}
	case e.Type.Kerberos != nil:
		ans.Type = method_v3{Kerberos: &kerberos_v2{
			ServerProfile: e.Type.Kerberos.ServerProfile,
			Realm:         e.Type.Kerberos.Realm,
		}}
	case e.Type.TacacsPlus != nil:
		ans.Type = method_v3{TacacsPlus: &tacacsPlus_v2{
			ServerProfile:     e.Type.TacacsPlus.ServerProfile,
			RetrieveUserGroup: util.YesNo(e.Type.TacacsPlus.RetrieveUserGroup),
		}}
	case e.Type.Saml != nil:
		ans.Type = method_v3{Saml: &saml{
			ServerProfile:             e.Type.Saml.ServerProfile,
			EnableSingleLogout:        util.YesNo(e.Type.Saml.EnableSingleLogout),
			RequestSigningCertificate: e.Type.Saml.RequestSigningCertificate,
			CertificateProfile:        e.Type.Saml.CertificateProfile,
			UsernameAttribute:         e.Type.Saml.UsernameAttribute,
			UserGroupAttribute:        e.Type.Saml.UserGroupAttribute,
			AdminRoleAttribute:        e.Type.Saml.AdminRoleAttribute,
			AccessDomainAttribute:     e.Type.Saml.AccessDomainAttribute,
		}}
	}

	if e.MultiFactorAuthentication != nil {
		ans.MultiFactorAuthentication = &mfa{
			Enabled: util.YesNo(e.MultiFactorAuthentication.Enabled),
			Factors: util.StrToMem(e.MultiFactorAuthentication.Factors),
		}
	}

	return ans
}
