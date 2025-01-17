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
	Suffix = []string{"address"}
)

type Entry struct {
	Name            string
	DisableOverride *string
	Type            *Type

	Misc map[string][]generic.Xml
}

type Type struct {
	Domain        *TypeDomain
	Imei          *TypeImei
	Imsi          *TypeImsi
	Ip            *TypeIp
	PredefinedIp  *TypePredefinedIp
	PredefinedUrl *TypePredefinedUrl
	Url           *TypeUrl
}
type TypeDomain struct {
	Auth               *TypeDomainAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	ExpandDomain       *bool
	Recurring          *TypeDomainRecurring
	Url                *string
}
type TypeDomainAuth struct {
	Password *string
	Username *string
}
type TypeDomainRecurring struct {
	Daily      *TypeDomainRecurringDaily
	FiveMinute *TypeDomainRecurringFiveMinute
	Hourly     *TypeDomainRecurringHourly
	Monthly    *TypeDomainRecurringMonthly
	Weekly     *TypeDomainRecurringWeekly
}
type TypeDomainRecurringDaily struct {
	At *string
}
type TypeDomainRecurringFiveMinute struct {
}
type TypeDomainRecurringHourly struct {
}
type TypeDomainRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
}
type TypeDomainRecurringWeekly struct {
	At        *string
	DayOfWeek *string
}
type TypeImei struct {
	Auth               *TypeImeiAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeImeiRecurring
	Url                *string
}
type TypeImeiAuth struct {
	Password *string
	Username *string
}
type TypeImeiRecurring struct {
	Daily      *TypeImeiRecurringDaily
	FiveMinute *TypeImeiRecurringFiveMinute
	Hourly     *TypeImeiRecurringHourly
	Monthly    *TypeImeiRecurringMonthly
	Weekly     *TypeImeiRecurringWeekly
}
type TypeImeiRecurringDaily struct {
	At *string
}
type TypeImeiRecurringFiveMinute struct {
}
type TypeImeiRecurringHourly struct {
}
type TypeImeiRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
}
type TypeImeiRecurringWeekly struct {
	At        *string
	DayOfWeek *string
}
type TypeImsi struct {
	Auth               *TypeImsiAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeImsiRecurring
	Url                *string
}
type TypeImsiAuth struct {
	Password *string
	Username *string
}
type TypeImsiRecurring struct {
	Daily      *TypeImsiRecurringDaily
	FiveMinute *TypeImsiRecurringFiveMinute
	Hourly     *TypeImsiRecurringHourly
	Monthly    *TypeImsiRecurringMonthly
	Weekly     *TypeImsiRecurringWeekly
}
type TypeImsiRecurringDaily struct {
	At *string
}
type TypeImsiRecurringFiveMinute struct {
}
type TypeImsiRecurringHourly struct {
}
type TypeImsiRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
}
type TypeImsiRecurringWeekly struct {
	At        *string
	DayOfWeek *string
}
type TypeIp struct {
	Auth               *TypeIpAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeIpRecurring
	Url                *string
}
type TypeIpAuth struct {
	Password *string
	Username *string
}
type TypeIpRecurring struct {
	Daily      *TypeIpRecurringDaily
	FiveMinute *TypeIpRecurringFiveMinute
	Hourly     *TypeIpRecurringHourly
	Monthly    *TypeIpRecurringMonthly
	Weekly     *TypeIpRecurringWeekly
}
type TypeIpRecurringDaily struct {
	At *string
}
type TypeIpRecurringFiveMinute struct {
}
type TypeIpRecurringHourly struct {
}
type TypeIpRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
}
type TypeIpRecurringWeekly struct {
	At        *string
	DayOfWeek *string
}
type TypePredefinedIp struct {
	Description   *string
	ExceptionList []string
	Url           *string
}
type TypePredefinedUrl struct {
	Description   *string
	ExceptionList []string
	Url           *string
}
type TypeUrl struct {
	Auth               *TypeUrlAuth
	CertificateProfile *string
	Description        *string
	ExceptionList      []string
	Recurring          *TypeUrlRecurring
	Url                *string
}
type TypeUrlAuth struct {
	Password *string
	Username *string
}
type TypeUrlRecurring struct {
	Daily      *TypeUrlRecurringDaily
	FiveMinute *TypeUrlRecurringFiveMinute
	Hourly     *TypeUrlRecurringHourly
	Monthly    *TypeUrlRecurringMonthly
	Weekly     *TypeUrlRecurringWeekly
}
type TypeUrlRecurringDaily struct {
	At *string
}
type TypeUrlRecurringFiveMinute struct {
}
type TypeUrlRecurringHourly struct {
}
type TypeUrlRecurringMonthly struct {
	At         *string
	DayOfMonth *int64
}
type TypeUrlRecurringWeekly struct {
	At        *string
	DayOfWeek *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName         xml.Name `xml:"entry"`
	Name            string   `xml:"name,attr"`
	DisableOverride *string  `xml:"disable-override,omitempty"`
	Type            *TypeXml `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeXml struct {
	Domain        *TypeDomainXml        `xml:"domain,omitempty"`
	Imei          *TypeImeiXml          `xml:"imei,omitempty"`
	Imsi          *TypeImsiXml          `xml:"imsi,omitempty"`
	Ip            *TypeIpXml            `xml:"ip,omitempty"`
	PredefinedIp  *TypePredefinedIpXml  `xml:"predefined-ip,omitempty"`
	PredefinedUrl *TypePredefinedUrlXml `xml:"predefined-url,omitempty"`
	Url           *TypeUrlXml           `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainXml struct {
	Auth               *TypeDomainAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string                 `xml:"certificate-profile,omitempty"`
	Description        *string                 `xml:"description,omitempty"`
	ExceptionList      *util.MemberType        `xml:"exception-list,omitempty"`
	ExpandDomain       *string                 `xml:"expand-domain,omitempty"`
	Recurring          *TypeDomainRecurringXml `xml:"recurring,omitempty"`
	Url                *string                 `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainAuthXml struct {
	Password *string `xml:"password,omitempty"`
	Username *string `xml:"username,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringXml struct {
	Daily      *TypeDomainRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *TypeDomainRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *TypeDomainRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *TypeDomainRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *TypeDomainRecurringWeeklyXml     `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringDailyXml struct {
	At *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringMonthlyXml struct {
	At         *string `xml:"at,omitempty"`
	DayOfMonth *int64  `xml:"day-of-month,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeDomainRecurringWeeklyXml struct {
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiXml struct {
	Auth               *TypeImeiAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string               `xml:"certificate-profile,omitempty"`
	Description        *string               `xml:"description,omitempty"`
	ExceptionList      *util.MemberType      `xml:"exception-list,omitempty"`
	Recurring          *TypeImeiRecurringXml `xml:"recurring,omitempty"`
	Url                *string               `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiAuthXml struct {
	Password *string `xml:"password,omitempty"`
	Username *string `xml:"username,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringXml struct {
	Daily      *TypeImeiRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *TypeImeiRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *TypeImeiRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *TypeImeiRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *TypeImeiRecurringWeeklyXml     `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringDailyXml struct {
	At *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringMonthlyXml struct {
	At         *string `xml:"at,omitempty"`
	DayOfMonth *int64  `xml:"day-of-month,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImeiRecurringWeeklyXml struct {
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiXml struct {
	Auth               *TypeImsiAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string               `xml:"certificate-profile,omitempty"`
	Description        *string               `xml:"description,omitempty"`
	ExceptionList      *util.MemberType      `xml:"exception-list,omitempty"`
	Recurring          *TypeImsiRecurringXml `xml:"recurring,omitempty"`
	Url                *string               `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiAuthXml struct {
	Password *string `xml:"password,omitempty"`
	Username *string `xml:"username,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringXml struct {
	Daily      *TypeImsiRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *TypeImsiRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *TypeImsiRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *TypeImsiRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *TypeImsiRecurringWeeklyXml     `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringDailyXml struct {
	At *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringMonthlyXml struct {
	At         *string `xml:"at,omitempty"`
	DayOfMonth *int64  `xml:"day-of-month,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeImsiRecurringWeeklyXml struct {
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpXml struct {
	Auth               *TypeIpAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string             `xml:"certificate-profile,omitempty"`
	Description        *string             `xml:"description,omitempty"`
	ExceptionList      *util.MemberType    `xml:"exception-list,omitempty"`
	Recurring          *TypeIpRecurringXml `xml:"recurring,omitempty"`
	Url                *string             `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpAuthXml struct {
	Password *string `xml:"password,omitempty"`
	Username *string `xml:"username,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringXml struct {
	Daily      *TypeIpRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *TypeIpRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *TypeIpRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *TypeIpRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *TypeIpRecurringWeeklyXml     `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringDailyXml struct {
	At *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringMonthlyXml struct {
	At         *string `xml:"at,omitempty"`
	DayOfMonth *int64  `xml:"day-of-month,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeIpRecurringWeeklyXml struct {
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypePredefinedIpXml struct {
	Description   *string          `xml:"description,omitempty"`
	ExceptionList *util.MemberType `xml:"exception-list,omitempty"`
	Url           *string          `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypePredefinedUrlXml struct {
	Description   *string          `xml:"description,omitempty"`
	ExceptionList *util.MemberType `xml:"exception-list,omitempty"`
	Url           *string          `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlXml struct {
	Auth               *TypeUrlAuthXml      `xml:"auth,omitempty"`
	CertificateProfile *string              `xml:"certificate-profile,omitempty"`
	Description        *string              `xml:"description,omitempty"`
	ExceptionList      *util.MemberType     `xml:"exception-list,omitempty"`
	Recurring          *TypeUrlRecurringXml `xml:"recurring,omitempty"`
	Url                *string              `xml:"url,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlAuthXml struct {
	Password *string `xml:"password,omitempty"`
	Username *string `xml:"username,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringXml struct {
	Daily      *TypeUrlRecurringDailyXml      `xml:"daily,omitempty"`
	FiveMinute *TypeUrlRecurringFiveMinuteXml `xml:"five-minute,omitempty"`
	Hourly     *TypeUrlRecurringHourlyXml     `xml:"hourly,omitempty"`
	Monthly    *TypeUrlRecurringMonthlyXml    `xml:"monthly,omitempty"`
	Weekly     *TypeUrlRecurringWeeklyXml     `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringDailyXml struct {
	At *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringFiveMinuteXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringHourlyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringMonthlyXml struct {
	At         *string `xml:"at,omitempty"`
	DayOfMonth *int64  `xml:"day-of-month,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TypeUrlRecurringWeeklyXml struct {
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.DisableOverride = o.DisableOverride
	var nestedType *TypeXml
	if o.Type != nil {
		nestedType = &TypeXml{}
		if _, ok := o.Misc["Type"]; ok {
			nestedType.Misc = o.Misc["Type"]
		}
		if o.Type.Domain != nil {
			nestedType.Domain = &TypeDomainXml{}
			if _, ok := o.Misc["TypeDomain"]; ok {
				nestedType.Domain.Misc = o.Misc["TypeDomain"]
			}
			if o.Type.Domain.Auth != nil {
				nestedType.Domain.Auth = &TypeDomainAuthXml{}
				if _, ok := o.Misc["TypeDomainAuth"]; ok {
					nestedType.Domain.Auth.Misc = o.Misc["TypeDomainAuth"]
				}
				if o.Type.Domain.Auth.Password != nil {
					nestedType.Domain.Auth.Password = o.Type.Domain.Auth.Password
				}
				if o.Type.Domain.Auth.Username != nil {
					nestedType.Domain.Auth.Username = o.Type.Domain.Auth.Username
				}
			}
			if o.Type.Domain.CertificateProfile != nil {
				nestedType.Domain.CertificateProfile = o.Type.Domain.CertificateProfile
			}
			if o.Type.Domain.Description != nil {
				nestedType.Domain.Description = o.Type.Domain.Description
			}
			if o.Type.Domain.ExceptionList != nil {
				nestedType.Domain.ExceptionList = util.StrToMem(o.Type.Domain.ExceptionList)
			}
			if o.Type.Domain.ExpandDomain != nil {
				nestedType.Domain.ExpandDomain = util.YesNo(o.Type.Domain.ExpandDomain, nil)
			}
			if o.Type.Domain.Recurring != nil {
				nestedType.Domain.Recurring = &TypeDomainRecurringXml{}
				if _, ok := o.Misc["TypeDomainRecurring"]; ok {
					nestedType.Domain.Recurring.Misc = o.Misc["TypeDomainRecurring"]
				}
				if o.Type.Domain.Recurring.Daily != nil {
					nestedType.Domain.Recurring.Daily = &TypeDomainRecurringDailyXml{}
					if _, ok := o.Misc["TypeDomainRecurringDaily"]; ok {
						nestedType.Domain.Recurring.Daily.Misc = o.Misc["TypeDomainRecurringDaily"]
					}
					if o.Type.Domain.Recurring.Daily.At != nil {
						nestedType.Domain.Recurring.Daily.At = o.Type.Domain.Recurring.Daily.At
					}
				}
				if o.Type.Domain.Recurring.FiveMinute != nil {
					nestedType.Domain.Recurring.FiveMinute = &TypeDomainRecurringFiveMinuteXml{}
					if _, ok := o.Misc["TypeDomainRecurringFiveMinute"]; ok {
						nestedType.Domain.Recurring.FiveMinute.Misc = o.Misc["TypeDomainRecurringFiveMinute"]
					}
				}
				if o.Type.Domain.Recurring.Hourly != nil {
					nestedType.Domain.Recurring.Hourly = &TypeDomainRecurringHourlyXml{}
					if _, ok := o.Misc["TypeDomainRecurringHourly"]; ok {
						nestedType.Domain.Recurring.Hourly.Misc = o.Misc["TypeDomainRecurringHourly"]
					}
				}
				if o.Type.Domain.Recurring.Monthly != nil {
					nestedType.Domain.Recurring.Monthly = &TypeDomainRecurringMonthlyXml{}
					if _, ok := o.Misc["TypeDomainRecurringMonthly"]; ok {
						nestedType.Domain.Recurring.Monthly.Misc = o.Misc["TypeDomainRecurringMonthly"]
					}
					if o.Type.Domain.Recurring.Monthly.At != nil {
						nestedType.Domain.Recurring.Monthly.At = o.Type.Domain.Recurring.Monthly.At
					}
					if o.Type.Domain.Recurring.Monthly.DayOfMonth != nil {
						nestedType.Domain.Recurring.Monthly.DayOfMonth = o.Type.Domain.Recurring.Monthly.DayOfMonth
					}
				}
				if o.Type.Domain.Recurring.Weekly != nil {
					nestedType.Domain.Recurring.Weekly = &TypeDomainRecurringWeeklyXml{}
					if _, ok := o.Misc["TypeDomainRecurringWeekly"]; ok {
						nestedType.Domain.Recurring.Weekly.Misc = o.Misc["TypeDomainRecurringWeekly"]
					}
					if o.Type.Domain.Recurring.Weekly.At != nil {
						nestedType.Domain.Recurring.Weekly.At = o.Type.Domain.Recurring.Weekly.At
					}
					if o.Type.Domain.Recurring.Weekly.DayOfWeek != nil {
						nestedType.Domain.Recurring.Weekly.DayOfWeek = o.Type.Domain.Recurring.Weekly.DayOfWeek
					}
				}
			}
			if o.Type.Domain.Url != nil {
				nestedType.Domain.Url = o.Type.Domain.Url
			}
		}
		if o.Type.Imei != nil {
			nestedType.Imei = &TypeImeiXml{}
			if _, ok := o.Misc["TypeImei"]; ok {
				nestedType.Imei.Misc = o.Misc["TypeImei"]
			}
			if o.Type.Imei.Auth != nil {
				nestedType.Imei.Auth = &TypeImeiAuthXml{}
				if _, ok := o.Misc["TypeImeiAuth"]; ok {
					nestedType.Imei.Auth.Misc = o.Misc["TypeImeiAuth"]
				}
				if o.Type.Imei.Auth.Password != nil {
					nestedType.Imei.Auth.Password = o.Type.Imei.Auth.Password
				}
				if o.Type.Imei.Auth.Username != nil {
					nestedType.Imei.Auth.Username = o.Type.Imei.Auth.Username
				}
			}
			if o.Type.Imei.CertificateProfile != nil {
				nestedType.Imei.CertificateProfile = o.Type.Imei.CertificateProfile
			}
			if o.Type.Imei.Description != nil {
				nestedType.Imei.Description = o.Type.Imei.Description
			}
			if o.Type.Imei.ExceptionList != nil {
				nestedType.Imei.ExceptionList = util.StrToMem(o.Type.Imei.ExceptionList)
			}
			if o.Type.Imei.Recurring != nil {
				nestedType.Imei.Recurring = &TypeImeiRecurringXml{}
				if _, ok := o.Misc["TypeImeiRecurring"]; ok {
					nestedType.Imei.Recurring.Misc = o.Misc["TypeImeiRecurring"]
				}
				if o.Type.Imei.Recurring.Daily != nil {
					nestedType.Imei.Recurring.Daily = &TypeImeiRecurringDailyXml{}
					if _, ok := o.Misc["TypeImeiRecurringDaily"]; ok {
						nestedType.Imei.Recurring.Daily.Misc = o.Misc["TypeImeiRecurringDaily"]
					}
					if o.Type.Imei.Recurring.Daily.At != nil {
						nestedType.Imei.Recurring.Daily.At = o.Type.Imei.Recurring.Daily.At
					}
				}
				if o.Type.Imei.Recurring.FiveMinute != nil {
					nestedType.Imei.Recurring.FiveMinute = &TypeImeiRecurringFiveMinuteXml{}
					if _, ok := o.Misc["TypeImeiRecurringFiveMinute"]; ok {
						nestedType.Imei.Recurring.FiveMinute.Misc = o.Misc["TypeImeiRecurringFiveMinute"]
					}
				}
				if o.Type.Imei.Recurring.Hourly != nil {
					nestedType.Imei.Recurring.Hourly = &TypeImeiRecurringHourlyXml{}
					if _, ok := o.Misc["TypeImeiRecurringHourly"]; ok {
						nestedType.Imei.Recurring.Hourly.Misc = o.Misc["TypeImeiRecurringHourly"]
					}
				}
				if o.Type.Imei.Recurring.Monthly != nil {
					nestedType.Imei.Recurring.Monthly = &TypeImeiRecurringMonthlyXml{}
					if _, ok := o.Misc["TypeImeiRecurringMonthly"]; ok {
						nestedType.Imei.Recurring.Monthly.Misc = o.Misc["TypeImeiRecurringMonthly"]
					}
					if o.Type.Imei.Recurring.Monthly.At != nil {
						nestedType.Imei.Recurring.Monthly.At = o.Type.Imei.Recurring.Monthly.At
					}
					if o.Type.Imei.Recurring.Monthly.DayOfMonth != nil {
						nestedType.Imei.Recurring.Monthly.DayOfMonth = o.Type.Imei.Recurring.Monthly.DayOfMonth
					}
				}
				if o.Type.Imei.Recurring.Weekly != nil {
					nestedType.Imei.Recurring.Weekly = &TypeImeiRecurringWeeklyXml{}
					if _, ok := o.Misc["TypeImeiRecurringWeekly"]; ok {
						nestedType.Imei.Recurring.Weekly.Misc = o.Misc["TypeImeiRecurringWeekly"]
					}
					if o.Type.Imei.Recurring.Weekly.DayOfWeek != nil {
						nestedType.Imei.Recurring.Weekly.DayOfWeek = o.Type.Imei.Recurring.Weekly.DayOfWeek
					}
					if o.Type.Imei.Recurring.Weekly.At != nil {
						nestedType.Imei.Recurring.Weekly.At = o.Type.Imei.Recurring.Weekly.At
					}
				}
			}
			if o.Type.Imei.Url != nil {
				nestedType.Imei.Url = o.Type.Imei.Url
			}
		}
		if o.Type.Imsi != nil {
			nestedType.Imsi = &TypeImsiXml{}
			if _, ok := o.Misc["TypeImsi"]; ok {
				nestedType.Imsi.Misc = o.Misc["TypeImsi"]
			}
			if o.Type.Imsi.Auth != nil {
				nestedType.Imsi.Auth = &TypeImsiAuthXml{}
				if _, ok := o.Misc["TypeImsiAuth"]; ok {
					nestedType.Imsi.Auth.Misc = o.Misc["TypeImsiAuth"]
				}
				if o.Type.Imsi.Auth.Password != nil {
					nestedType.Imsi.Auth.Password = o.Type.Imsi.Auth.Password
				}
				if o.Type.Imsi.Auth.Username != nil {
					nestedType.Imsi.Auth.Username = o.Type.Imsi.Auth.Username
				}
			}
			if o.Type.Imsi.CertificateProfile != nil {
				nestedType.Imsi.CertificateProfile = o.Type.Imsi.CertificateProfile
			}
			if o.Type.Imsi.Description != nil {
				nestedType.Imsi.Description = o.Type.Imsi.Description
			}
			if o.Type.Imsi.ExceptionList != nil {
				nestedType.Imsi.ExceptionList = util.StrToMem(o.Type.Imsi.ExceptionList)
			}
			if o.Type.Imsi.Recurring != nil {
				nestedType.Imsi.Recurring = &TypeImsiRecurringXml{}
				if _, ok := o.Misc["TypeImsiRecurring"]; ok {
					nestedType.Imsi.Recurring.Misc = o.Misc["TypeImsiRecurring"]
				}
				if o.Type.Imsi.Recurring.Daily != nil {
					nestedType.Imsi.Recurring.Daily = &TypeImsiRecurringDailyXml{}
					if _, ok := o.Misc["TypeImsiRecurringDaily"]; ok {
						nestedType.Imsi.Recurring.Daily.Misc = o.Misc["TypeImsiRecurringDaily"]
					}
					if o.Type.Imsi.Recurring.Daily.At != nil {
						nestedType.Imsi.Recurring.Daily.At = o.Type.Imsi.Recurring.Daily.At
					}
				}
				if o.Type.Imsi.Recurring.FiveMinute != nil {
					nestedType.Imsi.Recurring.FiveMinute = &TypeImsiRecurringFiveMinuteXml{}
					if _, ok := o.Misc["TypeImsiRecurringFiveMinute"]; ok {
						nestedType.Imsi.Recurring.FiveMinute.Misc = o.Misc["TypeImsiRecurringFiveMinute"]
					}
				}
				if o.Type.Imsi.Recurring.Hourly != nil {
					nestedType.Imsi.Recurring.Hourly = &TypeImsiRecurringHourlyXml{}
					if _, ok := o.Misc["TypeImsiRecurringHourly"]; ok {
						nestedType.Imsi.Recurring.Hourly.Misc = o.Misc["TypeImsiRecurringHourly"]
					}
				}
				if o.Type.Imsi.Recurring.Monthly != nil {
					nestedType.Imsi.Recurring.Monthly = &TypeImsiRecurringMonthlyXml{}
					if _, ok := o.Misc["TypeImsiRecurringMonthly"]; ok {
						nestedType.Imsi.Recurring.Monthly.Misc = o.Misc["TypeImsiRecurringMonthly"]
					}
					if o.Type.Imsi.Recurring.Monthly.At != nil {
						nestedType.Imsi.Recurring.Monthly.At = o.Type.Imsi.Recurring.Monthly.At
					}
					if o.Type.Imsi.Recurring.Monthly.DayOfMonth != nil {
						nestedType.Imsi.Recurring.Monthly.DayOfMonth = o.Type.Imsi.Recurring.Monthly.DayOfMonth
					}
				}
				if o.Type.Imsi.Recurring.Weekly != nil {
					nestedType.Imsi.Recurring.Weekly = &TypeImsiRecurringWeeklyXml{}
					if _, ok := o.Misc["TypeImsiRecurringWeekly"]; ok {
						nestedType.Imsi.Recurring.Weekly.Misc = o.Misc["TypeImsiRecurringWeekly"]
					}
					if o.Type.Imsi.Recurring.Weekly.At != nil {
						nestedType.Imsi.Recurring.Weekly.At = o.Type.Imsi.Recurring.Weekly.At
					}
					if o.Type.Imsi.Recurring.Weekly.DayOfWeek != nil {
						nestedType.Imsi.Recurring.Weekly.DayOfWeek = o.Type.Imsi.Recurring.Weekly.DayOfWeek
					}
				}
			}
			if o.Type.Imsi.Url != nil {
				nestedType.Imsi.Url = o.Type.Imsi.Url
			}
		}
		if o.Type.Ip != nil {
			nestedType.Ip = &TypeIpXml{}
			if _, ok := o.Misc["TypeIp"]; ok {
				nestedType.Ip.Misc = o.Misc["TypeIp"]
			}
			if o.Type.Ip.CertificateProfile != nil {
				nestedType.Ip.CertificateProfile = o.Type.Ip.CertificateProfile
			}
			if o.Type.Ip.Description != nil {
				nestedType.Ip.Description = o.Type.Ip.Description
			}
			if o.Type.Ip.ExceptionList != nil {
				nestedType.Ip.ExceptionList = util.StrToMem(o.Type.Ip.ExceptionList)
			}
			if o.Type.Ip.Recurring != nil {
				nestedType.Ip.Recurring = &TypeIpRecurringXml{}
				if _, ok := o.Misc["TypeIpRecurring"]; ok {
					nestedType.Ip.Recurring.Misc = o.Misc["TypeIpRecurring"]
				}
				if o.Type.Ip.Recurring.Weekly != nil {
					nestedType.Ip.Recurring.Weekly = &TypeIpRecurringWeeklyXml{}
					if _, ok := o.Misc["TypeIpRecurringWeekly"]; ok {
						nestedType.Ip.Recurring.Weekly.Misc = o.Misc["TypeIpRecurringWeekly"]
					}
					if o.Type.Ip.Recurring.Weekly.At != nil {
						nestedType.Ip.Recurring.Weekly.At = o.Type.Ip.Recurring.Weekly.At
					}
					if o.Type.Ip.Recurring.Weekly.DayOfWeek != nil {
						nestedType.Ip.Recurring.Weekly.DayOfWeek = o.Type.Ip.Recurring.Weekly.DayOfWeek
					}
				}
				if o.Type.Ip.Recurring.Daily != nil {
					nestedType.Ip.Recurring.Daily = &TypeIpRecurringDailyXml{}
					if _, ok := o.Misc["TypeIpRecurringDaily"]; ok {
						nestedType.Ip.Recurring.Daily.Misc = o.Misc["TypeIpRecurringDaily"]
					}
					if o.Type.Ip.Recurring.Daily.At != nil {
						nestedType.Ip.Recurring.Daily.At = o.Type.Ip.Recurring.Daily.At
					}
				}
				if o.Type.Ip.Recurring.FiveMinute != nil {
					nestedType.Ip.Recurring.FiveMinute = &TypeIpRecurringFiveMinuteXml{}
					if _, ok := o.Misc["TypeIpRecurringFiveMinute"]; ok {
						nestedType.Ip.Recurring.FiveMinute.Misc = o.Misc["TypeIpRecurringFiveMinute"]
					}
				}
				if o.Type.Ip.Recurring.Hourly != nil {
					nestedType.Ip.Recurring.Hourly = &TypeIpRecurringHourlyXml{}
					if _, ok := o.Misc["TypeIpRecurringHourly"]; ok {
						nestedType.Ip.Recurring.Hourly.Misc = o.Misc["TypeIpRecurringHourly"]
					}
				}
				if o.Type.Ip.Recurring.Monthly != nil {
					nestedType.Ip.Recurring.Monthly = &TypeIpRecurringMonthlyXml{}
					if _, ok := o.Misc["TypeIpRecurringMonthly"]; ok {
						nestedType.Ip.Recurring.Monthly.Misc = o.Misc["TypeIpRecurringMonthly"]
					}
					if o.Type.Ip.Recurring.Monthly.At != nil {
						nestedType.Ip.Recurring.Monthly.At = o.Type.Ip.Recurring.Monthly.At
					}
					if o.Type.Ip.Recurring.Monthly.DayOfMonth != nil {
						nestedType.Ip.Recurring.Monthly.DayOfMonth = o.Type.Ip.Recurring.Monthly.DayOfMonth
					}
				}
			}
			if o.Type.Ip.Url != nil {
				nestedType.Ip.Url = o.Type.Ip.Url
			}
			if o.Type.Ip.Auth != nil {
				nestedType.Ip.Auth = &TypeIpAuthXml{}
				if _, ok := o.Misc["TypeIpAuth"]; ok {
					nestedType.Ip.Auth.Misc = o.Misc["TypeIpAuth"]
				}
				if o.Type.Ip.Auth.Password != nil {
					nestedType.Ip.Auth.Password = o.Type.Ip.Auth.Password
				}
				if o.Type.Ip.Auth.Username != nil {
					nestedType.Ip.Auth.Username = o.Type.Ip.Auth.Username
				}
			}
		}
		if o.Type.PredefinedIp != nil {
			nestedType.PredefinedIp = &TypePredefinedIpXml{}
			if _, ok := o.Misc["TypePredefinedIp"]; ok {
				nestedType.PredefinedIp.Misc = o.Misc["TypePredefinedIp"]
			}
			if o.Type.PredefinedIp.Description != nil {
				nestedType.PredefinedIp.Description = o.Type.PredefinedIp.Description
			}
			if o.Type.PredefinedIp.ExceptionList != nil {
				nestedType.PredefinedIp.ExceptionList = util.StrToMem(o.Type.PredefinedIp.ExceptionList)
			}
			if o.Type.PredefinedIp.Url != nil {
				nestedType.PredefinedIp.Url = o.Type.PredefinedIp.Url
			}
		}
		if o.Type.PredefinedUrl != nil {
			nestedType.PredefinedUrl = &TypePredefinedUrlXml{}
			if _, ok := o.Misc["TypePredefinedUrl"]; ok {
				nestedType.PredefinedUrl.Misc = o.Misc["TypePredefinedUrl"]
			}
			if o.Type.PredefinedUrl.Description != nil {
				nestedType.PredefinedUrl.Description = o.Type.PredefinedUrl.Description
			}
			if o.Type.PredefinedUrl.ExceptionList != nil {
				nestedType.PredefinedUrl.ExceptionList = util.StrToMem(o.Type.PredefinedUrl.ExceptionList)
			}
			if o.Type.PredefinedUrl.Url != nil {
				nestedType.PredefinedUrl.Url = o.Type.PredefinedUrl.Url
			}
		}
		if o.Type.Url != nil {
			nestedType.Url = &TypeUrlXml{}
			if _, ok := o.Misc["TypeUrl"]; ok {
				nestedType.Url.Misc = o.Misc["TypeUrl"]
			}
			if o.Type.Url.Auth != nil {
				nestedType.Url.Auth = &TypeUrlAuthXml{}
				if _, ok := o.Misc["TypeUrlAuth"]; ok {
					nestedType.Url.Auth.Misc = o.Misc["TypeUrlAuth"]
				}
				if o.Type.Url.Auth.Username != nil {
					nestedType.Url.Auth.Username = o.Type.Url.Auth.Username
				}
				if o.Type.Url.Auth.Password != nil {
					nestedType.Url.Auth.Password = o.Type.Url.Auth.Password
				}
			}
			if o.Type.Url.CertificateProfile != nil {
				nestedType.Url.CertificateProfile = o.Type.Url.CertificateProfile
			}
			if o.Type.Url.Description != nil {
				nestedType.Url.Description = o.Type.Url.Description
			}
			if o.Type.Url.ExceptionList != nil {
				nestedType.Url.ExceptionList = util.StrToMem(o.Type.Url.ExceptionList)
			}
			if o.Type.Url.Recurring != nil {
				nestedType.Url.Recurring = &TypeUrlRecurringXml{}
				if _, ok := o.Misc["TypeUrlRecurring"]; ok {
					nestedType.Url.Recurring.Misc = o.Misc["TypeUrlRecurring"]
				}
				if o.Type.Url.Recurring.FiveMinute != nil {
					nestedType.Url.Recurring.FiveMinute = &TypeUrlRecurringFiveMinuteXml{}
					if _, ok := o.Misc["TypeUrlRecurringFiveMinute"]; ok {
						nestedType.Url.Recurring.FiveMinute.Misc = o.Misc["TypeUrlRecurringFiveMinute"]
					}
				}
				if o.Type.Url.Recurring.Hourly != nil {
					nestedType.Url.Recurring.Hourly = &TypeUrlRecurringHourlyXml{}
					if _, ok := o.Misc["TypeUrlRecurringHourly"]; ok {
						nestedType.Url.Recurring.Hourly.Misc = o.Misc["TypeUrlRecurringHourly"]
					}
				}
				if o.Type.Url.Recurring.Monthly != nil {
					nestedType.Url.Recurring.Monthly = &TypeUrlRecurringMonthlyXml{}
					if _, ok := o.Misc["TypeUrlRecurringMonthly"]; ok {
						nestedType.Url.Recurring.Monthly.Misc = o.Misc["TypeUrlRecurringMonthly"]
					}
					if o.Type.Url.Recurring.Monthly.At != nil {
						nestedType.Url.Recurring.Monthly.At = o.Type.Url.Recurring.Monthly.At
					}
					if o.Type.Url.Recurring.Monthly.DayOfMonth != nil {
						nestedType.Url.Recurring.Monthly.DayOfMonth = o.Type.Url.Recurring.Monthly.DayOfMonth
					}
				}
				if o.Type.Url.Recurring.Weekly != nil {
					nestedType.Url.Recurring.Weekly = &TypeUrlRecurringWeeklyXml{}
					if _, ok := o.Misc["TypeUrlRecurringWeekly"]; ok {
						nestedType.Url.Recurring.Weekly.Misc = o.Misc["TypeUrlRecurringWeekly"]
					}
					if o.Type.Url.Recurring.Weekly.DayOfWeek != nil {
						nestedType.Url.Recurring.Weekly.DayOfWeek = o.Type.Url.Recurring.Weekly.DayOfWeek
					}
					if o.Type.Url.Recurring.Weekly.At != nil {
						nestedType.Url.Recurring.Weekly.At = o.Type.Url.Recurring.Weekly.At
					}
				}
				if o.Type.Url.Recurring.Daily != nil {
					nestedType.Url.Recurring.Daily = &TypeUrlRecurringDailyXml{}
					if _, ok := o.Misc["TypeUrlRecurringDaily"]; ok {
						nestedType.Url.Recurring.Daily.Misc = o.Misc["TypeUrlRecurringDaily"]
					}
					if o.Type.Url.Recurring.Daily.At != nil {
						nestedType.Url.Recurring.Daily.At = o.Type.Url.Recurring.Daily.At
					}
				}
			}
			if o.Type.Url.Url != nil {
				nestedType.Url.Url = o.Type.Url.Url
			}
		}
	}
	entry.Type = nestedType

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
		entry.DisableOverride = o.DisableOverride
		var nestedType *Type
		if o.Type != nil {
			nestedType = &Type{}
			if o.Type.Misc != nil {
				entry.Misc["Type"] = o.Type.Misc
			}
			if o.Type.Domain != nil {
				nestedType.Domain = &TypeDomain{}
				if o.Type.Domain.Misc != nil {
					entry.Misc["TypeDomain"] = o.Type.Domain.Misc
				}
				if o.Type.Domain.ExpandDomain != nil {
					nestedType.Domain.ExpandDomain = util.AsBool(o.Type.Domain.ExpandDomain, nil)
				}
				if o.Type.Domain.Recurring != nil {
					nestedType.Domain.Recurring = &TypeDomainRecurring{}
					if o.Type.Domain.Recurring.Misc != nil {
						entry.Misc["TypeDomainRecurring"] = o.Type.Domain.Recurring.Misc
					}
					if o.Type.Domain.Recurring.Daily != nil {
						nestedType.Domain.Recurring.Daily = &TypeDomainRecurringDaily{}
						if o.Type.Domain.Recurring.Daily.Misc != nil {
							entry.Misc["TypeDomainRecurringDaily"] = o.Type.Domain.Recurring.Daily.Misc
						}
						if o.Type.Domain.Recurring.Daily.At != nil {
							nestedType.Domain.Recurring.Daily.At = o.Type.Domain.Recurring.Daily.At
						}
					}
					if o.Type.Domain.Recurring.FiveMinute != nil {
						nestedType.Domain.Recurring.FiveMinute = &TypeDomainRecurringFiveMinute{}
						if o.Type.Domain.Recurring.FiveMinute.Misc != nil {
							entry.Misc["TypeDomainRecurringFiveMinute"] = o.Type.Domain.Recurring.FiveMinute.Misc
						}
					}
					if o.Type.Domain.Recurring.Hourly != nil {
						nestedType.Domain.Recurring.Hourly = &TypeDomainRecurringHourly{}
						if o.Type.Domain.Recurring.Hourly.Misc != nil {
							entry.Misc["TypeDomainRecurringHourly"] = o.Type.Domain.Recurring.Hourly.Misc
						}
					}
					if o.Type.Domain.Recurring.Monthly != nil {
						nestedType.Domain.Recurring.Monthly = &TypeDomainRecurringMonthly{}
						if o.Type.Domain.Recurring.Monthly.Misc != nil {
							entry.Misc["TypeDomainRecurringMonthly"] = o.Type.Domain.Recurring.Monthly.Misc
						}
						if o.Type.Domain.Recurring.Monthly.At != nil {
							nestedType.Domain.Recurring.Monthly.At = o.Type.Domain.Recurring.Monthly.At
						}
						if o.Type.Domain.Recurring.Monthly.DayOfMonth != nil {
							nestedType.Domain.Recurring.Monthly.DayOfMonth = o.Type.Domain.Recurring.Monthly.DayOfMonth
						}
					}
					if o.Type.Domain.Recurring.Weekly != nil {
						nestedType.Domain.Recurring.Weekly = &TypeDomainRecurringWeekly{}
						if o.Type.Domain.Recurring.Weekly.Misc != nil {
							entry.Misc["TypeDomainRecurringWeekly"] = o.Type.Domain.Recurring.Weekly.Misc
						}
						if o.Type.Domain.Recurring.Weekly.At != nil {
							nestedType.Domain.Recurring.Weekly.At = o.Type.Domain.Recurring.Weekly.At
						}
						if o.Type.Domain.Recurring.Weekly.DayOfWeek != nil {
							nestedType.Domain.Recurring.Weekly.DayOfWeek = o.Type.Domain.Recurring.Weekly.DayOfWeek
						}
					}
				}
				if o.Type.Domain.Url != nil {
					nestedType.Domain.Url = o.Type.Domain.Url
				}
				if o.Type.Domain.Auth != nil {
					nestedType.Domain.Auth = &TypeDomainAuth{}
					if o.Type.Domain.Auth.Misc != nil {
						entry.Misc["TypeDomainAuth"] = o.Type.Domain.Auth.Misc
					}
					if o.Type.Domain.Auth.Password != nil {
						nestedType.Domain.Auth.Password = o.Type.Domain.Auth.Password
					}
					if o.Type.Domain.Auth.Username != nil {
						nestedType.Domain.Auth.Username = o.Type.Domain.Auth.Username
					}
				}
				if o.Type.Domain.CertificateProfile != nil {
					nestedType.Domain.CertificateProfile = o.Type.Domain.CertificateProfile
				}
				if o.Type.Domain.Description != nil {
					nestedType.Domain.Description = o.Type.Domain.Description
				}
				if o.Type.Domain.ExceptionList != nil {
					nestedType.Domain.ExceptionList = util.MemToStr(o.Type.Domain.ExceptionList)
				}
			}
			if o.Type.Imei != nil {
				nestedType.Imei = &TypeImei{}
				if o.Type.Imei.Misc != nil {
					entry.Misc["TypeImei"] = o.Type.Imei.Misc
				}
				if o.Type.Imei.Description != nil {
					nestedType.Imei.Description = o.Type.Imei.Description
				}
				if o.Type.Imei.ExceptionList != nil {
					nestedType.Imei.ExceptionList = util.MemToStr(o.Type.Imei.ExceptionList)
				}
				if o.Type.Imei.Recurring != nil {
					nestedType.Imei.Recurring = &TypeImeiRecurring{}
					if o.Type.Imei.Recurring.Misc != nil {
						entry.Misc["TypeImeiRecurring"] = o.Type.Imei.Recurring.Misc
					}
					if o.Type.Imei.Recurring.Daily != nil {
						nestedType.Imei.Recurring.Daily = &TypeImeiRecurringDaily{}
						if o.Type.Imei.Recurring.Daily.Misc != nil {
							entry.Misc["TypeImeiRecurringDaily"] = o.Type.Imei.Recurring.Daily.Misc
						}
						if o.Type.Imei.Recurring.Daily.At != nil {
							nestedType.Imei.Recurring.Daily.At = o.Type.Imei.Recurring.Daily.At
						}
					}
					if o.Type.Imei.Recurring.FiveMinute != nil {
						nestedType.Imei.Recurring.FiveMinute = &TypeImeiRecurringFiveMinute{}
						if o.Type.Imei.Recurring.FiveMinute.Misc != nil {
							entry.Misc["TypeImeiRecurringFiveMinute"] = o.Type.Imei.Recurring.FiveMinute.Misc
						}
					}
					if o.Type.Imei.Recurring.Hourly != nil {
						nestedType.Imei.Recurring.Hourly = &TypeImeiRecurringHourly{}
						if o.Type.Imei.Recurring.Hourly.Misc != nil {
							entry.Misc["TypeImeiRecurringHourly"] = o.Type.Imei.Recurring.Hourly.Misc
						}
					}
					if o.Type.Imei.Recurring.Monthly != nil {
						nestedType.Imei.Recurring.Monthly = &TypeImeiRecurringMonthly{}
						if o.Type.Imei.Recurring.Monthly.Misc != nil {
							entry.Misc["TypeImeiRecurringMonthly"] = o.Type.Imei.Recurring.Monthly.Misc
						}
						if o.Type.Imei.Recurring.Monthly.DayOfMonth != nil {
							nestedType.Imei.Recurring.Monthly.DayOfMonth = o.Type.Imei.Recurring.Monthly.DayOfMonth
						}
						if o.Type.Imei.Recurring.Monthly.At != nil {
							nestedType.Imei.Recurring.Monthly.At = o.Type.Imei.Recurring.Monthly.At
						}
					}
					if o.Type.Imei.Recurring.Weekly != nil {
						nestedType.Imei.Recurring.Weekly = &TypeImeiRecurringWeekly{}
						if o.Type.Imei.Recurring.Weekly.Misc != nil {
							entry.Misc["TypeImeiRecurringWeekly"] = o.Type.Imei.Recurring.Weekly.Misc
						}
						if o.Type.Imei.Recurring.Weekly.At != nil {
							nestedType.Imei.Recurring.Weekly.At = o.Type.Imei.Recurring.Weekly.At
						}
						if o.Type.Imei.Recurring.Weekly.DayOfWeek != nil {
							nestedType.Imei.Recurring.Weekly.DayOfWeek = o.Type.Imei.Recurring.Weekly.DayOfWeek
						}
					}
				}
				if o.Type.Imei.Url != nil {
					nestedType.Imei.Url = o.Type.Imei.Url
				}
				if o.Type.Imei.Auth != nil {
					nestedType.Imei.Auth = &TypeImeiAuth{}
					if o.Type.Imei.Auth.Misc != nil {
						entry.Misc["TypeImeiAuth"] = o.Type.Imei.Auth.Misc
					}
					if o.Type.Imei.Auth.Password != nil {
						nestedType.Imei.Auth.Password = o.Type.Imei.Auth.Password
					}
					if o.Type.Imei.Auth.Username != nil {
						nestedType.Imei.Auth.Username = o.Type.Imei.Auth.Username
					}
				}
				if o.Type.Imei.CertificateProfile != nil {
					nestedType.Imei.CertificateProfile = o.Type.Imei.CertificateProfile
				}
			}
			if o.Type.Imsi != nil {
				nestedType.Imsi = &TypeImsi{}
				if o.Type.Imsi.Misc != nil {
					entry.Misc["TypeImsi"] = o.Type.Imsi.Misc
				}
				if o.Type.Imsi.Recurring != nil {
					nestedType.Imsi.Recurring = &TypeImsiRecurring{}
					if o.Type.Imsi.Recurring.Misc != nil {
						entry.Misc["TypeImsiRecurring"] = o.Type.Imsi.Recurring.Misc
					}
					if o.Type.Imsi.Recurring.Daily != nil {
						nestedType.Imsi.Recurring.Daily = &TypeImsiRecurringDaily{}
						if o.Type.Imsi.Recurring.Daily.Misc != nil {
							entry.Misc["TypeImsiRecurringDaily"] = o.Type.Imsi.Recurring.Daily.Misc
						}
						if o.Type.Imsi.Recurring.Daily.At != nil {
							nestedType.Imsi.Recurring.Daily.At = o.Type.Imsi.Recurring.Daily.At
						}
					}
					if o.Type.Imsi.Recurring.FiveMinute != nil {
						nestedType.Imsi.Recurring.FiveMinute = &TypeImsiRecurringFiveMinute{}
						if o.Type.Imsi.Recurring.FiveMinute.Misc != nil {
							entry.Misc["TypeImsiRecurringFiveMinute"] = o.Type.Imsi.Recurring.FiveMinute.Misc
						}
					}
					if o.Type.Imsi.Recurring.Hourly != nil {
						nestedType.Imsi.Recurring.Hourly = &TypeImsiRecurringHourly{}
						if o.Type.Imsi.Recurring.Hourly.Misc != nil {
							entry.Misc["TypeImsiRecurringHourly"] = o.Type.Imsi.Recurring.Hourly.Misc
						}
					}
					if o.Type.Imsi.Recurring.Monthly != nil {
						nestedType.Imsi.Recurring.Monthly = &TypeImsiRecurringMonthly{}
						if o.Type.Imsi.Recurring.Monthly.Misc != nil {
							entry.Misc["TypeImsiRecurringMonthly"] = o.Type.Imsi.Recurring.Monthly.Misc
						}
						if o.Type.Imsi.Recurring.Monthly.At != nil {
							nestedType.Imsi.Recurring.Monthly.At = o.Type.Imsi.Recurring.Monthly.At
						}
						if o.Type.Imsi.Recurring.Monthly.DayOfMonth != nil {
							nestedType.Imsi.Recurring.Monthly.DayOfMonth = o.Type.Imsi.Recurring.Monthly.DayOfMonth
						}
					}
					if o.Type.Imsi.Recurring.Weekly != nil {
						nestedType.Imsi.Recurring.Weekly = &TypeImsiRecurringWeekly{}
						if o.Type.Imsi.Recurring.Weekly.Misc != nil {
							entry.Misc["TypeImsiRecurringWeekly"] = o.Type.Imsi.Recurring.Weekly.Misc
						}
						if o.Type.Imsi.Recurring.Weekly.At != nil {
							nestedType.Imsi.Recurring.Weekly.At = o.Type.Imsi.Recurring.Weekly.At
						}
						if o.Type.Imsi.Recurring.Weekly.DayOfWeek != nil {
							nestedType.Imsi.Recurring.Weekly.DayOfWeek = o.Type.Imsi.Recurring.Weekly.DayOfWeek
						}
					}
				}
				if o.Type.Imsi.Url != nil {
					nestedType.Imsi.Url = o.Type.Imsi.Url
				}
				if o.Type.Imsi.Auth != nil {
					nestedType.Imsi.Auth = &TypeImsiAuth{}
					if o.Type.Imsi.Auth.Misc != nil {
						entry.Misc["TypeImsiAuth"] = o.Type.Imsi.Auth.Misc
					}
					if o.Type.Imsi.Auth.Password != nil {
						nestedType.Imsi.Auth.Password = o.Type.Imsi.Auth.Password
					}
					if o.Type.Imsi.Auth.Username != nil {
						nestedType.Imsi.Auth.Username = o.Type.Imsi.Auth.Username
					}
				}
				if o.Type.Imsi.CertificateProfile != nil {
					nestedType.Imsi.CertificateProfile = o.Type.Imsi.CertificateProfile
				}
				if o.Type.Imsi.Description != nil {
					nestedType.Imsi.Description = o.Type.Imsi.Description
				}
				if o.Type.Imsi.ExceptionList != nil {
					nestedType.Imsi.ExceptionList = util.MemToStr(o.Type.Imsi.ExceptionList)
				}
			}
			if o.Type.Ip != nil {
				nestedType.Ip = &TypeIp{}
				if o.Type.Ip.Misc != nil {
					entry.Misc["TypeIp"] = o.Type.Ip.Misc
				}
				if o.Type.Ip.Auth != nil {
					nestedType.Ip.Auth = &TypeIpAuth{}
					if o.Type.Ip.Auth.Misc != nil {
						entry.Misc["TypeIpAuth"] = o.Type.Ip.Auth.Misc
					}
					if o.Type.Ip.Auth.Password != nil {
						nestedType.Ip.Auth.Password = o.Type.Ip.Auth.Password
					}
					if o.Type.Ip.Auth.Username != nil {
						nestedType.Ip.Auth.Username = o.Type.Ip.Auth.Username
					}
				}
				if o.Type.Ip.CertificateProfile != nil {
					nestedType.Ip.CertificateProfile = o.Type.Ip.CertificateProfile
				}
				if o.Type.Ip.Description != nil {
					nestedType.Ip.Description = o.Type.Ip.Description
				}
				if o.Type.Ip.ExceptionList != nil {
					nestedType.Ip.ExceptionList = util.MemToStr(o.Type.Ip.ExceptionList)
				}
				if o.Type.Ip.Recurring != nil {
					nestedType.Ip.Recurring = &TypeIpRecurring{}
					if o.Type.Ip.Recurring.Misc != nil {
						entry.Misc["TypeIpRecurring"] = o.Type.Ip.Recurring.Misc
					}
					if o.Type.Ip.Recurring.Hourly != nil {
						nestedType.Ip.Recurring.Hourly = &TypeIpRecurringHourly{}
						if o.Type.Ip.Recurring.Hourly.Misc != nil {
							entry.Misc["TypeIpRecurringHourly"] = o.Type.Ip.Recurring.Hourly.Misc
						}
					}
					if o.Type.Ip.Recurring.Monthly != nil {
						nestedType.Ip.Recurring.Monthly = &TypeIpRecurringMonthly{}
						if o.Type.Ip.Recurring.Monthly.Misc != nil {
							entry.Misc["TypeIpRecurringMonthly"] = o.Type.Ip.Recurring.Monthly.Misc
						}
						if o.Type.Ip.Recurring.Monthly.At != nil {
							nestedType.Ip.Recurring.Monthly.At = o.Type.Ip.Recurring.Monthly.At
						}
						if o.Type.Ip.Recurring.Monthly.DayOfMonth != nil {
							nestedType.Ip.Recurring.Monthly.DayOfMonth = o.Type.Ip.Recurring.Monthly.DayOfMonth
						}
					}
					if o.Type.Ip.Recurring.Weekly != nil {
						nestedType.Ip.Recurring.Weekly = &TypeIpRecurringWeekly{}
						if o.Type.Ip.Recurring.Weekly.Misc != nil {
							entry.Misc["TypeIpRecurringWeekly"] = o.Type.Ip.Recurring.Weekly.Misc
						}
						if o.Type.Ip.Recurring.Weekly.At != nil {
							nestedType.Ip.Recurring.Weekly.At = o.Type.Ip.Recurring.Weekly.At
						}
						if o.Type.Ip.Recurring.Weekly.DayOfWeek != nil {
							nestedType.Ip.Recurring.Weekly.DayOfWeek = o.Type.Ip.Recurring.Weekly.DayOfWeek
						}
					}
					if o.Type.Ip.Recurring.Daily != nil {
						nestedType.Ip.Recurring.Daily = &TypeIpRecurringDaily{}
						if o.Type.Ip.Recurring.Daily.Misc != nil {
							entry.Misc["TypeIpRecurringDaily"] = o.Type.Ip.Recurring.Daily.Misc
						}
						if o.Type.Ip.Recurring.Daily.At != nil {
							nestedType.Ip.Recurring.Daily.At = o.Type.Ip.Recurring.Daily.At
						}
					}
					if o.Type.Ip.Recurring.FiveMinute != nil {
						nestedType.Ip.Recurring.FiveMinute = &TypeIpRecurringFiveMinute{}
						if o.Type.Ip.Recurring.FiveMinute.Misc != nil {
							entry.Misc["TypeIpRecurringFiveMinute"] = o.Type.Ip.Recurring.FiveMinute.Misc
						}
					}
				}
				if o.Type.Ip.Url != nil {
					nestedType.Ip.Url = o.Type.Ip.Url
				}
			}
			if o.Type.PredefinedIp != nil {
				nestedType.PredefinedIp = &TypePredefinedIp{}
				if o.Type.PredefinedIp.Misc != nil {
					entry.Misc["TypePredefinedIp"] = o.Type.PredefinedIp.Misc
				}
				if o.Type.PredefinedIp.Url != nil {
					nestedType.PredefinedIp.Url = o.Type.PredefinedIp.Url
				}
				if o.Type.PredefinedIp.Description != nil {
					nestedType.PredefinedIp.Description = o.Type.PredefinedIp.Description
				}
				if o.Type.PredefinedIp.ExceptionList != nil {
					nestedType.PredefinedIp.ExceptionList = util.MemToStr(o.Type.PredefinedIp.ExceptionList)
				}
			}
			if o.Type.PredefinedUrl != nil {
				nestedType.PredefinedUrl = &TypePredefinedUrl{}
				if o.Type.PredefinedUrl.Misc != nil {
					entry.Misc["TypePredefinedUrl"] = o.Type.PredefinedUrl.Misc
				}
				if o.Type.PredefinedUrl.ExceptionList != nil {
					nestedType.PredefinedUrl.ExceptionList = util.MemToStr(o.Type.PredefinedUrl.ExceptionList)
				}
				if o.Type.PredefinedUrl.Url != nil {
					nestedType.PredefinedUrl.Url = o.Type.PredefinedUrl.Url
				}
				if o.Type.PredefinedUrl.Description != nil {
					nestedType.PredefinedUrl.Description = o.Type.PredefinedUrl.Description
				}
			}
			if o.Type.Url != nil {
				nestedType.Url = &TypeUrl{}
				if o.Type.Url.Misc != nil {
					entry.Misc["TypeUrl"] = o.Type.Url.Misc
				}
				if o.Type.Url.Recurring != nil {
					nestedType.Url.Recurring = &TypeUrlRecurring{}
					if o.Type.Url.Recurring.Misc != nil {
						entry.Misc["TypeUrlRecurring"] = o.Type.Url.Recurring.Misc
					}
					if o.Type.Url.Recurring.Daily != nil {
						nestedType.Url.Recurring.Daily = &TypeUrlRecurringDaily{}
						if o.Type.Url.Recurring.Daily.Misc != nil {
							entry.Misc["TypeUrlRecurringDaily"] = o.Type.Url.Recurring.Daily.Misc
						}
						if o.Type.Url.Recurring.Daily.At != nil {
							nestedType.Url.Recurring.Daily.At = o.Type.Url.Recurring.Daily.At
						}
					}
					if o.Type.Url.Recurring.FiveMinute != nil {
						nestedType.Url.Recurring.FiveMinute = &TypeUrlRecurringFiveMinute{}
						if o.Type.Url.Recurring.FiveMinute.Misc != nil {
							entry.Misc["TypeUrlRecurringFiveMinute"] = o.Type.Url.Recurring.FiveMinute.Misc
						}
					}
					if o.Type.Url.Recurring.Hourly != nil {
						nestedType.Url.Recurring.Hourly = &TypeUrlRecurringHourly{}
						if o.Type.Url.Recurring.Hourly.Misc != nil {
							entry.Misc["TypeUrlRecurringHourly"] = o.Type.Url.Recurring.Hourly.Misc
						}
					}
					if o.Type.Url.Recurring.Monthly != nil {
						nestedType.Url.Recurring.Monthly = &TypeUrlRecurringMonthly{}
						if o.Type.Url.Recurring.Monthly.Misc != nil {
							entry.Misc["TypeUrlRecurringMonthly"] = o.Type.Url.Recurring.Monthly.Misc
						}
						if o.Type.Url.Recurring.Monthly.At != nil {
							nestedType.Url.Recurring.Monthly.At = o.Type.Url.Recurring.Monthly.At
						}
						if o.Type.Url.Recurring.Monthly.DayOfMonth != nil {
							nestedType.Url.Recurring.Monthly.DayOfMonth = o.Type.Url.Recurring.Monthly.DayOfMonth
						}
					}
					if o.Type.Url.Recurring.Weekly != nil {
						nestedType.Url.Recurring.Weekly = &TypeUrlRecurringWeekly{}
						if o.Type.Url.Recurring.Weekly.Misc != nil {
							entry.Misc["TypeUrlRecurringWeekly"] = o.Type.Url.Recurring.Weekly.Misc
						}
						if o.Type.Url.Recurring.Weekly.At != nil {
							nestedType.Url.Recurring.Weekly.At = o.Type.Url.Recurring.Weekly.At
						}
						if o.Type.Url.Recurring.Weekly.DayOfWeek != nil {
							nestedType.Url.Recurring.Weekly.DayOfWeek = o.Type.Url.Recurring.Weekly.DayOfWeek
						}
					}
				}
				if o.Type.Url.Url != nil {
					nestedType.Url.Url = o.Type.Url.Url
				}
				if o.Type.Url.Auth != nil {
					nestedType.Url.Auth = &TypeUrlAuth{}
					if o.Type.Url.Auth.Misc != nil {
						entry.Misc["TypeUrlAuth"] = o.Type.Url.Auth.Misc
					}
					if o.Type.Url.Auth.Password != nil {
						nestedType.Url.Auth.Password = o.Type.Url.Auth.Password
					}
					if o.Type.Url.Auth.Username != nil {
						nestedType.Url.Auth.Username = o.Type.Url.Auth.Username
					}
				}
				if o.Type.Url.CertificateProfile != nil {
					nestedType.Url.CertificateProfile = o.Type.Url.CertificateProfile
				}
				if o.Type.Url.Description != nil {
					nestedType.Url.Description = o.Type.Url.Description
				}
				if o.Type.Url.ExceptionList != nil {
					nestedType.Url.ExceptionList = util.MemToStr(o.Type.Url.ExceptionList)
				}
			}
		}
		entry.Type = nestedType

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
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !matchType(a.Type, b.Type) {
		return false
	}

	return true
}

func matchTypePredefinedUrl(a *TypePredefinedUrl, b *TypePredefinedUrl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	return true
}
func matchTypeUrlAuth(a *TypeUrlAuth, b *TypeUrlAuth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Password, b.Password) {
		return false
	}
	if !util.StringsMatch(a.Username, b.Username) {
		return false
	}
	return true
}
func matchTypeUrlRecurringDaily(a *TypeUrlRecurringDaily, b *TypeUrlRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeUrlRecurringFiveMinute(a *TypeUrlRecurringFiveMinute, b *TypeUrlRecurringFiveMinute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeUrlRecurringHourly(a *TypeUrlRecurringHourly, b *TypeUrlRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeUrlRecurringMonthly(a *TypeUrlRecurringMonthly, b *TypeUrlRecurringMonthly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.Ints64Match(a.DayOfMonth, b.DayOfMonth) {
		return false
	}
	return true
}
func matchTypeUrlRecurringWeekly(a *TypeUrlRecurringWeekly, b *TypeUrlRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchTypeUrlRecurring(a *TypeUrlRecurring, b *TypeUrlRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeUrlRecurringFiveMinute(a.FiveMinute, b.FiveMinute) {
		return false
	}
	if !matchTypeUrlRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchTypeUrlRecurringMonthly(a.Monthly, b.Monthly) {
		return false
	}
	if !matchTypeUrlRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	if !matchTypeUrlRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	return true
}
func matchTypeUrl(a *TypeUrl, b *TypeUrl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	if !matchTypeUrlAuth(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !matchTypeUrlRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchTypeDomainRecurringMonthly(a *TypeDomainRecurringMonthly, b *TypeDomainRecurringMonthly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.Ints64Match(a.DayOfMonth, b.DayOfMonth) {
		return false
	}
	return true
}
func matchTypeDomainRecurringWeekly(a *TypeDomainRecurringWeekly, b *TypeDomainRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeDomainRecurringDaily(a *TypeDomainRecurringDaily, b *TypeDomainRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeDomainRecurringFiveMinute(a *TypeDomainRecurringFiveMinute, b *TypeDomainRecurringFiveMinute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeDomainRecurringHourly(a *TypeDomainRecurringHourly, b *TypeDomainRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeDomainRecurring(a *TypeDomainRecurring, b *TypeDomainRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeDomainRecurringMonthly(a.Monthly, b.Monthly) {
		return false
	}
	if !matchTypeDomainRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	if !matchTypeDomainRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchTypeDomainRecurringFiveMinute(a.FiveMinute, b.FiveMinute) {
		return false
	}
	if !matchTypeDomainRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	return true
}
func matchTypeDomainAuth(a *TypeDomainAuth, b *TypeDomainAuth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Password, b.Password) {
		return false
	}
	if !util.StringsMatch(a.Username, b.Username) {
		return false
	}
	return true
}
func matchTypeDomain(a *TypeDomain, b *TypeDomain) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.ExpandDomain, b.ExpandDomain) {
		return false
	}
	if !matchTypeDomainRecurring(a.Recurring, b.Recurring) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	if !matchTypeDomainAuth(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	return true
}
func matchTypeImeiAuth(a *TypeImeiAuth, b *TypeImeiAuth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Username, b.Username) {
		return false
	}
	if !util.StringsMatch(a.Password, b.Password) {
		return false
	}
	return true
}
func matchTypeImeiRecurringDaily(a *TypeImeiRecurringDaily, b *TypeImeiRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeImeiRecurringFiveMinute(a *TypeImeiRecurringFiveMinute, b *TypeImeiRecurringFiveMinute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeImeiRecurringHourly(a *TypeImeiRecurringHourly, b *TypeImeiRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeImeiRecurringMonthly(a *TypeImeiRecurringMonthly, b *TypeImeiRecurringMonthly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.Ints64Match(a.DayOfMonth, b.DayOfMonth) {
		return false
	}
	return true
}
func matchTypeImeiRecurringWeekly(a *TypeImeiRecurringWeekly, b *TypeImeiRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchTypeImeiRecurring(a *TypeImeiRecurring, b *TypeImeiRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeImeiRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchTypeImeiRecurringMonthly(a.Monthly, b.Monthly) {
		return false
	}
	if !matchTypeImeiRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	if !matchTypeImeiRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchTypeImeiRecurringFiveMinute(a.FiveMinute, b.FiveMinute) {
		return false
	}
	return true
}
func matchTypeImei(a *TypeImei, b *TypeImei) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeImeiAuth(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !matchTypeImeiRecurring(a.Recurring, b.Recurring) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	return true
}
func matchTypeImsiAuth(a *TypeImsiAuth, b *TypeImsiAuth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Password, b.Password) {
		return false
	}
	if !util.StringsMatch(a.Username, b.Username) {
		return false
	}
	return true
}
func matchTypeImsiRecurringDaily(a *TypeImsiRecurringDaily, b *TypeImsiRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeImsiRecurringFiveMinute(a *TypeImsiRecurringFiveMinute, b *TypeImsiRecurringFiveMinute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeImsiRecurringHourly(a *TypeImsiRecurringHourly, b *TypeImsiRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeImsiRecurringMonthly(a *TypeImsiRecurringMonthly, b *TypeImsiRecurringMonthly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.Ints64Match(a.DayOfMonth, b.DayOfMonth) {
		return false
	}
	return true
}
func matchTypeImsiRecurringWeekly(a *TypeImsiRecurringWeekly, b *TypeImsiRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchTypeImsiRecurring(a *TypeImsiRecurring, b *TypeImsiRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeImsiRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchTypeImsiRecurringFiveMinute(a.FiveMinute, b.FiveMinute) {
		return false
	}
	if !matchTypeImsiRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchTypeImsiRecurringMonthly(a.Monthly, b.Monthly) {
		return false
	}
	if !matchTypeImsiRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	return true
}
func matchTypeImsi(a *TypeImsi, b *TypeImsi) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeImsiAuth(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !matchTypeImsiRecurring(a.Recurring, b.Recurring) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	return true
}
func matchTypeIpAuth(a *TypeIpAuth, b *TypeIpAuth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Username, b.Username) {
		return false
	}
	if !util.StringsMatch(a.Password, b.Password) {
		return false
	}
	return true
}
func matchTypeIpRecurringHourly(a *TypeIpRecurringHourly, b *TypeIpRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeIpRecurringMonthly(a *TypeIpRecurringMonthly, b *TypeIpRecurringMonthly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.Ints64Match(a.DayOfMonth, b.DayOfMonth) {
		return false
	}
	return true
}
func matchTypeIpRecurringWeekly(a *TypeIpRecurringWeekly, b *TypeIpRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchTypeIpRecurringDaily(a *TypeIpRecurringDaily, b *TypeIpRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchTypeIpRecurringFiveMinute(a *TypeIpRecurringFiveMinute, b *TypeIpRecurringFiveMinute) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchTypeIpRecurring(a *TypeIpRecurring, b *TypeIpRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeIpRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	if !matchTypeIpRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchTypeIpRecurringFiveMinute(a.FiveMinute, b.FiveMinute) {
		return false
	}
	if !matchTypeIpRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchTypeIpRecurringMonthly(a.Monthly, b.Monthly) {
		return false
	}
	return true
}
func matchTypeIp(a *TypeIp, b *TypeIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeIpAuth(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !matchTypeIpRecurring(a.Recurring, b.Recurring) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	return true
}
func matchTypePredefinedIp(a *TypePredefinedIp, b *TypePredefinedIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.OrderedListsMatch(a.ExceptionList, b.ExceptionList) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	return true
}
func matchType(a *Type, b *Type) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchTypeDomain(a.Domain, b.Domain) {
		return false
	}
	if !matchTypeImei(a.Imei, b.Imei) {
		return false
	}
	if !matchTypeImsi(a.Imsi, b.Imsi) {
		return false
	}
	if !matchTypeIp(a.Ip, b.Ip) {
		return false
	}
	if !matchTypePredefinedIp(a.PredefinedIp, b.PredefinedIp) {
		return false
	}
	if !matchTypePredefinedUrl(a.PredefinedUrl, b.PredefinedUrl) {
		return false
	}
	if !matchTypeUrl(a.Url, b.Url) {
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
