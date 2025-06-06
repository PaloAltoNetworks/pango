package extdynlist

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
	suffix = []string{"external-list", "$name"}
)

type Entry struct {
	Name            string
	DisableOverride *string
	Type            *Type
	Misc            []generic.Xml
}
type Type struct {
	Domain        *TypeDomain
	Imei          *TypeImei
	Imsi          *TypeImsi
	Ip            *TypeIp
	PredefinedIp  *TypePredefinedIp
	PredefinedUrl *TypePredefinedUrl
	Url           *TypeUrl
	Misc          []generic.Xml
}
type TypeDomain struct {
	Auth               *TypeDomainAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	ExpandDomain       *bool
	Recurring          *TypeDomainRecurring
	Url                *string
	Misc               []generic.Xml
}
type TypeDomainAuth struct {
	Password *string
	Username *string
	Misc     []generic.Xml
}
type TypeDomainRecurring struct {
	Daily      *TypeDomainRecurringDaily
	FiveMinute *TypeDomainRecurringFiveMinute
	Hourly     *TypeDomainRecurringHourly
	Monthly    *TypeDomainRecurringMonthly
	Weekly     *TypeDomainRecurringWeekly
	Misc       []generic.Xml
}
type TypeDomainRecurringDaily struct {
	At   *string
	Misc []generic.Xml
}
type TypeDomainRecurringFiveMinute struct {
	Misc []generic.Xml
}
type TypeDomainRecurringHourly struct {
	Misc []generic.Xml
}
type TypeDomainRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
	Misc       []generic.Xml
}
type TypeDomainRecurringWeekly struct {
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type TypeImei struct {
	Auth               *TypeImeiAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeImeiRecurring
	Url                *string
	Misc               []generic.Xml
}
type TypeImeiAuth struct {
	Password *string
	Username *string
	Misc     []generic.Xml
}
type TypeImeiRecurring struct {
	Daily      *TypeImeiRecurringDaily
	FiveMinute *TypeImeiRecurringFiveMinute
	Hourly     *TypeImeiRecurringHourly
	Monthly    *TypeImeiRecurringMonthly
	Weekly     *TypeImeiRecurringWeekly
	Misc       []generic.Xml
}
type TypeImeiRecurringDaily struct {
	At   *string
	Misc []generic.Xml
}
type TypeImeiRecurringFiveMinute struct {
	Misc []generic.Xml
}
type TypeImeiRecurringHourly struct {
	Misc []generic.Xml
}
type TypeImeiRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
	Misc       []generic.Xml
}
type TypeImeiRecurringWeekly struct {
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type TypeImsi struct {
	Auth               *TypeImsiAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeImsiRecurring
	Url                *string
	Misc               []generic.Xml
}
type TypeImsiAuth struct {
	Password *string
	Username *string
	Misc     []generic.Xml
}
type TypeImsiRecurring struct {
	Daily      *TypeImsiRecurringDaily
	FiveMinute *TypeImsiRecurringFiveMinute
	Hourly     *TypeImsiRecurringHourly
	Monthly    *TypeImsiRecurringMonthly
	Weekly     *TypeImsiRecurringWeekly
	Misc       []generic.Xml
}
type TypeImsiRecurringDaily struct {
	At   *string
	Misc []generic.Xml
}
type TypeImsiRecurringFiveMinute struct {
	Misc []generic.Xml
}
type TypeImsiRecurringHourly struct {
	Misc []generic.Xml
}
type TypeImsiRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
	Misc       []generic.Xml
}
type TypeImsiRecurringWeekly struct {
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type TypeIp struct {
	Auth               *TypeIpAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeIpRecurring
	Url                *string
	Misc               []generic.Xml
}
type TypeIpAuth struct {
	Password *string
	Username *string
	Misc     []generic.Xml
}
type TypeIpRecurring struct {
	Daily      *TypeIpRecurringDaily
	FiveMinute *TypeIpRecurringFiveMinute
	Hourly     *TypeIpRecurringHourly
	Monthly    *TypeIpRecurringMonthly
	Weekly     *TypeIpRecurringWeekly
	Misc       []generic.Xml
}
type TypeIpRecurringDaily struct {
	At   *string
	Misc []generic.Xml
}
type TypeIpRecurringFiveMinute struct {
	Misc []generic.Xml
}
type TypeIpRecurringHourly struct {
	Misc []generic.Xml
}
type TypeIpRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
	Misc       []generic.Xml
}
type TypeIpRecurringWeekly struct {
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type TypePredefinedIp struct {
	Description   *string
	ExceptionList []string
	Url           *string
	Misc          []generic.Xml
}
type TypePredefinedUrl struct {
	Description   *string
	ExceptionList []string
	Url           *string
	Misc          []generic.Xml
}
type TypeUrl struct {
	Auth               *TypeUrlAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeUrlRecurring
	Url                *string
	Misc               []generic.Xml
}
type TypeUrlAuth struct {
	Password *string
	Username *string
	Misc     []generic.Xml
}
type TypeUrlRecurring struct {
	Daily      *TypeUrlRecurringDaily
	FiveMinute *TypeUrlRecurringFiveMinute
	Hourly     *TypeUrlRecurringHourly
	Monthly    *TypeUrlRecurringMonthly
	Weekly     *TypeUrlRecurringWeekly
	Misc       []generic.Xml
}
type TypeUrlRecurringDaily struct {
	At   *string
	Misc []generic.Xml
}
type TypeUrlRecurringFiveMinute struct {
	Misc []generic.Xml
}
type TypeUrlRecurringHourly struct {
	Misc []generic.Xml
}
type TypeUrlRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
	Misc       []generic.Xml
}
type TypeUrlRecurringWeekly struct {
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
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
	XMLName         xml.Name      `xml:"entry"`
	Name            string        `xml:"name,attr"`
	DisableOverride *string       `xml:"disable-override,omitempty"`
	Type            *typeXml      `xml:"type,omitempty"`
	Misc            []generic.Xml `xml:",any"`
}
type typeXml struct {
	Domain        *typeDomainXml        `xml:"domain,omitempty"`
	Imei          *typeImeiXml          `xml:"imei,omitempty"`
	Imsi          *typeImsiXml          `xml:"imsi,omitempty"`
	Ip            *typeIpXml            `xml:"ip,omitempty"`
	PredefinedIp  *typePredefinedIpXml  `xml:"predefined-ip,omitempty"`
	PredefinedUrl *typePredefinedUrlXml `xml:"predefined-url,omitempty"`
	Url           *typeUrlXml           `xml:"url,omitempty"`
	Misc          []generic.Xml         `xml:",any"`
}
type typeDomainXml struct {
	Auth               *typeDomainAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string                 `xml:"certificate-profile,omitempty"`
	Description        *string                 `xml:"description,omitempty"`
	ExceptionList      *util.MemberType        `xml:"exception-list,omitempty"`
	ExpandDomain       *string                 `xml:"expand-domain,omitempty"`
	Recurring          *typeDomainRecurringXml `xml:"recurring,omitempty"`
	Url                *string                 `xml:"url,omitempty"`
	Misc               []generic.Xml           `xml:",any"`
}
type typeDomainAuthXml struct {
	Password *string       `xml:"password,omitempty"`
	Username *string       `xml:"username,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type typeDomainRecurringXml struct {
	Daily      *typeDomainRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *typeDomainRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *typeDomainRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *typeDomainRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *typeDomainRecurringWeeklyXml     `xml:"weekly,omitempty"`
	Misc       []generic.Xml                     `xml:",any"`
}
type typeDomainRecurringDailyXml struct {
	At   *string       `xml:"at,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type typeDomainRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeDomainRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeDomainRecurringMonthlyXml struct {
	At         *string       `xml:"at,omitempty"`
	DayOfMonth *int64        `xml:"day-of-month,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type typeDomainRecurringWeeklyXml struct {
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type typeImeiXml struct {
	Auth               *typeImeiAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string               `xml:"certificate-profile,omitempty"`
	Description        *string               `xml:"description,omitempty"`
	ExceptionList      *util.MemberType      `xml:"exception-list,omitempty"`
	Recurring          *typeImeiRecurringXml `xml:"recurring,omitempty"`
	Url                *string               `xml:"url,omitempty"`
	Misc               []generic.Xml         `xml:",any"`
}
type typeImeiAuthXml struct {
	Password *string       `xml:"password,omitempty"`
	Username *string       `xml:"username,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type typeImeiRecurringXml struct {
	Daily      *typeImeiRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *typeImeiRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *typeImeiRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *typeImeiRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *typeImeiRecurringWeeklyXml     `xml:"weekly,omitempty"`
	Misc       []generic.Xml                   `xml:",any"`
}
type typeImeiRecurringDailyXml struct {
	At   *string       `xml:"at,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type typeImeiRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeImeiRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeImeiRecurringMonthlyXml struct {
	At         *string       `xml:"at,omitempty"`
	DayOfMonth *int64        `xml:"day-of-month,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type typeImeiRecurringWeeklyXml struct {
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type typeImsiXml struct {
	Auth               *typeImsiAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string               `xml:"certificate-profile,omitempty"`
	Description        *string               `xml:"description,omitempty"`
	ExceptionList      *util.MemberType      `xml:"exception-list,omitempty"`
	Recurring          *typeImsiRecurringXml `xml:"recurring,omitempty"`
	Url                *string               `xml:"url,omitempty"`
	Misc               []generic.Xml         `xml:",any"`
}
type typeImsiAuthXml struct {
	Password *string       `xml:"password,omitempty"`
	Username *string       `xml:"username,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type typeImsiRecurringXml struct {
	Daily      *typeImsiRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *typeImsiRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *typeImsiRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *typeImsiRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *typeImsiRecurringWeeklyXml     `xml:"weekly,omitempty"`
	Misc       []generic.Xml                   `xml:",any"`
}
type typeImsiRecurringDailyXml struct {
	At   *string       `xml:"at,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type typeImsiRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeImsiRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeImsiRecurringMonthlyXml struct {
	At         *string       `xml:"at,omitempty"`
	DayOfMonth *int64        `xml:"day-of-month,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type typeImsiRecurringWeeklyXml struct {
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type typeIpXml struct {
	Auth               *typeIpAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string             `xml:"certificate-profile,omitempty"`
	Description        *string             `xml:"description,omitempty"`
	ExceptionList      *util.MemberType    `xml:"exception-list,omitempty"`
	Recurring          *typeIpRecurringXml `xml:"recurring,omitempty"`
	Url                *string             `xml:"url,omitempty"`
	Misc               []generic.Xml       `xml:",any"`
}
type typeIpAuthXml struct {
	Password *string       `xml:"password,omitempty"`
	Username *string       `xml:"username,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type typeIpRecurringXml struct {
	Daily      *typeIpRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *typeIpRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *typeIpRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *typeIpRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *typeIpRecurringWeeklyXml     `xml:"weekly,omitempty"`
	Misc       []generic.Xml                 `xml:",any"`
}
type typeIpRecurringDailyXml struct {
	At   *string       `xml:"at,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type typeIpRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeIpRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeIpRecurringMonthlyXml struct {
	At         *string       `xml:"at,omitempty"`
	DayOfMonth *int64        `xml:"day-of-month,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type typeIpRecurringWeeklyXml struct {
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type typePredefinedIpXml struct {
	Description   *string          `xml:"description,omitempty"`
	ExceptionList *util.MemberType `xml:"exception-list,omitempty"`
	Url           *string          `xml:"url,omitempty"`
	Misc          []generic.Xml    `xml:",any"`
}
type typePredefinedUrlXml struct {
	Description   *string          `xml:"description,omitempty"`
	ExceptionList *util.MemberType `xml:"exception-list,omitempty"`
	Url           *string          `xml:"url,omitempty"`
	Misc          []generic.Xml    `xml:",any"`
}
type typeUrlXml struct {
	Auth               *typeUrlAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string              `xml:"certificate-profile,omitempty"`
	Description        *string              `xml:"description,omitempty"`
	ExceptionList      *util.MemberType     `xml:"exception-list,omitempty"`
	Recurring          *typeUrlRecurringXml `xml:"recurring,omitempty"`
	Url                *string              `xml:"url,omitempty"`
	Misc               []generic.Xml        `xml:",any"`
}
type typeUrlAuthXml struct {
	Password *string       `xml:"password,omitempty"`
	Username *string       `xml:"username,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type typeUrlRecurringXml struct {
	Daily      *typeUrlRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *typeUrlRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *typeUrlRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *typeUrlRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *typeUrlRecurringWeeklyXml     `xml:"weekly,omitempty"`
	Misc       []generic.Xml                  `xml:",any"`
}
type typeUrlRecurringDailyXml struct {
	At   *string       `xml:"at,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type typeUrlRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeUrlRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type typeUrlRecurringMonthlyXml struct {
	At         *string       `xml:"at,omitempty"`
	DayOfMonth *int64        `xml:"day-of-month,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type typeUrlRecurringWeeklyXml struct {
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DisableOverride = s.DisableOverride
	if s.Type != nil {
		var obj typeXml
		obj.MarshalFromObject(*s.Type)
		o.Type = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var typeVal *Type
	if o.Type != nil {
		obj, err := o.Type.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		typeVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		DisableOverride: o.DisableOverride,
		Type:            typeVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *typeXml) MarshalFromObject(s Type) {
	if s.Domain != nil {
		var obj typeDomainXml
		obj.MarshalFromObject(*s.Domain)
		o.Domain = &obj
	}
	if s.Imei != nil {
		var obj typeImeiXml
		obj.MarshalFromObject(*s.Imei)
		o.Imei = &obj
	}
	if s.Imsi != nil {
		var obj typeImsiXml
		obj.MarshalFromObject(*s.Imsi)
		o.Imsi = &obj
	}
	if s.Ip != nil {
		var obj typeIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	if s.PredefinedIp != nil {
		var obj typePredefinedIpXml
		obj.MarshalFromObject(*s.PredefinedIp)
		o.PredefinedIp = &obj
	}
	if s.PredefinedUrl != nil {
		var obj typePredefinedUrlXml
		obj.MarshalFromObject(*s.PredefinedUrl)
		o.PredefinedUrl = &obj
	}
	if s.Url != nil {
		var obj typeUrlXml
		obj.MarshalFromObject(*s.Url)
		o.Url = &obj
	}
	o.Misc = s.Misc
}

func (o typeXml) UnmarshalToObject() (*Type, error) {
	var domainVal *TypeDomain
	if o.Domain != nil {
		obj, err := o.Domain.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		domainVal = obj
	}
	var imeiVal *TypeImei
	if o.Imei != nil {
		obj, err := o.Imei.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		imeiVal = obj
	}
	var imsiVal *TypeImsi
	if o.Imsi != nil {
		obj, err := o.Imsi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		imsiVal = obj
	}
	var ipVal *TypeIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}
	var predefinedIpVal *TypePredefinedIp
	if o.PredefinedIp != nil {
		obj, err := o.PredefinedIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		predefinedIpVal = obj
	}
	var predefinedUrlVal *TypePredefinedUrl
	if o.PredefinedUrl != nil {
		obj, err := o.PredefinedUrl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		predefinedUrlVal = obj
	}
	var urlVal *TypeUrl
	if o.Url != nil {
		obj, err := o.Url.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		urlVal = obj
	}

	result := &Type{
		Domain:        domainVal,
		Imei:          imeiVal,
		Imsi:          imsiVal,
		Ip:            ipVal,
		PredefinedIp:  predefinedIpVal,
		PredefinedUrl: predefinedUrlVal,
		Url:           urlVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *typeDomainXml) MarshalFromObject(s TypeDomain) {
	if s.Auth != nil {
		var obj typeDomainAuthXml
		obj.MarshalFromObject(*s.Auth)
		o.Auth = &obj
	}
	o.CertificateProfile = s.CertificateProfile
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	o.ExpandDomain = util.YesNo(s.ExpandDomain, nil)
	if s.Recurring != nil {
		var obj typeDomainRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typeDomainXml) UnmarshalToObject() (*TypeDomain, error) {
	var authVal *TypeDomainAuth
	if o.Auth != nil {
		obj, err := o.Auth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authVal = obj
	}
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}
	var recurringVal *TypeDomainRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &TypeDomain{
		Auth:               authVal,
		CertificateProfile: o.CertificateProfile,
		Description:        o.Description,
		ExceptionList:      exceptionListVal,
		ExpandDomain:       util.AsBool(o.ExpandDomain, nil),
		Recurring:          recurringVal,
		Url:                o.Url,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *typeDomainAuthXml) MarshalFromObject(s TypeDomainAuth) {
	o.Password = s.Password
	o.Username = s.Username
	o.Misc = s.Misc
}

func (o typeDomainAuthXml) UnmarshalToObject() (*TypeDomainAuth, error) {

	result := &TypeDomainAuth{
		Password: o.Password,
		Username: o.Username,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringXml) MarshalFromObject(s TypeDomainRecurring) {
	if s.Daily != nil {
		var obj typeDomainRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.FiveMinute != nil {
		var obj typeDomainRecurringFiveMinuteXml
		obj.MarshalFromObject(*s.FiveMinute)
		o.FiveMinute = &obj
	}
	if s.Hourly != nil {
		var obj typeDomainRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.Monthly != nil {
		var obj typeDomainRecurringMonthlyXml
		obj.MarshalFromObject(*s.Monthly)
		o.Monthly = &obj
	}
	if s.Weekly != nil {
		var obj typeDomainRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o typeDomainRecurringXml) UnmarshalToObject() (*TypeDomainRecurring, error) {
	var dailyVal *TypeDomainRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var fiveMinuteVal *TypeDomainRecurringFiveMinute
	if o.FiveMinute != nil {
		obj, err := o.FiveMinute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fiveMinuteVal = obj
	}
	var hourlyVal *TypeDomainRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var monthlyVal *TypeDomainRecurringMonthly
	if o.Monthly != nil {
		obj, err := o.Monthly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monthlyVal = obj
	}
	var weeklyVal *TypeDomainRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &TypeDomainRecurring{
		Daily:      dailyVal,
		FiveMinute: fiveMinuteVal,
		Hourly:     hourlyVal,
		Monthly:    monthlyVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringDailyXml) MarshalFromObject(s TypeDomainRecurringDaily) {
	o.At = s.At
	o.Misc = s.Misc
}

func (o typeDomainRecurringDailyXml) UnmarshalToObject() (*TypeDomainRecurringDaily, error) {

	result := &TypeDomainRecurringDaily{
		At:   o.At,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringFiveMinuteXml) MarshalFromObject(s TypeDomainRecurringFiveMinute) {
	o.Misc = s.Misc
}

func (o typeDomainRecurringFiveMinuteXml) UnmarshalToObject() (*TypeDomainRecurringFiveMinute, error) {

	result := &TypeDomainRecurringFiveMinute{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringHourlyXml) MarshalFromObject(s TypeDomainRecurringHourly) {
	o.Misc = s.Misc
}

func (o typeDomainRecurringHourlyXml) UnmarshalToObject() (*TypeDomainRecurringHourly, error) {

	result := &TypeDomainRecurringHourly{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringMonthlyXml) MarshalFromObject(s TypeDomainRecurringMonthly) {
	o.At = s.At
	o.DayOfMonth = s.DayOfMonth
	o.Misc = s.Misc
}

func (o typeDomainRecurringMonthlyXml) UnmarshalToObject() (*TypeDomainRecurringMonthly, error) {

	result := &TypeDomainRecurringMonthly{
		At:         o.At,
		DayOfMonth: o.DayOfMonth,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeDomainRecurringWeeklyXml) MarshalFromObject(s TypeDomainRecurringWeekly) {
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o typeDomainRecurringWeeklyXml) UnmarshalToObject() (*TypeDomainRecurringWeekly, error) {

	result := &TypeDomainRecurringWeekly{
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *typeImeiXml) MarshalFromObject(s TypeImei) {
	if s.Auth != nil {
		var obj typeImeiAuthXml
		obj.MarshalFromObject(*s.Auth)
		o.Auth = &obj
	}
	o.CertificateProfile = s.CertificateProfile
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	if s.Recurring != nil {
		var obj typeImeiRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typeImeiXml) UnmarshalToObject() (*TypeImei, error) {
	var authVal *TypeImeiAuth
	if o.Auth != nil {
		obj, err := o.Auth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authVal = obj
	}
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}
	var recurringVal *TypeImeiRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &TypeImei{
		Auth:               authVal,
		CertificateProfile: o.CertificateProfile,
		Description:        o.Description,
		ExceptionList:      exceptionListVal,
		Recurring:          recurringVal,
		Url:                o.Url,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *typeImeiAuthXml) MarshalFromObject(s TypeImeiAuth) {
	o.Password = s.Password
	o.Username = s.Username
	o.Misc = s.Misc
}

func (o typeImeiAuthXml) UnmarshalToObject() (*TypeImeiAuth, error) {

	result := &TypeImeiAuth{
		Password: o.Password,
		Username: o.Username,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringXml) MarshalFromObject(s TypeImeiRecurring) {
	if s.Daily != nil {
		var obj typeImeiRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.FiveMinute != nil {
		var obj typeImeiRecurringFiveMinuteXml
		obj.MarshalFromObject(*s.FiveMinute)
		o.FiveMinute = &obj
	}
	if s.Hourly != nil {
		var obj typeImeiRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.Monthly != nil {
		var obj typeImeiRecurringMonthlyXml
		obj.MarshalFromObject(*s.Monthly)
		o.Monthly = &obj
	}
	if s.Weekly != nil {
		var obj typeImeiRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o typeImeiRecurringXml) UnmarshalToObject() (*TypeImeiRecurring, error) {
	var dailyVal *TypeImeiRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var fiveMinuteVal *TypeImeiRecurringFiveMinute
	if o.FiveMinute != nil {
		obj, err := o.FiveMinute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fiveMinuteVal = obj
	}
	var hourlyVal *TypeImeiRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var monthlyVal *TypeImeiRecurringMonthly
	if o.Monthly != nil {
		obj, err := o.Monthly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monthlyVal = obj
	}
	var weeklyVal *TypeImeiRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &TypeImeiRecurring{
		Daily:      dailyVal,
		FiveMinute: fiveMinuteVal,
		Hourly:     hourlyVal,
		Monthly:    monthlyVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringDailyXml) MarshalFromObject(s TypeImeiRecurringDaily) {
	o.At = s.At
	o.Misc = s.Misc
}

func (o typeImeiRecurringDailyXml) UnmarshalToObject() (*TypeImeiRecurringDaily, error) {

	result := &TypeImeiRecurringDaily{
		At:   o.At,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringFiveMinuteXml) MarshalFromObject(s TypeImeiRecurringFiveMinute) {
	o.Misc = s.Misc
}

func (o typeImeiRecurringFiveMinuteXml) UnmarshalToObject() (*TypeImeiRecurringFiveMinute, error) {

	result := &TypeImeiRecurringFiveMinute{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringHourlyXml) MarshalFromObject(s TypeImeiRecurringHourly) {
	o.Misc = s.Misc
}

func (o typeImeiRecurringHourlyXml) UnmarshalToObject() (*TypeImeiRecurringHourly, error) {

	result := &TypeImeiRecurringHourly{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringMonthlyXml) MarshalFromObject(s TypeImeiRecurringMonthly) {
	o.At = s.At
	o.DayOfMonth = s.DayOfMonth
	o.Misc = s.Misc
}

func (o typeImeiRecurringMonthlyXml) UnmarshalToObject() (*TypeImeiRecurringMonthly, error) {

	result := &TypeImeiRecurringMonthly{
		At:         o.At,
		DayOfMonth: o.DayOfMonth,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeImeiRecurringWeeklyXml) MarshalFromObject(s TypeImeiRecurringWeekly) {
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o typeImeiRecurringWeeklyXml) UnmarshalToObject() (*TypeImeiRecurringWeekly, error) {

	result := &TypeImeiRecurringWeekly{
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *typeImsiXml) MarshalFromObject(s TypeImsi) {
	if s.Auth != nil {
		var obj typeImsiAuthXml
		obj.MarshalFromObject(*s.Auth)
		o.Auth = &obj
	}
	o.CertificateProfile = s.CertificateProfile
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	if s.Recurring != nil {
		var obj typeImsiRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typeImsiXml) UnmarshalToObject() (*TypeImsi, error) {
	var authVal *TypeImsiAuth
	if o.Auth != nil {
		obj, err := o.Auth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authVal = obj
	}
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}
	var recurringVal *TypeImsiRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &TypeImsi{
		Auth:               authVal,
		CertificateProfile: o.CertificateProfile,
		Description:        o.Description,
		ExceptionList:      exceptionListVal,
		Recurring:          recurringVal,
		Url:                o.Url,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *typeImsiAuthXml) MarshalFromObject(s TypeImsiAuth) {
	o.Password = s.Password
	o.Username = s.Username
	o.Misc = s.Misc
}

func (o typeImsiAuthXml) UnmarshalToObject() (*TypeImsiAuth, error) {

	result := &TypeImsiAuth{
		Password: o.Password,
		Username: o.Username,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringXml) MarshalFromObject(s TypeImsiRecurring) {
	if s.Daily != nil {
		var obj typeImsiRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.FiveMinute != nil {
		var obj typeImsiRecurringFiveMinuteXml
		obj.MarshalFromObject(*s.FiveMinute)
		o.FiveMinute = &obj
	}
	if s.Hourly != nil {
		var obj typeImsiRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.Monthly != nil {
		var obj typeImsiRecurringMonthlyXml
		obj.MarshalFromObject(*s.Monthly)
		o.Monthly = &obj
	}
	if s.Weekly != nil {
		var obj typeImsiRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o typeImsiRecurringXml) UnmarshalToObject() (*TypeImsiRecurring, error) {
	var dailyVal *TypeImsiRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var fiveMinuteVal *TypeImsiRecurringFiveMinute
	if o.FiveMinute != nil {
		obj, err := o.FiveMinute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fiveMinuteVal = obj
	}
	var hourlyVal *TypeImsiRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var monthlyVal *TypeImsiRecurringMonthly
	if o.Monthly != nil {
		obj, err := o.Monthly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monthlyVal = obj
	}
	var weeklyVal *TypeImsiRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &TypeImsiRecurring{
		Daily:      dailyVal,
		FiveMinute: fiveMinuteVal,
		Hourly:     hourlyVal,
		Monthly:    monthlyVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringDailyXml) MarshalFromObject(s TypeImsiRecurringDaily) {
	o.At = s.At
	o.Misc = s.Misc
}

func (o typeImsiRecurringDailyXml) UnmarshalToObject() (*TypeImsiRecurringDaily, error) {

	result := &TypeImsiRecurringDaily{
		At:   o.At,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringFiveMinuteXml) MarshalFromObject(s TypeImsiRecurringFiveMinute) {
	o.Misc = s.Misc
}

func (o typeImsiRecurringFiveMinuteXml) UnmarshalToObject() (*TypeImsiRecurringFiveMinute, error) {

	result := &TypeImsiRecurringFiveMinute{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringHourlyXml) MarshalFromObject(s TypeImsiRecurringHourly) {
	o.Misc = s.Misc
}

func (o typeImsiRecurringHourlyXml) UnmarshalToObject() (*TypeImsiRecurringHourly, error) {

	result := &TypeImsiRecurringHourly{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringMonthlyXml) MarshalFromObject(s TypeImsiRecurringMonthly) {
	o.At = s.At
	o.DayOfMonth = s.DayOfMonth
	o.Misc = s.Misc
}

func (o typeImsiRecurringMonthlyXml) UnmarshalToObject() (*TypeImsiRecurringMonthly, error) {

	result := &TypeImsiRecurringMonthly{
		At:         o.At,
		DayOfMonth: o.DayOfMonth,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeImsiRecurringWeeklyXml) MarshalFromObject(s TypeImsiRecurringWeekly) {
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o typeImsiRecurringWeeklyXml) UnmarshalToObject() (*TypeImsiRecurringWeekly, error) {

	result := &TypeImsiRecurringWeekly{
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *typeIpXml) MarshalFromObject(s TypeIp) {
	if s.Auth != nil {
		var obj typeIpAuthXml
		obj.MarshalFromObject(*s.Auth)
		o.Auth = &obj
	}
	o.CertificateProfile = s.CertificateProfile
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	if s.Recurring != nil {
		var obj typeIpRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typeIpXml) UnmarshalToObject() (*TypeIp, error) {
	var authVal *TypeIpAuth
	if o.Auth != nil {
		obj, err := o.Auth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authVal = obj
	}
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}
	var recurringVal *TypeIpRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &TypeIp{
		Auth:               authVal,
		CertificateProfile: o.CertificateProfile,
		Description:        o.Description,
		ExceptionList:      exceptionListVal,
		Recurring:          recurringVal,
		Url:                o.Url,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *typeIpAuthXml) MarshalFromObject(s TypeIpAuth) {
	o.Password = s.Password
	o.Username = s.Username
	o.Misc = s.Misc
}

func (o typeIpAuthXml) UnmarshalToObject() (*TypeIpAuth, error) {

	result := &TypeIpAuth{
		Password: o.Password,
		Username: o.Username,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringXml) MarshalFromObject(s TypeIpRecurring) {
	if s.Daily != nil {
		var obj typeIpRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.FiveMinute != nil {
		var obj typeIpRecurringFiveMinuteXml
		obj.MarshalFromObject(*s.FiveMinute)
		o.FiveMinute = &obj
	}
	if s.Hourly != nil {
		var obj typeIpRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.Monthly != nil {
		var obj typeIpRecurringMonthlyXml
		obj.MarshalFromObject(*s.Monthly)
		o.Monthly = &obj
	}
	if s.Weekly != nil {
		var obj typeIpRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o typeIpRecurringXml) UnmarshalToObject() (*TypeIpRecurring, error) {
	var dailyVal *TypeIpRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var fiveMinuteVal *TypeIpRecurringFiveMinute
	if o.FiveMinute != nil {
		obj, err := o.FiveMinute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fiveMinuteVal = obj
	}
	var hourlyVal *TypeIpRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var monthlyVal *TypeIpRecurringMonthly
	if o.Monthly != nil {
		obj, err := o.Monthly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monthlyVal = obj
	}
	var weeklyVal *TypeIpRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &TypeIpRecurring{
		Daily:      dailyVal,
		FiveMinute: fiveMinuteVal,
		Hourly:     hourlyVal,
		Monthly:    monthlyVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringDailyXml) MarshalFromObject(s TypeIpRecurringDaily) {
	o.At = s.At
	o.Misc = s.Misc
}

func (o typeIpRecurringDailyXml) UnmarshalToObject() (*TypeIpRecurringDaily, error) {

	result := &TypeIpRecurringDaily{
		At:   o.At,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringFiveMinuteXml) MarshalFromObject(s TypeIpRecurringFiveMinute) {
	o.Misc = s.Misc
}

func (o typeIpRecurringFiveMinuteXml) UnmarshalToObject() (*TypeIpRecurringFiveMinute, error) {

	result := &TypeIpRecurringFiveMinute{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringHourlyXml) MarshalFromObject(s TypeIpRecurringHourly) {
	o.Misc = s.Misc
}

func (o typeIpRecurringHourlyXml) UnmarshalToObject() (*TypeIpRecurringHourly, error) {

	result := &TypeIpRecurringHourly{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringMonthlyXml) MarshalFromObject(s TypeIpRecurringMonthly) {
	o.At = s.At
	o.DayOfMonth = s.DayOfMonth
	o.Misc = s.Misc
}

func (o typeIpRecurringMonthlyXml) UnmarshalToObject() (*TypeIpRecurringMonthly, error) {

	result := &TypeIpRecurringMonthly{
		At:         o.At,
		DayOfMonth: o.DayOfMonth,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeIpRecurringWeeklyXml) MarshalFromObject(s TypeIpRecurringWeekly) {
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o typeIpRecurringWeeklyXml) UnmarshalToObject() (*TypeIpRecurringWeekly, error) {

	result := &TypeIpRecurringWeekly{
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *typePredefinedIpXml) MarshalFromObject(s TypePredefinedIp) {
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typePredefinedIpXml) UnmarshalToObject() (*TypePredefinedIp, error) {
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}

	result := &TypePredefinedIp{
		Description:   o.Description,
		ExceptionList: exceptionListVal,
		Url:           o.Url,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *typePredefinedUrlXml) MarshalFromObject(s TypePredefinedUrl) {
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typePredefinedUrlXml) UnmarshalToObject() (*TypePredefinedUrl, error) {
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}

	result := &TypePredefinedUrl{
		Description:   o.Description,
		ExceptionList: exceptionListVal,
		Url:           o.Url,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *typeUrlXml) MarshalFromObject(s TypeUrl) {
	if s.Auth != nil {
		var obj typeUrlAuthXml
		obj.MarshalFromObject(*s.Auth)
		o.Auth = &obj
	}
	o.CertificateProfile = s.CertificateProfile
	o.Description = s.Description
	if s.ExceptionList != nil {
		o.ExceptionList = util.StrToMem(s.ExceptionList)
	}
	if s.Recurring != nil {
		var obj typeUrlRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Url = s.Url
	o.Misc = s.Misc
}

func (o typeUrlXml) UnmarshalToObject() (*TypeUrl, error) {
	var authVal *TypeUrlAuth
	if o.Auth != nil {
		obj, err := o.Auth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authVal = obj
	}
	var exceptionListVal []string
	if o.ExceptionList != nil {
		exceptionListVal = util.MemToStr(o.ExceptionList)
	}
	var recurringVal *TypeUrlRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &TypeUrl{
		Auth:               authVal,
		CertificateProfile: o.CertificateProfile,
		Description:        o.Description,
		ExceptionList:      exceptionListVal,
		Recurring:          recurringVal,
		Url:                o.Url,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *typeUrlAuthXml) MarshalFromObject(s TypeUrlAuth) {
	o.Password = s.Password
	o.Username = s.Username
	o.Misc = s.Misc
}

func (o typeUrlAuthXml) UnmarshalToObject() (*TypeUrlAuth, error) {

	result := &TypeUrlAuth{
		Password: o.Password,
		Username: o.Username,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringXml) MarshalFromObject(s TypeUrlRecurring) {
	if s.Daily != nil {
		var obj typeUrlRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.FiveMinute != nil {
		var obj typeUrlRecurringFiveMinuteXml
		obj.MarshalFromObject(*s.FiveMinute)
		o.FiveMinute = &obj
	}
	if s.Hourly != nil {
		var obj typeUrlRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.Monthly != nil {
		var obj typeUrlRecurringMonthlyXml
		obj.MarshalFromObject(*s.Monthly)
		o.Monthly = &obj
	}
	if s.Weekly != nil {
		var obj typeUrlRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o typeUrlRecurringXml) UnmarshalToObject() (*TypeUrlRecurring, error) {
	var dailyVal *TypeUrlRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var fiveMinuteVal *TypeUrlRecurringFiveMinute
	if o.FiveMinute != nil {
		obj, err := o.FiveMinute.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		fiveMinuteVal = obj
	}
	var hourlyVal *TypeUrlRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var monthlyVal *TypeUrlRecurringMonthly
	if o.Monthly != nil {
		obj, err := o.Monthly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monthlyVal = obj
	}
	var weeklyVal *TypeUrlRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &TypeUrlRecurring{
		Daily:      dailyVal,
		FiveMinute: fiveMinuteVal,
		Hourly:     hourlyVal,
		Monthly:    monthlyVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringDailyXml) MarshalFromObject(s TypeUrlRecurringDaily) {
	o.At = s.At
	o.Misc = s.Misc
}

func (o typeUrlRecurringDailyXml) UnmarshalToObject() (*TypeUrlRecurringDaily, error) {

	result := &TypeUrlRecurringDaily{
		At:   o.At,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringFiveMinuteXml) MarshalFromObject(s TypeUrlRecurringFiveMinute) {
	o.Misc = s.Misc
}

func (o typeUrlRecurringFiveMinuteXml) UnmarshalToObject() (*TypeUrlRecurringFiveMinute, error) {

	result := &TypeUrlRecurringFiveMinute{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringHourlyXml) MarshalFromObject(s TypeUrlRecurringHourly) {
	o.Misc = s.Misc
}

func (o typeUrlRecurringHourlyXml) UnmarshalToObject() (*TypeUrlRecurringHourly, error) {

	result := &TypeUrlRecurringHourly{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringMonthlyXml) MarshalFromObject(s TypeUrlRecurringMonthly) {
	o.At = s.At
	o.DayOfMonth = s.DayOfMonth
	o.Misc = s.Misc
}

func (o typeUrlRecurringMonthlyXml) UnmarshalToObject() (*TypeUrlRecurringMonthly, error) {

	result := &TypeUrlRecurringMonthly{
		At:         o.At,
		DayOfMonth: o.DayOfMonth,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *typeUrlRecurringWeeklyXml) MarshalFromObject(s TypeUrlRecurringWeekly) {
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o typeUrlRecurringWeeklyXml) UnmarshalToObject() (*TypeUrlRecurringWeekly, error) {

	result := &TypeUrlRecurringWeekly{
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "type" || v == "Type" {
		return e.Type, nil
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
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !o.Type.matches(other.Type) {
		return false
	}

	return true
}

func (o *Type) matches(other *Type) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Domain.matches(other.Domain) {
		return false
	}
	if !o.Imei.matches(other.Imei) {
		return false
	}
	if !o.Imsi.matches(other.Imsi) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}
	if !o.PredefinedIp.matches(other.PredefinedIp) {
		return false
	}
	if !o.PredefinedUrl.matches(other.PredefinedUrl) {
		return false
	}
	if !o.Url.matches(other.Url) {
		return false
	}

	return true
}

func (o *TypeDomain) matches(other *TypeDomain) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Auth.matches(other.Auth) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !util.BoolsMatch(o.ExpandDomain, other.ExpandDomain) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeDomainAuth) matches(other *TypeDomainAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *TypeDomainRecurring) matches(other *TypeDomainRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.FiveMinute.matches(other.FiveMinute) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.Monthly.matches(other.Monthly) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *TypeDomainRecurringDaily) matches(other *TypeDomainRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *TypeDomainRecurringFiveMinute) matches(other *TypeDomainRecurringFiveMinute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeDomainRecurringHourly) matches(other *TypeDomainRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeDomainRecurringMonthly) matches(other *TypeDomainRecurringMonthly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.Ints64Match(o.DayOfMonth, other.DayOfMonth) {
		return false
	}

	return true
}

func (o *TypeDomainRecurringWeekly) matches(other *TypeDomainRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
		return false
	}

	return true
}

func (o *TypeImei) matches(other *TypeImei) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Auth.matches(other.Auth) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeImeiAuth) matches(other *TypeImeiAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *TypeImeiRecurring) matches(other *TypeImeiRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.FiveMinute.matches(other.FiveMinute) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.Monthly.matches(other.Monthly) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *TypeImeiRecurringDaily) matches(other *TypeImeiRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *TypeImeiRecurringFiveMinute) matches(other *TypeImeiRecurringFiveMinute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeImeiRecurringHourly) matches(other *TypeImeiRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeImeiRecurringMonthly) matches(other *TypeImeiRecurringMonthly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.Ints64Match(o.DayOfMonth, other.DayOfMonth) {
		return false
	}

	return true
}

func (o *TypeImeiRecurringWeekly) matches(other *TypeImeiRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
		return false
	}

	return true
}

func (o *TypeImsi) matches(other *TypeImsi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Auth.matches(other.Auth) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeImsiAuth) matches(other *TypeImsiAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *TypeImsiRecurring) matches(other *TypeImsiRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.FiveMinute.matches(other.FiveMinute) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.Monthly.matches(other.Monthly) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *TypeImsiRecurringDaily) matches(other *TypeImsiRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *TypeImsiRecurringFiveMinute) matches(other *TypeImsiRecurringFiveMinute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeImsiRecurringHourly) matches(other *TypeImsiRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeImsiRecurringMonthly) matches(other *TypeImsiRecurringMonthly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.Ints64Match(o.DayOfMonth, other.DayOfMonth) {
		return false
	}

	return true
}

func (o *TypeImsiRecurringWeekly) matches(other *TypeImsiRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
		return false
	}

	return true
}

func (o *TypeIp) matches(other *TypeIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Auth.matches(other.Auth) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeIpAuth) matches(other *TypeIpAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *TypeIpRecurring) matches(other *TypeIpRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.FiveMinute.matches(other.FiveMinute) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.Monthly.matches(other.Monthly) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *TypeIpRecurringDaily) matches(other *TypeIpRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *TypeIpRecurringFiveMinute) matches(other *TypeIpRecurringFiveMinute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeIpRecurringHourly) matches(other *TypeIpRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeIpRecurringMonthly) matches(other *TypeIpRecurringMonthly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.Ints64Match(o.DayOfMonth, other.DayOfMonth) {
		return false
	}

	return true
}

func (o *TypeIpRecurringWeekly) matches(other *TypeIpRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
		return false
	}

	return true
}

func (o *TypePredefinedIp) matches(other *TypePredefinedIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypePredefinedUrl) matches(other *TypePredefinedUrl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeUrl) matches(other *TypeUrl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Auth.matches(other.Auth) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.OrderedListsMatch[string](o.ExceptionList, other.ExceptionList) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}

	return true
}

func (o *TypeUrlAuth) matches(other *TypeUrlAuth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Password, other.Password) {
		return false
	}
	if !util.StringsMatch(o.Username, other.Username) {
		return false
	}

	return true
}

func (o *TypeUrlRecurring) matches(other *TypeUrlRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.FiveMinute.matches(other.FiveMinute) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.Monthly.matches(other.Monthly) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *TypeUrlRecurringDaily) matches(other *TypeUrlRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *TypeUrlRecurringFiveMinute) matches(other *TypeUrlRecurringFiveMinute) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeUrlRecurringHourly) matches(other *TypeUrlRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *TypeUrlRecurringMonthly) matches(other *TypeUrlRecurringMonthly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.Ints64Match(o.DayOfMonth, other.DayOfMonth) {
		return false
	}

	return true
}

func (o *TypeUrlRecurringWeekly) matches(other *TypeUrlRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
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
