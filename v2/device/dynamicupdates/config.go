package dynamicupdates

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

type Config struct {
	UpdateSchedule *UpdateSchedule
	Misc           []generic.Xml
}
type UpdateSchedule struct {
	AntiVirus                  *UpdateScheduleAntiVirus
	AppProfile                 *UpdateScheduleAppProfile
	GlobalProtectClientlessVpn *UpdateScheduleGlobalProtectClientlessVpn
	GlobalProtectDatafile      *UpdateScheduleGlobalProtectDatafile
	StatisticsService          *UpdateScheduleStatisticsService
	Threats                    *UpdateScheduleThreats
	WfPrivate                  *UpdateScheduleWfPrivate
	Wildfire                   *UpdateScheduleWildfire
	Misc                       []generic.Xml
}
type UpdateScheduleAntiVirus struct {
	Recurring *UpdateScheduleAntiVirusRecurring
	Misc      []generic.Xml
}
type UpdateScheduleAntiVirusRecurring struct {
	SyncToPeer *bool
	Threshold  *int64
	Daily      *UpdateScheduleAntiVirusRecurringDaily
	Hourly     *UpdateScheduleAntiVirusRecurringHourly
	None       *UpdateScheduleAntiVirusRecurringNone
	Weekly     *UpdateScheduleAntiVirusRecurringWeekly
	Misc       []generic.Xml
}
type UpdateScheduleAntiVirusRecurringDaily struct {
	Action *string
	At     *string
	Misc   []generic.Xml
}
type UpdateScheduleAntiVirusRecurringHourly struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleAntiVirusRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleAntiVirusRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type UpdateScheduleAppProfile struct {
	Recurring *UpdateScheduleAppProfileRecurring
	Misc      []generic.Xml
}
type UpdateScheduleAppProfileRecurring struct {
	SyncToPeer *bool
	Threshold  *int64
	Daily      *UpdateScheduleAppProfileRecurringDaily
	None       *UpdateScheduleAppProfileRecurringNone
	Weekly     *UpdateScheduleAppProfileRecurringWeekly
	Misc       []generic.Xml
}
type UpdateScheduleAppProfileRecurringDaily struct {
	Action *string
	At     *string
	Misc   []generic.Xml
}
type UpdateScheduleAppProfileRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleAppProfileRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpn struct {
	Recurring *UpdateScheduleGlobalProtectClientlessVpnRecurring
	Misc      []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpnRecurring struct {
	Daily  *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily
	Hourly *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly
	None   *UpdateScheduleGlobalProtectClientlessVpnRecurringNone
	Weekly *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringDaily struct {
	Action *string
	At     *string
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringHourly struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type UpdateScheduleGlobalProtectDatafile struct {
	Recurring *UpdateScheduleGlobalProtectDatafileRecurring
	Misc      []generic.Xml
}
type UpdateScheduleGlobalProtectDatafileRecurring struct {
	Daily  *UpdateScheduleGlobalProtectDatafileRecurringDaily
	Hourly *UpdateScheduleGlobalProtectDatafileRecurringHourly
	None   *UpdateScheduleGlobalProtectDatafileRecurringNone
	Weekly *UpdateScheduleGlobalProtectDatafileRecurringWeekly
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectDatafileRecurringDaily struct {
	Action *string
	At     *string
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectDatafileRecurringHourly struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleGlobalProtectDatafileRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleGlobalProtectDatafileRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
	Misc      []generic.Xml
}
type UpdateScheduleStatisticsService struct {
	ApplicationReports          *bool
	FileIdentificationReports   *bool
	HealthPerformanceReports    *bool
	PassiveDnsMonitoring        *bool
	ThreatPreventionInformation *bool
	ThreatPreventionPcap        *bool
	ThreatPreventionReports     *bool
	UrlReports                  *bool
	Misc                        []generic.Xml
}
type UpdateScheduleThreats struct {
	Recurring *UpdateScheduleThreatsRecurring
	Misc      []generic.Xml
}
type UpdateScheduleThreatsRecurring struct {
	NewAppThreshold *int64
	SyncToPeer      *bool
	Threshold       *int64
	Daily           *UpdateScheduleThreatsRecurringDaily
	Every30Mins     *UpdateScheduleThreatsRecurringEvery30Mins
	Hourly          *UpdateScheduleThreatsRecurringHourly
	None            *UpdateScheduleThreatsRecurringNone
	Weekly          *UpdateScheduleThreatsRecurringWeekly
	Misc            []generic.Xml
}
type UpdateScheduleThreatsRecurringDaily struct {
	Action            *string
	At                *string
	DisableNewContent *bool
	Misc              []generic.Xml
}
type UpdateScheduleThreatsRecurringEvery30Mins struct {
	Action            *string
	At                *int64
	DisableNewContent *bool
	Misc              []generic.Xml
}
type UpdateScheduleThreatsRecurringHourly struct {
	Action            *string
	At                *int64
	DisableNewContent *bool
	Misc              []generic.Xml
}
type UpdateScheduleThreatsRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleThreatsRecurringWeekly struct {
	Action            *string
	At                *string
	DayOfWeek         *string
	DisableNewContent *bool
	Misc              []generic.Xml
}
type UpdateScheduleWfPrivate struct {
	Recurring *UpdateScheduleWfPrivateRecurring
	Misc      []generic.Xml
}
type UpdateScheduleWfPrivateRecurring struct {
	SyncToPeer  *bool
	Every15Mins *UpdateScheduleWfPrivateRecurringEvery15Mins
	Every30Mins *UpdateScheduleWfPrivateRecurringEvery30Mins
	Every5Mins  *UpdateScheduleWfPrivateRecurringEvery5Mins
	EveryHour   *UpdateScheduleWfPrivateRecurringEveryHour
	None        *UpdateScheduleWfPrivateRecurringNone
	Misc        []generic.Xml
}
type UpdateScheduleWfPrivateRecurringEvery15Mins struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleWfPrivateRecurringEvery30Mins struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleWfPrivateRecurringEvery5Mins struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleWfPrivateRecurringEveryHour struct {
	Action *string
	At     *int64
	Misc   []generic.Xml
}
type UpdateScheduleWfPrivateRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleWildfire struct {
	Recurring *UpdateScheduleWildfireRecurring
	Misc      []generic.Xml
}
type UpdateScheduleWildfireRecurring struct {
	Every15Mins *UpdateScheduleWildfireRecurringEvery15Mins
	Every30Mins *UpdateScheduleWildfireRecurringEvery30Mins
	EveryHour   *UpdateScheduleWildfireRecurringEveryHour
	EveryMin    *UpdateScheduleWildfireRecurringEveryMin
	None        *UpdateScheduleWildfireRecurringNone
	RealTime    *UpdateScheduleWildfireRecurringRealTime
	Misc        []generic.Xml
}
type UpdateScheduleWildfireRecurringEvery15Mins struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
	Misc       []generic.Xml
}
type UpdateScheduleWildfireRecurringEvery30Mins struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
	Misc       []generic.Xml
}
type UpdateScheduleWildfireRecurringEveryHour struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
	Misc       []generic.Xml
}
type UpdateScheduleWildfireRecurringEveryMin struct {
	Action     *string
	SyncToPeer *bool
	Misc       []generic.Xml
}
type UpdateScheduleWildfireRecurringNone struct {
	Misc []generic.Xml
}
type UpdateScheduleWildfireRecurringRealTime struct {
	Misc []generic.Xml
}

type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}

