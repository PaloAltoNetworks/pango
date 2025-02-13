package dynamicupdates

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	UpdateSchedule *UpdateSchedule

	Misc map[string][]generic.Xml
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
}
type UpdateScheduleAntiVirus struct {
	Recurring *UpdateScheduleAntiVirusRecurring
}
type UpdateScheduleAntiVirusRecurring struct {
	SyncToPeer *bool
	Threshold  *int64
	Daily      *UpdateScheduleAntiVirusRecurringDaily
	Hourly     *UpdateScheduleAntiVirusRecurringHourly
	None       *UpdateScheduleAntiVirusRecurringNone
	Weekly     *UpdateScheduleAntiVirusRecurringWeekly
}
type UpdateScheduleAntiVirusRecurringDaily struct {
	Action *string
	At     *string
}
type UpdateScheduleAntiVirusRecurringHourly struct {
	Action *string
	At     *int64
}
type UpdateScheduleAntiVirusRecurringNone struct {
}
type UpdateScheduleAntiVirusRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
}
type UpdateScheduleAppProfile struct {
	Recurring *UpdateScheduleAppProfileRecurring
}
type UpdateScheduleAppProfileRecurring struct {
	SyncToPeer *bool
	Threshold  *int64
	Daily      *UpdateScheduleAppProfileRecurringDaily
	None       *UpdateScheduleAppProfileRecurringNone
	Weekly     *UpdateScheduleAppProfileRecurringWeekly
}
type UpdateScheduleAppProfileRecurringDaily struct {
	Action *string
	At     *string
}
type UpdateScheduleAppProfileRecurringNone struct {
}
type UpdateScheduleAppProfileRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
}
type UpdateScheduleGlobalProtectClientlessVpn struct {
	Recurring *UpdateScheduleGlobalProtectClientlessVpnRecurring
}
type UpdateScheduleGlobalProtectClientlessVpnRecurring struct {
	Daily  *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily
	Hourly *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly
	None   *UpdateScheduleGlobalProtectClientlessVpnRecurringNone
	Weekly *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringDaily struct {
	Action *string
	At     *string
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringHourly struct {
	Action *string
	At     *int64
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringNone struct {
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
}
type UpdateScheduleGlobalProtectDatafile struct {
	Recurring *UpdateScheduleGlobalProtectDatafileRecurring
}
type UpdateScheduleGlobalProtectDatafileRecurring struct {
	Daily  *UpdateScheduleGlobalProtectDatafileRecurringDaily
	Hourly *UpdateScheduleGlobalProtectDatafileRecurringHourly
	None   *UpdateScheduleGlobalProtectDatafileRecurringNone
	Weekly *UpdateScheduleGlobalProtectDatafileRecurringWeekly
}
type UpdateScheduleGlobalProtectDatafileRecurringDaily struct {
	Action *string
	At     *string
}
type UpdateScheduleGlobalProtectDatafileRecurringHourly struct {
	Action *string
	At     *int64
}
type UpdateScheduleGlobalProtectDatafileRecurringNone struct {
}
type UpdateScheduleGlobalProtectDatafileRecurringWeekly struct {
	Action    *string
	At        *string
	DayOfWeek *string
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
}
type UpdateScheduleThreats struct {
	Recurring *UpdateScheduleThreatsRecurring
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
}
type UpdateScheduleThreatsRecurringDaily struct {
	Action            *string
	At                *string
	DisableNewContent *bool
}
type UpdateScheduleThreatsRecurringEvery30Mins struct {
	Action            *string
	At                *int64
	DisableNewContent *bool
}
type UpdateScheduleThreatsRecurringHourly struct {
	Action            *string
	At                *int64
	DisableNewContent *bool
}
type UpdateScheduleThreatsRecurringNone struct {
}
type UpdateScheduleThreatsRecurringWeekly struct {
	Action            *string
	At                *string
	DayOfWeek         *string
	DisableNewContent *bool
}
type UpdateScheduleWfPrivate struct {
	Recurring *UpdateScheduleWfPrivateRecurring
}
type UpdateScheduleWfPrivateRecurring struct {
	SyncToPeer  *bool
	Every15Mins *UpdateScheduleWfPrivateRecurringEvery15Mins
	Every30Mins *UpdateScheduleWfPrivateRecurringEvery30Mins
	Every5Mins  *UpdateScheduleWfPrivateRecurringEvery5Mins
	EveryHour   *UpdateScheduleWfPrivateRecurringEveryHour
	None        *UpdateScheduleWfPrivateRecurringNone
}
type UpdateScheduleWfPrivateRecurringEvery15Mins struct {
	Action *string
	At     *int64
}
type UpdateScheduleWfPrivateRecurringEvery30Mins struct {
	Action *string
	At     *int64
}
type UpdateScheduleWfPrivateRecurringEvery5Mins struct {
	Action *string
	At     *int64
}
type UpdateScheduleWfPrivateRecurringEveryHour struct {
	Action *string
	At     *int64
}
type UpdateScheduleWfPrivateRecurringNone struct {
}
type UpdateScheduleWildfire struct {
	Recurring *UpdateScheduleWildfireRecurring
}
type UpdateScheduleWildfireRecurring struct {
	Every15Mins *UpdateScheduleWildfireRecurringEvery15Mins
	Every30Mins *UpdateScheduleWildfireRecurringEvery30Mins
	EveryHour   *UpdateScheduleWildfireRecurringEveryHour
	EveryMin    *UpdateScheduleWildfireRecurringEveryMin
	None        *UpdateScheduleWildfireRecurringNone
	RealTime    *UpdateScheduleWildfireRecurringRealTime
}
type UpdateScheduleWildfireRecurringEvery15Mins struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
}
type UpdateScheduleWildfireRecurringEvery30Mins struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
}
type UpdateScheduleWildfireRecurringEveryHour struct {
	Action     *string
	At         *int64
	SyncToPeer *bool
}
type UpdateScheduleWildfireRecurringEveryMin struct {
	Action     *string
	SyncToPeer *bool
}
type UpdateScheduleWildfireRecurringNone struct {
}
type UpdateScheduleWildfireRecurringRealTime struct {
}
type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}
type configXml struct {
	XMLName        xml.Name           `xml:"system"`
	UpdateSchedule *UpdateScheduleXml `xml:"update-schedule,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleXml struct {
	AntiVirus                  *UpdateScheduleAntiVirusXml                  `xml:"anti-virus,omitempty"`
	AppProfile                 *UpdateScheduleAppProfileXml                 `xml:"app-profile,omitempty"`
	GlobalProtectClientlessVpn *UpdateScheduleGlobalProtectClientlessVpnXml `xml:"global-protect-clientless-vpn,omitempty"`
	GlobalProtectDatafile      *UpdateScheduleGlobalProtectDatafileXml      `xml:"global-protect-datafile,omitempty"`
	StatisticsService          *UpdateScheduleStatisticsServiceXml          `xml:"statistics-service,omitempty"`
	Threats                    *UpdateScheduleThreatsXml                    `xml:"threats,omitempty"`
	WfPrivate                  *UpdateScheduleWfPrivateXml                  `xml:"wf-private,omitempty"`
	Wildfire                   *UpdateScheduleWildfireXml                   `xml:"wildfire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusXml struct {
	Recurring *UpdateScheduleAntiVirusRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusRecurringXml struct {
	SyncToPeer *string                                    `xml:"sync-to-peer,omitempty"`
	Threshold  *int64                                     `xml:"threshold,omitempty"`
	Daily      *UpdateScheduleAntiVirusRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly     *UpdateScheduleAntiVirusRecurringHourlyXml `xml:"hourly,omitempty"`
	None       *UpdateScheduleAntiVirusRecurringNoneXml   `xml:"none,omitempty"`
	Weekly     *UpdateScheduleAntiVirusRecurringWeeklyXml `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusRecurringDailyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusRecurringHourlyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAntiVirusRecurringWeeklyXml struct {
	Action    *string `xml:"action,omitempty"`
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAppProfileXml struct {
	Recurring *UpdateScheduleAppProfileRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAppProfileRecurringXml struct {
	SyncToPeer *string                                     `xml:"sync-to-peer,omitempty"`
	Threshold  *int64                                      `xml:"threshold,omitempty"`
	Daily      *UpdateScheduleAppProfileRecurringDailyXml  `xml:"daily,omitempty"`
	None       *UpdateScheduleAppProfileRecurringNoneXml   `xml:"none,omitempty"`
	Weekly     *UpdateScheduleAppProfileRecurringWeeklyXml `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAppProfileRecurringDailyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAppProfileRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleAppProfileRecurringWeeklyXml struct {
	Action    *string `xml:"action,omitempty"`
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnXml struct {
	Recurring *UpdateScheduleGlobalProtectClientlessVpnRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringXml struct {
	Daily  *UpdateScheduleGlobalProtectClientlessVpnRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly *UpdateScheduleGlobalProtectClientlessVpnRecurringHourlyXml `xml:"hourly,omitempty"`
	None   *UpdateScheduleGlobalProtectClientlessVpnRecurringNoneXml   `xml:"none,omitempty"`
	Weekly *UpdateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringDailyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringHourlyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml struct {
	Action    *string `xml:"action,omitempty"`
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileXml struct {
	Recurring *UpdateScheduleGlobalProtectDatafileRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileRecurringXml struct {
	Daily  *UpdateScheduleGlobalProtectDatafileRecurringDailyXml  `xml:"daily,omitempty"`
	Hourly *UpdateScheduleGlobalProtectDatafileRecurringHourlyXml `xml:"hourly,omitempty"`
	None   *UpdateScheduleGlobalProtectDatafileRecurringNoneXml   `xml:"none,omitempty"`
	Weekly *UpdateScheduleGlobalProtectDatafileRecurringWeeklyXml `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileRecurringDailyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *string `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileRecurringHourlyXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleGlobalProtectDatafileRecurringWeeklyXml struct {
	Action    *string `xml:"action,omitempty"`
	At        *string `xml:"at,omitempty"`
	DayOfWeek *string `xml:"day-of-week,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleStatisticsServiceXml struct {
	ApplicationReports          *string `xml:"application-reports,omitempty"`
	FileIdentificationReports   *string `xml:"file-identification-reports,omitempty"`
	HealthPerformanceReports    *string `xml:"health-performance-reports,omitempty"`
	PassiveDnsMonitoring        *string `xml:"passive-dns-monitoring,omitempty"`
	ThreatPreventionInformation *string `xml:"threat-prevention-information,omitempty"`
	ThreatPreventionPcap        *string `xml:"threat-prevention-pcap,omitempty"`
	ThreatPreventionReports     *string `xml:"threat-prevention-reports,omitempty"`
	UrlReports                  *string `xml:"url-reports,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsXml struct {
	Recurring *UpdateScheduleThreatsRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringXml struct {
	NewAppThreshold *int64                                        `xml:"new-app-threshold,omitempty"`
	SyncToPeer      *string                                       `xml:"sync-to-peer,omitempty"`
	Threshold       *int64                                        `xml:"threshold,omitempty"`
	Daily           *UpdateScheduleThreatsRecurringDailyXml       `xml:"daily,omitempty"`
	Every30Mins     *UpdateScheduleThreatsRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	Hourly          *UpdateScheduleThreatsRecurringHourlyXml      `xml:"hourly,omitempty"`
	None            *UpdateScheduleThreatsRecurringNoneXml        `xml:"none,omitempty"`
	Weekly          *UpdateScheduleThreatsRecurringWeeklyXml      `xml:"weekly,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringDailyXml struct {
	Action            *string `xml:"action,omitempty"`
	At                *string `xml:"at,omitempty"`
	DisableNewContent *string `xml:"disable-new-content,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringEvery30MinsXml struct {
	Action            *string `xml:"action,omitempty"`
	At                *int64  `xml:"at,omitempty"`
	DisableNewContent *string `xml:"disable-new-content,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringHourlyXml struct {
	Action            *string `xml:"action,omitempty"`
	At                *int64  `xml:"at,omitempty"`
	DisableNewContent *string `xml:"disable-new-content,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleThreatsRecurringWeeklyXml struct {
	Action            *string `xml:"action,omitempty"`
	At                *string `xml:"at,omitempty"`
	DayOfWeek         *string `xml:"day-of-week,omitempty"`
	DisableNewContent *string `xml:"disable-new-content,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateXml struct {
	Recurring *UpdateScheduleWfPrivateRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringXml struct {
	SyncToPeer  *string                                         `xml:"sync-to-peer,omitempty"`
	Every15Mins *UpdateScheduleWfPrivateRecurringEvery15MinsXml `xml:"every-15-mins,omitempty"`
	Every30Mins *UpdateScheduleWfPrivateRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	Every5Mins  *UpdateScheduleWfPrivateRecurringEvery5MinsXml  `xml:"every-5-mins,omitempty"`
	EveryHour   *UpdateScheduleWfPrivateRecurringEveryHourXml   `xml:"every-hour,omitempty"`
	None        *UpdateScheduleWfPrivateRecurringNoneXml        `xml:"none,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringEvery15MinsXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringEvery30MinsXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringEvery5MinsXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringEveryHourXml struct {
	Action *string `xml:"action,omitempty"`
	At     *int64  `xml:"at,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWfPrivateRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireXml struct {
	Recurring *UpdateScheduleWildfireRecurringXml `xml:"recurring,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringXml struct {
	Every15Mins *UpdateScheduleWildfireRecurringEvery15MinsXml `xml:"every-15-mins,omitempty"`
	Every30Mins *UpdateScheduleWildfireRecurringEvery30MinsXml `xml:"every-30-mins,omitempty"`
	EveryHour   *UpdateScheduleWildfireRecurringEveryHourXml   `xml:"every-hour,omitempty"`
	EveryMin    *UpdateScheduleWildfireRecurringEveryMinXml    `xml:"every-min,omitempty"`
	None        *UpdateScheduleWildfireRecurringNoneXml        `xml:"none,omitempty"`
	RealTime    *UpdateScheduleWildfireRecurringRealTimeXml    `xml:"real-time,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringEvery15MinsXml struct {
	Action     *string `xml:"action,omitempty"`
	At         *int64  `xml:"at,omitempty"`
	SyncToPeer *string `xml:"sync-to-peer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringEvery30MinsXml struct {
	Action     *string `xml:"action,omitempty"`
	At         *int64  `xml:"at,omitempty"`
	SyncToPeer *string `xml:"sync-to-peer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringEveryHourXml struct {
	Action     *string `xml:"action,omitempty"`
	At         *int64  `xml:"at,omitempty"`
	SyncToPeer *string `xml:"sync-to-peer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringEveryMinXml struct {
	Action     *string `xml:"action,omitempty"`
	SyncToPeer *string `xml:"sync-to-peer,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type UpdateScheduleWildfireRecurringRealTimeXml struct {
	Misc []generic.Xml `xml:",any"`
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func specifyConfig(o *Config) (any, error) {
	config := configXml{}
	var nestedUpdateSchedule *UpdateScheduleXml
	if o.UpdateSchedule != nil {
		nestedUpdateSchedule = &UpdateScheduleXml{}
		if _, ok := o.Misc["UpdateSchedule"]; ok {
			nestedUpdateSchedule.Misc = o.Misc["UpdateSchedule"]
		}
		if o.UpdateSchedule.AntiVirus != nil {
			nestedUpdateSchedule.AntiVirus = &UpdateScheduleAntiVirusXml{}
			if _, ok := o.Misc["UpdateScheduleAntiVirus"]; ok {
				nestedUpdateSchedule.AntiVirus.Misc = o.Misc["UpdateScheduleAntiVirus"]
			}
			if o.UpdateSchedule.AntiVirus.Recurring != nil {
				nestedUpdateSchedule.AntiVirus.Recurring = &UpdateScheduleAntiVirusRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleAntiVirusRecurring"]; ok {
					nestedUpdateSchedule.AntiVirus.Recurring.Misc = o.Misc["UpdateScheduleAntiVirusRecurring"]
				}
				if o.UpdateSchedule.AntiVirus.Recurring.SyncToPeer != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.SyncToPeer = util.YesNo(o.UpdateSchedule.AntiVirus.Recurring.SyncToPeer, nil)
				}
				if o.UpdateSchedule.AntiVirus.Recurring.Threshold != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.Threshold = o.UpdateSchedule.AntiVirus.Recurring.Threshold
				}
				if o.UpdateSchedule.AntiVirus.Recurring.Weekly != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.Weekly = &UpdateScheduleAntiVirusRecurringWeeklyXml{}
					if _, ok := o.Misc["UpdateScheduleAntiVirusRecurringWeekly"]; ok {
						nestedUpdateSchedule.AntiVirus.Recurring.Weekly.Misc = o.Misc["UpdateScheduleAntiVirusRecurringWeekly"]
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Weekly.Action != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Weekly.Action = o.UpdateSchedule.AntiVirus.Recurring.Weekly.Action
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Weekly.At != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Weekly.At = o.UpdateSchedule.AntiVirus.Recurring.Weekly.At
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek
					}
				}
				if o.UpdateSchedule.AntiVirus.Recurring.Daily != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.Daily = &UpdateScheduleAntiVirusRecurringDailyXml{}
					if _, ok := o.Misc["UpdateScheduleAntiVirusRecurringDaily"]; ok {
						nestedUpdateSchedule.AntiVirus.Recurring.Daily.Misc = o.Misc["UpdateScheduleAntiVirusRecurringDaily"]
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Daily.Action != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Daily.Action = o.UpdateSchedule.AntiVirus.Recurring.Daily.Action
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Daily.At != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Daily.At = o.UpdateSchedule.AntiVirus.Recurring.Daily.At
					}
				}
				if o.UpdateSchedule.AntiVirus.Recurring.Hourly != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.Hourly = &UpdateScheduleAntiVirusRecurringHourlyXml{}
					if _, ok := o.Misc["UpdateScheduleAntiVirusRecurringHourly"]; ok {
						nestedUpdateSchedule.AntiVirus.Recurring.Hourly.Misc = o.Misc["UpdateScheduleAntiVirusRecurringHourly"]
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Hourly.Action != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Hourly.Action = o.UpdateSchedule.AntiVirus.Recurring.Hourly.Action
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Hourly.At != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Hourly.At = o.UpdateSchedule.AntiVirus.Recurring.Hourly.At
					}
				}
				if o.UpdateSchedule.AntiVirus.Recurring.None != nil {
					nestedUpdateSchedule.AntiVirus.Recurring.None = &UpdateScheduleAntiVirusRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleAntiVirusRecurringNone"]; ok {
						nestedUpdateSchedule.AntiVirus.Recurring.None.Misc = o.Misc["UpdateScheduleAntiVirusRecurringNone"]
					}
				}
			}
		}
		if o.UpdateSchedule.AppProfile != nil {
			nestedUpdateSchedule.AppProfile = &UpdateScheduleAppProfileXml{}
			if _, ok := o.Misc["UpdateScheduleAppProfile"]; ok {
				nestedUpdateSchedule.AppProfile.Misc = o.Misc["UpdateScheduleAppProfile"]
			}
			if o.UpdateSchedule.AppProfile.Recurring != nil {
				nestedUpdateSchedule.AppProfile.Recurring = &UpdateScheduleAppProfileRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleAppProfileRecurring"]; ok {
					nestedUpdateSchedule.AppProfile.Recurring.Misc = o.Misc["UpdateScheduleAppProfileRecurring"]
				}
				if o.UpdateSchedule.AppProfile.Recurring.SyncToPeer != nil {
					nestedUpdateSchedule.AppProfile.Recurring.SyncToPeer = util.YesNo(o.UpdateSchedule.AppProfile.Recurring.SyncToPeer, nil)
				}
				if o.UpdateSchedule.AppProfile.Recurring.Threshold != nil {
					nestedUpdateSchedule.AppProfile.Recurring.Threshold = o.UpdateSchedule.AppProfile.Recurring.Threshold
				}
				if o.UpdateSchedule.AppProfile.Recurring.Daily != nil {
					nestedUpdateSchedule.AppProfile.Recurring.Daily = &UpdateScheduleAppProfileRecurringDailyXml{}
					if _, ok := o.Misc["UpdateScheduleAppProfileRecurringDaily"]; ok {
						nestedUpdateSchedule.AppProfile.Recurring.Daily.Misc = o.Misc["UpdateScheduleAppProfileRecurringDaily"]
					}
					if o.UpdateSchedule.AppProfile.Recurring.Daily.Action != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Daily.Action = o.UpdateSchedule.AppProfile.Recurring.Daily.Action
					}
					if o.UpdateSchedule.AppProfile.Recurring.Daily.At != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Daily.At = o.UpdateSchedule.AppProfile.Recurring.Daily.At
					}
				}
				if o.UpdateSchedule.AppProfile.Recurring.None != nil {
					nestedUpdateSchedule.AppProfile.Recurring.None = &UpdateScheduleAppProfileRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleAppProfileRecurringNone"]; ok {
						nestedUpdateSchedule.AppProfile.Recurring.None.Misc = o.Misc["UpdateScheduleAppProfileRecurringNone"]
					}
				}
				if o.UpdateSchedule.AppProfile.Recurring.Weekly != nil {
					nestedUpdateSchedule.AppProfile.Recurring.Weekly = &UpdateScheduleAppProfileRecurringWeeklyXml{}
					if _, ok := o.Misc["UpdateScheduleAppProfileRecurringWeekly"]; ok {
						nestedUpdateSchedule.AppProfile.Recurring.Weekly.Misc = o.Misc["UpdateScheduleAppProfileRecurringWeekly"]
					}
					if o.UpdateSchedule.AppProfile.Recurring.Weekly.Action != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Weekly.Action = o.UpdateSchedule.AppProfile.Recurring.Weekly.Action
					}
					if o.UpdateSchedule.AppProfile.Recurring.Weekly.At != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Weekly.At = o.UpdateSchedule.AppProfile.Recurring.Weekly.At
					}
					if o.UpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek
					}
				}
			}
		}
		if o.UpdateSchedule.GlobalProtectClientlessVpn != nil {
			nestedUpdateSchedule.GlobalProtectClientlessVpn = &UpdateScheduleGlobalProtectClientlessVpnXml{}
			if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpn"]; ok {
				nestedUpdateSchedule.GlobalProtectClientlessVpn.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpn"]
			}
			if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring != nil {
				nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring = &UpdateScheduleGlobalProtectClientlessVpnRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurring"]; ok {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurring"]
				}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly != nil {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly = &UpdateScheduleGlobalProtectClientlessVpnRecurringWeeklyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly"]; ok {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly"]
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek
					}
				}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily != nil {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily = &UpdateScheduleGlobalProtectClientlessVpnRecurringDailyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringDaily"]; ok {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringDaily"]
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At
					}
				}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly != nil {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly = &UpdateScheduleGlobalProtectClientlessVpnRecurringHourlyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringHourly"]; ok {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringHourly"]
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At
					}
				}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.None != nil {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.None = &UpdateScheduleGlobalProtectClientlessVpnRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringNone"]; ok {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.None.Misc = o.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringNone"]
					}
				}
			}
		}
		if o.UpdateSchedule.GlobalProtectDatafile != nil {
			nestedUpdateSchedule.GlobalProtectDatafile = &UpdateScheduleGlobalProtectDatafileXml{}
			if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafile"]; ok {
				nestedUpdateSchedule.GlobalProtectDatafile.Misc = o.Misc["UpdateScheduleGlobalProtectDatafile"]
			}
			if o.UpdateSchedule.GlobalProtectDatafile.Recurring != nil {
				nestedUpdateSchedule.GlobalProtectDatafile.Recurring = &UpdateScheduleGlobalProtectDatafileRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafileRecurring"]; ok {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Misc = o.Misc["UpdateScheduleGlobalProtectDatafileRecurring"]
				}
				if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily != nil {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily = &UpdateScheduleGlobalProtectDatafileRecurringDailyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafileRecurringDaily"]; ok {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Misc = o.Misc["UpdateScheduleGlobalProtectDatafileRecurringDaily"]
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At
					}
				}
				if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly != nil {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly = &UpdateScheduleGlobalProtectDatafileRecurringHourlyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafileRecurringHourly"]; ok {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Misc = o.Misc["UpdateScheduleGlobalProtectDatafileRecurringHourly"]
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At
					}
				}
				if o.UpdateSchedule.GlobalProtectDatafile.Recurring.None != nil {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring.None = &UpdateScheduleGlobalProtectDatafileRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafileRecurringNone"]; ok {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.None.Misc = o.Misc["UpdateScheduleGlobalProtectDatafileRecurringNone"]
					}
				}
				if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly != nil {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly = &UpdateScheduleGlobalProtectDatafileRecurringWeeklyXml{}
					if _, ok := o.Misc["UpdateScheduleGlobalProtectDatafileRecurringWeekly"]; ok {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Misc = o.Misc["UpdateScheduleGlobalProtectDatafileRecurringWeekly"]
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek
					}
				}
			}
		}
		if o.UpdateSchedule.StatisticsService != nil {
			nestedUpdateSchedule.StatisticsService = &UpdateScheduleStatisticsServiceXml{}
			if _, ok := o.Misc["UpdateScheduleStatisticsService"]; ok {
				nestedUpdateSchedule.StatisticsService.Misc = o.Misc["UpdateScheduleStatisticsService"]
			}
			if o.UpdateSchedule.StatisticsService.ThreatPreventionPcap != nil {
				nestedUpdateSchedule.StatisticsService.ThreatPreventionPcap = util.YesNo(o.UpdateSchedule.StatisticsService.ThreatPreventionPcap, nil)
			}
			if o.UpdateSchedule.StatisticsService.ThreatPreventionReports != nil {
				nestedUpdateSchedule.StatisticsService.ThreatPreventionReports = util.YesNo(o.UpdateSchedule.StatisticsService.ThreatPreventionReports, nil)
			}
			if o.UpdateSchedule.StatisticsService.UrlReports != nil {
				nestedUpdateSchedule.StatisticsService.UrlReports = util.YesNo(o.UpdateSchedule.StatisticsService.UrlReports, nil)
			}
			if o.UpdateSchedule.StatisticsService.ApplicationReports != nil {
				nestedUpdateSchedule.StatisticsService.ApplicationReports = util.YesNo(o.UpdateSchedule.StatisticsService.ApplicationReports, nil)
			}
			if o.UpdateSchedule.StatisticsService.FileIdentificationReports != nil {
				nestedUpdateSchedule.StatisticsService.FileIdentificationReports = util.YesNo(o.UpdateSchedule.StatisticsService.FileIdentificationReports, nil)
			}
			if o.UpdateSchedule.StatisticsService.HealthPerformanceReports != nil {
				nestedUpdateSchedule.StatisticsService.HealthPerformanceReports = util.YesNo(o.UpdateSchedule.StatisticsService.HealthPerformanceReports, nil)
			}
			if o.UpdateSchedule.StatisticsService.PassiveDnsMonitoring != nil {
				nestedUpdateSchedule.StatisticsService.PassiveDnsMonitoring = util.YesNo(o.UpdateSchedule.StatisticsService.PassiveDnsMonitoring, nil)
			}
			if o.UpdateSchedule.StatisticsService.ThreatPreventionInformation != nil {
				nestedUpdateSchedule.StatisticsService.ThreatPreventionInformation = util.YesNo(o.UpdateSchedule.StatisticsService.ThreatPreventionInformation, nil)
			}
		}
		if o.UpdateSchedule.Threats != nil {
			nestedUpdateSchedule.Threats = &UpdateScheduleThreatsXml{}
			if _, ok := o.Misc["UpdateScheduleThreats"]; ok {
				nestedUpdateSchedule.Threats.Misc = o.Misc["UpdateScheduleThreats"]
			}
			if o.UpdateSchedule.Threats.Recurring != nil {
				nestedUpdateSchedule.Threats.Recurring = &UpdateScheduleThreatsRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleThreatsRecurring"]; ok {
					nestedUpdateSchedule.Threats.Recurring.Misc = o.Misc["UpdateScheduleThreatsRecurring"]
				}
				if o.UpdateSchedule.Threats.Recurring.NewAppThreshold != nil {
					nestedUpdateSchedule.Threats.Recurring.NewAppThreshold = o.UpdateSchedule.Threats.Recurring.NewAppThreshold
				}
				if o.UpdateSchedule.Threats.Recurring.SyncToPeer != nil {
					nestedUpdateSchedule.Threats.Recurring.SyncToPeer = util.YesNo(o.UpdateSchedule.Threats.Recurring.SyncToPeer, nil)
				}
				if o.UpdateSchedule.Threats.Recurring.Threshold != nil {
					nestedUpdateSchedule.Threats.Recurring.Threshold = o.UpdateSchedule.Threats.Recurring.Threshold
				}
				if o.UpdateSchedule.Threats.Recurring.Daily != nil {
					nestedUpdateSchedule.Threats.Recurring.Daily = &UpdateScheduleThreatsRecurringDailyXml{}
					if _, ok := o.Misc["UpdateScheduleThreatsRecurringDaily"]; ok {
						nestedUpdateSchedule.Threats.Recurring.Daily.Misc = o.Misc["UpdateScheduleThreatsRecurringDaily"]
					}
					if o.UpdateSchedule.Threats.Recurring.Daily.Action != nil {
						nestedUpdateSchedule.Threats.Recurring.Daily.Action = o.UpdateSchedule.Threats.Recurring.Daily.Action
					}
					if o.UpdateSchedule.Threats.Recurring.Daily.At != nil {
						nestedUpdateSchedule.Threats.Recurring.Daily.At = o.UpdateSchedule.Threats.Recurring.Daily.At
					}
					if o.UpdateSchedule.Threats.Recurring.Daily.DisableNewContent != nil {
						nestedUpdateSchedule.Threats.Recurring.Daily.DisableNewContent = util.YesNo(o.UpdateSchedule.Threats.Recurring.Daily.DisableNewContent, nil)
					}
				}
				if o.UpdateSchedule.Threats.Recurring.Every30Mins != nil {
					nestedUpdateSchedule.Threats.Recurring.Every30Mins = &UpdateScheduleThreatsRecurringEvery30MinsXml{}
					if _, ok := o.Misc["UpdateScheduleThreatsRecurringEvery30Mins"]; ok {
						nestedUpdateSchedule.Threats.Recurring.Every30Mins.Misc = o.Misc["UpdateScheduleThreatsRecurringEvery30Mins"]
					}
					if o.UpdateSchedule.Threats.Recurring.Every30Mins.Action != nil {
						nestedUpdateSchedule.Threats.Recurring.Every30Mins.Action = o.UpdateSchedule.Threats.Recurring.Every30Mins.Action
					}
					if o.UpdateSchedule.Threats.Recurring.Every30Mins.At != nil {
						nestedUpdateSchedule.Threats.Recurring.Every30Mins.At = o.UpdateSchedule.Threats.Recurring.Every30Mins.At
					}
					if o.UpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent != nil {
						nestedUpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent = util.YesNo(o.UpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent, nil)
					}
				}
				if o.UpdateSchedule.Threats.Recurring.Hourly != nil {
					nestedUpdateSchedule.Threats.Recurring.Hourly = &UpdateScheduleThreatsRecurringHourlyXml{}
					if _, ok := o.Misc["UpdateScheduleThreatsRecurringHourly"]; ok {
						nestedUpdateSchedule.Threats.Recurring.Hourly.Misc = o.Misc["UpdateScheduleThreatsRecurringHourly"]
					}
					if o.UpdateSchedule.Threats.Recurring.Hourly.Action != nil {
						nestedUpdateSchedule.Threats.Recurring.Hourly.Action = o.UpdateSchedule.Threats.Recurring.Hourly.Action
					}
					if o.UpdateSchedule.Threats.Recurring.Hourly.At != nil {
						nestedUpdateSchedule.Threats.Recurring.Hourly.At = o.UpdateSchedule.Threats.Recurring.Hourly.At
					}
					if o.UpdateSchedule.Threats.Recurring.Hourly.DisableNewContent != nil {
						nestedUpdateSchedule.Threats.Recurring.Hourly.DisableNewContent = util.YesNo(o.UpdateSchedule.Threats.Recurring.Hourly.DisableNewContent, nil)
					}
				}
				if o.UpdateSchedule.Threats.Recurring.None != nil {
					nestedUpdateSchedule.Threats.Recurring.None = &UpdateScheduleThreatsRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleThreatsRecurringNone"]; ok {
						nestedUpdateSchedule.Threats.Recurring.None.Misc = o.Misc["UpdateScheduleThreatsRecurringNone"]
					}
				}
				if o.UpdateSchedule.Threats.Recurring.Weekly != nil {
					nestedUpdateSchedule.Threats.Recurring.Weekly = &UpdateScheduleThreatsRecurringWeeklyXml{}
					if _, ok := o.Misc["UpdateScheduleThreatsRecurringWeekly"]; ok {
						nestedUpdateSchedule.Threats.Recurring.Weekly.Misc = o.Misc["UpdateScheduleThreatsRecurringWeekly"]
					}
					if o.UpdateSchedule.Threats.Recurring.Weekly.At != nil {
						nestedUpdateSchedule.Threats.Recurring.Weekly.At = o.UpdateSchedule.Threats.Recurring.Weekly.At
					}
					if o.UpdateSchedule.Threats.Recurring.Weekly.DayOfWeek != nil {
						nestedUpdateSchedule.Threats.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.Threats.Recurring.Weekly.DayOfWeek
					}
					if o.UpdateSchedule.Threats.Recurring.Weekly.DisableNewContent != nil {
						nestedUpdateSchedule.Threats.Recurring.Weekly.DisableNewContent = util.YesNo(o.UpdateSchedule.Threats.Recurring.Weekly.DisableNewContent, nil)
					}
					if o.UpdateSchedule.Threats.Recurring.Weekly.Action != nil {
						nestedUpdateSchedule.Threats.Recurring.Weekly.Action = o.UpdateSchedule.Threats.Recurring.Weekly.Action
					}
				}
			}
		}
		if o.UpdateSchedule.WfPrivate != nil {
			nestedUpdateSchedule.WfPrivate = &UpdateScheduleWfPrivateXml{}
			if _, ok := o.Misc["UpdateScheduleWfPrivate"]; ok {
				nestedUpdateSchedule.WfPrivate.Misc = o.Misc["UpdateScheduleWfPrivate"]
			}
			if o.UpdateSchedule.WfPrivate.Recurring != nil {
				nestedUpdateSchedule.WfPrivate.Recurring = &UpdateScheduleWfPrivateRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleWfPrivateRecurring"]; ok {
					nestedUpdateSchedule.WfPrivate.Recurring.Misc = o.Misc["UpdateScheduleWfPrivateRecurring"]
				}
				if o.UpdateSchedule.WfPrivate.Recurring.SyncToPeer != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.SyncToPeer = util.YesNo(o.UpdateSchedule.WfPrivate.Recurring.SyncToPeer, nil)
				}
				if o.UpdateSchedule.WfPrivate.Recurring.None != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.None = &UpdateScheduleWfPrivateRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleWfPrivateRecurringNone"]; ok {
						nestedUpdateSchedule.WfPrivate.Recurring.None.Misc = o.Misc["UpdateScheduleWfPrivateRecurringNone"]
					}
				}
				if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins = &UpdateScheduleWfPrivateRecurringEvery15MinsXml{}
					if _, ok := o.Misc["UpdateScheduleWfPrivateRecurringEvery15Mins"]; ok {
						nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins.Misc = o.Misc["UpdateScheduleWfPrivateRecurringEvery15Mins"]
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Action != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Action
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.At != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.At
					}
				}
				if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins = &UpdateScheduleWfPrivateRecurringEvery30MinsXml{}
					if _, ok := o.Misc["UpdateScheduleWfPrivateRecurringEvery30Mins"]; ok {
						nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins.Misc = o.Misc["UpdateScheduleWfPrivateRecurringEvery30Mins"]
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Action != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Action
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.At != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.At
					}
				}
				if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins = &UpdateScheduleWfPrivateRecurringEvery5MinsXml{}
					if _, ok := o.Misc["UpdateScheduleWfPrivateRecurringEvery5Mins"]; ok {
						nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins.Misc = o.Misc["UpdateScheduleWfPrivateRecurringEvery5Mins"]
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Action != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Action
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.At != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.At
					}
				}
				if o.UpdateSchedule.WfPrivate.Recurring.EveryHour != nil {
					nestedUpdateSchedule.WfPrivate.Recurring.EveryHour = &UpdateScheduleWfPrivateRecurringEveryHourXml{}
					if _, ok := o.Misc["UpdateScheduleWfPrivateRecurringEveryHour"]; ok {
						nestedUpdateSchedule.WfPrivate.Recurring.EveryHour.Misc = o.Misc["UpdateScheduleWfPrivateRecurringEveryHour"]
					}
					if o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Action != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.EveryHour.Action = o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Action
					}
					if o.UpdateSchedule.WfPrivate.Recurring.EveryHour.At != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.EveryHour.At = o.UpdateSchedule.WfPrivate.Recurring.EveryHour.At
					}
				}
			}
		}
		if o.UpdateSchedule.Wildfire != nil {
			nestedUpdateSchedule.Wildfire = &UpdateScheduleWildfireXml{}
			if _, ok := o.Misc["UpdateScheduleWildfire"]; ok {
				nestedUpdateSchedule.Wildfire.Misc = o.Misc["UpdateScheduleWildfire"]
			}
			if o.UpdateSchedule.Wildfire.Recurring != nil {
				nestedUpdateSchedule.Wildfire.Recurring = &UpdateScheduleWildfireRecurringXml{}
				if _, ok := o.Misc["UpdateScheduleWildfireRecurring"]; ok {
					nestedUpdateSchedule.Wildfire.Recurring.Misc = o.Misc["UpdateScheduleWildfireRecurring"]
				}
				if o.UpdateSchedule.Wildfire.Recurring.Every15Mins != nil {
					nestedUpdateSchedule.Wildfire.Recurring.Every15Mins = &UpdateScheduleWildfireRecurringEvery15MinsXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringEvery15Mins"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.Misc = o.Misc["UpdateScheduleWildfireRecurringEvery15Mins"]
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.At != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.At = o.UpdateSchedule.Wildfire.Recurring.Every15Mins.At
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer = util.YesNo(o.UpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer, nil)
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Action != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.Action = o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Action
					}
				}
				if o.UpdateSchedule.Wildfire.Recurring.Every30Mins != nil {
					nestedUpdateSchedule.Wildfire.Recurring.Every30Mins = &UpdateScheduleWildfireRecurringEvery30MinsXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringEvery30Mins"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.Misc = o.Misc["UpdateScheduleWildfireRecurringEvery30Mins"]
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Action != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.Action = o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Action
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.At != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.At = o.UpdateSchedule.Wildfire.Recurring.Every30Mins.At
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer = util.YesNo(o.UpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer, nil)
					}
				}
				if o.UpdateSchedule.Wildfire.Recurring.EveryHour != nil {
					nestedUpdateSchedule.Wildfire.Recurring.EveryHour = &UpdateScheduleWildfireRecurringEveryHourXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringEveryHour"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.EveryHour.Misc = o.Misc["UpdateScheduleWildfireRecurringEveryHour"]
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryHour.Action != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryHour.Action = o.UpdateSchedule.Wildfire.Recurring.EveryHour.Action
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryHour.At != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryHour.At = o.UpdateSchedule.Wildfire.Recurring.EveryHour.At
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer = util.YesNo(o.UpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer, nil)
					}
				}
				if o.UpdateSchedule.Wildfire.Recurring.EveryMin != nil {
					nestedUpdateSchedule.Wildfire.Recurring.EveryMin = &UpdateScheduleWildfireRecurringEveryMinXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringEveryMin"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.EveryMin.Misc = o.Misc["UpdateScheduleWildfireRecurringEveryMin"]
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer = util.YesNo(o.UpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer, nil)
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryMin.Action != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryMin.Action = o.UpdateSchedule.Wildfire.Recurring.EveryMin.Action
					}
				}
				if o.UpdateSchedule.Wildfire.Recurring.None != nil {
					nestedUpdateSchedule.Wildfire.Recurring.None = &UpdateScheduleWildfireRecurringNoneXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringNone"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.None.Misc = o.Misc["UpdateScheduleWildfireRecurringNone"]
					}
				}
				if o.UpdateSchedule.Wildfire.Recurring.RealTime != nil {
					nestedUpdateSchedule.Wildfire.Recurring.RealTime = &UpdateScheduleWildfireRecurringRealTimeXml{}
					if _, ok := o.Misc["UpdateScheduleWildfireRecurringRealTime"]; ok {
						nestedUpdateSchedule.Wildfire.Recurring.RealTime.Misc = o.Misc["UpdateScheduleWildfireRecurringRealTime"]
					}
				}
			}
		}
	}
	config.UpdateSchedule = nestedUpdateSchedule

	config.Misc = o.Misc["Config"]

	return config, nil
}
func (c *configXmlContainer) Normalize() ([]*Config, error) {
	configList := make([]*Config, 0, len(c.Answer))
	for _, o := range c.Answer {
		config := &Config{
			Misc: make(map[string][]generic.Xml),
		}
		var nestedUpdateSchedule *UpdateSchedule
		if o.UpdateSchedule != nil {
			nestedUpdateSchedule = &UpdateSchedule{}
			if o.UpdateSchedule.Misc != nil {
				config.Misc["UpdateSchedule"] = o.UpdateSchedule.Misc
			}
			if o.UpdateSchedule.GlobalProtectClientlessVpn != nil {
				nestedUpdateSchedule.GlobalProtectClientlessVpn = &UpdateScheduleGlobalProtectClientlessVpn{}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Misc != nil {
					config.Misc["UpdateScheduleGlobalProtectClientlessVpn"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Misc
				}
				if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring != nil {
					nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring = &UpdateScheduleGlobalProtectClientlessVpnRecurring{}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Misc != nil {
						config.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurring"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Misc
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily = &UpdateScheduleGlobalProtectClientlessVpnRecurringDaily{}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringDaily"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Misc
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.Action
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Daily.At
						}
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly = &UpdateScheduleGlobalProtectClientlessVpnRecurringHourly{}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringHourly"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Misc
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.Action
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Hourly.At
						}
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.None != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.None = &UpdateScheduleGlobalProtectClientlessVpnRecurringNone{}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringNone"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly != nil {
						nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly = &UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly{}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly"] = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Misc
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.At
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.DayOfWeek
						}
						if o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action != nil {
							nestedUpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action = o.UpdateSchedule.GlobalProtectClientlessVpn.Recurring.Weekly.Action
						}
					}
				}
			}
			if o.UpdateSchedule.GlobalProtectDatafile != nil {
				nestedUpdateSchedule.GlobalProtectDatafile = &UpdateScheduleGlobalProtectDatafile{}
				if o.UpdateSchedule.GlobalProtectDatafile.Misc != nil {
					config.Misc["UpdateScheduleGlobalProtectDatafile"] = o.UpdateSchedule.GlobalProtectDatafile.Misc
				}
				if o.UpdateSchedule.GlobalProtectDatafile.Recurring != nil {
					nestedUpdateSchedule.GlobalProtectDatafile.Recurring = &UpdateScheduleGlobalProtectDatafileRecurring{}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Misc != nil {
						config.Misc["UpdateScheduleGlobalProtectDatafileRecurring"] = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Misc
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily = &UpdateScheduleGlobalProtectDatafileRecurringDaily{}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectDatafileRecurringDaily"] = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Misc
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.Action
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Daily.At
						}
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly = &UpdateScheduleGlobalProtectDatafileRecurringHourly{}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectDatafileRecurringHourly"] = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Misc
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.At
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Hourly.Action
						}
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.None != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.None = &UpdateScheduleGlobalProtectDatafileRecurringNone{}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectDatafileRecurringNone"] = o.UpdateSchedule.GlobalProtectDatafile.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly != nil {
						nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly = &UpdateScheduleGlobalProtectDatafileRecurringWeekly{}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Misc != nil {
							config.Misc["UpdateScheduleGlobalProtectDatafileRecurringWeekly"] = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Misc
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.At
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.DayOfWeek
						}
						if o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action != nil {
							nestedUpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action = o.UpdateSchedule.GlobalProtectDatafile.Recurring.Weekly.Action
						}
					}
				}
			}
			if o.UpdateSchedule.StatisticsService != nil {
				nestedUpdateSchedule.StatisticsService = &UpdateScheduleStatisticsService{}
				if o.UpdateSchedule.StatisticsService.Misc != nil {
					config.Misc["UpdateScheduleStatisticsService"] = o.UpdateSchedule.StatisticsService.Misc
				}
				if o.UpdateSchedule.StatisticsService.HealthPerformanceReports != nil {
					nestedUpdateSchedule.StatisticsService.HealthPerformanceReports = util.AsBool(o.UpdateSchedule.StatisticsService.HealthPerformanceReports, nil)
				}
				if o.UpdateSchedule.StatisticsService.PassiveDnsMonitoring != nil {
					nestedUpdateSchedule.StatisticsService.PassiveDnsMonitoring = util.AsBool(o.UpdateSchedule.StatisticsService.PassiveDnsMonitoring, nil)
				}
				if o.UpdateSchedule.StatisticsService.ThreatPreventionInformation != nil {
					nestedUpdateSchedule.StatisticsService.ThreatPreventionInformation = util.AsBool(o.UpdateSchedule.StatisticsService.ThreatPreventionInformation, nil)
				}
				if o.UpdateSchedule.StatisticsService.ThreatPreventionPcap != nil {
					nestedUpdateSchedule.StatisticsService.ThreatPreventionPcap = util.AsBool(o.UpdateSchedule.StatisticsService.ThreatPreventionPcap, nil)
				}
				if o.UpdateSchedule.StatisticsService.ThreatPreventionReports != nil {
					nestedUpdateSchedule.StatisticsService.ThreatPreventionReports = util.AsBool(o.UpdateSchedule.StatisticsService.ThreatPreventionReports, nil)
				}
				if o.UpdateSchedule.StatisticsService.UrlReports != nil {
					nestedUpdateSchedule.StatisticsService.UrlReports = util.AsBool(o.UpdateSchedule.StatisticsService.UrlReports, nil)
				}
				if o.UpdateSchedule.StatisticsService.ApplicationReports != nil {
					nestedUpdateSchedule.StatisticsService.ApplicationReports = util.AsBool(o.UpdateSchedule.StatisticsService.ApplicationReports, nil)
				}
				if o.UpdateSchedule.StatisticsService.FileIdentificationReports != nil {
					nestedUpdateSchedule.StatisticsService.FileIdentificationReports = util.AsBool(o.UpdateSchedule.StatisticsService.FileIdentificationReports, nil)
				}
			}
			if o.UpdateSchedule.Threats != nil {
				nestedUpdateSchedule.Threats = &UpdateScheduleThreats{}
				if o.UpdateSchedule.Threats.Misc != nil {
					config.Misc["UpdateScheduleThreats"] = o.UpdateSchedule.Threats.Misc
				}
				if o.UpdateSchedule.Threats.Recurring != nil {
					nestedUpdateSchedule.Threats.Recurring = &UpdateScheduleThreatsRecurring{}
					if o.UpdateSchedule.Threats.Recurring.Misc != nil {
						config.Misc["UpdateScheduleThreatsRecurring"] = o.UpdateSchedule.Threats.Recurring.Misc
					}
					if o.UpdateSchedule.Threats.Recurring.Threshold != nil {
						nestedUpdateSchedule.Threats.Recurring.Threshold = o.UpdateSchedule.Threats.Recurring.Threshold
					}
					if o.UpdateSchedule.Threats.Recurring.NewAppThreshold != nil {
						nestedUpdateSchedule.Threats.Recurring.NewAppThreshold = o.UpdateSchedule.Threats.Recurring.NewAppThreshold
					}
					if o.UpdateSchedule.Threats.Recurring.SyncToPeer != nil {
						nestedUpdateSchedule.Threats.Recurring.SyncToPeer = util.AsBool(o.UpdateSchedule.Threats.Recurring.SyncToPeer, nil)
					}
					if o.UpdateSchedule.Threats.Recurring.Daily != nil {
						nestedUpdateSchedule.Threats.Recurring.Daily = &UpdateScheduleThreatsRecurringDaily{}
						if o.UpdateSchedule.Threats.Recurring.Daily.Misc != nil {
							config.Misc["UpdateScheduleThreatsRecurringDaily"] = o.UpdateSchedule.Threats.Recurring.Daily.Misc
						}
						if o.UpdateSchedule.Threats.Recurring.Daily.Action != nil {
							nestedUpdateSchedule.Threats.Recurring.Daily.Action = o.UpdateSchedule.Threats.Recurring.Daily.Action
						}
						if o.UpdateSchedule.Threats.Recurring.Daily.At != nil {
							nestedUpdateSchedule.Threats.Recurring.Daily.At = o.UpdateSchedule.Threats.Recurring.Daily.At
						}
						if o.UpdateSchedule.Threats.Recurring.Daily.DisableNewContent != nil {
							nestedUpdateSchedule.Threats.Recurring.Daily.DisableNewContent = util.AsBool(o.UpdateSchedule.Threats.Recurring.Daily.DisableNewContent, nil)
						}
					}
					if o.UpdateSchedule.Threats.Recurring.Every30Mins != nil {
						nestedUpdateSchedule.Threats.Recurring.Every30Mins = &UpdateScheduleThreatsRecurringEvery30Mins{}
						if o.UpdateSchedule.Threats.Recurring.Every30Mins.Misc != nil {
							config.Misc["UpdateScheduleThreatsRecurringEvery30Mins"] = o.UpdateSchedule.Threats.Recurring.Every30Mins.Misc
						}
						if o.UpdateSchedule.Threats.Recurring.Every30Mins.Action != nil {
							nestedUpdateSchedule.Threats.Recurring.Every30Mins.Action = o.UpdateSchedule.Threats.Recurring.Every30Mins.Action
						}
						if o.UpdateSchedule.Threats.Recurring.Every30Mins.At != nil {
							nestedUpdateSchedule.Threats.Recurring.Every30Mins.At = o.UpdateSchedule.Threats.Recurring.Every30Mins.At
						}
						if o.UpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent != nil {
							nestedUpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent = util.AsBool(o.UpdateSchedule.Threats.Recurring.Every30Mins.DisableNewContent, nil)
						}
					}
					if o.UpdateSchedule.Threats.Recurring.Hourly != nil {
						nestedUpdateSchedule.Threats.Recurring.Hourly = &UpdateScheduleThreatsRecurringHourly{}
						if o.UpdateSchedule.Threats.Recurring.Hourly.Misc != nil {
							config.Misc["UpdateScheduleThreatsRecurringHourly"] = o.UpdateSchedule.Threats.Recurring.Hourly.Misc
						}
						if o.UpdateSchedule.Threats.Recurring.Hourly.Action != nil {
							nestedUpdateSchedule.Threats.Recurring.Hourly.Action = o.UpdateSchedule.Threats.Recurring.Hourly.Action
						}
						if o.UpdateSchedule.Threats.Recurring.Hourly.At != nil {
							nestedUpdateSchedule.Threats.Recurring.Hourly.At = o.UpdateSchedule.Threats.Recurring.Hourly.At
						}
						if o.UpdateSchedule.Threats.Recurring.Hourly.DisableNewContent != nil {
							nestedUpdateSchedule.Threats.Recurring.Hourly.DisableNewContent = util.AsBool(o.UpdateSchedule.Threats.Recurring.Hourly.DisableNewContent, nil)
						}
					}
					if o.UpdateSchedule.Threats.Recurring.None != nil {
						nestedUpdateSchedule.Threats.Recurring.None = &UpdateScheduleThreatsRecurringNone{}
						if o.UpdateSchedule.Threats.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleThreatsRecurringNone"] = o.UpdateSchedule.Threats.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.Threats.Recurring.Weekly != nil {
						nestedUpdateSchedule.Threats.Recurring.Weekly = &UpdateScheduleThreatsRecurringWeekly{}
						if o.UpdateSchedule.Threats.Recurring.Weekly.Misc != nil {
							config.Misc["UpdateScheduleThreatsRecurringWeekly"] = o.UpdateSchedule.Threats.Recurring.Weekly.Misc
						}
						if o.UpdateSchedule.Threats.Recurring.Weekly.Action != nil {
							nestedUpdateSchedule.Threats.Recurring.Weekly.Action = o.UpdateSchedule.Threats.Recurring.Weekly.Action
						}
						if o.UpdateSchedule.Threats.Recurring.Weekly.At != nil {
							nestedUpdateSchedule.Threats.Recurring.Weekly.At = o.UpdateSchedule.Threats.Recurring.Weekly.At
						}
						if o.UpdateSchedule.Threats.Recurring.Weekly.DayOfWeek != nil {
							nestedUpdateSchedule.Threats.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.Threats.Recurring.Weekly.DayOfWeek
						}
						if o.UpdateSchedule.Threats.Recurring.Weekly.DisableNewContent != nil {
							nestedUpdateSchedule.Threats.Recurring.Weekly.DisableNewContent = util.AsBool(o.UpdateSchedule.Threats.Recurring.Weekly.DisableNewContent, nil)
						}
					}
				}
			}
			if o.UpdateSchedule.WfPrivate != nil {
				nestedUpdateSchedule.WfPrivate = &UpdateScheduleWfPrivate{}
				if o.UpdateSchedule.WfPrivate.Misc != nil {
					config.Misc["UpdateScheduleWfPrivate"] = o.UpdateSchedule.WfPrivate.Misc
				}
				if o.UpdateSchedule.WfPrivate.Recurring != nil {
					nestedUpdateSchedule.WfPrivate.Recurring = &UpdateScheduleWfPrivateRecurring{}
					if o.UpdateSchedule.WfPrivate.Recurring.Misc != nil {
						config.Misc["UpdateScheduleWfPrivateRecurring"] = o.UpdateSchedule.WfPrivate.Recurring.Misc
					}
					if o.UpdateSchedule.WfPrivate.Recurring.SyncToPeer != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.SyncToPeer = util.AsBool(o.UpdateSchedule.WfPrivate.Recurring.SyncToPeer, nil)
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins = &UpdateScheduleWfPrivateRecurringEvery15Mins{}
						if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Misc != nil {
							config.Misc["UpdateScheduleWfPrivateRecurringEvery15Mins"] = o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Misc
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.At != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.At
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Action != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every15Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every15Mins.Action
						}
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins = &UpdateScheduleWfPrivateRecurringEvery30Mins{}
						if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Misc != nil {
							config.Misc["UpdateScheduleWfPrivateRecurringEvery30Mins"] = o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Misc
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Action != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.Action
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.At != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every30Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every30Mins.At
						}
					}
					if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins = &UpdateScheduleWfPrivateRecurringEvery5Mins{}
						if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Misc != nil {
							config.Misc["UpdateScheduleWfPrivateRecurringEvery5Mins"] = o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Misc
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Action != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins.Action = o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.Action
						}
						if o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.At != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.Every5Mins.At = o.UpdateSchedule.WfPrivate.Recurring.Every5Mins.At
						}
					}
					if o.UpdateSchedule.WfPrivate.Recurring.EveryHour != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.EveryHour = &UpdateScheduleWfPrivateRecurringEveryHour{}
						if o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Misc != nil {
							config.Misc["UpdateScheduleWfPrivateRecurringEveryHour"] = o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Misc
						}
						if o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Action != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.EveryHour.Action = o.UpdateSchedule.WfPrivate.Recurring.EveryHour.Action
						}
						if o.UpdateSchedule.WfPrivate.Recurring.EveryHour.At != nil {
							nestedUpdateSchedule.WfPrivate.Recurring.EveryHour.At = o.UpdateSchedule.WfPrivate.Recurring.EveryHour.At
						}
					}
					if o.UpdateSchedule.WfPrivate.Recurring.None != nil {
						nestedUpdateSchedule.WfPrivate.Recurring.None = &UpdateScheduleWfPrivateRecurringNone{}
						if o.UpdateSchedule.WfPrivate.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleWfPrivateRecurringNone"] = o.UpdateSchedule.WfPrivate.Recurring.None.Misc
						}
					}
				}
			}
			if o.UpdateSchedule.Wildfire != nil {
				nestedUpdateSchedule.Wildfire = &UpdateScheduleWildfire{}
				if o.UpdateSchedule.Wildfire.Misc != nil {
					config.Misc["UpdateScheduleWildfire"] = o.UpdateSchedule.Wildfire.Misc
				}
				if o.UpdateSchedule.Wildfire.Recurring != nil {
					nestedUpdateSchedule.Wildfire.Recurring = &UpdateScheduleWildfireRecurring{}
					if o.UpdateSchedule.Wildfire.Recurring.Misc != nil {
						config.Misc["UpdateScheduleWildfireRecurring"] = o.UpdateSchedule.Wildfire.Recurring.Misc
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every30Mins != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every30Mins = &UpdateScheduleWildfireRecurringEvery30Mins{}
						if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringEvery30Mins"] = o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Misc
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.At != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.At = o.UpdateSchedule.Wildfire.Recurring.Every30Mins.At
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer = util.AsBool(o.UpdateSchedule.Wildfire.Recurring.Every30Mins.SyncToPeer, nil)
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Action != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every30Mins.Action = o.UpdateSchedule.Wildfire.Recurring.Every30Mins.Action
						}
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryHour != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryHour = &UpdateScheduleWildfireRecurringEveryHour{}
						if o.UpdateSchedule.Wildfire.Recurring.EveryHour.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringEveryHour"] = o.UpdateSchedule.Wildfire.Recurring.EveryHour.Misc
						}
						if o.UpdateSchedule.Wildfire.Recurring.EveryHour.Action != nil {
							nestedUpdateSchedule.Wildfire.Recurring.EveryHour.Action = o.UpdateSchedule.Wildfire.Recurring.EveryHour.Action
						}
						if o.UpdateSchedule.Wildfire.Recurring.EveryHour.At != nil {
							nestedUpdateSchedule.Wildfire.Recurring.EveryHour.At = o.UpdateSchedule.Wildfire.Recurring.EveryHour.At
						}
						if o.UpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer != nil {
							nestedUpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer = util.AsBool(o.UpdateSchedule.Wildfire.Recurring.EveryHour.SyncToPeer, nil)
						}
					}
					if o.UpdateSchedule.Wildfire.Recurring.EveryMin != nil {
						nestedUpdateSchedule.Wildfire.Recurring.EveryMin = &UpdateScheduleWildfireRecurringEveryMin{}
						if o.UpdateSchedule.Wildfire.Recurring.EveryMin.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringEveryMin"] = o.UpdateSchedule.Wildfire.Recurring.EveryMin.Misc
						}
						if o.UpdateSchedule.Wildfire.Recurring.EveryMin.Action != nil {
							nestedUpdateSchedule.Wildfire.Recurring.EveryMin.Action = o.UpdateSchedule.Wildfire.Recurring.EveryMin.Action
						}
						if o.UpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer != nil {
							nestedUpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer = util.AsBool(o.UpdateSchedule.Wildfire.Recurring.EveryMin.SyncToPeer, nil)
						}
					}
					if o.UpdateSchedule.Wildfire.Recurring.None != nil {
						nestedUpdateSchedule.Wildfire.Recurring.None = &UpdateScheduleWildfireRecurringNone{}
						if o.UpdateSchedule.Wildfire.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringNone"] = o.UpdateSchedule.Wildfire.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.Wildfire.Recurring.RealTime != nil {
						nestedUpdateSchedule.Wildfire.Recurring.RealTime = &UpdateScheduleWildfireRecurringRealTime{}
						if o.UpdateSchedule.Wildfire.Recurring.RealTime.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringRealTime"] = o.UpdateSchedule.Wildfire.Recurring.RealTime.Misc
						}
					}
					if o.UpdateSchedule.Wildfire.Recurring.Every15Mins != nil {
						nestedUpdateSchedule.Wildfire.Recurring.Every15Mins = &UpdateScheduleWildfireRecurringEvery15Mins{}
						if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Misc != nil {
							config.Misc["UpdateScheduleWildfireRecurringEvery15Mins"] = o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Misc
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.At != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.At = o.UpdateSchedule.Wildfire.Recurring.Every15Mins.At
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer = util.AsBool(o.UpdateSchedule.Wildfire.Recurring.Every15Mins.SyncToPeer, nil)
						}
						if o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Action != nil {
							nestedUpdateSchedule.Wildfire.Recurring.Every15Mins.Action = o.UpdateSchedule.Wildfire.Recurring.Every15Mins.Action
						}
					}
				}
			}
			if o.UpdateSchedule.AntiVirus != nil {
				nestedUpdateSchedule.AntiVirus = &UpdateScheduleAntiVirus{}
				if o.UpdateSchedule.AntiVirus.Misc != nil {
					config.Misc["UpdateScheduleAntiVirus"] = o.UpdateSchedule.AntiVirus.Misc
				}
				if o.UpdateSchedule.AntiVirus.Recurring != nil {
					nestedUpdateSchedule.AntiVirus.Recurring = &UpdateScheduleAntiVirusRecurring{}
					if o.UpdateSchedule.AntiVirus.Recurring.Misc != nil {
						config.Misc["UpdateScheduleAntiVirusRecurring"] = o.UpdateSchedule.AntiVirus.Recurring.Misc
					}
					if o.UpdateSchedule.AntiVirus.Recurring.SyncToPeer != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.SyncToPeer = util.AsBool(o.UpdateSchedule.AntiVirus.Recurring.SyncToPeer, nil)
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Threshold != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Threshold = o.UpdateSchedule.AntiVirus.Recurring.Threshold
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Daily != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Daily = &UpdateScheduleAntiVirusRecurringDaily{}
						if o.UpdateSchedule.AntiVirus.Recurring.Daily.Misc != nil {
							config.Misc["UpdateScheduleAntiVirusRecurringDaily"] = o.UpdateSchedule.AntiVirus.Recurring.Daily.Misc
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Daily.Action != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Daily.Action = o.UpdateSchedule.AntiVirus.Recurring.Daily.Action
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Daily.At != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Daily.At = o.UpdateSchedule.AntiVirus.Recurring.Daily.At
						}
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Hourly != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Hourly = &UpdateScheduleAntiVirusRecurringHourly{}
						if o.UpdateSchedule.AntiVirus.Recurring.Hourly.Misc != nil {
							config.Misc["UpdateScheduleAntiVirusRecurringHourly"] = o.UpdateSchedule.AntiVirus.Recurring.Hourly.Misc
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Hourly.Action != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Hourly.Action = o.UpdateSchedule.AntiVirus.Recurring.Hourly.Action
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Hourly.At != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Hourly.At = o.UpdateSchedule.AntiVirus.Recurring.Hourly.At
						}
					}
					if o.UpdateSchedule.AntiVirus.Recurring.None != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.None = &UpdateScheduleAntiVirusRecurringNone{}
						if o.UpdateSchedule.AntiVirus.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleAntiVirusRecurringNone"] = o.UpdateSchedule.AntiVirus.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.AntiVirus.Recurring.Weekly != nil {
						nestedUpdateSchedule.AntiVirus.Recurring.Weekly = &UpdateScheduleAntiVirusRecurringWeekly{}
						if o.UpdateSchedule.AntiVirus.Recurring.Weekly.Misc != nil {
							config.Misc["UpdateScheduleAntiVirusRecurringWeekly"] = o.UpdateSchedule.AntiVirus.Recurring.Weekly.Misc
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Weekly.Action != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Weekly.Action = o.UpdateSchedule.AntiVirus.Recurring.Weekly.Action
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Weekly.At != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Weekly.At = o.UpdateSchedule.AntiVirus.Recurring.Weekly.At
						}
						if o.UpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek != nil {
							nestedUpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.AntiVirus.Recurring.Weekly.DayOfWeek
						}
					}
				}
			}
			if o.UpdateSchedule.AppProfile != nil {
				nestedUpdateSchedule.AppProfile = &UpdateScheduleAppProfile{}
				if o.UpdateSchedule.AppProfile.Misc != nil {
					config.Misc["UpdateScheduleAppProfile"] = o.UpdateSchedule.AppProfile.Misc
				}
				if o.UpdateSchedule.AppProfile.Recurring != nil {
					nestedUpdateSchedule.AppProfile.Recurring = &UpdateScheduleAppProfileRecurring{}
					if o.UpdateSchedule.AppProfile.Recurring.Misc != nil {
						config.Misc["UpdateScheduleAppProfileRecurring"] = o.UpdateSchedule.AppProfile.Recurring.Misc
					}
					if o.UpdateSchedule.AppProfile.Recurring.SyncToPeer != nil {
						nestedUpdateSchedule.AppProfile.Recurring.SyncToPeer = util.AsBool(o.UpdateSchedule.AppProfile.Recurring.SyncToPeer, nil)
					}
					if o.UpdateSchedule.AppProfile.Recurring.Threshold != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Threshold = o.UpdateSchedule.AppProfile.Recurring.Threshold
					}
					if o.UpdateSchedule.AppProfile.Recurring.Daily != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Daily = &UpdateScheduleAppProfileRecurringDaily{}
						if o.UpdateSchedule.AppProfile.Recurring.Daily.Misc != nil {
							config.Misc["UpdateScheduleAppProfileRecurringDaily"] = o.UpdateSchedule.AppProfile.Recurring.Daily.Misc
						}
						if o.UpdateSchedule.AppProfile.Recurring.Daily.Action != nil {
							nestedUpdateSchedule.AppProfile.Recurring.Daily.Action = o.UpdateSchedule.AppProfile.Recurring.Daily.Action
						}
						if o.UpdateSchedule.AppProfile.Recurring.Daily.At != nil {
							nestedUpdateSchedule.AppProfile.Recurring.Daily.At = o.UpdateSchedule.AppProfile.Recurring.Daily.At
						}
					}
					if o.UpdateSchedule.AppProfile.Recurring.None != nil {
						nestedUpdateSchedule.AppProfile.Recurring.None = &UpdateScheduleAppProfileRecurringNone{}
						if o.UpdateSchedule.AppProfile.Recurring.None.Misc != nil {
							config.Misc["UpdateScheduleAppProfileRecurringNone"] = o.UpdateSchedule.AppProfile.Recurring.None.Misc
						}
					}
					if o.UpdateSchedule.AppProfile.Recurring.Weekly != nil {
						nestedUpdateSchedule.AppProfile.Recurring.Weekly = &UpdateScheduleAppProfileRecurringWeekly{}
						if o.UpdateSchedule.AppProfile.Recurring.Weekly.Misc != nil {
							config.Misc["UpdateScheduleAppProfileRecurringWeekly"] = o.UpdateSchedule.AppProfile.Recurring.Weekly.Misc
						}
						if o.UpdateSchedule.AppProfile.Recurring.Weekly.Action != nil {
							nestedUpdateSchedule.AppProfile.Recurring.Weekly.Action = o.UpdateSchedule.AppProfile.Recurring.Weekly.Action
						}
						if o.UpdateSchedule.AppProfile.Recurring.Weekly.At != nil {
							nestedUpdateSchedule.AppProfile.Recurring.Weekly.At = o.UpdateSchedule.AppProfile.Recurring.Weekly.At
						}
						if o.UpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek != nil {
							nestedUpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek = o.UpdateSchedule.AppProfile.Recurring.Weekly.DayOfWeek
						}
					}
				}
			}
		}
		config.UpdateSchedule = nestedUpdateSchedule

		config.Misc["Config"] = o.Misc

		configList = append(configList, config)
	}

	return configList, nil
}

func SpecMatches(a, b *Config) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}

	// Don't compare Name.
	if !matchUpdateSchedule(a.UpdateSchedule, b.UpdateSchedule) {
		return false
	}

	return true
}

func matchUpdateScheduleThreatsRecurringDaily(a *UpdateScheduleThreatsRecurringDaily, b *UpdateScheduleThreatsRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.DisableNewContent, b.DisableNewContent) {
		return false
	}
	return true
}
func matchUpdateScheduleThreatsRecurringEvery30Mins(a *UpdateScheduleThreatsRecurringEvery30Mins, b *UpdateScheduleThreatsRecurringEvery30Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.DisableNewContent, b.DisableNewContent) {
		return false
	}
	return true
}
func matchUpdateScheduleThreatsRecurringHourly(a *UpdateScheduleThreatsRecurringHourly, b *UpdateScheduleThreatsRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.DisableNewContent, b.DisableNewContent) {
		return false
	}
	return true
}
func matchUpdateScheduleThreatsRecurringNone(a *UpdateScheduleThreatsRecurringNone, b *UpdateScheduleThreatsRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleThreatsRecurringWeekly(a *UpdateScheduleThreatsRecurringWeekly, b *UpdateScheduleThreatsRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	if !util.BoolsMatch(a.DisableNewContent, b.DisableNewContent) {
		return false
	}
	return true
}
func matchUpdateScheduleThreatsRecurring(a *UpdateScheduleThreatsRecurring, b *UpdateScheduleThreatsRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.NewAppThreshold, b.NewAppThreshold) {
		return false
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(a.Threshold, b.Threshold) {
		return false
	}
	if !matchUpdateScheduleThreatsRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchUpdateScheduleThreatsRecurringEvery30Mins(a.Every30Mins, b.Every30Mins) {
		return false
	}
	if !matchUpdateScheduleThreatsRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchUpdateScheduleThreatsRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleThreatsRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	return true
}
func matchUpdateScheduleThreats(a *UpdateScheduleThreats, b *UpdateScheduleThreats) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleThreatsRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurringEvery5Mins(a *UpdateScheduleWfPrivateRecurringEvery5Mins, b *UpdateScheduleWfPrivateRecurringEvery5Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurringEveryHour(a *UpdateScheduleWfPrivateRecurringEveryHour, b *UpdateScheduleWfPrivateRecurringEveryHour) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurringNone(a *UpdateScheduleWfPrivateRecurringNone, b *UpdateScheduleWfPrivateRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurringEvery15Mins(a *UpdateScheduleWfPrivateRecurringEvery15Mins, b *UpdateScheduleWfPrivateRecurringEvery15Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurringEvery30Mins(a *UpdateScheduleWfPrivateRecurringEvery30Mins, b *UpdateScheduleWfPrivateRecurringEvery30Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivateRecurring(a *UpdateScheduleWfPrivateRecurring, b *UpdateScheduleWfPrivateRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	if !matchUpdateScheduleWfPrivateRecurringEvery30Mins(a.Every30Mins, b.Every30Mins) {
		return false
	}
	if !matchUpdateScheduleWfPrivateRecurringEvery5Mins(a.Every5Mins, b.Every5Mins) {
		return false
	}
	if !matchUpdateScheduleWfPrivateRecurringEveryHour(a.EveryHour, b.EveryHour) {
		return false
	}
	if !matchUpdateScheduleWfPrivateRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleWfPrivateRecurringEvery15Mins(a.Every15Mins, b.Every15Mins) {
		return false
	}
	return true
}
func matchUpdateScheduleWfPrivate(a *UpdateScheduleWfPrivate, b *UpdateScheduleWfPrivate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleWfPrivateRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfireRecurringEvery15Mins(a *UpdateScheduleWildfireRecurringEvery15Mins, b *UpdateScheduleWildfireRecurringEvery15Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfireRecurringEvery30Mins(a *UpdateScheduleWildfireRecurringEvery30Mins, b *UpdateScheduleWildfireRecurringEvery30Mins) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfireRecurringEveryHour(a *UpdateScheduleWildfireRecurringEveryHour, b *UpdateScheduleWildfireRecurringEveryHour) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfireRecurringEveryMin(a *UpdateScheduleWildfireRecurringEveryMin, b *UpdateScheduleWildfireRecurringEveryMin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfireRecurringNone(a *UpdateScheduleWildfireRecurringNone, b *UpdateScheduleWildfireRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleWildfireRecurringRealTime(a *UpdateScheduleWildfireRecurringRealTime, b *UpdateScheduleWildfireRecurringRealTime) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleWildfireRecurring(a *UpdateScheduleWildfireRecurring, b *UpdateScheduleWildfireRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleWildfireRecurringEveryMin(a.EveryMin, b.EveryMin) {
		return false
	}
	if !matchUpdateScheduleWildfireRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleWildfireRecurringRealTime(a.RealTime, b.RealTime) {
		return false
	}
	if !matchUpdateScheduleWildfireRecurringEvery15Mins(a.Every15Mins, b.Every15Mins) {
		return false
	}
	if !matchUpdateScheduleWildfireRecurringEvery30Mins(a.Every30Mins, b.Every30Mins) {
		return false
	}
	if !matchUpdateScheduleWildfireRecurringEveryHour(a.EveryHour, b.EveryHour) {
		return false
	}
	return true
}
func matchUpdateScheduleWildfire(a *UpdateScheduleWildfire, b *UpdateScheduleWildfire) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleWildfireRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleAntiVirusRecurringDaily(a *UpdateScheduleAntiVirusRecurringDaily, b *UpdateScheduleAntiVirusRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleAntiVirusRecurringHourly(a *UpdateScheduleAntiVirusRecurringHourly, b *UpdateScheduleAntiVirusRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleAntiVirusRecurringNone(a *UpdateScheduleAntiVirusRecurringNone, b *UpdateScheduleAntiVirusRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleAntiVirusRecurringWeekly(a *UpdateScheduleAntiVirusRecurringWeekly, b *UpdateScheduleAntiVirusRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchUpdateScheduleAntiVirusRecurring(a *UpdateScheduleAntiVirusRecurring, b *UpdateScheduleAntiVirusRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(a.Threshold, b.Threshold) {
		return false
	}
	if !matchUpdateScheduleAntiVirusRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleAntiVirusRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	if !matchUpdateScheduleAntiVirusRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchUpdateScheduleAntiVirusRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	return true
}
func matchUpdateScheduleAntiVirus(a *UpdateScheduleAntiVirus, b *UpdateScheduleAntiVirus) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleAntiVirusRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleAppProfileRecurringNone(a *UpdateScheduleAppProfileRecurringNone, b *UpdateScheduleAppProfileRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleAppProfileRecurringWeekly(a *UpdateScheduleAppProfileRecurringWeekly, b *UpdateScheduleAppProfileRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchUpdateScheduleAppProfileRecurringDaily(a *UpdateScheduleAppProfileRecurringDaily, b *UpdateScheduleAppProfileRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleAppProfileRecurring(a *UpdateScheduleAppProfileRecurring, b *UpdateScheduleAppProfileRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.SyncToPeer, b.SyncToPeer) {
		return false
	}
	if !util.Ints64Match(a.Threshold, b.Threshold) {
		return false
	}
	if !matchUpdateScheduleAppProfileRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchUpdateScheduleAppProfileRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleAppProfileRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	return true
}
func matchUpdateScheduleAppProfile(a *UpdateScheduleAppProfile, b *UpdateScheduleAppProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleAppProfileRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpnRecurringDaily(a *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily, b *UpdateScheduleGlobalProtectClientlessVpnRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpnRecurringHourly(a *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly, b *UpdateScheduleGlobalProtectClientlessVpnRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpnRecurringNone(a *UpdateScheduleGlobalProtectClientlessVpnRecurringNone, b *UpdateScheduleGlobalProtectClientlessVpnRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpnRecurringWeekly(a *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly, b *UpdateScheduleGlobalProtectClientlessVpnRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpnRecurring(a *UpdateScheduleGlobalProtectClientlessVpnRecurring, b *UpdateScheduleGlobalProtectClientlessVpnRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpnRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpnRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpnRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpnRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectClientlessVpn(a *UpdateScheduleGlobalProtectClientlessVpn, b *UpdateScheduleGlobalProtectClientlessVpn) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpnRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafileRecurringNone(a *UpdateScheduleGlobalProtectDatafileRecurringNone, b *UpdateScheduleGlobalProtectDatafileRecurringNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafileRecurringWeekly(a *UpdateScheduleGlobalProtectDatafileRecurringWeekly, b *UpdateScheduleGlobalProtectDatafileRecurringWeekly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	if !util.StringsMatch(a.DayOfWeek, b.DayOfWeek) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafileRecurringDaily(a *UpdateScheduleGlobalProtectDatafileRecurringDaily, b *UpdateScheduleGlobalProtectDatafileRecurringDaily) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.StringsMatch(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafileRecurringHourly(a *UpdateScheduleGlobalProtectDatafileRecurringHourly, b *UpdateScheduleGlobalProtectDatafileRecurringHourly) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Action, b.Action) {
		return false
	}
	if !util.Ints64Match(a.At, b.At) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafileRecurring(a *UpdateScheduleGlobalProtectDatafileRecurring, b *UpdateScheduleGlobalProtectDatafileRecurring) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleGlobalProtectDatafileRecurringDaily(a.Daily, b.Daily) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectDatafileRecurringHourly(a.Hourly, b.Hourly) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectDatafileRecurringNone(a.None, b.None) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectDatafileRecurringWeekly(a.Weekly, b.Weekly) {
		return false
	}
	return true
}
func matchUpdateScheduleGlobalProtectDatafile(a *UpdateScheduleGlobalProtectDatafile, b *UpdateScheduleGlobalProtectDatafile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleGlobalProtectDatafileRecurring(a.Recurring, b.Recurring) {
		return false
	}
	return true
}
func matchUpdateScheduleStatisticsService(a *UpdateScheduleStatisticsService, b *UpdateScheduleStatisticsService) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.HealthPerformanceReports, b.HealthPerformanceReports) {
		return false
	}
	if !util.BoolsMatch(a.PassiveDnsMonitoring, b.PassiveDnsMonitoring) {
		return false
	}
	if !util.BoolsMatch(a.ThreatPreventionInformation, b.ThreatPreventionInformation) {
		return false
	}
	if !util.BoolsMatch(a.ThreatPreventionPcap, b.ThreatPreventionPcap) {
		return false
	}
	if !util.BoolsMatch(a.ThreatPreventionReports, b.ThreatPreventionReports) {
		return false
	}
	if !util.BoolsMatch(a.UrlReports, b.UrlReports) {
		return false
	}
	if !util.BoolsMatch(a.ApplicationReports, b.ApplicationReports) {
		return false
	}
	if !util.BoolsMatch(a.FileIdentificationReports, b.FileIdentificationReports) {
		return false
	}
	return true
}
func matchUpdateSchedule(a *UpdateSchedule, b *UpdateSchedule) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchUpdateScheduleThreats(a.Threats, b.Threats) {
		return false
	}
	if !matchUpdateScheduleWfPrivate(a.WfPrivate, b.WfPrivate) {
		return false
	}
	if !matchUpdateScheduleWildfire(a.Wildfire, b.Wildfire) {
		return false
	}
	if !matchUpdateScheduleAntiVirus(a.AntiVirus, b.AntiVirus) {
		return false
	}
	if !matchUpdateScheduleAppProfile(a.AppProfile, b.AppProfile) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectClientlessVpn(a.GlobalProtectClientlessVpn, b.GlobalProtectClientlessVpn) {
		return false
	}
	if !matchUpdateScheduleGlobalProtectDatafile(a.GlobalProtectDatafile, b.GlobalProtectDatafile) {
		return false
	}
	if !matchUpdateScheduleStatisticsService(a.StatisticsService, b.StatisticsService) {
		return false
	}
	return true
}
