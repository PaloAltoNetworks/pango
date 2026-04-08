package schedules

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
	suffix = []string{"schedule", "$name"}
)

type Entry struct {
	Name            string
	DisableOverride *string
	ScheduleType    *ScheduleType
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type ScheduleType struct {
	NonRecurring   []string
	Recurring      *ScheduleTypeRecurring
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ScheduleTypeRecurring struct {
	Daily          []string
	Weekly         *ScheduleTypeRecurringWeekly
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type ScheduleTypeRecurringWeekly struct {
	Friday         []string
	Monday         []string
	Saturday       []string
	Sunday         []string
	Thursday       []string
	Tuesday        []string
	Wednesday      []string
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
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	DisableOverride *string          `xml:"disable-override,omitempty"`
	ScheduleType    *scheduleTypeXml `xml:"schedule-type,omitempty"`
	Misc            []generic.Xml    `xml:",any"`
	MiscAttributes  []xml.Attr       `xml:",any,attr"`
}
type scheduleTypeXml struct {
	NonRecurring   *util.MemberType          `xml:"non-recurring,omitempty"`
	Recurring      *scheduleTypeRecurringXml `xml:"recurring,omitempty"`
	Misc           []generic.Xml             `xml:",any"`
	MiscAttributes []xml.Attr                `xml:",any,attr"`
}
type scheduleTypeRecurringXml struct {
	Daily          *util.MemberType                `xml:"daily,omitempty"`
	Weekly         *scheduleTypeRecurringWeeklyXml `xml:"weekly,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type scheduleTypeRecurringWeeklyXml struct {
	Friday         *util.MemberType `xml:"friday,omitempty"`
	Monday         *util.MemberType `xml:"monday,omitempty"`
	Saturday       *util.MemberType `xml:"saturday,omitempty"`
	Sunday         *util.MemberType `xml:"sunday,omitempty"`
	Thursday       *util.MemberType `xml:"thursday,omitempty"`
	Tuesday        *util.MemberType `xml:"tuesday,omitempty"`
	Wednesday      *util.MemberType `xml:"wednesday,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.DisableOverride = s.DisableOverride
	if s.ScheduleType != nil {
		var obj scheduleTypeXml
		obj.MarshalFromObject(*s.ScheduleType)
		o.ScheduleType = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var scheduleTypeVal *ScheduleType
	if o.ScheduleType != nil {
		obj, err := o.ScheduleType.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		scheduleTypeVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		DisableOverride: o.DisableOverride,
		ScheduleType:    scheduleTypeVal,
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *scheduleTypeXml) MarshalFromObject(s ScheduleType) {
	if s.NonRecurring != nil {
		o.NonRecurring = util.StrToMem(s.NonRecurring)
	}
	if s.Recurring != nil {
		var obj scheduleTypeRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o scheduleTypeXml) UnmarshalToObject() (*ScheduleType, error) {
	var nonRecurringVal []string
	if o.NonRecurring != nil {
		nonRecurringVal = util.MemToStr(o.NonRecurring)
	}
	var recurringVal *ScheduleTypeRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &ScheduleType{
		NonRecurring:   nonRecurringVal,
		Recurring:      recurringVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *scheduleTypeRecurringXml) MarshalFromObject(s ScheduleTypeRecurring) {
	if s.Daily != nil {
		o.Daily = util.StrToMem(s.Daily)
	}
	if s.Weekly != nil {
		var obj scheduleTypeRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o scheduleTypeRecurringXml) UnmarshalToObject() (*ScheduleTypeRecurring, error) {
	var dailyVal []string
	if o.Daily != nil {
		dailyVal = util.MemToStr(o.Daily)
	}
	var weeklyVal *ScheduleTypeRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &ScheduleTypeRecurring{
		Daily:          dailyVal,
		Weekly:         weeklyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *scheduleTypeRecurringWeeklyXml) MarshalFromObject(s ScheduleTypeRecurringWeekly) {
	if s.Friday != nil {
		o.Friday = util.StrToMem(s.Friday)
	}
	if s.Monday != nil {
		o.Monday = util.StrToMem(s.Monday)
	}
	if s.Saturday != nil {
		o.Saturday = util.StrToMem(s.Saturday)
	}
	if s.Sunday != nil {
		o.Sunday = util.StrToMem(s.Sunday)
	}
	if s.Thursday != nil {
		o.Thursday = util.StrToMem(s.Thursday)
	}
	if s.Tuesday != nil {
		o.Tuesday = util.StrToMem(s.Tuesday)
	}
	if s.Wednesday != nil {
		o.Wednesday = util.StrToMem(s.Wednesday)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o scheduleTypeRecurringWeeklyXml) UnmarshalToObject() (*ScheduleTypeRecurringWeekly, error) {
	var fridayVal []string
	if o.Friday != nil {
		fridayVal = util.MemToStr(o.Friday)
	}
	var mondayVal []string
	if o.Monday != nil {
		mondayVal = util.MemToStr(o.Monday)
	}
	var saturdayVal []string
	if o.Saturday != nil {
		saturdayVal = util.MemToStr(o.Saturday)
	}
	var sundayVal []string
	if o.Sunday != nil {
		sundayVal = util.MemToStr(o.Sunday)
	}
	var thursdayVal []string
	if o.Thursday != nil {
		thursdayVal = util.MemToStr(o.Thursday)
	}
	var tuesdayVal []string
	if o.Tuesday != nil {
		tuesdayVal = util.MemToStr(o.Tuesday)
	}
	var wednesdayVal []string
	if o.Wednesday != nil {
		wednesdayVal = util.MemToStr(o.Wednesday)
	}

	result := &ScheduleTypeRecurringWeekly{
		Friday:         fridayVal,
		Monday:         mondayVal,
		Saturday:       saturdayVal,
		Sunday:         sundayVal,
		Thursday:       thursdayVal,
		Tuesday:        tuesdayVal,
		Wednesday:      wednesdayVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
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
	if v == "schedule_type" || v == "ScheduleType" {
		return e.ScheduleType, nil
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
	if !o.ScheduleType.matches(other.ScheduleType) {
		return false
	}

	return true
}

func (o *ScheduleType) matches(other *ScheduleType) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.NonRecurring, other.NonRecurring) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *ScheduleTypeRecurring) matches(other *ScheduleTypeRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Daily, other.Daily) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *ScheduleTypeRecurringWeekly) matches(other *ScheduleTypeRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Friday, other.Friday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Monday, other.Monday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Saturday, other.Saturday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Sunday, other.Sunday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Thursday, other.Thursday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Tuesday, other.Tuesday) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Wednesday, other.Wednesday) {
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