func (o *configXmlContainer) Normalize() ([]*Config, error) {
	entries := make([]*Config, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyConfig(source *Config) (any, error) {
	var obj configXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type configXml struct {
	XMLName        xml.Name           `xml:"system"`
	UpdateSchedule *updateScheduleXml `xml:"update-schedule,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
}
type updateScheduleXml struct {
	AntiVirus                  *updateScheduleAntiVirusXml                  `xml:"anti-virus,omitempty"`
	AppProfile                 *updateScheduleAppProfileXml                 `xml:"app-profile,omitempty"`
	GlobalProtectClientlessVpn *updateScheduleGlobalProtectClientlessVpnXml `xml:"global-protect-clientless-vpn,omitempty"`
	GlobalProtectDatafile      *updateScheduleGlobalProtectDatafileXml      `xml:"global-protect-datafile,omitempty"`
	StatisticsService          *updateScheduleStatisticsServiceXml          `xml:"statistics-service,omitempty"`
	Threats                    *updateScheduleThreatsXml                    `xml:"threats,omitempty"`
	WfPrivate                  *updateScheduleWfPrivateXml                  `xml:"wf-private,omitempty"`
	Wildfire                   *updateScheduleWildfireXml                   `xml:"wildfire,omitempty"`
	Misc                       []generic.Xml                                `xml:",any"`
}
type updateScheduleAntiVirusXml struct {
	Recurring *updateScheduleAntiVirusRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                        `xml:",any"`
}
type updateScheduleAntiVirusRecurringXml struct {
	SyncToPeer *string                                    `xml:"sync-to-peer,omitempty"`
	Threshold  *int64                                     `xml:"threshold,omitempty"`
	Daily      *updateScheduleAntiVirusRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly     *updateScheduleAntiVirusRecurringHourlyXml `xml:"hourly,omitempty"`
	None       *updateScheduleAntiVirusRecurringNoneXml   `xml:"none,omitempty"`
	Weekly     *updateScheduleAntiVirusRecurringWeeklyXml `xml:"weekly,omitempty"`
	Misc       []generic.Xml                              `xml:",any"`
}
type updateScheduleAntiVirusRecurringDailyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *string       `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleAntiVirusRecurringHourlyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleAntiVirusRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleAntiVirusRecurringWeeklyXml struct {
	Action    *string       `xml:"action,omitempty"`
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type updateScheduleAppProfileXml struct {
	Recurring *updateScheduleAppProfileRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                         `xml:",any"`
}
type updateScheduleAppProfileRecurringXml struct {
	SyncToPeer *string                                     `xml:"sync-to-peer,omitempty"`
	Threshold  *int64                                      `xml:"threshold,omitempty"`
	Daily      *updateScheduleAppProfileRecurringDailyXml  `xml:"daily,omitempty"`
	None       *updateScheduleAppProfileRecurringNoneXml   `xml:"none,omitempty"`
	Weekly     *updateScheduleAppProfileRecurringWeeklyXml `xml:"weekly,omitempty"`
	Misc       []generic.Xml                               `xml:",any"`
}
type updateScheduleAppProfileRecurringDailyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *string       `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleAppProfileRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleAppProfileRecurringWeeklyXml struct {
	Action    *string       `xml:"action,omitempty"`
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnXml struct {
	Recurring *updateScheduleGlobalProtectClientlessVpnRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                                         `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnRecurringXml struct {
	Daily  *updateScheduleGlobalProtectClientlessVpnRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly *updateScheduleGlobalProtectClientlessVpnRecurringHourlyXml `xml:"hourly,omitempty"`
	None   *updateScheduleGlobalProtectClientlessVpnRecurringNoneXml   `xml:"none,omitempty"`
	Weekly *updateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml `xml:"weekly,omitempty"`
	Misc   []generic.Xml                                               `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnRecurringDailyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *string       `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnRecurringHourlyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml struct {
	Action    *string       `xml:"action,omitempty"`
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectDatafileXml struct {
	Recurring *updateScheduleGlobalProtectDatafileRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                                    `xml:",any"`
}
type updateScheduleGlobalProtectDatafileRecurringXml struct {
	Daily  *updateScheduleGlobalProtectDatafileRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly *updateScheduleGlobalProtectDatafileRecurringHourlyXml `xml:"hourly,omitempty"`
	None   *updateScheduleGlobalProtectDatafileRecurringNoneXml   `xml:"none,omitempty"`
	Weekly *updateScheduleGlobalProtectDatafileRecurringWeeklyXml `xml:"weekly,omitempty"`
	Misc   []generic.Xml                                          `xml:",any"`
}
type updateScheduleGlobalProtectDatafileRecurringDailyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *string       `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectDatafileRecurringHourlyXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectDatafileRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleGlobalProtectDatafileRecurringWeeklyXml struct {
	Action    *string       `xml:"action,omitempty"`
	At        *string       `xml:"at,omitempty"`
	DayOfWeek *string       `xml:"day-of-week,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type updateScheduleStatisticsServiceXml struct {
	ApplicationReports          *string       `xml:"application-reports,omitempty"`
	FileIdentificationReports   *string       `xml:"file-identification-reports,omitempty"`
	HealthPerformanceReports    *string       `xml:"health-performance-reports,omitempty"`
	PassiveDnsMonitoring        *string       `xml:"passive-dns-monitoring,omitempty"`
	ThreatPreventionInformation *string       `xml:"threat-prevention-information,omitempty"`
	ThreatPreventionPcap        *string       `xml:"threat-prevention-pcap,omitempty"`
	ThreatPreventionReports     *string       `xml:"threat-prevention-reports,omitempty"`
	UrlReports                  *string       `xml:"url-reports,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type updateScheduleThreatsXml struct {
	Recurring *updateScheduleThreatsRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                      `xml:",any"`
}
type updateScheduleThreatsRecurringXml struct {
	NewAppThreshold *int64                                        `xml:"new-app-threshold,omitempty"`
	SyncToPeer      *string                                       `xml:"sync-to-peer,omitempty"`
	Threshold       *int64                                        `xml:"threshold,omitempty"`
	Daily           *updateScheduleThreatsRecurringDailyXml       `xml:"daily,omitempty"`
	Every30Mins     *updateScheduleThreatsRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	Hourly          *updateScheduleThreatsRecurringHourlyXml      `xml:"hourly,omitempty"`
	None            *updateScheduleThreatsRecurringNoneXml        `xml:"none,omitempty"`
	Weekly          *updateScheduleThreatsRecurringWeeklyXml      `xml:"weekly,omitempty"`
	Misc            []generic.Xml                                 `xml:",any"`
}
type updateScheduleThreatsRecurringDailyXml struct {
	Action            *string       `xml:"action,omitempty"`
	At                *string       `xml:"at,omitempty"`
	DisableNewContent *string       `xml:"disable-new-content,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type updateScheduleThreatsRecurringEvery30MinsXml struct {
	Action            *string       `xml:"action,omitempty"`
	At                *int64        `xml:"at,omitempty"`
	DisableNewContent *string       `xml:"disable-new-content,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type updateScheduleThreatsRecurringHourlyXml struct {
	Action            *string       `xml:"action,omitempty"`
	At                *int64        `xml:"at,omitempty"`
	DisableNewContent *string       `xml:"disable-new-content,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type updateScheduleThreatsRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleThreatsRecurringWeeklyXml struct {
	Action            *string       `xml:"action,omitempty"`
	At                *string       `xml:"at,omitempty"`
	DayOfWeek         *string       `xml:"day-of-week,omitempty"`
	DisableNewContent *string       `xml:"disable-new-content,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type updateScheduleWfPrivateXml struct {
	Recurring *updateScheduleWfPrivateRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                        `xml:",any"`
}
type updateScheduleWfPrivateRecurringXml struct {
	SyncToPeer  *string                                         `xml:"sync-to-peer,omitempty"`
	Every15Mins *updateScheduleWfPrivateRecurringEvery15MinsXml `xml:"every-15-mins,omitempty"`
	Every30Mins *updateScheduleWfPrivateRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	Every5Mins  *updateScheduleWfPrivateRecurringEvery5MinsXml  `xml:"every-5-mins,omitempty"`
	EveryHour   *updateScheduleWfPrivateRecurringEveryHourXml   `xml:"every-hour,omitempty"`
	None        *updateScheduleWfPrivateRecurringNoneXml        `xml:"none,omitempty"`
	Misc        []generic.Xml                                   `xml:",any"`
}
type updateScheduleWfPrivateRecurringEvery15MinsXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleWfPrivateRecurringEvery30MinsXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleWfPrivateRecurringEvery5MinsXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleWfPrivateRecurringEveryHourXml struct {
	Action *string       `xml:"action,omitempty"`
	At     *int64        `xml:"at,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type updateScheduleWfPrivateRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleWildfireXml struct {
	Recurring *updateScheduleWildfireRecurringXml `xml:"recurring,omitempty"`
	Misc      []generic.Xml                       `xml:",any"`
}
type updateScheduleWildfireRecurringXml struct {
	Every15Mins *updateScheduleWildfireRecurringEvery15MinsXml `xml:"every-15-mins,omitempty"`
	Every30Mins *updateScheduleWildfireRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	EveryHour   *updateScheduleWildfireRecurringEveryHourXml   `xml:"every-hour,omitempty"`
	EveryMin    *updateScheduleWildfireRecurringEveryMinXml    `xml:"every-min,omitempty"`
	None        *updateScheduleWildfireRecurringNoneXml        `xml:"none,omitempty"`
	RealTime    *updateScheduleWildfireRecurringRealTimeXml    `xml:"real-time,omitempty"`
	Misc        []generic.Xml                                  `xml:",any"`
}
type updateScheduleWildfireRecurringEvery15MinsXml struct {
	Action     *string       `xml:"action,omitempty"`
	At         *int64        `xml:"at,omitempty"`
	SyncToPeer *string       `xml:"sync-to-peer,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type updateScheduleWildfireRecurringEvery30MinsXml struct {
	Action     *string       `xml:"action,omitempty"`
	At         *int64        `xml:"at,omitempty"`
	SyncToPeer *string       `xml:"sync-to-peer,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type updateScheduleWildfireRecurringEveryHourXml struct {
	Action     *string       `xml:"action,omitempty"`
	At         *int64        `xml:"at,omitempty"`
	SyncToPeer *string       `xml:"sync-to-peer,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type updateScheduleWildfireRecurringEveryMinXml struct {
	Action     *string       `xml:"action,omitempty"`
	SyncToPeer *string       `xml:"sync-to-peer,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type updateScheduleWildfireRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type updateScheduleWildfireRecurringRealTimeXml struct {
	Misc []generic.Xml `xml:",any"`
}

func (o *configXml) MarshalFromObject(s Config) {
	if s.UpdateSchedule != nil {
		var obj updateScheduleXml
		obj.MarshalFromObject(*s.UpdateSchedule)
		o.UpdateSchedule = &obj
	}
	o.Misc = s.Misc
}

func (o configXml) UnmarshalToObject() (*Config, error) {
	var updateScheduleVal *UpdateSchedule
	if o.UpdateSchedule != nil {
		obj, err := o.UpdateSchedule.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		updateScheduleVal = obj
	}

	result := &Config{
		UpdateSchedule: updateScheduleVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *updateScheduleXml) MarshalFromObject(s UpdateSchedule) {
	if s.AntiVirus != nil {
		var obj updateScheduleAntiVirusXml
		obj.MarshalFromObject(*s.AntiVirus)
		o.AntiVirus = &obj
	}
	if s.AppProfile != nil {
		var obj updateScheduleAppProfileXml
		obj.MarshalFromObject(*s.AppProfile)
		o.AppProfile = &obj
	}
	if s.GlobalProtectClientlessVpn != nil {
		var obj updateScheduleGlobalProtectClientlessVpnXml
		obj.MarshalFromObject(*s.GlobalProtectClientlessVpn)
		o.GlobalProtectClientlessVpn = &obj
	}
	if s.GlobalProtectDatafile != nil {
		var obj updateScheduleGlobalProtectDatafileXml
		obj.MarshalFromObject(*s.GlobalProtectDatafile)
		o.GlobalProtectDatafile = &obj
	}
	if s.StatisticsService != nil {
		var obj updateScheduleStatisticsServiceXml
		obj.MarshalFromObject(*s.StatisticsService)
		o.StatisticsService = &obj
	}
	if s.Threats != nil {
		var obj updateScheduleThreatsXml
		obj.MarshalFromObject(*s.Threats)
		o.Threats = &obj
	}
	if s.WfPrivate != nil {
		var obj updateScheduleWfPrivateXml
		obj.MarshalFromObject(*s.WfPrivate)
		o.WfPrivate = &obj
	}
	if s.Wildfire != nil {
		var obj updateScheduleWildfireXml
		obj.MarshalFromObject(*s.Wildfire)
		o.Wildfire = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleXml) UnmarshalToObject() (*UpdateSchedule, error) {
	var antiVirusVal *UpdateScheduleAntiVirus
	if o.AntiVirus != nil {
		obj, err := o.AntiVirus.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		antiVirusVal = obj
	}
	var appProfileVal *UpdateScheduleAppProfile
	if o.AppProfile != nil {
		obj, err := o.AppProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		appProfileVal = obj
	}
	var globalProtectClientlessVpnVal *UpdateScheduleGlobalProtectClientlessVpn
	if o.GlobalProtectClientlessVpn != nil {
		obj, err := o.GlobalProtectClientlessVpn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectClientlessVpnVal = obj
	}
	var globalProtectDatafileVal *UpdateScheduleGlobalProtectDatafile
	if o.GlobalProtectDatafile != nil {
		obj, err := o.GlobalProtectDatafile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectDatafileVal = obj
	}
	var statisticsServiceVal *UpdateScheduleStatisticsService
	if o.StatisticsService != nil {
		obj, err := o.StatisticsService.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		statisticsServiceVal = obj
	}
	var threatsVal *UpdateScheduleThreats
	if o.Threats != nil {
		obj, err := o.Threats.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		threatsVal = obj
	}
	var wfPrivateVal *UpdateScheduleWfPrivate
	if o.WfPrivate != nil {
		obj, err := o.WfPrivate.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		wfPrivateVal = obj
	}
	var wildfireVal *UpdateScheduleWildfire
	if o.Wildfire != nil {
		obj, err := o.Wildfire.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		wildfireVal = obj
	}

	result := &UpdateSchedule{
		AntiVirus:                  antiVirusVal,
		AppProfile:                 appProfileVal,
		GlobalProtectClientlessVpn: globalProtectClientlessVpnVal,
		GlobalProtectDatafile:      globalProtectDatafileVal,
		StatisticsService:          statisticsServiceVal,
		Threats:                    threatsVal,
		WfPrivate:                  wfPrivateVal,
		Wildfire:                   wildfireVal,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusXml) MarshalFromObject(s UpdateScheduleAntiVirus) {
	if s.Recurring != nil {
		var obj updateScheduleAntiVirusRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusXml) UnmarshalToObject() (*UpdateScheduleAntiVirus, error) {
	var recurringVal *UpdateScheduleAntiVirusRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleAntiVirus{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusRecurringXml) MarshalFromObject(s UpdateScheduleAntiVirusRecurring) {
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Threshold = s.Threshold
	if s.Daily != nil {
		var obj updateScheduleAntiVirusRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.Hourly != nil {
		var obj updateScheduleAntiVirusRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.None != nil {
		var obj updateScheduleAntiVirusRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Weekly != nil {
		var obj updateScheduleAntiVirusRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusRecurringXml) UnmarshalToObject() (*UpdateScheduleAntiVirusRecurring, error) {
	var dailyVal *UpdateScheduleAntiVirusRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var hourlyVal *UpdateScheduleAntiVirusRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var noneVal *UpdateScheduleAntiVirusRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var weeklyVal *UpdateScheduleAntiVirusRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &UpdateScheduleAntiVirusRecurring{
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Threshold:  o.Threshold,
		Daily:      dailyVal,
		Hourly:     hourlyVal,
		None:       noneVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusRecurringDailyXml) MarshalFromObject(s UpdateScheduleAntiVirusRecurringDaily) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusRecurringDailyXml) UnmarshalToObject() (*UpdateScheduleAntiVirusRecurringDaily, error) {

	result := &UpdateScheduleAntiVirusRecurringDaily{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusRecurringHourlyXml) MarshalFromObject(s UpdateScheduleAntiVirusRecurringHourly) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusRecurringHourlyXml) UnmarshalToObject() (*UpdateScheduleAntiVirusRecurringHourly, error) {

	result := &UpdateScheduleAntiVirusRecurringHourly{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusRecurringNoneXml) MarshalFromObject(s UpdateScheduleAntiVirusRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleAntiVirusRecurringNone, error) {

	result := &UpdateScheduleAntiVirusRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAntiVirusRecurringWeeklyXml) MarshalFromObject(s UpdateScheduleAntiVirusRecurringWeekly) {
	o.Action = s.Action
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o updateScheduleAntiVirusRecurringWeeklyXml) UnmarshalToObject() (*UpdateScheduleAntiVirusRecurringWeekly, error) {

	result := &UpdateScheduleAntiVirusRecurringWeekly{
		Action:    o.Action,
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAppProfileXml) MarshalFromObject(s UpdateScheduleAppProfile) {
	if s.Recurring != nil {
		var obj updateScheduleAppProfileRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleAppProfileXml) UnmarshalToObject() (*UpdateScheduleAppProfile, error) {
	var recurringVal *UpdateScheduleAppProfileRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleAppProfile{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAppProfileRecurringXml) MarshalFromObject(s UpdateScheduleAppProfileRecurring) {
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Threshold = s.Threshold
	if s.Daily != nil {
		var obj updateScheduleAppProfileRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.None != nil {
		var obj updateScheduleAppProfileRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Weekly != nil {
		var obj updateScheduleAppProfileRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleAppProfileRecurringXml) UnmarshalToObject() (*UpdateScheduleAppProfileRecurring, error) {
	var dailyVal *UpdateScheduleAppProfileRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var noneVal *UpdateScheduleAppProfileRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var weeklyVal *UpdateScheduleAppProfileRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &UpdateScheduleAppProfileRecurring{
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Threshold:  o.Threshold,
		Daily:      dailyVal,
		None:       noneVal,
		Weekly:     weeklyVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAppProfileRecurringDailyXml) MarshalFromObject(s UpdateScheduleAppProfileRecurringDaily) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleAppProfileRecurringDailyXml) UnmarshalToObject() (*UpdateScheduleAppProfileRecurringDaily, error) {

	result := &UpdateScheduleAppProfileRecurringDaily{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAppProfileRecurringNoneXml) MarshalFromObject(s UpdateScheduleAppProfileRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleAppProfileRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleAppProfileRecurringNone, error) {

	result := &UpdateScheduleAppProfileRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleAppProfileRecurringWeeklyXml) MarshalFromObject(s UpdateScheduleAppProfileRecurringWeekly) {
	o.Action = s.Action
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o updateScheduleAppProfileRecurringWeeklyXml) UnmarshalToObject() (*UpdateScheduleAppProfileRecurringWeekly, error) {

	result := &UpdateScheduleAppProfileRecurringWeekly{
		Action:    o.Action,
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpn) {
	if s.Recurring != nil {
		var obj updateScheduleGlobalProtectClientlessVpnRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpn, error) {
	var recurringVal *UpdateScheduleGlobalProtectClientlessVpnRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleGlobalProtectClientlessVpn{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnRecurringXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpnRecurring) {
	if s.Daily != nil {
		var obj updateScheduleGlobalProtectClientlessVpnRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.Hourly != nil {
		var obj updateScheduleGlobalProtectClientlessVpnRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.None != nil {
		var obj updateScheduleGlobalProtectClientlessVpnRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Weekly != nil {
		var obj updateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnRecurringXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpnRecurring, error) {
	var dailyVal *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var hourlyVal *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var noneVal *UpdateScheduleGlobalProtectClientlessVpnRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var weeklyVal *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &UpdateScheduleGlobalProtectClientlessVpnRecurring{
		Daily:  dailyVal,
		Hourly: hourlyVal,
		None:   noneVal,
		Weekly: weeklyVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnRecurringDailyXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpnRecurringDaily) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnRecurringDailyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpnRecurringDaily, error) {

	result := &UpdateScheduleGlobalProtectClientlessVpnRecurringDaily{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnRecurringHourlyXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpnRecurringHourly) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnRecurringHourlyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpnRecurringHourly, error) {

	result := &UpdateScheduleGlobalProtectClientlessVpnRecurringHourly{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnRecurringNoneXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpnRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpnRecurringNone, error) {

	result := &UpdateScheduleGlobalProtectClientlessVpnRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml) MarshalFromObject(s UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly) {
	o.Action = s.Action
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly, error) {

	result := &UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly{
		Action:    o.Action,
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafile) {
	if s.Recurring != nil {
		var obj updateScheduleGlobalProtectDatafileRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafile, error) {
	var recurringVal *UpdateScheduleGlobalProtectDatafileRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleGlobalProtectDatafile{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileRecurringXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafileRecurring) {
	if s.Daily != nil {
		var obj updateScheduleGlobalProtectDatafileRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.Hourly != nil {
		var obj updateScheduleGlobalProtectDatafileRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.None != nil {
		var obj updateScheduleGlobalProtectDatafileRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Weekly != nil {
		var obj updateScheduleGlobalProtectDatafileRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileRecurringXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafileRecurring, error) {
	var dailyVal *UpdateScheduleGlobalProtectDatafileRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var hourlyVal *UpdateScheduleGlobalProtectDatafileRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var noneVal *UpdateScheduleGlobalProtectDatafileRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var weeklyVal *UpdateScheduleGlobalProtectDatafileRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &UpdateScheduleGlobalProtectDatafileRecurring{
		Daily:  dailyVal,
		Hourly: hourlyVal,
		None:   noneVal,
		Weekly: weeklyVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileRecurringDailyXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafileRecurringDaily) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileRecurringDailyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafileRecurringDaily, error) {

	result := &UpdateScheduleGlobalProtectDatafileRecurringDaily{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileRecurringHourlyXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafileRecurringHourly) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileRecurringHourlyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafileRecurringHourly, error) {

	result := &UpdateScheduleGlobalProtectDatafileRecurringHourly{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileRecurringNoneXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafileRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafileRecurringNone, error) {

	result := &UpdateScheduleGlobalProtectDatafileRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleGlobalProtectDatafileRecurringWeeklyXml) MarshalFromObject(s UpdateScheduleGlobalProtectDatafileRecurringWeekly) {
	o.Action = s.Action
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.Misc = s.Misc
}

func (o updateScheduleGlobalProtectDatafileRecurringWeeklyXml) UnmarshalToObject() (*UpdateScheduleGlobalProtectDatafileRecurringWeekly, error) {

	result := &UpdateScheduleGlobalProtectDatafileRecurringWeekly{
		Action:    o.Action,
		At:        o.At,
		DayOfWeek: o.DayOfWeek,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleStatisticsServiceXml) MarshalFromObject(s UpdateScheduleStatisticsService) {
	o.ApplicationReports = util.YesNo(s.ApplicationReports, nil)
	o.FileIdentificationReports = util.YesNo(s.FileIdentificationReports, nil)
	o.HealthPerformanceReports = util.YesNo(s.HealthPerformanceReports, nil)
	o.PassiveDnsMonitoring = util.YesNo(s.PassiveDnsMonitoring, nil)
	o.ThreatPreventionInformation = util.YesNo(s.ThreatPreventionInformation, nil)
	o.ThreatPreventionPcap = util.YesNo(s.ThreatPreventionPcap, nil)
	o.ThreatPreventionReports = util.YesNo(s.ThreatPreventionReports, nil)
	o.UrlReports = util.YesNo(s.UrlReports, nil)
	o.Misc = s.Misc
}

func (o updateScheduleStatisticsServiceXml) UnmarshalToObject() (*UpdateScheduleStatisticsService, error) {

	result := &UpdateScheduleStatisticsService{
		ApplicationReports:          util.AsBool(o.ApplicationReports, nil),
		FileIdentificationReports:   util.AsBool(o.FileIdentificationReports, nil),
		HealthPerformanceReports:    util.AsBool(o.HealthPerformanceReports, nil),
		PassiveDnsMonitoring:        util.AsBool(o.PassiveDnsMonitoring, nil),
		ThreatPreventionInformation: util.AsBool(o.ThreatPreventionInformation, nil),
		ThreatPreventionPcap:        util.AsBool(o.ThreatPreventionPcap, nil),
		ThreatPreventionReports:     util.AsBool(o.ThreatPreventionReports, nil),
		UrlReports:                  util.AsBool(o.UrlReports, nil),
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsXml) MarshalFromObject(s UpdateScheduleThreats) {
	if s.Recurring != nil {
		var obj updateScheduleThreatsRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleThreatsXml) UnmarshalToObject() (*UpdateScheduleThreats, error) {
	var recurringVal *UpdateScheduleThreatsRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleThreats{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringXml) MarshalFromObject(s UpdateScheduleThreatsRecurring) {
	o.NewAppThreshold = s.NewAppThreshold
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Threshold = s.Threshold
	if s.Daily != nil {
		var obj updateScheduleThreatsRecurringDailyXml
		obj.MarshalFromObject(*s.Daily)
		o.Daily = &obj
	}
	if s.Every30Mins != nil {
		var obj updateScheduleThreatsRecurringEvery30MinsXml
		obj.MarshalFromObject(*s.Every30Mins)
		o.Every30Mins = &obj
	}
	if s.Hourly != nil {
		var obj updateScheduleThreatsRecurringHourlyXml
		obj.MarshalFromObject(*s.Hourly)
		o.Hourly = &obj
	}
	if s.None != nil {
		var obj updateScheduleThreatsRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Weekly != nil {
		var obj updateScheduleThreatsRecurringWeeklyXml
		obj.MarshalFromObject(*s.Weekly)
		o.Weekly = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurring, error) {
	var dailyVal *UpdateScheduleThreatsRecurringDaily
	if o.Daily != nil {
		obj, err := o.Daily.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dailyVal = obj
	}
	var every30MinsVal *UpdateScheduleThreatsRecurringEvery30Mins
	if o.Every30Mins != nil {
		obj, err := o.Every30Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every30MinsVal = obj
	}
	var hourlyVal *UpdateScheduleThreatsRecurringHourly
	if o.Hourly != nil {
		obj, err := o.Hourly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		hourlyVal = obj
	}
	var noneVal *UpdateScheduleThreatsRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var weeklyVal *UpdateScheduleThreatsRecurringWeekly
	if o.Weekly != nil {
		obj, err := o.Weekly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		weeklyVal = obj
	}

	result := &UpdateScheduleThreatsRecurring{
		NewAppThreshold: o.NewAppThreshold,
		SyncToPeer:      util.AsBool(o.SyncToPeer, nil),
		Threshold:       o.Threshold,
		Daily:           dailyVal,
		Every30Mins:     every30MinsVal,
		Hourly:          hourlyVal,
		None:            noneVal,
		Weekly:          weeklyVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringDailyXml) MarshalFromObject(s UpdateScheduleThreatsRecurringDaily) {
	o.Action = s.Action
	o.At = s.At
	o.DisableNewContent = util.YesNo(s.DisableNewContent, nil)
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringDailyXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurringDaily, error) {

	result := &UpdateScheduleThreatsRecurringDaily{
		Action:            o.Action,
		At:                o.At,
		DisableNewContent: util.AsBool(o.DisableNewContent, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringEvery30MinsXml) MarshalFromObject(s UpdateScheduleThreatsRecurringEvery30Mins) {
	o.Action = s.Action
	o.At = s.At
	o.DisableNewContent = util.YesNo(s.DisableNewContent, nil)
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringEvery30MinsXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurringEvery30Mins, error) {

	result := &UpdateScheduleThreatsRecurringEvery30Mins{
		Action:            o.Action,
		At:                o.At,
		DisableNewContent: util.AsBool(o.DisableNewContent, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringHourlyXml) MarshalFromObject(s UpdateScheduleThreatsRecurringHourly) {
	o.Action = s.Action
	o.At = s.At
	o.DisableNewContent = util.YesNo(s.DisableNewContent, nil)
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringHourlyXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurringHourly, error) {

	result := &UpdateScheduleThreatsRecurringHourly{
		Action:            o.Action,
		At:                o.At,
		DisableNewContent: util.AsBool(o.DisableNewContent, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringNoneXml) MarshalFromObject(s UpdateScheduleThreatsRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurringNone, error) {

	result := &UpdateScheduleThreatsRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleThreatsRecurringWeeklyXml) MarshalFromObject(s UpdateScheduleThreatsRecurringWeekly) {
	o.Action = s.Action
	o.At = s.At
	o.DayOfWeek = s.DayOfWeek
	o.DisableNewContent = util.YesNo(s.DisableNewContent, nil)
	o.Misc = s.Misc
}

func (o updateScheduleThreatsRecurringWeeklyXml) UnmarshalToObject() (*UpdateScheduleThreatsRecurringWeekly, error) {

	result := &UpdateScheduleThreatsRecurringWeekly{
		Action:            o.Action,
		At:                o.At,
		DayOfWeek:         o.DayOfWeek,
		DisableNewContent: util.AsBool(o.DisableNewContent, nil),
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateXml) MarshalFromObject(s UpdateScheduleWfPrivate) {
	if s.Recurring != nil {
		var obj updateScheduleWfPrivateRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateXml) UnmarshalToObject() (*UpdateScheduleWfPrivate, error) {
	var recurringVal *UpdateScheduleWfPrivateRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleWfPrivate{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurring) {
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	if s.Every15Mins != nil {
		var obj updateScheduleWfPrivateRecurringEvery15MinsXml
		obj.MarshalFromObject(*s.Every15Mins)
		o.Every15Mins = &obj
	}
	if s.Every30Mins != nil {
		var obj updateScheduleWfPrivateRecurringEvery30MinsXml
		obj.MarshalFromObject(*s.Every30Mins)
		o.Every30Mins = &obj
	}
	if s.Every5Mins != nil {
		var obj updateScheduleWfPrivateRecurringEvery5MinsXml
		obj.MarshalFromObject(*s.Every5Mins)
		o.Every5Mins = &obj
	}
	if s.EveryHour != nil {
		var obj updateScheduleWfPrivateRecurringEveryHourXml
		obj.MarshalFromObject(*s.EveryHour)
		o.EveryHour = &obj
	}
	if s.None != nil {
		var obj updateScheduleWfPrivateRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurring, error) {
	var every15MinsVal *UpdateScheduleWfPrivateRecurringEvery15Mins
	if o.Every15Mins != nil {
		obj, err := o.Every15Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every15MinsVal = obj
	}
	var every30MinsVal *UpdateScheduleWfPrivateRecurringEvery30Mins
	if o.Every30Mins != nil {
		obj, err := o.Every30Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every30MinsVal = obj
	}
	var every5MinsVal *UpdateScheduleWfPrivateRecurringEvery5Mins
	if o.Every5Mins != nil {
		obj, err := o.Every5Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every5MinsVal = obj
	}
	var everyHourVal *UpdateScheduleWfPrivateRecurringEveryHour
	if o.EveryHour != nil {
		obj, err := o.EveryHour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		everyHourVal = obj
	}
	var noneVal *UpdateScheduleWfPrivateRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}

	result := &UpdateScheduleWfPrivateRecurring{
		SyncToPeer:  util.AsBool(o.SyncToPeer, nil),
		Every15Mins: every15MinsVal,
		Every30Mins: every30MinsVal,
		Every5Mins:  every5MinsVal,
		EveryHour:   everyHourVal,
		None:        noneVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringEvery15MinsXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurringEvery15Mins) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringEvery15MinsXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurringEvery15Mins, error) {

	result := &UpdateScheduleWfPrivateRecurringEvery15Mins{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringEvery30MinsXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurringEvery30Mins) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringEvery30MinsXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurringEvery30Mins, error) {

	result := &UpdateScheduleWfPrivateRecurringEvery30Mins{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringEvery5MinsXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurringEvery5Mins) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringEvery5MinsXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurringEvery5Mins, error) {

	result := &UpdateScheduleWfPrivateRecurringEvery5Mins{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringEveryHourXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurringEveryHour) {
	o.Action = s.Action
	o.At = s.At
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringEveryHourXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurringEveryHour, error) {

	result := &UpdateScheduleWfPrivateRecurringEveryHour{
		Action: o.Action,
		At:     o.At,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWfPrivateRecurringNoneXml) MarshalFromObject(s UpdateScheduleWfPrivateRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleWfPrivateRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleWfPrivateRecurringNone, error) {

	result := &UpdateScheduleWfPrivateRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireXml) MarshalFromObject(s UpdateScheduleWildfire) {
	if s.Recurring != nil {
		var obj updateScheduleWildfireRecurringXml
		obj.MarshalFromObject(*s.Recurring)
		o.Recurring = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleWildfireXml) UnmarshalToObject() (*UpdateScheduleWildfire, error) {
	var recurringVal *UpdateScheduleWildfireRecurring
	if o.Recurring != nil {
		obj, err := o.Recurring.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		recurringVal = obj
	}

	result := &UpdateScheduleWildfire{
		Recurring: recurringVal,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringXml) MarshalFromObject(s UpdateScheduleWildfireRecurring) {
	if s.Every15Mins != nil {
		var obj updateScheduleWildfireRecurringEvery15MinsXml
		obj.MarshalFromObject(*s.Every15Mins)
		o.Every15Mins = &obj
	}
	if s.Every30Mins != nil {
		var obj updateScheduleWildfireRecurringEvery30MinsXml
		obj.MarshalFromObject(*s.Every30Mins)
		o.Every30Mins = &obj
	}
	if s.EveryHour != nil {
		var obj updateScheduleWildfireRecurringEveryHourXml
		obj.MarshalFromObject(*s.EveryHour)
		o.EveryHour = &obj
	}
	if s.EveryMin != nil {
		var obj updateScheduleWildfireRecurringEveryMinXml
		obj.MarshalFromObject(*s.EveryMin)
		o.EveryMin = &obj
	}
	if s.None != nil {
		var obj updateScheduleWildfireRecurringNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.RealTime != nil {
		var obj updateScheduleWildfireRecurringRealTimeXml
		obj.MarshalFromObject(*s.RealTime)
		o.RealTime = &obj
	}
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurring, error) {
	var every15MinsVal *UpdateScheduleWildfireRecurringEvery15Mins
	if o.Every15Mins != nil {
		obj, err := o.Every15Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every15MinsVal = obj
	}
	var every30MinsVal *UpdateScheduleWildfireRecurringEvery30Mins
	if o.Every30Mins != nil {
		obj, err := o.Every30Mins.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		every30MinsVal = obj
	}
	var everyHourVal *UpdateScheduleWildfireRecurringEveryHour
	if o.EveryHour != nil {
		obj, err := o.EveryHour.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		everyHourVal = obj
	}
	var everyMinVal *UpdateScheduleWildfireRecurringEveryMin
	if o.EveryMin != nil {
		obj, err := o.EveryMin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		everyMinVal = obj
	}
	var noneVal *UpdateScheduleWildfireRecurringNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var realTimeVal *UpdateScheduleWildfireRecurringRealTime
	if o.RealTime != nil {
		obj, err := o.RealTime.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		realTimeVal = obj
	}

	result := &UpdateScheduleWildfireRecurring{
		Every15Mins: every15MinsVal,
		Every30Mins: every30MinsVal,
		EveryHour:   everyHourVal,
		EveryMin:    everyMinVal,
		None:        noneVal,
		RealTime:    realTimeVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringEvery15MinsXml) MarshalFromObject(s UpdateScheduleWildfireRecurringEvery15Mins) {
	o.Action = s.Action
	o.At = s.At
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringEvery15MinsXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringEvery15Mins, error) {

	result := &UpdateScheduleWildfireRecurringEvery15Mins{
		Action:     o.Action,
		At:         o.At,
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringEvery30MinsXml) MarshalFromObject(s UpdateScheduleWildfireRecurringEvery30Mins) {
	o.Action = s.Action
	o.At = s.At
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringEvery30MinsXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringEvery30Mins, error) {

	result := &UpdateScheduleWildfireRecurringEvery30Mins{
		Action:     o.Action,
		At:         o.At,
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringEveryHourXml) MarshalFromObject(s UpdateScheduleWildfireRecurringEveryHour) {
	o.Action = s.Action
	o.At = s.At
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringEveryHourXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringEveryHour, error) {

	result := &UpdateScheduleWildfireRecurringEveryHour{
		Action:     o.Action,
		At:         o.At,
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringEveryMinXml) MarshalFromObject(s UpdateScheduleWildfireRecurringEveryMin) {
	o.Action = s.Action
	o.SyncToPeer = util.YesNo(s.SyncToPeer, nil)
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringEveryMinXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringEveryMin, error) {

	result := &UpdateScheduleWildfireRecurringEveryMin{
		Action:     o.Action,
		SyncToPeer: util.AsBool(o.SyncToPeer, nil),
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringNoneXml) MarshalFromObject(s UpdateScheduleWildfireRecurringNone) {
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringNoneXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringNone, error) {

	result := &UpdateScheduleWildfireRecurringNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *updateScheduleWildfireRecurringRealTimeXml) MarshalFromObject(s UpdateScheduleWildfireRecurringRealTime) {
	o.Misc = s.Misc
}

func (o updateScheduleWildfireRecurringRealTimeXml) UnmarshalToObject() (*UpdateScheduleWildfireRecurringRealTime, error) {

	result := &UpdateScheduleWildfireRecurringRealTime{
		Misc: o.Misc,
	}
	return result, nil
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func SpecMatches(a, b *Config) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Config) matches(other *Config) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.UpdateSchedule.matches(other.UpdateSchedule) {
		return false
	}

	return true
}

func (o *UpdateSchedule) matches(other *UpdateSchedule) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AntiVirus.matches(other.AntiVirus) {
		return false
	}
	if !o.AppProfile.matches(other.AppProfile) {
		return false
	}
	if !o.GlobalProtectClientlessVpn.matches(other.GlobalProtectClientlessVpn) {
		return false
	}
	if !o.GlobalProtectDatafile.matches(other.GlobalProtectDatafile) {
		return false
	}
	if !o.StatisticsService.matches(other.StatisticsService) {
		return false
	}
	if !o.Threats.matches(other.Threats) {
		return false
	}
	if !o.WfPrivate.matches(other.WfPrivate) {
		return false
	}
	if !o.Wildfire.matches(other.Wildfire) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirus) matches(other *UpdateScheduleAntiVirus) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirusRecurring) matches(other *UpdateScheduleAntiVirusRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirusRecurringDaily) matches(other *UpdateScheduleAntiVirusRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirusRecurringHourly) matches(other *UpdateScheduleAntiVirusRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirusRecurringNone) matches(other *UpdateScheduleAntiVirusRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleAntiVirusRecurringWeekly) matches(other *UpdateScheduleAntiVirusRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
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

func (o *UpdateScheduleAppProfile) matches(other *UpdateScheduleAppProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleAppProfileRecurring) matches(other *UpdateScheduleAppProfileRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *UpdateScheduleAppProfileRecurringDaily) matches(other *UpdateScheduleAppProfileRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleAppProfileRecurringNone) matches(other *UpdateScheduleAppProfileRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleAppProfileRecurringWeekly) matches(other *UpdateScheduleAppProfileRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
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

func (o *UpdateScheduleGlobalProtectClientlessVpn) matches(other *UpdateScheduleGlobalProtectClientlessVpn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectClientlessVpnRecurring) matches(other *UpdateScheduleGlobalProtectClientlessVpnRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily) matches(other *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly) matches(other *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectClientlessVpnRecurringNone) matches(other *UpdateScheduleGlobalProtectClientlessVpnRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly) matches(other *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
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

func (o *UpdateScheduleGlobalProtectDatafile) matches(other *UpdateScheduleGlobalProtectDatafile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectDatafileRecurring) matches(other *UpdateScheduleGlobalProtectDatafileRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectDatafileRecurringDaily) matches(other *UpdateScheduleGlobalProtectDatafileRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectDatafileRecurringHourly) matches(other *UpdateScheduleGlobalProtectDatafileRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectDatafileRecurringNone) matches(other *UpdateScheduleGlobalProtectDatafileRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleGlobalProtectDatafileRecurringWeekly) matches(other *UpdateScheduleGlobalProtectDatafileRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
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

func (o *UpdateScheduleStatisticsService) matches(other *UpdateScheduleStatisticsService) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.ApplicationReports, other.ApplicationReports) {
		return false
	}
	if !util.BoolsMatch(o.FileIdentificationReports, other.FileIdentificationReports) {
		return false
	}
	if !util.BoolsMatch(o.HealthPerformanceReports, other.HealthPerformanceReports) {
		return false
	}
	if !util.BoolsMatch(o.PassiveDnsMonitoring, other.PassiveDnsMonitoring) {
		return false
	}
	if !util.BoolsMatch(o.ThreatPreventionInformation, other.ThreatPreventionInformation) {
		return false
	}
	if !util.BoolsMatch(o.ThreatPreventionPcap, other.ThreatPreventionPcap) {
		return false
	}
	if !util.BoolsMatch(o.ThreatPreventionReports, other.ThreatPreventionReports) {
		return false
	}
	if !util.BoolsMatch(o.UrlReports, other.UrlReports) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreats) matches(other *UpdateScheduleThreats) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurring) matches(other *UpdateScheduleThreatsRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.NewAppThreshold, other.NewAppThreshold) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}
	if !o.Daily.matches(other.Daily) {
		return false
	}
	if !o.Every30Mins.matches(other.Every30Mins) {
		return false
	}
	if !o.Hourly.matches(other.Hourly) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Weekly.matches(other.Weekly) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurringDaily) matches(other *UpdateScheduleThreatsRecurringDaily) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.DisableNewContent, other.DisableNewContent) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurringEvery30Mins) matches(other *UpdateScheduleThreatsRecurringEvery30Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.DisableNewContent, other.DisableNewContent) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurringHourly) matches(other *UpdateScheduleThreatsRecurringHourly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.DisableNewContent, other.DisableNewContent) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurringNone) matches(other *UpdateScheduleThreatsRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleThreatsRecurringWeekly) matches(other *UpdateScheduleThreatsRecurringWeekly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.At, other.At) {
		return false
	}
	if !util.StringsMatch(o.DayOfWeek, other.DayOfWeek) {
		return false
	}
	if !util.BoolsMatch(o.DisableNewContent, other.DisableNewContent) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivate) matches(other *UpdateScheduleWfPrivate) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurring) matches(other *UpdateScheduleWfPrivateRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}
	if !o.Every15Mins.matches(other.Every15Mins) {
		return false
	}
	if !o.Every30Mins.matches(other.Every30Mins) {
		return false
	}
	if !o.Every5Mins.matches(other.Every5Mins) {
		return false
	}
	if !o.EveryHour.matches(other.EveryHour) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurringEvery15Mins) matches(other *UpdateScheduleWfPrivateRecurringEvery15Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurringEvery30Mins) matches(other *UpdateScheduleWfPrivateRecurringEvery30Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurringEvery5Mins) matches(other *UpdateScheduleWfPrivateRecurringEvery5Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurringEveryHour) matches(other *UpdateScheduleWfPrivateRecurringEveryHour) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}

	return true
}

func (o *UpdateScheduleWfPrivateRecurringNone) matches(other *UpdateScheduleWfPrivateRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfire) matches(other *UpdateScheduleWildfire) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Recurring.matches(other.Recurring) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurring) matches(other *UpdateScheduleWildfireRecurring) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Every15Mins.matches(other.Every15Mins) {
		return false
	}
	if !o.Every30Mins.matches(other.Every30Mins) {
		return false
	}
	if !o.EveryHour.matches(other.EveryHour) {
		return false
	}
	if !o.EveryMin.matches(other.EveryMin) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.RealTime.matches(other.RealTime) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringEvery15Mins) matches(other *UpdateScheduleWildfireRecurringEvery15Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringEvery30Mins) matches(other *UpdateScheduleWildfireRecurringEvery30Mins) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringEveryHour) matches(other *UpdateScheduleWildfireRecurringEveryHour) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.Ints64Match(o.At, other.At) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringEveryMin) matches(other *UpdateScheduleWildfireRecurringEveryMin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.BoolsMatch(o.SyncToPeer, other.SyncToPeer) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringNone) matches(other *UpdateScheduleWildfireRecurringNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *UpdateScheduleWildfireRecurringRealTime) matches(other *UpdateScheduleWildfireRecurringRealTime) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}
